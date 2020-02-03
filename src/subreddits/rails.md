# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/evmx0w/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/evmx0w/personal_projects_show_off_your_own_project_andor/
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
## [3][Receive payments for Side Project](https://www.reddit.com/r/rails/comments/ey5q90/receive_payments_for_side_project/)
- url: https://www.reddit.com/r/rails/comments/ey5q90/receive_payments_for_side_project/
---
What would you use in order to receive small payments ($3-$30)  for a side project ?

There are not different types of subscriptions and no complicated pricing schemas. Only pay by month ($3)  or pay annual price ($30).  

Stripe? Paypal? Braintree?  Any other suggestions?

Please keep in mind this is a low maintenance project that I may abandon after one month (If it fails to generate earnings) and I also want to find a way to accept payments quickly in case I build another and another and another project.
## [4][devise_ldap_authenticatable &amp; react_native](https://www.reddit.com/r/rails/comments/exqdrq/devise_ldap_authenticatable_react_native/)
- url: https://www.reddit.com/r/rails/comments/exqdrq/devise_ldap_authenticatable_react_native/
---
I'm adding a React Native App to my monolith rails app. I was wondering, what would be the best approach to dealing with user authorization?  I check the user based on their username/active directory group. How do I handle authenticating users through devise\_ldap\_authenticatable with React Native?
## [5][Nginx + Puma, shows nothing but the root page](https://www.reddit.com/r/rails/comments/expomw/nginx_puma_shows_nothing_but_the_root_page/)
- url: https://www.reddit.com/r/rails/comments/expomw/nginx_puma_shows_nothing_but_the_root_page/
---
First, I followed this :

