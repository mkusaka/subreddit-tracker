# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (33/2020)!](https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.45]](https://www.reddit.com/r/rust/comments/htnv19/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/htnv19/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://old.reddit.com/r/rust/comments/gz8ut5/official_rrust_whos_hiring_thread_for_jobseekers/).

This thread will be periodically stickied to the top of /r/rust for improved visibility. The thread will be refreshed and posted anew when the next version of Rust releases in six weeks.

Please adhere to the following rules when posting (adapted from /r/cpp's jobs thread).
# Rules for individuals:

 * Don't create top-level comments; those are for employers.
 * Feel free to reply to top-level comments with on-topic questions.
 * I will create a stickied top-level comment for individuals looking for work.
 * I will create an additional top-level comment for meta discussion.

# Rules for employers:

 * To find individuals seeking work, see the replies to the stickied top-level comment; you will need to click the "more comments" link at the bottom of the top-level comment in order to make these replies visible.
 * To make a top-level comment you must be hiring directly; no third-party recruiters.
 * One top-level comment per employer. If you have multiple job openings, please consolidate their descriptions or mention them in replies to your own top-level comment.
 * Proofread your comment after posting it and edit it if necessary to correct mistakes.
 * Please base your comment on the following template:

COMPANY: *[Company name; please link to your company's website or careers page.]*

TYPE: *[Full time, part time, internship, contract, etc.]*

DESCRIPTION: *[What does your company do, and what are you using Rust for? How much experience are you seeking and what seniority levels are you hiring for? The more details the better.]*

LOCATION: *[Where are your office or offices located? If your workplace language isn't English, please specify it.]*

REMOTE: *[Do you offer the option of working remotely? If so, do you require employees to live in certain areas or time zones?]*

VISA: *[Does your company sponsor visas?]*

CONTACT: *[How can someone get in touch with you?]*
## [3][generic-std: streaming iterators and other HKT-powered traits in stable Rust](https://www.reddit.com/r/rust/comments/if0a3u/genericstd_streaming_iterators_and_other/)
- url: https://www.reddit.com/r/rust/comments/if0a3u/genericstd_streaming_iterators_and_other/
---
tl;dr [https://crates.io/crates/generic-std](https://crates.io/crates/generic-std)

Not rarely I find myself trying to generalize my code, hitting some issues, searching for answers online, and then realizing the type system is not powerful enough to express what I need. And these blockers are always centered around the fact that Rust doesn't support higher-kinded types.

For those that don't know, higher-kinded types (HKTs) are *type constructors*. Rust has a limited implementation of HKTs in the form of generics. For example, `Vec` is a type constructor that takes a single (type) argument, while `Vec&lt;T&gt;` is a type resulting from the application of `T` to the `Vec` type constructor. The crucial bit that is missing is the ability to pass type constructors around like you can do with normal types.

One example where this would be useful is in a data structure that needs a reference-counted box. You may want to give the user a choice between `Rc`, `Arc`, or maybe even a non-std implementation. You could define, say, an [`Rcb&lt;T&gt;`](https://docs.rs/generic-std/0.1.0/generic_std/trait.Rcb.html) trait and let the user choose which implementations to use. This is fine for simple cases, but it can get unwieldy when you need reference counting for values of multiple types:

    struct Example&lt;RcString, RcStructA, RcStructB&gt;
    where
      RcString: Rcb&lt;String&gt;,
      RcStructA: Rcb&lt;StructA&gt;,
      RcStructB: Rcb&lt;StructB&gt;, {...}
    
    let example = Example::&lt;Arc&lt;String&gt;, Arc&lt;StructA&gt;, Arc&lt;StructB&gt;&gt;::new();

In an ideal world where Rust has native support for HKT this case might instead be expressible like this:

    struct Example&lt;Rc&gt;
    where
      for&lt;T in [String, StructA, StructB]&gt; Rc&lt;T&gt;: Rcb&lt;T&gt;, {...}
    
    let example = Example::&lt;Arc&gt;::new();

Here `Arc` is usable even if we don't specify its type argument. (spoiler: [`H1Arc`](https://docs.rs/generic-std/0.1.0/generic_std/sync/struct.H1Arc.html))

While in this example HKTs would be a nice to have, there are some cases that are (supposedly) inexpressible in stable Rust. The most famous example is probably the streaming iterator:

    trait StreamingIterator {
      type Item&lt;'a&gt;;
      fn next&lt;'a&gt;(&amp;'a mut self) -&gt; Option&lt;Self::Item&lt;'a&gt;&gt;;
    }

This is a trait for iterators that may return references to parts of itself. This trait as it is relies on a proposed feature called [associated type constructors](https://github.com/rust-lang/rfcs/blob/master/text/1598-generic_associated_types.md) (ATCs). In fact, the ATC RFC itself states that this is impossible to express with the current Rust language (spoiler: [`StreamingIterator`](https://docs.rs/generic-std/0.1.0/generic_std/trait.StreamingIterator.html)).

Lucky for us, Rust's type system is powerful enough to be turing-complete, which means we can do any computation with types, including emulating type constructors. This is a bit awkward because each value in a type function is a separate type that needs to be declared, and type functions are declared using traits and associated types, but it works.

I've published the [generic-std](https://crates.io/crates/generic-std) crate as a proof of concept. I've also included three [unit tests](https://github.com/Ereski/generic-std/blob/master/src/tests.rs) that show how the streaming iterator and reference counting examples can be handled. The current [docs](https://docs.rs/generic-std/0.1.0/generic_std/) should be enough to give a basic idea of how things are set up, and the way HKT is emulated is surprisingly simple.
## [4][More CPP](https://www.reddit.com/r/rust/comments/iewg37/more_cpp/)
- url: https://www.reddit.com/r/rust/comments/iewg37/more_cpp/
---
After having read the recent discussions about cxx and autocxx crates, I thought I'd bring to your attention another one,  which I consider to be a forgotten gem of the Rust-C++ ecosystem - the [cpp](https://crates.io/crates/cpp) crate.

What distinguishes `cpp` from other similar crates, is that it does not attempt to automatically generate Rust bindings;  instead, it allows one to inline c++ right into your Rust code.   Here's what the usage of `cpp` looks like  (borrowed from cpp's documentation):

    use cpp::cpp;

    cpp!{{
        #include &lt;iostream&gt;
    }}

    fn main() {
        let name = std::ffi::CString::new("World").unwrap();
        let name_ptr = name.as_ptr();
        let r = cpp!(unsafe [name_ptr as "const char *"] -&gt; u32 as "int32_t" {
            std::cout &lt;&lt; "Hello, " &lt;&lt; name_ptr &lt;&lt; std::endl;
            return 42;
        });
        assert_eq!(r, 42)
    }

`cpp` can invoke arbitrary c++ functions and methods, even templated ones.  Solutions like `bindgen` cannot use c++ functions defined in headers - that code simply hasn't been generated yet, so there's nothing to invoke in the library.  It uses an actual c++ compiler to generate glue code, so no problems with symbol mangling or matching the c++ ABI.

As we've just learnt in adjacent discussions, it is not possible to automatically generate sound Rust bindings for arbitrary c++ code, so you are going to have to create safe wrappers anyway.   And this is where the ability granted by `cpp` to call arbitrary c++ code with minimal fuss comes in extremely handy.

No good thing is without some downsides, of course:   

* `cpp` does not have built-in conversions even for primitive types, you'll have to annotate all parameter and return types manually,
* because it relies on `syn` crate to parse Rust source and extract C++ closures, the `cpp!` macro cannot be wrapped other macros, which is too bad, because in many cases macros would have helped to cut down on boilerplate needed due to #1.

Still, this doesn't compare too badly to `cxx`, where you also need to write out the types of all function parameters at least once.
(Actually, [the very first version of cpp](https://github.com/mystor/rust-cpp/tree/legacy_rustc_plugin) did not require users to provide types.  Unfortunately this functionality relied on a compiler plugin and would have never been able to work on stable, so this approach was abandoned.  I also wonder if we could get this convenience back eventually - when const generics become advanced enough, there has got to be a way to apply them to compile-time code generation.   Well, maybe with a bit of help from the compiler :-).
## [5][Why You Should Use Rust in 2020](https://www.reddit.com/r/rust/comments/iewznw/why_you_should_use_rust_in_2020/)
- url: https://serokell.io/blog/rust-guide
---

## [6][As above, so below: Using Rust generics to develop two bare-metal flash drivers](https://www.reddit.com/r/rust/comments/if21da/as_above_so_below_using_rust_generics_to_develop/)
- url: https://www.ecorax.net/as-above-so-below-1/
---

## [7][jlrs 0.6: Support for Julia 1.5, many new builtin types, massively extended support for accessing data from Rust, automatically generate valid Rust structs from Julia structs.](https://www.reddit.com/r/rust/comments/if1oe7/jlrs_06_support_for_julia_15_many_new_builtin/)
- url: https://www.reddit.com/r/rust/comments/if1oe7/jlrs_06_support_for_julia_15_many_new_builtin/
---
A few weeks ago Julia 1.5 was released. In this version the layout of structs was changed, before any struct with pointer fields couldn't be stored inline but had to be hidden behind a pointer (a `Value` in jlrs). Now, immutable structs with pointer fields will be inlined. One of the results of this change is that arrays that contain immutable structs with pointer fields (eg a stacktrace) is no longer stored as an array of `Value`s but as an inline array, in order to access its contents purely from Rust you have to provide a struct with the correct layout.

In order to make this possible the derive macros have received a massive overhaul in jlrs 0.6. First of all, `JuliaTuple` is gone and has been replaced with many generic tuple types. Second, `JuliaStruct` supports much more than bits types without type parameters and can now handle pointer fields, type parameters, and bits unions. Finally, you don't even need to figure out the proper layout yourself, the [JlrsReflect.jl](https://github.com/Taaitaaiger/JlrsReflect.jl) package can be used to generate valid bindings automatically. This package should be available through `Pkg` soon.

Other changes include the availability of all datatypes defined in the Julia C API like `TypeName` and `Expr`, a raw string `JuliaString`, and the ability to set module globals and constants from Rust.

jlrs 0.6 is only compatible with Julia 1.5, earlier versions will not work.

[Github](https://github.com/Taaitaaiger/jlrs)

[Docs.rs](https://docs.rs/jlrs/0.6.0/jlrs/index.html)

[Crates.io](https://crates.io/crates/jlrs)
## [8][[Rust Analyzer] cross platform check](https://www.reddit.com/r/rust/comments/iezpz3/rust_analyzer_cross_platform_check/)
- url: https://www.reddit.com/r/rust/comments/iezpz3/rust_analyzer_cross_platform_check/
---
I'm trying to develop a windows application from Linux, found the crate `native-windows-gui` which is good and can be compiled from linux.

But since I'm in Linux Rust Analyzer is using `cargo check` in the background to get the errors and other stuff. When running `cargo check` without changing the target it would spit a lot of error which is understandable as it only works for windows.

Is there any way to mark a specific module or package to use a specific target?
## [9][The CXX Debate](https://www.reddit.com/r/rust/comments/ielvxu/the_cxx_debate/)
- url: https://steveklabnik.com/blog/the-cxx-debate
---

## [10][The problem of safe FFI bindings in Rust](https://www.reddit.com/r/rust/comments/iev7za/the_problem_of_safe_ffi_bindings_in_rust/)
- url: https://www.abubalay.com/blog/2020/08/22/safe-bindings-in-rust
---

## [11][standback 0.2.10 has been released — use new stdlib APIs in old compilers](https://www.reddit.com/r/rust/comments/ieycn0/standback_0210_has_been_released_use_new_stdlib/)
- url: https://www.reddit.com/r/rust/comments/ieycn0/standback_0210_has_been_released_use_new_stdlib/
---
Some of you may remember [back in March](https://users.rust-lang.org/t/introducing-standback-portions-of-the-standard-library-backported-to-support-older-compilers/39136) (oh, how long ago that feels!) when I announced standback. The TL;DR of the whole thing is that a significant portion of, though not all, APIs stabilized in the standard library are fully capable of being implemented in userland. So…why not‽

I have just released standback 0.2.10, which includes support for Rust 1.45 and 1.46, the latter of which is due to be released this coming Thursday. I typically aim to have support for each release as they come out, but I was preoccupied with my capstone this summer.

I have been using standback as a direct dependency of time 0.2 since its release without issue (and over 600,000 downloads). If you are the author of a library and desire to maintain _both_ your MSRV and using new APIs, this crate is for you. If the end user is using a newer compiler, they pay extremely little compilation cost (near zero), as everything is gated behind automatic version detection.
## [12][google/autocxx - calling C++ from Rust in a heavily automated, but safe, fashion](https://www.reddit.com/r/rust/comments/iefeum/googleautocxx_calling_c_from_rust_in_a_heavily/)
- url: https://github.com/google/autocxx
---

