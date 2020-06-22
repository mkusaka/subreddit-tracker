# aws
## [1][PSA: Don't take remote exams offered by Pearson Vue (OnVue) for AWS Certifications!](https://www.reddit.com/r/aws/comments/fscq7v/psa_dont_take_remote_exams_offered_by_pearson_vue/)
- url: https://www.reddit.com/r/aws/comments/fscq7v/psa_dont_take_remote_exams_offered_by_pearson_vue/
---
I can't describe how horrible this experience was.  I am not looking forward to how much work I am going to have to do to get my money back.  This is not my first AWS certification (I have SA Pro and DevOps Pro), but is my first online exam.  The short version is: Don't take AWS exams via the Pearson Vue at home option, even if it is offered.  AWS should not be offering this option as I can attest it is a waste of time.  Ironically, AWS would have us use their services because of their high availability and scaling but apparently they don't ask their test partners to do the same!

It started off easy enough: I passed the initial 'checks' as it confirmed my internet speed, camera access, and microphone access.  I started the process 15+ minutes before my scheduled exam time.  I was able to open the app, it again verified the technical requirements passed, and I went to the next screen.  It asked for my cell phone number and texted me a link which opened a web page which requested to take my photo.  Easy enough.  I did that and then the web page went to 'Uploading and verifying photo'.  A spinning circle started spinning.  This is where my test experience ended, but not where the poor experience ends.  I tried again, and then a third time.  Same experience.  As I write this, I left it on that page and the spinning is continuing.  This screen has been spinning for no less than 45 minutes.  At 8 minutes before my scheduled exam, I tried finding the help link.  A chat window opened, and I waited, and waited, and waited.  Still waiting as I write this.  My chat window has been open for 52 minutes and still no one to help.  Every two minutes I get ' All agents are currently assisting others. Thank you for your patience.' written in the window.  OK - what next?  They make it harder to find, but I got a phone number I can call.  I tried calling that.  Busy signal.  For the next 20 minutes I called back and back, busy signal.  Finally, I got it to actually pick up, but of course no human yet.  No estimate of time to when I can be helped.  They don't even have nice elevator music to listen to.  Who knows when I will be able to talk to someone.  This has been an exceedingly poor experience.

If you value your time, please do yourself a favor and don't even attempt a online exam with Pearson.  I worked hard to prepare for this exam and rescheduled things to fit around it.  Now, I will have to do that all again.

u/jeffbarr Is this the experience AWS is hoping to get with their testing partners?  This was a waste of my time and money.  Amazon should seriously reevaluate the quality of their test partners.  I understand everyone is trying to deal with all the issues.  However, if you can't offer quality testing, then please don't offer the option at all.  It isn't respectful to people's time.  Pearson is well aware of their capacity and if it isn't up to requirements, they shouldn't be scheduling test slots.

&amp;#x200B;

*EDIT*: A few background items I didn't initially share that may be relevant for others.  For the computer, I used a fully up to date Windows 10 laptop.  The laptop itself is only about a month old and is in near pristine condition.  Other than a few applications like Office, there is barely anything installed on there yet.  I used a hard wired connection, like recommended by Pearson through the use of a usb-to-ethernet adapter.  I have Verizon FIOS (980Mbps/840Mbps) and did do a speed test way after it was apparent this would not work.  I forget the exact numbers, but I was still pulling in hundreds of Mbps in both directions, despite everyone being at home and using the USB ethernet adapater which does put a cap on my speed, but I can't see hundreds of Mbps not being sufficent by orders of magnatude.  My phone is a fully up to date pixel 3.  I tried using my wifi in my house first (connected through FIOS), and then using the phone 4G LTE connection.  I can't imagine this was caused by my end.  It seemed like Pearson's servers were jammed at that point in time.

&amp;#x200B;

