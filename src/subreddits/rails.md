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
## [2][Whats the best way to send a series of emails to users given each mail must be sent at a given time](https://www.reddit.com/r/rails/comments/g77lfv/whats_the_best_way_to_send_a_series_of_emails_to/)
- url: https://www.reddit.com/r/rails/comments/g77lfv/whats_the_best_way_to_send_a_series_of_emails_to/
---
I am trying to find a solution on how could this be best handled

I wanted to send emails to users at different times on a day and each user will be receiving multiple emails on the same day. I could have a background job and send an email but as far as I know I should run some check every few minutes to see whether an email should be sent to the user now and send it  but thats will query the db multiple times to find how many users should receive a mail now and then process all those users data to send a mail.

The other option would be run every few times and take the next few mails to be sent and schedule it to be sent but the problem here is when I make a change in the time on email to be sent for a user. I should be able to cancel that enqueued one and add it newly and once we fetch a list and enqueue there can be new user coming in at later point and I should also enqueue them

So trying to find what are the other options with which this can be done.
## [3][Multiple apps on a VPS?](https://www.reddit.com/r/rails/comments/g77ool/multiple_apps_on_a_vps/)
- url: https://www.reddit.com/r/rails/comments/g77ool/multiple_apps_on_a_vps/
---
Noob alert! 

I’ve only recently stepped into the world of Docker containers and I was wondering if part of scope of containerisation is to run multiple unrelated apps in the form of multiple containers, one in each. 

To make optimal use of the VPS one rents. 

For example, a development VPS which has multiple unrelated apps running with accessible endpoints instead of renting multiple VPSs 

Because an app would rarely, if ever, utilise the full potential of even a basic server when it’s for development or personal use

If yes, how does one do this? Could anyone point me to some relavent resources?
## [4][How to create dynamic changeable wizard for registration](https://www.reddit.com/r/rails/comments/g708fd/how_to_create_dynamic_changeable_wizard_for/)
- url: https://www.reddit.com/r/rails/comments/g708fd/how_to_create_dynamic_changeable_wizard_for/
---
What would be the best practice or method to create dynamic wizards for registrations. This is for an event registration app. Users can register for an event, but each event has different steps. Some might have multiple release forms, other an option to add buying a T-shirt, another to ask some random question. What’s the best way to have an admin create these steps and then when a user registers, they go through this process?

I believe I can figure out the models and have a model that admins use to fill out the “store “ module, selecting the the up-sell, etc. and I looked at Wicked for multi-step forms, but I feel like I’m missing some major. 

Any advice is greatly appreciated.
## [5][Looking to get started with Rails 6](https://www.reddit.com/r/rails/comments/g6owja/looking_to_get_started_with_rails_6/)
- url: https://www.reddit.com/r/rails/comments/g6owja/looking_to_get_started_with_rails_6/
---
I have been putting off learning Rails 6, specifically webpacker. I tried it a few times but kind of got distracted or found that I need to have a lot of background information on it to get started. But now with the lockdown I decided it might be the best time to build a test app and get started with Rails 6.

I am looking for some tutorial that helps me get started. Something does not focus on the basics of rails but instead focused on what is new between Rails 5 and Rails 6 and how to get started and use webpacker. Any videos/tutorials would be appreciated.

Thanks in advance.
## [6][I've published a new post about how Trix works from the inside](https://www.reddit.com/r/rails/comments/g6lxa8/ive_published_a_new_post_about_how_trix_works/)
- url: https://www.reddit.com/r/rails/comments/g6lxa8/ive_published_a_new_post_about_how_trix_works/
---
This  post is going into Trix internals. We'll break down its core components  and explain what their function is and how they interrelate. How Trix works:

