# golang
## [1][moshebe/gebug: A tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly.](https://www.reddit.com/r/golang/comments/hv2ncw/moshebegebug_a_tool_that_makes_debugging_of/)
- url: https://github.com/moshebe/gebug
---

## [2][RTSP to WebRTC Server for large scale broadcast benchmarks (written in Go!)](https://www.reddit.com/r/golang/comments/hv2yu4/rtsp_to_webrtc_server_for_large_scale_broadcast/)
- url: https://github.com/pion/rtsp-bench
---

## [3][Useful packages Gophers should know](https://www.reddit.com/r/golang/comments/hv608d/useful_packages_gophers_should_know/)
- url: https://www.golangprograms.com/go-programming-language-packages.html
---

## [4][Did you guys read all of the topics like "Animated GIFS" and "complex numbers" in a book called "The Go progamming Langauge"?](https://www.reddit.com/r/golang/comments/hv6yvq/did_you_guys_read_all_of_the_topics_like_animated/)
- url: https://www.reddit.com/r/golang/comments/hv6yvq/did_you_guys_read_all_of_the_topics_like_animated/
---
New to Go language and I am reading a book "Go Programming Language" by Alan A. A. and Brian W. I know it's popular book among all Golang developers. My question is "Do I need to learn the "Animated GIfs " Four Lissajous figures. " from the book as well as "complex numbers" parts? Are these necessary while doing the development works with GO? Did you guys skip this part or dive into it just to make your language better? Cause I feel it's too complicated and also useless in real development scenario.
## [5][Go compiler doesn't like unused variables](https://www.reddit.com/r/golang/comments/hufqpl/go_compiler_doesnt_like_unused_variables/)
- url: https://i.redd.it/1oobvnrzwub51.jpg
---

## [6][Embedding interface{} in interface](https://www.reddit.com/r/golang/comments/hv6lxr/embedding_interface_in_interface/)
- url: https://www.reddit.com/r/golang/comments/hv6lxr/embedding_interface_in_interface/
---
I was reading some code recently and ran into the following:

    type Expander interface{}
    
    type Foo interface {
        Bar()
        Expander
    }


How would I use the `Expander` concept? I know this isn't best practice but it was interesting.
## [7][Datadog Go API and Contexts](https://www.reddit.com/r/golang/comments/hv2aus/datadog_go_api_and_contexts/)
- url: https://www.reddit.com/r/golang/comments/hv2aus/datadog_go_api_and_contexts/
---
Reading up on [Datadog's announcement of their new official Go API client](https://www.datadoghq.com/blog/java-go-libraries/), and at first blush I like their implementation. I'm intrigued by their use of contexts to pass the credentials in. I've seen similar uses in `context` examples -- and I feel it's following idiomatic Go principles -- but this is the first time I've ever seen `context.WithValue()` used in an official library in the wild.

Curious about what everyone's opinions are on it. Do you see it as a solid pattern when building client libraries?
## [8][Anyone Willing to look over my Twitter Bot](https://www.reddit.com/r/golang/comments/huy9o3/anyone_willing_to_look_over_my_twitter_bot/)
- url: https://www.reddit.com/r/golang/comments/huy9o3/anyone_willing_to_look_over_my_twitter_bot/
---
Hey guys,

So I recently started learning Go.  I'm pretty early on in my programming journey.  Started learning python and have kind of jumped over to Go.  This is basically my first real program written in Go.  Based on what I could find online, I think I setup my project structure correctly.  I still need to add more to the README, but here's the gist of what the bot does.

Makes a daily API request to the NASA Insight Rover data and tweets out the daily weather.  If no new weather is available, it provides a list of the most recent readings.

It also will reply to a user if they tweet at the bot with specific tweets.  Similar style to the above sentence.

I've also stumbled through Docker and AWS and was able to get it deployed onto the free tier of an EC2 instance, so it is currently up and running.

I'm still trying to gain an understanding of pointers and such.  

Appreciate any criticism.  Here is the link to my github repo.

[Mars Weather Bot Github Repo](https://github.com/zacharygilliom/marsweatherbot)
## [9][How wrong is this?](https://www.reddit.com/r/golang/comments/hv608o/how_wrong_is_this/)
- url: https://www.reddit.com/r/golang/comments/hv608o/how_wrong_is_this/
---
Hi Gophers,

I'm trying to find a way to stop cluttering my functions and methods with all that error checking.

I know there are (in)famous articles about how errors are values and such, but there was still functionality I was missing there.

What I came up with is the following, and I am wondering how bad it is really that I am basically throwing panics all over the place for errors, however, I am recovering them and handling the error states.

[https://github.com/TheApeMachine/errnie](https://github.com/TheApeMachine/errnie)

I welcome any feedback before I continue working on this or even worse, actually use it.
## [10][Is there a pm2 like process manager to manage applications written in golang?](https://www.reddit.com/r/golang/comments/hv5i6f/is_there_a_pm2_like_process_manager_to_manage/)
- url: https://www.reddit.com/r/golang/comments/hv5i6f/is_there_a_pm2_like_process_manager_to_manage/
---
By managing, I mean daemonizing the process so that it starts when the system reboots, starting and stopping the process, see how much memory it consumes. I've tried [apm](https://github.com/topfreegames/apm) and [pmgo](https://github.com/struCoder/pmgo) but none of them worked properly or I was unable to get them to work. I follow the instructions to start the app using the process manager but the app does not start. There is not a single log telling that some of the things have gone wrong as well.
