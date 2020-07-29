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
## [2][What's everyone working on this week (31/2020)?](https://www.reddit.com/r/rust/comments/hynlvw/whats_everyone_working_on_this_week_312020/)
- url: https://www.reddit.com/r/rust/comments/hynlvw/whats_everyone_working_on_this_week_312020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-31-2020/46441?u=llogiq)!
## [3][Can we please stop the "Is it just me or is Rust literally the best language" posts?](https://www.reddit.com/r/rust/comments/hzxrs1/can_we_please_stop_the_is_it_just_me_or_is_rust/)
- url: https://www.reddit.com/r/rust/comments/hzxrs1/can_we_please_stop_the_is_it_just_me_or_is_rust/
---
I don't want to offend anyone here, and I know my opinion will be controversial, but I see a lot of posts like this on r/rust that get upvoted to the top. They contribute nothing, and we all know Rust is a great language (that's why we're here!)

I think such posts just add to the Rust community's reputation of being a circlejerk. We should be focusing on new developments in the language, interesting projects built with the language, etc. Posting "Is it just me or is Rust the best language ever" on r/rust is obviously going to get you nothing but yes-people, and as a long time lurker on this sub, I feel such posts are growing increasingly repetitive and pointless.

Once again, I apologize if I've offended anyone.
## [4][Emigui deserves more love](https://www.reddit.com/r/rust/comments/hzwvsk/emigui_deserves_more_love/)
- url: https://github.com/emilk/emigui
---

