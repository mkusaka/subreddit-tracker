# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (41/2020)!](https://www.reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j1jgum/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 359](https://www.reddit.com/r/rust/comments/j71tq7/this_week_in_rust_359/)
- url: https://this-week-in-rust.org/blog/2020/10/07/this-week-in-rust-359/
---

## [3][Built a tiny web tool with RustWasm](https://www.reddit.com/r/rust/comments/j8ht8a/built_a_tiny_web_tool_with_rustwasm/)
- url: https://www.reddit.com/r/rust/comments/j8ht8a/built_a_tiny_web_tool_with_rustwasm/
---
I'm still quite new to the language but I have found myself rather infatuated with RustWasm lately and I went ahead to create a simple [random password generator](https://wasm-pass.gq) with wasm-bindgen, rand and React.js. Feel free to give it a try, open to feedback and contributions.
## [4][A New Backend for Cranelift, Part 1: Instruction Selection – Mozilla Hacks](https://www.reddit.com/r/rust/comments/j8aqbs/a_new_backend_for_cranelift_part_1_instruction/)
- url: https://hacks.mozilla.org/2020/10/a-new-backend-for-cranelift-part-1-instruction-selection/
---

## [5][Memory Safe ‘curl’ for a More Secure Internet](https://www.reddit.com/r/rust/comments/j7yegb/memory_safe_curl_for_a_more_secure_internet/)
- url: https://www.abetterinternet.org/post/memory-safe-curl/
---

## [6][RustDesk: a remote desktop software written with Rust](https://www.reddit.com/r/rust/comments/j8dl4p/rustdesk_a_remote_desktop_software_written_with/)
- url: https://www.producthunt.com/posts/rustdesk
---

## [7][How To Send Keyboard Events To Another Process In Rust ?](https://www.reddit.com/r/rust/comments/j8fp23/how_to_send_keyboard_events_to_another_process_in/)
- url: https://www.reddit.com/r/rust/comments/j8fp23/how_to_send_keyboard_events_to_another_process_in/
---
In a game I play, there is a macro assigned to the key `4` and I have written a small rust program that uses [enigo lib](https://github.com/enigo-rs/enigo) to simulate the key press. When I run the program in cmd, I can see that `4` key is being pressed repeatedly (once a second) because it's written onto the console window, but when I switch to game process, while the rust code is running, it doesn't work, and windows says my rust program was blocked from making unwanted changes.

How is it that remote-control/conference applications like team-viewer/zoom etc. is able to send key clicks to another processes if windows is blocking them ?

EDIT: I'm able to send key events to the sublime text, but not to the game. But zoom is able to control the game as well.
## [8][Building entire toolchain from source](https://www.reddit.com/r/rust/comments/j8j454/building_entire_toolchain_from_source/)
- url: https://www.reddit.com/r/rust/comments/j8j454/building_entire_toolchain_from_source/
---
Hey, is there any way to use rustup or a tool like it to manage and build the entire toolchain from source? 

I can't use my default rustup toolchains, as the powerpc64-linux binaries use a different linker to my system and segfault.
## [9][Help needed: Moving a buffer and a Cow pointing to it in a struct.](https://www.reddit.com/r/rust/comments/j8irng/help_needed_moving_a_buffer_and_a_cow_pointing_to/)
- url: https://www.reddit.com/r/rust/comments/j8irng/help_needed_moving_a_buffer_and_a_cow_pointing_to/
---
Hey everyone,

I'm currently using the quick\_xml crate to parse an XML file. Being a pull parser, the central API looks like:  

```
let mut buffer = Vec::new();
let event = reader.read_event(&amp;mut buffer);

Element { buffer, event } // &lt;- this fails as buffer is mutable borrowed...
```

Now once the expected event occurs, I'd like to return a struct in which I move the buffer and the event (which internally contains a Cow which lifetime is bound to the buffer). However the borrow checker doesn't like this approach, as it considers buffer as mutably borrowed.

For completeness, the signature of read_event is: `pub fn read_event&lt;'a, 'b&gt;(&amp;'a mut self, buf: &amp;'b mut Vec&lt;u8&gt;) -&gt; Result&lt;Event&lt;'b&gt;&gt;`

Therefore packing the buffer and the resulting event together should be technically safe.

I hope I'm explaining enough to understand my problem. Is there a safe way to achive this in Rust? (without giving up the zero-copy capabilities). 

many thanks in advance :)
## [10][rust in curl with hyper](https://www.reddit.com/r/rust/comments/j7yw4o/rust_in_curl_with_hyper/)
- url: https://daniel.haxx.se/blog/2020/10/09/rust-in-curl-with-hyper/
---

## [11][Introducing instant.bible - An as-you-type Bible search engine written in Rust, Swift, Kotlin, and TypeScript!](https://www.reddit.com/r/rust/comments/j8anec/introducing_instantbible_an_asyoutype_bible/)
- url: https://knpw.rs/blog/instant-bible
---

## [12][Are we Bluetooth yet?](https://www.reddit.com/r/rust/comments/j81m4z/are_we_bluetooth_yet/)
- url: https://www.reddit.com/r/rust/comments/j81m4z/are_we_bluetooth_yet/
---
Hello everyone, I have finally the chance to use Rust at work but I'm kinda stumbling on an issue.

I'm starting the development of an application (I'd be pretty much the only developer and I can chose what to use, as long as it makes sense) which lets the user interact with a Bluetooth device my company made and I wanted to use Flutter strictly for the GUI side and Rust for the application logic (interfacing with the device, handling access to the filesystem, network calls and, eventually, some machine learning stuff either in the back end or the app itself, but TBD). The target platform would mainly be Android, but the goal in the long run would also be to support iOS and desktop (Windows mainly).

Well, so far so good, but surprisingly enough when I was looking for Bluetooth crates I didn't find anything that seems usable, either because unmaintained (last update years ago) or incomplete (not supporting the platforms I need). I didn't find any recent post talking about it here and no reference in [Awesome Rust](https://github.com/rust-unofficial/awesome-rust) or [Not Yet Awesome Rust](https://github.com/not-yet-awesome-rust/not-yet-awesome-rust). I was looking for something like [QT's Bluetooth](https://doc.qt.io/qt-5/qtbluetooth-index.html).

&amp;#x200B;

I'm posting here hoping you might have some good suggestions on the matter and convince me not to dump Rust already for this use case. Thanks in advance for your support.
