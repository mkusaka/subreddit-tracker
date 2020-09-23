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
## [2][September 2020 C++ standard mailing](https://www.reddit.com/r/cpp/comments/iy77zr/september_2020_c_standard_mailing/)
- url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/#mailing2020-09
---

## [3][Implementation Challenge: Replacing std::move and std::forward](https://www.reddit.com/r/cpp/comments/ixwtcm/implementation_challenge_replacing_stdmove_and/)
- url: https://foonathan.net/2020/09/move-forward/
---

## [4][Back to Basics: The Abstract Machine - Bob Steagall - CppCon 2020](https://www.reddit.com/r/cpp/comments/ixv8d3/back_to_basics_the_abstract_machine_bob_steagall/)
- url: https://www.youtube.com/watch?v=ZAji7PkXaKY
---

## [5][Back to Basics: The Structure of a Program - Bob Steagall - CppCon 2020](https://www.reddit.com/r/cpp/comments/ixz9gn/back_to_basics_the_structure_of_a_program_bob/)
- url: https://www.youtube.com/watch?v=3KoXeegncrs
---

## [6][Powerful C++ metaprogramming: find all shortest path between nodes in Graph at compile/run time!](https://www.reddit.com/r/cpp/comments/iy7usx/powerful_c_metaprogramming_find_all_shortest_path/)
- url: https://www.reddit.com/r/cpp/comments/iy7usx/powerful_c_metaprogramming_find_all_shortest_path/
---
for more details, see [https://godbolt.org/z/av4Yb1](https://godbolt.org/z/av4Yb1), [https://github.com/netcan/recipes/blob/master/cpp/metaproggramming/FindShortestPath.cpp](https://github.com/netcan/recipes/blob/master/cpp/metaproggramming/FindShortestPath.cpp)

demo:

    struct A: Node&lt;'A'&gt; {};
    struct B: Node&lt;'B'&gt; {};
    struct C: Node&lt;'C'&gt; {};
    struct D: Node&lt;'D'&gt; {};
    struct E: Node&lt;'E'&gt; {};
    using g = Graph&lt;
        __link(__node(A) -&gt; __node(B) -&gt; __node(C) -&gt; __node(D)),
        __link(__node(A) -&gt; __node(C)),
        __link(__node(B) -&gt; __node(A)),
        __link(__node(A) -&gt; __node(E)) &gt;;
    
    static_assert(g::getPath('A', 'D').sz == 3);
    static_assert(g::getPath('A', 'A').sz == 1);
    static_assert(g::getPath('B', 'D').sz == 3);
    static_assert(g::getPath('B', 'E').sz == 3);
    static_assert(g::getPath('D', 'E').sz == 0);
    
    int main(int argc, char** argv) {
        char from = 'A';
        char to = 'D';
        if (argc &gt; 2) {
            from = argv[1][0]; // A
            to = argv[2][0]; // D
        }
        auto path = g::getPath(from, to);
        std::cout &lt;&lt; "from " &lt;&lt; from &lt;&lt; " to " &lt;&lt; to &lt;&lt; " path size: " &lt;&lt; path.sz &lt;&lt; std::endl;
        for (size_t i = 0; i &lt; path.sz; ++i) {
            std::cout &lt;&lt; path.path[i];
            if (i != path.sz - 1) {
                std::cout &lt;&lt; "-&gt;";
            }
        }
        std::cout &lt;&lt; std::endl;
    
        return 0;
    }
## [7][What makes modern C++ modern?](https://www.reddit.com/r/cpp/comments/iy06mn/what_makes_modern_c_modern/)
- url: https://www.reddit.com/r/cpp/comments/iy06mn/what_makes_modern_c_modern/
---
I understand that "modern" C++began with C++11, but I would guess not every feature is required to be considered modern

What C++ features would you really consider the fundamental features for a C++ codebase to be considered modern?  Is it smart pointers?  Auto keyword? Constexpr?

(Go easy on me, I'm a C dev by trade)
## [8][CppCon 2020 videos](https://www.reddit.com/r/cpp/comments/ixyfs8/cppcon_2020_videos/)
- url: https://www.youtube.com/playlist?list=PLHTh1InhhwT6VxYHtoWIvOup9gz0p95Qr
---

## [9][September 23: Free Talk with Emery Berger on Optimizing Application Performance](https://www.reddit.com/r/cpp/comments/ixx2bz/september_23_free_talk_with_emery_berger_on/)
- url: https://www.reddit.com/r/cpp/comments/ixx2bz/september_23_free_talk_with_emery_berger_on/
---
On September 23, join 2019 ACM Fellow Emery Berger (Professor of Computer Science at UMass-Amherst) for the free ACM TechTalk, "[Performance (Really) Matters.](https://webinars.on24.com/acm/berger?partnerref=red)"

Learn why current approaches to evaluating and optimizing performance don't work; how complicated performance has become on modern systems, and how compiler optimizations have essentially run out of steam; and learn about a couple of radically new performance profilers that could help.

One is is Coz, a new "causal profiler" for C/C++/Rust that lets programmers optimize for throughput or latency, and which pinpoints and accurately predicts the impact of optimizations via what we call "virtual speedup" experiments. Coz's approach unlocks previously unknown optimization opportunities. Guided by Coz, we improved the performance of applications by as much as 68%; in most cases, this involved modifying less than 10 lines of code and took under half an hour (without any prior understanding of the programs!). Coz now ships as part of standard Linux distros (apt install coz-profiler). 

[Register](https://webinars.on24.com/acm/berger?partnerref=red) for free to attend live or be alerted when a recording is available.
## [10][Ideal ratio of library - number of files - number of lines per file - number of lines per function](https://www.reddit.com/r/cpp/comments/iy7cg6/ideal_ratio_of_library_number_of_files_number_of/)
- url: https://www.reddit.com/r/cpp/comments/iy7cg6/ideal_ratio_of_library_number_of_files_number_of/
---
For most of the questions the right answer is "depends", and I know there is no typical C++ project.
Nevertheless, in general, what would be for you the ideal ratio between the various metrics you see in a typical C++ project? and above all, what are the underlying motivations?

I would say:

-Functions -&gt; as small as possible, even [1-5] lines per function are ok
-Members per class -&gt; as few as possible: [1-5] members.. more members acceptable for complex algorithms
-Number of lines per files -&gt; in this case it depends on the files included (dependencies), but even here less is better.  You must always be able to include only what you need, not more.

-Number of Libraries per application: in this case,  in my opinion it is reasonable to have many (indirect) library dependencies, or better this is the output of my reasoning.

Why this reasoning? in two words: complexity management
## [11][This feels weird, but it does work. (std::map with a std::any, get/set templates). Please tell me what you think.](https://www.reddit.com/r/cpp/comments/ixn7lq/this_feels_weird_but_it_does_work_stdmap_with_a/)
- url: https://www.reddit.com/r/cpp/comments/ixn7lq/this_feels_weird_but_it_does_work_stdmap_with_a/
---
I need opinions if this is good/bad/ugly/python. If you would do a code review and saw this, what would you think? Do you feel uncomfortable with all this dynamic auto unknown? 

    /* some legacy class which there are a lot of instances off, which sometimes need to transfer data to another instance */

    std::map&lt;std::string, std::any&gt; _tVars;

    template &lt;typename T&gt;
    T getValue(const std::string &amp;key, T defaultValue) const
    {
        auto it = _tVars.find(key);
        if (it == _tVars.end())
            return std::any_cast&lt;T&gt;(defaultValue);

        return std::any_cast&lt;T&gt;(it-&gt;second);
    };

    template &lt;typename T&gt;
    void setValue(const std::string &amp;key, T value)
    {
        _tVars[key] = value;
    };


It's used as a sort of pseudo k/v:

    exampleClass.setValue("blah", aRandomTypeInstance);
    // somewhere else
    auto blah = exampleClass.getValue("blah", defaultValue);
   // blah should be aRandomTypeInstance, if it is not set, then it's a "defaultValue"
The programmer must ensure both sides use the correct key and validate the data in it.

The data that is passed around is unknown beforehand, either comes from a sybase database or user defined json (sometimes having base64 image in it). It's used purely for passing around data between class instances.
