# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ipfxw8/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ipfxw8/personal_projects_show_off_your_own_project_andor/
---
In this thread you can showcase your personal pet project to other redditors.

Need help with a specific problem or just wanna have some extra eyeballs on your code? Ask away!

A suggested format to get you started:

1. **Name of your project**
2. **A short description**
3. **Application stack**
4. **Link to Live app**
5. **Link to GitHub**
6. **You experience level**
7. **Other information or areas that you would like advice on**

 

^(Many thanks to Kritnc for getting the ball rolling.)
## [2][Whats the best thing to do? Write test after or before a feature?](https://www.reddit.com/r/rails/comments/irt5u9/whats_the_best_thing_to_do_write_test_after_or/)
- url: https://www.reddit.com/r/rails/comments/irt5u9/whats_the_best_thing_to_do_write_test_after_or/
---

## [3][I'm listing unmaintained Ruby projects](https://www.reddit.com/r/rails/comments/iredl3/im_listing_unmaintained_ruby_projects/)
- url: https://www.reddit.com/r/rails/comments/iredl3/im_listing_unmaintained_ruby_projects/
---
Hello everyone!

Yesterday  I was looking for some unmaintained open source Ruby project to put  some hours in and turns out they are quite hard to find.  
So I started listing and came up with a small collection.

I'm not thinking about building a huge list of projects, but to help those projects to meet their new maintainers.

Any help will be awesome!

The github repo is here: [https://github.com/attics/ruby\_attic](https://github.com/attics/ruby_attic)
## [4][Different Concerns with the same name for the same action for the same model](https://www.reddit.com/r/rails/comments/irubps/different_concerns_with_the_same_name_for_the/)
- url: https://www.reddit.com/r/rails/comments/irubps/different_concerns_with_the_same_name_for_the/
---
Suppose I have two concerns like this

    module Concerns::FirstConcern
      extend ActiveSupport::Concern
    
      included do
        after_commit :on_create, on: :create
      end
    
      private
    
      def on_create
        puts 'First Concern Create'
      end  
    end

and the second one like this

    module Concerns::SecondConcern
      extend ActiveSupport::Concern
    
      included do
        after_commit :on_create, on: :create
      end
    
      private
    
      def on_create
        puts 'Second Concern Create'
      end  
    end

And then in the user model have these both included

    class User &lt; ActiveRecord::Base

include Concerns::FirstConcern include Concerns::SecondConcern end

And then I see all the registered callback it gives me this.`User._commit_callbacks.select { |item| item.kind == :after }.map(&amp;:filter)`

I only get `[:on_create]`

I have two questions

1. Where can I find documentation of how these callbacks will be combined and which ones will be overridden? I'm assuming if it has the same name only the last included module ones will be included.
2. From the `_commit_callbacks` I can find out all the names of the registered callbacks. How do I find the location where they are defined.
## [5][What is the most efficient way to develop Achievements](https://www.reddit.com/r/rails/comments/irhnrk/what_is_the_most_efficient_way_to_develop/)
- url: https://www.reddit.com/r/rails/comments/irhnrk/what_is_the_most_efficient_way_to_develop/
---
Hey! So I'm developing an achievement system for my app. There are going to be lots of achievements, maybe 50.

What is the most efficient way to achieve this? For now, I created an `Achievement` model which logs achievement's: name, description, medal(gold, silver, bronze). And I have the `UserAchievement` model which tracks which achievement has user completed, and it contains only user_id and achievement_id.

Also, for now, I did something like this. Let's say there is an achievement that is completed after you complete 5 "jobs"(or anything). I added a function to the "Job" model called `check_achievement_completion`. And it basically checks whether the user who completed the Job has completed 5 of them, if so, I create `UserAchievement`.

And I soon realized that that function would get massive if there were 40 achievements only for Jobs.

What is the go-to strategy to develop the Achievement system and make it efficient?

Please let me know if you need any code snippets that could help you in understanding this issue.
## [6][Saving last 10 search results in cookies or local storage for ransack search/user](https://www.reddit.com/r/rails/comments/irgrtv/saving_last_10_search_results_in_cookies_or_local/)
- url: https://www.reddit.com/r/rails/comments/irgrtv/saving_last_10_search_results_in_cookies_or_local/
---
I'm simply trying to keep track of what the user searched for with ransack. Has anyone managed to pull this off? Even if the user is not logged in, their search results should show on the landing page.
## [7][cache warming strategy recs](https://www.reddit.com/r/rails/comments/ira56v/cache_warming_strategy_recs/)
- url: https://www.reddit.com/r/rails/comments/ira56v/cache_warming_strategy_recs/
---
In my rails app, I generate a lot of dashboards which aggregates data on the order of GBs. using the pull based caching strategy won’t work for me because the cache miss would take about 10 seconds to repopulate the cache.

the data is also high dimensionality. the dashboards can group by multiple columns over user specified dates.

I’m doing all sorts of gymnastics in order to warm up the cache in a lot of different places in the app, but it is getting unruly to the point where i’m not even sure if a cache populating line of code is getting used in the app. 

are there any gems that help with this problem?
## [8][Has anyone considered building a SIEM with (or based on) rails?](https://www.reddit.com/r/rails/comments/ircyvq/has_anyone_considered_building_a_siem_with_or/)
- url: https://www.reddit.com/r/rails/comments/ircyvq/has_anyone_considered_building_a_siem_with_or/
---
So, as i am browsing the web on my endless quest for perfection, i started to wonder if someone ever considered building a SIEM (Security information and event management) with rails as a component.

The way i see it, SIEM is all about getting data in a big database, and producing human readable stuff from that data (simplyfied explaination, obviously) so, using Rails to build a SIEM isnt all that far fetched i think?

There are some pretty expensive SIEMs out there that have a very “easy on the eyes” good looking interface, and good alerting. And then there are some pretty cheap SIEMs dat are build of “mix’n’match” opensource components, which is, if you have to look st it all day long, not easy on the eyes.

Surely, with RAILS one should be able to build a SIEM that is both low on budget and still easy on the eyes?

Preferable also scalable, and easy to extend?

Obviously, i know about ELK, Apache Metron, SIEMonster, AlienVault etc. Etc. But they are not based on Ruby on Rails ;-)
## [9][How to get the most out of GoRails as a newbie?](https://www.reddit.com/r/rails/comments/iqv8si/how_to_get_the_most_out_of_gorails_as_a_newbie/)
- url: https://www.reddit.com/r/rails/comments/iqv8si/how_to_get_the_most_out_of_gorails_as_a_newbie/
---
I managed to get a free year subscription of GoRails thanks to the Github Student Development Pack. I'm still learning Rails. I managed to get the basics of Rails down thanks to Michael Hartl Rails book.


