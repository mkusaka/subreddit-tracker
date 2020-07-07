# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (28/2020)!](https://www.reddit.com/r/rust/comments/hm1pws/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hm1pws/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (28/2020)?](https://www.reddit.com/r/rust/comments/hm1qi3/whats_everyone_working_on_this_week_282020/)
- url: https://www.reddit.com/r/rust/comments/hm1qi3/whats_everyone_working_on_this_week_282020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-28-2020/45440?u=llogiq)!
## [3][Growing Gold with Rust](https://www.reddit.com/r/rust/comments/hmqjvs/growing_gold_with_rust/)
- url: https://www.reddit.com/r/rust/comments/hmqjvs/growing_gold_with_rust/
---
Hi everyone, 

I’m a scientist working in the field of nano-optics/-technology and at work we are regularly growing gold flakes -- thin platelets out of single-crystalline gold. How it works is still a bit of magic and a better understanding would be great.

[Some real gold flakes.](https://preview.redd.it/7iil1whf8e951.png?width=800&amp;format=png&amp;auto=webp&amp;s=1b7d211abc523654aa73e1329fd2f3d78b5c1e33)

As I got interested in Rust and had a lot of time during the last couple of months, I started implementing a simulation of the growing process in Rust. And it was a pleasure!

The main challenges were to find the right data structures for handling up to billions of atoms (in the end I used ndarray with some bitshifting/masking) at an acceptable speed (I settled with BTreeSet as a store for the surface vacancies) and also to deal with the fcc lattice and its arbitrary number of stacking faults of the gold crystals. I learned a lot about proper programming and important system details e.g. that the stack is limited to only a couple of megabytes or that the OS is really lazy when allocating memory. I wasn’t able to get DTrace running on Windows, so no nice flame graphs here, but I believe the bottleneck is that there is no quick way to randomly pick an element from the BTreeSet. At least I didn’t find one. If you have an idea, please tell me!

The code can be found at [Github](https://github.com/Rene-007/flake_growth), a [Wasm version](https://www.kullock.de/flake_growth/) is also available to get a first impression and I would recommend everybody to first have a look at the [visual guide](https://www.kullock.de/flake-growth-guide/).

&amp;#x200B;

[A short animation.](https://i.redd.it/stgdksuk8e951.gif)

In the end I have come to like the language, the tooling and the values behind Rust a lot! I think it is really the way to go forward for not only systems programming. Nevertheless, I had some difficulties and got some ideas I would like to share with you:

* Initially it was quite difficult to get started because I missed the ability to play around with the data and e.g. see if my mapping from memory to the fcc lattice and back is right. It is not possible to write unit tests for that, and one must simply play around and see whether the results in the 3D scene fit or not. At work I usually use Matlab for things like that and it would be the much easier tool to figure out the mapping, but I deliberately decided not to do so. (During my PhD time I used python/numpy a lot and observed that it is (mentally) hard to switch to another language once you already invested a lot into your code. And when looking on examples such as the HipHop Virtual Machine others seems have the problem, too.) So, from my point of view it would be nice to have some “playing-around capabilities”. I think I do not mean rapid prototyping with that but rather some small loop/block you can put into your program where the compiled code is interrupted by an interpreted section. When you run the program, you will end in the interpreter loop, have direct access to all the data structure/functions and can play around with them using the Rust syntax. So, no bindings or another language needed. I am not sure if that is feasible, but it would be cool.
* The second issue circles around libraries. As a beginner/outsider it is hard to judge which library/crate is needed, which might be the best one, which one I can trust. I can sympathize with the decision to have a small and stable std library but have the feeling that there should be some additional “meta crate” which combines the most popular matured crates in a complete way, i.e. that no additional external dependencies are needed. This crate should provide a root namespace (e.g. pop::rand or pop::serde), all unsafe parts should be reviewed  (with reports/discussion openly available) and some common programming standards (documentation, api design, naming) would also be good. It should be a big honor when “your crate” gets part of this “popular crate” or when you yourself become an approved reviewer. In contrast to the standard lib the API should guarantee compatibility only within a Rust edition such that subcrates, which are not state of the art anymore, can simply be removed. But as they will still live on as a separate crate, existing users just need the remove the “pop::” prefix within their source code. Smaller incompatible API changes within a subcrate might be introduced similar to linux in a “pop.next” meta-crate to smoothen the transition to the next edition. I think this might be a good compromise for a trustworthy base which is stable but not hammered in stone forever. What do you think?
## [4][Microsoft Research's Project Freta: "Given the history and preponderance of memory-corruption exploits, we made the choice as a team to embrace Rust at the beginning, architecting the entire capability from scratch in Rust from line one and building upon no existing software."](https://www.reddit.com/r/rust/comments/hmjsvs/microsoft_researchs_project_freta_given_the/)
- url: https://www.microsoft.com/en-us/research/blog/toward-trusted-sensing-for-the-cloud-introducing-project-freta/
---

## [5][Statically Sized Higher-kinded Polymorphism](https://www.reddit.com/r/rust/comments/hmqrx8/statically_sized_higherkinded_polymorphism/)
- url: http://blog.ielliott.io/sized-hkts/
---

## [6][Announcing Rustup 1.22.0](https://www.reddit.com/r/rust/comments/hmaddj/announcing_rustup_1220/)
- url: https://blog.rust-lang.org/2020/07/06/Rustup-1.22.0.html
---

## [7][Rustacean Station: Mun](https://www.reddit.com/r/rust/comments/hmqhwo/rustacean_station_mun/)
- url: https://rustacean-station.org/episode/020-mun/
---

## [8][A high-performance and safe runtime for Rust functions in Node.js apps](https://www.reddit.com/r/rust/comments/hmsskr/a_highperformance_and_safe_runtime_for_rust/)
- url: https://www.reddit.com/r/rust/comments/hmsskr/a_highperformance_and_safe_runtime_for_rust/
---
Hello, Rustaceans! Checkout how to run your Rust functions in Node.js in under 10 minutes. The SSVM is a WebAssembly VM optimized for server-side applications.

[https://www.secondstate.io/ssvm/](https://www.secondstate.io/ssvm/)  
 

SSVM offers Rust developers a way to build web apps, combining with WebAssembly's safety &amp; portabilit*y*; and JavaScript's productivity. Getting started here: [https://www.secondstate.io/articles/getting-started-with-rust-function/](https://www.secondstate.io/articles/getting-started-with-rust-function/)  
 

Why not just use v8? For starters, SSVM runs much faster than V8 (or even plain Docker ...)

[https://www.secondstate.io/articles/ssvm-performance/](https://www.secondstate.io/articles/ssvm-performance/)  
 

And SSVM supports extensions for server-side apps. 

SSVM supports WASI to access operating system resources in a controlled manner (ie capability-based security). Similar to WASI, the SSVM supports extensions to WebAssembly to safely access underlying systems — such as a storage interface for attached key/value storage, an AI inference interface for specialized hardware, and the Ethereum interface for a host blockchain.

SSVM is an open source project. Check it out: [https://github.com/second-state/SSVM](https://github.com/second-state/SSVM)
## [9][Comparison of Mun and Dyon, two Rusty programming languages written in Rust?](https://www.reddit.com/r/rust/comments/hmtngr/comparison_of_mun_and_dyon_two_rusty_programming/)
- url: https://www.reddit.com/r/rust/comments/hmtngr/comparison_of_mun_and_dyon_two_rusty_programming/
---
[Mun](https://github.com/mun-lang/mun) is a programming language with a Rusty-syntax, written in Rust, statically typed, with emphases on embedding and hot reloading.  I first learned of Mun through an [r/rust post](https://old.reddit.com/r/rust/comments/gkzmab/mun_v020_released/) a month ago, and I see that [a new Mun-related post](https://old.reddit.com/r/rust/comments/hmqhwo/rustacean_station_mun/) was just submitted a few hours ago.  

[Dyon](https://github.com/PistonDevelopers/dyon) describes itself as a rusty dynamically typed programming language.  It is also written in Rust. I have only just learned about it from[ this r/rust discussion](https://old.reddit.com/r/rust/comments/hmqjvs/growing_gold_with_rust/).

I would be interested in a comparison of the two programming languages from those who are more familiar with them.  At first glance, they seem to have similar goals, although static vs. dynamic typing is one obvious differentiator.

While I'm at it, are there any other Rusty programming languages in the works?  Given their similarities, I'm somewhat surprised that Dyon was not mentioned by anyone in the Mun discussion a month ago.  Maybe I'm overestimating how similar they are to each other.
## [10][`grass`, a pure Rust Sass compiler, can now compile Bootstrap 4 and Bulma-scss](https://www.reddit.com/r/rust/comments/hmctqc/grass_a_pure_rust_sass_compiler_can_now_compile/)
- url: https://www.reddit.com/r/rust/comments/hmctqc/grass_a_pure_rust_sass_compiler_can_now_compile/
---
After 6 months of work \`grass\` is finally able to compile both Bootstrap 4 and Bulma-scss! It's still not ready for production use and there are a lot of outstanding issues, but as long as your input is well formed and not too complex it should work very well.

The repository is at [https://github.com/connorskees/grass](https://github.com/connorskees/grass). 

With regard to performance, it is roughly even to \`libsass\` and several times faster than \`dart-sass\`. \`libsass\` is a bit faster to compile Bootstrap, but a bit slower to compile Bulma-scss. 

&amp;#x200B;

\`grass\` also compiles to wasm, though the payload is about 1MB.

There is a (very minimal) repl at  [https://connorskees.github.io/grassmeister/](https://connorskees.github.io/grassmeister/) to test it out.
## [11][Creating a custom target in Rust](https://www.reddit.com/r/rust/comments/hminsg/creating_a_custom_target_in_rust/)
- url: https://rust-embedded.github.io/embedonomicon/custom-target.html
---

## [12][async-net: Async networking primitives for TCP/UDP/Unix communication](https://www.reddit.com/r/rust/comments/hmfsjr/asyncnet_async_networking_primitives_for/)
- url: https://docs.rs/async-net
---

