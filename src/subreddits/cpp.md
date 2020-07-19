# cpp
## [1][C++ Jobs - Q3 2020](https://www.reddit.com/r/cpp/comments/hjnaf2/c_jobs_q3_2020/)
- url: https://www.reddit.com/r/cpp/comments/hjnaf2/c_jobs_q3_2020/
---
Rules For Individuals
---------------------

* **Don't** create top-level comments - those are for employers.
* Feel free to reply to top-level comments with **on-topic** questions.
* I will create top-level comments for **meta** discussion and **individuals looking for work**.

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

\*\*Remote:\*\* [Do you offer the option of working remotely (permanently, or for the duration of the pandemic)? If so, do you require employees to live in certain areas or time zones?]

&amp;nbsp;

\*\*Visa Sponsorship:\*\* [Does your company sponsor visas?]

&amp;nbsp;

\*\*Technologies:\*\* [Required: do you mainly use C++98/03, C++11, C++14, C++17, or C++20? Optional: do you use Linux/Mac/Windows, are there languages you use in addition to C++, are there technologies like OpenGL or libraries like Boost that you need/want/like experience with, etc.]

&amp;nbsp;

\*\*Contact:\*\* [How do you want to be contacted? Email, reddit PM, telepathy, gravitational waves?]

---

Previous Post
--------------

