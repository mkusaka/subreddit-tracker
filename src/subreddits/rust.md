# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (13/2020)!](https://www.reddit.com/r/rust/comments/fnfky9/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fnfky9/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fjef12/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (13/2020)?](https://www.reddit.com/r/rust/comments/fnflsh/whats_everyone_working_on_this_week_132020/)
- url: https://www.reddit.com/r/rust/comments/fnflsh/whats_everyone_working_on_this_week_132020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-13-2020/39864?u=llogiq)!
## [3][This Week in Rust 331](https://www.reddit.com/r/rust/comments/fp9z5t/this_week_in_rust_331/)
- url: https://this-week-in-rust.org/blog/2020/03/24/this-week-in-rust-331/
---

## [4][A MPEG-4 audio metadata parser](https://www.reddit.com/r/rust/comments/fp4hhk/a_mpeg4_audio_metadata_parser/)
- url: https://www.reddit.com/r/rust/comments/fp4hhk/a_mpeg4_audio_metadata_parser/
---
[mp4ameta](https://crates.io/crates/mp4ameta) is a library for parsing MPEG-4 audio metadata (m4a and m4b) and was heavily influenced by [id3](https://crates.io/crates/id3) and [metaflac](https://crates.io/crates/metaflac).
## [5][Do We Need Desktop, Or Is The Web Sufficient?](https://www.reddit.com/r/rust/comments/fp9a9v/do_we_need_desktop_or_is_the_web_sufficient/)
- url: https://www.reddit.com/r/rust/comments/fp9a9v/do_we_need_desktop_or_is_the_web_sufficient/
---
I recently wrote an orga tool for myself in Rust. I currently use the tool only via a CLI, but would like to create a GUI for it.

To understand my situation a bit better, here are some features and capabilities, that I'm going to require:

\- Interface with the file system, e.g. open directories, read &amp; write files

\- Alternatively to the file system interaction, connect to a server and interact with the file system there via a WebSocket / Rest API / WebDAV layer.

\- Make interactive lists, search bars and the usual widgets that an application needs.

As I see it, I now have a couple of options:

1. Port the functionality to a server architecture and interface with the server over a web application via a REST API or WebSocket.
2. Create a native Desktop GUI in Rust and there will be no need to port the functionality, since everything can be called via native Rust.
3. Create a native Desktop GUI in some other language/framework, such as ElectronJS, C#, Java. But if I go down this route, I'd have to go back to the server architecture and then I might aswell just write a web application, since I already have a lot of knowledge developing them (hence ElectronJS, but more on that below).

Therefore my questions and toughts:

\- I would like to avoid ElectronJS, since I see it as a very overloaded technology. Although I have to admit that it seems quite simple and smooth to use and develop something with it.

\- From my research, it seems, that native GUI development in Rust is still very immature and in it's early stages. There are frameworks/libaries available, such as [conrod](https://github.com/PistonDevelopers/conrod) or [orbtk](https://github.com/redox-os/orbtk), but I have no idea how viable they are.

\- I have heard of people writing in ElectronJS, porting it to WebAssembly, then to Rust, and then back to ElectronJS (lmao) to achieve quite a significant speedup, which might help with my worries about using ElectronJS.

\- How the heck was Figma created? I've heard that it is a Desktop Application ported to WebAssembly, but I can't find anything about what it originally was written in. My guess is C++ with a GUI framework, such as Qt or wxWidgets or maybe even something custom. It seems so perfect for what I want to do.

Ultimately it boils down to:

Can I create a Desktop Application (preferably in Rust) and port it to WebAssembly to display it in the browser, or am I going to have to go the ElectronJS route?

And to expand on the question in the title, I feel like everytime I think of scenario where I absolutely need a Desktop application, it seems that you could just wrap the functionality into an API and call it via a web application or just port the application via web assembly to the web, without going through the hassle of a desktop application.
## [6][My rust-themed home office decoration](https://www.reddit.com/r/rust/comments/fopy3b/my_rustthemed_home_office_decoration/)
- url: https://i.redd.it/gps2brjqhto41.jpg
---

## [7][Quickly Find Rust Program Bottlenecks Online Using a Go Tool](https://www.reddit.com/r/rust/comments/fp1q6k/quickly_find_rust_program_bottlenecks_online/)
- url: https://pingcap.com/blog/quickly-find-rust-program-bottlenecks-online-using-a-go-tool/
---

## [8][Macro Rules for Optional Trailing Commas](https://www.reddit.com/r/rust/comments/fozfe8/macro_rules_for_optional_trailing_commas/)
- url: https://www.reddit.com/r/rust/comments/fozfe8/macro_rules_for_optional_trailing_commas/
---

Up until now I've been using two rules (either one forwarding to the other or just copy-paste), which just felt clumsy.

I've seen [other](https://stackoverflow.com/questions/43143327/how-to-allow-optional-trailing-commas-in-macros) [posts](https://reddit.com/r/rust/comments/8e33ns/question_macros_how_to_make_a_macro_optionnally/) suggest `$($t:ty),*$(,)*`, or the modern `$(,)?` variant, but the first allows infinite trailing commas and the second still allows a trailing comma on an empty list.

If you're like me, you find this obscene permissiveness to be almost as bad as the two rule solutions. Maybe even worse, since you're exposing your macro end-users to the shame of allowing their typo-ridden code to compile!

Fortunately, `?` does actually provide the solution we've all been waiting for:

    macro_rules! one_rule_to_rule_them_all {
        ( $($($t:ty),+ $(,)?)? ) =&gt; { stringify!( $($($t,)*)? ) };
    }

I've written a [little test suite](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=2ee9f1644dd5f95ede72226988253ea0) comparing all the results of the different methods.
## [9][Noob question: how can I return a value from a HashMap in a struct?](https://www.reddit.com/r/rust/comments/fp7bti/noob_question_how_can_i_return_a_value_from_a/)
- url: https://www.reddit.com/r/rust/comments/fp7bti/noob_question_how_can_i_return_a_value_from_a/
---
I'll just admit it: I consider myself a pretty decent programmer, but Rust's borrow and lifetime rules are not clicking with me at all.  Could someone maybe point me in the right direction?

I've got this bit of code:

    use std::collections::HashMap;
    
    pub struct SpatialHash {
        pub contents: HashMap&lt;u64, Vec&lt;i32&gt;&gt;,
    }

    impl SpatialHash {
        pub fn push(&amp;mut self, x: i32, y: i32, entity: i32) {
            let h = SpatialHash::hash(x, y);
            let entities = self.contents.entry(h).or_insert(Vec::new());
            entities.push(entity);
        }

        pub fn get&lt;'a&gt;(self, x: i32, y: i32) -&gt; Option&lt;&amp;'a Vec&lt;i32&gt;&gt; {
            let h = SpatialHash::hash(x, y);
            self.contents.get(&amp;h)
        }

        // Hash a position into a spatial hash key.
        fn hash(x: i32, y: i32) -&gt; u64 {
            let x = x as u64;
            let y = y as u64;

            x &lt;&lt; 32 | y
        }
    }

But I get this error when I call my `get` function:

    error[E0515]: cannot return value referencing local data `self.contents`
      --&gt; src/spatial_hash.rs:17:9
       |
    17 |         self.contents.get(&amp;h)
       |         -------------^^^^^^^^
       |         |
       |         returns a value referencing data owned by the current function
       |         `self.contents` is borrowed here

I don't understand why "self.contents" is "borrowed" at that point, when I'm trying to return a reference to the value in the HashMap.  Any chance someone could explain this to me and how I could get that function to work?  

(I tried using those lifetimes on the get, but I've probably gotten that wrong too.)
## [10][A generic lexer in Rust using a simple match function or closure](https://www.reddit.com/r/rust/comments/fp2hsa/a_generic_lexer_in_rust_using_a_simple_match/)
- url: https://www.reddit.com/r/rust/comments/fp2hsa/a_generic_lexer_in_rust_using_a_simple_match/
---
[https://crates.io/crates/generic-lexer](https://crates.io/crates/generic-lexer)
## [11][Data Oriented Programming - The best of all worlds?](https://www.reddit.com/r/rust/comments/foy7ql/data_oriented_programming_the_best_of_all_worlds/)
- url: https://www.reddit.com/r/rust/comments/foy7ql/data_oriented_programming_the_best_of_all_worlds/
---
Since Rust has a somewhat unique focus on memory and safety, it looks to me like a perfect candidate to explore data oriented design and programming.

From my research I've seen it come up mostly in game development, but I would like to expand the conversation of this approach to different kinds of Software.

But first, here is my understanding and definition of data oriented programming, to lay out the perspective for the points I am going to make.

&gt;Data Oriented Design is the shift of focus to the data itself, especially how it is acted upon, layed out and stored in memory. It seperates data and behaviour and is mostly concerned about the transformation of an input to a specified output.

I think this paradigm mostly comes up in game development, since you deal with a lot of different kinds of data - Players, maps, mops, maps, networks, items, ... Since OOP often fails or abstracts this process too much, of dealing with such a vast amount of complex data, the focus now shifts to looking at the raw data and then thinking about the behaviour, grouping and functionality of the available data. OOP gives you no way of separating these two worlds.

This to me basically looks like functional programming, which can be very abstract, pure and mathematical, applied to the real world (which is often what games try to simulate) and less mystified. It expands on the purity of functional programming, since it more so wraps around the idea of an outside world with mutable state and complex systems. Functional programming seems very focused to the inside, wanting an immutable world with pure functions.

Now to the Rust portion. I think Rust is already very data-oriented in it's nature. You have structs to group data and allow for building custom data types. And traits, that allow for implementation of behavior of data (structs).

Abstractions are also very important in data oriented architectures, since you're not going to want to always deal with primitive data all the time and be very declarative in your code. This is where rusts "Cost-free abstraction" are insanely useful and efficient.

Rust also allows for a range of programming paradigms and styles, be it functional, imperative or declarative, so you can always mix and match to whatever problems you're trying to solve. Or wrap imperative processes in declarative / functional code.

In the end I feel like all software, at the heart, is data oriented, since programming is all about transforming data (literally).

I would like to see a future that is less OS, Application and Platform focused and more data oriented, so that data can live and interact freely outside of those boxes.
## [12][Rust Pie!](https://www.reddit.com/r/rust/comments/fopj3r/rust_pie/)
- url: https://i.redd.it/1ig1geehcto41.jpg
---

