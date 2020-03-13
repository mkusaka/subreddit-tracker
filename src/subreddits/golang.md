# golang
## [1][VScode and Go suddenly don't get along](https://www.reddit.com/r/golang/comments/fhptjx/vscode_and_go_suddenly_dont_get_along/)
- url: https://youtu.be/gkHYhkqVhb8
---

## [2][Be careful fetching user-supplied URLs.](https://www.reddit.com/r/golang/comments/fhio12/be_careful_fetching_usersupplied_urls/)
- url: https://www.reddit.com/r/golang/comments/fhio12/be_careful_fetching_usersupplied_urls/
---
If you have a service which fetches user-supplied URLs, for example to convert to PDF, minify, or similar, you want to make sure people can't submit things like:

* file:////etc/passwd
* [http://localhost/server-status](http://localhost/server-status)
* [http://169.254.169.254/latest/meta-data/](http://169.254.169.254/latest/meta-data/)

To avoid these kind of things I put together a small wrapper, which denies resources to "local" resources.  It might be useful, although the general advice of sanitizing user-supplied input is hopefully more general, and more useful.  (I've reported **numerous** security issues like this over the years, for example [this markdown converter](https://blog.steve.fi/if_your_code_accepts_uris_as_input__.html).)

Anyway a simple library/wrapper:

[https://github.com/skx/remotehttp](https://github.com/skx/remotehttp)
## [3][Debug high memory allocations and GC cycles downloading large S3 objects](https://www.reddit.com/r/golang/comments/fhuxly/debug_high_memory_allocations_and_gc_cycles/)
- url: https://link.medium.com/OT9h008vO4
---

## [4][Live code reload for Go apps](https://www.reddit.com/r/golang/comments/fhmr82/live_code_reload_for_go_apps/)
- url: https://www.mailwizz.com/blog/live-code-reload-for-go-apps/
---

## [5][Open source projects that could use help](https://www.reddit.com/r/golang/comments/fhp9j8/open_source_projects_that_could_use_help/)
- url: https://www.reddit.com/r/golang/comments/fhp9j8/open_source_projects_that_could_use_help/
---
I've been using Go for a few years now outside of work and I'm trying to find a full time position where the majority of the work would involve Go. I've creating a few projects on GitHub but I'd like to move on to working with others on some larger scale projects so I learn the skills involved with that. If anyone knows of a project that can use some help and that would provide some good opportunities for learning more please let me know. Thanks.
## [6][Golang service to connect with Active Directory](https://www.reddit.com/r/golang/comments/fhpxsy/golang_service_to_connect_with_active_directory/)
- url: https://www.reddit.com/r/golang/comments/fhpxsy/golang_service_to_connect_with_active_directory/
---
Hi guys, im working on a Go microservice to connect with a Active Directory server (in aws) for authentication and get data and response back. Can you suggest which Golang library to use and how to approach this.

thanks
## [7][importing the same thing in multiple files with the same package](https://www.reddit.com/r/golang/comments/fhxeaz/importing_the_same_thing_in_multiple_files_with/)
- url: https://www.reddit.com/r/golang/comments/fhxeaz/importing_the_same_thing_in_multiple_files_with/
---
let's say we have package named "package"

and we have multiple files 

test1.go test2.go 

i need "fmt " in both of them  i have to import it in just one file or both of them
## [8][Build PWA with Go and WebAssembly](https://www.reddit.com/r/golang/comments/fhaneq/build_pwa_with_go_and_webassembly/)
- url: https://github.com/maxence-charriere/go-app
---

## [9][Build a backend server for Android Application in Golang?](https://www.reddit.com/r/golang/comments/fht5w9/build_a_backend_server_for_android_application_in/)
- url: https://www.reddit.com/r/golang/comments/fht5w9/build_a_backend_server_for_android_application_in/
---
I'm looking to build a backend server for an android application where it will send data from sensors from the mobile devices to the server every 3 hours. I want to build the server in Golang. Is it possible to build this server in Golang for Mobile? Would Go be the best programming language? I really want to do it in Golang, but wouldn't mind hearing other options. Can you also provide me resources to get started?
## [10][Go: Slice and Memory Management](https://www.reddit.com/r/golang/comments/fhh6jw/go_slice_and_memory_management/)
- url: https://medium.com/a-journey-with-go/go-slice-and-memory-management-670498bb52be
---

