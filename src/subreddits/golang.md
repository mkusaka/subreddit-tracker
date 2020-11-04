# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][Today, I presented Go to my team...](https://www.reddit.com/r/golang/comments/jnhzm5/today_i_presented_go_to_my_team/)
- url: https://www.reddit.com/r/golang/comments/jnhzm5/today_i_presented_go_to_my_team/
---
... and it went horribly wrong.

So to give you context, we’re a company having a SaaS written in PHP and are in the early stages of rewriting in from Yii to Laravel.

Today I had a presentation with a lot of the backend devs where I wanted to introduce Go to them and how it can help us writing better software.

I started by showing some code examples, how it fits very nicely into the web development world, the toolchain and a few other things like documentation, concurrency and speed. Then I continued by comparing it to PHP in terms of type safety and error handling, code refactoring and also in terms of simplicity. I moved on talking about how fast we are currently moving (in all company departments) and that Go would be a great fit for us. Finally I showed, how we could easily migrate the API with loadbalancing the endpoints.

Then started the Q&amp;A which was very very silent at first. One dev broke it by asking how it compares to node js. As I don’t have much experience in that, I could only tell what I‘ve heard of. But then we wanted to hear what our team lead had to say to all of this: 

He said that he was silent at first because he was googling. He followed up by thanking me for the presentation and immediately added, that my examples were heavily Go biased and not well worked out when comparing it to PHP. His google search resulted in articles he found about: 1. Go doesn’t work for complex projects and 2. that it’s very Memory hungry. I responded with 1. Docker/Kubernetes and 2. that it’s machine code with no Interpreter or VM, so what could be memory hungry then?

So the end result was, that nobody else said something to it and that I went out with a very bad feeling. I invested about a week into it, even creating a sample app, which we‘ve live deployed. Just to get very negative feedback on it. Feels like I could have taken the time creating something useful instead.

Did you have similar experience when trying to introduce Go to your team? 

Please share your experience as I am very very interested in that.

Thanks for making it to this point &lt;3

Edit: 
One other thing I found quite interesting: one dev asked about migrations and said that you would then probably need to be a SQL expert if you don’t have ORMs of a Library managing that for you. I replied that there was gorm (even though I’m personally not liking it.)

I didn’t want to say, that he 1. didn’t need to be an sql expert. 2. that with being a programmer in general you should know how to SQL anyway.
## [3][Etymology of Dial()](https://www.reddit.com/r/golang/comments/jnqs8e/etymology_of_dial/)
- url: https://www.reddit.com/r/golang/comments/jnqs8e/etymology_of_dial/
---
In many Go packages, `Dial()` is the common name for creating a connection to a server. But `Connect()` is a more common name for the same purpose in other languages. I wonder the reason why Go chose `Dial()` instead of `Connect()`.
## [4][How to Create a Go (Golang) API on Google App Engine](https://www.reddit.com/r/golang/comments/jnu3mm/how_to_create_a_go_golang_api_on_google_app_engine/)
- url: https://www.reddit.com/r/golang/comments/jnu3mm/how_to_create_a_go_golang_api_on_google_app_engine/
---
A [simple walkthrough](https://itnext.io/how-to-create-a-go-golang-api-on-google-app-engine-157e7cd33a93?source=friends_link&amp;sk=f6d5a07c82ed939c3432b789774015bf) on creating an API with a server, written in Go, to deploy to Google App Engine.
## [5][Clean Architecture by Robert C. Martin](https://www.reddit.com/r/golang/comments/jnwc28/clean_architecture_by_robert_c_martin/)
- url: https://www.reddit.com/r/golang/comments/jnwc28/clean_architecture_by_robert_c_martin/
---
Hi,

please read my article about Clean Architecture that I wrote to share some of my insights. I hope you find it useful, educational and interesting.

[https://adevelopersdiscourse.blogspot.com/2020/06/clean-architecture-demystified.html](https://adevelopersdiscourse.blogspot.com/2020/06/clean-architecture-demystified.html)

Clean Architecture was invented by Robert C. Martin that co-wrote the Agile Manifesto and coined the SOLID principles. It's well worth looking into.
## [6][Usage of pointers in golang in backend development](https://www.reddit.com/r/golang/comments/jnvdom/usage_of_pointers_in_golang_in_backend_development/)
- url: https://www.reddit.com/r/golang/comments/jnvdom/usage_of_pointers_in_golang_in_backend_development/
---
Hello everyone! I used nodejs/express in the past for developing personal projects for my github portfolio. I am really interested in Go programming but useage of \* and &amp; confuses me even though we are using C++ in university. I mean I know what are they and how they are used but never used them very effective in any projects.

I was following  Nic Jackson's microservices videos on youtube and he is using pointers A LOT! I feel like I am missing something. Where can I find a good resource that explains these concepts. Open for advices for learning backend development with Go.
## [7][C# to Go?](https://www.reddit.com/r/golang/comments/jnloxc/c_to_go/)
- url: https://www.reddit.com/r/golang/comments/jnloxc/c_to_go/
---
I have 9 years experience on C# and used it from mobile, web and back-end systems. Just currently, I got hired to a company that wants to transition from PHP to Go. I accepted the job since I'm interested in learning Go. Is this a good career move?
## [8][What is the best microservice framework in Go?](https://www.reddit.com/r/golang/comments/jnv4bd/what_is_the_best_microservice_framework_in_go/)
- url: https://www.reddit.com/r/golang/comments/jnv4bd/what_is_the_best_microservice_framework_in_go/
---
Go noob here so please don't yell at me for asking this question. I am interested to know which frameworks does a lot of code generation which alleviates the developer having to write boilerplate code. I have been a C# developer for more than 11 years now, but I'm finding Go extremely agile. To give a perspective Dotnet world has something like the Tye project. Was wondering if there are similar frameworks in Go.
## [9][Why doesn’t golang allow me to do this](https://www.reddit.com/r/golang/comments/jn6oo0/why_doesnt_golang_allow_me_to_do_this/)
- url: https://i.redd.it/z4jnotbkizw51.jpg
---

## [10][How to escape an url with zero allocations, I was quite surprised I couldnt find something like this out there already.](https://www.reddit.com/r/golang/comments/jncmsr/how_to_escape_an_url_with_zero_allocations_i_was/)
- url: https://github.com/Villenny/fastUrlEscape-go
---

## [11][IP camera feed](https://www.reddit.com/r/golang/comments/jnlwyf/ip_camera_feed/)
- url: https://www.reddit.com/r/golang/comments/jnlwyf/ip_camera_feed/
---
Recently bought an IP camera, and am wanting to write a server for storing/feeding live footage. After hours of research, I have come to 2 conclusions:

1. I want to avoid 3rd party libraries such as OpenCV
2. I have less networking knowledge than expected

I understand go/programming but very green with networking.

Where should I start? What protocols should I use?

Thank you.
