# golang
## [1][Look what my colleague gave me. It's so cute! 🥰](https://www.reddit.com/r/golang/comments/fdab62/look_what_my_colleague_gave_me_its_so_cute/)
- url: https://i.redd.it/w44c9ga5emk41.jpg
---

## [2][Building a terminal dashboard in Golang in 300 lines of code](https://www.reddit.com/r/golang/comments/fd722v/building_a_terminal_dashboard_in_golang_in_300/)
- url: https://levelup.gitconnected.com/building-a-terminal-dashboard-in-golang-in-300-lines-of-code-3b9f83f363a8
---

## [3][Golang is my new shell](https://www.reddit.com/r/golang/comments/fcumfu/golang_is_my_new_shell/)
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
## [4][Go is one of the most loved languages based on StackOverflow's developer survey results](https://www.reddit.com/r/golang/comments/fcyao2/go_is_one_of_the_most_loved_languages_based_on/)
- url: https://learnworthy.net/stackoverflows-developer-survey-results-for-2019/
---

## [5][gopls 0.3.3 update release notes](https://www.reddit.com/r/golang/comments/fcz14k/gopls_033_update_release_notes/)
- url: https://github.com/golang/go/issues/33030#issuecomment-593592737
---

## [6][Newbie needs some advice](https://www.reddit.com/r/golang/comments/fdcohn/newbie_needs_some_advice/)
- url: https://www.reddit.com/r/golang/comments/fdcohn/newbie_needs_some_advice/
---
Hey guys. 

I am currently doing a bit of freelance web development (front end only). Prior to this I practiced law for 13 years and decided to change careers at the ripe age of 38 at the end of 2017. I guess I was lucky to get a project fairly early on, but it is really simple and basically amounts to data entry with a bit of HTML and CSS throw in. Furthermore, the pay is not that good and I am slowly reaching the stage where I have to find a job and I'm getting kind of desperate. The good thing is that it "keeps the pot on the boil" while I learn further. 

I am fairly proficient at Javascript and Python, but I feel as if I'm just not at that stage where I can apply for jobs. Granted, I think I have been in tutorial purgatory for the past 2 years, having that fear of trying to build something on my own. 

Ironically, the first language I was introduced to when I first started out was Go. I have played around with it a bit before and I really enjoy it. I do, however, want to reach that stage where I am employable. I have both Todd Mcleod's courses (the Learn Go one and the Web Development with Go). Are there any other resources that are worth it? I also want to find out where one can find a mentor to, firstly, mentor me as I learn and, secondly, keep me accountable.  

Any help will be appreciated.

Thanks
## [7][How to find goroutines during debugging - aka goroutine labeling](https://www.reddit.com/r/golang/comments/fcvruj/how_to_find_goroutines_during_debugging_aka/)
- url: https://blog.jetbrains.com/go/2020/03/03/how-to-find-goroutines-during-debugging/
---

## [8][Verify Docker image exists using SDK?](https://www.reddit.com/r/golang/comments/fd9s98/verify_docker_image_exists_using_sdk/)
- url: https://www.reddit.com/r/golang/comments/fd9s98/verify_docker_image_exists_using_sdk/
---
Is there anyway to check if an image exists/is available using the Go SDK?

I know there is the ImageSearch functionality, but it seems to only be able to search by the image name, not by the name with repository and tag.

The following works.


	results, _ := cli.ImageSearch(context.Background(), "golangci-lint", types.ImageSearchOptions{Limit: 1})


But I'd like to be able to do search like this


	results, _ := cli.ImageSearch(context.Background(), "golangci/golangci-lint:v1.23.7", types.ImageSearchOptions{Limit: 1})

EDIT:

Here's a full script.

Only dependency is `github.com/docker/docker v1.13.1`

    package main
    
    import (
    	"context"
    	"github.com/docker/docker/api/types"
    	"github.com/docker/docker/client"
    	"os"
    )
    
    func main() {
    	cli, err := client.NewClientWithOpts(client.FromEnv)
    	if err != nil {
    		os.Exit(1)
    	}
    
    	results, _ := cli.ImageSearch(context.Background(), "golangci-lint", types.ImageSearchOptions{Limit: 1})
        // Len = 1
    	println(len(results))
    
    	results, _ = cli.ImageSearch(context.Background(), "golangci/golang-lint:v1.23.7", types.ImageSearchOptions{Limit: 1})
        // Len = 0
   	println(len(results))
    }

And to verify image really does exist - the following CLI command works.

    docker pull golangci/golangci-lint:v1.23.7
## [9][Ergo - a new framework for creating mesh networks with Erlang technologies. New release 1.0.0 https://github.com/halturin/ergo 🚀Details in comment](https://www.reddit.com/r/golang/comments/fcr3ez/ergo_a_new_framework_for_creating_mesh_networks/)
- url: https://i.redd.it/yddtl1f1tek41.gif
---

## [10][Jet is a tool written in Go to convert source code into Docker images. First release 0.1.0.](https://www.reddit.com/r/golang/comments/fcv8ba/jet_is_a_tool_written_in_go_to_convert_source/)
- url: https://github.com/lade-io/jet
---

