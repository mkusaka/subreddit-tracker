# ruby
## [1][Ruby 3.0 Added Leading Arguments Forwarding](https://www.reddit.com/r/ruby/comments/ivoker/ruby_30_added_leading_arguments_forwarding/)
- url: https://www.reddit.com/r/ruby/comments/ivoker/ruby_30_added_leading_arguments_forwarding/
---
[https://scriptday.com/blog/2020/09/12/ruby-3-0-added-leading-arguments-forwarding](https://scriptday.com/blog/2020/09/12/ruby-3-0-added-leading-arguments-forwarding)
## [2][Automate Rails server provisioning and deployment using Ansible and Capistrano](https://www.reddit.com/r/ruby/comments/ivpbfc/automate_rails_server_provisioning_and_deployment/)
- url: https://rubyyagi.com/rails-deploy-automate-ansible/
---

## [3][We Made Puma Faster With Sleep Sort](https://www.reddit.com/r/ruby/comments/ivasiq/we_made_puma_faster_with_sleep_sort/)
- url: https://www.speedshop.co/2020/09/17/we-made-puma-faster-with-sleep-sort.html
---

## [4][[EN] Running Rack and Rails Faster with TruffleRuby / Benoit Daloze @eregontp](https://www.reddit.com/r/ruby/comments/ivcyyk/en_running_rack_and_rails_faster_with_truffleruby/)
- url: https://www.youtube.com/watch?v=281YdMYRAsk
---

## [5][[EN] Don't Wait For Me! Scalable Concurrency for Ruby 3! / Samuel Williams @ioquatix](https://www.reddit.com/r/ruby/comments/iv1dk1/en_dont_wait_for_me_scalable_concurrency_for_ruby/)
- url: https://www.youtube.com/watch?v=Y29SSOS4UOc
---

## [6][Method to randomly select name from an array](https://www.reddit.com/r/ruby/comments/ivanf4/method_to_randomly_select_name_from_an_array/)
- url: https://www.reddit.com/r/ruby/comments/ivanf4/method_to_randomly_select_name_from_an_array/
---
Hi!

I am trying to solve a problem that uses a method to randomly select a name from array. Here is the code that I have generated thus far:

**# Please write your code between this line...**

**def random\_name(array)**

**array = \["Ana", "Ollie"\].sample**

**end**

**# ...and between this line**

**puts "My name is #{random\_name()}"**

With this code I get the error message: "wrong number of arguments (given 0, expected 1)"

I am new to Ruby and coding in general, so I am not sure where to go from here. Any help would be greatly appreciated!
## [7][Ruby Antipatterns - I don't agree with all of them but the list is worth reading](https://www.reddit.com/r/ruby/comments/iv1pnm/ruby_antipatterns_i_dont_agree_with_all_of_them/)
- url: https://www.alchemists.io/articles/ruby_antipatterns/
---

## [8][Document Template Solutions?](https://www.reddit.com/r/ruby/comments/iv6sn1/document_template_solutions/)
- url: https://www.reddit.com/r/ruby/comments/iv6sn1/document_template_solutions/
---
Hello! I work in a financial service company which generates a lot of documents. We implement documents as HTML and print them to PDF for consumption, which I think is a pretty standard approach.

Although, I'm finding this to be quite hard to maintain as we get a lot of minor changes on copy over time and errors that creep in because the owners of the documents are not able to edit the HTML.

I wonder, is this a common problem in other companies? Are there other more efficient solutions to manage document templates?

PS: Posting here because we use Ruby and Ruby specific solutions are welcome, but this is a generic problem that I'm sure is faced by people working with other technologies
## [9][How the way you define hash default value can generate side effects](https://www.reddit.com/r/ruby/comments/iv3llq/how_the_way_you_define_hash_default_value_can/)
- url: https://medium.com/rubycademy/hash-shift-using-default-value-b0d7cb62ffa2
---

## [10][Add "super" command in comware.rb [Oxidized]](https://www.reddit.com/r/ruby/comments/iv63fs/add_super_command_in_comwarerb_oxidized/)
- url: https://www.reddit.com/r/ruby/comments/iv63fs/add_super_command_in_comwarerb_oxidized/
---
Hello,

I'm using Oxidized to backup my swtiches.

I need to use to use the "super" command on some of my comware switches to have access to all the commands.

This command is unfortunately not natively used in the comware.rb file so I'm trying to edit this file. The """only""" issue is that I've never used ruby...

I've tried doing something like this but it's not working correctly. Any ideas ?

**comware.rb**

    class Comware &lt; Oxidized::Model
      # HP (A-series)/H3C/3Com Comware
    
      # sometimes the prompt might have a leading nul or trailing ASCII Bell (^G)
      prompt /^\0*(&lt;[\w.-]+&gt;).?$/
      comment '# '
    
      # example how to handle pager
      # expect /^\s*---- More ----$/ do |data, re|
      #  send ' '
      #  data.sub re, ''
      # end
    
      cmd :all do |cfg|
        # cfg.gsub! /^.*\e\[42D/, ''        # example how to handle pager
        # skip rogue ^M
        cfg = cfg.delete "\r"
        cfg.cut_both
      end
    
      cmd :secret do |cfg|
        cfg.gsub! /^( snmp-agent community).*/, '\\1 &lt;configuration removed&gt;'
        cfg.gsub! /^( password hash).*/, '\\1 &lt;configuration removed&gt;'
        cfg.gsub! /^( password cipher).*/, '\\1 &lt;configuration removed&gt;'
        cfg
      end
    
      cfg :telnet do
        username /^(Username|login):/
        password /^Password:/
      end
    
      cfg :telnet, :ssh do
        # enable command-line mode on SMB comware switches (HP V1910, V1920)
        # autodetection is hard, because the 'summary' command is paged, and
        # the pager cannot be disabled before _cmdline-mode on.
        if vars :comware_cmdline
          post_login do
            # HP V1910, V1920
            cmd '_cmdline-mode on', /(#{@node.prompt}|Continue)/
            cmd 'y', /(#{@node.prompt}|input password)/
            cmd vars(:comware_cmdline)
    
            # HP V1950
            cmd 'xtd-cli-mode on', /(#{@node.prompt}|Continue)/
            cmd 'y', /(#{@node.prompt}|input password)/
            cmd vars(:comware_cmdline)
          end
        end
    
        post_login 'screen-length disable'
        post_login 'undo terminal monitor'

&gt;`if vars(:super)cmd "super"cmd vars(:super)end`

        pre_logout 'quit'
      end
    
      cmd 'display version' do |cfg|
        cfg = cfg.each_line.reject { |l| l.match /uptime/i }.join
        comment cfg
      end
    
      cmd 'display device' do |cfg|
        comment cfg
      end
    
      cmd 'display device manuinfo' do |cfg|
        cfg = cfg.each_line.reject { |l| l.match 'FF'.hex.chr }.join
        comment cfg
      end
    
      cmd 'display current-configuration' do |cfg|
        cfg
      end
    end

**Oxidized config file**

    groups:
      switches-iutb:
        username: user
        password: pass
      switches-iuta:
        username: user
        password: pass
        vars:
          super: superpass

Here's what happens :

&gt;Sep 18 15:32:25 : lib/oxidized/worker.rb: Added switches-iuta/sw-iuta-00 to the job queue  
&gt;  
&gt;Sep 18 15:32:25 : lib/oxidized/job.rb: Starting fetching process for sw-iuta-00 at 2020-09-18 13:32:25 UTC  
&gt;  
&gt;Sep 18 15:32:25 : lib/oxidized/input/ssh.rb super @ sw-iuta-00 with expect: /\^\\0\*(&lt;\[\\w.-\]+&gt;).?$/  
&gt;  
&gt;Sep 18 15:32:25 : lib/oxidized/node.rb: Oxidized::SSH failed for sw-iuta-00

&amp;#x200B;
