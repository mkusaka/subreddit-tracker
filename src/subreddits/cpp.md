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
## [2][Create a fully functioning command line interface with 1.5 lines of code (wo. include statement)](https://www.reddit.com/r/cpp/comments/i5yo1g/create_a_fully_functioning_command_line_interface/)
- url: https://github.com/kongaskristjan/fire-hpp
---

## [3][a + b: how hard can it be?](https://www.reddit.com/r/cpp/comments/i5nmf7/a_b_how_hard_can_it_be/)
- url: https://www.reddit.com/r/cpp/comments/i5nmf7/a_b_how_hard_can_it_be/
---
Given...

    int a = 0, b = 0;
    std::cin &gt;&gt; a &gt;&gt; b;

What’s your preferred method of a safe, no UB addition of...

    a + b

Please show your workings.
## [4][What Is The Minimal Set Of Optimizations Needed For Zero-Cost Abstraction?](https://www.reddit.com/r/cpp/comments/i5il0j/what_is_the_minimal_set_of_optimizations_needed/)
- url: https://robert.ocallahan.org/2020/08/what-is-minimal-set-of-optimizations.html
---

## [5][The “array size constant” antipattern](https://www.reddit.com/r/cpp/comments/i5bl28/the_array_size_constant_antipattern/)
- url: https://quuxplusone.github.io/blog/2020/08/06/array-size/
---

## [6][Does virtual constexpr function make sense?](https://www.reddit.com/r/cpp/comments/i5xr9v/does_virtual_constexpr_function_make_sense/)
- url: https://www.reddit.com/r/cpp/comments/i5xr9v/does_virtual_constexpr_function_make_sense/
---
I was just thinking whether it makes sense to have virtual constexpr functions?

I understand, that `virtual` is a "runtime keyword" and the `constexpr` is a "compile-time keyword", but in some context, I think it would make the compiler to generate more optimized code?

For example, let's say I have an `Object` class, it provides a method `type`, and all the types of objects are listed in an enum

```
enum OBJ_TYPE
{
   String, Integer, Real
};

class Object
{
public:
   virtual OBJ_TYPE type() const = 0;
}
```

In this example, at compile-time, I definitely know what is the body and return values of the derived classes for this particular function. So, what do you think about this?
## [7][C++ Lambda Week: Some Tricks](https://www.reddit.com/r/cpp/comments/i5epa4/c_lambda_week_some_tricks/)
- url: https://www.bfilipek.com/2020/08/c-lambda-week-some-tricks.html
---

## [8][Range-v3 0.11.0 released: backports from C++20 and compile-time improvements](https://www.reddit.com/r/cpp/comments/i54w0o/rangev3_0110_released_backports_from_c20_and/)
- url: https://github.com/ericniebler/range-v3/releases/tag/0.11.0
---

## [9][ASCII Earth rotating in C++. I made this animation using projections and a bit of linear algebra.](https://www.reddit.com/r/cpp/comments/i4qtol/ascii_earth_rotating_in_c_i_made_this_animation/)
- url: https://youtu.be/wgdB9Dgb0Ro
---

## [10][Overload Journal #158 - August 2020](https://www.reddit.com/r/cpp/comments/i5ar4c/overload_journal_158_august_2020/)
- url: https://accu.org/index.php/journals/c413/
---

## [11][CppCast: Go](https://www.reddit.com/r/cpp/comments/i54jmn/cppcast_go/)
- url: https://cppcast.com/ian-taylor-go/
---

