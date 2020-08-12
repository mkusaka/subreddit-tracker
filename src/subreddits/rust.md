# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (33/2020)!](https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (33/2020)?](https://www.reddit.com/r/rust/comments/i6yras/whats_everyone_working_on_this_week_332020/)
- url: https://www.reddit.com/r/rust/comments/i6yras/whats_everyone_working_on_this_week_332020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-33-2020/47095?u=llogiq)!
## [3][Rd: A port of mozilla/rr to the Rust programming language](https://www.reddit.com/r/rust/comments/i8bmgq/rd_a_port_of_mozillarr_to_the_rust_programming/)
- url: https://github.com/sidkshatriya/me/blob/master/001-rd-intro.md
---

## [4][gitoxide 0.3 - clone from disk at 1.7x the speed of git*](https://www.reddit.com/r/rust/comments/i8ccqd/gitoxide_03_clone_from_disk_at_17x_the_speed_of/)
- url: https://github.com/Byron/gitoxide/releases/tag/v0.3.0
---

## [5][cargo-fuzzcheck 0.4.0 : now with a proc_macro to make custom structs/enums automatically fuzzable, and many other improvements](https://www.reddit.com/r/rust/comments/i8ck9h/cargofuzzcheck_040_now_with_a_proc_macro_to_make/)
- url: https://www.reddit.com/r/rust/comments/i8ck9h/cargofuzzcheck_040_now_with_a_proc_macro_to_make/
---
Hi everyone!

