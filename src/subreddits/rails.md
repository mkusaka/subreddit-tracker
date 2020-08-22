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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/id911w/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/id911w/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][How to make the comment’s content unique based on its parent post?](https://www.reddit.com/r/rails/comments/ie8zri/how_to_make_the_comments_content_unique_based_on/)
- url: https://www.reddit.com/r/rails/comments/ie8zri/how_to_make_the_comments_content_unique_based_on/
---
I don’t want to make it globally unique in the model file. I just want the content of the comment to be unique if the post already has a comment with same content.
## [4][Alternatives for Bootstrap framework?](https://www.reddit.com/r/rails/comments/iebb66/alternatives_for_bootstrap_framework/)
- url: https://www.reddit.com/r/rails/comments/iebb66/alternatives_for_bootstrap_framework/
---
I have a rails + stimulus app and using Bootstrap for styling. The thing is: Bootstrap depends on jquery, and  this last one is a heavy file after compilation. Because I don't use too much of Bootstrap components and neither Jquery, I would like something more lightweight after compilation. 
 
One more think, I'm not a expertise on css e.g I would need to search how to implement a system grid or how animate dropdown/cascading, for references. 

So my question is: There are others alternatives frameworks that are stable and lightweight? Also, do you guys usually wrote you on components (not talking about JS components e.g react), wrote your own system grid, etc...?

About the Jquery size, I don't remember right now about how much, but I use, sometime ago, a lib that shows by size the files compiled by webpack, and as result the Jquery was the biggest.
## [5][Is it possible to output custom text to the console/server log on startup?](https://www.reddit.com/r/rails/comments/idyetp/is_it_possible_to_output_custom_text_to_the/)
- url: https://www.reddit.com/r/rails/comments/idyetp/is_it_possible_to_output_custom_text_to_the/
---
Bit of a silly request but I want to add some ASCII art to my applications startup so I can see the app logo every time it runs.  

I’m sure it’s possible but where would I need to add my puts statement for this to work?
## [6][Wits-end aws elb ssl eks ingress rails not recognizing https scheme in auth0](https://www.reddit.com/r/rails/comments/ie0kum/witsend_aws_elb_ssl_eks_ingress_rails_not/)
- url: https://www.reddit.com/r/rails/comments/ie0kum/witsend_aws_elb_ssl_eks_ingress_rails_not/
---
First apologies for breaking any rules, we're in panic mode and I didn't have a chance to review rules.
We have in fact searched for a solution frequently and have attempted many fixes suggested on stack overflow / medium etc to know success. 

Long story short
Terminating ssl in aws l7 elb. 
Ingress controller routing host to service-&gt;deployment-&gt;pod-&gt;container of a rails app attempting to use auth0 as the IDP.

Initial site loads properly on https.
However it appears as if generated links (e.g. Redirect and call back urls) are missing awareness of the https scheme.

We've attempted modification of the ingress annotations but seems to be coming up short.

X forward for and proto and scheme and host headers have been tested but we lack visibility into whether or not they are actually getting passed to the rails app.

Working on adding verbosity to that effect now but curious if anyone has experience with this pattern.

Also forgive mobile formatting. 

Thank you!


Just a quick update:
We resolved the issue initially by moving ssl termination into the ingress controller and changing from L7 load balancer to L4 load balancer. This isn't ideal and will continue to investigate the L7 config.  I'll be sure to post an update when we resolve.
## [7][Newbie dev here, got some questions about Rails as a backend](https://www.reddit.com/r/rails/comments/idxys7/newbie_dev_here_got_some_questions_about_rails_as/)
- url: https://www.reddit.com/r/rails/comments/idxys7/newbie_dev_here_got_some_questions_about_rails_as/
---
\- Can you use Rails as a back end for a react native app? 

And if you can, could anyone link me to resources that explain how to do it
## [8][JSON API specification for rails api](https://www.reddit.com/r/rails/comments/idppfn/json_api_specification_for_rails_api/)
- url: https://www.reddit.com/r/rails/comments/idppfn/json_api_specification_for_rails_api/
---
Which one do you guys use or prefer using for rails api response in json api specification.

1. [fast_jsonapi](https://github.com/Netflix/fast_jsonapi)    now forked and mainted as      [jsonapi-serializer](https://github.com/jsonapi-serializer/jsonapi-serializer) 

2. [jsonapi-rb](https://github.com/jsonapi-rb/jsonapi-rb) (latest activity 3 years ago)

3. Others( would be great if you specify :) )
## [9][Best stack for a new SaaS side project?](https://www.reddit.com/r/rails/comments/ido300/best_stack_for_a_new_saas_side_project/)
- url: https://www.reddit.com/r/rails/comments/ido300/best_stack_for_a_new_saas_side_project/
---
I'm currently planning out a SaaS side project and I'm trying to figure out the best stack for a modern Rails app. I've done small projects with both a monolith and Rails APIs, so I'm really fine with either approach. One goal here is to up my frontend game. The options I'm considering are as follows:

* 1. Rails API with a separate React frontend
   * Pros: I've done this sort of stack before; API is relatively quick to set up and deploying a separate frontend to Netlify can be convenient.
   * Cons: There are some security concerns with a JWT auth solution, which is what I usually use for APIs. Additionally, I'll be implementing multi-tenancy/admin roles and views, so I'm not exactly sure how well that sort of thing would work with a decoupled frontend/backend.
* 2. Monolith Rails + React
   * Pros: Security would concern me less with a full-on Devise/sessions set up. No security solution is perfect, but I'd feel better about my ability to set up auth in a monolith vs putting together a JWT solution.
   * Cons: I'm not sure I'm 100% sold on the idea of using React within Rails; it would require a bit of set up but I know people do it all the time, so maybe I'm overthinking it. Using a robust FE library like that within Rails would just be a new thing, I guess.
* 3. Monolith Rails + StimulusReflex
   * Pros: I'm super interesting in learning StimulusReflex. Maybe it's just HEY-related hype, but I like the way HEY's app works and it seems like a great way to get a reactive frontend without React or \[insert other framework/library here\].
   * Cons: TBH, it's kind of hard to find stuff about SR beyond basic examples since it's relatively new. I'm also concerned about using CableReady, since my understanding is that disconnections can cause issues(?). I also haven't seen that many posts about SR on Reddit, so I'd be curious to get thoughts on the notion of using SR in a possible production environment.
* 4. Pure Monolith Rails with good ol' ERB.
   * Eh, it works. I've done it before. Gets the job done.

I'd like to take my Rails skills further and try something new, so I'm really leaning towards Rails + StimulusReflex; that seems to be a bit of a wave in terms of the future direction of Rails FE development, but I'd love to get some thoughts from other folks! I'm a relatively new Rails developer (been using Rails for about a year) and don't currently use Rails in the context of my job, so I'm sure there's stuff I'm missing/not taking into consideration.

TL;DR: If you were starting a brand new SaaS side project in Rails and weren't tied down stack-wise, which direction would you go for a modern Rails app?

Edit: some formatting/fixing a sentence or two
## [10][Follower/Following system using devise](https://www.reddit.com/r/rails/comments/idremq/followerfollowing_system_using_devise/)
- url: https://www.reddit.com/r/rails/comments/idremq/followerfollowing_system_using_devise/
---
I searched on the internet, I probably found 10 or more articles, saying the same thing and most of them are incomplete. 

Follower/Following system is important in most cases nowadays, a lot of employers think their website might look much fancier with this ability. Anyway, I'm not here to discuss how employers annoy me, I'm just curious, isn't there any easy and understandable way to implement follower/following system in a rails app?
## [11][Handling Multiple Points of Relation](https://www.reddit.com/r/rails/comments/idn1zs/handling_multiple_points_of_relation/)
- url: https://www.reddit.com/r/rails/comments/idn1zs/handling_multiple_points_of_relation/
---
Hey everyone!

I would appreciate your help with this. Not really sure what to look for on StackOverflow. My Googling hasn't led me far. 

&amp;#x200B;

*Example:*

**I have the following tables:** 

1. Records. Like vinyl records. 
2. Storage Locations. Think: A shelf or a cabinet. 
3. Boxes: A box that a vinyl record may or may not be in.

**Some parameters:** 

A box must have a storage location. 

A vinyl record may or may not be in a box. If it is not in a box, it is just sitting directly in the storage location, e.g. directly in the cabinet. 

&amp;#x200B;

**So . . .**

This means that I have two possible points of relation b/w records &amp; storage locations. In some cases they relate directly, and in others, they relate through the box. 

&amp;#x200B;

**Are there best practices for handling this type of relationship?** Right now, I've just been doing a lot of "*if* [record.box](https://record.box) *then* record.box.storage\_location *else* record.storage\_location" statements. 

Thoughts?

Any articles or best practices would be super helpful! Thanks :-D
## [12][Javascript Library loading twice in Rails using Turbolinks](https://www.reddit.com/r/rails/comments/idfohv/javascript_library_loading_twice_in_rails_using/)
- url: https://www.reddit.com/r/rails/comments/idfohv/javascript_library_loading_twice_in_rails_using/
---
My javascript library is loading twice. I have consulted..

[https://stackoverflow.com/questions/34388869/javascript-library-loads-twice-in-rails](https://stackoverflow.com/questions/34388869/javascript-library-loads-twice-in-rails) (taking include\_tag out of application.html.erb)

[https://github.com/turbolinks/turbolinks/issues/403](https://github.com/turbolinks/turbolinks/issues/403) (changing  &lt;%= javascript\_include\_tag 'application', 'data-turbolinks-track': 'reload' %&gt; in my application.html.erb file to  &lt;%= javascript\_include\_tag 'application', 'data-turbolinks-track': 'false' %&gt; )

&amp;#x200B;

Has anyone come along this issue, and have a solution?
