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
## [3][config.active_record.default_timezone = :local in rails 5.2.4 not working](https://www.reddit.com/r/rails/comments/hx1ca3/configactive_recorddefault_timezone_local_in/)
- url: https://www.reddit.com/r/rails/comments/hx1ca3/configactive_recorddefault_timezone_local_in/
---
I'd like to set all records in my database to my local timezone. I have had     

    config.active_record.default_timezone = :local
    
    in my application.rb for months, and it doesn't work. Does anybody have a solution?
## [4][Last used rails version 3.2](https://www.reddit.com/r/rails/comments/hwnyuf/last_used_rails_version_32/)
- url: https://www.reddit.com/r/rails/comments/hwnyuf/last_used_rails_version_32/
---
ending early 2013

what's the most important thing to know about version 6 compared to 3.2?

bonus question : I did notice people are charging for gems,  which is new.  How has that gone?
## [5][What's the difference between stimulus_reflex and stimulus.js?](https://www.reddit.com/r/rails/comments/hwm4et/whats_the_difference_between_stimulus_reflex_and/)
- url: https://www.reddit.com/r/rails/comments/hwm4et/whats_the_difference_between_stimulus_reflex_and/
---
Hello, I am kinda new to the rails world, and I am wondering what is the difference between  stimulus\_reflex and stimulus.js? since they're doing the same job.
## [6][Server Issue Serving](https://www.reddit.com/r/rails/comments/hwmmbs/server_issue_serving/)
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
## [7][Discovered SimpleCov, what can I do with it ?](https://www.reddit.com/r/rails/comments/hwf6um/discovered_simplecov_what_can_i_do_with_it/)
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
## [8][Good templates/examples for action mailer email templates?](https://www.reddit.com/r/rails/comments/hwgd1t/good_templatesexamples_for_action_mailer_email/)
- url: https://www.reddit.com/r/rails/comments/hwgd1t/good_templatesexamples_for_action_mailer_email/
---
Anyone have resources for nice/good-looking action mailer email templates? Thinking one that incorporates a logo, has social links at the bottom, centered text etc. Any/all resources are appreciated.
## [9][Having trouble with Capybara not waiting for remote form](https://www.reddit.com/r/rails/comments/hwdxnr/having_trouble_with_capybara_not_waiting_for/)
- url: https://www.reddit.com/r/rails/comments/hwdxnr/having_trouble_with_capybara_not_waiting_for/
---
I have a view where customers can change the delivery date of their shipments by selecting a date in [Flatpickr](https://flatpickr.js.org/). The flatpickr instance is connected to a remote `form_with`input field which submits the form on change via stimulus. 

This is all working manually in the browser but I obviously also want to test it automatically via system tests (I'm using Rails' system tests with rspec, selenium and capybara). However, it looks like Capybara doesn't wait long enough for the remote form to submit. I currently have the following test:

&gt; delivery_date = find("span", text: "27")

&gt; expect {

&gt; &gt;delivery_date.click

&gt; }.to change { shipment.reload.delivery_date }

which fails. However, if I change it to sleep for a second it works:

&gt; delivery_date = find("span", text: "27")

&gt; expect {

&gt; &gt; delivery_date.click

&gt; &gt; sleep 1

&gt; }.to change { shipment.reload.delivery_date }

Is there something that I'm doing wrong here like some syntax that I'm not following? The only thing that I manually added to rspec regarding system tests is the following in my `rails_helper.rb`:

&gt;  config.before(:each, type: :system) do

&gt;    driven_by :rack_test

&gt;  end

&gt;  config.before(:each, type: :system, js: true) do

&gt;    driven_by :selenium_chrome_headless

&gt;  end

which just changes the driver based on whether or not the test needs js (which I have enabled for the above test). I also tried changing the `Capybara.default_max_wait_time` to 10 in a before action without luck.
## [10][Trouble with Elasticsearch DSL in a Ruby on Rails Project](https://www.reddit.com/r/rails/comments/hw5p69/trouble_with_elasticsearch_dsl_in_a_ruby_on_rails/)
- url: https://www.reddit.com/r/rails/comments/hw5p69/trouble_with_elasticsearch_dsl_in_a_ruby_on_rails/
---
Hey All, I'm trying to implement fuzzy search with Elasticsearch. I'm able to return results when I type in the query correctly, but now I am having trouble with fuzziness. I'm not having any luck with any of the online guides, does anyone have suggestions for a rails article with elasticsearch and a simple fuzzy search set up?  


`def search`  
 `@elasticsearchresults = Car.__elasticsearch__.search(`  
`{`  
 `query: {`  
 `match: {`  
 `message: {`  
 `query: params[:q],`  
 `fuzziness: 1`  
`}`  
`}`  
`}`  
`}).records`  
`end`
## [11][Security issues with generating temporary passwords](https://www.reddit.com/r/rails/comments/hw2ile/security_issues_with_generating_temporary/)
- url: https://www.reddit.com/r/rails/comments/hw2ile/security_issues_with_generating_temporary/
---
Hello. So I am developing an app with an application system. The user applies to my platform and then an admin or somebody with high enough permissions decides whether to accept the application or deny it. 

Let's say he accepts it. When he accepts it, the user is automatically created. I am wondering if it is safe to generate a temporary password with something like [Passgen gem](https://github.com/cryptice/Passgen) and then store it as the user's password until he changes it.

I would also need to send a confirmation email to his email address with a success message and a temporary password. Is it safe to send a password like that?

I am also wondering. If I hash that random password, how do I send it in plain text for a user to understand in the email?

Or is there any other, better way to handle this situation. Thanks!
## [12][How to insert font-awesome icons into action mailer views?](https://www.reddit.com/r/rails/comments/hw7ppm/how_to_insert_fontawesome_icons_into_action/)
- url: https://www.reddit.com/r/rails/comments/hw7ppm/how_to_insert_fontawesome_icons_into_action/
---
No problem embedding images, but haven't had luck leveraging FA-icons for this.
