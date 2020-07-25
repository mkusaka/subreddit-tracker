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
## [2][BrainFuckStudio written on C++ using Qt](https://www.reddit.com/r/cpp/comments/hxkaep/brainfuckstudio_written_on_c_using_qt/)
- url: https://www.reddit.com/r/cpp/comments/hxkaep/brainfuckstudio_written_on_c_using_qt/
---
Pretty fast &amp; useful app for BF code writing &amp; executing

Supports working with files, personal interface customization, working with interpreter and etc.

[https://github.com/Djivs/BrainFuckStudio](https://github.com/Djivs/BrainFuckStudio)

That's my first post here, so I'm sorry if I did something wrong
## [3][How do you contribute to MSVC STL?](https://www.reddit.com/r/cpp/comments/hx764k/how_do_you_contribute_to_msvc_stl/)
- url: https://www.reddit.com/r/cpp/comments/hx764k/how_do_you_contribute_to_msvc_stl/
---
I recently ran into a bug in MSVC's &lt;random&gt; implementation:

[https://github.com/microsoft/STL/issues/1084](https://github.com/microsoft/STL/issues/1084)

I know how to fix their implementation, but was a bit unclear on the logistics of how to deliver the fix to them. Specifically:

1. What branch do I make a pull request into, `master`? How does Microsoft get these changes into the desired versions (15.9 vs 16.6)? I don't see any specific branches or tags that indicate specific MSVC versions.
2. I am used to the workflow (in our codebases) that every bug fix should come along with new unit tests to prevent the bug from ever regressing. Where is the appropriate place to add such unit tests, `std` or `libcxx`?

Not sure if /u/STL, another Microsoft dev, or anyone else can advise on this process?
## [4][C++ on Sea 2020 video: "Structured bindings uncovered" - Dawid Zalewski](https://www.reddit.com/r/cpp/comments/hxmhsj/c_on_sea_2020_video_structured_bindings_uncovered/)
- url: https://www.youtube.com/watch?v=uZCvz-E1heA
---

## [5][Embedding lightweight scripting into C++: Lua vs. QuickJS](https://www.reddit.com/r/cpp/comments/hxmctr/embedding_lightweight_scripting_into_c_lua_vs/)
- url: https://www.reddit.com/r/cpp/comments/hxmctr/embedding_lightweight_scripting_into_c_lua_vs/
---
Looking for an advice from those who have experience in embedding scripting langauges into C++, e.g., to provide app API. In my case, APIs are mostly for processing of unicode texts (including regexps usage for searching and replacements). My app is rather small, so I don't think I want to embed python or V8 javascript engine, as they will bloat it. Thus far, my choices are Lua (specifically, via [sol2](https://github.com/ThePhD/sol2)) and [QuickJS](https://bellard.org/quickjs/), developed a year ago by Fabrice Bellard; and claimed to be very efficient. Both sol2 and QuickJS are designed specificaly for embedding, both are quite lightweight and seem to yield comparable size overhead for the host app. Lua has no standard regexp, but that's not a big problem, as I can expose regexp-related functionaly through the host interface (e.g., using regexps from C++ standard library). What really interests me is how they would compare in performance when used as callback functions from the host app to crunch string data.

Based on your knowledge and/or experience, given these two options (Lua vs QuickJS), which one would you choose for an embedded API interface in general, and for string manipulations in particular?
## [6][Smart Pointers: When and Why? (walk-through video with motivational example)](https://www.reddit.com/r/cpp/comments/hwyhc8/smart_pointers_when_and_why_walkthrough_video/)
- url: https://www.youtube.com/watch?v=Zt0zUcDHi7I
---

## [7][Best C++ Alternatives to Pandas](https://www.reddit.com/r/cpp/comments/hx3fd9/best_c_alternatives_to_pandas/)
- url: https://www.reddit.com/r/cpp/comments/hx3fd9/best_c_alternatives_to_pandas/
---
Hi everyone,

I've been developing with python for years and have extensively used pandas.  I have a new project that requires me to code in C++ and I'm looking for a library that is similar to pandas.  I'd like to work with dataframes that have mixed data types.  It would be okay to have a fixed data type for each column in the dataframe but having columns with different data type is essential.  Ideally it could read data from csv files or json strings into the dataframe.  Speed is less important for me.  What do you guys suggest?  


Thanks!
## [8][C++ on Sea 2020 video - "Improving Readability With Class Template Argument Deduction" - A.J. Orians](https://www.reddit.com/r/cpp/comments/hx96c1/c_on_sea_2020_video_improving_readability_with/)
- url: https://www.youtube.com/watch?v=vgeOjM4X2K0
---

## [9][Avoid character-by-character processing when performance matters](https://www.reddit.com/r/cpp/comments/hx0o1g/avoid_characterbycharacter_processing_when/)
- url: https://lemire.me/blog/2020/07/21/avoid-character-by-character-processing-when-performance-matters/
---

## [10][C++ on Sea 2020 video - "Live Compiler Development with Cross-Platform Tooling" - Sy Brand](https://www.reddit.com/r/cpp/comments/hwyrh2/c_on_sea_2020_video_live_compiler_development/)
- url: https://www.youtube.com/watch?v=Jw2n5QSul34
---

## [11][fork of compiler explorer (godbolt.org) editing and displaying SDL in the browser with WebAssembly](https://www.reddit.com/r/cpp/comments/hwk1se/fork_of_compiler_explorer_godboltorg_editing_and/)
- url: https://v.redd.it/dsxz9xt07nc51
---

