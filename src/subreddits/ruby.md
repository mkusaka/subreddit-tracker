# ruby
## [1][The redo Keyword in Ruby](https://www.reddit.com/r/ruby/comments/f7pi1f/the_redo_keyword_in_ruby/)
- url: https://medium.com/rubycademy/the-redo-keyword-in-ruby-3f150d69e3c2
---

## [2][Aaron Patterson's thoughts on the pipeline emoji](https://www.reddit.com/r/ruby/comments/f7d117/aaron_pattersons_thoughts_on_the_pipeline_emoji/)
- url: https://github.com/ruby/ruby/pull/2273/
---

## [3][Namespacing Rails application?](https://www.reddit.com/r/ruby/comments/f7idfp/namespacing_rails_application/)
- url: https://www.reddit.com/r/ruby/comments/f7idfp/namespacing_rails_application/
---
I maintain a "majestic" monolith eCommerce platform  with 3 very distinct parts -- api, admin, store. And I'd love to namespace this application, so this distinction could be move obvious and I can easily use mutant testing.  


Has anyone went through such a hassle with their Rails apps? Any blog posts/ experience you can share?  


Namespacing seems obvious for controller code, but not that 'obvious' for code that is being shared between all these "distinct parts". Should I namespace models as well?   


Quite a lot of questions once I start thinking about it and very little information on such a 'transition' in the internet.
## [4][Need help with some refactoring in a cfndsl CloudFormation template](https://www.reddit.com/r/ruby/comments/f7kn5v/need_help_with_some_refactoring_in_a_cfndsl/)
- url: https://www.reddit.com/r/ruby/comments/f7kn5v/need_help_with_some_refactoring_in_a_cfndsl/
---
In [cfndsl](https://github.com/cfndsl/cfndsl), you can create an AWS CloudFormation template like this:

```
CloudFormation do
  EC2_Instance(:Example) do
    ImageId 'ami-12345678'
    Type 't1.micro'
  end
end
```

However real-world templates are usually much larger, but that's not the point.

Let's say that I have dozens of CloudFormation templates that share the same `EC2_Instance`, each with a possibly different `ImageId` and `Type`.

How can I extract this into a method (e.g. `ec2_instance`) so that I can refactor the `CloudFormation` block into this:

```
def ec2_instance(ami, type)
   # ???
end

CloudFormation do
  ec2_instance('ami-12345678', 't1.micro')
end

CloudFormation do
  ec2_instance('ami-90111213', 't2.small')
end

```

Thanks in advance.
## [5][Best Developer Tools Trick](https://www.reddit.com/r/ruby/comments/f7czxv/best_developer_tools_trick/)
- url: https://blog.driftingruby.com/best-developer-tool-trick/
---

## [6][Ruby Quiz - Code Challenge #18 - Up-to-Date? Version Check All Your Libraries](https://www.reddit.com/r/ruby/comments/f7dct8/ruby_quiz_code_challenge_18_uptodate_version/)
- url: https://github.com/planetruby/quiz/tree/master/018
---

## [7][What book/resources would you recommend for intermediate Rails developer?](https://www.reddit.com/r/ruby/comments/f75sei/what_bookresources_would_you_recommend_for/)
- url: https://www.reddit.com/r/ruby/comments/f75sei/what_bookresources_would_you_recommend_for/
---
I’ve been developing in Rails for about a year now and feel I’m ready to learn intermediate stuff. I have read Michael Hartl’s book, and did a couple of tutorials on udemy as well. What resources/books would you recommend? Thanks!
## [8][Arrays of empty arrays - can someone explain why my code was doing this please?](https://www.reddit.com/r/ruby/comments/f6ysig/arrays_of_empty_arrays_can_someone_explain_why_my/)
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
## [9][Help with gets method](https://www.reddit.com/r/ruby/comments/f77sov/help_with_gets_method/)
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
## [10][Calculate current week with offset?](https://www.reddit.com/r/ruby/comments/f77ljw/calculate_current_week_with_offset/)
- url: https://www.reddit.com/r/ruby/comments/f77ljw/calculate_current_week_with_offset/
---
Hi people,

I am currently working on a little project to plot some data from an google sheet. Because I need some history when the sheet changes I began to extract them an put it into a CSV file which got the current week as "primary index".

Now the weekly "rollover" inside the sheet is not monday or sunday, but wednesday at around 2-3pm.

My current way to approach is to get the current week and just +=1 when certain conditions are met.

    week_now  = Time.now.strftime("%W").to_i
    week_now += 1 if Time.now.wday &gt;= 4 or (Time.now.wday.to_i &gt;= 3 and Time.now.hour.to_i &gt;= 15)

Maybe there is some better way? 
