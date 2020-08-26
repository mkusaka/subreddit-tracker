# golang
## [1][Just released go-reddit v1.0.0, a Go library for accessing the Reddit API](https://www.reddit.com/r/golang/comments/igwj4l/just_released_goreddit_v100_a_go_library_for/)
- url: https://github.com/vartanbeno/go-reddit
---

## [2][The Ultimate Go Book is here](https://www.reddit.com/r/golang/comments/igf4o7/the_ultimate_go_book_is_here/)
- url: https://www.reddit.com/r/golang/comments/igf4o7/the_ultimate_go_book_is_here/
---
Hi friends,

Almost a year ago, in August 2019, I shared one of my projects called Ultimate Go on GitHub, [https://github.com/hoanhan101/ultimate-go](https://github.com/hoanhan101/ultimate-go), and surprisingly, it got a lot of attention from the community. Fast forward to August 2020, it now has over 11K stars, 900 folks with the help of more than 20 contributors.

The project is a collection of my notes while learning Go programming language. Different people have different learning styles. For me, I learn best by doing and walking through examples. That said, I take notes carefully, comment directly on the source code to make sure that I understand every single line of code as I am reading and also be mindful of the theories behind the scene.

As Ultimate Go keeps growing, there’s one issue that keeps coming up. That’s about the format of the project. Specifically, many people have requested an ebook version where the content is more streamlined and they can read it at their convenience.

That’s when The Ultimate Go Book is born. For the last 3 months or so, I have spent most of my free time putting together everything from Ultimate Go into a 200-page book. Other than all the good stuff from Ultimate Go, two new and better things in this version are:

* Follow-along code input and output.
* Diagrams.

Hope it makes your journey of learning Go a bit easier. And again, thank you all for your support. I really appreciate it.

**Here's the link for the book:** [**https://gum.co/bpUYF**](https://gum.co/bpUYF)**.**

Note that I've made it free, though I would really appreciate your support. Regardless, you will receive the book in 2 different formats: PDF and ePub. 

Happy reading!
## [3][Golang on Blue Pill Microcontroller](https://www.reddit.com/r/golang/comments/igx7k5/golang_on_blue_pill_microcontroller/)
- url: https://i.redd.it/n7v2hg4z1cj51.jpg
---

## [4][Go Modules and Project Structuring](https://www.reddit.com/r/golang/comments/igyry9/go_modules_and_project_structuring/)
- url: https://www.reddit.com/r/golang/comments/igyry9/go_modules_and_project_structuring/
---
Hi everyone, I've been working with golang for the past couple of months. I've been really happy with the language and its ecosystem, but there's still something that bothers me and that's why I came here to get different opinions and possibly past experience on dealing with this subject. My big problem is project structuring and the use of go modules.

\[CONTEXT\]

I have developed a couple of microservices with REST APIs. My view is that since each one of these exposes a generic API and is an application by itself, each microservice should be a module with a main package. Nevertheless, they depend on each other (for example if service A interacts with service B it has to import the module to know what's the port it should contact the service at, and what are the request bodies like). Since I work on all the modules at the same time I had difficulties  figuring out how I could reflect the changes made in one module in another (without versioning). That's when I came across the replace instruction in the go.mod file. This has worked pretty good so far (removing the need to version each change made), but I feel like IDEs such as the one I use (GoLand) are not ready to deal with multiple modules in one project, (since they adopt a module per project mentality). Am I missing something, doing something wrong, or is the language and the ecosystem not yet perfect handling multi-module projects?

I would like to thank you for your comments in advance :)
## [5][Just Created a Free Interview Prep Review](https://www.reddit.com/r/golang/comments/igyl10/just_created_a_free_interview_prep_review/)
- url: https://www.reddit.com/r/golang/comments/igyl10/just_created_a_free_interview_prep_review/
---
Howdy, all. The other week I was surfing around and looking for some Go interview preparation, or even just a good list of common interview questions. It seemed to me that all the top results on Google are pretty much garbage, low effort reposts of each other. Most of the top posts also were written in broken English :(

I just published a free-tier course on Qvault in hopes that it will help some people out there. I would love to iterate on this and have it grow to become a one-stop-shop for Go interview review. Anyhow, here is the link to the course page if you want to give it a shot and let me know what you think:

[https://qvault.io/2020/08/26/interview-prep-golang-course-released/](https://qvault.io/2020/08/26/interview-prep-golang-course-released/)
## [6][My first English article: In-memory data type security for Go (Secure Types)](https://www.reddit.com/r/golang/comments/igydpt/my_first_english_article_inmemory_data_type/)
- url: https://link.medium.com/CCzj0a5We9
---

## [7][Anyone interested in a book about graphics in Go?](https://www.reddit.com/r/golang/comments/ig7a0h/anyone_interested_in_a_book_about_graphics_in_go/)
- url: https://www.reddit.com/r/golang/comments/ig7a0h/anyone_interested_in_a_book_about_graphics_in_go/
---
This is more of a general inquiry at this point. I love both Go and graphics programming. At some point I thought, I might actually write a short book combining both topics. As we all know, Go is much more than just a language for writing Web services, yet most materials teach only this side of it. I thought that by using Go to teach a different context (e.g. computer graphics), I might give it a different spin and bring in some fresh minds into the community.

I will soon share a link to my first scratch. It is not really a full chapter at the moment. At this point, I am combining a wiki of things, from basic GP concepts to more advanced, like game programming. 

What do you think? Will this be something useful for the general Go community? Feel free to share feedback as well as resources (e.g. libraries, tutorials, etc) which you want to see featured there.
## [8][0/0 is -9223372036854775808 !!!](https://www.reddit.com/r/golang/comments/igwn2n/00_is_9223372036854775808/)
- url: https://www.reddit.com/r/golang/comments/igwn2n/00_is_9223372036854775808/
---
I found bug in my software (divide zero) but instead of panic I get a *nice* number...

[https://play.golang.org/p/YTaDA7cUXSB](https://play.golang.org/p/YTaDA7cUXSB)
## [9][Fuzzing in Go](https://www.reddit.com/r/golang/comments/igjte6/fuzzing_in_go/)
- url: https://lwn.net/SubscriberLink/829242/64cbde0531aaf166/
---

## [10][Experimenting with QUIC and WebTransport in Go](https://www.reddit.com/r/golang/comments/igg1os/experimenting_with_quic_and_webtransport_in_go/)
- url: https://centrifugal.github.io/centrifugo/blog/quic_web_transport/
---

