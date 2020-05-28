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
## [3][html.erb vs html.haml vs html.slim - which one do you use?](https://www.reddit.com/r/rails/comments/gs0x4b/htmlerb_vs_htmlhaml_vs_htmlslim_which_one_do_you/)
- url: https://www.reddit.com/r/rails/comments/gs0x4b/htmlerb_vs_htmlhaml_vs_htmlslim_which_one_do_you/
---


[View Poll](https://www.reddit.com/poll/gs0x4b)
## [4][Integrating with Google Calendar in a Rails app - The Right Way](https://www.reddit.com/r/rails/comments/gs3qov/integrating_with_google_calendar_in_a_rails_app/)
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
## [5][Monitoring multi-dyno Heroku app](https://www.reddit.com/r/rails/comments/gs3n7t/monitoring_multidyno_heroku_app/)
- url: https://www.reddit.com/r/rails/comments/gs3n7t/monitoring_multidyno_heroku_app/
---
Hello there.

I'm trying to set up some metrics monitoring on Heroku which would eventually be wired to Grafana Cloud. The formation on Heroku is multi-dyno (one web dyno running Rails and two dynos for some workers), with dynos potentially scaling up (especially web dyno).

I've gave a shot at setting up Prometheus, trying out different Ruby gems ([discourse/prometheus_exporter](https://github.com/discourse/prometheus_exporter/) and [prometheus/client_ruby](https://github.com/prometheus/client_ruby/)) but with no luck. The problem I'm facing is that dynos are isolated, so even both of these gems have multi process support, it just doesn't work on Heroku. E.g., the official Prometheus client has `DirectFileStore` which writes to the same file on the filesystem so that different processes can write to it and then expose the same metrics when `/metrics` is called. This doesn't work on Heroku because each dyno gets its own filesystem, so the file can't be shared.

Another option I've looked into potentially is using StatsD, but that will require running a deamon parallel with the Rails app or workers.

Was anyone in a similar situation before? Or do you use something totally different for your metrics for Heroku apps?

Thanks in advance!
## [6][How are you hosting small apps these days?](https://www.reddit.com/r/rails/comments/gs0hbh/how_are_you_hosting_small_apps_these_days/)
- url: https://www.reddit.com/r/rails/comments/gs0hbh/how_are_you_hosting_small_apps_these_days/
---
I have various slack bots and little project apps and I feel like I am a little bit antiquated as I host them on a digital ocean VPS and deploy using capistrano. All the cool kids these days are on about containers and I'm wondering if I'm missing out.

My way works... And it's relatively easy to understand. But I guess one big drawback is that there is a decent amount of setup and config on my host. If I had to start over, there are nginx config files to make, databases to manually install and set up, specific ruby versions to install, etc.

I've kind of given up on heroku because a $10 digital ocean VPS seems a better value than $7 per app. There are various cloud services out there, but pricing is not all that easy to figure out and they seem pretty complicated. 

I'm tempted to redo my apps using Dokku to get containers going easily on a digital ocean box, one nice thing there is that if I ever did need to burn down and start over containers make the setup a lot easier. 

I'm curious what others are doing for low cost hosting of small projects these days.
## [7][Anyone move from frontend to learn RoR?](https://www.reddit.com/r/rails/comments/gs4jsj/anyone_move_from_frontend_to_learn_ror/)
- url: https://www.reddit.com/r/rails/comments/gs4jsj/anyone_move_from_frontend_to_learn_ror/
---
Howdy all!   I'm currently working in the WordPress space on the frontend. I spent a few months about a decade ago deep-diving into Ruby and Rails before moving on to focusing more on FE and WordPress. That's paid off, as the work pays quite well and it seems like there's a decent amount of it out there. 

But I'm still curious about the "what-if" I'd stayed with RoR. I never got too-too deep into Rails back then, but I remember really loving Ruby.   I'm wondering if anyone has moved from this frontend track to RoR. 

A few questions if I may:

How long would it take to get decent with RoR — decent as in employable. Skill-wise now, I've been doing frontend for 10 years or so, I'm probably intermediate-level with React/Vue, to give you some benchmark to go on. I'm definitely not a beginner, but I'm not a backend programmer either.

There's so many posts about the job market from 1-2-3-4+ years ago, with lots of different opinions. The Covid19 situation not withstanding, anyone got an opinion on the 2020-2021 job market? 

Thanks!

ps: I recognize I'm talking about leaving one of the current hottest spaces (React / frontend) and going to a far less hot area, but still interested in perspectives.
## [8][OPEN: 2020 Ruby on Rails Developer Community Survey](https://www.reddit.com/r/rails/comments/grlcry/open_2020_ruby_on_rails_developer_community_survey/)
- url: https://www.reddit.com/r/rails/comments/grlcry/open_2020_ruby_on_rails_developer_community_survey/
---
Over ten years ago, we invited our community to participate in the first survey about the state of deploying Ruby on Rails applications. Over the years, we've evolved this to include questions about tools, frameworks, and workflows in order to see how the environment is changing.

To view previous results: [https://rails-hosting.com/](https://rails-hosting.com/)

To take the 2020 survey visit [https://planetargon.survey.fm/rails-survey-2020](https://planetargon.survey.fm/rails-survey-2020)

Thanks in advance for helping spread the word!
## [9][Would Rails be considered a good option as api backend for a social network mobile app?](https://www.reddit.com/r/rails/comments/gs1bxt/would_rails_be_considered_a_good_option_as_api/)
- url: https://www.reddit.com/r/rails/comments/gs1bxt/would_rails_be_considered_a_good_option_as_api/
---
Hi guys

The title more or less says it all. I am contemplating using graphql-ruby and rails for a project that involves a social network mobile app built in react native.

I put little stock in flavour of the month tech so I don't want just jump ship and go with something like Hasura or other BaaS solutions but I do want to be critical enough to ask if Rails is the best solution in this case?  


Any opinions are welcome!  
Thanks!
## [10][Firebase Authentication](https://www.reddit.com/r/rails/comments/gs2s0u/firebase_authentication/)
- url: https://www.reddit.com/r/rails/comments/gs2s0u/firebase_authentication/
---
Hello All,

I’m about to start a new project and I’m scoping out options for various parts. The project will contain a front end web app and a backend API which would be used by a mobile app. Normally, I’d use devise for the front end and devise token with for the backend, but I’m considering firebase auth to move the authentication part out of this application. 

The intention is to use a traditional rails front end with sessions, and the api is to be consumed solely for the mobile app. 

Has anybody got any tips on firebase authentication in this way? The firebase docs show javascript examples for authentication. I’m not keen on turning the front end into a full javascript consumer of the api since it increases the complexity of the application (but I’m not completely against this idea). The original plan was to use sprinklings of stimulus on the front end, but I’m happy to use react or vue. 

After spending a day looking at options I’m thinking that this isn’t realistic in 2020, with firebase, but it’d work with the devise solution. 

I’m just looking for thoughts and ideas. 

Thanks :)

*edit typo
## [11][Identifier of the bussines included in BLACKLIST](https://www.reddit.com/r/rails/comments/gs5map/identifier_of_the_bussines_included_in_blacklist/)
- url: https://www.reddit.com/r/rails/comments/gs5map/identifier_of_the_bussines_included_in_blacklist/
---
Good afternoon people! I hope this sub reddit is active because I have a problem and it would be great if you can help me.

It turns out that when trying to make an online purchase with a debit card, it pops up a notification on my bank app saying that the page is in BLACKLIST, exactly it says: "identifier of the business included in Blacklist".

What I need to know is the following:

\-Is it possible to alter the merchant's identifier without being part of the team of trade programmers? (like a VPN for web pages ID)

\-What solution can you think of to circumvent the system and be able to buy on this page (I swear it is not illegal! )

&amp;#x200B;

I appreciate your opinions in advance!
## [12][How to use your custom static files?](https://www.reddit.com/r/rails/comments/grxmgn/how_to_use_your_custom_static_files/)
- url: https://www.reddit.com/r/rails/comments/grxmgn/how_to_use_your_custom_static_files/
---
I am new to rails, and the version I am using is 6, I faced an issue with referring to my static files like CSS and javascript, I tried:

\-  typical ways

`&lt;link href='assets/stylesheets/customs.css' rel='stylesheet' type='text/css'&gt;`

or

`&lt;link href='javascript/stylesheets/customs.css' rel='stylesheet' type='text/css'&gt;`

or

`&lt;link href='vendor/stylesheets/customs.css' rel='stylesheet' type='text/css'&gt;`

\- import in application.css

`@import './theme/css/clean-blog.css';`

and none of them work, any suggestions, please?
