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
## [2][Tortellini: A really, really stupid INI file format for C++11 and above](https://www.reddit.com/r/cpp/comments/feaul5/tortellini_a_really_really_stupid_ini_file_format/)
- url: https://github.com/Qix-/tortellini
---

## [3][Modern std::byte stream IO for C++](https://www.reddit.com/r/cpp/comments/fe72kp/modern_stdbyte_stream_io_for_c/)
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
## [4][Fluent {C++}: How to Pass Class Member Functions to STL Algorithms](https://www.reddit.com/r/cpp/comments/febn53/fluent_c_how_to_pass_class_member_functions_to/)
- url: https://www.fluentcpp.com/2020/03/06/how-to-pass-class-member-functions-to-stl-algorithms/
---

## [5][C++20: Why the word "constinit"? Why not name the specifier complinit or staticinit?](https://www.reddit.com/r/cpp/comments/feb67q/c20_why_the_word_constinit_why_not_name_the/)
- url: https://www.reddit.com/r/cpp/comments/feb67q/c20_why_the_word_constinit_why_not_name_the/
---
If I understand correctly the specifier `constinit` forces initialization of a variable (declared with constinit) at compile time. The value of the variable can be changed later and is not const. I just find that naming the specifier "constinit" could be misunderstood or is misleading in the sense that the variable specified with it isn't const. Perhaps, complinit or staticinit is a better choice for what constinit does.
## [6][C++ random utilities and embedded rng](https://www.reddit.com/r/cpp/comments/fdtsg8/c_random_utilities_and_embedded_rng/)
- url: https://sqrtroot.com/blog/embedded_rng
---

## [7][a speaker that recognizes and plays that particular song on which it has been placed upon.](https://www.reddit.com/r/cpp/comments/fecn76/a_speaker_that_recognizes_and_plays_that/)
- url: https://www.reddit.com/r/cpp/comments/fecn76/a_speaker_that_recognizes_and_plays_that/
---
 So, My grandpa has a lyrical songbook that contains all his favorite songs lyrics in printed pages. I want to make a speaker that when put upon the particular song lyrics plays it. Please don't mind my bad English.
## [8][Invariants and Preconditions](https://www.reddit.com/r/cpp/comments/fdu8gz/invariants_and_preconditions/)
- url: https://www.justsoftwaresolutions.co.uk/cplusplus/invariants.html
---

## [9][CppCast: Packs and Pipelines](https://www.reddit.com/r/cpp/comments/fe7r8e/cppcast_packs_and_pipelines/)
- url: https://cppcast.com/barry-revzin-packs-pipelines/
---

## [10][Did someone propose std::get&lt;&gt; for aggregates and similar helper functions?](https://www.reddit.com/r/cpp/comments/fdwnh5/did_someone_propose_stdget_for_aggregates_and/)
- url: https://www.reddit.com/r/cpp/comments/fdwnh5/did_someone_propose_stdget_for_aggregates_and/
---
As a poor man's reflection, I think it'd be nice. Of course it'd require compiler support to implement.
## [11][Knot: seeing how far you can take a tie of a structs members in place of static reflection](https://www.reddit.com/r/cpp/comments/fdmgo6/knot_seeing_how_far_you_can_take_a_tie_of_a/)
- url: https://www.reddit.com/r/cpp/comments/fdmgo6/knot_seeing_how_far_you_can_take_a_tie_of_a/
---
https://github.com/fried-water/knot

Hey r/cpp,

Given that we lack proper static reflection I was curious how far you could get with a tie of a structs members.

    struct Point { int x; int y; }

    auto as_tie(const Point&amp; p) { return std::tie(p.x, p.y); } 

    void example(const Point&amp; p) {
      std::size_t hash = knot::hash_value(p);
      std::cout &lt;&lt; knot::debug_string(p) &lt;&lt; '\n';

      std::vector&lt;uint8_t&gt; bytes = knot::serialize(p);
      std::optional&lt;Point&gt; p2 = knot::deserialize&lt;Point&gt;(bytes.begin(), bytes.end());
    }

Turns out you can get pretty far, the biggest thing missing is the name of the members. I was able to implement a generic hash, serialize, deserialize, debug string (without member names), lexicographic operators. On top of that there is a `visit()` and `accumulate()` that recursively visits a structs members in a preorder traversal. In addition to structs there are overloads for many of the common std types. ~~This is all possible (names included) with something like BOOST_HANA_DEFINE_STRUCT, but this has the benefit of being non-intrusive.~~

Would love to hear any questions, suggestions, or other use cases you can come up with. Thanks for taking a look!

Edit: forgot BOOST_HANA_ADAPT_STRUCT was a thing
