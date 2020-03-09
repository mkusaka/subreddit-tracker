# ruby
## [1][Ruby 2.7 - Trying out without blowing up my dev.](https://www.reddit.com/r/ruby/comments/ffqy9h/ruby_27_trying_out_without_blowing_up_my_dev/)
- url: https://www.reddit.com/r/ruby/comments/ffqy9h/ruby_27_trying_out_without_blowing_up_my_dev/
---
Hello,

I'd like to start trying out the newest Ruby version, but I'd want to do so without breaking my current dev environment for work. At present I use *rbenv*. What would be the safest way to use 2.6 and 2.7 concurrently without issues?
## [2][Installing older version of ruby with RVM?](https://www.reddit.com/r/ruby/comments/ffpdcn/installing_older_version_of_ruby_with_rvm/)
- url: https://www.reddit.com/r/ruby/comments/ffpdcn/installing_older_version_of_ruby_with_rvm/
---
I factory reset my mac to the current version of Mac OS. I installed RVM because that's what I normally use. However, for work, we still have projects that run on ruby 2.2.4, so I tried installing it but it didn't work. New versions install correctly, however, such as 2.6.3.

When I try installing 2.2.4, I get this error:

    Error running 'env GEM_HOME=/Users/garcia/.rvm/gems/ruby-2.2.4@global GEM_PATH= /Users/garcia/.rvm/rubies/ruby-2.2.4/bin/ruby -d /Users/garcia/.rvm/src/rubygems-3.0.8/setup.rb --no-document',
    please read /Users/garcia/.rvm/log/1583728534_ruby-2.2.4/rubygems.install.log

However, I'm still able to see 2.2.4 when I run `rvm list`. When I switch to it via `rvm use --default 2.2.4`, everything seems to work, but when I run `gem update --system`. I get this error:

    ERROR:  Loading command: update (LoadError)
    	cannot load such file -- openssl
    ERROR:  While executing gem ... (NoMethodError)
        undefined method `invoke_with_build_args' for nil:NilClass

&amp;#x200B;

Any ideas? I submitted an issue ticket to the RVM Github repo like a week ago, but no one has responded.
## [3][Top 5 Reasons For Outsourcing Front-End Development](https://www.reddit.com/r/ruby/comments/ffucix/top_5_reasons_for_outsourcing_frontend_development/)
- url: http://selleoo.xyz
---

## [4][Ruby 2.2.3 in 2020](https://www.reddit.com/r/ruby/comments/ffscd5/ruby_223_in_2020/)
- url: https://www.reddit.com/r/ruby/comments/ffscd5/ruby_223_in_2020/
---
Ruby 2.2.3 has been dead for a while and yet I've been asked to look at a legacy project in order to update and fix it.  
I can't get 2.2.3 to install in order to get a local env up and running on either MacOS or Ubuntu. 

The error seems to be down to an openssl rubygems issue: 

from `rubygems.install.log`

    Exception `LoadError' at /Users/user/.rvm/rubies/ruby-2.2.3/lib/ruby/2.2.0/rubygems.rb:1222 - cannot load such file -- rubygems/defaults/operating_system
    Exception `LoadError' at /Users/user/.rvm/rubies/ruby-2.2.3/lib/ruby/2.2.0/rubygems.rb:1231 - cannot load such file -- rubygems/defaults/ruby
    /Users/user/.rvm/src/rubygems-3.0.8/lib/rubygems/core_ext/kernel_require.rb:54:in `require': cannot load such file -- openssl (LoadError)

Does anyone have any suggestions on how proceed?
## [5][opmlparser gem - read / parse outlines (incl. feed subscription lists) in the OPML (Outline Processor Markup Language) format in xml](https://www.reddit.com/r/ruby/comments/ffs6lk/opmlparser_gem_read_parse_outlines_incl_feed/)
- url: https://github.com/feedreader/pluto/tree/master/opmlparser
---

## [6][Splitting an array into sub arrays](https://www.reddit.com/r/ruby/comments/ffgx8h/splitting_an_array_into_sub_arrays/)
- url: https://www.reddit.com/r/ruby/comments/ffgx8h/splitting_an_array_into_sub_arrays/
---
Hi guys, 

I am working on my final for my Ruby class.  I have an array that I have created from reading in a file.  I am creating a word search puzzle.  I need to list the words under my word search grid.  I have 45 words and want 3 rows of 15 words each.  I am also using a Prawn to create a pdf file to output my word search grid, words, and a grid key.  Any advice on the best way to split my array into 3 equal columns to print out my words?  

   I was thinking about creating some kind of method that would calculate the length, the take that length and divide it by 3.  From there I would print a string with one column ljust, one column centered, and one column rjust.  

   The other option I am entertaining is to create a boundary with the prawn and somehow break the list into 3 parts and put each part in the created boundary.  

   Any ideas and direction here would be greatly appreciated.  

