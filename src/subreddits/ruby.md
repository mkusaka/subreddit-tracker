# ruby
## [1][After 10 years of Ruby I just realized you can step directly into a method with pry-byebug from IRB...](https://www.reddit.com/r/ruby/comments/h9o0e9/after_10_years_of_ruby_i_just_realized_you_can/)
- url: https://www.reddit.com/r/ruby/comments/h9o0e9/after_10_years_of_ruby_i_just_realized_you_can/
---
Adding a `binding.pry` to your source and running your source is one thing I've always known you can do, but that has always been a chore for stepping into methods starting from IRB. Normally my hack around this was to redefine `my_method` with a `binding.pry` on the first line of that method and load that back into my repl

Turns out you can just do `binding.pry; my_method` and Pry will step right in; no editing or redefinition required

Counting the hours wasted...

EDIT: source where I stumbled across this https://stackoverflow.com/a/36762059
## [2][football.csv - sportdb-importers Gem Update - Read League &amp; Match Datasets in Comma-Separated Values (CSV) Format into Any SQL Database](https://www.reddit.com/r/ruby/comments/ha022y/footballcsv_sportdbimporters_gem_update_read/)
- url: https://github.com/sportdb/sport.db/tree/master/sportdb-importers
---

## [3][Upgrade Path for Rails Apps](https://www.reddit.com/r/ruby/comments/ha3mgs/upgrade_path_for_rails_apps/)
- url: https://www.reddit.com/r/ruby/comments/ha3mgs/upgrade_path_for_rails_apps/
---
I need an upgrade path for Rails apps which need to scale and I'm thinking:

1. Rails/MRI
2. Rails/TruffleRuby
3. Rails/TruffleRuby with Kotlin components
4. Rewrite in Kotlin

I also thought of Clojure but it doesn't really fit the Rails model. Any other ideas? I believe TruffleRuby is almost there with Rails but until then JRuby can be substituted.
## [4][Pythonic failures - How to unlearn snaky behavior?](https://www.reddit.com/r/ruby/comments/h9y7yd/pythonic_failures_how_to_unlearn_snaky_behavior/)
- url: https://www.reddit.com/r/ruby/comments/h9y7yd/pythonic_failures_how_to_unlearn_snaky_behavior/
---
I know, it will go away at some point :D

But coming from Python it's just too tempting to assume this:

    str = "1324"
    str[0] == 1 # must be True 💆‍♂️

Do you also remember snaky behavior or any other "bad habits" you have taken over from other languages?
## [5][Using Redis Lua scripts in gems](https://www.reddit.com/r/ruby/comments/h9ooo0/using_redis_lua_scripts_in_gems/)
- url: http://www.wjwh.eu/posts/2020-06-15-redis-eval-trick.html
---

## [6][How to run slow RSpec files on Github Actions with parallel jobs by doing an auto split of the spec file by test examples](https://www.reddit.com/r/ruby/comments/h9beb4/how_to_run_slow_rspec_files_on_github_actions/)
- url: https://docs.knapsackpro.com/2020/how-to-run-slow-rspec-files-on-github-actions-with-parallel-jobs-by-doing-an-auto-split-of-the-spec-file-by-test-examples
---

## [7][datoji - Remote JSON storage server. Create, Read, Update, Delete and Search JSON data.](https://www.reddit.com/r/ruby/comments/h9fvqs/datoji_remote_json_storage_server_create_read/)
- url: https://datoji.dev
---

## [8][Allowing dots in Rails routes](https://www.reddit.com/r/ruby/comments/h9l7yf/allowing_dots_in_rails_routes/)
- url: https://prathamesh.tech/2020/06/15/allowing-dots-in-rails-routes/
---

## [9][My wishlist for improvements in the Ruby ecosystem](https://www.reddit.com/r/ruby/comments/h9h787/my_wishlist_for_improvements_in_the_ruby_ecosystem/)
- url: https://www.reddit.com/r/ruby/comments/h9h787/my_wishlist_for_improvements_in_the_ruby_ecosystem/
---
1) self-contained deployment
The best would be to compile to a self-contained binary but i think the closest you get to this with Ruby is the GraalVM/Substrate ecosystem.
What should be more achievable is something like a .NET self-contained archive which you can unpack on the system.

I'd like to build my Ruby and the c-extensions (maybe with musl) and just copy paste it to the system.


2) fast self-contained http server that we could use by default
I think having something that is production ready albeit not as configurable and robust as nginx by default would be really nice. Should also be a part of the self-contained deployment experience.

All of this could be multiplatform and rely on docker or a virtualization to build the necessary artifacts.

3) popular and robust solution to build mobile apps
RubyMotion is almost never a choice for anyone. It's mostly used because people dont want to write JavaScript (RN),native or Xamarin.
I want this to be the first option for a Ruby shop when they are considering a mobile app. Right now it is not an it's questionable to bet your business success on it.
It's not the cost it's the uncertainty and lack of popularity that concerns me
## [10][Noob ruby/Linux cron job question](https://www.reddit.com/r/ruby/comments/h93lk9/noob_rubylinux_cron_job_question/)
- url: https://www.reddit.com/r/ruby/comments/h93lk9/noob_rubylinux_cron_job_question/
---
Thanks for reading this. I have very little experience with Ruby.  I have a RedHat Linux server that has a homemade Ruby application running on it. The person who created it is long gone. It's started manually now, but I'd like it to start on reboot. I was going to add a cron job for the script, but it's started with the command "rvmsudo bundle exec puma". Is adding this to a "@reboot" cron job the best path here, and if so do I just put that path and command in a script?

&amp;#x200B;

Thanks again, any help is appreciated.
