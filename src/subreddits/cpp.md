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
## [2][Bjarne Stroustrup: The Beauty and Power of "Primitive" C++](https://www.reddit.com/r/cpp/comments/j1x14e/bjarne_stroustrup_the_beauty_and_power_of/)
- url: https://www.youtube.com/watch?v=ERzENfQ51Ck
---

## [3][c++20 coroutines, opinions?](https://www.reddit.com/r/cpp/comments/j1xhgv/c20_coroutines_opinions/)
- url: https://www.reddit.com/r/cpp/comments/j1xhgv/c20_coroutines_opinions/
---
Fundamentally this post is asking peoples opinion about the new C++ coroutines. Especially when it comes to application on embedded systems. 
And, when compared to the alternatives, such as Rusts async / await.

---
Please respond with your opinion and experiences before reading on.

---
The potential of these things conceptually is enormous. When I heard about the proposal I was extremely excited. My first tests also indicated really good things for embedded systems. The size of various coroutine handles was light enough to push into a queue from within a hardware interrupt!!! (where resume could be called). This could truly be game changing for embedded systems.

My experience started with real optimism! Imagine being able to handle DMA memory loading just like this:
```CPP
auto page = co_await memory.LoadPage(address);  
```  
Or to be able to wait on a button press like this:
```CPP
co_await button.ShortPress();
```
Or with mqtt connections:
```CPP
co_await mqtt.Subscribe(...);
```

This could lead to really easy to analyze code and virtually branch-less logic(sort of) in many cases. 

But when using them I came across dark sides. Allocation of the frames with sizes of hundreds of bytes meant that it was no longer feasible to use them for interrupt handling. This is on an embedded system where allocations are often simply intolerable. 

Interrupts occur when anything completes. There are a lot of these things though. If there were one co-routine frame for each (worst case), the cumulative allocated frame size is totally and utterly unacceptable.

I contacted Gor, who did the co-routine proposal. He kindly got back to me with a detailed example that showed HALO occurring, it occurred on all nested calls coroutines but there was always at least 1 allocation for first coroutine frame.

It seems like there is a disconnect between the communities excitement for this feature and my experiences. It feels unending, each day I try to figure this out, but each day I come up short. I'm hoping enough people agree with me that they are intolerable and we can get some improvements in the standard. Or someone can tell me why I'm wrong.
## [4][Draft proposal for compilation using C++ as build language](https://www.reddit.com/r/cpp/comments/j1ebld/draft_proposal_for_compilation_using_c_as_build/)
- url: https://www.reddit.com/r/cpp/comments/j1ebld/draft_proposal_for_compilation_using_c_as_build/
---
Hello,  


There are multiple tools to compile a C++ project: CMake, premake, autoconf, build2, etc.  
Each  of these tools require the learning of a new syntax, and interacting  between projects compiled with differents tools might be challenging.

Package managers face also this problem, because they must support all the tools that exist to support all libraries.  
Moreover, integrating package managers in the build tools is often difficult, or at least hacky.

My suggestion is to use C++ as the base language for building C++ projects.  
Motivations are:

\- A C++ compiler is all you need to compile projects.

\- Build tools would become C++ libraries.

\- C++ developers already now its syntax.  
I have made an informal draft which describes a basic API which would make this possible.  
Here is a link: [https://gist.github.com/J-Vernay/bd8ec49374987c628d02601ef85cd9a7](https://gist.github.com/J-Vernay/bd8ec49374987c628d02601ef85cd9a7)  


Let me know what you think :)
## [5][Adding Ask me Anything to Meeting C++ 2020](https://www.reddit.com/r/cpp/comments/j1weqh/adding_ask_me_anything_to_meeting_c_2020/)
- url: https://meetingcpp.com/meetingcpp/news/items/Adding-Ask-me-Anything-to-Meeting-Cpp-2020.html
---

