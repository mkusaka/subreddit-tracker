# golang
## [1][Writing an API using Golang](https://www.reddit.com/r/golang/comments/gt84cm/writing_an_api_using_golang/)
- url: https://www.reddit.com/r/golang/comments/gt84cm/writing_an_api_using_golang/
---
I am extremely new to Golang and I have been using a Udemy course to learn about the language. I am starting a new job as a back-end developer and my first task is to create an API using Golang. If anyone has any good guides/youtubes to throw my way that would be great. Or if anyone has any pointers or best practices please give me all the knowledge. Thanks!
## [2][Digging deeper into the analysis of Go-code](https://www.reddit.com/r/golang/comments/gtblh7/digging_deeper_into_the_analysis_of_gocode/)
- url: https://nakabonne.dev/posts/digging-deeper-into-the-analysis-of-go-code
---

## [3][How do you set up shared models for microservices in your project?](https://www.reddit.com/r/golang/comments/gtdys7/how_do_you_set_up_shared_models_for_microservices/)
- url: https://www.reddit.com/r/golang/comments/gtdys7/how_do_you_set_up_shared_models_for_microservices/
---
We have some models that are shared across some microservices and they are a pita to support. Changes to the models usually have to be merged in tandem or CI breaks, same for deployments. Some people are suggesting joining the repos into a monorepo for the services.
## [4][GoLand 2020.2 Early Access Program starts](https://www.reddit.com/r/golang/comments/gswlfj/goland_20202_early_access_program_starts/)
- url: https://blog.jetbrains.com/go/2020/05/29/goland-2020-2-eap-is-open/
---

## [5][immudb released version 0.6 - a lightweight, high-speed immutable database for systems and applications. New website https://immudb.io including documentation released as well. Open Source and wrtten in Go.](https://www.reddit.com/r/golang/comments/gsyjwa/immudb_released_version_06_a_lightweight/)
- url: https://github.com/codenotary/immudb
---

## [6][nitr-agent Remote Monitoring Tool](https://www.reddit.com/r/golang/comments/gt9euo/nitragent_remote_monitoring_tool/)
- url: https://www.reddit.com/r/golang/comments/gt9euo/nitragent_remote_monitoring_tool/
---
I want to share with you a project I am currently working on.

nitr-agent is a crossplatform remote monitoring tool written in Golang, providing system and hardware information through a JSON APIproviding system and hardware information through a JSON API.

[https://github.com/juanhuttemann/nitr-agent](https://github.com/juanhuttemann/nitr-agent)

Any comments on this will be highly appreciated.
## [7][Best way to cross compile a Golang app onto a raspberry pi?](https://www.reddit.com/r/golang/comments/gszi0d/best_way_to_cross_compile_a_golang_app_onto_a/)
- url: https://www.reddit.com/r/golang/comments/gszi0d/best_way_to_cross_compile_a_golang_app_onto_a/
---
Just to be clear, I've looked at typical options like gox and xgo which do cross compiling but ultimately I'm stuck with the issue my app has a few difficult caveats to deal with namely I need libnfc and libasound2 as there are some C libraries being used by the app's dependencies to make the app work.

I've already got the app building and running on a Raspberry Pi 4. This is mainly just an exercise of curiosity.

Is there a good way of cross compiling with dependencies in place? I'm writing my code on a Mac and I can do a little testing of it there but for simplicity I'd love to compile it their as well and be able to just push the build to a raspberry pi. Originally looked at just making making a docker image so I could boot up the container and compile it there, seems when I try though the build fails and I don't necessarily know a lot about compiling across different architectures. I'm still very much a beginner when it comes to using Go as it is.
## [8][Generate a REST API from a PostgreSQL database with pagination, sorting, filtering, and JWT-based authentication](https://www.reddit.com/r/golang/comments/gsl7sp/generate_a_rest_api_from_a_postgresql_database/)
- url: https://eatonphil.github.io/dbcore/
---

## [9][How to cut elements from []byte slice with minimum mallocs and time without changing it's total length?](https://www.reddit.com/r/golang/comments/gsyyl8/how_to_cut_elements_from_byte_slice_with_minimum/)
- url: https://www.reddit.com/r/golang/comments/gsyyl8/how_to_cut_elements_from_byte_slice_with_minimum/
---
Hi guys!

I'm implementing in-memory cache with golang. I store values as `[]byte` representation in a cache shards – that is `[]byte` slice with N length – to omit GC (similar to what BigCache do).

I need an opportunity to delete values from `[]byte` slice from A index to B index. 

I've tried several approaches to accomplish this task but encountered the following problems:
– long execution time
– a lof of memory allocations

The more memory I allocate to my `[]byte` slice (it's length must be fixed all the time) – the more time it takes to delete elements from it and more memory being allocated at this time. Which is quite logical, yes. 

For example, I initialize a 128mb shard – `shard := make([]byte, 128*1024*1024)` and store 3 responses as []byte:

– response `A` from `0` to `100 000` index;

– response `B` from `10000` to `15 000 000` index

– response `C` from `15 000 000` to `78 000 000` index


Now I want to delete `B` response, because it's TTL has expired. I've tried these approaches:

Approach #1. 

    valueSize := 14990000
    index := 10000
    shard = append(shard[:index], shard[index+valueSize:]...)
    shard = append(shard, make([]byte, valueSize)...)

Approach #2.

    valueSize := 14990000
    index := 10000
    for i := index; i &lt; len(shard) - valueSize; i++ {
        shard[i] = shard[i+valueSize]
    }

Approach #3.

    valueSize := 14990000
    index := 10000
    tmpBuffer := make([]byte, len(shard))
    copy(tmpBuffer[0:], shard[:index])
    copy(tmpBuffer[index:], shard[index+valueSize:])
    shard = tmpBuffer


In the #1 we have a lot of memory allocations because each `append` returns new slice. Things go really hard with big slices (even 128mb).

In the #2 there is no allocations, as far as I understand, but the more elements – the more time it'll take. With 128mb shard and ~14.29mb value it takes up to 5-7 seconds on my machine.

In the #3 we have only 1 additional allocation (right?) at the `make()`, but `copy()` takes quite much time – something around ~4-6 seconds.

None of the cases suits me: either it makes a lot of allocations and rss jumps above the roof, or it takes a lot of time.

This is why I'm looking for an optimised way to accomplish my task. I might be doing things **very** wrong at the moment, but I can't find other options. 

I will be really grateful for any kind of help! If you have any additional questions to clarify the things – I'll give a prompt reply!

Thanks in advance.
## [10][go-guardian an awesome authentication library for go](https://www.reddit.com/r/golang/comments/gt4anb/goguardian_an_awesome_authentication_library_for/)
- url: https://www.reddit.com/r/golang/comments/gt4anb/goguardian_an_awesome_authentication_library_for/
---
 Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication. 

here a beginner tutorial about it. 

[https://medium.com/@hajsanad/authentication-in-golang-using-go-guardian-b1cd47da47a0](https://medium.com/@hajsanad/authentication-in-golang-using-go-guardian-b1cd47da47a0)

github repo: 

[https://github.com/shaj13/go-guardian](https://github.com/shaj13/go-guardian)
