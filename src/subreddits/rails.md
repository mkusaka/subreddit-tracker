# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/
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
## [3][SSO server in rails, is there any gems?](https://www.reddit.com/r/rails/comments/i8vq7h/sso_server_in_rails_is_there_any_gems/)
- url: https://www.reddit.com/r/rails/comments/i8vq7h/sso_server_in_rails_is_there_any_gems/
---
Actually, I'm looking for a way to make my app the SSO server and users will be able to login to other services I made with a single account.
## [4][Looking for Rails tutor/mentor](https://www.reddit.com/r/rails/comments/i8lzk7/looking_for_rails_tutormentor/)
- url: https://www.reddit.com/r/rails/comments/i8lzk7/looking_for_rails_tutormentor/
---
I'm currently building an app for a friend, which I'd like to eventually turn into a multi-tenant SaaS app. I'm self-taught through online tutorials, so I've been able to make decent progress. However, I'm starting to hit a wall with certain features.

I plan to tackle these features over the next few weeks: multi-tenancy, setting up robust 3rd party API integrations (Stripe, Mailgun), background jobs, and deployment beyond Heroku.

I'm looking for someone that is able to meet once a week via Zoom and do short PR reviews async during the week.

Happy to pay a reasonable weekly or monthly fee.
## [5][Question about text encryption in Rails](https://www.reddit.com/r/rails/comments/i8zp1u/question_about_text_encryption_in_rails/)
- url: https://www.reddit.com/r/rails/comments/i8zp1u/question_about_text_encryption_in_rails/
---
I want to encrypt the content of a column in a model in Rails. Here's the scenario:

* There are Users (devise) and Posts. Each User can have many Posts
* Each Post has a title and a content column 
* When someone logs in to the application, then he has access to all posts/content 
* The content is of course a column of type text.

I want to limit access to certain Post content, leaving the title intact for search purposes. I want the app to prompt the user for a password (when he enables somehow this feature) and If he enters the right one, show the text. 

One silly idea is to redirect the user to an Unlock this Post page where there could be a single text\_field form that posts to a controller action. The verification of the right passcode is done and If legit, he will be redirected to the show decrypted Post page. 

I don't want to use a gem for this (lockbox or attr\_encrypted). The text can be encrypted with a before\_save callback in the Post model. 

Have you ever ran into a similar issue and If yes, what did you do ?
## [6][Can anyone help me find the feature I'm thinking of?](https://www.reddit.com/r/rails/comments/i8lspn/can_anyone_help_me_find_the_feature_im_thinking_of/)
- url: https://www.reddit.com/r/rails/comments/i8lspn/can_anyone_help_me_find_the_feature_im_thinking_of/
---
Hello! A few months back I swear saw a new Rails feature about auto-expiring keys, However, I can't for the life of me find any reference to it anymore.

It seemed to be intended for things like password resets, were you might want to set a value that only lasts in the database for 15 minutes.

The code looked something like this:

```ruby
user = User.first
user.update(password_reset_token: "abcdefg", expiry: 15.minutes) 
```

Does anyone know of the feature I'm thinking of, and what it's actually called?

Edit: It's called `signed_ids`. https://blog.saeloun.com/2020/05/20/rails-6-1-adds-support-for-signed-ids-to-active-record.html
## [7][Is it possible to configure an ActiveRecord model to read and write from different database tables?](https://www.reddit.com/r/rails/comments/i8mv31/is_it_possible_to_configure_an_activerecord_model/)
- url: https://www.reddit.com/r/rails/comments/i8mv31/is_it_possible_to_configure_an_activerecord_model/
---
I don't mean splitting reads/writes on different databases (e.g. master/replica) and this isn't an STI situation. I'm wondering if there's a way to configure an ActiveRecord model to write to table `A` but read from table `B`. Specifically, table `B` is a readonly postgresql view with a schema that is a superset of table `A`.

I have a model that has certain attributes that are written to the DB, but it also has calculated attributes implemented as instance methods of the Ruby model. Instead of calculating these attributes in Ruby, I want to push these calculations down and implement them in a postgres SQL view instead, primarily so these values can be accessed by other SQL queries/reports.

Simple example:

I have a `products` table with two columns, `price` and `weight`.

    create_table "products" do |t|
      t.decimal "price"
      t.decimal "weight"
    end

In the Ruby model, I have a calculated property `price_per_lb`.

    class Product &lt; ApplicationRecord
      def price_per_lb
        price / weight
      end
    end

I want to move `price_per_lb` to be calculated in SQL instead of Ruby. Using the Scenic gem I would create a view called something like `product_info` that would look like this:

    /* product_info view definition */
    SELECT *,
           price / weight AS price_per_lb
      FROM products;

