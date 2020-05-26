# golang
## [1][SpaceX Dragon ISS docking autopilot using Go and WebAssembly, well documented](https://www.reddit.com/r/golang/comments/gqk61u/spacex_dragon_iss_docking_autopilot_using_go_and/)
- url: https://github.com/mbertschler/dragon-iss-docking-autopilot
---

## [2][Instrumentation in Go](https://www.reddit.com/r/golang/comments/gquh39/instrumentation_in_go/)
- url: https://gbws.io/articles/instrumentation-in-go/
---

## [3][HTTP how to stream download response](https://www.reddit.com/r/golang/comments/gqw3yg/http_how_to_stream_download_response/)
- url: https://rafallorenz.com/go/go-http-stream-download/
---

## [4][What does the Go community think of Elixir?](https://www.reddit.com/r/golang/comments/gqsoq1/what_does_the_go_community_think_of_elixir/)
- url: https://www.reddit.com/r/golang/comments/gqsoq1/what_does_the_go_community_think_of_elixir/
---
It's a rather weird and perhaps, subjective question, but not without its grounds. I have been using both languages for quite some time, and like each one for its own merit. I think, there are many things that Go developers can learn from the Erlang/Elixir ecosystem.  


Are there other people in this community who equally like both languages? How do you distinguish what to use each for?
## [5][What is a goroutine? And what is their size?](https://www.reddit.com/r/golang/comments/gqat1d/what_is_a_goroutine_and_what_is_their_size/)
- url: https://tpaschalis.github.io/goroutines-size/
---

## [6][Webassembly demos using the Pigo face detection library](https://www.reddit.com/r/golang/comments/gqvj3s/webassembly_demos_using_the_pigo_face_detection/)
- url: https://github.com/esimov/pigo-wasm-demos
---

## [7][What is the proper way to get zeromq up and running in Go?](https://www.reddit.com/r/golang/comments/gqv77p/what_is_the_proper_way_to_get_zeromq_up_and/)
- url: https://www.reddit.com/r/golang/comments/gqv77p/what_is_the_proper_way_to_get_zeromq_up_and/
---
**Platform: Windows 10 x64**

