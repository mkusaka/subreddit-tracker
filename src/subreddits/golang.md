# golang
## [1][Harmony - A free and open source alternative to Discord made in Go [VERY WIP]](https://www.reddit.com/r/golang/comments/f52p6a/harmony_a_free_and_open_source_alternative_to/)
- url: https://github.com/harmony-development/harmony-server
---

## [2][🚀 Visualizing memory management in Golang](https://www.reddit.com/r/golang/comments/f55qml/visualizing_memory_management_in_golang/)
- url: https://deepu.tech/memory-management-in-golang/
---

## [3][pkg.go.dev will be open-sourced](https://www.reddit.com/r/golang/comments/f56ua9/pkggodev_will_be_opensourced/)
- url: https://groups.google.com/d/msg/golang-dev/mfiPCtJ1BGU/OZ9cU3SgBgAJ
---

## [4][Had written web service in Go, got no users, so decided to open the code (web server + background worker)](https://www.reddit.com/r/golang/comments/f56m6p/had_written_web_service_in_go_got_no_users_so/)
- url: https://www.reddit.com/r/golang/comments/f56m6p/had_written_web_service_in_go_got_no_users_so/
---
Hi, I've written an MVP web service in Go [link](https://noiseremover.net). It's a web service that cleans audio files from background noise. It uses simple smoothing algorithm for that so nothing fancy.

Apparently people are not interested in this type of service (at least in the current form). So I decided to open the code here [github.com/normie7/nore](https://github.com/normie7/nore)

Any feedback is appreciated.
## [5][sqlc - Generate your Go database code with ease](https://www.reddit.com/r/golang/comments/f4suu3/sqlc_generate_your_go_database_code_with_ease/)
- url: https://www.youtube.com/watch?v=uBPXNREhZZw
---

## [6][Hazelcast / Open Source Distributed Caching for Go](https://www.reddit.com/r/golang/comments/f5875o/hazelcast_open_source_distributed_caching_for_go/)
- url: https://www.reddit.com/r/golang/comments/f5875o/hazelcast_open_source_distributed_caching_for_go/
---
Hi all,

Hazelcast is a distributed in-memory object store and compute, supporting a wide variety of data structures such as Map, Set, List, MultiMap, RingBuffer, HyperLogLog. It is cloud &amp; Kubernetes friendly.

I wanted to let you know that we have prepared a Code Reference Card for Hazelcast Go client: [https://hazelcast.com/resources/go-client-ref-card/](https://hazelcast.com/resources/go-client-ref-card/)

You can visit the Github repo for [code samples and API documentation](https://github.com/hazelcast/hazelcast-go-client).

Here is the command to install the Hazelcast Go client:

    go get github.com/hazelcast/hazelcast-go-client

Currently, we are working very hard on the next major release, i.e v4.0. We'd be really happy to hear your feedback :)

Disclaimer: I'm working at Hazelcast as part of the Clients Team. If you have any feature requests or any feedback, please let me know!

All the best, Burak.
## [7][Micro In Action, Part 5: Message Broker](https://www.reddit.com/r/golang/comments/f57kj2/micro_in_action_part_5_message_broker/)
- url: https://itnext.io/micro-in-action-part-5-message-broker-a3decf07f26a
---

## [8][Any specific reason why the Golang source code uses variable names like "i", "j", "r", "w" instead of explicit alternatives like "index" "request", "writer"?](https://www.reddit.com/r/golang/comments/f5574h/any_specific_reason_why_the_golang_source_code/)
- url: https://www.reddit.com/r/golang/comments/f5574h/any_specific_reason_why_the_golang_source_code/
---
A good example is:

`http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {`   
`fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))`

`})`

But I see it throughout go packages and the golang source code iteself. Isn't it cleaner (naming wise) to use more explicit variables like "responseWriter" instead of "r" the function above could be

`http.HandleFunc("/bar", func(responseWriter http.ResponseWriter, request *http.Request) {`   
`fmt.Fprintf(responseWriter, "Hello, %q", html.EscapeString(request.URL.Path))`   
`})`

Or am I missing something?
## [9][A friend asked me how to write a resilient service worker in Go . I figured I should share it to anyone interested.](https://www.reddit.com/r/golang/comments/f4ov4d/a_friend_asked_me_how_to_write_a_resilient/)
- url: https://gist.github.com/System-Glitch/301e95975a2645b8ea57c47b0c7cfef4
---

## [10][Verify XML with XSD in Golang ? ( windows )](https://www.reddit.com/r/golang/comments/f585pi/verify_xml_with_xsd_in_golang_windows/)
- url: https://www.reddit.com/r/golang/comments/f585pi/verify_xml_with_xsd_in_golang_windows/
---
I am trying to get a verification process going of checking a XML against its XSD file.

Been googling abit and can see different Golang libraries that appears to do that - but none of them seems to be able to compile in windows due to pkg-config returning a 9 digit error code.

Trying to get pkg-config to work on windows been to no avail simply dont have enough info on how to anywhere.

But wondering if anyone sucessfully had verification done on an XML towards a XSD in Golang ( preferable a proof working example ).
