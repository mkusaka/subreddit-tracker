# golang
## [1][Go is a Pretty Average Language](https://www.reddit.com/r/golang/comments/gbez73/go_is_a_pretty_average_language/)
- url: https://blog.chewxy.com/2019/02/20/go-is-average/
---

## [2][Rob Pike interview for Evrone: “Go has become the language of cloud infrastructure”](https://www.reddit.com/r/golang/comments/gay4v3/rob_pike_interview_for_evrone_go_has_become_the/)
- url: https://evrone.com/rob-pike-interview
---

## [3][Go: Asynchronous Preemption](https://www.reddit.com/r/golang/comments/gbhnk1/go_asynchronous_preemption/)
- url: https://medium.com/a-journey-with-go/go-asynchronous-preemption-b5194227371c
---

## [4][Is there any Go based NFS Server mock](https://www.reddit.com/r/golang/comments/gbhmgk/is_there_any_go_based_nfs_server_mock/)
- url: https://www.reddit.com/r/golang/comments/gbhmgk/is_there_any_go_based_nfs_server_mock/
---
I want NFS Server Mock for my Unit tests
## [5][GoBDD hits v1.0](https://www.reddit.com/r/golang/comments/gbeg7x/gobdd_hits_v10/)
- url: https://www.reddit.com/r/golang/comments/gbeg7x/gobdd_hits_v10/
---
[GoBDD](https://github.com/go-bdd/gobdd) is a small framework focused on writing BDD tests using Gherkin syntax. I already [wrote about it before](https://www.reddit.com/r/golang/comments/cgv3km/gobdd_new_bdd_framework_for_go/). This time, I want to tell that the project has his first stable version. I simplified a lot of things.

The project is an alternative to [godog](https://github.com/cucumber/godog) but has a few differences. The most important thing is that godog is a separate binary that compiles our code under the hood. It has several disadvantages:

* no debugging (breakpoints) in the test. Sometimes it’s useful to go through the whole execution step by step
* metrics don’t count the test run this way
* some style checkers recognise tests as dead code
* it’s impossible to use built-in features like build constraints.
* no context in steps - so the state have to be stored somewhere else - in my opinion, it makes the maintenance harder

My motivation is to create a similar project but without magic whenever it's possible. I want to ask you to give it a try and suggest improvements/bug reports etc. I want to make GoBDD the best BDD testing tool in Go :) it won't happen without your help.
## [6][GoPing : Ping utility written in Go](https://www.reddit.com/r/golang/comments/gauvgj/goping_ping_utility_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/gauvgj/goping_ping_utility_written_in_go/
---
I am new to Go and also a sophomore undergrad. This is a traditional Ping utility written by me in Go using ICMP requests from scratch for both ipv4 and ipv6. [Github  repo link](https://github.com/shivamanipatil/GoPing)  . It is a small program and supports some of the functionalities of traditional Ping. It would be helpful if you could provide suggestions or feedback.
## [7][Using GitLab's CI for Go Packages : Stone Code](https://www.reddit.com/r/golang/comments/gb9wo6/using_gitlabs_ci_for_go_packages_stone_code/)
- url: http://stonecode.ca/posts/gitlab_ci_for_go/
---

## [8][flower: workflow engine that manages workflows composed of DAGs](https://www.reddit.com/r/golang/comments/gb7oaw/flower_workflow_engine_that_manages_workflows/)
- url: https://github.com/d-tsuji/flower
---

## [9][Learning Go](https://www.reddit.com/r/golang/comments/gbcsvx/learning_go/)
- url: https://www.reddit.com/r/golang/comments/gbcsvx/learning_go/
---
 I tried to jot down a list of concepts and code snippets that would help in learning Golang and applying in Web Development. Hope you find something useful. If you like it please give it a star  
[https://github.com/debck/Learning-Go](https://github.com/debck/Learning-Go)
## [10][If you use VSCode as your editor - and want free vulnerability scanning - we just put this out there!](https://www.reddit.com/r/golang/comments/gb2i5u/if_you_use_vscode_as_your_editor_and_want_free/)
- url: https://www.reddit.com/r/golang/comments/gb2i5u/if_you_use_vscode_as_your_editor_and_want_free/
---
[https://jfrog.com/blog/free-go-module-vulnerability-scanning-in-visual-studio-code/](https://jfrog.com/blog/free-go-module-vulnerability-scanning-in-visual-studio-code/)
