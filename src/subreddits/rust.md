# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (42/2020)!](https://www.reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 360](https://www.reddit.com/r/rust/comments/jbh0ci/this_week_in_rust_360/)
- url: https://this-week-in-rust.org/blog/2020/10/14/this-week-in-rust-360/
---

## [3][genpdf, a user-friendly PDF generator written in pure Rust](https://www.reddit.com/r/rust/comments/jbk94w/genpdf_a_userfriendly_pdf_generator_written_in/)
- url: https://www.reddit.com/r/rust/comments/jbk94w/genpdf_a_userfriendly_pdf_generator_written_in/
---
genpdf ([crates.io](https://crates.io/crates/genpdf), [lib.rs](https://lib.rs/crates/genpdf), [docs.rs](https://docs.rs/genpdf), [source code](https://git.sr.ht/~ireas/genpdf-rs)) is a high-level PDF generator built on top of [`printpdf`](https://lib.rs/crates/printpdf) and [`rusttype`](https://lib.rs/crates/rusttype).  It takes care of the page layout and text alignment and renders a document tree into a PDF document.  All of its dependencies are written in Rust, so you don’t need any pre-installed libraries or tools.

    // Create a document and set the default font family
    let mut doc = genpdf::Document::new("./fonts", "Liberation")
        .expect("Failed to create PDF document");
    // Change the default settings
    doc.set_margins(10);
    doc.set_title("Demo document");
    // Add one or more elements
    doc.push(genpdf::elements::Paragraph::new("This is a demo document."));
    // Render the document and write it to a file
    doc.render_to_file("output.pdf")
        .expect("Failed to write PDF file");

For a complete example with all supported elements, see the
[`examples/demo.rs`][] file that generates [this PDF document][].

[`examples/demo.rs`]: https://git.sr.ht/~ireas/genpdf-rs/tree/master/examples/demo.rs
[this PDF document]: https://git.sr.ht/~ireas/genpdf-rs/blob/master/examples/demo.pdf

For more information, see the [readme](https://sr.ht/~ireas/genpdf-rs) and the  [API documentation](https://docs.rs/genpdf).
## [4][We need to talk about StackOverflow](https://www.reddit.com/r/rust/comments/jb3ukm/we_need_to_talk_about_stackoverflow/)
- url: https://www.reddit.com/r/rust/comments/jb3ukm/we_need_to_talk_about_stackoverflow/
---
There's one thing I hate more than anything else about Rust - more than confusing lifetime errors, more than compile times, even more than `std::ops::Range`: asking questions on StackOverflow.

55% of the my questions are edited, and 15% are *erroneously* closed as duplicates/too broad by one single user. I won't name them but anyone who has posted a Rust question to StackOverflow will know who I am talking about.

This user often posts useful information, but I did not ask him to be my personal copy editor. If a *single person nitpicked more than half* of all the text he wrote I do not think he would appreciate it. And we are talking nitpicks. Here is a typical edit:

&gt; Convert SystemTime date to ISO 8601 in rust

to

&gt; How do I convert a SystemTime to ISO 8601 in Rust?

The question closures are worse than the edits though. StackOverflow has a meme-level problem with overzealous question closure, and it's especially infuriating because closed questions are almost impossible to reopen (only 6% are). Out of the 4 closed-as-duplicates I have been punished by, I would say only 1 was a genuine duplicate. The others have helpful answers. To have so many questions mistakenly closed by a single prolific user is very frustrating.

The Rust team seem to be keen to make the Rust community welcoming. This is not welcoming. It also does not happen with any other topic I ask about - only Rust.

The thought of asking a question on StackOverflow should not fill me with dread. It should not make me think "god I hope that guy is asleep".
## [5][Blog Post: Study of std::io::Error](https://www.reddit.com/r/rust/comments/jbdk5x/blog_post_study_of_stdioerror/)
- url: https://matklad.github.io/2020/10/15/study-of-std-io-error.html
---

## [6][How does HashMap return a reference to a value?](https://www.reddit.com/r/rust/comments/jbmc1y/how_does_hashmap_return_a_reference_to_a_value/)
- url: https://www.reddit.com/r/rust/comments/jbmc1y/how_does_hashmap_return_a_reference_to_a_value/
---
Hello.

I was wondering how does HashMap return a reference to a value without getting

`returns a value referencing data owned by the current function`

&amp;#x200B;

        pub fn get&lt;BK&gt;(&amp;self, key: &amp;BK) -&gt; Option&lt;&amp;V&gt;
        where
            BK: Hash + Eq + ?Sized,
            K: Borrow&lt;BK&gt;,
        {
            self.root
                .get(hash_key(&amp;*self.hasher, key), 0, key)
                .map(|&amp;(_, ref v)| v)
        }

&amp;#x200B;

I have a similar function `get` for a struct that owns a HashMap and I'm trying to return that value reference and I get a compilation error: `returns a value referencing data owned by the current function`

        fn get&lt;'a&gt;(&amp;self, key: &amp;'a dyn MyTrait) -&gt; Option&lt;&amp;'a dyn MyTrait&gt; {
            ...
            hm.get(&amp;key)
                .map(|v| v as &amp;dyn MyTrait)
        }

&amp;#x200B;

I think I understand the essence of the problem, where I'm a returning a reference to something that `hm` owns, but how does it work for HashMap?

&amp;#x200B;

EDIT: I think I understand something that is missing from my question and is the reason for this issue.

        fn get&lt;'a&gt;(&amp;self, key: &amp;'a dyn MyTrait) -&gt; Option&lt;&amp;'a dyn MyTrait&gt; {
            let hm = self.hm.lock().unwrap();
            hm.get(&amp;key)
                .map(|v| v as &amp;dyn MyTrait)
        }

