# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][Wombat: Cross Platform gRPC Client. Developed in Go.](https://www.reddit.com/r/golang/comments/jrudsh/wombat_cross_platform_grpc_client_developed_in_go/)
- url: https://i.redd.it/w03bdqwnghy51.jpg
---

## [3][galeone/tfgo: simplified TensorFlow's Go bindings with 2.3 support](https://www.reddit.com/r/golang/comments/js53od/galeonetfgo_simplified_tensorflows_go_bindings/)
- url: https://github.com/galeone/tfgo
---

## [4][a simple question about file server's routing](https://www.reddit.com/r/golang/comments/js7vvr/a_simple_question_about_file_servers_routing/)
- url: https://www.reddit.com/r/golang/comments/js7vvr/a_simple_question_about_file_servers_routing/
---
hello I'm learning net/http package by a book and there's some problem.

I have this structure

    main.go
    - images
       - dog.png

and this code

    func main() {
    	http.Handle("/images", http.FileServer(http.Dir("./images")))
    	http.ListenAndServe(":8080", nil)
    }

There's a page not found error when I go to "localhost:8080/images"

what's the problem??
## [5][Happy Birthday Go!](https://www.reddit.com/r/golang/comments/jrkn2m/happy_birthday_go/)
- url: https://www.reddit.com/r/golang/comments/jrkn2m/happy_birthday_go/
---
Was going through tour of golang  [here](https://tour.golang.org/welcome/4) and found that it's Go's birthday today!

Long live Go!

Edit: A little more details - 

So [this page](https://tour.golang.org/welcome/4) in tour says:

&gt;In the playground the time begins at 2009-11-10 23:00:00 UTC  (determining the significance of this date is an exercise for the  reader).

And a quick search for "2009-11-10" on internet tells that it's Go's birthday :)
## [6][openacid/slim: Surprisingly space efficient trie in Golang(11 bits/key; 100 ns/get).](https://www.reddit.com/r/golang/comments/js5n9d/openacidslim_surprisingly_space_efficient_trie_in/)
- url: https://github.com/openacid/slim
---

## [7][Eleven Years of Go](https://www.reddit.com/r/golang/comments/jrrlk4/eleven_years_of_go/)
- url: https://blog.golang.org/11years
---

## [8][UNIX `touch` command in go](https://www.reddit.com/r/golang/comments/js45ea/unix_touch_command_in_go/)
- url: https://www.reddit.com/r/golang/comments/js45ea/unix_touch_command_in_go/
---
UNIX type operating systems  have a terminal command called `touch` for creating a file or changing its modified time. Unfortunately Windows doesn't have such command.

I needed to have `touch` command in Windows so I decided to build it. It's been 2 days  that I'm learning go, and it is awesome. The go is really good for building cli tools, so I decided to build the create the command in go. Here's the repo

&gt;[https://github.com/spitfire-hash/go-touch](https://github.com/spitfire-hash/go-touch)

I have also wrote the tests for the tool. All the instructions about how to install from release page or building from source is  in the README. I really appreciate anyone who gives feedback about the code, about the code style, or tests itself.

I'm sorry if the showcasing in this subreddit is not allowed, I couldn't find anything about it in the rules section, so thought it's not a problem.
## [9][Simple video streaming service in Golang?](https://www.reddit.com/r/golang/comments/js7tic/simple_video_streaming_service_in_golang/)
- url: https://www.reddit.com/r/golang/comments/js7tic/simple_video_streaming_service_in_golang/
---
Hello, I am guite new in backend .. I do mostly frontend but for my sister I want to build a full stack webpage (react, golang, digital ocean - probably). One of the things she needs is to play a few hundreds of videos (downloaded twitch clips). At first I was thinking of using youtube or vimeo to host the files and provide player + stream service but now I wonder:  if it's just simpliest video streaming possible, could I implement it myself and host the files on my server? I would like to learn new things on this project.   


So I am asking: is this doable for beginner in golang .. or is to too hard? I thought that Go should have some libraries for making these kind of things quite simple. Right? Wrong?
## [10][Get key presses](https://www.reddit.com/r/golang/comments/js7nc6/get_key_presses/)
- url: https://www.reddit.com/r/golang/comments/js7nc6/get_key_presses/
---
I'm looking for a way to get key presses for a program I'm working on. I don't want the user to have to type something then press enter with something like the scanln function. I want it to listen to key presses in the background then react when it sees that a certain key has been pressed.

I found this library [https://github.com/eiannone/keyboard](https://github.com/eiannone/keyboard), but the listening functions it offers are blocking, and there isn't a function for detecting when a key isn't pressed.

I need a non-blocking way of detecting key presses and key releases. Is there a library for doing this? Or a (relatively) simple way of doing it manually?
## [11][Q/A: As a Go developer, would you switch to one of the new M1-based Macs this year?](https://www.reddit.com/r/golang/comments/js4cgg/qa_as_a_go_developer_would_you_switch_to_one_of/)
- url: https://www.reddit.com/r/golang/comments/js4cgg/qa_as_a_go_developer_would_you_switch_to_one_of/
---
My  2015 MBP is showing sings of old age already, so if I am going to  switch, I don't see a reason to do with an Intel-based Mac anymore. My  two options are either M1, or moving to a Linux laptop. Main main  concern with M1 is the availability of developer tooling for ARM.

This is a very good general "2 cents" write-up on yesterday's announcements: [https://sixcolors.com/post/2020/11/enter-the-m1-notes-on-tuesdays-big-event/](https://sixcolors.com/post/2020/11/enter-the-m1-notes-on-tuesdays-big-event/)

What's your take?

[View Poll](https://www.reddit.com/poll/js4cgg)
