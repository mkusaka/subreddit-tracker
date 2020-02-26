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
## [2][Want constexpr? Use naked new!](https://www.reddit.com/r/cpp/comments/f9qujo/want_constexpr_use_naked_new/)
- url: https://www.reddit.com/r/cpp/comments/f9qujo/want_constexpr_use_naked_new/
---
So I wanted to have a `constexpr` type erased class. I would usually use `std::unique_ptr` inside but I noticed that only its default constructor is `constexpr`, the rest is not.

That felt so awkward. In C++20 we got `constexpr new`, `constexpr delete`, `constexpr std::allocator_traits`, `constexpr std::vector`, `constexpr std::string` but not `constexpr std::unique_ptr`.

Well, in the end I had to use naked `new` and `delete` and it worked. Remember when C++11 brought us `std::make_shared` but not `std::make_unique`? Yep, C++20 is in the same boat.

Pls fix in Varna.
## [3][Epochs, ABI breakage, std::embed, and co-routines. A pulse on the state of satisfaction with c++](https://www.reddit.com/r/cpp/comments/f9e4pa/epochs_abi_breakage_stdembed_and_coroutines_a/)
- url: https://www.reddit.com/r/cpp/comments/f9e4pa/epochs_abi_breakage_stdembed_and_coroutines_a/
---
I feel like c++ is starting to get more disgruntled voices as the years go on, especially since the c++11 days. 

For example, the frustration with how hard it is to get std::embed into the language, or folks begging for a "clean up" of the language via epochs/abi breakage. We've got the graphics workgroup which some vehemently oppose while others want (very polarized opinions on this). Apparently ranges has some serious [foot guns](https://www.fluentcpp.com/2019/09/13/the-surprising-limitations-of-c-ranges-beyond-trivial-use-cases/) regarding performance (run time and compile. Time). Or even how move semantics warrant a book now due to their complexity.

