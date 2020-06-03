# ruby
## [1][Is there a GUI Toolkit for Ruby that is currently an active and simple project like "Shoes"? I find Shoes very intuitive, too bad it is no longer developed. Do you know anything like this? I have searched around but I find all the other more cumbersome toolkits ....](https://www.reddit.com/r/ruby/comments/gvq1ww/is_there_a_gui_toolkit_for_ruby_that_is_currently/)
- url: https://www.reddit.com/r/ruby/comments/gvq1ww/is_there_a_gui_toolkit_for_ruby_that_is_currently/
---

## [2][Ansible Rails: Deploy Ruby on Rails apps easily](https://www.reddit.com/r/ruby/comments/gvr7m9/ansible_rails_deploy_ruby_on_rails_apps_easily/)
- url: https://www.reddit.com/r/ruby/comments/gvr7m9/ansible_rails_deploy_ruby_on_rails_apps_easily/
---
I 've been using Heroku to deploy my Rails apps but I always wanted to learn how it all works under the hood. Over the last couple of months, I decided to learn more about how to set up a server and deploy a Rails app to production. I've made this open-source project in order to consolidate my learning.

Ansible Rails is an Ansible playbook for easily deploying Ruby on Rails applications.

[https://github.com/EmailThis/ansible-rails](https://github.com/EmailThis/ansible-rails)

It includes roles for performing the following tasks -

* Installation of common packages
* Ruby (via rbenv)
* Rails 6, Puma, Sidekiq
* Redis
* Nodejs/Webpack/yarn
* Postgresql and another role that saves backups to S3
* Install nginx, Certbot (for Letsencrypt SSL Certs)
* Deploying using Ansistrano

Let me know what you guys think about it 🙌🏻
## [3][Services and tools we use for development](https://www.reddit.com/r/ruby/comments/gvpwp9/services_and_tools_we_use_for_development/)
- url: https://www.reddit.com/r/ruby/comments/gvpwp9/services_and_tools_we_use_for_development/
---
Tools and services that help to automatize development flow. Easy to install with increasing team effectiveness.

[https://jtway.co/services-and-tools-we-use-for-development-2749af5ad08a](https://jtway.co/services-and-tools-we-use-for-development-2749af5ad08a)
## [4][Understanding Programs Using Graphs](https://www.reddit.com/r/ruby/comments/gvbn0p/understanding_programs_using_graphs/)
- url: https://engineering.shopify.com/blogs/engineering/understanding-programs-using-graphs
---

## [5][Designing helpful service objects. Part 1. Choosing the right design](https://www.reddit.com/r/ruby/comments/gvbrg2/designing_helpful_service_objects_part_1_choosing/)
- url: https://www.morozov.is/2020/06/01/helpful-service-objects-part-1-chosing-right-design.html
---

## [6][Integrating with Google Calendar - The Right Way](https://www.reddit.com/r/ruby/comments/gv5hdk/integrating_with_google_calendar_the_right_way/)
- url: https://www.reddit.com/r/ruby/comments/gv5hdk/integrating_with_google_calendar_the_right_way/
---
## Prerequisite

1. If using rails then use the gem `google-api-client`
2. I am considering here that you already have the `access_token` of the user. I will write a different blog to explain how to get that.

## 1) Do full initial synch of events

It has the following steps -

* Fetch a new `access_token` if the token has expired.
* Create the service authorization object which will be used for fetching the events.

Ref code for `service authorization`

    def create_service_auth
      #create service auth
      @service = Google::Apis::CalendarV3::CalendarService.new
      @service.authorization = token.google_secret.to_authorization
      return unless token.expired?
    
      new_access_token = @service.authorization.refresh! #refresh access_token
    end

* Fetching all calendar events(past, present and future).
   * The full sync is the original request for all the resources of the collection you want to synchronize.
   * In the response to the list operation, you will find a field called nextSyncToken representing a sync token. You'll need to store the value of nextSyncToken. If the result set is too large and the response gets paginated, then the nextSyncToken field is present only on the very last page.
   * Depending on what your use case is, it will be better to perform this job as a background task.
   * [Events: list API](https://developers.google.com/calendar/v3/reference/events/list) is used for this. The gem provides an easier method called `list_events`

Ref code for `syncing events`

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

## 2) Create a webhook to receive push notifications

After a full sync of events, the next step is to setup a Webhook so that google can inform us of the changes that we subscribe for. For every user that links their calendar to the app, we will create a subscription so that we can be informed whenever there is a change in their calendar.

It has the following steps -

* Fetch a new `access_token` if the token has expired.
* Create the service authorisation object which will be used for fetching the events, exactly same as shown above.
* Set up a Channel - It creates a channel with google and specifies the callback URL or the web-hook URL.
* Watch events - After the web-hook is set up, we need to specify what events we want to watch and also need to specify from which calendar.

&amp;#8203;

    def setup_channel
      @channel = Google::Apis::CalendarV3::Channel.new(address: callback_url, id: channel_id, type: "web_hook")
    end

`callback_url` \- It can't be localhost, it has to be a valid `https` url. For testing purposes you can use [ngrok](https://ngrok.com/). `channel_id` \- This is a UUID - `SecureRandom.uuid`

      def watch_events
        time_min = DateTime.now.rfc3339
        @webhook = @service.watch_event('primary', @channel, single_events: true, time_min: time_min)
      end

`primary` \- refers to the `primary` calendar of the user. `single_events` \- Setting it to true also gives all events belonging to 1 single recurring event.

Now, whenever there will be any change in the primary calendar of the user google will hit the registered web-hook for the user.

In the request Google will pass `X-Goog-Resource-ID` and `X-Goog-Channel-ID`. We would have to hit the `list_events` API again to fetch the changed events data for that user.

Only difference will be that instead of passing the page token like we did earlier, we would pass the `sync_token`.

      def get_events
        @events_list = @service.list_events('primary', single_events: true, max_results: 2500, sync_token: sync_token)
      end

## 3) Saving X-Goog-Resource-ID &amp; X-Goog-Channel-ID

When we created the web-hook google will return us with a `resource_id`, `resource_uri`, `id`(that we created). We need to save all this data so that we can get to know for which user the events have changed. Also the channel expires in around 1 week so we need to keep creating new web-hooks before it expires.

## 3) Deleting the events with status cancelled

This is the flow that took me some time to understand. So what happens when a user changes the time of their event or has the user changed a single event or all the events in a recurring event. What google does is that

* if the user changes a single event, then google keeps the `calendar_id` as same.
* if the user changes a recurring event and selects `all` or `following events` as option then the `calendar_id` changes for all the events. Hence, in this case we need to delete the old events and add new events in our system. So, this is a check that you will have to add when saving the calendar events in your system.

That's it - It's quite messy if you are trying to figure it out from scratch and I hope this article will help you all.
## [7][Discussing Rails Deployment and Hosting Options with Nate Berkopec - The Rails with Jason Podcast](https://www.reddit.com/r/ruby/comments/gv90wc/discussing_rails_deployment_and_hosting_options/)
- url: https://www.codewithjason.com/rails-with-jason-podcast/nate-berkopec-2/
---

## [8][Writing better StimulusJS controllers](https://www.reddit.com/r/ruby/comments/gullws/writing_better_stimulusjs_controllers/)
- url: https://boringrails.com/articles/better-stimulus-controllers/
---

## [9][New error after forking ruby Koans from a not "10 yr old repo"](https://www.reddit.com/r/ruby/comments/guyr6v/new_error_after_forking_ruby_koans_from_a_not_10/)
- url: https://www.reddit.com/r/ruby/comments/guyr6v/new_error_after_forking_ruby_koans_from_a_not_10/
---
After forking from git hub, running rake or  ruby path\_to\_enlightenment.rb  in  koans returns 

    Thinking AboutAsserts
      test_assert_truth has damaged your karma.
    
    You have not yet reached enlightenment ...
    &lt;false&gt; is not true.
    
    Please meditate on the following code:
    Traceback (most recent call last):
            2: from /mnt/c/Users/skywalker/dev/flatiron/projects/ruby_koans/koans/edgecase.rb:265:in `block in &lt;top (required)&gt;'
            1: from /mnt/c/Users/skywalker/dev/flatiron/projects/ruby_koans/koans/edgecase.rb:112:in `report'
    /mnt/c/Users/skywalker/dev/flatiron/projects/ruby_koans/koans/edgecase.rb:102:in `assert_failed?': uninitialized constant EdgeCase::Sensei::AssertionError (NameError)

what is this 

    uninitialized constant EdgeCase::Sensei::AssertionError (NameError)
## [10][Somecache - A micro library to create and reuse cache configurations](https://www.reddit.com/r/ruby/comments/guvelz/somecache_a_micro_library_to_create_and_reuse/)
- url: https://www.reddit.com/r/ruby/comments/guvelz/somecache_a_micro_library_to_create_and_reuse/
---
[https://github.com/VAGAScom/somecache](https://github.com/VAGAScom/somecache)