## [6][cmake rants, anyone?](https://www.reddit.com/r/cpp/comments/j1yxwf/cmake_rants_anyone/)
- url: https://www.reddit.com/r/cpp/comments/j1yxwf/cmake_rants_anyone/
---
CMake is definitely the best tool available for building on different platforms. However, I think it has a lot of shortcomings and we need something better. I am trying to gather a list of problems with cmake so that they can be kept in mind when creating an alternative. I have worked on cmake for 2 years only, so I'm looking to hear from more experienced folks on the sort of problems they face.

I'll start by enlisting my complaints :

  - Inconsistencies/ too much overloading in the syntax. Functions can have a mind boggling number of arguments which makes it dificult to understand.

  - Does not enforce much structure, which means cmake files of different projects can be written in very different ways, which makes difficult to understand quickly.

  - Hard to debug

Anything else?
## [7][Do you have any side projects made with C++? Do any of them earn you income?](https://www.reddit.com/r/cpp/comments/j1q36h/do_you_have_any_side_projects_made_with_c_do_any/)
- url: https://www.reddit.com/r/cpp/comments/j1q36h/do_you_have_any_side_projects_made_with_c_do_any/
---
Do any of you C++ devs have any side projects that bring you any income? Yes, yes, a job is the main way most make income. I thoght it might be a little more interesting to ask this question here. Please feel free to share. Thanks.
## [8][Guide: "A modern, open source C++ dev environment with Visual Studio Code, vcpkg, and CMake"](https://www.reddit.com/r/cpp/comments/j1dh9w/guide_a_modern_open_source_c_dev_environment_with/)
- url: https://www.reddit.com/r/cpp/comments/j1dh9w/guide_a_modern_open_source_c_dev_environment_with/
---
Hi there,

Almost exactly one week ago [this post](https://www.reddit.com/r/cpp/comments/ix9n1u/why_is_it_such_an_abysmal_pain_to_use_libraries/) has been posted on /r/cpp by someone who was asking why it is such a pain to use libraries in C++ compared to other languages. From there started a number of enthusiastic discussions in the comments regarding existing solutions to deal with libraries, people sharing their personal opinions and feeling, their experience, etc.

Reading the previous post and discussion motivated me to spend some time with both `vcpkg` and `CMake`, try to integrate as much as I could in Visual Studio Code (because that's my editor and general dev environment), and write down my notes during that process (for my future self and potentially others).

The result is a guide "A modern, open source C++ dev environment with Visual Studio Code, vcpkg, and CMake" that I published just a few minutes ago on my personal blog: https://sam.elborai.me/blog/vscode-cpp-dev-environment-2020.

I am neither a CMake, nor vcpkg, nor C++ expert but I think that can help some people navigate in the vast sea of C++ tooling. And because folks here are actually experts in these domains, I would be interested to get your feedback ✨.

In any case I had some fun and learned new things :)
## [9][The fickle aggregate: how the definition of what is an aggregate class has changed through various standard versions (C++11 through C++20), and how these rules can have somewhat surprising consequences.](https://www.reddit.com/r/cpp/comments/j1bhgg/the_fickle_aggregate_how_the_definition_of_what/)
- url: https://dfrib.github.io/the-fickle-aggregate/
---

## [10][Snippet for Pattern Support in Switch Statements](https://www.reddit.com/r/cpp/comments/j1pa4i/snippet_for_pattern_support_in_switch_statements/)
- url: https://www.reddit.com/r/cpp/comments/j1pa4i/snippet_for_pattern_support_in_switch_statements/
---
This is a pretty trivial 5 minute example of a switch case with regex support.  
[https://godbolt.org/z/6exWWT](https://godbolt.org/z/6exWWT)

Hope people can use this to neaten up code a little. Its only fractionally neater than raw regexes and isn't as neat as rust enums (what I'm looking to reproduce).
## [11][What is the standard Library?](https://www.reddit.com/r/cpp/comments/j1e8i5/what_is_the_standard_library/)
- url: https://cor3ntin.github.io/posts/std/
---

