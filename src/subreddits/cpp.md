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
## [2][Is cpp-httplib the simplest/highest-performance way to get started writing a small REST API in C++?](https://www.reddit.com/r/cpp/comments/igfn33/is_cpphttplib_the_simplesthighestperformance_way/)
- url: https://github.com/yhirose/cpp-httplib
---

## [3][Stream error functions instead of exceptions?](https://www.reddit.com/r/cpp/comments/igp033/stream_error_functions_instead_of_exceptions/)
- url: https://www.reddit.com/r/cpp/comments/igp033/stream_error_functions_instead_of_exceptions/
---
I'm just curious... why does C++ use utility functions for stream error conditions that return booleans, such as cin.bad(), [cin.fail](https://cin.fail)(), and cin.eof() instead of  just throwing related exceptions to the caller?  

If i is an int, and (cin &gt;&gt; i) fails, why not throw a related exception like a type casting exception, or an end of file exception, and allow the caller to catch it, or not?

Yet on the other hand, dynamic\_cast&lt;&gt; will throw bad\_cast exceptions?  Respectfully, it seems like it's requiring an admixture of programming styles in what can easily be the the same top level code.

P.S. I am only asking because I am trying to learn modern C++ after some professional experience over a decade ago in C++98.  I am really energized to learn the new concurrency support in C++17, but in dusting off my memory to prepare, I am coming across things like this... and going... hey.... wait...
## [4][2020-08 mailing available : Standard C++](https://www.reddit.com/r/cpp/comments/ig8rdk/202008_mailing_available_standard_c/)
- url: https://isocpp.org/blog/2020/08/2020-08-mailing-available
---

## [5][Virtual Table](https://www.reddit.com/r/cpp/comments/igy0ux/virtual_table/)
- url: https://www.reddit.com/r/cpp/comments/igy0ux/virtual_table/
---
Hello!

I am wondering if anyone has any great intuition behind the virtual table, and specifically how it results in different casting behavior between C++ and Java.

Thanks!
## [6][Pointers and Classes](https://www.reddit.com/r/cpp/comments/igxljc/pointers_and_classes/)
- url: https://www.reddit.com/r/cpp/comments/igxljc/pointers_and_classes/
---
I'm new to c++ and trying to learn the language right now. 

It seems tricky to understand Pointers in classes or as an argument.

Does anyone have a clarify explanation about class and pointer?
## [7][Is there bigdata stack using cpp?](https://www.reddit.com/r/cpp/comments/igs4h2/is_there_bigdata_stack_using_cpp/)
- url: https://www.reddit.com/r/cpp/comments/igs4h2/is_there_bigdata_stack_using_cpp/
---
As far as I know, clickhouse is OLAP system, and qfs vs hdfs. Any other data processing system using c++? There is a thrill on github but rarely update.
## [8][A Tutorial on Atomic Variables and Memory Models in C++](https://www.reddit.com/r/cpp/comments/igfeuj/a_tutorial_on_atomic_variables_and_memory_models/)
- url: https://www.youtube.com/watch?v=IE6EpkT7cJ4&amp;feature=youtu.be
---

## [9][match/wcmatch needs an at()](https://www.reddit.com/r/cpp/comments/igfe1f/matchwcmatch_needs_an_at/)
- url: https://www.reddit.com/r/cpp/comments/igfe1f/matchwcmatch_needs_an_at/
---
It would ever so nice if the regex match result had an alternative at() to use instead of the index operator. Some linting/analysis tools assume that this is an unchecked array operation and complain about it. This requires either lots of suppressions, or a lot of bogus warnings that hide real indexing operations that might be an issue.

It would be trivial to provide an at() that just turned around called the operator. Those who have this concern could avoid it and there would approximately zero cost to the implementors.
## [10][Schedule for Meeting C++ 2020 is online!](https://www.reddit.com/r/cpp/comments/ig8t55/schedule_for_meeting_c_2020_is_online/)
- url: https://meetingcpp.com/2020/Schedule.html
---

## [11][Why would I use std lib?](https://www.reddit.com/r/cpp/comments/igv1sp/why_would_i_use_std_lib/)
- url: https://www.reddit.com/r/cpp/comments/igv1sp/why_would_i_use_std_lib/
---
Not that anybody minds, but I would like to use standard C++. I am labouring under a strict single requirement to produce the smallest and fastest server-side executable possible. 

Why would I use std lib?
