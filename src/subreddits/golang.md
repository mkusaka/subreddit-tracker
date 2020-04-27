# golang
## [1][Telebot V2.1 released: Telegram bot framework in Go](https://www.reddit.com/r/golang/comments/g8xw2u/telebot_v21_released_telegram_bot_framework_in_go/)
- url: https://github.com/tucnak/telebot/releases/tag/v2.1
---

## [2][grobotstxt: A native Go port of Google's robots.txt parser and matcher C++ library.](https://www.reddit.com/r/golang/comments/g8v0to/grobotstxt_a_native_go_port_of_googles_robotstxt/)
- url: https://github.com/jimsmart/grobotstxt
---

## [3][What's coming in Go 1.15 — Slides By Daniel Martí (mvdan)](https://www.reddit.com/r/golang/comments/g8d8jk/whats_coming_in_go_115_slides_by_daniel_martí/)
- url: https://docs.google.com/presentation/d/1veyF0y6Ynr6AFzd9gXi4foaURlgbMxM-tmB4StDrdAM/edit#slide=id.g550f852d27_228_0
---

## [4][Whats the correct way to refer to an instantiated struct in documentation?](https://www.reddit.com/r/golang/comments/g8scow/whats_the_correct_way_to_refer_to_an_instantiated/)
- url: https://www.reddit.com/r/golang/comments/g8scow/whats_the_correct_way_to_refer_to_an_instantiated/
---
For example, "This function takes a pointer to an instantiated x struct and does y".

Is this the right term? would it be wrong to refer to it as an object? or an instance?
## [5][packagemain #18: Writing REST API Client in Go](https://www.reddit.com/r/golang/comments/g8nkhz/packagemain_18_writing_rest_api_client_in_go/)
- url: https://youtu.be/evorkFq3Y5k
---

## [6][Golang CMS and blog engines](https://www.reddit.com/r/golang/comments/g8unwf/golang_cms_and_blog_engines/)
- url: https://www.reddit.com/r/golang/comments/g8unwf/golang_cms_and_blog_engines/
---
Are there any good resources for golang available? I was looking for CMSes and blog engines written in go. Found ponzu and kabukky, but both seem a little dated. Any *go*to resources that you guys can recommend?
## [7][Code Review - Learning microservices](https://www.reddit.com/r/golang/comments/g8m9u7/code_review_learning_microservices/)
- url: https://www.reddit.com/r/golang/comments/g8m9u7/code_review_learning_microservices/
---
Hello,

I am a 19-year-old comp sci student from Montreal with a strong passion for the Go language.

I have recently been interested in learning more about microservices and have come across the go-kit library. I created a small test microservice that lists products from a database. 

I want to make sure I'm following best practices and that my code stays idiomatic. I am also unsure what the general opinion on go-kit is as I have seen mixed responses online.

I also would like maybe to make some friends here as I don't have people around me that share my interests. :)

Here is my source: https://github.com/sergiosegrera/go-kit-product

Thank you guys.
## [8][Chat with Go, React and k8s](https://www.reddit.com/r/golang/comments/g84yg2/chat_with_go_react_and_k8s/)
- url: https://www.reddit.com/r/golang/comments/g84yg2/chat_with_go_react_and_k8s/
---
Hey everybody! 
I built this simple chat in Golang and React and hosted it in a Kubernetes cluster in AWS. 
Though it might be useful to share a project with the full CI/CD pipeline.

[github repo](https://github.com/leartgjoni/go-chat-api)
## [9][[discussion] Why I'm hesitant / afraid of adopting go as main language](https://www.reddit.com/r/golang/comments/g8aujj/discussion_why_im_hesitant_afraid_of_adopting_go/)
- url: https://www.reddit.com/r/golang/comments/g8aujj/discussion_why_im_hesitant_afraid_of_adopting_go/
---
Is this normal behavior in Golang community? [https://i.imgur.com/wu9NAOb.png](https://i.imgur.com/wu9NAOb.png)

What's wrong with that framework? I'd love to use only stdlib but it feel too low level. Are we really expected for each project to manually parse the URL paths and map requests to their handlers? Or do we end up anyway making our own "framework" that we just keep private due to this kind of remarks?

What's wrong with having nice APIs like:

    app.Get("/", func (c *fiber.Ctx) {
        c.Send("Hello, World!")
    })

vs. needing to manually parse URL paths and their params?

Why don't we have a better router/mux in the stdlib if people hate so much on these frameworks?
## [10][Crud API Generator tool (gen)](https://www.reddit.com/r/golang/comments/g8oul9/crud_api_generator_tool_gen/)
- url: https://www.reddit.com/r/golang/comments/g8oul9/crud_api_generator_tool_gen/
---
I am looking for suggestions on a tool that I am contributing to on github. https://github.com/smallnest/gen 

The gen tool produces a CRUD (Create, read, update and delete) REST api project template from a given database. The gen tool will connect to the db connection string analyze the database and generate the code based on the flags provided. 

The code will be annotated with Swagger tags, for a Gin server.

The generated project will contain the following code under the `./example` directory.


  - Makefile
    - useful Makefile for installing tools building project etc. Issue `make` to display help output.
  - .gitignore
    - git ignore for go project
  - go.mod
    - go module setup, pass `--module` flag for setting the project module default `example.com/example` 
  - README.md
    - Project readme
  - app/server/main.go
    - Sample Gin Server, with swagger init and comments
  - api/*.go
    - REST crud controllers
  - dao/*.go
    - DAO functions providing CRUD access to database
  - model/*.go
    - Structs representing a row for each database table


I am looking for suggestions for improvement for the generated code as well as the project. Here are some simple steps to get started.


    ## install gen tool (should be installed to ~/go/bin, make sure ~/go/bin is in your path.
    $ go get -u github.com/smallnest/gen

    ## download sample sqlite database
    $ wget https://github.com/smallnest/gen/raw/master/example/sample.db

    ## generate code based on the sqlite database (project will be contained within the ./example dir)
    $ gen --sqltype=sqlite3 \
        --connstr "./sample.db" \
        --database main  \
        --json \
        --gorm \
        --guregu \
        --rest \
        --out ./example \
        --module example.com/rest/example \
        --mod \
        --server \
        --makefile \
        --json-fmt=snake \
        --overwrite

    ## build example code (build process will install packr2 if not installed)
    $ cd ./example
    $ make example

    ## binary will be located at ./bin/example
    ## when launching make sure that the sqlite file sample.db is located in the same dir as the binary 
    $ cp ../../sample.db  .
    $ ./example 

    ## Open a browser to http://localhost:8080/swagger/index.html

    ## Use wget/curl/httpie to fetch via command line
    http http://localhost:8080/albums
    curl http://localhost:8080/artists
