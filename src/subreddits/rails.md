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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/fgx7fz/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/fgx7fz/personal_projects_show_off_your_own_project_andor/
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
## [3][An architecture for a chat system that interacts with a telegram bot](https://www.reddit.com/r/rails/comments/fh652g/an_architecture_for_a_chat_system_that_interacts/)
- url: https://www.reddit.com/r/rails/comments/fh652g/an_architecture_for_a_chat_system_that_interacts/
---
Hey, guys!

As you can see in the image, I'm building a chat system (Just me interacting with a telegram bot)

Right now my architecture is the following:

Conversation(name: string) -&gt; has\_many :messages

Message(body: string, conversation\_id: integer) -&gt; belongs\_to :conversation

The main problem is saving the telegram bot messages because I need them to appear as history, so I have a few questions

Should I save the bot messages in the database at all? or use the method 'get\_updates' from telegram bot API and make them appear into the chat using the DateTime, for example, to order the messages correctly? I'm still trying to figure the best architecture for this issue

https://preview.redd.it/xbkutozoi4m41.png?width=800&amp;format=png&amp;auto=webp&amp;s=c68400833c2835a8ae7d2615427de88ba42cbebe
## [4][How do I use Devise in normal AND API mode](https://www.reddit.com/r/rails/comments/fgwc24/how_do_i_use_devise_in_normal_and_api_mode/)
- url: https://www.reddit.com/r/rails/comments/fgwc24/how_do_i_use_devise_in_normal_and_api_mode/
---
Hi guys,  
I want to use Devise for creating users with mobile app via REST API but I also want use it in classic fashion.  
So in API I want to return JSON and appropriate HTTP codes. Also I want to issue JWT token for authentication for mobile app.  
On web I want normal devise functionality.  
How I should structure controllers, router and other stuff? I cant find any good resources about this.  
Thank you guys.
## [5][Devise and Client ID and Secret Key authentication?](https://www.reddit.com/r/rails/comments/fh15pj/devise_and_client_id_and_secret_key_authentication/)
- url: https://www.reddit.com/r/rails/comments/fh15pj/devise_and_client_id_and_secret_key_authentication/
---
Hi everyone,

I currently use Devise for authentication, both by normal rails and JWT auth. I'm wondering how people have been using devise auth for API calls that use client id and secret?

My initial thought is to create a new user role `api` and when the company authenticates via client id and secret to an endpoint, I sign them into the api user and return a JWT.

Am I doing this right?
## [6][[HELP] Test if a method is called after create/update method on controller](https://www.reddit.com/r/rails/comments/fh0ykp/help_test_if_a_method_is_called_after/)
- url: https://www.reddit.com/r/rails/comments/fh0ykp/help_test_if_a_method_is_called_after/
---
Hi guys, I have the following controller

    #controller
    
    def create
      @obj = ...
      if @obj.save
        
        @obj.run_job
        
        ...
      else 
        ...
      end
    end
    
    def update
      @obj = ...
      if @obj.update(...)
        
        @obj.update_job
        
        ...
      else 
        ...
      end
    end

I would like to test if the methods run/update\_jobs are been called using Rspec's requests test. I already tested these methods with Rspec's model tests, I just I want to check now if they are been called on controller. Is there a simple way to do that? I try to use rspec receive method, but didn't work.

Thanks in advance.

**# EDIT**

The test would be something like this

    # request test
    
    it "run_jobs should be called on create" do
      sign_in_user
             
      document = FactoryBot.attributes_for(:user_document)
      expect{
        post user_documents_path, xhr: true, :params =&gt; {"user_document" =&gt; { ... }}
      }.to change{ User::Document.count }.by(1)
    
      expect(document).to receive(:run_jobs)
    end

I got the error message: "document ... does not implement: run\_jobs"

That error message is because document is a FactoryBot object, which doesn't correspond to the document object responsible for calling run\_jobs method.
## [7][Multiple shopping carts for a single user](https://www.reddit.com/r/rails/comments/fgy9kw/multiple_shopping_carts_for_a_single_user/)
- url: https://www.reddit.com/r/rails/comments/fgy9kw/multiple_shopping_carts_for_a_single_user/
---
I'm currently working on an app where you can book beauty and related services. Right now, I'm adding the shopping cart functionality and I'm facing an issue deciding what to do with it.

The issue is that the business requirement doesn't allow any user to book services from more than one venue at a time but they want each specific venue where the user has added a service to the shopping cart to "remember" which service was added.

So I can't have a traditional "Cart" model where you can add any item as long as the relations are set up because the cart is kinda venue specific.

