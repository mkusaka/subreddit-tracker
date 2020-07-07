# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/
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
## [2][Up to date resource for starting to learn rails](https://www.reddit.com/r/rails/comments/hmod8i/up_to_date_resource_for_starting_to_learn_rails/)
- url: https://www.reddit.com/r/rails/comments/hmod8i/up_to_date_resource_for_starting_to_learn_rails/
---
Hey guys. I'm looking to get into rails development and I'm wondering if there are any recommended up to date resources. The JHU Coursera course seemed tempting but it is from 2015 and uses rails 4 which I feel like is a bad place to start. There's also the learn enough [railstutorial.org](https://railstutorial.org) which seems tempting, but I'd like to look around for a free resource first. 

Any recommendations and advice are appreciated.

Cheers.
## [3][Rails is about giving the power of a 10 person development team to one person through simplicity](https://www.reddit.com/r/rails/comments/hm7mex/rails_is_about_giving_the_power_of_a_10_person/)
- url: https://www.reddit.com/r/rails/comments/hm7mex/rails_is_about_giving_the_power_of_a_10_person/
---
Having built SPA using front end frameworks before, and now building the same web app with Rails I was thinking about the following:

Rails is a great tool for solo / small team developers because it gives you the leverage to deal with the same set of problems as a larger team by providing a simplistic mental model through which to reason about the software application.

For example, let's discuss the often heated SPA / JSON Api versus Server Side Application debate. I think this debate is not actually a technology problem - but a personnel problem. For example, suppose you  have an application that is highly interactive. Giving a dedicated front end team (3-5 of 10 people) a front end framework makes complete sense because it allows those individuals to organize and handle the complexity of five parallel change sets at any one time. That is to say, the front end framework provides the opportunity to create a "\*communicable\* model" for this squad to maintain in their collective understanding \*and to reason about / collaborate on\*.  Without a \*communicable\* model (e.g. spaghetti javascript / css), the difficulty of proposing new change sets and evolving software in a multiplayer environment grows in a non-linear way. This is a really misrepresented risk in technology choices I think.

Frameworks linearize the learning curve for the application's mechanics, allowing \*new collaborators\* to ramp on and quickly evolve their end of the software.

Flipping the script now, if you are just one or two developers, it might not make sense to use a full blown front end framework. The reason is that the primary issue you're dealing with is not managing a \*communicable\* model between many contributors to handle parallel change sets, rather it is a question of "how much you can get done". Rails provides the \*leverage\* here by abstracting certain application decisions / choices and doing the "heavy lifting" allowing you, a solo developer, to hold the entire application model in \*your head\*. That is to say, the application model might not be a \*communicable\* model (e.g. Sever Generated JavaScript responses is probably not as easily grasped as parsing a JSON request) but it is simple enough to hold in one person's mind and still evolve the application with a sufficient speed.

There's some nuances here and I'm probably wrong in some regards, but an interesting thought (I think) none the less

&amp;#x200B;

edit; spelling
## [4][How to create a form to embed into other websites?](https://www.reddit.com/r/rails/comments/hmrp3q/how_to_create_a_form_to_embed_into_other_websites/)
- url: https://www.reddit.com/r/rails/comments/hmrp3q/how_to_create_a_form_to_embed_into_other_websites/
---
I want to create a form that lives in my app but I can then embed it into any other website that I give a piece of code to place into their website.

&amp;#x200B;

What is the best way to go about this with Rails?

&amp;#x200B;

I assume I should be using an iframe but wondering any obstacles I may need to watch out for and any opinions on how I should go about doing this?
## [5][Should I upgrade my Rails 4 API project to Rails 5 or Rails 6?](https://www.reddit.com/r/rails/comments/hmbm7k/should_i_upgrade_my_rails_4_api_project_to_rails/)
- url: https://www.reddit.com/r/rails/comments/hmbm7k/should_i_upgrade_my_rails_4_api_project_to_rails/
---
This API is not that big and it has about 87% test coverage. I was thinking it's better to upgrade to latest major to avoid a new upgrade in the next 2 years.

Also, I'd like to use Rails 6 parallel tests feature.

I'm more concerned about Active Model Serializers, I am currently using version 0.9.7 which seems outdated and there are probably breaking changes to next major 0.10

Also Paperclip is dead in Rais 6, I would need to migrate to Active Storage.

What do you guys think?
## [6][I'm trying to import a theme nav-bar with but it's has being difficult.](https://www.reddit.com/r/rails/comments/hm9g59/im_trying_to_import_a_theme_navbar_with_but_its/)
- url: https://www.reddit.com/r/rails/comments/hm9g59/im_trying_to_import_a_theme_navbar_with_but_its/
---
Hello, I'm trying to import a navbar from a theme with stim but I'm having difficulties such as the navbar getting too "thick" or the styles of the navbar not loading properly

Here are images

