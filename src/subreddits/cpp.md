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
* **Don't** use URL shorteners. [reddiquette](https://www.reddit.com/wiki/reddiquette) forbids them because they're opaque to the spam filter.
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
## [2][Advice on debugging C++ on Linux](https://www.reddit.com/r/cpp/comments/el6y9e/advice_on_debugging_c_on_linux/)
- url: https://www.reddit.com/r/cpp/comments/el6y9e/advice_on_debugging_c_on_linux/
---
I always use gdb for C/C++ programs. Unfortunately, latest features of C++ don't seem to fit well with gdb.

An example is a lambda function calling about 5 additional stack frames, and most STL stuff (std::unordered\_map) seems to hide away its internals, so that it makes it hard for me to find each element's exact memory location or its values.

It has not been much of an issue till now, since I don't really like using these modern features anyways. But right now the company is making me debug another person's program, and these lambdas and STL containers are really killing me right now...

Any general advice? A better debugger perhaps? Or gdb with plugins?
## [3][Waiting for std::embed: Very Large Arrays in Clang](https://www.reddit.com/r/cpp/comments/ekvk0n/waiting_for_stdembed_very_large_arrays_in_clang/)
- url: https://cor3ntin.github.io/posts/arrays/
---

## [4][A New Decade, A New Tool](https://www.reddit.com/r/cpp/comments/ekwb4y/a_new_decade_a_new_tool/)
- url: https://vector-of-bool.github.io/2020/01/06/new-decade.html
---

## [5][A priority queue implementation in CUDA applied to the many-to-many shortest path problem](https://www.reddit.com/r/cpp/comments/elac6i/a_priority_queue_implementation_in_cuda_applied/)
- url: https://github.com/crosetto/cupq/
---

## [6][On build systems: attempt at finding right abstractions](https://www.reddit.com/r/cpp/comments/el8o4m/on_build_systems_attempt_at_finding_right/)
- url: https://www.reddit.com/r/cpp/comments/el8o4m/on_build_systems_attempt_at_finding_right/
---
&gt;So, let me annoy everyone with my rant. If want to get to something more substantial just skip ahead.  
&gt;  
&gt;All this situation with build systems and hype around it can't help but remind me about good old [xkcd](https://xkcd.com/927/) panel. This situation seems completely paradoxical. We are ready to happily dump hundreds of human hours to discuss and design an obscure standard library feature which is required by 3.5 people and still are not willing to figure out and agree on way we build our software which is required by everyone. And we still claim to be grand C++ programmers who write most performant, efficient and elegant code out there. Bunch of us wrote most complicated systems in the world. And still look at the stuff we use every day. Ugh...  
&gt;  
&gt;There is only one conclusion that comes to my mind: we just SUCK at designing our own tools.

&amp;#x200B;

&gt;I don't want to come as all-knowing or claim I found a panacea. This post is mostly a brain dump of what was lazily brewing in the back of my mind for some time. It could be completely wrong, it could be far away from real world, I don't care. I just hope it will help us find the right direction.  
&gt;  
&gt;And no, there is not need to remind me. As a fellow C++ programmer I absolutely adore convoluted solutions and overly complex systems.

When we design something, first we need to decide our goals. For build system it would be:

1. Ease distribution and consumption of libraries. 
2. Reproducible builds. Holy grail of C++.

A lot of people focus on first point and completely forget about the second one, although it is equally as important.

Let's not dig too deep and take compilation model from what we have nowadays. It seems to be fairly stable and widely adopted. According to it there is a number of atomic (within reason) operations that need to be performed:

1. Compile TU into object file
2. Compile TU into AST dump\*
3. Assemble module interface\*\*
4. Assemble module body from related TUs\*\*
5. Link object files and module bodies into binary artifact
6. Configure build options\*\*\*
7. "Install" artifacts (move to reasonably unique and accessible location in filesystem)
8. Fetch sources/precompiled artifacts

