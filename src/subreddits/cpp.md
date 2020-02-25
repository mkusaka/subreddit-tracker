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
## [2][How I Declare My class And Why – Howard E. Hinnant](https://www.reddit.com/r/cpp/comments/f918oz/how_i_declare_my_class_and_why_howard_e_hinnant/)
- url: http://howardhinnant.github.io/classdecl.html
---

## [3][Would you use my ABI-breaking port of an existing std:: and/or compiler? (alt title: would you put your money where your mouth is?)](https://www.reddit.com/r/cpp/comments/f94zcr/would_you_use_my_abibreaking_port_of_an_existing/)
- url: https://www.reddit.com/r/cpp/comments/f94zcr/would_you_use_my_abibreaking_port_of_an_existing/
---
Hi folks,
Like many of you here, I was pretty disappointed in the committee's recent decision to effectively maintain ABI forever*. I'd like to gain a sense of how strong that disappointment is - paraphrasing an old favorite "most people say they would vote to send litterbugs to jail, but almost nobody would support a candidate just because they want to send litterbugs to jail."

We talk about ABI a lot, but not all ABI breaks are the same. Some ABI breaks are in the library alone, while others are in the compiler as well. The ABI also has a complex relationship with the standard:

 * the ABI is not part of the standard
 * the standard can mandate that some or all implementations change their ABI (e.g. the elimination of COW std::string) 
 * some ABI breaks require standard standard support first - either new language features or better library interfaces
 * but some ABI breaks can be taken on by library owners alone 

Today, the committee is effectively unwilling to make any changes to the language that would require any implementation to break ABI. And, std:: maintainers are generally unwilling to break the ABI of their implementation.

One of the things that's a bit silly here is that a new std:: implementation based on one of the existing std:: implementations could break ABI and thereby be faster (and more conformant!) than the std:: implementation it was based on. Of course, this has happened before - this was part of the impetus for creating libc++ instead of continuing to use stdlibc++. I'd like to make such a new implementation, but with ABI incompatibility across major versions as a expectation from day0.

The core question I'd like to ask is: would you use such an implementation?

The compiler comes in because we'd like to mangle the names differently, to make incompatibilities occur at link time rather than runtime. 

So, whats the big picture here, and how does it relate to the standard?

The simple fact is, a metric ton of applications could accept an ABI break no problem. If such applications had an option of voluntarily accepting an ABI break to gain perf, some would. Once 2-5% of applications have voluntarily chosen an ABI break to gain perf, the notion that the committee needs to maintain ABI begins to stand on a much weaker foundation.

So, would this work? Would you use such an implementation? Would your employer? For those with an understanding of the std::, how much perf do today's std:: implementations leave on the table compared to the total set of available perf from ABI breaks? Who would need to own / ship / use such an implementation in order for it to be viable?

For reference, I'm a senior engineer at huge tech company that uses C++ heavily, and my software expertise is in C++. I wouldn't consider myself a C++ expert (and sometimes I wonder if anybody truly is a C++ expert), but I know enough to know what I don't know and I know enough to write safe / efficient / secure std:: code.

I'd want to fork a popular std:: implementation on Linux that has an appropriate license. Of those, I'd like to choose the implementation that probably has more users willing to accept an ABI break. Frankly, that's probably libc++.

Very curious on thoughts here.

\* I'm aware that this isn't want the committee said, but it's the effective result of their decision on Linux, as other std maintainers have posted elsewhere.
## [4][The Day The Standard Library Died](https://www.reddit.com/r/cpp/comments/f8shr6/the_day_the_standard_library_died/)
- url: https://cor3ntin.github.io/posts/abi/
---

## [5][Stefan Petersen: Embedded RTOS in C++](https://www.reddit.com/r/cpp/comments/f96zjd/stefan_petersen_embedded_rtos_in_c/)
- url: https://youtu.be/NTATehY_n10
---

## [6][So for epochs. Is there a reason we don't just prefix with a namespace?](https://www.reddit.com/r/cpp/comments/f91s6o/so_for_epochs_is_there_a_reason_we_dont_just/)
- url: https://www.reddit.com/r/cpp/comments/f91s6o/so_for_epochs_is_there_a_reason_we_dont_just/
---
So the reason epochs didn't pass iirc was avoiding multiple definitions across translation units. Why don't we just create per epochs namespaces that automatically apply to included libraries from other epochs. You could even ignore that rule in the case where there is only one definition.
## [7][C++ Weekly - Ep 208 - The Ultimate CMake / C++ Quick Start](https://www.reddit.com/r/cpp/comments/f8szn3/c_weekly_ep_208_the_ultimate_cmake_c_quick_start/)
- url: https://www.youtube.com/watch?v=YbgH7yat-Jo
---

## [8][Anyone here tried out clang-10 and clang-tidy 10?](https://www.reddit.com/r/cpp/comments/f8yor4/anyone_here_tried_out_clang10_and_clangtidy_10/)
- url: https://www.reddit.com/r/cpp/comments/f8yor4/anyone_here_tried_out_clang10_and_clangtidy_10/
---
Hello,  


It looks like clang-10 will be released sometime soon, though the release/10.x branch has been available on the LLVM github mirror. I built it and tried out both this morning. Anyone else here build the branch and try out C++20 features that clang-10 supports? Also, anyone tried out clang-10 ?  


1. What have you tried out so far and how has your experience been? 
2. Any existing checks that clang-tidy 10 improves over clang-tidy 9?  


Thank You!
## [9][69-year-old father of C++: C++ has survived for decades, what keeps its vitality?](https://www.reddit.com/r/cpp/comments/f8jcuj/69yearold_father_of_c_c_has_survived_for_decades/)
- url: https://www.youtube.com/watch?v=mWtrLdR56oA
---

## [10][References, simply](https://www.reddit.com/r/cpp/comments/f8jrfk/references_simply/)
- url: https://herbsutter.com/2020/02/23/references-simply/
---

## [11][Is it worth to pay for DialogBlocks?](https://www.reddit.com/r/cpp/comments/f93942/is_it_worth_to_pay_for_dialogblocks/)
- url: https://www.reddit.com/r/cpp/comments/f93942/is_it_worth_to_pay_for_dialogblocks/
---
Do you think is it worth to pay for DialogBlocks? Or do you prefer any other UI Designer for complex GUIs? Like wxCrafter
