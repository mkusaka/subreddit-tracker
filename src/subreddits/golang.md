# golang
## [1][Olric v0.2.0 is out: Distributed cache and in-memory key/value data store. It can be used both as an embedded Go library and as a language-independent service.](https://www.reddit.com/r/golang/comments/gd7v1s/olric_v020_is_out_distributed_cache_and_inmemory/)
- url: https://github.com/buraksezer/olric#olric-
---

## [2][I try to create a minimalistic and flexible Go repository template. Feedback is more than welcome.](https://www.reddit.com/r/golang/comments/gdajt5/i_try_to_create_a_minimalistic_and_flexible_go/)
- url: https://github.com/golang-templates/seed
---

## [3][Wrote a Chip-8 emulator in Go. It is a WIP and has no tests written yet, but so far everything seems to work except for some audio bugs.](https://www.reddit.com/r/golang/comments/gcvy16/wrote_a_chip8_emulator_in_go_it_is_a_wip_and_has/)
- url: https://github.com/bradford-hamilton/chippy
---

## [4][Building a REST API With Go Gin - Part 2: Login/Register/Authentication](https://www.reddit.com/r/golang/comments/gd04vj/building_a_rest_api_with_go_gin_part_2/)
- url: https://www.youtube.com/watch?v=XxCr4QOD5Hs
---

## [5][Convert time.Minute to seconds(int)](https://www.reddit.com/r/golang/comments/gd8go8/convert_timeminute_to_secondsint/)
- url: https://www.reddit.com/r/golang/comments/gd8go8/convert_timeminute_to_secondsint/
---
I have a constant of type time.Duration: **Time\_Min = 10 \* time.Minute** (this holds value 10m). I want to convert this value to seconds (10\*60 = 600s) and store it as integer(int and not int64). 

for eg, Time\_Sec = Time\_Min \* 60 

OR Time\_Sec = Time\_Min \* time.Second

Where, Time\_sec should be of type int.

I am unable to find a way to perform type conversion for time.Duration. Any help will be highly appreciated.

Thanks
## [6][Supporting gRPC and HTTP Side by Side](https://www.reddit.com/r/golang/comments/gd0s9n/supporting_grpc_and_http_side_by_side/)
- url: https://www.reddit.com/r/golang/comments/gd0s9n/supporting_grpc_and_http_side_by_side/
---
Our organization is currently evaluating moving to gRPC for our microservices, and eventually moving to gRPC exclusively with the exception of services that are consumed from a web browser. We will have services that will need to continue support HTTP / JSON long term, but even in the interim we are being asked to support gRPC and HTTP side by side on all services that adopt gRPC. Most of our code bases are using GoKit with the HTTP transport. Adding gRPC shouldn't be a heavy lift but in many of the examples I've seen for GoKit gRPC services there are two sets of models being maintained, one for protobuf that is auto generated from protoc, and one for the application which the GoKit endpoints and services receive and return. This seems to lead to duplication as the gRPC transport ends up mapping the protobuf model to another set of very similar models. While being trivial this adds duplicative effort and begs the question what is the canonical source of truth for the models, the proto file, or the models the service expects.

What are the best practices around supporting gRPC and HTTP, particularly with GoKit? Should we maintain two sets of models and just do the mapping? Or can we just use the protobuf generated models through the endpoint and services?
## [7][Command line tool that runs a local local web server for facilitating bash script execution](https://www.reddit.com/r/golang/comments/gdbnlz/command_line_tool_that_runs_a_local_local_web/)
- url: https://github.com/atye/go-swizard
---

## [8][Handling MongoDB Errors](https://www.reddit.com/r/golang/comments/gdblea/handling_mongodb_errors/)
- url: https://www.reddit.com/r/golang/comments/gdblea/handling_mongodb_errors/
---
How can I differentiate timeout error from not found error when I am calling collection.FindOne()? Thank you.
## [9][I'm a noob, here a Eshop Cli I made, any suggestions?](https://www.reddit.com/r/golang/comments/gdavow/im_a_noob_here_a_eshop_cli_i_made_any_suggestions/)
- url: https://www.reddit.com/r/golang/comments/gdavow/im_a_noob_here_a_eshop_cli_i_made_any_suggestions/
---
Hello there,
here's a simple CLI I made to put together a couple of things I learned:

https://github.com/marcodenisi/eshop-tracker

It's a CLI retrieving and showing Italian Nintendo Eshop games with related prices. I've used boltdb to manage a key-value store and promptui to enrich a little bit the terminal UI. The idea was to write something that could be easily extended to be consumed via REST API as well.

I'd really like to get some advice on how I can improve my style. In fact, I can't really understand if there's some kind of naming conventions for files and folders and my directory structure seems a little odd.

I come from Java and I think my project somehow shows it, it's all a little bit to jav-ish, I'd really like to write more idiomatic Go code.

Thanks to all!
## [10][How does a Go WebAPI look like today?](https://www.reddit.com/r/golang/comments/gdarfm/how_does_a_go_webapi_look_like_today/)
- url: https://www.reddit.com/r/golang/comments/gdarfm/how_does_a_go_webapi_look_like_today/
---
Hi. I started to learn Go and got most of the basics now. Now it would be awesome to know how Go is used out in the big wide world. How does your workplace Go look like? What packages are you using to built your WebAPI's? Would love to hear what problems companies tackle with Go and how they do it.

Is there any "realistic" opensource projects out there to learn from? How does the structure look like? Hope anyone has the time to share something :) Don't be afraid to get down and dirty with details.
