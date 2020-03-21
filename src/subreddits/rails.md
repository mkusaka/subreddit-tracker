# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/fgx7fz/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/fgx7fz/personal_projects_show_off_your_own_project_andor/
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
## [2][Are Companies still hiring?](https://www.reddit.com/r/rails/comments/fm6nwe/are_companies_still_hiring/)
- url: https://www.reddit.com/r/rails/comments/fm6nwe/are_companies_still_hiring/
---
I am a part time Rails developer searching for a full time role. It looks like I am facing some tough trends with companies everywhere closing and doing layoffs. 
I am seeing Jr developers getting laid off everywhere. Do I stand a chance in the job search right now with less than a year of professional experience.
## [3][Wanting to help other RoR developers](https://www.reddit.com/r/rails/comments/fm6vfn/wanting_to_help_other_ror_developers/)
- url: https://www.reddit.com/r/rails/comments/fm6vfn/wanting_to_help_other_ror_developers/
---
I really enjoy helping other programmers so i'd like to offer my time to anyone is stuck on something in Rails. I've been developing in RoR for 6 years and own a small custom software development company working with some pretty large projects.

I mostly enjoy working with with the backend focusing on code design and testing in rspec, but if I can help you in something else, I'm happy to give it a shot. Even if you're brand new and it's just a discussion to ask some questions and maybe help connect the dots for you.

Some requirements:

* I'd like to video conference the session (you sharing your screen) so we can record it and pop it up on YouTube for others. I'd like to keep it less than 2 hours.
* You would need a mic/camera and good internet connection
* I think it would be best suited those who are a beginner/intermediate in the framework. Preferably I'd like us to focus on an actual problem in the session rather than it turning into a straight video tutorial.
* It would be helpful to get an idea on what the challenge you're facing, just so I know I can actually help you and not waste our time

If you're keen just leave a comment or send me a PM with your challenge and I'll be in touch. If it's not within my skillset maybe someone else might offer some time!

&amp;#x200B;

EDIT:  I'm not looking to charge anything, I'm just trying to stay busy, take a break from personal projects, and avoid video games.
## [4][How to make sure you are integrating google/outlook calendar correctly with your product](https://www.reddit.com/r/rails/comments/flyeuv/how_to_make_sure_you_are_integrating/)
- url: https://www.reddit.com/r/rails/comments/flyeuv/how_to_make_sure_you_are_integrating/
---
I have worked on calendar integrations in my current role for quite some time now and have a deep understanding of how they work. Hence, I thought it would be nice to help developers who would face this challenge in their own workplace.

Google calendar has decent documentation for most of its products and supports most of the libraries, even then, it's confusing enough to take the wrong approach while integrating with google calendar.

When I set out to build calendar integration for one of my products, I had never thought that it would end up becoming so complicated to do things which I had initially thought to be simple.

Product Requirements

1. During Signup/Login ask for permission to view/edit/create the events in their primary google calendar.
2. Fetch the calendar events(both recurring and 1 time) and show them to the users so that they can add more details to it(specific to the product).
3. Keep in sync with any new events that are created.
4. Sync existing events between the product and the calendar.
5. Updating the correct recurring events in google calendar for any updates done to them in my product.

Remember that, my product required 2-way sync of events.

The Incorrect Approach

