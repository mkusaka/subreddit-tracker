# golang
## [1][New to Go, not to programming: tutorial frustration](https://www.reddit.com/r/golang/comments/ejtm8n/new_to_go_not_to_programming_tutorial_frustration/)
- url: https://www.reddit.com/r/golang/comments/ejtm8n/new_to_go_not_to_programming_tutorial_frustration/
---
I've been programming for years in multiple languages, and I'm frustrated with the Go tutorials. They all teach the same thing: syntax. After learning the pieces (loops, slices, structs, etc.) in the end I'm left without knowing how to write an application. 

I have yet to see anything in the tutorials that is more than a single file,  package main,  func main. If I search for packages, I see nothing but importing packages into that single file. 

Am I missing something? Where can I find info on creating multi file applications? Code organization? I really want to delve into this but the tutorials are too basic to apply to useful functionality.

Am I not getting far enough? Writing switch statements and for loops is not very satisfying.
## [2][Going all in with Golang from PHP. Am I too ambitious ?](https://www.reddit.com/r/golang/comments/ejjppg/going_all_in_with_golang_from_php_am_i_too/)
- url: https://www.reddit.com/r/golang/comments/ejjppg/going_all_in_with_golang_from_php_am_i_too/
---
Sorry couldn't really construct the title well but this is not meant to be clickbaity. Our current SAAS product is PHP based and we are a 100% PHP team at the moment. I am in love with Golang and being the founder of the company, I decided to rewrite in Go due to the following reasons:

\- Better performance. We sell a platform and performance is very important.

\- Ease of deployment. Go Binaries are awesome.

\- I want to find more technical partners who can promote. Using Golang, I can just give them the binary to self host wherever.

\- Our current PHP stack is too bloated and since our platform can be customized for each client, the frameworks and bloat get in our way. It will be easier to have a slimmed down version and then customize on top for each client. Go's simplicity has won me over.

\- We need to remove/simplify some features that are not really needed or are coded in a complex way. 

\- Better suited to build REST APIs. I don't hate PHP (it has paid the bills) but I really like how simple Golang is with its http library

\- Attract better developers ? This one has been so controversial. I want to attract better talent. Does that help ? Or I am stupid.

&amp;#x200B;

The cons:

\- No one in our dev. team knows Golang. I  finally broke the news to my CTO and he was like WTF. He is very good at what he does and comes from a .NET background. So he can pick this up in 2-3 weeks BUT our team is mostly junior devs and they will need to learn Golang.

\- Existing customers on the PHP platform will need to be supported forever. We cannot retire the legacy platform that quickly due to the complex functions built for clients.

Am I committing business suicide ?

&amp;#x200B;

EDIT: Wow, so many great responses. I do appreciate everyone and I will think hard before putting our team at risk. It seems like the consensus is to not do a full re-write but consider chipping away slowly and may be build a few smaller components in Go as microservice if at all.
## [3][Running A/B test](https://www.reddit.com/r/golang/comments/ejur6h/running_ab_test/)
- url: https://www.reddit.com/r/golang/comments/ejur6h/running_ab_test/
---
I am not sure if it is right sub to ask this.

I want to perform A/B test with my go gRPC server so that

1. Code does not become messy
2. Can edit/add or stop the A/B test without deployment.

I cannot use service proxy like Envoy or any traffic management for A/B test. Also, I have already created a separate server that add header to every request about the variant details. 

Solution that I have thought:

