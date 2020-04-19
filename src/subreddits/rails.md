# rails
## [1][GitHub - styd/pry-diff-routes: Inspect routes changes in Rails console.](https://www.reddit.com/r/rails/comments/g45csl/github_stydprydiffroutes_inspect_routes_changes/)
- url: https://www.reddit.com/r/rails/comments/g45csl/github_stydprydiffroutes_inspect_routes_changes/
---
Hi fellow Rails developer,

First of all, let's keep being safe and healthy and enjoy staying at home. Secondly, I want to say thank you to those of you who have contributed to my older gem, [ApexCharts.RB](https://github.com/styd/apexcharts.rb), by using, reporting issues, sending pull requests, or just giving the repo a star (I couldn't believe it at first but even [Chartkick](https://chartkick.com) author Andrew Kane and [ApexCharts](https://apexcharts.com) author Juned Chhipa also gave a star. Thank you!). I sincerely appreciate all of your efforts.

Today, I want to tell you about a new gem that I've been working on since last month called [PryDiffRoutes](https://github.com/styd/pry-diff-routes). It's a Pry plugin that works in Rails console to see the route changes that we make.

The way it works is that it would save the current routes as Rails console is starting and then from that point on the routes you change -- don't forget to save the route file -- is comparable with the routes when you started the Rails console. When you type the Pry command `diff-routes`, it will output your changes as removed, modified, or new routes.

What's "modified"? And why don't I just display the **before** and the **after**?
IMHO, when we talk about an endpoint, what we really meant usually is a specific combination of a _verb_ and a _path_. When they don't change, we consider it as the same endpoint. By changing either the _verb_ or the _path_, we can say that we are removing the old route and creating a new route. When we just change where it routes to, we are modifying the route. So, what I'm doing is just group routes changes based on our intention for those changes.

I hope it is useful for you and let me know what you think about it.
Happy weekend!
## [2][Rails Associations best practices?](https://www.reddit.com/r/rails/comments/g46rgf/rails_associations_best_practices/)
- url: https://www.reddit.com/r/rails/comments/g46rgf/rails_associations_best_practices/
---
So this is something i've been curious about for awhile, and that is the best way to structure associations. Specifically when it comes to chaining has\_many down a tree or using a more spread out hierarchy.

For example:

Lets say I have a User that can leave comments. Comments can also have likes. I could say that a User has\_many comments and that a Comment has\_many likes and that a User has\_many Likes through comments.

OR

I Could say a User has\_many Likes and has\_many Comments (and I suppose also could say a Comment has\_many Likes as well).

Is one way strictly better than the other? And when we can avoid join tables should we?
## [3][Clients want to self-host my Rails app - high level advice](https://www.reddit.com/r/rails/comments/g3qaw3/clients_want_to_selfhost_my_rails_app_high_level/)
- url: https://www.reddit.com/r/rails/comments/g3qaw3/clients_want_to_selfhost_my_rails_app_high_level/
---
I launched a moderately successful SaaS app a couple of years ago which pays the bills. It's Rails 6, PostgreSQL, Redis &amp; Elasticsearch primarily, with a Grape API layer using JWTs. It uses Auth0/Devise for authentication. Lately I've had more and more clients asking to self-host my service. I've always steered away from this in the past. I'm looking for some advice so I don't come across as completely uninformed in conversations:

* Does self-hosting always imply that I have to share my source code with the client?
* Would I normally be expected to physically show up on site to get things set-up?
* If I get my Rails app into a Docker container, is that the best way to go? Assume it still shares my source code
* I'm reliant on 3rd party services - Auth0, Segment, Intercom to name a few. Do I straight up have to find alternatives to these services? I am guessing so since the driver behind self-hosting is that they would 'own' the data
* How would patching/updates work for a self-hosted solution? Right now I'm hosting on Heroku and my git repo is on Github - would the client fork the repo and manage updates themselves?
* Any other considerations?

Appreciate your feedback. Sorry if this is the wrong place to ask!

Thanks
## [4][I've just published a new post: You don't have to be a genius to contribute to the open source](https://www.reddit.com/r/rails/comments/g3nm6b/ive_just_published_a_new_post_you_dont_have_to_be/)
- url: https://www.reddit.com/r/rails/comments/g3nm6b/ive_just_published_a_new_post_you_dont_have_to_be/
---
I wrote a new post about a pull request I've made to the Rails source code yesterday ––&gt;

[https://www.thatweeklytech.com/posts/5-you-don-t-have-to-be-a-genius-to-contribute-to-the-open-source](https://www.thatweeklytech.com/posts/5-you-don-t-have-to-be-a-genius-to-contribute-to-the-open-source)

Thanks for reading!
## [5][Devise Token Auth + Omniauth](https://www.reddit.com/r/rails/comments/g3s0v7/devise_token_auth_omniauth/)
- url: https://www.reddit.com/r/rails/comments/g3s0v7/devise_token_auth_omniauth/
---
I’ve been trying to wrap my head around setting up social login in my Rails API application with a react native front end but I can’t can’t seem to figure out a solution.

Right now, I’m using devise token Auth as my authentication library but feel like removing it completely due to frustration. In particular, social login.

There seems to be NO proper guides available explaining how the social login flow works, whether it is the code grant flow in the browser or the implicit grant with native SDKs and I’ve searched all through the web trying to find any article that could help me with this headache.

Given the context, my question is how can you setup social login, with or without omniauth, using devise token Auth (or without it).

I’ve thought of setting up a custom Service model and OmniauthCallbacks controller but I didn’t know how that was supposed to integrate on the mobile end, am I supposed to redirect User -&gt; Rails -&gt; Login a provider -&gt; Back to Rails -&gt; Back to the app or do I send the access token generated on the client to the web?

Any help is very much appreciated
## [6][Best strategy for storing translations for frontend JS consumption](https://www.reddit.com/r/rails/comments/g3qwy6/best_strategy_for_storing_translations_for/)
- url: https://www.reddit.com/r/rails/comments/g3qwy6/best_strategy_for_storing_translations_for/
---
This is a mostly frontend question that needs backend insight, which is why I'm posting it here.

I'm a backend dev working on a Rails project with a heavy Vue frontend. The project uses the Rails i18n library for displaying text to users. The project currently uses this to inject translations into a rendered Vue component:

```
&lt;% # Some .erb file %&gt;
&lt;some-vue-component t="&lt;%= t('js.some_vue_component').to_json %&gt;"&gt;&lt;/some-vue-component&gt;
```

We can then use `polyglot` or whatever library to read the right translation.

This is OK for simple components. If there are a bunch of sub-components that need translations, you have to either (1) Inject a subset of the translations to those components, making them tedious to use; (2) store the translations in the global Vue component and have the components adhere a particular JS structure hierarchy, making it somewhat brittle; or (3) declare a global JS variable in a view elsewhere, as recommended by the [polyglot library](https://github.com/airbnb/polyglot.js#translation), also making it particularly brittle and dependent on understanding a js hierarchy.

However, the biggest drawback for the approaches above (IMO) is how tedious they make Javascript testing. Every test file then requires the defining of arbitrary translations to be injected into the component. And the translations you defined are not available in JS unit tests. (Testing in JS is already awful, and this makes it doubly so.)

Another approach that would solve this testing problem would be do away with injections, and store the translations as `.js` files, allowing each component to `import` the relevant translation. For example:

```
// translations/some_vue_component.js
export default {
  en: {
    hello: {
      world: "Hello, world!"
    }
  }
}

// some_vue_component.vue that knows it is in "en" locale
&lt;template&gt;
  &lt;div&gt;{{t('hello.world')}}&lt;/div&gt;
&lt;/template&gt;
&lt;script&gt;
  import translations from "../translations/some_vue_component";
  // ...
&lt;/script&gt;
```

This approach takes care of all JS testing issues, since translations are handled internally with the component. The big downside I see here is that it could (would) bloat your frontend. If one is accommodating a lot of languages, you could be serving clients with an awful lot of unnecessary data.

What do you guys do? Is there something I am missing?
## [7][ActiveStorage Direct Uploads throwing 422](https://www.reddit.com/r/rails/comments/g3p6je/activestorage_direct_uploads_throwing_422/)
- url: https://www.reddit.com/r/rails/comments/g3p6je/activestorage_direct_uploads_throwing_422/
---
I've been banging my head against this wall trying to figure out what's going on and would appreciate the help of someone with more experience.

Here's what's going on,

I'm trying to use the \`ActiveStorage::DirectUploadsController\` to upload an image. Here's what my custom controller looks like:

`class DirectUploadsController &lt; ActiveStorage::DirectUploadsController`

`protect_from_forgery with: :null_session`

`before_action :authenticate_request`

&amp;#x200B;

`private`

&amp;#x200B;

`def authenticate_request`

`user =` [`AuthorizeApiRequest.call`](https://AuthorizeApiRequest.call)`(request.headers).result`

`render json: { error: 'Not Authorized' }, status: 401 unless user`

`end`

`end`

&amp;#x200B;

However, when hitting the endpoint with a jpeg, I get the following:

`Started POST "/direct_uploads" for ::1 at 2020-04-17 17:42:40 -0400`

`Processing by DirectUploadsController#create as JSON`

`Parameters: {"blob"=&gt;{"filename"=&gt;"image_picker_1249334B-5119-4F5E-91FB-99D55063495C-37712-0000F8C231364FA5.jpg", "content_type"=&gt;"image/jpeg", "byte_size"=&gt;1476387, "checksum"=&gt;"33cpsUeaiJpTT+o6MkZlAQ=="}, "direct_upload"=&gt;{"blob"=&gt;{"filename"=&gt;"image_picker_1249334B-5119-4F5E-91FB-99D55063495C-37712-0000F8C231364FA5.jpg", "content_type"=&gt;"image/jpeg", "byte_size"=&gt;1476387, "checksum"=&gt;"33cpsUeaiJpTT+o6MkZlAQ=="}}}`

`Can't verify CSRF token authenticity.`

`User Load (0.5ms)  SELECT "users".* FROM "users" WHERE "users"."id" = ? LIMIT ?  [["id", 1], ["LIMIT", 1]]`

`↳ app/commands/authorize_api_request.rb:19:in \`user'`

`(0.1ms)  begin transaction`

`ActiveStorage::Blob Create (1.5ms)  INSERT INTO "active_storage_blobs" ("key", "filename", "content_type", "byte_size", "checksum", "created_at") VALUES (?, ?, ?, ?, ?, ?)  [["key", "2qcdc5dzs615rkxf5xgptki4l5pe"], ["filename", "image_picker_1249334B-5119-4F5E-91FB-99D55063495C-37712-0000F8C231364FA5.jpg"], ["content_type", "image/jpeg"], ["byte_size", 1476387], ["checksum", "33cpsUeaiJpTT+o6MkZlAQ=="], ["created_at", "2020-04-17 21:42:41.075445"]]`

`(10.2ms)  commit transaction`

`Disk Storage (4.8ms) Generated URL for file at key: 2qcdc5dzs615rkxf5xgptki4l5pe (`[`http://localhost:3000/rails/active_storage/disk/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaDdDVG9JYTJWNVNTSWhNbkZqWkdNMVpIcHpOakUxY210NFpqVjRaM0IwYTJrMGJEVndaUVk2QmtWVU9oRmpiMjUwWlc1MFgzUjVjR1ZKSWc5cGJXRm5aUzlxY0dWbkJqc0dWRG9UWTI5dWRHVnVkRjlzWlc1bmRHaHBBeU9IRmpvTlkyaGxZMnR6ZFcxSkloMHpNMk53YzFWbFlXbEtjRlJVSzI4MlRXdGFiRUZSUFQwR093WlUiLCJleHAiOiIyMDIwLTA0LTE3VDIxOjQ3OjQxLjExMVoiLCJwdXIiOiJibG9iX3Rva2VuIn19--b2c1b25a821e7ef4b150012ad33d28e5bb6752e8`](http://localhost:3000/rails/active_storage/disk/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaDdDVG9JYTJWNVNTSWhNbkZqWkdNMVpIcHpOakUxY210NFpqVjRaM0IwYTJrMGJEVndaUVk2QmtWVU9oRmpiMjUwWlc1MFgzUjVjR1ZKSWc5cGJXRm5aUzlxY0dWbkJqc0dWRG9UWTI5dWRHVnVkRjlzWlc1bmRHaHBBeU9IRmpvTlkyaGxZMnR6ZFcxSkloMHpNMk53YzFWbFlXbEtjRlJVSzI4MlRXdGFiRUZSUFQwR093WlUiLCJleHAiOiIyMDIwLTA0LTE3VDIxOjQ3OjQxLjExMVoiLCJwdXIiOiJibG9iX3Rva2VuIn19--b2c1b25a821e7ef4b150012ad33d28e5bb6752e8)`)`

`Completed 200 OK in 439ms (Views: 1.9ms | ActiveRecord: 14.4ms | Allocations: 20173)`

&amp;#x200B;

`Started PUT "/rails/active_storage/disk/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaDdDVG9JYTJWNVNTSWhNbkZqWkdNMVpIcHpOakUxY210NFpqVjRaM0IwYTJrMGJEVndaUVk2QmtWVU9oRmpiMjUwWlc1MFgzUjVjR1ZKSWc5cGJXRm5aUzlxY0dWbkJqc0dWRG9UWTI5dWRHVnVkRjlzWlc1bmRHaHBBeU9IRmpvTlkyaGxZMnR6ZFcxSkloMHpNMk53YzFWbFlXbEtjRlJVSzI4MlRXdGFiRUZSUFQwR093WlUiLCJleHAiOiIyMDIwLTA0LTE3VDIxOjQ3OjQxLjExMVoiLCJwdXIiOiJibG9iX3Rva2VuIn19--b2c1b25a821e7ef4b150012ad33d28e5bb6752e8" for ::1 at 2020-04-17 17:42:41 -0400`

`Processing by ActiveStorage::DiskController#update as HTML`

`Parameters: {"encoded_token"=&gt;"eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaDdDVG9JYTJWNVNTSWhNbkZqWkdNMVpIcHpOakUxY210NFpqVjRaM0IwYTJrMGJEVndaUVk2QmtWVU9oRmpiMjUwWlc1MFgzUjVjR1ZKSWc5cGJXRm5aUzlxY0dWbkJqc0dWRG9UWTI5dWRHVnVkRjlzWlc1bmRHaHBBeU9IRmpvTlkyaGxZMnR6ZFcxSkloMHpNMk53YzFWbFlXbEtjRlJVSzI4MlRXdGFiRUZSUFQwR093WlUiLCJleHAiOiIyMDIwLTA0LTE3VDIxOjQ3OjQxLjExMVoiLCJwdXIiOiJibG9iX3Rva2VuIn19--b2c1b25a821e7ef4b150012ad33d28e5bb6752e8"}`

`Completed 422 Unprocessable Entity in 1ms (ActiveRecord: 0.0ms | Allocations: 218)`

&amp;#x200B;

It looks like the first request succeeds, then an internal call is made from within the \`ActiveStorage\` gem that fails. I've done some digging, and the only similar issue I could find is here: [https://github.com/rails/rails/issues/34058](https://github.com/rails/rails/issues/34058)

It seems this check fails, causing the 422: [https://github.com/rails/rails/blob/bfea0af4ba7d717d6a065b4370e3ccfd8869dde6/activestorage/app/controllers/active\_storage/disk\_controller.rb#L22-L26](https://github.com/rails/rails/blob/bfea0af4ba7d717d6a065b4370e3ccfd8869dde6/activestorage/app/controllers/active_storage/disk_controller.rb#L22-L26)

After debugging, it seems this check is failing: \`token\[:content\_length\] == request.content\_length\`, because \`token\[:content\_length\]\` is correct but \`request.content\_length\` is \`0\`.

I'm not really sure where to look next, I think I found the source of the request but \`content\_length\` is set correctly from what I can tell. Any idea what's going on?

I'm making the request from a Flutter frontend using this package: [https://pub.dev/packages/active\_storage/](https://pub.dev/packages/active_storage/). I'd love to test the API directly using Postman but I can't seem to find any info on the request format.
## [8][Help with creating autocomplete search forms with different functions](https://www.reddit.com/r/rails/comments/g3ohjc/help_with_creating_autocomplete_search_forms_with/)
- url: https://www.reddit.com/r/rails/comments/g3ohjc/help_with_creating_autocomplete_search_forms_with/
---
I've gotten bogged down looking through documentation, so I was wondering if anyone might be able to help me. I'm working on a volunteer application which has jobs and description tags. For the search form, I'd like people to be able to have an autocomplete dropdown for either city or postcode. If no current match, you can still input either field and it will return list with the closest current locations.
For description tags, another dropdown that permits multiple tags to be entered.. The search results returned will first be the ones that match all the selected tags, then ones matching some of them.

Thanks in advance for any advice!
## [9][Has anyone ever had extremely high response times on the ActionCable endpoint?](https://www.reddit.com/r/rails/comments/g3hq8q/has_anyone_ever_had_extremely_high_response_times/)
- url: https://www.reddit.com/r/rails/comments/g3hq8q/has_anyone_ever_had_extremely_high_response_times/
---
On localhost, inspector's network tab either shows these /cable requests with *absurd* times:

https://preview.redd.it/q1vjzezt8it41.png?width=2704&amp;format=png&amp;auto=webp&amp;s=08bd5f10b2737ead8ef6cb0af3c9ada16b3d3f69

...or it gives me this nanoscale domino stack of pending requests:

https://preview.redd.it/skp07rw39it41.png?width=2718&amp;format=png&amp;auto=webp&amp;s=4c001441853c3c680129f8f1ffeca02148f57ab8

I have a similar problem with Heroku. I'll be the only user on the site, I'll be doing *nothing*, and the logentries addon will bombard my email with consecutive High Response Time Alerts:

    25 Dec 2019 14:13:31.061299 &lt;158&gt;1 2019-12-25T22:13:28.421009+00:00 heroku router - - at=info method=GET path="/cable" host=www.######.com request_id=#### fwd="####" dyno=web.1 connect=1ms service=531114ms status=101 bytes=174 protocol=https High Response Time

I'm using Ruby version 2.5.0, Rails version 5.2.0, and Puma version: 4.3.1. I've tried both Iodine and AnyCable, and neither fixes this problem. Leading me to believe that the problem is elsewhere. But my ActionCable setup is very standard - an almost exact copy+paste of the setups suggested in docs and tutorials everywhere.

I've found a very small handful of people experiencing a similar issue, but none of them have had any luck. So I'm trying here. Have you experienced this, and did you fix it? If so - how? If not - what could possibly be causing this? Thank you so much!
## [10][Jr dev here having trouble connect Postgres with Ruby](https://www.reddit.com/r/rails/comments/g3t9ks/jr_dev_here_having_trouble_connect_postgres_with/)
- url: https://www.reddit.com/r/rails/comments/g3t9ks/jr_dev_here_having_trouble_connect_postgres_with/
---
Aloha, i made a previous post and i now know that i have to connect my postgres database with my ruby server , but after 2 hours of research and trial and error i am still lost going down the rabbit hole. i do have postgresl setup with my user and password. Here is my config/database for the github repo im cloning : [https://pastebin.com/xgrp9H7s](https://pastebin.com/xgrp9H7s)

Here are the links i have used thus far : [https://www.digitalocean.com/community/tutorials/how-to-use-postgresql-with-your-ruby-on-rails-application-on-ubuntu-14-04](https://www.digitalocean.com/community/tutorials/how-to-use-postgresql-with-your-ruby-on-rails-application-on-ubuntu-14-04)

[https://coderwall.com/p/dovmsg/how-to-make-your-rails-app-start-talking-to-a-postgres-database](https://coderwall.com/p/dovmsg/how-to-make-your-rails-app-start-talking-to-a-postgres-database)

[https://www.dreamvps.com/tutorials/setup-ruby-on-rails-with-postgres/](https://www.dreamvps.com/tutorials/setup-ruby-on-rails-with-postgres/)

but all these refer to a new ruby app not an existing one, can anyone please assist me in connecting my postgres db to my ruby server so i can test out this cool etsy clone github repo. 

mahalo
