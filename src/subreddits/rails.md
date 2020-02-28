# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/f9t9kq/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/f9t9kq/personal_projects_show_off_your_own_project_andor/
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
## [2][Why does my variable behave differently when I assign using different ActiveRecord search method?](https://www.reddit.com/r/rails/comments/faoadl/why_does_my_variable_behave_differently_when_i/)
- url: https://www.reddit.com/r/rails/comments/faoadl/why_does_my_variable_behave_differently_when_i/
---
So I've noticed some unexpected behaviour when dealing with ActiveRecord.  


First off I assign my instance using 4 different ActiveRecord methods:  


    a = MyModel.first
    
    =&gt; #&lt;MyModel:0x00007f9ed138a4b0
     id: 1,
     name: "foobar",
     rating: 1
    
    
    
    b = MyModel.where(name: "foobar")
    
    =&gt; [#&lt;MyModel:0x00007f9ed706db00
      id: 1,
      name: "foobar",
      rating: 1
    
    
    
    c = MyModel.find(1)
    
    =&gt; #&lt;MyModel:0x00007f9ed7142da0
     id: 1,
     name: "foobar",
     rating: 1
    
    
    
    d = MyModel.where(rating: 1)
    
    =&gt; [#&lt;MyModel:0x00007f9ed1bc5328
      id: 1,
      name: "foobar",
      rating: 1

Now I try returning data from them:  


    a) WILL RETURN ALL DATA
    
    b) WILL ONLY RETURN NAME
        i.e. b.name =&gt; foobar
        b.id =&gt; undefined method `id
    
    c) WILL RETURN ALL DATA
    
    d) WILL ONLY RETURN NAME
        i.e. b.name =&gt; foobar
        b.rating =&gt; undefined method 'rating'

The actual instance seems to have been returned each time and I can see all the data when I just return the variable (e.g. just hitting a, b, c, or d), but I can only seem to access "name" for all of them. Only when assigning using .find(id) and .first will allow me to fully access all data.  


If someone could please explain why this is, that would be very much appreciated??
## [3][Ruby Hash#transform_keys now accepts a hash that maps existing keys to new keys](https://www.reddit.com/r/rails/comments/fae2xz/ruby_hashtransform_keys_now_accepts_a_hash_that/)
- url: https://blog.saeloun.com/2020/02/26/ruby-hash-transform_keys-now-accepts-a-hash-that-maps-existing-keys-to-new-keys
---

## [4][How can I disable migrations in the second database in Rails 6?](https://www.reddit.com/r/rails/comments/fad72g/how_can_i_disable_migrations_in_the_second/)
- url: https://www.reddit.com/r/rails/comments/fad72g/how_can_i_disable_migrations_in_the_second/
---
I'm doing a new project on a database in Rails 6.0.2.1, we're going to use two databases, one of them is the database of the Rails project and the other will be connecting to the database of an ERP we have, I want to completely disable migrations on this second database database (so a `rails db:drop` or a `rails db:reset` doesn't affect our ERP's database) and only use a `schema.rb` so that Rails can know what tables are in this db.

What can I do?
## [5][When developing a big app with different parts, is it smart to split every part into their own app?](https://www.reddit.com/r/rails/comments/faed8e/when_developing_a_big_app_with_different_parts_is/)
- url: https://www.reddit.com/r/rails/comments/faed8e/when_developing_a_big_app_with_different_parts_is/
---
I'm currently developing quite a big application (atleast for my standards). The application basically consists of a user-facing part with a "User" user model and a dashboard part with a "Vendor" user model (there will most likely also be a third part with it's own dashboard-like application in the future). The User part has the most business logic and is way bigger then the Vendor part, however, both parts are operating on and sharing the same data. But the data they share is the only thing they have in common. Everything else in the application, like for example the frontend or the domains, is pretty clear cut into these two parts. 

That's why I'm now wondering if it would be a better idea to really split these two parts into their own apps. So instead of having one big rails app, I would have one for the User side and one for the Vendor side which both share the same database. 

There are three big positives that I see with this approach:

1. **Seperation of concern.** Since they are essentially two different applications that just work on the same underlying data, having them on two different Rails apps should make everything more secure/stable and more easily extentable in the future.

2. **Simpler auth logic.** I was previously thinking about going with devise since both apps have entirely different user models and devise is the only auth gem that can handle multiple user models. However, I'm not a huge fan of devise. Splitting everything into two apps would allow me to choose a simpler auth gem like Clearance which has a way simpler api.

3. **Easily switchable technology.** Going with two apps instead of one big rails app would also allow me two choose and experiment with different technologies more easily. For example, I'm pretty interested in choosing Hanami but wouldn't choose it for the entire stack if everything would be a single app. However, splitting the one really big app into one big and one small app would allow me to choose Hanami for the smaller one and keep Rails for the big one. And if I find Hanami to be the wrong fit somewhere down the road I would only need to convert a small part of the original really big app back.

Does this all make sense and is this a common approach to handling bigger apps with multiple different parts or am I setting myself up for a lot of issues in the future if I go this path?
## [6][Freelancing in Rails](https://www.reddit.com/r/rails/comments/fa7ojc/freelancing_in_rails/)
- url: https://www.reddit.com/r/rails/comments/fa7ojc/freelancing_in_rails/
---
hey.. how are you all?? I hope you all doing fine.. I want advice from experience people.. Anyone here doing freelancing work in ruby on rails?? can you tell me the struggles of finding a rails job in freelancing?? if someone learning rails and want a career of rails in freelancing, what advice would you like to give??
## [7][How to make API calls to Rails server running on Ubuntu Virtual Machine from Windows?](https://www.reddit.com/r/rails/comments/faf1ps/how_to_make_api_calls_to_rails_server_running_on/)
- url: https://www.reddit.com/r/rails/comments/faf1ps/how_to_make_api_calls_to_rails_server_running_on/
---
Hello. So I am developing Rails back-end on my Ubuntu VM and React front-end on my Windows machine. Is there a way to make API calls to that server running on Ubuntu VM? If yes, how to achieve it?
## [8][Why heroku is saying the migration.sh is not present?](https://www.reddit.com/r/rails/comments/faehxi/why_heroku_is_saying_the_migrationsh_is_not/)
- url: https://www.reddit.com/r/rails/comments/faehxi/why_heroku_is_saying_the_migrationsh_is_not/
---
I am writing some github actions to deploy code using docker. I have the below job which kept failing and saying the file is not found. But I see the file is there. What am I missing here?

    deploy-to-heroku:
      runs-on: ubuntu-latest
      if: github.ref != 'refs/heads/master'
      env:
        HEROKU_API_KEY: ${{ secrets.HEROKU_API_KEY }}
      steps:
        - uses: actions/checkout@v2
        - name: Deploy container
          working-directory: backend
          run: |
            HEROKU_APP_NAME=`./scripts/appname.sh ${GITHUB_REF##*/}`
            heroku container:release web -a $HEROKU_APP_NAME
        - name: Run migrations
          working-directory: backend/scripts
          run: |
            HEROKU_APP_NAME=`./appname.sh ${GITHUB_REF##*/}`
            file_contents=`ls -l | cat`
            echo 'current working directory'
            echo $(pwd)
            echo 'file contents..'
            echo $file_contents
            heroku run -s hobby --type=web -a $HEROKU_APP_NAME -- migration.sh
    

error output:

    Run HEROKU_APP_NAME=`./appname.sh ${GITHUB_REF##*/}`
    current working directory
    /home/runner/work/DockingTestApp/DockingTestApp/backend/scripts
    file contents..
    total 20 -rwxr-xr-x 1 runner docker 198 Feb 27 16:12 appname.sh -rwxr-xr-x 1 runner docker 981 Feb 27 16:12 create-heroku-review-app.sh -rwxr-xr-x 1 runner docker 312 Feb 27 16:12 migration.sh -rwxr-xr-x 1 runner docker 772 Feb 27 16:12 review-app-setup.sh -rwxr-xr-x 1 runner docker 990 Feb 27 16:12 setup-review-app-bucket.sh
    Running migration.sh on review-github-actions11a778... starting, web.4853 (Hobby)
    Running migration.sh on review-github-actions11a778... connecting, web.4853 (Hobby)
    Running migration.sh on review-github-actions11a778... up, web.4853 (Hobby)
    /bin/sh: migration.sh: not found

I tried `heroku run -s hobby --type=web -a $HEROKU_APP_NAME -- ./migration.sh` this also, but no luck. What am I missing?
## [9][Looking For Projects to Build a Portfolio](https://www.reddit.com/r/rails/comments/fa6f2q/looking_for_projects_to_build_a_portfolio/)
- url: https://www.reddit.com/r/rails/comments/fa6f2q/looking_for_projects_to_build_a_portfolio/
---
Hi All,

I'm a Business Intelligence Consultant looking to pivot careers.   I'm looking for some projects I could work on that would look good as part of a portfolio.  Any recommendations are greatly appreciated, as well as any advice on getting my foot in the door with a job.
## [10][Would anyone be interested in taking over development of a passion project that is a tool for guitarists and other musicians?](https://www.reddit.com/r/rails/comments/f9y0ge/would_anyone_be_interested_in_taking_over/)
- url: https://www.reddit.com/r/rails/comments/f9y0ge/would_anyone_be_interested_in_taking_over/
---
The site is www.whatkeyamiin.com. It gets 2,000+ visits per month and has room to grow, but I haven't had time to work on it or develop new features in years. I'd love to see it continue on with someone else who wants to help give back to the guitarist community!

It is built on Rails 4 and jQuery, no front end framework (although it would be a good fit for one now).
## [11][Problems installing ActiveAdmin - allows assigning of a non existent parameter?](https://www.reddit.com/r/rails/comments/fa5ues/problems_installing_activeadmin_allows_assigning/)
- url: https://www.reddit.com/r/rails/comments/fa5ues/problems_installing_activeadmin_allows_assigning/
---
Hi all,  


 I'm trying to experiment with Active Admin but having problems with getting set up. I'm at the final step but can't seem to understand what's going on.  


So during the install it gives me the following line in my seeds file:  


    AdminUser.create!(email: 'admin@example.com', password: 'password', password_confirmation: 'password') if Rails.env.development?

This seems to create the AdminUser instance fine, but then when I try to log in, I get an error:

    Invalid Email or password.

So, I then checked the schema, but it appears that a password field doesn't even exist for that model:  


    create_table "admin_users", force: :cascade do |t|
        t.string "email", default: "", null: false
        t.string "encrypted_password", default: "", null: false
        t.string "reset_password_token"
        t.datetime "reset_password_sent_at"
        t.datetime "remember_created_at"
        t.datetime "created_at", null: false
        t.datetime "updated_at", null: false
        t.index ["email"], name: "index_admin_users_on_email", unique: true
        t.index ["reset_password_token"], name: "index_admin_users_on_reset_password_token", unique: true
      end

But for some reason, it still seems to accept the password being set when we create the instance in seeds. No error.  


When I check the model, it looks like this:  


    =&gt; AdminUser(id: integer, email: string, encrypted_password: string, reset_password_token: string, reset_password_sent_at: datetime, remember_created_at: datetime, created_at: datetime, updated_at: datetime)

So I guess my question is twofold:  
1) Why is it allowing me to assign a password when that column seemingly doesn't exist in the database.  
2) How do I get round this issue?  


Thanks.
