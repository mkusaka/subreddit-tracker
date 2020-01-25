# golang
## [1][Getting GOing -- Learning Go in parallel with Docker](https://www.reddit.com/r/golang/comments/etgixe/getting_going_learning_go_in_parallel_with_docker/)
- url: https://www.reddit.com/r/golang/comments/etgixe/getting_going_learning_go_in_parallel_with_docker/
---
I got my CS degree \~3 years ago and have been working as a Software Engineer since. The vast majority of my work has been in using Java. I've also done various Node.js powered stuff like React. 

I think a time is coming within the next year or so where my team at work is going to get merged into an organization that exclusively uses Go. So I basically need to learn Go for job security. Also I just want to learn Go anyway because I'm interested. Docker is something I also need to learn how to use for the aforementioned reasons.

My question is: Does it make sense to start learning Go in parallel with learning Docker? Like building some simple Go applications that use Docker. Or would this likely be too overwhelming and I should focus on learning one and then the other? For reference, I'm an average developer with a bit of a imposter syndrome.
## [2][schollz/httpfileserver: Wrapper for http.FileServer that is faster (by serving from memory) and uses less bandwidth (by gzipping when possible)](https://www.reddit.com/r/golang/comments/etqk50/schollzhttpfileserver_wrapper_for_httpfileserver/)
- url: https://github.com/schollz/httpfileserver
---

