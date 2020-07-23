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
## [2][GCC 10.2 Released](https://www.reddit.com/r/cpp/comments/hwc5sc/gcc_102_released/)
- url: https://www.reddit.com/r/cpp/comments/hwc5sc/gcc_102_released/
---
GCC 10.2 is a bug-fix release from the GCC 10 branch containing important fixes for regressions and serious bugs in GCC 10.1 with more than 94 bugs fixed since the previous release. Bug list: [https://gcc.gnu.org/bugzilla/buglist.cgi?bug\_status=RESOLVED&amp;resolution=FIXED&amp;target\_milestone=10.2](https://gcc.gnu.org/bugzilla/buglist.cgi?bug_status=RESOLVED&amp;resolution=FIXED&amp;target_milestone=10.2)
## [3][New 3.1.4 release of wxWidgets, free and open source library for creating portable native GUI applications, adding macOS 11/ARM support and much more, is now available.](https://www.reddit.com/r/cpp/comments/hvyk11/new_314_release_of_wxwidgets_free_and_open_source/)
- url: http://wxwidgets.org/news/2020/07/wxwidgets-3.1.4-released/
---

## [4][nonstd::uniform_real_distribution, a drop-in replacement for std::uniform_real_distribution (msvc) with proper mathematical properties](https://www.reddit.com/r/cpp/comments/hw4cdl/nonstduniform_real_distribution_a_dropin/)
- url: https://www.reddit.com/r/cpp/comments/hw4cdl/nonstduniform_real_distribution_a_dropin/
---
The standard is in a dilemma regarding `std::uniform_real_distribution(a, b)`: On one side it's asking for the distr. to generate uniformly distributed values on [a, b), on the other it's dictating an implementation which results in *non*-uniform values being generated on [a, b], meaning it's inconsistent and unfulfillable at this point.

Microsoft opts for the dictated implementation forfeiting uniformity. Seeing as nobody can be standard-conform anyway, true uniformity was more important for us, let me present [`nonstd::uniform_real_distribution`](https://github.com/denniskb/uniform_real_distr), which is a (close to) drop-in replacement for `std::uniform_real_distribution`, the only differences being:
- We do not support (de-)serialization from/to streams.
- We only support float and double for now.

Advantages over `std`:
- We generate perfectly distributed uniform values on [a, b) for a=0. (All values are equally far spaced apart and have an equal probability of occurring.)
- We generate values on [min(), max()] with min()==a and max()&lt;b for a=0.
- We guarantee to invoke your RNG at most `ceil({24|53} / log2(rng.max()-rng.min()+1))` times.
- We are just as fast as `std::uniform_real_distribution`.
- We can also generate right-inclusive values on (a, b] (for a=0) if desired, simply instantiate with `right_inc` equal to `true`: `nonstd::uniform_real_distribution&lt;float, true&gt; iid;`

This header was extracted from a current, thoroughly tested production-level codebase. If there's sufficient interest I'd be happy to extract the corresponding unit tests as well.
## [5][C++ on Sea 2020 video - "A refresher on the floating-point computations and the standard library" - Boguslaw Cyganek](https://www.reddit.com/r/cpp/comments/hwbpgj/c_on_sea_2020_video_a_refresher_on_the/)
- url: https://www.youtube.com/watch?v=7aZbYJ5UTC8
---

## [6][July update on COVID-19 and Meeting C++ 2020](https://www.reddit.com/r/cpp/comments/hwbr8z/july_update_on_covid19_and_meeting_c_2020/)
- url: http://meetingcpp.com/meetingcpp/news/items/July-update-on-COVID-19-and-Meeting-Cpp-2020.html
---

## [7][July 2020 C++ standard mailing](https://www.reddit.com/r/cpp/comments/hw2s5k/july_2020_c_standard_mailing/)
- url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/#mailing2020-07
---

## [8][C++ on Sea 2020 video - "Hypercritical C++ Code Review" - Yuri Minaev](https://www.reddit.com/r/cpp/comments/hwe16k/c_on_sea_2020_video_hypercritical_c_code_review/)
- url: https://www.youtube.com/watch?v=f1_Iwh33f9I
---

## [9][I Made an Esoteric Programming Language with ANTLR4 and C++ which let's you code in Cats' language.](https://www.reddit.com/r/cpp/comments/hvsdax/i_made_an_esoteric_programming_language_with/)
- url: https://github.com/bauripalash/mewmew
---

## [10][C++ on Sea video - "Quickly and Effectively Testing Legacy C++ Code with Approval Tests" - Clare Macrae](https://www.reddit.com/r/cpp/comments/hwcywb/c_on_sea_video_quickly_and_effectively_testing/)
- url: https://www.youtube.com/watch?v=tXEuf_3VzRE
---

## [11][C++ on Sea 2020 video: "C++20: All the small things" - Fergus Cooper](https://www.reddit.com/r/cpp/comments/hvxjyr/c_on_sea_2020_video_c20_all_the_small_things/)
- url: https://www.youtube.com/watch?v=LL_NrM7MY44
---

