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
## [2][What's everyone working on this week (41/2020)?](https://www.reddit.com/r/rust/comments/j57zzs/whats_everyone_working_on_this_week_412020/)
- url: https://www.reddit.com/r/rust/comments/j57zzs/whats_everyone_working_on_this_week_412020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-41-2020/49684?u=llogiq)!
## [3][Make a Language in Rust, Part Seven: A REPL](https://www.reddit.com/r/rust/comments/j6p8wt/make_a_language_in_rust_part_seven_a_repl/)
- url: https://arzg.github.io/lang/7/
---

## [4][Is it worth learning Rust as a junior dev?](https://www.reddit.com/r/rust/comments/j6m5j6/is_it_worth_learning_rust_as_a_junior_dev/)
- url: https://www.reddit.com/r/rust/comments/j6m5j6/is_it_worth_learning_rust_as_a_junior_dev/
---
Hi guys,

As a junior software developer (5 years of coding experience ), I am familiar with JS, C# and Python. I want to learn a low level programming language as I can see myself working in the IoT industry in a few years.

I am not sure if I should spend my time on C, C++, Go or Rust.

Thanks.
## [5][/r/rust is partnering with the Rust communities on Discord, Mozilla's Matrix, and Stack Overflow as venues for "IRC-like" real-time chat. Please read!](https://www.reddit.com/r/rust/comments/j6b03y/rrust_is_partnering_with_the_rust_communities_on/)
- url: https://www.reddit.com/r/rust/comments/j6b03y/rrust_is_partnering_with_the_rust_communities_on/
---
**TL;DR**: [The Rust Community Discord](https://discord.com/invite/aVESxV8), [The Official Rust Discord](https://discord.gg/rust-lang), [#rust on chat.mozilla.org](https://chat.mozilla.org/#/room/#rust:mozilla.org), [The Stack Overflow Rust Chat](https://chat.stackoverflow.com/rooms/62927/rust).

Gather 'round, Rustaceans, for a brief lesson in history. In the beginning, the Rust community lived in two places: Mozilla's IRC server, and the rust-dev mailing list.

Respectively, the two venues represented two different modes of communication: *synchronous* and *asynchronous* (sound familiar?). Broadly speaking, synchronous venues excel at high-volume, low-importance topics (e.g. casual conversation, simple questions), whereas asynchronous venues excel at high-importance, low-volume topics (e.g. news, announcements).

Reddit is an asynchronous venue. When I began popularizing /r/rust in 2012, I saw it as a user-friendly alternative to the mailing list and a complementary companion to IRC. Sadly, Mozilla's IRC server has since been decommissioned; even before then, IRC's famously ~~hostile~~ vintage UX had splintered Rust users onto many other synchronous chat platforms, which did not tend to share IRC's admirable qualities of openness and federation.

Meanwhile the decline of IRC has had consequences for the subreddit, most notably in the profusion of "how do I do X?" threads. We implemented a stopgap in the form of [the stickied weekly questions thread](https://www.reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/), but the subreddit's rapid growth continues to exacerbate the problem. In the future we may need to adopt a policy of disallowing simple "how do I do X?" questions on the front page entirely, but before going that far we would like to take a preventative step: encouraging new venues for synchronous communication, as a way to relieve pressure on our front page.

Thus, we are happy to endorse the Rust communities at the following venues (all of which have web-based interfaces; no app necessary) and encourage our readers to use them as a first resort for help and Q&amp;A:

* [The Rust Community Discord](https://discord.com/invite/aVESxV8): not to be confused with The Official Rust Discord, the community Discord server has a wide variety of channels for specific topics of interest (embedded dev, game dev, etc.) in addition to the usual channels for general discussion, offtopic chat, and beginner help.

* [The Official Rust Discord](https://discord.gg/rust-lang): although the official Discord server (similarly [the official Zulip](https://rust-lang.zulipchat.com/)) is predominantly used for working on Rust itself, its `#rust-usage` and `#beginners` channels are direct successors to the old `#rust` and `#rust-beginners` channels from Mozilla IRC and are open and welcome to all.

* [#rust on chat.mozilla.org](https://chat.mozilla.org/#/room/#rust:mozilla.org): users who have a distaste for the closed and proprietary nature of Discord will be happy to learn that Mozilla has chosen [Matrix](https://matrix.org/) as their IRC replacement and hosts a Rust channel there as well; the link here will have you using the default Element web client, but you can use any Matrix-compatible client to connect.

* [The Stack Overflow Rust Chat](https://chat.stackoverflow.com/rooms/62927/rust): a little-known feature of Stack Overflow is its real-time chat, which is host to its own active Rust community. Note that Stack Overflow requires a user to have 20 points of reputation (equivalent to two upvotes) in order to participate in chats.

We are "partnering" with these venues in the sense that we have observed that they are all helpful and actively moderated, and we have reached out to mods of each venue to ensure that they are okay with us directing the Reddit firehose at them. You will find permanent links to each venue in the header menus (newreddit) or sidebar (oldreddit). And don't worry, we're not mandating that you use any of these venues: the weekly questions thread is here to stay.

If you have any questions, feel free to ask them below.
## [6][Anyone requiring a collaborator](https://www.reddit.com/r/rust/comments/j6nsq3/anyone_requiring_a_collaborator/)
- url: https://www.reddit.com/r/rust/comments/j6nsq3/anyone_requiring_a_collaborator/
---
So I maybe a little odd here, if this is not a place for posting it but anyway I wanted to ask if anyone working on a open-source Rust project needs a collaborator. I am always ready to help open source software but apart from filing PRs and issues, I haven't been really close to a single project. I am generally very active on GitHub as well as Gitlab.

So at very end I just wanted to ask if anyone needs a collaborator, just tell me in comments. I will help him
## [7][in: a trivial utility for running a command in a directory](https://www.reddit.com/r/rust/comments/j6ljjq/in_a_trivial_utility_for_running_a_command_in_a/)
- url: https://www.reddit.com/r/rust/comments/j6ljjq/in_a_trivial_utility_for_running_a_command_in_a/
---
I've been annoyed for a long time that there's no non-shell way to run a command in a specific working directory. You can use `env` to run a command in a modified environment, `sudo` to run a command as a different user, `nice` to change a command's priority, but for some reason, there's program that runs a command in a directory. So I made `in`, a trivial command that runs a subcommand in a directory https://crates.io/crates/in-directory.
## [8][Rendered a signed distance field to the terminal](https://www.reddit.com/r/rust/comments/j6j3p8/rendered_a_signed_distance_field_to_the_terminal/)
- url: https://gfycat.com/fearlesslittledachshund
---

## [9][What languages do you use with Rust? What languages do you think compliment Rust?](https://www.reddit.com/r/rust/comments/j6jgbo/what_languages_do_you_use_with_rust_what/)
- url: https://www.reddit.com/r/rust/comments/j6jgbo/what_languages_do_you_use_with_rust_what/
---
As per the title.

What languages do you use along side Rust? Bonus points if you get paid to do so.

What languages do you think compliment Rust? Or vis-versa which languages do you think Rust compliments?
## [10][[Git] Repository summary on your terminal](https://www.reddit.com/r/rust/comments/j6c5nb/git_repository_summary_on_your_terminal/)
- url: https://i.redd.it/a7ryqatz3jr51.png
---

## [11][A new rust community](https://www.reddit.com/r/rust/comments/j6hp6f/a_new_rust_community/)
- url: https://www.reddit.com/r/rust/comments/j6hp6f/a_new_rust_community/
---
Hey /r/rust.

I've been working on this for a bit and have shared it on HN and other places, so some of you may have seen this already. 

Gurlic is a platform to build communities. 

Right now, most of the code is in Go, and the shortlink server is written in Rust. I'm working on replacing bits of Go code with Rust as I become more proficient with the language. It honestly took me several tries to 'grok' the language, but having something real-world to build helps I guess.

Anyway, here's the main page, https://gurlic.com, and an example community, https://gurlic.com/rust
  
And here are some random links to give you an idea of the site, without having to type a wall of text:  
 
 
A poll:  
https://gurlic.com/programming/post/322832131900309509
  
A post with images:  
https://gurlic.com/nomanssky/screenshots/post/322921012087750661
  
A custom-domain gallery:  
https://gurlic.art  
  
A custom-domain publication:  
https://classics.wtf  
(same as https://gurlic.com/classics or https://classics.gurlic.com)  
  
A post from it:  
https://classics.wtf/the-odyssey-book-iii
    
Another example community:  
https://gurlic.space  
  
I'm also working on having better/deeper matrix integration, but it's a work in progress. Right now new users get a @username:gurlic.com matrix account. There's a temporarily self-hosted instance of element at https://gurlic.com/chat, until I can integrate hydrogen-web into gurlic.

Anyway, I'd be happy to hear feedback and advice. Thanks.
## [12][The Rust 2021 Experience - Year of the Macro](https://www.reddit.com/r/rust/comments/j6e9m4/the_rust_2021_experience_year_of_the_macro/)
- url: https://casualhacks.net/blog/2020-10-05/the-rust-2021-experience-macros/
---

