# ruby
## [1][Good (can be paid) courses for learning Ruby + Rails for PHP developer](https://www.reddit.com/r/ruby/comments/h7p6tm/good_can_be_paid_courses_for_learning_ruby_rails/)
- url: https://www.reddit.com/r/ruby/comments/h7p6tm/good_can_be_paid_courses_for_learning_ruby_rails/
---
Hello! I am a PHP developer with few years experience and I want to switch to ruby and get a job as a junior ROR dev. I was hoping you could point me to some good and tested tutorials or courses where the focus is put on creating a real-life application.
## [2][Is my algorithm's time complexity O(n log n)?](https://www.reddit.com/r/ruby/comments/h7s4yg/is_my_algorithms_time_complexity_on_log_n/)
- url: https://www.reddit.com/r/ruby/comments/h7s4yg/is_my_algorithms_time_complexity_on_log_n/
---
    def two_sum?(array, value)
    	array.sort! # O(nlogn)
    	array.each do |element|
    		return true if bsearch(array - [element], value - element) == true
    	end
    	return false
    end
    
    def bsearch(array, value)
    	return false if array.empty?
    	mid_idx = array.length / 2
    	mid_value = array[mid_idx]
    	return true if mid_value == value
    	mid_value &lt; value ? bsearch(array[0...mid_idx], value) : bsearch(array[mid_idx+1..-1], value)
    end

I'm trying to create a function that finds a two unique numbers in an array whose sum equals the value in the second argument. I believe my implementation has a time complexity of O(n log n). However, when I run it with another function whose time complexity is also O(n log n), the total time is way different (calculated using the Benchmark gem) using the same input. For my function, it takes about 0.9 seconds. For the other function, it is taking 0.003 seconds. Is there any error in my algorithm analysis? Is my implementation not O(n log n)?
## [3][Reverting to an older string time format](https://www.reddit.com/r/ruby/comments/h7pq7v/reverting_to_an_older_string_time_format/)
- url: https://www.reddit.com/r/ruby/comments/h7pq7v/reverting_to_an_older_string_time_format/
---
Currently with ruby 2.7.1, when you do something like `Date.today.to_s(:long)` you'll get something like

`June 09, 2020` 

But with older versions of ruby (for instance 2.2.4), that format would look like this: 

`June 9, 2020`

Without that 0 padding on the day. Is there anyway to revert back to this old formatting, i.e **without** the 0 padding on the day? Without changing the `to_s(:long)` part and without going back to an older  version of ruby? 

Or is there anyway to specify that `:long` should follow this structure: `%B %e, %Y`

All help is appreciated!
## [4][Spark Joy by Running Fewer Tests](https://www.reddit.com/r/ruby/comments/h7cclo/spark_joy_by_running_fewer_tests/)
- url: https://engineering.shopify.com/blogs/engineering/spark-joy-by-running-fewer-tests
---

## [5][Fun facts about Ruby #4](https://www.reddit.com/r/ruby/comments/h12yg6/fun_facts_about_ruby_4/)
- url: https://i.redd.it/n22dyxmq8b451.png
---

## [6][RHUBARBCIPHER: A plausibly deniable multi-key file encryption tool written in Ruby for GNU/Linux and BSD systems.](https://www.reddit.com/r/ruby/comments/h7e299/rhubarbcipher_a_plausibly_deniable_multikey_file/)
- url: https://github.com/octetsplicer/RHUBARBCIPHER
---

## [7][Help me understand bubble_sort](https://www.reddit.com/r/ruby/comments/h77rzt/help_me_understand_bubble_sort/)
- url: https://www.reddit.com/r/ruby/comments/h77rzt/help_me_understand_bubble_sort/
---
https://github.com/Ubuntu19019/learnruby/blob/master/bubble_sort

The bubble sort problem is easy to implement, because it is very formulaic. I don't understand the sorted part though. I can do it, but I don't understand.

At the top sorted = false

so does while !sorted mean while sorted = true? That doesn't really make sense to me. Next it says sorted = true. Does it automatically do a pass after that. If it says sorted = true, why does it check the next loop? After it checks the next loop and makes changes sorted = false. Does that take it back to while !sorted? If so wouldn't that mean !sorted = false? If that's the case whats the point of setting sorted = false at the top?

Thanks for the help. I don't understand the logic.
## [8][Having trouble installing ruby in my Vagrant Box](https://www.reddit.com/r/ruby/comments/h777f6/having_trouble_installing_ruby_in_my_vagrant_box/)
- url: https://www.reddit.com/r/ruby/comments/h777f6/having_trouble_installing_ruby_in_my_vagrant_box/
---
Hello. I'm in desperate need of help with this. So I haven't done much on my own outside of my coding bootcamp I graduated from yet. They had me download a set of files that included everything set up for Vagrant. But I wanted to recreate the environment for myself so that everything came from me and doesn't have any files related to the bootcamp anymore.

