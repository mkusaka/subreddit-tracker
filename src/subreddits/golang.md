# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][Dotfile - Version Control System for Single Files](https://www.reddit.com/r/golang/comments/jt9yp7/dotfile_version_control_system_for_single_files/)
- url: https://github.com/knoebber/dotfile
---

## [3][Abstracting dependencies.. where do you draw the line?](https://www.reddit.com/r/golang/comments/jtevfj/abstracting_dependencies_where_do_you_draw_the/)
- url: https://www.reddit.com/r/golang/comments/jtevfj/abstracting_dependencies_where_do_you_draw_the/
---
Hey,

I'm not a real software engineer. Officially I'm a "senior product manager" of antispam for a freemail service with millions of daily active users (biggest in my country). I have scientific background in physics and during good days spend more time programming than doing manager stuff (delegating stuff and sitting on my ass waiting is not my style) so I dabble in programming a little bit.. (mostly in python + some C++ if need be).

&amp;#x200B;

I and my team learned go and trying to rewrite in it some critical parts of our architecture at work (not for fun, it has real benefits for us).

&amp;#x200B;

Also after a six years of fighthing an uphil battle agains the bad guys while being understafed and not exactly a company priority, I'm getting disenchanted and tired of working for someone else..

&amp;#x200B;

During those years my job title changed a bit, but I'm still doing basically the same thing since I started working there after college.

&amp;#x200B;

So during the little bit of time I have left at home (not easy with one year old son), I'm trying to expand my horizons on side projects and going from backend to full stack (learning JS without the framework mess and using golang for backend).

&amp;#x200B;

I have a project in my head, that I thing is missing on the web and people would apreciate.. But there is an overwhelming amount of work and studing I have to do in order to create it, so I'm kinda procrastinating (hello reddit).

&amp;#x200B;

The stuff I procrastinate at the moment is creating a backend server template (personal "framework"), that I could use in the future to quickly start making servers and not have to bootstap things over and over.

`----------------------------------------------------------------------------------`

Ok, enough of the rant.. **The real question starts here:**

&amp;#x200B;

Abstracting and isolating dependencies is considered good practice (for example [https://research.swtch.com/deps](https://research.swtch.com/deps) by Rob Pike), but I wonder where to draw the line of hiding dependencies with wrapper packages?

&amp;#x200B;

Like having conf package that allows populating arbitrary config structure from specified conf file or env variables - internally I'm using viper and the package exposes only the following so I could reimplement it at any time.

&amp;#x200B;

    func ReadFlagsHelper() (confFile, envPrefix string)
    func LoadYAML(file string, config interface{}) error
    func LoadENV(envPrefix string, config interface{}) error

&amp;#x200B;

Config is easy, how about logger? I'm using zerolog simply embedded inside my own structure so I can call zerolog methods, but also hopefully reimplement them over my Logger struct (in an unlikely scenario in which I would need to replace zerolog by something else, without the need to change logging all over the app).

    type Logger struct {
        zerolog.Logger
        file   *os.File
        mask   string
    }

Database, message broker are another candidates for wrapping.. How about router? Would you wrap something like [https://github.com/go-chi/chi](https://github.com/go-chi/chi) ?

Does anyone in the real world actually bothers with isolating and abstracting third party dependencies? Or do I just waste time on it?  


Thanks for your thoughts.
## [4][Vuejs frontend + golang gcp functions API + firebase db =&gt; viable or overkilled stack?](https://www.reddit.com/r/golang/comments/jtg9xk/vuejs_frontend_golang_gcp_functions_api_firebase/)
- url: https://www.reddit.com/r/golang/comments/jtg9xk/vuejs_frontend_golang_gcp_functions_api_firebase/
---
Hello all!

Mostly for the sake of learning new languages and tools, I want to build a classic blog website. It will handle authenticated users, articles and an access system based on users' roles. 

I would like to build the front end with vuejs (or elm, if I have enough spare time), the back end would be composed of a REST API made in Golang functions (FaaS), stored in GCP. And for the database and the auth, I would use Firebase.

But as I understand things, I could very well use Firebase directly with my frontend and bypass completely a homebrew backend.

Have you been in a similar situation? Would you use golang FaaS for an API? Should I use it only for some "complex" computing and use Firebase API for all the simple stuff?
Again, this is a basic project mostly for learning go, FaaS system and Firebase. But if it's viable, I might use it for a personal project.
## [5][Announcing the Ardan Labs Podcast with Bill Kennedy!! Say hello to our first guest &amp; security nerd, Andy Walker!](https://www.reddit.com/r/golang/comments/jswlix/announcing_the_ardan_labs_podcast_with_bill/)
- url: https://youtu.be/mQidWZbc8JM
---

## [6][teler: Real-time HTTP Intrusion Detection](https://www.reddit.com/r/golang/comments/jszbf3/teler_realtime_http_intrusion_detection/)
- url: https://github.com/kitabisa/teler
---

## [7][Building a WebRTC video and audio Broadcaster in Golang using ion-sfu and media devices](https://www.reddit.com/r/golang/comments/jszgcn/building_a_webrtc_video_and_audio_broadcaster_in/)
- url: https://gabrieltanner.org/blog/broadcasting-ion-sfu
---

## [8][A fast thread-safe in-memory cache server that supports a big number of entries in Go](https://www.reddit.com/r/golang/comments/jt13nd/a_fast_threadsafe_inmemory_cache_server_that/)
- url: https://github.com/ziyasal/distrox
---

## [9][Efficient struct packing guided pass for Go](https://www.reddit.com/r/golang/comments/jtgehp/efficient_struct_packing_guided_pass_for_go/)
- url: https://medium.com/orijtech-developers/efficient-struct-packing-guided-pass-for-go-92255872ec72
---

## [10][syscall package deprecated - question](https://www.reddit.com/r/golang/comments/jt1qfm/syscall_package_deprecated_question/)
- url: https://www.reddit.com/r/golang/comments/jt1qfm/syscall_package_deprecated_question/
---
While working on a lib using syscall package, I've stumbled upon a [deprecation notice](https://golang.org/pkg/syscall/).  It seems syscall is frozen [since go 1.4 release](https://golang.org/doc/go1.4#major_library_changes). The authors recommend using `golang.org/x/sys` instead.

So why `x/sys` is not part of the standard library yet? Are there plans to include it there?

(technically, /x/ packages *are* part of standard library, however I have to `go get` them instead of them being readily available in the distribution).
## [11][k6 v0.29.0 is out! 🎊🎉🥳](https://www.reddit.com/r/golang/comments/jssvos/k6_v0290_is_out/)
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
