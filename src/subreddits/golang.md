# golang
## [1][Controversial opinion : Why do people keep sharing gopher plush toy pictures?](https://www.reddit.com/r/golang/comments/fg69gv/controversial_opinion_why_do_people_keep_sharing/)
- url: https://www.reddit.com/r/golang/comments/fg69gv/controversial_opinion_why_do_people_keep_sharing/
---
The choice of a language is completely independent of the plush toy representation of it. Golang has absolutely nothing to do with a googly eyed gopher. Are we a bunch of a passionate software engineers or a bunch of toddlers? /rant

Go ahead share em here. Make me cry.
## [2][go test -v streaming output](https://www.reddit.com/r/golang/comments/fg9xda/go_test_v_streaming_output/)
- url: https://dave.cheney.net/2020/03/10/go-test-v-streaming-output
---

## [3][Building Web Server with Go - Part 9](https://www.reddit.com/r/golang/comments/fg92lh/building_web_server_with_go_part_9/)
- url: https://www.reddit.com/r/golang/comments/fg92lh/building_web_server_with_go_part_9/
---
How to use MySQL with Go : [https://www.gophersumit.com/series/web-9/](https://www.gophersumit.com/series/web-9/)
## [4][New friend at work](https://www.reddit.com/r/golang/comments/ffukbe/new_friend_at_work/)
- url: https://i.redd.it/u4t3x90henl41.jpg
---

## [5][Vecosy - Centralized configuration service - first release](https://www.reddit.com/r/golang/comments/fg9cpg/vecosy_centralized_configuration_service_first/)
- url: https://www.reddit.com/r/golang/comments/fg9cpg/vecosy_centralized_configuration_service_first/
---
Hi Gophers!

I would like to share with you my project [vecosy](https://github.com/vecosy/vecosy) that should help with services configurations.

Main features:

* configuration on GIT
* three different merging strategies
* auto search the nearest configuration version
* watch changes detection (only GRPC)
* REST / GRPC
* [Spring-Cloud-Conf](https://cloud.spring.io/spring-cloud-config/reference/html/) compatible
* [viper](https://github.com/spf13/viper) compatible
* JWS
* TLS
* Helm chart (work in progress)

There is a demo in docker to try it and various examples.

Comments, suggestions, contributions (especially contributions!) are absolutely welcome.

Thank you for your time!
## [6][I am comfortable writing go but that "Intuition" isn't there yet. What are some easy-to-read projects on github I could skim that are great examples of using Go in the real world?](https://www.reddit.com/r/golang/comments/fgc4ki/i_am_comfortable_writing_go_but_that_intuition/)
- url: https://www.reddit.com/r/golang/comments/fgc4ki/i_am_comfortable_writing_go_but_that_intuition/
---
Forgive the vague ask - I want to read some good and proper Go code but a lot of the projects shared here are way over my head.

What are some good examples of projects that use Go cleanly, but won't intimidate a beginner such as myself?

Thanks!
## [7][[Newbie question] How to have an src directory when using go mod](https://www.reddit.com/r/golang/comments/fgdl7c/newbie_question_how_to_have_an_src_directory_when/)
- url: https://www.reddit.com/r/golang/comments/fgdl7c/newbie_question_how_to_have_an_src_directory_when/
---
Hi,

I'm a Node and PHP developer starting to build a GraphQL API using golang.

To use go-pg, I had to do a \`go mod init\` and rename all my packages with my repository name. 

From that point, I've been unable to keep my src directory as go didn't resolved my imports of "[github.com/myuser/myrepo/types](https://github.com/myuser/myrepo/types)" for example (it tried to download the lib and failed).

    # Before
    docker/
    ci/
    doc/
    src/
    +--types/
    +----user.go
    +--resolvers/
    +----user.go
    +--main.go
    
    # After
    docker/
    ci/
    doc/
    types/
    resolvers/
    main.go

To run my script, I used \`go run src/\*.go\`

Do you know how can I setup go mod to make it aware of my src directory ? And in a repo, how do you keep source code separated from documentation and other configuration files ?
## [8][Top 5 Frontend Development Outsourcing Challenges And How To Overcome Them](https://www.reddit.com/r/golang/comments/fgcxip/top_5_frontend_development_outsourcing_challenges/)
- url: http://selleooo.xyz
---

## [9][go-pg cursor pagination utility](https://www.reddit.com/r/golang/comments/fg9zry/gopg_cursor_pagination_utility/)
- url: https://www.reddit.com/r/golang/comments/fg9zry/gopg_cursor_pagination_utility/
---
If you’re using 

    go-pg - https://github.com/go-pg/pg

ORM I made a cursor pagination utility - [https://github.com/SeamPay/gopgpager](https://github.com/SeamPay/gopgpager)

    PS: never use offset pagination.
## [10][How to become really good at golang?](https://www.reddit.com/r/golang/comments/ffsaou/how_to_become_really_good_at_golang/)
- url: https://www.reddit.com/r/golang/comments/ffsaou/how_to_become_really_good_at_golang/
---
I have about 1 year experience with the language, already did 2 projects with it (web apps). I still feel that I'm not doing it the 'idiomatic' way, and I didn't use channels, etc.
The thing is I'm developing alone, there's no 'senior' developer to guide me. And I can't find 'advanced golang' courses online, to help me become a really good golang developer. Also I don't think I'm good enough to contribute to open source projects.
Is there a book/course that would help me become a really good golang developer? Thank you
