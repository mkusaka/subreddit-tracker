# ruby
## [1][I'm reading Why's Poignant Guide to Ruby. The preeventualist.org website from Chapter 6 is down. How do i make use of this Github alternative provided by mistydemeo?](https://www.reddit.com/r/ruby/comments/fmxz1t/im_reading_whys_poignant_guide_to_ruby_the/)
- url: https://www.reddit.com/r/ruby/comments/fmxz1t/im_reading_whys_poignant_guide_to_ruby_the/
---
It seems mistydemeo created a replica of the original website to help people reading the Poignant Guide practise on it. Here it is:  [https://github.com/mistydemeo/preeventualist](https://github.com/mistydemeo/preeventualist) 

How do I make use of this github alternative? I'm a total newb at Ruby &amp; programming in general.

Also, Chapter 6 from the guide, jfyr:  [https://poignant.guide/book/chapter-6.html](https://poignant.guide/book/chapter-6.html)
## [2][Difference between “super” and “super()” in Ruby](https://www.reddit.com/r/ruby/comments/fmktyo/difference_between_super_and_super_in_ruby/)
- url: https://dannyh79.github.io/posts/ruby-super-super
---

## [3][Im trying to get into a learning curve for Ruby](https://www.reddit.com/r/ruby/comments/fmt2oe/im_trying_to_get_into_a_learning_curve_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fmt2oe/im_trying_to_get_into_a_learning_curve_for_ruby/
---
The company I was just laid off from uses Ruby in most applications, and if I want to move up and be of more use I need to get beyond the copy and paste of it all and start understanding how to write it and make it functional for myself.


I've gotten through TryRuby and I bought some books, but I'm a hands on learner who can see something done and then learn from that. Any resources you can recommend, or small projects I can start with?
## [4][Wrote a CLI app to fetch and compare numbers on coronavirus with my time in quarantine.](https://www.reddit.com/r/ruby/comments/fmg7u3/wrote_a_cli_app_to_fetch_and_compare_numbers_on/)
- url: https://github.com/siaw23/kovid
---

## [5][Ruby 2.5 acceptable for learning modern Ruby?](https://www.reddit.com/r/ruby/comments/fmqcgd/ruby_25_acceptable_for_learning_modern_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fmqcgd/ruby_25_acceptable_for_learning_modern_ruby/
---
I'm on OpenSUSE and the available version of Ruby by default is 2.5. Is that acceptable for learning modern Ruby, or am I missing on on serious goodies in 2.6? I'm not interested in learning anything Rails related, if that makes any difference.
## [6][Writing an array with %w](https://www.reddit.com/r/ruby/comments/fms17m/writing_an_array_with_w/)
- url: https://www.reddit.com/r/ruby/comments/fms17m/writing_an_array_with_w/
---
Hey y'all total noob here but something just occurred to me and I fired up pry to play with no luck.

OK so I know you can write

`my_string = %w(The brown fox is lost)`

`=&gt; ["The", "brown", "fox", "is", "lost"]`

But is there a way to inject a variable into the %w(variable) without it thinking it's a string? i.e.

`my_string = "The brown fox is lost"`

`new_arr = %w(my_string)`

Hopefully get this?

`new_arr`

`=&gt; ["The", "brown", "fox", "is", "lost"]` 

And I know it's no biggie I can always call the split(' ') on the string, I was only curious, seems like a nice lil shorthand especially when getting user input, no?

Thanks in advance.
## [7][Are Companies still hiring?](https://www.reddit.com/r/ruby/comments/fmi8oc/are_companies_still_hiring/)
- url: /r/rails/comments/fm6nwe/are_companies_still_hiring/
---

## [8][Does concatenating strings need a parenthesis around it?](https://www.reddit.com/r/ruby/comments/fm4zhe/does_concatenating_strings_need_a_parenthesis/)
- url: https://www.reddit.com/r/ruby/comments/fm4zhe/does_concatenating_strings_need_a_parenthesis/
---
I am following a ruby tutorial, and in the tutorial the instructor says if I want  to print out a string concatenated with a variable I need a parenthesis around the whole thing like this:

    name = "Joe"
    puts ("Hello my name is " + name)

However when I try this

    name = "Joe"
    puts "Hello my name is " + name

I don't get an error and the program runs fine. Is the instructor wrong? do I not need  the parenthesis?
## [9][[Rant] I've never written a line of of this language but I hate it with a passion.](https://www.reddit.com/r/ruby/comments/fmnau7/rant_ive_never_written_a_line_of_of_this_language/)
- url: https://www.reddit.com/r/ruby/comments/fmnau7/rant_ive_never_written_a_line_of_of_this_language/
---
WHY IS IT SO HARD TO INSTALL RVM AND RUBY ????? 

I was trying to help a friend with a project installation that had ruby as a dependency. I have never had so much trouble installing anything.

The installation process finds a very convenient way to fail somehow, I've tried so many tutorials and guides for installation but every time I make a little progress something fails in the next step.

I know I'm not a total idiot, no other installation process has given me so much pain. Compare this to nvm the tool for managing Node versions and it works like a charm without any hassle, why the fuck is rvm so special. I just ran the instructions like usual, somehow managed to install rvm!!!! but then tried installing Ruby and that step failed?! Because somehow `sudo apt update` has an issue that I never encountered with any other software and I've been using this particular installation of ubuntu for over a year. 

I'm sure there's great things about this language but I'd never know because I can't run it.
## [10][How to add class and method to rails logger](https://www.reddit.com/r/ruby/comments/flwc21/how_to_add_class_and_method_to_rails_logger/)
- url: https://www.reddit.com/r/ruby/comments/flwc21/how_to_add_class_and_method_to_rails_logger/
---
I have a rails app I want to add a custom log formatter to.  In the logs, I want to dynamically show the class where code was called and the method that called it.  Currently, in the initialize method of classes in my app I'm doing \`@logger = self.progname\` to get the class and \`@logger.level("#{\_\_callee\_\_}: some message"\` to  get the method.  

&amp;#x200B;

I can write a custom logging class and assign it to \`config.log\_formatter\` in .../config/environments/production.rb which is fine but how can I add a formatter that will dynamically tell me what class and what method I'm in when I log stuff?
