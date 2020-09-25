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
## [2][Meeting C++ online - Basit Ayantunde - Functional Error and Optional value Handling with STX](https://www.reddit.com/r/cpp/comments/izfknm/meeting_c_online_basit_ayantunde_functional_error/)
- url: https://www.youtube.com/watch?v=8CZhJa8UJk0
---

## [3][Plenary: The Beauty and Power of "Primitive" C++ - Bjarne Stroustrup - CppCon 2020](https://www.reddit.com/r/cpp/comments/iziw57/plenary_the_beauty_and_power_of_primitive_c/)
- url: https://youtu.be/ERzENfQ51Ck
---

## [4][New C++ features in GCC 10](https://www.reddit.com/r/cpp/comments/iyvgge/new_c_features_in_gcc_10/)
- url: https://developers.redhat.com/blog/2020/09/24/new-c-features-in-gcc-10/
---

## [5][Do you have a preferred library for constexpr data structures?](https://www.reddit.com/r/cpp/comments/izhmre/do_you_have_a_preferred_library_for_constexpr/)
- url: https://www.reddit.com/r/cpp/comments/izhmre/do_you_have_a_preferred_library_for_constexpr/
---
Do y'all have a preferred library of constexpr data structures? I'm trying to find something like the embedded STL [https://www.etlcpp.com/](https://www.etlcpp.com/)  but with constexpr support. I was wondering if there were any  alternatives. I've found quite a few header only implementation of  single data structures, but not a solid colllection. This also isn't  strictly embedded, but I assume the people likely to use this are people  working in embedded. 

I'm working on designing a real time system, and so I'm trying to put as much of the work at compile time. Not necessarily embedded, but doing any unnecessary work at runtime is not preferred. 

Also while I'm here, given that information, should I implement a threadpool, scheduler, and event system myself? I wasn't able to find a thread pool with considerations like this (liberal use of shared\_ptr on a few implmentations), so I'm curious to see if anyone could point me in the right direction.
## [6][Any recommendations for a decent and free memory leak detection tool?](https://www.reddit.com/r/cpp/comments/iz4w17/any_recommendations_for_a_decent_and_free_memory/)
- url: https://www.reddit.com/r/cpp/comments/iz4w17/any_recommendations_for_a_decent_and_free_memory/
---
For context, I have Visual Leak Detector, but it crashes my application. Not sure if this is due to incompatibility with C++17 (I'm running version 2.5.1) but it makes me sad. I also tried Dr. Memory, but am also getting an internal crash with that too! I'm filled with sorrow.

Edit: Since I'm a fool and failed to mention it, I am on Windows with MSVC (in case you couldn't tell from the fact that I'm using VLD).
## [7][std::exchange Patterns: Fast, Safe, Expressive, and Probably Underused](https://www.reddit.com/r/cpp/comments/izjgdw/stdexchange_patterns_fast_safe_expressive_and/)
- url: https://www.fluentcpp.com/2020/09/25/stdexchange-patterns-fast-safe-expressive-and-probably-underused/
---

## [8][LTL: A C++ 17 library to write functional code](https://www.reddit.com/r/cpp/comments/iyt4wo/ltl_a_c_17_library_to_write_functional_code/)
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
## [9][Generating compile_cemmands.json from eclipse project](https://www.reddit.com/r/cpp/comments/izdbi4/generating_compile_cemmandsjson_from_eclipse/)
- url: https://www.reddit.com/r/cpp/comments/izdbi4/generating_compile_cemmandsjson_from_eclipse/
---
How can i generate compile_commands.json for eclipse c/c++ project.
I used use bear and compiledb for other c projects that has make files.

 Eclipse also generate make files but i couldn't generate compile commands from that make file. Since eclipse completely builds the project, if i clean the build it will also clean make files

EDIT: solved it by just adding `bear` behind make in project builder settings
## [10][The Hidden Secrets of Move Semantics - Nicolai Josuttis - CppCon 2020](https://www.reddit.com/r/cpp/comments/iyr0wt/the_hidden_secrets_of_move_semantics_nicolai/)
- url: https://www.youtube.com/watch?v=TFMKjL38xAI
---

## [11][C++ for fun ... ctional programmers... ???](https://www.reddit.com/r/cpp/comments/iyivts/c_for_fun_ctional_programmers/)
- url: https://www.youtube.com/watch?v=2vJfJE4K0zg
---

