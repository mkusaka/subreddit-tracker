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
## [2][Advanced C/C++ side projects](https://www.reddit.com/r/cpp/comments/huh7d2/advanced_cc_side_projects/)
- url: https://www.reddit.com/r/cpp/comments/huh7d2/advanced_cc_side_projects/
---
Hi.

After a resume review by a fellow engineer, I was told that I should build some advanced C and C++ projects to show my competency.  Although my long term goal is to become a sensor fusion engineer, now I am just looking for entry-level positions.  The industry I want to get into is self-drivig cars/robotics/autonomy development.  What are some advanced C and C++ based projects that I can build along the lines of this industry?  These projects are also going to help me learn the languages a lot better as well.

I apologize if this is a dumb post. I figured that this is a good place to ask.

Thanks.
## [3][Looking for an UI framework easy to integrate with C++ for mobile platforms](https://www.reddit.com/r/cpp/comments/huhppm/looking_for_an_ui_framework_easy_to_integrate/)
- url: https://www.reddit.com/r/cpp/comments/huhppm/looking_for_an_ui_framework_easy_to_integrate/
---
Hi guys,

I've been working for the last couple of years in a codebase in C++.

 I did manage to crosscompile to different platforms: iOS &amp; Android and now I'm looking forward to work in the frontend of the application.

For the desktop application, I'm using Qt and it works just fine for all desktop platforms.

I've been doing some research about the different technologies that can be used for the mobile UIs and I would like to know about your experiences or any recomendation.

I'm considering:
- Qt (QML) and try to get the startup/indie license
- Flutter 
- Native development
- Whatever JavaScript framework

How hard was the integration of the C/C++ code with the framework? Limitations? Any license issues?

I want to create a serverless isolated application, so implementing a back-end service and query the data is something I'm not considering. 

Thanks
## [4][Looking for an argument parsing library](https://www.reddit.com/r/cpp/comments/hudlgw/looking_for_an_argument_parsing_library/)
- url: https://www.reddit.com/r/cpp/comments/hudlgw/looking_for_an_argument_parsing_library/
---
Any suggestions on argument parser for C++ just like argparse for Python?
## [5][Have you completed an Embedded C/C++ Internship?](https://www.reddit.com/r/cpp/comments/huhvnr/have_you_completed_an_embedded_cc_internship/)
- url: https://www.reddit.com/r/cpp/comments/huhvnr/have_you_completed_an_embedded_cc_internship/
---
Has anyone here ever taken part of an embedded c/c++ internship? I need some help finding tutorials for this, as I don't really know what should I search for. Please help if you do know.
## [6][Clang OpenMP for Loops](https://www.reddit.com/r/cpp/comments/hu9yl7/clang_openmp_for_loops/)
- url: https://www.reddit.com/r/cpp/comments/hu9yl7/clang_openmp_for_loops/
---
I wrote some code that needs parallelized (just a for loop), and am trying to make sure my code builds on GCC and Clang. I wrote my code on an Ubuntu system (20.04) and everything builds fine with both compilers. When I tried to build it on my Gentoo system, it builds fine with GCC, but fails with Clang. This is particularly odd since both systems have exactly the same versions of the compilers (GCC 10.0.1 and Clang 10.0.0 with libomp-10). Here is the relevant section of my code:

    #pragma omp parallel for default(shared)
    for (auto v = my_vector.begin(); v &lt; my_vector.end(); ++v) {
      // use v
    }

On my Gentoo system, if I use Clang it complains that I am not using a relational operator on \`v\`, and that I should use \`&lt;, &lt;=, &gt;, or &gt;=\`. If I re-write the loop to use indexing instead of iterators, it works fine.

libomp-10 has partial support for OpenMP 5.0, including support for ranged for loops. If I rewrite my code to use a range, I get exactly the same error.

Any idea why this error is happening?
## [7][Digging in member method pointers in C++](https://www.reddit.com/r/cpp/comments/hukjrf/digging_in_member_method_pointers_in_c/)
- url: https://www.reddit.com/r/cpp/comments/hukjrf/digging_in_member_method_pointers_in_c/
---
Hi guys

Recently  was dealing with member method pointers in C++  
I tried to deal with it a little bit deep, and somehow my brain generated an article.

[https://medium.com/@helloooooooooooooooooooooooooo/method-member-pointers-in-c-6a2e0190b3e0](https://medium.com/@helloooooooooooooooooooooooooo/method-member-pointers-in-c-6a2e0190b3e0)

Hope it will be useful
## [8][Good youtube playlist for someone who already knows C very well](https://www.reddit.com/r/cpp/comments/htxlyp/good_youtube_playlist_for_someone_who_already/)
- url: https://www.reddit.com/r/cpp/comments/htxlyp/good_youtube_playlist_for_someone_who_already/
---
Im going to learn C++ my next year of college, and I was wondering what are some good yt playlist to learn it, im already very good at C so if there are any that playlists that skip the basic stuff feel free to share them. Thanks!
## [9][Quick setup for a new C++ project in Visual Studio Code](https://www.reddit.com/r/cpp/comments/hu00cg/quick_setup_for_a_new_c_project_in_visual_studio/)
- url: https://github.com/IvanSafonov/cpp-boilerplate
---

## [10][Aku - Toy Backtesting/Trading Engine in C++.](https://www.reddit.com/r/cpp/comments/htwmiq/aku_toy_backtestingtrading_engine_in_c/)
- url: https://github.com/flouthoc/aku
---

## [11][Open sourcing my CQC automation platform](https://www.reddit.com/r/cpp/comments/hti9zs/open_sourcing_my_cqc_automation_platform/)
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
