# ruby
## [1][Is there a name for this code smell?](https://www.reddit.com/r/ruby/comments/iqntuy/is_there_a_name_for_this_code_smell/)
- url: https://www.reddit.com/r/ruby/comments/iqntuy/is_there_a_name_for_this_code_smell/
---
I have come across some code which calls a local private method with a parameter which is also a class variable.

      def initialize
        @record = find_record
        do_something(@record)
      end
    
      private
    
      def find_record
        'record yo'
      end
    
      def do_something(record)
      end

I feel like Sandi Metz would have something to say about that, or perhaps Rubocop would pick it up. How can I succinctly describe what is wrong here?
## [2][An SMS Reminder With Ruby &amp; A Raspberry Pi](https://www.reddit.com/r/ruby/comments/iqqk9e/an_sms_reminder_with_ruby_a_raspberry_pi/)
- url: https://emmanuelhayford.com/building-an-sms-reminder-with-ruby-raspberry-pi/
---

## [3][Enforcing Public/Private Access in Rails](https://www.reddit.com/r/ruby/comments/iqp2yk/enforcing_publicprivate_access_in_rails/)
- url: https://medium.com/opendoor-labs/enforcing-public-private-access-in-rails-57e9c21d141d
---

## [4][What happened to Jesse Storimer?](https://www.reddit.com/r/ruby/comments/iq1etj/what_happened_to_jesse_storimer/)
- url: https://www.reddit.com/r/ruby/comments/iq1etj/what_happened_to_jesse_storimer/
---
Does anyone know what happened to Jesse Storimer, the author of the brilliant books “Working With Unix Processes“, “Working With TCP Sockets“ and “Working With Ruby Threads“? His GitHub profile shows that the last activity was in 2015 and I couldn’t find any information on what he’s been doing since. It’s sad because he is such a great teacher and I feel like there’s so much more he could cover to the benefit of the Ruby community.
## [5][Ruby adds `Symbol#name` to return frozen string corresponding to the symbol name](https://www.reddit.com/r/ruby/comments/ipk3fu/ruby_adds_symbolname_to_return_frozen_string/)
- url: https://blog.saeloun.com/2020/09/09/ruby-adds-name-method-to-symbol
---

## [6][These Rails apps are overpacking their JavaScript bundles](https://www.reddit.com/r/ruby/comments/iplq0z/these_rails_apps_are_overpacking_their_javascript/)
- url: https://rossta.net/blog/rails-apps-overpacking-with-webpacker.html
---

## [7][Book suggestion for Rspec](https://www.reddit.com/r/ruby/comments/ipitcj/book_suggestion_for_rspec/)
- url: https://www.reddit.com/r/ruby/comments/ipitcj/book_suggestion_for_rspec/
---
I am looking for a book to master Rspec.

I already know the basics of RSpec testing but I am looking to learn good practices and advanced topics in Rspec.

I have not been able to find something online apart from basic tutorials. And was hoping for suggestions from you guys.
## [8][Choosing solution for Chat Bot](https://www.reddit.com/r/ruby/comments/ipm3xt/choosing_solution_for_chat_bot/)
- url: https://www.reddit.com/r/ruby/comments/ipm3xt/choosing_solution_for_chat_bot/
---
Hi, In the near future I have the plan to implement a Slack bot.   
I did some research and because of my favorite programming language is Ruby I think about using [https://www.lita.io](https://www.lita.io).  
Are any of you used this or can recommend better solution for building a Slack bot?  


Thanks in advance
## [9][Help with benchmark-ips](https://www.reddit.com/r/ruby/comments/ipgffm/help_with_benchmarkips/)
- url: https://www.reddit.com/r/ruby/comments/ipgffm/help_with_benchmarkips/
---
I am trying to benchmark a few ways I can check for equality, like so

    Benchmark.ips do |x|
      x.config(:time =&gt; 40, :warmup =&gt; 5)
    
      model = User.new
      x.report("equals") do
        model.class.name == 'User'
      end
    
      x.report("equals freeze") do
        model.class.name == 'User'.freeze
      end
    
      x.report("array") do
        ['User'].include?(model.class.name)
      end
    
      set = Set.new(['User'])
      x.report("set") do
        set.include?(model.class.name)
      end
    
      frozen_set = Set.new(['User']).freeze
      x.report("frozen set") do
        frozen_set.include?(model.class.name)
      end
    
      # Compare the iterations per second of the various reports!
      x.compare!
    end

And this is the result it gave  
 

    Warming up --------------------------------------
    equals  423.483k i/100ms
    equals freeze  573.846k i/100ms
    array  315.613k i/100ms
    set  459.956k i/100ms
    frozen set  461.620k i/100ms
    Calculating -------------------------------------
    equals 4.381M (± 5.7%) i/s - 174.898M in 40.074310s
    equals freeze 6.038M (± 3.5%) i/s - 241.589M in 40.059452s
    array 3.261M (± 5.0%) i/s - 130.348M in 40.074906s
    set 4.698M (± 3.3%) i/s - 188.122M in 40.087738s
    frozen set 4.741M (± 4.0%) i/s - 189.726M in 40.090792s
    Comparison:
    equals freeze: 6038148.2 i/s
    frozen set: 4741206.2 i/s - 1.27x (± 0.00) slower
    set: 4698138.7 i/s - 1.29x (± 0.00) slower
    equals: 4381153.2 i/s - 1.38x (± 0.00) slower
    array: 3261099.2 i/s - 1.85x (± 0.00) slower

I have a few questions regd. this

1. Is there any documentation on how I can understand the output from this better. I can see that more I/O is better but more information on these numbers would be great.
2. How do I measure the memory consumption of each of the approaches, i.e. the numbers of objects each approach creates.

Thanks.
## [10][Gem API design question](https://www.reddit.com/r/ruby/comments/ipce8i/gem_api_design_question/)
- url: https://www.reddit.com/r/ruby/comments/ipce8i/gem_api_design_question/
---
Hi,

I'm new to the Ruby world, and a project at work is using a gem implementing a wrapper to a remote service API. What troubles me is that this gem is exposing only class-level methods, including the "configure" method which does *not* return an object instance.

Its usage is something like this :

```ruby  
ZooApiWrapper.configure { |config|
  config.token = "abcdef1234"
  config.url = "https://my.api/"
}

lions = ZooApiWrapper.animals.lions

new_pet = ZooApiWrapper::Animal::Dog.new
new_pet.name = "Spot"
new_pet.save

# etc...
```

I would have expected to have the configure call return a wrapper object, and use this wrapper object to access the various resources exposed via the API. Is it usual to use global state for such things instead in Ruby ?
