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
## [3][Rails Live-streaming: Adding ‘admin’ functionality to a Rails app](https://www.reddit.com/r/rails/comments/j4xhn0/rails_livestreaming_adding_admin_functionality_to/)
- url: https://www.reddit.com/r/rails/comments/j4xhn0/rails_livestreaming_adding_admin_functionality_to/
---
Hi, every week I livestream some Ruby on Rails coding, building an app and looking at different approaches to features from a learners perspective. In recent weeks we [created our initial models and associations](https://youtu.be/vQjw5uUAx5k), last week we [added some forms along with `edit` and `update` actions on a controller](https://youtu.be/MWSngfjNKS0).

This week, I wanted to run through some different approaches for allowing admins (or any other type of user) to edit/update a model. We looked at where this type of behaviour can live inside your Rails app and then implemented one of the options.

Here's the link: https://www.youtube.com/watch?v=TbRw2jHpgg0

I'll add in some chapter markers as soon as I can to make this easier to navigate, I've already done some of the previous videos.

I usually stream 10:30am (UK Time) every Sunday. You can subscribe to [the YouTube channel](http://www.youtube.com/channel/UC6JPQW8LnAZyfC_rUZqDqqA) if you want to be reminded. Thanks!
## [4][s3 signed urls question](https://www.reddit.com/r/rails/comments/j4qcxj/s3_signed_urls_question/)
- url: https://www.reddit.com/r/rails/comments/j4qcxj/s3_signed_urls_question/
---
After uploading a file to S3 using Shrine and Uppy (which requires a presigned url), and saving its data to my model, I get to call `image_url` in image tags and such. However....something about having all that credential information as part of a url that I want to be a publicly readable file strikes me as sort of...odd? If my understanding is correct, and it's used for signing uploads, is there any reason why all this remains in the url? If I copy the image url from most sites, I do not see S3 credentials. Would this hurt me in any way?

I go to my console and type `model.image_url` (or just copy image url in the browser) and get this:

`https://bucket.s3.region.amazonaws.com/model/1/image/9457719e40fb416e01a153657d7ea4fb.png?X-Amz-Algorithm=###&amp;X-Amz-Credential=accesskey_date_region_s3aws4_request&amp;X-Amz-Date=###&amp;X-Amz-Expires=900&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=###`

Is it ok that X-Amz-Algorithm, X-Amz-Credential, X-Amz-Date, X-Amz-Expires, X-Amz-SignedHeaders, and X-Amz-Signature are public like this?

X-Amz-Credential is just my access key and date and such, so it doesn't seem harmful. But what about X-Amz-Signature?

And X-Amz-Expires...what is it that expires here? According to [this](https://forums.aws.amazon.com/thread.jspa?threadID=284629) a value of 900 is equal to 15 minutes, however I've been able to view this image hours later. Why would I want this?
## [5][Why not to use Rails on Windows?](https://www.reddit.com/r/rails/comments/j4va9b/why_not_to_use_rails_on_windows/)
- url: https://www.reddit.com/r/rails/comments/j4va9b/why_not_to_use_rails_on_windows/
---
I'm starting to learn rails and heard a lot of recommendations to use it on Linux, but no one tells me exactly why. So, I prepared a setup on Ubuntu but tried to give a chance do Windows and made my first small api, with login system, etc... but couldn't spot any problem using Windows.

So, can anyone give me reason to drop Windows now instead of running into big issues later, or is it just some kind of a programmer tabu?

&amp;#x200B;

EDIT: Also, I found it way easier to setup an developer environment to Rails on Windows than on Ubuntu, so I don't get the "configuration nightmare" that people tells about.
## [6][Would anyone be interested in a web jam?](https://www.reddit.com/r/rails/comments/j3v10t/would_anyone_be_interested_in_a_web_jam/)
- url: https://www.reddit.com/r/rails/comments/j3v10t/would_anyone_be_interested_in_a_web_jam/
---
For awhile I have been looking for an outlet for competitive web development but haven't found anything. In turn, we are hosting our own competitive web development jam tomorrow!

It will be open up tonight at 12am EST and you will have 24hours(all day Saturday) to complete the website!

Tonight we will be posting the theme we have picked and the criteria you must complete for your submission. It must be submitted by 11:59pm on Saturday Oct 3rd EST.

You are allowed to utilize any web technologies you want to build the website, except drag and drop programs. You do have to build the site from scratch and cannot utilize a template.

There is an optional $2 to get a ticket to compete for the cash prize! Otherwise it's free for anyone wanting to just get some experience or have some fun!

For anyone that's interested in still signing up, here is the link :) Good luck everyone!

[https://discord.gg/5rcKpTk](https://discord.gg/5rcKpTk)
## [7][anyone ever try to put authenticated views in cdn?](https://www.reddit.com/r/rails/comments/j3ydbr/anyone_ever_try_to_put_authenticated_views_in_cdn/)
- url: https://www.reddit.com/r/rails/comments/j3ydbr/anyone_ever_try_to_put_authenticated_views_in_cdn/
---
wondering how you approached security, which APIs you used (conditions, token).
## [8][Rails API only application; Native App + Doorkeeper = very confused fellow developer](https://www.reddit.com/r/rails/comments/j3wuu8/rails_api_only_application_native_app_doorkeeper/)
- url: https://www.reddit.com/r/rails/comments/j3wuu8/rails_api_only_application_native_app_doorkeeper/
---
Hey, I promise you I've gone through all the posts on this subreddit I could find on this topic. I have some beginner questions as I've only dived into my first rails project in May (which is running just fine :D)

I now need to add proper authentication to it

I initially used Devise + a custom JWT strategy but that wouldn't work for me in the long run. Long story short, I need an Oauth mechanism for creating an Alexa Skill. 

Anyway, my confusions are :

1. In my app (a flutter one in this case) I would like to start off with a simple sign in with email and later scale to social logins. Can I do this without leaving the app (not going into the browser?)

2. If not, do I need to create a custom view in my API only rails app just for handling the login that I need to send my native app users to, to authenticate?

3. What is the redirect URI in this case for doorkeeper? 

I can figure out the omniauth bit and the specifics as and when it comes, but my fundamentals here in understanding this process keeping rails and doorkeeper in mind aren't clear. 

I will buy you a beer because this has been really stressful figuring out and because I love this community in general haha

Thanks in advance!
## [9][View helpers - Is this approach "acceptable" or is it horrible?](https://www.reddit.com/r/rails/comments/j3utnl/view_helpers_is_this_approach_acceptable_or_is_it/)
- url: https://www.reddit.com/r/rails/comments/j3utnl/view_helpers_is_this_approach_acceptable_or_is_it/
---
Hi,

I have a few custom UI components in my app, such as a dropdown menu. Writing those up manually every time I need one is not so cool, considering that it's quite a bit of HTML to write and if I need to update the dropdown markup once, I have to do it across many, many views.

Therefore, I wrote those two helpers: [\*click\*](https://gist.github.com/ImMaax/9d9c6ba45eba6664b87a43e0c93fa168)

Do you think this is a good approach or is it completely horrible? Pretty much my first time working with blocks in Ruby. Thanks a lot!
## [10][Logging rendering time of templates retrieved with s custom template resolver](https://www.reddit.com/r/rails/comments/j3uy0r/logging_rendering_time_of_templates_retrieved/)
- url: https://www.reddit.com/r/rails/comments/j3uy0r/logging_rendering_time_of_templates_retrieved/
---
In my CMS I manage user editable templates that are stored in the database and retrieved with a custom template resolver. For these templates/views Rails' logs don't show rendering time like for the standard views. Is there a way to fix this? For example by hooking into ActionView LogSubscriber? Thanks!
## [11][Comments problem in my app](https://www.reddit.com/r/rails/comments/j3uzd1/comments_problem_in_my_app/)
- url: https://www.reddit.com/r/rails/comments/j3uzd1/comments_problem_in_my_app/
---
I made this app:

[https://github.com/prp-e/railstagram](https://github.com/prp-e/railstagram)

as a fun project. Each time I need to learn something new, I mess around with this project. The newest thing I wanted to learn was adding comments to the posts. So, I found this tutorial:

[https://guides.railsgirls.com/commenting](https://guides.railsgirls.com/commenting)

and I followed it and did a little bit of change to get it to work in my way. Now, when I want to add a comment I get this :

&amp;#x200B;

https://preview.redd.it/81rso25vtoq51.png?width=614&amp;format=png&amp;auto=webp&amp;s=cb241f2876af487c7ea2eec845761654bc380775

I can solve the "User" problem by adding `current_user.comments.build` to my comments controller. But I don't know what should I do for the post to get it to the work. I'd be thankful if you guys help me for this.
## [12][Graphql Authentication](https://www.reddit.com/r/rails/comments/j3mxrd/graphql_authentication/)
- url: https://www.reddit.com/r/rails/comments/j3mxrd/graphql_authentication/
---
Hello everyone,

Basically I know how to do everything with react/Apollo but fairly new with using rails/graphql, especially authentication. Like, I don’t even know where to start and couldn’t find any solid resources. I’m looking to use graphql and devise. The graphq-devise gem confused me. The boilerplate option was nice but I want to know how to build it from the ground up.
