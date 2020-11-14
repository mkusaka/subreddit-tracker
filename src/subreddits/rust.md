# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (46/2020)!](https://www.reddit.com/r/rust/comments/jqrkpa/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jqrkpa/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 364](https://www.reddit.com/r/rust/comments/jslo80/this_week_in_rust_364/)
- url: https://this-week-in-rust.org/blog/2020/11/11/this-week-in-rust-364/
---

## [3][Starship - The minimal, blazing-fast, and infinitely customizable prompt for any shell!](https://www.reddit.com/r/rust/comments/jtq2x8/starship_the_minimal_blazingfast_and_infinitely/)
- url: https://starship.rs/
---

## [4][mockalloc - a crate to test for leaks and other allocation bugs](https://www.reddit.com/r/rust/comments/jttpie/mockalloc_a_crate_to_test_for_leaks_and_other/)
- url: https://www.reddit.com/r/rust/comments/jttpie/mockalloc_a_crate_to_test_for_leaks_and_other/
---
Just published this tiny crate: https://docs.rs/mockalloc/0.1.0/mockalloc/

One of my projects involves implementing some custom container types (vec and map-like containers) and as well as testing that they function as expected, I wanted a simple way to detect bugs in the way they allocate and free memory.

Enter the `#[mockalloc::test]` attribute.

Simply replace `#[test]` with `#[mockalloc::test]`, set `Mockalloc` as your global allocator during `#[cfg(test)]` and common allocation bugs will be detected automatically.
## [5][Tide v0.15.0 has been released!](https://www.reddit.com/r/rust/comments/jtqep1/tide_v0150_has_been_released/)
- url: https://github.com/http-rs/tide/releases/tag/v0.15.0
---

