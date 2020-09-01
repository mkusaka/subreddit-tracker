# golang
## [1][I wrote a music video in Go](https://www.reddit.com/r/golang/comments/ikibl5/i_wrote_a_music_video_in_go/)
- url: https://www.youtube.com/watch?v=-_-2EpUqb9g
---

## [2][Go playground provides options like Multiple files and writing test cases 💡️](https://www.reddit.com/r/golang/comments/ik0m92/go_playground_provides_options_like_multiple/)
- url: https://i.redd.it/as61kin11dk51.png
---

## [3][Println debugging and tracing tool capable of visualising function invocation paths](https://www.reddit.com/r/golang/comments/ikiy1w/println_debugging_and_tracing_tool_capable_of/)
- url: https://www.reddit.com/r/golang/comments/ikiy1w/println_debugging_and_tracing_tool_capable_of/
---
Hello Gophers,

I am happy to share with you an idea i had for some time now and i finally managed to find time to make happen!

In short like many of us i have found myself struggling when onboarding on a new projects which has already a decent codebase. 

I realised that i am constantly adding some prints to see what’s going on in the flow and then the idea was born.

What if i can add all the prints i need in a matter of a second and when i am comfortable with the flow that i am interested in i can revert all of the prints in an automated fashion.

Even more what if i can see it in a visual representation from the function i am interested in down the rabbit hole. It would be great. And here it is  [prinTracer](https://github.com/DimitarPetrov/printracer).

&amp;#x200B;

[prinTracer](https://github.com/DimitarPetrov/printracer) is a go tool that instruments your codebase with fmt.Println(…) statements for every function execution and then visualize all function invocations in a sequence diagram. 

You can check the README.md for better visual explanation.

&amp;#x200B;

It will be great if it helps you and saves you some hard times!

Please don’t hesitate to give some feedback it is always appreciated.
## [4][Add drop shadow effect to images](https://www.reddit.com/r/golang/comments/ikif80/add_drop_shadow_effect_to_images/)
- url: https://www.reddit.com/r/golang/comments/ikif80/add_drop_shadow_effect_to_images/
---
A utility to add "drop-shadow" effect to images using Go.

I created this one for the documentation team for the images that they add into the documentation. And to avoid dependency on Photoshop or similar softwares.

[https://github.com/jerrymannel/imageDropShadow](https://github.com/jerrymannel/imageDropShadow)

[Original Image](https://preview.redd.it/h0z53wt6sik51.png?width=948&amp;format=png&amp;auto=webp&amp;s=cfbc93a181090a4c5cec23a74dfec322b3579e15)

[With a stroke and shadow](https://preview.redd.it/vnr4vwt6sik51.png?width=1023&amp;format=png&amp;auto=webp&amp;s=b5f41d7d08a35bb91be484dc18195f7b63980d9d)

Image courtesy of [u/Brandinious](https://www.reddit.com/user/Brandinious/) \- [Calvin and Hobbes \[1920x1200\]](https://www.reddit.com/r/wallpapers/comments/ie81pv/calvin_and_hobbes_1920x1200/)
## [5][Tuning GO Runtime for Large Heaps](https://www.reddit.com/r/golang/comments/ik7wjp/tuning_go_runtime_for_large_heaps/)
- url: https://www.reddit.com/r/golang/comments/ik7wjp/tuning_go_runtime_for_large_heaps/
---
I am currently working on a service that calls a number of other services, processes and aggregates the data into something meaningful. Because of the sheer amount of data and processing, by nature, it tends to use quite a bit of heap under load. The payloads from the upstream systems can be quite large. The issue I am currently facing is the CPU is running very high on this service, and by profiling with PPROF it looks like 40% of the CPU time is being spent doing garbage collection activities, although the heap usage is not surpassing 1 GB, and the container max memory is 4GB. The GO runtime appears to be aggressively trying to keep the memory consumption down at the cost of CPU. I'm curious if there is a guide for tuning the GO runtime for higher memory workloads. I've seen in the documentation there is an environment variable GOGC which can tune how aggressively the garbage collector runs. Is this the recommended route for tunning the runtime? I'm looking for sacrifice higher heap usage for the sake of lower CPU consumption by the GC.
## [6][Dedicated golang plugin for Geany IDE?](https://www.reddit.com/r/golang/comments/ikfp4b/dedicated_golang_plugin_for_geany_ide/)
- url: https://www.reddit.com/r/golang/comments/ikfp4b/dedicated_golang_plugin_for_geany_ide/
---
I am a user of Geany IDE and I personally think its the most under rated IDE despite it being open source, lite and fast. I mostly use it for python  programming.

I recently started to program in go and use geany editor. Its still fun to code with geany but I really  wish it had a dedicated golang plugin.

Features I wish,

Autocompletion: geany has auto completion but its not devoted for golang

Code completion: Able to fill multiplie lines of repetitive code, for example,  err codes.

Jumping to source code: when checking  a method name it should auto jump to the line of source file. If I check for fmt.Print from my current file then the file implementating fmt.Print should open in another tab of geany.

Other features native to golang.
## [7][obscure-go: Implement In-memory data type security for your projects (Anti-CheatEngine)](https://www.reddit.com/r/golang/comments/ik2uj9/obscurego_implement_inmemory_data_type_security/)
- url: https://github.com/Dentrax/obscure-go
---

## [8][A reflection based cli toolkit](https://www.reddit.com/r/golang/comments/ikbnbb/a_reflection_based_cli_toolkit/)
- url: https://www.reddit.com/r/golang/comments/ikbnbb/a_reflection_based_cli_toolkit/
---
I've been building a reflection based cli framework called [quack](https://github.com/eliothedeman/quack) (think rust's structopt + python's click).  
I haven't found anything like this for go before, so I'm interested to see if anyone knows of any prior art. Also interested to hear people's thoughts on the API
## [9][Ultimate Go video course for FREE through your public library](https://www.reddit.com/r/golang/comments/ik63w6/ultimate_go_video_course_for_free_through_your/)
- url: https://www.reddit.com/r/golang/comments/ik63w6/ultimate_go_video_course_for_free_through_your/
---
I just found out about this course a couple days ago due the post by the guy who repurposed his notes into a book. (don't necessarily agree w/ how he went about it but, it did lead me to the course).

Anyways,  I started googling the course to find out more and saw it was fairly expensive.  Then, I remember my library offers free access to Lynda and Udemy.  So, I thought, maybe they do the same for Oreilly materials.  Sure enough, answer is yup.

So if you don't have the $280 for the course or don't want to pay the monthly subscription for O'reilly Learning, check out your public library's page and see if they offer O'reilly material for free.
## [10][GORM V2 officially released!](https://www.reddit.com/r/golang/comments/ijlhiw/gorm_v2_officially_released/)
- url: https://gorm.io/docs/v2_release_note.html
---

