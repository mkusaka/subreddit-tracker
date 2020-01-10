# golang
## [1]["How to write Go code" has been entirely re-written, and now also covers Go modules!](https://www.reddit.com/r/golang/comments/emgpj8/how_to_write_go_code_has_been_entirely_rewritten/)
- url: https://golang.org/doc/code.html
---

## [2][XML-RPC client library for Go](https://www.reddit.com/r/golang/comments/empq6x/xmlrpc_client_library_for_go/)
- url: https://www.reddit.com/r/golang/comments/empq6x/xmlrpc_client_library_for_go/
---
I know, first thing you think about when you see XML-RPC is "*what year is this?!*" and "*but why?*".

However I needed to make some XML-RPC calls to a third-party service for one of my projects. I've looked around and found a few projects that kind-of work but had weird quirks and several bugs that I would hit.

Combine the above with me wanting to experiment a bit - I've created ~~yet another http router~~ XML-RPC client library for Go: [https://github.com/alexejk/go-xmlrpc](https://github.com/alexejk/go-xmlrpc)

It most likely has some issues and improvements one could make - any feedback or PRs are  welcome!

P.S: A more extended version of this post: [https://alexejk.io/article/handling-xmlrpc-in-go/](https://alexejk.io/article/handling-xmlrpc-in-go/)
## [3][It Looks Like GCC's Long-Awaited Git Conversion Could Happen This Weekend](https://www.reddit.com/r/golang/comments/emm8ae/it_looks_like_gccs_longawaited_git_conversion/)
- url: https://www.phoronix.com/scan.php?page=news_item&amp;px=GCC-Git-Possibly-This-Weekend
---

## [4][timex: a test-friendly replacement for golang's time package](https://www.reddit.com/r/golang/comments/emqosy/timex_a_testfriendly_replacement_for_golangs_time/)
- url: https://github.com/cabify/timex
---

## [5][What is the correct approach to do table driven tests for http handler functions?](https://www.reddit.com/r/golang/comments/emok37/what_is_the_correct_approach_to_do_table_driven/)
- url: https://www.reddit.com/r/golang/comments/emok37/what_is_the_correct_approach_to_do_table_driven/
---
Hi,

I have experimented with table driven tests by some functions. But http handlers have different things. They usually don't return something and the way this handlers take input is also not like normal functions, as they take user input data through http request object.

I don't know how to organize table driven tests that will give input to handler functions and checks numerous outputs like status codes, returned data and some other things (with table driven tests).

&amp;#x200B;

I know, many of you guys faced this kind of problems and solved it too. So, please share with us what is approach to solve this problem?

I'm using Echo framework, but examples of any framework or even standard library's Handler Function will also be appreciated.

&amp;#x200B;

Thanks
## [6][I built a FOSS service to protect against link shorteners tracking us. It's GPLv3 and build completely with golang. Frontend is build with go templates. Help is very welcomed :D](https://www.reddit.com/r/golang/comments/em9pwi/i_built_a_foss_service_to_protect_against_link/)
- url: https://unshort.link/
---

## [7][Go 1.13.6 and Go 1.12.15 are Released](https://www.reddit.com/r/golang/comments/emh2un/go_1136_and_go_11215_are_released/)
- url: https://groups.google.com/forum/#!topic/golang-announce/RLFrcJ_FZZs
---

## [8][Attaching go services to a Python / Django app?](https://www.reddit.com/r/golang/comments/emnjf3/attaching_go_services_to_a_python_django_app/)
- url: https://www.reddit.com/r/golang/comments/emnjf3/attaching_go_services_to_a_python_django_app/
---
I don’t have too much experience with the more complicated parts of managing a web app— the most I have is experience configuring monoliths to do what I want.

Right now I’m trying to figure out a way to use a golang service for the real-time parts of my django app. Django has channels, but doesn’t quite fit my performance needs. How do two services like that (in this case, say a Django app and a websockets golang server) typically authenticate and communicate?
## [9][Executing html template in an infinite loop, new templates append to the bottom of the previous](https://www.reddit.com/r/golang/comments/emo7hp/executing_html_template_in_an_infinite_loop_new/)
- url: https://www.reddit.com/r/golang/comments/emo7hp/executing_html_template_in_an_infinite_loop_new/
---
Hey all. Racking my brain on this, and I'm wondering what I can be doing differently to solve this problem. I'm not putting the correct combination of words into Google to solve it either. But essentially I have an http handler that loads some data from an external API to fill in an html template. 

So lets just say I'm running a goroutine that is updating a struct of data named `apiAnswer` and a for loop like the following:
     
    for {
       indexTmpl.Execute(w, apiAnswer)
       time.Sleep(10 * time.Second)
    }

It works! Sort of... Except the new version of the page gets reloaded under the previous, and repeatedly does so.. Is there any way to remove previous copies of the template execution?
## [10][Development of Virtual Assistants Made Simpler with RASA](https://www.reddit.com/r/golang/comments/emnpen/development_of_virtual_assistants_made_simpler/)
- url: https://medium.com/@Oodles_AI/development-of-virtual-assistants-made-simpler-with-rasa-d27bdb18e909
---

