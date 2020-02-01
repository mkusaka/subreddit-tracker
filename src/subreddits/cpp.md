# cpp
## [1][C++ Jobs - Q1 2020](https://www.reddit.com/r/cpp/comments/eiila4/c_jobs_q1_2020/)
- url: https://www.reddit.com/r/cpp/comments/eiila4/c_jobs_q1_2020/
---
Rules For Individuals
---------------------

* **Don't** create top-level comments - those are for employers.
* Feel free to reply to top-level comments with **on-topic** questions.
* I will create one top-level comment for **meta** discussion.
* I will create another top-level comment for **individuals looking for work** and **community groups looking for sponsors**.

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

\*\*Remote:\*\* [Do you offer the option of working remotely? If so, do you require employees to live in certain areas or time zones?]

&amp;nbsp;

\*\*Visa Sponsorship:\*\* [Does your company sponsor visas?]

&amp;nbsp;

\*\*Technologies:\*\* [Required: do you mainly use C++98/03, C++11, C++14, C++17, or the C++20 working draft? Optional: do you use Linux/Mac/Windows, are there languages you use in addition to C++, are there technologies like OpenGL or libraries like Boost that you need/want/like experience with, etc.]

&amp;nbsp;

\*\*Contact:\*\* [How do you want to be contacted? Email, reddit PM, telepathy, gravitational waves?]

---

Previous Post
--------------

