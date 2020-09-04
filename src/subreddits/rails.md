# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/
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
## [2][iCallendar issue](https://www.reddit.com/r/rails/comments/imdmnc/icallendar_issue/)
- url: https://www.reddit.com/r/rails/comments/imdmnc/icallendar_issue/
---
I am having some issues with icalendar gem,  it is adding events   in UTC+00 and I wanted it to be readjusted to match users  timezone, how can I do it?
## [3][missing /tmp/passenger.xxxxx directory with standalone passenger run by logrotate](https://www.reddit.com/r/rails/comments/imcai3/missing_tmppassengerxxxxx_directory_with/)
- url: https://www.reddit.com/r/rails/comments/imcai3/missing_tmppassengerxxxxx_directory_with/
---
Hi everyone.

I'm running a phusion passenger standalone instance and everything works fine except the automatic restart after the log rotation which occurs every night. When I wake up, the "built-in" nginx reports a 502 bad gateway.

I'm using a script to stop/start passenger and when I restart in an interactive ssh session, everything works fine. However, when logrotate executes the same script, passenger restarts but there's a missing directory in /tmp so the built-in nginx can't speak with passenger.

I'm at a loss. I don't know what's wrong and what to check next.

My script is just doing this :

    #!/bin/bash
    su -l myuser /home/myuser/scripts/passenger_stop.sh
    su -l myuser /home/myuser/scripts/passenger_start.sh

passenger\_start.sh :

    #!/bin/bash
    export RBENV_ROOT="/opt/rbenv"
    export PATH="/opt/rbenv/plugins/ruby-build/bin:/opt/rbenv/bin:$PATH"
    eval "$(rbenv init -)"
    alias node=nodejs
    SOCKET="/home/myuser/passenger-nginx.socket"
    PID_FILE="/home/myuser/project/shared/pids/passenger.pid"
    cd /home/myuser/project/current
    bundle exec passenger start -e production --socket=${SOCKET} -d --pid-file ${PID_FILE}

The passenger.log file :

    [ N 2020-09-04 00:05:31.7264 20421/T1 age/Cor/CoreMain.cpp:1340 ]: Starting Passenger core...
    [ N 2020-09-04 00:05:31.7265 20421/T1 age/Cor/CoreMain.cpp:256 ]: Passenger core running in multi-application mode.
    [ N 2020-09-04 00:05:31.7350 20421/T1 age/Cor/CoreMain.cpp:1015 ]: Passenger core online, PID 20421
    [ N 2020-09-04 00:05:33.8468 20421/T5 age/Cor/SecurityUpdateChecker.h:519 ]: Security update check: no update found (next check in 24 hours)
    /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler/shared_helpers.rb:71:in `getwd': No such file or directory - getcwd (Errno::ENOENT)
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler/shared_helpers.rb:71:in `pwd'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler/shared_helpers.rb:71:in `block in pwd'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/monitor.rb:235:in `mon_synchronize'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler/shared_helpers.rb:70:in `pwd'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler/shared_helpers.rb:272:in `search_up'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler/shared_helpers.rb:155:in `print_major_deprecations!'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler.rb:103:in `setup'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/bundler/setup.rb:20:in `&lt;top (required)&gt;'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
            from /opt/rbenv/versions/2.6.6/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
    2020/09/04 00:05:51 [crit] 20441#0: *3 connect() to unix:/passenger_core failed (2: No such file or directory) while connecting to upstream, client: unix:, server: _, request: "GET / HTTP/1.0", upstream: "passenger:unix:/passenger_core:", host: "myhost"
    2020/09/04 00:05:52 [crit] 20441#0: *5 connect() to unix:/tmp/passenger.Bh5wxh4/agents.s/core failed (2: No such file or directory) while connecting to upstream, client: unix:, server: _, request: "GET / HTTP/1.0", upstream: "passenger:unix:/tmp/passenger.Bh5wxh4/agents.s/core:", host: "myhost"

The last line repeats for every request.

The process command lines (for nginx and PassengerAgent) are referencing this "/tmp/passenger.something" directory but it doesn't exist.

I just have to execute the restart script again (in ssh interactive, not in my logrotate script) and both the command lines and the tmp directory are ok. The command lines are then exactly the same except for the random tmp path and the referenced pids.

