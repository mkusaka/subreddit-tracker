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
## [3][Is @ option or not in link_to?](https://www.reddit.com/r/rails/comments/g0hgo8/is_option_or_not_in_link_to/)
- url: https://www.reddit.com/r/rails/comments/g0hgo8/is_option_or_not_in_link_to/
---
I am a little confused on link\_to method. Let's say you have link\_to 'article', article\_path(@article). Do I have to use the @? Or is the @ only used in earlier Ruby versions?
## [4][Need help with Rails deployment](https://www.reddit.com/r/rails/comments/g0dlcg/need_help_with_rails_deployment/)
- url: https://www.reddit.com/r/rails/comments/g0dlcg/need_help_with_rails_deployment/
---
I am trying to deploy a small rails app to my VPS on DigitalOcean.

This is purely for learning more about devops and deployment. I have relied on Heroku for too long and I want to understand how it works underneath all the abstraction. 

Here's what I am trying to accomplish - 

    load balancer (NGNIX)
                |
                |
    ----------------------
    |                      |     
    |                      |
    Node 1             Node 2
           |            |
      Postgres DB server

It took me a couple of days but I have now setup everything (creating each server, installing nginx, ruby, rails, letsencrypt, postgres manually). Everything seems to be working.

I was wondering if there is a way to automate this. 
Let's say I want to add another node, is there a way to run a command and have it do all this? 

It seems to me that this is going to be a pain to maintain (e.g: update ruby version on each node etc). Is there a better way to do this? Most of the online tutorials seem to be pointing towards Docker. Is that the preferred way to deploy rails app nowadays? 

I don't have a lot of experience with docker. I only fiddled with Docker and Docker Compose a little bit and I was able to run my rails app locally. How can I push the docker-compose file to my server?

Lastly, is there a book or course that can help me understand more about devops and rails deployment? 

Thanks!
## [5][A faster version for different "if between"](https://www.reddit.com/r/rails/comments/g0gaub/a_faster_version_for_different_if_between/)
- url: https://www.reddit.com/r/rails/comments/g0gaub/a_faster_version_for_different_if_between/
---
Hi guys, I'm new on rails.

I'm adding a level system with different user titles. I used this system (*this is a short version. The levels are over 100*) :

    &lt;% if user.level &lt; 5 %&gt;
    	Beginner
    &lt;% elsif user.level.between? 5, 9 %&gt;
    	Recruit
    &lt;% elsif user.level.between? 10, 14 %&gt;
    	Enthusiast
    &lt;% elsif user.level.between? 15, 19 %&gt;
    	Apprentice
    &lt;% elsif user.level.between? 20, 24 %&gt;
    	Graduate
    &lt;% elsif user.level &gt; 25 %&gt;
    	The Best
    &lt;% end %&gt;

**It's work, but... is it the best/faster?**

**Can I improve it?**
## [6][Why does turbolinks redirect only work for non-GET xhr requests?](https://www.reddit.com/r/rails/comments/g0edjs/why_does_turbolinks_redirect_only_work_for_nonget/)
- url: https://www.reddit.com/r/rails/comments/g0edjs/why_does_turbolinks_redirect_only_work_for_nonget/
---
If `redirect_to` is used during an xhr POST/PUT/PATCH/DELETE request, turbolinks will perform a `Turbolinks.visit()` and therefore the redirect works correctly.

However this does not work if the original request was GET.

Diving into the turbolinks codebase, I can see this behaviour is clearly intentional:

```
if turbolinks != false &amp;&amp; request.xhr? &amp;&amp; !request.get?
  visit_location_with_turbolinks(location, turbolinks)
```

This piece of code is found inside the `redirect_to` method here: https://github.com/turbolinks/turbolinks-rails/blob/master/lib/turbolinks/redirection.rb

Can anyone shed some light on why this behaviour is limited to non-GET requests?

Thanks
## [7][Getting basic Rails testing terminology down](https://www.reddit.com/r/rails/comments/g0a4qn/getting_basic_rails_testing_terminology_down/)
- url: https://www.reddit.com/r/rails/comments/g0a4qn/getting_basic_rails_testing_terminology_down/
---
I just wanted to make sure I have a basic mental model with working terminology when it comes to Rails, this is very simple but If feel it important to articulate:

"Fixtures" (singular "Fixture") are simply test/sample data, and come built-in with Rails.

A common replacement for fixtures that come with Rails, are "factories"(singular "factory"), via the gem FactoryBot.

If we need to work with large sets of data, we may want to have those factories be dynamic and varied, we can do so using the Faker gem to - for example - create a bunch of User test objects with varied names, emails, and phone numbers.

Test data(factories) can be simply used in memory( via `FactoryBot.build()`) or stored/saved to the test DB (via `FactoryBot.create()`), the latter is needed for when a test is testing data retrieval from the DB.  

---------
Any enhancements to the above are appreciated, but would be good to get feedback, thanks!

