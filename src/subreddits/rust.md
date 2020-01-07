# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (2/2020)!](https://www.reddit.com/r/rust/comments/ekpqp7/hey_rustaceans_got_an_easy_question_ask_here_22020/)
- url: https://www.reddit.com/r/rust/comments/ekpqp7/hey_rustaceans_got_an_easy_question_ask_here_22020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/ehk18j/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What’s everyone working on this week (2/2020)?](https://www.reddit.com/r/rust/comments/ekpr6w/whats_everyone_working_on_this_week_22020/)
- url: https://www.reddit.com/r/rust/comments/ekpr6w/whats_everyone_working_on_this_week_22020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-2-2020/36546?u=llogiq)!
## [3][Translating Quake 3 into Rust](https://www.reddit.com/r/rust/comments/el2z0k/translating_quake_3_into_rust/)
- url: https://immunant.com/blog/2020/01/quake3/
---

## [4][Writing an OS in Rust: Updates in December 2019](https://www.reddit.com/r/rust/comments/elas6x/writing_an_os_in_rust_updates_in_december_2019/)
- url: https://os.phil-opp.com/status-update/2020-01-07/
---

## [5][What applications are better written in Rust than in other languages?](https://www.reddit.com/r/rust/comments/el9p2x/what_applications_are_better_written_in_rust_than/)
- url: https://www.reddit.com/r/rust/comments/el9p2x/what_applications_are_better_written_in_rust_than/
---
Currently I know Java/Kotlin for backends &amp; Android, Javascript/Typescript for web &amp; mobile development, and C++ for games programming &amp; microcontrollers.

I keep googling on what Rust can do, and I keep seeing results about people saying you can do anything. While that's true, I have no interest in writing backend servers or web frontends in Rust, when the frameworks &amp; libraries are far more mature in other languages. I would only write in Rust if there are clear benefits for doing so.

I presume Rust overlaps with C++, which means it would be a better choice for writing games and microcontrollers. However I don't see a lot of game engines supporting Rust. I also don't see any mainstream support for Rust with Microcontrollers either, especially with Arduino. Certainly not enough to give up the amount of C++ libraries which are currently supported.

So, in what applications does it make more sense to write in Rust? I really want to learn the language but it's hard to find a project where it would make more sense over the alternatives.
## [6][Informal survey: PGO](https://www.reddit.com/r/rust/comments/elaghi/informal_survey_pgo/)
- url: https://www.reddit.com/r/rust/comments/elaghi/informal_survey_pgo/
---
Curious how many people are using profile guided optimisation in Rust and how effective it has been for them.

