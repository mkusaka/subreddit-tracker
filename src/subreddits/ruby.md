# ruby
## [1][Has anyone here used DRb in the last few years? What was your experience?](https://www.reddit.com/r/ruby/comments/geihqa/has_anyone_here_used_drb_in_the_last_few_years/)
- url: https://www.reddit.com/r/ruby/comments/geihqa/has_anyone_here_used_drb_in_the_last_few_years/
---
I ended up browsing DRb's docs today. There isn't much written about it. Hence the curiosity whether anybody is using it.

I need a kind of micro-service that's part of the Rails app. A micro-service in the light that all puma/web-server/sidekiq instances talk to the same Class. I can imagine something like Redis but being part of Rails that I don't have to provision as a separate service. That would be straight-forward in Elixir/Phoenix, but there isn't a clear path in Ruby/Rails.

That's how I ended up reading about DRb... this obscure ruby tech.
## [2][Crystal yay or nay?](https://www.reddit.com/r/ruby/comments/ge4h03/crystal_yay_or_nay/)
- url: https://www.reddit.com/r/ruby/comments/ge4h03/crystal_yay_or_nay/
---
Just wondering what my fellow rubyists think of Crystal.
Do you see some potential in Crystal or you very much prefer the way Ruby works?

Crystal is young and full of flaws but that does not mean those cant be fixed.
## [3][Was tasked to write a function that reverses an array without using the reverse method or a second data structure, this was my solution.](https://www.reddit.com/r/ruby/comments/gdzksr/was_tasked_to_write_a_function_that_reverses_an/)
- url: https://i.redd.it/hdkidgtlqyw41.png
---

## [4]['RailsConf 2020.2 - Couch Edition' has been released, with 33 videos now available](https://www.reddit.com/r/ruby/comments/gdsw8y/railsconf_20202_couch_edition_has_been_released/)
- url: https://www.youtube.com/playlist?list=PLE7tQUdRKcyZ-TzxlxdLvh6tDUfZHqm76
---

## [5][Help beta test the Heroku "PlatformAPI" gem with Rate Throttling](https://www.reddit.com/r/ruby/comments/ge36pw/help_beta_test_the_heroku_platformapi_gem_with/)
- url: /r/Heroku/comments/ge32cr/help_beta_test_the_heroku_platformapi_gem_with/
---

## [6][Ruby Stream API](https://www.reddit.com/r/ruby/comments/ge0zl8/ruby_stream_api/)
- url: https://www.reddit.com/r/ruby/comments/ge0zl8/ruby_stream_api/
---
[https://github.com/ruby-ee/ruby-stream-api](https://github.com/ruby-ee/ruby-stream-api) 

Useful for both finite and (potentially) infinite streams of elements!
## [7][anyone have suggestions for turning a comma separated list of plain text (some entries with spaces) to an array of strings?](https://www.reddit.com/r/ruby/comments/ge5i6z/anyone_have_suggestions_for_turning_a_comma/)
- url: https://www.reddit.com/r/ruby/comments/ge5i6z/anyone_have_suggestions_for_turning_a_comma/
---
say I just have play text like

    High,
    Low,
    High Low,
    Medium

there are carriage returns to the new line so it looks like the above.


and I want to create an array of strings of it, I can't just remove the commends and replace them with spaces and do `%w(High Low High Low Medium)`, because that will split up the value "High Low" string that I want to "High" and "Low", meaning two "Low"s will be in the array.  

Any suggestions?

I also tried making it one big string

    "High,
    Low,
    High Low,
    Medium"


and doing `.split(',')` or `.split(',\n') but it doesn't give the desired result.  There are carriage returns to new lines hence I tried that \n
## [8][Exploring Method Arguments in Ruby (Part 3)](https://www.reddit.com/r/ruby/comments/gdzv6j/exploring_method_arguments_in_ruby_part_3/)
- url: https://www.ombulabs.com/blog/ruby/learning/methods-arguments-ruby-part3.html
---

## [9][RailsConf 2020: DHH Keynote](https://www.reddit.com/r/ruby/comments/gdnqxs/railsconf_2020_dhh_keynote/)
- url: https://www.youtube.com/watch?v=OltCr8AWpWw
---

## [10][Learn Ruby(Programming language) - YouTube](https://www.reddit.com/r/ruby/comments/gdutuo/learn_rubyprogramming_language_youtube/)
- url: https://www.youtube.com/playlist?list=PLHCuMScnJ-Q1HdY7KZUpTDuUDDW53K3zz
---

