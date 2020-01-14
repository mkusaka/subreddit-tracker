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
## [2][Little text game I made in C++](https://www.reddit.com/r/cpp/comments/eokn4p/little_text_game_i_made_in_c/)
- url: https://gamejolt.com/games/robotlife/434756
---

## [3][Speeding up C++ GitHub Actions using ccache](https://www.reddit.com/r/cpp/comments/eoctdc/speeding_up_c_github_actions_using_ccache/)
- url: https://cristianadam.eu/20200113/speeding-up-c-plus-plus-github-actions-using-ccache/
---

## [4][Why isn't there an RAII guard for 20's new semaphore types?](https://www.reddit.com/r/cpp/comments/eobqfh/why_isnt_there_an_raii_guard_for_20s_new/)
- url: https://www.reddit.com/r/cpp/comments/eobqfh/why_isnt_there_an_raii_guard_for_20s_new/
---
It doesn't seem like it would be too complex (following `std::lock_guard`) to have a type that does acquire/release. But cppreference has no mention of such a type. Was it considered for the standard?
## [5][The New ConanCenter Improves Search and Discovery](https://www.reddit.com/r/cpp/comments/eo4ivz/the_new_conancenter_improves_search_and_discovery/)
- url: https://blog.conan.io/2020/01/13/New-Conan-Center-Improves-Search-Discovery.html
---

## [6][Influence of C++ on Swift](https://www.reddit.com/r/cpp/comments/eo10jo/influence_of_c_on_swift/)
- url: https://www.quora.com/What-are-similarities-and-differences-between-C-and-Swift
---

## [7][Better Algorithm Intuition - Conor Hoekstra - code::dive 2019](https://www.reddit.com/r/cpp/comments/eo2mtl/better_algorithm_intuition_conor_hoekstra/)
- url: https://www.youtube.com/watch?v=0z-cv3gartw
---

## [8][abbreviated class templates?](https://www.reddit.com/r/cpp/comments/eoiscy/abbreviated_class_templates/)
- url: https://www.reddit.com/r/cpp/comments/eoiscy/abbreviated_class_templates/
---
since we now have abbreviated function templates in c++20 (with concepts), I wonder if we could have something similar for structs/classes, concretely

&amp;#x200B;

`struct test {`

`auto x;`

`auto y;`

`test(auto _x, auto _y) : x{ _x }, y{ _y } {}`

`};`

`auto obj = test{ 1, "aaa"s };`

`// _x deduced to int, _y deduced to std::string as function template deduction for the ctor.`

`// obj.x then deduced to int since it is initialized by _x, and _x is int, likewise for obj.y`

the whole thing should be equivalent to

`template&lt;typename T1, typename T2&gt;`

`struct test {`

`T1 x;`

`T2 y;`

`template&lt;typename T3, typename T4&gt;`

`test(T3 _x, T4 _y) : x{ _x }, y{ _y } {}`

`};`

`template&lt;typename...T&gt;`  
`test(T...)-&gt;test&lt;T...&gt;;`
## [9][Towards deprecating volatile: A Simplified volatile read and write implementation in C++](https://www.reddit.com/r/cpp/comments/eofdxu/towards_deprecating_volatile_a_simplified/)
- url: https://www.reddit.com/r/cpp/comments/eofdxu/towards_deprecating_volatile_a_simplified/
---
I have implemented a simplified header-only library to help make performing volatile read and writes easier in C++ codebases.
Really handy especially if you're into embedded systems.

It's a few lines but well documented and won't have to repeat yourself nor use the volatile keyword irresponsibly again. 

More documentation and justification on the repo. 

Contributions and reviews are highly welcome too 😀


https://github.com/lamarrr/volta
## [10][C++ Profiling with Perf and Flamegraphs](https://www.reddit.com/r/cpp/comments/enxeew/c_profiling_with_perf_and_flamegraphs/)
- url: http://www.mycpu.org/flamegraphs-on-c++/
---

## [11][latest c++ theory vs reality](https://www.reddit.com/r/cpp/comments/enuntq/latest_c_theory_vs_reality/)
- url: https://www.reddit.com/r/cpp/comments/enuntq/latest_c_theory_vs_reality/
---
Why is there a seemingly big discrepancy of existing and used c++ version? Despite good compiler support.

When I see the new features of current c++ currently (17 and even 20 has wide compiler support already) i am amazed.

But in reality:
Open recent and active c++ project on github.
It’s 98 or 11 in the best case
Why on earth is every c++ code I (with my limited horizon) find in the wild uses such old standards?
Also clang uses c++98 as default setting. Which is somewhat okay, because it can be easily adjusted. But it’s weird nonetheless.
