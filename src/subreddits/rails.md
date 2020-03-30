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
## [2][best implementation for a model with an owner and users?](https://www.reddit.com/r/rails/comments/frkvhe/best_implementation_for_a_model_with_an_owner_and/)
- url: https://www.reddit.com/r/rails/comments/frkvhe/best_implementation_for_a_model_with_an_owner_and/
---
I have a Team model in my rails app. The team should have an owner. The owner should be able to invite other users to his/her team.

I'm using a Team as an "account" for my app. This account can have many users on it, but ultimately belongs to the owner of the team.

Therefore a team should also have many users. The users can interact with the account/app a certain way. But, the owner has more authorization with what he/she can do with the account/app because of his/her owner status.

Also, the owner should be able to assign different roles to any of the users, with each role having it's own unique set of authorizations.

Lastly, users may have their own teams, separate from the ones the members of. Meaning, that users can have multiple teams, being members of some, and being owners of others (although, it's likely they'll only ever be the owner of one, for convenience reasons).

&amp;#x200B;

How can I implement this? My initial thought was this:

**3 models**: Team, User, UserTeam

**Many-to-many assocation (has\_many through):** using user\_team as the middle man.

Lastly, I thought to keep track of roles on the user\_team records.

But this doesn't *feel* right. Ideally, I'd like to be able to write

    team.users # =&gt; returns all users including owner
    team.owner # =&gt; returns user that is the owner
    
    user.teams # =&gt; returns all teams user belongs to (including owned teams)
    user.owned_teams # =&gt; returns only the teams that the user owns.
    
    # and most importantly
    team.owner == @owner # =&gt; true
    team.users.includes?(@owner) # =&gt; true

because owners are technically users with an owner role.

What do you all think? Was my first thought on the right track? is there a better way? or should I just completely give up programming and become a musician?

Thanks!
## [3][Building a rails app that can upload pictures to an S3 bucket via CarrierWave and fog-aws, but it's giving me a Errno::EPIPE at /posts Broken pipe error](https://www.reddit.com/r/rails/comments/frmi1y/building_a_rails_app_that_can_upload_pictures_to/)
- url: https://www.reddit.com/r/rails/comments/frmi1y/building_a_rails_app_that_can_upload_pictures_to/
---
Edit: Note to anyone with the same issue in the future: 

SOLVED - Remember to specify the region in fog.credentials 

&amp;#x200B;

Hi there.

First of all this is my first post, I am really new at both Ruby and Rails, so I am glad to be here.

The last few days I've tried following a youtube tutorial creating an "instagram clone" cause I for my second rails project wanted to learn about implementing a postgres db, devise user authentication and also file hosting on AWS.

For the last 2 days I have been stuck on this issue where I cannot push anything to S3, and I feel like either the error is extremely vague or I am just not equipped to understand it. I tried using better\_errors and binding\_of\_caller to get a better feel for what's going on, but I am as stuck as can be.

things I've done:

I generated a CarrierWave uploader via the cli, and changed it to use :fog for storageI created the CarrierWave initializer and filled in the required fog credentials and directory as such:

&amp;#x200B;

https://preview.redd.it/65f2byqgjrp41.png?width=1310&amp;format=png&amp;auto=webp&amp;s=9837ef56e4d1a0baec84e83d9a96405a6d10d183

I have the variables for my credentials and bucket in a separate application.yml file living in the config directory

I then have a Post model in which I mount the uploader mentioned, and I have a controller for it. It is the controller specifically that throws the error

&amp;#x200B;

https://preview.redd.it/emqer83ckrp41.png?width=1490&amp;format=png&amp;auto=webp&amp;s=94ad7b7745fbf906b4d2e0136406410decadfa6d

Please let me know what other information is required cause right now I am feeling desperate and that's never a fun experience :)

(also, if someone would be willing to help me out with screensharing or similar, I would appreciate and welcome it)

Best regards

&amp;#x200B;

ps: forgot to mention that it looks like it times out when trying to upload/connect to aws, and that I verified the credentials and bucket with the awscli tool
## [4][Ruby 2.7.0 Warnings](https://www.reddit.com/r/rails/comments/frjwgd/ruby_270_warnings/)
- url: https://www.reddit.com/r/rails/comments/frjwgd/ruby_270_warnings/
---
I started using Ruby 2.7.0 for my projects. However there seems to be a lot of warning messages with Rails 6 libraries like activesupport. Will this be an issue moving forward?
## [5][From %&lt;a href=... to &lt;%=link_to in text.gsub!](https://www.reddit.com/r/rails/comments/frmzyn/from_a_href_to_link_to_in_textgsub/)
- url: https://www.reddit.com/r/rails/comments/frmzyn/from_a_href_to_link_to_in_textgsub/
---
Hi, I'm customizing my  **markdown redcarpet (**`class MarkdownRenderer &lt; Redcarpet::Render::HTML`**)**.

