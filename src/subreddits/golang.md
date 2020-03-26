# golang
## [1][How much performance headroom is left / longer compile times](https://www.reddit.com/r/golang/comments/fp9mrk/how_much_performance_headroom_is_left_longer/)
- url: https://www.reddit.com/r/golang/comments/fp9mrk/how_much_performance_headroom_is_left_longer/
---
Hi all – As you know, the Go compiler is incredibly fast, and generated application code is decently fast. How much do you think application performance could be improved with longer compile times?

Say there was an option for longer, deeper compiles, for contexts where instant compiles were not needed. For example, for the release build of an application I'd want the best possible runtime performance – I'd be fine letting a compiler run over a weekend. Do you think there is still a lot of application performance headroom to be realized with longer build times?

The compiler is somewhat threadbare when it comes to optimization options. It lacks the suite of optimizations that gcc, clang, Visual Studio, and the Intel compiler offer for C/C++ applications. Where I see possible headroom is in targeting modern CPU instructions and vectorization. The last time I looked at the compiler code base I saw a bunch of old assembly targeting old CPUs – is it still not accounting for Haswell+ CPUs? Go support for modern assembly *by the application developer* is excellent, but I'm not sure if the Go compiler itself uses modern instructions in its generated code. Thoughts?

Thanks.
## [2][grailbio/bigslice: A serverless cluster computing system for the Go programming language](https://www.reddit.com/r/golang/comments/fp8424/grailbiobigslice_a_serverless_cluster_computing/)
- url: https://github.com/grailbio/bigslice
---

## [3][pion/ion, self-hosted conferencing software with single command deploy written entirely in Go!](https://www.reddit.com/r/golang/comments/fouryt/pionion_selfhosted_conferencing_software_with/)
- url: https://github.com/pion/ion
---

## [4][Custom JSON unmarshal to fingerprint objects](https://www.reddit.com/r/golang/comments/fp5dm5/custom_json_unmarshal_to_fingerprint_objects/)
- url: https://klotzandrew.com/blog/object-fingerprinting-for-efficient-data-ingestion
---

## [5][Gold: Reinforcement Learning in Go](https://www.reddit.com/r/golang/comments/fpa515/gold_reinforcement_learning_in_go/)
- url: https://github.com/aunum/gold
---

## [6][Go, the Go Community, and the Pandemic](https://www.reddit.com/r/golang/comments/fowx7s/go_the_go_community_and_the_pandemic/)
- url: https://blog.golang.org/pandemic
---

## [7][xcgo has what gophers crave (including snap)](https://www.reddit.com/r/golang/comments/fp4h9e/xcgo_has_what_gophers_crave_including_snap/)
- url: https://www.reddit.com/r/golang/comments/fp4h9e/xcgo_has_what_gophers_crave_including_snap/
---
[neilotoole/xcgo](https://github.com/neilotoole/xcgo) docker image will cross-compile and distribute (via `goreleaser`) your CGo app, even a `snap`. See the companion [neilotoole/sqlitr](https://github.com/neilotoole/sqlitr) project for an example of using `goreleaser` to publish to `brew`, `scoop`, `deb`, `rpm`, and also... `snap`.
## [8][Building Distributed Services with Go is available in beta. Learn how to build a distributed service from scratch.](https://www.reddit.com/r/golang/comments/fotz2h/building_distributed_services_with_go_is/)
- url: https://pragprog.com/book/tjgo/distributed-services-with-go
---

## [9][Using environment variables in your next GoLang project](https://www.reddit.com/r/golang/comments/fp9y53/using_environment_variables_in_your_next_golang/)
- url: https://www.hackdoor.io/articles/lMN5AwN1/using-environment-variables-in-your-next-golang-project
---

## [10][Quickly Find Rust Program Bottlenecks Online Using a Go Tool](https://www.reddit.com/r/golang/comments/fp3e2h/quickly_find_rust_program_bottlenecks_online/)
- url: https://pingcap.com/blog/quickly-find-rust-program-bottlenecks-online-using-a-go-tool/
---

