# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/fx6je4/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/fx6je4/personal_projects_show_off_your_own_project_andor/
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
## [3][Looking for a sales/demand forecasting gem/API](https://www.reddit.com/r/rails/comments/g1qu36/looking_for_a_salesdemand_forecasting_gemapi/)
- url: https://www.reddit.com/r/rails/comments/g1qu36/looking_for_a_salesdemand_forecasting_gemapi/
---
I have to build a product that does some sales forecasting. Currently the client is doing it in Excel with the Excel stat functions. I was wondering if anyone knew any gems that could do something similar where you feed it in a series of data and it forecasts how much sales will happen in the future. 

If no well maintained gem is available I am also open to a third party app that do something similar with an API I can connect to.

Thanks in advance.
## [4][Grid-based document creation tools/gems/solutions?](https://www.reddit.com/r/rails/comments/g1pm4j/gridbased_document_creation_toolsgemssolutions/)
- url: https://www.reddit.com/r/rails/comments/g1pm4j/gridbased_document_creation_toolsgemssolutions/
---
Excel is a grid-based document creation application. But it requires a lot of upfront planning because users can't split any cell into multiple cells to make sure the contents inside are aligned (only merge is allowed).

Is there anything else out there?
## [5][Choice of System test drivers](https://www.reddit.com/r/rails/comments/g1p7di/choice_of_system_test_drivers/)
- url: https://www.reddit.com/r/rails/comments/g1p7di/choice_of_system_test_drivers/
---
TLDR: I would like to know how you choose which driver you use in System Tests (or End-to-End).

