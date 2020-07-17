# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hrnm2o/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hrnm2o/personal_projects_show_off_your_own_project_andor/
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
## [3][Cryptic error. Any clue?](https://www.reddit.com/r/rails/comments/hsubkv/cryptic_error_any_clue/)
- url: https://www.reddit.com/r/rails/comments/hsubkv/cryptic_error_any_clue/
---
Hi Guys.

So far I have been using capistrano for my production app, and all was good, but trying to set up a staging environment, I get an error in the last moment and I don't know how to interpret it well.  


Can you give me any help?

&amp;#x200B;

https://preview.redd.it/t045d3jvoeb51.png?width=1666&amp;format=png&amp;auto=webp&amp;s=04de65012fc28e162b404dfcc78bfa2cfddfc95f

https://preview.redd.it/lg8z25jvoeb51.png?width=1666&amp;format=png&amp;auto=webp&amp;s=aaffabffcef25030478a08d0490178e86bd6e505
## [4][Writing a good Rails app CSP, and how to build a report_uri](https://www.reddit.com/r/rails/comments/hspydx/writing_a_good_rails_app_csp_and_how_to_build_a/)
- url: https://www.reddit.com/r/rails/comments/hspydx/writing_a_good_rails_app_csp_and_how_to_build_a/
---
The [official documentation](https://edgeguides.rubyonrails.org/security.html#content-security-policy) suggests starting with a global policy, and making changes as needed. In trying to understand this, I've made some very minor changes:

    Rails.application.config.content_security_policy do |policy|
      policy.default_src :self, :https
      config.font_src    :self, :https, 'fonts.googleapis.com'
      policy.img_src     :self, :https, :data, 'blahblah.s3-us-east-1.amazonaws.com'
      policy.object_src  :none
      config.script_src  :self, :https, 'stripe.com'
      policy.style_src   :self, :https
    
      # Specify URI for violation reports
      policy.report_uri "/csp-violation-report-endpoint"
    end

Two questions about this. 

1. Should excluding the `blahblah.s3-us-east-1.amazonaws.com` from `policy.img_src` effectively prevent those resources from loading? I would think so, but the resource still loads.
2. Same for `fonts_src.` Should excluding `fonts.googleapis.com` prevent that font from loading? I get a console warning that says it's not loading due to a Content Security Policy violation, but the font still seems to be used when checking the styles in Inspector.
3. Exactly what is this `report_uri`? I can seem to find anything that explains it. Presumably it needs a `POST` route, a controller, and a view accessible only to an admin?
## [5][How to use a Transaction Script(aka Service Objects) in Ruby on Rails. Simple example](https://www.reddit.com/r/rails/comments/hssf51/how_to_use_a_transaction_scriptaka_service/)
- url: https://www.reddit.com/r/rails/comments/hssf51/how_to_use_a_transaction_scriptaka_service/
---
The logic of small applications can be present as a series of transactions. Using the Transaction Scripts pattern, we get an application that is easier to maintain, to cover with tests, and to scale.

In the [tutorial](https://jtway.co/how-to-use-a-transaction-script-aka-service-objects-in-ruby-on-rails-simple-example-161b7e228942) we will develop an [application](https://github.com/dgorodnichy/transaction-script-example) that has Post, User, and Like models. Users should be able to like posts. The first version of the controller will contain extra code, which we will extract into a separate Transaction Script. We also describe when we need to use the Transaction Scripts and the pros of the transaction script usage.  


Full tutorial: [How to use a Transaction Script (aka Service Objects) in Ruby on Rails. Simple example](https://jtway.co/how-to-use-a-transaction-script-aka-service-objects-in-ruby-on-rails-simple-example-161b7e228942)
## [6][Best way to do this](https://www.reddit.com/r/rails/comments/hsuxs4/best_way_to_do_this/)
- url: https://www.reddit.com/r/rails/comments/hsuxs4/best_way_to_do_this/
---
Hello,

Am in building an app using activeadmin for the admin functionalities and am trying to integrate with Plivo (like Twilio). In order for the app to receive sms's sent to my Plivo number, I need to supply a post route to be called. My question is how to provide such a post route when all routes are behind the authenticated "admin" parent route? Thanks.
## [7][Hi everyone.](https://www.reddit.com/r/rails/comments/hsolu9/hi_everyone/)
- url: https://www.reddit.com/r/rails/comments/hsolu9/hi_everyone/
---
HI everyone, i hope that everyone is doing good through these times we are experiencing and everyone is safe, i was just wondering if ruby and rails is still the best framework for prototyping? coming from python :). 

any advice would be appreciated. Thank you
## [8][Is innovation needed anyway ?](https://www.reddit.com/r/rails/comments/hsskui/is_innovation_needed_anyway/)
- url: https://www.reddit.com/r/rails/comments/hsskui/is_innovation_needed_anyway/
---
Here is a tweet quote from excellent Chris Oliver (GoRails) "If we want the Ruby community to grow, we need to get back to innovating and talking about it." My feeling is : is there any need to innovate anyway ?   
Even the JS hype of the last years concern only a few pages that need more interactions... only small parts of a standard business app. My feeling is that innovation do not concern anymore the vast majority of every day problems. Any thought ?
## [9][How are background jobs scaled?](https://www.reddit.com/r/rails/comments/hsktek/how_are_background_jobs_scaled/)
- url: https://www.reddit.com/r/rails/comments/hsktek/how_are_background_jobs_scaled/
---
I am wondering on how backend jobs are scaled.

If you have a background job that is running but another user initiates a job at the same time, is there a way to "split" or multiple the job into 2 servers automatically so they run in parallel instead of linearly?

Like, what would an app do if 1000 users hit submit on a form that would then submit to a background job that each would take 10 mins?  Like, how would this be done in parallel so the 1000th user doesn't need to wait 10,000 minutes?

Would a situation like this just need to have thousands of workers just waiting? How can you implement this dynamically?

I've been looking up how to concurrently or in parallel to background jobs but can't find many with good explanations
## [10][First freelance site - Spree Commerce RoR - thoughts?](https://www.reddit.com/r/rails/comments/hscuc0/first_freelance_site_spree_commerce_ror_thoughts/)
- url: https://www.reddit.com/r/rails/comments/hscuc0/first_freelance_site_spree_commerce_ror_thoughts/
---
 Hi Guys,

Short time scroller first time poster. Started coding at the start of lockdown (late March here in the UK), first on Codecademy, then did a React Nanodegree with Udacity and found my way to Hartl's Rails tutorial. Reached out to a small business to offer to build an ecommerce site to assist with their business - pushed myself in the deep end to be honest but enough I felt I knew enough to start freelancing!

Let me know your thoughts: [https://voyage-fromage-60583.herokuapp.com/](https://voyage-fromage-60583.herokuapp.com/)

Modified Spree Commerce App. I have the .co.uk domain but not pushing until the client is ready.

If you're still reading (thanks), an eye on my portfolio site would be great! [https://www.badenashford.me](https://www.badenashford.me/)

Cheers!!
## [11][Rails &amp; iSeries DB2...Again](https://www.reddit.com/r/rails/comments/hsks6x/rails_iseries_db2again/)
- url: https://www.reddit.com/r/rails/comments/hsks6x/rails_iseries_db2again/
---
Alight guys...Has anyone successfully connected a rails app...hell, even from a ruby console...to the DB2 database (library) on an iSeries, AS400, System i? I'm trying to use the ODBC driver on Linux. When I try to connect it seems to hang indefinitely. I can use and ODBC query tool on my windows PC and it works great. I can't get rails on my PC to connect either. Any insights here?
## [12][Show a pre-filled form when clicking "Back" from a form submission?](https://www.reddit.com/r/rails/comments/hsjgbt/show_a_prefilled_form_when_clicking_back_from_a/)
- url: https://www.reddit.com/r/rails/comments/hsjgbt/show_a_prefilled_form_when_clicking_back_from_a/
---
Is this possible?

I have a scenario where submitting a form shows a calculation and would prefer to not show the form on the resulting screen. The form does not create a record so multiple submisisons are not an issue. 

The natural UX for updating for values is hitting back on the browser, fixing the values and re-submitting.

But, the default behavior of Rails in this case is to show the empty "new" state before entering any values.

Any way to show the form in its filled state on hitting the "back" button?

It's the default browser behavior for forms which do not involve a redirect.

I've explored other options, like a link to "edit" state, but would like to see if I can stick with the simplest possible most intuitive default.
