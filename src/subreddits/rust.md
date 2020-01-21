# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (4/2020)!](https://www.reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here_42020/)
- url: https://www.reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here_42020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/eo6pjy/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (4/2020)?](https://www.reddit.com/r/rust/comments/eraop8/whats_everyone_working_on_this_week_42020/)
- url: https://www.reddit.com/r/rust/comments/eraop8/whats_everyone_working_on_this_week_42020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-4-2020/37181?u=llogiq)!
## [3][Deno JavaScript Runtime for V8 Written in Rust](https://www.reddit.com/r/rust/comments/erok5p/deno_javascript_runtime_for_v8_written_in_rust/)
- url: https://deno.land/std/manual.md
---

## [4][Actix-web is back in the main repo with a note from the Nikolay Kim](https://www.reddit.com/r/rust/comments/erdklb/actixweb_is_back_in_the_main_repo_with_a_note/)
- url: https://github.com/actix/actix-web/issues/1289
---

## [5][Pushrod 0.4.26 - SDL2 + Textures = Happy Face!](https://www.reddit.com/r/rust/comments/erp1yh/pushrod_0426_sdl2_textures_happy_face/)
- url: https://www.reddit.com/r/rust/comments/erp1yh/pushrod_0426_sdl2_textures_happy_face/
---
Hi all,

Pushrod is a GUI library written in Rust using the SDL2 library. It is written to be simple, fast, efficient, and easy to understand.

This release contains a huge number of changes to the rendering library.  I am now using `Texture`s via GPU instead of drawing directly to the `Canvas`, which makes for a much more modern, performant, and compact library.  I had been struggling trying to get Textures implemented for a couple of months, but managed to solve it.  Now, all on-screen components are `Texture`s!

Now that I have `Texture`s in place, it's full steam ahead.  The next series of updates will include a menu bar system at the top of each window, a scrollable `Viewport`, tab navigation, and a few more features before a feature freeze and a 0.5.x branch.