When you are writing system tests the default is selenium (slow but controlling a real browser) but another common option is rack\_test (fast but it 'fakes'^(1) a browser). [The choice of Selenium was because it was zero setup and supported JavaScript out the box.](https://github.com/rails/rails/pull/26703#issuecomment-251649333)

&amp;#x200B;

What I have done: 

I have used rack\_test for most of my end to end testing unless I am writing JavaScript in which case I used Selenium (or one of the other JavaScript aware drivers).

Is rack\_test first common or should I be using Selenium more often? I guess I'm wondering if there's times when you reach for Selenium first (other than JS).

&amp;#x200B;

^(1)  \- not sure of the right term it is not a browser but is parses the HTML but doesn't know about JavaScript?
## [6][Authentication for both web app and mobile app](https://www.reddit.com/r/rails/comments/g1g1ne/authentication_for_both_web_app_and_mobile_app/)
- url: https://www.reddit.com/r/rails/comments/g1g1ne/authentication_for_both_web_app_and_mobile_app/
---
I have a Rails app in API mode serving both a JS web app and a mobile app. I did some research on authentication and went with JWT implementation with Sorcery.  I'm aware of Devise for authentication but since it's more for MVC Rails app instead. Now, everything works pretty well for both platforms for creating users, logging in etc.

However, now I have to implement token refreshing mechanism and I found this post [https://stackoverflow.com/questions/55486341/how-to-secure-a-refresh-token](https://stackoverflow.com/questions/55486341/how-to-secure-a-refresh-token) saying that storing a refresh token in a browser is not a good idea but for mobile devices it's ok.  So, if using JWT for web apps is not a good idea and storing the refresh token in a browser is a no no then I wonder if I should have both cookie-based auth for traffic from the web app and JWT-based auth for the mobile app instead. Any one has done something similar?

&amp;#x200B;

For JWT implementation I followed this [tutorial](https://qiita.com/kaishuu0123/items/27ba55e774ac1b7a94fc) (it's in Japanese but I mainly just followed the github)

&amp;#x200B;

PS. I'm not using OAuth at the moment since I didn't want 3rd party auth service
## [7][React: Handling associations from Rails API in state?](https://www.reddit.com/r/rails/comments/g1rdj4/react_handling_associations_from_rails_api_in/)
- url: https://www.reddit.com/r/rails/comments/g1rdj4/react_handling_associations_from_rails_api_in/
---
So hopefully this question will make sense, but lets say I have a Rails API serving JSON to a React front-end.

Lets say I also have a model: \`widget\`, and \`widget\` has a bunch of different associations (like has\_one \`widget\_type\` among other associations.

On the React end in my "Table" component I want to have a table that has all the \`widget\` attributes on a table row (and state in this case is just an array of objects). But lets say I also want to include some data from \`widgets\` associations on each row as well?

Whats the best way to do that? (Sidenote: Lets pretend I am using fetch/axios to make the call to the Rails API on the table component load).

I imagine I maybe use some sort of json serializer/builder to include the extra data? I can't find a distinct answer unfortunately.
## [8][Struggling with Rails api social logins using Devise.](https://www.reddit.com/r/rails/comments/g1dthu/struggling_with_rails_api_social_logins_using/)
- url: https://www.reddit.com/r/rails/comments/g1dthu/struggling_with_rails_api_social_logins_using/
---
I need to allow for signing up and logging in using Apple, Google and Facebook.  I've done it on a full stack Rails project years ago, but never with an api only setup.  My understanding is that the front end app will handle logging in, and then will pass a token to the server where I will then have to verify it.  I've seen tutorials, but they seem to be all over the place.  Is there a good starting point where I can see the correct way to implement this in rails?
## [9][Junior dev just recently got first job looking for good resources and advice](https://www.reddit.com/r/rails/comments/g1l1kn/junior_dev_just_recently_got_first_job_looking/)
- url: https://www.reddit.com/r/rails/comments/g1l1kn/junior_dev_just_recently_got_first_job_looking/
---
Hello all, I just recently got my first job as a rails developer and I was hoping to get some resources (could be books, sites, whatever) or advice on things you wish you would have known about when you first started or something that helped you as a junior now. This could be for rails, ruby in general, front end stuff, just kinda whatever really. 

Just as background, before getting the job I hadn't had any prior experience with ruby or rails, except a codecademy course in 7th grade because its what RPG maker used lol. In school we used alot of Java and C++ and of course SQL. I have expierence with other languages but mainly those. I feel comfortable in HTML and I like bootstrap alot, I want to improve my CSS skills I'd say that's my weakest point. I do enjoy JS, but I wouldnt say I'm great at it or anything.
## [10][Rails 6 CSS organization](https://www.reddit.com/r/rails/comments/g14b3n/rails_6_css_organization/)
- url: https://www.reddit.com/r/rails/comments/g14b3n/rails_6_css_organization/
---
I'm finally toying around Ruby-on-Rails 6 (better later than never), JavaScript has finally evolved with webpack, which is a good thing, but I'm quite surprised to see that "The CSS Rails way" didn't changed. It's \*\*still\*\* old-fashioned CSS by default (where is SASS ? BEM ?), using the "=require" directive. This is a bit annoying when setting up new rails project with the "rails new" command. Moreover, this is not consistent with using webpacker, this makes 2 front-end builds (webpack + sprockets) running side-by-side. Finally, it do not encourage "componentization of the front-end", unless you choose a "full JS Framework" like React or Vue that you may not need anyway... So my question is : why is CSS default organization \*\*still\*\* so old-fashioned after all these years ?
## [11][My new gem: Heya is a campaign mailer for Rails](https://www.reddit.com/r/rails/comments/g1hslx/my_new_gem_heya_is_a_campaign_mailer_for_rails/)
- url: /r/ruby/comments/g1ezhp/my_new_gem_heya_is_a_campaign_mailer_for_rails/
---

## [12][How do you approach client editable page layouts with multiple sections and varying columns?](https://www.reddit.com/r/rails/comments/g181yx/how_do_you_approach_client_editable_page_layouts/)
- url: https://www.reddit.com/r/rails/comments/g181yx/how_do_you_approach_client_editable_page_layouts/
---
I fell in love with Rails years ago but for quite some time I have been busy with long-term PHP projects the last few years, but I am now looking to make a return.

A common content management requirement I face is for clients to be able to manage their own page layouts. For example, a page may have sections that are styled differently - some full width, different colours, number of columns. While classes can solve the styling challenges, column management is harder to do.

In PHP land I have used Craft, Statamic, WP (Gutenberg, ACF, Elementor and Divi) which all solve this problem in different ways with their own pros/cons.

I'm not really looking for a visual builder, more a way of assigning pre-defined section templates, e.g. two-thirds &amp; one-third, two halves, full width column. But there needs to be a good way for the user to pick which section they want, populate and re-order as needed in the admin panel.

From my initial research Locomotive CMS seems to have a nice section concept. I'm not sure how easy it would be to combine with a wider Rails app yet. I wonder if any of the admin gems like active admin or administrate are capable of doing something similar or lay some of the foundations? I'd like to avoid reinventing the wheel if possible!
