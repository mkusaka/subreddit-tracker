# golang
## [1][[Q&amp;A] io/fs draft design](https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/
---
I posted a draft design today for new file system interfaces for Go.

Video: https://golang.org/s/draft-iofs-video

Design: https://golang.org/s/draft-iofs-design

Let's do the Q&amp;A about the design here in Reddit. My hope is that the threading support will help keep questions and answers matched.

Please start a new top-level comment for each new question.

See also the related [Q&amp;A for the //go:embed draft design](https://golang.org/s/draft-embed-reddit).
## [2][[Q&amp;A] //go:embed draft design](https://www.reddit.com/r/golang/comments/hv96ny/qa_goembed_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hv96ny/qa_goembed_draft_design/
---
I posted a draft design today for embedding files into Go programs.

Video: https://golang.org/s/draft-embed-video

Design: https://golang.org/s/draft-embed-design

Let's do the Q&amp;A about the design here in Reddit. My hope is that the threading support will help keep questions and answers matched.

Please start a new top-level comment for each new question.

See also the related [Q&amp;A for the io/fs draft design](https://golang.org/s/draft-iofs-reddit).
## [3][How to deploy a Go backend with a React frontend separately on Kubernetes with (silly) drawings!](https://www.reddit.com/r/golang/comments/hworna/how_to_deploy_a_go_backend_with_a_react_frontend/)
- url: https://skarlso.github.io/2020/07/23/kubernetes-deploy-golang-react-apps-separately-part1/
---

## [4][Boilerplate on the top of Fiber Web framework](https://www.reddit.com/r/golang/comments/hwz1y3/boilerplate_on_the_top_of_fiber_web_framework/)
- url: https://www.reddit.com/r/golang/comments/hwz1y3/boilerplate_on_the_top_of_fiber_web_framework/
---
I'm new to Golang world. While playing with different libraries and tools, I've created following boilterplate on that top of Fiber Framework [https://github.com/gofiber/fiber](https://github.com/gofiber/fiber)

Please use this boilerplate for any purpose and let me know if anyone have suggestions [https://github.com/itsursujit/fiber-boilerplate](https://github.com/itsursujit/fiber-boilerplate)

&amp;#x200B;

Any PRs are welcome
## [5][Me and my friend built a self-host monitoring, alert and incident system. We would love to hear some feedback from you!](https://www.reddit.com/r/golang/comments/hwpbss/me_and_my_friend_built_a_selfhost_monitoring/)
- url: https://www.reddit.com/r/golang/comments/hwpbss/me_and_my_friend_built_a_selfhost_monitoring/
---
Hello everyone. We want to show Reddit community our full open-source (and free) self-host monitoring, incident, and alerting system. We have chosen "Squzy" as a product name.

**Our goals were to**:

1. Provide amazing and free monitoring complex for a lot of cases
2. Provide product which doesn't collect data for external resources
3. Support by open-source community
4. Open protocol for support user own realisations/plugins

[**GITHUB**](https://github.com/squzy/squzy)

[**DEMO**](https://demo.squzy.app/)

[**DOCS**](https://squzy.app/)
## [6][Getting an error when trying to test with t.Errorf()](https://www.reddit.com/r/golang/comments/hx1pij/getting_an_error_when_trying_to_test_with_terrorf/)
- url: https://www.reddit.com/r/golang/comments/hx1pij/getting_an_error_when_trying_to_test_with_terrorf/
---
Here is my code ... [https://github.com/jsquardo/GoCards](https://github.com/jsquardo/GoCards) 

Whenever I run the test I get this error... 

    Errorf call has arguments but no formatting directives
    

This is my first day learning Go so I'm sure its something stupid I am missing. Thanks in advance.
## [7][proposal: Go 1 : accept multiple return values as the rest arguments · Issue #40387 · golang/go · GitHub](https://www.reddit.com/r/golang/comments/hx1ace/proposal_go_1_accept_multiple_return_values_as/)
- url: https://github.com/golang/go/issues/40387
---

## [8][[Advice] Rewrite partly finished project in go?](https://www.reddit.com/r/golang/comments/hx17wm/advice_rewrite_partly_finished_project_in_go/)
- url: https://www.reddit.com/r/golang/comments/hx17wm/advice_rewrite_partly_finished_project_in_go/
---
Hi there community,

this is my first post in /r/golang and I am quite new to go. I hope that some of you guys might be able to help me out a little with a decision. So this is kind of a combined question and I am also looking to get some insights into professional software programming strategies and design decisions. 

I have written a darts scoreboard application a time ago using python with flask and jinja2 templates as a viewing engine. That project is quite old and I decided to rewrite it using modern technology. Therefore I learned react and decided to use a combination of a RESTful API design and graphql for transfering big data junks to my frontend fpr easier parsing.

As the backend also needs to integrate the game logic the software is quite big afterwards and not that easy to handle an maintaine.

I do love the fact, that go is easy to learn (I did a few programs in go already) and that it is simple and straight forward. Also it is compilable which means I could release single file applications at the long run.

So do you guys think it might be worth it to start over using go as a backend language (which works quite well with react) and build out my API there?

I did not do that much yet, so there will be not much time lost by restarting once again.

Also I would like your opinion on design. To get you a little more detail information I want to point out what the software will do.

One is able to configure and start a game of darts (add player with profile and picture, choose game type and player playing, start the game).

One Index page will then be redirected to a scoreboard page and one admin page will be redirected to a controller page. The scoreboard page shows the game stats to the players playing darts and the controller page is there to add throws in. (If you do not use a recognition software which will send the api requests automatically).

So the backend needs to fetch the throws (/throw/20/1 - is single 20, /throw/20/2 is double twenty ..) and then run some logic.

If you know a game of darts there is quite a lot to handle. Winning conditions, Bust, etc... And I want to store the game state in a database (sqlite or something like that).

How would you guys approach this? Cause right now I have a hybrid API running which will add throws but also handle gameplay.

There are other difficulties like undoing falsly booked throws or changing to the next player and so on.

So I am quite unsure if I do that in a well design and if this will even be possible or advisable to do in go.

I will leave my "old" scoreboard code ([https://github.com/patrickhener/dart-o-mat-3000](https://github.com/patrickhener/dart-o-mat-3000), [https://dart-o-mat-3000.readthedocs.io/en/latest/](https://dart-o-mat-3000.readthedocs.io/en/latest/)) here. The new one is intended to work like the old one does if my design concept is good. Otherwise I am grateful and thankful for every guidance I can receive from you guys.

If someone wants to elaborate more on how to do it right I can also add you to the github repo for the new software approach. Just shoot me a dm.

Thanks in advance,  
best regards,  
Patrick
## [9][armon/go-metrics: instrument code, expose application metrics, and profile runtime performance in a flexible manner](https://www.reddit.com/r/golang/comments/hx0k21/armongometrics_instrument_code_expose_application/)
- url: https://www.reddit.com/r/golang/comments/hx0k21/armongometrics_instrument_code_expose_application/
---
[https://github.com/armon/go-metrics](https://github.com/armon/go-metrics)
## [10][Authentication in Golang with JWTs](https://www.reddit.com/r/golang/comments/hx06h4/authentication_in_golang_with_jwts/)
- url: https://auth0.com/blog/authentication-in-golang/
---

## [11][fullstorydev/grpcurl: a command-line tool that lets you interact with gRPC servers. It's basically curl for gRPC servers](https://www.reddit.com/r/golang/comments/hwbmdx/fullstorydevgrpcurl_a_commandline_tool_that_lets/)
- url: https://www.reddit.com/r/golang/comments/hwbmdx/fullstorydevgrpcurl_a_commandline_tool_that_lets/
---
[https://github.com/fullstorydev/grpcurl](https://github.com/fullstorydev/grpcurl)
## [12][Monolithic Component Architecture in Go](https://www.reddit.com/r/golang/comments/hwfo4v/monolithic_component_architecture_in_go/)
- url: https://www.reddit.com/r/golang/comments/hwfo4v/monolithic_component_architecture_in_go/
---
Caveat: I'm a long-term Java developer, so a lot of my thinking is coloured by that.

I'm coming back to Go and this time things are making a lot more sense to me - yay :)

What I'm trying to do right now is work out a sensible architecture for my project. It's a monolithic web app (monolith because I want to only install one thing), and is going to have a series of components that work together. As such, things that I want to be able to do are:

* Each component will have an internal API that is the only way for other components to interact with it.
* Many components - though not all - will expose HTTP endpoints for the public API.

So, for example, the `users` component might have use cases for: "GetUser", "UpdateUser" and "CreateUser". The first two of these are available both on the internal and HTTP API, but the third is only on the internal API, because it's only ever called by the code that logs in via Google.

So - presumably what I want is for each component to be it's own package. And for each component to only have as public the bits that are on the internal API. So in this case I would have as public: `GetUserUseCase`, `UpdateUserUseCase`, `CreateUserUseCase`, `UserID`, `UserData` and `UserModel`. The actual service that implements those use cases, the repositories that back it, and the HTTP stuff are all private to the package.

I can then repeat this pattern for each component as needed and it all works.

Is this reasonable? Is this in any way Go-like? Or is this too coloured by my Java experience?

I'm particularly concerned that this will mean each package is going to have files covering many different concerns - API, Service Layer, Repositories, HTTP Endpoints, etc - and that's going to make it harder to work with.

Or is there a more Go-like architecture that's recommended for something like this? Is there a standard way that this kind of thing is/should be structured to better fit with how Go works?

Cheers
