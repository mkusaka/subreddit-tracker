# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (2/2020)!](https://www.reddit.com/r/rust/comments/ekpqp7/hey_rustaceans_got_an_easy_question_ask_here_22020/)
- url: https://www.reddit.com/r/rust/comments/ekpqp7/hey_rustaceans_got_an_easy_question_ask_here_22020/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

The Rust-related IRC channels on irc.mozilla.org (click the links to open a web-based IRC client):

 - [#rust](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust) (general questions)
 - [#rust-beginners](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-beginners) (beginner questions)
 - [#cargo](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23cargo) (the package manager)
 - [#rust-gamedev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-gamedev) (graphics and video games, and see also [/r/rust_gamedev](https://www.reddit.com/r/rust_gamedev))
 - [#rust-osdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-osdev) (operating systems and embedded systems)
 - [#rust-webdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-webdev) (web development)
 - [#rust-networking](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-networking) (computer networking, and see also [/r/rust_networking](https://www.reddit.com/r/rust_networking))

Also check out [last week's thread](https://reddit.com/r/rust/comments/ehk18j/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What’s everyone working on this week (2/2020)?](https://www.reddit.com/r/rust/comments/ekpr6w/whats_everyone_working_on_this_week_22020/)
- url: https://www.reddit.com/r/rust/comments/ekpr6w/whats_everyone_working_on_this_week_22020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-2-2020/36546?u=llogiq)!
## [3][Let The Compiler Do The Work](https://www.reddit.com/r/rust/comments/eksduv/let_the_compiler_do_the_work/)
- url: http://cliffle.com/p/dangerust/6/
---

## [4][tomaka/redshirt: an OS prototype where binaries are WASM, run in ring 0](https://www.reddit.com/r/rust/comments/ekingn/tomakaredshirt_an_os_prototype_where_binaries_are/)
- url: https://github.com/tomaka/redshirt
---

## [5][Is anyone concerned about this deep, deep nesting of dependencies for basic web functionality in Rust?](https://www.reddit.com/r/rust/comments/ekpa3i/is_anyone_concerned_about_this_deep_deep_nesting/)
- url: https://www.reddit.com/r/rust/comments/ekpa3i/is_anyone_concerned_about_this_deep_deep_nesting/
---
Today, I wanted to know what it would take to issue a basic HTTP request using \`reqwest\`, the de-facto standard library:

    cargo new with_reqwest
    cd with_reqwest
    echo 'reqwest = "*"' &gt;&gt; Cargo.toml
    cargo build

This built 97 crates.

I tried another one with \`scraper\`, to scape HTML. 95 crates.

For basic manipulation of JSON, using \`serde\` and \`serde\_json\`. 18 crates.

That's a lot of dependencies. Are there any potential issues this could cause? Is anyone worried about this?
## [6][Announcing task_scope](https://www.reddit.com/r/rust/comments/ekqjih/announcing_task_scope/)
- url: https://docs.rs/task_scope/0.1.0/task_scope/
---

## [7][Actix Web: Optimization Amongst Optimizations](https://www.reddit.com/r/rust/comments/ekshih/actix_web_optimization_amongst_optimizations/)
- url: https://brandur.org/nanoglyphs/008-actix
---

## [8][Official announcement: glsl-4.0](https://www.reddit.com/r/rust/comments/eklo6l/official_announcement_glsl40/)
- url: https://www.reddit.com/r/rust/comments/eklo6l/official_announcement_glsl40/
---
[glsl-4.0](https://crates.io/crates/glsl/4.0.0) and [glsl-quasiquote](https://crates.io/crates/glsl-quasiquote/4.0.0) just got released.

&gt; `glsl` is a [GLSL](https://en.wikipedia.org/wiki/OpenGL_Shading_Language) parsing, syntax and transformation crate written in pure Rust. Its goals are to parse valid _GLSL450/460_ sources. `glsl-quasiquote` is a companion crate that is to `glsl` what `quote` is to Rust.

Among the interesting features, `glsl` got support for two _backward-compatible_ keyword: `attribute` and `varying`. Those are not part of the _GLSL450_ but adding them would allow to parse older versions of GLSL for free.

Another important part is that prior to 4.0, the `glsl` crate parsed binary operations as right-associative, which is not what the spec states (they should be left-associative). A patch has landed and this is now fixed. Thanks to [@jrmuizel](https://github.com/jrmuizel) for his commitment and help on that issue.

Keep the vibes!
## [9][How do you get project ideas and motivation for coding in your spare time?](https://www.reddit.com/r/rust/comments/eknvoq/how_do_you_get_project_ideas_and_motivation_for/)
- url: https://www.reddit.com/r/rust/comments/eknvoq/how_do_you_get_project_ideas_and_motivation_for/
---
I am 25 and I've been coding for 6 years now. I found out, that coding is something I really like and I think I am/have the potential to be good at it. While trying to find a solution for a concrete problem, I often think of how to generalize it and try to find a solution for not the problem itself, but for the type of problems in a more general way. Sometimes this thing happens on the next (more general) level again and again and I catch myself trying to find the most generic solution possible, which can be very counterproductive for solving the actual problem.

Because of that I figured that I would love to create libraries/frameworks/APIs instead of end user products. I love the feeling of starting a fresh new project with just a main, planning and creating everything from scratch. When I started learning how to code I always came to the point where I trashed the whole project and started it all over again, because I realized, that I messed up the structure and learned how to make it better until I realized, that I messed up again...

There are days, where I get really excited about programming (just thinking about it) and feel a strong urge to code and be productive. Its an energizing feeling, that I can't really describe and it comes over me just some days randomly, when I'm somehow in the mood. That feeling gets strong especially when learning something knew in the field of coding (reading the Rust book for example always gave me that feeling). Does anyone know what I mean?

Anyway, I think I unfortunately am a really uncreative person, so I don't have many ideas of what to do. I used to program simple 2d games like tetris, which was fun and really instructive, but I wanna create something that is useful. I need a project, that I can work on when I get in the above described mood. In the rare cases of having a good project idea, I don't really know if it already exists and how to meet all the rules and conventions in general.

I did some of those little programming exercises like adventOfCode or rustlings. Those little exercises are good to practice and get to know the basics of the language (syntax, standard library...), but are not really what I am looking for. I also thought of contributing to an existing project (Rust standard library or an existing crate), but every time I try to find something to contribute to, I feel overwhelmed by the things/rules that are necessary to follow and the reading into the existing code. For some reason I don't really know where to start and give up. I am not even aware of what exactly is stopping me.

I also really find it hard to find people, that share those interests. I study IT and work as a programmer, but I feel like most of the people around there just do IT, because of the career possibilities and not because they really are interested in that kinda stuff.

Does anyone have any suggestion on how to find a project to work on? How do you get the motivation or ideas for coding for fun? What do you code when you code in your free time and (how) are the projects different from the ones you work on at your job?
## [10][Can anybody explain these differences in test results between platforms?](https://www.reddit.com/r/rust/comments/ekt3ox/can_anybody_explain_these_differences_in_test/)
- url: https://www.reddit.com/r/rust/comments/ekt3ox/can_anybody_explain_these_differences_in_test/
---
I recently created [a PR](https://github.com/tov/libffi-sys-rs/pull/29) for `libffi-sys` to support MSVC building. My next step is [a PR](https://github.com/tov/libffi-rs/pull/12) for the `libffi` crate to build and test with MSVC too via Github Actions.

After playing CI whack-a-mole for a while I finally got it mostly working and found a single doc test [failing on Windows only](https://github.com/timfish/libffi-rs/runs/373201160):

    extern "C" fn add(x: f64, y: &amp;f64) -&gt; f64 {
        return x + y;
    }
    
    use libffi::middle::*;
    
    let args = vec![Type::f64(), Type::pointer()];
    let cif = Cif::new(args.into_iter(), Type::f64());
    
    let n = unsafe { cif.call(CodePtr(add as *mut _), &amp;[arg(&amp;5), arg(&amp;&amp;6)]) };
    assert_eq!(11, n);

Result:

    ---- src\middle\mod.rs - middle::Cif (line 71) stdout ----
    Test executable failed (exit code 101).
    
    stderr:
    thread 'main' panicked at 'assertion failed: `(left == right)`
      left: `11`,
     right: `6`', src\middle\mod.rs:14:1

I did some dirty `println!` debugging in the `add` function and found that at least on Windows, the `x: f64, y: &amp;f64` arguments were not 5 and 6 as expected. I quickly realised that there is no type safety in these ffi calls... I made the following change to the arguments to ensure that the numbers being passed were in fact `f64`:

    let n = unsafe { cif.call(CodePtr(add as *mut _), &amp;[arg(&amp;5f64), arg(&amp;&amp;6f64)]) };

Why was this test passing on macOS and Linux but failing on Windows?

After making this change, this test [now fails for the same reason on every platform](https://github.com/timfish/libffi-rs/runs/374800102), which I'll take as a partial win, but I have no idea how to fix it and it seems like I should probably understand what's going on here before I precede!

    ---- src\middle\mod.rs - middle::Cif (line 71) stdout ----
    Test executable failed (exit code 101).
    
    stderr:
    thread 'main' panicked at 'assertion failed: `(left == right)`
      left: `11`,
     right: `0`', src\middle\mod.rs:14:1
## [11][Performance of Rust when used with Java and in smartphone development](https://www.reddit.com/r/rust/comments/ekpzyw/performance_of_rust_when_used_with_java_and_in/)
- url: https://www.reddit.com/r/rust/comments/ekpzyw/performance_of_rust_when_used_with_java_and_in/
---
Hello, 

I am working on a project that requires a lot of hashing, encrypting and decrypting, as well as some networking. Since these operations can be computationally intense I was considering writing my library in Rust. 

I found a question on stack overflow where someone made a [.dll file from his Rust](https://stackoverflow.com/questions/30258427/calling-rust-from-java) code and used it in java as well as an article on [medium]( https://link.medium.com/ZcnTOMXk12) explaining how to use Rust in android and iOS dev. 


Another pro is that I could use the same code base for all the devices, that would gain me some time. 

Do you think it is worth learning Rust and write the API/library in it to fasten the app ?
## [12][Announcing the v0.11 release of Yew](https://www.reddit.com/r/rust/comments/ektzvd/announcing_the_v011_release_of_yew/)
- url: https://www.reddit.com/r/rust/comments/ektzvd/announcing_the_v011_release_of_yew/
---
Hello, I'm excited to share with all of you the latest Yew release :)

*(If you're not familiar, Yew is a framework for building client web apps with Rust &amp; WebAssembly!)*

The latest release contains quite a few breaking changes and has a [transition guide](https://github.com/yewstack/yew/blob/0.11.0/CHANGELOG.md#transition-guide) for projects that need to be migrated. The reason there are a number of breaking changes is because we are preparing for a 1.0 release of Yew in the near future. Additionally, there has been a lot of interest from the Yew community in building and using component libraries built on Yew and so this release adds some functionality to support that effort. For more info, check out the [full changelog](https://github.com/yewstack/yew/releases).

Also, I'd like to announce Yew's new documentation website that aims to help developers learn about Yew and start building Yew apps more quickly. You can find it here: [https://yew.rs](https://yew.rs). *Note that it's still a WIP, PR's welcome ;)*

For the next release, we will be adding support for `web-sys` (while preserving support for `stdweb` and Emscripten targets) and will start integrating with the [gloo](https://github.com/rustwasm/gloo) toolkit. Also in the works.. we have folks working on CSS tooling, component libraries, and server side rendering. Come hang out in our [Gitter chat](https://gitter.im/yewframework/Lobby) if you'd like to learn more!
