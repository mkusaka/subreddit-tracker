# ruby
## [1][Postgres Indexes for ActiveRecord Join Tables in Rails Apps](https://www.reddit.com/r/ruby/comments/hea2rf/postgres_indexes_for_activerecord_join_tables_in/)
- url: https://pawelurbanek.com/rails-postgres-join-indexes
---

## [2][How to Make a Perfect Pull Request](https://www.reddit.com/r/ruby/comments/hdzfeb/how_to_make_a_perfect_pull_request/)
- url: https://www.reddit.com/r/ruby/comments/hdzfeb/how_to_make_a_perfect_pull_request/
---
Once I had decided to improve my team’s performance and the first step was to find the bottleneck.

It turned out that coding speed wasn’t the limiting factor — code review was.

So, to speed up reviews, I compared two types of pull requests:

* Those that receive few comments and get merged quickly
* Those that get a lot of comments and require several review rounds.

Here’s the result: 9 ways to make pull requests easier to review. [https://medium.com/better-programming/how-to-make-a-perfect-pull-request-3578fb4c112](https://medium.com/better-programming/how-to-make-a-perfect-pull-request-3578fb4c112?source=friends_link&amp;sk=b326bf34cab38601f1fb994206b33137)

https://preview.redd.it/k2q8ba88oi651.jpg?width=4032&amp;format=pjpg&amp;auto=webp&amp;s=9fe61b1c1fa5ee10f73de999561be76896a8b746
## [3][Strings -- how does `\b` work?](https://www.reddit.com/r/ruby/comments/hee8tu/strings_how_does_b_work/)
- url: https://www.reddit.com/r/ruby/comments/hee8tu/strings_how_does_b_work/
---
source:

`s = "string"`

`s.size.times { |index| print(s[index]); sleep(0.01) }`  
`s.size.times { print("\b") }`

and this is ok, but if I replace the first `print()` with `puts()`, I get `s` printed out vertically and nothing deleted.  
basically, "\\b" can't delete new lines, is there a specific format character for doing that?
## [4][Building GitHub-style Hovercards with Stimulus and HTML-over-the-wire](https://www.reddit.com/r/ruby/comments/hdsmwj/building_githubstyle_hovercards_with_stimulus_and/)
- url: https://boringrails.com/articles/hovercards-stimulus/
---

## [5][Opal-Async Gem Maintenance Going Forward](https://www.reddit.com/r/ruby/comments/he6mtg/opalasync_gem_maintenance_going_forward/)
- url: https://www.reddit.com/r/ruby/comments/he6mtg/opalasync_gem_maintenance_going_forward/
---
[https://andymaleh.blogspot.com/2020/06/opal-async-gem-maintainer-going-forward.html](https://andymaleh.blogspot.com/2020/06/opal-async-gem-maintainer-going-forward.html)
## [6][Graceful Request Retries in Ruby Applications](https://www.reddit.com/r/ruby/comments/hdr91t/graceful_request_retries_in_ruby_applications/)
- url: https://medium.com/@kirill_shevch/graceful-request-retries-in-ruby-applications-7bbeac5ebd40
---

## [7][ruby on arm?](https://www.reddit.com/r/ruby/comments/hdnmn9/ruby_on_arm/)
- url: https://www.reddit.com/r/ruby/comments/hdnmn9/ruby_on_arm/
---
given that mac will move to arm procs
aws has his on arm processor
and somehow the performance is good(https://www.anandtech.com/show/15578/cloud-clash-amazon-graviton2-arm-against-intel-and-amd/10)

- what is the status of ruby on arm?
- can ruby leverage here, or again we are years behind all?
- what the community could do to improve it?
- what is the story with jruby, truffleruby on arm?
## [8][I used 7 as the input and it still outputs hi. How would I be able to fix this](https://www.reddit.com/r/ruby/comments/hdur6w/i_used_7_as_the_input_and_it_still_outputs_hi_how/)
- url: https://i.redd.it/x6h24pz0gh651.png
---

## [9][What are People Using for Localizing Date/Time?](https://www.reddit.com/r/ruby/comments/hdkw5t/what_are_people_using_for_localizing_datetime/)
- url: /r/rails/comments/hdkurj/what_are_people_using_for_localizing_datetime/
---

## [10][JRuby](https://www.reddit.com/r/ruby/comments/hdd1dw/jruby/)
- url: https://www.reddit.com/r/ruby/comments/hdd1dw/jruby/
---
So if I interpreted this correctly (no pun intended) I can use JRuby to write both Java and Ruby code?
