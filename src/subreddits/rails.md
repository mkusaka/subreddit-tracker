# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/
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
## [3][HOWTO: highlight link_to current_page](https://www.reddit.com/r/rails/comments/jiznev/howto_highlight_link_to_current_page/)
- url: https://www.reddit.com/r/rails/comments/jiznev/howto_highlight_link_to_current_page/
---
Short post on how to highlight `link_to` `current_page`: [https://blog.corsego.com/2020/10/ruby-on-rails-highlight-linkto.html](https://blog.corsego.com/2020/10/ruby-on-rails-highlight-linkto.html) I hope you find it useful :)

https://preview.redd.it/awsrnkrnlmv51.png?width=338&amp;format=png&amp;auto=webp&amp;s=05b5ace0671b8a100ada1fb7fe4d7082230fe532
## [4][Good guides/articles regarding Rails 6 deployment to AWS](https://www.reddit.com/r/rails/comments/jiob7g/good_guidesarticles_regarding_rails_6_deployment/)
- url: https://www.reddit.com/r/rails/comments/jiob7g/good_guidesarticles_regarding_rails_6_deployment/
---
Hello everyone!

I decided to keep all my Rails stuff inside of one service(AWS S3 and hosting on AWS) but I struggle to find a good guide on how to deploy Rails to AWS and that the guide is up to date.

It would be really cool if you could point me towards some good articles about this topic.

Thank you!
## [5][Any good courses/tutorials for Rails API + Vue out there?](https://www.reddit.com/r/rails/comments/jiozti/any_good_coursestutorials_for_rails_api_vue_out/)
- url: https://www.reddit.com/r/rails/comments/jiozti/any_good_coursestutorials_for_rails_api_vue_out/
---
Hi 

I'd like to get more into working with the combination of a Rails API and Vue, but I'm looking to follow a tutorial or course to get the hang of it some more before I dive into my own project. 

I already followed a [tutorial from WebCrunch](https://web-crunch.com/posts/ruby-on-rails-api-vue-js) but felt like a lot of corners were being cut there. 

Anybody have some up-to-date recommendations for where I could get started? Would really appreciate it!
## [6][Introducing jQuery in Rails 6 Using Webpacker](https://www.reddit.com/r/rails/comments/jiyqvc/introducing_jquery_in_rails_6_using_webpacker/)
- url: https://www.reddit.com/r/rails/comments/jiyqvc/introducing_jquery_in_rails_6_using_webpacker/
---
Hello Everyone,

First release candidate of Rails 6 is out with exciting features and refinements in existing features. Today we will look into one of those features, which is webpacker. It is now by default in Rails 6 as a Javascript Compiler.

Here’s a detailed blog on [Introducing jQuery in Rails 6 Using Webpacker](https://www.botreetechnologies.com/blog/introducing-jquery-in-rails-6-using-webpacker)

Hope this will be helpful to #railscommunity.
## [7][Linters](https://www.reddit.com/r/rails/comments/jij648/linters/)
- url: https://www.reddit.com/r/rails/comments/jij648/linters/
---
Hi, guys! Which linter do you use? Rubocop or standardrb? Or other?
## [8][devise host address](https://www.reddit.com/r/rails/comments/jifjwr/devise_host_address/)
- url: https://www.reddit.com/r/rails/comments/jifjwr/devise_host_address/
---
I have an app and I configured it to send confirmation data using GMail. It works fine, but instead of sending confirmation like "[mysite.com/users/](https://mysite.com/users/)..." it sends it from [gmail.com/users/](https://gmail.com/users/)... . 

How can I fix this? Thanks.
## [9][How to merge SimpleCov results on SemaphoreCI 2.0](https://www.reddit.com/r/rails/comments/jic8m8/how_to_merge_simplecov_results_on_semaphoreci_20/)
- url: https://www.reddit.com/r/rails/comments/jic8m8/how_to_merge_simplecov_results_on_semaphoreci_20/
---
Original post with images on [https://anonoz.github.io/tech/2020/10/26/simplecov-semaphoreci.html](https://anonoz.github.io/tech/2020/10/26/simplecov-semaphoreci.html)

----
We have been using [SemaphoreCI](https://semaphoreci.com) for over a year now. 

We like Semaphore a lot because of their bare metal performance and their fair per-second pricing, which is different from other CIs charging by maximum parallelisms.

Having sufficient test coverage has helped us in staying clear of nasty problems along the road, I highly encourage other projects to keep it in check too!

Feel free to copy and paste the following scripts and YAMLs, if there are any error and problems, please contact me and let me know.

## Step 1: Setup SimpleCov

Include it in your `Gemfile`, please find the latest version number and replace it: 

    gem 'simplecov', '~&gt; 0.18.0', require: false


To make sure SimpleCov tracks the code coverage in Capybara test, paste this in your `bin/rails` before the existing code:

    #!/usr/bin/env ruby
    if ENV['RAILS_ENV'] == 'test'
      require 'simplecov'
      SimpleCov.start 'rails'
    end
    # ^^ Copy Above
    
    APP_PATH = File.expand_path('../../config/application', __FILE__)
    require_relative '../config/boot'
    require 'rails/commands'

In your `specs/spec_helper.rb`, paste this between all the `require` and other configuration blocks:

    if ENV['CI']
      SimpleCov.command_name "#{ENV['SEMAPHORE_JOB_ID']}-#{ENV['TEST_ATTEMPT']}"
    end
    
    SimpleCov.start 'rails'

## Step 2: Here are the .yaml and .rb

Paste this in `scripts/simplecov-collate.rb` in your project directory:

    #!/usr/bin/env ruby
    require 'simplecov'
    
    SimpleCov.collate Dir["simplecov-resultset/*/.resultset.json"] do
      formatter SimpleCov::Formatter::MultiFormatter.new([
        SimpleCov::Formatter::SimpleFormatter,
        SimpleCov::Formatter::HTMLFormatter
      ])
    end

I assume:

1. You already have SimpleCov gem included and configured in your rspec.
2. Your project directory is mounted in `/project-mount` in the Docker image that you used to test. SimpleCov somehow uses absolute path inside their result JSON files, and this is the way to make the merge successful.
3. SemaphoreCI git clones into `~/reponame` when `checkout` command is being run.

Update your `.semaphore/semaphore.yml`:

1. The `epilogue` that runs after every parallel tests, and;
2. The block that runs after ALL the tests.

.semaphore/semaphore.yml: (something weird with reddit markdown indent)

    # .semaphore/semaphore.yml
    blocks:
      - name: 'Run tests'
        task:
          jobs:
            - name: 'E2E Part 1'
            - name: 'E2E Part 2'
            - name: 'E2E Part 3'
            - name: 'Models'
            - name: 'API'
          # Copy the commands in the epilogue, this runs after every different test.
          epilogue:
            always:
              commands:
                - |
                    if [ -d coverage ]; then
                      artifact push job --expire-in 3d --force coverage;
                      COV_DIRNAME=simplecov-resultset/$SEMAPHORE_JOB_ID;
                      sudo chown -R $USER:$USER coverage/;
                      mkdir -p simplecov-resultset;
                      mv coverage $COV_DIRNAME;
                      artifact push workflow --expire-in 2w --force simplecov-resultset;
                    fi
      
      # Feel free to copy this block below!
      - name: 'Check codebase quality'
        task:  
          jobs:
            - name: 'Test coverage'
              commands:
                - checkout
                - |
                    if ! rbenv install -s; then
                      git -C /home/semaphore/.rbenv/plugins/ruby-build pull &amp;&amp;
                      rbenv install;
                    fi
                - sudo mkdir -p /project-mount
         - sudo chown -R $USER:$USER /project-mount
                - cd ~
                - shopt -s dotglob; mv ~/reponame/* /project-mount
                - cd /project-mount
                - artifact pull workflow simplecov-resultset
                - gem install simplecov
                - ruby scripts/simplecov-collate.rb
                - cd coverage/
                - cd /project-mount
                - zip -r coverage-$SEMAPHORE_GIT_SHA coverage/
                - artifact push job --expire-in 2w coverage-$SEMAPHORE_GIT_SHA.zip

Once you are done, commit and push to trigger a CI run.

## Step 3: Get your test coverage!

If the configurations are correct and your test suites pass as usual, this is how you can find your test coverage results.

Go into the workflow for the latest commit, you should see there is a new block at the end after the parallel tests.

Click **Job artifacts** between the console outputs and the job name. That is not the most obvious place, I know. By the time you read this, the UI may have changed again, but it should be somewhere in the page.

Download and unzip it, then open `index.html` file in your browser to view.

Unlike CircleCI, every artifact on SemaphoreCI cannot be viewed directly in the browser, but rather must be downloaded. It is not very convenient.

---

1. **Why `gem install simplecov` separately instead of just using the Docker image I have been using?**

    The Docker images are usually large. Downloading just the gems needed is much quicker than pulling the image all over again. In our case, it takes almost a minute to pull the testing Docker image down, but only 20 seconds to run the whole SimpleCov merging job - from git cloning, to gem install, to uploading artifact.

2. **What if we did not run tests in Docker image?**

    You should be able to replace `/project-mount` with the directory pah the project is git cloned into in previous test jobs.

Please let me know if there are better ideas on how to go about this!
## [10][how does create! method work?](https://www.reddit.com/r/rails/comments/jidiju/how_does_create_method_work/)
- url: https://www.reddit.com/r/rails/comments/jidiju/how_does_create_method_work/
---
Hello, I have a serious situation.

My projects are using rails 5.2.3 with Delayed Job and a job creates a record and then another job perform with find a record by id.

The gap between perform another job and create! method called is maybe within 0.5s.

But sometimes another job could not find a record by id.

So I was wondering create! method does not guarantee that a transaction is committed?

If not, how guarantee after commit?
## [11][How to sanitize some field values before save?](https://www.reddit.com/r/rails/comments/jicrlg/how_to_sanitize_some_field_values_before_save/)
- url: https://www.reddit.com/r/rails/comments/jicrlg/how_to_sanitize_some_field_values_before_save/
---
Hello, Rails beginner here. 

Is there any built in method in Rails that works like Django model's clean() method?

I'm looking for a way to **sanitize user provided value for a url field**, by removing schema/trailing slash etc.

I'm saving a `Project` model &amp; `Site` models using a nested form. Cleaning it in the controller requires to loop through `params[:sites_attributes].` I can clean the url field in a validation method but it doesn't feel right.

What is the best practice regarding sanitising a single or multiple fields before save? Where should i do it?

\# Project Model

`class Project &lt; ApplicationRecord`

  `has_many :sites`

`end` 

\# Site Model

`class Site &lt; ApplicationRecord`

 `belongs_to :project` 

  `validate :must_be_valid_url`

`end`

\# Controller

`def update`  
 `@project = Project.find(params[:id])`  
 `if @project.update(project_params)`  
`redirect_to edit_project_path(@project)`  
 `else`  
`render :edit`  
 `end`  
`end`

`def project_params`  
 `params.require(:project).permit(sites_attributes: [:id, :url, ....])`  
`end`
## [12][Getting a job as a rails developer](https://www.reddit.com/r/rails/comments/ji5n8y/getting_a_job_as_a_rails_developer/)
- url: https://www.reddit.com/r/rails/comments/ji5n8y/getting_a_job_as_a_rails_developer/
---
I'm curious when is a good time to get a job as a rails developer? 

Mainly, I'm wondering how much experience you should have before you can feasibly land a rails gig. 

I've been building websites for 4 years, coding for 2 years, and doing rails for 1 year. I've built a few rails apps at this point, I also work as a software PM where I also do some light coding to contribute to the app we run. 

I still have a lot to learn, but I'm also starting to feel (for the first time ever) confident in my coding skills . I'm good with HTML, CSS, jquery ruby and rails. I've deployed apps using AWS to heroku with mailers and all sorts of fun working features. 

However, when I look up rails jobs, some of them say 2-5 years experience is a must, and that instantly makes me doubt myself. Does my experience sound like I might be able to get a job as a rails developer. 

LIke I said, I am a PM already and I love that position, but I've totally caught the coding bug and I find it's all I ever think about or want to do, wondering if I should try and make the jump to DEV.
