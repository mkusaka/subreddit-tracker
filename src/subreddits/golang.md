# golang
## [1][A Virtual Stock Market Exchange game in Golang](https://www.reddit.com/r/golang/comments/fc7olw/a_virtual_stock_market_exchange_game_in_golang/)
- url: https://www.reddit.com/r/golang/comments/fc7olw/a_virtual_stock_market_exchange_game_in_golang/
---
Dalal Street is a virtual stock market exchange game developed by a bunch of undergrads. It was built using golang.

We used

* Goroutines to handle streams / transactions / stocks matching simultaneously
* Gorm as ORM (mysql database)
* GRPC - to handle communication between frontend and backend
* Server side streams (using GRPC) to send updates at frontend
* Migrate CLI for handling migrations
* Mockgen for generating mocks for testing
* grpc-middleware for authentication
* caching for getting sessions and some other info
* Mutexes to handle race conditions and inconsistency

Code : [https://github.com/delta/dalal-street-server](https://github.com/delta/dalal-street-server)

&amp;#x200B;

In this you will be given a fixed amount of money and you aim is to maximise your wealth.

Game is live right now !! From 1st March - 7th March, 2:30PM - 5:30 PM UTC.

Web : [https://dalal.pragyan.org](https://dalal.pragyan.org/trade)

App :  [http://prgy.in/dalalapp](http://prgy.in/dalalapp)

&amp;#x200B;

Please do check it out !!
## [2][Library for hashing an interface{}](https://www.reddit.com/r/golang/comments/fcb9ul/library_for_hashing_an_interface/)
- url: https://www.reddit.com/r/golang/comments/fcb9ul/library_for_hashing_an_interface/
---
Hello reddit gophers,

I'm a software developer from Athens, Greece and a professional Golang developer for the last six months.

During work, I fell onto a case where I had to compare two huge structs and see if they are equal and also ignore some struct fields during comparison.

The first "naive" way to do so, is to write a huge function where this comparison takes place. I thought about avoiding this in the first place.

After that, I thought doing JSON Marshal to both structs that I had to compare, then use some regex and string modification to remove the fields, and then JSON unmashal again. It worked, but it seemed a bit "expensive".

So I did a small research and I decided to exploit the capabilities of Golang `reflect` package. I finally wrote this library below, which iterates recursively in all fields of the given struct (ignoring the ones that have the tag `"hash:"ignore"`) and generates a hash string.

[https://github.com/panospet/recursive-deep-hash](https://github.com/panospet/recursive-deep-hash)

I'm pretty sure it's not perfect yet, so I need your advice / help.Feel free to give feedback, comments, improvements suggestions, possible bugs, etc.
## [3][Which framework should i use for web development in go](https://www.reddit.com/r/golang/comments/fc9slb/which_framework_should_i_use_for_web_development/)
- url: https://www.reddit.com/r/golang/comments/fc9slb/which_framework_should_i_use_for_web_development/
---
Hello guys,

I would really appreciate your help if you could recommend me a good framework in go for web development. I write go in my free time but my main languages are php and java. My experience in go comes from:

1. Small projects i wrote on my own.
2. By reading [https://www.openmymind.net/The-Little-Go-Book/](https://www.openmymind.net/The-Little-Go-Book/)
3. By watching [https://www.udemy.com/course/learn-how-to-code/](https://www.udemy.com/course/learn-how-to-code/)
4. Attended a golang conference in London 2 years ago.

I know i can use net/http to create a web server in go but what i want to find out is if golang has something like  spring boot (java) and symfony (php). These are the 2 frameworks i am using.

&amp;#x200B;

When i attended the conference 2 years ago they were really trying to promote go bufallo but a lot of people are saying that it is mainly a code generator. I can see that gin [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) is really popular but i don't know if it is what i am looking for.

I am looking forward to see your recommendations.
## [4][Lessons Learned after 1 year of programming in Go as a C# developer | Robert Pajak](https://www.reddit.com/r/golang/comments/fc9els/lessons_learned_after_1_year_of_programming_in_go/)
- url: https://pajak.home.blog/2020/02/28/lessons-learned-after-1-year-of-programming-in-go-as-a-c-developer/
---

## [5][How to create a CLI in golang with cobra](https://www.reddit.com/r/golang/comments/fcc2e5/how_to_create_a_cli_in_golang_with_cobra/)
- url: https://schadokar.dev/posts/how-to-create-a-cli-in-golang-with-cobra/
---

## [6][Seeking testers for gotop](https://www.reddit.com/r/golang/comments/fbtyl0/seeking_testers_for_gotop/)
- url: https://www.reddit.com/r/golang/comments/fbtyl0/seeking_testers_for_gotop/
---
I'm looking for some volunteers to help testing the next release of gotop.

The original author of gotop decided to work on another project and flagged gotop as "unmaintained."  I've been working with him to assume maintainership, and in the past month have released a number of versions that merge long-outstanding pull requests and add new features and bug fixes.  I have a new major version ready to go out, but as I'm starting to get distribution packagers on board with the switch, testing becomes more of a concern for me.  I'm appealing to the community for some light assistance in smoke testing releases for other architectures and OSes that I'm unable to do myself, and for assistance with distribution packages.

- OS/Arch testing. I need testers for FreeBSD, Darwin, and Windows, and for architectures other than AMD64.
- Homebrew testing.  I'd appreciate help testing the Homebrew tap, and the instructions that are in the README.
- Device testing.  One of the features being added is instrumentation support for GPUs.  NVIDIA support is pending, and AMD and Intel are in the queue.
- Distribution packages.  The AUR packages have already been transferred, but I can only test that the packages work for my CPU architecture. I'm both unfamiliar and unsure how to start with Debian and Redhat derivatives; do they have AUR analogues?  What about Slackware and Alpine?
- Performance testing. Mostly, I'd like feedback about behaviors and would especially like to have folks out there watching for changes that negatively impact.  I think I'm pretty good about this, but I can only really watch how it performs on my machine.

Also, if any tuning gurus are willing to help with profiling, I'd love to get some tickets in about areas to investigate.  I'm capable of doing this myself, but will probably not get to it for a while as I continue to ramp up on the maintainer aspects.

Finally, I'm always happy to look at pull requests, and will try to prioritize those.

I'm getting some help already from a few folks who have been active in helping me with the transition, but I still have gaps (no ARM testers, for instance, and FreeBSD is a big worry).

To smooth the transition, I've [kept the project in github](https://github.com/xxxserxxx/gotop).

In addition to bug fixes, the features released (and currently tested) are:

- Instrumentation data export via HTTP.
- A battery gauge
- Support for GPUs
- Customizable widget layout via config file
- Error log rotation
- Config file support
- Support for multiple network interfaces (incl. filtering)
- Searching in the proc widget

Once again, the project URL is https://github.com/xxxserxxx/gotop .  An call for testers is in the issue tracker; if you're willing to help, please post to that ticket.  I'm only seeking sanity checks for things I can't test myself, so if someone has already volunteered to, e.g., test FreeBSD, don't feel obliged to also volunteer.

Thanks for reading the long post.
## [7][Can you use Golang for low-level system and embedded programming?](https://www.reddit.com/r/golang/comments/fbzd53/can_you_use_golang_for_lowlevel_system_and/)
- url: https://www.reddit.com/r/golang/comments/fbzd53/can_you_use_golang_for_lowlevel_system_and/
---
Can you use Golang for low-level system and embedded programming?
## [8][Build a CRUD application in Golang with PostgreSQL - CodeSource.io](https://www.reddit.com/r/golang/comments/fc7gi0/build_a_crud_application_in_golang_with/)
- url: https://codesource.io/build-a-crud-application-in-golang-with-postgresql/
---

## [9][zserge/webview](https://www.reddit.com/r/golang/comments/fc6qwt/zsergewebview/)
- url: https://www.reddit.com/r/golang/comments/fc6qwt/zsergewebview/
---
Cool, tiny cross-platform [webview library](https://github.com/zserge/webview) for Go and C/C++ by u/zserge. I had experimented with using it for demos at AWS and liked it, but was wary of committing to the library knowing big changes were coming. Looks like they finally landed, so time to revisit!
## [10][Looking for the most user-friendly template engine for my tool](https://www.reddit.com/r/golang/comments/fc6ivy/looking_for_the_most_userfriendly_template_engine/)
- url: https://www.reddit.com/r/golang/comments/fc6ivy/looking_for_the_most_userfriendly_template_engine/
---
Briefly about the project: I am creating a utility to generate examples for HTTP APIs. You write a configuration file, specifying for which requests you want to make examples and utility performs these requests and outputs their response.

Github: [api-example-generator](https://github.com/flygrounder/api-example-generator)

Now to the question: there is an ability to use response of one request to generate another request. You can access this data using templates. At the moment I use go default templates, which works but they seem not user-friendly for people not familiar with go. Which template engine you would recommend to use here?
