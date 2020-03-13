# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (11/2020)!](https://www.reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fc7h20/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 329](https://www.reddit.com/r/rust/comments/fhcuec/this_week_in_rust_329/)
- url: https://this-week-in-rust.org/blog/2020/03/10/this-week-in-rust-329/
---

## [3][Announcing Rust 1.42.0](https://www.reddit.com/r/rust/comments/fhhkqo/announcing_rust_1420/)
- url: https://blog.rust-lang.org/2020/03/12/Rust-1.42.html
---

## [4][Tips for approaching Rust code base](https://www.reddit.com/r/rust/comments/fhw6uu/tips_for_approaching_rust_code_base/)
- url: https://www.reddit.com/r/rust/comments/fhw6uu/tips_for_approaching_rust_code_base/
---
I previously programmed in Java and a little bit in C#.

With these languages, almost 99% of code-base you encounter have the convention that a source file holds a single class definition. Hence when you approach a file, cognitively when you are looking at definitions like methods, variables etc, you know it is part of a that very class you are introspecting.

Now that I am picking up Rust, I find things are done a little bit different. A source file could have multiple struct/enum definitions and definitions of methods/functions on structs are not as visually and cognitively linked together as it is with Java.

Hence, when I approach a code base in Rust and I am looking at definitions like methods, variables etc, it is a little bit harder to quickly tell which struct a function definition is for. And in some cases (which is not actually rare) there might be multiple structs defined in the file, making it harder to quickly tell where things belong. I am finding this to be a little confusing and straining cognitively.

How do experienced Rust developers deal with this? Any tips?
## [5][Traits, dynamic dispatch and upcasting](https://www.reddit.com/r/rust/comments/fhpas7/traits_dynamic_dispatch_and_upcasting/)
- url: https://articles.bchlr.de/traits-dynamic-dispatch-upcasting
---

## [6][Question about what happens when iterating over a vector](https://www.reddit.com/r/rust/comments/fhz6rk/question_about_what_happens_when_iterating_over_a/)
- url: https://www.reddit.com/r/rust/comments/fhz6rk/question_about_what_happens_when_iterating_over_a/
---
Hey

I have this simple code:

https://gist.github.com/rust-play/21158e9e5d4851abb8116a9707cae98e

Which gives me the error:

error[E0507]: cannot move out of `self.rules` which is behind a shared reference
  --&gt; src/lib.rs:16:21
   |
