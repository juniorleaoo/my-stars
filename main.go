package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	githubUsername := flag.String("github_username", "asd", "Github Username")
	githubToken := flag.String("github_token", "asd", "Github Token")
	flag.Parse()

	githubClient := GithubClient{
		Authorization: *githubToken,
	}
	var stars = githubClient.ListAllStars(*githubUsername)
	content := generateContent(stars)
	createFile(content)
}

func generateContent(stars []GithubStarred) (content string) {
	var maps = make(map[string][]GithubStarred)
	var langs []string

	content += "## Summary\n"

	for _, star := range stars {
		var language = star.Language
		if language == "" {
			language = "Others"
		}
		if !slices.Contains(langs, language) {
			langs = append(langs, language)
		}
		maps[language] = append(maps[language], star)
	}

	slices.Sort(langs)

	content += "\n"

	for _, language := range langs {
		content += fmt.Sprintf("- [%s](#%s)\n", language, strings.ReplaceAll(strings.ToLower(language), " ", "-"))
	}

	for _, lang := range langs {
		stars, _ := maps[lang]

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
