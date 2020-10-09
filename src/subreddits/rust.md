# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (41/2020)!](https://www.reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j1jgum/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 359](https://www.reddit.com/r/rust/comments/j71tq7/this_week_in_rust_359/)
- url: https://this-week-in-rust.org/blog/2020/10/07/this-week-in-rust-359/
---

## [3][Memory Safe ‘curl’ for a More Secure Internet](https://www.reddit.com/r/rust/comments/j7yegb/memory_safe_curl_for_a_more_secure_internet/)
- url: https://www.abetterinternet.org/post/memory-safe-curl/
---

## [4][density-mesh v1.2 | Now it supports working with mesh chunks, ideal for big worlds and real-time terrain generation!](https://www.reddit.com/r/rust/comments/j7onr7/densitymesh_v12_now_it_supports_working_with_mesh/)
- url: https://i.redd.it/1zo3jps4nyr51.png
---

## [5][Announcing Rust 1.47.0](https://www.reddit.com/r/rust/comments/j7d49v/announcing_rust_1470/)
- url: https://blog.rust-lang.org/2020/10/08/Rust-1.47.html
---

## [6][Repos with beginner issues](https://www.reddit.com/r/rust/comments/j7w20s/repos_with_beginner_issues/)
- url: https://www.reddit.com/r/rust/comments/j7w20s/repos_with_beginner_issues/
---
Hi! I was wondering if anyone could suggest me repositories with nice beginner issues to start working on something practical after having finished the book. Thanks!
## [7][dangerous: library for safely and explicitly handling untrusted aka dangerous data](https://www.reddit.com/r/rust/comments/j7u0ue/dangerous_library_for_safely_and_explicitly/)
- url: https://www.reddit.com/r/rust/comments/j7u0ue/dangerous_library_for_safely_and_explicitly/
---
First development release for a parsing library with optional verbose errors. Would love to hear any opinions/issues! I'm working on a couple of libraries that consume the dangerous crate, so hopefully with decent amount of dog-fooding I'll find some improvements (or bugs heh) along the way.

[https://github.com/avitex/rust-dangerous](https://github.com/avitex/rust-dangerous)

Updated with example verbose error below

```text
[ERRO]: ahhh!: error attempting to convert input to str: expected utf-8 code point
&gt; ['h' 'e' ff 'l' 'o']
           ^^
additional:
  error offset: 2, input length: 5
backtrace:
  1. `read all`
  2. `read` (expected message)
  3. `read` (expected body)
  4. `convert input to str` (expected utf-8 code point)
```
## [8][Compile-time guarantee for regexp captures?](https://www.reddit.com/r/rust/comments/j7w2s0/compiletime_guarantee_for_regexp_captures/)
- url: https://www.reddit.com/r/rust/comments/j7w2s0/compiletime_guarantee_for_regexp_captures/
---
Hypothetical question: what would it take for the compiler to be able to check that for `regex::Regex::new(r"^(\d{1,18})blargh$").captures(some_string).unwrap().get(1)` is not only always `Some()`, but also a valid `u64` after parsing?
## [9][Best Book/material for People who want to level up in systems programming](https://www.reddit.com/r/rust/comments/j7p7or/best_bookmaterial_for_people_who_want_to_level_up/)
- url: https://www.reddit.com/r/rust/comments/j7p7or/best_bookmaterial_for_people_who_want_to_level_up/
---
I’m looking to “level up” my understanding of more “advanced” topics in Rust at a higher level. Like when I should use ARC, channel, etc. Coming from Elixir I can look into erlang books and materials. I was wondering if anyone had any recommendations about picking up some higher level topics in rust (which I haven’t been able to find) or even in C.
## [10][Patterns of fallible iteration](https://www.reddit.com/r/rust/comments/j7n6ei/patterns_of_fallible_iteration/)
- url: https://morestina.net/blog/1607/fallible-iteration
---

## [11][serde_json benchmark](https://www.reddit.com/r/rust/comments/j7yopg/serde_json_benchmark/)
- url: https://www.reddit.com/r/rust/comments/j7yopg/serde_json_benchmark/
---
Hello everyone.

According to this, serde\_json has a rate of above 500 MB/s parsing a json string to a structure.

                                    DOM                STRUCT
    ======= serde_json ======= parse|stringify ===== parse|stringify ====
    data/canada.json         280 MB/s   370 MB/s   510 MB/s   320 MB/s
    data/citm_catalog.json   400 MB/s   420 MB/s   840 MB/s   660 MB/s
    data/twitter.json        270 MB/s   730 MB/s   530 MB/s   850 MB/s
    

[https://github.com/serde-rs/json-benchmark](https://github.com/serde-rs/json-benchmark)

&amp;#x200B;

However, when I try to parse an 18 MB file it takes me about half a second when built for release.

Here's my test code:

        let mut buffer = String::new();
        let mut f = File::open("data.json").unwrap();
        f.read_to_string(&amp;mut buffer).unwrap();
    
        let start = now(); // now is a fn that returns current time in ms
        let value: serde_json::Value = serde_json::from_str(buffer.as_str()).unwrap();
        println!("string to value: {}", now() - start);

Am I reading the benchmark wrong? Is my code wrong?  
Who's lying?
## [12][Woohoo, milestones being made. Updated 2 days ago](https://www.reddit.com/r/rust/comments/j76xgg/woohoo_milestones_being_made_updated_2_days_ago/)
- url: https://i.redd.it/o7w56ku9xsr51.png
---

