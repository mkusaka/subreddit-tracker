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
## [3][How do you include the lib folder in your spec ?](https://www.reddit.com/r/rails/comments/f44p47/how_do_you_include_the_lib_folder_in_your_spec/)
- url: https://www.reddit.com/r/rails/comments/f44p47/how_do_you_include_the_lib_folder_in_your_spec/
---
I have some custom PORO in the lib/ folder. I am using Rails 5 and Rspec. I would like to include files from lib in the spec. I found answers [1](https://stackoverflow.com/questions/4073856/rails-3-autoload), [2](https://stackoverflow.com/questions/16954989/how-to-include-lib-directory-in-rspec-tests) which is too old. So I am seeking some advices majority of rails community follows.
## [4][First project finished, but badly in need of criticism!!](https://www.reddit.com/r/rails/comments/f3z6zt/first_project_finished_but_badly_in_need_of/)
- url: https://www.reddit.com/r/rails/comments/f3z6zt/first_project_finished_but_badly_in_need_of/
---
##Hey there you wonderfully helpful subreddit, you!

Thanks to so much help from you guys over the last month+ I've finished and deployed my first webapp. [You can check it out here](https://www.chucksef.com) but it really won't be terribly interesting since none of you are admins (and hopefully there's no way for you to sneak your way in!)

[Here's a link to the repo on github](https://github.com/Chucksef/photo_app) if you're interested in looking at the code.

If you didn't read the title, I **very badly want criticism**! I'm very new to rails (just started learning Jan 3rd) so I'm sure there are an overwhelming number of things that I've done 100% wrong, and I'd like to know about them!

#Website concept: 
Squarespace-y platform I built myself to allow me to dynamically build my own portfolio without having to go back into my development environment and re-code pages as I finish more projects to show off!

Right now, however, I've filled it with some generic photo-studio stuff.

However, here are a few categories where I'm NOT looking for criticism:

* Webpage is very generic looking bootstrap bullshit - I know. That's not the point of this project!
* I didn't use ANY testing tools while building this. Oh man do I really regret glossing over this stuff. Now that I'm finished every time I make a change I cringe at not knowing if something else broke! I WILL be learning this in the near future.
* Website copy is all bad - I'm not trying to become a copywriter.

Anything else is fair game! So come on and help me become better by destroying my month+ of studying and hard work!
## [5][Highcharts in Ruby syntax for Rails app / Also, trying to get two series of data to display on a single highchart and it's not taking](https://www.reddit.com/r/rails/comments/f3ybw0/highcharts_in_ruby_syntax_for_rails_app_also/)
- url: https://www.reddit.com/r/rails/comments/f3ybw0/highcharts_in_ruby_syntax_for_rails_app_also/
---
Does anyone know of any good docs for highcharts in rails?  Seems sparse if much at all.  I don't have a JS background yet to translate much from JS to Ruby, if that is useful (I notice lots of demos in JS fiddle).

Additionally, Im trying to get two series of data to load on a highchart in rails.  One column type and one scatter type, each with nested arrays like `[[x1, y1], [x2, y2], [x3, y3],  ...]` and so on

I'm making accessibly charts, and I've tried `chart.series( { &lt; chart 1&gt; }, { &lt;chart 2&gt; })`, I've tried making separate `chart.series`, I've tried making two separate Accessible charts.  One chart is working but can't get to seem to get the other one going, trying to add the second plot in various was as mentioned above.  Page won't load properly when I try to get a second one in
## [6][How nested_attributes works?](https://www.reddit.com/r/rails/comments/f3ui2c/how_nested_attributes_works/)
- url: https://www.reddit.com/r/rails/comments/f3ui2c/how_nested_attributes_works/
---
Hi guys, I was wondering how Rails create &amp; update nested attributes.

I had a problem where I was creating 2 associated models at the same time. For that, I was using the rails callbacks :after\_create/update, and while I was doing that, I knowledge that Rails pushes the code to the Stack (execute asynchronous), hence, inside the method called by the callback, we only have access to database's data. Also, apparently, we can't pass arguments (I tried and didn't work) to those triggers.

For this problem, the only way I was able to pass data from one model to another was creating database attributes to the first model and passing them to the second model. I did try to pass attr\_acessor first, but that didn't work for the reason I described above.

Besides using nested\_attributes to this problem, there was something that I'm missing? I guess that Rails has more privileges to pass arguments to these triggers or maybe the arguments should be feasible of serialization - similar to how works Sidekiq's Workers.

I hope this example can clarify my explanation

    # inside First::Model.rb
    
    attr_accessor :second_model_attr_4
    
    after_create :create_solicitation
    
    def create_solicitation
        puts second_model_attr_4 # returns nothing
     
        Second::Model.create!(
          first_model_id: self.id,
          second_model_attr_1: self.second_model_attr_1 
          second_model_attr_2: self.second_model_attr_2
          second_model_attr_3: Second::Model::Status.find(
            self.second_model_attr_3
         ),
          # this attr_accessor won't be seem by the function  
          second_model_attr_4: self.second_model_attr_4 
        )
    end
    
    # P.s second_model_attr_1 to 3 were added as database attributes to the First::Model (did I migration and add these methods to First::Model)

P.s It's callbacks, not triggers - sorry.
## [7][bundle exec foreman start fails](https://www.reddit.com/r/rails/comments/f436wc/bundle_exec_foreman_start_fails/)
- url: https://www.reddit.com/r/rails/comments/f436wc/bundle_exec_foreman_start_fails/
---
&gt; \`parse\_numeric': Expected numeric value for '--port'; got "${PORT-3000}" (Thor::MalformattedArgumentError)  


All I am doing is creating the environment for an open source project.  
 [https://github.com/huginn/huginn](https://github.com/huginn/huginn)   


Gemfile:  
 [https://github.com/huginn/huginn/blob/master/Gemfile](https://github.com/huginn/huginn/blob/master/Gemfile)   


Gemfile.lock  
 [https://github.com/huginn/huginn/blob/master/Gemfile.lock](https://github.com/huginn/huginn/blob/master/Gemfile.lock)   


I don't really know why I am getting this.  


From online sources they told me to try downgrading thor from 0.19.4 to 0.19.1, didn't work, another told me to try doing 0.20.0, but didn't work.  


I'm running this on Windows 10.
## [8][Should ActiveRecord be used for already existing files?](https://www.reddit.com/r/rails/comments/f3yvdp/should_activerecord_be_used_for_already_existing/)
- url: https://www.reddit.com/r/rails/comments/f3yvdp/should_activerecord_be_used_for_already_existing/
---
Hey,

I was reading a lot about ActiveStorage in the past days and I'm still not sure if it's the right solution for my problem or if it's more aimed at simple user file uploads.

I have an Image model which has an "external_url" field which contains a random direct URL to an image scraped from the internet. It's linked to another model called "Show". I don't want to store that image blob as soon as I insert a new Image as it would be too many (in my case).

What I want is that as soon as someone requests that image by visiting a show page like: http://localhost:3000/show/example-show it'll download it from the "external_url", then upload it to S3 and create a variant of it. Just like ActiveStorage usually does, just with the difference that it has to download it first and then do it's usual thing.

Is that an unusual approach and should I just use an external image server like https://github.com/willnorris/imageproxy instead?

Any pointers into the right direction or usual approaches are appreciated!

Thanks!
## [9][Is there a complete overview about rails conventions?](https://www.reddit.com/r/rails/comments/f3iqbr/is_there_a_complete_overview_about_rails/)
- url: https://www.reddit.com/r/rails/comments/f3iqbr/is_there_a_complete_overview_about_rails/
---
Hey there,

I'm looking for a complete overview about all rails conventions. There is one on github: [https://gist.github.com/iangreenleaf/b206d09c587e8fc6399e](https://gist.github.com/iangreenleaf/b206d09c587e8fc6399e)

This overview has many points but not all. Recently I found a rails tutorial that mentioned that when we use render \`@users\` rails look for the partial \_user.erb that is rendered for every user in that collection. Pretty useful to know. And I bet there are many more conventions like this beside the overview above mentioned.

What else useful rails conventions do you know?
## [10][Just curious. How many req/s does you app handle?](https://www.reddit.com/r/rails/comments/f3i1pt/just_curious_how_many_reqs_does_you_app_handle/)
- url: https://www.reddit.com/r/rails/comments/f3i1pt/just_curious_how_many_reqs_does_you_app_handle/
---
Aldo how many instances of the app, how many servers etc? Like I said I'm just curious about the scale at which your apps operate. Thanks!
## [11][Scaling DB schema changes: How do you do these things?](https://www.reddit.com/r/rails/comments/f3esci/scaling_db_schema_changes_how_do_you_do_these/)
- url: https://www.reddit.com/r/rails/comments/f3esci/scaling_db_schema_changes_how_do_you_do_these/
---

I've built a number of applications. The largest perhaps 5-10 write/update heavy tables reaching over 50 million records, and the application server pool getting 1000 requests per second at peak load. The scales and teams I work with can stomach a few seconds, maybe a minute or two, of downtime during a deployment.

But I am wondering how teams handle bigger updates where downtime might nto be acceptable, or at scales where changes take significant amounts of time. Or in case where another team is dependent on your system, but they cannot change their system at the same moment you release your changes.


The simplest toy example I can imagine would be altering the decimal precision on a a large "orders" table where any such alteration to the table takes several minutes. 

Another scenario is where the business changes "first and last name" columns to a single "full name", where your system can be changed quickly, but other team's components cannot be modified alongside your changes.

A more complex might be changing from an Entity-Attribute-Value system (i.e., property bags) to a concrete set of tables. So you introduce significant schema changes on a large database, and also a signifiant amount of code changes. 

This is a very loose and hand-waving kind of question, but it's hard to find information on this kind of topic as I don't quite know how to phrase questions. Hope I'm being clear!

* edit 1: in the case of slowly changing schemas, do you just try and backwards-port the data where possible, throw in switch statements (or differing classes) to handle various versions, etc?
## [12][Bad conventions at my new job, what to do?](https://www.reddit.com/r/rails/comments/f31lal/bad_conventions_at_my_new_job_what_to_do/)
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
