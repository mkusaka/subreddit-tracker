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
## [2][PhysFS performance, a story of threading and locking · Mathieu Ropert](https://www.reddit.com/r/cpp/comments/hyoq8s/physfs_performance_a_story_of_threading_and/)
- url: https://mropert.github.io/2020/07/26/threading_with_physfs/
---

## [3][C++ on Sea video - "How to Refactor Millions of Line of Code Without Alienating your Colleagues" - Fred Tingaud](https://www.reddit.com/r/cpp/comments/hyqyt7/c_on_sea_video_how_to_refactor_millions_of_line/)
- url: https://www.youtube.com/watch?v=Fxe7JoGSyy8
---

## [4][Why can't we have void type variables?](https://www.reddit.com/r/cpp/comments/hyn2nf/why_cant_we_have_void_type_variables/)
- url: https://www.reddit.com/r/cpp/comments/hyn2nf/why_cant_we_have_void_type_variables/
---
Hello.

Sometime ago I was writing some code, which basically looked like this:

    T result = func();
    post_func_action();
    return result;

The problem is, my code doesn't compile when type `T` is `void`. Therefore I had to write a second template for handling cases where `func` returns `void`. In my opinion this is unnecessary, and reduces code reusability. I am aware that allowing `void` type variables would cause some inconvenience with `void*` pointers (which was, if you ask me, a very unfortunate choice for a name), but other than that, is there any particular reason why values of `void` type can't exist?
## [5][C++ on Sea 2020 video - "Combining Modern C++ and Lua" - James Pascoe](https://www.reddit.com/r/cpp/comments/hyei46/c_on_sea_2020_video_combining_modern_c_and_lua/)
- url: https://www.youtube.com/watch?v=QwfL72yHZEY
---

## [6][Differences between old (pre GCC6) and new-style function multi-versioning and how to migrate](https://www.reddit.com/r/cpp/comments/hyfhy3/differences_between_old_pre_gcc6_and_newstyle/)
- url: https://extensa.tech/blog/migrating-multiversioning/
---

## [7][Is reentrancy a problem because of recursive calls? When should you worry about code reentrancy?](https://www.reddit.com/r/cpp/comments/hydpyx/is_reentrancy_a_problem_because_of_recursive/)
- url: https://www.reddit.com/r/cpp/comments/hydpyx/is_reentrancy_a_problem_because_of_recursive/
---
I read [this](https://deadbeef.me/2017/09/reentrant-threadsafe) blog post (which seems to be a parrot of [Wiki: Reentrancy (computing)](https://en.wikipedia.org/wiki/Reentrancy_(computing))) about reentrancy and it appears to be a problem because of a recursive call (in their case, indirectly via `my_func()`).

I guess what I'm really asking is under what conditions should you worry about code reentrancy. The examples and wiki mentions two:

1. Recursive calls
2. An interrupt service routine which, separate from `main()` execution, executes a function that had already been called from `main()`.

If you're writing a multithreaded program and **Thread1** calls *foo()* and **Thread2** also calls *foo(),* would you need to take care to make *foo()* reentrant?

Is that it?
## [8][Is it sane to do iOS and Android dev in cpp these days?](https://www.reddit.com/r/cpp/comments/hyk0do/is_it_sane_to_do_ios_and_android_dev_in_cpp_these/)
- url: https://www.reddit.com/r/cpp/comments/hyk0do/is_it_sane_to_do_ios_and_android_dev_in_cpp_these/
---
Hi, I'm looking for subjective opinions and real battle stories here, I know it's \*possible\*. I know QT and JUCE and others say they do it, but I've also read lots of folks saying it's actually a fools game....

I'm making some music pedagogy applications. I prototype them on desktop in Max/MSP with the app logic in Scheme, using a Max external I wrote in C to embed S7 Scheme in Max. I love doing the domain logic in Scheme, and prototyping in Max is super productive. But I eventually want to release these as small apps for desktop on windows and OSX, and (if possible) for iOS and Android. With S7, I can plug my domain layer into any C framework, it's got a \*great\* FFI. My apps need to do things like midi i/o, make sounds, run a scheduler (ie schedule making sounds into the future), and draw custom widgets like keyboards and musical notation.  But nothing heavy duty in the audio or graphics realm, this is for practising, not playing live or anything performance critical. The hard part is in the scheme, the C/C++ part is plumbing really. 

I'm trying to get a sense of whether investing time on something like QT or JUCE for their supposed "deploy anywhere" is worthwhile, or if cross-platform mobile dev in C++ is bologna. And if you say these are not good solutions, can you tell me if there is a good way to do mobile in C++ for those platforms? I might be able to live with having to use a different toolkit for each platform, but these are apps for a very small market and it's likely to be just me or me plus 1 person. The only thing I'm married to is Scheme, which thus means C at some level for hosting the scheme interpreter. 

Thanks!
## [9][Code coverage problem](https://www.reddit.com/r/cpp/comments/hyp2b2/code_coverage_problem/)
- url: https://www.reddit.com/r/cpp/comments/hyp2b2/code_coverage_problem/
---
Hi all. I've been programming in C++ for a while now, but I've always had colleagues who work on the more devopsy side of things. To fill in my gaps in knowledge, I've decided to make my own project starter template (I know there are some great ones out there, but I figured starting from scratch would give me a better understanding of how things work). Long story short, there's one thing which I can't quite figure out how to do: accurate code coverage of functions which are not invoked in tests. I've included a minimum working example [here](https://github.com/kubagalecki/my-cpp-starter). When a function isn't called, I'd like gcov(r) to report 0% coverage of it. Instead, the compiler seems to be optimizing it away (despite -O0 and all the anti-inlining flags) which results in the tool reporting file coverage as 100%. When I split up the declaration and definition into separate files, the file with the definition is omitted in the coverage report altogether. There are some posts on various forums about this problem, but the answers there don't seem to do the trick for me, so I thought I'd take up some of the time of the fine folks on this sub. Any help would be highly appreciated.

Edit for posterity, see below for details: The problem seems to only occur with inlined functions. A possible solution is to mark anything that may be inlined with `__attribute__ ((used))`.  

## [10][C++ Compiler Optimization at work](https://www.reddit.com/r/cpp/comments/hymfod/c_compiler_optimization_at_work/)
- url: https://www.reddit.com/r/cpp/comments/hymfod/c_compiler_optimization_at_work/
---
I am currently in chapter 18 of PPP using C++.  Found the below observation about C++ compiler optimization interesting. Posting it for the benefit of fellow C++ beginners.

So the discussion was about how move constructor is better than copy constructor in the below scenario. 

1) We have an object factory function creating objects of large size and returning them.

2) 2nd function is utilizing the above object factory to create objects.

