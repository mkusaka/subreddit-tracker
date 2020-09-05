# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/)
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
## [2][how do you deal with user timezones?](https://www.reddit.com/r/rails/comments/imym4y/how_do_you_deal_with_user_timezones/)
- url: https://www.reddit.com/r/rails/comments/imym4y/how_do_you_deal_with_user_timezones/
---
in a previous project i was saving the user timezone when signing up with an hidden field filled up with javascript to get the user's timezone and then used that timezone in views and mailers

but in my current project users can sign up with facebook so i can't add an hidden field

do i need to store the user timezone?

how would you localize datetimes in mailers without the user timezone stored?

and more generally how do you deal with user timezones?
## [3][What is your solution for page-specific javascript in rails?](https://www.reddit.com/r/rails/comments/imrqlk/what_is_your_solution_for_pagespecific_javascript/)
- url: https://www.reddit.com/r/rails/comments/imrqlk/what_is_your_solution_for_pagespecific_javascript/
---
I've been researching this throughout the day.  The most straight forward solution I've found is to put it in a content\_for in the view, and pull it into the layout right before the body tag, but that doesn't seem very 'unobtrusive.'

Is there is Rails standard, or community preferred method for this?
## [4][Assets result in a 404](https://www.reddit.com/r/rails/comments/in0bqy/assets_result_in_a_404/)
- url: https://www.reddit.com/r/rails/comments/in0bqy/assets_result_in_a_404/
---
Hi. Tested whether my app would run on production (still on my local computer). However, all requests to them result in 404s.

What I did:

```sh
$ export RAILS_ENV=production
$ rails assets:precompile
$ rails s
```

At this point, neither my stylesheets nor my JavaScripts are loaded. I'm using the default configuration for `config/environments/production.rb`, meaning that the only assets settings is `compile`, which is set to `false`, which threw me off at first.

Tried turning it on, now the stylesheets loaded, but still no JavaScript. What is going on, I really can't figure it out?
## [5][Podcasting platforms or tools built with Ruby on Rails?](https://www.reddit.com/r/rails/comments/imib6f/podcasting_platforms_or_tools_built_with_ruby_on/)
- url: https://www.reddit.com/r/rails/comments/imib6f/podcasting_platforms_or_tools_built_with_ruby_on/
---
Hey guys,

There are tons of Podcast platforms out there, but can any of you tell me if some of them are built with Rails?  

I'd be surprised if there were none.  Also, doesn't have to be the full platform, but maybe support tools like RSS feed builders or episode/video deployment tools.
## [6][Is there a stylistic reason NOT to with_indifferent_access every hash?](https://www.reddit.com/r/rails/comments/imhxf7/is_there_a_stylistic_reason_not_to_with/)
- url: https://www.reddit.com/r/rails/comments/imhxf7/is_there_a_stylistic_reason_not_to_with/
---
Hi,

Curious what the community thinks on this point. In many Rails codebases I've worked with `with_indifferent_access` is used _defensively_ to avoid the pain of `nil` returns.

```
# my_hash = { string_key: 'whoever did this must be a troll..' }

my_hash['string_key'] 
# =&gt; nil 
# argh!!!
```

Wondering if there are any stylistic reasons or situations that one would NOT convert all hashes `with_indifferent_access` and use a string key over a symbol (or mix of both) to semantically imply something about a hash value -- or perhaps this is just a quirk of Ruby?
## [7][API question](https://www.reddit.com/r/rails/comments/imgph5/api_question/)
- url: https://www.reddit.com/r/rails/comments/imgph5/api_question/
---
I have a decently successful iOS application that I am adding team/subscriptions to. Basically issuing licenses by the number of seats they want. 

To do so, I built a simple rails application that does three main functions: it allows a company to be created, to add employees (or seats) to the company roster, and to allow them the pay a monthly fee. 

Now I need to create an API that can speak with my iOS application to: 
1) verify that an account has paid (which my rails app w/ stripe integrated handles)
2) see what users (email addresses) are part of the account 

I am wondering what services you might use for this? Doorkeeper or something like this. Or JWT? I’m not that experienced with API’s and I’d love to simply allow api access for my iOS app server to then ingest the data from my rails app (in a secure way)
## [8][iCallendar issue](https://www.reddit.com/r/rails/comments/imdmnc/icallendar_issue/)
- url: https://www.reddit.com/r/rails/comments/imdmnc/icallendar_issue/
---
I am having some issues with icalendar gem,  it is adding events   in UTC+00 and I wanted it to be readjusted to match users  timezone, how can I do it?
## [9][[HELP] Preset data in dynamic forms.](https://www.reddit.com/r/rails/comments/imcpmz/help_preset_data_in_dynamic_forms/)
- url: https://www.reddit.com/r/rails/comments/imcpmz/help_preset_data_in_dynamic_forms/
---
I have added dynamic form feature inside the rails application.The end-user can create form fields by selecting input fields from a bunch of available input options i.e text field, radio buttons, file field etc .    It works perfectly fine. After the form is created, end-users can submit this form without any issue. Now,I have a use case in which I would like to pre-set data in some of the input fields. For instance, an admin-user created a form with the following input fields: `First Name(String), Last Name(String),Father's First Name(String), Father's Last Name(String)`, etc.

