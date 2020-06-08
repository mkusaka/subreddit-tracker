# golang
## [1][How do you browse large GO codebase](https://www.reddit.com/r/golang/comments/gyrxeb/how_do_you_browse_large_go_codebase/)
- url: https://www.reddit.com/r/golang/comments/gyrxeb/how_do_you_browse_large_go_codebase/
---
So I am reading Kubernetes codebase. It is so large. Are there any tools to which can help me how functions interact with each other and what is the code flow.
## [2][What are your reasons for mastering Go?](https://www.reddit.com/r/golang/comments/gyxl86/what_are_your_reasons_for_mastering_go/)
- url: https://www.reddit.com/r/golang/comments/gyxl86/what_are_your_reasons_for_mastering_go/
---
Hello developers!

I am starting to get my hands on Go and I wanted to know why you guys are mastering this language. Currently I am mostly coding in Python, Java, C++ and R. I think Go is interesting but I would like to have opinions of people who work with it on a regular basis as to why it is worth mastering the language. By mastering I don't mean just being able to develop some web apps but to really be a capable and competent Go programmer. For Python it is the rapid prototyping and the ease of doing machine learning. For C++ the argument is obvious. But why should one master Go?

Please enlighten me! :)
## [3][Go has removed all uses of blacklist/whitelist and master/slave](https://www.reddit.com/r/golang/comments/gy9ylr/go_has_removed_all_uses_of_blacklistwhitelist_and/)
- url: https://go-review.googlesource.com/c/go/+/236857/
---

## [4][How to Write a CLI With “Just Enough” Architecture](https://www.reddit.com/r/golang/comments/gyxfvp/how_to_write_a_cli_with_just_enough_architecture/)
- url: https://blog.carlmjohnson.net/post/2020/go-cli-how-to-and-advice/
---

## [5][Tell me how wrong i am on database layers in GO](https://www.reddit.com/r/golang/comments/gyx4cn/tell_me_how_wrong_i_am_on_database_layers_in_go/)
- url: https://www.reddit.com/r/golang/comments/gyx4cn/tell_me_how_wrong_i_am_on_database_layers_in_go/
---
I'm pondering this for last year or so. So I made an example: [https://github.com/JakubOboza/go-database-model-example](https://github.com/JakubOboza/go-database-model-example)

I'm looking for most flexible and easy to use database layer for GO. (doesn't matter if it is NoSQL, or some sort of SQL database). I was pondering few solutions or types of solutions like DI, Having like a app/env wrapping object or just var in a package. All of them have some sort of issue usually around tests or usage. Here is my POV and I don't know if it is perfect. So I would like to ask you guys as gurus   
:D what is wrong with how i think about Database layers / access.

Example isn't perfect but i didn't want to make it complex because the purpose example.

The example of use is here: [https://github.com/JakubOboza/go-database-model-example/blob/master/example/moonlanding/moonlanding.go](https://github.com/JakubOboza/go-database-model-example/blob/master/example/moonlanding/moonlanding.go)

The example of test is here: [https://github.com/JakubOboza/go-database-model-example/blob/master/example/moonlanding/moonlanding\_test.go](https://github.com/JakubOboza/go-database-model-example/blob/master/example/moonlanding/moonlanding_test.go)  


The main question is: **"How would you improve this ? Where I'm making mistakes in your opinion."**  Obviously if you will tell me that i should have database package that models should register to and it would be only initialized one time. Yeah that is fine but for purpose of example i've cut it.
## [6][why append function returns new slice instead of mutating the underlying array of another slice?](https://www.reddit.com/r/golang/comments/gyyngm/why_append_function_returns_new_slice_instead_of/)
- url: https://www.reddit.com/r/golang/comments/gyyngm/why_append_function_returns_new_slice_instead_of/
---
A tour of Go says

&gt;If the backing array of slice is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

I've tried creating a slice and using go's built-in **append** function on that slice, it just returns a new slice instead of using the same underlying array.

  
This is my experiment regarding to my question: 

    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	d := make([]int, 0, 10)
    
    	printLenAndCap(d)
    
    	// returns new slice instead of append directly to underlying array of d
    	e := append(d, 1)
    
    	printLenAndCap(d)
    	printLenAndCap(e)
    }
    
    
    func printLenAndCap(ls []int) {
    	fmt.Printf("Length: %d Capacity: %d\n", len(ls), cap(ls))
    	fmt.Println(ls)
    }


which produces the following output:

    Length: 0 Capacity: 10
    []
    Length: 0 Capacity: 10
    []
    Length: 1 Capacity: 10
    [200]

so back to my question, why is it returning new slice? Does **d** and **e** have the same underlying array?
## [7][Context in Golang](https://www.reddit.com/r/golang/comments/gyxbeg/context_in_golang/)
- url: https://www.reddit.com/r/golang/comments/gyxbeg/context_in_golang/
---
Guys, I am trying to read about the ***goalng context***, but not getting exact tutorials or resources to learn about. 

As per my understanding the context allows to pass the context of program and it helps to identify the timeout or deadline so that program can be stopped to avoid any cascading effect. 

But still I am not clear about the context, 

Please suggest If I am missing something. 

I found some resources, 

 [http://p.agnihotry.com/post/understanding\_the\_context\_package\_in\_golang/](http://p.agnihotry.com/post/understanding_the_context_package_in_golang/) 

 [https://blog.golang.org/context](https://blog.golang.org/context) 

But still not clear about context. 

Please help me to learn about the context.
## [8][Nominal Typing vs Duck Typing vs Structural Typing](https://www.reddit.com/r/golang/comments/gyvfl1/nominal_typing_vs_duck_typing_vs_structural_typing/)
- url: https://medium.com/higher-order-functions/duck-typing-vs-structural-typing-vs-nominal-typing-e0881860bf10
---

## [9][Console #2: How to Write a JVM (in Go)](https://www.reddit.com/r/golang/comments/gydkev/console_2_how_to_write_a_jvm_in_go/)
- url: https://youtu.be/TbWTtJupz44
---

## [10][Golang Fprint, Fprintf and Fprintln | Go FMT Package](https://www.reddit.com/r/golang/comments/gywg6b/golang_fprint_fprintf_and_fprintln_go_fmt_package/)
- url: https://www.reddit.com/r/golang/comments/gywg6b/golang_fprint_fprintf_and_fprintln_go_fmt_package/
---
The Golang FMT package Fprint function takes a writer and variadic parameter in the form of empty interfaces and returns the number of bytes written as integer and an error with Golang error type. 

Read More [Golang Fprint](https://divyanshushekhar.com/golang-fprint-fprintf-fprintln-fmt-package/)
