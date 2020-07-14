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
## [2][What's everyone working on this week (29/2020)?](https://www.reddit.com/r/rust/comments/hq8j8i/whats_everyone_working_on_this_week_292020/)
- url: https://www.reddit.com/r/rust/comments/hq8j8i/whats_everyone_working_on_this_week_292020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-29-2020/45746?u=llogiq)!
## [3][Making a Game in 48 hours with Rust and WebAssembly](https://www.reddit.com/r/rust/comments/hqu6dj/making_a_game_in_48_hours_with_rust_and/)
- url: https://ianjk.com/rust-gamejam/
---

## [4][Raspberry Pi on steroids with Rust and WebAssembly](https://www.reddit.com/r/rust/comments/hqzjk8/raspberry_pi_on_steroids_with_rust_and_webassembly/)
- url: https://www.secondstate.io/articles/get-started-with-raspberry-pi-20200708/
---

## [5][Descriptive Statistics with NDArray](https://www.reddit.com/r/rust/comments/hqz2w7/descriptive_statistics_with_ndarray/)
- url: https://shahinrostami.com/posts/programming/rust-notebooks/descriptive-statistics-with-ndarray/
---

## [6][[Gtk-rs] gladis: easily import Glade-generated UI files into Rust code](https://www.reddit.com/r/rust/comments/hr0wmc/gtkrs_gladis_easily_import_gladegenerated_ui/)
- url: https://www.reddit.com/r/rust/comments/hr0wmc/gtkrs_gladis_easily_import_gladegenerated_ui/
---
Hello gals and guys,

As I was tired to call the builder methods when using [Gtk-rs](https://gtk-rs.org/) to [give life to a .glade UI file](https://gtk-rs.org/docs-src/tutorial/glade), I took the time to write a proc macro that will automatically import widgets from a declarative struct of them and their type.

[https://crates.io/crates/gladis](https://crates.io/crates/gladis)

This allows you to turn this "idiomatic" Gtk-rs code:

    fn build_ui(app: &amp;gtk::Application) {
        let builder = gtk::Builder::from_string(include_str!("main.glade"));
    
        let window: gtk::Window = builder
            .get_object("window")
            .expect("could not find window");
        let my_label: gtk::Label = builder
            .get_object("my_label")
            .expect("could not find my_label");
    
        my_label.set_label("hello from Rust!");
        window.set_application(Some(app));
        window.show_all();
    }

Into this "declarative" approach using a struct:

    use gladis::Gladis;
    use gladis_proc_macro::Gladis;
    
    #[derive(Gladis, Clone)]
    struct Ui {
        window: gtk::Window,
        my_label: gtk::Label,
    }
    
    fn build_ui(app: &amp;gtk::Application) {
        let ui = Ui::from_glade_src(include_str!("main.glade"));
    
        ui.label.set_label("hello from Rust!");
        ui.window.set_application(Some(app));
        ui.window.show_all();
    }

This gets handy especially when you start to have a good number of widgets. By deriving from `Clone`, the `ui` object can be cloned with all the widgets when necessary, for example when creating closures to handle events like on a button press.

The crate is named [gladis](https://crates.io/crates/gladis) (and [gladis\_proc\_macro](https://crates.io/crates/gladis_proc_macro)) as a [glade](https://crates.io/crates/glade) (and [glade\_derive](https://crates.io/crates/glade_derive)) crate already existed but they seemed to be limited. This is my very first proc-macro crate, I now realize that maybe I should have used *gladis\_derive* crate name but (one of?) the book encouraged me to use *gladis\_proc\_macro* instead.

Feedback welcome! I hope this will help users getting started with Gtk-rs faster.
## [7][PGX: Write postgres extensions in Rust](https://www.reddit.com/r/rust/comments/hqjevv/pgx_write_postgres_extensions_in_rust/)
- url: https://github.com/zombodb/pgx
---

## [8][Autovivification(basic) in Rust](https://www.reddit.com/r/rust/comments/hqxbtf/autovivificationbasic_in_rust/)
- url: https://www.reddit.com/r/rust/comments/hqxbtf/autovivificationbasic_in_rust/
---
[https://repl.it/talk/share/Autovivification-in-Rust/45262](https://repl.it/talk/share/Autovivification-in-Rust/45262)
## [9][rust-analyzer changelog #33](https://www.reddit.com/r/rust/comments/hqdgli/rustanalyzer_changelog_33/)
- url: https://rust-analyzer.github.io/thisweek/2020/07/13/changelog-33.html
---

## [10][IntelliJ Rust Changelog #126](https://www.reddit.com/r/rust/comments/hqijgt/intellij_rust_changelog_126/)
- url: https://intellij-rust.github.io/2020/07/13/changelog-126.html
---

## [11][The confidence you get when your code finally compiles is dangerous. Published my crate before I knew it even worked!](https://www.reddit.com/r/rust/comments/hqud3r/the_confidence_you_get_when_your_code_finally/)
- url: https://www.reddit.com/r/rust/comments/hqud3r/the_confidence_you_get_when_your_code_finally/
---
Did end up having a major bug.

[Announcing release of v0.2.0](https://github.com/intendednull/yew-state/releases)!
## [12][rust-sum-bot (my first Rust project)](https://www.reddit.com/r/rust/comments/hqzdzz/rustsumbot_my_first_rust_project/)
- url: https://imgur.com/a/jXAFgbC
---

