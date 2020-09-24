# golang
## [1][Rob's frustration resulted in this beautiful language](https://www.reddit.com/r/golang/comments/iypbin/robs_frustration_resulted_in_this_beautiful/)
- url: https://www.reddit.com/r/golang/comments/iypbin/robs_frustration_resulted_in_this_beautiful/
---
I've stumbled upon a paper written by Rob Pike in 2000.

[http://herpolhode.com/rob/utah2000.pdf](http://herpolhode.com/rob/utah2000.pdf) 

I think his frustration is still relevant and looks like go language is an outcome of that. Particularly reflecting on:

1. Not depending on existing tool chain
2. Incompatible (with C) ABI
3. unconventional assembly
4. goroutine scheduler
## [2][embed, cmd/go: add support for embedded files · Issue #41191 [Accepted !]](https://www.reddit.com/r/golang/comments/iys6wu/embed_cmdgo_add_support_for_embedded_files_issue/)
- url: https://github.com/golang/go/issues/41191
---

## [3][TamaGo – bare metal Go for ARM SoCs](https://www.reddit.com/r/golang/comments/iyv2zd/tamago_bare_metal_go_for_arm_socs/)
- url: https://github.com/f-secure-foundry/tamago
---

## [4][Writing API with swagger vs by hand](https://www.reddit.com/r/golang/comments/iynmtg/writing_api_with_swagger_vs_by_hand/)
- url: https://www.reddit.com/r/golang/comments/iynmtg/writing_api_with_swagger_vs_by_hand/
---
Hello
I have been writing API with go without swagger and I love the process. But recently I got a requirement to generate code with swagger, which seemed more complex and verbose. I wanted to know how many of you use swagger to do code generation and wanted to know thought of using swagger vs writing by hand
## [5][thinkpad-led tool manager](https://www.reddit.com/r/golang/comments/iywmd5/thinkpadled_tool_manager/)
- url: https://www.reddit.com/r/golang/comments/iywmd5/thinkpadled_tool_manager/
---
Hi, I've just written this simple tool to manage the red back led of your Thinkpad (under Linux).

If the topic arouses you some interest, this is the link of my project:

[https://github.com/alegrey91/thinkpad-led](https://github.com/alegrey91/thinkpad-led)

A really thank [u/sali20](https://www.reddit.com/u/sali20/) for his amazing post ([https://www.reddit.com/r/thinkpad/comments/7n8eyu/thinkpad\_led\_control\_under\_gnulinux/](https://www.reddit.com/r/thinkpad/comments/7n8eyu/thinkpad_led_control_under_gnulinux/)), which inspired me to make my own tool.
## [6][Interface argument - copy or reference](https://www.reddit.com/r/golang/comments/iyvr60/interface_argument_copy_or_reference/)
- url: https://www.reddit.com/r/golang/comments/iyvr60/interface_argument_copy_or_reference/
---
I have question about functions which accepts interface as a argument. Will the underlying struct be copied during function call or not. How to control it?
## [7][In-process caching in Go: scaling lakeFS to 100k requests/second](https://www.reddit.com/r/golang/comments/iy9wvq/inprocess_caching_in_go_scaling_lakefs_to_100k/)
- url: https://lakefs.io/2020/09/23/in-process-caching-in-go-scaling-lakefs-to-100k-requests-second/
---

## [8][I need a second opinion on a code snippet](https://www.reddit.com/r/golang/comments/iyxaux/i_need_a_second_opinion_on_a_code_snippet/)
- url: https://www.reddit.com/r/golang/comments/iyxaux/i_need_a_second_opinion_on_a_code_snippet/
---
I am writing a little program to learn channel usage in golang. The program gives you n number of tickets for a given period. At each period, tickets are replenished whatever you use it or not. I wrote a struct to hold the tickets and refill them periodically. To replenish the tickets, I started a goroutine when a new instance of the struct was initialized. My question is that is it acceptable to start a background goroutine to update some fields of the struct periodically or are there any better ways to accomplish the same task.  you can find the sample code here:

[https://play.golang.org/p/KbW9Eoq9V7B](https://play.golang.org/p/KbW9Eoq9V7B)

thnx for your help.
## [9][SQL Migration options (2020)](https://www.reddit.com/r/golang/comments/iywzd7/sql_migration_options_2020/)
- url: https://www.reddit.com/r/golang/comments/iywzd7/sql_migration_options_2020/
---
Hey everyone-

We're starting our first Go project coming off of two fully loaded frameworks:  Rails and Django.   I'm curious what people are having success with for SQL Migrations (Postgres in our case)?

Looking through Awesome Go, these seems like the projects that are most healthy:

[https://github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate) \- seems most active and maintained

[https://github.com/rubenv/sql-migrate](https://github.com/rubenv/sql-migrate) \- was the previous recommendation for the thread I read 3 years ago, but last commit was 5 months ago so not sure if maintained

[https://github.com/go-gorm/gorm](https://github.com/go-gorm/gorm) \- full ORM with support for migrations.  Feels closest to where we are coming from.

Others?

Thanks for any advice!
## [10][Go Syntax: Literal functions, closures, and the defer keyword](https://www.reddit.com/r/golang/comments/iywuvx/go_syntax_literal_functions_closures_and_the/)
- url: https://youtu.be/CTMxVSwB4o8
---