[https://postimg.cc/RqL0FmS7](https://postimg.cc/RqL0FmS7) here it's how supposed to be.

Here is how my RoR implementation result: [https://postimg.cc/PpNMqyZZ](https://postimg.cc/PpNMqyZZ)

here is the slim file: [https://github.com/LeoFragozo/notika\_slim\_test/blob/master/app/views/temas/index.html.slim](https://github.com/LeoFragozo/notika_slim_test/blob/master/app/views/temas/index.html.slim)

News: apparently it's not loading the "::before" selector how I can put these in rails? also the inside div is in the wrong size.

[https://postimg.cc/xJc3Mpz0](https://postimg.cc/xJc3Mpz0) here is the working version

[https://postimg.cc/5YTq35B8](https://postimg.cc/5YTq35B8) here is my version  

here is the css file: [https://github.com/LeoFragozo/notika\_slim\_test/blob/master/app/assets/stylesheets/temas.scss](https://github.com/LeoFragozo/notika_slim_test/blob/master/app/assets/stylesheets/temas.scss)
## [7][How make changes on production database](https://www.reddit.com/r/rails/comments/hm9h8o/how_make_changes_on_production_database/)
- url: https://www.reddit.com/r/rails/comments/hm9h8o/how_make_changes_on_production_database/
---
Can someone give me some guide here on how to change production database with good manners. I don't have experience at all with this.

This is the first time I will be maintaining a production server and I was thinking how is the right way to make changes on production DB. I have only been on development until now, and on deploying on staging I always reset the database then populate it with seeds files. As that is impossible on production, how should I proceed on this task. 

One thing to note about my project is that my migrations file are a mess. I don't usually create migration using the "up" and "down" (is this still used by rails?) and a make a lot of changes, so I'm not confident to used them for recreating the database. Futures migrations can be more "organized", but the past ones will probably raise exceptions. 

From my understating I should do something like:

* First deploy on production, database will be created from schema.
* After first deploy, create a BK if, don't already exist, before make changes
* New changes must be working on development before going to prod.
* For changing the DB, run the last migrations (not sure how to proceed here - maybe should delete the old migrations before?) 
* If works right, than good. Else, reverse migrations.

Any advice?
## [8][Containerized Development + CI -- is there need for instructional posts/videos?](https://www.reddit.com/r/rails/comments/hmbzfq/containerized_development_ci_is_there_need_for/)
- url: https://www.reddit.com/r/rails/comments/hmbzfq/containerized_development_ci_is_there_need_for/
---
Hi all .. 

I've been developing on a rails container based platform (Che Eclipse) and have had a CI/CD setup running using gitlab.

This is for personal development at home (not a business) -- i've been a developer for 20+ years and the last 5 on this setup have been the most satisfying from a toolbox pov (once setup, its easy to tweak without breaking something else).

I was hoping to gauge if this sort of info would be useful to some (if you are already an expert, i appreciate feedback too), there is already lots of info out there, this sort of setup has been around for many years, so before creating some material i just want to see if it would be useful to some. -- let me know if you would like to see some guides about how to set this up. -- I've never made instructional videos / posts, but i'd like to try.

I was thinking of making a multi-part series -- 

\- setting up eclipse/che, 

\- setting up gitlab, 

\- git basics, 

\- che/basics, 

\- setting up rails 6 on che (setting up your development container)

\- simple app -&gt; CI/CD for container based rails apps.

\- docker container registries (running your ow).

\- deployment on docker

thats already tons of info there -- my target would be basic knowledge, starting from almost zero -- so very basic instruction and basic examples.
## [9][Spree Commerce Help](https://www.reddit.com/r/rails/comments/hm6hfg/spree_commerce_help/)
- url: https://www.reddit.com/r/rails/comments/hm6hfg/spree_commerce_help/
---
Hi There,

Trying out Spree Commerce and having trouble rendering editted templates in my app/views/spree directory. I am able to override templates in Solidus by recreating them in app/views/spree but it doesn't seem to work in Spree.

Anyone experienced this issue?

Rails 6.0.3 &amp; Spree 4.1

&amp;#x200B;

Thank you in advance!
## [10][I’m starting a new project and having a hard time determining whether to use react, vuejs, or stimulus. It will start out as a simple photo sharing app but will eventually need to pull in data after image processing and provide visualization. Any thoughts on which route to go?](https://www.reddit.com/r/rails/comments/hmajmg/im_starting_a_new_project_and_having_a_hard_time/)
- url: https://www.reddit.com/r/rails/comments/hmajmg/im_starting_a_new_project_and_having_a_hard_time/
---

## [11][How to build guest cart/checkout?](https://www.reddit.com/r/rails/comments/hm6xjd/how_to_build_guest_cartcheckout/)
- url: https://www.reddit.com/r/rails/comments/hm6xjd/how_to_build_guest_cartcheckout/
---
Currently I have an ecommerce app that allows users to add items to a cart and checkout if they have an existing user account (the cart object is associated with user object and product object). I want to create the ability for guest checkout, but not sure how this would work from an architecture perspective if a user doesn't have an account. Any thoughts? Thanks in advance.
