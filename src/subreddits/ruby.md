# ruby
## [1][Fun facts about Ruby #4](https://www.reddit.com/r/ruby/comments/h12yg6/fun_facts_about_ruby_4/)
- url: https://i.redd.it/n22dyxmq8b451.png
---

## [2][Spark Joy by Running Fewer Tests](https://www.reddit.com/r/ruby/comments/h7cclo/spark_joy_by_running_fewer_tests/)
- url: https://engineering.shopify.com/blogs/engineering/spark-joy-by-running-fewer-tests
---

## [3][RHUBARBCIPHER: A plausibly deniable multi-key file encryption tool written in Ruby for GNU/Linux and BSD systems.](https://www.reddit.com/r/ruby/comments/h7e299/rhubarbcipher_a_plausibly_deniable_multikey_file/)
- url: https://github.com/octetsplicer/RHUBARBCIPHER
---

## [4][Help me understand bubble_sort](https://www.reddit.com/r/ruby/comments/h77rzt/help_me_understand_bubble_sort/)
- url: https://www.reddit.com/r/ruby/comments/h77rzt/help_me_understand_bubble_sort/
---
https://github.com/Ubuntu19019/learnruby/blob/master/bubble_sort

The bubble sort problem is easy to implement, because it is very formulaic. I don't understand the sorted part though. I can do it, but I don't understand.

At the top sorted = false

so does while !sorted mean while sorted = true? That doesn't really make sense to me. Next it says sorted = true. Does it automatically do a pass after that. If it says sorted = true, why does it check the next loop? After it checks the next loop and makes changes sorted = false. Does that take it back to while !sorted? If so wouldn't that mean !sorted = false? If that's the case whats the point of setting sorted = false at the top?

Thanks for the help. I don't understand the logic.
## [5][Having trouble installing ruby in my Vagrant Box](https://www.reddit.com/r/ruby/comments/h777f6/having_trouble_installing_ruby_in_my_vagrant_box/)
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
## [6][HABTM to has_many through](https://www.reddit.com/r/ruby/comments/h11m7g/habtm_to_has_many_through/)
- url: https://medium.com/rubycademy/habtm-to-has-many-through-43f68f50e50e
---

## [7][Best way to design grid for flight simultor ?](https://www.reddit.com/r/ruby/comments/h13o7f/best_way_to_design_grid_for_flight_simultor/)
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
## [8][Expected output of the debug flag (`-d` or `--debug`)](https://www.reddit.com/r/ruby/comments/h0vor5/expected_output_of_the_debug_flag_d_or_debug/)
- url: https://www.reddit.com/r/ruby/comments/h0vor5/expected_output_of_the_debug_flag_d_or_debug/
---
On a fresh install of macOS 10.14.6, opening the Interactive Ruby Shell with the debug flag enabled results in the following output:

```console
$ ruby --debug
Exception `LoadError' at /System/Library/Frameworks/Ruby.framework/Versions/2.3/usr/lib/ruby/2.3.0/rubygems.rb:1242 - cannot load such file -- rubygems/defaults/operating_system
Exception `LoadError' at /System/Library/Frameworks/Ruby.framework/Versions/2.3/usr/lib/ruby/2.3.0/rubygems.rb:1251 - cannot load such file -- rubygems/defaults/ruby
Exception `LoadError' at /System/Library/Frameworks/Ruby.framework/Versions/2.3/usr/lib/ruby/2.3.0/rubygems/core_ext/kernel_require.rb:55 - cannot load such file -- did_you_mean
```

For what it's worth, similar output is shown when running a Ruby installation installed via `rbenv`.

I did not expect to see any exceptions. The closest thing I could find relating to the output I saw on GitHub was [this issue](https://github.com/rubygems/rubygems/issues/358) in the rubygems repository on GitHub. A contributor closed the issue with [this](https://github.com/rubygems/rubygems/issues/358#issuecomment-7114343) to say:

&gt; This is expected output when you run with `ruby -d`. There's nothing wrong with it.

Why is it the expected output?

I also came across [this Stack Overflow post](https://stackoverflow.com/questions/35240324/something-wrong-with-my-ruby?noredirect=1#comment58199906_35240324). To quote [the accepted answer](https://stackoverflow.com/a/35240886/11805749):

&gt; The system ruby installation seems to be messed up. Try removing the offending `/Library/Ruby/Site/2.0.0/rubygems.rb` (and may be more of enclosing folders)
&gt; 
&gt; Ruby have its bundled rubygems at `/System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/lib/ruby/2.0.0` so it should work with these.
&gt; 
&gt; Once ruby is alive - you can try upgrading rubygems again by `sudo gem update --system` if needed, but I advise setting up a ruby version manager like `rvm` (rvm.io) and leave the system ruby be there only for emergences and backing the `brew`.
## [9][Help me to understand the triple equal with module](https://www.reddit.com/r/ruby/comments/h0s4gx/help_me_to_understand_the_triple_equal_with_module/)
- url: https://www.reddit.com/r/ruby/comments/h0s4gx/help_me_to_understand_the_triple_equal_with_module/
---
&amp;#x200B;

https://preview.redd.it/qyjc4bz6m7451.png?width=1058&amp;format=png&amp;auto=webp&amp;s=afbcc7dc64239ae644fee353e0a7171b195f9c95

Why `@controller.class.parent === Recall` doesn't returns `true`?

The `object_id` are the same but it is not equal. Why?

`Recall` is a module.
## [10][Ruby Desktop Development - Complete OSS Mac DMG App](https://www.reddit.com/r/ruby/comments/h0lnai/ruby_desktop_development_complete_oss_mac_dmg_app/)
- url: https://www.reddit.com/r/ruby/comments/h0lnai/ruby_desktop_development_complete_oss_mac_dmg_app/
---
A Ruby desktop complete application developed with [Glimmer](https://github.com/AndyObtiva/glimmer) has been released at version 2.1 (open-source and free):

[https://github.com/AndyObtiva/MathBowling](https://github.com/AndyObtiva/MathBowling)

Please use as a learning reference for how to leverage the [Glimmer](https://github.com/AndyObtiva/glimmer) desktop development GUI library.

Enjoy!!!
