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
## [2][[Question] What are convincing (preferably perf-unrelated) ways to argue against passing everything by shared_ptr?](https://www.reddit.com/r/cpp/comments/hv45lw/question_what_are_convincing_preferably/)
- url: https://www.reddit.com/r/cpp/comments/hv45lw/question_what_are_convincing_preferably/
---
Since C++11, I noticed that programmers in some circles started using `std::shared_ptr` for everything (possibly type-aliasing it as e.g.`Ptr`) and passing everything by `shared_ptr` copy, but this makes me personally uneasy due to the loss in visibility into RAII-based resource release times.

Unfortunately, I'm not good at coming up with the kind of concrete arguments that truly speak to people, so to say, so here I am.

Most notably, I'm having a hard time conveying why destruction times are Important™, and why you don't want your object relationships to turn into a strongly connected graph.

What's your guys' take? How would you (amiably) approach the issue?
## [3][Optimizing 128-bit Division](https://www.reddit.com/r/cpp/comments/hurl3f/optimizing_128bit_division/)
- url: https://danlark.org/2020/06/14/128-bit-division/
---

## [4][std::ref and std::reference_wrapper: common use cases](https://www.reddit.com/r/cpp/comments/hv6d9k/stdref_and_stdreference_wrapper_common_use_cases/)
- url: https://www.nextptr.com/tutorial/ta1441164581/stdref-and-stdreference_wrapper-common-use-cases
---

## [5][Versioning and other C++ Boilerplate](https://www.reddit.com/r/cpp/comments/huuisv/versioning_and_other_c_boilerplate/)
- url: https://thephd.github.io/versioning-and-other-boilerplate
---

## [6][TIL static_assert error messages are not printed consistently across compilers.](https://www.reddit.com/r/cpp/comments/hupykm/til_static_assert_error_messages_are_not_printed/)
- url: https://www.reddit.com/r/cpp/comments/hupykm/til_static_assert_error_messages_are_not_printed/
---
Obligatory compiler explorer: [https://godbolt.org/z/bcaxGs](https://godbolt.org/z/bcaxGs)

Ran into this today when I realized that my nice helpful static assertion messages were being turned into absolute garbage by Clang. Google yielded [this SO post](https://stackoverflow.com/questions/46376309/why-arent-newlines-other-control-sequences-printed-with-static-assert) quoting this excerpt from the standard:

&gt;the resulting diagnostic message (4.1) shall include the text of the string-literal, if one is supplied, except that characters not in the basic source character set (5.3) are not required to appear in the diagnostic message.

I ended up getting around it by offering a simpler version of my error messages on clang, which is... 'fine', but it almost defies understanding that the standard isn't stricter about it to begin with. Does anyone know the historical context behind this, or if there's any existing moves to clear it up? The current discrepancy between the way the messages are rendered significantly hinders the usefulness of allowing messages in static assertions to begin with, imo.
## [7][Using std::map Wisely With Modern C++](https://www.reddit.com/r/cpp/comments/hulelk/using_stdmap_wisely_with_modern_c/)
- url: http://www.vishalchovatiya.com/using-std-map-wisely-with-modern-cpp/
---

## [8][Advanced C/C++ side projects](https://www.reddit.com/r/cpp/comments/huh7d2/advanced_cc_side_projects/)
- url: https://www.reddit.com/r/cpp/comments/huh7d2/advanced_cc_side_projects/
---
Hi.

After a resume review by a fellow engineer, I was told that I should build some advanced C and C++ projects to show my competency.  Although my long term goal is to become a sensor fusion engineer, now I am just looking for entry-level positions.  The industry I want to get into is self-drivig cars/robotics/autonomy development.  What are some advanced C and C++ based projects that I can build along the lines of this industry?  These projects are also going to help me learn the languages a lot better as well.

I apologize if this is a dumb post. I figured that this is a good place to ask.

Thanks.
## [9][C++ audio projects](https://www.reddit.com/r/cpp/comments/hv2930/c_audio_projects/)
- url: https://www.reddit.com/r/cpp/comments/hv2930/c_audio_projects/
---
Hello,

I'm more of a Javascript developer but I've been getting deep into C++ the past few months. My goal is to be an audio developer. Can anyone suggest any good projects to build or good resources? I have built a synthesizer simulator on the keyboard but trying to think of others that are in a way easy to tackle.
## [10][build2 0.13.0 released, adds support for ad hoc recipes, source layout customization](https://www.reddit.com/r/cpp/comments/hum1nw/build2_0130_released_adds_support_for_ad_hoc/)
- url: https://build2.org/release/0.13.0.xhtml
---

## [11][The ABI stability matryoshka](https://www.reddit.com/r/cpp/comments/hum7oz/the_abi_stability_matryoshka/)
- url: https://nibblestew.blogspot.com/2020/07/the-abi-stability-matryoshka.html
---

