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
## [2][Sudde Errno::EEXIST in Main#index error](https://www.reddit.com/r/rails/comments/ib7jvh/sudde_errnoeexist_in_mainindex_error/)
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
## [3][another content security policy question](https://www.reddit.com/r/rails/comments/ib6shh/another_content_security_policy_question/)
- url: https://www.reddit.com/r/rails/comments/ib6shh/another_content_security_policy_question/
---
Given the extensive use of dynamically set inline styling in most front end frameworks, exactly how does one go about making their front end comply with a safe content security policy?

Semantic UI's [modules](https://semantic-ui.com/modules/modal.html) are one example. Their modal changes the `body` height with an inline style when the modal appears. A pull request from two years ago aimed to fix the issue by changing some instances of `attr()` to `css()`, but this has had no impact on the framework's dependence on inline styling.

I don't want to use `unsafe_inline` for `style_src`!!!
## [4][What's the best way to implement filters? Gem recommendations? (Ransack?)](https://www.reddit.com/r/rails/comments/iauk5c/whats_the_best_way_to_implement_filters_gem/)
- url: https://www.reddit.com/r/rails/comments/iauk5c/whats_the_best_way_to_implement_filters_gem/
---
I've started programming in Rails 5 a few months ago. I'm developing an app that has a couple models like Technology, Skill Level, Role, etc. (where a Role requires a certain Technology and Skill Level to be covered by a User).

Basically what I need is to apply filters to the Roles index, so I can filter those Roles for which I have the required technology and skill level to apply to it, so to speak. So for example, I should be able to filter Roles that require Ruby on Rails, HTML and CSS technologies, and a skill level of Junior.

What would be the best way to implement something like this? I've been looking at a gem called Ransack but I'm not sure if it's what I need.

Any help would be appreciated.
## [5][Updating An Existing App?](https://www.reddit.com/r/rails/comments/iapuzt/updating_an_existing_app/)
- url: https://www.reddit.com/r/rails/comments/iapuzt/updating_an_existing_app/
---
Hey guys, just a quick question. I uploaded my initial version of my app, and to be honest it was to mainly see how deployment works.

How do I go about updating the app when I have a major change I want to upload to the server?

For Heroku, I used the CLI to do this, but for servers like Digital Ocean and AWS are different.

Any tips or workflow you guys use will be very helpful!

Thanks.
## [6][To Fork or To Clone?](https://www.reddit.com/r/rails/comments/iauid0/to_fork_or_to_clone/)
- url: https://www.reddit.com/r/rails/comments/iauid0/to_fork_or_to_clone/
---
Good morning guys, my desktop is being super slow so I want to transition to my laptop. My JS with Rails API project is saved on my desktop, so on my laptop, would I fork or clone the repository from github? Thanks in advance!
## [7][rails aborted! NoMethodError: undefined method `split' for nil:NilClass](https://www.reddit.com/r/rails/comments/iasufk/rails_aborted_nomethoderror_undefined_method/)
- url: https://www.reddit.com/r/rails/comments/iasufk/rails_aborted_nomethoderror_undefined_method/
---
Hey, I've been running into this weird issue while trying to seed my database.

Data in the format that I want. It's an array and I'm trying to iterate through it and create bunch of objects with it. It throws this weird error. I've checked all the similar previous questions but all of them are related to some config or gem issues.

[My code](https://i.stack.imgur.com/Fvbrg.png)

[The error](https://i.stack.imgur.com/HpAcI.png)

[Byebug before split method](https://i.stack.imgur.com/wLwgY.png)

[Byebug after split method](https://i.stack.imgur.com/fPXc1.png)

&amp;#x200B;

Any help would be greatly appreciated.

Thanks.
## [8][Managing usage for a commercial gem](https://www.reddit.com/r/rails/comments/iam6tf/managing_usage_for_a_commercial_gem/)
- url: https://www.reddit.com/r/rails/comments/iam6tf/managing_usage_for_a_commercial_gem/
---
How could I manage usage of a gem or engine as a product?

Paid Sidekiq for example is tied to threads, but I don't know how this is managed. 

It's loaded in the Gemfile so is referenced when bundled, yet Ruby is of course an interpreted and open language so even if it is reading other configs, that can all be manipulated in different ways at runtime, right? Is it more the honor system? (There is always trust involved somewhere, though this is handled to different degrees depending on the product.)

This is not meant to be the 'hack Sidekiq' thread, I am more curious how this could be done with a commercial product, not even with threads but with number of users for instance to prevent abuse.

[I tried Google but didn't find much, it returns whatever it feels like now.]
## [9][How do I configure my models in this situation?](https://www.reddit.com/r/rails/comments/iaov9r/how_do_i_configure_my_models_in_this_situation/)
- url: https://www.reddit.com/r/rails/comments/iaov9r/how_do_i_configure_my_models_in_this_situation/
---
I'm writing an email web app. The app uses email templates to create dynamic emails depending on who they are sent to.

An email template has variables. For example an email template would look like the following:

    EmailTemplate
    
    name: string
    subject: string
    body: string
    variables: [string]

and an instance would look like this:

    email_template
    
    name: "Cool email"
    subject: "Hi, FIRST_NAME,"
    body: "I see you are a POSITION at COMPANY...."
    variables: ["NAME", "POSITION", "COMPANY"]

Then a recipient model to the email would look something like this:

    Recipient
    name: "Some guy"
    email: "guy@email.com"
    variables: ??????
    
The *variables* attribute in EmailTemplate are dynamic so how do I set up the variables on the recipient model?
## [10][Efficient to run SQL command instead of ruby?](https://www.reddit.com/r/rails/comments/iaex30/efficient_to_run_sql_command_instead_of_ruby/)
- url: https://www.reddit.com/r/rails/comments/iaex30/efficient_to_run_sql_command_instead_of_ruby/
---
I've having trouble recreating the trigam `similarity` function in ruby (due to Asian characters in my text). Is the following a good idea in my rails code?

`sql = "SELECT similarity('#{src}', '#{targ}');"`

`tm_ratio = ActiveRecord::Base.connection.execute(sql).getvalue(0,0)`
## [11][TDD Course or tutorial](https://www.reddit.com/r/rails/comments/iaejjw/tdd_course_or_tutorial/)
- url: https://www.reddit.com/r/rails/comments/iaejjw/tdd_course_or_tutorial/
---
Hey y'all, I feel that I have a pretty decent grasp of ruby and rails and I want to learn TDD what are some good resources for learning TDD?

Thanks in advance
