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
## [2][Is "C++ Template Metaprogramming" by Abrahams and Gurtovoy still relevant in the C++17/20 times?](https://www.reddit.com/r/cpp/comments/f16e9o/is_c_template_metaprogramming_by_abrahams_and/)
- url: https://www.reddit.com/r/cpp/comments/f16e9o/is_c_template_metaprogramming_by_abrahams_and/
---
I'm considering buying it, but am not sure whether it's still relevant or most of it outdated / implementing what's now already available in the language or std.  

Also feel free to name (better) alternatives if there are any. Thanks :)
## [3][fpm: a header-only fixed-point math library with trigonometry, etc](https://www.reddit.com/r/cpp/comments/f0ph40/fpm_a_headeronly_fixedpoint_math_library_with/)
- url: https://github.com/MikeLankamp/fpm
---

## [4][GitHub actions for vcpkg, CMake and Ninja](https://www.reddit.com/r/cpp/comments/f0oc24/github_actions_for_vcpkg_cmake_and_ninja/)
- url: https://www.reddit.com/r/cpp/comments/f0oc24/github_actions_for_vcpkg_cmake_and_ninja/
---
These are two GitHub actions that streamline building C++ sources with CMake, vcpkg, Ninja:

* [run-vcpkg](https://github.com/marketplace/actions/run-vcpkg)  
* [run-cmake](https://github.com/marketplace/actions/run-cmake)

Enjoy!
## [5][C++ Memory (Chrome University 2019)](https://www.reddit.com/r/cpp/comments/f0fx1x/c_memory_chrome_university_2019/)
- url: https://www.youtube.com/watch?v=UNJrgsQXvCA
---

## [6][A fast work-stealing queue template library in modern C++](https://www.reddit.com/r/cpp/comments/f0nu69/a_fast_workstealing_queue_template_library_in/)
- url: https://github.com/tsung-wei-huang/work-stealing-queue
---

## [7][C++ as a second language (Chrome University 2019)](https://www.reddit.com/r/cpp/comments/f0fxcq/c_as_a_second_language_chrome_university_2019/)
- url: https://www.youtube.com/watch?v=cN9c_JyvL1A
---

## [8][reproc v11.0.0 released!](https://www.reddit.com/r/cpp/comments/f0h238/reproc_v1100_released/)
- url: https://www.reddit.com/r/cpp/comments/f0h238/reproc_v1100_released/
---
I've just released [reproc](https://github.com/DaanDeMeyer/reproc) v11.0.0. reproc is a cross-platform process handling library for C and C++.

The big addition this release is that `reproc_poll` can now be used to poll more than one child process. This allows users to manage multiple processes from a single thread which is very useful to implement build systems (make, ninja, ...) that need to manage multiple child compiler processes from a single-threaded parent process.

While this was mostly trivial to implement on POSIX systems, Windows turned to be a lot harder due to the lack of a poll equivalent. I tried using I/O completion ports but they're extremely hard to adapt to the model of poll. I finally got lucky when I found out that child process standard streams can be redirected to sockets on Windows and Windows provides `WSAPoll` for use on sockets which has the same semantics as `poll` on POSIX. 

Unfortunately, Windows doesn't provide `socketpair` so instead we set up a TCP connection over localhost. I'm not 100% sure, but it seems TCP connections over localhost are somewhat optimized by Windows itself to avoid much of the overhead of the networking stack.

With access to `poll` on both platforms, it became doable to implement a `poll` for processes on top of it. To monitor child process exit, we have each child process inherit an extra socket which will be closed when the child process exits. This way, waiting for child process exit can be done as part of the main call to `poll`.

I'm hoping this can help developers with monitoring multiple child processes without having to pull in heavier dependencies such as boost process (along with boost asio) or libuv or be forced to start up a separate thread to monitor each separate child process. Of course, as soon as you have to poll more than just child processes, libuv or boost asio combined with boost process make a lot more sense than reproc.

I also added a bunch of extra options and a higher level `run` function which is on the same level as `system` but a lot safer.
    
    std::array&lt;std::string, 2&gt; args = { "echo", "Hello, World!" };
    auto [status, ec] = reproc::run(args);

I also tried to add support for pseudo-ttys using `openpty` on POSIX and the new ConPTY API on Windows but it seemed I'm the first one to try and use ConPTY with sockets on Windows which led to some [issues](https://github.com/microsoft/terminal/issues/4359).

I also made a (single-line) [contribution](https://github.com/microsoft/STL/pull/406) to the MSVC STL to fix a bug I found while working on reproc. It's amazing to be able to contribute to something as widely used as the MSVC STL. The experience was great so props to the developers at Microsoft that made this happen.

As always, any feedback is highly appreciated and if you have any questions about reproc, feel free to ask.
## [9][OpenSSL client and server from scratch series](https://www.reddit.com/r/cpp/comments/f099li/openssl_client_and_server_from_scratch_series/)
- url: https://quuxplusone.github.io/blog/2020/01/24/openssl-part-1/
---

## [10][Anti-Cheats - Taking an inside look at Razer's Synapse driver](https://www.reddit.com/r/cpp/comments/f0bmz0/anticheats_taking_an_inside_look_at_razers/)
- url: https://niemand.com.ar/2020/02/07/bypassing-anti-cheats-part-1-exploiting-razer-synapse-driver/
---

## [11][Quick Snake: A fast-faced snake-like game for the terminal in C++17](https://www.reddit.com/r/cpp/comments/f0f2fo/quick_snake_a_fastfaced_snakelike_game_for_the/)
- url: https://github.com/gregstula/quick-snake
---

