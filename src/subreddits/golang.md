# golang
## [1][Just released v1.9.0 of glab (a GitLab Cli tool written in golang) with new cool features including WATCHING A RUNNING PIPELINE AND VIEWING A JOB'S TRACE/LOG DIRECTLY FROM YOUR CLI.](https://www.reddit.com/r/golang/comments/ib3rhe/just_released_v190_of_glab_a_gitlab_cli_tool/)
- url: https://github.com/profclems/glab
---

## [2][Task v3.0.0 is released! 🎉🎉🎉](https://www.reddit.com/r/golang/comments/ib4px3/task_v300_is_released/)
- url: https://github.com/go-task/task/releases/tag/v3.0.0
---

## [3][airscan: Go package to scan paper documents 📄 from a scanner 🖨️ via the network 🕸️ using the Apple AirScan (eSCL) protocol](https://www.reddit.com/r/golang/comments/iascmx/airscan_go_package_to_scan_paper_documents_from_a/)
- url: https://github.com/stapelberg/airscan
---

## [4][[Fun project] Released hselect - a cli mirror selector in Go](https://www.reddit.com/r/golang/comments/ibe5il/fun_project_released_hselect_a_cli_mirror/)
- url: https://www.reddit.com/r/golang/comments/ibe5il/fun_project_released_hselect_a_cli_mirror/
---
Worked on a mirror selection cli utility **hselect** that can select a reliable mirror (or any endpoint) from a group of mirror sites. It is based on my earlier work on **harmonic** dispatch algorithm. For trial, I am using it for selecting a reliable go module mirror site. Comments/suggestions are welcome.

[https://github.com/gptankit/hselect](https://github.com/gptankit/hselect)
## [5][Rocketc : CSV Manipulation and 2D Matrix library](https://www.reddit.com/r/golang/comments/ibe3qr/rocketc_csv_manipulation_and_2d_matrix_library/)
- url: https://www.reddit.com/r/golang/comments/ibe3qr/rocketc_csv_manipulation_and_2d_matrix_library/
---
Finally, I am happy to  announce that I have completed the development of RocketC, a simple  library written in Go for computations involving 2D Matrices with proper  documentation and testing.  
RocketC  


Github Link : [https://github.com/aryanmaurya1/rocketc](https://github.com/aryanmaurya1/rocketc)  


Fast, Simple and Lightweight library for CSV data manipulation and mathematical computation involving 2D Matrices.  
This  library can be used in developing simple Machine Learning models like  Linear Regression, Logistic Regression, etc from scratch.
## [6][Help: Encrypting and Decrypting with golang.org/x/crypto/openpgp](https://www.reddit.com/r/golang/comments/ibdtac/help_encrypting_and_decrypting_with/)
- url: https://www.reddit.com/r/golang/comments/ibdtac/help_encrypting_and_decrypting_with/
---
Hi, 

i have been trying to decrypt and encrypt ASCII armor PGP Messages that i receive and send to a API. 

I Wrote a function for decrypting messages and another one for encrypting messages. The Decrypting function Works and i can decrypt the API's messages. But My encrypt function Seems to be broken as Neither the API or my decrypt function can decrypt the Messages encrypted by my encrypt function. Decrypting the encrypted messages from my encrypt function always result in the following error: "Error reading message: openpgp: unsupported feature: unknown SymmetricallyEncrypted version"

i have uploaded a example to the Go Play Ground :  [https://play.golang.org/p/JhhZ4xWWYvn](https://play.golang.org/p/JhhZ4xWWYvn)

Any help is Appreciated.
## [7][How do you handle errors in background job?](https://www.reddit.com/r/golang/comments/ibd5pv/how_do_you_handle_errors_in_background_job/)
- url: https://www.reddit.com/r/golang/comments/ibd5pv/how_do_you_handle_errors_in_background_job/
---
Hey gophers,

I'm looking for the best practices and advice on how to handle errors in background jobs.

For example, I have a HTTP server and a background job running in a separate goroutine. The job is writing the current time in some file every second.

    func writeTimeToFile(ctx context.Context, filename string) {
        var timer = time.NewTimer(time.Second)
        for {
            select {
            case &lt;-timer.C:
                // 1. Open or create the file filename.
                // 2. Write the current time.
                // 3. Close the file.
                timer.Reset(time.Second)
            case &lt;-ctx.Done():
                return
            }
        }
    }

Errors here could happen when

1. the file is being opened or created.
2. the current time is being written into the file.
3. the file is closing.

*What do you suggest is the best solution to handle errors here? Should the errors be processed in place, or should errors be propagated elsewhere on upper levels? Which is the best place for taking decision on what to do with the error?*

I'm asking because errors have different meaning.

* Some error can be ignored without harm.
* Some errors are recoverable but requires some additional actions to be done prior to continue.
* Some errors requires job restart.
* Some errors requires to halt the whole program with diagnostics.
* etc etc.
## [8][Create versatile Microservices in Golang - part 1 of 10 part series](https://www.reddit.com/r/golang/comments/iap0gv/create_versatile_microservices_in_golang_part_1/)
- url: https://ewanvalentine.io/microservices-in-golang-part-1/
---

## [9][Transitioning from DevOps to Go developer?](https://www.reddit.com/r/golang/comments/iardul/transitioning_from_devops_to_go_developer/)
- url: https://www.reddit.com/r/golang/comments/iardul/transitioning_from_devops_to_go_developer/
---
Hi,

This might be one of the stranger posts, as many seem to wish to enter my field (DevOps). But truly, I enjoy development a lot more, specifically Go development. I don't feel the same level of excitement writing Jenkins pipelines, managing Kubernetes clusters or maintaining Terraform modules. Any time I get to develop a solution in Go/Python I feel genuine excitement. Sadly I can't opensource alot of my tooling either. It feels like I'm a glue, like the Star Wars midichlorians, but I wish to be the force user. 

Has anyone made this transition before? Any advice? Go roles are quite scarce around London as it is.
## [10][CodeRover - A low code platform for developing cross language and framework applications fast.](https://www.reddit.com/r/golang/comments/iba3z4/coderover_a_low_code_platform_for_developing/)
- url: https://www.reddit.com/r/golang/comments/iba3z4/coderover_a_low_code_platform_for_developing/
---
CodeRover is an open source low code platform with the following goals :

1. One platform for developing microservices with any language and related frameworks
2. Easy to use CLI and UI to facilitate rapid development of microservices.

[https://github.com/coderover-dev/coderover](https://github.com/coderover-dev/coderover)

&amp;#x200B;

The initial development is currently targeted with Golang/Gin stack for the MVP.

The project is at an early stage right now, looking forward to collaborate and open for contributors to take this project forward.

Please feel free to ask questions, provide suggestions.
