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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hrnm2o/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hrnm2o/personal_projects_show_off_your_own_project_andor/
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
## [3][Deployment problem with HTTP Active Record adapter](https://www.reddit.com/r/rails/comments/htbgz9/deployment_problem_with_http_active_record_adapter/)
- url: https://www.reddit.com/r/rails/comments/htbgz9/deployment_problem_with_http_active_record_adapter/
---
I'm trying to install a Rails app (Spree E-commerce) on a Digital Ocean droplet using the [Tomo CLI](https://github.com/mattbrictson/tomo) tool. You run it on your local machine, give it the login credentials for the server, and it will deploy your Rails app for you. At least in theory.

I've had problem after problem with Tomo, most of which I've managed to fix. However, I'm at a loss for the current problem, which Tomo encounters when it tries to create the database with `bundle exec rake db:create`. I tried creating the database first, as Tomo seems to skip steps if they are already completed on the server, but that didn't fix anything.

Here is the error message provided by Tomo. Googling has brought me zilch so far. Any help would be great!

    • rails:db_create
    cd /tmp/tomo/20200718051237 &amp;&amp; bundle exec rake db:version
    cd /tmp/tomo/20200718051237 &amp;&amp; bundle exec rake db:create
    DEPRECATION WARNING: Including LoggerSilence is deprecated and will be removed in Rails 6.1. Please use `ActiveSupport::LoggerSilence` instead (called from block (2 levels) in require at /home/deployer/.rbenv/versions/2.6.5/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/lib/bundler/runtime.rb:74)
    rake aborted!
    LoadError: Could not load the 'http' Active Record adapter. Ensure that the adapter is spelled correctly in config/database.yml and that you've added the necessary adapter gem to your Gemfile.
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:34:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/polyglot-0.3.5/lib/polyglot.rb:65:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/zeitwerk-2.3.0/lib/zeitwerk/kernel.rb:23:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `block in require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:291:in `load_dependency'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/connection_adapters/connection_specification.rb:169:in `spec'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/connection_adapters/abstract/connection_pool.rb:1052:in `establish_connection'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/connection_handling.rb:51:in `establish_connection'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/railtie.rb:201:in `block (2 levels) in &lt;class:Railtie&gt;'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:71:in `class_eval'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:71:in `block in execute_hook'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:61:in `with_execution_control'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:66:in `execute_hook'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:43:in `block in on_load'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:42:in `each'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:42:in `on_load'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/railtie.rb:198:in `block in &lt;class:Railtie&gt;'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:32:in `instance_exec'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:32:in `run'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:61:in `block in run_initializers'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:60:in `run_initializers'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/application.rb:363:in `initialize!'
    /tmp/tomo/20200718051237/config/environment.rb:5:in `&lt;main&gt;'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:23:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:23:in `block in require_with_bootsnap_lfi'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/loaded_features_index.rb:92:in `register'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:22:in `require_with_bootsnap_lfi'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:31:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/polyglot-0.3.5/lib/polyglot.rb:65:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/zeitwerk-2.3.0/lib/zeitwerk/kernel.rb:23:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `block in require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:291:in `load_dependency'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/application.rb:339:in `require_environment!'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/application.rb:523:in `block in run_tasks_blocks'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/rake-13.0.1/exe/rake:27:in `&lt;top (required)&gt;'
    /home/deployer/.rbenv/versions/2.6.5/bin/bundle:23:in `load'
    /home/deployer/.rbenv/versions/2.6.5/bin/bundle:23:in `&lt;main&gt;'
    
    Caused by:
    LoadError: cannot load such file -- active_record/connection_adapters/http_adapter
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:34:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/polyglot-0.3.5/lib/polyglot.rb:65:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/zeitwerk-2.3.0/lib/zeitwerk/kernel.rb:23:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `block in require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:291:in `load_dependency'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/connection_adapters/connection_specification.rb:169:in `spec'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/connection_adapters/abstract/connection_pool.rb:1052:in `establish_connection'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/connection_handling.rb:51:in `establish_connection'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/railtie.rb:201:in `block (2 levels) in &lt;class:Railtie&gt;'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:71:in `class_eval'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:71:in `block in execute_hook'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:61:in `with_execution_control'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:66:in `execute_hook'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:43:in `block in on_load'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:42:in `each'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:42:in `on_load'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/railtie.rb:198:in `block in &lt;class:Railtie&gt;'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:32:in `instance_exec'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:32:in `run'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:61:in `block in run_initializers'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:60:in `run_initializers'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/application.rb:363:in `initialize!'
    /tmp/tomo/20200718051237/config/environment.rb:5:in `&lt;main&gt;'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:23:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:23:in `block in require_with_bootsnap_lfi'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/loaded_features_index.rb:92:in `register'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:22:in `require_with_bootsnap_lfi'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:31:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/polyglot-0.3.5/lib/polyglot.rb:65:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/zeitwerk-2.3.0/lib/zeitwerk/kernel.rb:23:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `block in require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:291:in `load_dependency'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `require'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/application.rb:339:in `require_environment!'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/application.rb:523:in `block in run_tasks_blocks'
    /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/rake-13.0.1/exe/rake:27:in `&lt;top (required)&gt;'
    /home/deployer/.rbenv/versions/2.6.5/bin/bundle:23:in `load'
    /home/deployer/.rbenv/versions/2.6.5/bin/bundle:23:in `&lt;main&gt;'
    Tasks: TOP =&gt; db:create =&gt; db:load_config =&gt; environment
    (See full trace by running task with --trace)
    
      ERROR: The following script failed on deployer@134.209.110.34 (exit status 1).
    
        cd /tmp/tomo/20200718051237 &amp;&amp; bundle exec rake db:create
    
      You can manually re-execute the script via SSH as follows:
    
        ssh -o LogLevel\=ERROR -A -o ConnectTimeout\=5 -o StrictHostKeyChecking\=accept-new -o ControlMaster\=auto -o ControlPath\=/var/folders/b1/mqpmy8jj4h74xl88g_wdxr3h0000gn/T/tomo_ssh_91c26631d427a22d -o ControlPersist\=30s -o PasswordAuthentication\=no deployer@134.209.110.34 -- cd\ /tmp/tomo/20200718051237\ \&amp;\&amp;\ bundle\ exec\ rake\ db:create
    
      For more troubleshooting info, run tomo again using the --debug option.
    
      DEPRECATION WARNING: Including LoggerSilence is deprecated and will be removed in Rails 6.1. Please use `ActiveSupport::LoggerSilence` instead (called from block (2 levels) in require at /home/deployer/.rbenv/versions/2.6.5/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/lib/bundler/runtime.rb:74)
      rake aborted!
      LoadError: Could not load the 'http' Active Record adapter. Ensure that the adapter is spelled correctly in config/database.yml and that you've added the necessary adapter gem to your Gemfile.
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:34:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/polyglot-0.3.5/lib/polyglot.rb:65:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/zeitwerk-2.3.0/lib/zeitwerk/kernel.rb:23:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `block in require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:291:in `load_dependency'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/connection_adapters/connection_specification.rb:169:in `spec'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/connection_adapters/abstract/connection_pool.rb:1052:in `establish_connection'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/connection_handling.rb:51:in `establish_connection'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/railtie.rb:201:in `block (2 levels) in &lt;class:Railtie&gt;'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:71:in `class_eval'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:71:in `block in execute_hook'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:61:in `with_execution_control'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:66:in `execute_hook'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:43:in `block in on_load'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:42:in `each'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:42:in `on_load'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/railtie.rb:198:in `block in &lt;class:Railtie&gt;'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:32:in `instance_exec'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:32:in `run'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:61:in `block in run_initializers'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:60:in `run_initializers'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/application.rb:363:in `initialize!'
      /tmp/tomo/20200718051237/config/environment.rb:5:in `&lt;main&gt;'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:23:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:23:in `block in require_with_bootsnap_lfi'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/loaded_features_index.rb:92:in `register'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:22:in `require_with_bootsnap_lfi'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:31:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/polyglot-0.3.5/lib/polyglot.rb:65:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/zeitwerk-2.3.0/lib/zeitwerk/kernel.rb:23:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `block in require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:291:in `load_dependency'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/application.rb:339:in `require_environment!'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/application.rb:523:in `block in run_tasks_blocks'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/rake-13.0.1/exe/rake:27:in `&lt;top (required)&gt;'
      /home/deployer/.rbenv/versions/2.6.5/bin/bundle:23:in `load'
      /home/deployer/.rbenv/versions/2.6.5/bin/bundle:23:in `&lt;main&gt;'
    
      Caused by:
      LoadError: cannot load such file -- active_record/connection_adapters/http_adapter
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:34:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/polyglot-0.3.5/lib/polyglot.rb:65:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/zeitwerk-2.3.0/lib/zeitwerk/kernel.rb:23:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `block in require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:291:in `load_dependency'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/connection_adapters/connection_specification.rb:169:in `spec'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/connection_adapters/abstract/connection_pool.rb:1052:in `establish_connection'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/connection_handling.rb:51:in `establish_connection'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/railtie.rb:201:in `block (2 levels) in &lt;class:Railtie&gt;'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:71:in `class_eval'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:71:in `block in execute_hook'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:61:in `with_execution_control'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:66:in `execute_hook'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:43:in `block in on_load'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:42:in `each'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/lazy_load_hooks.rb:42:in `on_load'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activerecord-6.0.3.2/lib/active_record/railtie.rb:198:in `block in &lt;class:Railtie&gt;'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:32:in `instance_exec'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:32:in `run'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:61:in `block in run_initializers'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/initializable.rb:60:in `run_initializers'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/application.rb:363:in `initialize!'
      /tmp/tomo/20200718051237/config/environment.rb:5:in `&lt;main&gt;'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:23:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:23:in `block in require_with_bootsnap_lfi'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/loaded_features_index.rb:92:in `register'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:22:in `require_with_bootsnap_lfi'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:31:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/polyglot-0.3.5/lib/polyglot.rb:65:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/zeitwerk-2.3.0/lib/zeitwerk/kernel.rb:23:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `block in require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:291:in `load_dependency'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/activesupport-6.0.3.2/lib/active_support/dependencies.rb:324:in `require'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/application.rb:339:in `require_environment!'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/railties-6.0.3.2/lib/rails/application.rb:523:in `block in run_tasks_blocks'
      /var/www/ncode-bkk-spree/shared/bundle/ruby/2.6.0/gems/rake-13.0.1/exe/rake:27:in `&lt;top (required)&gt;'
      /home/deployer/.rbenv/versions/2.6.5/bin/bundle:23:in `load'
      /home/deployer/.rbenv/versions/2.6.5/bin/bundle:23:in `&lt;main&gt;'
      Tasks: TOP =&gt; db:create =&gt; db:load_config =&gt; environment
      (See full trace by running task with --trace)