As always, [click here to see the project](https://www.github.com/KenSuenobu/rust-pushrod), and feel free to weigh-in, or provide a pull request if you see something that needs fixing, or just want to contribute.

I look forward to any and all comments - good or bad!

I also decided to get with the times and [created a blog.](https://kensuenobu.github.io/)
## [6][rweb: Yet another web server framework](https://www.reddit.com/r/rust/comments/errbq9/rweb_yet_another_web_server_framework/)
- url: https://www.reddit.com/r/rust/comments/errbq9/rweb_yet_another_web_server_framework/
---
**Note**: This is not a stable release.

Add rweb to the dependencies section
```toml
[dependencies]
rweb = "0.3.0-alpha.0"
```
and import as
```rust
use rweb::*;
```


---


rweb is a small http framework based on warp. It's fast, safe (thanks to warp and hyper), easy to use and extensible.

```rust
#[get("/")]
fn index() -&gt; Result&lt;String, Error&gt; {
    Ok(String::new())
}

#[get("/foo")]
fn foo() -&gt; Result&lt;String, Error&gt; {
    Ok(String::new())
}

// async works!
#[get("/param/{v}")]
async fn param_typed(v: u32) -&gt; Result&lt;String, Error&gt; {
    Ok(v.to_string())
}

#[get("/param/{name}/{value}")]
fn multiple_param(name: String, value: String) -&gt; Result&lt;String, Error&gt; {
    Ok(format!("{}={}", name, value))
}

#[tokio::main]
async fn main() {
    rweb::serve(index().or(foo()).or(param_typed()).or(multiple_param())).run(([127, 0, 0, 1], 3030)).await;
}

```

Note: There is a helper named #[router] to manage routes. I didn't use it for demonstrating purpose.


Repository: https://github.com/kdy1/rweb

See docs on docs.rs: https://docs.rs/rweb/0.3.0-alpha.0/rweb/
## [7][The Thank You Thread](https://www.reddit.com/r/rust/comments/ereh5h/the_thank_you_thread/)
- url: https://www.reddit.com/r/rust/comments/ereh5h/the_thank_you_thread/
---
Drop a comment here to say Thank You to all the open source maintainers of crates that you use.
## [8][Ready At Dawn Studios hiring contractors with preferred experience in Rust.](https://www.reddit.com/r/rust/comments/eribbz/ready_at_dawn_studios_hiring_contractors_with/)
- url: https://twitter.com/AndreaPessino/status/1219341435238895616
---

## [9][I'm looking for crates for a personal project](https://www.reddit.com/r/rust/comments/erq73f/im_looking_for_crates_for_a_personal_project/)
- url: https://www.reddit.com/r/rust/comments/erq73f/im_looking_for_crates_for_a_personal_project/
---
As the title suggests, I'm thinking through a personal project, and I'm wondering about crates and libraries. So, I thought I'd ask around for opinions and advice.

The project:

Most of the time, I spend my free time on reading and video games, but I recently got an itch to scratch. I recently finished a Nintendo Switch JRPG called [Atelier Ryza](https://store.steampowered.com/app/1121560/Atelier_Ryza_Ever_Darkness__the_Secret_Hideout/). Being a game about alchemy, it has rather involved item crafting mechanics, detailed ingredients, and flexible recipes for items. I often found myself yak-shaving (i.e. this recipe needs a (fuel) ingredient with a fire-element value of 3; I *could* just use some Ashen Sand I gathered earlier, but if I synthesize some Zettel paper first, I'll get a better result; hmm the Zettel paper recipe needs a Red Supplement to be awesome though, and I just ran out...) 

Yeah, I spend a lot of time in the crafting menu...

Anyway, the project!

So, there's a fair amount of details to the items and recipes to the game, with somewhat recursive relationships. There's (currently) no FAQ that really does this justice, in part because organizing and presenting the information as prose or individual tables only presents a slice of the whole picture.

So the project is to make a Atelier Ryza crafting information app.

Goals:

* step outside of my comfy skill set (C# web development)
* collect information about the items and recipes in the game
* create a stand-alone deployable app, that can be cheaply hosted, the items and recipes need to be stored with the app (I'm guessing some kind of schema file format; I'm considering JSON as a fallback option if I can't find something better)
* if the file format isn't conducive to manual creation, create a simple tool to help create the file
* host a simple UI to explore the data, and perhaps a GraphQL interface to do some more involved queries

Rough steps:

* record item and recipe information (likely by hand, which will take some time)
   * there's hundreds of items, and hundred of recipes--it'd likely only be a few megabytes in memory
* create a schema for storing items and recipes (I haven't touched on all of the things associated with items and recipes, but it's... involved)
* create the file to store the game data
* create a web host, using GraphQL to interact with the data
* create a UI that's not awful to look at (maybe play around with HSL-color design ideas  


I'm also wondering if Node.js would be a better fit for this project, rather than Rust. I'm interested in learning something new, and I'd love to hear from people that worked with both.

I'm also looking for libraries to help out. I'm not even sure what to do about a 'db-in-a-file'. I don't need a full sqlite solution, as there's only a few use cases.
## [10][Oxidize 2020 - The Embedded Rust Conference](https://www.reddit.com/r/rust/comments/erhv6n/oxidize_2020_the_embedded_rust_conference/)
- url: https://ferrous-systems.com/blog/oxidize-2020/
---

## [11][Writing an OS in Rust: Allocator Designs](https://www.reddit.com/r/rust/comments/erd5w1/writing_an_os_in_rust_allocator_designs/)
- url: https://os.phil-opp.com/allocator-designs/
---

## [12][More idiomatic way of passing a mutable reference to a data structure and immutable references to parts of the same structure?](https://www.reddit.com/r/rust/comments/erucgq/more_idiomatic_way_of_passing_a_mutable_reference/)
- url: https://www.reddit.com/r/rust/comments/erucgq/more_idiomatic_way_of_passing_a_mutable_reference/
---
Hi,

I'm currently implementing a board game, and I'm dividing various parts of the combat logic into functions wich look like this:

    fn apply_effect(board: &amp;mut Gameboard, effect: &amp;Effect, triggering_unit: &amp;Unit, unit_with_effect: &amp;Unit)
    fn attack_target(board: &amp;mut Gameboard, attacker: Unit, target_index: usize)

The typical way I'd use `apply_effect` is by passing in the board (since this function could create new units) with the other arguments being references to units from the board. For example, `&amp;board.army[0]`, which obviously violates the borrow rules.

The workarounds I've found is to clone the unit and pass a reference to the clone to the `apply_effect` function, or to use indices like in the second.

Am I doing something wrong here, or is there a more idiomatic way of structuring this code? Also, what exactly is the borrow checker trying to protect me from by insisting that I not do this? Because its pretty inconvenient right now unfortunately.
