# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (35/2020)!](https://www.reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (35/2020)?](https://www.reddit.com/r/rust/comments/ifif3r/whats_everyone_working_on_this_week_352020/)
- url: https://www.reddit.com/r/rust/comments/ifif3r/whats_everyone_working_on_this_week_352020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-35-2020/47702?u=llogiq)!
## [3][I printed "Rewrite it in Rust" swag for students](https://www.reddit.com/r/rust/comments/ifkrvg/i_printed_rewrite_it_in_rust_swag_for_students/)
- url: https://i.redd.it/wedjvxv3nwi51.jpg
---

## [4][My mother made me a plushie Rustacean 😁🦀](https://www.reddit.com/r/rust/comments/if7ns8/my_mother_made_me_a_plushie_rustacean/)
- url: https://i.redd.it/nzu7n80disi51.jpg
---

## [5][RsMixer - pulseaudio volume mixer written in rust](https://www.reddit.com/r/rust/comments/iflu4n/rsmixer_pulseaudio_volume_mixer_written_in_rust/)
- url: https://www.reddit.com/r/rust/comments/iflu4n/rsmixer_pulseaudio_volume_mixer_written_in_rust/
---
Hi! RsMixer is a pulsaudio volume mixer with a terminal user interface and vi-like keybindings. It is configurable (both keybindings and colors).

[GitHub](https://github.com/jantap/rsmixer)

  


https://preview.redd.it/1us0rjgl6xi51.png?width=935&amp;format=png&amp;auto=webp&amp;s=93d744de13437c2edc10b5ab2199eecdd2c70df5

This is my first big project written in rust, so any constructive criticism of the code is welcome.
## [6][rust-analyzer changelog #39](https://www.reddit.com/r/rust/comments/ifo5bs/rustanalyzer_changelog_39/)
- url: https://rust-analyzer.github.io/thisweek/2020/08/24/changelog-39.html
---

## [7][Let The Caller Decide The Return Value!](https://www.reddit.com/r/rust/comments/ifn33k/let_the_caller_decide_the_return_value/)
- url: https://www.justanotherdot.com/posts/let-the-caller-decide-the-return-value.html
---

## [8][One year of Nushell - a look back over a year of Nushell being public](https://www.reddit.com/r/rust/comments/if7rkw/one_year_of_nushell_a_look_back_over_a_year_of/)
- url: http://www.nushell.sh/blog/2020/08/23/year_of_nushell.html
---

## [9][Nano-ECS 0.4.0 released!](https://www.reddit.com/r/rust/comments/ifl003/nanoecs_040_released/)
- url: https://crates.io/crates/advancedresearch-nano_ecs
---

## [10][As above, so below: Using Rust generics to develop two bare-metal flash drivers](https://www.reddit.com/r/rust/comments/if21da/as_above_so_below_using_rust_generics_to_develop/)
- url: https://www.ecorax.net/as-above-so-below-1/
---

## [11][BEVY Rust Game Engine](https://www.reddit.com/r/rust/comments/if9jit/bevy_rust_game_engine/)
- url: https://www.youtube.com/watch?v=buUw0xAs4KM
---

## [12][Why can't I cast to trait like that?](https://www.reddit.com/r/rust/comments/ifnul0/why_cant_i_cast_to_trait_like_that/)
- url: https://www.reddit.com/r/rust/comments/ifnul0/why_cant_i_cast_to_trait_like_that/
---
        let doors: Vec&lt;_&gt; = [
            Door::new(),
            Door::new(),
            Door::new(),
        ]
        .into_iter()
        .map(|v| {
            Rc::new(RefCell::new(v)) as Rc&lt;RefCell&lt;dyn IDoor&gt;&gt;
        })
        .collect();

I want to create multiple instances of Door and wrap each of them into Rc&lt;RefCell&lt;&gt;&gt;. Resulting collection elements should have trait type IDoor instead of specific type.

What's the correct way of doing that?

&amp;#x200B;

The compiler error is: the trait \`IDoor\` is not implemented for \`&amp;Door\`
