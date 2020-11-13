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
## [2][Deprecating volatile - JF Bastien - CppCon 2019](https://www.reddit.com/r/cpp/comments/jtdk74/deprecating_volatile_jf_bastien_cppcon_2019/)
- url: https://www.youtube.com/watch?v=KJW_DLaVXIY
---

## [3][Compound assignment to volatile must be un-deprecated](https://www.reddit.com/r/cpp/comments/jswz3z/compound_assignment_to_volatile_must_be/)
- url: https://www.reddit.com/r/cpp/comments/jswz3z/compound_assignment_to_volatile_must_be/
---
To my horror I discovered that C++20 has deprecated compound assignments to a volatile. For those who are at a loss what that might mean: a compound assignment is += and its family, and a volatile is generally used to prevent the compiler from optimizing away reads from and/or writes to an object.

In close-to-the-metal programming volatile is the main mechanism to access memory-mapped peripheral registers. The manufacturer of the chip provides a C header file that contains things like

    #define port_a (*((volatile uint32_t *)409990))
    #define port_b (*((volatile uint32_t *)409994))

This creates the ‘register’ port\_a: something that behaves very much like a global variable. It can be read from, written to, and it can be used in a compound assignment. A very common use-case is to set or clear one bit in such a register, using a compound or-assignment or and-assignment:

    port_a |= (0x01 &lt;&lt; 3 ); // set bit 3
    port_b &amp;= ~(0x01 &lt;&lt; 4 ); // clear bit 4

In these cases the compound assignment makes the code a bit shorter, more readable, and less error-prone than the alterative with separate bit operator and assignment.  When instead of port\_a a more complex expression is used, like uart\[ 2 \].flags\[ 3 \].tx, the advantage of the compound expression is much larger.

As said, manufacturers of chips provide C header files for their chips. C, because as far as they are concerned, their chips should be programmed in C (and with \*their\* C tool only). These header files provide the register definitions, and operations on these registers, often implemented as macros. For me as C++ user it is fortunate that I can use these C headers files in C++, otherwise I would have to create them myself, which I don’t look forward to.

So far so good for me, until C++20 deprecated compound assignments to volatile. I can still use the register definitions, but my code gets a bit uglier. If need be, I can live with that. It is my code, so I can change it. But when I want to use operations that are provided as macros, or when I copy some complex manipulation of registers that is provided as an example (in C, of course), I am screwed.

Strictly speaking I am not screwed immediately, after all deprecated features only produce a warning, but I want my code to be warning-free, and todays deprecation is tomorrows removal from the language.

I can sympathise with the argument that some uses of volatile were ill-defined, but that should not result in removal from the language of a tool that is essential for small-system close-to-the-metal programming. The get a feeling for this: using a heap is generally not acceptable. Would you consider this a valid argument to deprecate the heap from C++23?

As it is, C++ is not broadly accepted in this field. Unjustly, in my opinion, so I try to make my small efforts to change this. Don’t make my effort harder and alienate this field even more by deprecating established practice.

So please, un-deprecate compound assignments to volatile.
## [4][CppCast: Video Games, Robotics and Audio](https://www.reddit.com/r/cpp/comments/jtahbn/cppcast_video_games_robotics_and_audio/)
- url: https://cppcast.com/joel-lamotte/
---

## [5][std::shared_mutex was broken in glibc&lt;2.30 distros like Ubuntu for two years](https://www.reddit.com/r/cpp/comments/jsynzn/stdshared_mutex_was_broken_in_glibc230_distros/)
- url: https://github.com/bitcoin/bitcoin/issues/16684#issuecomment-726214696
---

## [6][Overload Journal 159 · October 2020](https://www.reddit.com/r/cpp/comments/jt84v4/overload_journal_159_october_2020/)
- url: https://accu.org/journals/overload/overload159
---

## [7][Destructing outside the lock when removing items from C++ standard containers](https://www.reddit.com/r/cpp/comments/jswpfp/destructing_outside_the_lock_when_removing_items/)
- url: https://devblogs.microsoft.com/oldnewthing/20201112-00/?p=104444
---

## [8][New boost additions for 1.75 looking great](https://www.reddit.com/r/cpp/comments/jsarkz/new_boost_additions_for_175_looking_great/)
- url: https://www.reddit.com/r/cpp/comments/jsarkz/new_boost_additions_for_175_looking_great/
---
I just took a look at Boost and it has these 3 new libraries:

- Json (parse json, etc.)
- Leaf (lightweight error handling with try-catch like syntax)
- PFR (basic serialization without macros)

https://www.boost.org/users/history/in_progress.html

What do you think?
## [9][Using C++ as a scripting language, part 3](https://www.reddit.com/r/cpp/comments/jscnt7/using_c_as_a_scripting_language_part_3/)
- url: https://fwsgonzo.medium.com/using-c-as-a-scripting-language-part-3-b8f92206ef94
---

## [10][The hidden callout: The destructor](https://www.reddit.com/r/cpp/comments/jsbxff/the_hidden_callout_the_destructor/)
- url: https://devblogs.microsoft.com/oldnewthing/20201111-00/?p=104439
---

## [11][Miniselect: Practical and Generic Selection Algorithms](https://www.reddit.com/r/cpp/comments/jsba2m/miniselect_practical_and_generic_selection/)
- url: https://danlark.org/2020/11/11/miniselect-practical-and-generic-selection-algorithms/
---

