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
## [2][Why has reddit suspended Bjarne Stroustrup's account? Accounts are only suspended manually so somebody at reddit looked at u/bstroustrup's account and decided to suspend it.](https://www.reddit.com/r/cpp/comments/ia6u5t/why_has_reddit_suspended_bjarne_stroustrups/)
- url: https://www.reddit.com/user/bstroustrup
---

## [3][What are some of the best Cheat Sheets that you use?](https://www.reddit.com/r/cpp/comments/i9q480/what_are_some_of_the_best_cheat_sheets_that_you/)
- url: https://www.reddit.com/r/cpp/comments/i9q480/what_are_some_of_the_best_cheat_sheets_that_you/
---
In my experience having quick references is extremely helpful. What are some of the best cheat sheets that you have used or currently use?
## [4][Bizarre SSE/AVX codegen with MSVC Preview](https://www.reddit.com/r/cpp/comments/i9wa79/bizarre_sseavx_codegen_with_msvc_preview/)
- url: https://www.reddit.com/r/cpp/comments/i9wa79/bizarre_sseavx_codegen_with_msvc_preview/
---
I switched to the latest preview recently because I wanted access to more C++20 features to clean some things up, but looking at the assembly all of my aligned SIMD loads and stores through intrinsics have become unaligned, and no matter what hints I give the compiler it will never generate an aligned load or store. I'm not sure if this is a bug, or if there's a genuine reason to prefer unaligned loads and store instructions now on x86-64, but I'm just putting this out there because I know some of the MSVC guys lurk around here.
## [5][Boost 1.74 is out](https://www.reddit.com/r/cpp/comments/i9i4vi/boost_174_is_out/)
- url: https://www.boost.org/users/history/version_1_74_0.html
---

## [6][TIL: on x86-64 compilers will use addressing mode to replace simple multiplications](https://www.reddit.com/r/cpp/comments/i9mgvf/til_on_x8664_compilers_will_use_addressing_mode/)
- url: https://www.reddit.com/r/cpp/comments/i9mgvf/til_on_x8664_compilers_will_use_addressing_mode/
---
obligatory godbolt: [https://gcc.godbolt.org/z/csEnh4](https://gcc.godbolt.org/z/csEnh4)

i was testing a snippet, and i was delighted to see how the compiler optimizes simple multiplications using the x86-64 addressing mode

here's a little table

| op | ~ trasformation |
| --- | ---|
| var*=2 |  var &lt;&lt;= 1 |
| var*=3 | var = var+var*2 |
| var*=4 | var &lt;&lt;= 2 |
| var*=5 | var = var+var*4 |
| var*=6 | tmp = var*3; var = tmp+tmp |
| var*=7 | var = var*8-var |
| var*=8 | var &lt;&lt;= 3 |
| var*=9 | var = var*8+var |
| var*=10 | tmp = var+var*4; var = tmp+tmp |

and here's a nice explanation: [https://paul.bone.id.au/blog/2018/09/05/x86-addressing/](https://paul.bone.id.au/blog/2018/09/05/x86-addressing/)
## [7][CppCast: Modern C++ for Absolute Beginners](https://www.reddit.com/r/cpp/comments/i9dilw/cppcast_modern_c_for_absolute_beginners/)
- url: https://cppcast.com/modern-cpp-absolute-beginners/
---

## [8][Editable Flow Chart for choosing Containers / Data Structures](https://www.reddit.com/r/cpp/comments/i9lzdw/editable_flow_chart_for_choosing_containers_data/)
- url: https://www.reddit.com/r/cpp/comments/i9lzdw/editable_flow_chart_for_choosing_containers_data/
---
## Summary

After stumbling upon this C++ Container Flow Chart on stackoverflow [here](https://stackoverflow.com/questions/471432/in-which-scenario-do-i-use-a-particular-stl-container/22671607#22671607), I created an editable version of the Flow Chart after realizing no editable version currently exist.

## Download

Download the Editable Container Flow Chart [here](https://github.com/parkershamblin/container-flow-chart#Download).

## Contributions

Looking for community members who are experienced in C++ &amp; Python to help create a Python version of the Flow Chart.

Please submit your improvements/modifications using this [GitHub Repository](https://github.com/parkershamblin/container-flow-chart). I will be actively reviewing the Repository for any new pull request and updating the Repository frequently. If you need any help with downloading or modifying the Flow Chart please message me anytime. Thanks guys.
## [9][[CppCon] 2020 Back to Basics Track announced](https://www.reddit.com/r/cpp/comments/i90svg/cppcon_2020_back_to_basics_track_announced/)
- url: https://cppcon.org/b2b2020/
---

## [10][Reflection library for C/C++ (no RTTI)](https://www.reddit.com/r/cpp/comments/i9fi7i/reflection_library_for_cc_no_rtti/)
- url: https://www.reddit.com/r/cpp/comments/i9fi7i/reflection_library_for_cc_no_rtti/
---
Link to the project: [https://github.com/flecs-hub/flecs-meta](https://github.com/flecs-hub/flecs-meta)

flecs.meta is a reflection library that lets you import type definitions without code generation or manually describing field names. A simple example:

    /* ECS_STRUCT macro captures type information */
    ECS_STRUCT(Position, {
        float x;
        float y;
    });
    
    int main(int argc, char *argv[]) {
        // Initialize
        flecs::world world;
        flecs::import&lt;flecs::components::meta&gt;(world);
    
        // Register reflection data for Position
        flecs::meta&lt;Position&gt;(world);
    
        // Use it!
        Position p{10, 20};
        std::cout &lt;&lt; flecs::pretty_print(world, p) &lt;&lt; std::endl;
    }

Output:

    {x = 10, y = 20} 

The library has Flecs as a dependency, but does not require applications to be written in Flecs. A JSON serializer written for the library also exists: [https://github.com/flecs-hub/flecs-json](https://github.com/flecs-hub/flecs-json)

Because of how the library works, it can only be used for trivial C++ types(!)
## [11][Order of functions in a translation unit and compiler optimization](https://www.reddit.com/r/cpp/comments/i9ivfs/order_of_functions_in_a_translation_unit_and/)
- url: https://www.reddit.com/r/cpp/comments/i9ivfs/order_of_functions_in_a_translation_unit_and/
---
Consider this:

1)

    void foo() { some short code }
    (...)
    void bar() { foo(); }

Typically, a modern C++ compiler will inline foo() inside bar()

But if we have this:

2)

    void foo();
    void bar() { foo(); }
    (...)
    void foo() { some short code }

will the same happen? (in ancient times, this did not happen).
