# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (10/2020)!](https://www.reddit.com/r/rust/comments/fc7h20/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fc7h20/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/f8ney8/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 328](https://www.reddit.com/r/rust/comments/fdzspa/this_week_in_rust_328/)
- url: https://this-week-in-rust.org/blog/2020/03/03/this-week-in-rust-328/
---

## [3][Not Another Lifetime Annotations Post](https://www.reddit.com/r/rust/comments/feb31o/not_another_lifetime_annotations_post/)
- url: https://www.reddit.com/r/rust/comments/feb31o/not_another_lifetime_annotations_post/
---
Yeah, it is. But I've spend a few days on this and I'm starting to really worry about my brain, because I'm just getting nowhere.

To be clear, it's not *lifetimes* that are confusing. They aren't. Anyone who's spent any meaningful time writing C/C++ code understands the inherent problem that Rust solves re: dangling pointers and how strict lifetimes/ownership/borrowing all play a role in the solution.

But...lifetime annotations, I simply and fully cannot wrap my head around.

Among other things (reading the book, a few articles, and reading some of Spinoza's ethics because this shit is just about as cryptic as what he's got in there so maybe he mentioned lifetime annotations), I've watched [this video](https://www.youtube.com/watch?v=QoEX-Vu-R6k), and the presenter gave me hope early on by promising to dive into annotations specifically, not just lifetimes. But...then, just, nothing. Nothing clicks, not even a nudge in the direction of a click. Here are some of my moments of pure confusion:

* At one point, he spams an `'a` lifetime parameter across a function signature. But then it compiles, and he says "these are wrong". I have no idea what criteria for correctness he's even using at this point. What I'm understanding from this is that all of the responsibility for correctness here falls to the programmer, who can fairly easily "get it wrong", but with consequences that literally no one specifies anywhere that I've seen.
* He goes on to 'correct' the lifetime annotations...but he does this with **explicit knowledge of the calling context**. He says, "hey, look at this particular call - one of the parameters here has an entirely different lifetime than the other!" and then alters the lifetimes annotations *in the function signature* to reflect that particular call's scope context. How is this possibly a thing? There's no way I can account for every possible calling context as a means of deriving the "correct" annotations, and as soon as I do that, I might have created an invalid annotation signature with respect to some other calling context.
* He then says that we're essentially "mapping inputs to outputs" - alright, that's moving in the right direction, because the problem is now framed as one of relations between parameters and outputs, not of unknowable patterns of use. But he doesn't explain how they relate to each other, and it just seems completely random to me if you ignore the outer scope.

The main source I've been using, though, is The Book. Here are a couple moments from the annotations section where I went actually wait what:


&gt;We also don’t know the concrete lifetimes of the references that will be passed in, so we can’t look at the scopes...to determine whether the reference we return will always be valid.

Ok, so that sort of contradicts what the guy in the video was saying, if they mean this to be a general rule. But then:

&gt; For example, let’s say we have a function with the parameter `first` that is a reference to an `i32` with lifetime `'a`. The function also has another parameter named `second` that is another reference to an `i32` that also has the lifetime `'a`. The lifetime annotations indicate that the references `first` and `second` must both live as long as that generic lifetime.

Now, suddenly, it *is* the programmer's responsibility yet again to understand the "outer scope". I just don't understand what business it is of the function signature what the lifetimes are of its inputs - if they live longer than the function (they should inherently do so, right?) - why does it have to have an opinion? What is this informing as far as memory safety?

&gt; The constraint we want to express in this signature is that all the references in the parameters and the return value must have the same lifetime.

This is now dictatorial against the outer scope in a way that makes no sense to me. Again, why does the function signature care about the lifetimes of its reference parameters? If we're trying to resolve confusion around a returned reference, I'm still unclear on what the responsibility of the function signature is: if the only legal thing to do is return a reference that lives longer than the function scope, then that's all that either I or the compiler could ever guarantee, and it seems like all patterns in the examples reduce to "the shortest of the input lifetimes is the longest lifetime we can guarantee the output to be", which is a hard-and-fast rule that doesn't require programmer intervention. At best we could contradict the rule if we knew the function's return value related to only one of the inputs, but...that also seems like something the compiler could infer, because that guarantee probably means there's no ambiguity. Anything beyond seems to me to be asking the programmer, again, to reach out into outer scope to contrive to find a better suggestion than that for the compiler to run with. Which...we could get wrong, again, but I haven't seen the consequences of that described anywhere.

