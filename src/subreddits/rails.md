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
## [3][[Help] Slow JSON Post -&gt; JSONB Column in Postgres](https://www.reddit.com/r/rails/comments/htsmbw/help_slow_json_post_jsonb_column_in_postgres/)
- url: https://www.reddit.com/r/rails/comments/htsmbw/help_slow_json_post_jsonb_column_in_postgres/
---
HI!

I've dabbled in rails off and on over the years and I'm back to an old project trying to squeeze some life out of it.

The long and skinny of it is this:

I have a client that is posting a json array of 30,000'ish elements to Rails, and then rails is putting that straight into a JSONB column in postgres. This is slower than I like. 10 full seconds or so on Heroku...

It's not nearly that slow to query it back out?

Curious if I'm missing something obvious about possibly skipping the json object creation on the post or something. Or is it possible that posting this many records at once will always be slow?

Thanks! Here is the create method on the controller.

    @songbook = current_user.songbooks.build(songbook_params)

    if @songbook.save
      flash[:notice] = "Successfully created songbook."

      respond_to do |format|
        format.html { redirect_to :action =&gt; "index" }
        format.json { render :json =&gt; @songbook }
      end
    else
      flash[:error] = @songbook.errors.full_messages.join(": ")

      respond_to do |format|
        format.html { render :action =&gt; "new" }
        format.json { render :json =&gt; @songbook.errors }
      end
    end
## [4][How to generate a persistent URL for ActiveStorage Blob so that can be included in a content page.](https://www.reddit.com/r/rails/comments/httsxe/how_to_generate_a_persistent_url_for/)
- url: https://www.reddit.com/r/rails/comments/httsxe/how_to_generate_a_persistent_url_for/
---
Hi everyone,

I create a content page where I let people use any of their already uploaded media, for that I need to generate an URL. If people chose any media the URL of that media (e.g. image, video, doc) will be then attached to the content page.

    def url
       rails_blob_path(resource, only_path: true)`
    end

This is how I generate url for a media. I searched the documentation and found that urls generated from active stroage are not permanent and having a permanent url generated from a media is not a good practice.

I've `config.active_storage_service` set to local. I use local file system to store uploaded media. I use rails 6.0.3.  
I would be really glad if someone can help point me towards the right solution.  


Thank You.  

## [5][Demonstrating turbolinks?](https://www.reddit.com/r/rails/comments/htu9n5/demonstrating_turbolinks/)
- url: https://www.reddit.com/r/rails/comments/htu9n5/demonstrating_turbolinks/
---
Curious if anyone has any ideas here... I'm doing a presentation on doing responsive rails apps without a JS framework (building an app and enhancing it with turbolinks, UJS, stimulus and stimulus reflex). 

Most of this is fairly easy to show off, but turbolinks is kind of invisible. I mean really the only noticeable thing about it is that page "loads" seem faster, but with covidpalooza I'll be doing this all over a screen share, that's going to be hard to demo. 

Does anyone have any ideas on something that would come across as showing the benefit of turbolinks? I was thinking some kind of timer that displays the load time from the browser perspective might be neat but I don't know how to pull that off...
## [6][Solution for revision functionality](https://www.reddit.com/r/rails/comments/htjloy/solution_for_revision_functionality/)
- url: https://www.reddit.com/r/rails/comments/htjloy/solution_for_revision_functionality/
---
Hi there,

Imagine that we have a model, let's call it Article. An Article is created by someone and it can also have associations to other models like Tag, etc.

Now, a couple of years later, we need to support what I will call the "revision" functionality. When someone revisions an Article can  pretty much change anything (properties, assoations, etc). However, that revision should not replace the original article right away, it should be stored so someone else can see the changes and approve or reject then. Only then, upon approval, should the revision replace the original Article.

The revision functionality is not that hard to implement upfront, but implementing it on a long running project is like changing the rules of the game at half time: it won't be pretty. I don't want to have a parallel/replic model called ArticleRevision, and any assoation to a revision should not mess with the logic I have now. 

I was wondering, do you know any gem/solution that can magically help us achieving this functionality?
## [7][How to register an user with SteamAPI and custom email, username etc.](https://www.reddit.com/r/rails/comments/hthahm/how_to_register_an_user_with_steamapi_and_custom/)
- url: https://www.reddit.com/r/rails/comments/hthahm/how_to_register_an_user_with_steamapi_and_custom/
---
Hello! I am using SteamAPI to register the user(with `omniauth-steam` gem). I only need Steam UID and the profile image from their API. I want the user to enter the email address and the username and some other fields manually. 

At this point, I need to note that I am using Rails 6 API and React. 

When user clicks on the "Sign in with Steam" button I redirect him to `localhost:3000/auth/steam` which goes to Steam's own login page. After the login is successful, Omniauth automatically redirects the user to `localhost:3000/auth/steam/callback`. Here I can create a new user with let's say `User.create(uid: auth.uid, image: auth.info['profile_image'])` with `auth` being the Steam user object that just logged in. But I don't know how to add the user's own email, username, birth date, etc. which I would ask for in the text fields on my React front-end.

Since I can't just call `POST /auth/steam/callback` I can't pass some data to it let's say from Axios or something and add other custom fields. 

If it helps, here is the `auth_callback` function that get's called by omniauth:
```ruby
class SteamController &lt; ApplicationController
  skip_before_action :verify_authenticity_token, only: :auth_callback

  def auth_callback
    # Get Steam user data
    auth = request.env['omniauth.auth']
    # Create new User
    user = User.new(steam_uid: auth['uid'], birthday_date: DateTime.now.to_date)
    if user.save
      render json: user, status: 201
    else
      render json: user.errors.full_messages, status: 400
    end
  end
