# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/
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
## [2][Solopreneur owner of 20k/month Rails app wants to give inside view](https://www.reddit.com/r/rails/comments/ev2s0n/solopreneur_owner_of_20kmonth_rails_app_wants_to/)
- url: https://www.reddit.com/r/rails/comments/ev2s0n/solopreneur_owner_of_20kmonth_rails_app_wants_to/
---
Ten years ago I built a niche marketplace and have been living from the revenue ever since. It was designed to be a lifestyle business from the get-go so I spent a ton of time initially optimizing and automating everything.

I want to create a stream that gives a behind-the-scenes, nothing-hidden, look at what I've built and how I work, but I've no idea which topics are most interesting. If I learned one thing in creating my business, it's that it's all too easy to create something that no-one wants.

So what topics would you like to be covered?

Some ideas I had were

**Coding**

1. generating a maximally enjoyable development environment (e.g. seeding data, simulating cron, mirroring production as much as possible etc.)
2. removing brittleness from integration tests that run on circleci
3. dealing with the shitshow that is sales tax accounting across multiple currencies
4. detecting and recovering from production bugs asap
5. dealing with the real-world mess that is imperfect user input (e.g. when they type emails with a leading space or inconsistent capitalization; when they create a tag that is almost the same as a previous one — like E Guitar vs. Electric Guitar—and now your data is split across two areas)
6. discussing the 8+ year consequences of certain architectural/software design issues
7. streamlining massive amounts of config
8. multi-redundant systems of backup to prevent disaster
9. designing error messages and a logging strategy that speeds up recovery from errors
10. a tour of the most evil, insidious bugs I dealt with over the years (I keep a diary for them)
11. payment systems in-depth (refunds, errors etc.)
12. caching systems for performance

**Business/marketing**

