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
## [2][2020-02 Prague ISO C++ Committee Trip Report — 🎉 C++20 is Done! 🎉](https://www.reddit.com/r/cpp/comments/f47x4o/202002_prague_iso_c_committee_trip_report_c20_is/)
- url: https://www.reddit.com/r/cpp/comments/f47x4o/202002_prague_iso_c_committee_trip_report_c20_is/
---
**[A very special video report from Prague.](https://youtu.be/AvPiGstxV_g)**

&amp;nbsp;

**C++20, the most impactful revision of C++ in a decade, is done!** 🎉🎊🥳

At the ISO C++ Committee meeting in Prague, hosted by [Avast](https://www.avast.com/), we completed the C++20 Committee Draft and voted to send the Draft International Standard (DIS) out for final approval and publication. Procedurally, it's possible that the DIS could be rejected, but due to our procedures and process, it's very unlikely to happen. This means that C++20 is complete, and in a few months the standard will be published.

During this meeting, we also [adopted a plan for C++23](https://wg21.link/P0592), which includes prioritizing a modular standard library, library support for coroutines, executors, and networking.

A big thanks to everyone who made C++20 happen - the proposal authors, the minute takers, the implementers, and everyone else involved!

[This was the largest C++ committee meeting ever - 252 people attended!](https://i.imgur.com/umP0qke.jpg) Our generous host, [Avast](https://www.avast.com/), did an amazing job hosting the meeting and also organized a lovely evening event for everyone attending.

&amp;nbsp;

This week, we made the following changes and additions to the C++20 draft:

- [Improved the context-sensitive recognition of 'module' and 'import' to make it easier for non-compiler tools such as build systems to determine build dependencies](https://wg21.link/P1857).
- Added [several new rangified algorithms](https://wg21.link/p1243).
- Added [`ranges::ssize`](https://wg21.link/p1970).
- Refined the meaning of 'static' and 'inline' in module interfaces ([P1779](https://wg21.link/P1779) and [P1815](https://wg21.link/P1815)).
- Resolved a lot of open core language and library issues and made many substantial improvements to specification.

&amp;nbsp;

The following notable features are in C++20:

- [**Modules**](https://en.cppreference.com/w/cpp/language/modules).
- [**Coroutines**](https://en.cppreference.com/w/cpp/language/coroutines).
- [**Concepts**](https://en.cppreference.com/w/cpp/concepts).
- [**Ranges**](https://en.cppreference.com/w/cpp/ranges).
- **`constexpr`ification: [`constinit`](https://wg21.link/P1143), [`consteval`](https://wg21.link/P1073), [`std::is_constant_evaluated`](https://wg21.link/P0595), [`constexpr` allocation](https://wg21.link/P0784), [`constexpr` `std::vector`](https://wg21.link/P1004), [`constexpr` `std::string`](https://wg21.link/P0980), [`constexpr` `union`](https://wg21.link/P1330), [`constexpr` `try` and `catch`](https://wg21.link/P1002), [`constexpr` `dynamic_cast` and `typeid`](https://wg21.link/P1327).**
- **[`std::format("For C++{}", 20)`](https://wg21.link/P0645).**
-  [**`operator&lt;=&gt;`**](https://en.cppreference.com/w/cpp/utility#Relational_operators_and_comparison).
- [Feature test macros](https://en.cppreference.com/w/cpp/feature_test).
- [`std::span`](https://en.cppreference.com/w/cpp/container/span).
- [Synchronized output](https://en.cppreference.com/w/cpp/io#Synchronized_output).
- [`std::source_location`](https://wg21.link/P1208).
- [`std::atomic_ref`](https://en.cppreference.com/w/cpp/atomic/atomic_ref).
- [`std::atomic::wait`, `std::atomic::notify`, `std::latch`, `std::barrier`, `std::counting_semaphore`, etc](https://wg21.link/P1135).
- [`std::jthread` and `std::stop_*`](https://wg21.link/P0660).

&amp;nbsp;

****

# ABI Discussion

****

We had a very important discussion about [ABI stability and the priorities of C++](https://wg21.link/P1863) this week in a joint session of the Language Evolution and Library Evolution group.

Although there was strong interest in exploring how to evolve ABI in the future, we are not pursuing making C++23 a clean ABI breaking release at this time. We did, however, affirm that authors should be encouraged to bring individual papers for consideration, even if those would be an ABI break. Many in the committee are interested in considering targeted ABI breaks when that would signify significant performance gains.

&gt;
&gt; ‟How many C++ developers does it take to change a lightbulb?” — [@tvaneerd](https://twitter.com/tvaneerd/status/1060646821247119360)
&gt;
&gt; ‟None: changing the light bulb is an ABI break.” — [@LouisDionne](https://twitter.com/LouisDionne/status/1062724316905775106)
&gt;

&amp;nbsp;

*****

# Language Progress

*****

### Evolution Working Group Incubator (EWGI) Progress

*****

The EWG Incubator met for three days in Prague and looked at and gave feedback to 22 papers for C++23. 10 of those papers were forwarded to Evolution, possibly with some revisions requested. Notably:

- [Guaranteed copy elision for named return objects](https://wg21.link/p2025)
- [Generalized pack declaration and usage](https://wg21.link/P1858)
- [Member templates for local classes](https://wg21.link/p2044)
- [Object relocation in terms of move plus destroy](https://wg21.link/p1144)

Several papers received a lot of feedback and will return to the Incubator, hopefully in Varna:

- [A pipeline-rewrite operator](https://wg21.link/P2011)
- [Universal template parameters](https://wg21.link/P1985)
- [Partially mutable lambda captures](https://wg21.link/p2034)
- [C++ should support just-in-time compilation](https://wg21.link/p1609)
- [`move = bitcopies`](https://wg21.link/P1029)

Notably, [the proposed epochs language facility](https://wg21.link/p1881) received no consensus to proceed. One significant problem pointed out was that in a concepts and modules world, we really cannot make any language changes that may change the satisfaction of a concept for a set of types. If one TU thinks `C&lt;T&gt;` is true, but another TU in a later epoch thinks `C&lt;T&gt;` is false, that easily leads to ODR violations. Many of the suggested changes in the paper run afoul of this problem. However, we’re interested in solving the problem, so welcome an alternative approach.

*****

### Evolution Working Group (EWG) Progress

*****

The top priority of EWG was again fixing the final national body comments for C++20. Once that was done, we started looking at C++23 papers. We saw a total of 36 papers.

Papers of note:

* We adopted the [C++ IS schedule](https://wg21.link/P1000).
* We adopted a [plan for C++23](https://wg21.link/P0592).
* We adopted a [process for evolutionary proposals](https://wg21.link/P1999), to make sure that we reduce the chance that we’ll make mistakes
* We agreed to pursue the Undefined Behavior group’s effort to [document Core Undefined or Unspecified Behavior](https://wg21.link/P2118) going forward. They’re documenting all language undefined behavior that C++ contains today, and we agreed to document and justify any new language undefined behavior going forward.

We marked 3 papers as tentatively ready for C++23:

* [Make declaration order layout mandated](https://wg21.link/P1847)
* [Guaranteed copy elision for named return objects](https://wg21.link/P2025)
* [C++ Identifier Syntax using Unicode Standard Annex 31](https://wg21.link/P1949)

They’ll proceed to the Core language group at the next meeting if no issues are raised with these papers.

We continued reviewing [pattern matching](https://wg21.link/p1371). This is one of our top priorities going forward. It’s looking better and better as we explore the design space and figure out how all the corner cases should work. One large discussion point at the moment is what happens when no match occurs, and whether we should mandate exhaustiveness. There’s exploration around the expression versus statement form. We’re looking for implementation experience to prove the design.

We really liked [deducing `this`](https://wg21.link/P0847), a proposal that eliminates the boilerplate associated with having `const` and non-`const`, `&amp;` and `&amp;&amp;` member function overloads. It still needs wording and implementation experience, but has strong support.

We continue discussing floating-point [fixed-layout types](https://wg21.link/P1468) and [extended floating point types](https://wg21.link/P1467), which are mandating IEEE 754 support for the new C++ `float16_t`, `float32_t`, `float64_t`, and adding support for `bfloat16_t`.

[`std::embed`](https://wg21.link/P1040), which allows embedding strings from files, is making good progress.

In collaboration with the Unicode group, [named universal character escapes](https://wg21.link/P2071) got strong support.

[`if consteval`](https://wg21.link/P1938) was reviewed. We’re not sure this is exactly the right solution, but we’re interested in solving problems in this general area.

We saw a really cute paper on [deleting variable templates](https://wg21.link/P2041) and decided to expand its scope such that more things can be marked as `= delete` in the language. This will make C++ much more regular, and reduce the need for expert-only solutions to tricky problems.

&amp;nbsp;

*****

### Core Working Group (CWG) Progress

*****

The top priority of CWG was finishing processing national body comments for C++20. CWG spent most of its remaining time this week working through papers and issues improving the detailed specification for new C++20 features.

We finished reviewing four papers that fine-tune the semantics of modules:

- We [clarified the meaning of `static`](https://wg21.link/p1815) (and unnamed namespaces) in module interfaces: such entities are now kept internal and cannot be exposed in the interface / ABI of the module. In non-modules compilations, we deprecated cases where internal-linkage entities are used from external-linkage entities. (These cases typically lead to violations of the One Definition Rule.)

- We [clarified the meaning of `inline`](https://wg21.link/p1779) in module interfaces: the intent is that bodies of functions that are not explicitly declared `inline` are not part of the ABI of a module, even if those function bodies appear in the module interface. In order to give module authors more control over their ABI, member functions defined in class bodies in module interfaces are no longer implicitly `inline`.

- We tweaked the context-sensitive recognition of the `module` and `import` keyword in order to avoid changing the meaning of more existing code that uses these identifiers, and to make it more straightforward for a scanning tool to recognize these declarations without full preprocessing.

- We improved backwards compatibility with unnamed enumerations in legacy header files (particularly C header files). Such unnamed enumerations will now be properly merged across header files if they're reachable in multiple different ways via imports.

- We finalized some subtle rules for concepts: a [syntax gotcha](https://wg21.link/p2092r0) in `requires` expressions was fixed, and we allowed [caching of concept values](https://wg21.link/p2104r0), which has been shown to dramatically improve performance in some cases.

- We agreed to (retroactively, via the defect report process) [treat initialization of a `bool` from a pointer as narrowing](https://wg21.link/p1957r2), improving language safety.

- We added permission for a comparison function to be defaulted outside its class, so long as the comparison function is a member or friend of the class, for consistency and to allow a defaulted comparison function to be non-inline.

&amp;nbsp;

*****

# Library Progress

*****

### Library Evolution Working Group Incubator (LEWGI) Progress

*****

LEWGI met for three and a half days this week and reviewed 22 papers. Most of our work this week was on various numerics proposals during joint sessions with the Numerics group. A lot of this work may end up going into the proposed Numerics Technical Specification, whose [scope and goals we are working to define](https://wg21.link/P2004). We also spent a chunk of time working on [modern I/O](https://wg21.link/P1883) and concurrent data structures for the upcoming Concurrency Technical Specification Version 2.

LEWGI looked at the following proposals, among others:

- Numerics:
  - [Physical Units Library](https://wg21.link/P1930).
  - [Linear Algebra](https://wg21.link/P1385).
- Concurrency:
  - [Concurrent Queues](https://wg21.link/P1958).
- Low-level File I/O
  - [Mapped File Handle](https://wg21.link/P1883).
- Narrowing Conversions
  - [std::is_narrowing_conversion](https://wg21.link/P0870).
  - [std::narrow_cast](https://wg21.link/P1998).
- Random Numbers
  - [Improving std::random_device](https://wg21.link/P2058).
  - [Portable Distributions](https://wg21.link/P2059).
  - [Improving Engine Seeding](https://wg21.link/P2060).

&amp;nbsp;

*****

### Library Evolution Working Group (LEWG) Progress

*****

After handling the few remaining National Body comments to fix issues with C++20, LEWG focused on making general policy decisions about standard library design standards. For example, we [formally codified](https://wg21.link/P1851) the guidelines for concept names in the standard library, and clarified [SD-8](https://isocpp.org/std/standing-documents/sd-8-standard-library-compatibility), our document listing the compatibility guarantees we make to our users. Then we started looking at C++23 library proposals.

[Moved-from objects need not be valid](https://wg21.link/P2027) generated much internal discussion in the weeks leading up to the meeting as well as at the meeting itself.  While the exact solution outlined in the paper wasn’t adopted, we are tightening up the wording around algorithms on what operations are performed on objects that are temporarily put in the moved-from state during the execution of an algorithm.

The biggest C++23 news: LEWG spent an entire day with the concurrency experts of SG1 to review the [**executors** proposal](https://wg21.link/P0443) — we liked the direction! This is a huge step, which will enable **networking**, **audio**, **coroutine library support**, and more.

Other C++23 proposals reviewed include

- [a new `status_code` facility](https://wg21.link/P1028)
- [an ability for containers and allocators to communicate about the actual allocation size](https://wg21.link/P0401)
- [iterator range constructors for `std::stack` and `std::queue`](https://wg21.link/P1425)

We’ve also decided to [deprecate `std::string`’s assignment operator taking a `char`](https://wg21.link/P2037) (pending LWG).

&amp;nbsp;

*****

### Library Working Group (LWG) Progress

*****

The primary goals were to finish processing NB comments and to [rebase](https://wg21.link/P2081) the Library Fundamentals TS on C++20.  We met both of those goals.

We looked at all 48 open library-related NB comments and responded to them.  Some were accepted for C++20.  Some were accepted for C++20 with changes.  For some, we agreed with the problem but considered the fix to be too risky for C++20, so an issue was opened for consideration in C++23.  For many the response was “No consensus for change,” which can mean a variety of things from “this is not really a problem” to “the problem is not worth fixing.”

The last of the [mandating papers](https://wg21.link/P1460) was reviewed and approved.  All of the standard library should now be cleaned up to use the latest [library wording guidelines](https://wg21.link/P1369), such as using “Mandates” and “Constraints” clauses rather than “Requires” clauses.

Some time was spent going through the LWG open issues list.  We dealt with all open P1 issues (“must fix for C++20”).  Many of the open P2 issues related to new C++20 features were dealt with, in an attempt to fix bugs before we ship them.

This was Marshall Clow’s last meeting as LWG chair.  He received a standing ovation in plenary.

&amp;nbsp;

*****

### Concurrency and Parallelism Study Group (SG1) Progress

*****

SG1 focused on C++23 this week, primarily on driving executors, one of the major planned features on our roadmap. Executors is a foundational technology that we'll build all sorts of modern asynchronous facilities on top of, so it's important that we land it in the standard early in the C++23 cycle.

At this meeting, LEWG approved of the [executors](https://wg21.link/P0443R12) design, and asked  the authors to return with a full specification and wording for review at the next meeting.

SG1 reviewed and approved of a [refinement](https://wg21.link/P2006R0) to the design of the sender/receiver concepts. This change unifies the lifetime model of coroutines and sender/receiver and allows us to statically eliminate the need for heap allocations for many kinds of async algorithms.

Going forward, SG1 will start working on proposals that build on top of executors, such as concurrent algorithms, parallel algorithms work, networking, asynchronous I/O, etc.

&amp;nbsp;

*****

### Networking Study Group (SG4) Progress

*****

SG4 started processing review feedback on the [networking TS](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2018/n4771.pdf) aimed at modernizing it for inclusion in C++23. SG4 also reviewed a proposal to [unify low-level I/O with the high-level asynchronous abstractions](https://wg21.link/p2052) and gave feedback to the author.

&amp;nbsp;

*****

### Numerics Study Group (SG6) Progress

*****

The Numerics group met on Monday this week, and also jointly with LEWGI on Tuesday and Thursday, and with SG19 on Friday.

We reviewed papers on a number of topics, including:

- [The scope and goals for the Numerics TS](https://wg21.link/P2004).
- [Linear algebra](https://wg21.link/P1385).
- [Units library](https://wg21.link/P1935).

&amp;nbsp;

*****

### Compile-Time Programming Study Group (SG7) Progress

*****

[Circle](https://www.circle-lang.org/) is a fork of C++ that enables arbitrary compile-time execution (e.g. a compile-time `std::cout`), coupled with reflection to allow powerful meta-programming. SG7 was interested in it and [considered copying parts of it](https://wg21.link/P2043). However, concerns were [raised](https://wg21.link/P2062) about security and usability problems, so the ability to execute arbitrary code at compile-time was rejected.

Besides that, we also continued to make progress on C++ reflection including naming of reflection keywords and potential to enable lazy evaluation of function arguments.

We also looked at [the JIT proposal](https://wg21.link/P1609) and asked authors to try to unify the design with current reflection proposals.

&amp;nbsp;

*****

### Undefined Behavior Study Group (SG12)/Vulnerabilities Working Group (WG23) Progress

*****

We set out to [enumerate all undefined and unspecified behavior](https://wg21.link/P1705). We’ve decided that upcoming papers adding new undefined or unspecified behavior need to include rationale and examples.

SG12 also collaborated with the [MISRA standard](https://www.misra.org.uk/)  for coding standards in embedded systems to help them update the guidelines for newer C++ revisions.

&amp;nbsp;

*****

### Human Machine Interface and Input/Output Study Group (SG13) Progress

*****

SG13 had a brief presentation of extracts from the [2019 CppCon keynote](https://www.youtube.com/watch?v=5N4b-rU-OAA) featuring Ben Smith (from 1:05:00) 

We looked at [A Brief 2D Graphics Review](http://wg21.link/P2005R0) and encouraged exploration of work towards a separable color proposal.

Finally, we worked through the use cases in [Audio I/O Software Use Cases](http://wg21.link/P2054). We have a couple of weeks before the post meeting mailing deadline to collect additional use cases and will then solicit feedback on them from WG21 and the wider C++ community.

&amp;nbsp;



*****

### Tooling Study Group (SG15) Progress

*****

The Tooling study group met this week to continue work on the Module Ecosystem Technical Report. Three of the papers targeting the Technical Report are fairly mature at this point, so we've directed the authors of those papers to work together to create an initial draft of the Technical Report for the Varna meeting. Those papers are:

- [Dependency Information Format](https://wg21.link/P1689) 
- [Module Recipe and BMI Reuse](https://wg21.link/P1788)
- [User-Facing Lexicon and File Extensions](https://wg21.link/P1838)

This draft will give us a shared vehicle to start hammering out the details of the Technical Report, and a target for people to write papers against.

We also discussed two proposals, about [debugging C++ coroutines](https://wg21.link/P2073) and [asynchronous call stacks](https://wg21.link/P2074).

&amp;nbsp;

*****

### Machine Learning Study Group (SG19) Progress

*****

SG14 met in Prague in a joint session with SG19 (Machine Learning).

The freestanding library took a few steps forward, with some interesting proposals, including [Freestanding Language: Optional `::operator new`](https://wg21.link/P2013R0)

One of the biggest decisions was on [Low-Cost Deterministic C++ Exceptions for Embedded Systems](https://www.research.ed.ac.uk/portal/files/78829292/low_cost_deterministic_C_exceptions_for_embedded_systems.pdf) which got great reactions. We will probably hear more about it!

&amp;nbsp;

*****

### Unicode and Text Study Group (SG16) Progress

*****

Our most interesting topic of the week concerned the interaction of execution character set and compile-time programming.  Proposed features for `std::embed` and reflection require the evaluation of strings at compile time and this occurs at translation phase 7.  This is after translation phase 5 in which character and string literals are converted to the execution character set.  These features require interaction with file names or the internal symbol table of a compiler.  In cross compilation scenarios in which the target execution character set is not compatible with the compiler’s host system or internal encoding, interesting things happen.  As in so many other cases, we found an answer in UTF-8 and will be recommending that these facilities operate solely in UTF-8.

We forwarded [Named Universal Character Escapes](https://wg21.link/P2071) and [C++ Identifier Syntax using Unicode Standard Annex 31](https://wg21.link/P1949) to EWG.
Both papers were seen by EWG this week and are on track for approval for C++23 in meetings later this year.

We forwarded [Naming Text Encodings to Demystify Them](https://wg21.link/P1885) to LEWG.

We declined to forward a paper to enhance `std::regex` to better support Unicode due to severe ABI restrictions; the `std::regex` design exposes many internal details of the implementation to the ABI and implementers indicated that they cannot make any significant changes.  Given the current state of `std::regex` is such that we cannot fix either its interface or its well-known performance issues, a number of volunteers agreed to bring a paper to deprecate `std::regex` at a future meeting.

&amp;nbsp;

*****

### Machine Learning Study Group (SG19) Progress

*****

SG19 met for a full day, one half day with SG14 (Low Latency), and one half day with SG6 (Numerics).

Significant feedback from a ML perspective was provided on [Simple Statistics functions](https://wg21.link/P1708), especially regarding the handling of missing data, non-numeric data, and various potential performance issues.

There was an excellent presentation of "Review of P1708: Simple Statistical Functions" which presented an analysis across Python, R, SAS and Matlab for common statistical methods.

The [graph library](http://wg21.link/P1709) paper had a great reaction, was also discussed, and will proceed.

Also, support for differentiable programming in C++, important for well-integrated support for ML back-propagation, was discussed in the context of [differentiable programming for C++](https://wg21.link/P2072).

&amp;nbsp;

*****

### Contracts Study Group (SG21) Progress

*****

In a half-day session, we discussed one of the major points of contention from previous proposals, which was the relationship between “assume” and “assert”, disentangling colloquial and technical interpretations. We also discussed when one implies the other, and which combinations a future facility should support.

- [Previous Disagreements](https://wg21.link/P2076)
- [Assumptions](https://wg21.link/P2064)
- [Portable Assumptions](https://wg21.link/P1774)

&amp;nbsp;

*****

# C++ Release Schedule

*****

***NOTE: This is a plan not a promise. Treat it as speculative and tentative. See [P1000](https://wg21.link/P1000) for the latest plan.***

* IS = International Standard. The C++ programming language. C++11, C++14, C++17, etc.
* TS = Technical Specification. "Feature branches" available on some but not all implementations. Coroutines TS v1, Modules TS v1, etc.
* CD = Committee Draft. A draft of an IS/TS that is sent out to national standards bodies for review and feedback ("beta testing"). 

Meeting | Location | Objective
-|-|-
~~2018 Summer LWG Meeting~~ | ~~Chicago~~ | ~~Work on wording for C++20 features.~~
~~2018 Fall EWG Modules Meeting~~ | ~~Seattle~~ | ~~Design modules for C++20.~~
~~2018 Fall LEWG/SG1 Executors Meeting~~ | ~~Seattle~~ | ~~Design executors for C++20.~~
~~2018 Fall Meeting~~ | ~~San Diego~~ | ~~C++20 major language feature freeze.~~
~~2019 Spring Meeting~~ | ~~Kona~~ | ~~C++20 feature freeze. C++20 design is feature-complete.~~
~~2019 Summer Meeting~~ | ~~Cologne~~ | ~~Complete C++20 CD wording. Start C++20 CD balloting ("beta testing").~~
~~2019 Fall Meeting~~ | ~~Belfast~~ | ~~C++20 CD ballot comment resolution ("bug fixes").~~
**2020 Spring Meeting** | **Prague** | **C++20 CD ballot comment resolution ("bug fixes"), C++20 completed.**
2020 Summer Meeting | Varna | First meeting of C++23.
2020 Fall Meeting | New York | Design major C++23 features.
2021 Winter Meeting | Kona | Design major C++23 features.
2021 Summer Meeting | Montréal | Design major C++23 features.
2021 Fall Meeting | 🗺️ | C++23 major language feature freeze.
2022 Spring Meeting | Portland | C++23 feature freeze. C++23 design is feature-complete.
2022 Summer Meeting | 🗺️ | Complete C++23 CD wording. Start C++23 CD balloting ("beta testing").
2022 Fall Meeting | 🗺️ | C++23 CD ballot comment resolution ("bug fixes").
2023 Spring Meeting | 🗺️ | C++23 CD ballot comment resolution ("bug fixes"), C++23 completed.
2023 Summer Meeting | 🗺️ | First meeting of C++26.

&amp;nbsp;

*****

# Status of Major C++ Feature Development

*****

***NOTE: This is a plan not a promise. Treat it as speculative and tentative.***

* IS = International Standard. The C++ programming language. C++11, C++14, C++17, etc.
* TS = Technical Specification. "Feature branches" available on some but not all implementations. Coroutines TS v1, Modules TS v1, etc.
* CD = Committee Draft. A draft of an IS/TS that is sent out to national standards bodies for review and feedback ("beta testing"). 

**Changes since last meeting are in bold.**

Feature | Status | Depends On | Current Target (Conservative Estimate) | Current Target (Optimistic Estimate)
-|-|-|-|-
[Concepts](https://en.cppreference.com/w/cpp/language/constraints) | Concepts TS v1 published and merged into C++20 | | C++20 | C++20
[Ranges](https://en.cppreference.com/w/cpp/experimental/ranges) | Ranges TS v1 published and merged into C++20 | Concepts | C++20 | C++20
[Modules](https://wg21.link/P1103) | Merged design approved for C++20 | | C++20 | C++20
[Coroutines](https://wg21.link/N4723) | Coroutines TS v1 published and merged into C++20 | | C++20 | C++20
[Executors](https://wg21.link/P1658) | New compromise design approved for C++23 | | C++26 | **C++23 (Planned)**
[Contracts](https://wg21.link/P0542) | Moved to Study Group | | C++26 | C++23
[Networking](https://wg21.link/N4711) | Networking TS v1 published | Executors | C++26 | **C++23 (Planned)**
[Reflection](https://wg21.link/P0194) | Reflection TS v1 published | | C++26 |C++23
[Pattern Matching](https://wg21.link/P1371) | | | C++26 | C++23
**[Modularized Standard Library](https://wg21.link/P1453)** | | | **C++23** | **C++23 (Planned)**

&amp;nbsp;

[**Last Meeting's Reddit Trip Report.**](https://www.eddit.com/r/cpp/comments/dtuov8/201911_belfast_iso_c_committee_trip_report/)

&amp;nbsp;

**If you have any questions, ask them in this thread!**

**Report issues by replying to the top-level stickied comment for issue reporting.**

&amp;nbsp;

&amp;nbsp;

*/u/blelbach, Tooling (SG15) Chair, Library Evolution Incubator (SG18) Chair*

*/u/bigcheesegs*

*/u/c0r3ntin*

*/u/jfbastien, Evolution (EWG) Chair*

*/u/arkethos (aka code_report)*

*/u/vulder*

*/u/hanickadot, Compile-Time Programming (SG7) Chair*

*/u/tahonermann, Text and Unicode (SG16) Chair*

*/u/cjdb-ns, Education (SG20) Lieutenant*

*/u/nliber*

*/u/sphere991*

*/u/tituswinters, Library Evolution (LEWG) Chair*

*/u/HalFinkel, US National Body (PL22.16) Vice Chair*

*/u/ErichKeane, Evolution Incubator (SG17) Assistant Chair*

*/u/sempuki*

*/u/ckennelly*

*/u/mathstuf*

*/u/david-stone, Modules (SG2) Chair and Evolution (EWG) Vice Chair*

*/u/je4d, Networking (SG4) Chair*

*/u/FabioFracassi, German National Body Chair*

*/u/redbeard0531*

*/u/nliber*

*/u/foonathan*

*/u/InbalL, Israel National Body Chair*

*/u/zygoloid, C++ Project Editor*

*⋯ and others ⋯*
## [3][C++ is NOT a superset of C: tentative definitions, implicit conversions, implicit declarations &amp;more](https://www.reddit.com/r/cpp/comments/f86rdq/c_is_not_a_superset_of_c_tentative_definitions/)
- url: https://www.youtube.com/watch?v=s3Cv0-U5bXc
---

## [4][Created my First Space Invaders Game with Raylib. Feedback would be appreciated.](https://www.reddit.com/r/cpp/comments/f87lm1/created_my_first_space_invaders_game_with_raylib/)
- url: /r/gamedev/comments/f87les/created_my_first_space_invaders_game_with_raylib/
---

## [5][EnTT v3.3.0 is out: Gaming meets Modern C++](https://www.reddit.com/r/cpp/comments/f7tzbd/entt_v330_is_out_gaming_meets_modern_c/)
- url: https://www.reddit.com/r/cpp/comments/f7tzbd/entt_v330_is_out_gaming_meets_modern_c/
---
# What's EnTT

[`EnTT`](https://github.com/skypjack/entt) is a header-only library written in **modern C++**.

It's mainly known for its innovative and very performing [entity-component-system](https://en.wikipedia.org/wiki/Entity_component_system) (ECS) model. However, it offers also [many other things](https://github.com/skypjack/entt#introduction) useful for development, from flexible tools for managing signals to an integrated reflection system and so on. That's not all: some new modules are under development and will soon become part of the library (or at least I hope so).

`EnTT` is also a production-ready, fully documented and battle-tested library with a 100% coverage. [Among others](https://github.com/skypjack/entt/wiki/EnTT-in-Action), It's currently used in **Minecraft** by Mojang and the **ArcGIS Runtime SDK** by Esri.

# Why it's important for you

You should care because the **best feedback** I've ever received is perhaps when someone told me that this library was a great source for studying modern C++. Maybe he was right, maybe not. It's not for me to say, but it may still interest you.

In any case, I hope I've given other material to those interested in it in this sense and I'm open to any feedback, because ours is a language in which you never stop learning. So, go for it!

# What else

In my free time I'm running the [ECS back and forth](https://skypjack.github.io/tags/#ecs) series (along with [other posts](https://skypjack.github.io/)). I'd like to start something new about the internals of this library and the development of a software with `EnTT`. The latter should serve as a step by step guide to using all the feature offered by the library. I cannot really set a deadline for this but I'd like to know if there is interest in such a series before to start, so any feedback is appreciated!

If you are using `EnTT` and want to tell me *hello* or which of your products relies on it, do not hesitate to [contact me](https://github.com/skypjack)! For everyone else interested in the library, the [wiki](https://github.com/skypjack/entt/wiki) contains more than what I've said here and the [gitter](https://gitter.im/skypjack/entt) channel is a great place to come and ask your first question!

I'm looking forward to hearing from you. :)
## [6][How much could we gain by completely breaking ABI?](https://www.reddit.com/r/cpp/comments/f7zls8/how_much_could_we_gain_by_completely_breaking_abi/)
- url: https://www.reddit.com/r/cpp/comments/f7zls8/how_much_could_we_gain_by_completely_breaking_abi/
---
Quoting from 2020-02 Prague Meeting report:
&gt; Although there was strong interest in exploring how to evolve ABI in the future, we are not pursuing making C++23 a clean ABI breaking release at this time.

Is there any big list of things that could be improved, if it was decided for C++23 to be an ABI-breaking release?

I found a good paper mentioning some:
http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/p2028r0.pdf

But it only listed a few, and gave just vague estimates on what would be real gains.

Is there a bigger list, listing more of existing problems, with measured possible speed-ups or memory gains?

I encountered one myself, that is not widely known: 
- `sizeof(std::random_device)` on GCC is enormous, cause it is a union with Mersenne Twister. It was MT on Windows until it got implemented as indeterminate recently.

And I'm sure there are many many more.

I'd love to know more about it. Understand, where along the way there were made some small mistakes, and how big is the technical debt we are dragging behind us.

How much are we losing, and for how long can we afford not to break ABI? Are these changes big enough, that we can expect someone to fork the language just for performance reasons?
## [7][C UWP APP in Visual Studio 2019 (Getting Started)](https://www.reddit.com/r/cpp/comments/f883k3/c_uwp_app_in_visual_studio_2019_getting_started/)
- url: https://www.youtube.com/watch?v=mE9nVXWgIMo&amp;feature=share
---

## [8][Counting Nanoseconds: Microbenchmarking C++ Code - David Gross - Meeting C++ 2019](https://www.reddit.com/r/cpp/comments/f7swf9/counting_nanoseconds_microbenchmarking_c_code/)
- url: https://www.youtube.com/watch?v=Czr5dBfs72U
---

## [9][Experimental Binary Package Support for vcpkg](https://www.reddit.com/r/cpp/comments/f7unlb/experimental_binary_package_support_for_vcpkg/)
- url: https://github.com/rkitover/vcpkg-binpkg-prototype
---

## [10][Bjarne Stroustrup on C++20's significance](https://www.reddit.com/r/cpp/comments/f7h6xn/bjarne_stroustrup_on_c20s_significance/)
- url: https://isocpp.org//blog/2020/02/bjarne-stroustrup-on-cpp20s-significance
---

## [11][Those that started c++ before 11 : How do you keep in touch with new features?](https://www.reddit.com/r/cpp/comments/f7dz9s/those_that_started_c_before_11_how_do_you_keep_in/)
- url: https://www.reddit.com/r/cpp/comments/f7dz9s/those_that_started_c_before_11_how_do_you_keep_in/
---
I have been coding "old C++" for a while now and started C++17 about a couple of years ago. How do you learn about use the latest features? For instance, I realized I haven't used any lambda captures lately. Now, most code I've written in the last 2 months hasn't really needed them, but I realized that I'm only using trivial features of modern C++. How do you "remember" to do stuff in the modern ways?
## [12][How to design exceptions for a (midsize desktop) application?](https://www.reddit.com/r/cpp/comments/f7oznb/how_to_design_exceptions_for_a_midsize_desktop/)
- url: https://www.reddit.com/r/cpp/comments/f7oznb/how_to_design_exceptions_for_a_midsize_desktop/
---
I'm looking for a book chapter/presentation/blog post/article describing how to design exceptions architecture for a (midsize desktop) application (like [https://github.com/dvorka/mindforger](https://github.com/dvorka/mindforger)). Can you please suggest a good resource?

I know how to declare, implement, throw/catch exception.

I would like to learn **best practices** for how to design exceptions (class) hierarchy, throw/catch strategy, patterns, when to use built-in and when custom exceptions, how to design application-wide error handling/debugging strategy or tips for handy exception fields to declare.

If there is any(/your :) open source project whose exceptions architecture/handling you like, can you please share it?
