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
## [2][ASCII Earth rotating in C++. I made this animation using projections and a bit of linear algebra.](https://www.reddit.com/r/cpp/comments/i4qtol/ascii_earth_rotating_in_c_i_made_this_animation/)
- url: https://youtu.be/wgdB9Dgb0Ro
---

## [3][What happened to polymorphic_value?](https://www.reddit.com/r/cpp/comments/i4nxyx/what_happened_to_polymorphic_value/)
- url: https://www.reddit.com/r/cpp/comments/i4nxyx/what_happened_to_polymorphic_value/
---
I don't see this in the c++20 standard library. What happened to this proposal?
## [4][JSON for Modern C++ 3.9.1 (bug fix release)](https://www.reddit.com/r/cpp/comments/i4r7jr/json_for_modern_c_391_bug_fix_release/)
- url: https://github.com/nlohmann/json/releases/tag/v3.9.1
---

## [5][Writing files 5 to 9 times faster than fprintf](https://www.reddit.com/r/cpp/comments/i48j9k/writing_files_5_to_9_times_faster_than_fprintf/)
- url: http://www.zverovich.net/2020/08/04/optimal-file-buffer-size.html
---

## [6][C++ Lambda Week: Going Generic](https://www.reddit.com/r/cpp/comments/i4na1c/c_lambda_week_going_generic/)
- url: https://www.bfilipek.com/2020/08/lambda-generic.html
---

## [7][Volunteer Mac Developers needed for BOINC](https://www.reddit.com/r/cpp/comments/i4qvic/volunteer_mac_developers_needed_for_boinc/)
- url: https://boinc.berkeley.edu/trac/wiki/MacDeveloper
---

## [8][Categories for the Working C++ Programmer](https://www.reddit.com/r/cpp/comments/i4af6m/categories_for_the_working_c_programmer/)
- url: https://bartoszmilewski.com/2020/08/05/categories-for-the-working-c-programmer/
---

## [9][nonstd::narrow() - safely convert between arithmetic types](https://www.reddit.com/r/cpp/comments/i4q6xd/nonstdnarrow_safely_convert_between_arithmetic/)
- url: https://www.reddit.com/r/cpp/comments/i4q6xd/nonstdnarrow_safely_convert_between_arithmetic/
---
[`narrow&lt;To&gt;(x)`](https://github.com/denniskb/narrow) allows you to convert the value `x` to the type `To`. If `x` can be precisely represented as a `To`, `static_cast&lt;To&gt;(x)` is returned, otherwise an exception is thrown.

This came up at work. Sharing as it might be useful for others.
## [10][Python's blinker like library in C++](https://www.reddit.com/r/cpp/comments/i4mjj8/pythons_blinker_like_library_in_c/)
- url: https://www.reddit.com/r/cpp/comments/i4mjj8/pythons_blinker_like_library_in_c/
---
Hello gents,

I use blinker [https://pythonhosted.org/blinker/](https://pythonhosted.org/blinker/) in python to connect objects/function to other objects/functions. When the object containing these signals completes certain actions, they fire off signals, and the connected functions/objects can take actions. Simple stuff like

1. Assume an object has 4 signals people can connect to.
2. Run object.method
3. When method completes x1, fire off a signal. Someone subscribed to it can potentially start doing something.
4. When method completes x2, fire off a signal. Someone subscribed to it can start doing something.
5. and so on

I use the signals to send out data from an object to anyone subscribed so that they can save it or log it or use it themselves. I also use it to start other threads so it's a pretty big part of my async programming. Of course, since this is Python, the signals can send out all kind of crazy stuff, which may not be possible in C++.

&amp;#x200B;

So my question is, is there something like blinker for C++? If not, is there a different design paradigm that is preferred in C++?
## [11][[Tooling] Memcheck-cover: A valgrind's memcheck tool wrapper and HTML report generator -- v1.1](https://www.reddit.com/r/cpp/comments/i4bjrq/tooling_memcheckcover_a_valgrinds_memcheck_tool/)
- url: https://www.reddit.com/r/cpp/comments/i4bjrq/tooling_memcheckcover_a_valgrinds_memcheck_tool/
---
Hello,   


It's been 3 months since the initial release of my "Memcheck-Cover" project.

I've been improving it a bit lately and I'm proud to present you the 2nd release today : https://github.com/Farigh/memcheck-cover/releases/tag/release-1.1

This new version fixes some bugs and adds dark themes and display settings persistence.

Here is a demo : https://david-garcin.github.io/demos/memcheck-cover/index.html

Don't hesitate to try it, share it, suggest improvements to it and open issues :)

Thanks,
