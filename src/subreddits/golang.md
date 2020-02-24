# golang
## [1][Golang and architecture tips for larger applications](https://www.reddit.com/r/golang/comments/f8odq4/golang_and_architecture_tips_for_larger/)
- url: https://www.reddit.com/r/golang/comments/f8odq4/golang_and_architecture_tips_for_larger/
---
I have lots of experience in Java and PHP (the good kind, Laravel, Symfony). I write fully object oriented code. OO allows me to do things like:

\- Enforce a specific architecture. For example, to create a "Processor" for some event, I extend AbstractProcessor. This AbstractProcessor class has some abstract functions,  your editor will force you to implement them, giving the Abstract Code vital information that needs to work correctly, like getProcessorName() { return "SomeObjectProcessor" }, function getRepo() { return $this-&gt;someInjectedObject }. Then the abstract code can do its work. What is good with this approach is that, I only have to extent the base class. I don't have to remember anything else, abstract code is implemented one time, then whenever I create a new Child Class, I MUST provide the info, otherwise the program won't run at all. 

I have bad memory, so I usually use lots of tricks like this and that makes my life a lot easier. All these tricks user OO and Reflection on one level or other, and its either in RunTime or "tests" that check for the validity of the data. All the problems of the type "When I do that, I must remember to do a,b,c" are usually enforced by the architecture, or the tests fail. I found that this way of programming, takes more time to setup, but maintenance is 100x easier.

Now, I fucking love Go. I have rewritten some long running small tools, that use lots of parallel threads and IO, and the performance was x20 at least, CPU and RAM. All of them are fairly simple 200-500 lines of code per tool. Go is ideal for that, easy and has excellent performance. I also want to practice more for web apps in GO.

I know this is because of luck of experience with the language, of procedural languages in general, but I find writing larger applications without OO very difficult. I know have to refresh the basics, then try to write larger things and I will find my way. I also need to study composition more, but my time is limited, so I wanted to take advantage your wisdom:

\- Did anyone had the similar issues ?

\- Do you know any book/resource that goes beyond the basics and talks more about project architecture and good practices using go ?

\- What tricks do you use to make your life easier in the future of the project ?
## [2][Fun with Functions | Frank Müller](https://www.reddit.com/r/golang/comments/f8q5t9/fun_with_functions_frank_müller/)
- url: https://www.youtube.com/watch?v=dtOQ86tLr4I
---

