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
## [2][awesome-hpp: A curated list of awesome header-only C++ libraries](https://www.reddit.com/r/cpp/comments/icm8g1/awesomehpp_a_curated_list_of_awesome_headeronly_c/)
- url: https://github.com/p-ranav/awesome-hpp
---

## [3][Writing ideal class definition in C++](https://www.reddit.com/r/cpp/comments/icijxj/writing_ideal_class_definition_in_c/)
- url: https://www.reddit.com/r/cpp/comments/icijxj/writing_ideal_class_definition_in_c/
---
 [https://medium.com/a-devs-life/designing-an-ideal-class-in-c-d205516c03ab](https://medium.com/a-devs-life/designing-an-ideal-class-in-c-d205516c03ab) 

Guys, please provide your valuable comments.
## [4][A Rust-style guide for C++](https://www.reddit.com/r/cpp/comments/ickurq/a_ruststyle_guide_for_c/)
- url: https://www.reddit.com/r/cpp/comments/ickurq/a_ruststyle_guide_for_c/
---
Hello,

I have recently been perusing the Rust language book and found it really well-organized and presenting all the important topics concisely and ready-to-go.

Is there something similar for C++ language? I know there are some WG movements to create the ONE book for learning C++ as many other modern languages do. However, I feel that for now apart from some well-written books from gurus like Scott Meyers, there is a RTFM approach to learning this language. Or am I wrong?  


I would be happy to discuss this topic in detail, because I believe that without such book C++ is quite unapproachable for newcomers.
## [5][VS 2019 16.7.2 is now available](https://www.reddit.com/r/cpp/comments/icabt5/vs_2019_1672_is_now_available/)
- url: https://github.com/microsoft/STL/wiki/Changelog#vs-2019-167
---

## [6][sol 3: modern lua and C++ integration](https://www.reddit.com/r/cpp/comments/ic8ru0/sol_3_modern_lua_and_c_integration/)
- url: https://sol2.readthedocs.io/en/latest/
---

## [7][libstdc++ &lt;algorithm&gt; header transitively including &lt;range&gt;](https://www.reddit.com/r/cpp/comments/ibz0m2/libstdc_algorithm_header_transitively_including/)
- url: https://www.reddit.com/r/cpp/comments/ibz0m2/libstdc_algorithm_header_transitively_including/
---
I noticed some time ago that in c++20 mode the &lt;algorithm&gt; header in libstdc++ is substantially bigger than in c++11 mode :

`echo "#include &lt;algorithm&gt;" | gcc -std=c++11 -P -E -x c++ - | wc -l`

evaluates to 11760 loc, while

`echo "#include &lt;algorithm&gt;" | gcc -std=c++20 -P -E -x c++ - | wc -l`

evaluates to 45219 loc. Clicking through the header apparently in c++20 we have the include chain

&lt;algorithm&gt; -&gt; &lt;bits/ranges\_algo.h&gt; -&gt; &lt;bits/ranges\_algobase.h&gt; -&gt; &lt;range&gt;.

How could that happen? Measuring the compile time of file only including the &lt;algorithm&gt; header I get with the c++11 switch 120 ms and with the c++20 switch **600 ms** (empty file is 100 ms). I a bit baffled that while standardizing the &lt;range&gt; header this either slipped through the cracks or it was actively decided that this is ok. The &lt;algorithm&gt; header is virtually anywhere and with c++20 flag enabled every file including the &lt;algorithm&gt; header will compile half a second slower...

&amp;#x200B;

**Edit:**

Apparently the &lt;vector&gt; header also increased by a huge amount from 9530 loc to 19818 loc for libstdc++ due to now including &lt;bits/stl\_algo.h&gt; (this seems to be a known [issue](https://gcc.gnu.org/bugzilla/show_bug.cgi?id=92546) though). Compiling an empty header file only including &lt;vector&gt; is now up to 230 ms from 130 ms. 
## [8][C++20 SQLite wrapper](https://www.reddit.com/r/cpp/comments/ic629n/c20_sqlite_wrapper/)
- url: https://www.reddit.com/r/cpp/comments/ic629n/c20_sqlite_wrapper/
---
I am writing a C++20 SQLite wrapper with expressive code in mind. The code must run fast.

    sql::open("dev.db")
    | sql::query("select name, salary from person")
    | sql::for_each([](std::string_view name, float salary)
      { std::cout &lt;&lt; name &lt;&lt; "," &lt;&lt; salary &lt;&lt; std::endl; })
    | sql::onerror([](auto e){ std::cout &lt;&lt; e &lt;&lt; std::endl; });

I have some benchmarks to compare with solutions using the SQLite library. The usage of pipe operators to chain operations is optional and there is an API that throws C++ exceptions instead of using `result&lt;T&gt;` to report errors. 

[github.com/ricardocosme/msqlite](https://github.com/ricardocosme/msqlite)

What do you think? Any comments or feedbacks are welcome!
## [9][Why does c++ give a bare segmentation fault error instead of a more detailed one?](https://www.reddit.com/r/cpp/comments/ic4ol4/why_does_c_give_a_bare_segmentation_fault_error/)
- url: https://www.reddit.com/r/cpp/comments/ic4ol4/why_does_c_give_a_bare_segmentation_fault_error/
---
I started programming in python, then got into java and dart. Then I got into c++ and I'm always frustrated when I first run my code and get a segmentation fault error. Every other language I have learnt give a detaild, nice for debugging error message. Why c++ does not? Is it related to the fact that it's compiled?
## [10][Combining Modern C++ and Lua - James Pascoe](https://www.reddit.com/r/cpp/comments/ic8raa/combining_modern_c_and_lua_james_pascoe/)
- url: https://www.youtube.com/watch?v=QwfL72yHZEY
---

## [11][unique_ptr, shared_ptr, weak_ptr, or reference_wrapper for class relationships](https://www.reddit.com/r/cpp/comments/ibzqvj/unique_ptr_shared_ptr_weak_ptr_or_reference/)
- url: https://www.nextptr.com/tutorial/ta1450413058/unique_ptr-shared_ptr-weak_ptr-or-reference_wrapper-for-class-relationships
---

