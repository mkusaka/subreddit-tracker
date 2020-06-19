# ruby
## [1][Web Scraping with Ruby](https://www.reddit.com/r/ruby/comments/hbgajf/web_scraping_with_ruby/)
- url: https://www.scrapingbee.com/blog/web-scraping-ruby/
---

## [2][Maintenance Mode and End of Support Dates Announced for AWS SDK For Ruby V2](https://www.reddit.com/r/ruby/comments/hbsf1b/maintenance_mode_and_end_of_support_dates/)
- url: https://aws.amazon.com/blogs/developer/deprecation-schedule-for-aws-sdk-for-ruby-v2/
---

## [3][Rails 6.0.3.2 has been released!](https://www.reddit.com/r/ruby/comments/hbsilw/rails_6032_has_been_released/)
- url: https://weblog.rubyonrails.org/2020/6/17/Rails-6-0-3-2-has-been-released/
---

## [4][Looking into learning Ruby](https://www.reddit.com/r/ruby/comments/hbmy6e/looking_into_learning_ruby/)
- url: https://www.reddit.com/r/ruby/comments/hbmy6e/looking_into_learning_ruby/
---
Hey guys! I’m a full stack developer who’s coming from a JavaScript/Node/Express background. And I’m intrested in learning ruby, but a lot of the medium articles are 3+ years old. Would those articles still be relevant?
If not then any good resources?
## [5][Alternative to Sidekiq for long-running jobs](https://www.reddit.com/r/ruby/comments/hbi3gv/alternative_to_sidekiq_for_longrunning_jobs/)
- url: https://www.reddit.com/r/ruby/comments/hbi3gv/alternative_to_sidekiq_for_longrunning_jobs/
---
Hello everyone,

i'm developing a pretty old ruby (on rails) app which relies heavy on background processing using Sidekiq. We basically fetch a lot of data from external CRM systems, parses them and persists on our side. As the job sounds relatively easy from the logic perspective, it became a nightmare from the infrastructure/resources/timing perspective. Some CRMs we use have an API designed in a way that  doing one import run might even take up to **20 hours (!!).** As sidekiq has been designed to run a fast and simple jobs, we started to look for alternatives.

I'm not entirely sure what am i looking for to be honest but ideal solution that comes to my mind is: some sort of a cloud service i can send a HTTP request to with any params. This service receives it and runs any script i configured injecting the query params into the "job". Being it docker-based would be additional pro. Even better if i could clone any git repository (containing my ruby script) inside that job. It kinda reminds me of how CircleCI works - you can setup a project, assign github repo to it, prepare a  "workflow" with anything you want inside and trigger a build using a webhook (not sure about that part but you get the idea).

I even thought about setting up a [Drone.io](https://Drone.io) instance but i would prefer to have something in cloud (cloud is my strong requirement - we can't afford to maintain anything complex on our own)

Thanks in advance for any ideas!

PS. i beg you not to advice me to split my long running sidekiq jobs into smaller ones, we already tried that - let's focus on cloud-based solutions :)
## [6][Is there a full list of structs, unions and functions of the C API?](https://www.reddit.com/r/ruby/comments/hbdjd8/is_there_a_full_list_of_structs_unions_and/)
- url: https://www.reddit.com/r/ruby/comments/hbdjd8/is_there_a_full_list_of_structs_unions_and/
---
I am looking for a complete list of the \`struct\`s, \`typedef struct\`s, and functions in Ruby C API and it's body.  


I tried looking up but I found nothing.
## [7][How to use a string array for case labels and call methods with same name](https://www.reddit.com/r/ruby/comments/hbc4uo/how_to_use_a_string_array_for_case_labels_and/)
- url: https://www.reddit.com/r/ruby/comments/hbc4uo/how_to_use_a_string_array_for_case_labels_and/
---
Hi All -

Can't quite figure this one out, hoping the /r/ruby brains trust can help me out.

I have code similar to this:

    cmd = get\_string\_from\_requestor
    case cmd
    when 'FOO' then foo # foo() is a method on the same class
    when 'BAR' then bar # as is bar()
    when 'BAZ' then bar # etc
    else
        # error
    end

And I was hoping to switch it to something a bit more dynamic:

    cmd = get\_string\_from\_requestor
    commands = %w(FOO BAR BAZ QUX QUUX)
    case cmd
    # magic required here: iterate(?) through commands[], match cmd
    # to the commands[i] and then call commands[i].downcase as a
    # method on current class...?
    else
        # error
    end

Anyone point me in the right direction?

Thanks!
## [8][I wanna make fancy(!) TUIs for Linux using Ruby. Are my only options the ancient and largely undocumented ncurses wrappers?](https://www.reddit.com/r/ruby/comments/hb4uqq/i_wanna_make_fancy_tuis_for_linux_using_ruby_are/)
- url: https://www.reddit.com/r/ruby/comments/hb4uqq/i_wanna_make_fancy_tuis_for_linux_using_ruby_are/
---
Every time I've tried to use one of the many ncurses wrappers, there's like one blog post and if you're lucky some docs written for C programmers.

I've tried reading the 'sup' source and it wasn't really helpful. Is the windowed TUI the stuff of wizards in 2020?
## [9][An attempt at creating a "Heroku-like" Docker/Helm/GitHub Action setup for "classic-stack" Rails apps. Feedback wanted.](https://www.reddit.com/r/ruby/comments/haz3c0/an_attempt_at_creating_a_herokulike/)
- url: https://github.com/lewagon/rails-k8s-demo
---

## [10][Rails 6.1's ActiveModel Errors Revamp](https://www.reddit.com/r/ruby/comments/hap4u9/rails_61s_activemodel_errors_revamp/)
- url: https://code.lulalala.com/2020/0531-1013.html
---

