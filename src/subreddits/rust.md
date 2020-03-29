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
## [2][This Week in Rust 331](https://www.reddit.com/r/rust/comments/fp9z5t/this_week_in_rust_331/)
- url: https://this-week-in-rust.org/blog/2020/03/24/this-week-in-rust-331/
---

## [3][R2: A Router in Rust](https://www.reddit.com/r/rust/comments/fr50vl/r2_a_router_in_rust/)
- url: https://www.reddit.com/r/rust/comments/fr50vl/r2_a_router_in_rust/
---
Introducing R2 [**https://r2.rs/,**](https://r2.rs/,) a 'Router in Rust' that I wrote. Read my experience with Rust at [**https://r2.rs/blog/**](https://r2.rs/blog/)
## [4][Traits working group 2020 sprint 1 summary](https://www.reddit.com/r/rust/comments/fqrrv5/traits_working_group_2020_sprint_1_summary/)
- url: https://blog.rust-lang.org/inside-rust/2020/03/28/traits-sprint-1.html
---

## [5][Clokwerk run scheduler indefinitely](https://www.reddit.com/r/rust/comments/fr414j/clokwerk_run_scheduler_indefinitely/)
- url: https://www.reddit.com/r/rust/comments/fr414j/clokwerk_run_scheduler_indefinitely/
---
Hello!
Recently, I started learning Rust. I'm trying to run Clokwerk to go on forever, but the main function is dropping the reference to ScheduleHandle at the end. What's the proper way to keep it running indefinitely?

    fn main() {
        let yaml = load_yaml!("app.yaml");
        let matches = App::from_yaml(yaml).get_matches();
    
        let mut params = Params::new();
    
        let domain = matches.value_of("domain").unwrap();
        params.domain = String::from(domain);
    
        let subdomains = matches.value_of("subdomains").unwrap();
        params.subdomains = subdomains.split(",").map(|s| s.to_string()).collect();
    
        scheduler
            .every(10.seconds())
            .run(move || run_checker(&amp;params));
    
        let _thread_handle: ScheduleHandle = scheduler.watch_thread(Duration::from_millis(6000));
    }


I'm at the beginning of my adventure with Rust, so please be understanding :).
## [6][Implementing the lambda calculus in Rust](https://www.reddit.com/r/rust/comments/fqpwan/implementing_the_lambda_calculus_in_rust/)
- url: https://www.reddit.com/r/rust/comments/fqpwan/implementing_the_lambda_calculus_in_rust/
---
Here is my first post in what will hopefully be a series about type systems: https://christianpoveda.github.io/blog/untyped-lambda-calculus/

I'm happy to receive constructive feedback about it.
## [7][Unsure what graphics library to use in my Rust application](https://www.reddit.com/r/rust/comments/fqsw4t/unsure_what_graphics_library_to_use_in_my_rust/)
- url: https://www.reddit.com/r/rust/comments/fqsw4t/unsure_what_graphics_library_to_use_in_my_rust/
---
Hello fellow rustaceans!

I am coding the Conway's Game of Life as part of my journey to learn Rust.

I have already coded the "engine" and a nice terminal output with termion, but I would also like to have a windowed graphical display for it. I am unsure what dependency to use for this purpose. I did my research and gfx-rs seems the most popular option, but seems just too low level to just draw some squares in a window. Amethyst in the other end, looks too overfecthed, using a game engine for this seems unreasonable.

I am new to Rust and to graphics, so I am not sure if it is worth it to spend some ours learning gfx-rs or go directly to some higher level library, if so, which one? Also, if possible, I would like this library to be cross-platform and pure Rust.

Any good recommendations? Thank you!
## [8][derive_more 0.99.5 released: Support for Error derive similar to dtolnay/thiserror](https://www.reddit.com/r/rust/comments/fqkxtm/derive_more_0995_released_support_for_error/)
- url: https://www.reddit.com/r/rust/comments/fqkxtm/derive_more_0995_released_support_for_error/
---
derive_more allows you to derive lots of traits that are present in the standard library, but that you normally have to implement manually yourself. 
This release adds support for deriving `Error` and it's methods `source` and `backtrace`. This works very well with the `Display` derive that already exists in this library.

* Github: [https://github.com/JelteF/derive_more](https://github.com/JelteF/derive_more)
* Docs: [https://jeltef.github.io/derive_more/derive_more/index.html](https://jeltef.github.io/derive_more/derive_more/index.html)
* Docs for Error derive: [https://jeltef.github.io/derive_more/derive_more/error.html](https://jeltef.github.io/derive_more/derive_more/error.html)
*  Crates.io: [https://crates.io/crates/derive_more](https://crates.io/crates/derive_more)
## [9][I need help sending emails with rust](https://www.reddit.com/r/rust/comments/fr53jb/i_need_help_sending_emails_with_rust/)
- url: https://www.reddit.com/r/rust/comments/fr53jb/i_need_help_sending_emails_with_rust/
---
I want to make a login system with email verification. I found the crate lettre but I keep getting errors when sending.

The code I tried:

`extern crate lettre;`  
`use lettre::smtp::authentication::{Credentials, Mechanism};`  
`use lettre::{SendableEmail, Envelope, EmailAddress, Transport, SmtpClient};`  
`use lettre::smtp::extension::ClientId;`  
`use lettre::smtp::ConnectionReuseParameters;`  
`fn main() {`  
 `let email_1 = SendableEmail::new(`  
`Envelope::new(`  
 `Some(EmailAddress::new("user@localhost".to_string()).unwrap()),`  
 `vec![EmailAddress::new("root@localhost".to_string()).unwrap()],`  
 `).unwrap(),`  
 `"id1".to_string(),`  
 `"Hello world".to_string().into_bytes(),`  
 `);`  
`let email_2 = SendableEmail::new(`  
`Envelope::new(`  
 `Some(EmailAddress::new("user@localhost".to_string()).unwrap()),`  
 `vec![EmailAddress::new("root@localhost".to_string()).unwrap()],`  
 `).unwrap(),`  
 `"id2".to_string(),`  
 `"Hello world a second time".to_string().into_bytes(),`  
 `);`  
 `// Connect to a remote server on a custom port`  
 `let mut mailer = SmtpClient::new_simple("smtp.mailtrap.io").unwrap()`  
 `// Set the name sent during EHLO/HELO, default is \`localhost\``  
 `.hello_name(ClientId::Domain("my.hostname.tld".to_string()))`  
 `// Add credentials for authentication`  
 `.credentials(Credentials::new("username".to_string(), "password".to_string()))`  
 `// Enable SMTPUTF8 if the server supports it`  
 `.smtp_utf8(true)`  
 `// Configure expected authentication mechanism`  
 `.authentication_mechanism(Mechanism::Plain)`  
 `// Enable connection reuse`  
 `.connection_reuse(ConnectionReuseParameters::ReuseUnlimited).transport();`  
`let result_1 = mailer.send(email_1);`  
 `assert!(result_1.is_ok());`  
 `// The second email will use the same connection`  
 `let result_2 = mailer.send(email_2);`  
 `assert!(result_2.is_ok());`  
 `// Explicitly close the SMTP transaction as we enabled connection reuse`  
 `mailer.close();`  
`}`

&amp;#x200B;

It panics on the result asserts. The error in the debugger is:

[CLion debugger](https://preview.redd.it/58ie9nq8tlp41.png?width=740&amp;format=png&amp;auto=webp&amp;s=765c3100b210f9c90535bf849ad77709dc3c11ed)

Does anyone know what I am doing wrong?
## [10][How to handle money in diesel and rust?](https://www.reddit.com/r/rust/comments/fr0wvc/how_to_handle_money_in_diesel_and_rust/)
- url: https://www.reddit.com/r/rust/comments/fr0wvc/how_to_handle_money_in_diesel_and_rust/
---
So I'm making a graphql api and I'm going to need to use money as one of my fields.

I know postgres has a special money field. However, I'm not really sure how to get a value from whatever is thrown to my api and then convert it to Diesel's [Cents](https://docs.diesel.rs/diesel/pg/data_types/struct.Cents.html) class.

I would like to know if anybody has any advice. I would appreciate it.
## [11][[Help] Cursive T-UI Application structure with mutable states](https://www.reddit.com/r/rust/comments/fqybys/help_cursive_tui_application_structure_with/)
- url: https://www.reddit.com/r/rust/comments/fqybys/help_cursive_tui_application_structure_with/
---
# \# Intro
Hi, I have been using `Cursive` and some other crates together to make an auto-clicker.
A bit of a learning + useful projects started in this recent period of time.

UI: [cursive github](https://github.com/gyscos/cursive) + [cursive-tableview github](https://github.com/BonsaiDen/cursive_table_view)
KeySimulation: [rdev github](https://github.com/Narsil/rdev)
Hotkey: [hotkey github](https://github.com/unwrittenfun/hotkey-rs)

# \# Program General Structure
Code general structure, note that `new()` returns `Rc&lt;RefCell&lt;Controller&gt;&gt;` not `Controller`.
```
struct Controller {
siv: Cursive,
mode: Mode,
active_cell: Option&lt;ACAction&gt;,
}
impl Controller {
fn new_tableview() { ... } // get new tableview instance
fn new_siv() { ... } // get new cursive instance
fn new() -&gt; Rc&lt;RefCell&lt;Controller&gt;&gt; { ... } // get a shared controller
fn add_action(&amp;self) { ... } // add a new action to view, but check if mode is not RunningAction
fn run_selected_cell(&amp;self) { ... } // run current selected cell in tableview, but check if cell is None
}

```
# \# Issue
At first, I was having trouble to capture `Controller` in some of the callbacks, such as from key press n for add_action. This was resolved with #415 (comment), by using `Rc&lt;RefCell&lt;Controller&gt;&gt;`.

[gtk closure tutorial link i used for Rc&lt;RefCell&lt;_&gt;&gt;](https://gtk-rs.org/docs-src/tutorial/closures)

However, this restrict me to always `borrow_mut()` to get access to siv:

```
fn main() {
let mut controller = Controller::new();
controller.borrow_mut().siv.run();  // Starts the event loop.
}
```

Which means the callbacks that requires access to my controller by `borrow_mut()`, it will panic because the controller has been `borrow_mut()`ed forever in the main code.
I have thought of separate `siv` from `Controller`, but that makes it hard to update ui with just controller.
Using `Rc&lt;RefCell&lt;&amp;'a Controller&gt;&gt;` instead `Rc&lt;RefCell&lt;Controller&gt;&gt;` introduces a lifetime issue.

This seems like in a deadlock, `Rc&lt;RefCell&lt;_&gt;&gt;` is required in callback to ensure the program is in right 'state', but this prevents state updating because of `borrow_mut()` twice panic.

# \# End

Going through all the trouble to getting my application to work, I feel like my current way of modelling UI application is definitely not compatible with rust.

It just seems to be so much more easier to get started with programming languages such as C/C++/ObjC.

I would like to get 'trained' to think in the way of rust, similar to how I learned lisp with SICP, by reading some good books, or materials, that walks me through the ideology of rust.

I understand this would be the steep learning curve of rust that I am facing, but I also feel like this learning curve is not an introduction to something new, rather something confusing.

Thanks for reading. Any helpful advices given are appreciated in any ways.
## [12][Does Rustfmt support column alignment?](https://www.reddit.com/r/rust/comments/fqybka/does_rustfmt_support_column_alignment/)
- url: https://www.reddit.com/r/rust/comments/fqybka/does_rustfmt_support_column_alignment/
---
like 

```
struct T {
  id        : i32,
  active    : bool,
  username  : String
}
```

Also what happened to the weekly small questions thread?

Looks like reddit doesn't display the codeblocks properly in mobile clients. I really do not appreciate this. Ever since they introduced the new redesign, everything has been unpredictable. A complete UX disaster. For those who are confused, the `i32`, `bool` and `String` are intended to be aligned in one column.
