# golang
## [1][The Evolution of a Go Programmer](https://www.reddit.com/r/golang/comments/f1hy9a/the_evolution_of_a_go_programmer/)
- url: https://github.com/SuperPaintman/the-evolution-of-a-go-programmer
---

## [2][Excelize 2.1.0 is Released – Go lib for reading and writing Excel (XLSX) files](https://www.reddit.com/r/golang/comments/f1azwz/excelize_210_is_released_go_lib_for_reading_and/)
- url: https://github.com/360EntSecGroup-Skylar/excelize/releases/tag/v2.1.0
---

## [3][Looking for Practical Goroutine / Channels Exercises](https://www.reddit.com/r/golang/comments/f1bi58/looking_for_practical_goroutine_channels_exercises/)
- url: https://www.reddit.com/r/golang/comments/f1bi58/looking_for_practical_goroutine_channels_exercises/
---
Looking for practical exercises related to goroutines and channels. Most of the things I found online are tutorials and would love to get some more exercises. Then these exercises would cover different sub-topics (eg. waiting). It would be nice if these exercises have solution sets too.
## [4][[Rob Pike][2012] Less is exponentially more](https://www.reddit.com/r/golang/comments/f1nj16/rob_pike2012_less_is_exponentially_more/)
- url: https://commandcenter.blogspot.com/2012/06/less-is-exponentially-more.html
---

## [5][8chipgo - Chip 8 implemented in Go with pixel library](https://www.reddit.com/r/golang/comments/f1dhgv/8chipgo_chip_8_implemented_in_go_with_pixel/)
- url: https://i.imgur.com/6C46vtD.gif
---

## [6][Bosun - Kibana Automatic Index Pattern Discovery and Other Curating Tasks with Go](https://www.reddit.com/r/golang/comments/f1p7h7/bosun_kibana_automatic_index_pattern_discovery/)
- url: https://github.com/sherifabdlnaby/bosun
---

## [7][Go: Inlining Strategy &amp; Limitation](https://www.reddit.com/r/golang/comments/f18oi1/go_inlining_strategy_limitation/)
- url: https://medium.com/a-journey-with-go/go-inlining-strategy-limitation-6b6d7fc3b1be
---

## [8][Channels, structs, etc](https://www.reddit.com/r/golang/comments/f1kdem/channels_structs_etc/)
- url: https://www.reddit.com/r/golang/comments/f1kdem/channels_structs_etc/
---
Hey all,

Little frustrated as I can't quite put it together in my head on how to handle a scenario I have recently found myself in...

Basically I am sending an API request (in go) to a server.. in a go func, the response of which could be a large number of records. Now, in the go func, I am passed a channel which expects the record type I get back in the API response. What I am not sure of is if the channel should be designed to take an array of struct.. and pass it all back in one call.... or if per the nature of channels.. a stream.. should I do a loop through the records I get back, sending each one through the channel and requiring the receiver to then put them together as an array before doing something with them.

In this particular situation, the receiver end is itself part of an API req/resp.. so it would essentially be waiting for ALL the individual channel structs.. building up an array of the struct, to then return it as a single JSON response.

I think the right way to do this is the channel is set to stream one record at a time, and the receiver builds up the array of that struct, then marshalls the whole shabang in one JSON response.
## [9][Paste - a web app to keep encrypted text snippets, easily deployable to Heroku](https://www.reddit.com/r/golang/comments/f1ooo8/paste_a_web_app_to_keep_encrypted_text_snippets/)
- url: https://gitlab.com/opennota/paste
---

## [10][New JSON parser/interpreter](https://www.reddit.com/r/golang/comments/f1k1mt/new_json_parserinterpreter/)
- url: https://www.reddit.com/r/golang/comments/f1k1mt/new_json_parserinterpreter/
---
Hi gophers and other fellow life forms.

About 6 months ago we have started a new project that font-end side with `JS`/`Node.js`/`React` and back-end side  with `Golang`. I was working for the back-end side everything was O.K. But the project grows the `'encoding/json'` package get inefficient. For example we are not able to change a 'key' or change any value with something else. Thats because we switch to some custom packages for manipulating JSON. 

At this point. I figure out 'How about I wrote a packege for JSON manipulating'

And than I start working on it. But dealing with such a project was not easy.

First I had to find a way to validate my result. Manually writing unit tests was impossible for large data. Thats because I started to integrate Node.js to this project for test-case creation and validation. And it worked perfectly.

I use a continius integration platform for test automation. And wrote a detailed documentation on __[GoDoc](https://godoc.org/github.com/ecoshub/jin)__.

And It has a cool gopher logo you should see :)

Well today its finally ready for public release. 

Name of this package is __[JIN](https://github.com/ecoshub/jin)__

It is super easy to use and much more important its super fast.
I just want to share with this to you guys.

Let's look at with an example.

```go
	// this is the JSON we will work on
	data := []byte(`{"repo":{"name":"ecoshub/jin"},"others":["jin","dev.to"]}`)
```
we are going to try to access the 'dev.to' string in 'others' array.

```go
	// In order to Unmarshal a JSON.
	// first we have to define those structs properly.
	type Repository struct {
		Name string 		`json:"name"`
	}

	type Data struct {
		Repo   Repository 	`json:"repo"`
		Others []string 	`json:"others"`
	}
	// an empty data struct
	var newData Data

	// finally Unmarshaling.
	err := json.Unmarshal(data, &amp;newData)

	// standard error implementation.
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Prinln(newData.Others[1])

```
Or you can use this. with no struct defination.

```go
	// just one line of code.
	value, err := jin.Get(data, "others", "1")

	// standard error implementation.
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(value))
```

Do not miss-understand me I am not the first person who figure out this simple and elegant definition. I am just trying to expand and improve JSON manipulation.

This package has over __90__ functions/methods for ease JSON manipulation, build and formating needs. Some useful functions that i like to mention.
 
-`Add()`, `AddKeyValue()`, `Set()`, `SetKey()` `Delete()`, `Insert()`, `IterateArray()`, `IterateKeyValue()` `Tree()`. 

There is a very detailed explanations and lots of examples in __[GoDoc](https://godoc.org/github.com/ecoshub/jin)__ .

Also I think you have to  check out benchmark results :)

This is the link of repository:
https://github.com/ecoshub/jin

Please do not hesitate to fork/clone pull-request.

Thank you so much for your time.

Have a good day.
