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
## [2][C++ on Sea 2020 video - "Deep C Diving - Fast and Scalable Text Interfaces at the Bottom" - JeanHeyd Meneide](https://www.reddit.com/r/cpp/comments/hy3yzc/c_on_sea_2020_video_deep_c_diving_fast_and/)
- url: https://www.youtube.com/watch?v=X-FLGsa8LVc
---

## [3][subprocess library for c++](https://www.reddit.com/r/cpp/comments/hxqqqu/subprocess_library_for_c/)
- url: https://www.reddit.com/r/cpp/comments/hxqqqu/subprocess_library_for_c/
---
I want to share [subprocess](https://github.com/benman64/subprocess) library. Inspired by python subprocess api. If you have c++20 you can do like so

&amp;#x200B;

    // capture output. You can do this syntax if you have C++20
    CompletedProcess process = subprocess::run({"echo", "hello", "world"}, {
        .cout = PipeOption::pipe,
        // make true to throw exception
        .check = false
    });

or c++14+

    // simplest capture output.
    CompletedProcess process = subprocess::RunBuilder({"echo", "hello", "world"})
        .cout(PipeOption::pipe).run();

This is the best I could come up with given c++ is not python. I've been using it for a while in my personal projects and it's been a treat so far. I've grown to accept that it's a bit more wordy than python. API supports advanced techniques like piping one process to another. Works on Windows, Linux, &amp; Mac.

Some notes:

* Because stdout is a macro in C/C++, cout/cin/cerr terms are used instead.
* on mac/linux posix API is used for launching processes. It does have issues but should be fine for most applications.
* C++ threading is used to avoid deadlock conditions when necessary. And to help facilitate some convenience. Currently no way to provide your own threading control.
* has similar issues as python subprocess library. No API is perfect, this one is my favorite I have seen.
* On windows there is no good way to send signals. Signals work the same way as pythons library, which usually send an abrupt killing of process. Mac/Linux work as expected.

&amp;#x200B;

Thank you and enjoy.
## [4][Am I a bad C++ programmer because of prefering raw pointers over safer smart pointers?](https://www.reddit.com/r/cpp/comments/hy6uiq/am_i_a_bad_c_programmer_because_of_prefering_raw/)
- url: https://www.reddit.com/r/cpp/comments/hy6uiq/am_i_a_bad_c_programmer_because_of_prefering_raw/
---
So, this is the case: I am the kind of C++ programmer who comes from Java. I struggled at first, but I got used to C++ fast. However I made heavy use of raw pointers, I know how to properly manage memory, as I use C as well, and I wrote a basic operating system in C and Assembly. So long story short: I can manage my memory.

Another point is that I'm developing a game engine as a hobby project. In this engine I have a very specific style of writing the engine, namingly I have to allocate a lot of stuff on the heap as pointers because I make heavy use of inheritance (example: Texture - interface, GLTexture and VKTexture - opengl and vulkan abstractions of textures), and I have to delete them when 1) there are no references to an asset or the level is unloaded, and there are other criteria to meet.

So to sum it up: I like to use raw pointers. That's my style of programming. I make heavy use of the C++ standard library, don't get me wrong, it's just that I don't like smart pointers. And this gets me to my point. Many people think I'm a horrible person because of this. So what's your point on this? (If you're interested in the engine and wanna figure out why I don't use smart pointers, I can link the GitHub repository)
## [5][How do you write class templates? (Two question survey)](https://www.reddit.com/r/cpp/comments/hxvwov/how_do_you_write_class_templates_two_question/)
- url: https://forms.gle/ZukANmABYqRvWmbt9
---

## [6][C++ on Sea 2020 video - "Algorithmic and microarchitecture optimizations of C++ applications" - Alexander Maslennikov](https://www.reddit.com/r/cpp/comments/hy5836/c_on_sea_2020_video_algorithmic_and/)
- url: https://www.youtube.com/watch?v=QLHQhzy1W4Y
---

## [7][C++ on Sea video - "C++ STL best and worst performance features and how to learn from them" - Danila Kutenin](https://www.reddit.com/r/cpp/comments/hxom99/c_on_sea_video_c_stl_best_and_worst_performance/)
- url: https://www.youtube.com/watch?v=GRuX31P4Ric
---

## [8][A single-header library for build SQL query string in C++11](https://www.reddit.com/r/cpp/comments/hxy27v/a_singleheader_library_for_build_sql_query_string/)
- url: https://github.com/six-ddc/sql-builder
---

## [9][C++ on Sea 2020 video - "The Upcoming Evolution Of the Standard Library" - Arno Schoedl](https://www.reddit.com/r/cpp/comments/hxum56/c_on_sea_2020_video_the_upcoming_evolution_of_the/)
- url: https://www.youtube.com/watch?v=jvCjtGLuuvw
---

## [10][C++ project template with Meson build](https://www.reddit.com/r/cpp/comments/hxwwfy/c_project_template_with_meson_build/)
- url: https://github.com/michaelbrockus/cpp_project_template
---

## [11][Question about learning C++](https://www.reddit.com/r/cpp/comments/hy5fdk/question_about_learning_c/)
- url: https://www.reddit.com/r/cpp/comments/hy5fdk/question_about_learning_c/
---
Hi,

I am a Data Scientist intern (I know Python) and I need to learn C++for a project. I searched good learning resources online but there is a lot of contradictory informations and I would like to hear from you guys. I need to understand C++, not be proficient, so I don't want to read a 1500 pages book right now .

I don't have money to put on books rn so I'd like free online resource (If it's possible to find good ones). I found [that one](https://www.udemy.com/course/free-learn-c-tutorial-beginners) on Udemy but I'm afraid It would be a bit old.

So is there a major updates that make prior courses outdated ? Or do you guys have any reommandations for beginner ?

Thanks a lot !
