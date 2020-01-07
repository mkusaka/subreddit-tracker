# golang
## [1][101+ Coding Interview Problems with Detailed Solutions, Test Cases, and Program Analysis](https://www.reddit.com/r/golang/comments/el4vut/101_coding_interview_problems_with_detailed/)
- url: https://www.reddit.com/r/golang/comments/el4vut/101_coding_interview_problems_with_detailed/
---
Hi friends,

4 months ago, I made a [reddit post](https://www.reddit.com/r/learnprogramming/comments/ctyvbc/is_anyone_interested_in_weekly_coding_interview/) on creating a newsletter that sends out 3-6 coding interview problems with details solutions every week. To my surprise, the post received a lot of attention and positive feedback from the community.

**Since the total number of problems and solutions has just passed 101, I would like to say another big thank you all again for all the support. It really helps me keep going.**

Moving on, I am hoping to add more questions and then finally condense them into a list of most frequently appeared coding interview problems that I think are the most valuable and productive to spend time on. Again, my goal has always been to help you get good at algorithms and data structures so that you can prepare better for your next coding interviews.

For those who don’t know, [here is the link to all 101+ questions and solutions, sorted by resources →](https://github.com/hoanhan101/algo)

If you’re interested in getting updates for this, feel free to check out [my blog](https://hoanhan101.github.io/) and join [my mail list here](https://tinyletter.com/hoanhan).

Best,

Hoanh
## [2][Continuous profiling in Go with Profefe](https://www.reddit.com/r/golang/comments/el9alc/continuous_profiling_in_go_with_profefe/)
- url: https://gianarb.it/blog/go-continuous-profiling-profefe
---

## [3][Make resilient Go net/http servers using timeouts, deadlines and context cancellation](https://www.reddit.com/r/golang/comments/el0cxx/make_resilient_go_nethttp_servers_using_timeouts/)
- url: https://ieftimov.com/post/make-resilient-golang-net-http-servers-using-timeouts-deadlines-context-cancellation
---

## [4][Anyone read o'reilly "Head First Go" 2019? It's telling me to drink water, no seriously...](https://www.reddit.com/r/golang/comments/ekwqxc/anyone_read_oreilly_head_first_go_2019_its/)
- url: https://www.reddit.com/r/golang/comments/ekwqxc/anyone_read_oreilly_head_first_go_2019_its/
---
I just started reading it and I get how they're trying not to be "noring" by adding all these creative visuals &amp; comedy, but I don't know if this is the right approach to learning something especially since they needed an intro to convince me that this approach was better.

Also tip 6:

&gt; Drink water. Lots of it.

So wondering if this book actually helped anyone learn Go or even improved their Go skills.  Thanks.
## [5][[Hiring] Golang engineers in SoCal](https://www.reddit.com/r/golang/comments/ekzodh/hiring_golang_engineers_in_socal/)
- url: https://www.reddit.com/r/golang/comments/ekzodh/hiring_golang_engineers_in_socal/
---
Hey everyone,

I'm working with 3 product companies in OC hiring for Golang. 

The first company is just a short term contract that's 100% remote, this company is actually LA based but one of the biggest names in media, period. The second company is an IOT software product, they're in Irvine. The third opportunity is also in the media space, however with a software product company that builds the backend system for broadcasting news. Their office is in Newport.

I'd like to chat with anyone who wants to learn more about any of these three companies! Pass along to your friends 

[stacey.papernaya@jobspringpartners.com](mailto:stacey.papernaya@jobspringpartners.com) \&gt; my DMs.
## [6][tdewolff/canvas: Support for HTML Canvas through WASM](https://www.reddit.com/r/golang/comments/ekykv9/tdewolffcanvas_support_for_html_canvas_through/)
- url: https://tdewolff.github.io/canvas/examples/html-canvas/
---

## [7][Global variable vs. function for constant structs?](https://www.reddit.com/r/golang/comments/el49xz/global_variable_vs_function_for_constant_structs/)
- url: https://www.reddit.com/r/golang/comments/el49xz/global_variable_vs_function_for_constant_structs/
---
Hi. In my program I have a struct that represents the program's configuration, `Config`. I need a way of accessing the default configuration in my code somehow. I know that Go does not allow you to define constant structs, therefore, I am curious which of the following is considered better practice.


**Global Variable**

```
var DefaultConfig Config = Config{
	...
}
```

**Function**

```
func DefaultConfig() Config {
	return Config{
		...
	}
}
```

Thoughts?
## [8][Go Things I Love: Channels and Goroutines by Justin Fuller](https://www.reddit.com/r/golang/comments/el996g/go_things_i_love_channels_and_goroutines_by/)
- url: https://www.justindfuller.com/2020/01/go-things-i-love-channels-and-goroutines/
---

## [9][A Recap of Golang Taipei Meetup in December](https://www.reddit.com/r/golang/comments/el4ko3/a_recap_of_golang_taipei_meetup_in_december/)
- url: https://diode.io/diode/Diode-Presented-at-the-Golang-Taipei-Meetup-20002/
---

## [10][[Help] Could someone explain bufio.Scanner?](https://www.reddit.com/r/golang/comments/el2iks/help_could_someone_explain_bufioscanner/)
- url: https://www.reddit.com/r/golang/comments/el2iks/help_could_someone_explain_bufioscanner/
---
Hi all, I’m new to Go and I’m struggling with assigning input from the user which contains spaces (i.e. ‘Marble Arch Station’ from the console).

fmt.Scan isn’t correct - but I read online that bufio.Scanner is - but I’m struggling to implement it.

Could anyone explain to me how to use bufio.Scanner in this case?

Thanks!
