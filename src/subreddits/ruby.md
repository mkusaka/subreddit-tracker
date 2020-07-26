# ruby
## [1][Optimally distribute and execute RSpec suites among parallel workers; for faster CI builds](https://www.reddit.com/r/ruby/comments/hy4dvk/optimally_distribute_and_execute_rspec_suites/)
- url: https://github.com/skroutz/rspecq
---

## [2][Make API payloads generic - API only ruby on rails course (chapter 6)](https://www.reddit.com/r/ruby/comments/hxm279/make_api_payloads_generic_api_only_ruby_on_rails/)
- url: https://duetcode.io/rails-api-only-course/make-api-payloads-generic
---

## [3][Starting a monthly newsletter for cool stuff I encounter on Rails, React and Graphql.](https://www.reddit.com/r/ruby/comments/hxpads/starting_a_monthly_newsletter_for_cool_stuff_i/)
- url: https://www.reddit.com/r/ruby/comments/hxpads/starting_a_monthly_newsletter_for_cool_stuff_i/
---
Starting a monthly newsletter on Ruby on Rails, React and GraphQL developers to read about some cool stuff happening in the community.

In a month we come across a lot cool stuff happening in the community, learn new things and contribute to the open-source world and newsletters is the best way to share some of those finding/blogs/videos.

This newsletter would consist of the following things:

* Dev Joke
* My Blogs.
* Reading/watching list.
* Open-source libraries/contributions.
* Rails changelog.
* Watching/Reading/Reviews.

If you'd like to read my monthly newsletter. Please do subscribe here [https://buttondown.email/abhaynikam](https://buttondown.email/abhaynikam)
## [4][Failing to Deploy RoR at DO Droplet](https://www.reddit.com/r/ruby/comments/hxvixo/failing_to_deploy_ror_at_do_droplet/)
- url: /r/rubyonrails/comments/hxtwhv/failing_from_reach_index_page_on_ror_deployed_at/
---

## [5][Installing Ruby on VSCode Please Help!](https://www.reddit.com/r/ruby/comments/hxazpb/installing_ruby_on_vscode_please_help/)
- url: https://www.reddit.com/r/ruby/comments/hxazpb/installing_ruby_on_vscode_please_help/
---
Hey guys I've been trying to figure out how to use ruby on VSCode for 2 hours on windows and I have no idea. I have Ruby, ruby extension pack, and vscode ruby extensions installed. I made a test.rb file and wrote some code but theres no run option when i right click. If I click run - start debugging on the top all I dont see any code executed.  It says Node.js (preview) in the top right. I dont know what that means?

Please help! What do I do?

&amp;#x200B;

https://preview.redd.it/mr6nxvzvnvc51.png?width=1920&amp;format=png&amp;auto=webp&amp;s=f349e3f3bb85d4f3201e4fb3354a25425071eed6
## [6][sequel-activerecord_connection now fully supports Sequel's transaction API](https://www.reddit.com/r/ruby/comments/hx0pzd/sequelactiverecord_connection_now_fully_supports/)
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
## [7][How can I tell HOW I installed ruby on a Mac?](https://www.reddit.com/r/ruby/comments/hx77se/how_can_i_tell_how_i_installed_ruby_on_a_mac/)
- url: https://www.reddit.com/r/ruby/comments/hx77se/how_can_i_tell_how_i_installed_ruby_on_a_mac/
---
So, I've installed Ruby and rails and gems and all that awhile back. I had a directory of Jekyll themes I had been trying, and they worked. Now, when I try to run them I keep getting errors about not having a .bundle directory and other errors. I am trying to figure out what to do, and right now I am trying to figure HOW I installed it all in the first place. Can anyone tell me how I can figure out HOW I installed ruby and rails and whatnot please? For example, how can I tell if I used rbenv, or rmv, or the System, or brew, etc. ? Then, I just want to nuke it all and start over (again, for the 1000th time). Thank you!
## [8][Fun Facts about Ruby #11: The different ways to generate symbol with special characters](https://www.reddit.com/r/ruby/comments/hwy9c3/fun_facts_about_ruby_11_the_different_ways_to/)
- url: https://i.redd.it/lnb7dvf6prc51.png
---

## [9][Ideas for Cli app](https://www.reddit.com/r/ruby/comments/hxcx0c/ideas_for_cli_app/)
- url: https://www.reddit.com/r/ruby/comments/hxcx0c/ideas_for_cli_app/
---
So like the title says, in school we need to create a cli app using active record and at least 3 classes and tables. Any ideas for such a simple thing? I’m not trying to re create the wheel but something that can easily be done in a week. I have some ideas already but I am interested what others have to say! 

Thanks!!
## [10][Introducing GoodJob 1.0, a new Postgres-based, multithreaded, ActiveJob backend for Ruby on Rails](https://www.reddit.com/r/ruby/comments/hwqmlf/introducing_goodjob_10_a_new_postgresbased/)
- url: https://island94.org/2020/07/introducing-goodjob-1-0
---

