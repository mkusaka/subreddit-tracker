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
## [3][rust-analyzer changelog #33](https://www.reddit.com/r/rust/comments/hqdgli/rustanalyzer_changelog_33/)
- url: https://rust-analyzer.github.io/thisweek/2020/07/13/changelog-33.html
---

## [4][nnnoiseless: porting audio code from C to rust](https://www.reddit.com/r/rust/comments/hq2g8e/nnnoiseless_porting_audio_code_from_c_to_rust/)
- url: https://jneem.github.io/nnnoiseless/
---

## [5][Announcing yew-state: providing Yew components with ergonomic access to globally shared (and optionally persistent) state](https://www.reddit.com/r/rust/comments/hqbzgr/announcing_yewstate_providing_yew_components_with/)
- url: https://github.com/intendednull/yew-state
---

## [6][My first crate, StructConf! A structopt-like derive macro to combine argument parsing and an ini config file. Would love some feedback ;)](https://www.reddit.com/r/rust/comments/hqcetf/my_first_crate_structconf_a_structoptlike_derive/)
- url: https://github.com/vidify/structconf
---

## [7][Check if an email address exists without sending any email](https://www.reddit.com/r/rust/comments/hqbuqu/check_if_an_email_address_exists_without_sending/)
- url: https://www.reddit.com/r/rust/comments/hqbuqu/check_if_an_email_address_exists_without_sending/
---
[github repo](https://github.com/amaurymartiny/check-if-email-exists) | [docs](https://docs.rs/check-if-email-exists/0.8.11/check_if_email_exists/) | [live demo](https://reacher.email)

`check-if-email-exists` is a tool to check if an email exists without sending any email. The tool connects to the target email's SMTP server, and parses the response of the `RCPT TO:` command.

This lib is only half of the equation for email verification. Companies use a bunch of techniques (greylisting, IP blacklisting, ISPs block port 25...) to prevent spam. During this lockdown, I created a SaaS to get around these restrictions: https://reacher.email (also 100% open-source, backend in Rust). It scales with AWS Lambda, and to avoid IP blacklisting, I use Tor. I get decent results with it: ~85% of the emails I verify return a result.

Looking for feedback!
## [8][1.0 release of the Rust client for Apache Pulsar](https://www.reddit.com/r/rust/comments/hqbmz3/10_release_of_the_rust_client_for_apache_pulsar/)
- url: https://www.reddit.com/r/rust/comments/hqbmz3/10_release_of_the_rust_client_for_apache_pulsar/
---
I'm please to announce the stable release of the [pulsar crate](https://crates.io/crates/pulsar), a Rust client for [Apache Pulsar](https://pulsar.apache.org/), a distributed pub-sub system that is gaining steam this year.  


A few features:

&amp;#x200B;

* async/await interface
* TLS connections
* message batching
* compression (LZ4, zlib, zstd or Snappy)
* automatic reconnection with configurable exponential back off
* configurable executor: tokio, async-std, or any custom executor that wraps one of those
## [9][Rust is Surprisingly Good as a Server Language](https://www.reddit.com/r/rust/comments/hpzmeu/rust_is_surprisingly_good_as_a_server_language/)
- url: https://stu2b50.dev/posts/rust-is-surpris76171
---

## [10][Are there any latest best practices for no_std crates](https://www.reddit.com/r/rust/comments/hqdku7/are_there_any_latest_best_practices_for_no_std/)
- url: https://www.reddit.com/r/rust/comments/hqdku7/are_there_any_latest_best_practices_for_no_std/
---
Hey rustaceans,

I'm getting started with Rust and doing some hobby project for home automation. I've decided to split my project in 2 crates (so far) - a tiny core (mostly for data representation and some pieces of transport layer) that is supposed to be re-usable potentially on embedded devices (thus no\_std) and a regular binary crate.

I've read about no\_std crates in the book and I've checked some existing crates, however it's not entirely clear what's the best practice of writing a crate that is supposed to work with both std/no\_std. 

It definitely might be an overkill for me to play with no\_std at my current level of Rust knowledge, however I have a strong feeling that some experience with no\_std will help me understanding Rust way better.
## [11][Tour of Rust](https://www.reddit.com/r/rust/comments/hpt01d/tour_of_rust/)
- url: https://tourofrust.com/
---

## [12][The Soul of a New Debugger](https://www.reddit.com/r/rust/comments/hpy44d/the_soul_of_a_new_debugger/)
- url: https://nbaksalyar.github.io/2020/07/12/soul-of-a-new-debugger.html
---

