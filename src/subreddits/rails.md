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
## [3][Scoped associations in Rails](https://www.reddit.com/r/rails/comments/j6lxl8/scoped_associations_in_rails/)
- url: https://www.reddit.com/r/rails/comments/j6lxl8/scoped_associations_in_rails/
---
Howdy folks! 

I've never thought that ActiveRecord associations could be used as scopes until last week. I couldn't find a step-by-step explanation, so [I did a quick write up about how one can scope associations in Rails](https://remimercier.com/scoped-active-record-associations/). 

Some key take-aways: 
👀  Make your code much more readable (stylish, even)
✂  DRY things up
⚡️  Preload your scopes and remove your N+1 queries

Hope you'll like it!
## [4][RailsAdmin: How to disable edit action?](https://www.reddit.com/r/rails/comments/j6qhyq/railsadmin_how_to_disable_edit_action/)
- url: https://www.reddit.com/r/rails/comments/j6qhyq/railsadmin_how_to_disable_edit_action/
---
Hi there,

I'm working on a rails project with [rails\_admin](https://github.com/sferik/rails_admin) and multiple models. There are several people working on the backend and I want to remove the ability to edit some of the records which have a `imported` boolean set to `true`. This records should just be readable in rails\_admin.

Is it possible to remove the pencil-button in the list view of a record an make this records read only with the condition that the flag is `true`?

The documentation [shows up](https://github.com/sferik/rails_admin/wiki/Base-action#only) some similar behavior to disable action in general but there is no way to bind this configuration on a condition (I guess).

I just found out that you can define a method called `readonly?` on the model but that won't work while I have to update or change attributes while processing this record in my programming logic.

Many thanks!
## [5][active_storage in production](https://www.reddit.com/r/rails/comments/j6omlt/active_storage_in_production/)
- url: https://www.reddit.com/r/rails/comments/j6omlt/active_storage_in_production/
---
background: I am building a Rails API for my react frontend Instagram clone.

Everything was working perfectly in development but now that i pushed the code to my VPS I am having problems with active\_storage. The returned url for the image attachments looks like ´[mydomain.com/rails/active\_storage/blobs/somehash/somename.jpg](https://mydomain.com/rails/active_storage/blobs/somehash/somename.jpg)´ but placing that url in a img src (or trying to open it) gives a 404.   


Using Unicorn and Nginx on the server, active\_storage is set to local file and saving to a postgresql db. Any ideas?
## [6][Documentation or expertise with AWS Cognito SDK](https://www.reddit.com/r/rails/comments/j6desf/documentation_or_expertise_with_aws_cognito_sdk/)
- url: https://www.reddit.com/r/rails/comments/j6desf/documentation_or_expertise_with_aws_cognito_sdk/
---
I am currently implementing Cognito in a rails 6 app with a view to then using OIDC from Cognito onwards, enabling SSO. Unfortunately, the docs are lacklustre at best. They go into quite a lot of technical detail (which is excellent) but then skip crucial points. For example, `secret_hash` ([used here](https://docs.aws.amazon.com/sdk-for-ruby/v3/api/Aws/CognitoIdentityProvider/Client.html#sign_up-instance_method)) requires some boilerplate code which feels like it should belong inside the SDK. The docs don't tell you this, nor how to generate it. Thankfully, [StackOverflow](https://stackoverflow.com/questions/37438879/unable-to-verify-secret-hash-for-client-in-amazon-cognito-userpools) came to my rescue.

Of significant more concern is that everything I google seems to suggest that people try to get Cognito to work but then give up. I can find plenty of people who are trying to get it to work in Ruby but very few who have successfully done so. A lot of "it's not worth it" or "we looked into it but we couldn't get it to work" type comments.

1. What should the actual flow look like?
1. Should I be using the `admin_x` methods such as `admin_initiate_auth` or the standard `initiate_auth`? What's the difference?
1. If I'm using the flow that requires `SRP_A`, how do I generate it?
1. Almost as a show of hands - how many people have tried to implement it? Were you successful or not? Why?

Thanks.
## [7][Ruby on Rails &amp; Digital Ocean](https://www.reddit.com/r/rails/comments/j63loa/ruby_on_rails_digital_ocean/)
- url: https://www.reddit.com/r/rails/comments/j63loa/ruby_on_rails_digital_ocean/
---
I'm looking to migrate [my Ruby on Rails app](https://www.mugshotbot.com) from Heroku to Digital Ocean. I don't have any dev-ops experience and am very confused on where to start.

Does anyone have any beginner tutorials or links I could work through?
## [8][Active Storage -- Uploads via an API](https://www.reddit.com/r/rails/comments/j62zs0/active_storage_uploads_via_an_api/)
- url: https://www.reddit.com/r/rails/comments/j62zs0/active_storage_uploads_via_an_api/
---
Hi!

*Short version:* I'm planning on writing an editor with an image uploader. The images have to be uploaded seperately from the rest of the form via AJAX, so that the server can return permanent URLs that can then be used in the editor. How would I go about this? Do I need an extra Image model or can I still attach those images to my original Page/Post/whatever model?

*Long version:*

I want to create an image upload feature for my current project. More specifically, it should accept images via an API. The idea is to have 2 forms on a page:

1. the normal contents form
2. a form for images, submitted via JavaScript

After the images have been submitted via JavaScript, I need permanent URLs to those images being sent back so I can use them in the first form. Ugh, I hope you understand what I mean.
## [9][How to Deploy Ruby on Rails with MongoDB to Heroku?](https://www.reddit.com/r/rails/comments/j60455/how_to_deploy_ruby_on_rails_with_mongodb_to_heroku/)
- url: https://www.reddit.com/r/rails/comments/j60455/how_to_deploy_ruby_on_rails_with_mongodb_to_heroku/
---
Hi! I created a project by following this documentation of MongoDB [https://docs.mongodb.com/mongoid/current/tutorials/getting-started-rails/](https://docs.mongodb.com/mongoid/current/tutorials/getting-started-rails/)

&amp;#x200B;

I finished the project, but I don't know how to set it up to Heroku?
## [10][Integrating Bootstrap theme in Rails 6](https://www.reddit.com/r/rails/comments/j5v9xt/integrating_bootstrap_theme_in_rails_6/)
- url: https://www.reddit.com/r/rails/comments/j5v9xt/integrating_bootstrap_theme_in_rails_6/
---
So, to save time, I purchased a bootstrap theme online for my project. I am having difficulty integrating it . Confused by webpacker /asset pipeline file system. 
Could anyone walk me through the process or point me to a tutorial on how to properly integrate my bootstrap theme into rails 6 ?
## [11][How to implement CKEditor in Rails API](https://www.reddit.com/r/rails/comments/j61uay/how_to_implement_ckeditor_in_rails_api/)
- url: https://www.reddit.com/r/rails/comments/j61uay/how_to_implement_ckeditor_in_rails_api/
---
Hi! I hope everyone is fine,  
I have a task in which I have to implement CKEditor in rails, My work is from the backend side so I have to implement this on my API , After searching many tutorials I find each for the monolithic structure but just not for API,    
Maybe it is easy just for API but I don't know how to do it &amp; how it works with API?    
Can anyone please guide me?  or show me the path to implement this?
## [12][Just a big thank you!](https://www.reddit.com/r/rails/comments/j5hz32/just_a_big_thank_you/)
- url: https://www.reddit.com/r/rails/comments/j5hz32/just_a_big_thank_you/
---
tl;dr: lost interest in web dev, learnt Rails, have just completed rewriting my old homepage; thanks to the Rails community for helping me get through this.

---

Hi! Over the past month, I've completely rewritten my homepage from scratch, in Ruby using Rails. This was my very first real Rails project and I'm so proud that I can finally say that it's (nearly) completed.

I come from a background of many languages, but mostly php. Don't get me wrong, php is an awesome language and if you're interested in a typed scripting language with a Java-like OOP system, then php is probably your best bet.

I've been doing web development in my freetime for nearly 4 years now and I really lost interest lately. The only app that I was still maintaining was my homepage. It had a pretty old php code base from when I was a complete newbie to software development.

From side projects which I had long since discontinued, I gained much, much more knowledge in software development and became what I would consider to be quite a good developer. My homepage was my only project left that was not a library or small tool, and considering its age, I thought of rewriting it.

Long story short, I couldn't bring up the motivation. Started it a few times using different frameworks, but always discontinued the project after a few days. I think I just knew my stack too well and got bored of it. In the end, it was just repeating the same tasks I did a thousand times already, once again. I think you could say that I was in some kind of coding depression.

I came to the conclusion that I needed to learn a new framework or maybe even a whole new language. I tried PHP/Slim, PHP/Symfony (granted, that really was not new to me...), Python/Django and even a few more.

After looking for more frameworks to try, I found Rails. I already knew Rails (not really), because I dabbled with it around 2 years ago. Back then, I dropped it because I neither understood Rails nor Ruby. Having heard great things about Ruby and Rails and having seen the Ruby community as really friendly, calm devs whose projects "just work", I really got interested in it again.

I started learning Ruby using [the Programming Ruby book](http://ruby-doc.com/docs/ProgrammingRuby/) and Rails using several YouTube tutorials and blogs. I felt as if Rails is the right thing for my homepage, so I started working on it (based off of a 45min YouTube video, even though the code base has completely changed since then).

After having understood Rails' basics (where does it get the views/layouts from, how does webpacker work, how do all the helpers work, how do I validate input, how do I query a database, ...), it became pretty easy and I could work on it all by myself, no tutorial needed anymore.

I still ran into quite some problems (looking through my post history, I made 9 posts on here over the past month), which I consulted you guys for. All my questions got answered! And this is just a big thank you to all of you! This community gave me a really warm welcome and I feel right at home here. I joined this sub when I started my Rails journey and in the future I'd like to help others new to Rails too. Rails has an incredibly awesome community.

Also, the people from Stack Overflow, numerous blogs around the web, YouTubers making tutorials and developers writing libraries, deserve all an equally big thank you. To me, it feels like those resources are often not appreciated as much as they should be.

My site is now live and I'm going to Open Source it very soon. I'd love to post a link, but I don't want my Reddit account to be associated with my homepage/IRL personality.

So, again, thank you to all of you! ;)
