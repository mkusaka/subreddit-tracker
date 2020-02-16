# golang
## [1][A friend asked me how to write a resilient service worker in Go . I figured I should share it to anyone interested.](https://www.reddit.com/r/golang/comments/f4ov4d/a_friend_asked_me_how_to_write_a_resilient/)
- url: https://gist.github.com/System-Glitch/301e95975a2645b8ea57c47b0c7cfef4
---

## [2][esbuild - JavaScript bundler and minifier](https://www.reddit.com/r/golang/comments/f4hnt9/esbuild_javascript_bundler_and_minifier/)
- url: https://github.com/evanw/esbuild/
---

## [3][Encrypting data inside a png image file](https://www.reddit.com/r/golang/comments/f4gh02/encrypting_data_inside_a_png_image_file/)
- url: https://github.com/Trivernis/cryptpng/
---

## [4][GORM v2 is ongoing, suggestions welcome ~](https://www.reddit.com/r/golang/comments/f4jv2i/gorm_v2_is_ongoing_suggestions_welcome/)
- url: https://github.com/jinzhu/gorm/issues/2886
---

## [5][Total 2 years of backend experience, new to Go, but want to make a switch in Go. What is expected?](https://www.reddit.com/r/golang/comments/f4d428/total_2_years_of_backend_experience_new_to_go_but/)
- url: https://www.reddit.com/r/golang/comments/f4d428/total_2_years_of_backend_experience_new_to_go_but/
---
Hey there!

I have a total of 2 years of backend experience and am pretty new to Go. What is expected by an "experienced newbie"? 

A little background on me,

The organization I work for has a pretty great development framework written on top of .net, due to which we mostly had to configure the software and write minimal code. To be honest, I wrote very less code in the past two years, which is something that I would like to change.

Apart from my work, I have worked on a few hobby projects though, all written in go, namely,

1. Web Scrapper
2. APIs for Expense Tracker
3. TCP Server using the net package
4. Adaptive Video Streaming Server

Keeping this in mind, what else I should know about backend and development in general.

Any advice would be helpful. 

Thanks in Advance!
## [6][How I structure cloud functions in Go](https://www.reddit.com/r/golang/comments/f4q025/how_i_structure_cloud_functions_in_go/)
- url: https://www.reddit.com/r/golang/comments/f4q025/how_i_structure_cloud_functions_in_go/
---
[https://medium.com/@rizwaniqbal/how-i-structure-cloud-functions-in-go-61e151b278ac](https://medium.com/@rizwaniqbal/how-i-structure-cloud-functions-in-go-61e151b278ac)
## [7][zeebo/blake3: Pure Go implementation of BLAKE3 with AVX2 and SSE4.1 acceleration](https://www.reddit.com/r/golang/comments/f4pgok/zeeboblake3_pure_go_implementation_of_blake3_with/)
- url: https://github.com/zeebo/blake3
---

## [8][Learn How to use JSON in Golang - CodeSource.io](https://www.reddit.com/r/golang/comments/f4pav0/learn_how_to_use_json_in_golang_codesourceio/)
- url: https://codesource.io/learn-how-to-use-json-in-golang/
---

## [9][What is the right way to connect to a database?](https://www.reddit.com/r/golang/comments/f4k7y8/what_is_the_right_way_to_connect_to_a_database/)
- url: https://www.reddit.com/r/golang/comments/f4k7y8/what_is_the_right_way_to_connect_to_a_database/
---
Hey everyone,

I am new to Go, and I keep seeing people connect to databases by writing a file that contains some code similar to this:

    func ConnectSQL() *sql.DB {
    	db, err := sql.Open("mysql", "")
    	if err != nil {
    		fmt.Print(err)
    	}
    	return db
    }

then through out their models they will do something like this:

    var db * sql.DB
    
    func init() {
    	db = api.ConnectSQL()
    } 

is this the correct way or a bad way to connect to a database? It just seems strange to me while learning this language. Any advice here would be helpful!

this would open a DB connection on every API call and that seems strange too me.
## [10][Moving Towards DDD in Go](https://www.reddit.com/r/golang/comments/f480oe/moving_towards_ddd_in_go/)
- url: https://www.calhoun.io/moving-towards-domain-driven-design-in-go/
---

