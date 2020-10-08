# golang
## [1][Proposal to add errors.ErrUnsupported is accepted](https://www.reddit.com/r/golang/comments/j74okf/proposal_to_add_errorserrunsupported_is_accepted/)
- url: https://github.com/golang/go/issues/41198#issuecomment-705078431
---

## [2][goc: A Comprehensive Coverage Testing System for Go, especially for complex scenarios, like e2e or system tests.](https://www.reddit.com/r/golang/comments/j78yp0/goc_a_comprehensive_coverage_testing_system_for/)
- url: https://www.reddit.com/r/golang/comments/j78yp0/goc_a_comprehensive_coverage_testing_system_for/
---
[https://github.com/qiniu/goc](https://github.com/qiniu/goc)
## [3][Cross Compilation (Linux-&gt;Windows)](https://www.reddit.com/r/golang/comments/j788kw/cross_compilation_linuxwindows/)
- url: https://www.reddit.com/r/golang/comments/j788kw/cross_compilation_linuxwindows/
---
Hey there. I've successfully compiled an example project from Linux to Windows using Mingw-w64, however, I have to statically link the standard c and c++ libraries which drastically increases the executable size. If I don't statically link then my program throws an error on runtime specifying that it cannot locate "libstdc++6.dll" or "libgcc\_s\_seh-1.dll".

Has anyone else run into this issue?

Edit: I reduced the file size using the following commands:

    strip --strip-unneeded test.exe
    upx test.exe

The initial strip reduced the file size to about 5,998KB from 20MB and then UPX reduced it further to  2,127KB.
## [4][Gonum is now an "affiliated project" of NumFOCUS](https://www.reddit.com/r/golang/comments/j6uer7/gonum_is_now_an_affiliated_project_of_numfocus/)
- url: https://www.reddit.com/r/golang/comments/j6uer7/gonum_is_now_an_affiliated_project_of_numfocus/
---
https://numfocus.salsalabs.org/numfocusnewsletter_september_2020?wvpId=7ed9ad03-52fb-4e84-96ee-a53d4bcadeb3

how cool is that?
yay us!
## [5][Lateralus - Phishing campaign from terminal](https://www.reddit.com/r/golang/comments/j7cx2q/lateralus_phishing_campaign_from_terminal/)
- url: https://www.reddit.com/r/golang/comments/j7cx2q/lateralus_phishing_campaign_from_terminal/
---
Hey guys, i have created tool to run phishing campaigns from terminal at following repo https://github.com/XdaemonX/lateralus. 

Let me know what you think about it
## [6][Best Language for the “Simple App”](https://www.reddit.com/r/golang/comments/j7bb9i/best_language_for_the_simple_app/)
- url: https://www.reddit.com/r/golang/comments/j7bb9i/best_language_for_the_simple_app/
---
Benchmarking programming languages for Simple Hello World Rest Services. 

Golang, Rust, Java(Spring Boot, Quarkus, Micronaut), NodeJS, .net Core

[https://medium.com/@emreodabas\_20110/best-language-for-simple-app-979729d3e48d](https://medium.com/@emreodabas_20110/best-language-for-simple-app-979729d3e48d)
## [7][Why is there no Try/Catch.](https://www.reddit.com/r/golang/comments/j7c8ut/why_is_there_no_trycatch/)
- url: https://www.reddit.com/r/golang/comments/j7c8ut/why_is_there_no_trycatch/
---
I have been a developer for a long time and new to Go.  I understand the concept of trying to keep things to a minimum (1 way to do things).   However, why do we not get a Try/Catch paradigm for error handling?  Having used this in most other languages, its incredibly useful and saves a lot of coding rather than checking err != nil after many calls in a single function.   Any particular reason why Go leaves this out?
## [8][File storage abstraction libraries suggestion?](https://www.reddit.com/r/golang/comments/j78q95/file_storage_abstraction_libraries_suggestion/)
- url: https://www.reddit.com/r/golang/comments/j78q95/file_storage_abstraction_libraries_suggestion/
---
Hi,

For my next project I'm looking for a file abstraction library for supporting both Local and S3 compatible storages (AWS S3, DigitalOcean Spaces, etc.).

The only one I could found was [Stow](https://github.com/graymeta/stow) but it seems that it's no longer actively maintained (based on the unanswered issues and last commit date) and there is almost no documentation.

Could you suggest another one or would you just create an interface on your own and go with it?
## [9][construct : Go generators for low abstraction persistence with PostgreSQL - Not an ORM](https://www.reddit.com/r/golang/comments/j6mt1u/construct_go_generators_for_low_abstraction/)
- url: https://github.com/networkteam/construct
---

## [10][How do I use a bearer token with the twitter api](https://www.reddit.com/r/golang/comments/j771qe/how_do_i_use_a_bearer_token_with_the_twitter_api/)
- url: https://www.reddit.com/r/golang/comments/j771qe/how_do_i_use_a_bearer_token_with_the_twitter_api/
---
I've been trying to use the new twitter api with golang for a fun project and while everything works in insomina, I cannot get the bearer token to work in golang. How do I get authorized using the bearer token  with my request? I've tried doing r*eq.Header.Add("Authorization","Bearer " + bearer)* and making my request but had no luck.
