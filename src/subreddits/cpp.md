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
## [2][Trying to build an Open Source browser in 2020 - Patricia Aas](https://www.reddit.com/r/cpp/comments/j2ifjr/trying_to_build_an_open_source_browser_in_2020/)
- url: https://www.youtube.com/watch?v=a0aY3rOD0Ss
---

## [3][It Is Never a Compiler Bug Until It Is](https://www.reddit.com/r/cpp/comments/j2jf0s/it_is_never_a_compiler_bug_until_it_is/)
- url: http://r6.ca/blog/20200929T023701Z.html
---

## [4][Tool to merge multiple C/C++ files into a single source file](https://www.reddit.com/r/cpp/comments/j2hg59/tool_to_merge_multiple_cc_files_into_a_single/)
- url: https://www.reddit.com/r/cpp/comments/j2hg59/tool_to_merge_multiple_cc_files_into_a_single/
---
I've recently created a command line tool named **cpp-merge**. It can create a single C/C++ source file from multiple header and source files. It was developed mainly to use in programming contests in which solution must be submitted as a single source file. It has been written in TypeScript. You can install it using [NPM](https://www.npmjs.com/package/cpp-merge). Source code is available on [GitHub](https://github.com/RandomVoid/cpp-merge).
## [5][A simple library to count elements while iterating in for-range loop.](https://www.reddit.com/r/cpp/comments/j2it58/a_simple_library_to_count_elements_while/)
- url: https://www.reddit.com/r/cpp/comments/j2it58/a_simple_library_to_count_elements_while/
---
I often find myself wanting a counter while looping in a for(auto&amp; item : items). Introducing a new variable at the outer scope or rewriting for-range loop into the ordinary one was, frankly, annoying, thus this library was born https://github.com/itsuart/counted_range .

It let's you do this:
        
    std::vector&lt;int&gt; items{ 11, 22, 33, 44 };
    for (auto&amp; i : helpers::counted_range(items)) {
        std::cout &lt;&lt; i.count &lt;&lt; ": " &lt;&lt; i.value &lt;&lt; "\n";
    }

Output:
    
    0: 11
    1: 22
    2: 33
    3: 44

Starting counter value and step can be set via optional parameters.

The library is C++11 compatible, single header only, supports all kinds of iterables, and in public domain.


Since this is both my first submission to the subreddit, and first generic library intended to be used outside of my work projects, I'll be happy to receive any feedback :).


Thank you!
## [6][Confusion regarding co_yield. Does usage of co_yield always involve generators (I think it does) and if true, why does C++20 Standard not specify a generator?](https://www.reddit.com/r/cpp/comments/j2fpdr/confusion_regarding_co_yield_does_usage_of_co/)
- url: https://www.reddit.com/r/cpp/comments/j2fpdr/confusion_regarding_co_yield_does_usage_of_co/
---
(Edit:- Specify should be more like name or contain a Generator Class the way we have Vector/dequeue/exchange/move etc.)

So let me start off by saying, let's see a simple co\_yield example in MSVC

    std::experimental::generator&lt;std::size_t&gt; Generate(std::size_t const end){
       for(std::size_t i = 0; i &lt; end; ++i)
          co_yield i;
    }

So now, as you can imagine, MSVC provides an experimental/generator

I am not sure if anyone else provides this.

But I cannot find a generator in the C++ Standard.

So as a dev, are we expected to create our own generators?

(Because I would prefer using a standardized one in use cases like this)

