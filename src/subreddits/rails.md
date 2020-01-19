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
## [2][on: :create is not working with validate](https://www.reddit.com/r/rails/comments/eqip01/on_create_is_not_working_with_validate/)
- url: https://www.reddit.com/r/rails/comments/eqip01/on_create_is_not_working_with_validate/
---
I have a code like

    validate: :call_this_method, on: :create 

And when I call the create method using a curl post command this method is not being invoked. If I remove the on: part it works fine. The same happens if I just put on: :update. I'm not sure why this is happening or how to debug this.

The class that includes these both as well.

    include ActiveModel::Validations
    include ActiveModel::Validations::Callbacks

Could someone please help me with this?

Also if there is a way to find out if the validation is from create or update in the callback method would be great. Since it would solve this problem too.
## [3][What are you using for deployment these days?](https://www.reddit.com/r/rails/comments/eq6yld/what_are_you_using_for_deployment_these_days/)
- url: https://www.reddit.com/r/rails/comments/eq6yld/what_are_you_using_for_deployment_these_days/
---
I've been deploying Rails applications with some wrappers around Mina that I wrote a few years ago, but have been looking at more mature deployment workflows recently as those applications scale further and was wondering what everyone else have been doing recently?

Ideally I'd love to be able to able to go from writing code to deploying as quickly as possible so containerizing the app seems like it would slow that down if you have to wait for an image build? As compared to just swapping out code on a VM and restarting Puma?

What has been your experience with rolling deploys vs. blue/green? How about multi-region deploys?
## [4][data is in database but when viewed in browser app only shows data older than 6 months](https://www.reddit.com/r/rails/comments/eq8k6o/data_is_in_database_but_when_viewed_in_browser/)
- url: https://www.reddit.com/r/rails/comments/eq8k6o/data_is_in_database_but_when_viewed_in_browser/
---
Hey guys,

We use redmine and bitbucket to manage our product. In an effort to integrate the two systems, I am following the steps shown [here](https://www.redmine.org/plugins/bitbucket_reference_redmine).

After completing the steps beneath the heading "First, install the plugin in Redmine", I restarted Redmine and discovered an error trace to a file in the plugin directory: NoMethodError because the plugin was looking for config.secret_key_base but it wasn't configured.

No big deal, I'll add it to the environment. Generated one with SecureRandom, fixed the error, Redmine working. Except what I see in Redmine is from 7 months ago. My work issues are old, and each note's timestamp said 7 months ago.

I check the database and I find my current issues' records intact. I was at the point of backing up the database before executing `rake db:schema:load` to see if that would fix the issue (perhaps that alone, or perhaps after reloading the backed up data). Then my boss asked for an update and decided to have me revert redmine to what it was.

So I removed the plugin using the rake task to roll back the migration. I also deleted .bundle/ and rebundled. Despite my best efforts, Redmine in the browser still shows the old data.

I think this is probably related more to the fact that I added the secret key base rather than to the bitbucket plugin itself, but I don't know. I can load current records in MySQL as well as rails console, but Redmine in the browser is still time warped. I'm looking for clues within the Redmine source right now, as well as the source for the plugin.

I was hoping to get some suggestions as far as where to look to fix this issue, or if anyone has ever experienced something similar. Thanks in advance!

TL;DR: what could possibly cause me to be able to view current records in MySQL console and rails console, but when viewing the data in browser, the data is all &gt; 6 months old
## [5][Accessing a rail app without internet? Newbie question](https://www.reddit.com/r/rails/comments/eq5n5d/accessing_a_rail_app_without_internet_newbie/)
- url: https://www.reddit.com/r/rails/comments/eq5n5d/accessing_a_rail_app_without_internet_newbie/
---
Hi *!
So I am super excited since I found my very first customer!🎉
Long story short, he needs a dashboard type of app, where he can control various subjects (emailing, agenda, content of the project...).
He also asked if he could access the dashboard without internet... Just to consult his projets, without the need to edit anything. 
Not sure I use the correct wording but the idea is to have a local replicate of the latest version of his website. 
I will be using rails 5.2. And... I have no idea how to do this simply! 😁
May be a gem could help? 
Thanks a lot in advance! 🙏
## [6][Prevent app from caching in :development](https://www.reddit.com/r/rails/comments/eq5g0p/prevent_app_from_caching_in_development/)
- url: https://www.reddit.com/r/rails/comments/eq5g0p/prevent_app_from_caching_in_development/
---
I recently upgraded from Rails 5.2.3 to 6.0.2.1, everything went supper smoothly, tests are passing, etc.

However for some reason the app seems to now be cached in `:development` and changes to files (i.e. `.erb.html`) don't get picked up unless I restart the Rails server.
This was obviously not the case before the upgrade.

I'm only able to see changes to files in the `/assets` folder without restarting the server.

For the record my `development.rb` is set to `config.cache_classes = false`
In general, virtually all of the config files were left intact by the upgrade task (aside from some new additions in Rails 6).

As I was trying to troubleshoot this I saw similar cases where people were "blaming" `Spring` for similar behavior. `config/spring.rb` was one of the few files that actually change during the upgrade process, from:
```
%w[
  .ruby-version
  .rbenv-vars
  tmp/restart.txt
  tmp/caching-dev.txt
].each { |path| Spring.watch(path) }
```
to
```
Spring.watch(
  ".ruby-version",
  ".rbenv-vars",
  "tmp/restart.txt",
  "tmp/caching-dev.txt"
)
```

I honestly don't know how Spring really works or how it gets executed (I'm assuming `rails s` initializes it?). I'm on the latest `v2.1.0`. Could this be the issue?

