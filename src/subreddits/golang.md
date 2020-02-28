# golang
## [1][Actual arbitrary monkeypatching for Go. Yes really.](https://www.reddit.com/r/golang/comments/fatlme/actual_arbitrary_monkeypatching_for_go_yes_really/)
- url: https://github.com/bouk/monkey
---

## [2][Building a Web App With Go, Gin and React](https://www.reddit.com/r/golang/comments/faj4kr/building_a_web_app_with_go_gin_and_react/)
- url: https://www.reddit.com/r/golang/comments/faj4kr/building_a_web_app_with_go_gin_and_react/
---
In this tutorial, I'll show you how easy it is to build a web application with Go and the Gin framework and add authentication to it. 

[https://hakaselogs.me/2018-04-20/building-a-web-app-with-go-gin-and-react/](https://hakaselogs.me/2018-04-20/building-a-web-app-with-go-gin-and-react/)
## [3][A tcp/udp port forwarding/reverse forwarding tool](https://www.reddit.com/r/golang/comments/fashw1/a_tcpudp_port_forwardingreverse_forwarding_tool/)
- url: https://www.reddit.com/r/golang/comments/fashw1/a_tcpudp_port_forwardingreverse_forwarding_tool/
---
I write a tcp/udp port forwarding/reverse forwarding tool and hope it's useful for you :)

[https://github.com/xitongsys/gogw](https://github.com/xitongsys/gogw)
## [4][Does anyone know an good alternative for Scapy in golang?](https://www.reddit.com/r/golang/comments/faqoiu/does_anyone_know_an_good_alternative_for_scapy_in/)
- url: https://www.reddit.com/r/golang/comments/faqoiu/does_anyone_know_an_good_alternative_for_scapy_in/
---
I’m trying to find a framework that allows me to define and parse network packets. Does anyone is working on or using one?
## [5][v1.3 of sessionup - simple and straightforward HTTP session management](https://www.reddit.com/r/golang/comments/fatxgd/v13_of_sessionup_simple_and_straightforward_http/)
- url: https://github.com/swithek/sessionup
---

## [6][nginx and golang](https://www.reddit.com/r/golang/comments/fatt3y/nginx_and_golang/)
- url: https://www.reddit.com/r/golang/comments/fatt3y/nginx_and_golang/
---
We use nginx exclusively. We compile modules such as LDAP [https://github.com/kvspb/nginx-auth-ldap](https://github.com/kvspb/nginx-auth-ldap)

It works but was wondering if there was a golang service which handles authentication and gives the request back to nginx. I like golang service because its a single binary.
## [7][Best practice to Dockerize Golang 1.13](https://www.reddit.com/r/golang/comments/fah0v8/best_practice_to_dockerize_golang_113/)
- url: https://www.reddit.com/r/golang/comments/fah0v8/best_practice_to_dockerize_golang_113/
---
Hi folks, forgive me the question but I am very n00b.

I would like to properly dockerize a sample Go application (just a server for now listening on port 80 and displaying a message). 

I decided to try the multistage build approach inside Docker, so I build everything in a official golang image, then I switch to `scratch` and copy just the compiled binary, using it as `ENTRYPOINT`.

I began though to have some "File not found" errors, so I realized - at least I think - that some dependent libraries were missing. 

So I began to search if there were some proper `go build` parameters that could help me build a self-sufficient binary; but I found very contradictory statements about some environment vars, [workarounds](https://github.com/golang/go/issues/9344) that depend on the version used, ...

So, for specifically version 1.13 of Golang, which is the correct way to build the executable into the container so I can just copy it into a blank image?
## [8][I made a service for scoring go pull requests based on GoReportCard](https://www.reddit.com/r/golang/comments/fauibb/i_made_a_service_for_scoring_go_pull_requests/)
- url: https://www.reddit.com/r/golang/comments/fauibb/i_made_a_service_for_scoring_go_pull_requests/
---
I've made a microservice which checkout go repos and scores the PRs (using github api) based on GoReportCard.

Here is the readme for [README and Examples](https://github.com/gogitops/gogitops-examples) for running it against your own PRs.

I've been just developing this in my spare time and hosting it on gcloud at the moment.

Looking for feedback and opinion :) 

**Example Curl**

curl -X GET \\ 	https://gogitops.beau.cf/submit/$(GITHUB\_REPOSITORY)/pull/$(PR\_ID) \\ 	  
-H 'apikey: RlqchS8EnQ7edvgyNTLIaAdC2yncRQ8v'

**ofcourse theres no warranty or guarantee of service.**

&amp;#x200B;

https://preview.redd.it/gjvt6yptxnj41.png?width=800&amp;format=png&amp;auto=webp&amp;s=f91ac078d14646a472811bab078d1e546b2411fc
## [9][Will there be an official go gui framework?](https://www.reddit.com/r/golang/comments/fabea5/will_there_be_an_official_go_gui_framework/)
- url: https://www.reddit.com/r/golang/comments/fabea5/will_there_be_an_official_go_gui_framework/
---
Will there be an official go standard gui framework? Like tkinter is for python.
## [10][gnark: a fast zero knowledge proof library (written in Go)](https://www.reddit.com/r/golang/comments/fafdj1/gnark_a_fast_zero_knowledge_proof_library_written/)
- url: https://hackmd.io/@zkteam/gnark
---