## [5][I Rewrote The C Donut In Rust](https://www.reddit.com/r/rust/comments/hzkgep/i_rewrote_the_c_donut_in_rust/)
- url: https://www.reddit.com/r/rust/comments/hzkgep/i_rewrote_the_c_donut_in_rust/
---
You know that cool donut written in C that made a donut when ran? You know, [*that one*](https://www.a1k0n.net/2011/07/20/donut-math.html)? I decided to have some fun and rewrote it all in Rust! Now there are a few problems with it at the moment...

1. There are some rendering errors
2. All the code doesn't fit in the donut

I'm happy to say that both these errors are *only* in the compressed version that I wrote and not in the original. I could've remade the compressed version to fix these issues, but then I would've needed a larger donut, and although I know how to resize the donut, I'd have to fiddle with the size of the viewport to make sure it all shows, and I'm not sure how I'd go around that...

One more note, I also found out a bug that apparently both the original (C) and mine (Rust) have, and that's that it gets faster overtime and eventually just freezes. Since the original one had it I'm not gonna worry about it.

# Here's A Screenshot

https://preview.redd.it/xdyw60lx3nd51.png?width=814&amp;format=png&amp;auto=webp&amp;s=9e2ac715648eed4c796f099bd563ac36a1ca3155

And with the screenshot, here's [all of the Rust code](https://gist.github.com/Daniihh/cb550f402a6f4b80de8cf53184ea5625). I may make an improved version of the compressed code later.

# Credits

* Original code by [Andy Sloane](https://www.a1k0n.net/about.html)
* Help from [c2rust](https://c2rust.com/), which helped me find out what I incorrectly ported
## [6][Beginner's critiques of Rust](https://www.reddit.com/r/rust/comments/hzx1ak/beginners_critiques_of_rust/)
- url: https://www.reddit.com/r/rust/comments/hzx1ak/beginners_critiques_of_rust/
---
Hey all. I've been a Java/C#/Python dev for a number of years. I noticed Rust topping the StackOverflow most loved language list earlier this year, and I've been hearing good things about Rust's memory model and "free" concurrency for awhile. When it recently came time to rewrite one of my projects as a small webservice, it seemed like the perfect time to learn Rust.

I've been at this for about a month and so far I'm not understanding the love at all. I haven't spent this much time fighting a language in awhile. I'll keep the frustration to myself, but I do have a number of critiques I wouldn't mind discussing. Perhaps my perspective as a beginner will be helpful to someone. Hopefully someone else has faced some of the same issues and can explain why the language is still worthwhile.

Fwiw - I'm going to make a lot of comparisons to the languages I'm comfortable with. I'm not attempting to make a value comparison of the languages themselves, but simply comparing workflows I like with workflows I find frustrating or counterintuitive.

**Docs**

When I have a question about a language feature in C# or Python, I go look at the official language documentation. Python in particular does a really nice job of breaking down what a class is designed to do and how to do it. Rust's standard docs are little more than Javadocs with extremely minimal examples. There are more examples in the Rust Book, but these too are super simplified. Anything more significant requires research on third-party sites like StackOverflow, and Rust is too new to have a lot of content there yet.

It took me a week and a half of fighting the borrow checker to realize that `HashMap.get_mut()` was not the correct way to get and modify a map entry whose value was a non-primitive object. Nothing in the official docs suggested this, and I was actually on the verge of quitting the langugage over this until someone linked [Tour of Rust](https://tourofrust.com/00_en.html), which did have a useful map example, in a Reddit comment. (If any other poor soul stumbles across this - you need `HashMap.entry().or_insert()`, and you modify the resulting entry in place using `*my_entry.value = whatever`. The borrow checker doesn't allow getting the entry, modifying it, and putting it back in the map.)

**Pit of Success/Failure**

C# has the concept of a [pit of success](https://blog.codinghorror.com/falling-into-the-pit-of-success/): the most natural thing to do should be the correct thing to do. It should be easy to succeed and hard to fail.

Rust takes the opposite approach: every natural thing to do is a landmine. `Option.unwrap()` can and will terminate my program. `String.len()` sets me up for a crash when I try to do character processing because what I actually want is `String.chars.count()`. `HashMap.get_mut()` is only viable if I know ahead of time that the entry I want is already in the map, because `HashMap.get_mut().unwrap_or()` is a snake pit and simply calling `get_mut()` is apparently enough for the borrow checker to think the map is mutated, so reinserting the map entry afterward causes a borrow error. If-else statements aren't idiomatic. Neither is `return`.

**Language philosophy**

Python has the saying "we're all adults here." Nothing is truly private and devs are expected to be competent enough to know what they should and shouldn't modify. It's possible to monkey patch (overwrite) pretty much anything, including standard functions. The sky's the limit.
C# has visibility modifiers and the concept of sealing classes to prevent further extension or modification. You can get away with a lot of stuff using inheritance or even extension methods to tack on functionality to existing classes, but if the original dev wanted something to be private, it's (almost) guaranteed to be. (Reflection is still a thing, it's just understood to be dangerous territory a la Python's monkey patching.) This is pretty much "we're all professionals here"; I'm trusted to do my job but I'm not trusted with the keys to the nukes.
Rust doesn't let me so much as reference a variable twice in the same method. This is the functional equivalent of being put in a straitjacket because I can't be trusted to not hurt myself. It also means I can't do anything.

**The borrow checker**

This thing is legendary. I don't understand how it's smart enough to theoretically track data usage across threads, yet dumb enough to complain about variables which are only modified inside a single method. Worse still, it likes to complain about variables which aren't even modified.

Here's a fun example. I do the same assignment twice (in a real-world context, there are operations that don't matter in between.) This is apparently illegal unless Rust can move the value on the right-hand side of the assignment, even though the second assignment is technically a no-op.

	//let Demo be any struct that doesn't implement Copy.
    let mut demo_object: Option&lt;Demo&gt; = None;
    let demo_object_2: Demo = Demo::new(1, 2, 3);
    
    demo_object = Some(demo_object_2);
    demo_object = Some(demo_object_2);
	
Querying an Option's inner value via `.unwrap` and querying it again via `.is_none` is also illegal, because `.unwrap` seems to move the value even if no mutations take place and the variable is immutable:

    let demo_collection: Vec&lt;Demo&gt; = Vec::&lt;Demo&gt;::new();
    let demo_object: Option&lt;Demo&gt; = None;

    for collection_item in demo_collection {
        if demo_object.is_none() {
        }
        
        if collection_item.value1 &gt; demo_object.unwrap().value1 {
        }
    }
	
And of course, the HashMap example I mentioned earlier, in which calling `get_mut` apparently counts as mutating the map, regardless of whether the map contains the key being queried or not:

    let mut demo_collection: HashMap&lt;i32, Demo&gt; = HashMap::&lt;i32, Demo&gt;::new();
    
    demo_collection.insert(1, Demo::new(1, 2, 3));
    
    let mut demo_entry = demo_collection.get_mut(&amp;57);
    let mut demo_value: &amp;mut Demo;
    
    //we can't call .get_mut.unwrap_or, because we can't construct the default
    //value in-place. We'd have to return a reference to the newly constructed
	//default value, which would become invalid immediately. Instead we get to
	//do things the long way.
    let mut default_value: Demo = Demo::new(2, 4, 6);
    
    if demo_entry.is_some() {
        demo_value = demo_entry.unwrap();
    }
    else {
        demo_value = &amp;mut default_value;
    }
    
    demo_collection.insert(1, *demo_value);

None of this code is especially remarkable or dangerous, but the borrow checker seems absolutely determined to save me from myself. In a lot of cases, I end up writing code which is a lot more verbose than the equivalent Python or C# just trying to work around the borrow checker.

This is rather tongue-in-cheek, because I understand the borrow checker is integral to what makes Rust tick, but I think I'd enjoy this language a lot more without it.

**Exceptions**

I can't emphasize this one enough, because it's terrifying. The language flat up encourages terminating the program in the event of some unexpected error happening, forcing me to predict every possible execution path ahead of time. There is no forgiveness in the form of try-catch. The best I get is Option or Result, and nobody is required to use them. This puts me at the mercy of every single crate developer for every single crate I'm forced to use. If even one of them decides a specific input should cause a panic, I have to sit and watch my program crash.

Something like this came up in a Python program I was working on a few days ago - a web-facing third-party library didn't handle a web-related exception and it bubbled up to my program. I just added another except clause to the try-except I already had wrapped around that library call and that took care of the issue. In Rust, I'd have to find a whole new crate because I have no ability to stop this one from crashing everything around it.

**Pushing stuff outside the standard library**

Rust deliberately maintains a small standard library. The devs are [concerned](https://users.rust-lang.org/t/how-do-you-iterate-over-grapheme-clusters-of-a-string-in-rust/11338/4) about the commitment of adding things that "must remain as-is until the end of time."

This basically forces me into a world where I have to get 50 billion crates with different design philosophies and different ways of doing things to play nicely with each other. It forces me into a world where any one of those crates can and will be abandoned at a moment's notice; I'll probably have to find replacements for everything every few years. And it puts me at the mercy of whoever developed those crates, who has the language's blessing to terminate my program if they feel like it.

Making more stuff standard would guarantee a consistent design philosophy, provide stronger assurance that things won't panic every three lines, and mean that yes, I can use that language feature as long as the language itself is around (assuming said feature doesn't get deprecated, but even then I'd have enough notice to find something else.)

**Testing is painful**

Tests are definitively second class citizens in Rust. Unit tests [are expected](https://stackoverflow.com/questions/38995892/how-to-move-tests-into-a-separate-file-for-binaries-in-rusts-cargo/39009227#39009227) to sit in the same file as the production code they're testing. What?

There's no way to tag tests to run groups of tests later; tests can be run singly, using a wildcard match on the test function name, or can be ignored entirely using `[ignore]`. That's it.

**Language style**

This one's subjective. I expect to take some flak for this and that's okay.
* Conditionals with two possible branches should use if-else. Conditionals of three or more branches can use switch statements. Rust tries to wedge match into *everything*. Options are a perfect example of this - either a thing has a value (`is_some()`) or it doesn't (`is_none()`) but examples in the Rust Book only use match.
* Match syntax is virtually unreadable because the language encourages heavy match use (including nested matches) with large blocks of code and no language feature to separate different blocks. Something like C#'s `break`/`case` statements would be nice here - they signal the end of one case and start another. Requiring each match case to be a short, single line would also be good.
* Allowing functions to return a value without using the keyword return is awful. It causes my IDE to perpetually freak out when I'm writing a method because it thinks the last line is a malformed return statement. It's harder to read than a `return X` statement would be. It's another example of the Pit of Failure concept from earlier - the natural thing to do (`return X`) is considered non-idiomatic and the super awkward thing to do (`X`) is considered idiomatic.
* `return if {} else {}` is really bad for readability too. It's a lot simpler to put the return statement *inside* the if and else blocks, where you're actually returning a value.
## [7][A simple crud on rust (with rocket.rs and diesel.rs)](https://www.reddit.com/r/rust/comments/hzuoa8/a_simple_crud_on_rust_with_rocketrs_and_dieselrs/)
- url: https://medium.com/@luis_50157/a-simple-crud-on-rust-with-rocket-rs-and-diesel-rs-e885672cb23d
---

## [8][Do you miss Rust when you use another language?](https://www.reddit.com/r/rust/comments/hzm3sm/do_you_miss_rust_when_you_use_another_language/)
- url: https://www.reddit.com/r/rust/comments/hzm3sm/do_you_miss_rust_when_you_use_another_language/
---
I am currently writing a mircoservice that I wrote with Node to Golang.

I could find that I miss many features of Rust when I use other programming languages.(JavaScript, Python, Golang and Solidity a little bit.)

Am I the only one who feels like this? I don't miss other languages.

Edit - The post attracted more attentions than I expected. Thanks for sharing your thoughts. It is a decent opportunity to learn Rust with other languages even though that was not my intention for this post.
## [9][I love rust and RISC-V. Will it be possible to run rust programs on the SiFive-Learn-Inventor dev kit?](https://www.reddit.com/r/rust/comments/hzyju9/i_love_rust_and_riscv_will_it_be_possible_to_run/)
- url: https://devices.amazonaws.com/detail/a3G0h0000077I8lEAE/SiFive-Learn-Inventor
---

## [10][You Can Now Debug Programs Using GDB on Redox OS](https://www.reddit.com/r/rust/comments/hzd35z/you_can_now_debug_programs_using_gdb_on_redox_os/)
- url: https://www.redox-os.org/news/public-announcement-gdb/
---

## [11][deeply nested structs lifetime annotations](https://www.reddit.com/r/rust/comments/hzukuc/deeply_nested_structs_lifetime_annotations/)
- url: https://www.reddit.com/r/rust/comments/hzukuc/deeply_nested_structs_lifetime_annotations/
---
Yesterday, I was wondering how you can describe a very deeply nested structs lifetime for a function parameter that returns a reference that lives long enough. Can't imagine a real use case for it but how would you write the function?

    struct A&lt;'a&gt; {
        b: &amp;'a B&lt;'a&gt;
    }
    
    struct B&lt;'b&gt; {
        c: &amp;'b C&lt;'b&gt;
    }
    
    struct C&lt;'c&gt; {
        d: &amp;'c D&lt;'c&gt;
    }
    
    struct D&lt;'i&gt; {
        i: &amp;'i i32
    }
    
    fn my_int_ref&lt;'a&gt;(a: &amp;'a A) -&gt; &amp;'a i32 {
        a.b.c.d.i
    }
    
    fn main() {
        let i = &amp;5;
        let k;
        
        {
            let d = D { i };
            let c = C { d: &amp;d };
            let b = B { c: &amp;c };
            let a = A { b: &amp;b };
            k = my_int_ref(&amp;a);
        }
        
        dbg!(k);
    }

