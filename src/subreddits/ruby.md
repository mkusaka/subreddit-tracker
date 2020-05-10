# ruby
## [1][[Another] An alternative Ruby implementation by Rust.](https://www.reddit.com/r/ruby/comments/ggy0ew/another_an_alternative_ruby_implementation_by_rust/)
- url: https://github.com/sisshiki1969/ruruby
---

## [2][Project on Software development.](https://www.reddit.com/r/ruby/comments/gh0whe/project_on_software_development/)
- url: https://www.reddit.com/r/ruby/comments/gh0whe/project_on_software_development/
---
Looking forward for some software development project to work on.
## [3][Interesting ruby twitter accounts to follow?](https://www.reddit.com/r/ruby/comments/ggly5g/interesting_ruby_twitter_accounts_to_follow/)
- url: https://www.reddit.com/r/ruby/comments/ggly5g/interesting_ruby_twitter_accounts_to_follow/
---

## [4][Setting up an Automatic Book reader with Devise + Rails - 1](https://www.reddit.com/r/ruby/comments/ggzscb/setting_up_an_automatic_book_reader_with_devise/)
- url: https://www.reddit.com/r/ruby/comments/ggzscb/setting_up_an_automatic_book_reader_with_devise/
---
 

Hey guys,

I hate how much coding channels focus on a basic 'ToDo list / blog' so i wanted wanted to share my progress in building an automatic book reader (not that good in making vids but meh). Probably will do several more videos until the end goal of having a tracker for how many books I've read.

[https://www.youtube.com/watch?v=AnDLaaTXeWg&amp;feature=youtu.be](https://www.youtube.com/watch?v=AnDLaaTXeWg&amp;feature=youtu.be)

Thx!
## [5][What GUI library would you recommend for a simple to do ruby app?](https://www.reddit.com/r/ruby/comments/gg5ff5/what_gui_library_would_you_recommend_for_a_simple/)
- url: https://www.reddit.com/r/ruby/comments/gg5ff5/what_gui_library_would_you_recommend_for_a_simple/
---
I am not using Ruby on Rails as it will be just for pc
## [6][Is Ruby a good choice here?](https://www.reddit.com/r/ruby/comments/gg7774/is_ruby_a_good_choice_here/)
- url: https://www.reddit.com/r/ruby/comments/gg7774/is_ruby_a_good_choice_here/
---
I need to automate gathering certain data from Microsoft's reference pages and turn that into a neat C++ wrapper for parts of WinAPI. I could do it manually but at the current pace it would take me a better part of this year to finish it so I think it's time for me to learn how to work with the internet for once.

I've heard that Ruby is much better suited for dealing with this type of web stuff than Python. Is that true? If so, are there any tutorials from the "X for Y programmers" category? (I mean the ones that are at least serviceable, I tend to die a little bit inside every time someone explains to me what a variable is).
## [7][Looping thru Nested Arrays](https://www.reddit.com/r/ruby/comments/gg4lnf/looping_thru_nested_arrays/)
- url: https://www.reddit.com/r/ruby/comments/gg4lnf/looping_thru_nested_arrays/
---
Hi! Really new at this. About to start a bootcamp and need to complete some pre-work.

My current task is as follows: I have a nested array, all with integers, and **I must loop through** each array to iterate and return an array of the smallest numbers of each array.

I know of the map method, but have no idea how to use it. We haven't covered it yet.

Ideas?

**UPDATE:** solved! Thank you everyone for being so dope and helping me. I really appreciate it :)  
I ended up being stuck in an infinite loop. Had to do some screen sharing to figure it out. 
## [8][I'm supposed to find someone who has 10 years experience consistently using Chef - I'm wondering if that is that even possible? Do people even use it anymore and could they have used it that long?](https://www.reddit.com/r/ruby/comments/gfv7d3/im_supposed_to_find_someone_who_has_10_years/)
- url: https://www.reddit.com/r/ruby/comments/gfv7d3/im_supposed_to_find_someone_who_has_10_years/
---

## [9][Ruby Association -&gt; 2019 Grant Accomplishment Report](https://www.reddit.com/r/ruby/comments/gfo8bh/ruby_association_2019_grant_accomplishment_report/)
- url: https://www.ruby.or.jp/en/news/20200508
---

## [10][Interesting behavior of Hash.new](https://www.reddit.com/r/ruby/comments/gfoc59/interesting_behavior_of_hashnew/)
- url: https://www.reddit.com/r/ruby/comments/gfoc59/interesting_behavior_of_hashnew/
---
Hello,

&amp;#x200B;

I was building a script to format some hashes and I had something similar to this:

    data = { monday: { start: '08:00', end: '18:00' }, thuesday: { start: '08:00', end: '15:00' }}
    
    result = data.each_with_object(Hash.new([])) do |(day, schedule), result|
      result[schedule[:start]] &lt;&lt; day
    end

Nothing fancy, but I noticed that my result variable seemed empty, but if I tried to access a key, it was there.

See the following screenshot:

&amp;#x200B;

https://preview.redd.it/6kl2xkpgqhx41.png?width=472&amp;format=png&amp;auto=webp&amp;s=77d2f3a8c166fe4ef6c1e70b99834cac670e7275

I know how to fix it, I replaced [`Hash.new`](https://Hash.new)`([])` with `Hash.new { |hash, key| hash[key] = [] }` and it works as expected.

&amp;#x200B;

But I was still curious why it happens and how come I can access the keys. Any thoughts?

&amp;#x200B;

Thanks
