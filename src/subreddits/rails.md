# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/
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
## [3][I applied for a RoR job!](https://www.reddit.com/r/rails/comments/iiddg4/i_applied_for_a_ror_job/)
- url: https://www.reddit.com/r/rails/comments/iiddg4/i_applied_for_a_ror_job/
---
After leeching for almost 2 years, while learning and building own projects I applied for a REAL RoR job. I felt confident enough to do so. This is one of those rare companies who had a decent job description without requiring 5+ years experience with none existing language 

Thanks u/rails community for your alwasy friendly help

Wish me luck!
## [4][Anyone taking Ruby on Rails courses?](https://www.reddit.com/r/rails/comments/iisg11/anyone_taking_ruby_on_rails_courses/)
- url: https://www.reddit.com/r/rails/comments/iisg11/anyone_taking_ruby_on_rails_courses/
---
Hey guys, anyone taking [*Ruby on Rails Courses*](https://www.google.com/amp/s/onlinecoursesgalore.com/ruby-rails-udemy/amp/) on Udemy?
## [5][Issue with GCS Private Key using dotenv](https://www.reddit.com/r/rails/comments/iiqtkg/issue_with_gcs_private_key_using_dotenv/)
- url: https://www.reddit.com/r/rails/comments/iiqtkg/issue_with_gcs_private_key_using_dotenv/
---
I am trying to use active storage with GCS.  It all worked when I used keyfile.json, but when I wanted to make it work using dotenv-rails I am getting error with private key entry.

&amp;#x200B;

storage.yml

    ...
    google:
      service: GCS 
      project: &lt;%= ENV['GCS_PROJECT'] %&gt; 
      bucket: &lt;%= ENV['GCS_BUCKET'] %&gt; 
      credentials: type: "service_account" 
        project_id: &lt;%= ENV['GCS_PROJECT_ID'] %&gt; 
        private_key_id: &lt;%= ENV['GCS_PRIVATE_KEY_ID'] %&gt; 
        private_key: &lt;%= ENV['GCS_PRIVATE_KEY'] %&gt; 
        client_email: &lt;%= ENV['GCS_CLIENT_EMAIL'] %&gt; 
        client_id: &lt;%= ENV['GCS_CLIENT_ID'] %&gt; 
        auth_uri: "https://accounts.google.com/o/oauth2/auth" 
        token_uri: "https://accounts.google.com/o/oauth2/token"    
        auth_provider_x509_cert_url: "https://www.googleapis.com/oauth2/v1/certs"  
        client_x509_cert_url: &lt;%= ENV['GCS_CLIENT_X509_CERT_URL'] %&gt; 
    ...

.env

    ...
    GCS_PRIVATE_KEY="-----BEGIN PRIVATE KEY-----\n&lt; ACTUAL SECRET KEY &gt;\n-----END PRIVATE KEY-----\n"
    ...

&amp;#x200B;

error I get

    YAML syntax error occurred while parsing /Users/admin/Desktop/ruby/rails/launchzilla/config/storage.yml. Please note that YAML must be consistently indented using spaces. Tabs are not allowed. Error: (&lt;unknown&gt;): could not find expected ':' while scanning a simple key at line 29 column 1

I am pretty sure its the Private key part, because I have tested all the scenarios, and when I remove private key part the yml seem to parse corectly, but of course I get error when trying to upload, since GCS setting are incorrect then. Any ideas?
## [6][Why Rails and Not...](https://www.reddit.com/r/rails/comments/iia406/why_rails_and_not/)
- url: https://www.reddit.com/r/rails/comments/iia406/why_rails_and_not/
---
Hey all, I work for a company that's considering a major redo of their e-commerce platform. We're considering a bunch of frameworks for this: Rails, Laravel, Django, and Node. So I thought I'd ask the community here what would be some good reasons for Rails over these others? Thanks so much in advance. :)
## [7][What’s the best formula to calculate unicorn workers?](https://www.reddit.com/r/rails/comments/iijesu/whats_the_best_formula_to_calculate_unicorn/)
- url: https://www.reddit.com/r/rails/comments/iijesu/whats_the_best_formula_to_calculate_unicorn/
---
I am struggling to achieve the best number of unicorn for my Rails projects.
My apps are sub-utilizing CPU and RAM.
## [8][Newb trying to fix a CORS issue in WEBrick](https://www.reddit.com/r/rails/comments/iin972/newb_trying_to_fix_a_cors_issue_in_webrick/)
- url: https://www.reddit.com/r/rails/comments/iin972/newb_trying_to_fix_a_cors_issue_in_webrick/
---
I am getting a cors issue with DELETE, so I am trying to fix it, and can't understand how to do it with WEBrick. It looks like there is an application.rb file with environment development.rb. I am dealing with a legacy system I don't know about...

    'http://localhost:3006' has been blocked by CORS policy: Method DELETE is not allowed by Access-Control-Allow-Methods in preflight response.
    
    
    System:
      [2020-08-28 22:14:26] INFO  WEBrick 1.3.1 [2020-08-28 22:14:26] INFO 
      ruby 2.3.8 (2018-10-18) [x86_64-darwin19]

My development.rb looks like:

Where the heck do I add DELETE/PUT header info????

    *********::Application.configure do
      config.cache_classes = false
    
      config.whiny_nils = true
    
      config.logger = SkipResourcesLogger.new(File.join(Rails.root, 'log', 'development.log'))
    
      config.lograge.formatter = Lograge::Formatters::Logstash.new
      config.lograge.enabled = true
      config.lograge.custom_options = lambda do |event| stacktrace = {}
      if event.payload[:exception] &amp;&amp; strace = event.payload[:stacktrace]
          quoted_stacktrace = %('#{Array(strace).reject{|l| l.include?('/gems/') || 
      l.include?('/instrumenter')}.join("\n\t")}') # depending on if I'm logging 
      as logstash or not
          stacktrace = {stacktrace: quoted_stacktrace}
      end
        exceptions = %w(controller action format id)
        {
          id: event.payload[:ip],
          user_id: event.payload[:user_id],
          params: event.payload[:params].except(*exceptions)
        }.merge(stacktrace)
     end
    
      config.consider_all_requests_local       = true
      config.action_controller.perform_caching = true
    
      config.cache_store = :dalli_store, 'localhost:11211', { :compress =&gt; true, :value_max_bytes =&gt; 1024*1024*32 }
    
      config.active_support.deprecation = :log
    
      config.action_dispatch.best_standards_support = :builtin
    
      config.active_record.mass_assignment_sanitizer = :strict
    
      config.active_record.auto_explain_threshold_in_seconds = 0.5
    
      config.assets.compress = false
      config.assets.debug = true
    
       unless defined?(Rails::Console)
       end
       config.dependency_loading = true if $rails_rake_task
    
       config.assets.logger = false
    
       config.action_mailer.default_url_options = { host: 'localhost' }
       Rails.application.routes.default_url_options = { host: 'localhost' }
    
       config.after_initialize do
       Bullet.enable = true
       Bullet.alert = true
       Bullet.rails_logger = true
       Bullet.console = true
       Bullet.add_footer = true
      end
    
       config.action_controller.asset_host = "http://localhost"
    
     end