It's obvious why the compiler can't compile this. The function \`my\_int\_ref\` gets a lifetime \`'a\` which lives as long as the scope where \`A\` is created. Though, the return type \`&amp;i32\` lives longer than \`A\`. Is it possible to describe this use case without using raw pointers?  


[https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=b368e6fa13ec3e4cb4ce34559b9495cb](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=b368e6fa13ec3e4cb4ce34559b9495cb)
## [12][We need license features](https://www.reddit.com/r/rust/comments/i001mf/we_need_license_features/)
- url: https://www.reddit.com/r/rust/comments/i001mf/we_need_license_features/
---
I came across an interesting licensing issue in one of my crates and would be interested to share it.

If the solution I am describing already exists, I would be happy for anyone to point me in the right direction, if not I believe it is an addition to `cargo` that should be made.

Imagine your crate `whatever` has `dep_1`, `dep_2`, `dep_3` and `dep_4` as dependencies. `dep_1` could be replaced by `dep_2`. `dep_3` is optional, and so is `dep_4`.

You could have any of the following lists as dependencies:
- `dep_1` (some functionality from `dep_3` and `dep_4` missing)
- `dep_2` (same)
- `dep_1` and `dep_3` (some functionality from `dep_4` missing)
- `dep_2` and `dep_4` (some functionality from `dep_3` missing)
- `dep_1` and `dep_3` and `dep_4` (all functionality available)
- `dep_1` and `dep_2` and `dep_3` (some redundant imports from `dep_1` and `dep_2`, some functionality from `dep_4` missing)
- ...

It is quite easy to set up feature flags to control which of the dependencies you use.
```
# Cargo.toml
...

