# golang
## [1][Proposal: Register-based Go calling convention](https://www.reddit.com/r/golang/comments/i8x4xe/proposal_registerbased_go_calling_convention/)
- url: https://go.googlesource.com/proposal/+/refs/changes/78/248178/1/design/40724-register-calling.md
---

## [2][Static type checking for the empty interface{}](https://www.reddit.com/r/golang/comments/i8w8c2/static_type_checking_for_the_empty_interface/)
- url: https://github.com/siadat/interface-type-check
---

## [3][A weechat IRC client encryption system attempt](https://www.reddit.com/r/golang/comments/i8wp6y/a_weechat_irc_client_encryption_system_attempt/)
- url: https://git.sr.ht/~eau/wic
---

## [4][Recently started learning Go, and can't help wonder any good mature Packages or Frameworks for GUI development using Go, particularly for Windows? How about Linux?](https://www.reddit.com/r/golang/comments/i8gjol/recently_started_learning_go_and_cant_help_wonder/)
- url: https://www.reddit.com/r/golang/comments/i8gjol/recently_started_learning_go_and_cant_help_wonder/
---
Thank you for taking the time.
## [5][Given a multipart upload of an array of video clips, how can I merge them into one video?](https://www.reddit.com/r/golang/comments/i8wnck/given_a_multipart_upload_of_an_array_of_video/)
- url: https://www.reddit.com/r/golang/comments/i8wnck/given_a_multipart_upload_of_an_array_of_video/
---

## [6][Introduce a tool to do code coverage collecting for the API or e2e tests more easily](https://www.reddit.com/r/golang/comments/i8v5ad/introduce_a_tool_to_do_code_coverage_collecting/)
- url: https://www.reddit.com/r/golang/comments/i8v5ad/introduce_a_tool_to_do_code_coverage_collecting/
---
Open Source: [https://github.com/qiniu/goc](https://github.com/qiniu/goc)

`Goc` uses `go tool cover` to add statements counters into the source code, and what's more, it also injects HTTP APIs to aggregate all the counters, which makes it easily to pull coverage data from the applications under test at runtime. 

Enjoy and feel free to leave comments if you have any questions.
## [7][Should I use net/http when deploying?](https://www.reddit.com/r/golang/comments/i8x1qc/should_i_use_nethttp_when_deploying/)
- url: https://www.reddit.com/r/golang/comments/i8x1qc/should_i_use_nethttp_when_deploying/
---
Newbie question here. So far I did majority of my "web" development using django. I'm used to things like WSGI. I'm wondering whether I should rely on application servers like gunicorn or uwsgi and point my webserver (nginx instance) to them as well when writing apps with golang? Alternatively, can I just use net/http directly and use nginx as a reverse proxy and a load balancer instead? Is the latter a production ready solution and something that is commonly used?
## [8][Go and JSON encoding/decoding - Tit Petric](https://www.reddit.com/r/golang/comments/i8z0sy/go_and_json_encodingdecoding_tit_petric/)
- url: https://scene-si.org/2020/08/13/go-and-json-encoding-decoding/
---

## [9][How to run a command on a remote container and get STDOUT?](https://www.reddit.com/r/golang/comments/i8yim7/how_to_run_a_command_on_a_remote_container_and/)
- url: https://www.reddit.com/r/golang/comments/i8yim7/how_to_run_a_command_on_a_remote_container_and/
---
Hello. I'm trying to run a command on a remote  container and get the output stream on my local machine, in a secure way.

My first thought was to simply run ssh + the command. That would be secure, but I'm not sure how I would authorize the user who's running the remote command such that they can only execute this command and not do anything else on the running container.

So my next thought was, could I set up an RPC server and somehow stream the Stdout back to the client for the duration of the running program that was called? Of course the server would authorise the client through OAuth or something similar. 

I'm newer to Golang and mostly dealing with web applications day to day, so any help here would be appreciated. Thank you!
## [10][Clean API example with model injection and access policies](https://www.reddit.com/r/golang/comments/i8vpvl/clean_api_example_with_model_injection_and_access/)
- url: https://gist.github.com/philippta/a9fac29dccecd88fbfca643b5e5b7b68
---

