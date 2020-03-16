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
## [3][Standing out as a true Pro Rails Developer](https://www.reddit.com/r/rails/comments/fjedn3/standing_out_as_a_true_pro_rails_developer/)
- url: https://www.reddit.com/r/rails/comments/fjedn3/standing_out_as_a_true_pro_rails_developer/
---
Sorry for the long windedness, but here goes...

In my day job, I work as a JavaScript developer at a DotNet shop.  However, I've loved Rails - and Ruby - ever since I started playing around with them about 8 years ago.

People have always told me "don't waste your time on that, stick with dotnet (or Java), because that's where the money is."

But I honestly don't care about working in large enterprise organizations and I want to do the thing that I enjoy doing in programming.  Over the years I've done lots of Rails tutorials and made small projects.  Followed YouTube videos about various Rails subjects, and gone through books - like Sandi Metz excellent "Practical Object Oriented Design in Ruby".

Over the next few months I would like to transition into my first job as a Rails developer.  But I just have this hunch, some sneaky feeling that "real world" Rails development is much messier and different than what I'm seeing in the tutorials and books that I'm working with.

So I have a few questions for you Rails pros out there who have been in the trenches for a long time.

**What knowledge would I need to know to make me stand out from other Rails newbies?**  What do the tutorials **not** prepare you for?

For instance, in the DotNet world, we've had new developers who had experience working with Entity Framework - the dotnet equivalent of ActiveRecord.

But in most companies I've worked at, nobody actually uses Entity Framework, or any ORM for that matter.  The databases have generally been around for a long time and are super complicated and not very well designed.  To get the data you need, there are vast oceans of hand-made stored procedures and custom designed SQL - a lot of them pointed at OLAP cubes.

So a lot of newbies feel like there's this whole world they weren't prepared for.

In "real world" Rails development, are you really working with a SQLite database in Dev, updating migration files, and just running those against a PostGres or MySql database?  Or is it much more complicated than that?

Just looking for tips of what I need to study at a more advanced level that maybe the tutorials and books aren't preparing me for.
## [4][Date default value](https://www.reddit.com/r/rails/comments/fjk0mu/date_default_value/)
- url: https://www.reddit.com/r/rails/comments/fjk0mu/date_default_value/
---
So I have a model `Task` which has an attribute `deadline` of type `datetime`. I also have at some point in my views the following input:

    &lt;input type="date" id="start" name="task[deadline]" min="2020-01-01" max="2025-12-31" value="&lt;%= @task.deadline.strftime("%Y/%m/%d - %H:%M %p") %&gt;"&gt;

Well, apparently strftime does not work, as it shows `mm/dd/yyyy` instead of the corresponding date. I don't know how to fix this, but probably is a simple thing that right now I don't see.
## [5][SCSS changes not triggering recompile](https://www.reddit.com/r/rails/comments/fjfonm/scss_changes_not_triggering_recompile/)
- url: https://www.reddit.com/r/rails/comments/fjfonm/scss_changes_not_triggering_recompile/
---
Hi everyone, I'm having the a problem whereby my scss changes are not recompiling. Every time I make a change  I'm having to do the following:  


    rake assets:clean
    rake assets:precompile
    &lt;kill server&gt;
    rails s

  
I've tried all these suggested solutions I found online:  


    Deleting the contents of public/assets/
    
    # development.rb
    config.serve_static_assets = false
    
    Clearing browser cache

  
None of these things have worked.  


Any ideas? Thanks
## [6][Making Rails apps more beautiful? (using frontend technologies/frameworks)](https://www.reddit.com/r/rails/comments/fj5s94/making_rails_apps_more_beautiful_using_frontend/)
- url: https://www.reddit.com/r/rails/comments/fj5s94/making_rails_apps_more_beautiful_using_frontend/
---
Today I depolyed a rails app and what most of people tested say was : "Make it look more beautiful please". I was thinking about what are the ways to make a rails app look more and more beautiful? I did search and I found it's possible to use react or vue with webpacker on rails, but they still need some work to look beautiful. 

