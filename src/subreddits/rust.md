# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (4/2020)!](https://www.reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here_42020/)
- url: https://www.reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here_42020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/eo6pjy/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.40]](https://www.reddit.com/r/rust/comments/ecxd62/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/ecxd62/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/dvxw6u/official_rrust_whos_hiring_thread_for_jobseekers/).

This thread will be periodically stickied to the top of /r/rust every week, for improved visibility. The thread will be refreshed and posted anew when the next version of Rust releases in six weeks.

Please adhere to the following rules when posting (adapted from /r/cpp's jobs thread).

#### Rules for individuals:

* Don't create top-level comments; those are for employers.
* Feel free to reply to top-level comments with on-topic questions.
* I will create a stickied top-level comment for individuals looking for work.
* I will create an additional top-level comment for meta discussion.

#### Rules for employers:

* To find individuals seeking work, see the replies to the stickied top-level comment; you will need to click the "more comments" link below that comment in order to see them.
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
## [3][Kondo 🧹 my first Rust crate.](https://www.reddit.com/r/rust/comments/euinys/kondo_my_first_rust_crate/)
- url: https://www.reddit.com/r/rust/comments/euinys/kondo_my_first_rust_crate/
---
[**Kondo**](https://github.com/tbillington/kondo) is a command line tool for removing things from your drive that don't bring you joy.

It identifies projects on your drive that are using a lot of space that could be saved! Examples include *target* in Cargo projects and *node_modules* in Node projects.

This is especially useful for:

* if you're about to zip up a bunch of projects to move them/back them up
* if you're running low on disk space and just want to free some up
* if you're a serial hoarder and want to keep all those things you cloned from github 1 time without the disk cost...

I've been using Rust on and off for a few years now. Inspired by tools like [ripgrep](https://github.com/BurntSushi/ripgrep) I made a [basic command line tool for helping me at work](https://github.com/tbillington/pj) a while back and was quite happy with how it turned out and was motivated me to make more!

I wanted to make this tool because I've done this manually way too many times to count. It doesn't currently delete anything because I'm paranoid that I'll `rm -rf /` somehow but I do plan to add the feature at some point, perhaps behind a prompt/flag.

You should just be able to run `cargo install kondo` and be on your way!

Sample output:

    $ kondo ~
    Scanning "C:/Users/Trent"
    1000 projects found
    Calculating savings per project
      (redacted 1000~ lines)
      385.6MB UnityTestApp (Unity) C:\Users\Trent\code\UnityTestApp
      458.7MB tokio (Cargo) C:\Users\Trent\code\tokio
        1.5GB ui-testing (Node) C:\Users\Trent\code\ui-testing
        4.0GB rust-analyzer (Cargo) C:\Users\Trent\code\rust-analyzer
    9.5GB possible savings

Any code feedback is greatly appreciated :)
## [4][Units of Measure in Rust with Refinement Types](https://www.reddit.com/r/rust/comments/eun51n/units_of_measure_in_rust_with_refinement_types/)
- url: https://yoric.github.io/post/uom.rs/
---

## [5][Words are Hard - An Essay on Communicating With Non-Programmers](https://www.reddit.com/r/rust/comments/euhe2q/words_are_hard_an_essay_on_communicating_with/)
- url: http://adventures.michaelfbryan.com/posts/words-are-hard/?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=words-are-hard
---

## [6][Music Theory library written in Rust](https://www.reddit.com/r/rust/comments/eu5fop/music_theory_library_written_in_rust/)
- url: https://www.reddit.com/r/rust/comments/eu5fop/music_theory_library_written_in_rust/
---
I've been working on a music theory library for a while, the purpose is to cover all the theoretical knowledge in the code, slowly but surely getting there. Will be adding more functionality over time.

The choice of Rust turned out to be great for this kind of library. Helped covering the edge cases with all the musical notions like Scale, Chord etc. 

Wanted to share here for anyone interested in music;

[https://github.com/ozankasikci/rust-music-theory](https://github.com/ozankasikci/rust-music-theory)
## [7][[Help] I need a help to compare the performance of the web frameworks.](https://www.reddit.com/r/rust/comments/eum1pa/help_i_need_a_help_to_compare_the_performance_of/)
- url: https://www.reddit.com/r/rust/comments/eum1pa/help_i_need_a_help_to_compare_the_performance_of/
---
[I was testing web frameworks with React](https://github.com/steadylearner/Rust-Full-Stack/tree/master/React_Rust). They are Actix, Warp, Golang, Express etc.

I use [loadtest](https://www.npmjs.com/package/loadtest) and Linux command to see how many memory each of them use.

The loadtest command used here is **$loadtest http://0.0.0.0:8000/ -t 20 -c 10 --keepalive --rps 2000**.

You can modify rps(request per second), t(timelimit), c(concurrency) part if you want to test them on your own.

(I used this after I read [the loadtest documentation](https://www.npmjs.com/package/loadtest) until "Not bad at all: 2 krps with a single core, sustained. " and you can use ctrl + f to find the part in the page)

**I let links to the details of my tests here because it was difficult to let formatted ones here**. 

You can visit [React_Rust folder](https://github.com/steadylearner/Rust-Full-Stack/tree/master/React_Rust) instead also.

**1.** [Express](https://github.com/steadylearner/Rust-Full-Stack/tree/master/React_Rust#express)

Summary) Completed requests(14669) is good but memory usage **113 MB** is not low. There were no errors but compravely high mean latency        (1012.1 ms).

**Does high mean latency is because Node is single threaded?**

I think that the high memory usage(113 MB) is from its garabage collector.

**2.** [Warp](https://github.com/steadylearner/Rust-Full-Stack/tree/master/React_Rust#warp)

Summary) There are some errors while serving the static files(React app here) concurrently. But, memory usage(4.63 MB) is pretty low.

**Do I ask the author to imporve this?**

**3.** [Actix](https://github.com/steadylearner/Rust-Full-Stack/tree/master/React_Rust#actix)

Summary) Completed requests(13300) is good without errors. The mean latency is low(32.4 ms). The memory usage(13 Mb) is also low.

**Does low latency is from its actor system?**

**4.** [Golang](https://github.com/steadylearner/Rust-Full-Stack/tree/master/React_Rust#golang)

Summary) Completed requests(27881) is good without errors. The mean latency is very low also(5.4 ms). But, the memory usage of it (231 Mb) is pretty low.

**Where this high memory usage comes from? Garbage collector or green thread for concurrency?**

With theses results, I think that **Actix** will be the best option to serve React apps(static files).

**Do you have other opnions?**

I am thinking of writing blog posts similar to ["How to use React with Rust Actix"](https://www.steadylearner.com/blog/read/How-to-use-React-with-Rust-Actix) instead with Warp.

I think that table will be better to compare the results. But, are there better ways than this?

You may give other datas to compare if you have.

p.s I am not a native speaker. You may think that I can search them on web but I want opinions from more experts here that know various languages with low level prograaming knowledge.
## [8][[OXYGENGINE] Instantiate entities from prefab assets](https://www.reddit.com/r/rust/comments/eunppk/oxygengine_instantiate_entities_from_prefab_assets/)
- url: https://www.reddit.com/r/rust/comments/eunppk/oxygengine_instantiate_entities_from_prefab_assets/
---
All you do is you write prefab asset for scene:

```yaml
template_name: instance
entities:
  - Data:
      components:
        CompositeRenderable:
          Image:
            image: logo.png
            alignment:
              x: 0.5
              y: 0.5
        CompositeTransform:
          scale:
            x: 0.15
            y: 0.15
        NonPersistent: ~
```

and then you can instantiate it with Prefab Manager like this:

```rust
// instantiate object from prefab and store its entity.
let instance = world
    .write_resource::&lt;PrefabManager&gt;()
    .instantiate_world("instance", world)
    .unwrap()[0];
// LazyUpdate::exec() runs code after all systems are done, so it's perfect to
// modify components of entities created from prefab there.
// note this `move` within closure definition - since we ue `pos` and `instance`
// objects from outside of closure scope, rust has to be informed that we want
// to move ownership of that objects to inside of closure scope.
world.read_resource::&lt;LazyUpdate&gt;().exec(move |world| {
    // fetch CompositeTransform from instance and set its position.
    // note that we can fetch multiple components at once if we pack them in
    // tuple (up to 26 components) just like that:
    // ```
    // let (mut t, s) = &lt;(CompositeTransform, Speed)&gt;::fetch(world, instance);
    // let pos = t.get_translation() + t.get_direction() * s.0;
    // t.set_translation(pos);
    // ```
    let mut transform = &lt;CompositeTransform&gt;::fetch(world, instance);
    transform.set_translation(pos);
});
```

Link to example project: https://github.com/PsichiX/Oxygengine/tree/master/demos/basic-web-game

Link to the engine project: https://github.com/PsichiX/Oxygengine

BTW. The current milestone is about making Blueprints-like visual scripting, expect something before Global Game Jam! :D
## [9][I'm learning rust and livetweeting the whole experience](https://www.reddit.com/r/rust/comments/eunh0u/im_learning_rust_and_livetweeting_the_whole/)
- url: https://twitter.com/alcuadrado/status/1221517317236019202
---

## [10][Giving an alias to super module?](https://www.reddit.com/r/rust/comments/eun0dx/giving_an_alias_to_super_module/)
- url: https://www.reddit.com/r/rust/comments/eun0dx/giving_an_alias_to_super_module/
---
When importing other modules, I can give them aliases. This not only prevents name collisions, but also allows us to type less:

    use othermodule as om;

The same applies when using a sibling module, accessible with `super` keyword:

    use super::siblingmodule as sm;

Now, is there a way to give an alias to the `super` module itself? Something like:

    use super as su;

When I try that, compiler gives me "unresolved import `super`" error.
## [11][Future of ureq http client library](https://www.reddit.com/r/rust/comments/eu6qg8/future_of_ureq_http_client_library/)
- url: https://github.com/algesten/ureq/blob/future/THOUGHTS.md
---

## [12][Help understanding array slicing / postgres traits](https://www.reddit.com/r/rust/comments/euit93/help_understanding_array_slicing_postgres_traits/)
- url: https://www.reddit.com/r/rust/comments/euit93/help_understanding_array_slicing_postgres_traits/
---
Hi all, I'm a relative Rust newbie and I've run into some errors I'm hoping you folks can explain. I'm getting compilation errors trying to call [`postgres::Client.query_one`](https://docs.rs/postgres/0.17.0/postgres/struct.Client.html#method.query_one). Here's a minimal code sample:

    let arg: [u8; 4] = [0, 1, 2, 3];
    let opt_row = client.query_opt("some query", &amp;[&amp;arg]);

This is the compilation error:

	error[E0277]: the trait bound `[u8; 4]: postgres::types::ToSql` is not satisfied
	  --&gt; src/main.rs:11:25
	   |
	11 |     client.query_opt("", &amp;[&amp;arg]);
	   |                            ^^^^ the trait `postgres::types::ToSql` is not implemented for `[u8; 4]`
	   |
	   = help: the following implementations were found:
				 &lt;&amp;'a [T] as postgres::types::ToSql&gt;
				 &lt;&amp;'a [u8] as postgres::types::ToSql&gt;
	   = note: required for the cast to the object type `dyn postgres::types::ToSql + std::marker::Sync`

The compiler error tells me that there is an impl for `&amp;[u8]`, but `&amp;` doesn't appear to be slicing here; it says I'm passing `[u8; 4]`. I tried passing `&amp;&amp;arg` as well, and got the same error. I've found I can pass `&amp;arg.to_vec()` instead and it works, but there shouldn't be a need to create a `Vec`. What am I missing here?
