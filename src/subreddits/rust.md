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
## [2][This Week in Rust 351](https://www.reddit.com/r/rust/comments/i8iqb9/this_week_in_rust_351/)
- url: https://this-week-in-rust.org/blog/2020/08/11/this-week-in-rust-351/
---

## [3][Blog Post: Code Smell: Concrete Abstraction](https://www.reddit.com/r/rust/comments/iaic5w/blog_post_code_smell_concrete_abstraction/)
- url: https://matklad.github.io/2020/08/15/concrete-abstraction.html
---

## [4][Kosmonaut: Web browser from scratch in Rust](https://www.reddit.com/r/rust/comments/iab2sm/kosmonaut_web_browser_from_scratch_in_rust/)
- url: https://github.com/twilco/kosmonaut
---

## [5][Show /r/rust: A GUI Backup tool, written in Rust using wgpu](https://www.reddit.com/r/rust/comments/iaqpxs/show_rrust_a_gui_backup_tool_written_in_rust/)
- url: https://github.com/KongouDesu/BackupGUI
---

## [6][Why Rust's Unsafe Works](https://www.reddit.com/r/rust/comments/iab5y6/why_rusts_unsafe_works/)
- url: https://jam1.re/blog/why-rusts-unsafe-works
---

## [7][Saving Servo](https://www.reddit.com/r/rust/comments/iaeiov/saving_servo/)
- url: https://www.reddit.com/r/rust/comments/iaeiov/saving_servo/
---
Reflecting on the latest Mozilla layoffs, I'm sure Rust will bounce back, but I'm curious about the future of Servo. From what I understand the whole Servo team was laid off, which is a real shame. Whilst the project is open source and contributions are still ongoing, the reality is that developers need to pay their bills and it would appear likely that we'll see a slowdown in activity. What I'm wondering about is whether there is an opportunity to build a funding model to support it, either through donations (Patreon, etc...) or through commercialisation.

Regarding commercialisation, one idea that springs to mind is to use it in an embedded mobile browser product similar to Adobe PhoneGap ([https://phonegap.com/](https://phonegap.com/)), but with a stronger focus on WebAssembly performance. WebAssembly seems like an ideal base for cross-platform mobile development, and although Servo isn't quite ready for primetime as a general purpose web browser engine it seems that it could work well if support was targetted at best-in-class WebAssembly performance for mobile apps, which is also something I think mobile app developers would be willing to pay for.

Aside from the suggestion above, does anyone have any thoughts on how to help Servo to continue to develop?
## [8][Are the memory representations of T&lt;A&gt; and T&lt;B&gt; identical if A has the transparent representation of B?](https://www.reddit.com/r/rust/comments/iaqow5/are_the_memory_representations_of_ta_and_tb/)
- url: https://www.reddit.com/r/rust/comments/iaqow5/are_the_memory_representations_of_ta_and_tb/
---
Suppose that we have

    #[repr(transparent)]
    struct A(B);

Are the ABIs of T&lt;A&gt; and T&lt;B&gt; identical no matter how we define T, so that transmuting between T&lt;A&gt; and T&lt;B&gt; always has a guaranteed behavior?

I mean, transmuting might not be safe, however, it does result in a valid value. Am I right?
## [9][How does `impl` work when returning a trait-bound type?](https://www.reddit.com/r/rust/comments/iao53b/how_does_impl_work_when_returning_a_traitbound/)
- url: https://www.reddit.com/r/rust/comments/iao53b/how_does_impl_work_when_returning_a_traitbound/
---
Hello everyone,
Yesterday I tried to implement a simple `composition` function that takes a couple of other functions and returns their mathematical composition as a closure.

This was my first attempt:
```rust
fn composition&lt;A, B, C, F : Fn(B) -&gt; C, G : Fn(A) -&gt; B&gt;(f : F, g : G) -&gt; Fn(A) -&gt; C {
    move |a| f(g(a))
}
```

It doesn't work because the size of the return type (`Fn(A) -&gt; C`) is not known at compile time (`Error[E0746]: return type cannot have an unboxed trait object`). I was ready to try something involving `Box` when I stumbled upon the `impl` keyword applied to return types as a convenient way to apply trait bounds. The following iteration works like a charm:

```rust
fn composition&lt;A, B, C, F : Fn(B) -&gt; C, G : Fn(A) -&gt; B&gt;(f : F, g : G) -&gt; impl Fn(A) -&gt; C {
    move |a| f(g(a))
}
```

I would like to ask a little clarification on what is going on under the hood. I can see why my first try didn't work: `Fn(A) -&gt; C` by itself is not enough to know how much memory must be allocated on the stack, and Rust works mainly with stack memory (unless otherwise specified, with types like `Box` and `Vec`). How does adding `impl` fix this? Is it treated like as a generic/template/type constructor and the compiler infers the actual return type for each usage of the function? Or does it hide some kind of heap allocation?

My primary usage for rust is focused on embedded development where dynamic memory is not a given, so I'm particularly interested in this topic.
## [10][Temporarily opt-in to shared mutation](https://www.reddit.com/r/rust/comments/iaajiu/temporarily_optin_to_shared_mutation/)
- url: https://ryhl.io/blog/temporary-shared-mutation/
---

## [11][How can I change syntax highlighting of rust analyzer in VS Code?](https://www.reddit.com/r/rust/comments/iak2hk/how_can_i_change_syntax_highlighting_of_rust/)
- url: https://www.reddit.com/r/rust/comments/iak2hk/how_can_i_change_syntax_highlighting_of_rust/
---
Rust analyzer currently underlines mutable variables in vs code, but I'd rather make them bold. If this is possible, how do I do it?

[This](https://imgur.com/a/l4snYD1) is how it looks like at the moment. I don't like the underlines (they also hide error squiggles).
## [12][[Ask] Explicit ways to express argument mutability](https://www.reddit.com/r/rust/comments/ianmbk/ask_explicit_ways_to_express_argument_mutability/)
- url: https://www.reddit.com/r/rust/comments/ianmbk/ask_explicit_ways_to_express_argument_mutability/
---
Hi guys,

I would like to ask something. There is an annoying thing in Rust I could not solve yet nicely, namely call site mutability notation.

    pub fn func(vars: Vec&lt;Variable&gt;) {
        immut_call(&amp;vars);      // accepts &amp;Vec&lt;Variable&gt;
        mut_call(&amp;mut vars);    // accepts &amp;mut Vec&lt;Variable&gt;
    }

In this example it is obvious for the reader, which function modifies `vars`. Now let's change the input param type to a &amp;mut ptr.

    pub fn func(vars: &amp;mut Vec&lt;Variable&gt;) {
        immut_call(vars); 
        mut_call(vars);   
    }
My question is, is there a nice way to somehow make it explicit which call (immut_call and mut_call) will modify its input parameter? I can solve it by introducing a new scope around `immut_call` and shadowing the original `vars` variable with an immutable one, or by casting it to the right mutability in the argument list, but they are quite ugly solutions.
