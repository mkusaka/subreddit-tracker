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
## [2][CppCast: Modern C++ for Absolute Beginners](https://www.reddit.com/r/cpp/comments/i9dilw/cppcast_modern_c_for_absolute_beginners/)
- url: https://cppcast.com/modern-cpp-absolute-beginners/
---

## [3][Boost 1.74 is out](https://www.reddit.com/r/cpp/comments/i9i4vi/boost_174_is_out/)
- url: https://www.boost.org/users/history/version_1_74_0.html
---

## [4][What are the best resources for learning cpp](https://www.reddit.com/r/cpp/comments/i9l4ko/what_are_the_best_resources_for_learning_cpp/)
- url: https://www.reddit.com/r/cpp/comments/i9l4ko/what_are_the_best_resources_for_learning_cpp/
---
I'm trying to learn cpp but don't know how to also I've already learned python so will that help
## [5][Order of functions in a translation unit and compiler optimization](https://www.reddit.com/r/cpp/comments/i9ivfs/order_of_functions_in_a_translation_unit_and/)
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
## [6][[CppCon] 2020 Back to Basics Track announced](https://www.reddit.com/r/cpp/comments/i90svg/cppcon_2020_back_to_basics_track_announced/)
- url: https://cppcon.org/b2b2020/
---

## [7][Editable Flow Chart for choosing Containers / Data Structures](https://www.reddit.com/r/cpp/comments/i9lzdw/editable_flow_chart_for_choosing_containers_data/)
- url: https://www.reddit.com/r/cpp/comments/i9lzdw/editable_flow_chart_for_choosing_containers_data/
---
## Summary

After stumbling upon this C++ Container Flow Chart on stackoverflow [here](https://stackoverflow.com/questions/471432/in-which-scenario-do-i-use-a-particular-stl-container/22671607#22671607), I created an editable version of the Flow Chart after realizing no editable version currently exist.

## Download

Download the Editable Container Flow Chart [here](https://github.com/parkershamblin/container-flow-chart#Download).

## Request for Help

I'm reaching out to the Python/C++ community because I need help with creating a Python version of the flow chart. I’m new to C++ and I could use some help from community members who are experienced in Python and C++.

I know there is no perfect mapping between C++ and Python, but I’m confident we could complete 70-75% of the mappings successfully and build something valuable for the Python community.

## Contributions

Please submit your improvements/modifications using this [GitHub Repository](https://github.com/parkershamblin/container-flow-chart). I will be actively reviewing the Repository for any new pull request and updating the Repository frequently. If you need any help with downloading or modifying the Flow Chart please message me anytime. Thanks guys.
## [8][Reflection library for C/C++ (no RTTI)](https://www.reddit.com/r/cpp/comments/i9fi7i/reflection_library_for_cc_no_rtti/)
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
## [9][Any Free courses for begginers ?](https://www.reddit.com/r/cpp/comments/i9jh2r/any_free_courses_for_begginers/)
- url: https://www.reddit.com/r/cpp/comments/i9jh2r/any_free_courses_for_begginers/
---

## [10][container library for type and non-type template parameter packs](https://www.reddit.com/r/cpp/comments/i92fht/container_library_for_type_and_nontype_template/)
- url: https://www.reddit.com/r/cpp/comments/i92fht/container_library_for_type_and_nontype_template/
---
I have been working on a library that adds container functionality to template parameter packs. A couple examples below but it has a lot more functionality than these two.

```

using example = std::tuple&lt;int, short, float, long, float, bool, int, short, float, long, float, bool&gt;;

using example_as_pointers = ext::type_transform_t&lt;ext::type_begin&lt;example&gt;, ext::type_end&lt;example&gt;, std::add_pointer&gt;;

using tuple_binder = ext::type_bind&lt;std::tuple, int, ext::type_placeholder&lt;1&gt;, float, ext::type_placeholder&lt;3&gt;, short, ext::type_placeholder&lt;2&gt;&gt;;

using bound_tuple = tuple_binder::template type&lt;long, char, unsigned long&gt;; // output std::tuple&lt;int, long, float, unsigned long, short, char&gt;

```

It essentially implements the iterator, algorithm, and functional libraries for compile time template parameter packs. It's undocumented  but I am using it in projects at my employer. My question is would it be worth documenting and adding robust testing for sharing with the community. I don't want to waste my time if this is something that nobody will use.
## [11][Updates in Grisu-Exact](https://www.reddit.com/r/cpp/comments/i8pgzc/updates_in_grisuexact/)
- url: https://www.reddit.com/r/cpp/comments/i8pgzc/updates_in_grisuexact/
---
Grisu-Exact is a float-to-string conversion algorithm with

1. The round-trip guarantee,
2. The shortest representation guarantee, and
3. The correct rounding guarantee.

It is a variant of Grisu2 a classic float-to-string conversion algorithm developed by Florian Loitsch, and is largely inspired by Ryu, another float-to-string conversion algorithm developed by Ulf Adams. The main strength of Grisu-Exact over Ryu is that it performs better for short numbers.

Here are some updates:

(1) After having some code simplifications, dead branch eliminations, compression of redundant works, and other optimizations, now it got to outperform Ryu (in my benchmark) for both binary32 (\`float\`) and binary64 (\`double\`) in terms of uniformly randomly generated data.

(2) I developed a string-to-float conversion algorithm using a similar (but simpler) idea, and had a joint-test of it with Grisu-Exact for every possible \`float\`'s. It turned out that for every input data, feeding the data into Grisu-Exact and then into the reverse algorithm correctly generated the input data itself.  Hence, I'm pretty confident about the correctness of Grisu-Exact.

(3) The paper ([https://github.com/jk-jeon/Grisu-Exact/blob/master/other\_files/Grisu-Exact.pdf](https://github.com/jk-jeon/Grisu-Exact/blob/master/other_files/Grisu-Exact.pdf)) explaining the algorithm now includes the explanation on the improved min-max Euclid algorithm. Compared to the original min-max  Euclid algorithm developed by Adams, this improved algorithm produces the exact minimum and maximum not like the original one, has fewer preconditions, and at the same time runs much faster.

Please have a look at the repository [https://github.com/jk-jeon/Grisu-Exact](https://github.com/jk-jeon/Grisu-Exact) if you are interested!
