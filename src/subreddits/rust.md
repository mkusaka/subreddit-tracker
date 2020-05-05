# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (19/2020)!](https://www.reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (19/2020)?](https://www.reddit.com/r/rust/comments/gd6gvk/whats_everyone_working_on_this_week_192020/)
- url: https://www.reddit.com/r/rust/comments/gd6gvk/whats_everyone_working_on_this_week_192020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-19-2020/42034?u=llogiq)!
## [3][An experimental asynchronous runtime based on io-uring](https://www.reddit.com/r/rust/comments/gdrrl0/an_experimental_asynchronous_runtime_based_on/)
- url: https://github.com/quininer/ritsu
---

## [4][SuckIT, a fast, multithreaded website downloader](https://www.reddit.com/r/rust/comments/gdwuat/suckit_a_fast_multithreaded_website_downloader/)
- url: https://www.reddit.com/r/rust/comments/gdwuat/suckit_a_fast_multithreaded_website_downloader/
---
## Introduction 
[SuckIT](https://github.com/skallwar/suckit) is a multithreaded, open source web downloader written in Rust. It aims to recursively download webpages and allow offline browsing.

## Benchmark
As of right now, it's a _little_ faster (about 3460%) than HTTrack on a single core, 60 second run on http://book.toscrape.com/

name    | pages downloaded
------- | ----------------
suckit  |             2422
httrack |               70

http://book.toscrape.com/ is downloaded in 75s for 1 thread, 37s for 2 threads and 19s for 4 threads. It's more or less a linear time reduction

## Future

Some features are missing, such as a random delay between downloads (to avoid IP ban) or user agent support but they will come soon.

Any feedback, reviews or PRs are welcome !

Enjoy :)
## [5][Microsoft wants Rust instead of Go in the cloud](https://www.reddit.com/r/rust/comments/gdwxw1/microsoft_wants_rust_instead_of_go_in_the_cloud/)
- url: https://www.reddit.com/r/rust/comments/gdwxw1/microsoft_wants_rust_instead_of_go_in_the_cloud/
---
 [https://www.en24.news/en/2020/05/secure-programming-language-microsoft-wants-rust-instead-of-go-in-the-cloudhtml](https://www.en24.news/en/2020/05/secure-programming-language-microsoft-wants-rust-instead-of-go-in-the-cloudhtml)
## [6][Restart accel project, GPGPU Framework for Rust: 0.3.0 Release](https://www.reddit.com/r/rust/comments/gdmhoq/restart_accel_project_gpgpu_framework_for_rust/)
- url: https://users.rust-lang.org/t/restart-accel-project-gpgpu-framework-for-rust-0-3-0-release/42087
---

## [7][rust-analyzer Changelog #23](https://www.reddit.com/r/rust/comments/gdbrht/rustanalyzer_changelog_23/)
- url: https://rust-analyzer.github.io/thisweek/2020/05/04/changelog-23.html
---

## [8][Reinventing Asynchronous Rust](https://www.reddit.com/r/rust/comments/gdocgk/reinventing_asynchronous_rust/)
- url: https://www.reddit.com/r/rust/comments/gdocgk/reinventing_asynchronous_rust/
---
[https://aldaronlau.com/reinventing-async-rust/](https://aldaronlau.com/reinventing-async-rust/)
## [9][cargo crev and cargo audit 🦀 the second online meetup ☁](https://www.reddit.com/r/rust/comments/gdmlp6/cargo_crev_and_cargo_audit_the_second_online/)
- url: https://estada.ch/2020/4/30/cargo-crev-and-cargo-audit-the-second-online-meetup/
---

## [10][Point of WebGPU on native](https://www.reddit.com/r/rust/comments/gdbgoc/point_of_webgpu_on_native/)
- url: http://kvark.github.io/web/gpu/native/2020/05/03/point-of-webgpu-native.html
---

## [11][An ESP8266/ESP32 Over The Air (OTA) firmware server, written in Rust!](https://www.reddit.com/r/rust/comments/gdd030/an_esp8266esp32_over_the_air_ota_firmware_server/)
- url: https://github.com/Evander12345/rota
---

## [12][I made a hashmap that is keyed by types and maps vectors of the key-types](https://www.reddit.com/r/rust/comments/gdwmp1/i_made_a_hashmap_that_is_keyed_by_types_and_maps/)
- url: https://www.reddit.com/r/rust/comments/gdwmp1/i_made_a_hashmap_that_is_keyed_by_types_and_maps/
---
I was surprised I didn't need to use any unsafe code.

    use core::any::TypeId;
    use std::any::Any;
    use std::collections::hash_map::HashMap;

    struct VecStorage {
        vecs: HashMap&lt;TypeId, Box&lt;dyn Any&gt;&gt;
    }

    impl VecStorage {
        pub fn new() -&gt; Self {
            VecStorage {
                vecs: HashMap::new(),
            }
        }

        pub fn push&lt;T: 'static&gt;(&amp;mut self, val: T) {
            let key = (&amp;val as &amp;dyn Any).type_id();
            let vec: Option&lt;&amp;mut Vec&lt;T&gt;&gt; = self.vecs.get_mut(&amp;key).map(|x| {
                x.as_mut()
                    .downcast_mut::&lt;Vec&lt;T&gt;&gt;()
                    .expect("failed to downcast")
            });
            match vec {
                Some(v) =&gt; v.push(val),
                None =&gt; {
                    let mut v = Vec::new();
                    v.push(val);
                    self.vecs.insert(key, Box::new(v));
                }
            }
        }

        pub fn get&lt;T: 'static&gt;(&amp;mut self, index: usize) -&gt; Option&lt;&amp;mut T&gt; {
            let key = TypeId::of::&lt;T&gt;();
            let vec: &amp;mut Vec&lt;T&gt; = self.vecs.get_mut(&amp;key)?.as_mut().downcast_mut::&lt;Vec&lt;T&gt;&gt;()?;
            vec.get_mut(index)
        }
    }

    fn main() {
        let mut storage = VecStorage::new();
        storage.push(100);
        storage.push("Hello world".to_owned());
        storage.push(200);
        println!("{:?}", storage.get::&lt;i32&gt;(0));
        println!("{:?}", storage.get::&lt;String&gt;(0));
        println!("{:?}", storage.get::&lt;i32&gt;(1));
        println!("{:?}", storage.get::&lt;i32&gt;(2));
        // Some(100)
        // Some("Hello world")
        // Some(200)
        // None
    }
