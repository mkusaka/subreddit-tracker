# golang
## [1][Speeding up JSON processing in go](https://www.reddit.com/r/golang/comments/gn45jq/speeding_up_json_processing_in_go/)
- url: https://klotzandrew.com/blog/speeding-up-json-processing-in-go
---

## [2][A complete Terraform setup of a serverless application on Google Cloud Run and Firebase](https://www.reddit.com/r/golang/comments/gnar3n/a_complete_terraform_setup_of_a_serverless/)
- url: https://threedots.tech/post/complete-setup-of-serverless-application/
---

## [3][I built a reddit image downloader in Go. This is my first project using Golang](https://www.reddit.com/r/golang/comments/gn74r2/i_built_a_reddit_image_downloader_in_go_this_is/)
- url: https://www.reddit.com/r/golang/comments/gn74r2/i_built_a_reddit_image_downloader_in_go_this_is/
---
Please review my code.

[https://github.com/axiiomatic/reddit-downloader](https://github.com/axiiomatic/reddit-downloader)
## [4][Build a Microservice with Go and GoLand](https://www.reddit.com/r/golang/comments/gmuey6/build_a_microservice_with_go_and_goland/)
- url: https://youtu.be/arZiFSerU1k
---

## [5][How I Structure Web Servers in Go](https://www.reddit.com/r/golang/comments/gnb3ih/how_i_structure_web_servers_in_go/)
- url: https://www.dudley.codes/posts/2020.05.19-golang-structure-web-servers/
---

## [6][I made a super simple Gmail filter organizer tool with Go!](https://www.reddit.com/r/golang/comments/gnapgh/i_made_a_super_simple_gmail_filter_organizer_tool/)
- url: https://www.reddit.com/r/golang/comments/gnapgh/i_made_a_super_simple_gmail_filter_organizer_tool/
---
I was looking through Gmail a few days back and realized that I had created so many filters for my many different labels. Most of them had a really common criteria: FROM &lt;email&gt;, ARCHIVE it, then ADD a given label. 

I had nearly 20 filters created just to add a set of different emails to a label for my school work - and that was just for one label. I dreaded the thought of trying to manage and sort through them manually. So, in programmer fashion, I decided to hack together a quick and simple CLI tool that would authenticate a Gmail account and then proceed to clean up the filters list.

Essentially, the tool will look through all the available filters, find the ones that match the criteria outlined above, and then re-organize the filters by deleting all of them and creating a single large filter to assign a set of emails to a given label. 

I tested it out on my own email account and managed to reduce the number of filters from 172 to 63! 

The Github repository is [here.](https://github.com/woojiahao/gmail-filter-organiser) If you do want to try it out, please export a copy of your existing filters first - you can do so by following [this guide.](https://www.lifewire.com/how-to-save-export-and-back-up-gmail-filters-1172109)

I hope that someone is able to find this useful! Any feedback is much appreciated, this is my third project using Go and my experience with Go has been fantastic so far :)
## [7][Create repeatable byte array of Go struct which contains a pointer](https://www.reddit.com/r/golang/comments/gn8wmt/create_repeatable_byte_array_of_go_struct_which/)
- url: https://www.reddit.com/r/golang/comments/gn8wmt/create_repeatable_byte_array_of_go_struct_which/
---
I want to be able to create repeatable byte arrays of structs in Go so I can hash them and then verify that hash at some point.

I am currently following this simple approach to create a byte array from a struct with: 

        []byte(fmt.Sprintf("%v", struct))...)

This works perfectly until my struct holds an embedded struct with a pointer, for example:

        type testEmbeddedPointerStruct struct {
        	T *testSimpleStruct
        }

In my tests this creates a different byte array each time, I think it may be because with the pointer the address in memory changes each time?

Is there a way of creating a repeatable byte array digest even if the struct holds a pointer?

Thanks
## [8][Beginner-Intermediate level Golang projects](https://www.reddit.com/r/golang/comments/gmvw0m/beginnerintermediate_level_golang_projects/)
- url: https://www.reddit.com/r/golang/comments/gmvw0m/beginnerintermediate_level_golang_projects/
---
I have been working with Go for the past 2 months and I am looking for project ideas to do something a little more advanced than to-do applications in golang. Please recommend some projects ideas to work on from scratch. The things I have worked on, mostly reside in one file, so I have only worked on simple cli's etc.

I must say, I am a little afraid to take up advanced projects, probably because I have a habit of giving up on them if the project seems too difficult :/
## [9][CLI based stock tracker / graph visualizer written in Go](https://www.reddit.com/r/golang/comments/gmm8f1/cli_based_stock_tracker_graph_visualizer_written/)
- url: https://github.com/ericm/stonks
---

## [10][Parsing HTML fragments with Go](https://www.reddit.com/r/golang/comments/gnagbp/parsing_html_fragments_with_go/)
- url: https://www.reddit.com/r/golang/comments/gnagbp/parsing_html_fragments_with_go/
---
Hello,

The other day, I tried to use [x/net/html](https://godoc.org/golang.org/x/net/html) to create a small web scraper. I struggled quite a bit with fragment parsing for my unit tests, and figured I might not be the only one.

Quick, off the top of your head: what HTML nodes will the following produce?`html.Parse(strings.NewReader("&lt;tr&gt;&lt;td&gt;1&lt;/td&gt;&lt;/tr")`If (like me!) you thought

    &lt;tr&gt;&lt;td&gt;1&lt;/td&gt;&lt;/tr&gt;

Then, you're wrong :)

It took me a while to understand why it does not work like that (spoiler: it is not a bug in the parser), and I made a little blog post, where I explain how I managed to solve my problem and what it taught me about the HTML spec.

Feel free to [check it out](https://nikodoko.com/posts/html-table-parsing/)!
