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
## [3][Is there a complete overview about rails conventions?](https://www.reddit.com/r/rails/comments/f3iqbr/is_there_a_complete_overview_about_rails/)
- url: https://www.reddit.com/r/rails/comments/f3iqbr/is_there_a_complete_overview_about_rails/
---
Hey there,

I'm looking for a complete overview about all rails conventions. There is one on github: [https://gist.github.com/iangreenleaf/b206d09c587e8fc6399e](https://gist.github.com/iangreenleaf/b206d09c587e8fc6399e)

This overview has many points but not all. Recently I found a rails tutorial that mentioned that when we use render \`@users\` rails look for the partial \_user.erb that is rendered for every user in that collection. Pretty useful to know. And I bet there are many more conventions like this beside the overview above mentioned.

What else useful rails conventions do you know?
## [4][Scaling DB schema changes: How do you do these things?](https://www.reddit.com/r/rails/comments/f3esci/scaling_db_schema_changes_how_do_you_do_these/)
- url: https://www.reddit.com/r/rails/comments/f3esci/scaling_db_schema_changes_how_do_you_do_these/
---

I've built a number of applications. The largest perhaps 5-10 write/update heavy tables reaching over 50 million records, and the application server pool getting 1000 requests per second at peak load. The scales and teams I work with can stomach a few seconds, maybe a minute or two, of downtime during a deployment.

But I am wondering how teams handle bigger updates where downtime might nto be acceptable, or at scales where changes take significant amounts of time. Or in case where another team is dependent on your system, but they cannot change their system at the same moment you release your changes.


The simplest toy example I can imagine would be altering the decimal precision on a a large "orders" table where any such alteration to the table takes several minutes. 

Another scenario is where the business changes "first and last name" columns to a single "full name", where your system can be changed quickly, but other team's components cannot be modified alongside your changes.

A more complex might be changing from an Entity-Attribute-Value system (i.e., property bags) to a concrete set of tables. So you introduce significant schema changes on a large database, and also a signifiant amount of code changes. 

This is a very loose and hand-waving kind of question, but it's hard to find information on this kind of topic as I don't quite know how to phrase questions. Hope I'm being clear!

* edit 1: in the case of slowly changing schemas, do you just try and backwards-port the data where possible, throw in switch statements (or differing classes) to handle various versions, etc?
## [5][Just curious. How many req/s does you app handle?](https://www.reddit.com/r/rails/comments/f3i1pt/just_curious_how_many_reqs_does_you_app_handle/)
- url: https://www.reddit.com/r/rails/comments/f3i1pt/just_curious_how_many_reqs_does_you_app_handle/
---
Aldo how many instances of the app, how many servers etc? Like I said I'm just curious about the scale at which your apps operate. Thanks!
## [6][Bad conventions at my new job, what to do?](https://www.reddit.com/r/rails/comments/f31lal/bad_conventions_at_my_new_job_what_to_do/)
- url: https://www.reddit.com/r/rails/comments/f31lal/bad_conventions_at_my_new_job_what_to_do/
---
Hi, I'm from Peru and I wanted to ask something, at my new job there are a lot of bad conventions, just to name a few:

* Models, controllers and code in general are not in English.
* Different `database.yml` for development and production, also they don't use environment variables or credentials in any part of the project, just plain text everything.
* Controllers with more than 1000 lines (WET as fuck, also not a single concern in this project)
* No tests.
* Garbage files like `.byebug_history`, log files and temp files are not considered at `.gitignore`.
* I think `rubocop` killed himself.

This is not a technical question, but a more human one, what should I do? Say fuck off and don't trouble myself trying to bring this project to a decent level and only do the requirements they want me to do? I feel I can't tell them to do the things right since I'm the new guy or I can, but just a few, not all of them or they will think badly of me.
## [7][Deploy staging and production environments to single server using Capistrano — Rails](https://www.reddit.com/r/rails/comments/f3791t/deploy_staging_and_production_environments_to/)
- url: https://www.reddit.com/r/rails/comments/f3791t/deploy_staging_and_production_environments_to/
---
The author shows you several different configuration options to deploy different environments a single server. 

Link: https://medium.com/lalusaud_84531/deploy-staging-and-production-applications-to-single-server-using-capistrano-rails-1d5ab558d44f?source=friends_link&amp;sk=a686b1a3d16b7bb644be697136e9c5f3
## [8][Handling abandoned records](https://www.reddit.com/r/rails/comments/f2x1o5/handling_abandoned_records/)
- url: https://www.reddit.com/r/rails/comments/f2x1o5/handling_abandoned_records/
---
I'll describe a similar model to mine. I'm going to use an oversimplified music library.

Models:

`Album`

`has_many songs`

&amp;#x200B;

`Track`

`belongs_to album`

&amp;#x200B;

`AlbumTrack`

`references album`

`references song`

In this example, a user would enter an album, then while viewing the album have the option to add tracks (via the AlbumTracks table). This would allow an indefinite number of tracks to be added to each album.

It's possible a user would create the album, then never add any tracks. It's also possible the user would return at a later time and create a new album intended to contain the data for the previously abandoned album. In any case, there is now 1 or more albums without tracks.

Is there a conventional Rails way to approach preventing an ever growing list of unpopulated albums?

One approach I'm considering is running a check when a user attempts to add a new Album that first checks if there is an abandoned record, and if so, presents that record instead of a new record. (One possible problem is that if 2 users are using the app, they might both attempt to create a new record at the same time, and both be presented with the abandoned record.)

Edit: I'm also currently looking at the Orphanage gem ([https://rubygems.org/gems/orphanage/versions/0.3.0](https://rubygems.org/gems/orphanage/versions/0.3.0)) which might work for me. In this case, I imagine I could show forms for both Album and AlbumTrack on the same view, and use Orphanage to create each AlbumTrack and reassign them to Album once the Album is saved. Not sure how this would handle a scenario where the abandonment occurs right before the save (but the individual tracks are already created).
## [9][Can I savely use ActiveRecord transactions or do they have any negative impact?](https://www.reddit.com/r/rails/comments/f2tu58/can_i_savely_use_activerecord_transactions_or_do/)
- url: https://www.reddit.com/r/rails/comments/f2tu58/can_i_savely_use_activerecord_transactions_or_do/
---
I'm currently developing a feature where I need to create a bunch of records (between 5-10) all at once. The important part here is, that either all are created or none at all. Now this to me sounds like a perfect use-case for an `ActiveRecord::Transaction` but I don't have a lot of experience with using them and from all I've read it sounds like they are "blocking" the database until the transaction is done. To me this sounds pretty bad, because if I need to create for example 10 records and maybe even do a little bit of business logic all in one transaction then this could take quite some time in which the entire database would be locked. However, I also don't see a different way of achieving my desired result.

So before I go and use a transaction for this task I wanted to ask the following: Is my current picture of how a transaction works and the downsides that it has correct? Should I try and avoid transactions when possible or can I savely use them whenever I need to make sure that a task happens fully or not at all? And if I'm correct in asuming that they "lock" the database, then how do transaction scale? Because I could imagine that when there are multiple users/second which are all locking the database that this wouldn't be ideal.

Edit: I forgot to mention that the feature I'm currently developing will not be used a lot. Every user will probably only use it once per session.
## [10][Approach for saving lots of data in rails database](https://www.reddit.com/r/rails/comments/f2r5l1/approach_for_saving_lots_of_data_in_rails_database/)
- url: https://www.reddit.com/r/rails/comments/f2r5l1/approach_for_saving_lots_of_data_in_rails_database/
---
Hi! Im creating this scraper and would like to store the following data. I want to achieve something like this: Every time the scraper runs, save data in a new table and display it on the product page

I was thinking to go with this approach:

Create a controller/ model; `IntervalChecks inventory_qty:string product_price:string product:references`

And then i can display the related inventory data and the prices on the product show page.

Is this a good approach (with speed in mind?)

Or would you advice to store it as a key value pair in Json as hash (in that case I would need to read more about how that stuff works :)
## [11][First Rails Interview](https://www.reddit.com/r/rails/comments/f2q3bl/first_rails_interview/)
- url: https://www.reddit.com/r/rails/comments/f2q3bl/first_rails_interview/
---
Hey all!
New to this subreddit and just have a question about interviews. 

I have my first technical interview as a backend engineer. I’m going to have a technical interview on rails and active record. 

I’m just wondering what I should look out for and if anyone has any insights. It would be greatly appreciated. 

Thank you! :)
## [12][Ruby blockchain client](https://www.reddit.com/r/rails/comments/f2oe24/ruby_blockchain_client/)
- url: https://www.reddit.com/r/rails/comments/f2oe24/ruby_blockchain_client/
---
Hello, I'm looking for an api that would allow me to sign documents on a bitcoin or ethereum blockchain like [https://chainpoint.org/](https://chainpoint.org/) for example but would provide a ruby gem to use their services.

My application allows my customers to manage and centralize all kinds of documents, I would like to offer them the possibility to be able to sign and authenticate a set of documents whenever they want.

Thank you for your answers
