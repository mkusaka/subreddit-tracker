# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (15/2020)!](https://www.reddit.com/r/rust/comments/fw2hd8/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fw2hd8/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/frfduy/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 333](https://www.reddit.com/r/rust/comments/fxqdrs/this_week_in_rust_333/)
- url: https://this-week-in-rust.org/blog/2020/04/07/this-week-in-rust-333/
---

## [3][Flutter RS - Build desktop apps in Flutter (using Rust backend) on stable branch](https://www.reddit.com/r/rust/comments/fzvojh/flutter_rs_build_desktop_apps_in_flutter_using/)
- url: https://github.com/flutter-rs/flutter-rs
---

## [4][What's the idiomatic way to handle functions which mix different result types and optionals?](https://www.reddit.com/r/rust/comments/fzrvi1/whats_the_idiomatic_way_to_handle_functions_which/)
- url: https://www.reddit.com/r/rust/comments/fzrvi1/whats_the_idiomatic_way_to_handle_functions_which/
---
For instance, I want to get a value which can either be pulled from:

1. a command line arg

2. read from a JSON config file

3. a default value

Here I have something which looks like this (pseudocode):



    use serde::{Deserialize};

    fn get_arg() -&gt; Option&lt;String&gt; { ... }
 
    #[derive(Deserialize)]
    struct Config {
        arg: Option&lt;String&gt;
    }

    fn my_function() -&gt; String {
        let arg: Option&lt;String&gt; = get_arg()

        match arg {
            Some(value) =&gt; return value.clone(),
            None =&gt; println!("no option from command line, parsing config")
        };

        let config_string = fs::read_to_string("path/to/config.json");
        match config_string {
            Ok(json) =&gt; {
                let config: Option&lt;Config&gt; = serde_json::from_str(&amp;json).unwrap_or( Config { arg: "default" } ); // unwrapping a Serde error
                return config.arg.unwrap_or("default");
            },
            Err(_) =&gt; return  "default"; // This would be an IO error
        };
    }


If these were all optionals, or the same type of error, it would be easy to do it with the `?` operator.  Is there a cleaner way to handle situations like this?

edit: formatting
## [5][How often does Rust change?](https://www.reddit.com/r/rust/comments/fz8mwm/how_often_does_rust_change/)
- url: https://words.steveklabnik.com/how-often-does-rust-change
---

