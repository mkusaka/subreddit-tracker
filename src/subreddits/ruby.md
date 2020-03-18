# ruby
## [1][Lies, Damned Lies, and Averages: Perc50, Perc95 explained for Programmers](https://www.reddit.com/r/ruby/comments/fkaja8/lies_damned_lies_and_averages_perc50_perc95/)
- url: https://schneems.com/2020/03/17/lies-damned-lies-and-averages-perc50-perc95-explained-for-programmers/
---

## [2][Jason Gedge, Staff Production Engineer at Shopify - The Rails with Jason Podcast](https://www.reddit.com/r/ruby/comments/fk6rbh/jason_gedge_staff_production_engineer_at_shopify/)
- url: https://www.codewithjason.com/rails-with-jason-podcast/jason-gedge/
---

## [3][Spring Cleaning: Tidying up your Rails codebase](https://www.reddit.com/r/ruby/comments/fk5vco/spring_cleaning_tidying_up_your_rails_codebase/)
- url: https://boringrails.com/articles/spring-cleaning/
---

## [4][Understanding delegate in Ruby on Rails](https://www.reddit.com/r/ruby/comments/fk1fte/understanding_delegate_in_ruby_on_rails/)
- url: https://medium.com/@merajulislam/understanding-delegate-in-ruby-on-rails-i-wish-i-knew-before-5edd341bad47?source=friends_link&amp;sk=c5905885f140221f5c480d264ae2152d
---

## [5][Clean up downloads folder(My 3rd month of studying ruby)](https://www.reddit.com/r/ruby/comments/fk3x3a/clean_up_downloads_foldermy_3rd_month_of_studying/)
- url: https://www.reddit.com/r/ruby/comments/fk3x3a/clean_up_downloads_foldermy_3rd_month_of_studying/
---
After 3 month studying ruby 4 hours in week I finally made something useful.

[https://github.com/brian7346/Clean-downloads/blob/master/cleanup.rb](https://github.com/brian7346/Clean-downloads/blob/master/cleanup.rb)

Here's my program that clean my "Downloads" folder. It takes about 2-3 hours to make it.

https://preview.redd.it/jjr3bpoi18n41.png?width=324&amp;format=png&amp;auto=webp&amp;s=37c8389fe3e4fef8ba6c8bc1be75d9149ae3eb5c

https://preview.redd.it/vauytioi18n41.png?width=320&amp;format=png&amp;auto=webp&amp;s=60e4fdc3921efed4d309d4dd4c53885e7ad2d5dc
## [6][Corona Helper](https://www.reddit.com/r/ruby/comments/fkdom8/corona_helper/)
- url: https://www.reddit.com/r/ruby/comments/fkdom8/corona_helper/
---
Hello all,

Lots of people are having problems due to the corona virus. There are lots of people who are willing to help. I was hoping that something could be done to connect them efficiently so we can all get through Corona. 

Here is my idea for an app: 

	* On one side a user can post their need (such as: I am out of pasta/nappies and there are none in the store. Could someone check on my grandma, I have a burst pipe but there is no one around to help) along with their location.
	* On the other side people who are willing to help can post their location and info such if they have a car or extra supplies
	* Both users will post if they have been exposed to Corona or not. (this could go on for 12-18 months so people can show if they already have immunity etc)
	* The users can message each other.
	* This can be implemented worldwide using post codes/zip codes etc.
	* Some things like safeguarding and protecting vulnerable people from exploitation will need to be worked out - but I'm sure solutions can be found.



To build this we will need:
	* Backend Developers (Python/Ruby etc)
	* Frontend (Javascript)
	* IOS developers
	* Android
	* UX
	* DevOps
	* Project Management
	* Anyone else with skills they can bring to the table.



I've set up a discord server https://discord.gg/YQgZvfr - please join if you have some time to spare and would like to help.

(if you could upvote and comment on this post to keep it bumped it would be most apprecaited).
## [7][Unable to rename file...](https://www.reddit.com/r/ruby/comments/fk97is/unable_to_rename_file/)
- url: https://www.reddit.com/r/ruby/comments/fk97is/unable_to_rename_file/
---
f = [File.open](https://File.open)("test.txt" , 'w+')

&amp;#x200B;

    f.write("hello ruby")
    
    f.close
    
    
    
    f =  File.open("test.txt" , 'a+')
    
    
    
    f.write(". happy bday")
    
    	
    
    f.close
    
    
    
    f = File.open("test.txt" , 'a+')
    
    
    
    content = f.sysread(7)
    
    		
    
    puts content
    
    		
    
    f.syswrite("fresh")	
    
    f = File.open("test.txt" , 'r')			
    
    print f.read
    f.close
    
    File.rename("test.txt" , "new.txt")
    
    f = File.open("new.txt" , r)
    
    print f.read

when I try to rename the file , the compiler throws an error-

in \`rename': Permission denied @ rb\_file\_s\_rename - (test.txt, new.txt) (Errno::EACCES)

What can be the possible solution?
## [8][For everyone (yes, both of you) who's planning to make a Telegram bot with Ruby](https://www.reddit.com/r/ruby/comments/fjqpv0/for_everyone_yes_both_of_you_whos_planning_to/)
- url: https://www.reddit.com/r/ruby/comments/fjqpv0/for_everyone_yes_both_of_you_whos_planning_to/
---
&amp;#x200B;

[Example of how this is useful](https://preview.redd.it/w4f8klc533n41.png?width=1051&amp;format=png&amp;auto=webp&amp;s=a612f6e7f02fa1d0a8e8aa4afab1418fe1ba3e2f)

I've been learning Ruby for several months now, and making my own Telegram bot with Ruby (using [telegram-bot gem](https://github.com/atipugin/telegram-bot-ruby)). These are helpers I've built during that time to help myself with that. I'm not a professional developer (yet :3), so the choices in the structure and the code itself might seem questionable. I don't feel like building a full-fledged gem with it, so I'll just leave a link:  [https://github.com/bulgakke/telegram\_bot\_helpers](https://github.com/bulgakke/telegram_bot_helpers)

Even if you feel you don't need it and you can do it better than me, or you can't bother reading someone else's home project, just take a look and maybe you'll better understand how to encapsulate stuff while working with Telegram Bot API
## [9][Remote work for Ruby?](https://www.reddit.com/r/ruby/comments/fjmo4d/remote_work_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fjmo4d/remote_work_for_ruby/
---
Does anyone know where I can get remote work for ruby. Ideally small jobs that would pay about $30-50
## [10][gets.chomp waits for input before rest of the code runs](https://www.reddit.com/r/ruby/comments/fjud5z/getschomp_waits_for_input_before_rest_of_the_code/)
- url: https://www.reddit.com/r/ruby/comments/fjud5z/getschomp_waits_for_input_before_rest_of_the_code/
---
Is this a common issue? I have a gets.chomp towards the end but it won’t run the preceding code (most puts statements) until I enter a value for the gets statements. Is this a normal problem to run into?
