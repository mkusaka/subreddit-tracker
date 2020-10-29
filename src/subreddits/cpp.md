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
## [2][Is C++ really that hard for beginners?](https://www.reddit.com/r/cpp/comments/jk6jcz/is_c_really_that_hard_for_beginners/)
- url: https://www.reddit.com/r/cpp/comments/jk6jcz/is_c_really_that_hard_for_beginners/
---
hello, i have always heard that "C++ is hard for beginner, manual memory management and stuff...", but is it really that hard after the changes c++ had?, i always avoided c++ (started with it and i got fustrated) but since c++ is the norm for game dev i kinda wanna use it (well maybe now or untill i learn java or something idk...)

and yes i know c++ doesn't have a garbage collection and it has pointers and stuff..

thanks.
## [3][ewig, an efficient text editor written in C++, using immutable data structures](https://www.reddit.com/r/cpp/comments/jjsx29/ewig_an_efficient_text_editor_written_in_c_using/)
- url: https://github.com/arximboldi/ewig
---

## [4][Subscript Assignment Operator?](https://www.reddit.com/r/cpp/comments/jjzy4m/subscript_assignment_operator/)
- url: https://www.reddit.com/r/cpp/comments/jjzy4m/subscript_assignment_operator/
---
I had a random idea recently, which is to have a separate `operator[]=` for assigning to elements of a container. Code like `my_object[my_index] = data`, would result in a call to `operator[]=` instead of `operator[]` and `operator=`.

Currently, container authors (specifically for containers that implicitly create elements in `operator[]`, like `std::map`) need to ensure that `operator[]` returns a non-const reference, so that code like `my_map[my_index] = data;` works. But this then means that code like `auto data = my_map[my_index];` fails to compile when `my_map` is a const reference, which is common for function arguments. In theory the container author could provide a const-qualified overload of `operator[]` that returns a const reference (`std::vector` does this) but in the case of containers with implicit creation, things start to get weird. Code like `auto data = my_map[my_index];` would have different behaviour when `my_index` is not in the map depending on whether `my_map` is const or not. If it's const, an error gets reported (probably via exception), and if not, the item gets inserted. Smells like a footgun.

But if there is a separate `operator[]=`, that problem goes away. `operator[]` can have a const-qualified overload only, it can return a const reference (or a copy depending on the container and stored type), and could even potentially do something like return `std::optional` when the item isn't in the container. Then `operator[]=` can have a non-const overload only, and return a reference to the item, whether newly created or not.

Thoughts? Worth proposing? Any potential issues with this design?
## [5][while (auto x=y; z)](https://www.reddit.com/r/cpp/comments/jk9i4z/while_auto_xy_z/)
- url: https://quuxplusone.github.io/blog/2020/10/28/while-with-initializer/
---

## [6][Qt6 to ship with conan](https://www.reddit.com/r/cpp/comments/jjnqoq/qt6_to_ship_with_conan/)
- url: https://www.qt.io/blog/qt-6-additional-libraries-via-package-manager
---

## [7][StructLayout: Visual Studio Extension for C++ struct memory layout visualization](https://www.reddit.com/r/cpp/comments/jji9ej/structlayout_visual_studio_extension_for_c_struct/)
- url: https://github.com/Viladoman/StructLayout
---

## [8][Jonathan Boccaras Meeting C++ 2020 opening keynote: meta-polymorphism](https://www.reddit.com/r/cpp/comments/jk72u1/jonathan_boccaras_meeting_c_2020_opening_keynote/)
- url: https://meetingcpp.com/meetingcpp/news/items/Jonathan-Boccaras-opening-keynote-meta-polymorphism.html
---

## [9][Multi-Level Break in C++ via IIFE](https://www.reddit.com/r/cpp/comments/jjpgpe/multilevel_break_in_c_via_iife/)
- url: https://artificial-mind.net/blog/2020/10/28/multi-level-break-iife
---

## [10][The new Oat++ version 1.2.0 is finally released!](https://www.reddit.com/r/cpp/comments/jjdk5z/the_new_oat_version_120_is_finally_released/)
- url: https://github.com/oatpp/oatpp
---

## [11][Fun with Concepts: Do You Even Lift, Bool?](https://www.reddit.com/r/cpp/comments/jj48jm/fun_with_concepts_do_you_even_lift_bool/)
- url: https://vector-of-bool.github.io/2020/10/26/strong-bool.html
---

