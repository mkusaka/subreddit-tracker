# ruby
## [1][Best guide/tutorial/course for Ruby?](https://www.reddit.com/r/ruby/comments/enyhnr/best_guidetutorialcourse_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/enyhnr/best_guidetutorialcourse_for_ruby/
---
Hi, I'm a frontend developer, and this year I want to learn backend development, I choosed Ruby, with Ruby on Rails

But for now I want to learn Ruby, can you please recommend the best Ruby guide/tutorial/course for you?

Thanks.
## [2][Trying to understand advantages of dry-monads](https://www.reddit.com/r/ruby/comments/eo4829/trying_to_understand_advantages_of_drymonads/)
- url: https://www.reddit.com/r/ruby/comments/eo4829/trying_to_understand_advantages_of_drymonads/
---
Hey there!

Today i decided to try dry-monads for the first time but i don't really get if it's advantage is purely around visual aspects of the code or is there more to it? I'm basing my opinion on the only real-life examples i was able to find on the internet which was always about a following refactoring or similar:

Original code:

    if current_user &amp;&amp; current_user.info &amp;&amp; current_user.info.bio
      return current_user.info.bio 
    end

Refactored code (i'm pretty sure that's not the actual dry-monads code, just a pseudo-code to visualize what i'm talking about:

    return Maybe(current_user).map { |user| user.info }
                              .map { |info| info.bio }
                              .or_else { 'No bio set' }

And that makes me think...what's the profit? Is that better than writing it as below?

    return current_user&amp;.info&amp;.bio || 'No bio set'

Pardon my ignorance, it's probably because i can't find any real-life examples of some more complex refactors using dry family. No matter what - i have a strong feeling i'm not getting something in here fully. Thanks in advance for any clues!
## [3][Reality Show on Ruby Programming](https://www.reddit.com/r/ruby/comments/enywfx/reality_show_on_ruby_programming/)
- url: https://www.reddit.com/r/ruby/comments/enywfx/reality_show_on_ruby_programming/
---
Hi

I'm doing a video series on creating a web application using Ruby on Rails. It's not meant to be a tutorial but more of a reality TV show on how one would go about creating an application.

I was just curious if any of you know anything similar (preferably in Ruby) or anyone doing the same thing (coding each step of the way w/ mistakes). 

Playlist link here:

[https://www.youtube.com/playlist?list=PL2-7U6BzddIZ35bJdCFx6RZ-QR8n\_JD82](https://www.youtube.com/playlist?list=PL2-7U6BzddIZ35bJdCFx6RZ-QR8n_JD82)

Videos ongoing. Trying to do a video once a week or if I have time. For this one, I'm trying to build a book keeping system.
## [4][How to use Ruby with VsCode and RuboCop?](https://www.reddit.com/r/ruby/comments/envc0p/how_to_use_ruby_with_vscode_and_rubocop/)
- url: https://www.reddit.com/r/ruby/comments/envc0p/how_to_use_ruby_with_vscode_and_rubocop/
---
I am so confused at how to do this. I normally used iTerm with Ruby, but I've recently moved to VsCode.

I've installed Rubocop, and I'm trying to have it auto format when you press cmd + shift + p. However, in VsCode this opens up the command pallette. It doesn't auto format.

RuboCop is installed and it's working. I can see it making recommendations and underlining things in blue. But the cmd + shift + p command brings up the VsCode command pallette instead of auto formatting. How do I make this happen?

Thanks
## [5][Ruby Tutorial - Building a Blog Title Generator](https://www.reddit.com/r/ruby/comments/enojyf/ruby_tutorial_building_a_blog_title_generator/)
- url: https://www.reddit.com/r/ruby/comments/enojyf/ruby_tutorial_building_a_blog_title_generator/
---
Hi guys, I normally post a lot of Ruby on Rails builds on YouTube but this time I've decided to build something in pure Ruby. It's very beginner friendly but might be useful to those who dived straight into Rails without spending a ton of time learning the Ruby language. In the video I build a Ruby file that generates viral blog title ideas and returns them back to the user, the titles are based on the keyword that the user inputs via Terminal.

This could be easily adapted and used within a Rails app. Hopefully it's useful to some of you guys.

[https://www.youtube.com/watch?v=BC7VaxGkihY](https://www.youtube.com/watch?v=BC7VaxGkihY)

I'm planning to post some more Ruby content so would be open to suggestions about content ideas.
## [6][The Ruby Bibliography - Academic writing on the Ruby programming language](https://www.reddit.com/r/ruby/comments/eneopo/the_ruby_bibliography_academic_writing_on_the/)
- url: https://rubybib.org/
---

## [7][Profit and Loss Function: No Output](https://www.reddit.com/r/ruby/comments/enkzbx/profit_and_loss_function_no_output/)
- url: https://www.reddit.com/r/ruby/comments/enkzbx/profit_and_loss_function_no_output/
---
Hello all, good day to you.

I am trying to create a function that prints profit, loss, or break even based on input (integer) provided. For some reason I do not get any output. I feel like I am overlooking something that is very obvious. I apologise if I am and thank you for your help.

&amp;#x200B;

Please find below two approaches I have used:

&amp;#x200B;

Approach 1:

def accounts(input)  
   
   cost= 21000000  
   price= 3000000  
   revenue = sales \* price  
   totalCosts = cost + insurance  
   
 if revenue &gt; totalCosts  
 return "Profit"  
 elsif revenue ==totalCosts  
 return "Broke Even"  
 else  
 return "Loss"  
 end  
   
 return accounts  
   
end  
puts accounts(9)

&amp;#x200B;

Approach 2:

cost = 21000000

price = 3000000

def accounts()

result = ''

revenue = sales \* price

&amp;#x200B;

totalCosts = cost + insurance

if revenue &gt; totalCosts

result= 'Profit'

else if revenue == totalCosts

result ="Broke Even"

else

result ="Loss"

end

&amp;#x200B;

print result

end
## [8][puma is shipped with rails, why wasn't unicorn?](https://www.reddit.com/r/ruby/comments/enjjlm/puma_is_shipped_with_rails_why_wasnt_unicorn/)
- url: https://www.reddit.com/r/ruby/comments/enjjlm/puma_is_shipped_with_rails_why_wasnt_unicorn/
---
Why was webbrick (something not production ready) used instead of unicorn?
## [9][Is it just me or is the sidebar missing on the web version of this subreddit?](https://www.reddit.com/r/ruby/comments/enjfg8/is_it_just_me_or_is_the_sidebar_missing_on_the/)
- url: https://www.reddit.com/r/ruby/comments/enjfg8/is_it_just_me_or_is_the_sidebar_missing_on_the/
---
Noticed that in my Relay app i can access the sidebar, but when accessing the sub via my PC's browser, i don't see any sidebar at all.
## [10][Building an Asynchronous Scheduler for Ruby](https://www.reddit.com/r/ruby/comments/en4pyh/building_an_asynchronous_scheduler_for_ruby/)
- url: https://www.youtube.com/watch?v=OlmySf03GUo&amp;feature=youtu.be
---