Based on [this website](https://zeromq.org/languages/go/), I want to use **zeromq**.

Here is what I have done till now:

I installed **msys2** and in **MSYS2 MSYS** command window, issued the following command:

    pacman -Syuu

Then in the **MSYS2 MinGW 64-bit** command window:

    pacman -Sy libtool

then added ***"C:\\msys64\\mingw64\\bin"*** to the Path and finally in a command prompt window:

    go get "github.com/zeromq/goczmq"

This gave me this error:

    # github.com/zeromq/goczmq
    Documents\Golang\GoWorkspace\src\github.com\zeromq\goczmq\auth.go:4:10: fatal error: czmq.h: No such file or directory
        4 | #include "czmq.h"
          |          ^~~~~~~~
    compilation terminated.

So I downloaded **zeromq** from [here](https://github.com/zeromq/czmq/releases) copied all the contents of its **include** folder into the ***"Documents\\Golang\\GoWorkspace\\src\\***[***github.com***](https://github.com)***\\zeromq\\goczmq\\"*** path and got this error:

    # github.com/zeromq/goczmq
    In file included from ./czmq_library.h:23,
                     from ./czmq.h:37,
                     from Documents\Golang\GoWorkspace\src\github.com\zeromq\goczmq\auth.go:4:
    ./czmq_prelude.h:217:10: fatal error: zmq.h: No such file or directory
      217 | #include "zmq.h"
          |          ^~~~~~~
    compilation terminated.

this time I downloaded the  **ZeroMQ** library from [here](https://zeromq.org/download/) ([Visual Studio 15 2017 x64](https://dl.bintray.com/zeromq/generic/libzmq-v141-x64-4_3_2.zip)) and copied all of its content into the ***"Documents\\Golang\\GoWorkspace\\src\\***[***github.com***](https://github.com)***\\zeromq\\goczmq\\"*** path and got:

    # github.com/zeromq/goczmq
    C:/msys64/mingw64/bin/../lib/gcc/x86_64-w64-mingw32/10.1.0/../../../../x86_64-w64-mingw32/bin/ld.exe: cannot find -lsodium
    C:/msys64/mingw64/bin/../lib/gcc/x86_64-w64-mingw32/10.1.0/../../../../x86_64-w64-mingw32/bin/ld.exe: cannot find -lzmq
    C:/msys64/mingw64/bin/../lib/gcc/x86_64-w64-mingw32/10.1.0/../../../../x86_64-w64-mingw32/bin/ld.exe: cannot find -lczmq
    collect2.exe: error: ld returned 1 exit status

I was able to fix the missing -lsodium by using:

    pacman -Ss libsodium
    
    and then:
    
    pacman -S mingw64/mingw-w64-x86_64-libsodium

but I still get the error of missing **-lczmq** and **-lzmq**. How can I install those packages? Actually what is the proper way for installing and using packages such as **zeromq** in Go?
## [8][Which Go telegram bot framework to use?](https://www.reddit.com/r/golang/comments/gqrp0g/which_go_telegram_bot_framework_to_use/)
- url: https://www.reddit.com/r/golang/comments/gqrp0g/which_go_telegram_bot_framework_to_use/
---
I've looked at both

1. [https://github.com/go-telegram-bot-api/telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api)
2. [https://github.com/tucnak/telebot](https://github.com/tucnak/telebot)

The first one seems to have minimal abstraction (too bare) while the second one seems to rely too much on interfaces (possible runtime type safety issues). Does anyone have any experience with either of these frameworks and which one is more reliable?
## [9][What are some good tips for keeping an eye on blocked / stuck goroutines?](https://www.reddit.com/r/golang/comments/gqvt1f/what_are_some_good_tips_for_keeping_an_eye_on/)
- url: https://www.reddit.com/r/golang/comments/gqvt1f/what_are_some_good_tips_for_keeping_an_eye_on/
---
As applications grow, so does their complexity, and with it, the traceability of things running in parallel Go's race detector is good for keeping an eye on shared memory touched by multiple goroutines. It won't however, report on goroutines stuck on a channel that never gets emptied. In large concurrent applications, this use case is fairly common. Unless one adds log traces on each step, it is fairly easy to get a running app (no compile or runtime errors), which does nothing, because some goroutine somewhere gets stuck, and with it, a whole chain of actions.

What would be some cool tips for instrumenting Go code, such that stuck long-running goroutines can be pointed out immediately? The only thing I can think of right now is pprof
## [10][consuming grpc service using go (re-post from stackoverflow) - https://stackoverflow.com/questions/62021444/consuming-grpc-service-using-go](https://www.reddit.com/r/golang/comments/gqvqjb/consuming_grpc_service_using_go_repost_from/)
- url: https://www.reddit.com/r/golang/comments/gqvqjb/consuming_grpc_service_using_go_repost_from/
---
 

I plan to use grpc as an inter-service sync communication protocol. There are lots of different services and I have generated a pb.go  
 file with all the relevant code for client and server using protoc  
 with the go-rpc plugin.

Now I'm trying to figure out the best way or the common way of consuming this service from another service. Here is what I have so far:

**Option 1**

1. use the .proto  
 file from the service (download it)
2. run the protoc  
 compiler and generate the ...pb.go  
 file for the consumer to use

**Option 2**

1. because the ...pb.go  
 is already generated on the grpc service side to implement the server and my client is another service written in go I can expose this as a sub module (another .mod  
 file in a sub- directory)
2. use go get github.com/usr/my-cool-grpc-service/client

Option 2 seems more appealing to me because it makes the consumption of a service very easy and available for all other services that may require it.

On the other hand I know that the .proto  
 file is the contract that cann generate clients for many different languages and should be used the the source of truth.

I fear that by choosing option 2 I'm unaware of any possible pitfalls I might encounter with regards to backwards compatibility or any other topic..

So, what is the idiomatic way of consuming a gRPC service?
