# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Adding users &amp; authorisation to a Rails app | Live-streaming Ruby on Rails](https://www.reddit.com/r/rails/comments/je15gu/adding_users_authorisation_to_a_rails_app/)
- url: https://www.reddit.com/r/rails/comments/je15gu/adding_users_authorisation_to_a_rails_app/
---
I've done a couple more live-streams on the Sundae Club channel, where each week we live-stream building an app and look at features from a learner's perspective. In the last couple of weeks we've looked at:

[https://www.youtube.com/watch?v=b-J1TuimRtk](https://www.youtube.com/watch?v=b-J1TuimRtk) - Adding users to a Rails app

[https://www.youtube.com/watch?v=77Z_u2QDTZA](https://www.youtube.com/watch?v=77Z_u2QDTZA) - Adding user permissions to a Rails app

Both videos follow roughly the same format where we look at some of the options that are available and then implement one of them on the app and try to see it in action. Chapters are listed on each videos to help you navigate and save you having to watch the whole stream.

New streams are each Sunday at 10.30am (UK time) and available any time on YouTube.
## [3][I’ve recorded a walk-through video of Action Cable’s JavaScript side](https://www.reddit.com/r/rails/comments/je185l/ive_recorded_a_walkthrough_video_of_action_cables/)
- url: https://www.reddit.com/r/rails/comments/je185l/ive_recorded_a_walkthrough_video_of_action_cables/
---
I’ve recorded a walk-through video of Action Cable’s JavaScript side behind-the-scenes ––&gt;

[https://twitter.com/liroyleshed/status/1318173979983728643](https://twitter.com/liroyleshed/status/1318173979983728643)

(Couldn't upload the full video to Twitter, but working on a fix :)

Thank you for watching.
## [4][What is the best course for Ruby &amp; Ruby on Rails?](https://www.reddit.com/r/rails/comments/jdq7pu/what_is_the_best_course_for_ruby_ruby_on_rails/)
- url: https://www.reddit.com/r/rails/comments/jdq7pu/what_is_the_best_course_for_ruby_ruby_on_rails/
---
Hello everyone,

although I am already working a bit with Ruby on Rails, I want to focus myself more and more and deepen my knowledge with a deep, thorough course about Ruby, RoR (6), Stimulus JS etc.

What is the best course for that, in your opinion? Could be one entire course or something like a Website with multiple modules (I already know GoRails, but there is no real structure for me and also, Chris is too fast for me sometimes :D (I love his courses tho!)).

I learn the best when creating a project during the course, but it could be anything. Thanks a lot in advance!
## [5][Same filter with same callback but for different conditions for different transactions](https://www.reddit.com/r/rails/comments/jdukz2/same_filter_with_same_callback_but_for_different/)
- url: https://www.reddit.com/r/rails/comments/jdukz2/same_filter_with_same_callback_but_for_different/
---
Suppose I have two callbacks, with the same callback methods but to be called for different transactions for different conditions like below.  


     after_rollback :model_method, on: :create, if: -&gt; { self.name == "create" }
     after_rollback :model_method, on: :update, if: -&gt; { self.name == "update" }

I can see that this is being overridden.   


    User._rollback_callbacks.count // 1

How can I achieve something similar without it being overridden?

I was thinking something like this

    def should_do_model_method?
     if transaction_include_any_action([:update])
        self.name == "update"
     elsif transaction_include_any_action([:create])
        self.name == "create"
     end
    end

Is this the way to achieve this or is there any other way?
## [6][Devise .. Deleting the Account.](https://www.reddit.com/r/rails/comments/jdi7qa/devise_deleting_the_account/)
- url: https://www.reddit.com/r/rails/comments/jdi7qa/devise_deleting_the_account/
---
I added this in my views/settings/\_form.html.erb

    &lt;%= button_to "Cancel my account", registration_path(resource_name), data: { confirm: "Are you sure?"}, method: :delete %&gt;

If the user has no comment/votes/following/etc, it works perfectly and delete the account.

But if he posted just one comment, I have this error

    PG::ForeignKeyViolation: ERROR: update or delete on table "users" violates foreign key constraint "fk rails 40412334d9" on table "comments" DETAIL: Key (id)=(3872) is still referenced from table "comments". : DELETE FROM "users" WHERE "users"."id" = $1

How to delete every comment of the user and then to delete the account?
## [7][Best/secure way of having users (ROUTING)](https://www.reddit.com/r/rails/comments/jde1np/bestsecure_way_of_having_users_routing/)
- url: https://www.reddit.com/r/rails/comments/jde1np/bestsecure_way_of_having_users_routing/
---
Hi guys,

  
I am new to development and I have made a few ruby apps. I want to make something a bit more impressive for my CV. So I am making a finance app, this will be where a user can have many plans (these plans will have data etc.).  


What is the best way of setting up the routing for this, I am using devise due to getting online with users nice and quick but I am unsure of where to go next.  


These I think are the two options and I prefer the second (it gives less info and is cleaner), but I cant find a way of making the Plans specific to the user in a way that is 'secure'.  


1. users/user\_id/plans
2. /plans

I'd love some advice or someone to suggest some good content to read on this! Thanks
## [8][Ruby 3 introduces a new method 'name' on Symbol class](https://www.reddit.com/r/rails/comments/jcwa3o/ruby_3_introduces_a_new_method_name_on_symbol/)
- url: https://www.reddit.com/r/rails/comments/jcwa3o/ruby_3_introduces_a_new_method_name_on_symbol/
---
One of the important aspects of Ruby 3.0 is optimization. The part of that optimization is the introduction of the 'name' method for Symbol. [Ruby 3 Symbol Name](https://blog.bigbinary.com/2020/10/12/ruby-3-adds-symbol-name.html)
## [9][Generating (pfd, excel and docx) files on client side](https://www.reddit.com/r/rails/comments/jczhdk/generating_pfd_excel_and_docx_files_on_client_side/)
- url: https://www.reddit.com/r/rails/comments/jczhdk/generating_pfd_excel_and_docx_files_on_client_side/
---
Hi guys! I was wondering if do you guys use JS libraries to generate PDFs, Excel and Docs file on the client computer? If so, what are the options? I was thinking in something simple: after the client receives the data, the JS library would generate the file.
## [10][The character "-" in a xml builder](https://www.reddit.com/r/rails/comments/jcxl0l/the_character_in_a_xml_builder/)
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
## [11][How to merge redirect_to session.delete(:return_to) with (locale: @language.locale)](https://www.reddit.com/r/rails/comments/jcxv4n/how_to_merge_redirect_to_sessiondeletereturn_to/)
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
