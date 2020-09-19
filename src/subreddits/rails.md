# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][GCS File browser/viewer in Rails?](https://www.reddit.com/r/rails/comments/ivo2aq/gcs_file_browserviewer_in_rails/)
- url: https://www.reddit.com/r/rails/comments/ivo2aq/gcs_file_browserviewer_in_rails/
---
Hello, r/rails! 

I'd like to build a Google Cloud Storage bucket file browser and manager in Rails. The main functionality would be browsing files on the bucket (while recreating the folder hierarchy) and downloading one/many files from it.

How would you go about implementing that? What gems would you use etc..?

Thank you!
## [3][Validating inputs within the context of other inputs](https://www.reddit.com/r/rails/comments/ivi361/validating_inputs_within_the_context_of_other/)
- url: https://www.reddit.com/r/rails/comments/ivi361/validating_inputs_within_the_context_of_other/
---
I have three inputs that require the user to rank three resources from 1 to 3.  Validating that each input is in the range of 1..3 was easy, but I'm stumped on how to validate their uniqueness, or that the group adds up to 6.  

Any pointers would be appreciated.
## [4][Editing before Loading DateTime into Input Field](https://www.reddit.com/r/rails/comments/ivcefa/editing_before_loading_datetime_into_input_field/)
- url: https://www.reddit.com/r/rails/comments/ivcefa/editing_before_loading_datetime_into_input_field/
---
I want to have simple datetime picker, one field for date other for time. But when I am editing  it inside the form it loads the whole datetime object string date, would there be a way to format that long date and time to simple date in one field and simple time in other field.

**What I get:**

https://preview.redd.it/thu887zxfyn51.png?width=1432&amp;format=png&amp;auto=webp&amp;s=d65b4600f2b75a05f0e92ed0fa63404a60919dc4

**What I want to get:**

https://preview.redd.it/bb4aqu90gyn51.png?width=1472&amp;format=png&amp;auto=webp&amp;s=8fac06eafd7d485862960ef75b85dec8ecd05381
## [5][Please help me with my N+1 Problem](https://www.reddit.com/r/rails/comments/iv2ga7/please_help_me_with_my_n1_problem/)
- url: https://www.reddit.com/r/rails/comments/iv2ga7/please_help_me_with_my_n1_problem/
---
I have a model 

    class Posts
      has_many :comments
 
      def latest_comment
        comments.order(:created_at).last
      end
    end

If I do

    Post.includes(:comments).each do |post|
        puts post.latest_comment
    end

The ***instance method*** seems to ignore the ```includes``` and runs a query for each post to get the latest comment for each post.

I'm able to circumvent this via a static/class method which is:

    def self.get_latest_comment(post)
        post.comments.sort_by{|x|x.created_at}.last
    end

but this feels inelegant. Is this the only solution?

Thanks in advance!

**Edit** fixed my example
## [6][Any number of spaces after :: is okay?](https://www.reddit.com/r/rails/comments/iv3849/any_number_of_spaces_after_is_okay/)
- url: https://www.reddit.com/r/rails/comments/iv3849/any_number_of_spaces_after_is_okay/
---
How and why does this work?   


    class User &lt; ActiveRecord::Base
      CONSTANT = 1
    end

And when I add space in between class and const name it still prints the value.   


    User::CONSTANT # prints 1
    User::                   CONSTANT # prints 1
