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
## [2][ReSharper C++ 2020.2 is out with updates to Unreal Engine integration, improved C++/CLI support, and more](https://www.reddit.com/r/cpp/comments/iflc9j/resharper_c_20202_is_out_with_updates_to_unreal/)
- url: https://blog.jetbrains.com/rscpp/2020/08/13/resharper-cpp-2020-2/
---

## [3][CppCon 2020](https://www.reddit.com/r/cpp/comments/iffap0/cppcon_2020/)
- url: https://www.reddit.com/r/cpp/comments/iffap0/cppcon_2020/
---
Please join us this year online for Cppcon starting September 13th.
## [4][Any papers submitted for even less creation of temporaries in C++23/26?](https://www.reddit.com/r/cpp/comments/ifd1pv/any_papers_submitted_for_even_less_creation_of/)
- url: https://www.reddit.com/r/cpp/comments/ifd1pv/any_papers_submitted_for_even_less_creation_of/
---
We've come a long way with [http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2016/p0135r1.html](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2016/p0135r1.html) \- prvalues can pretty much stay in what we could call "ghost form"  until a variety of conditions forces materialization (the standard's term) into an actual temporary variable. Once you have a temporary though, you start paying for construction and destruction, and copies or moves.

One would hope that work continues to eliminate creation of temporaries when they are just passed along to initialize objects somewhere down the line.

The following scenarios would benefit the most from being optimized but I can't even find a discussion about those, much less a paper:

    struct Bar {}
    struct Foo {
        Foo(Bar&amp;&amp; bar) : m_bar(std::move(bar)) {}
        Bar m_bar;
    };
    
    void move3(Bar&amp;&amp; bar) { Bar myBar = std::move(bar); }
    void move2(Bar&amp;&amp; bar) { move3(std::move(bar)); }
    void move1(Bar&amp;&amp; bar) { move2(std::move(bar)); }
    
    void copy3(Bar bar) { Bar myBar = bar; }
    void copy2(Bar bar) { copy3(bar); }
    void copy1(Bar bar) { copy2(bar); }
    
    void Test()
    {
        //Counter-example: prvalue used to initialize b directly, no copy or move. We want more of this!
        Bar b = Bar{Bar{Bar{}}};
        
        //Ex 1: uses move constructor, instead of using "Bar{}" prvalue to initialize m_bar:
        Foo{Bar{}};
        
        //Ex 2: uses move constructor, instead of using "Bar{}" prvalue to initialize myBar:
        //At least it does only one move and elides the the extra ones
        move1(Bar{});
        
        //Ex 3: does 3 copies instead of eliding copies and using "Bar{}" prvalue to initialize myBar
        copy1(Bar{});
    }

*Equivalent code which logs constructor accesses, tested with gcc and clang on GodBolt:* [*https://godbolt.org/z/GfEYdd*](https://godbolt.org/z/GfEYdd)

 Are you guys aware of any work in progress to optimize those scenarios (paper links appreciated)? Are there specific obstacles?
## [5][Standard library development made easy with C++20](https://www.reddit.com/r/cpp/comments/ifotxd/standard_library_development_made_easy_with_c20/)
- url: https://cor3ntin.github.io/posts/tuple/
---

## [6][vector&lt;const T&gt; in C++20](https://www.reddit.com/r/cpp/comments/ifkasg/vectorconst_t_in_c20/)
- url: https://www.reddit.com/r/cpp/comments/ifkasg/vectorconst_t_in_c20/
---
/u/HowardHinnant notes in [this StackOverflow answer](https://stackoverflow.com/a/39652132) that the reason `vector&lt;const T&gt;` is illegal is due to the deprecated address overload on allocator. These overloads are now removed in C++20. Does this mean `vector&lt;const T&gt;` is now legal? Or does the allocator requirements still forbid this?
## [7][Can anyone recommend a good audio library for a game engine?](https://www.reddit.com/r/cpp/comments/ifowbe/can_anyone_recommend_a_good_audio_library_for_a/)
- url: https://www.reddit.com/r/cpp/comments/ifowbe/can_anyone_recommend_a_good_audio_library_for_a/
---
I've built my game so far just using SDL2 for rendering and audio - but moved rendering over to Vulkan and now want to upgrade the engine's audio library as well. Can anyone recommend a good audio library?

SDL2 audio is not great for many simultaneous sounds or sounds moving in 2D/3D coordinates - so that is the main feature I am looking to improve. I have read people recommend OpenAL though that library seems like it hasn't been used for many recent games according to their website. I also hear that irrklang is meant to be OK as well. I was also looking into [YSE](http://www.attr-x.net/yse/) but not sure how good it is.

Any help would be appreciated, thank you.
## [8][Glob for C++17: Unix-style pathname pattern expansion](https://www.reddit.com/r/cpp/comments/ife5w7/glob_for_c17_unixstyle_pathname_pattern_expansion/)
- url: https://github.com/p-ranav/glob
---

## [9][Purely online conferences](https://www.reddit.com/r/cpp/comments/ifkhge/purely_online_conferences/)
- url: https://www.reddit.com/r/cpp/comments/ifkhge/purely_online_conferences/
---
Upcoming conferences such as Meeting C++ and CppCon are held mostly or entirely online due to travel restrictions caused by the Corona virus. 

Aren't all of the talks uploaded on YouTube after the conference for free anyway? If so, what is the point of attending the online CppCon for $200? Genuinely wondering if I am missing something.
## [10][Compilation overhead of invoke](https://www.reddit.com/r/cpp/comments/ifnvo4/compilation_overhead_of_invoke/)
- url: https://eggs-cpp.github.io/invoke/benchmark.html
---

## [11][Should I buy the C++ Primer (5th edition) ?](https://www.reddit.com/r/cpp/comments/if5tzh/should_i_buy_the_c_primer_5th_edition/)
- url: https://www.reddit.com/r/cpp/comments/if5tzh/should_i_buy_the_c_primer_5th_edition/
---
Hello ! I'm new to coding and I want to learn C++. I searched on the net about some books and found out that C++ Primer (from lippman) is a great book. My only problem is that it's 2020 and I don't know if the book is still " up to date " or if it's a bit outdated ( I don't know when c++ get updated ). I hope people can help me ! Oh and if you have other suggestions please say it. Thanks !