*Update*: After a LONG time, I did eventually get someone to answer from Pearson.  They were nice enough and were fairly easy to understand, although there was an delay echo introduced where whatever I said was echoed a quarter to half second later which was annoying, but bearable.  I was just happy she was able to hear me.  She said she could open a trouble ticket for me, but as it was well over an hour trying to get through to any human and doubtful it was on my side, I just told her to schedule me for the next available in person appointment.  She had to cancel my appointment and then rebook it as their sub-standard system wouldn't let her reschedule an at home appointment to at a location.  Surprisingly, she said they would refund my money and rebook me.  It was painless enough, but when I asked for a reference number on the refund, all she could do is say I 'should' get an email.  Perhaps unsurprisingly, this morning I see a fully posted charge for the rescheduled exam, but no sign of a refund.  Sigh.  I will give it a few days and then start this process over.

For what its worth, people should IGNORE the advice that the web chat is the fastest way of getting help.  Find the phone number and dial and re-dial it as fast as you can when you get a busy signal.  Despite the fact that it took 20+ minutes to get the number to pickup (and was 'waiting' 20 minutes less from the phones point of view) I got a faster response from someone on the phone.  Web based chat never picked up, even though I left it running during my entire phone conversation.

*Update #2*: It took two more days than the charge, but the refund did show up in the correct amount on my credit card.  I am actually quite surprised.
## [2][Amazon introduces 'Distance Assistant' The company’s latest innovation provides real-time social distancing feedback and we plan to open source the technology By Brad Porter](https://www.reddit.com/r/aws/comments/hdeaqy/amazon_introduces_distance_assistant_the_companys/)
- url: https://www.reddit.com/r/aws/comments/hdeaqy/amazon_introduces_distance_assistant_the_companys/
---
[Amazon introduces 'Distance Assistant'](https://blog.aboutamazon.com/operations/amazon-introduces-distance-assistant)
The company’s latest innovation provides real-time social distancing feedback and Amazon plan to open source the technology
By Brad Porter

Excerpt from blog post

"As we've continued to learn and innovate to support the health and safety of our associates, we also saw an opportunity to evolve our tech even further and promote social distancing behavior in real-time. Given social distancing isn’t always natural, this team set out to use augmented reality to create a magic-mirror-like tool that helps associates see their physical distancing from others. Working backwards from a concept of immediate visual feedback, and inspired by existing examples like radar speed check signs, our 'Distance Assistant' provides employees with live feedback on social distancing via a 50 inch monitor, a camera, and a local computing device. The standalone unit uses machine learning models to differentiate people from their surroundings. Combined with depth sensors, it creates an accurate distance measurement between associates."

Here is [Amazon Distance Assistant in action](https://youtu.be/FexuGnXVkvE).
## [3][Which WAF ACL managed rules are best practice good starting point?](https://www.reddit.com/r/aws/comments/hdpnbn/which_waf_acl_managed_rules_are_best_practice/)
- url: https://www.reddit.com/r/aws/comments/hdpnbn/which_waf_acl_managed_rules_are_best_practice/
---
I have

* RDS
* Beanstalk Application Load Balancer
* VPC
* Node Linux
* Using Route 53
* Guard Duty
* 1 location country only, at least 300-1000 active users monthly
* Single API only

&amp;#x200B;
## [4][aws cli command not working when installed local folder](https://www.reddit.com/r/aws/comments/hds410/aws_cli_command_not_working_when_installed_local/)
- url: https://www.reddit.com/r/aws/comments/hds410/aws_cli_command_not_working_when_installed_local/
---
I installed the [aws-cli](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2-linux.html) program as mentioned in this doc without `sudo` to local folder. But when I try to run the command line utility I get error.

    $ ./aws/install -i awscli-app -b awscli-bin
    $ ls awscli-bin/
     aws  aws_completer
    $ ls -alh awscli-bin
    total 8.0K
    drwxrwxr-x  2 docking_prod docking_prod 4.0K Jun 22 14:14 .
    drwxr-xr-x 16 docking_prod docking_prod 4.0K Jun 22 14:14 ..
    lrwxrwxrwx  1 docking_prod docking_prod   29 Jun 22 14:14 aws -&gt; awscli-app/v2/current/bin/aws
    lrwxrwxrwx  1 docking_prod docking_prod   39 Jun 22 14:14 aws_completer -&gt; awscli-app/v2/current/bin/aws_completer
    $ ./awscli-bin/aws --version
    -bash: ./awscli-bin/aws: No such file or directory

What am I missing here, can anyone help me ?
## [5][How to control Ingress rules for an NLB ?](https://www.reddit.com/r/aws/comments/hdrdpl/how_to_control_ingress_rules_for_an_nlb/)
- url: https://www.reddit.com/r/aws/comments/hdrdpl/how_to_control_ingress_rules_for_an_nlb/
---
Hi,

I have an application composed of an API Gateway, 2 EC2 instance (A and B) and a ECS Cluster (Fargate). Fargate Tasks should be accessed ONLY by the API Gateway throught a VPC Link and an NLB and the EC2 Instance A.

UPDATE : Here is a diagram simplified of the APP : 

&amp;#x200B;

https://preview.redd.it/9d9xk1s4hg651.png?width=552&amp;format=png&amp;auto=webp&amp;s=2bca2f65cffb491c2b031c314ad9e58e2603904f

Now what i want to do is to secure the access to that NLB to make it only accessible from the API Gateway and from the EC2 Instance A.

The problem is that in the NLB there is no Security Group (unlike the ALB)

I already configured the Security Group of Fargate to accept trafic only form the Private IP of the NLB and from the EC2 A.

Now if i try to access directly the Fargate from the EC2 instance B it will not work, and that's what i want. **BUT**, if access the NLB from the EC2 B the traffic will be accepted and redirected to the Fargate Task.

So my question how to control the Ingress rules in the NLB ? Is there a way to make that ?

Thank you,
## [6][Is the connection between datacenters of a different region faster than a typical connection?](https://www.reddit.com/r/aws/comments/hdqo8q/is_the_connection_between_datacenters_of_a/)
- url: https://www.reddit.com/r/aws/comments/hdqo8q/is_the_connection_between_datacenters_of_a/
---
Would someone with a high speed internet connection in Singapore get similar latency and throughput accessing the EU-West-1 region; as to a inter-datacentre connection between Ap-Southeast-1 and EU-West-1.

Is there anything special about the connection between data centers. I'd imagine that Amazon has optimised and invested the fuck out of the connections between them.
## [7][Directory unavailable](https://www.reddit.com/r/aws/comments/hdqjy7/directory_unavailable/)
- url: https://www.reddit.com/r/aws/comments/hdqjy7/directory_unavailable/
---
Is anyone else having users being kicked out of AWS with the  'Directory unavailable' error?

How can I resolve this?
## [8][Elastic Beanstalk - Daphne is not serving static files](https://www.reddit.com/r/aws/comments/hdqcot/elastic_beanstalk_daphne_is_not_serving_static/)
- url: https://www.reddit.com/r/aws/comments/hdqcot/elastic_beanstalk_daphne_is_not_serving_static/
---
Following this guide [https://medium.com/@elspanishgeek/how-to-deploy-django-channels-2-x-on-aws-elastic-beanstalk-8621771d4ff0](https://medium.com/@elspanishgeek/how-to-deploy-django-channels-2-x-on-aws-elastic-beanstalk-8621771d4ff0) to set up a django project that uses websockets. Before I made the traffic go through Daphne, the static files were being served, but now I get on the console that the files are not found.

01\_env.config

        option_settings:  
          aws:elasticbeanstalk:application:environment:
            DJANGO_SETTINGS_MODULE: dashboard.settings
            PYTHONPATH: /opt/python/current/app/dashboard:$PYTHONPATH
          aws:elasticbeanstalk:container:python:
            WSGIPath: dashboard/wsgi.py
          "aws:elasticbeanstalk:container:python:staticfiles":
              /static/: "static/"
        
          aws:elbv2:listener:80:
            ListenerEnabled: 'true'
            Protocol: HTTP
          aws:elbv2:listener:5000:
            ListenerEnabled: 'true'
            Protocol: HTTP

02\_setup.config

        container_commands:
          00_pip_upgrade:
            command: "source /opt/python/run/venv/bin/activate &amp;&amp; pip install --upgrade pip"
            ignoreErrors: false
          01_migrate:
            command: "django-admin.py migrate"
            leader_only: true
          02_collectstatic:
            command: "django-admin.py collectstatic --noinput"
          03_wsgipass:
            command: 'echo "WSGIPassAuthorization On" &gt;&gt; ../wsgi.conf'
          04_celery_tasks:
            command: "cat .ebextensions/celery_configuration.txt &gt; /opt/elasticbeanstalk/hooks/appdeploy/post/run_supervised_celeryd.sh &amp;&amp; chmod 744 /opt/elasticbeanstalk/hooks/appdeploy/post/run_supervised_celeryd.sh"
            leader_only: true
          05_celery_tasks_run:
            command: "/opt/elasticbeanstalk/hooks/appdeploy/post/run_supervised_celeryd.sh"
            leader_only: true

03\_proxy.config

        files:
          "/etc/httpd/conf.d/proxy.conf":
            mode: "000644"
            owner: root
            group: root
            content: |
              ProxyPass /websockets/ ws://127.0.0.1:5000/websockets/
              ProxyPassReverse /websockets/ ws://127.0.0.1:5000/websockets/
              ProxyPass / http://127.0.0.1:5000/
              ProxyPassReverse / http://127.0.0.1:5000/

[settings.py](https://settings.py)

        ...
        
        STATIC_URL = '/static/'
        STATIC_ROOT = os.path.join(BASE_DIR, 'static/')

Running Daphne using:

        daphne -b 0.0.0.0 -p 5000 dashboard.asgi:application
## [9][When SSL is terminated at ALB, how the request to ECS from ALB is still encrypted?](https://www.reddit.com/r/aws/comments/hdh2e6/when_ssl_is_terminated_at_alb_how_the_request_to/)
- url: https://www.reddit.com/r/aws/comments/hdh2e6/when_ssl_is_terminated_at_alb_how_the_request_to/
---
In our infrastructure, a request is initiated from the client and hit the ALB which has an ACM for, say, [testdomain.com](https://testdomain.com). The listener on the ALB redirects the request to port 443 on one of the services running on ECS. 

I am having a hard time understanding that,

&amp;#x200B;

* When SSL is terminated at ALB, is the connection still encrypted between ALB to ECS service?
* How can I verify if the connection between ALB and Backend service is still encrypted if SSL is terminated at ALB?
## [10][Cognito token verification with API gateway](https://www.reddit.com/r/aws/comments/hdjlqp/cognito_token_verification_with_api_gateway/)
- url: https://www.reddit.com/r/aws/comments/hdjlqp/cognito_token_verification_with_api_gateway/
---
I have an authorizer set up in API gateway, it is connected to my Cognito user pool.

It has worked before, but as of recent, when I log in on my app, and get the id token:

userSession.getIdToken().getJWTToken()

I am running my app in debugger mode, and getting the value of this token, however, when I run it through the API gateway authorize test, it claims it is an unauthorized request.

I have no idea why it would be doing this, Also, I have noticed each time I log in with the same username, cognito gives me identical JWT tokens, is this how it should be?

&amp;#x200B;

I have noticed, when I make a new user, it will work for them for a while

Thank you!
## [11][CORS, why does the target domain authorise which source domains can call its API and not the other way round?](https://www.reddit.com/r/aws/comments/hdp7bi/cors_why_does_the_target_domain_authorise_which/)
- url: https://www.reddit.com/r/aws/comments/hdp7bi/cors_why_does_the_target_domain_authorise_which/
---
I've got to the point where I can configure API Gateway CORS so that my scripts can submit to my AWS HTTP API. I have a very basic understanding of cross site attacks. I don't understand why the target domain authorises the source domain and not the other way round or why both domains don't need to authorise each other.
