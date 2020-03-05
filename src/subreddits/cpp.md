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
## [2][C++ random utilities and embedded rng](https://www.reddit.com/r/cpp/comments/fdtsg8/c_random_utilities_and_embedded_rng/)
- url: https://sqrtroot.com/blog/embedded_rng
---

## [3][Knot: seeing how far you can take a tie of a structs members in place of static reflection](https://www.reddit.com/r/cpp/comments/fdmgo6/knot_seeing_how_far_you_can_take_a_tie_of_a/)
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

Turns out you can get pretty far, the biggest thing missing is the name of the members. I was able to implement a generic hash, serialize, deserialize, debug string (without member names), lexicographic operators. On top of that there is a `visit()` and `accumulate()` that recursively visits a structs members in a preorder traversal. In addition to structs there are overloads for many of the common std types. This is all possible (names included) with something like BOOST_HANA_DEFINE_STRUCT, but this has the benefit of being non-intrusive.

Would love to hear any questions, suggestions, or other use cases you can come up with. Thanks for taking a look!
## [4][Invariants and Preconditions](https://www.reddit.com/r/cpp/comments/fdu8gz/invariants_and_preconditions/)
- url: https://www.justsoftwaresolutions.co.uk/cplusplus/invariants.html
---

## [5][Low-Cost Deterministic C++ Exceptions for Embedded Systems](https://www.reddit.com/r/cpp/comments/fdd7vc/lowcost_deterministic_c_exceptions_for_embedded/)
- url: https://www.research.ed.ac.uk/portal/files/78829292/low_cost_deterministic_C_exceptions_for_embedded_systems.pdf
---

## [6][Getting rid of “volatile” in (some of) Qt](https://www.reddit.com/r/cpp/comments/fdip6a/getting_rid_of_volatile_in_some_of_qt/)
- url: https://www.kdab.com/getting-rid-of-volatile-in-some-of-qt/
---

## [7][StarDraft: A C++ project for Starcraft Map Visualization](https://www.reddit.com/r/cpp/comments/fdd3c3/stardraft_a_c_project_for_starcraft_map/)
- url: https://github.com/davechurchill/stardraft
---

## [8][Thoughts on “The C++ Rvalue Lifetime Disaster”](https://www.reddit.com/r/cpp/comments/fdi5pb/thoughts_on_the_c_rvalue_lifetime_disaster/)
- url: https://quuxplusone.github.io/blog/2020/03/04/rvalue-lifetime-disaster/
---

## [9][std::binary_search should return a std::optional&lt;iterator&gt;](https://www.reddit.com/r/cpp/comments/fdj49t/stdbinary_search_should_return_a/)
- url: https://www.reddit.com/r/cpp/comments/fdj49t/stdbinary_search_should_return_a/
---
(like rust)

Currently std::binary_search returns only the boolean value whether your container contains your value. This means that if it does *there's lost information*: the index of the element being searched for is already known but not accessible, you have to call std lower or upper bound to find it. This means that the interface presented to the programmer is less versatile for no good reason. Changing the return value to an iterator to the element found makes the standard library more useful at no performance cost. For example:

    if(auto b = std::binary_search(container.begin(), container.end(), value); b) {
        // in this branch *b is an iterator
    }
## [10][C++17 MySQL client Boost.Asio implementation](https://www.reddit.com/r/cpp/comments/fdes2i/c17_mysql_client_boostasio_implementation/)
- url: https://github.com/anarthal/mysql-asio
---

## [11][Super compact serialisation of C++ classes](https://www.reddit.com/r/cpp/comments/fd8jnf/super_compact_serialisation_of_c_classes/)
- url: https://www.reddit.com/r/cpp/comments/fd8jnf/super_compact_serialisation_of_c_classes/
---
When needing to save many different classes to disk into a human readable format and load them back (a pretty common but very boring task), I figured out this trick, which is probably the shortest way to do it without macros, working with any standard-compliant C++14 compiler (plus MSVC).

    struct Device : SerialisableBrief {
    	int timeout = key("timeout") = 1000;
    	std::string address = key("adress") = "192.168.32.28";
    	bool enabled = key("enabled") = false;
    	std::vector&lt;int&gt; ports = key("ports");
    }

With the inheritance, it gets methods `save()` and `load()` that allow saving it in JSON format as an object with keys `timeout`, `address`, `enabled` and `ports`.

Article how it works: [https://lordsof.tech/programming/super-compact-serialisation-of-c-classes/](https://lordsof.tech/programming/super-compact-serialisation-of-c-classes/)

Full code: [https://github.com/Dugy/serialisable/blob/master/serialisable\_brief.hpp](https://github.com/Dugy/serialisable/blob/master/serialisable_brief.hpp)
