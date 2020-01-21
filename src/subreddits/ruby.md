# ruby
## [1][Only 15% of the Basecamp operations budget is spent on Ruby](https://www.reddit.com/r/ruby/comments/erj73x/only_15_of_the_basecamp_operations_budget_is/)
- url: https://m.signalvnoise.com/only-15-of-the-basecamp-operations-budget-is-spent-on-ruby/
---

## [2][Rails alongside an API gateway + lambda API for a bursty/high-volume API](https://www.reddit.com/r/ruby/comments/erpito/rails_alongside_an_api_gateway_lambda_api_for_a/)
- url: https://www.reddit.com/r/ruby/comments/erpito/rails_alongside_an_api_gateway_lambda_api_for_a/
---
Hey folks--I'm building a new product and I'm interested to get your thoughts on an architecture I had in mind. My requirements are the following:

1. A high-throughput API called from the frontend and who-knows-where else (maybe 10-100 TPS. More if things go well.)
2. Traditional website with CRUD. Would have users etc, but are infrequently accessed

The data between the two would largely independent, but the views on the Rails site would call the API. The idea is as follows:

1. A Lambda+DDB API behind APIGW using reserved concurrency
2. Traditional rails app on Heroku or EBS for the website

I've found some success using both rails and lambda in the past. Rails for traditional user-based apps and lambda for APIs. I like not having to tweak autoscaling policies on EC2 boxes for my API, plus, with the new reserved concurrency feature on lambda, I can now reach my latency goals with it. I'm wondering if I should try to combine the best of both works?

Have you folks seen anything like this in the wild? 

And, if you have, was with with EBS or Heroku? 

Do you feel like the split locations for resources might be too confusing for newer devs?

Give me your thoughts
## [3][MIR: A lightweight JIT compiler project with CRuby](https://www.reddit.com/r/ruby/comments/ereh60/mir_a_lightweight_jit_compiler_project_with_cruby/)
- url: https://developers.redhat.com/blog/2020/01/20/mir-a-lightweight-jit-compiler-project/
---

## [4][6 Advanced Ruby Loops](https://www.reddit.com/r/ruby/comments/erresb/6_advanced_ruby_loops/)
- url: https://medium.com/better-programming/6-advanced-ruby-loops-13695c20d012
---

## [5][URI.escape is obsolete. How to fix Ruby 2.7.0 warning? Percent-encoding your query string](https://www.reddit.com/r/ruby/comments/era3cq/uriescape_is_obsolete_how_to_fix_ruby_270_warning/)
- url: https://docs.knapsackpro.com/2020/uri-escape-is-obsolete-percent-encoding-your-query-string
---

## [6][Dependency management and other good programming practices.](https://www.reddit.com/r/ruby/comments/erb471/dependency_management_and_other_good_programming/)
- url: https://www.reddit.com/r/ruby/comments/erb471/dependency_management_and_other_good_programming/
---
Today I was learning how to use Nokogiri from the book instant Nokogiri. Somewhere along the line, the author mentioned that dependency management is good as a programmer especially for ruby. And then went on ahead to explain how to create the gem file and include required gems on there for bundler later on. I am still learning about programming in ruby and I realised that there may be a lot more general practices that are good but I have no knowledge of. 

What are some of the good practices you have seenor learned in your line of work that you think should be done by everyone ?
## [7][GUI on Ruby based Web Application](https://www.reddit.com/r/ruby/comments/era5k0/gui_on_ruby_based_web_application/)
- url: https://www.reddit.com/r/ruby/comments/era5k0/gui_on_ruby_based_web_application/
---
Hello Guys, 

im developing a ressource planning tool for my company. Right now im deciding between Redmine and Task Juggler. 

For that I would like to put a GUI on it. Is that possible ?

&amp;#x200B;

Thanks in adavance

Regards

&amp;#x200B;

Joshi
## [8][Be Suspicious of Join Tables](https://www.reddit.com/r/ruby/comments/er51td/be_suspicious_of_join_tables/)
- url: https://andycroll.com/ruby/be-suspicious-of-join-tables/
---

## [9][Elixir Isn't Ruby](https://www.reddit.com/r/ruby/comments/ere6k4/elixir_isnt_ruby/)
- url: https://blog.joshsoftware.com/2020/01/20/elixir-isnt-ruby/
---

## [10][Yukihiro Matsumoto: "Ruby is designed for humans, not machines"](https://www.reddit.com/r/ruby/comments/equpgz/yukihiro_matsumoto_ruby_is_designed_for_humans/)
- url: https://evrone.com/yukihiro-matsumoto-interview
---

