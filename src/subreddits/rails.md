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
## [2][Adding Snippets for easier rails coding](https://www.reddit.com/r/rails/comments/i42vyt/adding_snippets_for_easier_rails_coding/)
- url: https://www.reddit.com/r/rails/comments/i42vyt/adding_snippets_for_easier_rails_coding/
---
Hi just a quick question can someone tell me how i create shortcuts for  the below rails tags in VS code i was able to get emmet working its just the below tags wont work. Any guide how to set this up in rails to save time? 

    &lt;%= %&gt;

or 

    &lt;% %&gt;
## [3][Does serverless rails work well and is it worth it?](https://www.reddit.com/r/rails/comments/i3v2lr/does_serverless_rails_work_well_and_is_it_worth_it/)
- url: https://www.reddit.com/r/rails/comments/i3v2lr/does_serverless_rails_work_well_and_is_it_worth_it/
---
What are the pros and cons for doing so with Rails?
## [4][How to log exceptions using Lograge as default logger](https://www.reddit.com/r/rails/comments/i43zyh/how_to_log_exceptions_using_lograge_as_default/)
- url: https://www.reddit.com/r/rails/comments/i43zyh/how_to_log_exceptions_using_lograge_as_default/
---
How to log exceptions using Lograge as default logger.

Basically, I am using Lograge as the default Rails logger that reduces the bulk of rails logs to formatted structure. But exceptions are not included by default. In the lograge configuration I included the following,

`opts[:exception] = event.payload[:exception] if event.payload[:exception]`  
`opts[:exceptionObject] = event.payload[:exception_object] if event.payload[:exception_object]`

This seem to pring the first like in code exceptions. It also does STDOUT and so i have the exceotions being logged normally with stacktrace. This does not include routing exceptions.

