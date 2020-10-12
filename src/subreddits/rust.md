# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (42/2020)!](https://www.reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (42/2020)?](https://www.reddit.com/r/rust/comments/j9l0o7/whats_everyone_working_on_this_week_422020/)
- url: https://www.reddit.com/r/rust/comments/j9l0o7/whats_everyone_working_on_this_week_422020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-42-2020/49997?u=llogiq)!
## [3][Proving that 1 + 1 = 2 in Rust](https://www.reddit.com/r/rust/comments/j9nnpv/proving_that_1_1_2_in_rust/)
- url: https://gist.github.com/gretingz/bc194c20a2de2c7bcc0f457282ba2662
---

## [4][Rust after the honeymoon](https://www.reddit.com/r/rust/comments/j99o0t/rust_after_the_honeymoon/)
- url: http://dtrace.org/blogs/bmc/2020/10/11/rust-after-the-honeymoon/
---

## [5][jlrs v0.7: write functions in Rust and call them from Julia, execute multiple tasks in parallel, use Julia arrays with ndarray](https://www.reddit.com/r/rust/comments/j9n0ee/jlrs_v07_write_functions_in_rust_and_call_them/)
- url: https://www.reddit.com/r/rust/comments/j9n0ee/jlrs_v07_write_functions_in_rust_and_call_them/
---
Last Saturday I released jlrs v0.7 which has two major new features.

First of all, you can now use functionality from jlrs when writing Rust code that is called from Julia using its `ccall` interface. This function can call arbitrary functions defined in dynamic libraries that use the C ABI (ie functions marked with `extern "C"`, and possibly `#[no_mangle]` in Rust). While this can be achieved without using jlrs, it's now possible to use features from jlrs such as borrowing array data.

The second major feature is the introduction of an (experimental) async runtime. This feature is highly experimental and only works on Linux. With the async runtime you can use `Value::call_async` to offload a function call to another thread (using `Base.Threads.@spawn`) and wait for the result; while the computation is running the runtime is free to work on other tasks. This feature is experimental because it's the first async code I've ever written and depends on some unstable features.

Additionally, some small features have been added. The biggest one of these is that there's a new arrya type, `TypedArray`, which is highly similar to `Array` except that the type of the elements is explicitly known. The crate `jlrs-ndarray` is now available to easily borrow array data from Julia as an `ArrayView(Mut)` from ndarray.

The repository now contains several examples, so if you're curious how to use these new features you can take a look [here](https://github.com/Taaitaaiger/jlrs/tree/master/examples).

