# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/f2r9mk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/f2r9mk/personal_projects_show_off_your_own_project_andor/
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
## [3][Exchanging sensitive data between two Rails apps](https://www.reddit.com/r/rails/comments/f8897p/exchanging_sensitive_data_between_two_rails_apps/)
- url: https://www.reddit.com/r/rails/comments/f8897p/exchanging_sensitive_data_between_two_rails_apps/
---
I have two Rails apps and they both use SSL/HTTPS. They have a REST api to exchange data between each other. I assume this data is encrypted because both apps are SSL/HTTPS? And if so, is this secure enough to exchange sensitive data like passwords?
## [4][Markdown redcarpet and mentions(link)](https://www.reddit.com/r/rails/comments/f88tu4/markdown_redcarpet_and_mentionslink/)
- url: https://www.reddit.com/r/rails/comments/f88tu4/markdown_redcarpet_and_mentionslink/
---
Following this tip [https://stackoverflow.com/questions/12381230/how-to-extend-redcarpet-to-support-auto-linking-user-mentions](https://stackoverflow.com/questions/12381230/how-to-extend-redcarpet-to-support-auto-linking-user-mentions)

... I edited my MarkdownRenderer in this way:

&amp;#x200B;

    class MarkdownRenderer &lt; Redcarpet::Render::HTML
    
      def preprocess(text)
        wrap_mentions(text)
      end
    
      def wrap_mentions(text)
        text.gsub! /(^|\s)(@\w+)/ do
          "#{$1}&lt;a href='/user/#{$2}'&gt;#{$2}&lt;/a&gt;"
        end
        text
      end

But then if I write u/marcus, it show me &lt;a href='/user/marcus'&gt;marcus&lt;/a&gt; **as a text and now as a link.** It blocks the html. How to solve?
## [5][Is the code for the GitHub web app on Github?](https://www.reddit.com/r/rails/comments/f88rba/is_the_code_for_the_github_web_app_on_github/)
- url: https://www.reddit.com/r/rails/comments/f88rba/is_the_code_for_the_github_web_app_on_github/
---
Hello, This might be a noob question but I"m wondering if the code for GitHub web application rails app is hosted on GitHub as an open-source project?

I'm asking this because I"m interested in how their roles and scopes are stored in the database and how it is being validated for each user. I want to draw inspiration from their implementation as I'm working on a rails app myself.
## [6][Help getting set up - can't find gem railties error](https://www.reddit.com/r/rails/comments/f82shv/help_getting_set_up_cant_find_gem_railties_error/)
- url: https://www.reddit.com/r/rails/comments/f82shv/help_getting_set_up_cant_find_gem_railties_error/
---
I've been trying to get rails set up and I keep running into problem - all of which I'm sure are just user error. I'm using Windows 10 and Ruby 2.7.

When I run rails -v, I get an error saying it can't find gem railties.

&gt;C:/Ruby27-x64/lib/ruby/2.7.0/rubygems.rb:275:in \`find\_spec\_for\_exe': can't find gem railties (&gt;= 0.a) with executable rails (Gem::GemNotFoundException)  
&gt;  
&gt;from C:/Ruby27x64/lib/ruby/2.7.0/rubygems.rb:252:in \`bin\_path  
&gt;  
&gt;from C:/RailsInstaller/Ruby2.3.3/bin/rails:22:in \`&lt;main&gt;'

Is it trying to get the wrong version of Ruby? I'm really stuck and would love some help! I ran into similar problem earlier but what I tried didn't help, I attempted to delete everything and start over, and here I am. 

I'd be happy to provide in additional files that could help!
## [7][Rails 6 Webpacker assets not compiling in production.](https://www.reddit.com/r/rails/comments/f80ffl/rails_6_webpacker_assets_not_compiling_in/)
- url: https://www.reddit.com/r/rails/comments/f80ffl/rails_6_webpacker_assets_not_compiling_in/
---
I have a website live in production that was working 10 minutes ago. Then I tried deploying a tiny little update (it was only in an erb file) and now the website is running without css or javascript.

Usually this would happen on every deployment. However, I would work around this by ssh-ing into my server and running `bin/webpack` and that would compile all the assets and everything would look great again. However, this time, it took forever to compile and I got this message:

&amp;#x200B;

WARNING in asset size limit: The following asset(s) exceed the recommended size limit (244 KiB).

This can impact web performance.

&amp;#x200B;

What's weird is that I never changed any javasript or css when I made the update. All i changed was html.

This website is hosted on a digital ocean droplet. I have multiple sites on this one droplet. And, i've noticed that i've started getting this same problem for the other sites as well. However, the website i'm working on right now is super important and I need to get it working again soon.

I appreciate any help

&amp;#x200B;

I'm temporarily compiling the packs manually and deploying that. It's working now, but it's a horrible work around. 
## [8][How to step up?](https://www.reddit.com/r/rails/comments/f80e2j/how_to_step_up/)
- url: https://www.reddit.com/r/rails/comments/f80e2j/how_to_step_up/
---
Hello,  
i have been working with ROR since the start of my career but i have no idea now how can i step up.  
I work with ROR in my daily job  60 % and 40 % is for front end stuff like vue.  
How can i step up my level?  
Thanks.
## [9][[Help] Performance/Improvements on model](https://www.reddit.com/r/rails/comments/f7xe78/help_performanceimprovements_on_model/)
- url: https://www.reddit.com/r/rails/comments/f7xe78/help_performanceimprovements_on_model/
---
I have a model, model\_A, and it has a *belongs\_to* association with another model - model\_B. Model\_B represents the possibles types of model\_A (e.g cars &amp; brands). Because model\_B is not suppose to change, I was think if is worth to reduce database costs of searching and joining Model\_B and Model\_A. I was thinking in something like that:

    # model_A
    
    belongs_to :model_b
    
    before/after_save :update_model_b_name
    
    def update_model_b_name
      if model_b.present? &amp;&amp; model_b_changed?
        self.model_b_name = self.model_b.name
      end 
    end
    
    validates :model_b_id, presence: true
    validates :model_b_name, presence: true # may create conflict here with new entity

I'm considering that because usually my problem query a list of model\_A, then I have the cost of join with model\_B. Using this way may lost a little of consistence between the models - e.g let say model\_B is updated, but the update won't reflects direct on model\_A -, but that wouldn't be a problem. Also, as I already said, model\_B shouldn't be edit frequently (almost never), just added new values.

What do you guys think about it? Do you guys think it's worthy to do something like that in therms of database cost? Is it a common practice?

\#Another Solution

Another possibility for this problem would be create an array attribute on model\_A to holds model\_B's data, hence excluding model\_B. I never did that, so a need some opinion too. Right now, I have model\_B just for the convince of editing &amp; creating new data, instead of consistence or searching purposes.

&amp;#x200B;
## [10][Approach to authorisation when using GraphQL relay node interface](https://www.reddit.com/r/rails/comments/f7vdrr/approach_to_authorisation_when_using_graphql/)
- url: https://www.reddit.com/r/rails/comments/f7vdrr/approach_to_authorisation_when_using_graphql/
---
As the title says, I’m wondering what’s the best approach to authenticated GraphQL queries on the “node” field, making sure one user or a malicious client can’t access the nodes that belong to a user.

I was thinking of creating the node field resolved myself and calling Pundit.authorize depending on the class type of the node, in like some massive switch statement.

I feel like there’s a better, cleaner more idiomatic approach but I can’t think of one so if you have one to share thag would be great
## [11][Trix Editor Extension](https://www.reddit.com/r/rails/comments/f7t1xc/trix_editor_extension/)
- url: https://www.reddit.com/r/rails/comments/f7t1xc/trix_editor_extension/
---
I’m wanting to add a checkbox-style input to the Trix editor / action text. Anyone have any ideas on how to implement this?
## [12][How can I use &amp; instead of try! in this case?](https://www.reddit.com/r/rails/comments/f7ngfi/how_can_i_use_instead_of_try_in_this_case/)
- url: https://www.reddit.com/r/rails/comments/f7ngfi/how_can_i_use_instead_of_try_in_this_case/
---
In a the view I show the first 120 letters of the description of some posts, but I have to safe navigate these since `post.description` can be `nil`, for now I'm using `post.description.try!(:[], 0..120)`.

Is there a way to do the same with `&amp;` syntax?
