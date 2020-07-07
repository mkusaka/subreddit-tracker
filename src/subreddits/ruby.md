# ruby
## [1][Hard Code Price - Shopify Script](https://www.reddit.com/r/ruby/comments/hmu1mh/hard_code_price_shopify_script/)
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
## [2][Why validation matchers are the only Shoulda matchers I use](https://www.reddit.com/r/ruby/comments/hmmdza/why_validation_matchers_are_the_only_shoulda/)
- url: https://www.codewithjason.com/validation-matchers-shoulda-matchers-use/
---

## [3][Best deep-dive Ruby course?](https://www.reddit.com/r/ruby/comments/hmg444/best_deepdive_ruby_course/)
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
## [4][All About Ruby](https://www.reddit.com/r/ruby/comments/hmgjql/all_about_ruby/)
- url: https://medium.com/@jinglis12/all-about-ruby-fc83a0a18bd3
---

## [5][Continuous Integration Testing: Basics + What to Test](https://www.reddit.com/r/ruby/comments/hmb5qu/continuous_integration_testing_basics_what_to_test/)
- url: https://www.youtube.com/watch?v=f-1Q896R1no
---

## [6][Soft Delete with Discard](https://www.reddit.com/r/ruby/comments/hm6wju/soft_delete_with_discard/)
- url: https://www.driftingruby.com/episodes/soft-delete-with-discard?utm_medium=social&amp;utm_campaign=weekly_episode&amp;utm_source=reddit
---

## [7][Looking for an old Ruby blog post.](https://www.reddit.com/r/ruby/comments/hmb4ut/looking_for_an_old_ruby_blog_post/)
- url: https://www.reddit.com/r/ruby/comments/hmb4ut/looking_for_an_old_ruby_blog_post/
---
Sorry if this seems insignificant, but I've been looking for this old article for ages and I just can't seem to find it. It was a blog post about Ruby's culture/community on a blog that was otherwise mainly about sushi, folk music, and Discworld. I think it was British. The article was linked to quite a bit back in the day. I've been trying to Google it for ages but not luck.
## [8][What is the use of Ruby and how its development is proceeding?](https://www.reddit.com/r/ruby/comments/hmt1he/what_is_the_use_of_ruby_and_how_its_development/)
- url: https://www.reddit.com/r/ruby/comments/hmt1he/what_is_the_use_of_ruby_and_how_its_development/
---
Ruby is a nice language to type, but I'm unable to understand its use case.  
C-C++ is used for performance  
R-Python is used for data science and AI.  
What is the use of Ruby?  
Is Ruby an evolving language?
## [9][A 65 LOC Instagram scraper powered by Kimurai](https://www.reddit.com/r/ruby/comments/hm5mh5/a_65_loc_instagram_scraper_powered_by_kimurai/)
- url: https://github.com/glaucocustodio/inspider
---

## [10][rubocop-graphql—a rubocop plugin that enforces graphql-ruby best practices](https://www.reddit.com/r/ruby/comments/hm2mrz/rubocopgraphqla_rubocop_plugin_that_enforces/)
- url: https://www.reddit.com/r/ruby/comments/hm2mrz/rubocopgraphqla_rubocop_plugin_that_enforces/
---
Source code is [here](https://github.com/DmitryTsepelev/rubocop-graphql).

For now, the gem can do following things:

- makes sure fields and arguments have snake–cased names;
- suggests to add descriptions to entities;
- forces grouping field definitions together/with their resolvers;
- suggests using input types in complex mutations;
- proposes replacing long resolver methods with resolver objects.
