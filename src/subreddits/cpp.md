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
## [2][Dangerous elements of modern C++](https://www.reddit.com/r/cpp/comments/fa9g4t/dangerous_elements_of_modern_c/)
- url: https://www.reddit.com/r/cpp/comments/fa9g4t/dangerous_elements_of_modern_c/
---
The company I work for is finally going to update from C++03 this year. One question to decide on will be which version from C++11 to C++20 to use then. Regarding this discussion I think it's very important to be aware of any dangerous (as in, they can be easily abused or overused) or controversial features, that need to be decided on, strictly excluded, or educated about.
Examples that I can think of right now are:

* auto; whether to use it everywhere, sparingly, or not at all.
* constexpr; have the potential to be abused and blow up compile times

I would appreciate any suggestions for this list, as my professional experience only goes as far as C++11 and my personal experience only to C++17.

(to be clear, I'm not calling these features bad per se. Just that it might be a bad idea to allow them in a huge code base of inexperienced developers without regulations)
## [3][GitQlient: Multi-platform Git client written with Qt](https://www.reddit.com/r/cpp/comments/fa88sv/gitqlient_multiplatform_git_client_written_with_qt/)
- url: https://www.reddit.com/r/cpp/comments/fa88sv/gitqlient_multiplatform_git_client_written_with_qt/
---
**I’m happy to announce the release of GitQlient 1.0.0**

Almost 5 months ago I started this project with the idea of taking an [old app](https://github.com/tibirna/qgit) and re-factor it. The main focus was to change the the look &amp; feel and creating a Qt Creator plugin based on it. But the most important thing was to have a client I’d feel comfortable to work with.

During the first two months I focused in the separation of the UI and the logic to set a nicer MVC pattern so I could start from there. After that, I started to add new features and removed all the old widgets and functionality that I didn’t need.

For the version 1.0.0 I’ve implemented all the features that were part of GitQlientPlugin with some fixes and improvements. A way more that I initially thought. I’ve been using GitQlient for the last two weeks in my day-to-day work and I feel it’s ready for this first version!

## The binaries

You can find the binaries for GitQlient 1.0.0 on the [release section](https://github.com/francescmm/GitQlient/releases/tag/v1.0.0) on the GitHub repo:

* [GitQlient for Linux](https://github.com/francescmm/GitQlient/releases/download/v1.0.0/GitQlient-1.0.0-x86_64.AppImage) (AppImage)
* [GitQlient for Windows](https://github.com/francescmm/GitQlient/releases/download/v1.0.0/GitQlientInstaller-1.0.0.exe) (Installer)
* Download Source code

## What’s new in GitQlient 1.0.0?

You can find everything that is [in the plugin](https://www.francescmm.com/gitqlient-preparing-the-first-version/), specially this major ones:

* Multi-repository support
* Submodules
* Tags and stashes
* Blame &amp; History
* [QLogger](https://github.com/francescmm/QLogger) as log system
* New graphic tree representation
* Improved Diff view

## Future

After these release, I’ll start working on the version 2.0.0 and the plan is to have it in 3-6 months. After that I’ll plan how often I want to release major versions: whether 3, 6 or 12 months period.

For the next version the plan is to include the following features:

* **Merge:** Full merge support
* **Configuration:** Provide UI for Git config file per project
* **Commit** additions:
   * Stage changes by hunk and by line
   * Squash commits
* **Diff** features:
   * Provide UI for file vs file diff
   * Browse file diff and full commit diff
* **UI:** Stylesheet for bright colours
* **More platforms:** GitQlient for MacOSX (planned for 2.0.0 if requested)

As always, if you’d like some feature or you’re missing something in GitQlient, check that it’s not yet in the backlog and [open an issue on GitHub](https://github.com/francescmm/GitQlient/issues)!
## [4][What's your favorite static code analyzer for c++ ?](https://www.reddit.com/r/cpp/comments/fa0sza/whats_your_favorite_static_code_analyzer_for_c/)
- url: https://www.reddit.com/r/cpp/comments/fa0sza/whats_your_favorite_static_code_analyzer_for_c/
---
Hi folks, 

What kind of tools are you using to ensure code quality?
## [5][Coz vs. Sampling Profilers](https://www.reddit.com/r/cpp/comments/fa1fsf/coz_vs_sampling_profilers/)
- url: https://easyperf.net/blog/2020/02/26/coz-vs-sampling-profilers
---

## [6][Want constexpr? Use naked new!](https://www.reddit.com/r/cpp/comments/f9qujo/want_constexpr_use_naked_new/)
- url: https://www.reddit.com/r/cpp/comments/f9qujo/want_constexpr_use_naked_new/
---
So I wanted to have a `constexpr` type erased class. I would usually use `std::unique_ptr` inside but I noticed that only its default constructor is `constexpr`, the rest is not.

That felt so awkward. In C++20 we got `constexpr new`, `constexpr delete`, `constexpr std::allocator_traits`, `constexpr std::vector`, `constexpr std::string` but not `constexpr std::unique_ptr`.

Well, in the end I had to use naked `new` and `delete` and it worked. Remember when C++11 brought us `std::make_shared` but not `std::make_unique`? Yep, C++20 is in the same boat.

Pls fix in Varna.
## [7][Our Approach to Testing a Large-Scale C++ Codebase](https://www.reddit.com/r/cpp/comments/f9yxea/our_approach_to_testing_a_largescale_c_codebase/)
- url: https://pspdfkit.com/blog/2020/our-approach-testing-large-scale-cpp/
---

## [8][Microsoft Research Reinforcement Learning Open Source Fest](https://www.reddit.com/r/cpp/comments/f9ukex/microsoft_research_reinforcement_learning_open/)
- url: https://www.microsoft.com/en-us/research/academic-program/rl-open-source-fest/
---

## [9][Possible to use int main(std::vector&lt;char*&gt; argv) as entrypoint?](https://www.reddit.com/r/cpp/comments/f9zbf4/possible_to_use_int_mainstdvectorchar_argv_as/)
- url: https://www.reddit.com/r/cpp/comments/f9zbf4/possible_to_use_int_mainstdvectorchar_argv_as/
---
Hi!

I am curious if the popular compilers like gcc, clang, msvc support safer entrypoint signatures?

* `int main(std::vector&lt;char*&gt; argv);`
* `int main(std::vector&lt;std::string&gt; argv);`

Compared to the classic `int main(int argc, char **argv);`, these signatures could reduce the likelihood of memory violations in common applications, and provide a more convenient syntax to manipulate CLI argument entries as well.
## [10][Epochs, ABI breakage, std::embed, and co-routines. A pulse on the state of satisfaction with c++](https://www.reddit.com/r/cpp/comments/f9e4pa/epochs_abi_breakage_stdembed_and_coroutines_a/)
- url: https://www.reddit.com/r/cpp/comments/f9e4pa/epochs_abi_breakage_stdembed_and_coroutines_a/
---
I feel like c++ is starting to get more disgruntled voices as the years go on, especially since the c++11 days. 

For example, the frustration with how hard it is to get std::embed into the language, or folks begging for a "clean up" of the language via epochs/abi breakage. We've got the graphics workgroup which some vehemently oppose while others want (very polarized opinions on this). Apparently ranges has some serious [foot guns](https://www.fluentcpp.com/2019/09/13/the-surprising-limitations-of-c-ranges-beyond-trivial-use-cases/) regarding performance (run time and compile. Time). Or even how move semantics warrant a book now due to their complexity.

Most of these seem to stem with the mentality of keeping c++ in an industry friendly state (don't break our old code base which cost millions to develop) and a very formal process for introducing new changes to the language (write a proposal, require someone physically advocating for you in reviews, etc).

Do you feel this is just a /r/c++ bubble (vocal minority), or is this frustration an accurate representation of influential c++ figures? I am honestly starting to worry that some tech giants might end up "forking" c++ even more than they currently do via their own standard library implementation, for example Google deciding to add proprietary features to a gcc fork beyond just absiel.

While I don't think c++ is going anywhere for a while, I do fear it will become the modern day perl, where it's slowly withers away over many many years into a "ugh damn, I have to deal with this language now? No bindings for language a/b/c at least?".
## [11][Perhaps it's the case of redesigning C++'s move semantics](https://www.reddit.com/r/cpp/comments/f9ahoc/perhaps_its_the_case_of_redesigning_cs_move/)
- url: https://www.reddit.com/r/cpp/comments/f9ahoc/perhaps_its_the_case_of_redesigning_cs_move/
---
Move semantics was really needed in C++. I never enjoyed its final design, though... In [my course on advanced C++11/14](http://ltcmelo.com), this is the lecture where a lot of doubts/frustrations/criticisms emerge -- what I hear from students are things like "I thought I understood how that worked". So in my perspective, only the superficial functionality of move semantics is comprehended by the average C++ developer. Typically, the root of confusion lies in the dual behavior of rvalue references, which, as pointed by Scott Meyers a decade ago, are actually "universal" references. In the ISO Standard of C++17 and in the community as a whole, there's been an attempt to reduce the friction on this matter by introducing the terminology of forwarding references for distinguishing such template contexts. Yet, this doesn't appear to be enough of a change, given that we continue to see questions, posts and articles trying to clarify those meanings; also, suspiciously wrong source-code can be seen out there. Afterall, the main "problem" -- which is, perhaps, to have aimed at a common solution to both move semantics and perfect forwarding, by stretching even further the template argument deduction rules -- is still present. I've always wondered whether my take on this situation was erroneous or a bit exaggerated; maybe it is... I admire the C++ committee, they generally have qualified people for the respective areas, and I'm sure that all of them strive to make good decisions. However, when I see that Nicolai Josuttis, an author who I greatly respect, is [writing a book entitled "C++ Move Semantics - The Complete Guide](http://www.cppmove.com)", I feel a strong reinforcement of my point of view: an important and cross-cutting feature of language should be thoroughly intuitive and straightforward to learn; if there's market for a book exclusively about it (and I'd unfortunately agree that there is), I suspect of a red herring. To put it shortly: I no longer believe that it'd be insane to start a more accessible redesign of C++'s move semantics and to deprecate the feature in its current state.
