# ruby
## [1][Testing Rails generators with ease.](https://www.reddit.com/r/ruby/comments/ir8tor/testing_rails_generators_with_ease/)
- url: https://www.abhaynikam.me/posts/rails_generator_testing/
---

## [2][.each method not working? RUBY](https://www.reddit.com/r/ruby/comments/ir8z7j/each_method_not_working_ruby/)
- url: https://www.reddit.com/r/ruby/comments/ir8z7j/each_method_not_working_ruby/
---
Hi, can someone please tell me why my each is not working to get all the character upcase? 

def asf(str)  
str.split("").each {|char| char.upcase}  
end   
puts asf("what a string")
## [3][Ruby Jard - Just Another Ruby Debugger](https://www.reddit.com/r/ruby/comments/ircolf/ruby_jard_just_another_ruby_debugger/)
- url: https://www.reddit.com/r/ruby/comments/ircolf/ruby_jard_just_another_ruby_debugger/
---
Hi all, I would like to show you guys a Ruby debugger, named Ruby  Jard, that I'm working on recently. Ruby Jard provides a rich Terminal  UI that visualizes everything your need, navigates your program with  pleasure, stops at matter places only, reduces manual and mental  efforts. You can now focus on real debugging.

Ruby Jard wraps around Byebug, and Pry and provides a series of features and enhancements. Some significant features:

* Rich intuitive Terminal UI, runs directly in terminal emulators.
* Navigate with pleasure. No more next, next, next, next, and have no idea what you are doing.
* Debug filter. Prevent your program from stepping into weird places.  You can specify which layer (application only, gems, standard libs,  etc.) you are interested in, and can combine with the included/excluded  list.
* Visualize everything you need, grab important things at one glance. A  totally new variable inspection that helps you comprehend any variable  data shape, and peek at its data.
* Powerful REPL console
* Flexible customization

It is still in Beta version. So, it would be great to hear the feedback so that I can improve it in the future.

You can see it in action at [https://rubyjard.org/](https://rubyjard.org/). Or watch the following video: [https://asciinema.org/a/358874](https://asciinema.org/a/358874).

Some screenshots:

https://preview.redd.it/q6yjcltmypm51.png?width=1240&amp;format=png&amp;auto=webp&amp;s=ad3853be5c6de571e6a8ee446cdde06a378b9d41

https://preview.redd.it/7c6i1yrpypm51.png?width=1272&amp;format=png&amp;auto=webp&amp;s=8dcbc1c1a04b48bd60c11725d07a08fa121e9a71

https://preview.redd.it/h0rzk0spypm51.png?width=1252&amp;format=png&amp;auto=webp&amp;s=6ceeacf144199ba5fd4f1c90fc88eb94c64d656d

https://preview.redd.it/yn0aa6spypm51.png?width=1248&amp;format=png&amp;auto=webp&amp;s=e3d940e7b3bad638537d2121959a2b7bdec02b6d

https://preview.redd.it/m5oh3azrypm51.png?width=1251&amp;format=png&amp;auto=webp&amp;s=cf94a3cecfb756d26824bf324f2a9536be654742

[Debug Rails on 2K screen](https://preview.redd.it/f0mt8bzsypm51.png?width=2559&amp;format=png&amp;auto=webp&amp;s=ca3522c663fa6f1b102ae5498d8db853f1137760)
## [4][Caeser cipher exercise Ruby](https://www.reddit.com/r/ruby/comments/ir9lk4/caeser_cipher_exercise_ruby/)
- url: https://www.reddit.com/r/ruby/comments/ir9lk4/caeser_cipher_exercise_ruby/
---
Can someone help me review this code? still having trouble solving this problem: It is a type of substitution cipher in which each letter in the plaintext is replaced by a letter some fixed number of positions down the alphabet. For example, with a left shift of 3, D would be replaced by A, E would become B, and so on. 

It needs to return "Bmfy f xywnsl!"

THANKS!

def caeser\_cipher(str, key)  
 lowercase = ("a".."z").to\_a  
 uppercase = ("A".."Z").to\_a  
 max = lowercase.length   
 output = ""  
   
   
str.split('').each {|char|   
case  
 when lowercase.include?(char) || uppercase.include?(char) then charNum = lowercase.find\_index(char.downcase)   
 when char == char.upcase &amp;&amp; charNum + n &lt;= max then newChar = uppercase\[charNum+n\] #index  
 when char == char.upcase &amp;&amp; charNum + n &gt; max then remainder = (charNum + n) - max   
 newChar = uppercase\[remainder\]   
 when char == char.downcase &amp;&amp; charNum + n &lt;= max then newChar = lowecase\[charNum+n\]  
 when char == char.downcase &amp;&amp; charNum + n &gt; max then   
remainder = (charNum + n) - max  
newChar = uppercase\[remainder\]   
 else output &lt;&lt; char  
 end  
 output+=newChar  
}  
 return output  
end  
caeser\_cipher("What a string!", 5)
## [5][Took a performance pass over DragonRuby Game Toolkit. Pretty decent results!](https://www.reddit.com/r/ruby/comments/iqrddq/took_a_performance_pass_over_dragonruby_game/)
- url: https://www.youtube.com/watch?v=UuY7CWdvyWM&amp;feature=youtu.be&amp;ab_channel=AmirRajan
---

## [6][Help VSCode RSpec](https://www.reddit.com/r/ruby/comments/ir1yar/help_vscode_rspec/)
- url: https://www.reddit.com/r/ruby/comments/ir1yar/help_vscode_rspec/
---
Hey guys I'm trying to do rspec on vscode but it isnt working. I did bundle exec rspec in the terminal but its just showing whats on the bottom. There should be clear rspec like in the second picture. Thanks!

&amp;#x200B;

https://preview.redd.it/g5wv6smitlm51.png?width=1920&amp;format=png&amp;auto=webp&amp;s=b3a87f02368926440f9b94b47492008fd42cb4f1

https://preview.redd.it/r5vormedulm51.png?width=850&amp;format=png&amp;auto=webp&amp;s=2efcd99bc4c4260d7fc3d39129175b48cd385bbf
## [7][An SMS Reminder With Ruby &amp; A Raspberry Pi](https://www.reddit.com/r/ruby/comments/iqqk9e/an_sms_reminder_with_ruby_a_raspberry_pi/)
- url: https://emmanuelhayford.com/building-an-sms-reminder-with-ruby-raspberry-pi/
---

## [8][Is there a name for this code smell?](https://www.reddit.com/r/ruby/comments/iqntuy/is_there_a_name_for_this_code_smell/)
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
## [9][Can someone please explain this line by line using num = 3?](https://www.reddit.com/r/ruby/comments/iqqnuq/can_someone_please_explain_this_line_by_line/)
- url: https://www.reddit.com/r/ruby/comments/iqqnuq/can_someone_please_explain_this_line_by_line/
---
I just started to learn Ruby. I would really appreciate any help! 

def fibs(num)  
 return arr\[0...num\] if num &lt;= 2  
  ary = \[0, 1\]  
 (3..num).each { |i| ary &lt;&lt; (ary\[i - 3\] + ary\[i - 2\])}  
  ary  
end  
p fibs(3)
## [10][Enforcing Public/Private Access in Rails](https://www.reddit.com/r/ruby/comments/iqp2yk/enforcing_publicprivate_access_in_rails/)
- url: https://medium.com/opendoor-labs/enforcing-public-private-access-in-rails-57e9c21d141d
---

