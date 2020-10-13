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
## [2][Question about packages coming from a Python Perspective](https://www.reddit.com/r/cpp/comments/ja8ybj/question_about_packages_coming_from_a_python/)
- url: https://www.reddit.com/r/cpp/comments/ja8ybj/question_about_packages_coming_from_a_python/
---
Okay, so I want to learn a little cpp, so I decided to do this facial recognition thing with OpenCV. I've done it before with python. Super easy, quick and simple. With Python you can just pip install the OpenCV module, and use it in seconds. However, when attempting with Cpp, it looks like I'm having to use this cmake command to build the OpenCV ordeal. This process it taking FOREVER. I am wondering if I am missing something or if this is just the process with cpp. I found out there isn't a pip for cpp unfortunetly. 

&amp;#x200B;

I guess what I am trying to figure out is:

\- Is this a normal application for cpp? (Facial Recognition)

\- Do you have to build each package you want to use? (ie OpenCV)

\- Is there some fundamental thing I am missing?
## [3][Clang 11.0.0 is out](https://www.reddit.com/r/cpp/comments/j9rq0y/clang_1100_is_out/)
- url: https://releases.llvm.org/11.0.0/tools/clang/docs/ReleaseNotes.html
---

## [4][Does this code snippet lead to undefined behavior?](https://www.reddit.com/r/cpp/comments/ja6xnq/does_this_code_snippet_lead_to_undefined_behavior/)
- url: https://www.reddit.com/r/cpp/comments/ja6xnq/does_this_code_snippet_lead_to_undefined_behavior/
---
I came across some code in android project

[https://cs.android.com/android/platform/superproject/+/master:system/netd/include/Fwmark.h;l=24?q=fwmark.h](https://cs.android.com/android/platform/superproject/+/master:system/netd/include/Fwmark.h;l=24?q=fwmark.h)

    union Fwmark {
        uint32_t intValue;
        struct {
            unsigned netId          : 16;
            bool explicitlySelected :  1;
            bool protectedFromVpn   :  1;
            Permission permission   :  2;
            bool uidBillingDone     :  1;
        };
        constexpr Fwmark() : intValue(0) {}
    
        static inline uint32_t getUidBillingMask() {
            Fwmark m;
            m.uidBillingDone = true;
            return m.intValue;
        }
    };

The intention here is to use union to do cast between bitfield and uint32\_t.

I suspect this is still undefined behavior.
## [5][C++ in a web app (series of blog posts)](https://www.reddit.com/r/cpp/comments/j9yjsj/c_in_a_web_app_series_of_blog_posts/)
- url: https://www.reddit.com/r/cpp/comments/j9yjsj/c_in_a_web_app_series_of_blog_posts/
---
Want to learn how to run your C++ code on the web? A series of blogs has been published to explain the steps to do this. Great for demonstration or to move the heavy computation from the server to the browser.

[Using C++ in a web app with WebAssembly](https://blog.esciencecenter.nl/using-c-in-a-web-app-with-webassembly-efd78c08469)

[Help! My C++ web app is not responding](https://blog.esciencecenter.nl/help-my-c-web-app-is-not-responding-b930ca3034ad)

[Interact with your C++ web app using React forms](https://blog.esciencecenter.nl/interact-with-your-c-web-app-using-react-forms-543e676a7634)

[Spice up your C++ web app with visualizations](https://blog.esciencecenter.nl/spice-up-your-c-web-app-with-visualizations-bcc1e888ec25)

[C++ web app with WebAssembly, Vega, Web Worker and React](https://blog.esciencecenter.nl/c-web-app-with-webassembly-vega-web-worker-and-react-1e5b750c88df)
## [6][Taruga is a simple single-header C++11 library for turtle graphics!](https://www.reddit.com/r/cpp/comments/j9yx0r/taruga_is_a_simple_singleheader_c11_library_for/)
- url: https://www.reddit.com/r/cpp/comments/j9yx0r/taruga_is_a_simple_singleheader_c11_library_for/
---
Examples (on Streamable): [drawing a Koch fractal](https://streamable.com/27ltmt) and [drawing a spiral](https://streamable.com/x5anqh).

Differently than other C++ implementations, Taruga offers a very simple API, it's code is easy to understand, needs only one (widely accessible) dependency and is hardware-accelerated.

More examples (and obviously, the code) are available at [the Taruga repo](https://github.com/vrmiguel/taruga/).

Any insights and (constructive) critique are welcome!
## [7][Faster C++ builds, simplified: a new metric for time](https://www.reddit.com/r/cpp/comments/j9tbeb/faster_c_builds_simplified_a_new_metric_for_time/)
- url: https://devblogs.microsoft.com/cppblog/faster-cpp-builds-simplified-a-new-metric-for-time/
---

## [8][Increased Complexity of C++20 Range Algorithms Declarations - Is It Worth?](https://www.reddit.com/r/cpp/comments/j9ndx1/increased_complexity_of_c20_range_algorithms/)
- url: https://www.bfilipek.com/2020/10/complex-ranges-algorithms.html?m=1
---

## [9][Upsetting Opinions about Static Analyzers](https://www.reddit.com/r/cpp/comments/j9ymgv/upsetting_opinions_about_static_analyzers/)
- url: https://www.viva64.com/en/b/0765/
---

## [10][`visitor&lt;Variants...&gt;` concept from scratch](https://www.reddit.com/r/cpp/comments/j9vu4c/visitorvariants_concept_from_scratch/)
- url: https://godbolt.org/z/xeW9Gf
---

## [11][When a coroutine ‘suspends’, does control flow halt or pass to the caller?](https://www.reddit.com/r/cpp/comments/j9r3ez/when_a_coroutine_suspends_does_control_flow_halt/)
- url: https://www.reddit.com/r/cpp/comments/j9r3ez/when_a_coroutine_suspends_does_control_flow_halt/
---
https://github.com/lewissbaker/cppcoro/blob/master/lib/async_manual_reset_event.cpp

I’m trying to understand what happens when await_suspend returns true. Does the thread of execution pass control back to where co_await was called by the user? Or does the thread of execution simply halt until we get a call to resume?

Put another way, after await_suspend returns true, where does the EIP register point? Are CPU resources returned to the OS?

After iterating through the linked list on a call to set(), how does resume() know which threads of execution to resume on? Does the coroutine handle hold this information?
