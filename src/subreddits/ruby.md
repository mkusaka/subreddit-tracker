# ruby
## [1][Integrating with Google Calendar - The Right Way](https://www.reddit.com/r/ruby/comments/gv5hdk/integrating_with_google_calendar_the_right_way/)
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
## [2][Writing better StimulusJS controllers](https://www.reddit.com/r/ruby/comments/gullws/writing_better_stimulusjs_controllers/)
- url: https://boringrails.com/articles/better-stimulus-controllers/
---

## [3][New error after forking ruby Koans from a not "10 yr old repo"](https://www.reddit.com/r/ruby/comments/guyr6v/new_error_after_forking_ruby_koans_from_a_not_10/)
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
## [4][ActiveRecord's STI gotchas](https://www.reddit.com/r/ruby/comments/gv2nuk/activerecords_sti_gotchas/)
- url: https://blog.capsens.eu/why-you-should-avoid-nested-sti-activerecord-rails-6-b180f1bcc029?source=friends_link&amp;sk=ad41cf73d1ec992f26ccede46400c1fa
---

## [5][Somecache - A micro library to create and reuse cache configurations](https://www.reddit.com/r/ruby/comments/guvelz/somecache_a_micro_library_to_create_and_reuse/)
- url: https://www.reddit.com/r/ruby/comments/guvelz/somecache_a_micro_library_to_create_and_reuse/
---
[https://github.com/VAGAScom/somecache](https://github.com/VAGAScom/somecache)
## [6][Full page caching in Rails with Nginx and Redis](https://www.reddit.com/r/ruby/comments/guqwnm/full_page_caching_in_rails_with_nginx_and_redis/)
- url: https://vitobotta.com/2020/06/01/full-page-caching-in-rails-with-nginx-and-redis/
---

## [7][Question about gem dependencies and versions](https://www.reddit.com/r/ruby/comments/gurldl/question_about_gem_dependencies_and_versions/)
- url: https://www.reddit.com/r/ruby/comments/gurldl/question_about_gem_dependencies_and_versions/
---
Hey, everyone.  Apologies in advance if this is a dumb question that has been answered a million times.

I'm looking for an explanation of how Ruby/Bundler loads the correct version of sub-dependencies across multiple projects or gems on the same system (for example, multiple one-off projects on my local machine).  For example:

\- Project A requires gem B (v1)  

\- Project C requires gem B (v2)

I can see multiple versions of the same gem installed (I apparently decided to use RVM at some point; not sure if that's affecting things), but if different versions of gem B need to be installed or used, how is the correct version chosen?  Is a common version chosen and installed, and projects A and C both use the version?  Or, is the system able to load different versions for projects A and C at runtime?

I hunted around on Google, but was not able to find any good explanations of how it works.  If anyone could help me out, I would greatly appreciate it!
## [8][Is there an equivalent framework to Tornado (Python) in Ruby?](https://www.reddit.com/r/ruby/comments/guhanp/is_there_an_equivalent_framework_to_tornado/)
- url: https://www.reddit.com/r/ruby/comments/guhanp/is_there_an_equivalent_framework_to_tornado/
---
I have been spending a bit of time in Python since discovering the Tornado framework, but I still can't get to love Python the way I love Ruby.

I am looking for a Ruby framework that supports non-blocking HTTP requests and also WebSockets.

From looking around, it seems like you have to cobble something together from different libraries. Or it is possible with an existing framework but that they don't mention it explicitly because it's a given that it would support it.
## [9][Is ruby koans compatible with ruby 2.6.1?](https://www.reddit.com/r/ruby/comments/gup0ot/is_ruby_koans_compatible_with_ruby_261/)
- url: https://www.reddit.com/r/ruby/comments/gup0ot/is_ruby_koans_compatible_with_ruby_261/
---
 

Just starting with Ruby Koans as a way to practice aside from traditional homework.

Following the outlined instructions on the Koans site it says to run

    ruby path_to_enlightenment.rb

which should return

    [ ruby_koans ] $ ruby  path_to_enlightenment.rb (in /Users/person/dev/ruby_koans) cd koans  Thinking AboutAsserts test_assert_truth has damaged your karma. You have not yet reached enlightenment ... &lt;false&gt; is not true. Please meditate on the following code: ./about_asserts.rb:10:in `test_assert_truth' path_to_enlightenment.rb:27  mountains are merely mountains

but instead is returning

    ruby path_to_enlightenment.rb  Thinking AboutAsserts test_assert_truth has damaged your karma. You have not yet reached enlightenment ... &lt;false&gt; is not true. Please meditate on the following code: Traceback (most recent call last): 2: from  /mnt/c/Users/skywalker/ruby_koans/koans/edgecase.rb:265:in `block in &lt;top (required)&gt;' 1: from /mnt/c/Users/skywalker/ruby_koans/koans/edgecase.rb:112:in `report' /mnt/c/Users/skywalker/ruby_koans/koans/edgecase.rb:102:in `assert_failed?': uninitialized constant  EdgeCase::Sensei::AssertionError (NameError)

Would anyone know why im getting this unitialized constant error ?

    uninitialized constant  EdgeCase::Sensei::AssertionError (NameError)

I highly doubt the very first test would ask be to fix an unitialized constant error and no where in the read does it say to do so. Would you know why im getting this error and how i could fix it ?

The only thing i can think of is maybe Koans is not compatible with Ruby 2.6 ?
## [10][gRPC concurrency](https://www.reddit.com/r/ruby/comments/guffav/grpc_concurrency/)
- url: https://www.reddit.com/r/ruby/comments/guffav/grpc_concurrency/
---
So, I have a lightweight ruby gRPC server running with docker and kubernetes. No threading or forking. I’ve used rails / rack apps for a long time, and I’m used to having some sense of concurrency via webrick, unicorn, puma, passenger. 

My question is around concurrency. Since this service has such a small footprint, of like 10m cpu and 10mb of ram, would it be best to scale up the pods and let the cluster handle the load balancing? My searches for “ruby grpc concurrency” were not fruitful, so there doesn’t seem to be anything out of the box for going the “traditional” way.
