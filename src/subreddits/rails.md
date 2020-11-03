# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/
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
## [2][adding CSS to a single view](https://www.reddit.com/r/rails/comments/jn9nfy/adding_css_to_a_single_view/)
- url: https://www.reddit.com/r/rails/comments/jn9nfy/adding_css_to_a_single_view/
---
I'm working on a rails 6 application ( i generated a controller with his view (index/show/edit) ) 

but on the show page i have to add a show some images so i'm going to use swiperJS  (like this exemple not mine btw) https://codepad.co/snippet/swiper-thumbs-gallery-with-two-way-control

so i need to add the JS and CSS to the show page only 

a did a:

    yarn add swiper 

and to the javascript folder i added a new file called GallerieSwiper.js which contant my JS 

and in my show.html.erb page i added:

    &lt;%= javascript_pack_tag 'GallerieSwiper' %&gt;


the issue is how to add the swiper.min.css to only the show view ? 

i can't added in the top of the page as it is a subview of the application.html.erb that contant the whole website structure
## [3][The best Rails developer that no one ever was](https://www.reddit.com/r/rails/comments/jmxfci/the_best_rails_developer_that_no_one_ever_was/)
- url: https://www.reddit.com/r/rails/comments/jmxfci/the_best_rails_developer_that_no_one_ever_was/
---
What resources helped you become a better rails developer? Or projects?
## [4][Have you implemented Elasticsearch as an indexing solution?](https://www.reddit.com/r/rails/comments/jn94yh/have_you_implemented_elasticsearch_as_an_indexing/)
- url: https://www.reddit.com/r/rails/comments/jn94yh/have_you_implemented_elasticsearch_as_an_indexing/
---
Hi All,

Hope everyone is staying safe.

I'm currently hiring for a Senior Ruby Engineer who has implemented Elasticsearch as an indexing solution. We’re an AI startup with offices in London and Berlin who have grown to 55 people in the last 5 years. We’re working on cutting edge AI technology, supported by a small team of machine learning experts and a collaborative product team that enables our clients to make data driven decisions using 100k+ customer feedback to improve their product.

