# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (15/2020)!](https://www.reddit.com/r/rust/comments/fw2hd8/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fw2hd8/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/frfduy/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (15/2020)?](https://www.reddit.com/r/rust/comments/fw2i84/whats_everyone_working_on_this_week_152020/)
- url: https://www.reddit.com/r/rust/comments/fw2i84/whats_everyone_working_on_this_week_152020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-15-2020/40539?u=llogiq)!
## [3][Most commonly ignored clippy lints](https://www.reddit.com/r/rust/comments/fwev22/most_commonly_ignored_clippy_lints/)
- url: https://github.com/rust-lang/rust-clippy/issues/5418
---

## [4][Futures Explained in 200 Lines of Rust](https://www.reddit.com/r/rust/comments/fwiet3/futures_explained_in_200_lines_of_rust/)
- url: https://cfsamson.github.io/books-futures-explained/
---

## [5][Ask Rust Experts - a live Q&amp;A series for your questions about Rust](https://www.reddit.com/r/rust/comments/fwj27s/ask_rust_experts_a_live_qa_series_for_your/)
- url: https://rust-experts.com/#askrustexperts
---

## [6][Security advisories for March 2020: bitvec, hyper and others](https://www.reddit.com/r/rust/comments/fw3vyg/security_advisories_for_march_2020_bitvec_hyper/)
- url: https://www.reddit.com/r/rust/comments/fw3vyg/security_advisories_for_march_2020_bitvec_hyper/
---
[RustSec](https://rustsec.org) is a community database of security advisories filed against crates published to crates.io. It is maintained by the [Rust Secure Code Working Group](https://github.com/rust-secure-code/wg).

These security bugs have been identified and corrected in Rust crates in March 2020:

 * [bitvec: use-after or double free of allocated memory](https://rustsec.org/advisories/RUSTSEC-2020-0007.html)
 * [hyper: Flaw in hyper allows request smuggling by sending a body in GET requests](https://rustsec.org/advisories/RUSTSEC-2020-0008.html)
 * [cbox: CBox API allows to de-reference raw pointers without `unsafe` code](https://rustsec.org/advisories/RUSTSEC-2020-0005.html)
 * [bumpalo: Flaw in `realloc` allows reading unknown memory](https://rustsec.org/advisories/RUSTSEC-2020-0006.html)

You can use [cargo-audit](https://github.com/RustSec/cargo-audit) to identify whether your code depends on vulnerable versions of these crates and upgrade. [A GitHub action](https://github.com/actions-rs/audit-check) that files bugs if your code depends on vulnerable crates is also available.

If you have found a vulnerability in your code, please [report it to RustSec](https://github.com/RustSec/advisory-db) so that your users can identify that they're running a vulnerable version and upgrade.

If you discover a bug in unsafe code and aren't sure if it has security implications, please get in touch and we'll work with you to figure out if could have security implications.
## [7][Congratulations](https://www.reddit.com/r/rust/comments/fvxghs/congratulations/)
- url: https://www.reddit.com/r/rust/comments/fvxghs/congratulations/
---
Hello rustaceans,

I started learning Go about one year ago. I started my journey in computer science/security from C so going back (from Python and JS) to a compiled, strongly typed language made me feel somehow home. Two things I appreciated a lot are the ecosystem and the Go community. The tools are well done and the community is very welcoming.

In the meantime I discovered this beautiful language called Rust. I know the public opinion is that Go is competing with Rust, but for me it served as a bridge. Reading the rust book I caught myself saying under my breath "No freaking way..." more than once. Everything just clicked and just made sense. Besides the content I was impressed by the initiative of writing such a wonderful resource to help virtually anyone transition to Rust. Cargo is so incredibly polished that it was a pleasure even to just watch source code compile. Now to the main point. I heard plenty of good stuff about Rust's community and they all proved to be true. The thing that pushed me to post this is seeing an anon's post here a few days ago, being quite on the offensive about Rust being unsafe. More than once he proved a lot of bias which kind of drove me insane just by reading it. I was shocked to see how rustaceans responded with information, sources and explanations pointing out the differences in perspective. It was just incredible. This is just an example, there are more where this came from.

Go's ecosystem and community are cool, I'm not trying to deny it. But for me,  Rust is superior in every way, except maybe the breadth and maturity of some packages. But this will come with time, I'm not worried at all.

Congratulations for building such a community. I am looking forward to finish the introductory stuff in the language and be a part of it. Continue to shock me and be at least as awesome as you are
## [8][From failure to Fehler](https://www.reddit.com/r/rust/comments/fw4jsx/from_failure_to_fehler/)
- url: https://boats.gitlab.io/blog/post/failure-to-fehler/
---

## [9][rustfmt is not working from VIM](https://www.reddit.com/r/rust/comments/fwjr6k/rustfmt_is_not_working_from_vim/)
- url: https://www.reddit.com/r/rust/comments/fwjr6k/rustfmt_is_not_working_from_vim/
---
I installed rust.vim plugin via Pathogen. Running rustfmt from the filesystem is fine but when I try to run :RustFmt from within VIM i get the following error:

Error detected while processing function rustfmt#Format:

line    1:

E121: Undefined variable: v:false

E116: Invalid arguments for function &lt;SNR&gt;33\_RunRustfmt

Also I would like to run rustfmt each time I'm saving changes. This doesnt work either.

Can somebody help?
## [10][Rust analyzer weekly changelog](https://www.reddit.com/r/rust/comments/fvygkf/rust_analyzer_weekly_changelog/)
- url: https://rust-analyzer.github.io/thisweek/2020/04/06/changelog-19.html
---

## [11][Why doesn't rust allow invoking "static" methods on trait objects?](https://www.reddit.com/r/rust/comments/fwfggk/why_doesnt_rust_allow_invoking_static_methods_on/)
- url: https://www.reddit.com/r/rust/comments/fwfggk/why_doesnt_rust_allow_invoking_static_methods_on/
---
Take the following code ([playground link](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=db51eb20e8355f95578b6ceba07b904f))

(kind of a dumb example, but it's just to demonstrate the point)

    trait Test {
        fn init() -&gt; Box&lt;dyn Test&gt;;
    }

If you try to compile this, you get the following error:

    error[E0038]: the trait `Test` cannot be made into an object...
    associated function `init` has no `self` parameter

The rustc explanation (`rustc --explain E0038`) for this isn't very satisfying:

&gt;Methods that do not take a \`self\` parameter can't be called since there won't be  
&gt;  
&gt;a way to get a pointer to the method table for them  
&gt;  
&gt;...  
&gt;  
&gt;This could be called as \`&lt;Foo as Foo&gt;::foo()\`, which would not be able to pick  
&gt;  
&gt;an implementation.  
&gt;  
&gt;...  
&gt;  
&gt;Adding a \`Self: Sized\` bound to these methods will generally make this compile.

If you think about it, there's nothing stopping the compiler from embedding function pointers to "static" methods in the trait object's vtable just like it does for any other trait method, and there isn't any ambiguity about the calling convention since both the caller and the callee are aware that the method doesn't borrow 'self'.

This is especially annoying if you're trying to make up for the lack of arbitrary self types, by doing something like:

    // Do something with my custom smart pointer
    trait DoSomething {
        // This method *needs* to consume "this" (self), so can't use a borrow
        fn do_something(this: SmartPointer&lt;dyn DoSomething&gt;, arg: f32) -&gt; String;
    }
    
    // Totally making up the syntax here
    let sp: SmartPointer&lt;dyn DoSomething&gt; = ...;
    let do_something = sp.do_something;
    (do_something)(sp);

Now, you could work around this by doing something like:

    trait DoSomething {
        fn get_do_something(&amp;self) -&gt; fn(SmartPointer&lt;dyn DoSomething&gt;, f32) -&gt; String;
    }

But then you're paying the cost of two dynamic function calls when you really only need one, which is unfortunate.

This will be kind of a moot point once arbitrary self types on trait objects get stabilized since the absence of that is the primary use case I'm running into for this, but it's unclear when that will happen.

I can imagine two arguments for why this was initially decided against:

* Don't want to bloat memory with vtable members that are unlikely to get called
   * This is kind of irrelevant now since all existing code uses the `where Self: Sized` bound on static methods, so enabling it wouldn't have any immediate effect.
   * It's also outweighed by the way `dyn(TraitA + TraitB)` is represented (by constructing a brand new vtable with the combination of TraitA's and TraitB's methods).
* If you're designing a polymorphic trait, you need to account for the fact that implementations may need to borrow 'self'
   * If you had to implement a trait from a library that didn't allow you to do that, you'd be hosed.
   * I think most API designers probably realize that, so this is kind of a weak argument.

This post is kind of a ramble, but I'm sitting here slightly frustrated since I need one of these two features  and neither are currently allowed. This one seems like the lower hanging fruit, so does anyone know why this was decided against?
## [12][no_error - an error library for no_std](https://www.reddit.com/r/rust/comments/fwenhz/no_error_an_error_library_for_no_std/)
- url: https://github.com/richardanaya/no_error
---

