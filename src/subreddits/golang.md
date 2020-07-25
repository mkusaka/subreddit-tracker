# golang
## [1][[Q&amp;A] io/fs draft design](https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/
---
I posted a draft design today for new file system interfaces for Go.

Video: https://golang.org/s/draft-iofs-video

Design: https://golang.org/s/draft-iofs-design

Let's do the Q&amp;A about the design here in Reddit. My hope is that the threading support will help keep questions and answers matched.

Please start a new top-level comment for each new question.

See also the related [Q&amp;A for the //go:embed draft design](https://golang.org/s/draft-embed-reddit).
## [2][[Q&amp;A] //go:embed draft design](https://www.reddit.com/r/golang/comments/hv96ny/qa_goembed_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hv96ny/qa_goembed_draft_design/
---
I posted a draft design today for embedding files into Go programs.

Video: https://golang.org/s/draft-embed-video

Design: https://golang.org/s/draft-embed-design

Let's do the Q&amp;A about the design here in Reddit. My hope is that the threading support will help keep questions and answers matched.

Please start a new top-level comment for each new question.

See also the related [Q&amp;A for the io/fs draft design](https://golang.org/s/draft-iofs-reddit).
## [3][A Growing Collection of Challenges to help you learn Go and Math!](https://www.reddit.com/r/golang/comments/hxkpxt/a_growing_collection_of_challenges_to_help_you/)
- url: https://tutorialedge.net/challenges/go/
---

## [4][Go 1.15 Release Candidate 1 is released](https://www.reddit.com/r/golang/comments/hx98am/go_115_release_candidate_1_is_released/)
- url: https://groups.google.com/forum/m/#!msg/golang-announce/irEtdvYep-Y/rOH3sqouCAAJ
---

## [5][A simple HTTP Server to share files over WiFi via Qr Code](https://www.reddit.com/r/golang/comments/hxlxfr/a_simple_http_server_to_share_files_over_wifi_via/)
- url: https://github.com/prdpx7/go-fileserver
---

## [6][Deploy a private GoDoc server on Google Kubernetes Engine](https://www.reddit.com/r/golang/comments/hx3uoq/deploy_a_private_godoc_server_on_google/)
- url: https://cloud.google.com/community/tutorials/godoc-for-github-on-k8s
---

## [7][what is the most productive way to learn concurrency with go?](https://www.reddit.com/r/golang/comments/hxjdmq/what_is_the_most_productive_way_to_learn/)
- url: https://www.reddit.com/r/golang/comments/hxjdmq/what_is_the_most_productive_way_to_learn/
---
Learning concurrency is little hard.

I have read "concurrency in go" book but i didn't work with me maybe good general idea.

What about you? What is the best way that make you fluent in write high performance go concurrency code?

Thanks in advance.
## [8][I created a Mandelbrot Set visualizer using distributed computing and parallelism with Go, gRPC and Raylib](https://www.reddit.com/r/golang/comments/hxl08d/i_created_a_mandelbrot_set_visualizer_using/)
- url: https://www.reddit.com/r/golang/comments/hxl08d/i_created_a_mandelbrot_set_visualizer_using/
---
I want to share with the Golang community a side-project I made to calculate the **Mandelbrot Set** using **parallelism** and **distributed computing** with **Go**, **gRPC** and **Raylib**. Here you can see a video-demo running a local cluster of 2 nodes with 8 virtual cores each one.

https://reddit.com/link/hxl08d/video/iz9elv21kzc51/player

Source code: [https://github.com/albertnadal/MandelbrotGoLang](https://github.com/albertnadal/MandelbrotGoLang)

YouTube: [https://youtu.be/pDbuClfEAIY](https://youtu.be/pDbuClfEAIY)

I hope you find it really interesting and educative.
## [9][Examples for Pointer Semantics](https://www.reddit.com/r/golang/comments/hxklwp/examples_for_pointer_semantics/)
- url: https://www.reddit.com/r/golang/comments/hxklwp/examples_for_pointer_semantics/
---
There are a lot of great resources discussing pointer vs value semantics. But I did not understand the recommendation to use pointer semantics when in doubt. So far I understood:

1. Build-in types: Value semantics
2. Reference types: Value semantics
3. Pointer semantics for decoding and unmarshalling
4. If there're already side effects like using File, which must be closed. So we need to work on the original data.

For custom types let's see this example ([https://play.golang.org/p/ySwEqcXgS0A](https://play.golang.org/p/ySwEqcXgS0A)):

    type Post struct {
    	votes int
    }
    
    func (p Post) upvoteValueSemantics(newVotes int) Post {
    	return Post{votes: p.votes + newVotes}
    }
    
    func (p *Post) upvotePointerSemantics(newVotes int) {
    	p.votes = p.votes + newVotes
    }

What semantic should be used here?
## [10][Beginning web development in Go](https://www.reddit.com/r/golang/comments/hxjr01/beginning_web_development_in_go/)
- url: https://www.reddit.com/r/golang/comments/hxjr01/beginning_web_development_in_go/
---
Hey guys,

I've been learning go and web development for a few months. But, I have been learning the web basics using **fastHttp** server and it seems like it can do everything. Then why do I see so many people use frameworks like **Gin, Gorilla and Beego**. What are some things that they have that makes people use them?
## [11][Authentication in Golang with JWTs](https://www.reddit.com/r/golang/comments/hx06h4/authentication_in_golang_with_jwts/)
- url: https://auth0.com/blog/authentication-in-golang/
---

## [12][Maximum Salary](https://www.reddit.com/r/golang/comments/hxg4k1/maximum_salary/)
- url: https://www.reddit.com/r/golang/comments/hxg4k1/maximum_salary/
---
I am doing the Algorithms and Data Structures Specialization on Coursera and I am hard stuck into problem 3.7. Maximum Salary.

Basically my code works on pretty all the test cases I can invent, but it failed in case 6 when I submit it.

`Failed case #6/11: Wrong answer (Time used: 0.00/1.50, memory used: 9138176/536870912.)`

Here you have the problem description:

&gt;Problem Description  
&gt;  
&gt;Task. Compose the largest number out of a set of integers.  
&gt;  
&gt;Input Format. The first line of the input contains an integer 𝑛. The second line contains integers  
&gt;  
&gt;𝑎1,𝑎2,...,𝑎𝑛.  
&gt;  
&gt;Constraints. 1≤𝑛≤100;1≤𝑎𝑖 ≤103 forall1≤𝑖≤𝑛.Output Format.  
&gt;  
&gt;Output the largest number that can be composed out of 𝑎1, 𝑎2, . . . , 𝑎𝑛.

And here my implementation:

    package main
    
    import (
    	"fmt"
    	"strconv"
    )
    
    func main() {
    	var n int
    	var numbers []int
    
    	_, _ = fmt.Scan(&amp;n)
    	for i := 0; i &lt; n; i++ {
    		var number int
    		_, _ = fmt.Scan(&amp;number)
    		numbers = append(numbers, number)
    	}
    
    	fmt.Print(maximumSalary(numbers))
    }
    
    func maximumSalary(numbers []int) string {
    	var answer string
    
    	for i := 0; i &lt; len(numbers); i++ {
    		maxDigit := 0
    		for j := i; j &lt; len(numbers); j++ {
    			if isGreaterOrEqual(numbers[j], maxDigit) {
    				maxDigit = numbers[j]
    				numbers[i], numbers[j] = numbers[j], numbers[i]
    			}
    		}
    
    		answer += strconv.Itoa(maxDigit)
    	}
    
    	return answer
    }
    
    func isGreaterOrEqual(i, j int) bool {
    	iString := strconv.Itoa(i)
    	jString := strconv.Itoa(j)
    	if len(iString) == len(jString) {
    		return i &gt;= j
    	} else {
    		n := 0
    		for {
    			if n == len(iString) || n == len(jString) {
    				if len(iString) &lt; len(jString) {
    					return true
    				} else {
    					return false
    				}
    			}
    			if iString[n] &gt; jString[n] {
    				return true
    			}
    			if n == 0 &amp;&amp; iString[n] &lt; jString[n] {
    				return false
    			}
    
    			n++
    		}
    	}
    }

I am really out of ideas if someone has achieved the problem before and can let me some advice, would be great.

Edit: More detailed explanation.

https://preview.redd.it/k7io9hyzuxc51.png?width=1804&amp;format=png&amp;auto=webp&amp;s=18686b14a57a56378db531cb5d675e789fd45361
