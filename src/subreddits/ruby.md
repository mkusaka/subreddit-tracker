# ruby
## [1][Looking for an online IDE for ruby](https://www.reddit.com/r/ruby/comments/ga7pr8/looking_for_an_online_ide_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/ga7pr8/looking_for_an_online_ide_for_ruby/
---
I'm a hobbyist programmer and I usually use Ruby for automating tasks, scarping, etc..

Sometimes I work on my desktop that has a dev environment setup, but sometimes I'm on my Chromebook.

I tried https://repl.it but I kept having problems. Do you guys have any recommendations?
## [2][Practical use of Ruby PStore — The story of imlementing github data source for static site on nanoc. How the octokit, faraday, concurrent-ruby and pstore fit together to deliver remote posts.](https://www.reddit.com/r/ruby/comments/g9tea8/practical_use_of_ruby_pstore_the_story_of/)
- url: https://blog.arkency.com/practical-use-of-ruby-pstore/
---

## [3][Super Basic Ruby Question (new programmer!)](https://www.reddit.com/r/ruby/comments/g9tptl/super_basic_ruby_question_new_programmer/)
- url: https://www.reddit.com/r/ruby/comments/g9tptl/super_basic_ruby_question_new_programmer/
---
Hey there - I am self teaching Ruby and am getting a syntax error for something that is definitely super basic and I am hoping someone can help:

I bolded where I am getting the syntax error. THANK YOU in advance! 

puts "Hey! Thanks for playing the numbers game. I am going to ask you to provide a few numbers. Please provide a number with no decimal points"  
firstnumber = gets.chomp.to\_i  
if firstnumber &gt;= #{20}  
  **puts "#{firstnumber}** is greater than or equal to 20!"  
**else**  
  puts "#{firstnumber} is less than 20. :("  
end
## [4][Script help with path objects](https://www.reddit.com/r/ruby/comments/ga0e28/script_help_with_path_objects/)
- url: https://www.reddit.com/r/ruby/comments/ga0e28/script_help_with_path_objects/
---
Hi all,

I'm playing around and doing some Ruby scripting for the first time. I've done plenty of PowerShell in the past so I know a little about scripting. Just need a hand on path's. I've got the below script to process some media files, I'm basically transcoding audio to ac3 as my TV will only play with that audio codec.

My main question is how paths are handled on the objects coming out of the globbing? Is is it a straight string or is there a method I can use for getting just the path and not the filename as well. This way I wouldn't need to split the 'file' object to get the current and new filenames.

Any help or commentary on structuring the script would be appreciated. It's my first one in Ruby so I only have a few hours experience in Ruby.

    
    require "streamio-ffmpeg"
    
    # Set path to recursively search directories
    # TODO: Convert to script args
    path = "/home/alex/Downloads/complete"
    
    # Search all subdirectories for mp4 and mkv files to process and begin loop
    Dir.glob(path + "/**/*.{mp4,mkv}").each do |file|
      
      # Create filenames
      old_filename = file.split("/")[-1]
      new_filename = file.split("/")[-2] + ".mkv"
      
      # Instantiate movie object for processing
      media = FFMPEG::Movie.new(file)
      
      # Check got ac3 audio in media file and convert if not found.
      if media.audio_codec != "ac3"
        puts "non ac3 audio codec detected on #{old_filename} converting audio to ac3."
        media.transcode("#{new_filename}", %w(-c:v copy -c:a ac3))
        puts "Completed conversion to ac3 audio for #{new_filename}"
      end
    end
## [5][Help with this ruby code please!](https://www.reddit.com/r/ruby/comments/ga4oxx/help_with_this_ruby_code_please/)
- url: https://www.reddit.com/r/ruby/comments/ga4oxx/help_with_this_ruby_code_please/
---
item = ""

basket = \[\]

while item != "quit"

  puts "What item do you want?"

  item = gets.chomp

  if shopping\_items.key?(item.to\_sym)

basket &lt;&lt; item

p basket

  else

puts "This item is not available"

  end

end

&amp;#x200B;

With this block of quit, when I type "quit" into the terminal it prints "This item is not available" before exiting the program. How do I fix this?
## [6][Ruby adds experimental support for end-less method definition](https://www.reddit.com/r/ruby/comments/g9hlar/ruby_adds_experimental_support_for_endless_method/)
- url: https://www.reddit.com/r/ruby/comments/g9hlar/ruby_adds_experimental_support_for_endless_method/
---
[https://blog.saeloun.com/2020/04/27/ruby-adds-endless-method-definition-experimental](https://blog.saeloun.com/2020/04/27/ruby-adds-endless-method-definition-experimental)

&amp;#x200B;

PS: Do not shoot the messenger. I come in peace.
## [7][Ruby GUI for Desktop App](https://www.reddit.com/r/ruby/comments/g9pu38/ruby_gui_for_desktop_app/)
- url: https://www.reddit.com/r/ruby/comments/g9pu38/ruby_gui_for_desktop_app/
---
Hi guys,
I am sorry for this stupid question but... 
Does it still exist an updated library for building Desktop app GUI in Ruby? 

I tried Tk, but it gives much problem ( and cannot abolish that annoying deprecated warning messages ), shoes, fxruby and others but all of these are not updated/ bad documented... 

I am looking for a library  easy to learn for a small university Project... 

Any idea(besides change programming language) ?
## [8][Read remote file](https://www.reddit.com/r/ruby/comments/g9pk00/read_remote_file/)
- url: https://www.reddit.com/r/ruby/comments/g9pk00/read_remote_file/
---
Hey, is it possible to read remote file from other server in Ruby?  
I don't mean read from uri, but log into other server and cat file from location for example
## [9][Ruby Service object](https://www.reddit.com/r/ruby/comments/g9rzxy/ruby_service_object/)
- url: https://medium.com/@igor04/ruby-service-object-cd64dbe986aa
---

## [10][The difference between system specs and feature specs](https://www.reddit.com/r/ruby/comments/g9rnck/the_difference_between_system_specs_and_feature/)
- url: https://www.codewithjason.com/difference-system-specs-feature-specs/
---

