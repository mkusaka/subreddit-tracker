# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/fx6je4/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/fx6je4/personal_projects_show_off_your_own_project_andor/
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
## [3][Background job to constantly listen for MQTT requests?](https://www.reddit.com/r/rails/comments/fzp4cn/background_job_to_constantly_listen_for_mqtt/)
- url: https://www.reddit.com/r/rails/comments/fzp4cn/background_job_to_constantly_listen_for_mqtt/
---
I don't know if this is even possible. But given how MQTT works I want to have my backend rails server subscribe to some topics (Which is mainly sensor data) from an MQTT broker (Mosquitto in this case).

However, based on my limited knowledge of Rails and specifically background jobs....Im not sure if what I want is possible, since most background jobs are meant to run ever x amount of seconds/minutes/etc..

But the way MQTT works (from my understanding) is that is needs to "block" and listen for published messages to the broker (by being subscribed to the topic). However this would mean I need to have a job that constantly listens. Im not sure if this even do-able...or is there some way I am not thinking of? 

&amp;#x200B;

Thanks!
## [4][Help with major refactor and data structures](https://www.reddit.com/r/rails/comments/fzvlk1/help_with_major_refactor_and_data_structures/)
- url: https://www.reddit.com/r/rails/comments/fzvlk1/help_with_major_refactor_and_data_structures/
---
Hi guys,  


I have to do a major refactor in a simple app that I own, and don't know well how.

So far, I have been using user with roles (I have several: business owners, editors, admins..)  
Thing is, every role needs custom attributes  very different from each others. In the case of business owners it is so specific that it needs its own associations (that also have to change from has\_one to has\_many)  


So far I am in a mess. What is the best way to order this kind of structures? (considering that has to work with the previous data)
## [5][How to serve external images locally?](https://www.reddit.com/r/rails/comments/fzo2bn/how_to_serve_external_images_locally/)
- url: https://www.reddit.com/r/rails/comments/fzo2bn/how_to_serve_external_images_locally/
---
Hi,
I am trying to load images from our s3 bucket and ,make this happen:
Serve the URL locally: instead of aws.com/image/xxxx, how can I make it mywebsite.com/image/xxxx?

I thought of using the tmp/ directory but not sure this would work because there is no guarantee how long it stays there. Any ideas?
Sorry if this is kind of a novice question.
## [6][best VSCode extensions for Ruby/Rails](https://www.reddit.com/r/rails/comments/fz3muy/best_vscode_extensions_for_rubyrails/)
- url: https://www.reddit.com/r/rails/comments/fz3muy/best_vscode_extensions_for_rubyrails/
---
thought it would be a great place to ask here, does anyone have a list of some of the best VSCode extensions to have for Ruby and Rails development?  

I actually have too many that I got from a recommended source and now my format selection keyboard shortcut isn't even working on one laptop, but may just strip out all the extensions and start bare on that. Though, that system is an  Ubuntu system.

EDIT: FWIW, I did actually find this

https://medium.com/better-programming/code-like-a-pro-tooling-to-supercharge-vs-code-for-ruby-bf2ae61df5e3
## [7][How do I reverse the associations on instances of a has-and-belongs-to-many-association?](https://www.reddit.com/r/rails/comments/fz70m3/how_do_i_reverse_the_associations_on_instances_of/)
- url: https://www.reddit.com/r/rails/comments/fz70m3/how_do_i_reverse_the_associations_on_instances_of/
---
  
Hi all.  


I have a model Venues, which has\_many Categories. I will also have many Categories which belongs\_to many Venues.  


I want to be able to do `venue_instance.categories`, as well as `category_instance.venue`  


I've had a look at the [documentation](https://guides.rubyonrails.org/association_basics.html#the-has-and-belongs-to-many-association) on this, which says:

    "if you need to work with the relationship model as an independent entity. If you don't need to do anything with the relationship model, it may be simpler to set up a has_and_belongs_to_many relationship"

So since I don't need to work with the relationship model as an independent entity at all, I've tried to implement the has\_and\_belongs\_to\_many relationship. Here's what I've done, again, following the docs:

    class Venue &lt; ApplicationRecord
      has_and_belongs_to_many :categories
    end
    
    class Category &lt; ApplicationRecord
      has_and_belongs_to_many :venues
    end
    
    class CreateVenuesAndCategories &lt; ActiveRecord::Migration[6.0]
      def change
        create_table :venues_categories, id: false do |t|
          t.belongs_to :venue
          t.belongs_to :category
        end
      end
    end

  
Over to the terminal. Here's what I've tried to test instances of Category associated with many Venues. (I already have Venues in the DB, but no Categories as yet).

    a = Category.new(name: "something")
    =&gt; #&lt;Category:0x00007fe55287ad40 id: nil, name: "something", created_at: nil, updated_at: nil&gt;
    a.save!
    
    a.venues
    =&gt; []
    
    Venue.last.categories
    =&gt; []

