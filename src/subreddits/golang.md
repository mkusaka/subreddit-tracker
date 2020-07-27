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

## [3][CGo-free SQLite database/sql driver for linux/amd64 v1.4.0-beta1 is released](https://www.reddit.com/r/golang/comments/hyenjh/cgofree_sqlite_databasesql_driver_for_linuxamd64/)
- url: https://www.reddit.com/r/golang/comments/hyenjh/cgofree_sqlite_databasesql_driver_for_linuxamd64/
---
From the change log at https://godoc.org/modernc.org/sqlite#hdr-Changelog

2020-07-26 v1.4.0-beta1:

The project has reached beta status while supporting linux/amd64 only at the moment. The 'extraquick' Tcl testsuite reports

    630 errors out of 200177 tests on  Linux 64-bit little-endian

and some memory leaks

    Unfreed memory: 698816 bytes in 322 allocations

Please try your production load tests and unit tests with the new version. Your help testing these pre-release versions is invaluable.

Report any problems using the issue tracker (requires a GitLab account):
https://gitlab.com/cznic/sqlite/-/issues/new?issue%5Bassignee_id%5D=&amp;issue%5Bmilestone_id%5D=
## [4][Examples of how to create a modern &amp; robust RESTful api server?](https://www.reddit.com/r/golang/comments/hy8a96/examples_of_how_to_create_a_modern_robust_restful/)
- url: https://www.reddit.com/r/golang/comments/hy8a96/examples_of_how_to_create_a_modern_robust_restful/
---
I'm looking into a developing a RESTful API server in Go. I'm having a hard time finding a lot of good examples, or common patterns/libraries people use. I can see how to string a bunch of libraries together but I'm looking for suggestions! Here's what I'd to do:

* OAuth2 server
* API serving. There's 1k ways of doing this, but very few integrate OAuth2
* Authorizing credential creation via OpenID
* OpenAPI support. At the very least generating a spec from the API.
* Desired but not required would be a developer portal

