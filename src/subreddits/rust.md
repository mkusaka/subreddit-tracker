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
## [2][This Week in Rust 328](https://www.reddit.com/r/rust/comments/fdzspa/this_week_in_rust_328/)
- url: https://this-week-in-rust.org/blog/2020/03/03/this-week-in-rust-328/
---

## [3][Considering Rust [talk]](https://www.reddit.com/r/rust/comments/fenspv/considering_rust_talk/)
- url: https://youtu.be/DnT-LUQgc7s
---

## [4][Running Actix integration tests with Github actions](https://www.reddit.com/r/rust/comments/fettdq/running_actix_integration_tests_with_github/)
- url: https://www.youtube.com/watch?v=TKCJl6YgsH8
---

## [5][looking at gdb how can I know whether its heap or stack allocated memory ?](https://www.reddit.com/r/rust/comments/fes6pp/looking_at_gdb_how_can_i_know_whether_its_heap_or/)
- url: https://i.redd.it/4rmmxnbfj7l41.png
---

## [6][Demo of a new GUI + 2D drawing crate](https://www.reddit.com/r/rust/comments/fejx5a/demo_of_a_new_gui_2d_drawing_crate/)
- url: https://www.reddit.com/r/rust/comments/fejx5a/demo_of_a_new_gui_2d_drawing_crate/
---
"Oh no, not another one."