I'm not searching for "immediate" ways of course, but I look for something look better than rails default views on start. How should I say ... a more stylish way of views we can see.
## [7][organizing/managing Routes for an MVP](https://www.reddit.com/r/rails/comments/fj16w9/organizingmanaging_routes_for_an_mvp/)
- url: https://www.reddit.com/r/rails/comments/fj16w9/organizingmanaging_routes_for_an_mvp/
---
so i start working on an MVP with rails, 
i added devise to manage users and created my first model/controller 
which added to the route as resource so i have /model/:id/[edit/show] and /model/add 

but now i'm not sure if i have to change it to account/model for the user submission and /model for the rest of user to see the object 

so my question is when starting a project do you think about routes ? 

should i make a simple text file to list all the best names/titles for routes or i can keep the default resources and don't think/let the user think about them ?
## [8][ran rufus scheduler in my rails app, now when I later went back to the app is hanging in the browser](https://www.reddit.com/r/rails/comments/fj2oyy/ran_rufus_scheduler_in_my_rails_app_now_when_i/)
- url: https://www.reddit.com/r/rails/comments/fj2oyy/ran_rufus_scheduler_in_my_rails_app_now_when_i/
---
So i created a job, and I wanted to test it in console (eh, I'm new to it so don't know if this is bad) so I ran the job there.  I noticed when I went back later that my app doesn't load in browser, and I can't even ctrl+c out of the rails server log.  So I had to kill the rails pid and then start up rails s again, but it still doesn't load in the browser, though I see the initial GET for the homepage.  

Is there some process or other thing I need to do in rails console to stop rufus?  I had exited all my console sessions too, including the one in which I was trying to test a rufus job.  I've even commented out code, etc
## [9][Deploying to Heroku, do you have to worry about HTTP server?](https://www.reddit.com/r/rails/comments/fiub0i/deploying_to_heroku_do_you_have_to_worry_about/)
- url: https://www.reddit.com/r/rails/comments/fiub0i/deploying_to_heroku_do_you_have_to_worry_about/
---
Deploying to Heroku, do I have to worry about adding my own apache/nginx server?
## [10][GraphQL Rails queries](https://www.reddit.com/r/rails/comments/fimrmk/graphql_rails_queries/)
- url: https://www.reddit.com/r/rails/comments/fimrmk/graphql_rails_queries/
---
I have been messing with GraphQL in Rails, and have created a simple API so far. I have a couple of questions oh how extend this a little.

\- How would I go about creating a query with multiple variables, say looking up posts by {x} user on {y} date?  
\- How can I add a total returned count like how Shopify does their GraphQL queries.  
\- Also in should I be putting all my queries within the \`query\_type.rb\` file? The tutorial I was following wasn't that explicit.   


Any help would be much appreciated, Thank you.
## [11][Creating a react on rails app based on existing rails web-app?](https://www.reddit.com/r/rails/comments/fir46o/creating_a_react_on_rails_app_based_on_existing/)
- url: https://www.reddit.com/r/rails/comments/fir46o/creating_a_react_on_rails_app_based_on_existing/
---
So I have an existing web application but there's only so much web design I can do to make it useable on mobile.

Is react on rails the best way to go about creating an MVP for mobile that I can use to publish an app on the app/Google store?

And if so, would you:

a.) Recreate the entire project in react-rails and use this new project for both web and mobile.

b.) Create seperate project for mobile and isolate each project from one another.

c.) Create seperate project for mobile but utilise the backend of existing project.

Let me know your thoughts!
## [12][How do you test multi-step forms using rspec?](https://www.reddit.com/r/rails/comments/fid8ww/how_do_you_test_multistep_forms_using_rspec/)
- url: https://www.reddit.com/r/rails/comments/fid8ww/how_do_you_test_multistep_forms_using_rspec/
---
**If this question is too general and not rails specific please let me know and I'll post it elsewhere :)**

Hi everyone. I wanted to get your thoughts on how you methodically write rspec tests for multi-step (multi-page) forms.

I want to be able to achieve 100% coverage, but I don't know if that's feasible considering how many different possible input combinations my form can have, and how different the options available from fields in the next page can be depending on what was inputted in the first page.

I was wondering if any of you have done anything like this, and if you can give general advice in how you can tackle it. I am a beginner in rspec and so far have only managed to write tests for simple features like logging in, logging out, changing password, etc. without giving up...
