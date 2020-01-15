# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/em8qtp/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/em8qtp/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/
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
## [3][The best approach for deleting existing association while updating](https://www.reddit.com/r/rails/comments/ep12dg/the_best_approach_for_deleting_existing/)
- url: https://www.reddit.com/r/rails/comments/ep12dg/the_best_approach_for_deleting_existing/
---
I'm writing an update API where when the user sends the updated parameter the already existing records for that association must be deleted and the ones in the API must be updated.   
If I have a relationship like Book has many Authors and if a particular book has a few authors and I'm trying to update the book by setting the new authors. The already existing authors for this book must be deleted and the new one updated. 

The association is defined as below

    has_many :authors, class_name: 'Authors', dependent: :destroy, autosave: true, after_add: :compute_something,
after_remove: :compute_some_other_thing

In my update method, I have something like this 

    book.authors.delete_all
    book.update_attributes(params[:book])

I'm creating new authors even if they exist after removing them by using the accepts\_nested\_attributes\_for property in rails.   


My question is do you see any pitfalls or bottlenecks with this approach in terms of the after\_add and after\_remove method execution or something else too.   
The authors aren't huge always it's mostly around 5 on average and around 25 at max.
## [4][Building a Hangman Game in Ruby in 25 minutes](https://www.reddit.com/r/rails/comments/eopxfx/building_a_hangman_game_in_ruby_in_25_minutes/)
- url: https://www.reddit.com/r/rails/comments/eopxfx/building_a_hangman_game_in_ruby_in_25_minutes/
---
Hey guys, I've launched a new Ruby build on YouTube. I know it's not Rails, but more Ruby on Rails content will be coming soon. I decided to do some Ruby language videos as a refresher and also to help promote the Ruby community on YouTube.

This build is a pretty interesting one as it involves a lot of interaction with the user and handling the user's input from the Terminal. Guess the wrong letter too many times and it's game over. Was a lot of fun writing this code for the video, hopefully it's useful to anyone getting into Ruby.

