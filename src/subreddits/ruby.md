# ruby
## [1][Conditional class opening in Ruby](https://www.reddit.com/r/ruby/comments/irtuw8/conditional_class_opening_in_ruby/)
- url: https://medium.com/rubycademy/conditional-class-opening-in-ruby-49a4524b7f85
---

## [2][Should I stick with Ruby or learn a more modern language like C#, Python, or PHP?](https://www.reddit.com/r/ruby/comments/irxxwb/should_i_stick_with_ruby_or_learn_a_more_modern/)
- url: /r/learnprogramming/comments/irlp9c/stick_with_learning_ruby_or_switch_over_to/
---

## [3][Did you know that Time#utc modifies the receiver?](https://www.reddit.com/r/ruby/comments/irkhj2/did_you_know_that_timeutc_modifies_the_receiver/)
- url: https://i.redd.it/vv5ac7il4sm51.png
---

## [4][Ruby Jard - Just Another Ruby Debugger](https://www.reddit.com/r/ruby/comments/ircolf/ruby_jard_just_another_ruby_debugger/)
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
## [5][The Most Popular Programming Languages - 1965/2020](https://www.reddit.com/r/ruby/comments/irh1dr/the_most_popular_programming_languages_19652020/)
- url: https://youtu.be/UNSoPa-XQN0
---

## [6][How would I implement Regexp#optional? ?](https://www.reddit.com/r/ruby/comments/irmbsq/how_would_i_implement_regexpoptional/)
- url: https://www.reddit.com/r/ruby/comments/irmbsq/how_would_i_implement_regexpoptional/
---
Given some random regular expression passed into a method, how do I determine if it's optional? ie. how do I see if it would match the empty string without trying to match the empty string?

This works, but seems like a bad idea:

    class Regexp
        def optional?
            not self.match('').nil?
        end
    end

Is there a better way?
## [7][Testing Rails generators with ease.](https://www.reddit.com/r/ruby/comments/ir8tor/testing_rails_generators_with_ease/)
- url: https://www.abhaynikam.me/posts/rails_generator_testing/
---

## [8][Caeser cipher exercise Ruby](https://www.reddit.com/r/ruby/comments/ir9lk4/caeser_cipher_exercise_ruby/)
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
## [9][.each method not working? RUBY](https://www.reddit.com/r/ruby/comments/ir8z7j/each_method_not_working_ruby/)
- url: https://www.reddit.com/r/ruby/comments/ir8z7j/each_method_not_working_ruby/
---
Hi, can someone please tell me why my each is not working to get all the character upcase? 

def asf(str)  
str.split("").each {|char| char.upcase}  
end   
puts asf("what a string")
## [10][Took a performance pass over DragonRuby Game Toolkit. Pretty decent results!](https://www.reddit.com/r/ruby/comments/iqrddq/took_a_performance_pass_over_dragonruby_game/)
- url: https://www.youtube.com/watch?v=UuY7CWdvyWM&amp;feature=youtu.be&amp;ab_channel=AmirRajan
---