## [3][Extension to viper configuration library enabling go templating inside config variables](https://www.reddit.com/r/golang/comments/f8pvx3/extension_to_viper_configuration_library_enabling/)
- url: https://www.reddit.com/r/golang/comments/f8pvx3/extension_to_viper_configuration_library_enabling/
---
I needed to extend [viper](https://github.com/spf13/viper) to be able to specify randomly allocated docker port inside the viper config (basically invoke `docker port &lt;container name&gt; &lt;private port&gt;` from inside the config).

As a nice side-effect [github.com/molecule-man/vipertpl](https://github.com/molecule-man/vipertpl) library has emerged which enables one to use go templating inside viper configs.

I hope someone will find it useful.
## [4][How to run a task continuously in the background and an API in front?](https://www.reddit.com/r/golang/comments/f8oppl/how_to_run_a_task_continuously_in_the_background/)
- url: https://www.reddit.com/r/golang/comments/f8oppl/how_to_run_a_task_continuously_in_the_background/
---
Hi, I am new to go and I need to write a microservice that exposes an API and also in the background runs some tasks continuously (every 5 minutes). I want to know what is the best practice for this in go.
## [5][Gopher made by my niece](https://www.reddit.com/r/golang/comments/f876c6/gopher_made_by_my_niece/)
- url: https://i.redd.it/plyj3iek3ni41.jpg
---

## [6][[Code Review] I need an opinion from more experienced Gophers.](https://www.reddit.com/r/golang/comments/f8p4xk/code_review_i_need_an_opinion_from_more/)
- url: https://www.reddit.com/r/golang/comments/f8p4xk/code_review_i_need_an_opinion_from_more/
---
Hi, [I wrote a simple Golang code to automate GitHub commit process.](https://github.com/steadylearner/Rust-Full-Stack/blob/master/commit.go)

The entire code snippet will be similar to this.

```
package main

import (
	"bufio"
	"fmt"
	"log"
	// 1.
	"os"
	"os/exec"
	"strings"
)

func handleCombinedOutput(out []byte, err error) {
	if err != nil {
		log.Fatalf("cmd.CombinedOutput() failed with '%s'\n", err)
	}
	fmt.Printf("%s", string(out))
}

func handleReadToString(err error) {
	if err != nil {
		log.Fatalf("reader.ReadString failed with '%s'\n", err)
	}
}

func execCommand(name string, args ...string) {
        cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	handleCombinedOutput(out, err)
}

func main() {
	execCommand("git", "add", ".")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to use the default message for this commit?([y]/n)\n")
	response, err := reader.ReadString('\n')
	handleReadToString(err)
	// Is this right approach to use declare mesage here and reassign?
	// What is the problem?

	// 2. 
	// var message := "Edit README.md"
	// var message = "Edit README.md"

	// if strings.HasPrefix(response, "n") {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print("What do you want then?\n")
	// 	response, err := reader.ReadString('\n')
	// 	handleReadToString(err)
	// 	message = response
	// }

	var message string

	if strings.HasPrefix(response, "n") {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("What do you want then?\n")
		response, err := reader.ReadString('\n')
		handleReadToString(err)
		message = response
	} else {
		message = "Edit README.md"
	}

	execCommand("git", "commit", "-m", message)
	execCommand("git", "push", "-u", "origin", "master", "-f")
}
```

My questions are 

1. I am repeating "os" keyword inside import (). Is there ways to use os::{Stdin, exec} etc only to import used submodules?

2. Why these don't work? When I uncomment this part and use **var message:= "Edit REAMD.md" instead of "declaring variable with **var message string** and assign value later" it shows error. Do I miss something here?

3. Is it normal you can not use **$go run filename** after you use **$go build filename**? I couldn't use the command in this case when I have binary files built with it in the directory.  

4. Is there something I can imporve here? I know that more experienced coders can find it.

[I have a Python version of it.](https://github.com/steadylearner/Rust-Full-Stack/blob/master/commit.py) and [a blog post for that also](https://www.steadylearner.com/blog/read/How-to-automatically-commit-files-to-GitHub-with-Python).

You may refer to them if necessary.

p.s This is the first Golang code I really use for my work. I am very new to the language and please help me.
## [7][I'd like to create another course in Go, I'm discounting my book to fund this](https://www.reddit.com/r/golang/comments/f8qq39/id_like_to_create_another_course_in_go_im/)
- url: https://www.reddit.com/r/golang/comments/f8qq39/id_like_to_create_another_course_in_go_im/
---
Hello fellow Go developers,

I'm looking for topic ideas for a new course I'd like to create in Go.

I was wondering if I could get some feedback on the following 10k foot ideas.

1. Function as a Service with Google Cloud and Go: How to build and deploy a system built in Go with GCP FaaS and other GCP services (all using Go).

2. Go for C# developers: From 2001-2012 I was a corporate .NET developer and I think I have lots to say on how to help C# devs adopt Go.

3. Real-world concurrency: I've used Go to build applications since 2014 and I've used its concurrency model on real-world situations like web scapping, email processing (IMAP), message queueing, ETL etc.  Maybe I could show each of those to help new Go developers use concurrency.

Is there one of those that sounds more interesting to you? Any suggestions?

I've published a book in September 2018 name Build SaaS apps in Go and I really liked the experienced.

This one was somewhat niche, but I think for a first book/course it was perfect for me because I built SaaS for the last decade so it is a subject I know of.

I've discounted my book to hopefully been able to fund for the creation of the next course. Hoping to get higher quantity of sales at that $13 price point so I can take 1 / 1.5 month to create the next course.

I received extremely high value feedback last time when posting in here, so hopefully I can pick your brain again and start the course.

Thanks for your time and let me know your thoughts.

p.s. I'm legally blind, but I'd still prefer to try a video course for this one vs. a book, hopefully it will not be a problem :).
## [8][Do I need to develop a service from scratch to handle sign-up, login, reset password or is there a framework to use that can handle all that?](https://www.reddit.com/r/golang/comments/f8q4gt/do_i_need_to_develop_a_service_from_scratch_to/)
- url: https://www.reddit.com/r/golang/comments/f8q4gt/do_i_need_to_develop_a_service_from_scratch_to/
---

## [9][The Zen of Go | Dave Cheney](https://www.reddit.com/r/golang/comments/f8804m/the_zen_of_go_dave_cheney/)
- url: https://dave.cheney.net/2020/02/23/the-zen-of-go
---

## [10][Tango - simple dep-free CLI tool to analyze your access logs](https://www.reddit.com/r/golang/comments/f8gds6/tango_simple_depfree_cli_tool_to_analyze_your/)
- url: https://www.reddit.com/r/golang/comments/f8gds6/tango_simple_depfree_cli_tool_to_analyze_your/
---
Tango process access logs to get insights about performance and security incidents with a help of 5 built-in reports and support of Apache/Nginx Combined Log Formats.

Ready to be installed on macOS, Linux, Windows:

[https://github.com/roma-glushko/tango](https://github.com/roma-glushko/tango)

Hope you find it useful 🙌
