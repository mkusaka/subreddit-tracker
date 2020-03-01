# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (9/2020)!](https://www.reddit.com/r/rust/comments/f8ney8/hey_rustaceans_got_an_easy_question_ask_here_92020/)
- url: https://www.reddit.com/r/rust/comments/f8ney8/hey_rustaceans_got_an_easy_question_ask_here_92020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/f5413m/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 327](https://www.reddit.com/r/rust/comments/fabtwg/this_week_in_rust_327/)
- url: https://this-week-in-rust.org/blog/2020/02/25/this-week-in-rust-327/
---

## [3][Pietro Albini has joined the core team](https://www.reddit.com/r/rust/comments/fbkmon/pietro_albini_has_joined_the_core_team/)
- url: https://blog.rust-lang.org/inside-rust/2020/02/27/pietro-joins-core-team.html
---

## [4][A half-hour to learn Rust](https://www.reddit.com/r/rust/comments/fbenua/a_halfhour_to_learn_rust/)
- url: https://fasterthanli.me/blog/2020/a-half-hour-to-learn-rust/
---

## [5][Spring Cleanup](https://www.reddit.com/r/rust/comments/fbsrvl/spring_cleanup/)
- url: https://vorner.github.io/2020/03/01/spring-cleanup.html
---

## [6][Shipyard 0.3 release](https://www.reddit.com/r/rust/comments/fbo8wf/shipyard_03_release/)
- url: https://www.reddit.com/r/rust/comments/fbo8wf/shipyard_03_release/
---
[Shipyard](https://crates.io/crates/shipyard) is an Entity Component System crate. ECS is a pattern mostly used in games but not only. It fits really well with Rust, allowing easy composition and lifetime management.

**What's new**

* [User guide](https://leudz.github.io/shipyard/book)
Learning to use a crate with the documentation can be hard. There's now a guide to explain what can be done with Shipyard and how.
* No need to register components anymore
Components had to be registered before accessing them by using `World::new` or `World::register`.  
Storages are now automatically created when they are first accessed.
* `!Send` and `!Sync` components
All components had to be `Send + Sync`, this is a strict limitation to make sure storages can use parallel features.
With 0.3 `!Send` and `!Sync` types can be stored and accessed while still following Rust's rules. These rules limit threading for these types but doesn't always prevent it.
* `Unique` components
When we only need a single instance of a component, keeping an id around to access it can be annoying. A `Unique` component won't be attached to any entity but will have the storage all for itself.
* Components sorting
* `no_std` support
* And more 😉

**Small Example**

    use shipyard::prelude::*;

    struct Health(f32);
    struct Position {
        x: f32,
        y: f32,
    }

    #[system(Populate)]
    fn run(mut entities: &amp;mut Entities, mut positions: &amp;mut Position, mut healths: &amp;mut Health) {
        entities.add_entity(
            (&amp;mut positions, &amp;mut healths),
            (Position { x: 0.0, y: 0.0 }, Health(1000.0)),
        );
    }

    #[system(InAcid)]
    fn run(pos: &amp;Position, mut health: &amp;mut Health) {
        (&amp;pos, &amp;mut health)
            .iter()
            .filter(|(pos, _)| is_in_acid(pos))
            .for_each(|(_, mut health)| {
                health.0 -= 1.0;
            });
    }

    fn is_in_acid(pos: &amp;Position) -&gt; bool {
        // it's wet season
        true
    }

    fn main() {
        let world = World::new();

        world.run_system::&lt;Populate&gt;();
        world.run_system::&lt;InAcid&gt;();
    }

Lastly, I want to thank dakom and eldyer for all their help and support!
## [7][Chrono English: Parsing all kinds of human date input](https://www.reddit.com/r/rust/comments/fbt0lu/chrono_english_parsing_all_kinds_of_human_date/)
- url: https://docs.rs/chrono-english/0.1.4/chrono_english/
---

## [8][Finishing Touches for Plotly Visualisation + Crate for Extras](https://www.reddit.com/r/rust/comments/fbrwzo/finishing_touches_for_plotly_visualisation_crate/)
- url: https://shahinrostami.com/posts/programming/rust-notebooks/finishing-touches-for-visualisation/
---

## [9][Suggested VSCode extensions](https://www.reddit.com/r/rust/comments/fbrxxs/suggested_vscode_extensions/)
- url: https://www.reddit.com/r/rust/comments/fbrxxs/suggested_vscode_extensions/
---
I am a bit overwhelmed by the amount of Rust extensions and I am afraid some may overlap.

What is a good set of extensions ?

In particular code completion is flaky with my current installation... :-/
## [10][A look into ways to implement and share data with interrupt handlers in Rust](https://www.reddit.com/r/rust/comments/fbk027/a_look_into_ways_to_implement_and_share_data_with/)
- url: https://therealprof.github.io/blog/interrupt-comparison/
---

## [11][Can WASM lead to accelerated adoption of rust?](https://www.reddit.com/r/rust/comments/fbo4eo/can_wasm_lead_to_accelerated_adoption_of_rust/)
- url: https://www.reddit.com/r/rust/comments/fbo4eo/can_wasm_lead_to_accelerated_adoption_of_rust/
---
Hello!


As per my observations, most languages have "killer apps" that lead to widespread use of the language and also large influx of new developers into the language. Ruby had Rails. Golang had docker and kubernetes. For python, it's scientific computing ( I know this isn't one particular project like rails or docker. But doesn't make any difference for our purpose) and web frameworks like django. But rust, imo, is yet to have such a project written in it. 

Although both golang and rust came about around the same time (10 years ago approx.), there is huge difference in the extent these two languages have been adopted. I understand that these two languages are unique, each with its own "domain of use". Also it's easier to get started with golang for new developers coming into the language, compared to rust (I don't think anyone will disagree with me here). But I still think that not having a "killer app" is a cause for the low adoption of rust language, even though rust offers highly beneficial features like memory safety and catching dataraces at compile time etc. 

But I think we've a "killer app" in rust. It's web assembly. Although web assembly isn't particular to rust, I don't see any other language community putting as much effort as rust community in this direction. Especially making use of web assembly outside the browser, for providing both sandboxing of applications and also easy sharing of them. So these are some of my  conclusions drawn from my observations as a polyglot, using many languages. I'd like to hear from other people in the community, what they think about recent developments in rust and wasm. 

Thanks.
## [12][[Call for help] - Let's make a truly native, truly cross platform GUI for rust a reality and I believe it can be done.](https://www.reddit.com/r/rust/comments/fbf3zw/call_for_help_lets_make_a_truly_native_truly/)
- url: https://twitter.com/ivanceras/status/1233791915792855040
---

