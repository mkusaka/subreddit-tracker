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
## [3][Sidekiq + NetSSH for running remote commands, or something else?](https://www.reddit.com/r/rails/comments/hkbjpu/sidekiq_netssh_for_running_remote_commands_or/)
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
## [4][Rails 6 with semantic ui and icons](https://www.reddit.com/r/rails/comments/hki5sm/rails_6_with_semantic_ui_and_icons/)
- url: https://www.reddit.com/r/rails/comments/hki5sm/rails_6_with_semantic_ui_and_icons/
---
Hi, i have a problem. I create a rails 6 project and add fomantic-ui(fork from semantic-ui) and its work. But i can't use a icon, i  instaled fontawesome(how to says a doc [https://semantic-ui.com/elements/icon.html](https://semantic-ui.com/elements/icon.html)) but system try found icons [http://localhost:3000/assets/themes/default/assets/fonts/icons.woff2](http://localhost:3000/assets/themes/default/assets/fonts/icons.woff2) how i can change path to my node\_module ? I think i need change my webpack settings, but i don't how
## [5][How do you organize Stripe Resources?](https://www.reddit.com/r/rails/comments/hk7qhb/how_do_you_organize_stripe_resources/)
- url: https://www.reddit.com/r/rails/comments/hk7qhb/how_do_you_organize_stripe_resources/
---
I'm building an app that uses Stripe as a backend for credit card issuing. I started off with an Account model that would keep a local copy of the Stripe Connected Account that I was associating with the User. I used a jsonb field to just store the json blob that Stripe returns via their API. 

Then, I need to create a Source (i.e. a bank account I could draw money from on behalf of the user). I did the hackiest thing I could and just added a jsonb "source" field to the Account model. Now I've got 2 json blobs that are keeping a local copy of Objects over at Stripe.

As I build this out, I'm going to need to cache a copy of a Stripe Balance Object maybe keep track of TopUPs, etc... 

I was just wondering how others deal with the complexity of Stripe. Are you creating separate models for each of the Stripe objects you need to cache? I've already gotten into a little trouble just with an Account model because I have to load it async into the front end. If I have a bunch of other models, how would I keep track of what's necessary in the front-end?
## [6][Rails &amp; Gatsby anyone?](https://www.reddit.com/r/rails/comments/hk22vq/rails_gatsby_anyone/)
- url: https://www.reddit.com/r/rails/comments/hk22vq/rails_gatsby_anyone/
---
I have a personal/portfolio website that I've built with a Rails 6 backend and a React frontend.  It's currently deployed to Heroku and using SendGrid for emails.  I'm also considering using the Comfortable Mexican Couch gem for blogs - but there is probably a Gatsby plugin for that too.

My biggest concern right now is search engine optimization and I would also like to increase the initial load time of the site.  

Since Gatsby just takes your dynamic React content and builds it into static resources, I figured that would take care of both issues.

