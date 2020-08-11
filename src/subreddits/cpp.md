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
## [2][AddressSanitizer for Windows: x64 and Debug Build Support | C++ Team Blog](https://www.reddit.com/r/cpp/comments/i7iydj/addresssanitizer_for_windows_x64_and_debug_build/)
- url: https://devblogs.microsoft.com/cppblog/asan-for-windows-x64-and-debug-build-support/
---

## [3][Concepts can’t do quantifiers](https://www.reddit.com/r/cpp/comments/i7mzga/concepts_cant_do_quantifiers/)
- url: https://quuxplusone.github.io/blog/2020/08/10/concepts-cant-do-quantifiers/
---

## [4][ModernCppStarter v0.13 released: m.css documentation and CMake formatting!](https://www.reddit.com/r/cpp/comments/i7oroj/moderncppstarter_v013_released_mcss_documentation/)
- url: https://github.com/TheLartians/ModernCppStarter
---

## [5][Build system, what’s your favorite?](https://www.reddit.com/r/cpp/comments/i7825h/build_system_whats_your_favorite/)
- url: https://www.reddit.com/r/cpp/comments/i7825h/build_system_whats_your_favorite/
---
Hey guys, I don’t know who saw the https://blog.jetbrains.com/clion/2020/06/dev-eco-cpp-2020/  where CMake is used by 53% of respondents. But would  like to see the opinions of this sub Reddit.

[View Poll](https://www.reddit.com/poll/i7825h)
## [6][TransCoder from Facebook Reserchers translates code from a C++ to another language! Check some examples at 3:10 in the video, or in the paper itself linked in the video description!](https://www.reddit.com/r/cpp/comments/i7q54w/transcoder_from_facebook_reserchers_translates/)
- url: https://www.youtube.com/watch?v=u6kM2lkrGQk
---

## [7][envy: Deserialize environment variables into type-safe structs](https://www.reddit.com/r/cpp/comments/i73c2a/envy_deserialize_environment_variables_into/)
- url: https://github.com/p-ranav/envy
---

## [8][Official C++ Logo is now under public domain](https://www.reddit.com/r/cpp/comments/i6u3xq/official_c_logo_is_now_under_public_domain/)
- url: https://www.reddit.com/r/cpp/comments/i6u3xq/official_c_logo_is_now_under_public_domain/
---
As you can know, Official C++ Logo ([isocpp/logos](https://github.com/isocpp/logos)) is not free since 18.04.2017 (※ [`a036ea6`](https://github.com/isocpp/logos/commit/a036ea65afa8b5f5ba7733f90d9aed8266eca6c1)).

Fortunatelly, I was recently cleaning up my hard drive space and found old C++ logo downloaded while C++ logo was under public domain, so in order to allow community use it freely, I am distributing it under public domain: [Benio101/cpp-logo](https://github.com/Benio101/cpp-logo) so…

**Everyone can use, modify and redistribute C++ Logo without any restrictions**.

⚠️ Note that in order to use unrestricted version of C++ Logo, you have to download it from my repository, not from isocpp once due to licensing.

Have a nice day.
## [9][Overwriting new and delete - sized delete never called](https://www.reddit.com/r/cpp/comments/i75qol/overwriting_new_and_delete_sized_delete_never/)
- url: https://www.reddit.com/r/cpp/comments/i75qol/overwriting_new_and_delete_sized_delete_never/
---
Hello Cpp!

&amp;#x200B;

My question is the following, i'm overwriting `operator new(size_t size)` and `operator delete(void *ptr, size_t sz)` to track my allocated memory.

The problem is, that the sized version of `operator delete` is never called, not matter the type of object i create previously. When i provice my own `operator delete(void *ptr)`, the version gets correctly called everytime. I tried with gcc and clang on ubuntu, maybe it was something the compiler did, but it seems the sized version of the operator do not work as intended. The documentation says they will be called when provided, but it seems just not true. Is there a way to force call these?

&amp;#x200B;

Is the only way around this to use the non-size version and do something like this for allocating, and read the size in the reverse way when deleting?

    void* operator new(size_t size) {
     size_t * ret = (size_t*) malloc(sizeof(size_t) + s);
     *ret = s;
     return (void*) &amp;ret[1];
    }
## [10][Alternatives to Visual Assist for C++ in Visual Studio](https://www.reddit.com/r/cpp/comments/i734hh/alternatives_to_visual_assist_for_c_in_visual/)
- url: https://www.reddit.com/r/cpp/comments/i734hh/alternatives_to_visual_assist_for_c_in_visual/
---
I'm looking for Visual Studio add-ons that improve intellisense and the refactoring tools like Visual Assist does. 

I search for something else than VA as VA has become a major annoyance lately. You cannot simply buy a key and download the software anymore but you need to register with some third party site and (at least I) had to go through a lengthy email conversation because I wanted to buy a private licence and my last name changed from years ago due to marriage.

I surely don't want to go through this again so I'm looking for a alternative to VA. What are you guys using?
## [11][How to Check String or String View Prefixes and Suffixes in C++20](https://www.reddit.com/r/cpp/comments/i71it6/how_to_check_string_or_string_view_prefixes_and/)
- url: https://www.bfilipek.com/2020/08/string-prefix-cpp20.html?m=1
---

