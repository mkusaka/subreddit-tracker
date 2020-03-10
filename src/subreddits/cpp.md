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
## [2][Fast float parsing in practice](https://www.reddit.com/r/cpp/comments/fgcufa/fast_float_parsing_in_practice/)
- url: https://lemire.me/blog/2020/03/10/fast-float-parsing-in-practice/
---

## [3][(Yet another) C++17 Signal-Slots (First Part)](https://www.reddit.com/r/cpp/comments/ffzp4o/yet_another_c17_signalslots_first_part/)
- url: https://dreamdota.com/c-17-signals/
---

## [4][Quick object instantiation question](https://www.reddit.com/r/cpp/comments/fgcy9c/quick_object_instantiation_question/)
- url: https://www.reddit.com/r/cpp/comments/fgcy9c/quick_object_instantiation_question/
---
Hi guys. I have been coding for a few months now, and I have a quick question about object instantiation.

&amp;#x200B;

I am using a class called cv::namedWindow, and in order to use it, I do not supply a name for the instantiated object. instead, I just pass values in to the constructor. This is weird. Have I instantiated the object or not and where does it exist in memory?

&amp;#x200B;

Eg.

cv::namedWindow("myWindow", cv::WINDOW\_NORMAL);

See, I havent supplied the name to the object, so how I am I to use this object later on?
## [5][Why does including existing libraries (e.g. fmt, ranges, etc.) into the standard matter?](https://www.reddit.com/r/cpp/comments/fg2kvq/why_does_including_existing_libraries_eg_fmt/)
- url: https://www.reddit.com/r/cpp/comments/fg2kvq/why_does_including_existing_libraries_eg_fmt/
---
After seeing all the excitement from people about upcoming changes to c++20 with the std::format and std::ranges, I'm a bit confused. The libraries exist already and you could make use of them before, yet people put a lot of effort and time into writing the technical specifications and proposals to include them into the std. I've seen Eric Niebler talking about including ranges in standard in 2015 and it only made its way in c++20.

Is it done to make the compiler developers work on supporting them, therefore making it more stable, or to increase awareness that such libraries exist? I'm just curious as I feel like missing a big part of reasoning)
## [6][Bool in C++](https://www.reddit.com/r/cpp/comments/fgdnuh/bool_in_c/)
- url: https://www.reddit.com/r/cpp/comments/fgdnuh/bool_in_c/
---
I'm going through PPP by Bjarne Stroustrup when I came across this line:

bool c = "fool"&lt;s;

I have two questions, how does bool process string inputs and what does the "&lt;s" do?
## [7][How would passing non-trivial types in registers preserve semantics?](https://www.reddit.com/r/cpp/comments/fgdcs9/how_would_passing_nontrivial_types_in_registers/)
- url: https://www.reddit.com/r/cpp/comments/fgdcs9/how_would_passing_nontrivial_types_in_registers/
---
Many recent posts in this subreddit about breaking ABI compatibility mention the ability to pass `unique_ptr` in a register as one of the potential benefits. While I agree this would be beneficial, I wonder how it can be done while preserving the following fundamental properties of objects:
- Construction / destruction order
- Object identity

I was taught that, in general, objects are destroyed in the opposite order of creation. This applies to class members, bases, function locals, and most crucially to temporaries.

In the expression:
   
    f(a(), b(), c());
   
The calls to `a()`, `b()`, and `c()` are indeterminately sequenced, but the temporary values returned live until the call to `f` returns, and then are destroyed in the opposite order. If one of those temporaries is a non-trivial object and is passed through a register, will this continue to be true?

I was also taught that, in general, the address of an object uniquely identifies that object (modulo base classes, first members, or objects that provide storage for other objects). In particular, if for some non-trivial type a constructor runs with `this` = some address, I can expect the destructor to eventually run with the same `this`. If we allow passing non-trivial objects through registers, how will this property be preserved?

I'm aware of one situation where the above is not the case - the proposed `[[trivially_relocatable]]` attribute. `unique_ptr` and even `shared_ptr` would benefit greatly from it, and its use "answers" my concerns. But will the imagined future ABI that allows passing non-trivial types in registers only allow it for types that opt in to this attribute?
## [8][Differences between std::string_view and std::span](https://www.reddit.com/r/cpp/comments/fg79hm/differences_between_stdstring_view_and_stdspan/)
- url: https://www.nextptr.com/question/qa1344313286/differences-between-stdstring_view-and-stdspan
---

## [9][How to remove cplusplus.com from Google search results?](https://www.reddit.com/r/cpp/comments/ffpd1s/how_to_remove_cpluspluscom_from_google_search/)
- url: https://www.reddit.com/r/cpp/comments/ffpd1s/how_to_remove_cpluspluscom_from_google_search/
---
it even \*pretends\* to be cppreference, that's how bad it is! &lt;/rant&gt;

Edit: answer in comments, works for Firefox which I happen to use: [https://www.reddit.com/r/cpp/comments/ffpd1s/how\_to\_remove\_cpluspluscom\_from\_google\_search/fk00gpl?utm\_source=share&amp;utm\_medium=web2x](https://www.reddit.com/r/cpp/comments/ffpd1s/how_to_remove_cpluspluscom_from_google_search/fk00gpl?utm_source=share&amp;utm_medium=web2x)

Edit2: personal block list [https://www.reddit.com/r/cpp/comments/ffpd1s/how\_to\_remove\_cpluspluscom\_from\_google\_search/fk00noj?utm\_source=share&amp;utm\_medium=web2x](https://www.reddit.com/r/cpp/comments/ffpd1s/how_to_remove_cpluspluscom_from_google_search/fk00noj?utm_source=share&amp;utm_medium=web2x)

edit3: how to improve the proper reference: [https://www.reddit.com/r/cpp/comments/ffpd1s/how\_to\_remove\_cpluspluscom\_from\_google\_search/fk06he5?utm\_source=share&amp;utm\_medium=web2x](https://www.reddit.com/r/cpp/comments/ffpd1s/how_to_remove_cpluspluscom_from_google_search/fk06he5?utm_source=share&amp;utm_medium=web2x)
## [10][Const collection of unique_ptr, options and design choices](https://www.reddit.com/r/cpp/comments/fg2t6i/const_collection_of_unique_ptr_options_and_design/)
- url: https://www.bfilipek.com/2020/03/const-collection-of-uniqueptr-options.html
---

## [11][Shared General Utilities](https://www.reddit.com/r/cpp/comments/fg0bv9/shared_general_utilities/)
- url: https://www.reddit.com/r/cpp/comments/fg0bv9/shared_general_utilities/
---
&gt;Code I never intended to be part of a library is now used by many  applications and in environments I never foresaw. ***A goal for a  programmer has to be to look beyond the product they are shipping and  recognize their obligation to create correct and efficient solutions*** and  understand that their code may well endure, for good or bad.   
&gt;  
&gt;\- Sean Parent.

A lot of us have written something particularly useful to our case, which we then generalized to all -or most- uses. The idea of the thread is to share and discuss such code, the sort that would fit in a "GeneralUtilities.h" header with the namespace `gu` .

Code that was sufficiently discussed and is nigh unanimously approved of will be pinned to this post. Better means of saving will be used as needed.
