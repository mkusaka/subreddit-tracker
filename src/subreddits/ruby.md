# ruby
## [1][Airbnb Clone Ruby Sinatra Microservice Plus Angular 9](https://www.reddit.com/r/ruby/comments/iaedzb/airbnb_clone_ruby_sinatra_microservice_plus/)
- url: https://www.youtube.com/watch?v=FENVw7LjDbY&amp;feature=share
---

## [2][Need help figuring out how to isolate integers from strings in a nested array.](https://www.reddit.com/r/ruby/comments/iam54q/need_help_figuring_out_how_to_isolate_integers/)
- url: https://www.reddit.com/r/ruby/comments/iam54q/need_help_figuring_out_how_to_isolate_integers/
---
Hi friends, first post! I'm learning Ruby and currently going through a bootcamp prework, but I am currently stuck on one of the labs. I need to write a method isolating strings from integers in a given nested array but don't have the vernacular to figure out just how to do that. 

&amp;#x200B;

mixed\_data **=** \[

1. \["The", 4, "quick"\],
2. \[**-**1, "brown", "fox", 30\],
3. \["studied", 101, 233, "Ruby"\]
4. \]

&amp;#x200B;

&amp;#x200B;

So far I have this as my method  


def join\_nested\_strings(src)

  string\_result = \[ \]

  row\_index = 0

  while row\_index &lt; src.count do

element\_index = 0

string\_elements = ""

while element\_index &lt; src\[row\_index\].count do

&amp;#x200B;

if src\[row\_index\]\[element\_index\]

&amp;#x200B;

string\_result &lt;&lt; src\[row\_index\]\[element\_index\]

&amp;#x200B;

element\_index

&amp;#x200B;

end

element\_index += 1

end

row\_index += 1

  end

  string\_result.join

end

&amp;#x200B;

&amp;#x200B;

I am almost 100% confident my issue lies in the lack of code directly underneath the "if" statement, however I'm just genuinely unsure of where to go from here. The output currently is "The4quick-1brownfox30studied101233Ruby" when I need "The quick brown fox studied Ruby."

&amp;#x200B;

Any and all help is appreciated!
## [3][Completing 100% Ruby Koans in 1 Stream](https://www.reddit.com/r/ruby/comments/ia949z/completing_100_ruby_koans_in_1_stream/)
- url: https://twitch.tv/robedcoder
---

## [4][Replace weird characters such as á → a, or Ü → U](https://www.reddit.com/r/ruby/comments/i9vlu5/replace_weird_characters_such_as_á_a_or_ü_u/)
- url: https://www.reddit.com/r/ruby/comments/i9vlu5/replace_weird_characters_such_as_á_a_or_ü_u/
---
Basically I have a requirement where part of the job I have to do is:

1. Receive a string of 200 characters long (which can be in UTF8)
2. Write this string in a text file replacing every weird character, producing 200 characters again, but using only 200 bytes (i.e. one byte per character so can't be UTF8)

Ideally I should only allow ASCII characters in the final file, replacing characters like this: á → a, or Ü → U, but replacing stuff like ¶ or ʧ for a white-space.

Is there a relatively well known way to do this?
## [5][12 ways to call a method in Ruby](https://www.reddit.com/r/ruby/comments/i9hvjj/12_ways_to_call_a_method_in_ruby/)
- url: https://www.notonlycode.org/12-ways-to-call-a-method-in-ruby/
---

## [6][Uploading files to root-owned directories with Capistrano](https://www.reddit.com/r/ruby/comments/i9miq1/uploading_files_to_rootowned_directories_with/)
- url: https://tryhexadecimal.com/journal/capistrano-upload-root-directory
---

## [7][Is this perfectly legal to do in Ruby? (Recently switched from PHP to Ruby/Sinatra)](https://www.reddit.com/r/ruby/comments/i98y7u/is_this_perfectly_legal_to_do_in_ruby_recently/)
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
## [8][Updated to Ubuntu 20.04, can't install gems without root access](https://www.reddit.com/r/ruby/comments/i98wok/updated_to_ubuntu_2004_cant_install_gems_without/)
- url: https://www.reddit.com/r/ruby/comments/i98wok/updated_to_ubuntu_2004_cant_install_gems_without/
---
So idk what happened. I updated to Ubuntu 20.04 and then I wanted to learn mechanize. Tried installing it and it said I don't have the write permissions for /var/lib/gems/2.7.0 directory. 

Now to install gems I have to prefix sudo, what fixes are there for this?
## [9][How environment check works in Ruby on Rails?](https://www.reddit.com/r/ruby/comments/i8uipt/how_environment_check_works_in_ruby_on_rails/)
- url: https://medium.com/rubycademy/how-environment-check-works-in-ruby-on-rails-4cfbd0434605
---

## [10][Error in Ruby Shoes When Trying to Run App](https://www.reddit.com/r/ruby/comments/i8ub0h/error_in_ruby_shoes_when_trying_to_run_app/)
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
