# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (24/2020)!](https://www.reddit.com/r/rust/comments/gyswpo/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/gyswpo/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/guo4c1/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (24/2020)?](https://www.reddit.com/r/rust/comments/gysxgq/whats_everyone_working_on_this_week_242020/)
- url: https://www.reddit.com/r/rust/comments/gysxgq/whats_everyone_working_on_this_week_242020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-24-2020/43988?u=llogiq)!
## [3][Are many companies using Rust outside of the US?](https://www.reddit.com/r/rust/comments/gyw3f5/are_many_companies_using_rust_outside_of_the_us/)
- url: https://www.reddit.com/r/rust/comments/gyw3f5/are_many_companies_using_rust_outside_of_the_us/
---
I live in Australia and haven't really seen any companies that mention they are using Rust in their products. 

From what I can tell, the vast majority of Rust jobs are based in the US (San Francisco, New York, etc.). Do you know of any companies outside of America that employ Rust programmers?
## [4][Absolute addresses in position-independent code (PIC)](https://www.reddit.com/r/rust/comments/gyrmam/absolute_addresses_in_positionindependent_code_pic/)
- url: https://www.reddit.com/r/rust/comments/gyrmam/absolute_addresses_in_positionindependent_code_pic/
---
Cross-posted to [StackOverflow](https://stackoverflow.com/questions/62254811/absolute-addresses-in-position-independent-code-pic) .

I am trying to build and link a single image to load as an OS kernel (ie. in QEMU) targeting aarch64-unknown-none-softfloat. I use a custom linker.ld file which sets the entry point for the kernel `ENTRY(_reset)` and positions the image

`. = 0x40080000`

where the program counter (PC) is on reset.

It works ok until I map the pages at 0x40080000 to high memory where the kernel will reside and enable virtual memory translation. To ensure the debugging information meshes after the switch, I change the nominal image position to

`. = 0xffffff8200000000`

and rebuild.

I have discovered that access:

* to some (pub extern) statics, and
* by certain core library functions

is by reading the absolute address from somewhere in `.rodata`. This breaks the code when it is running before mapping. And if I change it back it will break the code when I run it after mapping.

The code it is generating looks a bit like this at O1 (indirect through PC-relative page):

    adrp  x0, 0x10000 // page offset from PC up to rodata
    add   x0, 0x120   // byte offset from page in rodata
    ldr   x0, [x0]    // use as address

What I need is truly position independent code across code and data so that it works at *both* locations in memory without referring to any stored absolute addresses, even if those addresses are available relative to PC.

I've tried the other relocation-models including Pic and RopiRwpi but I can't see it generating different code.

Thanks!
## [5][Announcing fluent-templates: Easily Add Fluent Localisation To Your Rust Projects.](https://www.reddit.com/r/rust/comments/gyxl2g/announcing_fluenttemplates_easily_add_fluent/)
- url: https://github.com/XAMPPRocky/fluent-templates
---

## [6][What is the state of rust on the ESP8266?](https://www.reddit.com/r/rust/comments/gyxbkv/what_is_the_state_of_rust_on_the_esp8266/)
- url: https://www.reddit.com/r/rust/comments/gyxbkv/what_is_the_state_of_rust_on_the_esp8266/
---
There are a few older posts about this but there is a bunch of activity in the area since then and it looks like its at least possible to blink leds now. Has anyone tried to use rust here? What is it like? Are we at the point of being able to connect to wifi and make a http request?
## [7][GTK and Rust](https://www.reddit.com/r/rust/comments/gytgw8/gtk_and_rust/)
- url: https://www.reddit.com/r/rust/comments/gytgw8/gtk_and_rust/
---
I saw on the oficial GTK page on language bindings the support of Rust, but is developed for other people, not directly by Gnome, so this library port are developed by gnome too and have a really interest of moving forward or is only a honorable mention like alternative for rust language?
## [8][A Quick (and good) View of Iterators in Rust!](https://www.reddit.com/r/rust/comments/gyryth/a_quick_and_good_view_of_iterators_in_rust/)
- url: https://youtu.be/HZftwxCIXqE
---

## [9][This Month in Rust OSDev (May 2020)](https://www.reddit.com/r/rust/comments/gyi9p0/this_month_in_rust_osdev_may_2020/)
- url: https://rust-osdev.com/this-month/2020-05/
---

## [10][This month in rustsim #11 (April - May 2020): cross-platform deterministic physics using nphysics with fixed-point numbers!](https://www.reddit.com/r/rust/comments/gyeaik/this_month_in_rustsim_11_april_may_2020/)
- url: https://www.rustsim.org/blog/2020/06/01/this-month-in-rustsim/
---

## [11][[Ann] gemini a crate to make sync and async apis while letting your users choose what they want](https://www.reddit.com/r/rust/comments/gyl7h7/ann_gemini_a_crate_to_make_sync_and_async_apis/)
- url: https://github.com/mgattozzi/gemini
---

## [12][MiniCouchDB: implementing a subset of CouchDB in Rust](https://www.reddit.com/r/rust/comments/gynj0f/minicouchdb_implementing_a_subset_of_couchdb_in/)
- url: https://www.garrensmith.com/blogs/mini-couch-hack-week
---

