# ruby
## [1][Low Hanging Fruits in Frontend Performance Optimization](https://www.reddit.com/r/ruby/comments/jrhy6c/low_hanging_fruits_in_frontend_performance/)
- url: https://pawelurbanek.com/frontend-performance-optimization
---

## [2][How to GraphQL with Ruby, Rails,Active Record, and no N+1](https://www.reddit.com/r/ruby/comments/jr2748/how_to_graphql_with_ruby_railsactive_record_and/)
- url: https://evilmartians.com/chronicles/how-to-graphql-with-ruby-rails-active-record-and-no-n-plus-one
---

## [3][new football-sources gem - get football data via web pages or web api (json) calls (and convert to Football.CSV format)](https://www.reddit.com/r/ruby/comments/jriq7l/new_footballsources_gem_get_football_data_via_web/)
- url: https://github.com/yorobot/sport.db.more/tree/master/football-sources
---

## [4][The Leaky Bucket rate limiter](https://www.reddit.com/r/ruby/comments/jr0pqb/the_leaky_bucket_rate_limiter/)
- url: https://www.mikeperham.com/2020/11/09/the-leaky-bucket-rate-limiter/
---

## [5][OOP(s) a Rake – write Rake tasks with Ruby objects](https://www.reddit.com/r/ruby/comments/jr3boh/oops_a_rake_write_rake_tasks_with_ruby_objects/)
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
## [6][I'm a trainee ruby developer with some learning budget left, where should I spend it?](https://www.reddit.com/r/ruby/comments/jr8cyt/im_a_trainee_ruby_developer_with_some_learning/)
- url: https://www.reddit.com/r/ruby/comments/jr8cyt/im_a_trainee_ruby_developer_with_some_learning/
---
I still have about £300 left on my learning budget this year from my company, what were some good books/courses you did that is worth mentioning?
## [7][Specific Practice?](https://www.reddit.com/r/ruby/comments/jr74u9/specific_practice/)
- url: https://www.reddit.com/r/ruby/comments/jr74u9/specific_practice/
---
Is there a website or something that lets your practice certain sections of ruby? For example if I want to test passing through classes, has many relationships, etc... is there a place I can practice just those til I’m confident to move onto the next section?
## [8][What Is The Purpose Of Private?](https://www.reddit.com/r/ruby/comments/jqx0g5/what_is_the_purpose_of_private/)
- url: https://www.tomdalling.com/blog/software-design/what-is-the-purpose-of-private/
---

## [9][Introduction to Rails testing with Rspec and Capybara](https://www.reddit.com/r/ruby/comments/jqzyj8/introduction_to_rails_testing_with_rspec_and/)
- url: https://rubyyagi.com/intro-rspec-capybara-testing/
---

## [10][Can I run two versions of Ruby at the same time so old gem dependencies work?](https://www.reddit.com/r/ruby/comments/jr56kf/can_i_run_two_versions_of_ruby_at_the_same_time/)
- url: https://www.reddit.com/r/ruby/comments/jr56kf/can_i_run_two_versions_of_ruby_at_the_same_time/
---
I'm updating to Ruby 2.6 from 2.3, but some of my gem dependencies don't work in 2.6. I've been using URU to switch between Ruby versions and it got me thinking, are there any tools to run 2.3 and 2.6 at the same time with their associated gems?
