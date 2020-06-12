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
## [2][Deploying subdirectory projects to Heroku](https://www.reddit.com/r/rails/comments/h7iazy/deploying_subdirectory_projects_to_heroku/)
- url: https://www.reddit.com/r/rails/comments/h7iazy/deploying_subdirectory_projects_to_heroku/
---
When you have your project services (API backend, frontend, background processing, etc.) in one monorepo you might struggle with a simple way of deploying them to Heroku independently. Here's an article containing a way of doing just that.

TL;DR: The easiest solution is 
```
git subtree push --prefix path/to/app-subdir heroku master
```

To get more details and a more complex solution that provides the ability to set up a deployment from CI and utilizing Heroku Review Apps you can read the article.

https://jtway.co/deploying-subdirectory-projects-to-heroku-f31ed65f3f2
## [3][MVC vs 3-tier architecture](https://www.reddit.com/r/rails/comments/h17sez/mvc_vs_3tier_architecture/)
- url: https://www.reddit.com/r/rails/comments/h17sez/mvc_vs_3tier_architecture/
---
**So this has always been a debatable topic.**  
**Approach 1:** We use the classic rails MVC architecture with Server Side rendering.  
**Approach 2:** We use client-side rendering, with rails working as a backend providing the JSON response and client to be made using react or any frontend framework.

  
I would love to get some insights and have a discussion on what are the pros and cons of both these approaches.  


Some of my thoughts are:  
\* Using client-side rendering restricts us to leverage the magic provided by rails. (form auto completions, an automatic mapping between endpoint and controller methods)  
\* Client-side rendering makes the application a bit lighter, as the server is not responsible for rendering the UI for every instance.  
\* I personally am not very comfortable in Js, so using Rails MVC is always an easy way out for me.  


Let's discuss this!
## [4][Can I use a conditional to toggle between rendered partials?](https://www.reddit.com/r/rails/comments/h79ljs/can_i_use_a_conditional_to_toggle_between/)
- url: https://www.reddit.com/r/rails/comments/h79ljs/can_i_use_a_conditional_to_toggle_between/
---
I'm trying to render a partial/hide another partial on a page when a button is clicked. Is this possible with a conditional? Or possible at all? Something like:   


    -if @toggle == 'show'
      = render :partial =&gt; "partial_one"
    
    -elsif @toggle == 'hide'
      = render :partial =&gt; "partial_two"
## [5][Implementing search, should I go for Algolia or is Rails/Postgres enough?](https://www.reddit.com/r/rails/comments/h14rpx/implementing_search_should_i_go_for_algolia_or_is/)
- url: https://www.reddit.com/r/rails/comments/h14rpx/implementing_search_should_i_go_for_algolia_or_is/
---
As the title says, I am implementing search functionality on a social network project.

I want it to feel snappy and provide good suggestions on typing and since I just came back to Rails after a long hiatus in javascript land, I'm not sure if Rails provides something for this out of the box or if should use something like Algolia which is a search-as-a-service to save me the headache.

Any ideas/feeback are welcome!

&amp;#x200B;

Regards
## [6][Appropriate way to validate association](https://www.reddit.com/r/rails/comments/h128cd/appropriate_way_to_validate_association/)
- url: https://www.reddit.com/r/rails/comments/h128cd/appropriate_way_to_validate_association/
---
I'm building a Rails API app that deals with orchestrating Users within a Company, and I want to check my logic for attaching Settings to a User.

The idea is, an admin at a Company will create a Settings entry in the database, and then assign that to a bunch of users.

&amp;#x200B;

Because this is an API, I have one call to create a Setting, one call to create a User, and one call to assign a Setting to a User.

&amp;#x200B;

Here's what the User controller looks like:

    class UsersController &lt; ApplicationController
    
      def set_setting
        if @user.update(setting: Setting.find(params[:setting_id]))
          render :show
        else
          render json: @user.errors, status: :unprocessable_entity
        end
      end
    
    end

And here's how I'm doing the validation in the User model:

    class User &lt; ApplicationRecord
      belongs_to :organisation
      has_one :user_setting
      has_one :setting, through: :user_setting
    
      validate :setting_in_organisation, on: :update
    
      private
    
        def setting_in_organisation
          errors.add(:setting_id, "You can't assign a setting that's not part of your organisation") unless setting.organisation == organisation
        end
    end

So, my question is, is this a good way to approach the issue while sticking to Rails ethos and DRY principals? I think it's important to highlight the views for this app are handled in a different repo. And because it's an API I definitely want a separate call to assign a Setting to a User, rather than chucking everything in user\_params.

What's everyone's thoughts?
## [7][Navbar and Topbar in application.html.erb vs individual files](https://www.reddit.com/r/rails/comments/h16k7s/navbar_and_topbar_in_applicationhtmlerb_vs/)
- url: https://www.reddit.com/r/rails/comments/h16k7s/navbar_and_topbar_in_applicationhtmlerb_vs/
---
Hello everyone, I just came across a project where the navigation bar and topbar are being rendered on every page (thus fetching the username on every reload). I was thinking about adding them to application.html.erb instead.   
I was just wondering if this approach has any kind of downfalls, in terms of scalability or performance.  


