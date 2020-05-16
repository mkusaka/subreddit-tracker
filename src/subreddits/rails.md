# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/
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
## [2][friendly_id and slug_candidates work only for the first 2 times...](https://www.reddit.com/r/rails/comments/gksig2/friendly_id_and_slug_candidates_work_only_for_the/)
- url: https://www.reddit.com/r/rails/comments/gksig2/friendly_id_and_slug_candidates_work_only_for_the/
---
 Hi guys, i'm using the friendly\_id gem (version 5).

I tried to edit it in this way to have an url as title-YY, title-YY-2, title-YY-3, time-YY-4 etc

     extend FriendlyId
     friendly_id :slug_candidates, use: [:slugged, :finders]
    
       def slug_candidates
         [:title, :title_and_sequence]
       end
    
        def title_and_sequence
         slug = title.to_param
         sequence = Book.where("slung = '#{slung}-%").count + 2
         "#{slug}--#{sequence}"
       end
    

It works with title-YY and title-YY-2... **BUT** **when I create the third page** with "Title YY", it creates  


    mywebsite.com/pages/title-yy-asda-iiqej-asddas

How to fix it to make it "forever"?
## [3][Best way to update two records based on each others values.](https://www.reddit.com/r/rails/comments/gkto3w/best_way_to_update_two_records_based_on_each/)
- url: https://www.reddit.com/r/rails/comments/gkto3w/best_way_to_update_two_records_based_on_each/
---
Hi, I have two 'Creature' records and I would like to perform a calculation based on both records' stats, and then update both records.

This is my current Creatures controller:

    def hunt
        predator = Creature.find(params[:id])
        prey = Creature.find(predator.prey)
        @predator = predator.hunt_predator
        @prey = prey.hunt_prey
    end

And here is a snippet from my Creature model:

    def hunt_prey
        prey = self.clone
        prey.population = prey.population*0.8
        prey.save!
    end

My challenge is that the model changes the records irrespective of the other creatures stats. Is there a way of calling the `predator.hunt_predator` model that takes the other record as input so I can modify both?

Alternatively, am I approaching this in a non-Rails way, and is there a much better practice way of doing this?

Looking online I found a lot of suggestions to use `Creature.update_all`, but couldn't find anything for running the calculations using both records as input. I've considered running the calculations in the controller and then using `Creature.update_all` but I know having a bloated controller is bad practice.

Thanks so much for the help, I'm very new to Rails and been stumped on this for weeks.
## [4][Common issues deploy to production and how to fix that ?](https://www.reddit.com/r/rails/comments/gkrf01/common_issues_deploy_to_production_and_how_to_fix/)
- url: https://www.reddit.com/r/rails/comments/gkrf01/common_issues_deploy_to_production_and_how_to_fix/
---
Hey guys, I wanna deploy my first rails app into production this week or next week in GCP, before going do this I wanna know the common issues when you deploy to production and how to fix that?


In my Rails app, I am using Active Storage with GCP bucket, Whenever Cron for run rake task, Rbenv, Delayed Job, PostgreSql, I saw some articles how to deploy to Ubuntu server , and there are some ways to deploy 
From using Mina, Passenger , Unicorn, and Capistrano. Which you use for this ? And if we use this gem to help us to deploy , would it make issue on the things I use in my Rails app like active job , active storage , or else ? And will make some more configuration for it ??
Thank you for attention
## [5][How to implement SEO in Rails?](https://www.reddit.com/r/rails/comments/gkt6xt/how_to_implement_seo_in_rails/)
- url: https://www.reddit.com/r/rails/comments/gkt6xt/how_to_implement_seo_in_rails/
---
How do you make rails SEO friendly.? Have you implementation it and works very well in rails ?
## [6][Setting up an Automatic Book reader with Devise + Rails - 2](https://www.reddit.com/r/rails/comments/gkssv6/setting_up_an_automatic_book_reader_with_devise/)
- url: https://www.reddit.com/r/rails/comments/gkssv6/setting_up_an_automatic_book_reader_with_devise/
---
Hey guys,  I hate how much coding channels focus on a basic 'ToDo list / blog' so i wanted wanted to share my progress in building an automatic book reader (not that good in making vids but meh). 

Here is my second video where we explore the CRUD section of the application.

 [https://www.youtube.com/watch?v=dxFTUA1leRc&amp;feature=youtu.be](https://www.youtube.com/watch?v=dxFTUA1leRc&amp;feature=youtu.be) 

Videos will get better as time goes on ;)
## [7][how should i update my config/webpack/environment.js for bootstrap 4 ?](https://www.reddit.com/r/rails/comments/gkgsl9/how_should_i_update_my_configwebpackenvironmentjs/)
- url: https://www.reddit.com/r/rails/comments/gkgsl9/how_should_i_update_my_configwebpackenvironmentjs/
---
i've looked at three guides, and ended up confused even more:  


  


