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
## [3][Writing Python inside Rust](https://www.reddit.com/r/rust/comments/g3kxid/writing_python_inside_rust/)
- url: https://blog.m-ou.se/writing-python-inside-rust-1/
---

## [4][Rust Survey 2019 Results](https://www.reddit.com/r/rust/comments/g3bikn/rust_survey_2019_results/)
- url: https://blog.rust-lang.org/2020/04/17/Rust-survey-2019.html
---

## [5][Boa releases v0.7 with a much faster parser](https://www.reddit.com/r/rust/comments/g3m5i7/boa_releases_v07_with_a_much_faster_parser/)
- url: https://github.com/jasonwilliams/boa/blob/master/CHANGELOG.md
---

## [6][Zinc Framework: the ZK circuit programming language and VM](https://www.reddit.com/r/rust/comments/g3m2xb/zinc_framework_the_zk_circuit_programming/)
- url: https://github.com/matter-labs/zinc
---

## [7][Why does this lifetime mismatch occur? (expected enum `Node&lt;'a&gt;` - found enum `Node&lt;'static&gt;`)](https://www.reddit.com/r/rust/comments/g3kbje/why_does_this_lifetime_mismatch_occur_expected/)
- url: https://www.reddit.com/r/rust/comments/g3kbje/why_does_this_lifetime_mismatch_occur_expected/
---
The following code results in the error

    error[E0308]: mismatched types
     --&gt; src/lib.rs:5:9
      |
    5 |         br(),
      |         ^^^^ lifetime mismatch
      |
      = note: expected enum `Node&lt;'a&gt;`
                 found enum `Node&lt;'static&gt;`
    note: [...]

[Playground link](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=acdac6bbf31b657d3d53ec86c9018c9b)

    use std::borrow::Cow;
    
    fn vec&lt;'a&gt;(text: &amp;'a str) -&gt; Vec&lt;Node&lt;'a&gt;&gt; {
        vec![
            br(),
            p(text),
        ]
    }
    
    fn br() -&gt; Node&lt;'static&gt; {
        Node::Text("".into())
    }
    
    fn p(text: &amp;str) -&gt; Node {
        Node::Text(text.into())
    }
    
    #[derive(Clone)]
    enum Node&lt;'a&gt; {
        Element(Cow&lt;'a, Element&lt;'a&gt;&gt;),
        Text(Cow&lt;'a, str&gt;),
    }
    
    #[derive(Clone)]
    struct Element&lt;'a&gt; {
        children: Vec&lt;Node&lt;'a&gt;&gt;,
    }

I can actually resolve the error by replacing `Element(Cow&lt;'a, Element&lt;'a&gt;&gt;)` with `Element(Element&lt;'a)` or `Element(Cow&lt;'a, Element&lt;'static&gt;&gt;)` in `enum Node&lt;'a&gt;`. But I still don't understand what the problem is. Could you explain why the error occurs?

My understanding would be that any type with a `'static` lifetime could be assigned to any other type with a shorter lifetime.

