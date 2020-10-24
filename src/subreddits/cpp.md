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
## [2][Feedback on OpenGL Walkthrough &amp; Design](https://www.reddit.com/r/cpp/comments/jh0fmp/feedback_on_opengl_walkthrough_design/)
- url: https://www.reddit.com/r/cpp/comments/jh0fmp/feedback_on_opengl_walkthrough_design/
---
I was thinking about doing a weekly or semi weekly series were I detail my progress through OpenGL using modern C++ features and design.

I wanted to ask the community if they would be interested in such a series where I incorporate their feedback. Things ranging from C++ optimizations, code layout, design patterns, file structure, and of course graphics. 

If the community would be interested in such a concept, I would definitely love to try it and have you critique it.
## [3][std::once_flag is a glass hill](https://www.reddit.com/r/cpp/comments/jh7ogb/stdonce_flag_is_a_glass_hill/)
- url: https://quuxplusone.github.io/blog/2020/10/23/once-flag/
---

## [4][Version 0.6.0 of the simdjson C++ library is released](https://www.reddit.com/r/cpp/comments/jgnweb/version_060_of_the_simdjson_c_library_is_released/)
- url: https://github.com/simdjson/simdjson/releases/tag/v0.6.0
---

## [5][Parallelizing GPU-intensive Workloads via Multi-Queue Operations (2x+ Performance Improvements on a Single Graphics Card)](https://www.reddit.com/r/cpp/comments/jh8zq8/parallelizing_gpuintensive_workloads_via/)
- url: https://towardsdatascience.com/parallelizing-heavy-gpu-workloads-via-multi-queue-operations-50a38b15a1dc
---

## [6][Has a release() method for std::vector been proposed?](https://www.reddit.com/r/cpp/comments/jgilit/has_a_release_method_for_stdvector_been_proposed/)
- url: https://www.reddit.com/r/cpp/comments/jgilit/has_a_release_method_for_stdvector_been_proposed/
---
A method akin to

    auto std::vector&lt;T, Alloc&gt;::release() -&gt; std::unique_ptr&lt;T[], _deleter_t&gt;;

where `_deteter_t` is a deleter, that destructs in an appropriate way. The state of the vector after a call to release would be, as if it was moved from.

One often works with libraries that adopt raw pointers on construction of their internals. When coding the algorithms constructing the data to be handed over, `vector` is a convenient class to use, with the api it provides. Moreover, often one uses other libraries for said construction, and they do return a `vector`. But currently, one has to copy the data over to a new pointer before handing it over to be adopted by another structure. Such a `release` method would allow easier data transfer in described situations. Of course, when the `_deleter_t` does something special, one still could not simply hand the pointer over, but this would be implied by returning a `unique_ptr`.

Has anything like this been proposed? If yes, has it been rejected due to drawbacks? If no, is this a method deemed not useful by most?

Edit:

Returning `unique_ptr` instead of the raw pointer, is to be explicit about ownership transfer, and ensure safety until something else adopts ownership over the memory.

Edit II:

## Given the constructive commments bellow

The main use-case for `release()` would be, so one can use the convenience of `vector` in the construction of data, which is than adopted by a third-party library. But when the third-party  does so through a raw pointer (and not a `unique_ptr`)

    auto u_ptr = vec.release();
    // ...
    auto obj = lib::some_adopting_t(u_ptr.release(), size);

this would require that the pointer be deletable by `operator delete[]`, which is not the case even when default allocator is used. Hence `obj` would be deleting the memory in a non-valid way.

The above could still be achieved by using a custom allocator, which allocates in an appropriate way. But this would still leave us without the ability to use other third-party libraries that return `vector`s (using the default allocator) for constructing the data to be adopted (as `obj` above). So it seems the payoff is too small for this to be considered.

Thank you for your comments
## [7][Dealing with very large numbers (primality test)](https://www.reddit.com/r/cpp/comments/jgsgmh/dealing_with_very_large_numbers_primality_test/)
- url: https://www.reddit.com/r/cpp/comments/jgsgmh/dealing_with_very_large_numbers_primality_test/
---

I just learned about Fermat’s primality test and was wondering how it would be implemented with large numbers on a computer. 

Using an unsigned long long variable in C++ would allow you to deal with numbers up to 2^64 - 1, but since Fermat’s theorem uses the term a^(n-1), wouldn’t this mean that it cannot be used to evaluate the primality of any number bigger than 64? As in, if you tried to pass 65, the value stored in the variable would simply default to 0 since 2^64 &gt; 2^64 - 1 (the maximum possible stored value in a ULL variable)?

