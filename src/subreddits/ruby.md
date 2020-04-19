# ruby
## [1][Simple game engine with Ruby + Rust (help and feedback requested)](https://www.reddit.com/r/ruby/comments/g42b12/simple_game_engine_with_ruby_rust_help_and/)
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
## [2][RSS as RESTful API](https://www.reddit.com/r/ruby/comments/g4532n/rss_as_restful_api/)
- url: https://github.com/davidesantangelo/feedirss-api
---

## [3][Optimizing Ruby Lazy Initialization in TruffleRuby with Deoptimization](https://www.reddit.com/r/ruby/comments/g465y7/optimizing_ruby_lazy_initialization_in/)
- url: https://engineering.shopify.com/blogs/engineering/optimizing-ruby-lazy-initialization-in-truffleruby-with-deoptimization
---

## [4][Over 700 Malicious Typosquatted Libraries Found On RubyGems Repository](https://www.reddit.com/r/ruby/comments/g3qwx3/over_700_malicious_typosquatted_libraries_found/)
- url: https://thehackernews.com/2020/04/rubygem-typosquatting-malware.html
---

## [5][Having trouble with API pagination with HTTParty.](https://www.reddit.com/r/ruby/comments/g40jct/having_trouble_with_api_pagination_with_httparty/)
- url: https://www.reddit.com/r/ruby/comments/g40jct/having_trouble_with_api_pagination_with_httparty/
---
I am using an api and it says there are 120 pages, only 25 items displayed per page, but not sure how to pull all the pages. Ive been searching for a while now cant find anything useful. If I cant figure this out I will have to scrap this and start over but only have a few more days left to turn in a project.

Thanks for your help!
## [6][Ruby "main"](https://www.reddit.com/r/ruby/comments/g3px8i/ruby_main/)
- url: https://medium.com/@igor04/ruby-main-789ff58320f
---

## [7][Play around with Bezier curves in the browser using Ruby (source code link in the comments).](https://www.reddit.com/r/ruby/comments/g3wmy9/play_around_with_bezier_curves_in_the_browser/)
- url: https://v.redd.it/koobj5vpmnt41
---

## [8][Ruby Bitwise Operators](https://www.reddit.com/r/ruby/comments/g3oecm/ruby_bitwise_operators/)
- url: https://medium.com/rubycademy/ruby-bitwise-operators-da57763fa368
---

## [9][Ruby Hash Default Values](https://www.reddit.com/r/ruby/comments/g33oyz/ruby_hash_default_values/)
- url: https://www.reddit.com/r/ruby/comments/g33oyz/ruby_hash_default_values/
---
I've written a new doc section about Hash default values.  It has been merged into Ruby's [hash.c](https://github.com/ruby/ruby/commit/39c965f2306d29524feae04ff6e710b32702851a#diff-eff9999082c8ce7d8ba1fc1d79f439cf), but not yet released.

I wrote it over at my doc project [AboutRuby](https://github.com/BurdetteLamar/AboutRuby/blob/master/core/Hash/api/markdown.md#hash-api), then PRed it over to Ruby. (Thanks, @nobu, for the review and merge.)
## [10][Processing high volume of unique messages exactly once while preserving order in a queue using SQS and Ruby-on-Rails](https://www.reddit.com/r/ruby/comments/g37ygs/processing_high_volume_of_unique_messages_exactly/)
- url: https://medium.com/@rcdexta/processing-high-volume-of-unique-messages-exactly-once-while-preserving-order-in-a-queue-d8d6184ded01
---