Thanks in advance for any help :)
## [8][Crates implementing FromIterator&lt;&gt; for std collection types, best practices](https://www.reddit.com/r/rust/comments/g3iqan/crates_implementing_fromiterator_for_std/)
- url: https://www.reddit.com/r/rust/comments/g3iqan/crates_implementing_fromiterator_for_std/
---
I recently decided to add a new dependency into a crate. However, I started noticing that a downstream crate dependent on mine would break simply by \`use\`ing the new crate:

    use the_new_package; // All that was needed to trigger the downstream crate to fail

After some digging, I was able to reproduce the following, which simulates what was going on:

    // The crate I was importing had something like this...
    pub struct Test;
    
    impl FromIterator&lt;Test&gt; for Vec&lt;u8&gt; {
        fn from_iter&lt;I: IntoIterator&lt;Item = Test&gt;&gt;(iter: I) -&gt; Self {
            let v: Vec&lt;u8&gt; = iter.into_iter().skip(2).collect();
            v
        }
    }
    
    // The downstream crate that broke had something like this...
    use rand::thread_rng;
    use rand::distributions::{Distribution, Standard};
    
    fn main() {
      let mut rng = thread_rng();
      let v: Vec&lt;u8&gt; = Standard.sample_iter(&amp;mut rng).take(5).collect();
      println!("{:}", v);
    }

The error message became:

    error[E0282]: type annotations needed
      --&gt; src/main.rs:16:31
       |
    16 |     let v: Vec&lt;u8&gt; = Standard.sample_iter(&amp;mut rng).take(5).collect();
       |                               ^^^^^^^^^^^
       |                               |
       |                               cannot infer type for type parameter `T`
       |                               help: consider specifying the type argument in the method call: `sample_iter::&lt;R&gt;`

I now understand that the problem is that in trying to infer how to convert the `Iterator&lt;Item=T&gt;` that `take(5)` will return, there are now two options: std's definition of `FromIterator&lt;u8&gt; for &lt;Vec&lt;u8&gt;&gt;`, and the new `FromIterator&lt;Test&gt; for Vec&lt;u8&gt;`. Since `sample_iter` has its iterator item type generic, it's impossible for the type checker to infer  what type of items the iterator returned from `sample_iter` will have.

So, my question:

I believe that the `let v: Vec&lt;u8&gt; = ...;` is idiomatic Rust, and yet this crate was able to break that code simply by including it. Is the way that the crate did this considered an okay practice and this is something that sometimes just happens? Or would it be considered bad form for a crate to do something like this, and I should consider not using that crate or requesting a change? Or is this pretty contextual? [Here is the crate implementing FromIterator](https://github.com/tormol/encode_unicode/blob/ca2e5fc3393d440e6c933e3bdb094680241ddc93/src/utf8_char.rs#L127-L142), where it is seems to be trying to convert u8 iterators into utf8-aware iterators.

Thanks
## [9][Should I pass by immutable move or just by value](https://www.reddit.com/r/rust/comments/g3jsrz/should_i_pass_by_immutable_move_or_just_by_value/)
- url: https://www.reddit.com/r/rust/comments/g3jsrz/should_i_pass_by_immutable_move_or_just_by_value/
---
So I'm learning Rust now. When I was learning C/C++ in school, teachers were always telling us to pass by constant reference whenever we could because it was kinda safe and took less memory.

Should I do the same in rust? All other things equal, should I prefer to

    fn cool_function(cool_parameter: &amp;thing::one)

over

    fn cool_function(cool_parameter: thing::one)

I don't know a lot about how borrowing and changing ownership is implemented. If someone could explain that in the comments that would be much appreciated!!
## [10][yet another batch rename utility written in async Rust](https://www.reddit.com/r/rust/comments/g3729g/yet_another_batch_rename_utility_written_in_async/)
- url: https://github.com/yaa110/nomino
---

## [11][Guidance to learn Rust](https://www.reddit.com/r/rust/comments/g3k9xq/guidance_to_learn_rust/)
- url: https://www.reddit.com/r/rust/comments/g3k9xq/guidance_to_learn_rust/
---
Howdy Rustaceans! 

Hope everyone is safe and sound! I work at a major fintech as a Platform software engineer. We use Golang to manage our K8S abstraction layer and java/python for Jenkins-related utils. I'm a Java developer primarily. However, a colleague got me into Rust and since then I've joined this sub. To learn and understand this language, I've started reading [The Rust book](https://doc.rust-lang.org/book/). However, in order to get my hands, is there anything else I should do? Either/both on personal and professional front? Should I start writing some utils in Rust? I wish to grow in this domain (systems/platform programming) as I'm just 3 years into this profession. 

Thanks in advance!
## [12][Upgrade Rust's Android SDK to API level 16 (from 14)](https://www.reddit.com/r/rust/comments/g30bpg/upgrade_rusts_android_sdk_to_api_level_16_from_14/)
- url: https://github.com/rust-lang/rust/pull/71123
---

