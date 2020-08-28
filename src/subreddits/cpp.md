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
## [2][A line by line explanation on how to use a non-CMake library in your CMake project](https://www.reddit.com/r/cpp/comments/ii1o0y/a_line_by_line_explanation_on_how_to_use_a/)
- url: https://dominikberner.ch/cmake-find-library/
---

## [3][std::array indexed by any numeric-like type](https://www.reddit.com/r/cpp/comments/ii2xwt/stdarray_indexed_by_any_numericlike_type/)
- url: https://www.reddit.com/r/cpp/comments/ii2xwt/stdarray_indexed_by_any_numericlike_type/
---
Hi folks!

Is that popular use case to index `std::array` with enum values?

    enum Animal
    {
        Pig, Dog, Cat,
        SIZE
    };
    
    std::array&lt;Farm, Animal::SIZE&gt; farms;

For me it seems to be, but it would be great to leverage type-safety of `enum class`. So I come up with an idea:

    template&lt;typename T, auto Size&gt;
    class array
    {
    public:
        using Index = decltype(Size);
    
        T&amp; operator[](const Index i)
        {
            return _data[static_cast&lt;std::size_t&gt;(i)];
        }
    
    private:
        std::array&lt;T, static_cast&lt;std::size_t&gt;(Size)&gt; _data;
    };

[Godbolt link to play.](https://godbolt.org/z/3n6ndM)

It actually can work with anything convertible to `size_t` and it is just natural extension of `std::array`, so I would expect this to implemented already somewhere.

&amp;#x200B;

1. Would do you think about this?
2. Is there prod-ready implementation of this?
3. Looks like `std::array` could support this out of the box, just replacing `size_t` with `auto` and adding `static_cast` everywhere. May this be pushed to std lib?
## [4][Webinar: Write cleaner, safer, modern C++ with SonarQube](https://www.reddit.com/r/cpp/comments/ii5vgw/webinar_write_cleaner_safer_modern_c_with/)
- url: https://www.reddit.com/r/cpp/comments/ii5vgw/webinar_write_cleaner_safer_modern_c_with/
---
[https://sonarsource.zoom.us/webinar/register/1815977846611/WN\_6iMr512HQHa2CTvcdfv1vw](https://sonarsource.zoom.us/webinar/register/1815977846611/WN_6iMr512HQHa2CTvcdfv1vw)

&gt;As a C++ Developer, you know that writing clean, secure, modern C++ code is important for you and your users. At SonarSource, we know that only developers can truly impact Code Quality and Security, so we put the power in your hands.  
&gt;  
&gt;SonarQube makes C++ development easier with static code analysis that's powerful, fast, and accurate - right out of the box. Analysis is easy to integrate into your workflow and works with most common compilers, including many for embedded systems. Come see for yourself how you can make your C++ projects more reliable and secure.  
&gt;  
&gt;Join us on September 2nd at 10 am CDT / 3 pm GMT for a 30-minute webinar to learn more about:  
\- Code Quality &amp; Security for the individual: in-IDE and in PRs  
\- Code Quality &amp; Security for the team: in SonarQube  
\- What types of issues you can find on your C++ code: Bugs, Vulnerabilities, Security Hotspots, and Code Smells.  
\- How easy it is to get started - including a demo of SonarQube!

Even if you can't make it at that time, sign up and you'll get a link to the recording afterward.
## [5][Introducing vcperf /timetrace for C++ build time analysis](https://www.reddit.com/r/cpp/comments/ihnr1k/introducing_vcperf_timetrace_for_c_build_time/)
- url: https://devblogs.microsoft.com/cppblog/introducing-vcperf-timetrace-for-cpp-build-time-analysis/
---

## [6][concurrencpp v.0.0.4 - modern concurrency for C++](https://www.reddit.com/r/cpp/comments/ii3zef/concurrencpp_v004_modern_concurrency_for_c/)
- url: https://www.reddit.com/r/cpp/comments/ii3zef/concurrencpp_v004_modern_concurrency_for_c/
---
Hello reddit!

I just released version 0.0.4 of [concurrencpp](https://github.com/David-Haim/concurrencpp), a library for executors and coroutines.

The library is still very fresh and gradually matures more and more. 

There are still tons of features and optimizations that are scheduled for the future, so this library is far from being complete.

Suggestions, questions, reviews, and most importantly - stars, are greatly appreciated.

Let's make the way we deal with concurrency in C++ the best among all languages!
## [7][Can I start implementing modules in my code?](https://www.reddit.com/r/cpp/comments/ii7346/can_i_start_implementing_modules_in_my_code/)
- url: https://www.reddit.com/r/cpp/comments/ii7346/can_i_start_implementing_modules_in_my_code/
---
MSVC has partial support for them. I know the question is vague, but is anybody already using modules in their code, or everybody is waiting for full compilers support?
## [8][CppCast: Cross Platform Mobile Telephony](https://www.reddit.com/r/cpp/comments/ihxt9e/cppcast_cross_platform_mobile_telephony/)
- url: https://cppcast.com/telephony-dave-hagedorn/
---

## [9][Lyra 1.5 -- Create a full CLI parser in one statement, without globals or macros](https://www.reddit.com/r/cpp/comments/ihlaq0/lyra_15_create_a_full_cli_parser_in_one_statement/)
- url: https://www.reddit.com/r/cpp/comments/ihlaq0/lyra_15_create_a_full_cli_parser_in_one_statement/
---
Lyra ([https://bfgroup.github.io/Lyra/](https://bfgroup.github.io/Lyra/)) is a simple to use, composing, header only, command line arguments parser for C++ 11 and beyond. [Version 1.5](https://bfgroup.github.io/Lyra/lyra.html#_1_5) includes the time saving feature of a "[main](https://bfgroup.github.io/Lyra/lyra.html#_main)" utility to make creating simple CLIs with help quick and easy. For example:

    #include &lt;iostream&gt;
    #include &lt;lyra/lyra.hpp&gt;
    int main(int argc, const char ** argv)
    {
      return lyra::main()("-x", 0)("-y", 0)(argc, argv, [](lyra::main &amp; m)
      {
        std::cout &lt;&lt; int(m["-x"]) + int(m["-y"]) &lt;&lt; "\n";
        return 0;
      });
    }

Other notable new features for this release include:

* Direct support for [sub-commands](https://bfgroup.github.io/Lyra/lyra.html#_sub_commands).
* [Value](https://bfgroup.github.io/Lyra/lyra.html#lyra_val) holders.
* Argument [groups](https://bfgroup.github.io/Lyra/lyra.html#_argument_groups) to support alternate parsing, like sub-commands.
* The help output is, once again, nicely formatted following clang style help output.
* And there's even Cmake install support now.
## [10][Can we get MAC address using boost library?? if yes then how ??](https://www.reddit.com/r/cpp/comments/ii6bbs/can_we_get_mac_address_using_boost_library_if_yes/)
- url: https://www.reddit.com/r/cpp/comments/ii6bbs/can_we_get_mac_address_using_boost_library_if_yes/
---
 Can we get MAC address using boost library?? if yes then how ??
## [11][What's the status of P1729R0 Text Parsing?](https://www.reddit.com/r/cpp/comments/ihpc5u/whats_the_status_of_p1729r0_text_parsing/)
- url: https://www.reddit.com/r/cpp/comments/ihpc5u/whats_the_status_of_p1729r0_text_parsing/
---
 [http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1729r0.html](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1729r0.html)   


I would love to see a sscanf alternative in C++. Anyone knows the status of that proposal?
