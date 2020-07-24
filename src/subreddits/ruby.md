# ruby
## [1][Introducing GoodJob 1.0, a new Postgres-based, multithreaded, ActiveJob backend for Ruby on Rails](https://www.reddit.com/r/ruby/comments/hwqmlf/introducing_goodjob_10_a_new_postgresbased/)
- url: https://island94.org/2020/07/introducing-goodjob-1-0
---

## [2][Fun Facts about Ruby #11: The different ways to generate symbol with special characters](https://www.reddit.com/r/ruby/comments/hwy9c3/fun_facts_about_ruby_11_the_different_ways_to/)
- url: https://i.redd.it/lnb7dvf6prc51.png
---

## [3][sequel-activerecord_connection now fully supports Sequel's transaction API](https://www.reddit.com/r/ruby/comments/hx0pzd/sequelactiverecord_connection_now_fully_supports/)
- url: https://www.reddit.com/r/ruby/comments/hx0pzd/sequelactiverecord_connection_now_fully_supports/
---
A while ago I created the [sequel-activerecord_connection](https://github.com/janko/sequel-activerecord_connection) gem, which allows you to use [Sequel](https://github.com/jeremyevans/sequel) alongside ActiveRecord and reuse the same database connection. This makes it easier to try Sequel out or to use libraries that depend on Sequel (e.g. [Rodauth](https://github.com/jeremyevans/rodauth/)) in ActiveRecord-based projects.

Originally the transaction support was partially implemented, where the most common Sequel transaction options were emulated with the ActiveRecord API. In the latest version I've [rewritten the transaction code](https://github.com/janko/sequel-activerecord_connection/pull/8) to fully support Sequel's transaction API.

    require "sequel"

    DB = Sequel.postgres(test: false) # avoid creating a database connection
    DB.extension :activerecord_connection # use ActiveRecord's database connection

    DB.transaction(isolation: :serializable) do # handles isolation levels
      DB.after_commit { ... } # runs after transaction commits
      DB.transaction(savepoint: true) do # creates a savepoint
        DB.after_rollback { ... } # runs after savepoint rolls back
      end
      DB.in_transaction? #=&gt; true
    end

Supporting Sequel's transaction/savepoint hooks was especially important for me, because it's a very useful feature when implementing business logic, and it's something ActiveRecord doesn't really support (AR supports transactional _model callbacks_, which are tied to model's lifecycle and aren't usable when I want to keep my business logic away from AR callbacks).
## [4][Is this possible with ruby https://github.com/uber/piranha?](https://www.reddit.com/r/ruby/comments/hwwtv3/is_this_possible_with_ruby/)
- url: https://www.reddit.com/r/ruby/comments/hwwtv3/is_this_possible_with_ruby/
---
I was looking at [https://github.com/uber/piranha](https://github.com/uber/piranha) project on Github. Basically it can remove stale feature flags that are no longer used. It currently supports only three languages Java, Objective C, and Swift. I'm curious if something like this can be used in a language like Ruby to remove such stale flags from my RoR project. 

My concerns are since Ruby is not a statically typed language and with meta-programming, we cannot deterministically find all like in the code where it is being used. Can someone please help me understand if it is possible to port it to  Ruby?
## [5][Am I implementing Dragonfly Gem correctly to my Ruby on Rails app ? [joke]](https://www.reddit.com/r/ruby/comments/hwf5t1/am_i_implementing_dragonfly_gem_correctly_to_my/)
- url: https://github.com/markevans/dragonfly/issues/509
---

## [6][The best online course for beginners](https://www.reddit.com/r/ruby/comments/hwm9ks/the_best_online_course_for_beginners/)
- url: https://www.reddit.com/r/ruby/comments/hwm9ks/the_best_online_course_for_beginners/
---
I'm trying to learn Ruby (or RoR) from scratch, but I can't seem to find a tutorial that gets me started. I need one that starts from scratch with the setup for MAC, and everything. Could you please suggest something?
## [7][Ruby for Good (Online)](https://www.reddit.com/r/ruby/comments/hw718c/ruby_for_good_online/)
- url: https://www.reddit.com/r/ruby/comments/hw718c/ruby_for_good_online/
---
Hey all,

I'm delighted to share that [registration ](https://ti.to/codeforgood/rubyforgood)for our second virtual Ruby for Good (RfG) event is now open! This is an annual event held every year since 2014, based out of the DC-metro area, where Ruby programmers (of all skill levels), from all over the globe, get together to build technology solutions in support of nonprofits with meaningful missions that benefit our communities. 

Last year we built applications for diaper banks, refuge restrooms, the amazon conservation team, and many other nonprofits serving critical missions, that really need support from people with diverse technology skill-sets, that would not otherwise be able to afford this support. This year, we will be focusing on projects that affect our local communities, and given travel and gathering restrictions, the 2020 event will be virtual.

You can find more information about the event [here](https://rubyforgood.org/attend), a list of the most frequently asked questions [here](https://rubyforgood.org/faq), and the event schedule [here](https://rubyforgood.org/2020).I'd be thrilled to answer any questions!

Happiness,

Sean and the RfG team
## [8][How to speed up assets precompile for Ruby on Rails apps](https://www.reddit.com/r/ruby/comments/hwcu7e/how_to_speed_up_assets_precompile_for_ruby_on/)
- url: https://www.reddit.com/r/ruby/comments/hwcu7e/how_to_speed_up_assets_precompile_for_ruby_on/
---
You spend too much time to deploy your rails project especially on the assets:precompile step, or maybe sometimes you see the following error during the assets precompilation:

*FATAL ERROR: Ineffective mark-compacts near heap limit Allocation failed - JavaScript heap out of memory*

So for sure, this short article will help you to have 50% improvement only with some minor changes:

[https://medium.com/@ali.sepehri.kh/how-to-speed-up-assets-precompile-for-ruby-on-rails-apps-e0338d8d7301?source=friends\_link&amp;sk=995eaef8fdbc73e533c6384f6d47d37e](https://medium.com/@ali.sepehri.kh/how-to-speed-up-assets-precompile-for-ruby-on-rails-apps-e0338d8d7301?source=friends_link&amp;sk=995eaef8fdbc73e533c6384f6d47d37e)
## [9][How to handle STI with Fast JSON API serializer for Ruby objects](https://www.reddit.com/r/ruby/comments/hw1u9u/how_to_handle_sti_with_fast_json_api_serializer/)
- url: https://www.reddit.com/r/ruby/comments/hw1u9u/how_to_handle_sti_with_fast_json_api_serializer/
---
 A most popular and faster gem for serializing data in Ruby on Rails: [Fast JSON API](https://github.com/Netflix/fast_jsonapi)

Apparently, this gem is not supporting STI. Check issue: [Lookup for STI models](https://github.com/Netflix/fast_jsonapi/issues/225)

My only way around this is by [Conditional Attributes](https://github.com/Netflix/fast_jsonapi#conditional-attributes) and check if an object is some particular inherited class to parse additional attributes and relations.

How do you handle this?
## [10][Please help](https://www.reddit.com/r/ruby/comments/hwd5wz/please_help/)
- url: https://www.reddit.com/r/ruby/comments/hwd5wz/please_help/
---
I am fairly new to ruby and am trying to create a calculator.However the program enters into a loop and does not return to terminal or exits.Please help me.

code in imgur

[https://imgur.com/gallery/aCCxPtd](https://imgur.com/gallery/aCCxPtd)
