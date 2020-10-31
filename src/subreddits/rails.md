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
## [2][How to reference new active storage file after save in model?](https://www.reddit.com/r/rails/comments/jlbpt1/how_to_reference_new_active_storage_file_after/)
- url: https://www.reddit.com/r/rails/comments/jlbpt1/how_to_reference_new_active_storage_file_after/
---
I am trying to read meta information from mp3s uploaded using taglib-ruby gem. It doesn't seem to be finding the file to read in though. My active storage is set up to store in public/uploads. How would I reference the recently uploaded file in this model?

    class Song &lt; ApplicationRecord
      belongs_to :band
      has_one_attached :file, dependent: :destroy
    
      after_save :set_id3_tags_in_database
    
      require 'taglib'
    
      def set_id3_tags_in_database
      TagLib::MPEG::File.open(file.filename) do |file|
         tag = file.id3v2_tag
         tag.title
         tag.album
         tag.artist
       end
      end
    end

Not even sure if I should put that in the model or in the controller either.

&amp;#x200B;

Any thoughts are appreciated.
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
## [4][New article: saving address as a hash with serialization](https://www.reddit.com/r/rails/comments/jl4eco/new_article_saving_address_as_a_hash_with/)
- url: https://www.reddit.com/r/rails/comments/jl4eco/new_article_saving_address_as_a_hash_with/
---
I don't like having separate columns for `city`, `state`, `street_address_1` etc...So here's how I store them as a hash and *edit them in a form*: [https://blog.corsego.com/2020/10/ruby-on-rails-serialization-saving.html](https://blog.corsego.com/2020/10/ruby-on-rails-serialization-saving.html) **But... do you think this approach makes sence?**
## [5][devise_ldap_authenticatable API Based](https://www.reddit.com/r/rails/comments/jl00hg/devise_ldap_authenticatable_api_based/)
- url: https://www.reddit.com/r/rails/comments/jl00hg/devise_ldap_authenticatable_api_based/
---
Does anyone know how to use devise\_ldap\_authenticatable as a bare bones API?
## [6][Clueless about how to fetch data from an API gem](https://www.reddit.com/r/rails/comments/jkt9ab/clueless_about_how_to_fetch_data_from_an_api_gem/)
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
## [7][Unused methods slow down rspec?](https://www.reddit.com/r/rails/comments/jknjzu/unused_methods_slow_down_rspec/)
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
## [8][[Help] Using Hstore in Form](https://www.reddit.com/r/rails/comments/jknjx7/help_using_hstore_in_form/)
- url: https://www.reddit.com/r/rails/comments/jknjx7/help_using_hstore_in_form/
---
I have a **emails** field in my **contacts** table which is a hstore: `t.hstore :emails` . My goal is to store key value pairs of label and email. For example if a use has 2 emails:

    { "Home" =&gt; "email@example.com", "Work" =&gt; "work@example.com" }

How do I go about making a form where a user can select the label from a select dropdown and enter his email in a text field? Also, would it be possible for the user to add more emails if wanted?

I would also need to permit the emails param in the controller which is giving me issues.

Thanks
## [9][Read this medium post about "devise-jwt" and some question popped in my head.](https://www.reddit.com/r/rails/comments/jkm7gt/read_this_medium_post_about_devisejwt_and_some/)
- url: https://www.reddit.com/r/rails/comments/jkm7gt/read_this_medium_post_about_devisejwt_and_some/
---
I was searching for a good API authentication way in rails and I found this medium article: 

[https://medium.com/@nandhae/2019-how-i-set-up-authentication-with-jwt-in-just-a-few-lines-of-code-with-rails-5-api-devise-9db7d3cee2c0](https://medium.com/@nandhae/2019-how-i-set-up-authentication-with-jwt-in-just-a-few-lines-of-code-with-rails-5-api-devise-9db7d3cee2c0) 

It really explained the implementation phase very well, but I have a question, how can we make users? how can we authenticate users? 

These are unsolved for me. 

P.S : Haven't implemented what I read yet. I may do it in a better time (It's almost 4 A.M in my zone! and I'm too sleepy!)
## [10][Audited vs papertrail gem for auditing table changes?](https://www.reddit.com/r/rails/comments/jk9i0y/audited_vs_papertrail_gem_for_auditing_table/)
- url: https://www.reddit.com/r/rails/comments/jk9i0y/audited_vs_papertrail_gem_for_auditing_table/
---
I'm currently looking to use either one of these two for auditing changes made by specific user roles on a rails 6 ecommerce app using mysql. They both seem similar and I'm currently leaning towards papertrail due to the more frequent gem updates, but I was curious if there is any specific advantage of using one over the other(performance, ease of use...etc) or anyone knows of any pros/cons that stick out.

Would be keen to know if people have used any other gems that seemed good as well.
## [11][How can I make this fetch?](https://www.reddit.com/r/rails/comments/jkbozb/how_can_i_make_this_fetch/)
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
