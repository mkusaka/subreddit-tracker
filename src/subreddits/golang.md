# golang
## [1][A parallel SSH tool in Go](https://www.reddit.com/r/golang/comments/iecmjo/a_parallel_ssh_tool_in_go/)
- url: https://www.reddit.com/r/golang/comments/iecmjo/a_parallel_ssh_tool_in_go/
---
Hey folks! I have just started on Go, and I tried to make an implementation of PSSH tool.

[https://github.com/korde96/gossh](https://github.com/korde96/gossh)

It might not be the best, or the most optimised solution, so I would appreciate it if folks here would teach me the ways of Go and tell me how to make it better. Cheers!
## [2][Moving forward with the generics design draft](https://www.reddit.com/r/golang/comments/idwylv/moving_forward_with_the_generics_design_draft/)
- url: https://groups.google.com/forum/#!topic/golang-nuts/iAD0NBz3DYw
---

## [3][Postgresql PGX vs Database/SQL](https://www.reddit.com/r/golang/comments/iedksn/postgresql_pgx_vs_databasesql/)
- url: https://www.reddit.com/r/golang/comments/iedksn/postgresql_pgx_vs_databasesql/
---
I am using PGX driver to connect my go web app with Postgresql database. As I read on the package documentation, there are two options to write code - directly with PGX interface or through Database/SQL. I would like to know the pros and cons of both to decide what to use. One advantage of using Database/SQL I can see is to change the driver in future if there is any better driver or PGX driver is not maintained.
## [4][Look for feedback on a Go standard library for building microservices](https://www.reddit.com/r/golang/comments/ieiizy/look_for_feedback_on_a_go_standard_library_for/)
- url: https://www.reddit.com/r/golang/comments/ieiizy/look_for_feedback_on_a_go_standard_library_for/
---
I'm currently building a standard library for building microservices application. It's currently built around gRPC, expose both gRPC and REST API over 1 single port. This project is both for learning and using for my projects, hence I'm really appreciate if you guys can give your feedback so that I know what needs to be improved.

[https://github.com/pthethanh/micro](https://github.com/pthethanh/micro)

BTW, I've tried to ask some people I know but they didn't give much feedback, hence I think it's good to look for feedback here....

Thank you very much...and sorry if this annoy you.
## [5][gRPC or REST for a simple service](https://www.reddit.com/r/golang/comments/ieh8f0/grpc_or_rest_for_a_simple_service/)
- url: https://www.reddit.com/r/golang/comments/ieh8f0/grpc_or_rest_for_a_simple_service/
---
I am writing a service which will generate tokens. My personal requirement is no special tools are needed to access the service -- curl and wget will suffice. The server will be built using go, obviously. 

I can build a simple REST API using go-chi. Or, I can create a gRPC service and have a REST gateway. The later seems more tedious and complex. But it provides backward compatibility, schema, etc...Maybe in the future,  the service will become popular and people ask for more features. 

I like the REST gateway because I can have users access Swagger pages.
## [6][Looking for a Golang Buddy](https://www.reddit.com/r/golang/comments/ie562u/looking_for_a_golang_buddy/)
- url: https://www.reddit.com/r/golang/comments/ie562u/looking_for_a_golang_buddy/
---
I have been using Go for the past 4 months. 

  
Looking for someone with a similar level of experience (the more the better) and looking to dive into more complicated things by working on some side projects.  


If it matters, I have some experience in other languages as well, and just recently moved to go.
## [7][Need Urgent Help in defining an api endpoint using GIN](https://www.reddit.com/r/golang/comments/ieezr1/need_urgent_help_in_defining_an_api_endpoint/)
- url: https://www.reddit.com/r/golang/comments/ieezr1/need_urgent_help_in_defining_an_api_endpoint/
---
so want to make an endpoint like GET extract/&lt;url&gt; which gives a json response lile  
{text:"url"}, i.e whatever parameter , i.e the url of another website is passed in the endpoint  
it should return that paramter as a string  


I've done this so far

    r.GET("extract/:target",func(c *gin.Context){
    
    		url:=c.Param("target")
    		c.JSON(200, gin.H{
    			"text":url,
    		})
    	})

the problem is it works for trivial strings like "hello"  
but whenever there is special charecters such as "." or "/" it gives an error  
for example  


    extract/"facebook.com/profile"

gives an error  
PLease help anyone?
## [8][Learning to build restAPI in GOlang](https://www.reddit.com/r/golang/comments/iecy9v/learning_to_build_restapi_in_golang/)
- url: https://www.reddit.com/r/golang/comments/iecy9v/learning_to_build_restapi_in_golang/
---
I want to build a GET api endpoint which takes in url of another website and returns the html code of that url in a json object.

So far i've done this

    func main(){
 r:= gin.Default()
    r.GET("extract/:target",func(c *gin.Context){

     url:=c.Param("target")

     c.JSON(200, gin.H{
     "parsedText":url,
      })
    })
    r.Run()
}

when I pass extract/hello-world  
it returns hello-world and works fine  
but when I enter a website like extract/facebook.com or extract/faceboom.com/profile  
it doesnt take the whole url as a parameter because of "/" or ".com" how do I do this without breaking the query param , any help??
## [9][Newbie question: Why should I use templates instead of API?](https://www.reddit.com/r/golang/comments/idz3t4/newbie_question_why_should_i_use_templates/)
- url: https://www.reddit.com/r/golang/comments/idz3t4/newbie_question_why_should_i_use_templates/
---
Hey, I'm quite new to backend programming and I don't understand the next thing:

Why should I use template rendering if there is a frontender and it is much easier for him to just model his page and just fetch data from API? Can he use template I create?

If there is no frontend, then it is obviously good way to send page to client.

I may be completely wrong so please correct me
## [10][Gate: The extensible Minecraft proxy written in Go!](https://www.reddit.com/r/golang/comments/ideg2v/gate_the_extensible_minecraft_proxy_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/ideg2v/gate_the_extensible_minecraft_proxy_written_in_go/
---
# Today, I want to present "Gate", the extensible Minecraft proxy, to the Gophers part of the Minecraft community!

# [https://github.com/minekube/gate](https://github.com/minekube/gate)

[Every single 🌟 supports the project!](https://preview.redd.it/31xi5b36sci51.png?width=3500&amp;format=png&amp;auto=webp&amp;s=e71105ad91f79a2390a4e782c1b765052b03342f)

# Target audiences

* advanced **Minecraft networks that already (or want to) have a Go code base** for their Minecraft related workloads

Before you may ask: "*Why not use an existing proxy written in Java?*"

Because the less Java we need to maintain, the happier we are at [Minekube](https://discord.gg/6vMDqWE), since we need and work in a very fast paced and cloud-centric ecosystem with a lot of Kubernetes &amp; controllers, Protobuf &amp; GRPC, CockroachDB, Agones, Istio, Nats &amp; Stan and so forth, there is no space and time for Java. The ONLY Java code we must write is for paper/spigot plugins!

# What Gate does

*(for those who have never heard of a Minecraft proxy)*

**TL;DR**

* keep the player's session without disconnect to...
* move players between servers
* cross server functionalities (events such as chat/commands, send messages, ...builtin/custom session- &amp; packet handlers)

Gate presents itself as a normal Minecraft server in the player's server list, but once the player connects Gate forwards the connection to one of the actual game servers (e.g. Minecraft vanilla, paper, spigot, sponge, etc.) to play the game.

The player can be moved around the network of Minecraft servers **without** fully disconnecting, since we want the player to stay (and not want them to re-login via the server-list every time).

Therefore, Gate reads all packets sent between players (Minecraft client) and upstream servers, logs session state changes, emits different events like [Login, Disconnect, ServerConnect, Chat, Kick etc.](https://github.com/minekube/gate/blob/master/pkg/proxy/events.go) that custom plugins/code can react to.

The **advantages** for using a proxy should be clear now.

[Every single 🌟 supports the project!](https://preview.redd.it/rgnrvmi7sci51.png?width=1229&amp;format=png&amp;auto=webp&amp;s=d6723949d962b731ced7de9e4074cfa783dbab22)

# Features

* Fast
* Excellent server version support
   * 1.7 up to newest &amp; forge support
* Quick installation
   * Download the binary from the [releases](https://github.com/minekube/gate/releases)
   * `docker run -it --rm -p 25565:25565` `registry.gitlab.com/minekube/gate:latest`
* Simple API to [extend Gate](https://github.com/minekube/gate#extending-gate-with-custom-code)
* Built-in IP based rate limiter
* A detailed config with sane defaults

# ...please see more on the [GitHub repository](https://github.com/minekube/gate) to get started and feel free to support Gate with a 🌟 and contributions!
