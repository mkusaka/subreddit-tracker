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
## [2][On the relevance of debug performance](https://www.reddit.com/r/cpp/comments/iuzcha/on_the_relevance_of_debug_performance/)
- url: https://www.reddit.com/r/cpp/comments/iuzcha/on_the_relevance_of_debug_performance/
---
I had heard that some people really care about debug performance and out of pure curiosity I've recently been investigating code that is generated at -O0 and -Og and found the generated assembly to be shockingly bad.
I have a few simple questions based on that observation:

- Does it actually matter to you?
- Has this been a problem in your code base?
- What workarounds do you use?
- What is your optimization level in debug mode?

Edit: If you use -O0, have you tried -Og? If so, why do you prefer -O0?
## [3][CppCast: Microsoft Announcements at CppCon 2020](https://www.reddit.com/r/cpp/comments/iuyfyk/cppcast_microsoft_announcements_at_cppcon_2020/)
- url: https://cppcast.com/msvc-cppcon-2020/
---

## [4][Project-based C++ learning](https://www.reddit.com/r/cpp/comments/iukh9r/projectbased_c_learning/)
- url: https://learncppthroughprojects.com
---

## [5][Clang Format Extended, A Clang-Format Interface with support for taking Blobs as input and ignoring files in .gitignore and .clang-format-ignore](https://www.reddit.com/r/cpp/comments/iuxqfl/clang_format_extended_a_clangformat_interface/)
- url: https://www.npmjs.com/package/clang-format-ex
---

## [6][What are the problems with modules in C++20?](https://www.reddit.com/r/cpp/comments/iukqju/what_are_the_problems_with_modules_in_c20/)
- url: https://www.reddit.com/r/cpp/comments/iukqju/what_are_the_problems_with_modules_in_c20/
---
I'm new to C++, but one of things I've heard being talked about most is the new feature called modules in C++, which apparently are able to reduce compilation times and clean up code.

However, they're also a lot of posts about why C++ modules are actually bad, such as

https://vector-of-bool.github.io/2019/01/27/modules-doa.html
https://izzys.casa/2017/10/millennials-are-killing-the-modules-ts/

and a few other reddit posts. However, I can't really understand what all these posts are saying, so can somebody summarize all the perceived problems with modules here please?
## [7][I'm pretty sure there's a massive, widespread misconception about std::random (most particularly mt19937), and that it's actually trivially easy to use.](https://www.reddit.com/r/cpp/comments/iufxze/im_pretty_sure_theres_a_massive_widespread/)
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
## [8][Creating type traits using ADL (proof of concept with traits for bitmasks)](https://www.reddit.com/r/cpp/comments/iuojms/creating_type_traits_using_adl_proof_of_concept/)
- url: https://www.reddit.com/r/cpp/comments/iuojms/creating_type_traits_using_adl_proof_of_concept/
---
I got an idea to try and make type traits using a function found with ADL. So that way you could kind of get around the limitation of not being able to specialize templates in a non-enclosing namespace. The cool thing you can with this is that you can specialize, for example, bitmask operators for a scoped enum, and then use them in the namespace you defined them in without having to break out of the namespace to specialize a template.

Here's a link the gist: 

[https://gist.github.com/Hamondorf/66c37e519f8f7b55066dd6c10cc1b972](https://gist.github.com/Hamondorf/66c37e519f8f7b55066dd6c10cc1b972) 

I'm super open to feedback/criticism.
## [9][periodic-function: a small header only library to call a function at a specific time interval.](https://www.reddit.com/r/cpp/comments/iun0qh/periodicfunction_a_small_header_only_library_to/)
- url: https://www.reddit.com/r/cpp/comments/iun0qh/periodicfunction_a_small_header_only_library_to/
---
Hello all! 

I wanted to share a small C++ header only library I've been working on. I routinely run into situations where I need to repeatedly probe the state of hardware at a regular time interval and I wanted something I could use outside of other larger frameworks (i.e. `Qt` and `Boost.Signals`) and thus [periodic-function](https://github.com/DeveloperPaul123/periodic-function) was born!

I would be very grateful for any feedback/suggestions you have on anything at all (code organization, optimizations and so on).

Also, special thanks to u/TheLartians for taking the time to look at the code.
## [10][Where can I find Raspberry Pi projects written in C++?](https://www.reddit.com/r/cpp/comments/iuvcbb/where_can_i_find_raspberry_pi_projects_written_in/)
- url: https://www.reddit.com/r/cpp/comments/iuvcbb/where_can_i_find_raspberry_pi_projects_written_in/
---
Title. Does anyone know any site where can I find these projects?
## [11][Exception safety in the composite pattern](https://www.reddit.com/r/cpp/comments/iuu9ut/exception_safety_in_the_composite_pattern/)
- url: https://stackoverflow.com/questions/63945051/exception-safety-in-the-composite-pattern
---

