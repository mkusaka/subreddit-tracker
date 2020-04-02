# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (14/2020)!](https://www.reddit.com/r/rust/comments/frfduy/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/frfduy/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fnfky9/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (14/2020)?](https://www.reddit.com/r/rust/comments/frff7k/whats_everyone_working_on_this_week_142020/)
- url: https://www.reddit.com/r/rust/comments/frff7k/whats_everyone_working_on_this_week_142020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-14-2020/40159?u=llogiq)!
## [3][swc now has a typescript / javascript node visitor which works on stable](https://www.reddit.com/r/rust/comments/ftjmhh/swc_now_has_a_typescript_javascript_node_visitor/)
- url: https://www.reddit.com/r/rust/comments/ftjmhh/swc_now_has_a_typescript_javascript_node_visitor/
---
I've implemented a Visitor and code generator which works on stable rust.

If you are building a tool related to typescript or javascript in rust, please take a look at it.,

&amp;#x200B;

The PR: [https://github.com/swc-project/swc/pull/743](https://github.com/swc-project/swc/pull/743)
## [4][This Week in Rust 332](https://www.reddit.com/r/rust/comments/ftl9l3/this_week_in_rust_332/)
- url: https://this-week-in-rust.org/blog/2020/03/31/this-week-in-rust-332/
---

## [5][Forbid pineapple on pizza](https://www.reddit.com/r/rust/comments/ft2r71/forbid_pineapple_on_pizza/)
- url: https://github.com/rust-lang/rust/pull/70645
---

## [6][Tokio: Reducing tail latencies with automatic cooperative task yielding](https://www.reddit.com/r/rust/comments/ft98nz/tokio_reducing_tail_latencies_with_automatic/)
- url: https://tokio.rs/blog/2020-04-preemption/
---

## [7][New project: Crush, a command line shell](https://www.reddit.com/r/rust/comments/ftb7fc/new_project_crush_a_command_line_shell/)
- url: https://www.reddit.com/r/rust/comments/ftb7fc/new_project_crush_a_command_line_shell/
---
I've written Crush, a shell/programming language aimed at replacing shells like bash or fish. It has some ideas in common with nushell, but Crush focuses on being a fully featured and reasonably modern programming language with types, closures, etc. 

The code (and a longer description) is available [here](https://github.com/liljencrantz/crush/), and I'd be interested in patches, feedback and opinions. The shell is fully functional and should at least work out of the box on modern Linux systems, but it's still pretty rough around the edges. Also, this has been my way of teaching myself Rust. I am already aware that I'm doing error handling wrong and that I should look into Cow strings, I will be looking into fixing those issues, but any other rust anti-patterns and mistakes I'm doing would be helpful to hear about.
## [8][(Partially) fixing a bug in a Rust research database [video]](https://www.reddit.com/r/rust/comments/ftdljb/partially_fixing_a_bug_in_a_rust_research/)
- url: https://www.youtube.com/watch?v=kiMUI0y91YI
---

## [9][Help: Parsing/Deserializing straight from the network](https://www.reddit.com/r/rust/comments/ftkcl5/help_parsingdeserializing_straight_from_the/)
- url: https://www.reddit.com/r/rust/comments/ftkcl5/help_parsingdeserializing_straight_from_the/
---
I'm trying to implement a network protocol that makes use of some weird decisions, for example I am trying to deserialize the following.

    struct Foo {
      length: VarInt,
      data: Vec&lt;u8&gt;
    }

Where VarInt is an i32 represented between 1 and 5 bytes with the LSB indicating whether or not to read another byte.

Serde and Nom seem like the obvious choices but Serde appears to be lacking a way of pulling bytes off the line one at a time (for parsing `VarInt`) or pulling N at once without consuming a length prefix.

Nom only seems to want to take byte slices rather than anything that implements the Read trait.

Is this something that would need to be implemented by hand or is there another crate better suited to this that I've missed?
## [10][Why don't the format and println macros use the f-string syntax from Python?](https://www.reddit.com/r/rust/comments/ftkse0/why_dont_the_format_and_println_macros_use_the/)
- url: https://www.reddit.com/r/rust/comments/ftkse0/why_dont_the_format_and_println_macros_use_the/
---
I've recently started learning Rust. As far as I understand, macros work by actually parsing and analysing the tokens given to the macro. Then I wonder, why does the format and print macros use this syntax:

```println!("The person {} is {} years old", name, age);```

Rather than this syntax, inspired by f-strings from Python:

```println!("The person {name} is {age} years old");```

Is there something about the way macros work that prevents this kind of syntax? Would it be possible to make a macro which works this way?

Personally I find the second syntax much more readable, especially if there are many arguments to the format string. Why was the first syntax chosen?
## [11][What happened to the 2019 State of Rust Survey results?](https://www.reddit.com/r/rust/comments/ftch5o/what_happened_to_the_2019_state_of_rust_survey/)
- url: https://www.reddit.com/r/rust/comments/ftch5o/what_happened_to_the_2019_state_of_rust_survey/
---
The survey was closed on Dec 16th 2019 and the results were supposed to be released "a month or so afterwards" but there is no mention of it on the blog.  What happened?
## [12][kmon: Linux Kernel Manager and Activity Monitor written in Rust](https://www.reddit.com/r/rust/comments/fsz7ef/kmon_linux_kernel_manager_and_activity_monitor/)
- url: https://www.reddit.com/r/rust/comments/fsz7ef/kmon_linux_kernel_manager_and_activity_monitor/
---
[kmon](https://github.com/orhun/kmon) provides a text-based user interface for managing the Linux kernel modules and monitoring the kernel activities. By managing, it means loading, unloading, blacklisting and showing the information of a module. These updates in the kernel modules, logs about the hardware and other kernel messages can be tracked with the real-time activity monitor in kmon. Since the usage of different tools like dmesg and kmod are required for these tasks in Linux, kmon aims to gather them in a single terminal window and facilitate the usage as much as possible while keeping the functionality.

kmon is written in Rust and uses [tui-rs](https://github.com/fdehau/tui-rs) &amp; [termion](https://github.com/redox-os/termion) libraries for its text-based user interface.

[kmon on action](https://i.redd.it/frrmmacme7q41.gif)

**Project Homepage:** [https://github.com/orhun/kmon](https://github.com/orhun/kmon)  
**Rust Package:** [https://crates.io/crates/kmon](https://crates.io/crates/kmon)
