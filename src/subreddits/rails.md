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
## [3][cant use .save in model? (using rufus-scheduler)](https://www.reddit.com/r/rails/comments/f8q6nj/cant_use_save_in_model_using_rufusscheduler/)
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
## [4][AWS Elastic Beanstalk re-deploys can see the database, but the site is inaccessible to the web](https://www.reddit.com/r/rails/comments/f8lnc5/aws_elastic_beanstalk_redeploys_can_see_the/)
- url: https://www.reddit.com/r/rails/comments/f8lnc5/aws_elastic_beanstalk_redeploys_can_see_the/
---
Sorry if this is a little scattered, but I'm trying to upgrade my active Rails site from Ruby 2.3 to 2.6.

The site is deployed through AWS Elastic Beanstalk, and it's functioned smoothly for about 2 years without any incidents. I recently got an e-mail from AWS saying that they're discontinuing the Ruby 2.3 platform on March 1st. So I need to deploy another version of the site in the same environment using the Ruby 2.6 platform, make sure it works, then do a blue/green deployment to switch URLs and discontinue the old one.

For the last few weeks I've been trying to make another deployment using EB, but I keep running into the same issue. I can deploy the site, EB says the health is fine, and I can even SSH into the instance and verify through running the Rails console that it can see the database and run queries against it with ActiveRecord. However, I can never get the web site to load with the URL it gives me. It just sits there for a long time until it times out. If I specify it to use SSL, my browser will complain about the certificate, but once I click past that, it still just times out.

Has anyone encountered something similar? Could it be that puma isn't starting up? I can't find anything suspicious or fishy in any logs on the instance. In all the searching I've done, I haven't found anyone with a similar problem. There's no obvious error message so it's hard to tell what went wrong where.

To be honest, this has me really stressed out. The site has built up an audience of several hundred users daily, and we have people regularly signing up and contributing. This all comes after about a 3-month digression into Docker to see if I could deploy the site through ECS instead, but I've never been able to get that working either. Oddly enough, I encountered the same problem - inaccessible web server, but I could confirm that the rails console could see the database. The idea that the entire site and community could just die next week because I can't figure this out is pretty frustrating. Are there any sort of freelance consultants or people who are willing to help out Rails sites in their spare time?
## [5][Geared Pagination Gem Behind The Scenes](https://www.reddit.com/r/rails/comments/f8psui/geared_pagination_gem_behind_the_scenes/)
- url: https://www.reddit.com/r/rails/comments/f8psui/geared_pagination_gem_behind_the_scenes/
---
There’s a better, more user-friendly-driven approach to writing  pagination for your application. The idea is simple: Load more records  every time the user is asking to see more.

## Kickoff with GearedPagination::Controller

It all starts with the set\_page\_and\_extract\_portion\_from method which initiates the flow. It receives a set of records, and you  can specify how many records per page you’d like to present to the user.

