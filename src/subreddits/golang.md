# golang
## [1][So, we made MTProto full native implementation on pure go. This is protocol used by Telegram messenger](https://www.reddit.com/r/golang/comments/j5hffz/so_we_made_mtproto_full_native_implementation_on/)
- url: https://github.com/xelaj/mtproto
---

## [2][Hash Server 9000](https://www.reddit.com/r/golang/comments/j5c7eq/hash_server_9000/)
- url: https://www.reddit.com/r/golang/comments/j5c7eq/hash_server_9000/
---
I took an internal dev challenge at work a while back because I wanted to learn more Go since I have an OO background and was not familiar with the procedural way of thinking.  The goal was to build a simple [non-persistent password hashing service](https://github.com/JudeQuintana/hash_server) using only the standard library while demonstrating the use of concurrency with the endpoint requirements below. Reading the [blue and white book](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440) was definitely key to gaining deeper insight. It was super challenging and I tested as much as I could. Although, I still need to figure out benchmark tests. There are probably areas for improvement but it was fun and I learned a ton so I thought I'd share. I don't work with Go as much as I'd like to but building this service really helped me get a fundamental understanding and now it's much easier to read Go code. It's amazing how powerful the standard library is so that I can serve them hashes HOT!

## Endpoints

1. `POST /hash` \- Hash and encode a password string. The request must contain a `password` parameter. Returns the `id` of Base64 encoded string of the password that's been hashed with SHA512 with a 5 second delay to simulate asynchronous processing. Example: `curl --data "password=angryMonkey" http://localhost:8080/hash`
2. `GET /hash/:id` \- Retrieve a generated hash with the `id` (after approximately 5 seconds), otherwise you will receive error `id not found`.  Example: `curl http://localhost:8080/hash/1` will return: `ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==`.
3. `GET /stats` \- Statistics endpoint. Returns a JSON object with the `total` count of the number of password hash requests made to the server so far and the `average` time in milliseconds it has taken to process all of the requests.  Example: `curl http://localhost:8080/stats`, will return: `{"total": 1, "average": 123}`
4. `GET /shutdown` \- Graceful shutdown. When a request is made to `/shutdown` the server will reject new requests and wait for any pending/in-flight work to finish before exiting.
## [3][Go Regular Expressions](https://www.reddit.com/r/golang/comments/j5hn5o/go_regular_expressions/)
- url: https://codesalad.dev/blog/go-regular-expressions-5
---

## [4][Go-edlib: Golang edit distance and string comparison library fully compatible with Unicode (Levenshtein, LCS, Hamming, Damerau-Levenshtein, Jaro, Cosine, etc...)](https://www.reddit.com/r/golang/comments/j5hvsw/goedlib_golang_edit_distance_and_string/)
- url: https://github.com/hbollon/go-edlib
---

## [5][gimmeasearx - Configurable, JavaScript-less Neocities alternative, written in Go](https://www.reddit.com/r/golang/comments/j51x9l/gimmeasearx_configurable_javascriptless_neocities/)
- url: https://github.com/demostanis/gimmeasearx
---

## [6][All the go help ⛑️ documentation at one place 🔥️](https://www.reddit.com/r/golang/comments/j5ioi9/all_the_go_help_documentation_at_one_place/)
- url: https://golang.org/src/cmd/go/internal/help/helpdoc.go
---

## [7][Kalbi - Golang SIP Framework](https://www.reddit.com/r/golang/comments/j59jkw/kalbi_golang_sip_framework/)
- url: https://www.reddit.com/r/golang/comments/j59jkw/kalbi_golang_sip_framework/
---
Hey Guys

Looking for developers to help contribute to this project !!

its a SIP (Session Initiation Protocol) Framework to help build SIP application in golang 

[https://github.com/KalbiProject/Kalbi](https://github.com/KalbiProject/Kalbi)

&amp;#x200B;

EDIT - project has been updated with some issues
## [8][How to implement concurrency in Go](https://www.reddit.com/r/golang/comments/j5i468/how_to_implement_concurrency_in_go/)
- url: https://blog.umesh.wtf/how-to-implement-concurrency-in-go
---

## [9][Feature-complete framework or libs for rapid web dev with golang?](https://www.reddit.com/r/golang/comments/j54275/featurecomplete_framework_or_libs_for_rapid_web/)
- url: https://www.reddit.com/r/golang/comments/j54275/featurecomplete_framework_or_libs_for_rapid_web/
---
I'm a developper, I know a fair share on how to implement sessions, templating and REST APIs in Java, but I have touched Golang for some embedded project and I'd like to code my next (web) project with it.

I'm not quite up to speed on which frameworks to use, or if there are "standard" libraries that are state of the art in Go. What I know though is that I would like to have a maximum of "batteries included" because I'd rather focus on the business aspects of the code, and leave the rest to a trustable framework.

Do you guys know about something that could tie all these:

- session/authentication management  
- some SQLite abstraction (ideally with schema migrations)  
- Maybe a combined frontend / backend development experience ?
- Testable easily?

I've heard of Gin and Revel on the web side, and Gorm on the SQLite side, but I have no idea if I would footgun myself with these.

Anybody has experience with a combo that works?
Thank you to all people willing to discuss this with me!
## [10][slack-go/slack v0.7.0 released!](https://www.reddit.com/r/golang/comments/j52aa4/slackgoslack_v070_released/)
- url: https://www.reddit.com/r/golang/comments/j52aa4/slackgoslack_v070_released/
---
See also: [https://github.com/slack-go/slack/releases/tag/v0.7.0](https://github.com/slack-go/slack/releases/tag/v0.7.0)
