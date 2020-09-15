# ruby
## [1][Here is a list with a few job openings for ruby developers that auto-updates everyday. Maybe can help someone!](https://www.reddit.com/r/ruby/comments/it6uzw/here_is_a_list_with_a_few_job_openings_for_ruby/)
- url: https://docs.google.com/spreadsheets/d/e/2PACX-1vTjk1rnmRduqHpBl-3QxgGultV2FpKh96h8wfUx4gfKm5rfa1P7vyI5l8yVs7Q55K9yXAdIf6UstPyW/pubhtml
---

## [2][I wrote a Ruby circuit-breaker library for decoupling external dependencies and improving fault-tolerance in production applications at my company](https://www.reddit.com/r/ruby/comments/isrily/i_wrote_a_ruby_circuitbreaker_library_for/)
- url: https://github.com/ParentSquare/faulty
---

## [3][football.db Gem Family 2020.9 Update for Season 2020/21 incl. all FIFA countries, world's top football leagues &amp; clubs](https://www.reddit.com/r/ruby/comments/it4038/footballdb_gem_family_20209_update_for_season/)
- url: https://github.com/sportdb/football.db
---

## [4][Another Reason to Avoid constantize in Rails](https://www.reddit.com/r/ruby/comments/isz2db/another_reason_to_avoid_constantize_in_rails/)
- url: https://blog.presidentbeef.com/blog/2020/09/14/another-reason-to-avoid-constantize-in-rails/
---

## [5][How might I run RSpec on files that have recently changed?](https://www.reddit.com/r/ruby/comments/it260y/how_might_i_run_rspec_on_files_that_have_recently/)
- url: https://www.reddit.com/r/ruby/comments/it260y/how_might_i_run_rspec_on_files_that_have_recently/
---
I'm working on a project that has a LOT of files under `spec/`. Rather than typing out the full path to files I want to run when I want to run _just that test_, I'm lazy so I run `bundle exec rake spec` and it runs them all.

Well that's not efficient right? But I want to _stay lazy_. So I'd like to be able to run a command on files that only recently changed. But I don't want this done _automatically_, only via a rake task when I call it. I'm basically trying to create a one-liner "fire and forget" command to run specs that "have changed" since the last time this command was run.

So conceptually, this might be something like:

1. Run rake task, gather up `mtime` on each `_spec.rb` file
1. Save that somewhere
1. Run rake task again, and look for changed `mtime` and run specs for only those files

Does anyone know of a **simple** way to implement something like this?

I'd like to stay away from anything that shells out to git, or something like that, if possible. I have my reasons.

Appreciate your thoughts. Thank you.
## [6][Function Composition: Python vs Ruby](https://www.reddit.com/r/ruby/comments/it8spw/function_composition_python_vs_ruby/)
- url: https://www.reddit.com/gallery/it8spw
---

## [7][dry-system - why? | from manual dependency injection to a full-blown architecture](https://www.reddit.com/r/ruby/comments/isgjfz/drysystem_why_from_manual_dependency_injection_to/)
- url: https://www.youtube.com/watch?v=BQkAGgSCoZ4
---

## [8][cant run ruby files](https://www.reddit.com/r/ruby/comments/issvkl/cant_run_ruby_files/)
- url: https://www.reddit.com/r/ruby/comments/issvkl/cant_run_ruby_files/
---
Hello friend,

I am trying to run the following command on ruby in terminal. I am using a framework that's called "ruby2d" the file is called "rub.rb" when i do the following command in terminal "ruby rub.rb" it gives me an error saying "no file found" I am running windows 10. Here is the code...

&amp;#x200B;

require 'ruby2d'

&amp;#x200B;

show
## [9][3 ways to make class methods private in Ruby](https://www.reddit.com/r/ruby/comments/isfick/3_ways_to_make_class_methods_private_in_ruby/)
- url: https://medium.com/rubycademy/3-ways-to-make-class-methods-private-in-ruby-64b970e54613
---

## [10][Results form the 2020 Ruby on Rails developer community survey](https://www.reddit.com/r/ruby/comments/is35ga/results_form_the_2020_ruby_on_rails_developer/)
- url: https://rails-hosting.com/2020/
---

