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

## [3][Printable copy of 'A tour of Go'](https://www.reddit.com/r/golang/comments/i2x0av/printable_copy_of_a_tour_of_go/)
- url: https://www.reddit.com/r/golang/comments/i2x0av/printable_copy_of_a_tour_of_go/
---
Hi,

The tutorial on 'https://tour.golang.org/' is really helpful, but is there a way to print the tutorial in a nice way?

Sometimes it's easier to learn something that is printed.

Thx
## [4][Mission critical application needs unique IDs but due to constrains only 20 chars are available and my IT can't find a reliable system alternative to UUIDs](https://www.reddit.com/r/golang/comments/i2wzz6/mission_critical_application_needs_unique_ids_but/)
- url: https://www.reddit.com/r/golang/comments/i2wzz6/mission_critical_application_needs_unique_ids_but/
---
In our company we will need to uniquely identify millions of entries for a new job, however, due to technical and external constrains we have only 20 bytes free where to store a unique ID. No room for negotiation.

UUIDs obviusly are out of the game, and I'm browsing everywhere to find other unique ID systems that could fit the requirements.

I do know that we could develop in-house a unique ID generation using Go math and rand and some boilerplate, however given the mission critical application we would prefer to stick with a know scheme rather than a custom on-the-fly implementation.

After a day of  researching my two IT guys didn't come up with anything, hence I'm landing here to the subreddit for feedback from a wider userbase.
## [5][Using GitHub Actions to compile, sign, and notarize your MacOS Golang Binaries](https://www.reddit.com/r/golang/comments/i2a2fx/using_github_actions_to_compile_sign_and_notarize/)
- url: https://www.kencochrane.com/2020/08/01/build-and-sign-golang-binaries-for-macos-with-github-actions/
---

## [6][My new US addresses segments parser](https://www.reddit.com/r/golang/comments/i2jo1o/my_new_us_addresses_segments_parser/)
- url: https://www.reddit.com/r/golang/comments/i2jo1o/my_new_us_addresses_segments_parser/
---
Hey guys!I just created a library to get US addresses segments.

Like if you have an address and you want just the street number, prefix, direction suffix :D.It's based it off USPS's validation.

I do plan on updating it to validate address ranges and stuff later on.

https://github.com/0syntrax0/go-address-parser
## [7][Check out my Go app that lets you split files into horcruxes](https://www.reddit.com/r/golang/comments/i26l3g/check_out_my_go_app_that_lets_you_split_files/)
- url: https://www.reddit.com/r/golang/comments/i26l3g/check_out_my_go_app_that_lets_you_split_files/
---
Not too long ago I made a program in Go called Horcrux which allows you to split a file into multiple 'horcruxes' (aka fragments) which can then be recombined to obtain the original file (as opposed to encrypting a single file with a password).  


Somebody urged me to create a GUI for it so I've done just that! I've used [https://fyne.io/](https://fyne.io/) and it's a pretty lightweight frontend. I'd prefer to use a native file select dialog but fyne currently doesn't support that. Let me know your thoughts :)  


[https://github.com/jesseduffield/horcrux-ui](https://github.com/jesseduffield/horcrux-ui)
## [8][How can I identify the genuine click on URL in SMS Text?](https://www.reddit.com/r/golang/comments/i2tf62/how_can_i_identify_the_genuine_click_on_url_in/)
- url: https://www.reddit.com/r/golang/comments/i2tf62/how_can_i_identify_the_genuine_click_on_url_in/
---
Hey all,  Maybe any of you have any idea to fix the following scenario: I sent a SMS with URL. I want to identify if the URL has a genuine click. Currently some SMS App, renders the URL Content and the problem comes here. Rendering the URL content by the app is considered as a click (even though the user has not clicked on the link).

How can I identify the genuine click?
## [9][go-monads - implementation of basic monads on go generics draft](https://www.reddit.com/r/golang/comments/i2fqbs/gomonads_implementation_of_basic_monads_on_go/)
- url: https://www.reddit.com/r/golang/comments/i2fqbs/gomonads_implementation_of_basic_monads_on_go/
---
[https://github.com/OlegStotsky/go-monads](https://github.com/OlegStotsky/go-monads)
## [10][grule-rule-engine version 1.5. 0 released. Grule is a Rule Engine library for the Golang programming language. Inspired by the acclaimed JBOSS Drools, done in a much simple manner.](https://www.reddit.com/r/golang/comments/i2634a/gruleruleengine_version_15_0_released_grule_is_a/)
- url: https://github.com/hyperjumptech/grule-rule-engine/blob/master/CHANGELOG.md#150---2020-08-02
---

## [11][glog - a (w.i.p.) tool that turns any GitHub repository into a blog](https://www.reddit.com/r/golang/comments/i2jn1x/glog_a_wip_tool_that_turns_any_github_repository/)
- url: https://github.com/carlos-menezes/glog
---

## [12][Does we need DI in project?](https://www.reddit.com/r/golang/comments/i2rgg8/does_we_need_di_in_project/)
- url: https://www.reddit.com/r/golang/comments/i2rgg8/does_we_need_di_in_project/
---
Hi a new gopher here.  I feel that dependency injection is inconsistent with the idea of ​​go.  
I want to know what is your attitude towards DI? Is it frequently used in go projects?
