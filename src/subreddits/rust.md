# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (24/2020)!](https://www.reddit.com/r/rust/comments/gyswpo/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/gyswpo/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/guo4c1/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.44]](https://www.reddit.com/r/rust/comments/gz8ut5/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/gz8ut5/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/g6v98u/official_rrust_whos_hiring_thread_for_jobseekers/).

This thread will be periodically stickied to the top of /r/rust for improved visibility. The thread will be refreshed and posted anew when the next version of Rust releases in six weeks.

Please adhere to the following rules when posting (adapted from /r/cpp's jobs thread).
# Rules for individuals:

 * Don't create top-level comments; those are for employers.
 * Feel free to reply to top-level comments with on-topic questions.
 * I will create a stickied top-level comment for individuals looking for work.
 * I will create an additional top-level comment for meta discussion.

# Rules for employers:

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
## [3][Ropey 1.2 - an editable text buffer for Rust](https://www.reddit.com/r/rust/comments/h8o6p4/ropey_12_an_editable_text_buffer_for_rust/)
- url: https://www.reddit.com/r/rust/comments/h8o6p4/ropey_12_an_editable_text_buffer_for_rust/
---
[Ropey v1.2.0](https://crates.io/crates/ropey/1.2.0)

Ropey is a utf8 text rope for Rust, designed to be the backing text buffer for applications such as text editors. Ropey is fast, robust, and can handle huge texts and memory-incoherent edits with ease.

Since the last release announcement on Reddit (v1.0.0) there have been two main feature additions:

* Ropey's iterators have generally been made more flexible.  You can now directly create iterators that start anywhere within the text of the rope, and you can step backwards via a new `prev()` method.
* Ropey can now convert between scalar value indices and utf16 code unit indices.  This is useful for interop with APIs that may use utf16, such as the Language Server Protocol.
## [4][Fast 2D rendering on GPU](https://www.reddit.com/r/rust/comments/h8hx7e/fast_2d_rendering_on_gpu/)
- url: https://raphlinus.github.io/rust/graphics/gpu/2020/06/13/fast-2d-rendering.html
---

## [5][JetBrains 2020 dev ecosystem survey](https://www.reddit.com/r/rust/comments/h8pur7/jetbrains_2020_dev_ecosystem_survey/)
- url: https://www.reddit.com/r/rust/comments/h8pur7/jetbrains_2020_dev_ecosystem_survey/
---
Rust seems to be the only language (at the moment) where more people are planning on adopting it than are currently using it:

[https://www.jetbrains.com/lp/devecosystem-2020/](https://www.jetbrains.com/lp/devecosystem-2020/)

and some rust specific analysis:

[https://www.jetbrains.com/lp/devecosystem-2020/rust/](https://www.jetbrains.com/lp/devecosystem-2020/rust/)
## [6][Microsoft: Rust Is the Industry’s ‘Best Chance’ at Safe Systems Programming](https://www.reddit.com/r/rust/comments/h84zwp/microsoft_rust_is_the_industrys_best_chance_at/)
- url: https://thenewstack.io/microsoft-rust-is-the-industrys-best-chance-at-safe-systems-programming/
---

## [7][Eddit - A basic text editor with GTK and Rust](https://www.reddit.com/r/rust/comments/h8rl0l/eddit_a_basic_text_editor_with_gtk_and_rust/)
- url: https://www.reddit.com/r/rust/comments/h8rl0l/eddit_a_basic_text_editor_with_gtk_and_rust/
---
Hi y'all,

I would like to share my first project in Rust. The repo is available in [GitHub](https://github.com/maze-n/eddit). Please check it out guys.

[https://github.com/maze-n/eddit](https://github.com/maze-n/eddit)
## [8][Do I use gfx-rs or wgpu-rs?](https://www.reddit.com/r/rust/comments/h8skin/do_i_use_gfxrs_or_wgpurs/)
- url: https://www.reddit.com/r/rust/comments/h8skin/do_i_use_gfxrs_or_wgpurs/
---
I'm working on a 3D Rendering Engine as a component for my Game Engine, and I'm about to begin programming it. I originally intended to use gfx-hal, but I thought that it was a little overkill for what I'm making. I've narrowed my options down to gfx-rs or wgpu-rs (unless someone points out a better API that I didn't know of), and I'm not sure which to choose, does anyone have any suggestions for which?

&amp;#x200B;

Edit: I just found out that wgpu-rs is actually a sort of layer on top of gfx-rs
## [9][Compiler isn't printing line where erroneous code is](https://www.reddit.com/r/rust/comments/h8ptwc/compiler_isnt_printing_line_where_erroneous_code/)
- url: https://www.reddit.com/r/rust/comments/h8ptwc/compiler_isnt_printing_line_where_erroneous_code/
---
Hey y'all. I could use some help finding this warning. *Somewhere* in my code, I'm getting these warnings, but the compiler isn't giving me a line to jump to. I'd post my code, but I don't know where to look. Is there any flag I'm missing that could give me a line number and file?

    $ cargo build --features="rustls/dangerous_configuration"
    warning: variable does not need to be mutable
      |
      = note: `#[warn(unused_mut)]` on by default
    
    warning: unused `std::result::Result` that must be used
      |
      = note: `#[warn(unused_must_use)]` on by default
      = note: this `Result` may be an `Err` variant, which should be handled
    
    warning: 2 warnings emitted
    
        Finished dev [unoptimized + debuginfo] target(s) in 0.08s
## [10][Mun's "Make It or Break It" contest](https://www.reddit.com/r/rust/comments/h8d1vb/muns_make_it_or_break_it_contest/)
- url: https://www.reddit.com/r/rust/comments/h8d1vb/muns_make_it_or_break_it_contest/
---
In honour of the [Mun v0.2 release](https://github.com/mun-lang/mun/releases/tag/v0.2.0) we are launching the ***Make It or Break It*** contest. We want to invite both content creators and developers to use Mun v0.2 and *make* or *break* something with/in Mun.

There are two tracks in which you can participate:

***Make It***

Mun was designed to empower creation through iteration, by using natively supported hot reloading for data and functions. Even though Mun v0.2 is still lacking language features - such as enums and arrays - and is limited to single source files, we are excited to see what games and apps people come up with. After all, didn't great artists like Monet create their masterpieces by limiting available tools?

For inspiration, have a look at the classic game [Pong](https://github.com/mun-lang/example-rs/commit/7a9a3757e6360a60a5c37a831afceb48e16d9b96) ([Making of video](https://www.youtube.com/watch?v=n1ea4QUepSU)) that Mun Team devs made.

***Break It***

Mun consists of the Mun ABI[^(1)](https://github.com/mun-lang/mun/blob/master/crates/mun_abi/src/autogen_impl.rs), Mun Compiler[^(1)](https://github.com/mun-lang/mun/tree/master/crates/mun_syntax/src/tests)^(,)[^(2)](https://github.com/mun-lang/mun/blob/master/crates/mun_hir/src/tests.rs)^(,)[^(3)](https://github.com/mun-lang/mun/blob/master/crates/mun_codegen/src/test.rs), and Mun Runtime[^(1)](https://github.com/mun-lang/mun/tree/master/crates/mun_runtime/tests)^(,)[^(2)](https://github.com/mun-lang/mun/blob/master/crates/mun_runtime_capi/src/tests.rs); each with their own tests. Still, there are a lot of untested code paths and other (non-logical) ways that Mun can crash and 🔥. Can you find the craziest, most unexpected bugs?

For inspiration, have a look at breaking Mun's [struct hot reloading](https://mun-lang.org/blog/2020/05/01/memory-mapping) or try fuzzing ([\#152](https://github.com/mun-lang/mun/issues/152)).

***Award***

A contest wouldn't be a contest if there was no award. At the end of every season (i.e. release cycle), we'll release a poll to our Discord community to determine the most popular ***Make It or Break It*** demos. The top contender will be awarded with a mention on our website, but the Mun Team reserves the right to award multiple top contenders (always in order of popularity).

Additionally, all (merged) submissions will be credited on our [Rust](https://github.com/mun-lang/example-rs) example repository.

For more information, please have a look at the [GitHub Issue](https://github.com/mun-lang/mun/issues/220).
## [11][Can someone review my code? Have a function that seeds data, but I'm not quite sure if it's done in the best way.](https://www.reddit.com/r/rust/comments/h8ozan/can_someone_review_my_code_have_a_function_that/)
- url: https://www.reddit.com/r/rust/comments/h8ozan/can_someone_review_my_code_have_a_function_that/
---
So here is my code: 

    pub async fn seed_recipe(pool: &amp;PgPool, recipe: &amp;Recipe, ingredients: &amp;Vec&lt;Ingredient&gt;) -&gt; Uuid {
        let pool_ref = Rc::new(RefCell::new(pool));
        let mut sql_ingredients = "INSERT INTO ingredients(id, name, created_at) VALUES".to_owned();
        let mut sql_recipes_ingredients = "INSERT INTO recipes_ingredients(recipe_id, ingredient_id) VALUES".to_owned();
        let row: (Uuid,) = sqlx::query_as("
            INSERT INTO recipes (id, name, created_at)
            VALUES ($1, $2, $3)
            RETURNING id"
        )
        .bind(recipe.id)
        .bind(recipe.name.clone())
        .bind(recipe.created_at)
        .fetch_one(*pool_ref.borrow_mut()).await.expect("Error inserting recipe.");
    
        for ingredient in ingredients {
            let ingredient_values = &amp;format!("('{}', '{}', '{}'),", ingredient.id, ingredient.name, ingredient.created_at);
            let ingredient_recipe_values = &amp;format!("('{}', '{}'),", recipe.id, ingredient.id);
            sql_ingredients.push_str(ingredient_values);
            sql_recipes_ingredients.push_str(ingredient_recipe_values);
        }
    
        sql_ingredients.pop();
        sql_recipes_ingredients.pop();
    
        sqlx::query(sql_ingredients.as_str()).execute(*pool_ref.borrow_mut()).await.expect("Error inserting ingredients.");
        sqlx::query(sql_recipes_ingredients.as_str()).execute(*pool_ref.borrow_mut()).await.expect("Error inserting recipes_ingredients.");
    
        row.0
    
    }

Basically, this just does a couple of inserts. First, I insert a recipe, then a list of ingredients that belong to that recipe and finally a bunch of associations between those records within a many-to-many table. 

I'm not sure if the way I'm building the queries to insert is the most efficient or if I'm using `Rc&lt;T&gt;`and `RefCell&lt;T&gt;` properly. 

I was wondering if someone could help me out. Also, I'm using expects instead of handling the errors because this is supposed to be a utility function for my tests. If it fails, it fails, it shouldn't handle anything special in the case of an error. 

Thanks in advance.
## [12][Beginner question: Best practice for performing gaussian elimination on a sparse matrix (and keeping the borrowchecker happy)](https://www.reddit.com/r/rust/comments/h8sd46/beginner_question_best_practice_for_performing/)
- url: https://www.reddit.com/r/rust/comments/h8sd46/beginner_question_best_practice_for_performing/
---
I want to represent my matrix as Vec&lt;Vec&lt;(i32,f64)&gt;&gt; - this is a list gor each row which contains tuples (column index, value).

Now for the elimination I fetch a row (immutable borrow) and iterate over all rows below which of course is a mut borrow and makes the borrow checker quite unhappy.

Therefore the question is - what would be the best practice here? Wrap the rows in RefCell or use a unsafe block?

Any hints would be highly appreciated - I can of course post code if my question is unclear.
