# ruby
## [1][My new gem: Heya is a campaign mailer for Rails](https://www.reddit.com/r/ruby/comments/g1ezhp/my_new_gem_heya_is_a_campaign_mailer_for_rails/)
- url: https://www.reddit.com/r/ruby/comments/g1ezhp/my_new_gem_heya_is_a_campaign_mailer_for_rails/
---
Hello all! I wanted to share a side project I've been working on at Honeybadger. It's called [Heya](https://www.heya.email), and it's a sequence mailer for Rails. I think the README says it best:

&gt;Heya is a campaign mailer for Rails. Think of it like ActionMailer, but for timed email sequences. It can also perform other actions like sending a text message.

I built Heya to replace part of Intercom, which we finally migrated away from last month. 🎉 My long-term plan for Heya is to build more features around customer communication and automation in our Rails SaaS app.

There are some other gems out there that tackle this problem, but I think the approach I landed on with Heya is unique, and I like how well it fits into the Rails ecosystem (for instance, it can be combined with mailkick for managing subscriptions, maildown for markdown support, and it feels just like \`ActionMailer\` 👌).

Heya is completely open source and free for non-profits and small businesses. There's a $99/year commercial license for established companies after a 30 day trial (this may change in the future--that's what it is right now). Here are the links:

[https://www.heya.email](https://www.heya.email)

[https://github.com/honeybadger-io/heya](https://github.com/honeybadger-io/heya)
## [2][Exploring Method Arguments in Ruby (Part 2 )](https://www.reddit.com/r/ruby/comments/g1rri1/exploring_method_arguments_in_ruby_part_2/)
- url: https://www.ombulabs.com/blog/ruby/learning/methods-arguments-ruby-part2.html
---

## [3][Anyway Config: Keep your Ruby configuration sane and escape the ENV Hell](https://www.reddit.com/r/ruby/comments/g1d1n5/anyway_config_keep_your_ruby_configuration_sane/)
- url: https://evilmartians.com/chronicles/anyway-config-keep-your-ruby-configuration-sane
---

## [4][Question on the devise gem: How do I override SMTP?](https://www.reddit.com/r/ruby/comments/g1n6lk/question_on_the_devise_gem_how_do_i_override_smtp/)
- url: https://www.reddit.com/r/ruby/comments/g1n6lk/question_on_the_devise_gem_how_do_i_override_smtp/
---
By default Rails Devise sends emails through SMTP, but in my case I need to send them through an API instead. What's a good way to override Devise's defaults?

I'm thinking if I could knew how to generate the email texts based on `resource_params` I would be able to use that in `users/passwords_controller.rb` for example to send out the emails within the `create` method.

Any help would be greatly appreciated!
## [5][Dissecting Rails Migrations](https://www.reddit.com/r/ruby/comments/g1pl1y/dissecting_rails_migrations/)
- url: https://blog.appsignal.com/2020/04/14/dissecting-rails-migrationsl.html
---

## [6][How to build a command line application in Ruby](https://www.reddit.com/r/ruby/comments/g1didk/how_to_build_a_command_line_application_in_ruby/)
- url: https://www.youtube.com/watch?v=A4VP-suGiJI
---

## [7][Ruby Concurrency Final Report - not the end, just the beginning!](https://www.reddit.com/r/ruby/comments/g12yai/ruby_concurrency_final_report_not_the_end_just/)
- url: https://www.codeotaku.com/journal/2020-04/ruby-concurrency-final-report/index
---

## [8][Help with starting new project](https://www.reddit.com/r/ruby/comments/g1hlc8/help_with_starting_new_project/)
- url: https://www.reddit.com/r/ruby/comments/g1hlc8/help_with_starting_new_project/
---
Hello! I've just started learning ruby and so far it seems pretty nice, im a CS student at University currently so I've just been doing this on the side. My knowledge is still pretty basic however I've been doing the Odin Project and it just seems extremely slow so im trying to do some independent projects. 

Anyways, I've been reading this history book and id like to develop a "web scraper" to find web-pages where the family name is mentioned. For example their Wikipedia page if they have one, such as the Tudors. I don't need a step by step but where are some places I can look to head in the right direction. I'm not looking for an overnight solution or anything so I expect it to take a while. Sorry for the long post, thanks!
## [9][Using named captures to extract information from Strings - Aliou Diallo](https://www.reddit.com/r/ruby/comments/g13ud1/using_named_captures_to_extract_information_from/)
- url: https://aliou.me/posts/2020/04/regexp-named-captures/
---

## [10][hex2bin BUT then extract bit fields of binary, back to hex.](https://www.reddit.com/r/ruby/comments/g16dhr/hex2bin_but_then_extract_bit_fields_of_binary/)
- url: https://www.reddit.com/r/ruby/comments/g16dhr/hex2bin_but_then_extract_bit_fields_of_binary/
---
Hi. I'm new to Ruby, and what seems daunting to me, may be trivial to others.

&amp;#x200B;

I'm trying to create a script, in any language, but Ruby, Perl, maybe Python are available. But Ruby seems capable.

&amp;#x200B;

To help me understand data from my spec, I need to make it easier to extract bit fields from a 128 bit value. So here is my need:

&amp;#x200B;

Take input hex string 32 chars long.

Convert it to a binary array so I can access subset of the bits.

Then convert any binary range, between 127 and 0, back to hex.

&amp;#x200B;

For example:

\--------------

\# input

myhex = "ABCD0123ABCD0123ABCD0123ABCD0123"

&amp;#x200B;

\# conversion

mybin\[127:0\] = hex2bin(myhex)

&amp;#x200B;

\# extraction

value1 = bin2hex(mybin\[22:21\]) # 2 bits (range 0x0 to 0x3)

value2 = bin2hex(mybin\[65:60\]) # 6 bits (0x0 to 0x3F)

value3 = bin2hex(mybin\[127:110\]) # 18 bits (0x0 to 0x3\_FFFF)

&amp;#x200B;

print value1, value2, value3

\--------------

&amp;#x200B;

There seem to be plenty of hex2bin/bin2hex programs that utilize built in functions, but extracting fields from binary, then convert back to hex is harder to find examples.

&amp;#x200B;

Thoughts?

&amp;#x200B;

Thanks.
