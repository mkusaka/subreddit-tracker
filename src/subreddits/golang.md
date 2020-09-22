# golang
## [1][variations on fizzbuzz](https://www.reddit.com/r/golang/comments/ixedgq/variations_on_fizzbuzz/)
- url: https://www.reddit.com/r/golang/comments/ixedgq/variations_on_fizzbuzz/
---
I was watching a vid on youtube talking about function composition, and it reminded me of doing pipelines with channels in go, so I got the stupid idea to write fizzbuzz with channels.

[fizzbuzz in go with channels](https://github.com/pdk/fizzbuzz)

Showed it to some friends, and then one was like "I guess this means I have to write it in rust", so he wrote 

[fizzbuzz in rust](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=f614c481e9e075d259ef279c5f37006b)
and then [fizzbuzz in rust #2](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=4386bf1302376f3e45e4252eecb571e1)

Then I was like "I can write this in go more functional style" and got

[func-style fizzbuzz in go](https://bit.ly/3chFGoG)

which is slightly interesting, but incredibly ugly, so I rewrote as

[func-style fizzbuzz in go #2](https://bit.ly/3mFCmZh)

which I like much better.

Of course, all of this is reminiscent of [FizzBuzzEnterpriseEdition](https://github.com/EnterpriseQualityCoding/FizzBuzzEnterpriseEdition)

(edit to fix link)
## [2][AtomKV: in-memory, JSON, key-value service with compare-and-swap updates and event streams.](https://www.reddit.com/r/golang/comments/ix9ml9/atomkv_inmemory_json_keyvalue_service_with/)
- url: https://github.com/skeeto/atomkv
---

## [3][k6 v0.28.0 is out!](https://www.reddit.com/r/golang/comments/iwzy8u/k6_v0280_is_out/)
- url: https://www.reddit.com/r/golang/comments/iwzy8u/k6_v0280_is_out/
---
We are happy to announce that k6 v0.28.0 is released with new features and improvements. k6 is a modern open-source performance and load-testing tool, written in Go and scriptable in JavaScript. 🎊🎉🥳

- k6 Cloud execution logs
- Pushing logs to Loki
- Optional port to host mappings
- Support for specifying data types to InfluxDB fields
- Support for automatic gzip-ing of the CSV output result
- UX improvements
- Lots of bugfixes

You can read more about it in the [release notes](https://github.com/loadimpact/k6/releases/tag/v0.28.0). We'd be happy to hear your feedback about it.
## [4][Trying to find a way to code daily](https://www.reddit.com/r/golang/comments/ixn337/trying_to_find_a_way_to_code_daily/)
- url: https://www.reddit.com/r/golang/comments/ixn337/trying_to_find_a_way_to_code_daily/
---
Hello all. My process of learning Go has been ongoing and at this point im starting to feel like im getting stuck in tutorial hell. I did Todd Mcleods intro to go about a year ago and really enjoyed that class. Iv since moved in to doing cs50x to get some foundational CS knowledge and am working my way through jon Calhons Web Dev course. I know the old saying "work on any project" but I guess my issue is im not sure where to even begin. 

 At this point I felt comfortable enough to write a small go app at work for simple influxdb parsing but since my job doesnt not necessarily offer too many opportunities to code id like to find a project, opensource or even possibly some sort of contract work to really rely less on courses to continue learning and start implementing some really world solutions.

For context Iv been digging through many oreilly go books and even watched a couple of videos from the Ultimate Go courses to get better familiar with channels and go routines but I also feel that until im met with a task that requires them I wont fully grasp their usage. My day to day revolves around implementing "devops" practices via ci/cd pipelines and working with a java based team to maintain, deploy, and migrate their apps to k8s with more focus on GCP in the near future.
## [5][If you want to dive DEEP into quality code with Go, definitely check out golangci-lint](https://www.reddit.com/r/golang/comments/ix2xzd/if_you_want_to_dive_deep_into_quality_code_with/)
- url: https://www.reddit.com/r/golang/comments/ix2xzd/if_you_want_to_dive_deep_into_quality_code_with/
---
[https://github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint)

It's essentially a linter that combines all popular linters (gosec, gocritic, govet, etc.).

It's a bit harsh so you may have to disable some checks, for example

`golangci-lint run --disable unused --disable deadcode`

It installs on all major OSs if you want to run it locally: [https://golangci-lint.run/usage/install/#local-installation](https://golangci-lint.run/usage/install/#local-installation)

Or in a CICD environment: [https://golangci-lint.run/usage/install/#ci-installation](https://golangci-lint.run/usage/install/#ci-installation)
## [6][Go Time - Community Q&amp;A using Questions from Reddit](https://www.reddit.com/r/golang/comments/ix3a1w/go_time_community_qa_using_questions_from_reddit/)
- url: https://changelog.com/gotime/147
---

## [7][mongo-go-driver: How to get the database name from the connection string?](https://www.reddit.com/r/golang/comments/ixcut4/mongogodriver_how_to_get_the_database_name_from/)
- url: https://www.reddit.com/r/golang/comments/ixcut4/mongogodriver_how_to_get_the_database_name_from/
---
Hey all -- I'm trying to wrap my head around the mongo go driver, specifically getting the database name from the connection string!  All of the docs point to writing a function like this:

 `collection := client.Database("testing").Collection("numbers")` 

So in this example, in my configuration I would need to keep "testing" as its own string...  in NodeJS I could keep the my mongo connection string like this...

`'mongodb://localhost:27017/testing'`

... and the NodeJS mongo driver would know to connect to the "testing" database.  Is there any comparable feature in GoLang?  When I deploy this to Heroku, for instance, I'd like to keep the entire connection string as a single config variable and not have to worry about either specifically creating a DB with a specific name or needing a second config variable with the database name separately...  Is this possible without some string splitting magic using the default **mongo-go-driver** package?

Thanks in advance! (and sorry for my ignorance... this is day one of Golang for me....)
## [8][go-redis cannot set more than 9999 keys](https://www.reddit.com/r/golang/comments/ix44m6/goredis_cannot_set_more_than_9999_keys/)
- url: https://www.reddit.com/r/golang/comments/ix44m6/goredis_cannot_set_more_than_9999_keys/
---
I am trying to load a text/html response into Redis, however every time it seems to stop at 9999 keys.

I've split the response and iterate over it using if statement.

&amp;#x200B;

Example code: [https://pastebin.com/ARuE6v3t](https://pastebin.com/ARuE6v3t)

&amp;#x200B;

After the code finishes running I can get into redis-cli and manually add more keys using SET, but cannot understand why the loop will not get past 9999.
## [9][gomponents, a library for building reusable components of HTML DOM nodes](https://www.reddit.com/r/golang/comments/ix1lzh/gomponents_a_library_for_building_reusable/)
- url: https://github.com/maragudk/gomponents
---

## [10][A load testing tool with a real-time analyzer, written in Go](https://www.reddit.com/r/golang/comments/iwdx5v/a_load_testing_tool_with_a_realtime_analyzer/)
- url: https://i.redd.it/bf5f5nb42bo51.gif
---

