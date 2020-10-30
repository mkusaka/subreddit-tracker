# cpp
## [1][C++ Jobs - Q4 2020](https://www.reddit.com/r/cpp/comments/j3qems/c_jobs_q4_2020/)
- url: https://www.reddit.com/r/cpp/comments/j3qems/c_jobs_q4_2020/
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

* [C++ Jobs - Q3 2020](https://www.reddit.com/r/cpp/comments/hjnaf2/c_jobs_q3_2020/)
## [2][Forbidden C++](https://www.reddit.com/r/cpp/comments/jks8f3/forbidden_c/)
- url: https://www.reddit.com/r/cpp/comments/jks8f3/forbidden_c/
---
I have Just watched this entertaining video on Forbidden C++.  Its aimed at beginners but thought i would post it here just for the intro alone!

[https://youtu.be/j0\_u26Vpb4w](https://youtu.be/j0_u26Vpb4w)

The video is just a small list but as we all know, with every newer release of C++ we get better ways of doing things.  Hence the darker side of C++ will continue to grow!
## [3][std::visit is everything wrong with modern C++](https://www.reddit.com/r/cpp/comments/jkiqkz/stdvisit_is_everything_wrong_with_modern_c/)
- url: https://bitbashing.io/std-visit.html
---

## [4][Strings in switch statements using constexp hashing](https://www.reddit.com/r/cpp/comments/jkw84k/strings_in_switch_statements_using_constexp/)
- url: https://www.reddit.com/r/cpp/comments/jkw84k/strings_in_switch_statements_using_constexp/
---
I have a question which might be dumb.

If I am not mistaken, Java has switch-case statements that support strings. Such a thing is not possible with plain c++ but there is a workaround if we use a constexp hash function for converting a string to a size\_t value.

So lets say we have a program like this:


    constexpr size_t hash(const char* str){
        const long long p = 131;
        const long long m = 4294967291; // 2^32 - 5, largest 32 bit prime
        long long total = 0;
        long long current_multiplier = 1;
        for (int i = 0; str[i] != '\0'; ++i){
            total = (total + current_multiplier * str[i]) % m;
            current_multiplier = (current_multiplier * p) % m;
        }
        return total;
    }

    int main() {
        std::string val;
        std::cin &gt;&gt; val;
    
        switch (hash(val.c_str())){
            case hash("monday"):
                std::cout &lt;&lt; "Have a nice monday" &lt;&lt; std::endl;
                break;
            case hash("tuesday"):
                std::cout &lt;&lt; "Have a nice tuesday" &lt;&lt; std::endl;
                break;
            case hash("wednesday"):
                std::cout &lt;&lt; "Have a nice wednesday" &lt;&lt; std::endl;
                break;
            case hash("thursday"):
                std::cout &lt;&lt; "Have a nice thursday" &lt;&lt; std::endl;
                break;
            case hash("friday"):
                std::cout &lt;&lt; "Have a nice friday" &lt;&lt; std::endl;
                break;
            default:
                std::cout &lt;&lt; "It is the weekend" &lt;&lt; std::endl;
        }
        return 0;
    }

Would that be something bad to do. I don't really know a use case for now, but am interested if this would be OKish to do.

Edit: The edits were for making the code look like it should
## [5][CppCast: ThinLTO](https://www.reddit.com/r/cpp/comments/jkp9sx/cppcast_thinlto/)
- url: https://cppcast.com/teresa-johnson-thinlto/
---

## [6][A tale of a build system bug](https://www.reddit.com/r/cpp/comments/jkc1bf/a_tale_of_a_build_system_bug/)
- url: https://gist.github.com/Som1Lse/2fbb0e22cb59a158bb8e09bf6f527f7e
---

## [7][Is C++ really that hard for beginners?](https://www.reddit.com/r/cpp/comments/jk6jcz/is_c_really_that_hard_for_beginners/)
- url: https://www.reddit.com/r/cpp/comments/jk6jcz/is_c_really_that_hard_for_beginners/
---
hello, i have always heard that "C++ is hard for beginner, manual memory management and stuff...", but is it really that hard after the changes c++ had?, i always avoided c++ (started with it and i got fustrated) but since c++ is the norm for game dev i kinda wanna use it (well maybe now or untill i learn java or something idk...)

and yes i know c++ doesn't have a garbage collection and it has pointers and stuff..

thanks.
## [8][Colony v6 released](https://www.reddit.com/r/cpp/comments/jkikj4/colony_v6_released/)
- url: https://www.reddit.com/r/cpp/comments/jkikj4/colony_v6_released/
---
More details [here](https://plflib.org/blog.htm#colonyv6),

but yeah it's done.

[https://github.com/mattreecebentley/plf\_colony](https://github.com/mattreecebentley/plf_colony)  


Main features/changes:  
\* actually fully-functional reserve()  
\* performance improvements to lots of stuff  
\* assign!  
\* Misc!
## [9][atomicDEX Desktop 0.3.0: A C++ 17 Open-Source decentralized cryptocurrency exchange/wallet](https://www.reddit.com/r/cpp/comments/jkiphk/atomicdex_desktop_030_a_c_17_opensource/)
- url: https://www.reddit.com/r/cpp/comments/jkiphk/atomicdex_desktop_030_a_c_17_opensource/
---
[https://github.com/KomodoPlatform/atomicDEX-Desktop/releases](https://github.com/KomodoPlatform/atomicDEX-Desktop/releases)

[https://github.com/KomodoPlatform/atomicDEX-Desktop/wiki](https://github.com/KomodoPlatform/atomicDEX-Desktop/wiki)

build with: boost, spdlog, entt, fmt, nlohmann\_json, cpprestsdk, taskflow, reproc, Qt, doctest

I'm also the author of:

[https://github.com/KomodoPlatform/antara-gaming-sdk](https://github.com/KomodoPlatform/antara-gaming-sdk)

which was previously: [https://github.com/Milerius/shiva](https://github.com/Milerius/shiva)

Discord link of the project: [https://discord.gg/FR96rQh](https://discord.gg/FR96rQh)

One of the challenges of the project was to mix the elegance of the boost libraries with the power of the Qt framework for the user interface, we always try to use the open-source C++ libraries which allow us to provide you with a coherent and safe project. , the dependencies are managed by Conan / VCPKG through our forks (for vcpkg)

[https://i.imgur.com/KN4i6XS.gif](https://i.imgur.com/KN4i6XS.gif)
## [10][vcpkg and github actions are a nightmare](https://www.reddit.com/r/cpp/comments/jkgapv/vcpkg_and_github_actions_are_a_nightmare/)
- url: https://www.reddit.com/r/cpp/comments/jkgapv/vcpkg_and_github_actions_are_a_nightmare/
---
The available tooling for doing cross platform development has improved quite a bit in the last few years, but sadly it's not quite there yet.

Let's start out with a simple task: add a github action to enable ci for a c++ project that depends on boost and qt5.

The workflow file would look something like this:

    name: CompileMyProject
    
    on: [push, pull_request]
    
    jobs:
      build:
        strategy:
          matrix:
            platform: [windows-latest, ubuntu-latest, macos-latest]
        runs-on: ${{ matrix.platform }}
        steps:
        - uses: actions/checkout@v2
          with:
            submodules: 'recursive'
        - name: Run vcpkg
          uses: lukka/run-vcpkg@v4
          with:
            vcpkgArguments: 'boost qt5'
            vcpkgDirectory: '${{ github.workspace }}/vcpkg'

As soon as you push the commit containing this workflow to the repo everything will appear to work perfectly. Github will even helpfully build the three platforms in parallel for you (even though each job is limited to 2 cores so expect this dependency build to take a least 2 hours).

Don't worry about that 2 hour build time for your dependencies because the run-vcpkg script is smart enough to cache the compiled libraries...

...unless of course something goes wrong. On any job. At any stage of the build process.

You see, github actions do support artifact caching, but it's all-or-nothing.

If vcpkg wants to build 100 packages, and you get an error on the 99th package, well sucks to be you because the 98 packages which did build successfully are toast.

If any of the three platforms experiences an error, then the other two will abort if they have not finished. So unless the other platforms finish before the error on a different job, then you can kiss their artifacts good bye too.

Don't worry though: Microsoft is happy to bill you for all the hours of redundant dependency building that you have to endure as you go through the process of discovering a configuration that works on all platforms.

At least it's easy to debug though. Vcpkg helpfully saves all the build output to files so that you can read the exact build error message... if you were running this interactively on a console instead of as part of a ci script.

The two tools in question (github actions and vcpkg) are reasonably decent when each is considered in isolation, but somehow they manage to achieve negative synergy when combined.
## [11][Random Testing for C and C++ Compilers with YARPGen](https://www.reddit.com/r/cpp/comments/jklgh5/random_testing_for_c_and_c_compilers_with_yarpgen/)
- url: https://www.cs.utah.edu/~regehr/yarpgen-oopsla20.pdf
---

