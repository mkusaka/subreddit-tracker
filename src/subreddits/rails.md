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
## [3][Do I need to back up /dev/zero ?](https://www.reddit.com/r/rails/comments/fp28d7/do_i_need_to_back_up_devzero/)
- url: https://www.reddit.com/r/rails/comments/fp28d7/do_i_need_to_back_up_devzero/
---
My Ruby on Rails website application has a really big /dev/zero file (250+ GB). I'm not much of a developer, but a few times a year, I back up the site by FTPing all of the files to my desktop. Do I need this file?
## [4][DIY JWT Auth + Refresh Tokens + Blacklisting + Whitelisting](https://www.reddit.com/r/rails/comments/fov6tb/diy_jwt_auth_refresh_tokens_blacklisting/)
- url: https://www.reddit.com/r/rails/comments/fov6tb/diy_jwt_auth_refresh_tokens_blacklisting/
---
Hey all,

So I just made a gist with some code I extrapolated from a project I was working on. I ran into an issue, where I need my user to stay logged in indefinitely which lead to me needing to use refresh tokens/sliding sessions. Not only, I needed a way to both whitelist and blacklist tokens based on their jti for added security.

I spent a good while searching, found some excellent gems, but none could match exactly what I needed so I thought I'd do it DIY.

A couple hours, and modules later, I managed to get something working and thought I might share it for anyone in a similar predicament to me.

