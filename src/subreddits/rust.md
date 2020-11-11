# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (46/2020)!](https://www.reddit.com/r/rust/comments/jqrkpa/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jqrkpa/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What’s everyone working on this week (46/2020)?](https://www.reddit.com/r/rust/comments/jqrmqn/whats_everyone_working_on_this_week_462020/)
- url: https://www.reddit.com/r/rust/comments/jqrmqn/whats_everyone_working_on_this_week_462020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-46-2020/51226?u=llogiq)!
## [3][I present you baru](https://www.reddit.com/r/rust/comments/js6ypa/i_present_you_baru/)
- url: https://www.reddit.com/r/rust/comments/js6ypa/i_present_you_baru/
---
&amp;#x200B;

[baru](https://preview.redd.it/su7j8m6gmly51.png?width=1470&amp;format=png&amp;auto=webp&amp;s=cb98832b6eba256ce2f3d5c2e1b285d7fed52cf8)

Hi,

[baru](https://github.com/doums/baru) is a system monitor that I wrote in Rust and C  with performance in mind. I create it to use it as a filler for my windows manager status bar. Since it's just plain text output you can use it for other purposes I think. For the past few months I have been using it on a daily basis.

Text output is fully customizable. Label are dynamic. Currently it supports the following modules:

* date and time
* battery (level, status, design level based)
* wireless (state, essid, signal strength)
* wired (state)
* audio sink and source (level, muted)
* brightness
* cpu usage and temperature
* memory (percent or used/total in gigabyte/gibibyte)

Its configuration is in yaml.

If you are interested, you can install it just using the compiled [binary](https://github.com/doums/baru/releases). It is also present as an AUR [package](https://aur.archlinux.org/packages/baru-bin) for Arch Linux users.
## [4][Announcing InfluxDB IOx - The Future Core of InfluxDB Built with Rust and Arrow | InfluxData](https://www.reddit.com/r/rust/comments/jrstbt/announcing_influxdb_iox_the_future_core_of/)
- url: https://www.influxdata.com/blog/announcing-influxdb-iox/
---

## [5][Announcing once_cell 1.5.0: no_std non-blocking initialization](https://www.reddit.com/r/rust/comments/jrz6xi/announcing_once_cell_150_no_std_nonblocking/)
- url: https://docs.rs/once_cell/1.5.0/once_cell/race/index.html
---

## [6][Is there a list of Rust software that doesn't exist](https://www.reddit.com/r/rust/comments/js7g6v/is_there_a_list_of_rust_software_that_doesnt_exist/)
- url: https://www.reddit.com/r/rust/comments/js7g6v/is_there_a_list_of_rust_software_that_doesnt_exist/
---
Fairly easy question, I'm sure it's been asked before but I can't find it:

&amp;#x200B;

Is there a list of software that, in principle, would benefit from being written using Rust but no implementation using rust exists yet, or the implementation that exists lacks active maintenance and is outdated?

&amp;#x200B;

By this I mean everything ranging from drivers for a specific piece of hardware to coreutils that are complex and somewhat slower and less feature-rich rich than they could be (e.g. ripgrep being a lot better than grep in many ways) to high level apps like email clients.

&amp;#x200B;

Obviously any such list would be subjective, but I still think it might serve as inspiration for starting projects. Provided that every entry has some info attached to it (e.g. what are the alternatives, what are they lacking)

&amp;#x200B;

I'm asking this question to some extent for myself, but more so because I keep recommending people to learn Rust (instead of C or C++) as a systems programming language, and I believe in a learning-via-doing style approach.
## [7][the trait `std::convert::From&lt;i64&gt;` is not implemented for `f64`](https://www.reddit.com/r/rust/comments/js1avn/the_trait_stdconvertfromi64_is_not_implemented/)
- url: https://www.reddit.com/r/rust/comments/js1avn/the_trait_stdconvertfromi64_is_not_implemented/
---
I have an `i64` where I know that only 37 bits are used. I also know that it's non-negative. I want to convert it into an `f64`. Surprisingly, I can't find any simple way to do that. Neither the `From` nor the `TryFrom` traits are implemented in this direction. Is there something I'm missing?

I'd like to avoid an extra dependency for this, if at all possible.
## [8][Rust as a productive high-level language](https://www.reddit.com/r/rust/comments/jrtlhy/rust_as_a_productive_highlevel_language/)
- url: https://omarabid.com/rust-high-level-language
---

## [9][Tp-Note learns RestructuredText](https://www.reddit.com/r/rust/comments/js69fy/tpnote_learns_restructuredtext/)
- url: https://blog.getreu.net/20201101-tp-note-news3/
---

## [10][LLVM Code Coverage Instrumentation PR merged!](https://www.reddit.com/r/rust/comments/jrlnfk/llvm_code_coverage_instrumentation_pr_merged/)
- url: https://github.com/rust-lang/rust/issues/34701#issuecomment-723178087
---

## [11][Writing (Linux) BPF Code in Rust](https://www.reddit.com/r/rust/comments/jrukxs/writing_linux_bpf_code_in_rust/)
- url: https://blog.redsift.com/labs/writing-bpf-code-in-rust/
---

## [12][Q: Rust used to have a ton of pointer types; what were they and where did they go?](https://www.reddit.com/r/rust/comments/jrkiuw/q_rust_used_to_have_a_ton_of_pointer_types_what/)
- url: https://www.reddit.com/r/rust/comments/jrkiuw/q_rust_used_to_have_a_ton_of_pointer_types_what/
---
Today's Rust is more streamlined, but you may know that Rust used to have a ton of pointer types:

&gt; Older Rusts had many more pointer types, they’re gone now.
&gt;
&gt; — https://doc.rust-lang.org/guide-pointers.html

I'm looking to understand the chronology of them to understand the language's evolution a little more deeply.

Q1: What were the older pointer types? Ideally I'm looking for actual data from the Rust docs (and not specific memories people have), if possible. Are there previous versions of the Rust docs online that I can look at or do I have to recompile the old docs to do so?

Q2: What became of each pointer type as the language evolved (e.g. became redundant, functionality merged someplace else, dropped as a feature, etc.)?
