# cpp
## [1][C++ Jobs - Q1 2020](https://www.reddit.com/r/cpp/comments/eiila4/c_jobs_q1_2020/)
- url: https://www.reddit.com/r/cpp/comments/eiila4/c_jobs_q1_2020/
---
Rules For Individuals
---------------------

* **Don't** create top-level comments - those are for employers.
* Feel free to reply to top-level comments with **on-topic** questions.
* I will create one top-level comment for **meta** discussion.
* I will create another top-level comment for **individuals looking for work** and **community groups looking for sponsors**.

Rules For Employers
---------------------

* You must be hiring **directly**. No third-party recruiters.
* **One** top-level comment per employer. If you have multiple job openings, that's great, but please consolidate their descriptions or mention them in replies to your own top-level comment.
* **Don't** use URL shorteners. [reddiquette](https://www.reddithelp.com/en/categories/reddit-101/reddit-basics/reddiquette) forbids them because they're opaque to the spam filter.
* Templates are awesome. Please **use** the following template. As the "formatting help" says, use \*\*two stars\*\* to **bold text**. Use empty lines to separate sections.
* **Proofread** your comment after posting it, and edit any formatting mistakes.

---

\*\*Company:\*\* [Company name; also, use the "formatting help" to make it a link to your company's website, or a specific careers page if you have one.]

&amp;nbsp;

\*\*Type:\*\* [Full time, part time, internship, contract, etc.]

&amp;nbsp;

\*\*Description:\*\* [What does your company do, and what are you hiring C++ devs for? How much experience are you looking for, and what seniority levels are you hiring for? The more details you provide, the better.]

&amp;nbsp;

\*\*Location:\*\* [Where's your office - or if you're hiring at multiple offices, list them. If your workplace language isn't English, please specify it.]

&amp;nbsp;

\*\*Remote:\*\* [Do you offer the option of working remotely? If so, do you require employees to live in certain areas or time zones?]

&amp;nbsp;

\*\*Visa Sponsorship:\*\* [Does your company sponsor visas?]

&amp;nbsp;

\*\*Technologies:\*\* [Required: do you mainly use C++98/03, C++11, C++14, C++17, or the C++20 working draft? Optional: do you use Linux/Mac/Windows, are there languages you use in addition to C++, are there technologies like OpenGL or libraries like Boost that you need/want/like experience with, etc.]

&amp;nbsp;

\*\*Contact:\*\* [How do you want to be contacted? Email, reddit PM, telepathy, gravitational waves?]

---

Previous Post
--------------

* [C++ Jobs - Q4 2019](https://www.reddit.com/r/cpp/comments/dbqgbw/c_jobs_q4_2019/)
## [2][STL header token parsing benchmarks for VS2017 and VS2019](https://www.reddit.com/r/cpp/comments/eumou7/stl_header_token_parsing_benchmarks_for_vs2017/)
- url: https://www.reddit.com/r/cpp/comments/eumou7/stl_header_token_parsing_benchmarks_for_vs2017/
---
One of the MSVC devs poked me to update the results from [https://www.reddit.com/r/cpp/comments/860hya/stl\_header\_token\_parsing\_benchmarks\_for\_vs2008/](https://www.reddit.com/r/cpp/comments/860hya/stl_header_token_parsing_benchmarks_for_vs2008/) for VS2019:

Graph of VS2019: [https://raw.githubusercontent.com/ned14/stl-header-heft/master/graphs/msvs-2019.png](https://raw.githubusercontent.com/ned14/stl-header-heft/master/graphs/msvs-2019.png)

Comparative graph: [https://raw.githubusercontent.com/ned14/stl-header-heft/master/graphs/msvs-history.png](https://raw.githubusercontent.com/ned14/stl-header-heft/master/graphs/msvs-history.png)

Detailed notes: [https://github.com/ned14/stl-header-heft/blob/master/Readme.msvs.md](https://github.com/ned14/stl-header-heft/blob/master/Readme.msvs.md)

Project github: [https://github.com/ned14/stl-header-heft](https://github.com/ned14/stl-header-heft)

There is a lot of good news in this benchmark: overall VS2019 has 12% lower token processing times than VS2017 did. **That makes VS2019 as quick as VS2008 used to be!**

The single biggest surprise is surely now `&lt;array&gt;`, which is by far the biggest impact STL container now that Microsoft have greatly improved `&lt;string&gt;` and especially `&lt;vector&gt;` and `&lt;forward_list&gt;`.

I'll try to produce updated benchmarks for libstdc++ during this coming week.

I'd like to take this opportunity to thank the Visual C++ team for such outstanding work on slimming down their STL implementation, yet simultaneously implementing a large proportion of C++ 20! Great work guys, wish more STLs were like yours!
## [3][Hacking on Clang is surprisingly easy](https://www.reddit.com/r/cpp/comments/eumv9p/hacking_on_clang_is_surprisingly_easy/)
- url: https://mort.coffee/home/clang-compiler-hacking/
---

## [4][Jonas Minnberg: Tricks, Tips, and a couple of war stories](https://www.reddit.com/r/cpp/comments/eummvc/jonas_minnberg_tricks_tips_and_a_couple_of_war/)
- url: https://youtu.be/CJ5_a4JjbTw
---

## [5][neoGFX -- a new C++ game engine and app creation framework -- coming soon. A Qt killer?](https://www.reddit.com/r/cpp/comments/eunk89/neogfx_a_new_c_game_engine_and_app_creation/)
- url: https://neogfx.org
---

## [6][How do you explain NOT to write high performance code?](https://www.reddit.com/r/cpp/comments/eunny1/how_do_you_explain_not_to_write_high_performance/)
- url: https://www.reddit.com/r/cpp/comments/eunny1/how_do_you_explain_not_to_write_high_performance/
---
Apologize for using a throwaway account, but I don't want my current colleagues to be aware of this post because they are a good bunch of people and I like being here, but technical issues are hard to ignore.


I'm sort of a performance junkie myself, though I'm not on the level of expertise as people working in the HPC industry, but I feel confident enough to know when and how to optimize, but also when not to optimize. When I'm not sure, I do a minimal research to figure out what's the right path to go.


I have a problem at my current employer that I'm sure some others must have experienced as well. It's also visible in some public discussions as well. Performance is good and writing high performance code is good. But it seems like there's a bandwagon of i-must-always-write-most-performant-code people who don't understand most of the low level details to accomplish this. My colleagues are like this.


The codebase is a mess. While it's partially not their fault due to deadlines in the past, it also partially is due to their decisions. Something's slow? Oh, that's simple to fix, let's split that into multiple threads. We don't even know why it's slow, but if we parallelize it, we've optimized it. End result? A billion mutexes on a billion internal variables and noone understand logical relationship between them. Now there's some issue, let's run thread sanitizer. Nice, we found the problem! Let's fix it: randomly add shared and unique locks. If thread sanitizer does not complain with shared lock, shared lock it is, even though we don't really understand anymore what should be exclusive access or if shared lock introduces a sort of logic-race-condition that thread sanitizer will obviously not pick up.


Cache friendly optimizations. During the interview, I felt like everyone knows what it is about simply based on the fact that it was mentioned a couple of times. Well, the truth is they don't really understand it, just heard somewhere that the code needs to be cache friendly and are implementing stuff that makes literally no sense. Truth: most of the code does not need to care about such low level details.


As mentioned in the last paragraph, the most important thing is that most of the code really does not need to be high performance code. That does not stop people from doing micro optimizations that complicate the code even further just so they avoid a vector from being returned (the optimization often ends up complicating the code significantly). Then a few lines below you see a hundred string concatenations which are ugly, but are not an issue performance wise because it does not matter if this method takes 100us or 1ms to execute. Needless to say, they don't even know that it's "slow".


I could go on and on, but I think you get the point. Have you ever found yourself in a position where you need to explain to people in a _KIND_ way that they should stop attempting to do micro optimizations because they don't have basic understanding of most fundamental things required to perform them? If yes, how did you do it? How did you stop your co-workers from performing micro optimizations everywhere? Especially if they don't understand them.
## [7][C++ DataFrame (like Pandas, R) new release](https://www.reddit.com/r/cpp/comments/eunk7f/c_dataframe_like_pandas_r_new_release/)
- url: https://www.reddit.com/r/cpp/comments/eunk7f/c_dataframe_like_pandas_r_new_release/
---
[https://github.com/hosseinmoein/DataFrame](https://github.com/hosseinmoein/DataFrame)

Please see release notes

Your feedback is greatly appreciated
## [8][Performance of Handling Asynchronous Requests using Futures](https://www.reddit.com/r/cpp/comments/eu5jsi/performance_of_handling_asynchronous_requests/)
- url: http://www.mycpu.org/c++-threads-async-deferred/
---

## [9][The Gallows an open-source educational TUI video game in Monochrome Green.](https://www.reddit.com/r/cpp/comments/eu8m6i/the_gallows_an_opensource_educational_tui_video/)
- url: https://github.com/Notfromthisworld/TheGallows
---

## [10][Tell me about your worst cpp interview](https://www.reddit.com/r/cpp/comments/ette1a/tell_me_about_your_worst_cpp_interview/)
- url: https://www.reddit.com/r/cpp/comments/ette1a/tell_me_about_your_worst_cpp_interview/
---
Can anyone describe their worst interview experience as a cpp developer? When I say worst interview, I am mainly wondering about this from the standpoint of, “man I really just sucked on this particular day.” 
   
If the person interviewing you was bad, this is ok too, but my primary motivation is to learn from others, who are otherwise great engineers that had an off day, in an interview you went it expecting to crush. 

What, if any, feedback was given to you by the interviewer? Was there a certain way that you prepared, or practice projects that you focused on to regain your confidence, and completely rebound the next time around? Finally, what would you say is the number 1 thing that is specific to cpp interviews that you dread?

Thank you to anyone who feels like answering, I love hearing about, and learning from other people’s experience. It is a great resource to be able to have a community of experienced engineers to learn from, and I am grateful for everyone’s time.
## [11][When creating a side project, why choose C++?](https://www.reddit.com/r/cpp/comments/etuvu4/when_creating_a_side_project_why_choose_c/)
- url: https://www.reddit.com/r/cpp/comments/etuvu4/when_creating_a_side_project_why_choose_c/
---
First of all, I realize that there are a ton of cool things you can make in C++. This [site](https://github.com/danistefanovic/build-your-own-x) lists a bunch; for example, you can create a game engine, a game, a database, a compiler, or a web server. However, you can make these things in lots of other languages as well; I'm guessing most people aren't implementing their web servers in C++, and many use a game engine like Unity or GameMaker Studio for games (not a completely fair comparison, since these aren't languages). 

My question is: why do you choose to work on side projects *specifically in C++*? The main reason I can think of, which is the primary reason to work on most side projects, is to learn. I think that's a great reason! It's enough justification to work on any side project in any language. However, for many other languages, there are often secondary reasons.  For example, if I wanted to create a game, Python or Javascript (not to mention actual game frameworks) would let me iterate much more quickly than C++. If I wanted to create a command line tool, I would consider Rust over C++ because of the safety guarantees.

Overall, the idea I'm trying to get across is that C++ isn't a particularly practical language to use for side projects. Using it for the sake of learning is a great reason, but I can't think of any compelling secondary reasons. This doesn't hold true for many other languages, where a secondary reason might be development speed, safety, or abundance of frameworks (e.g. Node for JS).