* [C++ Jobs - Q4 2019](https://www.reddit.com/r/cpp/comments/dbqgbw/c_jobs_q4_2019/)
## [2][A syntax-based overview of C++20 Concepts](https://www.reddit.com/r/cpp/comments/ex4lm5/a_syntaxbased_overview_of_c20_concepts/)
- url: https://omnigoat.github.io/2020/01/19/cpp20-concepts/
---

## [3][A freestyle rap from the British comedian Chris Turner dedicated to the ISO C++ standards committee meeting in Prague](https://www.reddit.com/r/cpp/comments/ewsgg4/a_freestyle_rap_from_the_british_comedian_chris/)
- url: https://www.cameo.com/v/TXgs6dWbN
---

## [4][A Universal I/O Abstraction for C++](https://www.reddit.com/r/cpp/comments/ewr18m/a_universal_io_abstraction_for_c/)
- url: https://cor3ntin.github.io/posts/iouring/
---

## [5][Trying to understand C++20 modules](https://www.reddit.com/r/cpp/comments/ewpmzi/trying_to_understand_c20_modules/)
- url: https://www.reddit.com/r/cpp/comments/ewpmzi/trying_to_understand_c20_modules/
---
So I've been reading and watching videos, but don't feel that I've really understood how modules work. Or rather, I get the gist of it, and I'm all for having better modularisation, but I am troubled by how they are to be implemented. For example, I've seen examples where the module names do not match the files in which they reside. How would the compiler located an imported module in that case?

Many years ago, I wrote Delphi for a living, and recall being very impressed with Pascal units. A "unit" has the same name as its file, so the lookup is simple. The unit has distinct named "interface" and "implementation" sections. A unit can import the interface sections of other units with "uses". I may be misremembering, but that was more or less it. I can imagine how to implement this so that units are imported from a binary intermediate, and that the binary intermediate is created in a one-time on-demand compilation of the unit's interface section. There would seem to be no barrier to parallel compilation of the implementation sections of the units in a project. I'm probably over-simplifying. I do however recall that compiling Delphi projects was incredibly fast compared to C++ projects. 

I've read a number of concerns about how to find modules and how to parallelise compilation in C++20, but have no idea if those concerns are overstated, have been addressed already, are just plain wrong, or what. I'd be grateful for more information. 

Thanks.
## [6][DS323I RTC HAL Driver Written in C++](https://www.reddit.com/r/cpp/comments/ex4kzu/ds323i_rtc_hal_driver_written_in_c/)
- url: https://www.reddit.com/r/cpp/comments/ex4kzu/ds323i_rtc_hal_driver_written_in_c/
---
Hey, guys! 

Here is a HAL driver for the DS3231 RTC I wrote a few months back. I felt to share.
It can work on any platform via the I2C interface. Similar to Rust's embedded-hal I2C interface which I really wish we had.
 
I'm really excited about C++ 20's concepts and I feel it'd be highly useful in this type of scenario.

Contributions are welcome!

https://github.com/lamarrr/DS3231
## [7][How 'granular' should a CMake target be?](https://www.reddit.com/r/cpp/comments/ewyw9h/how_granular_should_a_cmake_target_be/)
- url: https://www.reddit.com/r/cpp/comments/ewyw9h/how_granular_should_a_cmake_target_be/
---
In John Lakos' [talk](https://youtu.be/K_fTl_hIEGY?t=1398) about physical designs of C++ projects, he showed an example of such an idea with four classes and three files. My question is how should this reflect in CMake? Should each logical component be a CMake library target for ease of testing or do they belong in a larger CMake library (maybe geometry in this case)? Doing the former seems to be clearer but it introduces lots of small targets that depend on each other and also requires lots of boilerplate CMake (specifying the files, compiler warnings, include directory for users, library dependencies, etc) which might slow down complication. Doing the latter might be difficult in cases when you just need to have some functionality but you're not sure if it actually belongs in such a library. For a personal executable project where nothing is exported, files are sometimes just a collection of code that are incapable of forming a cohesive library. 

&amp;#x200B;

Am I missing something completely or is the best practise a mixture of the two?
## [8][HPX Community Survey -- STE||AR Group](https://www.reddit.com/r/cpp/comments/ewo4tq/hpx_community_survey_stear_group/)
- url: http://stellar-group.org/2020/01/2020-hpx-community-survey
---

## [9][The Battle for C++ Knowledge Supremacy](https://www.reddit.com/r/cpp/comments/ewv9v6/the_battle_for_c_knowledge_supremacy/)
- url: https://cppbattle.net/battle/early-quelea-20-63f6
---

## [10][Relaunch your career in C++ at 39](https://www.reddit.com/r/cpp/comments/ewug2t/relaunch_your_career_in_c_at_39/)
- url: https://www.reddit.com/r/cpp/comments/ewug2t/relaunch_your_career_in_c_at_39/
---
Hello everyone,

this is not a motivational post, rather it's a call for motivational answers. Let me explain. I've been on the worst project of my entire life for almost 2 years now, and I need to stop this growing sensation of failure in life. My question is flatly simple: can I relaunch myself in a C++ career at 39?

I officially had experience in custom business software, but at least doing a lot of backend and system integration allowed me to cultivate strong C#/.Net skill. Now as I said, I don't know how, but I fell in this trap of fu\*\*ing project, which is basically read idiotic incidents on service-now and solve the problem, clean up corrupt data, and whatever. Coding is just one line a week at best (maintaining one of the worst C# code bases I ever seen in my life).

For anyone wondering why I am posting this on /r/cpp, it is because I am not new to the language, and looking back at my \~20y career, I now see that I felt satisfaction only when I was developing useful APIs in C, C with classes (more often), C++ (very few times), and when doing systems programming (if you can call it such).

I have a few notable examples of satisfying projects:

* It was the early 2000s, every enterprise reporting solution was fast in showing its limits (Crystal Reports 8 anyone?) so I had to build this "graphics engine" for text formatting and print previewing to meet customer's needs.
* Drive a call recording server, I developed the project with STL, and the job was to consume a proprietary Java-style C++ library and expose my API with COM (yes I know a lot of win32 and COM).
* 4 years ago, I developed in C# a CTI server software. It was structured to be PBX vendor independent, and the purpose was to manage call distribution to the users of another company product. Anyway it didn't end well, the PM had obsolete ideas about the project and removed me in favor of more obeying colleagues.

You see the problem is I have little to nothing of curricular experience with C++. So what chances do you think I may have to restart my career? Also I see that in the last years developing on Linux has become so easy that even a Windows developer could do it. I am based in Italy, which nation you think could be the best for me?

Thanks
## [11][Status of exception specification removal?](https://www.reddit.com/r/cpp/comments/ewwir3/status_of_exception_specification_removal/)
- url: https://www.reddit.com/r/cpp/comments/ewwir3/status_of_exception_specification_removal/
---
C++ exception specifications were deprecated in C++11. The latest update to Alisdair's [proposal](http://open-std.org/JTC1/SC22/WG21/docs/papers/2016/p0003r4.html) to remove them just says "Retargeted for C++Next-after-17." Is this onboard for C++20? Has it been punted to 23+?