## [7][Is this a banana solution?](https://www.reddit.com/r/rails/comments/iutmqo/is_this_a_banana_solution/)
- url: https://www.reddit.com/r/rails/comments/iutmqo/is_this_a_banana_solution/
---
A few days ago I posted a problem I had with [Nginx (also Apache)](https://www.reddit.com/r/rails/comments/ito2ou/trying_to_deploy_my_web_app_on_digitalocean/), so at the end what I did was:

1. Use standalone passenger and serve my web app on ports 5000 (HTTP) and 5001 (HTTPS)
2. redirect port `80` to `5000` and 443 to `5001` with `iptables`, example: `sudo iptables -t nat -I PREROUTING -p tcp --dport 80 -j REDIRECT --to-ports 5000`

Is this a banana solution?
## [8][webpacker breaks CSP](https://www.reddit.com/r/rails/comments/iusays/webpacker_breaks_csp/)
- url: https://www.reddit.com/r/rails/comments/iusays/webpacker_breaks_csp/
---
I would like to set webpacker to user a nonce. My app's CSP breaks ever since I installed it. I think it has something to do with `style-loader`. The browser console error looks like this:

`Content Security Policy: The page’s settings observed the loading of a resource at inline (“style-src”). A CSP report is being sent. injectStylesIntoStyleTag.js:117`

`Content Security Policy: The page’s settings observed the loading of a resource at inline (“style-src”). A CSP report is being sent. injectStylesIntoStyleTag.js:190`

The code in question:

    function insertStyleElement(options) {
      var style = document.createElement('style');      
      ...
      if (typeof options.insert === 'function') {
        options.insert(style);
      } else {
        var target = getTarget(options.insert || 'head');
        if (!target) {
          throw new Error("Couldn't find a style target. This probably means that the value for the 'insert' parameter is invalid.");
        }
        target.appendChild(style); //LINE 117//
      }
      return style;
    }

And:

    function applyToTag(style, options, obj) {
      var css = obj.css;
      ...
      if (style.styleSheet) {
        style.styleSheet.cssText = css;
      } else {
        while (style.firstChild) {
          style.removeChild(style.firstChild);
        }
        style.appendChild(document.createTextNode(css)); //LINE 190//
      }
    }

How do I whitelist this? Note that I'm using Turbolinks. This ([https://webpack.js.org/guides/csp/](https://webpack.js.org/guides/csp/)) says to add `__webpack_nonce__ = 'random'` to my entry file ( in this case `app/javascript/packs/application.js`), yet adding that nonce to my csp file has no effect on the style-src violation. Which in this case, looks like this: `config.style_src :self, '`[`https://fonts.googleapis.com`](https://fonts.googleapis.com)`', 'nonce-random'`
## [9][anyone successfully convert a model into a queue for ETL purposes?](https://www.reddit.com/r/rails/comments/iumw4v/anyone_successfully_convert_a_model_into_a_queue/)
- url: https://www.reddit.com/r/rails/comments/iumw4v/anyone_successfully_convert_a_model_into_a_queue/
---
Please correct me if you see any better way to organize my problem. 

&amp;#x200B;

I am doing data backup from a 3rd party provider. Their data is provided daily, and is not finalized for 14-days. 

&amp;#x200B;

Meaning we basically have 14 days of "fresh" data to fetch, and the oldest is more important than the newest.

&amp;#x200B;

Everyday we take all records for the last 14 days and mark them as "new"

Every 10 minutes we fetch data on records sorted oldest to newest mostly to make sure we don't hit API limits but i really don't have much observability into how often we retry jobs due to the API limits. 

&amp;#x200B;

Anyone have a rock-solid solution that has ETL experience and whatnot?
## [10][Zipping active storage has_one_attached &amp; has_many_attached files in one zip file for a model](https://www.reddit.com/r/rails/comments/iui5g4/zipping_active_storage_has_one_attached_has_many/)
- url: https://www.reddit.com/r/rails/comments/iui5g4/zipping_active_storage_has_one_attached_has_many/
---
I'm using the active\_storage-send-zip, but it seems to only focus on has\_many\_attached.  I need to zip all files from the model at once. I can't seem to combine the two, to get the results that I want.

&amp;#x200B;

For example:

User.rb

has\_one\_attached :avatar

has\_many\_attached :photos

&amp;#x200B;

zip all photos + their avatar image
## [11][rails code style for model callbacks](https://www.reddit.com/r/rails/comments/iujskb/rails_code_style_for_model_callbacks/)
- url: https://www.reddit.com/r/rails/comments/iujskb/rails_code_style_for_model_callbacks/
---
From robocops style guidelines this is what they suggest. [https://github.com/rubocop-hq/rails-style-guide#callbacks-order](https://github.com/rubocop-hq/rails-style-guide#callbacks-order)

    #bad
    class Person
      after_commit :after_commit_callback
      before_validation :before_validation_callback
    end
    
    #good
    class Person
      before_validation :before_validation_callback
      after_commit :after_commit_callback
    end

My question is suppose we have something like this  


    class Person
      after_save :after_user_update on: :update
      after_save :after_user_save on: :create
      before_validation :validate_user_update on: :update
      before_validation :validate_user_destroy on: :destroy
    end

In that case, what is the best approach? 

Option 1) Group all before\_validations for create, update, delete and then after\_validation and so on

    class Person
      before_validation :validate_user_update on: :update
      before_validation :validate_user_destroy on: :destroy
      after_save :after_user_save on: :create
      after_save :after_user_update on: :update
    end

Or   
go with the all the callbacks together for each of the action, first set for create and then update and delete.   
from the list here [https://guides.rubyonrails.org/active\_record\_callbacks.html#available-callbacks](https://guides.rubyonrails.org/active_record_callbacks.html#available-callbacks)

    class Person
      before_validation :validate_user_update on: :update
      after_save :after_user_save on: :create
      before_validation :validate_user_destroy on: :destroy
      after_save :after_user_update on: :update
    end

What are your thoughts?
