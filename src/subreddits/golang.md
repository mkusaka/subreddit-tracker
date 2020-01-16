# golang
## [1][This belongs here](https://www.reddit.com/r/golang/comments/epbny1/this_belongs_here/)
- url: https://gfycat.com/happygoluckyhelpfulgull
---

## [2][Go gRPC server project structure.](https://www.reddit.com/r/golang/comments/epfryk/go_grpc_server_project_structure/)
- url: https://www.reddit.com/r/golang/comments/epfryk/go_grpc_server_project_structure/
---
I have read the tutorial on the official gRPC site and seen videos on youtube about how to work with gRPC in Go. But in all those tutorials they put all the code in a single file and no pattern or guideline is followed regarding the project structure. How do we structure the project for better maintainability and scalability?
## [3][Go library for Parsing Ansible inventory files](https://www.reddit.com/r/golang/comments/epg4d2/go_library_for_parsing_ansible_inventory_files/)
- url: https://github.com/relex/aini
---

## [4][gRPC Test Framework For Microservices](https://www.reddit.com/r/golang/comments/ep7eml/grpc_test_framework_for_microservices/)
- url: https://github.com/smallinsky/mtf
---

## [5][[Question] Profiling/Benchmarking TCP server](https://www.reddit.com/r/golang/comments/epihgm/question_profilingbenchmarking_tcp_server/)
- url: https://www.reddit.com/r/golang/comments/epihgm/question_profilingbenchmarking_tcp_server/
---
Hi, I'm trying to implement a [SSH Tarpit](https://nullprogram.com/blog/2019/03/22/). It's a really simple SSH Server which just writes some data to the client every few seconds. 

    
    package main
    
    import (
    	"log"
    	"net"
    	"time"
    )
    
    const (
    	ADDRESS = ":2222"
    	DELAY   = time.Second * 3
    )
    func main() {
    	log.Println("Starting")
    	ll, err := net.Listen("tcp", ADDRESS)
    	defer close(ll)
    	if err != nil {
    		log.Println("err in opening socket: ", err)
    	}
    
    	for {
    		conn, err := ll.Accept()
    		defer conn.Close()
    		if err != nil {
    			log.Println("err in accepting socket: ", err)
    		}
    		go handle(conn)
    	}
    }
    func close(ll net.Listener) {
    	if err := ll.Close(); err != nil {
    		log.Println("err in closing socket: ", err)
    	}
    }
    func handle(conn net.Conn) {
    	log.Println("conn from: ", conn.LocalAddr().String())
    	for {
    		conn.Write([]byte{'a'})
    		time.Sleep(DELAY)
    	}
    }
I'm trying to figure out how memory grows as more clients connect to this. Is it possible to use the testing framework to view the memory of this application as the number of clients go? 

The only way I can think of doing this is to create a separate application which just listens to a connection on a new goroutine and run a profiler on the server application.

Any help is appreciated :)
## [6][gojsonq: A simple Go package to Query over JSON/YAML/XML/CSV Data, Released v2.5.0](https://www.reddit.com/r/golang/comments/epfrvk/gojsonq_a_simple_go_package_to_query_over/)
- url: https://github.com/thedevsaddam/gojsonq
---

## [7][Writing struct to a YAML file](https://www.reddit.com/r/golang/comments/epj1a6/writing_struct_to_a_yaml_file/)
- url: https://www.reddit.com/r/golang/comments/epj1a6/writing_struct_to_a_yaml_file/
---
I'm really new to go and it has been less than a week. I want to write a struct into a YAML file and I've tried to follow the documentation. It still doesn't work. Can someone please help?
## [8][Interceptors and middleware for database/sql](https://www.reddit.com/r/golang/comments/epiyig/interceptors_and_middleware_for_databasesql/)
- url: https://github.com/ngrok/sqlmw
---

## [9][series of articles: Micro in Action, part2](https://www.reddit.com/r/golang/comments/ephmho/series_of_articles_micro_in_action_part2/)
- url: https://itnext.io/micro-in-action-part-2-71230f01d6fb
---

