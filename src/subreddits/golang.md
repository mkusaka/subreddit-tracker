# golang
## [1][Squzy - is a high-performance open-source monitoring system written in Golang with Bazel and love. Help you monitoring critical api(http, grpc) and can be used for site monitoring.](https://www.reddit.com/r/golang/comments/et37ad/squzy_is_a_highperformance_opensource_monitoring/)
- url: https://github.com/squzy/squzy
---

## [2][termenv - Advanced ANSI style &amp; color support for your terminal applications](https://www.reddit.com/r/golang/comments/ess2vz/termenv_advanced_ansi_style_color_support_for/)
- url: https://github.com/muesli/termenv
---

## [3][Go and Robotics/AI](https://www.reddit.com/r/golang/comments/eta159/go_and_roboticsai/)
- url: https://www.reddit.com/r/golang/comments/eta159/go_and_roboticsai/
---
I listened a podcast in Talk Python channel and the topic was about robotics and Python. The speaker mentioned about they have to use C++ for complex things because of performance of Python(Because when it comes to Robots, reaction duration must be minimal). So, scale of Python in Robotics is not that large.

I think, Golang has best of both, access to microcontrollers without sacrificing coding productivity(unlike C++) and AI libraries without big performance penalty(unlike Python).

I'm entry level Go developer and learning it for gRPC but later I have plans for AI and Robotics. So, I wanted to know what you think about this topic? Is it trendy and is it possible to get best result with Go? Should I invest in Go for this area?

Podcast link: [https://podcasts.apple.com/tr/podcast/talk-python-to-me-python-conversations-for-passionate/id979020229?i=1000452355489](https://podcasts.apple.com/tr/podcast/talk-python-to-me-python-conversations-for-passionate/id979020229?i=1000452355489)
## [4][How to setup "go get" properly?](https://www.reddit.com/r/golang/comments/et6r2i/how_to_setup_go_get_properly/)
- url: https://www.reddit.com/r/golang/comments/et6r2i/how_to_setup_go_get_properly/
---
I followed few tutorials and all of them helped me to get go installed and setup env variable but am having trouble getting "get" running on my linux machine.

Can anyone help me or point to an article which talks about setting up go properly and get the "get" working.
## [5][java vs go question](https://www.reddit.com/r/golang/comments/et4t7o/java_vs_go_question/)
- url: https://www.reddit.com/r/golang/comments/et4t7o/java_vs_go_question/
---
Say you had several services taking in millions of requests a day, why would you write those services in go vs java. I have read through a few blog posts online but I am interested to hear what the community here has to say...
## [6][Go modules are now available in gomobile!](https://www.reddit.com/r/golang/comments/esu9bd/go_modules_are_now_available_in_gomobile/)
- url: https://github.com/golang/go/issues/27234#issuecomment-577481562
---

## [7][Announcing cleanup: Remove gone Git branches with ease.](https://www.reddit.com/r/golang/comments/et1i5e/announcing_cleanup_remove_gone_git_branches_with/)
- url: https://github.com/dominikbraun/cleanup
---

## [8][Instrumenting Go apps with uprobes and eBPF](https://www.reddit.com/r/golang/comments/esvgd9/instrumenting_go_apps_with_uprobes_and_ebpf/)
- url: https://www.reddit.com/r/golang/comments/esvgd9/instrumenting_go_apps_with_uprobes_and_ebpf/
---
Hi,

I've put together a write up on [instrumenting userland apps](https://sematext.com/blog/ebpf-userland-apps/) with uprobes and eBPF. There is a dedicated section about how to leverage eBPF superpowers to implant uprobe hook points in Go processes. I also built a small tool (in Go!) to trace `http.Get` calls with BCC. You can find it [here](https://github.com/sematext/uprobe-http-tracer).
## [9][Specifying DNS server for lookup](https://www.reddit.com/r/golang/comments/esus2j/specifying_dns_server_for_lookup/)
- url: https://www.reddit.com/r/golang/comments/esus2j/specifying_dns_server_for_lookup/
---
Is thee a way to specify which DNS server to use for a name lookup?

&amp;nbsp;

Looking at https://golang.org/pkg/net/#LookupHost seems it will use only the local resolver

    LookupHost looks up the given host using the local resolver. It returns a slice 
    of that host's addresses. 

&amp;nbsp;

Also earlier on that link

     It can use a pure Go resolver that sends DNS requests directly to 
     the servers listed in /etc/resolv.conf,

&amp;nbsp;

How could one do a lookup against arbitrary servers like one can do with dig?

    dig @8.8.8.8 google.com
## [10][GitHub - xandout/soxy: Golang WebSocket to TCP proxy bridge (like kubectl port-forward)](https://www.reddit.com/r/golang/comments/eszhhu/github_xandoutsoxy_golang_websocket_to_tcp_proxy/)
- url: https://github.com/xandout/soxy
---

