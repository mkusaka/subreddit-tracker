# ruby
## [1][Scope Gates in Ruby: Flat Scopes](https://www.reddit.com/r/ruby/comments/j92h27/scope_gates_in_ruby_flat_scopes/)
- url: https://medium.com/rubycademy/scope-gates-in-ruby-flat-scopes-bbf100f8e459
---

## [2][Middleman init fails with BigDecimal error](https://www.reddit.com/r/ruby/comments/j8thab/middleman_init_fails_with_bigdecimal_error/)
- url: https://www.reddit.com/r/ruby/comments/j8thab/middleman_init_fails_with_bigdecimal_error/
---
I just ran `middleman init` with Ruby 2.7.2 and hit an Active Support BigDecimal error "undefined method `new' for BigDecimal:Class". Anyone have any idea what this is all about?
## [3][hubba gem - little github api helper for (auto-)generating statistics reports (summary, stars, timeline, updates)](https://www.reddit.com/r/ruby/comments/j8fz13/hubba_gem_little_github_api_helper_for/)
- url: https://github.com/rubycoco/git/tree/master/hubba
---

## [4][Remote Ruby: Adam Wathan joins to talk TailwindCSS, Tailwind UI, and View Components](https://www.reddit.com/r/ruby/comments/j7zouw/remote_ruby_adam_wathan_joins_to_talk_tailwindcss/)
- url: https://share.transistor.fm/s/7bd5b843
---

## [5][Can someone help me out with what this does rand(999_999)](https://www.reddit.com/r/ruby/comments/j7vkhg/can_someone_help_me_out_with_what_this_does/)
- url: https://www.reddit.com/r/ruby/comments/j7vkhg/can_someone_help_me_out_with_what_this_does/
---
I'm trying to generate a fixed digit random number, and I came across this code. Could someone help me with what is happening here?  


I'm not sure i get much from the documentation for rand either.
## [6][Ruby 2.7.2 has revised Rdoc for Hash](https://www.reddit.com/r/ruby/comments/j7l51e/ruby_272_has_revised_rdoc_for_hash/)
- url: https://www.reddit.com/r/ruby/comments/j7l51e/ruby_272_has_revised_rdoc_for_hash/
---
I've spent a good part of this year working on revisions to the RDoc for Hash.  The revisions include all methods documentation in addition to much of the introductory text (class documentation).  Lots more example code.

Check it out:

\- 2.7.2: [https://ruby-doc.org/core-2.7.2/Hash.html](https://ruby-doc.org/core-2.7.2/Hash.html)

\- 2.7.1: [https://ruby-doc.org/core-2.7.1/Hash.html](https://ruby-doc.org/core-2.7.1/Hash.html)
## [7][Multithreading -- Parallelizing loops](https://www.reddit.com/r/ruby/comments/j80sz0/multithreading_parallelizing_loops/)
- url: https://www.reddit.com/r/ruby/comments/j80sz0/multithreading_parallelizing_loops/
---
I am looking for something like Julia's `Threads.@threads` macro, which allows to parallelize a for loop within multithreading, now I am looking for something like this for Ruby.
## [8][Fastest way to generating a n digit number](https://www.reddit.com/r/ruby/comments/j7w5xn/fastest_way_to_generating_a_n_digit_number/)
- url: https://www.reddit.com/r/ruby/comments/j7w5xn/fastest_way_to_generating_a_n_digit_number/
---
I found this code to generate 6 digit number 

    rand(100000...999999)

Is this approach fine?   
I'm not sure what the two dots(..) or three dots(...) mean and I hope it doesn't expand into an array and then rand selects a values
## [9][Speeding up Rails with Memoization](https://www.reddit.com/r/ruby/comments/j7v6l2/speeding_up_rails_with_memoization/)
- url: https://www.honeybadger.io/blog/ruby-rails-memoization/?utm_source=rubyweekly&amp;utm_medium=email&amp;utm_campaign=ruby
---

## [10][Beginner need some help understaind data pipeline](https://www.reddit.com/r/ruby/comments/j7vlbx/beginner_need_some_help_understaind_data_pipeline/)
- url: https://www.reddit.com/r/ruby/comments/j7vlbx/beginner_need_some_help_understaind_data_pipeline/
---
Context:  
 I am pulling in data from a social media site and storing it in my database. For eg. Everday I make an api req to fetch all tweets of a user. Then make another api req for each tweet to fetch its public\_metrics like tweets, likes,etc.   


If I am using Sidekiq, how should I build my data pipeline?  


Right now, I have a worker who just makes both the api calls from the same worker.  
Should I separate, "fetching all tweets" into a seprate worker and "fetch tweets metrics" into a separate wrker?
