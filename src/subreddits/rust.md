# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (21/2020)!](https://www.reddit.com/r/rust/comments/glvkc5/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/glvkc5/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/ghw4v6/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 339](https://www.reddit.com/r/rust/comments/gmyv8h/this_week_in_rust_339/)
- url: https://this-week-in-rust.org/blog/2020/05/19/this-week-in-rust-339/
---

## [3][New asm! syntax from RFC 2850 has been merged](https://www.reddit.com/r/rust/comments/gn6yh1/new_asm_syntax_from_rfc_2850_has_been_merged/)
- url: https://github.com/rust-lang/rust/pull/69171
---

## [4]["Contact me if you want to use this name!"](https://www.reddit.com/r/rust/comments/gmu6y1/contact_me_if_you_want_to_use_this_name/)
- url: https://www.reddit.com/r/rust/comments/gmu6y1/contact_me_if_you_want_to_use_this_name/
---
I had a project in mind so I thought I could start to do it. I was checking some package distribution platforms. I checked AUR if something like this one exists in the registry and then went to crates.io to see if the name `op` is taken or not. It was [already registered](https://crates.io/crates/op). It was saying:

 &gt; WIP. Contact me if you want to use this name!

I thought maybe it was a WIP but visited [the creator's crates.io page](https://crates.io/users/swmon) and he has, no jokes here, almost 90 packages registered saying the same thing. He even holds some important names saying "contact me if you'd like to take it", some of which are:

 - algorithm
 - any
 - bash
 - benchmark
 - bind
 - box
 - class
 - download
 - easy-oop
 - ecma
 - emulator
 - exception
 - fuck (yeah, I'm not kidding, this is taken)
 - and almost 80 more of which I have no time to write all

So I wanted to bring this to you guys' attention. Is this really ethic to do that, holding some names? I could say maybe these were WIP projects but almost all of them also state "Contact me if you want to take it.".

Has any of you guys/girls contacted this person? Does he demand anything? Are crates.io devs aware of this situation? If so, why do they not take any action? What do you think of this issue?
## [5][[ANN] Rusty Days - Virtual Rust Conference (27.07-02.08)](https://www.reddit.com/r/rust/comments/gnatik/ann_rusty_days_virtual_rust_conference_27070208/)
- url: https://www.reddit.com/r/rust/comments/gnatik/ann_rusty_days_virtual_rust_conference_27070208/
---
Hello, fellow Rustaceans!

I am one of the organizers of the **Rusty Days Virtual Conference** (https://rusty-days.org/) which is going to happen in about two months.

This event is free!

The conference starts at 27.07 (Monday) and for 5 days we want to deliver a talk at the same time each day (most probably 6 PM CEST - but we have not decided yet).

Talks will be recorded for those who will not be able to watch the live stream.

**We have opened the CfP!** If you are interested in becoming a speaker - please visit our site and submit your proposal. (Deadline for CfP will be announced soon).
We are also happy to inform you that we have been sent the first proposals already :)

For the last two days, we plan to run a hackathon (more details on that one will also be given a bit later).

So that's it. We hope this idea will find you excited and that you will be able to take part in Rusty Days.

And a small thing we would like to ask you for: could you please follow us and like/RT us on social media? It will definitely help with widening our ranges (which is pretty important in regards to event organization).

More info: https://rusty-days.org/

Twitter: https://twitter.com/rdconf

FB: https://www.facebook.com/Rusty-Days-113835850333252
## [6][Rocket can be compiled on stable Rust 1.45, last blocker has been solved](https://www.reddit.com/r/rust/comments/gmkpsg/rocket_can_be_compiled_on_stable_rust_145_last/)
- url: https://github.com/SergioBenitez/Rocket/issues/19#issuecomment-630650328
---

## [7][tokio, tonic and a thread pool...](https://www.reddit.com/r/rust/comments/gn9pyp/tokio_tonic_and_a_thread_pool/)
- url: https://www.reddit.com/r/rust/comments/gn9pyp/tokio_tonic_and_a_thread_pool/
---
Hi, I'm writing a small application that has an rpc server (using tonic) but also needs to listen for some fs changes.

My initial thought in my async main was to spawn a thread pool for my filesystem watchers - but I wanted to see if there was another way to leverage the tokio reactor for this.

The idea is my main(async) will create a channel and hand the Senders to each of the Watchers I create - though I'm a bit lost on how to go about this async.

Any help / suggestions appreciated!
## [8][🎉 🎉 The code is compiling, and the tests pass! 🎉 🎉](https://www.reddit.com/r/rust/comments/gmx7u8/the_code_is_compiling_and_the_tests_pass/)
- url: https://www.reddit.com/r/rust/comments/gmx7u8/the_code_is_compiling_and_the_tests_pass/
---
I'm currently learning to code rust with rustlings. I'm very pleased with with the progress I made up to now:

https://preview.redd.it/7h429cjocsz41.png?width=621&amp;format=png&amp;auto=webp&amp;s=11cd5a381325cf3d9caf78fe1c66624323e60e8b
## [9][What's the project of the cross-platform Rust editor using native UI components?](https://www.reddit.com/r/rust/comments/gn7u1d/whats_the_project_of_the_crossplatform_rust/)
- url: https://www.reddit.com/r/rust/comments/gn7u1d/whats_the_project_of_the_crossplatform_rust/
---
I remember watching a very nice talk a while ago about a project to implement an IDE or editor in Rust which would work across platform while using native UI components and text rendering, but I can't find the talk and I don't remember the name of the project.  Does anyone by chance know what I'm referring to?  Is the project already open-sourced?
## [10][Confusion around borrow checker analysis of loops](https://www.reddit.com/r/rust/comments/gn901c/confusion_around_borrow_checker_analysis_of_loops/)
- url: https://www.reddit.com/r/rust/comments/gn901c/confusion_around_borrow_checker_analysis_of_loops/
---
I've run into a case where borrow checker seems to think that returning from a function mid loop can result in a mutable borrow being held over into another iteration of a loop.

I've put together a [simplified example](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=4ee7019fa83b7206e48941b6fb1e43aa). The curious part is in the last`fn`, the rest is just fluff to justify the signature.

How I mentally check borrows in loops is to figure out what borrows survive over a loop iteration after the drops at the end of the scope occur and make sure that no incompatible borrows occur. In this case we can see that the candidate reference to be returned gets dropped before another iteration can occur and that there is no sub-borrow of `self` crossing iterations.
The compiler on the other hand seems to be working backwards from the lifetime of the return type and using that to assert that *any* borrow created from`self` that could be returned should be treated as though it *definitely is* living as long as a returned type and therefore self is permanently borrowed each iteration.

Is that whats happening? Does anyone have a workarounds to suggest? All I have are the ugly options of Arcs or Clones.

Edit: it looks like the consensus is that this is behaviour as intended, unfortunately. Is there any chance of creating UB if the loop returns a pointer which is then, after the loop, converted to a reference with a lifetime bound by self, [like so](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=9192689c8e23537575521b049bac963f)?
## [11][Refinery 0.3 released](https://www.reddit.com/r/rust/comments/gn0kaj/refinery_03_released/)
- url: https://www.reddit.com/r/rust/comments/gn0kaj/refinery_03_released/
---
Hi, I have just released [refinery 0.3](https://github.com/rust-db/refinery), a powerful SQL migration toolkit for Rust. 
New features include:

- Update `Runner.run` and `Runner.run_async` return signature, it's now `Result&lt;Report, Error&gt;` where `Report` contains applied Migration's
-  Add `Runner.get_applied_migrations_async` method
- Add `Runner.get_applied_migrations` method
- Add `Runner.get_last_applied_migration_async` method
- Add `Runner.get_last_applied_migration` method
- Add allow migrations to run up until a `Target` version,
- Impl `Migrate` trait for `Config` so that projects using drivers that aren't yet supported such as `SQLx` can still migrate from a database config
## [12][[pyo3] trouble with subclassing from Python](https://www.reddit.com/r/rust/comments/gn97kn/pyo3_trouble_with_subclassing_from_python/)
- url: https://www.reddit.com/r/rust/comments/gn97kn/pyo3_trouble_with_subclassing_from_python/
---
I'm using pyo3 and I'm having issues getting subclassing to work. Specifically a Python class that inherits from a Rust class.

When I try to construct the Python class I just get an instance of the Rust class.

Minimal example:

```
#[pyclass(subclass,dict)]
struct Rust {
}

#[pymethods]
impl Rust {
    #[new]
    fn new() -&gt; Rust {
        Rust {}
    }
}

#[pymodule]
fn test(_py: Python, m: &amp;PyModule) -&gt; PyResult&lt;()&gt; {
    m.add_class::&lt;Rust&gt;()?;
    Ok(())
}
```

Python code:

```
import test

class Python(test.Rust):
    pass

p = Python()
print('p is a', type(p))
```

It's printing `p is a &lt;class 'Rust'&gt;` whereas I would expect `p is a &lt;class 'Python'&gt;`.
`Python.__mro__` returns `[&lt;class Python&gt;, &lt;class Rust&gt;, &lt;class 'object'&gt;]` as I would expect.

Looking into the code I can see that the `__new__` method generated by pyo3 ignores the "cls" argument and always returns an instance of the Rust class. Is there a way to make this work?