end
```

As you can see for now I am creating the user with only Steam UID and birthday date which doesn't really matter right now.

I would really appreciate any help, tips, or resources to look into. Thanks!
## [8][Getting unique record Id's into iFrame form?](https://www.reddit.com/r/rails/comments/htn20o/getting_unique_record_ids_into_iframe_form/)
- url: https://www.reddit.com/r/rails/comments/htn20o/getting_unique_record_ids_into_iframe_form/
---
I have created an iframe that can be in other websites.

I am doing something like this:

    &lt;iframe id="unique_id" token="1235235" src="https://1234567890.ngrok.io/widgets/form" &gt;&lt;/iframe&gt;

the \`token\` will be a record of some sort so I can connect the form to a user or record so i know exactly what the form is related to when submitted.  Basically, I want to make sure i know which user the form is for so when it submits i can relate it correctly, etc.

**1st Attempt**

My first thought was to use nested resources for the iframe url like: /users/1/form/widget

Although, I am unsure how giving a url like that would benefit as none of the iframe url ID's come into the form submit.

**Second**

I then figured i could makea variable into the iframe element itself, as mentioned in the beginning, with "token" or any other name.

Although, I have an issue doing this with the javascript and jquery as i am unable to get it on load.  I also tried after form submit but that also isn't working.  It is odd because...

&amp;#x200B;

My issue is, seemingly, mainly in the javascript and jquery.  I am trying to get the form submit info so before the form submits, I grab the iframe token, create an input hidden field, add the token to the hidden field value and then submit.

&amp;#x200B;

When I use:

    $('iForm).submit(function() {
        console.log("A");
    });

I get nothing but the error ahead of time is: `$ is not defined`, which is understandable but how would i even get the attributes from the iframe?

I am able to within the browsers console.  

But if I do something like:

    document.createElement('input');
    input.setAttribute('type', 'hidden');
    ...;
    document.getElementById('iForm').appendChild(input);

It will create a hidden field and work.

&amp;#x200B;

How can I make this work?

&amp;#x200B;

Also, should I be doing this like that with the javascript or jquery, or is there a routing and/or rails way?
## [9][Deployment problem with HTTP Active Record adapter](https://www.reddit.com/r/rails/comments/htbgz9/deployment_problem_with_http_active_record_adapter/)
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
## [10][A little help from our friends](https://www.reddit.com/r/rails/comments/ht5g1c/a_little_help_from_our_friends/)
- url: https://www.reddit.com/r/rails/comments/ht5g1c/a_little_help_from_our_friends/
---
Would love if anyone can add what they have resource wise. I've have some solid feedback and help thus far. I no longer have to have 100 bookmarks or google something. 

Cheers!

Heres the repo  [https://github.com/tylertomlinson/crucial\_resources](https://github.com/tylertomlinson/crucial_resources)
## [11][Rspec undefined method ... for ?](https://www.reddit.com/r/rails/comments/ht5sth/rspec_undefined_method_for/)
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
## [12][Looking for opinions on Rails with Elasticsearch gems - does Elasticsearch work well for your application? Was it easy to integrate? Are there good alternatives?](https://www.reddit.com/r/rails/comments/hszo1o/looking_for_opinions_on_rails_with_elasticsearch/)
- url: https://www.reddit.com/r/rails/comments/hszo1o/looking_for_opinions_on_rails_with_elasticsearch/
---
I'm looking at elasticsearch-model gem and the elasticsearch-rails gem. I'm also curious if you had issues with elasticsearch? Any thoughts are appreciated!