Using that same logic, even if you created a variable to store a 128-bit number, isn’t it true that you would still only be able to evaluate relatively small prime numbers (up to 128)?

This puzzles me because Fermat Primality Test is sometimes used to deal with unimaginably large numbers (in RSA for example) .... right? 

I feel like I am misunderstanding something here.
## [8][I've rewritten the compiler for my programming language into C++](https://www.reddit.com/r/cpp/comments/jglz0x/ive_rewritten_the_compiler_for_my_programming/)
- url: https://www.reddit.com/r/cpp/comments/jglz0x/ive_rewritten_the_compiler_for_my_programming/
---
My largest C++ project, by far, is the new compiler for my programming language (called AEC), targeting WebAssembly: [https://github.com/FlatAssembler/AECforWebAssembly](https://github.com/FlatAssembler/AECforWebAssembly)

My  previous compiler was targeting x86 and was written in JavaScript.  However, I got a bit dissatisfied with JavaScript. When programming in  JavaScript, you spend a lot of your time debugging errors that, in most  other programming languages, would have been caught by the compiler and  the compiler would issue a warning (or refuse to compile the program).  For instance, in JavaScript, even in the strict mode, if you mistype the  name of a property of some object, no syntax error or reference error  would happen, your program will just continue to go and probably  misbehave in hard-to-debug ways. Dynamic typing, while it can make your  code shorter sometimes, opens a room for a whole class of bugs. In most  programming languages, supplying a different number of arguments to a  function than a function expects will lead to a compile-time warning or  even an error. In JavaScript, the JIT-compiler compiles your code into  something that crashes in hard-to-debug ways. In fact, it will even  accept obviously wrong code such as  \`functions\_that\_returns\_an\_integer()()\` (and, of course, produce  assembly code that crashes). I knew C++ has improved a lot since the  time I first learned it. In old versions of C++, for example, you needed  to write lots of code to convert a number to string, while today there  is a \`std::to\_string\` template function. So, I decided to write my new  compiler in C++. One of the things I like about C++ is that it makes it  very easy to do deep copying, which is something you need a lot when  writing a compiler, and you don't want to waste time thinking about how  you will do that (or so I thought, C++ has some surprising undefined behaviors regarding the default copy constructors: [https://stackoverflow.com/questions/63951270/using-default-copy-constructor-corrupts-a-tree-in-c](https://stackoverflow.com/questions/63951270/using-default-copy-constructor-corrupts-a-tree-in-c)). I knew about Rust, however, in my experience, Rust is  annoying in that it often refuses to compile code in the name of safety,  and sometimes there is no obvious way to do what you want in the way  Rust would consider safe. Rust has many false positives when trying to  detect semantically invalid code, JavaScript has many false negatives.  C++ appears, to me, to be the best in that regard.

The specification for my programming language is available here: [https://flatassembler.github.io/AEC\_specification.html](https://flatassembler.github.io/AEC_specification.html)

The example program I like the most is the Analog Clock in AEC: [https://flatassembler.github.io/analogClock.html](https://flatassembler.github.io/analogClock.html)

So, what do you think about my work?
## [9][Modern C++ Features - a few features that you may know from Python or JavaScript that you can use in C++ too! (For beginners)](https://www.reddit.com/r/cpp/comments/jg5oip/modern_c_features_a_few_features_that_you_may/)
- url: https://blog.yuvv.xyz/modern-cpp-features
---

## [10][CppCast: Programming History, JIT Compilations and Generic Algorithms](https://www.reddit.com/r/cpp/comments/jgfz75/cppcast_programming_history_jit_compilations_and/)
- url: https://cppcast.com/ben-deane-jit-history/
---

## [11][How to boost my learning journey with cpp - currently intermediate level](https://www.reddit.com/r/cpp/comments/jgnmzp/how_to_boost_my_learning_journey_with_cpp/)
- url: https://www.reddit.com/r/cpp/comments/jgnmzp/how_to_boost_my_learning_journey_with_cpp/
---
Hello, I promise you I am not asking about things I can find online.

I am working on my cpp skills currently and I am studying Bjarne Stroustrup books “programming: principles and practice using C++” (finished) and “The C++ Programming language” (currently).

I am trying to have more practice oriented approach by trying to apply what I learn on simple ideas. However, I feel that I can’t reach a good point of understanding if I don’t work on actual projects or have a resource that discuss real programming situations and applications. I am studying C++ for more hardware related applications.

I would appreciate any advice, useful resources, possible opportunities, and/or tips to get the most out of my learning journey. 

My goal is to fully grasp most of the C++ concepts and be able to integrate them.

Thank you so much 😊
