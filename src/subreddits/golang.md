# golang
## [1][ESP32 and ESP8266 support in TinyGo](https://www.reddit.com/r/golang/comments/j027qe/esp32_and_esp8266_support_in_tinygo/)
- url: https://aykevl.nl/2020/09/tinygo-esp32
---

## [2][GoLand 2020.3 Early Access Program Is Starting! Make goroutines dumps, initial support for table tests, upgrades for Testify support and code completion, UI improvements, and more!](https://www.reddit.com/r/golang/comments/izo4nu/goland_20203_early_access_program_is_starting/)
- url: https://blog.jetbrains.com/go/2020/09/25/goland-2020-3-eap/
---

## [3][Up to which struct size is passing by value quicker then passing by pointer?](https://www.reddit.com/r/golang/comments/j05usq/up_to_which_struct_size_is_passing_by_value/)
- url: https://www.reddit.com/r/golang/comments/j05usq/up_to_which_struct_size_is_passing_by_value/
---
I created some benchmarks which suggest that on my machine a struct consisting of up to 9 int 64 is faster to be passed as value. From 10 fields and more passing by pointer is faster.

[https://github.com/komkom/value-vs-ptr](https://github.com/komkom/value-vs-ptr)

Are these tests correct? And if they are correct why would it be that on both of my machines the breaking point is 9 fields?
## [4][spotrip: Open source Spotify ripper in Go - download an artists entire discography, individual albums, or tracks; works with free and premium accounts](https://www.reddit.com/r/golang/comments/izt3p7/spotrip_open_source_spotify_ripper_in_go_download/)
- url: https://github.com/chris124567/spotrip
---

## [5][I just finished and released v1.0 of my programming language, written entirely in Go! I have no idea how to properly write a language but I gave it my best shot](https://www.reddit.com/r/golang/comments/izgz5w/i_just_finished_and_released_v10_of_my/)
- url: https://github.com/odddollar/Leafscript
---

## [6][Any RunDeck like open source automation tools written in Go?](https://www.reddit.com/r/golang/comments/j05c7v/any_rundeck_like_open_source_automation_tools/)
- url: https://www.reddit.com/r/golang/comments/j05c7v/any_rundeck_like_open_source_automation_tools/
---
Are there any [RunDeck](https://www.rundeck.com/) like automation tools written in Go?
## [7][I am relatively experienced now programming but I am at a complete loss for what i want to program.](https://www.reddit.com/r/golang/comments/izspm3/i_am_relatively_experienced_now_programming_but_i/)
- url: https://www.reddit.com/r/golang/comments/izspm3/i_am_relatively_experienced_now_programming_but_i/
---
I had an idea about some kind of basic real time 3d rendering as math is a strong suit of mine but i couldn’t find any good resources for that. And all the project ideas i have seen online are meant for beginners. I feel like i’ve hit writers block but for programming.
## [8][Developing price and currency handling for Go](https://www.reddit.com/r/golang/comments/izk4kz/developing_price_and_currency_handling_for_go/)
- url: https://bojanz.github.io/price-currency-handling-go/
---

## [9][Using go-sqlmock and cannot get this to work..](https://www.reddit.com/r/golang/comments/j01eez/using_gosqlmock_and_cannot_get_this_to_work/)
- url: https://www.reddit.com/r/golang/comments/j01eez/using_gosqlmock_and_cannot_get_this_to_work/
---
I've setup my mockDb and pass in the db into the function to call it.

    t.mockDb.ExpectExec("UPDATE customers SET customer_id = $1").WithArgs(sqlmock.AnyArg()).WillReturnError(nil)
    
    err := t.store.UpdateCustomer(1)
    t.Equal(nil, err)

I am getting the following error:

    postgres exec error: ExecQuery: could not match actual sql: "UPDATE customers SET customer_id = $1"
## [10][Get a free copy of "For the Love of Go: Fundamentals" by John Arundel](https://www.reddit.com/r/golang/comments/izlmes/get_a_free_copy_of_for_the_love_of_go/)
- url: https://twitter.com/goinggodotnet/status/1309493600372379652?s=20
---

