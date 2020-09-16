# ruby
## [1][How To Publish Your First Rails API To Heroku](https://www.reddit.com/r/ruby/comments/itqjbd/how_to_publish_your_first_rails_api_to_heroku/)
- url: https://codepen.io/rvvergara/post/how-to-publish-your-first-rails-api-to-heroku
---

## [2][Rails Concerns: To Concern Or Not To Concern](https://www.reddit.com/r/ruby/comments/itulx6/rails_concerns_to_concern_or_not_to_concern/)
- url: https://blog.appsignal.com/2020/09/16/rails-concers-to-concern-or-not-to-concern.html
---

## [3][Ruby 2.8 adds endless method definition](https://www.reddit.com/r/ruby/comments/itrg19/ruby_28_adds_endless_method_definition/)
- url: https://blog.bigbinary.com/2020/09/15/ruby-2-8-adds-endless-method-definition.html
---

## [4][Function Composition: Python vs Ruby](https://www.reddit.com/r/ruby/comments/it8spw/function_composition_python_vs_ruby/)
- url: https://www.reddit.com/gallery/it8spw
---

## [5][Here is a list with a few job openings for ruby developers that auto-updates everyday. Maybe can help someone!](https://www.reddit.com/r/ruby/comments/it6uzw/here_is_a_list_with_a_few_job_openings_for_ruby/)
- url: https://docs.google.com/spreadsheets/d/e/2PACX-1vTjk1rnmRduqHpBl-3QxgGultV2FpKh96h8wfUx4gfKm5rfa1P7vyI5l8yVs7Q55K9yXAdIf6UstPyW/pubhtml
---

## [6][RFID reader/writer for raspberry pi using Ruby?](https://www.reddit.com/r/ruby/comments/ith8bc/rfid_readerwriter_for_raspberry_pi_using_ruby/)
- url: https://www.reddit.com/r/ruby/comments/ith8bc/rfid_readerwriter_for_raspberry_pi_using_ruby/
---
I got [clever card kit for raspberry pi ](https://github.com/simonmonk/clever_card_kit) from monkworks and the project code is in python. The python code works fine but I am trying to understand the lower level code better. I would like to rewrite the projects in Ruby.

I found [this Ruby gem](https://github.com/atitan/MFRC522_Ruby) but it does not work for me. I cannot figure out what is different other then the rst pin is 24 not 25

Anyone able to lend some guidance? I like to practice Ruby while playing with my pi but I am over my head when it comes to this hardware stuff.
## [7][football.db Gem Family 2020.9 Update for Season 2020/21 incl. all FIFA countries, world's top football leagues &amp; clubs](https://www.reddit.com/r/ruby/comments/it4038/footballdb_gem_family_20209_update_for_season/)
- url: https://github.com/sportdb/football.db
---

## [8][I wrote a Ruby circuit-breaker library for decoupling external dependencies and improving fault-tolerance in production applications at my company](https://www.reddit.com/r/ruby/comments/isrily/i_wrote_a_ruby_circuitbreaker_library_for/)
- url: https://github.com/ParentSquare/faulty
---

## [9][Another Reason to Avoid constantize in Rails](https://www.reddit.com/r/ruby/comments/isz2db/another_reason_to_avoid_constantize_in_rails/)
- url: https://blog.presidentbeef.com/blog/2020/09/14/another-reason-to-avoid-constantize-in-rails/
---

## [10][How might I run RSpec on files that have recently changed?](https://www.reddit.com/r/ruby/comments/it260y/how_might_i_run_rspec_on_files_that_have_recently/)
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
