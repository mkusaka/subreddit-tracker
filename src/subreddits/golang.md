# golang
## [1][LocalXpose: Bye Bye Localhost, Hello World.](https://www.reddit.com/r/golang/comments/gw6lo5/localxpose_bye_bye_localhost_hello_world/)
- url: https://www.reddit.com/r/golang/comments/gw6lo5/localxpose_bye_bye_localhost_hello_world/
---
&amp;#x200B;

https://preview.redd.it/ggl95vkghs251.png?width=2529&amp;format=png&amp;auto=webp&amp;s=c03d4a58cbc29d52d25132747c5f51410cb2a0d5

I created LocalXpose which is a reverse proxy that enables you to expose your localhost to the internet.

LocalXpose has many features like TCP, HTTP, TLS, UDP tunnels, built-in Let's Encrypt, unlimited connections, CLI &amp; GUI, SDK to start tunnels from your app and much more.

You can read more at [https://localxpose.io](https://localxpose.io)

&amp;#x200B;

&amp;#x200B;

https://i.redd.it/6qoalibihs251.gif
## [2][how to read the next 4 bytes in file ?](https://www.reddit.com/r/golang/comments/gwg7ek/how_to_read_the_next_4_bytes_in_file/)
- url: https://www.reddit.com/r/golang/comments/gwg7ek/how_to_read_the_next_4_bytes_in_file/
---
basically after i read starting 4 bytes of a file i want to read the next 4 bytes in the file without looping.basically specific next 4 bytes
## [3][How I solved Jepsen with OpenCensus Distributed Tracing: A personal journey](https://www.reddit.com/r/golang/comments/gwgaf6/how_i_solved_jepsen_with_opencensus_distributed/)
- url: https://dgraph.io/blog/post/solving-jepsen-with-opencensus/
---

## [4][Golang Authentication Service: allows setup authentication in minutes](https://www.reddit.com/r/golang/comments/gvwelw/golang_authentication_service_allows_setup/)
- url: https://github.com/maximthomas/gortas
---

## [5][Self-Host Complete Authentication Service in Go](https://www.reddit.com/r/golang/comments/gwh8vw/selfhost_complete_authentication_service_in_go/)
- url: https://www.reddit.com/r/golang/comments/gwh8vw/selfhost_complete_authentication_service_in_go/
---
Hi community, my first post here.

&amp;#x200B;

I've written an Authentication microservice, or more appropriately, a User microservice complete with all the features you'd expect. Here's some of what's included:

* JWT Based and API Key based authentication
* Support For Social Media Login Using Facebook/Google/Instagram/Twitter/LinkedIn.
* CRUD operations on user, with included endpoints for all the usual stuff such as Email Verification, Password Recovery, etc.
* Additional functionality on certain routes based on Admin role (Additional roles can easily be added).
* Complete session tracking with locations and time, ability for the user to manage their sessions.
* Created using Goa ([https://github.com/goadesign/goa](https://github.com/goadesign/goa)), making implementation of RPC super easy if needed.
* Ability to create Frontend SDKs for JavaScript/Android/iOS/others using openapi-generator within seconds.

This list is not exhaustive.

It uses MongoDB for database but the service is written in a manner where the database package can easily be swapped out for a different DB. Other features would include things such as Rate Limiting, Logging, etc.

&amp;#x200B;

**Do note that this is not a package that you can import into your code and use that way. This is a service in itself, which ideally someone can take, modify for their use and then deploy directly.**

&amp;#x200B;

In any case, here's the project: [https://github.com/anshap1719/authentication](https://github.com/anshap1719/authentication). Also, here's the API documentation for the same: [https://anshap1719.github.io/authentication/](https://anshap1719.github.io/authentication/).

&amp;#x200B;

Also note that this is not a 100% ready to use directly. You'll need to make some changes anyway such as setting up your own API keys and providing it with the Public and Private Key Files For JWT, etc.
## [6][Big File Uploader :- Upload big file by chunking the file using goroutines and channel](https://www.reddit.com/r/golang/comments/gw0axl/big_file_uploader_upload_big_file_by_chunking_the/)
- url: https://www.reddit.com/r/golang/comments/gw0axl/big_file_uploader_upload_big_file_by_chunking_the/
---
Hi community,

I have been learning golang for some time after office hours. I love the language. I have written a simple file uploader which writes the file(any type) onto disk. I would love a code review and feedback from the community.

[Github repo](https://github.com/atulkgupta9/big-file-uploader)
## [7][Requests and context](https://www.reddit.com/r/golang/comments/gwes7r/requests_and_context/)
- url: https://www.reddit.com/r/golang/comments/gwes7r/requests_and_context/
---
Hi,

I wonder how a server should deal with requests that take too much time to be handled.

I read the go's blog about the context. But the main thing that bothers me there is that the actual function that does the work should return anyhow. Either when &lt;-done is called or not. So the actual load will still be there, so what is the point in timeout (in server side perspective)?

I found it is OK if we don't wait for the function to send it's output to the channel, because we can buffer the channel. but in this situation the load is still there.

Is there a good way to handle it? Do I miss something?

link to the blog: [https://blog.golang.org/context](https://blog.golang.org/context)

Thanks everyone
## [8][Starting to learn Go](https://www.reddit.com/r/golang/comments/gw50ea/starting_to_learn_go/)
- url: https://www.reddit.com/r/golang/comments/gw50ea/starting_to_learn_go/
---
Hi, guys. I’m starting to learn Go. I have been coding for past few years in various languages and for networking I used Python as my choice. 
I have been looking at some stats about speed of Python and Go and figured that for what I’m doing (mostly networking and cli tools) Python is keeping me back. 

I want to learn this language.. from very start. So do you have any book or some advices what should I focus on?

Thank you :)
## [9][JMX Exporter in golang?](https://www.reddit.com/r/golang/comments/gwdmve/jmx_exporter_in_golang/)
- url: https://www.reddit.com/r/golang/comments/gwdmve/jmx_exporter_in_golang/
---
Hi Guys,

I want to create a custom Prometheus exporter for JMX monitoring of a java based application. is it possible to code it in golang?

 Prometheus is created with golang go there shouldn't be any issue there but my concern is can we extract Java memory management parameters (like heap/non heap memory, garbage collector information, native memory usage, metaspace etc.) using golang?

if so, is there any reference/source available in web (i checked but didn't find any lead)? 

Thanks in advance \^\_\^

Stay Safe!

Cheers!
## [10][go-guardian newly born library for golang authentication](https://www.reddit.com/r/golang/comments/gwd3sh/goguardian_newly_born_library_for_golang/)
- url: https://www.reddit.com/r/golang/comments/gwd3sh/goguardian_newly_born_library_for_golang/
---
Repo: [https://github.com/shaj13/go-guardian](https://github.com/shaj13/go-guardian)

GoDoc: [https://pkg.go.dev/github.com/shaj13/go-guardian@v1.1.1?tab=doc](https://pkg.go.dev/github.com/shaj13/go-guardian@v1.1.1?tab=doc)

It worth your star :)
