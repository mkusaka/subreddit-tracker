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
## [2][17 Smaller but Handy C++17 Features](https://www.reddit.com/r/cpp/comments/jdx37j/17_smaller_but_handy_c17_features/)
- url: https://www.bfilipek.com/2019/08/17smallercpp17features.html
---

## [3][Initialization in C++ is Seriously Bonkers](https://www.reddit.com/r/cpp/comments/jdz5cx/initialization_in_c_is_seriously_bonkers/)
- url: http://mikelui.io/2019/01/03/seriously-bonkers.html
---

## [4][Asio users, how do you deal with cancellation?](https://www.reddit.com/r/cpp/comments/jdy2gd/asio_users_how_do_you_deal_with_cancellation/)
- url: https://www.reddit.com/r/cpp/comments/jdy2gd/asio_users_how_do_you_deal_with_cancellation/
---
First, by cancellation I don't mean _just_ calling the `socket::cancel(...)` method (which has its own quirks, but let's not go into that). I mean any means of forcing a callback (or coroutine or whatever) to execute with an error code. E.g. calling `socket::close()`, `socket::cancel()`, `socket::~socket()`, `timer::cancel()`,...

The main problem I have is with [the remark section of timer::cancel function](https://www.boost.org/doc/libs/1_74_0/doc/html/boost_asio/reference/basic_waitable_timer/cancel/overload1.html):


    If the timer has already expired when cancel() is called, then the handlers for asynchronous wait operations will:

        * have already been invoked; or
        * have been queued for invocation in the near future.

    These handlers can no longer be cancelled, and therefore are passed an error code that indicates the successful completion of the wait operation. 


I presume the same applies for all other means of cancellation and for all other IO objects as well, although the Asio documentation is not explicit about this anywhere else (please someone correct me if I'm wrong).

This means that I can't rely on errors being passed to the caller and thus I need to perform an explicit check pretty much after each async command.

Note that Asio examples utilize std::enable_shared_from_this, which **partially** solves the issue in that I no longer need to worry about memory leaks. But one still has to explicitly check that things haven't been "cancelled" after each async command in order to exit the "cancelled" task. Which is annoying in the long run:

    struct MyAsyncTask : std::enable_shared_from_this&lt;MyAsyncTask&gt; {
        bool stopped = false;

        // To avoid boiler plate in the example, say we're using coroutines
        void operator()(asio::yield_context yield) {
            auto self = shared_from_this(); // to prevent destruction of `this`

            do_async_task(yield);
            if (stopped) throw runtime_error(asio::error::operation_aborted);

            do_another_async_task(yield);
            if (stopped) throw runtime_error(asio::error::operation_aborted);

            ...
        }
    
        void stop() {
            stopped = true;
            // close sockets, timers and any other io objects used in this class.
        }
  
        ...
    };

Thus I would like to ask Asio users on reddit: what is your preferred way of dealing with this annoyance?
## [5][A typedef for when an alias declaration cannot](https://www.reddit.com/r/cpp/comments/jdh549/a_typedef_for_when_an_alias_declaration_cannot/)
- url: https://dfrib.github.io/typedef-where-alias-decl-cannot/
---

## [6][Static Registration Macros (e.g. TEST_CASE("foo") { ... })](https://www.reddit.com/r/cpp/comments/jdgxcl/static_registration_macros_eg_test_casefoo/)
- url: https://artificial-mind.net/blog/2020/10/17/static-registration-macro
---

## [7][Upp sounds too good to be true, what is the catch?](https://www.reddit.com/r/cpp/comments/jdfp7z/upp_sounds_too_good_to_be_true_what_is_the_catch/)
- url: https://www.reddit.com/r/cpp/comments/jdfp7z/upp_sounds_too_good_to_be_true_what_is_the_catch/
---
I am new to C++ and I have find out about this project only recently. I have tried Qt - mostly in Qt Creator IDE and some console stuff, obviously, Vaguely tried wxWidget but it was a pain to setup it on Ubuntu. So, installign Ultimate++ was extremely simple, there is a ton of examples and it feels good.

I haven't found vim keybinding option, probably is not there but I am not sure.

Check this latest article about U++:

[https://www.codeproject.com/Articles/5268022/Getting-Started-with-Uplusplus-2020-1#why](https://www.codeproject.com/Articles/5268022/Getting-Started-with-Uplusplus-2020-1#why)

The author claims things like:

&gt;At the same time, U++ Core libraries provide about **3 times better performance**  than standard C++ library in areas like string operations and mapping,  which consist of a bulk of load in many scenarios (think e.g. about  processing JSON or parsing files).

What do you think about it and more importantly why is it not more known for creating apps that doesn't need flashy buttons and stuff and just work with a nice and simple BSD license?

Can somebody who is experienced in C++ check these comparisons and comment on them:

[https://www.ultimatepp.org/www$uppweb$comparison$en-us.html](https://www.ultimatepp.org/www$uppweb$comparison$en-us.html)

I am probably missing something. What is the catch?

P.S. The reason I read the article on codeproject was... [Jet-Story](https://www.youtube.com/watch?v=YTsF6_V3cT0). I was searching something about old ZX Spectrum games and wondering what my [three 80's heroes](https://doupe.zive.cz/clanek/10-osobnosti-ceske-herni-sceny-ktere-byste-meli-znat/ch-146306) of [Golden Triangle](https://cs.wikipedia.org/wiki/Golden_Triangle) (Cybexlab, T.R.C., Fuxoft) do nowadays. Then I saw some hidden comment or something from Frantisek Fuka that Miroslav Fídler still does programming. And yes, it turns out that [Mirek Fídler](https://cs.wikipedia.org/wiki/Miroslav_F%C3%ADdler) is the author of U++ and that codeproject article and... Jet-Story... and other things like [this](https://en.wikipedia.org/wiki/Belegost_(video_game)) or [these](https://spectrumcomputing.co.uk/list?label_id=2918).
## [8][GCC AVR](https://www.reddit.com/r/cpp/comments/jdl3a4/gcc_avr/)
- url: https://www.reddit.com/r/cpp/comments/jdl3a4/gcc_avr/
---
Is anyone here on the gcc dev team?  I understand that AVR has been deprecated in 10 and is gone in 11.  It looks like someone took on the task of making the conversion to MODE\_CC or whatever, but the last messages I see regarding the topic on the gcc list end in August.

I didn't want to join just to ask status, so I'm hoping someone here might have an idea?  Is it still a matter of getting someone to review the guy's code?

I'm interested in this conversion because I'm very interested in challenging the latest language features in an embedded, microcontroller environment.  Mainly Arduino.  Gcc appears to be way further along in implementing C++20 than clang.
## [9][Has anyone made a Standard C++ proposal for a package manager?](https://www.reddit.com/r/cpp/comments/jdftfn/has_anyone_made_a_standard_c_proposal_for_a/)
- url: https://www.reddit.com/r/cpp/comments/jdftfn/has_anyone_made_a_standard_c_proposal_for_a/
---

## [10][xmake v2.3.8 released, Add Intel C++/Fortran Compiler Support](https://www.reddit.com/r/cpp/comments/jdfk6r/xmake_v238_released_add_intel_cfortran_compiler/)
- url: https://github.com/xmake-io/xmake/wiki/xmake-v2.3.8-released,-Add-Intel-Compiler-Support
---

## [11][new and delete operator time complexity](https://www.reddit.com/r/cpp/comments/jdhsp4/new_and_delete_operator_time_complexity/)
- url: https://www.reddit.com/r/cpp/comments/jdhsp4/new_and_delete_operator_time_complexity/
---
I was wondering, what is the time complexity of new, new[], delete and delete[]. I assumed it was constant, but I was reading that it can go to be linear.
