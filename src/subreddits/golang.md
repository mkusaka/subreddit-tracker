# golang
## [1][Simple embedding of static files in your binary](https://www.reddit.com/r/golang/comments/hdiod7/simple_embedding_of_static_files_in_your_binary/)
- url: https://www.reddit.com/r/golang/comments/hdiod7/simple_embedding_of_static_files_in_your_binary/
---
I found the existing tools for resource embedding a little complicated, lacking features, hard to use and having low test coverage. 

So I wrote
[binclude](https://github.com/lu4p/binclude). 

I would appreciate getting some feedback on features you miss and suggestions for improvements in documentation and code.

Features:

- high test coverage 96%+
- add files directly in your code with just  `binclude.Include("./assets")` this includes the assets folder with all files and subdirectories
- files have the same paths as on your host os
- development mode for loading files from disk instead
- ioutil like functions
- ast for typsafe parsing and code generation
- the generated filesystem implements the http.Filesystem interface for standard conformity
- supports execution of included executables directly (wrapper for os/exec)
## [2][(Bad?) Experiments with Go generics](https://www.reddit.com/r/golang/comments/hdlu9g/bad_experiments_with_go_generics/)
- url: https://brbe.net/posts/experiments-with-go-generics/
---

## [3][PIN Verification API of a banking application - which approach is preferred?](https://www.reddit.com/r/golang/comments/hdptnu/pin_verification_api_of_a_banking_application/)
- url: https://www.reddit.com/r/golang/comments/hdptnu/pin_verification_api_of_a_banking_application/
---
**Option 1:**

`GET /api/v1/pin?code=123456`

**Option 2:**

`POST /api/v1/pin` with request body `{"code":"123456"}`

This is why I am asking your suggestion:

* The issue with first option is, sending sensitive data like PIN code in query string doesn't  seem to be right approach. Again, as sending through *https*, the code will not be visible to middleman (if any).
* And, the problem with second option is, "PIN Verification" is conceptually not a "POST" as it's not *creating resource* in REST terms.

Please let me which one will you prefer and why. Also, if want to suggest any alternative approach, totally welcome :)
## [4][go get -u &lt;github url&gt; not doing anything](https://www.reddit.com/r/golang/comments/hdsclr/go_get_u_github_url_not_doing_anything/)
- url: https://www.reddit.com/r/golang/comments/hdsclr/go_get_u_github_url_not_doing_anything/
---
This command would not throw any output, literally nothing. When I enter the program's name it will say no command found.

Why isn't it working?
## [5][Library for conversion and Forex of 6100+ Fiat &amp; Crypto currency with no API key](https://www.reddit.com/r/golang/comments/hds3td/library_for_conversion_and_forex_of_6100_fiat/)
- url: https://github.com/asvvvad/exchange
---

## [6][gRPC over WebRTC](https://www.reddit.com/r/golang/comments/hd6hsb/grpc_over_webrtc/)
- url: https://github.com/jsmouret/grpc-over-webrtc
---

## [7][Type safe query builder and struct mapper](https://www.reddit.com/r/golang/comments/hdq5dt/type_safe_query_builder_and_struct_mapper/)
- url: https://www.reddit.com/r/golang/comments/hdq5dt/type_safe_query_builder_and_struct_mapper/
---
I've been working on this type safe query builder for the past month, and now I would like feedback on it. Especially for people who write SQL heavily, does it cover the queries that you normally write?

Here's the package: [https://github.com/bokwoon95/go-structured-query](https://github.com/bokwoon95/go-structured-query)

- I wrote this initially out of frustration due to null handling, I didn't want to pepper my structs with NullStrings and NullInt64s. More often than not the zero value suited me fine.
- Most query builders fall back on strings, this one uses jOOQ-like columns generated from your database. So you gain some type-safety knowing that the columns you're putting into the query builder actually exist and you didn't mistype something.
- The most unique aspect is the struct mapping, you use a callback function to map columns to fields which doubles as the SELECT clause so you don't have to specify the columns twice:


    u := tables.USERS().As("u")

    var user User
    var users []User

    // SELECT u.user_id, u.name, u.email
    // FROM public.users AS u
    // WHERE u.name = 'bob'
    err := From(u).
        Where(u.NAME.EqString("bob")).
        Selectx(func(row *sq.Row) {
            user.UserID = row.Int(u.USER_ID)
            user.Name = row.String(u.NAME)
            user.Email = row.String(u.EMAIL)
        }, func() {
            users = append(users, user)
        }).
        Fetch(DB)


The code above is equivalent to running the SQL query, scanning each column into the corresponding user struct field and aggregating the user into the slice of users with rows.Next().

The code formatting seems borked, here is a link [playground](https://play.golang.org/p/LXTiWaJQOhI).
## [8][Youtube: Go Syntax with Bill Kennedy](https://www.reddit.com/r/golang/comments/hdfiiz/youtube_go_syntax_with_bill_kennedy/)
- url: https://www.youtube.com/playlist?list=PLADD_vxzPcZATO4tDU_nHApxTEJyxskiS
---

## [9][Manber&amp;Myers's Suffix Array implemented in Go.](https://www.reddit.com/r/golang/comments/hdpw79/manbermyerss_suffix_array_implemented_in_go/)
- url: https://github.com/d-tsuji/suffixarray
---

## [10][How productive is Golang?](https://www.reddit.com/r/golang/comments/hdq7i5/how_productive_is_golang/)
- url: https://www.reddit.com/r/golang/comments/hdq7i5/how_productive_is_golang/
---
Hi, I'm looking for a new web development stack and I had to choose between Elixir and Golang. For the time begin, I can only choose one language and pick a framework. My background is in PHP with laravel.

How productive it is compare to other language?

EDIT: I'm looking into Beego as it's the closest thing to Laravel and provide ORM out of the box. My main focus is realtime web application. Back in Laravel, I would use something like [socket.io](https://socket.io) \+ redis where Laravel broadcast event to redis and had [socket.io](https://socket.io) handle the websocket between server + client.
