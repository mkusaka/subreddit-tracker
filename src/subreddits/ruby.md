# ruby
## [1][render_async 2.1.6 released](https://www.reddit.com/r/ruby/comments/gi8xuf/render_async_216_released/)
- url: https://pragmaticpineapple.com/render-async-2-1-6-released/
---

## [2][TruffleRuby is now the first implementation besides CRuby to pass all of the 2180 #RubyGems' tests, no exclusions!](https://www.reddit.com/r/ruby/comments/giabdw/truffleruby_is_now_the_first_implementation/)
- url: https://www.reddit.com/r/ruby/comments/giabdw/truffleruby_is_now_the_first_implementation/
---
[https://github.com/rubygems/rubygems/pull/2797#issuecomment-626150446](https://github.com/rubygems/rubygems/pull/2797#issuecomment-626150446)

&amp;#x200B;

Hopefully we are much much closer to running Rails. 
## [3][sportdb-readers v1.0 - sport.db readers for leagues, seasons, clubs, match schedules and results, and more](https://www.reddit.com/r/ruby/comments/gi885t/sportdbreaders_v10_sportdb_readers_for_leagues/)
- url: https://github.com/sportdb/sport.db/tree/master/sportdb-readers
---

## [4][I'm creating a Ruby on Rails 6 course. What would you like to see it it?](https://www.reddit.com/r/ruby/comments/gi9wcg/im_creating_a_ruby_on_rails_6_course_what_would/)
- url: https://www.reddit.com/r/ruby/comments/gi9wcg/im_creating_a_ruby_on_rails_6_course_what_would/
---
I don't want to focus on the basic things like "what is Ruby, what is Rails, rails new, scaffolding, devise", - rather more advanced and interesting topics. What would you like to hear/learn more about?

[View Poll](https://www.reddit.com/poll/gi9wcg)
## [5][Books or video tutorials on how to use dry-CLI?](https://www.reddit.com/r/ruby/comments/gi6jlf/books_or_video_tutorials_on_how_to_use_drycli/)
- url: https://www.reddit.com/r/ruby/comments/gi6jlf/books_or_video_tutorials_on_how_to_use_drycli/
---
Looking to expand the things I know about dry in ruby, starting with dry-CLI. Any direction would be of help. Thanks.
## [6][Why hasn't Opal Native taken off?](https://www.reddit.com/r/ruby/comments/ghs1zm/why_hasnt_opal_native_taken_off/)
- url: https://www.reddit.com/r/ruby/comments/ghs1zm/why_hasnt_opal_native_taken_off/
---
https://github.com/zetachang/opal-native

While frameworks like hyperstack and volt let you write ruby on the frontend using opal, there seems to me to exist a use case to develop mobile apps quickly from a rails (or similar) backend, using only ruby on the backend and mobile fronted. Any ideas why this hasn't taken off? Currently if we want to develop mobile apps using ruby we have to start from scratch with RubyMotion. A monolithic full-ruby stack sounds very appealing to me.

Edit: maybe I should have used a different title. I’m not so fussed about this Opal Native repo in particular, rather I’m wondering why there isn’t a ruby framework to develop cross platform apps. There’s clearly a need for it considering what’s going on in the JS world, but ruby has significant advantages over JS, one of them being rapid web app development.
## [7][Adding an exclamation point to each element in an array](https://www.reddit.com/r/ruby/comments/ghve61/adding_an_exclamation_point_to_each_element_in_an/)
- url: https://www.reddit.com/r/ruby/comments/ghve61/adding_an_exclamation_point_to_each_element_in_an/
---
Hi!

I'm trying to add an exclamation point to each element in an array. Not to `.join` them, but to literally add ! to each element. For example, my array is `["a","b","c"]` and I want it to be `["a!", "b!", "c!"]`

Is there a Ruby enumerable to do this? I've surfed the web but it keeps leading me back to `bang`, which isn't what I want.
## [8][ELI5: The dry Gems](https://www.reddit.com/r/ruby/comments/ghjhi4/eli5_the_dry_gems/)
- url: https://www.reddit.com/r/ruby/comments/ghjhi4/eli5_the_dry_gems/
---
I keep seeing the dry gems pop up and have read the documentation for them so I have a loose understanding of what they do.

I was hoping someone could flesh that understanding out for me a bit. Specifically, what is the use case for reaching for the set of dry gems? When do you reach for them and why? Those sorts of questions and probably more I'm not even capable of conceptualising to the point of asking them.

Thanks!
## [9][Ruby on Rails: Add `gem 'sqlite3'` to your Gemfile](https://www.reddit.com/r/ruby/comments/ghe8xd/ruby_on_rails_add_gem_sqlite3_to_your_gemfile/)
- url: https://www.reddit.com/r/ruby/comments/ghe8xd/ruby_on_rails_add_gem_sqlite3_to_your_gemfile/
---
 

I've been working on a project for a term paper over MVC architecture and have been trying to set up the blog template suggested by 'Getting Started with Rails' webpage ( [https://guides.rubyonrails.org/getting\_started.html](https://guides.rubyonrails.org/getting_started.html) ). After fighting with tireless amounts of version issues, I have finally got the server starting through the command, 'ruby bin\\rails server'. However, now the Localhost:3000 is providing an error of:

Specified 'sqlite3' for database adapter, but the gem is not loaded. Add gem 'sqlite3' to your Gemfile (and ensure its version is at the minimum required by ActiveRecord).

I have tried all suggestions from the thread listed here ( [https://stackoverflow.com/questions/17350837/ruby-on-rails-add-gem-sqlite3-to-your-gemfile](https://stackoverflow.com/questions/17350837/ruby-on-rails-add-gem-sqlite3-to-your-gemfile) ), and here ( [http://www.ruby-forum.com/topic/4415126](http://www.ruby-forum.com/topic/4415126) ) to no avail. After seeing countless mentions about switching from Windows to Linux, I did. But that only proved to be more of a headache with more version issues, etc. I am on a Windows 10 OS and would greatly appreciate any suggestions or help.

&amp;#x200B;

Here is my Gemfile for reference:

source '[https://rubygems.org](https://rubygems.org)'

&amp;#x200B;

`git_source(:github) do |repo_name|`

  `repo_name = "#{repo_name}/#{repo_name}" unless repo_name.include?("/")`

  `"`[`https://github.com/#{repo_name}.git`](https://github.com/#{repo_name}.git)`"`

`end`

&amp;#x200B;

&amp;#x200B;

`# Bundle edge Rails instead: gem 'rails', github: 'rails/rails'`

`gem 'rails', '~&gt; 5.0.7', '&gt;=` [`5.0.7.1`](https://5.0.7.1)`'`

`# Use sqlite3 as the database for Active Record`

`gem 'sqlite3'`

`# Use Puma as the app server`

`gem 'puma', '~&gt; 3.0'`

`# Use SCSS for stylesheets`

`gem 'sass-rails', '~&gt; 5.0'`

`# Use Uglifier as compressor for JavaScript assets`

`gem 'uglifier', '&gt;= 1.3.0'`

`# Use CoffeeScript for .coffee assets and views`

`gem 'coffee-rails', '~&gt; 4.2'`

`# See` [`https://github.com/rails/execjs#readme`](https://github.com/rails/execjs#readme) `for more supported runtimes`

`# gem 'therubyracer', platforms: :ruby`

&amp;#x200B;

`# Use jquery as the JavaScript library`

`gem 'jquery-rails'`

`# Turbolinks makes navigating your web application faster. Read more:` [`https://github.com/turbolinks/turbolinks`](https://github.com/turbolinks/turbolinks)

`gem 'turbolinks', '~&gt; 5'`

`# Build JSON APIs with ease. Read more:` [`https://github.com/rails/jbuilder`](https://github.com/rails/jbuilder)

`gem 'jbuilder', '~&gt; 2.5'`

`# Use Redis adapter to run Action Cable in production`

`# gem 'redis', '~&gt; 3.0'`

`# Use ActiveModel has_secure_password`

`# gem 'bcrypt', '~&gt; 3.1.7'`

&amp;#x200B;

`# Use Capistrano for deployment`

`# gem 'capistrano-rails', group: :development`

&amp;#x200B;

`group :development, :test do`

  `# Call 'byebug' anywhere in the code to stop execution and get a debugger console`

  `gem 'byebug', platform: :mri`

`end`

&amp;#x200B;

`group :development do`

  `# Access an IRB console on exception pages or by using &lt;%= console %&gt; anywhere in the code.`

  `gem 'web-console', '&gt;= 3.3.0'`

`end`

&amp;#x200B;

`# Windows does not include zoneinfo files, so bundle the tzinfo-data gem`

`gem 'tzinfo-data', platforms: [:mingw, :mswin, :x64_mingw, :jruby]`
## [10][RSS Search Engine](https://www.reddit.com/r/ruby/comments/ghigs0/rss_search_engine/)
- url: https://github.com/davidesantangelo/feedirss-api
---

