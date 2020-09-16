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
## [2][I made this ASCII Tetris clone demo in console](https://www.reddit.com/r/cpp/comments/itscqo/i_made_this_ascii_tetris_clone_demo_in_console/)
- url: https://youtu.be/_rQ3RHeVDRQ
---

## [3][CppScript: the evolution of modern C++](https://www.reddit.com/r/cpp/comments/itsj2n/cppscript_the_evolution_of_modern_c/)
- url: https://www.reddit.com/r/cpp/comments/itsj2n/cppscript_the_evolution_of_modern_c/
---
This is a thread following [this thread](https://www.reddit.com/r/cpp/comments/i0i0e4/constexpr_if_and_requires_expressions_changed/) and [this post](https://www.reddit.com/r/cpp/comments/i15g4p/c_has_become_easier_to_write_than_java/fzuq0jp/?context=3).

There's an emerging modern cpp subset, I personally call it "CppScript", that closely resembles the coding style of most popular dynamically typed languages. consider the following Python program:

    class TypeA:
        def f(self):
            return 'TypeA'
        def g(self):
            return 'Hello!'
        
    class TypeB:
        def f(self):
            return 'TypeB'
        
    def Test(obj):
        if hasattr(obj, 'g'):
            print(obj.f() + ' says ' + obj.g())
        else:
            print('g() is missing from ' + obj.f())
            
    Test(TypeA())
    Test(TypeB())

Below are the corresponding C++ programs that accurately replicate the behavior of the Python program above, in different versions of C++ (I have also tried to further downgrade from C++11 to C++98 but I failed because SFINAE seems exceptionally hard in C++98, I can't figure out the syntax to describe `decltype(&amp;T::g)` in C++98):

C++20: [https://godbolt.org/z/xG5v4s](https://godbolt.org/z/xG5v4s)

C++17: [https://godbolt.org/z/v9nY9P](https://godbolt.org/z/v9nY9P)

C++14: [https://godbolt.org/z/5Eh1Mv](https://godbolt.org/z/5Eh1Mv)

C++11: [https://godbolt.org/z/961dnY](https://godbolt.org/z/961dnY)

as you can see, the C++20 version is bit-identical (by "bit-identical", I mean each line of the C++ code **directly** corresponds to a line of the Python code above) to the Python code and as the C++ version gets older, the code gets more intricate and the coding style gradually deviates from the script looking Python/C++20 to more traditional template metaprogramming.

The CppScript subset of modern C++ is very easy to learn and write, you just need to follow these rules:

1. `auto` everywhere, this is not that much about "I just don't bother with typing the types", `auto` is essential to duck typing and letting your code ***automatically*** generalize to any compatible type.
2. favor duck typing (static polymorphism) over dynamic polymorphism (type erasure + virtual functions) whenever possible. This is because C++ is after all, statically typed, and the type information is either mostly erased (if RTTI is enabled or if there's any virtual function) or completely erased at runtime, therefore any functionality that relies on runtime type information will be very limited. many extremely powerful features including duck typing and dependent types are compile time only for C++ and you will want to take advantage of these features.
3. Don't overcomplicate things with a million classes and design patterns, use `requires` expressions and `constexpr if` if you want to query the behavior of something and change the behavior of your code accordingly.
4. Think in dynamic typing, but be aware that any type related stuff must be determined at compile time. You don't need a highly constrained mindset like you do for other less permissive statically typed languages like Java. C++ is very flexible at compile time, for instance, functions (function templates to be precise) could have return values of potentially different types, as long as the exact return type is reachable at compile time.

&amp;#x200B;

    template&lt;auto x&gt;
    auto f() {
        if constexpr (x == 0)
            return 42;
        else
            return 2.71;
    }
    f&lt;0&gt;(); // returns int
    f&lt;1&gt;(); // returns double

personally I don't see a huge gap between the CppScript subset and dynamically typed languages like Python for most use cases, and I think it's a great startpoint for anyone that struggles to learn C++, it is much simpler, easier and more intuitive than the rest of C++
## [4][CppCon: Try out CppCon 2020 on the cheap](https://www.reddit.com/r/cpp/comments/itl0l8/cppcon_try_out_cppcon_2020_on_the_cheap/)
- url: https://www.reddit.com/r/cpp/comments/itl0l8/cppcon_try_out_cppcon_2020_on_the_cheap/
---
CppCon 2020 is off to a great start in its "online venue."

If you've been curious about attending CppCon or attending an online C++ conference, now may be the time.

Since the conference is about half over, registrations are now available at half-off the regular $300 registration fee (with discount code).

To register at the $150 discounted registration fee use this registration code: [https://cppcon2020.eventbrite.com/?discount=TasteOf\_CppCon\_Reddit](https://cppcon2020.eventbrite.com/?discount=TasteOf_CppCon_Reddit)
## [5][Happy new Jewish year - A C++20 Jewish Date library](https://www.reddit.com/r/cpp/comments/itct0i/happy_new_jewish_year_a_c20_jewish_date_library/)
- url: https://www.reddit.com/r/cpp/comments/itct0i/happy_new_jewish_year_a_c20_jewish_date_library/
---
[https://github.com/royalbee/jewish\_date](https://github.com/royalbee/jewish_date)

A header only std::chrono date like library to handle [Jewish Dates](https://en.wikipedia.org/wiki/Hebrew_calendar).

    #include &lt;jewish/date.h&gt;
    
    using namespace jewish::literals;
    
    static_assert(5781_y / Tishrei / 1 == jewish::year_month_day(date::year(2020) / 9 / 19));
    static_assert(date::year_month_day(5781_y / Tishrei / 1) == date::year(2020) / 9 / 19);

Since I Couldn't find any standard library supporting chrono date it Currently uses [Howard Hinnant's date library](https://github.com/HowardHinnant/date).

&amp;#x200B;

Happy Healthy Peaceful Year
## [6][What are the advantages of a weak ownership model in modules?](https://www.reddit.com/r/cpp/comments/ita6xn/what_are_the_advantages_of_a_weak_ownership_model/)
- url: https://www.reddit.com/r/cpp/comments/ita6xn/what_are_the_advantages_of_a_weak_ownership_model/
---
As the title says, there are two different ownership models allowed by the standard in modules

1. Strong Ownership Model
2. Weak Ownership Model

What are the advantages of the weak ownership model as compared to strong ownership models?
## [7][Structure Padding in C++](https://www.reddit.com/r/cpp/comments/itv4zw/structure_padding_in_c/)
- url: https://thoughts-on-coding.com/2020/09/14/structure-padding-in-cpp/
---

## [8][Refactoring from single to multi purpose](https://www.reddit.com/r/cpp/comments/ituksf/refactoring_from_single_to_multi_purpose/)
- url: http://meetingcpp.com/blog/items/Refactoring-from-single-to-multi-purpose.html
---

## [9][Bjarne Stroustrup: C++ Today (2016)](https://www.reddit.com/r/cpp/comments/it54o4/bjarne_stroustrup_c_today_2016/)
- url: https://www.youtube.com/watch?v=aPvbxuOBQ70&amp;list=WL&amp;index=11&amp;t=0s
---

## [10][C++ in Visual Studio Code reaches version 1.0! | C++ Team Blog](https://www.reddit.com/r/cpp/comments/ista1w/c_in_visual_studio_code_reaches_version_10_c_team/)
- url: https://devblogs.microsoft.com/cppblog/c-in-visual-studio-code-reaches-version-1-0/
---

## [11][Standard C++20 Modules support with MSVC in Visual Studio 2019 version 16.8 | C++ Team Blog](https://www.reddit.com/r/cpp/comments/ispeiy/standard_c20_modules_support_with_msvc_in_visual/)
- url: https://devblogs.microsoft.com/cppblog/standard-c20-modules-support-with-msvc-in-visual-studio-2019-version-16-8/
---

