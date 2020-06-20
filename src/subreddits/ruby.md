# ruby
## [1][Library to write, in pure Ruby, web apps that are automagically accessible from all the Internet (link to source code and online demos in comment)](https://www.reddit.com/r/ruby/comments/hch8e6/library_to_write_in_pure_ruby_web_apps_that_are/)
- url: https://i.redd.it/n9wu9gsrd0651.gif
---

## [2][Ruby Notes for Professionals "by the beautiful people at Stack Overflow"](https://www.reddit.com/r/ruby/comments/hc8btp/ruby_notes_for_professionals_by_the_beautiful/)
- url: https://www.dbooks.org/ruby-notes-for-professionals-5592543428/read/
---

## [3][Rails Vs Sinatra](https://www.reddit.com/r/ruby/comments/hcgsb7/rails_vs_sinatra/)
- url: https://www.reddit.com/r/ruby/comments/hcgsb7/rails_vs_sinatra/
---
Hey all!

I'm a long time ruby dev who's never gone into rails or Sinatra(mostly been on the Ops side), I've read some very old threads explaining that as people have scaled they've moved away from Sinatra into rails, but none of them actually go into much detail. Does anyone here have experience with reasons to migrate to rails? Thus far I've had zero problems with Sinatra so I intend on using it but am wondering what I may run into down the line.

Appreciate any advice!
## [4][Deep dive into Database transactions in RSpec](https://www.reddit.com/r/ruby/comments/hcfb0w/deep_dive_into_database_transactions_in_rspec/)
- url: https://www.reddit.com/r/ruby/comments/hcfb0w/deep_dive_into_database_transactions_in_rspec/
---
[https://medium.com/swlh/deep-dive-into-database-transactions-in-rspec-a3d883b35c0?source=friends\_link&amp;sk=257ed6980f4a6ea1d9ba2e34b6a1a071](https://medium.com/swlh/deep-dive-into-database-transactions-in-rspec-a3d883b35c0?source=friends_link&amp;sk=257ed6980f4a6ea1d9ba2e34b6a1a071)

&amp;#x200B;
## [5][Rails 6.0.3.2 has been released!](https://www.reddit.com/r/ruby/comments/hbsilw/rails_6032_has_been_released/)
- url: https://weblog.rubyonrails.org/2020/6/17/Rails-6-0-3-2-has-been-released/
---

## [6][Maintenance Mode and End of Support Dates Announced for AWS SDK For Ruby V2](https://www.reddit.com/r/ruby/comments/hbsf1b/maintenance_mode_and_end_of_support_dates/)
- url: https://aws.amazon.com/blogs/developer/deprecation-schedule-for-aws-sdk-for-ruby-v2/
---

## [7][Web Scraping with Ruby](https://www.reddit.com/r/ruby/comments/hbgajf/web_scraping_with_ruby/)
- url: https://www.scrapingbee.com/blog/web-scraping-ruby/
---

## [8][Looking into learning Ruby](https://www.reddit.com/r/ruby/comments/hbmy6e/looking_into_learning_ruby/)
- url: https://www.reddit.com/r/ruby/comments/hbmy6e/looking_into_learning_ruby/
---
Hey guys! I’m a full stack developer who’s coming from a JavaScript/Node/Express background. And I’m intrested in learning ruby, but a lot of the medium articles are 3+ years old. Would those articles still be relevant?
If not then any good resources?
## [9][Alternative to Sidekiq for long-running jobs](https://www.reddit.com/r/ruby/comments/hbi3gv/alternative_to_sidekiq_for_longrunning_jobs/)
- url: https://www.reddit.com/r/ruby/comments/hbi3gv/alternative_to_sidekiq_for_longrunning_jobs/
---
Hello everyone,

i'm developing a pretty old ruby (on rails) app which relies heavy on background processing using Sidekiq. We basically fetch a lot of data from external CRM systems, parses them and persists on our side. As the job sounds relatively easy from the logic perspective, it became a nightmare from the infrastructure/resources/timing perspective. Some CRMs we use have an API designed in a way that  doing one import run might even take up to **20 hours (!!).** As sidekiq has been designed to run a fast and simple jobs, we started to look for alternatives.

I'm not entirely sure what am i looking for to be honest but ideal solution that comes to my mind is: some sort of a cloud service i can send a HTTP request to with any params. This service receives it and runs any script i configured injecting the query params into the "job". Being it docker-based would be additional pro. Even better if i could clone any git repository (containing my ruby script) inside that job. It kinda reminds me of how CircleCI works - you can setup a project, assign github repo to it, prepare a  "workflow" with anything you want inside and trigger a build using a webhook (not sure about that part but you get the idea).

I even thought about setting up a [Drone.io](https://Drone.io) instance but i would prefer to have something in cloud (cloud is my strong requirement - we can't afford to maintain anything complex on our own)

Thanks in advance for any ideas!

PS. i beg you not to advice me to split my long running sidekiq jobs into smaller ones, we already tried that - let's focus on cloud-based solutions :)
## [10][Is there a full list of structs, unions and functions of the C API?](https://www.reddit.com/r/ruby/comments/hbdjd8/is_there_a_full_list_of_structs_unions_and/)
- url: https://www.reddit.com/r/ruby/comments/hbdjd8/is_there_a_full_list_of_structs_unions_and/
---
I am looking for a complete list of the \`struct\`s, \`typedef struct\`s, and functions in Ruby C API and it's body.  


I tried looking up but I found nothing.
