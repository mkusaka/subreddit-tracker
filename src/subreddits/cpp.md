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
## [2][Recursive Lambdas in C++](https://www.reddit.com/r/cpp/comments/irupel/recursive_lambdas_in_c/)
- url: https://artificial-mind.net/blog/2020/09/12/recursive-lambdas
---

## [3][The Most Popular Programming Languages - 1965/2020](https://www.reddit.com/r/cpp/comments/irfs7m/the_most_popular_programming_languages_19652020/)
- url: https://youtu.be/UNSoPa-XQN0
---

## [4][Any way to "merge" a const and non-const getter?](https://www.reddit.com/r/cpp/comments/irva9b/any_way_to_merge_a_const_and_nonconst_getter/)
- url: https://www.reddit.com/r/cpp/comments/irva9b/any_way_to_merge_a_const_and_nonconst_getter/
---
Like, now i have it setup like this:

    MyClass&amp; Get(int index) 
    {
    return ...;
    }
    
    const MyClass&amp; Get(int index) const
    {
    return...;
    }

Is there any way to merge the logic? Like something like this:

    MyClass&amp; Get(int index) 
    {
        return Get(index); // Explicitly call the const variant from below and return as non-const
    }
    
    const MyClass&amp; Get(int index) const
    {
    return...;
    }

Because right now i basically have all the logic from these getters duplicated.. which is shit
## [5][The Design of C++ , lecture by Bjarne Stroustrup](https://www.reddit.com/r/cpp/comments/irissi/the_design_of_c_lecture_by_bjarne_stroustrup/)
- url: https://youtu.be/69edOm889V4
---

## [6][CppCon: Classic STL Class Announcement](https://www.reddit.com/r/cpp/comments/irlm0w/cppcon_classic_stl_class_announcement/)
- url: https://quuxplusone.github.io/blog/2020/09/12/classic-stl-at-cppcon-2020/
---

## [7][lunasvg v1.2.0 released](https://www.reddit.com/r/cpp/comments/irore5/lunasvg_v120_released/)
- url: https://github.com/sammycage/lunasvg/releases/tag/v1.2.0
---

## [8][cmkr: Generate CMake from TOML](https://www.reddit.com/r/cpp/comments/ircxnr/cmkr_generate_cmake_from_toml/)
- url: https://www.reddit.com/r/cpp/comments/ircxnr/cmkr_generate_cmake_from_toml/
---
Hello

I was wondering how it would look like using a declarative language to generate CMakeLists. This repo is still a WIP:

[https://github.com/MoAlyousef/cmkr](https://github.com/MoAlyousef/cmkr)

It doesn't handle complex logic because of TOML's nature, however it should be easier for tooling to integrate into.  
The tool, for the moment, can init a cmake.toml project, generate CMakeLists and run cmake builds.
## [9][Async C++ with fibers](https://www.reddit.com/r/cpp/comments/ir87xf/async_c_with_fibers/)
- url: https://www.reddit.com/r/cpp/comments/ir87xf/async_c_with_fibers/
---
I would like to ask the community to share their thoughts and experience on building I/O bound C++ backend services on fibers (stackfull coroutines).

Asynchronous  responses/requests/streams (thinking of grpc-like server service) cycle is quite difficult to write in C++.

Callback-based (like original boost.asio approach) is quite a mess: difficult to reason about lifetimes, program flow and error handling.

C++20 Coroutines are not quite here and one needs to have some experience to rewrite "single threaded" code to coroutine based. And here is also a dangling reference problem could exist.

The last approach is fibers. It seems very easy to think about and work with (like boost.fibers). One writes just a "single threaded" code, which under the hood turned into interruptible/resumable code. The program flow and error handlings are the same like in the single threaded program.

&amp;#x200B;

What do you think about fibers approach to write i/o bound services? Did I forget some fibers drawbacks that make them not so attractive to use?
## [10][Any progress on clang concepts?](https://www.reddit.com/r/cpp/comments/irbrus/any_progress_on_clang_concepts/)
- url: https://www.reddit.com/r/cpp/comments/irbrus/any_progress_on_clang_concepts/
---
[short-circuiting](https://godbolt.org/z/E6r9z3) still not working and only std::same_as is implemented in [&lt;concepts&gt;](https://github.com/llvm/llvm-project/blob/master/libcxx/include/concepts), which makes it quite unusable.
## [11][How to perform automated performance tests on a framework?](https://www.reddit.com/r/cpp/comments/ir7ute/how_to_perform_automated_performance_tests_on_a/)
- url: https://www.reddit.com/r/cpp/comments/ir7ute/how_to_perform_automated_performance_tests_on_a/
---
Hi guys,

&amp;#x200B;

I am working on a bigger c++ software module. I am using GTest for testing, and valgrind for checking for memory leaks. I am wondering if there is a framework that allows me to write timing tests for high-performance code.

&amp;#x200B;

Any suggestions would be appreciated!
