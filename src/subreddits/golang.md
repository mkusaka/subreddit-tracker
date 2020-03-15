# golang
## [1][My wife made a cake for my birthday](https://www.reddit.com/r/golang/comments/fioou4/my_wife_made_a_cake_for_my_birthday/)
- url: https://i.redd.it/uq50cj544pm41.jpg
---

## [2][Re-learning slice](https://www.reddit.com/r/golang/comments/fizlot/relearning_slice/)
- url: https://kilabit.info/journal/2020/re-learning_slice/
---

## [3][If Go channels are FIFO, why this prints the last one first always?](https://www.reddit.com/r/golang/comments/fix1k1/if_go_channels_are_fifo_why_this_prints_the_last/)
- url: https://www.reddit.com/r/golang/comments/fix1k1/if_go_channels_are_fifo_why_this_prints_the_last/
---
```go
package main

import "fmt"

func main() {
	ch := make(chan string)

	go write(ch, "1")
	go write(ch, "2")
	go write(ch, "3")
	go write(ch, "last")

	for i := 0; i &lt; 4; i++ {
		fmt.Println(&lt;-ch)
	}
}

func write(ch chan string, str string) {
	ch &lt;- str
}
```

Output
```
last
1
2
3
```
https://play.golang.org/p/RGsEcdjzb-e
## [4][hostctl: manage your /etc/hosts like a pro](https://www.reddit.com/r/golang/comments/fiyp8q/hostctl_manage_your_etchosts_like_a_pro/)
- url: https://github.com/guumaster/hostctl
---

## [5][[net/url] invalid url escape %2f error after creating query using url](https://www.reddit.com/r/golang/comments/fj00h7/neturl_invalid_url_escape_2f_error_after_creating/)
- url: https://www.reddit.com/r/golang/comments/fj00h7/neturl_invalid_url_escape_2f_error_after_creating/
---
https://pastebin.com/UFBFaP9N - code

Basically the error i get is   

    "https://api.sample.org%2F1%2Fvalidate-token?token=0c433896-5351-4e6": invalid URL escape "%2F" 

what i expect is 

    https://api.sample.org/1/validate-token?token=0c433896-5351-4"

unable to wrap my head around what excatly is going on because even on stackoverflow this seems to be optimal solution to create a proper query url
## [6][Dynamic service discovery in Golang micro-service.](https://www.reddit.com/r/golang/comments/fizssa/dynamic_service_discovery_in_golang_microservice/)
- url: https://www.reddit.com/r/golang/comments/fizssa/dynamic_service_discovery_in_golang_microservice/
---
I develop a web application using golang. In my app, I have hard-coded all the other services (pretty ugly).

`_, err := http.PostForm("http://user:7071/checkemail", url.Values{"email": {email}, "uid": {validatemailID.String()}})`


Is there a way to do this dynamically using golang?

Without using other services such as Zookeeper?
## [7][GolangCI.com is closing](https://www.reddit.com/r/golang/comments/fiif1m/golangcicom_is_closing/)
- url: https://medium.com/golangci/golangci-com-is-closing-d1fc1bd30e0e
---

## [8][Understand unsafe in GoLang](https://www.reddit.com/r/golang/comments/fiy0gf/understand_unsafe_in_golang/)
- url: https://www.pixelstech.net/article/1584241521-Understand-unsafe-in-GoLang
---

## [9][Creating Database Entities &amp; SQL Migrations - Learn Web Development in Go](https://www.reddit.com/r/golang/comments/fiq8ly/creating_database_entities_sql_migrations_learn/)
- url: https://www.youtube.com/watch?v=wZ46iWeozM8
---

## [10][Using modules locally without publishing to VCS](https://www.reddit.com/r/golang/comments/fix89j/using_modules_locally_without_publishing_to_vcs/)
- url: https://www.reddit.com/r/golang/comments/fix89j/using_modules_locally_without_publishing_to_vcs/
---
Hi I am new to golang and trying to understand modules. I have a module  that I implemented locally and want to use it from a different module  locally. I am not planning to push to github or any other server. Does  anyone know the best way to do this ?