Phusion Passenger v6.0.6 with nginx 1.18.0 (I guess its the default for passenger 6.0.6) running on ubuntu server 20.04 up-to-date.

&amp;#x200B;

Thanks for your time and for the help :-)
## [4][[HELP] Preset data in dynamic forms.](https://www.reddit.com/r/rails/comments/imcpmz/help_preset_data_in_dynamic_forms/)
- url: https://www.reddit.com/r/rails/comments/imcpmz/help_preset_data_in_dynamic_forms/
---
I have added dynamic form feature inside the rails application.The end-user can create form fields by selecting input fields from a bunch of available input options i.e text field, radio buttons, file field etc .    It works perfectly fine. After the form is created, end-users can submit this form without any issue. Now,I have a use case in which I would like to pre-set data in some of the input fields. For instance, an admin-user created a form with the following input fields: `First Name(String), Last Name(String),Father's First Name(String), Father's Last Name(String)`, etc.

Now, a member logs in to fill in this form. I would like to preset  data in `First Name, Last Name` fields with the first\_name, last\_name of the logged in user  by fetching this information from the user's profile page. From a usability point of view, it would be a burden for the user to type this information once again. Since the form fields are created dynamically and the  name attribute of each of the input fields is a generated at run time so its hard to figure out that which input field corresponds to what user information. What could be the reasonable solution to preset data in dynamic form input fields? Any help will be highly appreciated.  Thanks.
## [5][Can't read "public/tmp/" in development &amp; production environment](https://www.reddit.com/r/rails/comments/im9wx6/cant_read_publictmp_in_development_production/)
- url: https://www.reddit.com/r/rails/comments/im9wx6/cant_read_publictmp_in_development_production/
---
It works well in localhost, but when I deployed it to development &amp; production, it returns **No such file or directory @ rb\_sysopen - /usr/src/app/public/tmp/file\_name.csv** error when I try to read the file.

The file is created with sidekiq job, and the sidekiq jobs are processed (success) so I assume the file is created since I don't have access to the server directory. This is how the file is created inside sidekiq job.

        dir_name = "#{Rails.root}/public/tmp"
        FileUtils.mkdir_p dir_name
        file = "#{dir_name}/file_name.csv"
        header = ["test"]  
    
        CSV.open( file, 'w' ) do |writer|
            writer &lt;&lt; header
            body.each do |b|
                writer &lt;&lt; [b["test"]]
            end
        end

And this is how the file is read in controller.

        file_path = "#{Rails.root}/public/tmp/#{filename}"
        File.open(file_path, 'r') do |f|
            send_data f.read,
                    :type =&gt; 'text/csv; charset=iso-8859-1; header=present',
                              :disposition =&gt; "attachment; filename=#{filename}"
        end
## [6][Multiple accounts with the same email address](https://www.reddit.com/r/rails/comments/im2mny/multiple_accounts_with_the_same_email_address/)
- url: https://www.reddit.com/r/rails/comments/im2mny/multiple_accounts_with_the_same_email_address/
---
I have devise set up for Users. I have a model for Companies, and Users have a field called company_relationship with an integer of the company.id saved to their account when they sign up. How can I check if the email and company ID are correct when a user logs in from a Company’s page.

Also if a user is signed in to another company’s page, how do I destroy the old session when they log in to the page of another company?

Really appreciate the help!
## [7][Is this the simplest way to create a Facebook-style friendship functionality?](https://www.reddit.com/r/rails/comments/im32av/is_this_the_simplest_way_to_create_a/)
- url: https://www.reddit.com/r/rails/comments/im32av/is_this_the_simplest_way_to_create_a/
---
A *symmetrical bi-directional self-referential association* is a bit of a mouthful, but I'd like to understand what the simplest way of designing something like a Facebook-style friendship would be. Using the principle of Occam's Razor, could I design this in a simpler way? To me this currently seems like a lot of work for the functionality it provides.

In the way I've made this, every time a friendship is created it sets the status to "requested", and creates a "reverse friendship" that is set to "pending" using callbacks. When the friend accepts, it updates both friendships to “accepted” and sets an “accepted at” datetime value.

When a friendship is destroyed it destroys both friendships.

The friendships and users controllers are essentially basic RESTful scaffold controllers.

