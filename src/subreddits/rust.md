# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (1/2020)!](https://www.reddit.com/r/rust/comments/ehk18j/hey_rustaceans_got_an_easy_question_ask_here_12020/)
- url: https://www.reddit.com/r/rust/comments/ehk18j/hey_rustaceans_got_an_easy_question_ask_here_12020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/eelr59/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 319](https://www.reddit.com/r/rust/comments/ej4eom/this_week_in_rust_319/)
- url: https://this-week-in-rust.org/blog/2019/12/31/this-week-in-rust-319/
---

## [3][Helps understanding a post on performances/syscalls](https://www.reddit.com/r/rust/comments/eka8rh/helps_understanding_a_post_on_performancessyscalls/)
- url: https://www.reddit.com/r/rust/comments/eka8rh/helps_understanding_a_post_on_performancessyscalls/
---
Recently (well, half an hour ago) I stumbled across this ( https://drewdevault.com/2020/01/04/Slow.html ) post where the OP compares the "ideal" number of syscalls for an hello world vs the number of syscalls in a bunch of languages. It's a fun little comparison and I got curious. Some tests impressed me for better (Zig for example) or worse (C glib vs the musl counterpart), but I got puzzled as soon as I reached Rust. Rust does a wopping total of 123 total syscalls, 21 unique (so far more than write/exit). I don't have much experience with strace and the similar, nor with the rust's internals, so I asks you experts: why this difference? whose "fault" is that? Is it possible to reduce the gap with the ideal case without inlining assembler?

The code:

```
fn main() {
    println!("hello world");
}
```

Compiled with: `rustc -C opt-levels=s test.rs`

It doesn't say with version of the compiler, so I assume the latest stable. It's just a tought experiment, I don't think the author strived for the most performant but a "classic" version for every language. Thank you for everyone that will take time to help me understand! :)
## [4][Projects for beginners](https://www.reddit.com/r/rust/comments/ek7mqr/projects_for_beginners/)
- url: https://www.reddit.com/r/rust/comments/ek7mqr/projects_for_beginners/
---
I'm started learning Rust, but the actual books are a little boring.

Does someone have an idea of projects that a beginner rust programmer can do?
## [5][Blog Post: Mutexes Are Faster Than Spinlocks](https://www.reddit.com/r/rust/comments/ejx7y8/blog_post_mutexes_are_faster_than_spinlocks/)
- url: https://matklad.github.io/2020/01/04/mutexes-are-faster-than-spinlocks.html
---

## [6][How can I prevent expanding [ T ] to [ [ T ] ] while implementing a trait?](https://www.reddit.com/r/rust/comments/ekbhqn/how_can_i_prevent_expanding_t_to_t_while/)
- url: https://www.reddit.com/r/rust/comments/ekbhqn/how_can_i_prevent_expanding_t_to_t_while/
---
I'm trying to overload methods using a specialization. 

&amp;#x200B;

Main trait is

`pub trait Validate&lt;T: ?Sized&gt; {`

`type Output: Output;`

`fn validate(&amp;mut self, node: &amp;T) -&gt; Self::Output;`

`}`

&amp;#x200B;

and it explodes when I add

&amp;#x200B;

`impl&lt;T, V, O, E&gt; Validate&lt;[T]&gt; for V`

`where`

`Self: Validate&lt;T, Output = Result&lt;O, E&gt;&gt;,`

`{`

`type Output = Result&lt;Vec&lt;O&gt;, E&gt;;`

`fn validate(&amp;mut self, nodes: &amp;[T]) -&gt; Self::Output {`

`nodes.iter().map(|node| self.validate(node)).collect()`

`}`

`}`

`impl Analyzer {`

`pub fn check&lt;T, O&gt;(&amp;mut self, node: &amp;T) -&gt; Option&lt;O&gt;`

`where`

`Self: Validate&lt;T, Output = Result&lt;O, Error&gt;&gt;,`

`{`

`let res: Result&lt;O, _&gt; = self.validate(node);`

`match res {`

`Ok(v) =&gt; Some(v),`

`Err(..) =&gt; {`

`// handle error`

`None`

`}`

`}`

`}`

`}`

&amp;#x200B;

&amp;#x200B;

(Output trait is a trick to bypass the current limitation of the specialization)

&amp;#x200B;

As rustc prints

= note: required because of the requirements on the impl of \`validator::Validate&lt;\[\_\]&gt;\` for \`analyzer::Analyzer&lt;'\_, '\_&gt;\`

= note: required because of the requirements on the impl of \`validator::Validate&lt;\[\[\_\]\]&gt;\` for \`analyzer::Analyzer&lt;'\_, '\_&gt;\`

= note: required because of the requirements on the impl of \`validator::Validate&lt;\[\[\[\_\]\]\]&gt;\` for \`analyzer::Analyzer&lt;'\_, '\_&gt;\`

&amp;#x200B;

I think it can be solved by preventing impl of \`Validate&lt;\[\[T\]\]&gt;\`. How can I do this?

&amp;#x200B;

&amp;#x200B;

&amp;#x200B;

&amp;#x200B;

Playground: [https://play.rust-lang.org/?version=nightly&amp;mode=debug&amp;edition=2018&amp;gist=3145207592ca5a8ad41fd691094ce7ec](https://play.rust-lang.org/?version=nightly&amp;mode=debug&amp;edition=2018&amp;gist=3145207592ca5a8ad41fd691094ce7ec)
## [7][Google's OSS-fuzz officially supports Rust since August](https://www.reddit.com/r/rust/comments/ek2l52/googles_ossfuzz_officially_supports_rust_since/)
- url: https://github.com/google/oss-fuzz
---

## [8][I just (sort of) finished my first project in rust, a simple image viewer and editor](https://www.reddit.com/r/rust/comments/ek4r9j/i_just_sort_of_finished_my_first_project_in_rust/)
- url: https://github.com/sam-barr/png-rs
---

## [9][pkger - automated .rpm and .deb packaging using docker](https://www.reddit.com/r/rust/comments/ekczda/pkger_automated_rpm_and_deb_packaging_using_docker/)
- url: https://www.reddit.com/r/rust/comments/ekczda/pkger_automated_rpm_and_deb_packaging_using_docker/
---
Hey Rustaceans!

I wanted to show off my new project that I've started working on. It's called `pkger` and the idea behind it is to make it simple to build simple binary packages for multiple os, versions and architectures. To achieve this we have to define a recipe (containing all necessary build, install and metadata info) based on which `pkger` will build the desired package using docker's api to spawn and manipulate containers. Let me know if you have any ideas, suggestions or tips :) 

Here is the repository: [pkger](https://github.com/wojciechkepka/pkger)
## [10][Explaining Atomics in Rust](https://www.reddit.com/r/rust/comments/ejz4rs/explaining_atomics_in_rust/)
- url: https://cfsamsonbooks.gitbook.io/explaining-atomics-in-rust/
---

## [11][chacha20_poly1305_rs: Implementation of ChaCha20, Poly1305, and an unsigned 256-bit integer (U256) in pure Rust with no additional dependencies](https://www.reddit.com/r/rust/comments/ek4r1x/chacha20_poly1305_rs_implementation_of_chacha20/)
- url: https://github.com/jmg292/chacha20_poly1305_rs
---

## [12][Introduce async-socks5: async/.await version of SOCKS5!](https://www.reddit.com/r/rust/comments/ejwbaq/introduce_asyncsocks5_asyncawait_version_of_socks5/)
- url: https://github.com/ark0f/async-socks5
---

