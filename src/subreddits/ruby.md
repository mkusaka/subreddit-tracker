# ruby
## [1][ActiveSupport is so awesome](https://www.reddit.com/r/ruby/comments/fzbu6f/activesupport_is_so_awesome/)
- url: https://i.redd.it/nd2q35ss98s41.png
---

## [2][Implementing Exclusive Fibers by Samuel Williams](https://www.reddit.com/r/ruby/comments/fze924/implementing_exclusive_fibers_by_samuel_williams/)
- url: https://www.youtube.com/watch?v=cNaqbeagqUw&amp;feature=youtu.be
---

## [3][Endless Ruby](https://www.reddit.com/r/ruby/comments/fz578c/endless_ruby/)
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
## [4][Hey guys is there any good exercise to do pass the interview test algorithms with their solutions? Ruby is preferred](https://www.reddit.com/r/ruby/comments/fzmo53/hey_guys_is_there_any_good_exercise_to_do_pass/)
- url: https://www.reddit.com/r/ruby/comments/fzmo53/hey_guys_is_there_any_good_exercise_to_do_pass/
---

## [5][How I Built a React Templating Tool in Ruby](https://www.reddit.com/r/ruby/comments/fz8xi8/how_i_built_a_react_templating_tool_in_ruby/)
- url: https://link.medium.com/EzTIbPooB5
---

## [6][What's the best library for game programming in Ruby?](https://www.reddit.com/r/ruby/comments/fz3g5k/whats_the_best_library_for_game_programming_in/)
- url: https://www.reddit.com/r/ruby/comments/fz3g5k/whats_the_best_library_for_game_programming_in/
---
I only used `ruby2D`. Do your have any other suggestions? Thanks!
## [7][Open Source Progress Report: Falcon, Bake, Variant, Thread Safety and more!](https://www.reddit.com/r/ruby/comments/fyh6wz/open_source_progress_report_falcon_bake_variant/)
- url: https://www.codeotaku.com/journal/2020-04/open-source-progress-report/index
---

## [8][Web Scraping Like Your Time is Valuable](https://www.reddit.com/r/ruby/comments/fylba9/web_scraping_like_your_time_is_valuable/)
- url: https://medium.com/@c.andrewlong/web-scraping-like-your-time-is-valuable-65194b1dfe02
---

## [9][Why is there no .dmg type of Mac install for Ruby?](https://www.reddit.com/r/ruby/comments/fyvrt7/why_is_there_no_dmg_type_of_mac_install_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fyvrt7/why_is_there_no_dmg_type_of_mac_install_for_ruby/
---
It seems like it's easier to install Ruby on Windows than it is on Mac (rubyinstaller.org has come a LONG way!). I'm not saying rbenv/rvm/chruby are complicated, but it does require someone to have some understanding of their zsh/bash profile (I think).

I'm just looking to get some beginners started and wanted the easiest solution out there for it for their macbooks. Is a .dmg a possibility?
## [10][The Ruby Blend Episode 8: Tests and Webpacker](https://www.reddit.com/r/ruby/comments/fyhyx8/the_ruby_blend_episode_8_tests_and_webpacker/)
- url: https://fireside.fm/s/ouBAUjGy+F6MV5gnV
---