[https://www.digitalocean.com/community/tutorials/how-to-deploy-a-rails-app-with-puma-and-nginx-on-ubuntu-14-04](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-rails-app-with-puma-and-nginx-on-ubuntu-14-04)

and my nginx configuration and puma.rb matches this (I actually modified them a bit, but important parts are just like this)

and now, when I go to the domain, it:

1. doesn't load the assets (and I wrote a lot of css to fix RTL)
2. doesn't load any route except root.

in both cases, it gives me error 404.

as an update, it only happens on https. http is ok :|
## [6][How to handle CRUD on join table in backend/frontend?](https://www.reddit.com/r/rails/comments/exph1s/how_to_handle_crud_on_join_table_in/)
- url: https://www.reddit.com/r/rails/comments/exph1s/how_to_handle_crud_on_join_table_in/
---
I'm trying to understand how to handle a many\_to\_many relationship on CRUD using Rails/React. 

On CRUD of the has\_many model, do I handle CRUD of the belongs\_to join table in the backend controller? Or do I route the join table to the frontend and handle the CRUD ops there with \`fetch\`?

In the backend, would the controller look something like this on create? 

&amp;#x200B;

    def create
        @table = Table.create(table_params)
        if @table.save
            JoinTable.create(table_id: params.permit!(:table_id), other_table_id: 
        params.permit!(:other_table_id))
          render :json =&gt; @table, status: @ok
        else  
           render :json =&gt; { errors: @table.erros }
        end
    end

or would I handle this in the frontend and treat the join table like the other tables.
## [7][Deploying rails app using Nginx](https://www.reddit.com/r/rails/comments/exnchp/deploying_rails_app_using_nginx/)
- url: https://www.reddit.com/r/rails/comments/exnchp/deploying_rails_app_using_nginx/
---
I just put my project on a server, then I ran its production env.

Now, I configured nginx like this :

           root /home/prp-e/dakhlokharj/public;
    
            # Add index.php to the list if you are using PHP
            index index.html index.htm index.nginx-debian.html;
    
            server_name _;
    
            location / {
                    # First attempt to serve request as file, then
                    # as directory, then fall back to displaying a 404.
                    #try_files $uri $uri/ =404;
                    proxy_pass       https://localhost:3000;
                    proxy_set_header Host      $host;
                    proxy_set_header X-Real-IP $remote_addr;
            }

But when I go to the domain, it still shows me the old static index.

UPDATE: 

It's okay now, but doesn't load assets (I added another location for assets) and also doesn't load other routes, only root! It gave me 404.

P.S : Ruby installed using RVM.
## [8][How to find working gem?](https://www.reddit.com/r/rails/comments/exg1ek/how_to_find_working_gem/)
- url: https://www.reddit.com/r/rails/comments/exg1ek/how_to_find_working_gem/
---
Been in ruby for a month now!  
How do you guys find gem? i am looking for a gem for appointment scheduling, but everything i see is at least 2 years old. How to gauge if it still works. Do you perform unit test for each case?
## [9][Scoped query generating duplicate conditions](https://www.reddit.com/r/rails/comments/ex9tyi/scoped_query_generating_duplicate_conditions/)
- url: https://www.reddit.com/r/rails/comments/ex9tyi/scoped_query_generating_duplicate_conditions/
---
I have a scope which is defined as below in the associations file of the Account model

     has_many :users, -&gt; { where(type: "Requester").order(:id) },
     class_name: 'Requester', dependent: :delete_all  

But when I see the query generated by this from the console I see many similar conditions being duplicated

    SELECT `users`.* FROM `users` WHERE 
    `users`.`type` IN ('Requester') AND 
    `users`.`account_id` = 1 AND 
    `users`.`type` IN ('Requester') AND 
    `users`.`type` = 'Requester'  
    ORDER BY `users`.`type` ASC

I'm not sure why there is multiple conditions being generated which might have an impact on the query performance.

I'm using rails 4.2.1

Could you please tell me how I can debug this or find out how this query generation works?
## [10][Unable to get rubocop to autoformat on save to work with VSCode](https://www.reddit.com/r/rails/comments/ex9yca/unable_to_get_rubocop_to_autoformat_on_save_to/)
- url: https://www.reddit.com/r/rails/comments/ex9yca/unable_to_get_rubocop_to_autoformat_on_save_to/
---
I'm trying to use the rubocop Code extension to format code on saving but I get the famous "An error occured during autocorrection"

This is the settings.json file related to what I have. Yet I keep getting this error. I'm using Macbook Pro if that is required.

    "editor.formatOnPaste": true,
    "editor.formatOnSave": true, 
    "editor.formatOnSaveTimeout": 5000, 
    "diffEditor.renderSideBySide": true,
    "ruby.format": "rubocop", 
    "ruby.rubocop.executePath": "pathto//bin/rubocop/", 
    "[ruby]": { "editor.defaultFormatter": "misogi.ruby-rubocop" }

This is the error message I get.

https://preview.redd.it/39z2f6rezfe41.png?width=1253&amp;format=png&amp;auto=webp&amp;s=3f2b8974ef7ed43e2e1c48f1ef2ae7dcebaba65a

The path to rubocop is correct in the actual settings.json file.

I followed every instruction from [https://stackoverflow.com/questions/48030698/vscode-vscode-ruby-rubocop-how-to-auto-correct-on-save](https://stackoverflow.com/questions/48030698/vscode-vscode-ruby-rubocop-how-to-auto-correct-on-save)

When I try to run it from the command line the rubocop command works just fine. This looks like more of a VS Code error.

Could you please help me resolve this?
## [11][rails 6: place to put a custom.scss file](https://www.reddit.com/r/rails/comments/ex7qic/rails_6_place_to_put_a_customscss_file/)
- url: https://www.reddit.com/r/rails/comments/ex7qic/rails_6_place_to_put_a_customscss_file/
---

If your frontend dev gave you a [custom.scss](https://gist.github.com/oystersauce8/b84abe760127474d5adaefd56e6a015d),
where do you place it in a rails 6 / bootstrap 4
application?

sidenote: I used [this gist](https://gist.github.com/bazzel/2c64e2e5804077f9a61938a93ed54823) when setting up the bootstrap on rails.
## [12][Is it possible to use Rails 6 with SQL Server (MSSQL)?](https://www.reddit.com/r/rails/comments/ewx2q2/is_it_possible_to_use_rails_6_with_sql_server/)
- url: https://www.reddit.com/r/rails/comments/ewx2q2/is_it_possible_to_use_rails_6_with_sql_server/
---
We were hoping to jump to the latest version of rails to rectify a bunch of identified vulnerabilities.  This Rails code is heavily hand edited, so if necessary, switching from tiny_tds would not be unthinkable, but the data we are fetching is, and will always be in an MS SQL database.  Must we use rails 5.x?
