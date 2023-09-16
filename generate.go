package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
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

func main() {
	subReddit := "PokemonGoFriends"
	redditThread := "13ph1zs"

	page := getHttp("https://reddit.com/r/" + subReddit + "/comments/" + redditThread + ".json")

	regex := regexp.MustCompile(`([0-9]{4}\s[0-9]{4}\s[0-9]{4})`)
	friendCodes := regex.FindAllString(page, -1)
	fmt.Println("Found", len(friendCodes), "friendcodes")

	jsConst := "const friendCodes = [\"" + strings.Join(friendCodes, "\", \"") + "\"]"

	writeToFile(jsConst, "friendcodes.js")
}