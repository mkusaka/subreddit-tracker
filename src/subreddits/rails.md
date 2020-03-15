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
## [3][organizing/managing Routes for an MVP](https://www.reddit.com/r/rails/comments/fj16w9/organizingmanaging_routes_for_an_mvp/)
- url: https://www.reddit.com/r/rails/comments/fj16w9/organizingmanaging_routes_for_an_mvp/
---
so i start working on an MVP with rails, 
i added devise to manage users and created my first model/controller 
which added to the route as resource so i have /model/:id/[edit/show] and /model/add 

but now i'm not sure if i have to change it to account/model for the user submission and /model for the rest of user to see the object 

so my question is when starting a project do you think about routes ? 

should i make a simple text file to list all the best names/titles for routes or i can keep the default resources and don't think/let the user think about them ?
## [4][Deploying to Heroku, do you have to worry about HTTP server?](https://www.reddit.com/r/rails/comments/fiub0i/deploying_to_heroku_do_you_have_to_worry_about/)
- url: https://www.reddit.com/r/rails/comments/fiub0i/deploying_to_heroku_do_you_have_to_worry_about/
---
Deploying to Heroku, do I have to worry about adding my own apache/nginx server?
## [5][GraphQL Rails queries](https://www.reddit.com/r/rails/comments/fimrmk/graphql_rails_queries/)
- url: https://www.reddit.com/r/rails/comments/fimrmk/graphql_rails_queries/
---
I have been messing with GraphQL in Rails, and have created a simple API so far. I have a couple of questions oh how extend this a little.

\- How would I go about creating a query with multiple variables, say looking up posts by {x} user on {y} date?  
\- How can I add a total returned count like how Shopify does their GraphQL queries.  
\- Also in should I be putting all my queries within the \`query\_type.rb\` file? The tutorial I was following wasn't that explicit.   


Any help would be much appreciated, Thank you.
## [6][Creating a react on rails app based on existing rails web-app?](https://www.reddit.com/r/rails/comments/fir46o/creating_a_react_on_rails_app_based_on_existing/)
- url: https://www.reddit.com/r/rails/comments/fir46o/creating_a_react_on_rails_app_based_on_existing/
---
So I have an existing web application but there's only so much web design I can do to make it useable on mobile.

Is react on rails the best way to go about creating an MVP for mobile that I can use to publish an app on the app/Google store?

And if so, would you:

a.) Recreate the entire project in react-rails and use this new project for both web and mobile.

b.) Create seperate project for mobile and isolate each project from one another.

c.) Create seperate project for mobile but utilise the backend of existing project.

Let me know your thoughts!
## [7][How do you test multi-step forms using rspec?](https://www.reddit.com/r/rails/comments/fid8ww/how_do_you_test_multistep_forms_using_rspec/)
- url: https://www.reddit.com/r/rails/comments/fid8ww/how_do_you_test_multistep_forms_using_rspec/
---
**If this question is too general and not rails specific please let me know and I'll post it elsewhere :)**

Hi everyone. I wanted to get your thoughts on how you methodically write rspec tests for multi-step (multi-page) forms.

I want to be able to achieve 100% coverage, but I don't know if that's feasible considering how many different possible input combinations my form can have, and how different the options available from fields in the next page can be depending on what was inputted in the first page.

I was wondering if any of you have done anything like this, and if you can give general advice in how you can tackle it. I am a beginner in rspec and so far have only managed to write tests for simple features like logging in, logging out, changing password, etc. without giving up...
## [8][What is the most Rails-savvy javascript framework for a green-field project in 2020?](https://www.reddit.com/r/rails/comments/fi431t/what_is_the_most_railssavvy_javascript_framework/)
- url: https://www.reddit.com/r/rails/comments/fi431t/what_is_the_most_railssavvy_javascript_framework/
---
When I build a web app, I want it to feel like a coherent whole with a consistent philosophy. I want a framework that can easily talk to a Rails-provided RESTful API without customization. I want it to be as "omakase" as possible so that I can focus on the challenges of my domain rather than researching alternative technical approaches. Over the years, I've used jquery, backbone, angular, and most recently react. None of these experiences was satisfying. For example, I've found client-side rendering to be a deeply disappointing trade: about a 75% slow-down in project velocity for a little polish that, frankly, most features don't need. I say all this to clarify my preferences, not to get in an argument about why I'm wrong. So, fresh thoughts about a client-side framework that plays nice with Rails 6?
## [9][[HELP] Creating a Inner join &amp; Left join SQL query using scopes](https://www.reddit.com/r/rails/comments/fi39jj/help_creating_a_inner_join_left_join_sql_query/)
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
## [10][undefined method 'method_name' for Controller:Module](https://www.reddit.com/r/rails/comments/fiavil/undefined_method_method_name_for_controllermodule/)
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
## [11][How I can build User invitation system?](https://www.reddit.com/r/rails/comments/fhmmor/how_i_can_build_user_invitation_system/)
- url: https://www.reddit.com/r/rails/comments/fhmmor/how_i_can_build_user_invitation_system/
---
Totally noob here, 

I am trying to build something similar to project management system. 

Admin where he can CRUD projects, reports, users. Also, add/invite user to projects (no sign up, I will seed him) 

Users can only  CRUD reports and view projects and  that they are added/invited into by admin. 

I have been looking online for a while, no luck! Most of the  tutorials doesn't help. 

Can someone please point me to the direction or help me with this.
## [12][why is rails more expensive to work with?](https://www.reddit.com/r/rails/comments/fhqtqi/why_is_rails_more_expensive_to_work_with/)
- url: https://www.reddit.com/r/rails/comments/fhqtqi/why_is_rails_more_expensive_to_work_with/
---
why is rails more expensive to work with? when a client starts working with a ruby developer, after once having worked with a php site, they often find

* ruby developers generally charge more (per hour)
* (usually) estimates to build a MVP with ruby on rails is higher

&gt;why is rails more expensive to work with?

More to the point, how does one justify the extra cost?
