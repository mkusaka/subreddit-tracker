# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (44/2020)!](https://www.reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/jdwuis/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 362](https://www.reddit.com/r/rust/comments/jk35ha/this_week_in_rust_362/)
- url: https://this-week-in-rust.org/blog/2020/10/28/this-week-in-rust-362/
---

## [3][Announcing Qovery Engine - A Rust lib to deploy microservices apps on any Cloud provider](https://www.reddit.com/r/rust/comments/jlyv25/announcing_qovery_engine_a_rust_lib_to_deploy/)
- url: https://www.reddit.com/r/rust/comments/jlyv25/announcing_qovery_engine_a_rust_lib_to_deploy/
---
Hi Rustaceans, after months of hard work with my team of 6. I am glad to announce that we have open-sourced our deployment engine.

[https://github.com/Qovery/engine](https://github.com/Qovery/engine)

Qovery engine is still under development, but more than 600 developers and dozens of successful companies use our Engine for 11 months through [Qovery](https://www.qovery.com).

## Features

* **Zero infrastructure management:** Qovery Engine initialized, configure, and manage your Cloud account for you.
* **Multi-Cloud:** Qovery Engine is built to work on AWS, GCP, Azure, and any Cloud provider.
* **On top of Kubernetes:** Qovery Engine takes advantage of the power of Kubernetes at a higher level of abstraction.
* **Terraform and Helm:** Qovery Engine uses Terraform and Helm files to manage the infrastructure and app deployment.
* **Powerful CLI:** Use the provided Qovery Engine CLI to deploy your app on your Cloud account seamlessly.

&amp;#x200B;

[compose your flow with Qovery engine](https://preview.redd.it/7d5msvtk6lw51.png?width=1722&amp;format=png&amp;auto=webp&amp;s=594771af8b1dc9496cfd303bafb356eb497b9031)

Qovery engine supports many different plugins to compose your own deployment flow.

As Qovery Engine is a plugin system, we are looking for contributors interested in helping us making Rust successful in the Cloud industry.
## [4][Stabilizing arc-swap](https://www.reddit.com/r/rust/comments/jlztu7/stabilizing_arcswap/)
- url: https://vorner.github.io/2020/10/31/stabilizing-arc-swap.html
---

## [5][Fixing bootstrap of rustc using cg_clif](https://www.reddit.com/r/rust/comments/jm0rte/fixing_bootstrap_of_rustc_using_cg_clif/)
- url: https://bjorn3.github.io/2020/11/01/fixing-rustc-bootstrap-with-cg_clif.html
---

## [6][Compile-Time Reflection in Rust](https://www.reddit.com/r/rust/comments/jm1h2x/compiletime_reflection_in_rust/)
- url: https://www.mn.uio.no/ifi/english/research/groups/psy/completedmasters/2020/gaarde/masterthesis-gaarde.pdf
---

## [7][Show r/rust: hstr-rs - A bash history suggest box](https://www.reddit.com/r/rust/comments/jlygxl/show_rrust_hstrrs_a_bash_history_suggest_box/)
- url: https://www.reddit.com/r/rust/comments/jlygxl/show_rrust_hstrrs_a_bash_history_suggest_box/
---
**hstr** is bash history suggest box. Like hstr, but with pages.  


It lets you quickly search, navigate, manage, and execute commands from your bash history.  


The issue with regex-mode is now fixed.  


[Check it out](https://github.com/adder46/hstr-rs) and let me know what you think!  


&amp;#x200B;

https://i.redd.it/fb0l7zl7ykw51.gif
## [8][rkvm - A tool to share keyboard and mouse over the network on Linux machines](https://www.reddit.com/r/rust/comments/jlhga1/rkvm_a_tool_to_share_keyboard_and_mouse_over_the/)
- url: https://github.com/htrefil/rkvm
---

## [9][C person question for RUST person.](https://www.reddit.com/r/rust/comments/jm0p1j/c_person_question_for_rust_person/)
- url: https://www.reddit.com/r/rust/comments/jm0p1j/c_person_question_for_rust_person/
---
In C often heap/memory needs to be looked at a different way.  also you have CPU states (ie: device drivers).  C makes all this pretty easy (but dangerous).  For instance you might need to optimize a string in such a way that you know that the end result will be smaller than the incoming string.  even better than that you know you need 64 bytes of stack and you can reuse the heap and return a different type.  with C your given this control and creativity... at a risk given.  so my question: does RUST allow semantics to address "low level" C semantics.  Could you write a kernel in RUST.  Could you natively talk to PCI physical addresses using RUST?  or is this still C land?
## [10][Working Group for better Rust support in ROS2](https://www.reddit.com/r/rust/comments/jlnp00/working_group_for_better_rust_support_in_ros2/)
- url: https://discourse.ros.org/t/proposal-rust-working-group-for-ros/16993
---

## [11][Projects in rust](https://www.reddit.com/r/rust/comments/jlp3kd/projects_in_rust/)
- url: https://www.reddit.com/r/rust/comments/jlp3kd/projects_in_rust/
---
Can someone guide me some beginner-medium project to learn rust in software development field?

Thanks :)
## [12][wasm-script - Build your on WebAssembly compiler in Rust and put it on the web!](https://www.reddit.com/r/rust/comments/jloq9n/wasmscript_build_your_on_webassembly_compiler_in/)
- url: https://github.com/richardanaya/wasm-script
---

