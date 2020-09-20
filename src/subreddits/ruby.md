# ruby
## [1][how does 1 + 2 work in ruby?](https://www.reddit.com/r/ruby/comments/iw8h06/how_does_1_2_work_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/iw8h06/how_does_1_2_work_in_ruby/
---
When I do `1.methods` I get a list of methods available like so  


    [:%, :&amp;, :*, :+, :-, :/, :&lt;, :&gt;, :^, :|, :~, :-@, :**, :&lt;=&gt;, :&lt;&lt;, :&gt;&gt;, :&lt;=, :&gt;=, :==, :===, :[], :inspect, :size, :succ, :to_int, :to_s, :to_i, :to_f, :next, :div, :upto, :chr, :ord, :coerce, :divmod, :fdiv, :modulo, :remainder, :abs, :magnitude, :integer?, :floor, :ceil, :round, :truncate, :odd?, :even?, :downto, :times, :pred, :bit_length, :digits, :to_r, :numerator, :denominator, :rationalize, :gcd, :lcm, :gcdlcm, :+@, :eql?, :singleton_method_added, :i, :real?, :zero?, :nonzero?, :finite?, :infinite?, :step, :positive?, :negative?, :quo, :arg, :rectangular, :rect, :polar, :real, :imaginary, :imag, :abs2, :angle, :phase, :conjugate, :conj, :to_c, :between?, :clamp, :remove_instance_variable, :instance_of?, :kind_of?, :is_a?, :tap, :public_send, :singleton_method, :instance_variable_defined?, :define_singleton_method, :method, :public_method, :instance_variable_set, :extend, :to_enum, :enum_for, :=~, :!~, :respond_to?, :freeze, :object_id, :send, :display, :nil?, :hash, :class, :singleton_class, :clone, :dup, :itself, :taint, :tainted?, :untaint, :untrust, :untrusted?, :trust, :frozen?, :methods, :singleton_methods, :protected_methods, :private_methods, :public_methods, :instance_variable_get, :instance_variables, :!, :!=, :__send__, :equal?, :instance_eval, :instance_exec, :__id__]

So all the operators are a method. 

so unlike `1.size` where we use a `.` operator how does `1 + 2` work without a dot operator while `1.+2` works
## [2][Ruby 3.0 Added Leading Arguments Forwarding](https://www.reddit.com/r/ruby/comments/ivoker/ruby_30_added_leading_arguments_forwarding/)
- url: https://www.reddit.com/r/ruby/comments/ivoker/ruby_30_added_leading_arguments_forwarding/
---
[https://scriptday.com/blog/2020/09/12/ruby-3-0-added-leading-arguments-forwarding](https://scriptday.com/blog/2020/09/12/ruby-3-0-added-leading-arguments-forwarding)
## [3][Automate Rails server provisioning and deployment using Ansible and Capistrano](https://www.reddit.com/r/ruby/comments/ivpbfc/automate_rails_server_provisioning_and_deployment/)
- url: https://rubyyagi.com/rails-deploy-automate-ansible/
---

## [4][We Made Puma Faster With Sleep Sort](https://www.reddit.com/r/ruby/comments/ivasiq/we_made_puma_faster_with_sleep_sort/)
- url: https://www.speedshop.co/2020/09/17/we-made-puma-faster-with-sleep-sort.html
---

## [5][[EN] Running Rack and Rails Faster with TruffleRuby / Benoit Daloze @eregontp](https://www.reddit.com/r/ruby/comments/ivcyyk/en_running_rack_and_rails_faster_with_truffleruby/)
- url: https://www.youtube.com/watch?v=281YdMYRAsk
---

## [6][[EN] Don't Wait For Me! Scalable Concurrency for Ruby 3! / Samuel Williams @ioquatix](https://www.reddit.com/r/ruby/comments/iv1dk1/en_dont_wait_for_me_scalable_concurrency_for_ruby/)
- url: https://www.youtube.com/watch?v=Y29SSOS4UOc
---

## [7][Method to randomly select name from an array](https://www.reddit.com/r/ruby/comments/ivanf4/method_to_randomly_select_name_from_an_array/)
- url: https://www.reddit.com/r/ruby/comments/ivanf4/method_to_randomly_select_name_from_an_array/
---
Hi!

I am trying to solve a problem that uses a method to randomly select a name from array. Here is the code that I have generated thus far:

**# Please write your code between this line...**

**def random\_name(array)**

**array = \["Ana", "Ollie"\].sample**

**end**

**# ...and between this line**

**puts "My name is #{random\_name()}"**

With this code I get the error message: "wrong number of arguments (given 0, expected 1)"

I am new to Ruby and coding in general, so I am not sure where to go from here. Any help would be greatly appreciated!
## [8][Ruby Antipatterns - I don't agree with all of them but the list is worth reading](https://www.reddit.com/r/ruby/comments/iv1pnm/ruby_antipatterns_i_dont_agree_with_all_of_them/)
- url: https://www.alchemists.io/articles/ruby_antipatterns/
---

## [9][Document Template Solutions?](https://www.reddit.com/r/ruby/comments/iv6sn1/document_template_solutions/)
- url: https://www.reddit.com/r/ruby/comments/iv6sn1/document_template_solutions/
---
Hello! I work in a financial service company which generates a lot of documents. We implement documents as HTML and print them to PDF for consumption, which I think is a pretty standard approach.

Although, I'm finding this to be quite hard to maintain as we get a lot of minor changes on copy over time and errors that creep in because the owners of the documents are not able to edit the HTML.

I wonder, is this a common problem in other companies? Are there other more efficient solutions to manage document templates?

PS: Posting here because we use Ruby and Ruby specific solutions are welcome, but this is a generic problem that I'm sure is faced by people working with other technologies
## [10][How the way you define hash default value can generate side effects](https://www.reddit.com/r/ruby/comments/iv3llq/how_the_way_you_define_hash_default_value_can/)
- url: https://medium.com/rubycademy/hash-shift-using-default-value-b0d7cb62ffa2
---