## [9][[HELP] What is the "rails way" of eager loading multiple joins with 'includes'](https://www.reddit.com/r/rails/comments/iighai/help_what_is_the_rails_way_of_eager_loading/)
- url: https://www.reddit.com/r/rails/comments/iighai/help_what_is_the_rails_way_of_eager_loading/
---
I'm working on a simple small 6 project to manage a competition. 

I have a **team**, **user** and **trip** models where the team has_many users and user has_many trips. Trip has a mileage field. 

I need the controller to get all the teams and their status. The status for now is the number of trips and total mileage per team.

my view iterates the @teams and displays them in a table:

    &lt;% @teams.each do |team|%&gt;
      &lt;%= content_tag :tr, id: dom_id(team) do %&gt;
          ... add td tags ...
      &lt;% end %&gt;
    &lt;% end %&gt;

I had what I need doing something like the following but as you know it writes a separate query for each team. 

    @team.users.joins(:trips).sum(:distance_miles)

After some research it seemed like I needed eager loading with "includes". I have the following in my teams_controller:
    
    def index
        # @teams = Team.includes(users: :trips).group(:id).sum(:distance_miles)
        @teams = Team.left_joins(users: :trips).select('teams.name, SUM(trips.distance_miles) as num_miles, COUNT(trips.id) as num_trips').group(:id)
    end

