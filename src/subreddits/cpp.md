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
## [2][LTL: A C++ 17 library to write functional code](https://www.reddit.com/r/cpp/comments/iyt4wo/ltl_a_c_17_library_to_write_functional_code/)
- url: https://www.reddit.com/r/cpp/comments/iyt4wo/ltl_a_c_17_library_to_write_functional_code/
---
I would like to present you a C++17 library I am working on since few months.

The idea is to bring the power of C++ ranges / range-v3 with some improvements, like the `&gt;&gt;` notation, and extends it to the `Option` monad and the `Error` monad.

Some differences with the C++ range :
  1. Supports rvalue references
  2. Supports `&gt;&gt;` notation
  3. Extends the _pipelining_ of algorithms to `std::optional` and an home craft `expected`
  4. Supports some action like finding a value inside a container
  5. Supports function composition
  6. There is far more functionality in ranges-v3 than in my implementation
  7. There is probably some bugs and missing functionality that will be fixed / added in a near future

I hope you will like it, and don't hesitate to give feedback :)

https://github.com/qnope/Little-Type-Library/
## [3][C++ for fun ... ctional programmers... ???](https://www.reddit.com/r/cpp/comments/iyivts/c_for_fun_ctional_programmers/)
- url: https://www.youtube.com/watch?v=2vJfJE4K0zg
---

## [4][The Hidden Secrets of Move Semantics - Nicolai Josuttis - CppCon 2020](https://www.reddit.com/r/cpp/comments/iyr0wt/the_hidden_secrets_of_move_semantics_nicolai/)
- url: https://www.youtube.com/watch?v=TFMKjL38xAI
---

## [5][New C++ features in GCC 10](https://www.reddit.com/r/cpp/comments/iyvgge/new_c_features_in_gcc_10/)
- url: https://developers.redhat.com/blog/2020/09/24/new-c-features-in-gcc-10/
---

## [6][How expensive is integer overflow trapping in C++?](https://www.reddit.com/r/cpp/comments/iyqz18/how_expensive_is_integer_overflow_trapping_in_c/)
- url: https://lemire.me/blog/2020/09/23/how-expensive-is-integer-overflow-trapping-in-c
---

## [7][ZenMake - A cross-platform build system for C/C++ and some other languages](https://www.reddit.com/r/cpp/comments/iytf9g/zenmake_a_crossplatform_build_system_for_cc_and/)
- url: https://www.reddit.com/r/cpp/comments/iytf9g/zenmake_a_crossplatform_build_system_for_cc_and/
---
Hi people! I develop my own build system for C/C++ and some other languages.

ZenMake is a build system based on the meta build system/framework [Waf](https://waf.io/). The main purpose of ZenMake is to be as simple to use as possible but remain flexible. And I will be glad if my project will be useful not only for me.

Why yet another build system? Yes, I know, we already have a lot of them. I decided to create this project because I couldn’t find a build tool for Linux which is quick and easy to use, flexible, ready to use, with declarative configuration, without the need to learn one more special language and suitable for my needs. Full description can be found here: [https://zenmake.readthedocs.io/en/latest/why.html](https://zenmake.readthedocs.io/en/latest/why.html)

Primary git repository of the project: [https://gitlab.com/pustotnik/zenmake](https://gitlab.com/pustotnik/zenmake)

Documentation: [https://zenmake.readthedocs.io](https://zenmake.readthedocs.io/)
## [8][Tips &amp; Tricks for Effective SYCL Development (SYCL Summer Sessions 2020)](https://www.reddit.com/r/cpp/comments/iysxim/tips_tricks_for_effective_sycl_development_sycl/)
- url: https://www.youtube.com/watch?v=4Jo-Sb1AX_c
---

## [9][C++20 ranges lazy set_intersection](https://www.reddit.com/r/cpp/comments/iyw7yk/c20_ranges_lazy_set_intersection/)
- url: https://www.reddit.com/r/cpp/comments/iyw7yk/c20_ranges_lazy_set_intersection/
---
It seems like there is no lazy version of set\_intersection and other similar algorithms that operate on ranges. However, rangev3 seems to have one under views::set\_intersection. Why is this discrepancy?
## [10][Collaborative C++ Development with Visual Code - Julia Reid - CppCon 2020](https://www.reddit.com/r/cpp/comments/iyt7dw/collaborative_c_development_with_visual_code/)
- url: https://youtu.be/qCn6zT76zTs
---

## [11][Back to Basics: Class Layout - Stephen Dewhurst - CppCon 2020](https://www.reddit.com/r/cpp/comments/iyi16h/back_to_basics_class_layout_stephen_dewhurst/)
- url: https://www.youtube.com/watch?v=SShSV_iV1Ko
---

