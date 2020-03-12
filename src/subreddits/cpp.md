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
## [2][Fold Expressions with cout - doubt](https://www.reddit.com/r/cpp/comments/fhfxwr/fold_expressions_with_cout_doubt/)
- url: https://www.reddit.com/r/cpp/comments/fhfxwr/fold_expressions_with_cout_doubt/
---
Why the the atempt to print with a break line using binary fold doesn't work but the second option with unary case it works?

    template&lt;typename... T&gt;
    void print(T&amp;&amp;... args)
    {
            // Binary fold
    	(cout &lt;&lt; ... &lt;&lt; (args &lt;&lt; '\n')) &lt;&lt; '\n'; // doesn't work
             
            // Unary fold
    	((cout &lt;&lt; args &lt;&lt; '\n'), ...); // prints with break line - Works fine!
    }
## [3][Question: Websocket in STL](https://www.reddit.com/r/cpp/comments/fh1vfb/question_websocket_in_stl/)
- url: https://www.reddit.com/r/cpp/comments/fh1vfb/question_websocket_in_stl/
---
I've found a number of libraries that do websockets in C++

* [https://docs.websocketpp.org/index.html](https://docs.websocketpp.org/index.html)
* [https://www.boost.org/doc/libs/1\_72\_0/libs/beast/doc/html/index.html](https://www.boost.org/doc/libs/1_72_0/libs/beast/doc/html/index.html)
* [https://libwebsockets.org](https://libwebsockets.org)

I was wondering what people used for this/is there an STL way of doing WebSockets easily?

The library needs to be permissive (Apache, MIT, BSD) because it's for an open source project.

Edit: I'm consuming a WebSocket API, not creating one.
## [4][Why there is no wchar_t variant of to_chars and from_chars?](https://www.reddit.com/r/cpp/comments/fh5h8q/why_there_is_no_wchar_t_variant_of_to_chars_and/)
- url: https://www.reddit.com/r/cpp/comments/fh5h8q/why_there_is_no_wchar_t_variant_of_to_chars_and/
---
It is especially painful on Windows where most text is `wchar_t`. `to_chars` and `from_chars` operate on `char*` buffers. But what is `char` anyway? It could be binary blob or it could be text. If it is text, what encoding is it in? It could be ASCII, UTF-8, ANSI (Windows codepage such as cp1250 or Shift-JIS), EBCDIC or something completely different. On some platforms, `char` is 16bits wide. This is bad for inter-compatibility. We already have `char` / `wchar_t` / `char8_t` / `char16_t` / `char32_t` types in Standard and also `std::string` is using them. Is there a reason why `to_chars` and `from_chars` didn't receive a bit of Unicode love? Is there a paper where I can read the rationale why this feature was not included?
## [5][C++ Benchmark: Timsort vs pdqsort vs quadsort vs std::stable_sort](https://www.reddit.com/r/cpp/comments/fgxfqa/c_benchmark_timsort_vs_pdqsort_vs_quadsort_vs/)
- url: https://www.reddit.com/r/cpp/comments/fgxfqa/c_benchmark_timsort_vs_pdqsort_vs_quadsort_vs/
---
[https://rextester.com/QDXH20310](https://rextester.com/QDXH20310)

Linked above is the benchmark source code, I set int max = 512 because it only allows 10 seconds of running time.

My main question for those interested at looking at the source code, is this a fair c++ benchmark or do any of the specific implementations have an unfair advantage? Also, any idea why quadsort is significantly faster on small arrays?

Output with max set to 1000000 and compiled with g++ -O3 -fpermissive source.cpp. Each result is the best run out of 100.

             quadsort: sorted 1000000 i32s in 0.070989 seconds. MO: 0 (random order)
            std::sort: sorted 1000000 i32s in 0.064780 seconds. MO: 0 (random order)
              timsort: sorted 1000000 i32s in 0.088847 seconds. MO: 0 (random order)
              pdqsort: sorted 1000000 i32s in 0.030218 seconds. MO: 0 (random order)
    
             quadsort: sorted 1000000 i32s in 0.000856 seconds. MO: 0 (ascending order)
            std::sort: sorted 1000000 i32s in 0.011282 seconds. MO: 0 (ascending order)
              timsort: sorted 1000000 i32s in 0.000365 seconds. MO: 0 (ascending order)
              pdqsort: sorted 1000000 i32s in 0.000833 seconds. MO: 0 (ascending order)
    
             quadsort: sorted 1000000 i32s in 0.000847 seconds. MO: 0 (uniform order)
            std::sort: sorted 1000000 i32s in 0.007948 seconds. MO: 0 (uniform order)
              timsort: sorted 1000000 i32s in 0.000361 seconds. MO: 0 (uniform order)
              pdqsort: sorted 1000000 i32s in 0.000903 seconds. MO: 0 (uniform order)
    
             quadsort: sorted 1000000 i32s in 0.003561 seconds. MO: 0 (descending order)
            std::sort: sorted 1000000 i32s in 0.008633 seconds. MO: 0 (descending order)
              timsort: sorted 1000000 i32s in 0.000782 seconds. MO: 0 (descending order)
              pdqsort: sorted 1000000 i32s in 0.001882 seconds. MO: 0 (descending order)
    
             quadsort: sorted 1000000 i32s in 0.017593 seconds. MO: 0 (random tail)
            std::sort: sorted 1000000 i32s in 0.026802 seconds. MO: 0 (random tail)
              timsort: sorted 1000000 i32s in 0.020676 seconds. MO: 0 (random tail)
              pdqsort: sorted 1000000 i32s in 0.021670 seconds. MO: 0 (random tail)
    
             quadsort: sorted 1000000 i32s in 0.010657 seconds. MO: 0 (wave order)
            std::sort: sorted 1000000 i32s in 0.026747 seconds. MO: 0 (wave order)
              timsort: sorted 1000000 i32s in 0.008884 seconds. MO: 0 (wave order)
              pdqsort: sorted 1000000 i32s in 0.024792 seconds. MO: 0 (wave order)
    
             quadsort: sorted    1000 i32s in 0.005088 seconds. KO: 0 (random range)
            std::sort: sorted    1000 i32s in 0.015931 seconds. KO: 0 (random range)
              timsort: sorted    1000 i32s in 0.014487 seconds. KO: 0 (random range)
              pdqsort: sorted    1000 i32s in 0.010379 seconds. KO: 0 (random range)

Summary:

quadsort wins 2 out of 7, timsort wins 4 out of 7, pdqsort wins 1 out of 7

quadsort is good at semi-ordered data and small arrays.

timsort is good at ordered / reverse ordered data.

pdqsort is good at random data and large arrays. (not a stable sort)

std::sort beats timsort and quadsort at random data. (not a stable sort)

The final test measures the average performance on 1000 arrays ranging from size 1 to 999 filled with random data.

quadsort source and description:

[https://github.com/scandum/quadsort](https://github.com/scandum/quadsort)

timsort source and description:

[https://github.com/timsort/cpp-TimSort](https://github.com/timsort/cpp-TimSort)

pdqsort source and description:

[https://github.com/orlp/pdqsort](https://github.com/orlp/pdqsort)
## [6][PythonFMU or how to call Python from C++](https://www.reddit.com/r/cpp/comments/fgthet/pythonfmu_or_how_to_call_python_from_c/)
- url: https://www.reddit.com/r/cpp/comments/fgthet/pythonfmu_or_how_to_call_python_from_c/
---
Wanted to share a library I co-wrote for exporting co-simulation models written in Python that is compatible with FMI 2.0.  [https://github.com/NTNU-IHB/PythonFMU](https://github.com/NTNU-IHB/PythonFMU)

The reason I share it in the cpp sub is that it is a nice example of how Python can be embedded from C++. See  [https://github.com/NTNU-IHB/PythonFMU/tree/master/pythonfmu/pythonfmu-export](https://github.com/NTNU-IHB/PythonFMU/tree/master/pythonfmu/pythonfmu-export) and especially  [https://github.com/NTNU-IHB/PythonFMU/blob/master/pythonfmu/pythonfmu-export/cpp/PySlaveInstance.cpp](https://github.com/NTNU-IHB/PythonFMU/blob/master/pythonfmu/pythonfmu-export/cpp/PySlaveInstance.cpp)
## [7][Moving from C to C++](https://www.reddit.com/r/cpp/comments/fhe4oi/moving_from_c_to_c/)
- url: https://www.reddit.com/r/cpp/comments/fhe4oi/moving_from_c_to_c/
---
Hello everybody,  
I'm a university student and I've been working with C for the past 2 years. At this point, I think I have a solid knowledge of C, but nothing super advanced. I've done some basic programming and data structures, also many different functions and algorithms. Right now I want to switch to C++, because I'll have to use it in the next semester, but I want to do some preparation before that. 

My question for you is: How to I make a transition to C++? Are there any good books or tutorials you can recommend? Do you have any tips?

Thanks guys!
## [8][I wrote a 3D math library for games](https://www.reddit.com/r/cpp/comments/fgmj30/i_wrote_a_3d_math_library_for_games/)
- url: https://www.reddit.com/r/cpp/comments/fgmj30/i_wrote_a_3d_math_library_for_games/
---
Hi everyone,

A while ago, I was looking for a math lib for game dev that supports my math conventions, all the features I need, and has a good syntax. I had problems with all of the existing libs, so I made my own. I also tried to solve other peoples' similar problems by making the library highly configurable with templates, so hopefully it supports your conventions and features as well.

I'm sharing this library in the hopes that it will be useful for your project. If you want to get into template meta-programming and modern C++, the source code may also be interesting to browse.

Let me know what you think: [https://github.com/petiaccja/Mathter](https://github.com/petiaccja/Mathter)
## [9][Fast float parsing in practice](https://www.reddit.com/r/cpp/comments/fgcufa/fast_float_parsing_in_practice/)
- url: https://lemire.me/blog/2020/03/10/fast-float-parsing-in-practice/
---

## [10][&lt;cuchar&gt; mbrtoc8 c8rtomb](https://www.reddit.com/r/cpp/comments/fgy4da/cuchar_mbrtoc8_c8rtomb/)
- url: https://www.reddit.com/r/cpp/comments/fgy4da/cuchar_mbrtoc8_c8rtomb/
---
When will the [full implementation](https://en.cppreference.com/w/cpp/header/cuchar) of \`&lt;cuchar&gt;\` be delivered by all three?

By all three you know who I mean. I am looking at you: clang, gcc and msvc  :)
## [11][How would passing non-trivial types in registers preserve semantics?](https://www.reddit.com/r/cpp/comments/fgdcs9/how_would_passing_nontrivial_types_in_registers/)
- url: https://www.reddit.com/r/cpp/comments/fgdcs9/how_would_passing_nontrivial_types_in_registers/
---
Many recent posts in this subreddit about breaking ABI compatibility mention the ability to pass `unique_ptr` in a register as one of the potential benefits. While I agree this would be beneficial, I wonder how it can be done while preserving the following fundamental properties of objects:
- Construction / destruction order
- Object identity

I was taught that, in general, objects are destroyed in the opposite order of creation. This applies to class members, bases, function locals, and most crucially to temporaries.

In the expression:
   
    f(a(), b(), c());
   
The calls to `a()`, `b()`, and `c()` are indeterminately sequenced, but the temporary values returned live until the call to `f` returns, and then are destroyed in the opposite order. If one of those temporaries is a non-trivial object and is passed through a register, will this continue to be true?

I was also taught that, in general, the address of an object uniquely identifies that object (modulo base classes, first members, or objects that provide storage for other objects). In particular, if for some non-trivial type a constructor runs with `this` = some address, I can expect the destructor to eventually run with the same `this`. If we allow passing non-trivial objects through registers, how will this property be preserved?

I'm aware of one situation where the above is not the case - the proposed `[[trivially_relocatable]]` attribute. `unique_ptr` and even `shared_ptr` would benefit greatly from it, and its use "answers" my concerns. But will the imagined future ABI that allows passing non-trivial types in registers only allow it for types that opt in to this attribute?
