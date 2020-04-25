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
## [2][AuthTrail is a gem for tracking Devise login activity](https://www.reddit.com/r/rails/comments/g7e20w/authtrail_is_a_gem_for_tracking_devise_login/)
- url: https://www.reddit.com/r/rails/comments/g7e20w/authtrail_is_a_gem_for_tracking_devise_login/
---
I was looking for a way how to provide my users with an add additional level of visibility into their login activity and came across [AuthTrail](https://github.com/ankane/authtrail). Brilliant gem , opensourced by Instacart, that add to your db new table \`LoginActivity\` and record is created every time a user tries to login. You can then use this information to detect suspicious behavior.
## [3][Another way to handle complexity in Rails application](https://www.reddit.com/r/rails/comments/g7tkzq/another_way_to_handle_complexity_in_rails/)
- url: https://www.reddit.com/r/rails/comments/g7tkzq/another_way_to_handle_complexity_in_rails/
---
 [https://github.com/andriy-baran/steel\_wheel/wiki/Getting-started-on-Rails](https://github.com/andriy-baran/steel_wheel/wiki/Getting-started-on-Rails)
## [4][a better way to do git commit -m "wip"](https://www.reddit.com/r/rails/comments/g7dj8g/a_better_way_to_do_git_commit_m_wip/)
- url: https://www.reddit.com/r/rails/comments/g7dj8g/a_better_way_to_do_git_commit_m_wip/
---
&gt;git commit -m "wip"

I find myself doing that \^ a lot, specially after I starting cleanup up a of branch that was put together rather hastily.

Using this function

    function c() {
      str="$*" # https://unix.stackexchange.com/a/197794
      git add --all
      git commit -m "$str"
      git log --oneline -n2
    }

I can commit like this:

    c wip

c stands for commit,  
no quotes needed around the commit message, and git add is done for you already.,  
and you can use multiple words obviously
## [5][Developing Rails/React app on remote ubuntu machine?](https://www.reddit.com/r/rails/comments/g7gp5n/developing_railsreact_app_on_remote_ubuntu_machine/)
- url: https://www.reddit.com/r/rails/comments/g7gp5n/developing_railsreact_app_on_remote_ubuntu_machine/
---
I have everything setup and built on an AWS EC2 Ubuntu 18.04 machine for remote development/testing/production. On my local development I'm using foreman to combine \`rails s\` and \`bin/webpack-dev-server\` on port 5000. How does this work using a remote machine and having an accessible URL to view app in browser? I'm using ssh to access the remote machine. 
## [6][Whats the best way to send a series of emails to users given each mail must be sent at a given time](https://www.reddit.com/r/rails/comments/g77lfv/whats_the_best_way_to_send_a_series_of_emails_to/)
- url: https://www.reddit.com/r/rails/comments/g77lfv/whats_the_best_way_to_send_a_series_of_emails_to/
---
I am trying to find a solution on how could this be best handled

I wanted to send emails to users at different times on a day and each user will be receiving multiple emails on the same day. I could have a background job and send an email but as far as I know I should run some check every few minutes to see whether an email should be sent to the user now and send it  but thats will query the db multiple times to find how many users should receive a mail now and then process all those users data to send a mail.

The other option would be run every few times and take the next few mails to be sent and schedule it to be sent but the problem here is when I make a change in the time on email to be sent for a user. I should be able to cancel that enqueued one and add it newly and once we fetch a list and enqueue there can be new user coming in at later point and I should also enqueue them

So trying to find what are the other options with which this can be done.
## [7][Whats the optimal maintainable &amp; cost-efficient way to setup a cronjob for alerts for a todo app?](https://www.reddit.com/r/rails/comments/g7detz/whats_the_optimal_maintainable_costefficient_way/)
- url: https://www.reddit.com/r/rails/comments/g7detz/whats_the_optimal_maintainable_costefficient_way/
---
I'm building a simple todo app in rails 6.

I have work_item that has a title, deadline and finished_at. I want to send an alert on 0, 5 and 15 minutes from the deadline, if finished_at is nil. I've looked at using ActiveJob with Sidekiq + Redis. Now looking at the whenever gem like this:
```
every 1.minute do 
  rake work_item:alert
end
```

```
namespace :work_item do
  task :alert do

     // Some pseudo code
    minutes = [0, 5, 15]
    time_left = work_item.deadline - DateTime.current
    if work_item.finished_at.nil?
      minutes.each do |number|
        if time_left == number.minutes
          AlertJob.perform_later(work_item)
        end
      end
    end
  end
end
```


Is this the optimal maintainable &amp; cost-efficient way to setup a cronjob for a todo app?
## [8][Multiple apps on a VPS?](https://www.reddit.com/r/rails/comments/g77ool/multiple_apps_on_a_vps/)
- url: https://www.reddit.com/r/rails/comments/g77ool/multiple_apps_on_a_vps/
---
Noob alert! 

I’ve only recently stepped into the world of Docker containers and I was wondering if part of scope of containerisation is to run multiple unrelated apps in the form of multiple containers, one in each. 

To make optimal use of the VPS one rents. 

For example, a development VPS which has multiple unrelated apps running with accessible endpoints instead of renting multiple VPSs 

Because an app would rarely, if ever, utilise the full potential of even a basic server when it’s for development or personal use

If yes, how does one do this? Could anyone point me to some relavent resources?
## [9][Can you write more than one condition in an if statement on one line?](https://www.reddit.com/r/rails/comments/g7bmw1/can_you_write_more_than_one_condition_in_an_if/)
- url: https://www.reddit.com/r/rails/comments/g7bmw1/can_you_write_more_than_one_condition_in_an_if/
---
I'm looking to make my program do something like this, what is the right way to do this? 

So for example:

If var.count == 5 then - 4
## [10][How to create dynamic changeable wizard for registration](https://www.reddit.com/r/rails/comments/g708fd/how_to_create_dynamic_changeable_wizard_for/)
- url: https://www.reddit.com/r/rails/comments/g708fd/how_to_create_dynamic_changeable_wizard_for/
---
What would be the best practice or method to create dynamic wizards for registrations. This is for an event registration app. Users can register for an event, but each event has different steps. Some might have multiple release forms, other an option to add buying a T-shirt, another to ask some random question. What’s the best way to have an admin create these steps and then when a user registers, they go through this process?

I believe I can figure out the models and have a model that admins use to fill out the “store “ module, selecting the the up-sell, etc. and I looked at Wicked for multi-step forms, but I feel like I’m missing some major. 

Any advice is greatly appreciated.
## [11][Looking to get started with Rails 6](https://www.reddit.com/r/rails/comments/g6owja/looking_to_get_started_with_rails_6/)
- url: https://www.reddit.com/r/rails/comments/g6owja/looking_to_get_started_with_rails_6/
---
I have been putting off learning Rails 6, specifically webpacker. I tried it a few times but kind of got distracted or found that I need to have a lot of background information on it to get started. But now with the lockdown I decided it might be the best time to build a test app and get started with Rails 6.

I am looking for some tutorial that helps me get started. Something does not focus on the basics of rails but instead focused on what is new between Rails 5 and Rails 6 and how to get started and use webpacker. Any videos/tutorials would be appreciated.

Thanks in advance.
