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
## [2][CLion 2020.2: Makefile Projects, C++20, Enhanced Code analysis, Doctest, and Other Unit Testing Support Improvements](https://www.reddit.com/r/cpp/comments/i0czrr/clion_20202_makefile_projects_c20_enhanced_code/)
- url: https://blog.jetbrains.com/clion/2020/07/clion-2020-2-makefile-cpp20-doctest/
---

## [3][constexpr if and requires expressions changed everything about template metaprogramming](https://www.reddit.com/r/cpp/comments/i0i0e4/constexpr_if_and_requires_expressions_changed/)
- url: https://www.reddit.com/r/cpp/comments/i0i0e4/constexpr_if_and_requires_expressions_changed/
---
then: [https://gist.github.com/ChemiaAion/4118a3598f0b120b7f9c5884e9799a8b](https://gist.github.com/ChemiaAion/4118a3598f0b120b7f9c5884e9799a8b)

now: [https://godbolt.org/z/KKrPMr](https://godbolt.org/z/KKrPMr)

we can finally template-metaprogram in a human readable language.
## [4][libstud-json: JSON pull-parser/push-serializer library for C++](https://www.reddit.com/r/cpp/comments/i0l6ao/libstudjson_json_pullparserpushserializer_library/)
- url: https://github.com/libstud/libstud-json/
---

## [5][Getters and Setters, and Member functions vs the alternative](https://www.reddit.com/r/cpp/comments/i06xfh/getters_and_setters_and_member_functions_vs_the/)
- url: https://www.reddit.com/r/cpp/comments/i06xfh/getters_and_setters_and_member_functions_vs_the/
---
I took a class last semester about OOP with C++, but to be honest, I'm still struggling with understanding what good practices are when it comes to making classes and their functions. When I code C++, I usually just make all the variables private, and make member functions to get/set the data, as well as perform other tasks. However, I've heard that using getters and setters is bad practice, but so is public variables, so what should I do? Also, I've seen people not use member functions and instead create global functions where you pass in a constant reference to the class. Is this good practice? What about if you're working with private variables? I'm not too familiar with how to properly structure and organize classes in different ways (I did my whole final project just with private data and getters/setters) so any help would be appreciated.
## [6][Why does one of these code samples get optimized but not the other one?](https://www.reddit.com/r/cpp/comments/i01cwd/why_does_one_of_these_code_samples_get_optimized/)
- url: https://www.reddit.com/r/cpp/comments/i01cwd/why_does_one_of_these_code_samples_get_optimized/
---
I was playing around with inlining in godbolt and I came across a slightly puzzling difference, that I've summarised in these two samples:

* [inlined](https://godbolt.org/z/rc6rao)
* [not inlined](https://godbolt.org/z/5v54Gn)

As you can see, the only difference is that in the first example the variable whose reference I'm returning in the private function lives at global scope for this TU, while in the other example it's a static method variable.

It seems to me like there's no reason why the compiler shouldn't be able to infer that the memory location is going to be irrelevant in both cases (since the only way to access the variable is through the accessor, which discards the address). So why does it get inlined in one case and not the other? Is it a limitation of the compiler, or is there a language design reason for this?

Thanks!
## [7][Updated version of C++ Coding Standards By Herb Sutter and Andrei Alexandrescu](https://www.reddit.com/r/cpp/comments/i01lt3/updated_version_of_c_coding_standards_by_herb/)
- url: https://www.reddit.com/r/cpp/comments/i01lt3/updated_version_of_c_coding_standards_by_herb/
---
Hey guys, just wondering if there is an updated version of this book (or a similar book which talks about good C++ practices). Also, for people who have read this can you please confirm if they methodologies mentioned in the book still apply to modern C++.   
Linked the book below. 

LINK : [https://www.amazon.com/dp/0321113586](https://www.amazon.com/dp/0321113586)
## [8][Should you use the inline keyword or not?](https://www.reddit.com/r/cpp/comments/i0haek/should_you_use_the_inline_keyword_or_not/)
- url: https://www.reddit.com/r/cpp/comments/i0haek/should_you_use_the_inline_keyword_or_not/
---
I usually add the `inline` keyword when I think a function ought to be inlined. However, I've received some review comments and criticism about this.

Some people argue that `inline` is superfluous because these days compilers do a good job deciding when to inline things by default, and maybe even make a better decision as to what is worth to inline and what isn't.

A few years ago I fiddled around with the compiler explorer and found that not every compiler does inlining without the `inline` keyword, even when optimizations are enabled. However, I can't actually reproduce this anymore with more recent versions.

What do you guys think? Should `inline` be avoided because today's compilers are smart enough? Or should we keep using it "just in case"?
## [9][C++ on Sea 2020 video - "Code samples that actually compile" (lightning talk) - Clare Macrae](https://www.reddit.com/r/cpp/comments/hzktac/c_on_sea_2020_video_code_samples_that_actually/)
- url: https://www.youtube.com/watch?v=KtxTlBeVYiE
---

## [10][eCAL](https://www.reddit.com/r/cpp/comments/hzc1lo/ecal/)
- url: https://www.reddit.com/r/cpp/comments/hzc1lo/ecal/
---
[eCAL](https://github.com/continental/ecal) is a  modern C++ interprocess communication framework designed for low latency and massive data throughput. It was developed by a Continental R&amp;D team to manage distributed high performance computing for autonomous driving prototyping cars.

The API design is inspired by [ROS](https://www.ros.org/) but it has built in shared memory transport from the start and supports modern message protocols like [Google Protobuf](https://developers.google.com/protocol-buffers), [Flatbuffers](https://google.github.io/flatbuffers/) or [Cap'n Proto](https://capnproto.org/) (infinitely fast ;-)).

Finally there are some powerful applications for monitoring, recording and replaying the whole interprocess communication in a distributed manner.
## [11][C++ on Sea 2020 video - "Decision Fatigue and coding guidelines" (lightning talk) - Sandor Dargo](https://www.reddit.com/r/cpp/comments/hzhebp/c_on_sea_2020_video_decision_fatigue_and_coding/)
- url: https://www.youtube.com/watch?v=B-VnIakCOsw
---

