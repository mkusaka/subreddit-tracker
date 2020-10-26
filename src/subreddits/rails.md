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
## [3][How to merge SimpleCov results on SemaphoreCI 2.0](https://www.reddit.com/r/rails/comments/jic8m8/how_to_merge_simplecov_results_on_semaphoreci_20/)
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
           date: 2020-10-26
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
## [4][Getting a job as a rails developer](https://www.reddit.com/r/rails/comments/ji5n8y/getting_a_job_as_a_rails_developer/)
- url: https://www.reddit.com/r/rails/comments/ji5n8y/getting_a_job_as_a_rails_developer/
---
I'm curious when is a good time to get a job as a rails developer? 

Mainly, I'm wondering how much experience you should have before you can feasibly land a rails gig. 

I've been building websites for 4 years, coding for 2 years, and doing rails for 1 year. I've built a few rails apps at this point, I also work as a software PM where I also do some light coding to contribute to the app we run. 

I still have a lot to learn, but I'm also starting to feel (for the first time ever) confident in my coding skills . I'm good with HTML, CSS, jquery ruby and rails. I've deployed apps using AWS to heroku with mailers and all sorts of fun working features. 

However, when I look up rails jobs, some of them say 2-5 years experience is a must, and that instantly makes me doubt myself. Does my experience sound like I might be able to get a job as a rails developer. 

LIke I said, I am a PM already and I love that position, but I've totally caught the coding bug and I find it's all I ever think about or want to do, wondering if I should try and make the jump to DEV.
## [5][how does create! method work?](https://www.reddit.com/r/rails/comments/jidiju/how_does_create_method_work/)
- url: https://www.reddit.com/r/rails/comments/jidiju/how_does_create_method_work/
---
Hello, I have a serious situation.

My projects are using rails 5.2.3 with Delayed Job and a job creates a record and then another job perform with find a record by id.

The gap between perform another job and create! method called is maybe within 0.5s.

But sometimes another job could not find a record by id.

So I was wondering create! method does not guarantee that a transaction is committed?

If not, how guarantee after commit?
## [6][How to sanitize some field values before save?](https://www.reddit.com/r/rails/comments/jicrlg/how_to_sanitize_some_field_values_before_save/)
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
## [7][MVP : start with Devise (or equivalent) or auth0 (or equivalent) ?](https://www.reddit.com/r/rails/comments/jhqo9r/mvp_start_with_devise_or_equivalent_or_auth0_or/)
- url: https://www.reddit.com/r/rails/comments/jhqo9r/mvp_start_with_devise_or_equivalent_or_auth0_or/
---
If you had to build a MVP for a new startup, will you choose an authentication provider (auth0 or whatever  equivalent), or will you integrate it inside your application (Clearance, Devise, Sorcery...). I have this choice to do and I find it hard to make the right decision. In theory less software is better, but I feel I loose control over what is happening, and the main interest of Rails is precisely ability to integrate anything. Thanks for your thoughts!
## [8][Help: Good way to handle mail templates stored in database](https://www.reddit.com/r/rails/comments/jhomuk/help_good_way_to_handle_mail_templates_stored_in/)
- url: https://www.reddit.com/r/rails/comments/jhomuk/help_good_way_to_handle_mail_templates_stored_in/
---
Hi all, 

I have an administration system for customers, where they can control the product we are providing. We need to send emails on different actions, to confirm then changes has been done. 

Sometimes it need to sent an email with different content to 3 receptions. 

* The user who did the action
* The administrator(s) for the users customer
* Our administrators (or helpdesk shared email)

The emails contains variables, with data from the execution. 

Currently  we are handing it directly in rails. But now we have partners which also have customers, and might want to make changes in the templates. 

Therefore, we need someway smart to handle all the mail templates in a user friendly interface together with the rest of the application. Maybe some markdown support in a simple wysiwyg editor, however that is more a frontend question. First we need the backend to support til mail templates. 

To spice it, we also need to handle multiple languages, however currently we only support two. 

We are using Rails 6 api only as backend and React as frontend.
## [9][What are some of the most beautiful rails made websites you've visited?](https://www.reddit.com/r/rails/comments/jhf166/what_are_some_of_the_most_beautiful_rails_made/)
- url: https://www.reddit.com/r/rails/comments/jhf166/what_are_some_of_the_most_beautiful_rails_made/
---

## [10][Can I use Pundit to shadow ban?](https://www.reddit.com/r/rails/comments/jhozhz/can_i_use_pundit_to_shadow_ban/)
- url: https://www.reddit.com/r/rails/comments/jhozhz/can_i_use_pundit_to_shadow_ban/
---
Can you make me an example to hide the posts only for the shadow banned user using pundit gem?
## [11][Automatically check whether the current user is allowed to modify a record](https://www.reddit.com/r/rails/comments/jhfm2h/automatically_check_whether_the_current_user_is/)
- url: https://www.reddit.com/r/rails/comments/jhfm2h/automatically_check_whether_the_current_user_is/
---
Solution: I'm disappointed at myself. It was right in [the docs](https://guides.rubyonrails.org/action_controller_overview.html#filters):

