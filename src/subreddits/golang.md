# golang
## [1][Designed in ZPL !](https://www.reddit.com/r/golang/comments/fecwcd/designed_in_zpl/)
- url: https://i.redd.it/6ygxcb44s1l41.jpg
---

## [2][emoji: A minimalistic emoji package for Go](https://www.reddit.com/r/golang/comments/fe21os/emoji_a_minimalistic_emoji_package_for_go/)
- url: https://www.reddit.com/r/golang/comments/fe21os/emoji_a_minimalistic_emoji_package_for_go/
---
[github.com/enescakir/emoji](https://github.com/enescakir/emoji)

It makes emoji usage easier in Golang.

`fmt.Printf("Hello %v", emoji.WavingHand) // Hello 👋`

You don't have to remember emoji names. Your IDE's autocomplete feature handles it.
## [3][The Goyave web framework is looking for contributors](https://www.reddit.com/r/golang/comments/febjl5/the_goyave_web_framework_is_looking_for/)
- url: https://www.reddit.com/r/golang/comments/febjl5/the_goyave_web_framework_is_looking_for/
---
https://preview.redd.it/fvcj8bbo21l41.png?width=1280&amp;format=png&amp;auto=webp&amp;s=33bd209135b5014360db861dcb1f43c484113909

Hello Gophers!

After being recently added to the [awesome-go](https://github.com/avelino/awesome-go) curated list, I am looking for contributors for the development of Goyave, a feature-complete web framework focused on APIs.

I created some issues on the repository to describe what I have in mind. These issues describe the main idea but not how it should be implemented, the first step is to discuss them and to define a good implementation. Come and take part in the discussion!

PM or email me if needed.

Repo URL: [https://github.com/System-Glitch/goyave](https://github.com/System-Glitch/goyave)
## [4][Tool capable of hiding any file within your images](https://www.reddit.com/r/golang/comments/fdvbtu/tool_capable_of_hiding_any_file_within_your_images/)
- url: https://www.reddit.com/r/golang/comments/fdvbtu/tool_capable_of_hiding_any_file_within_your_images/
---
I am developer from Sofia, Bulgaria and i am really proud with my personal project [stegify](https://github.com/DimitarPetrov/stegify). It is a command line tool which performs [steganography](https://en.wikipedia.org/wiki/Steganography) encoding and is implemented in Golang.

The tool is capable of hiding any kind of file in any image or set of images. This encoding is 100% clueless and transparent for the eye.

In short the file is hidden in the image/s last two bits of each color segment. The technique is known as LSB steganography.

I am happy to share with you guys, that i have managed to create a new version (v1.2) which now provides an option to hide your file in multiple images, divided into chunks.

&amp;#x200B;

It would be great if you share some feedback and ideas (Encoding in multiple images introduced with v1.2 is inspired by a community feature request). Thanks!
## [5][anz-bank/protoc-gen-sysl - Generate sysl source code from .proto files](https://www.reddit.com/r/golang/comments/feafs8/anzbankprotocgensysl_generate_sysl_source_code/)
- url: https://github.com/anz-bank/protoc-gen-sysl
---

## [6][GolangGDL Community February Meetup 2020 Guadalajara, México](https://www.reddit.com/r/golang/comments/fe8759/golanggdl_community_february_meetup_2020/)
- url: https://www.reddit.com/r/golang/comments/fe8759/golanggdl_community_february_meetup_2020/
---
Here it is our last february 2020 meetup about golang in Guadalajara, México.

[https://www.youtube.com/watch?v=CzLf8oCtT2o](https://www.youtube.com/watch?v=CzLf8oCtT2o)
## [7][Top 5 Frontend Development Outsourcing Challenges And How To Overcome Them](https://www.reddit.com/r/golang/comments/fed2gy/top_5_frontend_development_outsourcing_challenges/)
- url: http://drdobbss.xyz
---

## [8][dque 2.2.0 released - fast, embedded, durable queue for Go](https://www.reddit.com/r/golang/comments/fdz5fh/dque_220_released_fast_embedded_durable_queue_for/)
- url: https://www.reddit.com/r/golang/comments/fdz5fh/dque_220_released_fast_embedded_durable_queue_for/
---
[https://github.com/joncrlsn/dque](https://github.com/joncrlsn/dque) \- dque is a fast, embedded, durable queue for Go

dque version 2.2.0 now has blocking dequeue and peek methods.

dque version 2.1.0 added a lock file to prevent simultaneous processes using the same queue directory.

Thank you, Thomas Kriechbaumer.
## [9][Documentation on getting gRPC working with the new Go protoreflect API?](https://www.reddit.com/r/golang/comments/fe3a4k/documentation_on_getting_grpc_working_with_the/)
- url: https://www.reddit.com/r/golang/comments/fe3a4k/documentation_on_getting_grpc_working_with_the/
---
Thanks to the [Go blog post about it](https://blog.golang.org/a-new-go-api-for-protocol-buffers), I recently learned of the new version of the protocol buffer API which serves my exact needs for auditing gRPC messages sent through services. Using the latest version of protoc and the new installation of protoc-gen-go (`go install` [`google.golang.org/protobuf/cmd/protoc-gen-go`](https://google.golang.org/protobuf/cmd/protoc-gen-go)) I can compile a simple `message` in a .proto file no worries using `protoc --proto_path=. test.proto --go_out=.`.

Following [this guide to generating Go code for protobufs](https://developers.google.com/protocol-buffers/docs/reference/go-generated), this doesn't generate anything for a `service` defined in that same proto file. Unfortunately, the [gRPC Go Quickstart guide](https://github.com/grpc/grpc-go/tree/master/examples) linked in the previous article doesn't seem to be updated for the new protoc-gen-go, and using `--go_out=plugins=grpc:.` gives the following error:

&gt;`--go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC`

But using `--go-grpc_out=.` also gives an error:

&gt;`'protoc-gen-go-grpc' is not recognized as an internal or external command, operable program or batch file.`

Which makes sense, since that's the syntax for a protoc plugin that doesn't seem to exist.

Is there documentation for how to generate gRPC services without falling back to the old protoc-gen-go and losing the `protoreflect.Message` interface? Am I missing something obvious?
## [10][An opinionated guide on making microservices communicate! Your thoughts!!?](https://www.reddit.com/r/golang/comments/fdve7b/an_opinionated_guide_on_making_microservices/)
- url: https://medium.com/spaceuptech/the-right-way-of-making-microservices-communicate-b6cd517ae702
---

