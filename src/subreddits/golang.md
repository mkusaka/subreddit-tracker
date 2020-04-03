# golang
## [1][Quarantine has driven me crazy - I'm writing ArnoldC as a UNIX shell, in Go](https://www.reddit.com/r/golang/comments/ftxrwx/quarantine_has_driven_me_crazy_im_writing_arnoldc/)
- url: https://www.reddit.com/r/golang/comments/ftxrwx/quarantine_has_driven_me_crazy_im_writing_arnoldc/
---
For those unfamiliar: [https://esolangs.org/wiki/ArnoldC](https://esolangs.org/wiki/ArnoldC)

&amp;#x200B;

I'm a lowly CS student with no experience in Go and no business writing any systems stuff, but what else would I do with all my free time now? Anyway, here's the repo with only one command implemented. **Any criticism is welcome, because honestly I have no idea what I'm doing.**

&amp;#x200B;

[https://github.com/John123Allison/ArnoldShell](https://github.com/John123Allison/ArnoldShell)
## [2][Help setting up ECHO framework](https://www.reddit.com/r/golang/comments/fu80zv/help_setting_up_echo_framework/)
- url: https://www.reddit.com/r/golang/comments/fu80zv/help_setting_up_echo_framework/
---
I'm using echo's quick guide that you can follow along to see where I'm at([labstack.com/guide](https://labstack.com/guide)) here.  
So far I've ran "go get -u github.com/labstack/echo/..." on my terminal. Then, I CD'd into the directory of my project and ran "touch server.go". I opened this directory on VS Code, and copied and pasted the server file code into the server.go file. I saved it and was prompted with an error that says the following:

    server.go:6:2: cannot find package "github.com/labstack/echo/v4" in any of:
    	/usr/local/go/src/github.com/labstack/echo/v4 (from $GOROOT)
    	/Users/[my_comp_name]/go/src/github.com/labstack/echo/v4 (from $GOPATH)

I've tried going back to my terminal and run "go get ./..." to no avail. What am I doing wrong here?
## [3][A tiny Go library which can help you to keep your configurations structured. Currently you can init fields from: environment variables, flags, files(json and yaml) and 'default' tag.](https://www.reddit.com/r/golang/comments/fu7s9o/a_tiny_go_library_which_can_help_you_to_keep_your/)
- url: https://github.com/BoRuDar/configuration
---

## [4][Announcing Go-TinyTime, Go-TinyDate's Sister Package](https://www.reddit.com/r/golang/comments/ftlmxo/announcing_gotinytime_gotinydates_sister_package/)
- url: https://qvault.io/2020/04/02/announcing-go-tinytime-go-tinydates-sister-package/
---

## [5][Open source tool to auto-generate API tests using text/template](https://www.reddit.com/r/golang/comments/ftxdfg/open_source_tool_to_autogenerate_api_tests_using/)
- url: https://www.reddit.com/r/golang/comments/ftxdfg/open_source_tool_to_autogenerate_api_tests_using/
---
Hi all! I'm a software engineer in quality + automation tooling at Intuit, and today my team would like to announce our newly open source tool Replay Zero, a developer productivity tool to assist in authoring functional API tests. As the title says we use the standard library's templating engine to generate our output and have had a great experience with it so far.

Read my Medium post given a bit of backstory

[https://medium.com/@wtait1/introducing-replay-zero-77c9e5a54466](https://medium.com/@wtait1/introducing-replay-zero-77c9e5a54466)

  
Or check out the repo directly here  
[https://github.com/intuit/ReplayZero](https://github.com/intuit/ReplayZero)
## [6][Generating Go code in Kubebuilder style](https://www.reddit.com/r/golang/comments/ftnxak/generating_go_code_in_kubebuilder_style/)
- url: https://banzaicloud.com/blog/generating-go-code/
---

## [7][alash3al/go-pubsub: An in-process simple publish/subscribe library for golang](https://www.reddit.com/r/golang/comments/fu59zy/alash3algopubsub_an_inprocess_simple/)
- url: https://github.com/alash3al/go-pubsub
---

## [8][The best kafka tool to debug brokers written in Go !](https://www.reddit.com/r/golang/comments/ftgxel/the_best_kafka_tool_to_debug_brokers_written_in_go/)
- url: https://github.com/birdayz/kaf
---

## [9][Golang on Mips](https://www.reddit.com/r/golang/comments/ftpy4s/golang_on_mips/)
- url: https://www.reddit.com/r/golang/comments/ftpy4s/golang_on_mips/
---
Can Golang run on this hardware? I tried to compile with GOOS=linux GOARCH=mips, but when I run compiled app, I get syntax error.

`root@dm500hd:~# uname -a
Linux dm500hd 3.2-dm500hd #1 SMP Sun Oct 29 08:22:40 CET 2017 mips GNU/Linux
root@dm500hd:~# cat /proc/cpuinfo
system type             : BCM7413B1 STB platform
processor               : 0
cpu model               : Brcm4380 V4.4  FPU V0.1
BogoMIPS                : 403.45
cpu MHz                 : 405.016
wait instruction        : yes
microsecond timers      : yes
tlb_entries             : 32
extra interrupt vector  : yes
hardware watchpoint     : no
ASEs implemented        : mips16
shadow register sets    : 1
kscratch registers      : 0
core                    : 0
VCED exceptions         : not available
VCEI exceptions         : not available

processor               : 1
cpu model               : Brcm4380 V4.4  FPU V0.1
BogoMIPS                : 403.45
cpu MHz                 : 405.016
wait instruction        : yes
microsecond timers      : yes
tlb_entries             : 32
extra interrupt vector  : yes
hardware watchpoint     : no
ASEs implemented        : mips16
shadow register sets    : 1
kscratch registers      : 0
core                    : 0
VCED exceptions         : not available
VCEI exceptions         : not available

root@dm500hd:~# `
## [10][Trying to make changes to an interface](https://www.reddit.com/r/golang/comments/ftuult/trying_to_make_changes_to_an_interface/)
- url: https://www.reddit.com/r/golang/comments/ftuult/trying_to_make_changes_to_an_interface/
---
I have a variable  named \`\`\`result\`\`\` that is an interface but when I use 

    fmt.Printf("(%v, %T)\n", result, result)
    (&amp;[0x6d 0x00000 0x01f401f5f 0x03], *[]string)

I am trying to change the \`\`\`0x03\`\`\` to a different value but  \`\`\`result\[3\] = "3faf"\`\`\` doesn't work.

Using golang, the variable type is [this](https://prnt.sc/rrv3ae).
