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
## [2][A Buffers Library for C++20: Part 1](https://www.reddit.com/r/cpp/comments/ijatqb/a_buffers_library_for_c20_part_1/)
- url: https://vector-of-bool.github.io/2020/08/29/buffers-1.html
---

## [3][Why I like C++ attributes](https://www.reddit.com/r/cpp/comments/ijas18/why_i_like_c_attributes/)
- url: https://mariusbancila.ro/blog/2020/08/30/why-i-like-cpp-attributes/
---

## [4][Question regarding optional virtual destructor in C++20](https://www.reddit.com/r/cpp/comments/ijbaf5/question_regarding_optional_virtual_destructor_in/)
- url: https://www.reddit.com/r/cpp/comments/ijbaf5/question_regarding_optional_virtual_destructor_in/
---
Given

    template&lt;typename T, bool needVirtual&gt;
    struct C
    {
        virtual ~C() requires(needVirtual)
        {}
    
        ~C() = default;
    };
    
    int main()
    {
        C&lt;int, false&gt; c1;
        C&lt;int, true&gt; c2;
    
        return 0;
    }

([https://godbolt.org/z/oa4vjc](https://godbolt.org/z/oa4vjc))

I would expect `c1` to not have virtual destructor and, well, to not be a virtual class.

Currently however **gcc** does create a virtual class:

    vtable for C&lt;int, false&gt;:
            .quad   0
            .quad   0

And **clang** seems to fail to compile:

    error: virtual function cannot have a requires clause

What is the correct behavior according to the standard?    
Thank you!    
P.S. I know one can have the desired behavior with conditional base classes.
## [5][Uberswitch: a header-only, unobtrusive, almighty alternative to the C++ switch statement that looks just like the original but can do so much more.](https://www.reddit.com/r/cpp/comments/iithnh/uberswitch_a_headeronly_unobtrusive_almighty/)
- url: https://www.reddit.com/r/cpp/comments/iithnh/uberswitch_a_headeronly_unobtrusive_almighty/
---
TL;DR: here's the code: [https://github.com/falemagn/uberswitch](https://github.com/falemagn/uberswitch)

It was one boring day of a few years ago at work in which I found myself copying and pasting lots of if-then-else's and wished I'd have the ability to "switch over any type I wanted".

I knew other people had given a go at it, but all the solutions I could find were not of my liking (ugly syntax, lacking functionalities, too many compromises, etc.), so I thought  I might have some fun trying to invent my own.

It turns out the "holy grail" of the switches, the string switch, was totally achievable.

    int string2num(std::string s) {
        uberswitch (s) {
            case ("one"):
                return 1;
    
            case ("two"):
                return 2;
    
            case ("three"):
                return 3;
    
            // fallthrough works too
            case ("four"):
            case ("f0ur"):
                return 4
    
            default: return -1;
        }
    }

And just like one can use it to switch over a string, it can also be used to switch over an object of any other type, or a sequence of objects of any other type,  for which the `operator==()` is properly implemented (see the [README.md](https://github.com/falemagn/uberswitch/blob/master/README.md) for more examples).

So, I've had this thing sitting on github for about 2 years and in production for more than 3 years, and given the talks that there are about making [pattern matching](https://github.com/mpark/patterns) part of the language, I thought I might polish the code, get rid of some parts that were overkill/badly designed and integrate [fameta::counter](https://github.com/falemagn/fameta-counter) into it, to achieve the ability to nest `uberswitch` statemements within one another.

And here it is: [https://github.com/falemagn/uberswitch](https://github.com/falemagn/uberswitch)

The tool uses the preprocessor in a way that shouldn't give anybody goose bumps and gcc is able to optimize the code down to something which seems to be at least on par with equivalent code that doesn't make use of `uberswitch`. Clang comes close second in the optimization race. See [on godbolt](https://godbolt.org/z/WE3E1q)

PS: originally the code had a version of \`uberswitch\` that allowed to build a static map of to-be-matched-with items, but that part was to be rethought. Perhaps with your help?
## [6][pclmulqdq Tricks](https://www.reddit.com/r/cpp/comments/ijcgw3/pclmulqdq_tricks/)
- url: https://wunkolo.github.io/post/2020/05/pclmulqdq-tricks/
---

## [7][inheritance in C++](https://www.reddit.com/r/cpp/comments/ijc7le/inheritance_in_c/)
- url: https://www.reddit.com/r/cpp/comments/ijc7le/inheritance_in_c/
---
Is it possible to call a virtual function in child class?? for more clearance of the question ,find the below code.

class Person{

private:

int age;

string name;

public:

virtual void getdata()

{

cin&gt;&gt;name&gt;&gt;age;   }

virtual void putdata()

{

   cout&lt;&lt;name&lt;&lt;" "&lt;&lt;age;   }        };

class student: public Person

{

int cur\_id;

public:

void getdata()

{

getdata(); 

cin&gt;&gt;cur\_id;    }

void putdata()  {

putdata();

cout&lt;&lt;cur\_id;    }      };

int main()  {...}  
 

&amp;#x200B;

in the main function i make a ptr of base class, is it correct to call a virtual function in the child class like this???
## [8][Matplot++: A C++ Graphics Library for Data Visualization](https://www.reddit.com/r/cpp/comments/iikdwq/matplot_a_c_graphics_library_for_data/)
- url: https://www.reddit.com/r/cpp/comments/iikdwq/matplot_a_c_graphics_library_for_data/
---
Data visualization can help programmers and scientists identify trends in their data and efficiently communicate these results with their peers. Modern C++ is being used for a variety of scientific applications, and this environment can benefit considerably from graphics libraries that attend the typical design goals toward scientific data visualization. Besides the option of exporting results to other environments, the customary alternatives in C++ are either non-dedicated libraries that depend on existing user interfaces or bindings to other languages. Matplot++ is a graphics library for data visualization that provides interactive plotting, means for exporting plots in high-quality formats for scientific publications, a compact syntax consistent with similar libraries, dozens of plot categories with specialized algorithms, multiple coding styles, and supports generic backends.

&amp;#x200B;

[https://github.com/alandefreitas/matplotplusplus](https://github.com/alandefreitas/matplotplusplus)
## [9][Custom 'diagnostics' for concepts](https://www.reddit.com/r/cpp/comments/iixi6x/custom_diagnostics_for_concepts/)
- url: https://www.reddit.com/r/cpp/comments/iixi6x/custom_diagnostics_for_concepts/
---
I'd like to 'propose' the following simple attribute for concepts:

        template &lt;typename T&gt; [[requirement_failed("reason")]]
        concept MyConcept = /*some expression*/;
        
        template&lt;typename T&gt;
        concept MyConcept2 = requires(T t) {    
            /*some expression*/;
        } [[requirement_failed("reason")]];

The purpose should be quite obvious: "reason" is issued as a compiler note, when the requirements are not met.

Notes: 

There is [P1267R0](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2018/p1267r0.pdf), but I think it's just not put or presented it in a simple enough way. Allowing the attribute on any definition that 'requires' will likely result in a mess.

Allowing it *on concept definitions only* will keep things clean. It avoids duplication for various functions making use of the same concept. It avoids over-use of the feature, as there should not be an excessive amount of concepts.

In my view, concepts that are part of an interface should offer 'diagnostic' messages as a best practice.

I don't think I have the guts, insight or time to make this into an official proposal, even if it is a very simple thing. So I thought, I'd just mention it here. It's also hard to imagine that something similar has not been considered yet, by some concepts committee.

Anyway, I wanted to draw some attention to it.
## [10][Lithium is now the fastest web framework (techempower.com)](https://www.reddit.com/r/cpp/comments/iiqnz5/lithium_is_now_the_fastest_web_framework/)
- url: https://www.reddit.com/r/cpp/comments/iiqnz5/lithium_is_now_the_fastest_web_framework/
---
Link to the benchmark: [https://www.techempower.com/benchmarks/#section=test&amp;runid=57b25c85-082a-4013-b572-b0939006eaff&amp;hw=ph&amp;test=composite&amp;a=2](https://www.techempower.com/benchmarks/#section=test&amp;runid=57b25c85-082a-4013-b572-b0939006eaff&amp;hw=ph&amp;test=composite&amp;a=2) 

Link to lithium: [https://github.com/matt-42/lithium/](https://github.com/matt-42/lithium/)
## [11][Phil Nash interview about Accelerated Test Driven Design For More Productive C++ prior CppCon Academy](https://www.reddit.com/r/cpp/comments/ij7hat/phil_nash_interview_about_accelerated_test_driven/)
- url: https://www.reddit.com/r/cpp/comments/ij7hat/phil_nash_interview_about_accelerated_test_driven/
---
[https://cppcon.org/instructor-interview-phil-nash-accelerated-test-driven-design/](https://cppcon.org/instructor-interview-phil-nash-accelerated-test-driven-design/)
