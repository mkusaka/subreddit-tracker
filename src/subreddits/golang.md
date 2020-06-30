# golang
## [1][The absence of const reference parameters in Go](https://www.reddit.com/r/golang/comments/hikq9s/the_absence_of_const_reference_parameters_in_go/)
- url: https://www.reddit.com/r/golang/comments/hikq9s/the_absence_of_const_reference_parameters_in_go/
---
Why does Go not support const reference parameters similar to C++? Passing by value has overhead in making a copy and passing by pointer sacrifices the benefits of immutability. For a language that prides itself on performance and concurrency, const reference seems like such an obviously needed feature.

    // Simple Go struct
    type Player struct {
    	Name            string
    	Team            string
    	Age             int
    	PointsPerGame   float64
    	AssistsPerGame  float64
    	ReboundsPerGame float64
    }
    
    // An entire new Player Struct is created everytime function is called
    func (p Player) SomeFunction() string {
        return fmt.Sprintf("%s is %d years old.", p.Name, p.Age)
    }
    
    // Vs pointer where copy is avoided but now allowing function to alter the struct
    // when that's not necessary.
    func (p *Player) SomeFunction() string {
        return fmt.Sprintf("%s is %d years old.", p.Name, p.Age)
    }
    
    // Why not have the best of both worlds with something like
    func (p const *Player) SomeFunction() string {
        return fmt.Sprintf("%s is %d years old.", p.Name, p.Age)
    }

Maybe I am overestimating the amount of resources spent on making a copy. But even in cases where the struct is small, it bothers me to think that anytime at all is being spent on making an unnecessary copy. Do you guys find yourselves using predominantly pointers (even when the function doesn't modify the struct) or is worrying about this just me trying to prematurely optimize?

Any input is appreciated, I come from a Java and C++ background so the idea of pass by value for anything outside of primitive variable types is foreign to me.
## [2][A sample task engine](https://www.reddit.com/r/golang/comments/hifopt/a_sample_task_engine/)
- url: https://www.reddit.com/r/golang/comments/hifopt/a_sample_task_engine/
---
I am writing a multi-coroutine task processing engine. Anyone want with me together.

I have 2+ years experience with golang, but i am just like a beginner. Maybe this is the philosophy of golang.

https://github.com/90634/gotaskengine

This is my first post. Is this correct behavior?
## [3][What are the limits of channels, and just how 'fast' are they?](https://www.reddit.com/r/golang/comments/hilolk/what_are_the_limits_of_channels_and_just_how_fast/)
- url: https://tpaschalis.github.io/channels-limitations-speed/
---

## [4][Question about Go channels.](https://www.reddit.com/r/golang/comments/hikks9/question_about_go_channels/)
- url: https://www.reddit.com/r/golang/comments/hikks9/question_about_go_channels/
---
Hi, I am learning Go and I have the following example

    package main
    import (
    	"fmt"
    	"sync"
    )
    var wg = sync.WaitGroup{}
    func main() {
    	ch := make(chan int)
    	counter := 0
    	for j := 0; j &lt; 5; j++ {
    		wg.Add(2)
    		go func() {
    			i := &lt;-ch
    			fmt.Println(i)
    			wg.Done()
    		}()
    		go func() {
    			ch &lt;- counter
    			counter++
    			wg.Done()
    		}()
    	}
    	wg.Wait()
    }

the output is   


    0
    1
    1
    0
    2

How can the output contain two 0's and two 1's, because numbers 0,1,2,3,4 enter the channel exactly once and after receiving a number from the channel the channel becomes empty. Thanks in advance.
## [5][Senior Remote Golang Job](https://www.reddit.com/r/golang/comments/hil0ox/senior_remote_golang_job/)
- url: https://www.works-hub.com/jobs/remote-senior-go-engineer-e69?utm_source=Linkedin&amp;utm_medium=Recruiter_Social&amp;utm_campaign=p.gubbey
---

## [6][Multi-Select Facet with Solr, Vue and Go](https://www.reddit.com/r/golang/comments/hijv8c/multiselect_facet_with_solr_vue_and_go/)
- url: https://stevenferrer.github.io/posts/multi-select-facet-solr-vue-go
---

## [7][[Q] files with _debug.go suffix](https://www.reddit.com/r/golang/comments/hikygn/q_files_with_debuggo_suffix/)
- url: https://www.reddit.com/r/golang/comments/hikygn/q_files_with_debuggo_suffix/
---
I've come across a project with files ending with _debug.go. For example there are two files in a folder. One is called security.go, the other is security_debug.go. Both these files have the same function signatures, so I'm not sure how the package even compiles The debug file has slightly modified function definitions.   

I can't really find a reference to this. Are _debug files a golang feature similar to _test.go, or am I looking at some in-house solution, maybe an IDE feature?
## [8][The How and Why of Go, Part 1: Tooling](https://www.reddit.com/r/golang/comments/hin7sn/the_how_and_why_of_go_part_1_tooling/)
- url: http://okigiveup.net/the-how-and-why-of-go-part-1-tooling/
---

## [9][Vendor is suggesting concurrency for api calls](https://www.reddit.com/r/golang/comments/hin4ke/vendor_is_suggesting_concurrency_for_api_calls/)
- url: https://www.reddit.com/r/golang/comments/hin4ke/vendor_is_suggesting_concurrency_for_api_calls/
---
 I have an api project using go-chi/chi and one of the calls is registering a user by inserting into a table. The performance is quite poor about 20 users/sec and the hosting vendor suggested we use concurrency libraries to fix this issue because the app is not using more memory available on the machine. Maybe I'm missing something but I've not seen concurrency for simple api calls and such usage on the interwebs.
## [10][My second golang project.](https://www.reddit.com/r/golang/comments/himyba/my_second_golang_project/)
- url: https://www.reddit.com/r/golang/comments/himyba/my_second_golang_project/
---
Check out [https://github.com/FarzamAlam/short-url](https://github.com/FarzamAlam/short-url) 

Thank you golang community for making such amazing libraries and learning content free and opensource.
