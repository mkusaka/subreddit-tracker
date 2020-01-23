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
## [2][How to introduce exceptions into historically exception free code base?](https://www.reddit.com/r/cpp/comments/espg2t/how_to_introduce_exceptions_into_historically/)
- url: https://www.reddit.com/r/cpp/comments/espg2t/how_to_introduce_exceptions_into_historically/
---
We have a massive C++ code base where much of the core was written in the early 2000's using error codes and a custom template library to support an exception free environment.

Over the years we've lost support in the team to continue to maintain and debug the custom template library and there's been a desire to "modernize" our code.  One aspect being the use of the STL.  However, this creates some issues as we can now have functions that can return error codes and throw exceptions which makes things increasingly difficult to reason about.

Has anyone has experience going from exception free code to introducing something like the STL?  If so, what was your approach?

A couple things I could think of off the top of my head.  Have one master catch statement in our runtime code.  Or perhaps create macro that wraps every STL call that can throw in a try catch (though this feels cumbersome).  I'll admit, I don't have much experience with exceptions as I've mainly worked in this codebase for the past couple years.
## [3][Standard C++ library on GCC's AVR target](https://www.reddit.com/r/cpp/comments/esk2wy/standard_c_library_on_gccs_avr_target/)
- url: https://www.reddit.com/r/cpp/comments/esk2wy/standard_c_library_on_gccs_avr_target/
---
I was doing several projects on AVR micro controllers with C++ and my custom build of GCC 9.2, just to get modern C++ core language features available (C++17). The problem arise when I try to get standard library support. I do not need containers (maybe except std::array) but things like type\_traits or tuple would be really appreciated. I came up with small subset of things that I need, mainly by "stealing" small snippets from cppreference or creating things by myself: [https://github.com/lukaszgemborowski/avrcpp](https://github.com/lukaszgemborowski/avrcpp) 

The approach is cool to get know some C++ internals but is really annoying in terms of development and being C++ standard compliant. So my question is: do you have any ideas how to approach that problem? Maybe subset of standard C++ library is available on AVR but I just can't get it to work? Anyone have a similar experience?

PS. hello community, this is my first post here! ;-)
## [4][Concepts merged to clang trunk](https://www.reddit.com/r/cpp/comments/esamj3/concepts_merged_to_clang_trunk/)
- url: https://www.reddit.com/r/cpp/comments/esamj3/concepts_merged_to_clang_trunk/
---
[https://twitter.com/saarraz1/status/1219809819781222401?s=20](https://twitter.com/saarraz1/status/1219809819781222401?s=20)

I just saw this and got excited - I hope it’s legit as I think it’s from the person actually writing the code for Concepts in Clang.
## [5][SObjectizer-5.7 with support of send_case in select() and so5extra-1.4 relicensed under BSD-license](https://www.reddit.com/r/cpp/comments/esrn1x/sobjectizer57_with_support_of_send_case_in_select/)
- url: https://www.reddit.com/r/cpp/comments/esrn1x/sobjectizer57_with_support_of_send_case_in_select/
---
[SObjectizer](https://github.com/Stiffstream/sobjectizer) is one of the few alive and evolving "actor frameworks" for C++ (others are [QP/C++](https://www.state-machine.com/qpcpp/), [CAF: C++ Actor Framework](http://actor-framework.org/) and very young project [rotor](https://github.com/basiliscos/cpp-rotor)). A brief introduction for SObjectizer can be found in [this presentation](https://sourceforge.net/projects/sobjectizer/files/sobjectizer/Slides/What%20is%20SObjectizer-5.7%20%28at%20v.5.7.0%29.pdf/download) or in this [overview](https://habr.com/en/post/458202/). It is worth mentioning that SObjectizer supports not the only Actor, but also Publish-Subscribe and Communicating Sequential Processes models.

New versions of SObjectizer and its companion project [so5extra](https://github.com/Stiffstream/so5extra) are released.

TL;DR:

* SObjectizer-5.7 now supports the `send_case` in `select()` function. It makes SObjectizer's `select()` much more similar to Golang's `select` statement. But it breaks the compatibility with v.5.6 because old `case_` function is now renamed to `receive_case`;
* v.5.7 of SObjectizer fixes a flaw in enveloped message delivery mechanism with respect to `transfer_to_state()` and `suppress()` functionality;
* so5extra is now relicensed and distributed under BSD-3-CLAUSE license. Previous versions of so5extra were distributed under GNU Affero GPL v.3 and commercial licenses;
* so5extra-1.4 now implements fixed-capacity message chains that capacity is known at the compile-time.

Speaking more lengthy the main new feature of SObjectizer-5 is the support of `send_case` in `select()` function (somewhat similar to Golang's `select` statement with sending an outgoing message in a channel). Now it is possible to do things like that:

```cpp
using namespace so_5;

struct Greetings {
   std::string msg_;
};

// Try to send messages to the corresponding chains,
// but wait no more than 250ms for all the operations.
select(from_all().handle_n(3).total_time(250ms),
   send_case(chAlice,
      message_holder_t&lt;Greetings&gt;::make("Hello, Alice!"),
      []{ std::cout &lt;&lt; "message sent to chAlice" &lt;&lt; std::endl; }),
   send_case(chBob,
      message_holder_t&lt;Greetings&gt;::make("Hello, Bob!"),
      []{ std::cout &lt;&lt; "message sent to chBob" &lt;&lt; std::endl; }),
   send_case(chEve,
      message_holder_t&lt;Greeting&gt;::make("Hello, Eve!"),
      []{ std::cout &lt;&lt; "message sent to chEve" &lt;&lt; std::endl; }));
```

`send_case()` can be used with `receive_case()` in the same `select()`. For example, this is SObjectizer's version of calculation of Fibonacci numbers in separate thread from [Golang's tour](https://tour.golang.org/concurrency/5):

```cpp
using namespace std;
using namespace std::chrono_literals;
using namespace so_5;

struct quit {};

void fibonacci( mchain_t values_ch, mchain_t quit_ch )
{
   int x = 0, y = 1;
   mchain_select_result_t r;
   do
   {
      r = select(
         from_all().handle_n(1),
         // Sends a new message of type 'int' with value 'x' inside
         // when values_ch is ready for a new outgoing message.
         send_case( values_ch, message_holder_t&lt;int&gt;::make(x),
               [&amp;x, &amp;y] { // This block of code will be called after the send().
                  auto old_x = x;
                  x = y; y = old_x + y;
               } ),
         // Receive a 'quit' message from quit_ch if it is here.
         receive_case( quit_ch, [](quit){} ) );
   }
   // Continue the loop while we send something and receive nothing.
   while( r.was_sent() &amp;&amp; !r.was_handled() );
}

int main()
{
   wrapped_env_t sobj;

   thread fibonacci_thr;
   auto thr_joiner = auto_join( fibonacci_thr );

   // The chain for the Fibonacci numbers will have limited capacity.
   auto values_ch = create_mchain( sobj, 1s, 1,
         mchain_props::memory_usage_t::preallocated,
         mchain_props::overflow_reaction_t::abort_app );

   auto quit_ch = create_mchain( sobj );
   auto ch_closer = auto_close_drop_content( values_ch, quit_ch );

   fibonacci_thr = thread{ fibonacci, values_ch, quit_ch };

   // Read the first 10 numbers from values_ch.
   receive( from( values_ch ).handle_n( 10 ),
         // And show every number to the standard output.
         []( int v ) { cout &lt;&lt; v &lt;&lt; endl; } );

   send&lt; quit &gt;( quit_ch );
}
```

The full release notes for SObjectizer-5.7.0 can be found [here](https://github.com/Stiffstream/sobjectizer/wiki/v.5.7.0).

The main new feature of so5extra library is the new license. Since v.1.4.0 so5extra is distributed under BSD-3-CLAUSE license and all so5extra's stuff (like Asio's based dispatcher and environment infrastructures, different kinds of message boxes and so on) can now be used for free. Even in closed-source commercial projects.

The only one new addition to so5extra is the implementation of message chain that maximal capacity is known at the compile time. Usage of such message chains can be useful, for example, in request-reply scenarios where just only reply message is expected:

```cpp
#include &lt;so_5_extra/mchains/fixed_size.hpp&gt;
#include &lt;so_5/all.hpp&gt;
...
using namespace so_5;

// The chain for the reply message.
auto reply_ch = extra::mchains::fixed_size::create_mchain&lt;1&gt;(env,
   mchain_props::overflow_reaction_t::drop_newset);
// Send a request.
send&lt;SomeRequest&gt;(target, ..., reply_ch, ...);
// Wait and handle the response.
receive(so_5::from(reply_ch).handle_n(1), [](const SomeReply &amp; reply) { ... });
```

The full release notes for so5extra can be found [here](https://github.com/Stiffstream/so5extra/wiki/v.1.4.0).

I hope that information about SObjectizer/so5extra will be useful for someone and will be glad to answer questions if any.
## [6][How to run build cpp project using gradle in docker container?](https://www.reddit.com/r/cpp/comments/esrk8j/how_to_run_build_cpp_project_using_gradle_in/)
- url: https://www.reddit.com/r/cpp/comments/esrk8j/how_to_run_build_cpp_project_using_gradle_in/
---
Hi everyone, 

I am trying to build my cpp project via gradle on docker. Without docker, it works fine, but i need to run in docker. Untill now i haven't seen any proper example or helpful documentation at all. 

My objective is so,

`$ gradle build`

then build image and run container then compile my project there.
## [7][Fantastic Bugs and Where to Find Them](https://www.reddit.com/r/cpp/comments/es9urc/fantastic_bugs_and_where_to_find_them/)
- url: https://medium.com/tanker-blog/fantastic-bugs-and-where-to-find-them-3c95afe06357
---

## [8][C++ UPnP client library using Boost.Asio](https://www.reddit.com/r/cpp/comments/eseaw4/c_upnp_client_library_using_boostasio/)
- url: https://www.reddit.com/r/cpp/comments/eseaw4/c_upnp_client_library_using_boostasio/
---
As the title says, [cpp-upnp](https://github.com/equalitie/cpp-upnp) is a UPnP library written in C++ using Boost.Asio. UPnP is a big set of protocols and this library currently only supports creating, removing and listing of IPv4 TCP and UDP port mappings.

The API is based around Asio coroutines, which suffices for our purposes ATM, but if there is interest I'm happy to add support for other idioms using Asio's async result machinery.
## [9][How about a much simpler C++ abbreviated lambda expression syntax?](https://www.reddit.com/r/cpp/comments/esevzz/how_about_a_much_simpler_c_abbreviated_lambda/)
- url: https://www.reddit.com/r/cpp/comments/esevzz/how_about_a_much_simpler_c_abbreviated_lambda/
---
A week ago, this article about why the abbreviated lambda syntax was rejected was circulating: https://brevzin.github.io/c++/2020/01/15/abbrev-lambdas/ ([discussion on reddit](https://old.reddit.com/r/cpp/comments/epq4ui/why_were_abbrev_lambdas_rejected/)).

C++ still desperately needs a shorter lambda syntax for the common case where the lambda is a simple predicate, or is used as an alternative to std::bind, or any other situation where only one expression is necessary. However:

* The abbreviated syntax doesn't have to be (and shouldn't be) semantically different from long lambdas. `[](int const *p) =&gt; *p` and `[](int const *p) { return *p; }` should be identical. If `decltype(auto)` semantics are needed, one can just write `[](int const *p) -&gt; decltype(auto) { return *p; }`.
* I think it's unnecessary to play "syntax golf"; to make the syntax as short as possible. Requiring `[](auto x) =&gt; x` instead of `[](x) =&gt; x` is an entirely acceptable compromise (and, I would argue, more in line with the rest of C++, so it's easier to learn).

I don't think abbreviated lambdas have to be fancy. They should act as a way to write shorter, more readable code. They don't have to enable the tersest code possible, and they don't have to change the defaults of lambdas. Just compare these three uses of `std::any_of`, where the first uses C++11 lambdas, the second uses a very basic terse lambda syntax, and the third uses P0573R2's terse lambdas:

    std::any_of(arr.begin(), arr.end(),
        [](auto &amp;x) { return x.somePred(); });

    std::any_of(arr.begin(), arr.end(),
        [](auto &amp;x) =&gt; x.somePred());

    std::any_of(arr.begin(), arr.end(),
        [](x) =&gt; x.somePred());

To my eyes, both the second and third call to `any_of` have a clear readability advantage over the first. However, the first and second call to `any_of` have an obvious meaning to someone who knows the rest of C++; the third call is less obvious (for example, is x a reference or a value?).

There may be things I'm missing, maybe the C++ lambdas are way more subtly broken than I think because their return type defaults to `auto` and not `decltype(auto)`, or maybe there are parsing challenges even with the simpler abbreviated lambdas I'm proposing. However, every time I write `someFunction([](auto &amp;x) { return x.someMethod(); });`, I feel the pain of the loss of readability.

I'm sure these thoughts aren't new, what I'm proposing is essentially the most obvious possible way to abbreviate the current lambda syntax, but I really think this option should be seriously considered.
## [10][What's wrong wity my project that using gtest?](https://www.reddit.com/r/cpp/comments/esqpjm/whats_wrong_wity_my_project_that_using_gtest/)
- url: https://www.reddit.com/r/cpp/comments/esqpjm/whats_wrong_wity_my_project_that_using_gtest/
---
Hi. i'm noob to googletest. I setup a test project:[https://github.com/LinArcX/tddcpp](https://github.com/LinArcX/tddcpp)

Go to test directory and create build directory in it and then:

    cmake ..
    make

This gives me this error:

    /bin/ld: CMakeFiles/tddcpp_test.dir/fizzbuzz/fizzbuzz_test.cpp.o: in function `FizzBuzzTest_canAssertTrue_Test::TestBody()':
    fizzbuzz_test.cpp:(.text+0x23): undefined reference to `Fizzbuzz::Fizzbuzz()'
    /bin/ld: fizzbuzz_test.cpp:(.text+0x2f): undefined reference to `Fizzbuzz::~Fizzbuzz()'
    collect2: error: ld returned 1 exit status
    make[2]: *** [CMakeFiles/tddcpp_test.dir/build.make:102: tddcpp_test] Error 1
    make[2]: Leaving directory '/mnt/D/WorkSpace/C++/cpp/projects/active/tddcpp/test/build'
    make[1]: *** [CMakeFiles/Makefile2:79: CMakeFiles/tddcpp_test.dir/all] Error 2
    make[1]: Leaving directory '/mnt/D/WorkSpace/C++/cpp/projects/active/tddcpp/test/build'
    make: *** [Makefile:87: all] Error 2
## [11][Why modules instead of definable compilation units?](https://www.reddit.com/r/cpp/comments/esn7kk/why_modules_instead_of_definable_compilation_units/)
- url: https://www.reddit.com/r/cpp/comments/esn7kk/why_modules_instead_of_definable_compilation_units/
---
C++20 brings modules into cpp, if I understand correctly, it's meant for better compilation time and better segregate interfaces.

I was looking into unity build for some reasons and thought it may be a good idea to segregate a project in terms of unity build units.

Each small unit that does not change much can use the same compiled output as if it was all shared library. This may also inline a few functions (and the compiler will know more about the whole thing, perform optimizations and checks like it would do normally...).

This should be able to map to the problem of modules (though identifier lifetime should be managed as if it was static in the c sense?),  and really easy to implement using a few preprocessors. The order of include in unity build can be defined using numbers in the syntax, etc.

Am I missing something?  It seems that this is much easier to implement, and can move to complete modules eventually.

EDIT:

I did thought about the definition problems,  and it seems it will need a new keyword to restrict lifetime and use preprocessor to do the job (e.g. a static in the C sense when there's no branches for module private variables.) 

Compiler can be like a library in the C sense, but the compiler can do all the integrity checks and merge the dependencies and remove unexported  leaf exports.

Libraries of course is already doing what modules should be doing, but they typically need to follow a guideline of defining opaque types and interface in public headers, and rely on the linker to do the job.

The idea, if I understand correctly, is perhaps making this easier and more cost-effective within projects by offering a better interface, and making the compiler more responsibile in checking integrity by providing guarantees.
