# ruby
## [1][Anti-IF framework - if/else based on type](https://www.reddit.com/r/ruby/comments/ilt25m/antiif_framework_ifelse_based_on_type/)
- url: https://blog.arkency.com/anti-if-framework---if-slash-else-based-on-type/
---

## [2][[Ruby, Rails, Sorbet] Cleo and Robb clean up a Ruby on Rails codebase with Sorbet](https://www.reddit.com/r/ruby/comments/ilk9f4/ruby_rails_sorbet_cleo_and_robb_clean_up_a_ruby/)
- url: https://www.reddit.com/r/ruby/comments/ilk9f4/ruby_rails_sorbet_cleo_and_robb_clean_up_a_ruby/
---
[https://www.twitch.tv/lawiscode](https://www.twitch.tv/lawiscode)
## [3][Using methods with super when there is no parent class other than active record](https://www.reddit.com/r/ruby/comments/ilocdh/using_methods_with_super_when_there_is_no_parent/)
- url: https://www.reddit.com/r/ruby/comments/ilocdh/using_methods_with_super_when_there_is_no_parent/
---
Hi. I came across this piece of code:

`class XYZ &lt; ActiveRecord::Base`

  `def method1`

`print("hello world")`

`super`

  `end`

My question is this. Since there is no defined parent class other than active record, what does this super do ? is it a bug in the system or does it actually do something. This code is a few years old.

Edit: the answer is that there was a module included in the model which had 'method1' defined.
## [4][when you paste the name of a variable and press enter in irb, which method is utilized by irb? .to_s, .inspect or .display](https://www.reddit.com/r/ruby/comments/ill8pw/when_you_paste_the_name_of_a_variable_and_press/)
- url: https://www.reddit.com/r/ruby/comments/ill8pw/when_you_paste_the_name_of_a_variable_and_press/
---
when you paste the name of a variable and press enter in irb, which one is called? to\_s, .inspect or .display
## [5][Debugging a bug when using rand](https://www.reddit.com/r/ruby/comments/ilp0sl/debugging_a_bug_when_using_rand/)
- url: https://www.reddit.com/r/ruby/comments/ilp0sl/debugging_a_bug_when_using_rand/
---
The ruby faker library used this below method to generate random characters in one of their much older releases

     def self.characters(char_count = 255)
       rand(36**char_count).to_s(36)
     end

However, they changed it to something like this in the next release

     def characters(char_count = 255)
       rand(36**char_count).to_s(36).rjust(char_count, '0').chars.to_a.shuffle.join
     end 

Now, when I use the first release sometimes get a lesser number of characters that specified. it's probably a bug that got fixed later on. But could here help me identify why the first code would fail sometimes randomly when called say a thousand times. Thanks.
## [6][What are the Code Coverage Metrics for Ruby on Rails?](https://www.reddit.com/r/ruby/comments/il7r5h/what_are_the_code_coverage_metrics_for_ruby_on/)
- url: https://www.fastruby.io/blog/rails/what-is-code-coverage-ruby-on-rails.html
---

## [7][Introduction to DSL by implementing Rake](https://www.reddit.com/r/ruby/comments/il40do/introduction_to_dsl_by_implementing_rake/)
- url: https://medium.com/rubycademy/introduction-to-dsl-by-implementing-rake-99439bdc3513
---

## [8][Rails 6 adds support to persist timezones of Active Job | The Official BigBinary Blog | BigBinary](https://www.reddit.com/r/ruby/comments/il2gy4/rails_6_adds_support_to_persist_timezones_of/)
- url: https://blog.bigbinary.com/2020/09/01/rails-6-add-timezone-support-in-active-job.html
---

## [9][Moving from EventMachine to Async](https://www.reddit.com/r/ruby/comments/ikjcux/moving_from_eventmachine_to_async/)
- url: https://blog.joshsoftware.com/2020/06/19/moving-from-eventmachine-to-async/
---

## [10][My experience re-creating Xonix32 using Ruby with Gosu, Ruby Packer, Platypus.](https://www.reddit.com/r/ruby/comments/ikhvga/my_experience_recreating_xonix32_using_ruby_with/)
- url: https://youtu.be/VqO5XDNauws
---

