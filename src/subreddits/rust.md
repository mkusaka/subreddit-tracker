# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (35/2020)!](https://www.reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 353](https://www.reddit.com/r/rust/comments/ih8uai/this_week_in_rust_353/)
- url: https://this-week-in-rust.org/blog/2020/08/26/this-week-in-rust-353/
---

## [3][Announcing GlueSQL: SQL database fully written in Rust &amp; WebAssembly support](https://www.reddit.com/r/rust/comments/ihdx7o/announcing_gluesql_sql_database_fully_written_in/)
- url: https://www.reddit.com/r/rust/comments/ihdx7o/announcing_gluesql_sql_database_fully_written_in/
---
GlueSQL is the new sql database that I am working on.

It's still in pretty early stage, but it supports basic SQL queries:  aggregation, join and nested select.

Please take a look: [https://github.com/gluesql/gluesql](https://github.com/gluesql/gluesql)

Key features

* Fully written in Rust
* Pure functional execution layer
* Easily-swappable storage (default storage is [sled](https://github.com/spacejam/sled))
* WebAssembly support: [https://github.com/gluesql/gluesql-js](https://github.com/gluesql/gluesql-js)

In case of wasm, please try the Web Dashboard Demo: [https://gluesql.org/playground](https://gluesql.org/playground)

&amp;#x200B;
## [4][openrazer-rs, Rust bindings for the OpenRazer daemon](https://www.reddit.com/r/rust/comments/ihijbw/openrazerrs_rust_bindings_for_the_openrazer_daemon/)
- url: https://crates.io/crates/openrazer
---

## [5][Fixing include_bytes!](https://www.reddit.com/r/rust/comments/ih7ecn/fixing_include_bytes/)
- url: https://jack.wrenn.fyi/blog/include-transmute/
---

## [6][Is rust suitable for competitive programming ?](https://www.reddit.com/r/rust/comments/ihk7o8/is_rust_suitable_for_competitive_programming/)
- url: https://www.reddit.com/r/rust/comments/ihk7o8/is_rust_suitable_for_competitive_programming/
---
Hello community ,I hope you're doing good . As a beginner on rust , I had the idea of learning the langage by participating into competitive programming  contest ( like binary search ,reverse strings etc ..).

And I was wondering ,if it was the proper manner to learn  Rust. Should I keep on the cookbook made by Rust itself to master all the idea behind the langage , or should I learn by project or by training by participating into contest like competitive programming ?

&amp;#x200B;
## [7][Rust Deref coercion confusion.](https://www.reddit.com/r/rust/comments/ihjnjo/rust_deref_coercion_confusion/)
- url: https://www.reddit.com/r/rust/comments/ihjnjo/rust_deref_coercion_confusion/
---
When I started C I got scared of pointers. later i found out that how amazing and strong pointers are.

Now I am learning Rust and  in rust references are like on steroids or are like sneaky. Coming from C reading rust references seems quite difficult for me.

    enum List{
        Cons(Rc&lt;RefCell&lt;i32&gt;&gt;, Rc&lt;List&gt;),
        Nil,
    }
    
    use crate::List::{Cons, Nil};
    use std::cell::RefCell;
    use std::rc::Rc;
    
    fn main() {
        let value = Rc::new(RefCell::new(5));
    
        let a = Rc::new(Cons(Rc::clone(&amp;value), Rc::new(Nil)));
    
        let b = Cons(Rc::new(RefCell::new(6)), Rc::clone(&amp;a));
        let c = Cons(Rc::new(RefCell::new(10)), Rc::clone(&amp;a));
    
        *value.borrow_mut() += 10;

This code is copied from Rust docs. on last line they used a single \* , I understand that how coercion works but its quite confusing that i think since value is wrapped twice in smart pointers shouldn't it be like \*\*.

Anyone can share any links that explains this concepts clearly. (videos would be even great or talks would be amazing)
## [8][Deep Learning in Rust](https://www.reddit.com/r/rust/comments/igz8iv/deep_learning_in_rust/)
- url: https://www.reddit.com/r/rust/comments/igz8iv/deep_learning_in_rust/
---
I am in a bit of dilemma , I learned C++ to implement deep learning algorithms , I am using DL for the purpose of macroeconomic simulations, recently I came across rust and instantly fall in love with the syntax of the language. Now I am in dilemma if i should implement DL algorithms in Rust or C++ and if Rust have any advantage over C++ in future ?? Thanks in advance to the vibrant community
## [9][ANN: Ocypod 0.5.0 (language-agnostic, Redis-backed job queue server, now fully async)](https://www.reddit.com/r/rust/comments/ihkjr0/ann_ocypod_050_languageagnostic_redisbacked_job/)
- url: https://www.reddit.com/r/rust/comments/ihkjr0/ann_ocypod_050_languageagnostic_redisbacked_job/
---
* [user guide](https://ocypod.readthedocs.io/en/latest/)
* [source](https://github.com/davechallis/ocypod)
* [docker images](https://hub.docker.com/repository/docker/davechallis/ocypod/)

Ocypod is an open source job queue server implemented using Rust/Actix/Redis. It's been used production at my workplace for over a year now, having processed ~500m jobs without any issues.

This release focused entirely on updating the internals. The previous version was built on Actix Web 1.0 and Actix actors were used to provide access to Redis via a sync connection.

This update now uses Actix Web 2.0 and async/await throughout (including Redis communication), so no more actors/Actix needed.

The whole async/await migration process was less painful than I'd anticipated, bar a couple of minor stumbling blocks to figure out.
## [10][Polars DataFrame library; Missing features?](https://www.reddit.com/r/rust/comments/ihgazu/polars_dataframe_library_missing_features/)
- url: https://www.reddit.com/r/rust/comments/ihgazu/polars_dataframe_library_missing_features/
---
Yesterday I released version `0.4` of [Polars](https://github.com/ritchie46/polars), a DataFrame library in Rust. I've been working on it for a few months now. I believe it is already pretty feature complete. However, I'd like some feedback and get my head out of the ivory tower. 

What do you guys think? Which operations should definitely be supported by DataFrame library.
## [11][Spawning Tasks on Multiple Thread Pools in Rust](https://www.reddit.com/r/rust/comments/ihep0e/spawning_tasks_on_multiple_thread_pools_in_rust/)
- url: https://pkolaczk.github.io/multiple-threadpools-rust/
---

## [12][Motivations for NoProto Library](https://www.reddit.com/r/rust/comments/ih4qm6/motivations_for_noproto_library/)
- url: https://www.reddit.com/r/rust/comments/ih4qm6/motivations_for_noproto_library/
---
I'm working on a new database and went through the process of trying to find a good serialization format that meets the performance and technical needs of the project. 

One of the frequent operations with secondary indexes is to grab a value out of a row and insert it into an index table.  This operation needs to be efficient and fast to keep the insert hot path running as quickly as possible.  

This requirement automatically excludes a majority of serialization formats.  Even high performance/efficient formats like BSON and Message Pack are designed to deserialize the entire buffer before allowing access to any internal values.  This makes these formats not much better than JSON from a performance standpoint.  The crux of the issue with these formats is if you need to read 4 bytes from a buffer that is a kilobyte in size, you will need to deserialize the whole kilobyte object before you can get your 4 byte value.

This is one of the major advantages of projects like Flatbuffers and CapN Proto, you can read internal values from the buffer without deserializing the whole thing.  So even if you have a large buffer that is several kilobytes large, you will only need to deserialize the data you need and maybe a few other bytes to traverse the buffer.

Unfortunately there's a few tradeoffs with Flatbuffers &amp; CapN Proto.  The first is that new schemas must be compiled into the source language you're using (Rust, C, Java, etc) and then included into your  binary.  This makes these formats a non starter for my use case, I can't ask the end user to recompile the database binary every time there is a schema change.

The second tradeoff is that the these buffers cannot be mutated.  So once you create a FlatBuffer or CapNProto buffer the data is essentially in stone.  If you want to update or delete a value inside a buffer after it's been created or after opening an existing buffer you simply can't.  The correct process to perform an update would be to recursively copy all values in the existing buffer and either don't include the value in the new buffer to delete it or update the value into the new buffer.  As you can imagine, this would be extremely tedious and there's no abstract way to set this process up.  You'd need an "update" or "delete" function for every single value in the buffer.

Now Flatbuffers &amp; CapN Proto is off the table.  The "table" is now... empty.

We need a format that:

1. You can traverse efficiently without parsing the whole buffer first.
2. Allows efficient mutations in place.
3. Allows the schema of the buffers to change without a compilation step

As far as I know, this simply hasn't existed until now.  I wrote the [NoProto](https://crates.io/crates/no_proto) crate to solve exactly this problem.

How does it work?

Schemas are specially formatted JSON objects parsed at runtime.  The parsed schemas are converted into "factories" which can be used to serialize and deserialize any number of buffers. 

```rs
let user_factory = NP_Factory::new(r#"{
    "type": "table",
    "columns": [
        ["name",   {"type": "string"}],
        ["age",    {"type": "u16", "default": 0}],
        ["tags",   {"type": "list", "of": {
            "type": "string"
        }}]
    ]
}"#).unwrap();
```

The schemas are validated to follow a specific format, which means the validation or JSON parsing can fail, thus the call to `.unwrap()`.

Like Protocol Buffers, CapN Proto &amp; Flatbuffers there are ways you can safely mutate the schema without breaking it's ability to parse previous buffers. 

```rs
// safe mutation from previous schema, adding new column at the bottom
let user_factory = NP_Factory::new(r#"{
    "type": "table",
    "columns": [
        ["name",   {"type": "string"}],
        ["age",    {"type": "u16", "default": 0}],
        ["tags",   {"type": "list", "of": {
            "type": "string"
        }}],
        ["phone",    {"type": "string"}],
    ]
}"#)?;

// UNSAFE mutation from previous schema, deleting columns entirely (changes column order).
let user_factory = NP_Factory::new(r#"{
    "type": "table",
    "columns": [
        ["name",   {"type": "string"}],
        ["tags",   {"type": "list", "of": {
            "type": "string"
        }}]
    ]
}"#)?;
```

Once you have a factory, creating and opening buffers is easy and efficient.
```rs
// create a new empty buffer
let mut user_buffer = user_factory.empty_buffer(None, None);

// assign the name column
user_buffer.deep_set("name", String::from("Billy Joel")).unwrap();

// assign nested internal values, sets the first tag element
user_buffer.deep_set("tags.0", String::from("first tag")).unwrap();

// close buffer and get internal bytes
// this is essentially a free operation, just moves the internal bytes out of the buffer
let user_bytes: Vec&lt;u8&gt; = user_buffer.close();

// user_bytes can be sent across the network, saved to disk, or whatever else you'd like... 
// then to open it again...

// open the buffer, essentially free operation as it just moves the Vec.  
// Zero parsing happens here
let user_buffer = user_factory.open_buffer(user_bytes);

// efficiently traverse the buffer to get internal value
// only reads the value and a few other bytes from the buffer.
let tag = user_buffer.deep_get::&lt;String&gt;("tags.0").unwrap();
assert_eq!(tag, Some(Box::new(String::from("first tag"))));

// get nested internal value, the age field
let age = user_buffer.deep_get::&lt;u16&gt;("age")?;
// returns default value from schema
assert_eq!(age, Some(Box::new(0u16)));
```

So we get the dynamic capabilities of something like JSON with the performance of CapNProto &amp; Flatbuffers.  

The library is also fully documented and supports `no_std`, so it will work with WASM as long as you include an allocator.

I hope you'll find this library useful and any feedback / suggestions are welcome!

https://crates.io/crates/no_proto
