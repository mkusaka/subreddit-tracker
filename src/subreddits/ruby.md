# ruby
## [1][Maintaining Sanity with Ruby Under A Lockdown.](https://www.reddit.com/r/ruby/comments/fs535x/maintaining_sanity_with_ruby_under_a_lockdown/)
- url: https://emmanuelhayford.com/maintaining-sanity-with-ruby-under-a-lockdown/
---

## [2][Exercising During Lockdown: a look at Ruby’s Array each_cons and each_slice methods](https://www.reddit.com/r/ruby/comments/fsa6wv/exercising_during_lockdown_a_look_at_rubys_array/)
- url: https://www.mesomorphic.co.uk/blog/2020/03/31/lockdown-push-up/
---

## [3][Learn Ruby on Rails](https://www.reddit.com/r/ruby/comments/fs8uhf/learn_ruby_on_rails/)
- url: https://medium.com/rubycademy/learn-ruby-on-rails-28f5ceb609f8?sk=16eda3bb701f75eff2b37165b5466a29
---

## [4][How to Rspec this?](https://www.reddit.com/r/ruby/comments/fsbowg/how_to_rspec_this/)
- url: /r/rails/comments/fsbo6p/how_to_rspec_this/
---

## [5][If you use Gruvbox as color scheme with vim/nvim, please check this](https://www.reddit.com/r/ruby/comments/fsatdz/if_you_use_gruvbox_as_color_scheme_with_vimnvim/)
- url: https://www.reddit.com/r/ruby/comments/fsatdz/if_you_use_gruvbox_as_color_scheme_with_vimnvim/
---
I am wondering if this is happening to others or it is something with my setup.

[https://github.com/morhetz/gruvbox/issues/323](https://github.com/morhetz/gruvbox/issues/323)
## [6][The first version of dry-rails was released!](https://www.reddit.com/r/ruby/comments/frofpd/the_first_version_of_dryrails_was_released/)
- url: https://www.reddit.com/r/ruby/comments/frofpd/the_first_version_of_dryrails_was_released/
---
If you're using Rails and you'd like to try out [dry-rb](https://dry-rb.org/) then I have great news for you - the first version of dry-rails was released today and it provides built-in support for dry-system that works OOTB with both Rails 5 and 6 (works with Zeitwerk enabled too). It also gives you additional features that are integrated with the controller layer:

* `safe_params` - an alternative to strong-parameters, powered by dry-schema
* `application_contract` - an alternative to ActiveModel/Record validations powered by dry-validation
* `controller_helpers` - a couple of convenient methods that make it easier to access your application container

## Links

* [0.1.0 release on GitHub](https://github.com/dry-rb/dry-rails/releases/tag/v0.1.0) and feel
* [API docs](https://rubydoc.info/gems/dry-rails)
* [User docs](https://dry-rb.org/gems/dry-rails)

Check it out, provide feedback and feel free to ask any questions!
## [7][How do you properly test a selenium web scrapper?](https://www.reddit.com/r/ruby/comments/fs0442/how_do_you_properly_test_a_selenium_web_scrapper/)
- url: https://www.reddit.com/r/ruby/comments/fs0442/how_do_you_properly_test_a_selenium_web_scrapper/
---
I'm writing a scraper that uses selenium to navigate &amp; login to a certain website; search for the newest data and then store it into a database. I'm using selenium-webdriver to navigate the website, and now I'm trying to write tests for the most important edge cases.

I downloaded the HTML and built a fake Sinatra website, that mimics the behavior of the original site so that I can test my code. However, I have to run the puma server separately in an environment independent of my code. 

I need to be able to mock everything in the same environment so that I can have better control of how the application behaves. I think I can take the same approach as the guys from Capybara do but I don't know how to start.

I created a short mocking class, and it runs but as soon as puma starts RSpec is halted waiting for puma to stop its execution.

What's the best approach that I can take to actually test this scraper correctly, are there any technologies that already exist and that I can use?

Edit: Grammar
## [8][Main usage for Ruby](https://www.reddit.com/r/ruby/comments/fryvuh/main_usage_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fryvuh/main_usage_for_ruby/
---
I was thinking about implementing Ruby on a project that I'm working on, it involves a lot of console applications for Linux, however I started to think: What's the main usage of Ruby? And if it's ok to use it for desktop/terminal applications or scripts.

So I'm asking for your opinion on this matter.
What do you believe is the right usage for Ruby?

If your answer is other, please, feel free to comment your thoughts.

[View Poll](https://www.reddit.com/poll/fryvuh)
## [9][footballdb-clubs gem v2020.3.30 - the world's top football clubs (2500+ and counting) from around the world](https://www.reddit.com/r/ruby/comments/frvoeg/footballdbclubs_gem_v2020330_the_worlds_top/)
- url: https://github.com/sportdb/sport.db/tree/master/footballdb-clubs
---

## [10][If condition is off, returning the incorrect value](https://www.reddit.com/r/ruby/comments/frv3m1/if_condition_is_off_returning_the_incorrect_value/)
- url: https://www.reddit.com/r/ruby/comments/frv3m1/if_condition_is_off_returning_the_incorrect_value/
---
I am trying to figure out why this keeps returning the else statement. I assume it's because the array\[i\] is having a hard time comparing the element value. I want it to assume to take any element (string, integer, etc..)

`def includes_index(array, element)`  
  `limit = array.length`  
 `for i in 0..limit`  
 `if array[i] == element`  
 `return i`  
 `else`  
 `return -1`  
 `end`  
 `end`  
`end`  
`array1 = ["baseball", "basketball", "football", "hockey", "pokemon", "lacrosse", "soccer"]`  
`includes_index(array1, "lacrosse")`
