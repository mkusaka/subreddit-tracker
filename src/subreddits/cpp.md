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
## [2][Klein: A SIMD-optimized C++17 Geometry Library](https://www.reddit.com/r/cpp/comments/fb8y1n/klein_a_simdoptimized_c17_geometry_library/)
- url: /r/gamedev/comments/fb8wbq/klein_a_simdoptimized_c17_geometry_library/
---

## [3][Should there be a standard C++ pattern for this? transform_to | The Old New Thing](https://www.reddit.com/r/cpp/comments/fawq17/should_there_be_a_standard_c_pattern_for_this/)
- url: https://devblogs.microsoft.com/oldnewthing/20200228-00/?p=103498
---

## [4][Meaning of Ref implementations](https://www.reddit.com/r/cpp/comments/fbcajx/meaning_of_ref_implementations/)
- url: https://www.reddit.com/r/cpp/comments/fbcajx/meaning_of_ref_implementations/
---
I've seen several `template &lt;typename T&gt; class Ref` implementations in various code bases (For instance [this one](https://dawn.googlesource.com/dawn/+/refs/heads/master/src/dawn_native/RefCounted.h) used e.g. [here](https://dawn.googlesource.com/dawn/+/refs/heads/master/src/dawn_native/ComputePipeline.h)) and I always wondered what's the reasoning behind this? Why not simply declare `ShaderModuleBase &amp;mModule` as a private class member if the author wants to enforce that it references  _something_ as soon as the class is instantiated?
## [5][A quick primer on type traits in modern C++](https://www.reddit.com/r/cpp/comments/fauhzf/a_quick_primer_on_type_traits_in_modern_c/)
- url: https://www.internalpointers.com/post/quick-primer-type-traits-modern-cpp
---

## [6][The C++ rvalue lifetime disaster, by Arno Schödl, CoreHard Autumn 2019](https://www.reddit.com/r/cpp/comments/fbag4p/the_c_rvalue_lifetime_disaster_by_arno_schödl/)
- url: https://www.youtube.com/watch?v=s9vBk5CxFyY
---

## [7][How I use references](https://www.reddit.com/r/cpp/comments/faszel/how_i_use_references/)
- url: https://cor3ntin.github.io/posts/reference/
---

## [8][C++20 designated initializers](https://www.reddit.com/r/cpp/comments/faudmh/c20_designated_initializers/)
- url: https://mariusbancila.ro/blog/2020/02/27/c20-designated-initializers/
---

## [9][Is MinGW linker way too slow?](https://www.reddit.com/r/cpp/comments/fb3ylq/is_mingw_linker_way_too_slow/)
- url: https://www.reddit.com/r/cpp/comments/fb3ylq/is_mingw_linker_way_too_slow/
---
I have a relatively small project that makes use of some templated libraries like Cereal. MinGW linker times are really slow for me, especially in debug. It takes over 20 seconds to link the program. Release is bearable (4-5 seconds), but the whole project links in under a second with msvc.  I tried a bunch of linker options, but nothing seems to really help (besides reducing the debug info level with -g flag, but that kinda defeats the point of the debug build). Is MinGW just that slow?  How does clang comapre on Windows?
## [10][Crashing clang in 52 bytes: template&lt;class&gt;struct a;a&lt;typename b::template c&lt;&gt;&gt;](https://www.reddit.com/r/cpp/comments/fan0c5/crashing_clang_in_52_bytes_templateclassstruct/)
- url: https://www.reddit.com/r/cpp/comments/fan0c5/crashing_clang_in_52_bytes_templateclassstruct/
---
Managed to distill this down from some deep deep template muck today trying to chain things in weird ways with a promises library. With some help from creduce I got it down from over a million lines preprocessed to maybe 2000 and then just started deleting random things and making sure it still segfaulted each time. GCC handles this correctly; I haven't looked through what's listed in the stack trace but my guess is that I tricked it into doing a dependent name lookup on `b` to see if it has a `c`, but `b` was never declared in the first place so it's a null pointer or something.

Who says C++ can't be concise?

Godbolt for proof: [https://godbolt.org/z/ZrFCcY](https://godbolt.org/z/ZrFCcY)

Edit: ugh just realized it can be 51 if I used `class` instead of `struct`...
## [11][Why is LLVM written entirely in C++?](https://www.reddit.com/r/cpp/comments/fbaqos/why_is_llvm_written_entirely_in_c/)
- url: https://www.reddit.com/r/cpp/comments/fbaqos/why_is_llvm_written_entirely_in_c/
---
(by LLVM I mean LLVM, LLDB, Clang, etc, everything)

sure C++ could be useful or even nessicary for some parts, but the entire thing does not need to be written in C++, and in my opinion it shouldn't be.

compiling takes 200 times longer (I compiled my 50,000 SLOC project vs a 1 character change to libFormat to get this timing, and this was a rebuild, not from scratch; my project took 2 seconds, Clang took 400.)

I assume that's because of the extensive use of templates.

Also, there's all the issues with readability, the libFormat change was near a function called `parseBlock` which has hidden default bool values that are only visible in the UnwrappedLineParser header.

and then there's all the other hidden complexity, operator overloading, function overloading, and probably a bunch more I'm not even thinking of.

Don't get me wrong, I'm not saying Templates shouldn't exist, or operator overloading, or function overloading, or even default variables.

all of those features can be useful.

I'm just saying they should be used sparingly.

so, as the title says, why is/should LLVM be written entirely in C++, instead of a healthy mix of C and C++?
