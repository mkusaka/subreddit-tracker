# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Ruby 3 introduces a new method 'name' on Symbol class](https://www.reddit.com/r/rails/comments/jcwa3o/ruby_3_introduces_a_new_method_name_on_symbol/)
- url: https://www.reddit.com/r/rails/comments/jcwa3o/ruby_3_introduces_a_new_method_name_on_symbol/
---
One of the important aspects of Ruby 3.0 is optimization. The part of that optimization is the introduction of the 'name' method for Symbol. [Ruby 3 Symbol Name](https://blog.bigbinary.com/2020/10/12/ruby-3-adds-symbol-name.html)
## [3][Best/secure way of having users (ROUTING)](https://www.reddit.com/r/rails/comments/jde1np/bestsecure_way_of_having_users_routing/)
- url: https://www.reddit.com/r/rails/comments/jde1np/bestsecure_way_of_having_users_routing/
---
Hi guys,

  
I am new to development and I have made a few ruby apps. I want to make something a bit more impressive for my CV. So I am making a finance app, this will be where a user can have many plans (these plans will have data etc.).  


What is the best way of setting up the routing for this, I am using devise due to getting online with users nice and quick but I am unsure of where to go next.  


These I think are the two options and I prefer the second (it gives less info and is cleaner), but I cant find a way of making the Plans specific to the user in a way that is 'secure'.  


1. users/user\_id/plans
2. /plans

I'd love some advice or someone to suggest some good content to read on this! Thanks
## [4][Generating (pfd, excel and docx) files on client side](https://www.reddit.com/r/rails/comments/jczhdk/generating_pfd_excel_and_docx_files_on_client_side/)
- url: https://www.reddit.com/r/rails/comments/jczhdk/generating_pfd_excel_and_docx_files_on_client_side/
---
Hi guys! I was wondering if do you guys use JS libraries to generate PDFs, Excel and Docs file on the client computer? If so, what are the options? I was thinking in something simple: after the client receives the data, the JS library would generate the file.
## [5][The character "-" in a xml builder](https://www.reddit.com/r/rails/comments/jcxl0l/the_character_in_a_xml_builder/)
- url: https://www.reddit.com/r/rails/comments/jcxl0l/the_character_in_a_xml_builder/
---
Hi guys,

i am working on a xml file (feed rss) fro Yandex turbo pages.

This is the code in my view/feed/yandex.rss.builder

    xml.instruct! :xml, version: '1.0'
    xml.rss  "xmlns:yandex": 'http://news.yandex.ru', "xmlns:media": 'http://search.yahoo.com/mrss/', "xmlns:turbo": 'http://turbo.yandex.ru', version: '2.0' do
      xml.channel do
        xml.title 'My Title'
        xml.description 'My Description'
        xml.link root_url(locale: nil)
        xml.language 'ru'
        xml.turbo:analytics, type: 'Yandex', id: '123456', params:"{ 'param' : 'val' }" do end
        xml.turbo:adNetwork type: 'Yandex', id: 'R-A-123456-7', turbo-ad-id: 'first_ad_place' do end
      end
    end

BUT I have a bug only on the line `xml.turbo:adNetwork`.

It seems that I can not add "`-`" character. I have a syntax report. How to solve?
## [6][How to merge redirect_to session.delete(:return_to) with (locale: @language.locale)](https://www.reddit.com/r/rails/comments/jcxv4n/how_to_merge_redirect_to_sessiondeletereturn_to/)
- url: https://www.reddit.com/r/rails/comments/jcxv4n/how_to_merge_redirect_to_sessiondeletereturn_to/
---
I use this system to switch the language on my website (warning about  `def update` and `redirect_to` )

    class LanguagesController &lt; ApplicationController
      def edit
        @language = Language.find_by(locale: I18n.locale)
        session[:return_to] ||= request.referer
        render layout: false
      end
    
      def update
        @language = Language.find(update_params[:id])
        if user_signed_in?
          setting = current_user.setting || current_user.build_setting
          setting.language = @language.id
          setting.save
          redirect_to session.delete(:return_to)
        else
          redirect_to root_path(locale: @language.locale)
        end
      end
    
      private
    
      def update_params
        params.require(:language).permit(:id)
      end
    end
    

As you can see in  `def update`, there is an if condition to check if the user is logged or not.

But as you can see there are two different `redirect_to` . If the user is logged, he is redirected into the previous page. If not, in homepage ()

I used this solution because I am not able to redirect the not-logged-user in the previous page with `(locale: @language.locale)`

**How to do?**

I tried with 

      def update
        @language = Language.find(update_params[:id])
        if user_signed_in?
          setting = current_user.setting || current_user.build_setting
          setting.language = @language.id
          setting.save
        end
        redirect_to session.delete(:return_to, locale: @language.locale)
      end

but it doesn't work (  I can not add two elements there)
## [7][Gem/Ideas for seeing outliers in a dataset.](https://www.reddit.com/r/rails/comments/jcpo2k/gemideas_for_seeing_outliers_in_a_dataset/)
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
## [8][Implementing HATEOAS in Rails API](https://www.reddit.com/r/rails/comments/jct7rc/implementing_hateoas_in_rails_api/)
- url: https://www.reddit.com/r/rails/comments/jct7rc/implementing_hateoas_in_rails_api/
---
I'm trying to learn Rails a bit more, and I'm not sure how to do this 'properly'.
From what I've seen, the routing url helpers should be able to do this, but I'm not sure if I'm misusing them or they simply can't work like this. For the record, I'm using `active_model_serializer`, and struggling with the following issues:

1) Getting the self reference URL. I could easily hardcode the link to the resource, then interpolate the `:id`. Easy, but not maintainable. I tried using a url_helper, but `todo_url(object_id)` ends up producing `"self": "http://localhost:3000/todos.17740"`

2) Getting URLs of resources not related to the object being serialized. For example, let's say in the case the client didn't send an auth token, return a link to the login/register endpoints.
## [9][Avo - Configuration-based, no-maintenance, extendable Ruby on Rails admin](https://www.reddit.com/r/rails/comments/jc85sc/avo_configurationbased_nomaintenance_extendable/)
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
## [10][Strategies for deploying Rails engines containing webpacker-compiled assets](https://www.reddit.com/r/rails/comments/jchwst/strategies_for_deploying_rails_engines_containing/)
- url: https://www.reddit.com/r/rails/comments/jchwst/strategies_for_deploying_rails_engines_containing/
---
Hello all,

I'm working on a Rails engine that contains assets that are compiled using webpacker. I'm wondering what the correct strategy is for packaging and deploying this gem. I've read the instructions [here](https://github.com/rails/webpacker/blob/master/docs/engines.md), but I find Step 7 a bit confusing.  I tried the instructions under "Use a separate middleware", but that didn't work. So my questions are:

1. Should I be pre-compiling the assets in my gem and including the compiled assets (in the gem's public/packs) with the published gem and load those assets from there in the app? OR
2. Should I *not* include the engine's assets, but run rake my-engine:webpacker:compile in the app?

Anyone have any experience with this?
## [11][What is the right gem to ban users? Devise, Pundit or Rolify?](https://www.reddit.com/r/rails/comments/jcdqwj/what_is_the_right_gem_to_ban_users_devise_pundit/)
- url: https://www.reddit.com/r/rails/comments/jcdqwj/what_is_the_right_gem_to_ban_users_devise_pundit/
---
What is your experience?
