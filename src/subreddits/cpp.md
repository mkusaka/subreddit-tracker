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
## [2][A common C/C++ core specification](https://www.reddit.com/r/cpp/comments/ffaato/a_common_cc_core_specification/)
- url: https://gustedt.wordpress.com/2020/03/08/a-common-c-c-core-specification/
---

## [3][Software optimization resources. C++ and assembly. Windows, Linux, BSD, Mac OS X](https://www.reddit.com/r/cpp/comments/fev58b/software_optimization_resources_c_and_assembly/)
- url: https://www.agner.org/optimize/
---

## [4][PSA: gsl::span iterators are not just pointers](https://www.reddit.com/r/cpp/comments/feuhzr/psa_gslspan_iterators_are_not_just_pointers/)
- url: https://www.reddit.com/r/cpp/comments/feuhzr/psa_gslspan_iterators_are_not_just_pointers/
---
*Edit: as per this [comment](https://www.reddit.com/r/cpp/comments/feuhzr/psa_gslspan_iterators_are_not_just_pointers/fjwnv3e) from the current GSL maintainer, `gsl::span` iterators will soon be more pointer-like.*

*Edit: as per comments in the thread, `gslite::span` iterators are pointers, and `std::span` iterators will be pointer-like in the upcoming MSVC release.*

---

A colleague stumbled upon a weird optimization in our application; a reduced example is:

    auto index = /*...*/;
    gsl::span&lt;T&gt; view = container.get_view();

    assert(view.size() == 1);
    assert(index == 0);

    view[index] // triggers `Ensure` about the index being without bounds.

After some digging around, it turned out that *later down*, we had a loop:

    for (auto pair : zip(/*...*/, container.get_view())) {
    }

Where `zip` is defined as:

    template &lt;typename... Cs&gt;
    auto zip(Cs const&amp;... containers) { /*...*/ }

Which instantiates two zip iterators (Boost) and returns a range.

Notice that `container.get_view()` returns a `gsl::span` *by value*. Lifetime extension then kicks in so that it lives for as long as `zip(/*...*/, container.get_view())` takes to evaluate, **and no longer**.

Why do we care, though, when the *actual container* lives long enough? Well, it turns out that a `span_iterator&lt;T&gt;` is NOT `T*`, instead it is:

    template &lt;typename T&gt;
    class span_iterator&lt;T&gt; {
         span&lt;T&gt;* __span;
         std::ptrdiff_t __index;
    };

Which is necessary to validate for the `Ensure` machinery used to ensure that only legal operations are applied. This has a number of implications, along which:

 - Iterators of `gsl::span` are only valid as long as their underlying `gsl::span` lives *and* the underlying container lives.
 - By default, `++`, `*`, ... are all **checked**.
 - Even if checks are turned off, `span_iterator` are still twice as heavy as regular `T*`.

It's unclear whether the `std::span` version will take the same approach; if so, we'll likely re-implement `span` ourselves to avoid the overhead and the surprising lifetime implications.
## [5][Analyze your builds programmatically with the C++ Build Insights SDK | Visual C++ Team Blog](https://www.reddit.com/r/cpp/comments/fegtup/analyze_your_builds_programmatically_with_the_c/)
- url: https://devblogs.microsoft.com/cppblog/analyze-your-builds-programmatically-with-the-c-build-insights-sdk/
---

## [6][A few experimental features for C++](https://www.reddit.com/r/cpp/comments/fef2d0/a_few_experimental_features_for_c/)
- url: https://cor3ntin.github.io/posts/qol23/
---

## [7][Tortellini: A really, really stupid INI file format for C++11 and above](https://www.reddit.com/r/cpp/comments/feaul5/tortellini_a_really_really_stupid_ini_file_format/)
- url: https://github.com/Qix-/tortellini
---

## [8][Field-testing “Down with lifetime extension!”](https://www.reddit.com/r/cpp/comments/fefk20/fieldtesting_down_with_lifetime_extension/)
- url: https://quuxplusone.github.io/blog/2020/03/04/field-report-on-lifetime-extension/
---

## [9][Modern std::byte stream IO for C++](https://www.reddit.com/r/cpp/comments/fe72kp/modern_stdbyte_stream_io_for_c/)
- url: https://www.reddit.com/r/cpp/comments/fe72kp/modern_stdbyte_stream_io_for_c/
---
Previous post: https://www.reddit.com/r/cpp/comments/avcalo/a_proposal_to_add_stdbytebased_io_to_the_c/

Direct link to PDF: https://github.com/Lyberta/cpp-io/raw/master/generated/Paper.pdf

Paper repository: https://github.com/Lyberta/cpp-io

Reference implementation: https://github.com/Lyberta/cpp-io-impl

This paper proposes fundamental IO concepts, customization points for serialization and deserialization and streams for memory and file IO.

It's been a year since the last post and quite a few things have changed:

* IO and serialization are separate now. You can use `std::io::read_raw` and `std::io::write_raw` for raw IO. They have much less overloads and will take less time to compile. `std::io::read` and `std::io::write` are more heavy because they support endianness/floating-point-format conversion as well as dispatch to custom serialization functions.
* Concepts instead of inheritance + virtual functions. You no longer need to pay the cost of virtual function calls in generic code.
* Custom serialization functions can now be variadic.
* IO contexts as a way to not have format as part of the stream as well as providing local format support during nested [de]serialization.
* File streams now support buffering.
* Added type erased streams for cases where you need dynamic polymorphism.
* `std::io::in()`, `std::io::out()` and `std::io::err()` for byte IO with standard streams. The objects are type erased so you can redirect them to any stream you want. This is similar to `std::cout` et al.

## It is faster than both `&lt;iostream&gt;` and `&lt;stdio&gt;`

During benchmarking of sequential file IO on Linux proposed `std::io::input_file_stream` was found to be ~30% faster than `std::FILE` and ~45% faster than `std::ifstream` while proposed `std::io::output_file_stream` was found to be ~38% faster than `std::FILE` and ~60% faster than `std::ofstream`. Raw numbers can be found in the paper.

This post was made to gather a round of feedback before I publish R0 targeted for Varna. I'm also looking for a champion to present this proposal in Varna.
## [10][Fluent {C++}: How to Pass Class Member Functions to STL Algorithms](https://www.reddit.com/r/cpp/comments/febn53/fluent_c_how_to_pass_class_member_functions_to/)
- url: https://www.fluentcpp.com/2020/03/06/how-to-pass-class-member-functions-to-stl-algorithms/
---

## [11][C++20: Why the word "constinit"? Why not name the specifier complinit or staticinit?](https://www.reddit.com/r/cpp/comments/feb67q/c20_why_the_word_constinit_why_not_name_the/)
- url: https://www.reddit.com/r/cpp/comments/feb67q/c20_why_the_word_constinit_why_not_name_the/
---
If I understand correctly the specifier `constinit` forces initialization of a variable (declared with constinit) at compile time. The value of the variable can be changed later and is not const. I just find that naming the specifier "constinit" could be misunderstood or is misleading in the sense that the variable specified with it isn't const. Perhaps, complinit or staticinit is a better choice for what constinit does.
