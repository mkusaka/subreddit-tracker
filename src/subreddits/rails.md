# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/
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
## [3][Moving to Rails from the Front End](https://www.reddit.com/r/rails/comments/j7llb1/moving_to_rails_from_the_front_end/)
- url: https://www.reddit.com/r/rails/comments/j7llb1/moving_to_rails_from_the_front_end/
---
Hello friends! I'm currently a senior front end developer. I've been using Rails in personal projects for a little over 3 years now and I'm reaching out to you all for some advice. I'm finding that I have greatly enjoyed working on the server side and have fallen in love with Ruby and Rails and I'd like to attempt a pivot in my career and move to a position that is more focused on those technologies. I have worked on the server side professionally with node but it's been in a very limited capacity so it's hard to point to that experience as some kind of proof.

I understand that the competition is fierce and that these domains can be pretty different but I'm hoping someone can help me think a little creatively about moving into this awesome community.
## [4][React/Rails + devise question: how the F*** do I render current_user data in my React front-end?](https://www.reddit.com/r/rails/comments/j7plfk/reactrails_devise_question_how_the_f_do_i_render/)
- url: https://www.reddit.com/r/rails/comments/j7plfk/reactrails_devise_question_how_the_f_do_i_render/
---
Need some help! 

I am trying to render my current\_user data on my React front-end. 

The JSON renders fine directly in the browser at localhost:3000/current\_user (a custom route I wrote), but when I try to fetch that data in react, I get a null response. I am using devise+google\_oauth2 with a Rails API. The authentication is working, user\_signed\_in? returns true, current user checks out, all of that.

 I have been googling for hours and haven't found anything that works. I tried installing the devise-token-authentication gem and overwriting the standard devise current\_user with this method in my application controller like so:

`def current_user`  
  `token = request.headers["Authorization"].to_s`  
  `User.find_for_database_authentication(authentication_token: token)`  
`end` 

but that didn't work either (stack overflow solution). I am using the `before_action :authenticate_user!` in my custom controller, which seems to be doing something as I get a 401 response without it. No matter how I try to tackle this, the JSON response from my fetch returns null. If anyone can help. I would greatly appreciate it! I'm assuming it has something to do with the request from my front-end not being authenticated but I can't figure this one out for the life of me...
## [5][Storing user settings in rails 5 API](https://www.reddit.com/r/rails/comments/j79doh/storing_user_settings_in_rails_5_api/)
- url: https://www.reddit.com/r/rails/comments/j79doh/storing_user_settings_in_rails_5_api/
---
I'm wondering what the best way to store user preferences / settings in a rails 5 api application. Most of my search results use deprecated gems, or are behind a paywall. (looking at you gorails). Any info on the topic would be greatly appreciated, thanks
## [6][link_to generates path like /controller?id=1234, I need /controller/1234](https://www.reddit.com/r/rails/comments/j711dn/link_to_generates_path_like_controllerid1234_i/)
- url: https://www.reddit.com/r/rails/comments/j711dn/link_to_generates_path_like_controllerid1234_i/
---
I am having some trouble rendering my show.html.erb view. I have the following line to create a button on the page:

`&lt;%= link_to('Click Here', books_path(id: book.id), :class =&gt; 'button') %&gt;`

which generated a path like /books?id=1234. This does not render my show view or hit my show action in the controller. However, if I manually type /books/1234 into the address bar it seems to work and render my view.

in routes.rb I have:

`resources :books`

Can anyone shed any light on what I'm doing wrong here?
## [7][Scoped associations in Rails](https://www.reddit.com/r/rails/comments/j6lxl8/scoped_associations_in_rails/)
- url: https://www.reddit.com/r/rails/comments/j6lxl8/scoped_associations_in_rails/
---
Howdy folks! 

