# golang
## [1][Go compiler doesn't like unused variables](https://www.reddit.com/r/golang/comments/hufqpl/go_compiler_doesnt_like_unused_variables/)
- url: https://i.redd.it/1oobvnrzwub51.jpg
---

## [2][autocert vs certmagic](https://www.reddit.com/r/golang/comments/huiffe/autocert_vs_certmagic/)
- url: https://www.reddit.com/r/golang/comments/huiffe/autocert_vs_certmagic/
---
Lately I have been experimenting with [autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert) for my reverse proxy project ([ServiceQ](https://github.com/gptankit/serviceq)). I feel it satisfies all my current requirements - auto cert generation, tls-alpn-01 challenge, auto renewals etc. It also provides me fine grained control over tls config. I have also heard a lot about [certmagic](https://github.com/caddyserver/certmagic) but never got to use it any of the projects. Is there any reason why someone should be choosing certmagic over autocert given that autocert is supported by core team and might also some day make its way into the standard library? Looking to find real feature differences.
## [3][Brian Kernighan: UNIX, C, AWK, AMPL, and Go Programming | AI Podcast #109 with Lex Fridman](https://www.reddit.com/r/golang/comments/htwr7e/brian_kernighan_unix_c_awk_ampl_and_go/)
- url: https://www.youtube.com/watch?v=O9upVbGSBFo
---

## [4][Create a Color Cli in Golang](https://www.reddit.com/r/golang/comments/hujpan/create_a_color_cli_in_golang/)
- url: https://schadokar.dev/posts/create-a-color-cli-in-golang/
---

## [5][Interested in Contribution to Go [New to Open-Source]](https://www.reddit.com/r/golang/comments/hu514j/interested_in_contribution_to_go_new_to_opensource/)
- url: https://www.reddit.com/r/golang/comments/hu514j/interested_in_contribution_to_go_new_to_opensource/
---
Hi,

I and a couple of my friends recently started with Go and are excited to contribute to Open-source. We do not have a lot of experience in the same and were hoping for some advice as to how to proceed with Open-source. I searched for some projects on Github but could not decide on how to proceed or choose a project to contribute to. Any suggestions or projects to contribute to would be helpful. Thanks.

TLDR: New to open-source, looking to contribute
## [6][HTTP/REST gateway to gRPC](https://www.reddit.com/r/golang/comments/hujf2k/httprest_gateway_to_grpc/)
- url: https://www.reddit.com/r/golang/comments/hujf2k/httprest_gateway_to_grpc/
---
Can you share some examples of code that is an endpoint / gateway between a web client and microservices.  
I do not quite understand, I have to write a service (for example, in go-gin) to receive http requests on it and somehow process them. Then make requests by gRPC to internal services.

Waiting for your advice
## [7][Checking Kafka advertised.listeners with Go](https://www.reddit.com/r/golang/comments/huidxo/checking_kafka_advertisedlisteners_with_go/)
- url: /r/apachekafka/comments/huid7o/checking_kafka_advertisedlisteners_with_go/
---

## [8][Cannot send email using Amazon SES and Golang](https://www.reddit.com/r/golang/comments/huhp1s/cannot_send_email_using_amazon_ses_and_golang/)
- url: https://www.reddit.com/r/golang/comments/huhp1s/cannot_send_email_using_amazon_ses_and_golang/
---
On my local Ubuntu machine, I'm trying [this](https://stackoverflow.com/questions/41176256/how-to-integrate-aws-sdk-ses-in-golang/41181934#41181934) example snippet on Stackoverflow with my existing SES credentials,  I get

`MissingEndpoint: 'Endpoint' configuration is required for this service.`

My go version is \`1.14.6\` and I've just got the aws package from github.  I'm wondering what is wrong here? Has SES API changed since then? Or I'm missing something?
## [9][Suggest database package to work well with graphql server](https://www.reddit.com/r/golang/comments/huh4ii/suggest_database_package_to_work_well_with/)
- url: https://www.reddit.com/r/golang/comments/huh4ii/suggest_database_package_to_work_well_with/
---
Hi, 


I build a graphql server with [gqlgen](https://github.com/99designs/gqlgen) for generate api from schema. 
For fetching data from database, I get two options: 


+ using go orm (like gorm, xorm or sqlboiler).
+ write SQL query and using. [sqlc](https://github.com/kyleconroy/sqlc) for generate Go code.



How do you fetch data from database(like mysql or postgres) and fill it to api? Anyone, whose done some project like that. Do you have any recommendation about db package to work well with grapql.
## [10][GoSlice - an experimental slicer for 3d printing](https://www.reddit.com/r/golang/comments/hu5u86/goslice_an_experimental_slicer_for_3d_printing/)
- url: https://www.reddit.com/r/golang/comments/hu5u86/goslice_an_experimental_slicer_for_3d_printing/
---
Hi,

as I am interested in both, 3d printing and Go programming, the best way to go was creating a new slicer software in Go. So...

... I would like to introduce you [GoSlice!](https://github.com/aligator/GoSlice)

If you are not into 3d printing: A slicer is basically a software which converts a 3d model into commands for the printer.

It is very experimental and there is still much to do (and there are many things, I have currently no idea how to do it, so any help would be great), but it can already slice stl files.I also got a great print of the well known 3DBenchy out of it -&gt; So the POC is done :-)

If you are interested in how such a slicer (can) work, look at the source and docs folder as I tried to comment and explain everything because there is nearly no information about the inner workings of a slicer yet. Basically the code of the existing ones, e.g. PrusaSlicer, Cura, Slic3r is nearly the only docu I could find about this topic and it's not always obvious code...