Also, why do people prefer FactoryBot over fixtures, the former just reads better and is more versatile?
## [8][[ANN 🚀 ] Ralix - A JavaScript micro-framework for building and organizing Rails front-end code](https://www.reddit.com/r/rails/comments/fzz5i0/ann_ralix_a_javascript_microframework_for/)
- url: https://www.reddit.com/r/rails/comments/fzz5i0/ann_ralix_a_javascript_microframework_for/
---
Ralix is a JavaScript microframework with the goal to provide barebones and utilities to help enhance your current Rails (server-side rendered) views. It pairs really well with Turbolinks.

Following the spirit of frameworks like Stimulus, Ralix doesn't seek to take over your entire front-end logic and rendering. Instead, it's designed to enhance your HTML ✨

It has no dependencies, a really small codebase (\~350 LOC) and it's actually used in production for around 1 year 🚀

\- GitHub: [https://github.com/ralixjs/ralix](https://github.com/ralixjs/ralix)  
\- Example app (js app): [https://github.com/ralixjs/ralix-todomvc/tree/master/app/javascript](https://github.com/ralixjs/ralix-todomvc/tree/master/app/javascript)  
\- Live version: [https://ralix-todomvc.herokuapp.com/](https://ralix-todomvc.herokuapp.com/)
## [9][From selected to disabled (input) - FrontEnd](https://www.reddit.com/r/rails/comments/g0gicz/from_selected_to_disabled_input_frontend/)
- url: https://www.reddit.com/r/rails/comments/g0gicz/from_selected_to_disabled_input_frontend/
---
I have a page with this

        &lt;%= render color, selected: (@color == color) %&gt;

and the color page with this

        &lt;%= button_to color.name, color_path, selected: selected %&gt;

in this way I have the "selected color input" like this

        &lt;form class="button_to" etc. etc.&gt;&lt;input selected="selected" type="submit" value="Red"&gt; etc. etc.

**I have to relace the selected="selected" only with disabled. How to do?**
## [10][Ruby on rails vs "Enterprise language and frameworks": some suggestions needed](https://www.reddit.com/r/rails/comments/fzx2uv/ruby_on_rails_vs_enterprise_language_and/)
- url: https://www.reddit.com/r/rails/comments/fzx2uv/ruby_on_rails_vs_enterprise_language_and/
---
Hi community,

I am a professional developer that at work uses "traditional and boring enterprise languages and frameworks" ( started in J2EE and now it's .NET core  since the last two years). Now for a personal project I am considering whether to adopt a framework based on more dynamic languages and my eye got caught on Ruby on Rails. What are the advantages I might get with RoR with respect to MVC .net core? I am interested to get informed opinions only in the aspects that affects speed/ease of development. For example the build-release-cycle in .NET core is still a little bit slow even when using Azure pipelines and editing code on the fly is limited only to changes within a method and when in debug mode. Also profiling properly a web application is a little bit complicated and requires external tools or major versions of VS ( with the corresponding license costs). The obvious advantage of .NET core is that the ecosystem is impressive and performances really rock. Possibly I would like to hear an answer with somebody that has some experience with both frameworks ( classic .NET instead of .NET core would do as well ). Thanks for the support.
## [11][Seeking Rails Developers to help with an open source project.](https://www.reddit.com/r/rails/comments/g07jw5/seeking_rails_developers_to_help_with_an_open/)
- url: https://www.reddit.com/r/rails/comments/g07jw5/seeking_rails_developers_to_help_with_an_open/
---
Hey fellow Rails-devs!

**About me:** I am a technical wizard with nearly 20 years development experience, everything from Perl to Java to Rails. I love Rails because it's so easy to prototype something completely from zero. I build everything on a stack of Ubuntu, NGinx, Puma, HAML, right now.

**The project:** I have a project that I released as open source, and I'd be interested in gathering collaborators. It helps people who can't attend anime tradeshows (due to COVID-19) make a virtual store. We don't charge people nor vendors to use it - it's currently just a "Buy Now" button on Paypal.

I'd love support implementing other features from a good Rails developer - plus it's a great way to show public contributions on your Github profile (if you care about that sort of thing... some recruiters check your profile).

Project url: [https://github.com/ryankopf/cononline](https://github.com/ryankopf/cononline)

I've got a small budget of $500 for the right dev(s) to help a bit.

Anyone interested?

Ryan
## [12][Help with major refactor and data structures](https://www.reddit.com/r/rails/comments/fzvlk1/help_with_major_refactor_and_data_structures/)
- url: https://www.reddit.com/r/rails/comments/fzvlk1/help_with_major_refactor_and_data_structures/
---
Hi guys,  


I have to do a major refactor in a simple app that I own, and don't know well how.

So far, I have been using user with roles (I have several: business owners, editors, admins..)  
Thing is, every role needs custom attributes  very different from each others. In the case of business owners it is so specific that it needs its own associations (that also have to change from has\_one to has\_many)  


So far I am in a mess. What is the best way to order this kind of structures? (considering that has to work with the previous data)
