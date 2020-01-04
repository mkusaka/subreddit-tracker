# ruby
## [1][Extracting a tidy PORO from a messy Active Record model - Code with Jason](https://www.reddit.com/r/ruby/comments/ejlnt0/extracting_a_tidy_poro_from_a_messy_active_record/)
- url: https://www.codewithjason.com/extracting-tidy-poro-messy-active-record-model/
---

## [2][Help getting started with ruby. I can't get the directories created.](https://www.reddit.com/r/ruby/comments/ejr8i0/help_getting_started_with_ruby_i_cant_get_the/)
- url: https://www.reddit.com/r/ruby/comments/ejr8i0/help_getting_started_with_ruby_i_cant_get_the/
---
I can't get `rails new blog` to work. I have all the permissions. Removing --pretend flag isn't working as expected. 

    dev@vmdev:~$ ruby -v
    ruby 2.5.1p57 (2018-03-29 revision 63029) [x86_64-linux-gnu]
    dev@vmdev:~$ rails --version
    Rails 6.0.2.1
    dev@vmdev:~$ rails new blog -p
          create
          create  README.md
          create  Rakefile
          create  .ruby-version
          create  config.ru
          create  .gitignore
          create  Gemfile
          create  package.json
          create  app
          create  app/assets/config/manifest.js
          create  app/assets/stylesheets/application.css
          create  app/channels/application_cable/channel.rb
          create  app/channels/application_cable/connection.rb
          create  app/controllers/application_controller.rb
          ...........[truncated]..............
          create  storage
          create  storage/.keep
          create  tmp/storage
          create  tmp/storage/.keep
          remove  config/initializers/cors.rb
          remove  config/initializers/new_framework_defaults_6_0.rb
           rails  webpacker:install
    dev@vmdev:~$ rails new blog
          create
          create  README.md
          create  Rakefile
          create  .ruby-version
          create  config.ru
          create  .gitignore
          create  Gemfile
             run  git init from "."
    dev@vmdev:~$ ll blog/
    total 32
    drwxrwxr-x  2 dev dev 4096 Jan  3 22:40 ./
    drwxr-xr-x 20 dev dev 4096 Jan  3 22:40 ../
    -rw-rw-r--  1 dev dev  130 Jan  3 22:40 config.ru
    -rw-rw-r--  1 dev dev 1976 Jan  3 22:40 Gemfile
    -rw-rw-r--  1 dev dev  604 Jan  3 22:40 .gitignore
    -rw-rw-r--  1 dev dev  227 Jan  3 22:40 Rakefile
    -rw-rw-r--  1 dev dev  374 Jan  3 22:40 README.md
    -rw-rw-r--  1 dev dev   11 Jan  3 22:40 .ruby-version
    dev@vmdev:~$
## [3][Ruby 2.7 NEWS: Commentary by Cookpad’s Full Time Ruby Comitters](https://www.reddit.com/r/ruby/comments/ejmrz5/ruby_27_news_commentary_by_cookpads_full_time/)
- url: https://sourcediving.com/ruby-2-7-news-commentary-by-cookpads-full-time-ruby-comitters-bdbaacb36d0c
---

## [4][Load test your rack-based web apps with this shadow-requesting middleware](https://www.reddit.com/r/ruby/comments/ejgper/load_test_your_rackbased_web_apps_with_this/)
- url: https://medium.com/carwow-product-engineering/shadow-requesting-for-great-good-92cde331363a
---

## [5][UnicodePlot - plot your data on the terminal](https://www.reddit.com/r/ruby/comments/ej7p0p/unicodeplot_plot_your_data_on_the_terminal/)
- url: https://i.redd.it/0h29qquorg841.png
---

## [6][Any tools and tutorials on rails and activitypub?](https://www.reddit.com/r/ruby/comments/ejfpsh/any_tools_and_tutorials_on_rails_and_activitypub/)
- url: /r/rails/comments/ejfj4p/any_tools_and_tutorials_on_rails_and_activitypub/
---

