# ruby
## [1][Programming languages not required!](https://www.reddit.com/r/ruby/comments/ftk7kf/programming_languages_not_required/)
- url: https://medium.com/the-developers-journey/programming-languages-not-required-6fd0422e9dec?source=friends_link&amp;sk=227ce47f96f0aab7dcd803538d500953
---

## [2][Veteran developers, how do you get acclimated to a large new codebase?](https://www.reddit.com/r/ruby/comments/ft5fbc/veteran_developers_how_do_you_get_acclimated_to_a/)
- url: https://www.reddit.com/r/ruby/comments/ft5fbc/veteran_developers_how_do_you_get_acclimated_to_a/
---
I just landed my second Ruby job a few weeks ago. The codebase is HUGE and (for me) mind-bendingly complex, with layers upon layers of ... interesting ... architectural decisions in various stages of deprecation.

For those of you who've been around the block a few times, how do you get acclimated to a new codebase when a new place onboards you? What questions are you asking yourself? Are you looking for or noticing anything in particular that might help you down the line? Are you taking notes?

It definitely takes time and practice. But are there any other ideas that might help in the meantime?

EDIT: The amount of great advice in this thread blows me away. Thanks!
## [3][Terminal Snake](https://www.reddit.com/r/ruby/comments/ft8pxs/terminal_snake/)
- url: https://www.reddit.com/r/ruby/comments/ft8pxs/terminal_snake/
---
I have written a game in Ruby 2.6:

`require'io/console';X,Y=15,10;U,D,L,R=-X,X,-1,1;f=Array.new S=X*Y,'.';`
`t=[h=X*(Y+1)/2];f[h]='*';(g=-&gt;{f[f.zip(0..).select{|v,_|v=='.'}.sample[1]]='+'}).();`
`m,s=U,0;Thread.new{loop{k=STDIN.getch;`
`m=k=='q'&amp;&amp;exit||k=='w'&amp;&amp;U||k=='s'&amp;&amp;D||k=='a'&amp;&amp;L||k=='d'&amp;&amp;R||m}};`
`begin;print"\e[2J\e[?25l";loop{h+=m;m==U&amp;&amp;h&lt;0&amp;&amp;h+=S;m==D&amp;&amp;h&gt;=S&amp;&amp;h-=S;`
`m==R&amp;&amp;h%X==0&amp;&amp;h-=X;m==L&amp;&amp;h%X==X-1&amp;&amp;h+=X;`
`f[h]=='*'&amp;&amp;break;f[h]!='+'?(f[t[0]]='.';t=t.drop(1)):(s+=1;g.());t&lt;&lt;h;f[h]='*';print"\e[H";`
`f.each_slice(X){|c|puts"#{c.join}\e[#{X}D"};puts"+#{s}\e[#{X}D";sleep 0.2}ensure;print"\e[?25h";end`

It should work in any ANSI compatible terminal.

EDIT:

