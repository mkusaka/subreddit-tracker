# golang
## [1][Just because you can, doesn't mean you should.](https://www.reddit.com/r/golang/comments/is81vi/just_because_you_can_doesnt_mean_you_should/)
- url: https://i.redd.it/70cmh75yozm51.jpg
---

## [2][Tutorial: Running Go Code on iOS and Android](https://www.reddit.com/r/golang/comments/isk65w/tutorial_running_go_code_on_ios_and_android/)
- url: https://rogchap.com/2020/09/14/running-go-code-on-ios-and-android/
---

## [3][Can anybody explain me the following video?](https://www.reddit.com/r/golang/comments/isjttb/can_anybody_explain_me_the_following_video/)
- url: https://www.reddit.com/r/golang/comments/isjttb/can_anybody_explain_me_the_following_video/
---
[https://www.youtube.com/watch?v=sjd2ePF3CuQ](https://www.youtube.com/watch?v=sjd2ePF3CuQ)

I'm  a beginner in Go. I dont have much experience with OOPs so i' having difficulty understanding interfaces.

I dont know why but I seriously dont get interfaces. I dont know how to use them correctly. Interfaces really confuses me( also dependency injection using interfaces)

&amp;#x200B;

In the video he explains how to create a struct which implements the interface and now you can access different methods.

&amp;#x200B;

But I could have simple a method and accessed it with that struct instead.

Can anybody please explain in not so beginner friendly way?
## [4][BitMealum: an e2e encrypted email alternative written in Go](https://www.reddit.com/r/golang/comments/isg3a7/bitmealum_an_e2e_encrypted_email_alternative/)
- url: https://www.reddit.com/r/golang/comments/isg3a7/bitmealum_an_e2e_encrypted_email_alternative/
---
I want to introduce BitMaelum ([https://github.com/bitmaelum](https://github.com/bitmaelum)), which is a secure end-to-end email infrastructure. It's a work-in-progress and far from ready for even a beta release, and hopefully I will find some collaborators in both code and architecture.

It tries to combat (many of) the pain-points found in the current email system: spam, privacy, mail/account ownership, etc. I don't expect this project ever to be implemented. Still, I find it's a good exercise to come up with solutions without worrying about the existing email infrastructure. It also is a lovely way to get introduced into a more complex go architecture and code.

It's entirely written in Go and consists of a mail-server (bm-server), a client (bm-client), and configuration tools. It also contains a "key-server" which is a (centralized) system that holds information about addresses and their destinations (a bit like a DNS system), which I'm trying to figure out how to decentralize this (I've looked in detail to blockchains, DHT (ipfs) and others, but it's a challenging problem).

There is no central command (except for the key-server for now), so you are free to use any address you like and even set up organizations where you can add other users. All emails, including meta-data like headers, are fully encrypted securely (as best to my knowledge, so I guess there are many glitches).

This project is my first large project in Go, so I am still struggling a bit with all the Go idioms, and I like to have input from contributors in not only code but also on sharing ideas on how to solve some of the problems I still haven't solved.

Please take a look at [https://github.com/bitmaelum](https://github.com/bitmaelum), and I look forward to collaboration with others!
## [5][Best way to create folder structure of REST service project](https://www.reddit.com/r/golang/comments/isfv2o/best_way_to_create_folder_structure_of_rest/)
- url: https://www.reddit.com/r/golang/comments/isfv2o/best_way_to_create_folder_structure_of_rest/
---
Hi

What is the best way to create folder structure of a REST service?

Thank you.
## [6][Using Dependency Inversion in Go](https://www.reddit.com/r/golang/comments/isk490/using_dependency_inversion_in_go/)
- url: https://medium.com/@cobbinma/using-dependency-inversion-in-go-31d8bf9b3760?source=friends_link&amp;sk=ea07f366baea72eb992f06391301a317
---

## [7][Go-Diagrams: Create architecture diagrams with Go](https://www.reddit.com/r/golang/comments/is8eo9/godiagrams_create_architecture_diagrams_with_go/)
- url: https://github.com/blushft/go-diagrams
---

## [8][Converts 'go mod graph' output into Graphviz's DOT language](https://www.reddit.com/r/golang/comments/irupxd/converts_go_mod_graph_output_into_graphvizs_dot/)
- url: https://www.reddit.com/r/golang/comments/irupxd/converts_go_mod_graph_output_into_graphvizs_dot/
---
[https://github.com/lucasepe/modgv](https://github.com/lucasepe/modgv)

&amp;#x200B;

Graphically view the dependencies of your latest project.

https://preview.redd.it/9t2h8xumnvm51.jpg?width=1280&amp;format=pjpg&amp;auto=webp&amp;s=fa3b692eb4ddfb50f5a3192f7f820f68acc36a05
## [9][Working with Slices in Go (Golang) - Understanding How append, copy and Slice Expressions Work](https://www.reddit.com/r/golang/comments/is42mo/working_with_slices_in_go_golang_understanding/)
- url: http://www.tugberkugurlu.com/archive/working-with-slices-in-go-golang-understanding-how-append-copy-and-slicing-syntax-work
---

## [10][Try to run Jenkins job from the command line](https://www.reddit.com/r/golang/comments/isi5l4/try_to_run_jenkins_job_from_the_command_line/)
- url: https://www.reddit.com/r/golang/comments/isi5l4/try_to_run_jenkins_job_from_the_command_line/
---
Best way to run any Jenkins Job/Build from the Command Line/Terminal without the Web interface: [https://github.com/gocruncher/jbuilder](https://github.com/gocruncher/jbuilder)
