# cpp
## [1][C++ Jobs - Q3 2020](https://www.reddit.com/r/cpp/comments/hjnaf2/c_jobs_q3_2020/)
- url: https://www.reddit.com/r/cpp/comments/hjnaf2/c_jobs_q3_2020/
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

* [C++ Jobs - Q2 2020](https://www.reddit.com/r/cpp/comments/ft77lv/c_jobs_q2_2020/)
## [2][bitflags - C++17 single-header library for easily managing set of flags](https://www.reddit.com/r/cpp/comments/im9opm/bitflags_c17_singleheader_library_for_easily/)
- url: https://github.com/m-peko/bitflags
---

## [3][Almost 30 year old optimization bug in MSVC causing incorrect code (In specific circumstances) found](https://www.reddit.com/r/cpp/comments/ilzbqd/almost_30_year_old_optimization_bug_in_msvc/)
- url: https://developercommunity.visualstudio.com/content/problem/1096974/msvc-global-optimization-of-program-results-in-inc.html?childToView=1172202#comment-1172202
---

## [4][Meeting C++ online: Ivan Čukić - Move-only types can save the API](https://www.reddit.com/r/cpp/comments/imcw6n/meeting_c_online_ivan_čukić_moveonly_types_can/)
- url: https://www.youtube.com/watch?v=ojTsSGlhVLo
---

## [5][It turns out `std::tuple` is not a zero-cost abstraction](https://www.reddit.com/r/cpp/comments/ilujab/it_turns_out_stdtuple_is_not_a_zerocost/)
- url: https://www.reddit.com/r/cpp/comments/ilujab/it_turns_out_stdtuple_is_not_a_zerocost/
---
I used tuples for a long time now in C++, always assuming they're just a struct generator and thus have no cost. Surely this:

```
struct two_ints { int a; int b };
```

and this:

```
std::tuple&lt;int, int&gt;
```

will compile to the same assembly. Well, turns out this is not true. I stumbled across this SO question:

https://stackoverflow.com/questions/63719249/why-stdtuple-breaks-small-size-struct-call-convention-optimization-in-c

Serves me right, as I didn't follow the general "don't assume, measure" advice. So:

    foo(std::tuple&lt;int, int&gt;);
    
    movabs  rax, 4294967298
    sub     rsp, 24
    lea     rdi, [rsp+8]
    mov     QWORD PTR [rsp+8], rax
    call    func_tuple(std::tuple&lt;int, int&gt;)
    add     rsp, 24
    ret

is more expensive to call than:

    foo(two_ints);
    
    movabs  rdi, 8589934593
    jmp     func_struct(two_ints)

([Godbolt link](https://godbolt.org/z/nrTors))

Except on clang with libc++, where `tuple` produces the same code as a struct.

I've never seen this tuple overhead mentioned anywhere. Is this a big surprise to anyone else, or just me?
## [6][CppCast: Unit Testing](https://www.reddit.com/r/cpp/comments/im8d2s/cppcast_unit_testing/)
- url: https://cppcast.com/testing-oleg-rabaev/
---

## [7][Preprocessor Embed and Language Embed - The Last Sprint](https://www.reddit.com/r/cpp/comments/ilomtk/preprocessor_embed_and_language_embed_the_last/)
- url: https://thephd.github.io/preprocessor-embed-std-embed-the-last-spring
---

## [8][Concept archetypes | Andrzej's C++ blog](https://www.reddit.com/r/cpp/comments/ilp8sx/concept_archetypes_andrzejs_c_blog/)
- url: https://akrzemi1.wordpress.com/2020/09/02/concept-archetypes/
---

## [9][Using Vim for C++ development](https://www.reddit.com/r/cpp/comments/ils6ob/using_vim_for_c_development/)
- url: https://gist.github.com/p1v0t/42a34744b5e4f5980e5f4e1c980ec859
---

## [10][glyph v0.1.0 released for open-source](https://www.reddit.com/r/cpp/comments/im3ffa/glyph_v010_released_for_opensource/)
- url: https://www.reddit.com/r/cpp/comments/im3ffa/glyph_v010_released_for_opensource/
---
Hey everyone! I've been working on a command line program that would be used for encryption purposes. It's in very early development, so I thought it would be a good time to make it open source. Documentation is available in the README, and it's licensed with the GNU GPLv3 license. Anyone can send a pull request, or an issue if you find anything. If you have any questions, feel free to contact me! [Here's the link.](https://github.com/levi02321/glyph)

Edit: In similar fashion to my git commits, I forgot the main feature! I've included the link at the end of the post in case anyone wants to check out the project on GitHub.
## [11][What's the most useless keyword in c++?](https://www.reddit.com/r/cpp/comments/ilq2f0/whats_the_most_useless_keyword_in_c/)
- url: https://www.reddit.com/r/cpp/comments/ilq2f0/whats_the_most_useless_keyword_in_c/
---
is it `register` ?
