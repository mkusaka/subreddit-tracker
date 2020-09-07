# golang
## [1][lsoc-overlay: an overlay of recording webcams and its processes on Linux](https://www.reddit.com/r/golang/comments/io2wo2/lsocoverlay_an_overlay_of_recording_webcams_and/)
- url: https://www.reddit.com/r/golang/comments/io2wo2/lsocoverlay_an_overlay_of_recording_webcams_and/
---
As summer has ended, many people are rushing back to their usual lives of online meetings and Zoom calls. This overlay was written to give some people the transparency of when their webcams are being used and by what applications.

I came up with this idea when I was fiddling with FFmpeg to create a dummy webcam from an existing one, but this caused the webcam's activity light to be constantly on. Since I can no longer rely on that indicator, I decided to write one myself.

The application uses Gtk to draw an overlay, and it is hard-coded to use Video4Linux2 (v4l2). As such, the application will only work on Linux. There is also a small one-shot CLI at `./cmd/lsof`.

GitHub link: [https://github.com/diamondburned/lsoc-overlay](https://github.com/diamondburned/lsoc-overlay)
## [2][TIL: make+copy to clone a slice is optimized in Go 1.15, which is not mentioned in 1.15 release notes.](https://www.reddit.com/r/golang/comments/inorfa/til_makecopy_to_clone_a_slice_is_optimized_in_go/)
- url: https://www.reddit.com/r/golang/comments/inorfa/til_makecopy_to_clone_a_slice_is_optimized_in_go/
---
Reads:

* https://go-review.googlesource.com/c/go/+/146719/ (by @martisch)

* https://github.com/golang/go/commit/6ed4661807b219781d1aa452b7f210e21ad1974b

* https://github.com/golang/go/issues/26252

* https://github.com/go101/go-benchmarks/tree/master/append-vs-make

* https://github.com/go101/go101/wiki/How-to-efficiently-clone-a-slice%3F
## [3][Is Go still relevant for high-performance infrastructure applications?](https://www.reddit.com/r/golang/comments/inxo4v/is_go_still_relevant_for_highperformance/)
- url: https://www.reddit.com/r/golang/comments/inxo4v/is_go_still_relevant_for_highperformance/
---
With the rise of Rust and the improvement of C++ is Go still relevant at the development of high-performance infrastructure applications?

For example, I tried to write an API gateway in Go and even optimizing it as much as I could I was reaching 1/3 of Nginx or HAProxy performance. Rust on the other side is able to archive the same level of performance or even more.

I'm still using Go for my services, but they're at the business logic side. For things like high-performance proxies? Hmm, I don't know if it's the best choice anymore. Does anyone share the same feeling?
## [4][Sloc Cloc and Code a fast code counter written in Go. Released v2.13.0 with many new features, bug fixes and 10% performance improvement thanks to Go 1.15](https://www.reddit.com/r/golang/comments/inz4id/sloc_cloc_and_code_a_fast_code_counter_written_in/)
- url: https://github.com/boyter/scc
---

## [5][Graviton Database: ZFS for Key-Value Stores](https://www.reddit.com/r/golang/comments/inoipt/graviton_database_zfs_for_keyvalue_stores/)
- url: https://github.com/deroproject/graviton
---

## [6][xs: remote shell and copy written from scratch in Go](https://www.reddit.com/r/golang/comments/io136e/xs_remote_shell_and_copy_written_from_scratch_in/)
- url: https://github.com/Russtopia/xs
---

## [7][Iterating through arrays](https://www.reddit.com/r/golang/comments/inqqm2/iterating_through_arrays/)
- url: https://youtu.be/t65Vu91jO5A
---

## [8][Building a game engine with golang!](https://www.reddit.com/r/golang/comments/inlvqi/building_a_game_engine_with_golang/)
- url: https://link.medium.com/bzx0GDeOy9
---

## [9][internet speedmeter written in golang](https://www.reddit.com/r/golang/comments/io2ute/internet_speedmeter_written_in_golang/)
- url: https://www.reddit.com/r/golang/comments/io2ute/internet_speedmeter_written_in_golang/
---
Meter to check the current bandwidth use(\*\*not a speed tester\*\*)

github: [https://github.com/amalshaji/speedmeter](https://github.com/amalshaji/speedmeter)

screenshot:

https://preview.redd.it/u5b9usx7dol51.png?width=439&amp;format=png&amp;auto=webp&amp;s=0c3a28f3a9056d6fadb6694995c59ba2e8898061
## [10][I wrote a web &amp; rpc framework, that makes developing microservices much easier!](https://www.reddit.com/r/golang/comments/io43x0/i_wrote_a_web_rpc_framework_that_makes_developing/)
- url: https://github.com/tal-tech/go-zero
---

