# cpp
## [1][C++ Jobs - Q4 2020](https://www.reddit.com/r/cpp/comments/j3qems/c_jobs_q4_2020/)
- url: https://www.reddit.com/r/cpp/comments/j3qems/c_jobs_q4_2020/
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

* [C++ Jobs - Q3 2020](https://www.reddit.com/r/cpp/comments/hjnaf2/c_jobs_q3_2020/)
## [2][Herb Sutter Trip report: Autumn ISO C++ standards meeting](https://www.reddit.com/r/cpp/comments/jtt9cb/herb_sutter_trip_report_autumn_iso_c_standards/)
- url: https://herbsutter.com/2020/11/13/trip-report-autumn-iso-c-standards-meeting-virtual/
---

## [3][Stop learning C++ or keep going?](https://www.reddit.com/r/cpp/comments/ju1ikd/stop_learning_c_or_keep_going/)
- url: https://www.reddit.com/r/cpp/comments/ju1ikd/stop_learning_c_or_keep_going/
---
I picked C++ as my first language because I want to get into a video game company. It was difficult at first, but it gets easier over time, but still, hard nonetheless.

But then, something happened that force me to stay in my hometown until I die probably. 

So, the ways for me to earn money from programming now are: create my own software/app/game, freelancing, or remote job. 

Everybody wants to have a successful software that they can sell it for a stockpile of $$, but I'm going to go with the realistic view here and aim for the freelancing and remote job.

My question is, is C++ is a great language for me if I'm going to be a freelance and remote job? Or should I focus on another language? I'm still going to learn C++ though. I fell in love with it. I just going to assign more of my time to the new language and make C++ become my second language. 

Now, a little background of me. I'm 20 years old boy from a third world country living in a very very small city that I think about 50% still using a non-smartphone as their phone and there's absolutely no company that engaged in technology. Most people either work as a farmer, livestock breeder, and things like that. 

Don't get me wrong, I love my hometown. It is so peaceful here. I always said that my hometown is a perfect city for retirement. 

But, that's the thing. I'm not planning for retirement, I'm still trying to get some money.

Also, I will not attend college. Well, "can't" is the more proper word here.

So yeah, any advice would be appreciated. Thank you so much
## [4][With the new module system in C++20 will there be centralized repositories for modules?](https://www.reddit.com/r/cpp/comments/ju1h1e/with_the_new_module_system_in_c20_will_there_be/)
- url: https://www.reddit.com/r/cpp/comments/ju1h1e/with_the_new_module_system_in_c20_will_there_be/
---
Like nuget's, npm's or crates?

Examples:

* Rust Crates [https://crates.io/](https://crates.io/)
* Node NPM's [https://www.npmjs.com/](https://www.npmjs.com/) 
* .NET Nuget's [https://www.nuget.org/](https://www.nuget.org/)
* Java [https://mvnrepository.com/](https://mvnrepository.com/)
## [5][It Is Time for Compiler Developers to Step In](https://www.reddit.com/r/cpp/comments/jts223/it_is_time_for_compiler_developers_to_step_in/)
- url: https://www.reddit.com/r/cpp/comments/jts223/it_is_time_for_compiler_developers_to_step_in/
---
Warning: this is not a strictly-technical post. More like a rant.

TL;DR: actual implementors of C++ should be the ones responding to C++ developers' needs and pushing C++ forward according to them, not some committee authority. Things to be standardized should be battle-tested first.

Firstly, I want to express a great admiration and respect for men and women behind C++ implementations and tools as well as towards those who participate in various working groups and do the hard work of keeping the language both modern and standardized, the language I, among the plethora of others, use daily or from time to time, to make a living or to enjoy a little hobby or a pet project, big or small. I do not find myself in a position to make demands on them.

What was a tipping point for me to write this post was recent discussion ([this](https://www.reddit.com/r/cpp/comments/jtdk74/deprecating_volatile_jf_bastien_cppcon_2019/), [this](https://www.reddit.com/r/cpp/comments/jswz3z/compound_assignment_to_volatile_must_be/) and [this](https://www.reddit.com/r/cpp/comments/dk542b/cppcon_2019_jf_bastien_deprecating_volatile/)) about [deprecation](https://wg21.link/p1152r0) of some uses of `volatile` in C++20. Numerous embedded developers say that this strives to solve a language-lawyer-world problem of "ambiguous semantics" at the cost of real world usage and public image of C++. Meanwhile, C++ standard fails to address important real-world issues due to obvious time limitations and complexities of designing single universal solution to *hard problems*. The most prominent latest example is C++20 modules. They are standardized, but there are both a split between implementations with respect to strong/weak ownership model and [feeling](https://www.reddit.com/r/cpp/comments/akihlv/c_modules_might_be_deadonarrival/) that C++20 modules just failed to deliver. Oh, and do you remember `std::error_code` or `std::regex`?  


I think that the general mindset of "1. we need X in the Standard 2. only Standard-compliant code is OK" is wrong. Look at `__restrict` \- all major compilers support some form of it, yet it is non-standard. The fact that it is non-standard has not harmed anyone! Would anyone be harmed if major compiler's team explicitly says "we do not agree with Standard here and there and will not comply with it" about some controversial part of the C++ Standard? In fact, we do program real hardware, not C++ Abstract Machine. That's why `std::vector` \- like class can be implemented in C++ as understood by every compiler in the world despite all the [UB](https://wg21.link/p0593).

However, optimizing compilers can do all sorts of weird stuff and we want them to do this weird stuff to make our code fast. But we also want our code to be expressive. Would it matter that `std::start_lifetime_as` has not made it into C++20 Standard if every major compiler supported it? No. Would it allow us to write code which is more clear to both humans and optimizers? Yes.

Now is the time C++ must be up to the challenges because of increasing competition. Do not let authorities prevent you from using language extensions and non-standard features, we need them for  the language to evolve and find its way forward. It is better to standardize things when there is a clear path proven by a group of pioneers to be successful, than when we don't have *anything* but find ourselves in need of having *something*. It is better to have things in compilers than not have them just because of working group meetings schedule.
## [6][An industrial-grade RPC framework used throughout Baidu](https://www.reddit.com/r/cpp/comments/jtlebx/an_industrialgrade_rpc_framework_used_throughout/)
- url: https://github.com/apache/incubator-brpc
---

## [7][Deprecating volatile - JF Bastien - CppCon 2019](https://www.reddit.com/r/cpp/comments/jtdk74/deprecating_volatile_jf_bastien_cppcon_2019/)
- url: https://www.youtube.com/watch?v=KJW_DLaVXIY
---

## [8][Having a bit of trouble with inputs. not sure how I can put the two inputs together to create a line.](https://www.reddit.com/r/cpp/comments/ju0bl9/having_a_bit_of_trouble_with_inputs_not_sure_how/)
- url: https://www.reddit.com/r/cpp/comments/ju0bl9/having_a_bit_of_trouble_with_inputs_not_sure_how/
---
    #include &lt;iostream&gt;
    using namespace std;
    int main() {
         string name;
         int age;
         cout &lt;&lt; "Enter your name"&lt;&lt;endl;
         cout &lt;&lt; "Enter your age"&lt;&lt;endl;
         cin &gt;&gt; age;
    
         getline(cin,name);
         cout &lt;&lt; name &lt;&lt; " is " &lt;&lt; age &lt;&lt; endl;
        return 0;
    }
    
    This is the output 
    
    Enter your name
    Enter your age
     davaughn  23
     is 0
## [9][Tiny build tool : µmake (quick makefile alternative)](https://www.reddit.com/r/cpp/comments/jtlalr/tiny_build_tool_µmake_quick_makefile_alternative/)
- url: https://www.reddit.com/r/cpp/comments/jtlalr/tiny_build_tool_µmake_quick_makefile_alternative/
---
I used to make compile my projects from makefiles (had a template that was quiet flexible) but grew tired of the syntax and issues.

I can't get myself to like autoconf nor cmake (I quiet hate the second as a user, even if not as bad as bazel).

Lately because I have to code a lot in python, I realize this could be just right for the job. I heard of scons, it looks great but it feels an overkill for most of my projects.

So after using a small python script for months now in different projects I refactored it for easy git submodule usage and publish it.

µmake ([https://gitlab.com/corentin-pro/umake](https://gitlab.com/corentin-pro/umake)) is ~150 lines of python to compile easily and quickly. A caveat I forgot to write in the README is that it is mainly made for UNIX system.

I share this in case it would be inspirational for your own small build script or even better if you find any improvement/feedback to be made.

Also I put a BSD 3-clause in there, I am not sure what is the best to make it easy to use (creative-common? I want to avoid users to even have to credit it). My understanding is that if you add as submodule the license is also there so it can be used/modified/anything.
## [10][CppCast: Video Games, Robotics and Audio](https://www.reddit.com/r/cpp/comments/jtahbn/cppcast_video_games_robotics_and_audio/)
- url: https://cppcast.com/joel-lamotte/
---

## [11][Compound assignment to volatile must be un-deprecated](https://www.reddit.com/r/cpp/comments/jswz3z/compound_assignment_to_volatile_must_be/)
- url: https://www.reddit.com/r/cpp/comments/jswz3z/compound_assignment_to_volatile_must_be/
---
To my horror I discovered that C++20 has deprecated compound assignments to a volatile. For those who are at a loss what that might mean: a compound assignment is += and its family, and a volatile is generally used to prevent the compiler from optimizing away reads from and/or writes to an object.

In close-to-the-metal programming volatile is the main mechanism to access memory-mapped peripheral registers. The manufacturer of the chip provides a C header file that contains things like

    #define port_a (*((volatile uint32_t *)409990))
    #define port_b (*((volatile uint32_t *)409994))

This creates the ‘register’ port\_a: something that behaves very much like a global variable. It can be read from, written to, and it can be used in a compound assignment. A very common use-case is to set or clear one bit in such a register, using a compound or-assignment or and-assignment:

    port_a |= (0x01 &lt;&lt; 3 ); // set bit 3
    port_b &amp;= ~(0x01 &lt;&lt; 4 ); // clear bit 4

In these cases the compound assignment makes the code a bit shorter, more readable, and less error-prone than the alterative with separate bit operator and assignment.  When instead of port\_a a more complex expression is used, like uart\[ 2 \].flags\[ 3 \].tx, the advantage of the compound expression is much larger.

As said, manufacturers of chips provide C header files for their chips. C, because as far as they are concerned, their chips should be programmed in C (and with \*their\* C tool only). These header files provide the register definitions, and operations on these registers, often implemented as macros. For me as C++ user it is fortunate that I can use these C headers files in C++, otherwise I would have to create them myself, which I don’t look forward to.

So far so good for me, until C++20 deprecated compound assignments to volatile. I can still use the register definitions, but my code gets a bit uglier. If need be, I can live with that. It is my code, so I can change it. But when I want to use operations that are provided as macros, or when I copy some complex manipulation of registers that is provided as an example (in C, of course), I am screwed.

Strictly speaking I am not screwed immediately, after all deprecated features only produce a warning, but I want my code to be warning-free, and todays deprecation is tomorrows removal from the language.

I can sympathise with the argument that some uses of volatile were ill-defined, but that should not result in removal from the language of a tool that is essential for small-system close-to-the-metal programming. The get a feeling for this: using a heap is generally not acceptable. Would you consider this a valid argument to deprecate the heap from C++23?

As it is, C++ is not broadly accepted in this field. Unjustly, in my opinion, so I try to make my small efforts to change this. Don’t make my effort harder and alienate this field even more by deprecating established practice.

So please, un-deprecate compound assignments to volatile.
