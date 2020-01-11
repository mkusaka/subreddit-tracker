# golang
## [1][How we optimized our DNS server using go tools](https://www.reddit.com/r/golang/comments/en4pds/how_we_optimized_our_dns_server_using_go_tools/)
- url: https://medium.com/@arash.cordi/how-we-optimized-our-dns-server-using-go-tools-d753e1a5e709
---

## [2][A dead simple proxy server that allows you to support HTTPS automatically with the Let's Encrypt certificate authority](https://www.reddit.com/r/golang/comments/en2yvk/a_dead_simple_proxy_server_that_allows_you_to/)
- url: https://www.reddit.com/r/golang/comments/en2yvk/a_dead_simple_proxy_server_that_allows_you_to/
---
If anyone is interested, I just wrote a simply proxy server (called gossl) that allows you to easily support HTTPS through Let's Encrypt. Here's an example of serving static files with HTTPS:

/home/ubuntu&gt; sudo ./gossl -staticDir=/path/to/your/static/content/dir -domains=[mydomain.com](https://slack-redir.net/link?url=http%3A%2F%2Fmydomain.com),[www.mydomain.com](https://slack-redir.net/link?url=http%3A%2F%2Fwww.mydomain.com)

That's really all you need to do. No libraries to install. No certificate renewal service to set up. It uses the autocert library in Go. I have instructions on my GitHub page explaining how to set it up as a service using systemd on Linux and how to write a proxy mapping file. You can proxy to several different servers and still serve up static files. One good use for it is running your primary server on localhost on a non-public port and using gossl to provide HTTPS support. I'm not sure if this is useful to anyone else or if I'm reinventing the wheel (don't care--was still interesting to write), but if so, it's open source. There's also a link to a binary (built for Ubuntu 18.04) on the page and instructions for how to build it: [https://github.com/metaphyze/gossl](https://slack-redir.net/link?url=https%3A%2F%2Fgithub.com%2Fmetaphyze%2Fgossl)
## [3][Multiplexing Channels in Go](https://www.reddit.com/r/golang/comments/en6nus/multiplexing_channels_in_go/)
- url: https://katcipis.github.io/blog/mux-channels-go/
---

## [4][Go 1.4 significant reduction in defer performance footprint](https://www.reddit.com/r/golang/comments/emr462/go_14_significant_reduction_in_defer_performance/)
- url: https://twitter.com/janiszt/status/1215601972281253888?s=19
---

## [5][ok-zoomer: A Go program that takes an image, uses pigo to detect a face, and creates a gif that zooms in on the face](https://www.reddit.com/r/golang/comments/en615f/okzoomer_a_go_program_that_takes_an_image_uses/)
- url: https://github.com/jbirms/ok-zoomer
---

## [6][Complete Guide to Create Docker Container for Your Golang Application](https://www.reddit.com/r/golang/comments/emxfxo/complete_guide_to_create_docker_container_for/)
- url: https://medium.com/@afdolriski/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e
---

## [7][Complete Go + Typescript application with proper system design, testing, and modern architecture explanations](https://www.reddit.com/r/golang/comments/emyffo/complete_go_typescript_application_with_proper/)
- url: https://github.com/short-d/short#system-design
---

## [8][Ann: Repotrace and srctrace](https://www.reddit.com/r/golang/comments/en6mtd/ann_repotrace_and_srctrace/)
- url: https://www.reddit.com/r/golang/comments/en6mtd/ann_repotrace_and_srctrace/
---
friends

The following is the first (of hopefully a few more to come) tools that I built for DevOps. Hoping it will be of use to the community. Both written in go and support a number of languages besides go.

[https://github.com/RajaSrinivasan/srctrace.git](https://github.com/RajaSrinivasan/srctrace.git)

[https://github.com/RajaSrinivasan/repotrace.git](https://github.com/RajaSrinivasan/repotrace.git)

best, srini
## [9][An experimental &amp; simple key-value service written in Go/C++](https://www.reddit.com/r/golang/comments/en6l5m/an_experimental_simple_keyvalue_service_written/)
- url: https://www.reddit.com/r/golang/comments/en6l5m/an_experimental_simple_keyvalue_service_written/
---
This is our fresher training project in [ZaloPay.](https://github.com/zalopay-oss) [GodBee](https://github.com/zalopay-oss/godbee) is a Key-Value Store Service project. In this project, we choose B-Tree and B+Tree data structures to organize and manipulate data. Key-Value Storage is written in C++ and Service layer is written in Golang programming language. We use gRPC services to handle requests from client and use CGO to access data from C++ storage.
## [10][maildir-tools scripting tools for Maildirs, along with a simple TUI based email-client.](https://www.reddit.com/r/golang/comments/en5z7a/maildirtools_scripting_tools_for_maildirs_along/)
- url: https://github.com/skx/maildir-tools
---

