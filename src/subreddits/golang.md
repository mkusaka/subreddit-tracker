# golang
## [1][Go Proxyless GRPC Load Balancing on Kubernetes](https://www.reddit.com/r/golang/comments/j8dtnz/go_proxyless_grpc_load_balancing_on_kubernetes/)
- url: https://link.medium.com/7t4I71YAsab
---

## [2][grofer: A system profiler written in golang!](https://www.reddit.com/r/golang/comments/j852wz/grofer_a_system_profiler_written_in_golang/)
- url: https://www.reddit.com/r/golang/comments/j852wz/grofer_a_system_profiler_written_in_golang/
---
In the spirit of Hacktoberfest, inspired by gotop, here's a TUI based system/resource monitor, written purely in golang that we've been working on: [grofer](https://github.com/pesos/grofer)  

It provides basic stats about your system along with a few helpful metrics such as number of voluntary and involuntary context switches, time spent servicing different types of interrupts, etc. And we're currently working on building export functionality for the monitored system data in hopes of it being usable to profile remote machines. 

We've been having tons of fun developing it and learning quite a lot ourselves from the contributions that people send in! 

If anyone's interested in contributing to a beginner friendly, golang repository, please do check it out!
## [3][Python Dev needing some help with implementing a constructor.](https://www.reddit.com/r/golang/comments/j8fvim/python_dev_needing_some_help_with_implementing_a/)
- url: https://www.reddit.com/r/golang/comments/j8fvim/python_dev_needing_some_help_with_implementing_a/
---
Hi there,

I'm a Python dev but am learning Go at work. I've been seeing this type of constructor pattern seen [here](https://stackoverflow.com/questions/18125625/constructors-in-go#:~:text=In%20Go%2C%20a%20constructor%20can,pointer%20to%20a%20modified%20structure.&amp;text=For%20weak%20dependencies%20and%20better,interface%20that%20this%20structure%20implements.&amp;text=Golang%20is%20not%20OOP%20language%20in%20its%20official%20documents). 

I'm trying to implement and practice how I've been seeing this pattern at work. Here's my example code

    import (
        "fmt"
    )

    type colors interface {
        colorMethod1() byte
        colorMethod2(byte) byte
    }

    type Colors struct {
        R   byte
        G   byte
        B   byte
    }

    type Params struct {
        R   byte
        G   byte
        B   byte
    }

    func (c *Colors) colorMethod1() byte {
        return c.R
    }

    func (c *Colors) colorMethod2(b byte) byte {
        return c.B
    }

    // Constructor
    func NewColors (p Params) colors {
        return &amp;Colors{
            R:p.R,
            G:p.G,
            B:p.B,
        }
    }


    // I'm using a Jupyter Notebook so I don't need a main() func
    b := NewColors(Params{4, 6 ,7})
    b.colorMethod1()
    &gt;&gt;&gt; 4

My issue here is that when I try `b.R` I end up with 

     type main.colors has no field or method "R": b.R

Could I get some insight as to why? Here's me trying to implement it with a doggo.

https://play.golang.org/p/0MeZRePd0xC
## [4][kubecolor: a CLI tool to colorize kubectl output written in Go](https://www.reddit.com/r/golang/comments/j8khzf/kubecolor_a_cli_tool_to_colorize_kubectl_output/)
- url: https://www.reddit.com/r/golang/comments/j8khzf/kubecolor_a_cli_tool_to_colorize_kubectl_output/
---
[https://github.com/dty1er/kubecolor](https://github.com/dty1er/kubecolor)
## [5][Pagination and dynamic filtering arrays and slices for restful](https://www.reddit.com/r/golang/comments/j8hpt5/pagination_and_dynamic_filtering_arrays_and/)
- url: https://www.reddit.com/r/golang/comments/j8hpt5/pagination_and_dynamic_filtering_arrays_and/
---
[https://github.com/kohestanimahdi/go-pagination-filtering](https://github.com/kohestanimahdi/go-pagination-filtering)
## [6][A Go package for marshaling and unmarshaling map[string]string with struct tags.](https://www.reddit.com/r/golang/comments/j88glg/a_go_package_for_marshaling_and_unmarshaling/)
- url: https://github.com/chanced/labeler
---

## [7][Can gRPC be completely used in place of websockets?](https://www.reddit.com/r/golang/comments/j7wjlf/can_grpc_be_completely_used_in_place_of_websockets/)
- url: https://www.reddit.com/r/golang/comments/j7wjlf/can_grpc_be_completely_used_in_place_of_websockets/
---
I'm coming from nodejs and after implementing API in gRPC I found its more fun in building APIs with gRPC. I already knew golang for a while and was switching to golang. I am wondering if I want to implement a chat application or a server monitoring app with grpc instead of websockets, is it going to be a good idea?
## [8][Understand more about Go basics with one interview question](https://www.reddit.com/r/golang/comments/j8h35x/understand_more_about_go_basics_with_one/)
- url: https://www.pixelstech.net/article/1602313746-Understand-more-about-Go-basics-with-one-interview-question
---

## [9][todocheck v0.3.0 released](https://www.reddit.com/r/golang/comments/j8gjsp/todocheck_v030_released/)
- url: https://github.com/preslavmihaylov/todocheck/releases/tag/v0.3.0
---

## [10][Finding the single occurring int in a slice of duplicates](https://www.reddit.com/r/golang/comments/j8byrc/finding_the_single_occurring_int_in_a_slice_of/)
- url: https://www.reddit.com/r/golang/comments/j8byrc/finding_the_single_occurring_int_in_a_slice_of/
---
Hey All,

I recently came across this challenge with my friends and this is the solution we came up with.  But there are always multiple ways to come up with a solution so I thought to post here to see what you guys can come up with.

Prompt:

Given a slice of positive integers greater than 0, every element appears more than once except for one. Find the odd one out. If no such element exists, return -1.

ex.

\[\]int {10, 10, 20, 30, 30}

output:

20

&amp;#x200B;

My groups code:

    package main
    
    import "fmt"
    
    // import "fmt"
    
    func main() {
    	nums := []int{10, 10, 20, 30, 30}
    	finalAnswer := findSingleNum(nums)
    	fmt.Println(finalAnswer)
    }
    
    func findSingleNum(slice []int) int {
    	answer := -1
    	for _, v := range slice {
    		var results []int
    		for ind, val := range slice {
    			if val == v {
    				results = append(results, val)
    			}
    			if ind == len(slice)-1 &amp;&amp; len(results) == 1 {
    				answer = results[0]
    			}
    		}
    	}
    	if answer == -1 {
    		return answer
    	}
    	return answer
    }

&amp;#x200B;
