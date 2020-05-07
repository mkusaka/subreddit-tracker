# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (19/2020)!](https://www.reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 337](https://www.reddit.com/r/rust/comments/geagy0/this_week_in_rust_337/)
- url: https://this-week-in-rust.org/blog/2020/05/05/this-week-in-rust-337/
---

## [3][Buttplug Sex Toy Control Project](https://www.reddit.com/r/rust/comments/gezz84/buttplug_sex_toy_control_project/)
- url: https://buttplug.io/
---

## [4][Look who made it front and center for a new Microsoft/GitHub announcement!](https://www.reddit.com/r/rust/comments/geoaf8/look_who_made_it_front_and_center_for_a_new/)
- url: https://i.redd.it/3n5mzytqh6x41.png
---

## [5][KAS GUI 0.4 release](https://www.reddit.com/r/rust/comments/gf42ip/kas_gui_04_release/)
- url: https://github.com/kas-gui/kas
---

## [6][Embedded async/await in the real world: The Polymer Keyboard](https://www.reddit.com/r/rust/comments/geuzk0/embedded_asyncawait_in_the_real_world_the_polymer/)
- url: https://www.reddit.com/r/rust/comments/geuzk0/embedded_asyncawait_in_the_real_world_the_polymer/
---
I've had this in the works for far too long. It stalled a bit when I took a ~8 month break from it, but it's finally complete(ish)! At any rate, it's complete enough to have supplanted the ErgoDox as my daily driver.

Overall, my experience has been extremely positive! Being able to abstract the core logic over familiar `Read`, `Write`, `Sink`, and `Stream` traits helps immensely for being able to test and reason about the code. I suspect that my IO/Timer Future abstractions have some flaws, but they work well enough until someone who actually knows what they're doing builds some better ones. The beauty of having an abstract core is that it should allow for a drop-in replacement of the IO backend!