https://preview.redd.it/v3gbl91ymzy41.png?width=2400&amp;format=png&amp;auto=webp&amp;s=e61348825d05d95585dd5230def6679ec0a295a8
## [8][Build and test docker image in gitlab CI](https://www.reddit.com/r/rails/comments/gkacsb/build_and_test_docker_image_in_gitlab_ci/)
- url: https://www.reddit.com/r/rails/comments/gkacsb/build_and_test_docker_image_in_gitlab_ci/
---
Hi,

i'm trying to configure gitlab ci in such way, that after push, gitlab runner builds an docker image. On next step this image is pushed to gitlab container registry. On last step I run on this image test suite.

My .gitlab-ci.yml:

`image: docker:19.03.5`

`services:`

`- docker:dind`

`stages:`

`- build`

`- test`

`variables:`

`GITLAB_HOST:` [`mygitlab.com`](https://mygitlab.com)

`GITLAB_PORT: 1111`

`CONTAINER_TEST_IMAGE: $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_SLUG`

`POSTGRES_USER: test`

`POSTGRES_PASSWORD: test-password`

`POSTGRES_DB: test`

`POSTGRES_HOST: postgres`

`before_script:`

`- docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $GITLAB_HOST:$GITLAB_PORT`

`after_script:`

`- docker logout $GITLAB_HOST:$GITLAB_PORT`

`build:`

`stage: build`

`script:`

`- docker build -t $CONTAINER_TEST_IMAGE -f docker/app/DockerFile .`

`- docker push $CONTAINER_TEST_IMAGE`

`test:`

`stage: test`

`services:`

`- name: postgres:latest`

`- name: redis:latest`

`variables:`

`DATABASE_URL: postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}/${POSTGRES_DB}`

`RAILS_ENV: test`

`REDIS_URL: redis://redis:6379`

`script:`

`- docker run --env RAILS_ENV --env REDIS_URL --env DATABASE_URL --rm $CONTAINER_TEST_IMAGE rails db:migrate`

`- docker run --env RAILS_ENV --env REDIS_URL --env DATABASE_URL --rm $CONTAINER_TEST_IMAGE rails test`

The problem is that, I cannot connect to gitlab postgres service (`could not translate host name "postgres" to address: Name or service not known`). I suspect the reason is that i'm running docker image from script and the containers (app and postgres) can't see each other (they are not in same network?). Below is a modifiaction for test role:

`test:`

`stage: test`

`image: $CONTAINER_TEST_IMAGE`

`services:`

`- name: postgres:latest`

`- name: redis:latest`

`variables:`

`DATABASE_URL: postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}/${POSTGRES_DB}`

`RAILS_ENV: test`

`REDIS_URL: redis://redis:6379`

`script:`

`- rails db:migrate`

`- rails test`

In this case, i'm facing an error  rails command not found (bundle exec also does not work). Maybe someone has solved such problem?
## [9][trying to ensure I properly understand how caxlsx gem works for spreadsheets](https://www.reddit.com/r/rails/comments/gkansz/trying_to_ensure_i_properly_understand_how_caxlsx/)
- url: https://www.reddit.com/r/rails/comments/gkansz/trying_to_ensure_i_properly_understand_how_caxlsx/
---
I'm using this gem to make a spreadsheet:

https://github.com/caxlsx/caxlsx_rails

1.  I just had a question, the caxlsx didn't appear to indicate it explicitly, but when I do an `add_row ["sadf", "adsfsd"]` each of those command separated strings is a column, correct?  I guess it's supposed to be really evident, but surprised that the github page for it doesn't even mention the word "column"

2.  Another question is, I"ll be adding some comma separated data per single cell as well such that I'll want in a single cell of the spreadsheet data like 

`"Bob, Cindy, Layla, Ahmad"` 

But the issue is say I have some values that have first and last names with comma, so we will have a "Smith, Joe" appended to the above and still in that single cell:

    "Bob, Cindy, Layla, Ahmad, Smith, Joe"

I guess I want my question 1 and 2 to work in conjunction together, such that I can properly enter in comma-separated values in a single cell properly in the spreadsheet, but also distinguish the separation between "Smith, Joe" in the above and the names preceding it, preferably with a comma too. 

If there is any misunderstanding in how data is inserted into a row and it's columns using caxlsx, please let me know too, thanks!
## [10][How to Fix Slow Code in Ruby](https://www.reddit.com/r/rails/comments/gjwyl9/how_to_fix_slow_code_in_ruby/)
- url: https://engineering.shopify.com/blogs/engineering/how-fix-slow-code-ruby
---

## [11][301 redirect](https://www.reddit.com/r/rails/comments/gk6soo/301_redirect/)
- url: https://www.reddit.com/r/rails/comments/gk6soo/301_redirect/
---
I'm using Nginx and I would like to do a 301 redirect to one domain as I have 4. I'm using the word example as my domain example in this cry for help 😥. 

I don't understand why it is not going to the root file path when I use 301 redirect, can someone help please? 

Inside the server block I have this:

server_name www.example.com example.com www.example.co.uk;

return 301 example.co.uk;

root /var/www/example;
