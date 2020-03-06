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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/fduax9/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/fduax9/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Is there a more efficient way to add hundreds of things to the database?](https://www.reddit.com/r/rails/comments/fe82w2/is_there_a_more_efficient_way_to_add_hundreds_of/)
- url: https://www.reddit.com/r/rails/comments/fe82w2/is_there_a_more_efficient_way_to_add_hundreds_of/
---
Let's say I have a JSON object containing an array of 300+ objects. I want to create an instance of a given model for each of these then save that instance to the database.

Now, I know I can do this in a loop, but that seems pretty inefficient for several hundred objects. **Is there a more efficient way to do this, and if so, what should I be looking** (googling) **for?**

```
# This seems like a very inefficient way to do things with 300-1000
# things in params[:something] (and no I can't reduce or predict that
# number of things).
json = JSON.parse(params[:something])
json['foo'].each do |foo|
  bar = Foo.new(thing1: foo['thing1'], thing2: foo['thing2'])
  bar.save
end
```
## [4][Deploying Hundreds of Applications to AWS](https://www.reddit.com/r/rails/comments/fe448t/deploying_hundreds_of_applications_to_aws/)
- url: https://www.reddit.com/r/rails/comments/fe448t/deploying_hundreds_of_applications_to_aws/
---
Hey gang, I'm having a bit of trouble researching anything truly applicable to my specific case.  For context, my company has \~150 different applications (different code, different purpose, no reliance on each other) each deployed to its own set of EC2 servers based on the needs of the application.  To do this, our deployment stack uses Capistrano 2 and an internal version of Rubber.  This has worked for years but management is pushing modernization and I want to make sure that it's done with the best available resources that will avoid as many blockers down the road.

Everything I find is mainly designed under the context that all containers are generally related and grouped as such.  When that's not the case, there's only a small number.  


Still, all research points to Docker. Creating an image that we could use as a base for all applications then each application would be created as its own container.  That seems like just as much management of resources at the end of the day but with slightly simpler deployment.  


To help with said management, I've seen suggestions of setting up Kubernetes, turning each application into its own cluster and using Rancher (or alternatives).  While this sounds good in theory, Kubernetes isn't exactly designed for this purpose.  It would work but I'm not sure it's the best solution.

&amp;#x200B;

So I'm hoping someone out there may have insight or advice.  Anything at all is greatly appreciated.
## [5][Mapbox App linked to SQLite3 database?](https://www.reddit.com/r/rails/comments/fe49n6/mapbox_app_linked_to_sqlite3_database/)
- url: https://www.reddit.com/r/rails/comments/fe49n6/mapbox_app_linked_to_sqlite3_database/
---
Hey all, first post here. Amateur Ruby on Rails user working on my college senior project. We are working on building a site that resembles the Toronto Poetry Map: [https://torontopoetry.ca/](https://torontopoetry.ca/) \- except for the monuments in D.C. instead. We have the project on Rails and have a Mapbox app to display the locations of monuments and such and an SQLite3 database for the backend. Only issue is we're unsure of how (or if) we are able to have the map pull information from the database instead of from a json file. Google searches didn't help me much on this issue -- if anyone could provide some advice we would greatly appreciate it. Thanks!
## [6][Can't install Rails on Windows 10](https://www.reddit.com/r/rails/comments/fe2aoj/cant_install_rails_on_windows_10/)
- url: https://www.reddit.com/r/rails/comments/fe2aoj/cant_install_rails_on_windows_10/
---
I've been running into the "sqlite error code"  and not being able to render the server. I've gone through about 20 different solutions trying to update that, using Ubunto on Cloud9 (through Amazon Web Services) and still no luck. 

&amp;#x200B;

Any Windows users have any luck installing and successfully deploying Rails?

&amp;#x200B;

Also my Ruby version was 2.6.5 and my Rails version was 5.0 (?) I just uninstalled it and im trying to reinstall it again. (don't know if these versions are too advanced and causing errors?)

&amp;#x200B;

Thanks!
## [7][Help to find the right callback method to be used](https://www.reddit.com/r/rails/comments/feb4og/help_to_find_the_right_callback_method_to_be_used/)
- url: https://www.reddit.com/r/rails/comments/feb4og/help_to_find_the_right_callback_method_to_be_used/
---
I have a `user` model that has a callback method `do_something` in the `before_update`

that has an association to roles model

    has_many :roles, autosave: true, after_add: :added, after_remove: :removed

and in the do\_something callback I am trying to compute some values based on the `user.roles` association.

Currently, during an update on the user model, I first remove all the roles that the user wants to be deleted like so `user.roles.destroy(each_individual_role)` and in the `do_something` callback when I try `user.roles` I get only the valid users i.e. all the deleted roles don't show up. This deletion of the roles happen before the `user.update_attributes` so if `update_attributes` fails the associated roles would have been deleted and it is an inconsistent state.

To avoid the above problem and to do the roles creation/deletion in one go I decided to delete the roles via the nested attributes like so

    [{id: 54, _destroy: true}, {id: nil, name: "New Role to be added"}]

Now, I have two problems

1. in the do\_something callback when I try and access `user.roles` I still get the role corresponding with `"_destroy": true` since I believe the association hasn't been broken yet. and this results in incorrect computation in the `do_something` callback
2. the `before_update` is being called twice. I believe one after saving each entry to the role table.

To resolve this I converted do\_something at after\_save and then do a `self.update_column` in `do_something`. This works fine as at `after_save` the roles attribute has been saved. But I do not like this solution of updating after saving that too by passing validations and I require the callbacks it uses as well

Is there a better way to resolve this?

Edit: ideally I would like to have the updates roles attribute from the roles_attribute on the before_update callback and the privileges I apply will have to be used for other callbacks as well.
## [8][Best practices for blocking bot signups?](https://www.reddit.com/r/rails/comments/fdwbuc/best_practices_for_blocking_bot_signups/)
- url: https://www.reddit.com/r/rails/comments/fdwbuc/best_practices_for_blocking_bot_signups/
---
My side project has a single-step signup: you just put in your email and hit subscribe. (Behind the scenes, all email is delivered via SendGrid and I automatically suppress any emails that are dropped/undeliverable. I also stop the automated weekly emails if the user has not opened or clicked any emails within the past 31 days. According to SendGrid, I've got a 99% sender reputation score...)

Currently, I see a couple hundred valid signups per day (thanks, SEO!) and I can confirm they're "real" because they clearly engage with the emails, update their account settings and do other things that normal users do on my app.

Lately, however, I've noticed a couple of weird signups creeping into the mix each day... they signup and then immediately request a password reset email. Additionally, their emails get deferred because of "receiving too much email currently..." and there is zero additional login attempts or any other "real" behavior. A quick look at their IP addresses seems to indicate that they're coming from Eastern Europe and Asia. My site is mostly relevant to US-based users but I don't want to entirely block outsiders because I also see a lot of traffic from military bases around the world. (I should also mention that I'm not seeing very many repeat IP addresses... they're definitely rotating them.)

Before I start messing with recaptchas (which would add friction to my "real" users), I'm looking for other solutions...

I'm guessing I should be using something like rack-attack... but I'm not quite sure what the best practices might be to systematically block offending IPs when they're rotating, etc. Help? :)
## [9][Creating an OpenID Connect Server](https://www.reddit.com/r/rails/comments/fdzs0y/creating_an_openid_connect_server/)
- url: https://www.reddit.com/r/rails/comments/fdzs0y/creating_an_openid_connect_server/
---
I have an existing Rails server that I'd like to function as an OpenID Connect Provider. I've seen that there are at least 2 gems that might help with this: [https://github.com/doorkeeper-gem/doorkeeper-openid\_connect](https://github.com/doorkeeper-gem/doorkeeper-openid_connect) and [https://github.com/nov/openid\_connect](https://github.com/nov/openid_connect). I'm wondering if any of you had had any experience with these gems for this purpose and could tell me about your experiences?

I'm running Rails 6/Ruby 2.6.5

Thanks!
## [10][Redirect all get requests except for these routes...?](https://www.reddit.com/r/rails/comments/fe2wk9/redirect_all_get_requests_except_for_these_routes/)
- url: https://www.reddit.com/r/rails/comments/fe2wk9/redirect_all_get_requests_except_for_these_routes/
---
Title says it all. Right now Devise will catch the get request upon initial visit to check for user session. 

&amp;#x200B;

Once in my app, I'd like to be able to signout, but the below block redirects the request to my react application instead of hitting devise/sessions/destroy. I would like to make an exception for this route in the below block.

&amp;#x200B;

 `get '*page', to: 'static#index', constraints: -&gt; (req) do`  
`!req.xhr? &amp;&amp; req.format.html?`  
 `end`

&amp;#x200B;

Here's my routes.rb. Thanks for the help!

&amp;#x200B;

`# API`  
  `namespace :v1, defaults: { format: 'json' } do`  
`# get 'users', 'users#index'`  
  `end`  
  `devise_for :users, :path =&gt; '', :path_names =&gt; {`  
`:sign_in =&gt; 'sign-in',`  
`:sign_out =&gt; 'sign-out',`  
`:sign_up =&gt; 'sign-up'`  
  `},`  
 `controllers: {`  
`sessions: 'sessions',`  
`registrations: 'registrations'`  
  `}`  
 `# Forward all requests to StaticController#index but requests`  
 `# must be non-ajax (!req.xhr?) and HTML Mime type (req.format.html?)`  
 `# This does not include the root ("/") path.`  
 `get '*page', to: 'static#index', constraints: -&gt; (req) do`  
`!req.xhr? &amp;&amp; req.format.html?`  
 `end`  
 `# Forward root to StaticController#index`  
  `root "static#index"`
## [11][Layer Breeding Equipment Market 2020 Top Key Vendors are | Big Dutchman, Texha, Guangdong Guangxing, Big Herdsman Machinery, Guangzhou Huanan Poultry Equipment, Chore-Time Brock, Qindao Tianrui Poultry Equipment](https://www.reddit.com/r/rails/comments/febc2v/layer_breeding_equipment_market_2020_top_key/)
- url: https://www.reddit.com/r/rails/comments/febc2v/layer_breeding_equipment_market_2020_top_key/
---
**Global  Layer Breeding Equipment Market 2020 Research Report**

This report explains the key market drivers, trends, restraints and opportunities to give a precise data which is required and expected. It also analyzes how such aspects affect the market existence globally helping make a wider and better choice of market establishment. The  Layer Breeding Equipment markets growth and developments are studied and a detailed overview is been given.

This Report covers the companies data, including: shipment, price, revenue, gross profit, interview record, business distribution etc., these data help the consumer know about the competitors better. This report also covers all the regions and countries globally, which shows a regional development status, including market size, volume and value, as well as price data.

**Thoroughly Studied Key Players ||   Big Dutchman, Texha, Guangdong Guangxing, Big Herdsman Machinery, Guangzhou Huanan Poultry Equipment, Chore-Time Brock, Qindao Tianrui Poultry Equipment, Shanghai Extra Machinery, Facco, Langfang Yanbei Animal Husbandry Machinery Group, Henan Jinfeng Poultry Equipment, GARTECH Equipment, HYTEM, and Fienhage Poultry-Solutions** 

The increasing demand in the well-established and emerging regions as well as latest technological advent  Layer Breeding Equipment, and the growing insistence of the end-use industries are all together driving the growth of the  Layer Breeding Equipment Industry.

This report however describes a brief summary of market and explains the major terminologies of the  Layer Breeding Equipment Industry. However an accurate analysis of the market trends, drivers, challenges and opportunities has derived the most reasonable outlook of the  Layer Breeding Equipment market.

**Get sample copy of  Layer Breeding Equipment :** [**https://www.reportsandmarkets.com/sample-request/global-layer-breeding-equipment-market-2020-by-manufacturers-regions-type-and-application-forecast-to-2025?utm\_source=reddit&amp;utm\_medium=47**](https://www.reportsandmarkets.com/sample-request/global-layer-breeding-equipment-market-2020-by-manufacturers-regions-type-and-application-forecast-to-2025?utm_source=reddit&amp;utm_medium=47) 

**The study objectives are:**

To analyze and research the global  Layer Breeding Equipment status and future forecast involving, production, revenue, consumption, historical and forecast.

To present the key  Layer Breeding Equipment  manufacturers, production, revenue, market share, and recent development.

To split the breakdown data by regions, type, manufacturers and applications.

To analyze the global and key regions market potential and advantage, opportunity and challenge, restraints and risks.

To identify significant trends, drivers, influence factors in global and regions.

To analyze competitive developments such as expansions, agreements, new product launches, and acquisitions in the market.

And more.

**In this study, the years considered to estimate the market size of  Layer Breeding Equipment**

History Year: 2015–2019

Base Year: 2019

Estimated Year: 2020

Forecast Year: 2020–2027

**The report offers in-depth assessment of the growth and other aspects of the  Layer Breeding Equipment market in important countries (regions), including:**

North America

Europe

Asia Pacific

Middle East &amp; Africa

Latin America

America Country (United States, Canada)

South America

Asia Country (China, Japan, India, Korea)

Europe Country (Germany, UK, France, Italy)

Other Country (Middle East, Africa, GCC)

**Get Complete Report  Layer Breeding Equipment @** [**https://www.reportsandmarkets.com/reports/global-layer-breeding-equipment-market-2020-by-manufacturers-regions-type-and-application-forecast-to-2025?utm\_source=reddit&amp;utm\_medium=47**](https://www.reportsandmarkets.com/reports/global-layer-breeding-equipment-market-2020-by-manufacturers-regions-type-and-application-forecast-to-2025?utm_source=reddit&amp;utm_medium=47) 

1 Report Overview

2 Global Growth Trends

3 Market Share by Key Players

4 Breakdown Data by Type and Application

5 United States

6 Europe

7 China

8 Japan

9 Southeast Asia

10 India

11……

12…….

13……..

14 Central &amp; South America

15 International Players Profiles

16 Market Forecast 2020-2027

17 Analyst's Viewpoints/Conclusions

18 Appendix

**About Us:**

Market research is the new buzzword in the market, which helps in understanding the market potential of any product in the market. Reports And Markets is not just another company in this domain but is a part of a veteran group called Algoro Research Consultants Pvt. Ltd. It offers premium progressive statistical surveying, market research reports, analysis &amp; forecast data for a wide range of sectors both for the government and private agencies all across the world.

**Contact Us:**

Sanjay Jain

Manager – Partner Relations &amp; International Marketing

https://www.reportsandmarkets.com/

Ph: +1-352-353-0818 (US)
## [12][The tricky situation with rails before_validation callbacks](https://www.reddit.com/r/rails/comments/fdwrl2/the_tricky_situation_with_rails_before_validation/)
- url: https://www.reddit.com/r/rails/comments/fdwrl2/the_tricky_situation_with_rails_before_validation/
---
I have a `user` model that can have many `emails` model. I have a before validation callback method that sets the primary email from the list of `emails` for that user. 

So when I create an email in the `user` model the `user_email` object with primary email will automatically be updated as part of the `before_validation` callback. 

say a `user` A has the following entries in the `emails` table

    user_id email primary 
    A.      EM    true
    C       MF   false

If I wanted to update the email for user A to "MF", I'll set the `email_attributes` to `{ id: 45, "_destroy" : 1 }` and call `item.update_attributes({email : "MF"})`

Now user A  will have email "MF" first in the user table. Then via callbacks, the email object will have its email updated to "MF"  for the email with primary = true email entry which will fail as the before validation is called first and the email\_attributes which destroy the entry "MF" which causes duplicate hasn't been destroyed yet. 

  
To overcome this I'll have to explicitly call Email.find(45).destroy and then there will be no duplicates. 

I'm not sure how else to solve this. Any help would be great.
