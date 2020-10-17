# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Gem/Ideas for seeing outliers in a dataset.](https://www.reddit.com/r/rails/comments/jcpo2k/gemideas_for_seeing_outliers_in_a_dataset/)
- url: https://www.reddit.com/r/rails/comments/jcpo2k/gemideas_for_seeing_outliers_in_a_dataset/
---
I know that it is simple to just "sort by value" and notice the upper end of a dataset. However, my use case is rather complicated and doesn't necessarily have to do with an integer "value" per se. For example, say I have a set of students from several locations across the U.S. and a list of grades they received in a class and I'm trying to see what were the outlying cases in the class like so:

&amp;#x200B;

|Name|Class|Location|Grade(s)|
|:-|:-|:-|:-|
|John Wick|Genocide 101|Russia|\[100,82,0,78,80,85,10\]|

I would want to average out the numbers, which would yield me something along the lines of 62.14%. However, this is in no way representative of the individual's real performance. So I would want to point out the outliers - in this specific case being the values of 100 and 0 - in order to get a more accurate representation of values. 

I would also like to use this logic in a column comparison, though not to compare the students, to see that "Oh, these two went to Russia and studied genocide? The average of that location is xyz and these performed phenomenal and these performed terribly. They're outliers."

That in mind, I do need to implement some form of standard deviation for the values, which I could do just fine.

&amp;#x200B;

Anyways, I know the maths I need and have that all mapped out, I'm just wondering if there's some sort of ML/AI or dataset analysis gem that I could use to assist in this - if one exists.

Keyword: **Defining Outliers.** 

TL;DR

Need some help popping (though from hash/model form) some outliers from a dataset a la leetcode but in actual production with more complex values.
## [4][Implementing HATEOAS in Rails API](https://www.reddit.com/r/rails/comments/jct7rc/implementing_hateoas_in_rails_api/)
- url: https://www.reddit.com/r/rails/comments/jct7rc/implementing_hateoas_in_rails_api/
---
I'm trying to learn Rails a bit more, and I'm not sure how to do this 'properly'.
From what I've seen, the routing url helpers should be able to do this, but I'm not sure if I'm misusing them or they simply can't work like this. For the record, I'm using `active_model_serializer`, and struggling with the following issues:

1) Getting the self reference URL. I could easily hardcode the link to the resource, then interpolate the `:id`. Easy, but not maintainable. I tried using a url_helper, but `todo_url(object_id)` ends up producing `"self": "http://localhost:3000/todos.17740"`

2) Getting URLs of resources not related to the object being serialized. For example, let's say in the case the client didn't send an auth token, return a link to the login/register endpoints.
## [5][Avo - Configuration-based, no-maintenance, extendable Ruby on Rails admin](https://www.reddit.com/r/rails/comments/jc85sc/avo_configurationbased_nomaintenance_extendable/)
- url: https://www.reddit.com/r/rails/comments/jc85sc/avo_configurationbased_nomaintenance_extendable/
---
Hi guys,

Today I'd like to show you [Avo](https://avohq.io), a beautiful next-generation framework that empowers you, the developer, to create fantastic admin panels for your Ruby on Rails apps with the flexibility to fit your needs as you grow.

Out of the box, it has an excellent CRUD interface, ordering, filters, and actions. It even knows how to handle your Active Record model relations.

It's super easy to configure. There's one configuration file per model and one configuration line of code per field. You can add simple fields like text, textarea, dropdowns, and more complex ones like datetime, badges, loaders, currency, and others. There's even a cool one-liner single or multi-file Active Storage integration 🤯.

**Avo's mission is to make the job of developers easier and help them and companies move faster.**

Try it in your app and let me know what you think.

Thank you,  
Adrian

