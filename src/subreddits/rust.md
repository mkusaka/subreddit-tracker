# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (39/2020)!](https://www.reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ismh8n/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.46]](https://www.reddit.com/r/rust/comments/iix8vw/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/iix8vw/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/htnv19/official_rrust_whos_hiring_thread_for_jobseekers/).

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

COMPANY: *[Company name; please link to your company's website or careers page.]*

TYPE: *[Full time, part time, internship, contract, etc.]*

DESCRIPTION: *[What does your company do, and what are you using Rust for? How much experience are you seeking and what seniority levels are you hiring for? The more details the better.]*

LOCATION: *[Where are your office or offices located? If your workplace language isn't English, please specify it.]*

REMOTE: *[Do you offer the option of working remotely? If so, do you require employees to live in certain areas or time zones?]*

VISA: *[Does your company sponsor visas?]*

CONTACT: *[How can someone get in touch with you?]*
## [3][[Discussion] From C++ to Rust: everything's just better (to me)](https://www.reddit.com/r/rust/comments/j0jlob/discussion_from_c_to_rust_everythings_just_better/)
- url: https://www.reddit.com/r/rust/comments/j0jlob/discussion_from_c_to_rust_everythings_just_better/
---
# Some background

Until 3 months ago, C++ had always been my main language both for work and personal projects. I was told by my friends countless times to give Rust a try but I always thought Rust was nothing but a fancy flavour of  C++. (Which is a VERY wrong thought, shame on myself). Also I stuck with OOP for most of my life, switching to Rust may be painful at the beginning. Besides, my personal projects are mostly games, and I have already familiar with some mainstream frameworks in C++. However most of them haven't got support on Rust yet.

# What's changed?

Coding in C++ can be very painful in some situations (actually in many situations), especially when doing memory management for a fairly large project. Some unintentional bad design may lead to horrible memory related errors. For me, most common one is sharing raw pointers. It doesn't happen often, smart pointers should be used in most cases, but sometimes it's just unavoidable. (E.g. a deep clone is needed for the pointer which is pointing to a super class). Yeah I can implement "Clone" class for it but it's a heavy work when the inheritance tree goes crazy. 

Apparently OOP has a lot of downside in game programming besides inheritance hell, and it becomes worse and worse, and finally it drives me crazy one day. (I know OOP isn't the only way to use C++, but come on, using Data oriented programming in C++ is as hard as coding in C, I think.)

Now, I have pretty many other choices to make my life easier, C# can be one of them, but languages with GC are just not my dishes, (Except PFP languages. Love you Haskell.) even though the language efficiency isn't a big problem in current days for game development. What I want is a language with high efficiency like C/C++ as well as human friendly syntax and memory management system.

# Rust? Rust!

Pattern match, which is the one thing I always want in C++ (I think there isn't a official way to do pattern match in C++? Correct me if I'm wrong), is the first thing that got my attention. In most FP languages pattern match is a must-have, e.g. Haskell, Scala, etc., it makes code much readable and clean. So I decided to give Rust a try just for pattern match. Then I spent a whole weekend going through the tutorial, and man, that's impressive! I can't summarize my surprise how this language is well designed in a few words.

# The syntax is way more comforting than C++'s (to me)

I love pattern match, I love Option, I love using Result to handle errors (Exception is just a nightmare) and I love macros, especially derive! They are too good I don't even know why other languages don't have all of them.

Ok this section is just my opinion, but I really love how beautiful my code can be.

# The compiler is way more smarter than C++'s

C++ compiler is like, raw pointers? Read write a reference at the same time? No problem go ahead I don't care. Then you need to prepare for several hours debugging on multithread code to find out you shouldn't have written the code that way from the beginning.

Rust compiler is like, WHAT ARE YOU DOING? STOP! NO YOU SHOULDN'T DO THAT! YOU WILL REGRET IT! It seems too sensitive about what you are writing at the beginning, but it is  truely reasonable and force you to use a better design on your code. After a while, you will find out how beautiful your code has been.

It's like taking care of your teeth. If you have a very good habit of brushing your teeth and using proper medical goods, you won't need to meet dentist for almost whole life. But it's difficult and most people don't do well. Rust is like a personal dentist and monitor you everyday, force you to follow all steps cleaning teeth. It may be annoying but it just works.

# What now?

I am still getting familiar with Rust since there are still a lot of gimmicks I haven't touched. I'm also making a game with Amethyst. It feels really good how I don't need to worry about , well, a lot of things anymore. Data oriented programming is also fairly new to me but I found it not bad at all while coding in Rust.

I haven't been so excited about a programming language before even when I was learning Haskell, which is a pretty fun experience. I now think I may recommend more people to try Rust at least once. It may not be suitable for everyone but it must be fun for anyone who like a good programming language.

# At Last

I wonder how people in this sub got into Rust? How was your experience while learning it? Am I the only one being so excited about a programming language?
## [4][So you want to live-reload Rust](https://www.reddit.com/r/rust/comments/j0ajdy/so_you_want_to_livereload_rust/)
- url: https://fasterthanli.me/articles/so-you-want-to-live-reload-rust
---

## [5][`#[non_exhaustive]` got me worried about the direction of rust](https://www.reddit.com/r/rust/comments/j0ldal/non_exhaustive_got_me_worried_about_the_direction/)
- url: https://www.reddit.com/r/rust/comments/j0ldal/non_exhaustive_got_me_worried_about_the_direction/
---
First of all, very very glad to see a "practical" language with a Hashell-ish type system and immutable by default (yes, yes, Haskell is practical too :). I think it is one of the "good" languages (others are elm, erlang, haskell, python for what it good at, ocaml?), and I choose to use it in production.

I was reading the RFC of \`#\[non\_exhaustive\]\`, it seems that the main rational behind it is to avoid "breaking" changes when a new variant is added to the enum (in the enum case).

It seems to me that, the "breaking" mentioned in the RFC has a very naive meaning: still compiles after the variant is added, which is far unimportant compared to the other meaning of "breaking": the domain logic is still correct after the new variant is added.

A major role the exhaustive checked plays is that, when a new variant is added, it forces you to go to each use site and check that your implementation (regarding to domain logic and semantics) is still correct. In Elm and Haskell, the wildcard pattern (which you have to use if \`#\[non\_exhaustive\]\`) is discouraged, for it practically disables exhaustive check, which is a valuable tool for large projects and for maintaining projects not written by you (also for projects you wrote long time ago)

If a breaking change is introduced on a library function, you don't want the compiler omits the call to the breaking function, you want the compilation to fail, so that you can handle the change manually. In a ideal world, the compiler will detect semantic changes of a function, (compared to function signature changes), when the semantic of a function changed, your compilation will fail, how awesome would that be? I think, to a certain degree, change-on-enum reflects change-on-semantics.

I think that the compilation error the \`#\[non\_exhaustive\]\` tries to avoid, helps to introduce the worst type of error: error that is hidden (because the wildcard pattern handles the newly added variant, and the newly added variant requires special domain logic to handle and shouldn't be wildcarded, resulting the domain logic is not audited after the new variant is added)

You can argue that we are all grown ups here, and you should be careful when use \`#\[non\_exhaustive\]\`, that is not the point (and you can say the same thing about c++, and not everyone is a grown up), the point is that such an RFC shouldn't be passed, for it introduces bad practice, and de-sound the compiled program. I think one big reason that people like languages such as elm and haskell is for the confidence they have on their programs, if you finish a program, you finish a program, you can go to sleep knowing that your program will behave the way you wanted
## [6][A tale of two libcs](https://www.reddit.com/r/rust/comments/j0puje/a_tale_of_two_libcs/)
- url: https://drewdevault.com/2020/09/25/A-story-of-two-libcs.html
---

## [7][Dicetest, an alternative to proptest and quickcheck](https://www.reddit.com/r/rust/comments/j0pjdo/dicetest_an_alternative_to_proptest_and_quickcheck/)
- url: https://www.reddit.com/r/rust/comments/j0pjdo/dicetest_an_alternative_to_proptest_and_quickcheck/
---
[https://github.com/jakoschiko/dicetest](https://github.com/jakoschiko/dicetest)

I'm using ScalaCheck at work and don't quite like it. It's nice for testing functional code, but that's not enough. Scala is both, functional and imperative. Many of my tests look like this: Mutate some state, make assertions, mutate some state, make assertions... In these cases I just want to write simple unit tests with random test data instead of inventing elegant properties. It just feels clumsy to use ScalaCheck for this.

Then I started to learn Rust and more of my code became stateful. I was curious if I could implement a less functional property testing library that I would enjoy to use. I think I succeeded. Writing the library was fun, too, and it helped me to understand Rust. E.g. it was very enlightening to realize that I need two generator traits, similar to Fn and FnOnce, if I don't want to clone or reference count everything.

There are obvious similarities to proptest and quickcheck. They have shrinking, but I never liked shrinking in ScalaCheck. Instead, I'm not forcing the user to generate all test data at the beginning, which I prefer.

What do you think?
## [8][Towards principled reactive UI](https://www.reddit.com/r/rust/comments/j0840q/towards_principled_reactive_ui/)
- url: https://raphlinus.github.io/rust/druid/2020/09/25/principled-reactive-ui.html
---

## [9][[Academic] Open Source Development Survey (Software Developers using GitHub)](https://www.reddit.com/r/rust/comments/j0q584/academic_open_source_development_survey_software/)
- url: /r/SampleSize/comments/iybrqr/academic_open_source_development_survey_software/
---

## [10][SublimeText rust-enhanced: Is it possible to show the filename where the error occurs?](https://www.reddit.com/r/rust/comments/j0ozb4/sublimetext_rustenhanced_is_it_possible_to_show/)
- url: https://www.reddit.com/r/rust/comments/j0ozb4/sublimetext_rustenhanced_is_it_possible_to_show/
---
Hello,  
Rust Enhanced works very well, but on larger projects I've some trouble to find the file where an error occurs. In my bottom bar, Rust Enhanced only shows me that there is an error, but not the filename. I've to say I'm relying on the "check on save" function.

https://preview.redd.it/ix2wlyr04op51.png?width=348&amp;format=png&amp;auto=webp&amp;s=bc352410ab1b0adfb8cbde971ee975e86e55070a

Is it possible to highlight the file in the sidebar, similar to the git-plugin which will draw (in my case) a simple dot to on the files I've modified. Or at least show the file names where the error occurs in the bottom bar.

Maybe it's a setting I've missed?

Thanks for help!

PS: I posted my question on Github without any luck to reach someone

x-post: [https://github.com/rust-lang/rust-enhanced/issues/446](https://github.com/rust-lang/rust-enhanced/issues/446)
## [11][Unimpressed with lolcat's startup performance, I've rewritten it in Rust - Introducing dotacat!](https://www.reddit.com/r/rust/comments/j0aaus/unimpressed_with_lolcats_startup_performance_ive/)
- url: https://crates.io/crates/dotacat
---

## [12][Type checking process.exit()](https://www.reddit.com/r/rust/comments/j0qyvp/type_checking_processexit/)
- url: https://www.reddit.com/r/rust/comments/j0qyvp/type_checking_processexit/
---
I am working through the introductory rust book, and got to the following example.

    let config = Config::new(&amp;args).unwrap_or_else(|err| {
        println!("Problem parsing arguments: {}", err);
        process::exit(1);
    });

unwrap_or_else expects a closure of type `&amp;str -&gt; Config`, but our closure is of type  `&amp;str -&gt; ()`, so how does the compiler handle this edge case?