[Read on Medium...](https://medium.com/squeezerhq/geared-pagination-in-rails-behind-the-scenes-61d9e227540e)
## [6][Upgrading to Rails 5, need help with strong params?](https://www.reddit.com/r/rails/comments/f8fe35/upgrading_to_rails_5_need_help_with_strong_params/)
- url: https://www.reddit.com/r/rails/comments/f8fe35/upgrading_to_rails_5_need_help_with_strong_params/
---
Hello, I'm new to this, but I'm studying and trying to upgrade an old rails app (was on rails 3 so there's attr\_accessible everywhere in every single model) to rails 5. I have all of my specs passing on Rails 5.1, and I am using the "protected\_attributes\_continued" gem.

I know this gem isn't really maintained/endorsed by the Rails team, so I want to convert to "strong params". I am so confused at how to do this, I've already read a bunch of guides.

Half of the models do not even have controllers where I could put the strong params. And the majority of them are heavily nested with belongs\_to, etc.

\*\*\* I have experimented with just deleting every single attr\_accessible in my entire application, removing the protected\_attributes\_continued gem, and removing the whitelist line in my config file....

My specs PASS when I do all of this, and I've tested the application and nothing really seems to be wrong, but I didn't move ANYTHING from attr\_accessible to the controllers.

Is this okay? I'm just confused at how my application still seems to be working fine.
## [7]["VFS Connection does not exist" on Cloud9](https://www.reddit.com/r/rails/comments/f8gjeb/vfs_connection_does_not_exist_on_cloud9/)
- url: https://www.reddit.com/r/rails/comments/f8gjeb/vfs_connection_does_not_exist_on_cloud9/
---
As the title says, when i try to run the $rails server command and preview running application, it tells me "Oops, VFS Connection does not exist". Any way to fix this?
## [8][Exchanging sensitive data between two Rails apps](https://www.reddit.com/r/rails/comments/f8897p/exchanging_sensitive_data_between_two_rails_apps/)
- url: https://www.reddit.com/r/rails/comments/f8897p/exchanging_sensitive_data_between_two_rails_apps/
---
I have two Rails apps and they both use SSL/HTTPS. They have a REST api to exchange data between each other. I assume this data is encrypted because both apps are SSL/HTTPS? And if so, is this secure enough to exchange sensitive data like passwords?
## [9][Is the code for the GitHub web app on Github?](https://www.reddit.com/r/rails/comments/f88rba/is_the_code_for_the_github_web_app_on_github/)
- url: https://www.reddit.com/r/rails/comments/f88rba/is_the_code_for_the_github_web_app_on_github/
---
Hello, This might be a noob question but I"m wondering if the code for GitHub web application rails app is hosted on GitHub as an open-source project?

I'm asking this because I"m interested in how their roles and scopes are stored in the database and how it is being validated for each user. I want to draw inspiration from their implementation as I'm working on a rails app myself.
## [10][Markdown redcarpet and mentions(link)](https://www.reddit.com/r/rails/comments/f88tu4/markdown_redcarpet_and_mentionslink/)
- url: https://www.reddit.com/r/rails/comments/f88tu4/markdown_redcarpet_and_mentionslink/
---
Following this tip [https://stackoverflow.com/questions/12381230/how-to-extend-redcarpet-to-support-auto-linking-user-mentions](https://stackoverflow.com/questions/12381230/how-to-extend-redcarpet-to-support-auto-linking-user-mentions)

... I edited my MarkdownRenderer in this way:

&amp;#x200B;

    class MarkdownRenderer &lt; Redcarpet::Render::HTML
    
      def preprocess(text)
        wrap_mentions(text)
      end
    
      def wrap_mentions(text)
        text.gsub! /(^|\s)(@\w+)/ do
          "#{$1}&lt;a href='/user/#{$2}'&gt;#{$2}&lt;/a&gt;"
        end
        text
      end

and in product\_helper i have this

      def parse_markdown(text)
        markdown = Redcarpet::Markdown.new(MarkdownRenderer, hard_wrap: true, autolink: true, space_after_headers: true)
        markdown.render(text)
      end

But then if I write u/marcus, it show me &lt;a href='/user/marcus'&gt;marcus&lt;/a&gt; **as a text and now as a link.** It blocks the html. How to solve?
## [11][How to step up?](https://www.reddit.com/r/rails/comments/f80e2j/how_to_step_up/)
- url: https://www.reddit.com/r/rails/comments/f80e2j/how_to_step_up/
---
Hello,  
i have been working with ROR since the start of my career but i have no idea now how can i step up.  
I work with ROR in my daily job  60 % and 40 % is for front end stuff like vue.  
How can i step up my level?  
Thanks.
## [12][Help getting set up - can't find gem railties error](https://www.reddit.com/r/rails/comments/f82shv/help_getting_set_up_cant_find_gem_railties_error/)
- url: https://www.reddit.com/r/rails/comments/f82shv/help_getting_set_up_cant_find_gem_railties_error/
---
I've been trying to get rails set up and I keep running into problem - all of which I'm sure are just user error. I'm using Windows 10 and Ruby 2.7.

When I run rails -v, I get an error saying it can't find gem railties.

&gt;C:/Ruby27-x64/lib/ruby/2.7.0/rubygems.rb:275:in \`find\_spec\_for\_exe': can't find gem railties (&gt;= 0.a) with executable rails (Gem::GemNotFoundException)  
&gt;  
&gt;from C:/Ruby27x64/lib/ruby/2.7.0/rubygems.rb:252:in \`bin\_path  
&gt;  
&gt;from C:/RailsInstaller/Ruby2.3.3/bin/rails:22:in \`&lt;main&gt;'

Is it trying to get the wrong version of Ruby? I'm really stuck and would love some help! I ran into similar problem earlier but what I tried didn't help, I attempted to delete everything and start over, and here I am. 

I'd be happy to provide in additional files that could help!
