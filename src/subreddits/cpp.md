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
## [2][Increased Complexity of C++20 Range Algorithms Declarations - Is It Worth?](https://www.reddit.com/r/cpp/comments/j9ndx1/increased_complexity_of_c20_range_algorithms/)
- url: https://www.bfilipek.com/2020/10/complex-ranges-algorithms.html?m=1
---

## [3][A very short war story on too much overloading](https://www.reddit.com/r/cpp/comments/j9e9k7/a_very_short_war_story_on_too_much_overloading/)
- url: https://quuxplusone.github.io/blog/2020/10/11/overloading-considered-harmful/
---

## [4][Modern C++ Undo/Redo framework](https://www.reddit.com/r/cpp/comments/j9of7d/modern_c_undoredo_framework/)
- url: https://www.reddit.com/r/cpp/comments/j9of7d/modern_c_undoredo_framework/
---
Hello everyone, just sharing my latest development.  
[https://github.com/ryder052/History](https://github.com/ryder052/History)

Happy coding :)
## [5][std::i/ostream's are a bit too oversized for both usage and implementation. Are there lighter alternatives widely used?](https://www.reddit.com/r/cpp/comments/j9mqsu/stdiostreams_are_a_bit_too_oversized_for_both/)
- url: https://www.reddit.com/r/cpp/comments/j9mqsu/stdiostreams_are_a_bit_too_oversized_for_both/
---
Disclaimer: I'm not solving any specific production case.

Standard input/output stream specifications include a lot of details - locales, buffering, large set of virtual functions. And they're known for their performance implications. So I was wondering if there's some simpler version of i/o interface, coming from some 3rd party library, which is widely used. Something like this:

    struct reader {
        virtual std::size_t read(std::span&lt;std::byte&gt; buffer) = 0;
    }
    struct writer {
        virtual std::size_t write(std::span&lt;std::byte&gt; buffer) = 0;
        virtual void flush() = 0;
    }

Or does everyone either just use std::i/o stream's or reimplements concept above from scratch?

EDIT: I'd like to clarify that I'm not looking for alternative text formatting facilities. And I'm well aware about `std::format`. What I'm looking for is a widely used thin interface for reading bytes from some source or writing bytes to some sink.
## [6][bitflags v1.3.0 - raw flags support](https://www.reddit.com/r/cpp/comments/j9mh4y/bitflags_v130_raw_flags_support/)
- url: https://www.reddit.com/r/cpp/comments/j9mh4y/bitflags_v130_raw_flags_support/
---
Hi everybody,

New version of [bitflags](https://github.com/m-peko/bitflags) library is released these days!

Previously, you could create your flags only by writing something like:

    BEGIN_BITFLAGS(Flags)
        FLAG(none)
        FLAG(flag_a)
        FLAG(flag_b)
        FLAG(flag_c)
    END_BITFLAGS(Flags)

Each flag is consisted of raw value (bits) and flag's name. So you could get a string representation of each flag by writing: `Flags::flag_a.name` .

In the new version, you can create raw flags, i.e. flags without string representation. Purpose of this feature is to provide light version that would take up as little space as possible:

    BEGIN_RAW_BITFLAGS(RawFlags)
        RAW_FLAG(none)
        RAW_FLAG(flag_a)
        RAW_FLAG(flag_b)
    END_RAW_BITFLAGS(RawFlags)

What do you think of this feature?

**What's next?**

In the next version, support for C++11 and C++14 will be added.

Thank you all!
## [7][I made a vector that maintains polymorphism without being (I hope) too janky](https://www.reddit.com/r/cpp/comments/j9mcum/i_made_a_vector_that_maintains_polymorphism/)
- url: https://github.com/friedkeenan/poly_vector
---

## [8][Overloading by Return Type in C++](https://www.reddit.com/r/cpp/comments/j94jd8/overloading_by_return_type_in_c/)
- url: https://artificial-mind.net/blog/2020/10/10/return-type-overloading
---

## [9][Fegeya Scrift - Colorized &amp; Customizable CLI Shell Project for everyone. Written in C++17.](https://www.reddit.com/r/cpp/comments/j9nwvl/fegeya_scrift_colorized_customizable_cli_shell/)
- url: https://www.reddit.com/r/cpp/comments/j9nwvl/fegeya_scrift_colorized_customizable_cli_shell/
---
Hello everyone, This is Scrift, my CLI shell project for GNU/Linux and FreeBSD (maybe OpenBSD and NetBSD (for GNU C and Clang)).

This project's all features listed on GitHub (README).

[GitHub repository is here! Click me](https://github.com/ferhatgec/scrift)

[Example video is here, Click me! (YouTube)](https://www.youtube.com/watch?v=ao5OqL-smzs&amp;t=8s)
## [10][CppCon 2020 presentation recommendations?](https://www.reddit.com/r/cpp/comments/j9atb4/cppcon_2020_presentation_recommendations/)
- url: https://www.reddit.com/r/cpp/comments/j9atb4/cppcon_2020_presentation_recommendations/
---
I've seen four videos so far and I can only recommend one of them as being interesting and/or well presented: [The Many Shades of reference_wrapper - Zhihao Yuan - CppCon 2020](https://www.youtube.com/watch?v=EKJMZCL00Ak)

Do you guys have recommendations?
## [11][what's the actual C++20 features compilers support status?](https://www.reddit.com/r/cpp/comments/j9q25y/whats_the_actual_c20_features_compilers_support/)
- url: https://www.reddit.com/r/cpp/comments/j9q25y/whats_the_actual_c20_features_compilers_support/
---
according to my experience with the trunk nightly builds and vs preview:

* **concepts/range**: GCC has the most complete and highest quality implementation. MSVC doesn't support requires expression in other contexts. Clang still doesn't support requires-clause short-circuiting, and lacks concepts and ranges library features.

* **modules**: Clang and MSVC merged it in trunk, but feels buggier than GCC modules branch.

* **coroutine**: Clang has the highest quality implementation. GCC is almost on par with Clang, though the implementation started latest. MSVC is almost unusable.
