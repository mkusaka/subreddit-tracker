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
## [2][Branch Prediction - Fundamentals Every Programmer Need Not Know](https://www.reddit.com/r/cpp/comments/exm0vl/branch_prediction_fundamentals_every_programmer/)
- url: http://www.mycpu.org/branch-prediction-basics/
---

## [3][Looking for review: yet another RPC library](https://www.reddit.com/r/cpp/comments/exjf8i/looking_for_review_yet_another_rpc_library/)
- url: https://www.reddit.com/r/cpp/comments/exjf8i/looking_for_review_yet_another_rpc_library/
---
I would like to present [packio](https://github.com/qchateau/packio), an asynchronous library based on Boost.asio that lets you write msgpack-RPC clients and servers. 

You could say it's just another RPC library, and yes it is just another one, so let me explain what my goals were and what features you will find:

1. The library is async-oriented, implementing synchronous calls on top of it is simple but it's up to you.
2. The library heavily uses Boost.asio, meaning io\_context, sockets and connections are handled with asio. For thoses used to using asio, you don't have yet another API to learn, and the integration of packio inside of asio-based code feels more natural.
3. The server procedures can be either synchronous or asynchronous. This allows a server to process multiple RPC calls concurrently on a single thread. This can be useful if the server procedure requires IO.
4. The library requires boost &gt;=1.70 and C++17, and is available on conan for ease of use.

The library is currently in version 0.8.0 as I finished defining its identity and API not long ago. Although I did not release a 1.0.0 version yet, I believe the API and philosophy won't change too much from now on. I've been using the library for my own need for over a month now and I'm satisfied with the results. 

This is the reason I'm posting today: I would really appreciate any review or feedback as it is always difficult to have a proper view on its own work. I'd like this library to fit other people's need, not just my own.

The library is available under MPL-2, so feel free to use it !
## [4][AVX2 Vectorized Multithreaded Mandelbrot renderer](https://www.reddit.com/r/cpp/comments/exbc1c/avx2_vectorized_multithreaded_mandelbrot_renderer/)
- url: https://www.reddit.com/r/cpp/comments/exbc1c/avx2_vectorized_multithreaded_mandelbrot_renderer/
---
Hi all! I am an electrical engineering student and for fun have written a vectorized mandelbrot renderer. It uses AVX2 and multiple threads to compute 8 pixels in parallel per thread. Any advice on code and performance is welcome!

[https://github.com/voldemoriarty/Qbrot](https://github.com/voldemoriarty/Qbrot)

&amp;#x200B;

EDIT:

Thanks you guys/gals for the awesome tips! Following the points you gave allowed me to cut down the render time for 1080p 256 iterations from 41ms to 19ms (more than 50%). More importantly I learned some new stuff. Thanks for taking the time and reviewing the code. Stay blessed

&amp;#x200B;

Special Thanks:

u/TheDeviloper

u/xurxoham

u/anders987 

u/corysama
## [5][After std::ranges we need std::containers](https://www.reddit.com/r/cpp/comments/exd81l/after_stdranges_we_need_stdcontainers/)
- url: https://www.reddit.com/r/cpp/comments/exd81l/after_stdranges_we_need_stdcontainers/
---
I'm working on [a Unicode proposal based on strong types](https://github.com/Lyberta/cpp-unicode) which introduces new containers and views and I couldn't avoid the issue that current containers are specified in terms of `Cpp17NamedRequirements` and not Concepts which is Very Wrong™.

Now that we have redone iterators and algorithms using concepts and ranges, we need to do the same to containers.

In particular, in my rough translation of old requirements I got the following concepts that we need to have:

* `std::allocator` (oops, name taken)
* `std::default_insertable_into`
* `std::move_insertable_into`
* `std::copy_insertable_into`
* `std::emplace_constructible_into`
* `std::erasable_from`
* `std::container`
* `std::reversible_container`
* `std::allocator_aware_container`
* `std::sequence_container`
* `std::associative_container`
* `std::unordered_associative_container`

This is just rough mechanical translation so far and will require refinement. I thought the logical site would be to put new containers in `std2` namespace but then it hit me that they should be in `std::containers` namespace because then we can have `std::containers::vector` and `std::math::vector` (to a great pleasure of Guy Davidson and others).

I don't have enough experience with concepts and ranges yet so I can't write a paper at this time but I wanted to share my observations trying to write my own container in C++20.

EDIT: And to argue my point better, here's how the API of `std::containers::vector` may look like. This is the "diff view" which only shows changed API:

	template &lt;typename T, std::allocator_concept&lt;T&gt; A = std::allocator&lt;T&gt;&gt;
	class vector
	{
	public:
		vector() noexcept(noexcept(A{}))
		requires std::default_constructible&lt;A&gt;;
		
		explicit vector(size_type amount, const A&amp; allocator = A{})
		requires std::default_insertable_into&lt;T, vector&gt;;
		
		vector(size_type amount, const T&amp; value, const A&amp; allocator = A{})
		requires std::copy_insertable_into&lt;T, vector&gt;;
		
		template &lt;typename I, std::sentinel_for&lt;I&gt; S&gt;
		requires std::input_iterator&lt;I&gt; &amp;&amp;
			std::emplace_constructible_into&lt;T, vector, std::iter_reference_t&lt;I&gt;&gt;
		vector(I first, S last, const A&amp; allocator = A{});
		
		template &lt;typename R&gt;
		requires std::ranges::input_range&lt;R&gt; &amp;&amp; /*TODO*/
		Vector(R&amp;&amp; r, const A&amp; allocator = A{});
		
		vector(const vector&amp; other)
		requires std::copy_insertable_into&lt;T, vector&gt;;
		
		vector(const vector&amp; other, const A&amp; allocator)
		requires std::copy_insertable_into&lt;T, vector&gt;;
		
		vector(vector&amp;&amp; other, const A&amp; new_allocator)
		requires std::move_insertable_into&lt;T, vector&gt;;
		
		vector&amp; operator=(const vector&amp; other)
		requires std::copy_insertable_into&lt;T, vector&gt; &amp;&amp; std::copyable&lt;T&gt;;
		
		vector&amp; operator=(vector&amp;&amp; other) noexcept(
			std::allocator_traits&lt;A&gt;::
			propagate_on_container_move_assignment::value ||
			std::allocator_traits&lt;A&gt;::is_always_equal::value)
		requires (std::allocator_traits&lt;A&gt;::propagate_on_container_move_assignment::
			value == true) || (std::move_insertable_into&lt;T, vector&gt; &amp;&amp;
			std::movable&lt;T&gt;);
		
		template &lt;typename I, std::sentinel_for&lt;I&gt; S&gt;
		requires std::input_iterator&lt;I&gt; &amp;&amp;
			std::assignable_from&lt;T, std::iter_reference_t&lt;I&gt;&gt;
		void assign(I first, S last)
		requires std::emplace_constructible_into&lt;T, vector&gt; &amp;&amp;
			std::move_insertable_into&lt;T, vector&gt;;
		
		template &lt;typename R&gt;
		requires std::ranges::input_range&lt;R&gt; &amp;&amp; /*TODO*/
		void assign(R&amp;&amp; r);
		
		void assign(size_type amount, const T&amp; value)
		requires std::copy_insertable_into&lt;T, vector&gt; &amp;&amp; std::copyable&lt;T&gt;;
		
		void resize(size_type new_size)
		requires std::move_insertable_into&lt;T, vector&gt; &amp;&amp;
			std::default_insertable_into&lt;T, vector&gt;;
		
		void resize(size_type new_size, const T&amp; value)
		requires std::copy_insertable_into&lt;T, vector&gt;;
		
		void reserve(size_type new_capacity)
		requires std::move_insertable_into&lt;T, vector&gt;;
		
		void shrink_to_fit()
		requires std::move_insertable_into&lt;T, vector&gt;;
		
		template &lt;typename... Args&gt;
		reference emplace_back(Args&amp;&amp;... args)
		requires std::emplace_constructible_into&lt;T, vector, Args...&gt; &amp;&amp;
			std::move_insertable_into&lt;T, vector&gt;;
		
		void push_back(const T&amp; value)
		requires std::copy_insertable_into&lt;T, vector&gt;;
		
		void push_back(T&amp;&amp; value)
		requires std::move_insertable_into&lt;T, vector&gt;;
		
		template &lt;typename... Args&gt;
		iterator emplace(const_iterator position, Args&amp;&amp;... args)
		requires std::emplace_constructible_into&lt;T, vector, Args...&gt; &amp;&amp;
			std::move_insertable_into&lt;T, vector&gt; &amp;&amp; std::movable&lt;T&gt;;
		
		iterator insert(const_iterator position, const T&amp; value)
		requires std::copy_insertable_into&lt;T, vector&gt; &amp;&amp; std::copyable&lt;T&gt;;
		
		iterator insert(const_iterator position, T&amp;&amp; value)
		requires std::move_insertable_into&lt;T, vector&gt; &amp;&amp; std::movable&lt;T&gt;;
		
		iterator insert(const_iterator position, size_type amount, const T&amp; value)
		requires std::copy_insertable_into&lt;T, vector&gt; &amp;&amp; std::copyable&lt;T&gt;;
		
		template &lt;typename I, std::sentinel_for&lt;I&gt; S&gt;
		requires std::input_iterator&lt;I&gt; &amp;&amp; /*TODO*/
		iterator insert(const_iterator position, I first, S last)
		requires std::emplace_constructible_into&lt;T, vector&gt; &amp;&amp;
			std::move_insertable_into&lt;T, vector&gt; &amp;&amp; std::movable&lt;T&gt; &amp;&amp;
			std::swappable&lt;T&gt;;
		
		template &lt;typename R&gt;
		requires std::ranges::input_range&lt;R&gt; &amp;&amp; /*TODO*/
		iterator insert(const_iterator position, R&amp;&amp; r);
		
		iterator erase(const_iterator position)
		requires std::movable&lt;T&gt;;
		
		iterator erase(const_iterator first, const_iterator last)
		requires std::movable&lt;T&gt;;
	};

As you can see, old API required a lot of changes to be in the spirit of C++20.
## [6][A syntax-based overview of C++20 Concepts](https://www.reddit.com/r/cpp/comments/ex4lm5/a_syntaxbased_overview_of_c20_concepts/)
- url: https://omnigoat.github.io/2020/01/19/cpp20-concepts/
---

## [7][Runtime Compiled C++ Dear ImGui and DirectX11 Tutorial](https://www.reddit.com/r/cpp/comments/exoaip/runtime_compiled_c_dear_imgui_and_directx11/)
- url: https://www.enkisoftware.com/devlogpost-20200202-1-Runtime-Compiled-C++-Dear-ImGui-and-DirectX11-Tutorial
---

## [8][Why should I care about the New Cool Thing(TM)?](https://www.reddit.com/r/cpp/comments/exnfa1/why_should_i_care_about_the_new_cool_thingtm/)
- url: https://www.reddit.com/r/cpp/comments/exnfa1/why_should_i_care_about_the_new_cool_thingtm/
---
This may very well be a divisive topic, so I want to make it clear I'm not trolling or trying to start an endless argument.

Why should I care about C++20?

I'm a Windows dev (no choice) using MSVS.

For years now I've been forced to "upgrade" from VS2008/C++07 to VS2010/C++11, then VS2012/C++11 again, etc.. to VS2017 as of current.

Through all these updates I've not done anything new or different, but I have needed to contend with bugs either in the compiler or in the development environment itself.

The bottom line is I've updated for seemingly zero benefit, but only because someone higher up the chain thinks "latest is greatest".

It also appears to me that recent C++ updates have just been to add more functionality in some vague attempt to guess at what future software or features some program might require, and write a bunch of libraries to handle these things.

Why???

I haven't seen anything really useful appear in the language in some time, nor am I aware of any great improvements that affect existing code.

I haven't even seen any real  arguments over performance gains - merely the Next Cool Thing(TM) having some mystical power that I must use it in my code.

So... why should I "upgrade", and why should I care?
## [9][A freestyle rap from the British comedian Chris Turner dedicated to the ISO C++ standards committee meeting in Prague](https://www.reddit.com/r/cpp/comments/ewsgg4/a_freestyle_rap_from_the_british_comedian_chris/)
- url: https://www.cameo.com/v/TXgs6dWbN
---

## [10][A Universal I/O Abstraction for C++](https://www.reddit.com/r/cpp/comments/ewr18m/a_universal_io_abstraction_for_c/)
- url: https://cor3ntin.github.io/posts/iouring/
---

## [11][DS323I RTC HAL Driver Written in C++](https://www.reddit.com/r/cpp/comments/ex4kzu/ds323i_rtc_hal_driver_written_in_c/)
- url: https://www.reddit.com/r/cpp/comments/ex4kzu/ds323i_rtc_hal_driver_written_in_c/
---
Hey, guys! 

Here is a HAL driver for the DS3231 RTC I wrote a few months back. I felt to share.
It can work on any platform via the I2C interface. Similar to Rust's embedded-hal I2C interface which I really wish we had.
 
I'm really excited about C++ 20's concepts and I feel it'd be highly useful in this type of scenario.

Contributions are welcome!

https://github.com/lamarrr/DS3231
