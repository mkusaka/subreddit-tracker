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
## [2][Async C++ with fibers](https://www.reddit.com/r/cpp/comments/ir87xf/async_c_with_fibers/)
- url: https://www.reddit.com/r/cpp/comments/ir87xf/async_c_with_fibers/
---
I would like to ask the community to share their thoughts and experience on building I/O bound C++ backend services on fibers (stackfull coroutines).

Asynchronous  responses/requests/streams (thinking of grpc-like server service) cycle is quite difficult to write in C++.

Callback-based (like original boost.asio approach) is quite a mess: difficult to reason about lifetimes, program flow and error handling.

C++20 Coroutines are not quite here and one needs to have some experience to rewrite "single threaded" code to coroutine based. And here is also a dangling reference problem could exist.

The last approach is fibers. It seems very easy to think about and work with (like boost.fibers). One writes just a "single threaded" code, which under the hood turned into interruptible/resumable code. The program flow and error handlings are the same like in the single threaded program.

&amp;#x200B;

What do you think about fibers approach to write i/o bound services? Did I forget some fibers drawbacks that make them not so attractive to use?
## [3][Any progress on clang concepts?](https://www.reddit.com/r/cpp/comments/irbrus/any_progress_on_clang_concepts/)
- url: https://www.reddit.com/r/cpp/comments/irbrus/any_progress_on_clang_concepts/
---
[short-circuiting](https://godbolt.org/z/E6r9z3) still not working and only std::same_as is implemented in [&lt;concepts&gt;](https://github.com/llvm/llvm-project/blob/master/libcxx/include/concepts), which makes it quite unusable.
## [4][How to perform automated performance tests on a framework?](https://www.reddit.com/r/cpp/comments/ir7ute/how_to_perform_automated_performance_tests_on_a/)
- url: https://www.reddit.com/r/cpp/comments/ir7ute/how_to_perform_automated_performance_tests_on_a/
---
Hi guys,

&amp;#x200B;

I am working on a bigger c++ software module. I am using GTest for testing, and valgrind for checking for memory leaks. I am wondering if there is a framework that allows me to write timing tests for high-performance code.

&amp;#x200B;

Any suggestions would be appreciated!
## [5][For what inputs is std::sort unstable?](https://www.reddit.com/r/cpp/comments/iqwnug/for_what_inputs_is_stdsort_unstable/)
- url: https://quuxplusone.github.io/blog/2020/09/11/unstable-sort-inputs/
---

## [6][Herb Sutter: My plans at CppCon](https://www.reddit.com/r/cpp/comments/iqxhhi/herb_sutter_my_plans_at_cppcon/)
- url: https://herbsutter.com/2020/09/11/my-plans-at-cppcon/
---

## [7][Tips on progression](https://www.reddit.com/r/cpp/comments/ir8nb9/tips_on_progression/)
- url: https://www.reddit.com/r/cpp/comments/ir8nb9/tips_on_progression/
---
I know a good amount of libraries, how to set things up manage memory and can do lots with SFML. Should I start making bigger projects using my knowledge of graphics programming and basic c++? I’m just kind of out of things to learn. I suppose there’s more GUI design to learn and implementation of two languages at once.
## [8][Return to C++ after a few years of Java](https://www.reddit.com/r/cpp/comments/iqo7z6/return_to_c_after_a_few_years_of_java/)
- url: https://www.reddit.com/r/cpp/comments/iqo7z6/return_to_c_after_a_few_years_of_java/
---
After about 6 years of professional C++ development, including top tech companies (think FAANG), I decided to move to finance. Since 2017, I've been working for one of the top-tier investment banks, developing trading systems in Java. While I'm totally enjoying working in this field, I'm not particularly happy about writing in Java and am missing C++ a lot.

Unfortunately, my bank has nearly zero C++ openings, and even those are legacy projects. I am thinking to switch back to C++ by getting a job at another bank/prop trading/hedge fund.

My main question - what is the best way to brush up my C++ after a few years of Java? I certainly haven't forgotten the basics, always use C++ for my Leetcode practice, but apparently that couldn't replace real work experience. Should I try and start a C++ pet project as a preparation / read Effective C++ / ... something else?

Also, the last standard that I used in my work was C++11 (with a bit of '14). I have almost no experience with '17. What's the best way to  learn it, provided I'm not a complete noob?

Thank you.
## [9][Parsing floats in C++: benchmarking strtod vs. from_chars](https://www.reddit.com/r/cpp/comments/iqlnt4/parsing_floats_in_c_benchmarking_strtod_vs_from/)
- url: https://lemire.me/blog/2020/09/10/parsing-floats-in-c-benchmarking-strtod-vs-from_chars/
---

## [10][CppCon 2020 starts next week and new keynote speaker has been published](https://www.reddit.com/r/cpp/comments/iqgqmt/cppcon_2020_starts_next_week_and_new_keynote/)
- url: https://www.reddit.com/r/cpp/comments/iqgqmt/cppcon_2020_starts_next_week_and_new_keynote/
---
Within 3 days to start CppCon, **Emery Berger** had been presented as keynote speaker with the talk titled: "**Performance Matters"**

[https://cppcon.org/cppcon-2020-keynote-performance-matters-by-emery-berger/](https://cppcon.org/cppcon-2020-keynote-performance-matters-by-emery-berger/)

This talk is scheduled on **Wednesday**, September 16 •    10:30 am- 11:30 am (MDT)
## [11][A few details and thoughts on online events like Meeting C++ online](https://www.reddit.com/r/cpp/comments/iqmukd/a_few_details_and_thoughts_on_online_events_like/)
- url: http://meetingcpp.com/meetingcpp/news/items/A-few-details-and-thoughts-on-online-events-like-Meeting-Cpp-online.html
---

