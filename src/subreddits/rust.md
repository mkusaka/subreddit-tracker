# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (10/2020)!](https://www.reddit.com/r/rust/comments/fc7h20/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fc7h20/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/f8ney8/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (10/2020)?](https://www.reddit.com/r/rust/comments/fc7hn2/whats_everyone_working_on_this_week_102020/)
- url: https://www.reddit.com/r/rust/comments/fc7hn2/whats_everyone_working_on_this_week_102020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-10-2020/38939?u=llogiq)!
## [3][Inner loops are faster if they don't span a cache line](https://www.reddit.com/r/rust/comments/fcjx43/inner_loops_are_faster_if_they_dont_span_a_cache/)
- url: https://www.reddit.com/r/rust/comments/fcjx43/inner_loops_are_faster_if_they_dont_span_a_cache/
---
I was intrigued by the post [https://www.reddit.com/r/rust/comments/fbev13/the\_cost\_of\_indirection/](https://www.reddit.com/r/rust/comments/fbev13/the_cost_of_indirection/) , where using a trait seemed to slow things down. After all, traits should ideally be zero-cost abstractions in rust, and the case described seemed very simple.

After some analysis I've come to the conclusion that what is happening is that minor changes to the program file, compiler version etc, can cause the inner loop to be put at different places in memory, basically at random.

And it turns out that if the inner loop spans two cache lines, it is not as fast. On my machine 1 billion iterations take 450 ms or 900 ms, depending on the memory layout.

I've created a repo demonstrating this [https://github.com/avl/strange\_performance\_repro](https://github.com/avl/strange_performance_repro) .

Edit 2: In a previous edit I claimed even a loop spanning two cache lines could be fast, sometimes. I think I was wrong - I had accidentally checked objdump for an .o-file instead of the executable, and drawn the wrong conclusions.

Edit 3: D'oh! The title really should have been "loops are faster if they don't span two cachelines". Hopefully no one was confused by this bad title.
## [4][Commodity and Doublecount - libraries for representing currencies, and a double entry accounting system](https://www.reddit.com/r/rust/comments/fcrka3/commodity_and_doublecount_libraries_for/)
- url: https://www.reddit.com/r/rust/comments/fcrka3/commodity_and_doublecount_libraries_for/
---
Hi Rust Community,  
I created the following libraries while developing a cost sharing web application, and figured I could tidy them up and share them.

+ `doublecount` - [crates.io](https://crates.io/crates/doublecount) | [github](https://github.com/kellpossible/doublecount) - A double entry accounting system/library for Rust.
+ `commodity` - [crates.io](https://crates.io/crates/commodity) | [github](https://github.com/kellpossible/commodity) - A library for representing commodities/currencies, and exchange rates/conversions between them in Rust. Values are backed by the `rust_decimal` library.

Doublecount is inspired by the [beancount](http://furius.ca/beancount/) plain text accounting system, but currently has a simpler model, and exists only as a library at the moment. I developed it to serve as a base for calculations in the work in progress [coster](https://github.com/kellpossible/coster) cost sharing web application (aiming for something similar to [splittypie](https://github.com/cowbell/splittypie) but implemented entirely in Rust). I'm hoping that doublecount can be generic enough to serve the needs of some others, would like to hear if anyone else has some ideas about how they would like to use it, and improve the API.
## [5][rust-analyzer Changelog #14](https://www.reddit.com/r/rust/comments/fcdlq6/rustanalyzer_changelog_14/)
- url: https://rust-analyzer.github.io/thisweek/2020/03/02/changelog-14.html
---

## [6][First CLI "published" amazed by simplicity](https://www.reddit.com/r/rust/comments/fckb4o/first_cli_published_amazed_by_simplicity/)
- url: https://www.reddit.com/r/rust/comments/fckb4o/first_cli_published_amazed_by_simplicity/
---
I was struggling with my 40% keyboard not having numbers accessible directly and i3wm shortcuts. Fell into a rabbit hole and finished writing a CLI just to move my current window to the next workspace, creating it if it doesn't exist.

First I was so happy to find someone has [already wrapped i3 ipc socket](https://github.com/tmerr/i3ipc-rs) into a crate.

And then I wondered how I'll manage actually using my CLI. I see that cargo install just magically does the thing, and that it can work from git directly so that it'll be simple to use on another pc... That was a woah moment.

I've been able to do all that in less than two hours. Rust has such a wonderful ecosystem, thank you all &lt;3

Edit: I did not talked of StructOpt as it seemed to evident, but damn that's neat to use too!

PS: if anyone is curious, the two command cli is here https://gitlab.com/MonsieurMan/i3-utils
## [7][how can i send a closure to another thread without a 'static life time](https://www.reddit.com/r/rust/comments/fcoozh/how_can_i_send_a_closure_to_another_thread/)
- url: https://www.reddit.com/r/rust/comments/fcoozh/how_can_i_send_a_closure_to_another_thread/
---
thread\_pool.rs: [https://pastebin.com/wgxaJHq1](https://pastebin.com/wgxaJHq1)

worker.rs: [https://pastebin.com/rGqGXpDf](https://pastebin.com/rGqGXpDf)

(thread\_pool\_error.rs was omitted, its just an enum implementing Display, and Error)

I currently get an error on line 39 of thread\_pool.rs where Rust complains that `the parameter type \`F\` may not live long enough`.

Rust recommends that I add the constraint of `'static` to the type `F`, however I want be able to send closures that borrow non static references to my thread pool. Take this code for example in a `Server` which handles incoming connections.

    //implement a way for the loop to end later
    pub fn run(&amp;self) -&gt; Result&lt;(), Box&lt;dyn Error&gt;&gt;{
    	let listener = TcpListener::bind(format!("127.0.0.1:{}", self.port))?;
    	let thread_pool = ThreadPool::new(10)?;
    
    	for stream in listener.incoming(){
    		let stream = stream.unwrap();
    		let job = move ||{
    			self.handle_connection(stream);
    		};
    		thread_pool.execute(Box::new(job));
    	};
    
    	Ok(())
    }

As you can see from the `Drop` trait being implemented the references need only to live as long or longer than the thread\_pool struct, but I'm not sure how I can convey this information to Rust. What should I do?
## [8][IntelliJ Rust Changelog #117](https://www.reddit.com/r/rust/comments/fcdql8/intellij_rust_changelog_117/)
- url: https://intellij-rust.github.io/2020/03/02/changelog-117.html
---

## [9][Writing an OS in Rust: Updates in February 2020](https://www.reddit.com/r/rust/comments/fcadax/writing_an_os_in_rust_updates_in_february_2020/)
- url: https://os.phil-opp.com/status-update/2020-03-02/
---

## [10][Is there an IndexMut impl for HashMap collection struct in standard library?](https://www.reddit.com/r/rust/comments/fcmt8l/is_there_an_indexmut_impl_for_hashmap_collection/)
- url: https://www.reddit.com/r/rust/comments/fcmt8l/is_there_an_indexmut_impl_for_hashmap_collection/
---
One of my friend is a new Rust learner, he has some experience on other languages and began to learn Rust recently. He just raised this question to me but I can't decide a reason to it. Is this design absent by accident or on purpose? Why is it missing? Thank you all Redditors!
## [11][gsrs - PoC crate for creating generic movable self referencing structs on stable rust](https://www.reddit.com/r/rust/comments/fchwus/gsrs_poc_crate_for_creating_generic_movable_self/)
- url: https://www.reddit.com/r/rust/comments/fchwus/gsrs_poc_crate_for_creating_generic_movable_self/
---
Making some nontrivial workarounds about not being able to safely create movable self referencing structs made me thinking: what is the actual problem with self referencing structs in Rust. I know that you can create some self referencing structs locally. But when you do it locally compiler forces all items inside to have the same lifetime but having lifetime means it is not movable as a whole. So if all fields have the same lifetime, why can't we just erase this lifetime, move the struct, and apply new lifetime when we are using it in new place?

Well, looks like we can. Of course it is not that simple, but it is the general idea behind [gsrs](https://docs.rs/gsrs) - crate for creating generic movable self referencing structs(SRS). Basically after creation, when borrowing SRS to use it, it will apply lifetime of that borrow to all references inside. But to ensure safety it restricts ways you can interact with data.

API is not rich yet. It is in more of PoC/MVP state currently. I am pretty sure that my idea/implementation is sound(it even passes miri, although there is not quite many tests yet) but it is my second dive into unsafe Rust. Although I have a decent low level background I still could have missed something subtle. I have a lot of ideas of how to improve it so I would really appreciate a soundness review(it is quite small, ~150 lines of actual code), so I can be fully confident in my approach before fully committing on it and recommending someone else to use it. 

And some close to real world example of what you can do - movable Unicode suffix array
```rust
    use gsrs::*;
    struct TestRef&lt;'a&gt;(Vec&lt;&amp;'a str&gt;);
    deref_with_lifetime!(TestRef);

    let mut suffix_array = SRS::&lt;String, TestRef&gt;::create_with(
        "testtest こんにちは".to_owned(),
        |owner| {
            TestRef({
                let mut a = owner.chars();
                let mut vec = Vec::new();
                vec.push(a.as_str());
                while let Some(_) = a.next() {
                    vec.push(a.as_str());
                }
                vec.sort();
                vec
            })
       },
    );
    //move it wherever you like
    let mut suffix_array = black_box(suffix_array);
    fn contains(vec: &amp;[&amp;str], str: &amp;str) -&gt; bool {
        vec.binary_search_by(|&amp;it| {
            if it.starts_with(str) {
                Ordering::Equal
            } else {
                it.cmp(str)
            }
        }).is_ok()
    }
    let a = suffix_array.with(|user, _| contains(&amp;user.0, "ttes"));
    let c = suffix_array.with(|user, _| contains(&amp;user.0, "こん"));
    let str = "aa".to_owned();
    let b = suffix_array.with(move |user, _| contains(&amp;user.0, &amp;*str));
    assert!(a &amp;&amp; c &amp;&amp; !b);
```
## [12][WHY do you think IntelliJ is a good development tool?](https://www.reddit.com/r/rust/comments/fcshgy/why_do_you_think_intellij_is_a_good_development/)
- url: https://www.reddit.com/r/rust/comments/fcshgy/why_do_you_think_intellij_is_a_good_development/
---
Hi,

So I’ve been asking lately, in several communities, why people seem to agree that IntelliJ is a good tool. And… I haven’t seen anything convincing yet. Some context: I’m a neovim+tmux user, I use the UNIX philosophy on everything I do, and **my own quality criteria is that something must have a small scope and be excellent at what it pretends.** That has yielded a collection of tools I use on a daily basis, such as `nvim`, `fzf`, `git`, `fd`, `rg`, `nm`, `lldb`, `rr`, `tmux`, `i3wm`, etc.

So, I’m asking people who are already feeling comfortable with IntelliJ: why do you think it’s worth it? What makes it worth its price? This is a genuine question I’ve been asking and all I got was “It works for me” or “I’m faster with it.” I could say “I’m faster” too because I use EasyMotion, but if I were asked _why I’m faster with EasyMotion_, I’d explain that it integrates easily in neovim; allows me to jump everywhere in buffers with only two or three keystrokes, and then reduces muscle tiredness and stress; I’m also moving more “smoothly” and more comfortably.

Can someone tell me, in the same idea, about IntelliJ? Lots of people seem to like that tool, so I wanna know, especially because if I’m missing something, I might learn new workflow / techniques to add to my own workflow.

Thanks in advance!
