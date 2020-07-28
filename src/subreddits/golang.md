# golang
## [1][[Q&amp;A] io/fs draft design](https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/
---
I posted a draft design today for new file system interfaces for Go.

Video: https://golang.org/s/draft-iofs-video

Design: https://golang.org/s/draft-iofs-design

Let's do the Q&amp;A about the design here in Reddit. My hope is that the threading support will help keep questions and answers matched.

Please start a new top-level comment for each new question.

See also the related [Q&amp;A for the //go:embed draft design](https://golang.org/s/draft-embed-reddit).
## [2][Design Draft: First Class Fuzzing](https://www.reddit.com/r/golang/comments/hvpr96/design_draft_first_class_fuzzing/)
- url: https://go.googlesource.com/proposal/+/refs/heads/master/design/40307-fuzzing.md
---

## [3][A complete Golang and Nuxt(VueJS) boilerplate for your project with backend API, frontend, tests and CI/CD pipelines.](https://www.reddit.com/r/golang/comments/hyz5nt/a_complete_golang_and_nuxtvuejs_boilerplate_for/)
- url: https://gitlab.com/gadelkareem/skeleton
---

## [4][I made a video about Delve, the native Go debugger, and it's basics.](https://www.reddit.com/r/golang/comments/hz05yl/i_made_a_video_about_delve_the_native_go_debugger/)
- url: https://www.youtube.com/watch?v=r033vEzL6a4
---

## [5][Hitting an endpoint multiple times within a short period](https://www.reddit.com/r/golang/comments/hzemqz/hitting_an_endpoint_multiple_times_within_a_short/)
- url: https://www.reddit.com/r/golang/comments/hzemqz/hitting_an_endpoint_multiple_times_within_a_short/
---
Hello all,  
I'm trying to create an analysis of Github users, repos and their statistics using the [Golang Github API](https://pkg.go.dev/github.com/google/go-github/v31/github?tab=doc) . According to the library's docs, if I want to get the commit counts of user, I need to do the following:

1. Get all the repos of the specified user in a slice or a similar DS.  
2. Iterate through the slice and get individual repo data from the contributors service which gives me the list of contributors and their commits.  
3. I need to check for my particular user and then retrieve his/her commits or I can total them up and get the total commits of the repo.

Now here's the problem.  
If a user has 100 repos, I'm supposed to do 100 individual calls to the github API. The first three of four calls finish quickly (&lt;200ms) but the rest of the 95 calls take more than 1 sec. Which slows down my overall function in a very disgusting way.

So my question is, is there any way where I can achieve this without consecutively hitting the public API with such frequency? Has anyone encountered such a problem before? I tried optimising my Go code as far as I can. Any feedback on how to treat public APIs along with having a significant performance of an application would be very useful as well.

Thank you very much. :)
## [6][Basic CSV manipulation and Linear Algebra library - RocketC](https://www.reddit.com/r/golang/comments/hzeg55/basic_csv_manipulation_and_linear_algebra_library/)
- url: https://www.reddit.com/r/golang/comments/hzeg55/basic_csv_manipulation_and_linear_algebra_library/
---
I am working on a library which will provide functionality to program  machine learning algorithms from scratch. It’s in very early stage and  contains only basic functionalities, I am working on it. Please go  through it once and provide valuable feedback and suggestions.

[https://github.com/aryanmaurya1/RocketC](https://github.com/aryanmaurya1/RocketC)
## [7][GitHub - gin-boilerplate: The fastest way to deploy a restful api's with Gin Framework with a structured project that defaults to PostgreSQL database and JWT authentication middleware stored in Redis](https://www.reddit.com/r/golang/comments/hz2owl/github_ginboilerplate_the_fastest_way_to_deploy_a/)
- url: https://github.com/Massad/gin-boilerplate
---

## [8][My First go application, a very simple slack bot](https://www.reddit.com/r/golang/comments/hyrezs/my_first_go_application_a_very_simple_slack_bot/)
- url: https://www.reddit.com/r/golang/comments/hyrezs/my_first_go_application_a_very_simple_slack_bot/
---
[https://github.com/rimonmostafiz/frodobot](https://github.com/rimonmostafiz/frodobot)

I have been playing with GO for the last couple of days and wrote this slack bot. In my workplace, we use slack and we have a channel where we post our daily status update (kind of scrum). Team members sometimes forgot to post their status. So I wrote this bot to remind them to post their status.

1. The bot will start at every day 10:45 AM
2. Read all the messages of the particular channel from 6.00 AM to 10:45 AM
3. List out users who didn't post their status
4. Send a soft reminder message to the channel tagging those users

Please suggest to me how can I improve this project?   
Also If you like this project consider giving it a ⭐star ⭐. Thanks!
## [9][Please welcome, my first golang package](https://www.reddit.com/r/golang/comments/hzaxch/please_welcome_my_first_golang_package/)
- url: https://www.reddit.com/r/golang/comments/hzaxch/please_welcome_my_first_golang_package/
---
This is my very first golang package, please give me feedbacks 

[https://github.com/ramabmtr/asynctask](https://github.com/ramabmtr/asynctask)

this package aim to simplify the concurrent process on the goroutine that have response to process in main thread,

please visit and give me feedback how I can improve this package
## [10][Convert JSON to XLSX or XLSX to JSON with one HTTP request](https://www.reddit.com/r/golang/comments/hyvgfq/convert_json_to_xlsx_or_xlsx_to_json_with_one/)
- url: https://www.reddit.com/r/golang/comments/hyvgfq/convert_json_to_xlsx_or_xlsx_to_json_with_one/
---
After multiple weeks of working on this side-project, we finally achieve a first version of our Golang API.

[https://github.com/Los-Crackitos/Excelante](https://github.com/Los-Crackitos/Excelante)

This Golang based API aims to have an interaction between XLSX files and other type of output/input values.

This first version allow you to extract XLSX files into JSON or create XLSX files from JSON.

Currently only JSON is supported but we will implement other value type soon. We are also working on a deployed and usable free environment for single/demo conversions.

  
Feel free to star it if you think it deserve it. 

Thanks!
## [11][secp256k1 — A Go wrapper for libsecp256](https://www.reddit.com/r/golang/comments/hz7xwj/secp256k1_a_go_wrapper_for_libsecp256/)
- url: https://github.com/renproject/secp256k1
---

## [12][How can I validate a file upload before its uploaded or while its uploading? I would like to prevent myself from any file upload exploitation’s.](https://www.reddit.com/r/golang/comments/hz7iak/how_can_i_validate_a_file_upload_before_its/)
- url: https://www.reddit.com/r/golang/comments/hz7iak/how_can_i_validate_a_file_upload_before_its/
---