Now, a member logs in to fill in this form. I would like to preset  data in `First Name, Last Name` fields with the first\_name, last\_name of the logged in user  by fetching this information from the user's profile page. From a usability point of view, it would be a burden for the user to type this information once again. Since the form fields are created dynamically and the  name attribute of each of the input fields is a generated at run time so its hard to figure out that which input field corresponds to what user information. What could be the reasonable solution to preset data in dynamic form input fields? Any help will be highly appreciated.  Thanks.
## [10][missing /tmp/passenger.xxxxx directory with standalone passenger run by logrotate](https://www.reddit.com/r/rails/comments/imcai3/missing_tmppassengerxxxxx_directory_with/)
- url: https://www.reddit.com/r/rails/comments/imcai3/missing_tmppassengerxxxxx_directory_with/
---
Hi everyone.

I'm running a phusion passenger standalone instance and everything works fine except the automatic restart after the log rotation which occurs every night. When I wake up, the "built-in" nginx reports a 502 bad gateway.

I'm using a script to stop/start passenger and when I restart in an interactive ssh session, everything works fine. However, when logrotate executes the same script, passenger restarts but there's a missing directory in /tmp so the built-in nginx can't speak with passenger.

I'm at a loss. I don't know what's wrong and what to check next.

My script is just doing this :

    #!/bin/bash
    su -l myuser /home/myuser/scripts/passenger_stop.sh
    su -l myuser /home/myuser/scripts/passenger_start.sh

passenger\_start.sh :

    #!/bin/bash
    export RBENV_ROOT="/opt/rbenv"
    export PATH="/opt/rbenv/plugins/ruby-build/bin:/opt/rbenv/bin:$PATH"
    eval "$(rbenv init -)"
    alias node=nodejs
    SOCKET="/home/myuser/passenger-nginx.socket"
    PID_FILE="/home/myuser/project/shared/pids/passenger.pid"
    cd /home/myuser/project/current
    bundle exec passenger start -e production --socket=${SOCKET} -d --pid-file ${PID_FILE}

The passenger.log file :

    [ N 2020-09-04 00:05:31.7264 20421/T1 age/Cor/CoreMain.cpp:1340 ]: Starting Passenger core...
    [ N 2020-09-04 00:05:31.7265 20421/T1 age/Cor/CoreMain.cpp:256 ]: Passenger core running in multi-application mode.
    [ N 2020-09-04 00:05:31.7350 20421/T1 age/Cor/CoreMain.cpp:1015 ]: Passenger core online, PID 20421
    [ N 2020-09-04 00:05:33.8468 20421/T5 age/Cor/SecurityUpdateChecker.h:519 ]: Security update check: no update found (next check in 24 hours)
    /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler/shared_helpers.rb:71:in `getwd': No such file or directory - getcwd (Errno::ENOENT)
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler/shared_helpers.rb:71:in `pwd'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler/shared_helpers.rb:71:in `block in pwd'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/monitor.rb:235:in `mon_synchronize'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler/shared_helpers.rb:70:in `pwd'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler/shared_helpers.rb:272:in `search_up'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler/shared_helpers.rb:155:in `print_major_deprecations!'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler.rb:103:in `setup'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler/setup.rb:20:in `&lt;top (required)&gt;'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
    2020/09/04 00:05:51 [crit] 20441#0: *3 connect() to unix:/passenger_core failed (2: No such file or directory) while connecting to upstream, client: unix:, server: _, request: "GET / HTTP/1.0", upstream: "passenger:unix:/passenger_core:", host: "myhost"
    2020/09/04 00:05:52 [crit] 20441#0: *5 connect() to unix:/tmp/passenger.Bh5wxh4/agents.s/core failed (2: No such file or directory) while connecting to upstream, client: unix:, server: _, request: "GET / HTTP/1.0", upstream: "passenger:unix:/tmp/passenger.Bh5wxh4/agents.s/core:", host: "myhost"

The last line repeats for every request.

The process command lines (for nginx and PassengerAgent) are referencing this "/tmp/passenger.something" directory but it doesn't exist.

I just have to execute the restart script again (in ssh interactive, not in my logrotate script) and both the command lines and the tmp directory are ok. The command lines are then exactly the same except for the random tmp path and the referenced pids.

Phusion Passenger v6.0.6 with nginx 1.18.0 (I guess its the default for passenger 6.0.6) running on ubuntu server 20.04 up-to-date.

&amp;#x200B;

Thanks for your time and for the help :-)
## [11][Can't read "public/tmp/" in development &amp; production environment](https://www.reddit.com/r/rails/comments/im9wx6/cant_read_publictmp_in_development_production/)
- url: https://www.reddit.com/r/rails/comments/im9wx6/cant_read_publictmp_in_development_production/
---
It works well in localhost, but when I deployed it to development &amp; production, it returns **No such file or directory @ rb\_sysopen - /usr/src/app/public/tmp/file\_name.csv** error when I try to read the file.

The file is created with sidekiq job, and the sidekiq jobs are processed (success) so I assume the file is created since I don't have access to the server directory. This is how the file is created inside sidekiq job.

        dir_name = "#{Rails.root}/public/tmp"
        FileUtils.mkdir_p dir_name
        file = "#{dir_name}/file_name.csv"
        header = ["test"]  
    
        CSV.open( file, 'w' ) do |writer|
            writer &lt;&lt; header
            body.each do |b|
                writer &lt;&lt; [b["test"]]
            end
        end

And this is how the file is read in controller.

        file_path = "#{Rails.root}/public/tmp/#{filename}"
        File.open(file_path, 'r') do |f|
            send_data f.read,
                    :type =&gt; 'text/csv; charset=iso-8859-1; header=present',
                              :disposition =&gt; "attachment; filename=#{filename}"
        end
