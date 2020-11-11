# aws
## [1][re:Invent registration is now open](https://www.reddit.com/r/aws/comments/jkenu3/reinvent_registration_is_now_open/)
- url: https://register.virtual.awsevents.com/
---

## [2][Week of Nov 9th - What have you learned recently about AWS?](https://www.reddit.com/r/aws/comments/jqya79/week_of_nov_9th_what_have_you_learned_recently/)
- url: https://www.reddit.com/r/aws/comments/jqya79/week_of_nov_9th_what_have_you_learned_recently/
---
Share your learnings with the community
## [3][AWS Cognito Multi-Region Failover](https://www.reddit.com/r/aws/comments/js6tmv/aws_cognito_multiregion_failover/)
- url: https://www.reddit.com/r/aws/comments/js6tmv/aws_cognito_multiregion_failover/
---
I am trying to build in AWS a platform that covers multiple regions

I will have users signing up in EU and signing up in US

I will use AWS Cognito to handle user auth

&amp;#x200B;

My question is: if I failover a region - how do we migrate users across to the nearest (lowest latency) available region?

&amp;#x200B;

I have a secondary question around S3 too:

I am using route53 latency based routing to route API customer requests to the Lambda/ElasticSearch that is closest to them.  To do the same with S3 too means I must make our bucket public (not acceptable) - or only have the user write to one bucket (and authenticate with Cognito) but then the write latency will be very high

&amp;#x200B;

Thanks for your time!
## [4][Introducing AWS Gateway Load Balancer – Easy Deployment, Scalability, and High Availability for Partner Appliances | Amazon Web Services](https://www.reddit.com/r/aws/comments/jrp7c3/introducing_aws_gateway_load_balancer_easy/)
- url: https://aws.amazon.com/blogs/aws/introducing-aws-gateway-load-balancer-easy-deployment-scalability-and-high-availability-for-partner-appliances/
---

## [5][State of the Art Landing Zone](https://www.reddit.com/r/aws/comments/js6coz/state_of_the_art_landing_zone/)
- url: https://www.reddit.com/r/aws/comments/js6coz/state_of_the_art_landing_zone/
---
We are wondering how big corporations are managing their multi-account environments in AWS? We are currently evaluating Landing Zone and Control Tower, but neither seems to be suitable for us.

Landing Zone seems to be a collection of Cloud Formation Templates and a couple of step functions, and Control Tower is a managed solution offered by AWS. However, both solutions are limited to new environments and don't allow for importing existing architectures (which I guess would be hard to implement).

* Create new environments and import existing environments
* Implement the CIS Benchmarks for AWS
* Create customizable, optimal account structure (e.g., Root, Backup, VPC Management, Security ( IAM ), N Workloads, Shared Services, Audit &amp; Logs, CI / CD, etc. )
* Sane DevSecOps defaults: Cloudtrail, KMS, Macie, Guard Duty, Config, DNS Flow Logs, etc.
* IAM: SSO / Federation, Cross Account Access Management
* SASS Bindings, e.g., Elastic.co VPC Peering

Thinking about this a bit further, if not existing yet, a product for creating and managing Landing Zone in multiple clouds would be cool. Also, to enable Dev Teams to provision accounts in a self-service manner. This could be done with a rather cloud-agnostic technology ( e.g., Terraform ) and an abstraction layer on top of it.

Is anyone aware of such a solution? Also, what are your struggles when working with Landing Zones?
## [6][How to make a POST request to a S3 bucket using a Cloudfront distribution?](https://www.reddit.com/r/aws/comments/js7lry/how_to_make_a_post_request_to_a_s3_bucket_using_a/)
- url: https://www.reddit.com/r/aws/comments/js7lry/how_to_make_a_post_request_to_a_s3_bucket_using_a/
---
I want to limit HTTP requests to my S3 bucket trough a Cloudfront distribution.

I've been able to get this to work to some extend: now GET request will only succeed when making a req. like this: `https://CLOUDFRONT_URL.cloudfront.net/OBJECT_KEY.jpeg`

However, setting up the Cloudfront distribution caused my server (NodeJS/Express/MulterS3) not to be able to make POST requests, i.a. adding objects to the bucket.

