# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (6/2020)!](https://www.reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here_62020/)
- url: https://www.reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here_62020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.41]](https://www.reddit.com/r/rust/comments/eyw94s/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/eyw94s/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/ecxd62/official_rrust_whos_hiring_thread_for_jobseekers/).

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
## [3][[FOSDEM] sled and rio: modern database engineering with io_uring](https://www.reddit.com/r/rust/comments/f06y7m/fosdem_sled_and_rio_modern_database_engineering/)
- url: https://fosdem.org/2020/schedule/event/rust_techniques_sled/
---

## [4][Announcing Cooked Wakers!](https://www.reddit.com/r/rust/comments/f05qiy/announcing_cooked_wakers/)
- url: https://www.reddit.com/r/rust/comments/f05qiy/announcing_cooked_wakers/
---
After finding that consuming `RawWaker`s and `RawWakerVTable`s was bad for my health, I created [cooked-waker](https://docs.rs/cooked-waker). `cooked-waker` is a safe library with *minimal* abstraction cost that allows you to create a `Waker` out of purely safe rust traits. It aims to be exactly what you "would have" written if you were implementing a `Waker` by hand; in particular, it aims to minimize boxing and other kinds of indirection whenever possible.

It works be providing the `Wake` and `WakeRef` traits, which correspond to the `wake` and `wake_by_ref` methods on `std::task::Waker`. Once you have a *concrete* type that implements these traits, you can derive `IntoWaker` on it, which will generate all of the boilerplate code necessary to create a working `Waker` struct.

Under the hood it leverages [`stowaway`](https://docs.rs/stowaway/1.1.1/stowaway/)¹ to pack your waker struct into the pointer in an `std::task::Waker`. Stowaway prevents extra boxing by packing your struct directly into the bytes of a pointer if it will fit; the most common case is that your struct simply contains a single pointer, like an `Arc&lt;Handle&gt;`.

¹ I wrote this crate as well

# Example:

    use cooked_waker::{Wake, WakeRef, IntoWaker};
    use std::sync::atomic::{AtomicUsize, Ordering};
    use std::sync::Arc;
    use std::task::Waker;

    /// A simple struct that counts the number of times it is awoken. Can't
    /// be awoken by value (because that would discard the counter), so we
    /// must instead wrap it in an Arc (see CounterHandle)
    #[derive(Debug, Default)]
    struct Counter {
        // We use atomic usize because we need Send + Sync and also interior
        // mutability
        count: AtomicUsize,
    }

    impl Counter {
        fn get(&amp;self) -&gt; usize {
            self.count.load(Ordering::SeqCst)
        }
    }

    impl WakeRef for Counter {
        fn wake_by_ref(&amp;self) {
            let _prev = self.count.fetch_add(1, Ordering::SeqCst);
        }
    }

    /// A shared handle to a Counter.
    ///
    /// We can derive Wake and WakeRef because the inner field implements
    /// them, and we can derive IntoWaker because this is a concrete type
    /// with Wake + Clone + Send + Sync. Note that *any* concrete type can have
    /// IntoWaker implemented for it; it doesn't have to be "pointer-sized"
    #[derive(Debug, Clone, Default, WakeRef, Wake, IntoWaker)]
    struct CounterHandle {
        counter: Arc&lt;Counter&gt;,
    }

    impl CounterHandle {
        fn get(&amp;self) -&gt; usize {
            self.counter.get()
        }
    }

    let counter = CounterHandle::default();

    // Create an std::task::Waker
    let waker: Waker = counter.clone().into_waker();

    waker.wake_by_ref();
    waker.wake_by_ref();

    let waker2 = waker.clone();
    waker2.wake_by_ref();

    // This calls Counter::wake_by_ref because the Arc doesn't have exclusive
    // ownership of the underlying Counter
    waker2.wake();

    assert_eq!(counter.get(), 4);
## [5][Announcing the Cleanup Crew ICE-breaker group | Inside Rust Blog](https://www.reddit.com/r/rust/comments/ezxmu4/announcing_the_cleanup_crew_icebreaker_group/)
- url: https://blog.rust-lang.org/inside-rust/2020/02/06/Cleanup-Crew-ICE-breakers.html
---

## [6][Is actix still the choice to make?](https://www.reddit.com/r/rust/comments/f07n4b/is_actix_still_the_choice_to_make/)
- url: https://www.reddit.com/r/rust/comments/f07n4b/is_actix_still_the_choice_to_make/
---
I'm starting a new project, I wanted to use rust and I was wondering if, after what happened recently, actix-* was still a good option?
## [7][Wrapper types for beginners](https://www.reddit.com/r/rust/comments/f00vtt/wrapper_types_for_beginners/)
- url: https://manishearth.github.io/blog/2015/05/27/wrapper-types-in-rust-choosing-your-guarantees/
---

## [8][Idiomatic rust](https://www.reddit.com/r/rust/comments/f08231/idiomatic_rust/)
- url: https://www.reddit.com/r/rust/comments/f08231/idiomatic_rust/
---
Hi there.

In order to teach myself some *Rust* using examples, I imagined it'd be nice to do what I did with *Haskell* : go as far as possible in the [99 questions](https://wiki.haskell.org/H-99:_Ninety-Nine_Haskell_Problems).

The [seventh question](https://wiki.haskell.org/99_questions/1_to_10#Problem_7) is about flattening a nested list structure :

```rust
pub enum NestedListEnum&lt;T&gt; {
    Elem(T),
    List(Vec&lt;NestedListEnum&lt;T&gt;&gt;),
}
```

So, my `enum` might not conform *Rust*'s best practices but my question is more about the `flatten` implementation. Here are two :

```rust
pub fn flatten_0&lt;T: Copy&gt;(nested_list: &amp;[NestedListEnum&lt;T&gt;]) -&gt; Vec&lt;T&gt; {
    nested_list.split_first().map_or(vec![], |(head, tail)| {
        [
            match head {
                NestedListEnum::Elem(e) =&gt; vec![*e],
                NestedListEnum::List(l) =&gt; flatten(l),
            },
            flatten(tail),
        ]
        .concat()
    })
```

and

```rust
pub fn flatten_1&lt;T: Copy&gt;(nested_list: &amp;[NestedListEnum&lt;T&gt;]) -&gt; Vec&lt;T&gt; {
    nested_list.split_first().map_or(vec![], |(head, tail)| {
        let mut output = match head {
            NestedListEnum::Elem(e) =&gt; vec![*e],
            NestedListEnum::List(l) =&gt; flatten(l),
        };
        output.extend(flatten(tail));
        output
    })
}
```

So, my question is, given *Rust*'s internals and or best practice, which one is best?

Also, I'm not against other feedback if you're at it.

Many thanks in advance!

EDIT : u/po8 and u/wishthane pointed out that `Clone` might be better than `Copy` so the code could well be as follow :

```rust
pub fn flatten&lt;T: Clone&gt;(nested_list: &amp;[NestedListEnum&lt;T&gt;]) -&gt; Vec&lt;T&gt; {
    nested_list.split_first().map_or(vec![], |(head, tail)| {
        [
            match head {
                NestedListEnum::Elem(e) =&gt; vec![e.clone()],
                NestedListEnum::List(l) =&gt; flatten(l),
            },
            flatten(tail),
        ]
        .concat()
    })
}
```

Also, thanks to u/wishthane, in two times to reduce allocations :

```rustd
pub fn flatten&lt;T: Clone&gt;(nested_list: &amp;[NestedListEnum&lt;T&gt;]) -&gt; Vec&lt;T&gt; {
    let mut output = vec![];
    flatten_util(&amp;nested_list, &amp;mut output);
    output
}

fn flatten_util&lt;T: Clone&gt;(nested_list: &amp;[NestedListEnum&lt;T&gt;], output: &amp;mut Vec&lt;T&gt;) {
    for nl in nested_list {
        match nl {
            NestedListEnum::Elem(e) =&gt; output.push(e.clone()),
            NestedListEnum::List(l) =&gt; flatten_util(l, output),
        }
    }
}

```
## [9][Rust Game Development - Ecosystem Survey](https://www.reddit.com/r/rust/comments/ezrk0y/rust_game_development_ecosystem_survey/)
- url: https://rust-gamedev.github.io/posts/survey-01/
---

## [10][Find random long/lat for a given long/lat within a radius](https://www.reddit.com/r/rust/comments/f08lqu/find_random_longlat_for_a_given_longlat_within_a/)
- url: https://www.reddit.com/r/rust/comments/f08lqu/find_random_longlat_for_a_given_longlat_within_a/
---
I have a location (longitude and latitude). I am wondering if it is possible to get a random location for a given radius, so e.g.: give me a random longitude and latitude within X meters of longitude Y and latitude Z.

I had a look in the examples of crate  [https://github.com/georust/geo/tree/master/geo/examples](https://github.com/georust/geo/tree/master/geo/examples), but I did not saw a method that can do this.
## [11][rust web benchmark](https://www.reddit.com/r/rust/comments/f06zzq/rust_web_benchmark/)
- url: https://www.reddit.com/r/rust/comments/f06zzq/rust_web_benchmark/
---
Hi,

I need to benchmark the impact of some kuberbetes addon at work. We mostly use http services for whatever is deployed there coded in a few languages.

I took the opportunity to testout hello world http servers in different languages:

go, java and of course, rust.

For each language, the logic is the same, create a http handler and make it wait 26ms before writing hello in the http response.What I do is use 6 replicas of the pods, and simulate 50k reqs with 500 concurrent workers.

the code for rust is the following (and I tried different servers before)

    use actix_web::{web, App, HttpRequest, HttpServer};
    use std::{thread, time};
    
    async fn index(_req: HttpRequest) -&gt; &amp;'static str {
        thread::sleep(time::Duration::from_millis(26));
        "hi rust!"
    }
    
    #[actix_rt::main]
    async fn main() -&gt; std::io::Result&lt;()&gt; {
        HttpServer::new(|| App::new().service(web::resource("/*").to(index)))
            .bind("0.0.0.0:8080")?
            .run()
            .await
    }

Results are kind of expected:

rust uses 10x less cpu and 30x less memory than java, 3x less cpu than go.

**BUT**

1. There are a few request errors
2. The number of reqs/s is half of what both java and go can handle (for the same number f replicas)

Of course, if I add more rust replicas, it scales up and can come on par with go and java, but why is my rust web server isn't better at managing concurrent requests? Is it the way I thread sleep? Is it a config issue?

Help me use this opportunity to introduce rust in my company!

**EDIT**

Indeed, the blocking thread was the killer here. Now rust has 4x more requests handled!thanks [K900\_](https://www.reddit.com/user/K900_/) and [pbspbsingh](https://www.reddit.com/user/pbspbsingh/)!
## [12][Announcing shaku, a dependency injection library](https://www.reddit.com/r/rust/comments/ezv5zf/announcing_shaku_a_dependency_injection_library/)
- url: https://www.reddit.com/r/rust/comments/ezv5zf/announcing_shaku_a_dependency_injection_library/
---
[crate], [docs], [repo], [rocket integration]

I've been working on this crate for the past few months, and now I think it's ready to see what others think of it. Shaku is a [dependency injection] library which, unlike most other Rust DI libraries, allows you to have both long living and temporary services, and checks dependencies at application startup.

For example, a database connection pool lives for the application lifetime, but the connections it provides may only live for the lifetime of a request (temporary service). If a service is registered that requires a connection, but no registered service provides a connection, then an error will be thrown on startup (when building the DI container).

Most of the work is taken care of by a (optional) derive macro, so there is minimal boilerplate. See the [getting started guide] on the [docs] for a walkthrough!

I'm really interested in what you think of the API, usability, documentation, etc, so please open issues! I plan on continuing development of the crate and getting it to 1.0 (and beyond), so your thoughts would be greatly appreciated.

[users.rust-lang.org thread]

[crate]: https://crates.io/crates/shaku
[docs]: https://docs.rs/shaku/0.1.0/shaku/
[repo]: https://github.com/Mcat12/shaku
[rocket integration]: https://crates.io/crates/shaku_rocket
[dependency injection]: https://en.wikipedia.org/wiki/Dependency_injection
[getting started guide]: https://docs.rs/shaku/0.1.0/shaku/#getting-started
[users.rust-lang.org thread]: https://users.rust-lang.org/t/announcing-shaku-a-dependency-injection-library/37924
