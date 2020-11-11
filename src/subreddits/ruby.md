# ruby
## [1][Rebuilding Redis in Ruby - New chapters available](https://www.reddit.com/r/ruby/comments/jrt2up/rebuilding_redis_in_ruby_new_chapters_available/)
- url: https://www.reddit.com/r/ruby/comments/jrt2up/rebuilding_redis_in_ruby_new_chapters_available/
---
About four months ago I [posted about a project](https://www.reddit.com/r/ruby/comments/hmxcgv/redis_in_ruby_a_work_in_progress_online_book/) I've been working on, [Rebuilding Redis in Ruby](https://redis.pjam.me). There are six more chapters available now, so I figured I would post an update. Especially since the first three chapters were pretty short, and the later ones contain way more information:

1. [Chapter 4](https://redis.pjam.me/post/chapter-4-adding-missing-options-to-set/): Completing the SET command, that includes handling expiration and all the good stuff
2. [Chapter 5](https://redis.pjam.me/post/chapter-5-redis-protocol-compatibility/): Making the server RESP (Redis Protocol) compatible
3. [Chapter 6](https://redis.pjam.me/post/chapter-6-building-a-hash-table/): Building a Hash replacement, from scratch (including replicating the hashing algorithm used in Redis, SipHash)
4. [Chapter 7](https://redis.pjam.me/post/chapter-7-adding-list-commands/): Adding all List commands
5. [Chapter 8](https://redis.pjam.me/post/chapter-8-adding-hash-commands/): Adding all hash commands\* (\*except HSCAN)
6. [Chapter 9](https://redis.pjam.me/post/chapter-9-adding-set-commands/): Adding all set commands\* (\*except SSCAN)

If you have any feedback, please feel free to share it here or anywhere else (twitter, mail, etc ...), and if you're interested you can check what the future chapters will be [on the site](https://redis.pjam.me/chapters/):

1. Chapter 10: Sorted Sets
2. Chapter 11: Bitmaps:
3. Chapter 12: HyperLogLogs
4. Chapter 13: Geo Commands
5. Chapter 14: Pub/Sub
6. Chapter 15: Redis as an LRU cache
7. Chapter 16: Transactions

Hope you'll enjoy it!

PS: I definitely don't want to spam this sub, I'm only really on posting one more time about this, when all sixteen chapters will be completed.
## [2][Copy config.yml (bundled in .gem) to home directory after installation](https://www.reddit.com/r/ruby/comments/js6obr/copy_configyml_bundled_in_gem_to_home_directory/)
- url: https://www.reddit.com/r/ruby/comments/js6obr/copy_configyml_bundled_in_gem_to_home_directory/
---
Hey there, ruby community!       

I'd like to include a config.yml file with my gem, which should get copied to the users home directory, either before running the gem or on  the first gem run.

I have thought of the following options:

* Create a default config hash in ruby and write that to a file on the first run of the program. Drawbacks: No comments in yaml
* Fetch the config from a public repo. Drawbacks: Online only

Is there a way to bundle the gem with the config file included, and reference that file from within the code after the gem has been installed with gem install?

Thanks! :)
## [3][The art of errors in Ruby](https://www.reddit.com/r/ruby/comments/jrnk8v/the_art_of_errors_in_ruby/)
- url: https://longliveruby.com/articles/art-of-errors-in-ruby
---

## [4][What other things equals nil](https://www.reddit.com/r/ruby/comments/js2nel/what_other_things_equals_nil/)
- url: https://www.reddit.com/r/ruby/comments/js2nel/what_other_things_equals_nil/
---
I'm checking for equality between two variables, first and second. The first variable can be nil.   
I wanted to be sure other than second being nil also is there any cases where first == second would be true when first is nil?

i.e. first, second = nil, nil

first == second # true 

Is there another other cases for first being nil and second being anything else that can make first == second true? 

As far as I know, it can't be, but I'm not 100 % sure.
## [5][Low Hanging Fruits in Frontend Performance Optimization](https://www.reddit.com/r/ruby/comments/jrhy6c/low_hanging_fruits_in_frontend_performance/)
- url: https://pawelurbanek.com/frontend-performance-optimization
---

## [6][Ruby + appium + cucumber in AWS device farm setup](https://www.reddit.com/r/ruby/comments/jrlx0d/ruby_appium_cucumber_in_aws_device_farm_setup/)
- url: https://www.reddit.com/r/ruby/comments/jrlx0d/ruby_appium_cucumber_in_aws_device_farm_setup/
---
Hello,   
I'm trying to run my test suite in AWS device farm, I have successfully managed to run it on a custom environment. 

However, it feels as if cucumber is acting weird, it seems like some steps are skipped and sometimes the hooks don't run.

  
The whole behaviour of the suite seems really inconsistent, sometimes the test run sometimes they don't, even though they always run locally.  
Does anyone have experience with this type of set up for mobile automation?

Thank you very much for the answers :)
## [7][new football-sources gem - get football data via web pages or web api (json) calls (and convert to Football.CSV format)](https://www.reddit.com/r/ruby/comments/jriq7l/new_footballsources_gem_get_football_data_via_web/)
- url: https://github.com/yorobot/sport.db.more/tree/master/football-sources
---

## [8][How to GraphQL with Ruby, Rails,Active Record, and no N+1](https://www.reddit.com/r/ruby/comments/jr2748/how_to_graphql_with_ruby_railsactive_record_and/)
- url: https://evilmartians.com/chronicles/how-to-graphql-with-ruby-rails-active-record-and-no-n-plus-one
---

## [9][The Leaky Bucket rate limiter](https://www.reddit.com/r/ruby/comments/jr0pqb/the_leaky_bucket_rate_limiter/)
- url: https://www.mikeperham.com/2020/11/09/the-leaky-bucket-rate-limiter/
---

## [10][OOP(s) a Rake – write Rake tasks with Ruby objects](https://www.reddit.com/r/ruby/comments/jr3boh/oops_a_rake_write_rake_tasks_with_ruby_objects/)
- url: https://www.reddit.com/r/ruby/comments/jr3boh/oops_a_rake_write_rake_tasks_with_ruby_objects/
---
Hello 👋,

I made a mini-gem called [oops\_a\_rake](https://github.com/odlp/oops_a_rake) which I'd like to share. It lets you write Rake tasks as Ruby objects:

    class GreetingTask
      include OopsARake::Task
    
      description "An enthusiastic greeting"
    
      def call
        puts "Hello!"
      end
    end

This example would define a `greeting` task.

My motivation for writing this gem is to make testing Rake tasks easier ([more context](https://github.com/odlp/oops_a_rake#motivation)). I'd really appreciate any feedback or thoughts. Thanks!
