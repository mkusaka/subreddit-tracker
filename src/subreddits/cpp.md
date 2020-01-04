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
* **Don't** use URL shorteners. [reddiquette](https://www.reddit.com/wiki/reddiquette) forbids them because they're opaque to the spam filter.
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
## [2][A &lt;windows.h&gt; wrapper for Windows apps that use UTF-8 as the Windows ANSI codepage and the C++ execution character set.](https://www.reddit.com/r/cpp/comments/ejtcot/a_windowsh_wrapper_for_windows_apps_that_use_utf8/)
- url: https://alfps.wordpress.com/2020/01/03/a-windows-h-wrapper-for-utf-8-windows-apps/
---

## [3][Dry-comparisons: A C++ Library to Shorten Redundant If Statements](https://www.reddit.com/r/cpp/comments/ejkwqb/drycomparisons_a_c_library_to_shorten_redundant/)
- url: https://www.fluentcpp.com/2020/01/03/dry-comparisons-a-c-library-to-shorten-redundant-if-statements/
---

## [4][C++ Pattern Matching proposal](https://www.reddit.com/r/cpp/comments/ejg1yc/c_pattern_matching_proposal/)
- url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1371r1.pdf
---

## [5][The most useless piece of code I have ever written](https://www.reddit.com/r/cpp/comments/ejwmwr/the_most_useless_piece_of_code_i_have_ever_written/)
- url: https://www.reddit.com/r/cpp/comments/ejwmwr/the_most_useless_piece_of_code_i_have_ever_written/
---
\#include &lt;iostream&gt;

double b(double d)

{

return d / d;

}

int main()

{

double num;

std::cout &lt;&lt; "Choose a number to divide by itself." &lt;&lt; std::endl;

std::cin &gt;&gt; num;

std::cout &lt;&lt; b(num);

}
## [6][Free and open-source C++ online code editor](https://www.reddit.com/r/cpp/comments/ejwcul/free_and_opensource_c_online_code_editor/)
- url: https://ide.judge0.com/?7U55
---

## [7][Oh No! More Modern CMake - Deniz Bahadir - Meeting C++ 2019](https://www.reddit.com/r/cpp/comments/eje57u/oh_no_more_modern_cmake_deniz_bahadir_meeting_c/)
- url: https://www.youtube.com/watch?v=y9kSr5enrSk
---

## [8][A c++ based synchronous multi-peer music player](https://www.reddit.com/r/cpp/comments/ejs99e/a_c_based_synchronous_multipeer_music_player/)
- url: https://github.com/VedantParanjape/audio-streamer
---

## [9][The Debug heap that created bugs](https://www.reddit.com/r/cpp/comments/ej39ma/the_debug_heap_that_created_bugs/)
- url: http://lectem.github.io/windows/heap/appverifier/detours/2020/01/02/The-debug-heap-that-created-bugs.html
---

## [10][CppCast: C++ 2020 News](https://www.reddit.com/r/cpp/comments/ej98qn/cppcast_c_2020_news/)
- url: https://cppcast.com/cpp-2020-news/
---

## [11][Any active C++ alternative operating system projects underway, or does Redox OS written in Rust have zero competition?](https://www.reddit.com/r/cpp/comments/ej403h/any_active_c_alternative_operating_system/)
- url: https://www.reddit.com/r/cpp/comments/ej403h/any_active_c_alternative_operating_system/
---
Am aware of Haiku, and I guess it may be using C++, not sure if it's used for their kernel layers, though...

Redox OS is using Rust for everything. They've even been re-writing the popular UNIX flavored command line tools in Rust. However, Redox OS will allow for, say, device drivers to be written in other languages (because it's a micro kernel so drivers don't directly link to the kernel). And the OS API is POSIX so user programs can be written in anything.

MINIX3 is all C code (though does some microkernel architecture stuff too).

My interest is to identify new operating system implementations that can replace Linux as a user-facing operating system (for the long term - not the near term, of course).

I'd like to see a microkernel architecture where drivers aren't linked to the kernel. But to me the holy grail would be ability to use Linux device drivers. Kind of tough, right? Given the Linux device driver model. But I have some ideas of how to solve that.