Most of these seem to stem with the mentality of keeping c++ in an industry friendly state (don't break our old code base which cost millions to develop) and a very formal process for introducing new changes to the language (write a proposal, require someone physically advocating for you in reviews, etc).

Do you feel this is just a /r/c++ bubble (vocal minority), or is this frustration an accurate representation of influential c++ figures? I am honestly starting to worry that some tech giants might end up "forking" c++ even more than they currently do via their own standard library implementation, for example Google deciding to add proprietary features to a gcc fork beyond just absiel.

While I don't think c++ is going anywhere for a while, I do fear it will become the modern day perl, where it's slowly withers away over many many years into a "ugh damn, I have to deal with this language now? No bindings for language a/b/c at least?".
## [4][Perhaps it's the case of redesigning C++'s move semantics](https://www.reddit.com/r/cpp/comments/f9ahoc/perhaps_its_the_case_of_redesigning_cs_move/)
- url: https://www.reddit.com/r/cpp/comments/f9ahoc/perhaps_its_the_case_of_redesigning_cs_move/
---
Move semantics was really needed in C++. I never enjoyed its final design, though... In [my course on advanced C++11/14](http://ltcmelo.com), this is the lecture where a lot of doubts/frustrations/criticisms emerge -- what I hear from students are things like "I thought I understood how that worked". So in my perspective, only the superficial functionality of move semantics is comprehended by the average C++ developer. Typically, the root of confusion lies in the dual behavior of rvalue references, which, as pointed by Scott Meyers a decade ago, are actually "universal" references. In the ISO Standard of C++17 and in the community as a whole, there's been an attempt to reduce the friction on this matter by introducing the terminology of forwarding references for distinguishing such template contexts. Yet, this doesn't appear to be enough of a change, given that we continue to see questions, posts and articles trying to clarify those meanings; also, suspiciously wrong source-code can be seen out there. Afterall, the main "problem" -- which is, perhaps, to have aimed at a common solution to both move semantics and perfect forwarding, by stretching even further the template argument deduction rules -- is still present. I've always wondered whether my take on this situation was erroneous or a bit exaggerated; maybe it is... I admire the C++ committee, they generally have qualified people for the respective areas, and I'm sure that all of them strive to make good decisions. However, when I see that Nicolai Josuttis, an author who I greatly respect, is [writing a book entitled "C++ Move Semantics - The Complete Guide](http://www.cppmove.com)", I feel a strong reinforcement of my point of view: an important and cross-cutting feature of language should be thoroughly intuitive and straightforward to learn; if there's market for a book exclusively about it (and I'd unfortunately agree that there is), I suspect of a red herring. To put it shortly: I no longer believe that it'd be insane to start a more accessible redesign of C++'s move semantics and to deprecate the feature in its current state.
## [5][Today I learned that you can declare a virtual method to be pure even if it has a definition in a base class](https://www.reddit.com/r/cpp/comments/f9dgxf/today_i_learned_that_you_can_declare_a_virtual/)
- url: https://www.reddit.com/r/cpp/comments/f9dgxf/today_i_learned_that_you_can_declare_a_virtual/
---
Today I was asked by someone new to C++ whether `= 0` could be used anywhere other than the first declaration of a virtual method. I said that I'd only ever seen it used on the base-most declaration, but that I honestly didn't know whether it could be used elsewhere.

So after some quick testing, it turns out it can: https://godbolt.org/z/Kqv-wx

I'm curious if anyone has ever had a good, real-life use case for this feature? I can't help but feel that if you encounter the need to "repurify" a virtual, you've probably already gone too deep down the inheritance rabbit hole.
## [6][IMO very few languages have as cleaner way of dealing with const as C++ does.](https://www.reddit.com/r/cpp/comments/f9i0el/imo_very_few_languages_have_as_cleaner_way_of/)
- url: https://whackylabs.com/swift/cpp/2019/11/28/swift-vs-c-mutability/
---

## [7][Calling shell commands from C++](https://www.reddit.com/r/cpp/comments/f9q465/calling_shell_commands_from_c/)
- url: https://freepo.st/freepost.cgi/post/8tK9EqO5lM
---

## [8][YAWYSIWYGEE - MIT licensed Qt widget for typeset equations](https://www.reddit.com/r/cpp/comments/f9ium9/yawysiwygee_mit_licensed_qt_widget_for_typeset/)
- url: https://www.reddit.com/r/cpp/comments/f9ium9/yawysiwygee_mit_licensed_qt_widget_for_typeset/
---
GitHub: [https://github.com/JohnDTill/YAWYSIWYGEE](https://github.com/JohnDTill/YAWYSIWYGEE)

WebAssembly demo:  [https://johndtill.github.io/YAWYSIWYGEE\_WASM/](https://johndtill.github.io/YAWYSIWYGEE_WASM/) 

Incorporate mathematical notation and easily access Unicode in your scientific application using YAWYSIWYGEE (yet another what-you-see-is-what-you-get equation editor). I have long held an interest in computer algebra systems, but there exists a shortage of ways to include typesetting in a C++ app of your own making. YAWYSIWYGEE solves this problem by providing a typesetting widget with all the basic interaction methods of a text editor. The widget defines a serialized representation with `toCode` and `setCode` methods so that backend logic in your app can work with a stable string format, and code defining typeset constructs can be injected into the document, e.g. pressing a GUI button creates a matrix.

The widget has reached a point where it is useful and mostly stable, although I would appreciate input on the codebase as this is my first medium-sized project! The GitHub wiki contains a few short articles describing usage and high-level design.
## [9][Would you use my ABI-breaking port of an existing std:: and/or compiler? (alt title: would you put your money where your mouth is?)](https://www.reddit.com/r/cpp/comments/f94zcr/would_you_use_my_abibreaking_port_of_an_existing/)
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
## [10][HPX 1.4.1 Released! -- The STE||AR Group](https://www.reddit.com/r/cpp/comments/f9cnt5/hpx_141_released_the_stear_group/)
- url: http://stellar-group.org/2020/02/hpx-1-4-1-released/
---

## [11][How I Declare My class And Why – Howard E. Hinnant](https://www.reddit.com/r/cpp/comments/f918oz/how_i_declare_my_class_and_why_howard_e_hinnant/)
- url: http://howardhinnant.github.io/classdecl.html
---