* [C++ Jobs - Q2 2020](https://www.reddit.com/r/cpp/comments/ft77lv/c_jobs_q2_2020/)
## [2][Aku - Toy Backtesting/Trading Engine in C++.](https://www.reddit.com/r/cpp/comments/htwmiq/aku_toy_backtestingtrading_engine_in_c/)
- url: https://github.com/flouthoc/aku
---

## [3][Good youtube playlist for someone who already knows C very well](https://www.reddit.com/r/cpp/comments/htxlyp/good_youtube_playlist_for_someone_who_already/)
- url: https://www.reddit.com/r/cpp/comments/htxlyp/good_youtube_playlist_for_someone_who_already/
---
Im going to learn C++ my next year of college, and I was wondering what are some good yt playlist to learn it, im already very good at C so if there are any that playlists that skip the basic stuff feel free to share them. Thanks!
## [4][Quick setup for a new C++ project in Visual Studio Code](https://www.reddit.com/r/cpp/comments/hu00cg/quick_setup_for_a_new_c_project_in_visual_studio/)
- url: https://github.com/IvanSafonov/cpp-boilerplate
---

## [5][Open sourcing my CQC automation platform](https://www.reddit.com/r/cpp/comments/hti9zs/open_sourcing_my_cqc_automation_platform/)
- url: https://www.reddit.com/r/cpp/comments/hti9zs/open_sourcing_my_cqc_automation_platform/
---
A while back I opened sourced by CIDLib C++ development platform:

[https://github.com/DeanRoddey/CIDLib](https://github.com/DeanRoddey/CIDLib)

This is a large (450K'ish lines of code) general purpose C++ development system. It's basically a virtual OS, written in C++, very powerful and very clean.

Built on top of CIDLib is my CQC automation platform, which is an even bigger (about 650K'ish lines of code), full on, commercial quality home/business automation system. This is something I tried to make a success of for a long time, but ultimately have failed, after sacrificing the bulk of my adult life. That sucks on a fairly cosmic scale, but it is what it is.

It's sort of a poster-boy for why technical people shouldn't make marketing decisions I guess. I selected automation because I thought it was cool and interesting, not because I had any really solid proof that I could sell it. Ultimately, it ended up being an amazing product that just didn't have a commercial market. It's too powerful (aka complex) for the hoi polloi to really be interested in it. It's very well suited technically for professionally installed systems, but there are a number of big players in that space and no way to break into it without big investment that was never going to materialize.

The folks who would be potential customers were mostly folks who won't spend a dime if they can avoid it, i.e. more technical hobbyist types. There are other open source products out there, and those folks would use anything free, no matter how less refined, rather than pay for something.

So, anyhoo, my abject failure is the community's potential gain. I'm moving towards open sourcing it. It may be another month before I'm ready since I have to do things like strip out all of the licensing stuff and make various other tweaks, get the docs updated to reflect those changes, get a repo set up up, build environment documented, etc...

But, I'm heading in that direction and anyone who might be interested in contributing to it as an open product would probably want to be getting familiar with it as a product before trying to contribute as a developer so that you know what it is and how it works. The current commercial version has a 40 day trial period so it should be good until the open source version arrives.

[https://www.youtube.com/user/CharmedQuarkSystems](https://www.youtube.com/user/CharmedQuarkSystems)

[https://www.charmedquark.com/](https://www.charmedquark.com/)

Ultimately, because CQC is built on top of CIDLib, which is incredibly powerful (as evidenced by the fact that I, a single person, could create such a massive product with it), it could be a lot more than an automation system. It really could become a home IT infrastructure system in general. With some more resources it could go in a number of directions.

Currently it is Windows only (in terms of the back end and system management UI, there's a web based touch screen client in addition to the Windows client.) But CIDLib it highly portable. It was designed from day one to support Windows and Linux and actually used to support Linux. And I mean in a very clean way, not conditional code all over the place. That's why it was created as a virtual OS type system.

So, if some Linux folks got on board, we could get the back end cleanly supporting Windows and Linux in a heterogenous network way. I got started on resurrecting the Linux platform support but my Linux skills sort of bottomed out. I have my own build infrastructure which seamlessly supports both platforms, and inherently understands the needs of CIDLib and CQC. This makes it easy to develop on both. I use Visual Studio Code since it's equivalent on both, but that's not a requirement.

Anyhoo, if anyone is interested, start delving into it as a user. I'm going to eat the cost of keeping the web site and forums up. I'll start a new section on the forums for development discussion, so you can sign up there and ask questions, make suggestions, etc...
## [6][What's in a function pointer? Machine code, actually!](https://www.reddit.com/r/cpp/comments/htekog/whats_in_a_function_pointer_machine_code_actually/)
- url: https://youtu.be/I_g7CSoNNxc?t=758
---

## [7][Introduction - redo: a recursive build system](https://www.reddit.com/r/cpp/comments/htjmtj/introduction_redo_a_recursive_build_system/)
- url: https://redo.readthedocs.io/en/latest/
---

## [8][Web framework recommendations](https://www.reddit.com/r/cpp/comments/htht3h/web_framework_recommendations/)
- url: https://www.reddit.com/r/cpp/comments/htht3h/web_framework_recommendations/
---
I've decided to learn some C++ mostly for my own amusement. My hobby projects tend to be web focussed so I was going to rewrite a couple of those in order to learn but I'm having a little bit of an issue finding the right framework. I've come across things like Crow which look like what I'm after, only that doesn't appear to be maintained anymore, likewise Wayward. 

Drogon I think is the leading contender at the moment given WT looks way too heavyweight for what I'm after. Drogon does look more convoluted than I'd ideally like, though I'm not necessarily expecting things to be a walk in the park 

Coming from a python background I'm ideally looking for something in the realms of Flask or Django, ie backend only and delivering either rest or full content via templates.

Are there any others I should look at before I sink my teeth into Drogon or wt?
## [9][Not going to impress you, but I wanted to share: Structuring your code is so rewarding :)](https://www.reddit.com/r/cpp/comments/ht3htx/not_going_to_impress_you_but_i_wanted_to_share/)
- url: https://www.reddit.com/r/cpp/comments/ht3htx/not_going_to_impress_you_but_i_wanted_to_share/
---
So I have this project - biggest one I have worked on so far, but probably small in comparison to any software I have ever heard of.

I am working on this alone, but I will hand it over with full and proper documentation + handover training at the end, with the aim that other people can support and debug it as well.

*Not counting empty lines and pure comment lines, the project now has 33 header files with 1274 loc, and 24 code modules with 4353 loc.* 

*I have a ratio of 0.46 comments per line of code and a Makefile of 168 lines*

I spent the first two months working on the framework:

* Setting up a nice build folder, with subfolders for each module, according obj folder and subfolders, and a manually generated well-structured Makefile without redundant information.
* learning &amp; "proof-of-concepting" the (for me new) technologies I was going to use: snmp agents, boost program\_options, dynamic loading of a plugin (dlopen), atomic variables, function pointers
* implementing the application interaction: basic telemetry from the diagnostics module, inter-application communication via posix message queues, a general logging framework and configuration file logic
* ensuring full compatibility with both gcc on linux and gcc on cygwin64/Windows 10
* building proper test scenarios and test scripts

It was sometimes frustratingly slow, and especially debugging net-snmp mistakes I had in my code was a challenge (documentation is none too great, not even all of the tutorial examples are complete).

**But now**, now I am actually implementing the application functionality and the experience is like none before: I have so many utility functions at my disposal, thanks to the clear structure of the project I always know where to find stuff, and I have parameter checks in my functions left and right so that I always get meaningful error messages if I use a function wrongly.

I can work on a module and change a hundred lines of code in three different interacting functions - I build my project, and once I eliminate all the obvious bugs that the compiler complains about, the executables just work as they are intended to! I hardly ever manage to produce a crash bug - my thread handling is 99% clean and I don't cause race conditions or deadlocks. And when I decide to implement something more complex, it ends up falling apart into many easy steps that can be done with the tools of my framework.

&amp;#x200B;

This post is not intended to pat myself on the back, but rather to re-emphasize

**TL;DR Damn, does a clean and structured project feel rewarding! :)**
## [10][A lib for adding a stacktrace to every c++ exception in windows platform](https://www.reddit.com/r/cpp/comments/ht93d5/a_lib_for_adding_a_stacktrace_to_every_c/)
- url: https://www.reddit.com/r/cpp/comments/ht93d5/a_lib_for_adding_a_stacktrace_to_every_c/
---
 exceptions-stacktrace is a c++17 library which its purpose is to add a stacktrace to exception, even if it was thrown from some external lib you use like stl/boost/any other third party lib, with no code changes to your original code.  

https://twitter.com/ykfre1/status/1284308246472339457?s=20
## [11][C++ Template Library for Probabilistic Programming](https://www.reddit.com/r/cpp/comments/htbv8e/c_template_library_for_probabilistic_programming/)
- url: https://www.reddit.com/r/cpp/comments/htbv8e/c_template_library_for_probabilistic_programming/
---
Hi everyone,

I just wanted to share this library [autoppl](https://github.com/JamesYang007/autoppl) that a couple of my friends and I started for a class final project. We found that there was quite a lack of low-level tools for probabilistic programming and wanted to try making something for C++. I have been recently working on it more and have found it to be pretty successful for some examples. Any comments or feedbacks would be appreciated!
