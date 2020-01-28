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
## [2][Light and Powerful modern web-framework for C++](https://www.reddit.com/r/cpp/comments/ev2obp/light_and_powerful_modern_webframework_for_c/)
- url: https://oatpp.io/
---

## [3][Ninja v1.10.0 released!](https://www.reddit.com/r/cpp/comments/ev2zng/ninja_v1100_released/)
- url: https://github.com/ninja-build/ninja/releases
---

## [4][STL header token parsing benchmarks for libstdc++ 7, 8 and 9](https://www.reddit.com/r/cpp/comments/ev4e7g/stl_header_token_parsing_benchmarks_for_libstdc_7/)
- url: https://www.reddit.com/r/cpp/comments/ev4e7g/stl_header_token_parsing_benchmarks_for_libstdc_7/
---
To match [the results for VS2019 posted yesterday](https://www.reddit.com/r/cpp/comments/eumou7/stl_header_token_parsing_benchmarks_for_vs2017/) where we found that the VS2019 STL has become 9% lighter in preprocessor parsing time than the VS2017 STL, here are the same for libstdc++ 7, 8 and 9:

Graph of libstdc++-9: https://raw.githubusercontent.com/ned14/stl-header-heft/master/graphs/libstdc++-9.png

Comparative graph: https://raw.githubusercontent.com/ned14/stl-header-heft/master/graphs/libstdc++-history.png

Detailed notes: https://github.com/ned14/stl-header-heft/blob/master/Readme.libstdc++.md

Project github: https://github.com/ned14/stl-header-heft

You may remember from [the last test performed two years ago](https://www.reddit.com/r/cpp/comments/83rf8o/stl_header_token_parsing_benchmarks_for_libstdc_5/?st=jf0t6mob&amp;sh=817e7a37) that libstdc++ had engorged itself between libstdc++-5 and libstdc+-7 by **+25.7%**. I can report that libstdc++-9 has further engorged itself over libstdc++-7 by **+16%**. That's almost linear with time, which I find fascinating.

There is good news here though. Whilst the average and median has become much worse, all of the pathological individual headers in libstdc++-9 have been fixed e.g. `array` and `iterator` were dragging in lots of unneeded stuff which they no longer do. In this sense, both the libstdc++ and VS STLs have greatly improved in the past two years by not sub-including unnecessary headers. My congratulations to the libstdc++ authors and maintainers for this work!

If there was any one area where I think libstdc++ is particularly deficient, it is that any of its threading headers basically drag in the exact same stuff, most of which isn't actually needed by the individual header. MSVC's STL proves that this isn't necessary, `&lt;thread&gt;` is much lighter weight than any of the other threading headers on MSVC, but is the heaviest weight threading header on libstdc++, for example.

All that said, libstdc++ has *far* more very light weight headers than the MSVC STL e.g. `&lt;type_traits&gt;` on libstdc++ really is a class leader compared to the hefty `&lt;type_traits&gt;` on MSVC. MSVC *could* do much better here, a huge include for MSVC is `&lt;sal.h&gt;` which absolutely everything drags in, and is very low hanging fruit to prune in my opinion.
## [5][stdgpu 1.2.0 released!](https://www.reddit.com/r/cpp/comments/ev5ctx/stdgpu_120_released/)
- url: https://github.com/stotko/stdgpu/releases/tag/1.2.0
---

## [6][Qt restricts the LTS releases to commercial customers, :(](https://www.reddit.com/r/cpp/comments/eupd06/qt_restricts_the_lts_releases_to_commercial/)
- url: https://www.qt.io/blog/qt-offering-changes-2020
---

## [7][C++ Modules conformance improvements with MSVC in Visual Studio 2019 16.5](https://www.reddit.com/r/cpp/comments/eus3j6/c_modules_conformance_improvements_with_msvc_in/)
- url: https://devblogs.microsoft.com/cppblog/c-modules-conformance-improvements-with-msvc-in-visual-studio-2019-16-5/
---

## [8][Difference between string and stringstream](https://www.reddit.com/r/cpp/comments/ev4ttw/difference_between_string_and_stringstream/)
- url: https://www.reddit.com/r/cpp/comments/ev4ttw/difference_between_string_and_stringstream/
---
Hi! I'm taking a course on C++ so I'm still on the basics. I couldn't grasp the difference between a stringstream and a string. Would anyone care to explain? Thank you!
## [9][Hacking on Clang is surprisingly easy](https://www.reddit.com/r/cpp/comments/eumv9p/hacking_on_clang_is_surprisingly_easy/)
- url: https://mort.coffee/home/clang-compiler-hacking/
---

## [10][STL header token parsing benchmarks for VS2017 and VS2019](https://www.reddit.com/r/cpp/comments/eumou7/stl_header_token_parsing_benchmarks_for_vs2017/)
- url: https://www.reddit.com/r/cpp/comments/eumou7/stl_header_token_parsing_benchmarks_for_vs2017/
---
One of the MSVC devs poked me to update the results from [https://www.reddit.com/r/cpp/comments/860hya/stl\_header\_token\_parsing\_benchmarks\_for\_vs2008/](https://www.reddit.com/r/cpp/comments/860hya/stl_header_token_parsing_benchmarks_for_vs2008/) for VS2019:

Graph of VS2019: [https://raw.githubusercontent.com/ned14/stl-header-heft/master/graphs/msvs-2019.png](https://raw.githubusercontent.com/ned14/stl-header-heft/master/graphs/msvs-2019.png)

Comparative graph: [https://raw.githubusercontent.com/ned14/stl-header-heft/master/graphs/msvs-history.png](https://raw.githubusercontent.com/ned14/stl-header-heft/master/graphs/msvs-history.png)

Detailed notes: [https://github.com/ned14/stl-header-heft/blob/master/Readme.msvs.md](https://github.com/ned14/stl-header-heft/blob/master/Readme.msvs.md)

Project github: [https://github.com/ned14/stl-header-heft](https://github.com/ned14/stl-header-heft)

There is a lot of good news in this benchmark: overall VS2019 has 9% lower token processing times than VS2017 did. **That makes VS2019 almost as quick as VS2008 used to be!**

The single biggest surprise is surely now `&lt;array&gt;`, which is by far the biggest impact STL container now that Microsoft have greatly improved `&lt;string&gt;` and especially `&lt;vector&gt;` and `&lt;forward_list&gt;`.

I'll try to produce updated benchmarks for libstdc++ during this coming week.

I'd like to take this opportunity to thank the Visual C++ team for such outstanding work on slimming down their STL implementation, yet simultaneously implementing a large proportion of C++ 20! Great work guys, wish more STLs were like yours!
## [11][OpenSSL client and server from scratch, part 3](https://www.reddit.com/r/cpp/comments/ev3vbq/openssl_client_and_server_from_scratch_part_3/)
- url: https://quuxplusone.github.io/blog/2020/01/26/openssl-part-3/
---

