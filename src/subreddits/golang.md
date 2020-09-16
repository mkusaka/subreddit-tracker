# golang
## [1][gopls v0.5.0](https://www.reddit.com/r/golang/comments/itq7ih/gopls_v050/)
- url: https://github.com/golang/tools/releases/tag/gopls%2Fv0.5.0
---

## [2][Pebble: A RocksDB Inspired Key-Value Store Written in Go](https://www.reddit.com/r/golang/comments/itgy27/pebble_a_rocksdb_inspired_keyvalue_store_written/)
- url: https://www.cockroachlabs.com/blog/pebble-rocksdb-kv-store/
---

## [3][My favorite tattoo.](https://www.reddit.com/r/golang/comments/itllt2/my_favorite_tattoo/)
- url: https://i.redd.it/7v95dyu3sen51.jpg
---

## [4][A cool little Context trick for Gophers](https://www.reddit.com/r/golang/comments/itfjs2/a_cool_little_context_trick_for_gophers/)
- url: https://v.redd.it/oy08hvk33dn51
---

## [5][Why is the Go CLI help so weird?](https://www.reddit.com/r/golang/comments/ittitm/why_is_the_go_cli_help_so_weird/)
- url: https://www.reddit.com/r/golang/comments/ittitm/why_is_the_go_cli_help_so_weird/
---
This is just a short question, because I stumble upon this sooo often: Why did the Go designers decide to make `--help` basically useless?

    $ go clean --help
    usage: go clean [clean flags] [build flags] [packages]
    Run 'go help clean' for details.

    $ go help clean
    usage: go clean [clean flags] [build flags] [packages]

    Clean removes object files from package source directories.
    The go command builds most objects in a temporary directory,
    so go clean is mainly concerned with object files left by other
    tools or by manual invocations of go build.
    .....

This seems to passive aggressive and intentionally unhelpful. It's like Go knows "oh yeah, the user wants to see how this command works", but instead of just providing the help, it tells me "nope, you asked the wrong way, try again!"

The first version is useless to me, because `go help clean flags` does not work, neither does `go help build flags`, so I *must* run `go help clean` first to find out how to even get help for the build flags (`go help build`).

What's the reasoning behind this? Am I just holding it wrong?
## [6][idm - Golang wrapper for Internet Download Manager (IDM) CLI](https://www.reddit.com/r/golang/comments/itvedt/idm_golang_wrapper_for_internet_download_manager/)
- url: https://github.com/Navid2zp/idm
---

## [7][docx: Simple pure Go (golang) library for creating DOCX file](https://www.reddit.com/r/golang/comments/itv2qw/docx_simple_pure_go_golang_library_for_creating/)
- url: https://github.com/gingfrederik/docx
---

## [8][[Q] Communication in a microservice architecture](https://www.reddit.com/r/golang/comments/ituhql/q_communication_in_a_microservice_architecture/)
- url: https://www.reddit.com/r/golang/comments/ituhql/q_communication_in_a_microservice_architecture/
---
Howdy! Newbie learning microservices here. Currently I'm trying to use `gRPC`, but `grpc-web` is really PITA to use. I'm kind of looking for alternatives, what else do you guys use, messaging like RabbitMQ/Kafka or simply REST?

Please enlighten me.

---

Stack: Go, React/TS
## [9][When the Go garbage collector will panic over bad pointer values](https://www.reddit.com/r/golang/comments/it85va/when_the_go_garbage_collector_will_panic_over_bad/)
- url: https://utcc.utoronto.ca/~cks/space/blog/programming/GoGCBadPointerPanics
---

## [10][Authentication for REST APIs in Go](https://www.reddit.com/r/golang/comments/itttn4/authentication_for_rest_apis_in_go/)
- url: https://www.maragu.dk/blog/authentication-for-rest-apis-in-go/
---