\* Probably required for module bodies (templates?). AST is just an example - can be any other format that makes sense (including toolchain-specific)

\*\* Haven't been following module hype train. Are there any details on concrete implementations yet? I assume this is more or less what's supposed to be going on, but some operations here might be unnecessary (like interface assembly).

\*\*\* I believe I've seen some people (indirectly?) objecting to this step ("we don't need meta-build systems like CMake", although they might have a different objection point to this), but realistically this step is absolutely necessary as it includes stuff like choice between release/debug build, choice between target architecture etc.

What's come to mind when looking at the list? I guess

* 1-5 are just tiring invocations of compiler toolchain tools
* 6
* 7, 8 do basically the same thing: manage files in your system

When we talk about toolchain invocations we actually need to know which compiler we use to correctly issue compilation flags. Like call gcc for TUs but ld for linking exe, pass /Ox when handling msvc, but sometimes its just  clang++ for everything, and yeah don't tell linker weird stuff when compiling with asan. I mean... WTF is this mess? Why in all three hells your building systems needs to know all of that?

It's actually a complete disaster if you squint your eyes. We leak implementation details all over the place. Who cares if your toolchain has one executable or 10? Who cares if its dashes or slashes? All I want to ask is "do (one of 1, 2, 3, 4, 5)", why anything else should concern me? 

I mean at this point it makes sense to just have one executable which can be told what I *want* to do, not *how* to do it. Isn't it how good APIs work? Btw clang team apparently come to similar idea much earlier than me and hid all their stuff behind `clang++` executable. Anyways this is where I think we should start counting our tools: **Toolchain Abstraction Tool**.

&amp;#x200B;

The second part is resolving dependencies. TUs consume modules, module consume modules, all that. Basically speaking, if modules stuff is represented as files, we can explicitly feed those files in invocations where they are needed. Then ordering invocation is "trivial". Well at the very least in context of state-of-the-art contemporary systems: we already do this with headers. Needs some tweaking at most. 

