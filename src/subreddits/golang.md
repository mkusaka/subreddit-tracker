# golang
## [1][Is there a good GraphQL client for go?](https://www.reddit.com/r/golang/comments/gez4h1/is_there_a_good_graphql_client_for_go/)
- url: https://www.reddit.com/r/golang/comments/gez4h1/is_there_a_good_graphql_client_for_go/
---
I looked around but couldn't find anything useful. Do we not consume GraphQL apis from go? Or do we simply use raw HTTP requests?
## [2][smex - a cli app to extract and convert sitemaps](https://www.reddit.com/r/golang/comments/gf43cd/smex_a_cli_app_to_extract_and_convert_sitemaps/)
- url: https://github.com/hbish/smex
---

## [3][Are there are Go bindings to libosrm or any high performance routing alternatives?](https://www.reddit.com/r/golang/comments/gf45ae/are_there_are_go_bindings_to_libosrm_or_any_high/)
- url: https://www.reddit.com/r/golang/comments/gf45ae/are_there_are_go_bindings_to_libosrm_or_any_high/
---
I need to get road direction data into my go project. I'm using http requests to OSRM but it is a bottleneck and wanted to bring it into the application as a library. As OSRM is written in C++14 it seems challenging to integrate into Go with SWIG for example.

Need to make a large number of requests as part of a pre processing pass for a web service. The faster the better.

On a related note if I was to tackle the bindings is SWIG the right approach for C++ libraries in Go?
## [4][Rob Pike interview: “Go has indeed become the language of cloud infrastructure“](https://www.reddit.com/r/golang/comments/gegxe4/rob_pike_interview_go_has_indeed_become_the/)
- url: https://evrone.com/rob-pike-interview
---

## [5][Learn go with tests: HTTP handlers revisited](https://www.reddit.com/r/golang/comments/gf6d3u/learn_go_with_tests_http_handlers_revisited/)
- url: https://quii.gitbook.io/learn-go-with-tests/questions-and-answers/http-handlers-revisited
---

## [6][ent: a graph-based ORM for Go, from Facebook](https://www.reddit.com/r/golang/comments/gem8el/ent_a_graphbased_orm_for_go_from_facebook/)
- url: https://entgo.io
---

## [7][Packing args into a type for passing data into functions easily?](https://www.reddit.com/r/golang/comments/gf4vzz/packing_args_into_a_type_for_passing_data_into/)
- url: https://www.reddit.com/r/golang/comments/gf4vzz/packing_args_into_a_type_for_passing_data_into/
---
Hi guys.

I've been working professionally for about 4 months (as a backend developer working primarily on Rails), and I've been looking into Golang lately.

I thought I'd try building a simple LRU Cache as a pet project. In the process, I realised how the flexibility Ruby offers you can be extremely helpful. The ability to pass named arguments into a function, and then retrieve them easily, or choose a default value if the value is absent is pretty cool. For instance: -

&amp;#x200B;

    def initialize(name, **args): 

args is basically a Hash (hash map). All named arguments after the first argument to initialize() will be a part of the Hash.

Assign 'name' to instance variable '@name'

        @name = name 

If 'height' is present in args, assign it to '@height'. Else args\[:height\]  will evaluate to nil, and DEFAULT\_HEIGHT is assigned to '@height' instead.

        @height = args[:height] || DEFAULT_HEIGHT

&amp;#x200B;

I thought I'd build a small package that allows me to at least come close to this and this resulted in: [https://github.com/dannypz97/arghpacker](https://github.com/dannypz97/arghpacker). I can create a variable of type *ArgPacker* (which is basically a map) and can then add/modify values, or retrieve values  by key (or a default value if the key is not present in ArgPacker).

The code itself is very small - about 30 lines.

I plan to use this in my projects. I want to know if it's a bad idea since I don't want any weird ideas (or too many weird ideas) in my projects. I understand that this might have some overhead of its own, but I think this might have some benefits as well.

I want to know what you guys think. Constructive criticism is appreciated.
## [8][GOTTY error opening terminal: unknow](https://www.reddit.com/r/golang/comments/gf4onq/gotty_error_opening_terminal_unknow/)
- url: https://www.reddit.com/r/golang/comments/gf4onq/gotty_error_opening_terminal_unknow/
---
Hello everyone!

I have this problem and I am not able to solve it and I want to discuss it with you in case someone can give me another point of view.

I usually have a remote terminal open against one of the servers where I run the dnstop command to get query statistics to DNS.

So far everything is clear, but as my head never stops I began to think that instead of having a terminal it would be practical to be able to consult or run the dnstop but via the website. And well ... I got down to work, I searched ..... I documented ... and in the end I found what for me was the most suitable solution, which was to use GoTTY.

The idea itself was good, from the console it executed `gotty dnstop -4 -6 -Q -R l3 eth0`and automatically had on the web page [`http://xxx.xxx.xxx.xxx:8080`](http://xxx.xxx.xxx.xxx:8080/) the command dnstop showing me statistics.

So far fantastic, but my desire for perfectionism led me to automate it, in reality a not complex process and I solved it as follows by creating this script:

`#! / bin / bash`

`set -e`

`set -x`

`TERM = linux`

`echo "Executing GoTTY for DNStop ..."`

`/ usr / local / go_projects / bin / gotty --term xterm dnstop -4 -6 -Q -R -l 3 eth0`

then with crontab -e create the cron:

`@ reboot (sleep 30 ; /usr/local/execute_gotty.sh)`

If I manually launch the script it works, but if I let it auto-run after a reboot, when accessing via web, this is the result.

**error opening terminal: unknow**

Does anyone come up with a solution?

Thanks for the help
## [9][Could golang be a good alternative to rails in terms of occupying 80%+ of a backend of a livestreaming webapp](https://www.reddit.com/r/golang/comments/gevbmd/could_golang_be_a_good_alternative_to_rails_in/)
- url: https://www.reddit.com/r/golang/comments/gevbmd/could_golang_be_a_good_alternative_to_rails_in/
---

## [10][[question] How to properly manage common constants for a project in Go?](https://www.reddit.com/r/golang/comments/geyqwl/question_how_to_properly_manage_common_constants/)
- url: https://www.reddit.com/r/golang/comments/geyqwl/question_how_to_properly_manage_common_constants/
---
So a software project has strings/integers/other type constants that need to be referenced all over the place. Whats the best practice for handling this in go? 

Say for example you're rebuilding an application like [Plaid](https://plaid.com/?utm_source=google&amp;utm_medium=search&amp;utm_campaign=Search_G_Brand_Exact&amp;utm_content=Com&amp;utm_term=plaid&amp;utm_creative=205872015494) and you have several banks you need to store the name of and reference later

Is creating a single "const" package a bad practice here ?

    package const
    
    type BankName string
    
    const (
      BOA BankName = "bank of america"
      HSBC BankName = "the hongkong and shanghai banking corporation"
      BangkokBank BankName = "bangkok bank"  
      ...
    )
    
    // elsewhere
    func DebitAccount(bank BankName) error{
      switch bank {
       case: BOA
       // ...
      }
    }
