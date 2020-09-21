# ruby
## [1][Compile Ruby to C](https://www.reddit.com/r/ruby/comments/iwr3pn/compile_ruby_to_c/)
- url: https://github.com/agrafix/rubyspeed
---

## [2][Secure way of doing OAuth for SPA &amp; Native Apps](https://www.reddit.com/r/ruby/comments/iwztni/secure_way_of_doing_oauth_for_spa_native_apps/)
- url: https://blog.joshsoftware.com/2020/05/18/secure-way-of-doing-oauth-for-spa-native-apps/
---

## [3][What's a snippet of code that makes you happiest that you learnt Ruby?](https://www.reddit.com/r/ruby/comments/iwy3c4/whats_a_snippet_of_code_that_makes_you_happiest/)
- url: https://www.reddit.com/r/ruby/comments/iwy3c4/whats_a_snippet_of_code_that_makes_you_happiest/
---
For me currently it's:

    puts`dd`

8 character identity algorithm (outputs any input without any alteration). :) Keen to see other examples.
## [4][Did you know that block_given? is slower than defined?(yield) ?!](https://www.reddit.com/r/ruby/comments/iwxaai/did_you_know_that_block_given_is_slower_than/)
- url: https://www.reddit.com/gallery/iwxaai
---

## [5][An SMS Reminder Service With Ruby &amp; A Raspberry Pi](https://www.reddit.com/r/ruby/comments/iwoh4f/an_sms_reminder_service_with_ruby_a_raspberry_pi/)
- url: https://emmanuelhayford.com/building-an-sms-reminder-with-ruby-raspberry-pi/
---

## [6][An RSpec custom matcher to test application code that logs information into log files](https://www.reddit.com/r/ruby/comments/iwl1fq/an_rspec_custom_matcher_to_test_application_code/)
- url: https://github.com/juanmanuelramallo/rspec-log_matcher
---

## [7][NoMethodError - Calling private method?](https://www.reddit.com/r/ruby/comments/iwqkmq/nomethoderror_calling_private_method/)
- url: https://www.reddit.com/r/ruby/comments/iwqkmq/nomethoderror_calling_private_method/
---
 I'm following along with a tutorial about writing an shell in Ruby. The code below breaks when ever I type in a command at the prompt. The code attempts to call the 'execute' method then errors saying that this method is private. I don't understand why it's not working. Maybe I'm missing something really simple.   I'm sure I had this problem once before, I just don't remember how I solved it. 

&amp;#x200B;

$ ruby shell.rb

&gt;ls -a Traceback (most recent call last): 3: from shell.rb:59:in &lt;main&gt;' 2: from shell.rb:6:in   
main' 1: from shell.rb:6:in loop' shell.rb:10:in   
block in main': private method \`execute' called for #&lt;Command:0x00005637cb75b1e8 u/args=\["ls"@0, "-a"@3\]&gt; (NoMethodError)  
 

Ruby Code:

    require 'readline'
    require 'parslet'
    
    #PEG (Parsing Expressing Grammer)
    def main
        loop do
            cmdline = Readline.readline("&gt; ", true) # true = remember command history
           tree = parse_cmdline(cmdline)
            # p tree
            tree.execute
        end
    end
    
    def parse_cmdline(cmdline)
        raw_tree = Parser.new.parse(cmdline)
        # p raw_tree
        Transform.new.apply(raw_tree)
    end
    
    class Parser &lt; Parslet::Parser
        root :cmdline #Root being the object that will be parsed. 
    
        rule(:cmdline) { command }
        rule(:command) { arg.as(:arg).repeat(1).as(:command) }
        rule(:arg) { match[%q{^\s}].repeat(1) &gt;&gt; space? }
    
        rule(:space?) {space.maybe } #maybe - provided by parslet
        rule(:space) { match[%q{\s}].repeat(1).ignore }
    
    end
      
    class Transform &lt; Parslet::Transform
    
        rule(command: sequence(:args)) { Command.new(args) } #ls
        rule(arg: simple(:arg)) { arg } # -a
        
    end
    
    
    class Command
        def initialize(args)
            @args = args
        end
    
    end
    
    def execute
        spawn(*@args)
        
    
    end
    
    main
## [8][rodauth-select-account: manage multiple accounts within the same session.](https://www.reddit.com/r/ruby/comments/iwns99/rodauthselectaccount_manage_multiple_accounts/)
- url: https://gitlab.com/honeyryderchuck/rodauth-select-account
---

## [9][how does 1 + 2 work in ruby?](https://www.reddit.com/r/ruby/comments/iw8h06/how_does_1_2_work_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/iw8h06/how_does_1_2_work_in_ruby/
---
When I do `1.methods` I get a list of methods available like so  


    [:%, :&amp;, :*, :+, :-, :/, :&lt;, :&gt;, :^, :|, :~, :-@, :**, :&lt;=&gt;, :&lt;&lt;, :&gt;&gt;, :&lt;=, :&gt;=, :==, :===, :[], :inspect, :size, :succ, :to_int, :to_s, :to_i, :to_f, :next, :div, :upto, :chr, :ord, :coerce, :divmod, :fdiv, :modulo, :remainder, :abs, :magnitude, :integer?, :floor, :ceil, :round, :truncate, :odd?, :even?, :downto, :times, :pred, :bit_length, :digits, :to_r, :numerator, :denominator, :rationalize, :gcd, :lcm, :gcdlcm, :+@, :eql?, :singleton_method_added, :i, :real?, :zero?, :nonzero?, :finite?, :infinite?, :step, :positive?, :negative?, :quo, :arg, :rectangular, :rect, :polar, :real, :imaginary, :imag, :abs2, :angle, :phase, :conjugate, :conj, :to_c, :between?, :clamp, :remove_instance_variable, :instance_of?, :kind_of?, :is_a?, :tap, :public_send, :singleton_method, :instance_variable_defined?, :define_singleton_method, :method, :public_method, :instance_variable_set, :extend, :to_enum, :enum_for, :=~, :!~, :respond_to?, :freeze, :object_id, :send, :display, :nil?, :hash, :class, :singleton_class, :clone, :dup, :itself, :taint, :tainted?, :untaint, :untrust, :untrusted?, :trust, :frozen?, :methods, :singleton_methods, :protected_methods, :private_methods, :public_methods, :instance_variable_get, :instance_variables, :!, :!=, :__send__, :equal?, :instance_eval, :instance_exec, :__id__]

So all the operators are a method. 

so unlike `1.size` where we use a `.` operator how does `1 + 2` work without a dot operator while `1.+2` works
## [10][Ruby 3.0 Added Leading Arguments Forwarding](https://www.reddit.com/r/ruby/comments/ivoker/ruby_30_added_leading_arguments_forwarding/)
- url: https://www.reddit.com/r/ruby/comments/ivoker/ruby_30_added_leading_arguments_forwarding/
---
[https://scriptday.com/blog/2020/09/12/ruby-3-0-added-leading-arguments-forwarding](https://scriptday.com/blog/2020/09/12/ruby-3-0-added-leading-arguments-forwarding)
