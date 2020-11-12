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
## [2][This Week in Rust 364](https://www.reddit.com/r/rust/comments/jslo80/this_week_in_rust_364/)
- url: https://this-week-in-rust.org/blog/2020/11/11/this-week-in-rust-364/
---

## [3][rkyv: a zero-copy deserialization framework for Rust](https://www.reddit.com/r/rust/comments/jss6h4/rkyv_a_zerocopy_deserialization_framework_for_rust/)
- url: https://www.reddit.com/r/rust/comments/jss6h4/rkyv_a_zerocopy_deserialization_framework_for_rust/
---
Hi everyone! I'm a long-time lurker and first poster so please be gentle. :)

I just released the first version of rkyv (*archive*), a zero-copy deserialization framework for Rust:

Source code: [https://github.com/djkoloski/rkyv](https://github.com/djkoloski/rkyv)  
Docs: [https://docs.rs/rkyv](https://docs.rs/rkyv), [https://docs.rs/rkyv\_dyn](https://docs.rs/rkyv_dyn), [https://docs.rs/rkyv\_typename](https://docs.rs/rkyv_typename)

rkyv is similar to other zero-copy deserialization frameworks like Cap'n Proto and FlatBuffers, but it's 100% pure rust and uses macro magic to build its serialization functions like serde does. The main feature is that it's zero-copy, meaning that all you have to do to "deserialize" your data is just cast a pointer. All of the data is serialized in a way that makes its in-memory representation the same as its archived representation.

rkyv sports a couple of neat features:

* Derive macros, even for complex and generic types
* \#\[no\_std\] support
* Hashmap support through a custom implementation based off of hashbrown
* Trait object serialization through the accompanying rkyv\_dyn crate. You can serialize out a trait object then use it with just a pointer cast!
* Plenty of examples and tests to make sure everything's working right.

rkyv was primarily made with an eye toward game development, where lots of static data needs to be read in and load times negatively impact player experience. Speaking from experience, deserialization takes up a big chunk of load times so a world without deserialization is a faster one!

I've been writing rust for a while but this is my first contribution to the community. If you're interested, take a look and leave me some feedback if you're interested. For example, I've only tested on Windows due to hardware constraints but if some tests are failing on other toolchains I'll find a way to get them fixed.

Thanks for taking a look!
## [4][Inside Rust Blog: Exploring PGO for the Rust compiler](https://www.reddit.com/r/rust/comments/jscfbc/inside_rust_blog_exploring_pgo_for_the_rust/)
- url: https://blog.rust-lang.org/inside-rust/2020/11/11/exploring-pgo-for-the-rust-compiler.html
---

## [5][Etebase - An open source and end-to-end encrypted Firebase alternative](https://www.reddit.com/r/rust/comments/jsa0ve/etebase_an_open_source_and_endtoend_encrypted/)
- url: https://www.reddit.com/r/rust/comments/jsa0ve/etebase_an_open_source_and_endtoend_encrypted/
---
Hey everyone, I'm Tom, the lead developer of Etebase.

The idea behind Etebase is to make an easy-to-use API to enable developers to easily build encrypted applications, and enable more privacy-first and encrypted applications to be built.

It's fully open-source, and is what powers my other project, EteSync, and its integrations with GNOME, KDE and the likes.

One thing that's really cool and extra relevant to r/rust is that almost all of our client libraries are automatically generated from our Rust implementation using flapigen!

We already support Rust (docs coming soon), though we would love to get your feedback on the API. Did we do anything wrong? Can anything be better? Would love help from from experienced Rustaceans!

Website: [https://www.etebase.com](https://www.etebase.com/)

Docs: [https://docs.etebase.com](https://docs.etebase.com/) (no Rust docs yet)

Source code: [https://github.com/etesync/etebase-rs](https://github.com/etesync/etebase-rs)

And other repos generated from Rust:

[https://github.com/etesync/libetebase](https://github.com/etesync/libetebase)

[https://github.com/etesync/etebase-py](https://github.com/etesync/etebase-py)

[https://github.com/etesync/etebase-java](https://github.com/etesync/etebase-java)
## [6][New doc comment handling in rustdoc](https://www.reddit.com/r/rust/comments/jsenc4/new_doc_comment_handling_in_rustdoc/)
- url: https://blog.guillaume-gomez.fr/articles/2020-11-11+New+doc+comment+handling+in+rustdoc
---

## [7][reacty_yew: Generating Yew components from React components via Typescript type definitions](https://www.reddit.com/r/rust/comments/jsgcui/reacty_yew_generating_yew_components_from_react/)
- url: https://www.hobofan.com/blog/2020-11-10-reacty_yew/
---

## [8][building rust without linking against libgcc](https://www.reddit.com/r/rust/comments/jst1kk/building_rust_without_linking_against_libgcc/)
- url: https://www.reddit.com/r/rust/comments/jst1kk/building_rust_without_linking_against_libgcc/
---
Hi everyone.  


I'm trying to build Rust from source via MinGW/LLVM toolchain. This toolchain doesn't contain any GNU thing.  Build fails with the message:  


`error: linking with \`x86_64-w64-mingw32-gcc\` failed: exit code: 1`

  `|`

  `= note: "x86_64-w64-mingw32-gcc" "-fno-use-linker-plugin" "-Wl,--nxcompat" "-m64" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\rsbegin.o" "-L" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.0.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.1.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.10.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.11.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.12.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.13.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.14.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.15.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.2.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.3.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.4.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.5.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.6.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.7.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.8.rcgu.o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.build_script_build.3wohcmss-cgu.9.rcgu.o" "-o" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.exe" "D:/ffbuild/rust/build/bootstrap\\debug\\build\\winapi-7e7bfc6f9a473ad4\\build_script_build-7e7bfc6f9a473ad4.5bur67vuae78v5lf.rcgu.o" "-Wl,--gc-sections" "-nodefaultlibs" "-L" "D:/ffbuild/rust/build/bootstrap\\debug\\deps" "-L" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib" "-Wl,--start-group" "-Wl,-Bstatic" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libstd-d240a4eaf263e7ee.rlib" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libpanic_unwind-532e1afcf302a81f.rlib" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libobject-9ed07451000470bb.rlib" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libaddr2line-0e3cb4dd94196903.rlib" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libgimli-f9eac35cb1c317be.rlib" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\librustc_demangle-aef25cf51f320dc4.rlib" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libhashbrown-68579e8fe92f9e82.rlib" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\librustc_std_workspace_alloc-553f19b951e308cd.rlib" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libunwind-ba4484196f7f2f72.rlib" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libcfg_if-3aff0bfb60559090.rlib" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\liblibc-250280e4fea8429b.rlib" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\liballoc-cd5549db1faddca4.rlib" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\librustc_std_workspace_core-e902813c4f3ce918.rlib" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libcore-8ef472dd7f6b9a72.rlib" "-Wl,--end-group" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libcompiler_builtins-9973e814ab9d04e7.rlib" "-Wl,-Bdynamic" "-ladvapi32" "-lws2_32" "-luserenv" "-lgcc_eh" "-l:libpthread.a" "-lmsvcrt" "-lmingwex" "-lmingw32" "-lgcc" "-lmsvcrt" "-luser32" "-lkernel32" "D:\\ffbuild\\rust\\build\\x86_64-pc-windows-gnu\\stage0\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\rsend.o"`

  `= note: lld: error: unable to find library -lgcc_eh`

`lld: error: unable to find library -lgcc`

`clang-12: error: linker command failed with exit code 1 (use -v to see invocation)`

I am not familiar with Rust ecosystem so I don't know how to remove this flag from build files. How may I build Rust without libgcc? Thanks for help.
## [9][I present you baru](https://www.reddit.com/r/rust/comments/js6ypa/i_present_you_baru/)
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
## [10][Is there a list of Rust software that doesn't exist](https://www.reddit.com/r/rust/comments/js7g6v/is_there_a_list_of_rust_software_that_doesnt_exist/)
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
## [11][Idiomatic way to implement `From&lt;T&gt;` in terms of `From&lt;&amp;T&gt;`?](https://www.reddit.com/r/rust/comments/jsa7ei/idiomatic_way_to_implement_fromt_in_terms_of_fromt/)
- url: https://www.reddit.com/r/rust/comments/jsa7ei/idiomatic_way_to_implement_fromt_in_terms_of_fromt/
---
If I have:

    impl From&lt;&amp;T&gt; for S {
        fn from(other: &amp;T) -&gt; Self {
            ...
        }
    }

And I want `From&lt;T&gt;` to be piggyback on top of that, is there any more idiomatic way to do it than writing the boilerplate:

    impl From&lt;T&gt; for S {
        fn from(other: T) -&gt; Self {
            (&amp;other).into()
        }
    }
## [12][Is there any recommended rust crawler solution?](https://www.reddit.com/r/rust/comments/js84e4/is_there_any_recommended_rust_crawler_solution/)
- url: https://www.reddit.com/r/rust/comments/js84e4/is_there_any_recommended_rust_crawler_solution/
---
Some basic features that I think are necessary:

* Post/Get methods
* Headers(UA)/Cookies setter
* css selector/xpath selector
* Asynchronous/Multithreaded

It would be better if had the following advanced features:

* Proxy pool + load balancing
* Advanced retry strategy (used to break through the crawler block)
* Database integration

Some features that I don’t need, but also very useful:

* Distributed crawler
* js execution
