# ruby
## [1][How To Integrate Video Conferencing API to Your Application](https://www.reddit.com/r/ruby/comments/jsstuu/how_to_integrate_video_conferencing_api_to_your/)
- url: http://integratevideoconferencing.space/
---

## [2][Most Popular Backend Frameworks (2012/2020)](https://www.reddit.com/r/ruby/comments/jsru4g/most_popular_backend_frameworks_20122020/)
- url: https://youtu.be/94LokRYL5n0
---

## [3][Creating a Weekly iOS Release Train](https://www.reddit.com/r/ruby/comments/jsukp6/creating_a_weekly_ios_release_train/)
- url: https://medium.com/pipedrive-engineering/welcome-aboard-the-pipedrive-ios-release-train-40fd9123ceac?source=friends_link&amp;sk=9c1d9183fb49b1648ccefe44deb233e3
---

## [4][I have to create a custom linter for an assignment using Ruby. After a lot of research, I have come across the gem called Parser. Since I am a beginner, I am afraid I don't completely understand it.](https://www.reddit.com/r/ruby/comments/jsr9v1/i_have_to_create_a_custom_linter_for_an/)
- url: https://www.reddit.com/r/ruby/comments/jsr9v1/i_have_to_create_a_custom_linter_for_an/
---
I have to build a very basic linter that checks for whitespaces, indentations, naming conventions and other styling issues. I want to do it the right way and use a parser as linters use parsers. I am excited to learn it but I just am not able to wrap my head around it. Can someone please explain to me how parser gem works and how can I use it to create linting methods?
## [5][Ruby pdf-forms gem and upload images to file field.](https://www.reddit.com/r/ruby/comments/js93sw/ruby_pdfforms_gem_and_upload_images_to_file_field/)
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
## [6][Copy config.yml (bundled in .gem) to home directory after installation](https://www.reddit.com/r/ruby/comments/js6obr/copy_configyml_bundled_in_gem_to_home_directory/)
- url: https://www.reddit.com/r/ruby/comments/js6obr/copy_configyml_bundled_in_gem_to_home_directory/
---
Hey there, ruby community!       

I'd like to include a config.yml file with my gem, which should get copied to the users home directory, either before running the gem or on  the first gem run.

I have thought of the following options:

* Create a default config hash in ruby and write that to a file on the first run of the program. Drawbacks: No comments in yaml
* Fetch the config from a public repo. Drawbacks: Online only

Is there a way to bundle the gem with the config file included, and reference that file from within the code after the gem has been installed with gem install?

Thanks! :)
## [7][Ruby :: How to add a try/catch clause to a TCP Socket? (to catch timeouts)](https://www.reddit.com/r/ruby/comments/jsemk2/ruby_how_to_add_a_trycatch_clause_to_a_tcp_socket/)
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
## [8][Rebuilding Redis in Ruby - New chapters available](https://www.reddit.com/r/ruby/comments/jrt2up/rebuilding_redis_in_ruby_new_chapters_available/)
- url: https://www.reddit.com/r/ruby/comments/jrt2up/rebuilding_redis_in_ruby_new_chapters_available/
---
About four months ago I [posted about a project](https://www.reddit.com/r/ruby/comments/hmxcgv/redis_in_ruby_a_work_in_progress_online_book/) I've been working on, [Rebuilding Redis in Ruby](https://redis.pjam.me). There are six more chapters available now, so I figured I would post an update. Especially since the first three chapters were pretty short, and the later ones contain way more information:

1. [Chapter 4](https://redis.pjam.me/post/chapter-4-adding-missing-options-to-set/): Completing the SET command, that includes handling expiration and all the good stuff
2. [Chapter 5](https://redis.pjam.me/post/chapter-5-redis-protocol-compatibility/): Making the server RESP (Redis Protocol) compatible
3. [Chapter 6](https://redis.pjam.me/post/chapter-6-building-a-hash-table/): Building a Hash replacement, from scratch (including replicating the hashing algorithm used in Redis, SipHash)
4. [Chapter 7](https://redis.pjam.me/post/chapter-7-adding-list-commands/): Adding all List commands
5. [Chapter 8](https://redis.pjam.me/post/chapter-8-adding-hash-commands/): Adding all hash commands\* (\*except HSCAN)
6. [Chapter 9](https://redis.pjam.me/post/chapter-9-adding-set-commands/): Adding all set commands\* (\*except SSCAN)

If you have any feedback, please feel free to share it here or anywhere else (twitter, mail, etc ...), and if you're interested you can check what the future chapters will be [on the site](https://redis.pjam.me/chapters/):

1. Chapter 10: Sorted Sets
2. Chapter 11: Bitmaps:
3. Chapter 12: HyperLogLogs
4. Chapter 13: Geo Commands
5. Chapter 14: Pub/Sub
6. Chapter 15: Redis as an LRU cache
7. Chapter 16: Transactions

Hope you'll enjoy it!

PS: I definitely don't want to spam this sub, I'm only really on posting one more time about this, when all sixteen chapters will be completed.
## [9][What other things equals nil](https://www.reddit.com/r/ruby/comments/js2nel/what_other_things_equals_nil/)
- url: https://www.reddit.com/r/ruby/comments/js2nel/what_other_things_equals_nil/
---
I'm checking for equality between two variables, first and second. The first variable can be nil.   
I wanted to be sure other than second being nil also is there any cases where first == second would be true when first is nil?

i.e. first, second = nil, nil

first == second # true 

Is there another other cases for first being nil and second being anything else that can make first == second true? 

As far as I know, it can't be, but I'm not 100 % sure.
## [10][The art of errors in Ruby](https://www.reddit.com/r/ruby/comments/jrnk8v/the_art_of_errors_in_ruby/)
- url: https://longliveruby.com/articles/art-of-errors-in-ruby
---

