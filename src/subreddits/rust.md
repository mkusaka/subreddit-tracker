# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (17/2020)!](https://www.reddit.com/r/rust/comments/g4nu6j/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g4nu6j/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 335](https://www.reddit.com/r/rust/comments/g6d0ac/this_week_in_rust_335/)
- url: https://this-week-in-rust.org/blog/2020/04/21/this-week-in-rust-335/
---

## [3][Rust concurrency: five easy pieces.](https://www.reddit.com/r/rust/comments/g8b46i/rust_concurrency_five_easy_pieces/)
- url: https://www.reddit.com/r/rust/comments/g8b46i/rust_concurrency_five_easy_pieces/
---
How to structure concurrent workflows in Rust, via five simple examples.         

[https://medium.com/@polyglot\_factotum/rust-concurrency-five-easy-pieces-871f1c62906a](https://medium.com/@polyglot_factotum/rust-concurrency-five-easy-pieces-871f1c62906a)
## [4][Rust needs a wiki](https://www.reddit.com/r/rust/comments/g7s4ss/rust_needs_a_wiki/)
- url: https://www.reddit.com/r/rust/comments/g7s4ss/rust_needs_a_wiki/
---
Rust needs a wiki, or at least a hub for community documentation and other miscellaneous knowledge about the Rust language.

Currently when using the Rust language, it's quite easy to be efficient if you already know what you should be using, thanks to many things: cargo-doc which generates the documentation locally or online on docs.rs, the extensive book and docs for the standard library, and the community which often won't leave your questions answered on github/discord/IRC.

However, it's been really, really difficult over the last few years to start and get to that point where your Cargo.toml is stable and you start iterating quickly. Often, there will be a long process of trial and error where a user might try several crates for a given problem, or search extensively for a feature that might exist in rustc/cargo/other tool, but they just didn't know where to find it.

Some examples:

* The async story. What's the difference between tokio and async-std again? Do I need to read a blog post for that? Do I need to read a reddit post from a year ago that might already be outdated?
* The HTTP Server / Client story. Hyper ? Reqwest ? Iron ? Actix-web ? Rocket ? Tide ? What is the difference between all of those, what do they do?
* The 2D / 3D graphics story. I want to make a pong game. Piston? gfx? lyon? glium? What are the difference between all of these? By the way, I have this weird thing on windows where the command line still pops when I run my game on windows, what's up with that?
* Many other stories: Embedded, WASM, GUI, ... Even the IDE experience is a mess! RLS, rust-analyzer? When I want to know what are the recommended stuff for vim/nvim on google, I only find a post from a year ago, and even then there are multiple answers!

(Note that for most of the questions above, I did manage to answer them, but it took quite a bit of time that I wish I didn't spend. As we speak, some people might be struggling to understand those topics as well)

There have been attempts to alleviate the problem. http://www.arewewebyet.org/ https://areweguiyet.com/ https://arewegameyet.com/ those are great to know the rough state for each of those topics, but it stops there. Want to know which ones of those crates you should use together? Well, you could try to see if there is a real-life example that uses the crate you intend to use, if you're lucky it's been updated less than 2 years ago.

**The real solution is to have a community wiki**, akin to the one the Arch Linux community has. For reference, the Arch Linux wiki is the best tech-related hub of documentation linux users might even encounter, period. It is sober in design, and does its job. You need something? Look at the wiki, you'll probably find your answer there. You know something? Since it's overall maintained by the community, anyone can fix outdated stuff in a snap.

This would fix so many problems current Rust has:

* Comparative lists / tables of crates (or frameworks) to help new users find their way, without making a new website (website... that might not even be known by new rust users)
* Common issues within the rust community would be almost always up-to-date. Up-to-date "embedded" page, up-to-date "IDE Experience" page, up-to-date "GUI" page, ect.
* Community-driven: no risk of a useful blogpost being shutdown by its author. No risk of an important but outdated piece of info being on a website everyone knows the fix to, but nobody has the rights to change them.
* One unique source of knowledge, no more juggling between multiple blog posts to get what you want. (Bonus: backup friendly, less risk to lose something forever).


Of course the book, docs.rs and the official documentation should stay as they are. The wiki and the official rust website should just cross-reference each other to make sure everything is consistent. But there are so, so many topics that can be in a wiki, it really feels like there is a strong need for one. Look at me in the eyes and tell me [this awesome github repo](https://github.com/johnthagen/min-sized-rust) doesn't have the format of a wiki page.

What are your toughts on this? Any other points I might have missed? I don't know where to submit this as a serious proposal to the rust teams (I don't even know if I can!), but I do intend to make sure this idea gets discussed thoroughly and hopefully approved.
## [5][Proof Of Concept: Error Return Traces in Rust aka lightweight backtraces](https://www.reddit.com/r/rust/comments/g80c4h/proof_of_concept_error_return_traces_in_rust_aka/)
- url: https://twitter.com/yaahc_/status/1253771822920634369?s=19
---

## [6][Debugging on windows with the msvc toolchain?](https://www.reddit.com/r/rust/comments/g8c53a/debugging_on_windows_with_the_msvc_toolchain/)
- url: https://www.reddit.com/r/rust/comments/g8c53a/debugging_on_windows_with_the_msvc_toolchain/
---
AHoy!

I'm wondering what folks are using to debug on Windows. LLDB doesn't seem to work for me at all, I don't think there is an MSVC compatible GDB, I've had some early luck with remedybg (but lacks visualizers for Rust currently).

WinDBG was an issue for me because I need to have an environment setup that I didn't easily find a way to configure. So I don't know how well it works.

Visual Studio seems ok-ish... but yeah, not exactly awesome.

&amp;#x200B;

What's everyone else using?
## [7][[pre-release]: Introducing color-eyre, a custom context for for colorful error reports via the eyre crate](https://www.reddit.com/r/rust/comments/g7xf1r/prerelease_introducing_coloreyre_a_custom_context/)
- url: https://raw.githubusercontent.com/yaahc/color-eyre/master/pictures/full.png
---

## [8][Building a simple JIT in Rust](https://www.reddit.com/r/rust/comments/g83aan/building_a_simple_jit_in_rust/)
- url: https://www.jonathanturner.org/2015/12/building-a-simple-jit-in-rust.html
---

## [9][Simple snake game using SDL2](https://www.reddit.com/r/rust/comments/g84zvk/simple_snake_game_using_sdl2/)
- url: https://www.reddit.com/r/rust/comments/g84zvk/simple_snake_game_using_sdl2/
---
I wrote a simple snake game in rust using sdl2. Am a newbie rustacean and haven't been programming for long. I just program as a hobby. Feedback is very much appreciated :D 

[https://github.com/Whity25/Rust-Snake-sdl2](https://github.com/Whity25/Rust-Snake-sdl2)
## [10][A full stack application written in Rust (yew, yew-router, stdweb)](https://www.reddit.com/r/rust/comments/g7ztre/a_full_stack_application_written_in_rust_yew/)
- url: https://www.reddit.com/r/rust/comments/g7ztre/a_full_stack_application_written_in_rust_yew/
---
Hey r/rust! 

My friend and I wrote a simple full stack application in Rust, to showcase what can be done using Rust and wanted to share it with the community. 

Rusty Connect 4 is a full-stack project written completely in Rust. It uses:

* rocket on the backend
* yew for creating front-end webapps with WebAssembly (yew is a great component-based framework!)
* yew-router for routing
* stdweb to provide Rust bindings for Web APIs

The backend is only used for requests and saving your progress. You can still play the game with just the frontend.

Any feedback is welcome! Also feel free to fork or create any pull request.

repo can be found [here](https://github.com/ahmedelgohary/rusty-connect4)
## [11][Shipyard 0.4 release](https://www.reddit.com/r/rust/comments/g82tza/shipyard_04_release/)
- url: https://www.reddit.com/r/rust/comments/g82tza/shipyard_04_release/
---
[Shipyard](https://crates.io/crates/shipyard) is an Entity Component System crate. ECS is a pattern mostly used in games but not only. It fits really well with Rust, allowing easy composition and lifetime management.

**What's new**

* Rework of systems and types used to borrow storage
* Workloads had to be reworked to handle this change
* Workloads can now return errors
* `Iterator` and `IntoIterator` are now supported

For 0.3 users there's a [migration guide](https://leudz.github.io/shipyard/book/recipes/0.4-migration.html) to help with all the changes.

**Example:**

    use shipyard::*;

    struct Health(f32);
    struct Position {
        _x: f32,
        _y: f32,
    }

    fn in_acid(positions: View&lt;Position&gt;, mut healths: ViewMut&lt;Health&gt;) {
        for (_, mut health) in (&amp;positions, &amp;mut healths)
            .iter()
            .filter(|(pos, _)| is_in_acid(pos))
        {
            health.0 -= 1.0;
        }
    }

    fn is_in_acid(_: &amp;Position) -&gt; bool {
        // it's wet season
        true
    }

    fn main() {
        let world = World::new();

        world.run(
            |mut entities: EntitiesViewMut,
            mut positions: ViewMut&lt;Position&gt;,
            mut healths: ViewMut&lt;Health&gt;| {
                entities.add_entity(
                    (&amp;mut positions, &amp;mut healths),
                    (Position { _x: 0.0, _y: 0.0 }, Health(1000.0)),
                );
            },
        );

        world.run(in_acid);
    }
## [12][Help cross-compiling for Raspberry Pi Zero W](https://www.reddit.com/r/rust/comments/g86tc1/help_crosscompiling_for_raspberry_pi_zero_w/)
- url: https://www.reddit.com/r/rust/comments/g86tc1/help_crosscompiling_for_raspberry_pi_zero_w/
---
Hello! 

If this isn't the right place for this question, my apologies!

Background: I built a head-mounted computer. For lightness' sake, it's a Raspberry Pi Zero W. 

I have a project that I'm trying to cross-compile to run on the Pi, but I'm having trouble. My understanding is that the Pi Zero W has an Arm V6 processor. When I cross-compile my code and upload the executable, I get "\[1\]    2044 illegal hardware instruction" 

Compiling the project directly on the Pi *was* working, right up until I started getting out of memory errors trying to compile my most recent commit.

I'm new to Rust (and cross-compiling, come to that) and still trying to sort this all out. I'm not sure what I'm doing wrong, or what information might be relevant. So here's an info dump of everything I can think of.work

I'm using [https://github.com/rust-embedded/cross](https://github.com/rust-embedded/cross) to build a basic cross-compiling container, which I then extend to actually perform the compilation. I'm extending the base image because I also need a cross-compiled OpenSSL.

&amp;#x200B;

* Project: [https://github.com/swordsmanluke/todor/](https://github.com/swordsmanluke/todor/)
* Cross Compilation container base: rustembedded/cross:arm-unknown-linux-gnueabihf-0.2.0
* Host: Ubuntu 18.04
* RustC: 1.43
* GCC: 5.4.0
* Cross.toml:

&amp;#8203;

    [target.arm-unknown-linux-gnueabihf]
    image = "my/pi-build:latest"

* DockerFile:

&amp;#8203;

    FROM rustembedded/cross:arm-unknown-linux-gnueabihf-0.2.0
    ENV DEBIAN_FRONTEND=noninteractive
    ENV PKG_CONFIG_PATH=/usr/lib/arm-linux-gnueabihf/pkgconfig
    ENV RPI_TOOLS=/rpi_tools
    ENV MACHINE=armv6
    ENV ARCH=armv6
    ENV CC=gcc
    ENV OPENSSL_DIR=/openssl
    ENV CROSSCOMP_DIR=/rpi_tools/arm-bcm2708/arm-rpi-4.9.3-linux-gnueabihf/bin
    
    RUN dpkg --add-architecture armhf
    RUN apt-get update &amp;&amp;\
        apt-get install -y wget openssl:armhf libssl-dev:armhf pkg-config libudev-dev:armhf
    
    # Get Raspberry Pi cross-compiler tools
    RUN git -C "/" clone -q --depth=1 https://github.com/raspberrypi/tools.git "${RPI_TOOLS}"
    
    # Manually cross-compile OpenSSL to link with
    
    # 1) Download OpenSSL 1.1.0
    RUN mkdir -p $OPENSSL_DIR
    RUN cd /tmp &amp;&amp;\
        wget --no-check-certificate https://www.openssl.org/source/openssl-1.1.0h.tar.gz &amp;&amp;\
        tar xzf openssl-1.1.0h.tar.gz
    
    # 2) Compile
    RUN cd /tmp/openssl-1.1.0h &amp;&amp;\
        ./Configure linux-generic32 shared\
          --prefix=$INSTALL_DIR --openssldir=$OPENSSL_DIR/openssl\
          --cross-compile-prefix=$CROSSCOMP_DIR/arm-linux-gnueabihf- &amp;&amp;\
          make depend &amp;&amp; make &amp;&amp; make install
