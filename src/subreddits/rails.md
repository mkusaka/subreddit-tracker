# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/gs526t/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/gs526t/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Can we parallel upload to s3 while store is file system using carrierwave?](https://www.reddit.com/r/rails/comments/gsr6wp/can_we_parallel_upload_to_s3_while_store_is_file/)
- url: https://www.reddit.com/r/rails/comments/gsr6wp/can_we_parallel_upload_to_s3_while_store_is_file/
---
Hi,

We are using `carrierwave` uploader to upload files and storing them to file system. Now we are thinking to move them to s3 from file system. But before doing that, we want a parallel upload to local file system and s3 as well. So when I have a upload like this:

    # encoding: utf-8
    
    class FileUploader &lt; CarrierWave::Uploader::Base
    
      # Include RMagick or MiniMagick support:
      # include CarrierWave::RMagick
      include CarrierWave::MiniMagick
      # include CarrierWave::Processing::MiniMagick
    
      # Choose what kind of storage to use for this uploader:
      # storage :file
      storage :file
    
    
      # Override the directory where uploaded files will be stored.
      # This is a sensible default for uploaders that are meant to be mounted:
      def store_dir
        # "uploads/#{model.class.to_s.underscore}/#{mounted_as}/#{model.id}"
        "#{model.class.to_s.underscore}/#{model.company.id}/#{mounted_as}/#{model.id}"
      end
    
      version :report, if: :is_image? do
        process :quality =&gt; 80
        process resize_to_fill: [250, 250]
      end
    
      def is_image?(image)
        # causing error, so commented it out for now.
        ## ["png", "jpg", "jpeg"].any? {|mime| image.content_type.include?(mime)}
      end
    
      # Provide a default URL as a default if there hasn't been a file uploaded:
      # def default_url
      #   # For Rails 3.1+ asset pipeline compatibility:
      #   # ActionController::Base.helpers.asset_path("fallback/" + [version_name, "default.png"].compact.join('_'))
      #
      #   "/images/fallback/" + [version_name, "default.png"].compact.join('_')
      # end
    
      # Process files as they are uploaded:
      # process :scale =&gt; [200, 300]
      #
      # def scale(width, height)
      #   # do something
      # end
    
      # Create different versions of your uploaded files:
      # version :thumb do
      #   process :resize_to_fit =&gt; [50, 50]
      # end
    
      # Add a white list of extensions which are allowed to be uploaded.
      # For images you might use something like this:
      # def extension_white_list
      #   %w(jpg jpeg gif png)
      # end
    
      # Override the filename of the uploaded files:
      # Avoid using model.id or version_name here, see uploader/store.rb for details.
      # def filename
      #   "something.jpg" if original_filename
      # end
    
    end

How do I upload to S3 in background when local filesystem store is done? I was checking the [store!](https://github.com/carrierwaveuploader/carrierwave/blob/0f733d25e2aa9c3541c4a04fb114ee526c5ec78e/lib/carrierwave/uploader/store.rb#L62) method, but I don't see any way to tell uploader that I want to upload to s3 even if the it is mentioned as `file` statically inside the code. Is there any of doing what I am looking for?
## [4][React on Rails](https://www.reddit.com/r/rails/comments/gshg3g/react_on_rails/)
- url: https://www.reddit.com/r/rails/comments/gshg3g/react_on_rails/
---
For those of you using React for the front end and Rails on the back end, what have you found are the biggest tradeoffs/gains on the development side of things? Has it become a go-to tool in your arsenal, or a one-off specific client/job thing?
## [5][Claim Your Profile Functionality](https://www.reddit.com/r/rails/comments/gsp76i/claim_your_profile_functionality/)
- url: https://www.reddit.com/r/rails/comments/gsp76i/claim_your_profile_functionality/
---
So I have a Model Profiles with people names gathered from facebook groups. I want to make a functionality where users can claim their profiles. I have implemented devise omniauth facebook registrations. So I basically want it to check if user name match  profile name and if so assign that Profile to that User. For example, when a user gets to their registration edit page they see the matching Facebook profile and a button saying 'claim profile', and when button is clicked profile gets assigned. Can you help me code it the right way and avoid spaghetti code? :)
## [6][Mistakes I will avoid in my next Ruby-on-Rails project](https://www.reddit.com/r/rails/comments/gs7ias/mistakes_i_will_avoid_in_my_next_rubyonrails/)
- url: https://www.reddit.com/r/rails/comments/gs7ias/mistakes_i_will_avoid_in_my_next_rubyonrails/
---
Full article here : [http://bdavidxyz.com/blog/mistakes-i-will-avoid-in-my-next-rails-project/](http://bdavidxyz.com/blog/mistakes-i-will-avoid-in-my-next-rails-project/) Nothing really fancy actually, but I'll be happy to discuss here if needed.
## [7][anyone have experience integrating klaviyo with rails?](https://www.reddit.com/r/rails/comments/gsi5g9/anyone_have_experience_integrating_klaviyo_with/)
- url: https://www.reddit.com/r/rails/comments/gsi5g9/anyone_have_experience_integrating_klaviyo_with/
---
  
trying to decide between actionmailer vs klaviyo
## [8][Integrating with Google Calendar in a Rails app - The Right Way](https://www.reddit.com/r/rails/comments/gs3qov/integrating_with_google_calendar_in_a_rails_app/)
- url: https://www.reddit.com/r/rails/comments/gs3qov/integrating_with_google_calendar_in_a_rails_app/
---
### Prerequisite
1) If using rails then use the gem `google-api-client`
2) I am considering here that you already have the `access_token` of the user. I will write a different blog to explain how to get that.

