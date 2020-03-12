# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (11/2020)!](https://www.reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fc7h20/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (11/2020)?](https://www.reddit.com/r/rust/comments/fg0q9l/whats_everyone_working_on_this_week_112020/)
- url: https://www.reddit.com/r/rust/comments/fg0q9l/whats_everyone_working_on_this_week_112020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-11-2020/39269?u=llogiq)!
## [3][This Week in Rust 329](https://www.reddit.com/r/rust/comments/fhcuec/this_week_in_rust_329/)
- url: https://this-week-in-rust.org/blog/2020/03/10/this-week-in-rust-329/
---

## [4][Improving spotify-tui: going async](https://www.reddit.com/r/rust/comments/fheutk/improving_spotifytui_going_async/)
- url: https://www.reddit.com/r/rust/comments/fheutk/improving_spotifytui_going_async/
---
[spotify-tui](https://github.com/Rigellute/spotify-tui) is a terminal client for Spotify written in Rust.

The latest release fixes slow, laggy, or frozen UI experienced by users on slow internet connections.

This fix involved spawning a thread to handle network requests, and sharing state with the main rendering loop using `Arc`, `Mutex` and `sync::channel`.

At the same time, I implemented `async/await` and used `tokio`!

Sharing here in case others face similar challenges in their applications.

If you are interested, you can read more in this [blog post](https://keliris.dev/improving-spotify-tui/).

I would also be interested to hear any thoughts on a better architecture!
## [5][An interactive cheatsheet tool for the command-line written in Rust](https://www.reddit.com/r/rust/comments/fgzifj/an_interactive_cheatsheet_tool_for_the/)
- url: https://i.redd.it/nktdx1akh2m41.gif
---

## [6][Guide on how to write (good) documentation for a Rust crate](https://www.reddit.com/r/rust/comments/fhegra/guide_on_how_to_write_good_documentation_for_a/)
- url: https://blog.guillaume-gomez.fr/articles/2020-03-11+Guide+on+how+to+write+documentation+for+a+Rust+crate
---

## [7][OBS Studio "focus current window" filter and wrapper in Rust](https://www.reddit.com/r/rust/comments/fheh9u/obs_studio_focus_current_window_filter_and/)
- url: https://v.redd.it/4wj22pxj28m41
---

## [8][Tired of trying to debug mysterious panics? Dump a core file on every panic with this code snippet.](https://www.reddit.com/r/rust/comments/fhc1x2/tired_of_trying_to_debug_mysterious_panics_dump_a/)
- url: https://gist.github.com/epilys/a6caba03cb02cfd2880fd80755cd08b8
---

## [9][Be friendly :) A short introduction to Rust for programmers familiar with Go](https://www.reddit.com/r/rust/comments/fh8zb4/be_friendly_a_short_introduction_to_rust_for/)
- url: https://www.reddit.com/r/golang/comments/fgy1fo/a_short_introduction_to_rust_for_programmers/
---

## [10][How to translate this C++ program to Rust?](https://www.reddit.com/r/rust/comments/fhdque/how_to_translate_this_c_program_to_rust/)
- url: https://www.reddit.com/r/rust/comments/fhdque/how_to_translate_this_c_program_to_rust/
---
Hello Rust community!  

I wrote a tutorial how to create a shared library from Clojure code via GraalVM and how to call this C library from C++. The tutorial is [here](https://github.com/borkdude/sci/blob/master/doc/libsci.md).

Clarification:

I'm not looking to do FFI with C++. The only thing I want is to translate this piece of C++ (which interoperates with a C lib) to Rust.

The C++ code:

``` C++
#include &lt;iostream&gt;
#include &lt;libsci.h&gt;

int main(int argc, char* argv[]) {
  graal_isolate_t *isolate = NULL;
  graal_isolatethread_t *thread = NULL;

  if (graal_create_isolate(NULL, &amp;isolate, &amp;thread) != 0) {
    fprintf(stderr, "initialization error\n");
    return 1;
  }

  char *result = eval_string((long)thread, &amp;argv[1][0]);
  std::cout &lt;&lt; result &lt;&lt; std::endl;
  return 0;
}
```

I want to translate this code to Rust. I made a start using the `libloading` crate:

``` rust
extern crate libloading as lib;

fn call_dynamic() -&gt; lib::Result&lt;String&gt; {

    let lib = lib::Library::new("target/libsci.dylib")?;
    // how do I refer to the C type graal_isolate_t from Rust?
    let mut foo: *mut graal_isolate_t = std::ptr::null_mut();
    // same for graal_isolatethread_t
    let mut bar: *mut graal_isolatethread_t = std::ptr::null_mut();

    unsafe {
        let func: lib::Symbol&lt;unsafe extern fn(String) -&gt; String&gt; = lib.get(b"eval_string")?;
        Ok(func(String::from("(+ 1 2 3)")))
    }
}

fn main() {
    println!("{:?}", call_dynamic())
}
```

So the first issue I run into is how to access the type `graal_isolate_t` from Rust.
I suspect the next issue will be how to translate the pointer bits from C++ to Rust.

Your help is welcome.
## [11][Intro to WebAssembly via Rust](https://www.reddit.com/r/rust/comments/fhdwhy/intro_to_webassembly_via_rust/)
- url: https://medium.com/@lironhazan/intro-to-webassembly-via-rust-7365464868bf?source=friends_link&amp;sk=80c9fb8994131b711b02577c615a6a38
---

## [12][An embedded-hal driver for the PZEM004T energy monitor](https://www.reddit.com/r/rust/comments/fhfwbb/an_embeddedhal_driver_for_the_pzem004t_energy/)
- url: https://github.com/iostapyshyn/pzem004t
---