&amp;#x200B;

Here is my method to create my array  (I have to put a space after my instance variable so it would allow me to include it on this post)

@ words = \[\]  
*file* = File.open(*file*)  
until *file*.eof?  
 *word* = *file*.gets.chomp  
@ words&lt;&lt; *word*.delete(" ")  
@ words = @ words.sort\_by{ |*x*|-*x*.length }  


end  
*file*.close

&amp;#x200B;

My list of words I am including

Apple  
Apricot  
Avocado  
Breadfruit  
Banana  
Blackberry  
Blackcurrant  
Blueberry  
Cherimoya  
Cherry  
Clementine  
Coconut  
Cranberry  
Custard Apple  
Durian  
Fig  
Grapefruit  
Grape  
Guava  
Jackfruit  
Kiwi  
Lemon  
Lime  
Loganberry  
Mandarin  
Mango  
Mangosteen  
Melon  
Nectarine  
Orange  
Papaya  
Peach  
Pear  
Persimmon  
Pineapple  
Plum  
Pomegranate  
Quince  
Satsuma  
Sharon  
Strawberry  
Tamarillo  
Tangerine  
Ugli  
Watermelon  


My prawn pdf creation

 def create\_pdf  
 Prawn::Document.generate('puzzle2.pdf') do |*pdf*|  
 *pdf*.font 'Courier', size: 20  
 *pdf*.text 'Word Search', align: :center  
 *pdf*.font 'Courier', size: 8  
 *pdf*.move\_down(80)  


*pdf*.text(print\_puzzle, align: :center)  
 *pdf*.start\_new\_page  
 *pdf*.text(print\_puzzle, align: :center)  
 end  
  end  
end
## [7][Started my first roguelike with Dragon Ruby](https://www.reddit.com/r/ruby/comments/fexv07/started_my_first_roguelike_with_dragon_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fexv07/started_my_first_roguelike_with_dragon_ruby/
---
I decided to try out Dragon Ruby while also learning to develop my first roguelike. I was able to finish generating a basic dungeon map using the assets from [https://0x72.itch.io/dungeontileset-ii](https://0x72.itch.io/dungeontileset-ii) over the last few days. So far it's been pretty fun and I've already learned a lot about the Ruby programming language and 2D game development in general.

https://i.redd.it/bo2cbv0h7al41.gif
## [8][Improving my arrays](https://www.reddit.com/r/ruby/comments/ff3rqu/improving_my_arrays/)
- url: https://www.reddit.com/r/ruby/comments/ff3rqu/improving_my_arrays/
---
For a friend's game in Ruby 1.8 I'm coding something and while visually I've got it under control, I feel like I'm not organizing my data properly:

The idea is that the player can unlock several page entries in a book, and then, within each entry, there are nine other pieces of information to unlock. The thing is, I don't think I'm thinking right when it comes to how should this information be unlocked.

For example, an array should be all false at first, or should it? I need to initialize the data, and right now, I'm doing it like this:

[https://i.imgur.com/nR3eziA.png](https://i.imgur.com/nR3eziA.png)

I hope someone can help me figure out a better way of arranging this example $unlockedEntries array, as to store the data. 

https://preview.redd.it/hpr80kex6cl41.png?width=739&amp;format=png&amp;auto=webp&amp;s=167710e831f6c14956e7110771a1b3b3e1fd59ed
## [9][Learning Ruby without webdev?](https://www.reddit.com/r/ruby/comments/fer4mz/learning_ruby_without_webdev/)
- url: https://www.reddit.com/r/ruby/comments/fer4mz/learning_ruby_without_webdev/
---
Ruby seems pretty huge for web development type stuff, but I've also heard of its scripting capabilities applicable to ops and security, which interest me as I've come off knowing some rudimentary programming knowledge in Java and Python, not yet having JS knowledge(HTML/CSS of course). 

Could someone really get a feel for the Ruby language, hence "learning" it to a degree, without stepping much into web development?
## [10][How to parse nested json in ruby?](https://www.reddit.com/r/ruby/comments/feg2a1/how_to_parse_nested_json_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/feg2a1/how_to_parse_nested_json_in_ruby/
---
Hey guys,

I'm trying to use an API that returns nested json. I'm using json.parse but it is splitting the json file into two large hashes, when I want to go down one or two more levels. Does anyone know how I could do this? Thanks.
