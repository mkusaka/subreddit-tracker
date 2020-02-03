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
## [2][Libc++’s implementation of std::string](https://www.reddit.com/r/cpp/comments/ey464c/libcs_implementation_of_stdstring/)
- url: https://joellaity.com/2020/01/31/string.html
---

## [3][Activity Indicators - Example of a Modern C++ Library](https://www.reddit.com/r/cpp/comments/ey4c0g/activity_indicators_example_of_a_modern_c_library/)
- url: https://www.bfilipek.com/2020/02/inidicators.html
---

## [4][What kind of static code analyser tools do you use?](https://www.reddit.com/r/cpp/comments/ey508p/what_kind_of_static_code_analyser_tools_do_you_use/)
- url: https://www.reddit.com/r/cpp/comments/ey508p/what_kind_of_static_code_analyser_tools_do_you_use/
---
Hi,

I would like to use a static code analyser tool in my team but there are several options to choose from, so I would like to ask you, what static code analyser tool do you use, and how do you like it?

I am more interested in these: Gimpel's PC-Lint, SonarQube, and Coverity.

Thanks
## [5][Does C++ allow you to express the following wait-free RMW operations?](https://www.reddit.com/r/cpp/comments/exvuev/does_c_allow_you_to_express_the_following/)
- url: https://www.reddit.com/r/cpp/comments/exvuev/does_c_allow_you_to_express_the_following/
---
I have implemented a lock free linked list by following the work from this [paper](https://www.cl.cam.ac.uk/research/srg/netos/papers/2001-caslists.pdf).

To achieve all the reference counting and atomicity requirements, you rely on the following central data structure:

    struct ref_ptr {
        T* ptr;
        int count;
        bool mark;
    }

which can be represented in (under) 16 bytes.

The main operations are:

1. &amp;#x200B;

&amp;#x200B;

    auto old = some_ref_ptr.load();
    while(! some_ref_ptr.compare_exchange_weak(old, modify(old)));
    
    2.
    
    auto old = some_ref_ptr.load();
    while(! some_ref_ptr.compare_exchange_weak(old, {old.ptr, old.count+1, old.mark});

3.

    auto old = some_ref_ptr.load();
    while(! some_ref_ptr.compare_exchange_weak(old, {old.ptr, old.count, true});

The first operation is a true compare and swap. `modify` needs to perform a non-trivial operation to reattach a pointer by traversing a linked list. `LOCK CMPXCHG16B` is the sledgehammer of an assembly instruction we need. This suffers from the CAS retry problem, and so loses the 'wait freedom' property, but that is the price we pay for performing essentially an arbitrary atomic read-modify-write operation.

The second and third operations can be performed with a `LOCK XADD`, and an `XCHG`. These operations are *wait free* and do not need to act on the whole 16 bytes in a tight loop. Furthermore, `XCHG`, is simply read-write, not read-modify-write.

Are C++20 atomic\_refs the answer? Will I be able to get an atomic reference to *both* a ref\_ptr object and a ref\_ptr::count? Is there a trick in C++17 I have missed?

Does inline assembly allow full control over memory ordering without having to write an entire project in assembly?

Almost all use cases for std::compare_exchange_weak suffer from this issue.
## [6][Watching for software inefficiencies with Valgrind](https://www.reddit.com/r/cpp/comments/exqg9n/watching_for_software_inefficiencies_with_valgrind/)
- url: https://kristerw.blogspot.com/2020/02/watching-for-software-inefficiencies.html
---

## [7][One flew over the matrix](https://www.reddit.com/r/cpp/comments/ey6upj/one_flew_over_the_matrix/)
- url: https://github.com/hosseinmoein/Matrix
---

## [8][C++20 Concepts demo (using simple linear algebra example)](https://www.reddit.com/r/cpp/comments/ey6qki/c20_concepts_demo_using_simple_linear_algebra/)
- url: https://youtu.be/B_KjoLid5gw
---

## [9][Variadic operators informal proposal](https://www.reddit.com/r/cpp/comments/ey6hdk/variadic_operators_informal_proposal/)
- url: https://www.reddit.com/r/cpp/comments/ey6hdk/variadic_operators_informal_proposal/
---
Hi, this idea popped up in my mind but I’m not sure whether it can be widely useful except of several cases below.

Imagine that we have a function for strings concatenation

    std::string concat(const std::string&amp; a, const std::string&amp; b);

At some point we understand that we can concat more than two strings, what do we usually do? Make it variadic:

    std::string concat(T&amp;&amp;… strings)
    {
    	//we can calculate length of a result string and allocate only once,
    	//eliminate temporaries, etc.
    }

In general it’s pretty easy for function to extend its behavior to handle multiple arguments. Operators differ from function only with semantics, `std::string s = s1 + s2 + s3;` means the same as `std::string s = concat(s1, s2, s3);` but looks nicer.

My idea is to introduce *variadic operators* that allows compiler to treat expressions in form of `a op b op c …` as `op(a, b, c, …)` or `a.op(b, c, ...)`. Allowed operators are all ones for fold expression plus operator\[\]. If expression contains another operators then only normal form is allowed:

    x = a + b + c;	 // calls operator+(a, b, c); or a::operator+(b, c);
    x = a + b + c*d; // calls normal operators

As I understand it right now rules for overload resolution shouldn’t be changed, it’s mostly semantics.

Where it might be useful:

* multidimensional array operator\[\], e.g. md\[1\]\[2\]...
* json-like objects operator\[\]
* string concatenation
* I/O-like operations with operator&gt;&gt;/&lt;&lt;
* pipes chain?
* cases when we can handle same multiple operations in optimized way

What do you think, potential problems, usefulness?
## [10][Why doesn't this auto cast to shared_ptr&lt;vector&lt;int&gt;&gt;?](https://www.reddit.com/r/cpp/comments/ey63gv/why_doesnt_this_auto_cast_to_shared_ptrvectorint/)
- url: https://www.reddit.com/r/cpp/comments/ey63gv/why_doesnt_this_auto_cast_to_shared_ptrvectorint/
---
```cpp
#include &lt;iostream&gt;
#include &lt;memory&gt;
#include &lt;typeinfo&gt;
#include &lt;vector&gt;

using int_vector = std::vector&lt;int&gt;;
using int_vector_ptr = std::shared_ptr&lt;int_vector&gt;;

void foo(int_vector_ptr v)
{
}

int main()
{
    auto v = std::make_shared&lt;int_vector&gt;;
    foo(v);
    return 0;
}

```

Here's the compile output:

```cpp
$ g++ main.cpp 
main.cpp: In function ‘int main()’:
main.cpp:14:9: error: could not convert ‘v’ from ‘std::shared_ptr&lt;std::vector&lt;int&gt; &gt; (*)()’ to ‘int_vector_ptr’ {aka ‘std::shared_ptr&lt;std::vector&lt;int&gt; &gt;’}
   14 |     foo(v);
      |         ^
      |         |
      |         std::shared_ptr&lt;std::vector&lt;int&gt; &gt; (*)()
```

What's going on here?
## [11][Does UB affect static_assert?](https://www.reddit.com/r/cpp/comments/ey5dk2/does_ub_affect_static_assert/)
- url: https://www.reddit.com/r/cpp/comments/ey5dk2/does_ub_affect_static_assert/
---
I'm writing a wrapper around builtin_clz. The results of builtin_clz is undefined if the argument is zero (https://gcc.gnu.org/onlinedocs/gcc/Other-Builtins.html). I add a static_assert that makes sure the result is zero and it compiles fine.

What is going on here? Does it actually return zero? Or is the compiler detecting the UB and just removing the static_assert?

Code:
https://godbolt.org/z/VKMu3m

A similar question was asked on SO, but nobody really answered the question
https://stackoverflow.com/questions/55031184/does-undefined-behavior-affect-static-assert
