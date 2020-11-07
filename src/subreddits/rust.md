# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (45/2020)!](https://www.reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 363](https://www.reddit.com/r/rust/comments/joxy7n/this_week_in_rust_363/)
- url: https://this-week-in-rust.org/blog/2020/11/04/this-week-in-rust-363/
---

## [3][Diagram of Async Architectures](https://www.reddit.com/r/rust/comments/jpcv2s/diagram_of_async_architectures/)
- url: https://i.redd.it/6kxvfm94kox51.png
---

## [4][Theseus is a new OS written from scratch in Rust to experiment with novel OS structure, better state management, and how to shift OS responsibilities like resource management into the compiler.](https://www.reddit.com/r/rust/comments/jpfuwe/theseus_is_a_new_os_written_from_scratch_in_rust/)
- url: https://github.com/theseus-os/Theseus
---

## [5][Simple brainf*** transpiler written in Rust (PoC)](https://www.reddit.com/r/rust/comments/jpoa4x/simple_brainf_transpiler_written_in_rust_poc/)
- url: https://github.com/Wafelack/brainrust
---

## [6][RustFest Global 2020 is starting!](https://www.reddit.com/r/rust/comments/jphgqq/rustfest_global_2020_is_starting/)
- url: https://rustfest.global/
---

## [7][40 Millisencond Bug](https://www.reddit.com/r/rust/comments/jp23jq/40_millisencond_bug/)
- url: https://vorner.github.io/2020/11/06/40-ms-bug.html
---

## [8][rust book](https://www.reddit.com/r/rust/comments/jpqvw2/rust_book/)
- url: https://www.reddit.com/r/rust/comments/jpqvw2/rust_book/
---
is there any up to date book about programming in rust?

I know there is  The Rust Programming Language but it only covers 2018 rust.

thanks in advance.
## [9][Build gtk rust apps with Builder](https://www.reddit.com/r/rust/comments/jpb7qs/build_gtk_rust_apps_with_builder/)
- url: https://wiki.gnome.org/Apps/Builder
---

## [10][Rust vs Go](https://www.reddit.com/r/rust/comments/jp4orc/rust_vs_go/)
- url: https://bitfieldconsulting.com/golang/rust-vs-go
---

## [11][Defining a Fn bound with generic return type that requires a lifetime](https://www.reddit.com/r/rust/comments/jpg0pp/defining_a_fn_bound_with_generic_return_type_that/)
- url: https://www.reddit.com/r/rust/comments/jpg0pp/defining_a_fn_bound_with_generic_return_type_that/
---
Today I stumbled upon a rather specific problem, where I'm unsure whether it is currently possible to do this.

I want to define a function where one parameter is a closure (or rather, a generic parameter bound with a Fn trait). However, the closure should itself return a generic parameter with a trait containing a lifetime (to be precise, an associated type should contain the lifetime). This lifetime needs to be bound to the input of the closure.

The desired function signature might look something like this:

    fn take_closure_with_trait&lt;R&gt;(&amp;mut self, f: impl FnOnce(&amp;Self) -&gt; R)
    where
        R: IntoIterator&lt;Item=&amp;Self&gt;

Indeed, this does not compile due to a missing lifetime specifier.

For a closure returning a concrete type with a lifetime all is fine. The problem is that I know of no syntax to specifiy the lifetime within the trait bound, as a higher-ranked lifetime is required here. If `impl trait` would be possible for closures, I could write something like `for&lt;'a&gt; F: FnOnce(&amp;'a Self) -&gt; impl IntoIterator&lt;Item=&amp;'a Self&gt;`, but `impl trait` is not available in this position.

Playground link: https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=c9b33b336d4540649b6007e5c5e923f1

Does anyone know whether this is even possible to express currently and if yes, how?
## [12][ngrok 0.1.0](https://www.reddit.com/r/rust/comments/jp8eci/ngrok_010/)
- url: https://www.reddit.com/r/rust/comments/jp8eci/ngrok_010/
---
I just published [https://github.com/nkconnor/ngrok](https://github.com/nkconnor/ngrok) to [crates.io](https://crates.io). Often times I run into scenarios where I need to communicate with some API asynchronously via callbacks. Ngrok makes it pretty trivial to open a public URL to my dev server. This library lets you run ngrok and retrieve callback URLs at runtime using the free plan. An example  


```  
let ngrok = ngrok::builder()  
.http()  
.port(3030)  
.run().unwrap();  


let callback: url::Url = ngrok.tunnel()?.http();
```  


Hope it's useful to somebody here. Bug reports and contributions welcome.