[Rustc PGO docs](https://rust-lang.github.io/rustc-guide/profile-guided-optimization.html)
## [7][[GitPorn] Repository summary on your terminal](https://www.reddit.com/r/rust/comments/elas5f/gitporn_repository_summary_on_your_terminal/)
- url: https://i.redd.it/g69jqzgtlc941.jpg
---

## [8][Let The Compiler Do The Work](https://www.reddit.com/r/rust/comments/eksduv/let_the_compiler_do_the_work/)
- url: http://cliffle.com/p/dangerust/6/
---

## [9][Problem with RefCell](https://www.reddit.com/r/rust/comments/el9cy0/problem_with_refcell/)
- url: https://www.reddit.com/r/rust/comments/el9cy0/problem_with_refcell/
---
```
use std::rc::Rc;
use std::cell::RefCell;

type Work = Box&lt;dyn FnMut() + Send + 'static&gt;;
struct Inst {
    work: Work,
}

fn main() {
    let work = Box::new(||{ println!("hello"); });
    let a = Rc::new(RefCell::new(Inst{work}));
    (a.borrow_mut().work)();
}
```
This code does not compile. The error is:
```
error[E0596]: cannot borrow data in a dereference of `std::cell::RefMut&lt;'_, Inst&gt;` as mutable
  --&gt; src/main.rs:11:5
   |
11 |     (a.borrow_mut().work)();
   |     ^^^^^^^^^^^^^^^^^^^^^ cannot borrow as mutable
   |
   = help: trait `DerefMut` is required to modify through a dereference, but it is not implemented for `std::cell::RefMut&lt;'_, Inst&gt;`
```

What's going on here? Doesn't RefMut implement DerefMut?
## [10][Functional vs imperative styles - which of these is better form in Rust and how can this code be improved?](https://www.reddit.com/r/rust/comments/el7wpv/functional_vs_imperative_styles_which_of_these_is/)
- url: https://www.reddit.com/r/rust/comments/el7wpv/functional_vs_imperative_styles_which_of_these_is/
---
Hello, I'm a long time systems developer but new to Rust. I'm writing my first crate which will be a cryptographic library. Please advise me which of these functions is better written and how I can improve them:

```
            pub fn multiparty_keygen(&amp;self) -&gt; (Vec&lt;SecretKey&gt;, Vec&lt;VerifyKey&gt;) {
                let attributes_size = self.params.hs.len();
                assert!(self.authorities_total &gt;= self.threshold);
                assert!(attributes_size &gt; 0);
        
                let create_n_random_scalars = |n| -&gt; Vec&lt;_&gt; {
                    (0..n).map(|_| self.params.random_scalar()).collect()
                };
        
                // Generate polynomials
                let v_poly = create_n_random_scalars(self.threshold);
                let w_poly: Vec&lt;Vec&lt;bls::Scalar&gt;&gt; =
                    (0..attributes_size).map(|_| create_n_random_scalars(self.threshold)).collect();
        
                // Generate shares
                let x_shares: Vec&lt;bls::Scalar&gt; = (1..=self.authorities_total).map(
                    |i| compute_polynomial(&amp;v_poly, i as u64)).collect();
                let y_shares: Vec&lt;Vec&lt;bls::Scalar&gt;&gt; = (1..=self.authorities_total).map(
                    |i| w_poly.iter().map(move |w_coefficients|
                        compute_polynomial(&amp;w_coefficients, i as u64)).collect()).collect();
        
                // Set the keys
                // sk_i = (x, (y_1, y_2, ..., y_q))
                // vk_i = (g2^x, (g2^y_1, g2^y_2, ..., g2^y_q)) = (a, (B_1, B_2, ..., B_q))
                let verify_keys: Vec&lt;VerifyKey&gt; =
                    x_shares.iter()
                        .zip(y_shares.iter())
                        .map(
                            |(x, y_share_parts)| 
                            VerifyKey {
                                alpha: self.params.g2 * x,
                                beta: y_share_parts.iter().map(|y| self.params.g2 * y).collect()
                        }).collect();
                // We are moving out of x_shares into SecretKey, so this line happens
                // after creating verify_keys to avoid triggering borrow checker.
                let secret_keys: Vec&lt;SecretKey&gt; =
                    x_shares.iter().zip(y_shares)
                        .map(move |(&amp;x, y)| SecretKey{ x: x, y: y }).collect();
        
                (secret_keys, verify_keys)
            }
```

And here is the imperative version:

```
            pub fn multiparty_keygen2(&amp;self) -&gt; (Vec&lt;SecretKey&gt;, Vec&lt;VerifyKey&gt;) {
                let attributes_size = self.params.hs.len();
                assert!(self.authorities_total &gt;= self.threshold);
                assert!(attributes_size &gt; 0);
        
                let create_n_random_scalars = |n| -&gt; Vec&lt;_&gt; {
                    (0..n).map(|_| self.params.random_scalar()).collect()
                };
        
                // Generate polynomials
                let v_poly = create_n_random_scalars(self.threshold);
                let w_poly: Vec&lt;Vec&lt;bls::Scalar&gt;&gt; =
                    (0..attributes_size).map(|_| create_n_random_scalars(self.threshold)).collect();
        
                // Generate shares
                let mut x_shares = Vec::new();
                for i in 1..=self.authorities_total {
                    x_shares.push(compute_polynomial(&amp;v_poly, i as u64));
                }
        
                let mut y_shares = Vec::new();
                for i in 1..=self.authorities_total {
                    let mut y_parts = Vec::new();
                    for w_coefficients in &amp;w_poly {
                        y_parts.push(compute_polynomial(w_coefficients, i as u64));
                    }
                    y_shares.push(y_parts);
                }
        
                // Set the keys
                // sk_i = (x, (y_1, y_2, ..., y_q))
                // vk_i = (g2^x, (g2^y_1, g2^y_2, ..., g2^y_q)) = (a, (B_1, B_2, ..., B_q))
                let mut verify_keys = Vec::new();
                for (x, y_parts) in x_shares.iter().zip(y_shares.iter()) {
                    verify_keys.push(VerifyKey {
                        alpha: self.params.g2 * x,
                        beta: {
                            let mut beta = Vec::new();
                            for y in y_parts {
                                beta.push(self.params.g2 * y);
                            }
                            beta
                        }
                    });
                }
        
                // We are moving out of x_shares into SecretKey, so this line happens
                // after creating verify_keys to avoid triggering borrow checker.
                let mut secret_keys = Vec::new();
                for (&amp;x, y) in x_shares.iter().zip(y_shares) {
                    secret_keys.push(SecretKey{ x: x, y: y });
                }
        
                (secret_keys, verify_keys)
            }
```

What do you guys think between the two? Which is better written and how can we improve either of them?

Thanks
## [11][dtool 0.5.0 - A command line tool collection to assist development written in RUST](https://www.reddit.com/r/rust/comments/ela0wu/dtool_050_a_command_line_tool_collection_to/)
- url: https://www.reddit.com/r/rust/comments/ela0wu/dtool_050_a_command_line_tool_collection_to/
---
dtool 0.5.0 released

**About dtool**

dtool is a command line tool collection to assist development

 [Full description on github](https://github.com/guoxbin/dtool)

**New features**

* Pbkdf2
* Case conversion (upper, lower, title, camel, pascal, snake, shouty snake, kebab)
* New hashing algorithm (crc, blake2b)

**Examples**

|Sub command| Desc |                                                                   Example                                                                   |Remark|Since |
|-----------|------|---------------------------------------------------------------------------------------------------------------------------------------------|------|------|
|  pbkdf2   |Pbkdf2|$ dtool pbkdf2 -a sha2_256 -s 0x646566 -i 2 -l 256\ &lt;br&gt; 0x616263&lt;br&gt;0x51a30556d0d133d859d3f3da86f861b7b12546c4f9a193eb\ &lt;br&gt;b374397467872514|      |v0.5.0|
|   case    |Case conversion|$ dtool case -t upper &amp;#x27;good tool&amp;#x27;&lt;br&gt;GOOD TOOL|   Upper case    |v0.5.0|
|   case    |Case conversion|$ dtool case -t lower &amp;#x27;GOOD TOOL&amp;#x27;&lt;br&gt;good tool|   Lower case    |v0.5.0|
|   case    |Case conversion|$ dtool case -t title &amp;#x27;GOOD TOOL&amp;#x27;&lt;br&gt;Good Tool|   Title case    |v0.5.0|
|   case    |Case conversion|$ dtool case -t camel &amp;#x27;GOOD TOOL&amp;#x27;&lt;br&gt;goodTool |   Camel case    |v0.5.0|
|   case    |Case conversion|$ dtool case -t pascal &amp;#x27;GOOD TOOL&amp;#x27;&lt;br&gt;GoodTool|   Pascal case   |v0.5.0|
|   case    |Case conversion|      $ dtool case -t snake GoodTool&lt;br&gt;good_tool       |   Snake case    |v0.5.0|
|   case    |Case conversion|   $ dtool case -t shouty_snake GoodTool&lt;br&gt;GOOD_TOOL   |Shouty snake case|v0.5.0|
|   case    |Case conversion|      $ dtool case -t kebab GoodTool&lt;br&gt;good-tool       |   Kebab case    |v0.5.0|


**Installation**

&gt; cargo install dtool

**Suggestion or contribution**

[github](https://github.com/guoxbin/dtool)
## [12][Announcing the v0.11 release of Yew](https://www.reddit.com/r/rust/comments/ektzvd/announcing_the_v011_release_of_yew/)
- url: https://www.reddit.com/r/rust/comments/ektzvd/announcing_the_v011_release_of_yew/
---
Hello, I'm excited to share with all of you the latest Yew release :)

*(If you're not familiar, Yew is a framework for building client web apps with Rust &amp; WebAssembly!)*

The latest release contains quite a few breaking changes and has a [transition guide](https://github.com/yewstack/yew/blob/0.11.0/CHANGELOG.md#transition-guide) for projects that need to be migrated. The reason there are a number of breaking changes is because we are preparing for a 1.0 release of Yew in the near future. Additionally, there has been a lot of interest from the Yew community in building and using component libraries built on Yew and so this release adds some functionality to support that effort. For more info, check out the [full changelog](https://github.com/yewstack/yew/releases).

Also, I'd like to announce Yew's new documentation website that aims to help developers learn about Yew and start building Yew apps more quickly. You can find it here: [https://yew.rs](https://yew.rs). *Note that it's still a WIP, PR's welcome ;)*

For the next release, we will be adding support for `web-sys` (while preserving support for `stdweb` and Emscripten targets) and will start integrating with the [gloo](https://github.com/rustwasm/gloo) toolkit. Also in the works.. we have folks working on CSS tooling, component libraries, and server side rendering. Come hang out in our [Gitter chat](https://gitter.im/yewframework/Lobby) if you'd like to learn more!  


EDIT: Just created a Twitter account for Yew ([@yewstack](https://twitter.com/yewstack)). Feel free to follow if you would like more frequent updates about the project :)
