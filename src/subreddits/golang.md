# golang
## [1][wk8/go-ordered-map: Ordered Maps in Go](https://www.reddit.com/r/golang/comments/hs55vw/wk8goorderedmap_ordered_maps_in_go/)
- url: https://github.com/wk8/go-ordered-map
---

## [2][Open-source tool to automatically validate import boundaries for Go projects](https://www.reddit.com/r/golang/comments/hs5t71/opensource_tool_to_automatically_validate_import/)
- url: https://www.reddit.com/r/golang/comments/hs5t71/opensource_tool_to_automatically_validate_import/
---
Import Boundary Checker is a tool to automatically check if import boundaries are violated or not. Examples of where this tool is useful:

* **Hexagonal/clean architecture** where domain logic cannot import infrastructure code
* **Mono-repositories with multiple microservices** where you don't allow services to import code from other microservices
* **Layered architecture** where layers cannot import certain other layers

Why is this tool useful, and why consider using it over alternative tools?

* Import boundaries are checked **automatically** (meaning you don't have to spend time in code review on manually checking)
* It is extremely **fast** (sub-second speeds with medium sized projects)
* Configuration is **easy** (within a few minutes you can define forbidden imports for your projects and have everything set up)
* The tool is **independent** (you don't need to change any production or testing code for the tool to work)
* Usage of the tool requires **zero dependencies** (so no outside dependencies are needed for running the tool, only the source code of your project)

**GitHub link to project:** [**https://github.com/BytecodeAgency/import-boundary-checker**](https://github.com/BytecodeAgency/import-boundary-checker)

Feel free to give feedback on the project!
## [3][Benchmarking gRPC in Rust &amp; Go](https://www.reddit.com/r/golang/comments/hs7rp4/benchmarking_grpc_in_rust_go/)
- url: https://medium.com/@Rustling_gopher/benchmarking-grpc-in-rust-go-184545e7688a
---

## [4][Newbie going through the golang tour on slices. Which solution is better?](https://www.reddit.com/r/golang/comments/hrwz6o/newbie_going_through_the_golang_tour_on_slices/)
- url: https://www.reddit.com/r/golang/comments/hrwz6o/newbie_going_through_the_golang_tour_on_slices/
---
For [https://tour.golang.org/moretypes/18](https://tour.golang.org/moretypes/18) , I wrote two solutions.

Which is preferred and why?  And also, is there an even better way?

**SOLUTION 1**

&gt; // Use nil slices and append  
func Pic(dx, dy int) \[\]\[\]uint8 {  
	// s is nil slice referencing an array of array of uint8s  
	var s \[\]\[\]uint8  
	for i := 0; i &lt; dy; i++ {  
		// t is nil slice referencing an array of uint8s  
		var t \[\]uint8  
		for j := 0; j &lt; dx; j++ {  
			t = append(t,uint8((i + j) / 2))  
		}  
		s = append(s,t)  
	}  
	fmt.Println("dx=",dx,"dy=",dy,"s=",s)  
	return s  
}

**SOLUTION 2**

&gt; // Use make to create zeroed arrays and then overwrite  
func Pic(dx, dy int) \[\]\[\]uint8 {  
	// s is a slice to a zeroed array of size dy  
	s := make(\[\]\[\]uint8,dy)  
	for i := 0; i &lt; dy; i++ {  
		// t is a slice to a zeroed array of size dx   
		t := make(\[\]uint8, dx)  
		for j := 0; j &lt; dx; j++ {  
			t\[j\] = uint8((i + j) / 2)  
		}  
		s\[i\] = t  
	}  
	fmt.Println("dx=",dx,"dy=",dy,"s=",s)  
	return s  
}

&amp;#x200B;
## [5][The Repository pattern: a painless way to simplify your Go service logic](https://www.reddit.com/r/golang/comments/hs71iy/the_repository_pattern_a_painless_way_to_simplify/)
- url: https://threedots.tech/post/repository-pattern-in-go
---

## [6][Utility tool to print JSON / CSV formatted as table](https://www.reddit.com/r/golang/comments/hrwcxh/utility_tool_to_print_json_csv_formatted_as_table/)
- url: https://www.reddit.com/r/golang/comments/hrwcxh/utility_tool_to_print_json_csv_formatted_as_table/
---
[https://github.com/elwin/table](https://github.com/elwin/table)

While there are already tools available that do something similar I thought I'd be a nice exercise to build a utility tool that prints JSON or CSV documents formatted as a table. It's my first project that I  share publicly, and I'm very glad for any kind of feedback!
## [7][Go generics will use square brackets [] not parenthesis ()](https://www.reddit.com/r/golang/comments/hrce8e/go_generics_will_use_square_brackets_not/)
- url: https://groups.google.com/forum/#!topic/golang-nuts/7t-Q2vt60J8
---

## [8][highperforming tcp socket server with many clients and 'broadcast' functionality in go ?](https://www.reddit.com/r/golang/comments/hs3z0z/highperforming_tcp_socket_server_with_many/)
- url: https://www.reddit.com/r/golang/comments/hs3z0z/highperforming_tcp_socket_server_with_many/
---
i am in need of an ultrafast tcp-server that should handle up to 10000 live connections at the same time.

There are 3 basic task of this server :

1) continous ping during idle time between server &amp; client ( 10.000 active clients online all the time ) 

2) if one socket sends data to the server, then this data has to be sent to all other 9999 clients ( if one client 'speaks' then this will be sent to all connected clients right away ) - the requirement for this is that it has to be ultra fast.

3) small client management, the client will know if its allowed to speak or its allowed to listen to broadcasts

And ofcourse theres a packet-protocol handling of the data sent back and forth, so there will be tcp packet processing to handle the data that is being sent and passed on - but as little processing as possible here.

&amp;#x200B;

So the task itself is not very complicated to do and i can easily do a small tcp server that is doing this, but i am wondering about :

1) what is the correct term of  such a server, i would call it tcp-router, or multiplexer or broadcaster but all these seems to point to other terms, so wondering if it has a a real name

2) speed is essential here of the broadcast as i want it to be as instant as possible that data is broadcasted til all clients and want it to be snappy to connect and recognize disconnects etc.