What do you suggest I should look for/change to prevent this?
Thanks!
## [7][Signing cookies externally with the secret_key_base](https://www.reddit.com/r/rails/comments/eq7g6b/signing_cookies_externally_with_the_secret_key/)
- url: https://www.reddit.com/r/rails/comments/eq7g6b/signing_cookies_externally_with_the_secret_key/
---
I'm performing a penetration test, and managed to get a hold of the secret\_key\_base. I know I can use this to perform session tampering, and perhaps even RCE, but I just can't figure out how to pull it off.

I found [this script](https://gist.githubusercontent.com/cyberheartmi9/7fe85b61621f4126462d2125c4b19dfe/raw/cc39783a5a5bacda2762a7db6e163fffe91e8db9/Attacking%2520Ruby%2520on%2520Rails%2520Applications)  (see section 2.1), along with an explanation of the concept which will  do the same thing for Rails &lt;4.1.0, using Marshal - I just need to do  the same using JSON encoding, but cannot for the life of me figure it  out. I could very much use a hand from those of you who are more experienced in Rails!
## [8][How to build a live, face-to-face video chat app in Ruby on Rails 6.0.2.1 ​](https://www.reddit.com/r/rails/comments/eprjo9/how_to_build_a_live_facetoface_video_chat_app_in/)
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
## [9][Tips for writing readable system tests in Rails](https://www.reddit.com/r/rails/comments/epzag1/tips_for_writing_readable_system_tests_in_rails/)
- url: https://www.reddit.com/r/rails/comments/epzag1/tips_for_writing_readable_system_tests_in_rails/
---
Want to make system tests easy to main tain? We have selected some best practice tips to help achieve this. [https://jtway.co/best-tips-for-writing-integration-tests-in-rails-d1f56081f249](https://jtway.co/best-tips-for-writing-integration-tests-in-rails-d1f56081f249)
## [10][How to you find help when you have problem ?](https://www.reddit.com/r/rails/comments/epy25w/how_to_you_find_help_when_you_have_problem/)
- url: https://www.reddit.com/r/rails/comments/epy25w/how_to_you_find_help_when_you_have_problem/
---
Hello everyone,

I'm struggling since three days with a problem on my Rails app.

I first tried to solve it by myself so and finally got a bit of result. I know exactly where is the cause but I don't understand why.

I asked to a friend, and a slack, no result. I asked on [stackoverflow](https://stackoverflow.com/questions/59730179/devise-api-sign-in-doesnt-work-with-postman) I've got no answer, I put a bounty on my problem still no result.

So here is my question. What do you do when you are struggling on something and you don't find any help for it ?
## [11][I'm learning Docker and I'm trying to understand the logic behind its use](https://www.reddit.com/r/rails/comments/epopbc/im_learning_docker_and_im_trying_to_understand/)
- url: https://www.reddit.com/r/rails/comments/epopbc/im_learning_docker_and_im_trying_to_understand/
---
I have experience of Rails for quite some time and I feel comfortable with it, but since a couple of days I've starting to learn and study Docker.

As I understand it, with the Dockerfile you want to install all the necessary libraries and programs so that rails can be run, but if it uses services like `postgres`, `redis`, `mysql` we should define them at the `docker-compose.yml` file? Because  we could (in the same container) install postgres and RUN the commands needed to make it available there, can't we?, but I always see that the minimum commands for installing rails are executed with RUN.

I think I'm a little confused about what the role of Dockerfile and Docker Compose is, what's the rule to know what sould be in docker-compose.yml and what should be in the Dockerfile.
