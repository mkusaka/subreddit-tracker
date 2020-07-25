# ruby
## [1][Make API payloads generic - API only ruby on rails course (chapter 6)](https://www.reddit.com/r/ruby/comments/hxm279/make_api_payloads_generic_api_only_ruby_on_rails/)
- url: https://duetcode.io/rails-api-only-course/make-api-payloads-generic
---

## [2][Installing Ruby on VSCode Please Help!](https://www.reddit.com/r/ruby/comments/hxazpb/installing_ruby_on_vscode_please_help/)
- url: https://www.reddit.com/r/ruby/comments/hxazpb/installing_ruby_on_vscode_please_help/
---
Hey guys I've been trying to figure out how to use ruby on VSCode for 2 hours on windows and I have no idea. I have Ruby, ruby extension pack, and vscode ruby extensions installed. I made a test.rb file and wrote some code but theres no run option when i right click. If I click run - start debugging on the top all I dont see any code executed.  It says Node.js (preview) in the top right. I dont know what that means?

Please help! What do I do?

&amp;#x200B;

https://preview.redd.it/mr6nxvzvnvc51.png?width=1920&amp;format=png&amp;auto=webp&amp;s=f349e3f3bb85d4f3201e4fb3354a25425071eed6
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
## [4][How can I tell HOW I installed ruby on a Mac?](https://www.reddit.com/r/ruby/comments/hx77se/how_can_i_tell_how_i_installed_ruby_on_a_mac/)
- url: https://www.reddit.com/r/ruby/comments/hx77se/how_can_i_tell_how_i_installed_ruby_on_a_mac/
---
So, I've installed Ruby and rails and gems and all that awhile back. I had a directory of Jekyll themes I had been trying, and they worked. Now, when I try to run them I keep getting errors about not having a .bundle directory and other errors. I am trying to figure out what to do, and right now I am trying to figure HOW I installed it all in the first place. Can anyone tell me how I can figure out HOW I installed ruby and rails and whatnot please? For example, how can I tell if I used rbenv, or rmv, or the System, or brew, etc. ? Then, I just want to nuke it all and start over (again, for the 1000th time). Thank you!
## [5][Fun Facts about Ruby #11: The different ways to generate symbol with special characters](https://www.reddit.com/r/ruby/comments/hwy9c3/fun_facts_about_ruby_11_the_different_ways_to/)
- url: https://i.redd.it/lnb7dvf6prc51.png
---

## [6][Ideas for Cli app](https://www.reddit.com/r/ruby/comments/hxcx0c/ideas_for_cli_app/)
- url: https://www.reddit.com/r/ruby/comments/hxcx0c/ideas_for_cli_app/
---
So like the title says, in school we need to create a cli app using active record and at least 3 classes and tables. Any ideas for such a simple thing? I’m not trying to re create the wheel but something that can easily be done in a week. I have some ideas already but I am interested what others have to say! 

Thanks!!
## [7][Introducing GoodJob 1.0, a new Postgres-based, multithreaded, ActiveJob backend for Ruby on Rails](https://www.reddit.com/r/ruby/comments/hwqmlf/introducing_goodjob_10_a_new_postgresbased/)
- url: https://island94.org/2020/07/introducing-goodjob-1-0
---

## [8][Is this possible with ruby https://github.com/uber/piranha?](https://www.reddit.com/r/ruby/comments/hwwtv3/is_this_possible_with_ruby/)
- url: https://www.reddit.com/r/ruby/comments/hwwtv3/is_this_possible_with_ruby/
---
I was looking at [https://github.com/uber/piranha](https://github.com/uber/piranha) project on Github. Basically it can remove stale feature flags that are no longer used. It currently supports only three languages Java, Objective C, and Swift. I'm curious if something like this can be used in a language like Ruby to remove such stale flags from my RoR project. 

My concerns are since Ruby is not a statically typed language and with meta-programming, we cannot deterministically find all like in the code where it is being used. Can someone please help me understand if it is possible to port it to  Ruby?
## [9][The best online course for beginners](https://www.reddit.com/r/ruby/comments/hwm9ks/the_best_online_course_for_beginners/)
- url: https://www.reddit.com/r/ruby/comments/hwm9ks/the_best_online_course_for_beginners/
---
I'm trying to learn Ruby (or RoR) from scratch, but I can't seem to find a tutorial that gets me started. I need one that starts from scratch with the setup for MAC, and everything. Could you please suggest something?
## [10][Am I implementing Dragonfly Gem correctly to my Ruby on Rails app ? [joke]](https://www.reddit.com/r/ruby/comments/hwf5t1/am_i_implementing_dragonfly_gem_correctly_to_my/)
- url: https://github.com/markevans/dragonfly/issues/509
---

