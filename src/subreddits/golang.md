# golang
## [1][Ergo - a new framework for creating mesh networks with Erlang technologies. New release 1.0.0 https://github.com/halturin/ergo 🚀Details in comment](https://www.reddit.com/r/golang/comments/fcr3ez/ergo_a_new_framework_for_creating_mesh_networks/)
- url: https://i.redd.it/yddtl1f1tek41.gif
---

## [2][The Go Blog: A new Go API for Protocol Buffers](https://www.reddit.com/r/golang/comments/fciob7/the_go_blog_a_new_go_api_for_protocol_buffers/)
- url: https://blog.golang.org/a-new-go-api-for-protocol-buffers
---

## [3][RESTful API in GO](https://www.reddit.com/r/golang/comments/fctlkq/restful_api_in_go/)
- url: https://www.reddit.com/r/golang/comments/fctlkq/restful_api_in_go/
---
Hey reader,   
I want to learn how to write RESTful APIs in GOLang.
## [4][The httpgovernor package provides an HTTP request concurreny limiter](https://www.reddit.com/r/golang/comments/fcqv0k/the_httpgovernor_package_provides_an_http_request/)
- url: https://github.com/juju/httpgovernor
---

## [5][Golang is my new shell](https://www.reddit.com/r/golang/comments/fcumfu/golang_is_my_new_shell/)
- url: https://www.reddit.com/r/golang/comments/fcumfu/golang_is_my_new_shell/
---
This is not something informative or sophisticated, I just wanted to share my experience with Go over past couple of years.

I am a software engineer, that breed that sometimes is called "data engineer", in particular on AWS stack. Along with writing some basic "get data out of CSV and put that into parquet" stuff on EMR clusters ( got a chance to launch a cluster of X1 EC2s, that feels really weird ) - I also do many various automations/monitoring/etc with either Ansible or good old shell scripts.

For the past couple year I wrote close to zero bash scripts. 

I have a ton of a small utilities like "calculate the average data volume in this S3 folder", or "get the parquet schema from these folders and compare it with the reference", or "get into the history of a EMR cluster and visualize the run time trend". All of them are written in Go.

Considering that I have some Scala / functional programming background ( been doing that for almost 10 years by now) - it would be naturally to expect me suffering from lack of generics. And in fact I was missing them from time to time. But not anymore. I started to realize that most of the times I don't really need generics when I write some utility code. And even for more complex services that involve complex business logic and communications with other services I never really found myself struggling over something.

Even that `if err := ...; err != nil { ... }` makes much more sense to me over `Try` or some fancy Cats effects.

The tooling is great, and being able to compile things to run on MacOS or EC2 with Linux, or even Raspberry Pi is so amazing ( tried that on Rust, failed miserably ) that I take it for granted now.

So basically I just wanted to say thanks to this community for your existence, I learned a lot and keep learning. And even if my work is circling around data processing and data engineering, write Spark applications and consume data streams from Kafka with Akka Streams or FS2 - there are still opportunities to write some things in Go.
## [6][efficient json.MarshalJSON interface usage for streaming?](https://www.reddit.com/r/golang/comments/fcuf7i/efficient_jsonmarshaljson_interface_usage_for/)
- url: https://www.reddit.com/r/golang/comments/fcuf7i/efficient_jsonmarshaljson_interface_usage_for/
---
We encountered the problem that we need to connect a new application which transfers files as base64 encoded string in a JSON, like following:

    {
        "attachments": [
            {
                "name": "asdasdasd",
                "value": "base64encodedSomethingValue"
            },
            ....
        ]
    }

So i tried to upload a 200MB file, which caused the memory to dramatically increase.

Because the currently solution i found is to use the [https://golang.org/pkg/encoding/json/#RawMessage.MarshalJSON](https://golang.org/pkg/encoding/json/#RawMessage.MarshalJSON) on the "attachement" item struct but, for that i need to allocate the whole base64 into a \`\[\]byte\`, before it will be written to the json encoder, which consumes a lot of memory for nothing.

So my question how, can i stream the output of `base64.NewEncoder` to the json encoder? 

I checked this [https://github.com/mailru/easyjson](https://github.com/mailru/easyjson) but it also uses a `byte.Buffer` internally.
## [7][More or less what versions of Go support what OpenBSD releases (as of March 2020)](https://www.reddit.com/r/golang/comments/fcu4uo/more_or_less_what_versions_of_go_support_what/)
- url: https://utcc.utoronto.ca/~cks/space/blog/programming/GoWhatOpenBSDs-2020-03
---

## [8][Low level SSD read write in golang](https://www.reddit.com/r/golang/comments/fct56d/low_level_ssd_read_write_in_golang/)
- url: https://www.reddit.com/r/golang/comments/fct56d/low_level_ssd_read_write_in_golang/
---
Hi Guys,
I am new to golang and I wanted to know if there is a way to access low level system operations in golang.
What I need specifically:
1.  How can I access free Blocks and pages?
2. Write-read-earse blocks

Thanks a lot.
## [9][Simple CLI to schedule tweets](https://www.reddit.com/r/golang/comments/fcsw4s/simple_cli_to_schedule_tweets/)
- url: https://www.reddit.com/r/golang/comments/fcsw4s/simple_cli_to_schedule_tweets/
---
Hello everybody!

I developed recently a CLI which send tweets via a CSV, at a precise date. I needed a quick way to be able to schedule tweets.

I think about extending the functionality to other social media, if somebody is interested.

Any feedback is welcome!

https://github.com/Phantas0s/ottosocial
## [10][GoLand help](https://www.reddit.com/r/golang/comments/fcsm7k/goland_help/)
- url: https://www.reddit.com/r/golang/comments/fcsm7k/goland_help/
---
I'm done with trying to make VSCode+Gopls work as it's been crashing multiple times a day, randomly stops autocompleting, gets stuck trying to save files and so forth. I've two minor issues with GoLand and am wondering if anyone has any ideas.

1. When saving files, goimports/gofmt takes way, way longer than VSCode which is near instant. It's almost as if it's formatting across multiple files, and not just the current one. Any way around that?
2. Is there are way to setup a shortcut to test the current package? E.g. so that if I am looking at a file, it'll run all tests in the current directory? I have tons of packages and love the "Go: Test Package" feature in the VSCode integration.
