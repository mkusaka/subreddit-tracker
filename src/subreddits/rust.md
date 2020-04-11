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

## [3][Intermodal: A nice torrent creator written in Rust](https://www.reddit.com/r/rust/comments/fytczt/intermodal_a_nice_torrent_creator_written_in_rust/)
- url: https://www.reddit.com/r/rust/comments/fytczt/intermodal_a_nice_torrent_creator_written_in_rust/
---
Hi everyone!

I just finished the first release of a project I've been working on for a while, Intermodal.

Intermodal is a command-line BitTorrent `.torrent` file utility. The binary is called `imdl`, and at the moment, it can create, verify, and display the contents of torrents.

It has a whole bunch of nice features, like cute progress bars, file inclusion and exclusion and exclusion with globs, and an automatic piece length picker that will chose a good piece length for torrents based on the size of their contents.

This is the first step in a much larger project to try to improve the state of decentralized content creation, distribution, and consumption. I go into a lot of detail about the current state of the project, and where I hope to take it in the future, in a blog post [here](https://rodarmor.com/blog/intermodal/).

The current version of Intermodal would not be nearly as good if it wasn't for the Rust language, and all the wonderful crates that the community has created.

In particular, I want to thank the creators of `globset`, `regex`, `indicatif`, `ansi_term`, `serde`, `snafu`, `tempfile`, `walkdir`, `structopt`, `clap`, and `bendy`. (Although I honestly feel bad leaving out all the little guys in my `Cargo.toml`, like `atty`. What would I do without `atty`!)

[`bendy`](https://github.com/P3KI/bendy), an excellent crate for encoding and decoding bencode, which is the encoding format that BitTorrent uses for `.torrent` files, didn't initially have Serde support. The maintainers were supportive of the idea of adding it, and were super friendly and and responsive, which made contributing a breeze. So thanks to @thequux and @0ndorio on GitHub!

Development is hosted [here](https://github.com/casey/intermodal), and there are a bunch of good first issues if you're interested in contributing:

- [Setting up code coverage on CI.](https://github.com/casey/intermodal/issues/9)
- [Adding web seeds to torrents.](https://github.com/casey/intermodal/issues/92).
- [Adding another kind of web seeds to torrents.](https://github.com/casey/intermodal/issues/93)
- [Supporting the addition of arbitrary keys to created torrents.](https://github.com/casey/intermodal/issues/23)
- [Adding a config file containing profiles to use for torrent file creation.](https://github.com/casey/intermodal/issues/36)
- [Adding support for generating file padding.](https://github.com/casey/intermodal/issues/99)
- [Adding a  whole new subcommand to edit existing `.torrent` files.](https://github.com/casey/intermodal/issues/124)
- [Fixing a no-doubt silly bug causing tests to leave behind `.torrent` files in `/tmp`.](https://github.com/casey/intermodal/issues/344)
- [Adding file selections to magnet links.](https://github.com/casey/intermodal/issues/245)
- [Creating `.torrent` files from magnet links.](https://github.com/casey/intermodal/issues/255)
- [Verifying multiple `.torrent` files at a time.](https://github.com/casey/intermodal/issues/165)
- [Showing nonstandard fields in `imdl torrent show`.](https://github.com/casey/intermodal/issues/168)
- [Adding a `--quiet` flag to `imdl torrent create`.](https://github.com/casey/intermodal/issues/174)
- [Showing corrupted piece information during verification.](https://github.com/casey/intermodal/issues/192)
- [Supporting BitTorrent V2 torrents.](https://github.com/casey/intermodal/issues/101)

If you need to create a torrent, definitely give it a shot!
## [4][Programming Servo: My own private runtime.](https://www.reddit.com/r/rust/comments/fywv7n/programming_servo_my_own_private_runtime/)
- url: https://www.reddit.com/r/rust/comments/fywv7n/programming_servo_my_own_private_runtime/
---
An overview of how Servo, a large parallel, concurrent, and multiprocess web engine written in Rust, embeds a JS/Wasm VM.

[https://medium.com/programming-servo/programming-servo-my-own-private-runtime-8a5ba74c63c8](https://medium.com/programming-servo/programming-servo-my-own-private-runtime-8a5ba74c63c8)
## [5][Is there something I don't understand here ?](https://www.reddit.com/r/rust/comments/fz40nu/is_there_something_i_dont_understand_here/)
- url: https://www.reddit.com/r/rust/comments/fz40nu/is_there_something_i_dont_understand_here/
---
Note that I'm just giving a quick thought about Ok wrapping and try blocks from an end user perspective. Not sorry for another post on the subject.

Rust is the most attractive language for me for many reasons, including the fact that it's not \*too much\* magic, clear enough and obviously all the reasons it's Rust. But I want to emphasize the fact that it's clear and readable, I always felt comfortable writing Ok and Err, I know what it means and it gives a good look to my code, there is a good amount of helper methods in std that makes it bearable and I know when I look at the end of my function and I see Ok(thing) that at this moment everything went well, it's easy to write, I don't want to save 3 keystrokes if it makes my code looks bad and I really think it's gonna make my code look bad.

What is the problem with Ok(\_) ? Nothing has yet convinced me that there was a problem with this, and I can't see what is the advantage of going the mainstream way... and I hate JS, why would I want to see JS style try blocks ? or remove the Ok wrapping that makes my code feel clear ? Maybe I'm not informed well enough or not experienced enough, I don't know, but my first thought about this thing is this.

tl;dr, Try blocks are ugly and make me feel like I'm writing JS and I hate it, and Ok(\_) looks good to me.
## [6][The differences between Ok-wrapping, try blocks, and function level try](https://www.reddit.com/r/rust/comments/fyj43p/the_differences_between_okwrapping_try_blocks_and/)
- url: https://yaah.dev/try-blocks
---

## [7][native Android app with Kotlin and Rust](https://www.reddit.com/r/rust/comments/fywgwj/native_android_app_with_kotlin_and_rust/)
- url: https://gitlab.com/dpezely/native-android-kotlin-rust
---

## [8][A new programming language for malayalees,based on malayalam memes,written completely in rust](https://www.reddit.com/r/rust/comments/fyvl10/a_new_programming_language_for_malayaleesbased_on/)
- url: https://github.com/Sreyas-Sreelal/malluscript
---

## [9][Using struct method on other thread?](https://www.reddit.com/r/rust/comments/fz2pmi/using_struct_method_on_other_thread/)
- url: https://www.reddit.com/r/rust/comments/fz2pmi/using_struct_method_on_other_thread/
---
I have seen questions like this before, but I don't get my head around the answers.  
So, please, could someone tell me why my code is not working and how to get this piece of code working.  


    use std::thread;
    use std::thread::JoinHandle;
    
    pub struct TestStruct {
        string_vec: Vec&lt;String&gt;
    }
    
    impl TestStruct {
    
        pub fn new() -&gt; Self {
            TestStruct {
                string_vec: Vec::new()
            }
        }
    
        pub fn vec_push(&amp;mut self) {
            self.string_vec.push(String::from("+1"));
        }
    
        pub fn spawn_thread(&amp;mut self) -&gt; JoinHandle&lt;()&gt; {
            thread::spawn(move || {
                self.vec_push();
            })
        }
    
    }
## [10][Create HashMap of Rust struct properties](https://www.reddit.com/r/rust/comments/fz345k/create_hashmap_of_rust_struct_properties/)
- url: https://www.reddit.com/r/rust/comments/fz345k/create_hashmap_of_rust_struct_properties/
---
I want to create a HashMap of a Rust struct. The key is the name of the property and the value is the value of the property. This is my current attempt:

    pub fn into_query_values(self) -&gt; cdrs::query::QueryValues {
        use std::collections::HashMap;
        let mut values: HashMap&lt;String, cdrs::types::value::Value&gt; = HashMap::new();
    
        #(
            values.insert(stringify!(#idents), self.#idents);
        )*
    
        cdrs::query::QueryValues::NamedValues(values)
    }

This code (and the other code) can be cloned from this repo:  [https://github.com/Jasperav/cdrs-helpers-derive/tree/auto\_into\_query\_values](https://github.com/Jasperav/cdrs-helpers-derive/tree/auto_into_query_values). Trying to cargo expend the 'SomeStruct' inside examples/src/main.rs will reproduce the compile time error.

I expect it will do something like this:

    values.insert("pk", self.pk);
    values.insert("name", self.name);
## [11][Why can’t you ? inside a let x = {} block?](https://www.reddit.com/r/rust/comments/fyx8va/why_cant_you_inside_a_let_x_block/)
- url: https://www.reddit.com/r/rust/comments/fyx8va/why_cant_you_inside_a_let_x_block/
---
I'm trying to avoid using a ton of `if let`'s  inside of each other here. So I thought I could do this:

	let x: Option&lt;type&gt; = {
		let a = someFunctionThatReturnsOption()?;
		let b = someOtherOptionFunction()?;
		Some(a + b)
	}

Why can't I do this?

Instead of doing something elegant like that I've had to do:

    let dist: Option&lt;f32&gt; = {
        if let Some(mouse_position) = input.mouse_position() {
            if let ActiveCamera(Some(camera_ent)) = *active_camera {
                if let Some(camera) = camera_comp.get(camera_ent) {
                    if let Some(camera_transform) = transforms.get(camera_ent) {
                        let camera_trans = camera_transform.translation();
                        Some(
                            ((camera_trans.x - mouse_position.0).powi(2)
                                + (camera_trans.y - mouse_position.1).powi(2))
                            .sqrt(),
                        )
                    } else {
                        None
                    }
                } else {
                    None
                }
            } else {
                None
            }
        } else {
            None
        }
    };

It’s horrible.
## [12][Swift: Google's bet on differentiable programming](https://www.reddit.com/r/rust/comments/fyt1ou/swift_googles_bet_on_differentiable_programming/)
- url: https://tryolabs.com/blog/2020/04/02/swift-googles-bet-on-differentiable-programming/
---

