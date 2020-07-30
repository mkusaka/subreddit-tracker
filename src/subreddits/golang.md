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
## [2][Design Draft: First Class Fuzzing](https://www.reddit.com/r/golang/comments/hvpr96/design_draft_first_class_fuzzing/)
- url: https://go.googlesource.com/proposal/+/refs/heads/master/design/40307-fuzzing.md
---

## [3][Flattening and Filtering JSON for Cleaner Types](https://www.reddit.com/r/golang/comments/i0k7rc/flattening_and_filtering_json_for_cleaner_types/)
- url: https://kgrz.io/go-json-flatten-filter-cleaner-types.html
---

## [4][Improving my Golang without a mentor or code-reviews](https://www.reddit.com/r/golang/comments/i0kpta/improving_my_golang_without_a_mentor_or/)
- url: https://www.reddit.com/r/golang/comments/i0kpta/improving_my_golang_without_a_mentor_or/
---
Hello, I spent the last month learning go from courses on pluralsight, learnt how to write tests, benchmark tests and familiarized myself with the tooling, etc.  
I'm dying to get to the level where I can understand/contribute to Kubernetes and Istio (I have some experience in consuming/using them - now I want to contribute to these projects), but before that I'm looking into ways that can help me improve my golang skills? how can I do that without a mentor or code-reviews? "since I'm not working, last year SE student".
## [5][Optimizing structs, tail padding makes it faster?](https://www.reddit.com/r/golang/comments/i0blo6/optimizing_structs_tail_padding_makes_it_faster/)
- url: https://www.reddit.com/r/golang/comments/i0blo6/optimizing_structs_tail_padding_makes_it_faster/
---
I noticed a pretty big performance regression in my code when I made one structure smaller. The object is being allocated often (1 million times in the benchmark), but there are other smaller objects that are allocated up to 20 million times (objects between 24 to 64 bytes).

The object used to be 216 bytes and I decreased it to 144 bytes, which gave a performance hit of 5%. I then added tail padding and noticed that padding 144+96=240 or 144+112=256 was giving that boost of 5%. Funnily, padding 144+368=512 is even faster by 2%!

