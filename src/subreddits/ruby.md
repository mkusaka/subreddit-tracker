# ruby
## [1][Why Pry is one of the most important tools a junior Rubyist can learn](https://www.reddit.com/r/ruby/comments/fdh5gh/why_pry_is_one_of_the_most_important_tools_a/)
- url: https://www.reddit.com/r/ruby/comments/fdh5gh/why_pry_is_one_of_the_most_important_tools_a/
---
As programmers we often have to mentally run code. To imagine how a program will behave given certain inputs. This is hard enough for experienced developers. But for juniors? It can seem impossible. In this article, Melissa Williams argues that pry is an invaluable tool for junior rubyists because it allows them to see exactly what is going on as their code is run. [https://www.honeybadger.io/blog/debugging-ruby-with-pry/](https://www.honeybadger.io/blog/debugging-ruby-with-pry/)
## [2][Mistakes I've made treating file paths as strings](https://www.reddit.com/r/ruby/comments/fdnxz8/mistakes_ive_made_treating_file_paths_as_strings/)
- url: https://www.reddit.com/r/ruby/comments/fdnxz8/mistakes_ive_made_treating_file_paths_as_strings/
---
It bugs me about how long I've manipulated paths as strings without having a failure, but in the last couple of months things have broken in Ruby and Node projects. So I wrote about [the issues I've faced and how we should approach path names to avoid them](https://philna.sh/blog/2020/03/04/mistakes-treating-paths-as-strings/) (TL;DR use the standard lib `Pathname` and `File.join` in Ruby).
## [3][Can I make games or apps with ruby](https://www.reddit.com/r/ruby/comments/fdk0vs/can_i_make_games_or_apps_with_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fdk0vs/can_i_make_games_or_apps_with_ruby/
---

## [4][Is it possible for a method invoked on superclass to execute differently based on the subclass which called it?](https://www.reddit.com/r/ruby/comments/fdqdth/is_it_possible_for_a_method_invoked_on_superclass/)
- url: https://www.reddit.com/r/ruby/comments/fdqdth/is_it_possible_for_a_method_invoked_on_superclass/
---
Hi everyone,

Can anyone tell me if there's a way to do the following?

    class Animal
      def initialize
      end
    
      def sayhello
        # subclass is cat ? "meow" : "woof"
      end
    end
    
    class Cat &lt; Animal
    end
    
    class Dog &lt; Animal
    end
    
    p Cat.new.sayhello
    =&gt; "meow"
    
    p Dog.new.sayhello
    =&gt; "woof"

Thanks
## [5][Hanami::API on Amazon AWS Lambda](https://www.reddit.com/r/ruby/comments/fdctli/hanamiapi_on_amazon_aws_lambda/)
- url: https://lucaguidi.com/2020/03/04/hanamiapi-on-amazon-aws-lambda/
---

## [6][can somebody explain how does the following code execute?](https://www.reddit.com/r/ruby/comments/fdjnif/can_somebody_explain_how_does_the_following_code/)
- url: https://www.reddit.com/r/ruby/comments/fdjnif/can_somebody_explain_how_does_the_following_code/
---
I am following a linked tutorial from the odin project, its about  blocks and procs in ruby. i cant quite understand how does the following  code works.

    class Array def eachEven(&amp;wasABlock_nowAProc) # We start with "true" because arrays start with 0, which is even.     isEven = true self.each do |object| if isEven         wasABlock_nowAProc.call object       end        isEven = (not isEven) # Toggle from even to odd, or odd to even. end end end ['apple', 'bad apple', 'cherry', 'durian'].eachEven do |fruit|   puts 'Yum!  I just love '+fruit+' pies, don\'t you?' end # Remember, we are getting the even-numbered elements # of the array, all of which happen to be odd numbers, # just because I like to cause problems like that. [1, 2, 3, 4, 5].eachEven do |oddBall|   puts oddBall.to_s+' is NOT an even number!' end

is \`\['apple', 'bad apple', 'cherry', 'durian'\]\`a block in this context and are we calling the method \`isEven\`on that block? Does \`isEven\` used to only return true or false and if true \`do |fruit|   puts 'Yum!  I just love '+fruit+' pies, don\\'t you?' end\`is executed? also what is this line doing \`self.each do |object|       if isEven         wasABlock\_nowAProc.call object       end\`? if isEven is true then call \`\[1, 2, 3, 4, 5\]\`with the object??? what does calling that block with object mean?

&amp;#x200B;
## [7][Building a Rails App With Multiple Subdomains](https://www.reddit.com/r/ruby/comments/fddbrx/building_a_rails_app_with_multiple_subdomains/)
- url: https://blog.appsignal.com/2020/03/04/building-a-rails-app-with-multiple-subdomains.html
---

## [8][JRuby on Windows Day 0 - Install and Hello World](https://www.reddit.com/r/ruby/comments/fddrbf/jruby_on_windows_day_0_install_and_hello_world/)
- url: http://notepad.onghu.com/2020/jruby-win-day0-install-hello_world/
---

## [9][What tool would you recomend to document a Ruby Web Api?](https://www.reddit.com/r/ruby/comments/fdhnty/what_tool_would_you_recomend_to_document_a_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fdhnty/what_tool_would_you_recomend_to_document_a_ruby/
---
I have a Rubi Web api but it doesnt have documentation or any form for a person to know how it works (only the developers knows how it works), i need to document it to allow people online to use it in the right way.
## [10][`insert_all` and `upsert_all` ActiveRecord methods (a short recap)](https://www.reddit.com/r/ruby/comments/fd9iim/insert_all_and_upsert_all_activerecord_methods_a/)
- url: https://frontdeveloper.pl/2020/03/insert_all-and-upsert_all-activerecord-methods/
---

