# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/g616hm/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/g616hm/personal_projects_show_off_your_own_project_andor/
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
## [2][Beginning with Rails and I have a few questions](https://www.reddit.com/r/rails/comments/g8xv6i/beginning_with_rails_and_i_have_a_few_questions/)
- url: https://www.reddit.com/r/rails/comments/g8xv6i/beginning_with_rails_and_i_have_a_few_questions/
---
Hi everyone, I started to learn Rails last week, and I am so happy with it. I managed to get from a "hello world" level, to an actually usable app in less than a week, great experience.
That said, I have plenty of questions, and I can't really find some clear answers online. I read a lot of threads here, and learned a lot, so I was hoping some of you might help answer these questions.

First, some context:
I'm an artist manager, and my rails app is basically a CRM for my business (a team of 4 people). It has some basic CRUD features + a calendar, and allows me to create invoices etc. There is not a tons of datas (few hundred of entries on each table) but there is a lot of relations between these datas (e.g: a manager has many artists, an artist has many performances, a performance has many events (rehearsals, setup, travel), a performance has an invoice etc.).

I would like to actually use this project for my business, but before involving all my team, it would help a lot to have some answers to the following questions:


1 - For my database I use SQLite3. Being file-based, I think it is perfect for a beginner like who might benefit from having this file under version control, and in case of problem I could just revert to the previous version. I'm not that confident in doing this easily and quickly with MySQL. So is there a real benefit from using a better database for my small business ? Will I really see some better performances given that there is not a lof of entries, and 4 simultaneous user maximum ?

2 - What is the best strategy for assets, should I manage these on a per-page basis or should I just send the whole pack on the first page load ?

3 - I installed a few gems, and now my JS asset file is around 3Mb. Is that considered normal for a rails app (there is not a tons of gems: wicked_pdf, shrine, materializejs, jquery, actioncable, cocoon) ?

4 - I plan to add to my app a "settings" page where I can define some settings such as menu color etc. Basically, it is an array of settings. 
What is the best way to store this kind of datas ? It feels weird to create a table given that I will only use the first entry of this table. What do you think?

Edit:
5 - For my app I made a dashboard which display some general statistics. In order to display all the required values, I need to pass to this view all my datas (e.g: Artist.all, Performance.all, Invoice.all etc.).
Does this affect performance ? 

Many thanks in advance !
## [3][What is the explanation of blob in ActiveStorage?](https://www.reddit.com/r/rails/comments/g8y2cj/what_is_the_explanation_of_blob_in_activestorage/)
- url: https://www.reddit.com/r/rails/comments/g8y2cj/what_is_the_explanation_of_blob_in_activestorage/
---
I have a model named document. In the migration of document model, I have a column named t.blob :file. There is an error about blob ib my model migration. "Unknown method". How do you implement my model document to activestorage?
## [4][Validates. Removing the url video.](https://www.reddit.com/r/rails/comments/g8xplp/validates_removing_the_url_video/)
- url: https://www.reddit.com/r/rails/comments/g8xplp/validates_removing_the_url_video/
---
Hi guys, I added this lines to validate the youtube video id.

It check only if the video is added and the length of the video is of 11 characters. Very easy.

        validate :check_videos_url!
    
        def check_videos_url!
          errors.add(
            :videos, :wrong_youtube_id
          ) if videos.any? &amp;&amp; videos.first.try(:url).length != 11
        end

But now the users reported me a bug.

If I create a page and I don't add any video -&gt; It works and creates the page.

If I create a page and I add a video of only 5 characters -&gt; It works, reporting me the error.

If I create a page and I add a video with 11 characters -&gt; It works and creates the page.

If I edit the page with no video, adding a new video -&gt; It works and edits the page.

If I edit the page with a video, adding a new video -&gt; It works and edits the page.

**If I edit the page with a video and I want to remove it -&gt; BUG. It say "it should be 11 characters" and I can't remove it.**

&amp;#x200B;

How to solve?
## [5][Optimal way to define Chrome + Capybara integration in 2020?](https://www.reddit.com/r/rails/comments/g8m0xb/optimal_way_to_define_chrome_capybara_integration/)
- url: https://www.reddit.com/r/rails/comments/g8m0xb/optimal_way_to_define_chrome_capybara_integration/
---
Hey guys!

I'm working on redoing some of  my company's feature test suite, and I wanted to update our Capybara  &amp; Webdriver configuration to match the latest recommended standards by the Capybara team.

I'm trying to define three types of drivers for my suite:

1. Local, regular driver that will be able to spin up Chrome sessions and test in the same GUI/window server as normal
2. Same thing as #1, but headless
3. Completely remote setup where we can run our feature tests on a  separate Selenium-Chrome node in the same network (for CI usage).

