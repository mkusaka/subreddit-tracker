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
## [2][The abstraction penalty for wide integer math on x86-64 ("Down with Boost.Multiprecision"?)](https://www.reddit.com/r/cpp/comments/f3mgun/the_abstraction_penalty_for_wide_integer_math_on/)
- url: https://quuxplusone.github.io/blog/2020/02/13/wide-integer-proof-of-concept/
---

## [3][Even Better C++ Performance and Productivity: Enhancing Clang to Support Just-in-Time Compilation of Templates](https://www.reddit.com/r/cpp/comments/f3jpuw/even_better_c_performance_and_productivity/)
- url: https://www.youtube.com/watch?v=pDagqR0jAvQ
---

## [4][A Fast Bitboard-Based Chess Move Generator](https://www.reddit.com/r/cpp/comments/f3pzfu/a_fast_bitboardbased_chess_move_generator/)
- url: https://www.reddit.com/r/cpp/comments/f3pzfu/a_fast_bitboardbased_chess_move_generator/
---
Over the past few months, I've been developing a fast purely legal chess move generator as my first major project in C++. It uses:

* Magic Bitboard sliding attacks
* Make-Unmake Position class (inspired by Stockfish)
* A completely original legal move generation function
* 16-bit move representation

The best part is that perft(6) and perft(7) from the starting position clock in at an impressive \~200,000,000 million NPS on my system (about the same as the leading C generator)

I am planning to incorporate this in a fully functional chess engine. I'd love it if you guys could check it out at [https://github.com/DiehardTheTryhard/surge](https://github.com/DiehardTheTryhard/surge). Thanks :)

^((sorry for the lack of comments))
## [5][I created a library that will make your C++ tests pass when run on CI](https://www.reddit.com/r/cpp/comments/f3bzof/i_created_a_library_that_will_make_your_c_tests/)
- url: https://github.com/elnormous/volkswagencpp
---

## [6][What is the exact status of ABI compatibility in modern GCC-based (including Clang) C++?](https://www.reddit.com/r/cpp/comments/f3my1u/what_is_the_exact_status_of_abi_compatibility_in/)
- url: https://www.reddit.com/r/cpp/comments/f3my1u/what_is_the_exact_status_of_abi_compatibility_in/
---
I have read in many places that C++ and its compilers make no guarantees about a stable ABI between versions of the compiler (or across platforms, even similar ones like across Linux distributions). Yet, in Linux, it is extreamly common to download C++ libraries in the package manager as binaries, and yet I don't get every library on the system updating every time my compiler updates while the new compiler keeps on linking to the existing libraries correctly. Additionally, I have read conflicting information saying that modern GCC and clang both have a standard ABI (how standard? How does it work across OS's on the same CPU architeture? how modern is modern?). What is going on here? Is there some definitive reference that will shed some light on this, in detail?

EDIT: "GCC-based" includes mingw too, but not MSVC.
## [7][New videos from C++ Russia 2019 Piter](https://www.reddit.com/r/cpp/comments/f3s12v/new_videos_from_c_russia_2019_piter/)
- url: https://www.youtube.com/playlist?list=PLZN9ZGiWZoZp8_XKK1WMcAt2Dr9-QkfU6
---

## [8][Zero, one, two, Freddy's coming for you](https://www.reddit.com/r/cpp/comments/f3ro02/zero_one_two_freddys_coming_for_you/)
- url: https://isocpp.org/blog/2020/02/zero-one-two-freddys-coming-for-you
---

## [9][As many keywords in a function definition as possible](https://www.reddit.com/r/cpp/comments/f3h21j/as_many_keywords_in_a_function_definition_as/)
- url: https://www.reddit.com/r/cpp/comments/f3h21j/as_many_keywords_in_a_function_definition_as/
---
Please note: this is just for fun.

I would like to create a function definition with as many keywords or modifiers as possible, while ignoring the actual body, and this is the best I can come up with:

    virtual inline operator const volatile unsigned long long int &amp;&amp; () const &amp;&amp; noexcept final override try {} catch (...) {}

Could I still add something without resorting to infinite `decltype` loops and while still having each keyword bring meaning to the definition?

&amp;#x200B;

EDIT: This is the new longest one when not counting infinite loops:

    explicit virtual inline operator const volatile unsigned long long int* const volatile &amp;&amp; () const volatile &amp;&amp; noexcept final override try {} catch (...) {}

Also this still ignores conditional `noexcept` as that could also be considered an infinite loop

&amp;#x200B;

EDIT 2: The longest one without any technically redundant keywords or infinite loops:

    explicit inline operator const volatile unsigned long long* const volatile &amp;&amp; () const volatile &amp;&amp; noexcept final try {} catch (...) {}
## [10][CppCast: C++ on a Watch](https://www.reddit.com/r/cpp/comments/f3n15h/cppcast_c_on_a_watch/)
- url: https://cppcast.com/brad-larson-cpp-watch/
---

## [11][Herb Sutter: Quantifying Accidental Complexity: An Empirical Look at Teaching and Using C++](https://www.reddit.com/r/cpp/comments/f2xq0l/herb_sutter_quantifying_accidental_complexity_an/)
- url: https://www.youtube.com/watch?v=qx22oxlQmKc
---

