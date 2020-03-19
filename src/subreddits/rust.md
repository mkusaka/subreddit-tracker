# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (12/2020)!](https://www.reddit.com/r/rust/comments/fjef12/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fjef12/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

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
## [3][Announcing Rust IPFS, and a call for contributors](https://www.reddit.com/r/rust/comments/fl6ij9/announcing_rust_ipfs_and_a_call_for_contributors/)
- url: https://blog.ipfs.io/2020-03-18-announcing-rust-ipfs/
---

## [4][Unpopular opinion: Hyper is plenty good. You probably don't need a framework.](https://www.reddit.com/r/rust/comments/fkz7ui/unpopular_opinion_hyper_is_plenty_good_you/)
- url: https://www.reddit.com/r/rust/comments/fkz7ui/unpopular_opinion_hyper_is_plenty_good_you/
---
Whenever people ask here for a framework, people recommend actix, rocket, or maybe gotham or warp.

I just used hyper for a small project. It's amazing. It:

1. Does exactly what you ask for, no more and no less. There's no magic.
2. Is async
3. Works on stable rust.

Honestly, when I heard that it's minimalistic, I took it to mean that I'll be playing with sockets or something. In reality, it's _almost_ one step under the standard Go or servlet API (the only thing is that you'll have to import a crate to decode query_string and POST data). That's it. Debugging is a breeze (you know why it doesn't work because you write the router), footspace (in code) is minimal (so when it upgrades you just need to fix your router - no more need to go restructure your code base). It's just a breath of fresh air.
## [5][ttf-parser 0.5 - now with variable fonts support and C API](https://www.reddit.com/r/rust/comments/fl7kfx/ttfparser_05_now_with_variable_fonts_support_and/)
- url: https://www.reddit.com/r/rust/comments/fl7kfx/ttfparser_05_now_with_variable_fonts_support_and/
---
A new major version of [ttf-parser](https://github.com/RazrFalcon/ttf-parser) has been released, and since it gained a lot of new features, I wanted to advertise it a bit.

`ttf-parser` is a TrueType/OpenType parser like `stb_truetype` and `FreeType`, but unlike FreeType it doesn't render glyphs (stb_truetype also supports rendering, but a very primitive one). But the main difference from those projects is that `ttf-parser` is the only memory-safe, thread-safe and zero-allocation library (not sure about proprietary one from MS, Apple and Adobe). And more importantly, despite all the safety guarantees, it still the fastest (at least in my benchmarks). You can find more info in the readme.

In the latest version, `ttf-parser` gained a complete variable font support and C API. Right now it supports way more features than `stb_truetype` and has roughly the same TrueType support as FreeType (except rendering).

I hope a new version of rusttype with ttf-parser will be released soon, which will help with testing and adoption in the Rust ecosystem.
## [6][Dropbox gives up trying to port its sync engine to Python3 and rewrites it in Rust instead](https://www.reddit.com/r/rust/comments/fl9pdq/dropbox_gives_up_trying_to_port_its_sync_engine/)
- url: https://dropbox.tech/infrastructure/rewriting-the-heart-of-our-sync-engine
---

## [7][Announcing const-tweaker! Change your constants live from from a web GUI](https://www.reddit.com/r/rust/comments/fl8wpl/announcing_consttweaker_change_your_constants/)
- url: /r/rust_gamedev/comments/fl7593/announcing_consttweaker_change_your_constants/
---

## [8][Apple hiring Rust engineers for storage and networking groups](https://www.reddit.com/r/rust/comments/fkngza/apple_hiring_rust_engineers_for_storage_and/)
- url: https://twitter.com/benwilliamson/status/1240113606374686721
---

## [9][All Hands Retrospective | Inside Rust Blog](https://www.reddit.com/r/rust/comments/fkvcsr/all_hands_retrospective_inside_rust_blog/)
- url: https://blog.rust-lang.org/inside-rust/2020/03/18/all-hands-retrospective.html
---

## [10][Code Review AVL-Tree](https://www.reddit.com/r/rust/comments/fl8cwn/code_review_avltree/)
- url: https://www.reddit.com/r/rust/comments/fl8cwn/code_review_avltree/
---
Hey Guys,

I am pretty new to Rust. Currently I am doing daily coding challenges in mutiple programming languages. A few days ago there was the question to programm an AVL-tree. I had no problems in the other languages but Rust was completely different, even though i worked throught the  'The Book'. Now I am ready, and i would love to get some critics about my code and what i could do better:  


    struct Node{
        key: i32,
        left: Option&lt;Box&lt;Node&gt;&gt;,
        right: Option&lt;Box&lt;Node&gt;&gt;,
    }
    
    impl Node{
        fn new(key: i32) -&gt; Node{
            Node{
                key,
                left:None,
                right:None,
            }
        }
    
    
        fn height(pot_node: &amp;mut Option&lt;Box&lt;Node&gt;&gt;) -&gt; i32{
            match pot_node {
                Some(node) =&gt;{
                    let left_height = Node::height(&amp;mut node.left);
                    let right_height = Node::height(&amp;mut node.right);
    
                    let height = max(left_height, right_height) +1;
                    height
                },
                None =&gt; 0,
            }
        }
    
        fn rotate_right(mut node: Box&lt;Node&gt;) -&gt; Box&lt;Node&gt;{
            let mut y = node.left.take().unwrap();
            let mut b = y.right.take();
    
            node.left = b.take();
            y.right = Some(node).take();
    
            y
        }
    
        fn left_rotate(mut node: Box&lt;Node&gt;) -&gt; Box&lt;Node&gt;{
            let mut y = node.right.take().unwrap();
            let mut b = y.left.take();
    
            node.right = b.take();
            y.left =  Some(node).take();
    
            y
        }
    
        fn balance(pot_node: Option&lt;Box&lt;Node&gt;&gt;) -&gt; (Option&lt;Box&lt;Node&gt;&gt;, i32){
            match pot_node{
                None =&gt;{
                    return (None, 0);
                },
                Some(mut node) =&gt;{
                    let balance = Node::height(&amp;mut node.left)-Node::height(&amp;mut node.right);
    
                    return (Some(node),balance);
                }
            }
        }
    
        fn insert(pot_node : Option&lt;Box&lt;Node&gt;&gt;, key : i32) -&gt; Option&lt;Box&lt;Node&gt;&gt; {
            match pot_node{
                None =&gt;{
                  return Some(Box::new(Node::new(key)));
                },
                Some(mut node) =&gt;{
                    if key &lt; node.key{
                        node.left = Node::insert(node.left, key);
                    }else if key &gt; node.key{
                        node.right = Node::insert(node.right, key);
                    }else{
                        return Some(node);
                    }
    
                    let (tmp_node,balance) = Node::balance(Some(node).take());
                    let mut node = tmp_node.unwrap();
    
                    if balance &lt; -1{
                        match node.right.take(){
                            None =&gt;{
                                node.right = None;
                            },
                            Some(right_node) =&gt;{
                                if key &lt; right_node.key{
                                    node.right = Some(right_node);
                                    node.right = Some(Node::rotate_right(node.right.take().unwrap()));
                                    node = Node::left_rotate(node);
                                }else if key &gt; right_node.key{
                                    node.right = Some(right_node);
                                    node = Node::left_rotate(node);
                                }else{
                                    node.right = Some(right_node);
                                }
                            },
                        }
                    }else if balance &gt; 1{
                        match node.left.take() {
                            None =&gt;{
                                node.left = None;
                            },
                            Some(left_node) =&gt; {
                                if key &gt; left_node.key{
                                    node.left = Some(left_node);
                                    node.left = Some(Node::left_rotate(node.left.take().unwrap()));
                                    node = Node::rotate_right(node);
                                }else if key &lt; left_node.key{
                                    node.left = Some(left_node);
                                    node = Node::rotate_right(node);
                                }else{
                                    node.right = Some(left_node);
                                }
                            },
                        }
                    }
                    return Some(node);
                }
            };
        }
    
        fn pre_order_print(pot_node : Option&lt;Box&lt;Node&gt;&gt;){
            match pot_node {
                Some(node) =&gt;{
                    print!("{} ", node.key);
                    Node::pre_order_print(node.left);
                    Node::pre_order_print(node.right);
                },
                None =&gt; (),
            }
        }
    }

Best regards :)
## [11][Blog post: Parsing Library in Rust pt. 3 - TextSize crate](https://www.reddit.com/r/rust/comments/fky8pt/blog_post_parsing_library_in_rust_pt_3_textsize/)
- url: https://blog.frondeus.pl/parser-3/
---

