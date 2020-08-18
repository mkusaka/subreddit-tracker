# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/
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
## [2][Help me decide on a search solution (between algolia, searchkick and pg_search)](https://www.reddit.com/r/rails/comments/ibx0pz/help_me_decide_on_a_search_solution_between/)
- url: https://www.reddit.com/r/rails/comments/ibx0pz/help_me_decide_on_a_search_solution_between/
---
Hi guys

&amp;#x200B;

I'm looking for a good search solution for my application.

The project has rails as a graphql api and a react-native app as client.

What I am specifically looking for is a way to search for users but also users in specific contexts, like search a list of friends or make sure that blocked users are not showing up. Preferrably, while also providing some other data from the db, such as the amount of posts or other data, as part of the returned results.

Currently I was looking at the following solutions:

\- **Algolia:** Easy to implement to some extent but it seems a bit finicky having to add a list of friends/blocked users etc. as part of the saved indices in Algolia itself. Plus it seems more difficult to include other data from the db when needed. Plus (and this is the biggest downside) it's not free and cost quite a lot when an application scales. 

\- **Searchkick/Elasticsearch**: 2 months ago I asked a similar yet more vague question regarding search in Rails and some people pointed out to me how easy it is to use elasticsearch with searchkick. At the same time there were people who thought it was overkill and warned me of the complexities of Elastic Search. So I'm not sure what to believe but I feel this is not necessarily more complex than setting something similar up on Algolia?

\- **pg\_search**: This is a gem I just discovered today and although I'm not sure, it might be good enough for my purposes? It uses postgres full text search but also allows search scopes/facetting, which would be the answer when searching through someone's friends or blocked users.

I would love to know what your experience is. I'm not sure if it matters too much but this a production app and not a simple learning project, so having a scalable solution is needed.   


Thank you!
## [3][whenever cron tasks do not append to existing log file, but create .gz files instead](https://www.reddit.com/r/rails/comments/ibq7i8/whenever_cron_tasks_do_not_append_to_existing_log/)
- url: https://www.reddit.com/r/rails/comments/ibq7i8/whenever_cron_tasks_do_not_append_to_existing_log/
---
I am using the "whenever" gem to schedule cron tasks. They seem to work, but they do not append to my "cron\_log.log" file but instead go on to create "cron\_log.1.gz", "cron\_log.2.gz" etc.

Does anyone here who uses the whenever gem know why this is happening and how to fix it?

