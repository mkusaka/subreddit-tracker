# ruby
## [1][16 New ML Gems for Ruby](https://www.reddit.com/r/ruby/comments/et2bvc/16_new_ml_gems_for_ruby/)
- url: https://ankane.org/new-ml-gems
---

## [2][Ruby ML for Python Coders](https://www.reddit.com/r/ruby/comments/et2bbh/ruby_ml_for_python_coders/)
- url: https://ankane.org/ruby-ml-for-python-coders
---

## [3][How to Have Heroku Automatically Run Database Migrations Before a Rails App is Deployed](https://www.reddit.com/r/ruby/comments/est6ks/how_to_have_heroku_automatically_run_database/)
- url: https://scottbartell.com/2020/01/22/automating-rails-database-migrations-on-heroku/
---

## [4][Trying to learn Polymorphism/Duck Tying/Inheritance file structure](https://www.reddit.com/r/ruby/comments/et3izy/trying_to_learn_polymorphismduck_tyinginheritance/)
- url: https://www.reddit.com/r/ruby/comments/et3izy/trying_to_learn_polymorphismduck_tyinginheritance/
---
Hello,

I'm trying to learn more about Polymorphism in Ruby, and my main question is **what should the file structure look like?** Most examples online don't really include the file structure, they just list it all out.

As of right now, I have all of my classes in a single file for this Polymorphism example. This file is in the **lib** directory.

My nooby example:

I have a base "generic class" which is called FoodType. This has a single method called get\_food which takes in a "food type" as an argument.

I have multiple various "food types" that are different classes (Fruit, Vegetable, Dairy, Meat etc.) and these all have their own get\_food methods (which just gets some different strings with examples of the food).

So essentially, when I call FoodType.get\_food( give it a food\_type class here) it ends up using the correct food type sub class and outputs the correct example. I have this working. This is duck typed.

* But my question is HOW do I structure this in the lib directory?
* I'm guessing I should first make a food\_type.rb, and this file should only contain my generic FoodType class. I should also make separate files for meat.rb, vegetable.rb, etc.
* Do I make another directory in lib for these files? Do I also have to change the class of the subtypes to like FoodType::Meat ??
* My code is working and it's duck typed, so I'm not using inheritance, so I'm not using this notation "&lt;" in the class names. But should I instead use inheritance?
* I have no idea what the file structure/class name syntax is suppose to look like when I have to separate this into multiple files/directories. Please please please give me an example of how this would work.

Thank you
## [5][What are some important ruby topics to study before a Ruby Software Developer interview?](https://www.reddit.com/r/ruby/comments/esusjq/what_are_some_important_ruby_topics_to_study/)
- url: https://www.reddit.com/r/ruby/comments/esusjq/what_are_some_important_ruby_topics_to_study/
---
If you are interviewing someone for the role of an experienced ruby developer, what topics would you want the candidate to know thoroughly in Ruby / Rails ?
## [6][Help with running Jekyll - having trouble with gems and ruby set up - MAC OS](https://www.reddit.com/r/ruby/comments/et06vx/help_with_running_jekyll_having_trouble_with_gems/)
- url: https://www.reddit.com/r/ruby/comments/et06vx/help_with_running_jekyll_having_trouble_with_gems/
---
Hello,

I am a beginner to ruby and gems. I was trying to do `jekyll -v`  then it showed this error message:

