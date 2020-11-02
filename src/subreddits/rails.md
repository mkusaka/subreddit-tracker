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
## [2][Rails app working with nginx locally, but redirecting 301 when deployed to EC2](https://www.reddit.com/r/rails/comments/jmlm7x/rails_app_working_with_nginx_locally_but/)
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
## [3][Seo Optimizer gem](https://www.reddit.com/r/rails/comments/jmmmnu/seo_optimizer_gem/)
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
## [4][[RoR + k8] How do I manage to keep the same DB for multiple pods?](https://www.reddit.com/r/rails/comments/jmbc4q/ror_k8_how_do_i_manage_to_keep_the_same_db_for/)
- url: https://www.reddit.com/r/rails/comments/jmbc4q/ror_k8_how_do_i_manage_to_keep_the_same_db_for/
---
I'm trying to run more than one version of a rails project in different containers and some of those versions might have different migration files.

Is there a way to use the same database for all of the pods?

I saw in a presentation someone mentioning a "migration" service that keeps all the running version valid, but I could not find how to achieve this.

Does any of you had a similar problem or some guidance of how to do this?
## [5][Rails API architecture approach for default images](https://www.reddit.com/r/rails/comments/jmc9dk/rails_api_architecture_approach_for_default_images/)
- url: https://www.reddit.com/r/rails/comments/jmc9dk/rails_api_architecture_approach_for_default_images/
---
I am building api that hits a Vue frontend using ActiveStorage to handle images. If a record doesn’t have an image after retrieving from the db, I’d like to use a default image for the record to get passed to the frontend.  I’m unsure if this the right approach.  Should this be the responsibility of the api or the frontend?  The records aren’t required to have an image on create/update so I’m leaning towards the idea the frontend only care and should handle whether a record has one.
## [6][Audio streaming](https://www.reddit.com/r/rails/comments/jm30hd/audio_streaming/)
- url: https://www.reddit.com/r/rails/comments/jm30hd/audio_streaming/
---
Hi everyone, i want to add audio player in my apps, but i can't find simple example. I user rails 6, active storage (save all files in my disk) with standard  views (use jquery). Please give me links or idea how to build this
## [7][How to use AJAX on an devise protected route?](https://www.reddit.com/r/rails/comments/jlzduy/how_to_use_ajax_on_an_devise_protected_route/)
- url: https://www.reddit.com/r/rails/comments/jlzduy/how_to_use_ajax_on_an_devise_protected_route/
---
I want to make some post requests to my app. I’m not using any front end framework. Just straight up rails and erb (well, i guess this isn’t entirely true. I am using stimulus js).

In stimulus, I’m trying to make ajax requests which is working great. I know that I have to send an authenticity token along with my request. I just grab that from the meta tags. However, I don’t know what to send so that devise can authenticate my request with the current user. 

Do I have to send something in the header of the request? Or maybe some other additional parameters? 

```javascript
$.ajax({
    type: 'POST',
    url: url,
    headers: {
        "X-CSRF-Token": CSRF_TOKEN,
    }
}).done(function(data) { 
    alert(data);
});

```

Any help is appreciated. Thanks!
## [8][Bookmarks on Ruby or Rails](https://www.reddit.com/r/rails/comments/jlt3bo/bookmarks_on_ruby_or_rails/)
- url: https://www.reddit.com/r/rails/comments/jlt3bo/bookmarks_on_ruby_or_rails/
---
Hi everyone, I want to make a bookmarks system like on [Myanimelist.net](https://Myanimelist.net), here is a screenshot of what I am talking about [https://prnt.sc/vayxio](https://prnt.sc/vayxio)

Do you know if there is already a gem, plugin or module that has most of functionality covered?
## [9][Simple but Useful Rails Engines](https://www.reddit.com/r/rails/comments/jlkgx6/simple_but_useful_rails_engines/)
- url: https://www.reddit.com/r/rails/comments/jlkgx6/simple_but_useful_rails_engines/
---
I have a series of simple and useful plugins without complex logic to checkout. Wouldn't mind on getting some collaboration and suggestions from other developers.

[https://github.com/phcdevworks](https://github.com/phcdevworks)

[https://rubygems.org/profiles/phcdevworks](https://rubygems.org/profiles/phcdevworks)
## [10][How can I send truly empty body when hitting a API destroy endpoint?](https://www.reddit.com/r/rails/comments/jlyjq5/how_can_i_send_truly_empty_body_when_hitting_a/)
- url: https://www.reddit.com/r/rails/comments/jlyjq5/how_can_i_send_truly_empty_body_when_hitting_a/
---
I am using Rails 5.

I want to send an empty body after I hit the delete end point, but right now probably by rails default code, sending the \`id\` of the deleted resource. But i want it to be empty. What code I need to change here:

    def destroy
      @project.destroy
    
      respond_to do |format|
        format.json { head :no_content }
      end
    end

See the screenshot:

&amp;#x200B;

https://preview.redd.it/17wfwg190lw51.png?width=1844&amp;format=png&amp;auto=webp&amp;s=51c15fad17fb19356e2f7d1e321d681cd4bc20d7
## [11][DOM Parsing of XML file how to approach ?](https://www.reddit.com/r/rails/comments/jlovhj/dom_parsing_of_xml_file_how_to_approach/)
- url: https://www.reddit.com/r/rails/comments/jlovhj/dom_parsing_of_xml_file_how_to_approach/
---
Hi There Experts, 

some weeks ago I asked help from community regarding processing XML  files on RUBY with this post :

[https://www.reddit.com/r/ruby/comments/j6ht5w/is\_there\_a\_gem\_for\_xml\_to\_json\_convertion/](https://www.reddit.com/r/ruby/comments/j6ht5w/is_there_a_gem_for_xml_to_json_convertion/)

After read a lot I discover that what I need is parse a XML file and the DOM Parsing method seems to me more appropriate ( as I still do not have correct understanding of SAX Parsing) ... 

I have checked Rexml but as is not active library with a lot of open issues I am checking for alternatives to accomplish this task ... Rigth now I checking xml\_to\_hash but decided to ask to RoR community if also have a suggestion ... as my requirements is more simplicity than performance as i will not need to work with HUGE files ...

&amp;#x200B;

thanks in advance,
