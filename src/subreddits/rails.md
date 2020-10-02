# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/iya619/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/iya619/personal_projects_show_off_your_own_project_andor/
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
## [3][Graphql Authentication](https://www.reddit.com/r/rails/comments/j3mxrd/graphql_authentication/)
- url: https://www.reddit.com/r/rails/comments/j3mxrd/graphql_authentication/
---
Hello everyone,

Basically I know how to do everything with react/Apollo but faulty new with using rails/graphql, especially authentication. Like, I don’t even know where to start and couldn’t find any solid resources. I’m looking to use graphql and devise. The graphq-devise gem confused me. The boilerplate option was nice but I want to know how to build it from the ground up.
## [4][How to block users from getting messages?](https://www.reddit.com/r/rails/comments/j3eszd/how_to_block_users_from_getting_messages/)
- url: https://www.reddit.com/r/rails/comments/j3eszd/how_to_block_users_from_getting_messages/
---
Hey everyone,

I have an internal messaging system in Rails 6.

I was wondering how can I add the logic for a button to block the user in conversation that I'm messaging with.

&amp;#x200B;

[This is my schema](https://preview.redd.it/77mdw4g06jq51.png?width=965&amp;format=png&amp;auto=webp&amp;s=ee978661b4368e764942b42a3524d768fab21387)

&amp;#x200B;

[MessagesController](https://preview.redd.it/xefpdwnicjq51.png?width=828&amp;format=png&amp;auto=webp&amp;s=bc3932ac6a751504eae9c08d478bd2b0b1e20bd8)

&amp;#x200B;

[Messages Index](https://preview.redd.it/6vin53ancjq51.png?width=1108&amp;format=png&amp;auto=webp&amp;s=249910e94cba48f3386cd3981ea1db682b736f08)

&amp;#x200B;

[Conversations Controller](https://preview.redd.it/ddttqjdpcjq51.png?width=828&amp;format=png&amp;auto=webp&amp;s=5efc31c7f8acb40124c4facdd91ecd2021e3b56b)

&amp;#x200B;

[Conversations Index](https://preview.redd.it/hfqqb3ktcjq51.png?width=1331&amp;format=png&amp;auto=webp&amp;s=6dad341ffe8cbc0c0abb52111ea5c58fbeafe8eb)

&amp;#x200B;

Edit: Controllers and views are added per request.

Can I do it with my current setup or do you advice me to create another table like friendships and use that table to log the blocked users?

&amp;#x200B;

Thank you.
## [5][Does hosting streaming video get pretty expensive?](https://www.reddit.com/r/rails/comments/j2xzir/does_hosting_streaming_video_get_pretty_expensive/)
- url: https://www.reddit.com/r/rails/comments/j2xzir/does_hosting_streaming_video_get_pretty_expensive/
---
I've had an idea for a site that would host instructional videos and possibly integrate with a more traditional web app. I was thinking about doing recurring billing through Stripe and having a mix of free videos on YouTube and subscriber videos on the site.

If I were to host the streaming videos on something like S3, would the costs be pretty significant? I currently host a site using ActiveStorage through S3 and my bandwidth costs are reasonable, but not nothing. I pay more for the database and load balancer. Would I need to encode them in something friendly to streaming and find a player that isn't awful? Am I better off going with a service that facilitates paid access to videos, even if they take their cut?
## [6][Debugging why sometimes response is slow sometimes fast.](https://www.reddit.com/r/rails/comments/j31v72/debugging_why_sometimes_response_is_slow/)
- url: https://www.reddit.com/r/rails/comments/j31v72/debugging_why_sometimes_response_is_slow/
---
I have rails product which exposes certain API. We have a sister application that needs to call certain number of api in sequence. For example

Complete procedure - Sync a user

\- Create a new user  - v1/users

\- Sync his credit cards - v1/cards

\- Generate his bills - v1/bills

Now in have a staging server, with 5 instances, 2 threads each (unicorn). So thats total 10 requests per second. The sister application might need to do the complete sync for maybe 500 users in one go. So its hitting 2 users concurrenctly. Waiting for it to finish then 2 more, and so on. The problem is when the requests start the response times are pretty good. But as the users keep syncing in the response times get incredibly slow. So variations include 300ms, 4 seconds, 16 seconds, back to 4 seconds. Since all of this is test data, almost all users/cards/bills are pretty much the same(API are hit similar params).

How do I find what the bottleneck is? Could it be the db , Im running a \`db.t2.micro\` on aws for it. But even on peak requests the cpu utilization never exceeded 15 %. , 10 DB connections.

Same goes for the EC2 instances, CPU utilisation never seems to exceed 20 percent
## [7][Are monoliths becoming cool again?](https://www.reddit.com/r/rails/comments/j2kuxq/are_monoliths_becoming_cool_again/)
- url: https://www.reddit.com/r/rails/comments/j2kuxq/are_monoliths_becoming_cool_again/
---
Just a gut feeling, but it seems like micro-service everything for small-medium companies is getting a bit out of flavour, and monoliths start to make more sense again. The hype pendulum keeps swinging.

This together with some JS fatigue and the new Deno project (which might deprecate Node) makes me think choosing boring old monolith tech like Rails may become cool again. Or am I reading too much into this?
## [8][[help] `invalid_grant` - Google authorization for grabbing GMails from the API.](https://www.reddit.com/r/rails/comments/j2njw9/help_invalid_grant_google_authorization_for/)
- url: https://www.reddit.com/r/rails/comments/j2njw9/help_invalid_grant_google_authorization_for/
---
I am working on a Rails backend, where I'm implementing support for grabbing GMails from a Google service API. When using [google-auth and its guide](https://github.com/googleapis/google-auth-library-ruby#example-web) for the project, I get to the point where I have a token stored in Redis, but when trying to use it I get a BadRequest and \`invalid\_grant\` thrown back.

It is not due to a clock mismatch and the tokens are fresh. Does anyone have experience with this problem?

Optionally; does anyone know of alternative ways of grabbing mails through the Google API?
## [9][ApplicationController.render could generate an invalid CSRF token– and you have one more way for a CSRF token to annoy you.](https://www.reddit.com/r/rails/comments/j2hl1o/applicationcontrollerrender_could_generate_an/)
- url: https://www.reddit.com/r/rails/comments/j2hl1o/applicationcontrollerrender_could_generate_an/
---
A recent annoying CSRF problem in our platform. Hope you could find the information useful.

Basically there was an ActionController::InvalidAuthenticityToken error on deleting a group.

This was occurring because we were using ApplcationController.render to render one of the buttons for a delete request. This was not working. ApplicationController has no information about the session. We should have used 'render'.

Be careful when using ApplicationController.render for rendering forms. 

[https://kmitov.com/posts/applicationcontroller-render-one-more-way-for-a-csrf-token-to-annoy-you/](https://kmitov.com/posts/applicationcontroller-render-one-more-way-for-a-csrf-token-to-annoy-you/)
## [10][Deploy Ruby on Rails on Google cloud](https://www.reddit.com/r/rails/comments/j26lkz/deploy_ruby_on_rails_on_google_cloud/)
- url: https://www.reddit.com/r/rails/comments/j26lkz/deploy_ruby_on_rails_on_google_cloud/
---
Hi, does anyone know what's the best way to deploy a ruby on rails application on google cloud also using the SQL instance on Google Cloud and a process to use sidekiq.

Do you guys have any pratical experience?
## [11][Trying to upload the file to rails server from React? I receive permitted: false in parameters of file](https://www.reddit.com/r/rails/comments/j25sdg/trying_to_upload_the_file_to_rails_server_from/)
- url: https://www.reddit.com/r/rails/comments/j25sdg/trying_to_upload_the_file_to_rails_server_from/
---
I'm using active\_storage to upload pictures. The file doesn't show on rails. Is the something with rails or Javascript? Where should I be looking

&amp;#x200B;

&lt;ActionController::Parameters {} permitted: false&gt;

&amp;#x200B;

Update: I was able to send formData and get the file through. What I did differently was I did fetch request separate for formData separately instead of sending it within a hash of other data like {username: username, password: password, file: formData}. I don't know why this is though and would appreciate an answer
## [12][Using Carrierwave, cannot access image_url in view](https://www.reddit.com/r/rails/comments/j2694q/using_carrierwave_cannot_access_image_url_in_view/)
- url: https://www.reddit.com/r/rails/comments/j2694q/using_carrierwave_cannot_access_image_url_in_view/
---
Images are saving to the database and are nested within a parent. Parent model is Nitrogen, Child model is Photos. For example, in the view when I try to access the image by  &lt;%=nitrogen.photos.ids%&gt;, the id is displayed.  Attached is a image of what I get in the rails console when I call Nitrogen.last.photos. How can I extract :picture within the ActiveRecord::Associations::CollectionProxy?
