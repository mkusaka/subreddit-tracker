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
## [3][What is the most Rails-savvy javascript framework for a green-field project in 2020?](https://www.reddit.com/r/rails/comments/fi431t/what_is_the_most_railssavvy_javascript_framework/)
- url: https://www.reddit.com/r/rails/comments/fi431t/what_is_the_most_railssavvy_javascript_framework/
---
When I build a web app, I want it to feel like a coherent whole with a consistent philosophy. I want a framework that can easily talk to a Rails-provided RESTful API without customization. I want it to be as "omakase" as possible so that I can focus on the challenges of my domain rather than researching alternative technical approaches. Over the years, I've used jquery, backbone, angular, and most recently react. None of these experiences was satisfying. For example, I've found client-side rendering to be a deeply disappointing trade: about a 75% slow-down in project velocity for a little polish that, frankly, most features don't need. I say all this to clarify my preferences, not to get in an argument about why I'm wrong. So, fresh thoughts about a client-side framework that plays nice with Rails 6?
## [4][How do you test multi-step forms using rspec?](https://www.reddit.com/r/rails/comments/fid8ww/how_do_you_test_multistep_forms_using_rspec/)
- url: https://www.reddit.com/r/rails/comments/fid8ww/how_do_you_test_multistep_forms_using_rspec/
---
**If this question is too general and not rails specific please let me know and I'll post it elsewhere :)**

Hi everyone. I wanted to get your thoughts on how you methodically write rspec tests for multi-step (multi-page) forms.

I want to be able to achieve 100% coverage, but I don't know if that's feasible considering how many different possible input combinations my form can have, and how different the options available from fields in the next page can be depending on what was inputted in the first page.

I was wondering if any of you have done anything like this, and if you can give general advice in how you can tackle it. I am a beginner in rspec and so far have only managed to write tests for simple features like logging in, logging out, changing password, etc. without giving up...
## [5][[HELP] Creating a Inner join &amp; Left join SQL query using scopes](https://www.reddit.com/r/rails/comments/fi39jj/help_creating_a_inner_join_left_join_sql_query/)
- url: https://www.reddit.com/r/rails/comments/fi39jj/help_creating_a_inner_join_left_join_sql_query/
---
Hi guys, I need to create an query with inner join &amp; left join using scopes

What I have is something like this:

    scope :query_show, -&gt; { joins(:user, :document, requirements: [:status, :model] }

Because requirements may not exist on db for some users, this query may return empty. I don't want that behavior. I need a Left Join for the Requirement table and its association with the Inner Join. Something like this:

    scope :query_show, -&gt; { joins(:user, :document).select(...) }
    scope :requirements, -&gt; { includes(requirements: [:status, :model]).select(...) }
    
    User.query_show.requeriments # I don't think this work...
    

I know how to do it in raw SQL, but I'm trying to build with scopes. How can I achieve that? 

Thanks in advance.
## [6][undefined method 'method_name' for Controller:Module](https://www.reddit.com/r/rails/comments/fiavil/undefined_method_method_name_for_controllermodule/)
- url: https://www.reddit.com/r/rails/comments/fiavil/undefined_method_method_name_for_controllermodule/
---
Hi all, can't seem to resolve this error:  


    undefined method 'api_call' for CoffeeShopsController:Foursquare

Here's what my code looks like.  
Controller:

    class CoffeeShopsController &lt; ApplicationController
      include Foursquare
    
      def venue_search
        Foursquare.api_call(location, search)
      end
    
      ...
    
    emd

Module (app/lib/foursquare.rb)

    module Foursquare
      def api_call(location, search)
        ...
      end
    end

I've tried putting the following into application.rb which seemed to be the suggested solution on StackOverflow, but this has not worked:

    config.eager_load_paths += %W(#{config.root}/lib)
    config.autoload_paths &lt;&lt; Rails.root.join('lib')
    config.autoload_paths &lt;&lt; Rails.root.join('lib/notifier')
## [7][How I can build User invitation system?](https://www.reddit.com/r/rails/comments/fhmmor/how_i_can_build_user_invitation_system/)
- url: https://www.reddit.com/r/rails/comments/fhmmor/how_i_can_build_user_invitation_system/
---
Totally noob here, 

I am trying to build something similar to project management system. 

Admin where he can CRUD projects, reports, users. Also, add/invite user to projects (no sign up, I will seed him) 

Users can only  CRUD reports and view projects and  that they are added/invited into by admin. 

I have been looking online for a while, no luck! Most of the  tutorials doesn't help. 

Can someone please point me to the direction or help me with this.
## [8][why is rails more expensive to work with?](https://www.reddit.com/r/rails/comments/fhqtqi/why_is_rails_more_expensive_to_work_with/)
- url: https://www.reddit.com/r/rails/comments/fhqtqi/why_is_rails_more_expensive_to_work_with/
---
why is rails more expensive to work with? when a client starts working with a ruby developer, after once having worked with a php site, they often find

* ruby developers generally charge more (per hour)
* (usually) estimates to build a MVP with ruby on rails is higher

&gt;why is rails more expensive to work with?

More to the point, how does one justify the extra cost?
## [9][Does Rails auto-load all classes from all Gems?](https://www.reddit.com/r/rails/comments/fhhyz6/does_rails_autoload_all_classes_from_all_gems/)
- url: https://www.reddit.com/r/rails/comments/fhhyz6/does_rails_autoload_all_classes_from_all_gems/
---
Does Rails auto-load all classes from all Gems? So I can reference any class from any Gem's lib folder, and it will work?
## [10][How do I use Devise in normal AND API mode](https://www.reddit.com/r/rails/comments/fgwc24/how_do_i_use_devise_in_normal_and_api_mode/)
- url: https://www.reddit.com/r/rails/comments/fgwc24/how_do_i_use_devise_in_normal_and_api_mode/
---
Hi guys,  
I want to use Devise for creating users with mobile app via REST API but I also want use it in classic fashion.  
So in API I want to return JSON and appropriate HTTP codes. Also I want to issue JWT token for authentication for mobile app.  
On web I want normal devise functionality.  
How I should structure controllers, router and other stuff? I cant find any good resources about this.  
Thank you guys.
## [11][Devise and Client ID and Secret Key authentication?](https://www.reddit.com/r/rails/comments/fh15pj/devise_and_client_id_and_secret_key_authentication/)
- url: https://www.reddit.com/r/rails/comments/fh15pj/devise_and_client_id_and_secret_key_authentication/
---
Hi everyone,

I currently use Devise for authentication, both by normal rails and JWT auth. I'm wondering how people have been using devise auth for API calls that use client id and secret?

My initial thought is to create a new user role `api` and when the company authenticates via client id and secret to an endpoint, I sign them into the api user and return a JWT.

Am I doing this right?
## [12][[HELP] Test if a method is called after create/update method on controller](https://www.reddit.com/r/rails/comments/fh0ykp/help_test_if_a_method_is_called_after/)
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