I followed these instructions for logging:  [https://github.com/javan/whenever/wiki/Output-redirection-aka-logging-your-cron-jobs](https://github.com/javan/whenever/wiki/Output-redirection-aka-logging-your-cron-jobs) 

     set :output, {:error =&gt; 'path/to/app/cron_error_log.log', :standard =&gt; 'path/to/app/cron_log.log'}
## [4][Rails 6 Bcrypt vs Passenger error on Production](https://www.reddit.com/r/rails/comments/iblyvt/rails_6_bcrypt_vs_passenger_error_on_production/)
- url: https://www.reddit.com/r/rails/comments/iblyvt/rails_6_bcrypt_vs_passenger_error_on_production/
---
After changing my ssl settings and restarting my Nginx, I started to receive the following error in the Passenger startup:

    Before process_action callback :ensure_user_signed_in has not been defined (ArgumentError) 

I am running Rails 6, Nginx/Passenger. The protected area is a single namespace only.

sessions\_controller

    class SessionsController &lt; NamespaceController
         skip_before_action :ensure_user_signed_in, only: [:new, :create]      
    
    # Present login form     
        def new 
        end      
    
    # Create Session     
        def create         
            user = User.where(email: params[:email]).first 
    
             if user &amp;&amp; user.authenticate(params[:password]) 
                 session[:user_id] = user.id             
                 redirect_to '/namespace/adminhub'
             else
                 redirect_to new_sessions_path, alert: 'Unable to authenticate'
             end
         end 
    
         # Logout
         def destroy
             reset_session
             redirect_to root_path
         end
      end 

Namespace\_controller

    class NamespaceController &lt; ApplicationController
         before_action :ensure_user_signed_in
    
          private
             def ensure_user_signed_in
                 unless current_user.present?
                 redirect_to new_sessions_path, alert: 'Must be signed in.'
             end
         end
    
          def current_user
             if session.has_key? :user_id
                 @current_user ||=User.find(session[:user_id])
             end
         end
         helper_method :current_user
       end 

I have attempted to undo my ssl changes in nginx and also to  restart passenger, neither seem to be the cause of this issue.  Interestingly, when I first pushed out the changes with the Bcrpyt, the  page loaded without issue and ran properly, as it does in my development  area. It was not until I had to restart the nginx process that this  error has come to light.
## [5][Downloading in JS after using send_file](https://www.reddit.com/r/rails/comments/ibttqr/downloading_in_js_after_using_send_file/)
- url: https://www.reddit.com/r/rails/comments/ibttqr/downloading_in_js_after_using_send_file/
---
I am a noob when it comes to web dev, pls no attack

So I am working on trying to download a file via JS + ruby and send\_file specifically. JS because I am only able to tell from browser which files to download. I have a method that works and now I'm just trying to figure out how stuff works with send\_file

Previously, my controller was basically:

    def download 
        obj = Obj[params[:id].to_i]
        if obj 
            render :json =&gt; {filename: obj.file_name, contents: obj.contents}
    end 

Now, I'm trying to do it with send\_file. What I have in my controller is basically:

    def download 
        send_file filename_of_curr_obj
    end 

I have in my routes:

    ...
    collection do 
        post :download
    end
    ...

Changed post to get for when I tried to switch to send\_file.

This is what I can tell I'm supposed to do from searching online. But I don't understand how it's supposed to link up in HTML/JS. Everything I've seen online usually stops at around here, as if the steps after this are obvious(?).

So far I have a download button:

    a :id =&gt; :download_link. :download =&gt; 'filename' do
        button :id =&gt; download_button do 
            text 'Download'
        end
    end 

(My company has action.html.rb files instead of html.erb which is everywhere online. I have never seen html.rb files online in the wild thus far.)

In the JS, I was able to do something like this that worked previously:

    $('#download_button').click(function() {
        $.post('...obj/download', {id: id}, function(data) {
            ...
            $('#download_link').attr('download', data.filename);
            ...
            //created url from Blob and data.contents
            $('#download_link').attr('href', url');
            ...
        });
    });

So basically it worked like this: on browser, some actions affect ID of obj's file, so on click, send the ID via POST request to download and then download file with the controller passing the filename and file contents.

Trying to redo it now with send\_file, I am confused about two main things:

1. How to get the filename for the download attr of the a tag?
2. How to get the URL for the href attr of the a tag? I thought the URL would be of the form 

&amp;#8203;

    .../obj/download/1

or 

    .../obj/1/download

But I am getting errors when I try doing a $.post request with either form saying the URL doesn't exist. I saw online that to have a custom action have the first url form, in the routes, I should have:

    ...
    member do 
        get :download
    end 
    collection do 
        get :download
    end
    ...

But this doesn't work.
## [6][RSpec model, request and system specs?](https://www.reddit.com/r/rails/comments/ibof0z/rspec_model_request_and_system_specs/)
- url: https://www.reddit.com/r/rails/comments/ibof0z/rspec_model_request_and_system_specs/
---
I've seen people recommend these three specs in Rails due to controller specs and feature specs being surpassed by request and system specs.

So are request specs doing the same as controller specs and test the methods of a controller individually? e.g. my request spec should test #index, #show, #new etc?

If so system specs should test the application from a users perspective right? E.g. the whole registration flow or creating a new post flow?

If this is the case for the above where do integration specs fall? Wouldn't request, integration and system specs all overlap and create duplicate tests?

I've tried to search but the books I've read or articles mainly point to controller and feature specs.
## [7][Is Cucumber and Capybara still being used?](https://www.reddit.com/r/rails/comments/ibenlh/is_cucumber_and_capybara_still_being_used/)
- url: https://www.reddit.com/r/rails/comments/ibenlh/is_cucumber_and_capybara_still_being_used/
---
Hi! I'm on a Rails course at EDX, but the course is not new. They talk about BDD and TDD, using Cucumber, Capybara and RSpec. Are these tools still in use? Is it something worth learning today?
## [8][model.save transaction committed but DB still empty (Devise)](https://www.reddit.com/r/rails/comments/ibozkp/modelsave_transaction_committed_but_db_still/)
- url: https://www.reddit.com/r/rails/comments/ibozkp/modelsave_transaction_committed_but_db_still/
---
I am trying to create a user from console but getting the following:

Console output:

    2.5.0 :021 &gt; User.create!(:name =&gt; "admin", :email =&gt; "admin@test.com", :password =&gt; "password", :password_confirmation =&gt; "password")
       (0.1ms)  begin transaction
      User Exists? (0.3ms)  SELECT 1 AS one FROM "users" WHERE "users"."email" = ? AND "users"."id" IS NOT NULL LIMIT ?  [["email", "admin@test.com"], ["LIMIT", 1]]
      User Update (0.2ms)  UPDATE "users" SET "name" = ?, "email" = ?, "encrypted_password" = ?, "updated_at" = ? WHERE "users"."id" IS NULL  [["name", "admin"], ["email", "admin@test.com"], ["encrypted_password", "$2a$12$X0aDvaz6P7ZM7L66Tyans.jmcGvDDXm8hk/ddS2Fms2tx9o9AMf5S"], ["updated_at", "2020-08-17 22:43:25.388300"]]
       (0.1ms)  commit transaction
     =&gt; #&lt;User id: nil, name: "admin", email: "admin@test.com", customer_id: 0, created_at: nil, updated_at: "2020-08-17 22:43:25"&gt; 
    2.5.0 :022 &gt; User.all
      User Load (0.2ms)  SELECT "users".* FROM "users" LIMIT ?  [["LIMIT", 11]]
     =&gt; #&lt;ActiveRecord::Relation []&gt; 
    

DeviseCreateUsers Migration

    class DeviseCreateUsers &lt; ActiveRecord::Migration[6.0]
      def change
        create_table :users do |t|
          ## Database authenticatable
          t.string :name,               null: false, default: ""
          t.string :email,              null: false, default: ""
          t.string :encrypted_password, null: false, default: ""
    
          t.integer :customer_id, default: 0
    
          ## Recoverable
          t.string   :reset_password_token
          t.datetime :reset_password_sent_at
    
          ## Rememberable
          t.datetime :remember_created_at
    
          ## Trackable
          t.integer  :sign_in_count, default: 0, null: false
          t.datetime :current_sign_in_at
          t.datetime :last_sign_in_at
          t.string   :current_sign_in_ip
          t.string   :last_sign_in_ip
    
          ## Lockable
          t.integer  :failed_attempts, default: 0 #, null: false # Only if lock strategy is :failed_attempts
          t.string   :unlock_token # Only if unlock strategy is :email or :both
          t.datetime :locked_at
    
    
          t.timestamps null: false
        end
    
        add_index :users, :email,                unique: true
        add_index :users, :reset_password_token, unique: true
      end
    end

&amp;#x200B;

app/models/user.rb

    class User &lt; ApplicationRecord
      devise :database_authenticatable,
             :timeoutable, :lockable, :rememberable, :trackable, :validatable
      validates :name, format: { with: /\A[a-zA-Z0-9]+\Z/ }
    end

Any ideas?  
Thank you
## [9][Greetings and respect to the pro here, please i need help with my rails app.](https://www.reddit.com/r/rails/comments/ibl77n/greetings_and_respect_to_the_pro_here_please_i/)
- url: https://www.reddit.com/r/rails/comments/ibl77n/greetings_and_respect_to_the_pro_here_please_i/
---
Firstly, am a newbie to rails. I love working with template since am not into ui/ux but more of back end Dev.

I've been able to add bootstrap, jquery and popper.js successfully but have a hell of issues linking the CSS/JAVASCRIPT of the html template am using.

I will be grateful if can get a step by step procedure or a link to any resource that can help me resolve this issue, thanks.
## [10][Which web server is best to run Kubernetes for Ruby on Rails?](https://www.reddit.com/r/rails/comments/ibhhij/which_web_server_is_best_to_run_kubernetes_for/)
- url: https://www.reddit.com/r/rails/comments/ibhhij/which_web_server_is_best_to_run_kubernetes_for/
---
Hi all

One of my favoured setups for deploying Rails on a static server is by using the combination of Phusion Passenger and Nginx.

However now I'm embarking on a quest into cloud computing, and have a question around which web server I might want to serve Rails with from within a Kubernetes pod.

From my understanding, the options are:

* Unicorn
* Puma
* Thin
* Passenger

The big questions for me is concurrency. If I go with a Puma web server, can it handle more than one request at a time? If ten users are on the site at once, do I need to auto-scale to ten containers?

I have done some research and digging so far, but haven't found many people talking about what web server technology they need inside their production container.

So, which web server is best to run inside Kubernetes for Ruby on Rails?
## [11][Sudde Errno::EEXIST in Main#index error](https://www.reddit.com/r/rails/comments/ib7jvh/sudde_errnoeexist_in_mainindex_error/)
- url: https://www.reddit.com/r/rails/comments/ib7jvh/sudde_errnoeexist_in_mainindex_error/
---
Hello. I'm creating my web dev portfolio in Rails using Docker. I got my app created and I was able to start adding content to the index.html.erb file. I restarted my server because my changes weren't popping up and I've had to do that once before when I changed my test of "Hello World!" to make sure the controller was connected properly to the view. I'm not sure I changed anything else. I tried googling and looked like maybe it had to do with the sprockets gem. I added it to my gemfile cause I didn't see it there. But it's still having this issue. Below is my output once my Rails server starts and tries to load the only view index.html.erb: 

    web_1  | Started GET "/" for 172.18.0.1 at 2020-08-17 04:20:23 +0000                                web_1  | Cannot render console from 172.18.0.1! Allowed networks: 127.0.0.0/127.255.255.255, ::1    
    web_1  | Processing by MainController#index as HTML                                                 web_1  |   Rendering main/index.html.erb within layouts/application                                 web_1  |   Rendered main/index.html.erb within layouts/application (Duration: 2.8ms | Allocations:  
    web_1  | /usr/local/bundle/gems/sprockets-rails-3.2.1/lib/sprockets/rails/helper.rb:355: warning: U  last argument as keyword parameters is deprecated; maybe ** should be added to the call            
    web_1  | /usr/local/bundle/gems/sprockets-4.0.2/lib/sprockets/base.rb:118: warning: The called meth is defined here                                                                                     web_1  | Completed 500 Internal Server Error in 200ms (ActiveRecord: 0.0ms | Allocations: 10332)    web_1  |  web_1  |  web_1  |  web_1  | ActionView::Template::Error (File exists @ dir_s_mkdir - /my-portfolio/tmp/cache/assets/sp v4.0.0/TI):                                                                                         web_1  |      5:     &lt;%= csrf_meta_tags %&gt;                                                          web_1  |      6:     &lt;%= csp_meta_tag %&gt;                                                            web_1  |      7:                                                                                    web_1  |      8:     &lt;%= stylesheet_link_tag 'application', media: 'all', 'data-turbolinks-track':   %&gt;                                                                                                 web_1  |      9:     &lt;%= javascript_pack_tag 'application', 'data-turbolinks-track': 'reload' %&gt;    web_1  |     10:                                                                             web_1  |     11:                                                                                    web_1  |  web_1  | app/views/layouts/application.html.erb:8   

Also, here's the Github repo:  [https://github.com/Kyle-Williamson-Dev/My-Portfolio](https://github.com/Kyle-Williamson-Dev/My-Portfolio) 

Also included an image of the error in the browser.

Let me know if I'm missing something or need to make any changes. 

I see the first warning where it says last argument as keyword parameters is depreciated and to add \*\*. Not sure how or where I would do that. Then I also see just where it says ActionView::Template::Error (File exists @ dir\_s\_mkdir - my-portfolio/tmp/cache/assets/sp v4.0.0/TI) 

Not sure where I'd find that to either delete that file or how I'd stop it from trying to recreate that file, if that's even what it's doing. 

Any help is greatly appreciated.
