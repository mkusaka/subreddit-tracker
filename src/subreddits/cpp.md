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
## [2][Official C++ Logo is now under public domain](https://www.reddit.com/r/cpp/comments/i6u3xq/official_c_logo_is_now_under_public_domain/)
- url: https://www.reddit.com/r/cpp/comments/i6u3xq/official_c_logo_is_now_under_public_domain/
---
As you can know, Official C++ Logo ([isocpp/logos](https://github.com/isocpp/logos)) is not free since 18.04.2017 (※ [`a036ea6`](https://github.com/isocpp/logos/commit/a036ea65afa8b5f5ba7733f90d9aed8266eca6c1)).

Fortunatelly, I was recently cleaning up my hard drive space and found old C++ logo downloaded while C++ logo was under public domain, so in order to allow community use it freely, I am distributing it under public domain: [Benio101/cpp-logo](https://github.com/Benio101/cpp-logo) so…

**Everyone can use, modify and redistribute C++ Logo without any restrictions**.

⚠️ Note that in order to use unrestricted version of C++ Logo, you have to download it from my repository, not from isocpp once due to licensing.

Have a nice day.
## [3][Cross-platform process launching: reproc vs. tiny-process-library](https://www.reddit.com/r/cpp/comments/i6zxua/crossplatform_process_launching_reproc_vs/)
- url: https://www.reddit.com/r/cpp/comments/i6zxua/crossplatform_process_launching_reproc_vs/
---
I need a standalone cross-platform library for starting processes and communicating with them via stdin/stoud. Leaving `boost::process` (and other beasts like Qt's `QProcess`) aside, I've found only two such libraries: [reproc](https://github.com/DaanDeMeyer/reproc) and [tiny-process-library](https://github.com/eidheim/tiny-process-library). Both claim to be cross-platform and to provide communications via stdin/stdout , both have comparable numbers of github stars (though `tiny-process-library` moved to gitlab), both seem to be actively maintained. Has anyone had experience with both of them to provide unbiased advise on which one to choose? (Possibly biased) arguments from the developers ([u/DaanDeMeyer](https://www.reddit.com/user/DaanDeMeyer/)?) are also welcomed :)
## [4][Is C++ too conservative with UB rules?](https://www.reddit.com/r/cpp/comments/i6xn76/is_c_too_conservative_with_ub_rules/)
- url: https://www.reddit.com/r/cpp/comments/i6xn76/is_c_too_conservative_with_ub_rules/
---
Say, you need to write a function, which takes an array / vector, and returns bool if there exists any element, satisfying certain predicate. You create a local bool variable, and conditionally may set it to true, from multiple threads. Without any synchronization.

From C++ point of view, it's UB.

But realistically, given common architectures and compilers, what could possibly go wrong? Can an update - setting it to true - be lost?
## [5][envy: Deserialize environment variables into type-safe structs](https://www.reddit.com/r/cpp/comments/i73c2a/envy_deserialize_environment_variables_into/)
- url: https://github.com/p-ranav/envy
---

## [6][Alternatives to Visual Assist for C++ in Visual Studio](https://www.reddit.com/r/cpp/comments/i734hh/alternatives_to_visual_assist_for_c_in_visual/)
- url: https://www.reddit.com/r/cpp/comments/i734hh/alternatives_to_visual_assist_for_c_in_visual/
---
I'm looking for Visual Studio add-ons that improve intellisense and the refactoring tools like Visual Assist does. 

I search for something else than VA as VA has become a major annoyance lately. You cannot simply buy a key and download the software anymore but you need to register with some third party site and (at least I) had to go through a lengthy email conversation because I wanted to buy a private licence and my last name changed from years ago due to marriage.

I surely don't want to go through this again so I'm looking for a alternative to VA. What are you guys using?
## [7][What's the deal with ConstexprIterator and what's a MeowIterator ?](https://www.reddit.com/r/cpp/comments/i6hcbx/whats_the_deal_with_constexpriterator_and_whats_a/)
- url: https://www.reddit.com/r/cpp/comments/i6hcbx/whats_the_deal_with_constexpriterator_and_whats_a/
---
I was checking the new features of c++20 on cppreference and got intrigued by the [ConstexprIterator](https://en.cppreference.com/mwiki/index.php?title=cpp/named_req/ConstexprIterator&amp;oldid=119337) concept. I mean iterators are pretty cool but iterators that work in constexpr/consteval function are even cooler! And then I read the requirements:

&gt; **Requirements**

&gt; The type It satisfies *ConstexprIterator* if

&gt; * The type It satisfies some iterator requirements *MeowIterator*

&gt; And, for every

&gt; * purr, an operation on It that is required to be supported by *MeowIterator*,

&gt; * kittens..., a set of arguments to purr that meets the requirements for that operaton,

&gt; Then

&gt; * purr(kittens...) may be used in a constant expression if kittens... can be so used

which is kind of funny but brings the question what is a *MeowIterator*? I really hope this isn't a joke but I have very little faith in that.

Seeing that array::iterator satisfies ConstexprIterator on cppreference I tried to compile a small program using g++-10 but std::ConstexprIterator isn't recognized...
## [8][How to Check String or String View Prefixes and Suffixes in C++20](https://www.reddit.com/r/cpp/comments/i71it6/how_to_check_string_or_string_view_prefixes_and/)
- url: https://www.bfilipek.com/2020/08/string-prefix-cpp20.html?m=1
---

## [9][Experienced C++ role for innovative open trading platform based on automatic trading strategies](https://www.reddit.com/r/cpp/comments/i745p4/experienced_c_role_for_innovative_open_trading/)
- url: https://www.reddit.com/r/cpp/comments/i745p4/experienced_c_role_for_innovative_open_trading/
---
**Experienced C ++ developer who helps BOTS realize its worldwide ambition**

**About the company**

Our goal is to make the exclusive technology of algorithmic trading accessible for everyone. By providing the tools to trade as a pro, we hope to shake up the financial world and give everyone equal chances. 

Your contribution to democratize the world of trading will be at the core of our product, therefore, we are looking for an experienced C++ developer who makes a differcence..

According to Emerce.nl, BOTS by RevenYOU is ‘’One of the top seven fintech companies to look out for this year!’’ and dutch media company Sprout called us ‘’The google of the stock exchange!’’. None of this changes the fact that we are a young and fun company with a shared goal: make BOTS a success! 

BOTS  is a financial, but informal start-up, where you work with ambitious colleagues. Investing, quotes, stocks, asset markets. At BOTS we thought: it all sounds complicated and many people find it difficult. How do we make this more accessible? An app of course! BOTS has developed an app that makes investing more accessible, arranged and clearer. It is an innovative open trading platform based on automatic trading strategies and machine learning. Invest like a pro, without being one!

**The role**

We are looking for an experienced and driven C ++ developer who will work on the core of our product. C++ is the most important language that you need to have mastered already for this function. 

Your role:

* You are responsible for the development and functionality of our exchange. Many users rely on your coding abilities and you carry a large responsibility in the form of millions of transactions.
* You will be defining, running and monitoring large computational workloads, running on low latency infrastructure. 
* Researching, we expect that you have great research skills and that this is also something you find challenging. We have to find ways to improve differents part of our systeem. This can be about defining and implementing financial models orderbooks, api’s ect.

BOTS’ raised new capital and the continued development of our systems is essential. We want to grow to 120 countries with a lot of users. You will be in the core of this expansion, an exciting time of an exciting company to be part of. This is a once in a lifetime opportunity. 

An important task will be to further shape our own exchange It is an advantage if you already have experience as a developer. For the exchange you need to be very innovative and reliable,. This area could be overwhelming, so you need to be confident.. 

**What do you need for this role?**

* Preferably someone who has completed an MSc in computer science
* 6 year + experience with C++
* Knowledge of algorithms and data structures
* Core experience as engineer with trading and gaming structures
* Ability to provide clean and reliable code
* A proactive approach to problem solving
* Someone who can deal well with conflict and field-critical questions
* An independent thinker who is creative in finding alternative solutions

**Our offer**

* We offer an excellent, competitive salary based on your work experience 
* 25 vacation days per year 
* If you are very ambitious, growth is possible! 
* A pleasant workplace, with an open, fun and young international team! 
* A Friday afternoon drink with beer! 
* Table football clashes between employees. 
* Room for your personal development to learn and discover a lot, with a focus on craftsmanship. 
* A lot of independence, responsibility and a lot of room for initiative.
* Virtual shares certificates, so you really become part of a great journey to change the world of trading. 

Are you interested in this position? Sign up quickly! Send an email with your motivation and CV to jobs@revenyou.io
## [10][Quill Asynchronous Low Latency Logging Library - v1.3.3](https://www.reddit.com/r/cpp/comments/i6lbsw/quill_asynchronous_low_latency_logging_library/)
- url: https://www.reddit.com/r/cpp/comments/i6lbsw/quill_asynchronous_low_latency_logging_library/
---
It’s been a few months since the [initial post](https://www.reddit.com/r/cpp/comments/fcbflb/quill_an_asynchronous_low_latency_logging_library) about [quill](https://github.com/odygrd/quill) and now the library is more mature. 

I am posting again to get some new feedback.  

There were a few suggestions/feature requests in previous post added to the project. 

* It is now possible to disable exceptions (@[dextorious\_](https://www.reddit.com/user/dextorious_/))
* It is now possible to use a bounded queue and drop logs (@[matthieum](https://www.reddit.com/user/matthieum/))
* There was some discussion (@[matthieum](https://www.reddit.com/user/matthieum/)) about copying user defined types with mutable shared references and call to ‘operator&lt;&lt;‘ on the backend logging thread. There is now compile time protection against non-trivial user defined types that are passed to the logger. see [https://github.com/odygrd/quill/wiki/5.-User-Defined-Types](https://github.com/odygrd/quill/wiki/5.-User-Defined-Types) 
* The library now can be found in conan (@[gocarlos](https://www.reddit.com/user/gocarlos/)) , vcpkg (@[tracyma](https://www.reddit.com/user/tracyma/)) and homebrew package managers

Complete changelog is [here](https://github.com/odygrd/quill/blob/master/CHANGELOG.md)

Current version is 1.3.3 with version 1.4 being work in progress.

Looking forward to some feedback and please report any issues if you decide to use the library.
## [11][std::valstat - an metastate return type](https://www.reddit.com/r/cpp/comments/i6zxg2/stdvalstat_an_metastate_return_type/)
- url: https://www.reddit.com/r/cpp/comments/i6zxg2/stdvalstat_an_metastate_return_type/
---
## metastate

"A paradigm is a standard, perspective, or set of ideas. A paradigm is a way of looking at something ... When you change paradigms, you're changing how you think about something..." vocabulary.com

metastate is a software development paradigm. metastate is not yet another error return handling idiom. More precisely metastate is a paradigm about obtaining more information as a result of a function call. 

## valstat

`std::valstat` is an metastate C++ implementation. `std::valstat` is instrumental to the success of metastate in C++ programs.

Together metastate and valstat are assembling a programming behaviour pattern.

## [Proposal P2192](https://gitlab.com/dbjdbj/valstat/-/blob/master/P2192R1.md)