[https://www.thatweeklytech.com/posts/16-how-trix-works](https://www.thatweeklytech.com/posts/16-how-trix-works)

Thank you for reading! Hope this helps.
## [7][Unpermitted parameters after trying to AJAXify my request](https://www.reddit.com/r/rails/comments/g6wb0q/unpermitted_parameters_after_trying_to_ajaxify_my/)
- url: https://www.reddit.com/r/rails/comments/g6wb0q/unpermitted_parameters_after_trying_to_ajaxify_my/
---
Hi everyone,  


I've been hiding from AJAX for too long and I finally decided to face my fear, but it isn't going well.  


So I have this route:  
`/venues?commit=Search&amp;order_by=rating`  


And previously everything was working fine without AJAX. But now with AJAX, all of a sudden I'm getting an unpermitted parameters error. Here's the error:  


    Parameters: {"commit"=&gt;"Search", "location"=&gt;"", "order_by"=&gt;"rating", "utf8"=&gt;"✓"}
    Unpermitted parameters: :commit, :utf8
    Unpermitted parameters: :commit, :utf8
    Unpermitted parameters: :commit, :utf8
    Unpermitted parameters: :commit, :utf8
    Unpermitted parameters: :commit, :utf8
    Unpermitted parameters: :commit, :utf8
    Unpermitted parameters: :commit, :utf8
    Unpermitted parameters: :commit, :utf8
    Unpermitted parameters: :commit, :location, :order_by, :utf8

This is the request url which I'm doing in my fetch request.  
`/venues?commit=Search&amp;location=&amp;order_by=rating&amp;utf8=%E2%9C%93`  


This is my fetch request:

    async executeAjaxRequest(url) {
        await fetch(url, { headers: { accept: "application/json" } })
          .then((response) =&gt; response.json())
          .then((data) =&gt; {
            console.log('Success:', data);
          })
      }

When I just go to that URL in my browser and get the HTML response everything works fine still. So I don't understand why the parameters are suddenly unpermitted?  
Thanks in advance.
## [8][How to make changes to a production app](https://www.reddit.com/r/rails/comments/g6vx0t/how_to_make_changes_to_a_production_app/)
- url: https://www.reddit.com/r/rails/comments/g6vx0t/how_to_make_changes_to_a_production_app/
---
Hey y'all, Rails student here, after a few months of tutorials I'm done with them, figured it was time to get my hands dirty so I'm working on an MVP app and I've got a bunch of friends testing it out and suggesting tweaks, what works what doesn't, etc... so ... but I keep running into an issue that when I need to make changes to the model(s) the only thing I know is to drop the db and recreate it, this cannot be how it is done in a real production app, so I need some advice how do I make changes without absolutely destroying the db every single time, my focus group is getting annoyed HAHA.

Thanks in advance.
## [9][Should credentials/production.yml.enc be on the git repository?](https://www.reddit.com/r/rails/comments/g6ts6u/should_credentialsproductionymlenc_be_on_the_git/)
- url: https://www.reddit.com/r/rails/comments/g6ts6u/should_credentialsproductionymlenc_be_on_the_git/
---
That's my question, should `credentials/production.yml.enc` be on the git repository? I know it shouldn't matter if the key isn't there, but maybe I'm not considering something.
## [10][Catchup subscriptions with Rails Event Store](https://www.reddit.com/r/rails/comments/g6j09z/catchup_subscriptions_with_rails_event_store/)
- url: https://www.reddit.com/r/rails/comments/g6j09z/catchup_subscriptions_with_rails_event_store/
---
[https://blog.arkency.com/catchup-subscriptions-with-rails-event-store/](https://blog.arkency.com/catchup-subscriptions-with-rails-event-store/)
## [11][Is it ok to post blogpost links here?](https://www.reddit.com/r/rails/comments/g6m7yn/is_it_ok_to_post_blogpost_links_here/)
- url: https://www.reddit.com/r/rails/comments/g6m7yn/is_it_ok_to_post_blogpost_links_here/
---
I posted a link to my blogpost but it was removed for an unknown reason:

https://preview.redd.it/qd2kq28ifku41.png?width=751&amp;format=png&amp;auto=webp&amp;s=10569ed2523a34d3b5b14314162e91d221054cb6

I can't see many links on this reddit (contrary to r/rubyonrails where there's more links and less subscribers). Is there a rule I don't know about? I could only find this: [https://old.reddit.com/r/rails/wiki/how\_to\_ask\_for\_help](https://old.reddit.com/r/rails/wiki/how_to_ask_for_help)
