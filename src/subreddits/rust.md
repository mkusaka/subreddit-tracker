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

The Rust-related IRC channels on irc.mozilla.org (click the links to open a web-based IRC client):

 - [#rust](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust) (general questions)
 - [#rust-beginners](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-beginners) (beginner questions)
 - [#cargo](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23cargo) (the package manager)
 - [#rust-gamedev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-gamedev) (graphics and video games, and see also [/r/rust_gamedev](https://www.reddit.com/r/rust_gamedev))
 - [#rust-osdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-osdev) (operating systems and embedded systems)
 - [#rust-webdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-webdev) (web development)
 - [#rust-networking](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-networking) (computer networking, and see also [/r/rust_networking](https://www.reddit.com/r/rust_networking))

Also check out [last week's thread](https://reddit.com/r/rust/comments/f8ney8/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (10/2020)?](https://www.reddit.com/r/rust/comments/fc7hn2/whats_everyone_working_on_this_week_102020/)
- url: https://www.reddit.com/r/rust/comments/fc7hn2/whats_everyone_working_on_this_week_102020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-10-2020/38939?u=llogiq)!
## [3][Writing an OS in Rust: Updates in February 2020](https://www.reddit.com/r/rust/comments/fcadax/writing_an_os_in_rust_updates_in_february_2020/)
- url: https://os.phil-opp.com/status-update/2020-03-02/
---

## [4][No Nonsense Neovim Client in Rust](https://www.reddit.com/r/rust/comments/fc1l3y/no_nonsense_neovim_client_in_rust/)
- url: https://github.com/Kethku/neovide
---

## [5][git-trim: A git tool to trim merged local/remote branches written in Rust.](https://www.reddit.com/r/rust/comments/fc4ee6/gittrim_a_git_tool_to_trim_merged_localremote/)
- url: https://github.com/foriequal0/git-trim
---

## [6][v0.7 of Gleam, a statically typed language written in Rust for the Erlang VM, is out](https://www.reddit.com/r/rust/comments/fbz7lg/v07_of_gleam_a_statically_typed_language_written/)
- url: https://lpil.uk/blog/gleam-v0.7-released/
---

## [7][Announcing the v0.13 release of Yew!](https://www.reddit.com/r/rust/comments/fbweym/announcing_the_v013_release_of_yew/)
- url: https://www.reddit.com/r/rust/comments/fbweym/announcing_the_v013_release_of_yew/
---
Hello, I'm excited to share with all of you the latest Yew release :)

*(If you're not familiar, Yew is a framework for building client web apps with Rust &amp; WebAssembly!)*

In [this release](https://github.com/yewstack/yew/releases/tag/0.13.0), we have added support for building web apps with the foundational `web-sys` crate from the [Rust and Web Assembly Working Group](https://rustwasm.github.io/). We have also started integrating the `gloo` crate (also from the rust/wasm working group) for event listeners.

Another big change in this release is an update to how `Component` properties are specified. For context, Yew allows **compile time** property checking when declaring components with "JSX" style syntax. Before this release, properties were treated as optional by default, and struct fields could be annotated with a macro attribute if they were to be treated as required (forgetting to pass a required property results in a compile error). For this release, we flipped the default behavior. Properties are treated as required by default and optional if annotated as such. The new syntax takes advantage of [arbitrary token custom attributes released in Rust 1.34](https://blog.rust-lang.org/2019/04/11/Rust-1.34.0.html#custom-attributes-accept-arbitrary-token-streams) and looks like this:

    #[derive(Clone, Properties)]
    struct Props {
      #[prop_or(3)],
      countdown: usize,
    
      #[prop_or_else(Callback::noop)]
      on_click: Callback&lt;()&gt;,
    
      #[prop_or(true)]
      display: bool,
    
      #[prop_or_default]
      highlight: bool,
    
      // implicitly required
      required: MyRequiredValue,
    
      #[prop_or_default]
      opt_value: Option&lt;Value&gt;,
    
      // implicitly required
      opt_required: Option&lt;Value&gt;,
    }

In other news, we have officially kicked off our [Chinese Gitter](https://gitter.im/yewframework/%E4%B8%AD%E6%96%87) and the Chinese community has translated a lot of the new documentation site into Chinese! If you are interested in translating Yew docs into a new language, please reach out!
## [8][Having a hard time understanding what 'static means for a closure (rust book ch20).](https://www.reddit.com/r/rust/comments/fc5gbo/having_a_hard_time_understanding_what_static/)
- url: https://www.reddit.com/r/rust/comments/fc5gbo/having_a_hard_time_understanding_what_static/
---
I'm going through the rust book and finally got to chapter 20, and as I'm going through the multithreading section I noticed that the `Job` type the book defines looks very odd.

    type Job = Box&lt;dyn FnOnce() + Send + 'static&gt;;

I can understand pretty much of all of this except the `'static` bit. What does it mean for a closure/function to have a static life time?
## [9][ExpressJS vs Actix-Web: performance and running cost comparison](https://www.reddit.com/r/rust/comments/fbu5tt/expressjs_vs_actixweb_performance_and_running/)
- url: https://www.reddit.com/r/rust/comments/fbu5tt/expressjs_vs_actixweb_performance_and_running/
---
[https://medium.com/@maxsparr0w/performance-of-node-js-compared-to-actix-web-37f20810fb1a](https://medium.com/@maxsparr0w/performance-of-node-js-compared-to-actix-web-37f20810fb1a)
## [10][How to read from stdin without blocking?](https://www.reddit.com/r/rust/comments/fc71ju/how_to_read_from_stdin_without_blocking/)
- url: https://www.reddit.com/r/rust/comments/fc71ju/how_to_read_from_stdin_without_blocking/
---
I’m writing a program where I need to handle stdin, but don’t want to be blocked while waiting. To be honest I’m having trouble wrapping my head around the async stuff, and would like some help figuring that out.

Edit: the project I’m working on is a status bar (like xmobar or polybar). It should be able to display text fed in through stdin, but also update be updating a clock and other elements.
## [11][Deduplicating code only different over mutability. is there a better than using unsafe](https://www.reddit.com/r/rust/comments/fc67ad/deduplicating_code_only_different_over_mutability/)
- url: https://www.reddit.com/r/rust/comments/fc67ad/deduplicating_code_only_different_over_mutability/
---
I implemented a simple trie structure:

```rust
struct Try&lt;T, V&gt; {
	heads: Vec&lt;Node&lt;T, V&gt;&gt;,
}

/// Some(V) marks a complete item.
/// T simply implements PartialEq
struct Node&lt;T, V&gt; {
	key: T,
	value: Option&lt;V&gt;,
	children: Vec&lt;Node&lt;T, V&gt;&gt;,
}
```

This structure has a function `get_matches` that takes a prefix and returns a vector (`Vec&lt;(Vec&lt;&amp;T&gt;, &amp;E&gt;`)of all the values which match the said prefix. The vector contains the complete path and the `&amp;E`, a reference to the item.

```rust
	fn get_matches&lt;'a&gt;(&amp;'a self, path: &amp;[T]) -&gt; Vec&lt;(Vec&lt;&amp;T&gt;, &amp;E)&gt; {
		let mut start_path: Vec&lt;&amp;'a T&gt; = Vec::new();
		let mut o = Vec::new();
		let mut pivot: &amp;Vec&lt;Node&lt;T, E&gt;&gt; = &amp;self.heads;
		for key in path {
			if let Some(child) = pivot.iter().find(|c| &amp;c.key == key) {
				pivot = &amp;child.children;
				start_path.push(&amp;child.key);
			} else {
				return o;
			}
		}
		for (mut v, e) in Node::untree(pivot) {
			let mut c = start_path.clone();
			c.append(&amp;mut v);
			o.push((c, e));
		}

		o
	}

	#[allow(mutable_transmutes)]
	fn get_matches_mut&lt;'a&gt;(&amp;'a mut self, path: &amp;[T]) -&gt; Vec&lt;(Vec&lt;&amp;T&gt;, &amp;mut E)&gt; {
		use std::mem::transmute;
		self.get_matches(path)
			.into_iter()
			.map(|(a, b)| (a, unsafe { transmute(b) }))
			.collect()
	}
```

However, as you can see, I have a second function, `get_matches_mut` that would be a perfect clone of `get_matches` (except for mutability) if I didn't use `unsafe/transmute`. In theory, since `get_matches_mut` already takes a mutable reference, I ...think... the `unsafe` shouldn't cause any problem.

Is there a cleaner way of keeping this code deduplicated ?
## [12][Idiomatic approach to tree like data](https://www.reddit.com/r/rust/comments/fc8joc/idiomatic_approach_to_tree_like_data/)
- url: https://www.reddit.com/r/rust/comments/fc8joc/idiomatic_approach_to_tree_like_data/
---
Hi All,

I am building a parser and validator for a modelling language called PDDL in Rust. PDDL is used as a way to model AI Planning problems for AI Planning software to solve. The intricacies of the language aren't hugely important but here are a couple of trivial examples of what I want to represent

**Typing**
```
(:types
animal plant - object
dog cat - animal
flower tree - plant
)
```
(please forgive the trivial example)

**expressions**
```
(+ (* a b) c)
```
becomes a * b + c (PDDL uses prefix notation, please don't ask why, I didn't write the language I just work with it)

Now the former typing I can trivially collapse away the tree like behaviour by explicitly listing all the inherited types of a given object e.g.

if rover is a dog then rover is {dog, animal, object}

although this seems more expensive we typically don't have thousands of different types expressed in PDDL, and if we do the problem becomes non-trivial to solve.

The bigger issues are expressions, because we need to be able to maintain them until the solver/validator can actually determine concrete values for a, b and c.

**Bottom line**

Is there any idiomatic approach to tree like structures in rust? it seems to me that having references in structs is very expensive

Is there a rusty approach to this kind of data representation problem?

Thanks in advanced
