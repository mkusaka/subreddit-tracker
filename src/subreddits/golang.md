# golang
## [1][A (toy) web browser written from scratch entirely in golang](https://www.reddit.com/r/golang/comments/gcge4k/a_toy_web_browser_written_from_scratch_entirely/)
- url: https://www.reddit.com/r/golang/comments/gcge4k/a_toy_web_browser_written_from_scratch_entirely/
---
I've started this project to learn GO and to learn how web browsers work.

https://preview.redd.it/rnf59t1ktfw41.png?width=614&amp;format=png&amp;auto=webp&amp;s=0641365d5fdcfd349a9d0569a96be4badfa9014d

This is the project link:

[https://github.com/danfragoso/thdwb](https://github.com/danfragoso/thdwb)

I tried to not use any dependency on this project, and write everything myself.

I wrote the following components:

* ketchup (html parser and DOM Tree builder)
* mayo (css parser and Render Tree builder)
* mustard (UI Toolkit, events and OpenGL)
* sauce (requests, cache and filesystem)
* bun (css layout calculator)

The dependencies are:

* [https://github.com/fogleman/gg](https://github.com/fogleman/gg) (fork)
* [https://github.com/go-gl/gl](https://github.com/go-gl/gl)
* [https://github.com/go-gl/glfw](https://github.com/go-gl/glfw)

It builds and runs on Linux and macOS.

It's is for the most part completely broken. But I would love any kind of feedback and suggestions on the features I should work next.
## [2][Beginners Guide to gRPC in Go!](https://www.reddit.com/r/golang/comments/gc9k6s/beginners_guide_to_grpc_in_go/)
- url: https://youtu.be/BdzYdN_Zd9Q
---

## [3][Building a REST API with Go Gin - Part 1](https://www.reddit.com/r/golang/comments/gci9nf/building_a_rest_api_with_go_gin_part_1/)
- url: https://youtu.be/CzxEUDq9xiQ
---

## [4][A nice simple thread-safe progress bar / spinner (v3)](https://www.reddit.com/r/golang/comments/gc7j0a/a_nice_simple_threadsafe_progress_bar_spinner_v3/)
- url: https://github.com/schollz/progressbar
---

## [5][How to check whether a struct implements an interface in GoLang](https://www.reddit.com/r/golang/comments/gcqhby/how_to_check_whether_a_struct_implements_an/)
- url: https://www.pixelstech.net/article/1588481241-How-to-check-whether-a-struct-implements-an-interface-in-GoLang
---

## [6][Blimp: Cloud Docker Compose Written in Go](https://www.reddit.com/r/golang/comments/gcczf5/blimp_cloud_docker_compose_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/gcczf5/blimp_cloud_docker_compose_written_in_go/
---
I was struggling with how slow my laptop gets when I run Docker with a lot of microservices, so I built a Docker Compose alternative that runs in the cloud rather than on my local machine. It's called Blimp: [https://kelda.io/blimp/](https://kelda.io/blimp/)

* Written entirely in Go: we took advantage of the recently released [Go Docker-Compose spec](https://github.com/compose-spec/compose-go)
* Runs on Kubernetes under the hood. I thought we could just use Kompose, but ends up there’s tons of subtle differences between Kube and Docker Compose that aren't captured in the yaml
* We use an encrypted GRPC tunnel to forward localhost to the cloud environment. It was important to us that it *just worked* without having to change the workflow at all.
* For most volumes, we just use standard Kube volumes. But for bind mounts [we automatically use Syncthing](https://kelda.io/blimp/docs/compose-diff/#syncthing). Traditional volume mounts don't perform well over the WAN, and aren't necessary anyways.

Would love feedback, what do you think??
## [7][OpenCodenames: a Jackbox-style version of Codenames made with React &amp; Go!](https://www.reddit.com/r/golang/comments/gcbgig/opencodenames_a_jackboxstyle_version_of_codenames/)
- url: https://github.com/rodohanna/OpenCodenames
---

## [8][Four Steps to Daemonize Your Go Programs](https://www.reddit.com/r/golang/comments/gcgzwo/four_steps_to_daemonize_your_go_programs/)
- url: https://ieftimov.com/post/daemonizing-golang-programs-by-example/
---

## [9][Connecting ZMQ microservices within AWS using Fargate](https://www.reddit.com/r/golang/comments/gcnn6d/connecting_zmq_microservices_within_aws_using/)
- url: https://www.reddit.com/r/golang/comments/gcnn6d/connecting_zmq_microservices_within_aws_using/
---
I am attempting to use ZeroMQ ([https://github.com/pebbe/zmq4](https://github.com/pebbe/zmq4)) within ECS running on Fargate in awsvpc mode. I have 2 different services, each running its own task and service discovery enabled.

I create my Router and Dealer in a microservice called broker.

        front, _ := zmq.NewSocket(zmq.ROUTER)
        defer front.Close()
        front.Bind("tcp://*:4070")
        
        back, _ := zmq.NewSocket(zmq.DEALER)
        defer back.Close()
        back.Bind("tcp://*:4080")

I then add these 2 sockets to a poller and have a for loop that waits for messages. 

I have a separate microservice that connects to the socket and attempts to send a message to the dealer. I have set service discovery so I assume the address I connect to would be: 

&amp;#x200B;

    &gt; "tcp://broker:4070"
    
    Below is the code from 'serviceA'
    
        func New(ZMQ models.ZMQ) *Requester {
        	s, err := zmq.NewSocket(zmq.REQ)
        	if err != nil {
        		log.Fatalln("shareholder/requester zmq.NewSocket", err)
        	}
        	p := zmq.NewPoller()
        	p.Add(s, zmq.POLLIN)
        
        	log.Println("Requester", ZMQ.Req)
        	err = s.Connect("tcp://broker:4070")
        	if err != nil {
        		log.Print(fmt.Errorf("err is %w", err))
        	}
        
        	req := &amp;Requester{
        		Poller:  p,
        		Retries: 2,
        		Socket:  s,
        		Timeout: time.Duration(time.Minute),
        	}
        	runtime.SetFinalizer(req, (*Requester).Close)
        	return req
        }

I then use the above code to send a message with my socket connection

        _, err := r.Socket.SendMessage(req)

However, my message is never received within my broker service. I can hit my REST APIs on the network with their hostnames I register during service discovery, is there something I am missing here with Fargate/ECS/ZeroMQ???
## [10][Go Cheatsheet / repo](https://www.reddit.com/r/golang/comments/gch6ov/go_cheatsheet_repo/)
- url: https://www.reddit.com/r/golang/comments/gch6ov/go_cheatsheet_repo/
---
Hi. I have made a small repository with a bunch of code sample showcasing Go. Personally sometimes when I want to learn something new, I just want to see a bunch of code samples and compare them to what I already know, from other languages I have mastered. The link is right here: [https://github.com/TopSwagCode/Go.CheatRepo](https://github.com/TopSwagCode/Go.CheatRepo) 

I would love all the feedback that I can get. I am totally new to Go myself and just wanted to make this CheatRepo public, so others also could benefit it. I already have a Roadmap with some ideas to grow the repository with more content.