Here is the full source: https://pastebin.com/BzW1WUDP
## [4][Why doesn't my defined function return anything?](https://www.reddit.com/r/ruby/comments/ftb2xe/why_doesnt_my_defined_function_return_anything/)
- url: https://www.reddit.com/r/ruby/comments/ftb2xe/why_doesnt_my_defined_function_return_anything/
---
I am pretty new to Ruby and coding in general so my first project was to make a functional calculator that allows you to input two numbers as well as decide what you wanted to do with them(+, -, \*, /). The calculator was also supposed to keep asking to enter a number until the user asks an actual number and then turn it into either an integer or a float depending on what the user initially entered. (Whether it has a "." or not). The calculator would also allow you to enter the math signs and would keep asking until you entered an actual math sign as apposed to something random. I was able to successfully pull this off without using defined functions however I wanted to update the code with defined functions for fun, since I recently learned them. For some reason when I defined my function and tried running it, it doesn't return anything.

    num1=" "
     
    num2=" "
     
    def onlynums(number, word)
      until number.to_i.to_s == number or number.to_f.to_s == number
        print "Please enter your"+" "+word+" "+"number: "
        number = gets.chomp()
      end
    return number
    end
     
     
    def stringtofloatinteger(number)
      if number.include? "."
        number = number.to_f
      else
        number = number.to_i
      end
    return number
    end
     
     
     
    onlynums(num1, "first")
     
    stringtofloatinteger(num1)
     
     
    onlynums(num2, "second")
     
    stringtofloatinteger(num2)
     
     
    print "Please enter the mathicmatical sign you desire: "
    sign = gets.chomp.to_s
     
    until sign == "+" or sign == "-" or sign == "*" or sign == "/"
      print "Please enter the mathicmatical sign you desire: "
      sign = gets.chomp.to_s
    end
     
     
    if sign == "+"
      answer = num1 + num2
      print "#{num1} + #{num2} = #{answer}"
    elsif sign == "-"
        answer = num1 - num2
        print "#{num1} - #{num2} = #{answer}"
    elsif sign == "*"
        answer = num1 * num2
        print "#{num1} * #{num2} = #{answer}"
    elsif sign == "/"
        answer = num1 / num2
        print "#{num1} / #{num2} = #{answer}"
    else
      print "Error"
    end
## [5][Exploring Method Arguments in Ruby](https://www.reddit.com/r/ruby/comments/ft09xy/exploring_method_arguments_in_ruby/)
- url: https://www.ombulabs.com/blog/ruby/learning/method-s-arguments-pt-1.html
---

## [6][Changing the Approach to Debugging in Ruby with TracePoint](https://www.reddit.com/r/ruby/comments/ft01e8/changing_the_approach_to_debugging_in_ruby_with/)
- url: https://blog.appsignal.com/2020/04/01/changing-the-approach-to-debugging-in-ruby-with-tracepoint.html
---

## [7][Swapping Out Ruby Minor Version Without Re-running Bundler 2 Leads to "Missing Gems"?](https://www.reddit.com/r/ruby/comments/ft97fg/swapping_out_ruby_minor_version_without_rerunning/)
- url: https://www.reddit.com/r/ruby/comments/ft97fg/swapping_out_ruby_minor_version_without_rerunning/
---
We deploy our applications using Capistrano-Bundler and have noticed the last couple of times we've updated a patch Ruby version (e.g. 2.6.5 to 2.6.6) in an existing environment, applications using Bundler 2 cannot find the previously installed gems in the existing /app/shared/bundle directory. Swapping out the Ruby version works fine for applications using Bundler 1. Redeploying the same code (no changes) results in a full \`bundle install\` (unexpected) running successfully with no issues, and the applications come online as expected. While this is a solution to the issue, it's not ideal for simply doing light security-related patching as it causes an outage for the applications.

I was curious if anyone else had run into this sort of situation and maybe could point me in the right direction. I'm guessing it's a Capistrano-Bundler configuration setting that I can adjust?

&amp;#x200B;

&amp;#x200B;

Please note: I do understand how containerization or possibly another alternative deployment method would solve this issue, and while that's another project I'm working on, I'm not switching things up right now. I'm trying to get our customer there, but that's not something that is going to happen soon. Thanks for understanding.
## [8][A 10-day long hackathon called JamCraft 5 starts in less than 9 days. DragonRuby is sponsoring this jam and is providing all participants with a license to Game Toolkit. Here's your chance to build a game with Ruby.](https://www.reddit.com/r/ruby/comments/fsr5s8/a_10day_long_hackathon_called_jamcraft_5_starts/)
- url: http://jamcraft.dragonruby.org/
---

## [9][RailsConf 2020.2 Couch Edition](https://www.reddit.com/r/ruby/comments/fsk4w3/railsconf_20202_couch_edition/)
- url: https://twitter.com/railsconf/status/1245056735892525057?s=20
---

## [10][A new feature proposal today—Endless method definition: def: value(args) = expression](https://www.reddit.com/r/ruby/comments/fstnrz/a_new_feature_proposal_todayendless_method/)
- url: https://bugs.ruby-lang.org/issues/16746
---

