# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (17/2020)!](https://www.reddit.com/r/rust/comments/g4nu6j/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g4nu6j/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 335](https://www.reddit.com/r/rust/comments/g6d0ac/this_week_in_rust_335/)
- url: https://this-week-in-rust.org/blog/2020/04/21/this-week-in-rust-335/
---

## [3][3D graphics, Rust, Vulkan, ash](https://www.reddit.com/r/rust/comments/g7pypb/3d_graphics_rust_vulkan_ash/)
- url: https://hoj-senna.github.io/ashen-aetna/
---

## [4][Rust needs a wiki](https://www.reddit.com/r/rust/comments/g7s4ss/rust_needs_a_wiki/)
- url: https://www.reddit.com/r/rust/comments/g7s4ss/rust_needs_a_wiki/
---
Rust needs a wiki, or at least a hub for community documentation and other miscellaneous knowledge about the Rust language.

Currently when using the Rust language, it's quite easy to be efficient if you already know what you should be using, thanks to many things: cargo-doc which generates the documentation locally or online on docs.rs, the extensive book and docs for the standard library, and the community which often won't leave your questions answered on github/discord/IRC.

However, it's been really, really difficult over the last few years to start and get to that point where your Cargo.toml is stable and you start iterating quickly. Often, there will be a long process of trial and error where a user might try several crates for a given problem, or search extensively for a feature that might exist in rustc/cargo/other tool, but they just didn't know where to find it.

Some examples:

* The async story. What's the difference between tokio and async-std again? Do I need to read a blog post for that? Do I need to read a reddit post from a year ago that might already be outdated?
* The HTTP Server / Client story. Hyper ? Reqwest ? Iron ? Actix-web ? Rocket ? Tide ? What is the difference between all of those, what do they do?
* The 2D / 3D graphics story. I want to make a pong game. Piston? gfx? lyon? glium? What are the difference between all of these? By the way, I have this weird thing on windows where the command line still pops when I run my game on windows, what's up with that?
* Many other stories: Embedded, WASM, GUI, ... Even the IDE experience is a mess! RLS, rust-analyzer? When I want to know what are the recommended stuff for vim/nvim on google, I only find a post from a year ago, and even then there are multiple answers!

