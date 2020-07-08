# ruby
## [1][Ruby income from side projects. Do you?](https://www.reddit.com/r/ruby/comments/hnav7i/ruby_income_from_side_projects_do_you/)
- url: https://www.reddit.com/r/ruby/comments/hnav7i/ruby_income_from_side_projects_do_you/
---
I've seen this question on a few other subreddits, so I thought i'd ask here. Apart from your main job or full time work, do you make any side income from Ruby? Even if it's only a small amount, I would love to hear about it. Come tell all!!
## [2][Performance -- What about making a GCC front-end for Ruby?](https://www.reddit.com/r/ruby/comments/hnfz58/performance_what_about_making_a_gcc_frontend_for/)
- url: https://www.reddit.com/r/ruby/comments/hnfz58/performance_what_about_making_a_gcc_frontend_for/
---
I am working on my own programming language, with a native compiler.  


Since LLVM is kinda slow, I thought that maybe making a GCC front-end will be faster (and indeed it is).  


Yesterday, I saw an article about Ruby3\*3 and all that stuff, including optimizations of MJIT (the current JIT compiler for Ruby) and for a second I had a crazy idea: making Ruby as fast as C/Crystal by compiling it to native code, with a simple garbage collector and runtime.  


Thanks to the gem \`parser\` a Ruby to abstract syntax tree compiler isn't even needed, I can call the MRI from C via \`fork()\` and then \`execv()\` (will probably take up to 10 seconds more than a parser written in pure C).  


I already know about the pain of porting the core library and standard library and a package manager like RubyGems, but probably the most painful thing will be adding the possibility for creating C extensions, of course.  


What do you guys about this idea?
## [3][Redis in Ruby - A work in progress online book about reimplementing redis ... in ruby](https://www.reddit.com/r/ruby/comments/hmxcgv/redis_in_ruby_a_work_in_progress_online_book/)
- url: https://www.reddit.com/r/ruby/comments/hmxcgv/redis_in_ruby_a_work_in_progress_online_book/
---
I'm really excited to share a project I've been working on for a few months now, as the title says, I'm writing an online book about reimplementing Redis, but in Ruby:

[https://redis.pjam.me/](https://redis.pjam.me/)

Why you may ask? Because I think it's a lot of fun and there's a lot to learn on the way, tcp servers, tcp sockets, system calls (select, read, ...), threads, processes, we'll touch on all of these.

There are three chapters currently available, "Creating a Basic TCP server", "Implementing the first commands, GET &amp; SET", and "Handling multiple clients".

I laid out what I think the next chapters will be: [https://redis.pjam.me/chapters](https://redis.pjam.me/chapters)

If you have feedback, please let me know, and I hope you'll enjoy it as much as I enjoy writing it.

One of the inspirations for this is this ebook that I love and absolutely recommend: [https://shop.jcoglan.com/building-git/](https://shop.jcoglan.com/building-git/)
## [4][A Fast Car Needs Good Brakes: How We Added Client Rate Throttling to the Platform API Gem](https://www.reddit.com/r/ruby/comments/hngu5n/a_fast_car_needs_good_brakes_how_we_added_client/)
- url: https://blog.heroku.com/rate-throttle-api-client
---

## [5][Anonymous Struct Literals Might Be Coming To Ruby](https://www.reddit.com/r/ruby/comments/hmvn1k/anonymous_struct_literals_might_be_coming_to_ruby/)
- url: https://supergood.software/ruby-anonymous-struct-literals/
---

## [6][Hard Code Price - Shopify Script](https://www.reddit.com/r/ruby/comments/hmu1mh/hard_code_price_shopify_script/)
- url: https://www.reddit.com/r/ruby/comments/hmu1mh/hard_code_price_shopify_script/
---
So I used the Shopify Script Generator to generate a script that discounts Shipping Rates by a number or percentage. However I want to adapt this code so I can just hardcode the price rather than discount it, so for example I can explicitly set it to £5 for arguments sake, I'm more of a Python dev so I'm having alot of trouble adapting! Code below  


    class Campaign
      def initialize(condition, *qualifiers)
        @condition = (condition.to_s + '?').to_sym
        @qualifiers = PostCartAmountQualifier ? [] : [] rescue qualifiers.compact
        @line_item_selector = qualifiers.last unless @line_item_selector
        qualifiers.compact.each do |qualifier|
          is_multi_select = qualifier.instance_variable_get(:@conditions).is_a?(Array)
          if is_multi_select
            qualifier.instance_variable_get(:@conditions).each do |nested_q|
              @post_amount_qualifier = nested_q if nested_q.is_a?(PostCartAmountQualifier)
              @qualifiers &lt;&lt; qualifier
            end
          else
            @post_amount_qualifier = qualifier if qualifier.is_a?(PostCartAmountQualifier)
            @qualifiers &lt;&lt; qualifier
          end
        end if @qualifiers.empty?
      end
      def qualifies?(cart)
        return true if @qualifiers.empty?
        @unmodified_line_items = cart.line_items.map do |item|
          new_item = item.dup
          new_item.instance_variables.each do |var|
            val = item.instance_variable_get(var)
            new_item.instance_variable_set(var, val.dup) if val.respond_to?(:dup)
          end
          new_item
        end if @post_amount_qualifier
        @qualifiers.send(@condition) do |qualifier|
          is_selector = false
          if qualifier.is_a?(Selector) || qualifier.instance_variable_get(:@conditions).any? { |q| q.is_a?(Selector) }
            is_selector = true
          end rescue nil
          if is_selector
            raise "Missing line item match type" if @li_match_type.nil?
            cart.line_items.send(@li_match_type) { |item| qualifier.match?(item) }
          else
            qualifier.match?(cart, @line_item_selector)
          end
        end
      end
      def run_with_hooks(cart)
        before_run(cart) if respond_to?(:before_run)
        run(cart)
        after_run(cart)
      end
      def after_run(cart)
        @discount.apply_final_discount if @discount &amp;&amp; @discount.respond_to?(:apply_final_discount)
        revert_changes(cart) unless @post_amount_qualifier.nil? || @post_amount_qualifier.match?(cart)
      end
      def revert_changes(cart)
        cart.instance_variable_set(:@line_items, @unmodified_line_items)
      end
    end
    class ShippingDiscount &lt; Campaign
      def initialize(condition, customer_qualifier, cart_qualifier, li_match_type, line_item_qualifier, rate_selector, discount)
        super(condition, customer_qualifier, cart_qualifier, line_item_qualifier)
        @li_match_type = (li_match_type.to_s + '?').to_sym
        @rate_selector = rate_selector
        @discount = discount
      end
      def run(rates, cart)
        raise "Campaign requires a discount" unless @discount
        return unless qualifies?(cart)
        rates.each do |rate|
          next unless @rate_selector.nil? || @rate_selector.match?(rate)
          @discount.apply(rate)
        end
      end
    end
    class Selector
      def partial_match(match_type, item_info, possible_matches)
        match_type = (match_type.to_s + '?').to_sym
        if item_info.kind_of?(Array)
          possible_matches.any? do |possibility|
            item_info.any? do |search|
              search.send(match_type, possibility)
            end
          end
        else
          possible_matches.any? do |possibility|
            item_info.send(match_type, possibility)
          end
        end
      end
    end
    class RateNameSelector &lt; Selector
      def initialize(match_type, match_condition, names)
        @match_condition = match_condition
        @invert = match_type == :does_not
        @names = names.map(&amp;:downcase)
      end
      def match?(shipping_rate)
        name = shipping_rate.name.downcase
        case @match_condition
          when :match
            return @invert ^ @names.include?(name)
          else
            return @invert ^ partial_match(@match_condition, name, @names)
        end
      end
    end
    class FixedDiscount
      def initialize(amount, message)
        @amount = Money.new(cents: amount * 100)
        @message = message
      end
      def apply(rate)
        discount_amount = rate.price - @amount &lt; Money.zero ? rate.price : @amount
        rate.apply_discount(discount_amount, { message: @message })
      end
    end
    CAMPAIGNS = [
      ShippingDiscount.new(
        :any,
        nil,
        nil,
        :any,
        nil,
        RateNameSelector.new(
          :does,
          :match,
          ["Saturday Delivery (order by 11pm Fri)"]
        ),
        FixedDiscount.new(
          3,
          ""
        )
      ),
      ShippingDiscount.new(
        :any,
        nil,
        nil,
        :any,
        nil,
        RateNameSelector.new(
          :does,
          :match,
          ["Saturday Delivery (order by 11pm Fri)"]
        ),
        FixedDiscount.new(
          2,
          ""
        )
      )
    ].freeze
    CAMPAIGNS.each do |campaign|
      campaign.run(Input.shipping_rates, Input.cart)
    end
    
    Output.shipping_rates = Input.shipping_rates
## [7][Why validation matchers are the only Shoulda matchers I use](https://www.reddit.com/r/ruby/comments/hmmdza/why_validation_matchers_are_the_only_shoulda/)
- url: https://www.codewithjason.com/validation-matchers-shoulda-matchers-use/
---

## [8][Best deep-dive Ruby course?](https://www.reddit.com/r/ruby/comments/hmg444/best_deepdive_ruby_course/)
- url: https://www.reddit.com/r/ruby/comments/hmg444/best_deepdive_ruby_course/
---
Hey,

I’m looking for suggestions on a really good advanced Ruby course. It can start with the basics, but it should also ramp up and get into the weeds.

To better “explain” what I’m looking for, I once followed a course for JavaScript that went into the source code for V8 several times to explain concepts. It explained It also outlined how some of the major JS frameworks (at the time) where built.

At the end of this, I want to understand all the major mechanics surrounding the ruby language.

Thank you for your suggestions!

——

Edit:

I’m currently reading Why’s guide to Ruby. And I’m learning a bunch. I’ve also considered just reading a bunch of docs etc. But sometimes something in video format is easier to digest. Or a more carefully written (series of) articles.
## [9][All About Ruby](https://www.reddit.com/r/ruby/comments/hmgjql/all_about_ruby/)
- url: https://medium.com/@jinglis12/all-about-ruby-fc83a0a18bd3
---

## [10][Continuous Integration Testing: Basics + What to Test](https://www.reddit.com/r/ruby/comments/hmb5qu/continuous_integration_testing_basics_what_to_test/)
- url: https://www.youtube.com/watch?v=f-1Q896R1no
---

