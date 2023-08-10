package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var stars = ListAllStars(os.Getenv("GITHUB_USERNAME"))
	content := generateContent(stars)
	createFile(content)
}

func generateContent(stars []GithubStarred) (content string) {
	var maps = make(map[string][]GithubStarred)

	content += "## Summary\n"

	for _, star := range stars {
		var language = star.Language
		if language == "" {
			language = "Others"
		}
		maps[language] = append(maps[language], star)
	}

	content += "\n"

	for language := range maps {
		content += fmt.Sprintf("- [%s](#%s)\n", language, strings.ReplaceAll(strings.ToLower(language), " ", "-"))
	}

	for lang, stars := range maps {
		content += fmt.Sprintf("## %s\n\n", lang)

		for _, star := range stars {
			content += fmt.Sprintf("- [%s](%s) - %s\n", star.Name, star.HtmlUrl, star.Description)
		}

		content += "\n"
	}

	return content
}

func createFile(content string) {
	f, err := os.Create("README.md")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.WriteString(content)
}
