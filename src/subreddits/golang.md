# golang
## [1][Learn Golang by building a fintech banking app - Lesson3: User registration](https://www.reddit.com/r/golang/comments/gvmtxi/learn_golang_by_building_a_fintech_banking_app/)
- url: https://morioh.com/p/743f14006552?f=5c21fb01c16e2556b555ab32
---

## [2][Love at first sight](https://www.reddit.com/r/golang/comments/gvhywi/love_at_first_sight/)
- url: https://www.reddit.com/r/golang/comments/gvhywi/love_at_first_sight/
---
Sorry, this is a complete ramble.

I have recently given golang an actually chance. I had an opportunity to write golang at work over a year ago when we decided to use it in one of our services. I immediately brushed away the chance because of how golang didn't look so familiar (like C syntax languages). It breaks a few conventions that I'm familiar with.

Fast forward to literally yesterday around lunchtime and something in my head entices me to learn go. I immediately find "gobyexample" and I'm hooked. Spent the entire lunch learning about go.

Once I got home, all I can think about is learning go. I immediately join this subreddit. I don't know a lot of go, I've only written sub basic things in go (client/server ping pong deal), but I just love it.

I haven't felt this way about a language since I learned Lua all those years ago. And when I learned Java at first... I hated it (I've been doing it for 5 years now though,  so its like a spoken language to me at this point).

Do you guys still feel this way about go?
## [3][Build a distributed website with Hugo](https://www.reddit.com/r/golang/comments/gvm9mi/build_a_distributed_website_with_hugo/)
- url: https://levelup.gitconnected.com/build-a-distributed-website-with-hugo-1183bb098057
---

## [4][I wrote a post on how to write a lexer in Go](https://www.reddit.com/r/golang/comments/gv8ql1/i_wrote_a_post_on_how_to_write_a_lexer_in_go/)
- url: https://www.reddit.com/r/golang/comments/gv8ql1/i_wrote_a_post_on_how_to_write_a_lexer_in_go/
---
Hey everyone, this is my first blog post! I'd love to hear your feedback!

[https://www.aaronraff.dev/blog/how-to-write-a-lexer-in-go](https://www.aaronraff.dev/blog/how-to-write-a-lexer-in-go)
## [5][Handling Signals in Go](https://www.reddit.com/r/golang/comments/gvsnho/handling_signals_in_go/)
- url: https://www.reddit.com/r/golang/comments/gvsnho/handling_signals_in_go/
---
Hello, I just started writing blog posts and would really appreciate feedback on my first blog post in Go, titled [Handling Signals in Go](https://arberiii.github.io/go/2020/06/02/signals-go.html).
## [6][Reducing Slice Length](https://www.reddit.com/r/golang/comments/gvsjsa/reducing_slice_length/)
- url: https://www.reddit.com/r/golang/comments/gvsjsa/reducing_slice_length/
---
Hello,  I am trying to reduce a slice’s length using a maximum count parameter. Here is the code. It is quite ok. But I am open for advices.

&amp;#x200B;

    package main
    
    import "fmt"
    
    func main() {
    	records := selectRecords(1000)
    	fmt.Println(len(records))
    }
    
    func selectRecords(maximumRecordCount float64) []string {
    	records := make([]string, 5987)
    	recordCount := float64(len(records))
    	if recordCount &lt;= maximumRecordCount {
    		return records
    	}
    	step := recordCount / maximumRecordCount
    	var filteredRecords []string
    	for i := 0.0; i &lt; recordCount; i += step {
    		filteredRecords = append(filteredRecords, records[int(i)])
    	}
    	return filteredRecords
    }
## [7][Open Source MIT license Spreading Disease Simulation in golang](https://www.reddit.com/r/golang/comments/gvtluz/open_source_mit_license_spreading_disease/)
- url: https://www.reddit.com/r/golang/comments/gvtluz/open_source_mit_license_spreading_disease/
---
&amp;#x200B;

https://preview.redd.it/36syz7ht2p251.png?width=1766&amp;format=png&amp;auto=webp&amp;s=212a48b056d626f9e9b23b491bebd0e837785a6b

Hi gophers,

for an University project me and another classmate created a spreading disease simulation tool in golang, I decided to post here because I thought it could be a useful starting point for someone. At this stage it is a bit tricky to edit parameters, by the way if someone is interested can contribute in order to improve it. Star it if you like it!

[https://github.com/carbogninalberto/spreading-disease-simulation](https://github.com/carbogninalberto/spreading-disease-simulation)

**TL;DR**: there are a lot of customizable parameters, Montecarlo Simulation with Confidence Intervals calculation, Containment Measures and many more features...

The software is able to create and allocate a big graph using a relative low amount of RAM; by default the graph consist of around 4mln nodes and 150 edges per node, note that 32gb RAM is needed. Generally you can use a 400k nodes 150 edges with around 4 gb ram usage; if you suggest some feedback to reduce the amount of usage is appreciated.

We build a 4mln nodes graph because is the population in Veneto Region in Italy, and 150 edges per node is the Dunbar number, We know it's a rough rappresentation, but for our purpose was good enough; you can still load your graph in json format.

We provide a simple Montecarlo Simulation that consists on running multiple time the spreading simulation, collect data and save them on csv.

After that we run a python script to calculate confidence intervals with Asymptotic Formulas and Bootstrap algorithm. As final step the script plot some matplotlib graphs of the measured metrics.
## [8][logrusiowriter adapts logrus.Logger to be an io.Writer](https://www.reddit.com/r/golang/comments/gvt79c/logrusiowriter_adapts_logruslogger_to_be_an/)
- url: https://github.com/cabify/logrusiowriter
---

## [9][Golang Code Smells](https://www.reddit.com/r/golang/comments/gvpte8/golang_code_smells/)
- url: https://www.reddit.com/r/golang/comments/gvpte8/golang_code_smells/
---
Where can I find the code smells related to Golang?
## [10][You should not build your own authentication. Let Firebase do it for you](https://www.reddit.com/r/golang/comments/gvs4q8/you_should_not_build_your_own_authentication_let/)
- url: https://threedots.tech/post/firebase-cloud-run-authentication/
---

