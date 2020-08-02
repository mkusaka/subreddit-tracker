# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (31/2020)!](https://www.reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hurj77/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 349](https://www.reddit.com/r/rust/comments/i094wo/this_week_in_rust_349/)
- url: https://this-week-in-rust.org/blog/2020/07/28/this-week-in-rust-349/
---

## [3][Programming language popularity: Python overtakes Java – as Rust reaches top 20 | ZDNet](https://www.reddit.com/r/rust/comments/i241sd/programming_language_popularity_python_overtakes/)
- url: https://www.zdnet.com/article/programming-language-popularity-python-overtakes-java-as-rust-reaches-top-20/
---

## [4][Blue Team Rust: What is "Memory Safety", really?](https://www.reddit.com/r/rust/comments/i268wh/blue_team_rust_what_is_memory_safety_really/)
- url: https://tiemoko.com/blog/blue-team-rust/
---

## [5][gurk: Signal Messenger client for terminal](https://www.reddit.com/r/rust/comments/i2avtq/gurk_signal_messenger_client_for_terminal/)
- url: https://github.com/boxdot/gurk-rs
---

## [6][An efficient bit pattern matcher parsed at compiletime](https://www.reddit.com/r/rust/comments/i1z6nt/an_efficient_bit_pattern_matcher_parsed_at/)
- url: https://play.rust-lang.org/?version=beta&amp;mode=release&amp;edition=2018&amp;gist=2ff849086024a0a01b958060c3434570
---

## [7][A team for Rustacean Live Coders on Twitch - Join us if you stream coding in Rust.](https://www.reddit.com/r/rust/comments/i245g4/a_team_for_rustacean_live_coders_on_twitch_join/)
- url: https://twitch.tv/team/rustaceans
---

## [8][Beginner's guide to Error Handling in Rust](https://www.reddit.com/r/rust/comments/i2bv2m/beginners_guide_to_error_handling_in_rust/)
- url: http://www.sheshbabu.com/posts/rust-error-handling/
---

## [9][Weighty - A simple library and utility for reading from USB HID scales (for the two people out there who have those).](https://www.reddit.com/r/rust/comments/i23nxu/weighty_a_simple_library_and_utility_for_reading/)
- url: https://crates.io/crates/weighty
---

## [10][How to build a PinePhone game using Rust, Gtk, &amp; Cairo](https://www.reddit.com/r/rust/comments/i23m16/how_to_build_a_pinephone_game_using_rust_gtk_cairo/)
- url: https://github.com/richardanaya/pinephone-cairo-game-starter
---

## [11][Webference Rusty Days 2020: all recorded talks](https://www.reddit.com/r/rust/comments/i1satq/webference_rusty_days_2020_all_recorded_talks/)
- url: https://www.youtube.com/playlist?list=PLf3u8NhoEikhTC5radGrmmqdkOK-xMDoZ
---

## [12][Did I ran into the "First class type" issue?](https://www.reddit.com/r/rust/comments/i2bu77/did_i_ran_into_the_first_class_type_issue/)
- url: https://www.reddit.com/r/rust/comments/i2bu77/did_i_ran_into_the_first_class_type_issue/
---
Hello,

I am noobing around again. Can someone explain why the compiler cringes in the example below? Does it relate to the infamous first class type/higher order type/functional wizardry thing?

[Playground](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=e387a8a57e7108ec1e74f44f887d7af0) [Gist](https://gist.github.com/e387a8a57e7108ec1e74f44f887d7af0) 

    // This example demonstrates the fact you can (somewhat) use
    // generics of generics in function signature
    // but it is weird when you want to call it.
    
    // Here is an enum with 2 variants
    // Note that it bridges into some API
    // in which a third None variant makes no sense
    #[derive(Debug)]
    enum Dual {
        True,
        False,
    }
    
    // So it is natural to want to convert bools into this enum
    // for API friendliness purposes
    impl From&lt;bool&gt; for Dual {
        fn from(b: bool) -&gt; Self {
            match b {
                true =&gt; Self::True,
                false =&gt; Self::False,
            }
        }
    }
    
    // For my API, it is Option&lt;Dual&gt; that makes sense:
    
    // This does not compile
    // error[E0666]: nested `impl Trait` is not allowed
    // fn api0(x: impl Into&lt;Option&lt;impl Into&lt;Dual&gt;&gt;&gt;) {}
    
    // Let's try that
    // Spoiler: does not work
    fn api1&lt;I&gt;(x: impl Into&lt;Option&lt;I&gt;&gt;)
    where
        I: Into&lt;Dual&gt;
    {
        let x = x.into().map(|x| x.into());
        println!("api1 {:#?}", x);
    }
    
    // Ok last chance
    fn api2&lt;I, J&gt;(x: I)
    where
        I: Into&lt;Option&lt;J&gt;&gt;,
        J: Into&lt;Dual&gt;,
    {
        let x = x.into().map(|x| x.into());
        println!("api2 {:#?}", x);
    }
    
    pub fn main() {
        api1(true);
        api1(Dual::True);
        // api1(Some(true)); // cannot infer type for type parameter `I`
        // api1(Some(Dual::True)); // cannot infer type for type parameter `I`
        // api1(None); // cannot infer type for type parameter `I`
        // hum...
        // api1::&lt;bool&gt;(Some(true)); // explicit generic argument not allowed
        // api1::&lt;Dual&gt;(Some(Dual::True)); // explicit generic argument not allowed
        // api1::&lt;Dual&gt;(None); // explicit generic argument not allowed
        
        api2(true);
        api2(Dual::True);
        // api2(Some(true)); cannot infer type for type parameter `J`
        // api2(Some(Dual::True)); // cannot infer type for type parameter `J`
        // api2(None); // cannot infer type for type parameter `J`
        // api2::&lt;Option&lt;Dual&gt;, Dual&gt;(Some(true)); // expected enum `Dual`, found `bool`
        api2::&lt;Option&lt;bool&gt;, bool&gt;(Some(true));
        api2::&lt;Option&lt;Dual&gt;, Dual&gt;(None);
        api2::&lt;Option&lt;bool&gt;, bool&gt;(None);
        
        // So we can use api2, but the whole point of Into&lt;Option&lt;Into&lt;Dual&gt;&gt;&gt;
        // was to provide a nice API...
        
        // Thoughts?
    }