What deserves much more attention is optimization techniques. Currently we have 3 (correct me if I'm wrong):

* Parallel command execution. Will not go anywhere, although will be a little bit more tricky.
* Incremental builds. Always watching over us.
* Databases of precompiled files. Newcomer. Problematic to integrate as it stands now.

To manage all this we can have another tool: **Script Execution Tool**.

&amp;#x200B;

Next comes configuring our project. **Project Configuration Tool**. Part of CMake fills similar role, and you can easily witness how screwed up everything is when you are trying to support multiple targets and multiple toolchains.

&amp;#x200B;

**File Management Tool**. Tucks your files where they need to be, fetches from internet, didn't give it enough thought, something else?

&amp;#x200B;

So overall we can round up the zoo to whole... **five** tools:

1. **Toolchain abstraction tool**
2. **Script execution tool**
3. **Project configuration tool**
4. **File management tool**
5. **Virtual environment management tool**

Yes, I added a fifth one, but everything in order.

&amp;#x200B;

**Toolchain Abstraction**

Responsibilities:

* Define and provide unified "high"-level toolchain invocation commands (similar to listed above) and transform those into toolchain-specific ones. Yes I also cringe at unified, but let's *temporary* agree on something at least.
* Provide similar support for various widely used compiler switches (optimization nobs, warnings, `-fno-rtti`, other fun stuff)
* Bring toolchain switching  down to change in one flag.
* Provide query API for various features (like "do you have asan?")
* Toolchain verification. I've seen CMake and autotools do something like this, but a tool which is more closely related to the subject will probably do better?
* Toolchain versioning. Maintain multiple versions, tool selections (ld vs gold vs...), standard libararies (clang has access to full 3), custom patches and whatnot - without destroying any of your projects.
* Produce/consume *toolchain description scripts*.
* Provide helpers, explanations and search over supported options, information about toolchain support/compliance. I feel like compiler options explanations are often as cryptic as flags themselves. And yeah, don't forget, there are new people too, for them modern toolchain is just black magic box as it stands now.

**Script execution**

Responsibilities:

* Takes a script which contains a list of tool invocations. Assumptions: 1) no specific order 2) each invocation is tool+config flags, input files, output files 3) as long as tool is the same, config is the same, input is the same = output is also same
* There are no conditions or configuration at this step! Adding this makes many features much more difficult to correctly implement. Also it happens right after project configuration. Need something - put it there.
* 2nd point provides a way to infer dependencies between files and therefore order in which commands must be executed. And no, specifying order doesn't achieve much since if we want to split the work over multiple threads, to do it correctly we need to build dependency graph anyways.
* 2nd point is also a reason to have modules build as separate invocations (so we can tie module object files, module body/interface and TUs that consume it). Btw its also a reason *to have module files to be explicitly specified in invocation instead of letting toolchain to infer it magically.* This way we can employ power of dependency guessing and optimization techniques from this tool. Also its a bookkeeping for reproducible builds.
* 3rd point is required to enable incremental builds. Also means that builds are deterministic.
* Handle optimization techniques (parallelization, incremental builds, object file database, whatever else we invent)
* Note that we are not limited to toolchain executables, anything can be invoked as long as it complies to conditions.
* There already exist tools which look in this direction, for ex. [Ninja](https://ninja-build.org/).

**Project configuration**

Responsibilities:

* Tool operates with *two* scripts: project description script and project configuration script.
* *Project description script*, well, it describes your project. Which files to use, what artifacts to create, toolchain flags, dependencies, project-specific flags, other meta-information.
* *Project configuration script* fills the blanks in description script. Most prominent are: toolchain choice, optimization levels, debug information, project-specific flags, custom flags/additions.
* Supports, again, two important operations: produce configurations for dependencies and produce build script for further execution.
* Note that we don't try and recursively execute on dependencies! We only tell what do we want to get, it's not our job to find and deliver them!

**File management**

Responsibilities:

* Fetch stuff from net: library lists, sources, precompiled binaries etc.
* Locate and tuck stuff locally. Each artifact so far can be identified with 3 configuration files: project description, project configuration, toolchain description. The first one can maybe be left out but then customly modified variants can be an issue.
* Create installations of the project, i. e. assemble all those dependencies from across filesystem into one spot, so it can be sanely distributed.
* Probably strongly integrated into virtual environment.

**Virtual environment management**

Responsibilities:

* Isolate project from system and other projects.
* Primary function: being asked to make a specific project/tool with specific requirements (configuration) available. Environment then decides how to handle it: find something locally, download, build - whatever it takes.
* Primary function: Manage PATH and dynamic library locations. Bunch of backstage magic is begging to happen here.
* Secondary function: Set up defaults to keep things DRY (like your toolchain).

This may seem unnecessary, but it is essential step towards reproducible builds: now we have (almost) perfect information about what is available to the project. True, that some of it can inferred (direct project dependencies) but some are not (like exact toolchain configuration, or other tools).

Also others (for ex. Python) already realized usefulness of such abstraction and are making tools to handle it while we are still doing our woodoo dances around compilers.

You may ask now, WTF I want to make my project in one line. And yes, now you can have it, the 6th tool: **Project management tool**. Let it make your life easy:

* Set up new projects, define project structure, conventions whatever else.
* Generate project description/configuration files.
* Execute all necessary steps to actually build a project.
* Integrate into your IDE nicely so you never even have to touch command line.

But seriously I don't care if it's easy to set up project in one line or easily link a famous library when everything is a dumpster fire and any step left or right leads to hours of suffering and guide/manual smoking (yes, CMake, I look at you).

I don't know, maybe it's just me, but everyone for some reason jumps directly to last item on the list and handwaves the rest. It's literally the least important one, and obviously most desirable one: a nice facade hiding complex machine. Except such build tools usually find crutches behind itself instead.

This is where my mind is now. Maybe some abstraction levels are wrong. Maybe we need to merge/split something, or completely rework the workflow? I don't know. If you know - help everyone to figure it out.

Anyways I'm super tired right now, hope this wall-of-text makes at least some sense.
## [7][Anjuta and Test Driven Development](https://www.reddit.com/r/cpp/comments/elbiiy/anjuta_and_test_driven_development/)
- url: https://www.reddit.com/r/cpp/comments/elbiiy/anjuta_and_test_driven_development/
---
I love working with Anjuta. It was with me during my early days of programming on Linux and taught me a few basics regarding creating programming projects.

Now, the question. I'm working on a rather large C++ library, and I want to implement Test Driven Development, but haven't really found anything that will work natively on Anjuta. I've previously worked with Qt's TDD tools, but have very little experience using others. So this is a kind of growth opportunity for me and I'd like to see what I can use.

I know both Google and Boost provide TDD tools, but which would you recommend or are there others you'd think about using with Anjuta?

Any and all help is appreciated.
## [8][Consteval build system: what if we do not build systems anymore, just a compiler](https://www.reddit.com/r/cpp/comments/el0xvv/consteval_build_system_what_if_we_do_not_build/)
- url: https://www.reddit.com/r/cpp/comments/el0xvv/consteval_build_system_what_if_we_do_not_build/
---
Just tought today about having a file with a name like compileme.cpp in the project which has a bunch of consteval functions (consteval std::filesystem evtl. needed) ala cmake e.g. add\_executable, add\_library, target\_sources etc.

those functions could be executed at compile time by the compiler-&gt; no need for a build system ala cmake, make etc. anymore...

what do you think? which problems do you see with such an approach? Could this be probably standartized?
## [9][Why CMake encourages distributing it's *.cmake files instead of *.pc for libs configuration.](https://www.reddit.com/r/cpp/comments/el9b0u/why_cmake_encourages_distributing_its_cmake_files/)
- url: https://www.reddit.com/r/cpp/comments/el9b0u/why_cmake_encourages_distributing_its_cmake_files/
---
Some time ago it was a standard that must libraries distributed XYZ.pc files for pkg-config that could be later consumed by any build system (or none). Recently I see trend that it's encouraged to distribute appropriate FindXYZ.cmake or XYZConfig.cmake.

Why is that so? What additional information can be conveyed through cmake files that is crucial for libraries?
## [10][A C++ memory hunter review](https://www.reddit.com/r/cpp/comments/eku762/a_c_memory_hunter_review/)
- url: https://oopscenities.net/2020/01/06/deleaker-part-0-intro/
---

## [11][Flecs 1.2: Introducing a C++11 API](https://www.reddit.com/r/cpp/comments/ekv776/flecs_12_introducing_a_c11_api/)
- url: https://www.reddit.com/r/cpp/comments/ekv776/flecs_12_introducing_a_c11_api/
---
Link to the release: [https://github.com/SanderMertens/flecs/releases/tag/v1.2](https://github.com/SanderMertens/flecs/releases/tag/v1.2)

Flecs is an Entity Component System written for C89, C99 and now C++11! Its unique features are a super fast SoA-based storage, a fully-fledged (nested) prefab workflow, hierarchies, time management and a plug&amp;play module system.

If you'd like to know more about ECS, see [https://github.com/SanderMertens/ecs-faq](https://github.com/SanderMertens/ecs-faq).

Flecs 1.2 features a new header-only C++11 API that follows modern C++ best practices. For example code see: [https://github.com/SanderMertens/flecs/tree/master/examples/cpp](https://github.com/SanderMertens/flecs/tree/master/examples/cpp)

Other new features in 1.2 are:

* Snapshots, a fast &amp; lightweight mechanism for restoring a game to a previous point in time
* A new blob serializer for streaming games to/from a disk or network
* On demand systems that are only executed when there is interest in their output
* A new statistics API for monitoring server-side Flecs applications
* 24 new API functions
* And much more, see the [release notes](https://github.com/SanderMertens/flecs/releases/tag/v1.2) for more information.
