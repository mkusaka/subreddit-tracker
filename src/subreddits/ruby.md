# ruby
## [1][Ruby Structure](https://www.reddit.com/r/ruby/comments/g4sieh/ruby_structure/)
- url: https://www.reddit.com/r/ruby/comments/g4sieh/ruby_structure/
---
Hi Guy and Girls,

I have a beginner question I find really hard to find the right answer for, how do you structure in Ruby.

I have a very bad habit of writing all my code in 1 file, after creating a gem with bundle. 

I will write almost all of my code in the main file in the lib folder, but how do you know what to split up in different files, is where some kind of rule you follow then you split the code up in different files. ?

And is it best practice to write all your projects in a gem structure format if its only for private use and not to ever be published or is where a better structure for not creating gems but just a project?
## [2][I find it surprising (in the pejorative sense) that Enumerator#next can't take an argument. Is there some good reason for that, or has it just kinda slipped through the cracks?](https://www.reddit.com/r/ruby/comments/g4ples/i_find_it_surprising_in_the_pejorative_sense_that/)
- url: https://www.reddit.com/r/ruby/comments/g4ples/i_find_it_surprising_in_the_pejorative_sense_that/
---
I have an infinite sequence (`Enumerator.new { |yielder| ... loop { yielder &lt;&lt; ... } }`) that I know will continue producing values, and I'd like to grab two at a time. `a, b = enum.next, enum.next` works, of course, but `enum.next(2)` would be even better.

