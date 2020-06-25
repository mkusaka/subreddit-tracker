# ruby
## [1][HEY’s Gemfile](https://www.reddit.com/r/ruby/comments/hfhrcb/heys_gemfile/)
- url: https://gist.github.com/dhh/782fb925b57450da28c1e15656779556
---

## [2][Exception tracking: Airbrake vs Honeybadger vs Sentry vs AppSignal](https://www.reddit.com/r/ruby/comments/hfk47b/exception_tracking_airbrake_vs_honeybadger_vs/)
- url: https://www.reddit.com/r/ruby/comments/hfk47b/exception_tracking_airbrake_vs_honeybadger_vs/
---
Which one do you use / have recently decided for? And why?

What do you hate about the current ones that you (have to) use?

Did I miss an important one for the ruby ecosystem?
## [3][Lingohub - removing unused translation keys](https://www.reddit.com/r/ruby/comments/hfk697/lingohub_removing_unused_translation_keys/)
- url: https://www.reddit.com/r/ruby/comments/hfk697/lingohub_removing_unused_translation_keys/
---
Is there any way to remove unused keys on Lingohub?

I'm just deleting a bunch of unused keys in the translation yml file. Is there any good and safe way to do this?

edit:

Answer from Lingohub team:

&gt;Hi xyz,LingoHub will automatically deactivate keys when they are removed from a file. So for example you had a file with the following keys:A B CAnd you then upload the same file with the keysA CThe B key would be deactivated and not shown in the editor or exported anymore.Please note, this behavior is changeable through project settings. So make sure you check them before.Hope that helps.Best,

This is a partial answer that I wanted. Can I somehow reverse these keys if the mistake is created?or create more environments?