## [3][How to run script on data submitted through REST Api ?](https://www.reddit.com/r/golang/comments/etpymz/how_to_run_script_on_data_submitted_through_rest/)
- url: https://www.reddit.com/r/golang/comments/etpymz/how_to_run_script_on_data_submitted_through_rest/
---
Hey gophers how can I run scriots on submitted data to server backend ? Without using golang's inbuilt command execute method.
## [4][Another circuit breaker to avoid accidents!](https://www.reddit.com/r/golang/comments/etnvs6/another_circuit_breaker_to_avoid_accidents/)
- url: https://www.reddit.com/r/golang/comments/etnvs6/another_circuit_breaker_to_avoid_accidents/
---
Went through the other circuit breakers currently available in Go, and decided to make one for myself with simplicity in mind, and features that is going to be make consuming third-party services safer. Ended up with this:  
[Cutout](https://github.com/Anondo/cutout) 

Do check it out and please  have your say, would be really helpful.
## [5][Go and Robotics/AI](https://www.reddit.com/r/golang/comments/eta159/go_and_roboticsai/)
- url: https://www.reddit.com/r/golang/comments/eta159/go_and_roboticsai/
---
I listened a podcast in Talk Python channel and the topic was about robotics and Python. The speaker mentioned about they have to use C++ for complex things because of performance of Python(Because when it comes to Robots, reaction duration must be minimal). So, scale of Python in Robotics is not that large.

I think, Golang has best of both, access to microcontrollers without sacrificing coding productivity(unlike C++) and AI libraries without big performance penalty(unlike Python).

I'm entry level Go developer and learning it for gRPC but later I have plans for AI and Robotics. So, I wanted to know what you think about this topic? Is it trendy and is it possible to get best result with Go? Should I invest in Go for this area?

Podcast link: [https://podcasts.apple.com/tr/podcast/talk-python-to-me-python-conversations-for-passionate/id979020229?i=1000452355489](https://podcasts.apple.com/tr/podcast/talk-python-to-me-python-conversations-for-passionate/id979020229?i=1000452355489)
## [6][GoMVCtemplate idea: 0 non-native dependencies, `go build &amp; go run controller.go` and your server is running your entire website](https://www.reddit.com/r/golang/comments/etnap4/gomvctemplate_idea_0_nonnative_dependencies_go/)
- url: https://www.reddit.com/r/golang/comments/etnap4/gomvctemplate_idea_0_nonnative_dependencies_go/
---
I have an entire website written in Go.

the http server is a simple library for now, so I'll have to update that to http/2 or use a reverse proxy through another http2.0 - capable server, but....

Templating in go for html is beautiful. And you can run side-by-side with static pages through the handler. 

Dynamic form data handling is really easy, and you can write a separate view controller for each posted form. 

It's really easy to reach into the OS to read/write and so the entire model can be fractured into private and public copies and set permissions for access to the private database for development of model scenarios for developing with the server still running the site (e.g. private development of the model API)

and I'm about to start writing the entire database schema into basic folders with individual JSON files scattered into folders by date and then by id or potentially a hash value fractured into descending folders, like site.domain/public/users/bio/{hash(user\_id)}

and since everyone understands DB/table/column/row/value as a descending "database schema" of relational SQL engines, it wouldn't be hard to explain this idea in the model.go file.   


What do you think? Am I Midas re-inventing the MVC wheel or am I Medusa and this a nightmare waiting to happen?
## [7][Rate a modern go web stack boilerplate I built](https://www.reddit.com/r/golang/comments/etil22/rate_a_modern_go_web_stack_boilerplate_i_built/)
- url: https://www.reddit.com/r/golang/comments/etil22/rate_a_modern_go_web_stack_boilerplate_i_built/
---
Hey everyone,

When learning Go to replace C as the go-to low-level programming language, I decided to explore web development with Go and ended up [creating a boilerplate to use in hackathons](https://github.com/msanvarov/gin-rest-prisma-boilerplate).  The purpose of this boilerplate is to quickly scaffold a REST API with basics like authentication, crud operations on profiles available from the get-go so I can forgo spending the first night setting everything up and get straight to implementing the business requirements that our team came up with.

I would love for yall to take a look, and provide suggestions on improving it. I'm currently trying to master goroutines and parallel programming with Go, to elevate this boilerplate to the next level.

Link: [https://github.com/msanvarov/gin-rest-prisma-boilerplate](https://github.com/msanvarov/gin-rest-prisma-boilerplate)

Cheers,

Sal
## [8][Squzy - is a high-performance open-source monitoring system written in Golang with Bazel and love. Help you monitoring critical api(http, grpc) and can be used for site monitoring.](https://www.reddit.com/r/golang/comments/et37ad/squzy_is_a_highperformance_opensource_monitoring/)
- url: https://github.com/squzy/squzy
---

## [9][Need help using the "debug/elf" package](https://www.reddit.com/r/golang/comments/etekz9/need_help_using_the_debugelf_package/)
- url: https://www.reddit.com/r/golang/comments/etekz9/need_help_using_the_debugelf_package/
---
I'm trying to learn Go and more about ELF binaries at the same time. I figured I'd write something similar to \`readelf\`. Nothing anybody would actually use. I started by reading an ELF file as bytes and parsing through it manually. Then I discovered the "debug/elf" package. However, I'm having trouble with figuring out how to verify the "magic" string (first 4 bytes of the file) of the file.

So I have a basic program that looks like this:

    package main
    
    import (
        "fmt"
        "os"
        "log"
        "debug/elf"
    )
    
    func main() {
        if len(os.Args) &lt; 2 {
            fmt.Fprintf(os.Stderr, "Error: Provide file name\n")
            os.Exit(1)
        }
    
        file, err := elf.Open(os.Args[1])
        defer file.Close()
        if err != nil {
            log.Fatal(err)
        }
    }

Now, looking at [https://golang.org/pkg/debug/elf/](https://golang.org/pkg/debug/elf/) I'm trying to figure out how to access the [Header64](https://golang.org/pkg/debug/elf/#Header64) struct that contains all of the header information I want. I can use `file.FileHeader` and get access to the members of that struct, but the [FileHeader](https://golang.org/pkg/debug/elf/#FileHeader) type does not contain all of the header information (e.g. the magic string).

Any help would be greatly appreciated!
## [10][The Battle for Go Knowledge Supremacy](https://www.reddit.com/r/golang/comments/etgvvj/the_battle_for_go_knowledge_supremacy/)
- url: https://go-battle.net/battle/disastrous-wombat-3-ef29
---