How can I enable POST requests from my server?
## [7][Authorization for API Gateway](https://www.reddit.com/r/aws/comments/js1l5v/authorization_for_api_gateway/)
- url: https://www.reddit.com/r/aws/comments/js1l5v/authorization_for_api_gateway/
---
I’m working on a REST API that queries the Cognito User Group and is able to change the users.  I was able to make this through Amplify. 

Amplify sets up the API to only have the Authorization to a specific User Group. I’d like to access the API from the specific User Group and a Lamda. 

How do I set up the authorization for that?
## [8][Restoring an Aurora Database from a Snapshot](https://www.reddit.com/r/aws/comments/js6w6g/restoring_an_aurora_database_from_a_snapshot/)
- url: https://www.reddit.com/r/aws/comments/js6w6g/restoring_an_aurora_database_from_a_snapshot/
---
I had to restore our test database a few weeks ago (developers writing crap code... what can I say?), and it frustrated me a little. Our infrastructure (including the Aurora instance) is in terraform plans, as it should be. However, I can't restore a snapshot to an existing instance of a database, so I renamed the existing instance and tried to re-create what was in the terraform plan by hand. That turned out to be in-exact and time-consuming. I can live with this for test (and maybe stage), but should the need arise - this is not a great process for our production database.

So - how do you handle this? Am I missing a trick?

We're using Aurora/Postgres 10.11.
## [9][Glacier or S3 very slow upload](https://www.reddit.com/r/aws/comments/js6nwf/glacier_or_s3_very_slow_upload/)
- url: https://www.reddit.com/r/aws/comments/js6nwf/glacier_or_s3_very_slow_upload/
---
Hi Guys, just wanted to put this info out there:

I regularly upload files to S3.

Today I created a new bucket to be used with Glacier and the upload speed was 150 Kbps instead of my usual 2Mbps.

Searched threads online and no solutions.

Then I figured it out:

In my usual uploads the program I use (Syncback Pro) compressess the files to zip before uploading.

In this profile I was uploading uncompressed files (eg: mp4 instead of zip)

As soon as I activated compression again, the upload speed returned to normal. Made several tests with and without file compression activated and this was the only factor changing the upload speed.

Glacier or S3 storage class did not make any difference, just the .zip file extension.

