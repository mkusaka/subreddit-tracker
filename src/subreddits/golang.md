# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][NSQTracer to watch published NSQ messages](https://www.reddit.com/r/golang/comments/jf97h3/nsqtracer_to_watch_published_nsq_messages/)
- url: https://www.reddit.com/r/golang/comments/jf97h3/nsqtracer_to_watch_published_nsq_messages/
---
I thought it would be good if I share some of my CLI based utility for tracing **NSQ** messages built with Golang here. In case someone needs it.

&amp;#x200B;

[Github Preview](https://preview.redd.it/lmhfnnmpveu51.png?width=1113&amp;format=png&amp;auto=webp&amp;s=2f4ef1bb502b17fa578c46a6fa9933eaa17c394c)

[https://github.com/slaveofcode/nsqtracer](https://github.com/slaveofcode/nsqtracer)
## [3][Represent Message Length as a 2-byte binary](https://www.reddit.com/r/golang/comments/jf9xkk/represent_message_length_as_a_2byte_binary/)
- url: https://www.reddit.com/r/golang/comments/jf9xkk/represent_message_length_as_a_2byte_binary/
---
Am doing integration with a third party and the communication is being done over TCP sockets. They have this requirement(**Message variable length indicator**) where am supposed to represent the XML message length as a 2 binary byte message length.

For example, my message is 1024 in length. How do I represent 1024 as binary two-bytes?

&amp;#x200B;

Here is a similar issue in ruby [https://stackoverflow.com/questions/13794817/ruby-how-to-represent-message-length-as-2-binary-bytes](https://stackoverflow.com/questions/13794817/ruby-how-to-represent-message-length-as-2-binary-bytes)
## [4][A simple helper to restart Docker containers with newer versions of images pulled from registry.](https://www.reddit.com/r/golang/comments/jfcus1/a_simple_helper_to_restart_docker_containers_with/)
- url: https://www.reddit.com/r/golang/comments/jfcus1/a_simple_helper_to_restart_docker_containers_with/
---
Helps k8s-less provisioning and updating to the newer version of a stateless container without pains of recalling the command-line options used to start the container.

[https://github.com/jdevelop/repull](https://github.com/jdevelop/repull)
## [5][Best golang framework for REST API?](https://www.reddit.com/r/golang/comments/jf9ih7/best_golang_framework_for_rest_api/)
- url: https://www.reddit.com/r/golang/comments/jf9ih7/best_golang_framework_for_rest_api/
---
Hello everyone, I'm new to golang. The question is in the title. Right now I settled on the Iris framework [https://www.iris-go.com/](https://www.iris-go.com/) because it is the newest and seems to be powerful. What do you use in commercial projects and what do you like the most?
## [6][Go and GUI, what to do?](https://www.reddit.com/r/golang/comments/jeywfv/go_and_gui_what_to_do/)
- url: https://www.reddit.com/r/golang/comments/jeywfv/go_and_gui_what_to_do/
---
I'm toying with the idea of writing a GUI for a terminal application I've created in the past year - what are the best current options for creating a GUI with Go today?
## [7][A golang tool —- for fast generating icon for MacOS App](https://www.reddit.com/r/golang/comments/jf9hu9/a_golang_tool_for_fast_generating_icon_for_macos/)
- url: https://github.com/scott-x/icns
---

## [8][Learn Go test-first with 'For the Love of Go: Fundamentals'](https://www.reddit.com/r/golang/comments/jelx8o/learn_go_testfirst_with_for_the_love_of_go/)
- url: https://bitfieldconsulting.com/books/fundamentals
---

## [9][hashicorp/waypoint: A tool to build, deploy, and release any application on any platform](https://www.reddit.com/r/golang/comments/jeq7c6/hashicorpwaypoint_a_tool_to_build_deploy_and/)
- url: https://github.com/hashicorp/waypoint
---

## [10][Luks.go: pure-Golang implementation of LUKS partition manager](https://www.reddit.com/r/golang/comments/jey9fq/luksgo_puregolang_implementation_of_luks/)
- url: https://www.reddit.com/r/golang/comments/jey9fq/luksgo_puregolang_implementation_of_luks/
---
Hi folks,

I would like to present my project that I was working on for the last couple of months. Luks.go - a pure-Golang library to manage partitions encrypted with LUKS [https://github.com/anatol/luks.go](https://github.com/anatol/luks.go)

If you need to deal with LUKS partition in your golang application then luks.go is your friend. Luks.go allows you to unlock a LUKS partition without using system dynamic libraries or standalone tools like cryptsetup. Pulling such system dependencies is a PITA especially if your tool needs to work across different OS versions.

Currently luks.go supports unlocking only. In the future it might contain operations that modify luks metadata (e.g. adding/removing keyslots). Though this functionality is more dangerous and requires more time for development and testing.

This work has been sponsored by my employer - Twitter. The project has been started as a part of Twitter hackathon week. I want to say thanks to my company and my manager for supporting me in my open-source work.
## [11][Rare realtime log aggregator released 0.1.27 with bug fixes and better histogram support](https://www.reddit.com/r/golang/comments/jf482g/rare_realtime_log_aggregator_released_0127_with/)
- url: https://www.reddit.com/r/golang/comments/jf482g/rare_realtime_log_aggregator_released_0127_with/
---
[https://github.com/zix99/rare](https://github.com/zix99/rare)

This is a golang project I've been working on off-and-on for a while.  Originally, it was to help suite my needs to analyze massive amounts of log files (gigabytes+) and get incremental output, rather than waiting for something like \`zcat | uniq\` to run.  I still use it from time to time, but wanted to share with you all

Here's an example of analyzing nginx log files for statuses:

    ./rare histo -m '" (\d{3})' -e "{bucket {1} 100}" -xz testdata/*
    400                 5,807,761  [69.4%] ||||||||||||||||||||||||||
    200                 2,565,032  [30.6%] |||||||||||
    300                 535        [ 0.0%] 
    
    
    Matched: 8,373,328 / 8,383,717 (Groups: 3)

It supports a range of features similar to other programs you're already used to: gunzipping, regex search, simple handlebars-like expressions, ignore patterns, etc.

Hope you like my project.  Always happy to take feedback!  I went through great pains to optimize this, and learned a lot about go in the process.
