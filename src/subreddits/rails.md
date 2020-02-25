# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/f2r9mk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/f2r9mk/personal_projects_show_off_your_own_project_andor/
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
## [3][Help understanding Devise](https://www.reddit.com/r/rails/comments/f8zue1/help_understanding_devise/)
- url: https://www.reddit.com/r/rails/comments/f8zue1/help_understanding_devise/
---
Greetings all,

I'm having a bit of confusion on how to work effectively with the Devise gem. I have a Rails app with a React front end. I'm just starting to build it out, and the first thing I want to do is make an Admin page for the app. I created an Admin model, and then created an admins_controller.rb. I updated my routes.rb as follow:

    Rails.application.routes.draw do
      root 'homepage#index'
      devise_for :admins, controllers: { admins: 'admins', sessions: 'sessions' }
      # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
    end

On my front end, the first method I want to hit is to check if admin is logged in, so I created this function:

      getAdminPromise() {
        const url = '/admins/check_for_admin';

        fetch(url)
          .then(response =&gt; {
            if (response.ok) {
              console.log(response)
              return response.json();
            }
            throw new Error("Network response was not ok.");
          })
          .then(response =&gt; this.setState({ currentUser: response }))
          .catch(() =&gt; this.props.history.push("/"));
      }

I ran into "404 not found" and the rails console told me there was no corresponding route for /admins/check_for_admin. So I ran rake routes to see what was available, and I was surprised to see a ton of Devise routes. Awesome! 

After changing the route in my promise to hit one of the prefabricated Devise routes such as '/admins/sign_in'  I was able to hit the endpoint no problem. However, I am confused that there are no apparent controllers for looking at the code, or writing other methods. Do I simply need to rely on the provided Devise methods and just write calls that hit those end points? For instance there doesn't seem to be a route that implies checking if Admin is logged in??
## [4][Wondering about techniques for dynamic routing](https://www.reddit.com/r/rails/comments/f8t4ik/wondering_about_techniques_for_dynamic_routing/)
- url: https://www.reddit.com/r/rails/comments/f8t4ik/wondering_about_techniques_for_dynamic_routing/
---
I have a site I'm working on and each page is a resource in a table of site_items. Thus, all of the relative paths will be something along the lines of **/site_items/4**

I'm wondering if there's a BETTER way to dynamically link these items to prettier URL names.

For instance, In order to coarsely mimic "homepage" functionality, I have a routes_controller that dynamically redirects the root to the first site_item, sorted by the table's :order column:

    def root
        redirect_to site_item_path(SiteItem.order(order: :asc).all.first)
    end

However, this means that instead of the user going to **domain.com/home**, my users are forced to go to **domain.com/site_items/6** for the home page. This is bad SEO and overall just ugly.

I've looked into using slugs briefly, but—from what I can tell so far!—that looks like it'll result in something more like **domain.com/site_items/home**. This still includes much more than I would like to.

I'd love it if I could pop into routes and just put something like...

    get '/signup', to: SiteItem.order(order: :asc).all.first

But it doesn't look like Routes work that way!