Then I would remove the `price_per_lb` method from the `Product` model.

For this to work, it seems there needs to be a way to configure the `Product` model to write to the `products` table but read from the `product_info` view so that the `price_per_lb` property is available to the Ruby model.

    class Product &lt; ApplicationRecord
      read_table :product_info
      write_table :products
    end

Is something like this possible? Is there an alternative way of implementing something similar?

Thanks!

Edit: typo
## [8][Creating multiple objects through one form](https://www.reddit.com/r/rails/comments/i8dl92/creating_multiple_objects_through_one_form/)
- url: https://www.reddit.com/r/rails/comments/i8dl92/creating_multiple_objects_through_one_form/
---
I am a beginner rails developer and I am stuck with what seems to be a very specific problem.

I have three tables: User, UserApp, and App.

When a user logs in I want them to have input for 20 apps that they use (netflix, amazon, chase etc.) and each record to be saved in UserApp as you can imagine. I know the simpler way is to just use MongoDB but the client I am working for insisted on using a relational database. 

How do I create such a form where the user can input multiple apps (if the app doesn't already exist on the database then a new app should be created)? Are there any articles or prior posts that I can refer to? Thank you for your help
## [9]["Most purposeful code I've ever written" - Recruiting Volunteer Developers for Mental Health Project](https://www.reddit.com/r/rails/comments/i8f3ku/most_purposeful_code_ive_ever_written_recruiting/)
- url: https://www.reddit.com/r/rails/comments/i8f3ku/most_purposeful_code_ive_ever_written_recruiting/
---
Hello,

I'm a product owner for HeartSupport, a mental health non profit.

* We are recruiting volunteer developers
* 6-week cycle from 9/1 through 10/15 
* Help scale mental health support to every platform on the internet

Code base is in Rails, we have an MVP and a team of 3 Rails devs.

[Please visit this page to apply!](https://heartsupport.com/devteam)

\-Nate
## [10][I curated all the remote job openings from Hacker News who is hiring thread - August](https://www.reddit.com/r/rails/comments/i7pl8b/i_curated_all_the_remote_job_openings_from_hacker/)
- url: https://www.reddit.com/r/rails/comments/i7pl8b/i_curated_all_the_remote_job_openings_from_hacker/
---
This list contains 398 remote jobs and you can filter them by location or skills.

Here I would like to share the entire remote jobs list from the big list of opportunities. All these are 100% remote jobs not just allowed to work from home during this crisis. These are 100% remote jobs and will continue to follow that after the crisis.

https://remoteleaf.com/whoishiring.   

Note: Select "ruby" in the skills filter to see ruby/rails dev jobs.

- 100% remote full-time jobs
- Spent more than 14 hours to curate this information
## [11][Deploying Rails 6, React, Ruby-2.7.1 via Elastic Beanstalk. Unable to install a later version of nodejs via ebextensions to enable assets:precompile](https://www.reddit.com/r/rails/comments/i7t0js/deploying_rails_6_react_ruby271_via_elastic/)
- url: https://www.reddit.com/r/rails/comments/i7t0js/deploying_rails_6_react_ruby271_via_elastic/
---
Webpacker seems to want a nodejs greater then version 8 so I've tried installing it via ebextensions. 

In the application's configuration:

    RAILS_SKIP_ASSET_COMPILATION true

I've stripped down the ebextension's file to jus tonstall nodejs...

    commands:
      01_node_get:
        cwd: /tmp
        command: 'sudo curl --silent --location https://rpm.nodesource.com/setup_13.x | sudo bash -'

      02_node_install:
        cwd: /tmp
        command: 'sudo yum -y install nodejs'

I have even tried the commands on the AWS instance but it still insists on installing nodejs 6.17

    % sudo yum remove nodejs
    % sudo yum clean all
    % sudo rm /var/cache/yum
    % cd /tmp
    % sudo curl --silent --location https://rpm.nodesource.com/setup_13.x | sudo bash -
    % sudo yum install nodejs

What am I doing wrong. How can I get round this? Any pointers greatly appreciated.
## [12][Limiting Users from AD with Devise](https://www.reddit.com/r/rails/comments/i7sptq/limiting_users_from_ad_with_devise/)
- url: https://www.reddit.com/r/rails/comments/i7sptq/limiting_users_from_ad_with_devise/
---
I have been looking to install Devise into my app to a limited access namespace, but as I look at Devise, I am not seeing if there is a way to limit the users that able to log in. I work for a larger company with a more than substantial AD. I really only want 4-6 people to be able to log in out of the entire AD. Am I approaching this the best way? Is there a better option? Is there an easy way to accomplish this within Devise?