I'm looking to learn more about Rails and follow along step by step projects, but for the most part, GoRails seems to be just a collection of mini videos on how stuff works. I'm not really sure what to start watching or if the site has any "how to build X". I'm still exploring the site though. 

Anyways, what's the best way to use this site to get the most out of my free 1 year subscription and make myself a better Rails Dev?

Thanks.
## [10][Real life application with Rails and Stimulus Reflex compared to the SPA approach](https://www.reddit.com/r/rails/comments/iqqnvg/real_life_application_with_rails_and_stimulus/)
- url: https://www.reddit.com/r/rails/comments/iqqnvg/real_life_application_with_rails_and_stimulus/
---
Hi Rails and JS devs, 

&amp;#x200B;

I never really understood the hype around SPAs and "modern JavaScript" development workflow.

It always seemed so much complicated to me for basically no real advantages.

I will talk about it more precisely in a future blog post. 

So I created a repo with the two stacks to compare how much they are different in a real life scenario. 

&amp;#x200B;

Turbolinks was a game changer and Stimulus Reflex adds the last thing to create reactive apps.

&amp;#x200B;

Here's the repo: [https://github.com/guillaumebriday/modern-datatables](https://github.com/guillaumebriday/modern-datatables)

&amp;#x200B;

Let me know how do you build reactive apps these days!
## [11][Managing Rails cache low-level concurrency](https://www.reddit.com/r/rails/comments/iqy3i3/managing_rails_cache_lowlevel_concurrency/)
- url: https://www.reddit.com/r/rails/comments/iqy3i3/managing_rails_cache_lowlevel_concurrency/
---
Let's say I have the following Worker:

```
class ModelWorker
  include Sidekiq::Worker

  def perform(id)
    cached_request = do_request
    Model.find(id).process_using_request(cached_request)
  end

  def do_request
    Rails.cache.fetch('key') do
      Client.new.fetch_data
    end
  end
end
```

Ideally, what I would like to happen, is that the code wrapped by the `fetch` call is only called once. Is this possible if I fire, simultaneously, tons of `ModelWorker` at the same time? Or will all jobs call `Client.new.fetch_data`?
