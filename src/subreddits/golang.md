# golang
## [1][New Case Studies About Google’s Use of Go](https://www.reddit.com/r/golang/comments/ii300l/new_case_studies_about_googles_use_of_go/)
- url: https://opensource.googleblog.com/2020/08/new-case-studies-about-googles-use-of-go.html
---

## [2][Elf binary parser written in Go](https://www.reddit.com/r/golang/comments/ii1r9f/elf_binary_parser_written_in_go/)
- url: https://github.com/sad0p/go-readelf
---

## [3][Using Go at Google](https://www.reddit.com/r/golang/comments/ihppee/using_go_at_google/)
- url: https://go.dev/solutions/google/
---

## [4][GLab is an open source Gitlab Cli tool written in Go to help work seamlessly with Gitlab from the command line](https://www.reddit.com/r/golang/comments/ii4yzm/glab_is_an_open_source_gitlab_cli_tool_written_in/)
- url: https://github.com/profclems/glab
---

## [5][Commandline tool that makes building tilesets and rendering static square tilemaps super easy!](https://www.reddit.com/r/golang/comments/ihlo2s/commandline_tool_that_makes_building_tilesets_and/)
- url: https://i.redd.it/gl77fociyjj51.png
---

## [6][Writing byte streams in a type-safe and extensible way](https://www.reddit.com/r/golang/comments/ii5ir9/writing_byte_streams_in_a_typesafe_and_extensible/)
- url: https://github.com/maxim2266/stout
---

## [7][How/why is concurrent code faster than sequential code?](https://www.reddit.com/r/golang/comments/ii64br/howwhy_is_concurrent_code_faster_than_sequential/)
- url: https://www.reddit.com/r/golang/comments/ii64br/howwhy_is_concurrent_code_faster_than_sequential/
---
Hello, I’m a beginner at Go (kind of a beginner at programming in general) and I’m following a udemy course to learn Go and I’ve reached a section on concurrency. I didn’t feel like I fully understood what concurrency is or how it’s different from parallelism so I did more research and here’s what I understood so far: 

Concurrency is a way of writing a program to have independent units of code that can be run in an arbitrary order, now these units can either be run on one processor by just switching between the different units until they are all done or they can be run on multiple processors (each unit on one processor). Running these units on multiple processors is what parallelism is.

My question is, when there is only one processor how/why does concurrent code run faster than sequential code? since in both cases only one thing is running at a time. 

In everything I read or watched everyone would say that concurrency is faster but never explain how/why... I’m guessing because the reason is supposed to be intuitive once you understand concurrency but I still don’t get it :/

Thank you in advance for any answers.
## [8][Passing arguments as values vs pointers](https://www.reddit.com/r/golang/comments/ihyi63/passing_arguments_as_values_vs_pointers/)
- url: https://www.reddit.com/r/golang/comments/ihyi63/passing_arguments_as_values_vs_pointers/
---
Couldn’t find much information on this. If I have a function that takes an argument and doesn’t mutate it, is there any reason to not use a value?

In this case, in the wrapping scope, the variable is a pointer, and I dereference it to pass it to the function as a value. 

Ex

    Func foo() {
        bar := getBar() // returns pointer to a decent sized struct
        DoSomething(*bar)
    }

My reasoning is that passing the value is very explicit in that the caller knows it will not be mutated. A co worker suggested to use a pointer in a code review. 

Genuinely interested in if this is a good practice. 

Is there a cost to dereferencing? 

The other argument I could find was that passing a large struct by value is more expensive than a pointer to that struct. 

Is explicitly passing values to functions that don’t mutate a best practice?
## [9][Fast AVX2/SSE2 positional popcount routines for Go](https://www.reddit.com/r/golang/comments/ihqmxr/fast_avx2sse2_positional_popcount_routines_for_go/)
- url: https://github.com/fuzxxl/pospop
---

## [10][monsoon - a fast and flexible HTTP enumerator written in Go](https://www.reddit.com/r/golang/comments/ihnmri/monsoon_a_fast_and_flexible_http_enumerator/)
- url: https://github.com/RedTeamPentesting/monsoon
---