As a Senior Ruby Engineer at Chattermill, you’ll be working on greenfield products as well as interesting performance optimisation problems. You will be involved in all steps of the project including design, architecture, implementation, automated tests development and monitoring the health of the systems in production. This role can be fully remote (within EU timezone) or be based in our London or Berlin office (when it's safe to open up again).

This opportunity would suit: 

* Someone who values autonomy and would relish working on new challenges
* Has a growth mindset and wants to work with likeminded people
* Wants to work on complex scaling problems
* Wants to work in a high impact role within a startup but at the same time has a good work/ life balance

What we’re after:

* Someone with good knowledge of software design and architecture 
* Proficient with PostgreSQL and has experience with large datasets.
* Experience with Elasticsearch as an indexing solution
* Someone who values working within a close knit team and is willing to help and support others 

You can apply via: [https://grnh.se/a87508f63us](https://grnh.se/a87508f63us) or contact me directly if you have any questions via [jason@chattermill.io](mailto:jason@chattermill.io).
## [5][How to fix Rendered ActiveModel::Serializer::Null with Hash](https://www.reddit.com/r/rails/comments/jn0wim/how_to_fix_rendered_activemodelserializernull/)
- url: https://www.reddit.com/r/rails/comments/jn0wim/how_to_fix_rendered_activemodelserializernull/
---
Hi guys, I'm using Active Model Serializer and I'm having trouble updating a user. I'm not sure what to do to solve this problem? 

I get this error message.   

`app/controllers/registrations_controller.rb:19
[active_model_serializers] Rendered ActiveModel::Serializer::Null with Hash (0.09ms)` 

My registrations_controller is:

	`def update
		@user = User.find(params[:id])
		if @user.assign_attributes(registration_params)
			render json: @user
		else
			render json: { status: :bad_request }
		end	
	end`	


Here is my User Serializer:

`
class UserSerializer &lt; ActiveModel::Serializer
  include Rails.application.routes.url_helpers
  attributes  :id, :first_name, :last_name, :email, :photo_url, :login_status
  
  attributes :photo_url
  def photo_url
    variant = object.photo.variant(resize: "80x80")
    return rails_representation_url(variant, only_path: true)
  end  

  def login_status
    {
      status: :created,
      logged_in: true
    }
  end   
end`
## [6][AssetsNotPrecompiled not resolved with :clobber and :precompile](https://www.reddit.com/r/rails/comments/jn6qja/assetsnotprecompiled_not_resolved_with_clobber/)
- url: https://www.reddit.com/r/rails/comments/jn6qja/assetsnotprecompiled_not_resolved_with_clobber/
---
Hi all,

I am the creator of [https://github.com/chrisvel/wreeto\_official](https://github.com/chrisvel/wreeto_official).

There's an issue pending for a couple of days, but I am unable to save.

[https://github.com/chrisvel/wreeto\_official/issues/47](https://github.com/chrisvel/wreeto_official/issues/47)

Heads up:

`AssetsNotPrecompiled` appeared after a `docker-compose pull` of latest image and doesn't seem to go away with `rake assets:clobber` and `rake assets:precompile`, even If the public assets folder is removed manually with `rm -rf`.

Does anyone have any idea that might help ?
## [7][Seo Optimizer gem](https://www.reddit.com/r/rails/comments/jmmmnu/seo_optimizer_gem/)
- url: https://www.reddit.com/r/rails/comments/jmmmnu/seo_optimizer_gem/
---
Hey,

I wrote a gem to manage sitemap, robots.txt, error pages, etc.

Things to optimize SEO, you know \^\^ ...

It uses the "sitemap\_generator" gem. 

I would like to have your opinions, comments and contributions on this. 

The first goal is to easily and efficiently manage SEO on a Rails application. 

You can find the source code here: [https://github.com/RonanLOUARN/seo\_optimizer](https://github.com/RonanLOUARN/seo_optimizer)

Have a great day!
## [8][DateRangePciker and Jquery: Uncaught TypeError: $(…).daterangepicker is not a function](https://www.reddit.com/r/rails/comments/jmxpc2/daterangepciker_and_jquery_uncaught_typeerror/)
- url: https://www.reddit.com/r/rails/comments/jmxpc2/daterangepciker_and_jquery_uncaught_typeerror/
---
Hey guys, having a really puzzling time with everyones favorite daterangepicker. I had it working completely fine, and then all of a sudden the jquery has gone haywire on me! Im hoping someone ehre has experience with this issue.

Ive started getting the 

    Uncaught TypeError: $(...).daterangepicker is not a function

error in my console, and the page breaks. I havent changed my application.js, gemfile, nor stylesheets since it was working, hoping to get some assistance here. From doing some research it seems that the function is being loaded before jquery is, but I am mystified as to why this is happening.

Here are my files:

Stack Overflow Post: [https://stackoverflow.com/questions/64654020/uncaught-typeerror-daterangepicker-is-not-a-function-rails](https://stackoverflow.com/questions/64654020/uncaught-typeerror-daterangepicker-is-not-a-function-rails)
## [9][Incomplete response received from application](https://www.reddit.com/r/rails/comments/jmunb9/incomplete_response_received_from_application/)
- url: https://www.reddit.com/r/rails/comments/jmunb9/incomplete_response_received_from_application/
---
I am deploying a new rails 6 app to a linux server. I have never seen that message before but haven't been able to figure out a way around it. I was thinking originally it was something to do with my credentials. I ran rails credentials:edit --environment production and it generated the files.

Here is where I am getting confused. Is the string of characters I see in conig/credentials/production.yml.enc the secret key or is that just the contents of the file encrypted?

Do I need a secret key base in that file?

On the linux server I have the /home/deploy/myapp/.rbenv-vars file with my DB info and the secret\_key\_base. I'm not sure I am putting the right value in this. What exactly goes there?
## [10][Am I being too careful?](https://www.reddit.com/r/rails/comments/jmr89x/am_i_being_too_careful/)
- url: https://www.reddit.com/r/rails/comments/jmr89x/am_i_being_too_careful/
---
I have a User Model that has Organizations.  Each User will have two specific organizations (x and y) that are not to be modified in any way; Users can create more Organizations.  Would it be enough to just not show options for deletion/modification on the front-end or would I need to prevent deletion on the back-end?
For example, show deletion/modification buttons on the front-end for all Organizations except x and y organization.  Also on the back-end, make sure x and y can not be modified in any way in my Organization controller.
## [11][Rails app working with nginx locally, but redirecting 301 when deployed to EC2](https://www.reddit.com/r/rails/comments/jmlm7x/rails_app_working_with_nginx_locally_but/)
- url: https://www.reddit.com/r/rails/comments/jmlm7x/rails_app_working_with_nginx_locally_but/
---
I have a dockerized Rails app, and the nginx container with some custom config on runtime that I have setup. The setup works perfectly fine on my local machine.


I fire up docker-compose and the nginx container pointing to the app network, everything works like a charm. I'm able to access the app at localhost instead of localhost:8000.

However, I'm currently deploying this on AWS ECS and there's a loadbalancer with HTTPS involved here.


The app gets deployed and both containers are running fine, but urls that I hit are either returning a 301 Moved Permanently.

Here's my default.conf


    upstream PLACEHOLDER_BACKEND_NAME {
    server PLACEHOLDER_BACKEND_NAME:PLACEHOLDER_BACKEND_PORT;
    }
    
    server {
    listen 80;
    server_name www.PLACEHOLDER_VHOST;
    return 301 https://$host$request_uri;
    }
    
    server {
    listen 80 default deferred;
    server_name PLACEHOLDER_VHOST;
    
    root /PLACEHOLDER_BACKEND_NAME/public;
    
    # level due to Rail's asset pipeline.
    location ~ ^/assets/ {
    gzip_static on;
    
    expires max;
    add_header Cache-Control public;
    add_header Last-Modified "";
    add_header ETag "";
    }
    
    
    
    location ~ /\. {
    return 404;
    access_log off;
    log_not_found off;
    }
    
    try_files $uri $uri/index.html $uri.html u/PLACEHOLDER_BACKEND_NAME;
    
    location = /favicon.ico {
    try_files /favicon.ico = 204;
    access_log off;
    log_not_found off;
    }
    
    
    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains;";
    
    
    location u/PLACEHOLDER_BACKEND_NAME {
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header Host $http_host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_redirect off;
    
    if ($http_x_forwarded_proto = "http") {
    return 301 https://$host$request_uri;
    }
    proxy_pass http://PLACEHOLDER_BACKEND_NAME;
    }
    }

&amp;#x200B;

and here's my script that I'm running when the container starts :

&amp;#x200B;

    #!/usr/bin/env bash
    
    set -e
    
    PLACEHOLDER_BACKEND_NAME="my-api"
    PLACEHOLDER_BACKEND_PORT="8000"
    
    
    PLACEHOLDER_VHOST="$(curl http://169.254.169.254/latest/meta-data/public-hostname)"
    
    
    DEFAULT_CONFIG_PATH="/etc/nginx/conf.d/default.conf"
    
    sed -i "s/PLACEHOLDER_VHOST/${PLACEHOLDER_VHOST}/g" "${DEFAULT_CONFIG_PATH}"
    sed -i "s/PLACEHOLDER_BACKEND_NAME/${PLACEHOLDER_BACKEND_NAME}/g" "${DEFAULT_CONFIG_PATH}"
    sed -i "s/PLACEHOLDER_BACKEND_PORT/${PLACEHOLDER_BACKEND_PORT}/g" "${DEFAULT_CONFIG_PATH}"
    
    # Execute the CMD from the Dockerfile and pass in all of its arguments.
    exec "$@"

&amp;#x200B;

I even tried setting PLACEHOLDER\_VHOST in the script to my ELB public DNS but to no avail. What could be the issue here?
