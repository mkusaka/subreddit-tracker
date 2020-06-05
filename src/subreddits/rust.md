# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (23/2020)!](https://www.reddit.com/r/rust/comments/guo4c1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/guo4c1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/gr3r32/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 341](https://www.reddit.com/r/rust/comments/gvwvep/this_week_in_rust_341/)
- url: https://this-week-in-rust.org/blog/2020/06/02/this-week-in-rust-341/
---

## [3][I wrote a small virtual machine in rust](https://www.reddit.com/r/rust/comments/gx3emt/i_wrote_a_small_virtual_machine_in_rust/)
- url: https://www.reddit.com/r/rust/comments/gx3emt/i_wrote_a_small_virtual_machine_in_rust/
---
This is my first rust program. Would love some feedback :)

[github.com/tarekwiz/smallvm](https://github.com/tarekwiz/smallvm)
## [4][Announcing Rust 1.44.0](https://www.reddit.com/r/rust/comments/gwmmum/announcing_rust_1440/)
- url: https://blog.rust-lang.org/2020/06/04/Rust-1.44.0.html
---

## [5][toshi: A full-text search engine in rust](https://www.reddit.com/r/rust/comments/gwvu6x/toshi_a_fulltext_search_engine_in_rust/)
- url: https://github.com/toshi-search/Toshi
---

## [6][Jubi Taneja presents "Improving Compiler Construction Using Formal Methods"](https://www.reddit.com/r/rust/comments/gwy3d9/jubi_taneja_presents_improving_compiler/)
- url: https://www.youtube.com/watch?v=de8Ak0nY1hA
---

## [7][Tour of WebAssembly - Chapter 2 Released](https://www.reddit.com/r/rust/comments/gwt37e/tour_of_webassembly_chapter_2_released/)
- url: https://tourofrust.com/webassembly/chapter_2_en.html
---

## [8][My first blog, about "how to read Rust docs"](https://www.reddit.com/r/rust/comments/gwnl0w/my_first_blog_about_how_to_read_rust_docs/)
- url: https://www.reddit.com/r/rust/comments/gwnl0w/my_first_blog_about_how_to_read_rust_docs/
---
Hi everyone, 

I'm learning Rust and I found somehow rust docs hard to understand for a newbie. So I summarize something that may help reading rust docs, mainly about the function signatures and navigate modules with the aware of some key traits. 

[How to read Rust docs](http://blog.frqc.info/post/how-to-read-rust-docs/)

This is my post about rust. I know it's far from a conclusive summary and please correct me if there is anything not correct or precise enough in the post. 

&amp;#x200B;

Thank you and looking forward to your feedback!
## [9][ct_python!, a Rust macro to execute Python code at compile time and to insert the outputs as Rust code back into your Rust program, by Mara Bos](https://www.reddit.com/r/rust/comments/gwe7ut/ct_python_a_rust_macro_to_execute_python_code_at/)
- url: https://twitter.com/m_ou_se/status/1268469932544733186
---

## [10][Is it anti pattern to define a method in a trait to return a struct(to avoid lots of getter/setter)?](https://www.reddit.com/r/rust/comments/gx0esy/is_it_anti_pattern_to_define_a_method_in_a_trait/)
- url: https://www.reddit.com/r/rust/comments/gx0esy/is_it_anti_pattern_to_define_a_method_in_a_trait/
---
Hi! I am rather new to rust and quite confused about the way OOP is implemented in rust. I've got the following code to implement a basic class inheritance:

    #[derive(Debug, Default)]
    pub struct BaseData {
        a: f64,
        b: f64,
        c: f64,
        d: f64,
    }
    
    pub trait BaseCompute {
        fn get_base(&amp;self) -&gt; &amp;BaseData;
        fn get_mut_base(&amp;mut self) -&gt; &amp;mut BaseData;
        // Thus avoiding lots of "getter/setter"
    
        // Default implementation
        fn do_something(&amp;self) {
            // Do computation with: self.get().a, self.get().b, self.get().c
        }
    }
    
    pub trait CrazyCompute:BaseCompute{
        fn do_something_crazy(&amp;self);
    }
    
    pub struct Crazy{
        base: BaseData,
        x: f64,
    }
    
    // We only need to implement get_base and get_mut_base to implement BaseCompute
    impl BaseCompute for Crazy {
        fn get_base(&amp;self)-&gt;&amp;BaseData{
            &amp;self.base
        }
        fn get_mut_base(&amp;mut self)-&gt;&amp;mut BaseData{
            &amp;mut self.base
        }
    }
    
    impl CrazyCompute for Crazy{
        fn do_something_crazy(&amp;self){
    
        }
    }

I've defined a `get_base` and a `get_mut_base` in trait `BaseCompute` so that when another trait needs to **inherit** from `BaseCompute`, they don't need to implement a lot of `get_a`, `get_b` etc. I wonder if this is the correct way of OOP in rust, or if there's better way to implement inheritance in rust?
## [11][Rust has no goto? Now you can at least jump forward.](https://www.reddit.com/r/rust/comments/gwqrv6/rust_has_no_goto_now_you_can_at_least_jump_forward/)
- url: https://www.reddit.com/r/rust/comments/gwqrv6/rust_has_no_goto_now_you_can_at_least_jump_forward/
---
I made a crate that enables you to safely (borrow checker friendly!) jump forward in the control-flow [forward\_goto](https://crates.io/crates/forward_goto)

```rust
use forward_goto::rewrite_forward_goto;

#[rewrite_forward_goto]
fn decode_msg(luck: impl Fn() -&gt; bool, is_alan_turing: bool) {
    if is_alan_turing {
        // jump forward from here
        forward_goto!('turing_complete);
    }

    println!("You'll need a lot of luck for this");

    if luck() {
        println!("Seems you were lucky");

        forward_label!('turing_complete);
        // directly to here
        println!("Message decoded!");
    } else {
        println!("No luck today...");
    }
}
```

Should you use it? No, probably not...
## [12][Rather interesting take on visualizing ownership/borrowing from 2017](https://www.reddit.com/r/rust/comments/gwaqlx/rather_interesting_take_on_visualizing/)
- url: https://rufflewind.com/2017-02-15/rust-move-copy-borrow
---