[https://www.youtube.com/watch?v=uBwGfswwRL4](https://www.youtube.com/watch?v=uBwGfswwRL4)

As always, I'm planning to post some more Ruby / Rails content to my channel to promote the Ruby language, so would be open to suggestions about content ideas. Thanks :)
## [5][Best way to implement demo functionality that resets per login?](https://www.reddit.com/r/rails/comments/eoseel/best_way_to_implement_demo_functionality_that/)
- url: https://www.reddit.com/r/rails/comments/eoseel/best_way_to_implement_demo_functionality_that/
---
The company I work for will be demoing one of our products (written in Rails) to potential clients. The sales team that will be performing the demo will need the data to be reset on a per login basis. I've been trying to come up with a way to store the data in the user's session, and tell the application to persist the data on session (instead of in the production db) when specific users sign into the system.

The approach outlined in this article is interesting, but wanted to get feedback from any/all of you first! [https://www.smashingmagazine.com/2017/12/using-sqlite-demo-web-apps/](https://www.smashingmagazine.com/2017/12/using-sqlite-demo-web-apps/)

The system uses Devise for authentication, and PostgreSQL for the db 

Do any of you have recommendations?
## [6][Python Tutorial Similar to Hartl's Rails Tutorial?](https://www.reddit.com/r/rails/comments/eovdxy/python_tutorial_similar_to_hartls_rails_tutorial/)
- url: https://www.reddit.com/r/rails/comments/eovdxy/python_tutorial_similar_to_hartls_rails_tutorial/
---
I first learned Rails by going through Michael Hartl's Ruby on Rails tutorial. Previously I had taught myself new languages by reading books or online docs, but his tutorial was so frequently recommended that I gave it a shot and loved it. 

Anyhow, question is: have any Rails folks here that took a similar path found a tutorial employing a similar approach but for Python? There are some machine learning things I'd like to experiment with, but would like to get some foundational Python first. I know there are tons of resources out there, and I've found plenty already, but there seems to be less of a consensus of a good place to start in the Python world (unless I missed it). 

If it helps, the things I found most useful about the Rails tutorial was: 

\- Building an actual functioning app as a "spine" to the course  
\- Teaching how things are generally done in a professional production environment, actually deploying to Heroku, etc.  
\- Learning things the "hard way" (not really using rails g scaffold, etc.)

Posting this here because I'm hoping some other Rails devs have taken a similar path!
## [7][What to keep in mind if I change an existing deployed Rails app to start using Docker?](https://www.reddit.com/r/rails/comments/eojyk1/what_to_keep_in_mind_if_i_change_an_existing/)
- url: https://www.reddit.com/r/rails/comments/eojyk1/what_to_keep_in_mind_if_i_change_an_existing/
---
I've got an existing Rails app that's been running for a year or so. It's running Rails 5.2 and I deploy it to AWS using Elastic Beanstalk, with the database living in RDS. I was learning Rails (and a lot of other web dev stuff) as I made it, so now that I look back, I probably could have gotten away with a cheaper DigitalOcean instance or something similar. As is the way of things, there's enough of the app hooked into the AWS ecosystem that pulling it out would be pretty daunting.

I was attempting to upgrade the app to Rails 6, but my existing instance was running Ruby 2.3, so I tried to make a new instance and do a blue-green deployment, but after quite a lot of trial-and-error, I was never able to get Rails 6 to work on Elastic Beanstalk. I tried searching around and trying arcane EB configuration tweaks, but I couldn't ever deploy it successfully. I'm guessing that webpacker and/or yarn weren't installing correctly, and those are two things I have only the most basic knowledge of.

It occurred to me that having to tweak Elastic Beanstalk this much was sort of defeating the point of using a deployment system that's supposed to be simple, so I thought it was time I bit the bullet and start experimenting with Docker.

I've been able to get development and test environments running on my local systems with docker-compose using a very basic [Dockerfile](https://github.com/HeadBeeGuy/alt-activities/blob/rails-6-docker/Dockerfile) and [docker-compose](https://github.com/HeadBeeGuy/alt-activities/blob/rails-6-docker/docker-compose.yml). I'm realizing that I can't just push these up to a webserver and expect it to take the place of the old instance, though. At the very least, I'd have to tell it that it's running in production, tell it how to talk to the production database, and make sure it's using a web server appropriate for production instead of development. I've got a couple newbie questions that I wanted to ask anyone who uses Docker in production:

* Do I specify separate Dockerfiles or docker-compose.yml files for production? Is there a way to test these before deploying them?
* Do I need the "db" image for production? Since the database lives in RDS, I imagine Rails will just talk to it and it doesn't need to be handled by Docker?
* From my research so far, it looks like I can use Elastic Beanstalk, ECS, or Fargate to deploy Rails apps. Is there one that's generally preferable?
* Should I use Puma? Passenger? nginx? Several in combination?

The app has been able to handle its traffic thus far using only one instance, so for now I don't think I need to worry about auto-scaling. Raw performance hasn't been as important as being to iterate on features rapidly.

Sorry this is scattered! Any advice would be welcome!
## [8][How did you turn your Rails app into an Android and iOS app?](https://www.reddit.com/r/rails/comments/eoiq8q/how_did_you_turn_your_rails_app_into_an_android/)
- url: https://www.reddit.com/r/rails/comments/eoiq8q/how_did_you_turn_your_rails_app_into_an_android/
---
Hello!

How did you turn your Rails app into an Android and iOS app? What were the pros/cons? Was it easy to get accepted into App Store and Google Play?

Im very new to the game but from what I understand, for a Rails app to be turned into mobile apps it would have to use PhoneGap to act as a browser basically? And then set up an API to communicate with the user's mobile phone and the Rails server?

Thanks!
## [9][Uploading images from JS Front End to Rails API](https://www.reddit.com/r/rails/comments/eom1gy/uploading_images_from_js_front_end_to_rails_api/)
- url: https://www.reddit.com/r/rails/comments/eom1gy/uploading_images_from_js_front_end_to_rails_api/
---
Currently working on a Instagram type clone to learn how to use rails as an API. I have a post model that accepts a caption and has attached image for active storage. I'm having trouble figuring out how to upload an image using a form on the vanilla js front end and sending it via json to the back. What is the best way to do this, judging by the lack of resources I'm guessing there is a better way to implement this. Thanks.
## [10][Anyone willing to help an old man troubleshoot my rails 6 app?](https://www.reddit.com/r/rails/comments/eobzel/anyone_willing_to_help_an_old_man_troubleshoot_my/)
- url: https://www.reddit.com/r/rails/comments/eobzel/anyone_willing_to_help_an_old_man_troubleshoot_my/
---
I'm 99 per cent sure the issue is with how I set up webpacker, despite my best efforts. User can't sign in or log out. Dropdown menus don't function and modal won't fire.

Not sure what I did! I'm really an old man (77), volunteer webmaster for my church website. Is someone willing to clone it and give it a spin? [github.com/ThomasConnolly/paul](https://github.com/ThomasConnolly/paul)

The error message I get when trying to log out is "ActiveRecord::RecordNotFound - Couldn't find User with 'id'=sign\_out:"

The dropdowns &amp; modal just don't filre; no error message.
## [11][¿ How to shared current_user ?](https://www.reddit.com/r/rails/comments/eoll07/how_to_shared_current_user/)
- url: https://www.reddit.com/r/rails/comments/eoll07/how_to_shared_current_user/
---
Hi, i'm trying to shared my `current_user context` in grapqhl :

I have same method in mutation and query base...

Base Query

    # frozen_string_literal: true
    
    module Queries
      class BaseQuery &lt; GraphQL::Schema::Resolver
        private
    
        def current_user
          context[:current_user] || (raise GraphQL::ExecutionError, 'You need to authenticate to perform this action')
        end
      end
    end

Base Mutation

    # frozen_string_literal: true
    
    module Mutations
      class BaseMutation &lt; GraphQL::Schema::RelayClassicMutation
        argument_class Types::BaseArgument
        field_class Types::BaseField
        input_object_class Types::BaseInputObject
        object_class Types::BaseObject
    
        def current_user
          context[:current_user] || (raise GraphQL::ExecutionError, 'You need to authenticate to perform this action')
        end
      end
    end

How to refactor this and use in mutations and queries? Regards
## [12][Cache Crispies - Fast, Flexible Rails Serializer](https://www.reddit.com/r/rails/comments/eoes6h/cache_crispies_fast_flexible_rails_serializer/)
- url: https://www.reddit.com/r/rails/comments/eoes6h/cache_crispies_fast_flexible_rails_serializer/
---
Picking a method of doing JSON serialization in Rails has not been an easy decision as of late. Especially if you're not able to break your APIs by moving to a JSON API structure. And trying to mix in a caching strategy, certainly doesn't help. That's the problem the new Cache Crispies gem was written to fix.  


[https://codenoble.com/blog/introducing-cache-crispies/](https://codenoble.com/blog/introducing-cache-crispies/)  
[https://github.com/codenoble/cache-crispies](https://github.com/codenoble/cache-crispies)
