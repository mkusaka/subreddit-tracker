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
**C++20, the most impactful revision of C++ in a decade, is done!** 🎉🎊🥳

At the ISO C++ Committee meeting in Prague 🇨🇿, hosted by [Avast](https://www.avast.com/), we completed the C++20 Committee Draft and voted to send the Draft International Standard (DIS) out for final approval and publication. Procedurally, it's possible that the DIS could be rejected, but due to our procedures and process, it's very unlikely to happen. This means that C++20 is complete, and in a few months the standard will be published.

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
- [Partially mutable lambda captures](https://wg21.link/P2034)

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

- We [clarified the meaning of `static`](https://wg21.link/p1815r2) (and unnamed namespaces) in module interfaces: such entities are now kept internal and cannot be exposed in the interface / ABI of the module. In non-modules compilations, we deprecated cases where internal-linkage entities are used from external-linkage entities. (These cases typically lead to violations of the One Definition Rule.)

- We [clarified the meaning of `inline`](https://wg21.link/p1779r3) in module interfaces: the intent is that bodies of functions that are not explicitly declared `inline` are not part of the ABI of a module, even if those function bodies appear in the module interface. In order to give module authors more control over their ABI, member functions defined in class bodies in module interfaces are no longer implicitly `inline`.

- We tweaked the context-sensitive recognition of the `module` and `import` keyword in order to avoid changing the meaning of more existing code that uses these identifiers, and to make it more straightforward for a scanning tool to recognize these declarations without full preprocessing.

- We improved backwards compatibility with unnamed enumerations in legacy header files (particularly C header files). Such unnamed enumerations will now be properly merged across header files if they're reachable in multiple different ways via imports.

- We finalized some subtle rules for concepts: a [syntax gotcha](https://wg21.link/p2092r0) in `requires` expressions was fixed, and we allowed [caching of concept values](https://wg21.link/p2104r0), which has been shown to dramatically improve performance in some cases.

- We agreed to (retroactively, via the defect report process) [treat initialization of a `bool` from a pointer as narrowing](https://wg21.link/p2104r0), improving language safety.

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
  - [Improving Engine Seeding](https://wg21.link/P2059).

&amp;nbsp;

*****

### Library Evolution Working Group (LEWG) Progress

*****

After handling the few remaining National Body comments to fix issues with C++20, LEWG focused on making general policy decisions about standard library design standards. For example, we [formally codified](https://wg21.link/P1851) the guidelines for concept names in the standard library, and clarified [SD-8](https://isocpp.org/std/standing-documents/sd-8-standard-library-compatibility), our document listing the compatibility guarantees we make to our users. Then we started looking at C++23 library proposals.

[Moved-from objects need not be valid](https://wg21.link/P2027) generated much internal discussion in the weeks leading up to the meeting as well as at the meeting itself.  While the exact solution outlined in the paper wasn’t adopted, we are tightening up the wording around algorithms on what operations are performed on objects that are temporarily put in the moved-from state during the execution of an algorithm.

The biggest C++23 news: LEWG spent an entire day with the concurrency experts of SG1 to review the [**executors** proposal](https://wg21.link/P0443) — we liked the direction! This is a huge step, which will enable **networking**, **audio**, **coroutine library support**, and more.

Other C++23 proposals reviewed include

- [a new `status_code` facility](https://wg21.link/P1208)
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

The freestanding library took a few steps forward, with some interesting proposals, including [Freestanding Language: Optional `::operator new`](https://wg21.link/P2013R0)

One of the biggest decisions was on [Low-Cost Deterministic C++ Exceptions for Embedded Systems](https://www.research.ed.ac.uk/portal/files/78829292/low_cost_deterministic_C_exceptions_for_embedded_systems.pdf) which got great reactions. We will probably hear more about it!

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
~~2018 Summer LWG Meeting~~ | ~~Chicago 🇺🇸~~ | ~~Work on wording for C++20 features.~~
~~2018 Fall EWG Modules Meeting~~ | ~~Seattle 🇺🇸~~ | ~~Design modules for C++20.~~
~~2018 Fall LEWG/SG1 Executors Meeting~~ | ~~Seattle 🇺🇸~~ | ~~Design executors for C++20.~~
~~2018 Fall Meeting~~ | ~~San Diego 🇺🇸~~ | ~~C++20 major language feature freeze.~~
~~2019 Spring Meeting~~ | ~~Kona 🇺🇸~~ | ~~C++20 feature freeze. C++20 design is feature-complete.~~
~~2019 Summer Meeting~~ | ~~Cologne 🇩🇪~~ | ~~Complete C++20 CD wording. Start C++20 CD balloting ("beta testing").~~
~~2019 Fall Meeting~~ | ~~Belfast 🇬🇧~~ | ~~C++20 CD ballot comment resolution ("bug fixes").~~
**2020 Spring Meeting** | **Prague 🇨🇿** | **C++20 CD ballot comment resolution ("bug fixes"), C++20 completed.**
2020 Summer Meeting | Varna 🇧🇬 | First meeting of C++23.
2020 Fall Meeting | New York 🇺🇸 | Design major C++23 features.
2021 Winter Meeting | Kona 🇺🇸 | Design major C++23 features.
2021 Summer Meeting | Montréal 🇨🇦 ⚜️ | Design major C++23 features.
2021 Fall Meeting | 🗺️ | C++23 major language feature freeze.
2022 Spring Meeting | Portland 🇺🇸 | C++23 feature freeze. C++23 design is feature-complete.
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
## [3][My Favorite Unknown C++20 Feature](https://www.reddit.com/r/cpp/comments/f50ag6/my_favorite_unknown_c20_feature/)
- url: https://www.reddit.com/r/cpp/comments/f50ag6/my_favorite_unknown_c20_feature/
---
### My Favorite Unknown C++20 Feature

*by Nathan Myers*

C++2020 was voted out yesterday, its final stamp of technical
approval before it goes to ISO for legal certification.
This is a big, big release, comparable to C++2011 a decade
ago, that will change, again, how we write, and think about,
our programs.
The big features--Concepts, Coroutines, Modules--will get
plenty of coverage elsewhere in the days ahead, but the 
myriad smaller features, taken together, may make as great
a difference in the code that most of us will write. 

Here I will call your attention to one small feature that, 
nonetheless, represents a change to the very foundations of
the language.
It is usually called, cryptically, the "unified comma operator",
a name that belies it importance.
It represents only the *second* time, since the ISO Standard C++
effort began in 1990, that the language grammar has ever been
simplified.
(The first was eliminating "implicit int", the fix for a wart
inherited from C, a fix that C has since also adopted.)

On its face, comma unification is simple enough.
Up through C++17, the comma has had several different meanings
in expression contexts, each with its own place in the formal
language grammar.
The most frequent use separates function actual arguments.
Similar uses appear in aggregate initializers--between elements 
being initialized in curly brackets--and in initializer-lists,
likewise.
Finally, it is a sequence operator, which may also be overloaded
by user code for, usually, expression-template libraries.

There was no motivation to unify these meanings before we got 
another important language feature: parameter packs.
Parameter packs were introduced in C++11 with very restricted
capabilities, only just enough to support integrating variadic
argument lists into the type system.
Those restrictions have been progressively relaxed in each
subsequent Standard, with new syntax greatly amplifying their
value.  We have lately dropped "parameter" from the name, because
of their expanding usefulness, and just call them "packs".

Packs have come increasingly to resemble a language-native, 
first-class version of the library feature `std::tuple&lt;&gt;`.
Comma-operator unification takes the next step toward this future.
The comma has now become, effectively, the pack construction operator.
Appearing in a function-call argument context, it constructs an
argument pack.
(Every function, now, takes just a single actual, pack, argument.)
Appearing in an aggregate initialization, the pack it constructs
conforms to the shape of the target aggregate.

C++ has a long history, and a complicated inheritance.
C++20 inherits not only from C, but also C++98, C++03, C++11,
C++14, and C++17.
Nothing new in C++ can be as simple as it might be if we could
start from a clean slate.
The comma already means things.
What it means, for observable effects of existing code, mustn't
change, if backward compatibility is to be preserved.

So, when constructing an argument pack, the values have to be
converted to match formal arguments of the selected function
overload, and the initial type and number of the pack elements
drive the choice of that overload.
In an aggregate initializer, similarly, the pack's values must
be made to conform to the aggregate they initialize.
Aggregate initialization still needs to be disambiguated from
a uniform initializer-list.
In a scalar expression, all but the last element of the pack must
be computed and then discarded--except where they need to become
arguments to an overloaded comma-operator call, instead.

When totting up new features in C++20, comma unification is often
neglected, because it it doesn't yet change what programs can be
written, or how.
But you will hear more about it as each new feature it enables
surfaces.
Comma unification is all about what will come to be.
First-class packs, in C++23 and moreso in C++26, will enable
coding in C++ to feel more like using newer scripting and 
functional languages, and enable writing more powerful libraries
with cleaner, simpler, safer interfaces.

Sometimes it's not the marquee features that have the biggest impact, in the longer term.
## [4][An Extraterrestrial Guide to C++20 Text Formatting](https://www.reddit.com/r/cpp/comments/f56u0v/an_extraterrestrial_guide_to_c20_text_formatting/)
- url: https://www.bfilipek.com/2020/02/extra-format-cpp20.html
---

## [5][C++ conference in Madrid](https://www.reddit.com/r/cpp/comments/f57egn/c_conference_in_madrid/)
- url: https://www.reddit.com/r/cpp/comments/f57egn/c_conference_in_madrid/
---
We have opened today the registrations for the C++ conference using std::cpp 2020 to be held in Madrid on April 16th.

The complete program is still to be announced but some talks are already confirmed.

And the best of all registration is FREE.

https://eventos.uc3m.es/go/usingstdcpp2020

Come to Spain for a  one day conference 100% devoted to C++.
## [6][Qt car game with genetic algorithm](https://www.reddit.com/r/cpp/comments/f53w80/qt_car_game_with_genetic_algorithm/)
- url: https://www.reddit.com/r/cpp/comments/f53w80/qt_car_game_with_genetic_algorithm/
---
I just created Qt car game, where you can compete with a genetic algorithm, or just see how it evolves. Let me know your thoughts.
https://github.com/swordey/2DCarGameWithGA
## [7][Herb Sutter's Trip Winter ISO C++ Standards Meeting (Prague) Trip Report](https://www.reddit.com/r/cpp/comments/f4p2nn/herb_sutters_trip_winter_iso_c_standards_meeting/)
- url: https://herbsutter.com/2020/02/15/trip-report-winter-iso-c-standards-meeting-prague/
---

## [8][an efficient immutable vector](https://www.reddit.com/r/cpp/comments/f4szkn/an_efficient_immutable_vector/)
- url: https://www.reddit.com/r/cpp/comments/f4szkn/an_efficient_immutable_vector/
---
I recently had an idea for an efficient immutable vector.  
It is basically a copy-on-write array behind a shared pointer. The array grows exponentially and the `push_back` operation reuses the array if possible (if there is still space left and  if it is the first `push_back` for a given array and size).
The result is a vector with an O(1) copy operation and `push_back` (also `pop_back`) operations with an amortized O(1) best case complexity. The worst case complexity for `push_back` is O(n) but only happens if you create a copy of a given vector and then push_back to both the original vector and the copy, in which case it is no worse than the standard vector since the standard vector has an O(n) complexity for copying. All other operations should have the same complexity as the standard vector: O(1) lookup, O(n) insert and erase.

You can find a prototype implementation of this data structure here: https://gist.github.com/eyelash/87a3fe849b31a37c58eca4ef398a71dc

I believe that from a (somewhat theoretical) big-O perspective this vector is strictly better than the standard vector, since there is no sequence of operations that has a worse complexity. (Is this actually correct?)

I'm surely not the first to come up with this data structure. Is there a name for it? Is there somewhere I can find more information or a more robust implementation? I'm aware of a few implementations of immutable vectors (e.g. [immer](https://github.com/arximboldi/immer)) but they all use trees internally. What is the reason for that?

**Edit**: Thinking about it some more I realize the big difference is in updating a single element which can be done in O(log n) in a tree and will be O(n) for an array-based immutable vector.
## [9][Choose optimization in C++? (Preferally clang)](https://www.reddit.com/r/cpp/comments/f50etd/choose_optimization_in_c_preferally_clang/)
- url: https://www.reddit.com/r/cpp/comments/f50etd/choose_optimization_in_c_preferally_clang/
---
-Edit- Learn programming linked me to a stackoverflow answer that gave me enough information to figure it out https://www.reddit.com/r/learnprogramming/comments/f50dvb/choose_optimization_in_c/ https://stackoverflow.com/questions/15548023/clang-optimization-levels Basically I need -inline to get all the data into a single function then -sroa to reduce the struct into allocs which eliminates the struct. So it appears that optimizers can get rid of structs outright but may not reduce the members in it (at least I don't see any passes for that)

I  want to double check if an optimization works in a very specific situation. The optimization doesn't happen if I use -O1 but the compiler optimizes out pretty much the  entire app if I do -O2 (so I can't completely confirm if it does a thing without another thing). Is there a way to specify which passes I want? I  see a list of passes here but I have no idea how to specify it [https://llvm.org/docs/Passes.html](https://llvm.org/docs/Passes.html)

I prefer the solution to be in clang

-Edit- What I'd like to check is if the optimizer will reduce the size of the struct. If I copy b.c into a.c the optimizer will remove the struct completely. If they're in seperate files the struct is the full size. I tried to use undefined behavior to see if it might trigger but clang knows I'm passing the struct and leaves it intact

----


    //a.c
    #include &lt;stdio.h&gt;

    struct Dummy
    {
        int a, b, c, d, e, f;
    };

    static void test3(struct Dummy*p)
    {
        printf("%d %d\n", p-&gt;e, sizeof(struct Dummy));
    }

    //void test2(void *p, void (*f)(void *));
    void test2(long int z, void (*f)(void *));
    static void test(int c)
    {
        struct Dummy z;
        z.e = c*2;
        test2((long int)&amp;z, test3);
    }

    int main(int argc, char *argv[])
    {
        test(argc);
        return 0;
    }

    //b.c
    void test2(long int z, void (*f)(void*))
    {
        f((void*)z);
    }
## [10][One problem that C++ surprisingly doesn't seem to have](https://www.reddit.com/r/cpp/comments/f4x7fc/one_problem_that_c_surprisingly_doesnt_seem_to/)
- url: https://www.reddit.com/r/cpp/comments/f4x7fc/one_problem_that_c_surprisingly_doesnt_seem_to/
---
For all of its complexity, bad design decisions, and interactions between features that weren't explicitly considered, it all somehow *works*. The compilers and standard library implementations aren't riddled with bugs. Everything works as designed, even if that design sometimes has flaws. The compilers generate correct code and the standard library functions don't crash.

How do they do it?
## [11][Improving C++ compile time: precompiling generic code](https://www.reddit.com/r/cpp/comments/f4pt34/improving_c_compile_time_precompiling_generic_code/)
- url: https://www.reddit.com/r/cpp/comments/f4pt34/improving_c_compile_time_precompiling_generic_code/
---
I appologize if this is not well thought out/has already been proposed and for poor english.

#Precompiling generic code


With C++20 concepts, and with upcoming reflections/metaclasses, C++ is encouraging even more generic programming. 
However, generic programming is already bottlenecked by compile time. Adding more feature to generic programming
will only make it worse as more people uses it (and convenient features might encourage costly operations for the compiler). 

The main reason for that is that you cannot precompile generic code, and I think we can improve that.

##Templates as a compiler function:

A templated function
 
    f&lt;type A,type B&gt;(A a, B b) -&gt; C

can be reinterpreted as 

    f: &lt;type A, type B&gt; -&gt; fAB = ( (A a,B b) -&gt; C )

where fAB is non templated function. This was described in a paper which I cannot remember, I appologize for that.

During compilation, the function `f` will be constructued from the code, then evaluated for some types A and B, and a  function fAB will be generated by the compiler.
All this is essentialy working as an *interpreter* with all the associated costs.

Now imagine that this function `f` was hardcoded in the compiler. Given some internal representation for type A and type B, it would generate a function fAB.
This would go much faster, since this hardcoded function `f` would not require to parse any AST, and more importantly, it would be *compiled*! 
We can expect such function to perform much faster than it's interpreted counterpart.

A natural question is what should be the target language for fAB? should it be some internal representation for the compiler (llvm IR)? or assembly?
The first leaves out time consuming optimisation, but should be much easier to implement.

##Precompiling generic code: a dynamic library for the compiler
In what was said above, we assumed there was a hardcoded function f in the compiler. This is obviously not realistic.
However, such function could be loaded dynamically as a Plugin by the compiler, and the execute the function f from that plugin.
Essentially, a precompiled generic library would be targeting the compiler instead of the executable.

The previous example was demonstrated with a function, but it should also work with templated data type generation.   

  
&amp;nbsp;  

What do you all think?

Edit: added in a comment bellow:

If we look at what module brings, we have from to https://vector-of-bool.github.io/2019/01/27/modules-doa.html:


1. Tokenization caching - Because of TU isolation, tokenizing a TU can be cached when the module is later imported into another TU.
2. Parse-tree caching - Same as above. Tokenization and parsing are some of the most expensive operations in compiling C++. In my own tests, parsing can consume up to 30% of compilation time for files with a large preprocessed output.
3. Lazy re-generation - If foo imports bar, and we later make changes to the implementation of bar, we may be able to omit recompilation of foo. Only changes to the interface of bar necessitate recompilation of foo.
4. Template specialization memoization - This one is a bit more subtle and may take more work to implement, but the potential speedups are enormous. In short: A specialization of a class or function template appearing in the module interface unit can be cached and loaded from disk later.
5. Inline function codegen caching - Codegen results for inline functions (including function templates and member functions of class templates) can be cached and later re-loaded by the 6. Inline function codegen elision - extern template allows the compiler to omit performing codegen for function and class templates. This is hugely beneficial for the linker when it performs de-dup. Modules may allow the compiler to perform more extern template-style optimizations implicitly.

4 and 5 are pretty independant from what I'm proposing. 

What about 1,2 and 3? Well my proposal would essentially be seen as a possible implementation of those points! 

However, this approach has an important other benefit: using a template function/type from a library should be **almost as fast as having this function/type hard coded in the compiler!**
## [12][Declarative support in CMakeSL](https://www.reddit.com/r/cpp/comments/f4ge4k/declarative_support_in_cmakesl/)
- url: https://www.reddit.com/r/cpp/comments/f4ge4k/declarative_support_in_cmakesl/
---
Hi, I just wanted to share an article about declarative format that I'm working on in CMakeSL. Here's a teaser:

&amp;#x200B;

https://reddit.com/link/f4ge4k/video/ull8d9f3s5h41/player

And the art: [https://gist.github.com/stryku/4c69aa510711c9da6705fa4df4545515](https://gist.github.com/stryku/4c69aa510711c9da6705fa4df4545515)

Let me know what do you think about all this (:
