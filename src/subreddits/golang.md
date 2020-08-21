# golang
## [1][Gate: The extensible Minecraft proxy written in Go!](https://www.reddit.com/r/golang/comments/ideg2v/gate_the_extensible_minecraft_proxy_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/ideg2v/gate_the_extensible_minecraft_proxy_written_in_go/
---
# Today, I want to present "Gate", the extensible Minecraft proxy, to the Gophers part of the Minecraft community!

# [https://github.com/minekube/gate](https://github.com/minekube/gate)

[Every single 🌟 supports the project!](https://preview.redd.it/31xi5b36sci51.png?width=3500&amp;format=png&amp;auto=webp&amp;s=e71105ad91f79a2390a4e782c1b765052b03342f)

# Target audiences

* advanced **Minecraft networks that already (or want to) have a Go code base** for their Minecraft related workloads

Before you may ask: "*Why not use an existing proxy written in Java?*"

Because the less Java we need to maintain, the happier we are at [Minekube](https://discord.gg/6vMDqWE), since we need and work in a very fast paced and cloud-centric ecosystem with a lot of Kubernetes &amp; controllers, Protobuf &amp; GRPC, CockroachDB, Agones, Istio, Nats &amp; Stan and so forth. The ONLY Java code we must write is for paper/spigot plugins!

# What Gate does

*(for those who have never heard of a Minecraft proxy)*

**TL;DR**

* keep the player's session without disconnect to...
* move players between servers
* cross server functionalities (events such as chat/commands, send messages, ...builtin/custom session- &amp; packet handlers)

Gate presents itself as a normal Minecraft server in the player's server list, but once the player connects Gate forwards the connection to one of the actual game servers (e.g. Minecraft vanilla, paper, spigot, sponge, etc.) to play the game.

The player can be moved around the network of Minecraft servers **without** fully disconnecting, since we want the player to stay (and not want them to re-login via the server-list every time).

Therefore, Gate reads all packets sent between players (Minecraft client) and upstream servers, logs session state changes, emits different events like [Login, Disconnect, ServerConnect, Chat, Kick etc.](https://github.com/minekube/gate/blob/master/pkg/proxy/events.go) that custom plugins/code can react to.

The **advantages** for using a proxy should be clear now.

[Every single 🌟 supports the project!](https://preview.redd.it/rgnrvmi7sci51.png?width=1229&amp;format=png&amp;auto=webp&amp;s=d6723949d962b731ced7de9e4074cfa783dbab22)

# Features

* Fast
* Excellent server version support
   * 1.7 up to newest &amp; forge support
* Quick installation
   * Download the binary from the [releases](https://github.com/minekube/gate/releases)
   * `docker run -it --rm -p 25565:25565` `registry.gitlab.com/minekube/gate:latest`
* Simple API to [extend Gate](https://github.com/minekube/gate#extending-gate-with-custom-code)
* Built-in IP based rate limiter
* A detailed config with sane defaults

# ...please see more on the [GitHub repository](https://github.com/minekube/gate) to get started and feel free to support Gate with a 🌟 and contributions!
## [2][Mocking Techniques for Go](https://www.reddit.com/r/golang/comments/idip3d/mocking_techniques_for_go/)
- url: https://dmv.myhatchpad.com/insight/mocking-techniques-for-go/
---

## [3][Does profefe (a system for continuous profiling for Go) prefers "push" over "pull"?](https://www.reddit.com/r/golang/comments/idwi4d/does_profefe_a_system_for_continuous_profiling/)
- url: https://vladimir.varank.in/notes/2020/08/does-profefe-prefers-push-over-pull/
---

## [4][What do you use for authentication in golang?](https://www.reddit.com/r/golang/comments/idsvuv/what_do_you_use_for_authentication_in_golang/)
- url: https://www.reddit.com/r/golang/comments/idsvuv/what_do_you_use_for_authentication_in_golang/
---
Hey guys, I'm wondering how do you do authentication in golang?

I found goth ([https://github.com/markbates/goth](https://github.com/markbates/goth)) and authboss ([https://github.com/volatiletech/authboss](https://github.com/volatiletech/authboss)) - both have thousands of stars - which one would you recommend?

I've tried finding tutorials on auth in golang online, but every implementation is very different. I'm looking for something like Passport.js for golang or a more standard and recommended practice.

P/S: I'm aware this has been asked in a similar 5-year-old post, but it is quite old and I think it doesn't have enough discussions going on, so I'd like to bring up the topic again.

[https://www.reddit.com/r/golang/comments/2saim3/a\_good\_auth\_library\_for\_golang/](https://www.reddit.com/r/golang/comments/2saim3/a_good_auth_library_for_golang/)

If you have found a good piece of blog post that you'll refer to everytime you're coding authentication module, please share it with us!

Thanks a lot, gophers!
## [5][Implemented LZW compression as my first non-trivial Go program. Looking for constructive criticism on the code!](https://www.reddit.com/r/golang/comments/ido11x/implemented_lzw_compression_as_my_first/)
- url: https://www.reddit.com/r/golang/comments/ido11x/implemented_lzw_compression_as_my_first/
---
[https://github.com/davidcrosby/lzw-go](https://github.com/davidcrosby/lzw-go)

Just started messing around with Go this week. Seems like a great language with a great community!

P.S. I know it's in the standard library
## [6][Is it possible to run a loop without getting stuck to it?](https://www.reddit.com/r/golang/comments/idv01i/is_it_possible_to_run_a_loop_without_getting/)
- url: https://www.reddit.com/r/golang/comments/idv01i/is_it_possible_to_run_a_loop_without_getting/
---
Guys, all right?

I  am making a countdown timer, I pass a "--start 25m" argument and the  countdown begins. I use simple logic for this, I take the time in  minutes and convert to seconds, I loop and I subtract those seconds that  I got:

    arg := os.Args[1]
    period := convertToSeconds(arg)
    for period != 0 {
    	sendNotification(formatTime(period).String())
    	period--
    	sleep()
    }

This part works perfectly. With the help of dunst, a little linux program, I can display the time continuously on my screen.

The  problem with using the loop is that I get stuck with it until it ends, I  don't want this, I want the execution to be separated and I can execute  other commands: "--stop", "--reset".

I  was a bit naive and separated my logic in two parts, in one file I put  all the logic of the command line arguments and in another file I put  all the logic of the regressive timer.

Can you help me with that? How can I run a countdown through a loop without getting stuck in it?
## [7][🤖Building a Telegram bot with Apache Kafka, Go, and ksqlDB [SLIDES/CODE/RECORDING]](https://www.reddit.com/r/golang/comments/id76l5/building_a_telegram_bot_with_apache_kafka_go_and/)
- url: /r/apachekafka/comments/id7224/building_a_telegram_bot_with_apache_kafka_go_and/
---

## [8][Store bitmap in an array](https://www.reddit.com/r/golang/comments/idln85/store_bitmap_in_an_array/)
- url: https://www.reddit.com/r/golang/comments/idln85/store_bitmap_in_an_array/
---
I'm taking a screenshot of my display using robotgo. I want to save the resulting bitmap to an array for processing later. What data type would I use, if it's actually possible to do this?
## [9][`go build` vs `go mod vendor`](https://www.reddit.com/r/golang/comments/idl01w/go_build_vs_go_mod_vendor/)
- url: https://www.reddit.com/r/golang/comments/idl01w/go_build_vs_go_mod_vendor/
---
I came across an issue today where \`go build\` succeeds and produces an executable on a new go setup (\`rm -rf \~/go\` and \`rm -rf .caches/go-builds\`) but \`go mod vendor\` fails with:  
 "&lt;dep&gt;@latest found but does not contain package &lt;dep&gt;/package-name". Is there a reason that one would succeed in finding and downloading all dependencies but the other would fail?
## [10][Making Asteroids Game with Golang, Lorca/Webview and WASM](https://www.reddit.com/r/golang/comments/idihc8/making_asteroids_game_with_golang_lorcawebview/)
- url: https://maori.geek.nz/making-asteroids-game-with-golang-lorca-webview-and-wasm-9a8bed30cf72
---

