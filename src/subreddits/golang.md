# golang
## [1][[Q&amp;A] //go:build draft design](https://www.reddit.com/r/golang/comments/hitexe/qa_gobuild_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hitexe/qa_gobuild_draft_design/
---
I posted a draft design today for updating the // +build lines to fix some usability problems. 

Video: [https://golang.org/s/go-build-video](https://golang.org/s/go-build-video)\
Design: [https://golang.org/s/go-build-design](https://golang.org/s/go-build-design)

As an experiment, let's try doing Q&amp;A about the design here in Reddit.
My hope is that the threading support will help keep questions and answers matched.

**Please start a new top-level comment for each new question.**
## [2][Improved docker Go module dependency cache for faster builds in CI/CD](https://www.reddit.com/r/golang/comments/hj4n44/improved_docker_go_module_dependency_cache_for/)
- url: https://github.com/montanaflynn/golang-docker-cache
---

## [3][When microservices in Go are not enough: introduction to DDD Lite](https://www.reddit.com/r/golang/comments/hj9cxw/when_microservices_in_go_are_not_enough/)
- url: https://threedots.tech/post/ddd-lite-in-go-introduction/
---

## [4][The How and Why of Go, Part 1: Tooling](https://www.reddit.com/r/golang/comments/hin7sn/the_how_and_why_of_go_part_1_tooling/)
- url: http://okigiveup.net/the-how-and-why-of-go-part-1-tooling/
---

## [5][Setting up CircleCI and Sonarcloud integration with Go projects](https://www.reddit.com/r/golang/comments/hj9sz2/setting_up_circleci_and_sonarcloud_integration/)
- url: https://akondas.com/Tales-of-Devops-and-Go-Part-III/
---

## [6][How to send POST request service when using Traefik and Docker](https://www.reddit.com/r/golang/comments/hj9rst/how_to_send_post_request_service_when_using/)
- url: https://www.reddit.com/r/golang/comments/hj9rst/how_to_send_post_request_service_when_using/
---
Hi!

I recently started project based on microservices architecture and I have problem with calling from service A service B. I set up docker-compose and traefik, but I have problem with calling it when I use docker. Talk is cheap, so I'll show you code:

    req, err := http.NewRequest(http.MethodPost, http://127.0.0.1/api/v1/user/validate", bytes.NewBuffer(jsonCreds))
    if err != nil {
        log.Errorf("Couldn't prepare request: %v", err)
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")
    req.Host = "usersvc"
    res, err := s.httpClient.Do(req)
    if err != nil {
        log.Errorf("Couldn't to execute request: %v", err)
        return nil, err
    }

That's the part when I'm calling service B. When I run service A and service B separately and call `http://127.0.0.1/api/v1/user/validate` with `Host=usersvc` everything works ok, but when I want to execute the same code when I use \`docker-compose up\` I get this error: `{"level":"error","msg":"Couldn't to execute request: Post \"[``http://127.0.0.1/api/v1/user/validate\](http://127.0.0.1/api/v1/user/validate)``": dial tcp` [`127.0.0.1:80`](https://127.0.0.1:80)`: connect: connection refused","time":"2020-07-01T14:59:58Z"}`.

I guess there is problem with networking, but I don't really know how to fix it. How can I call that service using \`docker\`?

Thanks in advance for the help!
## [7][Rate Limiting in Golang](https://www.reddit.com/r/golang/comments/hj9ejw/rate_limiting_in_golang/)
- url: https://medium.com/@justin.graber/rate-limiting-in-golang-f3ed2c62df36
---

## [8][Pointer and value semantics in Go](https://www.reddit.com/r/golang/comments/hj98n4/pointer_and_value_semantics_in_go/)
- url: https://developer20.com/pointer-and-value-semantics-in-go/
---

## [9][Learning GO - version numbers in go.mod confusing me...](https://www.reddit.com/r/golang/comments/hj8z0q/learning_go_version_numbers_in_gomod_confusing_me/)
- url: https://www.reddit.com/r/golang/comments/hj8z0q/learning_go_version_numbers_in_gomod_confusing_me/
---
Hi,

Love the language - but versioning in go.mod is doing my head in.

[**https://github.com/alexellis/blinkt\_go**](https://github.com/alexellis/blinkt_go) **has a tag of "v1.0" but if I specify v1.0 in go.mod:**

    require github.com/alexellis/blinkt_go v1.0

**I get**

     /home/pi/projects/blinkt_test/go.mod:5: invalid module version "v1.0": no matching versions for query "v1.0"

Substituting **v1.0.0 for v1.0 works though.**

*If I remove the require line and leave the replace line, go get adds this:*

    require github.com/alexellis/blinkt_go v0.0.0-20180120180744-cc0ca163e0bc

**How does module versioning work?** I don't seem to be able to find a good explanation on the web. cc0ca16 is the commit ID of HEAD. Not sure where the rest is coming from - I see a date maybe.

**Am I better off just omitting require lines from go.mod, but supplying replace lines with versions that match tags I add to my forks\[1\]- and letting go get do what it wants?**

\[1\] In another Reddit post, I explain why I fork copies of code to my own repos - but in essence as insurance against them going away.

Cheers folks
## [10][Replacing GO import repos with my own forked copies - correct approach?](https://www.reddit.com/r/golang/comments/hj8sum/replacing_go_import_repos_with_my_own_forked/)
- url: https://www.reddit.com/r/golang/comments/hj8sum/replacing_go_import_repos_with_my_own_forked/
---
I am a paranoid type, so when building code, I don't want to assume a 3rd party repo will always be around. So I clone what I use into my own gitlab repo. eg:

[https://github.com/alexellis/blinkt\_go](https://github.com/alexellis/blinkt_go) =&gt; 

[https://gitlab.com/&lt;me&gt;-grp/ext/go/github.com/alexellis/blinkt\_go](https://gitlab.com/timjwatts-grp/ext/go/github.com/alexellis/blinkt_go)

(using gitlab groups to keep some symmetry of naming)

&amp;#x200B;

**1) So... if I add a** 

***replace*** [***https://github.com/alexellis/blinkt\_go***](https://github.com/alexellis/blinkt_go) ***=&gt;*** [***https://gitlab.com/&lt;me&gt;-grp/ext/go/github.com/alexellis/blinkt\_go***](https://gitlab.com/timjwatts-grp/ext/go/github.com/alexellis/blinkt_go) ***&lt;version&gt;***

**into my go.mod, this is the correct approach? (Seems to work)**

&amp;#x200B;

**2) If a library brings in other libraries indirectly - do I just need to add more replace statements to my final go.mod after cloning, or do I need to add go.mod files to my repo clone that is itself bringing in other libraries? It's the internals of the "go get" methodology I am struggling with.**

&amp;#x200B;

Cheers folks :)
## [11][How do I fix this error when running 'go build' or 'go install'](https://www.reddit.com/r/golang/comments/hj7t82/how_do_i_fix_this_error_when_running_go_build_or/)
- url: https://www.reddit.com/r/golang/comments/hj7t82/how_do_i_fix_this_error_when_running_go_build_or/
---
I'm getting this error when I try to run go install or go build. Would someone mind helping me out? Thanks.

can't load package: package .: no Go files in /Users/myname/projectname
