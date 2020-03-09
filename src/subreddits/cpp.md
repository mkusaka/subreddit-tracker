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
## [2][How to remove cplusplus.com from Google search results?](https://www.reddit.com/r/cpp/comments/ffpd1s/how_to_remove_cpluspluscom_from_google_search/)
- url: https://www.reddit.com/r/cpp/comments/ffpd1s/how_to_remove_cpluspluscom_from_google_search/
---
it even \*pretends\* to be cppreference, that's how bad it is! &lt;/rant&gt;

Edit: answer in comments, works for Firefox which I happen to use: [https://www.reddit.com/r/cpp/comments/ffpd1s/how\_to\_remove\_cpluspluscom\_from\_google\_search/fk00gpl?utm\_source=share&amp;utm\_medium=web2x](https://www.reddit.com/r/cpp/comments/ffpd1s/how_to_remove_cpluspluscom_from_google_search/fk00gpl?utm_source=share&amp;utm_medium=web2x)

Edit2: personal block list [https://www.reddit.com/r/cpp/comments/ffpd1s/how\_to\_remove\_cpluspluscom\_from\_google\_search/fk00noj?utm\_source=share&amp;utm\_medium=web2x](https://www.reddit.com/r/cpp/comments/ffpd1s/how_to_remove_cpluspluscom_from_google_search/fk00noj?utm_source=share&amp;utm_medium=web2x)

edit3: how to improve the proper reference: [https://www.reddit.com/r/cpp/comments/ffpd1s/how\_to\_remove\_cpluspluscom\_from\_google\_search/fk06he5?utm\_source=share&amp;utm\_medium=web2x](https://www.reddit.com/r/cpp/comments/ffpd1s/how_to_remove_cpluspluscom_from_google_search/fk06he5?utm_source=share&amp;utm_medium=web2x)
## [3][Demo: C++20 Concepts Feature](https://www.reddit.com/r/cpp/comments/ffdsvd/demo_c20_concepts_feature/)
- url: https://youtu.be/B_KjoLid5gw
---

## [4][Tips for working with GCC in a MSVC compatible way](https://www.reddit.com/r/cpp/comments/ffqvda/tips_for_working_with_gcc_in_a_msvc_compatible_way/)
- url: https://www.reddit.com/r/cpp/comments/ffqvda/tips_for_working_with_gcc_in_a_msvc_compatible_way/
---
Hello everyone, 

First time post in here, I don't know whether this is the right sub to ask, but alas, here I go.

So I am a CompSci freshman, and i have OOP this semester, which is done in C/C++.

As I have an old laptop(and I prefer GNU/Linux as an OS) I am doing pretty much everything in Emacs and with GCC, but my OOP lab teacher has an online platform for checking our assignments and "demands" us working in Visual Studio under MSVC, which my computer cannot run at all.

So, could you point me towards some material, or give me some advice, as to how could I finish this semester with GCC without failing?

&amp;#x200B;

Many thanks in advance.
## [5][Blacksmith - Library for in-place class construction.](https://www.reddit.com/r/cpp/comments/ffqc7l/blacksmith_library_for_inplace_class_construction/)
- url: https://www.reddit.com/r/cpp/comments/ffqc7l/blacksmith_library_for_inplace_class_construction/
---
Hi there, I've been toying with this idea for a while now, and have settled on a simple solution.

Basically, I wanted to build classes in-place using named arguments (I know designated initialisers are coming, but in the mean time...)

    auto p = build([](Point&amp; _) { _.x = 5; _.y = 6;});

Or for pointers:

    build_shared([](Point&amp; _) { _.x = 5; _.y = 6;}); // returns shared_ptr&lt;Point&gt;
    
    build_unique([](Point&amp; _) { _.x = 5; _.y = 6;}); // returns unique_ptr&lt;Point&gt;
    
    build_new([](Point&amp; _) { _.x = 5; _.y = 6;}); // returns a new Point* (must delete)

This is my first open source c++ contribution, so I'd love some feedback! Cheers.

[https://github.com/martypapa/blacksmith](https://github.com/martypapa/blacksmith)
## [6][A common C/C++ core specification](https://www.reddit.com/r/cpp/comments/ffaato/a_common_cc_core_specification/)
- url: https://gustedt.wordpress.com/2020/03/08/a-common-c-c-core-specification/
---

## [7][Software optimization resources. C++ and assembly. Windows, Linux, BSD, Mac OS X](https://www.reddit.com/r/cpp/comments/fev58b/software_optimization_resources_c_and_assembly/)
- url: https://www.agner.org/optimize/
---

## [8][PSA: gsl::span iterators are not just pointers](https://www.reddit.com/r/cpp/comments/feuhzr/psa_gslspan_iterators_are_not_just_pointers/)
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
## [9][Analyze your builds programmatically with the C++ Build Insights SDK | Visual C++ Team Blog](https://www.reddit.com/r/cpp/comments/fegtup/analyze_your_builds_programmatically_with_the_c/)
- url: https://devblogs.microsoft.com/cppblog/analyze-your-builds-programmatically-with-the-c-build-insights-sdk/
---

## [10][A few experimental features for C++](https://www.reddit.com/r/cpp/comments/fef2d0/a_few_experimental_features_for_c/)
- url: https://cor3ntin.github.io/posts/qol23/
---

## [11][Tortellini: A really, really stupid INI file format for C++11 and above](https://www.reddit.com/r/cpp/comments/feaul5/tortellini_a_really_really_stupid_ini_file_format/)
- url: https://github.com/Qix-/tortellini
---