During my research, most of the blog posts I read pointed to using the [Calendar Events List API](https://developers.google.com/calendar/v3/reference/events/list) and my brain started to work in that direction. I took the following approach which is inefficient and unscalable -

1. Integrated the Calendar Events List API which got me the event information for all the upcoming events.
2. Saved the ones which matched as per the algorithm of the product
3. If the user performed any update to these events then update the calendar events accordingly(changes to date/time/description)
4. Write a Cron job which runs every 30 minutes to check if any new events were added/existing events changed(biggest blunder) for all the users in the database.

The Correct Approach

Soon, I realised that this was not the most efficient and scalable way to achieve the results I wanted.  
I dug deeper into the google documentation to realize that there is a [synch API](https://developers.google.com/calendar/v3/sync) which helps in incremental synch of data and works as a [webhook](https://en.wikipedia.org/wiki/Webhook). This is what made complete sense and would fit my requirements perfectly.  
I am going to talk about this approach in a new series. Stay Tuned

Would be nice to hear your comments.

I plan to write a series about how to actually do it correctly as well
## [5][Auto creating memberships and assigning roles](https://www.reddit.com/r/rails/comments/fm3qz0/auto_creating_memberships_and_assigning_roles/)
- url: https://www.reddit.com/r/rails/comments/fm3qz0/auto_creating_memberships_and_assigning_roles/
---
I have an application where I want user to create workspaces/companies. I have a many\_through join model setup that's working fine. I have the rolify gem handling roles for me. They way I have setup creating a company is as follows.

    def create
    @company = Company.new(company_params)
    if @company.save
    .... things going on after save...
            
    @company.memberships &lt;&lt; Membership.create(user_id: current_user.id,     company_id: @company)
            Membership.last.add_role :owner, @company
    end

This way when a user creates a company the membership is created and set to the right company and user. My concern is the role adding part. Calling `Membership.last`, doesn't seem like a proper solution to handle this. I have tried things with @ sings etc but that didnt seem to work..

&amp;#x200B;

How to handle this correctly?
## [6][How to block access from most of the application until users finish stripe payments?](https://www.reddit.com/r/rails/comments/fm4w26/how_to_block_access_from_most_of_the_application/)
- url: https://www.reddit.com/r/rails/comments/fm4w26/how_to_block_access_from_most_of_the_application/
---
Not sure how much explanation is needed in conjunction with the title, but I have a signup process where users can be a parent or a student of a parent. They have slightly different processes to complete based on their selections. 

**The flows goes like this for students:**
Select create new account -&gt;
choose student account option -&gt;
fill out new user form -&gt;
submit -&gt;
stripe payment on stripe website -&gt;
should be able to see their dashboard 

**&amp; for parents:**
Select create new account -&gt;
choose parent account option -&gt;
fill out new user form -&gt;
submit -&gt;
add children to your account -&gt;
finish that process -&gt;
stripe payment on stripe website -&gt;
should be able to see their dashboard 

My main concern is that they can access their accounts by typing in the url or by clicking on their profile before they finish the process, and I want to make that impossible until they’ve paid. They’ll have stripe info rows on their table after paying, and I assume I can use that to make it impossible to see things. I have pundit for authorization, and a home rolled authorization (I’ve been pushing for devise, I know, I know). This seems like it should go in an authorization step, but I’m not totally sure. 

I’d be happy to provide more info, but I didn’t want to have a wall of text. Thanks in advance for any help!
## [7][Will a nosql database more suitable when it comes to a personal app like a todo list?](https://www.reddit.com/r/rails/comments/flxzre/will_a_nosql_database_more_suitable_when_it_comes/)
- url: https://www.reddit.com/r/rails/comments/flxzre/will_a_nosql_database_more_suitable_when_it_comes/
---
What I mean by a personal app is that when a user logs in, all they can see is their data, so they don't need to know anything else in the database.

What I understand is that if I use a relational database like postgres, to do lists of different users (for example) will be meshed in a single to do table, so I'll have to query by user id. Is there a more efficient way to do this, will mongodb better to solve this type of problem?
## [8][Hello! to all RoR experts here](https://www.reddit.com/r/rails/comments/fm95uk/hello_to_all_ror_experts_here/)
- url: https://www.reddit.com/r/rails/comments/fm95uk/hello_to_all_ror_experts_here/
---
How much should I pay for a project very similar to this website?   


[https://www.investagrams.com/](https://www.investagrams.com/)  


I think you can try and log in with your Facebook account. Thank you!!
## [9][Where to store file paths](https://www.reddit.com/r/rails/comments/flxmag/where_to_store_file_paths/)
- url: https://www.reddit.com/r/rails/comments/flxmag/where_to_store_file_paths/
---
I work on a rails app that has a number of hard coded paths. For example, on app initialization a class might load a YAML file like this:

@@variable = YAML.load_file('path/to/the/file.yaml')

Is this bad practice? I have been taught it is better to place paths in an environment variable, but my develop lead disagrees. I am trying to write tests that run in many different environments (dev, test server, etc.). This is an example of our configuration making the process more difficult than it should be. Is there a better solution?

Thanks!
## [10]["The Rails way" to deal with a lot of tables with custom columns](https://www.reddit.com/r/rails/comments/flw9ly/the_rails_way_to_deal_with_a_lot_of_tables_with/)
- url: https://www.reddit.com/r/rails/comments/flw9ly/the_rails_way_to_deal_with_a_lot_of_tables_with/
---
Hi,

I have very recently started learning Rails and I feel, after quite a few tutorials, that I am ready to tackle more complicated projects to learn Rails beyond basic stuff.

The project that I am building is some kind of inventory tracking system. The most challenging thing right now is structuring database because items have different columns based on their type. For example book has ISBN and screen resolution. I have tried to do some research on that topic but I don't even know what to search to find answers. I am currently thinking about having table with all universal columns and separate tables for each type with custom fields.

How should I approach this problem in "The Rails way". Should I create model for each type and is there a way to add new types (with their tables and models) without explicitly typing \`rails g model...\` on command line?

Thanks for helping me learn,

Bartol
## [11][Postgres client / server version mismatch](https://www.reddit.com/r/rails/comments/flkqlo/postgres_client_server_version_mismatch/)
- url: https://www.reddit.com/r/rails/comments/flkqlo/postgres_client_server_version_mismatch/
---
I’m toying with Rails in containers using [this page](https://docs.docker.com/compose/rails/) as my guide.

I picked the latest Ruby, the latest Debian, and the latest Postgres (12) and bumped into a problem.

I create a table and then from the Rails container, go into “rails dbconsole” and do “\d accounts” (my table) and I get [this error](https://stackoverflow.com/questions/58461178/how-to-fix-error-column-c-relhasoids-does-not-exist-in-postgres)

Sure enough, “psql —version” from inside the Rails container says 11 while the DB container says 12.

My question is if I can somehow force the error to happen (and perhaps be more obvious) when Rails (using a version 11 library) connects to a version 12 DB server?
