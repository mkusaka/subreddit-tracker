# ruby
## [1][A 10-day long hackathon called JamCraft 5 starts in less than 9 days. DragonRuby is sponsoring this jam and is providing all participants with a license to Game Toolkit. Here's your chance to build a game with Ruby.](https://www.reddit.com/r/ruby/comments/fsr5s8/a_10day_long_hackathon_called_jamcraft_5_starts/)
- url: http://jamcraft.dragonruby.org/
---

## [2][RailsConf 2020.2 Couch Edition](https://www.reddit.com/r/ruby/comments/fsk4w3/railsconf_20202_couch_edition/)
- url: https://twitter.com/railsconf/status/1245056735892525057?s=20
---

## [3][A new feature proposal today—Endless method definition: def: value(args) = expression](https://www.reddit.com/r/ruby/comments/fstnrz/a_new_feature_proposal_todayendless_method/)
- url: https://bugs.ruby-lang.org/issues/16746
---

## [4][Ruby versions 2.4.10, 2.5.8, 2.6.6, and 2.7.1](https://www.reddit.com/r/ruby/comments/fsf83w/ruby_versions_2410_258_266_and_271/)
- url: https://www.ruby-lang.org/en/news/
---

## [5][I wrote an RCON client gem in Ruby](https://www.reddit.com/r/ruby/comments/fspsh7/i_wrote_an_rcon_client_gem_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fspsh7/i_wrote_an_rcon_client_gem_in_ruby/
---
[https://github.com/hernanat/rconrb](https://github.com/hernanat/rconrb)

Background:

My younger siblings wanted us all to play Minecraft together while everyone is sheltering in place, so I ended up configuring a server for us on AWS. I also created an administrative app in Rails to manage that which I'll eventually open-source.

I know there is an existing RCON gem, but the project appears to be dead and I wasn't particularly fond of the implementation.

I plan to write a Minecraft specific wrapper around this as well :)
## [6][How to manage Ruby dependencies from the point of view of a NodeJS / npm dev?](https://www.reddit.com/r/ruby/comments/fssk1h/how_to_manage_ruby_dependencies_from_the_point_of/)
- url: https://www.reddit.com/r/ruby/comments/fssk1h/how_to_manage_ruby_dependencies_from_the_point_of/
---
I come from NodeJS where dependencies are managed:  


* Locally with \`npm install X\`
* Optionally globally with \`npm install --global X\`

With Ruby, it seems that its by default global with \`gem install X\`

How to install locally? I dont feel comfortable installing anything globally unless its a CLI tool that I use for different projects.
## [7][What's the best way to add logging to a gem?](https://www.reddit.com/r/ruby/comments/fshcil/whats_the_best_way_to_add_logging_to_a_gem/)
- url: https://www.reddit.com/r/ruby/comments/fshcil/whats_the_best_way_to_add_logging_to_a_gem/
---
Especially when rails is involved (or expected to be, I suppose) I see a lot of `if defined?(Rails)`. While googling, and the way I prefer honestly, is injection.

    module MyGem
      class &lt;&lt; self
        attr_writer :logger
    
        def logger
          @logger ||= Logger.new($stdout).tap do |log|
            log.progname = self.name
          end
        end
      end
    end

with an initializer with `MyGem.logger = Rails.logger`.

What's preferred these days? Anyone have an opinion on the topic?
## [8][Rails adoption higher than combined PHP frameworks on large sites](https://www.reddit.com/r/ruby/comments/fsg94m/rails_adoption_higher_than_combined_php/)
- url: https://www.reddit.com/r/ruby/comments/fsg94m/rails_adoption_higher_than_combined_php/
---
Comparing stats for PHP and Rails on https://www.similartech.com/compare/php-vs-ruby-on-rails Rails has 485,094 to PHP's 7,153,038. However, the combined total for PHP frameworks - Laravel, Symfony, CodeIgniter, Yii, MODX &amp; CakePHP - is 447,995 so when a framework is used (not a CMS) Rails adoption is slightly higher than PHP on large sites. Am I missing something?
## [9][Security Issue: CVE-2020-10933: Heap exposure vulnerability in the socket library](https://www.reddit.com/r/ruby/comments/fsdskg/security_issue_cve202010933_heap_exposure/)
- url: https://www.reddit.com/r/ruby/comments/fsdskg/security_issue_cve202010933_heap_exposure/
---
[https://www.ruby-lang.org/en/news/2020/03/31/heap-exposure-in-socket-cve-2020-10933/](https://www.ruby-lang.org/en/news/2020/03/31/heap-exposure-in-socket-cve-2020-10933/)
## [10][Maintaining Sanity with Ruby Under A Lockdown.](https://www.reddit.com/r/ruby/comments/fs535x/maintaining_sanity_with_ruby_under_a_lockdown/)
- url: https://emmanuelhayford.com/maintaining-sanity-with-ruby-under-a-lockdown/
---

