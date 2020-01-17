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
## [2][How can I handle both structured exceptions and C++ exceptions potentially coming from the same source?](https://www.reddit.com/r/cpp/comments/epwpx3/how_can_i_handle_both_structured_exceptions_and_c/)
- url: https://devblogs.microsoft.com/oldnewthing/20200116-00/?p=103333
---

## [3][Why were abbrev. lambdas rejected?](https://www.reddit.com/r/cpp/comments/epq4ui/why_were_abbrev_lambdas_rejected/)
- url: https://brevzin.github.io/c++/2020/01/15/abbrev-lambdas/
---

## [4][New Accuracy Improvements for Remote Linux IntelliSense in both CMake and VS projects | C++ Team Blog](https://www.reddit.com/r/cpp/comments/epvue6/new_accuracy_improvements_for_remote_linux/)
- url: https://devblogs.microsoft.com/cppblog/improvements-to-accuracy-and-performance-of-linux-intellisense/
---

## [5][CppCast: Conference Organizing](https://www.reddit.com/r/cpp/comments/epv9v3/cppcast_conference_organizing/)
- url: https://cppcast.com/cpp-conference-organizing/
---

## [6][I built a site to search C++ libraries based on the awesome C++ github list, with sort, filtering, and extra meta information for each library](https://www.reddit.com/r/cpp/comments/epikzx/i_built_a_site_to_search_c_libraries_based_on_the/)
- url: https://www.reddit.com/r/cpp/comments/epikzx/i_built_a_site_to_search_c_libraries_based_on_the/
---
 [https://lucidindex.com/cpp](https://lucidindex.com/cpp)
## [7][Software Development Blog](https://www.reddit.com/r/cpp/comments/epyj39/software_development_blog/)
- url: https://www.reddit.com/r/cpp/comments/epyj39/software_development_blog/
---
All, I have been working on my blog for over two years now. I switched to a C++ job little over a year ago and thus have been focusing on C++ a bit more.

If you are interested in reading some things. The topics are not very deep since I lack the proper knowledge about C++. There are also plenty of topic not related to C++.

Let me know what you think so I can grow.

[https://cvesters.wordpress.com](https://cvesters.wordpress.com)
## [8][Excel file to C++ mobile app??](https://www.reddit.com/r/cpp/comments/eq09k7/excel_file_to_c_mobile_app/)
- url: https://www.reddit.com/r/cpp/comments/eq09k7/excel_file_to_c_mobile_app/
---
Hi,

I create a really good financial excel file that have a lot of potential and I would like to create a mobile app with it. I am a beginner with C++ and really not that good with programming. I was wondering if it's the good way to create it by using C++ and also if there was a easier way to create it by just convert my excel file into a C++ language?   


If you guys have a good tutorial that I can follow or just give me good advice to start that project, it will help me a lot!   


Thank you!
## [9][C++ 2020 News](https://www.reddit.com/r/cpp/comments/epzaw9/c_2020_news/)
- url: https://cppcast.com/cpp-2020-news/
---

## [10][POD Struct Serialisation Using Type Punning](https://www.reddit.com/r/cpp/comments/epz9ex/pod_struct_serialisation_using_type_punning/)
- url: https://www.reddit.com/r/cpp/comments/epz9ex/pod_struct_serialisation_using_type_punning/
---
I wanted to figure out a way to serialise a POD struct into a vector of bytes. My application is to serialise some game state objects in my game so they can be saved to file, and at some point be retrieved again using the reverse process.

I'm using #pragma pack(1)to force alignment to 1 byte.

To do the serialisation, I'm simply casting the whole struct to void\*then to unsigned char\*and then using the resulting pointer to initialise the vector using make\_unique.

I'm relatively new to c++, and I'm wondering if it is safe to do these kind of things. Here are a few of my burning questions:

1. Is the struct memory layout guaranteed?
2. Is the struct size guaranteed given that I'm forcing packing to the nearest byte?
3. How might I handle when more complex data is inside the struct, e.g. a vector or another struct? I guess I'd need to do some kind of recursive serialisation in that case.
4. Any neat preprocessor tricks that might help me flip the endianness with little effort?
5. Would it be safe to go in the opposite direction; deserialising by casting an unsigned char\[\]back to the struct type?

Code below is an example of what I'm talking about. Runs in [Godbolt](https://godbolt.org/z/yoUxfZ). Clone of my question on [Stackoverflow](https://stackoverflow.com/questions/59786566/pod-struct-serialisation-using-type-punning).

    #include &lt;iostream&gt;
    #include &lt;memory&gt;
    #include &lt;vector&gt;
    
    #pragma pack(1)
    struct PODStruct {
        unsigned char a;
        unsigned int b;
        unsigned int c;
    
        std::unique_ptr&lt;std::vector&lt;unsigned char&gt;&gt; serialise() const;
    };
    #pragma pack()
    
    std::unique_ptr&lt;std::vector&lt;unsigned char&gt;&gt; PODStruct::serialise() const {
        auto rawChar = (unsigned char*)((void*)this);
        auto serDataPtr = std::make_unique&lt;std::vector&lt;unsigned char&gt;&gt;(rawChar, rawChar + sizeof(PODStruct));
        return serDataPtr;
    }
    
    int main() {
        PODStruct foo = {0x38, 0x34353637, 0x30313233};
    
        auto bar = foo.serialise();
    
        for (auto byte : *bar) {
            std::cout &lt;&lt; byte &lt;&lt; std::endl; 
        }
    
        std::cout &lt;&lt; "Size (bytes): " &lt;&lt; bar-&gt;size() &lt;&lt; std::endl;
        return 0;
    }
## [11][C++ Objects, their lifetimes and pointer to objects](https://www.reddit.com/r/cpp/comments/epgxj6/c_objects_their_lifetimes_and_pointer_to_objects/)
- url: https://blog.panicsoftware.com/objects-their-lifetimes-and-pointers/
---

