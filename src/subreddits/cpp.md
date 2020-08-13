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
## [2][Updates in Grisu-Exact](https://www.reddit.com/r/cpp/comments/i8pgzc/updates_in_grisuexact/)
- url: https://www.reddit.com/r/cpp/comments/i8pgzc/updates_in_grisuexact/
---
Grisu-Exact is a float-to-string conversion algorithm with

1. The round-trip guarantee,
2. The shortest representation guarantee, and
3. The correct rounding guarantee.

It is a variant of Grisu2 a classic float-to-string conversion algorithm developed by Florian Loitsch, and is largely inspired by Ryu, another float-to-string conversion algorithm developed by Ulf Adams. The main strength of Grisu-Exact over Ryu is that it performs better for short numbers.

Here are some updates:

(1) After having some code simplifications, dead branch eliminations, compression of redundant works, and other optimizations, now it got to outperform Ryu (in my benchmark) for both binary32 (\`float\`) and binary64 (\`double\`) in terms of uniformly randomly generated data.

(2) I developed a string-to-float conversion algorithm using a similar (but simpler) idea, and had a joint-test of it with Grisu-Exact for every possible \`float\`'s. It turned out that for every input data, feeding the data into Grisu-Exact and then into the reverse algorithm correctly generated the input data itself.  Hence, I'm pretty confident about the correctness of Grisu-Exact.

(3) The paper ([https://github.com/jk-jeon/Grisu-Exact/blob/master/other\_files/Grisu-Exact.pdf](https://github.com/jk-jeon/Grisu-Exact/blob/master/other_files/Grisu-Exact.pdf)) explaining the algorithm now includes the explanation on the improved min-max Euclid algorithm. Compared to the original min-max  Euclid algorithm developed by Adams, this improved algorithm produces the exact minimum and maximum not like the original one, has less preconditions, and at the same time runs much faster.

Please have a look at the repository [https://github.com/jk-jeon/Grisu-Exact](https://github.com/jk-jeon/Grisu-Exact) if you are interested!
## [3][std::atomic_ref is awesome](https://www.reddit.com/r/cpp/comments/i8ckxr/stdatomic_ref_is_awesome/)
- url: https://www.reddit.com/r/cpp/comments/i8ckxr/stdatomic_ref_is_awesome/
---
I only recently discovered that gcc 10 supports std::atomic_ref and I’ve been playing around with it. I have to say, it might be in my top 5 favorite C++20 features. But it begs the question: before std::atomic_ref, was there a way to atomically modify a variable of non-atomic type in-place? I don’t really do lock-free programming in my every-day life, so I haven’t really ever had to think about it, but I’m curious. Also Jason Turner if you’re reading this please consider std::atomic_ref for a C++ Weekly episode, maybe there are other people who also slept on this one.
## [4][What does "public class LSFR" in Java translates to cpp?](https://www.reddit.com/r/cpp/comments/i8zs1p/what_does_public_class_lsfr_in_java_translates_to/)
- url: https://www.reddit.com/r/cpp/comments/i8zs1p/what_does_public_class_lsfr_in_java_translates_to/
---

## [5][No more plain old data](https://www.reddit.com/r/cpp/comments/i8a5xv/no_more_plain_old_data/)
- url: https://mariusbancila.ro/blog/2020/08/10/no-more-plain-old-data/
---

## [6][Closer to the Edge: Testing Compilers More Thoroughly by Being Less Conservative About Undefined Behaviour](https://www.reddit.com/r/cpp/comments/i8hl4x/closer_to_the_edge_testing_compilers_more/)
- url: https://srg.doc.ic.ac.uk/files/papers/csmithedge-ase-nier-20.pdf
---

## [7][Feasibility of non-ABI breaking, but syntax breaking proposals](https://www.reddit.com/r/cpp/comments/i8hypo/feasibility_of_nonabi_breaking_but_syntax/)
- url: https://www.reddit.com/r/cpp/comments/i8hypo/feasibility_of_nonabi_breaking_but_syntax/
---
Have there ever been any proposals that suggest breaking syntactical compatibility, but have no effect on the ABI? If not, would it possible for something like the following rule to pass in some future standard (note: I do realize `final` does not imply override)?

&gt;All overriding virtual methods must have a `final` or `override` specifier.

Basically this would mean that all currently working projects would not only have to modify their own code, but also all external dependency header files as part of some pre-build step. However, if all virtual methods just had the `override` specifier slapped on to them, it would be still compatible ABI-wise (I believe), even though the textual representation of the code would change. It would not impact the already compiled objects/libraries, nor would it necessarily force the 3rd party to release a new version of the artifacts (a frequent argument againts ABI breaks). This would mean some work would have to be done when upgrading, but it would be a one time thing on each upgraded project (i.e. creation of a pre-build script for header file patching).

I realize that there is already some work in progress regarding epochs, but that's a much more ambitious idea. I'm suggesting a blanket removal of some language features (in the above example, making override optional) without any file/module/TU-based discriminator mechanism.

Another thing this could be applicable to is the possibility of disallowing non-function-pointer and non-pointer-to-array variable declarations with parentheses, which was the source of plenty of discussion twice in the past month ([This comment thread](https://www.reddit.com/r/cpp/comments/ht93d5/a_lib_for_adding_a_stacktrace_to_every_c/fyg2cfb?utm_source=share&amp;utm_medium=web2x) and [This post](https://www.reddit.com/r/cpp/comments/i7v7yl/why_stdunique_lock_still_isnt_nodiscard/)) and just creates a gateway for nasty bugs only because of (basically) forced backwards compatibility.
## [8][recpp, a neither user-friendly nor rocket science tool to generate c++ pseudo code annotated with c++ guideline references](https://www.reddit.com/r/cpp/comments/i8n2c9/recpp_a_neither_userfriendly_nor_rocket_science/)
- url: https://github.com/kenavolic/recpp
---

## [9][toml++ v2.1.0 released](https://www.reddit.com/r/cpp/comments/i7za0j/toml_v210_released/)
- url: https://marzer.github.io/tomlplusplus/
---

## [10][Why std::unique_lock still isn't [[nodiscard]]?](https://www.reddit.com/r/cpp/comments/i7v7yl/why_stdunique_lock_still_isnt_nodiscard/)
- url: https://www.reddit.com/r/cpp/comments/i7v7yl/why_stdunique_lock_still_isnt_nodiscard/
---
The following code: [https://godbolt.org/z/4ovWhv](https://godbolt.org/z/4ovWhv)

Is most likely a bug, if you find it somewhere in the code base.

Now, [p1771r1](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1771r1.pdf) introduced a  `[[nodiscard]]` for constructors. The change was  applied retroactively to previously published C++ standards. 

Shouldn't both unique\_lock and lock\_guard, as well as other RAII wrappers, be marked `[[nodiscard]]`?  

Is there an ongoing proposal to do that, and is there a chance for it to be retroactively applied? It seems like a bug that's way too easy to make, I did it myself a few years ago and later seen it in code reviews of other people. I believe it needs attention.
## [11][AddressSanitizer for Windows: x64 and Debug Build Support | C++ Team Blog](https://www.reddit.com/r/cpp/comments/i7iydj/addresssanitizer_for_windows_x64_and_debug_build/)
- url: https://devblogs.microsoft.com/cppblog/asan-for-windows-x64-and-debug-build-support/
---

