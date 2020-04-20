# golang
## [1][Code golfing in Go: Tips and Tricks](https://www.reddit.com/r/golang/comments/g4iwzl/code_golfing_in_go_tips_and_tricks/)
- url: https://www.reddit.com/r/golang/comments/g4iwzl/code_golfing_in_go_tips_and_tricks/
---
A few days ago there was an announcement that [code-golf.io](https://code-golf.io/) added Go support. For those of you unfamiliar with "code golf" the goal is to write a program to solve a given problem with the fewest characters possible. You can track you scores on the leaderboards and try to fight for the top spot.

Since the announcement I've been golfing away in Go and have found a few tricks I'd like to share. Hoping others can share their own tricks too!

Minimum program
===

Here is the minimum amount of code you need every single time. It's 39 characters.

    package main
    import."fmt"
    func main(){}

You need the `"fmt"` package because you will have to print something to stdout at some point, so you'll need to call Print/Printf/Println. As far as I can tell there is no shorter method for printing. Using a function in `io` package seems to always be longer. Calling the built-in `print` goes to stderr, but the requirements are to print to stdout.

Because you have to at least call `Print()` and that's another 7 characters, you're looking at a minimum of 46 characters before you even add any logic or data. This is a good number to keep in mind, so that when you're looking at scores you want to beat, you can subtract this number and it gives you an idea of how many characters you have to work with to actually solve the problem.


Assign functions that you use often to shorter variable names
===

    P:=Printf
    P("calling once")
    P("calling twice")


Never use "\n"
===

Instead of using `"\n"`, use backticks and a newline to save a character.

    // this
    Print(`
    `)
    
    // not this
    Print("\n")

Use assignment to multiple variables
===

    // 8 chars
    a,b:=0,1
    
    // 9 chars (because of required newline or semicolon)
    a:=0
    b:=1

Sometimes there are opportunities to do this in loops too.

    // 34 chars
    P:=Println
    for i:=0;i&lt;20;i++{P(i)}
    
    // 33 chars if we assign 'i' along with P
    P,i:=Println,0
    for i&lt;20{P(i);i++}


Import packages in the global namespace
===

You can import packages into the global namespace so you don't have to prefix calls with the package name.

    import(."fmt"
    ."strings")
    func main(){Println(ToUpper("xyz"))}

Notice you don't have to write `fmt.Println` or `strings.ToUpper`.


Compare many values with switch
===

    // 37 chars
    if x==3||x==6||x==8||x==9{Println(x)}
    
    // 33 chars
    switch x{case 3,6,8,9:Println(x)}


Use the unused index variable in a loop for something else
===

You often need to loop with `range` and only care about the elements of the slice or string but not the indexes. Imagine a scenario where you're looping over `Args` and you need to do a map lookup on each one to get an integer.

Normally you would do something like this:

    // 78 chars
    m:=map[string]int{'arg1':1,'arg2':2,'arg3':5}
    for _,a:=range Args[1:]{x:=m[a]}

But you can save a character by taking advantage of the fact that there's that free assignment operator in the loop you can use.

    // 77 chars
    m:=map[string]int{'arg1':1,'arg2':2,'arg3':5}
    for x,a:=range Args[1:]{x=m[a]}

Notice we can remove the colon from `x:=m[a]` since we initialized the variable in the for loop to an index value we don't care about and can just throw away immediately.

More Tips?
===

Anyone have some more tips specific to Go?
## [2][Thanos Metrics - Coding Style Guide](https://www.reddit.com/r/golang/comments/g4plmy/thanos_metrics_coding_style_guide/)
- url: https://thanos.io/contributing/coding-style-guide.md/
---

## [3][how to convert AST of toy language to Golang's AST?](https://www.reddit.com/r/golang/comments/g4pyve/how_to_convert_ast_of_toy_language_to_golangs_ast/)
- url: https://www.reddit.com/r/golang/comments/g4pyve/how_to_convert_ast_of_toy_language_to_golangs_ast/
---
so I am working through Writing an Interpreter in Go. 

I was wondering if it possible to parse some toy language and convert it into Go's code? 

And would it be possible at runtime (not at compile time)? like I would have a go webserver running, to which I will send my code and server would send me back the results.
## [4][week 4 from my learning go series is up! this week I cover Slices and Maps 👍🏻](https://www.reddit.com/r/golang/comments/g4hq56/week_4_from_my_learning_go_series_is_up_this_week/)
- url: https://www.martincartledge.io/learning-go-week-4/
---

## [5][qrcp: transfer files over wifi from your computer to your mobile device by scanning a QR code without leaving the terminal.](https://www.reddit.com/r/golang/comments/g45e7z/qrcp_transfer_files_over_wifi_from_your_computer/)
- url: https://github.com/claudiodangelis/qrcp
---

## [6][elk - 🦌 Minimalist yaml based task runner](https://www.reddit.com/r/golang/comments/g4lslg/elk_minimalist_yaml_based_task_runner/)
- url: https://github.com/jjzcru/elk
---

## [7][OpenCensus Prometheus's push gateway exporter](https://www.reddit.com/r/golang/comments/g4sk2t/opencensus_prometheuss_push_gateway_exporter/)
- url: https://github.com/hqt/opencensus-pushgateway-exporter
---

## [8][golang - docker client image operations](https://www.reddit.com/r/golang/comments/g4rybv/golang_docker_client_image_operations/)
- url: https://www.reddit.com/r/golang/comments/g4rybv/golang_docker_client_image_operations/
---
Even if the content is in Turkish language, the code is clear and understandable.

[Golang - Docker Client Image Operations](https://medium.com/trendyol-tech/golang-docker-client-ile-container-i%CC%87%C5%9Flemleri-6417884f4dbb)
## [9][Go CLI for real world container numbers - Good model/template for other Go CLI projects?](https://www.reddit.com/r/golang/comments/g4p6g6/go_cli_for_real_world_container_numbers_good/)
- url: https://www.reddit.com/r/golang/comments/g4p6g6/go_cli_for_real_world_container_numbers_good/
---
This CLI for real world container numbers was build to fulfill technical and UI expectations. What am I missing? Criticism is welcome!

[https://github.com/meyermarcel/icm](https://github.com/meyermarcel/icm)
## [10][Pxy – A Go server that proxies websocket livestreams to RTMP servers](https://www.reddit.com/r/golang/comments/g4r5gj/pxy_a_go_server_that_proxies_websocket/)
- url: https://github.com/chuabingquan/pxy
---

