# ruby
## [1][What book/resources would you recommend for intermediate Rails developer?](https://www.reddit.com/r/ruby/comments/f75sei/what_bookresources_would_you_recommend_for/)
- url: https://www.reddit.com/r/ruby/comments/f75sei/what_bookresources_would_you_recommend_for/
---
I’ve been developing in Rails for about a year now and feel I’m ready to learn intermediate stuff. I have read Michael Hartl’s book, and did a couple of tutorials on udemy as well. What resources/books would you recommend? Thanks!
## [2][Arrays of empty arrays - can someone explain why my code was doing this please?](https://www.reddit.com/r/ruby/comments/f6ysig/arrays_of_empty_arrays_can_someone_explain_why_my/)
- url: https://www.reddit.com/r/ruby/comments/f6ysig/arrays_of_empty_arrays_can_someone_explain_why_my/
---
So I came across some behavior in my code today I wanted to ask about.

I was trying to create an array of n empty arrays. My first attempt was:

    arr = Array.new(5, [])

On the surface this worked. However, the problem I had with this was that when I pushed to one of the subarrays, it pushed to all of them.

So:

    arr[0].push(1)

gave me:

    [[1],[1],[1],[1],[1]]

 I got the same result with:

    arr = Array.new(5, Array.new(0))

I'm guessing the reason for this is because creating the array using those methods means all the array elements are pointing to the same empty subarray, but I don't understand why. Could someone explain what's going on behind the scenes?

I ended up solving my problem by with:

    arr = []
    5.times do
    arr.push([])
    end

but again, I'm not sure why this worked and the other approach didn't.

Any help would be greatly appreciated, including whether there's a more idiomatic way of achieving this.
## [3][Help with gets method](https://www.reddit.com/r/ruby/comments/f77sov/help_with_gets_method/)
- url: https://www.reddit.com/r/ruby/comments/f77sov/help_with_gets_method/
---
Hey everyone,

I'm a rookie needing some help. I am currently going through [Learn to Program, Second Edition by Chris Pine](http://www.r-5.org/files/books/computers/languages/ruby/main/Chris_Pine-Learn_to_Program-EN.pdf). I'm on chapter 5.3 where the gets method is introduced. Here is the code the books says to input:

    puts gets Is there an echo in here?

Now I'm pretty sure when I type my program into terminal my output should be the same text right? It should say "Is there an echo in here?"
Instead I am getting an error message:

    get.rb:1: syntax error, unexpected in, expecting end-of-input
     puts gets Is there an echo in here?

I'm using Sublime and the default Ruby version that came on my MacBook. If anyone can help or steer me in the right direction that would be very much appreciated.

EDIT: So I made sure to press enter or "return" as it is on mac for the code:

    puts gets 
    Is there an echo in here?

I get first a blank line on the terminal.
So i run the program again and then get this message:

    get.rb:2: syntax error, unexpected in, expecting end-of-input
    Is there an echo in here?

If i change the text:

    puts gets
    random text

I get a blank line as before.
Then I run again and get this message:
    ruby get.rb
    ruby get.rb
    Traceback (most recent call last):
    get.rb:2:in `&lt;main&gt;': undefined local variable or method `text' for main:Object (NameError)
    Did you mean?  test
## [4][Calculate current week with offset?](https://www.reddit.com/r/ruby/comments/f77ljw/calculate_current_week_with_offset/)
- url: https://www.reddit.com/r/ruby/comments/f77ljw/calculate_current_week_with_offset/
---
Hi people,

I am currently working on a little project to plot some data from an google sheet. Because I need some history when the sheet changes I began to extract them an put it into a CSV file which got the current week as "primary index".

Now the weekly "rollover" inside the sheet is not monday or sunday, but wednesday at around 2-3pm.

My current way to approach is to get the current week and just +=1 when certain conditions are met.

    week_now  = Time.now.strftime("%W").to_i
    week_now += 1 if Time.now.wday &gt;= 4 or (Time.now.wday.to_i &gt;= 3 and Time.now.hour.to_i &gt;= 15)

Maybe there is some better way? 
## [5][Defibrilating Web Security [PDF]](https://www.reddit.com/r/ruby/comments/f749h1/defibrilating_web_security_pdf/)
- url: https://conference.hitb.org/hitbsecconf2012kul/materials/D1T2%20-%20Meder%20Kydyraliev%20-%20Defibrilating%20Web%20Security.pdf
---

## [6][Clean your Rails routes: nesting](https://www.reddit.com/r/ruby/comments/f77n9c/clean_your_rails_routes_nesting/)
- url: https://medium.com/@farsi_mehdi/clean-your-rails-routes-nesting-cd8528bbc344
---

## [7][Extended ruby grammar and Gruvbox-ish theme for vscode.](https://www.reddit.com/r/ruby/comments/f6z10w/extended_ruby_grammar_and_gruvboxish_theme_for/)
- url: https://www.reddit.com/r/ruby/comments/f6z10w/extended_ruby_grammar_and_gruvboxish_theme_for/
---
Hey there.

I'm trying to expand Ruby syntax highlighting and make it possible to highlight different code structures that are closes with `end` keyword in different colors.

For now I created a gruvbox-like theme for vscode which uses this extended grammar file. In this theme `module`, `class` and `def` structures highlights with another color. If one of these structures closing with red colored `end` keyword, it means something wrong with grammar file and I'd be glad if you could show what code caused highlighting error.

If anybody likes this color scheme and wants to help test this grammar extension, I'd like to ask you to send me any feedback if you find any bugs in syntax highlighting.

[Link to the vscode theme](https://marketplace.visualstudio.com/items?itemName=GracefulPotato.gruvbox-ish)

[Link to PR where you can leave your feedback about bugs as well](https://github.com/textmate/ruby.tmbundle/pull/131)
## [8][Clean your Rails routes: grouping](https://www.reddit.com/r/ruby/comments/f6pvrp/clean_your_rails_routes_grouping/)
- url: https://medium.com/@farsi_mehdi/how-to-keep-your-routes-clean-in-ruby-on-rails-f7cf348ec13b
---

## [9][RUBY NEWBIE HERE](https://www.reddit.com/r/ruby/comments/f6t56y/ruby_newbie_here/)
- url: https://www.reddit.com/r/ruby/comments/f6t56y/ruby_newbie_here/
---
run proc {
[200, {'Content-Type' =&gt; 'text/html'}, ["Hello, world!"]]
}

The above is a piece of code that I’m looking at
Correct me and complete me
I am assuming run is a method to which we are passing proc and the other hash?
What is going on up there?
Run is a method for sure ,what is proc and the hash?
Btw proc is not defined anywhere in the program
Cheers !
## [10][If you're to build a library for the language, what would you build?](https://www.reddit.com/r/ruby/comments/f6h4ho/if_youre_to_build_a_library_for_the_language_what/)
- url: https://www.reddit.com/r/ruby/comments/f6h4ho/if_youre_to_build_a_library_for_the_language_what/
---
I'm looking for a concept for my 3rd year project and I thought building a library would be fun. Help a little undergrad out guys?