I found this part

      def paragraph(text)
        text.gsub!(/@(\w+)/) do |match|
          %(&lt;a href="/user/#{match[1..-1]}"&gt;#{match}&lt;/a&gt;)
        end

Can I replace `%(&lt;a href="/user/#{match[1..-1]}"&gt;#{match}&lt;/a&gt;)` using `&lt;%= link_to #{match}, user_path(#{match}) etc. etc.` ?

**How to do?** What is the right syntax?
## [6][Live podcasting in rails](https://www.reddit.com/r/rails/comments/frbetq/live_podcasting_in_rails/)
- url: https://www.reddit.com/r/rails/comments/frbetq/live_podcasting_in_rails/
---
Hello, I want to ask how I can build a live podcasting app on rails. Something like facebook's live video. Thanks!
## [7][Dtos in rails](https://www.reddit.com/r/rails/comments/frn6cd/dtos_in_rails/)
- url: https://www.reddit.com/r/rails/comments/frn6cd/dtos_in_rails/
---
Hi, im from c#. How to create dto in ruby on rails?
## [8][Is there any way to run a command automatically after ```rails s``` is run?](https://www.reddit.com/r/rails/comments/frhstu/is_there_any_way_to_run_a_command_automatically/)
- url: https://www.reddit.com/r/rails/comments/frhstu/is_there_any_way_to_run_a_command_automatically/
---
I want to ensure that whenever the server starts up it destroys all entries in the Conversations model ```Conversations.destroy_all```. This is to prevent some after effects up the server shutting down abruptly (there are some methods in a channels unsubscribe that need to run).
## [9][Rails Integration Test (Internal Response Error)](https://www.reddit.com/r/rails/comments/freuf4/rails_integration_test_internal_response_error/)
- url: https://www.reddit.com/r/rails/comments/freuf4/rails_integration_test_internal_response_error/
---
I am trying to write an integration test for a simple create route, and could use some help. Everything works just fine, including the redirect, but when I check to see if the post had a valid response, it is saying that it turned a 500 response rather than a 200. Here is the code:

Integration Test:

```
require 'test\_helper'  
class *CreateEventTest* &lt; ActionDispatch::IntegrationTest  
 fixtures :users  
 test "can login and create an event" do  
 get "/login"  
 assert\_response :success  
 post "/login", params: {user\_name: users(:one).user\_name, password: 'secret'}  
follow\_redirect!  
assert\_equal 200, status  
assert\_equal "/", path  


get "/events/new"  
 assert\_response :success  
 post '/events', params: { event: { name: "Event Title", description: "Description", location: "Search Results  
1600 Pennsylvania Ave NW, Washington, DC 20500", date\_from: DateTime.now + 10, date\_to: DateTime.now + 15, latitude: -35.000000, longitude: 100.000000}}  
assert\_response :redirect  
 follow\_redirect!  
assert\_response :created  
 end  
end

```

Events Controller:

```
# POST /events
  # POST /events.json
  def create
    @event = Event.new(event_params)
    respond_to do | format |
      if @event.save
        if(params.require(:event).key?("tags"))
          tags = Tag.find(params.require(:event)['tags'])
        else
          tags = Tag.find_by(name: "Other") # Set default tag if none was selected
        end
        @event.tags &lt;&lt; tags
        UserEventRelationship.create(event_id: @event.id, user_id: current_user.id, role_type_id: 0)
        format.html { redirect_to @event, notice: 'Event was successfully created.' }
        format.json { render json: @event, status: :created, location: @event }
      else
        format.html {render 'new'}
        format.json {render json: @event.errors, status: :unprocessable_entity }
      end
    end
  end

private

    def event_params
      params.require(:event).permit(:name, :date_from,
      :location, :date_to, :description, :picture, :tags, :latitude, :longitude)
    end

```

Here are my routes:
```
Rails.application.routes.draw do
  resources :users
  resources :events
  resources :account_activations, only: [:edit]
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
  get  '/events/new',  to: 'events#new'
  get  '/browse',  to: 'events#index'
  post '/events/:id/signup', to: 'events#register', as: 'register'
  delete '/events/:id/unregister', to: 'events#unregister', as: 'unregister'
  get '/home', to: 'home#index'

  get 'sessions/new'
  get  '/signup',  to: 'users#new'
  get    '/login',   to: 'sessions#new'
  get    '/profile',   to: 'users#profile'
  post   '/login',   to: 'sessions#create'
  delete '/logout',  to: 'sessions#destroy'
  post '/setUserLocation', to: 'users#setUserLocation'

  root 'home#index'


end

```
 
Here is the error I get:
```
Failure:
CreateEventTest#test_can_login_and_create_an_event [/Users/user/Documents/GitHub/my_app/test/integration/create_event_test.rb:21]:
Expected response to be a &lt;201: created&gt;, but was a &lt;500: Internal Server Error&gt;.
Expected: 201
  Actual: 500
```
## [10][Warning: The running version of Bundler is older than the version that created the lockfile](https://www.reddit.com/r/rails/comments/fr6ulk/warning_the_running_version_of_bundler_is_older/)
- url: https://www.reddit.com/r/rails/comments/fr6ulk/warning_the_running_version_of_bundler_is_older/
---
Hi all,  


I can't seem to get my system to use the version of bundler specified when I run "ruby -v".  


I keep getting the error "the running version of Bundler (2.1.2) is older than the version that created the lockfile (2.1.4)".  


They appear to match, but I can't seem to get rid of the error:  


    gem install bundler:2.1.4
    
    Successfully installed bundler-2.1.4
    1 gem installed
    
    ➜  project git:(upgrade_ruby) ✗ bundle -v
    Bundler version 2.1.4
    
    ➜  project git:(upgrade_ruby) ✗ bundle exec bundle -v
    Bundler version 2.1.4
    
    ➜  project git:(upgrade_ruby) ✗ rails s
    
    Warning: the running version of Bundler (2.1.2) is older than the version that created the lockfile (2.1.4). We suggest you to upgrade to the version that created the lockfile by running `gem install bundler:2.1.4`.

Can anyone help me to identify why my system is seemingly using a version of bundler other than that which is returned with bundle -v, and also how to change it to the latest?  


Thanks.
## [11][Failed to run multiple Rails apps with Unicorn + Nginx on single AWS EC2 Instance](https://www.reddit.com/r/rails/comments/frar43/failed_to_run_multiple_rails_apps_with_unicorn/)
- url: https://www.reddit.com/r/rails/comments/frar43/failed_to_run_multiple_rails_apps_with_unicorn/
---
I am going to run multiple Rails apps on AWS EC2 Instance with Unicorn and Nginx.

I could run one rails app on **mydomain.com**

So project will be on **mydomain.com/app1** and **mydomain.com/app2**

Projects are in `/home/ubuntu/work/app1` and `/home/ubuntu/work/app2`

***/etc/nginx/sites-available/default***

`upstream app1 {`

`server unix:/home/ubuntu/work/app1/shared/sockets/unicorn.sock fail_timeout=0;`

`}`

`upstream app2 {`

`server unix:/home/ubuntu/work/app2/shared/sockets/unicorn.sock fail_timeout=0;`

`}`

`server {`

`listen 80;`

`server_name localhost;`

`root /home/ubuntu/work;`

`access_log /home/ubuntu/work/log/nginx.access.log;`

`error_log /home/ubuntu/work/log/nginx.error.log;`

`location /app1/ {`

`root /home/ubuntu/work/app1/public;`

`rewrite ^/app1/(.*)$ /$1 break;`

`try_files /app1/$uri/index.html /app1/$uri.html /app1/$uri @app1;`

`}`

`location /app2/ {`

`root /home/ubuntu/work/app2/public;`

`rewrite ^/app2/(.*)$ /$1 break;`

`try_files /app2/$uri/index.html /app2/$uri.html /app2/$uri @app2;`

`}`

`location @app1 {`

`proxy_pass` [`http://app1`](http://app1)`;`

`proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;`

`proxy_set_header Host $http_host;`

`proxy_redirect off;`

`}`

`location @app2 {`

`proxy_pass` [`http://app2`](http://app2)`;`

`proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;`

`proxy_set_header Host $http_host;`

`proxy_redirect off;`

`}`

`error_page 500 502 503 504 /500.html;`

`client_max_body_size 4G;`

`keepalive_timeout 10;`

`}`

&amp;#x200B;

The error I am getting now. (Looks like the internal errors)

    We're sorry, but something went wrong. If you are the application owner check the logs for more information.  

I have checked these files to check logs.

    /home/ubuntu/work/log/nginx.error.log /home/ubuntu/work/fastland1/shared/log/unicorn.stderr.log  /home/ubuntu/work/fastland2/shared/log/unicorn.stderr.log  

But I can't find any log in nginx.error.log

And there were just warning in 2 unicorn.stderr.log files.

There were not any errors.

Anyone can help me?