User Model: 

    class User &lt; ApplicationRecord
      has_many :friendships, dependent: :destroy
      has_many :friends, -&gt; { where "status = 'accepted'" }, through: :friendships
      has_many :requested_friends, -&gt; { where "status = 'requested'" }, through: :friendships, source: :friend
      has_many :pending_friends, -&gt; { where "status = 'pending'" }, through: :friendships, source: :friend
    end
Friendship Model:
    
    class Friendship &lt; ApplicationRecord
      belongs_to :user
      belongs_to :friend, class_name: 'User', foreign_key: 'friend_id'
      scope :friends, -&gt; { where("status = 'accepted'", true) }
      scope :not_friends, -&gt; { where("status = 'accepted'", false) }
   
      attribute :status, :string, default: 'requested'
      after_create :create_inverse
        # - VALIDATIONS
      validates_presence_of :user_id, :friend_id, :status
      validate :user_is_not_equal_friend
      validates_uniqueness_of :user_id, scope: :friend_id
      after_destroy :destroy_inverse

      private
    
      def self.accept(user, friend)
        transaction do
          accepted_at = Time.now
          accept_one_side(user, friend, accepted_at)
          accept_one_side(friend, user, accepted_at)
        end
      end
    
      def self.accept_one_side(current_user, friend, accepted_at)
        request = find_by_user_id_and_friend_id(current_user, friend)
        request.status = 'accepted'
        request.accepted_at = accepted_at
        request.save!
      end
    
     
    
      def create_inverse
        self.class.create(user_id: self.friend.id, friend_id: self.user.id, status: 'pending')
      end
    
      def destroy_inverse
        self.class.delete_by(user_id: self.friend.id, friend_id: self.user.id)
      end
    
      def user_is_not_equal_friend
        errors.add(:friend, "can't be the same as the user") if user == friend
      end 
    end

Schema:

      create_table "friendships", id: :serial, force: :cascade do |t|
        t.integer "friend_id"
        t.datetime "created_at"
        t.datetime "updated_at"
        t.bigint "user_id"
        t.string "status"
        t.datetime "accepted_at"
        t.index ["user_id"], name: "index_friendships_on_user_id"
      end
    
      create_table "users", force: :cascade do |t|
        t.string "email", default: "", null: false
        t.string "encrypted_password", default: "", null: false
        t.string "reset_password_token"
        t.datetime "reset_password_sent_at"
        t.datetime "remember_created_at"
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.string "first_name"
        t.string "last_name"
        t.string "provider"
        t.string "uid"
        t.index ["email"], name: "index_users_on_email", unique: true
        t.index ["reset_password_token"], name: "index_users_on_reset_password_token", unique: true
      end

Views:

    &lt;%= link_to 'Add Friend', friendships_path(friendship: {friend_id: @user.id, user_id: current_user.id}), method: :post%&gt;
    
    &lt;% if friendship.status == "pending" %&gt; &lt;!-- If receiving friendship request ---&gt;
      &lt;%= link_to 'Accept', friendship, friendship:{status: "accepted", accepted_at:"#{Time.now}"}, remote: true, method: :patch" %&gt;&lt;%= link_to 'Delete', friendship, friendship:{user_id: friendship.friend, friend_id: current_user }, remote: true, method: :delete, data: { confirm: 'Are you sure?' }" %&gt;
       &lt;% elsif friendship.status == "accepted" &amp;&amp; friendship.friend == current_user %&gt; &lt;!-- If accepted friendship request and friend ---&gt;
       &lt;%= link_to 'Delete', friendship, friendship:{user_id: friendship.friend, friend_id: current_user }, remote: true, method: :delete, data: { confirm: 'Are you sure?' }" %&gt;
       &lt;% elsif friendship.status == "accepted" &amp;&amp; friendship.user == current_user %&gt; &lt;!-- If accepted friendship request and user ---&gt;
       &lt;%= link_to 'Delete', friendship, friendship:{user_id: current_user, friend_id: friendship.friend }, remote: true, method: :delete, data: { confirm: 'Are you sure?' }" %&gt;
       &lt;% else %&gt; &lt;!-- if sending friendship request --&gt;
       &lt;%= link_to 'Delete', friendship, friendship:{user_id: current_user, friend_id: friendship.friend}, remote: true, method: :delete, data: { confirm: 'Are you sure?' }" %&gt;
     &lt;% end %&gt;
