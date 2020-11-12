# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][How Riot Games uses Golang for game development and operations](https://www.reddit.com/r/golang/comments/jsdulp/how_riot_games_uses_golang_for_game_development/)
- url: https://technology.riotgames.com/news/leveraging-golang-game-development-and-operations
---

## [3][25 Lines of Code Face Detection in Go](https://www.reddit.com/r/golang/comments/jslv32/25_lines_of_code_face_detection_in_go/)
- url: https://www.reddit.com/r/golang/comments/jslv32/25_lines_of_code_face_detection_in_go/
---
`Pion/mediadevices`, is a cross platform library which takes input from arbitrary media input devices (i.e. camera, microphone, screen share), with a simple and elegant API to create media pipelines.

* Want to create a security system, the library handles viewing from tons of devices at once
* Want to send your screen peer to peer, super easy!
* Want to create a robot to track you around the house and remind you to take the dog out, slightly less easy, but getting the camera and microphone and processing it in an elegant pipeline, super duper easy! 

# MediaDevices

`mediadevices` provides access to media input devices like cameras, microphones, and screen capture. It can also be used to encode your video/audio stream to various codec selections. `mediadevices` abstracts away the complexities of interacting with things like hardware and codecs allowing you to focus on building applications, interacting only with an amazingly **simple**, **easy**, and **elegant** API!

Github: [https://github.com/pion/mediadevices](https://github.com/pion/mediadevices)

[Face detection demo](https://i.redd.it/u0v95emkupy51.gif)

Examples:

* [Webrtc](https://github.com/pion/mediadevices/blob/master/examples/webrtc) \- Use Webrtc to create a realtime peer-to-peer video call
* [Face Detection](https://github.com/pion/mediadevices/blob/master/examples/facedetection) \- Use a machine learning algorithm to detect faces in a camera stream
* [RTP Stream](https://github.com/pion/mediadevices/blob/master/examples/rtp) \- Capture camera stream, encode it in H264/VP8/VP9, and send it to a RTP server
* [HTTP Broadcast](https://github.com/pion/mediadevices/blob/master/examples/http) \- Broadcast camera stream through HTTP with MJPEG
* [Archive](https://github.com/pion/mediadevices/blob/master/examples/archive) \- Archive H264 encoded video stream from a camera
## [4][k6 v0.29.0 is out! 🎊🎉🥳](https://www.reddit.com/r/golang/comments/jssvos/k6_v0290_is_out/)
- url: https://www.reddit.com/r/golang/comments/jssvos/k6_v0290_is_out/
---
We are happy and excited to announce that k6 v0.29.0 is released with new features and improvements. k6 is a modern open-source performance and load-testing tool, written in Go and scriptable in JavaScript. 🎊🎉🥳 This release contains the work of over 10 contributors split over more than 100 commits.

- Initial support for gRPC
- New options for configuring DNS resolution
- Support for Go extensions
- Support for setting local IPs, potentially from multiple NICs
- New option for blocking hostnames
- UX improvements
- Lots of bugfixes

This You can read more about it in the [release notes](https://github.com/loadimpact/k6/blob/master/release%20notes/v0.29.0.md). We'd be happy to hear your feedback about it.
## [5][[YOUTUBE] Fuzzing Go package in 5 min using go-fuzz &amp; libfuzzer](https://www.reddit.com/r/golang/comments/jsswu3/youtube_fuzzing_go_package_in_5_min_using_gofuzz/)
- url: https://www.youtube.com/watch?v=EsSebOAD5yw
---

## [6][leetcode-helper helps you debug your leetcode problems code. it just do two things: 1. convert leetcode string input to golang data structure which leetcode used . 2. display the data and structures.](https://www.reddit.com/r/golang/comments/jsu3yq/leetcodehelper_helps_you_debug_your_leetcode/)
- url: https://github.com/EchoUtopia/leetcode-helper
---

## [7][Best way to deploy a worker app to AWS](https://www.reddit.com/r/golang/comments/jssrzd/best_way_to_deploy_a_worker_app_to_aws/)
- url: https://www.reddit.com/r/golang/comments/jssrzd/best_way_to_deploy_a_worker_app_to_aws/
---
I have an app that subscribes to data through websocket and sends the incoming messages to several SQS queues. It works fine as a compiled binary, but deploying it to EC2 is a manual process for me - I have to copy the binary to the instance and run it manually. If it exits, I need to restart it or do some bash scripting to make it restart.

I was thinking about deploying it to ECS as a docker container, but something about dockerizing a binary feels wrong. Is there a better way?
## [8][haunt98/changeloguru: Tool to generate CHANGELOG.md from Conventional Commits.](https://www.reddit.com/r/golang/comments/jsawxc/haunt98changeloguru_tool_to_generate_changelogmd/)
- url: https://www.reddit.com/r/golang/comments/jsawxc/haunt98changeloguru_tool_to_generate_changelogmd/
---
[https://github.com/haunt98/changeloguru](https://github.com/haunt98/changeloguru)

This is my side project. Because my company project use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/), so I create a tool to auto generate CHANGELOG.md from them :D
## [9][Wombat: Cross Platform gRPC Client. Developed in Go.](https://www.reddit.com/r/golang/comments/jrudsh/wombat_cross_platform_grpc_client_developed_in_go/)
- url: https://i.redd.it/w03bdqwnghy51.jpg
---

## [10][EFK on a GCP VM instance](https://www.reddit.com/r/golang/comments/jsq1pa/efk_on_a_gcp_vm_instance/)
- url: https://www.reddit.com/r/golang/comments/jsq1pa/efk_on_a_gcp_vm_instance/
---
I am new to GCP. I have a GCE VM instance running (e2-small (2 vCPUs, 2 GB memory)) with an Ubuntu boot disk, which is hosting a golang API server. I want to install EFK onto that vm in order to capture and store the logs generated by the golang server (efk integrated through the elastic-go library in the code itself). Is it possible to install EFK on that VM instance/How to install EFK on that VM instance? What other options do I have?
## [11][GoThrough.dev - Searchable index of exported members in the standard library](https://www.reddit.com/r/golang/comments/jsaxda/gothroughdev_searchable_index_of_exported_members/)
- url: https://gothrough.dev/
---