[https://www.ory.sh/](https://www.ory.sh/) is the closest project I've found and I'm currently evaluating it. Any other recommendations?
## [5][k6 plugin to load test Apache Kafka with support for Avro messages](https://www.reddit.com/r/golang/comments/hyp9rv/k6_plugin_to_load_test_apache_kafka_with_support/)
- url: https://www.reddit.com/r/golang/comments/hyp9rv/k6_plugin_to_load_test_apache_kafka_with_support/
---
I wrote a k6 plugin to load test Apache Kafka with support for Avro messages. k6 is a FOSS performance and load-testing tool, written in Go and scriptable in JavaScript. The k6 [plugin support PR](https://github.com/loadimpact/k6/pull/1396) is not merged yet, and is highly experimental. If you have feedback and questions, reach out to me directly.
&lt;https://github.com/mostafa/k6-plugin-kafka&gt;
## [6][really need an advice](https://www.reddit.com/r/golang/comments/hys99n/really_need_an_advice/)
- url: https://www.reddit.com/r/golang/comments/hys99n/really_need_an_advice/
---
first of all, i live somewhere in middle east, i don't have any work experience and i don't consider myself as an real backend developer yet, cause i still need some time. do you think i can get a Go job in an developed country as a junior with no background? maybe a remote job could also work for me. if not in Go, do i have more chance in other langs?
## [7][JWT and microservices](https://www.reddit.com/r/golang/comments/hyotg8/jwt_and_microservices/)
- url: https://www.reddit.com/r/golang/comments/hyotg8/jwt_and_microservices/
---
Hello everyone ! It seems that people involved in microservices development use JWT or somehow connected with it because most products such as service meshes (e.g. Istio) , so-called gateways (Kraken) and others relies heavily on jwt. 

Today i read https://dchest.com/authbook/ and it seems that author is a well-known developer. 
Let me quote some things from this book:

"There is a huge weakness: tokens are valid until they expire. There is no real log out or a possibility to revoke other sessions before they expire"

"These weaknesses make this scheme completely unusable for anything other than toy projects.
The funny thing is that some developers try to fix the first problem by introducing a list of revoked tokens, which the server consults when validating tokens. Then they add bloom filters and other complications to make this list manageable"

Do you agree with an author? How do you manage jwts, including bans, update fields included in jwt, multiple devices? Do you have a resource explaining all the complex stuff without highlighting only advantages? Do you consider jwt as unusable in practice?
## [8][My First go application, a very simple slack bot](https://www.reddit.com/r/golang/comments/hyrezs/my_first_go_application_a_very_simple_slack_bot/)
- url: https://www.reddit.com/r/golang/comments/hyrezs/my_first_go_application_a_very_simple_slack_bot/
---
[https://github.com/rimonmostafiz/frodobot](https://github.com/rimonmostafiz/frodobot)

I have been playing with GO for the last couple of days and wrote this slack bot.In my workplace, we use slack we have a channel where we post our daily status update (kind of scrum). Team members sometimes forgot to post their status.

1. The bot will start at every day 10:45 AM
2. Read all the messages from 6.00 AM to 10:45 AM
3. List out users who didn't post their status
4. Send a soft reminder message to the channel tagging those users

Please suggest to me how can I improve this project? Thanks :) 
## [9][What’s the best place to find a go dev who has experience making mobile apps with go? Small demo for a FOSS project I’d like help building, negotiable on the pay.](https://www.reddit.com/r/golang/comments/hyr2su/whats_the_best_place_to_find_a_go_dev_who_has/)
- url: https://www.reddit.com/r/golang/comments/hyr2su/whats_the_best_place_to_find_a_go_dev_who_has/
---

## [10][multi-gitter: Personal project I've been working on that allows for changes in multiple git repositories at the same time](https://www.reddit.com/r/golang/comments/hypr3u/multigitter_personal_project_ive_been_working_on/)
- url: https://github.com/lindell/multi-gitter
---

## [11][Dynamically layout images of various sizes](https://www.reddit.com/r/golang/comments/hypnod/dynamically_layout_images_of_various_sizes/)
- url: https://www.reddit.com/r/golang/comments/hypnod/dynamically_layout_images_of_various_sizes/
---
First time poster, long time lurker. I have come across a problem that I would like to solve in go. I am sure I am not the first or the last person to want to do this, yet I am unable to find anything for go. I am hoping the community can assist me. 

So the idea is to dynamically layout images of different sizes. Yes, this can be done with JS libraries or in-browser with relative ease, but I would like to have a go implementation. I have found an [article](https://medium.com/@jtreitz/the-algorithm-for-a-perfectly-balanced-photo-gallery-914c94a5d8af) doing exactly this, but for web. It also references an SO [post](https://stackoverflow.com/questions/7938809/how-to-understand-the-dynamic-programming-solution-in-linear-partitioning/7942946#7942946) containing some python examples on solving this using linear partitions. Attempting to implement it in go myself however, my math and python skills have let me down. 

If we could have an example of this in go, I am sure others can benefit from it, and we can make a pull request to have it added to [TheAlgorithms/Go](https://github.com/TheAlgorithms/Go)
## [12][Selectively marshal JSON](https://www.reddit.com/r/golang/comments/hyouiw/selectively_marshal_json/)
- url: https://www.reddit.com/r/golang/comments/hyouiw/selectively_marshal_json/
---
If I have a big struct with lots of nested objects, and want to apply a “whitelist” of fields that should be returned in my API, what would be the best way to accomplish this?

E.g. for:

    
    type User struct {
    	Email   string   `json:"email"`
    	Address *Address `json:"address"`
    }
    
    type Address struct {
    	Street  string `json:"street"`
    	Zipcode string `json:"zipcode"`
    }

I might want to marshal using something like:

    Ser(user, []string{"email", "address.street"}) =&gt; 
    
    {
        "email": ...,
        "address": { "street": ... }
    }
