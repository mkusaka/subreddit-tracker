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
## [2][C++/Programming - budget for trainings during COVID-19. How do you utilize it right now?](https://www.reddit.com/r/cpp/comments/j78grv/cprogramming_budget_for_trainings_during_covid19/)
- url: https://www.reddit.com/r/cpp/comments/j78grv/cprogramming_budget_for_trainings_during_covid19/
---
Hello guys, normally I would join some conference in real. In my workplace I have a dedicated budget for training/self-development.  Nowadays I have no clue how to utilize it.  Can you recommend any good trainings, conferences in 2020 in Europe?
## [3][rotor v0.09 release (c++ actor framework, but not caf)](https://www.reddit.com/r/cpp/comments/j7cyai/rotor_v009_release_c_actor_framework_but_not_caf/)
- url: https://habr.com/ru/post/522588/
---

## [4][There is a need for more organized, open system with standard proposals](https://www.reddit.com/r/cpp/comments/j6nl77/there_is_a_need_for_more_organized_open_system/)
- url: https://www.reddit.com/r/cpp/comments/j6nl77/there_is_a_need_for_more_organized_open_system/
---
Once again there is a question "Does anyone know what is the status of XXX proposal?" posted on this subreddit. Myself, I had such a question about different proposals many times, and always found it hard to answer. The information seems to be all over the place - reddit posts, meeting reports, mailing lists... I think there is a dire need for system, where one can find all ongoing / closed proposals with their statuses, current revisions, and links to all the previous revisions. One could argue, that papers on http://www.open-std.org/ have all the information, but it is impossible to:
- find the newest revision of given paper,
- search for any particular paper, without knowing the year it was published,
- filter papers by subgroups / topics,
- there is no easy to find information, that the paper got rejected, abandoned etc. there only is an information, if the paper got adopted.

I think there is a need for website (maybe community-driven/created?) with all that information. Maybe also as a forum for discussion on certain proposals, rather than obscure mailing lists. Is it only me, or do anyone else thinks the same?
## [5][Other than generators, what do coroutines solve better?](https://www.reddit.com/r/cpp/comments/j6pi07/other_than_generators_what_do_coroutines_solve/)
- url: https://www.reddit.com/r/cpp/comments/j6pi07/other_than_generators_what_do_coroutines_solve/
---
I posted to cpp_questions and got answers about what coroutines do: split a thread’s *execution* from its *state*. But I still don’t understand when they are the right tool.

I wrote a chess engine. I have an alpha-beta pruned minimax tree. I don’t need a full list of child nodes (chess moves) at every depth (board), only a few up to an alpha-beta cutoff. Generators solve this problem elegantly.

I’m struggling to come up with any other situations where coroutines are the right tool. When is polling, std::async, a thread-pool, a functor, or a thread-safe container a not a more performant or simple option?

In other words, co_yield is great. But why co_await?
## [6][Winning the race against TOCTOU vulnerabilities in C &amp; C++](https://www.reddit.com/r/cpp/comments/j6vyvb/winning_the_race_against_toctou_vulnerabilities/)
- url: https://blog.sonarsource.com/winning-the-race-against-toctou-vulnerabilities
---

## [7][The Community](https://www.reddit.com/r/cpp/comments/j6oc4q/the_community/)
- url: https://thephd.github.io/the-community
---

## [8][Analyzing RSS memory leak for a C++ application](https://www.reddit.com/r/cpp/comments/j6m1xn/analyzing_rss_memory_leak_for_a_c_application/)
- url: https://www.reddit.com/r/cpp/comments/j6m1xn/analyzing_rss_memory_leak_for_a_c_application/
---
I am currently looking into what seems like a memory leak with a multithreaded C++ 11 application. The weird part about this is that when I use Valgrind Massif or Heaptrack I do not see a leak in the heap. The heap memory stays at about 15 MB but what I am noticing is that when I continue to use the application, I see that the RSS memory keeps increasing. I have also analyzed the number of threads to see if there is a thread leak, but the total number of threads does remain constant ( though the application does create and join threads a bunch of times)

Could someone give me pointers on how I can go about exploring this leak. I have used pmap to list out the pages and I do see few blocks marked as [anon] that are contributing for the increased RSS. But I am not sure how to track them back to the corresponding code. Is there a way to set break points for these blocks or is there some other way of analyzing what is contributing to this?

I read that fragmentation could cause this or that this could happen when malloc does not release memory back to the system. I would like to track down which function calls cause these bumps so that I can work on optimizing those.

Please forgive me if this is too obvious but would love to know more about this.
## [9][Announcing the closing keynote of Meeting C++ 2020](https://www.reddit.com/r/cpp/comments/j6o72y/announcing_the_closing_keynote_of_meeting_c_2020/)
- url: https://www.meetingcpp.com/meetingcpp/news/items/Announcing-the-closing-keynote-of-Meeting-Cpp-2020.html
---

## [10][Tuplify, std::tie but cooler (I hope). Structured bindings to specific members using their names or their getters.](https://www.reddit.com/r/cpp/comments/j6d3s6/tuplify_stdtie_but_cooler_i_hope_structured/)
- url: https://www.reddit.com/r/cpp/comments/j6d3s6/tuplify_stdtie_but_cooler_i_hope_structured/
---
 [https://gist.github.com/Hamondorf/5e227ae6c204321842a4b0138630e76c](https://gist.github.com/Hamondorf/5e227ae6c204321842a4b0138630e76c) 

Feedback and opinions welcome! Examples are given towards the bottom of the gist.

Edit:
I realized if a class has a reference member you can’t for a pointer to member to it, but what you can do instead is have either a member or free function (taking the class reference as an argument) that just acts an accessory for the reference. Glad the way I wrote it allows for that without changing anything :P
## [11][Does anyone know what the status of TartanLlama's p0798R3: Monadic operations for std::optional is?](https://www.reddit.com/r/cpp/comments/j6i82o/does_anyone_know_what_the_status_of_tartanllamas/)
- url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p0798r3.html
---

