# ruby
## [1][Need study buddy](https://www.reddit.com/r/ruby/comments/iopopr/need_study_buddy/)
- url: https://www.reddit.com/r/ruby/comments/iopopr/need_study_buddy/
---
Hi there! I am currently learning Ruby, solving challenges on codewars kyu 7-5. Looking for a study buddy to do live sessions to practice coding in front of someone else. Language doesn’t really matter, it’s more about practicing coding in an ‘interview-like’ manner.

DM me to arrange a session.
## [2][my deployment of rails api to digitalocean is failing. can someone point me in a direction to hunt down why?](https://www.reddit.com/r/ruby/comments/iorexx/my_deployment_of_rails_api_to_digitalocean_is/)
- url: https://www.reddit.com/r/ruby/comments/iorexx/my_deployment_of_rails_api_to_digitalocean_is/
---
i've been trying to deploy my rails api (to be paired with a react front end) to digitalocean and have been getting stuck at the same 2 spots and i'm just not really sure why. i feel like there's something obvious i'm missing, but i don't think i know enough to spot it.

i've been following this guide: https://gorails.com/deploy/ubuntu/20.04#guide

i get this error:

    00:07 deploy:assets:precompile
    01 $HOME/.rbenv/bin/rbenv exec bundle exec rake assets:precompile
    01 rake aborted!
    01 URI::InvalidURIError: bad URI(is not URI?): postgresql://deploy:nicetry@127.0.0.1/my-api-prod

is there an issue with me using rails as API only mode? If so, I tried this according to a github issue I found:

in capfile you add these lines (but it fails on migration after getting thru precompile) -

    # "require 'capistrano/rails'"

    require 'capistrano/bundler'
    require 'capistrano/rails/migrations'

i went thru a bootcamp but they never touched on deploying and only suggested using heroku, so i'm curious if i just didn't learn something obvious. i've spent a ton of time on trying to get it to work but at this point i want to know what i'm not understanding before just using an easy setup deployment service.

thanks for any help!
## [3][GitHub Actions](https://www.reddit.com/r/ruby/comments/io6g0t/github_actions/)
- url: https://www.driftingruby.com/episodes/github-actions?utm_medium=social&amp;utm_campaign=weekly_episode&amp;utm_source=reddit
---

## [4][backup - a github repo backup command line tool (powered by gitti/backup)](https://www.reddit.com/r/ruby/comments/iobrxj/backup_a_github_repo_backup_command_line_tool/)
- url: https://github.com/rubycoco/gitti/tree/master/gitti-backup
---

## [5][Help to optimize a method that gets called frequently.](https://www.reddit.com/r/ruby/comments/ioa2em/help_to_optimize_a_method_that_gets_called/)
- url: https://www.reddit.com/r/ruby/comments/ioa2em/help_to_optimize_a_method_that_gets_called/
---
There is a method that gets called for every request in my rails application and I optimized it to cut a few iterations inside a loop on every method call. The thing is I'm not sure how much of a better code this is.  I ran some benchmarks on the newer and existing code and this is how the data looks 

    10M iterations
                    user     system      total        real
    Exisintg =&gt; 55.180000   0.930000  56.110000 ( 56.397975)
     New =&gt; 46.250000   0.570000  46.820000 ( 46.973160)

    20M Iterations
                    user     system      total        real
    Exising =&gt; 105.620000   1.700000 107.320000 (108.221156)
    New =&gt;.     87.840000   1.000000  88.840000 ( 88.970525)

I used the [Benchmark.bm](https://Benchmark.bm) method to measure this data. What I'm trying to understand is how can I processed to more accurately find the performance improvements so I can go ahead with this approach.   
Things I'm worried about. 

1. the accuracy of bm isn't the best and wouldn't translate to real-world systems and I'm not able to get it to work with bmbm which is more accurate than this. 
2. Have to justify the testing and development effort is worth the gains. 
3. Does shaving a few iterations for every request actually worth it. For an individual request, I would say no but when a system receives say 30M requests it would increase the throughput of the system right? even if the average response time for a call doesn't go down by much?

Can anyone give me more insight into what I need to do to move forward with this?
## [6][retest: a gem to help you refactor your ruby projects faster](https://www.reddit.com/r/ruby/comments/io0uae/retest_a_gem_to_help_you_refactor_your_ruby/)
- url: https://github.com/AlexB52/retest
---

## [7][Survey finds only 3% of Ruby on Rails developers use Windows](https://www.reddit.com/r/ruby/comments/inpgox/survey_finds_only_3_of_ruby_on_rails_developers/)
- url: https://developers.slashdot.org/story/20/09/06/0028214/survey-finds-only-3-of-ruby-on-rails-developers-use-windows
---

## [8][HashWithIndifferentAccess in Rails](https://www.reddit.com/r/ruby/comments/io1trg/hashwithindifferentaccess_in_rails/)
- url: https://www.sandipmane.dev/hash-with-indifferent-access-in-rails
---

## [9][Command-line Arguments in Ruby: Part I](https://www.reddit.com/r/ruby/comments/inmgds/commandline_arguments_in_ruby_part_i/)
- url: https://medium.com/rubycademy/command-line-arguments-in-ruby-part-i-8a89479cb70f
---

## [10][Integrate Doorkeeper gem – API only ruby on rails course (chapter 10)](https://www.reddit.com/r/ruby/comments/inne5q/integrate_doorkeeper_gem_api_only_ruby_on_rails/)
- url: https://duetcode.io/rails-api-only-course/integrate-doorkeeper-gem
---

