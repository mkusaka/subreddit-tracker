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
## [2][Dragonbox: yet another float-to-string conversion algorithm](https://www.reddit.com/r/cpp/comments/ika6ml/dragonbox_yet_another_floattostring_conversion/)
- url: https://www.reddit.com/r/cpp/comments/ika6ml/dragonbox_yet_another_floattostring_conversion/
---
Hi, I'm the author of Grisu-Exact ([https://github.com/jk-jeon/Grisu-Exact](https://github.com/jk-jeon/Grisu-Exact)). I developed a new algorithm, *Dragonbox*, which performs better than Grisu-Exact, still at the same time have all of the followings:

1. Roundtrip guarantee
2. Shortest length guarantee
3. Correct rounding guarantee

Please have a look at the repository if you are interested: [https://github.com/jk-jeon/dragonbox](https://github.com/jk-jeon/dragonbox)

(To be fair, I could perhaps apply some of the optimization ideas I got from implementing Dragonbox to Grisu-Exact and save maybe about `1-2ns`, but I nevertheless believe Dragonbox is a better algorithm.)

&amp;#x200B;

Here are some backgrounds on how I came up with it.

Recently I realized there is a faster algorithm called *Schubfach* ([https://drive.google.com/file/d/1luHhyQF9zKlM8yJ1nebU0OgVYhfC6CBN/view](https://drive.google.com/file/d/1luHhyQF9zKlM8yJ1nebU0OgVYhfC6CBN/view); a C++ implementation is here: [https://github.com/abolz/Drachennest](https://github.com/abolz/Drachennest)) developed by  Raffaello Giulietti in (maybe around?) 2017-2018. For some reason I don't know about, this algorithm didn't get much attention compared to Ryu, but AFAICT, Schubfach is wonderful. It is really simple, elegant, and fast!

It was sad but I had to admit that Schubfach is a better algorithm than Grisu-Exact, so I decided to look at it more deeply, and eventually, came up with Dragonbox.

The beauty of Schubfach is that it does not perform any iterative search for the shortest representation. Rather, it just magically finds the "optimal number" at only one trial. But there are two downsides:

1. The output might contain trailing decimal zeros. For example, for the input `1.234f` the output might be something like `123400000 * 10^-8`.
2. It always computes all of the left endpoint, the center, and the right endpoint of the relevant search interval, which means there are three 128bit x 64bit -&gt; high 64bit multiplications which are expensive. (This is for the case of `double`; for `float`, we instead have 64bit x 32bit -&gt; high 32bit multiplications, but this is still expensive.)

However, by applying some Grisu-like method, it is possible to reduce the amortized number of 128bit x 64bit -&gt; high 64bit  multiplications into one, at the cost of having some branches and divisions-by-constants. This is the core idea of Dragonbox.

&amp;#x200B;

And about trailing decimal zeros. Like Schubfach, the output of Dragonbox might contain trailing decimal zeros. The implementation provides an option to specify what to do with that: to remove them completely, to report if there might be trailing decimal zeros, or to just don't care about it.

Not removing trailing zeros actually can make sense, because for the string generation procedure, we still need to divide by powers of 10 and compute both the quotient and the remainder. Hence, it might be possible to delay the removal of unnecessary trailing zeros to the string generation procedure and get a better performance. However, I didn't have enough experiment on that and I just removed all trailing zeros for the benchmark.

Thanks!
## [3][The problem with C](https://www.reddit.com/r/cpp/comments/ik3o10/the_problem_with_c/)
- url: https://cor3ntin.github.io/posts/c/index.html
---

## [4][C++17 Zero allocation Coroutine/[Resumable function] library](https://www.reddit.com/r/cpp/comments/ikhlvw/c17_zero_allocation_coroutineresumable_function/)
- url: https://www.reddit.com/r/cpp/comments/ikhlvw/c17_zero_allocation_coroutineresumable_function/
---
I'm releasing a stack only coroutine library. Its still got plenty of work to make it feature complete.

[https://github.com/seppeon/SCoro](https://github.com/seppeon/SCoro)  


Why, because C++20 coroutines:

* are extremely difficult to develop with.
* make nested co\_await very clumbsy and have large memory overhead to do it right.
* HALO optimization seems to regularly fail.
* require heap allocation to be avaliable in the general case, making them unusable on many systems.

Any feedback is welcome, I have a long todo here. I'm looking at improving the actual standard C++ coroutines eventually with a similar approach. If anybody sees advantages and disadvantages of my approach I'd love to hear about them!

The ultimate idea is to have the language be syntactic sugar around this sort of pattern.
## [5][libfev - a small library for events and fibers](https://www.reddit.com/r/cpp/comments/ik8c1p/libfev_a_small_library_for_events_and_fibers/)
- url: https://www.reddit.com/r/cpp/comments/ik8c1p/libfev_a_small_library_for_events_and_fibers/
---
Hello,

I implemented a small [library](https://github.com/patrykstefanski/libfev) for fibers and nonblocking I/O, so that it is possible to write programs in a simple blocking style. Currently, it provides:

* Few multithreaded schedulers
* Backends for epoll and kqueue (and experimental io_uring backend)
* Basic socket operations
* Timers
* Synchronization primitives (mutex, condition variable and semaphore)

I did some [benchmarks](https://github.com/patrykstefanski/async-bench) and in a throughput benchmark libfev can handle up to 172% more requests per second than Boost.Asio, up to 77% more than Tokio, up to 40% more than async-std and up to 16% more than Go.
However, I am not an expert in these libraries, so I might have implemented something in a suboptimal way. Any feedback is welcome.
## [6][Online C++ User Group meetings in September 2020](https://www.reddit.com/r/cpp/comments/ikhuy7/online_c_user_group_meetings_in_september_2020/)
- url: https://www.meetingcpp.com/meetingcpp/news/items/Online-Cpp-User-Group-meetings-in-September-2020.html
---

## [7][Freestanding in my home office](https://www.reddit.com/r/cpp/comments/ik40qx/freestanding_in_my_home_office/)
- url: https://www.reddit.com/r/cpp/comments/ik40qx/freestanding_in_my_home_office/
---
[Prague trip report](https://www.reddit.com/r/cpp/comments/f5hgqm/freestanding_in_prague/)

As it turns out, you don't actually need to go on a trip to write a trip report.  Surprise!

Over the last six months, Library Evolution (and other groups) have been holding regular telecons to continue progress on C++.  Telecons have worked out well for me personally, as I am typically only able to travel to two of the three face-to-face meetings a year.  This has let me contribute more officially during the spring-summer time period, as I wasn't going to be able to attend the Varna meeting that was originally scheduled.  That's not to say that telecons are great for everyone, as time zones, work commitments, preparation time, and general burnout have been issues for many.  Even though there aren't any scheduling conflicts, I have skipped almost all Language Evolution meetings, and most Library meetings.  Core meetings are right out.  Progress on the paper backlog has also been much slower compared to what we get done during face-to-face meetings.

C++20 was finished in Prague (Feb 2020), so all of this work is going towards C++23.

# Executors

As one of the newly minted LEWG Vice Chairs, I get to participate in paper scheduling decisions.  For C++23, we are following [The Plan](https://wg21.link/p0592).  That means paper scheduling will prioritize:

* [Modularizing the standard library](https://wg21.link/P0581) (`import std;`)
* [Executors](https://wg21.link/P0443) (a framework for dealing with concurrency and parallelism)
* [Library support for coroutines](https://wg21.link/P1056) (a way of expressing concurrency)
* [Networking](https://wg21.link/N4734) (coupled to concurrency and parallelism framework)

This means that an awful lot of the prioritized features for C++23 rely on making Executors successful.  The executors paper is big by wg21 standards.  60 printed pages, 11 authors, [8 years](https://wg21.link/n3378) of history, and lots of home-grown frameworks playing in the same space.

For most of the last eight years, executors have been discussed in the concurrency study group.  There is substantial overlap in membership between the concurrency study group and Library Evolution, but the set of concerns of the groups is different.  In order for Library Evolution to effectively review and understand executors, the Library Evolution chair (/u/blelbach) chopped the paper up into smaller pieces and sent those pieces to smaller groups to review and discuss them.  Thread pools goes to one group, the "property" mechanism goes to another, and so on until you end up with six groups, meeting for 4-10 weeks.  It's been an enormous amount of effort poured into reviewing the paper, and you can see some of the fruits of those efforts [in](https://wg21.link/P2202) [the](https://wg21.link/p2203) [various](https://wg21.link/p2205) [reports](https://wg21.link/p2207) that are in the [mailing](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/).  The executors authors now have a lot of feedback to incorporate into the next paper revision.

# Freestanding

Freestanding has made some progress in the last six months.

[P2013: Freestanding Language: Optional ::operator new](https://wg21.link/p2013) is a feature test macro decision away from moving to Core.  The bad news here is that many of the smaller study groups aren't meeting right now, so getting feedback on whether a feature test macro is required for this feature may take some time.

[P1641: Freestanding Library: Rewording the Status Quo](https://wg21.link/p1641) and [P1642: Freestanding Library: Easy \[utilities\], \[ranges\], and \[iterators\]](https://wg21.link/p1642r3) were seen and viewed favorably.  Library Evolution doesn't want me to walk through each of the inclusions / exclusions listed in those papers again, which is progress.  Yet again, the feature test macro approach proved somewhat controversial.  [Library Evolution has thus guided me to reorganize these papers](https://wg21.link/p1642/github).  You can see the reorganization in a new revision of [P1642](https://wg21.link/p1642) that folds in the non-feature-test macro parts of P1641.  The feature-test macro parts are in [P2198: Freestanding Feature-Test Macros and Implementation-Defined Extensions](https://wg21.link/P2198).

# Feature-test macros

When I started doing freestanding work in 2017 with [P0829](https://wg21.link/p0829), feature test macros were just something in a standing document, and weren't in the standard.  Now in 2020, I'm pretty sure that freestanding papers have consumed more than half of all the feature-test macro study group's time.  With regards to freestanding, the feature test macros are in a bit of a mess.  I feel that, in 2030, it will be nice having all the feature test macro stuff worked out, but in 2020, it certainly is painful.

Here's an optimistic answer to a reasonable question that should be answerable with feature test macros.

*Is the non-parallel version of uninitialized\_default\_construct available?*

    #if defined(__cpp_lib_raw_memory_algorithms) &amp;&amp; __cpp_lib_raw_memory_algorithms &gt;= 201606L
    // hosted success path and freestanding extension path
    #elif defined(__cpp_lib_freestanding_memory) &amp;&amp; __cpp_lib_freestanding_memory &gt;= 202007L
    // freestanding path
    #else
    // fallback path
    #endif

I say this is optimistic, because it is also assuming that `&lt;version&gt;` is available.  For pre-C++20 compilers, that isn't a guarantee either, and that detection code isn't trivial either.

There are much sadder questions that can be asked in freestanding contexts though...

*Is thread available?*

`// No good answers here. Currently calling this out of scope.`

Getting feature testing right is surprisingly tricky and error prone, especially since the whole point of the feature is to be usable in codebases that need to work on old toolchains.  A time machine would be really handy here.

# What's next for freestanding?

The most important aspect of standard C++ affecting freestanding is the lack of suitable error handling. This is also the area that is the most difficult to gain any kind of consensus.  Keep an eye out for a paper in the next few months that looks to discuss the many trade-offs in the various standard and standard-track error handling approaches.

Other than that, there are a lot of papers that need to be written to cover the aspects of [P0829](https://wg21.link/P0829) that haven't been covered in [P1642](https://wg21.link/P1642).  There are questions in these future papers that need to be resolved by Library Evolution, like "are we willing to split overload sets that look like `std::abs`", "are we willing to have partial class support for classes like `std::array`", and "are we willing to require more C support for freestanding C++ than C requires from freestanding C".  These should be straightforward papers, but they would all be competing with a very clogged back log of LEWG papers.

Language Evolution isn't quite as pressed for time, so now would be a good time to get some freestanding features in here.  I have new ideas on how to tackle issues involving `thread_local`, as well as constructors and destructors called during startup and termination, but those ideas would need some implementation experience that I haven't found the time for.

The area that most needs attention (that I don't have time to provide) is compile time programming for freestanding environments.  `std::vector` is a poor candidate for runtime kernel and embedded programming, but it *should* be just fine for compile-time kernel and embedded programming.  What we need is for someone to see what happens when you take all the non-freestanding `constexpr`'s in the standard library, and make them `consteval`.  Do bad things happen?  Do we end up splitting overload sets in nasty ways?  Does the change in object destruction timing frequently cause bugs?  I don't know, and would only really know the answer to these questions with some implementation and usage experience.

On top of all that, the Microsoft STL implementation now has much of the test infrastructure open sourced as well.  I need to get around to contributing a kernel test harness that can exercise freestanding C++ in the Windows kernel.  So much to do, and so little time.
## [8][EnTT v3.5.0 is out: Gaming meets Modern C++](https://www.reddit.com/r/cpp/comments/ik7tl2/entt_v350_is_out_gaming_meets_modern_c/)
- url: /r/gamedev/comments/ik7sk1/entt_v350_is_out_gaming_meets_modern_c/
---

## [9][Standard Library include times (measured in MSVC)](https://www.reddit.com/r/cpp/comments/ik5inf/standard_library_include_times_measured_in_msvc/)
- url: https://github.com/s9w/cppslit
---

## [10][CppCon 2020 Keynote: C++20 by Marc Gregoire](https://www.reddit.com/r/cpp/comments/ikb57w/cppcon_2020_keynote_c20_by_marc_gregoire/)
- url: https://www.reddit.com/r/cpp/comments/ikb57w/cppcon_2020_keynote_c20_by_marc_gregoire/
---
A new announcement from CppCon 2020. **Marc Gregoire** keynote talk titled: "*C++20: An (Almost) Complete Overview*."

[https://cppcon.org/cppcon-2020-keynote-c20-by-marc-gregoire/](https://cppcon.org/cppcon-2020-keynote-c20-by-marc-gregoire/)
## [11][spdlog 1.8.0 released](https://www.reddit.com/r/cpp/comments/ikk819/spdlog_180_released/)
- url: https://github.com/gabime/spdlog/releases/tag/v1.8.0
---

