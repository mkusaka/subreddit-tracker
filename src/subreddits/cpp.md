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
## [2][Asio C++ coroutines in Qt UI, communicating with Asio C++ coroutines in service thread, with deterministic cancellation. My October blog.](https://www.reddit.com/r/cpp/comments/jrg3gi/asio_c_coroutines_in_qt_ui_communicating_with/)
- url: https://cppalliance.org/richard/2020/10/31/RichardsOctoberUpdate.html
---

## [3][GitHub API library for c++](https://www.reddit.com/r/cpp/comments/jr6xsn/github_api_library_for_c/)
- url: https://www.reddit.com/r/cpp/comments/jr6xsn/github_api_library_for_c/
---
Some time ago I was looking for a library which would let me use GitHub's API.  
I could not find any so I wrote my own. Today I have decided to extract it from my project and make it easily reusable so here it is:

[https://github.com/Kicer86/github\_api](https://github.com/Kicer86/github_api)

It's api is very limited. I have just covered my needs, but it can be easily extended.  
I hope you find it useful.
## [4][Butano: a modern C++ high level engine for the GBA](https://www.reddit.com/r/cpp/comments/jrjj2n/butano_a_modern_c_high_level_engine_for_the_gba/)
- url: https://www.reddit.com/r/cpp/comments/jrjj2n/butano_a_modern_c_high_level_engine_for_the_gba/
---
Hi!

Lately, I have been working on [Butano](https://github.com/GValiente/butano), a C++20 engine for the Game Boy Advance.

It brings its own standard library, based on the awesome [ETL](https://www.etlcpp.com/). It differs from the standard library mostly in the usage of asserts instead of exceptions and in the usage of fixed size stack buffers instead of the heap.

Since the GBA doesn't have a file system, all data has to be compiled into the ROM, which allows constexpr to be used with ease. Butano brings its own [asserts system](https://gvaliente.github.io/butano/group__assert.html) which can be used in constant-evaluated contexts too thanks to [std::is\_constant\_evaluated](https://en.cppreference.com/w/cpp/types/is_constant_evaluated).

The engine provides detailed [documentation](https://gvaliente.github.io/butano/), multiple [examples](https://gvaliente.github.io/butano/examples.html) of most aspects of the engine an even the source code and assets of a full game: [Butano Fighter](https://gvaliente.itch.io/butano-fighter).

If you have time please tell me what you think. Thanks!
## [5][A libc written in C++](https://www.reddit.com/r/cpp/comments/jqxan9/a_libc_written_in_c/)
- url: https://llvm.org/docs/Proposals/LLVMLibC.html
---

## [6][braft: An industrial-grade C++ implementation of the RAFT consensus algorithm open sourced by Baidu](https://www.reddit.com/r/cpp/comments/jr6zxm/braft_an_industrialgrade_c_implementation_of_the/)
- url: https://github.com/baidu/braft
---

## [7][rotor v0.10 release (new: timers, request cancellation facility)](https://www.reddit.com/r/cpp/comments/jr39db/rotor_v010_release_new_timers_request/)
- url: https://github.com/basiliscos/cpp-rotor
---

## [8][Structured Concurrency](https://www.reddit.com/r/cpp/comments/jqt4z2/structured_concurrency/)
- url: https://ericniebler.com/2020/11/08/structured-concurrency/
---

## [9][I tried to efficiently render text with SDL_ttf, let me know what you think](https://www.reddit.com/r/cpp/comments/jr8s11/i_tried_to_efficiently_render_text_with_sdl_ttf/)
- url: /r/sdl/comments/jr8qdl/i_tried_to_efficiently_render_text_with_sdl_ttf/
---

## [10][Meaning of C++ Stanard Revisions](https://www.reddit.com/r/cpp/comments/jr0xsk/meaning_of_c_stanard_revisions/)
- url: https://www.reddit.com/r/cpp/comments/jr0xsk/meaning_of_c_stanard_revisions/
---
Is there any meaning of the revisions of the C++ standard, likeN4661, N4668?
## [11][Conservative instantiation in variadic templates](https://www.reddit.com/r/cpp/comments/jqx1oa/conservative_instantiation_in_variadic_templates/)
- url: https://www.reddit.com/r/cpp/comments/jqx1oa/conservative_instantiation_in_variadic_templates/
---
A use case of variadic templates that often arises in my libraries and applications is the following:

    void f(std::convertible_to&lt;double&gt; auto const&amp;... x);

Or put into words: wanting to accept an arbitrary number of arguments that all have to be convertible to a certain type. The issue with the provided implementation is that it leads to an arbitrary number of template instantiations even in the case that they all have the same number of parameters:

    f(1.0, 2.0), f(1.0, 2); // etc.

More instantiations lead to bloated binaries and longer compile and link times, so it is self-explanatory that one wants to avoid instantiations wherever possible. I would propose using a syntax along the lines of this:

    void f(double... x);

This, too, should just be an abbreviated function template but the deduced type parameter pack `T` ought to satisfy the constraint `std::conjunction_v&lt;std::is_same&lt;T, double&gt;...&gt;`. In other words, only the number of types in the pack should vary between instantiations, not the types themselves.

This proposed syntax may clash however with plain old C variadics (in the case that the argument pack isn't named) because they - for some reason - aren't required to be separated by a comma. A possible solution is requiring such typed parameter packs to always be named.

Has anybody seen or written a proposal regarding this? If not, how can I write a proposal?