I know the answer seems a bit obvious, but just trying to get some more insights.
## [8][[HELP] How to show specific jsonapi-resources attributes scheme on Rswag's swagger ui?](https://www.reddit.com/r/rails/comments/h77fda/help_how_to_show_specific_jsonapiresources/)
- url: https://www.reddit.com/r/rails/comments/h77fda/help_how_to_show_specific_jsonapiresources/
---
Hi everyone, I have a json api using rswag to create the features/resources/endpoint tests and generate the swagger ui at the same time. The lib is amazing, but I'm having a small problem: Because I'm using jsonapi-resources I downloaded the jsonschema for jsonapi and use it as my "base schema". I want to show it as a response schema for every endpoint, but having the "attributes" item specified to that specific endpoint. I tried using allOf [baseSchema, specificSchema] , but I didn't have much success. I am trying to create the schema for "data" dynamically but It feels like I am overreaching and doing something that otherwise should be rather simple. 

Has anyone used both these tools at the same time?

How did you handle the response format of the jsonapi in the swagger ui?

PS: I've been to their GitHub continuously for the past days, but it seems like the problem is related to how I organize and compose the schemas using the rswag lib.

Thanks a lot in advance for any help! (I am so glad that I found this sub on my first days of Ruby/rails!)
## [9][How do I preserve/pass URL parameters with render](https://www.reddit.com/r/rails/comments/h17rrp/how_do_i_preservepass_url_parameters_with_render/)
- url: https://www.reddit.com/r/rails/comments/h17rrp/how_do_i_preservepass_url_parameters_with_render/
---
Fairly run-of-the mill configuration. Only thing is that from the index view I pass a parameter in my URL on the :new action.

    &lt;%= link_to 'Add Model', new_my_model_path(utm: @utm), class: 'btn btn-outline-dark btn-large mb-3', role: 'button' %&gt;

My create action is fairly standard:

    def create
      respond_to do |format|
        if @my_model.save
          format.html { redirect_to my_models_path(utm: @my_model.reference_id), notice: 'Model was successfully created.' }
          format.json { render :show, status: :created, location: my_modelss_path(utm: @my_model.reference_id) }
        else
          format.html { render :new }
          format.json { render json: @my_model.errors, status: :unprocessable_entity }
        end
      end
    end

Problem is, on that 'else', I lose the url parameter (which I want to continue to pass).

I'm guessing that I have two options:

1. Discover a little-known way to maintain url parameters with 'render'
2. Pass the data in a different way (pass the instance variable?)
3. Ensure we never hit 'else'

Obviously, I'm hoping there is a valid solution to 1, with the intention of later updating to 2. But I haven't been able to find a way yet.

Would love some input on my possible path(s) forward.

\*Edited for code block after foolishly using inline code...
## [10][What are you using for requests?](https://www.reddit.com/r/rails/comments/h0vva7/what_are_you_using_for_requests/)
- url: https://www.reddit.com/r/rails/comments/h0vva7/what_are_you_using_for_requests/
---
A lot of times API's won't have Ruby SDK's or gems.

I'm wondering what you all do in these cases...

&amp;#x200B;

I see some API's do:

    require 'uri' 
    require 'net/http' 
    require 'openssl' 

What is your normal procedure for doing this?

&amp;#x200B;

Basically, something similar to Pythons requests library
## [11][Having a hard time with saving a date with date_select](https://www.reddit.com/r/rails/comments/h11z2p/having_a_hard_time_with_saving_a_date_with_date/)
- url: https://www.reddit.com/r/rails/comments/h11z2p/having_a_hard_time_with_saving_a_date_with_date/
---
Trying to save a date of birth using date\_select.  I'm kinda lost.  Any advice would be greatly appreciated!

&amp;#x200B;

schema:

    t.date "age"

&amp;#x200B;

form:

    &lt;%= f.date_select :age, :start_year=&gt;1900,:end_year=&gt;2020 %&gt;

&amp;#x200B;

when saved:

    Child Create (0.9ms)  INSERT INTO "children" ("age", "created_at", "updated_at") VALUES (?, ?, ?, ?, ?, ?, ?)  [["age", "2012-03-02"], ["created_at", "2020-06-11 15:43:05.548019"], ["updated_at", "2020-06-11 15:43:05.548019"]]

&amp;#x200B;

when i try to access the age:

    from (pry):17:in `create'
    [8] pry(#&lt;ChildrenController&gt;)&gt; @child.age
    =&gt; Fri, 02 Mar 2012

edit: I was hoping it would return what it appeared to save:  "2012-03-02" 

but should it be a string? 

&amp;#x200B;