## [8][Calculate leaves of Employee of year](https://www.reddit.com/r/rails/comments/im9hyz/calculate_leaves_of_employee_of_year/)
- url: https://www.reddit.com/r/rails/comments/im9hyz/calculate_leaves_of_employee_of_year/
---
I want To calculate leave of Employee from joining Date on yearly basis

    class Leave &lt; ApplicationRecord
     after_update :calculate_availed_leave 
     def calculate_availed_leave
        date = self.employee.joining_date.to_i + 365
        apply_date = self.leave_date.to_i
        if apply_date &lt; date
          taken_leaves = self.employee.leaves.where(status: "Approved").count
          availed_leaves = self.employee.allowed_leaves - taken_leaves
          self.employee.update(availed_leaves: availed_leaves)
        elsif apply_date &lt; new_date
          allowed_leaves = self.employee.allowed_leaves
          self.employee.update(availed_leaves: allowed_leaves)
        end
     end 
    end
## [9][Page client rendering too slow](https://www.reddit.com/r/rails/comments/im32bl/page_client_rendering_too_slow/)
- url: https://www.reddit.com/r/rails/comments/im32bl/page_client_rendering_too_slow/
---
We have a Rails Monolith in Heroku. We serve assets in the same server since we don't have many users and all the workload is managed by Sidekiq.
We use Sprockets and Webpacker for our frontend. We use react-webpacker to "prerender" the props of some root components that are later rendered in the client.

We have a listing view that's fairly slow: we render some filters and two lists. The list items are "prerendered" react-webpacker components. The filters are rendered by the old pipeline by using chosenjs.

The thing is that in Heroku, when accessing this page, all we can see for the first **40 seconds** are the filters as raw html select elements. Then, it seems that the JS executes, chosenjs and react components render.

The problem is that I have no clue why it hangs. All that I did was run a benchmark in Chrome's Performance Tab and I get that 37 seconds are spent in Layout recalculations.  So at first I thought that maybe prerendering too many react components would have something to do with this and therefore I changed that to a single prerendered component that would be in charge of building everything else. That didn't affect performance at all.

Then I took a closer look and found that the HTML page is **10mb** and that Time To First Byte is around 10 seconds. I guess this last metric has something to do with Request Queue Time, but I have no clue where to retrieve this in Heroku. I checked New Relic and it says that the worst requests to this page take 4 seconds max.

I am kind of lost what it could be, I ran into a Nate Berkopec blog post saying that Heroku buffer is of 1024 bytes and that on big requests that might affect the queue time. However, I am using Puma so I guess it shouldn't really affect it?

I am clueless on where to look. We are thinking about removing the logic from the html rendering and just add API endpoints that are fetched as soon as React components mount, but that's a huge change and I want to discard everything else before moving into that.
## [10][Is Spring meant to work when I run my full rspec tests suite?](https://www.reddit.com/r/rails/comments/iluo64/is_spring_meant_to_work_when_i_run_my_full_rspec/)
- url: https://www.reddit.com/r/rails/comments/iluo64/is_spring_meant_to_work_when_i_run_my_full_rspec/
---
I'm new to using the spring gem and have been impressed at how it has made my tests run quicker. I've followed the documentation:

1. Included the spring and spring-command-rspec gems in the development and test groups
2. ran $ bundle install
3. ran $ bundle exec spring binstub --all

I have direnv, so can type rspec and it runs it with spring. So when I do

&amp;#x200B;

$ rspec ./spec/path/to/file\_spec.rb 

It runs it much quicker! Should I be expecting the same behaviour when I run:

$ rspec ./spec/ 

Because currently, it runs at the same speed as when I run it without spring.

Any pointers and advice would be appreciated, thanks.
## [11][How to add hours to datetime?](https://www.reddit.com/r/rails/comments/ilsvo3/how_to_add_hours_to_datetime/)
- url: https://www.reddit.com/r/rails/comments/ilsvo3/how_to_add_hours_to_datetime/
---
I want to add (or subtract) hours like these  '+0330' to datetime. Let's say if I have '2020-09-03 15:30:20'

for '+0330' I want it to become '2020-09-03 19:00:20'. 

for '-0330' I want it to become '2020-09-03 12:00:20'. 

I came up with this simple method, but it seems to be prone to errors, any suggestions? :)

    def timezoneize(timedate, timezone_digits)
    digits = timezone_digits.gsub('0', '')
     if digits[0] = '+'
       timedate + digits.to_i.hours
     else
       timedate - digits.to_i.hours
     end
 end
