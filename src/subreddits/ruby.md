# ruby
## [1][12 ways to call a method in Ruby](https://www.reddit.com/r/ruby/comments/i9hvjj/12_ways_to_call_a_method_in_ruby/)
- url: https://www.notonlycode.org/12-ways-to-call-a-method-in-ruby/
---

## [2][Is this perfectly legal to do in Ruby? (Recently switched from PHP to Ruby/Sinatra)](https://www.reddit.com/r/ruby/comments/i98y7u/is_this_perfectly_legal_to_do_in_ruby_recently/)
- url: https://www.reddit.com/r/ruby/comments/i98y7u/is_this_perfectly_legal_to_do_in_ruby_recently/
---
Is this perfectly legal to do in Ruby? (Recently switched from PHP to Ruby/Sinatra)

Then call `ClientHelper.get_instance` whenever I need a client anywhere. So that I won't be establishing duplicate new connections to an API all the time. and just use one.

    class ClientHelper
      
      def initialize
        if @client == nil
          @client = establish_new_connection(ENV['TOKEN'])
        end
      end
    
      def get_instance
        @client
      end
    
    end
## [3][Updated to Ubuntu 20.04, can't install gems without root access](https://www.reddit.com/r/ruby/comments/i98wok/updated_to_ubuntu_2004_cant_install_gems_without/)
- url: https://www.reddit.com/r/ruby/comments/i98wok/updated_to_ubuntu_2004_cant_install_gems_without/
---
So idk what happened. I updated to Ubuntu 20.04 and then I wanted to learn mechanize. Tried installing it and it said I don't have the write permissions for /var/lib/gems/2.7.0 directory. 

Now to install gems I have to prefix sudo, what fixes are there for this?
## [4][How environment check works in Ruby on Rails?](https://www.reddit.com/r/ruby/comments/i8uipt/how_environment_check_works_in_ruby_on_rails/)
- url: https://medium.com/rubycademy/how-environment-check-works-in-ruby-on-rails-4cfbd0434605
---

## [5][Error in Ruby Shoes When Trying to Run App](https://www.reddit.com/r/ruby/comments/i8ub0h/error_in_ruby_shoes_when_trying_to_run_app/)
- url: https://www.reddit.com/r/ruby/comments/i8ub0h/error_in_ruby_shoes_when_trying_to_run_app/
---
I'm getting the following error when I'm trying to run my Ruby Shoes app:

&amp;#x200B;

**Error in &lt;unknown&gt; line 0 | 2020-08-13 15:38:53 +1000**

**cannot load such file -- ocr\_space**

**/Applications/Shoes.app/Contents/MacOS/lib/ruby/2.3.0/rubygems/core\_ext/kernel\_require.rb:55:in \`require'**

**/Applications/Shoes.app/Contents/MacOS/lib/ruby/2.3.0/rubygems/core\_ext/kernel\_require.rb:55:in \`require'**

**shoes.rb:1:in \`&lt;main&gt;'**

**/Applications/Shoes.app/Contents/MacOS/lib/shoes.rb:353:in \`eval'**

**/Applications/Shoes.app/Contents/MacOS/lib/shoes.rb:353:in \`visit'**

**/Applications/Shoes.app/Contents/MacOS/lib/shoes.rb:139:in \`show\_selector'**

**/Applications/Shoes.app/Contents/MacOS/lib/shoes.rb:169:in \`block (4 levels) in splash'**
## [6][How to build a one-time passcode protected conference line with Twilio Verify and Ruby](https://www.reddit.com/r/ruby/comments/i8owzn/how_to_build_a_onetime_passcode_protected/)
- url: https://www.reddit.com/r/ruby/comments/i8owzn/how_to_build_a_onetime_passcode_protected/
---
Protecting a conference line is one thing, but in writing this tutorial I actually discovered how little work it is to verify a user's phone number (and with a little bit of setup, email) with Twilio Verify. [Check out the post here](https://www.twilio.com/blog/one-time-passcode-protected-conference-line-twilio-verify-ruby) and let me know what you think.
## [7][HexaPDF 0.12.0 released - With support for interactive forms and a CLI command for filling forms](https://www.reddit.com/r/ruby/comments/i8fmri/hexapdf_0120_released_with_support_for/)
- url: https://hexapdf.gettalong.org/news/2020/hexapdf-0-12-0-acroform.html
---

## [8][First project](https://www.reddit.com/r/ruby/comments/i8ihi8/first_project/)
- url: https://www.reddit.com/r/ruby/comments/i8ihi8/first_project/
---
So I keep being told that the best way to teach myself the language is to pic projects and learn by doing. The idea that I've had for my first project is just a simple ticker to help keep track of my workout habits.  
I want it to be two simple buttons, one that +1 to the running total of like days in a row, and one that I can click to clear the total. I punched a few questions into google to try and find some videos or explanations on how to do it, but I can't seem to get anything that describes what I'm looking for. Can anyone help with maybe proper search terms? I don't want to make anyone hold my hand through it, I just want to dig in.
## [9][Say hello to RuboCop::Packaging! 👋](https://www.reddit.com/r/ruby/comments/i87l2l/say_hello_to_rubocoppackaging/)
- url: https://www.reddit.com/r/ruby/comments/i87l2l/say_hello_to_rubocoppackaging/
---
Hello all,

We (the Debian core devs (DDs)) have been facing a lot of problems whilst maintaining Ruby libraries and applications in Ruby. We maintain around 1600 packages in the Debian archive.  
Some of the problems, for example, are using a relative path for `require_relative` calls from specs (tests directory) to the lib directory (`require_relative '../lib/foo'`) and others include using `git` in the gemspec files. And so on..

So we came up with the idea of writing a linter (an extension to RuboCop) for flagging those calls which can be fixed upstream (so a win-win situation for both, Downstream and Upstream!).  
And this is how RuboCop::Packaging (https://github.com/utkarsh2102/rubocop-packaging) is born! 💖

Please give it a look, let me know what you think about it?  
In case you find it good, please give it a ⭐, so the Downstream maintainers know that it was worth it!    
Suggestions and opinions are heartily welcomed! 🤗
## [10][Learn Ruby Regexp with hundreds of examples and exercises](https://www.reddit.com/r/ruby/comments/i85hm1/learn_ruby_regexp_with_hundreds_of_examples_and/)
- url: https://www.reddit.com/r/ruby/comments/i85hm1/learn_ruby_regexp_with_hundreds_of_examples_and/
---
Hello!

I updated my Ruby Regexp book recently, changes include updating to 2.7.1, adding epub/html version, new sections on escape sequences, conditional grouping, `\R`, etc. There's also significant increase in number of exercises and this time I added solutions for reference as well. The entire book can be read online at

https://learnbyexample.github.io/Ruby_Regexp/

Or, you can get free pdf/epub versions from:

* https://gumroad.com/l/rubyregexp
* https://leanpub.com/rubyregexp

I made **all my books free** at the end of March when the pandemic hit my country. And then I spent more than four months to update all my books, this is the last one. You can get all my books as a bundle (Python/Ruby/JS regex and grep/sed/awk cli tools) for free until this weekend:

* https://gumroad.com/l/regex
* https://leanpub.com/b/regex

I’d highly appreciate your feedback. That’s been a major motivating factor to keep writing as well as for improving the content.

Happy learning :)