So the statement that uses include doesn't get the team name so I came up with the left_join. I want to add pagination, filtering, and sorting soon and so I don't want to over complicate the initial query.

That gets me what I'm looking for but is that the rails way to do something like this?
## [10][How can I restore a PG dump from Heroku to AWS RDS instance?](https://www.reddit.com/r/rails/comments/iifcet/how_can_i_restore_a_pg_dump_from_heroku_to_aws/)
- url: https://www.reddit.com/r/rails/comments/iifcet/how_can_i_restore_a_pg_dump_from_heroku_to_aws/
---
Hey, I'm migrating my rails app from Heroku to AWS (after suffering at Heroku for so long smh). The app is deployed and is working on an empty database. However, I'm finding it hard to migrate the database snapshot to the RDS instance that's created with the Beanstalk Enviornment. How do I do that using the console?
Thanks in advance :)
## [11][Can you share some stats of your app?](https://www.reddit.com/r/rails/comments/iifa2x/can_you_share_some_stats_of_your_app/)
- url: https://www.reddit.com/r/rails/comments/iifa2x/can_you_share_some_stats_of_your_app/
---
I'm curious. What kind of app is it? How many req/sec does it handle? How many app instances, servers etc? Can you say something about the infrastructure required? Thanks!
## [12][Navbar background color change on scroll in Rails portfolio?](https://www.reddit.com/r/rails/comments/iiiy26/navbar_background_color_change_on_scroll_in_rails/)
- url: https://www.reddit.com/r/rails/comments/iiiy26/navbar_background_color_change_on_scroll_in_rails/
---
Hi guys! I'm creating a simple web dev portfolio in Rails. I've added a banner that has half of it being a black background with white text including the links in the navbar with a transparent background so my image doesn't get cut off by it. But once you start scrolling I'd like the navbar to have a black background at least for now due to the fact that sections below the banner will have lighter backgrounds. Like my "About me" section is a very light grey. I also won't want the white links to go over any text in other sections obviously.

I have done some Googling of course and found similar JavasScript code blocks that seem very simple. I have very little JS knowledge, but looking at one example on codepen, I just copied and pasted and replaced the selector names and colors with my own from my own CSS file. I'll put the CSS I have for the navbar, which is originally "header" for me. And then it looks like you just create a new selector with the same css except for the color so that the JS code can make the page switch to that CSS on scroll. I have put the JS code just in my "application.js" file and it's not working, so I'm not sure if I have the code in the right place, or if the code is even correct at all. Any help will be greatly appreciated! Also, don't worry. Once I have this completed, I'm definitely going to be learning JS ASAP as there's an optional section from my bootcamp to teach me :). If you need to see my repository for all the code [here it is.](https://github.com/Kyle-Williamson-Dev/My-Portfolio)

CSS and Javascript for color changing on scroll: 

    header {
      position: fixed;
      top: 0;
      left: 0;
      width: 100%;
      padding: 15px 100px;
      z-index: 1000;
      display: flex;
      background: transparent;
    }
    
    .nav-bg {
      position: fixed;
      top: 0;
      left: 0;
      width: 100%;
      padding: 15px 100px;
      z-index: 1000;
      display: flex;
      background: black;
    }
     
    $(document).ready(function(){
      $(window).scroll(function(){
      	var scroll = $(window).scrollTop();
    	  if (scroll &gt; 100) {
    	    $(".nav-bg").css("background" , "black");
    	  }
    
    	  else{
                $("header").css("background" , "transparent");  	
    	  }
      })
    })
