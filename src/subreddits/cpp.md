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
## [2][Introducing fp library](https://www.reddit.com/r/cpp/comments/j8h81u/introducing_fp_library/)
- url: https://www.reddit.com/r/cpp/comments/j8h81u/introducing_fp_library/
---
Hi, the author of [Dragonbox](https://github.com/jk-jeon/dragonbox) again. I'm now working on a more comprehensive [library](https://github.com/jk-jeon/fp) with features related to conversions between strings and floating-point numbers. It's in an early stage and is not really usable right now AFAICT, but I want to gather some opinions from you guys. Please take a look at the [repo](https://github.com/jk-jeon/fp) if you are interested. Any idea/feedback is very welcome, and contribution is even more welcome.

As stated in the README, the goal of the library is to provide conversion algorithms between decimal and binary floating-point  numbers rather than string parsing/formatting. String parsing/formatting routines provided in the library should be considered as proof-of-concepts, although I'll continue to work on improving them as well. Therefore, I put much effort on cleanly separate decimal/binary conversion routines from string parsing/formatting as much as possible.

Currently, the library has the following features (with rudimentary/subject-to-change API's):

1. **Dragonbox**: shortest-roundtrip binary-to-decimal conversion algorithm for formatting. The version in this library is not really different from the original implementation, except that it uses a larger table (in order to share it with the parsing algorithm).
2. **Ryu-printf**: fixed-precision binary-to-decimal conversion algorithm for formatting. This is nothing more than just an implementation of the algorithm as written in the [paper](https://dl.acm.org/doi/pdf/10.1145/3360595) with a few modifications, but somewhat mysteriously it uses way fewer amount of static data compared to the original implementation by Ulf Adams. (I honestly have no idea why this happened.) But the size of static data is still huge; it's about 40KB.
3. **Dooly**: limited-precision (9 for `float`'s, 17 for `double`'s) decimal-to-binary conversion algorithm for parsing. This is the one that I mentioned in the [previous post](https://www.reddit.com/r/cpp/comments/ishdj9/updates_in_dragonbox/) about Dragonbox. Ryu library also provides a similar algorithm, and it seems that the performances of these algorithms are similar.
4. Combining Dooly and (a slight extension of) Ryu-printf, string-to-float parsing for arbitrary length input is also possible. A gist of the idea is explained in the comment of /u/STL in [this post](https://www.reddit.com/r/cpp/comments/iqlnt4/parsing_floats_in_c_benchmarking_strtod_vs_from/). (I originally got the idea from discussions with fast\_io's author, though.) /u/STL said that it might be agonizingly slow, but it turned out to be a lot faster than the current implementation of MS STL's  `std::from_chars`. (To be fair, my implementation does way fewer error handlings so the actual performance gap should be somewhat smaller.)
## [3][Crux: Symbolic Testing for C, C++, and Rust](https://www.reddit.com/r/cpp/comments/j8iu5v/crux_symbolic_testing_for_c_c_and_rust/)
- url: https://crux.galois.com/
---

## [4][Inheritance is for sharing an interface (and so is overloading)](https://www.reddit.com/r/cpp/comments/j87m1s/inheritance_is_for_sharing_an_interface_and_so_is/)
- url: https://quuxplusone.github.io/blog/2020/10/09/when-to-derive-and-overload/
---

## [5][periodic_function v0.1.0 released! Thank you all for the helpful feedback!](https://www.reddit.com/r/cpp/comments/j83v8l/periodic_function_v010_released_thank_you_all_for/)
- url: https://www.reddit.com/r/cpp/comments/j83v8l/periodic_function_v010_released_thank_you_all_for/
---
A while back I made my initial [post](https://www.reddit.com/r/cpp/comments/iun0qh/periodicfunction_a_small_header_only_library_to/?utm_source=share&amp;utm_medium=web2x&amp;context=3) about the `periodic_function` library I had created and I receive a **ton** of feedback.

*Thank you so much to everyone who commented.* Your suggestions and critical eyes have made the library more robust and better overall. The suggestions and comments have been addressed in the `v0.1.0` release and I hope some of you can find it useful. 

Thank you again! I really appreciate those who took the time to nitpick my code; it was eye-opening and a great learning experience.

Release is [here](https://github.com/DeveloperPaul123/periodic-function/releases/tag/v0.1.0).
## [6][Compile time conditional struct/class member variable in C++](https://www.reddit.com/r/cpp/comments/j8gesp/compile_time_conditional_structclass_member/)
- url: https://www.reddit.com/r/cpp/comments/j8gesp/compile_time_conditional_structclass_member/
---
[https://medium.com/@saleem.iitg/c-compile-time-conditional-struct-member-variables-c3761a35eca9](https://medium.com/@saleem.iitg/c-compile-time-conditional-struct-member-variables-c3761a35eca9)
## [7][Potential way to solve dangling reference to temporary?](https://www.reddit.com/r/cpp/comments/j89kzz/potential_way_to_solve_dangling_reference_to/)
- url: https://www.reddit.com/r/cpp/comments/j89kzz/potential_way_to_solve_dangling_reference_to/
---
When I am writing [when() in SugarPP](https://github.com/HO-COOH/SugarPP), I kept wondering how to deal with returning reference of possibly lvalue references and rvalue reference. To make it simpler, it's the same as
```cpp
int x{};
auto&amp;&amp; bigger = std::max(x, -1);
//... Use bigger, OK, but dangling if -1 is bigger
```
I gave up by just return by value. But it's obviously not optimal if I also want to be able to modify through the reference (both the lvalue or the temporary), as if the temporary is also lvalue.

Of course I can just give it a name to trivially solve it but still it's amazingly unintuitive. What if the compiler can implicitly generate a name to make the temporary to a lvalue when the function accepts reference (because temporary will still "materialize" anyway) at the current calling scope, but still interpret it as rvalues like this:
```cpp
int x{};
int __temp = -1;
auto&amp;&amp; bigger = std::max(x, std::move(__temp));
//... Use bigger, OK now
```
## [8][toml++ v2.2.0 released](https://www.reddit.com/r/cpp/comments/j7wg45/toml_v220_released/)
- url: https://marzer.github.io/tomlplusplus/
---

## [9][cpp.chat : episode 75 - 'I Really like Sugar', with Conor Hoekstra](https://www.reddit.com/r/cpp/comments/j81djk/cppchat_episode_75_i_really_like_sugar_with_conor/)
- url: https://cpp.chat/75/
---

## [10][What C++ Programmers Need to Know about Header ＜random＞| CppCon 2016: Walter E. Brown](https://www.reddit.com/r/cpp/comments/j802n2/what_c_programmers_need_to_know_about_header/)
- url: https://www.youtube.com/watch?v=6DPkyvkMkk8
---

## [11][Why do Rust fanboys claims that Modern C++ is unsafe?](https://www.reddit.com/r/cpp/comments/j8i51r/why_do_rust_fanboys_claims_that_modern_c_is_unsafe/)
- url: https://www.reddit.com/r/cpp/comments/j8i51r/why_do_rust_fanboys_claims_that_modern_c_is_unsafe/
---