I've never thought that ActiveRecord associations could be used as scopes until last week. I couldn't find a step-by-step explanation, so [I did a quick write up about how one can scope associations in Rails](https://remimercier.com/scoped-active-record-associations/). 

Some key take-aways: 
👀  Make your code much more readable (stylish, even)
✂  DRY things up
⚡️  Preload your scopes and remove your N+1 queries

Hope you'll like it!
## [8][RailsAdmin: How to disable edit action?](https://www.reddit.com/r/rails/comments/j6qhyq/railsadmin_how_to_disable_edit_action/)
- url: https://www.reddit.com/r/rails/comments/j6qhyq/railsadmin_how_to_disable_edit_action/
---
Hi there,

I'm working on a rails project with [rails\_admin](https://github.com/sferik/rails_admin) and multiple models. There are several people working on the backend and I want to remove the ability to edit some of the records which have a `imported` boolean set to `true`. This records should just be readable in rails\_admin.

Is it possible to remove the pencil-button in the list view of a record an make this records read only with the condition that the flag is `true`?

The documentation [shows up](https://github.com/sferik/rails_admin/wiki/Base-action#only) some similar behavior to disable action in general but there is no way to bind this configuration on a condition (I guess).

I just found out that you can define a method called `readonly?` on the model but that won't work while I have to update or change attributes while processing this record in my programming logic.

Many thanks!
## [9][active_storage in production](https://www.reddit.com/r/rails/comments/j6omlt/active_storage_in_production/)
- url: https://www.reddit.com/r/rails/comments/j6omlt/active_storage_in_production/
---
background: I am building a Rails API for my react frontend Instagram clone.

Everything was working perfectly in development but now that i pushed the code to my VPS I am having problems with active\_storage. I can upload pictures but when I try to get them the returned url for the image looks like ´[mydomain.com/rails/active\_storage/blobs/somehash/somename.jpg](https://mydomain.com/rails/active_storage/blobs/somehash/somename.jpg)´ -   placing that url in a img src (or trying to open it) returns a 404.

Using Unicorn and Nginx on the server, active\_storage is set to **local file** and saving to a postgresql db. Any ideas?
## [10][Documentation or expertise with AWS Cognito SDK](https://www.reddit.com/r/rails/comments/j6desf/documentation_or_expertise_with_aws_cognito_sdk/)
- url: https://www.reddit.com/r/rails/comments/j6desf/documentation_or_expertise_with_aws_cognito_sdk/
---
I am currently implementing Cognito in a rails 6 app with a view to then using OIDC from Cognito onwards, enabling SSO. Unfortunately, the docs are lacklustre at best. They go into quite a lot of technical detail (which is excellent) but then skip crucial points. For example, `secret_hash` ([used here](https://docs.aws.amazon.com/sdk-for-ruby/v3/api/Aws/CognitoIdentityProvider/Client.html#sign_up-instance_method)) requires some boilerplate code which feels like it should belong inside the SDK. The docs don't tell you this, nor how to generate it. Thankfully, [StackOverflow](https://stackoverflow.com/questions/37438879/unable-to-verify-secret-hash-for-client-in-amazon-cognito-userpools) came to my rescue.

Of significant more concern is that everything I google seems to suggest that people try to get Cognito to work but then give up. I can find plenty of people who are trying to get it to work in Ruby but very few who have successfully done so. A lot of "it's not worth it" or "we looked into it but we couldn't get it to work" type comments.

1. What should the actual flow look like?
1. Should I be using the `admin_x` methods such as `admin_initiate_auth` or the standard `initiate_auth`? What's the difference?
1. If I'm using the flow that requires `SRP_A`, how do I generate it?
1. Almost as a show of hands - how many people have tried to implement it? Were you successful or not? Why?

Thanks.
## [11][Ruby on Rails &amp; Digital Ocean](https://www.reddit.com/r/rails/comments/j63loa/ruby_on_rails_digital_ocean/)
- url: https://www.reddit.com/r/rails/comments/j63loa/ruby_on_rails_digital_ocean/
---
I'm looking to migrate [my Ruby on Rails app](https://www.mugshotbot.com) from Heroku to Digital Ocean. I don't have any dev-ops experience and am very confused on where to start.

Does anyone have any beginner tutorials or links I could work through?
## [12][Active Storage -- Uploads via an API](https://www.reddit.com/r/rails/comments/j62zs0/active_storage_uploads_via_an_api/)
- url: https://www.reddit.com/r/rails/comments/j62zs0/active_storage_uploads_via_an_api/
---
Hi!

*Short version:* I'm planning on writing an editor with an image uploader. The images have to be uploaded seperately from the rest of the form via AJAX, so that the server can return permanent URLs that can then be used in the editor. How would I go about this? Do I need an extra Image model or can I still attach those images to my original Page/Post/whatever model?

*Long version:*

I want to create an image upload feature for my current project. More specifically, it should accept images via an API. The idea is to have 2 forms on a page:

1. the normal contents form
2. a form for images, submitted via JavaScript

After the images have been submitted via JavaScript, I need permanent URLs to those images being sent back so I can use them in the first form. Ugh, I hope you understand what I mean.