`hm` is actually a local variable that unlocks [self.hm](https://self.hm) (is behind a mutex)

I redacted this code my original question as I didn't understand this is the problem.

&amp;#x200B;

How do I get around that?

My guess is there's no way other than cloning the value as it's a reference to something behind a mutex. Problem is I cannot have MyTrait be Cloneable because then I get errors for not being about to have MyTrait as a trait object....

&amp;#x200B;

&amp;#x200B;

Thank you!
## [7][Why is it so hard to convert a Stream&lt;Item = Vec&lt;u8&gt;&gt; into a linewise Stream&lt;Item = Result&lt;String&gt;&gt;?](https://www.reddit.com/r/rust/comments/jbkde7/why_is_it_so_hard_to_convert_a_streamitem_vecu8/)
- url: https://www.reddit.com/r/rust/comments/jbkde7/why_is_it_so_hard_to_convert_a_streamitem_vecu8/
---
I'm working on this for about two hours now.

I have a tokio application where I get a `Stream&lt;Item = Vec&lt;u8&gt;&gt;` (actually it is a bit more complex because shiplift (from master branch) returns an enum of buffers, but for the sake...) and I need to convert this to a `Stream&lt;Item = Result&lt;String&gt;&gt;` (parsing linewise into UTF8).

(the ultimate goal is to parse linewise using nom, but having a stream over lines is the first step towards this).

I am lost right now. This seems to be a task that is way too hard for me to implement. Maybe some of you gurus around here can help.

My code is https://gist.github.com/matthiasbeyer/6b5f3a79f75a68c2bdef5536bcd8f57d (I hope I included all relevant parts). Note that the code example actually handles the shiplift return value which is an enum over three different types of buffers.
## [8][Backlog Bonanza -- notes from lang team's October 14th meeting](https://www.reddit.com/r/rust/comments/jbj6xj/backlog_bonanza_notes_from_lang_teams_october/)
- url: https://paper.dropbox.com/doc/Backlog-bonanza--A9jkzp9e4KaNBFyMC0HwCkwkAg-2IcACiM0KX1up1thIeiWh
---

## [9][declio v0.1.0 - Declarative encoders and decoders for binary formats.](https://www.reddit.com/r/rust/comments/jbizl2/declio_v010_declarative_encoders_and_decoders_for/)
- url: https://crates.io/crates/declio
---

## [10][Rye - a tiny experiment in building safe Rust fibers](https://www.reddit.com/r/rust/comments/jb6nhp/rye_a_tiny_experiment_in_building_safe_rust_fibers/)
- url: https://github.com/mpdn/rye
---

## [11][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.47]](https://www.reddit.com/r/rust/comments/jb4cds/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/jb4cds/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/iix8vw/official_rrust_whos_hiring_thread_for_jobseekers/).

This thread will be periodically stickied to the top of /r/rust for improved visibility. The thread will be refreshed and posted anew when the next version of Rust releases in six weeks.

Please adhere to the following rules when posting:
# Rules for individuals:

 * Don't create top-level comments; those are for employers.
 * Feel free to reply to top-level comments with on-topic questions.
 * Anyone seeking work should reply to my stickied top-level comment.
 * Meta-discussion should be reserved for the distinguished comment at the very bottom.

# Rules for employers:

 * To find individuals seeking work, see the replies to the stickied top-level comment; you will need to click the "more comments" link at the bottom of the top-level comment in order to make these replies visible.
 * To make a top-level comment you must be hiring directly; no third-party recruiters.
 * One top-level comment per employer. If you have multiple job openings, please consolidate their descriptions or mention them in replies to your own top-level comment.
 * Proofread your comment after posting it and edit it if necessary to correct mistakes.
 * Please base your comment on the following template:

COMPANY: *[Company name; optionally link to your company's website or careers page.]*

TYPE: *[Full time, part time, internship, contract, etc.]*

DESCRIPTION: *[What does your company do, and what are you using Rust for? How much experience are you seeking and what seniority levels are you hiring for? The more details the better.]*

LOCATION: *[Where are your office or offices located? If your workplace language isn't English-speaking, please specify it.]*

ESTIMATED COMPENSATION: *[Be courteous to your potential future colleagues by attempting to provide at least a rough expectation of wages/salary. If you are listing several positions in the "Description" field above, then feel free to include this information inline above, and put "See above" in this field. If compensation is negotiable, please attempt to provide at least a base estimate from which to begin negotiations. If compensation is highly variable, then feel free to provide a range. If compensation is expected to be offset by other benefits, then please include that information here as well. If you don't have firm numbers but do have relative expectations of candidate expertise (e.g. entry-level, senior), then you may include that here. If you truly have no information, then put "Uncertain" here.* ***This is a new field in our template; please see the meta comment below to discuss it***.*]*

REMOTE: *[Do you offer the option of working remotely? If so, do you require employees to live in certain areas or time zones?]*

VISA: *[Does your company sponsor visas?]*

CONTACT: *[How can someone get in touch with you?]*
## [12][Is there a clever way to deserialize this nightmarish json into either structs or vectors?](https://www.reddit.com/r/rust/comments/jb1bdm/is_there_a_clever_way_to_deserialize_this/)
- url: https://www.reddit.com/r/rust/comments/jb1bdm/is_there_a_clever_way_to_deserialize_this/
---
I'm hoping a Serde wizard can kindly help me as I've been stuck on this all day. I'm trying to deserialize json of a format I've not encountered before which I am struggling with, due to it being a weird mixture of stucts and arrays. These are the 5 overall different message types:

    {"event": "event_name", ...other fields} // other fields vary depending on event
    ["heartbeat", "id"]     //heartbeart
    ["channel_id",  "price",  "count",  "amount" ]     //trade
    ["channel_id", [  ["price",  "count",  "amount"], ["price",  "count",  "amount" ] , ... ] ]    //snapshot
    ["channel_id", [  ["order",  "level",  ], ["order", "level" ] , ... ] ] //other_snapshot

So far I have got something like this, where I am using untagged and tag

    #[derive(Serialize, Deserialize, Debug)]
    [serde(untagged)]
    enum WebSocketMessages { 
       Heartbeat(String, String), 
       Trade(String, String, String, String),
       Snapshot(String, Vec&lt;Vec&lt;String&gt;&gt;),
       OtherSnapshot(String, Vec&lt;Vec&lt;String&gt;&gt;)
       Events(Events), 
    }
    
    [derive(Serialize, Deserialize, Debug)]
    [serde(tag = "event")]
    enum Events {
      //  I know I can put structs in here depending on event_names 
    }

I think the Events part will work with using the tag of event to put different structs in the Events enum. My big problem I think is with the two "snapshots", as the structure of both is very similar and yet I don't know how to neatly specify one has arrays of length 3 (price, count amount) and the other has arrays of length 2 (order, level). Also I'm mixing up String and Vec together in the enum which doesn't feel right.

If anyone can give me a tip or help I would really appreciate it.
