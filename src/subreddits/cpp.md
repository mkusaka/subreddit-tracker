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
## [2][CppCon Program is available NOW](https://www.reddit.com/r/cpp/comments/il05m8/cppcon_program_is_available_now/)
- url: https://www.reddit.com/r/cpp/comments/il05m8/cppcon_program_is_available_now/
---
The main program is now available. Please follow this [link](https://cppcon.org/2020-program-announcement/) for general information.

Main Program for CppCon 2020: [https://cppcon.org/program2020/](https://cppcon.org/program2020/)

Information about the CppCon 2020 tracks:

1. Embedded Track: [https://cppcon.org/embedded/](https://cppcon.org/embedded/)
2. Back to Basics Track: [https://cppcon.org/b2b/](https://cppcon.org/b2b/)
3. One surprise track to be defined.
## [3][in_template [c++20/GCC] constexpr bool function returning true if called in a template context](https://www.reddit.com/r/cpp/comments/il3ppw/in_template_c20gcc_constexpr_bool_function/)
- url: https://godbolt.org/z/jxjn55
---

## [4][eCAL 5.7.1 released](https://www.reddit.com/r/cpp/comments/il4zvk/ecal_571_released/)
- url: https://www.reddit.com/r/cpp/comments/il4zvk/ecal_571_released/
---
New version with lots of bug fixes and some nice new recording features released. Your communication middleware is to slow, your DDS system to complex ?

[eCALize it](https://github.com/continental/ecal) ;-)
## [5][On Modern Hardware the Min-Max Heap beats a Binary Heap](https://www.reddit.com/r/cpp/comments/ikpupr/on_modern_hardware_the_minmax_heap_beats_a_binary/)
- url: https://probablydance.com/2020/08/31/on-modern-hardware-the-min-max-heap-beats-a-binary-heap/
---

## [6][Two musings on the design of compiler warnings](https://www.reddit.com/r/cpp/comments/il3aix/two_musings_on_the_design_of_compiler_warnings/)
- url: https://quuxplusone.github.io/blog/2020/09/02/wparentheses/
---

## [7][=nil; Crypto3. Concept-based pure C++ cryptography library. From cryptography developers for cryptography developers. Headed to Boost and ISO/C++.](https://www.reddit.com/r/cpp/comments/ikxgs5/nil_crypto3_conceptbased_pure_c_cryptography/)
- url: https://www.reddit.com/r/cpp/comments/ikxgs5/nil_crypto3_conceptbased_pure_c_cryptography/
---
I always wondered, why C++ standard lacks of cryptographic primitives being standardized by NIST long time ago. Even of those, which became widely known and literally immortal. So here is the answer I figured: generic programming and cryptography are quite hard to marry. The first point is about being as abstract as it could be, second point is about being as particular as it could be.

But I also always wanted to encrypt/decrypt like this:

`
std::vector&lt;std::uint32_t&gt; in = {\x0123, \0x23145}, out;  
std::encrypt&lt;std::cipher::aes&gt;(in, out);
std::decrypt&lt;std::cipher::aes&gt;(out, in);
`

Or sign/verify like this:

`
std::string in = "Message", sig, key = "12345";
std::sign&lt;std::signature::ecdsa&gt;(in, key, sig);
bool valid = std::verify&lt;std::signature::ecdsa&gt;(sig, key);
`

I also always wanted to have an extensible cryptography suite. So I could easily embed my new super fancy hash in there simply complying to some concept, constructing the scheme from ready to use components (e.g. Merkle-Damgard or Sponge  constructions) with simply substituting classes to templates. So, you can do that now.

Do you want to know what I also always wondered? Why don't we have a C++ collection of classic-notion cryptography (ciphers, hashes, kdfs etc) along with modern constructions (ZK-cryptography, Verifiable Delay Functions, threshold cryptography)? So, now we have it. Work in progress, but still better than nothing.

So here is the motivation. And here is the implementation: https://github.com/nilfoundation/crypto3.git. Check it out. You can do encryption/hashing/encoding like this:

`
using namespace nil::crypto3;
std::string in = "Message", res;
encrypt&lt;ciphers::rijndael&lt;128, 256&gt;&gt;(in.begin(), in.end(), res);
`

Or like this:

`
using namespace nil::crypto3;
std::array&lt;std::uint8_t, 256&gt; in, out;
in.fill(1);
hash&lt;hashes::sha3&gt;(in, out);
`

And yes, there is a Boost-ified version in here: https://github.com/nilfoundaton/boost-crypto3.git. And the discussion in Boost mailing list in here: https://lists.boost.org/Archives/boost//2020/09/249672.php.

I would really like to keep the introduction post short and simple, so I'm going to keep taking a look from time to time at this thread to reply to any further questions.
## [8][Accelerating Map Multi-Lookup with Coroutines](https://www.reddit.com/r/cpp/comments/il5m9i/accelerating_map_multilookup_with_coroutines/)
- url: https://turingcompl33t.github.io/Coroutine-Map/
---

## [9][spdlog 1.8.0 released](https://www.reddit.com/r/cpp/comments/ikk819/spdlog_180_released/)
- url: https://github.com/gabime/spdlog/releases/tag/v1.8.0
---

## [10][Should I use C++ exceptions?](https://www.reddit.com/r/cpp/comments/ikv9kv/should_i_use_c_exceptions/)
- url: https://www.reddit.com/r/cpp/comments/ikv9kv/should_i_use_c_exceptions/
---
Many programming guides warns against using exceptions in C++. Exceptions are allegedly very slow and dangerous. But most modern languages use exceptions so why is C++ different? Does exceptions slow down the code only if they actually occur or does exeptions make all the code slow even if no exceptions happen to occur at runtime?

I'm aware there are many pitfalls but judicious use of RAII and obeying the golden rule of C++ maybe mitigates most of them. Also breaking out of deeply nested structures is much easier with exceptions and makes for more readable code.  Java programmers use exceptions all the time.
## [11][Use CMake to sync files in folders reading configuration from INI file.](https://www.reddit.com/r/cpp/comments/il0jkr/use_cmake_to_sync_files_in_folders_reading/)
- url: https://www.reddit.com/r/cpp/comments/il0jkr/use_cmake_to_sync_files_in_folders_reading/
---
[Use CMake and INI files to configure folder synchronisation](http://goorep.se:1001/changelog/report/rSelect/PAGE_result.htm?alias=guest&amp;set=api&amp;query=Book+pages&amp;$$TArticleBook1.ArticleBookK=7119&amp;link=%5B%5B%229F1E006D78894848838A0970E2FF0BE9zoom%22,%22Object1%22,7119%5D,%5B%2271C91DEE3C5A4FDC8EC1114C7C18033Bzoom%22,%22TArticleBook1%22,7119%5D%5D&amp;rows=25)

CMake has some strong commands, commands that can be used not only for generating builds. I haven't seen much of this. Updating files, removing files like cleaning a project from files generated during compiles etc. CMake is very useful for this.  
Link shows a simple script to update selected files specified in INI files to target folders is found in link with description.

**is there any good samples or tips on what you can do with CMake in addition to generating builds?**

Documentation in this link differs some from normal blogs. Here is a tutorial that briefly shows how it works [Getting started with changelog and selection](https://www.youtube.com/watch?v=Dmaue86qseo).