## [6][Users of SQLx, we're looking for your opinions on the future of the crate!](https://www.reddit.com/r/rust/comments/jtjffr/users_of_sqlx_were_looking_for_your_opinions_on/)
- url: https://www.reddit.com/r/rust/comments/jtjffr/users_of_sqlx_were_looking_for_your_opinions_on/
---
For those unaware, [SQLx](https://github.com/launchbadge/sqlx) is an async-first SQL client library conceived, developed and maintained by my boss, a couple coworkers, and myself, at [Launchbadge, LLC](https://launchbadge.com/), a development services company, primarily for use in our own web development stack utilizing Postgres and Actix-web.

We finally kicked out our 0.4 release yesterday after many months in beta. 

We're sorry that it took so long, one of the main reasons being that we had a new project come up with some tight deadlines that we've been working nonstop to deliver on, and haven't had much time or energy to spare for SQLx besides commenting on issues and merging PRs.

Also, since the beginning of the shelter-in-place order for COVID in California (where we're based) we've all been working from home. At the office, SQLx was a nice side-project to work on as a break from our primary tasks but at home there's a lot more distractions to deal with and it's a lot harder (for me, personally) to stay in the working mindset for a solid eight hours a day. There's also much fewer opportunities for collaboration as we'd often sit down to chat about SQLx at lunch and during breaks, which we simply can't do anymore.

Additionally, we're concerned about being able to maintain the quality of the MySQL and SQLite drivers in SQLx, compared to the Postgres driver which we're using all the time. We're very grateful for all the outside contributions but still it's pretty painfully obvious that the former databases don't get quite as much TLC as Postgres does.

And so, we're considering spinning off SQLx into its own entity with one or more full-time developers who will be dedicated to developing and maintaining it. It would most likely be a 501(c) non-profit corporation with public reporting for costs and revenue. To do this, though, it needs to be capable of supporting itself. 

The idea is that, like [jOOQ](http://www.jooq.org/download/) which is a database library for Java that we've used quite a bit in the past, we would charge a yearly license fee to use SQLx with various proprietary databases such as MSSQL, OracleDB and IBM's DB2, as well as provide fee-based consulting and support for enterprises using SQLx. 

SQLx would of course remain free to use for Postgres, MySQL and SQLite but the company would be capable of accepting tax-deductible donations as well for those who wish to support development.

To have some data to inform our decision, we've created a very short survey that we'd like anyone currently using or would be interested in using SQLx to fill out. If you have any feedback that isn't covered by the survey, please also feel free to share it below, or [on our Discord server](https://discord.gg/narftnn7).

https://forms.gle/M1Us31brnveUFAnX9
## [7][gping: ping, but with a graph](https://www.reddit.com/r/rust/comments/jthdba/gping_ping_but_with_a_graph/)
- url: https://github.com/orf/gping
---

## [8][RestQ - a simpler alternative to graphQL suited for database specific operation on a rest api.](https://www.reddit.com/r/rust/comments/jtxxkj/restq_a_simpler_alternative_to_graphql_suited_for/)
- url: https://github.com/ivanceras/restq
---

## [9][speed issue with rust backend](https://www.reddit.com/r/rust/comments/ju1g94/speed_issue_with_rust_backend/)
- url: https://www.reddit.com/r/rust/comments/ju1g94/speed_issue_with_rust_backend/
---
when running this code should give me a performance boost in existing lib which is purely in python

but now the performance is worst

Something I am doing wrong

&amp;#x200B;

https://preview.redd.it/ye3ef7m5c7z51.png?width=1112&amp;format=png&amp;auto=webp&amp;s=c5cc928a2c15468761b5d18566207c3db850a27c

    #[pyfunction]
    pub fn match_string_percentage_list_fn(py: Python, source: &amp;PyString, target: Vec&lt;String&gt;) -&gt; PyResult&lt;HashMap&lt;String, f64&gt;&gt; {
        let s = source.extract().unwrap();
        let obj = StringProcessing::new();
        let s_p = obj.replace_non_letters_non_numbers_with_whitespace(s);
    
        let mut temp_vec: Vec&lt;String&gt; = target.par_iter().map(|i| {
            obj.replace_non_letters_non_numbers_with_whitespace(i.to_string())
        }).collect();
        let mut rez_vec: HashMap&lt;String, f64&gt; = HashMap::new();
    
        for i in temp_vec.iter() {
            let z = normalized_levenshtein(s_p.trim(), i.as_str().trim());
            rez_vec.insert(i.clone(), z);
        }
        Ok(rez_vec)
    }
    
    #[pyfunction]
    pub fn match_string_percentage_fn(py: Python,source: &amp;PyString, target: &amp;PyString) -&gt; f64 {
        let s = source.to_string_lossy().to_string();
        let t = target.to_string_lossy().to_string();
        let obj = StringProcessing::new();
        // let s_p = obj.replace_non_letters_non_numbers_with_whitespace(s);
        // let t_p = obj.replace_non_letters_non_numbers_with_whitespace(t);
        normalized_levenshtein(s.trim(), t.trim())
    }
## [10][Make a Language in Rust, Part Ten: Starting Again](https://www.reddit.com/r/rust/comments/ju1vhl/make_a_language_in_rust_part_ten_starting_again/)
- url: https://arzg.github.io/lang/10/
---

## [11][Knurling-rs changelog #6: the defmt logging framework is now on crates.io 🎉](https://www.reddit.com/r/rust/comments/jtk2g0/knurlingrs_changelog_6_the_defmt_logging/)
- url: https://ferrous-systems.com/blog/knurling-changelog-6/
---

## [12][High school student using rust, looking for help (code review)](https://www.reddit.com/r/rust/comments/jtl9tk/high_school_student_using_rust_looking_for_help/)
- url: https://www.reddit.com/r/rust/comments/jtl9tk/high_school_student_using_rust_looking_for_help/
---
Hi all! 

I'm a high school student using Rust for my school's [Computer Graphics course](https://stuycs.org/courses/graphics/). My teacher provides template code for C and Python, which are all taught in my school. Rust is obviously not, but I'm not the first person to use it for this class, and I believe I won't be the last.

So after finishing the course, I thought it would nice to write templates for Rust so people who are not so familiar with this can do it without setting up everything themselves.

I've finished the first template by adding parts from my own work, and then refactoring, and then heavily documenting. But I don't know if I'm doing things in the "Rust" way, or if I'm writing clear and concise docs. I would like some feedback on this (or anything). I'll be working to finish the rest of the templates soon.

Links:

\- [My work on the course](https://github.com/ruoshui-git/stuy-graphics)

\- [Templates in C and Python, which I'll be referencing](https://github.com/mks66/66source)

\- [My first Rust template](https://github.com/ruoshui-git/stuygfx-rust-src-1), for which I really need feedback

Thanks in advance!
