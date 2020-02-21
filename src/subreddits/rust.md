# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (8/2020)!](https://www.reddit.com/r/rust/comments/f5413m/hey_rustaceans_got_an_easy_question_ask_here_82020/)
- url: https://www.reddit.com/r/rust/comments/f5413m/hey_rustaceans_got_an_easy_question_ask_here_82020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/f1ucwh/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 326](https://www.reddit.com/r/rust/comments/f6sy2s/this_week_in_rust_326/)
- url: https://this-week-in-rust.org/blog/2020/02/18/this-week-in-rust-326/
---

## [3][Made a rusty rust wallpaper. The other image was larger than 20 MB (24 Mb) and I don't know how to compress PNG images in photoshop.](https://www.reddit.com/r/rust/comments/f796ds/made_a_rusty_rust_wallpaper_the_other_image_was/)
- url: https://i.redd.it/9t12lnng69i41.png
---

## [4][Update about rustbus, (one of the) pure rust dbus libraries](https://www.reddit.com/r/rust/comments/f7a5qc/update_about_rustbus_one_of_the_pure_rust_dbus/)
- url: https://www.reddit.com/r/rust/comments/f7a5qc/update_about_rustbus_one_of_the_pure_rust_dbus/
---
Heya r/rust!

I had posted a while ago about my dbus library [Repo link](https://github.com/KillingSpark/rustbus). 

I have since, with the help of PRs (mostly by [okeri](https://github.com/okeri)), fixed hopefully all bugs and would love some feedback on the API. Maybe there is a better way to convert the rust types into the type enums I use for the dbus marshalling that I missed?

This lib is somewhat similar to libdbus in that it gives the user pretty lowlevel access to the connection and the messages sent. The next natural step would be to build a generic proxy-library on top of this for convenient RPC.


There are currently some other efforts in this direction [as collected by the dbus-rs crate maintainer](https://github.com/diwic/dbus-rs/issues/100#issuecomment-580411226):

* https://gitlab.freedesktop.org/zeenix/zbus/
* https://github.com/Arnavion/dbus-pure
* https://github.com/srwalter/dbus-bytestream

I have not yet done a comparison in features/speed/correctness but I might do that in the future.
## [5][Using Google Analytics from a Rust Wasm Front-End Web App](https://www.reddit.com/r/rust/comments/f7563j/using_google_analytics_from_a_rust_wasm_frontend/)
- url: https://medium.com/@matt_78276/using-google-analytics-from-a-rust-wasm-front-end-web-app-12840c420dad
---

## [6][Passthrough field access in unions like in C?](https://www.reddit.com/r/rust/comments/f7ae2j/passthrough_field_access_in_unions_like_in_c/)
- url: https://www.reddit.com/r/rust/comments/f7ae2j/passthrough_field_access_in_unions_like_in_c/
---
Let's say I am writing an emulator and I have two 8 bit registers
(a and f) which can also be interpreted as one 16 bit register.

In C, this could be modeled as


    struct register {
        union {
            struct {
                unsigned char f;
                unsigned char a;
            };
            unsigned short af;
        };
    };


which is convenient since you can just use `.f` or `.a` to access or set the 8 bit registers
or `.af` to access or set the 16 bit register.

In Rust (which I am quite new to), I guess I can do something like:


    #[derive(Copy, Clone)]
    struct FA {
        f: u8,
        a: u8,
    }


    union Register {
        _fa: FA,
        fa: u16,
    }


but the syntax to access the `f` and `a` registers are here much more awkward since
I can't pass through the `.f` access to the `FA` struct. I can, of course, just store the
`a` and `f` registers as is and then write the methods `get_af(&amp;self)` and `set_af(&amp;mut self, v: u16)` that does the bit
shifts needed but it doesn't feel really clean (need to write getter and setters).

Is there any nice way to do this or should I just implement the bit shift methods?
## [7][Hank Green explains the Rust programming language](https://www.reddit.com/r/rust/comments/f6tpig/hank_green_explains_the_rust_programming_language/)
- url: https://www.reddit.com/r/rust/comments/f6tpig/hank_green_explains_the_rust_programming_language/
---
Maybe we will see a "Crash Course Rust?"

[https://www.youtube.com/watch?v=IwjlCxwcuIc](https://www.youtube.com/watch?v=IwjlCxwcuIc)
## [8][PKCS11 Engine for Rust TLS](https://www.reddit.com/r/rust/comments/f77eqk/pkcs11_engine_for_rust_tls/)
- url: https://www.reddit.com/r/rust/comments/f77eqk/pkcs11_engine_for_rust_tls/
---
Hi folks,

I have a SHM(Secure Hardware Module) which houses my certs and does the signing. I was wondering has anyone here ever worked with a PKCS11 Engine using Rust. The Rust-TLS library doesn't support it, and I was wondering about integrating PKCS11 engine with our TLS library. Has anyone done anything similar here so that I might have some context? Currently I am going through RUST-TLS code and have identified the traits that I would need to implement(Thanks to the author of that crate).
## [9][Is it possible to get rid of those lines? I am using the rust-analyzer VSCode plugin and I find those extra lines to be more distracting than helpful :/](https://www.reddit.com/r/rust/comments/f6yuql/is_it_possible_to_get_rid_of_those_lines_i_am/)
- url: https://i.redd.it/gk2wwhj2y4i41.png
---

## [10][[ANN] I wrote a tool in sauron to convert html to sauron view syntax](https://www.reddit.com/r/rust/comments/f79nf9/ann_i_wrote_a_tool_in_sauron_to_convert_html_to/)
- url: https://www.reddit.com/r/rust/comments/f79nf9/ann_i_wrote_a_tool_in_sauron_to_convert_html_to/
---
[tool site](https://ivanceras.github.io/html2sauron/)

[repo](https://github.com/ivanceras/html2sauron)

This is a WIP, there will be bugs.
## [11][Experience with Rust In Action?](https://www.reddit.com/r/rust/comments/f74mga/experience_with_rust_in_action/)
- url: https://www.reddit.com/r/rust/comments/f74mga/experience_with_rust_in_action/
---
I was just wondering if anyone has experience with [Rust In Action](https://www.manning.com/books/rust-in-action?query=rust)? Is it worth it, or would I be better off just with the [Rust book](https://doc.rust-lang.org/stable/book/).
## [12][STM32F0](https://www.reddit.com/r/rust/comments/f784wq/stm32f0/)
- url: https://www.reddit.com/r/rust/comments/f784wq/stm32f0/
---
Fellow no-std Rustaceans! :)

I'm pretty new at both #rustlang and embedded, so I must be asking some obvious questions, but please help me here.

My application at this point is just supposed to continuously update a counter and show it on an SSD1306 display. The update must be triggered every second by a timer. I understand more or less how it works:  [https://github.com/nebelgrau77/STM32F0\_blinky\_IRQ](https://github.com/nebelgrau77/STM32F0_blinky_IRQ)  \- this is based on an example from the STM32F0xx-HAL crate.

&amp;#x200B;

Here's my currently not working code:  [https://github.com/nebelgrau77/STM32F0-SSD1306-IRQ\_DRAFT](https://github.com/nebelgrau77/STM32F0-SSD1306-IRQ_DRAFT)

The problem is the following: in order to be able to display my constantly updated value, the display must be in a loop outside of the critical section, but it also has to be available out there (lines 100 - 112), so it somehow has to be made global. I would assume this has to be done with static/Mutex/RefCell, just like the LEDs are global in the other example:

`type LEDPIN = gpioa::PA4&lt;Output&lt;PushPull&gt;&gt;;` 

`static GLED: Mutex&lt;RefCell&lt;Option&lt;LEDPIN&gt;&gt;&gt; = Mutex::new(RefCell::new(None));`

I don't understand how to do it, though: should everything be made global, e.g. RCC and SDA/SCL pins, to set I2C and the display outside of the CS? Or just the display instance, but in that case how do I put it in the Mutex?

&amp;#x200B;

Thanks for any help on this!!!

&amp;#x200B;

*ps. contrary to some other crates, e.g. F1 or F4, the F0xx crate specifically requires the Critical Section for setting up the pins, therefore I cannot take the whole display setup out.*
