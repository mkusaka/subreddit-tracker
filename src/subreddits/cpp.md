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
* **Don't** use URL shorteners. [reddiquette](https://www.reddit.com/wiki/reddiquette) forbids them because they're opaque to the spam filter.
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
## [2][Dry-comparisons: A C++ Library to Shorten Redundant If Statements](https://www.reddit.com/r/cpp/comments/ejkwqb/drycomparisons_a_c_library_to_shorten_redundant/)
- url: https://www.fluentcpp.com/2020/01/03/dry-comparisons-a-c-library-to-shorten-redundant-if-statements/
---

## [3][C++ Pattern Matching proposal](https://www.reddit.com/r/cpp/comments/ejg1yc/c_pattern_matching_proposal/)
- url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1371r1.pdf
---

## [4][Oh No! More Modern CMake - Deniz Bahadir - Meeting C++ 2019](https://www.reddit.com/r/cpp/comments/eje57u/oh_no_more_modern_cmake_deniz_bahadir_meeting_c/)
- url: https://www.youtube.com/watch?v=y9kSr5enrSk
---

## [5][A c++ based synchronous multi-peer music player](https://www.reddit.com/r/cpp/comments/ejs99e/a_c_based_synchronous_multipeer_music_player/)
- url: https://github.com/VedantParanjape/audio-streamer
---

## [6][The Debug heap that created bugs](https://www.reddit.com/r/cpp/comments/ej39ma/the_debug_heap_that_created_bugs/)
- url: http://lectem.github.io/windows/heap/appverifier/detours/2020/01/02/The-debug-heap-that-created-bugs.html
---

## [7][CppCast: C++ 2020 News](https://www.reddit.com/r/cpp/comments/ej98qn/cppcast_c_2020_news/)
- url: https://cppcast.com/cpp-2020-news/
---

## [8][The way to do error handling](https://www.reddit.com/r/cpp/comments/ejk5ui/the_way_to_do_error_handling/)
- url: https://www.reddit.com/r/cpp/comments/ejk5ui/the_way_to_do_error_handling/
---
[the link](https://github.com/Quaentor/status)

If you could be oh so kind, please tell me what you think.

This thing is a pretty general thing, and also highly composable. Because it is not only a set of classes, but a radical idea on how should be written code that deals with error handling, it is just impossible to describe it fully in a sane amount of words.

Considering the extra classes the repo contains, it already gives you several ways to handle the problem. It was designed to be extended for the needed functionality to handle the problem just how You want to. But with the extra classes you probably would already have all you need.

&amp;#x200B;

The repo contains:

* the main class `basic_status`, with which you can determine what type of error it was initialized with. It only stores the type. The class contains only one member of integral type. And all operations with it in assembly should take only one instruction.

&amp;#x200B;

    struct Error1{}; struct Error2{};

To write a function that returns a status you would write

    status&lt;Error1,Error2&gt;
    foo(){
        if(...) return Error1{};
        else    return Error2{};
        return {};
    }

To determine if it failed you would write

    if(auto s = foo(); !s)

To check if it returned a particular status you would write

    foo().is&lt;Error1&gt;()

If `Error2` was to be defined as `struct Error2 : Error1{};`, `foo().is&lt;Error1&gt;()` would return `true,` if `foo` returned `Error1` or `Error2`.

&amp;#x200B;

* the extra classes which all derive from `basic_status` :

`basic_holding_status` for storing the object of the current type. The object would be of type `Error1` or `Error2`, so you could initialize them in `foo` with some data.

`basic_acting_status` for giving a way to write a function that would be called when the status was initialized with some type(unsuccess state). You can choose when exactly the function is called -- in constructor of the status, in destructor, in destructor if the result was not checked.

I lied that they derive from `basic_status`. They derive from a template parameter, meaning you can combine them (an example can be seen in the [first link](https://wandbox.org/permlink/POagc0Sact7acFEo) in the repo, in file test\_acting)

&amp;#x200B;

The repo contains a couple of links, where you can see some examples.

The main reason why there aren't too many (though I think it's enough) is it's not mature enough yet to give examples on how to do things. And there could be two types of examples -- how to use given functionality, how to extend this functionality.

It could become mature enough with your suggestions and requests.

Maybe it's not that of a good idea, and it's not worth it. So tell me what you think.
## [9][Any active C++ alternative operating system projects underway, or does Redox OS written in Rust have zero competition?](https://www.reddit.com/r/cpp/comments/ej403h/any_active_c_alternative_operating_system/)
- url: https://www.reddit.com/r/cpp/comments/ej403h/any_active_c_alternative_operating_system/
---
Am aware of Haiku, and I guess it may be using C++, not sure if it's used for their kernel layers, though...

Redox OS is using Rust for everything. They've even been re-writing the popular UNIX flavored command line tools in Rust. However, Redox OS will allow for, say, device drivers to be written in other languages (because it's a micro kernel so drivers don't directly link to the kernel). And the OS API is POSIX so user programs can be written in anything.

MINIX3 is all C code (though does some microkernel architecture stuff too).

My interest is to identify new operating system implementations that can replace Linux as a user-facing operating system (for the long term - not the near term, of course).

I'd like to see a microkernel architecture where drivers aren't linked to the kernel. But to me the holy grail would be ability to use Linux device drivers. Kind of tough, right? Given the Linux device driver model. But I have some ideas of how to solve that.
## [10][Guaranteed copy elision for named return values](https://www.reddit.com/r/cpp/comments/ej3fvy/guaranteed_copy_elision_for_named_return_values/)
- url: https://www.reddit.com/r/cpp/comments/ej3fvy/guaranteed_copy_elision_for_named_return_values/
---
Two days ago, I've [posted](https://www.reddit.com/r/cpp/comments/ei83am/alias_expressions/) a proposal draft on "alias expressions". I've got two primary reactions:

1. Yay, less moves 👍
2. But it breaks some current rules and requres us to write extra weird syntax in order to achieve that 😕

I'm not sure, but it looks like I've come up with some promising wording to make copy elision guaranteed. This will apply to existing code, no changes required.

[The proposal link](https://gist.github.com/Anton3/594141354ff9625db0b85775799312c7)
## [11][phmap now provides btree containers in addition to the fast hash maps.](https://www.reddit.com/r/cpp/comments/eipcq1/phmap_now_provides_btree_containers_in_addition/)
- url: https://www.reddit.com/r/cpp/comments/eipcq1/phmap_now_provides_btree_containers_in_addition/
---
Phmap is a header only library which has provided very efficient (both time and space) hash maps for a while. Now it also provides similarly efficient btree alternatives for std::map and std::set, for when ordered containers are needed. 

See https://github.com/greg7mdp/parallel-hashmap.
