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
## [3][SQLx 0.3 released! Now with support for SQLite, zero-copy/amortized-allocation row decoding, nullable columns, user-defined types and LISTEN/NOTIFY in Postgres, plus a bunch more data types!](https://www.reddit.com/r/rust/comments/frww6f/sqlx_03_released_now_with_support_for_sqlite/)
- url: https://github.com/launchbadge/sqlx/blob/master/CHANGELOG.md#030---2020-03-29
---

## [4][How does Rust or the LLVM optimize code so well?](https://www.reddit.com/r/rust/comments/fsbwya/how_does_rust_or_the_llvm_optimize_code_so_well/)
- url: https://www.reddit.com/r/rust/comments/fsbwya/how_does_rust_or_the_llvm_optimize_code_so_well/
---
[Solved] I'm an idiot and didn't use the variabels. LLVM optimized and removed them.

For my school I'm doing research on Rust.
I've been comparing the performance of Rust and Java.
I'm doing this by creating a few test applications and comparing the execution time and memory usages.

One of the applications is calculating the Fibonacci sequence in [Rust](https://github.com/SemtexError/rust-fibonacci/blob/master/src/main.rs) and in [Java](https://github.com/SemtexError/java-fibonacci/blob/master/src/com/semtex/Main.java).
The execution time is calculated with [Hyperfine](https://github.com/sharkdp/hyperfine) and adding a random number to the sequence so LLVM can't optimize it.

After compiling and running both applications I noticed a big difference.
Calculating the sequence once will give the follow result.

Language|Time (ms)
:--|:--
Java|139.9
Rust|0.8

I thought maybe the JVM, the GC will take some time to start up and the JIT need some time to optimize the code.
So we calculate the sequence 1000 times within a single application.

Language|Time (ms)
:--|:--
Java|29397
Rust|1.8

This is a huge difference.
I would expect Rust to be way faster but this difference is huge.
The un-optimized Rust code has comparable results to Java.

How can Rust or the LLVM optimize this code so well?
How can it be this much faster then other languages?

Can anyone of you help me answer this question or point me in the right direction?

PS: English is not my native language so sorry for that.

Edit: Table didn't mark like normal markdown.
## [5][so long rust, why i built ZZ](https://www.reddit.com/r/rust/comments/fsc6p7/so_long_rust_why_i_built_zz/)
- url: http://aep.github.io/from-rust-to-zz/
---

## [6][Reducing memory consumption in librsvg, part 4: compact representation for Bézier paths](https://www.reddit.com/r/rust/comments/fs24cq/reducing_memory_consumption_in_librsvg_part_4/)
- url: https://people.gnome.org/~federico/blog/reducing-memory-consumption-in-librsvg-4.html
---

## [7][embedded Rust: writing to file on SD card?](https://www.reddit.com/r/rust/comments/fsaiic/embedded_rust_writing_to_file_on_sd_card/)
- url: https://www.reddit.com/r/rust/comments/fsaiic/embedded_rust_writing_to_file_on_sd_card/
---
fellow Rustaceans,

As I'm learning #rustlang on embedded devices, I'm trying to port some of my earlier MicroPython project to embedded Rust. 

I have a simple embedded (STM32) contraption that measures the temperature and displays it on an OLED display: [https://github.com/nebelgrau77/stm32f411\_TMP36](https://github.com/nebelgrau77/stm32f411_TMP36) 

 I would like to log these readings to a simple text file on an SD Card, to retrieve them later and use for some data analysis in Python. In MicroPython that was pretty simple, I would just create a new line with comma-separated values, and append to file on the card. 

I know there are some crates for SD cards in no\_std Rust environment, such as  [**embedded-sdmmc-rs**](https://github.com/thejpster/embedded-sdmmc-rs) and  [**rust-fatfs**](https://github.com/rafalh/rust-fatfs). What I need is a simple example, and the examples in the crates themselves are a bit confusing. While I'm trying to figure it out, does anyone have an example of a working code that writes to a file in embedded no\_std application?
## [8][Rust auto use in Vim ?](https://www.reddit.com/r/rust/comments/fsbje1/rust_auto_use_in_vim/)
- url: /r/vim/comments/fsbj5i/rust_auto_use_in_vim/
---

## [9][How to iterate trees nicely?](https://www.reddit.com/r/rust/comments/fsbqwp/how_to_iterate_trees_nicely/)
- url: https://www.reddit.com/r/rust/comments/fsbqwp/how_to_iterate_trees_nicely/
---
Whenever I'm writing trees, or any similar recursive structure, I repeatedly run into a problem when I want to iterate over the leaves of the tree. The "obvious" way to find the leaves recursively like this:

    fn iter(&amp;self) -&gt; impl Iterator&lt;Item = &amp;ItemType&gt; {
            match &amp;self {
                Leaf(items) =&gt; items.iter(),
                Children(c) =&gt; { c.iter().map(Deref:deref).flat_map(Self::iter) }
            }
    }

Doesn't work because the return type can't be made consistent and finite. It does work if you box everything:

    fn iter&lt;'a&gt;(&amp;'a self) -&gt; Box&lt;dyn Iterator&lt;Item = &amp;'a ItemType&gt; + 'a&gt; {
        match &amp;self {
            Leaf(items) =&gt; Box::new(items.iter()),
            Children(c) =&gt; Box::new(c.iter().map(Deref::deref).flat_map(Self::iter)),
        }
    }

(For context, `c` in the Children variant is `[Box&lt;Self&gt;, _]`, so the Deref call is there to unbox everything. I also think the second example shouldn't work because we end up with one too many layers of iterator, but the compiler lets it pass)

This second way is workable, but feels inefficient - if I ask for all the leaves in the tree, I'd create one box per tree node, when in principle I only need O(depth) space. Is there any way to do this more efficiently, perhaps with generators, that doesn't go so far that I need to write a completely non-recursive iterator from scratch? (I can imagine recursive iterators might somehow work, but can't get my head around how to do the allocations efficiently)

I've also ran into this problem dealing with "trees" that are implicitly generated in constraint solving and the like, where we take a collection of states and apply a function to produce all the "next" states possible from a given starting point, which might be as few as zero. Repeatedly flat\_mapping this function creates an iterator that yields the "leaves"  of an implicitly calculated tree search... except the call to flat\_map is very awkward because the iterator changes static type each iteration. Is this possible to do without creating lots of intermediate boxes? Or is it possible to replace a `mut Box&lt;dyn _&gt;`'s contents with a completely different type, so long as it implements the same trait?
## [10][Article about the DW1000 (IEEE 802.15.4-based radio transceiver) driver I wrote in Rust](https://www.reddit.com/r/rust/comments/fs9q54/article_about_the_dw1000_ieee_802154based_radio/)
- url: https://braun-embedded.com/dw1000/
---

## [11][More advanced aspects of pattern matching in Rust](https://www.reddit.com/r/rust/comments/fs0jg8/more_advanced_aspects_of_pattern_matching_in_rust/)
- url: https://notes.iveselov.info/programming/refs-and-pattern-matching-in-rust
---

## [12][rust-analyzer Changelog #18](https://www.reddit.com/r/rust/comments/frrboc/rustanalyzer_changelog_18/)
- url: https://rust-analyzer.github.io/thisweek/2020/03/30/changelog-18.html
---