Overview blog post:
[The Rest of the Keyboard](https://josh.robsonchase.com/rest-of-the-keyboard)

It's a little light on actual Rust stuff, but I got a little overwhelmed with everything I could have gone into more detail about.
## [7][has anyone used this (rust notebooks) for projects in the real world?](https://www.reddit.com/r/rust/comments/gf28is/has_anyone_used_this_rust_notebooks_for_projects/)
- url: https://shahinrostami.com/posts/programming/rust-notebooks/setup-anaconda-jupyter-and-rust/
---

## [8][Rust Analyzer is now in official Arch repos](https://www.reddit.com/r/rust/comments/gejsho/rust_analyzer_is_now_in_official_arch_repos/)
- url: https://www.reddit.com/r/rust/comments/gejsho/rust_analyzer_is_now_in_official_arch_repos/
---
I'd been maintaining Rust Analyzer AUR (Arch User Repository) packages since the [first binary release](https://www.reddit.com/r/rust/comments/eo5bs2/rustanalyzer_changelog_7_with_binary_releases/fe9gxu2) in January, but since today it's in official repos:

&gt; polyzen [1] filed a deletion request for rust-analyzer [2]:

&gt; This package has been added to the repos. Thank you.

&gt; [1] https://aur.archlinux.org/account/polyzen  
&gt; [2] https://aur.archlinux.org/pkgbase/rust-analyzer

To install it or replace AUR package:

    pacman -S rust-analyzer

and you're good to go.

---

edit: [relevant PR](https://github.com/rust-analyzer/rust-analyzer/pull/4333)
## [9][my first time programming ever...question about something that doesnt matter...](https://www.reddit.com/r/rust/comments/gf4nu1/my_first_time_programming_everquestion_about/)
- url: https://www.reddit.com/r/rust/comments/gf4nu1/my_first_time_programming_everquestion_about/
---
so, i made my first attempt at learning programming...after i successfully compiled and ran it, i wondered how big the compiled application actually was...so i just made an empty .rs file and tried to compile it, and it didnt work...so i just did this:

fn main () {

}

&amp;#x200B;

and after i compiled it, it was 2.6 Mb...my question is, what does a compiled application actually have in it? like, did rust compile nothing, and it has built in safechecks and error codes that it spits out even if nothing is in the actual code i wrote? just curious...thanks!!!
## [10][color-eyre: a colorful error reporter based on color-backtrace and eyre is now available on crates.io](https://www.reddit.com/r/rust/comments/gf69mv/coloreyre_a_colorful_error_reporter_based_on/)
- url: https://crates.io/crates/color-eyre
---

## [11][Many of Herb Sutter's C++ talks are surprisingly relevant to Rust](https://www.reddit.com/r/rust/comments/geqfhi/many_of_herb_sutters_c_talks_are_surprisingly/)
- url: https://www.reddit.com/r/rust/comments/geqfhi/many_of_herb_sutters_c_talks_are_surprisingly/
---
I don't write much C++ myself, but I find that listening to Herb Sutter's talks about the language's evolution (past, present, and future) is extremely informative. What C++ I know, I learned more than a decade ago. It turns out that a lot of things have changed since then, often in ways that bear a striking similarity to Rust. I wanted to make a list of some highlights that I've noticed:

**[Essentials of Modern C++ Style (2014)](https://youtu.be/xnqTKD8uD64)**: There are several parallels to Rust in this talk. At 27m45s, there's a slide that explicitly discusses "lifetimes" and "ownership". There's also a long discussion of move semantics, which works somewhat differently in C++ but addresses similar problems.

**[Making Exceptions More Affordable and Usable (2019)](https://youtu.be/ARYP83yNAWk)**: Discusses a proposal for a new exception throwing mechanism. These errors would be statically typed, and at 1h4m35s there's a slide describing a hypothetical explicit `try` syntax that would visually highlight function calls that could fail. Sutter explicitly calls out the similarity to Rust's `Result` type, though he doesn't mention the `?` operator. There's also a long discussion of how memory allocation failure is handled in C++, and how that might change in the future, which has also been a tricky topic in Rust since before 1.0.

**[Writing Good C++14... By Default (2015)](https://youtu.be/hEx5DNLWGgA)**: Discusses a proposal for defining "safe" and "unsafe" subsets of C++. At 41m42s there's a demo of a static analysis tool that's very similar to Rust's borrow checker, and then there's a high level discussion of lifetime inference and annotations.

**[Atomic Weapons (2012)](https://youtu.be/A8eCGOqgvH4)**: Rust atomics and C++ atomics are very similar. This talk is a deep dive into what the different atomic orderings mean and why they're defined the way they are. If you hear the term "memory model" and feel like you kinda sorta know what that means, but maybe not very concretely, this talk is extremely clarifying.
## [12][Init large array of struct and mutable static](https://www.reddit.com/r/rust/comments/gf5eww/init_large_array_of_struct_and_mutable_static/)
- url: https://www.reddit.com/r/rust/comments/gf5eww/init_large_array_of_struct_and_mutable_static/
---
Hello everyone,  
with the purpose of learning rust I've decided to develop a simple address book but I'm facing some issues and I'd like to know which can be a good approach to write this code.  
Please see my code below:

    struct Contact {
        name: String,
        id: u32,
        phone: u32,
    }
    
    impl&lt;'a&gt; Contact {
        fn new() -&gt; &amp;'a Contact {
            let c = &amp;mut address_book[idx];
    
            c.id = id as u32;
            c.phone = 0123456789;
            id += 1;
            c
        }
    }
    
    static mut address_book: [Contact; 2] = [
        Contact {name:"Unknown".to_string(), id:0, phone:0},
        Contact {name:"Unknown".to_string(), id:0, phone:0},
    ];
    static mut id: usize = 0


I've declared a simple Contact structure with few fields, declared an array of contacts and a function that should update some fields of a contact through its reference and return it.

What's unclear to me is how I can declare a large array without initializing every element.  
I was expecting to be able something like:

    static mut address_book: [Contact; 128] = [
        Default::default()
    ]

The only solution I found is to use the \`Vec&lt;T&gt;\` type but I'd prefer to allocate this array on the stack.

Another issue is that, as the compiler says, the use of mutable static is unsafe and require an unsafe block:

    error[E0133]: use of mutable static is unsafe and requires unsafe function or block
      --&gt; src/main.rs:10:35
       |
    10 |         let c = &amp;mut address_book[id];
       |                                   ^^^ use of mutable static
       |
       = note: mutable statics can be mutated by multiple threads: aliasing violations or data races will cause undefined behavior
    

I think that I should use a box here, but not sure becase I haven't tried yet.

I'm looking forward to reading your clarifications, thanks!
