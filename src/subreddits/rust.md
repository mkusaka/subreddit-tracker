# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (3/2020)!](https://www.reddit.com/r/rust/comments/eo6pjy/hey_rustaceans_got_an_easy_question_ask_here_32020/)
- url: https://www.reddit.com/r/rust/comments/eo6pjy/hey_rustaceans_got_an_easy_question_ask_here_32020/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

The Rust-related IRC channels on irc.mozilla.org (click the links to open a web-based IRC client):

 - [#rust](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust) (general questions)
 - [#rust-beginners](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-beginners) (beginner questions)
 - [#cargo](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23cargo) (the package manager)
 - [#rust-gamedev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-gamedev) (graphics and video games, and see also [/r/rust_gamedev](https://www.reddit.com/r/rust_gamedev))
 - [#rust-osdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-osdev) (operating systems and embedded systems)
 - [#rust-webdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-webdev) (web development)
 - [#rust-networking](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-networking) (computer networking, and see also [/r/rust_networking](https://www.reddit.com/r/rust_networking))

Also check out [last week's thread](https://reddit.com/r/rust/comments/ekpqp7/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 321](https://www.reddit.com/r/rust/comments/epinfr/this_week_in_rust_321/)
- url: https://this-week-in-rust.org/blog/2020/01/14/this-week-in-rust-321/
---

## [3][Announcing feed-rs 0.2.0](https://www.reddit.com/r/rust/comments/eqsadx/announcing_feedrs_020/)
- url: https://www.reddit.com/r/rust/comments/eqsadx/announcing_feedrs_020/
---
Kumamoto-san started feed-rs a while ago, but moved on to other things.

He was kind enough to add me as a maintainer last year, but unfortunately I haven't had much time to learn Rust and keep working on feed-rs in the interim.

I finally got around to publishing 0.2.0.

Its my first "real" Rust project - I wanted something a little more than the exercism examples to learn some Rust, so its probably full of mistakes and poor code but clippy seems happy. Key changes:

* 2018 edition
* Align feed model around Atom spec
* Switch to event-based parser (xml-rs) to reduce peak memory usage and use of clone()
* Expanded test coverage
* Documentation improvements

I do plan on using this in a different project (a replacement for fever, a web+app based RSS/Atom reader that is no longer maintained), so I'm expecting to issue a few more releases to feed-rs to address bugs, handle more of the RSS spec etc.

Happy to take feedback, bug reports, broken RSS/Atom feeds etc at the GitHub repo.

Thanks!
## [4][New Rust user questions about the `actix` situation](https://www.reddit.com/r/rust/comments/eqh3sy/new_rust_user_questions_about_the_actix_situation/)
- url: https://www.reddit.com/r/rust/comments/eqh3sy/new_rust_user_questions_about_the_actix_situation/
---
I've been curious about rust for a couple years, but just last month I finally took the time to start my first full application. I am working with high-frequency tick data from stock exchanges, and naturally the tipping point for switching was the performance of `actix-web` in the TechEmpower benchmarks.


With actix-web under fire, I don't see any other framework in rust that [even compares in performance on the benchmarks](https://www.techempower.com/benchmarks/#section=data-r18&amp;hw=ph&amp;test=fortune&amp;l=yyku7z-f). These were important to me not only to demonstrate speed, but the [benchmarks are open source](https://github.com/TechEmpower/FrameworkBenchmarks/tree/master/frameworks/Rust) and act as a nice Rosetta Stone for when I'm moving services to new languages/frameworks.


My questions are:

* Was actix-web inherently fast due to design, or will fixing the "unsafe" issues kill it's performance? (if that even happens)
* Are there any comparable alternatives that aren't posted in TechEmpower? Excluding actix, rust isn't even beating out many Python frameworks for performance.


I feel discouraged at continuing in rust after reading this situation. I know the benchmarks are somewhat arbitrary, but I find them to be incredibly helpful when learning new languages and hope that more rust packages will be added for reference purposes.
## [5][Project Structure of a Web Application](https://www.reddit.com/r/rust/comments/eqw1o8/project_structure_of_a_web_application/)
- url: https://www.reddit.com/r/rust/comments/eqw1o8/project_structure_of_a_web_application/
---
I am looking forward to use **warp** for making a backend server while using **ReactJs** on the client side. I will be using **MySql** as my primary database with a layer of caching using **memcached.**

I have looked through few of the articles, but none of them were clear enough of the Project Structures they follow. 

Should all the handlers go to a single folder and further dividing the folder according to the domain problem ? Where does the **View** file goes ? Is **DataAccessLayer** a different library project or just a separate folder should do ? 

I am making a basic application, but in future I will be referencing this project. Therefore, I still want to keep it atleast moderately organized.

&amp;#x200B;

Thanks for the help in advance.
## [6][The most stable (and ergonomic) Rust web framework](https://www.reddit.com/r/rust/comments/eqtf5v/the_most_stable_and_ergonomic_rust_web_framework/)
- url: https://www.reddit.com/r/rust/comments/eqtf5v/the_most_stable_and_ergonomic_rust_web_framework/
---
With the whole actix debacle, I have been looking into different other web frameworks that exist and I have found some seemingly good ones such as Warp and Tower, and of course Rocket looks pretty neat too. 

In comments about actix, as well as in the repo READMEs, I see performance always being mentioned almost as a main selling point. Also, benchmarks (for which I have learnt framework teams can and will gladly overoptimize). But since, in my opinion, most apps will not really benefit from performance because of other components such as the database inducing overhead, I’m quite surprised people obsess over performance of the web server. Surprisingly nobody really mentions “the safest web framework”. I’d rather use a safe and correct, but slower server, than something making the safety/speed trade-off in a way actix did...

I’m a hobbyist though, so I don’t really know much. Thus I wanted to ask more experienced developers: what is, in your opinion, the most correct / stable / production-ready web framework suitable for production?
## [7][Drop is not equivalent to the "toilet closure"](https://www.reddit.com/r/rust/comments/eqlx7z/drop_is_not_equivalent_to_the_toilet_closure/)
- url: https://www.reddit.com/r/rust/comments/eqlx7z/drop_is_not_equivalent_to_the_toilet_closure/
---
You would think that the standard library's `fn drop&lt;T&gt;(_: T) {}` and the "toilet closure" `|_| {}` would be equivalent. But that's actually not the case.

Consider this callback taking function:

    fn foo(_: impl FnOnce(&amp;())) {}

For [this function](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=63f661682d894d226519c9f06a624407), it is valid to call `foo(|_| {})` but not `foo(drop)`, which fails with

    error[E0631]: type mismatch in function arguments
    --&gt; src/main.rs:4:9
    |
    1 | fn foo(_: impl FnOnce(&amp;())) {}
    |    ---         ----------- required by this bound in `foo`
    ...
    4 |     foo(drop);
    |         ^^^^
    |         |
    |         expected signature of `for&lt;'r&gt; fn(&amp;'r ()) -&gt; _`
    |         found signature of `fn(_) -&gt; _`
    
    error[E0271]: type mismatch resolving `for&lt;'r&gt; &lt;fn(_) {std::mem::drop::&lt;_&gt;} as std::ops::FnOnce&lt;(&amp;'r (),)&gt;&gt;::Output == ()`
    --&gt; src/main.rs:4:5
    |
    1 | fn foo(_: impl FnOnce(&amp;())) {}
    |    ---         ----------- required by this bound in `foo`
    ...
    4 |     foo(drop);
    |     ^^^ expected bound lifetime parameter, found concrete lifetime
## [8][How will the rust community change after the actix situation?](https://www.reddit.com/r/rust/comments/eqw8ub/how_will_the_rust_community_change_after_the/)
- url: https://www.reddit.com/r/rust/comments/eqw8ub/how_will_the_rust_community_change_after_the/
---
Hello, 

now that the dust is starting to settle(?) for the whole actix situation, how do you think this will change the rust community, if at all? Do you think companies might stray from rust because of this?
## [9][Seeking assistance: Implement a function returning the local UTC offset (time crate, PRs accepted)](https://www.reddit.com/r/rust/comments/eqnti0/seeking_assistance_implement_a_function_returning/)
- url: https://github.com/time-rs/time/issues/203
---

## [10][Community thoughts about introducing a breaking change (remove the implicit into()) to the typed builder crate](https://www.reddit.com/r/rust/comments/eqjktn/community_thoughts_about_introducing_a_breaking/)
- url: https://www.reddit.com/r/rust/comments/eqjktn/community_thoughts_about_introducing_a_breaking/
---
I'm the author of the [typed-builder](https://crates.io/crates/typed-builder) crate. Following [some discussion in an pull request](https://github.com/idanarye/rust-typed-builder/pull/29), I'm considering making a breaking change, and would like to hear the community's opinion before I make a decision.

Currently, the setter functions generated by the `TypedBuilder` derive macro use `into()` on the argument. So, if we have something like:

    #[derive(TypedBuilder)]
    pub struct Foo {
        bar: String,
    }
    
The generated setter will look like (simplified a bit - the actual implementation has lots of generics ceremony):

    pub fn bar(self, bar: impl Into&lt;String&gt;) -&gt; SelfWithBar {
        SelfWithBar {
            bar: bar.into(),
        }
    }

This improves ergonomics for `String` and `Option` fields, but messes up the type inference a bit. The author of the PR I linked to earlier, for example, complains that this prevents Rust from inferring the builded struct type when it has fields with generic types. And with my own usages, I've also encountered inference problems induced by this behavior. Consider this:

    #[derive(typed_builder::TypedBuilder)]
    struct Foo {
        bar: usize,
        baz: usize,
    }

    Foo::builder()
        .bar(1usize) // will not compile without the literal suffix
        .baz("2".parse::&lt;usize&gt;().unwrap()) // will not compile with the turbofish
        .build();

`.bar(1)` would not work, because Rust by default treats integer literals as signed, and signed literal cannot `into()` unsigned ones because you'd lose the sign. Without the automatic `into()` it would work, because the argument of `bar()` would be a `usize` so Rust will know to make the literal a `usize`.

`.baz("2".parse().unwrap())` would not work, because Rust cannot determine what to parse the string as. It needs to parse it to something, and then convert that something to `usize`, but it has no way to determine the concrete type of that something.

I'm also starting to have second thoughts about doing it for strings. Yes, it's convenient, but it hides heap allocation - and the Rust philosophy is to make expensive operations stand out. So removing it entirely will make the typed-builder crate more rusty.

So, I'm thinking of removing the implicit `into()`, or maybe make it opt-in. This, of course, will be a breaking change - but I think it's worth it, because it'll force users to strap `.to_owned()` on the strings they pass to it, which makes the heap allocation explicit. But I want some community input before I do it - are you using typed-builder? What do you think about making this change?
## [11][The Soundness Pledge](https://www.reddit.com/r/rust/comments/eqcefv/the_soundness_pledge/)
- url: https://raphlinus.github.io/rust/2020/01/18/soundness-pledge.html
---

## [12][Rust Design Pattern: Sharing objects over async thread pool](https://www.reddit.com/r/rust/comments/eqraw2/rust_design_pattern_sharing_objects_over_async/)
- url: https://www.reddit.com/r/rust/comments/eqraw2/rust_design_pattern_sharing_objects_over_async/
---
I am new to Rust. I am writing an async HTTP/2 service with `hyper` `0.13`. I want to share a global configure object. Because the configuration can be modified at runtime. I protect it by wrapping it with a Mutex. 

```
pub struct Cluster {
    ring: HashRing&lt;String&gt;,
    nodes: HashMap&lt;String, NodeClient&gt;,
}

impl Send for Cluster {}

lazy_static! {
    static ref CLUSTER: Mutex&lt;Cluster&gt; = Mutex::new(Cluster::new(Vec::new()));
}
```

However, I am getting the following compile error. 

```
error[E0277]: `std::sync::MutexGuard&lt;'_, cluster::Cluster&gt;` cannot be sent between threads safely
  --&gt; src/server.rs:15:38
   |
15 |     let server = Server::bind(&amp;addr).serve(service);
   |                                      ^^^^^ `std::sync::MutexGuard&lt;'_, cluster::Cluster&gt;` cannot be sent between threads safely
   |
```
The `std::sync::MutexGuard` is can't be `Send` between threads. What would be the design pattern that is similar to global objects that can be used in the thread pool of an async application? Ordinary objects can only be used by one single thread. Cloning the object for each async call will be unnecessarily expensive. What is the ideal design pattern in Rust to share an object that can be modified at runtime across threads?
