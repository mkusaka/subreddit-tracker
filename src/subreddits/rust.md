# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (16/2020)!](https://www.reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fw2hd8/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.42]](https://www.reddit.com/r/rust/comments/fjsj1l/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/fjsj1l/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://en.reddit.com/r/rust/comments/eyw94s/official_rrust_whos_hiring_thread_for_jobseekers/).

This thread will be periodically stickied to the top of /r/rust for improved visibility. The thread will be refreshed and posted anew when the next version of Rust releases in six weeks.

Please adhere to the following rules when posting (adapted from /r/cpp's jobs thread).

#### Rules for individuals:

* Don't create top-level comments; those are for employers.
* Feel free to reply to top-level comments with on-topic questions.
* I will create a stickied top-level comment for individuals looking for work.
* I will create an additional top-level comment for meta discussion.

#### Rules for employers:

* To find individuals seeking work, see the replies to the stickied top-level comment; you will need to click the "more comments" link at the bottom of the top-level comment in order to make these replies visible.
* To make a top-level comment you must be hiring directly; no third-party recruiters.
* One top-level comment per employer. If you have multiple job openings, please consolidate their descriptions or mention them in replies to your own top-level comment.
* Proofread your comment after posting it and edit it if necessary to correct mistakes.
* Please base your comment on the following template:

COMPANY: *[Company name; please link to your company's website or careers page.]*

TYPE: *[Full time, part time, internship, contract, etc.]*

DESCRIPTION: *[What does your company do, and what are you using Rust for? How much experience are you seeking and what seniority levels are you hiring for? The more details the better.]*

LOCATION: *[Where are your office or offices located? If your workplace language isn't English, please specify it.]*

REMOTE: *[Do you offer the option of working remotely? If so, do you require employees to live in certain areas or time zones?]*

VISA: *[Does your company sponsor visas?]*

CONTACT: *[How can someone get in touch with you?]*
## [3][Yew developer survey](https://www.reddit.com/r/rust/comments/g43ld4/yew_developer_survey/)
- url: https://www.reddit.com/r/rust/comments/g43ld4/yew_developer_survey/
---
Hello, the Yew framework is looking for feedback from devs who are familiar with the framework and have a few minutes to share their thoughts anonymously. Link here: [https://yewstack.typeform.com/to/OR5rQG](https://yewstack.typeform.com/to/OR5rQG)

*(If you're not familiar,* [Yew](https://github.com/yewstack/yew) *is a framework for building client web apps with Rust &amp; WebAssembly!)*

The main goals of this survey are:

1. Prioritize upcoming features
2. Get feedback on progress and management of the project
3. Learn more about why devs decide to use Yew for a project

Thank you!!
## [4][Wired Logic - a pixel-based digital circuit simulator running in a browser (Rust compiled into WASM).](https://www.reddit.com/r/rust/comments/g478yt/wired_logic_a_pixelbased_digital_circuit/)
- url: https://iostapyshyn.github.io/wired-logic/
---

## [5][tide 0.7.0 introduces a new http-types](https://www.reddit.com/r/rust/comments/g43wxh/tide_070_introduces_a_new_httptypes/)
- url: https://www.reddit.com/r/rust/comments/g43wxh/tide_070_introduces_a_new_httptypes/
---
I updated from [tide](https://crates.io/crates/tide) 0.6.0 to 0.7.0 and was sad to find that they have switched out from the somewhat de-facto standard [http](https://crates.io/crates/http) crate to roll their own [http-types](https://crates.io/crates/http-types).

As a consumer of these crates, I think the split in the eco-system is sad. There's also some frustration in having to deal with yet another type called `Request` and `Response`. I'm sure the creators have their reasons, and I was hunting around their git repo for a motivation, but can't find any.

What is the reason? Does others also find the http crate lacking?
## [6][OsStr Bytes is now completely safe!](https://www.reddit.com/r/rust/comments/g3v5jc/osstr_bytes_is_now_completely_safe/)
- url: https://www.reddit.com/r/rust/comments/g3v5jc/osstr_bytes_is_now_completely_safe/
---
[Documentation](https://docs.rs/os_str_bytes) | [Repository](https://github.com/dylni/os_str_bytes) | [crates.io](https://crates.io/crates/os_str_bytes)

Except for testing, the newest version of OsStr Bytes doesn't use any unsafe code and has no dependencies!

Frequently, [`mem::transmute`](https://doc.rust-lang.org/std/mem/fn.transmute.html) is used to convert between [`OsStr`](https://doc.rust-lang.org/std/ffi/struct.OsStr.html) and [`[u8]`](https://doc.rust-lang.org/std/slice/), but that's [incredibly unsafe](https://doc.rust-lang.org/std/mem/fn.transmute.html). It assumes [`OsStr`](https://doc.rust-lang.org/std/ffi/struct.OsStr.html) has only one field and is annotated with [`repr(C)`](https://doc.rust-lang.org/nomicon/other-reprs.html), which [it isn't](https://github.com/rust-lang/rust/pull/49456). Using it this way is undefined behavior.

This crate doesn't make any of those assumptions but can still convert losslessly. It's meant to help make [`OsStr`](https://doc.rust-lang.org/std/ffi/struct.OsStr.html) and similar structs easier to use safely.
## [7][The cost of `Borrow`](https://www.reddit.com/r/rust/comments/g42tnz/the_cost_of_borrow/)
- url: https://www.reddit.com/r/rust/comments/g42tnz/the_cost_of_borrow/
---
The `Borrow` trait abstracts over different "borrowable" datatypes. I'm interested in it as a way of speeding up cloning of my game's state during AI training---the clone is needed by the learning framework, and accounts for something like 75% of the runtime of training. I'd like to use a copy-on-write approach to speed this up, since in most iterations of learning most map tiles are unchanged. So I've been looking at wrapping values inside `Cow`.

But I don't need this cloning to be so fast during normal gameplay where it is used rarely, and it seems using `Cow` in all cases will put a drag on performance. So I got interested in `Borrow` and `BorrowMut` as a way of abstracting over owned map tiles vs. `Cow` map tiles. Use the owned version in normal gameplay, and use the `Cow` version during AI training for faster cloning.

But I wondered whether `Borrow` itself has a cost and, sadly, it seems it does, even when the compiler could know statically it's an owned value and not a `Cow`.

I used `cargo-asm` to see how much code each of the following functions generated:

```
pub fn do_thing_normal&lt;T:fmt::Debug&gt;(t: &amp;T) {
    println!("{:?}", t);
}

pub fn do_thing_borrow&lt;T:fmt::Debug+Borrow&lt;T&gt;&gt;(t: T) {
    println!("{:?}", t.borrow());
}

pub fn actually_do_thing_normal() {
    let x = String::from("hi");
    do_thing_normal(&amp;x);
}

pub fn actually_do_thing_borrow_cow() {
    let x: Cow&lt;String&gt; = Cow::Owned(String::from("hi"));
    do_thing_borrow(x);
}

pub fn actually_do_thing_borrow_owned() {
    let x = String::from("hi");
    do_thing_borrow(x);
}
```

The number of lines of asm code generated:

* actually\_do\_thing\_normal: 36
* actually\_do\_thing\_borrow\_owned: 42
* actually\_do\_thing\_borrow\_cow: 43

So even used on owned values, `Borrow` incurs a real cost. This was disappointing to see. Is there some reason that the compiler couldn't (in theory) optimize this so that owned values are borrowed using the exact same borrowing mechanism used when borrowing statically? Or is there some other approach I might use?
## [8][Logos 0.11 is out: Iterators, callbacks, and stateful tokens](https://www.reddit.com/r/rust/comments/g3pfaw/logos_011_is_out_iterators_callbacks_and_stateful/)
- url: https://github.com/maciejhirsz/logos/releases/tag/v0.11.0
---

## [9][I feel stupid, but where does the Rand trait come from?](https://www.reddit.com/r/rust/comments/g4789z/i_feel_stupid_but_where_does_the_rand_trait_come/)
- url: https://www.reddit.com/r/rust/comments/g4789z/i_feel_stupid_but_where_does_the_rand_trait_come/
---
I want to implement the Rand trait for a custom type. But I can't figure out how to get the trait in scope. `std::rand` does not exist (rust stable), and the `rand` crate only provides the `Rng` trait.

I found [https://doc.rust-lang.org/1.0.0/rand/index.html](https://doc.rust-lang.org/1.0.0/rand/index.html), which only states 

&gt;It is not recommended to use this library directly, but rather the official interface through `std::rand`.

But rustc says

&gt;could not find \`rand\` in \`std\`

Where can I import Rand from???
## [10][Book recommendation needed please!](https://www.reddit.com/r/rust/comments/g45ax0/book_recommendation_needed_please/)
- url: https://www.reddit.com/r/rust/comments/g45ax0/book_recommendation_needed_please/
---
The textbook should be such that it contains a problem set after every chapter, about that chapter. This is the university style learning that I prefer.  


 It should also be self contained in explaining concepts unlike for people who know C++. 
## [11][How to make a Text Editor in Rust!](https://www.reddit.com/r/rust/comments/g45bim/how_to_make_a_text_editor_in_rust/)
- url: https://www.reddit.com/r/rust/comments/g45bim/how_to_make_a_text_editor_in_rust/
---
I want to make text editor in rust, but first I want to learn rust. I will learn rust then how shall I begin to proceed with making of text editor. I am web developer, iOS developer but I haven’t developed apps for desktop. How shall I do it? How to proceed? I want it to be like vim as well as gvim. What shall I do?
## [12][Install Latest Rust Diesel in Windows 10 and Fix Issue](https://www.reddit.com/r/rust/comments/g44xae/install_latest_rust_diesel_in_windows_10_and_fix/)
- url: https://www.reddit.com/r/rust/comments/g44xae/install_latest_rust_diesel_in_windows_10_and_fix/
---
Today I have faced plenty of issues in installing Diesel CLI in windows 10. Some alternative like changing my Windows language back to English, install Visual Studio C++ tools and re-installing rust is not worked.

Finally, I found a way to solve it (i use Postgres database) and documented it here: [https://www.yodiw.com/install-rust-diesel-in-windows-10-and-fix-issue/](https://www.yodiw.com/install-rust-diesel-in-windows-10-and-fix-issue/) 

I hope it's helping someone who faces the same issue.
