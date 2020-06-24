# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (26/2020)!](https://www.reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/h98zfz/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 344](https://www.reddit.com/r/rust/comments/hepkfq/this_week_in_rust_344/)
- url: https://this-week-in-rust.org/blog/2020/06/23/this-week-in-rust-344/
---

## [3][NLnet funding, and Lemmy 0.7.0 with new image hosting!](https://www.reddit.com/r/rust/comments/hezuy1/nlnet_funding_and_lemmy_070_with_new_image_hosting/)
- url: https://www.reddit.com/r/rust/comments/hezuy1/nlnet_funding_and_lemmy_070_with_new_image_hosting/
---
[Post on Lemmy](https://dev.lemmy.ml/post/35293)

Let's start with the biggest news first: Lemmy is receiving funding from the [NLnet foundation](https://nlnet.nl/)! The funding is for a total amount of 45.000 €, which will allow [/u/dessalines](/u/dessalines)  and me ([/u/nutomic](/u/nutomic) ) to work on Lemmy full-time for at least half a year.

We have created various milestones for the work we are planning to do. Most of them are about getting ActivityPub federation ready for production. In addition, we will work on:
- better accessibility
- private communities and instances
- reworking search
- creating a `joinlemmy.ml` type site
-  the option to block other users or communities

The details of the milestones will be posted on our github issue tracker soon.

We're very excited about this opportunity, and can't wait to finish federation. 


In other news, we have just released [Lemmy v0.7.0.](https://github.com/LemmyNet/lemmy/blob/master/RELEASES.md#lemmy-v070-release-2020-06-2x) Most importantly, this update switches to [Pict-rs](https://git.asonix.dog/asonix/pict-rs/) for image hosting, due to various performance-related issues with Pictshare. Pict-rs was coded from scratch in Rust by the amazing @asonix, who also created the ActivityPub library for Rust. We can't thank him enough for all the work he is doing for Lemmy!

We'd also like to thank the following people for their contributions:

- @iav for their work in building arm compatible rust docker images and builds.
- @ernestwisniewski and @bytesnake for code contributions.
- Many others for contributing translations via the [Lemmy weblate.](https://weblate.yerbamate.dev/projects/lemmy/)
- Our [Patreon](https://www.patreon.com/dessalines) and [Liberapay](https://liberapay.com/Lemmy/) supporters who help us devote more time to Lemmy (We're still very far from these being able to sustain two developers)
- Everyone else who contributes to Lemmy, be it by coding, hosting instances or just using it and spreading the word!

Other than that, since v0.6.0 in January [we've closed over 100 issues](https://github.com/LemmyNet/lemmy/milestone/16?closed=1), fixed tons of bugs and added many new features.

[You can find the full changelog and upgrade instructions here](https://github.com/LemmyNet/lemmy/blob/master/RELEASES.md#lemmy-v070-release-2020-06-2x).
## [4][RLS vs rust-analyzer in 2020](https://www.reddit.com/r/rust/comments/hf07lk/rls_vs_rustanalyzer_in_2020/)
- url: https://www.reddit.com/r/rust/comments/hf07lk/rls_vs_rustanalyzer_in_2020/
---
What is the current state of RLS vs. rust-analyzer in 2020? I ask because all the reddit posts from 7 months ago say to use rust-analyzer as it's faster and better. But my experience in VSCode is the opposite. Like I'm halfway through writing a line and rust-analyzer will lint it "expected SEMICOLON" then spin my CPU fan up for 10 seconds doing something, then the error will disappear. It's basically unusable.

I've been using RLS up to now and haven't noticed a lot of differences, except rust-analyzer can underline variables. The 'Show implementations' and 'Run tests' functions don't work (I get "Command 'rust-analyzer.$CMD' not found" for both 'runSingle' and 'showReferences').

I'm using the 'Rust' VScode extension which has only just added support for rust-analyzer, is there a better way to integrate it into VScode? Also, my cargo project is using beta rust, does that make any difference over nightly?
## [5][Four years of Rust at OneSignal](https://www.reddit.com/r/rust/comments/hekt04/four_years_of_rust_at_onesignal/)
- url: https://onesignal.com/blog/4-years-of-rust-at-onesignal/
---

## [6][Help: Lifetimes on async functions with callbacks](https://www.reddit.com/r/rust/comments/hey4oa/help_lifetimes_on_async_functions_with_callbacks/)
- url: https://www.reddit.com/r/rust/comments/hey4oa/help_lifetimes_on_async_functions_with_callbacks/
---
When running code on the network it is often needed to run some async code until it succeeds.  In the small toy example below I have a connection object \`Conn\` that has an operation \`may\_fail\` that may fail some times, indicated by returning None, and succeed sometimes.  This is then used in a larger async function such as \`named\_fn\` which can then also optionally succeed.  I want to add a function \`retry\` that will call the \`named\_fn\` until it succeeds. But I cannot figure out how to get the lifetimes to play along.  I guess that what we want to hold for \`F\` is that the lifetime of the \`Conn\` object must outlive the returned Fut.  Are there anyone here familiar enough with higher order lifetimes and stuff to find a solution?

    struct Conn {
        value: u32,
    }
    
    impl Conn {
        async fn may_fail(&amp;mut self) -&gt; Option&lt;u32&gt; {
            self.value += 1;
            if self.value &lt; 10 {
                None
            } else {
                Some(self.value)
            }
        }
    }
    async fn named_fn(conn: &amp;mut Conn) -&gt; Option&lt;u32&gt; {
        conn.may_fail().await
    }
    
    async fn retry&lt;Fut: std::future::Future&lt;Output = Option&lt;u32&gt;&gt;, F: Fn(&amp;mut Conn) -&gt; Fut&gt;(
        conn: &amp;mut Conn,
        cb: F,
    ) -&gt; u32 {
        loop {
            if let Some(v) = cb(conn).await {
                return v;
            }
        }
    }
    
    #[tokio::main]
    async fn main() {
        let conn = Conn { value: 0 };
        let v = retry(&amp;mut conn, named_fn).await;
    
        // let v = retry(&amp;mut conn, |conn| {
        //     async {
        //         conn.may_fail().await
        //     }
        // });
        println!("v {}", v);
    }
## [7][Rust’s struct update syntax made simple](https://www.reddit.com/r/rust/comments/hezw9t/rusts_struct_update_syntax_made_simple/)
- url: https://medium.com/@Dylan.Kerler/rusts-struct-update-syntax-made-simple-569f78476460
---

## [8][View APNG with emulsion and image](https://www.reddit.com/r/rust/comments/heyr7j/view_apng_with_emulsion_and_image/)
- url: https://blog.image-rs.org/2020/06/23/apng-in-image.html
---

## [9][lemmy (federated alternative to reddit) v 0.7.0 released!](https://www.reddit.com/r/rust/comments/hemrd9/lemmy_federated_alternative_to_reddit_v_070/)
- url: https://github.com/LemmyNet/lemmy/releases/tag/v0.7.0
---

## [10][Simd By Cheating](https://www.reddit.com/r/rust/comments/helq3e/simd_by_cheating/)
- url: https://vorner.github.io/2020/06/21/simd-by-cheating.html
---

## [11][IPv6 and Rust](https://www.reddit.com/r/rust/comments/hee938/ipv6_and_rust/)
- url: https://blog.apnic.net/2020/06/02/ipv6-and-rust/
---

## [12][Announcing Bincode 1.3](https://www.reddit.com/r/rust/comments/hejjnt/announcing_bincode_13/)
- url: https://dos.cafe/blog/bincode-1.3.html
---