### 1) Do full initial synch of events
It has the following steps - 
  * Fetch a new `access_token` if the token has expired.
  * Create the service authorization object which will be used for fetching the events.

Ref code for `service authorization`
```ruby
def create_service_auth
  #create service auth
  @service = Google::Apis::CalendarV3::CalendarService.new
  @service.authorization = token.google_secret.to_authorization
  return unless token.expired?

  new_access_token = @service.authorization.refresh! #refresh access_token
end
```

  * Fetching all calendar events(past, present and future).
 
    * The full sync is the original request for all the resources of the collection you want to synchronize.
    * In the response to the list operation, you will find a field called nextSyncToken representing a sync token. You'll need to store the value of nextSyncToken. If the result set is too large and the response gets paginated, then the nextSyncToken field is present only on the very last page.
    * Depending on what your use case is, it will be better to perform this job as a background task.
    * [Events: list API](https://developers.google.com/calendar/v3/reference/events/list) is used for this. The gem provides an easier method called `list_events`

Ref code for `syncing events`
```ruby
  def get_events
    @events_arr = []
    @events_list = @service.list_events('primary', single_events: true, max_results: 500)
    @sync_token = @events_list.next_sync_token
    @page_token = @events_list.next_page_token
    @events_arr &lt;&lt; @events_list.items
    while @sync_token.blank?
      @events_list = @service.list_events('primary', single_events: true, max_results: 500, page_token: @page_token)
      @sync_token = @events_list.next_sync_token
      @page_token = @events_list.next_page_token
      @events_arr &lt;&lt; @events_list.items
    end
  end
```

### 2) Create a webhook to receive push notifications

After a full sync of events, the next step is to setup a Webhook so that google can inform us of the changes that we subscribe for.
For every user that links their calendar to the app, we will create a subscription so that we can be informed whenever there is a change in their calendar.

It has the following steps - 
  * Fetch a new `access_token` if the token has expired.
  * Create the service authorisation object which will be used for fetching the events, exactly same as shown above.
  * Set up a Channel - It creates a channel with google and specifies the callback URL or the web-hook URL.
  * Watch events - After the web-hook is set up, we need to specify what events we want to watch and also need to specify from which calendar.

```ruby
def setup_channel
  @channel = Google::Apis::CalendarV3::Channel.new(address: callback_url, id: channel_id, type: "web_hook")
end
```

`callback_url` - It can't be localhost, it has to be a valid `https` url. For testing purposes you can use [ngrok](https://ngrok.com/).
`channel_id` - This is a UUID - `SecureRandom.uuid`


```ruby
  def watch_events
    time_min = DateTime.now.rfc3339
    @webhook = @service.watch_event('primary', @channel, single_events: true, time_min: time_min)
  end
```
`primary` - refers to the `primary` calendar of the user.
`single_events` - Setting it to true also gives all events belonging to 1 single recurring event.

Now, whenever there will be any change in the primary calendar of the user google will hit the registered web-hook for the user.

In the request Google will pass `X-Goog-Resource-ID` and `X-Goog-Channel-ID`. We would have to hit the `list_events` API again to fetch the changed events data for that user.

Only difference will be that instead of passing the page token like we did earlier, we would pass the `sync_token`.

```ruby
  def get_events
    @events_list = @service.list_events('primary', single_events: true, max_results: 2500, sync_token: sync_token)
  end
```



### 3) Saving X-Goog-Resource-ID &amp; X-Goog-Channel-ID

When we created the web-hook google will return us with a `resource_id`, `resource_uri`, `id`(that we created). We need to save all this data so that we can get to know for which user the events have changed.
Also the channel expires in around 1 week so we need to keep creating new web-hooks before it expires.

### 3) Deleting the events with status `cancelled`

This is the flow that took me some time to understand. So what happens when a user changes the time of their event or has the user changed a single event or all the events in a recurring event. What google does is that 
* if the user changes a single event, then google keeps the `calendar_id` as same.
* if the user changes a recurring event and selects `all` or `following events` as option then the `calendar_id` changes for all the events. Hence, in this case we need to delete the old events and add new events in our system. So, this is a check that you will have to add when saving the calendar events in your system.


That's it - It's quite messy if you are trying to figure it out from scratch and I hope this article will help you all.
## [9][html.erb vs html.haml vs html.slim - which one do you use?](https://www.reddit.com/r/rails/comments/gs0x4b/htmlerb_vs_htmlhaml_vs_htmlslim_which_one_do_you/)
- url: https://www.reddit.com/r/rails/comments/gs0x4b/htmlerb_vs_htmlhaml_vs_htmlslim_which_one_do_you/
---


[View Poll](https://www.reddit.com/poll/gs0x4b)
## [10][Anyone move from frontend to learn RoR?](https://www.reddit.com/r/rails/comments/gs4jsj/anyone_move_from_frontend_to_learn_ror/)
- url: https://www.reddit.com/r/rails/comments/gs4jsj/anyone_move_from_frontend_to_learn_ror/
---
Howdy all!   I'm currently working in the WordPress space on the frontend. I spent a few months about a decade ago deep-diving into Ruby and Rails before moving on to focusing more on FE and WordPress. That's paid off, as the work pays quite well and it seems like there's a decent amount of it out there. 

But I'm still curious about the "what-if" I'd stayed with RoR. I never got too-too deep into Rails back then, but I remember really loving Ruby.   I'm wondering if anyone has moved from this frontend track to RoR. 

A few questions if I may:

How long would it take to get decent with RoR — decent as in employable. Skill-wise now, I've been doing frontend for 10 years or so, I'm probably intermediate-level with React/Vue, to give you some benchmark to go on. I'm definitely not a beginner, but I'm not a backend programmer either.

There's so many posts about the job market from 1-2-3-4+ years ago, with lots of different opinions. The Covid19 situation not withstanding, anyone got an opinion on the 2020-2021 job market? 

Thanks!

ps: I recognize I'm talking about leaving one of the current hottest spaces (React / frontend) and going to a far less hot area, but still interested in perspectives.
## [11][Blog with Posts and few categories to the same post](https://www.reddit.com/r/rails/comments/gsbjki/blog_with_posts_and_few_categories_to_the_same/)
- url: https://www.reddit.com/r/rails/comments/gsbjki/blog_with_posts_and_few_categories_to_the_same/
---
Hello!

I'm trying to create an blog site, very simple:

With:

Posts

Categories

The thing is The Post can have category 1,2,3 etc. 

I want to show Post Title etc, and categories.

After clicking the category i want to show all the posts with specified category.

Also I want to add a functionality with creating new category or choosing (I want to choose 1-10 or more categories to one post, THAT'S important! ) during creating new Post. 

Adding image to post is secondary thing. I've tried scaffold and simple-form gem i have found this tutorials, but i have no idea how to connect it!

[https://www.youtube.com/watch?v=\_xKglx3ox0Q&amp;list=PLefL1yCRLHTHnYWgFEdsp35aMqNQWA1MI&amp;index=4&amp;t=0s](https://www.youtube.com/watch?v=_xKglx3ox0Q&amp;list=PLefL1yCRLHTHnYWgFEdsp35aMqNQWA1MI&amp;index=4&amp;t=0s)

[https://learn.co/lessons/has-many-through-forms-rails](https://learn.co/lessons/has-many-through-forms-rails)

Please help bc it's very important for me and I'm doing it for 3 month now ...

If anyone could help me doing this I would be very grateful.

&amp;#x200B;

PEACE
## [12][Monitoring multi-dyno Heroku app](https://www.reddit.com/r/rails/comments/gs3n7t/monitoring_multidyno_heroku_app/)
- url: https://www.reddit.com/r/rails/comments/gs3n7t/monitoring_multidyno_heroku_app/
---
Hello there.

I'm trying to set up some metrics monitoring on Heroku which would eventually be wired to Grafana Cloud. The formation on Heroku is multi-dyno (one web dyno running Rails and two dynos for some workers), with dynos potentially scaling up (especially web dyno).

I've gave a shot at setting up Prometheus, trying out different Ruby gems ([discourse/prometheus_exporter](https://github.com/discourse/prometheus_exporter/) and [prometheus/client_ruby](https://github.com/prometheus/client_ruby/)) but with no luck. The problem I'm facing is that dynos are isolated, so even both of these gems have multi process support, it just doesn't work on Heroku. E.g., the official Prometheus client has `DirectFileStore` which writes to the same file on the filesystem so that different processes can write to it and then expose the same metrics when `/metrics` is called. This doesn't work on Heroku because each dyno gets its own filesystem, so the file can't be shared.

Another option I've looked into potentially is using StatsD, but that will require running a deamon parallel with the Rails app or workers.

Was anyone in a similar situation before? Or do you use something totally different for your metrics for Heroku apps?

Thanks in advance!
