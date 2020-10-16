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
## [2][Google's TensorFlow management is garbage](https://www.reddit.com/r/cpp/comments/jbzywe/googles_tensorflow_management_is_garbage/)
- url: https://www.reddit.com/r/cpp/comments/jbzywe/googles_tensorflow_management_is_garbage/
---
This is a follow-up post to [https://www.reddit.com/r/cpp/comments/hwr6u9/google\_c\_libraries\_are\_garbage](https://www.reddit.com/r/cpp/comments/hwr6u9/google_c_libraries_are_garbage)

It's now been almost 4 months since I posted a PR to the TensorFlow GitHub repo. All I've seen so far, was people being assigned as reviewers, then assigning other people as reviewers. I mean we already have to patch our TensorFlow build in multiple ways to make it function at least somewhat as a proper C++ library. Maybe they're ignoring me because of these Reddit posts, but it's academic at this point I think. Until today, all the activity I've seen was one Google employee begging for others to take action.

Anyway, here's the PR, feel free to take it apart,  critique it, maybe tell me how stupid an incompetent I really am (aka review the PR): [https://github.com/tensorflow/tensorflow/pull/41733](https://github.com/tensorflow/tensorflow/pull/41733)
## [3][GUI libraries with support for RTL languages](https://www.reddit.com/r/cpp/comments/jc7sb2/gui_libraries_with_support_for_rtl_languages/)
- url: https://www.reddit.com/r/cpp/comments/jc7sb2/gui_libraries_with_support_for_rtl_languages/
---
Was wondering if any c++ gui libraries apart from Qt also support RTL languages.
## [4][How to add C++ structured binding support to your own types](https://www.reddit.com/r/cpp/comments/jbwkiy/how_to_add_c_structured_binding_support_to_your/)
- url: https://devblogs.microsoft.com/oldnewthing/20201015-00/?p=104369
---

## [5][Writing terminal moo: On C++, its ecosystem, performance optimization and more](https://www.reddit.com/r/cpp/comments/jc8y8q/writing_terminal_moo_on_c_its_ecosystem/)
- url: https://github.com/s9w/terminal_moo/blob/master/text/text.md
---

## [6][LLVM 11.0.0 Released - highlights for C/C++ developers](https://www.reddit.com/r/cpp/comments/jbn7cp/llvm_1100_released_highlights_for_cc_developers/)
- url: https://tobias.hieta.se/llvm11-release
---

## [7][Part 2: Upsetting Opinions about Static Analyzers](https://www.reddit.com/r/cpp/comments/jc71hl/part_2_upsetting_opinions_about_static_analyzers/)
- url: https://isocpp.org/blog/2020/10/part-2-upsetting-opinions-about-static-analyzers
---

## [8][Does this code triggers undefined behavior?](https://www.reddit.com/r/cpp/comments/jc4kkq/does_this_code_triggers_undefined_behavior/)
- url: /r/cpp_questions/comments/jbiuwd/does_this_code_triggers_undefined_behavior/
---

## [9][CppCast: Bazel](https://www.reddit.com/r/cpp/comments/jc28af/cppcast_bazel/)
- url: https://cppcast.com/bazel/
---

## [10][non-const ref copy ctor should be defaulted if the const ref copy ctor is available](https://www.reddit.com/r/cpp/comments/jc69ob/nonconst_ref_copy_ctor_should_be_defaulted_if_the/)
- url: https://www.reddit.com/r/cpp/comments/jc69ob/nonconst_ref_copy_ctor_should_be_defaulted_if_the/
---
The non-const ref copy ctor `T(T&amp;)` should be defaulted to `T(T&amp; x): T{ const_cast&lt;const T&amp;&gt;(x) } {}` if it is not defined by the user and the const ref copy ctor `T(const T&amp;)` is available, whether the const ref copy ctor is defined by the user or automatically generated by the compiler, otherwise, unexpected behavior might happen when there's a templated ctor...

see:  [https://godbolt.org/z/vo6xda](https://godbolt.org/z/vo6xda)
## [11][C Linked-List Visualization (GDBFrontend v0.3.0-beta)](https://www.reddit.com/r/cpp/comments/jbx28j/c_linkedlist_visualization_gdbfrontend_v030beta/)
- url: https://www.youtube.com/watch?v=z44KJDYZOoE
---

