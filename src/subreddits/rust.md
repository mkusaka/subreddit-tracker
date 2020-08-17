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
## [3][Veloren, an open-source multiplayer voxel RPG written in Rust, just had its 0.7 release with a tonne of new features!](https://www.reddit.com/r/rust/comments/ib83bl/veloren_an_opensource_multiplayer_voxel_rpg/)
- url: https://veloren.net/
---

## [4][🎉 The Embedded Working Group Newsletter - 24](https://www.reddit.com/r/rust/comments/ibcvnz/the_embedded_working_group_newsletter_24/)
- url: https://rust-embedded.github.io/blog/newsletter-24/
---

## [5][Stackage for Rust?](https://www.reddit.com/r/rust/comments/ib9orc/stackage_for_rust/)
- url: https://www.snoyman.com/blog/2020/08/stackage-for-rust
---

## [6][Cranelift can now compile rustc- giving nearly 7x faster compilations than LLVM!](https://www.reddit.com/r/rust/comments/iat25g/cranelift_can_now_compile_rustc_giving_nearly_7x/)
- url: https://github.com/bjorn3/rustc_codegen_cranelift/issues/743#issuecomment-674526510
---

## [7][Short Post: Using Long Paths in Windows and Rust](https://www.reddit.com/r/rust/comments/ibbv4t/short_post_using_long_paths_in_windows_and_rust/)
- url: https://gal.hagever.com/posts/windows-long-paths-in-rust/
---

## [8][Tunshell - A drop-in remote shell client and server written in Rust for easy shell access to deployment pipelines or other ephemeral environments, working behind NAT or firewall](https://www.reddit.com/r/rust/comments/ibcvpe/tunshell_a_dropin_remote_shell_client_and_server/)
- url: https://github.com/TimeToogo/tunshell
---

## [9][[announce] wide-0.5.0: now with more types!](https://www.reddit.com/r/rust/comments/ib5vok/announce_wide050_now_with_more_types/)
- url: https://www.reddit.com/r/rust/comments/ib5vok/announce_wide050_now_with_more_types/
---
The [wide](https://docs.rs/wide) crate has a new 0.5 version out! It's a crate to help make SIMD-friendly programming easier on you, the programmer.

Major features include:
* Supports all the common 128-bit SIMD types: f32x4, f64x2, i8x16, u8x16, i16x8, u16x8, i32x4, u32x4, i64x2, u64x2
* Supports many of the things you'd want to do with those types. What you can currently do is largely driven by the availability of hardware support on x86 / x86_64 machines.
* Software fallbacks for when there isn't hardware support are sometimes understood by LLVM and optimized appropriately, and sometimes LLVM doesn't get it and performance is just worse. This will improve some day in the far future when the ARM/Aarch64 intrinsics become Stable.
* `no_std`
* All cpu feature selection is done at compile time only via `#[cfg()]` and RUSTFLAGS and so forth, there's no runtime feature detection mechanism.
* Essentially no docs at the moment, but method names are the same as normal math function names, so it's fairly clear on its own.

If you want to contribute, I'm open to PRs as long as the crate continues to build on Stable by default.
## [10][Webassembly Without The Browser (Part 1)](https://www.reddit.com/r/rust/comments/iayqyv/webassembly_without_the_browser_part_1/)
- url: https://alexene.dev/2020/08/17/webassembly-without-the-browser-part-1.html
---

## [11][Using the quote crate, how to use the quote macro on &amp;str?](https://www.reddit.com/r/rust/comments/ibbw4v/using_the_quote_crate_how_to_use_the_quote_macro/)
- url: https://www.reddit.com/r/rust/comments/ibbw4v/using_the_quote_crate_how_to_use_the_quote_macro/
---
I am using the quote crate to write some Rust code. I have variables of type &amp;str which I want include in the quote macro (quote!).

Now the problem is that the &amp;str remains &amp;str. I want the quotes to be removed. So let's say I have this code:

    let some_import = "use crate::SomeStruct;";
    let generated = quote::quote! { #some_import };

The result is that it is interpretated literally with the quotes, so my import literally looks like this:

    "use crate::SomeStruct;"

But I want this:

    use crate::SomeStruct;

Normally I use format\_ident for &amp;str/String to remove the quotes, but it crashes saying ' use crate::SomeStruct;' is not a valid identifier.

How do I use the quote::quote macro to remove quotes from a &amp;str/String?
## [12][A steganographic tool for hiding messages in x86 binaries, written in Rust](https://www.reddit.com/r/rust/comments/iatp22/a_steganographic_tool_for_hiding_messages_in_x86/)
- url: https://github.com/woodruffw/steg86
---