16 |         for rule in self.rules {
   |                     ^^^^^^^^^^
   |                     |
   |                     move occurs because `self.rules` has type `std::vec::Vec&lt;Rule&gt;`, which does not implement the `Copy` trait

Can someone please help me understand why a for over the vector causes a move?
## [7][The project that brute forced every possible melody used Rust](https://www.reddit.com/r/rust/comments/fhkji0/the_project_that_brute_forced_every_possible/)
- url: https://www.reddit.com/r/rust/comments/fhkji0/the_project_that_brute_forced_every_possible/
---
[The project](https://www.musictech.net/news/programmers-generate-every-possible-melody-in-midi-to-prevent-lawsuits/) has been getting lots of press, but the first mention of Rust I'd heard in relation to it was on last night's (3/11/2020) episode of [Off the Hook](https://www.2600.com/offthehook/2020/0320.html). The whole segment is worthwhile, but the developer talks about how and why they used Rust around 20:45.
## [8][Command error: request error nvim_call_function - Vim(let):E15: Invalid expression #59](https://www.reddit.com/r/rust/comments/fhy4vl/command_error_request_error_nvim_call_function/)
- url: https://www.reddit.com/r/rust/comments/fhy4vl/command_error_request_error_nvim_call_function/
---
Hello! I'm trying to figure out how to use coc-rls.  
I have setup cargo, have updated (:CocCommand rls.update) coc-rls latest version u/1.1.4, neovim :checkhealth is all ok, but when try to run .rs file i get this nasty message:

``` 
[coc.nvim] Command error: request error nvim_call_function - Vim(let):E15: Invalid expression: 'ps -o state= -o comm= -t '#{pane_tty}'  
 | grep -iqE '^[^TXZ ]+ +(\S+\/)?g?(view|n?vim?x?)(diff)?$''
```

CocInfo:
```
vim version: NVIM v0.5.0-139-g610755ff6
node version: v12.16.1
coc.nvim version: 0.0.75-d5e12d81b2
term: screen-256color
platform: linux

[coc.nvim] Setting sysroot to/home/sky/.rustup/toolchains/stable-x86_64-unknown-linux-gnu
[coc.nvim] running: rustup run stable-x86_64-unknown-linux-gnu rls, at /home/sky/rust_learning/a3_rust-by-example

## Output channel: prettier
## Output channel: snippets
```

CocOpenLog (selected printout of Error message):
```
2020-03-08T22:33:53.834 ERROR (pid:19851) [commands] - Error: request error nvim_call_function - Vim(let):E15: Invalid expression: 'ps -o state= -o comm= -t '#{pane_tty}'   | grep -iqE '^[^TXZ ]+ +(\S+\/)?g?(view|n?vim?x?)(diff)?$''
    at /home/sky/.config/nvim/plugged/coc.nvim/build/index.js:14243:32
    at NvimTransport.parseMessage (/home/sky/.config/nvim/plugged/coc.nvim/build/index.js:10847:17)
    at DecodeStream.&lt;anonymous&gt; (/home/sky/.config/nvim/plugged/coc.nvim/build/index.js:10817:18)
    at DecodeStream.emit (events.js:311:20)
    at addChunk (_stream_readable.js:294:12)
    at readableAddChunk (_stream_readable.js:275:11)
    at DecodeStream.Readable.push (_stream_readable.js:209:10)
    at DecodeStream.Transform.push (_stream_transform.js:152:32)
    at DecodeBuffer.DecodeStream.decoder.push (/home/sky/.config/nvim/plugged/coc.nvim/build/index.js:13767:12)
    at DecodeBuffer.flush (/home/sky/.config/nvim/plugged/coc.nvim/build/index.js:12757:12)
```

How to fix the above problem, and have coc-rls start to work? (PS. I posted on `github.com/coc.rls` but the site was not helpful)
## [9][Async Interview #7: Withoutboats](https://www.reddit.com/r/rust/comments/fhiizd/async_interview_7_withoutboats/)
- url: http://smallcultfollowing.com/babysteps/blog/2020/03/10/async-interview-7-withoutboats/
---

## [10][Someday Rust will form their own foundation and standardize the language?](https://www.reddit.com/r/rust/comments/fhueqp/someday_rust_will_form_their_own_foundation_and/)
- url: https://www.reddit.com/r/rust/comments/fhueqp/someday_rust_will_form_their_own_foundation_and/
---
The standarization of a language and libraries is important to evolve the future in the right way, I guess, so, Why Rust doesn't have it?  (Please be nice is my first post :D)
## [11][Confy 0.4 is released!](https://www.reddit.com/r/rust/comments/fhzolr/confy_04_is_released/)
- url: https://www.reddit.com/r/rust/comments/fhzolr/confy_04_is_released/
---
Hi folks

We the Rust CLI Working Group (a.k.a Clique) are pleased to announce the release of Confy 0.4. Confy is a boilerplate-free configuration library. 

Links:
crates.io: https://crates.io/crates/confy
Github: https://github.com/rust-cli/confy

Changes:
* Deal with bad TOML data
* Replace panics with a custom error type
* Ability to use crate name in the directory
* fix issues with dependencies, runtime cargo-manifest, etc
* Add `load_path` &amp; `save_path`
* Migrate to 2018 edition


Example:  
```rust
#[derive(Default, Debug, Serialize, Deserialize)]
struct MyConfig {
    version: u8,
    api_key: String,
}

fn main() -&gt; Result&lt;(), ::std::io::Error&gt; {
    let cfg: MyConfig = confy::load("my-app-name")?;
    dbg!(cfg);
    Ok(())
}
```

Thanks to all  our contributors :) 
You can reach out to us at #wg-cli on both the official discord &amp; zulip servers
## [12][Self-hosted bot for sending messages to yourself in Telegram](https://www.reddit.com/r/rust/comments/fhv1uj/selfhosted_bot_for_sending_messages_to_yourself/)
- url: https://github.com/Mr-Andersen/notify-tg
---

