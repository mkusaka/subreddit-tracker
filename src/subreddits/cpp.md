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
## [2][Here I Stand, Free - Allocators and an Inclusive STL](https://www.reddit.com/r/cpp/comments/idtnn7/here_i_stand_free_allocators_and_an_inclusive_stl/)
- url: https://thephd.github.io/freestanding-noexcept-allocators-vector-memory-hole
---

## [3][CppCast: STX](https://www.reddit.com/r/cpp/comments/ido5v4/cppcast_stx/)
- url: https://cppcast.com/stx-basit-ayantunde/
---

## [4][Calling Functions: A Tutorial - Munich C++ User Group 2020 - Klaus Iglberger](https://www.reddit.com/r/cpp/comments/idt9dw/calling_functions_a_tutorial_munich_c_user_group/)
- url: https://www.youtube.com/watch?v=B9RT5sVunmk
---

## [5][[Article]But I was helping the compiler!](https://www.reddit.com/r/cpp/comments/idg1vd/articlebut_i_was_helping_the_compiler/)
- url: https://pankajraghav.com/2020/08/16/RVO.html
---

## [6][Dear PyGui - A Simple Python GUI Library built with Dear ImGui in C++](https://www.reddit.com/r/cpp/comments/idbi9j/dear_pygui_a_simple_python_gui_library_built_with/)
- url: https://www.reddit.com/r/cpp/comments/idbi9j/dear_pygui_a_simple_python_gui_library_built_with/
---
I'd like to show case our python GUI library we built on top of Dear ImGui. We'd love some feedback and love from the community. We are in a weird place because its a C++ project built for python users. This is my first large open source project.

https://github.com/hoffstadt/DearPyGui
## [7][Inside the Microsoft STL: The std::exception_ptr](https://www.reddit.com/r/cpp/comments/idijf4/inside_the_microsoft_stl_the_stdexception_ptr/)
- url: https://devblogs.microsoft.com/oldnewthing/20200820-00/?p=104097
---

## [8][neix - a news reader for your terminal](https://www.reddit.com/r/cpp/comments/idi1ik/neix_a_news_reader_for_your_terminal/)
- url: https://github.com/tomschwarz/neix
---

## [9][Rust and C++ interoperability - Chromium Projects](https://www.reddit.com/r/cpp/comments/idfbq5/rust_and_c_interoperability_chromium_projects/)
- url: https://www.chromium.org/Home/chromium-security/memory-safety/rust-and-c-interoperability
---

## [10][bsc::parser - C++20 command line arguments parser](https://www.reddit.com/r/cpp/comments/id8mq0/bscparser_c20_command_line_arguments_parser/)
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
## [11][Announcing Meeting C++ online!](https://www.reddit.com/r/cpp/comments/idbeyk/announcing_meeting_c_online/)
- url: https://meetingcpp.com/meetingcpp/news/items/Announcing-Meeting-Cpp-online-.html
---