1. According to the request header, add a key value in context where key is some flag. (used [https://github.com/mwitkow/go-flagz](https://github.com/mwitkow/go-flagz)). 
2. Create an informationMap that will be loaded from configMap in kubernetes. So, I can edit, add, stop the A/B test. 
3. I have assumed that we have to use if/else for checking flag from context. 

Can anyone suggest/ review this solution ?
## [4][Using local development modules without pushing to Github everytime](https://www.reddit.com/r/golang/comments/ejsgl0/using_local_development_modules_without_pushing/)
- url: https://www.reddit.com/r/golang/comments/ejsgl0/using_local_development_modules_without_pushing/
---
So I am trying to learn the new module system where I put my codebase outside of GOPATH and I am trying to figure out if my development flow is incorrect or not.

E.g. I developed this package/library: https://github.com/psykidellic/gomath. I checkout the repo $HOME/projects/gomath and start hacking on it.

Now, I want to use this library in another of my side project $HOME/projects/"gomathtest". Now if I push changes from gomath to github and use it in my new project 

    import (
        "github.com/psykidellic/gomath/calc"
    )

everything works.

But this means, during development I have to constantly push to github. Is there a way I can use the library gomath in gomathtest with the same import but then use the code from my local projects folder during development?

In my final ci/cd, obviously I want to do the correct way where code is actually imported from github.
## [5][Go: g0, Special Goroutine](https://www.reddit.com/r/golang/comments/ejjglc/go_g0_special_goroutine/)
- url: https://medium.com/a-journey-with-go/go-g0-special-goroutine-8c778c6704d8
---

## [6][Im at a point in learning go where I just don't understand it anymore (Dependency Injection) what is the wisest move to do?](https://www.reddit.com/r/golang/comments/ejn44m/im_at_a_point_in_learning_go_where_i_just_dont/)
- url: https://www.reddit.com/r/golang/comments/ejn44m/im_at_a_point_in_learning_go_where_i_just_dont/
---
Would reading the previous chapters in the tutorial I follow work or is it just a matter of time and excercise?
## [7][koanf: Light weight extensible library for reading config (file, S3 etc.) in Go applications. Built in support for JSON, TOML, YAML, env, command line.](https://www.reddit.com/r/golang/comments/ejjyz8/koanf_light_weight_extensible_library_for_reading/)
- url: https://github.com/knadh/koanf
---

## [8][Where to place project source code?](https://www.reddit.com/r/golang/comments/ejkbx0/where_to_place_project_source_code/)
- url: https://www.reddit.com/r/golang/comments/ejkbx0/where_to_place_project_source_code/
---
Hi, newbie Gopher here. I'm starting to learn Go for fun and I'm a bit confused about the relationship between the GOPATH directory and where my project source code should live. From the Go wiki on GitHub, I see that:

_The GOPATH environment variable is used to specify directories outside of $GOROOT that contain the source for Go projects and their binaries._

Does this include my own projects? Am I breaking a convention if my Go source lives in some other directory? I'm using Go `1.13.5`
## [9][Chaos-mesh: A Chaos Engineering Platform for Kubernetes](https://www.reddit.com/r/golang/comments/ejbgmn/chaosmesh_a_chaos_engineering_platform_for/)
- url: https://github.com/pingcap/chaos-mesh
---

## [10][Using channel to wait for message bus response.. good or bad idea?](https://www.reddit.com/r/golang/comments/ejk9wr/using_channel_to_wait_for_message_bus_response/)
- url: https://www.reddit.com/r/golang/comments/ejk9wr/using_channel_to_wait_for_message_bus_response/
---
TL;DR Right way to have API endpoint send async message across message bus, wait for async msg resp, before sending response back to consumer of API, using go func and channels while handling threaded API requests

Still learning/experimenting. Looking to understand what is considered good/bad (regardless if it works).

Most of my APIs will be async by default, but some, like the web UI waiting on the response from an API in a sync manner, that on the server side uses a message bus to communicate between the API service (the UI called) and the back end service that does some logic and eventually (but most likely within milliseconds) responds back to the API (via another message).

So, to facilitate a mostly async architecture, using a message bus  and some form of consumer event/push notifications is my preferred architecture. With Web hooks, web sockets, and consumers that run in servers being able to provide callback endpoints.. it's trivial to do for the most part. Regardless of the way the consumer, be it another server or a UI, retrieves push notifications.. sometimes a consumer can't handle push (or is not configured to do so, etc) and needs a sync API call.. or maybe it just makes sense in some cases NOT to add the complexity of matching up an async request to an async push notification when the flow dictates that it is sync.

One example is email verification. When a user gets an email with a link to verify their email.. and clicks it, it takes them to my site, which in turn, before the user sees the "next step" makes a SYNC API request to my server to let the server know the user has indeed verified their email. This on the back end, hits the API e.g. /verifyEmail endpoint. At this point, my back end user service (let's not argue/talk about how I have services set up for the sake of this post) is tied to the message bus via a few topics. One is user.verify, and the handler there takes in the email, looks up the user by email (at this point the user is nothing more than just an email, waiting on verification, before the UI next step allows the user to fill in username, password, details, etc).. finds the email (hopefully), and then returns a response via a message BACK to the API /verifyEmail endpoint.

SO.. it is this I wanted to discuss.. or rather ask if the way I think I should implement this is good, bad, if it works so be it, etc.

I am still trying to grasp/understand go funcs and channels. My thought is this. The API endpoint /verifyEmail sends a message (to the user.verifyEmail as I said), and then waits on the response of a channel. So prior, it creates a new channel, then does the &lt;- waitOnResp (or whatever the name is). In the same API endpoint, in the init() func presumably, I have set up a subscription to the message bus, to topic user.verifyEmailResp.  Now, the handler, when it receives a message, would then write the resp payload of the message to the channel... which I assume, if I understand it right, would then cause the currently waiting API endpoint handler to get the data, and then be able to send a response back.. right?

The one issue I am unsure of is an API handler needs to handle any number of incoming threads. So I assume I have to put this bit of code.. the part that sends the msg, then waits on the channel, in a go func(). But.. I am unsure if it will keep the web socket from the initial API request open.. so that when it receives the data from the channel and writes the response back out.. if it will go back to the API consumer (my UI that made the request).

Hence.. is what I am trying to do here the "right way", and more so, if so, how do I properly set it up so it can continue to handle any number of API requests (because one day I hope to have millions of users on my web site :D) and properly send responses back to the write sockets.. or rather hopefully there isn't some way to screw things up in code to cause a response to go back to the wrong user.

It is a bit winded.. hopefully it makes sense?

Appreciate/thank you for any help.
