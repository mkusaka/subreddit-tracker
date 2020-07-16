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
## [2][SwedenCpp website turned into, kind of, C++ community service page](https://www.reddit.com/r/cpp/comments/hs53um/swedencpp_website_turned_into_kind_of_c_community/)
- url: https://www.reddit.com/r/cpp/comments/hs53um/swedencpp_website_turned_into_kind_of_c_community/
---
SwedenCpp used the meetup free corona time to improve the website

* What upcoming C++ meetups are scheduled around the world?  
visit [https://www.swedencpp.se/worldwide](https://www.swedencpp.se/worldwide)
* What new C++ blogs are available?  
visit [https://www.swedencpp.se/blogs](https://www.swedencpp.se/blogs)
* What new C++ videos have been uploaded?  
visit [https://www.swedencpp.se/videos](https://www.swedencpp.se/videos)
* What new C++ user group talks have been made public?  
visit [https://www.swedencpp.se/ugvideos](https://www.swedencpp.se/ugvideos)

Originally, only meetups for the Swedish meetup groups have been shown, but during the last weeks, several content collections have been added, and I think this can be a nice overview for a lot of people.

Please note that only info is shown that can be collected automated.  Therefore it shows less, but always up to date and nearly in real time (just some minutes delay), content to what is already available on the isocpp website or the meetingcpp website, or here on reddit.

If you think anything is missing, please inform the maintainer.
## [3][Initial Support For C++20 Ranges in MSVC](https://www.reddit.com/r/cpp/comments/hrq7rs/initial_support_for_c20_ranges_in_msvc/)
- url: https://devblogs.microsoft.com/cppblog/initial-support-for-c20-ranges/
---

## [4][I made a scripting language](https://www.reddit.com/r/cpp/comments/hrwehr/i_made_a_scripting_language/)
- url: /r/github/comments/hrq432/i_made_a_scripting_language_for_c_projects/
---

## [5][CMake 3.18 Released!](https://www.reddit.com/r/cpp/comments/hrpi71/cmake_318_released/)
- url: https://cmake.org/download/
---

## [6][The Next Steps for Single Ownership and RAII](https://www.reddit.com/r/cpp/comments/hryer9/the_next_steps_for_single_ownership_and_raii/)
- url: https://vale.dev/blog/raii-next-steps
---

## [7][8 years of Meeting C++](https://www.reddit.com/r/cpp/comments/hs70hi/8_years_of_meeting_c/)
- url: https://meetingcpp.com/blog/items/8-years-of-Meeting-Cpp.html
---

## [8][Is this a legal use of pointer casting? (strict aliasing rule question)](https://www.reddit.com/r/cpp/comments/hs6ut6/is_this_a_legal_use_of_pointer_casting_strict/)
- url: https://www.reddit.com/r/cpp/comments/hs6ut6/is_this_a_legal_use_of_pointer_casting_strict/
---
I wanted to write a generic way to get from `SomePodType` -&gt; `std::byte *` (for the purposes of hashing).  I'm just curious if the below code breaks C++ strict aliasing.  It's confusing because `void *` is "related" to all types. And `std::byte *` is allowed to alias any type.  But I'm wondering if by jumping from source -&gt; `void *` -&gt; `std::byte *`, I end up somehow breaking aliasing rules... and thus the optimizer may not realize that it can't reorder stuff or remove "dead code" that is not really dead, etc...

Here is the code snippet:

    template &lt;typename CharT, typename T, typename = std::enable_if_t&lt;sizeof(CharT) == 1&gt;&gt;
    inline constexpr const CharT *ptr_cast(const T *t) {
        // this trick gets past the "no reinterpret_cast in constexpr" rule.
        // TODO: figure out if this is an aliasing rule violation.
        return static_cast&lt;const CharT *&gt;(static_cast&lt;const void *&gt;(t));
    }
    template &lt;typename T, std::enable_if_t&lt;!std::is_pointer_v&lt;T&gt; &amp;&amp; std::is_pod_v&lt;T&gt;, int&gt; = 0&gt;
    inline constexpr const std::byte *to_std_byte_ptr(const T &amp;t) {
        return ptr_cast&lt;std::byte&gt;(&amp;t);
    }
    
Is `to_std_byte_ptr` above a safe way to get to the bytes of some POD type T (say a double), in order to hash them, let's say?

Oh.. and I want this to all be `constexpr` if possible (which is why I didn't straight up `reinterpret_cast` without the intermediate `void *` -- it needed to all be `constexpr`-compatible `static_cast`s).
## [9][Postpone function checking till usage place](https://www.reddit.com/r/cpp/comments/hs5pa1/postpone_function_checking_till_usage_place/)
- url: https://www.reddit.com/r/cpp/comments/hs5pa1/postpone_function_checking_till_usage_place/
---
Hello all

Just wanted to share the problem/solution and ask for suggestion.

# Problem

So, here is [code](https://godbolt.org/z/1KbzKv):

    class A;
    
    class B
    {
    public:
        B(A* a, int value): _a{a}, _value{value} {}
    
        // template&lt;typename=void&gt;
        inline int total_value() const
        {
            return _a-&gt;value() + _value;
                  // ^~ invalid use of incomplete type 'class A'
        }
    private:
        A* _a;
        int _value;
    };

Important I cannot include `A` to `B`, and I *do want* `total_value` be inlined.

And I basically wanted the compiler to postpone checking of function `total_value` to the place where it's used, when definition of `A` is available.

This can be achieved with templates, but simply putting `template&lt;typename=void&gt;` still produces warning, which in case of "treat warning as errors" is still not sufficient.

# Solution

[Here](https://godbolt.org/z/jre8v8) is possible solution which uses additional class as indirection:

    #include &lt;utility&gt;
    
    struct Indirection
    {
        template&lt;typename T&gt;
        static decltype(auto) delay(T&amp;&amp; a)
        {
            return std::forward&lt;T&gt;(a);
        }
    };
    
    class A;
    
    class B
    {
    public:
        B(A* a, int value): _a{a}, _value{value} {}
    
        template&lt;typename I=Indirection&gt;
        int total_value() const
        {
            return I::delay(_a)-&gt;value() + _value;
        }
    private:
        A* _a;
        int _value;
    };

Actual implementation should probably be hidden by some macros.

# Questions

1. I am wondering if some libraries (e.g. boost) or some compilers (e.g through custom attribute) already contain such utilities?
2. Should one try to implement this as a custom attribute which would tell the compiler to not check function *at all* till usage? Would this violate some C++ rules?
3. And the main question. Is this something which anyone would really want?
## [10][Is downcasting for callback/event information also bad?](https://www.reddit.com/r/cpp/comments/hs05lf/is_downcasting_for_callbackevent_information_also/)
- url: https://www.reddit.com/r/cpp/comments/hs05lf/is_downcasting_for_callbackevent_information_also/
---
People often suggest that downcasting is a sign of a poor design.

Originating from a C# background myself, it's not at all uncommon to downcast information passed along from callbacks and/or events.

Probably a well known example is the UI events. Take the `onClick` for example. This callback expects an `Object` and an `EventArgs` as parameters.

The `Object` can then be downcast to a `Button` object, and the `EventArgs` contains event-specific information such as which mouse button as used (as an example).

Is doing something similar considered bad in C++, or is this a valid use-case of downcasting?
## [11][I'm looking for a C++ formatting tool that will ONLY change leading whitespace (for indentation) and does absolutely NOTHING ELSE. The popular tools are all highly configurable but I can't seem to narrow any of them down to that one essential feature. Please advise.](https://www.reddit.com/r/cpp/comments/hrxph1/im_looking_for_a_c_formatting_tool_that_will_only/)
- url: https://www.reddit.com/r/cpp/comments/hrxph1/im_looking_for_a_c_formatting_tool_that_will_only/
---