So the question is should the object factory function return references/values of the created objects?

Object factory returning object reference:-

    object* object_factory()
    {
        object *o = new object;
        return o;
    }
    
    int main()
    {
        object *new_obj = object_factory();
        delete new_obj; //we have to remember to delete this
    }

The above definition of object\_factory requires the user to "remember to manually free-up space allocated to the object", which is always a bad option. So the alternative is to return the object instead of its reference.

Object factory returning object:-

    object object_factory()
    {
        object o{}; //invoking default constructor
        return o; //invokes destructor implicitly
    }
    
    int main()
    {
        object new_obj = object_factory(); //invokes copy-constructor twice
    }

The above code is "neat". Below is the sequence of events that happen

1) default constructor is invoked when o is created

2) copy-constructor is invoked to make a 'copy of o', when return is called in object\_factory

3) destructor is invoked to destroy o when it goes out of scope at end of object\_factory

4) copy-constructor is invoked to make new\_obj from the 'copy of o'

5) destructor is invoked to destroy 'copy of o'

6) destructor is invoked to destroy new\_obj when it goes out of scope at the end of main.

As you can see, copy-constructor is called "twice". If the object happens to be large, copying it twice is an expensive operation. Can we avoid copying? Actually when o goes out of scope at the end of object\_factory we can no longer use it. If we could somehow move the data elements from 'o to new\_obj' our purpose is served. So if we have a move-constructor in the class definition of "object", the above same code produces a different sequence of events.

&amp;#x200B;

1) default constructor is invoked when o is created

2) move-constructor is invoked to move 'data elements from o to create "copy of o"', when return is called in object\_factory

3) destructor is invoked to destroy o when it goes out of scope at end of object\_factory

4) move-constructor is invoked to move 'data elements from "copy of o" to create "new\_obj"'

5) destructor is invoked to destroy "copy of o"

6) destructor is invoked to destroy new\_obj when it goes out of scope at end of main

All copy-constructors have been replaced by less-expensive move-constructors. But when you compile the above code without "special compiler flags" the C++ compiler optimizes further. Below is the actual sequence of events you will see.

1) default constructor is invoked when o is created

2) reference of o is passed to obj\_new

3) destructor is invoked to destroy obj\_new/o(o and obj\_new both point to same object)

This is super cool. The C++ compiler has avoided 2 move-constructions &amp; 2 destructions!(steps 2 to 5)

    Note:- Inorder to force the compiler not to make the above optimization(to see move-constructors in action) pass the '-fno-elide-constructors' flag to C++ compiler.
## [11][C++ on Sea 2020 video - "C++ ecosystem: the renaissance edition" - Anastasia Kazakova](https://www.reddit.com/r/cpp/comments/hyoa89/c_on_sea_2020_video_c_ecosystem_the_renaissance/)
- url: https://www.youtube.com/watch?v=5NuEX6cUpFI
---