What's going on here? Is this due to CPU cache line size, or due to Go's sizeclasses? Neither really explains why 512 bytes is faster than 256 bytes. If it were up to sizeclasses, surely according to [https://github.com/golang/go/blob/master/src/runtime/sizeclasses.go](https://github.com/golang/go/blob/master/src/runtime/sizeclasses.go) the original size of 144 should be sufficient? What should I be looking at? Is the benchmark function causing this behaviour (i.e. allocating the same objects for 10\*120 times)?

Interesting fact, it seems to correlate with another large object that is slightly larger (192 bytes). The following configurations are fast: 144+72=216 and 192+72=264, and 144+112=256 and 192+64=256. So they do not necessarily need to be the same size!

Extra information:

* I have an Intel i5-6300U 64-bit, and I run Arch Linux with one terminal open and Chrome, closed down most other background processes.
* The allocation is part of a JS parser that is building the AST, so lots of small allocations and pointers. Most objects are 24 bytes up to 64 bytes big, but I have four objects of around 256 bytes.
* The benchmark is parsing jquery.js about 120 times, I run the benchmark 10 times to get consistent results.
* The total memory that gets allocated is 6GB
* The structure in question is FuncDecl at [https://github.com/tdewolff/parse/blob/develop/js/ast.go#L856](https://github.com/tdewolff/parse/blob/develop/js/ast.go#L856), the other interesting structure is MethodDecl at [https://github.com/tdewolff/parse/blob/develop/js/ast.go#L881](https://github.com/tdewolff/parse/blob/develop/js/ast.go#L881). Both are being allocated often.
* The benchmark function is at [https://github.com/tdewolff/minify/blob/develop/js/js\_test.go#L478](https://github.com/tdewolff/minify/blob/develop/js/js_test.go#L478)
## [6][Kyber-K2SO is a clean implementation of the Kyber IND-CCA2-secure key encapsulation mechanism (KEM),](https://www.reddit.com/r/golang/comments/i0hnbd/kyberk2so_is_a_clean_implementation_of_the_kyber/)
- url: https://github.com/SymbolicSoft/kyber-k2so
---

## [7][Rod, a simpler way to drive headless browser](https://www.reddit.com/r/golang/comments/i03nte/rod_a_simpler_way_to_drive_headless_browser/)
- url: https://www.reddit.com/r/golang/comments/i03nte/rod_a_simpler_way_to_drive_headless_browser/
---
The project link: [https://github.com/go-rod/rod](https://github.com/go-rod/rod)

Here a side-by-side comparison between rod and chromedp, they will print the identical result (both will fail-fast and report useful message if an error happens): finding the specified section sect on the [awesome-go](https://github.com/avelino/awesome-go) page, and printing the associated projects.

For more comparison, please goto [here](https://github.com/go-rod/rod/tree/master/lib/examples/compare-chromedp).

About error handling please read our [doc](https://github.com/go-rod/rod#q-why-functions-dont-return-error-values).

[rod vs chromedp](https://preview.redd.it/6rjs12aghtd51.jpg?width=3932&amp;format=pjpg&amp;auto=webp&amp;s=21f5824613e51f68c3cf29450c858ab6d755ce9e)
## [8][4 practical principles of high-quality database integration tests in Go](https://www.reddit.com/r/golang/comments/i0ji1i/4_practical_principles_of_highquality_database/)
- url: https://threedots.tech/post/database-integration-testing/
---

## [9][[Question] Custom context.Context implementation](https://www.reddit.com/r/golang/comments/i0m0hv/question_custom_contextcontext_implementation/)
- url: https://www.reddit.com/r/golang/comments/i0m0hv/question_custom_contextcontext_implementation/
---
Hi, I'm considering to implement my MyCtx struct what would implement context.Context interface.MyCtx would be enriched with Logger and maybe the other stuff.In our app, we would pass it as MyCtx type as first func argument.So we can easily access the Logger withou using the (context.Context).Value and type cast.For (database/sql).ExecContext it could be passed also because MyCtx implents context.Context interface.

Do you think it's OK?
## [10][Finding repeating numbers](https://www.reddit.com/r/golang/comments/i0lsak/finding_repeating_numbers/)
- url: https://www.reddit.com/r/golang/comments/i0lsak/finding_repeating_numbers/
---
Hi 

&amp;#x200B;

I hope someone can help me with figuring out the logic for this function.

I need to write a function that looks at a string of numbers and identifies if any individual numbers next to each other are repeating.

&amp;#x200B;

Example string: 0123444567889

&amp;#x200B;

In the example string, there is a 1 triple of digit 4 and 1 double of the digit 8.

I have tried to do this by converting the string to a list with integers and comparing each value to the previous one. This does seem to work fine for doubles e.g. 88, however, when there is a triple-digit repeat it can still see the double inside.

&amp;#x200B;

Example code:

    package main
    
    import (
    	"fmt"
    	"strconv"
    	"strings"
    )
    
    // This turns the string into a list of integers
    func makeDigitList(number string) []int {
    	list := strings.Split(number, "")
    
    	var digitList = []int{}
    
    	for _, v := range list {
    		theInteger, err := strconv.Atoi(v)
    		if err != nil {
    			panic(err)
    		}
    		digitList = append(digitList, theInteger)
    	}
    	return digitList
    }
    
    func main() {
    
    	digitList := makeDigitList("0123444567889")
    
    	for i, v := range digitList {
                    
                    // On index zero there will be no previous value - skip
    		if i == 0 {
    
    		} else {
                            // If the index is bigger or equal to 2 look for a tripple
    			// Else look for a double
                            if i &gt;= 2 &amp;&amp; digitList[i-1] == v &amp;&amp; digitList[i-2] == v {
    				fmt.Println("Found 3")
    			} else if digitList[i-1] == v {
    				fmt.Println("Found 2")
    			}
    
    		}
    
    	}
    
    }

This returns:

Found 2

Found 3

Found 2

&amp;#x200B;

Where I need it to return: 

Found 3

Found 2

&amp;#x200B;

Any help on this is greatly appreciated.
## [11][Benchmarks - results vary when running multiple v.s. single benchmark functions](https://www.reddit.com/r/golang/comments/i0l02h/benchmarks_results_vary_when_running_multiple_vs/)
- url: https://www.reddit.com/r/golang/comments/i0l02h/benchmarks_results_vary_when_running_multiple_vs/
---
Hi All:I'm relatively new to Go, so I apologize if this is a trivial question ( [https://www.reddit.com/r/golang/comments/6ux8gc/newbie\_question\_why\_does\_golang\_benchmark\_result/](https://www.reddit.com/r/golang/comments/6ux8gc/newbie_question_why_does_golang_benchmark_result/)  seems similar, but distinct).

I am trying to benchmark using interface methods v.s. type assertions, and am seeing a strange behavior in the benchmarks:

I have 4 test cases, when I run all 4 together, the results are more or less consistently:

    BenchmarkMain-12     	  244844	      4902 ns/op	       3 B/op	       0 allocs/op
    BenchmarkMain2-12    	  122421	      9747 ns/op	       7 B/op	       0 allocs/op
    BenchmarkMain3-12    	  244813	      4892 ns/op	       3 B/op	       0 allocs/op
    BenchmarkMain4-12    	  239946	      4840 ns/op	       3 B/op	       0 allocs/op

&amp;#x200B;

If I sequentially comment out benchmarks, the results vary:

    BenchmarkMain-12     	  244836	      4793 ns/op	       3 B/op	       0 allocs/op 
    BenchmarkMain2-12    	  123693	      9768 ns/op	       6 B/op	       0 allocs/op 
    BenchmarkMain3-12    	  244839	      4943 ns/op	       3 B/op	       0 allocs/op

..   


    BenchmarkMain-12     	  249944	      4840 ns/op	       3 B/op	       0 allocs/op 
    BenchmarkMain2-12    	  121185	      9959 ns/op	       7 B/op	       0 allocs/op

&amp;#x200B;

Until the cases generally are consistent if I run them individually:

    BenchmarkMain-12    	  239943	      4936 ns/op	       3 B/op	       0 allocs/op

...

     BenchmarkMain2-12    	  230719	      4961 ns/op	       3 B/op	       0 allocs/op

..

    BenchmarkMain3-12    	  232940	      4872 ns/op	       3 B/op	       0 allocs/op
    
     ... 
    
    BenchmarkMain4-12    	  244856	      4828 ns/op	       3 B/op	       0 allocs/op

I suspect I could be running into cache coherency or similar mechanics that may explain the variation, but I'd like to ask if someone more experienced with Go would be able to offer any insights.

main.go:

    package main
    
    var baseList []baseInterface
    var embedList []embedImpl
    
    type typeID int
    
    const (
    	C_T1 typeID = iota
    	C_T2
    )
    
    type baseInterface interface {
    	GetList() []typeID
    }
    
    type baseImpl struct {
    	List []typeID
    }
    
    func (a baseImpl) GetList() []typeID {
    	return a.List
    }
    
    type embedImpl struct {
    	baseImpl
    	childList []childType
    }
    
    type childType struct {
    	A string
    }
    
    func main() {
    
    }

main\_test.go:

    package main
    
    import (
    	"log"
    	"testing"
    )
    
    func BenchmarkMain(b *testing.B) {
    	b.ReportAllocs()
    	b.ResetTimer()
    	baseList = append(baseList, baseImpl{})
    
    	a := embedImpl{}
    	for t := 0; t &lt; 10000; t++ {
    		a.List = append(a.List, C_T1)
    		a.List = append(a.List, C_T2)
    	}
    
    	baseList = append(baseList, a)
    
    	for i := 0; i &lt; b.N; i++ {
    		m := baseList[len(baseList)-1].GetList()
    		for k, _ := range m {
    			if k == 30000 {
    				log.Println(k)
    			}
    		}
    	}
    }
    
    func BenchmarkMain2(b *testing.B) {
    	b.ReportAllocs()
    	b.ResetTimer()
    	baseList = append(baseList, baseImpl{})
    
    	a := embedImpl{}
    	for t := 0; t &lt; 10000; t++ {
    		a.List = append(a.List, C_T1)
    		a.List = append(a.List, C_T2)
    	}
    
    	baseList = append(baseList, a)
    
    	for i := 0; i &lt; b.N; i++ {
    		at := baseList[len(baseList)-1].(embedImpl)
    		m := at.GetList()
    		for k, _ := range m {
    			if k == 30000 {
    				log.Println(k)
    			}
    		}
    	}
    }
    
    func BenchmarkMain3(b *testing.B) {
    	b.ReportAllocs()
    	b.ResetTimer()
    	baseList = append(baseList, baseImpl{})
    
    	a := embedImpl{}
    	for t := 0; t &lt; 10000; t++ {
    		a.List = append(a.List, C_T1)
    		a.List = append(a.List, C_T2)
    	}
    
    	baseList = append(baseList, a)
    
    	for i := 0; i &lt; b.N; i++ {
    		at := baseList[len(baseList)-1].(embedImpl)
    		m := at.List
    		for k, _ := range m {
    			if k == 30000 {
    				log.Println(k)
    			}
    		}
    	}
    }
    
    func BenchmarkMain4(b *testing.B) {
    	b.ReportAllocs()
    	b.ResetTimer()
    	a := embedImpl{}
    	for t := 0; t &lt; 10000; t++ {
    		a.List = append(a.List, C_T1)
    		a.List = append(a.List, C_T2)
    	}
    
    	embedList = append(embedList, a)
    
    	for i := 0; i &lt; b.N; i++ {
    		at := embedList[len(embedList)-1]
    		m := at.List
    		for k, _ := range m {
    			if k == 30000 {
    				log.Println(k)
    			}
    		}
    	}
    }

&amp;#x200B;

Thanks!
## [12][Go for pen-testing](https://www.reddit.com/r/golang/comments/i0knnn/go_for_pentesting/)
- url: https://www.reddit.com/r/golang/comments/i0knnn/go_for_pentesting/
---
hello guys is go lang is good for pen-testing ? and if is good what are the good lib that i should learn?