&amp;#x200B;

I looked at the C10K problem and the Go version of handling millions of connections but i fell his examples are not quite hitting me as they arent doing anything - so maybe someone could help me along and would it be right to use Go in this case or should i go towards C or Erlang ?

&amp;#x200B;

Any examples of not a simple TCP server but something that actually scales for speed and many concurrent connections that are handling it right for scaling would also be super nice to look at as most articles is about doing basic servers which wont scale very well with real numbers.
## [9][How to use {{ range }} within a {{ range }} with two []maps?](https://www.reddit.com/r/golang/comments/hrwc15/how_to_use_range_within_a_range_with_two_maps/)
- url: https://www.reddit.com/r/golang/comments/hrwc15/how_to_use_range_within_a_range_with_two_maps/
---
HI! 

I want to execute a template and gave it this struct as data:

```
type AllData struct {
	Cards    []map[string]interface{}
	Comments []map[string]interface{}
}
```

`Cards` and `Comments` are both structs themselves with multiple strings inside them. Now when I execute the template I want to iterate over all Cards structs and for each of them iterate also over all Comments (and check if a comment is matching a card).

So I tried something like this:

```
tmpl.ExecuteTemplate(w, "home.tmpl", data)
// data is of type AllData.
```

```
{{ range.Cards }}
&lt;b&gt;{{.username}}&lt;/b
    {{ range.Comments }}
          &lt;b&gt;{{.author}}&lt;/b&gt;
          {{.comment}}
      {{ else }}
            &lt;b&gt;No comments yet... :(&lt;/b&gt;
      {{ end }}
{{ else }}
Nothing to see here...
{{ end }}
```

When I debug I can see there is a Comment map inside the Comments part of the passed AllData struct, but no comments show in the rendered html file.
The iterating over the Cards works fine.
## [10][performance cost of dockerizing a Go app](https://www.reddit.com/r/golang/comments/hr5yzx/performance_cost_of_dockerizing_a_go_app/)
- url: https://www.reddit.com/r/golang/comments/hr5yzx/performance_cost_of_dockerizing_a_go_app/
---
In [this](https://levelup.gitconnected.com/docker-for-go-development-a27141f36ba9) Go+Docker tutorial, one the the [comments](https://medium.com/p/a27141f36ba9/responses/show) says:

&gt;Docker is great for development, but the performance gains you get from Golang, you lose with Docker, you may as well go php or python, since the whole point of golang is a single distributable binary. build in Docker, deploy to bare metal is my advice.

I can't help but to think that this is inaccurate. It seems like the runtime performance cost of dockerizing would be independent of what gets run inside it. You would still incur that cost regardless of the app's language.

Furthermore, the performance cost is practically negligible, according to [this SO answer](https://stackoverflow.com/questions/21889053/what-is-the-runtime-performance-cost-of-a-docker-container#answer-26149994):

&gt;An excellent 2014 IBM research paper “An Updated Performance Comparison of Virtual Machines and Linux Containers” by Felter et al. provides a comparison between bare metal, KVM, and Docker containers. The general result is: **Docker is nearly identical to native performance** and faster than KVM in every category.

Lastly, containerization is so common practice, and I've never heard/read about not containerizing due to performance concerns.

As a fairly inexperienced backend dev, I was hoping to get a few thoughts on this from more experienced devs. Thanks
