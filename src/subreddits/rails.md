# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/
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
## [3][Rails 6 multiple databases support in Rails Event Store](https://www.reddit.com/r/rails/comments/hbaxq8/rails_6_multiple_databases_support_in_rails_event/)
- url: https://www.reddit.com/r/rails/comments/hbaxq8/rails_6_multiple_databases_support_in_rails_event/
---
Rails 6 released in August 2019 has brought us several new features. One of the notable changes is support for multiple databases. All details have been described in Rails guides and I’ve read already several blog posts describing how to do it. But how to use this feature to allow Rails Event Store data to be stored in a separate database? Check my [blog post where I document my experiments with it](https://blog.arkency.com/rails-multiple-databases-support-in-rails-event-store/).
## [4][Would You Build a Near Real-Time Feed/Channel Using Action Cable?](https://www.reddit.com/r/rails/comments/hay1n8/would_you_build_a_near_realtime_feedchannel_using/)
- url: https://www.reddit.com/r/rails/comments/hay1n8/would_you_build_a_near_realtime_feedchannel_using/
---
I have a newbie question.  I was watching a GoRails episode that shows how Action Cable and Stimulus could be used to [create real-time group chats](https://gorails.com/episodes/rails-group-chat-revised-part-1).

I'm thinking about creating something with feeds on topics - like a Reddit, Twitter, or Slack.

However, the channels/feeds would be around events occurring now. So I'd like to have something that is showing new messages/posts automatically near real-time (the channel could update every 30 seconds to one minute instead of real-time) at the bottom of the screen.

I'd like to build something that has the ability to handle thousands, if not tens of thousands of people viewing and posting on the channels concurrently.

I'm wondering where in the stack I might right into bottlenecks from volume (if any place), if I built it with Action Cable and Stimulus. I'm also wondering how much more expensive it would be than a different approach that may be near real-time (some sort of polling), if it at all.

I'm curious how others would approach building something like this.  Thank you!
## [5][Any gems or tips for handling tag input fields?](https://www.reddit.com/r/rails/comments/haww7c/any_gems_or_tips_for_handling_tag_input_fields/)
- url: https://www.reddit.com/r/rails/comments/haww7c/any_gems_or_tips_for_handling_tag_input_fields/
---
Similar to these are there any gems or methods to set up a field that can hold a list of tags then add them to the db?

 [https://bootstrap-tagsinput.github.io/bootstrap-tagsinput/examples/](https://bootstrap-tagsinput.github.io/bootstrap-tagsinput/examples/)
## [6][Newbie question: data items not included in Model?](https://www.reddit.com/r/rails/comments/hasx9t/newbie_question_data_items_not_included_in_model/)
- url: https://www.reddit.com/r/rails/comments/hasx9t/newbie_question_data_items_not_included_in_model/
---
Hi all, non-practicing CS-grad from before iPods were invented. Working through the *Rails Novice to Ninja* book to learn Rails. Something I didn't understand was how the Model interacts with the data items in the database, and how to access those data items in the Model.

I understand that when I use a generator to create the Model (with data attributes specified during generation), Rails creates an empty Model stub. I understand it also create the migration file with the data attributes specified. Why aren't these attributes also defined in the Model? How do I interact with these items in the Model?

Hope that made sense. Flipping through the book quickly, and the guides on RubyOnRails.org, I can't seem to find any examples of a Model interacting with individual data items in the database.
## [7][Cheap VPS for Rails App](https://www.reddit.com/r/rails/comments/havjfz/cheap_vps_for_rails_app/)
- url: https://www.reddit.com/r/rails/comments/havjfz/cheap_vps_for_rails_app/
---
Hello guys, I’m new to Rails and ruby. I have almost finished my first app and it will take some memory so I’m planing to use DigitalOcean(4GM ram, 2vCPUs) for $20/month and dedicated managed database $15/month. But today I found some cheap VPS hosting like - VPSDime, OVS. 

VPSDime offering 24 GB ram 4vCPU for $28/month, its really very cheap compared to DigitalOcean. So I have searched for its review specially for rails app but couldn’t found any so I’m confused. Pricing is nearly same on OVS too. 

I am already using DO for my other apps(not ruby/rails) and I’m happy with them but compare for new pricing on other hosting I’m confused. Can anyone suggest me something. Thanks in advance.
## [8][Best PWA options for Rails?](https://www.reddit.com/r/rails/comments/hal2l7/best_pwa_options_for_rails/)
- url: https://www.reddit.com/r/rails/comments/hal2l7/best_pwa_options_for_rails/
---
Looking for the current and best methods for building Progressive Web Apps with Rails. Any suggestions?

Thinking about maybe dabbling in StimulusReflex, but not seeing any solid discussion on PWA implementation / planning.

Goal is to build an app that feels as native as possible with a cohesive UI that eliminates building and maintaining multiple user interface styles for various devices.
## [9][How to optimize "the query of the day"?](https://www.reddit.com/r/rails/comments/has21u/how_to_optimize_the_query_of_the_day/)
- url: https://www.reddit.com/r/rails/comments/has21u/how_to_optimize_the_query_of_the_day/
---
In my home page, I want to run every day a different query.

&amp;#x200B;

In view/home/landing

I have this

&amp;#x200B;

    &lt;% if t.monday? %&gt;
    	&lt;% if @over_1000_followers.any? %&gt;
    		&lt;div&gt;Monday&lt;/div&gt;
    		&lt;%= render 'movies/movies', movies: @over_1000_followers, skip_pages: true %&gt;
    	&lt;% end %&gt;
    &lt;% elsif t.tuesday? %&gt;
    	&lt;div&gt;Tuesday&lt;/div&gt;
    	&lt;%= render 'movies/movies', movies: @over_1000_votes, skip_pages: true %&gt;
    &lt;% elsif t.wednesday? %&gt;
    	&lt;div&gt;Wednesday&lt;/div&gt;
    	&lt;%= render 'movies/movies', movies: @online_movies, skip_pages: true %&gt;
    &lt;% end %&gt;
    

Where `@over_1000_followers`, `over_1000_votes` and `online_movies` are definited in `models/concerns/scopes/movie.rb` with a queries like those

    module Scopes
      module Movie
          included do
            scope :over_1000_followers, lambda {
            where('followers_count &gt; ?', 1000)
              .limit(3)
          }
    
            scope :over_1000_votes, lambda {
            where('rating_count &gt; ?', 1000)
              .limit(3)
          }
    
            scope :online_movies, lambda {
            where(online: true)
              .joins(:download_links)
              .where(download_links: { play_online: true })
              .order(downloads_count: :desc).limit(3)
              .distinct
          }
    
        end
      end
    end

and I use them because they are added in controllers/home\_controller.rb

class HomeController &lt; ApplicationController

      def landing
        @over_1000_followers = landing_movies.over_1000_followers.pluck(:id)
        @over_1000_votes = landing_movies.over_1000_votes.pluck(:id)
        @online_movies = landing_movies.online_movies.pluck(:id)
      end

&amp;#x200B;

**My question is: in this way ...is the website running only "the query of the day" or everyone?**

**How to optimize it?**
## [10][Integrating Keycloak](https://www.reddit.com/r/rails/comments/han42k/integrating_keycloak/)
- url: https://www.reddit.com/r/rails/comments/han42k/integrating_keycloak/
---
Has anyone integrated [Keycloak](https://keycloak.org) with a Rails app?   

I'm not very experienced with authentication in Rails, aside from the basics of setting up Devise, and it is unclear to me what I need to do to make the first connection.

There is a Rack middleware gem, [keycloak-api-rails](https://github.com/looorent/keycloak-api-rails), which handles JWT token management for Keycloak.  However, I must do something in my app that will start the authN process.

I believe it should go something like this:

1. User arrives at my Rails site, but user is not yet authenticated
2. User is automatically redirected to my Keycloak server, and Keycloak provides the login page
3. User logs in with a local Keycloak username/password, or they login via one of the configured social integrations (Github, for example)
4. Upon successful authN, Keycloak redirects user back to my Rails app
5. The keycloak gem looks at the header and sees the Bearer token, then does the appropriate inspection to determine that it is valid; if so, it allows Rails to process the request fully.

**I don't know what must be done to get from step 1 to step 2.**  There are many gems for doing many things regarding user authentication, but I don't know which is right for this use case.  I would like to be as precise as possible with my gem choices such that I do not include a lot of unnecessary dependencies or configurations.

Suggestions?  TIA.
## [11][Help with configuring PostgreSQL and locale in Rails app](https://www.reddit.com/r/rails/comments/hamond/help_with_configuring_postgresql_and_locale_in/)
- url: https://www.reddit.com/r/rails/comments/hamond/help_with_configuring_postgresql_and_locale_in/
---
Hi, I'm using a Mac running Catalina to develop Rails applications. Since I'm going to use PostgreSQL in production, I've switched to using it in development (from sqlite). However, despite the fact that my Rails app specifies 'unicode' in the database.yml file, I cannot use it with a PostgreSQL database with UTF8 encoding. It keeps giving me the following error message:

    ERROR:  encoding "UTF8" does not match locale "en_NZ.ISO8859-15" DETAIL:  The chosen LC_CTYPE setting requires encoding "LATIN9". 

What can I do to allow PostgreSQL to let me use UTF8 encoding, despite the fact that my Mac is set to NZ English?

This is definitely a problem for me as I've been tasked with creating a website that will have both English and Thai options.
## [12][Ruby on Rails cheatsheet from Michael Hartl tutorials](https://www.reddit.com/r/rails/comments/ham7y6/ruby_on_rails_cheatsheet_from_michael_hartl/)
- url: https://www.reddit.com/r/rails/comments/ham7y6/ruby_on_rails_cheatsheet_from_michael_hartl/
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
