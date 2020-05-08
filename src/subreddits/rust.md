# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (19/2020)!](https://www.reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 337](https://www.reddit.com/r/rust/comments/geagy0/this_week_in_rust_337/)
- url: https://this-week-in-rust.org/blog/2020/05/05/this-week-in-rust-337/
---

## [3][grout | A very simple tiling window manager for Windows](https://www.reddit.com/r/rust/comments/gfieg2/grout_a_very_simple_tiling_window_manager_for/)
- url: https://www.reddit.com/r/rust/comments/gfieg2/grout_a_very_simple_tiling_window_manager_for/
---
[https://github.com/tarkah/grout](https://github.com/tarkah/grout)

demo: [https://i.imgur.com/ugPMvlA.mp4](https://i.imgur.com/ugPMvlA.mp4)

&amp;#x200B;

I love Budgie's Window Shuffler grid functionality for quickly and easily tiling windows and I really wanted to have this on my Window's machine, so I decided to take a stab at building it. Going into this I knew basically nothing about the Windows API and overall this program feels super hacky, but it works! 

Hope others will find this as useful as me. I've currently only tested it on a single monitor, but it shouldn't be hard to get it working over any number of them. In the coming days, I plan to add a multi-monitor support, a system tray icon, simple settings page, and the ability to autostart, all of which should make it as familiar to use as any windows program.
## [4][An overview of foreign language interop libraries for Rust](https://www.reddit.com/r/rust/comments/gfp7dd/an_overview_of_foreign_language_interop_libraries/)
- url: https://www.hobofan.com/rust-interop
---

## [5][Announcing Rust 1.43.1](https://www.reddit.com/r/rust/comments/gf8om1/announcing_rust_1431/)
- url: https://blog.rust-lang.org/2020/05/07/Rust.1.43.1.html
---

## [6][Benchmarking slotmap, slab, stable_vec etc.](https://www.reddit.com/r/rust/comments/gfo1uw/benchmarking_slotmap_slab_stable_vec_etc/)
- url: https://www.reddit.com/r/rust/comments/gfo1uw/benchmarking_slotmap_slab_stable_vec_etc/
---
I made a small and simple implementation of a slotmap/slab whatever you want to call it. I wanted to see how the performance would be if you use a union instead of an enum for the slots and store the tag (occupied or not) in a bitvec. I did this because of curiosity, not because I actually needed performance. Then I started adding benchmarks to compare with existing implementations of such a container. I've added all the competitors that I could find on [crates.io](https://crates.io). The benchmark is probably of limited usefulness, anybody choosing which of these to use should probably make initial choice based on API/features/use of unsafe, etc.

The benchmarking code is available at [github](https://github.com/spersson/bvmap/). My prototype is called BvMap. All times in microseconds.

**Insert 10 000 usize into newly created container**

1. ExternStableVec	49.10
2. IdVec	56.77
3. BvMap	65.18
4. BeachMap	67.66
5. CompactMap	88.38
6. Slab	92.69
7. Stash	93.26
8. InlineStableVec	114.25
9. SlotMap	118.23
10. UniqueStash	126.58
11. HopSlotMap	156.29
12. DenseSlotMap	177.05
13. Froggy	222.66

**Remove 10 000 usize**

1. CompactMap	14.11
2. InlineStableVec	14.16
3. ExternStableVec	14.63
4. Slab	15.60
5. Stash	17.32
6. UniqueStash	25.98
7. BvMap	37.99
8. SlotMap	40.92
9. BeachMap	52.83
10. HopSlotMap	61.23
11. DenseSlotMap	69.14
12. Froggy	201.56
13. IdVec	465.20

**Re-insert 10 000 usize**

1. UniqueStash	28.30
2. Slab	29.35
3. Stash	29.72
4. BvMap	32.34
5. SlotMap	57.70
6. BeachMap	59.10
7. HopSlotMap	63.68
8. CompactMap	71.25
9. ExternStableVec	103.13
10. InlineStableVec	134.73
11. DenseSlotMap	151.28
12. Froggy	213.28
13. IdVec	5234.30

**Get 10 000 random items from collection of size 10 000**

1. ExternStableVec	169.28
2. Slab	177.15
3. BvMap	177.70
4. Stash	178.64
5. CompactMap	181.87
6. DenseSlotMap	186.79
7. InlineStableVec	190.80
8. SlotMap	210.83
9. HopSlotMap	217.73
10. UniqueStash	223.76
11. BeachMap	241.17
12. IdVec	316.86
13. Froggy	362.39

**Iterate over all 10 000 items**

1. ExternStableVec	10.85
2. Stash	11.10
3. InlineStableVec	12.02
4. CompactMap	12.71
5. DenseSlotMap	12.75
6. Slab	13.04
7. BeachMap	13.94
8. UniqueStash	16.40
9. BvMap	20.47
10. HopSlotMap	27.82
11. SlotMap	29.38
12. IdVec	158.04
13. Froggy	164.22

**Iterate over 5 000 items after 5 000 random items have been removed**

1. DenseSlotMap	6.48
2. BeachMap	16.72
3. CompactMap	34.60
4. Slab	35.24
5. Stash	35.41
6. HopSlotMap	35.61
7. InlineStableVec	35.87
8. ExternStableVec	37.17
9. UniqueStash	38.51
10. SlotMap	43.56
11. BvMap	49.79
12. Froggy	119.40
13. IdVec	212.02

&amp;#x200B;

**Observations**

Not much. IdVec stands out for its bad performance while it doesn't seem to offer any advantage to motivate it (although the idea to return Id&lt;T&gt; index type is interesting). The thing that caught my eye was that iterating over a full container was such a difference between UniqueStash and SlotMap. They have almost identical implementations but UniqueStash was almost double speed. Could be something for slotmap author to look into.
## [7][Rust-like language with static memory management and no ownership](https://www.reddit.com/r/rust/comments/gfgt1b/rustlike_language_with_static_memory_management/)
- url: https://www.reddit.com/r/rust/comments/gfgt1b/rustlike_language_with_static_memory_management/
---
[https://github.com/doctorn/micro-mitten](https://github.com/doctorn/micro-mitten)

This may be of interest - a fragment of Rust developed as a research language to support static memory management without ownership. The compiler for the language is also written in Rust!
## [8][Verco | A simple Git/Hg tui client focused on keyboard shortcuts](https://www.reddit.com/r/rust/comments/gfspio/verco_a_simple_githg_tui_client_focused_on/)
- url: https://www.reddit.com/r/rust/comments/gfspio/verco_a_simple_githg_tui_client_focused_on/
---
https://github.com/matheuslessarodrigues/verco

After many years being frustrated with slow gui clients for git/hg, I switched to using them from the command line.
However, I found myself also frustrated with having to type the same mundane commands everytime (even with aliases).
Also, having used SourceTree for some time, I liked the idea of a single interface for both git and mercurial.

So here I am now after writing my own tui client. My workflow is leaving verco open and when I need to quickly
push, pull, commit, update/checkout, I just need to type a single key (two at most) to execute a command.

Some keybindings: `p`: pull, `P`: push, `cc`: commit all, `u`: update/checkout, more at the repository page.

Verco has been my daily driver for quite some time and now I think it's stable enough to share with the internet.
Hope you find it useful as well :) Thanks!
## [9][Google Summer of Code 2020 project to integrate Pathfinder into Inkscape](https://www.reddit.com/r/rust/comments/gffhnx/google_summer_of_code_2020_project_to_integrate/)
- url: https://summerofcode.withgoogle.com/projects/#5859756641615872
---

## [10][Cow in a Box &amp; Friends [Rust Wrocław Webinar, 29.04.2020]](https://www.reddit.com/r/rust/comments/gfpefn/cow_in_a_box_friends_rust_wrocław_webinar_29042020/)
- url: https://www.youtube.com/watch?v=EjbfoOCL2to
---

## [11][color-eyre: a custom context for eyre which provides colorful error reports and backtraces on stable is now available on crates.io 🎉🥳🎉](https://www.reddit.com/r/rust/comments/gf6z3x/coloreyre_a_custom_context_for_eyre_which/)
- url: https://i.imgur.com/ZDVN9sf.png
---

## [12][v0.8 of Gleam, a statically typed language written in Rust for the Erlang VM, is out](https://www.reddit.com/r/rust/comments/gfe4n7/v08_of_gleam_a_statically_typed_language_written/)
- url: https://lpil.uk/blog/gleam-v0.8-released/
---

