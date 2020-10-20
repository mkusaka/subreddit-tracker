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
## [2][Qt 6.0 Beta Released](https://www.reddit.com/r/cpp/comments/jelc9l/qt_60_beta_released/)
- url: https://www.qt.io/blog/qt-6.0-beta-released
---

## [3][[update] CodingCards.org did some improvements as suggested by the C++/reddit community](https://www.reddit.com/r/cpp/comments/jek7xg/update_codingcardsorg_did_some_improvements_as/)
- url: https://CodingCards.org
---

## [4][Should c++ just standardize an interpreted step for compilation?](https://www.reddit.com/r/cpp/comments/jee26l/should_c_just_standardize_an_interpreted_step_for/)
- url: https://www.reddit.com/r/cpp/comments/jee26l/should_c_just_standardize_an_interpreted_step_for/
---
Constexpr is your compiler, interpreting your code.  
Right now it's weakness is it's inability and "Platformness" as described in a recent post.   
But an interpreter step would solve that and many more problems.  
  
How about actual macro system, that's not just a fancy find and replace.  
Just run a for loop the types you want declared, no need for fancy template tricks.  
Or how about that enum to string,  string to enum generator.  
Or automatic struct printing, serializing, deserializing.  
\#embed? How about you just open a file and generate whatever you need with it.  
  
It could allow for a powerful reflection step, if we let actual code touch our code.  
Perhaps the @ symbol could mark areas for interpretation, just like # marks the pre processor.  
  
C++ just has to run compiled, nothing says we can't interpret some code before hand.  
  
So what are some problems this would fix for you today?  
  
  
Also I guess i just came up with https://github.com/seanbaxter/circle.   
I just remembered it last sentence, I wonder how that's going.
## [5][Join the online C++ job fair by Meeting C++ now (live event link)](https://www.reddit.com/r/cpp/comments/jeo38i/join_the_online_c_job_fair_by_meeting_c_now_live/)
- url: https://live.remo.co/e/meeting-c-online-job-fair
---

## [6][Migrating C and C++ applications from Red Hat Enterprise Linux version 7 to version 8 - Red Hat Developer](https://www.reddit.com/r/cpp/comments/jeklip/migrating_c_and_c_applications_from_red_hat/)
- url: https://developers.redhat.com/blog/2020/10/08/migrating-c-and-c-applications-from-red-hat-enterprise-linux-version-7-to-version-8/
---

## [7][Dependency Graph Generator?](https://www.reddit.com/r/cpp/comments/jemlr9/dependency_graph_generator/)
- url: https://www.reddit.com/r/cpp/comments/jemlr9/dependency_graph_generator/
---
Hi, I have this huge project with multiple folders several layers deep. I need to start paring the thing down and getting rid of dependencies I don't need anymore. Can anyone point me to a good dependency graph generator I can use to arrange the source and header files in a fashion where I can understand how they interact and know where to cut? 

I'm working on Linux. I prefer something ready to go as possible with has little monkeying around and compiling but I'll accept it if its the best way.
## [8][A design pattern for running a lot of low priority periodic functions](https://www.reddit.com/r/cpp/comments/jeepyl/a_design_pattern_for_running_a_lot_of_low/)
- url: https://www.reddit.com/r/cpp/comments/jeepyl/a_design_pattern_for_running_a_lot_of_low/
---
Hi guys,

I'm a python programmer, so I have some biases on how I would do this in python, but was wondering how people would solve this problem in C++.

&amp;#x200B;

Sometimes in an application, you have to run a lot of low priority periodic tasks like the following:

1. Every second, broadcast X on this IO.
2. Every 0.5 seconds, take a camera image and save it.
3. Every 10 seconds, see if a file exists.
4. Every 0.2 seconds XYZ

This is all stuff that happens alongside a main program, which is cpu/timing intensive and should be prioritized.

&amp;#x200B;

If I wanted to do the above in python, I can just spawn a separate process and make it launch a thread for each one of the above. I would provide some API to users to submit their monitoring job to this process (which would just pickle the method they would like my proces to run at a certain Hz). Since each thread runs on the same process, I don't have to worry about this thing consuming too much cpu besides on the one it is assigned. I can also do this via asyncio but that would require some more effort from the user.

&amp;#x200B;

So how would one do this in the world of C++? In C++ each thread is much more powerful than in Python, since each thread can run on a separate CPU. So launching a separate thread for each of these low priority jobs especially since all of them would be sleeping most of the time seems kinda wasteful. I'm sure people have solved this problem before so I want to get some inspiration before I cook anything up.
## [9][constexpr is a Platform](https://www.reddit.com/r/cpp/comments/je3me7/constexpr_is_a_platform/)
- url: https://foonathan.net/2020/10/constexpr-platform/
---

## [10][17 Smaller but Handy C++17 Features](https://www.reddit.com/r/cpp/comments/jdx37j/17_smaller_but_handy_c17_features/)
- url: https://www.bfilipek.com/2019/08/17smallercpp17features.html
---

## [11][Initialization in C++ is Seriously Bonkers](https://www.reddit.com/r/cpp/comments/jdz5cx/initialization_in_c_is_seriously_bonkers/)
- url: http://mikelui.io/2019/01/03/seriously-bonkers.html
---

