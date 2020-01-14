# golang
## [1][rqlite, the lightweight distributed DB built on SQLite and written in Go, now scales read performance with read-only nodes](https://www.reddit.com/r/golang/comments/eoe99h/rqlite_the_lightweight_distributed_db_built_on/)
- url: https://github.com/rqlite/rqlite/releases/tag/v5.1.0
---

## [2][A tool/library for graph visualising](https://www.reddit.com/r/golang/comments/eok7sl/a_toollibrary_for_graph_visualising/)
- url: https://www.reddit.com/r/golang/comments/eok7sl/a_toollibrary_for_graph_visualising/
---
I'm looking for a tool/library that given a list of vertexes and edges with their labels, can visualise a graph.

I have found different draw libraries, but with them I need to create everything from scratch and hope there can be an automated way already implemented.

Some context is that I've found the [promviz](https://github.com/nghialv/promviz) library, that gets data from Prometheus and creates a graph where Prometheus labels are the vertexes and metrics values are edge labels. Unfortunately it's not updated since 2017. I'm trying to find something similar, or if not, to implement on my own.

Thanks
## [3][Test your CLIs easily with simple YAML definitions](https://www.reddit.com/r/golang/comments/eoknw8/test_your_clis_easily_with_simple_yaml_definitions/)
- url: https://github.com/SimonBaeumer/commander
---

## [4][Inline debugging of Lambda using AWS SAM CLI](https://www.reddit.com/r/golang/comments/eol46x/inline_debugging_of_lambda_using_aws_sam_cli/)
- url: https://www.reddit.com/r/golang/comments/eol46x/inline_debugging_of_lambda_using_aws_sam_cli/
---
(TL;DR at the bottom) I've been tinkering around with Go for a little bit now, and I have been playing with the idea to develop some serverless applications using golang. However, my experience in debugging with Go using VSCode have been not so great, at least when compared to NodeJS.

In NodeJS, I write some code that starts with `exports.handler`, use the following template as the `template.yaml`:

    AWSTemplateFormatVersion: '2010-09-09'
    Transform: 'AWS::Serverless-2016-10-31'
    Resources:
      MyFunction:
        Type: 'AWS::Serverless::Function'
        Properties:
          Handler: index.handler
          CodeUri: ./
          Runtime: nodejs8.10

Run `sam local invoke -d 5858 -e some-event.json`, attach the debugger to the process using the vscode `launch.json`

    {
      "version": "0.2.0",
      "configurations": [
        {
        "name": "Attach to SAM CLI",
        "type": "node",
        "request": "attach",
        "address": "localhost",
        "port": 5858,
        "localRoot": "",
        "remoteRoot": "/var/task",
        "protocol": "inspector",
        "stopOnEntry": false
        }
      ]
    }

and I'm ready to go.

In Go however, this is not all to easy to accomplish. I start off with a simple Go lambda application using `sam init --runtime go1.x`.

Running the lambda locally using sam was not a problem - I ran the commands as below.

    make build
    touch event.json # this I followed a standard API event structure
    sam local invoke HelloWorldFunction -e event.json
    
    {"statusCode":200,"headers":null,"multiValueHeaders":null,"body":"Hello, 2.49.15.103\n"}

With that, I refer to the documentation written here: [https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-using-debugging-golang.html](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-using-debugging-golang.html)

&amp;#x200B;

After setting up everything, including recompiling the delve debugger, and setting delve api to 2 (or not, it does not matter), I would always get this error message:

&amp;#x200B;

    sam local invoke HelloWorldFunction -e event.json -d 5986 --debugger-path ./tmp --debug-args "-delveAPI=2"
    Invoking hello-world (go1.x)
    2020-01-14 17:06:03 Found credentials in shared credentials file: ~/.aws/credentials
    
    Fetching lambci/lambda:go1.x Docker container image......
    Mounting /repos/go-lambda-sam/hello-world as /var/task:ro,delegated inside runtime container
    Could not create config directory: mkdir /home/sbx_user1051: permission denied.API server listening at: [::]:5986
    2020-01-14T13:06:07Z info layer=debugger launching process with args: [/var/task/hello-world]

and with that (I have also tried removing and reinstalling golang), I gave up doing inline debugging with go on AWS SAM CLI.   


TL;DR: I am wondering if anyone has other experiences dealing with this (inline debugging of go lambda using vscode and AWS SAM CLI) before, and I wonder why is setting up tooling for go more troublesome that it needs to be, or if I am looking at the wrong places.
## [5][Goroutines don't scale linearly but go programs do?](https://www.reddit.com/r/golang/comments/eol1i1/goroutines_dont_scale_linearly_but_go_programs_do/)
- url: https://www.reddit.com/r/golang/comments/eol1i1/goroutines_dont_scale_linearly_but_go_programs_do/
---
Hi,

I'm looking for some general advice on optimizing for performance.

I have a rather complex bit of code, but the code lives in a function (no globals), so I can either a) call this function in multiple goroutines:

`output := make(chan int, 4)`

`for w := 0; w &lt; 4; w++ {`  
 `go DoStuff(output)`  
`}`

`for a := 0; a &lt; 4; a++ {`  
 `fmt.Println(&lt;-output)`  
`}`

Or b) use the same kind of function, compile my go program and run the thing 4 times in parallel.

&amp;#x200B;

When I use 4 goroutines, it takes about 2.5x as long and running the compiled program in 4x in parallel! And 4 copies of the program running in parallel seem to run at the same speed as a single running version, i.e. near perfect scaling.

As far as I can tell (still a Go-newb though), none of the goroutines accesses the same data, and I made deep copies for all the structs they process.

I would love and tips on how to proceed. I want to scale to 32-cores, but after testing on up to 12-cores so far,  I get rapidly diminishing returns.
## [6][goban: a linter to ban "bad" functions](https://www.reddit.com/r/golang/comments/eoknxi/goban_a_linter_to_ban_bad_functions/)
- url: https://github.com/utrack/goban
---

## [7][GoatCounter: privacy-friendly web analytics written in Go](https://www.reddit.com/r/golang/comments/eog4mt/goatcounter_privacyfriendly_web_analytics_written/)
- url: https://github.com/zgoat/goatcounter
---

## [8][Continuous Benchmark for Go Project](https://www.reddit.com/r/golang/comments/eo5u5h/continuous_benchmark_for_go_project/)
- url: https://github.com/knqyf263/cob
---

## [9][cards - Playing card library (updated since receiving subreddit feedback)](https://www.reddit.com/r/golang/comments/eokum4/cards_playing_card_library_updated_since/)
- url: https://git.sr.ht/~tslocum/cards?0.1.1
---

## [10][Hey gophers what are you working on in 2020?](https://www.reddit.com/r/golang/comments/eof2nd/hey_gophers_what_are_you_working_on_in_2020/)
- url: https://www.reddit.com/r/golang/comments/eof2nd/hey_gophers_what_are_you_working_on_in_2020/
---
EDIT : 
I am working on reddit api wrapper.