The code is public here -&gt; [https://gist.github.com/jesster2k10/e626ee61d678350a21a9d3e81da2493e](https://gist.github.com/jesster2k10/e626ee61d678350a21a9d3e81da2493e)

so anyone can take it, improve on it, use it or just look at it for fun.

I'm a relatively new ruby developer, only got into it last year but I think my code could help someone out.
## [5][Seeking paid mentor for my startup](https://www.reddit.com/r/rails/comments/fovc42/seeking_paid_mentor_for_my_startup/)
- url: https://www.reddit.com/r/rails/comments/fovc42/seeking_paid_mentor_for_my_startup/
---
Hello!

I'm looking for a long-term paid Rails mentor.

My name is Oliver Curting and I'm building a wedding planner platform for the Danish market, [www.bryllupsplaner.dk](https://www.bryllupsplaner.dk).

I'm looking for a mentor since I find it increasingly difficult to keep momentum as a solo founder. Programming the wedding planner is a behemoth of a project for me since my forte is in online marketing (self-employed for five years). I often find myself stumbling into difficulties I just can't seem to crack – even with Stackoverflow's help. They slow me unnecessarily down, as well as giving me a sense of loneliness. I'm doing elite sport and I know how important it is to have a trainer/coach/mentor and I could really use one here.

It is my second Rails app. My first one is [www.brudevalsen.dk](https://www.brudevalsen.dk), which I launched a few years ago (and is now the most popular way to learn the bridal waltz in Denmark 🥳). A simple paid video course with Devise and Stripe integration. Just to give you an idea of my skill level.

Please DM me with your experience and rates if you're interested. Thank you! :-)
## [6][Building rails api with multiple Auth strategies](https://www.reddit.com/r/rails/comments/fosxp4/building_rails_api_with_multiple_auth_strategies/)
- url: https://www.reddit.com/r/rails/comments/fosxp4/building_rails_api_with_multiple_auth_strategies/
---
Hi, I've stated working on rails api-only app which will be backend for mobile app. I need to provide sign-in with my own authentication (email + password), google, and Facebook. Also I need jwt as token format. How can I implement that in safe and as easy as possible authentication? I see how can I accomplish that in classic web app but I don't know how can I implement that in api-only app. I can use devise user, google and Facebook clients and jwt gem, but I don't know if it's secure and elegant way, there is devise-jwt but I'm not sure if it possible to use it with Google and Facebook authenticators.
## [7][Strategy for a model with all 50 states and account selection?](https://www.reddit.com/r/rails/comments/fot9k8/strategy_for_a_model_with_all_50_states_and/)
- url: https://www.reddit.com/r/rails/comments/fot9k8/strategy_for_a_model_with_all_50_states_and/
---
I have a model named \`account\` and on creation of the account, the user can select any of the 50 states (51 with D.C.). I want to treat the states as a separate model. How would implement this? Would I create a model named \`state\` and then create all 51 instances, each with a unique id? And would it be \`account has\_many states\`?
## [8][Scintilla is a gem for one-time scripts on Rails apps](https://www.reddit.com/r/rails/comments/fonh4l/scintilla_is_a_gem_for_onetime_scripts_on_rails/)
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
## [9][for 22x times the slowness, what extras am I getting by doing `bundle exec rails dbconsole` instead of `psql yourapp_development`?](https://www.reddit.com/r/rails/comments/foqhb8/for_22x_times_the_slowness_what_extras_am_i/)
- url: https://www.reddit.com/r/rails/comments/foqhb8/for_22x_times_the_slowness_what_extras_am_i/
---

```
$ time bundle exec rails dbconsole
psql (12.1)
Type "help" for help.

myapp_development=# \q

real	0m8.890s
user	0m3.787s
sys	0m3.098s
$ 
```

VERSUS

```
$ time psql myapp_development
psql (12.1)
Type "help" for help.

myapp_development=# \q

real	0m0.461s
user	0m0.006s
sys	0m0.008s
$ 
```
## [10][MarkdownRender, redcarpet and custimization text...](https://www.reddit.com/r/rails/comments/fosafx/markdownrender_redcarpet_and_custimization_text/)
- url: https://www.reddit.com/r/rails/comments/fosafx/markdownrender_redcarpet_and_custimization_text/
---
I edited my MarkdownRenderer in this way.

    class MarkdownRenderer &lt; Redcarpet::Render::HTML
      def preprocess(text)
        text_formatting(text)
      end
    
      def text_formatting(text)
        text.gsub! /(^|\&gt;|\s)@(\w+)/ do
          "#{$1}&lt;a href=\"/user/#{$2}\"&gt;@#{$2}&lt;/a&gt;"
       end
       text
      end
    
      def autolink(link, link_type)
        case link
          when /((?:https?)?:\/\/\S*\.(?i:png|jpe?g|gif))/ then image_link(link)
          else normal_link(link)
        end
      end
    
        def image_link(link)
        "&lt;img src=\"#{link}\"&gt;"
      end
    
      def normal_link(link)
        "&lt;a href=\"#{link}\"&gt;#{link}&lt;/a&gt;"
      end

If I add a link or an image-link, it works perfectly. But if I add `@mark` (for example), it show me `&lt;a href="/user/mark"&gt;@mark&lt;/a&gt;` with the html tags. 

**In image\_link or normal\_link, it doesn't show the html tags, but in the text\_formatting, yes. I don't know how to solve it.**

&amp;#x200B;

p.s. In the views I added `&lt;%= parse_markdown(comment.text).html_safe %&gt;`

and in the comments\_helper I have this

    module CommentsHelper
      def parse_markdown(text)
        markdown = Redcarpet::Markdown.new(MarkdownRenderer, hard_wrap: true, autolink: true, space_after_headers: true)
        markdown.render(text)
      end
## [11][Registration Form - Skills](https://www.reddit.com/r/rails/comments/fowg7f/registration_form_skills/)
- url: https://www.reddit.com/r/rails/comments/fowg7f/registration_form_skills/
---
I'm looking to add the option for users to be able to list their skills upon registration/account edit but am unable to think what the best way to do so is.   


The skills entry will have multiple inputs from a user from 0 to 10 lets say.    
What is the best way to log this information to the user table, JSON Obj?   
or is it best to have a skills table linked to each user?
## [12][ActiveStorage DB relation explainer please](https://www.reddit.com/r/rails/comments/foop3u/activestorage_db_relation_explainer_please/)
- url: https://www.reddit.com/r/rails/comments/foop3u/activestorage_db_relation_explainer_please/
---
Hi there. I’m migrating from Paperclip to AS and so far doing fine but got stuck converting a gem that uses multiple images per record. 

It’s a has_many_attached relation, with an Offer model having many OfferImages. 

Should I have one Attachment per record, then multiple Blobs per Attachment? Or one Attachment per Blob, meaning many Blobs and Attachments per record? 

I can’t find this info anywhere, but I suspect that because each Attachment has a blob_id there’s actually many Attachment and many Blobs for each Offer. 

Please enlighten me as to how the Attachment and Blob tables should look.
