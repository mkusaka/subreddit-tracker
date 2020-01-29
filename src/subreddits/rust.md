# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (5/2020)!](https://www.reddit.com/r/rust/comments/eupam3/hey_rustaceans_got_an_easy_question_ask_here_52020/)
- url: https://www.reddit.com/r/rust/comments/eupam3/hey_rustaceans_got_an_easy_question_ask_here_52020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (5/2020)?](https://www.reddit.com/r/rust/comments/eupbbn/whats_everyone_working_on_this_week_52020/)
- url: https://www.reddit.com/r/rust/comments/eupbbn/whats_everyone_working_on_this_week_52020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-5-2020/37515?u=llogiq)!
## [3][Jonathan Blow on Rust](https://www.reddit.com/r/rust/comments/evgl81/jonathan_blow_on_rust/)
- url: https://www.reddit.com/r/rust/comments/evgl81/jonathan_blow_on_rust/
---
Podcast link: https://oxide.computer/blog/on-the-metal-9-jonathan-blow/

Submission statement: Bryan Cantrill (of the ["Is It Time to Rewrite the Operating System in Rust?" presentation](https://www.infoq.com/presentations/os-rust/)) interviewed Jonathan Blow, developer of [Braid](https://en.wikipedia.org/wiki/Braid_\(video_game\)) and the Jai compiler on a podcast called On The Metal. The whole thing is engrossing, but there's a quote in there regarding Rust I wanted to get opinions on from the /r/rust crowd at around 1:40:50:

&gt; Blow: And so Rust has a good set of ingredients there. **The problem that I have with it is when I'm working on really hard stuff, I don't exactly know what I'm doing for a long time. And so if the cost of experimentation is driven too high, it actually impairs my ability to get work done.**
&gt; 
&gt; Blow: And so my approach that we're doing in the language that we're working on is different. It's like, you have very very very extensive metaprogramming facilities and you can use those to build your own correctness checking for your program that you traditionally would have needed to make a compiler extension to do, which is pretty far away from what people do day to day and they don't end up doing it even if they could have. So what happens is your metaprogram gets information that normally would only be internal to the compiler.

He goes on to describe a use of that at 1:45:34:

&gt; Blow: So in games the way that's traditionally handled is there's a very natural barrier which is the end of the frame, right? So we do a bunch of stuff over and over 60 times per second or 200 times per second, whatever. And at the end of the frame, we pretty much know it's all... 
&gt; 
&gt; Cantrill: -garbage now.
&gt; 
&gt; Blow: Well, we're not really hanging on to anything. So it's very clean cut you could say. So, well, it's a very reasonable thing to say it's fine to do that stuff as long as you don't hold an entity pointer across a frame boundary, okay? Now that's not a very generalizable rule. Because what's an entity pointer versus a pointer to something else? Your compiler doesn't know that right? But you know that. And so when the new guy shows up or the summer intern shows up and writes some code where he puts that entity pointer into a data structure that...
&gt; 
&gt; Cantrill: -*is* surviving-
&gt; 
&gt; Blow: ...some hash table that's sitting around, that's going to be a problem, right? And that's a very practical, real kind of problem, it's not an academic, made up problem. You actually can solve that - your metaprogram can say well, you put an entity pointer in this data structure and that data strucutre is not in this whitelist that I have over here and so, error, don't compile, right? And so you can start solving the very specific problems that you have without introducing general friction. Now, is that better than the Rust approach? I don't know, I think it will be but I certainly offer you no proof of that statement.
## [4][We wrote about our first embedded Rust project (for a real client!). It is about piano technology.](https://www.reddit.com/r/rust/comments/evkkaf/we_wrote_about_our_first_embedded_rust_project/)
- url: https://jitter.company/blog/2020/01/28/measuring-space-time-behaviours-of-piano-keys-with-rust/
---

## [5][Nushell 0.9.0 - lots of improvements, new commands, and the return of shelling out.](https://www.reddit.com/r/rust/comments/evhzcn/nushell_090_lots_of_improvements_new_commands_and/)
- url: https://www.nushell.sh/blog/2020/01/28/nushell-0_9_0.html
---

## [6][thirtyfour - A new selenium library for Rust, for automated website testing](https://www.reddit.com/r/rust/comments/evlc49/thirtyfour_a_new_selenium_library_for_rust_for/)
- url: https://www.reddit.com/r/rust/comments/evlc49/thirtyfour_a_new_selenium_library_for_rust_for/
---
Last year I decided to build a selenium library for Rust. I've been working with the python one for several years and I thought it would be an interesting challenge to build a similar library for Rust. I learned a lot about selenium in the process, and I'm pretty happy with how it has turned out.

Github: [https://github.com/stevepryde/thirtyfour](https://github.com/stevepryde/thirtyfour)

Crate: [https://crates.io/crates/thirtyfour](https://crates.io/crates/thirtyfour)

  
I've aimed to implement the full feature set of the W3C WebDriver spec, but no doubt there are a few bits missing around the edges.

Here's the feature list: 

* Async / await support
* Synchronous support
* Create new browser session via Selenium Standalone or Grid
* Automatically close browser session on drop
* All W3C WebDriver and WebElement methods supported
* Find elements (via all common selectors)
* Send keys to elements, including key-combinations
* Execute Javascript (needs more testing)
* Action Chains
* Cookies
* Switch to frame/window/element/alert
* Alert support
* Capture / Save screenshot of browser or individual element

Up to this point it has been a personal project so my time is limited but I'd be happy to fix any bugs if you run into any (raise issues on github please). One point of note is that this library does not aim to be a full-featured automation suite. Just like the python library, the goal is to provide the base functionality on which larger high-level frameworks can be built.

I hope you find it useful.
## [7][How Zebra used Tower to replace Zcash's legacy Bitcoin C++ networking code with modern, async Rust](https://www.reddit.com/r/rust/comments/evc3ug/how_zebra_used_tower_to_replace_zcashs_legacy/)
- url: https://www.zfnd.org/blog/a-new-network-stack-for-zcash/
---

## [8][My Rust sticker finally arrived!](https://www.reddit.com/r/rust/comments/evhcp6/my_rust_sticker_finally_arrived/)
- url: https://i.redd.it/fqxanfnc4nd41.png
---

## [9][Use ? operator on a vec that returns result?](https://www.reddit.com/r/rust/comments/evmel0/use_operator_on_a_vec_that_returns_result/)
- url: https://www.reddit.com/r/rust/comments/evmel0/use_operator_on_a_vec_that_returns_result/
---
Below is some pseudo code. I am looping through a vec. Within the loop, I am transforming the type to another type. However, an error can occur. When that happens, I want to return as early as possible, so if an element fails to transform, I want to return a custom error. I can not get it to work, an I doubt if I can use the ? operator inside a closure.

&amp;#x200B;

This is my non compiling code. If this code would compile, it would still complete the loop before returning an error, even if the first element can not be parsed. An alternative would be initializing a mutable vec, but I hope I can do it in a functional way.

    struct MyErr;
    
    fn transform(v: Vec&lt;String&gt;) -&gt; Result&lt;Vec&lt;i32&gt;, MyErr&gt; { 
        let my_vec_i32 = v.into_iter().map(|s| s.parse().map_err(|_| MyErr)).collect()?;
        // Work with my_vec...
        return Ok(my_vec_i32) 
    }
## [10][Mapping the power of Mozilla’s Rebel Alliance – The Mozilla Blog](https://www.reddit.com/r/rust/comments/evah7y/mapping_the_power_of_mozillas_rebel_alliance_the/)
- url: https://blog.mozilla.org/blog/2020/01/27/mapping-the-power-of-mozillas-rebel-alliance/
---

## [11][Pokemon Gen 8 dataset](https://www.reddit.com/r/rust/comments/eve09h/pokemon_gen_8_dataset/)
- url: https://www.reddit.com/r/rust/comments/eve09h/pokemon_gen_8_dataset/
---
[https://github.com/blargg/pokemon](https://github.com/blargg/pokemon)

Posting here in case anyone might find this useful, or just wanted to play around with it. It's provided as is.

This is a data set for the Gen 8 pokemon and moves, available as basic rust types. It should have all their game stats available, but there may be some mistakes.

Personally, I've been using this to help build teams. This allows searches like "get all the pokemon that can learn this move set, and sort them by base stats". I haven't been able to find another program that allows for that much specificity, so I wrote this for fun.
## [12][Way Cooler postmortem, a tiling window manager for Wayland written in Rust](https://www.reddit.com/r/rust/comments/evmp27/way_cooler_postmortem_a_tiling_window_manager_for/)
- url: http://way-cooler.org/blog/2020/01/09/way-cooler-post-mortem.html
---