&gt; The lifetimes might be different each time the function is called. This is why we need to annotate the lifetimes manually.

Well, yeah, Rust, that is exactly the problem that *I* have. We have a lot in common, I guess. I'm currently mulling the idea of what happens when you have some sort of struct-implemented function that takes in references that the function intends to take some number of immutable secondary references to (are these references of references? Presumably ownership rules are the same with actual references?) and distribute them to bits of internal state, but I'm seeing this problem just explode in complexity so quickly that I'm gonna not do that anymore.

That's functions, I guess, and I haven't even gotten to how confused I am about annotations in structs (why on earth would the struct care about anything other than "these references outlive me"??) I'm just trying to get a handle on one ask: how the hell do I know what the 'correct' annotations are? If they're call-context derived, I'm of the opinion that the language is simply adding too much cognitive load to the programmer to justify any attention at all, or at least that aspect of the language is and it should be avoided at all costs. I cannot track the full scope context of every possible calling point all the time forever. How do library authors even exist if that's the case?

Of course it isn't the case - people use the language, write libraries and work with lifetime annotations perfectly fine, so I'm just missing something very fundamental here. If I sound a bit frustrated, that's because I am. I've written a few thousand lines of code for a personal project and have used 0 lifetime annotations, partially because I feel like most of the potential use-cases I've encountered present much better solutions in the form of transferring ownership, but mostly because I don't get it. And I just hate the feeling that such a central facet of the language I'm using is a mystery to me - it just gives me no creative confidence, and that hurts productivity.

*edit for positivity: I am genuinely enjoying learning about Rust and using it in practice. I'm just very sensitive to my own ignorance and confusion.
## [4][I have just made my first open source contribution to rust-analyzer. You should contribute to open source too!](https://www.reddit.com/r/rust/comments/fe4bs1/i_have_just_made_my_first_open_source/)
- url: https://avishay.dev/2020/03/04/oss-contribution/
---

## [5][I released my Rust Galaxy Simulator with GPU acceleration powered by wgpu. Try it out and give it a star if you like it 🎉](https://www.reddit.com/r/rust/comments/fdxdhb/i_released_my_rust_galaxy_simulator_with_gpu/)
- url: https://github.com/timokoesters/nbodysim
---

## [6][Introducing standback: portions of the standard library backported to support older compilers](https://www.reddit.com/r/rust/comments/fe7bna/introducing_standback_portions_of_the_standard/)
- url: https://www.reddit.com/r/rust/comments/fe7bna/introducing_standback_portions_of_the_standard/
---
# Introducing [standback](https://crates.io/crates/standback)

