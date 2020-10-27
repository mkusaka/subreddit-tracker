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
## [2][Hi, I'm new to C++. I was just wondering why the Qt framework is more than 1GB but the compiled app is very small.](https://www.reddit.com/r/cpp/comments/jiszf3/hi_im_new_to_c_i_was_just_wondering_why_the_qt/)
- url: https://www.reddit.com/r/cpp/comments/jiszf3/hi_im_new_to_c_i_was_just_wondering_why_the_qt/
---

## [3][Semantic requirements in concepts - Andrzej's C++ blog](https://www.reddit.com/r/cpp/comments/jiq75r/semantic_requirements_in_concepts_andrzejs_c_blog/)
- url: https://akrzemi1.wordpress.com/2020/10/26/semantic-requirements-in-concepts/
---

## [4][{fmt} 7.1 released with faster floating point formatting, file output, improved chrono and ranges support and more](https://www.reddit.com/r/cpp/comments/jierin/fmt_71_released_with_faster_floating_point/)
- url: https://github.com/fmtlib/fmt/releases/tag/7.1.0
---

## [5][I made this ASCII Terrain generator with Parallax effect in C++, I'll share my code here soon.](https://www.reddit.com/r/cpp/comments/jibs2z/i_made_this_ascii_terrain_generator_with_parallax/)
- url: https://youtu.be/ZJsgidL1lWs
---

## [6][Why no std::variant or std::tuple .get&lt;T&gt;() member function?](https://www.reddit.com/r/cpp/comments/jifw3w/why_no_stdvariant_or_stdtuple_gett_member_function/)
- url: https://www.reddit.com/r/cpp/comments/jifw3w/why_no_stdvariant_or_stdtuple_gett_member_function/
---
It would just be syntactic sugar, `xx.get&lt;yy&gt;() -&gt; std::get&lt;yy&gt;(xx)`, but it seems inconsistent with most other containers in the standard for accessing a member (its also easier to write if you dont `using std::get`). Its easy to make a wrapper for this but it would be nice if it were in the standard instead.
## [7][C++20 Ranges, Projections, std::invoke and if constexpr](https://www.reddit.com/r/cpp/comments/jib2wl/c20_ranges_projections_stdinvoke_and_if_constexpr/)
- url: https://www.bfilipek.com/2020/10/understanding-invoke.html
---

## [8][Will the default parameter syntax ever be fixed?](https://www.reddit.com/r/cpp/comments/jixfdl/will_the_default_parameter_syntax_ever_be_fixed/)
- url: https://www.reddit.com/r/cpp/comments/jixfdl/will_the_default_parameter_syntax_ever_be_fixed/
---
int myFunc (int a = 3, int b = 5)

I cant specify b without specifying a. Why cant this syntax just be like this

myFunc(, 4)

Just put a empty comma where you dont want to specify much better and ur not gonna be forced to have all default params in the end and u can specify only specific params

Will it ever be improved?
## [9][What is the debate between OOP and Functional programming?](https://www.reddit.com/r/cpp/comments/jiua22/what_is_the_debate_between_oop_and_functional/)
- url: https://www.reddit.com/r/cpp/comments/jiua22/what_is_the_debate_between_oop_and_functional/
---
Why have people been saying for the past few years that OOP is becoming obsolete? They seem to be two completely different styles of programming for two completely unrelated tasks, so I'm not sure how one is replacing the other. Is it true that OOP is becoming obsolete? Please provide evidence, lots of people don't like OOP so they say "OOP is obsolete", at least from what I've seen so far on Reddit, similar to how they say that a youtuber, a political party, or a celebrity that they don't like is "losing followers and will die out soon" lol.
## [10][Roi Barkan will talk about "C++20 Concepts" next Monday, November 2nd](https://www.reddit.com/r/cpp/comments/jipxc6/roi_barkan_will_talk_about_c20_concepts_next/)
- url: https://www.reddit.com/r/cpp/comments/jipxc6/roi_barkan_will_talk_about_c20_concepts_next/
---
Hi Everyone,

CPPDUG, the Dublin C/C++ Meetup Group, will host an online talk on "C++20 Concepts" by Roi Barkan next Monday, November 2nd. If the topic is of interest to you, please RSVP here: [https://www.meetup.com/cppdug/events/274196059](https://www.meetup.com/cppdug/events/274196059)

Talk description: C++20 concepts give us a brand new way to specify how our generic libraries and algorithms should be used and allow library users to better understand what is expected of them. However, concepts didn't rise from chaos, and before their adoption C++ library authors had other mechanisms to meet that goal. In this talk I'll briefly discuss concepts, and focus on their pros and cons compared to more mature mechanisms. Special attention will be given to the 'semantic' vs. 'syntactic' nature of concepts, and I'll try to give motivating examples of semantic-only concepts. I'll then explore best practices, and show potential pitfalls in various approaches to the challenge of defining interactions between generic code authors and users.

Talk outline:  
\- Problem statement, examples of 'concepts in the wild' (iterator traits, trivially-copyable, ordering)  
\- Brief overview of C++20 concepts - syntax alternatives, common examples  
\- Pre-concepts typical approaches - mainly type-traits, iterator\_traits, void\_t, enable\_if  
\- Tricky cases: semantics vs. syntax; common/subtle 'bugs'; explicit call-site override  
\- 'Philosophical' discussion on core concepts - cheaply-copyable, trivially-copyable

Speaker bio: Professional software developer and architect since 2000, Roi's main focus throughout his career was on high performance and distributed systems, implementing complex and innovative algorithms. Roi has been the VP technologies of Istra Research since 2014, where he helps creating low latency financial systems. Prior to working for Istra Research, Roi spent 12 years in software development, architecture and management in the IT Security field. Roi received his B.A in Computer Science with high honors from the Technion in Israel, and his executive MBA from Tel Aviv University.
## [11][I made command line parser library a long time ago, now I decided to share it with the world.](https://www.reddit.com/r/cpp/comments/jiease/i_made_command_line_parser_library_a_long_time/)
- url: https://gitlab.com/mildlyparallel/dashh
---

