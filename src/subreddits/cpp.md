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
## [2][I Made an Esoteric Programming Language with ANTLR4 and C++ which let's you code in Cats' language.](https://www.reddit.com/r/cpp/comments/hvsdax/i_made_an_esoteric_programming_language_with/)
- url: https://github.com/bauripalash/mewmew
---

## [3][Improving code generation time with C++ Build Insights](https://www.reddit.com/r/cpp/comments/hvi7l8/improving_code_generation_time_with_c_build/)
- url: https://devblogs.microsoft.com/cppblog/improving-code-generation-time-with-cpp-build-insights/
---

## [4][Doxygen comments and Logpoints support for VS Code C++ Extension](https://www.reddit.com/r/cpp/comments/hvc1gw/doxygen_comments_and_logpoints_support_for_vs/)
- url: https://devblogs.microsoft.com/cppblog/visual-studio-code-c-extension-july-2020-update-doxygen-comments-and-logpoints/
---

## [5][grpc_bench: a clear, objective gRPC benchmark](https://www.reddit.com/r/cpp/comments/hvs4f0/grpc_bench_a_clear_objective_grpc_benchmark/)
- url: /r/grpc/comments/hv7h6i/grpc_bench_a_clear_objective_grpc_benchmark/
---

## [6][How Relocations and Thread Local Store are Implemented](https://www.reddit.com/r/cpp/comments/hvi1s2/how_relocations_and_thread_local_store_are/)
- url: https://stffrdhrn.github.io/software/toolchain/openrisc/2020/07/21/relocs_tls_impl.html
---

## [7][std::ref and std::reference_wrapper: common use cases](https://www.reddit.com/r/cpp/comments/hv6d9k/stdref_and_stdreference_wrapper_common_use_cases/)
- url: https://www.nextptr.com/tutorial/ta1441164581/stdref-and-stdreference_wrapper-common-use-cases
---

## [8][[Question] What are convincing (preferably perf-unrelated) ways to argue against passing everything by shared_ptr?](https://www.reddit.com/r/cpp/comments/hv45lw/question_what_are_convincing_preferably/)
- url: https://www.reddit.com/r/cpp/comments/hv45lw/question_what_are_convincing_preferably/
---
Since C++11, I noticed that programmers in some circles started using `std::shared_ptr` for everything (possibly type-aliasing it as e.g.`Ptr`) and passing everything by `shared_ptr` copy, but this makes me personally uneasy due to the loss in visibility into RAII-based resource release times.

Unfortunately, I'm not good at coming up with the kind of concrete arguments that truly speak to people, so to say, so here I am.

Most notably, I'm having a hard time conveying why destruction times are Important™, and why you don't want your object relationships to turn into a strongly connected graph.

What's your guys' take? How would you (amiably) approach the issue?

Edit: Thank you all for the responses! I didn't go through quite all of them yet, but I will over the next day or so. I didn't expect this to get as much attention as it did, but I'm happy that it did :)
## [9][Macro for accessors methods](https://www.reddit.com/r/cpp/comments/hvor0t/macro_for_accessors_methods/)
- url: https://www.reddit.com/r/cpp/comments/hvor0t/macro_for_accessors_methods/
---
Hello,

In a new project where I am working now there are a lot of macros that help to write accessors and other things.
As example you can write GETSET(int, counter) in a class definition an the macro will exapnds into the code of set / get and member definition.
I don't like very much this but you don't have to write a lot of trivial code.

What do you think? Is it a bad practice? Or you like this approach?
## [10][Better IDE for linux](https://www.reddit.com/r/cpp/comments/hvelig/better_ide_for_linux/)
- url: https://www.reddit.com/r/cpp/comments/hvelig/better_ide_for_linux/
---
I am using g++ through terminal, it works fine for small projects but for complicated project like that using cmake , g++ gave me little choices compared to VS on Windows. What is the best IDE to use in Ubuntu?
## [11][static code analyzer for annotated TODO comments \w C++ support](https://www.reddit.com/r/cpp/comments/hv8p56/static_code_analyzer_for_annotated_todo_comments/)
- url: https://github.com/preslavmihaylov/todocheck
---

