# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (18/2020)!](https://www.reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g4nu6/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 336](https://www.reddit.com/r/rust/comments/gae1nt/this_week_in_rust_336/)
- url: https://www.reddit.com/r/rust/comments/gae1nt/this_week_in_rust_336/
---
[https://this-week-in-rust.org/blog/2020/04/29/this-week-in-rust-336/](https://this-week-in-rust.org/blog/2020/04/29/this-week-in-rust-336/)
## [3][Rust/WinRT Public Preview](https://www.reddit.com/r/rust/comments/gb049m/rustwinrt_public_preview/)
- url: https://blogs.windows.com/windowsdeveloper/2020/04/30/rust-winrt-public-preview/
---

## [4][Can i make Pokemon Gold in 10 minutes with Oxygengine? No, that's not even remotely possible.](https://www.reddit.com/r/rust/comments/gbahqx/can_i_make_pokemon_gold_in_10_minutes_with/)
- url: https://v.redd.it/uz9rh4ghm2w41
---

## [5][Zooming in on Observability with Rust and eBPF](https://www.reddit.com/r/rust/comments/gbghok/zooming_in_on_observability_with_rust_and_ebpf/)
- url: https://blog.redsift.com/labs/zooming-in-on-observability/
---

## [6][How to write a GUI for a async codebase?](https://www.reddit.com/r/rust/comments/gbeexl/how_to_write_a_gui_for_a_async_codebase/)
- url: https://www.reddit.com/r/rust/comments/gbeexl/how_to_write_a_gui_for_a_async_codebase/
---
I did some research and I'm a bit concerned that this might be not as easy.

I have a codebase I'm experiementing with [(see here)](https://git.sr.ht/~matthiasbeyer/distrox) that has a async API (mainly [this](https://git.sr.ht/~matthiasbeyer/distrox/tree/master/src/app.rs#L23)).

I want to start writing a GUI for the app. I'm not sure which framework to use (even which approach - native GUI or web UI - the app is single-user and supposed to run on a users device).

Because all my interfaces for the application code will be `async`, I tend to go for web UI because there are already `async` (some). But I am not sure whether I'm missing something, so I figured to ask here whether there's a way to write a native GUI application with `async`.

Thanks for any replies.
## [7][Pushrod - Rust Native UI with SDL2 - Still going!](https://www.reddit.com/r/rust/comments/gb7lre/pushrod_rust_native_ui_with_sdl2_still_going/)
- url: https://www.reddit.com/r/rust/comments/gb7lre/pushrod_rust_native_ui_with_sdl2_still_going/
---
Hey all,

Happily, I'm still working on Pushrod; had a few events take place that made me switch focus for a few months, but I'm back on track.

As I've been developing this library, a few things have come to light; mainly the design of the library itself.

I have been focused on trying to make the code work with strictly using callback mechanisms to handle features, but I've found this to be quite the limitation.  Mainly because the current object, and its referencing objects cannot be shared, and there is a possibility of locked access at the time of callback.

Therefore, I'm going to be switching the library design over to an event-based system.

My question is this: if you were to use a library built from the ground up for UI, would you prefer it to be an event-based system, or would you prefer callbacks (via inline functions and closures)?

I will eventually do a writeup on my blog to explain the decision going forward, but I wanted to hear from you to see what your thoughts were on the subject.

As always, my [library is available at this link](https://www.github.com/KenSuenobu/rust-pushrod), but the library is being split into subsections so that the library as a whole is easier to maintain.

Looking forward to any responses!  Thank you all!
## [8][Best way to pass shared data/service to Warp handlers (or threads in general)?](https://www.reddit.com/r/rust/comments/gbdwdx/best_way_to_pass_shared_dataservice_to_warp/)
- url: https://www.reddit.com/r/rust/comments/gbdwdx/best_way_to_pass_shared_dataservice_to_warp/
---
Hello there,

I'm pretty new to Rust, with a background of C++, Go and  C#.

 I'm currently trying to implement basic RESTful service using Warp. Usually, in the above mentioned languages, I use a very similar pattern when writing services:

*handler/controller* \-&gt; *uses a service* \-&gt; *which uses a data access layer (DAL)*

Pretty much standard. So, this service (a struct in my case, containing the DAL, ans couple of methods ) must be passed to all Warp handlers functions.

I've read Warp doc, and this seem possible using some magic which should look like this:

 `fn with_service(`  
    `service: api::service::Service,`  
`) -&gt; impl Filter&lt;Extract = (api::service::Service,), Error = Infallible&gt; + Clone {`  
    `warp::any().map(move || service.clone())`  
`}`

 And now, it appears a requirement is to have this service Cloneable, if I understand Rust correctly, ok:

`#[derive(Clone)]`  
`pub struct Service {`  
    `repo_provider: dal::ItemRepoProvider,`  
`}`

Finally it's justed with wrap this way:

`warp::path!("my-resource")`  
        `.and(warp::get())`  
        `.and(with_service(service.clone()))`  
        `.and_then(api::get_all_resources);`

This is where I'm starting to question many things... Because my service itself, at this moment, is literally just forwarding to the DAL, which does MongoDB access. There is no state/cache yet. I assume the MongoDB driver itself manages connection pool and is thread safe, but I have no clue, it seems to be in alpha.

So why do I need to "clone" my service? Can I, instead, let's say somehow tell Rust it's ok to access it from threads, by just passing a reference? For instance, I am aware of Arc, could I wrap it in a Arc and then, clone this arc so that Warp and the Rust compiler are happy? Maybe I'll need an additional Mutex to protect the DAL if compiler complains, I have not tried yet...

In Go or C#, for instance, my service would be a reference, single instance, passed somehow to the controller/handler, and I would simply protect some of its state, like cache, by a mutex. And of course compilers don't care whether I do it correctly or not.

This problem is not specific to Warp I guess, and and can be asked for any web framework and thread spawn logic.

Basically: what is the best way, or the idiomatic way, to share a "service" abstracting data access in HTTP handlers/closures? And does the answer change if the service has a state, like a cache?
## [9][Why care about Minimum Rust Version?](https://www.reddit.com/r/rust/comments/gb1bhx/why_care_about_minimum_rust_version/)
- url: https://www.reddit.com/r/rust/comments/gb1bhx/why_care_about_minimum_rust_version/
---
A number of Rustaceans (including many I deeply respect) care a lot about ensuring that their code continues to compile on fairly old versions of Rust.  Given the number of good programmers who care, there _must_ be a good reason for valuing that sort of compatibility, but I haven't been able to figure out what it is.  I've come up with a few potential advantages of maintaining compatibility with old compilers, but none seem significant:

* Allowing a very casual user of Rust (i.e., one who has not researched the language enough to have learned about Rustup) to compile the code after running `apt install rust` or similar from an LTS distribution.

* Supporting enterprise users who are dealing with a (significantly misguided) corporate security policy that forces them to run an extremely out-of-date version of Rust (I don't _think_ this really happens anymore, but maybe I'm mistaken?)

* Supporting users who have found a way to write code that (despite Rust's famously strong backwards compatibility guarantees) does not compile on a newer version of Rust.

All of those strike me as implausible, but I'm sure there is a good reason.  So am I missing a better reason, or is there some reason on of those isn't as implausible as it seems to be?
## [10][Help with llvm_sys?](https://www.reddit.com/r/rust/comments/gbcdv0/help_with_llvm_sys/)
- url: https://www.reddit.com/r/rust/comments/gbcdv0/help_with_llvm_sys/
---
After sifting through many doc pages and articles, I still have no idea what is causing my code to hang:

```rust
let pm = LLVMCreatePassManager();
llvmt::scalar::LLVMAddInstructionCombiningPass(pm);

assert!(
    LLVMRunPassManager(pm, self.module) == 1,
    "Optimization failed!"
);
LLVMDisposePassManager(pm);
```

It is supposed to run an optimization pass on an LLVM IR module. If I use any other pass it runs fine; the full code it was taken from used the SROA, early CSE, reassociation, GVN, and CFG simplification passes in that order. However, once I try adding the instruction combining pass into the mix, no matter where I add it in, the program freezes with 100% CPU when LLVMRunPassManager is called. When I try running this pass using LLVM's `opt` tool over the same piece of IR, it works fine. I've also tried copying parts of the PassManagerBuilder-related code from the `opt` tool to build a standard O3 pipeline instead, which displays the exact same 100% cpu freeze. 

Is there something wrong with how I'm setting it up or using the passes?
## [11][Trait type aliases](https://www.reddit.com/r/rust/comments/gbhkw9/trait_type_aliases/)
- url: https://www.reddit.com/r/rust/comments/gbhkw9/trait_type_aliases/
---
Is there any way to deal with this monstrosity?

    pub trait Vec3Operations&lt;T&gt; {
      
      
            fn dot&lt;'a, 'b: 'a&gt;(&amp;'a self, other: &amp;'b Self) -&gt; &lt;&lt;&lt;&lt;&amp;'a Self as XYZ&gt;::Item as std::ops::Mul&lt;&lt;&amp;'a Self as XYZ&gt;::Item&gt;&gt;::Output as std::ops::Add&lt;&lt;&lt;&amp;Self as XYZ&gt;::Item as std::ops::Mul&lt;&lt;&amp;'a Self as XYZ&gt;::Item&gt;&gt;::Output&gt;&gt;::Output as std::ops::Add&lt;&lt;&lt;&amp;'a Self as XYZ&gt;::Item as std::ops::Mul&gt;::Output&gt;&gt;::Output
      
      
            where &amp;'a Self: XYZ,
      
      
            &lt;&amp;'a Self as XYZ&gt;::Item: ops::Mul&lt;&lt;&amp;'a Self as XYZ&gt;::Item&gt;,
      
      
            &lt;&lt;&amp;'a Self as XYZ&gt;::Item as std::ops::Mul&lt;&lt;&amp;'a Self as XYZ&gt;::Item&gt;&gt;::Output: ops::Add&lt;&lt;&lt;&amp;'a Self as XYZ&gt;::Item as std::ops::Mul&lt;&lt;&amp;'a Self as XYZ&gt;::Item&gt;&gt;::Output&gt;,
      
      
            &lt;&lt;&lt;&amp;'a Self as XYZ&gt;::Item as std::ops::Mul&lt;&lt;&amp;'a Self as XYZ&gt;::Item&gt;&gt;::Output as std::ops::Add&lt;&lt;&lt;&amp;'a Self as XYZ&gt;::Item as std::ops::Mul&gt;::Output&gt;&gt;::Output: ops::Add&lt;&lt;&lt;&amp;'a Self as XYZ&gt;::Item as std::ops::Mul&lt;&lt;&amp;'a Self as XYZ&gt;::Item&gt;&gt;::Output&gt;
      
      
            {
      
      
                self.x() * other.x() + self.y() * other.y() + self.z() * other.z()
      
      
            }

It's fun and it's working, but it's unreadable and I don't know what to do with this.

[Full code](https://github.com/Lex98/raytracer)

Ps It would be great if someone could review my code.
## [12][slip 0.1.0: protect your error strings against reverse-engineering](https://www.reddit.com/r/rust/comments/gayjvz/slip_010_protect_your_error_strings_against/)
- url: https://github.com/Moxinilian/slip
---

