# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/
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
## [2][Clueless about how to fetch data from an API gem](https://www.reddit.com/r/rails/comments/jkt9ab/clueless_about_how_to_fetch_data_from_an_api_gem/)
- url: https://www.reddit.com/r/rails/comments/jkt9ab/clueless_about_how_to_fetch_data_from_an_api_gem/
---
Hi, I want to start utilizing APIs in my Rails apps but I am completely lost on how to actually get data from APIs when using a gem.

I am trying to make a simple app where you search for a query and it brings up a list of items with the matching keyword. Here are some API gems I have been trying to work with:

Giant Bomb: [https://github.com/games-directory/api-giantbomb](https://github.com/games-directory/api-giantbomb)

Steam: [https://github.com/Olgagr/steam-web-api](https://github.com/Olgagr/steam-web-api)

PokeAPI: [https://github.com/rdavid1099/poke-api-v2](https://github.com/rdavid1099/poke-api-v2)

I'm mostly focusing on the Giant Bomb API. I've added the gem and initialized my API key, but after that I'm pretty clueless in what to do. I was thinking I need to make a search method in my Games controller like so:

  def search

@ games = Game.find\_by(params\[:id\])

search = GiantBomb::Search.new().query(:query)

  end

I'm not sure what to do beyond this or even if it's right (my search is not working so I assume the latter).

Any tips or guides? I've googled for API guides but most seem to be fetching data via an API key and now on how to use an API gem.
## [3][What are your dislikes about the "rails new" command ?](https://www.reddit.com/r/rails/comments/jkwpxr/what_are_your_dislikes_about_the_rails_new_command/)
- url: https://www.reddit.com/r/rails/comments/jkwpxr/what_are_your_dislikes_about_the_rails_new_command/
---
The "rails new" command is very handy when you have to test something in complete isolation. Here are my complains about that command :

\* Too much things included. It would have been better to start with raw defaults and then only include  actual needs.

\* To the contrary, very basic things (or things that I consider very basic) are missing : env var, code coverage, http mocking... I know there are gems for this, but I find it weird that ActionText is a default, but env var not.

\* Webpacker seems unfinished feature so far. Webpack manage JS, but Sprockets others assets.

\* I miss a default HomeController or WelcomeController. Same can be said for the view.

\* .gitignore is not bad but still miss a few lines

\* No auto-refresh of the browser when coding the view (or backend) part.

\* sqlite by default. I don't know a Rails coders who uses it.

\* sassc is too slow to install. Defaut compilation (without tweaking) of webpacker assets are also too slow.

Bottom line : a **very big thank you** to Rails contributors, too much is better than not enough, and there are good ways to start with better defaults that match your needs anyway. Just wanted to share my thoughts with the community. Do you have also things you dislike about the Rails default tools ?
## [4][Unused methods slow down rspec?](https://www.reddit.com/r/rails/comments/jknjzu/unused_methods_slow_down_rspec/)
- url: https://www.reddit.com/r/rails/comments/jknjzu/unused_methods_slow_down_rspec/
---
I've been working on upgrading an old large Rails 4.2 project to 5.2 and hit this debugging some stuff... if I leave the last two methods (descendants/debug) uncommented, rspec tests take 3-4 times longer than if these methods are commented out. The methods aren't called anywhere (I only use this in console) ... any idea why that is? I wonder if I'm doing this unintentionally in other places.  

ApplicationRecord is inherited by all models (part of Rails 5 upgrade)

```
class ApplicationRecord &lt; ActiveRecord::Base
  self.abstract_class = true

  def self.descendants
    ObjectSpace.each_object(Class).select { |klass| klass &lt; self }
  end

  def self.debug
    Rails.logger.level = 3
    descendants.sort_by(&amp;:name).each_with_index do |obj, idx|
      puts "#{idx}-#{obj.name}: #{obj.count}"
    end
    nil
  end
end
```
## [5][[Help] Using Hstore in Form](https://www.reddit.com/r/rails/comments/jknjx7/help_using_hstore_in_form/)
- url: https://www.reddit.com/r/rails/comments/jknjx7/help_using_hstore_in_form/
---
I have a **emails** field in my **contacts** table which is a hstore: `t.hstore :emails` . My goal is to store key value pairs of label and email. For example if a use has 2 emails:

    { "Home" =&gt; "email@example.com", "Work" =&gt; "work@example.com" }

How do I go about making a form where a user can select the label from a select dropdown and enter his email in a text field? Also, would it be possible for the user to add more emails if wanted?

I would also need to permit the emails param in the controller which is giving me issues.

Thanks
## [6][Read this medium post about "devise-jwt" and some question popped in my head.](https://www.reddit.com/r/rails/comments/jkm7gt/read_this_medium_post_about_devisejwt_and_some/)
- url: https://www.reddit.com/r/rails/comments/jkm7gt/read_this_medium_post_about_devisejwt_and_some/
---
I was searching for a good API authentication way in rails and I found this medium article: 

[https://medium.com/@nandhae/2019-how-i-set-up-authentication-with-jwt-in-just-a-few-lines-of-code-with-rails-5-api-devise-9db7d3cee2c0](https://medium.com/@nandhae/2019-how-i-set-up-authentication-with-jwt-in-just-a-few-lines-of-code-with-rails-5-api-devise-9db7d3cee2c0) 

It really explained the implementation phase very well, but I have a question, how can we make users? how can we authenticate users? 

These are unsolved for me. 

P.S : Haven't implemented what I read yet. I may do it in a better time (It's almost 4 A.M in my zone! and I'm too sleepy!)
## [7][Audited vs papertrail gem for auditing table changes?](https://www.reddit.com/r/rails/comments/jk9i0y/audited_vs_papertrail_gem_for_auditing_table/)
- url: https://www.reddit.com/r/rails/comments/jk9i0y/audited_vs_papertrail_gem_for_auditing_table/
---
I'm currently looking to use either one of these two for auditing changes made by specific user roles on a rails 6 ecommerce app using mysql. They both seem similar and I'm currently leaning towards papertrail due to the more frequent gem updates, but I was curious if there is any specific advantage of using one over the other(performance, ease of use...etc) or anyone knows of any pros/cons that stick out.

Would be keen to know if people have used any other gems that seemed good as well.
## [8][How can I make this fetch?](https://www.reddit.com/r/rails/comments/jkbozb/how_can_i_make_this_fetch/)
- url: https://www.reddit.com/r/rails/comments/jkbozb/how_can_i_make_this_fetch/
---
Hey Rails stars,

Am working with ActiveAdmin and I have a custom page am calling "Approvals" with an Arbe table that's supposed to display data from a join of 2 tables. My code is as follows:

**AA custom page:**

    ActiveAdmin.register_page 'Approvals' do
        content do render partial: 'approval_requests', locals: { request_items:         ApprovalRequest.joins(:approval_items) } 
    end 
    end

&amp;#x200B;

**..the partial**

    &lt;%=
    Arbre::Context.new({}, self) do 
    table_for(request_items, sortable: true, class: 'index_table') do 
        column :requesting_party 
        column :approving_party 
        column 'Resource Type' do |ar| 
            ar.resource_type 
        end 
        end 
    end 
    %&gt;

&amp;#x200B;

**ApprovalRequest**

    class ApprovalRequest &lt; ApplicationRecord
        #fields: :requesting_party, :approving_party...
        has_many :approval_items 
    end

&amp;#x200B;

**ApprovalItem**

    class ApprovalItem &lt; ApplicationRecord
        #fields: :request_id, :resource_type....
        belongs_to :approval_request, foreign_key: :request_id 
    end

&amp;#x200B;

With that, am getting the error "PG::UndefinedColumn: ERROR:  column approval\_items.approval\_request\_id does not exist.."

So where could I be missing it? Thanks.
## [9][Designed a simple rails app that communicates with Google Drive API.](https://www.reddit.com/r/rails/comments/jkb9uk/designed_a_simple_rails_app_that_communicates/)
- url: https://www.reddit.com/r/rails/comments/jkb9uk/designed_a_simple_rails_app_that_communicates/
---
[https://github.com/AbhishekSinhaCoder/Ruby-on-Rails-Challenge](https://github.com/AbhishekSinhaCoder/Ruby-on-Rails-Challenge)
## [10][How is github able to provide go to definition and find references while I can't do it in VS Code](https://www.reddit.com/r/rails/comments/jjpqnw/how_is_github_able_to_provide_go_to_definition/)
- url: https://www.reddit.com/r/rails/comments/jjpqnw/how_is_github_able_to_provide_go_to_definition/
---
In VS code I struggle to get something like below to get to work. I have installed solargraph but it takes a long time to find something and most of the time it doesn't work?   


Any idea on how this works in github and if it's possible to get this to work in vs code? 

https://preview.redd.it/iuy85bnvruv51.png?width=3096&amp;format=png&amp;auto=webp&amp;s=628207d307a41c12920cf466b5d76b483f520774
## [11][A Gem for Cloudimage Image Optimizer](https://www.reddit.com/r/rails/comments/jjzf5z/a_gem_for_cloudimage_image_optimizer/)
- url: https://www.reddit.com/r/rails/comments/jjzf5z/a_gem_for_cloudimage_image_optimizer/
---
**Hi! Our tech team at Scaleflex just published a gem for our image optimizer Cloudimage** ([https://www.cloudimage.io/](https://www.cloudimage.io/))-&gt; Discover it here: Cloudimage-gem ([https://rubygems.org/gems/cloudimage](https://rubygems.org/gems/cloudimage))

What you can achieve with this gem:

* supports Ruby 2.4 and above, JRuby, and TruffleRuby
* on-the-fly resizing of all your images (150+ transformations)
* image optimization and WebP/AVIF (tba) compression
* responsive design, progressive/lazy loading
* worldwide media delivery via CDN

Happy to get your feedback!

Greetings!
