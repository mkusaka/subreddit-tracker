# golang
## [1][Nine-year-old bug in the Go standard library enables DoS](https://www.reddit.com/r/golang/comments/i59jmo/nineyearold_bug_in_the_go_standard_library/)
- url: https://github.com/ethereum/public-attacknets/issues/12
---

## [2][CLI tool for viewing/writing to and from Kafka, RabbitMQ and more](https://www.reddit.com/r/golang/comments/i54dnl/cli_tool_for_viewingwriting_to_and_from_kafka/)
- url: https://github.com/batchcorp/plumber
---

## [3][ReadUvarint unlimited input (CVE-2020-16845) - Go 1.14.7 and Go 1.13.15 security update](https://www.reddit.com/r/golang/comments/i4wyfk/readuvarint_unlimited_input_cve202016845_go_1147/)
- url: https://groups.google.com/u/1/g/golang-announce/c/NyPIaucMgXo/m/GdsyQP6QAAAJ?pli=1
---

## [4][Create a dummy REST API from a json file with zero coding in seconds](https://www.reddit.com/r/golang/comments/i4t3rd/create_a_dummy_rest_api_from_a_json_file_with/)
- url: https://www.reddit.com/r/golang/comments/i4t3rd/create_a_dummy_rest_api_from_a_json_file_with/
---
Hello guys, long time lurker on this sub but never posted before. 

I created **json-server**, a CLI tool to create a dummy REST API from a provided json file with zero coding in seconds. For each provided resource 6 full functional endpoints are created (GET x2, POST, PUT, PATCH, DELETE), that you can use right away. 

Inspired by the javascript package [json-server](https://github.com/typicode/json-server) that's where the name comes from. 

The next step is to create the first release, which will include binary files for Windows, Linux and macOS. Any comments/suggestions are really welcomed. You can find more info at README file.

[https://github.com/chanioxaris/json-server](https://github.com/chanioxaris/json-server)
## [5][gocv recognise ellipse](https://www.reddit.com/r/golang/comments/i5bmts/gocv_recognise_ellipse/)
- url: https://i.redd.it/lipmzj5m8kf51.jpg
---

## [6][How to programmatically get the size of a struct, incl its data?](https://www.reddit.com/r/golang/comments/i5bjie/how_to_programmatically_get_the_size_of_a_struct/)
- url: https://www.reddit.com/r/golang/comments/i5bjie/how_to_programmatically_get_the_size_of_a_struct/
---
I have been playing around with `unsafe.Sizeof` and `binary.Size` but neither are working for this use case. I have a struct as follows, and I am looking to find out the memory impact of it as the program continues. The issue is that structX can vary in size so `binary.Size` doesnt like it.

    type struct sample {
        data [] structX
    }

I have tried using pprof but the memory usage used varies too much, weird I know. Running a benchmark once gives 1GB usage, running it again gives 500MB, its too inconsistent. I tried using memStats but there are so many go routines running that I cant pin down the memory used by this one struct. The struct is allocated on the heap and has a lifespan of most of the running of the program. Any recommendations?
## [7][How to proxy the keystrokes from one keyboard to another?](https://www.reddit.com/r/golang/comments/i5ck1a/how_to_proxy_the_keystrokes_from_one_keyboard_to/)
- url: https://www.reddit.com/r/golang/comments/i5ck1a/how_to_proxy_the_keystrokes_from_one_keyboard_to/
---
I want to proxy the keystrokes from one keyboard to another keyboard, how can I do this with Linux? Any suggestion?
## [8][AWS Lambda in GoLang Guide](https://www.reddit.com/r/golang/comments/i59mlm/aws_lambda_in_golang_guide/)
- url: https://www.softkraft.co/aws-lambda-in-golang/
---

## [9][Go: Introduction to the Escape Analysis](https://www.reddit.com/r/golang/comments/i4w3zd/go_introduction_to_the_escape_analysis/)
- url: https://medium.com/a-journey-with-go/go-introduction-to-the-escape-analysis-f7610174e890
---

## [10][What is the best error handle in web applicaiton.](https://www.reddit.com/r/golang/comments/i57ubx/what_is_the_best_error_handle_in_web_applicaiton/)
- url: https://www.reddit.com/r/golang/comments/i57ubx/what_is_the_best_error_handle_in_web_applicaiton/
---
Hi everyone. How do you handle the error in web application?. In my case. I often have to pass errors from the `repository` layer to the `handler` layer. This error is an error about operate database. It should not display to users. So I must return a 500 in handler if error not equal nil.
```go
/repositories/user_repository.go. send error to the service layer.
tx := r.conn.Begin()
if err := tx.Error; err != nil {
	return userCreated, err
}
if err := tx.Create(&amp;user).Scan(&amp;userCreated).Error; err != nil {
	tx.Rollback()
	return userCreated, err
}
tx.Commit()
return userCreated, nil

/services/user_service.go receive the error 
_, err := u.userRepository.CreateUser(user)
if err != nil {
	return err
}
return nil

/handles/user_handler.go
if err != nil {
	ginresp.InternalError(c, "Create failed", nil, err)
}
```
If I panic it in service layer. then the recovery middlewares will handle it. Does this way is best?
