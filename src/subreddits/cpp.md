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
## [2][Conan 2.0 announcement](https://www.reddit.com/r/cpp/comments/i13i7m/conan_20_announcement/)
- url: https://blog.conan.io/2020/07/28/Launching-Conan-2.0-Tribe.html
---

## [3][TIL: You can use initializer lists directly in a range-based for loop](https://www.reddit.com/r/cpp/comments/i15o5k/til_you_can_use_initializer_lists_directly_in_a/)
- url: https://www.reddit.com/r/cpp/comments/i15o5k/til_you_can_use_initializer_lists_directly_in_a/
---
Maybe everybody else already knew this but for me this was a revelation:

    for (int64_t entity_id: {2, 1460, 3525, 5177}) {
      std::cout &lt;&lt; entity_id &lt;&lt; std::endl;
    }
## [4][jwt-cpp: header only JWT library with generic JSON library support (pre-release)](https://www.reddit.com/r/cpp/comments/i0yvib/jwtcpp_header_only_jwt_library_with_generic_json/)
- url: https://www.reddit.com/r/cpp/comments/i0yvib/jwtcpp_header_only_jwt_library_with_generic_json/
---
# [JWT-CPP](https://github.com/Thalhammer/jwt-cpp/tree/v0.5.0-rc.0)

## New Release Candidate 0.5.0-rc.0

This new pre-release has a few key features:

- support for *any* JSON library
- support std::error_code workflow 

Beyond the additional features we have also been working hard to improve the quality so you (but mostly me) can rely on this code in mission critical applications. The doxygen API documentation is now hosted in GitHub Pages to improve readability. 

&gt; Hoping to get your feedback! Looking forward to your contributions

