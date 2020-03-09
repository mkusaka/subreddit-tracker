# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/f9t9kq/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/f9t9kq/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/fduax9/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/fduax9/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Monolith serving React through webpack or seperate services.](https://www.reddit.com/r/rails/comments/fftzma/monolith_serving_react_through_webpack_or/)
- url: https://www.reddit.com/r/rails/comments/fftzma/monolith_serving_react_through_webpack_or/
---
Hello there...

I need some advice if you don't mind.

I have a service I have created as a backend only in rails with actioncable and multischema (tenant) postgres db.

I'm hating having to deal with CORS and extra security and running two services.

Would it be bad/not scalable to have it as a monolith?

Pros? Cons? 

I will be deploying this to ECS in AWS if that helps...

Please and thank you all for your help.
## [4][Help with associations](https://www.reddit.com/r/rails/comments/ffev1f/help_with_associations/)
- url: https://www.reddit.com/r/rails/comments/ffev1f/help_with_associations/
---
I'm building an application which consists on two modesl: `Task` and `Project`. Projects may have several zero or more tasks (`has_many` association), but tasks don't necessarily belong to a project. Well, I don't know how to program that last part in terms of routing, because if I implement this:

    Rails.application.routes.draw do
        resources :tasks
        resources :projects do
            resources :tasks
        end
    end

Rails assigns the same `Controller#Action` to both tasks, so I don't know how to distinguish between them:

    Controller#Action    Prefix            Verb   URI Pattern
    
    tasks#index          tasks             GET    /tasks(.:format)
    tasks#create                           POST   /tasks(.:format)
    
    tasks#index          project_tasks     GET    /projects/:project_id/tasks(.:format)
    tasks#create                           POST   /projects/:project_id/tasks(.:format)

Is there any way to separate them? Maybe creating different actions?
## [5][Help splitting an array into equal subarrays](https://www.reddit.com/r/rails/comments/ffgyjk/help_splitting_an_array_into_equal_subarrays/)
- url: https://www.reddit.com/r/rails/comments/ffgyjk/help_splitting_an_array_into_equal_subarrays/
---
Hi guys,

I am working on my final for my Ruby class(Ruby II not rails). I have an array that I have created from reading in a file. I am creating a word search puzzle. I need to list the words under my word search grid. I have 45 words and want 3 rows of 15 words each. I am also using a Prawn to create a pdf file to output my word search grid, words, and a grid key. Any advice on the best way to split my array into 3 equal columns to print out my words?

I was thinking about creating some kind of method that would calculate the length, the take that length and divide it by 3. From there I would print a string with one column ljust, one column centered, and one column rjust.

The other option I am entertaining is to create a boundary with the prawn and somehow break the list into 3 parts and put each part in the created boundary.

Any ideas and direction here would be greatly appreciated.

Here is my method to create my array (I have to put a space after my instance variable so it would allow me to include it on this post)

@ words = \[\]  
*file* = File.open(*file*)  
until *file*.eof?  
*word* = *file*.gets.chomp  
@ words&lt;&lt; *word*.delete(" ")  
@ words = @ words.sort\_by{ |*x*|-*x*.length }

end  
*file*.close

My list of words I am including

Apple  
Apricot  
Avocado  
Breadfruit  
Banana  
Blackberry  
Blackcurrant  
Blueberry  
Cherimoya  
Cherry  
Clementine  
Coconut  
Cranberry  
Custard Apple  
Durian  
Fig  
Grapefruit  
Grape  
Guava  
Jackfruit  
Kiwi  
Lemon  
Lime  
Loganberry  
Mandarin  
Mango  
Mangosteen  
Melon  
Nectarine  
Orange  
Papaya  
Peach  
Pear  
Persimmon  
Pineapple  
Plum  
Pomegranate  
Quince  
Satsuma  
Sharon  
Strawberry  
Tamarillo  
Tangerine  
Ugli  
Watermelon

My prawn pdf creation

def create\_pdf  
Prawn::Document.generate('puzzle2.pdf') do |*pdf*|  
*pdf*.font 'Courier', size: 20  
*pdf*.text 'Word Search', align: :center  
*pdf*.font 'Courier', size: 8  
*pdf*.move\_down(80)

*pdf*.text(print\_puzzle, align: :center)  
*pdf*.start\_new\_page  
*pdf*.text(print\_puzzle, align: :center)  
end  
end  
end
## [6][Minitest decorate test to change setup method](https://www.reddit.com/r/rails/comments/fffozb/minitest_decorate_test_to_change_setup_method/)
- url: https://www.reddit.com/r/rails/comments/fffozb/minitest_decorate_test_to_change_setup_method/
---
In Minitest I can use the `def setup` method to initialize some values before each test is run. My setup looks like this. 

    def setup 
    find_agent
    create_session
    end

I would want to modify the setup method to take a different agent from the `find_agent` method based on which test is running. Is there a way I can decorate each test so that the common setup method picks the right type of agent based on the test and creates a session?
## [7][Handling transactions with update_attributes!](https://www.reddit.com/r/rails/comments/ffi1b0/handling_transactions_with_update_attributes/)
- url: https://www.reddit.com/r/rails/comments/ffi1b0/handling_transactions_with_update_attributes/
---
I have a two below method that I want to wrap around a transaction so if update\_model fails destroy\_roles will be rolled back

    def update_model 
      if model.update_attributes(params)
         #do something
      else
         render /errors
      end
    end
    
    def destroy_roles
        role.destroy
    end

