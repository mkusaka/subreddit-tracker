# ruby
## [1][Gem API design question](https://www.reddit.com/r/ruby/comments/ipce8i/gem_api_design_question/)
- url: https://www.reddit.com/r/ruby/comments/ipce8i/gem_api_design_question/
---
Hi,

I'm new to the Ruby world, and a project at work is using a gem implementing a wrapper to a remote service API. What troubles me is that this gem is exposing only class-level methods, including the "configure" method which does *not* return an object instance.

Its usage is something like this :

```ruby  
ZooApiWrapper.configure { |config|
  config.token = "abcdef1234"
  config.url = "https://my.api/"
}

lions = ZooApiWrapper.animals.lions

new_pet = ZooApiWrapper::Animal::Dog.new
new_pet.name = "Spot"
new_pet.save

# etc...
```

I would have expected to have the configure call return a wrapper object, and use this wrapper object to access the various resources exposed via the API. Is it usual to use global state for such things instead in Ruby ?
## [2][The difference between let, let! and instance variables in RSpec](https://www.reddit.com/r/ruby/comments/ip0se6/the_difference_between_let_let_and_instance/)
- url: https://www.codewithjason.com/difference-let-let-instance-variables-rspec/
---

## [3][Writing a Ractor-based web server](https://www.reddit.com/r/ruby/comments/iozx5h/writing_a_ractorbased_web_server/)
- url: https://kirshatrov.com/2020/09/08/ruby-ractor-web-server/
---

## [4][Rails 6.1 adds --minimal option support](https://www.reddit.com/r/ruby/comments/iovpbj/rails_61_adds_minimal_option_support/)
- url: https://blog.bigbinary.com/2020/09/08/rails-6-1-adds-minimal-option-support.html
---

## [5][Working Set - a new tool for your programmer workbench. Written in ruby with actor-model concurrency, unix sockets for IPC, and an ncurses UI. Integrates with Vim, other editors possible.](https://www.reddit.com/r/ruby/comments/ip0b1e/working_set_a_new_tool_for_your_programmer/)
- url: https://github.com/coderifous/working_set
---

## [6][Introducing Eventide Fixtures: Testing So Easy It Feels Like Cheating](https://www.reddit.com/r/ruby/comments/iow78t/introducing_eventide_fixtures_testing_so_easy_it/)
- url: https://blog.eventide-project.org/articles/introducing-eventide-fixtures/
---

## [7][Integrating Chargebee subscriptions and recurring billing with a Rails app](https://www.reddit.com/r/ruby/comments/ioym5b/integrating_chargebee_subscriptions_and_recurring/)
- url: https://vitobotta.com/2020/09/08/integrating-chargebee-subscriptions-and-recurring-billing-with-a-rails-app/
---

## [8][New to ruby. I've been stuck on the coding part for a minute. I've been following this guys tutorial, following the exact same coding. His works but on my end, I get the 204 error code. In my sublime text, can someone help me with this coding? Post is not redirecting to my show controller.](https://www.reddit.com/r/ruby/comments/ip5j4q/new_to_ruby_ive_been_stuck_on_the_coding_part_for/)
- url: https://www.reddit.com/gallery/ip5j4q
---

## [9][my deployment of rails api to digitalocean is failing. can someone point me in a direction to hunt down why?](https://www.reddit.com/r/ruby/comments/iorexx/my_deployment_of_rails_api_to_digitalocean_is/)
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
## [10][Need study buddy](https://www.reddit.com/r/ruby/comments/iopopr/need_study_buddy/)
- url: https://www.reddit.com/r/ruby/comments/iopopr/need_study_buddy/
---
Hi there! I am currently learning Ruby, solving challenges on codewars kyu 7-5. Looking for a study buddy to do live sessions to practice coding in front of someone else. Language doesn’t really matter, it’s more about practicing coding in an ‘interview-like’ manner.

DM me to arrange a session.
