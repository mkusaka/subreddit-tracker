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
## [2][C++ Objects, their lifetimes and pointer to objects](https://www.reddit.com/r/cpp/comments/epgxj6/c_objects_their_lifetimes_and_pointer_to_objects/)
- url: https://blog.panicsoftware.com/objects-their-lifetimes-and-pointers/
---

## [3][I built a site to search C++ libraries based on the awesome C++ github list, with sort, filtering, and extra meta information for each library](https://www.reddit.com/r/cpp/comments/epikzx/i_built_a_site_to_search_c_libraries_based_on_the/)
- url: https://www.reddit.com/r/cpp/comments/epikzx/i_built_a_site_to_search_c_libraries_based_on_the/
---
 [https://lucidindex.com/cpp](https://lucidindex.com/cpp)
## [4][Preferred solution for making projects work on anyone's machine?](https://www.reddit.com/r/cpp/comments/epdgn7/preferred_solution_for_making_projects_work_on/)
- url: https://www.reddit.com/r/cpp/comments/epdgn7/preferred_solution_for_making_projects_work_on/
---
This is more a CMake / build system question than a C++ one, but I just wanted to get some opinions.

I'm working on a few side projects and I have them on my Github, but it's fairly hard for other people to build them. In particular, the flow looks something like this:

&amp;#x200B;

&gt;1) Checkout the git repo  
&gt;  
&gt;2) Install dependencies manually (from README)  
&gt;  
&gt;3) Edit the CMake module files (Find&lt;Lib&gt;.cmake) to point to installed dependency on their PC  
&gt;  
&gt;4) Run CMake

What is the best way to skip step 3 (and preferably step 2)?

Should I just write a build script, as there's no package manager? How would it work across OSes (I want to support Mac and Linux at the very least - and then you have the various distros).

Thanks!
## [5][2 Lines Of Code and 3 C++17 Features - The overload Pattern (updated with C++20 notes)](https://www.reddit.com/r/cpp/comments/ep7gjm/2_lines_of_code_and_3_c17_features_the_overload/)
- url: https://www.bfilipek.com/2019/02/2lines3featuresoverload.html
---

## [6][What is this communities opinion on c++11 as a first language?](https://www.reddit.com/r/cpp/comments/ep1s2i/what_is_this_communities_opinion_on_c11_as_a/)
- url: https://www.reddit.com/r/cpp/comments/ep1s2i/what_is_this_communities_opinion_on_c11_as_a/
---
A bit of context, I have been teaching programming for about two years now and it’s all been through a c/c++ based curriculum that eventually also teaches web development via python later. As expected whenever I talk to other swe's the first question is often why not python. As I’m venturing into the new year and continuing I’m debating on the best route forward and would love to hear the thoughts of people in the community.
## [7][cl /PH or #pragma file_hash](https://www.reddit.com/r/cpp/comments/ep77ey/cl_ph_or_pragma_file_hash/)
- url: https://www.reddit.com/r/cpp/comments/ep77ey/cl_ph_or_pragma_file_hash/
---
Visual Studio 2019 Command Prompt:

    cl /?
    Microsoft (R) C/C++ Optimizing Compiler Version 19.24.28315 for x86
    Copyright (C) Microsoft Corporation.  All rights reserved.
    ...
                                  -PREPROCESSOR-
    ...
    /PH generate #pragma file_hash when preprocessing

I couldn't find anything about either `/PH` or `#pragma file_hash` online, but it seems that it calculates the file hash in one or more formats, MD5 by default:

    cl 1.cpp /E /PH
    1.cpp
    #line 1 "1.cpp"
    #pragma file_hash("1.cpp", MD5="ecaa88f7fa0bf610a5a26cf545dcd3aa")

MSVC guys, what exactly does this option do? Can it be used to speed up our builds, something ccache-like? Or it's just some internal stuff, accidentally leaked into the help page?
## [8][naïve function call with named params in c++20 (designated init + macros)](https://www.reddit.com/r/cpp/comments/ep2dd5/naïve_function_call_with_named_params_in_c20/)
- url: https://www.reddit.com/r/cpp/comments/ep2dd5/naïve_function_call_with_named_params_in_c20/
---
[https://godbolt.org/z/LY3oS8](https://godbolt.org/z/LY3oS8)

shame I cannot find a way to make the syntax prettier.
## [9][Great examples of C++ repos for learning?](https://www.reddit.com/r/cpp/comments/epeyzz/great_examples_of_c_repos_for_learning/)
- url: https://www.reddit.com/r/cpp/comments/epeyzz/great_examples_of_c_repos_for_learning/
---
My Cpp skills are rusty, and I'm trying to brush up on the language as it is best used in 2020.

Are there any github repos you could recommend?
## [10][HPX 1.4.0 released](https://www.reddit.com/r/cpp/comments/ep3e47/hpx_140_released/)
- url: https://stellar-group.org/2020/01/hpx-1-4-0-released/
---

## [11][potential pitfall of abbreviated templates and generic lambdas on default arguments?](https://www.reddit.com/r/cpp/comments/ep3uep/potential_pitfall_of_abbreviated_templates_and/)
- url: https://www.reddit.com/r/cpp/comments/ep3uep/potential_pitfall_of_abbreviated_templates_and/
---
see: [https://godbolt.org/z/yGIypM](https://godbolt.org/z/yGIypM)

I think we might want to change the wording for `auto` parameters when a default argument is provided in the next C++ version to fix this corner case.

&amp;#x200B;

while some of you might wonder what's the use of `auto` parameters with default arguments, here I give you another less trivial example: [https://godbolt.org/z/5uoak4](https://godbolt.org/z/5uoak4)

here we have a function `aggregate()` that takes a list of values and somehow merges them into one, by default it merges values by summation (operator+) which is very intuitive, but you can also specify your own merging function to calculate the product or the mathematical expectation of the values.

since the type of the merging function is totally unknown, it could be anything callable, it might even be stateful like `dot` which we use to calculate the expectation, we simply declare the parameter as `auto`, however it should be addition by default unless the user specified a custom merging function, and here, we have an `auto` parameter, and it makes sense for it to have a default argument.
