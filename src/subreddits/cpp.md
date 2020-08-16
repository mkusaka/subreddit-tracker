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
## [2][Why has reddit suspended Bjarne Stroustrup's account? Accounts are only suspended manually so somebody at reddit looked at u/bstroustrup's account and decided to suspend it.](https://www.reddit.com/r/cpp/comments/ia6u5t/why_has_reddit_suspended_bjarne_stroustrups/)
- url: https://www.reddit.com/user/bstroustrup
---

## [3][Throwing Out the Kitchen Sink - Output Ranges](https://www.reddit.com/r/cpp/comments/iaesk7/throwing_out_the_kitchen_sink_output_ranges/)
- url: https://thephd.github.io/output-ranges
---

## [4][Asio vs Boost.Asio](https://www.reddit.com/r/cpp/comments/iai9ae/asio_vs_boostasio/)
- url: https://www.reddit.com/r/cpp/comments/iai9ae/asio_vs_boostasio/
---
Are there any reasons to use Boost.Asio over Asio?

I prefer to use only headers in my project but are there any disadvantages to this specifically if I'm writing a game server?

Thank you
## [5][A compiler/os detecting library for modern C++](https://www.reddit.com/r/cpp/comments/iadas4/a_compileros_detecting_library_for_modern_c/)
- url: https://www.reddit.com/r/cpp/comments/iadas4/a_compileros_detecting_library_for_modern_c/
---
I wrote a small single-header and header-only C++ library for detecting the host compiler/os using `if` and/or `if constexpr` (no need for `#if`/`#endif`). An example is included on `README.md`. The library supports some of the most used operating systems nowadays. Help/suggestions are appreciated.
## [6][Are small memory allocations in C++ STL containers still a concern?](https://www.reddit.com/r/cpp/comments/iaervc/are_small_memory_allocations_in_c_stl_containers/)
- url: https://blog.kaetemi.be/2020/08/15/are-small-memory-allocations-in-cpp-stl-containers-still-a-concern/
---

## [7][What are some of the best Cheat Sheets that you use?](https://www.reddit.com/r/cpp/comments/i9q480/what_are_some_of_the_best_cheat_sheets_that_you/)
- url: https://www.reddit.com/r/cpp/comments/i9q480/what_are_some_of_the_best_cheat_sheets_that_you/
---
In my experience having quick references is extremely helpful. What are some of the best cheat sheets that you have used or currently use?
## [8][Easy dependency management for C++ with CMake's FetchContent](https://www.reddit.com/r/cpp/comments/ia43ir/easy_dependency_management_for_c_with_cmakes/)
- url: https://bewagner.github.io/programming/2020/05/02/cmake-fetchcontent/
---

## [9][Bizarre SSE/AVX codegen with MSVC Preview](https://www.reddit.com/r/cpp/comments/i9wa79/bizarre_sseavx_codegen_with_msvc_preview/)
- url: https://www.reddit.com/r/cpp/comments/i9wa79/bizarre_sseavx_codegen_with_msvc_preview/
---
I switched to the latest preview recently because I wanted access to more C++20 features to clean some things up, but looking at the assembly all of my aligned SIMD loads and stores through intrinsics have become unaligned, and no matter what hints I give the compiler it will never generate an aligned load or store. I'm not sure if this is a bug, or if there's a genuine reason to prefer unaligned loads and store instructions now on x86-64, but I'm just putting this out there because I know some of the MSVC guys lurk around here.
## [10][Boost 1.74 is out](https://www.reddit.com/r/cpp/comments/i9i4vi/boost_174_is_out/)
- url: https://www.boost.org/users/history/version_1_74_0.html
---

## [11][TIL: on x86-64 compilers will use addressing mode to replace simple multiplications](https://www.reddit.com/r/cpp/comments/i9mgvf/til_on_x8664_compilers_will_use_addressing_mode/)
- url: https://www.reddit.com/r/cpp/comments/i9mgvf/til_on_x8664_compilers_will_use_addressing_mode/
---
obligatory godbolt: [https://gcc.godbolt.org/z/csEnh4](https://gcc.godbolt.org/z/csEnh4)

i was testing a snippet, and i was delighted to see how the compiler optimizes simple multiplications using the x86-64 addressing mode

here's a little table

| op | ~ trasformation |
| --- | ---|
| var*=2 |  var &lt;&lt;= 1 |
| var*=3 | var = var+var*2 |
| var*=4 | var &lt;&lt;= 2 |
| var*=5 | var = var+var*4 |
| var*=6 | tmp = var*3; var = tmp+tmp |
| var*=7 | var = var*8-var |
| var*=8 | var &lt;&lt;= 3 |
| var*=9 | var = var*8+var |
| var*=10 | tmp = var+var*4; var = tmp+tmp |

and here's a nice explanation: [https://paul.bone.id.au/blog/2018/09/05/x86-addressing/](https://paul.bone.id.au/blog/2018/09/05/x86-addressing/)
