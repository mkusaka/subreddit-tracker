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
## [3][Ruby blockchain client](https://www.reddit.com/r/rails/comments/f2oe24/ruby_blockchain_client/)
- url: https://www.reddit.com/r/rails/comments/f2oe24/ruby_blockchain_client/
---
Hello, I'm looking for an api that would allow me to sign documents on a bitcoin or ethereum blockchain like [https://chainpoint.org/](https://chainpoint.org/) for example but would provide a ruby gem to use their services.

My application allows my customers to manage and centralize all kinds of documents, I would like to offer them the possibility to be able to sign and authenticate a set of documents whenever they want.

Thank you for your answers
## [4][First Rails Interview](https://www.reddit.com/r/rails/comments/f2q3bl/first_rails_interview/)
- url: https://www.reddit.com/r/rails/comments/f2q3bl/first_rails_interview/
---
Hey all!
New to this subreddit and just have a question about interviews. 

I have my first technical interview as a backend engineer. I’m going to have a technical interview on rails and active record. 

I’m just wondering what I should look out for and if anyone has any insights. It would be greatly appreciated. 

Thank you! :)
## [5][Approach for saving lots of data in rails database](https://www.reddit.com/r/rails/comments/f2r5l1/approach_for_saving_lots_of_data_in_rails_database/)
- url: https://www.reddit.com/r/rails/comments/f2r5l1/approach_for_saving_lots_of_data_in_rails_database/
---
Hi! Im creating this scraper and would like to store the following data. I want to achieve something like this: Every time the scraper runs, save data in a new table and display it on the product page

I was thinking to go with this approach:

Create a controller/ model; `IntervalChecks inventory_qty:string product_price:string product:references`

And then i can display the related inventory data and the prices on the product show page.

Is this a good approach (with speed in mind?)

Or would you advice to store it as a key value pair in Json as hash (in that case I would need to read more about how that stuff works :)
## [6][How can I convince my company to send me to RailsConf?](https://www.reddit.com/r/rails/comments/f2c87e/how_can_i_convince_my_company_to_send_me_to/)
- url: https://www.reddit.com/r/rails/comments/f2c87e/how_can_i_convince_my_company_to_send_me_to/
---
Hi all

I've never been to any big conference really; mainly meetups. I wanted to go to RubyConf last year, but was saving up PTO hours to go on another vacation.

I work for a Mortgage banking company, so tech isn't exactly the main focus here but it is a very important part of the company; we're actually sponsoring an upcoming Linux expo.

There's a budget in my department to potentially send me and maybe a few other teammates to a conference, and I really want to go to RailsConf. My team (and several others in the company) use Rails; although my team uses it rather unconventionally (we use [Netzke](https://netzke.org/) and ExtJS, which means we don't have a bunch of controllers everywhere).

I've watched several talks from RailsConf and different conferences and I love learning new information; it makes me feel valuable and gives me ideas to push the company/my team forward. If I went to RailsConf, I could pick those speakers' brains and thank them personally after every talk.

But alas, my manager and his manager are not huge fans of conferences; their perspective is that you can watch all the talks online anyway, and that conferences are mostly for networking and job hunting.

&amp;#x200B;

Since I've never been to a conference before and don't know all the little details, I wanted to ask who here has gone to conferences (particularly RailsConf) and has used the information to really bring meaningful changes to their organization?

Would appreciate any insight. Thank y'all!
## [7][Rails and Sharepoint API Calendar](https://www.reddit.com/r/rails/comments/f2cq5o/rails_and_sharepoint_api_calendar/)
- url: https://www.reddit.com/r/rails/comments/f2cq5o/rails_and_sharepoint_api_calendar/
---
Google calendar has failed me for the last time, and I need to transition my apps public facing calendar to Sharepoint. My umbrella company does not allow public facing Sharepoint calendars so I have to pull the data via the API and post it into a calendar face that I program. I am looking for the best way to approach this for both a basic calendar layout to pull the data into and the best way to pull the data via the API. I have not every worked with API's and I am not sure where to start. Thanks for any and all help!!
## [8][Migrating from Rails 5 to Rails 6](https://www.reddit.com/r/rails/comments/f26c6v/migrating_from_rails_5_to_rails_6/)
- url: https://www.reddit.com/r/rails/comments/f26c6v/migrating_from_rails_5_to_rails_6/
---
I have a long running Rails 5 application with "a lot" of javascript files. It currently has webpack as well but not the default configuration used by Rails 6. Anyone who could recommend an article / tutorial / best way to migrate from sprockets asset pipeline Rails 5 to webpack based Rails 6?
## [9][Managing dependencies in a Gem that will be part of a Rails application](https://www.reddit.com/r/rails/comments/f28qg5/managing_dependencies_in_a_gem_that_will_be_part/)
- url: https://www.reddit.com/r/rails/comments/f28qg5/managing_dependencies_in_a_gem_that_will_be_part/
---
Hi all,

I am breaking out some functionality from a Rails app into separate gems so that they can be used in other Rails projects. We expect that these gems will always be used in a Rails environment.  In fact, some of the code in these gems depends on various Rails components.  For example, we use `#camelize` and other methods from ActiveSupport.  There are utility methods that apply a certain sort order to collections and they call `#reorder` which assumes ActiveRecord.

So I'm wondering how I should manage the dependencies these gems have on those Rails modules.  I don't want to include all of Rails in this gem.  So I assume that there's no "magic" here and the best practice is to simply determine which Rails gems I need, reference those gems directly in my Gemfile/gemspec and require the specific modules I need in my gems' source.  Correct?
## [10][Completely new to rails](https://www.reddit.com/r/rails/comments/f2bdgp/completely_new_to_rails/)
- url: https://www.reddit.com/r/rails/comments/f2bdgp/completely_new_to_rails/
---
Hello all,

As the title says, I'm completely new to rails and I'm currently trying to figure out what I need to do/download for my computer to be able to read/run files. I currently have a file hosted online I'm trying to read and edit for work. I've tried google and haven't had much luck yet. Any sort of help would be appreciated!
## [11][Problems with rails server](https://www.reddit.com/r/rails/comments/f2a1j5/problems_with_rails_server/)
- url: https://www.reddit.com/r/rails/comments/f2a1j5/problems_with_rails_server/
---
Hey. I'm a rails newbie so i'll try to explain the best i can.                                         When i start up a rails server in cloud9, everything goes well, but when i try to preview the running application, it tells me "Blocked host: (host) To allow requests to (host), add the following to your environment configuration: config.hosts &lt;&lt; (host)". I tried to do what it told me but it didn't work. Anyone know how to fix this?
## [12][Manage Redis by Ruby On Rails](https://www.reddit.com/r/rails/comments/f1x280/manage_redis_by_ruby_on_rails/)
- url: https://www.reddit.com/r/rails/comments/f1x280/manage_redis_by_ruby_on_rails/
---
I come to present a tool that I have developed with 2 others developers [https://github.com/OpenGems/redis\_web\_manager](https://github.com/OpenGems/redis_web_manager)

This tool allows you to manage Redis, thanks to a web interface you  can easily manage your Redis instance (see your keys, the memory used,  clients connected, etc …).

Enjoy !