(Note that for most of the questions above, I did manage to answer them, but it took quite a bit of time that I wish I didn't spend. As we speak, some people might be struggling to understand those topics as well)

There have been attempts to alleviate the problem. http://www.arewewebyet.org/ https://areweguiyet.com/ https://arewegameyet.com/ those are great to know the rough state for each of those topics, but it stops there. Want to know which ones of those crates you should use together? Well, you could try to see if there is a real-life example that uses the crate you intend to use, if you're lucky it's been updated less than 2 years ago.

**The real solution is to have a community wiki**, akin to the one the Arch Linux community has. For reference, the Arch Linux wiki is the best tech-related hub of documentation linux users might even encounter, period. It is sober in design, and does its job. You need something? Look at the wiki, you'll probably find your answer there. You know something? Since it's overall maintained by the community, anyone can fix outdated stuff in a snap.

This would fix so many problems current Rust has:

* Comparative lists / tables of crates (or frameworks) to help new users find their way, without making a new website (website... that might not even be known by new rust users)
* Common issues within the rust community would be almost always up-to-date. Up-to-date "embedded" page, up-to-date "IDE Experience" page, up-to-date "GUI" page, ect.
* Community-driven: no risk of a useful blogpost being shutdown by its author. No risk of an important but outdated piece of info being on a website everyone knows the fix to, but nobody has the rights to change them.
* One unique source of knowledge, no more juggling between multiple blog posts to get what you want. (Bonus: backup friendly, less risk to lose something forever).


Of course the book, docs.rs and the official documentation should stay as they are. The wiki and the official rust website should just cross-reference each other to make sure everything is consistent. But there are so, so many topics that can be in a wiki, it really feels like there is a strong need for one. Look at me in the eyes and tell me [this awesome github repo](https://github.com/johnthagen/min-sized-rust) doesn't have the format of a wiki page.

What are your toughts on this? Any other points I might have missed? I don't know where to submit this as a serious proposal to the rust teams (I don't even know if I can!), but I do intend to make sure this idea gets discussed thoroughly and hopefully approved.
## [5][How feasible is it to write kernel modules in Rust ?](https://www.reddit.com/r/rust/comments/g7rag0/how_feasible_is_it_to_write_kernel_modules_in_rust/)
- url: https://www.reddit.com/r/rust/comments/g7rag0/how_feasible_is_it_to_write_kernel_modules_in_rust/
---
I am currently working on a PCIE FPGA, and I'll soon have to write a linux driver for it.
Is the ecosystem mature enough to write complicated kernel modules in Rust ? Is there really anything to gain from it, or do you spend your time doing FFI calls with C ?
## [6][What is the history of Ferris?](https://www.reddit.com/r/rust/comments/g7nmxs/what_is_the_history_of_ferris/)
- url: https://www.reddit.com/r/rust/comments/g7nmxs/what_is_the_history_of_ferris/
---
I am a Marylander and I love crabs.  Hello!  


I was wondering where Ferris came from.  I started programming rust because I saw that it was a language that loved crabs.  I noticed that Brendan Eich was from Maryland, as well, so I was wondering if that may have factored into it at all.  


Here's a picture of a crab.

https://imgur.com/K3E9PcK
## [7][Is there a build service for cargo to {deb,rpm,...}?](https://www.reddit.com/r/rust/comments/g7qxaw/is_there_a_build_service_for_cargo_to_debrpm/)
- url: https://www.reddit.com/r/rust/comments/g7qxaw/is_there_a_build_service_for_cargo_to_debrpm/
---
I would like to offer by crate (a linux binary) as packages for Ubunut, RedHat, CentOS, etc.

Is there a build service or a similar tool that make this easy?
## [8][Embedded Rust pattern: Zero Sized References](https://www.reddit.com/r/rust/comments/g79ib1/embedded_rust_pattern_zero_sized_references/)
- url: https://ferrous-systems.com/blog/zero-sized-references/
---

## [9][WIP UF2 bootloader in rust (also USB mass storage/scsi)](https://www.reddit.com/r/rust/comments/g7rlc6/wip_uf2_bootloader_in_rust_also_usb_mass/)
- url: https://www.reddit.com/r/rust/comments/g7rlc6/wip_uf2_bootloader_in_rust_also_usb_mass/
---
Thought I'd share this [rust UF2 Bootloader](https://github.com/cs2dsb/stm32-usb.rs/tree/master/firmware/usb_bootloader) I've been working on. 

It's basically a port of this [Bluepill UF2 Bootloader](https://github.com/lupyuen/bluepill-bootloader) but I've attempted to make the USB mass storage class, USB bulk-only transport protocol and USB scsi sub-class usable for proper FAT or anything else you might want to implement as MSC/BOT/SCSI over USB.

The above bootloader crate works for the bluepill board but could work with any embedded-hal that implements [usb-device](https://github.com/mvirkkunen/usb-device) without much extra work - the main thing you'd have to check is the `impl Flash for FlashWrapper` in usb\_bootloader/src/bin/msc.rs. I think eventually the Flash trait or something similar could be moved into embedded-hal so that you can read/write slices of bytes to your flash without worrying about page size and the other stuff you need to do to unlock/etc. 

The various usb layers I've implemented are here (links to crates.io &amp; docs.rs in readmes):

* [Mass storage class](https://github.com/cs2dsb/stm32-usb.rs/tree/master/firmware/usbd_mass_storage)
* [Bulk only transport protocol](https://github.com/cs2dsb/stm32-usb.rs/tree/master/firmware/usbd_bulk_only_transport)
* [SCSI transparent command set subclass](https://github.com/cs2dsb/stm32-usb.rs/tree/master/firmware/usbd_scsi)

There's also [uf2\_util](https://github.com/cs2dsb/stm32-usb.rs/tree/master/firmware/uf2_util) which will convert ELF/BIN to UF2. There's an example of the commands to do it in [build\_uf2](https://github.com/cs2dsb/stm32-usb.rs/blob/master/firmware/blink/build_uf2).

&amp;#x200B;

The rationale behind developing the bootloader was I was considering selling some PCB kits with a preflashed bootloader so people could update the firmware over USB rather than having to get a debug probe and getting the whole lot configured. I could have just used the C UF2 bootloader or similar but why make things easy? I also wanted to develop some application firmware that used USB mass storage for config/data transfer and a USB bootloader seemed like a concrete way to get started with lots of working open source examples to crib off.

&amp;#x200B;

Any feedback/suggestions/comments gratefully received :)
## [10][Announcing Tide 0.8.0!](https://www.reddit.com/r/rust/comments/g79gag/announcing_tide_080/)
- url: https://github.com/http-rs/tide/releases/tag/v0.8.0
---

## [11][Struggling with arithmetic overflow checking](https://www.reddit.com/r/rust/comments/g7ryba/struggling_with_arithmetic_overflow_checking/)
- url: https://www.reddit.com/r/rust/comments/g7ryba/struggling_with_arithmetic_overflow_checking/
---
I want to rewrite Jenkins hash function in Rust, in particular \`hashword\`: [https://android.googlesource.com/platform/external/jenkins-hash/+/75dbeadebd95869dd623a29b720678c5c5c55630/lookup3.c#173](https://android.googlesource.com/platform/external/jenkins-hash/+/75dbeadebd95869dd623a29b720678c5c5c55630/lookup3.c#173)

The problem I'm facing is that these assignments causes an arithmetic overflow:  
[https://android.googlesource.com/platform/external/jenkins-hash/+/75dbeadebd95869dd623a29b720678c5c5c55630/lookup3.c#181](https://android.googlesource.com/platform/external/jenkins-hash/+/75dbeadebd95869dd623a29b720678c5c5c55630/lookup3.c#181)

For instance, this is a statement I wrote:

    let a: u32 = 0xdeadbeef + (5&lt;&lt;2) + 4271423877;

and obviously rust compiler is complaining:

    error: this arithmetic operation will overflow
      --&gt; src/main.rs:34:18
       |
    34 |     let a: u32 = 0xdeadbeef + (5&lt;&lt;2) + 4271423877;
       |                  ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ attempt to add with overflow
       |
       = note: `#[deny(arithmetic_overflow)]` on by default

I tried to use \`std::num::Wrapper\` and disabling overflow checks and I still get the same error:

    rustc -Z force-overflow-checks=no src/main.rs
    
    error: this arithmetic operation will overflow
      --&gt; src/main.rs:34:18
       |
    34 |     let a: u32 = 0xdeadbeef + (5&lt;&lt;2) + 4271423877;
       |                  ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ attempt to add with overflow
       |
       = note: `#[deny(arithmetic_overflow)]` on by default

Any clue about how to fix that? Thanks!
## [12][Advanced Beginners Rust: Ownership with Threads and Closures [video]](https://www.reddit.com/r/rust/comments/g7aww2/advanced_beginners_rust_ownership_with_threads/)
- url: https://www.youtube.com/watch?v=2mwwYbBRJSo
---

