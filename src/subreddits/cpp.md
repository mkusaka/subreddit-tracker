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
## [2][What's in a function pointer? Machine code, actually!](https://www.reddit.com/r/cpp/comments/htekog/whats_in_a_function_pointer_machine_code_actually/)
- url: https://youtu.be/I_g7CSoNNxc?t=758
---

## [3][Not going to impress you, but I wanted to share: Structuring your code is so rewarding :)](https://www.reddit.com/r/cpp/comments/ht3htx/not_going_to_impress_you_but_i_wanted_to_share/)
- url: https://www.reddit.com/r/cpp/comments/ht3htx/not_going_to_impress_you_but_i_wanted_to_share/
---
So I have this project - biggest one I have worked on so far, but probably small in comparison to any software I have ever heard of.

I am working on this alone, but I will hand it over with full and proper documentation + handover training at the end, with the aim that other people can support and debug it as well.

*Not counting empty lines and pure comment lines, the project now has 33 header files with 1274 loc, and 24 code modules with 4353 loc.* 

*I have a ratio of 0.46 comments per line of code and a Makefile of 168 lines*

I spent the first two months working on the framework:

* Setting up a nice build folder, with subfolders for each module, according obj folder and subfolders, and a manually generated well-structured Makefile without redundant information.
* learning &amp; "proof-of-concepting" the (for me new) technologies I was going to use: snmp agents, boost program\_options, dynamic loading of a plugin (dlopen), atomic variables, function pointers
* implementing the application interaction: basic telemetry from the diagnostics module, inter-application communication via posix message queues, a general logging framework and configuration file logic
* ensuring full compatibility with both gcc on linux and gcc on cygwin64/Windows 10
* building proper test scenarios and test scripts

It was sometimes frustratingly slow, and especially debugging net-snmp mistakes I had in my code was a challenge (documentation is none too great, not even all of the tutorial examples are complete).

**But now**, now I am actually implementing the application functionality and the experience is like none before: I have so many utility functions at my disposal, thanks to the clear structure of the project I always know where to find stuff, and I have parameter checks in my functions left and right so that I always get meaningful error messages if I use a function wrongly.

I can work on a module and change a hundred lines of code in three different interacting functions - I build my project, and once I eliminate all the obvious bugs that the compiler complains about, the executables just work as they are intended to! I hardly ever manage to produce a crash bug - my thread handling is 99% clean and I don't cause race conditions or deadlocks. And when I decide to implement something more complex, it ends up falling apart into many easy steps that can be done with the tools of my framework.

&amp;#x200B;

This post is not intended to pat myself on the back, but rather to re-emphasize

**TL;DR Damn, does a clean and structured project feel rewarding! :)**
## [4][A lib for adding a stacktrace to every c++ exception in windows platform](https://www.reddit.com/r/cpp/comments/ht93d5/a_lib_for_adding_a_stacktrace_to_every_c/)
- url: https://www.reddit.com/r/cpp/comments/ht93d5/a_lib_for_adding_a_stacktrace_to_every_c/
---
 exceptions-stacktrace is a c++17 library which its purpose is to add a stacktrace to exception, even if it was thrown from some external lib you use like stl/boost/any other third party lib, with no code changes to your original code.  

https://twitter.com/ykfre1/status/1284308246472339457?s=20
## [5][C++ Template Library for Probabilistic Programming](https://www.reddit.com/r/cpp/comments/htbv8e/c_template_library_for_probabilistic_programming/)
- url: https://www.reddit.com/r/cpp/comments/htbv8e/c_template_library_for_probabilistic_programming/
---
Hi everyone,

I just wanted to share this library [autoppl](https://github.com/JamesYang007/autoppl) that a couple of my friends and I started for a class final project. We found that there was quite a lack of low-level tools for probabilistic programming and wanted to try making something for C++. I have been recently working on it more and have found it to be pretty successful for some examples. Any comments or feedbacks would be appreciated!
## [6][The Forgotten Art of Structured Programming - Kevlin Henney [C++ on Sea 2019]](https://www.reddit.com/r/cpp/comments/hszgh9/the_forgotten_art_of_structured_programming/)
- url: https://www.youtube.com/watch?v=SFv8Wm2HdNM
---

## [7][A day in the life of a C++ compiler developer](https://www.reddit.com/r/cpp/comments/hsr382/a_day_in_the_life_of_a_c_compiler_developer/)
- url: https://twitter.com/TartanLlama/status/1283816537778982914
---

## [8][hdoc: a modern documentation tool for C++](https://www.reddit.com/r/cpp/comments/ht0fmn/hdoc_a_modern_documentation_tool_for_c/)
- url: https://www.reddit.com/r/cpp/comments/ht0fmn/hdoc_a_modern_documentation_tool_for_c/
---
Hello, I wanted to announce the alpha release of a project I've been working on.

[hdoc](https://hdoc.io) is a modern documentation tool for C++.
What this means in practice is hdoc generates API documentation like Doxygen, but also allows you to bundle separate Markdown pages that will be rendered to HTML and put alongside your API docs.
In effect, hdoc will create a single source of documentation for your project.

You can see a demo of the documentation hdoc generates from LLVM at [demo.hdoc.io](https://demo.hdoc.io).

Key features:

- Autogenerated API docs, using a subset of Doxygen commands
- Integrated Markdown pages
- Instant code search throughout your codebase, similar to Algolia

The tool is currently in alpha stage, and we're looking for feedback from users to shape future development.
If you're interested in trying the alpha and providing feedback, reach out to contact@hdoc.io or check out [hdoc.io](https://hdoc.io).
## [9][lazycsv : A blazing fast single-header library for reading and parsing csv files in c++](https://www.reddit.com/r/cpp/comments/hss3ws/lazycsv_a_blazing_fast_singleheader_library_for/)
- url: https://www.reddit.com/r/cpp/comments/hss3ws/lazycsv_a_blazing_fast_singleheader_library_for/
---
[https://github.com/ashtum/lazycsv](https://github.com/ashtum/lazycsv)

This is my first public cpp library, it is \~300 line of code (i think it is easy to read).

I would like to have your feedback on it 😋.
## [10][CppCast: LLVM Hacking and CPU Instruction Sets](https://www.reddit.com/r/cpp/comments/hsobf3/cppcast_llvm_hacking_and_cpu_instruction_sets/)
- url: https://cppcast.com/bruno-cardoso/
---

## [11][Compiler Explorer now supports linking with some of the libraries](https://www.reddit.com/r/cpp/comments/hsctwj/compiler_explorer_now_supports_linking_with_some/)
- url: https://github.com/compiler-explorer/compiler-explorer/issues/2079
---

