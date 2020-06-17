# ruby
## [1][Interview with DHH: the Ruby on Rails creator, co-founder &amp; CTO of Basecamp and Le Mans class-winning race car driver](https://www.reddit.com/r/ruby/comments/hami35/interview_with_dhh_the_ruby_on_rails_creator/)
- url: https://evrone.com/dhh-interview
---

## [2][Rails 6.1's ActiveModel Errors Revamp](https://www.reddit.com/r/ruby/comments/hap4u9/rails_61s_activemodel_errors_revamp/)
- url: https://code.lulalala.com/2020/0531-1013.html
---

## [3][Ruby on rails cheatsheet from Michael Hartl tutorials](https://www.reddit.com/r/ruby/comments/ham5u8/ruby_on_rails_cheatsheet_from_michael_hartl/)
- url: https://www.reddit.com/r/ruby/comments/ham5u8/ruby_on_rails_cheatsheet_from_michael_hartl/
---
Hi folks,   

Hope everyone is staying safe!

I started learning Ruby-on-Rails framework two weeks before &amp; made a cheat sheet with minimal explanations of different concepts, referred from **Michael Hartl**’s "[Learn Web Development with Rails - Fourth Edition](https://www.pdfdrive.com/ruby-on-rails-tutorial-learn-web-development-with-rails-4th-edition-e184254589.html)" tutorials.  

