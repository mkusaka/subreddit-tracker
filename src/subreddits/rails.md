# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/)
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
## [3][Job Queue Design for Multiple Small Transactions](https://www.reddit.com/r/rails/comments/hkz9pf/job_queue_design_for_multiple_small_transactions/)
- url: https://www.reddit.com/r/rails/comments/hkz9pf/job_queue_design_for_multiple_small_transactions/
---
Hi everyone

I have a task that I run via Sidekiq + ActiveJob + Redis. The task is a collection of ActiveRecord instances with a one to many relation to another model which performs updates on the "to many" table for each "belongs\_to". Is it considered best practice to have a single task for it in a loop or should I break it down to multiple smaller tasks to be put on queue?

Thanks
## [4][[Help] Link_to Remote, Cookies, and Jquery](https://www.reddit.com/r/rails/comments/hkqqfj/help_link_to_remote_cookies_and_jquery/)
- url: https://www.reddit.com/r/rails/comments/hkqqfj/help_link_to_remote_cookies_and_jquery/
---
I have an older application I'm updating where a user can rotate the order of a set of images and turn the visibility of those images on and off. I am using jquery to accomplish this and link_to remote. 

When I toggle the images on/off it works correctly. 
When I refresh the page manually the toggle state of the images is correct - I store them in cookies. 
If I refresh the partial to change the order, it changes the order but all images appear regardless of their toggle state. 

Is this an issue of using link_to remote with the partial and not passing some type of information to it?

Jquery Code:

    $("#distractorselecta").each(function() {
        $(this).toggle($.cookie('show-' + this.id) != 
    'collapsed');
    });

    $( ".toggle-distractora" ).click(function() {
      $( "#distractorselecta" ).toggle( "slow", function() {
    	  $.cookie('show-' + $(this).attr("id"), 
    $(this).is(":hidden") ? 'collapsed' : 'expanded');  
      });
    });

js.erb file for the partials and link_to remote

    $("#matching").html("&lt;%= escape_javascript(render(:partial =&gt; "fsets/matching")) %&gt;");
    $(".position-box").html("&lt;%=escape_javascript(render(:partial =&gt;"fsets/matchingposition")) %&gt;");
    $(".navbox").html("&lt;%=escape_javascript(render(:partial =&gt; "fsets/matchingnavigation")) %&gt;"); setMatching();