```rb
class BoardsController &lt; ApplicationController
    before_action :authenticate_user!
    before_action :must_be_board_owner
    skip_before_action :must_be_board_owner, only: [:new, :create, :index]

    # ...

    def must_be_board_owner
        if current_user != Board.find_by(slug: params[:slug]).user
            redirect_to root_path
        end
    end
end
```

---

Hi. I currently have the following problem:
On my controllers, I have this before_action filter:

```rb
# as you can see, I'm using Devise
before_action :authenticate_user!
```

This checks whether the user is signed in or not.
However, I also need to check whether the current user is allowed to see/modify the current record.
Right now, I have this (horrible) method:

```rb
    private def redirect_unless_board_owner! owner
        redirect_to root_path unless current_user == owner
    end
```

Using it in an action looks like this:

```rb
        @board = Board.find_by slug: params[:board_slug]
        redirect_unless_board_owner! @board.user

        @task = Task.find_by slug: params[:slug]
```

Is there any way that I can make this check a before_action filter? This just looks so ugly in my code, ugh. Thanks.
## [12][Some columns are not created after going to MySQL](https://www.reddit.com/r/rails/comments/jh9oaz/some_columns_are_not_created_after_going_to_mysql/)
- url: https://www.reddit.com/r/rails/comments/jh9oaz/some_columns_are_not_created_after_going_to_mysql/
---
I made a bunch of migrations, and this is my schema :

    # This file is auto-generated from the current state of the database. Instead
    # of editing this file, please use the migrations feature of Active Record to
    # incrementally modify your database, and then regenerate this schema definition.
    #
    # This file is the source Rails uses to define your schema when running `rails
    # db:schema:load`. When creating a new database, `rails db:schema:load` tends to
    # be faster and is potentially less error prone than running all of your
    # migrations from scratch. Old migrations may fail to apply correctly if those
    # migrations use external dependencies or application code.
    #
    # It's strongly recommended that you check this file into your version control system.
    
    ActiveRecord::Schema.define(version: 2020_10_14_075849) do
    
      create_table "comments", options: "ENGINE=InnoDB DEFAULT CHARSET=utf8", force: :cascade do |t|
        t.text "body"
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.integer "post_id"
        t.integer "user_id"
        t.index ["post_id"], name: "index_comments_on_post_id"
        t.index ["user_id"], name: "index_comments_on_user_id"
      end
    
      create_table "favorites", options: "ENGINE=InnoDB DEFAULT CHARSET=utf8", force: :cascade do |t|
        t.string "favoritable_type", null: false
        t.bigint "favoritable_id", null: false
        t.string "favoritor_type", null: false
        t.bigint "favoritor_id", null: false
        t.string "scope", default: "favorite", null: false
        t.boolean "blocked", default: false, null: false
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.index ["blocked"], name: "index_favorites_on_blocked"
        t.index ["favoritable_id", "favoritable_type"], name: "fk_favoritables"
        t.index ["favoritable_type", "favoritable_id"], name: "index_favorites_on_favoritable_type_and_favoritable_id"
        t.index ["favoritor_id", "favoritor_type"], name: "fk_favorites"
        t.index ["favoritor_type", "favoritor_id"], name: "index_favorites_on_favoritor_type_and_favoritor_id"
        t.index ["scope"], name: "index_favorites_on_scope"
      end
    
      create_table "likes", options: "ENGINE=InnoDB DEFAULT CHARSET=utf8", force: :cascade do |t|
        t.bigint "post_id", null: false
        t.bigint "user_id", null: false
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.index ["post_id"], name: "index_likes_on_post_id"
        t.index ["user_id"], name: "index_likes_on_user_id"
      end
    
      create_table "posts", options: "ENGINE=InnoDB DEFAULT CHARSET=utf8", force: :cascade do |t|
        t.string "image"
        t.text "caption"
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.integer "user_id"
        t.index ["user_id"], name: "index_posts_on_user_id"
      end
    
      create_table "users", options: "ENGINE=InnoDB DEFAULT CHARSET=utf8", force: :cascade do |t|
        t.string "email", default: "", null: false
        t.string "encrypted_password", default: "", null: false
        t.string "reset_password_token"
        t.datetime "reset_password_sent_at"
        t.datetime "remember_created_at"
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.string "username"
        t.index ["email"], name: "index_users_on_email", unique: true
        t.index ["reset_password_token"], name: "index_users_on_reset_password_token", unique: true
        t.index ["username"], name: "index_users_on_username", unique: true
      end
    
      add_foreign_key "likes", "posts"
      add_foreign_key "likes", "users"
    end

Then, I added this to the `database.yml` file to make it work with MySQL :

    default: &amp;default
      adapter: mysql2
      encoding: utf8
      pool: &lt;%= ENV.fetch("RAILS_MAX_THREADS") { 5 } %&gt;
      username: railstagram
      password: 
      socket: /tmp/mysql.sock

First, I'm a bit confused why it didn't ask for port or host, but I assume it's handled somewhere else.

By the way, when I run `rails db:create` and `rails db:migrate` everything seems just fine, but when I create a user, it seems it doesn't care for the *username* field.

This is a [link](https://github.com/prp-e/railstagram) to my project. I need to deploy it tonight and I have this problem with MySQL. My development environment is macOS Catalina.