## [10][How to do SQL properly and clean and tidy in Golang?](https://www.reddit.com/r/golang/comments/ep6q6d/how_to_do_sql_properly_and_clean_and_tidy_in/)
- url: https://www.reddit.com/r/golang/comments/ep6q6d/how_to_do_sql_properly_and_clean_and_tidy_in/
---
Hi!

**TL;DR: Basically the title. This got a lot longer than I thought it would so feel free to ignore all of that and focus on the title! Thanks. it got so long because I really had a lot of problems with this the last time I tried Go for webdev and never found answers to a lot of questions I've had ever since.**

So I'm starting a new pet project and for once have actual use for it even in the MVP stage and it might keep me busy and interested for a long time because it has a lot of potential (in terms of a platform for my personal interest. Not monetary).

I really liked what I did so far in Golang but I really, really hated the SQL part of it. I thought about going for MongoDB for a while but had to admit to myself that I'm not looking for a reason to use MongoDB but for something that gives me an excuse to not do SQL in Go. So here we are back in SQL land.

I just have a really tough time seeing how to do SQL properly. I used sqlx last time and it helped a lot cutting down that boiler plate but then you start adding stuff like transactions and it felt like I'm defining my data twice (once for the DB and then I morph it into something that works for my GraphQL API. I might be able to do that on one struct but I'm not really sure if that works so easily).

Then I also write queries myself so I'm either formatting everything into the query to avoid typos or copying a lot. I got really good with multicursors in vscode just copying row definitions and then transforming them into whatever I need for my queries. I'm just missing some kind of pattern here that makes this easier to handle.

And then you have joins. What's the best way to handle queries over multiple tables? Do I need a struct for every query? What if tables have columns with the same name? It's not a problem with SQL since you can just do table.row or add an alias but the bit of documentation I found for sqlx suggests that this is actually not something structscan can handle properly.

And transactions! It all feels dirty. I think most frameworks are doing one transaction per request. Easy to understand, for sure, but what's the best way to do that in Go? Some handler? Well I don't get a status code in my handler so I don't know if I threw an error and if I should rollback or commit. I can make everything panic but then the frontend doesn't get a proper error code. I've found some way to get the HTTP status in handlers but then I found out that the graphql library I was using actually doesn't throw a 400 status on error. Which is good because graphql is supposed to be detached from HTTP but bad because now I have to filter my response for any indication for an error. What I did found was some function magic that basically wraps the code that calls code for queries in a function and then does some things (I'm doing this from memory. Sorry) and it looked amazing... for a 10 liner in a block post. I don't want to do this in every gql mutation or query handler when we have middleware that is so easy to use and extend. What I ended up doing is passing a pointer to a bool in the context object and setting that manually to true at the end of every handler. Which is also super dirty.

I know about Gorm but a colleague of mine was very much not impressed with performance and the ergonomics of the API. Since Gorm sees to be not that popular anyway I assumed it's not worth trying and I'm absolutely not uncomfortable with writing SQL. In fact I prefer it over everything but a really well thought out ORM. So I went for sqlx.

I don't know if I'm missing some very obvious pattern here that makes this easier or if I'm just doing something else wrong but it feels like the further you go away from documentation examples and blog post snippets things just start to get ugly.

However, as a little disclaimer, in my professional life I've only ever seen really bad and really good stuff concerning databases. Never something in between. 

The really bad was a 15 years old Java application for internal enterprise only (so every user was paid to use the software and that of course shaped our priorities) that had 4 (!) ways to access the database. Static functions that would build certain objects defined by business logic, queries stuffed in Struts beans (basically data holders for the used framework), a tens of thousands of lines long property file with queries that would be stuffed into some fake DAO / ORM thingy (and I put a lot of reflection code on top of it because I'm not gonna fill 200 fields by hand thankyouverymuch. First thing I did after uni on a real job) and then finally some real ORM.

The really good was PHP and Python with Symfony and Django and developers that really knew how to work those frameworks to do what we want.

But I have no experience with good solutions if I don't have a good ORM! Like, what do people do?
