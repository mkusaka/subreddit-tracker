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
## [2][How to use AJAX on an devise protected route?](https://www.reddit.com/r/rails/comments/jlzduy/how_to_use_ajax_on_an_devise_protected_route/)
- url: https://www.reddit.com/r/rails/comments/jlzduy/how_to_use_ajax_on_an_devise_protected_route/
---
I want to make some post requests to my app. I’m not using any front end framework. Just straight up rails and erb (well, i guess this isn’t entirely true. I am using stimulus js).

In stimulus, I’m trying to make ajax requests which is working great. I know that I have to send an authenticity token along with my request. I just grab that from the meta tags. However, I don’t know what to send so that devise can authenticate my request with the current user. 

Do I have to send something in the header of the request? Or maybe some other additional parameters? 

```javascript
$.ajax({
    type: 'POST',
    url: url,
    headers: {
        "X-CSRF-Token": CSRF_TOKEN,
    }
}).done(function(data) { 
    alert(data);
});

```

Any help is appreciated. Thanks!
## [3][.distinct and postgres point type failure](https://www.reddit.com/r/rails/comments/jlynvx/distinct_and_postgres_point_type_failure/)
- url: https://www.reddit.com/r/rails/comments/jlynvx/distinct_and_postgres_point_type_failure/
---
I am using point type for keeping some coordinates in Rails. All the queries I've been using that have a `.distinct` fail because there are no operators in postgres for equality. 

What If I use `.select('DISTINCT ON(...` on the `.id` instead? What are any other alternatives?
## [4][Bookmarks on Ruby or Rails](https://www.reddit.com/r/rails/comments/jlt3bo/bookmarks_on_ruby_or_rails/)
- url: https://www.reddit.com/r/rails/comments/jlt3bo/bookmarks_on_ruby_or_rails/
---
Hi everyone, I want to make a bookmarks system like on [Myanimelist.net](https://Myanimelist.net), here is a screenshot of what I am talking about [https://prnt.sc/vayxio](https://prnt.sc/vayxio)

Do you know if there is already a gem, plugin or module that has most of functionality covered?
## [5][Simple but Useful Rails Engines](https://www.reddit.com/r/rails/comments/jlkgx6/simple_but_useful_rails_engines/)
- url: https://www.reddit.com/r/rails/comments/jlkgx6/simple_but_useful_rails_engines/
---
I have a series of simple and useful plugins without complex logic to checkout. Wouldn't mind on getting some collaboration and suggestions from other developers.

[https://github.com/phcdevworks](https://github.com/phcdevworks)

[https://rubygems.org/profiles/phcdevworks](https://rubygems.org/profiles/phcdevworks)
## [6][How can I send truly empty body when hitting a API destroy endpoint?](https://www.reddit.com/r/rails/comments/jlyjq5/how_can_i_send_truly_empty_body_when_hitting_a/)
- url: https://www.reddit.com/r/rails/comments/jlyjq5/how_can_i_send_truly_empty_body_when_hitting_a/
---
I am using Rails 5.

I want to send an empty body after I hit the delete end point, but right now probably by rails default code, sending the \`id\` of the deleted resource. But i want it to be empty. What code I need to change here:

    def destroy
      @project.destroy
    
      respond_to do |format|
        format.json { head :no_content }
      end
    end

See the screenshot:

&amp;#x200B;

https://preview.redd.it/17wfwg190lw51.png?width=1844&amp;format=png&amp;auto=webp&amp;s=51c15fad17fb19356e2f7d1e321d681cd4bc20d7
## [7][DOM Parsing of XML file how to approach ?](https://www.reddit.com/r/rails/comments/jlovhj/dom_parsing_of_xml_file_how_to_approach/)
- url: https://www.reddit.com/r/rails/comments/jlovhj/dom_parsing_of_xml_file_how_to_approach/
---
Hi There Experts, 

some weeks ago I asked help from community regarding processing XML  files on RUBY with this post :

[https://www.reddit.com/r/ruby/comments/j6ht5w/is\_there\_a\_gem\_for\_xml\_to\_json\_convertion/](https://www.reddit.com/r/ruby/comments/j6ht5w/is_there_a_gem_for_xml_to_json_convertion/)

After read a lot I discover that what I need is parse a XML file and the DOM Parsing method seems to me more appropriate ( as I still do not have correct understanding of SAX Parsing) ... 

I have checked Rexml but as is not active library with a lot of open issues I am checking for alternatives to accomplish this task ... Rigth now I checking xml\_to\_hash but decided to ask to RoR community if also have a suggestion ... as my requirements is more simplicity than performance as i will not need to work with HUGE files ...

&amp;#x200B;

thanks in advance,
## [8][How to reference new active storage file after save in model?](https://www.reddit.com/r/rails/comments/jlbpt1/how_to_reference_new_active_storage_file_after/)
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
## [9][What are your dislikes about the "rails new" command ?](https://www.reddit.com/r/rails/comments/jkwpxr/what_are_your_dislikes_about_the_rails_new_command/)
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
## [10][New article: saving address as a hash with serialization](https://www.reddit.com/r/rails/comments/jl4eco/new_article_saving_address_as_a_hash_with/)
- url: https://www.reddit.com/r/rails/comments/jl4eco/new_article_saving_address_as_a_hash_with/
---
I don't like having separate columns for `city`, `state`, `street_address_1` etc...So here's how I store them as a hash and *edit them in a form*: [https://blog.corsego.com/2020/10/ruby-on-rails-serialization-saving.html](https://blog.corsego.com/2020/10/ruby-on-rails-serialization-saving.html) **But... do you think this approach makes sence?**
## [11][devise_ldap_authenticatable API Based](https://www.reddit.com/r/rails/comments/jl00hg/devise_ldap_authenticatable_api_based/)
- url: https://www.reddit.com/r/rails/comments/jl00hg/devise_ldap_authenticatable_api_based/
---
Does anyone know how to use devise\_ldap\_authenticatable as a bare bones API?
