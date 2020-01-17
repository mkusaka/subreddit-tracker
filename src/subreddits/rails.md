# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/
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
## [2][How to build a live, face-to-face video chat app in Ruby on Rails 6.0.2.1 ​](https://www.reddit.com/r/rails/comments/eprjo9/how_to_build_a_live_facetoface_video_chat_app_in/)
- url: https://www.reddit.com/r/rails/comments/eprjo9/how_to_build_a_live_facetoface_video_chat_app_in/
---
&amp;#x200B;

[home page](https://preview.redd.it/024hrik638b41.png?width=2560&amp;format=png&amp;auto=webp&amp;s=87a1067b498863a674d65a4293a68abf3f90c830)

#### DISCLAIMER: This tutorial is not for beginners.

Want to know how to build apps like skype, zoom etc... using Ruby on rails 6? I got you!!

Building a video chat feature on a ruby on rails application is not an easy thing to do. That’s one of the reasons why I decided to do it, another reason being that I didn’t want to build another blog, todo-app etc… I wanted something challenging. 

You can check the app on heroku at [https://dev-video.herokuapp.com/](https://dev-video.herokuapp.com/).

For a demo, you can either do it yourself (by logging into two separate accounts on different browsers), or check this small demo on youtube at [https://www.youtube.com/watch?v=WXouJ8w7rQg](https://www.youtube.com/watch?v=WXouJ8w7rQg).

Here is a link to the first 7 chapters of the tutorial: [https://www.dropbox.com/s/2aqdsvm721tr01s/Ebook2020.pdf?dl=0](https://www.dropbox.com/s/2aqdsvm721tr01s/Ebook2020.pdf?dl=0)

You can get the full ebook on Gumroad at [https://gumroad.com/l/lNEhE](https://gumroad.com/l/lNEhE) for the price of a Starbucks coffee.

If you want to follow along with this tutorial, you are going to need a few things:

* Ruby 2.7
* Rails 6.0.2
* AWS S3 bucket
* AWS credentials (access\_key\_id and secret\_access\_key)

HAPPY CODING!!!
## [3][Tips for writing readable system tests in Rails](https://www.reddit.com/r/rails/comments/epzag1/tips_for_writing_readable_system_tests_in_rails/)
- url: https://www.reddit.com/r/rails/comments/epzag1/tips_for_writing_readable_system_tests_in_rails/
---
Want to make system tests easy to main tain? We have selected some best practice tips to help achieve this. [https://jtway.co/best-tips-for-writing-integration-tests-in-rails-d1f56081f249](https://jtway.co/best-tips-for-writing-integration-tests-in-rails-d1f56081f249)
## [4][How to you find help when you have problem ?](https://www.reddit.com/r/rails/comments/epy25w/how_to_you_find_help_when_you_have_problem/)
- url: https://www.reddit.com/r/rails/comments/epy25w/how_to_you_find_help_when_you_have_problem/
---
Hello everyone,

I'm struggling since three days with a problem on my Rails app.

I first tried to solve it by myself so and finally got a bit of result. I know exactly where is the cause but I don't understand why.

I asked to a friend, and a slack, no result. I asked on [stackoverflow](https://stackoverflow.com/questions/59730179/devise-api-sign-in-doesnt-work-with-postman) I've got no answer, I put a bounty on my problem still no result.

So here is my question. What do you do when you are struggling on something and you don't find any help for it ?
## [5][I'm learning Docker and I'm trying to understand the logic behind its use](https://www.reddit.com/r/rails/comments/epopbc/im_learning_docker_and_im_trying_to_understand/)
- url: https://www.reddit.com/r/rails/comments/epopbc/im_learning_docker_and_im_trying_to_understand/
---
I have experience of Rails for quite some time and I feel comfortable with it, but since a couple of days I've starting to learn and study Docker.

As I understand it, with the Dockerfile you want to install all the necessary libraries and programs so that rails can be run, but if it uses services like `postgres`, `redis`, `mysql` we should define them at the `docker-compose.yml` file? Because  we could (in the same container) install postgres and RUN the commands needed to make it available there, can't we?, but I always see that the minimum commands for installing rails are executed with RUN.

I think I'm a little confused about what the role of Dockerfile and Docker Compose is, what's the rule to know what sould be in docker-compose.yml and what should be in the Dockerfile.
## [6][Remote work subreddit](https://www.reddit.com/r/rails/comments/epsib5/remote_work_subreddit/)
- url: https://www.reddit.com/r/rails/comments/epsib5/remote_work_subreddit/
---
[https://www.reddit.com/r/remotetechjobs/](https://www.reddit.com/r/remotetechjobs/)
## [7][Anybody use a page speed measurement tool that can be called from bash?](https://www.reddit.com/r/rails/comments/eplqt7/anybody_use_a_page_speed_measurement_tool_that/)
- url: https://www.reddit.com/r/rails/comments/eplqt7/anybody_use_a_page_speed_measurement_tool_that/
---
  
how can i measure the "speed" of a given url?  
What I have now is \`time curl [http://example.com/terms-of-service.html\`](http://example.com/terms-of-service.html`)  
but looking for a more sophisticated (but cli friendly)  way -   
for example:  
  \- show me time taken from request initiation to LVC (last visual change)  
 \- breakdown of slowness - in the client side vs server side  
 \- breakdown of whether time was spent on ActiveRecord or ActionViewRendering
## [8][Validation in rails 4.2.1](https://www.reddit.com/r/rails/comments/epmbpf/validation_in_rails_421/)
- url: https://www.reddit.com/r/rails/comments/epmbpf/validation_in_rails_421/
---
I'm writing validations for an API.  I have two attributes old\_attr and new\_attr.

The validation rules are,

1. either of them should be present during create
2. during update both of them can be present or not but cannot be an empty array if present. 

I'm finding it hard to write this logic 

    validates :old_attr, required: { old_attr: true }, if: -&gt; { new_attr.blank? }
validates :new_attr, required: { new_attr: true }, if: -&gt; { old_attr.blank? }

This works well for the create part but in the update part one of them is still mandatory is also applied. 

How do I include the validations in point 2 as well?
## [9][How to uninstall the apartment gem?](https://www.reddit.com/r/rails/comments/eplmp7/how_to_uninstall_the_apartment_gem/)
- url: https://www.reddit.com/r/rails/comments/eplmp7/how_to_uninstall_the_apartment_gem/
---
I have an app where I installed the apartment gem for multi tenancy, but I no longer need it. 

I deleted the apartment.rb initalizer and cleaned up the application.rb file. And I removed the apartment gem and ran bundle install.

But when I try to run rake migrate I get the following error

    LoadError: cannot load such file -- apartment/elevators/subdomain

Any one know how I can remove the apartment gem completely? Thanks in advance.
## [10][How y'all doing performance wise with Rails 6 (compared to 5.2)?](https://www.reddit.com/r/rails/comments/ep8odb/how_yall_doing_performance_wise_with_rails_6/)
- url: https://www.reddit.com/r/rails/comments/ep8odb/how_yall_doing_performance_wise_with_rails_6/
---
We recently upgraded our monolith to Rails 6 and unfortunately there is a clear trend: https://i.imgur.com/6Ca3pT4.png  
*edit/*  
x-axis = time (this is like 4 days)  
y-axis = response time (mean, 90th, 95th)  
*/edit*

We did a few things at once (updated several gems to new major versions, finally using cache versioning) but no real change in *our* code. We switched back to non-versioned caching for a while but that doesn't change anything in the timings.

I don't want to exclude the possibility that it's not Rails but any of the gems we are using but before I do some more testing I thought I'd ask if any of you guys have some insights. Did your apps perform better or worse after Rails 5&gt;6 upgrade? Did you encounter anything particular while migrating?

Thanks in advance :)
## [11][Programmatically generate new email addresses?](https://www.reddit.com/r/rails/comments/epkzn9/programmatically_generate_new_email_addresses/)
- url: https://www.reddit.com/r/rails/comments/epkzn9/programmatically_generate_new_email_addresses/
---
Anyone know of a service that lets you programmatically generate email addresses to (essentially) act as forwarding addresses?  


I want to be able to programmatically generate email addresses to act as "middlemen" between external services and a private, "master" email account I own.  I know next to nothing about how email "really" works beyond hooking up to a transactional email service like SendGrid.  


I'd like to do this in Rails just because I prefer Ruby/Rails, but if another language/framework makes this easier to do, I'll definitely consider it.