I've made a [standalone demo](https://github.com/dabreegster/abstreet/blob/master/ezgui/examples/demo.rs) of the [GUI library](https://github.com/dabreegster/abstreet/tree/master/ezgui) that I've built for [A/B Street](https://abstreet.org). I think the library has enough properties that set it apart from existing crates that it might be useful to the Rust community. The API needs lots of cleanup before a 0.1 release, but I'll only put in the effort if some folks have use cases for it.

Some I'm here to gauge interest. Tell me what cool thing you'd build with this and why the other existing crates don't work for you.

Some highlights:

*  Runs on native or web, modulo text support a few months away
*  Fully vectorized text (usvg -&gt; lyon -&gt; triangulated polygons)
*  [All features](https://github.com/dabreegster/abstreet/tree/master/ezgui)

Thanks!
-Dustin
## [7][Show r/rust: Belay - Run your CI checks locally to git push with confidence](https://www.reddit.com/r/rust/comments/feq2x8/show_rrust_belay_run_your_ci_checks_locally_to/)
- url: https://github.com/JoshMcguigan/belay
---

## [8][Naming a module 'meta' results in weird behavior. I got bitten by this recently](https://www.reddit.com/r/rust/comments/fekol0/naming_a_module_meta_results_in_weird_behavior_i/)
- url: https://internals.rust-lang.org/t/is-the-module-name-meta-forbidden/9587
---

## [9][Simple Rust program using 100% CPU.](https://www.reddit.com/r/rust/comments/feu3ru/simple_rust_program_using_100_cpu/)
- url: https://www.reddit.com/r/rust/comments/feu3ru/simple_rust_program_using_100_cpu/
---
Hello, I wrote simple program which downloads multiple files at once using `threadpool` and `reqwest` and also calls API. When I run it on my main PC i get 0% CPU usage and everything works correctly. However I have a server I want to run it on, and when I run it there it uses 100% CPU constantly. I tried running it inside a VM with 1 CPU core (to simulate low performance enviromnent) and the CPU usage was 2-3%. I don't know what I am doing wrong.

1. I tried using async first, but switched to blocking `threadpool, reqwest::blocking`. No change
2. Tried building with `target-cpu native` and `lto`. No change
3. Tried building directly on the server, also no change.

Part of my download code: [https://pastebin.com/RvE0jWqK](https://pastebin.com/RvE0jWqK)

I am new to Rust, so sorry for my begginer mistakes. Thank you.
## [10][Not Another Lifetime Annotations Post](https://www.reddit.com/r/rust/comments/feb31o/not_another_lifetime_annotations_post/)
- url: https://www.reddit.com/r/rust/comments/feb31o/not_another_lifetime_annotations_post/
---
Yeah, it is. But I've spend a few days on this and I'm starting to really worry about my brain, because I'm just getting nowhere.

To be clear, it's not *lifetimes* that are confusing. They aren't. Anyone who's spent any meaningful time writing C/C++ code understands the inherent problem that Rust solves re: dangling pointers and how strict lifetimes/ownership/borrowing all play a role in the solution.

But...lifetime annotations, I simply and fully cannot wrap my head around.

Among other things (reading the book, a few articles, and reading some of Spinoza's ethics because this shit is just about as cryptic as what he's got in there so maybe he mentioned lifetime annotations), I've watched [this video](https://www.youtube.com/watch?v=QoEX-Vu-R6k), and the presenter gave me hope early on by promising to dive into annotations specifically, not just lifetimes. But...then, just, nothing. Nothing clicks, not even a nudge in the direction of a click. Here are some of my moments of pure confusion:

* At one point, he spams an `'a` lifetime parameter across a function signature. But then it compiles, and he says "these are wrong". I have no idea what criteria for correctness he's even using at this point. What I'm understanding from this is that all of the responsibility for correctness here falls to the programmer, who can fairly easily "get it wrong", but with consequences that literally no one specifies anywhere that I've seen.
* He goes on to 'correct' the lifetime annotations...but he does this with **explicit knowledge of the calling context**. He says, "hey, look at this particular call - one of the parameters here has an entirely different lifetime than the other!" and then alters the lifetimes annotations *in the function signature* to reflect that particular call's scope context. How is this possibly a thing? There's no way I can account for every possible calling context as a means of deriving the "correct" annotations, and as soon as I do that, I might have created an invalid annotation signature with respect to some other calling context.
* He then says that we're essentially "mapping inputs to outputs" - alright, that's moving in the right direction, because the problem is now framed as one of relations between parameters and outputs, not of unknowable patterns of use. But he doesn't explain how they relate to each other, and it just seems completely random to me if you ignore the outer scope.

The main source I've been using, though, is The Book. Here are a couple moments from the annotations section where I went actually wait what:


&gt;We also don’t know the concrete lifetimes of the references that will be passed in, so we can’t look at the scopes...to determine whether the reference we return will always be valid.

Ok, so that sort of contradicts what the guy in the video was saying, if they mean this to be a general rule. But then:

&gt; For example, let’s say we have a function with the parameter `first` that is a reference to an `i32` with lifetime `'a`. The function also has another parameter named `second` that is another reference to an `i32` that also has the lifetime `'a`. The lifetime annotations indicate that the references `first` and `second` must both live as long as that generic lifetime.

Now, suddenly, it *is* the programmer's responsibility yet again to understand the "outer scope". I just don't understand what business it is of the function signature what the lifetimes are of its inputs - if they live longer than the function (they should inherently do so, right?) - why does it have to have an opinion? What is this informing as far as memory safety?

&gt; The constraint we want to express in this signature is that all the references in the parameters and the return value must have the same lifetime.

This is now dictatorial against the outer scope in a way that makes no sense to me. Again, why does the function signature care about the lifetimes of its reference parameters? If we're trying to resolve confusion around a returned reference, I'm still unclear on what the responsibility of the function signature is: if the only legal thing to do is return a reference that lives longer than the function scope, then that's all that either I or the compiler could ever guarantee, and it seems like all patterns in the examples reduce to "the shortest of the input lifetimes is the longest lifetime we can guarantee the output to be", which is a hard-and-fast rule that doesn't require programmer intervention. At best we could contradict the rule if we knew the function's return value related to only one of the inputs, but...that also seems like something the compiler could infer, because that guarantee probably means there's no ambiguity. Anything beyond seems to me to be asking the programmer, again, to reach out into outer scope to contrive to find a better suggestion than that for the compiler to run with. Which...we could get wrong, again, but I haven't seen the consequences of that described anywhere.

&gt; The lifetimes might be different each time the function is called. This is why we need to annotate the lifetimes manually.

Well, yeah, Rust, that is exactly the problem that *I* have. We have a lot in common, I guess. I'm currently mulling the idea of what happens when you have some sort of struct-implemented function that takes in references that the function intends to take some number of immutable secondary references to (are these references of references? Presumably ownership rules are the same with actual references?) and distribute them to bits of internal state, but I'm seeing this problem just explode in complexity so quickly that I'm gonna not do that anymore.

That's functions, I guess, and I haven't even gotten to how confused I am about annotations in structs (why on earth would the struct care about anything other than "these references outlive me"??) I'm just trying to get a handle on one ask: how the hell do I know what the 'correct' annotations are? If they're call-context derived, I'm of the opinion that the language is simply adding too much cognitive load to the programmer to justify any attention at all, or at least that aspect of the language is and it should be avoided at all costs. I cannot track the full scope context of every possible calling point all the time forever. How do library authors even exist if that's the case?

Of course it isn't the case - people use the language, write libraries and work with lifetime annotations perfectly fine, so I'm just missing something very fundamental here. If I sound a bit frustrated, that's because I am. I've written a few thousand lines of code for a personal project and have used 0 lifetime annotations, partially because I feel like most of the potential use-cases I've encountered present much better solutions in the form of transferring ownership, but mostly because I don't get it. And I just hate the feeling that such a central facet of the language I'm using is a mystery to me - it just gives me no creative confidence, and that hurts productivity.

***

*edit for positivity: I am genuinely enjoying learning about Rust and using it in practice. I'm just very sensitive to my own ignorance and confusion.

*edit 2: just woke up and am reading through comments, thanks to all for helping me out. I think there are a couple standout concepts I want to highlight as really doing work against my confusion:

- Rust expects your function signature to completely and unambiguously describe the contract, lifetimes, types, etc., without relying on inference, because that allows for unmarked API changes - but it does validate your function body against the signature when actually compiling the function.

- 'Getting it wrong' means that your function might be overly or unusably constrained. The job of the programmer is to consider what's happening in the body of the function (which inputs are ACTUALLY related to the output in a way that I can provide the compiler with a less constrained guarantee?) to optimize those constraints for more general use.

I feel quite a bit better about the function-signature side of things. I'm going to go back and try to find some of the places I actively avoided using intermediate reference-holding structs to see if I can figure that out.
## [11][Work in progress java virtual machine written in rust](https://www.reddit.com/r/rust/comments/feeqpl/work_in_progress_java_virtual_machine_written_in/)
- url: https://github.com/douchuan/jvm
---

## [12][Introducing Yew Form, a model binder for HTML forms with Yew](https://www.reddit.com/r/rust/comments/fedgwj/introducing_yew_form_a_model_binder_for_html/)
- url: https://www.reddit.com/r/rust/comments/fedgwj/introducing_yew_form_a_model_binder_for_html/
---
I love Yew, but if you're like me, you probably find that Yew is not great at binding  HTML form control to Rust `struct`.  Since my current project requires me to create rich form that require complex validation, I decided to complement Yew with  [yew\_form](https://github.com/jfbilodeau/yew_form), a set of mildly opinionated component to automatically map and validate a model to a HTML form.

Presently, it supports:

* Two way binding between Rust `struct` and HTML form control
* Validation using Keats Validator ([https://github.com/Keats/validator](https://github.com/Keats/validator))
* Support structs via `#[derive(Model)]`
* Support for nested structs
* Support for any scalar datatypes that implements `ToString` and `From&lt;Str&gt;`
* Support for any custom types by implementing `yew_form::FormValue`
* Build in `Field` and `CheckBox` component for binding

There's an online demo at [http://chronogears.com/yew-form/](http://chronogears.com/yew-form/) and the source for the demo is available [here](https://github.com/jfbilodeau/yew_form/tree/master/examples/form)

There's still work to be done, but it's already quite usable at my end. I hope this can be of help to others Yew users out there!

J-F
