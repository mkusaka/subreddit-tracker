# golang
## [1][gopls/v0.4.1](https://www.reddit.com/r/golang/comments/gjhsp3/goplsv041/)
- url: https://github.com/golang/go/milestone/146?closed=1
---

## [2][When to use pointers in Go?](https://www.reddit.com/r/golang/comments/gj20kj/when_to_use_pointers_in_go/)
- url: https://medium.com/@meeusdylan/when-to-use-pointers-in-go-44c15fe04eac
---

## [3][Handy Gitlab CI YAML job templates for Go projects](https://www.reddit.com/r/golang/comments/gjipr5/handy_gitlab_ci_yaml_job_templates_for_go_projects/)
- url: /r/gitlab/comments/gjimb1/handy_gitlab_ci_yaml_job_templates_for_go_projects/
---

## [4][Are there any Benchmark results available for various Golang Websocket?](https://www.reddit.com/r/golang/comments/gjjb9e/are_there_any_benchmark_results_available_for/)
- url: https://www.reddit.com/r/golang/comments/gjjb9e/are_there_any_benchmark_results_available_for/
---
I am confused about choosing which WebSocket is fast for "room based chat app" to use in production apps:

1. [Kataras neffos](https://github.com/kataras/neffos)
2. [Gorilla Websocket](https://github.com/gorilla/websocket/)
3. [Gobwas Websocket](https://github.com/gobwas/ws)
4. [go-socket-io](https://github.com/googollee/go-socket.io)
## [5][HTTP GET with json payload](https://www.reddit.com/r/golang/comments/gjl063/http_get_with_json_payload/)
- url: https://www.reddit.com/r/golang/comments/gjl063/http_get_with_json_payload/
---
Hi!

I'm currently trying to learn golang while using it for something useful. So i have a elasticsearch database where i want to create a query to the database and return the results but i'm having difficulty understand how to construct such a query with golang.

Here is the cURL command i used:

&gt;curl -XGET "[https://elasticsearch:9200/index-\*/\_search](https://caroli-es01.caroli.svc.clu01.ilog.local:9200/ishield-dns-*/_search)" -H 'Content-Type: application/json' -d'{  "size": 1,  "sort": {      "@timestamp": {        "order": "desc"      }    }}'

And here is the code i've been trying to puzzle together:

```
package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//"time"

const (
	ElasticURL    string = "https://elasticsearch:9200/index-*/_search"
	ElasticUser   string = "admin"
	ElasticPasswd string = "password"
)

func main() {
	// Map variable data to struct
	//var data string
	type Payload struct {
		Size int  `json:"size"`
		Sort Sort `json:"sort"`
	}
	type Timestamp struct {
		Order string `json:"order"`
	}
	type Sort struct {
		Timestamp Timestamp `json:"@timestamp"`
	}

	data := Payload{
		Size: 1,
		Sort: Sort{Timestamp: Timestamp{Order: "desc"}},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	PayloadBody := bytes.NewReader(payloadBytes)
	//Http client to make API request with basic auth
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &amp;tls.Config{InsecureSkipVerify: true}
	ElasticClient := &amp;http.Client{}
	ElasticReq, err := http.NewRequest("GET", ElasticUrl, nil)
	ElasticReq.SetBasicAuth(ElasticUser, ElasticPasswd)
	//ElasticReq.Header.Set("Content-Type", "application/json")
	resp, err := ElasticClient.Do(ElasticReq)
	if err != nil {
		log.Fatalln(err)
	}

	//Open the API response and unmarshal to struct
	body, err := ioutil.ReadAll(resp.Body)
	//json.Unmarshal(body, &amp;data)
	output := ioutil.NopCloser(bytes.NewBuffer(body))
	fmt.Println(output)

}
```

But im just getting an error where it says "Sort" is undeclared. Is there a problem with how i populate the Sort struct?
## [6][How to handle errors in Go? [5 rules]](https://www.reddit.com/r/golang/comments/gj35yc/how_to_handle_errors_in_go_5_rules/)
- url: https://web3.coach/golang-how-to-handle-errors-five-rules
---

## [7][I've been out of the Go loop since 2017, what did I miss?](https://www.reddit.com/r/golang/comments/gj0u58/ive_been_out_of_the_go_loop_since_2017_what_did_i/)
- url: https://www.reddit.com/r/golang/comments/gj0u58/ive_been_out_of_the_go_loop_since_2017_what_did_i/
---
Go was my daily driver from 2014 - 2017, and I'm gearing up to write some Go again. 

What has happened since 2017? Best practices, language features, design patterns, ecosystem tooling, etc. Thanks Reddit!
## [8][Too modern Go application? Building a serverless application with Google Cloud Run and Firebase](https://www.reddit.com/r/golang/comments/giwx9q/too_modern_go_application_building_a_serverless/)
- url: https://threedots.tech/post/serverless-cloud-run-firebase-modern-go-application/
---

## [9][A new way to create plots with Go with go-plotly](https://www.reddit.com/r/golang/comments/gj438z/a_new_way_to_create_plots_with_go_with_goplotly/)
- url: https://www.reddit.com/r/golang/comments/gj438z/a_new_way_to_create_plots_with_go_with_goplotly/
---
hi all! This is my first post in reddit. Let's see how this platform works.

I fell in love with Plotly when I learned some data science with python, pandas and jupyter notebooks and when I started my adventure learning go, I missed a lot the Plotly features. That's why, after a few year of experience, I bring to you the library go-plotly inspired by python plotly to help you bulding plots in your favourite language :)

[https://github.com/MetalBlueberry/go-plotly](https://github.com/MetalBlueberry/go-plotly)

Let me know if you like It and go to my blog if you want to know more about the story behind.

[https://metalblueberry.github.io/post/blog/2020-05-11\_go\_plotly/](https://metalblueberry.github.io/post/blog/2020-05-11_go_plotly/)
## [10][lsdy - a cmdline tool for querying DynamoDB tables. Heavily used both at home and work. Maybe useful to anybody out there as well.](https://www.reddit.com/r/golang/comments/gj2ywe/lsdy_a_cmdline_tool_for_querying_dynamodb_tables/)
- url: https://github.com/flowerinthenight/lsdy
---

