# ruby
## [1][ELI5: The dry Gems](https://www.reddit.com/r/ruby/comments/ghjhi4/eli5_the_dry_gems/)
- url: https://www.reddit.com/r/ruby/comments/ghjhi4/eli5_the_dry_gems/
---
I keep seeing the dry gems pop up and have read the documentation for them so I have a loose understanding of what they do.

I was hoping someone could flesh that understanding out for me a bit. Specifically, what is the use case for reaching for the set of dry gems? When do you reach for them and why? Those sorts of questions and probably more I'm not even capable of conceptualising to the point of asking them.

Thanks!
## [2][Ruby on Rails: Add `gem 'sqlite3'` to your Gemfile](https://www.reddit.com/r/ruby/comments/ghe8xd/ruby_on_rails_add_gem_sqlite3_to_your_gemfile/)
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
## [3][Assignments in Conditional Expressions](https://www.reddit.com/r/ruby/comments/ghfrau/assignments_in_conditional_expressions/)
- url: https://julienchien.com/posts/assignments-in-conditional-expressions/
---

## [4][What is a good Ruby book for an experienced non-Ruby developer?](https://www.reddit.com/r/ruby/comments/gh7mhb/what_is_a_good_ruby_book_for_an_experienced/)
- url: https://www.reddit.com/r/ruby/comments/gh7mhb/what_is_a_good_ruby_book_for_an_experienced/
---
Hi all. I was wondering if anyone could help provide some suggestions for a good Ruby book. I'm an experienced developer and have mostly worked in Java and Python over the last 10 years. Now I'm learning Ruby, but I'm finding that the online tutorials only cover the surface level syntax. I'm looking for a book that does a quick but comprehensive overview of the basics, but also focuses on teaching how to write good Ruby code and apply Ruby-isms and good Ruby conventions. Also, I'm not looking to learn Rails for now, though I'm open to a book that covers it at a high level.

Here are the three I'm debating between so far:

1. Practical Object-Oriented Design: An Agile Primer Using Ruby (2nd Edition)
2. Eloquent Ruby (1st Edition)
3. The Well-Grounded Rubyist (3rd Edition)

Do folks have any recommendations for which of these would be good for someone who has a lot of experience in other languages but is having trouble grokking Ruby? Thanks!
## [5][RSS Search Engine](https://www.reddit.com/r/ruby/comments/ghigs0/rss_search_engine/)
- url: https://github.com/davidesantangelo/feedirss-api
---

## [6][Ruby client for Factom blockchain](https://www.reddit.com/r/ruby/comments/gh4dlb/ruby_client_for_factom_blockchain/)
- url: https://github.com/kompendium-ano/factom-ruby-client
---

## [7][football.db gem family (2020.5.10 update) - zero-config (pre-packaged) open football datasets / catalogs for countries, leagues &amp; cups, clubs and more](https://www.reddit.com/r/ruby/comments/gh5wou/footballdb_gem_family_2020510_update_zeroconfig/)
- url: https://github.com/sportdb/football.db
---

## [8][Book with exercises](https://www.reddit.com/r/ruby/comments/gh6ua5/book_with_exercises/)
- url: https://www.reddit.com/r/ruby/comments/gh6ua5/book_with_exercises/
---
Hello,
I am currently trying to learn ruby with "learning Ruby" by P.Cooper, however i am missing the format of books that have exercises at the end of chapters, in order to consolidate learning of the content of the chapter.Is there any such book for Ruby?
## [9][[Another] An alternative Ruby implementation by Rust.](https://www.reddit.com/r/ruby/comments/ggy0ew/another_an_alternative_ruby_implementation_by_rust/)
- url: https://github.com/sisshiki1969/ruruby
---

## [10][Project on Software development.](https://www.reddit.com/r/ruby/comments/gh0whe/project_on_software_development/)
- url: https://www.reddit.com/r/ruby/comments/gh0whe/project_on_software_development/
---
Looking forward for some software development project to work on.
