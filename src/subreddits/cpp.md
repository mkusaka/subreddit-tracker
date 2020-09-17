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
## [2][I'm pretty sure there's a massive, widespread misconception about std::random (most particularly mt19937), and that it's actually trivially easy to use.](https://www.reddit.com/r/cpp/comments/iufxze/im_pretty_sure_theres_a_massive_widespread/)
- url: https://www.reddit.com/r/cpp/comments/iufxze/im_pretty_sure_theres_a_massive_widespread/
---
I hear several things parroted over and over about mt19937. Some points are fair (it has a large internal state, which can be overkill in certain applications like embedded computing), some I can't be sure of (PRNG experts have purportedly found flaws in it in the time since its conception), and I'd like some clarification on that (i.e. further reading and specification of what cases it's now considered unsuitable for).

--But some points made don't gibe with my readings on cppreference.com. Perhaps a bit of code will help clarify what I mean:

	auto random_point_on_a_circle(double radius = 1.0) -&gt; std::array&lt;double, 2&gt;
	{
		using namespace std;
		
		static random_device rd{};
		static seed_seq seeder {rd(), rd(), rd(), rd(), 0xDEAD'FACE}; // 128 bits and a bit of salt.
		static mt19937 rand_eng {seeder};
		
		static uniform_real_distribution choose_angle {0.0, 2.0 * numbers::pi};
		
		auto angle = choose_angle(rand_eng);
		return {cos(angle) * radius, sin(angle) * radius};
	}

Here, I've decided that one "state collision" in 2^128 is plenty of distance between hypothetical runs, so I call upon random_device to supply 32 bits to my seed sequence 4 times. For good measure, in case random_device is deterministic and a user recognizes the output, I tack on a silly magic number term with no semantic meaning. If you like, chuck on `std::chrono::steady_clock::now().count()` (modulo 2^32 is implicit, maybe. I'd have to play around/look deeper).

The constructor for mt19937 then takes that seed sequence and calls on `seed_seq.generate()` for us, with the proper values required for its particular needs in order to eliminate any practical chance that its internal state is "degenerate", e.g. containing mostly null or uninitialized bytes.

Then we use the vastly superior `std::distribution`s, saying a permanent goodbye to all that rampant `rand() % 100` garbage that was never very elegant.

Alternatively, if we hate typing `rd()` multiple times, a seed_seq can be constructed by passing in *any* pair of iterators that dereference to an int. Notably, they can be forward-only, and no assumption is made that state is preserved after an increment. You can wrap `rd()` in a loop. You can use a string, standard input (just mash your keyboard), some low bits from your camera (if you know how), etc. `std::seed_seq` just uses the lowest 32 bits of each de-referenced element.

Alternatively, if we don't care about seeding the engine at all (e.g. we're just setting values in a game's particle system), mt19937 has a default constructor with a default seed.

	// Literally just this, no need for a random_device and a seed_seq
	static mt19937 rand_eng {};
## [3][I just started a Highschool programming course which focuses primarily on C++ so I’m very new to all this, and I just want to ask what some of the possibilities of C++ are](https://www.reddit.com/r/cpp/comments/iubwvb/i_just_started_a_highschool_programming_course/)
- url: https://www.reddit.com/r/cpp/comments/iubwvb/i_just_started_a_highschool_programming_course/
---

## [4][Any GSOC Alternatives Focused on C++](https://www.reddit.com/r/cpp/comments/iujdkj/any_gsoc_alternatives_focused_on_c/)
- url: https://www.reddit.com/r/cpp/comments/iujdkj/any_gsoc_alternatives_focused_on_c/
---
Hi,

Are there any organisations, or maybe a list somewhere, for organisations looking for C++ developers to be mentored / volunteer? I know Google Summer of Code is a good one, but they are not C++ focused, and are for uni students. Id also just be interested in maybe a IRC / Discord channel where people get C++ stuff built, or maybe some other active C++ community.

Some background below as to why I ask this, and perhaps someone knows better alternatives - I know, this is not cscareerquestions...

I'll be honest, I am career changer, and have no commercial C++ experience (though used it a lot during my grad degree). I see a lot of jobs and their job requirements are "can seepeepee anywhere, has been doing so for 5+ years". Now obviously I lack the experience and so keep getting rejected; my plan was to maybe just volunteer myself and learn and build experience that way.
## [5][Two new ticket types for Meeting C++!](https://www.reddit.com/r/cpp/comments/iuf512/two_new_ticket_types_for_meeting_c/)
- url: https://meetingcpp.com/meetingcpp/news/items/Two-new-ticket-types-for-Meeting-Cpp-.html
---

## [6][I made this ASCII Tetris clone demo in console](https://www.reddit.com/r/cpp/comments/itscqo/i_made_this_ascii_tetris_clone_demo_in_console/)
- url: https://youtu.be/_rQ3RHeVDRQ
---

## [7][Reading large code bases](https://www.reddit.com/r/cpp/comments/iu05bb/reading_large_code_bases/)
- url: https://www.reddit.com/r/cpp/comments/iu05bb/reading_large_code_bases/
---
Hello,

I'm using C++ since a few years now and I know and understand most of the C++11 Standard. I would call myself decently competent and experienced in writing C++ programs from scratch. But now, I'm working on someone else's code for the first time in my life (680 k lines of code and 270 k lines of header code, not that much documentation, called OpenFOAM). I still can write self-contained code, but I'm completely unable to read and understand that code and writing new code that interacts with the old code is extremely hard for me. What is the best way to learn to read and understand large codebases?
## [8][is the {fmt} library and std::format from c++20 the same?](https://www.reddit.com/r/cpp/comments/ityf5m/is_the_fmt_library_and_stdformat_from_c20_the_same/)
- url: https://www.reddit.com/r/cpp/comments/ityf5m/is_the_fmt_library_and_stdformat_from_c20_the_same/
---
was {fmt} adopted by c++20 like some boost libraries get adopted into the standard? they seem very close when it comes to using.
 
and if it was, can i expect the same performance from std::format implementation as the {fmt} library?
 
thanks in advance.
## [9][Utility Library for Imaging System : ULIS - C++ graphic libary](https://www.reddit.com/r/cpp/comments/ity88h/utility_library_for_imaging_system_ulis_c_graphic/)
- url: https://www.reddit.com/r/cpp/comments/ity88h/utility_library_for_imaging_system_ulis_c_graphic/
---
Hi people ! I just wanted to share a graphic library developed by Praxinos team, named ULIS.  We developed this library to answer specific needs that were not fullfilled by other graphic libraries (OpenCV, littleCMS, etc.)

ULIS reunites a large set of algorithm to manage color models, color spaces, image files formats, blending modes (like in Photoshop), animated sequences, etc. It is written in C++, it is optimized and well-documented.

You can access the sources for free and, its use is also free-of-charge for non-commercial purposes only. Hopefully this library can help you to make your own plugins or software.

Here is the link to Github : [https://github.com/Praxinos/ULIS](https://github.com/Praxinos/ULIS)

If you have any question, feel free answer this thread or contact us on our website :[https://praxinos.coop/contact.php](https://praxinos.coop/contact.php)
## [10][CppScript: the evolution of modern C++](https://www.reddit.com/r/cpp/comments/itsj2n/cppscript_the_evolution_of_modern_c/)
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
## [11][CppCon: Try out CppCon 2020 on the cheap](https://www.reddit.com/r/cpp/comments/itl0l8/cppcon_try_out_cppcon_2020_on_the_cheap/)
- url: https://www.reddit.com/r/cpp/comments/itl0l8/cppcon_try_out_cppcon_2020_on_the_cheap/
---
CppCon 2020 is off to a great start in its "online venue."

If you've been curious about attending CppCon or attending an online C++ conference, now may be the time.

Since the conference is about half over, registrations are now available at half-off the regular $300 registration fee (with discount code).

To register at the $150 discounted registration fee use this registration code: [https://cppcon2020.eventbrite.com/?discount=TasteOf\_CppCon\_Reddit](https://cppcon2020.eventbrite.com/?discount=TasteOf_CppCon_Reddit)
