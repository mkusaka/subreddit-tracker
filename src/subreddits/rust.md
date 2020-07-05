# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (27/2020)!](https://www.reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.44]](https://www.reddit.com/r/rust/comments/gz8ut5/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/gz8ut5/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/g6v98u/official_rrust_whos_hiring_thread_for_jobseekers/).

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
## [3][Is it me or does it feel like Rust writes itself sometimes?](https://www.reddit.com/r/rust/comments/hlhm4n/is_it_me_or_does_it_feel_like_rust_writes_itself/)
- url: https://www.reddit.com/r/rust/comments/hlhm4n/is_it_me_or_does_it_feel_like_rust_writes_itself/
---
Just spent a few hours following the compiler around and eventually it led me to an answer. Learned a lot too!
## [4][mdBook diagrams-as-code preprocessors](https://www.reddit.com/r/rust/comments/hljkw9/mdbook_diagramsascode_preprocessors/)
- url: https://bryce.fisher-fleig.org/markdown-diagrams/
---

## [5][Announcing the v0.17 release of Yew!](https://www.reddit.com/r/rust/comments/hl73bk/announcing_the_v017_release_of_yew/)
- url: https://www.reddit.com/r/rust/comments/hl73bk/announcing_the_v017_release_of_yew/
---
Hello, I'm excited to share with all of you the [latest Yew release](https://github.com/yewstack/yew/releases/tag/0.17.0) :)

*(If you're not familiar, Yew is a framework for building client web apps with Rust &amp; WebAssembly!)*

This release cycle, a lot of effort was [focused on performance](https://github.com/yewstack/yew/projects/7)!

**Performance Improvements**

* Cut number of DOM API calls when creating new DOM elements
* Discovered that downcasting elements with `web-sys` was calling out to JS and so elided those calls
* Restructured the rendering process to batch DOM updates to allow the browser to better optimize how it updates the screen (HUGE impact)
* Added the `key` attribute for elements in lists so that you can tell Yew when an element or Component has moved within a list. Without keys, Yew would have to fully re-render the moved nodes. Now, it is able to move existing nodes into their new position in the DOM

A lot of times, people point out these [JS benchmarks](https://github.com/krausest/js-framework-benchmark) to make claims about the speed of different frameworks. We've found those benchmarks to be useful but they're certainly not comprehensive and don't enforce too much standardization of each framework's implementation. We aim to keep Yew's standing competitive but hope to build upon those relatively simple benchmarks in our own fork to uncover more performance wins in the future.

**Other Changes**

* Expanded the number of listener attributes that devs can use in the `html!` macro. For example, we now support `onerror`, `onloadend`, etc.
* HTML button element `type` can finally be specified. (We are looking into improving our general handling of element properties for the next release!)
* `Agent` API ergonomics have been improved. We will continue this effort in the next release, follow along [here](https://github.com/yewstack/yew/projects/6).
* Void elements like `&lt;br&gt;` and `&lt;input&gt;` are no longer allowed to have children.
* Improved API ergonomics for many of the services included within Yew like `FetchService` and `TimeoutService`.

This release was made possible by many contributors in the Yew community but a special shout to two new contributors who made significant contributions! [@siku2](https://github.com/siku2) and [@mkawalec](https://github.com/mkawalec)
## [6][New gtk-rs release: New crates, better APIs](https://www.reddit.com/r/rust/comments/hl5778/new_gtkrs_release_new_crates_better_apis/)
- url: https://gtk-rs.org/blog/2020/07/04/new-release.html
---

## [7][How do I recompile std lib?](https://www.reddit.com/r/rust/comments/hlhp8c/how_do_i_recompile_std_lib/)
- url: https://www.reddit.com/r/rust/comments/hlhp8c/how_do_i_recompile_std_lib/
---
Hello, fellow rustaceans, I want to test a change to one of the std lib functions. How do I recompile these changes? Mine is a default rustup installation.
## [8][Borrowing self twice: alternate solutions?](https://www.reddit.com/r/rust/comments/hlffvx/borrowing_self_twice_alternate_solutions/)
- url: https://www.reddit.com/r/rust/comments/hlffvx/borrowing_self_twice_alternate_solutions/
---
I'm borrowing self twice here:

    self.reg.set_hl(self.get_word());

I _think_ this is because the self borrow doesn't get returned until the end of the block, which in this case is the line. So when  the leftmost `self` is borrowed, even though logically `self.get_word()` is resolved to a concrete value, the rightmost `self` is still borrowed.

An obvious solution is this:

    let w = self.get_word();
    self.reg.set_hl(w);

Because once `w` is assigned a value and we move to the next line, `self` is not borrowed anymore and can be borrowed again.

Hypothetically, say I'm very stubborn and don't want to solve it this way. How else might I solve it?   I'm mostly curious to discover if there's any sort of convenient shorthand to better control the lifetime of borrowing `self`.
## [9][An optimized sparse merkle tree implementation](https://www.reddit.com/r/rust/comments/hljw95/an_optimized_sparse_merkle_tree_implementation/)
- url: https://github.com/jjyr/sparse-merkle-tree
---

## [10][Back to old tricks .. (or, baby steps in Rust)](https://www.reddit.com/r/rust/comments/hl7eht/back_to_old_tricks_or_baby_steps_in_rust/)
- url: https://donsbot.wordpress.com/2020/07/04/back-to-old-tricks-or-baby-steps-in-rust/
---

## [11][Text and voice chat protocol for client-server-client](https://www.reddit.com/r/rust/comments/hlkeb7/text_and_voice_chat_protocol_for/)
- url: https://www.reddit.com/r/rust/comments/hlkeb7/text_and_voice_chat_protocol_for/
---
Hello Rustaceans, I'm trying to build as a side project a simple text and voice chat in Rust, as a learning project (like discord), it needs to work on the web and it should go through a server (not p2p), so I think the only options I have are Websockets and WebRTC, but neither of those are ideal, websockets are TCP and WebRTC is difficult to implement (from what I've read online).

I've looked at [webrtc-unreliable](https://github.com/kyren/webrtc-unreliable) but the maximum length is 1200 bytes for a message, and I doubt it's ideal for transferring of audio, and doesn't look like there are other libraries for this use-case.

So what are the alternatives? Should I just settle in with websockets and/or webrtc?
## [12][OpenID Provider Libraries in Rust](https://www.reddit.com/r/rust/comments/hlmbp9/openid_provider_libraries_in_rust/)
- url: https://www.reddit.com/r/rust/comments/hlmbp9/openid_provider_libraries_in_rust/
---
Is there any open identity provider library in rust? 

I request the community to recommend open identity libraries (Example: IdentityServer4(C#))
