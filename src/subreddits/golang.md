# golang
## [1][Breakout: browser game written on pure Go and compiled into WASM. No JS added.](https://www.reddit.com/r/golang/comments/jb5kn7/breakout_browser_game_written_on_pure_go_and/)
- url: https://gweb.orsinium.dev/breakout/
---

## [2][Go 1.15.3 and Go 1.14.10 are released](https://www.reddit.com/r/golang/comments/jb9wpq/go_1153_and_go_11410_are_released/)
- url: https://golang.org/doc/devel/release.html#go1.15.minor
---

## [3][[ANN] Fast JSON lexer with zero heap allocations](https://www.reddit.com/r/golang/comments/jbmsx3/ann_fast_json_lexer_with_zero_heap_allocations/)
- url: https://www.reddit.com/r/golang/comments/jbmsx3/ann_fast_json_lexer_with_zero_heap_allocations/
---
I had a requirement for a fast JSON lexer (lexer only) and did not found one that matches my needs. So I wrote one - which at least is a bit faster than the default encoding/json/Decoder.Token() and garbage collector friendly. Maybe someone requires this as well \[1\].

    $ typex github.com/dtgorski/jsonlex
    └── github.com
        └── dtgorski
            └── jsonlex
                ├── Lexer struct {}
                ├── Token uint8
                └── Yield func(token jsonlex.Token, load []byte, pos uint)

[https://github.com/dtgorski/jsonlex](https://github.com/dtgorski/jsonlex)

regards dtg
## [4][Free offline solution to convert PDFs into audiobooks -](https://www.reddit.com/r/golang/comments/jbmmga/free_offline_solution_to_convert_pdfs_into/)
- url: https://github.com/Harry-027/go-audio
---

## [5][Casbin: An authorization library that supports access control models like ACL, RBAC, ABAC in Golang](https://www.reddit.com/r/golang/comments/jberp8/casbin_an_authorization_library_that_supports/)
- url: https://github.com/casbin/casbin
---

## [6][Is Go a good first server language to learn?](https://www.reddit.com/r/golang/comments/jba2w6/is_go_a_good_first_server_language_to_learn/)
- url: https://www.reddit.com/r/golang/comments/jba2w6/is_go_a_good_first_server_language_to_learn/
---
Title says it, but is Go a good first language to learn for creating writing server code (APIs, either REST or GraphQL), or would I be screwing myself by attempting that first? I'm attempting to make a full-stack application (client, server). For my client, I would be using some JavaScript framework (probably React), but I've been up in the air about what to use for the server code. The typical recommendations I see is to go for Node.js first or Python, but Go caught my attention since I heard that its a language specifically made to be good to use for server code.

The only problem I see with that is if I get too comfortable with it. For a bit of contetx, most of my development experience is in Java but I've recently learned Kotlin which I enjoy, but I wasn't sure if learning Go first would be like learning Kotlin first and then having to use Java later. Doable, but would miss a lot of the niceties/faster productivity that came with it.

I'm also wondering how much companies are starting to adopt it lately compared to those other two, particularly in SV.
## [7][Structuring and testing HTTP handlers in Go](https://www.reddit.com/r/golang/comments/jbk959/structuring_and_testing_http_handlers_in_go/)
- url: https://www.maragu.dk/blog/structuring-and-testing-http-handlers-in-go/
---

## [8][How To Correctly Validate Passwords, go-password-validator released](https://www.reddit.com/r/golang/comments/jbmwbb/how_to_correctly_validate_passwords/)
- url: https://qvault.io/2020/10/15/how-to-correctly-validate-passwords-most-websites-do-it-wrong/
---

## [9][Self recursive struct](https://www.reddit.com/r/golang/comments/jbmur6/self_recursive_struct/)
- url: https://www.reddit.com/r/golang/comments/jbmur6/self_recursive_struct/
---
I am looking in a way to recuse into a struct and return the full tree. e.g  


    type Item stuct{
    ID int64
    Subitems []Item
    }

So the return i am looking for is:  
Item.Subitems.Subitems.Subitems..... filled with all the attached stuctures.

Although a simple recursion works, i cannot get it to fill anything other than the first level.

Thank you.
## [10][Go- Web API Framework](https://www.reddit.com/r/golang/comments/jblnwh/go_web_api_framework/)
- url: https://www.reddit.com/r/golang/comments/jblnwh/go_web_api_framework/
---
What's the best minimalist API web framework for Golang.
