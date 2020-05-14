# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/
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
## [2][Database-driven authorization in Rails using CanCanCan - Abilities in DB and MetaProgramming](https://www.reddit.com/r/rails/comments/gjkha1/databasedriven_authorization_in_rails_using/)
- url: https://www.reddit.com/r/rails/comments/gjkha1/databasedriven_authorization_in_rails_using/
---
Hi ruby family,

As an initiative to give back to the community, I have started writing a series of blogs on ruby and ruby on rails. A few days back, I published a post on **authorizing resources in rails using CanCanCan**. As a continuation of the previous post, I have recently published another post on how to **implement database-driven authorization using CanCanCan**.

Some of the key issues that I tried to solve was :

1. The Growing size of the ability file
2. Abilities being hard to maintain.
3. Redeployment of the application for every change in the ability file
4. Storing abilities in a database

If you think this can be extrapolated and be made into a gem, let me know, and let's work together to create an awesome library.

[https://addytalks.tech/2020/05/14/rails-cancancan-database-driven-authorization/](https://addytalks.tech/2020/05/14/rails-cancancan-database-driven-authorization/)

You check out my previous post here - 

[https://addytalks.tech/2020/05/03/ruby-on-rails-authorization-with-cancancan/](https://addytalks.tech/2020/05/03/ruby-on-rails-authorization-with-cancancan/)
## [3][using database cleaner gem to clean doesn't save the objects to database](https://www.reddit.com/r/rails/comments/gjfsrj/using_database_cleaner_gem_to_clean_doesnt_save/)
- url: https://www.reddit.com/r/rails/comments/gjfsrj/using_database_cleaner_gem_to_clean_doesnt_save/
---
I am running minitests and I want to use database cleaner gem with transaction strategy to clean up items I created during the tests. I'm using MySQL btw. 

I followed this snippet [https://stackoverflow.com/questions/15675125/database-cleaner-not-working-in-minitest-rails](https://stackoverflow.com/questions/15675125/database-cleaner-not-working-in-minitest-rails) and set my code something like this. 

    module AroundEachTest
      def before_setup
        super
        DatabaseCleaner.start
      end
    
      def after_teardown
        super
        DatabaseCleaner.clean
      end
    end
    
    class Minitest::Test
      include AroundEachTest
      def test_index
        # create 5 articles in the database
        # call index action
        # assert the response has five objects
      end
    end

The trouble I'm facing is in the test I'm able to get all 5 objects. but when I hit the index action and try to query for all objects I get 0 records. And none of the records is actually being saved in the database. 

And hence the tests fails. 

My question is

1. why isn't the record being saved in the database in the first place?
2. And why isn't it available in the index action 

My guesses are 

Since we started a transaction in the before\_setup method it doesn't get saved? I'm not sure though.
## [4][I have a CLI-tool which downloads content from a website...How could I run (localhost) this in an app?](https://www.reddit.com/r/rails/comments/gjgsuq/i_have_a_clitool_which_downloads_content_from_a/)
- url: https://www.reddit.com/r/rails/comments/gjgsuq/i_have_a_clitool_which_downloads_content_from_a/
---
I’m messing around with a CLI-tool crunchyroll download from this [repo](https://github.com/DasKraken/CR-dl/blob/master/README.md) which allows me to login in my crunchyroll and download the video from given URL.

I’m having hard time to figure it out how things will work. I mean, it’s not like the blog app or finnacial tracker I have done so far. I don’t wanna make a crud or store anything.

I actually don’t know if this is possible. Just because I can download from my shell doesn’t mean I could in an app, right ?

I tried making a Download model with a “url” column, a form to submit the URL and in my create action calling the “system” ruby command to execute my CLI-tool.
The command to login into my account works fine but the command to download does not. 

I would like to understand what is missing. If the problem is the CLI-tool that actually does not work for this kind of finality or there is something that I should be doing in my app.
## [5][Standalone ActiveAdmin app](https://www.reddit.com/r/rails/comments/gj1vhk/standalone_activeadmin_app/)
- url: https://www.reddit.com/r/rails/comments/gj1vhk/standalone_activeadmin_app/
---
Hi, I'm working on rails API app as a backend for mobile app, I've deployed an app on AWS ECS, and it's work. Now I want to add an admin panel as a separate service, I decided to use ActiveAdmin because I'm know that framework, now I'm thinking how to set up active admin as a separate app, I have to share models between apps? If so which is the best way to do this, separate git repo for models? Or maybe there is some Smart way to do this?
## [6][How do you handle input data for service objects (or more specifically Interactors/Organizers)](https://www.reddit.com/r/rails/comments/gj10xm/how_do_you_handle_input_data_for_service_objects/)
- url: https://www.reddit.com/r/rails/comments/gj10xm/how_do_you_handle_input_data_for_service_objects/
---
Like the title suggests I'm currently using the [Interactor gem](https://github.com/collectiveidea/interactor) in my project. For those who are new to this gem, Interactors are simple, single-purpose objects (basically like service objects with some extra stuff, so I think this question applies to service objects as well). The interactor gem also has this concept of Organizers, objects which single purpose is to run other interactors. 

Now to my problem: I have an `Orders` controller with a corresponding model and an `#update` method. Each Order also has multiple `Shipments`. This #update method calls an Organizer which in turn calls multiple interactors to handle the business logic. However, before any of this happens I need to find the Order and it's Shipments via the provided `id` and need to check whether they can be updated (not every shipment and order can be). This is pretty easy as I just need to do two queries:

* For the Order: `order = current_user.orders.where(id: params[:id], status: :finalized)`
* For the shipment: `shipments = order.shipments.where(status: :pending)`

Because these queries essentially represent some "logic" I want to test them (for example, does it correctly fail when there are no pending shipments in an order). Now I'm wondering: where do I put these queries?

My first idea was to put them in the Organizer because this is the part that calls all other interactors but then again I'm not sure whether this is still within the purpose of the Organizer. And even if, then I would need to create organizer specs, something that isn't mentioned anywhere in the Interactor gem's docs and makes me wonder whether this is the correct place for these queries. 

My second idea was to create another Interactor which "sets up" all the input for the following interactors. However, since these queries are quite specialized and aren't going to be use elsewhere I'm wondering if this makes sense.

My last idea was to put them in the controller and let the controller pass this data into the Organizer. But this doesn't seem all to clean with the controller and would I then test this with controller tests or system tests?

What do you guys think? Should I test these queries with unit-tests or should they be part of my system tests and where should I put these queries in the first place?
## [7][How to handle social login from a react native app to rails api](https://www.reddit.com/r/rails/comments/givaz1/how_to_handle_social_login_from_a_react_native/)
- url: https://www.reddit.com/r/rails/comments/givaz1/how_to_handle_social_login_from_a_react_native/
---
Hi guys

I'm trying to set up social login for my react native app which has a rails backend.

I've read a lot of tutorials but most of them focus on rails as a web app solution, not so much as an api. For example, I think I might be needing the omniauth gem but I'm not sure if I also need devise when I'm not making use of most of its features (views, controllers, ...) and when I'm just issuing JWT's.  


So I'm not sure what else I need or how best to approach this.

I was also contemplating using something like auth0 if using something like omniauth is not secure/scalable enough.

&amp;#x200B;

Any thoughts or help would be greatly appreciated!  
Stay safe!
## [8][How do sites like Freecodecamp work?](https://www.reddit.com/r/rails/comments/gixx9u/how_do_sites_like_freecodecamp_work/)
- url: https://www.reddit.com/r/rails/comments/gixx9u/how_do_sites_like_freecodecamp_work/
---
How do I go about building a coding challenge site like freecodecamp. I know how to embed code editors using code mirror or ace.c9. But these editors don’t compile the code and I also don’t know how to add the “run the tests” button.

Where can I learn to implement a site like this (preferably using ruby on rails)? Is there anything I need to learn or is there any API or ruby gem I can use?
## [9][Transacted db:migrate:redo rake task](https://www.reddit.com/r/rails/comments/givoeb/transacted_dbmigrateredo_rake_task/)
- url: https://www.reddit.com/r/rails/comments/givoeb/transacted_dbmigrateredo_rake_task/
---
If you are familiar with the `db:migrate:redo` rake task you've probably noticed how it's convenient for fixing one migration, usually that's latest one.

If run without params, at first the task rolls back the latest applied migration, i.e. runs `down` , and then immediately migrates it, i.e. runs `up`. Now, guess what happens if `up` step fails? Right, the `down` step remains performed, what means the latest applied migration now is not the one was a moment ago but the previous one. Next `db:migrate:redo` run, usually after the latest migration fix, surprisingly tries to run the previous migration. That's something unexpected, drives made, and makes the task useless in the scenario when a migration is being fixed that usually requires numerous of changes and `redo`'ing it several times.

I promote a [fix](https://github.com/rails/rails/pull/39026) for that problem. That wraps entire `db:migrate:redo` task into a transaction. It won't change current migration to the previous one when `down` step fails as in the scenario above. That makes fixing a certain migration with more comfort since redo task is idempotent and can run as many times as needed.

This post aims to gather feedback on the specified problem and the [pull request](https://github.com/rails/rails/pull/39026).

All participants, thank you for your attention!
## [10][How to create file(output.pdf) to table database in rails console?](https://www.reddit.com/r/rails/comments/givqm5/how_to_create_fileoutputpdf_to_table_database_in/)
- url: https://www.reddit.com/r/rails/comments/givqm5/how_to_create_fileoutputpdf_to_table_database_in/
---
I have Document table having a column of file:binary status:integer and expiration\_date:datetime.   
Document.create(file: Rack::Test::UploadedFile.new("#{Rails.root}/spec/fixtures/files/testing\_file\_#{rand(100)}.csv", 'application/csv'), status: 0, expiration\_date= "2020-05-13 16:48:21")  
After that, when i want to see the record the file is nil and the rest column has a record.
## [11][Deploying Rails 6 w/ Webpacker to Nginx VPS](https://www.reddit.com/r/rails/comments/gitcbu/deploying_rails_6_w_webpacker_to_nginx_vps/)
- url: https://www.reddit.com/r/rails/comments/gitcbu/deploying_rails_6_w_webpacker_to_nginx_vps/
---
I have a Rails 6 project that uses webpacker to manage its assets which I'm trying to deploy in a VPS w/ nginx + passenger as my application server. However, the css assets won't seem to display properly. JS files seem to be working ok.

What I tried to do:

```

$ RAILS_ENV=production bundle exec rails assets:precompile

$ RAILS_ENV=production RACK_ENV=production NODE_ENV=production ./bin/webpack

```

Why won't my css load correctly?