[features]
default = ["use-1"]
use-1 = ["dep_1"]
use-2 = ["dep_2"]
use-3 = ["dep_3"]
use-4 = ["dep_4"]

[dependencies]
dep_1 = { version = "*", optional = true }
...
dep_4 = { version = "*", optional = true }

...
```
With this you can add `dep_3` with `cargo build --features use-3` and you can switch out `dep_1` for `dep_2` and add `dep_4` with `cargo build --no-default-features --features use-2 use-4`.

So far, so good.

Now let's add licenses for `dep_1`, `dep_2`, and `dep_3`:
- `dep_1` is MIT/Apache2
- `dep_2` is GPLv2
- `dep_3` is GPLv3
- `dep_4` is GPLv2 or later

Let's say you would like to publish your crate under MIT because you don't want all of crates.io to be infected by GPL.

Unfortunately :
- `dep_2`, `dep_3` and `dep_4` prevent you from licensing under MIT
- `dep_2` is not compatible with `dep_3`

Therefore I believe we need something along the lines of the following:
```
# Cargo.toml
...
license = "MIT OR GPLv2 OR GPLv3"

[features]
default = ["use-1"]
use-1 = ["dep_1"]
use-2 = ["dep_2", { license = "GPLv2" }]
use-3 = ["dep_3", { license = "GPLv3" }]
use-4 = ["dep_4", { license = "GPLv2 OR GPLv3" }]

...
```

This would indicate that if the feature flag `"use-2"` is on (i.e. if `dep_2` is used), then the license has to be `"GPLv2"`, if `"use-3"` is on then it has to be `"GPLv3"`, and if neither of these features are on then the license can be `"MIT OR GPLv2 OR GPLv3"` as indicated in the `[package]` section.

In short, the `license` attribute of the `[package]` section would become a default license that feature flags can override. Dependent crates would then be limited to a subset of licenses depending on the features they require. Any crate that needs a feature from `"dep_4"` could not use the `"MIT"` and would be restricted to `"GPLv2 or later"`

Of course the proposed syntax above is only a draft. Once again, if there is already a solution to this I would love to hear about it, but I have searched the Rust forum, the `cargo-lichking` repository, and this subreddit without finding anything.

Do you have better ideas ? Or have you experienced situations where this could have made things easier for you ?

_Side note: I really don't want this discussion to turn into a war about which license to choose, I think we can all agree that whichever license we believe is best we all ought to be able to better choose the license under which we want to distribute parts of our code._
