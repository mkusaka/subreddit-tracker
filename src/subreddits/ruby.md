# ruby
## [1][Endless Ruby](https://www.reddit.com/r/ruby/comments/fz578c/endless_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fz578c/endless_ruby/
---
A very interesting feature was proposed to be added to Ruby. If you can imagine from the title - Yes! endless Ruby. But in what context? Let's see in this blog post. 

Method definition in Ruby is done this way :

```ruby
def some_method
# one-line method body
end 
```
The proposal is to allow the below syntax for method definitions

```ruby
def value(args) = expression
```

You see, there is no end keyword hence endless. 
This just applies to method definitions.

With this in place, the below examples can be considered

```ruby
def hello(name) = 
  puts "Hello #{name}!"

hello('endless Ruby') #=&gt; Hello endless Ruby!
```
and

```ruby
def square(x) = x * x

p square(4) #=&gt; 16
```
and 

```ruby
x = Object.new
def x.foo = "FOO"

p x.foo #=&gt; "FOO"
```

I felt it would a nice thing to add, as we often write one-line methods.

Discussions on this proposal are happening [here](https://bugs.ruby-lang.org/issues/16746). Feel free to visit and express your thoughts.
## [2][What's the best library for game programming in Ruby?](https://www.reddit.com/r/ruby/comments/fz3g5k/whats_the_best_library_for_game_programming_in/)
- url: https://www.reddit.com/r/ruby/comments/fz3g5k/whats_the_best_library_for_game_programming_in/
---
I only used `ruby2D`. Do your have any other suggestions? Thanks!
## [3][Open Source Progress Report: Falcon, Bake, Variant, Thread Safety and more!](https://www.reddit.com/r/ruby/comments/fyh6wz/open_source_progress_report_falcon_bake_variant/)
- url: https://www.codeotaku.com/journal/2020-04/open-source-progress-report/index
---

## [4][Web Scraping Like Your Time is Valuable](https://www.reddit.com/r/ruby/comments/fylba9/web_scraping_like_your_time_is_valuable/)
- url: https://medium.com/@c.andrewlong/web-scraping-like-your-time-is-valuable-65194b1dfe02
---

## [5][Why is there no .dmg type of Mac install for Ruby?](https://www.reddit.com/r/ruby/comments/fyvrt7/why_is_there_no_dmg_type_of_mac_install_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fyvrt7/why_is_there_no_dmg_type_of_mac_install_for_ruby/
---
It seems like it's easier to install Ruby on Windows than it is on Mac (rubyinstaller.org has come a LONG way!). I'm not saying rbenv/rvm/chruby are complicated, but it does require someone to have some understanding of their zsh/bash profile (I think).

I'm just looking to get some beginners started and wanted the easiest solution out there for it for their macbooks. Is a .dmg a possibility?
## [6][The Ruby Blend Episode 8: Tests and Webpacker](https://www.reddit.com/r/ruby/comments/fyhyx8/the_ruby_blend_episode_8_tests_and_webpacker/)
- url: https://fireside.fm/s/ouBAUjGy+F6MV5gnV
---

## [7][Been learning Ruby at a code camp for one week, felt like I haven't learned anything](https://www.reddit.com/r/ruby/comments/fyilhb/been_learning_ruby_at_a_code_camp_for_one_week/)
- url: https://www.reddit.com/r/ruby/comments/fyilhb/been_learning_ruby_at_a_code_camp_for_one_week/
---
I am participating in a code camp right now and we are learning Ruby. I'm one week into the camp, and honestly I felt like I haven't learned anything about Ruby outside of some vocabulary that I already most knew from codecademy and reading articles.  I feel like 50% of the time I am googling for answers, 40% of the time just staring at the screen, and 10% actually typing code.  I don't want to complain about the camp because the issue is 100% me as other students seem to be making progress, but out of the dozens of challenges I think there was maybe one I actually did by myself and the rest were solved by my partner or teacher.  We've touched on conditionals, arrays, hashes, iterators, blocks, etc., and while I understand what those elements are, actually implementing them into working code is beyond my grasp.  There are so many "how was I supposed to know that?" or "I would've never figured that out" moments.  I don't expect to be an expert after a week, but I don't think I could even complete the past exercises if I did them now.

I keep searching for "Ruby for beginners" types of articles and videos, but I feel most of them just rattle off concepts instead of actual useful strategies when it comes to writing code.

If you have had a similar experience with learning Ruby (even if you didn't go to a code camp), what did you do to help grasp the core concepts?  What are some things you wish you knew when first starting out?  No matter how the camp turns out I would still like to improve my Ruby skills in my free time and hopefully make something interesting in the future.
## [8][What installs activerecord-sqlserver-adapter?](https://www.reddit.com/r/ruby/comments/fyq7ud/what_installs_activerecordsqlserveradapter/)
- url: https://www.reddit.com/r/ruby/comments/fyq7ud/what_installs_activerecordsqlserveradapter/
---
# Hi,

I see that something is installing activerecord-sqlserver-adapter-5.2.1 on our server.

I did "gem dependency activerecord-sqlserver-adapter --reverse-dependencies"

and it shows:

\-----

Gem activerecord-sqlserver-adapter-5.2.1

activerecord (\~&gt; 5.2.0)

tiny\_tds (&gt;= 0)

\-----

However, activerecord doesn't require that gem.

tiny\_tds isn't specified anywhere in my app's Gemfile.

Why is that gem is installed? How can I find out?

My problem is there is a tiny\_tds gem comes with it and I cannot figure out why it wasn't needed before and our app is working perfectly fine on production without it.

Now tiny\_tds comes with it and it requires freetds and freetds-devel system packages.
## [9][Platformer game written in Ruby - Super Bombinhas](https://www.reddit.com/r/ruby/comments/fxqvdf/platformer_game_written_in_ruby_super_bombinhas/)
- url: https://www.reddit.com/r/ruby/comments/fxqvdf/platformer_game_written_in_ruby_super_bombinhas/
---
Hi!

I'd like to share a demo of my platformer game, it is [open source](https://github.com/victords/super-bombinhas) and completely written in Ruby, with the help of the [Gosu](https://www.libgosu.org/) and [MiniGL](https://github.com/victords/minigl) gems. Also, the Windows executable was created with the [Ocra](https://github.com/larsch/ocra/) gem.

Windows: [https://drive.google.com/open?id=1KWRfnx6eGbj37OK7dd6Pk8uVzRtIWWTx](https://drive.google.com/open?id=1KWRfnx6eGbj37OK7dd6Pk8uVzRtIWWTx)

Linux (Debian-based): [https://drive.google.com/open?id=1tjh\_4akR5b1QpAMqtG1zJHJ1YhW5i2T8](https://drive.google.com/open?id=1tjh_4akR5b1QpAMqtG1zJHJ1YhW5i2T8)

I'd appreciate any feedback. Thanks!
## [10][Minimalist tool to automate the writing of unit tests.](https://www.reddit.com/r/ruby/comments/fxevwc/minimalist_tool_to_automate_the_writing_of_unit/)
- url: https://github.com/fixrb/brutal
---