A while ago I posted here to announce that I released a [structure-aware &amp; coverage-guided fuzzing engine written in Rust](https://github.com/loiclec/fuzzcheck-rs). I also asked a few questions about how I could go forward with the project.

Since then, I have set up a GitHub Sponsors [page](https://github.com/sponsors/loiclec) and Embark Studios was kind enough to sponsor me, which has helped the project a lot. So thank to them!

A few of you suggested projects I could fuzz-test. I did it shortly for a couple of them to profile fuzzcheck’s performance and got some huge gains thanks to it. But I am not ready to create proper benchmarks yet, due to a lack of resources and time.

I have also improved the corpus shrinking algorithm, created a configuration file to allow greater control of the instrumentation used, and added preliminary support for stack-depth-guided fuzzing.

Finally, I was also told providing a proc_macro to generate mutators would help drive adoption. With the release of fuzzcheck 0.4.0, that proc_macro is finally there!

Now you can write things like:

    #[fuzzcheck_derive_mutator(DefaultMutator)]
    #[derive(Clone, Serialize, Deserialize)]
    pub enum SampleEnum&lt;T&gt; {
        A(u8),
        B { x: bool , y: Option&lt;T&gt; },
        C,
        D
    }

    #[fuzzcheck_derive_mutator(DefaultMutator)]
    #[derive(Clone, Serialize, Deserialize)]
    pub struct SampleStruct&lt;A, B, C&gt; {
        a: A,
        b: Vec&lt;B&gt;,
        c: C
    }

to make these types fuzz-able. This means that, finally, fuzzcheck is approaching `cargo-fuzz`’s ease of use! I am very excited about it, even though it is not very polished yet. I have carried some initial tests, and it seems so far that fuzzcheck’s structure-aware mutators are much, much better and faster than what cargo-fuzz can provide with the `Arbitrary`trait. On many very simple tests, fuzzcheck would find a solution almost instantly while cargo-fuzz was still running 1 minute later. Of course, it doesn't mean that it will consistently outperform it, but it is encouraging nonetheless.

If you'd like to try it, it is easy and only takes a couple minute to run the sample fuzz-test. Go to https://github.com/loiclec/fuzzcheck-rs and follow the usage instructions :)

That's it for the update. There is still so much to do, but for the rest of the month, I think I might write a “book” explaining the design as well the work that remains to be done in details. 

Thanks for reading! I will be in the comments to answer questions if you have some :)
## [6][How do Mozilla layoffs affect Rust?](https://www.reddit.com/r/rust/comments/i7stjy/how_do_mozilla_layoffs_affect_rust/)
- url: https://www.reddit.com/r/rust/comments/i7stjy/how_do_mozilla_layoffs_affect_rust/
---
Mozilla sadly announced a [large round of layoffs](https://blog.mozilla.org/blog/2020/08/11/changing-world-changing-mozilla/) today. Rust has grown beyond just Mozilla but certainly still has close ties to it. Any immediate sense of the impact of this announcement on people working on Rust or Rust projects at Mozilla? To anyone who was impacted, I’m sorry, and I hope you’re able to land on your feet soon.
## [7][unsure of whether to prototype in python first](https://www.reddit.com/r/rust/comments/i8cbe6/unsure_of_whether_to_prototype_in_python_first/)
- url: https://www.reddit.com/r/rust/comments/i8cbe6/unsure_of_whether_to_prototype_in_python_first/
---
hi rustaceans!

i’m new to rust, coming from a python and java background with very little system-level language experience, just for some context. i’m halfway through TRPL atm.

right now i’m planning to start work on a curses-based REST web api client that stores data in an sqlite db. i felt like it would be a good opportunity to practice coding in rust, instead of just reading the book.

however, coming from python, the rust syntax is downright terrifying, and i’m fairly sure i’m gonna spend more time agonizing over syntax than actually working on the business logic.

i’m debating whether i should prototype the app in python before i actually implement it in rust, and work on TRPL’s grep imitation as practice alongside the python prototype. but, i’m impatient and i plan to release a working version of the app by january next year. also, the curses libraries i’m using for python and rust implementations are wildly different and would require learning both in depth. i also have an idea of how the backend logic already works, so it’s not like i need to try my sqlite database or my api request handling anyway.

if you could help me through this dilemma, i’d appreciate it. additionally, if you could help me with a rust-like (like pythonic) way of structuring the app using things like lifetimes, traits and generics, i’d appreciate it.

thanks!

my current dependencies:

- cursive (for the tui)

- reqwest (api querying)

- serde (json deserialization)

- rusqlite (database)

- indicatif (progressbar)
## [8][Building just rust crates for arm](https://www.reddit.com/r/rust/comments/i8cf4q/building_just_rust_crates_for_arm/)
- url: https://www.reddit.com/r/rust/comments/i8cf4q/building_just_rust_crates_for_arm/
---
Hi

I'm cross-building and packaging rust for our mobile/linux-based OS. I'm very new to rust as a language so excuse my misuse of terms.

I want to find an incantation to just build the arm target version of the std/core crates that come with rust. I do not need an arm-hosted rustc. We rebuild the entire OS regularly using a dependency cycle and reproducable bit-identical packages. Because these builds happen so much we try to only build what we need. This is the high-level reason we don't want to build an arm rustc.

My assumption is that given an x86 rustc binary (--print target-list has armv7) and no ARM binary crates, I can build the ARM crates from scratch from source. If this is wrong then I can stop looking :)

(aside: We use an unusual cross-building system (which works well with rustc). Executables run on x86 hw in a qemu chroot. For build-time performance we use low-level system call intercepts to invoke x86-hosted binaries to bypass emulation. It's not important but it's the reason I don't need an arm rustc compiler.)

The build system will build x86 rust (with stages) and crates just fine. It also builds (very slowly, using emulation) arm rust and crates. I can also build arm rust binaries using the x86-host rustc.

So this is my incantation of ./configure and ./x.py :

    ./configure --program-prefix= --prefix=/usr --exec-prefix=/usr --bindir=/usr/bin --sbindir=/usr/sbin --sysconfdir=/etc --datadir=/usr/share --includedir=/usr/include --libdir=/usr/lib --libexecdir=/usr/libexec --localstatedir=/var --sharedstatedir=/var/lib --mandir=/usr/share/man --infodir=/usr/share/info \
      --disable-option-checking \
      --libdir=/usr/lib \
      --build=armv7-unknown-linux-gnueabihf --host=armv7-unknown-linux-gnueabihf --target=armv7-unknown-linux-gnueabihf \
      --python=python3 \
      --local-rust-root=/usr \  &lt;&lt; this is actually an x86 rustc
      --enable-local-rebuild \
      --enable-llvm-link-shared \
      --enable-optimize \
      --disable-docs \
      --disable-compiler-docs \
      --disable-rpath \
      --disable-codegen-tests \
      --debuginfo-level=0 --debuginfo-level-std=2 --disable-debuginfo \
      --disable-debuginfo-only-std --disable-debuginfo-tools --disable-debuginfo-lines \
      --enable-extended \
      --enable-vendor \
      --llvm-root=/usr/ \
      --enable-parallel-compiler

Then the x.py to just build the crates using the:stage 0 (system) rustc:

      python3 ./x.py build src/libstd --stage 0
      python3 ./x.py build src/libcore --stage 0

This however seems to fail looking with `can't find crate for \`std\``
## [9][Am I over-using traits?](https://www.reddit.com/r/rust/comments/i83kzn/am_i_overusing_traits/)
- url: https://www.reddit.com/r/rust/comments/i83kzn/am_i_overusing_traits/
---
Hello all,

I've recently started with Rust, coming from a Scala background. My natural tendency when coding is to write most functions to be as generic as possible - and this means that both the arguments to and the return types from most functions are trait instances. I tend to start a project by having a single file which lays out the skeleton of the project in terms of traits. I initially approached rust in this way, but I'm increasingly wondering if this is the wrong attitude for the language, and I need to start thinking in less of an FP-type way.

A few clues that I'm doing strange stuff is the slightly awkward syntax for dealing with trait objects (`&amp;dyn Trait` etc), and the fact that there is no syntactic sugar for deriving traits from other traits. I also understand how both trait objects and closures are abstractions that do not necessarily fit well in a low-level language with quite granular memory management. So before I spend too long trying to make a square peg fit in a round hole, I thought I'd ask if people thought I should just replace most of my trait objects with struct types, and shift a little away from the functional style of coding?

Any thoughts gratefully received.
## [10][Rust monolith as micro service killer](https://www.reddit.com/r/rust/comments/i8dg35/rust_monolith_as_micro_service_killer/)
- url: https://www.reddit.com/r/rust/comments/i8dg35/rust_monolith_as_micro_service_killer/
---
Greetings! I wanted to share a bit of our current story and some thoughts of the architecture I'm pursuing to hear the opinion of the community about possible limitations, areas of improvement or any random comment you guys have. I'm kind of having a bit of impostor syndrome trying to find reassurance in the decisions I'm proposing as it's hard to tell if in a 5 year time it was worth making those choices or a complete nightmare, I feel it fits well my team's and product needs but since it goes a bit against common practices is intimidating and hard to always look all cool and secure when proposing this changes.

My team and I are working on a fintech product in its early days that needs to be very reliable and secure as it will be dealing with user's real money. I'm responsible mostly for the system's security which in our early stages mostly means making design decisions that will keep our attack surface as small as possible by making certain attack vectors impossible by design. Big part of it is the use of Rust which the team is already onboard with and started learning already, other design decisions include for the longer term "blockchainizing" and decentralizing our product.

Now comes the monolith vs micro-service discussion. Typical story here, like with most start-ups the company started with its rushed MVP quality monolith that grew and lasted longer than it should have, with the many mistakes that curse monoliths like bad coupling among components, not well defined domains and dependencies, shared global state, etc, then it transitioned(perhaps to early) to  a micro-service kubernetes based architecture like all "cool kids" do nowadays, sadly and not surprisingly implemented in not the best way, with still a lot of coupling plus the many not well handled complexities of the new infrastructure and it's elevated costs. Now this annoying guy🖐🏽 comes in trying to defy common knowledge and suggesting that perhaps is better to go back to a monolith ... wait WHAAAT!?

First I think it was too early to think of having multiple services, the team didn't(and still doesn't) have the resources and expertise needed to deal with the many complexities of micro services, specially in the operational/infra side, and now we suffer from the expected issues of hard to debug bugs with no logging or tracing infrastructure, weak CI/CD pipeline, no end-to-end tests, etc.... but but kubernetes and scalability and and X startup did Y ... yeah but did I mention everything still has been running on one single EC2(expensive) instance, that there's one container per service at a time and traffic is in the tens or less req/s not hundreds? why do we even need k8s and it's learning curve? do we need the many data-bases? Doesn't this look like just a bigger more complex monolith?  

Now with Rust already being acknowledged as a reality and necessity for the core of our product I propose to extend its reach to basically everything, one mono repo, one service service to rule them all. My rationale and assumption being, a Rust monolith "done well" can scale just as good or better, scaling not by adding more metal to the mix but by making things faster and simpler. It's not only about the time we cut form a node.js service handling some logic vs what Rust takes in comparison but also the network round trips, (de)serialization or data base look-ups(sled? embedded db dangerous?) etc, I think we can shave up lots of precious CPU cycles there! 

Then there's still the issue of scalability on the code base level, how to grow a healthy code base and avoid things getting all tangled and coupled together again, minimize global mutable state and retain simplicity? For this I think restricting developer's "creativity" is a good approach, to avoid foot shoots by following well known design patterns. The first building block I thought about was an actor system that could help with reliable delivery of messages between components, and keeping things running, also to enforce go's mantra of sharing by communicating and not communicating by sharing memory(avoid global state at least among the main components of the system). Didn't find a lot of offer in the actor frameworks department but started to play with [Riker](https://riker.rs/) which seemed simple enough and not too focused on HTTP+REST like actix, could be more mature but so far is proving itself useful.  Then to further "restrict creativity" started to play with an [event sourcing+CQRS abstraction](https://github.com/Valiuapp/riker-es) on top of the actor framework hoping it can be a good balance of flexibility with "there should be just one obvious way to do things". When it comes to storing data, I'm very tempted to have a [Sled](https://github.com/spacejam/sled) backed, embedded event store to align with my future hopes of decentralization and the whole system being a single binary, double-click kind of thing that others can easily deploy and run in their own machines to serve a few hundreds or thousands of customers. It's well received in the company by now that our future might look more as like decentralized organization, focusing mostly on a specific geography and letting others like us pop up in other parts of the world to build a bigger whole. Am I being to idealistic thinking some young libraries can replace a more traditional architecture? specially on the DB side not using one of the (no)SQL DBs that we all know and love. How much of a risky move is it to build on top of experimental tech for a thing that is built to last? How far can this get us before we really need more instances? How long can vertical scaling take me? What if management says a couple of years down the line that we are not betting on decentralization, we need to serve millions of users on our own, shouldn't be that hard to break down the monolith down the line but there might be troubles I can't see?  What other bottle necks am I likely to find? Am I trusting Rust's super powers too much? 🤷🏻‍😅 

Lot's of uncertainty mixed with excitement, I feel is a good plan specially to achieve the goal of a secure system with minimized attack surface standing on the shoulders of giants ... but let's see what surprises await in the future 😱
## [11][Is this a bug?](https://www.reddit.com/r/rust/comments/i8c3gc/is_this_a_bug/)
- url: https://www.reddit.com/r/rust/comments/i8c3gc/is_this_a_bug/
---
So I was playing around in [https://godbolt.org/](https://godbolt.org/) and I tried to write an extremely simple program in some different languages to compare how well the different optimization levels held up. Long story short, I wrote this:

`pub fn main(stringlist: Vec&lt;String&gt;, name: String){`  
 `for s in stringlist {`  
 `if s == name {`  
 `println!("Found one!");`  
        `}`  
    `}`  
`}`

At first I tried -C opt-level=3 for speed, then I tried -C opt-level=s for size. Well, unfortunately for me, opt-level=s was larger than opt-level=3 for some reason. Not just by a little bit either, it was almost twice as long at 220 lines vs 416. What's going on here? This isn't exactly a complex piece of code, what is it getting stuck on? I know some assembly instructions take a few more or less bytes, and there might be some other minor things making this not a perfect measurement, but that's a big enough disparity that I can't imagine 3 overtaking s once assembled.
## [12][Rewriting existing functionalities of Mender using Rust](https://www.reddit.com/r/rust/comments/i89z49/rewriting_existing_functionalities_of_mender/)
- url: https://mender.io/blog/a-whirlwind-tour-of-the-mender-client-architecture-using-rust
---

