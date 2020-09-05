# golang
## [1][I wrote a piece about (lightweight) design patterns in Go, hope anyone finds it useful](https://www.reddit.com/r/golang/comments/iml937/i_wrote_a_piece_about_lightweight_design_patterns/)
- url: https://link.medium.com/tGADr7pOv9
---

## [2][golang performance for lambda cold restart](https://www.reddit.com/r/golang/comments/imqx3p/golang_performance_for_lambda_cold_restart/)
- url: https://www.reddit.com/r/golang/comments/imqx3p/golang_performance_for_lambda_cold_restart/
---
hey all, I was surprised to see golang lag python and js for lambda coldstart performance. [https://medium.com/the-theam-journey/benchmarking-aws-lambda-runtimes-in-2019-part-i-b1ee459a293d](https://medium.com/the-theam-journey/benchmarking-aws-lambda-runtimes-in-2019-part-i-b1ee459a293d)

seems like the binary size is the issue? i thought golang's single binary and speed would trump  interpreted languages. not sure if anyone has explored this topic and has any topics on optimizing golang in a lambda environment.

it's not an issue for me yet, was just curious.
## [3][Series on integrating a Go application with ELK](https://www.reddit.com/r/golang/comments/in0el8/series_on_integrating_a_go_application_with_elk/)
- url: https://pmihaylov.com/series-integrating-go-with-elk/
---

## [4][I made rainbow progress bar](https://www.reddit.com/r/golang/comments/im9t1v/i_made_rainbow_progress_bar/)
- url: https://v.redd.it/t5f4qk4ub2l51
---

## [5][Unit Test Logs Only For Failed Inputs](https://www.reddit.com/r/golang/comments/imxl3c/unit_test_logs_only_for_failed_inputs/)
- url: https://kaue.me/wealth/career/unit-test-logs-only-for-failed-inputs/
---

## [6][Infinite loop: read file, do something with lines, wait and restart - What am I doing wrong?](https://www.reddit.com/r/golang/comments/imunh1/infinite_loop_read_file_do_something_with_lines/)
- url: https://www.reddit.com/r/golang/comments/imunh1/infinite_loop_read_file_do_something_with_lines/
---
Description pretty much says it all. I have replicated the behaviour with this simple code snippet:

    func main() {
    	files, _ := os.Open("hosts.txt")
    	defer files.Close()
    	scanner := bufio.NewScanner(files)
    	scanner.Split(bufio.ScanLines)
    	for {
    		for scanner.Scan() {
    			fmt.Printf("Host: %s\n", scanner.Text())
    		}
    		time.Sleep(3 * time.Second)
    	}
    }

Code will run, but will never run second, third and so on times, and I don't understand what is wrong flow-wise. Am I using bufio.Scan() wrong?
## [7][An Alternative Approach to BDD in Go](https://www.reddit.com/r/golang/comments/iml318/an_alternative_approach_to_bdd_in_go/)
- url: https://medium.com/@elliotchance/an-alternative-approach-to-bdd-in-go-776bbbc24be9
---

## [8][Implement struct no copy in GoLang](https://www.reddit.com/r/golang/comments/imu58q/implement_struct_no_copy_in_golang/)
- url: https://www.pixelstech.net/article/1599275392-Implement-struct-no-copy-in-GoLang
---

## [9][Best practice on how to build a scalable Shell](https://www.reddit.com/r/golang/comments/imk3f6/best_practice_on_how_to_build_a_scalable_shell/)
- url: https://www.reddit.com/r/golang/comments/imk3f6/best_practice_on_how_to_build_a_scalable_shell/
---
Hey Guys  


I'm new to Go and want to build my own Shell as the first Project.  
Now I'm wondering whats the best way to structure such code I had the following things in mind:  


If you implement custom shell commands like:  "cd" etc. write those as Plugin as Goshell does.  
But what I would like to have is just a simple file called "cd" in the folder modules.  
But how can I then import all those Modules into the main file so I don't have to maintain a list of commands?  


Any tips on how to learn those best practices?  


Greets
## [10][GORM 2.0 Release Note | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://www.reddit.com/r/golang/comments/imjvqw/gorm_20_release_note_gorm_the_fantastic_orm/)
- url: https://gorm.io/docs/v2_release_note.html
---