[Github](https://github.com/Taaitaaiger/jlrs)

[Docs.rs](https://docs.rs/jlrs/0.7.0/jlrs/index.html)

[Crates.io](https://crates.io/crates/jlrs)
## [6][rust-analyzer changelog #46](https://www.reddit.com/r/rust/comments/j9q5hn/rustanalyzer_changelog_46/)
- url: https://rust-analyzer.github.io/thisweek/2020/10/12/changelog-46.html
---

## [7][Have you stopped using Rust? And if yes, why?](https://www.reddit.com/r/rust/comments/j9ogc4/have_you_stopped_using_rust_and_if_yes_why/)
- url: https://www.reddit.com/r/rust/comments/j9ogc4/have_you_stopped_using_rust_and_if_yes_why/
---
Hi there,

I would be interested to read some of the stories of people who were using Rust for some time, but then decided it wasn't the good choice for their project or situation. There is a lot of praise around the language, for good reasons, and a lot of focus on its unique features (or lack of terrible ones). I feel it could be interesting to hear a bit about the other stories, about situations where some of you decided their problem wasn't in fact a good fit.

You know the saying: "use the right tool for the job", so what are jobs where you **actually tried** to use Rust but found out it wasn't the right tool?
## [8][Learn Rust by building the game Snake](https://www.reddit.com/r/rust/comments/j9o6lu/learn_rust_by_building_the_game_snake/)
- url: https://blog.scottlogic.com/2020/10/08/lets-build-snake-with-rust.html
---

## [9][I just published my first crate: imperative-rs, deriving instruction sets from enums](https://www.reddit.com/r/rust/comments/j9ol07/i_just_published_my_first_crate_imperativers/)
- url: https://www.reddit.com/r/rust/comments/j9ol07/i_just_published_my_first_crate_imperativers/
---
Hey everyone,

after seeing the post about the [Game Boy Color emulator for the browser](https://www.reddit.com/r/rust/comments/j8vw84/i_made_a_gameboygameboy_color_emulator_which_runs/) last night, I thought it was finally time to publish the little project I've been working on for the last few months.

[imperative-rs](https://crates.io/crates/imperative-rs) can be used to automatically derive instruction sets from enums:

    #[derive(InstructionSet)] 
    enum Instruction {
        #[opcode="0x00_xx_yy"]
        Add{x:u8, y:u8},
        #[opcode="0x01_xx_yy"]
        Mov{x:u8, y:u8}
    }
    fn exec(&amp;mut self) -&gt; impl Error
        let num_bytes, instr = Instruction::decode(&amp;mem[self.pc..])?;
        self.pc += num_bytes;
        match instr {
            Instruction::Add(x, y) =&gt; self.acc = x + y,
            Instruction::Mov(x, y) =&gt; self.reg[y] = self.mem[x],
        }
    }

So instead of matching on tons of in literals like this:

    match opcode {
        0x00 =&gt; self.nop(),
        0x01 =&gt; {
            match next_byte {
                //some more instructions here
            }
        }
    }

you can decode the instruction from memory and just match on your enum with all the instructions in it.

    enum Instruction {
        Nop,
        Mov{s:usize, t:usize},
        Ldi{....},
        Add{....},
        //etc.
    }

and

    match instr {
        Instruction::Nop =&gt; {},
        Instruction::Mov{s, t} =&gt; self.mem[t] = self.mem[s],
        //etc.
    }

You can define opcodes in hex or bin format, with or without underscores. the macro checks for collisions (i.e. opcodes that can't be distinguished from one another). Only restriction is that variable names need to be one symbol long and can't be hex digits for now.

 To decode instructions from your memory you just hand over a memory slice starting at the current program counter and it give you back an instruction and the number of bytes used or an error if no instruction could be decoded.
## [10][Announcing httpmock v0.5.0](https://www.reddit.com/r/rust/comments/j9o8zq/announcing_httpmock_v050/)
- url: https://www.reddit.com/r/rust/comments/j9o8zq/announcing_httpmock_v050/
---
Hi all!

A while ago, I released an HTTP mocking library called [httpmock](https://crates.io/crates/httpmock).  It got a major overhaul recently with a strong focus on improving usability and ease of use. I am very excited to announce the following changes in the new [version 0.5.0](https://github.com/alexliesenfeld/httpmock/releases/tag/v0.5.0).

**Most important changes:**

* A new API that makes it easy to read and write tests (see [examples](https://crates.io/crates/httpmock)).
* Assertion failures now print very useful and smart error messages that will help you identify and fix bugs in your tests.
* Many new helpers for easy request matching and response definition (such as cookies, redirects, binary and file body, response delays, and more).
* Documentation improvements (a lot!).

[This article](https://medium.com/@alexliesenfeld/rust-http-testing-with-httpmock-3d411200669c) demonstrates many of the new features in the form of a tutorial.
## [11][Contract as Code as Contract: using Rust to unify specification and implementation of web services APIs](https://www.reddit.com/r/rust/comments/j9qagk/contract_as_code_as_contract_using_rust_to_unify/)
- url: https://www.youtube.com/watch?v=EmSjZbSzA3A
---

## [12][hors v0.7.2 is released, with higher performance (compared to previous version, and origin howdoi) Any code suggestions (especially about tokio usage code) is very appreciated](https://www.reddit.com/r/rust/comments/j9oeet/hors_v072_is_released_with_higher_performance/)
- url: https://i.redd.it/kmfunojx3ns51.jpg
---

