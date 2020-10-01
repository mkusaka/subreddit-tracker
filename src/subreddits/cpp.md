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
## [2][Do you enforce Core Guidelines and/or use the Guidelines Support Library (GSL)? If no, why?](https://www.reddit.com/r/cpp/comments/j3752j/do_you_enforce_core_guidelines_andor_use_the/)
- url: https://www.reddit.com/r/cpp/comments/j3752j/do_you_enforce_core_guidelines_andor_use_the/
---
I am watching videos from the latest CppCon event, speakers mention from time to time the C++ Core Guidelines and the support library (GSL) as a way to help modernize codebases and avoid known flaws. I'm trying to use those guidelines for my personal projects (enforced by Microsoft's analyzer, I don't know other analyzers), mostly because I'm still quite new to the language and see it as a good way to learn over time. But I'm wondering if it is actually used by the rest of the community, either in open source projects or at companies. I haven't seen the GSL used anywhere outside of my own code, and I don't often see it coming up in online discussions outside of public talks (though as I'm still new I may just not know about projects that are using them).

  
For the sake of discussion I would be interested to know who is using these (either the guidelines, or the support library, or both?), what are your opinions and experience with them (do you have rules you always disable?), and if you don't use them what are the reasons?  


For people who don't already know about them:  
\- C++ Core Guidelines:  [https://isocpp.github.io/CppCoreGuidelines/CppCoreGuidelines](https://isocpp.github.io/CppCoreGuidelines/CppCoreGuidelines)   
\- Guidelines Support Library by Microsoft:  [https://github.com/microsoft/GSL](https://github.com/microsoft/GSL)   
\- Lite Guidelines Support Library:  [https://github.com/gsl-lite/gsl-lite](https://github.com/gsl-lite/gsl-lite)   
\- How to use the Core Guidelines checker in Visual Studio:  [https://docs.microsoft.com/en-us/cpp/code-quality/using-the-cpp-core-guidelines-checkers?view=vs-2019](https://docs.microsoft.com/en-us/cpp/code-quality/using-the-cpp-core-guidelines-checkers?view=vs-2019)
## [3][Meeting C++ 2020 tickets for Berlin are available](https://www.reddit.com/r/cpp/comments/j36xv2/meeting_c_2020_tickets_for_berlin_are_available/)
- url: https://meetingcpp.com/meetingcpp/news/items/Meeting-Cpp-2020-tickets-for-Berlin-are-available.html
---

## [4][[Hacktoberfest] Looking for reviewers / contributors for a modern C++ game engine](https://www.reddit.com/r/cpp/comments/j37km1/hacktoberfest_looking_for_reviewers_contributors/)
- url: https://www.reddit.com/r/cpp/comments/j37km1/hacktoberfest_looking_for_reviewers_contributors/
---
Hello r/cpp !

As you might already know, Hacktoberfest started today and I'm here to talk about a personal project of mine.

## Hacktoberfest

If you don't know what Hacktoberfest is about, I'll make it short : complete 4 pull-requests during the month of october and get a free T-shirt or a tree planted.

## The project

[ÖbEngine](https://github.com/Sygmei/ObEngine) is a 2d game engine made with C++, Lua and eggplants. It strive to be user-friendly, versatile and fun to use.

I started this project with little-to-no knowledge of C++ and committed many deadly sins. After a long time of continuous improvement (and there is still a lot of room to improve) I can say that I am proud of most of my codebase.
I'd love to have some honest reviews from people of r/cpp as I know that there must be flaws that I'm not yet aware of in my project codebase.

I tried to embrace modern C++ and I use features from C++11, 14 and 17 :

|**C++ 17 Feature**|**Example**|
|:-|:-|
|Nested namespaces|[*src/Core/Bindings/Animation/Easing/Easing.cpp*](https://github.com/Sygmei/ObEngine/blob/master/src/Core/Bindings/obe/Animation/Easing/Easing.cpp#L7)|
|Structured bindings|[*tests/src/Transform/UnitVectorTests.cpp*](https://github.com/Sygmei/ObEngine/blob/master/tests/src/Transform/UnitVectorTests.cpp)|
|Selection statements with initializers|[*src/Core/Animation/Animation.cpp*](https://github.com/Sygmei/ObEngine/blob/c3608ec2b848da9ddbfa5f7d44e50fa7c9bc241e/src/Core/Animation/Animation.cpp#L57)|
|`[[fallthrough]]` statements|[*src/Core/Engine/Engine.cpp*](https://github.com/Sygmei/ObEngine/blob/8e6222c05c925b5793ebf82e8d9506865c9edf12/src/Core/Engine/Engine.cpp#L170)|
|`[[nodiscard]]` statements|[*include/Core/Transform/Referential.hpp*](https://github.com/Sygmei/ObEngine/blob/master/include/Core/Transform/Referential.hpp#L42)|
|`std::variant`|[*include/Core/Input/InputButton.hpp*](https://github.com/Sygmei/ObEngine/blob/8e6222c05c925b5793ebf82e8d9506865c9edf12/include/Core/Input/InputButton.hpp#L44)|
|`std::optional`|[*include/Core/Transform/Polygon.hpp*](https://github.com/Sygmei/ObEngine/blob/8e6222c05c925b5793ebf82e8d9506865c9edf12/include/Core/Transform/Polygon.hpp#L139)|
|`std::any`|[*include/Core/Graphics/DrawUtils.hpp*](https://github.com/Sygmei/ObEngine/blob/8e6222c05c925b5793ebf82e8d9506865c9edf12/include/Core/Graphics/DrawUtils.hpp#L40)|
|`std::string_view`|[*include/Core/Exception.hpp*](https://github.com/Sygmei/ObEngine/blob/8e6222c05c925b5793ebf82e8d9506865c9edf12/include/Core/Exception.hpp#L16)|
|`std::filesystem`|[*src/Core/Utils/FileUtils.cpp*](https://github.com/Sygmei/ObEngine/blob/master/src/Core/Utils/FileUtils.cpp)|

## How can I help ?

Thank you for reading this post first of all ! Here are some things you could help with :

- Review the codebase and give advices on how to make the code better, that would help a lot

- Test the engine, create game prototypes with it and give me your feedback, what could I improve, what features do you want

- Contribute directly by coding features, there is a list of Hacktoberfest labelled issues right here : https://github.com/Sygmei/ObEngine/issues?q=is%3Aissue+is%3Aopen+label%3Ahacktoberfest

- If you have any suggestions, I'd be happy to hear it !

Check out for the links in the comments !
## [5][Graph-based execution runtime?](https://www.reddit.com/r/cpp/comments/j2y95s/graphbased_execution_runtime/)
- url: https://www.reddit.com/r/cpp/comments/j2y95s/graphbased_execution_runtime/
---
Been looking for a library that can run a graph of a data processing flow. In particular, a set of input nodes (could be any data, i.e. templated class), processing nodes (arbitrary code, with or without state), and output nodes. Execution controlled by a runtime that sequences all processing steps as needed by the graph, running things in parallel (or on GPUs etc) where possible. Other key features:

* visualization tools, render the graph
* execution tracing &amp; critical path analysis
* low overhead on thread locks, mutex's etc.
* No unnecessary data copies between processing nodes.
* Some notion of multi-iteration, i.e. run the graph over and over on new 
input, and stateful processing nodes update as they run.
* Some kind of constant/global config data that all nodes can access but not modify.

This kind of thing seems generally useful in HPC, robotics, etc. Surprising that I haven't found some widely used, actively developed framework.
## [6][It Is Never a Compiler Bug Until It Is](https://www.reddit.com/r/cpp/comments/j2jf0s/it_is_never_a_compiler_bug_until_it_is/)
- url: http://r6.ca/blog/20200929T023701Z.html
---

## [7][[Hacktoberfest] We need C++ and Qt developers for the SimulIDE Free and Open Source software](https://www.reddit.com/r/cpp/comments/j2uwoa/hacktoberfest_we_need_c_and_qt_developers_for_the/)
- url: https://www.reddit.com/r/cpp/comments/j2uwoa/hacktoberfest_we_need_c_and_qt_developers_for_the/
---
Hey folks,

The holy month of October has already begun in many parts of the globe, and many programmers and geeks are rushing to get their hands on some nice FLOSS. Whether for publicity and reputation, programing and teamwork experience, or for the sheer joy of contributing back to the community, many participate in the Hachtoberfest pilgrimage. Here I have a nice piece of software you can dive in, or who knows you might become the core developer soon. ;)

[SimulIDE](https://github.com/SimulIDE) is a hardware emulation software. You may consider it as a FLOSS alternative to [Proteus Virtual System Modelling (VSM)](https://www.labcenter.com/simulation/). [Here](https://www.reddit.com/r/electronics/comments/apysi7/integrated_simulation_of_an_arduino_plus_analogue/?utm_source=share&amp;utm_medium=web2x&amp;context=3) you may see a demo of the software emulating an Arduino circuit. The software was created by [Santiago González](santigoro@gmail.com) and is still being actively developed by him. However, I had to fork the software simply because:

- Sadly the original developer doesn't use any version control. Right now zip files and the binaries are published on Patreon, only to the premium users. 
- There is no easy way for the community to contribute to the project, open new issues, or even ask questions. 

If you like to help:

- Please join our D i s c o r d channel (in the comments 👇)
- Take a look at [the source code](https://github.com/SimulIDE/SimulIDE), try to compile on any platform you have (Linux, macOS, Windows...), open new issues 
- Share your thoughts and ideas on how we can initiate the fork and make it more contributor-friendly. The final goal is to make it a self-sustaining software developed by a community. 
- Help me make some `good first issues` like the ones folks are doing for the [MyLibreLab software](https://github.com/MyLibreLab/MyLibreLab/issues) already.
- Share this withing the C++ and Qt community you have

Looking forward to having many of you over there
## [8][C++ Grad Scheme Opportunity](https://www.reddit.com/r/cpp/comments/j36wup/c_grad_scheme_opportunity/)
- url: /r/CompSciStudents/comments/j36prw/graduate_scheme_opportunity/
---

## [9][Trying to build an Open Source browser in 2020 - Patricia Aas](https://www.reddit.com/r/cpp/comments/j2ifjr/trying_to_build_an_open_source_browser_in_2020/)
- url: https://www.youtube.com/watch?v=a0aY3rOD0Ss
---

## [10][Plenary: Performance Matters - Emery Berger - CppCon 2020](https://www.reddit.com/r/cpp/comments/j2nxi1/plenary_performance_matters_emery_berger_cppcon/)
- url: https://www.youtube.com/watch?v=koTf7u0v41o
---

## [11][Simple graphical API for kids?](https://www.reddit.com/r/cpp/comments/j2sqg7/simple_graphical_api_for_kids/)
- url: https://www.reddit.com/r/cpp/comments/j2sqg7/simple_graphical_api_for_kids/
---
My son (7th grade) is learning modern C/C++. I'd like to mix it up a bit with some elementary physics simulations since he's also interested in physics. But it's quite tedious with just cin/cout. Is there a C++ library with a simple API I could set up for him to create some visualization, like moving point masses and such? It would allow him to understand physics better, as well as having more interesting programming tasks.

He already did some simulations in Scratch and some in Python, but he really wants to improve his C++ skills, so what can you do.
