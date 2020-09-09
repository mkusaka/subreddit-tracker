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
## [2][C++ is now the fastest-growing programming language](https://www.reddit.com/r/cpp/comments/ip9i8h/c_is_now_the_fastestgrowing_programming_language/)
- url: https://www.reddit.com/r/cpp/comments/ip9i8h/c_is_now_the_fastestgrowing_programming_language/
---
[https://www.techrepublic.com/article/c-is-now-the-fastest-growing-programming-language/](https://www.techrepublic.com/article/c-is-now-the-fastest-growing-programming-language/)
## [3][Why are there no genuine type-lists (aka packs)?](https://www.reddit.com/r/cpp/comments/ip6nyj/why_are_there_no_genuine_typelists_aka_packs/)
- url: https://www.reddit.com/r/cpp/comments/ip6nyj/why_are_there_no_genuine_typelists_aka_packs/
---
Template parameter packs are such a powerful thing, but IMHO a very basic/fundamental thing is missing.

To get to the point, consider the ubiquitous pattern:

    template &lt;typename T&gt;
    struct MyThing
    {
        using Type = T;     // saves extra matching along the road
    };

Now, the question simply is: Why is there no built-in way to have

    template &lt;typename... Ts&gt;
    struct MyThing
    {
        using Types = Ts;
    };

On two quite different occasions the existence of such a 'type list' would have simplified my code a lot. It's hard to see how a compiler that can handle all the intricacies of packs, fold expressions, etc. could not handle a simple list like that.

Note: Please don't mention tuple before thinking about it.
## [4][Why are gcc and clang behaving differently here for brace-initialization?](https://www.reddit.com/r/cpp/comments/iotebm/why_are_gcc_and_clang_behaving_differently_here/)
- url: https://godbolt.org/z/b74MP6
---

## [5][Algorithm Books C++](https://www.reddit.com/r/cpp/comments/ipbvwz/algorithm_books_c/)
- url: https://www.reddit.com/r/cpp/comments/ipbvwz/algorithm_books_c/
---
 

This a poll for the most recommended book for **BEGINNERS** in Algorithms (CLRS and art of computer programming are for experienced). Please vote and type the reason in comments so it is easier for other beginners to start their journey. If any other book pls write in comments. Please remember this is language specific. In this post, I have kept it to C++ .

Thank you.

[View Poll](https://www.reddit.com/poll/ipbvwz)
## [6][C++ package managers](https://www.reddit.com/r/cpp/comments/ioytr0/c_package_managers/)
- url: https://www.reddit.com/r/cpp/comments/ioytr0/c_package_managers/
---
Hello

Is there a recommended package manager for C++ nowadays?
What do you use for your projects/CI? 
What would you recommend for mobile dev?

From my limited experience, I’ve tried vcpkg, which works nice until it doesn’t, then it’s a nightmare to fix. Has no clear package versioning scheme. Lacks good support for mobile. 

Conan lacks in graphics and gui libraries, it also tries to pull system libraries even when it really shouldn’t. Can fail if the binaries differ, then it tries to build from source. It also requires updating config files when a compiler is updated. It also lacks good support for mobile. 

CMake (via git submodules or fetch content), works unless you’re dependency doesn’t use cmake. 

I haven’t tried hunter nor build2’s package manager. I’m not sure if meson has some form of package management.
## [7][KDevelop 5.6 released](https://www.reddit.com/r/cpp/comments/ioz5d4/kdevelop_56_released/)
- url: https://www.kdevelop.org/news/kdevelop-560-released
---

## [8][Crash Course in C++](https://www.reddit.com/r/cpp/comments/ip4ryw/crash_course_in_c/)
- url: https://www.reddit.com/r/cpp/comments/ip4ryw/crash_course_in_c/
---
I just finished reading the extensive C++ tutorials at learncpp.com. Is there anything I should know about the language or naming conventions that I can’t learn from a book before I fully enter the world of writing code? Any advice or tips or concepts are greatly appreciated!!
## [9][ABI versioning](https://www.reddit.com/r/cpp/comments/iorwq2/abi_versioning/)
- url: https://www.reddit.com/r/cpp/comments/iorwq2/abi_versioning/
---
Going by recent discussions on ABI stability/breakage, I believe ABI versioning is the way forward. What are the reasons why the community is not working with vendors to put in place mechanisms to support versioning the C++ ABI on each platform? This is the direction Herb Sutter proposed in N4028 if I’m correct.
## [10][Best C++ book](https://www.reddit.com/r/cpp/comments/iordgd/best_c_book/)
- url: https://www.reddit.com/r/cpp/comments/iordgd/best_c_book/
---
Now I have decided to start C++. I am complete beginner. I can learn fast. My mathematical background is very good. I prefer learning from books as they are more detailed and contain a smooth learning curve. What do you feel is the best book: C++ Primer ,  Programming: Principles and Practice Using C++, 2nd Edition, by Bjarne Stroustrup , The programming language C++ , C++ for dummies , C++ how to program , C++ crash course :a fast paced introduction, Absolute c++.

&amp;#x200B;

Hey I have finally narrowed the list of easiest and detailed C++ books to these 3 books :

1. Programming: Principles and Practice Using C++ is a Book by Bjarne Stroustrup
2. C++ Primer is a Book by  Stanley B. Lippman
3. Absolute C++ is a Book by Walter S.

Feel free to express your opinion. ☺☻

apart from these we have (this part contains other books and the reasons why they aren't in the top 3):

1)The C++ programming language -- more of a reference for experienced programmer.

2)C++ crash course -- fast paced and tried to give a lot of extra info in tables which was not shown in codes. confuses a beginner. however if you can take the speed , go for it.

3)a tour of C++ -- too short .(good for a brush up)

4) effective C++ -- too short .(good for a brush up)

5) thinking in C++ - SUPERB for concepts. Just a bit too old (2003) . no new editions.

6)C++:how to program --- Very very very long(3200 pages.) . But Superb if you have the time.

7)C++ for dummies -- not great reviews. But check it out yourself

NOTE: IF YOU WANT ANY BOOK IN ANY PROGRAMMING LANGUAGE , PRIVATE MESSAGE ME, I WILL GIVE YOU A FREE LINK. 

[View Poll](https://www.reddit.com/poll/iordgd)
## [11][C++ code for neural network and CUDA](https://www.reddit.com/r/cpp/comments/ioy954/c_code_for_neural_network_and_cuda/)
- url: https://github.com/dwha/SimpleCudaNeuralNet
---

