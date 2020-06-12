package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("data/articles.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	type Article struct {
		Id int
		Slug, Title, Teaser, Date_posted, Content string
	}
	dec := json.NewDecoder(jsonFile)
	_ = os.MkdirAll("content/articles/", 0775)

	// read open bracket
		t, err := dec.Token()
		if err != nil {
			fmt.Println(err)
			fmt.Println(t)
		}

		// while the array contains values
		for dec.More() {
			var a Article
			var out = ""
			// decode an array value (Article)
			err := dec.Decode(&a)
			if err != nil {
				fmt.Println(err)
			}

			thisId := fmt.Sprintf("%v", a.Id)

			out += fmt.Sprintf("---\n")
			out += fmt.Sprintf("id: %v\n", thisId)
			out += fmt.Sprintf("title: |\n  %v\n", a.Title)
			out += fmt.Sprintf("teaser: |\n  %v\n", a.Teaser)
			out += fmt.Sprintf("date: %v\n", a.Date_posted)
			out += fmt.Sprintf("---\n")
			out += fmt.Sprintf("%v\n", a.Content)

			file, err := os.Create("content/articles/" + a.Slug + ".md")
			if err != nil {
				fmt.Println(err)
			}
			file.WriteString(out)
			file.Close()
		}
}