Our old preferences can be seen here: [https://github.com/arman000/marty/blob/master/spec/support/chromedriver.rb](https://github.com/arman000/marty/blob/master/spec/support/chromedriver.rb)

Now,  it's a bit all over the place, but one of the main things I wanted to  ask is what's the difference between using `add_preference` &amp;  `add_argument` to using  `Selenium::WebDriver::Remote::Capabilities.chrome`? Should I be using  that option set for locally running feature specs as well or just for setups #2 and #3?

Appreciate any help/feedback!
## [6][Social Login + Rails API + Mobile Client](https://www.reddit.com/r/rails/comments/g8g69p/social_login_rails_api_mobile_client/)
- url: https://www.reddit.com/r/rails/comments/g8g69p/social_login_rails_api_mobile_client/
---
I'm making this post in response to the last one I made where I was stuck on figuring out how to get social login working when dealing with a RoR backend and a native mobile app. ([https://www.reddit.com/r/rails/comments/g3s0v7/devise\_token\_auth\_omniauth/](https://www.reddit.com/r/rails/comments/g3s0v7/devise_token_auth_omniauth/))

I tried working with Omniauth but things got quite messy considering the number of redirects required to get everything going so I decided to go DIY and try write up my own solution.

The result is [https://gist.github.com/jesster2k10/ff96b5adbce72abae5fc603bd17c1843](https://gist.github.com/jesster2k10/ff96b5adbce72abae5fc603bd17c1843)

I go into a good bit of detail in the gist of the code and how everything works but to summarise it here:

* The user signs in with the native sdks on the mobile client
* The mobile SDK generates an access/id token
* A POST /identities/:provider request is sent with the token in the body
* The server fetches the user info from the token
* A new user is created based on that user info

The main benefit of this is that it's a much simpler implementation on the native side than setting up a web view and dealing with it the "traditional" way. However, if you are working with a Rails application or even an SPA, there's not much benefit to this over Omniauth so I would go with that.

I've written specs for about 65% of the code right now but just testing it with Postman shows it's working. I'll update the gist with the new specs as I write more of them

Hope this can help somebody as frustrated as I was.
## [7][Heroku Deploy Error Message](https://www.reddit.com/r/rails/comments/g8vxbl/heroku_deploy_error_message/)
- url: https://www.reddit.com/r/rails/comments/g8vxbl/heroku_deploy_error_message/
---
Just completed this tutorial app which was running near flawless, but heroku is being a jerk and not actually running it and instead asking me if I'm the owner of the app. What did I do wrong? (no it's not a carbon copy of the tutorial app).

  

# "We're sorry, but something went wrong.

If you are the application owner check the logs for more information."
## [8][profit loss graph in rails](https://www.reddit.com/r/rails/comments/g8vu2f/profit_loss_graph_in_rails/)
- url: https://www.reddit.com/r/rails/comments/g8vu2f/profit_loss_graph_in_rails/
---
 have income,expense, user controller i want to plot a graph of currrent user profit i calculated it in model.  
def total\_expense  
expenses.pluck(:amount).sum  
end  
def total\_incomes  
incomes.pluck(:amount).sum  
end  
def profit\_loss  
total\_incomes - total\_expense  
end  
how i plot a graph of profit\_loss its value is not save in data base
## [9][Naming convention: folder containing algorithms that look up and/or calculate values](https://www.reddit.com/r/rails/comments/g8skfa/naming_convention_folder_containing_algorithms/)
- url: https://www.reddit.com/r/rails/comments/g8skfa/naming_convention_folder_containing_algorithms/
---
Do any of you contain such a folder? For example, I want to write an algorithm that determines the nearest year a vehicle could be based on VIN number (so an amateur decoder of sorts). Where would I put this logic? So far I only intend to use this method in multiple controllers, so is it simply a controller concern?

Edit: Or maybe I should just make a folder named "decoders" and name the file "vin\_decoder.rb"?
## [10][Storing Videos For My Rails Application](https://www.reddit.com/r/rails/comments/g8hbmy/storing_videos_for_my_rails_application/)
- url: https://www.reddit.com/r/rails/comments/g8hbmy/storing_videos_for_my_rails_application/
---
Hello!

A quick question regarding **storing videos** that I created for my rails application.

I'm wondering how I would be able to store videos in a database. From what I gathered, I saw that a lot of people are saying to store the media in the "filesystem" rather than in the database directly.

I'm wondering if the right way of doing this is by storing my videos on a server and then copying the path to that file into the database.

Also, I saw that Amazon S3 offers buckets. I only have about 1TB of videos that need to be uploaded. Would this be a better alternative?

**The services I'm considering are:**  
\- Amazon S3  
\- Digital Ocean Spaces Object Storage

Am I on the right path here? Or am I way off?

Thanks in advance guys!
## [11][Beefing security of Rails API](https://www.reddit.com/r/rails/comments/g7zhjp/beefing_security_of_rails_api/)
- url: https://www.reddit.com/r/rails/comments/g7zhjp/beefing_security_of_rails_api/
---
What are the best ways to make a rails API secure?

I'm considering whitelisting my client's IP address. How useful is hashing or encrypting the json data? Any other ideas?