## [JSON Traits](https://github.com/Thalhammer/jwt-cpp/tree/v0.5.0-rc.0#providing-your-own-json-traits)
The challenge with a lot of the C++ JWT libraries (check out [jwt.io](https://jwt.io/) for other options) is they are fixed to a single JSON library. This forces many developers to rely on 2, 3 or more JSON library to build their application. It's a pain.

A JWT claim, is simply `jwt::basic_claim&lt;my_json_traits&gt;`, which allows the creation and decoding/verification to be generic to the underlying JSON library you are already using. Check out the [read me](https://github.com/Thalhammer/jwt-cpp/tree/v0.5.0-rc.0#providing-your-own-json-traits) for more information!

In order to maintain the API, there's still a `jwt::claim` which is the template specialization of `picojson_traits`. I am in the camp of changing this is a major release for the much more common `nlohmann/json` that everyone is familiar with.

&gt; What JSON libraries would you like to see supported?

## std::error_code overloads
Exception handling is expensive, for many reasons, and it's one of the reasons behind the addition of these new API overloads. You are no longer forced to do exception handling. If you application is based around error_codes, this offers you a very nice fit.

## What's next?
- Support more algorithms that are available and we are trying to expend support.
- Ensure the `json_traits` are sufficient to meet *most* JSON libraries 
- Testing, testing, testing.
## [5][Difference between 'new' and 'std::allocator&lt;T&gt;'](https://www.reddit.com/r/cpp/comments/i12cma/difference_between_new_and_stdallocatort/)
- url: https://www.reddit.com/r/cpp/comments/i12cma/difference_between_new_and_stdallocatort/
---
Hi Guys,

I understand that new  operator allocates+calls the default constructor. This basically initializes the memory.

Whereas std::allocator&lt;T&gt;::allocate allocates memory and std::allocator&lt;T&gt;::construct initializes memory.

Below is the code demonstrating my understanding.

    #include&lt;iostream&gt;
    #include&lt;memory&gt;
    
    int main()
    {  
      double *p1 = new double(1.0); //initialized memory = allocated+constructed                                                                          
      std::allocator&lt;double&gt; alloc;
      double *p2 = alloc.allocate(1); //uninitialized memory = allocated                                                                                  
      std::cout&lt;&lt;"Value in memory allocated by new is :"&lt;&lt;p1[0]&lt;&lt;"\n";
      std::cout&lt;&lt;"Value in memory allocated by std::allocator is :"&lt;&lt;*p2&lt;&lt;"\n";
      alloc.construct(p2,double(1.0)); //initialized memory = constructed                                                                                 
      std::cout&lt;&lt;"Value in memory initialized by std::construct is :"&lt;&lt;*p2&lt;&lt;"\n";
      delete p1;
      alloc.destroy(p2);
      alloc.deallocate(p2,1);
    }

When I run the above code, output is:-

    Value in memory allocated by new is :1
    Value in memory allocated by std::allocator is :-1.28823e-231
    Value in memory initialized by std::construct is :1

This confirms my understanding. But the question is where is the value add in splitting allocation &amp; initialization steps?  

My understanding is that new invokes only the default constructor, but in the example above I am able to call "new double(1.0)" with my own default value.  In this case what is the value addition of using std::allocator&lt;T&gt; class?
## [6][Inconsistent incomplete type problem](https://www.reddit.com/r/cpp/comments/i173a6/inconsistent_incomplete_type_problem/)
- url: https://www.reddit.com/r/cpp/comments/i173a6/inconsistent_incomplete_type_problem/
---
Hello all,

I'm trying to make a simple `Config` class for a project I am working on and have run into weird compatibility issue. Issue is that code compiles in MSVC and clang (using libc++), but not in GCC. I do understand why it is an issue and that technically GCC is correct and others imply more lenient.

Code is here: [https://godbolt.org/z/xEjcM1](https://godbolt.org/z/xEjcM1)

What I am trying to figure out how to make this work without loosing the semantics of it. One way I figures is encasing value in `std::unique_ptr`, but that introduces extra allocation and indirection.

Is there a way to refactor it so it is:
- c++ standard compliant
- not loose semantics
- not require extra indirection?

Weirdly enough if I use `std::map` then GCC is happy.

If anyone interested my full implementation is here: https://github.com/albeva/fbide/blob/master/src/Config/Config.hpp
## [7][Conan 1.28 released](https://www.reddit.com/r/cpp/comments/i1709w/conan_128_released/)
- url: https://github.com/conan-io/conan/releases/tag/1.28.0
---

## [8][CppCast: Visual Effects](https://www.reddit.com/r/cpp/comments/i0zbpa/cppcast_visual_effects/)
- url: https://cppcast.com/josh-filstrup-vfx/
---

## [9][On the search for ideas](https://www.reddit.com/r/cpp/comments/i16kqj/on_the_search_for_ideas/)
- url: https://www.reddit.com/r/cpp/comments/i16kqj/on_the_search_for_ideas/
---
Hi guys!
I'm about to start finding a job as a c++ dev. I'm 15 and I have about yearly experience in STL, console apps and algorithms. I'm into maths and I'm good at it. Becouse every project I've done so far was either implementations of algorithms, data structures or solving exercises I'm looking for ideas for a projects I can put in my portfolio, I'm starting to study WxWidgets and if needed I'll learn some graphics engine. If you have any interesting ideas for one-person project I would love to see them in the comments.
Thanks!
## [10][libstud-json: JSON pull-parser/push-serializer library for C++](https://www.reddit.com/r/cpp/comments/i0l6ao/libstudjson_json_pullparserpushserializer_library/)
- url: https://github.com/libstud/libstud-json/
---

## [11][C++ has become easier to write than Java](https://www.reddit.com/r/cpp/comments/i15g4p/c_has_become_easier_to_write_than_java/)
- url: https://www.reddit.com/r/cpp/comments/i15g4p/c_has_become_easier_to_write_than_java/
---
After a while I started to do some Android development with Java (not into Kotlin yet). It's nice to see that after getting used to C++11 features like auto etc, Java now feels VERY cumbersome.

I hate this pattern in Java:

MyLongNamedClass foo = new MyLongNamedClass();

In C++ I can, e.g.:

auto foo = new MyLongNamedClass;

Or just:

MyLongNamedClass foo;

This was only a single example. Java is lagging behind C++.
