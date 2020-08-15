# ruby
## [1][Replace weird characters such as á → a, or Ü → U](https://www.reddit.com/r/ruby/comments/i9vlu5/replace_weird_characters_such_as_á_a_or_ü_u/)
- url: https://www.reddit.com/r/ruby/comments/i9vlu5/replace_weird_characters_such_as_á_a_or_ü_u/
---
Basically I have a requirement where part of the job I have to do is:

1. Receive a string of 200 characters long (which can be in UTF8)
2. Write this string in a text file replacing every weird character, producing 200 characters again, but using only 200 bytes (i.e. one byte per character so can't be UTF8)

Ideally I should only allow ASCII characters in the final file, replacing characters like this: á → a, or Ü → U, but replacing stuff like ¶ or ʧ for a white-space.

Is there a relatively well known way to do this?
## [2][12 ways to call a method in Ruby](https://www.reddit.com/r/ruby/comments/i9hvjj/12_ways_to_call_a_method_in_ruby/)
- url: https://www.notonlycode.org/12-ways-to-call-a-method-in-ruby/
---

## [3][Uploading files to root-owned directories with Capistrano](https://www.reddit.com/r/ruby/comments/i9miq1/uploading_files_to_rootowned_directories_with/)
- url: https://tryhexadecimal.com/journal/capistrano-upload-root-directory
---

## [4][Is this perfectly legal to do in Ruby? (Recently switched from PHP to Ruby/Sinatra)](https://www.reddit.com/r/ruby/comments/i98y7u/is_this_perfectly_legal_to_do_in_ruby_recently/)
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
## [5][Updated to Ubuntu 20.04, can't install gems without root access](https://www.reddit.com/r/ruby/comments/i98wok/updated_to_ubuntu_2004_cant_install_gems_without/)
- url: https://www.reddit.com/r/ruby/comments/i98wok/updated_to_ubuntu_2004_cant_install_gems_without/
---
So idk what happened. I updated to Ubuntu 20.04 and then I wanted to learn mechanize. Tried installing it and it said I don't have the write permissions for /var/lib/gems/2.7.0 directory. 

Now to install gems I have to prefix sudo, what fixes are there for this?
## [6][How environment check works in Ruby on Rails?](https://www.reddit.com/r/ruby/comments/i8uipt/how_environment_check_works_in_ruby_on_rails/)
- url: https://medium.com/rubycademy/how-environment-check-works-in-ruby-on-rails-4cfbd0434605
---

## [7][Error in Ruby Shoes When Trying to Run App](https://www.reddit.com/r/ruby/comments/i8ub0h/error_in_ruby_shoes_when_trying_to_run_app/)
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
## [8][How to build a one-time passcode protected conference line with Twilio Verify and Ruby](https://www.reddit.com/r/ruby/comments/i8owzn/how_to_build_a_onetime_passcode_protected/)
- url: https://www.reddit.com/r/ruby/comments/i8owzn/how_to_build_a_onetime_passcode_protected/
---
Protecting a conference line is one thing, but in writing this tutorial I actually discovered how little work it is to verify a user's phone number (and with a little bit of setup, email) with Twilio Verify. [Check out the post here](https://www.twilio.com/blog/one-time-passcode-protected-conference-line-twilio-verify-ruby) and let me know what you think.
## [9][HexaPDF 0.12.0 released - With support for interactive forms and a CLI command for filling forms](https://www.reddit.com/r/ruby/comments/i8fmri/hexapdf_0120_released_with_support_for/)
- url: https://hexapdf.gettalong.org/news/2020/hexapdf-0-12-0-acroform.html
---

## [10][First project](https://www.reddit.com/r/ruby/comments/i8ihi8/first_project/)
- url: https://www.reddit.com/r/ruby/comments/i8ihi8/first_project/
---
So I keep being told that the best way to teach myself the language is to pic projects and learn by doing. The idea that I've had for my first project is just a simple ticker to help keep track of my workout habits.  
I want it to be two simple buttons, one that +1 to the running total of like days in a row, and one that I can click to clear the total. I punched a few questions into google to try and find some videos or explanations on how to do it, but I can't seem to get anything that describes what I'm looking for. Can anyone help with maybe proper search terms? I don't want to make anyone hold my hand through it, I just want to dig in.
