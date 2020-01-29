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
## [2][MSVC Backend Updates in Visual Studio 2019 Versions 16.3 and 16.4 | C++ Team Blog](https://www.reddit.com/r/cpp/comments/evh16n/msvc_backend_updates_in_visual_studio_2019/)
- url: https://devblogs.microsoft.com/cppblog/msvc-backend-updates-in-visual-studio-2019-versions-16-3-and-16-4/
---

## [3][Requires-expression | Andrzej's C++ blog](https://www.reddit.com/r/cpp/comments/evfbaa/requiresexpression_andrzejs_c_blog/)
- url: https://akrzemi1.wordpress.com/2020/01/29/requires-expression/
---

## [4][About "Qt Offering Changes 2020"](https://www.reddit.com/r/cpp/comments/evb5v1/about_qt_offering_changes_2020/)
- url: https://valdyas.org/fading/software/about-qt-offering-changes-2020/
---

## [5][Light and Powerful modern web-framework for C++](https://www.reddit.com/r/cpp/comments/ev2obp/light_and_powerful_modern_webframework_for_c/)
- url: https://oatpp.io/
---

## [6][single source as header &amp; source](https://www.reddit.com/r/cpp/comments/evhzqb/single_source_as_header_source/)
- url: https://www.reddit.com/r/cpp/comments/evhzqb/single_source_as_header_source/
---
I was thinking about that for header only you sometimes need to do tricks when having a compilation unit would make it easier. Maybe next level is header &amp; source combined. I went over to [documentation of gcc](https://gcc.gnu.org/onlinedocs/cpp/Common-Predefined-Macros.html) and found \_\_INCLUDE\_LEVEL\_\_.

&amp;#x200B;

base.cpp

    #if __INCLUDE_LEVEL__
    
    int function();
    #else
    
    int function() {
        return 42;
    }
    #endif
    

main.cpp

    #include "base.cpp"
    
    #include &lt;iostream&gt;
    
    int main() {
        std::cout &lt;&lt; function() &lt;&lt; "\n";
    }

CMakeLists.txt

    cmake_minimum_required(VERSION 3.5)
    
    project(Combined VERSION 1.0.0 LANGUAGES CXX)
    add_executable(main base.cpp main.cpp)

I have compiled and successfully linked this. I think it's interesting that this is possible, and also interesting that this hasn't caught on. Whether it's a good idea or not I don't know as I haven't done this in practice. But after trying this proof of concept, I'm more tempted to do this. [Clang](https://clang.llvm.org/docs/LanguageExtensions.html) seems to support it too, but can't find anything similar to [MSVC](https://docs.microsoft.com/en-us/cpp/preprocessor/predefined-macros?view=vs-2019). What's nice is no extra parameters needed to be invented. Obviously the downside that if implementation changes then its parents need to be compiled too. But for distributing something like a header only library. and if you only support Clang and gcc, then this works kind of nicely. Maybe something like this can have a new extension something like `.cpph`.
## [7][stdgpu 1.2.0 released!](https://www.reddit.com/r/cpp/comments/ev5ctx/stdgpu_120_released/)
- url: https://github.com/stotko/stdgpu/releases/tag/1.2.0
---

## [8][STL header token parsing benchmarks for libstdc++ 7, 8 and 9](https://www.reddit.com/r/cpp/comments/ev4e7g/stl_header_token_parsing_benchmarks_for_libstdc_7/)
- url: https://www.reddit.com/r/cpp/comments/ev4e7g/stl_header_token_parsing_benchmarks_for_libstdc_7/
---
To match [the results for VS2019 posted yesterday](https://www.reddit.com/r/cpp/comments/eumou7/stl_header_token_parsing_benchmarks_for_vs2017/) where we found that the VS2019 STL has become 9% lighter in preprocessor parsing time than the VS2017 STL, here are the same for libstdc++ 7, 8 and 9:

Graph of libstdc++-9: https://raw.githubusercontent.com/ned14/stl-header-heft/master/graphs/libstdc++-9.png

Comparative graph: https://raw.githubusercontent.com/ned14/stl-header-heft/master/graphs/libstdc++-history.png

Detailed notes: https://github.com/ned14/stl-header-heft/blob/master/Readme.libstdc++.md

Project github: https://github.com/ned14/stl-header-heft

You may remember from [the last test performed two years ago](https://www.reddit.com/r/cpp/comments/83rf8o/stl_header_token_parsing_benchmarks_for_libstdc_5/?st=jf0t6mob&amp;sh=817e7a37) that libstdc++ had engorged itself between libstdc++-5 and libstdc+-7 by **+25.7%**. I can report that libstdc++-9 has further engorged itself over libstdc++-7 by **+16%**. That's almost linear with time, which I find fascinating.

There is good news here though. Whilst the average and median has become much worse, all of the pathological individual headers in libstdc++-9 have been fixed e.g. `array` and `iterator` were dragging in lots of unneeded stuff which they no longer do. In this sense, both the libstdc++ and VS STLs have greatly improved in the past two years by not sub-including unnecessary headers. My congratulations to the libstdc++ authors and maintainers for this work!

If there was any one area where I think libstdc++ is particularly deficient, it is that any of its threading headers basically drag in the exact same stuff, most of which isn't actually needed by the individual header. MSVC's STL proves that this isn't necessary, `&lt;thread&gt;` is much lighter weight than any of the other threading headers on MSVC, but is the heaviest weight threading header on libstdc++, for example.

All that said, libstdc++ has *far* more very light weight headers than the MSVC STL e.g. `&lt;type_traits&gt;` on libstdc++ really is a class leader compared to the hefty `&lt;type_traits&gt;` on MSVC. MSVC *could* do much better here, a huge include for MSVC is `&lt;sal.h&gt;` which absolutely everything drags in, and is very low hanging fruit to prune in my opinion.
## [9][Ninja v1.10.0 released!](https://www.reddit.com/r/cpp/comments/ev2zng/ninja_v1100_released/)
- url: https://github.com/ninja-build/ninja/releases
---

## [10][OpenSSL client and server from scratch, part 3](https://www.reddit.com/r/cpp/comments/ev3vbq/openssl_client_and_server_from_scratch_part_3/)
- url: https://quuxplusone.github.io/blog/2020/01/26/openssl-part-3/
---

## [11][Type trait I made wont work in MSVC](https://www.reddit.com/r/cpp/comments/ev7r87/type_trait_i_made_wont_work_in_msvc/)
- url: https://www.reddit.com/r/cpp/comments/ev7r87/type_trait_i_made_wont_work_in_msvc/
---
So I made a type trait as a learning exercise that is supposed to check if the first type of the template parameters is the same as any of the other types in the template parameter pack.

    //Checks if type T is any of the types provided as ARGS...
    //I could have done this using a much simpler function template but I wanted the interface to be similar to stl type traits
    namespace TRAIT_TYPE_HELPERS
    {
    	template &lt;typename T, typename = void, typename ... ARGS&gt;
    	struct is_one_of_helper : std::false_type {};
    
    	template &lt;typename T, typename ... ARGS&gt;
    	struct is_one_of_helper &lt;T, std::enable_if_t&lt;(std::is_same_v&lt;T, ARGS&gt; || ...)&gt;, ARGS...&gt; : std::true_type {};
    }
    
    //Main type trait
    template &lt;typename T, typename ... ARGS&gt;
    using is_one_of = TRAIT_TYPE_HELPERS::is_one_of_helper&lt;T, void, ARGS...&gt;;
    
    //convenience template
    template &lt;typename T, typename ... ARGS&gt;
    constexpr bool is_one_of_v = is_one_of&lt;T, ARGS...&gt;::value;
    
    int main()
    {
    	std::cout &lt;&lt; is_one_of&lt;int, char, float, double, int&gt;::value;
    }

The above test code prints out zero in MSVC (Visual Studio 16.4.3) but it prints out 1 in GCC 7.4.0.
When I change the SFINAE line from std::is_same_v&lt;T, ARGS&gt; to std::is_same&lt;T, ARGS&gt;::value, it prints out 1 on both platforms.

I was just wondering if this is a bug in MSVC or maybe them not implementing the standard fully yet or just me incorrectly making assumptions on what should work and what shouldn't.

Discussion is appreciated.
