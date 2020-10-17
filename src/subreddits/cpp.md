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
## [2][What's the future gonna be c++?](https://www.reddit.com/r/cpp/comments/jcqgy1/whats_the_future_gonna_be_c/)
- url: https://www.reddit.com/r/cpp/comments/jcqgy1/whats_the_future_gonna_be_c/
---
Hi, I just got my first job as c++ developer and am yet to go through training. But I just wanted to hear from the community about what the future is for c++?? While doing my bachelor's. everyone was like c++ was only used for learning the fundamentals and  soon it will cease to exist. I just want to know what all can a cpp developer can do? Are there any opportunity?
## [3][Writing terminal moo: On C++, its ecosystem, performance optimization and more](https://www.reddit.com/r/cpp/comments/jc8y8q/writing_terminal_moo_on_c_its_ecosystem/)
- url: https://github.com/s9w/terminal_moo/blob/master/text/text.md
---

## [4][GPU Accelerated Machine Learning on Android Phones using the Android C++ NDK and Vulkan Kompute](https://www.reddit.com/r/cpp/comments/jcv5my/gpu_accelerated_machine_learning_on_android/)
- url: https://towardsdatascience.com/gpu-accelerated-machine-learning-in-your-mobile-applications-using-the-android-ndk-vulkan-kompute-1e9da37b7617
---

## [5][std::get_temporary_buffer down sides](https://www.reddit.com/r/cpp/comments/jcuxs5/stdget_temporary_buffer_down_sides/)
- url: https://www.reddit.com/r/cpp/comments/jcuxs5/stdget_temporary_buffer_down_sides/
---
Hello guys, I'm currently writting a memory allocator, and i was thinking of using get_temporary_buffer for the memory allocation. So I was wondering what are the downsides of using it?
## [6][std::filesystem::path::string() (and related) should have a `&amp;&amp;` and `const&amp;` overloads](https://www.reddit.com/r/cpp/comments/jcc2av/stdfilesystempathstring_and_related_should_have_a/)
- url: https://www.reddit.com/r/cpp/comments/jcc2av/stdfilesystempathstring_and_related_should_have_a/
---
One of my projects is finally moving to C++17. One of the things I wanted to do was something like this:

    fs::path p = get_path(); // Could be anything, just as long as you don't have the string up front.
    return p.string();

