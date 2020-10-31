# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (44/2020)!](https://www.reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/jdwuis/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 362](https://www.reddit.com/r/rust/comments/jk35ha/this_week_in_rust_362/)
- url: https://this-week-in-rust.org/blog/2020/10/28/this-week-in-rust-362/
---

## [3][rkvm - A tool to share keyboard and mouse over the network on Linux machines](https://www.reddit.com/r/rust/comments/jlhga1/rkvm_a_tool_to_share_keyboard_and_mouse_over_the/)
- url: https://github.com/htrefil/rkvm
---

## [4][Announcing Octane, a high performance web server modelled after express](https://www.reddit.com/r/rust/comments/jl56nq/announcing_octane_a_high_performance_web_server/)
- url: https://www.reddit.com/r/rust/comments/jl56nq/announcing_octane_a_high_performance_web_server/
---
[https://github.com/OctaneWeb/Octane](https://github.com/OctaneWeb/Octane)  


How is octane different?

1. Lightweight HTTP implementation is designed to be fast and correct
2. 100% safe
3. A low dependency web server
4. Is scalable: Octane's architecture allows other crates to integrate with it seamlessly, allowing many custom middleware implementations
5. Has all the good of express and safety of rust, tries to be very intuitive and easy to use
6. Features things like extended queries, plain queries, \`path!\` macros out of the box
7. SSL is easy to set up and easy to switch too, you will not have to change any existing rust code in order to switch between OpenSSL or rustls
## [5][Announcing colo, a terminal color viewer](https://www.reddit.com/r/rust/comments/jld6cr/announcing_colo_a_terminal_color_viewer/)
- url: https://www.reddit.com/r/rust/comments/jld6cr/announcing_colo_a_terminal_color_viewer/
---
[colo](https://github.com/Aloso/colo) is a small CLI tool that can show colors and convert them between various color spaces. It supports RGB, CMY, CMYK, HSV, HSL, LCH, LUV, CIELAB, Hunter lab, CIE 1931 XYZ and CIE YXY. Big thanks for the [color\_space](https://docs.rs/color_space/0.5.3/color_space/) crate, which is used for the conversions!

https://preview.redd.it/ryoqg2iebdw51.png?width=602&amp;format=png&amp;auto=webp&amp;s=7fa6fc80c1cfe065c2c62d132d826052ba75ae26
## [6][Beginners - Rust Variables](https://www.reddit.com/r/rust/comments/jlh2y1/beginners_rust_variables/)
- url: https://www.golangprograms.com/rust-variables.html
---

## [7][Parse XML with NIF (Using Rust)](https://www.reddit.com/r/rust/comments/jlfy4d/parse_xml_with_nif_using_rust/)
- url: https://youtu.be/lm_6WzWazjU
---

## [8][Newbie here: just squashed my first bug, but learned that I don't understand mutable references. Please help me](https://www.reddit.com/r/rust/comments/jlf1vj/newbie_here_just_squashed_my_first_bug_but/)
- url: https://www.reddit.com/r/rust/comments/jlf1vj/newbie_here_just_squashed_my_first_bug_but/
---
I've been following along with *The Rust Programming Language* by Steve Klabnik and Carol Nichols et al, and up until the end of Chapter 3 I was following along really well. Then I tried to make a temperature converter between Fahrenheit and Celcius.

**The Bug:** The program compiles, but only works correctly if the user's input is successfully parsed on the first loop. For instance, entering "46" returns a value of "7.77...8" and then exits the program, which is exactly what I expected. Entering "asdf" prompts the user to retry, but now entering "46" prompts the user to retry *again,* even though it should parse correctly on the second try. This results in an endless loop where the user tries number after number and only ever gets more tries instead of an actual result, all because they messed up the input the first time.

**The Code:**

    let f2c = true;
    let mut input = String::new();

    loop {
    if f2c {
        println!("Current Mode: Fahrenheit to Celcius");
    }
    else {
        println!("Current Mode: Celcius to Fahrenheit");
    }

    io::stdin().read_line(&amp;mut input)
        .expect("Could not read line!");
    let mut temp: f64 = match input.trim().parse() {
        Ok(num) =&gt; num,
        Err(_) =&gt; continue,
    };

    temp = convert(f2c, temp);

    println!("&gt; {}", temp);
    break;
    }

**The Patch:** I realized that my variable "input" wasn't being reset at the end of the loop, so I moved its initialization into the body of the loop, just before my io::stdin().read_line...

...and it worked! My code now works properly. The "input" variable now goes out of scope at the end of the loop and is reborn each time the user inputs something. Because it is a "new" variable, my read_line() function can change it. But...

**The Question:** ...why is it that, even though I'm feeding read_line() a mutable reference, it's unable to change it the second time? I read ahead a bit in the book and found out that mutable references can only be changed one time, but the book also says that changing the scope can circumvent that and allow multiple changes (if not simultaneous ones).

Doesn't the loop { } change the scope? If I'm only changing the variable once per loop, doesn't that satisfy both requirements?

This is breaking my brain.
## [9][Rust Binary Tree: A Refactor](https://www.reddit.com/r/rust/comments/jlgpm1/rust_binary_tree_a_refactor/)
- url: https://rossketeer.medium.com/rust-binary-tree-a-refactor-1b090a88e24
---

## [10][Rust Design-for-Testability: a survey](https://www.reddit.com/r/rust/comments/jl2xlg/rust_designfortestability_a_survey/)
- url: https://alastairreid.github.io/rust-testability/
---

## [11][These Weeks In Actix | Sep-Oct '20](https://www.reddit.com/r/rust/comments/jkv5xu/these_weeks_in_actix_sepoct_20/)
- url: https://www.reddit.com/r/rust/comments/jkv5xu/these_weeks_in_actix_sepoct_20/
---
Since the [release of Actix Web v3.0](https://www.reddit.com/r/rust/comments/iqq8k9/announcing_actixweb_v30/) in September, we've been very happy to see the positive response to the numerous internal safety improvements. The uptake of v3 has been swift; over 50% of crate downloads are for the newer versions.

# Notable Changes

## actix-web v3.1 &amp; v3.2

* New `#[route(...)]` multi-method macro.
* A request-local data extractor `web::ReqData`.
* Enable access to client TLS certificates and other advanced use cases where extra socket data is needed (via `HttpServer::on_connect` hook).
* An `exclude_regex` method for the Logger middleware.
* Register function for custom request-derived logging units.

&gt;[Full Changelog](https://github.com/actix/actix-web/blob/HEAD/CHANGES.md)

&amp;#x200B;

## actix-cors v0.5

On the path to stabilization, v0.5 is being considered a v1.0 beta. Includes a new way to filter allowed origins using closures and a major internal rewrite. Any feedback is appreciated.

&gt;[Full Changelog](https://github.com/actix/actix-extras/blob/HEAD/actix-cors/CHANGES.md)

&amp;#x200B;

# An Official Discord Community

&gt;Server Invite: [https://discord.gg/NWpN5mmg3x](https://discord.gg/NWpN5mmg3x)

We're still keeping Gitter around as a support channel (and I'm hopeful it will improve more quickly under the Matrix banner).
## [12][With the new Sifive Unmatched, Risc-V is within reach of consumers or at least power linux users. What is the status of Rust for risc-v?](https://www.reddit.com/r/rust/comments/jlgkdb/with_the_new_sifive_unmatched_riscv_is_within/)
- url: https://www.reddit.com/r/rust/comments/jlgkdb/with_the_new_sifive_unmatched_riscv_is_within/
---
Im thinking about getting a board and help out with testing and or porting. https://youtu.be/VXqQuxmB76M 

What is the status of risc-v support?
Thanks
