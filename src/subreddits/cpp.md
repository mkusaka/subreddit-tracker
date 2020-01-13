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
* **Don't** use URL shorteners. [reddiquette](https://www.reddit.com/wiki/reddiquette) forbids them because they're opaque to the spam filter.
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
## [2][Influence of C++ on Swift](https://www.reddit.com/r/cpp/comments/eo10jo/influence_of_c_on_swift/)
- url: https://www.quora.com/What-are-similarities-and-differences-between-C-and-Swift
---

## [3][C++ Profiling with Perf and Flamegraphs](https://www.reddit.com/r/cpp/comments/enxeew/c_profiling_with_perf_and_flamegraphs/)
- url: http://www.mycpu.org/flamegraphs-on-c++/
---

## [4][Better Algorithm Intuition - Conor Hoekstra - code::dive 2019](https://www.reddit.com/r/cpp/comments/eo2mtl/better_algorithm_intuition_conor_hoekstra/)
- url: https://www.youtube.com/watch?v=0z-cv3gartw
---

## [5][The New ConanCenter Improves Search and Discovery](https://www.reddit.com/r/cpp/comments/eo4ivz/the_new_conancenter_improves_search_and_discovery/)
- url: https://blog.conan.io/2020/01/13/New-Conan-Center-Improves-Search-Discovery.html
---

## [6][latest c++ theory vs reality](https://www.reddit.com/r/cpp/comments/enuntq/latest_c_theory_vs_reality/)
- url: https://www.reddit.com/r/cpp/comments/enuntq/latest_c_theory_vs_reality/
---
Why is there a seemingly big discrepancy of existing and used c++ version? Despite good compiler support.

When I see the new features of current c++ currently (17 and even 20 has wide compiler support already) i am amazed.

But in reality:
Open recent and active c++ project on github.
It’s 98 or 11 in the best case
Why on earth is every c++ code I (with my limited horizon) find in the wild uses such old standards?
Also clang uses c++98 as default setting. Which is somewhat okay, because it can be easily adjusted. But it’s weird nonetheless.
## [7][How does module linking works?](https://www.reddit.com/r/cpp/comments/enzvn1/how_does_module_linking_works/)
- url: https://www.reddit.com/r/cpp/comments/enzvn1/how_does_module_linking_works/
---
So i've read about modules, but I can't wrap my head around how they're linked

Are compiled modules stored as a custom format or good ol' .so / .dll files? Since either way seems to break the structure of using ld (at least on Linux) to load a shared object dynamically and using some sort a header as a map of that shared object.

I am curious as to how modules will be handled as dynamic objects
## [8][Should there be a compiler flag to pretend const_cast doesn't exist?](https://www.reddit.com/r/cpp/comments/enmo8c/should_there_be_a_compiler_flag_to_pretend_const/)
- url: https://www.reddit.com/r/cpp/comments/enmo8c/should_there_be_a_compiler_flag_to_pretend_const/
---
Considering how the possibility of const_casts can make certain optimisations impossible, it seems like there's a potential benefit from having a flag that can tell the compiler to assume that nothing will ever have its constness casted away.

There are plenty of applications where you can be confident that no const_casting will ever happen, so to have a small amount of performance lost because of a rarely-used feature is a bit frustrating.

It's not like there's no precedent for this sort of thing either: The -ffast-math flag allows extra optimisations by assuming that no float can ever be NaN or Infinity, and that arguably has a much more significant effect.

The biggest problem I see is that it's not easy to be sure if a const-cast will happen when you call someone else's code (or the standard library implementation you're using) so a flag like this might break something you're not aware of. Maybe that would stop this from ever being useful.

Anyway, those are just some thoughts. What do you think?
## [9][A header-only library for easy interactive parameters](https://www.reddit.com/r/cpp/comments/enq4lc/a_headeronly_library_for_easy_interactive/)
- url: https://github.com/sharkdp/painless
---

## [10][A quirk of qualified member lookup](https://www.reddit.com/r/cpp/comments/enhvos/a_quirk_of_qualified_member_lookup/)
- url: https://quuxplusone.github.io/blog/2020/01/11/qualified-member-lookup/
---

## [11][Cpp-Taskflow v2.3.0 pre-release: A new Conditional Tasking interface for cyclic and dynamic control flows](https://www.reddit.com/r/cpp/comments/enk1s5/cpptaskflow_v230_prerelease_a_new_conditional/)
- url: https://github.com/cpp-taskflow/cpp-taskflow
---

