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
## [2][Design Draft: First Class Fuzzing](https://www.reddit.com/r/golang/comments/hvpr96/design_draft_first_class_fuzzing/)
- url: https://go.googlesource.com/proposal/+/refs/heads/master/design/40307-fuzzing.md
---

## [3][lockgate, a locking library for Go, has been extended with locks using a HTTP server](https://www.reddit.com/r/golang/comments/hzxwl9/lockgate_a_locking_library_for_go_has_been/)
- url: https://github.com/werf/lockgate
---

## [4][zeebo/xxh3 - A port of the extremely fast non-cryptographic hash xxh3 (xxhash) to Go.](https://www.reddit.com/r/golang/comments/hzzlzr/zeeboxxh3_a_port_of_the_extremely_fast/)
- url: https://github.com/zeebo/xxh3
---

## [5][How well do you know Go history?](https://www.reddit.com/r/golang/comments/hzfksf/how_well_do_you_know_go_history/)
- url: https://i.redd.it/asmsp647yld51.png
---

## [6][Let's build a Full-Text Search engine](https://www.reddit.com/r/golang/comments/hzosh9/lets_build_a_fulltext_search_engine/)
- url: https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/
---

## [7][[Hiring] SRE with GoLang - NYC](https://www.reddit.com/r/golang/comments/i00ohv/hiring_sre_with_golang_nyc/)
- url: https://www.reddit.com/r/golang/comments/i00ohv/hiring_sre_with_golang_nyc/
---
This NYC food/meal service company is looking for a local Mid Level Site Reliability Engineer to the team. This role will be remote until it is safe to return to the office in Manhattan - looking for local NYC candidates.

You will be joining a nimble and impressive team of engineers. The ideal candidate will be proficient in Golang and have experience with Docker and CI/CD tools (ie Jenkins, TeamCity, CircleCI, etc). You will also be leveraging Prometheus and Grafana.

Interviews are available for this week :)

Required Skills &amp; Experience

* 3+ years of experience
* Proficiency in GoLang
* Docker/containerization
* CI/CD - Jenkins, TeamCity, etc

The Offer

* Competitive Salary: $130k/year, DOE

You will receive the following benefits:

* Full benefits – Medical, Dental, Life, HSA
* 401(k)
* PTO
* Pre-tax Commuter Benefit
* And more!

Sound interesting? Have questions? Reach out to me at [annelise.hudson@workbridgeassociates.com](mailto:annelise.hudson@workbridgeassociates.com)

Applicants must be currently authorized to work in the United States on a full-time basis now and in the future.
## [8][how to extract data from microsoft visio with go?](https://www.reddit.com/r/golang/comments/hzz9lu/how_to_extract_data_from_microsoft_visio_with_go/)
- url: https://www.reddit.com/r/golang/comments/hzz9lu/how_to_extract_data_from_microsoft_visio_with_go/
---
Is there any way to fetch data from microsoft visio (without visio application), it doesn't have to be go.
## [9][Run go-git commands over ssh](https://www.reddit.com/r/golang/comments/hzz6os/run_gogit_commands_over_ssh/)
- url: https://www.reddit.com/r/golang/comments/hzz6os/run_gogit_commands_over_ssh/
---
Hello !

I recently started using [go-git](https://github.com/go-git/go-git) open source project to handle my git commands, and I'm having an issue figuring how I could use it on a remote server (over ssh).

I used to do my git work using \`[golang.org/x/crypto/ssh\`](https://golang.org/x/crypto/ssh`) package by simply running raw git commands over SSH session. And it would be truly awesome if I could use the go-git package to do the exact same work I do locally on my remote machine.

For example, I would use this (simplified) code to clone a repository to my local machine :

    _, err := git.PlainClone("/tmp/myrepo", false, &amp;git.CloneOptions{
        URL: "https://github.com/user/repo",
    })

How could I hack the thing to make it work over ssh ?

Thanks alot guys
## [10][gRPC vs Websocket performance](https://www.reddit.com/r/golang/comments/hzy68w/grpc_vs_websocket_performance/)
- url: https://www.reddit.com/r/golang/comments/hzy68w/grpc_vs_websocket_performance/
---
I'm still bad with gRPC and I wonder is anyone made benchmark to compare gRPC stream and websocket and which one is faster?
## [11][Server-Sent Event using Kubernetes sidecar to enable HTTP2](https://www.reddit.com/r/golang/comments/hzxooc/serversent_event_using_kubernetes_sidecar_to/)
- url: https://link.medium.com/d2rTHPuOv8
---

## [12][Password Handling Practices](https://www.reddit.com/r/golang/comments/hzwmqm/password_handling_practices/)
- url: https://www.reddit.com/r/golang/comments/hzwmqm/password_handling_practices/
---
Hey gophers,

 I am playing with go around half a year and now I feel my self confident enought in the language to write a complete almost large project in it. I wrote a few CLI tools and REST APIs but somehow I never get around password handling.

So the question is, what are the "best" way to handle user passwords in Go?  
Should I use bcrypt or should I use encryption-as-a-service like Vault, maybe some third party password less authentication?  
What is your oppinion about the listed options, is there anything thats better then the list above?

Thank you in advance!