## [7][Relevance of Ruby today aside from Rails.](https://www.reddit.com/r/ruby/comments/ej5yz9/relevance_of_ruby_today_aside_from_rails/)
- url: https://www.reddit.com/r/ruby/comments/ej5yz9/relevance_of_ruby_today_aside_from_rails/
---
I was talking to a friend recently and I told them I'm learning Ruby, to which they responded by saying Ruby has no actual market value today aside from Rails. Is that true? And Ruby Devs in the field today, what is your take on Ruby's relevance?
## [8][Ruby Conferences &amp; Camps in 2020 - What's Upcoming? Anything Missing? Updates Welcome](https://www.reddit.com/r/ruby/comments/ej2eu4/ruby_conferences_camps_in_2020_whats_upcoming/)
- url: https://www.reddit.com/r/ruby/comments/ej2eu4/ruby_conferences_camps_in_2020_whats_upcoming/
---
Hello,

   I've updated the [Ruby Conferences &amp; Camps in 2020 - What's
Upcoming?](http://planetruby.github.io/calendar/2020) page. Try the [rubyconf command line tool](https://github.com/textkit/whatson) (packaged up in the whatson gem) e.g.

    $ rubyconf

printing as of Jan/2:

    Upcoming Ruby Conferences:
    
    in 29d   Birmingham on Rails, Fri Jan/31 (1d) @ Birmingham, Alabama, United States
    in 35d   Rubyfuza, Thu-Sat Feb/6-8 (3d) @ Cape Town, South Africa
    in 47d   ParisRB Conf, Tue+Wed Feb/18+19 (2d) @ Paris, France
    in 49d   RubyConf Australia, Thu+Fri Feb/20+21 (2d) @ Melbourne, Victoria, Australia
    in 78d   Wrocław &lt;3 Ruby (wroclove.rb), Fri-Sun Mar/20-22 (3d) @ Wrocław, Poland
    in 91d   RubyDay Italy, Thu Apr/2 (1d) @ Verona, Veneto, Italy
    in 98d   RubyKaigi, Thu-Sat Apr/9-11 (3d) @ Nagano, Japan
    in 114d  RubyConf India, Sat+Sun Apr/25+26 (2d) @ Goa, India
    in 124d  RailsConf (United States), Tue-Thu May/5-7 (3d) @ Portland, Oregon, United States
    in 134d  Balkan Ruby, Fri+Sat May/15+16 (2d) @ Sofia, Bulgaria
    in 156d  Ruby Unconf Hamburg, Sat+Sun Jun/6+7 (2d) @ Hamburg, Germany
    in 183d  Brighton RubyConf, Fri Jul/3 (1d) @ Brighton, Sussex, England, United Kingdom
    in 203d  RubyConf Kenya, Thu-Sat Jul/23-25 (3d) @ Nairobi, Kenya
    in 232d  European Ruby Konference (EuRuKo), Fri+Sat Aug/21+22 (2d) @ Helsinki, Finnland
    in 288d  RubyConf Thailand (TH), Fri+Sat Oct/16+17 (2d) @ Bangkok, Thailand
    in 320d  RubyConf (United States), Tue-Thu Nov/17-19 (3d) @ Houston, Texas, United States

Anything missing? Updates welcome, see [`data/conferences2020.yml`](https://github.com/planetruby/calendar/blob/master/_data/conferences2020.yml) file
in the `planetruby/calendar` repo.

Cheers. Prosit 2020! Happy new year!
## [9][Testing Ruby on Rails application with RSpec](https://www.reddit.com/r/ruby/comments/ej4fie/testing_ruby_on_rails_application_with_rspec/)
- url: https://hixonrails.com/ruby-on-rails-tutorials/ruby-on-rails-testing-rspec-configuration/
---

## [10][Ruby 2.7 allows calling a private method with self.](https://www.reddit.com/r/ruby/comments/ejd2pi/ruby_27_allows_calling_a_private_method_with_self/)
- url: https://www.reddit.com/r/ruby/comments/ejd2pi/ruby_27_allows_calling_a_private_method_with_self/
---
[https://blog.saeloun.com/2019/12/24/ruby-2-7-allows-calling-a-private-method-with-self](https://blog.saeloun.com/2019/12/24/ruby-2-7-allows-calling-a-private-method-with-self)
