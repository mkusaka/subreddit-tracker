# cpp
## [1][C++ Jobs - Q4 2020](https://www.reddit.com/r/cpp/comments/j3qems/c_jobs_q4_2020/)
- url: https://www.reddit.com/r/cpp/comments/j3qems/c_jobs_q4_2020/
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

* [C++ Jobs - Q3 2020](https://www.reddit.com/r/cpp/comments/hjnaf2/c_jobs_q3_2020/)
## [2][Ridiculously fast unicode (UTF-8) validation](https://www.reddit.com/r/cpp/comments/jf0tx4/ridiculously_fast_unicode_utf8_validation/)
- url: https://lemire.me/blog/2020/10/20/ridiculously-fast-unicode-utf-8-validation/
---

## [3][json_dto-0.2.11: custom formatters added to an easy to use wrapper around RapidJSON](https://www.reddit.com/r/cpp/comments/jf82ka/json_dto0211_custom_formatters_added_to_an_easy/)
- url: https://eao197.blogspot.com/2020/10/progc-jsondto-0211-released.html
---

## [4][2020-10 C++ Committee Mailing](https://www.reddit.com/r/cpp/comments/jf4wsw/202010_c_committee_mailing/)
- url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/#mailing2020-10
---

## [5][What to use for writing large documentations?](https://www.reddit.com/r/cpp/comments/jexkdk/what_to_use_for_writing_large_documentations/)
- url: https://www.reddit.com/r/cpp/comments/jexkdk/what_to_use_for_writing_large_documentations/
---
This is not strictly CPP, but hey, I like you guys.

I'm planning some projects and will need to have quite large and extensive documentation.

Immediately I thought I'd just make a git repo full of markdown files for documentation as it's not going to be about specific functions but instead overviews of technology and systems (It will be updated before code so I wasn't worried about it being out of date).  
If I did that I could just host it with a github IO page or even just leave it as a repo.

Then I thought a little further and decided maybe I should use an actual system to track it since it would be handy to have more automatic references and such. Some documentation tracker / generator.

So do you guys have any suggestions for writing docs for large systems? Repos, doc generators, a CMS?



A little further clarification:  
The code itself will be commented like usual and I'll use Doxygen or something to generate that data.  
The documentation I need to write is more of an overview of the whole code. End users are not touching the code for the project, just using it. So it will not be "this function does X", rather it's "Press this button to do Y".  
That's why I'm not worried about one being out of date, they are separate pieces of data (implementation vs behavior).

Edit:
The documentation is going to be very technical, just not actual code.  
I also think I would prefer for the documentation to be in git not in a database.
## [6][Does anybody use std::list?](https://www.reddit.com/r/cpp/comments/jf3dwo/does_anybody_use_stdlist/)
- url: https://www.reddit.com/r/cpp/comments/jf3dwo/does_anybody_use_stdlist/
---
I have never used std::list the last 15 years,  mainly for performance reasons. If I need to store large objects, I use a vector of pointers. Does anybody have a use case that can share? 
Thanks!
## [7][Performance of std::pmr](https://www.reddit.com/r/cpp/comments/jf0dse/performance_of_stdpmr/)
- url: https://www.reddit.com/r/cpp/comments/jf0dse/performance_of_stdpmr/
---
I tried to experiment a bit by replacing several container types with their pmr:: counterparts on a medium-sized project. This change alone made the performance drop substantially. From around ~~180~~158 fps to 145 fps.

I am really surprised by this. With no additional changes, it should still delegate to new/delete, but does so via a virtual call. Could this indirection be the source of the overhead? Or could it be the extra pointer which is needed in all the containers?

This library is mainly single-threaded, so I replaced the default memory resource with an [unsynchronized pool](https://en.cppreference.com/w/cpp/memory/unsynchronized_pool_resource), and then the performance was back up again closer to the original 180 fps. I was hoping to see something much better here when I started, but it seems that the gain is completely eaten up by the overhead.

Does anyone have any experience using the polymorphic allocators on a larger codebase? Any idea what is going on here?

Edit: In the process, I accidentally changed out a non-std unordered map with `std::pmr::unordered_map`. This turned out to be the largest part of the performance drop, the `std::unordered_map` is just a whole lot slower. After correcting this, there is still a 10-15 fps performance drop purely due to using the pmr containers.
## [8][Warrior1: A Performance Sanitizer for C++](https://www.reddit.com/r/cpp/comments/jeqgvu/warrior1_a_performance_sanitizer_for_c/)
- url: https://arxiv.org/abs/2010.09583
---

## [9][Qt 6.0 Beta Released](https://www.reddit.com/r/cpp/comments/jelc9l/qt_60_beta_released/)
- url: https://www.qt.io/blog/qt-6.0-beta-released
---

## [10][I got "me knickers in a twist" while trying to compile and link using VS Code and clang that is part of Visual Studio 2019. But I was cool calm and collected and this is my story.](https://www.reddit.com/r/cpp/comments/jfa1x2/i_got_me_knickers_in_a_twist_while_trying_to/)
- url: https://github.com/dbj-data/bench/tree/main/.vscode/readme
---

## [11][macOS performance tester wanted (eCAL)](https://www.reddit.com/r/cpp/comments/jf9vmq/macos_performance_tester_wanted_ecal/)
- url: https://www.reddit.com/r/cpp/comments/jf9vmq/macos_performance_tester_wanted_ecal/
---
For our pub/sub middleware we are looking for a performance tester for macOS. I will spend a bottle of one of the best local beer in Frankfurt Rhein Main area (shipping world wide), maybe two ;-)

[Taunus Naturtrueb](https://imgur.com/YwchS0H)

What we would like to have is a performance measuring on a native Mac following these [instructions](https://continental.github.io/ecal/advanced/performance.html). The dmg darwin image you can find [here](https://github.com/continental/ecal/releases). Happy testing !
