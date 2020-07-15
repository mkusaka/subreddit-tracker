# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (29/2020)!](https://www.reddit.com/r/rust/comments/hq8id7/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hq8id7/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hm1pws/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 347](https://www.reddit.com/r/rust/comments/hrc4dt/this_week_in_rust_347/)
- url: https://this-week-in-rust.org/blog/2020/07/14/this-week-in-rust-347/
---

## [3][IntelliJ Rust 0.3: New Macro Expansion Engine](https://www.reddit.com/r/rust/comments/hrknxy/intellij_rust_03_new_macro_expansion_engine/)
- url: https://blog.jetbrains.com/clion/2020/07/intellij-rust-0-3-new-macro-expansion-engine/
---

## [4][Why even unused data needs to be valid](https://www.reddit.com/r/rust/comments/hrlozh/why_even_unused_data_needs_to_be_valid/)
- url: https://www.ralfj.de/blog/2020/07/15/unused-data.html
---

## [5][Security advisory for crates.io](https://www.reddit.com/r/rust/comments/hr7pt1/security_advisory_for_cratesio/)
- url: https://blog.rust-lang.org/2020/07/14/crates-io-security-advisory.html
---

## [6][pre: a crate to offer compile-time assistance for working with unsafe code](https://www.reddit.com/r/rust/comments/hrixq7/pre_a_crate_to_offer_compiletime_assistance_for/)
- url: https://www.reddit.com/r/rust/comments/hrixq7/pre_a_crate_to_offer_compiletime_assistance_for/
---
[pre](https://github.com/aticu/pre) is a library that helps you keep tabs on what exactly needs to be true in order for `unsafe` code to be safe.

You can think of it essentially like a `Safety` comment at the definition and call site, except that the compiler checks that the comments actually talk about the same thing.
This does not incur any cost at runtime in a release build.

It works by annotating functions with preconditions that need to be upheld when calling them.
These preconditions are transformed into an additional (*zero-sized*) argument to the function.
Callers then have to assure that the preconditions are upheld. Those assurances make up the additional argument at the call site and must match the original preconditions in order for the argument to type-check.
This avoids missing any changes in the safety requirements later. So it is most useful in larger projects with multiple people or for library authors, who want to make their preconditions part of the API.
The burden of proof still lies with the programmer however.

You can also think of it as a combination of the [`safety-guard` crate](https://crates.io/crates/safety-guard) and the [`safe` crate](https://crates.io/crates/safe).


## Example

Suppose you have this code:

```rust
fn init_foo() {
    /* ... */
}

/// # Safety
/// 
/// `init_foo` must be called before you can call `use_foo`.
unsafe fn use_foo() {
    /* ... */
}

fn main() {
    init_foo();

    unsafe {
        // Safety: This is safe, because `init_foo` was called first.
        use_foo();
    }
}
```

Using pre, it could be rewritten into this code:

```rust
fn init_foo() {
    /* ... */
}

use pre::pre;

#[pre("is only called after `init_foo` was called")]
unsafe fn use_foo() {
    /* ... */
}

#[pre]
fn main() {
    init_foo();

    unsafe {
        #[assure(
            "is only called after `init_foo` was called",
            reason = "`init_foo` was called first"
        )]
        use_foo();
    }
}
```

Now if someone who isn't aware of the relationship between the two functions writes this code somewhere else:

```rust
use pre::pre;

#[pre]
fn not_main() {
    unsafe {
        use_foo(); // Oops, didn't call `init_foo` first
    }
}
```

The code will fail to compile:

```text
error[E0061]: this function takes 1 argument but 0 arguments were supplied
  --&gt; src/main.rs:28:9
   |
7  |   #[pre("is only called after `init_foo` was called")]
   |  _______-
8  | | unsafe fn use_foo() {
9  | |     /* ... */
10 | | }
   | |_- defined here
...
28 |           use_foo(); // Oops, didn't call `init_foo` first
   |           ^^^^^^^-- supplied 0 arguments
   |           |
   |           expected 1 argument
```

The error message is not the most readable one, but I think it's better to spend a little time to understand an error than it is to have bugs in my program.

## Background

I'm writing this library for my bachelor's thesis and I'd be very happy if you could give me feedback about it.

Would you use it? If not, why not?

Feel free to [give it a try](https://github.com/aticu/pre) and let me know how it went.

Also if you are an author of a crate that might benefit from something like this, but don't have the time to integrate it yourself, feel free to contact me.
## [7][My journey of stumbling upon Rust](https://www.reddit.com/r/rust/comments/hrjt11/my_journey_of_stumbling_upon_rust/)
- url: https://www.reddit.com/r/rust/comments/hrjt11/my_journey_of_stumbling_upon_rust/
---
[DISCLAIMER: Story time]

Hello everyone, I’m new to reddit and I heard the rust community is welcoming so I really thought I’d share my story of getting to use Rust.

So I’m a Computing student (just finished my first year) and I kept searching for an interesting language to pick up and master. I’ve done about 3-4 years of C# and Java (.NET, Xamarin, Unity frameworks and Java in AndroidStudio and console apps) accumulated and by this point I’ve grown tired of it. I’ve also done C, Haskell and SQL and tried some html/vanilla js but the front-end dev life is not for me. I’ve seen really nice features (especially in Haskell, but it’s too purely functional for my taste) in each of those technologies but each of them lacked something. 

So, I kept searching while doing my webdev vanilla js course when I stumbled upon a load of talks about this “rust lang” thing. Binged like 10 talks for the fun of it and it looked really interesting, but I thought “ehhh maybe later”. After a week or two I ran out of things to watch and was too burned out from a game emulator I was working on in C# and thought “hey, what the hell, I’ve done C I liked it, Rust is just as efficient and has a load of interesting features I’ve never seen before [ownership, borrowing, lifetimes], so might as well try it” *also the cargo system is easy to use and extensible

I’ve gone through half of the official book in like 2 nights from 10pm-2am and I can’t begin to tell you how interesting this language is. It’s not like I have a lot of experience in the industry, but I haven’t seen anything like it, so it feels like it fills the hole. Also, a big reassurance is the fairly formal (i.e fairly abusive :D) compiler, which is a great help. Most of the time if the code compiles I can be pretty sure I won’t be able to break the code by dereferencing null pointers (because there are no nulls hah) or segfault it randomly (maybe I’m just bad at C), and I can be pretty sure the code works as intended. People talk a lot about rust and the compiler, comparing it with C gcc, doing benchmarks (which is a very fair match for the benchmarks I’ve seen, rust usually winning it out), but code safety is just as important and then cherry on top.

This was a long post, thanks for everyone who tagged along, but I just wanted to share my absolute and total hype.

For the ones still here, can you recommend me some good rust books that I can follow to become an expert in rust (yeah obviously will take a while but I think in 5-6 years’ time Rust will attract even more popularity, so it’s worth investing the time, since I’m a beginner-intermediate in most languages I’ve used)?

Good day everyone
## [8][Ayu theme is now available on nightly rustdoc!](https://www.reddit.com/r/rust/comments/hrfi8k/ayu_theme_is_now_available_on_nightly_rustdoc/)
- url: https://doc.rust-lang.org/nightly/std/
---

## [9][Adventures rewriting scientific FORTRAN software in Rust](https://www.reddit.com/r/rust/comments/hr9xgn/adventures_rewriting_scientific_fortran_software/)
- url: https://mckeogh.tech/post/shallow-water/
---

## [10][Learn Rust: Comprehensive Study Notes](https://www.reddit.com/r/rust/comments/hrkbzz/learn_rust_comprehensive_study_notes/)
- url: https://github.com/inancgumus/learnrust/
---

## [11][ReTeX - Typesetting mathematics](https://www.reddit.com/r/rust/comments/hrd8hx/retex_typesetting_mathematics/)
- url: https://www.reddit.com/r/rust/comments/hrd8hx/retex_typesetting_mathematics/
---
A few years ago I worked on a project to for typesetting a subset of LaTeX. When I started the project, my primary goal was to reach feature parity with KaTeX. This never happened, but I was relatively close. I never took the time to polish up a command line interface or develop appropriate APIs to ever consider it a released project. Like always, life got in the way and eventually the project was mostly forgotten about and remained dormant. But I think there is value in this prototype nonetheless.

So, I want to take a few minutes to describe this project, with the hope that it might inspire someone to pick it back up. I would be willing to mentor someone if they were curious about the overall structure if someone is interested. I only ask that we keep the mascot for a little bit longer :)

Github link: [https://github.com/ReTeX/ReX](https://github.com/ReTeX/ReX) 

There four main components to ReX:

1. Lexer -- This is a very simple parsing passing that streams to a parser. It categorizes everything as either a symbol, command, whitespace, or the end of the file.
2. Parser -- This parses the input into a abstract syntax tree. If I remember correctly, commands, like \\frac{1}{2} are resolved here into something more primitive -- a ParseNode::GenFraction in this case.
3. Layout -- The job of the layout engine is to take the abstract syntax tree and essentially output a format that can be used to actually render content. Generally you can think of this as a bunch of vertical and horizontal boxes of glyphs -- a collection of glyphs in a vertical box will be displayed stacked on top of eachother, and a collection of glyphs in a horizontal box will place glyphs one after the other. It's helpful to think of this as div versus span.
4. Render -- The renderer will take horizontal and vertical boxes of glyphs and produce something user facing. Only SVG is supported right now, but you could easily write a new renderer to rasterize this glyphs if you happen to have access to code that knows how to rasterize opentype fonts.

And that's basically it, it's relatively simple. There are a few other things worth mentioning:

1. The unit test framework is kind of nice. It will take a yaml of expressions to layout, then compare the layout with the last recorder layout, and generate an SVG diff for any expression that has changed. One expression will be shown in green (I think the old one?) and the other in red.
2. ReTeX needs to reference various attributes in the opentype font file. It wouldn't be terribly difficult to parse these out from the file themselves (except for maybe the dimensions of a glyph?), but during development I just wrote a python script to create lookup tables. This made it possible for the executable to run without being provided a reference to a valid font file. These are stored in hashmaps that are generated at compile time:
   1.  [https://github.com/ReTeX/ReX/blob/master/fonts/stix2/src/symbols.rs](https://github.com/ReTeX/ReX/blob/master/fonts/stix2/src/symbols.rs) \- Maps common commands to code points.
   2.  [https://raw.githubusercontent.com/ReTeX/ReX/master/fonts/stix2/src/glyphs.rs](https://raw.githubusercontent.com/ReTeX/ReX/master/fonts/stix2/src/glyphs.rs) \- Maps codepoints to glyph dimensions.
   3.  [https://github.com/ReTeX/ReX/blob/master/fonts/stix2/src/variants.rs](https://github.com/ReTeX/ReX/blob/master/fonts/stix2/src/variants.rs) \- Encodes the glyphs that should be used for various constructs -- like how should glyphs be pasted together to construct an infinitely tall parenthesis.
3. Oh yes, font parsing. That's a big part of this project. There are a bunch of python scripts that did a bunch of things to prepare fonts for use:  [https://github.com/ReTeX/rex-fonts/tree/master/tools](https://github.com/ReTeX/rex-fonts/tree/master/tools)
   1. I guess that some glyphs aren't accessible and it wouldn't be possible to put them in an SVG, so I wrote a script to take all inaccessible glyphs and place them in an private region:  [https://github.com/ReTeX/rex-fonts/blob/master/tools/accessible.py](https://github.com/ReTeX/rex-fonts/blob/master/tools/accessible.py) 
   2. The other scripts in that folder extract the relevant parameters that are needed from a font and exports them .rs file as a static hashmap.

So that's it for now. It used to compile. I don't know if it still does -- have there been any breaking changes in rustc in the past 2 years? I would like to work on this again, but I don't have much time. But more importantly, this project is now looking for *consumers*. If we can find consumers, then maybe we can get around to developing that command line interface or library. Currently, it's just an over complicated project that lets you run tests :)
## [12][Questions about licenses when distributing binary](https://www.reddit.com/r/rust/comments/hrmoj3/questions_about_licenses_when_distributing_binary/)
- url: https://www.reddit.com/r/rust/comments/hrmoj3/questions_about_licenses_when_distributing_binary/
---
So, as we know, when you add some library in Cargo, that library almost always pulls several other libraries and those pull other libraries as well. The resulting list of libraries used in your program can be a bit lengthy. We also know that rustc by default links everything statically. This of course means that if you distribute the resulting binaries, you're also distributing binary form of all those libraries.

Does that mean that in order to legally distribute binaries of my program, I have to obtain license texts of all of those libraries?

Edit The title is supposed to say "Question"
