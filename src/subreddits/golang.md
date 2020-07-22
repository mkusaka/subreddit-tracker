# golang
## [1][[Q&amp;A] io/fs draft design](https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/
---
I posted a draft design today for new file system interfaces for Go.

Video: https://golang.org/s/draft-iofs-video

Design: https://golang.org/s/draft-iofs-design

Let's do the Q&amp;A about the design here in Reddit. My hope is that the threading support will help keep questions and answers matched.

Please start a new top-level comment for each new question.

See also the related [Q&amp;A for the //go:embed draft design](https://golang.org/s/draft-iofs-reddit).
## [2][[Q&amp;A] //go:embed draft design](https://www.reddit.com/r/golang/comments/hv96ny/qa_goembed_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hv96ny/qa_goembed_draft_design/
---
I posted a draft design today for embedding files into Go programs.

Video: https://golang.org/s/draft-embed-video

Design: https://golang.org/s/draft-embed-design

Let's do the Q&amp;A about the design here in Reddit. My hope is that the threading support will help keep questions and answers matched.

Please start a new top-level comment for each new question.

See also the related [Q&amp;A for the io/fs draft design](https://golang.org/s/draft-iofs-reddit).
## [3][Go2 playground now supports square brackets for Generics](https://www.reddit.com/r/golang/comments/hvpis5/go2_playground_now_supports_square_brackets_for/)
- url: https://www.reddit.com/r/golang/comments/hvpis5/go2_playground_now_supports_square_brackets_for/
---
[https://go2goplay.golang.org](https://go2goplay.golang.org/) 

&gt;// The playground now supports parentheses or square brackets (only one at  
&gt;  
&gt;// a time) for generic type and function declarations and instantiations.  
&gt;  
&gt;// By default, parentheses are expected. To switch to square brackets,  
&gt;  
&gt;// the first generic declaration in the source must use square brackets.

Sample: [https://go2goplay.golang.org/p/7zFKUcpzhvZ](https://go2goplay.golang.org/p/7zFKUcpzhvZ)
## [4][Design Draft: First Class Fuzzing](https://www.reddit.com/r/golang/comments/hvpr96/design_draft_first_class_fuzzing/)
- url: https://go.googlesource.com/proposal/+/refs/heads/master/design/40307-fuzzing.md
---

## [5][Useful packages Gophers should know](https://www.reddit.com/r/golang/comments/hv608d/useful_packages_gophers_should_know/)
- url: https://www.golangprograms.com/go-programming-language-packages.html
---

## [6][moshebe/gebug: A tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly.](https://www.reddit.com/r/golang/comments/hv2ncw/moshebegebug_a_tool_that_makes_debugging_of/)
- url: https://github.com/moshebe/gebug
---

## [7][trouble with godoc](https://www.reddit.com/r/golang/comments/hvirja/trouble_with_godoc/)
- url: https://www.reddit.com/r/golang/comments/hvirja/trouble_with_godoc/
---
Hey everyone, I am having trouble getting godoc to work and figured I would ask if anyone has any insight.

I have a folder name test under my src folder with one package "main".

&amp;#x200B;

When I run the command

go doc -all -u ./test 

It creates documentation just fine. However when running godoc I cant seem to get anything. The documentation says to control presentation mode with the url parameter m and gives this as an example to use to get all unexported methods

For instance, [https://golang.org/pkg/math/big/?m=all](https://golang.org/pkg/math/big/?m=all) shows the documentation for all (not just the exported) declarations of package big.

&amp;#x200B;

so I ran the http server with 

godoc -http=localhost:6060   

and it hosts correctly, however when i navigate to my package both

[http://localhost:6060/pkg/test/?m=all](http://localhost:6060/pkg/test/?m=all) 

and

[http://localhost:6060/pkg/test/](http://localhost:6060/pkg/test/)

display an empty page with the header Command Test. 

Any insight would be great appreciated.
## [8][get random from given numbers](https://www.reddit.com/r/golang/comments/hvqnct/get_random_from_given_numbers/)
- url: https://www.reddit.com/r/golang/comments/hvqnct/get_random_from_given_numbers/
---
i'm new to go, how do i get a random value from a choice of numbers

ex: (1,9,36) and i get 1,9 or 36
## [9][[poll] what's your i18n/l11n solution](https://www.reddit.com/r/golang/comments/hvneur/poll_whats_your_i18nl11n_solution/)
- url: https://www.reddit.com/r/golang/comments/hvneur/poll_whats_your_i18nl11n_solution/
---
Go programmers - what do you use when writing software in Go that needs to support multiple languages? I'm tryna decide which to choose.

I'm leaning towards po files just because I don't want to ask an external translation company to write Go or manage gotext.json (as good as that looks, almost nobody knows about it and there's no docs)

[View Poll](https://www.reddit.com/poll/hvneur)
## [10][How to add a header to every request in Go](https://www.reddit.com/r/golang/comments/hvn7u3/how_to_add_a_header_to_every_request_in_go/)
- url: https://developer20.com/add-header-to-every-request-in-go/?utm_source=reddit&amp;utm_medium=link&amp;utm_campaign=headers-in-go
---

## [11][query2metric - A tool to run db queries in defined frequency and expose the count as Prometheus metrics.](https://www.reddit.com/r/golang/comments/hvdluz/query2metric_a_tool_to_run_db_queries_in_defined/)
- url: https://github.com/yolossn/query2metric/
---

## [12][Is Golang really viable for back end dev?](https://www.reddit.com/r/golang/comments/hvr0ol/is_golang_really_viable_for_back_end_dev/)
- url: https://www.reddit.com/r/golang/comments/hvr0ol/is_golang_really_viable_for_back_end_dev/
---
Hello, i did some google research and they said golang was viable for writing apis and stuff (am new to programing bare with me), but is it really viable for back end, i heard about the fiber framework is it any good?

also i know i have to learn a lot of thing in web dev (i mean a roadmap would be hepful if you guys have one)

am not looking for a job am only a 16 year old kid who wants to dive into programming (\*cough\* tried 3 years ago btw)

&amp;#x200B;

also i heard alot of people hating on golang whats the problem?

sorry if this question is asked again, Thanks for your anwsers.
