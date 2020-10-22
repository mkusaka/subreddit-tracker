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
## [2][Fast, typo tolerant instant search engine written in C++](https://www.reddit.com/r/cpp/comments/jfddui/fast_typo_tolerant_instant_search_engine_written/)
- url: https://github.com/typesense/typesense
---

## [3][clang-tidy support in VSCode is possible](https://www.reddit.com/r/cpp/comments/jfktsf/clangtidy_support_in_vscode_is_possible/)
- url: https://www.reddit.com/r/cpp/comments/jfktsf/clangtidy_support_in_vscode_is_possible/
---
if you're an exceptionally attractive individual like me, the only things preventing you from retiring clion and adopting vscode (for everything) are **clang-tidy** and **cppcheck**. we can get clang-tidy by upvoting this feature request for adding clang-tidy support to the cpp extension for vscode: [https://github.com/microsoft/vscode-cpptools/issues/2908](https://github.com/microsoft/vscode-cpptools/issues/2908)

clang-tidy good. more clang-tidy.
## [4][Qt and idiomatic smart pointer usage](https://www.reddit.com/r/cpp/comments/jfoeht/qt_and_idiomatic_smart_pointer_usage/)
- url: https://www.reddit.com/r/cpp/comments/jfoeht/qt_and_idiomatic_smart_pointer_usage/
---
One of the projects I work on is a gui application based on a c++ library which is written in c++17 style. The gui use Qt, but the library itself doesn't use any Qt classes at all except for optional model objects which inherit QAbstractItemModel. Any pointers that are part of the library api are std::unique_ptr or (rarely) std::shared_ptr.

This creates some friction because Qt's ownership model doesn't mesh very well with modern coding styles. Even though other patterns are grudgingly tolerated, Qt wants you to create widget objects on the stack with new and pass in a parent pointer which will take ownership of the object. This feels like a major step backwards when the non-gui parts of the project have successfully eliminated usage of raw new / delete usage.

The solution we came up with is based on a helper class called ScopeGuard (*):

    class ScopeGuard
    {
    public:
        using Callback = std::function&lt;void()&gt;;

        ScopeGuard(Callback cb) noexcept;
            : cb_(cb)
        {
        }
        ~ScopeGuard()
        {
            if (cb_) { cb_(); }
        }

    private:
        const SimpleCallback cb_;
    };

With that class available, it's possible to write code for creating and displaying a modal dialog that looks like this:

    {
        ...
        auto dialog = std::make_unique&lt;MyDialog&gt;(this);
        auto postcondition = ScopeGuard{[&amp;]() {
            dialog-&gt;deleteLater();
            dialog.release();
        }};
        connect(dialog.get(), &amp;MyDialog::signal, this, &amp;MyType::slot)
        ...
        dialog-&gt;exec();
    }

Using this pattern I still allow Qt to control object lifetime on its own terms. In particular I don't need to worry about whether the signal/slot connections will be cleaned up before the dialog object.

At the same time, raw usage of new is avoided and the owership semantics are more clearly conveyed to anyone reading the code. Any coder looking at this function may not realize **why** ownership of the object is being given up in this way, but it is clear that what is happening is deliberate even to someone unfamiliar with Qt's ownership model.

(*) It's probably obvious but the class name "ScopeGuard" was invented by a team member who is a fan of Dlang.
## [5][Ridiculously fast unicode (UTF-8) validation](https://www.reddit.com/r/cpp/comments/jf0tx4/ridiculously_fast_unicode_utf8_validation/)
- url: https://lemire.me/blog/2020/10/20/ridiculously-fast-unicode-utf-8-validation/
---

## [6][Learning C before C++](https://www.reddit.com/r/cpp/comments/jfjdzy/learning_c_before_c/)
- url: https://www.reddit.com/r/cpp/comments/jfjdzy/learning_c_before_c/
---
Is there any benefit in learning C before C++? Is C lower level than C++, meaning I would learn more about how computers work (on a lower level)? I am planning on taking CS50 soon, which I believe teaches C too, so if there is benefits in learning C before C++, I would learn it there and could afterwards directly start C++, or should I learn C in more depth?

I am mostly interested in cybersecurity, AI, game development, simulations, and more.

What are the advantages of C over C++ in general, as in what kind of programs is C more suited for than C++?
## [7][I created a technical tutorial demonstrating how to design a language-agnostic cross-platform computer Vision SDK (written in C++).](https://www.reddit.com/r/cpp/comments/jfn1a8/i_created_a_technical_tutorial_demonstrating_how/)
- url: https://www.reddit.com/r/cpp/comments/jfn1a8/i_created_a_technical_tutorial_demonstrating_how/
---
In the tutorial, I explain how to:

* Build a basic computer vision library in C++
* Compile and cross-compile the library for AMD64, ARM64, and ARM32
* Package the library and all the dependencies as a single static library
* Automate unit testing
* Set up a continuous integration (CI) pipeline
* Write python bindings for our library
* Generate documentation directly from our API

The tutorial has a corresponding video explanation and all the code is open source an available on github. 

Check it out, I'm open to any feedback or suggestions too.

Hope it helps you get started with your next big C++ project!

[https://medium.com/trueface-ai/how-to-design-a-language-agnostic-cross-platform-computer-vision-sdk-e437ecac8b4e](https://medium.com/trueface-ai/how-to-design-a-language-agnostic-cross-platform-computer-vision-sdk-e437ecac8b4e)
## [8][2020-10 C++ Committee Mailing](https://www.reddit.com/r/cpp/comments/jf4wsw/202010_c_committee_mailing/)
- url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/#mailing2020-10
---

## [9][json_dto-0.2.11: custom formatters added to an easy to use wrapper around RapidJSON](https://www.reddit.com/r/cpp/comments/jf82ka/json_dto0211_custom_formatters_added_to_an_easy/)
- url: https://eao197.blogspot.com/2020/10/progc-jsondto-0211-released.html
---

## [10][What to use for writing large documentations?](https://www.reddit.com/r/cpp/comments/jexkdk/what_to_use_for_writing_large_documentations/)
- url: https://www.reddit.com/r/cpp/comments/jexkdk/what_to_use_for_writing_large_documentations/
---
This is not strictly CPP, but hey, I like you guys.

I'm planning some projects and will need to have quite large and extensive documentation.

Immediately I thought I'd just make a git repo full of markdown files for documentation as it's not going to be about specific functions but instead overviews of technology and systems (It will be updated before code so I wasn't worried about it being out of date).  
If I did that I could just host it with a github IO page or even just leave it as a repo.

Then I thought a little further and decided maybe I should use an actual system to track it since it would be handy to have more automatic references and such. Some documentation tracker / generator.

So do you guys have any suggestions for writing docs for large systems? Repos, doc generators, a CMS?



A little further clarification:  
The code itself will be commented like usual and I'll use Doxygen or something to generate that data.  
The documentation I need to write is more of an overview of the whole code. End users are not touching the code for the project, just using it. So it will not be "this function does X", rather it's "Press this button to do Y".  
That's why I'm not worried about one being out of date, they are separate pieces of data (implementation vs behavior).

Edit:
The documentation is going to be very technical, just not actual code.  
I also think I would prefer for the documentation to be in git not in a database.
## [11][Does anybody use std::list?](https://www.reddit.com/r/cpp/comments/jf3dwo/does_anybody_use_stdlist/)
- url: https://www.reddit.com/r/cpp/comments/jf3dwo/does_anybody_use_stdlist/
---
I have never used std::list the last 15 years,  mainly for performance reasons. If I need to store large objects, I use a vector of pointers. Does anybody have a use case that can share? 
Thanks!