stackoverflow q: [https://stackoverflow.com/questions/62556336/lingohub-removing-unused-translation-keys](https://stackoverflow.com/questions/62556336/lingohub-removing-unused-translation-keys)

  
Lingohub service for translates: [https://lingohub.com/](https://lingohub.com/)
## [4][Moving from EventMachine to Async](https://www.reddit.com/r/ruby/comments/hfgsdf/moving_from_eventmachine_to_async/)
- url: https://blog.joshsoftware.com/2020/06/19/moving-from-eventmachine-to-async/
---

## [5][til-rb, a gem to help you write and manage a TIL repo](https://www.reddit.com/r/ruby/comments/hflzrs/tilrb_a_gem_to_help_you_write_and_manage_a_til/)
- url: https://www.reddit.com/r/ruby/comments/hflzrs/tilrb_a_gem_to_help_you_write_and_manage_a_til/
---
Hey there 👋

&amp;#x200B;

I recently published an early version of [a small gem](https://github.com/pjambet/til-rb/) I've been working on to manage a "TIL repo", like [this one](https://github.com/jbranchaud/til) that made the rounds recently. It essentially creates a new commit for you, using the GH API, with a new file containing the content of the new TIL, and maintains a table of content in the [README.md](https://README.md) file.

&amp;#x200B;

Building it was really interesting and [I wrote a post about it](https://blog.pjam.me/posts/til-cli/), thought it would be worth sharing here!

&amp;#x200B;

I've been using it for a week or so, and while I'm sure it has a lot of bugs (almost total lack of validation on user input for instance), it has worked reasonably well so far.

&amp;#x200B;

Let me know if you questions or feedback.
## [6][Introducing UndercoverCI - a GitHub code review app that finds untested methods you've changed](https://www.reddit.com/r/ruby/comments/hf8070/introducing_undercoverci_a_github_code_review_app/)
- url: https://medium.com/@mrgrodo/introducing-undercoverci-98c6c50793f2
---

## [7][Fun facts about Ruby #5](https://www.reddit.com/r/ruby/comments/hf04se/fun_facts_about_ruby_5/)
- url: https://i.redd.it/3tbszye9qu651.png
---

## [8][Ruby in Vagrant setup HELP!!](https://www.reddit.com/r/ruby/comments/hfd83f/ruby_in_vagrant_setup_help/)
- url: https://www.reddit.com/r/ruby/comments/hfd83f/ruby_in_vagrant_setup_help/
---
Hi there. I've been having this issue a while now. I graduated bootcamp and they used vagrant for teaching Rails development. That's what I'm used to. But I wanted to try to make my own Vagrant box to start my personal development environment instead of using the one the bootcamp provided. And the first project is going to be my Web Dev portfolio in Ruby on Rails.

Well now I'm running into issues installing the latest ruby version (or even the recommented one if you don't know what your'e doing which is 2.6.6 with devkit. (Not sure how I'd go about manually installing ruby with devkit in ubuntu from the terminal.)

I'm at the point where I'm ssh'd into my Vagrant box, and I'm trying to run the command rvm install 2.7.1 and it's throwing an error at me. I'm not sure how to determine exactly what's wrong, but here's the output after that command:

    vagrant@ubuntu-bionic:~$ rvm install 2.7.1
    Searching for binary rubies, this might take some time.
    Found remote file https://rvm_io.global.ssl.fastly.net/binaries/ubuntu/18.04/x86_64/ruby-2.7.1.tar.bz2
    Checking requirements for ubuntu.
    Requirements installation successful.
    ruby-2.7.1 - #configure
    ruby-2.7.1 - #download
    Downloaded archive checksum did not match!
    ruby-2.7.1 - #validate archive
    
    bzip2: Data integrity error when decompressing.
            Input file = (stdin), output file = (stdout)
    
    It is possible that the compressed file(s) have become corrupted.
    You can use the -tvv option to test integrity of such files.
    
    You can use the `bzip2recover' program to attempt to recover
    data from undamaged sections of corrupted files.
    
    tar: Child returned status 2
    tar: Error is not recoverable: exiting now
    
    bzip2: Data integrity error when decompressing.
            Input file = (stdin), output file = (stdout)
    
    It is possible that the compressed file(s) have become corrupted.
    You can use the -tvv option to test integrity of such files.
    
    You can use the `bzip2recover' program to attempt to recover
    data from undamaged sections of corrupted files.
    
    tar: Child returned status 2
    tar: Error is not recoverable: exiting now
    
    bzip2: Data integrity error when decompressing.
            Input file = (stdin), output file = (stdout)
    
    It is possible that the compressed file(s) have become corrupted.
    You can use the -tvv option to test integrity of such files.
    
    You can use the `bzip2recover' program to attempt to recover
    data from undamaged sections of corrupted files.
    
    tar: Child returned status 2
    tar: Error is not recoverable: exiting now
    The downloaded package for https://rvm_io.global.ssl.fastly.net/binaries/ubuntu/18.04/x86_64/ruby-2.7.1.tar.bz2,
    Does not contains single 'bin/ruby' or 'ruby-2.7.1',
    Only '' were found instead.
    Mounting remote ruby failed with status 4, trying to compile.
    Checking requirements for ubuntu.
    Requirements installation successful.
    Installing Ruby from source to: /home/vagrant/.rvm/rubies/ruby-2.7.1, this may take a while depending on your cpu(s)...
    ruby-2.7.1 - #downloading ruby-2.7.1, this may take a while depending on your connection...
    Downloaded archive checksum did not match!
    ruby-2.7.1 - #extracting ruby-2.7.1 to /home/vagrant/.rvm/src/ruby-2.7.1......
    Error running '__rvm_package_extract /home/vagrant/.rvm/archives/ruby-2.7.1.tar.bz2 /home/vagrant/.rvm/tmp/rvm_src_2151',
    please read /home/vagrant/.rvm/log/1593047102_ruby-2.7.1/extract.log
    There has been an error while trying to extract the source. Halting the installation.
    There has been an error fetching the ruby interpreter. Halting the installation.
    vagrant@ubuntu-bionic:~$

Here is my Vagrantfile, also. Pretty basic:

    Vagrant.configure("2") do |config|
    config.vm.box = "ubuntu/bionic64" config.vm.network "forwarded_port", guest: 3000, host: 3000
     config.vm.provision "shell", privileged: false, inline: &lt;&lt;-SHELL
       echo "==&gt; Installing RVM..."
       # https://rvm.io/rvm/install
       curl -sSL https://rvm.io/mpapis.asc | gpg --import -
       curl -sSL https://rvm.io/pkuczynski.asc | gpg --import -
       curl -sSL https://get.rvm.io | bash -s stable
       source "$HOME/.rvm/scripts/rvm"
       rvm install 2.6.6
       gem install rails
     SHELL

config.vm.synced\_folder ".", "/vagrant\_files" config.vm.provider "virtualbox" do |vb| vb.memory = "4096" end config.vm.provider "virtualbox" do |vb| vb.cpus = 8 end end

Any help will be greatly appreciated. I've been trying to get help on this and this is the furthest so far I think I've gotten besides this - I did get ruby and rails to install by running sudo apt-install ruby, which gives me a really old version, and then the same with Rails when I try to install that the same way, it gives me version 4.something when it's pas 6.0.3!

If I could just get a newer version of ruby working and I can install rails, then I'm set!
## [9][UnluckyWretchedTrust](https://www.reddit.com/r/ruby/comments/hf64x5/unluckywretchedtrust/)
- url: https://www.reddit.com/r/ruby/comments/hf64x5/unluckywretchedtrust/
---
So im learning ruby and I found this strange file on my second monitor on the desktop all alone called [UnluckyWretchedTrust.zip](https://UnluckyWretchedTrust.zip). I scanned it for virus's and came back safe. So I unzipped and it has a main.rb file inside. it was created at 4am. So i open it in atom and it is the code I wrote for a blackjack game. Why was it named UnluckyWretchedTrust? That seems odd to me.
## [10][Hey, guys. I recently did a task for interview, and as usual I didn't get answer about the code, can u critic my code pls?](https://www.reddit.com/r/ruby/comments/hey2oq/hey_guys_i_recently_did_a_task_for_interview_and/)
- url: https://www.reddit.com/r/ruby/comments/hey2oq/hey_guys_i_recently_did_a_task_for_interview_and/
---
[https://gitlab.com/abdullaachilov/my\_test\_task](https://gitlab.com/abdullaachilov/my_test_task)

Here is the link, Id love to know - how I can improve myself.  


UPD: The task was simple - create API for product with tags (there are was few requests in postman file)  
The condition also was - it will be work with ember on front-end, so I have some formats restricted.