&gt;Ignoring eventmachine-1.2.7 because its extensions are not built. Try: gem pristine eventmachine --version 1.2.7  
&gt;  
&gt;Ignoring ffi-1.11.1 because its extensions are not built. Try: gem pristine ffi --version 1.11.1  
&gt;  
&gt;Ignoring http\_parser.rb-0.6.0 because its extensions are not built. Try: gem pristine http\_parser.rb --version 0.6.0  
&gt;  
&gt;Ignoring sassc-2.2.1 because its extensions are not built. Try: gem pristine sassc --version 2.2.1  
&gt;  
&gt;Traceback (most recent call last):  
&gt;  
&gt;21: from /Users/H.QIAN/gems/bin/jekyll:23:in \`&lt;main&gt;'  
&gt;  
&gt;20: from /Users/H.QIAN/gems/bin/jekyll:23:in \`load'  
&gt;  
&gt;19: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/jekyll-4.0.0/exe/jekyll:8:in \`&lt;top (required)&gt;'  
&gt;  
&gt;18: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;17: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;16: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/jekyll-4.0.0/lib/jekyll.rb:206:in \`&lt;top (required)&gt;'  
&gt;  
&gt;15: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;14: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;13: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/jekyll-sass-converter-2.0.1/lib/jekyll-sass-converter.rb:4:in \`&lt;top (required)&gt;'  
&gt;  
&gt;12: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;11: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;10: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/jekyll-sass-converter-2.0.1/lib/jekyll/converters/scss.rb:3:in \`&lt;top (required)&gt;'  
&gt;  
&gt;9: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;8: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;7: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/sassc-2.2.1/lib/sassc.rb:31:in \`&lt;top (required)&gt;'  
&gt;  
&gt;6: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/sassc-2.2.1/lib/sassc.rb:31:in \`require\_relative'  
&gt;  
&gt;5: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/sassc-2.2.1/lib/sassc/native.rb:3:in \`&lt;top (required)&gt;'  
&gt;  
&gt;4: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;3: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;2: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/ffi-1.11.1/lib/ffi.rb:4:in \`&lt;top (required)&gt;'  
&gt;  
&gt;1: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;/Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require': cannot load such file -- 2.6/ffi\_c (LoadError)  
&gt;  
&gt;22: from /Users/H.QIAN/gems/bin/jekyll:23:in \`&lt;main&gt;'  
&gt;  
&gt;21: from /Users/H.QIAN/gems/bin/jekyll:23:in \`load'  
&gt;  
&gt;20: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/jekyll-4.0.0/exe/jekyll:8:in \`&lt;top (required)&gt;'  
&gt;  
&gt;19: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;18: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;17: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/jekyll-4.0.0/lib/jekyll.rb:206:in \`&lt;top (required)&gt;'  
&gt;  
&gt;16: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;15: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;14: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/jekyll-sass-converter-2.0.1/lib/jekyll-sass-converter.rb:4:in \`&lt;top (required)&gt;'  
&gt;  
&gt;13: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;12: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;11: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/jekyll-sass-converter-2.0.1/lib/jekyll/converters/scss.rb:3:in \`&lt;top (required)&gt;'  
&gt;  
&gt;10: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;9: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;8: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/sassc-2.2.1/lib/sassc.rb:31:in \`&lt;top (required)&gt;'  
&gt;  
&gt;7: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/sassc-2.2.1/lib/sassc.rb:31:in \`require\_relative'  
&gt;  
&gt;6: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/sassc-2.2.1/lib/sassc/native.rb:3:in \`&lt;top (required)&gt;'  
&gt;  
&gt;5: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;4: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;3: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/ffi-1.11.1/lib/ffi.rb:3:in \`&lt;top (required)&gt;'  
&gt;  
&gt;2: from /Users/H.QIAN/.gem/ruby/2.6.0/gems/ffi-1.11.1/lib/ffi.rb:6:in \`rescue in &lt;top (required)&gt;'  
&gt;  
&gt;1: from /Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require'  
&gt;  
&gt;/Users/H.QIAN/.rbenv/versions/2.6.2/lib/ruby/2.6.0/rubygems/core\_ext/kernel\_require.rb:54:in \`require': incompatible library version - /Users/H.QIAN/.gem/ruby/2.6.0/gems/ffi-1.11.1/lib/ffi\_c.bundle (LoadError)

then, I tried to do `gem install ffi` then it shows different message when I tried to run `jekyll -v`

&amp;#x200B;

 

&gt;Ignoring http\_parser.rb-0.6.0 because its extensions are not built. Try: gem pristine http\_parser.rb --version 0.6.0  
&gt;  
&gt;Ignoring sassc-2.2.1 because its extensions are not built. Try: gem pristine sassc --version 2.2.1  
&gt;  
&gt;Ignoring ffi-1.11.1 because its extensions are not built. Try: gem pristine ffi --version 1.11.1  
&gt;  
&gt;jekyll 4.0.0

which actually shows me the version name but I am not sure what I should do with all the error messages I have received above. 

&amp;#x200B;

Can someone explain what is going on? And will it be possible to restart everything and start to install gems again and a new version of ruby too?

&amp;#x200B;

Thanks in advance for any help.
## [7][Sharing an online tool to generate dry-structs from JSON / JSON Schema / GraphQL definitions](https://www.reddit.com/r/ruby/comments/esrvbr/sharing_an_online_tool_to_generate_drystructs/)
- url: https://app.quicktype.io/
---

## [8][Rails 6 adds touch option to has_one association](https://www.reddit.com/r/ruby/comments/esveb5/rails_6_adds_touch_option_to_has_one_association/)
- url: https://www.abhaynikam.in/posts/rails-6-adds-touch-option-to-has-one-association/
---

## [9][I'm a Ruby developer looking to relocate (EUR or USA), please give me feedback on my CV](https://www.reddit.com/r/ruby/comments/esw7eb/im_a_ruby_developer_looking_to_relocate_eur_or/)
- url: https://www.reddit.com/r/ruby/comments/esw7eb/im_a_ruby_developer_looking_to_relocate_eur_or/
---
[https://pablo-vizcay.tech/](https://pablo-vizcay.tech/)  


Any feedback appreciated.. good, bad, suggestions.
## [10][Would sorting a sorting an array linearly using multi threading speed up it's sorting time?](https://www.reddit.com/r/ruby/comments/esrrwr/would_sorting_a_sorting_an_array_linearly_using/)
- url: https://www.reddit.com/r/ruby/comments/esrrwr/would_sorting_a_sorting_an_array_linearly_using/
---
Say for instance, we have an array X. If we split X in half and then sort each of those halves concurrently using different threads and finally combine them, would the sorting time drop from O(n) to O(n/2)?

Or is O(n/2) ideally just O(n) with ½ as the constant factor?
