# ruby
## [1][Understanding Rails secrets/credentials](https://www.reddit.com/r/ruby/comments/f33dby/understanding_rails_secretscredentials/)
- url: https://www.codewithjason.com/understanding-rails-secrets-credentials/
---

## [2][Rails introduces disallowed deprecations in ActiveSupport](https://www.reddit.com/r/ruby/comments/f3580f/rails_introduces_disallowed_deprecations_in/)
- url: https://blog.saeloun.com/2020/02/12/rails-active-support-disallowed-deprecations
---

## [3][Opening The Ruby Concurrency Toolbox](https://www.reddit.com/r/ruby/comments/f300tx/opening_the_ruby_concurrency_toolbox/)
- url: https://www.honeybadger.io/blog/ruby-concurrency-parallelism/
---

## [4][I'm trying to scrap this json file with nokogiri](https://www.reddit.com/r/ruby/comments/f37h5s/im_trying_to_scrap_this_json_file_with_nokogiri/)
- url: https://www.reddit.com/r/ruby/comments/f37h5s/im_trying_to_scrap_this_json_file_with_nokogiri/
---
The webpage I'm trying to scrape is this one [here](https://webfec.org.br/Utils/GetCentrobyId?cod=1). The data I want is between the strong tag and the rest between labels tag.

Here my current code which is returning an array empty

require 'nokogiri'  
require 'rest-client'  
html = RestClient.get('https://webfec.org.br/Utils/GetCentrobyId?cod=1')  
doc = Nokogiri::HTML.parse(html)  
p emails = doc.xpath("/table/tbody/tr/td/div/span/strong").text
## [5][best ruby api there!](https://www.reddit.com/r/ruby/comments/f38k3n/best_ruby_api_there/)
- url: https://rubyapi.org/
---

## [6][Minimalist gem wrapper for Amazon Product Advertising API 5.0 Cart Form functionality](https://www.reddit.com/r/ruby/comments/f2zcfe/minimalist_gem_wrapper_for_amazon_product/)
- url: https://github.com/skatkov/carriage
---

## [7][Ruby API: Easily Find Ruby documentation](https://www.reddit.com/r/ruby/comments/f2rk50/ruby_api_easily_find_ruby_documentation/)
- url: http://rubyapi.org/
---

## [8][Building a Sumo Logic Dashboard for a Rails App on Heroku](https://www.reddit.com/r/ruby/comments/f2v2ew/building_a_sumo_logic_dashboard_for_a_rails_app/)
- url: https://scottbartell.com/2020/02/12/building-sumologic-dashboard-for-rails-heroku-app/
---

## [9][How to Calculate Tech Debt Using Skunk on GitHub Actions](https://www.reddit.com/r/ruby/comments/f2tjzd/how_to_calculate_tech_debt_using_skunk_on_github/)
- url: https://www.fastruby.io/blog/code-quality/calculate-tech-debt-using-skunk-on-github-actions.html
---

## [10][Getting Started With System Tests in Rails With Minitest](https://www.reddit.com/r/ruby/comments/f2qmkh/getting_started_with_system_tests_in_rails_with/)
- url: https://blog.appsignal.com/2020/02/12/getting-started-with-system-tests-in-ruby-with-minitest.html
---

