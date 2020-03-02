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
## [2][ABI Breaks: Not just about rebuilding](https://www.reddit.com/r/cpp/comments/fc2qqv/abi_breaks_not_just_about_rebuilding/)
- url: https://www.reddit.com/r/cpp/comments/fc2qqv/abi_breaks_not_just_about_rebuilding/
---
Related reading:

[What is ABI, and What Should WG21 Do About It?](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/p2028r0.pdf)

[The Day The Standard Library Died](https://cor3ntin.github.io/posts/abi/)

# Q: What does the C++ committee need to do to fix large swaths of ABI problems?

## A: Absolutely nothing

On current implementations, `std::unique_ptr`'s calling convention causes some inefficiencies compared to raw pointers.  The standard doesn't dictate the calling convention of `std::unique_ptr`, so implementers could change that if they chose to.

On current implementations, `std::hash` will return the same result for the same input, even across program invocations.  This makes it vulnerable to cache poisoning attacks.  Nothing in the standard requires that different instances of a program produce the same output.  An implementation could choose to have a global variable with a per-program-instance seed in it, and have `std::hash` mix that in.

On current implementations, `std::regex` is extremely slow.  Allegedly, this could be improved substantially without changing the API of `std::regex`, though most implementations don't change `std::regex` due to ABI concerns.  An implementation could change if it wanted to though.  However, very few people have waded into the guts of `std::regex` and provided a faster implementation, ABI breaking or otherwise.  Declaring an ABI break won't make such an implementation appear.

None of these issues are things that the C++ committee claims to have any control over.  They are dictated by vendors and by the customers of the vendors.  A new vendor could come along and have a better implementation.  For customers that prioritize QoI over ABI stability, they could switch and recompile everything.

Even better, the most common standard library implementations are all open source now.  You could fork the standard library, tweak the mangling, and be your own vendor.  You can then be in control of your own ~~destiny~~ ABI, and without taking the large up-front cost of reinventing the parts of the standard library that you are satisfied with.  libc++ has a [LIBCXX\_ABI\_UNSTABLE](https://github.com/llvm/llvm-project/blob/master/libcxx/CMakeLists.txt#L125) configuration flag, so that you always get the latest and greatest optimizations.  libstdc++ has a `--enable-symvers=gnu-versioned-namespace` configuration flag that is ABI unstable, and it goes a long way towards allowing multiple libstdc++ instances coexist simultaneously.  Currently the libc++ and libstdc++ unstable ABI branches don't have many new optimizations because there aren't many contributions and few people use it.  I will choose to be optimistic, and assume that they are unused because people were not aware of them.

If your only concern is ABI, and not API, then vendors and developers can fix this on their own without negatively affecting code portability or conformance.  If the QoI gains from an ABI break are worth a few days / weeks to you, then that option is available *today*.

# Q: What aspects of ABI makes things difficult for the C++ committee.

## A: API and semantic changes that would require changes to the ABI are difficult for the C++ committee to deal with.

There are a lot of things that you can do to a type or function to make it ABI incompatible with the old type.  The C++ committee is reluctant to make these kinds of changes, as they have a substantially higher cost.  Changing the layout of a type, adding virtual methods to an existing class, and changing template parameters are the most common operations that run afoul of ABI.

# Q: Are ABI changes difficult for toolchain vendors to deal with?

## A1: For major vendors, they difficulty varies depending on the magnitude of the break.

Since GCC 5 dealt with the std::string ABI break, [GCC has broken the language ABI 6 other times](https://gcc.gnu.org/onlinedocs/gcc/C_002b_002b-Dialect-Options.html), and most people didn't even notice.  There were several library ABI breaks (notably return type changes for std::complex and associative container erase) that went smoothly as well.  Quite a few people noticed the GCC 5 `std::string` ABI changes though.

In some cases, there are compiler heroics that can be done to mitigate an library ABI change.  You will get varying responses as to whether this is a worthwhile thing to do, depending on the vendor and the change.

If the language ABI changes in a large way, then it can cause substantially more pain.  GCC had a major language ABI change in GCC 3.4, and that rippled out into the library.  Dealing with libstdc++.so.5 and libstdc++.so.6 was a major hassle for many people, myself included.

## A2: For smaller vendors, the difficulty of an ABI break depends on their customer base.

These days, it's easier than ever to be your own toolchain vendor.  That makes you a vendor with excellent insight into how difficult an ABI change would be.

# Q: Why don't you just rebuild after an ABI change?

## A1: Are you rebuilding the standard library too?

Many people will recommend not passing standard library types around, and not throwing exceptions across shared library boundaries.  They often forget that at least one very commonly used shared library does exactly that... your C++ standard library.

On many platforms, there is usually a system C++ standard library.  If you want to use that, then you need to deal with standard library types and exceptions going across shared library boundaries.  If OS version N+1 breaks ABI in the system C++ standard library, the program you shipped and tested with for OS version N will not work on the upgraded OS until you rebuild.

## A2: Sometimes, rebuilding isn't enough

Suppose your company distributes pre-built programs to customers, and this program supports plugins (e.g. Wireshark dissector plugins).  If the plugin ABI changes, in the pre-built program, then all of the plugins need to rebuild.  The customer that upgrades the program is unlikely to be the one that does the rebuilding, but they will be responsible for upgrading all the plugins as well.  The customer cannot effectively upgrade until the entire ecosystem has responded to the ABI break.  At best, that takes a lot of time.  More likely, some parts of the ecosystem have become unresponsive, and won't ever upgrade.

This also requires upgrading large swaths of a system at once.  In certain industries, it is very difficult to convince a customer to upgrade anything at all, and upgrading an entire system would be right out.

Imagine breaking ABI on a system library on a phone.  Just getting all of the apps that your company owns upgraded and deployed at the same time as the system library would be a herculean effort, much less getting all the third party apps to upgrade as well.

There are things you can do to mitigate these problems, at least for library and C++ language breaks on Windows, but it's hard to mitigate this if you are relying on a system C++ standard library.  Also, these mitigations usually involve writing more error prone code that is less expressive and less efficient than if you just passed around C++ standard library types.

## A3: Sometimes you can't rebuild everything.

Sometimes, business models revolve around selling pre-built binaries to other people.  It is difficult to coordinate ABI changes across these businesses.

Sometimes, there is a pre-built binary, and the company that provided that binary is no longer able to provide updates, possibly because the company no longer exists.

Sometimes, there is a pre-built binary that is a shared dependency among many companies (e.g. OpenSSL).  Breaking ABI on an upgrade of such a binary will cause substantial issues.

# Q: What tools do we have for managing ABI changes?

## A: Several, but they all have substantial trade-offs.

The most direct tool is to just make a new thing and leave the old one alone.  Don't like `std::unordered_map`? Then make `std::open_addressed_hash_map`.  This technique allows new and old worlds to intermix, but the translations between new and old must be done explicitly.  You don't get to just rebuild your program and get the benefits of the new type.  Naming the new things becomes increasingly difficult, at least if you decide to not do the "lazy" thing and just name the new class `std::unordered_map2` or `std2::unordered_map`.  Personally, I'm fine with slapping a version number on most of these classes, as it gives a strong clue to users that we may need to revise this thing again in the future, and it would mean we might get an incrementally better hash map without needing to wait for hashing research to cease.

`inline` namespaces are another tool that can be used, but they solve far fewer ABI problems than many think.  Upgrading a type like `std::string` or `std::unordered_map` via `inline` namespaces generally wouldn't work, as user types holding the upgraded types would also change, breaking those ABIs.  `inline` namespaces can probably help add / change parameters to functions, and may even extend to updating empty callable objects, but neither of those are issues that have caused many problems in the C++ committee in the past.

Adding a layer of indirection, similar to COM, does a lot to address stability *and* extensibility, at a large cost to performance.  However, one area that the C++ committee hasn't explored much in the past is to look at the places where we already have a layer of indirection, and using COM-like techniques to allow us to add methods in the future.  Right now, I don't have a good understanding of the performance trade-offs between the different plug-in / indirect call techniques that we could use for things like `std::pmr::memory_resource` and `std::error_category`.

# Q: What can I do if I don't want to pay the costs for ABI stability?

## A: Be your own toolchain vendor, using the existing open-source libraries and tools.

If you are able to rebuild all your source, then you can point all your builds at a custom standard library, and turn on (or even make your own) ABI breaking changes.  You now have a competitive advantage, and you didn't even need to amend an international treaty (the C++ standard) to make it happen!  If your changes were only ABI breaking and not API breaking, then you haven't even given up on code portability.

Note that libc++ didn't need to get libstdc++'s permission in order to coexist on Linux.  You can have multiple standard libraries at the same time, though there are some technical challenges created when you do that.

# Q: What can I do if I want to change the standard in a way that is ABI breaking?

## A1: Consider doing things in a non-breaking way.

## A2: Talk to compiler vendors and the ABI Review Group (ARG) to see if there is a way to mitigate the ABI break.

## A3: Demonstrate that your change is so valuable that the benefit outweighs the cost, or that the cost isn't necessarily that high.

# Assorted points to make before people in the comments get them wrong

* I'm neither advocating to freeze ABI, nor am I advocating to break ABI.  In fact, I think those questions are too broad to even be useful.
* Fixing `std::unordered_map`'s performance woes would require an API break, as well as an ABI break.
* I have my doubts that `std::vector` could be made substantially faster with only an ABI break.  I can believe it if it is also coupled with an API break in the form of different exception safety guarantees.  Others are free to prove me wrong though.
* Making `&lt;cstring&gt;` `constexpr` will probably be fine.  The ABI issues were raised and addressed for `constexpr` `&lt;cmath&gt;`, and that paper is waiting in LWG.
* Filters on recursive\_directory\_iterators had additional concerns beyond ABI, and there wasn't consensus to pursue, even if we chose a different name.
* Making destructors implicitly `virtual` in polymorphic classes would be a massive cross-language ABI break, and not just a C++ ABI break, thanks to COM.  You'd be breaking the entire Windows ecosystem.  At a minimum, you'd need a way to opt out of `virtual` destructors.
* Are you sure that you are arguing against ABI stability?  Maybe you are arguing against backwards compatibility in general.
## [3][In-class Member Initialisation: From C++11 to C++20](https://www.reddit.com/r/cpp/comments/fc9iiz/inclass_member_initialisation_from_c11_to_c20/)
- url: https://www.bfilipek.com/2015/02/non-static-data-members-initialization.html
---

## [4][Quill - An asynchronous low latency logging library (C++14)](https://www.reddit.com/r/cpp/comments/fcbflb/quill_an_asynchronous_low_latency_logging_library/)
- url: https://github.com/odygrd/quill
---

## [5][In the unlikely event you want to abuse macros to test glsl shader code](https://www.reddit.com/r/cpp/comments/fc4v95/in_the_unlikely_event_you_want_to_abuse_macros_to/)
- url: https://www.reddit.com/r/cpp/comments/fc4v95/in_the_unlikely_event_you_want_to_abuse_macros_to/
---
Working (but not complete) [GLSL shim](https://github.com/jeremyong/klein/blob/master/test/glsl_shim.hpp)

Sample GLSL code [being tested](https://github.com/jeremyong/klein/blob/master/glsl/klein.glsl)

Sample [tests](https://github.com/jeremyong/klein/blob/master/test/test_glsl.cpp)

In short, code like the following function works in C++ with a bunch of macros that expand to a union with 261 members. The rest of the GLSL API could likely be supported in the same manner (and it could be adapted to HLSL as well).


    motor mm_mul(in motor a, in motor b)
    {
        motor c;
        vec4 a_zyzw = a.p1.zyzw;
        vec4 a_ywyz = a.p1.ywyz;
        vec4 a_wzwy = a.p1.wzwy;
        vec4 c_wwyz = b.p1.wwyz;
        vec4 c_yzwy = b.p1.yzwy;

        c.p1 = a.p1.x * b.p1;
        vec4 t = a_ywyz * c_yzwy;
        t += a_zyzw * b.p1.zxxx;
        t.x = -t.x;
        c.p1 += t;
        c.p1 -= a_wzwy * c_wwyz;

        c.p2 = a.p1.x * b.p2;
        c.p2 += a.p2 * b.p1.x;
        c.p2 += a_ywyz * b.p2.yzwy;
        c.p2 += a.p2.ywyz * c_yzwy;
        t = a_zyzw * b.p2.zxxx;
        t += a_wzwy * b.p2.wwyz;
        t += a.p2.zxxx * b.p1.zyzw;
        t += a.p2.wzwy * c_wwyz;
        t.x = -t.x;
        c.p2 -= t;
        return c;
    }
## [6][COAT: C++ EDSL for easier just-in-time code generation](https://www.reddit.com/r/cpp/comments/fbvkxp/coat_c_edsl_for_easier_justintime_code_generation/)
- url: https://tetzank.github.io/posts/coat-edsl-for-codegen/
---

## [7][Programming 3D Reconstruction Real Time](https://www.reddit.com/r/cpp/comments/fc72zr/programming_3d_reconstruction_real_time/)
- url: https://medium.com/@taying.cheng/understanding-real-time-3d-reconstruction-and-kinectfusion-33d61d1cd402
---

## [8][Looking for a good place to start.](https://www.reddit.com/r/cpp/comments/fc9xq2/looking_for_a_good_place_to_start/)
- url: https://www.reddit.com/r/cpp/comments/fc9xq2/looking_for_a_good_place_to_start/
---
I have a very basic knowledge of c++ and I’m looking to learn it. I have just started watching TheNewBostons c++ playlist on YouTube [playlist here](https://www.youtube.com/playlist?list=PL6gx4Cwl9DGAKIXv8Yr6nhGJ9Vlcjyymq)

I ultimately want to get into developing cheats for games I play. Do you think his tutorials will teach me all the essentials there is to c++ leaving me competent enough to develop my own cheat?
## [9][A journey to searching Have I Been Pwned database in 49μs (C++)](https://www.reddit.com/r/cpp/comments/fc1dsi/a_journey_to_searching_have_i_been_pwned_database/)
- url: http://stryku.pl/poetry/okon.php
---

## [10][How to create an LED rave mask using Arduino and C++ to automatically animate effects while using little memory.](https://www.reddit.com/r/cpp/comments/fbxfak/how_to_create_an_led_rave_mask_using_arduino_and/)
- url: https://armaizadenwala.com/blog/how-to-create-a-led-rave-mask-using-arduino/
---

## [11][Do you need to use an IDE to fit in in a professional C++ team?](https://www.reddit.com/r/cpp/comments/fc3tg5/do_you_need_to_use_an_ide_to_fit_in_in_a/)
- url: https://www.reddit.com/r/cpp/comments/fc3tg5/do_you_need_to_use_an_ide_to_fit_in_in_a/
---
The answer might seem like an obvious "no" but please hear me out. 

A while back I worked at a company where \*everyone\* used intelliJ for the [play](https://www.playframework.com/)\-[angular](https://angular.io/) stack we were working on. I've never liked IDEs because 1. they're bloated resource hogs, 2. they take a long time to master, and 3. they hide a ton of details under the hood. I much prefer using smart text editors like vim and vscode. I also much prefer creating bash scripts and config files so that every command/setting is in a clear place, rather than being buried within a complex sequence of mouse clicks.

So I was the only one at this company not using intelliJ and I found that this made it really difficult to work with everyone else since everything they did was done through intelliJ, including VC and CI. This experience and feeling that I could only work on a professional java stack if I used intelliJ really put me off Java.

I'm considering diving into C/C++, but first I'd like to know whether I would practically speaking need to master an IDE in order to fit in with a professinal crew of C++ developers. According to [this survey](https://www.jetbrains.com/lp/devecosystem-2019/cpp/), about 1/3 of C++ developers use a smart text editor, but I'd like to hear the thoughts of people within the actual industry. Thanks!