Now this works fine, but it copies the string, because [`path::string()`](https://en.cppreference.com/w/cpp/filesystem/path/string) is const qualified.

&amp;nbsp;

Let's take a look at [`std::stringstream::str()`](https://en.cppreference.com/w/cpp/io/basic_stringstream/str) in C++20. Before C++20, `str()` was just

    std::string str() const;

But [P0408](https://wg21.link/P0408) made some changes and now we get:

    std::string str() const&amp;; // Same as before
    std::string str() &amp;&amp;; // Moves the string from the stream's buffer

The overload like the above allows one to write:

    std::string f(std::stringstream ss) {
        return std::move(ss).str();  // Moves the string
    }

I see no reason why `std::filesystem::path` can't get the same treatment. If we had a similar overload set, like this:

    std::string string() const&amp;;
    std::string string() &amp;&amp;;

then we could have more efficient retrieval of file names by letting us write:

    return std::move(path).string();

And I guess all the other [format observers](https://en.cppreference.com/w/cpp/filesystem/path#Format_observers) (except for `c_str()`) can benefit from the same treatment.

&amp;nbsp;

Have I overlooked something? Is this not feasible for any reason? Or was this just overlook when `std::filesystem` got standardized?
## [7][Forward-declaring customization points and 'simple' vocab types](https://www.reddit.com/r/cpp/comments/jcrcge/forwarddeclaring_customization_points_and_simple/)
- url: https://www.reddit.com/r/cpp/comments/jcrcge/forwarddeclaring_customization_points_and_simple/
---
News to no-one: adding declarations to `namespace std` is UB (with a few standard-sanctioned exceptions).

[Extending the namespace std](https://en.cppreference.com/w/cpp/language/extending_std) suggests that a compile-time-friendly way of handling this is to not necessarily use the header that corresponds directly to the type you want, but find *some other* more lightweight header that *also* provides a definition:

    // Get the declaration of the primary std::hash template.
    // We are not permitted to declare it ourselves.
    // &lt;typeindex&gt; is guaranteed to provide such a declaration,
    // and is much cheaper to include than &lt;functional&gt;.

That is, frankly, not very cash money. I'd like to not pay for what I'm not using!

It seems as though a lightweight `namespace std` forward declaration header could be put together fairly easily, e.g. [this gist](https://gist.github.com/marzer/f455ddcfcdad06afd3ae552ab15abada). I know that's potentially problematic since different stdlibs will use inline namespaces and/or add additional template arguments so it'd be next to impossible to do in a portable way, but assuming I'm not interested in being portable: what's the potential harm of this *really*? Do compilers somehow bless code from the standard headers and treat it differently?
## [8][Does C++ make *you* feel awesome?](https://www.reddit.com/r/cpp/comments/jcat30/does_c_make_you_feel_awesome/)
- url: https://www.reddit.com/r/cpp/comments/jcat30/does_c_make_you_feel_awesome/
---
I'm not asking if you think the language is awesome.

I'm not asking for examples of awesome applications written in it.

I'm not asking if the community is awesome.

I'm asking you how *you* feel when reading, writing and learning C++. Do *you* feel awesome?

Me? Over the years, it seems like I've witnessed the "marketing" (be it coporate or community-based) of other languages focus on what its practitioners can achieve, on how bankable the language's ecosystem is. We may think ourselves above all that but...

I feel that C++, on the other hand, has focused solely on its technicity. It's a good argument to convince yourself of the language's power but I don't think it's enough. Or, at least, it's not enough to convince the next generation of developers^(1) to take it on. I feel that over time the learning curve is *not* flattening, that I'm less and less productive with it and that I'm more and more confused about where the line between C++-for-day-to-day-practitioners and C++-for-std-library-developers-mega-minds stands. 

How long before C++ is no longer the luckily sole common language between iOS and Android? How long before Rust with its built-in memory safety and sane package management makes it to Playstation 5 and Xbox [latest]? Python has already taken entire domains (ML^(2), robotics, RasberryPi-based projects, etc.) because, hell, you're gonna be just that much more productive with it. C++ is becoming the right tool for fewer and fewer jobs.

I want to see more "Only in C++" marketing and not the "because you had no choice" kind.

p.s. About the title. I wish I could find that video again: a presentation by a marketing lady from about 10 years ago that made the point that's it's not the awesome CEO, the awesome company or the awesome product that wins the marketing battle. It's who or whatever makes *the customer* feel awesome.

[1] "Next generation" here englobes the young students and the 35 year old accountants eyeing bootcamps.

[2] I *know* "the underlying code is C++". You don't need to tell me. Also, that's not my point.
## [9][Google's TensorFlow management is garbage](https://www.reddit.com/r/cpp/comments/jbzywe/googles_tensorflow_management_is_garbage/)
- url: https://www.reddit.com/r/cpp/comments/jbzywe/googles_tensorflow_management_is_garbage/
---
This is a follow-up post to [https://www.reddit.com/r/cpp/comments/hwr6u9/google\_c\_libraries\_are\_garbage](https://www.reddit.com/r/cpp/comments/hwr6u9/google_c_libraries_are_garbage)

It's now been almost 4 months since I posted a PR to the TensorFlow GitHub repo. All I've seen so far, was people being assigned as reviewers, then assigning other people as reviewers. I mean we already have to patch our TensorFlow build in multiple ways to make it function at least somewhat as a proper C++ library. Maybe they're ignoring me because of these Reddit posts, but it's academic at this point I think. Until today, all the activity I've seen was one Google employee begging for others to take action.

Anyway, here's the PR, feel free to take it apart,  critique it, maybe tell me how stupid an incompetent I really am (aka review the PR): [https://github.com/tensorflow/tensorflow/pull/41733](https://github.com/tensorflow/tensorflow/pull/41733)

P.S. I want to thank everyone for a great discussion. It's been truly interesting, inspiring, and insightful. I don't know if I have another clickbaty Google-hating (I don't actually hate Google) post in me, but it's been a lot of fun, and I've learned a lot too. Thank you, and keep trolling :)
## [10][GUI libraries with support for RTL languages](https://www.reddit.com/r/cpp/comments/jc7sb2/gui_libraries_with_support_for_rtl_languages/)
- url: https://www.reddit.com/r/cpp/comments/jc7sb2/gui_libraries_with_support_for_rtl_languages/
---
Was wondering if any c++ gui libraries apart from Qt also support RTL languages.
## [11][AnyOf and AnyBut concepts](https://www.reddit.com/r/cpp/comments/jcg3wf/anyof_and_anybut_concepts/)
- url: https://www.reddit.com/r/cpp/comments/jcg3wf/anyof_and_anybut_concepts/
---
While I was searching for an elegant solution for [this](https://www.reddit.com/r/cpp/comments/jc69ob/nonconst_ref_copy_ctor_should_be_defaulted_if_the/), I realized that `T(AnyBut&lt;T&gt; auto&amp;&amp;)` is the most elegant solution currently available. I was a bit surprised that something this fundamental and basic wasn't included in the standard library though.

see: [https://godbolt.org/z/4T518h](https://godbolt.org/z/4T518h)
