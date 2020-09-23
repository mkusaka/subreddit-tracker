# ruby
## [1][My First open-source Project: Building AmazonPay Gem](https://www.reddit.com/r/ruby/comments/iy7m2f/my_first_opensource_project_building_amazonpay_gem/)
- url: https://www.reddit.com/r/ruby/comments/iy7m2f/my_first_opensource_project_building_amazonpay_gem/
---
I built my first open-sourc project and I wrote about it and my motivations here: [https://blog.everistus.xyz/my-first-open-source-project-building-amazonpay-gem-ckff8s5ov043o65s13ngx0pil](https://blog.everistus.xyz/my-first-open-source-project-building-amazonpay-gem-ckff8s5ov043o65s13ngx0pil)
## [2][Rails will allow a module with extend ActiveSupport::Concern to be prepended](https://www.reddit.com/r/ruby/comments/iy91jy/rails_will_allow_a_module_with_extend/)
- url: https://blog.saeloun.com/2020/09/23/prepend-class-methods-for-concerns.html
---

## [3][Triage with Me - 11 issues &amp; 2 PRs in 1.5 hours](https://www.reddit.com/r/ruby/comments/ixrqh4/triage_with_me_11_issues_2_prs_in_15_hours/)
- url: https://schneems.com/2020/09/22/triage-with-me-11-issues-2-prs-in-15-hours/
---

## [4][tty-editor - new release supports opening multiple files and searches more editor commands](https://www.reddit.com/r/ruby/comments/ixyq3n/ttyeditor_new_release_supports_opening_multiple/)
- url: https://github.com/piotrmurach/tty-editor
---

## [5][Writing Good Test Descriptions](https://www.reddit.com/r/ruby/comments/ixtuar/writing_good_test_descriptions/)
- url: https://guilhermesimoes.github.io/blog/writing-good-test-descriptions
---

## [6][Rails 6.1 deprecates rails db:structure:dump and rails db:structure:load | The Official BigBinary Blog | BigBinary](https://www.reddit.com/r/ruby/comments/ixjs28/rails_61_deprecates_rails_dbstructuredump_and/)
- url: https://blog.bigbinary.com/2020/09/22/rails-6-1-deprecates-rails-db-structure-dump.html
---

## [7][api code structure ideas](https://www.reddit.com/r/ruby/comments/ixqj0r/api_code_structure_ideas/)
- url: https://www.reddit.com/r/ruby/comments/ixqj0r/api_code_structure_ideas/
---
Hello world,
 let's imagine we have a rails app:

- 200 models 
- 400 controllers
- 300 services

and is a mess :)

I want to move to roda, but I dont know how to structure the code properly
using:

 - pg
 - pg replicas
 - elasticsearch
 - redis
 - influxdb

we have api's for:

 - clients
 - integrations
 - desktop app
 - frontend app

connections via:

 - oauth
 - token
 - sessions

api response:

 - json
 - xml

we also have:

 - jobs (200)
 - mailers(mailgun, sendgrid, default email from AR)
 - reports
 - import tools

tests that take forever

currently we do follow the rails mvc, but is a mess, boot time is slow, tests are slow

dreams?

- roda
- sequel
- ??

any ideas will be appreciated :)
## [8][Testing elasticsearch](https://www.reddit.com/r/ruby/comments/ixvm64/testing_elasticsearch/)
- url: https://www.reddit.com/r/ruby/comments/ixvm64/testing_elasticsearch/
---
So in my current project we are using elastic to index data and in specs we are doing sleep after every import and sleep function to let the import finish, which slows the whole test suite a lot. What is the best way to have the data imported to elastic when testing queries, controllers?
## [9][Ruby Courses - Advanced Ruby: Behind the Magic](https://www.reddit.com/r/ruby/comments/ix35gf/ruby_courses_advanced_ruby_behind_the_magic/)
- url: https://www.reddit.com/r/ruby/comments/ix35gf/ruby_courses_advanced_ruby_behind_the_magic/
---
It seems like biggest ruby resources are screencasts, blogs and books. And yes, they offer tons of value. 

But I’m wondering if you folks know of effective courses to learn advanced ruby/rails concepts?

Also, has anyone tried this one “Advanced Ruby: Behind the Magic”

https://courses.gorails.com/advanced-ruby-for-rails-devs

(To me, 300 bucks seems to be on the higher end for 30 screencasts..)
## [10][Any downsides storing database host/username/password in ruby ENV?](https://www.reddit.com/r/ruby/comments/ixazdi/any_downsides_storing_database/)
- url: https://www.reddit.com/r/ruby/comments/ixazdi/any_downsides_storing_database/
---
I have this code

YAML.load(File.open(env\_file))\[ENV\["RACK\_ENV"\]\].each { |key, value| ENV\[key.to\_s\] = value } if File.exists?(env\_file)

This file loads all the secret variables, client API secrets, db secrets and everything into ENV.

&amp;#x200B;

Then, different gems just call ENV\['aa\_secret\_key'\] to get the secret and boot up the app

secret\_env\_file.yml is not uploaded on source control. So, everything stays safe.

&amp;#x200B;

Are there any downsides of using ENV to store all those information while the app boots up? I think ENV\['anything\_password'\] is still accessible later on as long as the app is still up and running. Like for sinatra/rails etc.