If I wanted to achieve the below using a transaction. 

    object.transaction do 
        destroy_roles
        update_model
    end
    
    def update_model
      begin
        model.udpate_attributes!(params
      resuce =&gt; e
        generate_bad_request(model.errors)
    end

Here is what I'm assuming will happen, 

1)  `destroy_roles` will destroy the roles of that particular model, in `update_model` the `update_attributes` raises an exception and we'll generate a bad request back for the user and also since it is inside a transaction the destroyed users in destroy\_roles will be restored as well. 

Could you please help me if my understanding is correct, I'm confused since we have handled the exception whether the transaction will roll back the destroy\_roles or do I have to re-raise the exception and so on.   


Any links to tutorial/documentation for the same would be really helpful as well.
## [8][Pls help: Upgraded from Rails 3 to 5, getting a Javascript error: Error: Unknown provider: tProvider &lt;- t](https://www.reddit.com/r/rails/comments/ffmvc6/pls_help_upgraded_from_rails_3_to_5_getting_a/)
- url: https://www.reddit.com/r/rails/comments/ffmvc6/pls_help_upgraded_from_rails_3_to_5_getting_a/
---
New to Rails and I'm getting this error:  **Error: Unknown provider: tProvider &lt;- t** 

and it seems to be a minification error? This error is also only happening on production, development environment is working fine. We don't have a local production server set up, so I have no idea how to test this before I deploy? 

I can produce this error on dev by just commenting out the **config.assets.debug**  line... I've commented this out and have tried to remove the js uglifier ( config.assets.js\_compressor = :uglifier), and have tried to put  config.assets.js\_compressor = Uglifier.new(mangle: false), but this doesn't work either.

I've also tried to make my controllers minification safe, but that doesn't seem to be doing anything either.

Maybe this things will work in production (and just the config.assets.debug line is messing up the output), but I don't know how to test this without deploying if I don't have a local prod server up?

I'm very confused at this whole error. Prod is broken, dev is working. Please help
## [9]["VFS Connection does not exist"](https://www.reddit.com/r/rails/comments/ffd94j/vfs_connection_does_not_exist/)
- url: https://www.reddit.com/r/rails/comments/ffd94j/vfs_connection_does_not_exist/
---
As the title says, when i try to run the $rails server command andpreview running application, it tells me "Oops, VFS Connection does notexist". Any way to fix this? I'm on Firefox if it helps.
## [10][How to view all available methods in template or controller?](https://www.reddit.com/r/rails/comments/ff8spe/how_to_view_all_available_methods_in_template_or/)
- url: https://www.reddit.com/r/rails/comments/ff8spe/how_to_view_all_available_methods_in_template_or/
---
Hi, I have been learning Rails for about a week and I am having very hard time with understanding which methods can be used when because everything seems like magically appearing in templates and controllers.

Is there a way to dump all available methods in debug statement or something to see and better understand them?

Also, if you have other tips on understanding this topic, your comment would be really helpful!
## [11][How do I handle this join?](https://www.reddit.com/r/rails/comments/ff3raq/how_do_i_handle_this_join/)
- url: https://www.reddit.com/r/rails/comments/ff3raq/how_do_i_handle_this_join/
---
I have a side project with an interesting problem. I have two tables:

1. Books
2. Restrictions

I want to be able to restrict certain books with a scope. I'm pretty new to joins, and I feel like this would be easy, *except* I also have another unique constraint. I have set up restrictions to be polymorphic. So, there is a record\_id and record\_type column on restrictions. This is so, that in the future, I have the ability to restrict other things than just books.

My attempt is something like this:

`Book.all.joins("INNER JOIN restrictions ON restrictions.record_id =` `books.id` `AND restrictions.record_type = 'Book'")`

But this is just returning all books, and I am trying to get a list of books that isn't restricted. Any ideas?
## [12][How do I efficiently query telemetry data?](https://www.reddit.com/r/rails/comments/feswoh/how_do_i_efficiently_query_telemetry_data/)
- url: https://www.reddit.com/r/rails/comments/feswoh/how_do_i_efficiently_query_telemetry_data/
---
I have this project where there are N devices in the field that beam back telemetry data **once every minute** per device. The ingest API for that is fine, but I'm running into some concerns about _reading_ that data, specifically for building a graph.

Data is in PostgreSQL and roughly resembles this:

```
    create_table :telemetry do |t|
      # some unrelated stuff here
      t.float         :some_value, null: false, default: 0.0
      t.timestamps
    end
```

So there's one row per minute per device in this database. At 24 hours in one day times 60 rows per hour (one per minute), we have **1,440 rows of telemetry data!**

Now, querying this is easy enough if we don't care about how evenly distributed that data is:

```
@telemetry = Telemetry.where('created_at &gt; ?', 24.hours.ago).order(created_at: :asc)
```

I could of course throw a limit on that, but then we'd just get the first N minutes of the data. So for example if I were to limit that to 50 rows, I'd only get back less than the first hour of all 24 hours.

Is there a better way to do this? For example, let's say I want two data points per hour (so one every 30 minutes). I _could_ try creating a loop then using the loop to craft and execute a new query for every trip through said loop, but that also sounds pretty inefficient.

How do I _efficiently_ do something like this? What should I be googling for?
