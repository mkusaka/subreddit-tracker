# cpp
## [1][C++ Jobs - Q1 2020](https://www.reddit.com/r/cpp/comments/eiila4/c_jobs_q1_2020/)
- url: https://www.reddit.com/r/cpp/comments/eiila4/c_jobs_q1_2020/
---
Rules For Individuals
---------------------

* **Don't** create top-level comments - those are for employers.
* Feel free to reply to top-level comments with **on-topic** questions.
* I will create one top-level comment for **meta** discussion.
* I will create another top-level comment for **individuals looking for work** and **community groups looking for sponsors**.

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

\*\*Remote:\*\* [Do you offer the option of working remotely? If so, do you require employees to live in certain areas or time zones?]

&amp;nbsp;

\*\*Visa Sponsorship:\*\* [Does your company sponsor visas?]

&amp;nbsp;

\*\*Technologies:\*\* [Required: do you mainly use C++98/03, C++11, C++14, C++17, or the C++20 working draft? Optional: do you use Linux/Mac/Windows, are there languages you use in addition to C++, are there technologies like OpenGL or libraries like Boost that you need/want/like experience with, etc.]

&amp;nbsp;

\*\*Contact:\*\* [How do you want to be contacted? Email, reddit PM, telepathy, gravitational waves?]

---

Previous Post
--------------

* [C++ Jobs - Q4 2019](https://www.reddit.com/r/cpp/comments/dbqgbw/c_jobs_q4_2019/)
## [2][Investigating the Performance Overhead of C++ Exceptions](https://www.reddit.com/r/cpp/comments/f2e8l3/investigating_the_performance_overhead_of_c/)
- url: https://pspdfkit.com/blog/2020/performance-overhead-of-exceptions-in-cpp/
---

## [3][What's the most annoying bug you've come across?](https://www.reddit.com/r/cpp/comments/f2m9vm/whats_the_most_annoying_bug_youve_come_across/)
- url: https://www.reddit.com/r/cpp/comments/f2m9vm/whats_the_most_annoying_bug_youve_come_across/
---

## [4][tser - Tiny Serialization for C++](https://www.reddit.com/r/cpp/comments/f2b0eo/tser_tiny_serialization_for_c/)
- url: https://www.reddit.com/r/cpp/comments/f2b0eo/tser_tiny_serialization_for_c/
---
[https://github.com/KonanM/tser](https://github.com/KonanM/tser)

I thought some of you might find this library useful - I developed this library for solving challenges on [codingame.com](https://codingame.com) . For the challenges there you have a 100KB size limit for all the code, so I needed a tiny serialization/pretty printing library and wanted to develop something under 20KB. That's how I came up with tser. 

**What does tser solve?**

* serialization of nearly all of the STL containers and types, c-arrays and raw pointer
* base64 encoding support for printing serialized objects to std::cout
* pretty printing of objects with member names
* automatically implement comparison operators
* all in a single header with \~350 lines of code

As stated in the readme of this library - there are much better non intrusive, macro free libraries out there. Pick tser only if you need something small and placing a macro into the classes you need to serialize is acceptable.

What would be interesting for me is to get some feedback if I should make the library even smaller (or if there is a crucial feature missing?)... for example I personally didn't need the pointer or c-array support yet, so I'm not sure if I should drop support for them.  I could also see dropping support for custom save/load methods (e.g. when you really need a compact storage for that big vector&lt;bool&gt;).
## [5][Combining ZeroMQ &amp; POSIX signals: Use ppoll to handle EINTR once and for all](https://www.reddit.com/r/cpp/comments/f2r1i6/combining_zeromq_posix_signals_use_ppoll_to/)
- url: https://blog.esciencecenter.nl/combining-zeromq-posix-signals-b754f6f29cd6
---

## [6][Avast C++ Meetup with Bjarne Stroustrup, Tony van Eerd, and Herb Sutter (Stream on 11th February 2020 6pm CET, 5pm UTC, 9am PCT, 12pm EST)](https://www.reddit.com/r/cpp/comments/f2409z/avast_c_meetup_with_bjarne_stroustrup_tony_van/)
- url: https://www.youtube.com/watch?v=o2V544cHNKE
---

## [7][Question about good OOP object usage technique.](https://www.reddit.com/r/cpp/comments/f2hnpl/question_about_good_oop_object_usage_technique/)
- url: https://www.reddit.com/r/cpp/comments/f2hnpl/question_about_good_oop_object_usage_technique/
---
Hello there. I am semi-new to C++ OOP. There are several things i still can't wrap my head around. For example i have class "Engine", which has an object 'player' of a class "Player". The "Engine" class controls input and thus updates player's position based on this input. However, my second class called "World" is used for rendering stuff, in which i would like player to be also rendered. How should i handle situation like this? Pass the reference to the player object using constructor of "World" class (But that would look messy imho). Any suggestion, tips and good advices are welcome.
## [8][What they should improve about C++ ?](https://www.reddit.com/r/cpp/comments/f2ifpt/what_they_should_improve_about_c/)
- url: https://www.reddit.com/r/cpp/comments/f2ifpt/what_they_should_improve_about_c/
---
What they should improve about C++ that you find in higher level managed language like C#, JAVA, Javascript etc or is there things you would change or do differently if it was up to you ? I mean anything like a better STL ? better thread ? string manipulation like managed language ? less fragmentations in the new standard features ? For example do you love the way the standard evolve C++ or you would prefer a direction like Microsoft do with C# a corporation who decide on the roadmap ?
## [9][Is there any Github repo that shows you every trick you need to know in order to manage the memory of any C++ application?](https://www.reddit.com/r/cpp/comments/f2kpk6/is_there_any_github_repo_that_shows_you_every/)
- url: https://www.reddit.com/r/cpp/comments/f2kpk6/is_there_any_github_repo_that_shows_you_every/
---
I think the best way to learn something is to find the simplest example that encompasses everything you need to know. That's why I am wondering if there's any good repository on Github with a C++ project that use every tool in the toolkit for memory management.
## [10][Cpp Image Processing Library](https://www.reddit.com/r/cpp/comments/f29b08/cpp_image_processing_library/)
- url: https://www.reddit.com/r/cpp/comments/f29b08/cpp_image_processing_library/
---
Hey so a couple of months ago i started developing a simple library for image processing and developing for c++ ,i got tired of including many additional dependencies  to my project as i had to do with opencv and other existing API , as for now the project is still under development but already it provides the users with a vast amount of tools , if you are looking for a simple library which only requires of you to include some header files and start working i really think my library can give you the tools needed , beside that i would be happy to get some requests and reviews on what should improve and so on , 

&amp;#x200B;

[https://github.com/MuteJester/Cpp\_Simple\_Image\_Processing\_Library](https://github.com/MuteJester/Cpp_Simple_Image_Processing_Library)

&amp;#x200B;

in the projects wiki tab you can find all instructions for working with the library .
## [11][Five Awesome C++ Papers for the Prague ISO Meeting and C++20 Status](https://www.reddit.com/r/cpp/comments/f1njuc/five_awesome_c_papers_for_the_prague_iso_meeting/)
- url: https://www.bfilipek.com/2020/02/prague.html
---