I have thought about having a cart with "venue\_id" and "user\_id" created each time the user adds a product to the shopping cart and have a "valid\_up\_to" field so I can make a cron job to delete the invalid carts.

I think that this approach can solve the issue but I can't avoid to feel it as a hacky solution.

What do you think would the best approach here?
## [8][Need help designing](https://www.reddit.com/r/rails/comments/fh2sgh/need_help_designing/)
- url: https://www.reddit.com/r/rails/comments/fh2sgh/need_help_designing/
---
I want to make a web application similar to a ToDo list. Thus, I have two models: projects and tasks. Projects are composed of several tasks, but tasks may be created without belonging to any project.

I don't know how to implement this association. I think project `has_many` tasks, but tasks don't necessarily `belongs_to` project, because they may be free.

I've tried to implement it with an optional belongs\_to, but when calling `@project.tasks` it throws `NoMethodError`, so I really don't know how to access the tasks belonging to a specific project.
## [9][Question: Rails 5.2 Testing Javascript and Manipulating the Session](https://www.reddit.com/r/rails/comments/fgmzfk/question_rails_52_testing_javascript_and/)
- url: https://www.reddit.com/r/rails/comments/fgmzfk/question_rails_52_testing_javascript_and/
---
Hey everyone,

I've been at this problem for too long, so I'm reaching out for a steer on it.  
I have some code that works in the browser, but I can't work out how to write a test for it.

Basically, I have a Bootstrap 4 modal in a view. There's some Coffeescript that will open the modal if the user has come from a particular HTTP\_REFERRER.

I cannot find a test framework that supports JS as well as supporting manipulating request headers. I've looked at MiniTest, RSpec and Capybara.

I've seen a StackOverflow post where BrowserMob was suggested - but it seems to have been abandoned, and I'm loathe to add another dependency for this one feature.

Any suggestions? I'm probably overlooking something obvious.  The aim is be able to manipulate the HTTP\_REFERRER header in a test to cause the JS to fire and show the modal.

Thanks in advance.
## [10][Deployment Questions](https://www.reddit.com/r/rails/comments/fgnxqg/deployment_questions/)
- url: https://www.reddit.com/r/rails/comments/fgnxqg/deployment_questions/
---
Hello, I have a few questions about your preferred way to deploy. I typically just launch an EC2, ssh in, configure capistrano, nginx, route 53 &amp; deploy.

I want to explore some other areas of deployment to try and speed this process up without increasing monthly cost too much. What i am thinking is to start using docker, deploy with beanstalk, RDS for a DB but i have one last area to figure out. Redis.   


ElastiCache is the price of a small EC2. Using this service could double the bill which isn't the goal. Maybe i don't understand the pricing the model. If i turn one on, am i charged $0.05/hr like a EC2 or is it on-demand? My other thought is to put redis into the docker container. My fear is something like this: 50 jobs in queue, deploy, queue is gone. Is this possible? Maybe i don't understand how the deployment works nor redis? Would the queue be gone if i put redis in the container?   


Regardless, what's your preferred way deploying RoR while not being backed by a corporate budget? Open to exploring other options!
## [11][Rails with nginx/unicorn - 2gb per process??](https://www.reddit.com/r/rails/comments/fgkwuf/rails_with_nginxunicorn_2gb_per_process/)
- url: https://www.reddit.com/r/rails/comments/fgkwuf/rails_with_nginxunicorn_2gb_per_process/
---
Hi all,

My unicorn processes are taking up about 2GB each of memory. The internet tells me this is rather abnormal and high, but I'm not sure what size app people are benchmarking against. 

I have a production consumer web app which does about a million people per month in traffic. My AWS instance is 16gb memory. I have my unicorn config set to `worker_processes 5`. 

When I run a command like `ps -aux --sort -rss | head -n 10` the top six are the unicorn master and 5 workers, each taking up between 2 and 2.5GB of memory. 

I've already gone down several rabbit holes here, but before following any one path I would love to get a sense if this is at all normal for the kind of traffic my website does. 

Thanks!
## [12][Migrate a Rails Project from Paperclip to ActiveStorage](https://www.reddit.com/r/rails/comments/fgiswf/migrate_a_rails_project_from_paperclip_to/)
- url: https://www.reddit.com/r/rails/comments/fgiswf/migrate_a_rails_project_from_paperclip_to/
---
I've recently started a blog and the first article I wrote was based around my experience migrating one of my personal projects from Paperclip to ActiveStorage.  Sharing here in hopes someone finds it useful!

[https://www.kevin-custer.com/blog/migrate-a-rails-project-from-paperclip-to-active-storage/](https://www.kevin-custer.com/blog/migrate-a-rails-project-from-paperclip-to-active-storage/)
