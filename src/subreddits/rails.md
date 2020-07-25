# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hrnm2o/personal_projects_show_off_your_own_project_andor/)
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/hwehh6/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/hwehh6/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][What's the 'Rails' way of writing clean code (and adhering to the single responsibility principle)?](https://www.reddit.com/r/rails/comments/hxl330/whats_the_rails_way_of_writing_clean_code_and/)
- url: https://www.reddit.com/r/rails/comments/hxl330/whats_the_rails_way_of_writing_clean_code_and/
---
I just finished the [lesson 1 of writing clean code by "Uncle Bob"](https://www.youtube.com/watch?v=7EmboKQH8lM), and my biggest takeaway was the single responsibility principle — that "a function does one thing only" — and "polite code": all lines in a function should be on the same abstraction level, and they should call functions on a lower abstraction level, going all the way down to implementation.

It's like newspaper: you read the heading to see what it's about; if you're interesting in learning more, you read the first paragraph; then the next; etc. It's being polite to the person trying to understand the codebase — they can take a look at the high-level function and understand its purpose without having to learn the unnecessary details in the process.

His way of avoiding thousands of little functions as a result of this was to put everything into its own class, and to find those sometimes-non-obvious classes. In a Java-like language, that makes sense, but should we do the same in Ruby on Rails?

I'm a beginner at Rails, but from what I can see it has a pretty strict structure ("convention over configuration") of MVC, where each "C" has several actions where the code goes. Should we split that code into other classes and compose everything inside the action? Or is there some other, better, more "Rails" way of doing things?
## [4][Ruby on Rails and IoT](https://www.reddit.com/r/rails/comments/hxhhti/ruby_on_rails_and_iot/)
- url: https://www.reddit.com/r/rails/comments/hxhhti/ruby_on_rails_and_iot/
---
Hello Everyone,

I'm new to ruby on rails and I'm trying to build a web app that lets me subscribe to a topic through mqtt and recieve a payload from a simulated node-red IoT device that send de data to AWS IoT core. So what I'm trying to do is to have a web page that lets me click on a subscribe button and I can start getting data from the simulated device. My question is first of all if this is possible and second can somebody guide me to where I can learn to do this.

Notes: What I've found is that for ruby there are gems for mqtt. But I'm not sure how to integrate with the web page. I've already sent a payload using irb and it was succesful. This is the repo I'm using: [njh/ruby-mqtt](https://github.com/njh/ruby-mqtt)

Thanks.
## [5][Having issues with adding gallery to a listing website am working on for some time now. Rails 6. Version](https://www.reddit.com/r/rails/comments/hxmsb4/having_issues_with_adding_gallery_to_a_listing/)
- url: https://www.reddit.com/r/rails/comments/hxmsb4/having_issues_with_adding_gallery_to_a_listing/
---
Please ur opinions will be of a great help to me.
## [6][How to Ignore raise error sidekiq worker | Capybara](https://www.reddit.com/r/rails/comments/hx8u87/how_to_ignore_raise_error_sidekiq_worker_capybara/)
- url: https://www.reddit.com/r/rails/comments/hx8u87/how_to_ignore_raise_error_sidekiq_worker_capybara/
---
I am pretty new to rails and my team recently moved to use sidekiq, I'm currently having troubles trying to create a Worker that retries upon a given condition (fails) and retrying it.

Calling a worker with this instruction within model

    CoolClassJob.perform_async(...) 

I'm using a worker with a code similar to this:

    class CoolClassJob   
      include Sidekiq::Worker   
        sidekiq_options queue: "payment", retry: 5    
        sidekiq_retry_in do |count| 
          10 
        end
     
        def perform() 
          ...         
          whatever = {...} 
          if whatever.status == 'successful'           
            thisCoolFunction                  # successfully ends job 
          elsif whatever.status == 'failed'           
            anotherCoolFunction               # successfully ends job 
          else whatever.pending?              # I want to retry if it falls in this condition since it is "waiting" for another task to complete. 
            raise 'We are trying again' 
          end 
          ... 
        end 
      ... 
      end
    ... 

I tried with

    begin raise 'We are trying again!' rescue nil end

But when I run my tests, I get this error:

    Failure/Error: raise 'We are trying again!' RuntimeError: 'We are trying again!' ...

Which of course, makes sense to me, since I'm raising the error, I tried searching but wasn't able to come up with a solution. I'm wondering whether its able toa) retry again without raising an error orb) tell Capybara (rspec) to keep trying without throwing an error.

I posted the question here: [https://stackoverflow.com/questions/63068356/rails-how-to-ignore-raise-error-sidekiq-worker-capybara](https://stackoverflow.com/questions/63068356/rails-how-to-ignore-raise-error-sidekiq-worker-capybara)
## [7][config.active_record.default_timezone = :local in rails 5.2.4 not working](https://www.reddit.com/r/rails/comments/hx1ca3/configactive_recorddefault_timezone_local_in/)
- url: https://www.reddit.com/r/rails/comments/hx1ca3/configactive_recorddefault_timezone_local_in/
---
I'd like to set all records in my database to my local timezone. I have had     

    config.active_record.default_timezone = :local
    
    in my application.rb for months, and it doesn't work. Does anybody have a solution?
## [8][Last used rails version 3.2](https://www.reddit.com/r/rails/comments/hwnyuf/last_used_rails_version_32/)
- url: https://www.reddit.com/r/rails/comments/hwnyuf/last_used_rails_version_32/
---
ending early 2013

what's the most important thing to know about version 6 compared to 3.2?

bonus question : I did notice people are charging for gems,  which is new.  How has that gone?
## [9][What's the difference between stimulus_reflex and stimulus.js?](https://www.reddit.com/r/rails/comments/hwm4et/whats_the_difference_between_stimulus_reflex_and/)
- url: https://www.reddit.com/r/rails/comments/hwm4et/whats_the_difference_between_stimulus_reflex_and/
---
Hello, I am kinda new to the rails world, and I am wondering what is the difference between  stimulus\_reflex and stimulus.js? since they're doing the same job.
## [10][Server Issue Serving](https://www.reddit.com/r/rails/comments/hwmmbs/server_issue_serving/)
- url: https://www.reddit.com/r/rails/comments/hwmmbs/server_issue_serving/
---
I am in dire need of help, as I am about to pull out what is left of my hair. I am trying to get my new server up and running, and I realize this may not be the right place for this post, but I am hopefully at least someone can point me in the right direction.

I have transferred my app to my server, still working on getting my deployment working properly, but that is an issue for another day. Currently, when I attempt to `curl` `127.0.1.1` I get the following:

`403 Forbidden`

`Forbidden`

`You don't have permission to access this resource.`

`Apache/2.4.38 (Debian) Server at` `127.0.1.1` `Port 80`

&amp;#x200B;

I am at my wits end with this server! My apache2/sites-available/app.conf

`&lt;VirtualHost *:80&gt;`

    `DocumentRoot /var/www/app/releases/current/`

&amp;#x200B;

    `&lt;Location /&gt;`

`Require all granted`

    `&lt;/Location&gt;`	

`&lt;/VirtualHost&gt;`

I am currently only trying to access the site via it's IP address as it will be taking over for an older server. I'll happily add anything that might be needed. Please reddit, you are my only hope!

&amp;#x200B;

Apache error.log

`[Thu Jul 23 15:34:50.275696 2020] [autoindex:error] [pid 7964:tid 140662913619712] [client` [`127.0.0.1:50404`](https://127.0.0.1:50404)`] AH01276: Cannot serve directory /var/www/app/releases/current/: No matching DirectoryIndex (index.html,index.cgi,`[`index.pl`](https://index.pl)`,index.php,index.xhtml,index.htm) found, and server-generated directory index forbidden by Options directive`

&amp;#x200B;

Apache.conf

&amp;#x200B;

`# Global configuration`

`DefaultRuntimeDir ${APACHE_RUN_DIR}`

`PidFile ${APACHE_PID_FILE}`

`Timeout 300`

`KeepAlive On`

`MaxKeepAliveRequests 100`

`KeepAliveTimeout 5`

`User ${APACHE_RUN_USER}`

`Group ${APACHE_RUN_GROUP}`

`HostnameLookups Off`

`ErrorLog ${APACHE_LOG_DIR}/error.log`

`LogLevel warn`

`IncludeOptional mods-enabled/*.load`

`IncludeOptional mods-enabled/*.conf`

`Include ports.conf`

`&lt;Directory /&gt;`

	`Options FollowSymLinks`

	`AllowOverride None`

	`Require all denied`

`&lt;/Directory&gt;`

`&lt;Directory /usr/share&gt;`

	`AllowOverride None`

	`Require all granted`

`&lt;/Directory&gt;`

`&lt;Directory /var/www/&gt;`

	`Options FollowSymLinks`

	`AllowOverride None`

	`Require all granted`

`&lt;/Directory&gt;`

`&lt;Directory /var/www/app&gt;`

	`Options FollowSymLinks`

	`AllowOverride None`

	`Require all granted`

`&lt;/Directory&gt;`

`AccessFileName .htaccess`

`&lt;FilesMatch "^\.ht"&gt;`

	`Require all denied`

`&lt;/FilesMatch&gt;`

`LogFormat "%v:%p %h %l %u %t \"%r\" %&gt;s %O \"%{Referer}i\" \"%{User-Agent}i\"" vhost_combined`

`LogFormat "%h %l %u %t \"%r\" %&gt;s %O \"%{Referer}i\" \"%{User-Agent}i\"" combined`

`LogFormat "%h %l %u %t \"%r\" %&gt;s %O" common`

`LogFormat "%{Referer}i -&gt; %U" referer`

`LogFormat "%{User-agent}i" agent`

`IncludeOptional conf-enabled/*.conf`

`IncludeOptional sites-enabled/*.conf`
## [11][Good templates/examples for action mailer email templates?](https://www.reddit.com/r/rails/comments/hwgd1t/good_templatesexamples_for_action_mailer_email/)
- url: https://www.reddit.com/r/rails/comments/hwgd1t/good_templatesexamples_for_action_mailer_email/
---
Anyone have resources for nice/good-looking action mailer email templates? Thinking one that incorporates a logo, has social links at the bottom, centered text etc. Any/all resources are appreciated.
## [12][Discovered SimpleCov, what can I do with it ?](https://www.reddit.com/r/rails/comments/hwf6um/discovered_simplecov_what_can_i_do_with_it/)
- url: https://www.reddit.com/r/rails/comments/hwf6um/discovered_simplecov_what_can_i_do_with_it/
---
Hi, I just discovered [SimpleCov](https://github.com/colszowka/simplecov) and looove the html output and the conclusions we can draw with it.

However, I can't really find a way to use it effectively. Thus I wonder, how do you use it (if you do of course) and what for ? 

&amp;#x200B;

I learnt about SimpleCov looking around Skunk that uses it at some point; this is IMO a great usage of SimpleCov.

&amp;#x200B;

I'd love a regular check of it on one of our projects that needs to be really well covered, maybe looking around Github action ? Anyway, just wanted to know if anyone uses it and what for. 

&amp;#x200B;

Thanks ! 

      _      _      _
    &gt;(.)__ &lt;(.)__ =(.)__
     (___/  (___/  (___/
