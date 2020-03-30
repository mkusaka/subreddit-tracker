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
## [2][What's everyone working on this week (14/2020)?](https://www.reddit.com/r/rust/comments/frff7k/whats_everyone_working_on_this_week_142020/)
- url: https://www.reddit.com/r/rust/comments/frff7k/whats_everyone_working_on_this_week_142020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-14-2020/40159?u=llogiq)!
## [3][Rust Logo should be part of Font-Awesome](https://www.reddit.com/r/rust/comments/frmw35/rust_logo_should_be_part_of_fontawesome/)
- url: https://www.reddit.com/r/rust/comments/frmw35/rust_logo_should_be_part_of_fontawesome/
---
Please leave a thumps up so that the Rust logo gets added to Font-Awesome !  


[https://github.com/FortAwesome/Font-Awesome/issues/13764](https://github.com/FortAwesome/Font-Awesome/issues/13764)
## [4][How many error types should a library have?](https://www.reddit.com/r/rust/comments/frk192/how_many_error_types_should_a_library_have/)
- url: https://www.reddit.com/r/rust/comments/frk192/how_many_error_types_should_a_library_have/
---
I've recently been thinking about error type best practices and came across [this github issue](https://github.com/rust-lang-nursery/failure/issues/7) asking about how many error types a crate should have.

The discussion is from 2017, and I was wondering what current opinions are on how many error types a library should have (thinking of a large crate)? One per library, one per module, or even one per function?

Or a combo like the stdlib based on how much usage can be grouped together?
## [5][dbcrossbar 0.3.1: Copy large tables between BigQuery, PostgreSQL, RedShift, CSV, S3, etc. (preview release, uses async Rust)](https://www.reddit.com/r/rust/comments/froyzx/dbcrossbar_031_copy_large_tables_between_bigquery/)
- url: https://www.reddit.com/r/rust/comments/froyzx/dbcrossbar_031_copy_large_tables_between_bigquery/
---
Hello! Faraday has generously agreed to [open source dbcrossbar](https://www.dbcrossbar.org/), their table-copying tool. This has been heavily used in production for over a year now, and it's ready for a few intrepid Rust developers to take a look. (There will be a larger announcement later for version 1.0.)

You can use `dbcrossbar` to load raw CSVs into PostgreSQL, or to mirror PostgreSQL tables in BigQuery for analysis. It has drivers for [several popular databases](https://www.dbcrossbar.org/dbcrossbar_guide_0.generated.svg). It's designed to operate efficiently on datasets in the 1GB to 500GB range. (Beyond that, it may be better to use a cluster for distributed copies.) No temporary files are required, because dbcrossbar uses multiple async Rust streams and [backpressure](https://ferd.ca/queues-don-t-fix-overload.html) to control data flow. Internally, dbcrossbar represents a table as multiple streams of CSV data, to avoid the need to store an entire table in a single CSV file, and to make efficient use of cloud buckets.

`dbcrossbar` supports common scalar data types, plus arrays, JSON, GeoJSON and UUIDs. It can convert these types between databases. `dbcrossbar` can filter data using `--where`. It can overwrite tables, append to them, or (for BigQuery and PostgreSQL) perform upserts. It knows how to automatically transform a PostgreSQL table definition into a BigQuery one, and vice versa.

Async Rust has proven to be an excellent language for this project. There have definitely been some challenges along the way. But Rust's discipline around ownership and borrowing has made it possible to mix async functions and threads with very few surprises. And solid performance is the default in Rust. Special thanks go to u/burntsushi for his excellent `csv` library, and to the rest of the Rust community for excellent libraries and tools. The list of people to thank would be enormous.

Your bug reports and PRs are very much appreciated. To learn more, see the [manual and installation instructions](https://www.dbcrossbar.org/). And please let me know about any rough edges!
## [6][Game Development in Rust: Monsters and AI!](https://www.reddit.com/r/rust/comments/frkb38/game_development_in_rust_monsters_and_ai/)
- url: https://youtu.be/JxT3r56aqcA
---

## [7][Many-to-many relationships in Diesel? Does anybody know how to do this?](https://www.reddit.com/r/rust/comments/frkta2/manytomany_relationships_in_diesel_does_anybody/)
- url: https://www.reddit.com/r/rust/comments/frkta2/manytomany_relationships_in_diesel_does_anybody/
---
So, I basically have a many-to-many association between two tables. I'm not quite sure how you can handle this with diesel. I think there isn't a particular macro for doing this. 

Is there any standard way of doing? Should I implement it myself?

I would appreciate any suggestions on this.
## [8][Let’s Write a Web Assembly Interpreter (Part 2)](https://www.reddit.com/r/rust/comments/frq583/lets_write_a_web_assembly_interpreter_part_2/)
- url: https://medium.com/@richardanaya/lets-write-a-web-assembly-interpreter-part-2-6c430f3f4bfd
---

## [9][Is there a way to use png's StreamWriter with actix' SyncArbiter?](https://www.reddit.com/r/rust/comments/frp28n/is_there_a_way_to_use_pngs_streamwriter_with/)
- url: https://www.reddit.com/r/rust/comments/frp28n/is_there_a_way_to_use_pngs_streamwriter_with/
---
I'm currently experimenting with actix to create big images. These images won't fit into memory (at least not on my machine) so I'm trying to use the StreamWriter provide by the png crate. My Problem is that it borrows the underlying writer so the borrow checker complains that references in my Actor type need to have static lifetimes. You can find the code [here](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=2e2bb738e719a9383b0914608d1c6ce3), though it won't run online since it requires extern crates.

I know this sub isn't ideal for tech support but since the IRCs are down I don't know where to go with this.
## [10][Global variables for database connection for websockets?](https://www.reddit.com/r/rust/comments/frnwp5/global_variables_for_database_connection_for/)
- url: https://www.reddit.com/r/rust/comments/frnwp5/global_variables_for_database_connection_for/
---
I have clients connected by a websocket to my Rust server. Currently, all those clients hold a Rc&lt;MyDatabase&gt; variable. I am wondering if this is good practice, because it is not really needed for the lifetime of the websocket connection to hold a reference to the database connection. It is only used when the client requests something.

I am wondering what other options are available. I can create a global variable, or a closure that returns a connection. Any thoughts which option I should use?
## [11][R2: A Router in Rust](https://www.reddit.com/r/rust/comments/fr50vl/r2_a_router_in_rust/)
- url: https://www.reddit.com/r/rust/comments/fr50vl/r2_a_router_in_rust/
---
Introducing R2 [**https://r2.rs/,**](https://r2.rs/,) a 'Router in Rust' that I wrote. Read my experience with Rust at [**https://r2.rs/blog/**](https://r2.rs/blog/)
## [12][gdbstub: An implementation of the GDB Remote Serial Protocol in Rust](https://www.reddit.com/r/rust/comments/franjy/gdbstub_an_implementation_of_the_gdb_remote/)
- url: https://docs.rs/gdbstub/0.1.1/gdbstub/
---