## [6][I ripgrepped all crates on crates.io for profanity](https://www.reddit.com/r/rust/comments/fzc9fo/i_ripgrepped_all_crates_on_cratesio_for_profanity/)
- url: https://www.reddit.com/r/rust/comments/fzc9fo/i_ripgrepped_all_crates_on_cratesio_for_profanity/
---
Following [the recent article](https://www.reddit.com/r/rust/comments/fxxued/) on how to download all of crates.io I and did that and used `ripgrep` to search for profanity. It has unearthed things ranging from passionate rants about cryptography standards to insulting chat bots to TODOs on unsafe code.

Results:

[`rg --iglob '*.rs' -i fuck | awk 'length &lt;= 2048' fuck | grep -vi 'brainfuck' | grep -vi 'THE FUCK YOU WANT TO PUBLIC LICENSE' | grep -v 'DO WHAT THE FUCK YOU WANT TO'`](https://pastebin.com/4MaNZzyv)

[`rg --iglob '*.rs' -i shit | awk 'length &lt;= 2048' shit | grep -vi 'hashit' | grep -vi MATSUSHITA | grep -vi isHit`](https://pastebin.com/1w0JZF5W)
## [7][RFC: a practical mechanism for applying Machine Learning for optimization policies in LLVM](https://www.reddit.com/r/rust/comments/fzjf2d/rfc_a_practical_mechanism_for_applying_machine/)
- url: http://lists.llvm.org/pipermail/llvm-dev/2020-April/140763.html
---

## [8][My first Rust project. An RSA implementation. There's a lot left to do, but I was too thrilled to put it out there!](https://www.reddit.com/r/rust/comments/fzkcs1/my_first_rust_project_an_rsa_implementation/)
- url: https://github.com/rsarky/og-rsa
---

## [9][Introducing Dors -- makefiles for cargo that treat workspaces as first-class citizens](https://www.reddit.com/r/rust/comments/fzj945/introducing_dors_makefiles_for_cargo_that_treat/)
- url: https://www.reddit.com/r/rust/comments/fzj945/introducing_dors_makefiles_for_cargo_that_treat/
---
If you've ever tried to use cargo-make in a cargo workspace before, you'll know how frustrating it is to get working. Running tasks on all members of a workspace involves setting environment variables, you run into a circular dependency trying to later issue those same tasks on just one crate. Lastly, cargo-make implements a bunch of default behavior that can be hard to track down and work around.

So, I'd like to introduce an alternative: [https://github.com/aklitzke/dors](https://github.com/aklitzke/dors) . It's a task runner for cargo, but without a lot of those problems. I'm trying to make something that is easy to predict, read, use, and integrates well with cargo workspaces. It has:

\- Workspace support  
\- Autocompletion  
\- The ability to pass command-line arguments to your task  
\- The ability to set workspace-wide or crate-specific environment variables  
\- And a whole host more!

It's very new, so any feedback or feature requests would be welcome!

Thanks!
## [10][Programming Generic Interrupt Controller and Timer Interrupt for my AArch64 OS in Rust](https://www.reddit.com/r/rust/comments/fzj4vq/programming_generic_interrupt_controller_and/)
- url: https://lowenware.com/blog/osdev/aarch64-gic-and-timer-interrupt/
---

## [11][Is there something I don't understand here ?](https://www.reddit.com/r/rust/comments/fz40nu/is_there_something_i_dont_understand_here/)
- url: https://www.reddit.com/r/rust/comments/fz40nu/is_there_something_i_dont_understand_here/
---
Note that I'm just giving a quick thought about Ok wrapping and try blocks from an end user perspective. Not sorry for another post on the subject.

Rust is the most attractive language for me for many reasons, including the fact that it's not \*too much\* magic, clear enough and obviously all the reasons it's Rust. But I want to emphasize the fact that it's clear and readable, I always felt comfortable writing Ok and Err, I know what it means and it gives a good look to my code, there is a good amount of helper methods in std that makes it bearable and I know when I look at the end of my function and I see Ok(thing) that at this moment everything went well, it's easy to write, I don't want to save 3 keystrokes if it makes my code looks bad and I really think it's gonna make my code look bad.

What is the problem with Ok(\_) ? Nothing has yet convinced me that there was a problem with this, and I can't see what is the advantage of going the mainstream way... and I hate JS, why would I want to see JS style try blocks ? or remove the Ok wrapping that makes my code feel clear ? Maybe I'm not informed well enough or not experienced enough, I don't know, but my first thought about this thing is this.

tl;dr, \[Edit: I think\] Try blocks are ugly and make me feel like I'm writing JS and I hate it, and Ok(\_) looks good to me.
## [12][jlrs v0.2 has been released](https://www.reddit.com/r/rust/comments/fzhdvh/jlrs_v02_has_been_released/)
- url: https://www.reddit.com/r/rust/comments/fzhdvh/jlrs_v02_has_been_released/
---
Some time ago I released the first version of `jlrs`, a crate that provides (mostly) safe bindings to the Julia C API. This first version works, but introduces a lot of unnecessary overhead, complexity and unnecessary distinctions. The second version is a major rewrite, but addresses many of those shortcomings.

If you've used the first version, you'll know that version includes different contexts that separate allocating data and calling functions and data is exposed through handles. That has changed in v0.2; the different contexts have been replaced with frames that let you freely mix allocating data and calling functions, handles have been replaced by values which expose the data directly but safely. In general, things have been renamed to better reflect their names in the C API.  

There's only one new feature, really. You can now check whether some value is (an array) of a specific type with the methods `Value::is&lt;T&gt;`, `Value::is_array`, and `Values::is_array_of&lt;T&gt;` respectively.

As an example, this is how you can add two numbers with a dynamically growing frame:

    // Initialize Julia. Read the documentation to learn
    // more about the details
    let mut julia = unsafe { Julia::init(16).unwrap() };

    // We can only do things when we have access to a frame
    let x = julia.dynamic_frame(|frame| {
        // Create the two arguments
        let i = Value::new(frame, 2u64)?;
        let j = Value::new(frame, 1u32)?;

        // We can find the addition-function in the base 
        // module
        let func = Module::base(frame).function("+")?;

        // Call the function and unbox the result
        let output = func.call2(frame, i, j)?;
        output.try_unbox::&lt;u64&gt;()
    }).unwrap();

    assert_eq!(x, 3);

[Crate](https://crates.io/crates/jlrs)

[Documentation](https://docs.rs/jlrs)

[Repo](https://github.com/Taaitaaiger/jlrs)
