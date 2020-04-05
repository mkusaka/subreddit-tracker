# golang
## [1][A CLI Podcast player that I built after learning the basics of Go](https://www.reddit.com/r/golang/comments/fvaa5r/a_cli_podcast_player_that_i_built_after_learning/)
- url: https://github.com/goulinkh/podcast-cli
---

## [2][A job scheduler with WASM support I've been working on](https://www.reddit.com/r/golang/comments/fuzd1v/a_job_scheduler_with_wasm_support_ive_been/)
- url: https://github.com/suborbital/hive
---

## [3][Disposable chat room web app](https://www.reddit.com/r/golang/comments/fvcn0o/disposable_chat_room_web_app/)
- url: https://github.com/knadh/niltalk
---

## [4][Simple todo list with golang and vue](https://www.reddit.com/r/golang/comments/fv028j/simple_todo_list_with_golang_and_vue/)
- url: https://github.com/Lallassu/doit
---

## [5][Can I enforce a 'list' to be homogeneous?](https://www.reddit.com/r/golang/comments/fvdrdm/can_i_enforce_a_list_to_be_homogeneous/)
- url: https://www.reddit.com/r/golang/comments/fvdrdm/can_i_enforce_a_list_to_be_homogeneous/
---
I'm trying to learn the ropes of how Go manages its data structures without generics.

For instance, Go has a 'container/list' package that implements a Linked List. As I suspected, with the empty interface hack, the list is actually heterogeneous. I can tack on an integer, then a string, then any other struct with perfect utility.

But to me it's weird seeing as how I'm a big fan of a static strong type system. Using empty interfaces kind of throws away the type checking.

Is there any way to enforce that my particular list is homogeneous?
## [6][Dependency Injection in Go using Fx](https://www.reddit.com/r/golang/comments/fva01s/dependency_injection_in_go_using_fx/)
- url: https://pmihaylov.com/dependency-injection-go-fx/
---

## [7][An example project with sql.Tx and separation of concerns](https://www.reddit.com/r/golang/comments/fvc8mo/an_example_project_with_sqltx_and_separation_of/)
- url: https://github.com/ar3s3ru/go-transaction-example
---

## [8][Slacker, A Slack Bot Framework](https://www.reddit.com/r/golang/comments/futvtf/slacker_a_slack_bot_framework/)
- url: https://github.com/shomali11/slacker
---

## [9][Which framework to use for a multiplayer mobile game?](https://www.reddit.com/r/golang/comments/fuqwqm/which_framework_to_use_for_a_multiplayer_mobile/)
- url: https://www.reddit.com/r/golang/comments/fuqwqm/which_framework_to_use_for_a_multiplayer_mobile/
---
Hi, I have to develop the back-end of a multiplayer mobile game. I specify that the real time part of the game will be managed with a separate service, so in Go I should just develop the services (REST API) that interface with the database (for example user registration, writing matches completed on db etc.), In particular we will use Mongodb.

So far I've simply used Gorilla/Mux and net/http to create APIs. I wanted to ask if you recommend a framework (e.g. Gin Gonic, Chi etc.) that allows me to better structure the code and allows me to develop faster than using only Gorilla/mux and net/http.

Thanks in advance.
## [10][Is there a way to do a global fuzzy-search using vim-go, just like cmd-shift-F/double shift in Goland IDE?](https://www.reddit.com/r/golang/comments/fv90qo/is_there_a_way_to_do_a_global_fuzzysearch_using/)
- url: https://www.reddit.com/r/golang/comments/fv90qo/is_there_a_way_to_do_a_global_fuzzysearch_using/
---
I'm trying to move away from proprietary Goland IDE, and these commands are one of the commands I use extensively. Is there a way to replicate that behavior in Vim using vim-go, or any other Vim plugin? I couldn't find anything like that by doing a web search or vim-go helptext search.

&amp;#x200B;

Thanks so much for your time!
