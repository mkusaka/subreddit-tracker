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
## [2][GCC: C++ coroutines - Initial implementation pushed to master.](https://www.reddit.com/r/cpp/comments/eqrv1n/gcc_c_coroutines_initial_implementation_pushed_to/)
- url: https://gcc.gnu.org/ml/gcc-patches/2020-01/msg01096.html
---

## [3][P2064: "Assumptions" by Herb Sutter](https://www.reddit.com/r/cpp/comments/eqv5i8/p2064_assumptions_by_herb_sutter/)
- url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/p2064r0.pdf
---

## [4][Account of search for a minimalistic bounded blocking MPMC queue](https://www.reddit.com/r/cpp/comments/equle2/account_of_search_for_a_minimalistic_bounded/)
- url: https://morestina.net/blog/1400/minimalistic-blocking-bounded-queue-for-c
---

## [5][Survey about C++ IO library status. Are iostreams allowed in your project?](https://www.reddit.com/r/cpp/comments/eqtn2b/survey_about_c_io_library_status_are_iostreams/)
- url: https://docs.google.com/forms/d/e/1FAIpQLSfruC_zT1d8VZsjcQNZ_nC_9NgY3OrfWcVWsSxKogVr2ifd6A/viewform
---

## [6][Prefetching - practical applications](https://www.reddit.com/r/cpp/comments/eqm4qc/prefetching_practical_applications/)
- url: https://www.reddit.com/r/cpp/comments/eqm4qc/prefetching_practical_applications/
---
 [https://extensa.tech/blog/prefetching/](https://extensa.tech/blog/prefetching/) 

Continuation of previous post:

 [https://www.reddit.com/r/cpp/comments/e6fuu0/blogposts\_about\_various\_methods\_of\_packing/](https://www.reddit.com/r/cpp/comments/e6fuu0/blogposts_about_various_methods_of_packing/)
## [7][Exploiting C++ string literal operator: GLSL vector swizzling in modern C++17](https://www.reddit.com/r/cpp/comments/eqoc0y/exploiting_c_string_literal_operator_glsl_vector/)
- url: https://github.com/TheMaverickProgrammer/Compile-Time-Vector-Swizzling
---

## [8][C++ Standards Committee Papers: 2020-01 pre-Prague mailing](https://www.reddit.com/r/cpp/comments/eqc3kz/c_standards_committee_papers_202001_preprague/)
- url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/#mailing2020-01
---

## [9][Ref target lifetime/ownership management?](https://www.reddit.com/r/cpp/comments/eqld7n/ref_target_lifetimeownership_management/)
- url: https://www.reddit.com/r/cpp/comments/eqld7n/ref_target_lifetimeownership_management/
---
Is there a general principle that lets you easily reason about the lifetime and "proper" ownership of a ref's target?

Despite being a relatively experienced C++ programmer, I still occasionally find myself fighting lifetime-management. Especially with refs. I very often want to write code that looks like this:

```
struct foo {
  const string&amp; x;
};
foo create_foo() {
  string y = get_string();
  return foo{y};
}
```

E.g. the ref is to a stack-allocated object that may go out of scope.

In the above example I could obviously have `foo` own the string, but there are lots of cases where I really do just want a ref, and it's often hard at first to know exactly where to pin that ownership. Of course `shared_ptr` is also an option but I try to avoid sharing `shared_ptr` across types or instances.

Experience helps, but usually I just have to sit and reason very carefully about it even though the compiler has all the information it needs to spot this bug in many cases. Usually the ownership becomes more obvious over time but only after being bit multiple times in debugging by "null" refs that cascaded from far-away-stacks.

I only realized C++ was really missing this when I started poking more at Rust recently. This seems to be something that Rust gets right with its *lifetime is part of the signature* / borrow-checker and one thing that C++ gets "wrong" (doesn't capture directly) despite the common "use smart ptrs and you won't have memory issues" saying.
## [10][DynamicArray: Yet Another C++ Array](https://www.reddit.com/r/cpp/comments/eqovjn/dynamicarray_yet_another_c_array/)
- url: https://www.reddit.com/r/cpp/comments/eqovjn/dynamicarray_yet_another_c_array/
---
Github link:  [https://github.com/LeonineKing1199/sleip](https://github.com/LeonineKing1199/sleip)

Have you ever needed a fixed-sized dynamic allocation? If so, `sleip::dynamic_array` is the class for you! `dynamic_array` is fully Allocator-aware and also supports array objects properly.

Meaning the following is possible:

    auto x = sleip::dynamic_array&lt;int[2][2](...);
    auto y = x; // properly copied over

The container requires C++17 as a minimum and is available via a custom vcpkg port repo in addition to supporting regular CMake.

Make sure to read the docs [here](https://github.com/LeonineKing1199/sleip/blob/develop/doc/dynamic_array.adoc#dynamic_array--fixed-size-dynamic-array)
## [11][Allocating large blocks of memory: bare-metal C++ speeds](https://www.reddit.com/r/cpp/comments/eqhckl/allocating_large_blocks_of_memory_baremetal_c/)
- url: https://lemire.me/blog/2020/01/17/allocating-large-blocks-of-memory-bare-metal-c-speeds/
---

