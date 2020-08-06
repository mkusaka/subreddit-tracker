# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/
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
## [2][Has the tech for hey.com been officially released yet?](https://www.reddit.com/r/rails/comments/i4mi5t/has_the_tech_for_heycom_been_officially_released/)
- url: https://www.reddit.com/r/rails/comments/i4mi5t/has_the_tech_for_heycom_been_officially_released/
---
If so, what is it? If not, when will it be?
## [3][Docker Rails app error with webpacker's config? "Webpacker::Manifest::MissingEntryError](https://www.reddit.com/r/rails/comments/i4koht/docker_rails_app_error_with_webpackers_config/)
- url: https://www.reddit.com/r/rails/comments/i4koht/docker_rails_app_error_with_webpackers_config/
---
Hi there. Trying to display hello world on my homepage to make sure my app is going to run properly. I was able to get to the Rails welcome page, but as soon as I created the first controller for the homepage I'm getting this error output:

     Started GET "/" for 172.20.0.1 at 2020-08-06 03:46:18 +0000
    web_1  | Cannot render console from 172.20.0.1! Allowed networks: 127.0.0.0/127.255.255.255, ::1
    web_1  | Processing by HomeController#index as HTML
    web_1  |   Rendering home/index.html.erb within layouts/application
    web_1  |   Rendered home/index.html.erb within layouts/application (Duration: 36.4ms | Allocations: 297)
    web_1  | /usr/local/bundle/gems/sprockets-rails-3.2.1/lib/sprockets/rails/helper.rb:355: warning: Using the last argument as keyword parameters is deprecated; maybe ** should be added to the call
    web_1  | /usr/local/bundle/gems/sprockets-4.0.2/lib/sprockets/base.rb:118: warning: The called method `[]' is defined here
    web_1  | [Webpacker] Compiling...
    web_1  | [Webpacker] Compilation failed:
    web_1  | yarn run v1.22.4
    web_1  | info Visit https://yarnpkg.com/en/docs/cli/run for documentation about this command.
    web_1  |
    web_1  |
    web_1  | error Command "webpack" not found.
    web_1  |
    web_1  | Completed 500 Internal Server Error in 3013ms (ActiveRecord: 0.0ms | Allocations: 10122)
    web_1  |
    web_1  |
    web_1  |
    web_1  | ActionView::Template::Error (Webpacker can't find application.js in /CanaBnb/public/packs/manifest.json. Possible causes:
    web_1  | 1. You want to set webpacker.yml value of compile to true for your environment
    web_1  |    unless you are using the `webpack -w` or the webpack-dev-server.
    web_1  | 2. webpack has not yet re-run to reflect updates.
    web_1  | 3. You have misconfigured Webpacker's config/webpacker.yml file.
    web_1  | 4. Your webpack configuration is not creating a manifest.
    web_1  | Your manifest contains:
    web_1  | {
    web_1  | }
    web_1  | ):
    web_1  |      6:     &lt;%= csp_meta_tag %&gt;
    web_1  |      7:
    web_1  |      8:     &lt;%= stylesheet_link_tag 'application', media: 'all', 'data-turbolinks-track': 'reload' %&gt;
    web_1  |      9:     &lt;%= javascript_pack_tag 'application', 'data-turbolinks-track': 'reload' %&gt;
    web_1  |     10:   &lt;/head&gt;
    web_1  |     11:
    web_1  |     12:   &lt;body&gt;
    web_1  |
    web_1  | app/views/layouts/application.html.erb:9

My eyes go to those "Possible causes" it lists, but I'm not sure if that's it because I also see the Webpacker compilation failed above that. Also the error Command "webpack" not found. It's like it has multiple reasons for this error and I'm just not sure how to fix this. Any help is greatly appreciated. I made a github repo  [here](https://github.com/Kyle-Williamson-Dev/CanaBnb)

&amp;#x200B;
## [4][Summernote on Rails 6](https://www.reddit.com/r/rails/comments/i4ovvm/summernote_on_rails_6/)
- url: https://www.reddit.com/r/rails/comments/i4ovvm/summernote_on_rails_6/
---
The original gem is for rails 3, I can't really figure out how to make it work on Rails 6. Can somebody help?

The guide I am trying to follow:

[https://github.com/summernote/summernote-rails](https://github.com/summernote/summernote-rails)
## [5][Safe Way to Save HTML data to Database](https://www.reddit.com/r/rails/comments/i49iii/safe_way_to_save_html_data_to_database/)
- url: https://www.reddit.com/r/rails/comments/i49iii/safe_way_to_save_html_data_to_database/
---
I want users to be able to save HTML editor's data to the database. Which would be the best approach to do it?
## [6][Trying to use Docker for Rails - not loading web correctly](https://www.reddit.com/r/rails/comments/i4e9du/trying_to_use_docker_for_rails_not_loading_web/)
- url: https://www.reddit.com/r/rails/comments/i4e9du/trying_to_use_docker_for_rails_not_loading_web/
---
Hi there. I've been having issues just getting things going with Rails and Docker. I'm getting the following output after following basically everything in the Docker-compose quickstart guide [here](https://docs.docker.com/compose/rails/).

Looks like the first thing I see is the web\_1 line that says exec: rails: not found. Tried Googling and can't find a solution that works. I've had this working before on my portfolio, but then I merged some changes someone was trying to help with and that got all messed up. So I'm just trying to create one of my projects FOR my portfolio just to start fresh with a Rails app to try again. I'll post the outcome below as well as my dockerfile and docker-compose.yml.

    Kyle_@LAPTOP-KSPCT1K1 MINGW64 /c/users/kyle_/desktop/canabnb (master)
    $ docker-compose up
    Creating canabnb_db_1  ... done
    Creating canabnb_web_1 ... done
    Attaching to canabnb_db_1, canabnb_web_1
    web_1  | /usr/bin/entrypoint.sh: line 8: exec: rails: not found
    db_1   | The files belonging to this database system will be owned by user "postgres".
    db_1   | This user must also own the server process.
    db_1   |
    db_1   | The database cluster will be initialized with locale "en_US.utf8".
    db_1   | The default database encoding has accordingly been set to "UTF8".
    db_1   | The default text search configuration will be set to "english".
    db_1   |
    db_1   | Data page checksums are disabled.
    db_1   |
    db_1   | fixing permissions on existing directory /var/lib/postgresql/data ... ok
    db_1   | creating subdirectories ... ok
    db_1   | selecting dynamic shared memory implementation ... posix
    db_1   | selecting default max_connections ... 100
    db_1   | selecting default shared_buffers ... 128MB
    db_1   | selecting default time zone ... Etc/UTC
    db_1   | creating configuration files ... ok
    canabnb_web_1 exited with code 127
    db_1   | running bootstrap script ... ok
    db_1   | performing post-bootstrap initialization ... ok
    db_1   | syncing data to disk ... ok
    db_1   |
    db_1   |
    db_1   | Success. You can now start the database server using:
    db_1   |
    db_1   |     pg_ctl -D /var/lib/postgresql/data -l logfile start
    db_1   |
    db_1   | initdb: warning: enabling "trust" authentication for local connections
    db_1   | You can change this by editing pg_hba.conf or using the option -A, or
    db_1   | --auth-local and --auth-host, the next time you run initdb.
    db_1   | waiting for server to start....2020-08-05 21:20:26.416 UTC [46] LOG:  starting PostgreSQL 12.3 (Debian 12.3-1.pgdg100+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 8.3.0-6) 8.3.0, 64-bit
    db_1   | 2020-08-05 21:20:26.420 UTC [46] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
    db_1   | 2020-08-05 21:20:26.451 UTC [47] LOG:  database system was shut down at 2020-08-05 21:20:26 UTC
    db_1   | 2020-08-05 21:20:26.459 UTC [46] LOG:  database system is ready to accept connections
    db_1   |  done
    db_1   | server started
    db_1   | CREATE DATABASE
    db_1   |
    db_1   |
    db_1   | /usr/local/bin/docker-entrypoint.sh: ignoring /docker-entrypoint-initdb.d/*
    db_1   |
    db_1   | waiting for server to shut down...2020-08-05 21:20:26.764 UTC [46] LOG:  received fast shutdown request
    db_1   | .2020-08-05 21:20:26.766 UTC [46] LOG:  aborting any active transactions
    db_1   | 2020-08-05 21:20:26.767 UTC [46] LOG:  background worker "logical replication launcher" (PID 53) exited with exit code 1
    db_1   | 2020-08-05 21:20:26.768 UTC [48] LOG:  shutting down
    db_1   | 2020-08-05 21:20:26.786 UTC [46] LOG:  database system is shut down
    db_1   |  done
    db_1   | server stopped
    db_1   |
    db_1   | PostgreSQL init process complete; ready for start up.
    db_1   |
    db_1   | 2020-08-05 21:20:26.879 UTC [1] LOG:  starting PostgreSQL 12.3 (Debian 12.3-1.pgdg100+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 8.3.0-6) 8.3.0, 64-bit
    db_1   | 2020-08-05 21:20:26.880 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
    db_1   | 2020-08-05 21:20:26.880 UTC [1] LOG:  listening on IPv6 address "::", port 5432
    db_1   | 2020-08-05 21:20:26.885 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
    db_1   | 2020-08-05 21:20:26.903 UTC [64] LOG:  database system was shut down at 2020-08-05 21:20:26 UTC
    db_1   | 2020-08-05 21:20:26.909 UTC [1] LOG:  database system is ready to accept connections

Dockerfile:

    FROM ruby:2.7.1
    RUN apt-get update -qq &amp;&amp; apt-get install -y nodejs postgresql-client WORKDIR /CanaBnb COPY Gemfile /CanaBnb/Gemfile COPY Gemfile.lock /CanaBnb/Gemfile.lock RUN gem install bundler COPY . /CanaBnb
     Add a script to be executed every time the container starts.
    COPY entrypoint.sh /usr/bin/ RUN chmod +x /usr/bin/entrypoint.sh ENTRYPOINT ["entrypoint.sh"] EXPOSE 3000
     Start the main process.
    CMD ["rails", "server", "-b", "0.0.0.0"]

docker-compose.yml:

    version: '3'
    services: 
    db: 
    image: postgres 
    environment: 
    POSTGRES_USER: postgres 
    POSTGRES_PASSWORD: password 
    POSTGRES_DB: CanaBnb_development 
    web: build: . 
    volumes:       - .:/CanaBnb 
    ports:       - "3000:3000"

If you need more than this to help, just let me know. I've been stuck on docker issues for a long time now just trying to get going after graduating bootcamp in February. Thing is I got my portfolio app working before and just decided to try a change (can't recall for the life of me what that change was) that messed everything up. I'm still so new to this and especially new to Docker. Thank you in advance.
## [7][Mina Deploy Precompile Issue](https://www.reddit.com/r/rails/comments/i49wo8/mina_deploy_precompile_issue/)
- url: https://www.reddit.com/r/rails/comments/i49wo8/mina_deploy_precompile_issue/
---
 

I am running a Debian 10 (Buster) Server for my Rails 6.0.3, Ruby  2.6.3, Mina 1.2.3 app. When I run Mina deploy it errors out but does not  give me enough to help me understand what I need to do to resolve the  issue. And thanks to u/[d4be4st](https://www.reddit.com/user/d4be4st/) for getting me this far!!

Mina Deploy output

    Precompiling asset files        
    $ RAILS_ENV="production" bundle exec rake assets:precompile        
    Traceback (most recent call last):         
    3: from /var/www/site/tmp/build-15965768145748/vendor/bundle/ruby/2.6.0/bin/ruby_executable_hooks:24:in `&lt;main&gt;'         
    2: from /var/www/site/tmp/build-15965768145748/vendor/bundle/ruby/2.6.0/bin/ruby_executable_hooks:24:in `eval'         
    1: from /usr/local/rvm/gems/ruby-2.6.3/bin/rake:23:in `&lt;main&gt;'        
    /usr/local/rvm/gems/ruby-2.6.3/bin/rake:23:in `load': cannot load such file -- /usr/local/rvm/rubies/ruby-2.6.3/lib/ruby/gems/2.6.0/specifications/default/exe/rake (LoadError)  
    !     ERROR: Deploy failed. 

Deploy.rb

    require 'mina/bundler' 
    require 'mina/rails' 
    require 'mina/git' 
    require 'mina/rvm'  
    
    set :application_name, 'site' 
    set :domain, 'server' 
    set :user, 'deployer' 
    set :deploy_to, "/var/www/site" 
    set :repository, "git@github.com:git.git" 
    set :branch, 'master' 
    set :rvm_use_path, '/etc/profile.d/rvm.sh'  
    set :rails_env, 'production'
    
    set :shared_files, fetch(:shared_files, []).push('config/database.yml', 'config/secrets.yml')  
    
    task :remote_environment do 
       ruby_version = File.read('.ruby-version').strip
        raise "Couldn't determine Ruby version: Do you have a file .ruby-version in your project root?" if ruby_version.empty?
    
         invoke :'rvm:use', 'ruby-2.6.3' 
    end  
    
    desc "Deploys the current version to the server." 
    task :deploy do    
    deploy do
          invoke :'git:clone'
          invoke :'deploy:link_shared_paths'
          invoke :'bundle:install'
          invoke :'rails:db_migrate'
          invoke :'rails:assets_precompile'
          invoke :'deploy:cleanup'
    
          on :launch do 
          end
        end
     end  
    
    namespace :passenger do 
    desc "Restart Passenger"
        task :restart do
          queue %{
            echo "-----&gt; Restarting passenger"
            cd #{deploy_to}/current
            #{echo_cmd %[mkdir -p tmp]}
            #{echo_cmd %[touch tmp/restart.txt]}
          }
        end
     end 

I have made sure that yarn is installed and working on my Production  machine. When I run the precompile on the Production machine, it runs  through successfully. But still errors out in the Mina Deploy with no  change in the error text. I have looked into [this article](https://github.com/mina-deploy/mina/blob/master/docs/faq.md#--my_program-not-found-but-is-already-installed)   In doing so I have commented out the suggested code snippet just to  test, but that did not change anything. In looking at my bashrc, I do  not see any snippets with EXPORT or SOURCE at the beginning of them.

    RAILS_ENV="production" bundle exec rake assets:precompile
     yarn install v1.22.4 
    [1/4] Resolving packages... 
    [2/4] Fetching packages... 
    info fsevents@1.2.13: The platform "linux" is incompatible with this module. 
    info "fsevents@1.2.13" is an optional dependency and failed compatibility check. Excluding it from installation. 
    info fsevents@2.1.3: The platform "linux" is incompatible with this module. 
    info "fsevents@2.1.3" is an optional dependency and failed compatibility check. Excluding it from installation. 
    [3/4] Linking dependencies... 
    [4/4] Building fresh packages... 
    Done in 2.95s. 
    yarn install v1.22.4 
    [1/4] Resolving packages... 
    [2/4] Fetching packages... 
    info fsevents@1.2.13: The platform "linux" is incompatible with this module. 
    info "fsevents@1.2.13" is an optional dependency and failed compatibility check. Excluding it from installation. 
    info fsevents@2.1.3: The platform "linux" is incompatible with this module. 
    info "fsevents@2.1.3" is an optional dependency and failed compatibility check. Excluding it from installation. 
    [3/4] Linking dependencies... 
    [4/4] Building fresh packages... 
    Done in 2.85s. 
    Everything's up-to-date. Nothing to do
## [8][3 level nested resources](https://www.reddit.com/r/rails/comments/i4hit5/3_level_nested_resources/)
- url: https://www.reddit.com/r/rails/comments/i4hit5/3_level_nested_resources/
---
Briefly , I stumble on a stackoverflow comment which he says to have always at most 2 resources nested levels. He didn't explain much and I'm trying to figure out how would be the difference doing

`resources :x do`

`resources :y` 

`end`

`resources :y do`

   `resources :z`

`end`

 instead 

`resources :x do`

`resources :y do`

`resources :z`

`end`

`end`

&amp;#x200B;

I have a ToDo App which has a TodoItem resources nested in TodoList resources. Today I naively tried to nested these resources into my User resources. I spent a few hours solving the routing paths in my views and at end I stumble on a issue with forms. The paths were long and confusing.

I would like to understand technically the difference between the first and the second way of dealing with nested resources.
## [9][Is the cache information stored in the browser?](https://www.reddit.com/r/rails/comments/i4d05b/is_the_cache_information_stored_in_the_browser/)
- url: https://www.reddit.com/r/rails/comments/i4d05b/is_the_cache_information_stored_in_the_browser/
---
Hello sorry if this question is silly (I'm a newbie in rails), I'm trying to fragment cache some views and I don't know if the info that is cached is saved on the browser or in the rails server
## [10][Adding Snippets for easier rails coding](https://www.reddit.com/r/rails/comments/i42vyt/adding_snippets_for_easier_rails_coding/)
- url: https://www.reddit.com/r/rails/comments/i42vyt/adding_snippets_for_easier_rails_coding/
---
Hi just a quick question can someone tell me how i create shortcuts for  the below rails tags in VS code i was able to get emmet working its just the below tags wont work. Any guide how to set this up in rails to save time? 

    &lt;%= %&gt;

or 

    &lt;% %&gt;
## [11][How to log exceptions using Lograge as default logger](https://www.reddit.com/r/rails/comments/i43zyh/how_to_log_exceptions_using_lograge_as_default/)
- url: https://www.reddit.com/r/rails/comments/i43zyh/how_to_log_exceptions_using_lograge_as_default/
---
How to log exceptions using Lograge as default logger.

Basically, I am using Lograge as the default Rails logger that reduces the bulk of rails logs to formatted structure. But exceptions are not included by default. In the lograge configuration I included the following,

`opts[:exception] = event.payload[:exception] if event.payload[:exception]`  
`opts[:exceptionObject] = event.payload[:exception_object] if event.payload[:exception_object]`

This seem to pring the first like in code exceptions. It also does STDOUT and so i have the exceotions being logged normally with stacktrace. This does not include routing exceptions.

Is there a way I can get the exceptions logged as part of lograge output, with appropriate status, backtrace and not duplicate the exception once again?