[avohq.io](https://avohq.io)

[Twitter Account](https://twitter.com/avo_hq)

[Github repo](https://github.com/avo-hq/avo)

[Discord community server](https://discord.gg/pkTF6y8)
## [6][Strategies for deploying Rails engines containing webpacker-compiled assets](https://www.reddit.com/r/rails/comments/jchwst/strategies_for_deploying_rails_engines_containing/)
- url: https://www.reddit.com/r/rails/comments/jchwst/strategies_for_deploying_rails_engines_containing/
---
Hello all,

I'm working on a Rails engine that contains assets that are compiled using webpacker. I'm wondering what the correct strategy is for packaging and deploying this gem. I've read the instructions [here](https://github.com/rails/webpacker/blob/master/docs/engines.md), but I find Step 7 a bit confusing.  I tried the instructions under "Use a separate middleware", but that didn't work. So my questions are:

1. Should I be pre-compiling the assets in my gem and including the compiled assets (in the gem's public/packs) with the published gem and load those assets from there in the app? OR
2. Should I *not* include the engine's assets, but run rake my-engine:webpacker:compile in the app?

Anyone have any experience with this?
## [7][What is the right gem to ban users? Devise, Pundit or Rolify?](https://www.reddit.com/r/rails/comments/jcdqwj/what_is_the_right_gem_to_ban_users_devise_pundit/)
- url: https://www.reddit.com/r/rails/comments/jcdqwj/what_is_the_right_gem_to_ban_users_devise_pundit/
---
What is your experience?
## [8][Unique URLs For Filter Combinations](https://www.reddit.com/r/rails/comments/jcgoul/unique_urls_for_filter_combinations/)
- url: https://www.reddit.com/r/rails/comments/jcgoul/unique_urls_for_filter_combinations/
---
Hey everyone!

I'm trying to improve the SEO of the site I'm working on, and there's one method in particular that would be promising with the niche I'm in. Basically, the idea revolves around having unique pages with content for combinations of filters. 

As an example, this is something that [Nomad List](https://nomadlist.com/) implements and has success with. Let's say you select the "Cold now" and "Clean air" filters on NomadList - the page refreshes with a new URL '/cool-places-with-clean-air' and the new page has a matching H1 tag with a brief description. These custom filter pages rank well for a lot of long-tail keywords.

I'm curious if anyone has experience implementing something similar in Rails? I currently have a filtering system setup in my app, but it uses query parameters in the form of '?filter_by_param1&amp;filter_by_param2'. I'd like to migrate this to a system based on unique urls so that they have potential to rank on Google. 

I'm not really sure about how I could go about setting up the routing for something like this. The filtering itself as well as generating the content is something I can manage. 

Would really appreciate any ideas! Thanks in advance :)
## [9][Federated rails apps, how they can be made?](https://www.reddit.com/r/rails/comments/jcci9o/federated_rails_apps_how_they_can_be_made/)
- url: https://www.reddit.com/r/rails/comments/jcci9o/federated_rails_apps_how_they_can_be_made/
---
I'm a member of Mastodon. And I'm on one instance. Most of my friends are on other instances. A lot of people from different networks, such as plume, can join us on Mastodon and share their thoughts and blogs or videos from plume or peertube can be shared using mastodon. 

The ActivityPub protocol is really a mind-blowing thing to me. I love it. It helps people on different social networks to connect together (wasn't this what "thefacebook" meant to be?) and this is amazing. 

I searched a lot about the protocol itself and read some parts of mastodon's code. But never saw a tutorial or centralized document about adding this feature to a rails app. If there will be one, that'd be the most helpful document on decentralized networks in my opinion.
## [10][How do you generate the different combinations of product variants?](https://www.reddit.com/r/rails/comments/jcgq11/how_do_you_generate_the_different_combinations_of/)
- url: https://www.reddit.com/r/rails/comments/jcgq11/how_do_you_generate_the_different_combinations_of/
---
Say I have a variant model that has the variants and values like this:

color: {\[blue, red, green, etc.\]}

Size: {\[XL,L,M,S\]}

If I have a variant\_values model that belongs to variant how do I loop through those and generate variant\_value records for each possible combination?
## [11][Rails 6 with Bootstrap and Webpacker: Quick guide](https://www.reddit.com/r/rails/comments/jbv43i/rails_6_with_bootstrap_and_webpacker_quick_guide/)
- url: https://www.reddit.com/r/rails/comments/jbv43i/rails_6_with_bootstrap_and_webpacker_quick_guide/
---
I guess this exact tutorial topic is already a "mauvais ton", but after reading 10+ guides on "How to install bootstrap on Rails 6" I wrote a quick guide: [https://blog.corsego.com/2020/10/rails-6-with-bootstrap-and-webpacker.html](https://blog.corsego.com/2020/10/rails-6-with-bootstrap-and-webpacker.html)

I hope you'll find it useful!
## [12][Rails help](https://www.reddit.com/r/rails/comments/jc95ti/rails_help/)
- url: https://www.reddit.com/r/rails/comments/jc95ti/rails_help/
---
 hey guys, so i have this and i am trying to get each position(entertainment, src president) have its own page. lets say if i vote for the candidates of entertainment position it directs me to src president to vote and so on. how do i do this please? 

&amp;#x200B;

https://preview.redd.it/g0sljvp7dgt51.png?width=1366&amp;format=png&amp;auto=webp&amp;s=849b2c46173b6aa28f04e9b8279cce0aa042ebda

&gt;&lt;table&gt;  
&gt;  
&gt;  &lt;tbody&gt;  
&gt;  
&gt;&lt;%= form\_for u/election do |f| %&gt;  
&gt;  
&gt;  
&gt;  
&gt;&lt;% Position.includes(:candidates).order(:name).each do |position| %&gt;  
&gt;  
&gt;  
&gt;  
&gt;  &lt;tr&gt;  
&gt;  
&gt;&lt;th colspan="5"&gt;&lt;%= position.name %&gt;&lt;/th&gt;  
&gt;  
&gt;  &lt;/tr&gt;  
&gt;  
&gt;  &lt;% position.candidates.each do |candidate| %&gt;  
&gt;  
&gt;  &lt;tr&gt;  
&gt;  
&gt;&lt;td&gt;&lt;%= image\_tag(candidate.image, :size =&gt; '50x50') if candidate.image.attached? %&gt;&lt;/td&gt;  
&gt;  
&gt;&lt;td&gt;&lt;%= candidate.name %&gt;&lt;/td&gt;  
&gt;  
&gt;&lt;td&gt;&lt;%= candidate.info %&gt;&lt;/td&gt;  
&gt;  
&gt;&lt;td&gt;

&lt;%= f.check\_box :c\_votes %&gt;

&amp;#x200B;

&gt;&lt;/td&gt;  
&gt;  
&gt;  &lt;/tr&gt;  
&gt;  
&gt;  &lt;% end %&gt;  
&gt;  
&gt;  &lt;td&gt;  &lt;%= f.submit "Submit"  %&gt;&lt;/td&gt;  
&gt;  
&gt;&lt;% end %&gt;  
&gt;  
&gt;  &lt;/tbody&gt;  
&gt;  
&gt;&lt;/table&gt;  
&gt;  
&gt;  
&gt;  
&gt; &lt;% end %&gt;

here is the form please