## [12][Quick question about unused generic type parameter](https://www.reddit.com/r/rust/comments/fkulrf/quick_question_about_unused_generic_type_parameter/)
- url: https://www.reddit.com/r/rust/comments/fkulrf/quick_question_about_unused_generic_type_parameter/
---
I'm trying to write a type that allows me to call a function once and save the result of that one function call.

I have a definition written for the struct like this

    struct CallOnce&lt;T, U, V&gt;
    where
        T: Fn(U) -&gt; V,
    {
        value: Option&lt;V&gt;,
        calculate: T,
    }

&amp;#x200B;

but I'm getting this error

    parameter `U` is never used
    
    unused parameter
    
    help: consider removing `U`, referring to it in a field, or using a marker such as `std::marker::PhantomData`rustc(E0392)main.rs(1, 18): unused parameter

This is confusing to me because it looks like I'm using U as a generic input parameter for type T (Fn(U) -&gt; V). I'm wondering how I can fix this error.

For what it's worth, here's all my code if that helps.

    struct CallOnce&lt;T, U, V&gt;
    where
        T: Fn(U) -&gt; V,
    {
        value: Option&lt;V&gt;,
        calculate: T,
    }
    
    impl&lt;T, U, V&gt; CallOnce&lt;T, U, V&gt;
    where
        T: Fn(U) -&gt; V,
    {
        fn new(calculate: T) -&gt; CallOnce&lt;T, U, V&gt; {
            CallOnce {
                value: None,
                calculate,
            }
        }
    
        fn call(&amp;mut self, arg: U) -&gt; V {
            match self.value {
                Some(value) =&gt; value,
                None =&gt; {
                    let value = (self.calculate)(arg);
                    self.value = Some(value);
                    value
                }
            }
        }
    }
    
    fn main() {
        let mut add_one = CallOnce::new(|x| {
            println!("Calling with {}", x);
            x + 1
        });
        println!("Result 1: {}", add_one.call(1));
        println!("Result 2: {}", add_one.call(2));
    }
