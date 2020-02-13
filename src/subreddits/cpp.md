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
## [2][Herb Sutter: Quantifying Accidental Complexity: An Empirical Look at Teaching and Using C++](https://www.reddit.com/r/cpp/comments/f2xq0l/herb_sutter_quantifying_accidental_complexity_an/)
- url: https://www.youtube.com/watch?v=qx22oxlQmKc
---

## [3][Modern C++ and Qt are ❤️](https://www.reddit.com/r/cpp/comments/f2wu98/modern_c_and_qt_are/)
- url: https://www.reddit.com/r/cpp/comments/f2wu98/modern_c_and_qt_are/
---
Just wanted to thank everybody who made this simplicity possible:

```cpp

std::optional&lt;QSize&gt; mudlet::getImageSize(const QString &amp;imageLocation)
{
    QImage image(imageLocation);

    if (image.isNull()) {
        return {};
    }

    return {image.size()};
}

// ...

if (auto size = mudlet::self()-&gt;getImageSize(imageLocation); size) {
    lua_pushnumber(L, size-&gt;width());
    lua_pushnumber(L, size-&gt;height());
    return 2;
} else {
    lua_pushnil(L);
    lua_pushfstring(L, "couldn't retrieve image size. Is the location '%s' correct?", imageLocation.toUtf8().constData());
    return 2;
}
```

This is amazing!
## [4][Improving Compilation Time of C/C++ Projects](https://www.reddit.com/r/cpp/comments/f2wsl4/improving_compilation_time_of_cc_projects/)
- url: https://interrupt.memfault.com/blog/improving-compilation-times-c-cpp-projects
---

## [5][My approach to compile-time reflection in C++](https://www.reddit.com/r/cpp/comments/f2xa28/my_approach_to_compiletime_reflection_in_c/)
- url: http://bluescreenofdoom.com/post/code/Reflection/
---

## [6][How to replace __FILE__ with source_location in a logging macro](https://www.reddit.com/r/cpp/comments/f2x461/how_to_replace_file_with_source_location_in_a/)
- url: https://quuxplusone.github.io/blog/2020/02/12/source-location/
---

## [7][Announcing TCMalloc](https://www.reddit.com/r/cpp/comments/f2yk94/announcing_tcmalloc/)
- url: https://abseil.io/blog/20200212-tcmalloc
---

## [8][Qt World Summit 2019 talk videos are online - KDAB](https://www.reddit.com/r/cpp/comments/f37t96/qt_world_summit_2019_talk_videos_are_online_kdab/)
- url: https://www.kdab.com/qt-world-summit-2019-talk-videos-are-online/
---

## [9][What's the best workflow for forking and iterating on external packages in vcpkg?](https://www.reddit.com/r/cpp/comments/f3317k/whats_the_best_workflow_for_forking_and_iterating/)
- url: https://www.reddit.com/r/cpp/comments/f3317k/whats_the_best_workflow_for_forking_and_iterating/
---
Let's say I have a program that depends on some commonly available 3rd party package that's available in vcpkg. Let's say I've configured the cmake build for my program to use the vcpkg cmake build system for finding everything.

Let's say I discover that I have a need to iterate on the code of an external package. The loop i want is:

* Compile and link my program, which links the vcpkg-provided locally-built external dependency. (No, it's not header-only.)
* Run it. Set and hit a breakpoint in the external dependency. Realize something needs to change.
* Edit the external source file?   (This is probably a great time to fork a separate version...what workflow?)
* ...ideally one button from here builds the external dependency (and anything else it depends on) incrementally, and my program also. 

Is such a thing possible? 

How do you do this with your local forks of external libraries?
## [10][Literal Classes as Non-type Template Parameters in C++20](https://www.reddit.com/r/cpp/comments/f2s4ut/literal_classes_as_nontype_template_parameters_in/)
- url: https://blog.keha.dev/posts/cpp20-class-as-non-type-template-param/
---

## [11][Combining ZeroMQ &amp; POSIX signals: Use ppoll to handle EINTR once and for all](https://www.reddit.com/r/cpp/comments/f2r1i6/combining_zeromq_posix_signals_use_ppoll_to/)
- url: https://blog.esciencecenter.nl/combining-zeromq-posix-signals-b754f6f29cd6
---

