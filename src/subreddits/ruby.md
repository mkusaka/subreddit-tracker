# ruby
## [1][Why the Release of Ruby 3 Will Be Monumental](https://www.reddit.com/r/ruby/comments/jtfzdp/why_the_release_of_ruby_3_will_be_monumental/)
- url: https://www.ruby3.dev/the-art-of-code/2020/11/12/ruby-3-monumental/
---

## [2][DragonRuby Game Toolkit Sound Synthesis in Pure Ruby ^_^](https://www.reddit.com/r/ruby/comments/jtbtkb/dragonruby_game_toolkit_sound_synthesis_in_pure/)
- url: https://www.youtube.com/watch?v=zEzovM5jT-k&amp;feature=youtu.be&amp;ab_channel=AmirRajan
---

## [3][Deep Dive: Moving ruby projects from Travis to Github Actions for CI](https://www.reddit.com/r/ruby/comments/jt2uub/deep_dive_moving_ruby_projects_from_travis_to/)
- url: https://bibwild.wordpress.com/2020/11/12/deep-dive-moving-ruby-projects-from-travis-to-github-actions-for-ci/
---

## [4][footballdata-12xpert gem - download, convert &amp; import 22+ top football leagues from 25 seasons back to 1993/94 from Joseph Buchdahl (12Xpert)'s Football Data website](https://www.reddit.com/r/ruby/comments/jtee59/footballdata12xpert_gem_download_convert_import/)
- url: https://github.com/sportdb/sport.db.sources/tree/master/footballdata-12xpert
---

## [5][Custom exception on mailers deliver_later - question](https://www.reddit.com/r/ruby/comments/jtebo4/custom_exception_on_mailers_deliver_later_question/)
- url: https://www.reddit.com/r/ruby/comments/jtebo4/custom_exception_on_mailers_deliver_later_question/
---
Let's say we have a simple Mailer as:

    class PostmanMailer &lt; ApplicationMailer   
        rescue_from CustomError do |exception|
        ... do something ...
        end
         def invitation(user)
            mail(to: user.email,subject: user)
        end 
    end

And invoke that mailer with the line:

    PostmanMailer.invitation(user).deliver_later

ActionMailer is using delivery jobs and enqueues email delivery as a job through Active Job, so I can't wrap this with being/rescue to insert custom exception.

How would you handle this? Is there any way for not monkey patching?
## [6][Ruby refinements have ONE good use case](https://www.reddit.com/r/ruby/comments/jsur5u/ruby_refinements_have_one_good_use_case/)
- url: http://www.soulcutter.com/articles/ruby-refinements-have-one-good-use-case.html
---

## [7][lazaronixon/react-native-turbolinks](https://www.reddit.com/r/ruby/comments/jt0cd9/lazaronixonreactnativeturbolinks/)
- url: https://github.com/lazaronixon/react-native-turbolinks
---

## [8][Creating a Weekly iOS Release Train](https://www.reddit.com/r/ruby/comments/jsukp6/creating_a_weekly_ios_release_train/)
- url: https://medium.com/pipedrive-engineering/welcome-aboard-the-pipedrive-ios-release-train-40fd9123ceac?source=friends_link&amp;sk=9c1d9183fb49b1648ccefe44deb233e3
---

## [9][Ruby pdf-forms gem and upload images to file field.](https://www.reddit.com/r/ruby/comments/js93sw/ruby_pdfforms_gem_and_upload_images_to_file_field/)
- url: https://www.reddit.com/r/ruby/comments/js93sw/ruby_pdfforms_gem_and_upload_images_to_file_field/
---
Hi,

I have to write some ruby code to fill the pdf form. I found this gem [https://github.com/jkraemer/pdf-forms](https://github.com/jkraemer/pdf-forms) which is good to fill the text fields. But how do I upload images as it has some input field to attach images. I tried something like this but none of them worked.

&amp;#x200B;

    [1] pry(main)&gt; pdftk = PdfForms.new('/usr/local/bin/pdftk')
    =&gt; #&lt;PdfForms::PdftkWrapper:0x007fbf617a0b48 @options={}, @pdftk="/usr/local/bin/pdftk"&gt;
    [2] pry(main)&gt; photo_file = File.open('/Users/aruprakshit/Documents/testdocs/download.jpeg')
    =&gt; #&lt;File:/Users/aruprakshit/Documents/testdocs/download.jpeg&gt;
    [3] pry(main)&gt; photo_file = File.open('/Users/aruprakshit/Documents/testdocs/ship-289664_960_720.jpg')
    =&gt; #&lt;File:/Users/aruprakshit/Documents/testdocs/ship-289664_960_720.jpg&gt;
    [4] pry(main)&gt; pdftk.fill_form file.to_s, Rails.root.join('public', 'photo.pdf').to_s, {'Photo 1' =&gt; photo_file.read, 'Project ID': 123334}
    =&gt; nil
    [5] pry(main)&gt; pdftk.fill_form file.to_s, Rails.root.join('public', 'photo.pdf').to_s, {'Photo 1' =&gt; '/Users/aruprakshit/Documents/testdocs/ship-289664_960_720.jpg', 'Project ID': 123334}
    =&gt; nil

Can anyone share some ideas if you worked on this?
## [10][Ruby :: How to add a try/catch clause to a TCP Socket? (to catch timeouts)](https://www.reddit.com/r/ruby/comments/jsemk2/ruby_how_to_add_a_trycatch_clause_to_a_tcp_socket/)
- url: https://www.reddit.com/r/ruby/comments/jsemk2/ruby_how_to_add_a_trycatch_clause_to_a_tcp_socket/
---
  

Hi everyone,

I’m a programmer who is also a beginner in Ruby.  I’ve inherited a Ruby program from another colleague, and I need to figure out the syntax to safely modify the original code.

The original code accepts a data record (a linked list, perhaps, I’m not sure) from an external source. The code then opens a socket on TCP 12345, sends the entire record, and then listens to the socket for a response from the remote host. The remote guy does some analysis and then sends back a string. The Ruby code then takes that string and appends it to the data record in a new field called “new\_information”. Here’s the code:

`require 'socket'`

`...accept data record "event"...`

`socket = TCPSocket.new("192.168.3.1", 12345)`

`socket.write (event.to_hash).to_s`

`response = socket.recv(1000)`

`socket.close`

`event.set("new_information", response)`

All of this works flawlessly. (For those who are curious, this is part of a Logstash implementation. Logstash allows Ruby code to help process incoming log records.)

But now I’m worried about the scenario where the remote host is down and unavailable. I need to include an if/else or try/catch structure so that if the Ruby socket times out and no answer is ever received, the “new\_information” field can still be populated with something meaningful. Here’s my stab at the modified code:

`require 'socket'`

`...accept data record "event"...`

`begin`

   `socket = TCPSocket.new("192.168.3.1", 12345)`

   `socket.write (event.to_hash).to_s`

   `response = socket.recv(1000)`

   `socket.close`

   `event.set("new_information", response)`

`rescue Errno::ETIMEDOUT`

   `p 'timeout'`

   `event.set("new_information", "remote_svr_down!")`

`end`

As you can probably tell, I don’t really know what I’m doing. Trouble is, the Logstash system with which I am working does not allow you to see the precise syntax error when the code crashes. So I have no idea if my modified code is nearly correct or way, way off.

Another question I have is… how long is the socket timeout here? I assume it is the default value. Is there a way to set that?

Thank you!
