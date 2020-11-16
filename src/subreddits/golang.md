# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][Are there any advance programming tutorials or guides where someone guide you to build say a network protocal from sctach.](https://www.reddit.com/r/golang/comments/jv5sbp/are_there_any_advance_programming_tutorials_or/)
- url: https://www.reddit.com/r/golang/comments/jv5sbp/are_there_any_advance_programming_tutorials_or/
---
Not just networks but any lower level system stuff that is generally abstracted away.
Like a framework or something
## [3][Turns out my mom is pretty talented](https://www.reddit.com/r/golang/comments/jv6e1j/turns_out_my_mom_is_pretty_talented/)
- url: https://i.redd.it/x892hdevolz51.jpg
---

## [4][Conquering the Cache Stampede](https://www.reddit.com/r/golang/comments/jv3ifq/conquering_the_cache_stampede/)
- url: https://medium.com/sidneywijngaarde/conquering-the-cache-stampede-3a4c9feb691c?source=friends_link&amp;sk=ed94d5ac9af040e296e8f667c0d6867f
---

## [5][Casbin： The authority management system with the most complete language types](https://www.reddit.com/r/golang/comments/jv3lqw/casbin_the_authority_management_system_with_the/)
- url: https://www.reddit.com/r/golang/comments/jv3lqw/casbin_the_authority_management_system_with_the/
---
&amp;#x200B;

