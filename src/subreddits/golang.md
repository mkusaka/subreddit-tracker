# golang
## [1][Would you use unmarshal nested json package like this ?](https://www.reddit.com/r/golang/comments/h8m01d/would_you_use_unmarshal_nested_json_package_like/)
- url: https://i.redd.it/gu1in1wnks451.jpg
---

## [2][Geek-life: A full featured CLI Task / To-do List Manager for Geeks](https://www.reddit.com/r/golang/comments/h8f6jl/geeklife_a_full_featured_cli_task_todo_list/)
- url: https://www.reddit.com/r/golang/comments/h8f6jl/geeklife_a_full_featured_cli_task_todo_list/
---
It's a lightweight, fast, **Terminal-UI based Golang app**.

[https://github.com/ajaxray/geek-life](https://github.com/ajaxray/geek-life)

I was finding for a generic To-do list manager in CLI with "Project =&gt; TaskList" view. Didn't find any. So, I started developing one for myself and anyone who -

* Loves and live with his dark terminal
* Thinks and writes in markdown

If you're a geek, try it. It's free and open source.

If you're a go developer, contribution and suggestion is much appreciated! (Also, don't forget to 🌟 the repo)

&amp;#x200B;

[Geek-life task manager screenshot](https://preview.redd.it/7pb7oze84r451.png?width=1270&amp;format=png&amp;auto=webp&amp;s=e5f5b2c951de6143731d0828d2118294bd2975b1)
## [3][unexpected-go.com: my wip list of stuff that a beginner doesn't expect in golang](https://www.reddit.com/r/golang/comments/h851d9/unexpectedgocom_my_wip_list_of_stuff_that_a/)
- url: https://unexpected-go.com/
---

## [4][Ardan Labs: Up &amp; Running with Dgraph](https://www.reddit.com/r/golang/comments/h8tsj8/ardan_labs_up_running_with_dgraph/)
- url: https://www.youtube.com/watch?v=BVj4b7UUHyA
---

## [5][[fsweeper] The file management automation tool](https://www.reddit.com/r/golang/comments/h8qq2b/fsweeper_the_file_management_automation_tool/)
- url: https://github.com/reugn/fsweeper
---

## [6][An Introduction to Producing and Consuming Kafka Messages in Go](https://www.reddit.com/r/golang/comments/h898df/an_introduction_to_producing_and_consuming_kafka/)
- url: https://www.aaronraff.dev/blog/an-introduction-to-producing-and-consuming-kafka-messages-in-go
---

## [7][Recommended Vue/go/mysql project structure?](https://www.reddit.com/r/golang/comments/h8qbsc/recommended_vuegomysql_project_structure/)
- url: https://www.reddit.com/r/golang/comments/h8qbsc/recommended_vuegomysql_project_structure/
---
Hey everyone. I'm new to go and come from a PHP MVC background. I have a very basic project working in go that accesses a table in mysql and passes the content via an API endpoint to a Vue component. All of the go code is written in main.go

Before I continue learning and experimenting I would like to get the basic structure of the project files correct so that I practice good habits from the get go.

Would any of you be able to recommend a decent project structure and a good example of how to split code into different kinds of responsibilities (and what those responsibilities might be).

An example in PHP that I would often use would be a folder for controllers, a router file that calls the controllers, a model folder that holds instances from the database, resources including views+js+css, repositories (and factories) which generate and manipulate models, services which abstract out controller logic into reusable methods

Many thanks
## [8][Proposal: selector-specific control flow &amp; gofmt rules to improve error-checking readability](https://www.reddit.com/r/golang/comments/h8q8o0/proposal_selectorspecific_control_flow_gofmt/)
- url: https://github.com/golang/go/issues/39372
---

## [9][Forever channel terminating the program and go channel with RabbitMQ consumer](https://www.reddit.com/r/golang/comments/h8psve/forever_channel_terminating_the_program_and_go/)
- url: https://www.reddit.com/r/golang/comments/h8psve/forever_channel_terminating_the_program_and_go/
---
I am new to GO and writing program in GO and also RMQ consumer with go channel and came to these scenarios.

**"go forever channel" blocks the main thread until it gets stop signal from other go routine".**

But below program 1 tells deadlock error, and **program 2** works fine without deadlock error,

why it's happening?  


**Program 1** : Go routine printing element in loop and deadlock error

        package main
        import "fmt"
        
        func main() {
            stopProgram := make(chan bool)
            
            go func() {
              for i := 0; i &lt; 5; i++ {
                 fmt.Println("hello ",i)
              }
              // Send signal through stopProgram to stop loop
              //stopProgram &lt;- true
            }()
    
    
           // your problem will wait here until it get stop signal through channel
           &lt;-stopProgram
           fmt.Println("after forever channel")
        }

**Output**                                                                                                                                            

    hello  0                                                                                                                                                      
        hello  1                                                                                                                                                      
        hello  2                                                                                                                                                      
        hello  3                                                                                                                                                      
        hello  4                                                                                                                                                      
        fatal error: all goroutines are asleep - deadlock!                                                                                                            
                                                                                                                                                                      
        goroutine 1 [chan receive]:                                                                                                                                   
        main.main()                                                                                                                                                   
                /home/main.go:26 +0x73                                                                                                                                
                                        

**Program 2:** Go routing receiving RMQ deliveries in loop and no deadlock

          
        package main
        import (
            "fmt"
            )
        func main() {
            conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
            failOnError(err, "Failed to connect to RabbitMQ")
            defer conn.Close()
    
            ch, err := conn.Channel()
            failOnError(err, "Failed to open a channel")
            defer ch.Close()
    
            err = ch.ExchangeDeclare(
                    "stn-msg-exc",   // name
                    "topic", // type
                    true,     // durable
                    false,    // auto-deleted
                    false,    // internal
                    false,    // no-wait
                    nil,      // arguments
            )
            failOnError(err, "Failed to declare an exchange")
    
            q, err := ch.QueueDeclare(
                    "stn-msg-queue",    // name
                    false, // durable
                    false, // delete when unused
                    true,  // exclusive
                    false, // no-wait
                    nil,   // arguments
            )
            failOnError(err, "Failed to declare a queue")
    
            err = ch.QueueBind(
                    q.Name, // queue name
                    "stn.#",     // binding key
                    "stn-msg-exc", // exchange
                    false,
                    nil,
            )
            failOnError(err, "Failed to bind a queue")
        
           msgs, err := ch.Consume(
                        q.Name, // queue
                        "",     // consumer
                        true,   // auto-ack
                        false,  // exclusive
                        false,  // no-local
                        false,  // no-wait
                        nil,    // args
            )
            
            stopProgram := make(chan bool)
        
            go func() {
                    for d := range msgs {
                          fmt.Println("reveived message ",d.Body)
                            
                    }
            }()
          
            // your problem will wait here until it get stop signal through channel
            &lt;-stopProgram
            fmt.Println("after forever channel")
        }

Can anybody clear the things here(i am new to GO) that how forever channel works here ?

**My assumption** \- In program 1, go routing ends after printing hello 5 times and there is no infinite execution in current routine/any other routine or stop signal to forever go channel.

If we want to use forever go channel(or block the main Go routine to stay in a specific go routine), we have to make sure these things 

 1. Either go routines ensures the infinite execution OR 

 2. Go routine sends the stop signal to forever channel.
## [10][GitHub Actions for Go Developers!](https://www.reddit.com/r/golang/comments/h88ted/github_actions_for_go_developers/)
- url: https://youtu.be/KVrL_UHJ7kQ
---

