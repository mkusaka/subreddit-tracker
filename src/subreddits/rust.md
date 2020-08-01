# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (31/2020)!](https://www.reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hurj77/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 349](https://www.reddit.com/r/rust/comments/i094wo/this_week_in_rust_349/)
- url: https://this-week-in-rust.org/blog/2020/07/28/this-week-in-rust-349/
---

## [3][Alacritty Terminal Emulator Version 0.5.0 now with vi mode and search](https://www.reddit.com/r/rust/comments/i1j5lk/alacritty_terminal_emulator_version_050_now_with/)
- url: https://blog.christianduerr.com/alacritty_0_5_0_announcement.html
---

## [4][Rewritten in Rust: Modern Alternatives of Command-Line Tools](https://www.reddit.com/r/rust/comments/i1abpg/rewritten_in_rust_modern_alternatives_of/)
- url: https://zaiste.net/posts/shell-commands-rust/
---

## [5][Smooth Voxel Mapping: a Technical Deep Dive on Real-time Surface Nets and Texturing](https://www.reddit.com/r/rust/comments/i1obro/smooth_voxel_mapping_a_technical_deep_dive_on/)
- url: https://medium.com/@bonsairobo/smooth-voxel-mapping-a-technical-deep-dive-on-real-time-surface-nets-and-texturing-ef06d0f8ca14
---

## [6][Headcrab, a modern Rust debugger library: July 2020 progress report](https://www.reddit.com/r/rust/comments/i1bc5a/headcrab_a_modern_rust_debugger_library_july_2020/)
- url: https://headcrab-rs.github.io/2020/07/31/july-update.html
---

## [7][How does new error handling work?](https://www.reddit.com/r/rust/comments/i1lyy5/how_does_new_error_handling_work/)
- url: https://www.reddit.com/r/rust/comments/i1lyy5/how_does_new_error_handling_work/
---
I read [RFC 2504 `fix_error`](https://github.com/rust-lang/rfcs/blob/master/text/2504-fix-error.md), but it doesn't seem to say anything about how *using* the new API works.

My impression of the status after reading issues [53487](https://github.com/rust-lang/rust/issues/53487) and [58520](https://github.com/rust-lang/rust/issues/58520) is that the RFC is pretty much implemented in nightly (which I'm using). I've also read several articles about the changes and all the third-party error crates that are being deprecated now. But I'm still not finding any explanation of how to *attach* context to errors.

Here's a concrete example. How would I use the new error features to have the error returned from `do_stuff` include information about which step resulted in an error?

```rust
type MyErr = Box&lt;dyn std::error::Error&gt;;

fn main() -&gt; Result&lt;(), MyErr&gt; {
	let args: Vec&lt;String&gt; = std::env::args().collect();
	do_stuff(&amp;args[1], &amp;args[2], &amp;args[3])
}

fn do_stuff(file1: &amp;str, file2: &amp;str, file3: &amp;str) -&gt; Result&lt;(), MyErr&gt; {
	let text = std::fs::read_to_string(file1)?;
	std::fs::write(file2, text)?;
	std::fs::remove_file(file3)?;
	Ok(())
}
```
## [8][Announcing Rust.app: A new macOS App for Rustaceans](https://www.reddit.com/r/rust/comments/i1fvka/announcing_rustapp_a_new_macos_app_for_rustaceans/)
- url: https://github.com/XAMPPRocky/Rust.app/releases/tag/v0.1.0
---

## [9][The State of Rust with Steve Klabnik](https://www.reddit.com/r/rust/comments/i1ab0i/the_state_of_rust_with_steve_klabnik/)
- url: https://www.reddit.com/r/rust/comments/i1ab0i/the_state_of_rust_with_steve_klabnik/
---
Hey, folks. Just wanted to let everyone know that I just dropped the latest episode of my podcast, The Virtual World. In this episode I was able to sit down with Steve Klabnik, a member of the Rust core team. I asked him a ton of questions about himself and the Rust programming language and ecosystem. Some of the info may surprise you!

[https://anchor.fm/the-virtual-world/episodes/Ep-6--The-State-of-Rust-with-Steve-Klabnik-ehf8mk](https://anchor.fm/the-virtual-world/episodes/Ep-6--The-State-of-Rust-with-Steve-Klabnik-ehf8mk)

&amp;#x200B;

P.S. I'll be interviewing Alex Chrichton from the core team later today as well!
## [10][event-listener: a synchronization primitive to turn non-blocking data structures into async or blocking data structures](https://www.reddit.com/r/rust/comments/i16oly/eventlistener_a_synchronization_primitive_to_turn/)
- url: https://docs.rs/event-listener
---

## [11][async-io: Async I/O and timers (alternative to mio with a focus on simplicity)](https://www.reddit.com/r/rust/comments/i198og/asyncio_async_io_and_timers_alternative_to_mio/)
- url: https://docs.rs/async-io
---

## [12][📮 teloxide v0.3.0 released: An elegant Telegram bots framework for Rust](https://www.reddit.com/r/rust/comments/i1a7sr/teloxide_v030_released_an_elegant_telegram_bots/)
- url: https://github.com/teloxide/teloxide
---

