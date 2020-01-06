# golang
## [1][rqlite 5.0.0 released - the lightweight distributed DB built on SQLite, now with new Raft code](https://www.reddit.com/r/golang/comments/ekn0xk/rqlite_500_released_the_lightweight_distributed/)
- url: http://www.philipotoole.com/rqlite-5-0-0-released/
---

## [2][How to build a real-time collaborative text editor using go?](https://www.reddit.com/r/golang/comments/ekt6tp/how_to_build_a_realtime_collaborative_text_editor/)
- url: https://www.reddit.com/r/golang/comments/ekt6tp/how_to_build_a_realtime_collaborative_text_editor/
---
Hi,

I was searching on the google to find any resource that will help me to build a real time collaborative text editor. Such as, ot.js for node JS.

Unfortunately I didn't find anything like this. I want to integrate my own custom text editor front end with golang backend.

&amp;#x200B;

Do you have any suggestions?
## [3][Micro - a modern and intuitive terminal-based text editor](https://www.reddit.com/r/golang/comments/ekfwu9/micro_a_modern_and_intuitive_terminalbased_text/)
- url: https://micro-editor.github.io/
---

## [4][Chime A Go editor for macOS](https://www.reddit.com/r/golang/comments/ekdbiy/chime_a_go_editor_for_macos/)
- url: https://www.chimehq.com/
---

## [5][Go Things I Love: Channels and Goroutines](https://www.reddit.com/r/golang/comments/eku5xl/go_things_i_love_channels_and_goroutines/)
- url: https://www.justindfuller.com/2020/01/go-things-i-love-channels-and-goroutines/?utm_source=reddit&amp;utm_medium=referral&amp;utm_campaign=go_things_i_love
---

## [6][Tutorial Microservices in Go](https://www.reddit.com/r/golang/comments/ekj87w/tutorial_microservices_in_go/)
- url: https://www.reddit.com/r/golang/comments/ekj87w/tutorial_microservices_in_go/
---
Someone knows a nice tutorial of Microservices in go?
## [7][Introducing ZincECS. An entity component system package.](https://www.reddit.com/r/golang/comments/ektb7c/introducing_zincecs_an_entity_component_system/)
- url: https://github.com/SirMetathyst/zinc
---

## [8][Text To Speech in Golang with AWS Polly in 15 minutes](https://www.reddit.com/r/golang/comments/ekt61y/text_to_speech_in_golang_with_aws_polly_in_15/)
- url: https://medium.com/@mssr/text-to-speech-in-golang-with-aws-polly-in-15-minutes-b5f30eb34c1f
---

## [9][Still stuck on apis and channels and threads...](https://www.reddit.com/r/golang/comments/ekq7nk/still_stuck_on_apis_and_channels_and_threads/)
- url: https://www.reddit.com/r/golang/comments/ekq7nk/still_stuck_on_apis_and_channels_and_threads/
---
Posted recently about using channels to wait for an async response from a message bus. I verified that it DOES work.. but I am still not sure it works as I think.

First, the basics of what I am trying to do (as short as I can). Microservices architecture, message bus between services, MQTT protocol with JSON payloads. I have rest APIs that are mostly going to be async based, but in some cases, sync. One such case would be a login flow. This may be an unorthodox way to do this, btw, and I am open to suggestions if there is a better way.. but basically I have a "user" service, which maintains... users (naturally). Each user has an email/password (for login), a role, name, and other various bits of info. On login, the consumer makes an API request.. /login  with email/password has request body. The API handler is a separate microservice.. so it does NOT have access to the user tables. Thus, it needs to send a message to the user service to try to find the user that is logging in. Now, message buses are async, right (unless there is some sync way of doing this?) so the /login API service sends a message to the use service.. and also has a func handler subscribed to a topic that the user service can send back on if it finds the user (or not). Simple enough.. login API sends msg, user service gets it, does some work, sends msg back, login API gets msg, and does something with it.

The problem is that I have this login API sync request.. waiting... while the API sends async message to the use service, and then has to wait on the user service to find the user (or not) and send message back. The real problem is.. having the login handler code "wait" while it waits for the response from the user service. I solved this by creating a channel (on every single request I do this.. which is one of my concerns), then waits on the channel to return some data. The subscribed func when it receives a message from the user service, then writes the user data (assuming found user) to the channel, and the currently waiting channel gets the data and returns the user result to the consumer. 

I actually have this working.. it's pretty neat. 

The concerns. First, creating a channel on each and every request.. is that a bad thing? I am not knowledgeable enough with Go and runtime to know if that ends up being a potential memory hog, cpu hog, both, some sort of unforeseen runtime issue that can cause the API service to crash, etc.

Second, probably the harder issue, is I \*think\* each request is wrapped in a Go func by the standard library (I am using Chi for API/routing/middleware, but it is based on the standard http.\*). I think this is the case. Haven't been able to confirm this, but I assume in order to handle tons of requests simultaneously, it must be doing something like this. That's great.. makes sense. However, what I am unsure of, if this is the case.. is if the channel I create tied to the specific go routine/thread that is handling a specific request, such that when my subscribe func gets  a response from the user service and writes to the channel.. if it is writing to the SAME channel as was created for that specific request?

If it is.. that is fantastic (assuming creating a channel for every request is ALSO ok). However, there remains another issue.. ensuring the response from the user service matches the request lookup (e.g. email/password) before writing the user data to the channel. For this I would need some way, per request/instance/thread to "keep track" of the request email/password... or something like a UUID tied to the creation of the channel and the email/password... to ensure that the correct found user is returned thru the correct channel. Here I am a bit stumped. I am creating the channel in the request handler, but another func that listens for messages from the message bus, writes to the channel. To make this work so far, I am declaring a channel outside of the functions, which seems wrong (I am still trying to learn/understand scope within golang as well). I am worried that this one variable ends up being shared across threads/go routines/requests, and thus is more like a singleton.

Another thought I had was to create an anonymous func IN the req handler to tie it to the channel/email/password data, so that it has access to those values to compare to the values of the user it responds to in a message from the user service. But I also very much dislike this route because if there are 1000 requests running at a time, that would mean every single request would have to potentially process up to 1000 messages until it finds the matching returned/found user matching the request user.

Damn.. this is a pain in the ass. I think the idea of it is awesome. But I am unsure of if it is even possible, or if it is, if it is a potentially bottleneck in terms of memory or performance.

Anyway, would love any help from someone that understands the whole API thread/go routine stuff with creating channels, matching them up correctly (if they dont automatically), etc.

Thank you.
## [10][How to mock OAuth 2.0 in Go](https://www.reddit.com/r/golang/comments/ekpket/how_to_mock_oauth_20_in_go/)
- url: https://blog.seriesci.com/how-to-mock-oauth-in-go/
---

