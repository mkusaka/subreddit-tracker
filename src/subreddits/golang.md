# golang
## [1][Why Go’s Error Handling is Awesome](https://www.reddit.com/r/golang/comments/hmnhkz/why_gos_error_handling_is_awesome/)
- url: https://rauljordan.com/2020/07/06/why-go-error-handling-is-awesome.html
---

## [2][gosh now supports snippets](https://www.reddit.com/r/golang/comments/hmujvh/gosh_now_supports_snippets/)
- url: https://www.reddit.com/r/golang/comments/hmujvh/gosh_now_supports_snippets/
---
`gosh`, the tool for writing command line scripts in Go now lets you keep snippets of code and embed them by name in your scripts. For instance you can check for errors and report them by having a file called `iferr` and then after every line that sets the error you can add the snippet.

You install `gosh` by running this command

`go get -u` [`github.com/nickwells/utilities/gosh`](https://github.com/nickwells/utilities/gosh)

Once you have installed it you can see a usage message by running

`gosh -help`
## [3][Help with an algorithm to cut string into multiple strings](https://www.reddit.com/r/golang/comments/hmugl2/help_with_an_algorithm_to_cut_string_into/)
- url: https://www.reddit.com/r/golang/comments/hmugl2/help_with_an_algorithm_to_cut_string_into/
---
Hi!

&amp;#x200B;

I have a string of characters like so:

&amp;#x200B;

&amp;#x200B;

999923999999866544487221888888899988

&amp;#x200B;

And would like to cut it so that I only have the numbers greater than 2 into a slice of strings:

&amp;#x200B;

9999

3999999866544487

888888899988

&amp;#x200B;

Is this possible please?

Thank you in advance!
## [4][A Data Engineering Perspective on Go vs. Python (Part 2 - Dataflow)](https://www.reddit.com/r/golang/comments/hmu216/a_data_engineering_perspective_on_go_vs_python/)
- url: https://chollinger.com/blog/2020/07/a-data-engineering-perspective-on-go-vs.-python-part-2-dataflow/
---

## [5][How I Structure Go Packages](https://www.reddit.com/r/golang/comments/hm34kq/how_i_structure_go_packages/)
- url: https://bencane.com/stories/2020/07/06/how-i-structure-go-packages/
---

## [6][Little Go CLI tool to flatten JSON input](https://www.reddit.com/r/golang/comments/hmo8my/little_go_cli_tool_to_flatten_json_input/)
- url: https://gist.github.com/benhoyt/05e22f0d52f7feaf88b0a37f8b739f91
---

## [7][dgraph-io/ristretto v0.0.3 Released - Ristretto is a fast, concurrent cache (Go) library built with a focus on performance and correctness.](https://www.reddit.com/r/golang/comments/hmdwzn/dgraphioristretto_v003_released_ristretto_is_a/)
- url: https://github.com/dgraph-io/ristretto/blob/master/CHANGELOG.md#003---2020-07-06
---

## [8][Help, portscanner cant get past port 1017!](https://www.reddit.com/r/golang/comments/hmohoo/help_portscanner_cant_get_past_port_1017/)
- url: https://www.reddit.com/r/golang/comments/hmohoo/help_portscanner_cant_get_past_port_1017/
---
[https://github.com/blackhat-go/bhg/blob/master/ch-2/tcp-scanner-final/main.go](https://github.com/blackhat-go/bhg/blob/master/ch-2/tcp-scanner-final/main.go)

I have been following a book and recreated the code above. I added a print statement right above "port := &lt;-results" in the third for loop. I wouldn't go past 1017. I believe this is because the requests channel only has 1017 but i don't know why this could be. I'm also not sure if this is just a problem with my computer or the code. The book this is from is quite new and from a publisher with a good reputation so I think it might be my computer.
## [9][Python and Go : Part II - Extending Python With Go](https://www.reddit.com/r/golang/comments/hmakgs/python_and_go_part_ii_extending_python_with_go/)
- url: https://www.ardanlabs.com/blog/2020/07/extending-python-with-go.html
---

## [10][[PROJECT] GhostDB - A distributed, in-memory, general purpose datastore](https://www.reddit.com/r/golang/comments/hmbwqs/project_ghostdb_a_distributed_inmemory_general/)
- url: https://www.reddit.com/r/golang/comments/hmbwqs/project_ghostdb_a_distributed_inmemory_general/
---
TLDR; We have built a distributed in-memory datastore, would love feedback and if you're interested - some contributions (beginners welcome)!

Hi everyone,

First off, I'd like to thank a lot of people on this subreddit, you've been a huge help throughout the development of this project.

We have been working on this project for a little while now and we're ready to release it. It was originally developed as a final year project for university but it turned out to be pretty useful. When we started development on this, we had 0 knowledge of golang and have learned a tonne since we started.

Below is a description but you can check out the repo [here](https://github.com/GhostDB/GhostDB):

GhostDB is a distributed, in-memory, general purpose key-value data store that delivers microsecond performance at any scale.

GhostDB provides a very large hash table that is distributed across multiple machines and stores large numbers of key-value pairs within the hash table. However, it's more than just a big hash table. GhostDB has multiple levels of on-disk persistence in the form of point-in-time snapshots and an append-only-file used to restore a cache node if it crashes, concurrent crawlers that will automatically evict stale data, fine grained metrics, and if configured correctly, can serve 200k+ requests per second from more than 1.5M concurrent keep-alive connections per physical server.

Users of GhostDB can configure GhostDB nodes on machines they have control over. The clients application servers will be capable of interacting with GhostDB through the GhostDB SDKs much the same way their application servers may interface with a MySQL or MongoDB instance.

Our SDKs are available for Python and Node.js (Java and Golang SDKs in development) and are simple to use. With just 4 lines of extra code in your application servers, you are able to achieve up to a 25x increase in data retrieval speeds when compared to databases such as MongoDB and MySQL.

We have a vision for GhostDB and have spent some time refactoring it before releasing it as open source in order to facilitate extending it further.

In the future we plan to add: direct support for queues and other data structures, transactional read/write operations and a whole bunch more.

We would love to hear feedback and if it's something you're interested in, we would love you to contribute - we're beginner friendly!

Check out the repo here:  [https://github.com/GhostDB/GhostDB](https://github.com/GhostDB/GhostDB)
