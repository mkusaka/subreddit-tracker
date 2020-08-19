# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/
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
## [2][Rails on Heroku: Guide to how many dynos and which size](https://www.reddit.com/r/rails/comments/icnl4q/rails_on_heroku_guide_to_how_many_dynos_and_which/)
- url: https://www.reddit.com/r/rails/comments/icnl4q/rails_on_heroku_guide_to_how_many_dynos_and_which/
---
I just published an [exhaustive and opinionated guide](https://railsautoscale.com/how-many-dynos/) to dynos on Heroku. It answers the questions I've been hearing over and over for years:

* How many dynos should you be running?
* Which dyno type is right for your app?

I hope you find it helpful!
## [3][Rails API 401 Unauthorized](https://www.reddit.com/r/rails/comments/icmvdf/rails_api_401_unauthorized/)
- url: https://www.reddit.com/r/rails/comments/icmvdf/rails_api_401_unauthorized/
---
Hey, I have a Rails API that works wonderfully on localhost.

Authentication with jwt and knock session tokens.

Just pushed my code to Netlify and Heroku. Changed the config for it to communicate with Netlify app and only call that gets 401 Unauthorized is get users/me endpoint.

It lets you login and it sets the jwt token as expected but it can not set the user.

&amp;#x200B;

&amp;#x200B;

[The failing request](https://preview.redd.it/uqwgsamzeyh51.png?width=1014&amp;format=png&amp;auto=webp&amp;s=6d05f993b3eccd907b5001637ef4d2821fa9bc4f)

&amp;#x200B;

&amp;#x200B;

[Request headers of the failing request](https://preview.redd.it/p9seyqzokyh51.png?width=1204&amp;format=png&amp;auto=webp&amp;s=ca85f2b62fa281b5c1360ec6d27c401735fc4e1c)

&amp;#x200B;

[Login and users\/me API calls](https://preview.redd.it/zk4rdkm1fyh51.png?width=1210&amp;format=png&amp;auto=webp&amp;s=917e76f293de4f85dddd777853d01e603f6734d8)

&amp;#x200B;

[My User Controller](https://preview.redd.it/y7ytgnl5fyh51.png?width=1089&amp;format=png&amp;auto=webp&amp;s=b00e548279397a95654f293d86910f4c1004ee0b)

&amp;#x200B;

&amp;#x200B;

[Heroku logs](https://preview.redd.it/wdm4vf0cgyh51.png?width=1351&amp;format=png&amp;auto=webp&amp;s=f0836fd76df74379730947b4d02912724a9f4f30)

Full code can be found in here :

Frontend =&gt; [https://github.com/gitwithuli/white-curtain-frontend](https://github.com/gitwithuli/white-curtain-frontend)

Backend =&gt; [https://github.com/gitwithuli/white-curtain-backend](https://github.com/gitwithuli/white-curtain-backend)

&amp;#x200B;

Any help would be greatly appreciated.

Thanks.
## [4][How to include a concern inside of my controller_spec file?](https://www.reddit.com/r/rails/comments/ic9sof/how_to_include_a_concern_inside_of_my_controller/)
- url: https://www.reddit.com/r/rails/comments/ic9sof/how_to_include_a_concern_inside_of_my_controller/
---
RSpec newbie here. I have a before action method 'get\_user'  that I moved into it's own concern from the user\_controller file. Now my user\_controller\_spec file is broken. How do I include the concern, and thus my 'get\_user' method, in the spec file?
## [5][Parsing Input for Employee Data](https://www.reddit.com/r/rails/comments/ic3uuk/parsing_input_for_employee_data/)
- url: https://www.reddit.com/r/rails/comments/ic3uuk/parsing_input_for_employee_data/
---
I’m working on an app that helps with employee scheduling and want to find an easy way for users to load employee data into the application (I currently have a form that adds them one at a time).  The data would include name, start date, role, ect, and be in an unpredictable format.  I was thinking of having a textarea and finding a way for the application to parse relevant data and create resources from it.  

Is such a feature a pipe dream?  This is my first web-dev personal project.  Any advice would be appreciated.
## [6][how can I mix slim and erb in same file?](https://www.reddit.com/r/rails/comments/ick27c/how_can_i_mix_slim_and_erb_in_same_file/)
- url: https://www.reddit.com/r/rails/comments/ick27c/how_can_i_mix_slim_and_erb_in_same_file/
---
    
    = form_for @reg, url: update_tbd_path(@reg.state) do |f|
    
      .form-group
        = f.submit "Next", class: 'form-control btn btn-green tbd', id: 'form-submit', 'data-test': 'register-submit-button'
    
    
          &lt;div class="form-group hidden" id="tbd-date-picker"&gt;
            &lt;p class="question-text"&gt;What was the date?&lt;/p&gt;
              &lt;%= ft.label 'tbd_start_date', 'Date', id: 'tbd-start-date' %&gt;
            &lt;div class="input-group date"&gt;
              &lt;%= ft.text_field 'tbd_start_date', class: 'form-control tbd datepicker', placeholder: 'MM/DD/YYYY' %&gt;
              &lt;span class="input-group-addon calendar-icon"&gt;&lt;%= image_tag('icons/calendar@3x.png')%&gt;&lt;/span&gt;
            &lt;/div&gt;
          &lt;/div&gt;
    
    
    h1 Info
    
    p When do you intend to?
    
    p DatePicker
    
    p TimePicker
## [7][Help me decide on a search solution (between algolia, searchkick and pg_search)](https://www.reddit.com/r/rails/comments/ibx0pz/help_me_decide_on_a_search_solution_between/)
- url: https://www.reddit.com/r/rails/comments/ibx0pz/help_me_decide_on_a_search_solution_between/
---
Hi guys

&amp;#x200B;

I'm looking for a good search solution for my application.

The project has rails as a graphql api and a react-native app as client.

What I am specifically looking for is a way to search for users but also users in specific contexts, like search a list of friends or make sure that blocked users are not showing up. Preferrably, while also providing some other data from the db, such as the amount of posts or other data, as part of the returned results.

Currently I was looking at the following solutions:

\- **Algolia:** Easy to implement to some extent but it seems a bit finicky having to add a list of friends/blocked users etc. as part of the saved indices in Algolia itself. Plus it seems more difficult to include other data from the db when needed. Plus (and this is the biggest downside) it's not free and cost quite a lot when an application scales. 

\- **Searchkick/Elasticsearch**: 2 months ago I asked a similar yet more vague question regarding search in Rails and some people pointed out to me how easy it is to use elasticsearch with searchkick. At the same time there were people who thought it was overkill and warned me of the complexities of Elastic Search. So I'm not sure what to believe but I feel this is not necessarily more complex than setting something similar up on Algolia?

\- **pg\_search**: This is a gem I just discovered today and although I'm not sure, it might be good enough for my purposes? It uses postgres full text search but also allows search scopes/facetting, which would be the answer when searching through someone's friends or blocked users.

I would love to know what your experience is. I'm not sure if it matters too much but this a production app and not a simple learning project, so having a scalable solution is needed.   


Thank you!
## [8][whenever cron tasks do not append to existing log file, but create .gz files instead](https://www.reddit.com/r/rails/comments/ibq7i8/whenever_cron_tasks_do_not_append_to_existing_log/)
- url: https://www.reddit.com/r/rails/comments/ibq7i8/whenever_cron_tasks_do_not_append_to_existing_log/
---
I am using the "whenever" gem to schedule cron tasks. They seem to work, but they do not append to my "cron\_log.log" file but instead go on to create "cron\_log.1.gz", "cron\_log.2.gz" etc.

Does anyone here who uses the whenever gem know why this is happening and how to fix it?

I followed these instructions for logging:  [https://github.com/javan/whenever/wiki/Output-redirection-aka-logging-your-cron-jobs](https://github.com/javan/whenever/wiki/Output-redirection-aka-logging-your-cron-jobs) 

     set :output, {:error =&gt; 'path/to/app/cron_error_log.log', :standard =&gt; 'path/to/app/cron_log.log'}
## [9][Downloading in JS after using send_file](https://www.reddit.com/r/rails/comments/ibttqr/downloading_in_js_after_using_send_file/)
- url: https://www.reddit.com/r/rails/comments/ibttqr/downloading_in_js_after_using_send_file/
---
I am a noob when it comes to web dev, pls no attack

So I am working on trying to download a file via JS + ruby and send\_file specifically. JS because I am only able to tell from browser which files to download. I have a method that works and now I'm just trying to figure out how stuff works with send\_file

Previously, my controller was basically:

    def download 
        obj = Obj[params[:id].to_i]
        if obj 
            render :json =&gt; {filename: obj.file_name, contents: obj.contents}
    end 

Now, I'm trying to do it with send\_file. What I have in my controller is basically:

    def download 
        send_file filename_of_curr_obj
    end 

I have in my routes:

    ...
    collection do 
        post :download
    end
    ...

Changed post to get for when I tried to switch to send\_file.

This is what I can tell I'm supposed to do from searching online. But I don't understand how it's supposed to link up in HTML/JS. Everything I've seen online usually stops at around here, as if the steps after this are obvious(?).

So far I have a download button:

    a :id =&gt; :download_link. :download =&gt; 'filename' do
        button :id =&gt; download_button do 
            text 'Download'
        end
    end 

(My company has action.html.rb files instead of html.erb which is everywhere online. I have never seen html.rb files online in the wild thus far.)

In the JS, I was able to do something like this that worked previously:

    $('#download_button').click(function() {
        $.post('...obj/download', {id: id}, function(data) {
            ...
            $('#download_link').attr('download', data.filename);
            ...
            //created url from Blob and data.contents
            $('#download_link').attr('href', url');
            ...
        });
    });

So basically it worked like this: on browser, some actions affect ID of obj's file, so on click, send the ID via POST request to download and then download file with the controller passing the filename and file contents.

Trying to redo it now with send\_file, I am confused about two main things:

1. How to get the filename for the download attr of the a tag?
2. How to get the URL for the href attr of the a tag? I thought the URL would be of the form 

&amp;#8203;

    .../obj/download/1

or 

    .../obj/1/download

But I am getting errors when I try doing a $.post request with either form saying the URL doesn't exist. I saw online that to have a custom action have the first url form, in the routes, I should have:

    ...
    member do 
        get :download
    end 
    collection do 
        get :download
    end
    ...

But this doesn't work.
## [10][Rails 6 Bcrypt vs Passenger error on Production](https://www.reddit.com/r/rails/comments/iblyvt/rails_6_bcrypt_vs_passenger_error_on_production/)
- url: https://www.reddit.com/r/rails/comments/iblyvt/rails_6_bcrypt_vs_passenger_error_on_production/
---
After changing my ssl settings and restarting my Nginx, I started to receive the following error in the Passenger startup:

    Before process_action callback :ensure_user_signed_in has not been defined (ArgumentError) 

I am running Rails 6, Nginx/Passenger. The protected area is a single namespace only.

sessions\_controller

    class SessionsController &lt; NamespaceController
         skip_before_action :ensure_user_signed_in, only: [:new, :create]      
    
    # Present login form     
        def new 
        end      
    
    # Create Session     
        def create         
            user = User.where(email: params[:email]).first 
    
             if user &amp;&amp; user.authenticate(params[:password]) 
                 session[:user_id] = user.id             
                 redirect_to '/namespace/adminhub'
             else
                 redirect_to new_sessions_path, alert: 'Unable to authenticate'
             end
         end 
    
         # Logout
         def destroy
             reset_session
             redirect_to root_path
         end
      end 

Namespace\_controller

    class NamespaceController &lt; ApplicationController
         before_action :ensure_user_signed_in
    
          private
             def ensure_user_signed_in
                 unless current_user.present?
                 redirect_to new_sessions_path, alert: 'Must be signed in.'
             end
         end
    
          def current_user
             if session.has_key? :user_id
                 @current_user ||=User.find(session[:user_id])
             end
         end
         helper_method :current_user
       end 

I have attempted to undo my ssl changes in nginx and also to  restart passenger, neither seem to be the cause of this issue.  Interestingly, when I first pushed out the changes with the Bcrpyt, the  page loaded without issue and ran properly, as it does in my development  area. It was not until I had to restart the nginx process that this  error has come to light.

&amp;#x200B;

EDIT: I removed private from the controller, but nothing changed on the production side. What I find strange is the with private in place, this works for both my developmental side and my Rails 4 production server running passenger. 
## [11][RSpec model, request and system specs?](https://www.reddit.com/r/rails/comments/ibof0z/rspec_model_request_and_system_specs/)
- url: https://www.reddit.com/r/rails/comments/ibof0z/rspec_model_request_and_system_specs/
---
I've seen people recommend these three specs in Rails due to controller specs and feature specs being surpassed by request and system specs.

So are request specs doing the same as controller specs and test the methods of a controller individually? e.g. my request spec should test #index, #show, #new etc?

If so system specs should test the application from a users perspective right? E.g. the whole registration flow or creating a new post flow?

If this is the case for the above where do integration specs fall? Wouldn't request, integration and system specs all overlap and create duplicate tests?

I've tried to search but the books I've read or articles mainly point to controller and feature specs.
