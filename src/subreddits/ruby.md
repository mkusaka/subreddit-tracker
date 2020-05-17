# ruby
## [1][How to Fix Slow Code in Ruby](https://www.reddit.com/r/ruby/comments/gldh8k/how_to_fix_slow_code_in_ruby/)
- url: https://engineering.shopify.com/blogs/engineering/how-fix-slow-code-ruby
---

## [2][New ruby on rails fullstack template Repo](https://www.reddit.com/r/ruby/comments/glf19d/new_ruby_on_rails_fullstack_template_repo/)
- url: https://www.reddit.com/r/ruby/comments/glf19d/new_ruby_on_rails_fullstack_template_repo/
---
Hey, I created this template for myself  and using it for some while, I extracted it for whoever wanna make use of it.

[https://github.com/ahmgeek/fullstack](https://github.com/ahmgeek/fullstack)
## [3][Game of Life in one Ruby statement... inspired by APL](https://www.reddit.com/r/ruby/comments/gkvxaq/game_of_life_in_one_ruby_statement_inspired_by_apl/)
- url: https://zverok.github.io/blog/2020-05-16-ruby-as-apl.html
---

## [4][Sending multiple commands to a list of servers using Net::SSH - What can I do better?](https://www.reddit.com/r/ruby/comments/gl7mb9/sending_multiple_commands_to_a_list_of_servers/)
- url: https://www.reddit.com/r/ruby/comments/gl7mb9/sending_multiple_commands_to_a_list_of_servers/
---
**The idea**

At my work, I'm tasked with retrieving logs for a list of online orders from a total of 12 servers. To do this, I SSH into them (with a jump host in between) one by one, run the script for every order, then parse the output. Overall, a very tedious process that I'm sure can be automated. After trying to use a few libraries in Powershell and Java with no luck, I stumbled upon a suggestion to use Net::SSH in Ruby because it uses your ssh config which simplifies things significantly. The only downside is that I have very very limited experience in Ruby; like I made one script in Ruby 4 years ago to check it out and that's it. As such, I need some help.

So the topology looks like this:

    W -&gt; J -&gt; S
    
    W = My Windows VM
    J = Red Hat jump box
    S = The Red Hat production servers

And the basic idea is the following:

1. Connect to every server by SSH

2. On every server, run the commands for all orders which are passed by command line

3. Associate each output with the order (eventually, auto parse and output to files)

Here is what I've come up with so far, and it is working on my home lab. But before I bring it in to work to propose it and try it on production servers, I want to know how the code can be cleaned up or done more efficiently.

**Code**

    require 'net/ssh'
    require 'io/console'
    require 'pp'
    
    ORDERS = ARGV.dup()     
    ARGV.clear()            
    
    SERVERS = [
        "prod1",
        "prod2",
        "prod3"
    ]
    
    COMMAND = "aux_search.sh"
    
    CONFIG = "C:/Users/myuser/.ssh/config"
    
    RESULTS = ORDERS.each_with_object({}) { |k, v| v[k] = "" }
    
    puts("Enter your username: ")
    USERNAME = gets().chomp()
    puts("Enter your password for accessing the servers: ")
    SERVERPASSWORD = STDIN.noecho(&amp;:gets).chomp()
    puts("Enter your sudo password: ")
    SUDOPASSWORD = STDIN.noecho(&amp;:gets).chomp()
    
    SESSIONS = []
    SERVERS.each do |server|
        session = Net::SSH.start(server, USERNAME, :password =&gt; SERVERPASSWORD, :config =&gt; CONFIG, 
            :verbose =&gt; :debug, :auth_methods =&gt; ["publickey", "password"], :verify_host_key =&gt; :accept_new)
        
        SESSIONS &lt;&lt; session
    end
    
    SESSIONS.each do |ssh|
        ORDERS.each do |order|
            ssh.open_channel() do |channel|
                channel.on_data do |_, data|
                    RESULTS[order].concat(data)
                    if data.match(/\[sudo\]/)  
                        channel.send_data("#{SUDOPASSWORD}\n")
                    end
                end
    
                channel.request_pty() do |_, success|
                    raise "Failed to request TTY for order #{order}" unless success
                end
    
                channel.exec("sudo ./#{COMMAND} #{order}") do |_, success| 
                    raise "Could not execute command #{COMMAND} #{order}" unless success
                end
            end
        end
        ssh.loop()
    end
    
    puts("===Results===")
    pp(RESULTS)
    
    #TODO: Parse results

And the ssh config I am using

**ssh config**
    
    Host jumpbox
        HostName 192.168.16.2
        User johrus2
        IdentityFile C:\Users\myuser\.ssh\id_rsa3
        Port 22
    
    Host prod1
        HostName 192.168.16.3
        User johrus2
        IdentityFile C:\Users\myuser\.ssh\id_rsa3
        ProxyJump jumpbox
        Port 22
    
    Host prod2
        HostName 192.168.16.4
        User johrus2
        IdentityFile C:\Users\myuser\.ssh\id_rsa3
        ProxyJump jumpbox
        Port 22
    
    Host prod3
        HostName 192.168.16.5
        User johrus2
        IdentityFile C:\Users\myuser\.ssh\id_rsa3
        ProxyJump jumpbox
        Port 22
    
Any help or advice is appreciated!
## [5][Need help running a script](https://www.reddit.com/r/ruby/comments/gl662g/need_help_running_a_script/)
- url: https://www.reddit.com/r/ruby/comments/gl662g/need_help_running_a_script/
---
I'm attempting to run a Ruby script ([linked here](https://gist.github.com/dhh/d3c8cf9309b662047257b7e583c3f595#file-dcp-converter-rb-L8)) that was shared by DHH to convert a number of .dcp Leica Q camera profiles to Leica M10 camera profiles.

I have a directory I've titled `qtom10`. Within it, I have the `dcp-converter.rb` script along with two directories—One titled `input` with all of the input files within it and one that's empty titled `output`.

I have the latest version of Ruby Gems installed, Bundler and Nokigiri.

When I run the script using `ruby dcp-converter.rb`, here's the error I get:

    ruby dcp-converter.rb
    Decompiling ./input/LEICA Q (Typ 116) Kodak Royal 400 - L.dcp into XML
    sh: ./bin/dcpTool: No such file or directory
    Replacing camera model: LEICA Q (Typ 116) -&gt; LEICA M10
    Traceback (most recent call last):
    	4: from dcp-converter.rb:26:in `&lt;main&gt;'
    	3: from dcp-converter.rb:26:in `each'
    	2: from dcp-converter.rb:37:in `block in &lt;main&gt;'
    	1: from dcp-converter.rb:20:in `replace_camera_model'
    dcp-converter.rb:20:in `read': No such file or directory @ rb_sysopen - ./output/LEICA M10 Kodak Royal 400 - L.xml (Errno::ENOENT)

Not sure where to go from here or what I'm doing wrong... HELP!

Here's the actual script from DHH:

    #!/usr/bin/env ruby
    
    # Requires that you have ./bin/dcpTool from https://sourceforge.net/projects/dcptool/
    
    require 'rubygems'
    require 'bundler/setup'
    require 'nokogiri'
    
    input_camera_model  = ARGV[0] || "LEICA Q (Typ 116)"
    output_camera_model = ARGV[1] || "LEICA M10"
    
    input_dir  = ARGV[2] || "./input"
    output_dir = ARGV[3] || "./output"
    
    def convert_profile_name(profile_name, input_camera_model, output_camera_model)
      File.basename(profile_name.gsub(/#{input_camera_model.gsub(/\(/, "\\(").gsub(/\)/, "\\)")}/, output_camera_model), ".dcp")
    end
    
    def replace_camera_model(xml_profile_filename, output_camera_model)
      profile_doc = Nokogiri::XML(File.read(xml_profile_filename))
      profile_doc.xpath('//UniqueCameraModelRestriction').first.content = output_camera_model  
    
      File.open(xml_profile_filename, "w+") { |file| file.write(profile_doc.to_xml) }
    end
    
    Dir.entries(input_dir).reject { |file| file =~ /^(\.|\.\.)$/ }.each do |existing_profile|
      converted_profile = convert_profile_name(existing_profile, input_camera_model, output_camera_model)
      existing_dcp_filename  = File.join(input_dir, existing_profile)
      xml_filename = "#{File.join(output_dir, converted_profile)}.xml"
    
      decompile_command = "./bin/dcpTool -d '#{existing_dcp_filename}' '#{xml_filename}'"
    
      puts "Decompiling #{existing_dcp_filename} into XML"
      `#{decompile_command}`
    
      puts "Replacing camera model: #{input_camera_model} -&gt; #{output_camera_model}"
      replace_camera_model(xml_filename, output_camera_model)
    
    
      converted_dcp_filename = "#{File.join(output_dir, converted_profile)}.dcp"
      recompile_command = "./bin/dcpTool -c '#{xml_filename}' '#{converted_dcp_filename}'"
    
      puts "Recompiling XML into #{converted_dcp_filename}"
      `#{recompile_command}`
    
      File.delete(xml_filename)
    
      puts
    end
## [6][Can you have a "literal" string literal in Ruby? Can I make single-quoted string *not* escape \\?](https://www.reddit.com/r/ruby/comments/gkuphz/can_you_have_a_literal_string_literal_in_ruby_can/)
- url: https://www.reddit.com/r/ruby/comments/gkuphz/can_you_have_a_literal_string_literal_in_ruby_can/
---
I've been doing the 2015 Advent of Code problems to practice my Ruby and I ran across an interesting issue on Day 8.  Day 8 gives you a list of strings and a list of "escape sequences" --- `\"`, `\\`, and `\x[0-9a-f]{2}`. You need to count how many characters in each string are escape characters.

The problem I ran into is that I can't figure out how to get Ruby to read a literal string as having two backslashes (`\\`) in a row. I thought single-quoted strings in Ruby ignored escape sequences, but even in single-quoted Ruby strings, `\\` is an escaped backslash character.

    irb(main):003:0&gt; string1 = 'test\test'
    =&gt; "test\\test"
    irb(main):004:0&gt; string2 = 'test\\test'
    =&gt; "test\\test"
    irb(main):005:0&gt; string1 == string2
    =&gt; true
    irb(main):006:0&gt; string1.size
    =&gt; 9
    irb(main):007:0&gt; string2.size
    =&gt; 9
    irb(main):008:0&gt; string1.bytes
    =&gt; [116, 101, 115, 116, 92, 116, 101, 115, 116]
    irb(main):009:0&gt; string2.bytes
    =&gt; [116, 101, 115, 116, 92, 116, 101, 115, 116]

I solved the AoC problem by reading the strings character by character into an array and working with that, but that doesn't feel like a very Ruby-like solution to me.

My question is: Is there a way to create a string literal in Ruby that ignores \*all\* escaped characters, or a way to get at the "raw" unescaped data that was passed into the string when it was created?
## [7][Why Fullstaq Ruby?](https://www.reddit.com/r/ruby/comments/gk5i19/why_fullstaq_ruby/)
- url: https://www.joyfulbikeshedding.com/blog/2020-05-15-why-fullstaq-ruby.html
---

## [8][sportdb-importers Gem v1.0 - tools 'n' scripts for importing sports (incl. football) data in comma-separated values (csv) format](https://www.reddit.com/r/ruby/comments/gke7v6/sportdbimporters_gem_v10_tools_n_scripts_for/)
- url: https://github.com/sportdb/sport.db/tree/master/sportdb-importers
---

## [9][A heavily tested (2k lines) and commented classic Red Black Tree implementation in Python and Ruby. Great for learning the material.](https://www.reddit.com/r/ruby/comments/gk51nf/a_heavily_tested_2k_lines_and_commented_classic/)
- url: https://www.reddit.com/r/ruby/comments/gk51nf/a_heavily_tested_2k_lines_and_commented_classic/
---
Back when I was trying to implement the structure, I could not find any open source implementations that were well written and commented. I did not manage to find any implementation that had any significant amount of tests and as such was not sure if it even worked correctly.

I tried my best to describe the different operations needed thoroughly and have written a lot of tests (functional too) covering all operations, with drawn out trees in comments.

https://github.com/stanislavkozlovski/Red-Black-Tree

Any feedback is greatly appreciated :)
## [10][Foot Traffic: pure Ruby DSL for Chrome scripting based on Ferrum](https://www.reddit.com/r/ruby/comments/gjqaji/foot_traffic_pure_ruby_dsl_for_chrome_scripting/)
- url: https://github.com/lewagon/foot_traffic
---