Has anyone here attempted such a project?  Any advice or potential hurdles to look out for?
## [7][How to “call” custom fonts within with SLIM templates](https://www.reddit.com/r/rails/comments/hk9fxl/how_to_call_custom_fonts_within_with_slim/)
- url: https://www.reddit.com/r/rails/comments/hk9fxl/how_to_call_custom_fonts_within_with_slim/
---
Hello I've placed custom fonts for my app use, but I have no clue of how to display the asset pipeline fonts using SLIM templates in the *temas/index* pages, I was successful with images but I don't have any clue if (it's even possible?) to use custom fonts with asset pipeline using slim templates.

[https://github.com/LeoFragozo/notika\_slim\_test/blob/master/app/views/temas/index.html.slim](https://github.com/LeoFragozo/notika_slim_test/blob/master/app/views/temas/index.html.slim)

I've tried to put fonts in the asset pipeline, creating the font .scss file, and it's apparently not loading hence [https://postimg.cc/QVytpngz](https://postimg.cc/QVytpngz), so I'm here to ask for help.
## [8][Trailblazer 2.1 - new features, landingpage and documentation](https://www.reddit.com/r/rails/comments/hjyz5r/trailblazer_21_new_features_landingpage_and/)
- url: https://www.reddit.com/r/rails/comments/hjyz5r/trailblazer_21_new_features_landingpage_and/
---
For those who know trailblazer - you know that TRB evolved few times, introduced a lot of great ideas but also had his problems:  missing documentation, lack of communication with the community, and not consistent approach to some problems. We strongly believe that the current version with the support of core team and the new landing page with complete documentation is something that is worth working with.  


For those who somehow didn't hear about Trailblazer - check it out, probably for some complex projects where Rails aren't sufficient that tool could be really helpful and joyfull.  


Either you know TRB already or you never heard about it, check out the blog post about the history and changes of trailblazer 2.1 : [https://trailblazer.to/2.1/blog.html](https://trailblazer.to/2.1/blog.html) and give us feedback!
## [9][How would you set this up?](https://www.reddit.com/r/rails/comments/hk671l/how_would_you_set_this_up/)
- url: https://www.reddit.com/r/rails/comments/hk671l/how_would_you_set_this_up/
---
Hey Rails guys (and gals),

Part of a project am working on (basically a job board) should have the functionality for sending out a weekly email to a bunch of users. The email should contain a list of jobs the user can apply to. 

The list of jobs in the emails sent is different for each user since a user can create a subscription for the types of job notifications they want to be getting e.g. a subscription for "marketing jobs", "developer jobs" but the differences are based on "job category" 

The models are as below:

&amp;#x200B;

https://preview.redd.it/d990df9iii851.png?width=3693&amp;format=png&amp;auto=webp&amp;s=886f760d837b3224fe9019b347fabf9efc5c19c9

I've already created the notification mailer as below:

    class NotifyMailer &lt; ApplicationMailer
      default from: 'notify@domain.com'
      def self.send_request(payload)
        # first grab all users who are non-admins..then fetch all their active subscriptions...
        active_subscriptions = User.normal_users.with_active_subs
        emails = []
        active_subscriptions.each do |sub|
          emails &lt;&lt; sub.email 
        end
    
        emails.each do |email|
          new_request(email, payload).deliver_later
        end
      end
    end

...I've also created a background job as below:

    class SendNotificationsJob &lt; ApplicationJob
      queue_as :default
      def perform(*args)
        # Do something later
        NotifyMailer.send_request(payload)
      end
    end

Now this is where things feel a bit over my head - Essentially, I would like to send the emails once per week on a particular day, say Wednesday at 9am but I'm wondering:

&amp;#x200B;

1. Have I even set this up correctly? 
2. Where do I tell the app that the emails should be sent "once per week on a particular day, say Wednesday at 9am"? (Forgive the rather newbish question...am still learning -  All tutorials regarding sending email that I've come across mostly discuss the aspect of sending an email like after a user registers so this methodology is easily input in that particular controller method....but where do I put functionality for automatic email sending on particular dates/times? 

Thanks. 

PS: In case I've missed out on a detail that may help me out lemme know.
## [10][Nested form question](https://www.reddit.com/r/rails/comments/hk2rll/nested_form_question/)
- url: https://www.reddit.com/r/rails/comments/hk2rll/nested_form_question/
---
Hello people. 
I have a composite question. Have asked this here a few weeks ago, haven’t solved this since.
Part 1:
I am creating an e-commerce website with a cart. Cart has many LineItems. Both inherit from ApplicationRecord. I want LineItem to belong_to a Product, where a Product would be a class of external API, in this case ShopifyAPI::Product. Is it accomplishable?
(Hacky fix which I now did and don’t like: Right now the LineItem has a product_id column, respective product I find when needed through an API call)
Part 2: 
Product has two Options. Option has a range of possible values, say [‘1’, ‘2’, ‘3’]. Cart has many LineItems. On the cart page ‘carts/show’ I have a listing of all the LineItems. Each LineItem is rendered as it’s id and combination of options of the Product that belongs to the LineItem. So far so good. Each option is rendered as a dropdown of it’s range.
When the button ‘Submit’ on the Cart page is clicked, I try to save the cart with chosen options, I want the LineItems with respective selected options to be passed as the params in the next way: 
cart: { 
  line_items: [
    {...}, {...}
  ]
}. Where {...} are the combination of options. Would you know how to accomplish this? 

Thanks for reading. My guess is the second question would be easier solved with the first one, hence at first i thought of the second question, neither of which I know how to answer.
## [11][Pick up rails again after a break: advice needed](https://www.reddit.com/r/rails/comments/hjsu59/pick_up_rails_again_after_a_break_advice_needed/)
- url: https://www.reddit.com/r/rails/comments/hjsu59/pick_up_rails_again_after_a_break_advice_needed/
---
A few years ago (2015) I learned to develop full-stack web applications by following an online bootcamp. The bootcamp used Ruby on Rails as a framework. After the bootcamp I was able to develop my own web applications by using RoR.

Since 2017 I had hardly the time to develop since I had other career opportunities. However, now I would like to start building a new small application as a hobby project.

The app will mostly consist of CRUD operations and a real-time chat functionality. Nothing groundbreaking though. However I would like to experiment with providing the content in a true mobile app as well for iOS &amp; Android.

Before starting out though I have 2 questions where I hope you guys can help me with.

&amp;#x200B;

**Rails refresher**  
Before starting out I would like to refresh and recap my RoR skills. I was looking at the following two possibilities:

* Udemy course:  [The Complete Ruby on Rails Developer Course](https://www.udemy.com/course/the-complete-ruby-on-rails-developer-course/)
* Book: [Agile Web Development with Rails 6](https://pragprog.com/titles/rails6/)

Which one would you guys recommend? Any other recommendations?

&amp;#x200B;

**Feedback on application structure**  
My first idea was to follow the JS hype and build an API in Express and use VUE for the frontend. However after following some courses, I had the feeling that this setup was way too complex for a one-man hobby project. Furthermore my content should be optimised for SEO as well. In the JS world that would mean that I should use a Framework like Nuxt for example to serve SSR-content.

That's why I would chose for a "classic" RoR application. I could use Stimulus for the limited "interactive" functionalities and developing a native mobile app could be fairly easy [by wrapping web views](https://www.youtube.com/watch?v=SWEts0rlezA).

Is this a good idea or should I follow the trend by using Rails as the API and develop a separate VUE front-end app?

&amp;#x200B;

Thanks a lot for the feedback!
## [12][Rails 6 deployment with Passenger and Docker](https://www.reddit.com/r/rails/comments/hjuu1x/rails_6_deployment_with_passenger_and_docker/)
- url: https://www.reddit.com/r/rails/comments/hjuu1x/rails_6_deployment_with_passenger_and_docker/
---
Hi all,

Is there any good guide/resource on how to prepare and deploy Rails 6 application with Passenger/NGINX as Docker containers?
