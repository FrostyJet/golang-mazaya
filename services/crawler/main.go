package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang-mazaya/crawler/internal/dialer"
	"golang-mazaya/crawler/internal/models"

	"github.com/gocolly/colly"
)

const baseURL = "https://foodnetwork.co.uk"

func handler(res http.ResponseWriter, req *http.Request) {
	log.Println("=========== Start Crawl ===================")

	limit := 1

	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	c.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting url %s\n", r.URL)
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a.card-link", func(e *colly.HTMLElement) {
		if limit <= 0 {
			return
		}

		limit--
		link := e.Attr("href")

		// Visit link found on page
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Actual recipe page
	c.OnHTML("section.article-head[itemtype='https://schema.org/Recipe']", func(element *colly.HTMLElement) {
		log.Println("=========== Found recipe ===================")

		r := &models.Recipe{}

		r.Title = element.DOM.Find(".page__title").Text()

		// Attributes
		{
			head := element.DOM.Find(".recipe-head")

			// Preparation time
			prepTime := head.Find("li:first-child > strong").Text()
			prepTimeWords := strings.Fields(prepTime)
			if time, err := strconv.Atoi(prepTimeWords[0]); err == nil {
				r.PreparationTime = time
				r.PreparationTimeUnit = prepTimeWords[1]
			}

			// Cooking time
			cookingTime := head.Find("li:nth-of-type(2) > strong").Text()
			cookingTimeWords := strings.Fields(cookingTime)
			if time, err := strconv.Atoi(cookingTimeWords[0]); err == nil {
				r.CookingTime = time
				r.CookingTimeUnit = cookingTimeWords[1]
			}

			// Servings
			servingsStr := head.Find("li:nth-of-type(3) > strong").Text()
			if servingsInt, err := strconv.Atoi(servingsStr); err == nil {
				r.Serves = servingsInt
			}

			// Difficulty
			r.Difficulty = head.Find("li:nth-of-type(4) > strong").Text()
		}

		// Hero Image
		r.Poster = element.DOM.Find("meta[itemprop='image']").AttrOr("content", "")

		// Description
		r.DescriptionText = element.DOM.Find(".recipe-text").Text()
		descriptionHTML, err := element.DOM.Find(".recipe-text").Html()
		if err == nil {
			r.DescriptionHTML = descriptionHTML
		}

		// Keywords
		keywordsMeta := element.DOM.Find(".tags > meta")
		keywordsStr := keywordsMeta.AttrOr("content", "")
		keywordsFields := strings.Fields(keywordsStr)
		r.Keywords = make([]models.RecipeKeyword, 0, len(keywordsFields))
		for _, k := range keywordsFields {
			r.Keywords = append(r.Keywords, models.RecipeKeyword(k))
		}

		body, err := json.Marshal(r)
		if err != nil {
			log.Printf("Could not create json from data: %+v\n", r)
		}

		err = dialer.Publish("recipe.created", body, "application/json")
		if err != nil {
			log.Println("Could not publish generated recipe")
			log.Panicln(err.Error())
			return
		}

		log.Printf("[x] new recipe has been published to rabbitMQ: '%s'", r.Title)
	})

	c.Visit(fmt.Sprintf("%s/dessert-recipes/", baseURL))
}

func main() {
	addr := ":8080"

	http.HandleFunc("/crawl", handler)

	log.Println("Crawler service started!: listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