This is very weird!
## [10][Getting error while deploying django with celery](https://www.reddit.com/r/aws/comments/js6ie7/getting_error_while_deploying_django_with_celery/)
- url: https://www.reddit.com/r/aws/comments/js6ie7/getting_error_while_deploying_django_with_celery/
---
So, i am trying to deploy my django project. I am using supervisor to invoke celery in the background. Below is the supervisor.conf file:  


    ; Sample supervisor config file.
    ;
    ; For more information on the config file, please see:
    ; http://supervisord.org/configuration.html
    ;
    ; Notes:
    ;  - Shell expansion ("~" or "$HOME") is not supported.  Environment
    ;    variables can be expanded using this syntax: "%(ENV_HOME)s".
    ;  - Quotes around values are not supported, except in the case of
    ;    the environment= options as shown below.
    ;  - Comments must have a leading space: "a=b ;comment" not "a=b;comment".
    ;  - Command will be truncated if it looks like a config file comment, e.g.
    ;    "command=bash -c 'foo ; bar'" will truncate to "command=bash -c 'foo ".
    ;
    ; Warning:
    ;  Paths throughout this example file use /tmp because it is available on most
    ;  systems.  You will likely need to change these to locations more appropriate
    ;  for your system.  Some systems periodically delete older files in /tmp.
    ;  Notably, if the socket file defined in the [unix_http_server] section below
    ;  is deleted, supervisorctl will be unable to connect to supervisord.
    
    [unix_http_server]
    file=/tmp/supervisor.sock   ; the path to the socket file
    ;chmod=0700                 ; socket file mode (default 0700)
    ;chown=nobody:nogroup       ; socket file uid:gid owner
    ;username=user              ; default is no username (open server)
    ;password=123               ; default is no password (open server)
    
    ; Security Warning:
    ;  The inet HTTP server is not enabled by default.  The inet HTTP server is
    ;  enabled by uncommenting the [inet_http_server] section below.  The inet
    ;  HTTP server is intended for use within a trusted environment only.  It
    ;  should only be bound to localhost or only accessible from within an
    ;  isolated, trusted network.  The inet HTTP server does not support any
    ;  form of encryption.  The inet HTTP server does not use authentication
    ;  by default (see the username= and password= options to add authentication).
    ;  Never expose the inet HTTP server to the public internet.
    
    ;[inet_http_server]         ; inet (TCP) server disabled by default
    ;port=127.0.0.1:9001        ; ip_address:port specifier, *:port for all iface
    ;username=user              ; default is no username (open server)
    ;password=123               ; default is no password (open server)
    
    [supervisord]
    logfile=/tmp/supervisord.log ; main log file; default $CWD/supervisord.log
    logfile_maxbytes=50MB        ; max main logfile bytes b4 rotation; default 50MB
    logfile_backups=10           ; # of main logfile backups; 0 means none, default 10
    loglevel=info                ; log level; default info; others: debug,warn,trace
    pidfile=/tmp/supervisord.pid ; supervisord pidfile; default supervisord.pid
    nodaemon=false               ; start in foreground if true; default false
    silent=false                 ; no logs to stdout if true; default false
    minfds=1024                  ; min. avail startup file descriptors; default 1024
    minprocs=200                 ; min. avail process descriptors;default 200
    ;umask=022                   ; process file creation umask; default 022
    ;user=supervisord            ; setuid to this UNIX account at startup; recommended if root
    ;identifier=supervisor       ; supervisord identifier, default is 'supervisor'
    ;directory=/tmp              ; default is not to cd during start
    ;nocleanup=true              ; don't clean up tempfiles at start; default false
    ;childlogdir=/tmp            ; 'AUTO' child log dir, default $TEMP
    ;environment=KEY="value"     ; key value pairs to add to environment
    ;strip_ansi=false            ; strip ansi escape codes in logs; def. false
    
    ; The rpcinterface:supervisor section must remain in the config file for
    ; RPC (supervisorctl/web interface) to work.  Additional interfaces may be
    ; added by defining them in separate [rpcinterface:x] sections.
    
    [rpcinterface:supervisor]
    supervisor.rpcinterface_factory = supervisor.rpcinterface:make_main_rpcinterface
    
    ; The supervisorctl section configures how supervisorctl will connect to
    ; supervisord.  configure it match the settings in either the unix_http_server
    ; or inet_http_server section.
    
    [program:celeryd]
    command=/var/app/venv/staging-LQM1lest/bin/celery -A main worker --loglevel=info
    stdout_logfile=/mountpoint/packages/celeryd.log
    stderr_logfile=/mountpoint/packages/celeryd.log
    autostart=true
    autorestart=true
    startsecs=10
    stopwaitsecs=600
    
    [program:celerybeat]
    command=/var/app/venv/staging-LQM1lest/bin/celery -A main beat --pidfile="/tmp/celerybeat.pid"
    stdout_logfile=/mountpoint/packages/celerybeat.log
    stderr_logfile=/mountpoint/packages/celerybeat.log
    autostart=true
    autorestart=true
    startsecs=10
    stopwaitsecs=600
    
    [supervisorctl]
    serverurl=unix:///tmp/supervisor.sock ; use a unix:// URL  for a unix socket
    ;serverurl=http://127.0.0.1:9001 ; use an http:// url to specify an inet socket
    ;username=chris              ; should be same as in [*_http_server] if set
    ;password=123                ; should be same as in [*_http_server] if set
    ;prompt=mysupervisor         ; cmd line prompt (default "supervisor")
    ;history_file=~/.sc_history  ; use readline history if available
    
    ; The sample program section below shows all possible program subsection values.
    ; Create one or more 'real' program: sections to be able to control them under
    ; supervisor.
    
    
    ;[include]
    ;files = relative/directory/*.ini
    

and this is my ".ebextensions/build.config" file commands the last command is to invoke the supervisord

    commands:
        command1:
            command: "test -f /usr/bin/google-chrome || wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm"
        command2:
            command: "test -f /usr/bin/google-chrome || sudo yum -y install ./google-chrome-stable_current_x86_64.rpm"
        command3:
            command: "sudo cp -r /home/ec2-user/nltk_data /usr/share/nltk_data"
        command4:
            command: "test -f /usr/bin/supervisord || sudo amazon-linux-extras install -y epel"
        command5:
            command: "test -f /usr/bin/supervisord || sudo yum install -y supervisor"
        command6:
            command: "test -e /tmp/supervisor.sock &amp;&amp; sudo unlink /tmp/supervisor.sock"
        command7:
            command: "/usr/bin/supervisord -c /var/app/current/supervisord.conf"
            cwd: "/var/app/current" 

My celery worker is working but in the celerybeat log i get this error:

    shell-init: error retrieving current directory: getcwd: cannot access parent directories: No such file or directory
    Traceback (most recent call last):
      File "/var/app/venv/staging-LQM1lest/bin/celery", line 8, in &lt;module&gt;
        sys.exit(main())
      File "/var/app/venv/staging-LQM1lest/lib/python3.7/site-packages/celery/__main__.py", line 16, in main
        _main()
      File "/var/app/venv/staging-LQM1lest/lib/python3.7/site-packages/celery/bin/celery.py", line 322, in main
        cmd.execute_from_commandline(argv)
      File "/var/app/venv/staging-LQM1lest/lib/python3.7/site-packages/celery/bin/celery.py", line 499, in execute_from_commandline
        super(CeleryCommand, self).execute_from_commandline(argv)))
      File "/var/app/venv/staging-LQM1lest/lib/python3.7/site-packages/celery/bin/base.py", line 289, in execute_from_commandline
        argv = self.setup_app_from_commandline(argv)
      File "/var/app/venv/staging-LQM1lest/lib/python3.7/site-packages/celery/bin/base.py", line 509, in setup_app_from_commandline
        self.app = self.find_app(app)
      File "/var/app/venv/staging-LQM1lest/lib/python3.7/site-packages/celery/bin/base.py", line 531, in find_app
        return find_app(app, symbol_by_name=self.symbol_by_name)
      File "/var/app/venv/staging-LQM1lest/lib/python3.7/site-packages/celery/app/utils.py", line 373, in find_app
        sym = symbol_by_name(app, imp=imp)
      File "/var/app/venv/staging-LQM1lest/lib/python3.7/site-packages/celery/bin/base.py", line 534, in symbol_by_name
        return imports.symbol_by_name(name, imp=imp)
      File "/var/app/venv/staging-LQM1lest/lib/python3.7/site-packages/kombu/utils/imports.py", line 57, in symbol_by_name
        module = imp(module_name, package=package, **kwargs)
      File "/var/app/venv/staging-LQM1lest/lib/python3.7/site-packages/celery/utils/imports.py", line 110, in import_from_cwd
        with cwd_in_path():
      File "/usr/lib64/python3.7/contextlib.py", line 112, in __enter__
        return next(self.gen)
      File "/var/app/venv/staging-LQM1lest/lib/python3.7/site-packages/celery/utils/imports.py", line 61, in cwd_in_path
        cwd = os.getcwd()
    FileNotFoundError: [Errno 2] No such file or directory

Could really use a help here.
## [11][Multiple stacks on AWS](https://www.reddit.com/r/aws/comments/js67np/multiple_stacks_on_aws/)
- url: https://www.reddit.com/r/aws/comments/js67np/multiple_stacks_on_aws/
---
Hello, I’m building a web app, and am primarily a php developper. What I’d like to do is have 
- information / brochure part of the site on wordpress
- the admin of the app on LAMP stack 
- the user facing side of the app on node to have a react native component

I’m completely new to aws and am struggling with the documentation. 

Is it possible to configure aws to point to different systems depending on the url?

Eg website.com to the wordpress, website.com/login to the admin, and website.com/app to the node.js?
## [12][Static and dynamic compliance scans (Terraform + live AWS account)](https://www.reddit.com/r/aws/comments/js4j20/static_and_dynamic_compliance_scans_terraform/)
- url: https://www.reddit.com/r/aws/comments/js4j20/static_and_dynamic_compliance_scans_terraform/
---
Hi,

I'm looking for a solution (preferably open-source, even if the integration takes more time) allowing to write **custom compliance rules for AWS which can be reused between static and dynamic analysis**.

The ideal workflow would be:

* Write rules such as "No public S3 bucket", "No EC2 instance with a public IP", etc.
* Run static analysis scans against Terraform plans
* Run dynamic analysis scans against a live AWS account

I believe Snyk and Prisma Cloud have these features, but they cost $$$. I know about several tools capable of performing this kind of static analysis (though few allow to write *custom* rules, tfsec doesn't) and dynamic analysis (AWS Config, CloudCustodian...), but having to write twice the same rule and properly test it sounds sub-optimal at best.

Anything you'd recommend? Thank you!