https://preview.redd.it/i5d019fd8kz51.png?width=323&amp;format=png&amp;auto=webp&amp;s=5796057e3b9b0236f13ceef13854b156ff2879fe

 Casbin is an open source library of access control designed to help complex systems solve access management problems.Casbin adopts the design idea of metamodel, which not only supports classic access control models such as ACL (access control list), RBAC (role-based access control) and ABAC (attribution-based access control), but also supports users to define permissions flexibly according to their own requirements.Casbin has been used open source by Intel, IBM, Tencent Cloud, VMware, RedHat, T-Mobile, etc., and closed source by Cisco, Verizon, etc.See the Casbin home page (https://casbin.org/) for details.

Casbin was originally an open source lightweight unified access control framework built in the Go language.The development, has gradually extended to the Go, Java, Node. Js, Javascript, P你们，ython, PHP, (React). The.net, Delphi, Rust and other languages, in making open source (https://github.com/casbin/casbin), the main project 7000 stars + on the lot.The project now has a stable maintenance team of about 10 people, and is in continuous development.

[Casbin's official website](https://casbin.org/)
## [6][what's the difference of []int() and []int{}](https://www.reddit.com/r/golang/comments/jv4fx4/whats_the_difference_of_int_and_int/)
- url: https://www.reddit.com/r/golang/comments/jv4fx4/whats_the_difference_of_int_and_int/
---

## [7][🚀 A real-time Golang runtime stats visualization profiler](https://www.reddit.com/r/golang/comments/jur6ms/a_realtime_golang_runtime_stats_visualization/)
- url: https://www.reddit.com/r/golang/comments/jur6ms/a_realtime_golang_runtime_stats_visualization/
---
Hey, gophers! I bring you guys a cool project again.

project link: https://github.com/go-echarts/statsview

[**Statsview**](https://github.com/go-echarts/statsview) is a real-time Golang runtime stats visualization profiler. It is built top on another open-source project, [go-echarts](https://github.com/go-echarts/go-echarts), which helps statsview to show its graphs on the browser.

## Installation

```shell
$ go get -u github.com/go-echarts/statsview/...
```

## Usage

Statsview is quite simple to use.

```golang
import (
    "time"

    "github.com/go-echarts/statsview"
)

func main() {
    go func() {
        mgr := statsview.New()

        // Start() runs a HTTP server at `localhost:18066` by default.
        mgr.Start()

        // Stop() will shutdown the http server gracefully
        // mgr.Stop()
    }()

    // busy working....
    time.Sleep(time.Minute)
}

// Visit your browser at http://localhost:18066/statsview/debug
```

## Configuration

Statsview gets a variety of configurations for the users. Everyone could customize their favorite charts style.

```golang
// WithInterval sets the interval(in millisecond) of collecting and pulling metrics
// default -&gt; 1500
WithInterval(interval int)

// WithMaxPoints sets the maximum points of each chart series
// default -&gt; 40
WithMaxPoints(n int)

// WithTemplate sets the rendered template which fetching stats from the server and
// handling the metrics data
WithTemplate(t string)

// WithAddr sets the listen address
// default -&gt; "localhost:18066"
WithAddr(addr string)

// WithTimeFormat sets the time format for the line-chart Y-axis label
// default -&gt; "15:04:05"
WithTimeFormat(s string)

// WithTheme sets the theme of the charts
// default -&gt; Macarons
//
// Optional:
// * ThemeWesteros
// * ThemeWalden
// * ThemeMacarons
WithTheme(theme Theme)
```

#### Set the options

```golang
import (
    "github.com/go-echarts/statsview/viewer"
)

// set configurations before calling the `Start()` method
viewer.SetConfiguration(viewer.WithTheme(viewer.ThemeWalden), view.WithAddr("localhost:8087"))
```

## Viewers

Viewer is the abstraction of a Graph which in charge of collecting metrics from somewhere. Statsview provides some default viewers as below.

* `GCCPUFractionViewer`
* `GCNumViewer`
* `GCSizeViewer`
* `GoroutinesViewer`
* `HeapViewer`
* `StackViewer`

Viewer wraps a go-echarts [Line instance](https://github.com/go-echarts/go-echarts/blob/master/charts/line.go) which means you can use all of the options/features on it. To be honest, I think that is the most charming thing about this project.

## Snapshot

[Macarons](https://user-images.githubusercontent.com/19553554/99192859-45943400-27b0-11eb-8096-8a9e76fba3a1.png)

[Westeros](https://user-images.githubusercontent.com/19553554/99193211-78d7c280-27b2-11eb-96c8-cbcb6792e68a.png)

🐶 PR are always welcome :)
## [8][Golang web service framework with DI, Telemetry and more](https://www.reddit.com/r/golang/comments/juvfaf/golang_web_service_framework_with_di_telemetry/)
- url: https://github.com/go-masonry/mortar
---

## [9][A Distributed Background Task Runner based on RabbitMQ and Redis](https://www.reddit.com/r/golang/comments/jv37d7/a_distributed_background_task_runner_based_on/)
- url: https://www.reddit.com/r/golang/comments/jv37d7/a_distributed_background_task_runner_based_on/
---
I have worked on this amazing job processing service based on [\#rabbitmq](https://twitter.com/hashtag/rabbitmq?src=hashtag_click) and [\#redis](https://twitter.com/hashtag/redis?src=hashtag_click) in the last job which was open source but never promoted. I forked it and made some improvements and now open sourced it. Use it [\#golang](https://twitter.com/hashtag/golang?src=hashtag_click) people. PRs and contributions welcome.  Please create issues if you find any. Also, comment if you need help with the documentation or using the project.

[https://github.com/joker666/cogman](https://github.com/joker666/cogman)

https://preview.redd.it/cqgqc3qe9kz51.png?width=1640&amp;format=png&amp;auto=webp&amp;s=37231c6fe07e88dc15b33d61518a23d812c92e81
## [10][Should You Commit the Vendor Folder in Go? - Qvault](https://www.reddit.com/r/golang/comments/jv6p8e/should_you_commit_the_vendor_folder_in_go_qvault/)
- url: https://qvault.io/2020/11/16/should-you-commit-the-vendor-folder-in-go/
---

## [11][If Go could turn off its GC optionally like Nim/Crystal, what benefits would you expect? Would it be viable like C/Rust performance for systems dev? Would a company like Discord not have swtiched to Rust from Go if it had this? What are your thoughts?](https://www.reddit.com/r/golang/comments/junupo/if_go_could_turn_off_its_gc_optionally_like/)
- url: https://www.reddit.com/r/golang/comments/junupo/if_go_could_turn_off_its_gc_optionally_like/
---
I'm interested in seeing what other's ideas on this are. 
Go certainly does its job and does it well....but Discord's issue with Go was its latency spikes that they were unable to deal with satisfactorily at their scale due to the GC and changed their stack from Elixir+Go to Elixir+Rust in serving 20 million concurrent users. 

Hearing that Nim and Crystal have features where one can turn off the GC and also write unsafe code as well...if optional GC was implemented in Go, do you think this will benefit Go in any way or would it actually go against the philosophy of being simple/productive?

[Link to explanation behind Discord's decision and their attempts](https://blog.discord.com/why-discord-is-switching-from-go-to-rust-a190bbca2b1f)&lt;--(Link leaves Reddit) 

Current verdict:

- One can turn off the Go GC to gain further low latency benefits, linking the github issue in the comment [refer to comment by u/pxrage](https://www.reddit.com/r/golang/comments/junupo/if_go_could_turn_off_its_gc_optionally_like/gcelmcb?utm_medium=android_app&amp;utm_source=share&amp;context=3) and his experience with facing and solving a similar problem to Discord. &lt;--(I recommend following this comment thread, healthy discussion, decent links) 

- We can already make the cache impervious to the GC by C allocation. (link leaves Reddit) --&gt; [Link to its implementation ](https://dgraph.io/blog/post/manual-memory-management-golang-jemalloc/)

- [u/0xnml mentioned that turning off GC is possible with Go including link](https://www.reddit.com/r/golang/comments/junupo/if_go_could_turn_off_its_gc_optionally_like/gcel6ms?utm_medium=android_app&amp;utm_source=share&amp;context=3) 

- So why rewrite the entire Go service in Rust instead of only using Rust where required?..we're not certain. However, [u/slantview mentions that Discord is now moving all new projects to Rust](https://www.reddit.com/r/golang/comments/junupo/if_go_could_turn_off_its_gc_optionally_like/gceqrkk?utm_medium=android_app&amp;utm_source=share&amp;context=3) and could possibly be why they made this decision, who knows?

- [u/painya comment links to Twitch's approach to GC management ](https://www.reddit.com/r/golang/comments/junupo/if_go_could_turn_off_its_gc_optionally_like/gcesw4h?utm_medium=android_app&amp;utm_source=share&amp;context=3)

- [u/doomfrog666 linking and mentioning that the Go team is aware of GC issues and are working on tuning it](https://www.reddit.com/r/golang/comments/junupo/if_go_could_turn_off_its_gc_optionally_like/gcfq4ud?utm_medium=android_app&amp;utm_source=share&amp;context=3)

- Note: Go is awesome, some of us even learned today that we can even turn off the GC to accomplish some things we'd might need to otherwise code in Rust/C and yes, they are better than Go at certain things. This post was to discuss benefits of Go having an optional GC and its application/implementation and not for language wars. The philosophy of Go is its simplicity and productivity (along with using a GC) which might be the reason why this won't necessarily catch on; and by using Go, you might not ever need this option, but it is here and you can use it. The language you choose and whether you decide to keep the complexity up the stack or down the stack is up to you. Languages are tools yes, but developer skills/proficiency/productivity and maintenance of code are very important as well.

- Also: do keep in mind, we are looking at the scale of Discord. They currently use Elixir+Go+Rust (with Rust instead of Go for new projects) to scale 20 million concurrent users. Using Rust, Elixir NIFs and a tuned/turned off Go GC would be overkill for most of our cases. However, if we do reach that scale, we know that we have options to deal with this✌️