## [4][A little help from our friends](https://www.reddit.com/r/rails/comments/ht5g1c/a_little_help_from_our_friends/)
- url: https://www.reddit.com/r/rails/comments/ht5g1c/a_little_help_from_our_friends/
---
Would love if anyone can add what they have resource wise. I've have some solid feedback and help thus far. I no longer have to have 100 bookmarks or google something. 

Cheers!

Heres the repo  [https://github.com/tylertomlinson/crucial\_resources](https://github.com/tylertomlinson/crucial_resources)
## [5][Looking for opinions on Rails with Elasticsearch gems - does Elasticsearch work well for your application? Was it easy to integrate? Are there good alternatives?](https://www.reddit.com/r/rails/comments/hszo1o/looking_for_opinions_on_rails_with_elasticsearch/)
- url: https://www.reddit.com/r/rails/comments/hszo1o/looking_for_opinions_on_rails_with_elasticsearch/
---
I'm looking at elasticsearch-model gem and the elasticsearch-rails gem. I'm also curious if you had issues with elasticsearch? Any thoughts are appreciated!
## [6][Rspec undefined method ... for ?](https://www.reddit.com/r/rails/comments/ht5sth/rspec_undefined_method_for/)
- url: https://www.reddit.com/r/rails/comments/ht5sth/rspec_undefined_method_for/
---
Hi everyone, So i try to run my test but i've got this failure message that i just don't understand. 

         Failure/Error: before { post '/api/v0/posts', params: valid_attributes }
         
         NoMethodError:
           undefined method `set_encoding' for #&lt;Post:0x00005595c63e7068&gt;

my valid params look like this:

        let(:valid_attributes) { create :post, title: 'Learn Elm', content:'azertyuop', created_by:'User', entry: 4, category_id: category.id, rdv: DateTime.now + 1.week, tag1:'ayo', tag2:'tag2', tag3: '', user_id: user.id }

And this is my spec:

            context 'when the request is valid' do
                before { post '/api/v0/posts', params: valid_attributes }
    
                it 'creates a post' do
                    expect(JSON.parse(response.body)).to eq('Learn Elm')
                end

I don't have any clue so if you have any tips Post it, thanks.
## [7][Cryptic error. Any clue?](https://www.reddit.com/r/rails/comments/hsubkv/cryptic_error_any_clue/)
- url: https://www.reddit.com/r/rails/comments/hsubkv/cryptic_error_any_clue/
---
Hi Guys.

So far I have been using capistrano for my production app, and all was good, but trying to set up a staging environment, I get an error in the last moment and I don't know how to interpret it well.  


Can you give me any help?

&amp;#x200B;

https://preview.redd.it/t045d3jvoeb51.png?width=1666&amp;format=png&amp;auto=webp&amp;s=04de65012fc28e162b404dfcc78bfa2cfddfc95f

https://preview.redd.it/lg8z25jvoeb51.png?width=1666&amp;format=png&amp;auto=webp&amp;s=aaffabffcef25030478a08d0490178e86bd6e505
## [8][How to use a Transaction Script(aka Service Objects) in Ruby on Rails. Simple example](https://www.reddit.com/r/rails/comments/hssf51/how_to_use_a_transaction_scriptaka_service/)
- url: https://www.reddit.com/r/rails/comments/hssf51/how_to_use_a_transaction_scriptaka_service/
---
The logic of small applications can be present as a series of transactions. Using the Transaction Scripts pattern, we get an application that is easier to maintain, to cover with tests, and to scale.

In the [tutorial](https://jtway.co/how-to-use-a-transaction-script-aka-service-objects-in-ruby-on-rails-simple-example-161b7e228942) we will develop an [application](https://github.com/dgorodnichy/transaction-script-example) that has Post, User, and Like models. Users should be able to like posts. The first version of the controller will contain extra code, which we will extract into a separate Transaction Script. We also describe when we need to use the Transaction Scripts and the pros of the transaction script usage.  


Full tutorial: [How to use a Transaction Script (aka Service Objects) in Ruby on Rails. Simple example](https://jtway.co/how-to-use-a-transaction-script-aka-service-objects-in-ruby-on-rails-simple-example-161b7e228942)
## [9][Logic/db help for a New Programmer](https://www.reddit.com/r/rails/comments/hsxl3p/logicdb_help_for_a_new_programmer/)
- url: https://www.reddit.com/r/rails/comments/hsxl3p/logicdb_help_for_a_new_programmer/
---
Making a simple community website for a niche community of gamers. 

Hung up on how to associate and display
Users to Games and platforms. 

Do I store a game ID and platform ID in the user profile, or store the platform specific username in the game card. 


User has many games
Game has many users

Do I just store it twice?
User gamelist
Game userlist 

When UserA adds PS-Fortnite to their list, there is also an entry made for PS-Fortnite registering UserA?

Does it need to be one or the other, or both?
## [10][Writing a good Rails app CSP, and how to build a report_uri](https://www.reddit.com/r/rails/comments/hspydx/writing_a_good_rails_app_csp_and_how_to_build_a/)
- url: https://www.reddit.com/r/rails/comments/hspydx/writing_a_good_rails_app_csp_and_how_to_build_a/
---
The [official documentation](https://edgeguides.rubyonrails.org/security.html#content-security-policy) suggests starting with a global policy, and making changes as needed. In trying to understand this, I've made some very minor changes:

    Rails.application.config.content_security_policy do |policy|
      policy.default_src :self, :https
      config.font_src    :self, :https, 'fonts.googleapis.com'
      policy.img_src     :self, :https, :data, 'blahblah.s3-us-east-1.amazonaws.com'
      policy.object_src  :none
      config.script_src  :self, :https, 'stripe.com'
      policy.style_src   :self, :https
    
      # Specify URI for violation reports
      policy.report_uri "/csp-violation-report-endpoint"
    end

Two questions about this.

1. Should excluding the `blahblah.s3-us-east-1.amazonaws.com` from `policy.img_src` effectively prevent those resources from loading? I would think so, but the resource still loads.
2. Same for `fonts_src.` Should excluding `fonts.googleapis.com` prevent that font from loading? I get a console warning that says it's not loading due to a Content Security Policy violation, but the font still seems to be used when checking the styles in Inspector.
3. Exactly what is this `report_uri`? I can't seem to find anything that explains it. Presumably it needs a `POST` route, a controller, and a view accessible only to an admin?
## [11][Is innovation needed anyway ?](https://www.reddit.com/r/rails/comments/hsskui/is_innovation_needed_anyway/)
- url: https://www.reddit.com/r/rails/comments/hsskui/is_innovation_needed_anyway/
---
Here is a tweet quote from excellent Chris Oliver (GoRails) "If we want the Ruby community to grow, we need to get back to innovating and talking about it." My feeling is : is there any need to innovate anyway ?   
Even the JS hype of the last years concern only a few pages that need more interactions... only small parts of a standard business app. My feeling is that innovation do not concern anymore the vast majority of every day problems. Any thought ?
## [12][Hi everyone.](https://www.reddit.com/r/rails/comments/hsolu9/hi_everyone/)
- url: https://www.reddit.com/r/rails/comments/hsolu9/hi_everyone/
---
HI everyone, i hope that everyone is doing good through these times we are experiencing and everyone is safe, i was just wondering if ruby and rails is still the best framework for prototyping? coming from python :). 

any advice would be appreciated. Thank you