I'm going through the tutorial of just creating the folder you want your code environment in, running the vagrant init command int he terminal and creating everything that way. I'm at the point of installing RVM and Ruby, specifically 2.7.1 which is the latest as of the time of this post. I'm getting an error though.

Here's the output when I run "rvm install ruby-2.7.1":

Searching for binary rubies, this might take some time.

    Found remote file https://rvm_io.global.ssl.fastly.net/binaries/ubuntu/18.04/x86_64/ruby-2.7.1.tar.bz2
    Checking requirements for ubuntu.
    Requirements installation successful.
    ruby-2.7.1 - #configure
    ruby-2.7.1 - #download
    Archive bin-ruby-2.7.1.tar.bz2 checksum did not match, downloading again.
      % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                     Dload  Upload   Total   Spent    Left  Speed
    100 21.0M  100 21.0M    0     0  9572k      0  0:00:02  0:00:02 --:--:-- 9572k
    Downloaded archive checksum did not match, archive was removed!
    If you wish to continue with not matching download add '--verify-downloads 2' after the command.
    
    Downloading https://rvm_io.global.ssl.fastly.net/binaries/ubuntu/18.04/x86_64/ruby-2.7.1.tar.bz2 failed.
    Mounting remote ruby failed with status 2, trying to compile.
    Checking requirements for ubuntu.
    Requirements installation successful.
    Installing Ruby from source to: /usr/share/rvm/rubies/ruby-2.7.1, this may take a while depending on your cpu(s)...
    ruby-2.7.1 - #downloading ruby-2.7.1, this may take a while depending on your connection...
    Archive ruby-2.7.1.tar.bz2 checksum did not match, downloading again.
      % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                     Dload  Upload   Total   Spent    Left  Speed
      0 14.0M    0 58100    0     0   267k      0  0:00:53 --:--:--  0:00:53  266k
    curl: (16) Error in the HTTP2 framing layer
    There was an error(16).
    Checking fallback: https://ftp.ruby-lang.org/pub/ruby/2.7/ruby-2.7.1.tar.bz2
    ** Resuming transfer from byte position 58100
      % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                     Dload  Upload   Total   Spent    Left  Speed
    100 13.9M  100 13.9M    0     0   9.9M      0  0:00:01  0:00:01 --:--:--  9.9M
    Downloaded archive checksum did not match, archive was removed!
    If you wish to continue with not matching download add '--verify-downloads 2' after the command.
    
    There has been an error fetching the ruby interpreter. Halting the installation.

I just need to try and get this fixed so I can continue to work on my Web Dev portfolio which is based on Ruby on Rails. I'm just not understanding why I'm now having all this trouble with my development environment when using the files that the bootcamp gave me is fine. It's just that those install older versions of ruby and rails and stuff and I want the latest working versions of everything. Any help is greatly appreciated.
## [9][HABTM to has_many through](https://www.reddit.com/r/ruby/comments/h11m7g/habtm_to_has_many_through/)
- url: https://medium.com/rubycademy/habtm-to-has-many-through-43f68f50e50e
---

## [10][Best way to design grid for flight simultor ?](https://www.reddit.com/r/ruby/comments/h13o7f/best_way_to_design_grid_for_flight_simultor/)
- url: https://www.reddit.com/r/ruby/comments/h13o7f/best_way_to_design_grid_for_flight_simultor/
---
The general idea is that:

\- Plane will fly at regulating speeds, altitude, and grid space as long as nothing obstructs its path. With "blockers" as AI  bots appearing at through out the path. 

OBJECTS 

plane \\\_\_\_/

blockers |-|  (yeah i wanted it to look like a tie fighter"

THE ISSUE

should the board grid be ordered from 0-12 across? or start from the bottom 0 - 12 ? On a smaller scale what im saying is what i more practicle

model A

    def display
        puts " #{@space[0]} | #{@space[1]} | #{@space[2]} "
        puts "-----------"
        puts " #{@space[3]} | #{@space[4]} | #{@space[5]} "
        puts "-----------"
        puts " #{@space[6]} | #{@space[7]} | #{@space[8]} "
      end

or 

model B

    def display
        puts " #{@space[6]} | #{@space[7]} | #{@space[8]} "
        puts "-----------"
        puts " #{@space[3]} | #{@space[4]} | #{@space[5]} "
        puts "-----------"
        puts " #{@space[0]} | #{@space[1]} | #{@space[2]} "
      end

to me, model B makes more sense visually. but being fairly new to programming. What are your thoughts ?

obviously the actualy grid is 12 x 12 not 3 x 3
