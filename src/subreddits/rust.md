# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (36/2020)!](https://www.reddit.com/r/rust/comments/ijvwsk/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ijvwsk/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (36/2020)?](https://www.reddit.com/r/rust/comments/ijvxwz/whats_everyone_working_on_this_week_362020/)
- url: https://www.reddit.com/r/rust/comments/ijvxwz/whats_everyone_working_on_this_week_362020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-36-2020/48100?u=llogiq)!
## [3][Supporting Linux kernel development in Rust](https://www.reddit.com/r/rust/comments/ik42b4/supporting_linux_kernel_development_in_rust/)
- url: https://lwn.net/SubscriberLink/829858/281103f9c6fd0dc2/
---

## [4][Are you interested in a tutorial series on making your own language using Rust?](https://www.reddit.com/r/rust/comments/ikiag7/are_you_interested_in_a_tutorial_series_on_making/)
- url: https://www.reddit.com/r/rust/comments/ikiag7/are_you_interested_in_a_tutorial_series_on_making/
---
Hi,

Would anyone be interested in a tutorial series on creating an interpreted language using Rust? My plan is to start the tutorials as simple and easy to follow as possible, and gradually work upwards from there to improve the language. I’ve only created one language (still WIP at the moment), so I’m not very experienced in this area, but I have a strong grasp on the fundamentals as a result of rewriting my pet language from scratch many times. Creating the series will hopefully also be a learning experience for me :)

At first I thought that the tutorials should be in text form, but then I remembered how much more engaged and interested I am when watching tutorial content in video form. I believe that this may also be easier to make for me, as I find things easier to explain with speech rather than writing. However, I’ve heard quite a few people say they dislike video content, as it doesn’t allow them to skim and slows them down. Another reason that text might be better is because I’m 16, so it would probably be strange watching a tutorial where the person teaching isn’t an adult (I know I would find it strange).

So, would you be interested?

[View Poll](https://www.reddit.com/poll/ikiag7)
## [5][Why can't we declare fields of type Box&lt;impl Trait&gt;?](https://www.reddit.com/r/rust/comments/ikjgl6/why_cant_we_declare_fields_of_type_boximpl_trait/)
- url: https://www.reddit.com/r/rust/comments/ikjgl6/why_cant_we_declare_fields_of_type_boximpl_trait/
---
Why am I forced to use dynamic dispatch when boxing a trait?

Consider this code:

`struct S { a: Box&lt;impl Trait&gt; }`

It doesn't compile, I have to write it like:

`struct S&lt;T:Trait&gt; { a: Box&lt;T&gt; }`

Imo it should not require specifying a generic parameter each time I use this struct.

Same for assigning variables - why `Box::new(something)` should be assigned to a `Box&lt;dyn Trait&gt;` if it could easily be stored as `Box&lt;impl Trait&gt;`?

&amp;#x200B;

Another example:

    fn test() -&gt; Box&lt;impl Trait&gt; {
        Box::new(1)
    }
    
    trait Trait{
        
    }
    
    impl Trait for i32{}
    
    fn main() {
        let z: Box&lt;impl Trait&gt; = Box::new(1); // compiler error,should be dyn
        let z/*:Box&lt;impl Trait&gt;*/ = test(); // implicit static dispatch
    }

&amp;#x200B;
## [6][Avoid Build Cache Bloat By Sweeping Away Artifacts](https://www.reddit.com/r/rust/comments/iki036/avoid_build_cache_bloat_by_sweeping_away_artifacts/)
- url: https://www.justanotherdot.com/posts/avoid-build-cache-bloat-by-sweeping-away-artifacts.html
---

## [7][.take() is my favorite method.](https://www.reddit.com/r/rust/comments/ik3npj/take_is_my_favorite_method/)
- url: https://www.reddit.com/r/rust/comments/ik3npj/take_is_my_favorite_method/
---
I discovered today that .take() allows me to take ownership of objects inside options and this changed everything for me. It's just one of those things where rust found the perfect balance between convenience and safety.
## [8][Should library crate also implement CLI?](https://www.reddit.com/r/rust/comments/ike2qx/should_library_crate_also_implement_cli/)
- url: https://www.reddit.com/r/rust/comments/ike2qx/should_library_crate_also_implement_cli/
---
I'm building a library which provides some functionality to users. I also want to build a CLI for that library, so those users, who don't care about writing their own programs, can use that functionality just by running CLI tool through command line and Bash.

What is idiomatic way of implementing this kind of library? Should CLI be part of the library crate? Or should CLI be a separate create that takes dependency on the library?

Can you please recommend some good real-world example of this kind of crate?
## [9][Thought I'd share this demo of a retro-futuristic UI in rust.](https://www.reddit.com/r/rust/comments/ik04zr/thought_id_share_this_demo_of_a_retrofuturistic/)
- url: https://ivanceras.github.io/futuristic-ui/
---

## [10][Why don't the methods into() and try_into support the turbofish?](https://www.reddit.com/r/rust/comments/ikahj4/why_dont_the_methods_into_and_try_into_support/)
- url: https://www.reddit.com/r/rust/comments/ikahj4/why_dont_the_methods_into_and_try_into_support/
---
Methods like `collect()` and `parse()` allow you to specify the type you're converting to and sometimes require you to specify the type you're converting to using the [turbofish](https://turbo.fish/).

`into()` and `try_into()` are both very similar. Sometimes they know the type you want to convert to, but often they don't. However they do not allow usage of the turbofish.

For a variable assignment it doesn't make much of a difference, you can just do:

    let foo: u8 = 0;
    let upcasted: usize = u8.into();

But for more complex expressions, you aren't given the option:

    let foo: u8 = 0;
    let items: [u32; 4] = [4, 5, 6, 7];
    let item = items[foo.into()];  // throws compile error, requires usize, not u8
    let item = items[foo.into::&lt;usize&gt;()];  // not valid, throws compile error

And the only option here is to use as:

    let item = items[foo as usize];

Which works, but feels less ergonomic. I suppose a lot of my complaints would go away if I could index into arrays and pals using number types other than usize, but that aside, I feel like into could really benefit from the turbofish.

Is there a reason `into()` and `try_into()` don't have it?
## [11][rust-analyzer changelog #40](https://www.reddit.com/r/rust/comments/ijwdcc/rustanalyzer_changelog_40/)
- url: https://rust-analyzer.github.io/thisweek/2020/08/31/changelog-40.html
---

## [12][Building a modular Ruby with traits and crates](https://www.reddit.com/r/rust/comments/ikk86o/building_a_modular_ruby_with_traits_and_crates/)
- url: https://twitter.com/artichokeruby/status/1300786903768203266
---

