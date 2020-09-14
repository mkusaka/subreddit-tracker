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
## [2][Updates in Dragonbox](https://www.reddit.com/r/cpp/comments/ishdj9/updates_in_dragonbox/)
- url: https://www.reddit.com/r/cpp/comments/ishdj9/updates_in_dragonbox/
---
Hi all, Dragonbox ([https://github.com/jk-jeon/dragonbox](https://github.com/jk-jeon/dragonbox)) is a float-to-string conversion algorithm with:

1. Round-trip guarantee,
2. Shortest number of digits guarantee, and
3. Correct rounding guarantee,

which performs better than Ryu and also my previous algorithm Grisu-Exact.

I already have posted about Dragonbox two weeks ago here ([https://www.reddit.com/r/cpp/comments/ika6ml/dragonbox\_yet\_another\_floattostring\_conversion/](https://www.reddit.com/r/cpp/comments/ika6ml/dragonbox_yet_another_floattostring_conversion/)), but I shouldn't have done that at that time because it was quite incomplete, though I couldn't really realize. Sorry for spamming! :(

Now I think it is finally sorta completed, and perhaps this is a better timing to share it with you guys.

Here are some updates:

1. I finished writing a [paper](https://github.com/jk-jeon/dragonbox/blob/master/other_files/Dragonbox.pdf) for the algorithm. I hope it is readable to anyone who are interested in this topic, but I know my writing is horrible, so please give me some feedback if you ever have tried reading it.
2. Fixed some bugs in less-tested rounding modes. I believe they are all correct now.
3. Added a version using a compressed cache table for `double`'s. It uses only about 11% of static data table compared to the original one, at the cost of being about 20% slower. (The original one uses about 10KB, while the new one uses about 1.1KB.) It is but still way faster than the double-conversion library, for example.
4. Did more tests, including exhaustive test for `float`'s against a reverse (decimal-to-binary) conversion algorithm and confirmed successful round-trip of every single `float`. Now I'm pretty confident about the correctness.
5. Simplified the way of passing policy arguments.

Please check the repository ([https://github.com/jk-jeon/dragonbox](https://github.com/jk-jeon/dragonbox)) if you are interested, thanks.

I'm planning to work on things like Ryu-printf and string-to-float algorithms from now on, if time permits. I would appreciate it if you guys can share some interesting ideas or thoughts.
## [3][bitflags v1.2.0 - single-header C++ library for easily managing set of auto-generated type-safe flags](https://www.reddit.com/r/cpp/comments/isekmx/bitflags_v120_singleheader_c_library_for_easily/)
- url: https://www.reddit.com/r/cpp/comments/isekmx/bitflags_v120_singleheader_c_library_for_easily/
---
Hi everybody!

About a week ago, I posted a link to the initial implementation of my [bitflags](https://github.com/m-peko/bitflags) library. Here is the [link](https://www.reddit.com/r/cpp/comments/im9opm/bitflags_c17_singleheader_library_for_easily/) to that post which gave me an inspiration to other things that would be good to have implemented in the library.

Since then I have implemented few things and now, you can specify your flags as simple as:

    BEGIN_BITFLAGS(Flags)
        FLAG(none)
        FLAG(flag_a)
        FLAG(flag_b)
        FLAG(flag_c)
    END_BITFLAGS(Flags)

It looks cool, right? No more specifying flag values and no more specifying appropriate type for flags :)

Please take a look at my library on [Github](https://github.com/m-peko/bitflags) (don't forget to give a star if you like it :) ) and let me know your thoughts!

Thanks in advance!
## [4][libclang-static-build: Static Builds Of Libclang](https://www.reddit.com/r/cpp/comments/iskit5/libclangstaticbuild_static_builds_of_libclang/)
- url: https://github.com/deech/libclang-static-build
---

## [5][Recursive Lambdas in C++](https://www.reddit.com/r/cpp/comments/irupel/recursive_lambdas_in_c/)
- url: https://artificial-mind.net/blog/2020/09/12/recursive-lambdas
---

## [6][How I Improved My Legacy C++ Project with PVS-Studio](https://www.reddit.com/r/cpp/comments/ish6bd/how_i_improved_my_legacy_c_project_with_pvsstudio/)
- url: https://www.bfilipek.com/2020/09/pvs-studio-checking.html
---

## [7][Any way to "merge" a const and non-const getter?](https://www.reddit.com/r/cpp/comments/irva9b/any_way_to_merge_a_const_and_nonconst_getter/)
- url: https://www.reddit.com/r/cpp/comments/irva9b/any_way_to_merge_a_const_and_nonconst_getter/
---
Like, now i have it setup like this:

    MyClass&amp; Get(int index) 
    {
    return ...;
    }
    
    const MyClass&amp; Get(int index) const
    {
    return...;
    }

Is there any way to merge the logic? Like something like this:

    MyClass&amp; Get(int index) 
    {
        return Get(index); // Explicitly call the const variant from below and return as non-const
    }
    
    const MyClass&amp; Get(int index) const
    {
    return...;
    }

Because right now i basically have all the logic from these getters duplicated.. which is shit
## [8][Should I use make or python to build my source?](https://www.reddit.com/r/cpp/comments/ise4he/should_i_use_make_or_python_to_build_my_source/)
- url: https://www.reddit.com/r/cpp/comments/ise4he/should_i_use_make_or_python_to_build_my_source/
---
I have 3 different builds, release/optimize, debug1, debug2 (which is even slower), debug3 (like debug1 but changes some dependencies so valgrind will work correctly)

1. I have no idea how to tell make how to compile my code. For example how do I say all of these builds should have their object files compiled into different folders? Looking at make 5 [here](https://www.cs.colby.edu/maxwell/courses/tutorials/maketutor/) and reading [this](https://felixcrux.com/blog/creating-basic-makefile) didn't help.
2. Next issue I have is validating if the test pass/failed. If I can't write a one line bash script to use with make I'd need to put it into a python file? Almost feels like I should just do everything in the python file
3. Last is if I make a git hook I'll need to use return codes to say if it passed/failed (unless I remember wrong?). Can I do that in make?

I haven't tried doing all of this in python but have tried in bash and hated it. So I was wondering if I should use python, if there's going to be any pain points if I try to do it in python, if make can do what I need and if there's another tool I should look into?
## [9][The Most Popular Programming Languages - 1965/2020](https://www.reddit.com/r/cpp/comments/irfs7m/the_most_popular_programming_languages_19652020/)
- url: https://youtu.be/UNSoPa-XQN0
---

## [10][The Design of C++ , lecture by Bjarne Stroustrup](https://www.reddit.com/r/cpp/comments/irissi/the_design_of_c_lecture_by_bjarne_stroustrup/)
- url: https://youtu.be/69edOm889V4
---

## [11][CppCon: Classic STL Class Announcement](https://www.reddit.com/r/cpp/comments/irlm0w/cppcon_classic_stl_class_announcement/)
- url: https://quuxplusone.github.io/blog/2020/09/12/classic-stl-at-cppcon-2020/
---

