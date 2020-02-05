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
## [2][Debugging with LLVM: A quick introduction to LLDB and LLVM sanitizers](https://www.reddit.com/r/cpp/comments/eyzj4k/debugging_with_llvm_a_quick_introduction_to_lldb/)
- url: https://fosdem.org/2020/schedule/event/debugging_with_llvm/
---

## [3][Effective Memory Reclamation for Lock-Free Data Structures in C++](https://www.reddit.com/r/cpp/comments/ez7jlv/effective_memory_reclamation_for_lockfree_data/)
- url: http://repositum.tuwien.ac.at/obvutwhs/download/pdf/2582190?originalFilename=true
---

## [4][Idea: Instead of implementing Visitor Pattern, define a conversion to a std::variant](https://www.reddit.com/r/cpp/comments/eyqxgp/idea_instead_of_implementing_visitor_pattern/)
- url: https://www.reddit.com/r/cpp/comments/eyqxgp/idea_instead_of_implementing_visitor_pattern/
---
Manually implementing the Visitor Pattern (visit and accept) requires a lot of repeated code.  Attempting to use template programming to ease this boilerplate is complicated.

So why not just define a virtual user defined cast operator to a std::variant?  Then you can use std::visit with your OOP class hierarchy.

By combining this with a trick of forward declaring a user class which turns out to derive from the class in std::, we can even hide the implementation (type list) of the variant.

This is something that is easier to just give the code, so here's a complete example:

[https://godbolt.org/z/TSCITM](https://godbolt.org/z/TSCITM)

But in case you don't want to read it all, here's the gist of it:

    // forward declared; variant typelist hidden.
    class FruitVariant;
    
    class Fruit 
    { 
    public:
    	// User-defined cast to a forward declared variant.
    	virtual operator FruitVariant() = 0;
    };
    
    class FruitVariant : public std::variant&lt;Apple*, Orange*&gt; {}
    
    Apple::operator FruitVariant() { return this; }
    Orange::operator FruitVariant() { return this; }
    
    void eat(Fruit&amp; fruit)
    {
    	// Fruit is convertible to std::variant via an operator
    	// (Unfortunately can't pass raw fruit directly to std::visit.)
    	std::variant&lt;Apple*, Orange*&gt; const&amp; fruit_var = fruit;
    		
    	std::visit(
    		overload {
    			[](Apple* ){ std::cout &lt;&lt; "Just bite in\n"; },
    			[](Orange*){ std::cout &lt;&lt; "Peel first\n;"; }
    		},
    		fruit_var
    	);
    }
## [5][Developer Ecosystem Survey 2020](https://www.reddit.com/r/cpp/comments/eytlnl/developer_ecosystem_survey_2020/)
- url: https://surveys.jetbrains.com/s3/a18-developer-ecosystem-survey-2020
---

## [6][C++ Weekly - Understanding C++ Lambdas Through Lambdas (online Course) - YouTube Playlist](https://www.reddit.com/r/cpp/comments/eyitr7/c_weekly_understanding_c_lambdas_through_lambdas/)
- url: https://www.youtube.com/watch?v=3hGSlUGEXtA&amp;list=PLs3KjaCtOwSY_Awyliwm-fRjEOa-SRbs-
---

## [7][KDevelop 5.5 released](https://www.reddit.com/r/cpp/comments/eye3r1/kdevelop_55_released/)
- url: https://www.kdevelop.org/news/kdevelop-550-released
---

## [8][Portable string SSO Challenge](https://www.reddit.com/r/cpp/comments/eyul4u/portable_string_sso_challenge/)
- url: https://www.reddit.com/r/cpp/comments/eyul4u/portable_string_sso_challenge/
---
Inspired by the post about libc++ std::string SSO

[https://www.reddit.com/r/cpp/comments/ey464c/libcs\_implementation\_of\_stdstring/](https://www.reddit.com/r/cpp/comments/ey464c/libcs_implementation_of_stdstring/)

&amp;#x200B;

Write and post a link to your implementation of a string with SSO. The caveat is that there has to be no undefined behavior (ie no union type punning, etc). You can assume a 64bit system and 24 byte string objects (8 bytes each for capacity, begin, end). Factors to consider will be the size of the short string that can be represented as well as "elegance".
## [9][ABI - Now or Never](https://www.reddit.com/r/cpp/comments/ey8y8j/abi_now_or_never/)
- url: https://wg21.link/P1863
---

## [10][What is ABI, and What Should WG21 Do About It?](https://www.reddit.com/r/cpp/comments/eyaee0/what_is_abi_and_what_should_wg21_do_about_it/)
- url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/p2028r0.pdf
---

## [11][Libc++’s implementation of std::string](https://www.reddit.com/r/cpp/comments/ey464c/libcs_implementation_of_stdstring/)
- url: https://joellaity.com/2020/01/31/string.html
---