Here is a video of me testing it - the little refresh icon is the link_to remote which resets the images but ignores their toggle state. 
[video](https://imgur.com/a/mqznHSx)

Thank you.

Update: 

After messing around with everything a bit more, if I place the query code in the partial with some &lt;script&gt; tags then everything works as I want it to. I just don’t really understand why that would be. Currently the code is in the application.js file, but moving it to the partial next to the divs that get replaced keeps the toggle state working correctly. 

It just feels wrong...can anyone explain what’s happening here?
## [5][Vue slider component replacing on a vanilla POST rails form?](https://www.reddit.com/r/rails/comments/hkqpyn/vue_slider_component_replacing_on_a_vanilla_post/)
- url: https://www.reddit.com/r/rails/comments/hkqpyn/vue_slider_component_replacing_on_a_vanilla_post/
---
I am building a rails form with a number of input type="range" entries (https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/range) which I need to style in a specific way.

This need let me to the excellent https://nightcatsama.github.io/vue-slider-component/#/ component which meets my needs perfectly.

What I'd like to do is to use the slider component for entry but still submit a regular rails view form via POST.

Currently the form uses numerical values.

My plan is

1. install vue and stimulus
2. change the form to use hidden inputs for the slider-controlled values
3. change the form view to display the slider component
4. use a stimulus controller to copy the hidden fields data into the vue app data store on connect()
5. use a stimulus controller to copy the vue app data store back to the hidden fields on form submit

Is this feasible? Have you done something like this? I am trying to go as simple as possible, and only use Vue for the component I need.

Any advice on how to make this even simpler/bulletproof?

I am not 100% sure on how to combine both stimulus and vue in one js bundle, but found this gist of using a stimulus controller to load a view app: https://gist.github.com/crispinheneise/c9de022dc94c16131971b64d49c0778d
## [6][Building SMS with Twilio for bulk SMS text messages?](https://www.reddit.com/r/rails/comments/hkw8bp/building_sms_with_twilio_for_bulk_sms_text/)
- url: https://www.reddit.com/r/rails/comments/hkw8bp/building_sms_with_twilio_for_bulk_sms_text/
---
Yes, I did say SMS text message.

&amp;#x200B;

I am building an SMS platform.

let us assume they have 100,000 phone numbers.

&amp;#x200B;

I am using Twilio as well with Active Job and will later implement Redis with a 3rd party server alongside with it.

&amp;#x200B;

For anyone with experience, is this enough?

It's fairly simple, I have a form to enter the message body and then submit it.  It then does something like:

    Subscriber.each do |s|
        NotificationsJob.perform_later(message, s)
    end 

Then in the qued job, it sends the message.

Nothing needs to look pretty, it just needs to actually function and work.

Aside from regulatory stuff .. from a programming standpoint, is there anything else do or worry about when it comes to doing the background job task I need to worry about?
## [7][Cannot get RSpec/Capybara test to pass for user password change](https://www.reddit.com/r/rails/comments/hkkemy/cannot_get_rspeccapybara_test_to_pass_for_user/)
- url: https://www.reddit.com/r/rails/comments/hkkemy/cannot_get_rspeccapybara_test_to_pass_for_user/
---
Hi all, I've got several tests similar to this, but for some reason, I can't get this one to pass.

    Failure/Error: expect(u.valid_password?('new password')).to be true
           expected true
                got false

Here's the test:

    it "allows a user to change their password", focus: true do
      sign_in u
      visit '/users/edit'
      within("#edit_user") do
        fill_in 'user[password]', with: "new password"
        fill_in 'user[password_confirmation]', with: "new password"
        fill_in 'user[current_password]', with: u.password
      end
      click_button 'Save changes'
      u.reload
      expect(u.valid_password?('new password')).to be true
      expect(page).to have_content 'Your account has been updated successfully'
    end

[Here](https://imgur.com/xVH8pJ9) is a screenshot of  `tail -f log/test.log`

I don't get any errors when I run the test, it just returns false rather than true. When I test for what `u.password` is, it is the old password (it remains unchanged). I'm using devise.

Thanks in advance.  


  
EDIT: FIXED  


Turns out it was to do with my dodgy override for registrations/update, which was supposed to be allowing certain account updates without requiring a password. I had the following:  


    class Users::RegistrationsController &lt; Devise::RegistrationsController
      
      protected
    
      def update_resource(resource, params)
        params.delete :current_password
        resource.update_without_password(params)
      end
    
    end

I needed to change this to:  


    class Users::RegistrationsController &lt; Devise::RegistrationsController
      
      protected
    
      def update_resource(resource, params)
        # Require current password if user is trying to change password.
        return super if params["password"]&amp;.present?
    
        # Allows user to update registration information without password.
        resource.update_without_password(params.except("current_password"))
      end
    end

&amp;#x200B;
## [8][Sidekiq + NetSSH for running remote commands, or something else?](https://www.reddit.com/r/rails/comments/hkbjpu/sidekiq_netssh_for_running_remote_commands_or/)
- url: https://www.reddit.com/r/rails/comments/hkbjpu/sidekiq_netssh_for_running_remote_commands_or/
---
Right now, I am using the net-ssh gem to run commands remotely on multiple linux systems via SSH, starting within sidekiq workers. I used to have alternative approach that involved HTTP POST/GET requests to an API + ActiveRecord, but that became way too complicated -- database pooling issues, constant database polling issues, managing the agents, etc.

I just recently started running into an issue to where a job that runs in Sidekiq every minute and lasts 2-3 seconds would actually start hanging while I'm pushing commands to remote agents via net-ssh. Here's an example of the issue:

    [root:a8cd613bc7a5:~/myapp]# sidekiq -e production -q production
    2020-07-02T21:44:45.024Z pid=1807172 tid=go3qqgj28 INFO: Booted Rails 5.2.4 application in production environment
    2020-07-02T21:44:45.024Z pid=1807172 tid=go3qqgj28 INFO: Running in ruby 2.5.8p224 (2020-03-31 revision 67882) [x86_64-linux]
    2020-07-02T21:44:45.024Z pid=1807172 tid=go3qqgj28 INFO: See LICENSE and the LGPL-3.0 for licensing details.
    2020-07-02T21:44:45.024Z pid=1807172 tid=go3qqgj28 INFO: Upgrade to Sidekiq Pro for more features and support: http://sidekiq.org
    2020-07-02T21:44:45.024Z pid=1807172 tid=go3qqgj28 INFO: Booting Sidekiq 6.0.4 with redis options {:id=&gt;"Sidekiq-server-PID-1807172", :url=&gt;nil}
    2020-07-02T21:44:45.027Z pid=1807172 tid=go3qqgj28 INFO: Loading Schedule
    2020-07-02T21:44:45.028Z pid=1807172 tid=go3qqgj28 INFO: Scheduling every_minute_worker {"cron"=&gt;"0 * * * * *", "class"=&gt;"EveryMinuteWorker", "queue"=&gt;"production"}
    2020-07-02T21:44:45.031Z pid=1807172 tid=go3qqgj28 INFO: Scheduling daily_job_worker {"cron"=&gt;"00 12 * * * America/Detroit", "class"=&gt;"DailyJobWorker", "queue"=&gt;"production"}
    2020-07-02T21:44:45.036Z pid=1807172 tid=go3qqgj28 INFO: Schedules Loaded
    2020-07-02T21:44:51.599Z pid=1807172 tid=go3rto2tc class=EnumerationWorker jid=6c151ebb7275ac6d3eda1f24 INFO: start
    2020-07-02T21:45:00.044Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T21:45:00.044Z pid=1807172 tid=go3rto384 class=EveryMinuteWorker jid=a54d8cacf904da8599eafb98 INFO: start
    2020-07-02T21:45:02.313Z pid=1807172 tid=go3rto384 class=EveryMinuteWorker jid=a54d8cacf904da8599eafb98 elapsed=2.268 INFO: done
    2020-07-02T21:45:12.601Z pid=1807172 tid=go3rto384 class=HttpWorker jid=949ba99e2dae36872d3a6bfd INFO: start
    2020-07-02T21:45:12.602Z pid=1807172 tid=go3rto19s class=HttpsWorker jid=981f165c4daa82b4c0c065bc INFO: start
    2020-07-02T21:46:00.102Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T21:46:00.102Z pid=1807172 tid=go3rto0nc class=EveryMinuteWorker jid=1aecdcc734c173c953bb6052 INFO: start
    2020-07-02T21:46:02.501Z pid=1807172 tid=go3rto0nc class=EveryMinuteWorker jid=1aecdcc734c173c953bb6052 elapsed=2.399 INFO: done
    2020-07-02T21:47:00.160Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T21:47:00.160Z pid=1807172 tid=go3rto1us class=EveryMinuteWorker jid=c5d878b4884d6bffeed222ba INFO: start
    2020-07-02T21:47:03.009Z pid=1807172 tid=go3rto1us class=EveryMinuteWorker jid=c5d878b4884d6bffeed222ba elapsed=2.849 INFO: done
    2020-07-02T21:48:00.222Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T21:48:00.223Z pid=1807172 tid=go3rto0gk class=EveryMinuteWorker jid=c05a9e7548c76a986ff90b2c INFO: start
    2020-07-02T21:48:02.581Z pid=1807172 tid=go3rto0gk class=EveryMinuteWorker jid=c05a9e7548c76a986ff90b2c elapsed=2.358 INFO: done
    2020-07-02T21:49:00.282Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T21:49:00.283Z pid=1807172 tid=go3rto0gk class=EveryMinuteWorker jid=1c00160fa0d8c909dc874ca0 INFO: start
    2020-07-02T21:49:02.653Z pid=1807172 tid=go3rto0gk class=EveryMinuteWorker jid=1c00160fa0d8c909dc874ca0 elapsed=2.369 INFO: done
    2020-07-02T21:50:00.042Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T21:50:00.043Z pid=1807172 tid=go3rto1h0 class=EveryMinuteWorker jid=3eae22a8af2730a0c524727a INFO: start
    2020-07-02T21:51:00.102Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T21:51:00.102Z pid=1807172 tid=go3rto218 class=EveryMinuteWorker jid=6ccbbd6e7dec18c5b0cd270d INFO: start
    2020-07-02T21:52:00.159Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T21:52:00.160Z pid=1807172 tid=go3rto0co class=EveryMinuteWorker jid=521026754ce0d517fa2dcd33 INFO: start
    2020-07-02T21:53:00.218Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T21:53:00.219Z pid=1807172 tid=go3rto0gk class=EveryMinuteWorker jid=ce17658e7bf76c25adfeabae INFO: start
    2020-07-02T21:54:00.275Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T21:54:00.276Z pid=1807172 tid=go3rto2f0 class=EveryMinuteWorker jid=a9ba8969acae516b07d837bd INFO: start
    2020-07-02T21:55:00.033Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T21:55:00.034Z pid=1807172 tid=go3rto0nc class=EveryMinuteWorker jid=2d66f0045c8ab7bcb28fa28b INFO: start
    2020-07-02T21:56:00.093Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T21:56:00.094Z pid=1807172 tid=go3rto1us class=EveryMinuteWorker jid=285067110201c2586a7fd07c INFO: start
    2020-07-02T21:57:00.151Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T21:58:00.207Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T21:59:00.263Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:00:00.020Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:01:00.077Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:01:02.176Z pid=1807172 tid=go3rto384 class=HttpWorker jid=949ba99e2dae36872d3a6bfd elapsed=949.575 INFO: done
    2020-07-02T22:01:02.177Z pid=1807172 tid=go3rto384 class=EveryMinuteWorker jid=a856fc66c1a0e81046714c71 INFO: start
    2020-07-02T22:02:00.136Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:03:00.196Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:04:00.254Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:05:00.012Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:06:00.072Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:07:00.131Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:08:00.190Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:09:00.248Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:10:00.006Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:11:00.062Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:12:00.120Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:13:00.178Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:14:00.234Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:15:00.291Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:16:00.050Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:17:00.109Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:18:00.167Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:19:00.224Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:20:00.282Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:20:49.618Z pid=1807172 tid=go3rto19s class=HttpsWorker jid=981f165c4daa82b4c0c065bc elapsed=2137.017 INFO: done
    2020-07-02T22:20:49.619Z pid=1807172 tid=go3rto19s class=EveryMinuteWorker jid=72cc6c5060105e45e872c7f6 INFO: start
    2020-07-02T22:20:53.160Z pid=1807172 tid=go3rto2tc class=EnumerationWorker jid=6c151ebb7275ac6d3eda1f24 elapsed=2161.56 INFO: done
    2020-07-02T22:20:53.161Z pid=1807172 tid=go3rto2tc class=EveryMinuteWorker jid=358f556d7f5e49ad0130b1b5 INFO: start
    2020-07-02T22:20:56.982Z pid=1807172 tid=go3rto1h0 class=EveryMinuteWorker jid=3eae22a8af2730a0c524727a elapsed=1856.938 INFO: done
    2020-07-02T22:20:56.983Z pid=1807172 tid=go3rto1h0 class=EveryMinuteWorker jid=fecc3022964b8c2833381033 INFO: start
    2020-07-02T22:21:00.066Z pid=1807172 tid=otv92mi6o INFO: queueing EveryMinuteWorker (every_minute_worker)
    2020-07-02T22:21:04.267Z pid=1807172 tid=go3rto0nc class=EveryMinuteWorker jid=2d66f0045c8ab7bcb28fa28b elapsed=1564.233 INFO: done
    2020-07-02T22:21:04.268Z pid=1807172 tid=go3rto0nc class=EveryMinuteWorker jid=4dc55fec933e63d8a2375928 INFO: start
    2020-07-02T22:21:04.396Z pid=1807172 tid=go3rto0gk class=EveryMinuteWorker jid=ce17658e7bf76c25adfeabae elapsed=1684.177 INFO: done
    2020-07-02T22:21:04.406Z pid=1807172 tid=go3rto1h0 class=EveryMinuteWorker jid=fecc3022964b8c2833381033 elapsed=7.423 INFO: done
    2020-07-02T22:21:04.507Z pid=1807172 tid=go3rto384 class=EveryMinuteWorker jid=a856fc66c1a0e81046714c71 elapsed=1202.33 INFO: done
    2020-07-02T22:21:04.543Z pid=1807172 tid=go3rto2f0 class=EveryMinuteWorker jid=a9ba8969acae516b07d837bd elapsed=1624.267 INFO: done
    2020-07-02T22:21:04.842Z pid=1807172 tid=go3rto0co class=EveryMinuteWorker jid=521026754ce0d517fa2dcd33 elapsed=1744.682 INFO: done
    2020-07-02T22:21:05.051Z pid=1807172 tid=go3rto2tc class=EveryMinuteWorker jid=358f556d7f5e49ad0130b1b5 elapsed=11.89 INFO: done
    2020-07-02T22:21:05.077Z pid=1807172 tid=go3rto19s class=EveryMinuteWorker jid=72cc6c5060105e45e872c7f6 elapsed=15.458 INFO: done
    2020-07-02T22:21:05.097Z pid=1807172 tid=go3rto1us class=EveryMinuteWorker jid=285067110201c2586a7fd07c elapsed=1505.003 INFO: done
    2020-07-02T22:21:05.223Z pid=1807172 tid=go3rto218 class=EveryMinuteWorker jid=6ccbbd6e7dec18c5b0cd270d elapsed=1805.121 INFO: done
    2020-07-02T22:21:06.930Z pid=1807172 tid=go3rto0nc class=EveryMinuteWorker jid=4dc55fec933e63d8a2375928 elapsed=2.663 INFO: done

As you can see, the worker that runs every minute for 2-3 seconds eventually stopped running and started overlapping. It never finished until the last worker (HttpsWorker) completed, which is just simply pushing commands to a linux system via net-ssh.

Now that this is happening, I'm starting to question this entire process. My rails app provides a UI for customers to schedule tasks, and Sidekiq processes those tasks on a Linux VM that sits on their network at a predefined time. Sidekiq is really nice for me because I can queue multiple simultaneous tasks, and also see straight from the console what's failing, manually re-queue if necessary, etc.

Here are the few questions that I have regarding this:

1. Are there potential scalability issues that may exist with this method? So far, so good, but I have a small customer base, so no issues yet other than jobs hanging.
2. Is there a better way to manage something like this while still being able to enjoy the benefits that I appreciate (e.g. tracking crashes, re-tries, etc.)
3. If this method is actually something that should be OK, what could be causing Sidekiq to do this with the jobs? I thought each job was 100% independent of any other job as long it's running, but this hiccup and lag is very intermittent and only when I have other jobs running, and not all of the time.
## [9][Rails 6 with semantic ui and icons](https://www.reddit.com/r/rails/comments/hki5sm/rails_6_with_semantic_ui_and_icons/)
- url: https://www.reddit.com/r/rails/comments/hki5sm/rails_6_with_semantic_ui_and_icons/
---
Hi, i have a problem. I create a rails 6 project and add fomantic-ui(fork from semantic-ui) and its work. But i can't use a icon, i  instaled fontawesome(how to says a doc [https://semantic-ui.com/elements/icon.html](https://semantic-ui.com/elements/icon.html)) but system try found icons [http://localhost:3000/assets/themes/default/assets/fonts/icons.woff2](http://localhost:3000/assets/themes/default/assets/fonts/icons.woff2) how i can change path to my node\_module ? I think i need change my webpack settings, but i don't how
## [10][How do you organize Stripe Resources?](https://www.reddit.com/r/rails/comments/hk7qhb/how_do_you_organize_stripe_resources/)
- url: https://www.reddit.com/r/rails/comments/hk7qhb/how_do_you_organize_stripe_resources/
---
I'm building an app that uses Stripe as a backend for credit card issuing. I started off with an Account model that would keep a local copy of the Stripe Connected Account that I was associating with the User. I used a jsonb field to just store the json blob that Stripe returns via their API. 

Then, I need to create a Source (i.e. a bank account I could draw money from on behalf of the user). I did the hackiest thing I could and just added a jsonb "source" field to the Account model. Now I've got 2 json blobs that are keeping a local copy of Objects over at Stripe.

As I build this out, I'm going to need to cache a copy of a Stripe Balance Object maybe keep track of TopUPs, etc... 

I was just wondering how others deal with the complexity of Stripe. Are you creating separate models for each of the Stripe objects you need to cache? I've already gotten into a little trouble just with an Account model because I have to load it async into the front end. If I have a bunch of other models, how would I keep track of what's necessary in the front-end?
## [11][Rails &amp; Gatsby anyone?](https://www.reddit.com/r/rails/comments/hk22vq/rails_gatsby_anyone/)
- url: https://www.reddit.com/r/rails/comments/hk22vq/rails_gatsby_anyone/
---
I have a personal/portfolio website that I've built with a Rails 6 backend and a React frontend.  It's currently deployed to Heroku and using SendGrid for emails.  I'm also considering using the Comfortable Mexican Couch gem for blogs - but there is probably a Gatsby plugin for that too.

My biggest concern right now is search engine optimization and I would also like to increase the initial load time of the site.  

Since Gatsby just takes your dynamic React content and builds it into static resources, I figured that would take care of both issues.

Has anyone here attempted such a project?  Any advice or potential hurdles to look out for?
## [12][How to “call” custom fonts within with SLIM templates](https://www.reddit.com/r/rails/comments/hk9fxl/how_to_call_custom_fonts_within_with_slim/)
- url: https://www.reddit.com/r/rails/comments/hk9fxl/how_to_call_custom_fonts_within_with_slim/
---
Hello I've placed custom fonts for my app use, but I have no clue of how to display the asset pipeline fonts using SLIM templates in the *temas/index* pages, I was successful with images but I don't have any clue if (it's even possible?) to use custom fonts with asset pipeline using slim templates.

[https://github.com/LeoFragozo/notika\_slim\_test/blob/master/app/views/temas/index.html.slim](https://github.com/LeoFragozo/notika_slim_test/blob/master/app/views/temas/index.html.slim)

I've tried to put fonts in the asset pipeline, creating the font .scss file, and it's apparently not loading hence [https://postimg.cc/QVytpngz](https://postimg.cc/QVytpngz), so I'm here to ask for help.
