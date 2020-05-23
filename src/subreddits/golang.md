# golang
## [1][exhaustive: a tool to ensure exhaustiveness of enum switch statements](https://www.reddit.com/r/golang/comments/gp2u4z/exhaustive_a_tool_to_ensure_exhaustiveness_of/)
- url: https://github.com/nishanths/exhaustive
---

## [2][Why is Golang so popular for systems programming?](https://www.reddit.com/r/golang/comments/got0yp/why_is_golang_so_popular_for_systems_programming/)
- url: https://www.reddit.com/r/golang/comments/got0yp/why_is_golang_so_popular_for_systems_programming/
---
I see projects like Docker, Kubernetes, and a few other tools written in Go so I was just curious.

A project I'm working on requires using the Docker sdk extensively and I was considering if I should learn Go and use it instead of Python.
## [3][Getting Started With LDAP in Go](https://www.reddit.com/r/golang/comments/goleyg/getting_started_with_ldap_in_go/)
- url: https://cybernetist.com/2020/05/18/getting-started-with-go-ldap/
---

## [4][Horizontally scalable web scraper written in Go](https://www.reddit.com/r/golang/comments/goze9i/horizontally_scalable_web_scraper_written_in_go/)
- url: https://github.com/sfproductlabs/scrp
---

## [5][A Closer Look at Go (golang) Type System.](https://www.reddit.com/r/golang/comments/goy6v5/a_closer_look_at_go_golang_type_system/)
- url: https://medium.com/@ankur_anand/a-closer-look-at-go-golang-type-system-3058a51d1615
---

## [6][A simple terminal UI for Go linters](https://www.reddit.com/r/golang/comments/gp4sdx/a_simple_terminal_ui_for_go_linters/)
- url: https://www.reddit.com/r/golang/comments/gp4sdx/a_simple_terminal_ui_for_go_linters/
---
[https://github.com/nakabonne/golintui](https://github.com/nakabonne/golintui)

I made a TUI tool to operate golangci-lint intuitively.

A noteworthy feature is that you can open a file by specifying the issue line. I’d be happy if you could try it, and give me any feedback.

https://i.redd.it/mojzzjd4mi051.gif
## [7][I just put Clean Code's teachings to work!](https://www.reddit.com/r/golang/comments/gp1fjg/i_just_put_clean_codes_teachings_to_work/)
- url: https://www.reddit.com/r/golang/comments/gp1fjg/i_just_put_clean_codes_teachings_to_work/
---
Just bought [Clean Code](https://www.amazon.com/Clean-Code-Handbook-Software-Craftsmanship-ebook/dp/B001GSTOAM) a few days ago, only been through the first few chapters but figured I would try applying it.

The least I can say is, it feels pretty good to crunch that big function! I can still see possible improvements (like create a `isTr(n *html.Node) { return isElement(n, "tr")}`), but I'm [pretty proud of it](https://github.com/nicolascouvrat/javaimports/commit/ffffd8a8a64580b75fa3e75aadfa5dc7c9416bd2).

Do you spot anything else I should improve there?
## [8][Variable declarations: A question of style](https://www.reddit.com/r/golang/comments/gp1agj/variable_declarations_a_question_of_style/)
- url: https://www.reddit.com/r/golang/comments/gp1agj/variable_declarations_a_question_of_style/
---
Hi, I've really just started using go seriously and I'm interested in people's opinions on best practices/common conventions in variable declarations. 

I have the following function which does exactly the same thing as far as I'm aware, but has a different syntax. Is there a preferred style that is more accepted among the go community?

Method 1:

```
// GetDb connects to the postgres database and returns a pointer to a gorm database object
func GetDb() *gorm.DB {
	var db *gorm.DB
	var dbURL string
	var err error
	var parsedURL *dburl.URL

	if dbURL = os.Getenv("DATABASE_URL"); dbURL == "" {
		dbURL = "postgresql://localhost/mydb?sslmode=disable"
	}

	if parsedURL, err = dburl.Parse(dbURL); err != nil {
		log.Fatal(err)
	}

	if db, err = gorm.Open("postgres", parsedURL.DSN); err != nil {
		log.Fatal(err)
	}

	return db
}
```
Method 2:
```
// GetDb connects to the postgres database and returns a pointer to a gorm database object
func GetDb() *gorm.DB {

	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		dbURL = "postgresql://localhost/mydb?sslmode=disable"
	}

	parsedURL, err := dburl.Parse(dbURL)

	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("postgres", parsedURL.DSN)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
```

From what I've seen, method 2 seems to be the more popular in tutorials but I personally find method 1 to be clearer.
## [9][I made a basic Windows version of neofetch in Go after picking up the language last night! I'm in love!!](https://www.reddit.com/r/golang/comments/gonc7b/i_made_a_basic_windows_version_of_neofetch_in_go/)
- url: https://github.com/M4cs/winfetch
---

## [10][10 part series to build web app with Go](https://www.reddit.com/r/golang/comments/gobfe7/10_part_series_to_build_web_app_with_go/)
- url: https://www.gophersumit.com/series/web/
---

