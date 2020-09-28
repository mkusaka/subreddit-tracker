# ruby
## [1][The Complexity of Active Record Transactions](https://www.reddit.com/r/ruby/comments/j1a97j/the_complexity_of_active_record_transactions/)
- url: https://janko.io/the-complexity-of-activerecord-transactions/
---

## [2][Hash to Struct in 💎](https://www.reddit.com/r/ruby/comments/j0vpyf/hash_to_struct_in/)
- url: https://i.redd.it/6bxehl5idqp51.png
---

## [3][Sandwich Scheduling: a look at Ruby’s Array cycle method](https://www.reddit.com/r/ruby/comments/j195jp/sandwich_scheduling_a_look_at_rubys_array_cycle/)
- url: https://www.mesomorphic.co.uk/blog/2020/09/28/sandwich-scheduling/
---

## [4][Rails Sidekiq configuration for micro services on reverse proxy.](https://www.reddit.com/r/ruby/comments/j1bsas/rails_sidekiq_configuration_for_micro_services_on/)
- url: https://blog.joshsoftware.com/2019/07/31/rails-sidekiq-configuration-for-micro-services-on-reverse-proxy/
---

## [5][To string punctuation issue](https://www.reddit.com/r/ruby/comments/j1azlf/to_string_punctuation_issue/)
- url: https://www.reddit.com/r/ruby/comments/j1azlf/to_string_punctuation_issue/
---
I have a Person object, with a to string method that returns output like `Sean is 20 years old`, it was returning that sentence with a full stop at the end, but I removed it because I also have a Couple object that stores two person objects and was returning text like `Sean is 20 years old., Sorcha is 21 years old.`

With the full stop now removed from the return string from the person object, in the main program I add the fullstop when calling to string on just a single person object with `puts "#{my_person1.to_s}."` which works fine, but when I try something similar like `puts "#{my_couple.to_s}."` the output has a full stop on a separate line:

    Sean is 20 years old.
    Sean is 20 years old, Sorcha is 21 years old
    .
    

Can I change this some how to put the full stop in the correct place?

The couple to string is: `puts "#{@person1.to_s}, #{@person2.to_s}"`
## [6][REPL playground for various Ruby Engines](https://www.reddit.com/r/ruby/comments/j0nsey/repl_playground_for_various_ruby_engines/)
- url: https://rubyapi.org/repl
---

## [7][Production ready project with Sinatra, Puma and Docker](https://www.reddit.com/r/ruby/comments/j0rw1a/production_ready_project_with_sinatra_puma_and/)
- url: https://www.reddit.com/r/ruby/comments/j0rw1a/production_ready_project_with_sinatra_puma_and/
---
Hello everyone! 😀

I developed a production ready (is it?!) project using **Sinatra** 🚀! It's open-source and the it has a specific purpose (see the README). Some features:

* Modular approach to setup everything. You can check [the main class](https://github.com/willianantunes/runner-said-no-one-ever/blob/d00ad82bc7e982f144fd0087567944cfcdca90c2/lib/runner_said_no_one_ever.rb#L16-L28) here where everything is setup;
* [Dedicated class](https://github.com/willianantunes/runner-said-no-one-ever/blob/d00ad82bc7e982f144fd0087567944cfcdca90c2/lib/runner_said_no_one_ever/config/settings.rb#L2) to store environment variables, so if someone is missing, it throws an error to you and other things;
* Basic [health check endpoint](https://github.com/willianantunes/runner-said-no-one-ever/blob/ba4046ef71d47ea05d8cdf517d3470979244cb38/lib/runner_said_no_one_ever/controllers/healthcheck_controller.rb#L2). I'm using [it on the Dockerfile](https://github.com/willianantunes/runner-said-no-one-ever/blob/6037b862a5a0f092e2e5597cce1d59c8c142deaf/Dockerfile#L21) as well!
* Talking about Dockerfile, it's using a dedicated user to run the web server installing only production dependencies (see more [here](https://github.com/willianantunes/runner-said-no-one-ever/blob/6037b862a5a0f092e2e5597cce1d59c8c142deaf/Dockerfile#L14-L15));
* As the project has [Docker Compose](https://github.com/willianantunes/runner-said-no-one-ever/blob/9c3a003c83f66bfb9a3c2bb02ce39b13b5258791/docker-compose.yaml#L3-L18), you can easily run the app and even run the tests;
* I implemented a strategy to log the requests and the messages from the business implementation;
* It has tests for all the endpoints, including the business logic and some infrastructure code;
* I created a dedicated configuration file for Puma 5.0.0 (latest version so far).

[Sample usage. See the project to know more! \(drawn on Excalidraw\)](https://preview.redd.it/asfdk2fz7pp51.png?width=1450&amp;format=png&amp;auto=webp&amp;s=a123dfad2e3e69fcc4d93261886f221e45d27300)

I would love to be critized by someone which has knowledge on Ruby and Sinatra, because I started coding on Ruby 2 months ago! Lessons learned are welcome 👍

  
[https://github.com/willianantunes/runner-said-no-one-ever/](https://github.com/willianantunes/runner-said-no-one-ever/)
## [8][Livestream: Rails CI with Raspberry Pi + GitHub Actions](https://www.reddit.com/r/ruby/comments/j0v91t/livestream_rails_ci_with_raspberry_pi_github/)
- url: https://www.twitch.tv/dogweather
---

## [9][Working on a mixer for DragonRuby Game Toolkit](https://www.reddit.com/r/ruby/comments/j0w0vc/working_on_a_mixer_for_dragonruby_game_toolkit/)
- url: https://youtu.be/b1pL4pEWymI
---

## [10][Best way to handle money?](https://www.reddit.com/r/ruby/comments/j0gsrt/best_way_to_handle_money/)
- url: https://www.reddit.com/r/ruby/comments/j0gsrt/best_way_to_handle_money/
---
I have a simple ruby script to generate payroll data. I am adding and multiplying but at the end of the day, I only need to support USD (for now). 

My database type is decimal(8,4) because I was told that’s accounting standards, so future proofing. 

However, what is the best way to handle addition and multiplication within my script for money? Is it BigDecimal? I don’t expect the numbers to be big (max $10,000) but doing it the “right way” is probably best. 

Thank you!
