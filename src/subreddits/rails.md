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
## [2][Beefing security of Rails API](https://www.reddit.com/r/rails/comments/g7zhjp/beefing_security_of_rails_api/)
- url: https://www.reddit.com/r/rails/comments/g7zhjp/beefing_security_of_rails_api/
---
What are the best ways to make a rails API secure?

I'm considering whitelisting my client's IP address. How useful is hashing or encrypting the json data? Any other ideas?
## [3][what do you think about my ruby/js coding?](https://www.reddit.com/r/rails/comments/g8crzn/what_do_you_think_about_my_rubyjs_coding/)
- url: https://www.reddit.com/r/rails/comments/g8crzn/what_do_you_think_about_my_rubyjs_coding/
---
&amp;#x200B;

im trying to get a code review done for [this code i wrote](https://codereview.stackexchange.com/q/241237/152349)   
can pay $ if thats what it takes
## [4][Another way to handle complexity in Rails application](https://www.reddit.com/r/rails/comments/g7tkzq/another_way_to_handle_complexity_in_rails/)
- url: https://www.reddit.com/r/rails/comments/g7tkzq/another_way_to_handle_complexity_in_rails/
---
 [https://github.com/andriy-baran/steel\_wheel/wiki/Getting-started-on-Rails](https://github.com/andriy-baran/steel_wheel/wiki/Getting-started-on-Rails)
## [5][[Newbie] Help Rails 6](https://www.reddit.com/r/rails/comments/g7x236/newbie_help_rails_6/)
- url: https://www.reddit.com/r/rails/comments/g7x236/newbie_help_rails_6/
---
Hi,

First giving a bit of introduction, thanks to Corona have enough time to start learning rails and its components.

&amp;#x200B;

I'm trying to incorporate ajax calls into my rails 6 newbie app ( but wans't able to understand this instruction [here](https://guides.rubyonrails.org/working_with_javascript_in_rails.html#unobtrusive-javascript) )

&amp;#x200B;

Was able to play with:

format.html

format.js

    respond_to do |format|
            if @problem.save
              format.html { redirect_to :root
                flash[:success] = "Added #{@problem.description}}
              format.json { render :show, status: :created, location: @ticket }
              format.js   { render :js =&gt; "window.location='#{ problems_path }'"  
              flash[:success] = "Added #{@problem.description} by #{current_user.email}" } 
            else
              format.js {}
              format.html { render 'problems/new' }
              format.json { render json: @problem.errors, status: :unprocessable_entity }          
            end
          end 

I understood know how html/js works when the request is being process from views and how to create a proper response from server.

Now, I want to completely understand as well the ajax processing.

&amp;#x200B;

As per link, when form\_with model:@model sets the processing as AJAX but I think it's just the same as JS, now, it's a bit cloudy on my side how can I exacly integrate AJAX scripts.

&amp;#x200B;

Any newbie help?
## [6][Creating Recipient model before payment session](https://www.reddit.com/r/rails/comments/g7wirm/creating_recipient_model_before_payment_session/)
- url: https://www.reddit.com/r/rails/comments/g7wirm/creating_recipient_model_before_payment_session/
---
Hi all, I have created a simple checkout for my rails application. For now, it works like this. A customer clicks the button below:

    &lt;div class="buynow-wrapper has-text-centered"&gt;
        &lt;%= button_to payments_create_path, params: { id: @product.id}, remote: true,        class: 'button is-info button-shadow' do %&gt;
        &lt;span class="icon"&gt;
            &lt;i class="fas fa-check"&gt;&lt;/i&gt;
        &lt;/span&gt;
        &lt;span&gt;Pay&lt;/span&gt;
        &lt;% end %&gt;
    &lt;/div&gt;

This will create a payment session(stripe), created in my Payments Controller. Now I want have a form with an e-mail just before submitting this all to my payment session create action. Can a button\_to go to two post paths? Im kinda stuck where and how to do this.  
Would it be possible to send the form data, in this case en e-mailadress, directly to my create action in some variable? Main goal is having the e-mailadres available from the person who pressen 'Pay'.
## [7][AuthTrail is a gem for tracking Devise login activity](https://www.reddit.com/r/rails/comments/g7e20w/authtrail_is_a_gem_for_tracking_devise_login/)
- url: https://www.reddit.com/r/rails/comments/g7e20w/authtrail_is_a_gem_for_tracking_devise_login/
---
I was looking for a way how to provide my users with an add additional level of visibility into their login activity and came across [AuthTrail](https://github.com/ankane/authtrail). Brilliant gem , opensourced by Instacart, that add to your db new table \`LoginActivity\` and record is created every time a user tries to login. You can then use this information to detect suspicious behavior.
## [8][a better way to do git commit -m "wip"](https://www.reddit.com/r/rails/comments/g7dj8g/a_better_way_to_do_git_commit_m_wip/)
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
## [9][Whats the best way to send a series of emails to users given each mail must be sent at a given time](https://www.reddit.com/r/rails/comments/g77lfv/whats_the_best_way_to_send_a_series_of_emails_to/)
- url: https://www.reddit.com/r/rails/comments/g77lfv/whats_the_best_way_to_send_a_series_of_emails_to/
---
I am trying to find a solution on how could this be best handled

I wanted to send emails to users at different times on a day and each user will be receiving multiple emails on the same day. I could have a background job and send an email but as far as I know I should run some check every few minutes to see whether an email should be sent to the user now and send it  but thats will query the db multiple times to find how many users should receive a mail now and then process all those users data to send a mail.

The other option would be run every few times and take the next few mails to be sent and schedule it to be sent but the problem here is when I make a change in the time on email to be sent for a user. I should be able to cancel that enqueued one and add it newly and once we fetch a list and enqueue there can be new user coming in at later point and I should also enqueue them

So trying to find what are the other options with which this can be done.
## [10][Developing Rails/React app on remote ubuntu machine?](https://www.reddit.com/r/rails/comments/g7gp5n/developing_railsreact_app_on_remote_ubuntu_machine/)
- url: https://www.reddit.com/r/rails/comments/g7gp5n/developing_railsreact_app_on_remote_ubuntu_machine/
---
I have everything setup and built on an AWS EC2 Ubuntu 18.04 machine for remote development/testing/production. On my local development I'm using foreman to combine \`rails s\` and \`bin/webpack-dev-server\` on port 5000. How does this work using a remote machine and having an accessible URL to view app in browser? I'm using ssh to access the remote machine. 
## [11][Whats the optimal maintainable &amp; cost-efficient way to setup a cronjob for alerts for a todo app?](https://www.reddit.com/r/rails/comments/g7detz/whats_the_optimal_maintainable_costefficient_way/)
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
