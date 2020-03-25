# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/fgx7fz/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/fgx7fz/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/)
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
## [3][Scintilla is a gem for one-time scripts on Rails apps](https://www.reddit.com/r/rails/comments/fonh4l/scintilla_is_a_gem_for_onetime_scripts_on_rails/)
- url: https://www.reddit.com/r/rails/comments/fonh4l/scintilla_is_a_gem_for_onetime_scripts_on_rails/
---
Hi all!

I just extracted a gem from a couple projects I'm running in production. It allows running scripts only once, exactly as you do with migrations.

Here's the link: [https://github.com/metalelf0/scintilla](https://github.com/metalelf0/scintilla)

My typical use case is this:

\- I realize I need to do some operations involving production data, and I want to keep track of them;

\- I create a \`Scintilla\` script with \`rails g:scintilla \[my\_script\_name\]\`, and put my ruby code in the \`do\_the\_stuff\` method;

\- I commit and push the code to staging env; an hook in my deployment flow runs \`rake scintilla:run\`, that finds the new script and runs it;

\- I check it did what intended and push the code to production.

It has some advantages over other approaches:

\- if you're currently using \`seeds.rb\` for every data operation, \`Scintilla\` will allow to have a better history of scripts you have ran; also scripts no longer need to be idempotent - they are ran only once;

\- if you're using migrations, \`Scintilla\` will allow to better separate concerns - you will be able to use migrations for db related operations and \`Scintilla\` for everything else. Also, you can safely use rails models in \`Scintilla\`: each script is \`require\`d conditionally, so if it has already been ran, it won't be required anymore.

What do you think about it? Feel free to post any suggestion, criticism and anything you find useful!
## [4][ActiveStorage DB relation explainer please](https://www.reddit.com/r/rails/comments/foop3u/activestorage_db_relation_explainer_please/)
- url: https://www.reddit.com/r/rails/comments/foop3u/activestorage_db_relation_explainer_please/
---
Hi there. I’m migrating from Paperclip to AS and so far doing fine but got stuck converting a gem that uses multiple images per record. 

It’s a has_many_attached relation, with an Offer model having many OfferImages. 

Should I have one Attachment per record, then multiple Blobs per Attachment? Or one Attachment per Blob, meaning many Blobs and Attachments per record? 

I can’t find this info anywhere, but I suspect that because each Attachment has a blob_id there’s actually many Attachment and many Blobs for each Offer. 

Please enlighten me as to how the Attachment and Blob tables should look.
## [5][clockworkd on aws](https://www.reddit.com/r/rails/comments/fohlp4/clockworkd_on_aws/)
- url: https://www.reddit.com/r/rails/comments/fohlp4/clockworkd_on_aws/
---
Has anyone ever tried to daemonize clockwork process on an aws environment? For some reason it doesn't save on the processes.

OS is in ubuntu 
## [6][Trying to hit a third party API from a controller action](https://www.reddit.com/r/rails/comments/folhcn/trying_to_hit_a_third_party_api_from_a_controller/)
- url: https://www.reddit.com/r/rails/comments/folhcn/trying_to_hit_a_third_party_api_from_a_controller/
---
I am using a rails application to fetch some data from a third party API and want to dedicate a controller action specifically for one action, how can I possibly do something like this?  
For example:- I want to list a blog post from WordPress API and link it with blogs action in my post controller.  
any help?
## [7][Rails 6 API template](https://www.reddit.com/r/rails/comments/fo9kii/rails_6_api_template/)
- url: https://www.reddit.com/r/rails/comments/fo9kii/rails_6_api_template/
---
Recently I've created rails API template. Because it seems that there no any well configured and modern template around GitHub. Anyway the template still pretty raw. So I will be grateful for any feedback/pull requests/github stars and you comments. Thanks :)  [rails-api-template](https://github.com/pandwoter/rails-api-template)
## [8][Best lightweight alternative for text search SQL / Sphinx / Solr / Elastic Search ?](https://www.reddit.com/r/rails/comments/fondtn/best_lightweight_alternative_for_text_search_sql/)
- url: https://www.reddit.com/r/rails/comments/fondtn/best_lightweight_alternative_for_text_search_sql/
---
Been developing with Rails for years and years.   


For my projects, I tend to keep my dependencies and architecture as light as possible. I own some **website with mid side database** (dozens of thousands of entries). I keep looking for best practices for searches.  


Currently, I use to do so :

&amp;#x200B;

1. SQL only search (mysql) with Match queries (and good indices)  

2. Thinking Sphinx (works quite well out of the box and do not require tons of dependencies)

If some of you use things like Elastic Search (seems like very huge for small website, require java AFAIR) or Solr, what would be your advises ?
## [9][Enums as constants in Ruby DSL](https://www.reddit.com/r/rails/comments/fomw4d/enums_as_constants_in_ruby_dsl/)
- url: https://www.reddit.com/r/rails/comments/fomw4d/enums_as_constants_in_ruby_dsl/
---
Define enums as constants in Ruby easily and maintain the code with pleasure. See how in my latest [post](https://railsguides.net/enums-as-constants-in-ruby-dsl/)
## [10][Is there a way to connect ActionCable channels to SocketIO on the client side?](https://www.reddit.com/r/rails/comments/foj5ar/is_there_a_way_to_connect_actioncable_channels_to/)
- url: https://www.reddit.com/r/rails/comments/foj5ar/is_there_a_way_to_connect_actioncable_channels_to/
---

## [11][How important is cache? Does reddit use cache?](https://www.reddit.com/r/rails/comments/fo6iih/how_important_is_cache_does_reddit_use_cache/)
- url: https://www.reddit.com/r/rails/comments/fo6iih/how_important_is_cache_does_reddit_use_cache/
---
It is very strange to me the loading times even if nothing changed, maybe up/down votes only? I admit it might be my internet, but it is not very likely.

&amp;#x200B;

So I want to ask you guys, working on Production grade sites/apps, how important is caching? Fragment caching seems the only reasonably way for Rails, and it is both simple and can go complicated choosing the keys.

&amp;#x200B;

So answer to the first question is that cache is probably important.

I am most interested is there any technique for testing it, like in Capybara? Or is this one of those esoteric topics that it just comes with experience? That always kind of frightened me before I really studied TR5W.

&amp;#x200B;

How does it work in reality, on production site? Let's say you got to implement a small feature, how do you go about caching your implementation?
## [12][Anti-spam system for Comments. Wait 60 seconds](https://www.reddit.com/r/rails/comments/fo6a6d/antispam_system_for_comments_wait_60_seconds/)
- url: https://www.reddit.com/r/rails/comments/fo6a6d/antispam_system_for_comments_wait_60_seconds/
---
The website is growing up. I am tring to create an anti-spam system.

The users have to wait 60 seconds to post again.

in `models/user.rb` I added this

    def anti_spam_comment
    	@anti_spam_comment ||= (comments.last.created_at.to_time - Time.zone.today.to_datetime.to_time).floor
    end

and in `views/comments/index` I added this

    &lt;%= current_user.anti_spam_comment &gt; 60 %&gt;
    	You can not post
    &lt;% else %&gt;
    	Post Button
    &lt;% end %&gt;

But I have always the same value (and it is something like 48929)

**What is wrong?**

**Is there a best way to do it? Should I use the scope?**