I [dug around](https://github.com/ruby/ruby/blob/bfe706716f7692f0a672a15a261ffffa25526cb1/enumerator.c#L903-L908) and found that `Enumerator#next` specifically makes sure it's returning a single value (the "sv" in `ary2sv()`, presumably), though that value is permitted to be an `Array`. It wouldn't make sense to rewrite this particular enumeration to yield a two-element array, so I won't pursue that avenue.

There's [this feature request](https://bugs.ruby-lang.org/issues/9557) from many moons ago that wanted `#next`'s argument to skip forward in the enumeration, but nothing much came of it, so I wouldn't know where to begin tacking this on without breaking things. Would it even be a good idea, or is there something I'm missing?
## [3][Reactive, realtime web apps in Ruby on Rails without JS using Stimulus Reflex](https://www.reddit.com/r/ruby/comments/g48g4b/reactive_realtime_web_apps_in_ruby_on_rails/)
- url: https://www.reddit.com/r/ruby/comments/g48g4b/reactive_realtime_web_apps_in_ruby_on_rails/
---
Thought you guys would be interested in this.

Stimulus Reflex lets you build realtime, reactive apps in Rails similar to Phoenix LiveView. You don't have to write any Javascript and it will use Rails to render all the HTML updates server-side. It uses ActionCable to trigger updates and then tells Rails to render the current page again and send that back over the websocket. It'll then use DOM diffing to update the page automatically for you.

Definitely one of the coolest projects going on right now in the Rails world I'd say. It looks like we might see some similar things in Turbolinks 6 when that comes out too which is exciting.

[https://gorails.com/episodes/stimulus-reflex-basics?autoplay=1](https://gorails.com/episodes/stimulus-reflex-basics?autoplay=1)
## [4][Optimizing Ruby Lazy Initialization in TruffleRuby with Deoptimization](https://www.reddit.com/r/ruby/comments/g465y7/optimizing_ruby_lazy_initialization_in/)
- url: https://engineering.shopify.com/blogs/engineering/optimizing-ruby-lazy-initialization-in-truffleruby-with-deoptimization
---

## [5][Simple game engine with Ruby + Rust (help and feedback requested)](https://www.reddit.com/r/ruby/comments/g42b12/simple_game_engine_with_ruby_rust_help_and/)
- url: https://www.reddit.com/r/ruby/comments/g42b12/simple_game_engine_with_ruby_rust_help_and/
---
[Dathos Game Engine](https://github.com/BrianMWest/dathos-game-engine)

## Background

I've been working on a game over the past several months and eventually needed to incorporate a dynamic language to make it easier to iterate over features and mechanics. I chose Ruby because I've been infatuated with it for awhile but hadn't been able to work with it much. I feel that there are many idioms in Ruby that have inspired the syntax of Rust. After the COVID-19 spread and subsequent lockdown, I put together a small Slack-based game for a community orchestra I am a part of so that we could keep hanging out even though we couldn't meet. When I did that, I took the code I had been using and have been modularizing it into a game engine that could be used for other games. Ideally, I'd like to see this be something that others could use for game jams and small games.

## Engine

The game engine is a combination of Ruby and Rust. The intention is that the game objects would live in Ruby and the behavior logic would exist in Ruby classes. Rust can be used to extend the game with extra functionality by creating Ruby modules and hooking into lifecycle events for `init`, `pre_update`, `update`, and `post_update`. Examples would be a collision system, tilemap system, Slack integration module or a Twitch integration module.

## Help I need

I've created some GitHub issues of things I _know_ I need. Some of these are more Rust-centric. I know this community is a Ruby community, and a lot of the engine is written in Rust, but I could use help in a few different ways from the Ruby community.

1. idiomatic feedback. I've created a few GitHub issues around this, but I'd like to get feedback from Ruby users on keeping the usage of Ruby consistent with the language's best practices.
2. Ruby packaging. If there are any Ruby users who are exceptionally good at packaging Ruby applications, some feedback on how we might create a utility to package a game with the Ruby runtime and any gems would be greatly appreciated. This is especially a concern given that I'd like to be able to support as many target operating systems / architectures as possible.
3. Ruby modules. Not everything in the game engine needs to exist in Rust. If there are any additional features, utilities, modules, etc. that can be included that don't require Rust (file loading utilities, game math or physics utilities?), I'm perfectly happy to include plain Ruby in the engine (I did this with a `GameUtils` module, for example.
4. Users! I would love to just get people to try this out to find problems with it and see where things break down or users run into roadblocks.
5. Editor? Bit of a pipe dream, but if any enterprising individuals wanted to make some sort of interface to work with this to make it easier to set up a new project or add game objects, that'd be super cool.

Here is a link to the engine and the Slack-based game I made with it:

[Dathos Game Engine](https://github.com/BrianMWest/dathos-game-engine)

[Melody Madness / Chordal Catastrophe](https://github.com/BrianMWest/melody-madness)
## [6][RSS as RESTful API](https://www.reddit.com/r/ruby/comments/g4532n/rss_as_restful_api/)
- url: https://github.com/davidesantangelo/feedirss-api
---

## [7][Over 700 Malicious Typosquatted Libraries Found On RubyGems Repository](https://www.reddit.com/r/ruby/comments/g3qwx3/over_700_malicious_typosquatted_libraries_found/)
- url: https://thehackernews.com/2020/04/rubygem-typosquatting-malware.html
---

## [8][Having trouble with API pagination with HTTParty.](https://www.reddit.com/r/ruby/comments/g40jct/having_trouble_with_api_pagination_with_httparty/)
- url: https://www.reddit.com/r/ruby/comments/g40jct/having_trouble_with_api_pagination_with_httparty/
---
I am using an api and it says there are 120 pages, only 25 items displayed per page, but not sure how to pull all the pages. Ive been searching for a while now cant find anything useful. If I cant figure this out I will have to scrap this and start over but only have a few more days left to turn in a project.

Thanks for your help!
## [9][Ruby "main"](https://www.reddit.com/r/ruby/comments/g3px8i/ruby_main/)
- url: https://medium.com/@igor04/ruby-main-789ff58320f
---

## [10][Play around with Bezier curves in the browser using Ruby (source code link in the comments).](https://www.reddit.com/r/ruby/comments/g3wmy9/play_around_with_bezier_curves_in_the_browser/)
- url: https://v.redd.it/koobj5vpmnt41
---

