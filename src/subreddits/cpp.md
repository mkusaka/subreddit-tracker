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
## [2][I made this ASCII Ray Tracing program in C++, any feedback and suggestions are welcome](https://www.reddit.com/r/cpp/comments/ibcvks/i_made_this_ascii_ray_tracing_program_in_c_any/)
- url: https://youtu.be/QkETiyYWh2o
---

## [3][Eliminating the static overhead of C++ ranges (2019)](https://www.reddit.com/r/cpp/comments/ib7tt9/eliminating_the_static_overhead_of_c_ranges_2019/)
- url: https://vector-of-bool.github.io/2019/10/21/rngs-static-ovr.html
---

## [4][structopt: Parse command line arguments by defining a struct](https://www.reddit.com/r/cpp/comments/ibdp1p/structopt_parse_command_line_arguments_by/)
- url: https://github.com/p-ranav/structopt
---

## [5][TIL: T{} is not the same as T()](https://www.reddit.com/r/cpp/comments/iatj87/til_t_is_not_the_same_as_t/)
- url: https://www.reddit.com/r/cpp/comments/iatj87/til_t_is_not_the_same_as_t/
---
C++ is a language of surprises. Consider this code:

    struct A
    {
        A() {}
        int v;
    };
    
    struct B
    {
        A a;
    };
    
    int main()
    {
        B b1 = B();
        B b2 = B{};
    }

b1.a.v is 0 while b2.a.v is uninitialized!

See here [https://godbolt.org/z/svhGzP](https://godbolt.org/z/svhGzP)

b1 is value initialized while b2 is aggregate initialized.

If I create a vector of B objects and resize it, it will be initialized by STL using `B()` so v will be zero.

However in many codebases people seem to use {} to initialize existing objects `x = {};`

simply because `x = ();` does not compile. So if you want to initialize a variable to its default

without repeating its type it seems a helper function is needed like:

    template &lt;typename T&gt;
    void init(T&amp; x) { x = T(); }

So now I can write `init(x);` to ensure default/value initialization.

What do you think?
## [6][Blog: New features in SYCL 2020](https://www.reddit.com/r/cpp/comments/ibbdjc/blog_new_features_in_sycl_2020/)
- url: https://www.reddit.com/r/cpp/comments/ibbdjc/blog_new_features_in_sycl_2020/
---
SYCL (pronounced ‘sickle’) is a royalty-free, cross-platform abstraction  layer that enables code for heterogeneous processors to be written in a  “single-source” style using completely standard C++. SYCL enables  single source development where C++ template functions can contain both  host and device code to construct complex algorithms that use the  acceleration provided by processors such as GPUs, and then re-use them  throughout their source code on different types of data. 

A new version, SYCL 2020, is now out for community feedback and this [blog post](https://www.codeplay.com/portal/blogs/2020/08/13/sycl-2020.html) outlines the major new features and enhancements.
## [7][Expansion statements in C++20](https://www.reddit.com/r/cpp/comments/ibc6xv/expansion_statements_in_c20/)
- url: https://www.reddit.com/r/cpp/comments/ibc6xv/expansion_statements_in_c20/
---
Hi,

I do not understand what is the status of the expansion statements feature ( [http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1306r1.pdf](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1306r1.pdf)). 

This report says it was added to C++20:  [https://www.reddit.com/r/cpp/comments/au0c4x/201902\_kona\_iso\_c\_committee\_trip\_report\_c20/](https://www.reddit.com/r/cpp/comments/au0c4x/201902_kona_iso_c_committee_trip_report_c20/) 

This thread basically says it is not:  [https://www.reddit.com/r/cpp/comments/dr1at2/the\_state\_of\_expansion\_statements/](https://www.reddit.com/r/cpp/comments/dr1at2/the_state_of_expansion_statements/) 

Here,  [https://github.com/cplusplus/papers/issues/156](https://github.com/cplusplus/papers/issues/156), the feature is labeled as needs-revision.

Can anybody clarify this?

Thank you.
## [8][Allowing CMake functions to return(value)](https://www.reddit.com/r/cpp/comments/ibbyro/allowing_cmake_functions_to_returnvalue/)
- url: https://oleksandrkvl.github.io/2020/08/09/allowing-cmake-functions-to-return-value.html
---

## [9][Looking for a in object-Oriented analysis and design using C++ for beginners.](https://www.reddit.com/r/cpp/comments/ibcc0c/looking_for_a_in_objectoriented_analysis_and/)
- url: https://www.reddit.com/r/cpp/comments/ibcc0c/looking_for_a_in_objectoriented_analysis_and/
---
Any help will be appreciated very much.
## [10][Why has reddit suspended Bjarne Stroustrup's account? Accounts are only suspended manually so somebody at reddit looked at u/bstroustrup's account and decided to suspend it.](https://www.reddit.com/r/cpp/comments/ia6u5t/why_has_reddit_suspended_bjarne_stroustrups/)
- url: https://www.reddit.com/user/bstroustrup
---

## [11][Throwing Out the Kitchen Sink - Output Ranges](https://www.reddit.com/r/cpp/comments/iaesk7/throwing_out_the_kitchen_sink_output_ranges/)
- url: https://thephd.github.io/output-ranges
---

