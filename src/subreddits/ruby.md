# ruby
## [1][AwesomePrint forked to AmazingPrint after it became stale.](https://www.reddit.com/r/ruby/comments/fvq3dv/awesomeprint_forked_to_amazingprint_after_it/)
- url: https://github.com/amazing-print/amazing_print
---

## [2][Google Maps API with StimulusJS](https://www.reddit.com/r/ruby/comments/fvx7xg/google_maps_api_with_stimulusjs/)
- url: https://www.driftingruby.com/episodes/google-maps-api-with-stimulusjs?utm_medium=social&amp;utm_campaign=weekly_episode&amp;utm_source=reddit
---

## [3][I'm a Java/JS dev learning Ruby to transition to a fully remote job. What should I focus on after the basics?](https://www.reddit.com/r/ruby/comments/fvm54k/im_a_javajs_dev_learning_ruby_to_transition_to_a/)
- url: https://www.reddit.com/r/ruby/comments/fvm54k/im_a_javajs_dev_learning_ruby_to_transition_to_a/
---
I'm most way through the Ruby On Rails Tutorial (Michael Hartl) book and it gives a good foundation on how to make a MVC site with Rails. Is this representative of most professional rails development or are there other skills I need to learn? I might spend some time learning to make JSON APIs with Rails aftewards.
## [4][Optimised Docker builds for Rails apps](https://www.reddit.com/r/ruby/comments/fvxzaa/optimised_docker_builds_for_rails_apps/)
- url: https://vitobotta.com/2020/04/06/optimised-docker-builds-rails-apps/
---

## [5][How to restart Sidekiq automatically for each deployment](https://www.reddit.com/r/ruby/comments/fvxnzk/how_to_restart_sidekiq_automatically_for_each/)
- url: https://www.codewithjason.com/restart-sidekiq-automatically-deployment/
---

## [6][How do I clean up objects created in a shared context](https://www.reddit.com/r/ruby/comments/fvuloq/how_do_i_clean_up_objects_created_in_a_shared/)
- url: https://www.reddit.com/r/ruby/comments/fvuloq/how_do_i_clean_up_objects_created_in_a_shared/
---
I am using \`RSpec.shared\_context\` which and I am setting up a few variables using let 

    RSpec.shared_context "common" do 
    let(:name) { #creates a database object }
    end

Now after I invoke it from describe block 

    describe "common test" do 
    include_context "common"
    #run few tests
    end

Now after running the describe block I want to clean it up. How do I rollback all the objects created in the shared context?
## [7][Understanding arel](https://www.reddit.com/r/ruby/comments/fvug4d/understanding_arel/)
- url: https://www.reddit.com/r/ruby/comments/fvug4d/understanding_arel/
---
Anyone knows what's a good way to understand how arel works? good blog post or MR maybe?

There's no special reason for me to do this other than it's how activerecord works and I like understanding the underlying system. Usually I can figure stuff out from the source code or the test suite but it's hard with arel.

Stuff I'm trying to figure out - how is the arel tree actually being transformed to actual sql? Couldn't find the class/classes where that happens.

Is there an equivalent library in other frameworks?
## [8][Stimulus Reflex](https://www.reddit.com/r/ruby/comments/fvd1pb/stimulus_reflex/)
- url: /r/rails/comments/fvd1cs/stimulus_reflex/
---

## [9][I've written a tiny gem which allows you to rerun your last IRB command using the "did you mean" suggestion if you make a typo. Hopefully it's a handy time-saver!](https://www.reddit.com/r/ruby/comments/fuyo4u/ive_written_a_tiny_gem_which_allows_you_to_rerun/)
- url: https://github.com/AaronC81/yes_i_did
---

## [10][Small web crawler ruby-based for practice purpose. How should I approach my project ?](https://www.reddit.com/r/ruby/comments/fv4dcw/small_web_crawler_rubybased_for_practice_purpose/)
- url: https://www.reddit.com/r/ruby/comments/fv4dcw/small_web_crawler_rubybased_for_practice_purpose/
---
So, paralleling with my rails course, I do like to code as practice few ruby scripts to get used to more and more ruby gems.  

I did a few projects using Nokogiri, rest client, open-uri, watir, mechanize and few others to scrape things I like. But crawling intere website it’s far more difficult. 

About the project...

I do like buy mangas. I have a list of 20 online stores to buy mangas. Most of it is small websites, so a few people know them. Every time I’m looking for a rare manga I do a search on every store from this list. 

I would like to do a web crawler to automate this for me.  But I would reduce to 3 or 4 stores which the program will crawl to find if the manga  I’m looking for is available.  

I need some guidance on how should I do that...


The scrape programa I made was based on inputs from the user and built the url using these values.  
I don’t think this is a good approach to use in this case. I would have to deal with a lot problems messing with the URL of every website. 

So, I thought using the home page of every store and use the Mechanize gem to find the search field from each store.  (Its that a good way of thinking ?)

So I though two program structure for this project and I don’t know if those are the best pra ctice for doing that.

Every website has its own html structure. Acessing the page results through css classes will be different for each one. 

So how should I approach ? Make exclusively functions to each store? Like each store will be a class and it has its own functions to return what I looking for.

Or the best way is to “global” functions which will have to cover a lot of exceptions due the different structure each website has. 

You can check out my scrape codes [here](https://github.com/nicweeaboo) . It’s really nothing special.
