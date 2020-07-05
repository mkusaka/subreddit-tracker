# golang
## [1][[Q&amp;A] //go:build draft design](https://www.reddit.com/r/golang/comments/hitexe/qa_gobuild_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hitexe/qa_gobuild_draft_design/
---
I posted a draft design today for updating the // +build lines to fix some usability problems. 

Video: [https://golang.org/s/go-build-video](https://golang.org/s/go-build-video)\
Design: [https://golang.org/s/go-build-design](https://golang.org/s/go-build-design)

As an experiment, let's try doing Q&amp;A about the design here in Reddit.
My hope is that the threading support will help keep questions and answers matched.

**Please start a new top-level comment for each new question.**
## [2][Mocking your SQL database in Go tests has never been easier.](https://www.reddit.com/r/golang/comments/hlhj8v/mocking_your_sql_database_in_go_tests_has_never/)
- url: https://github.com/cockroachdb/copyist
---

## [3][Building Go Services with DDD Approach / Eddy Kiselman](https://www.reddit.com/r/golang/comments/hlj8dc/building_go_services_with_ddd_approach_eddy/)
- url: https://youtu.be/YfLPZOpJQjY
---

## [4][I just found a website to create custom gopher's](https://www.reddit.com/r/golang/comments/hl55cc/i_just_found_a_website_to_create_custom_gophers/)
- url: https://gopherize.me/
---

## [5][Logo for a Go project](https://www.reddit.com/r/golang/comments/hlk1dr/logo_for_a_go_project/)
- url: https://www.reddit.com/r/golang/comments/hlk1dr/logo_for_a_go_project/
---
Hey folks! I'd like to resurrect [an old hobby project of mine](https://github.com/jaksi/sshesame), and I'd like a somewhat fancy logo incorporating the gopher mascot for the project (like for example [the logo](https://raw.githubusercontent.com/hybridgroup/gobot-site/master/source/images/elements/gobot-logo-small.png) for the [gobot project](https://github.com/hybridgroup/gobot/)).

Do y'all know anyone doing this kind of work? Obviously happy to pay for it.
## [6][HTML Frameworks](https://www.reddit.com/r/golang/comments/hlljvx/html_frameworks/)
- url: https://www.reddit.com/r/golang/comments/hlljvx/html_frameworks/
---
Hello Gophers! I just wanted to quickly ask if it would make any sense to build a desktop application with something like Sciter or Wails even if its just an offline app, just because i dont see any other method of creating good looking modern (really modern) GUIs
## [7][Roast my first ever Go project | A terminal-based websites monitor](https://www.reddit.com/r/golang/comments/hl516z/roast_my_first_ever_go_project_a_terminalbased/)
- url: https://www.reddit.com/r/golang/comments/hl516z/roast_my_first_ever_go_project_a_terminalbased/
---
Hello

&amp;#x200B;

I recently started fiddling around with Go. I'm loving it so far, I find myself writing code in a different, and I dare say cleaner way than I used to with other languages.

&amp;#x200B;

Now, I want the opinion of real gophers. I want to know what are the things that I'm not doing in an idiomatic way. What are the things that are so wrong about my project, that make you grunt audibly? You can roast away, but I would love it if it's based on actual constructive criticism. Thank you all in advance!

&amp;#x200B;

Link to the project: [https://gitlab.com/AyoubEd/httpmonitor](https://gitlab.com/AyoubEd/httpmonitor/)

&amp;#x200B;

Have a good weekend!
## [8][An opinionated guideline for structuring web applications/services](https://www.reddit.com/r/golang/comments/hl2fnw/an_opinionated_guideline_for_structuring_web/)
- url: https://www.reddit.com/r/golang/comments/hl2fnw/an_opinionated_guideline_for_structuring_web/
---
hello all, I made an opinionated guideline for creating/structuring a web application. I’ve seen a few out there some felt too complex, and some seemed confusing. I tried using a sample (non-functional app) app to explain how to make use of the structure. And explain why things are setup the way they are. Hope you find it useful! [https://github.com/bnkamalesh/goapp](https://github.com/bnkamalesh/goapp)
## [9][Recommendations for collecting periodic heap dumps](https://www.reddit.com/r/golang/comments/hlcln4/recommendations_for_collecting_periodic_heap_dumps/)
- url: https://www.reddit.com/r/golang/comments/hlcln4/recommendations_for_collecting_periodic_heap_dumps/
---
Just wondering if anyone here has found a tool which allows you to easily collect heap dumps of Go applications running in production/production-like environments.

I'm thinking about adding code to my services to do this and dump the heap data to an S3 bucket. However, I'm guessing that there are off the shelf tools which can do this i.e. SaaS products.

I'm fine with either approach and willing to pay for a service if it's convenient/flexible.

This is definitely something I don't want to run all the time. If I go the homegrown direction, I'll add instrumentation which is toggled via an environment variables and dumps the heap every N seconds.
## [10][Functional programming in Go [case study]](https://www.reddit.com/r/golang/comments/hl5l3o/functional_programming_in_go_case_study/)
- url: https://yourbasic.org/golang/your-basic-func/
---

## [11][What is the most idiomatic way of getting input from the terminal?](https://www.reddit.com/r/golang/comments/hkz1im/what_is_the_most_idiomatic_way_of_getting_input/)
- url: https://www.reddit.com/r/golang/comments/hkz1im/what_is_the_most_idiomatic_way_of_getting_input/
---
I have been spoilt by the simple `input()` function in Python. 

What is the most simple way of getting input from an external source? I have made several google searches and everyone often seems to recommend using different methods.

There is a `bufio.NewReader`. There is also a `bufio.NewScanner` , and then there's also a `fmt.Scanln`...
