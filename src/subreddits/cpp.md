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
## [2][Avast C++ Meetup with Bjarne Stroustrup, Tony van Eerd, and Herb Sutter (Stream on 11th February 2020 6pm CET, 5pm UTC, 9am PCT, 12pm EST)](https://www.reddit.com/r/cpp/comments/f2409z/avast_c_meetup_with_bjarne_stroustrup_tony_van/)
- url: https://www.youtube.com/watch?v=o2V544cHNKE
---

## [3][Five Awesome C++ Papers for the Prague ISO Meeting and C++20 Status](https://www.reddit.com/r/cpp/comments/f1njuc/five_awesome_c_papers_for_the_prague_iso_meeting/)
- url: https://www.bfilipek.com/2020/02/prague.html
---

## [4][Thought/question on library malfunction because of mismatched version](https://www.reddit.com/r/cpp/comments/f24t3i/thoughtquestion_on_library_malfunction_because_of/)
- url: https://www.reddit.com/r/cpp/comments/f24t3i/thoughtquestion_on_library_malfunction_because_of/
---
Today our software malfunction and the cause found was because someone updated the 3rd party library (cryptoPP) binary compiled from the latest source code and added the binary to our project WITHOUT updating the include headers. 

It compiles, links with no crash, warning, or anything. The hash function always returns empty. 

Normally when this happens, it either cause linker error during linking or runtime crash. But this is one of two times that happen to us. And the first time took us very long time to figure out.

I understand that if we don’t actually use anything that is different from the header than in the library binary, it will still compile and link. But this certainly isn’t the case.

But the question is how to better handle and spot this problem? Besides remembering updating the headers.
## [5][metastring mini library](https://www.reddit.com/r/cpp/comments/f1tyg9/metastring_mini_library/)
- url: https://www.reddit.com/r/cpp/comments/f1tyg9/metastring_mini_library/
---
https://github.com/ConorWilliams/metastring/tree/master

I've been playing around with compile time strings for another project and have put together a mini library to make things like this easy:
 
    using type_string = meta::stom_t&lt;"abcd"&gt;;
    
    //--$ type_string = meta::string&lt;'a', 'b', 'c', 'd'&gt;
## [6][GitHub - google/marl: A hybrid thread / fiber task scheduler written in C++ 11](https://www.reddit.com/r/cpp/comments/f1gwa3/github_googlemarl_a_hybrid_thread_fiber_task/)
- url: https://github.com/google/marl
---

## [7][Memory Bandwidth Napkin Math](https://www.reddit.com/r/cpp/comments/f1f7eq/memory_bandwidth_napkin_math/)
- url: https://www.forrestthewoods.com/blog/memory-bandwidth-napkin-math/
---

## [8][revdoor: A single-file C++ library for visiting revolving door combinations.](https://www.reddit.com/r/cpp/comments/f1t7d6/revdoor_a_singlefile_c_library_for_visiting/)
- url: https://github.com/dstein64/revdoor
---

## [9][Core C++ 2020 Announces Bjarne Stroustrup as Opening Keynote Speaker](https://www.reddit.com/r/cpp/comments/f1gljc/core_c_2020_announces_bjarne_stroustrup_as/)
- url: http://corecpp.org
---

## [10][Is "C++ Template Metaprogramming" by Abrahams and Gurtovoy still relevant in the C++17/20 times?](https://www.reddit.com/r/cpp/comments/f16e9o/is_c_template_metaprogramming_by_abrahams_and/)
- url: https://www.reddit.com/r/cpp/comments/f16e9o/is_c_template_metaprogramming_by_abrahams_and/
---
I'm considering buying it, but am not sure whether it's still relevant or most of it outdated / implementing what's now already available in the language or std.

Also feel free to name (better) alternatives if there are any. Thanks :)  

EDIT: Thanks everyone for the responses so far
## [11][LinqPlus:Single header C++ syntax implementing some of C#'s Linq core functions](https://www.reddit.com/r/cpp/comments/f1u3ah/linqplussingle_header_c_syntax_implementing_some/)
- url: https://github.com/PontiacGTX/LinqPlus
---

