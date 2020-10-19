# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (43/2020)!](https://www.reddit.com/r/rust/comments/jdwuis/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jdwuis/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (43/2020)?](https://www.reddit.com/r/rust/comments/jdwv7h/whats_everyone_working_on_this_week_432020/)
- url: https://www.reddit.com/r/rust/comments/jdwv7h/whats_everyone_working_on_this_week_432020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-43-2020/50310?u=llogiq)!
## [3][Java's error handling system is better than that of Rust](https://www.reddit.com/r/rust/comments/jdvtu4/javas_error_handling_system_is_better_than_that/)
- url: https://www.reddit.com/r/rust/comments/jdvtu4/javas_error_handling_system_is_better_than_that/
---
I realize that this is a controversial statement, but hear me out. I really like most things in Rust, but the error handling leaves a lot to be desired. 

Rust's error handling system is often praised for its explicitness — the possible error types from a method are right there in the signature, and you can see how the control flow in the program handles it just by reading the function (see an example of this argument [here](http://rantsideasstuff.com/posts/2018/04/20-rust-error-handling-awesome/)). Yet, this system isn't without flaw: one needs to deal with complicated and verbose custom error types (and also complicated macros to make these types — at least compared to Java's system), and one needs to spam their code with `?`, `Result`, and other error handling boilerplate. Furthermore, if an error propagates through several libraries, it's underlying type can be obscured since each time it can be wrapped. 

Java accomplishes the same explicitness and compile-time-correctness-enforcement (mostly, since `RuntimeException` exists, and polymorphism can make it much less explicit), but without these problems. To understand this, consider a toy function that reads a file and returns an ArrayList of lines. In Java, it might look something like this:

	public ArrayList&lt;String&gt; readLinesInFile(File path) throws IOException{
		var list = new ArrayList&lt;String&gt;();
		try( //note that this try exists primarily to give something like RAII to close the reader when the method ends, even if it ends with an exception — it won't do any error handling aside from that
				var reader = new BufferedReader(path);
		){
			String line;
			while((line = reader.readLine()) != null){
				list.add(line);
			}
		}
		return list;
	}

In this example, `throws IOException` is part of the method signature, so a user of this method can trivially determine what exceptions, if any, it can throw. If a function that does not have a `throws &lt;some exception&gt;` in its declaration attempts to call any other function that does claim to throw that exception will result in a compiler error, just as one would expect from a good language (and the kind of behavior that Rust rightly tries to do). This example, though, only shows that Java ensures correctness as Rust does — it isn't really any better than how this would work in Rust, where `std::io::Result` with the unary `?` operator could be used.

The first place where Java's error system really shines is when multiple types of errors can take place. Consider a toy example that will make an HTTP request and then parse JSON from it. In this case, there are fictitious functions (assumed to be from libraries): `String makeHttpRequest(String url) throws IOException, HttpStatusCodeNot200Exception`, `Object parseJson(String json) throw JsonParseException`, each doing what one would expect. The HTTP request and parse function may look like this:

    public Object parseJsonFromHttp(String url) throws IOException, HttpStatusCodeNot200Exception, JsonParseException{
        return parseJson(makeHttpRequest(url));
    }

That's it. That's the entire function, with all error handling laid out there. There is no complicated and verbose error chain type squirreled away in some module, no complicated and non-standard macro to generate that type in a highly implicit manner, and no additional dependencies. Clearly, it's much nicer than Rust to write this function. 

But what if the author of this function didn't want the user of their function to have to worry about HTTP status codes? No problem, that's easily handled:

    public Object easyParseJsonFromHttp(String url) throws IOException, JsonParseException{
        try{
            return parseJson(makeHttpRequest(url));
        }catch(HttpStatusCodeNot200Exception e){
            throw new IOException("HTTP returned " + e.getStatusCode() ".", e);
        }
    }

Now, the user only has to worry about IOException as one . This is a bit nicer than Rust's system of using pattern matching for the error since in the `try`'s main body one can see the normal error-free control flow of the function, and at the bottom with each `catch` (Java allows multiple catches with a single try) one can see what happens in each scenario where it goes wrong. In Java, one sometimes needs to have some additional control flow here to properly handle it, but Rust's pattern matching in the `catch` blocks would greatly reduce this to nearly nothing. For example, with this system, it's also possible for a user of the function to care about the status code if they so desire:

	try{
	    Object parsedJson = easyParseJsonFromHttp("https://example.com/some/json");
	    //do something with the parsed json
	}catch(IOException e){
	    if(e.getCause() instanceof HttpStatusCodeNot200Exception){
	        //do something with this information
	    }else{
	        throw e;
	    }
	}

The last advantage of Java's system is that the errors, when printed (either because `main` throws some exception type(s), or because `printStackTrace()` was called), are both more informative and easier to read. Consider this example stack trace (taken from [here](https://www.twilio.com/blog/how-to-read-and-understand-a-java-stacktrace)):

	Exception in thread "main" com.myproject.module.MyProjectFooBarException: The number of FooBars cannot be zero
	    at com.myproject.module.MyProject.anotherMethod(MyProject.java:19)
	    at com.myproject.module.MyProject.someMethod(MyProject.java:12)
	    at com.myproject.module.MyProject.main(MyProject.java:8)
	Caused by: java.lang.ArithmeticException: The denominator must not be zero
	    at org.apache.commons.lang3.math.Fraction.getFraction(Fraction.java:143)
	    at com.myproject.module.MyProject.anotherMethod(MyProject.java:17)
	    ... 2 more

It gives a lot of information: the thread that the exception took place in, the exact exception, a human-readable message to go with it (Java can bundle machine readable information in an exception type too), a full stack trace giving a lot of detail of where the exception took place, and a very useful `Caused by…` line, with its own stack trace, that can help one drill down to the root cause of an exception. Compare this to Rust, which, in the best case may look like this: `MyProjectFooBarError` (in the case of a unit-type-like struct being used for errors), or, in the best case like this: `MyProjectFooBarError{cause: ArithmeticException("The denominator must not be zero"), message: "The number of FooBars cannot be zero"}`. Sure, there's *probably* enough information there, but it's all packed in on a mostly-machine-readable single line as opposed to Java where it's more for humans (this message is generated from a machine-readable data structure that code can easily query — all of that information is accessible to code too), aside from the extra information of course.

Now, one may object and claim that Rust's system is superior becuase it has clearer control flow — there are no apparantly magical jumps from one method to 10 methods up the call stack where the corresponding catch is. To this, I have four things to say:

 - It's not magical. There are well-defined rules for how this works, and a debugger can instantly tell you where the catch is taking place. 
 - It's not really any better in Rust if an error type is passed several methods up the call chain before its handled, except a debugger can't instantly tell you where that handling is taking place.
 - When looking at a function in Rust, the only clue that implicit error handling is taking place is `?`, which isn't very informative and is easily missed. Furthermore, in Java the exact list of exceptions that can be thrown from this particular method is listed in the method header. Whereas in Rust, at best you need to dig through to find the error type for that crate (not very hard with a goto feature in an IDE), and at worst you don't know which error enum variations from that crate's type can actually result from this method, since most to all erroring functions in a crate are likely to share one type. 
 - If you're writing your Java code correctly, this won't matter (very much), since at the throwing end all you need to worry about is that something erroneous happened, and at the recieving end all you need to worry about is that the function that you just called had an error of the particular exception type.

Lastly, one may object and claim that many to all of Java's advantages are negated by polymorphism — where a function says that it throws some very broad base class of many different concrete exeptions (in the worst case `throws Throwable`, which means that the caller needs to handle literally any error), and/or by `RuntimeException` which tells the compiler to ignore any enforcement of handling. For Rust, these two flaws should not be present, although it is arguable in Java, so, essentially, on these specific points, I agree with the criticism.
## [4][This month and a half in Rune](https://www.reddit.com/r/rust/comments/jdvc8r/this_month_and_a_half_in_rune/)
- url: https://rune-rs.github.io/posts/tmir1/
---

## [5][Why is the box keyword unstable?](https://www.reddit.com/r/rust/comments/jdvzcg/why_is_the_box_keyword_unstable/)
- url: https://www.reddit.com/r/rust/comments/jdvzcg/why_is_the_box_keyword_unstable/
---
I came across the `box` keyword, which seemed like some really cool syntactic sugar. Then I found it was an unstable feature. Surely it should just compile to a call of `Box::new`, or, in a pattern, a dereference?
## [6][Flask Creator &amp; Rust Developer Armin Ronacher Interview](https://www.reddit.com/r/rust/comments/je0f47/flask_creator_rust_developer_armin_ronacher/)
- url: https://evrone.com/armin-ronacher-interview
---

## [7][No, C++ still isn't cutting it.](https://www.reddit.com/r/rust/comments/jdgm6u/no_c_still_isnt_cutting_it/)
- url: https://da-data.blogspot.com/2020/10/no-c-still-isnt-cutting-it.html
---

## [8][Osmflat - efficient random access OpenStreetMap file format](https://www.reddit.com/r/rust/comments/jdmtgf/osmflat_efficient_random_access_openstreetmap/)
- url: https://www.reddit.com/r/rust/comments/jdmtgf/osmflat_efficient_random_access_openstreetmap/
---
We are happy to announce a first release of [Osmflat](https://github.com/boxdot/osmflat-rs): A file format that allows for easy and fast random access to OpenStreetMap data.

The project consists of:

* A `flatdata` format definition, directly usable by all languages `flatdata` supports
* A thin Rust helper library for working with the format
* A compiler (written in Rust), that converts OpenStreetMap's `PBF` format into `osmflat`
* C++ and Rust usage example / benchmarks

# Why another format?

The current OpenStreetMap formats have one major weakness: Objects are referenced by an id that does not relate any information about where in the file that object is stored. `osmflat` allows every every reference to be followed immediately, making working with the data a breeze. Other formats often require the user to build expensive mapping tables before processing the data.

Example: to print all pub names and addresses from an `osmflat` file

    let archive = Osm::open(FileResourceStorage::new(archive_dir))?;
    let nodes_tags = archive.nodes().iter().map(|node| node.tags());
    let ways_tags = archive.ways().iter().map(|way| way.tags());
    for tag_range in nodes_tags.chain(ways_tags) {
        if has_tag(&amp;archive, tag_range.clone(), b"amenity", b"pub") {
            let name = find_tag(&amp;archive, tag_range.clone(), b"name");
            let name = name.map(|s| str::from_utf8(s).unwrap_or("broken pub name"));
            println!("{}", name.unwrap_or("unknown pub name"));
    
            let addrs = iter_tags(&amp;archive, tag_range).filter(|(k, _)| k.starts_with(b"addr:"));
            for (k, v) in addrs {
                if let (Ok(addr_type), Ok(addr)) = (str::from_utf8(k), str::from_utf8(v)) {
                    println!("  {}: {}", addr_type, addr)
                }
            }
        }
    }

The project includes a few other examples that showcase how one can work with the format, e.g. a SVG renderer.

# How to compile an osmflat archive

1. Download a `PBF` archive, e.g. from [https://download.geofabrik.de/index.html](https://download.geofabrik.de/index.html)
2. Compile / install `osmflatc`, e.g. with cargo: `cargo install osmflatc`
3. Run `osmflatc` (be aware you need enough RAM for it to builds the id mapping tables), e.g.: `osmflatc ~/Downloads/berlin-latest.osm.pbf berlin.osmflat`

# How does it work

The format is defined using `flatdata`: *a library providing data structures for convenient creation, storage and access of packed memory-mappable structures with minimal overhead.*

`flatdata` takes care of most low-level details: Handling I/O (using memory mapped files), portable reading of bit-packed structures, etc. It also provides a bit of tooling, e.g. visualizations of the format, or a python-based inspector that allows for quick data debugging.

The format follows the general structure of existing OSM formats (e.g. `PBF`), but it replaces id-based references by index-based references that can be resolved/followed immediately.

# Current state of the project

The format / library is fully working, but we do not expect the format specification to be final: We are hoping for several rounds of feedback from real-world use-cases so that we can refine it.

# Disclaimer

I am working for HERE, a mapping company, but this project is not affiliated with it.

# Resources

* [osmflat](https://github.com/boxdot/osmflat-rs)
* [flatdata](https://github.com/heremaps/flatdata)
## [9][Proving that 1 + 1 = 10 in Rust](https://www.reddit.com/r/rust/comments/je1ajy/proving_that_1_1_10_in_rust/)
- url: https://tavianator.com/2020/one_plus_one.html
---

## [10][Rust Firebird Client updated to v0.12.0 with : Named params support, statements without cursor and many other fixes](https://www.reddit.com/r/rust/comments/jdxtd2/rust_firebird_client_updated_to_v0120_with_named/)
- url: https://github.com/fernandobatels/rsfbclient/releases/tag/v0.12.0
---

## [11][pipe-op a crate that brings the functionality of elixir pipe operator to Rust as a macro.](https://www.reddit.com/r/rust/comments/jdwf4t/pipeop_a_crate_that_brings_the_functionality_of/)
- url: https://www.reddit.com/r/rust/comments/jdwf4t/pipeop_a_crate_that_brings_the_functionality_of/
---
[Github](https://github.com/McRaeAlex/pipe-op)

[Crates.io](https://crates.io/crates/pipe-op)

I created this crate to learn more about proc macros in Rust but found it had practical uses outside a learning experience!

Take a look at the code if you want. I will update it if any bugs are found and accept pull requests. The only thing I hope if if your going to change something please explain why.

It currently supports functions, methods, chained methods, field access, .await, try `?` and blocks `{}`

Example:

    use pipe_op::pipe;
    
    struct Adder {
        value: usize,
    }
    
    impl Adder {
        fn add(&amp;self, num: usize) -&gt; Result&lt;usize, Box&lt;dyn std::error::Error&gt;&gt;  {
            Ok(self.value + num)
        }
    }
    
    fn add(a: usize, b: usize) -&gt; usize {
        a + b
    }
    
    fn main() -&gt; Result&lt;(), Box&lt;dyn std::error::Error&gt;&gt; {
        let a = Adder {value: 10};
        assert_eq!(
            pipe!(
                1,
                add(10),
                add(10),
                add(10),
                add(10),
                a.add(10)?,
                add(10),
                add(10),
                add(10)
            ),
            81
        );
    }

Let me know what you think.
## [12][Design Patterns: Private Methods on a Public Trait](https://www.reddit.com/r/rust/comments/jdlhub/design_patterns_private_methods_on_a_public_trait/)
- url: https://jack.wrenn.fyi/blog/private-trait-methods/
---

