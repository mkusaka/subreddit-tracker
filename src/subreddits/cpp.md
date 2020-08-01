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
## [2][I've made a CMake code formatter and I would appreciate some feedback from user POV](https://www.reddit.com/r/cpp/comments/i1rn6t/ive_made_a_cmake_code_formatter_and_i_would/)
- url: https://www.reddit.com/r/cpp/comments/i1rn6t/ive_made_a_cmake_code_formatter_and_i_would/
---
I've made [gersemi](https://github.com/BlankSpruce/gersemi) with these two principles in mind:

- it should "just work"
- it shouldn't require much attention from a user

The second principle leads to opinionated style of formatting code. I took some inspiration from how `black` does it for Python and this talk by Kevlin Henney where he mentions a bit about how clear visual form might improve working with the code: [ITT 2016 - Kevlin Henney - Seven Ineffective Coding Habits of Many Programmers](https://youtu.be/ZsHMHukIlJY?t=978). If you are familiar with [cmake_format](https://github.com/cheshirekow/cmake_format) then it's quite an opposite approach to full configurability provided by cmake_format. Since both project are about formatting CMake code this comparison naturally arises. I'd certainly appreciate some feedback from potential users for this tool and C++ community seems like a good starting point for that.

To try out `gersemi` you need python 3.7+ and just invoke `pip install gersemi`.
## [3][I made this ASCII renderer in C++ (I'm planning on improving it soon)](https://www.reddit.com/r/cpp/comments/i1bbfq/i_made_this_ascii_renderer_in_c_im_planning_on/)
- url: https://youtu.be/I-2UK4tFk0U
---

## [4][C++ on Sea video - "These two ints have different types. A Data Oriented Design story" - Richard Fabian](https://www.reddit.com/r/cpp/comments/i1s1c9/c_on_sea_video_these_two_ints_have_different/)
- url: https://www.youtube.com/watch?v=fv43GuesjuM
---

## [5][TIL: You can use initializer lists directly in a range-based for loop](https://www.reddit.com/r/cpp/comments/i15o5k/til_you_can_use_initializer_lists_directly_in_a/)
- url: https://www.reddit.com/r/cpp/comments/i15o5k/til_you_can_use_initializer_lists_directly_in_a/
---
Maybe everybody else already knew this but for me this was a revelation:

    for (int64_t entity_id: {2, 1460, 3525, 5177}) {
      std::cout &lt;&lt; entity_id &lt;&lt; std::endl;
    }
## [6][Source files ending in a backslash](https://www.reddit.com/r/cpp/comments/i1rwsy/source_files_ending_in_a_backslash/)
- url: https://www.reddit.com/r/cpp/comments/i1rwsy/source_files_ending_in_a_backslash/
---
&gt;A source file that is not empty and that does not end in a new-line character, or that ends in a new-line character immediately preceded by a backslash character before any such splicing takes place, shall be processed as if an additional new-line character were appended to the file[.](https://eel.is/c++draft/lex.phases#1.2.sentence-4)

A file ending in a backslash does not end in a new-line character, therefore new-line should be appended. Is that before or after line splicing? At the end of phase 2, should the file end with [backslash, new-line] or should the original backslash be removed by splicing?
## [7][Conan 2.0 announcement](https://www.reddit.com/r/cpp/comments/i13i7m/conan_20_announcement/)
- url: https://blog.conan.io/2020/07/28/Launching-Conan-2.0-Tribe.html
---

## [8][Is there any reason to use old standard algorithms over constrained ones from std::ranges::?](https://www.reddit.com/r/cpp/comments/i1b069/is_there_any_reason_to_use_old_standard/)
- url: https://www.reddit.com/r/cpp/comments/i1b069/is_there_any_reason_to_use_old_standard/
---
As I look over all the new additions to &lt;algorithm&gt;, it seems that all the standard algorithms from std:: are superseded by algorithms from std::ranges::.

Is that so? Is there any reason to use the old ones, when the new ones are provided?

I believe that might be the case only if the definition of new ones are more constrained than the ones we already had. I heard somewhere that some of the new ones might have more restrictions in terms of required ordering, but I can't back it up with any example.

If the old algorithms from std:: really became obsolete, are they going to be deprecated?
## [9][C vs C++ for making a programming language](https://www.reddit.com/r/cpp/comments/i1rjzi/c_vs_c_for_making_a_programming_language/)
- url: https://www.reddit.com/r/cpp/comments/i1rjzi/c_vs_c_for_making_a_programming_language/
---
Which language (out of C and C++) is better to make a compiled programming language?

[View Poll](https://www.reddit.com/poll/i1rjzi)
## [10][Conan 1.28 released](https://www.reddit.com/r/cpp/comments/i1709w/conan_128_released/)
- url: https://github.com/conan-io/conan/releases/tag/1.28.0
---

## [11][jwt-cpp: header only JWT library with generic JSON library support (pre-release)](https://www.reddit.com/r/cpp/comments/i0yvib/jwtcpp_header_only_jwt_library_with_generic_json/)
- url: https://www.reddit.com/r/cpp/comments/i0yvib/jwtcpp_header_only_jwt_library_with_generic_json/
---
# [JWT-CPP](https://github.com/Thalhammer/jwt-cpp/tree/v0.5.0-rc.0)

## New Release Candidate 0.5.0-rc.0

This new pre-release has a few key features:

- support for *any* JSON library
- support std::error_code workflow 

Beyond the additional features we have also been working hard to improve the quality so you (but mostly me) can rely on this code in mission critical applications. The doxygen API documentation is now hosted in GitHub Pages to improve readability. 

&gt; Hoping to get your feedback! Looking forward to your contributions

## [JSON Traits](https://github.com/Thalhammer/jwt-cpp/tree/v0.5.0-rc.0#providing-your-own-json-traits)
The challenge with a lot of the C++ JWT libraries (check out [jwt.io](https://jwt.io/) for other options) is they are fixed to a single JSON library. This forces many developers to rely on 2, 3 or more JSON library to build their application. It's a pain.

A JWT claim, is simply `jwt::basic_claim&lt;my_json_traits&gt;`, which allows the creation and decoding/verification to be generic to the underlying JSON library you are already using. Check out the [read me](https://github.com/Thalhammer/jwt-cpp/tree/v0.5.0-rc.0#providing-your-own-json-traits) for more information!

In order to maintain the API, there's still a `jwt::claim` which is the template specialization of `picojson_traits`. I am in the camp of changing this is a major release for the much more common `nlohmann/json` that everyone is familiar with.

&gt; What JSON libraries would you like to see supported?

## std::error_code overloads
Exception handling is expensive, for many reasons, and it's one of the reasons behind the addition of these new API overloads. You are no longer forced to do exception handling. If you application is based around error_codes, this offers you a very nice fit.

## What's next?
- Support more algorithms that are available and we are trying to expend support.
- Ensure the `json_traits` are sufficient to meet *most* JSON libraries 
- Testing, testing, testing.