The only Standard Generator Proposal I can find is [std::generator is by Lewis Baker from CppCoro](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/p2168r0.pdf)
## [7][Using a constexpr std::array as a list of non-template type parameters in C++20](https://www.reddit.com/r/cpp/comments/j25xgn/using_a_constexpr_stdarray_as_a_list_of/)
- url: https://www.reddit.com/r/cpp/comments/j25xgn/using_a_constexpr_stdarray_as_a_list_of/
---
I recently found myself in the position of needing to use a `constexpr` array as a list of NTTPs. This has always been a thing, as evidenced by e.g. these stack overflow questions: [1](https://stackoverflow.com/questions/24560547/c1y-c14-converting-static-constexpr-array-to-non-type-template-parameter-pa), [2](https://stackoverflow.com/questions/49516020/using-an-instance-of-an-stdarraytype-n-as-a-template-parameter?noredirect=1&amp;lq=1), [3](https://stackoverflow.com/questions/55603440/convert-static-constexpr-array-to-template-parameter-in-c11). However, the code I wrote to solve this problem turned out to be an excellent demonstration of a cool new feature of C++20: the expansion of what can be a template parameter. Namely, we can now use a `std::array` as a template parameter! Since I don't have a blog, I thought I'd share my thoughts here, perhaps someone will find them interesting.

First, let me provide a use case. Say we have the class

    template &lt; size_t NUM_WHEELS &gt;
    class Vehicle { /* ... */ };

And let's assume that we have some idea about how many wheels a vehicle can have

    constexpr std::array admissible_wheels = {2, 4};

We'd now like to store a number of vehicles in one container. To avoid dynamic allocation for each object, inheritance, etc., we want to use a `std::vector` of `std::variant`s. So we can write

    using Garage = std::vector&lt; std::variant&lt;
        Vehicle&lt;admissible_wheels[0]&gt;,
        Vehicle&lt;admissible_wheels[1]&gt;
    &gt; &gt;;

This seems... problematic. What if we want to add unicycles? And tuk-tuks? Maybe a 6-wheeled Mars rover? Clearly, we should employ a bit of meta-programming to automate things. If this example seems conceited, imagine you have a type which is parametrized using several nttps, and you'd like to get the variant of the type templated over all possible combinations of the nttps (my exact use case, I thought I'd spare the extra complexity in this example). At some point, manually typing everything out becomes unfeasible and you have to resort to meta-programming.

 Loosely speaking, we'd like to have a "meta-function", which takes an array and a type template, and returns the type of the template stamped out over all elements of the array. In pseudo-code, we'd like to have something like

    // Pseudo-code is putting it kindly
    vehicle_variant_t = apply_nttp(admissable_nwheels, std::variant&lt; Vehicle&lt; &gt; &gt;)
    
    // Now in real C++
    using Garage = std::vector&lt; vehicle_variant_t &gt;;

The proper way to do that is a template of a template and... what? We can't use `std::array` as a template parameter, right? Well, in C++20, we can! We can define the following alias:

    template &lt; auto Parameter_Array, template &lt; auto... &gt; typename Type_To_Parametrize &gt;
    using apply_nttp_t = /* implementation details */;

And then simply define

    template &lt; auto ...NW &gt;
    using vehicle_variant_t = std::variant&lt; Vehicle&lt; NW &gt;... &gt;;
    
    using Garage = std::vector&lt; apply_nttp_t&lt; admissible_wheels, vehicle_variant_t &gt; &gt;;

Now, whenever we want to modify the set of wheels a vehicle is allowed to have, we only need to modify one array, and the changes automatically propagate to the rest of the code! While we're at it, we might as well constrain the template described above using a concept of `std::array`:

    template &lt; typename T &gt;
    struct is_std_array : std::false_type
    {};
    
    template &lt; typename T, size_t N &gt;
    struct is_std_array&lt; std::array&lt; T, N &gt; &gt; : std::true_type
    {};
    
    template &lt; typename T &gt;
    concept std_array = is_std_array&lt; std::decay_t&lt; T &gt; &gt;::value;
    
    template &lt; std_array auto Parameter_Array,
               template &lt; typename std::decay_t&lt; decltype(Par_Array) &gt;::value_type... &gt;
               typename Type_To_Parametrize &gt;
    using apply_nttp_t = /* implementation details */;

And that's it. You can find the implementation details [here](https://godbolt.org/z/b7zKr4). I've even included the use of a lambda template parameter list for all you C++20 freaks out there. I hope this was at least a little interesting. Thanks for reading!
## [8][How to get clang_format not to put spaces first and last inside curly braces](https://www.reddit.com/r/cpp/comments/j2hpgu/how_to_get_clang_format_not_to_put_spaces_first/)
- url: https://www.reddit.com/r/cpp/comments/j2hpgu/how_to_get_clang_format_not_to_put_spaces_first/
---
I want this; 

`enum EnumType {value1, value2};`

and not this (which my clang-format produces);

`enum EnumType { value1, value2 };`

I have searched the clang-format style options without finding an option that does this. Does anybody know if it exist?

I suspect it's the same for function and method one-liners so it might be a block thing but I'm not sure.
## [9][Classically polymorphic visit replaces some uses of dynamic_cast](https://www.reddit.com/r/cpp/comments/j2gukw/classically_polymorphic_visit_replaces_some_uses/)
- url: https://quuxplusone.github.io/blog/2020/09/29/oop-visit/
---

## [10][[Research] Super-resolution ESRGAN model C++!](https://www.reddit.com/r/cpp/comments/j27id2/research_superresolution_esrgan_model_c/)
- url: /r/MachineLearning/comments/j27gj7/research_superresolution_esrgan_model_c/
---

## [11][Bjarne Stroustrup: The Beauty and Power of "Primitive" C++](https://www.reddit.com/r/cpp/comments/j1x14e/bjarne_stroustrup_the_beauty_and_power_of/)
- url: https://www.youtube.com/watch?v=ERzENfQ51Ck
---

