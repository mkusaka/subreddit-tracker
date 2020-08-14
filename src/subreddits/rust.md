# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (33/2020)!](https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 351](https://www.reddit.com/r/rust/comments/i8iqb9/this_week_in_rust_351/)
- url: https://this-week-in-rust.org/blog/2020/08/11/this-week-in-rust-351/
---

## [3][Clap: Faster development with tiny bounties](https://www.reddit.com/r/rust/comments/i9ihh0/clap_faster_development_with_tiny_bounties/)
- url: https://opencollective.com/clap/updates/faster-development-with-tiny-bounties
---

## [4][Core team statement so far; we are sad for our colleagues, Rust will be okay. We'll be posting more in the near future.](https://www.reddit.com/r/rust/comments/i994km/core_team_statement_so_far_we_are_sad_for_our/)
- url: https://twitter.com/rustlang/status/1294024734804508679
---

## [5][This Week in Conduit (Matrix Homeserver)](https://www.reddit.com/r/rust/comments/i9kp35/this_week_in_conduit_matrix_homeserver/)
- url: https://www.reddit.com/r/rust/comments/i9kp35/this_week_in_conduit_matrix_homeserver/
---
&gt; Conduit is a new implementation of a Matrix homeserver: https://conduit.rs

&gt; Matrix is a secure, private, decentralized, open source messaging platform. Check out https://matrix.org and https://element.io

This week I worked a bit on federation! Last week we were able to query room directories over federation; This week I looked into what's needed to join one of those rooms:

- Resolve room aliases over federation
- Send make_join request when a user tries to join an unknown room
- Use make_join response to create send_join request (insert origin, origin_server_ts and add hashes and signatures)

After implementing all those things, I could actually see Conduit's join
event on Synapse!

Of course a lot more work is needed to make this really work,
like creating the room on the Conduit side and sending and receiving federation
transactions, but this is a start.

Other news:

- Docker support (thanks to Weasy, Valkum and Paul)
- Room upgrade support (almost done, thanks to Faelar)

Thanks to everyone who supports me on [Liberapay](https://liberapay.com/timokoesters) or Bitcoin!
## [6][Why Rust is a great fit for embedded software](https://www.reddit.com/r/rust/comments/i9hpbh/why_rust_is_a_great_fit_for_embedded_software/)
- url: https://tweedegolf.nl/blog/39/why-rust-is-a-great-fit-for-embedded-software
---

## [7]["Much" of the Rust/Wasmtime team hit by layoffs at Mozilla](https://www.reddit.com/r/rust/comments/i8yfwj/much_of_the_rustwasmtime_team_hit_by_layoffs_at/)
- url: https://twitter.com/tschneidereit/status/1293868141953667074
---

## [8][Renom: Unreal Engine 4 Project Renamer](https://www.reddit.com/r/rust/comments/i9jv7s/renom_unreal_engine_4_project_renamer/)
- url: https://imgur.com/efEzpaX
---

## [9][error[E0207]: : the type parameter `...` is not constrained by the impl trait, self type, or predicates](https://www.reddit.com/r/rust/comments/i9jc85/errore0207_the_type_parameter_is_not_constrained/)
- url: https://www.reddit.com/r/rust/comments/i9jc85/errore0207_the_type_parameter_is_not_constrained/
---
[Playground](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=dccab07178f20db9574655175073f854)

How can I survive that? (without associated types)

    trait GenericTrait&lt;A&gt; {}
    trait OtherTrait {}
    
    // error[E0207]: the type parameter `A` is not constrained
    // by the impl trait, self type, or predicates
    impl&lt;T, A&gt; OtherTrait for T
    where
        T: GenericTrait&lt;A&gt;
    {}
## [10][Oh great power of reddit and its people, please help me find this crate by the description!](https://www.reddit.com/r/rust/comments/i9l302/oh_great_power_of_reddit_and_its_people_please/)
- url: https://www.reddit.com/r/rust/comments/i9l302/oh_great_power_of_reddit_and_its_people_please/
---
Hi, lovely people!

I once used _in now deleted project_ a crate that was able to get the file specified by 
the path and if built in release mode encode the content of that file into the binary, 
so the end user could have like web assets or configs already in the memory. If it's a
debug mode, however, it would just read stuff from the actual file instead.

If someone knows or seen it lately, please help! I've been trying to find it for about 3 
hours, but failed, I give up at this point! 

Rough API description: ~~short, blond hair, green eyes~~ you would define a struct
and and specify a path in the `#[]`.

Thanks

UPD: found it, it's called `rust-embed`
## [11][zookeeper-async - first release](https://www.reddit.com/r/rust/comments/i9iu08/zookeeperasync_first_release/)
- url: https://www.reddit.com/r/rust/comments/i9iu08/zookeeperasync_first_release/
---
Hi,

I would like to present my first published crate - [zookeeper-async](https://crates.io/crates/zookeeper-async). It's an async port of the original zookeeper crate, also updated to Rust 2018. Any feedback is welcome😊
## [12][Hardware recommendations for interfacing with a camera module](https://www.reddit.com/r/rust/comments/i9k8zg/hardware_recommendations_for_interfacing_with_a/)
- url: https://www.reddit.com/r/rust/comments/i9k8zg/hardware_recommendations_for_interfacing_with_a/
---
I tried posting this [in the easy questions post](https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/g19ph9m/), but to no avail.

I'm looking to try my hand at embedded programming and came up with the following toy project: a 2020 version of the [Trojan Room coffee pot](https://en.wikipedia.org/wiki/Trojan_Room_coffee_pot) monitoring system.

Essentially, what I want to do is something along these lines:

* Control a camera module using a microcontroller (ie. make it take pictures every X seconds)
* Read the image off of the camera
* Send the image to a server via wifi/bluetooth for processing

Eventually I want to convert the image of the coffee machine into a reading of the amount of coffee left and use the data for a wildly over-engineered Grafana/Icinga monitoring setup, but for starters I'd be happy just to be able to interface with the camera module!

The STM32F3 or similar Cortex-M MCU boards seem like a good bet considering the amount of learning materials available. Would this be a decent choice for interfacing with simple camera modules? What communication protocol is most suitable? Are there any recommended board + camera combinations?

I tried searching for answers and only found someone [struggling to get the STM32F3 to talk to an OV7670 via I2C](https://users.rust-lang.org/t/stm32f3-discovery-ov7670/27484), though that may be due to me knowing very, very little about embedded programming and not knowing what search terms to use.

Any other pointers including (but not limited to) "this is a bad idea, here's why:" would be very welcome!
