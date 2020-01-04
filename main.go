package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	client *http.Client

	Logger *log.Logger
}

func NewClient(timeoutSecond time.Duration) *Client {
	client := &http.Client{
		Timeout: timeoutSecond,
	}
	return &Client{
		client: client,
		Logger: log.New(ioutil.Discard, "go-client: ", log.LstdFlags),
	}
}

var (
	Subreddits = []string{
		"androiddev",
		"Kotlin",
		"reduxjs",
		"reactjs",
		"typescript",
		"javascript",
		"cpp",
		"golang",
		"ruby",
		"rust",
		"aws",
		"rails",
	}
)

const baseURL = "https://www.reddit.com"
const baseSubredditURL = baseURL + "/r/"

// @param subreddit ex) Kotlin
// @return url ex) https://www.reddit.com/r/Kotlin.json
func URLGenerator(subreddit string) string {
	limit := 10 // default 25. its too long for me.
	return baseSubredditURL + subreddit + ".json" + "?" + "limit=" + strconv.Itoa(limit)
}

func JSONFilenameGnerator(subreddit string) string {
	return "src/raw/" + subreddit + ".json"
}

func MarkDownfilenameGenerator(subreddit string) string {
	return "src/subreddits/" + subreddit + ".md"
}

func FetchAllJSON() {
	c := NewClient(10 * time.Second)
	for _, subreddit := range Subreddits {
		url := URLGenerator(subreddit)
		log.Println(url + "\n")
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		// golang client has strict limitation from reddit. so we use custom header.
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
		resp, err := c.client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.Body != nil {
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			jsonfile, err := os.Create(JSONFilenameGnerator(subreddit))
			if err != nil {
				log.Fatal(err)
			}
			defer jsonfile.Close()

			jsonfile.WriteString(string(body))
		}
		time.Sleep(1 * time.Second)
	}
}

type PostData struct {
	name      string
	Title     string
	Subreddit string
	Selftext  string
	Permalink string
	Url       string
	Created   int
}

type Post struct {
	Kind string
	Data PostData
}

type Data struct {
	Modhash  string
	Dist     int
	Children []Post
}

type SubReddit struct {
	Kind string
	Data Data
}

func ParseSubreddit(rawSubReddits string) []Post {
	var subreddit SubReddit
	json.Unmarshal([]byte(rawSubReddits), &subreddit)
	return subreddit.Data.Children
}

func generateMarkDown(posts []Post, subreddit string, depth int, withBody bool) string {
	md := strings.Repeat("#", depth) + " " + subreddit + "\n"
	for idx, post := range posts {
		data := post.Data
		parmalink := baseURL + data.Permalink
		// TODO: show created time

		md += strings.Repeat("#", depth+1) + " [" + strconv.Itoa(idx+1) + "][" + data.Title + "](" + parmalink + ")\n" + "- url: " + data.Url + "\n"
		if withBody {
			md += "---\n" + data.Selftext + "\n"
		}
	}
	return md
}

func StoreEachSubredditMarkdownFromJSON() {
	for _, subreddit := range Subreddits {
		jsonByteString, err := ioutil.ReadFile(JSONFilenameGnerator(subreddit))
		if err != nil {
			log.Fatal(err)
		}
		parsed := ParseSubreddit(string(jsonByteString))
		md := generateMarkDown(parsed, subreddit, 1, true)
		f, err := os.Create(MarkDownfilenameGenerator(subreddit))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		f.WriteString(md)
	}
}

func StoreAllMarkdownFromJSON() {
	md := "# all\n"
	f, err := os.Create(MarkDownfilenameGenerator("all"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, subreddit := range Subreddits {
		jsonByteString, err := ioutil.ReadFile(JSONFilenameGnerator(subreddit))
		if err != nil {
			log.Fatal(err)
		}
		parsed := ParseSubreddit(string(jsonByteString))
		md += generateMarkDown(parsed, subreddit, 2, false)
	}
	f.WriteString(md)
}

func main() {
	// FetchAllJSON()
	StoreEachSubredditMarkdownFromJSON()
	StoreAllMarkdownFromJSON()
}
