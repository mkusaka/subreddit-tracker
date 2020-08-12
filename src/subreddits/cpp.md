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
## [2][No more plain old data](https://www.reddit.com/r/cpp/comments/i8a5xv/no_more_plain_old_data/)
- url: https://mariusbancila.ro/blog/2020/08/10/no-more-plain-old-data/
---

## [3][std::atomic_ref is awesome](https://www.reddit.com/r/cpp/comments/i8ckxr/stdatomic_ref_is_awesome/)
- url: https://www.reddit.com/r/cpp/comments/i8ckxr/stdatomic_ref_is_awesome/
---
I only recently discovered that gcc 10 supports std::atomic_ref and I’ve been playing around with it. I have to say, it might be in my top 5 favorite C++20 features. But it begs the question: before std::atomic_ref, was there a way to atomically modify a variable of non-atomic type in-place? I don’t really do lock-free programming in my every-day life, so I haven’t really ever had to think about it, but I’m curious. Also Jason Turner if you’re reading this please consider std::atomic_ref for a C++ Weekly episode, maybe there are other people who also slept on this one.
## [4][toml++ v2.1.0 released](https://www.reddit.com/r/cpp/comments/i7za0j/toml_v210_released/)
- url: https://marzer.github.io/tomlplusplus/
---

## [5][Why std::unique_lock still isn't [[nodiscard]]?](https://www.reddit.com/r/cpp/comments/i7v7yl/why_stdunique_lock_still_isnt_nodiscard/)
- url: https://www.reddit.com/r/cpp/comments/i7v7yl/why_stdunique_lock_still_isnt_nodiscard/
---
The following code: [https://godbolt.org/z/4ovWhv](https://godbolt.org/z/4ovWhv)

Is most likely a bug, if you find it somewhere in the code base.

Now, [p1771r1](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1771r1.pdf) introduced a  `[[nodiscard]]` for constructors. The change was  applied retroactively to previously published C++ standards. 

Shouldn't both unique\_lock and lock\_guard, as well as other RAII wrappers, be marked `[[nodiscard]]`?  

Is there an ongoing proposal to do that, and is there a chance for it to be retroactively applied? It seems like a bug that's way too easy to make, I did it myself a few years ago and later seen it in code reviews of other people. I believe it needs attention.
## [6][Time complexity of set in this case?](https://www.reddit.com/r/cpp/comments/i8a8rj/time_complexity_of_set_in_this_case/)
- url: https://www.reddit.com/r/cpp/comments/i8a8rj/time_complexity_of_set_in_this_case/
---
Consider this fragment of code:

    int main()
    { 
        string s1;
        getline(cin,s1);
        set&lt;char&gt; mySet;
        for(int i=0;i&lt;s1.length();i++)
            mySet.insert(s2[i]);
    }

What will be the time complexity of this code? I'm aware that the cost of insertion in a set in average case is O(log N). However when placed inside this for loop running through N characters, will the time complexity balloon up to O(N log N) or O(N) or will it be O(log N)?
## [7][AddressSanitizer for Windows: x64 and Debug Build Support | C++ Team Blog](https://www.reddit.com/r/cpp/comments/i7iydj/addresssanitizer_for_windows_x64_and_debug_build/)
- url: https://devblogs.microsoft.com/cppblog/asan-for-windows-x64-and-debug-build-support/
---

## [8][ModernCppStarter v0.13 released: m.css documentation and CMake formatting!](https://www.reddit.com/r/cpp/comments/i7oroj/moderncppstarter_v013_released_mcss_documentation/)
- url: https://github.com/TheLartians/ModernCppStarter
---

## [9][Any good C++ tutorials coming from C# and Python?](https://www.reddit.com/r/cpp/comments/i7walr/any_good_c_tutorials_coming_from_c_and_python/)
- url: https://www.reddit.com/r/cpp/comments/i7walr/any_good_c_tutorials_coming_from_c_and_python/
---
I am stepping up my learning after deciding what I want to do for a job. I make games in Unity and am at the point where coding doesn't stop me making things any more. I rarely need to look up documentation on the things I know, and I am confident that I could program most gameplay things.

I also use Python to automate things, and I can do stuff with that too, but not as well as C#. The problem with most tutorials on any language is that they tend to go over variables and functions, and then stop. I can already guess how to do those in most languages, so I don't want anything like that. Can anyone recommend some good tutorials that are for people who program already, but not in C++?  


I'm running G++, Windows 10 and my code editor of choice is Visual Studio Code.
## [10][Concepts can’t do quantifiers](https://www.reddit.com/r/cpp/comments/i7mzga/concepts_cant_do_quantifiers/)
- url: https://quuxplusone.github.io/blog/2020/08/10/concepts-cant-do-quantifiers/
---

## [11][Build system, what’s your favorite?](https://www.reddit.com/r/cpp/comments/i7825h/build_system_whats_your_favorite/)
- url: https://www.reddit.com/r/cpp/comments/i7825h/build_system_whats_your_favorite/
---
Hey guys, I don’t know who saw the https://blog.jetbrains.com/clion/2020/06/dev-eco-cpp-2020/  where CMake is used by 53% of respondents. But would  like to see the opinions of this sub Reddit.

[View Poll](https://www.reddit.com/poll/i7825h)
