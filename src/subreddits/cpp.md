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
## [2][Oh No! More Modern CMake - Deniz Bahadir - Meeting C++ 2019](https://www.reddit.com/r/cpp/comments/eje57u/oh_no_more_modern_cmake_deniz_bahadir_meeting_c/)
- url: https://www.youtube.com/watch?v=y9kSr5enrSk
---

## [3][C++ Pattern Matching proposal](https://www.reddit.com/r/cpp/comments/ejg1yc/c_pattern_matching_proposal/)
- url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1371r1.pdf
---

## [4][The Debug heap that created bugs](https://www.reddit.com/r/cpp/comments/ej39ma/the_debug_heap_that_created_bugs/)
- url: http://lectem.github.io/windows/heap/appverifier/detours/2020/01/02/The-debug-heap-that-created-bugs.html
---

## [5][CppCast: C++ 2020 News](https://www.reddit.com/r/cpp/comments/ej98qn/cppcast_c_2020_news/)
- url: https://cppcast.com/cpp-2020-news/
---

## [6][Any active C++ alternative operating system projects underway, or does Redox OS written in Rust have zero competition?](https://www.reddit.com/r/cpp/comments/ej403h/any_active_c_alternative_operating_system/)
- url: https://www.reddit.com/r/cpp/comments/ej403h/any_active_c_alternative_operating_system/
---
Am aware of Haiku, and I guess it may be using C++, not sure if it's used for their kernel layers, though...

Redox OS is using Rust for everything. They've even been re-writing the popular UNIX flavored command line tools in Rust. However, Redox OS will allow for, say, device drivers to be written in other languages (because it's a micro kernel so drivers don't directly link to the kernel). And the OS API is POSIX so user programs can be written in anything.

MINIX3 is all C code (though does some microkernel architecture stuff too).

My interest is to identify new operating system implementations that can replace Linux as a user-facing operating system (for the long term - not the near term, of course).

I'd like to see a microkernel architecture where drivers aren't linked to the kernel. But to me the holy grail would be ability to use Linux device drivers. Kind of tough, right? Given the Linux device driver model. But I have some ideas of how to solve that.
## [7][Guaranteed copy elision for named return values](https://www.reddit.com/r/cpp/comments/ej3fvy/guaranteed_copy_elision_for_named_return_values/)
- url: https://www.reddit.com/r/cpp/comments/ej3fvy/guaranteed_copy_elision_for_named_return_values/
---
Two days ago, I've [posted](https://www.reddit.com/r/cpp/comments/ei83am/alias_expressions/) a proposal draft on "alias expressions". I've got two primary reactions:

1. Yay, less moves 👍
2. But it breaks some current rules and requres us to write extra weird syntax in order to achieve that 😕

I'm not sure, but it looks like I've come up with some promising wording to make copy elision guaranteed. This will apply to existing code, no changes required.

[The proposal link](https://gist.github.com/Anton3/594141354ff9625db0b85775799312c7)
## [8][phmap now provides btree containers in addition to the fast hash maps.](https://www.reddit.com/r/cpp/comments/eipcq1/phmap_now_provides_btree_containers_in_addition/)
- url: https://www.reddit.com/r/cpp/comments/eipcq1/phmap_now_provides_btree_containers_in_addition/
---
Phmap is a header only library which has provided very efficient (both time and space) hash maps for a while. Now it also provides similarly efficient btree alternatives for std::map and std::set, for when ordered containers are needed. 

See https://github.com/greg7mdp/parallel-hashmap.
## [9][File Local Namespace - Header Only Library](https://www.reddit.com/r/cpp/comments/eidzh1/file_local_namespace_header_only_library/)
- url: https://www.reddit.com/r/cpp/comments/eidzh1/file_local_namespace_header_only_library/
---
What do you guys think of this little thing I wrote? Avoid global scope name pollution by using file local namespaces! https://github.com/Cresspresso/file_local_namespace
(Although I still recommend writing the full namespaces for the sake of being explicit, because the global namespace is polluted already.)
## [10][Alias expressions](https://www.reddit.com/r/cpp/comments/ei83am/alias_expressions/)
- url: https://www.reddit.com/r/cpp/comments/ei83am/alias_expressions/
---
Hi! It's 1 hour till the New Year where I live, but I'm sitting there finishing a proposal draft. It aims to provide "guaranteed copy elision" for more cases, and together with [P0927 on lazy parameters](https://wg21.link/p0927) and [P1144 on relocation](https://wg21.link/p01144) could make writing move constructors a thing of the past.

[Link to the proposal](https://gist.github.com/Anton3/f62dde10fdc6c3ae9d21650f54656157)

There may be some horrible typos, but feel free to check it out, help brainstorm the problem, bikeshed the syntax or to express great indignation. Happy New Year! 🎄
## [11][C++ at the end of 2019](https://www.reddit.com/r/cpp/comments/ei0zut/c_at_the_end_of_2019/)
- url: https://www.bfilipek.com/2019/12/cpp-status-2019.html
---

