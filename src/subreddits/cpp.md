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
## [2][concurrencpp v.0.0.6 - modern concurrency for C++](https://www.reddit.com/r/cpp/comments/j04one/concurrencpp_v006_modern_concurrency_for_c/)
- url: https://github.com/David-Haim/concurrencpp
---

## [3][Simple C++ Socket Programming Base / Library: Works for Windows and Linux](https://www.reddit.com/r/cpp/comments/izvy9i/simple_c_socket_programming_base_library_works/)
- url: https://www.reddit.com/r/cpp/comments/izvy9i/simple_c_socket_programming_base_library_works/
---
 Thought this might be useful for everyone doing entry-level socket programming in C++. Currently C++ does not have a unified way to write socket programming for both Windows and Linux. I tried to bridge that gap and write a simple server/client library that can be extended to do more.

The alternatives out there tend to use or depend on external libraries/runtime software which I did not want in my project. Here is what I came up with; feel free to leave feedback and contribute to it:

&amp;#x200B;

[https://github.com/masesk/easy-socket](https://github.com/masesk/easy-socket)
## [4][std::exchange Patterns: Fast, Safe, Expressive, and Probably Underused](https://www.reddit.com/r/cpp/comments/izjgdw/stdexchange_patterns_fast_safe_expressive_and/)
- url: https://www.fluentcpp.com/2020/09/25/stdexchange-patterns-fast-safe-expressive-and-probably-underused/
---

## [5][Structured binding may be the new hotness, but we'll always have std::tie](https://www.reddit.com/r/cpp/comments/izr5pf/structured_binding_may_be_the_new_hotness_but/)
- url: https://devblogs.microsoft.com/oldnewthing/20200925-00/?p=104297
---

## [6][Free hands-on online Conan C++ package manager training](https://www.reddit.com/r/cpp/comments/iznhw6/free_handson_online_conan_c_package_manager/)
- url: https://blog.conan.io/2020/09/24/New-conan-training-series.html
---

## [7][Meeting C++ online - Basit Ayantunde - Functional Error and Optional value Handling with STX](https://www.reddit.com/r/cpp/comments/izfknm/meeting_c_online_basit_ayantunde_functional_error/)
- url: https://www.youtube.com/watch?v=8CZhJa8UJk0
---

## [8][Plenary: The Beauty and Power of "Primitive" C++ - Bjarne Stroustrup - CppCon 2020](https://www.reddit.com/r/cpp/comments/iziw57/plenary_the_beauty_and_power_of_primitive_c/)
- url: https://youtu.be/ERzENfQ51Ck
---

## [9][Tests dont make code higher quality.](https://www.reddit.com/r/cpp/comments/j04ws9/tests_dont_make_code_higher_quality/)
- url: https://www.reddit.com/r/cpp/comments/j04ws9/tests_dont_make_code_higher_quality/
---
The code is what it is, writing a seperate test that mocks bad input that you already took in account and handled, achieves absolutely nothing, passing a bunch of tests is nothing more than an illusion. So if you break something by accident, the tests will tell you exactly where it is, provided you have 100% coverage, saving you at most 1 minute in a debugger, is it really worth it to increase your codebase by 75% to maybe save a minute later? You cant test for what you don't know, so code will still break for edge cases you are not yet aware of, instead of just fixing it when they arise, you need to fix it then fix a test or write a new test, if you decide to change something in the code you have to change the tests. Its insanity, seems like it was invented by programmers to get extra milage out of jobs and to trick management into a false sense of security.
## [10][Do you have a preferred library for constexpr data structures?](https://www.reddit.com/r/cpp/comments/izhmre/do_you_have_a_preferred_library_for_constexpr/)
- url: https://www.reddit.com/r/cpp/comments/izhmre/do_you_have_a_preferred_library_for_constexpr/
---
Do y'all have a preferred library of constexpr data structures? I'm trying to find something like the embedded STL [https://www.etlcpp.com/](https://www.etlcpp.com/)  but with constexpr support. I was wondering if there were any  alternatives. I've found quite a few header only implementation of  single data structures, but not a solid colllection. This also isn't  strictly embedded, but I assume the people likely to use this are people  working in embedded. 

I'm working on designing a real time system, and so I'm trying to put as much of the work at compile time. Not necessarily embedded, but doing any unnecessary work at runtime is not preferred. 

Also while I'm here, given that information, should I implement a threadpool, scheduler, and event system myself? I wasn't able to find a thread pool with considerations like this (liberal use of shared\_ptr on a few implmentations), so I'm curious to see if anyone could point me in the right direction.
## [11][The Sneaky base class constructor trick](https://www.reddit.com/r/cpp/comments/izqm9r/the_sneaky_base_class_constructor_trick/)
- url: https://www.reddit.com/r/cpp/comments/izqm9r/the_sneaky_base_class_constructor_trick/
---
Hello good people of r/cpp,

I'd like to share a simple trick for sneakily initializing a base class before the derived ctor, and without the derived class knowing anything about it.

Consider : you're writing an API for an interface class that a lot of people are going to use to build their derived objects.   
\[EDIT 2\] : These derived objects will be registered in a Factory which will create them at runtime in the environment you are developing. 

Obviously, the cleanest way to initialize it is to encapsulate the ctor arguments of your base class into a struct, that way you can silently adapt it without breaking client code. But ! The size of the client codebase might be pretty big, hundreds, perhaps thousands of derived objects will be written, that's a lot of "Derived() : Base( BaseCtorArgs )" to be copy pasted everywhere.  A minor issue, but why do your users should remember the name of your struct + type it ?

A common alternative is two steps initialization : tell the users of your API to write their construction in a virtual setup() function where the state of the base object is guaranteed to be correct. That's not great : errors will certainly happen, and perhaps the member of their derived classes do not have a default ctor, which would force them to write both a ctor and a setup function. Eurgh.

A less common alternative is what i call "the sneaky base constructor", add whatever data you need  to perform base initialization as a static member somewhere (e.g. in your Factory or Base class), have your factory set it up, and use it from your Base class constructor. [See demo on coliru.](http://coliru.stacked-crooked.com/a/c9f45931e1495814)

The pros : no one has to type anything to construct the Base class, and correct state is guaranteed.The cons : no serious cons that i can see of... except for the fringe case where you have a lot of factories in separate threads trying to create at the same time at a great rate, and the sync mechanism + indirection might hit performance, but even that can be taken care of by having an "active factory" variable for each thread.

What do you think ?

(ps : i previously tried to do this by allocating raw bytes, call static\_cast&lt;Base\*&gt; and initialize, then call a placement new Derived() on the address i had setup, which didn't worked)

EDIT : Here is the second version : no static member, but raw allocation + static\_cast + placement new. As expected, some type will be rewritten by the constructor, some won't (hence why i used void\* instead of std::string). This looks tricky and dangerous, but i don't understand why that would be undefined behavior, it's simply interpreting bytes and doing assignment.[http://coliru.stacked-crooked.com/a/66c416f1da4a32bc](http://coliru.stacked-crooked.com/a/66c416f1da4a32bc)