So it returns the array of relations successfully (currently empty). Now I've followed [this S.O. post](https://stackoverflow.com/questions/33693556/creating-a-habtm-relationship) to learn how to create the associations between instances:

    Venue.last.categories &lt;&lt; a
    Venue.last.categories
    =&gt; [#&lt;Category:0x00007f913e83a2d0 id: 4, name: "something", created_at: Sat, 11 Apr 2020 14:51:18 UTC +00:00, updated_at: Sat, 11 Apr 2020 14:51:18 UTC +00:00&gt;]

Great. So my Venue instance will return my Category instance. But here's where it breaks down:

    a.venues
    =&gt; []

See, the relation is only working one way. Now, I could manually do the reverse, i.e :

    a.venues &lt;&lt; Venue.last

And this will work. But it seems very undesirable. I essentially want that reverse relationship to always be true. I shouldn't have to explicitly make the connection for each model. And it would be easy for the relations to drift out of synchronisation doing it this way.  


Now, I have also looked at doing a has\_many through relationship. But this would involve me creating an additional model which I don't intent to interact with, and which will make my ActiveRecord queries more complicated. So this doesn't seem like an optimum solution.  


I'm probably thinking about the problem wrong, and this seems like an extremely basic thing to achieve, but my brain really doesn't like relational database stuff, so thanks in advance to anyone who can help me untangle this. The closer to ELI5, the better!
## [8][Anyone using Hivemind, and how do you properly close all processes?](https://www.reddit.com/r/rails/comments/fz81bp/anyone_using_hivemind_and_how_do_you_properly/)
- url: https://www.reddit.com/r/rails/comments/fz81bp/anyone_using_hivemind_and_how_do_you_properly/
---
Hi, I stumbled upon [Hivemind](https://github.com/DarthSim/hivemind) when I was trying to setup Sidekiq and Redis for a project. After some setup I was able to get the processes running, and all seemed fine. But when I tried to shutdown the Hivemind process with Ctrl+c, I got error messages and I was not sure if this is the correct way to shut things down. I've searched around but couldn't find any information on it.

Here is my Procfile:

    redis:  redis-server /usr/local/etc/redis.conf
    worker: bundle exec sidekiq
    rails:  PORT=3000 bundle exec rails server

From what I can guess from the error messages (see below), the Redis server is shutdown first, but the Sidekiq is still running, so Sidekiq generates errors and finally gets shutdown. Is this normal behavior for Hivemind? Because I am not too comfortable with it... Ideally, I'd want to close rails and worker process first, then close the redis server. Should I just move to Overmind? Any help will be appreciated.

Output for the shutdown process:

    ^C
    rails  | Interrupting...
    redis  | Interrupting...
    worker | Interrupting...
    redis  | 65248:signal-handler (1586619108) Received SIGINT scheduling shutdown...
    rails  | [65249] - Gracefully shutting down workers...
    worker | pid=65247 tid=19in INFO: Shutting down
    worker | pid=65247 tid=19in INFO: Terminating quiet workers
    worker | pid=65247 tid=wuz INFO: Scheduler exiting...
    redis  | 65248 # User requested shutdown...
    ...
    redis  | 65248 # Redis is now ready to exit, bye bye...
    redis  | Process exited
    worker | pid=65247 tid=wuf ERROR: Error fetching job: Error connecting to Redis on localhost:6379 (Errno::ECONNREFUSED)
    ......
    worker | /Users//.rbenv/versions/2.7.0/lib/ruby/gems/2.7.0/gems/redis-4.1.3/lib/redis/client.rb:362:in `rescue in establish_connection': Error connecting to Redis on localhost:6379 (Errno::ECONNREFUSED) (Redis::CannotConnectError)
    worker | Process exited
## [9][To improve languages_controller for logged users](https://www.reddit.com/r/rails/comments/fz5hok/to_improve_languages_controller_for_logged_users/)
- url: https://www.reddit.com/r/rails/comments/fz5hok/to_improve_languages_controller_for_logged_users/
---
in languages\_controller.rb I **had** this

    def update
      u/language = Language.find(update_params[:id])
      redirect_to root_path(locale: u/language.locale)
    end

**I edited it with an if condition (**`if user_signed_in?`)**.** In this way, if the user is logged, in the homepage (views/language/edit) he can change the the language also in the "Settings Area".

    def update
      u/language = Language.find(update_params[:id])
      if user_signed_in?
        current_user.setting.language = u/language.id
        current_user.setting.update(language: u/language.id)
      end
      redirect_to root_path(locale: u/language.locale)
    end

But in this way I had an error if the user never choosed his language in the "Settings Area".

For this reason I (*temporarily*) solved adding `&amp;&amp; !current_user.setting.blank?`

The last version: 

    def update
      u/language = Language.find(update_params[:id])
      if user_signed_in? &amp;&amp; !current_user.setting.blank?
        current_user.setting.language = u/language.id
        current_user.setting.update(language: u/language.id)
      end
      redirect_to root_path(locale: u/language.locale)
    end

In this way it works, but I want to fix it.

**How to edit it to left the user to change the language (in his settings) also in the homepage?**

I don't know if I'm very clear.
## [10][Hi guys is there any Vscode extension where if I want to Hover on some function and when I click it, it takes me to the page where that function is calling from!](https://www.reddit.com/r/rails/comments/fyvx8g/hi_guys_is_there_any_vscode_extension_where_if_i/)
- url: https://www.reddit.com/r/rails/comments/fyvx8g/hi_guys_is_there_any_vscode_extension_where_if_i/
---

## [11][Any thoughts on Devise Token Auth vs. Devise Doorkeeper gems?](https://www.reddit.com/r/rails/comments/fyxqz3/any_thoughts_on_devise_token_auth_vs_devise/)
- url: https://www.reddit.com/r/rails/comments/fyxqz3/any_thoughts_on_devise_token_auth_vs_devise/
---
Trying to determine which gem or combination to use to allow an enterprise customer's server authenticate to both my web app and API. 

Any thoughts or things worth knowing?
## [12][Web Hosting](https://www.reddit.com/r/rails/comments/fyy7f1/web_hosting/)
- url: https://www.reddit.com/r/rails/comments/fyy7f1/web_hosting/
---
So... basically all sources I’ve seen use Heroku for hosting or “production/deployment.” Why not AWS, GoDaddy, WordPress, etc. what’s the trade off aside from the Heroku integration?
