# golang
## [1][Learnt Go, teenager in dire need looking for advice.](https://www.reddit.com/r/golang/comments/gxltew/learnt_go_teenager_in_dire_need_looking_for_advice/)
- url: https://www.reddit.com/r/golang/comments/gxltew/learnt_go_teenager_in_dire_need_looking_for_advice/
---
Hi Everyone,

I'm 19 years old, decided to learn programming on my own, and bought two Go books and gave it a go, pun intended :) Finished them both, last topic I learnt was the basics of Concurrent Programming.

I've decided that this is something I could see myself doing for a living, but I come from a very poor family and due to this crazy pandemic, both of my parents have lost their jobs, we living off unemployment benefits and we won't last too long like this so I am doing everything I can to be able to help out as soon as humanly possible.

I've decided I wont be going to college because I dont have 4 years, not right now, I need to produce money like yesterday.  Long story short I am teaching myself programming using Go and now Ive been reading classic basic Data Structure and Algorithms book which btw, there are none available in Go, so been
reading a Data Structures book in Java and sort of converting them to Go as added exercise, algorithms too.  I took very basic Java in high school, so it's enough to understand some of it, the rest I sort of figure out as I progress.

Ok, is there a question here?

There sure is!

I want to become a paid developer, after I finish converting this Data Structures and Algorithms book to Go and reading and learning about these concepts, what's my next step? I dig Go, I like it, so anything that allows me to use it would be great.  I've read the web development market is flooding with juniors, so I dont even want to go there, what do you guys suggest? Where I could move to next?

Read I should learn Git and get it going, and well, that's about all I know as far as my next step.

Oh before I forget, where did I get the idea to learn Go and DS and Algorithms? I emailed my high school Java teacher and he said it's a great language, he was right, absolutely love it.

So there you have it,  your help is VERY VERY MUCH APPRECIATED.

BTW my computer kind of sucks but been using Visual Studio Code and no issues compiling or reading PDFs!  

Good thing Im not a gamer, I hear they use a lot of resources, but, no time for gaming, need to become a coder!

Sorry if I made typos wrote this on a very old iPhone.

Cheers.
## [2][My first app in Go](https://www.reddit.com/r/golang/comments/gxoe5k/my_first_app_in_go/)
- url: https://github.com/0xfederama/water-reminder
---

## [3][srv: minimalist alternative to python -m http.server](https://www.reddit.com/r/golang/comments/gxjkjy/srv_minimalist_alternative_to_python_m_httpserver/)
- url: https://github.com/joshuarli/srv/releases/tag/0.2
---

## [4][Webgo v4.0.1](https://www.reddit.com/r/golang/comments/gxn297/webgo_v401/)
- url: https://www.reddit.com/r/golang/comments/gxn297/webgo_v401/
---
After 2yrs of posting here about [webgo v2](https://www.reddit.com/r/golang/comments/898zt8/webgo_v2_a_micro_web_framework/), this is an update. I just released [v4 of webgo](https://github.com/bnkamalesh/webgo). I know a lot of people consider "web frameworks, router etc." a joke/amateur/reinventing the wheel. But I've been able to learn a lot from it and am happy maintaining this package. Hope you like it!

Webgo maintains its promise of standard library compliance, and getting out of the execution path asap. And this version comes with quite a lot of performance improvement in terms of latency &amp; memory usage.

[https://github.com/bnkamalesh/webgo](https://github.com/bnkamalesh/webgo)
## [5][Generate High Level Microservice Architecture diagrams for GraphViz using simple YAML syntax.](https://www.reddit.com/r/golang/comments/gx7syh/generate_high_level_microservice_architecture/)
- url: https://github.com/lucasepe/draft
---

## [6][Don't defer Close() on writable files](https://www.reddit.com/r/golang/comments/gxpf4l/dont_defer_close_on_writable_files/)
- url: https://www.joeshaw.org/dont-defer-close-on-writable-files/
---

## [7][[Update] TermBackTime (v1.0.0-beta.2) - Live terminal sharing via WebRTC.](https://www.reddit.com/r/golang/comments/gxlck5/update_termbacktime_v100beta2_live_terminal/)
- url: https://www.reddit.com/r/golang/comments/gxlck5/update_termbacktime_v100beta2_live_terminal/
---
I've recently pushed v1.0.0-beta.2 of my project [TermBackTime](https://termbackti.me/), adding support for live terminal sharing via WebRTC. As well, there are now ARMv6 and ARMv7 builds for Raspberry Pi support. Because this uses WebRTC it also allows you to use an official TURN server provided by this project, your own server, or you can attempt to use none at all.

When using the a provided TURN server, a request is made to the TermBackTime API for TURN authorization credentials. This API is running via Cloudflare Workers. I've been debating a change of domain for this project's live sharing feature, something more like "https://&lt;token&gt;.&lt;tbt&gt;.live" if anyone has any suggestions on this for perhaps a shorter URL.

A future goal is to share between two TermBackTime CLI clients still using WebRTC.

&amp;#x200B;

* Use an official TURN server provided by TermBackTime:
   * `termbacktime live`
* Provide your own TURN server credentials:
   * `termbacktime live --turn &lt;username&gt;:&lt;password&gt;@&lt;server&gt;[:&lt;port&gt;]`
   * `termbacktime live --user &lt;username&gt; --password &lt;password&gt; --addr &lt;server&gt;[:&lt;port&gt;]`
* Attempt to share without any TURN server:
   * `termbacktime live --no-turn`

&amp;#x200B;

[TermBackTime - Live terminal sharing](https://reddit.com/link/gxlck5/video/eaqr0lmi78351/player)

Prebuilt binaries are available on the [releases](https://github.com/termbacktime/termbacktime/releases) page. As always, any feedback would be greatly appreciated
## [8][Allez Go !! Premier Pas](https://www.reddit.com/r/golang/comments/gxqepb/allez_go_premier_pas/)
- url: https://fredericschmidt.fr/2020/06/06/allez-go-premier-pas/
---

## [9][goscreencasts.io | Go source file anatomy](https://www.reddit.com/r/golang/comments/gxn0do/goscreencastsio_go_source_file_anatomy/)
- url: https://www.goscreencasts.io/videos/03-anatomie-d-un-fichier-go
---

## [10][Trouble with concurrency](https://www.reddit.com/r/golang/comments/gxmk4l/trouble_with_concurrency/)
- url: https://www.reddit.com/r/golang/comments/gxmk4l/trouble_with_concurrency/
---
Hi everyone,  
I'm writing an API in Go which makes use of the Spotify API. I was able to make most of the things work fine except for when I tried adding some tracks to a playlist.  
I had in my hand an array of track IDs and had to add them to a playlist one by one. Instead of adding them one by one I thought this would be the right time to make use of concurrency.  


So, I wrote up a goroutine as follows:  

```
        type result struct {
		err   error
		resId string
	}
	resultCh := make(chan result)
	for _, j := range allTracks {
		go func(id string) {
			snap, err := client.AddTracksToPlaylist(newPlaylist.ID, id)
			if err != nil {
				resultCh &lt;- result{err: err}
			}
			resultCh &lt;- result{resId: snap}
		}(j)
	}
```

This seemed to do the job for me, except when I started receiving 429 - Too Many Requests from the Spotify API. It made sense, I could well be shooting too many requests in a second because of the goroutines. So I thought instead of having an unknown number of goroutines getting created I could batch my requests into groups of 2, 5, or 10 and then add a ticker between these batches so I don't hit that error again. But I still ended up getting Too Many Requests again.
Can someone help me with this and also if this was indeed the right place to make use of concurrency?
Thanks :)
