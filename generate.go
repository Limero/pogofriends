package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func getHttp(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	panicOnError(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:81.0) Gecko/20100101 Firefox/81.0")

	resp, err := client.Do(req)
	panicOnError(err)
	defer resp.Body.Close()

	var body strings.Builder
	_, err = io.Copy(&body, resp.Body)
	panicOnError(err)

	return body.String()
}

func formatFriendCodes(friendCodes []string) []string {
	formattedFriendCodes := make([]string, len(friendCodes))
	for i, c := range friendCodes {
		if len(c) == 12 {
			formattedFriendCodes[i] = c[:4] + " " + c[4:8] + " " + c[8:12]
			continue
		}
		formattedFriendCodes[i] = c
	}
	return formattedFriendCodes
}

func removeDuplicates(originals []string) []string {
	var uniques []string

	for _, s := range originals {
		isUnique := true

		for _, unique := range uniques {
			if s == unique {
				isUnique = false
				break
			}
		}

		if isUnique {
			uniques = append(uniques, s)
		}
	}

	return uniques
}

func writeToFile(content, path string) {
	textfile, err := os.Create(path)
	panicOnError(err)
	defer textfile.Close()

	_, err = textfile.WriteString(content)
	panicOnError(err)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

var regex = regexp.MustCompile(`([0-9]{4}\s[0-9]{4}\s[0-9]{4}|[0-9]{12})`)

func findFriendCodes(content string) []string {
	return regex.FindAllString(content, -1)
}

func main() {
	subReddit := "PokemonGoFriends"
	redditThread := "19f9hne"

	page := getHttp("https://reddit.com/r/" + subReddit + "/comments/" + redditThread + ".json")

	friendCodes := findFriendCodes(page)
	friendCodes = formatFriendCodes(friendCodes)
	friendCodes = removeDuplicates(friendCodes)
	fmt.Println("Found", len(friendCodes), "friendcodes")

	lastUpdatedConst := fmt.Sprintf("const lastUpdated = %q", time.Now().UTC().String())
	jsConst := "const friendCodes = [\"" + strings.Join(friendCodes, "\", \"") + "\"]"

	writeToFile(lastUpdatedConst+"\n"+jsConst, "friendcodes.js")
}
