# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (35/2020)!](https://www.reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (35/2020)?](https://www.reddit.com/r/rust/comments/ifif3r/whats_everyone_working_on_this_week_352020/)
- url: https://www.reddit.com/r/rust/comments/ifif3r/whats_everyone_working_on_this_week_352020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-35-2020/47702?u=llogiq)!
## [3][Announcing Rapier: 2D and 3D physics engines focused on performances!](https://www.reddit.com/r/rust/comments/igkul2/announcing_rapier_2d_and_3d_physics_engines/)
- url: https://www.dimforge.com/blog/2020/08/25/announcing-the-rapier-physics-engine/
---

## [4][Announcing REVC — Rewrite MPEG-5 EVC (ETM) in Rust](https://www.reddit.com/r/rust/comments/igu8be/announcing_revc_rewrite_mpeg5_evc_etm_in_rust/)
- url: https://www.reddit.com/r/rust/comments/igu8be/announcing_revc_rewrite_mpeg5_evc_etm_in_rust/
---
[https://github.com/revcx/revc](https://github.com/revcx/revc)

# Overview

MPEG-5 Essential Video Coding (EVC) baseline profile is royalty-free. It includes only technologies that are more than 20 years old or that were submitted with a royalty-free declaration. Compared to H.264/AVC (JM19.0), MPEG-5 EVC (ETM baseline) provides about 30% BD-rate reduction with comparable computation complexity.

REVC is a Rust-based EVC (baseline) video codec implementation.

## Roadmap

* ~~0.1 Translation:~~
   * ~~Translate ETM baseline decoder from C to Rust~~
   * ~~Translate ETM baseline encoder from C to Rust~~
* 0.2 Optimization:
   * profiling and benchmarking
   * rust safe code optimization
   * assembly optimization
      * armeabi-v7a
      * arm64-v8a
      * x86
      * x86\_64
   * multi-threading optimization
* 0.3 Modernization
   * rate control
   * practical usecases: RTC, Live Streaming, VOD, etc

## Contributing

Contributors or Pull Requests are Welcome!!!
## [5][tokio-fix: Unlock yourself from the tokio walled garden](https://www.reddit.com/r/rust/comments/igxxoi/tokiofix_unlock_yourself_from_the_tokio_walled/)
- url: https://docs.rs/tokio-fix
---

## [6][Announcing Trunk — Build, bundle &amp; ship your Rust WASM application to the web.](https://www.reddit.com/r/rust/comments/igrq90/announcing_trunk_build_bundle_ship_your_rust_wasm/)
- url: https://www.reddit.com/r/rust/comments/igrq90/announcing_trunk_build_bundle_ship_your_rust_wasm/
---
I am happy to announce the very first release of Trunk. Trunk is a CLI tool, written in Rust, which provides a simple, zero-config pattern for building Rust WebAssembly applications, bundling application assets (sass, css, images &amp;amp;c) and shipping it all to the web.

Trunk is designed for creating progressive, single-page web applications, written in Rust, compiled to WebAssembly, without any JS (though today JS is still needed for loading WASM modules). Trunk follows a simple paradigm: declare an `index.html` file describing the single page of your application, then Trunk will parallelize bundling assets declared in your HTML, will build your WASM app, hash resources for cache control ... all without any extra config files.

- [release notes](https://github.com/thedodd/trunk/releases/tag/v0.1.0).
- [github repo](https://github.com/thedodd/trunk).

If you are interested in getting involved (which I hope you are), I would love for you to help out. There are lots of great features planned, and many still in the design phase. I hope you will stop by, give the issues a read, share any thoughts if you feel so inclined, and if you want to write some code, please do!

”Pack your things, we’re going on an adventure!” ~ Ferris
## [7][CraftQL - A CLI tool to manipulate GraphQL schemas and to output a graph data structure as a graphviz .dot format](https://www.reddit.com/r/rust/comments/igwg6u/craftql_a_cli_tool_to_manipulate_graphql_schemas/)
- url: https://www.reddit.com/r/rust/comments/igwg6u/craftql_a_cli_tool_to_manipulate_graphql_schemas/
---
Hello!

I made a small CLI tool (again :P) to output a graph data structure as a graphviz .dot format. It can also help you get one or multiple nodes and to find orphans. This is already usable though quite a work in progress as I have other ideas to implement...

You can play with it here  [https://github.com/yamafaktory/craftql](https://github.com/yamafaktory/craftql)

N.B.: you can use [https://edotor.net/](https://edotor.net/) to render the graph
## [8][Wordpress/Drupal like CMS-plugin-system. Which web-framework?](https://www.reddit.com/r/rust/comments/igwsy6/wordpressdrupal_like_cmspluginsystem_which/)
- url: https://www.reddit.com/r/rust/comments/igwsy6/wordpressdrupal_like_cmspluginsystem_which/
---
Hello!

First off: I am well aware that an interpreted language may be better suited to the whole "plug&amp;play"-type plugin-system I have in mind. I am interested in learning Rust though.

My vision in a nutshell: A small, extendable CMS where plugins can be registered, enabled and disabled at runtime (like e.g. in Wordpress, Drupal, ...) based on e.g. wasm plugins or even libloading, where all the different plugin-modules would need to adhere to some trait.

I've had a look at warp (love the Akka vibe!) and Rocket (very nice API) so far. (Also a sneak-peek at tide, which is btw. not mentioned on [the top list]: arewewebyet.org and a long time ago I experimented with actix[-web])

So far - just a couple of hours of research and tinkering with rusty Rust-skills - I've not managed to create a plugin system in either warp, nor rocket.

In Warp I had the problem that I couldn't combine the filters in a loop (e.g: for plugin in plugins { routes.or(plugin.route())} ) 

while in Rocket (master branch) I got a bit further but couldn't make it so that the plugins/routes can be registered when the rocket is already launched (love the jargon/humor ;))

Any advice (other than "Use an interpreted language") would be appreciated

PS: I am _not_ a total Rust noobie, I am far from being an expert though ;)
## [9][[Pre-RFC] Safe(r) Transmutation](https://www.reddit.com/r/rust/comments/igi6p0/prerfc_safer_transmutation/)
- url: https://internals.rust-lang.org/t/pre-rfc-safer-transmutation/12926
---

## [10][ANN: mles-webproxy 0.8](https://www.reddit.com/r/rust/comments/igx44n/ann_mleswebproxy_08/)
- url: https://www.reddit.com/r/rust/comments/igx44n/ann_mleswebproxy_08/
---
Ever wondered why there can't be a simple static WWW-server in Rust that would fetch TLS certificates from *Let's encrypt* out-of-the-box, renew them in timely manner, and in addition allow to use distributed messaging over WebSocket? And do this by opening just ports 80 and 443, nothing extra?

At least now there is: [https://crates.io/crates/mles-webproxy](https://crates.io/crates/mles-webproxy) (yet another, if there already is one..)

This is based on Warp 0.1. WebSocket connection is upgraded from TLS session. It has application level keepalive, so the sessions won't stay hanging there. Feedback is welcome!
## [11][Random crate episode 2: Aion](https://www.reddit.com/r/rust/comments/igye17/random_crate_episode_2_aion/)
- url: https://www.reddit.com/r/rust/comments/igye17/random_crate_episode_2_aion/
---
Hello everyone, i recently took an  initiative to make a blog post per week in which i talk about lesser  known but awesome rust crates on [crates.io](https://crates.io/), i pick the crates randomly and select those that the most interesting, i then write a little review on my blog.

Second episode: aion, it's a utility crate, inspired by rails, that allows you to write Duration and DateTime in a friendly way 

link: [https://blog.snow-blade.vercel.app/post/random-crate-ep-2](https://blog.snow-blade.vercel.app/post/random-crate-ep-2)
## [12][Uutils coreutils - GNU coreutils rewritten in safe Rust](https://www.reddit.com/r/rust/comments/igcs2o/uutils_coreutils_gnu_coreutils_rewritten_in_safe/)
- url: https://github.com/uutils/coreutils
---

