# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (39/2020)!](https://www.reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ismh8n/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (39/2020)?](https://www.reddit.com/r/rust/comments/iwxjdh/whats_everyone_working_on_this_week_392020/)
- url: https://www.reddit.com/r/rust/comments/iwxjdh/whats_everyone_working_on_this_week_392020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-39-2020/49088?u=llogiq)!
## [3][Neovim and Rust - An effective development experience](https://www.reddit.com/r/rust/comments/iy15dw/neovim_and_rust_an_effective_development/)
- url: https://sharksforarms.dev/posts/neovim-rust/
---

## [4][OMG WTF RS - Resources to help you get started with Rust](https://www.reddit.com/r/rust/comments/iy81jn/omg_wtf_rs_resources_to_help_you_get_started_with/)
- url: https://ferrous-systems.com/blog/omg-wtf-rs-resources-to-help-you-get-started-with-rust/
---

## [5][iou version 0.3](https://www.reddit.com/r/rust/comments/ixvltf/iou_version_03/)
- url: https://without.boats/blog/iou-0-3/
---

## [6][PyO3 based modelling &amp; simulation framework (~300 loc, 4X speed of Python implementation - let's make it faster!)](https://www.reddit.com/r/rust/comments/iy6j8p/pyo3_based_modelling_simulation_framework_300_loc/)
- url: https://www.reddit.com/r/rust/comments/iy6j8p/pyo3_based_modelling_simulation_framework_300_loc/
---
Hey everyone! I'm quite new to Rust. I've started converting a "dynamical systems" modelling and simulation framework  from Python to Rust using PyO3 to create a native Python module.

What's  amazing is that it's only 300 loc so far, and already implements most of  the features seen in the existing Python module, at 4X the speed (which  is important for simulations). The reason it isn't faster, in my Rust  newbie mind, is because it relies on making calls back to Python from Rust (this is a requirement, to utilize the Python stack). It's a simple code-base, and I'd love a code review if anyone is up for it,  specifically about how I can optimize it for speed and memory  efficiency: [https://github.com/BenSchZA/radCAD](https://github.com/BenSchZA/radCAD)

I think it's a great use-case for Rust, and although the code-base is simple, it has interesting use-cases and opportunities for benchmarking - so I'm keen to see just how fast we can get these simulations to run using PyO3.
## [7][Async Iteration Semantics](https://www.reddit.com/r/rust/comments/ixwazy/async_iteration_semantics/)
- url: https://blog.yoshuawuyts.com/async-iteration/
---

## [8][Nushell Survey 2020 (for users and non-users)](https://www.reddit.com/r/rust/comments/ixwyhv/nushell_survey_2020_for_users_and_nonusers/)
- url: https://docs.google.com/forms/d/e/1FAIpQLScEFzDh7j3jfAuVMBCQtQE-qfKAhugLCiUaaGL583QtGwz5fw/viewform
---

## [9][I created a Window Hello style authentication for Linux like Howdy, but with rust.](https://www.reddit.com/r/rust/comments/iy16di/i_created_a_window_hello_style_authentication_for/)
- url: https://github.com/saanuregh/hola
---

## [10][Porting EBU R128 audio loudness analysis from C to Rust](https://www.reddit.com/r/rust/comments/ixn7ju/porting_ebu_r128_audio_loudness_analysis_from_c/)
- url: https://coaxion.net/blog/2020/09/porting-ebu-r128-audio-loudness-analysis-from-c-to-rust/
---

## [11][Trunk 0.5.0 | Proxy System](https://www.reddit.com/r/rust/comments/iya08b/trunk_050_proxy_system/)
- url: https://www.reddit.com/r/rust/comments/iya08b/trunk_050_proxy_system/
---
Build, bundle &amp; ship your Rust WASM application to the web.

- [Github repo](https://github.com/thedodd/trunk)
- [Release notes](https://github.com/thedodd/trunk/releases/tag/v0.5.0)

Trunk now ships with a built-in proxy which can be enabled when running `trunk serve`. There are two ways to configure the proxy, each discussed below. All Trunk proxies will transparently pass along the request body, headers, and query parameters to the proxy backend.

### proxy cli flags
The `trunk serve` command accepts two proxy related flags.

`--proxy-backend` specifies the URL of the backend server to which requests should be proxied. The URI segment of the given URL will be used as the path on the Trunk server to handle proxy requests. E.G., `trunk serve --proxy-backend=http://localhost:9000/api/` will proxy any requests received on the path `/api/` to the server listening at `http://localhost:9000/api/`. Further path segments or query parameters will be seamlessly passed along.

`--proxy-rewrite` specifies an alternative URI on which the Trunk server is to listen for proxy requests. Any requests received on the given URI will be rewritten to match the URI of the proxy backend, effectively stripping the rewrite prefix. E.G., `trunk serve --proxy-backend=http://localhost:9000/ --proxy-rewrite=/api/` will proxy any requests received on `/api/` over to `http://localhost:9000/` with the `/api/` prefix stripped from the request, while everything following the `/api/` prefix will be left unchanged.

### config file
The `Trunk.toml` config file accepts multiple `[[proxy]]` sections, which allows for multiple proxies to be configured. Each section requires at least the `backend` field, and optionally accepts the `rewrite` field, both corresponding to the `--proxy-*` CLI flags discussed above.

As it is with other Trunk config, a proxy declared via CLI will take final precedence and will cause any config file proxies to be ignored, even if there are multiple proxies declared in the config file.

The following is a snippet from the `Trunk.toml` file in this repo:

```toml
[[proxy]]
rewrite = "/api/v1/"
backend = "http://localhost:9000/"
```
## [12][Are there any gotchas to converting '&amp;T' into '&amp;[T; 1]'?](https://www.reddit.com/r/rust/comments/ixqzkf/are_there_any_gotchas_to_converting_t_into_t_1/)
- url: https://www.reddit.com/r/rust/comments/ixqzkf/are_there_any_gotchas_to_converting_t_into_t_1/
---
Hi rustaceans,

I'm working on a project where I have a function that has to return a slice from a reference of the same type.  The type is `Copy`, but I can't return it using `&amp;[*x]` because doing so would store the slice on the stack frame of the function, which gets invalidated when the function returns.  So I'm using `std::slice::from_raw_parts` to convert the reference into an immutable slice of length `1` with an `unsafe` block.

So, my question is: are there any gotchas that I'm letting myself in on?  I can't see how a `&amp;'a [T; 1]` is any different from an `&amp;'a T`, but there might be something I've overlooked.
This code currently only happens inside the test suite so I'm OK with a bit of potential UB there (since UB is highly likely to just cause my tests to break which fails safe), but in the future I'm going to have **a lot** of these conversions going on and I'm making a text editor with some very spicy data structures so memory corruption would be really bad.  As a result, I'd like to avoid using `unsafe` unless forced.  Edit: my tests do seem to be working fine...

Thanks!
