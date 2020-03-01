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
## [2][is_printable: Check if a value is printable at compile-time with modern C++](https://www.reddit.com/r/cpp/comments/fblqwd/is_printable_check_if_a_value_is_printable_at/)
- url: https://www.reddit.com/r/cpp/comments/fblqwd/is_printable_check_if_a_value_is_printable_at/
---
I came up with a helper metafunction named `is_printable` that can be used to check if a value can be printed, at compile-time. Please have a read about it at [https://nafe.es/posts/2020-02-29-is-printable/](https://nafe.es/posts/2020-02-29-is-printable/).

Looking forward to your questions, suggestions, and thoughts. Thank you!
## [3][Unity build test with Meson &amp; LibreOffice](https://www.reddit.com/r/cpp/comments/fbmp0w/unity_build_test_with_meson_libreoffice/)
- url: https://nibblestew.blogspot.com/2020/02/unity-build-test-with-meson-libreoffice.html
---

## [4][Klein: A SIMD-optimized C++17 Geometry Library](https://www.reddit.com/r/cpp/comments/fb8y1n/klein_a_simdoptimized_c17_geometry_library/)
- url: /r/gamedev/comments/fb8wbq/klein_a_simdoptimized_c17_geometry_library/
---

## [5][What does XSTRING mean?](https://www.reddit.com/r/cpp/comments/fbsmov/what_does_xstring_mean/)
- url: https://www.reddit.com/r/cpp/comments/fbsmov/what_does_xstring_mean/
---
Hi. I was writing a very simple code with arrays to try them after I watched a tutorial, but at a certain point the code stops executing and Visual Studio opens another file, called xstring.

It's more or less 4000 lines long and I can't understand anything of it (even though I can clearly see it's C++).

Where can I learn more about xstrings (since I didn't find any tutorial online) and what the hell are they?

Thanks
## [6][Relaxing the One Definition Rule in Interpreted C++](https://www.reddit.com/r/cpp/comments/fbf0dg/relaxing_the_one_definition_rule_in_interpreted_c/)
- url: https://dl.acm.org/doi/abs/10.1145/3377555.3377901
---

## [7][Meaning of Ref implementations](https://www.reddit.com/r/cpp/comments/fbcajx/meaning_of_ref_implementations/)
- url: https://www.reddit.com/r/cpp/comments/fbcajx/meaning_of_ref_implementations/
---
I've seen several `template &lt;typename T&gt; class Ref` implementations in various code bases (For instance [this one](https://dawn.googlesource.com/dawn/+/refs/heads/master/src/dawn_native/RefCounted.h) used e.g. [here](https://dawn.googlesource.com/dawn/+/refs/heads/master/src/dawn_native/ComputePipeline.h)) and I always wondered what's the reasoning behind this? Why not simply declare `ShaderModuleBase &amp;mModule` as a private class member if the author wants to enforce that it references  _something_ as soon as the class is instantiated?
## [8][The C++ rvalue lifetime disaster, by Arno Schödl, CoreHard Autumn 2019](https://www.reddit.com/r/cpp/comments/fbag4p/the_c_rvalue_lifetime_disaster_by_arno_schödl/)
- url: https://www.youtube.com/watch?v=s9vBk5CxFyY
---

## [9][Should there be a standard C++ pattern for this? transform_to | The Old New Thing](https://www.reddit.com/r/cpp/comments/fawq17/should_there_be_a_standard_c_pattern_for_this/)
- url: https://devblogs.microsoft.com/oldnewthing/20200228-00/?p=103498
---

## [10][Is MinGW linker way too slow?](https://www.reddit.com/r/cpp/comments/fb3ylq/is_mingw_linker_way_too_slow/)
- url: https://www.reddit.com/r/cpp/comments/fb3ylq/is_mingw_linker_way_too_slow/
---
I have a relatively small project that makes use of some templated libraries like Cereal. MinGW linker times are really slow for me, especially in debug. It takes over 20 seconds to link the program. Release is bearable (4-5 seconds), but the whole project links in under a second with msvc.  I tried a bunch of linker options, but nothing seems to really help (besides reducing the debug info level with -g flag, but that kinda defeats the point of the debug build). Is MinGW just that slow?  How does clang comapre on Windows?
## [11][A quick primer on type traits in modern C++](https://www.reddit.com/r/cpp/comments/fauhzf/a_quick_primer_on_type_traits_in_modern_c/)
- url: https://www.internalpointers.com/post/quick-primer-type-traits-modern-cpp
---

