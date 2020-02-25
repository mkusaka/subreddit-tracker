# ruby
## [1][Ruby on Rails Mistakes could kill Your Production Servers](https://www.reddit.com/r/ruby/comments/f96x3n/ruby_on_rails_mistakes_could_kill_your_production/)
- url: https://pawelurbanek.com/rails-mistakes-downtime
---

## [2][Deep dive into rackup](https://www.reddit.com/r/ruby/comments/f978se/deep_dive_into_rackup/)
- url: https://deepdive.sh/deep-dive-into-rackup
---

## [3][Rails 6 fixes a bug where after_commit callbacks are called on failed update in a transaction block | The Official BigBinary Blog | BigBinary](https://www.reddit.com/r/ruby/comments/f997tk/rails_6_fixes_a_bug_where_after_commit_callbacks/)
- url: https://blog.bigbinary.com/2020/02/25/rails-6-fixes-a-bug-where-after_commit-callbacks-are-called-on-failed-update-in-a-transaction-block.html
---

## [4][Hate for Medium](https://www.reddit.com/r/ruby/comments/f927ed/hate_for_medium/)
- url: https://www.reddit.com/r/ruby/comments/f927ed/hate_for_medium/
---
Several times, I've considered biting the bullet, putting my hand in my pocket, ponying up the $5, or whatever it is, for a Medium subscription. Some of the content is pretty good. I've enjoyed content by /u/mehdifarsi, /u/eric_programmer, /u/vfreefly and others in the past.

But, today, the top article in my *Medium Daily Digest* points me to \***medium link alert**\* - [Tracking User Streaks in Ruby](https://levelup.gitconnected.com/tracking-user-streaks-in-ruby-a49e90ce46a1). If you don't want to follow that link, fair enough, but in it, the author presents a 14 line poorly formatted method containing errors (both syntactic and semantic) that purports to count user activity streaks.

Now, I don't want to hate on the author - if this post was on a personal blog or a free aggregation site, I'd be happy to comment with some positive constructive criticism - but for a service that pesters me regularly for $5 per month to deliver this kind of "Ideas and perspectives you won’t find anywhere else" from "the world’s most insightful writers, thinkers, and storytellers" and call it "the smartest takes on topics that matter" is an outright scam!

If this outfit is taking $5/month off what must be an awful lot of people, surely they can afford to pay some subject matter experts to do the most cursory of proof reading? Can't they?

Anyway, rant over. Happy Birthday Ruby 🎂 🎉
## [5][Happy Birthday, Ruby!](https://www.reddit.com/r/ruby/comments/f8t108/happy_birthday_ruby/)
- url: https://twitter.com/yukihiro_matz/status/1231730031589675008?s=21
---

## [6][Pass STDOUT/ERR/IN via socket to a ruby process from any fast language ?](https://www.reddit.com/r/ruby/comments/f96kkp/pass_stdouterrin_via_socket_to_a_ruby_process/)
- url: https://www.reddit.com/r/ruby/comments/f96kkp/pass_stdouterrin_via_socket_to_a_ruby_process/
---
I want to pass IO handles to a socket, so the process on the other end can take them over (for a faster / more generic [spring](https://rubygems.org/gems/spring) alternative)

Atm I'm using this, which takes \~100ms and is very simple:

ruby --disable-gems -rsocket -e "s = UNIXSocket.new('$socket'); s.send\_io STDOUT; s.send\_io STDERR; s.send\_io STDIN" [on github](https://github.com/grosser/ruby-cli-daemon/blob/master/bin/ruby-cli-daemon.sh#L56-L64)

[send\_io source](https://apidock.com/ruby/UNIXSocket/send_io) but I was unable to re-create it via awk/bash/perl etc (any language that is pre-installed on most users machines), only the ruby version seems to work

Any idea how to make this faster ?
## [7][Senior Ruby Developer Job in NYC - 50% Remote](https://www.reddit.com/r/ruby/comments/f8y82r/senior_ruby_developer_job_in_nyc_50_remote/)
- url: https://www.reddit.com/r/ruby/comments/f8y82r/senior_ruby_developer_job_in_nyc_50_remote/
---
 Must live in Greater NYC area - Can work from home 2-3 days a week.

This is an opportunity to build and design microservices, as well as convert a monolith into microservices architecture. The company currently deals with over 24 million different online inventory items every day and is scaling quickly.

Details: - Allows for 2-3 days a week remote (work from home)  
Salary: 150K – 170K for senior level

They are an eCommerce Platform that is used both B2B and C2C with millions of ticket listings with the largest ticket partners in the world. They are looking for a Senior Ruby Developer to help them migrate a ruby monolith to microservices. This person will be heavy on writing APIs.

Tech Stack:  
\- Ruby  
\- AWS  
\- Docker  
\- Kubernetes  
\- Microservices

Must be a permanent resident of the United States to be considered. Send your resume to [ian.murphy@thirdrepublic.com](mailto:ian.murphy@thirdrepublic.com) to learn more
## [8][do ya ever wish `retry` took a max tries arg?](https://www.reddit.com/r/ruby/comments/f8txf1/do_ya_ever_wish_retry_took_a_max_tries_arg/)
- url: https://www.reddit.com/r/ruby/comments/f8txf1/do_ya_ever_wish_retry_took_a_max_tries_arg/
---
I just realized a ruby feature I *really* wish existed is that the `retry` keyword could take an argument for "max retries", like `retry(3)` or something. 

I almost always want to do this, it almost never seems safe to retry without a max retries guard, seems risking infinite loops. And trying to implement a max-retries myself generally involves so much apparatus (sometimes having to add a really ugly `current_retry_count` arg to a method that really ought not to be public api!), that it basically eliminates any advantage of the built-in `retry` keyword in the first place, I might as well not be using it I've had to (re)implement the convenient parts. 

Thoughts on `retry`?  Who likes it and uses it all the time?

I suspect `retry` may be a little-known and little-used ruby keyword... I wonder if it's because of a lack of max-tries? Or if it'd end up not being all that useful anyway? Or if I'm wrong and people use it all the time already?

----

https://docs.ruby-lang.org/en/2.4.0/syntax/exceptions_rdoc.html

https://blog.appsignal.com/2018/06/05/redo-retry-next.html

https://www.honeybadger.io/blog/how-to-try-again-when-exceptions-happen-in-ruby/

http://blog.mirthlab.com/2012/05/25/cleanly-retrying-blocks-of-code-after-an-exception-in-ruby/
## [9][Setting up Rails project with RSpec - #1 testing framework from scratch](https://www.reddit.com/r/ruby/comments/f8p1ik/setting_up_rails_project_with_rspec_1_testing/)
- url: https://hixonrails.com/ruby-on-rails-tutorials/ruby-on-rails-testing-rspec-configuration/
---

## [10][Hash#shift using default value](https://www.reddit.com/r/ruby/comments/f8puzj/hashshift_using_default_value/)
- url: https://medium.com/@farsi_mehdi/hash-shift-using-default-value-b0d7cb62ffa2
---

