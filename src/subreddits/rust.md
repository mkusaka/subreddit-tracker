# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (3/2020)!](https://www.reddit.com/r/rust/comments/eo6pjy/hey_rustaceans_got_an_easy_question_ask_here_32020/)
- url: https://www.reddit.com/r/rust/comments/eo6pjy/hey_rustaceans_got_an_easy_question_ask_here_32020/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

The Rust-related IRC channels on irc.mozilla.org (click the links to open a web-based IRC client):

 - [#rust](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust) (general questions)
 - [#rust-beginners](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-beginners) (beginner questions)
 - [#cargo](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23cargo) (the package manager)
 - [#rust-gamedev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-gamedev) (graphics and video games, and see also [/r/rust_gamedev](https://www.reddit.com/r/rust_gamedev))
 - [#rust-osdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-osdev) (operating systems and embedded systems)
 - [#rust-webdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-webdev) (web development)
 - [#rust-networking](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-networking) (computer networking, and see also [/r/rust_networking](https://www.reddit.com/r/rust_networking))

Also check out [last week's thread](https://reddit.com/r/rust/comments/ekpqp7/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 321](https://www.reddit.com/r/rust/comments/epinfr/this_week_in_rust_321/)
- url: https://this-week-in-rust.org/blog/2020/01/14/this-week-in-rust-321/
---

## [3][The Soundness Pledge](https://www.reddit.com/r/rust/comments/eqcefv/the_soundness_pledge/)
- url: https://raphlinus.github.io/rust/2020/01/18/soundness-pledge.html
---

## [4][A sad day for Rust](https://www.reddit.com/r/rust/comments/eq11t3/a_sad_day_for_rust/)
- url: https://words.steveklabnik.com/a-sad-day-for-rust
---

## [5][actix_web repository cleared by author who says he’s done with open source](https://www.reddit.com/r/rust/comments/epzukc/actix_web_repository_cleared_by_author_who_says/)
- url: https://www.reddit.com/r/rust/comments/epzukc/actix_web_repository_cleared_by_author_who_says/
---
He tweeted [“I am done with open source.”](https://twitter.com/fafhrd91/status/1218135374339301378) and has moved `actix-web` and `actix-net` to his personal account with a postmortem left in the original repo: https://github.com/actix/actix-web. I’ll leave the quoted text here as well:

&gt; ##Actix prject postmortem

&gt; Another day, another "unsafe shitstorm”, I guess I get used to it.

&gt; It is interesting how easy to move comment out of context and how hard to comment with very clear intention (especially if you are not native speaker) What was the patch? It was very strait forward, simple, uncreative change, intention was just to remove unsafe not to fix existing code. I believe software development is one of the most creative work we do, and creativity is part of why we love software development, why it is fun. Especially if you combine it with real world projects constraints. “creative constrains” could be source of very interesting solutions. Being on the edge of your abilities is super fun. So uncreative change felt boring (oh! And author gave up copyright claims for that patch (a bit irony and sarcasm)). I’ve never used unsafe unintentionally, I use it because I believe usage is safe. There was no any malicious intentions. I believed it held mutable aliasing invariant and I was very happy that someone found real problem. I wanted to solve the problem, just with a bit of creativity. And use RefCell solution only if it would be not possible to solve it with any other way. Btw, I like the solution I found, it is in master and solves the problem at least one from the issue. If you want to push boundaries you have to touch this boundaries and sometimes you push too hard.

&gt; Be a maintainer of large open source project is not a fun task. You alway face with rude and hate, everyone knows better how to build software, nobody wants to do home work and read docs and think a bit and very few provide any help. Seems everyone believes there is large team behind actix with unlimited time and budget. (Btw thanks to everyone who provided prs and other help!) For example, async/await took three weeks 12 hours/day work stint, quite exhausting, and what happened after release, I started to receive complaints that docs are not updated and i have to go fix my shit. Encouraging. You could notice after each unsafe shitstorm, i started to spend less and less time with the community. You felt betrayed after you put so much effort and then to hear all this shit comments, even if you understand that that is usual internet behavior. Anyway, removing issue was a stupid idea. But I was pissed off with last two personal comments, especially while sitting and thinking how to solve the problem. I am sorry for doing that.

&gt; It’s been three years since I started actix project (time flies). I learnt a lot, i meet new people, I found language that I really like and want to use it fulltime, I found fun job. But damage to the project's reputation is done and I don’t think it is possible to recover. Actix always will be “shit full of UB” and “benchmark cheater”. (Btw, with tfb benchmark I just wanted to push rust to the limits, I wanted it to be on the top, I didn’t want to push other rust frameworks down.) Everything started with actix, then actix-web and then actix-net. It took a lot of time to design api and architecture. Each of this projects was rewritten from scratch at least 4-5 time. I hope I expanded some boundaries and found few new patterns, I hope other developers will check source code and find inspiration to move even further. Nowadays supporting actix project is not fun, and be part of rust community is not fun as well.

&gt; I am done with open source.

&gt; P.S. I moved actix-net and actix-web project to my personal github account. I will make decision during next couple days what to do. I don’t want to see the project becomes ghost of what it was. Maintainers must understand how everything work, but don’t anyone who does and those who could are busy with other projects. At the moment I am planing to make repos private and then delete them (will remove benchmarks as well), unless others suggest better ideas.

&gt; Everything has to come to the end. It was fun trip but now is time to move on. Life should be fun.
## [6][Gauging interest in an actix-web (and siblings) fork.](https://www.reddit.com/r/rust/comments/eq4xsu/gauging_interest_in_an_actixweb_and_siblings_fork/)
- url: https://www.reddit.com/r/rust/comments/eq4xsu/gauging_interest_in_an_actixweb_and_siblings_fork/
---
You’ve heard the news. You might even depend on the framework for a personal project. Maybe even a work project.

It seems in my communication with Nikolay on Gitter that we won’t be seeing a return to the Actix org on GitHub. Maybe his concerns are valid. And maybe his reasons for exiling the code from @actix we’re warranted. It’s unfortunate, but we have an opportunity to take this amazing project onwards.

A fork is not the most ideal approach in my opinion but I believe that this project doesn’t deserve to be abandoned and that some of us can do that.

A few months ago myself and probably about 30 others joined the @actix/contributors team on GitHub. There are people willing to help. How about it?
## [7][hyper on async-std proof-of-concept](https://www.reddit.com/r/rust/comments/eq8x17/hyper_on_asyncstd_proofofconcept/)
- url: https://github.com/leo-lb/hyper-async-std
---

## [8][Microsoft's Rust inspired research language has been released](https://www.reddit.com/r/rust/comments/eq089q/microsofts_rust_inspired_research_language_has/)
- url: https://github.com/microsoft/verona
---

## [9][actix-support/letter](https://www.reddit.com/r/rust/comments/eq5l6s/actixsupportletter/)
- url: https://github.com/actix-support/letter
---

## [10][Introducing sksg - The static site generation pipeline](https://www.reddit.com/r/rust/comments/eqcnp3/introducing_sksg_the_static_site_generation/)
- url: https://www.reddit.com/r/rust/comments/eqcnp3/introducing_sksg_the_static_site_generation/
---
(So, I've been working on this since December, and for the tiny scope, it certainly took a long time to get it out. Even now it's still rough around the edges, so feedback is very much appreciated, especially regarding docs.)

sksg is a static site generator that's inspired by the Unix philosophy of gluing smaller tools to do one large thing.

The core of sksg will just locate your pages and create the appropriate "pipelines", while the actual page generation logic happens on the other tools, (which don't need to be designed for sksg specifically)

One advantage of this is that you aren't locked into a specific language's ecosystem, as everything that supports stdin/stdout can be used as a pipeline step (including shell scripts)

There are also some simple tools distributed alongside sksg to get people started. However, they might not be as polished or as "batteries included" as one might expect.

If this sounds interesting to you, feel free to check it out [here](https://issizler.club/sksg/)
## [11][How To Write Fast Rust Code](https://www.reddit.com/r/rust/comments/eq168d/how_to_write_fast_rust_code/)
- url: http://likebike.com/posts/How_To_Write_Fast_Rust_Code.html
---

## [12][false ban](https://www.reddit.com/r/rust/comments/eqgd2d/false_ban/)
- url: https://www.reddit.com/r/rust/comments/eqgd2d/false_ban/
---
i was banned from rustralasia.net au long my most successful server ever i had so much stuff and was having the most fun ive ever had and logged on im almost out of data so i was lagging really bad and getting teleported back i jumped off a roof and got banned for fly hacking

my steam name is chocolatemilk if anybody can help
