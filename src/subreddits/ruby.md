# ruby
## [1][[EN] Don't Wait For Me! Scalable Concurrency for Ruby 3! / Samuel Williams @ioquatix](https://www.reddit.com/r/ruby/comments/iv1dk1/en_dont_wait_for_me_scalable_concurrency_for_ruby/)
- url: https://www.youtube.com/watch?v=Y29SSOS4UOc
---

## [2][Refactoring nested if/else code - 1 hour video based on the Gilded Rose kata](https://www.reddit.com/r/ruby/comments/iuv7if/refactoring_nested_ifelse_code_1_hour_video_based/)
- url: https://www.youtube.com/watch?v=2fhvENjz7pM
---

## [3][Ruby Antipatterns - I don't agree with all of them but the list is worth reading](https://www.reddit.com/r/ruby/comments/iv1pnm/ruby_antipatterns_i_dont_agree_with_all_of_them/)
- url: https://www.alchemists.io/articles/ruby_antipatterns/
---

## [4][How the way you define hash default value can generate side effects](https://www.reddit.com/r/ruby/comments/iv3llq/how_the_way_you_define_hash_default_value_can/)
- url: https://medium.com/rubycademy/hash-shift-using-default-value-b0d7cb62ffa2
---

## [5][Add "super" command in comware.rb [Oxidized]](https://www.reddit.com/r/ruby/comments/iv63fs/add_super_command_in_comwarerb_oxidized/)
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

&gt;`if vars(:super)`  
`cmd "super"`  
`cmd vars(:super)`  
`end`

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

&amp;#x200B;
## [6][Programming only with classes](https://www.reddit.com/r/ruby/comments/iukgi0/programming_only_with_classes/)
- url: https://weird-programming.dev/oop/classes-only.html
---

## [7][Introducing Hanami::API](https://www.reddit.com/r/ruby/comments/iuffbz/introducing_hanamiapi/)
- url: http://hanamirb.org/blog/2020/02/26/introducing-hanami-api/
---

## [8][Wrong number of arguments error](https://www.reddit.com/r/ruby/comments/iust7s/wrong_number_of_arguments_error/)
- url: https://www.reddit.com/r/ruby/comments/iust7s/wrong_number_of_arguments_error/
---
FIXED!

Hello! I am starting a Ruby assignment for school and I keep running into the wrong number of arguments error when im trying to set up a basic testing file.

tests.rb:

`def test1()`

    `puts "1"`
    
    `sides = 6`
    
    `colour = ':red'`
    
    `die =` [`Die.new`](https://Die.new)`(6, ":red")`

`end`

`test1`

die.rb

`class Die &amp;lt; Randomizer`

    `def intitialize(sides, colour)`
    
    	`u/sides = sides`
    
    	`u/colour = colour`
    
    `end`

`...other methods....`

`end`

Error:

`tests.rb:9:in \`initialize': wrong number of arguments (given 2, expected 0) (ArgumentError)\`

`from tests.rb:9:in \`new'\`

`from tests.rb:9:in \`test1'\`

`from tests.rb:13:in \`&amp;lt;main&amp;gt;'\`

Could someone help me with whats going on?

&amp;amp;#x200B;

UPDATE: if i take out the initialize() and just use attr\_writer I can get it to work, but I feel like that is bad form.
## [9][Rails Concerns: To Concern Or Not To Concern](https://www.reddit.com/r/ruby/comments/iuh33i/rails_concerns_to_concern_or_not_to_concern/)
- url: https://blog.appsignal.com/2020/09/16/rails-concers-to-concern-or-not-to-concern.html
---

## [10][Ruby 3.0 Interpolated Strings Are No Longer Frozen](https://www.reddit.com/r/ruby/comments/iuc4iz/ruby_30_interpolated_strings_are_no_longer_frozen/)
- url: https://scriptday.com/blog/2020/09/16/ruby-3-0-interpolated-strings-are-no-longer-frozen
---

