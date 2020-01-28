# ruby
## [1][How to Perform Concurrent HTTP Requests in Ruby and Rails](https://www.reddit.com/r/ruby/comments/ev2m2q/how_to_perform_concurrent_http_requests_in_ruby/)
- url: https://pawelurbanek.com/ruby-concurrent-requests
---

## [2][Wrangling slow reports and large file exports in Rails with ActiveJob](https://www.reddit.com/r/ruby/comments/euw03c/wrangling_slow_reports_and_large_file_exports_in/)
- url: https://boringrails.com/articles/large-exports-and-slow-reports-with-activejob/
---

## [3][Sidekiq-like background job library in Go](https://www.reddit.com/r/ruby/comments/ev54rg/sidekiqlike_background_job_library_in_go/)
- url: https://www.reddit.com/r/ruby/comments/ev54rg/sidekiqlike_background_job_library_in_go/
---
Asynq is a simple Go library for queueing tasks and processing them in the background with workers.  
It is backed by Redis and it is designed to have a low barrier to entry. It should be integrated in your web stack easily.

[github.com/hibiken/asynq](https://github.com/hibiken/asynq)
## [4][How to pass arguments to methods in ruby and how it affects their arity](https://www.reddit.com/r/ruby/comments/euoj8j/how_to_pass_arguments_to_methods_in_ruby_and_how/)
- url: https://blog.saeloun.com/2020/01/27/how-to-pass-arguments-to-methods-in-ruby-and-how-it-affects-the-arity
---

## [5][StrongVersions - validate/auto-correct Gemfile for Semantic Version best practices](https://www.reddit.com/r/ruby/comments/eut8o1/strongversions_validateautocorrect_gemfile_for/)
- url: https://www.reddit.com/r/ruby/comments/eut8o1/strongversions_validateautocorrect_gemfile_for/
---
Hi. I am the author of a gem called StrongVersions. I have been using it successfully on quite a number of projects and thought it may be worth sharing in case others find it useful. The gem provides a command which validates the version definitions in a Gemfile (more info in links below).

I find it especially useful in continuous integration builds when working with large teams.

RubyGems:  [https://rubygems.org/gems/strong\_versions](https://rubygems.org/gems/strong_versions) 

Github:  [https://github.com/bobf/strong\_versions](https://github.com/bobf/strong_versions) 

Bug reports/suggestions are welcome. Have a nice day.
## [6][Evaluating Ruby in Ruby](https://www.reddit.com/r/ruby/comments/euasxs/evaluating_ruby_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/euasxs/evaluating_ruby_in_ruby/
---
[https://ilyabylich.svbtle.com/evaluating-ruby-in-ruby](https://ilyabylich.svbtle.com/evaluating-ruby-in-ruby)
## [7][tty-box ignoring everything past a certain point.](https://www.reddit.com/r/ruby/comments/euh1dt/ttybox_ignoring_everything_past_a_certain_point/)
- url: https://www.reddit.com/r/ruby/comments/euh1dt/ttybox_ignoring_everything_past_a_certain_point/
---
Anyone work with tty-box? It seems to be ignoring everything past the title: line. It prints the box to the console, shows the title in both area's, but ignores everything past the first costomzation. 

Is it the formatting I'm using? There's no good examples in the documentation that has multple comtomizations applied. 

    def title_box
      main_title = TTY::Box.frame(
        width: 30, height: 10,
        title: {top_left: 'Title', bottom_right: 'v0.01',
        border: {
        top_left: :cross,
        top_right: :cross,
        bottom_left: :cross,
        bottom_right: :cross,
        style: {
      fg: :bright_yellow,
      bg: :blue,
      border: {
        fg: :bright_yellow,
        bg: :blue
        }}}
      }
    )
    
        print main_title
          main
      end
## [8][Ruby variable or method](https://www.reddit.com/r/ruby/comments/eua27z/ruby_variable_or_method/)
- url: https://medium.com/@igor04/ruby-variable-or-method-ac6814b4e8b1
---

## [9][planet.rb quick starter script updated (v1.0) - add articles, blogs to your static (jekyll &amp; friends) website via feeds (and planet pluto)](https://www.reddit.com/r/ruby/comments/euch1c/planetrb_quick_starter_script_updated_v10_add/)
- url: https://github.com/feedreader/planet.rb
---

## [10][A Template for Starting New Projects on Ruby on Rails](https://www.reddit.com/r/ruby/comments/eu683v/a_template_for_starting_new_projects_on_ruby_on/)
- url: https://hackernoon.com/a-template-for-starting-new-projects-on-ruby-on-rails-t66c36iz
---