([documentation](https://docs.rs/standback))

There are [an increasing number of stabilizations](https://github.com/rust-lang/rust/blob/master/RELEASES.md) with every release. Ever wanted to use newly-stabilized features in a crate or published binary, but decided against it because it would mean bumping the minimum supported Rust version? This is a situation many authors have to deal with, and it's not always clear _when_ to bump the MSRV, even when soundness is an issue (see `mem::uninitialized()` vs `mem::MaybeUninit`).

Standback (from "standard library" and "backport") helps you out here. It has a minimum supported Rust version of 1.31, and has _a ton_ of APIs that have been backported. Nearly all of this code comes directly from the standard library itself, so it can definitely be trusted. Oh, and it's also `#![no_std]` compatible to a large extent.

If all you're doing is using a newly stabilized method (not an import), just add `use standback::prelude::*;` to the top of your file. If what you're using is something that must be imported, replace `std` with `standback`. For example, to import the `TryFrom` trait in Rust 1.31: `use standback::convert::TryFrom;`.

Standback will automatically determine the compiler version in use, and will happily re-export an import if it has been stabilized. This ensures that you can confidently use a type even if it has been stabilized — and do so at nearly zero cost!

Hopefully this proves to be a useful project to the Rust community at large. Surprisingly, it actually wasn't too much effort to put this together. I intend on keeping this up-to-date with each Rust release, so the functionality will continue to improve at (nearly) the same speed as `std` :slightly_smiling_face:. After some looking things over to ensure nothing is terribly buggy (or plain doesn't compile), my intent is to maintain a 0.X release, where X is the current compiler version (so 0.41 right now).

---

NB: If you're actively compiling on an older compiler (in CI, for example), you'll likely want to `#![allow(unstable_name_collisions)]` on those older versions. That is the whole point of this crate, after all.
## [7][WPC - Wallpaper Changer written in Rust](https://www.reddit.com/r/rust/comments/feadnx/wpc_wallpaper_changer_written_in_rust/)
- url: https://github.com/jkotra/wpc
---

## [8][Looking for a number of testing tools](https://www.reddit.com/r/rust/comments/fea9ry/looking_for_a_number_of_testing_tools/)
- url: https://www.reddit.com/r/rust/comments/fea9ry/looking_for_a_number_of_testing_tools/
---
I have a course in university called software testing. Normally, we have to use Java in this course, but I asked my professor If I can use Rust instead. Now I need to find equivalent technologies to do the same testing and I think this is the best place to consolidate such a list.

These are purposes for tools Im looking for:

* Random testing (JAVA tool: Randoop) (Rust tool: ?)
   * [https://github.com/altsysrq/proptest](https://github.com/altsysrq/proptest) (Thanks @etareduce)
* Input Space Partitioning (Combinatorial Testing) (Java tool: ACTS) (Rust tool: ?)
* Graph Coverage (Java tool: EclEmma) (Rust tool: ?)
* Logic Coverage (Java tool: EclEmma) (Rust tool: ?)
   * [https://github.com/xd009642/tarpaulin](https://github.com/xd009642/tarpaulin) (?)
* Syntax Coverage / Mutation Testing (Java tool: MuJava, Major) (Rust tool: ?)
   * [https://github.com/Geal/mutant](https://github.com/Geal/mutant)
   * [https://github.com/llogiq/mutagen](https://github.com/llogiq/mutagen)
   * [https://llogiq.github.io/2018/02/14/mutagen.html](https://llogiq.github.io/2018/02/14/mutagen.html)

&amp;#x200B;

I could not find much info for them. I mean I can always code tests to do this kind of testing myself, but an automated way, using these tools would be nice.

Cheers!
## [9][Found this to be extremely interesting](https://www.reddit.com/r/rust/comments/fdzgc8/found_this_to_be_extremely_interesting/)
- url: /r/FlutterDev/comments/fdgrdh/finally_running_rust_natively_on_a_flutter_plugin/
---

## [10][Introducing Yew Form, a model binder for HTML forms with Yew](https://www.reddit.com/r/rust/comments/fedgwj/introducing_yew_form_a_model_binder_for_html/)
- url: https://www.reddit.com/r/rust/comments/fedgwj/introducing_yew_form_a_model_binder_for_html/
---
I love Yew, but if you're like me, you probably find that Yew is not great at binding  HTML form control to Rust `struct`.  Since my current project requires me to create rich form that require complex validation, I decided to complement Yew with  [yew\_form](https://github.com/jfbilodeau/yew_form), a set of mildly opinionated component to automatically map and validate a model to a HTML form.

Presently, it supports:

* Two way binding between Rust `struct` and HTML form control
* Validation using Keats Validator ([https://github.com/Keats/validator](https://github.com/Keats/validator))
* Support structs via `#[derive(Model)]`
* Support for nested structs
* Support for any scalar datatypes that implements `ToString` and `From&lt;Str&gt;`
* Support for any custom types by implementing `yew_form::FormValue`
* Build in `Field` and `CheckBox` component for binding

There's an online demo at [http://chronogears.com/yew-form/](http://chronogears.com/yew-form/) and the source for the demo is available [here](https://github.com/jfbilodeau/yew_form/tree/master/examples/form)

There's still work to be done, but it's already quite usable at my end. I hope this can be of help to others Yew users out there!

J-F
## [11][This Month in Rust GameDev #7 - February 2020](https://www.reddit.com/r/rust/comments/fe03sb/this_month_in_rust_gamedev_7_february_2020/)
- url: https://rust-gamedev.github.io/posts/newsletter-007
---

## [12][WebGPU, Frame Capturing and more in the latest update for Nannou - an open source, creative coding framework for Rust!](https://www.reddit.com/r/rust/comments/fduj17/webgpu_frame_capturing_and_more_in_the_latest/)
- url: https://nannou.cc/posts/nannou_v0.13
---

