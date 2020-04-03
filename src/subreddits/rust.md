# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (14/2020)!](https://www.reddit.com/r/rust/comments/frfduy/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/frfduy/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fnfky9/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 332](https://www.reddit.com/r/rust/comments/ftl9l3/this_week_in_rust_332/)
- url: https://this-week-in-rust.org/blog/2020/03/31/this-week-in-rust-332/
---

## [3][Iced, a cross-platform GUI library — New release featuring custom styling, event subscriptions, additional widgets, new examples, and more!](https://www.reddit.com/r/rust/comments/ftpld3/iced_a_crossplatform_gui_library_new_release/)
- url: https://github.com/hecrj/iced/pull/253
---

## [4][Like alloca but not quite. Hello to the community and please take a look at my library.](https://www.reddit.com/r/rust/comments/fu37q9/like_alloca_but_not_quite_hello_to_the_community/)
- url: https://www.reddit.com/r/rust/comments/fu37q9/like_alloca_but_not_quite_hello_to_the_community/
---
Hi people. This is my first post on reddit, so congratulate me on losing my reddit-virginity and don’t kick me very much for post design.

I recently wrote one library for rust, because I didn’t find any analogues on the network, and because I had an idea for implementing this feature.

And recently, I discovered that Google somehow does not know about it, so I decided to post it here so that at least someone would look at it and give me feedback or something like that.

And so, there it is. https://crates.io/crates/inplace_it

It solves the problem of allocating arrays of previously unknown length on the stack memory. Well, at least, it's trying to solve.
## [5][[OC] Fondant - macro based, configuration management library for CLI utils](https://www.reddit.com/r/rust/comments/fu8c0p/oc_fondant_macro_based_configuration_management/)
- url: https://i.redd.it/akvcc4e7hlq41.png
---

## [6][Using Rust to follow along with a book directed towards C](https://www.reddit.com/r/rust/comments/fu7q1a/using_rust_to_follow_along_with_a_book_directed/)
- url: https://www.reddit.com/r/rust/comments/fu7q1a/using_rust_to_follow_along_with_a_book_directed/
---
I started studying Computer Organisation and Design (ARM Edition). All codes and other things are directed towards C.
So if anyone has read the book, can I follow along with the book with Rust knowledge?
Can I implement what is being taught using Rust?
## [7][A recursive-descent parser for the Lambda Calculus in Rust](https://www.reddit.com/r/rust/comments/fu041f/a_recursivedescent_parser_for_the_lambda_calculus/)
- url: https://www.reddit.com/r/rust/comments/fu041f/a_recursivedescent_parser_for_the_lambda_calculus/
---
[This](https://christianpoveda.github.io/blog/parsing/) is the second post in my "Writing a Lambda Calculus Implementation in Rust" series. It's about writing a simple parser for the Lambda Calculus implementation I did in the [first post](https://christianpoveda.github.io/blog/untyped-lambda-calculus/)

I'm happy to receive suggestions, opinions and comments in general about it :)
## [8][watson - a minimalistic no_std + alloc web assembly parser](https://www.reddit.com/r/rust/comments/ftwyud/watson_a_minimalistic_no_std_alloc_web_assembly/)
- url: https://github.com/richardanaya/watson
---

## [9][Is there a cleaner way to write this code while appeasing the borrow checker?](https://www.reddit.com/r/rust/comments/fu2xeh/is_there_a_cleaner_way_to_write_this_code_while/)
- url: https://www.reddit.com/r/rust/comments/fu2xeh/is_there_a_cleaner_way_to_write_this_code_while/
---
Sorry for the noob question :) I have a performance critical piece of code similar to this:

    fn parse(data: Vec&lt;u8&gt;) {
      let prefix = &amp;data[0..2];
      let header = &amp;data[3..6];
      // .. some more processing of prefix and header ..
      if prefix[0] == 0 &amp;&amp; header[1] == 1 {
        let text = &amp;data[7..];
        let lowercase_text: Vec&lt;u8&gt; = text.to_ascii_lowercase();
        if lowercase_text[0] == 'a' &amp;&amp; lowercase_text[5] == 'x' &amp;&amp; .. {
          process(prefix, header, &amp;lowercase_text);
        }
      }
    }

What I would like to do is be able to mutate the 'text\_bytes' portion of the current Vec without having to allocate a new Vec:

    fn parse(data: Vec&lt;u8&gt;) {
      let prefix = &amp;data[0..2];
      let header = &amp;data[3..6];
      // .. some more processing of prefix and header ..
      if prefix[0] == 0 &amp;&amp; header[1] == 1 {
        let text = &amp;mut data[7..];
        text.make_ascii_lowercase();
        if text[0] == 'a' &amp;&amp; text[5] == 'x' &amp;&amp; .. {
          process(prefix, header, text);
        }
      }
    }

however, because I have the two immutable borrows in prefix and header, I can't do this. One option is to drop prefix/header first before calling make\_ascii\_lowercase(), and then re-assigning prefix and header again. What would be the most efficient way to do this? Thanks!
## [10][Movie-thumbnailer: create wallpapers from movie scenes. Super minimalist glue for ffmpeg+imagemagick](https://www.reddit.com/r/rust/comments/fu7lrk/moviethumbnailer_create_wallpapers_from_movie/)
- url: https://github.com/indiscipline/movie-thumbnailer
---

## [11][Is there any way of converting futures::Future and std future back and forth?](https://www.reddit.com/r/rust/comments/fu786p/is_there_any_way_of_converting_futuresfuture_and/)
- url: https://www.reddit.com/r/rust/comments/fu786p/is_there_any_way_of_converting_futuresfuture_and/
---
I am stuct with a function that returns a futures::Future. I have a bunch of async functions but apparently I cannot use them in a function that is not async but returns a Future. 
This sounds so stupid really but here it is.
What are my options I cant abandon using async functions because of actix but also I cant use them in from_request apparently because it is not an async function.
the most stupid thing is that I am forced to use blocking functions inside from_request which defeats the entire idea of using async functions for request since the extractors are blocking
## [12][Dataframes in Rust that support String and numeric types in the same value?](https://www.reddit.com/r/rust/comments/fu6p8k/dataframes_in_rust_that_support_string_and/)
- url: https://www.reddit.com/r/rust/comments/fu6p8k/dataframes_in_rust_that_support_string_and/
---
I've just started exploring Pandas and Numpy from the Python world, and I've found that the dataframes there are tremendously useful because you can load all kinds of real world datasets into them that include strings. Its also easy to dynamically add, remove, slice and filter these dataframes too.

However, when I looked for a good library in Rust that had even basic support for Strings, I could find any. Do they exist?