Any Ideas?
## [5][Need help switching Ruby version in Rails](https://www.reddit.com/r/rails/comments/f8zds1/need_help_switching_ruby_version_in_rails/)
- url: https://www.reddit.com/r/rails/comments/f8zds1/need_help_switching_ruby_version_in_rails/
---
Just doing it with git bash, I have been following [this guide](https://youtu.be/pPy0GQJLZUM) but run into a problem when I try to make a make a new project:

&gt;An error occurred while installing byebug (11.1.1), and Bundler cannot continue.  
Make sure that \`gem install byebug -v '11.1.1'\` succeeds before bundling.

When I run that I get

&gt;byebug requires Ruby version &gt;= 2.4.0.

It looks like Rails is using Ruby 2.3.3. I have 2.6 installed but I don't know how to change it. I've looked into rvm to fix it but can't get it installed. When I run the first line from this guide, I get 

&gt;03113804BB82D39DC0E3 7D2BAF1CF37B13E2069D6956105BD0E739499BDB  
&gt;  
&gt;gpg: invalid size of lockfile '/c/Users/hsand/.gnupg/pubring.gpg.lock'  
&gt;  
&gt;gpg: cannot read lockfile  
&gt;  
&gt;gpg: can't lock '/c/Users/hsand/.gnupg/pubring.gpg'  
&gt;  
&gt;gpg: key 105BD0E739499BDB: public key not found: General error  
&gt;  
&gt;gpg: error reading '\[stream\]': General error  
&gt;  
&gt;gpg: Total number processed: 0

It seems like for everything I have to install, there are two other things I've had to install first. I don't doubt I've messed something up along the way but I would really appreciate some help!
## [6][cant use .save in model? (using rufus-scheduler)](https://www.reddit.com/r/rails/comments/f8q6nj/cant_use_save_in_model_using_rufusscheduler/)
- url: https://www.reddit.com/r/rails/comments/f8q6nj/cant_use_save_in_model_using_rufusscheduler/
---
hi All!

Im using rufus scheduler and created model/scrape.rb

in the model i added this method:

`def self.track_product_price(product, tracker)`

`puts "Tracker scheduled for product"`

`puts product.inspect`

`puts tracker.inspect`

`@new =` [`Tracker.new`](https://Tracker.new)

`@product = Product.find(`[`product.id`](https://product.id)`)`

`puts @product.inspect`

`x = tracker.qty = 100 #for testing &lt; setting here to 100 to test how to save it in the DB`

[`x.save`](https://x.save) `#for testing`

the problem im facing is that I cant use .save in the method to save the new tracked product price and qty in the database.

The error im getting is

`{ 73200 rufus-scheduler intercepted an error:`

`73200 job:`

`73200 Rufus::Scheduler::EveryJob "10s" {}`

`73200 error:`

`73200 73200`

`73200 NoMethodError`

`\`73200 undefined method \\save' for 100:Integer\`\`\`

&amp;#x200B;

How can i save the data in the model method?

# ## edit 1

I got a step further with this:

Tracker.create(id:1000, qty:1000, price:1000, product\_id:[product.id](https://product.id/) )

(i want to id of the new created tracker to be automatically set, for now set it to 1000 to test.....

But the log is saying:

(0.1ms) BEGIN

↳ app/models/scrape.rb:89:in \\track\_product\_price'\`

Product Load (0.2ms) SELECT "products".\* FROM "products" WHERE "products"."id" = $1 LIMIT $2 \[\["id", 29\], \["LIMIT", 1\]\]

↳ app/models/scrape.rb:89:in \\track\_product\_price'\`

(0.2ms) ROLLBACK

Cant see what is causing the rollback

# ## edit 2 : whoops forgot to set the user ID in the create!

`Tracker.create!(id:1000, qty:1000, price:1000, product_id:` [`product.id`](https://product.id)`, user_id:1 )`

# ## edit 3: sick!

I got it working. But how to automatically increase the trackerID? Can i use something like "trackerid +1"

# ## edit 4: need to figure out how to dynamically increase the tracker id.....

`73380 PG::UniqueViolation: ERROR: duplicate key value violates unique constraint "trackers_pkey"`

`DETAIL: Key (id)=(1) already exists.`

Solved!     `Tracker.create!(id:tracker, qty:1000, price:1000, product_id:` [`product.id`](https://product.id)`, user_id:2 )`

This did the trick... now to remove the hard coded user id....

After that to dynamically save the price and qty.....
## [7][AWS Elastic Beanstalk re-deploys can see the database, but the site is inaccessible to the web](https://www.reddit.com/r/rails/comments/f8lnc5/aws_elastic_beanstalk_redeploys_can_see_the/)
- url: https://www.reddit.com/r/rails/comments/f8lnc5/aws_elastic_beanstalk_redeploys_can_see_the/
---
Sorry if this is a little scattered, but I'm trying to upgrade my active Rails site from Ruby 2.3 to 2.6.

The site is deployed through AWS Elastic Beanstalk, and it's functioned smoothly for about 2 years without any incidents. I recently got an e-mail from AWS saying that they're discontinuing the Ruby 2.3 platform on March 1st. So I need to deploy another version of the site in the same environment using the Ruby 2.6 platform, make sure it works, then do a blue/green deployment to switch URLs and discontinue the old one.

For the last few weeks I've been trying to make another deployment using EB, but I keep running into the same issue. I can deploy the site, EB says the health is fine, and I can even SSH into the instance and verify through running the Rails console that it can see the database and run queries against it with ActiveRecord. However, I can never get the web site to load with the URL it gives me. It just sits there for a long time until it times out. If I specify it to use SSL, my browser will complain about the certificate, but once I click past that, it still just times out.

Has anyone encountered something similar? Could it be that puma isn't starting up? I can't find anything suspicious or fishy in any logs on the instance. In all the searching I've done, I haven't found anyone with a similar problem. There's no obvious error message so it's hard to tell what went wrong where.

To be honest, this has me really stressed out. The site has built up an audience of several hundred users daily, and we have people regularly signing up and contributing. This all comes after about a 3-month digression into Docker to see if I could deploy the site through ECS instead, but I've never been able to get that working either. Oddly enough, I encountered the same problem - inaccessible web server, but I could confirm that the rails console could see the database. The idea that the entire site and community could just die next week because I can't figure this out is pretty frustrating. Are there any sort of freelance consultants or people who are willing to help out Rails sites in their spare time?
## [8][Geared Pagination Gem Behind The Scenes](https://www.reddit.com/r/rails/comments/f8psui/geared_pagination_gem_behind_the_scenes/)
- url: https://www.reddit.com/r/rails/comments/f8psui/geared_pagination_gem_behind_the_scenes/
---
There’s a better, more user-friendly-driven approach to writing  pagination for your application. The idea is simple: Load more records  every time the user is asking to see more.

## Kickoff with GearedPagination::Controller

It all starts with the set\_page\_and\_extract\_portion\_from method which initiates the flow. It receives a set of records, and you  can specify how many records per page you’d like to present to the user.

[Read on Medium...](https://medium.com/squeezerhq/geared-pagination-in-rails-behind-the-scenes-61d9e227540e)
## [9][Upgrading to Rails 5, need help with strong params?](https://www.reddit.com/r/rails/comments/f8fe35/upgrading_to_rails_5_need_help_with_strong_params/)
- url: https://www.reddit.com/r/rails/comments/f8fe35/upgrading_to_rails_5_need_help_with_strong_params/
---
Hello, I'm new to this, but I'm studying and trying to upgrade an old rails app (was on rails 3 so there's attr\_accessible everywhere in every single model) to rails 5. I have all of my specs passing on Rails 5.1, and I am using the "protected\_attributes\_continued" gem.

I know this gem isn't really maintained/endorsed by the Rails team, so I want to convert to "strong params". I am so confused at how to do this, I've already read a bunch of guides.

Half of the models do not even have controllers where I could put the strong params. And the majority of them are heavily nested with belongs\_to, etc.

\*\*\* I have experimented with just deleting every single attr\_accessible in my entire application, removing the protected\_attributes\_continued gem, and removing the whitelist line in my config file....

My specs PASS when I do all of this, and I've tested the application and nothing really seems to be wrong, but I didn't move ANYTHING from attr\_accessible to the controllers.

Is this okay? I'm just confused at how my application still seems to be working fine.
## [10]["VFS Connection does not exist" on Cloud9](https://www.reddit.com/r/rails/comments/f8gjeb/vfs_connection_does_not_exist_on_cloud9/)
- url: https://www.reddit.com/r/rails/comments/f8gjeb/vfs_connection_does_not_exist_on_cloud9/
---
As the title says, when i try to run the $rails server command and preview running application, it tells me "Oops, VFS Connection does not exist". Any way to fix this?
## [11][Exchanging sensitive data between two Rails apps](https://www.reddit.com/r/rails/comments/f8897p/exchanging_sensitive_data_between_two_rails_apps/)
- url: https://www.reddit.com/r/rails/comments/f8897p/exchanging_sensitive_data_between_two_rails_apps/
---
I have two Rails apps and they both use SSL/HTTPS. They have a REST api to exchange data between each other. I assume this data is encrypted because both apps are SSL/HTTPS? And if so, is this secure enough to exchange sensitive data like passwords?
## [12][Is the code for the GitHub web app on Github?](https://www.reddit.com/r/rails/comments/f88rba/is_the_code_for_the_github_web_app_on_github/)
- url: https://www.reddit.com/r/rails/comments/f88rba/is_the_code_for_the_github_web_app_on_github/
---
Hello, This might be a noob question but I"m wondering if the code for GitHub web application rails app is hosted on GitHub as an open-source project?

I'm asking this because I"m interested in how their roles and scopes are stored in the database and how it is being validated for each user. I want to draw inspiration from their implementation as I'm working on a rails app myself.
