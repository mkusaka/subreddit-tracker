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
## [2][Tobias Hieta: Compiling C++ is slow - let's go faster](https://www.reddit.com/r/cpp/comments/evrbmv/tobias_hieta_compiling_c_is_slow_lets_go_faster/)
- url: https://youtu.be/X4pyOtawqjg
---

## [3][C++20 Reference Card + examples!](https://www.reddit.com/r/cpp/comments/evobp3/c20_reference_card_examples/)
- url: https://www.bfilipek.com/2020/01/cpp20refcard.html
---

## [4][Any thoughts on/experience with modules?](https://www.reddit.com/r/cpp/comments/evqluj/any_thoughts_onexperience_with_modules/)
- url: https://www.reddit.com/r/cpp/comments/evqluj/any_thoughts_onexperience_with_modules/
---
I would like to know peoples thoughts on modules, whether you think it solves the issues with headers and if they met your expectations. 

I'm very intrigued and my expectations are things like:
* Isolation of concerns. Only export what you want to export, but "globals" defined within a module are not accessible outside of it. As such, globals, defines, classes and functions with the same name in different modules shouldn't cause issues, unless they are both exported. I'm not yet sure if that's the case, but this is kind of a huge advantage of the way modules work in ES6. 
* Resolving the mess that are header imports and import order. Ideally, import order shouldn't matter because every module that needs something, should import that thing itself.
## [5][MSVC Backend Updates in Visual Studio 2019 Versions 16.3 and 16.4 | C++ Team Blog](https://www.reddit.com/r/cpp/comments/evh16n/msvc_backend_updates_in_visual_studio_2019/)
- url: https://devblogs.microsoft.com/cppblog/msvc-backend-updates-in-visual-studio-2019-versions-16-3-and-16-4/
---

## [6][Bitsery- binary serialization library v5.0.3 released](https://www.reddit.com/r/cpp/comments/evni9z/bitsery_binary_serialization_library_v503_released/)
- url: https://www.reddit.com/r/cpp/comments/evni9z/bitsery_binary_serialization_library_v503_released/
---
There are several key features that distinguish [bitsery](https://github.com/fraillt/bitsery) from all the other binary serialization libraries:

* is very fast
* serialized data is very small
* by default is safe when deserializing untrusted data
* provides A LOT of customization options via so-called "extensions" and bitsery config, few examples include:
   * fine-grained bit-level serialization control.
   * forward/backward compatibility for your types.
   * smart and raw pointers with allocators support and customizable runtime polymorphism.

I would like to provide you with an example of what bitsery can do for you.

Out of the box, bitsery is already very fast, but you can provide a non-resizable buffer for serialization, you can disable error checking during deserialization if input buffer is trusted, you can use bit-packing "extensions" to compress the data. What you get is at least +2x size and +2x speed (+4x with GCC) improvements compared to cereal!

More interesting results can be found [here](https://github.com/fraillt/cpp_serializers_benchmark).
The battle between GCC and Clang is very interesting.
GCC can do amazing things with bitsery :)
## [7][Organizing modules in C++20: ideas for conventions for module names, namespaces and module evolution.](https://www.reddit.com/r/cpp/comments/evp7kf/organizing_modules_in_c20_ideas_for_conventions/)
- url: https://www.reddit.com/r/cpp/comments/evp7kf/organizing_modules_in_c20_ideas_for_conventions/
---
Hello everyone,

I am trying to figure out how to organize a program with C++20 modules and settle on something that will make sense. My priorities would be:

    - Be reasonably ABI-Stable when evolving the code
    - Do not clash with others
       - But do::not::use::many::nested::namespaces. 2 levels of namespacing enough.
    
So here is my first try, tell me what you think:

Module-interface unit:
```
export module mycompany.domain.product:
export module mycompany.domain.product:model;
export module mycompany.domain.product:view;
...
``` 

Module partitions:

``` 
export module mycomany.domain.product:model;

namespace product::model {

export class MyModel {
...
};

}
```

I did not code this idea yet, since this is existing software and will need to take more things into account, but what I am wondering is how:

1. If the module grows a lot, let us say, I want to move mycompany.domain.product:models to its own module mycompany.domain.product.models, what happens to the ABI for external symbols? I guess it will be broken?
2. What happens if someone chooses the same namespace as mine from an outside module? I guess that the overload sets are merged, as usual.

These questions matter from an organizational point of view, since I have a monorepo that maybe later could be split in additional modularized repos.
## [8][Requires-expression | Andrzej's C++ blog](https://www.reddit.com/r/cpp/comments/evfbaa/requiresexpression_andrzejs_c_blog/)
- url: https://akrzemi1.wordpress.com/2020/01/29/requires-expression/
---

## [9][About "Qt Offering Changes 2020"](https://www.reddit.com/r/cpp/comments/evb5v1/about_qt_offering_changes_2020/)
- url: https://valdyas.org/fading/software/about-qt-offering-changes-2020/
---

## [10][Light and Powerful modern web-framework for C++](https://www.reddit.com/r/cpp/comments/ev2obp/light_and_powerful_modern_webframework_for_c/)
- url: https://oatpp.io/
---

## [11][single source as header &amp; source](https://www.reddit.com/r/cpp/comments/evhzqb/single_source_as_header_source/)
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
