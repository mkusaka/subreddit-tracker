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
## [2][what concepts/libraries of c++ are used in industries](https://www.reddit.com/r/cpp/comments/jpiuc3/what_conceptslibraries_of_c_are_used_in_industries/)
- url: https://www.reddit.com/r/cpp/comments/jpiuc3/what_conceptslibraries_of_c_are_used_in_industries/
---
I'm trying to learn more about C++. In my university, the courses do a good job of teaching the basic foundations of C++. But I feel like there is such a big void/gap between what I know and what is expected in industries. I feel like the only things I know are building classes, functions, using the STL and that's about it :(

What I am trying to ask is, what concepts or libraries are useful for me to learn before entering an industry? I know the word "industry" is super broad, but I'm open to anything because I want to learn more.
## [3][Created a convenient (albeit trivial) `concept` for keeping all parameters/arguments explicit—useful for embedded, perhaps.](https://www.reddit.com/r/cpp/comments/jpqtcd/created_a_convenient_albeit_trivial_concept_for/)
- url: https://www.reddit.com/r/cpp/comments/jpqtcd/created_a_convenient_albeit_trivial_concept_for/
---
    // Prevents all implicit conversions.
    //
    // intendend usage (example):
    //   auto aFunction(Strict&lt;ExactlyOneConcreteType&gt; auto parameter) {...}
    //
    template &lt;typename ValueType, typename StrictType&gt;
    concept Strict = std::is_same&lt;StrictType, ValueType&gt;::value;

Please see [this code on the Compiler Explorer](https://godbolt.org/z/e4qdWo) for more examples.

Essentially, it forces you to code even built-in types with type safety. You can't even pass `/\d+/` to a function with the signature `someFunction(Strict&lt;unsigned int&gt; auto)`, you're forced to append a `u`.

Or you could code in assembly, if you want. ¯\\\_(ツ)_/¯.

Please, anyone, speak up if this kills compiler optimization or something like that (I don't see it, but I'm not an expert). Also, yeah, you're not saving *very* much by skipping all the implicit conversions, but I think it's got potential for simply keeping you working *strictly* within the types you intend for things like embedded applications. You don't have to worry about templates/concepts exploding your function count, allowing you to use `concepts` to represent [concepts](https://en.wikipedia.org/wiki/Concept) even in limited environments, allowing you to say what you mean (and no more).

Maybe. I'm going to try it out on my keyboard (customizable firmware)—it's got me excited.
## [4][Using the function std::transform](https://www.reddit.com/r/cpp/comments/jppcwc/using_the_function_stdtransform/)
- url: https://www.reddit.com/r/cpp/comments/jppcwc/using_the_function_stdtransform/
---
I would like to learn modern cpp.. I have a vector of strings and I'd like to convert them into floats. I could use this code:

    std::vector&lt;float&gt; values;
    for (auto number : numbers) 
        values.push_back(std::stof(number));

But I would like to do the same in one line using just std::transform or std::for\_each, maybe instead of putting everything into a new vector (values) just changing straight into numbers. How would you do that?
## [5][[Q]Passing arrays by reference](https://www.reddit.com/r/cpp/comments/jpp7zr/qpassing_arrays_by_reference/)
- url: https://www.reddit.com/r/cpp/comments/jpp7zr/qpassing_arrays_by_reference/
---
I came across the following code:

    template&lt;typename _Ret,   //Return type
             typename _Coll&gt;   //Collection type
    _Ret Sum(const _Coll&amp; c) {
    
     _Ret sum = 0;
     for(auto&amp; v : c)
        sum += v;
     return sum;
    }
    
    int main() {
     //With regular array. Passed as reference
     int arr[] = {1,2,3,4,5};
     std::cout &lt;&lt; Sum&lt;int64_t&gt;(arr) &lt;&lt; "\n"; //15
    
     //With vector
     std::vector&lt;int&gt; vec = {1,2,3,4,5};
     std::cout &lt;&lt; Sum&lt;int64_t&gt;(vec) &lt;&lt; "\n"; //15
     return 0;
    }

If I remove the template and hardcore parameter's datatype [check this](https://godbolt.org/z/5hrdW6) , it doesn't seems to work. Can anyone please explain the reason?
## [6][I built a site to Instant-Search through 32 Million Songs in Milliseconds, using Typesense - a search engine written in C++](https://www.reddit.com/r/cpp/comments/jowldc/i_built_a_site_to_instantsearch_through_32/)
- url: https://songs-search.typesense.org/
---

## [7][Fast open-source intrusion detection](https://www.reddit.com/r/cpp/comments/jpk70l/fast_opensource_intrusion_detection/)
- url: https://github.com/cmu-snap/pigasus
---

## [8][BrainFuck compiler in c++/rust metaprogramming(constexpr/procedural macro solution)](https://www.reddit.com/r/cpp/comments/jp7k0u/brainfuck_compiler_in_crust/)
- url: https://www.reddit.com/r/cpp/comments/jp7k0u/brainfuck_compiler_in_crust/
---
In my last post [BrainFuck compiler in c++ template metaprogramming](https://www.reddit.com/r/cpp/comments/jnz5p1/brainfuck_compiler_in_c_template_metaprogramming/), I used template metaprogramming solution to do this, but now I have done it in c++ and rust in constexpr and procedural macro way (~80 lines), and supports output and nest loop.

c++ version: [BrainFuckConstexpr.cpp](https://github.com/netcan/recipes/blob/master/cpp/metaproggramming/brain_fuck/BrainFuckConstexpr.cpp), [https://godbolt.org/z/EYn7PG](https://godbolt.org/z/EYn7PG), below:

```cpp
// compile time
constexpr auto res = brain_fuck(R"(
    ++++++++[&gt;++++[&gt;++&gt;+++&gt;+++&gt;+&lt;&lt;&lt;&lt;-]&gt;+&gt;+&gt;-&gt;&gt;+[&lt;]&lt;-]&gt;&gt;.
    &gt;---.+++++++..+++.&gt;&gt;.&lt;-.&lt;.+++.------.--------.&gt;&gt;+.&gt;++.
)");
puts(res);

// runtime
if (argc &gt; 1) puts(brain_fuck(argv[1]));
```

rust version: [https://github.com/netcan/recipes/blob/master/rust/BrainFuck/src/lib.rs](https://github.com/netcan/recipes/blob/master/rust/BrainFuck/src/lib.rs), below:

```rust
extern crate brain_fuck;
use brain_fuck::brain_fuck;
fn main() {
    println!("{}", brain_fuck!(
        ++++++++[&gt;++++[&gt;++&gt;+++&gt;+++&gt;+&lt;&lt;&lt;&lt;-]&gt;+&gt;+&gt;-&gt;&gt;+[&lt;]&lt;-]&gt;&gt;.
        &gt;---.+++++++..+++.&gt;&gt;.&lt;-.&lt;.+++.------.--------.&gt;&gt;+.&gt;++.
    ));
}
```
## [9][Project Ideas](https://www.reddit.com/r/cpp/comments/jpktw7/project_ideas/)
- url: https://www.reddit.com/r/cpp/comments/jpktw7/project_ideas/
---
I've been learning computer science for a little over 6 months now. I read and completed a lot of the exercises in the C Programming book by K&amp;R. I then built a personal website in python using flask running on heroku which was cool but ehh. I'm currently working through Cal Berkley's CS61A, CS61B, and CS61C classes in order now.

Since I am getting a little burnt out on formal learning I wanted to take a break to start a finance-related project as I find investing interesting. When I created the CONOPS of projects and played around with APIs in python I usually get stuck at the API level for pulling data. No free API will allow me to make that many calls on that many tickers enough to really build out any somewhat respectable project. I want to use C++ so I can improve the performance and optimize the code rather than making a slow python project. I imagine future employers would be more interested in good code where the person understands the underlying implementation than a handful of python wrappers.

With all that being said, I was looking for some ideas for projects to showcase the performance of C++.

\-Thanks
## [10][FTL - A functional template library for containers processing in C++](https://www.reddit.com/r/cpp/comments/jor2hd/ftl_a_functional_template_library_for_containers/)
- url: https://www.reddit.com/r/cpp/comments/jor2hd/ftl_a_functional_template_library_for_containers/
---
Hi everyone!

I would like to present you a C++17 library that me and my friend were working on recently. The library is still very much WIP so any suggestions are welcome.

The main idea was to provide a more elegant functional API for containers processing. The API itself is heavily inspired by the way the Rust programming language deals with container processing and iterators.

If you have any suggestions or feedback do not hesitate to open issues/pull requests or reach out to us!

[https://github.com/ftlorg/ftl](https://github.com/ftlorg/ftl)
## [11][matio-cpp: a C++ wrapper for the matio library](https://www.reddit.com/r/cpp/comments/jp134a/matiocpp_a_c_wrapper_for_the_matio_library/)
- url: https://github.com/dic-iit/matio-cpp
---

