# ruby
## [1][JRuby 9.2.10.0 released!](https://www.reddit.com/r/ruby/comments/f605sa/jruby_92100_released/)
- url: https://www.jruby.org/2020/02/18/jruby-9-2-10-0.html
---

## [2][ROM + Dry Showcase: Part 3 - Testing](https://www.reddit.com/r/ruby/comments/f67vk3/rom_dry_showcase_part_3_testing/)
- url: https://ryanbigg.com/2020/02/rom-and-dry-showcase-part-3
---

## [3][Python(ish) to Ruby. Any suggested books/resources?](https://www.reddit.com/r/ruby/comments/f66yjc/pythonish_to_ruby_any_suggested_booksresources/)
- url: https://www.reddit.com/r/ruby/comments/f66yjc/pythonish_to_ruby_any_suggested_booksresources/
---
Hello everyone! I have made it a professional goal to learn Ruby this year. I have a short background in data management with a little bit of dabbling in python. I have been suggested \_why's poignant guide from two coworkers, both of which mentioned it is outdated but still very good for learning. In addition, it's what they both used when learning Ruby.  

I'm not completely new to programming: I've written a few object-oriented python scripts and whatnot. I'm just curious what you folks used when you first started and what sort of resources/underlying concepts on which I should focus. My main goals are to transpose my Python scripts into Ruby. They include API calls and web scraping. 

Thanks and if I have broken a rule by posting this then please send this post to oblivion (also I apologize).
## [4][How I MITM'd rubygems.org ... Kinda](https://www.reddit.com/r/ruby/comments/f64rav/how_i_mitmd_rubygemsorg_kinda/)
- url: http://gavinmiller.io/2020/how-i-mitmd-rubygems-org-kinda/
---

## [5][Migrating user passwords from Django to Ruby](https://www.reddit.com/r/ruby/comments/f66teg/migrating_user_passwords_from_django_to_ruby/)
- url: http://jetrockets.pro/blog/migrating-user-passwords-from-django-to-ruby
---

## [6][Decoupling Ruby: Delegation vs Dependency Injection](https://www.reddit.com/r/ruby/comments/f5ysjs/decoupling_ruby_delegation_vs_dependency_injection/)
- url: https://www.reddit.com/r/ruby/comments/f5ysjs/decoupling_ruby_delegation_vs_dependency_injection/
---
We've all worked with tightly-coupled code. If a butterfly flaps its wings in China, the unit tests break. Maintaining a system like this is...unpleasant. In this article, Jonathan Miles dives into the origins of tight-coupling. He demonstrates how you can use dependency injection (DI) to decouple code. Then he introduces a novel decoupling technique based on delegation that can be useful when DI is not an option. [https://www.honeybadger.io/blog/decoupling-ruby-delegation-dependency-injection/](https://www.honeybadger.io/blog/decoupling-ruby-delegation-dependency-injection/)
## [7][awesome-stimulusjs: List of modestly-awesome StimulusJS stuff](https://www.reddit.com/r/ruby/comments/f5r2re/awesomestimulusjs_list_of_modestlyawesome/)
- url: https://github.com/skatkov/awesome-stimulusjs
---

## [8][Ruby 2.7 removes taint checking mechanism](https://www.reddit.com/r/ruby/comments/f67vj6/ruby_27_removes_taint_checking_mechanism/)
- url: https://blog.saeloun.com/2020/02/18/ruby-2-7-access-and-setting-of-safe-warned-will-become-global-variable
---

## [9][Image classification in ruby](https://www.reddit.com/r/ruby/comments/f5rbkt/image_classification_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/f5rbkt/image_classification_in_ruby/
---
Anyone have any experience with image classification in ruby?

I want to add it Image classification to an existing rails app but i’m not sure if i should just make a micro service in a different language or is there a good solution in ruby?
## [10][Approach for associated model searching with Elasticsearch [Ruby, ActiveRecord, elasticsearch-model]](https://www.reddit.com/r/ruby/comments/f5zmxu/approach_for_associated_model_searching_with/)
- url: https://www.reddit.com/r/ruby/comments/f5zmxu/approach_for_associated_model_searching_with/
---
Hi community, I have a question about choosing the best approach in case of searching.

I have Ruby off Rails API, with `ActiveRecord`, together with `Elasticsearch` ( `elasticsearch-model` gem).

It's a simple API which returns `projects` (Project AR model), where I set up indexes:

`mapping dynamic: false do`  
`indexes :created_at, type: 'date'`  
 `end`

Then I simply make a search on `Elasticsearch` and return `AR` relations directly from the controller. It's working perfectly.

Now I am trying to add categories to projects, as `categories has_many projects`, and projects belongs\_to categories. I am wondering about two things:

1. How now should I make a query to get projects from specific categories, should I reimplement it to return `result = Category.search(...)` and return result.jobs, or still seeking by `projects`, but searching by `category_id`?
2. How to combine Category and Project in Elasticsearch to make possible to search for projects from a specific category, and from various multiple categories? Merge mappings?

Thanks in advance!
