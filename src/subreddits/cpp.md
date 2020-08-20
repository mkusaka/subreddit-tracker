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
## [2][bsc::parser - C++20 command line arguments parser](https://www.reddit.com/r/cpp/comments/id8mq0/bscparser_c20_command_line_arguments_parser/)
- url: https://www.reddit.com/r/cpp/comments/id8mq0/bscparser_c20_command_line_arguments_parser/
---
I'd like to present a small library from my bsc project. It can be used to parse command line arguments by defining a simple struct. It's a wrapper aroung GNU argp, but I didn't like how you have to manually define callbacks there. I also didn't like how boost::program\_options or other libraries where defining things, so I wrote my own thing.

[https://github.com/stilgarpl/bsc/tree/master/src/parser](https://github.com/stilgarpl/bsc/tree/master/src/parser)

Fields in the struct hold the information if the field was set and how many times it was set (for things like -vvvvv) and of course the value. operator() returns std::optional&lt;T&gt; for Parameter&lt;T&gt; and const T&amp; for DefaultParameter&lt;T&gt; (which always have value)

A sample of how it works:

    struct HelloWorldParameters : CommandLineParameters {
        Flag flag                    = {{.shortKey = 'f', .longKey = "flag", .doc = "Flag"}};
        DefaultParameter&lt;int&gt; number = {{.shortKey = 'i', .longKey = "number", .argumentName = "INT", .doc = "An integer number", .defaultValue = 5}};
        Parameter&lt;float&gt; realNumber = {{.shortKey = 'F', .longKey = "float", .argumentName="FLOAT", .doc = "A floating point number"}};
        Group g                     = {"Path group:"};
        DefaultParameter&lt;std::map&lt;short, std::filesystem::path&gt;&gt; numbersToPathsMap = {{.longKey = "map", .argumentName="MAP", .doc = "Numbers to path map"}};
        Argument&lt;std::string&gt; stringArgument;
        Argument&lt;int&gt; intArgument;
    };
    
    int main(int argc, char* argv[]) {
        const auto&amp; params = CommandLineParser::defaultParse&lt;HelloWorldParameters&gt;(argc, argv);
        
            if (params.flag()) {
                std::cout &lt;&lt; "Flag is set" &lt;&lt; std::endl;
            }
        
            std::cout &lt;&lt; "Number is " &lt;&lt; std::to_string(params.number()) &lt;&lt; std::endl;
            if (params.realNumber()) {
                std::cout &lt;&lt; "Real number was set and it is: " &lt;&lt; std::to_string(*params.realNumber()) &lt;&lt; std::endl;
            } else {
                std::cout &lt;&lt; "Real number was not set" &lt;&lt; std::endl;
            }
        
            std::cout &lt;&lt; "Path map:  "&lt;&lt; std::endl;
            for (const auto&amp; [key, value] : params.numbersToPathsMap()) {
                std::cout &lt;&lt; std::to_string(key) &lt;&lt; "=[" &lt;&lt; value.string() &lt;&lt; "]"&lt;&lt; std::endl;
            }
        
            std::cout &lt;&lt; "Required argument 1 (string): " &lt;&lt; params.stringArgument() &lt;&lt; std::endl
                      &lt;&lt; "Required argument 2 (int): " &lt;&lt; params.intArgument() &lt;&lt; std::endl;
        return 0;
    }

Of course, this project isn't finished yet, so things may still evolve, may maybe someone will find it useful.
## [3][awesome-hpp: A curated list of awesome header-only C++ libraries](https://www.reddit.com/r/cpp/comments/icm8g1/awesomehpp_a_curated_list_of_awesome_headeronly_c/)
- url: https://github.com/p-ranav/awesome-hpp
---

## [4][Regression in gcc's memory allocation elision optimization?](https://www.reddit.com/r/cpp/comments/ics6dj/regression_in_gccs_memory_allocation_elision/)
- url: https://www.reddit.com/r/cpp/comments/ics6dj/regression_in_gccs_memory_allocation_elision/
---
I recently re-watched Jason Turner's *C++ Code Smells* talk, and in it he has an example of declaring a string which, when cleaned up, essentially boils down to

    #include &lt;string&gt;

    int main() {
        std::string const greet1 = "Hello";
        std::string const greet2 = ", world!";
    }

At that time, he was using gcc-9 in godbolt, and the codegen was just two instructions: all of the memory allocation business in `std::string` was elided. I was very excited as clang has had this for some time. I tried it again in [godbolt](https://godbolt.org/z/vb9oKq) with gcc-10.X, and the elisions were no longer working!

I tried searching the gcc Bugzilla, but I couldn't find anything that quite matched (admittedly, my Bugzilla-search-foo may well be insufficient). Is this a known regression?

Particularly confounding is that it *does* work in gcc-10.X when you have only one string object, and it *doesn't* work in any gcc when you have three or more string objects. clang handles all of these cases. This suggests it may be a QoI issue.
## [5][The implication of const or reference member variables in C++](https://www.reddit.com/r/cpp/comments/icw0gk/the_implication_of_const_or_reference_member/)
- url: https://lesleylai.info/en/const-and-reference-member-variables/
---

## [6][A Rust-style guide for C++](https://www.reddit.com/r/cpp/comments/ickurq/a_ruststyle_guide_for_c/)
- url: https://www.reddit.com/r/cpp/comments/ickurq/a_ruststyle_guide_for_c/
---
Hello,

I have recently been perusing the Rust language book and found it really well-organized and presenting all the important topics concisely and ready-to-go.

Is there something similar for C++ language? I know there are some WG movements to create the ONE book for learning C++ as many other modern languages do. However, I feel that for now apart from some well-written books from gurus like Scott Meyers, there is a RTFM approach to learning this language. Or am I wrong?  


I would be happy to discuss this topic in detail, because I believe that without such book C++ is quite unapproachable for newcomers.
## [7][Writing ideal class definition in C++](https://www.reddit.com/r/cpp/comments/icijxj/writing_ideal_class_definition_in_c/)
- url: https://www.reddit.com/r/cpp/comments/icijxj/writing_ideal_class_definition_in_c/
---
 [https://medium.com/a-devs-life/designing-an-ideal-class-in-c-d205516c03ab](https://medium.com/a-devs-life/designing-an-ideal-class-in-c-d205516c03ab) 

Guys, please provide your valuable comments.
## [8][VS 2019 16.7.2 is now available](https://www.reddit.com/r/cpp/comments/icabt5/vs_2019_1672_is_now_available/)
- url: https://github.com/microsoft/STL/wiki/Changelog#vs-2019-167
---

## [9][sol 3: modern lua and C++ integration](https://www.reddit.com/r/cpp/comments/ic8ru0/sol_3_modern_lua_and_c_integration/)
- url: https://sol2.readthedocs.io/en/latest/
---

## [10][libstdc++ &lt;algorithm&gt; header transitively including &lt;range&gt;](https://www.reddit.com/r/cpp/comments/ibz0m2/libstdc_algorithm_header_transitively_including/)
- url: https://www.reddit.com/r/cpp/comments/ibz0m2/libstdc_algorithm_header_transitively_including/
---
I noticed some time ago that in c++20 mode the &lt;algorithm&gt; header in libstdc++ is substantially bigger than in c++11 mode :

`echo "#include &lt;algorithm&gt;" | gcc -std=c++11 -P -E -x c++ - | wc -l`

evaluates to 11760 loc, while

`echo "#include &lt;algorithm&gt;" | gcc -std=c++20 -P -E -x c++ - | wc -l`

evaluates to 45219 loc. Clicking through the header apparently in c++20 we have the include chain

&lt;algorithm&gt; -&gt; &lt;bits/ranges\_algo.h&gt; -&gt; &lt;bits/ranges\_algobase.h&gt; -&gt; &lt;range&gt;.

How could that happen? Measuring the compile time of file only including the &lt;algorithm&gt; header I get with the c++11 switch 120 ms and with the c++20 switch **600 ms** (empty file is 100 ms). I a bit baffled that while standardizing the &lt;range&gt; header this either slipped through the cracks or it was actively decided that this is ok. The &lt;algorithm&gt; header is virtually anywhere and with c++20 flag enabled every file including the &lt;algorithm&gt; header will compile half a second slower...

&amp;#x200B;

**Edit:**

Apparently the &lt;vector&gt; header also increased by a huge amount from 9530 loc to 19818 loc for libstdc++ due to now including &lt;bits/stl\_algo.h&gt; (this seems to be a known [issue](https://gcc.gnu.org/bugzilla/show_bug.cgi?id=92546) though). Compiling an empty header file only including &lt;vector&gt; is now up to 230 ms from 130 ms. 
## [11][C++20 SQLite wrapper](https://www.reddit.com/r/cpp/comments/ic629n/c20_sqlite_wrapper/)
- url: https://www.reddit.com/r/cpp/comments/ic629n/c20_sqlite_wrapper/
---
I am writing a C++20 SQLite wrapper with expressive code in mind. The code must run fast.

    sql::open("dev.db")
    | sql::query("select name, salary from person")
    | sql::for_each([](std::string_view name, float salary)
      { std::cout &lt;&lt; name &lt;&lt; "," &lt;&lt; salary &lt;&lt; std::endl; })
    | sql::onerror([](auto e){ std::cout &lt;&lt; e &lt;&lt; std::endl; });

I have some benchmarks to compare with solutions using the SQLite library. The usage of pipe operators to chain operations is optional and there is an API that throws C++ exceptions instead of using `result&lt;T&gt;` to report errors. 

[github.com/ricardocosme/msqlite](https://github.com/ricardocosme/msqlite)

What do you think? Any comments or feedbacks are welcome!
