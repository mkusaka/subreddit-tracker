# golang
## [1][Free Go programming ebook from Digital Ocean](https://www.reddit.com/r/golang/comments/hft9zb/free_go_programming_ebook_from_digital_ocean/)
- url: https://www.digitalocean.com/community/books/how-to-code-in-go-ebook
---

## [2][New tool written in Golang for dynamic configuration and management of Wireguard network overlays](https://www.reddit.com/r/golang/comments/hfz83c/new_tool_written_in_golang_for_dynamic/)
- url: https://www.reddit.com/r/golang/comments/hfz83c/new_tool_written_in_golang_for_dynamic/
---
Hi there!

I’m working on a new open source tool for dynamic configuration of [Wireguard](https://www.wireguard.com/) network overlays. It has many different applications, but is originally meant for establishing hybrid virtual clouds spanning different providers and edge devices.

Since it's written in Go, the project might be of interest to some people in the community. We're still in a very early-stage, and any feedback is much appreciated :-)

Check it out here: [https://github.com/seashell/drago](https://github.com/seashell/drago)

Cheers!
## [3][memplot: generate .PNG memory usage plots of any process from a single binary](https://www.reddit.com/r/golang/comments/hg7gb4/memplot_generate_png_memory_usage_plots_of_any/)
- url: https://github.com/0x0f0f0f/memplot
---

## [4][Simple tool to make downloaded movie file, Plex friendly.](https://www.reddit.com/r/golang/comments/hg7dsa/simple_tool_to_make_downloaded_movie_file_plex/)
- url: https://github.com/m4ns0ur/plexize
---

## [5][Some questions about Golang...](https://www.reddit.com/r/golang/comments/hg720b/some_questions_about_golang/)
- url: https://www.reddit.com/r/golang/comments/hg720b/some_questions_about_golang/
---
Hi there, python developer here. Started learning Go yesterday and am really impressed with it. I generally learn languages by completing projects with them. Anyone have any ideas for beginner friendly golang projects? 

Thanks!
## [6][How to do "frame" based profiling in GO?](https://www.reddit.com/r/golang/comments/hg71ji/how_to_do_frame_based_profiling_in_go/)
- url: https://www.reddit.com/r/golang/comments/hg71ji/how_to_do_frame_based_profiling_in_go/
---
I have a game server in Go and need to profile CPU costs of individual frames but sadly can't find a way to do that properly. Hopefully, someone here can help.

I thought that the \`pprof\` CPU profile should be the right tool. However, at least the profile visualization tools always aggregate the methods execution costs between \`StartCPUProfile\` and \`StopCPUProfile\`. While in some scenarios this is great, we have currently have big spikes in for frame duration and would love to see the difference in times per frame. Ideally in a flame-graph like visualization but grouped by frames. So we'd somehow need to mark the frame start/end and tell the visualization to use that.

Ideally, like the profiler in the Unity3D Editor does:

[Unity3D Profiling - Per frame method execution times](https://preview.redd.it/jck3mr3hz8751.png?width=800&amp;format=png&amp;auto=webp&amp;s=1e7092f877c76b8416a9c35eb0dd0bcadadc5258)

I tried use \`Start/StopCPUProfile\` at the  beginning and end of a frame to create an individual capture per frame. That's not ideal for comparison but should be good enough and simple in concept. Sadly the cost of \`StartCPUProfile\` is so high, we can't call that per frame.

Looking at the code, the the 100 millisecond sleep in \`profileWriter\` explains that it. So what we'd need is a way to mark the frame start/end somehow for visual grouping. Is that possible with \`pprof\`?

[pprof.go](https://preview.redd.it/pahav4l229751.png?width=1203&amp;format=png&amp;auto=webp&amp;s=41671383f2aa5fcd8bb39204701a9f6a5bb0797a)

Further, I wonder if it's even possible to profile a game server with a simulation framerate of 120Hz. Since the 100Hz CPU profile rate is hard coded in \`StartCPUProfile\` I would guess we even per frame aggregation it make not work of a frame duration is usually below 1s / 120, right?
## [7][Go on Windows Problem](https://www.reddit.com/r/golang/comments/hfmy0d/go_on_windows_problem/)
- url: https://www.reddit.com/r/golang/comments/hfmy0d/go_on_windows_problem/
---
I was wondering if anyone using Windows has successfully gotten Go to fully work without having to use WSL. By that I mean in PowerShell I have issues with modules that require CGo (uses gcc) not wanting to install such as the go-sqlite3. I have gotten it so that PowerShell recognizes the gcc command, but it still won't install the module when I use "go get," it exits with a status of 2. The main reason why I am asking is because I have other development environments that don't work in WSL and I would prefer to use one command prompt rather than have to switch terminals depending on which language I am using.
## [8][Kertish-DFS - Highly scalable &amp; available distributed file system](https://www.reddit.com/r/golang/comments/hfo3f1/kertishdfs_highly_scalable_available_distributed/)
- url: https://github.com/freakmaxi/kertish-dfs
---

## [9][Error Handling Question](https://www.reddit.com/r/golang/comments/hg1odn/error_handling_question/)
- url: https://www.reddit.com/r/golang/comments/hg1odn/error_handling_question/
---
In my `main.go`:

How should I handle a database connection or `http.ListenAndServe` error? Should I log and panic? Or should I handle it via a defer recover? I am running a web application and right now, I just log the error, then trigger a panic, not sure if this is the best way to go though.
## [10][How best to work with large exec outputs?](https://www.reddit.com/r/golang/comments/hg1a77/how_best_to_work_with_large_exec_outputs/)
- url: https://www.reddit.com/r/golang/comments/hg1a77/how_best_to_work_with_large_exec_outputs/
---
I have a binary which outputs log lines that I want to do some processing on. There might be a large amount of logs, so it reasons to me that for a well behaved program I don't want to store it all in a buffer all at once.

It would be awesome to have the program write to a buffer and block until I could read the lines from it, but I'm not sure that is possible. Is there a way to call exec.Command and have it output to a buffer with a fixed size? That way I could at least drop the rest of the log lines instead of a possible OOM?