1. how I use data to decide to add/remove a feature
2. AB testing a web app
3. technical SEO (microdata, site structure for internal links, google's tools, sitemaps, etc.) — I get 85% of my traffic (and therefore revenue) from SEO, so I know a thing or two
4. how I use JS and integration tests on all tracking code (critical to get right in my business)
5. auto-email systems to previous customers for extra sales
6. Adwords workflow to drive revenue
7. Analytics workflow to figure out what content working
8. Writing copy that gets sales (what worked for me vs. didn't)

\_If there's something you're interested in but don't see mentioned above, please do share.\_
## [3][How to build a Real Estate Platform in Rails 6 [FULL WALKTHROUGH]](https://www.reddit.com/r/rails/comments/ev1zvv/how_to_build_a_real_estate_platform_in_rails_6/)
- url: https://www.reddit.com/r/rails/comments/ev1zvv/how_to_build_a_real_estate_platform_in_rails_6/
---
Hi guys, some of you here may have seen my previous posts over the past few weeks where I created several videos on building a Tinder clone and more recently a Reddit clone in Ruby on Rails 6. One of my most popular Rails builds to date is the Real Estate Platform which I haven't shared here before. If you enjoy this style of video where you can build along and create a working product then this video series will be worth your time. 

Sharing here for anyone interested. Hopefully this is helpful to some of you guys.

**Part 1**: Set up user accounts using the devise gem, add functionality for users to create a property listing (image upload using carrierwave gem to Amazon S3), then on the front end we build a latest properties widget and property profile view.

[https://www.youtube.com/watch?v=qvGGz-2WHpU](https://www.youtube.com/watch?v=qvGGz-2WHpU)

**Part 2**: Create an attractive property profile showing a see more/see less property details section, agent profile picture and company, related properties from that agent, responsive design and more.

[https://www.youtube.com/watch?v=SCXiNYoDibw](https://www.youtube.com/watch?v=SCXiNYoDibw)

**Part 3:** On property profile we'll add a click to view telephone number, a modal contact form to contact the agent and a google map that is generated from the properties address

[https://www.youtube.com/watch?v=3UrvdFlWakw](https://www.youtube.com/watch?v=3UrvdFlWakw)

**Part 4**: Add an agent profile page to show a cover image, user details and stats about properties leased and sold. Users can now change the status of properties to mark them as leased or sold from their user account when signed in. We will also add blog posts to the app and improve the homepage styling. In the backend we add an accounts page for admin accounts and ability to add/edit/update blog entries.

[https://www.youtube.com/watch?v=j9fbJ7\_3N1Y](https://www.youtube.com/watch?v=j9fbJ7_3N1Y)

I'll be posting new web app builds on YouTube using Ruby on Rails 6 so feel free to check out the channel and add some recommendations for upcoming builds. I'm already collecting a list of user suggestions so feel free to leave me some info on the video comments.

Cheers! :)
## [4][Rails 6 w Bootstrap Template (Customizable)](https://www.reddit.com/r/rails/comments/euywh2/rails_6_w_bootstrap_template_customizable/)
- url: https://www.reddit.com/r/rails/comments/euywh2/rails_6_w_bootstrap_template_customizable/
---
For everyone that want's to save some boiler-plating I've implemented  bootstrap on a rails 6 app and created a template where you can  customize bootstrap features, add/remove them and of course add your  customs styles. Feel free to fork or clone and use it whatever you like.   
[https://github.com/MarioDena/Rails6-Bootstrap-Template](https://slack-redir.net/link?url=https%3A%2F%2Fgithub.com%2FMarioDena%2FRails6-Bootstrap-Template&amp;v=3)
## [5][Polymorphic associations to model has_many with different types](https://www.reddit.com/r/rails/comments/ev575w/polymorphic_associations_to_model_has_many_with/)
- url: https://www.reddit.com/r/rails/comments/ev575w/polymorphic_associations_to_model_has_many_with/
---
I'm learning Rails and currently working on my first application. 

I have a Chapter that has\_many Activities. There are different types of activities and they all have very different attributes. I have a base Activity and different subclasses to represent each activity type. 

I'm having issues modeling this scenario without using STI (since I don't want a single table with many null columns). I read about polymorphic associations but I'm not sure how to use this feature to model this situation (I'm not even sure if this scenario is a good fit for polymorphic associations). 

Has anyone modeled something like this using this feature?

&amp;#x200B;

Thanks!
## [6][What/How/When do you set tokens for records?](https://www.reddit.com/r/rails/comments/ev0v4m/whathowwhen_do_you_set_tokens_for_records/)
- url: https://www.reddit.com/r/rails/comments/ev0v4m/whathowwhen_do_you_set_tokens_for_records/
---
I've always been told not to use ID's to get records.

When, and how do you set your tokens and in your opinion what is the best practice?

&amp;#x200B;

I generate tokens like this:

&amp;#x200B;

    validates :token, presence: true
    validates :token, uniqueness: true
    
    def generate_token
        begin
            self.token = "#{SecureTandom.random_number(9e5) + 1e5).to_i}"
        end while self.class.find_by(token: token)
    end

Is this a good way, is there a better way?

&amp;#x200B;

What's your opinions on this? 

Should EVERY model entry have a token, only forward facing models (ones that would be in hidden\_fields, forms, passing through console, etc.)?
## [7][How to generate a PDF file with information entered into a form using Ruby](https://www.reddit.com/r/rails/comments/euzx0s/how_to_generate_a_pdf_file_with_information/)
- url: https://www.reddit.com/r/rails/comments/euzx0s/how_to_generate_a_pdf_file_with_information/
---
Hello guys,

I would like to use Ruby on rails to build an application that allows me to enter details into a form , and have such form spit out a PDF file containing the information of each field.

Can someone point me in the right direction? I would gladly appreciate it.
## [8][AJAX calls for continuous scrolling - how do I pass the page parameter to render the correct portion of the list?](https://www.reddit.com/r/rails/comments/euydfc/ajax_calls_for_continuous_scrolling_how_do_i_pass/)
- url: https://www.reddit.com/r/rails/comments/euydfc/ajax_calls_for_continuous_scrolling_how_do_i_pass/
---
Ok so in rails 6.0 I am trying to use AJAX requests for continuous scrolling on my Subs show view, loading posts on the sub as I scroll down. I used the Kaminari gem for pagination, and I am trying to append the next page of results with the url in the "Next" link. It is working... sort of. It doesn't append the correct page but instead appends the first page repeatedly.

I am a beginner to AJAX so please, baby me. I imagine maybe the solution has to do with the controller but I don't know. I've gone through a LOT of troubleshooting just to get to this point.

&amp;#x200B;

Here are the relevant segments of my code:

subs\_controller.erb:

      def show
        @posts = @sub.posts.order('created_at DESC').page(params[:page]).per(8)
        respond_to do |format|
          format.html
          format.js
        end
      end

subs/show.html.erb:

    &lt;div id="subsViewPosts" class="posts"&gt;
      &lt;%# display in descending order to show newest first %&gt;
      &lt;span id="renderPosts"&gt;
        &lt;%= render @posts %&gt;
      &lt;/span&gt;
    
      &lt;div class="paginate pb-4"&gt;
        &lt;%= paginate @posts, remote: true %&gt;
      &lt;/div&gt;
    &lt;/div&gt;
    
    
    &lt;script type="text/javascript"&gt;
      $(document).on('turbolinks:load', function () {
        $(window).on('scroll', function () {
          let url = $('.pagination .next a').attr('href');
          if ($(window).scrollTop() &gt; $(document).height() - $(window).height() - 50) {
            // Near bottom! load stuff:
            $('.pagination').text("Fetching...");
            $.getScript(url);
          }
        })
      });
    &lt;/script&gt;

show.js.erb:

    $('#subsViewPosts #renderPosts').append("&lt;%= j render @posts %&gt;");
    $('.pagination').replaceWith('&lt;%= j paginate(@posts) %&gt;');

and the partial posts/\_post.html.erb is a nice lengthy template for displaying pretty posts with options and whatnot, and it renders just fine.

&amp;#x200B;

The result of all of this is that it does dynamically load posts as I scroll near the bottom, but it only loads the first page each time, instead of loading the next page.

I simply need to figure out how to pass the correct page of posts (the number of which is contained in the url variable in the code I wrote) to the render method. Whether that should be via the controller or via the JavaScript itself, idk. Any help would be greatly appreciated. Thanks for reading my spiel.
## [9][How to create a raffle/spinner site using rails?](https://www.reddit.com/r/rails/comments/euy8mv/how_to_create_a_rafflespinner_site_using_rails/)
- url: https://www.reddit.com/r/rails/comments/euy8mv/how_to_create_a_rafflespinner_site_using_rails/
---
Hi, I am trying to create a site where a user can input a code which is unique then allows them to play a spinner type game(each reward have different probability).. where they spin a wheel, and the rewards they win get saved to the backend. I cannot find a source/example on where to start on creating a spinner wheel... if anyone knows a source to start from.. it will be helpful.. Thanks!
## [10][How can i build a never ending job in rails for web scraping?](https://www.reddit.com/r/rails/comments/eupbpd/how_can_i_build_a_never_ending_job_in_rails_for/)
- url: https://www.reddit.com/r/rails/comments/eupbpd/how_can_i_build_a_never_ending_job_in_rails_for/
---
Goal: I want to make a web scraper in a Rails app that runs indefinitely and can be scaled.

Current stack app is running on: ROR/Heroku/Redis/Postgres

Idea: I was thinking of running a Sidekiq Job that runs every n minutes and checks if there are any proxies available to scrape with (these will be stored in a table with status sleeping/scraping).

Assuming there is a proxy available to scrape it will then check (using Sidekiq API) if there is any available workers to start up another job to scrape with the available proxy.

This means i could scale the scraper by increasing number of workers and the number of available proxies. If for any reason the Job fails the Job that looks for available proxies will just start it again.

Questions: Is this the best solution for my goal? Is utilizing long running Sidekiq jobs the best idea or could this blow up?
## [11][Using knockout JS for a Rails app, how is the “view model” different from the model, if at all?](https://www.reddit.com/r/rails/comments/eupxm5/using_knockout_js_for_a_rails_app_how_is_the_view/)
- url: https://www.reddit.com/r/rails/comments/eupxm5/using_knockout_js_for_a_rails_app_how_is_the_view/
---
https://knockoutjs.com/documentation/observables.html

I understand with the “view model” nothing is it persisted to the database.  

Isn’t that the same when we do like a create controller action and say ‘Order.new’, or is that different than Knockout JS cause that’s direct interaction with the rails Order model?

I guess from the page above, with a “view model” AJAX is used to call some server-side code to write to the stored data model.  Perhaps that’s the notes difference from a regular model?