Is there a way I can get the exceptions logged as part of lograge output, with appropriate status, backtrace and not duplicate the exception once again?
## [5][Working on a React project and hit a snag.](https://www.reddit.com/r/rails/comments/i3wsv4/working_on_a_react_project_and_hit_a_snag/)
- url: https://www.reddit.com/r/rails/comments/i3wsv4/working_on_a_react_project_and_hit_a_snag/
---
ActionController::UnknownFormat (WelcomeController#index is missing a template for this request format and variant.

I tried just doing respond\_to :html in the Controller#Index, but it's being weird. Specifically, I'm trying to load a react component and it's not happening. I believe I have all my routes set correctly, so I'm not sure what this error message means exactly. Can anyone clarify?
## [6][Rails 5.2 Credentials API](https://www.reddit.com/r/rails/comments/i3sz7n/rails_52_credentials_api/)
- url: https://www.reddit.com/r/rails/comments/i3sz7n/rails_52_credentials_api/
---
Does anyone know if it's possible to only use the encrypted configs for production/staging/uat?

i'd like to use plain text config yamls for dev and test environments but it appears to be an all or nothing thing w/ the new credentials api.
## [7][How do you prevent users editing &amp; deleting what isn't theirs?](https://www.reddit.com/r/rails/comments/i3xk7o/how_do_you_prevent_users_editing_deleting_what/)
- url: https://www.reddit.com/r/rails/comments/i3xk7o/how_do_you_prevent_users_editing_deleting_what/
---
I heard of Pundit, but I'm new to Rails and don't get all the docs. And not sure its applicable. I want it that sellers cannot edit or delete other sellers' items.
## [8][Pros and cons of Rails with Docker?](https://www.reddit.com/r/rails/comments/i3xxjf/pros_and_cons_of_rails_with_docker/)
- url: https://www.reddit.com/r/rails/comments/i3xxjf/pros_and_cons_of_rails_with_docker/
---
I mainly develop by myself and am wondering if Docker would add too much complexity for me.
## [9][Mina Deploy issue](https://www.reddit.com/r/rails/comments/i3mzmu/mina_deploy_issue/)
- url: https://www.reddit.com/r/rails/comments/i3mzmu/mina_deploy_issue/
---
I have been fighting Capistrano for weeks trying to get deployment to work, so I decide to try switching to Mina. I am running into the following issue but I am not sure where the issue is. Any help would be greatly appreciated. 

Mina Output:

`Bundle complete! 22 Gemfile dependencies, 70 gems now installed.`

`Gems in the groups development and test were not installed.`

`Bundled gems are installed into \`./vendor/bundle\``

`Files /var/www/site/current/db/migrate/20200518185620_create_table1.rb and ./db/migrate/20200518185620_create_table1.rb differ`

`Files /var/www/site/current/db/migrate/20200518185712_create_table2.rb and ./db/migrate/20200518185712_create_table2.rb differ`

`Files /var/www/site/current/db/migrate/20200710220508_create_table3.rb and ./db/migrate/20200710220508_create_table3.rb differ`

`Files /var/www/site/current/db/migrate/20200710220643_create_table4.rb and ./db/migrate/20200710220643_create_table4.rb differ`

`Files /var/www/site/current/db/migrate/20200714180204_add_forgotten_to_table1.rb and ./db/migrate/20200714180204_add_forgotten_to_table1.rb differ`

`Files /var/www/site/current/db/migrate/20200714184239_add_image_to_table4.rb and ./db/migrate/20200714184239_add_image_to_table4.rb differ`

`Files /var/www/site/current/db/migrate/20200715171629_change_field_on_table1.rb and ./db/migrate/20200715171629_change_field_on_table1.rb differ`

`Files /var/www/site/current/db/migrate/20200721190108_devise_create_users.rb and ./db/migrate/20200721190108_devise_create_users.rb differ`

`-----&gt; Migrating database`

`Traceback (most recent call last):`

`3: from /var/www/site/tmp/build-159655897320394/vendor/bundle/ruby/2.6.0/bin/ruby_executable_hooks:24:in \`&lt;main&gt;'`

`2: from /var/www/site/tmp/build-159655897320394/vendor/bundle/ruby/2.6.0/bin/ruby_executable_hooks:24:in \`eval'`

`1: from /usr/local/rvm/gems/ruby-2.6.3/bin/rake:23:in \`&lt;main&gt;'`

`/usr/local/rvm/gems/ruby-2.6.3/bin/rake:23:in \`load': cannot load such file -- /usr/local/rvm/rubies/ruby-2.6.3/lib/ruby/gems/2.6.0/specifications/default/exe/rake (LoadError)`

 `!     ERROR: Deploy failed.`

`-----&gt; Cleaning up build`

`Unlinking current`

`OK`

`Connection to server closed.`

 `!     Run Error`

&amp;#x200B;

My Mina deploy.rb 

`require 'mina/bundler'`

`require 'mina/rails'`

`require 'mina/git'`

`require 'mina/rvm'`

&amp;#x200B;

`set :application_name, 'site'`

`set :domain, 'server'`

`set :user, 'deployer'`

`set :deploy_to, "/var/www/site"`

`set :repository,` [`"git@github.com`](mailto:"git@github.com)`:git.git"`

`set :branch, 'master'`

`set :rvm_use_path, '/etc/profile.d/rvm.sh'`

&amp;#x200B;

`set :shared_files, fetch(:shared_files, []).push('config/database.yml', 'config/secrets.yml')`

&amp;#x200B;

`task :remote_environment do`

  `ruby_version =` [`File.read`](https://File.read)`('.ruby-version').strip`

  `raise "Couldn't determine Ruby version: Do you have a file .ruby-version in your project root?" if ruby_version.empty?`

&amp;#x200B;

  `invoke :'rvm:use', 'ruby-2.6.3'`

`end`

&amp;#x200B;

`desc "Deploys the current version to the server."`

`task :deploy do`

  

  `deploy do`



`invoke :'git:clone'`

`invoke :'deploy:link_shared_paths'`

`invoke :'bundle:install'`

`invoke :'rails:db_migrate'`

`invoke :'rails:assets_precompile'`

`invoke :'deploy:cleanup'`

&amp;#x200B;

`on :launch do`

`end`

  `end`

&amp;#x200B;

`end`

&amp;#x200B;

`namespace :passenger do`

  `desc "Restart Passenger"`

  `task :restart do`

`queue %{`

`echo "-----&gt; Restarting passenger"`

`cd #{deploy_to}/current`

`#{echo_cmd %[mkdir -p tmp]}`

`#{echo_cmd %[touch tmp/restart.txt]}`

`}`

  `end`

`end`
## [10][Sending Thousands of Request at the Same Time?](https://www.reddit.com/r/rails/comments/i3iyu5/sending_thousands_of_request_at_the_same_time/)
- url: https://www.reddit.com/r/rails/comments/i3iyu5/sending_thousands_of_request_at_the_same_time/
---
I need to develop an endpoint that will send (possibly) hundreds/thousands of request, process it then return/render the json to user/client.

I've tried using thread pool with the size of 5, but it took forever, but when I tried increasing the size to the number of request, it threw \`ThreadError: can't create Thread: Resource temporarily unavailable\` exception.

I don't think I can use background job/worker for this because I should return the result.

So what should I do?

I was thinking that I should wrap the process in 20sec timeout so it doesn't reach rails 30sec limit, and if it's still not finished in 20sec, it will return the unfinished result. It goes like this [https://i.imgur.com/FolcJV6.png](https://i.imgur.com/FolcJV6.png)

But it's still not working, the process still keep going even after it timeout.
## [11][Running tests in containers with docker-compose](https://www.reddit.com/r/rails/comments/i3fzbz/running_tests_in_containers_with_dockercompose/)
- url: https://www.reddit.com/r/rails/comments/i3fzbz/running_tests_in_containers_with_dockercompose/
---
The main advantages of this way are to have an independent environment for the tests running and to reduce the complexity of the test environment setup.

What we want to achieve:

* Running the tests should be easy.
* Test runs should be isolated and repeatable.
* Test environment should be as close to the production environment as possible.

How to setup and use docker-compose for Ruby on Rails tests you can find in the article: [https://jtway.co/running-tests-in-containers-with-docker-compose-97480726c1e3](https://jtway.co/running-tests-in-containers-with-docker-compose-97480726c1e3)
