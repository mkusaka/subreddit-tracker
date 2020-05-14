# ruby
## [1][Matz is calling for feedback on Ruby 2.7/3.0 keyword argument pain](https://www.reddit.com/r/ruby/comments/gjjayp/matz_is_calling_for_feedback_on_ruby_2730_keyword/)
- url: https://discuss.rubyonrails.org/t/new-2-7-3-0-keyword-argument-pain-point/74980
---

## [2][Thread scheduler for light weight concurrency. by ioquatix · Pull Request #3032 · was merged. :heart:](https://www.reddit.com/r/ruby/comments/gjjtla/thread_scheduler_for_light_weight_concurrency_by/)
- url: https://github.com/ruby/ruby/pull/3032
---

## [3][Database-driven authorization in Rails using CanCanCan - Abilities in DB and MetaProgramming](https://www.reddit.com/r/ruby/comments/gjkhs9/databasedriven_authorization_in_rails_using/)
- url: https://www.reddit.com/r/ruby/comments/gjkhs9/databasedriven_authorization_in_rails_using/
---
Hi ruby family,

As an initiative to give back to the community, I have started writing a series of blogs on ruby and ruby on rails. A few days back, I published a post on **authorizing resources in rails using CanCanCan**. As a continuation of the previous post, I have recently published another post on how to **implement database-driven authorization using CanCanCan**.

Some of the key issues that I tried to solve was :

1. The Growing size of the ability file
2. Abilities being hard to maintain.
3. Redeployment of the application for every change in the ability file
4. Storing abilities in a database

If you think this can be extrapolated and be made into a gem, let me know, and let's work together to create an awesome library.

[https://addytalks.tech/2020/05/14/rails-cancancan-database-driven-authorization/](https://addytalks.tech/2020/05/14/rails-cancancan-database-driven-authorization/)

You check out my previous post here - 

[https://addytalks.tech/2020/05/03/ruby-on-rails-authorization-with-cancancan/](https://addytalks.tech/2020/05/03/ruby-on-rails-authorization-with-cancancan/)
## [4][Need help coming up with some ideas](https://www.reddit.com/r/ruby/comments/gjhb5f/need_help_coming_up_with_some_ideas/)
- url: https://www.reddit.com/r/ruby/comments/gjhb5f/need_help_coming_up_with_some_ideas/
---
Hey guys and girls,

I'm doing a programming unit at uni where we are using ruby. We need to create a custom program, and I'm not very good at this course, I know the basics and that's about it. I was wondering if you guys would be able to give me some ideas on a program that I'd be able to make that wouldn't need to much work or learning involved.
## [5][corrupted rubygems manager?](https://www.reddit.com/r/ruby/comments/gjci58/corrupted_rubygems_manager/)
- url: https://www.reddit.com/r/ruby/comments/gjci58/corrupted_rubygems_manager/
---
My Lenovo ideapad (running Ubuntu 18.04 LTS) started acting strangely yesterday morning. I dropped it off with my local fixit shop to check for mechanical &amp; disk drive problems - clean bill of health.

When I got home, I decided to test things by creating a new Rails app and got the following load error:

&lt;internal:gem\_prelude&gt;:2:in \`require': cannot load such file -- rubygems.rb (LoadError)

Here's what I saw:

* $gem -h: --&gt; same error.
* I downloaded &amp; uppacked rubygems 3.1.3, thinking that the old version was somehow broken. $ruby setup.rb --&gt; same error.
* I re-installed ruby 2.5.1 (via rbenv), then $ruby setup.rb --&gt; same error.
* I changed the local ruby executable to 2.6.0 (via rbenv), then $ruby setup.rb --&gt; same error
* calling ruby with no cmndline args --&gt; same error

By now I was starting to get worried, and started the requisite Google search for hints. One hint came from [https://stackoverflow.com/questions/38577603/internalgem-prelude1in-require](https://stackoverflow.com/questions/38577603/internalgem-prelude1in-require) \- using $rvm fix-permissions. (I hadn't touched rvm in several years, in favor of rbenv.) Sure enough, this seems (based on creating a blank Rails 6.0.3 app) that the problem is solved.

But I'm not sure \*what\* I solved. Can anyone shed light on this behavior?
## [6][Animation of the SHA-256 hash function in your terminal by Greg Walker (of Learn Me a Bitcoin fame)](https://www.reddit.com/r/ruby/comments/giyavu/animation_of_the_sha256_hash_function_in_your/)
- url: https://github.com/in3rsha/sha256-animation
---

## [7][Ruby hashing algorithm could be improved using Tabulation Hashing](https://www.reddit.com/r/ruby/comments/giuzz6/ruby_hashing_algorithm_could_be_improved_using/)
- url: https://bugs.ruby-lang.org/issues/16851
---

## [8][Churn vs. Complexity vs. Code Coverage](https://www.reddit.com/r/ruby/comments/gizu6r/churn_vs_complexity_vs_code_coverage/)
- url: https://www.fastruby.io/blog/code-quality/churn-vs-complexity-vs-coverage.html
---

## [9][Upgrading a trivial Rails app from Ruby 2.3.1 to 2.7 and from Rails 4.2.6 to 6.0.3](https://www.reddit.com/r/ruby/comments/git6m0/upgrading_a_trivial_rails_app_from_ruby_231_to_27/)
- url: https://blog.arkency.com/upgrading-a-trivial-rails-app-from-ruby-2-dot-3-dot-1-to-2-dot-7-and-from-rails-4-dot-2-dot-6-to-6-dot-0-dot-3/
---

## [10][Question about Gosu?](https://www.reddit.com/r/ruby/comments/gitli7/question_about_gosu/)
- url: https://www.reddit.com/r/ruby/comments/gitli7/question_about_gosu/
---
I am trying to create a game using the Gosu library. The question is that i have multiple screens such as the welcome screen which the user must see once they run the program. The. I want to implement a functionality where when the user clicks on a button and it gets redirected to the screen where my actual game is. The part of redirecting is what I am confused with.
