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
## [2][libstdc++ &lt;algorithm&gt; header transitively including &lt;range&gt;](https://www.reddit.com/r/cpp/comments/ibz0m2/libstdc_algorithm_header_transitively_including/)
- url: https://www.reddit.com/r/cpp/comments/ibz0m2/libstdc_algorithm_header_transitively_including/
---
I noticed some time ago that in c++20 mode the &lt;algorithm&gt; header in libstdc++ is substantially bigger than in c++11 mode :

`echo "#include &lt;algorithm&gt;" | gcc -std=c++11 -P -E -x c++ - | wc -l`

evaluates to 11760 loc, while

`echo "#include &lt;algorithm&gt;" | gcc -std=c++20 -P -E -x c++ - | wc -l`

evaluates to 45219 loc. Clicking through the header apparently in c++20 we have the include chain

&lt;algorithm&gt; -&gt; &lt;bits/ranges\_algo.h&gt; -&gt; &lt;bits/ranges\_algobase.h&gt; -&gt; &lt;range&gt;.

How could that happen? Measuring the compile time of file only including the &lt;algorithm&gt; header I get with the c++11 switch 120 ms and with the c++20 switch **600 ms** (empty file is 100 ms). I a bit baffled that while standardizing the &lt;range&gt; header this either slipped through the cracks or it was actively decided that this is ok. The &lt;algorithm&gt; header is virtually anywhere and with c++20 flag enabled every file including the &lt;algorithm&gt; header will compile half a second slower...
## [3][I made this ASCII Ray Tracing program in C++, any feedback and suggestions are welcome](https://www.reddit.com/r/cpp/comments/ibcvks/i_made_this_ascii_ray_tracing_program_in_c_any/)
- url: https://youtu.be/QkETiyYWh2o
---

## [4][unique_ptr, shared_ptr, weak_ptr, or reference_wrapper for class relationships](https://www.reddit.com/r/cpp/comments/ibzqvj/unique_ptr_shared_ptr_weak_ptr_or_reference/)
- url: https://www.nextptr.com/tutorial/ta1450413058/unique_ptr-shared_ptr-weak_ptr-or-reference_wrapper-for-class-relationships
---

## [5][zip iterator implementation feedback](https://www.reddit.com/r/cpp/comments/ibxxs3/zip_iterator_implementation_feedback/)
- url: https://www.reddit.com/r/cpp/comments/ibxxs3/zip_iterator_implementation_feedback/
---
I wrote a zip iterator as a challenge of my template metaprogramming knowledge and to make working with the standard parallel algorithms easier for numerical simulations: [https://github.com/mfdeakin/zip](https://github.com/mfdeakin/zip)

Goals of this project are ease of use (ideally as easy as python's zip), high (runtime) performance (though recommendations for improving compile times would be appreciated), and interoperability with the standard library algorithms. I think I've largely achieved these goals, though GPU performance is still unknown and something that concerns me due to my use of functors.

One somewhat hidden issue I'd like feedback/recommendations on is that I had to implement std::swap(tuple&lt;\_Elements...&gt; &amp;&amp;, tuple&lt;\_Elements...&gt; &amp;&amp;) to support working with std::sort and potentially other algorithms; ideally I'd be able to specialize iter\_swap instead, but std::sort in gcc's implementation makes the mistake (IMO) of using the qualified std::iter\_swap rather than just iter\_swap. (also, why doesn't swap doesn't use universal references? or support rvalues in general?)

I avoided exceptions with this code; this is largely for supporting GPUs.
## [6][vcpkg.info](https://www.reddit.com/r/cpp/comments/ibjgqq/vcpkginfo/)
- url: https://www.reddit.com/r/cpp/comments/ibjgqq/vcpkginfo/
---
[https://vcpkg.info](https://vcpkg.info) is a simple page where you can browse all libraries from vcpkg repository.

You can search for library using full text search on names or descriptions, or you can select libs using tags.

Library information page contains a list of all versions of particular library with appropriate commit hash, which you need to checkout to specific revision, if you want to use that version.

Until we get [promised proper versioning](https://devblogs.microsoft.com/cppblog/vcpkg-2020-04-update-and-product-roadmap/) in vcpkg, this is the easies way to find out about all available versions.

Please bear in mind, that this is still very early version, and may contain some errors here and there.

Happy browsing!
## [7][structopt: Parse command line arguments by defining a struct](https://www.reddit.com/r/cpp/comments/ibdp1p/structopt_parse_command_line_arguments_by/)
- url: https://github.com/p-ranav/structopt
---

## [8][How expensive is a shared_lock?](https://www.reddit.com/r/cpp/comments/ibsc8f/how_expensive_is_a_shared_lock/)
- url: https://www.reddit.com/r/cpp/comments/ibsc8f/how_expensive_is_a_shared_lock/
---
If I have an unordered_set (or a pointer to an unordered_set), and gate the writes with a unique lock on a shared mutex:

    std::unique_lock&lt;std::shared_mutex&gt; ul(mutex);
    std::swap(set_, newSet);

and the reads with a shared_lock

    std::shared_lock&lt;std::shared_mutex&gt; sl(mutex);
    set_.find(_) != set_.end();

If I'm doing thousands of membership checks (on a single thread) on the underlying set, is it beneficial to group these together under a single shared_lock or can I use these more freely per check? What's the overhead here?

Edit: A lot of people are giving me more general mutex and lock advice which I appreciate but it's stuff I already know. I'm more curious about the underlying cost in a shared_lock implementation if anyone is familiar with that. More particularly how expensive acquiring a shared_lock is because technically there is no contention with other shared_lock's. (Assume very infrequent writes)

And yes, I will profile both options. Was looking for some deeper insights into shared_lock implementation details.
## [9][c++ svg library](https://www.reddit.com/r/cpp/comments/ibf3zh/c_svg_library/)
- url: https://www.reddit.com/r/cpp/comments/ibf3zh/c_svg_library/
---
I wrote a standalone c++ library to create, animate, manipulate and render SVG files.

[https://github.com/sammycage/lunasvg](https://github.com/sammycage/lunasvg)

    #include &lt;lunasvg/svgdocument.h&gt;
    
    using namespace lunasvg;
    
    int main()
    {
        SVGDocument document;
        document.loadFromFile("tiger.svg");
        
        Bitmap bitmap = document.renderToBitmap();
        
        // do something useful with the bitmap.
        
        return 0;
    }
    

Any contributions and questions are welcome.
## [10][Eliminating the static overhead of C++ ranges (2019)](https://www.reddit.com/r/cpp/comments/ib7tt9/eliminating_the_static_overhead_of_c_ranges_2019/)
- url: https://vector-of-bool.github.io/2019/10/21/rngs-static-ovr.html
---

## [11][Mastering C++: Books | Courses | Tools | Tutorials | Blogs | Communities](https://www.reddit.com/r/cpp/comments/ibenc3/mastering_c_books_courses_tools_tutorials_blogs/)
- url: http://www.vishalchovatiya.com/mastering-c-books-courses-tools-tutorials-blogs-communities/
---

