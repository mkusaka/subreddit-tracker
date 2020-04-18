# golang
## [1][I wrote a terminal UI for npm in Go!](https://www.reddit.com/r/golang/comments/g3fkci/i_wrote_a_terminal_ui_for_npm_in_go/)
- url: https://github.com/jesseduffield/lazynpm
---

## [2][Why slice hasn't default '==' operator?](https://www.reddit.com/r/golang/comments/g3mmp0/why_slice_hasnt_default_operator/)
- url: https://www.reddit.com/r/golang/comments/g3mmp0/why_slice_hasnt_default_operator/
---
Is there any design decision for this?
## [3][Go reflections deep dive— from structs and interfaces](https://www.reddit.com/r/golang/comments/g3lvg0/go_reflections_deep_dive_from_structs_and/)
- url: https://medium.com/@snird/go-reflections-deep-dive-from-structs-and-interfaces-e1931f0c99af?sk=0f430b8906e583630bdcd4ce9a452943
---

## [4][tinysms Simple SMS sending library for small projects or prototyping using SMTP](https://www.reddit.com/r/golang/comments/g3dn7l/tinysms_simple_sms_sending_library_for_small/)
- url: https://www.reddit.com/r/golang/comments/g3dn7l/tinysms_simple_sms_sending_library_for_small/
---
I wasn't too happy with the current SMS libraries out there. They largely lacked GoLang standards (or they just didn't work at all) So I wrote one in a few hours for a trading bot I'm working on to let me know if any eventful happened while away and I didn't want to pay for SMS service or worry about getting cut off for using a demo API token. You just need an SMTP server (which most of you have available for free by your email provider) and you are all set!

  
[https://github.com/ajm113/go-tinysms](https://github.com/ajm113/go-tinysms)  


 Let me know if there is anything I can improve on or if you just want to let me know it works out for your project. I'm always open to PR requests!
## [5][I have created a Golang tool to scale persistent volume in Kubernetes dynamically](https://www.reddit.com/r/golang/comments/g3lt9g/i_have_created_a_golang_tool_to_scale_persistent/)
- url: https://www.reddit.com/r/golang/comments/g3lt9g/i_have_created_a_golang_tool_to_scale_persistent/
---
[https://github.com/opstree/dynamic-pv-scaler](https://github.com/opstree/dynamic-pv-scaler)
## [6][tdewolff/minify - Go minifiers for web formats](https://www.reddit.com/r/golang/comments/g3ljn8/tdewolffminify_go_minifiers_for_web_formats/)
- url: https://github.com/tdewolff/minify
---

## [7][Build a custom blockchain in Go from scratch](https://www.reddit.com/r/golang/comments/g3m2uk/build_a_custom_blockchain_in_go_from_scratch/)
- url: https://github.com/web3coach/the-blockchain-way-of-programming-newsletter-edition
---

## [8][Rate limit http middleware with few algorithms (Sliding Window, Leaky Bucket)](https://www.reddit.com/r/golang/comments/g308cp/rate_limit_http_middleware_with_few_algorithms/)
- url: https://www.reddit.com/r/golang/comments/g308cp/rate_limit_http_middleware_with_few_algorithms/
---
Hi, if you need rate limit [library](https://github.com/Shareed2k/go_limiter) or middleware for rate limit with algorithms like (Sliding Window, Leaky Bucket) , redis as store, you welcome to try test and give me a feedback... thanks ;)

[fiber framework middleware](https://github.com/Shareed2k/fiber_limiter)

[echo framework middleware](https://github.com/Shareed2k/echo_limiter)

[http middleware](https://github.com/Shareed2k/http_limiter)
## [9][Best practice for calling multiple API’s](https://www.reddit.com/r/golang/comments/g3jwyo/best_practice_for_calling_multiple_apis/)
- url: https://www.reddit.com/r/golang/comments/g3jwyo/best_practice_for_calling_multiple_apis/
---
Hi, I am new at golang, I worked with other languages such as Java, Ruby or rails, now I am doing a project where I need to call different api's but all serve the same thing. but I can't figure out what is the best practice to do it

what I have basically is API's for pharmacies all of them serve 3 things

get stock for medicine
order certain medicine
list all available medicines
you should know, some of them is XML some is SOAP and others JSON, each of them different structure .. etc

but all of them I need to retain unified api structure for the 3 methods.

What I would do in different language is to build it using repository design pattern

any tips could be helpful Thanks
## [10][Start an ssh session with Golang!](https://www.reddit.com/r/golang/comments/g3jcsm/start_an_ssh_session_with_golang/)
- url: https://www.reddit.com/r/golang/comments/g3jcsm/start_an_ssh_session_with_golang/
---
Hi I have a bunch of servers and a bunch of keys to ssh into the server.

 I have implemented a basic functionality using the **promptui** library where I can select a server from the list.

But I am unable to figure out how to ssh to that server. Here is what I have tried:

Using os/exec to run the command: **ssh -i &lt;path to key&gt; user@ip.**

But the program just exits with some code. Can anyone guide me on this?
