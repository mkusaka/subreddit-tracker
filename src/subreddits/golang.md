# golang
## [1][Modern Jenkins written in Go](https://www.reddit.com/r/golang/comments/iijvsg/modern_jenkins_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/iijvsg/modern_jenkins_written_in_go/
---
Recently, I used Jenkins for creating a Canary infrastructure where it needs to be able to connect 2 or more different nodes. I found Jenkins to be very flexible and meet the requirements, but I also found that Jenkins seems to be outdated. Jenkins feels really slow, the UI doesn't work very well for complex scenarios (e.g. nested parallel pipelines), etc.

Is there any Jenkins alternatives out there? I think Jenkins is different from any other CI/CDs that I know of due to:

1. Self hosted
2. Scalable, master-slave architecture makes it really easy to scale
3. Easy to add slaves, only needs SSH access and Java runtime
4. Very flexible pipeline, using Groovy for pipeline scripting makes it easy to create a very complex pipeline
5. Durable, auto recover steps can survive from network connection lost
6. Useful UI to show pipelines (mainly referring to Blue Ocean)
7. Plugin system

I'm thinking to reinvent the wheels, create a Jenkins alternative in Go. I think using Go has a lot of benefits:

1. Slaves don't need to have Java runtime to run the Jenkins agent
2. Static binary means it's very easy to install and no dependencies needed
3. No slow start up time from JVM
4. Pipeline scripting can be done in Go, thanks to [Yaegi](https://github.com/containous/yaegi)
5. Plugins can also be written in Go, again thanks to [Yaegi](https://github.com/containous/yaegi). This is similar to how [Traefik](https://containo.us/traefik/) uses Yaegi themselves

Is this worth it? Am I missing important pieces? If this sounds good, I'll start the project and be very welcome to contributors!

&amp;#x200B;

Note: I've explored many CI/CDs but they're missing some of Jenkins features that I listed above.

&amp;#x200B;

Note 2: Many people have suggested many alternatives. But, I found that none of them supports powerful scripting language like Groovy in Jenkins. I used Groovy to help me queue a task to 2 different nodes, and synchronize the 2 nodes so that they can start and finish at the expected time.
## [2][Discord server for Golang Developers](https://www.reddit.com/r/golang/comments/iiq3f4/discord_server_for_golang_developers/)
- url: https://www.reddit.com/r/golang/comments/iiq3f4/discord_server_for_golang_developers/
---
Hello Gophers, 

Over the last week our new Golang friends discord community has grown to 75 members, many of whom have since started group Github projects and shared their tech learning wins. Join us, all skill-levels are welcome!

Join us. We are very welcoming to newcomers and have members of all skill levels to answer your questions. We value all opinions, suggestions, rants, and anything else!

You can play an integral role in shaping our small tight-knit community. Get to know our members in a safe space where all discussions are welcome. 

Invite link below.
https://discord.gg/VmrBgg9
## [3][Testing Database Transactions in Go](https://www.reddit.com/r/golang/comments/iis3gm/testing_database_transactions_in_go/)
- url: https://marvinblum.de/blog/testing-database-transactions-in-go-jEaOGXravM
---

## [4][Windows compile with libusb](https://www.reddit.com/r/golang/comments/iiqziw/windows_compile_with_libusb/)
- url: https://www.reddit.com/r/golang/comments/iiqziw/windows_compile_with_libusb/
---
So I've made a generic 'driver' for a Logitech Flight Panel - a box of buttons for Flight-Sims.

This make it emulate a keyboard, so can be used in any flight-sim (or any other game), and not just the select few that Logitech provide plug-ins for.

It works wonderfully in Linux, no problem. But I just can't get it to compile under Windows.

Yes, I have installed pkg\_config, I've unpacked the libusb tar file to where pkg\_config finds it no problem. Yet trying to compile under windows it reports cgo invalid flags. Oh, and I have MinGW with the proper C compiler, all is fine there.

My little project resides at: [https://github.com/rDybing/switchPanel](https://github.com/rDybing/switchPanel)

If anyone could help me get this compiled under windows, that'd be fantastic!
## [5][Simplified HTTP requests and response unmarshaling](https://www.reddit.com/r/golang/comments/iiqyol/simplified_http_requests_and_response_unmarshaling/)
- url: https://github.com/Navid2zp/easyreq
---

## [6][Help me make SpaceCloud, an open source Firebase alternative, better! Starting a bi-weekly community call for the project!](https://www.reddit.com/r/golang/comments/iicfoq/help_me_make_spacecloud_an_open_source_firebase/)
- url: https://www.reddit.com/r/golang/comments/iicfoq/help_me_make_spacecloud_an_open_source_firebase/
---
Its been almost two years since I started working on SpaceCloud. Its been an incredible journey. Could never have imagined that the tool which was initially meant for internal use only has become a full fledged open source project!!

I've decided to adopt a more open governance model for SpaceCloud and open doors for active community participation. I really want to _imbibe your experiences and recommended practices in SpaceCloud's architecture_.

I understand it will be a long and tedious process to become truly open. I want to take the first step by inviting you to the Community call I'm starting. The next call is tomorrow at 6pm IST (GMT+5:30).

The agenda for the call: https://doc.clickup.com/d/h/3e4vc-186/ec5ebbf7fb270ba/3e4vc-291

Meeting details: https://zoom.us/j/91600633833

GitHub Repo: https://github.com/spaceuptech/space-cloud

Discord: https://discord.gg/RkGjW93

You can always drop by to say hello. A shoutout or GitHub star would work as well! 🙈

P.S.- I know i should've given more explanation of what the call or SpaceCloud is about. Bare with me. It's my first time doing this.
## [7][New Case Studies About Google’s Use of Go](https://www.reddit.com/r/golang/comments/ii300l/new_case_studies_about_googles_use_of_go/)
- url: https://opensource.googleblog.com/2020/08/new-case-studies-about-googles-use-of-go.html
---

## [8][Thanks for all your support!! The Community call is going to kick off in 15 mins!](https://www.reddit.com/r/golang/comments/iirdly/thanks_for_all_your_support_the_community_call_is/)
- url: /r/golang/comments/iicfoq/help_me_make_spacecloud_an_open_source_firebase/
---

## [9][A new Go API for Protocol Buffers](https://www.reddit.com/r/golang/comments/iioh28/a_new_go_api_for_protocol_buffers/)
- url: https://blog.golang.org/protobuf-apiv2
---

## [10][GLab is an open source Gitlab Cli tool written in Go to help work seamlessly with Gitlab from the command line](https://www.reddit.com/r/golang/comments/ii4yzm/glab_is_an_open_source_gitlab_cli_tool_written_in/)
- url: https://github.com/profclems/glab
---

