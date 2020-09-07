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
## [2][C++20 approved, C++23 meetings and schedule update](https://www.reddit.com/r/cpp/comments/inw0ez/c20_approved_c23_meetings_and_schedule_update/)
- url: https://herbsutter.com/2020/09/06/c20-approved-c23-meetings-and-schedule-update/
---

## [3][Replacing unique_ptr with C++17's std::variant a Practical Experiment](https://www.reddit.com/r/cpp/comments/io36xt/replacing_unique_ptr_with_c17s_stdvariant_a/)
- url: https://www.bfilipek.com/2020/09/replacing-into-variant.html?m=1
---

## [4][Qt 6.0 Feature Freeze Milestone Reached](https://www.reddit.com/r/cpp/comments/int1rp/qt_60_feature_freeze_milestone_reached/)
- url: https://www.qt.io/blog/qt-6.0-feature-freeze-milestone-reached
---

## [5][Named Parameters in C++20 · Peter Dimov](https://www.reddit.com/r/cpp/comments/iny133/named_parameters_in_c20_peter_dimov/)
- url: https://pdimov.github.io/blog/2020/09/07/named-parameters-in-c20/
---

## [6][CppCon 2018: Walter E. Brown “C++ Function Templates: How Do They Really Work?”](https://www.reddit.com/r/cpp/comments/io6mp1/cppcon_2018_walter_e_brown_c_function_templates/)
- url: https://www.youtube.com/watch?v=NIDEjY5ywqU
---

## [7][Why You Should Use the Boost Software License · Peter Dimov](https://www.reddit.com/r/cpp/comments/inot7e/why_you_should_use_the_boost_software_license/)
- url: https://pdimov.github.io/blog/2020/09/06/why-use-the-boost-license/
---

## [8][Best way to start learning C++(20)](https://www.reddit.com/r/cpp/comments/io57yz/best_way_to_start_learning_c20/)
- url: https://www.reddit.com/r/cpp/comments/io57yz/best_way_to_start_learning_c20/
---
Hey,

  
I want to learn C++, because I want to work for Tesla when their new Gigafactory and developer place in Germany is ready and I read that they need C/ C++ programmers mostly.  
I am studying at the moment and already have some experience with Python in Machine Learning, some Java from university and I worked with Javascript/ React in the Web Development area for a company. 

My plan is to learn the foundation of C++ in some good tutorials and then improve my knowledge by helping in Open Source projects.

What sites/ tutorials would you recommend a beginner to learn C++? I am worried that many tutorials will be somewhat outdated now that C++20 is approved and I read that it will change how people write C++.  
Would you recommend to visit certain university courses?  


Thank you in advance :)
## [9][Visual Studio Code + Cpp](https://www.reddit.com/r/cpp/comments/inqkz6/visual_studio_code_cpp/)
- url: https://www.reddit.com/r/cpp/comments/inqkz6/visual_studio_code_cpp/
---
Hi all, 

i love working with VS Code and i'm also a passionate C++ developer. So i started a small project to gather nice and valuable extensions and configuration examples how to enhance developing C++ with VS Code. 

It can be found on github: [https://github.com/fbaeuerlein/cpp-vscode-guide](https://github.com/fbaeuerlein/cpp-vscode-guide)

Feel free to contribute and help to convince people that VSC is a really good IDE also for C++.
## [10][C++20 has been approved](https://www.reddit.com/r/cpp/comments/imzkl4/c20_has_been_approved/)
- url: https://twitter.com/patriceroy1/status/1302055575140945921?s=21
---

## [11][inline namespace bug in msvc](https://www.reddit.com/r/cpp/comments/inmo3a/inline_namespace_bug_in_msvc/)
- url: https://www.reddit.com/r/cpp/comments/inmo3a/inline_namespace_bug_in_msvc/
---
I went to report this using the IDE and it asked me to "log in" which I tried but Microsoft's browser widget broke my password manager so I said to login with GitHub and when I pressed that button another browser window popped up asking me to log in to GitHub but at the top it said "GitHub does not support this browser (lol)" so I gave up and I am posting this code snippet here:

This fails to compile in latest msvc:

    namespace boost {
    namespace json {
    
    inline namespace standalone {
    namespace detail {
    struct V{};
    } // detail
    } // standalone
    
    namespace detail {
    struct T{};
    } // detail
    struct U : detail::T{};
    
    } // json
    } // boost

See:   [https://godbolt.org/z/qbPesG](https://godbolt.org/z/qbPesG) 

It compiles everywhere else
