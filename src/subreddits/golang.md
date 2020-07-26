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

## [3][I made a Telegram bot to control Docker](https://www.reddit.com/r/golang/comments/hy2gbl/i_made_a_telegram_bot_to_control_docker/)
- url: https://github.com/MrMarble/teledock
---

## [4][This is my first project in Go: pong in terminal. I come from higher level languages, I am still getting familiar with the "ways of Go". Any feedback regarding code quality (and gameplay) is appreciated!](https://www.reddit.com/r/golang/comments/hxvnzd/this_is_my_first_project_in_go_pong_in_terminal_i/)
- url: https://github.com/spinzed/tpong
---

## [5][A Growing Collection of Challenges to help you learn Go and Math!](https://www.reddit.com/r/golang/comments/hxkpxt/a_growing_collection_of_challenges_to_help_you/)
- url: https://tutorialedge.net/challenges/go/
---

## [6][Another generics proposal based on their (Go team) experemental branch](https://www.reddit.com/r/golang/comments/hxvm6d/another_generics_proposal_based_on_their_go_team/)
- url: https://www.reddit.com/r/golang/comments/hxvm6d/another_generics_proposal_based_on_their_go_team/
---
In short, Go team considers a possibility not to make `type` keyword for generic params mandatory when it can be state safely this is not a slice declaration. These are the following situations:

* more than one generic type parameter
* generic type parameter is followed with a trailing comma
* when a single generic type parameter has a constraint

And this already works in the Go2 playground.

Some person (Clara Pfaff) [proposed](https://groups.google.com/g/golang-nuts/c/pxDVtPjatXo) an idea to drop leading `type` keyword completely,  introduce alias `type any = interface{}` in `builtin` and make type constraint mandatory, with using `any` as an empty constraint what allows any type.

How it looks like

https://go2goplay.golang.org/p/IQV5LTAIuDr

I must say I found this a lot better for readbility. Just look at that:


    type OrderedMap[T any] struct {
        …
    }

vs

    type OrderedMap[type T] struct {
        …
    }
## [7][This repository provides a mechanism for constructing multiple, isolated, IPFS storage instances (blockstore, filestore, DAGService) on top of a single go-datastore instance.](https://www.reddit.com/r/golang/comments/hy1dqt/this_repository_provides_a_mechanism_for/)
- url: https://github.com/filecoin-project/go-multistore
---

## [8][Introduction to Fiber - An Express-inspired web framework](https://www.reddit.com/r/golang/comments/hxp9mf/introduction_to_fiber_an_expressinspired_web/)
- url: https://youtu.be/MfFi4Gt-tos
---

## [9][Yet another flag parsing library](https://www.reddit.com/r/golang/comments/hy4ygx/yet_another_flag_parsing_library/)
- url: https://www.reddit.com/r/golang/comments/hy4ygx/yet_another_flag_parsing_library/
---
Hello!

I made a flag parsing library for some of my yet to be released open-source project. Perhaps some of you would have a use for it as well!

https://github.com/Mattemagikern/Flags

Best regards,
## [10][A simple HTTP Server to share files over WiFi via Qr Code](https://www.reddit.com/r/golang/comments/hxlxfr/a_simple_http_server_to_share_files_over_wifi_via/)
- url: https://github.com/prdpx7/go-fileserver
---

## [11][Help ID this please](https://www.reddit.com/r/golang/comments/hy3bia/help_id_this_please/)
- url: https://www.reddit.com/r/golang/comments/hy3bia/help_id_this_please/
---
Total beginner to GO here. I really like the language. I'm learning to implement a stream with go lang and stumbled across a simple example online. I understand everything else except the line with the "-&lt;" 

var waitc = make(chan struct{})

close waitc() // This was inside a conditional for the error

&lt;-waitc

What does the "&lt;-" symbol mean? Any resources where I can find this?
## [12][Best practices testing functions with large output](https://www.reddit.com/r/golang/comments/hxz2mx/best_practices_testing_functions_with_large_output/)
- url: https://www.reddit.com/r/golang/comments/hxz2mx/best_practices_testing_functions_with_large_output/
---
I work on a code base which has a substantial amount of code that is basically a pure transformation of some input to some output. Normally these are the easiest to unit test, especially with table driven tests. The issue I run into is the outputs of our code are *massive*, often 10k+ lines when written to JSON, and this is mostly a result of large, deeply nested objects rather than large lists.


These objects also happen to be protobuf generated, and changes to the expected output is fairly common - additions or minor tweaks to fields happen regularly, and every few months the entire structure may dramatically change.

I was wondering if there are any recommended patterns/examples of good tests for this type of code. Note this is a large ~100k LOC codebase with thousands of tests, so we need a fairly scalable approach.

Historically, we have had a few different patterns for testing, all of which have substantial issues:


1. Golden tests

In our case this generally means having the output dumped to file as json, then asserting that file doesn't change. This is problematic because each test, even for a tiny change or feature, results in thousands of lines of code added/changed. Updates to a single part of the code may result in many, many tests updated. Because the expect output is so easily changed, its easy to accidentally change a test and break it; essentially every change requires the developer to fully understand every test to ensure the new result is valid. Because every change impacts many tests, this also leads to frequent merge conflicts, and huge PRs that are hard to review. This does have the benefit of being very visible as to what the change was, which can make **some** PRs easy to review.

2. Golden files + JsonPath Filtering

Basically an extension of the above, where instead of dumping the whole object we just look at a portion of the object, selected by JsonPath (or similar). For example, we may check the result of `foo.bar | select(.name == "something")`. This resolves some of the problems above, but does add some complexity to learning jsonpath, as well as issues when developing the tests as there is no types/IDE support for this.

3. JsonPath directly validating

Similar to the above, we can also define standard assertions, like `foo.bar == true`. This has similar issues with above

4. Direct, specific golang code for each test

For example, writing code to check specific parts of the output. This one is really common for us, and results in a lot of fragile tests checking things like `foo[0].bar[1].baz["key"] == "something"` which makes updating the ordering/structure extremely challenging, while also not giving much real coverage. There is also a lot of code with verbose and complex loops to extract the relevant info we need

5. Matchers/Extractors

This involves basically a bunch of helper functions that either validate some data, or extract it for later validation. Possibly using things like https://onsi.github.io/gomega/#adding-your-own-matchers. We would define a lot of functions like `AssertListEqualsUnordered(GetNames(FindAllFooBars()), []string{"name1", "name2"})`. This generally leads to some pretty readable tests, and overall seems like a reasonable approach. It can be challenging to come up with the right abstractions for these matcher functions though, as well as careful maintenance to ensure we don't end up with hyper-specific things like `FindAllFooBarsWithFieldIsGreenAndSizeIsLarge()`

6. Property Testing

This would involve basically defining a series of invariants about our data output and then asserting these apply to any inputs. Ideally with some fuzzing as well. We do a little bit of this, but generally its hard/impossible to define may invariants about our data.

This ended up being a lot longer than I expected... tl;dr testing large objects is hard, are there better ways to do it?
