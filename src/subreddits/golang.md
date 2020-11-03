# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][Why doesn’t golang allow me to do this](https://www.reddit.com/r/golang/comments/jn6oo0/why_doesnt_golang_allow_me_to_do_this/)
- url: https://i.redd.it/z4jnotbkizw51.jpg
---

## [3][What's so great about Go? - Stack Overflow Blog](https://www.reddit.com/r/golang/comments/jmtx0h/whats_so_great_about_go_stack_overflow_blog/)
- url: https://stackoverflow.blog/2020/11/02/go-golang-learn-fast-programming-languages/
---

## [4][I made CLI that can get GitHub's pull request URL.](https://www.reddit.com/r/golang/comments/jn8qhb/i_made_cli_that_can_get_githubs_pull_request_url/)
- url: https://www.reddit.com/r/golang/comments/jn8qhb/i_made_cli_that_can_get_githubs_pull_request_url/
---
[https://github.com/skanehira/getpr](https://github.com/skanehira/getpr)

&amp;#x200B;

https://i.redd.it/s33spemwg0x51.gif
## [5][go-df data frames for Golang](https://www.reddit.com/r/golang/comments/jn6rlv/godf_data_frames_for_golang/)
- url: https://www.reddit.com/r/golang/comments/jn6rlv/godf_data_frames_for_golang/
---
After searching for a decent data-frames solution for Golang, we've decided (based on other implementations of data frames for go) to write our own "complete" solution for data frames.

[https://github.com/AdikaStyle/go-df](https://github.com/AdikaStyle/go-df)

feedback is appreciated :)
## [6][I made an animation tool with go (Ebiten)!](https://www.reddit.com/r/golang/comments/jmjica/i_made_an_animation_tool_with_go_ebiten/)
- url: https://i.redd.it/ch0y9vl59sw51.jpg
---

## [7][Bell - language based on Lisp written in Go](https://www.reddit.com/r/golang/comments/jmue7p/bell_language_based_on_lisp_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/jmue7p/bell_language_based_on_lisp_written_in_go/
---
Hey everyone. I made a pre-release of my first programming language (that I actually finished in somewhat usable form), [Bell](https://github.com/BranislavLazic/bell).   
Few notes...

Lexer and parser are written from scratch. No parser combinators were used.  
Cucumber for e2e tests prove itself to be a fantastic choice. Bringing a lot of value with very little effort. Go is totally fine, but sometimes, I missed a language with expressive ADT's and pattern matching.  


P.S. Big thanks to Thorsten Ball.
## [8][What's the status of XML/SOAP support in Go?](https://www.reddit.com/r/golang/comments/jn24j8/whats_the_status_of_xmlsoap_support_in_go/)
- url: https://www.reddit.com/r/golang/comments/jn24j8/whats_the_status_of_xmlsoap_support_in_go/
---
Unfortunately, I need to interact with an external system using the [HTNG](https://www.htng.org/) standard. As part of that, everything needs to be transmitted using SOAP (as both server and client for 2-way communication), using WS-Addressing and WS-Security. Searching around, the majority of what I've found have been projects that haven't been updated for at least 2 years, or writeups saying to do everything manually.

What recommendations would you have with regards to interacting with a SOAP interface (preferably using Go)?
## [9][I made a small Python package, which allows to use Golang (Go) functions inside python code (and run into big issues)](https://www.reddit.com/r/golang/comments/jna5r1/i_made_a_small_python_package_which_allows_to_use/)
- url: https://www.reddit.com/r/golang/comments/jna5r1/i_made_a_small_python_package_which_allows_to_use/
---
GitHub page - [https://github.com/hermanTenuki/goinpy](https://github.com/hermanTenuki/goinpy).

I like the idea in speeding up Python by moving hard computations to fast C-type language (Golang), so i made this small package based on "ctypes" module.

It's working in 3 steps:

1. Write Golang code with "C" package imported, create export comments for function that you want to have access to;
2. Compile it to c-shared library;
3. Load library in Python and use lib's functions with proper file types setup.

As turned out, this is not so easy, i quickly ran into issues when Golang is starting to mess with mutable data types and pointers.

GitHub issue - [https://github.com/hermanTenuki/goinpy/issues/6](https://github.com/hermanTenuki/goinpy/issues/6)

As said in [https://golang.org/cmd/cgo/#hdr-Passing\_pointers](https://golang.org/cmd/cgo/#hdr-Passing_pointers):

&gt; Go is a garbage collected language, and the garbage collector needs to know the location of every pointer to Go memory. Because of this, there are restrictions on passing pointers between Go and C.   
&gt;  
&gt;. . .  
&gt;  
&gt; It is possible to defeat this enforcement by using the unsafe package, and of course there is nothing stopping the C code from doing anything it likes. However, programs that break these rules are likely to fail in unexpected and unpredictable ways. 

For now, my package **can be used for simple cases**, look at [https://github.com/hermanTenuki/goinpy#how-to-use](https://github.com/hermanTenuki/goinpy#how-to-use).

I'm definitely going to continue working on this project, i want this idea to be real xd
## [10][Problem Solving in Golang](https://www.reddit.com/r/golang/comments/jn8wtz/problem_solving_in_golang/)
- url: https://shareablecode.com/browse/tags/golang+codingchallenge
---

## [11][Is assigning a large struct thread-safe?](https://www.reddit.com/r/golang/comments/jn0s21/is_assigning_a_large_struct_threadsafe/)
- url: https://www.reddit.com/r/golang/comments/jn0s21/is_assigning_a_large_struct_threadsafe/
---
I have a global variable containing a rather large struct. Several concurrent goroutines do frequent reads on this variable/struct. Also there's an update func that returns a new struct (by value) with updated content which is assigned to the global variable.

There are no writes to the variable/struct other than the re-assignment of the variable itself by the update func.

```
type RatherLargeFoo struct {
// ...
}

var globalVar RatherLargeFoo

func updateFooStruct() RatherLargeFoo {
  // assemble a refreshed RatherLargeFoo
  return RatherLargeFoo{...}
}
```

Question: Is `globalVar = updateFooStruct()` safe to do while several goroutines are reading from `globalVar` or is there an instant of time where undefined behaviour can occur? I know this is not "atomic" in the technical sense (as nothing is atomic as long as it doesn't claim to be) but as long as goroutines read either the complete old content or the complete new content, it's fine. I just don't want mixed, partial or weird results.

Also note that the assingment is by value. Would having `updateFooStruct()` return a Pointer to a new struct change anything for the better or worse in my scenario?

Thanks!