The textbook is really awesome &amp; helped me to build a simple blogging website similar to the basic version of Twitter and the source code is available in github: [https://github.com/ddlogesh/rails-tutorial](https://github.com/ddlogesh/rails-tutorial) 

I would like to share these two cheat sheets, which may be a good start for preparing ruby-on-rails framework or recollect the concepts learned from the above-mentioned textbook.  

Rails: [https://www.notion.so/Ruby-on-Rails-Cheat-Sheet-61106a73031d46ec81c285daaebf1409](https://www.notion.so/Ruby-on-Rails-Cheat-Sheet-61106a73031d46ec81c285daaebf1409)

Ruby: [https://www.notion.so/Ruby-Cheat-Sheet-7c8aabc9268b4e65b5691245dd19068c](https://www.notion.so/Ruby-Cheat-Sheet-7c8aabc9268b4e65b5691245dd19068c)

If you are a beginner to Ruby, please watch this 4-hours tutorial from **freeCodeCamp**: [https://www.youtube.com/watch?v=t\_ispmWmdjY](https://www.youtube.com/watch?v=t_ispmWmdjY). Though it's quite a long, it's worth watching it!

Please do share these cheatsheets with your Ruby-on-Rails enthusiasts, who may be interested to learn Ruby-on-Rails.
## [4][Using Service Objects in Ruby on Rails](https://www.reddit.com/r/ruby/comments/haqv41/using_service_objects_in_ruby_on_rails/)
- url: https://blog.appsignal.com/2020/06/17/using-service-objects-in-ruby-on-rails.html
---

## [5][Are "nonclass singleton methods as rare as class methods are common"](https://www.reddit.com/r/ruby/comments/ham3k3/are_nonclass_singleton_methods_as_rare_as_class/)
- url: https://www.reddit.com/r/ruby/comments/ham3k3/are_nonclass_singleton_methods_as_rare_as_class/
---
In his 2011 book "Eloquent Ruby", Russ Olsen makes the states that 

&gt;"nonclass singleton methods \[he refers to instance methods\] are as rare as class methods are common."

He makes the point that the most common usage of "instance methods" is the definition of static methods (because this is just adding an instance method to a Class object instance). 

And then there is the usage of instance methods in RSpec and Mocha in stubs...

Is anyone using instance methods on a regular basis for anything else?
## [6][Introducing EventStoreClient - a Ruby implementation for Greg's EventStore](https://www.reddit.com/r/ruby/comments/haqcpj/introducing_eventstoreclient_a_ruby/)
- url: https://blog.arkency.com/introducing-eventstoreclient/
---

## [7][Heroku Add-ons I am working on](https://www.reddit.com/r/ruby/comments/haq6o2/heroku_addons_i_am_working_on/)
- url: https://www.reddit.com/r/ruby/comments/haq6o2/heroku_addons_i_am_working_on/
---
Hi Devs!

If you are on Heroku and use Heroku Scheduler, I am working on 2 projects that might interest you! Feel free to reach out!

[https://schedulerctl.com](https://schedulerctl.com/)

[https://schedulermon.com](https://schedulermon.com/)
## [8][How to add an index to a returned list after scraping?](https://www.reddit.com/r/ruby/comments/hakj39/how_to_add_an_index_to_a_returned_list_after/)
- url: https://www.reddit.com/r/ruby/comments/hakj39/how_to_add_an_index_to_a_returned_list_after/
---
Scraping a website to return a list of fighter jets. How can i iterate through the return and add an index to each element in the array ?

    class Flight::CLI       # :: Namespacing is a way of bundling logically related objects together.
      def call
        puts "\nTop US Attack Aircraft.\n"
        get_attack_aircraft
        # list_aircraft
        # get_more_info_for_aircraft
      end
    
      def get_attack_aircraft
        Flight::Scraper.scrape_aircraft
        @aircraft = Flight::Aircraft.all
        @aircraft.sort  
      end

returns

    Top US Attack Aircraft.
    "A-10 Thunderbolt II"
    "F-22 Raptor"
    "F-35A Lightning II"
    "F/A-18C/D Hornet"
    "F-16 Fighting Falcon"
    "EA-18G Growler | U.S. Navy Aircraft"
    "F/A-18E/F Super Hornet"
    "OH-58D Kiowa Warrior"
    "AC-130W Stinger II"
    "MQ-9 Reaper"
    "MQ-1B Predator"
    "F-35C Lightning II"
    "F-35B Lightning II"
    "F-15E Strike Eagle"
    "AH-64 Apache Longbow"
    "AH-1W Super Cobra"
    "AH-1Z Viper"
    "AV-8B Harrier II"

The disired return is 

    Top US Attack Aircraft.
    1. "A-10 Thunderbolt II"
    2. "F-22 Raptor"
    3. "F-35A Lightning II"
    4. "F/A-18C/D Hornet"
    5. "F-16 Fighting Falcon"
    6. "EA-18G Growler | U.S. Navy Aircraft"
    7. "F/A-18E/F Super Hornet"
    8. "OH-58D Kiowa Warrior"
    9. "AC-130W Stinger II"
    10. "MQ-9 Reaper"
    11. "MQ-1B Predator"
    12. "F-35C Lightning II"
    13. "F-35B Lightning II"
    14. "F-15E Strike Eagle"
    15. "AH-64 Apache Longbow"
    16. "AH-1W Super Cobra"
    17. "AH-1Z Viper"
    18. "AV-8B Harrier II"

an explination/link would be much appreciated rather then just the answer, thanks !
## [9][football.csv - sportdb-importers Gem Update - Read League &amp; Match Datasets in Comma-Separated Values (CSV) Format into Any SQL Database](https://www.reddit.com/r/ruby/comments/ha022y/footballcsv_sportdbimporters_gem_update_read/)
- url: https://github.com/sportdb/sport.db/tree/master/sportdb-importers
---

## [10][After 10 years of Ruby I just realized you can step directly into a method with pry-byebug from IRB...](https://www.reddit.com/r/ruby/comments/h9o0e9/after_10_years_of_ruby_i_just_realized_you_can/)
- url: https://www.reddit.com/r/ruby/comments/h9o0e9/after_10_years_of_ruby_i_just_realized_you_can/
---
Adding a `binding.pry` to your source and running your source is one thing I've always known you can do, but that has always been a chore for stepping into methods starting from IRB. Normally my hack around this was to redefine `my_method` with a `binding.pry` on the first line of that method and load that back into my repl

Turns out you can just do `binding.pry; my_method` and Pry will step right in; no editing or redefinition required

Counting the hours wasted...

EDIT: source where I stumbled across this https://stackoverflow.com/a/36762059
