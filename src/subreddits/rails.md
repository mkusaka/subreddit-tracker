# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][File types as part of ActiveStorage uploads](https://www.reddit.com/r/rails/comments/ivyqcw/file_types_as_part_of_activestorage_uploads/)
- url: https://www.reddit.com/r/rails/comments/ivyqcw/file_types_as_part_of_activestorage_uploads/
---
Hi /r/rails

I am working on an internal Rails app around adding products to our warehouse inventory to replace a paper based filing system.

When an employee adds a product, they can (if its available) add the OEM's product data sheet to their upload.

There are also other PDF documents that become part of the product's folder including:

* Procurement (who would add receipts for buys from the suppliers/manufacturers)
* Occupational safety team (They have a PDF form they fill out OK'ing warehousing the item and if it has to be stored at different heights or separated from other items)
* Receiving (they scan items when received and would upload an exported PDF from their receiving system)
* Loss prevention team (some items are kept caged and they identify these)
* Legal (A few items are export controlled and legal provides an opinion on if they need to be identified as such)

I'd like to be able to identify these files as part of display of files and workflow through the app

What is the easiest way to manage this in rails?

My first thought was to have my product model have

`has_many :files`

Then create a file model and table that has an enum for all the types of files ("MSDS", "Receipt", "Receiving", "Safety", "Security", "Legal") and then this file model to have:

    has_many_attached :msds
    has_many_attached :receipt
    has_many_attached :receiving_reports
    has_many_attached :safety
    has_many_attached :security
    has_many_attached :legal

Is this the best way to go about this? Am I over or under thinking this?
## [3][Difficulty with application.js syntax (//= vs require)](https://www.reddit.com/r/rails/comments/ivy7r8/difficulty_with_applicationjs_syntax_vs_require/)
- url: https://www.reddit.com/r/rails/comments/ivy7r8/difficulty_with_applicationjs_syntax_vs_require/
---
I'm having a popper.js issue with an import. A lot of documentation I see online has different syntax. I see two types:

//= import\_name

and

require(import\_name)
## [4][GCS File browser/viewer in Rails?](https://www.reddit.com/r/rails/comments/ivo2aq/gcs_file_browserviewer_in_rails/)
- url: https://www.reddit.com/r/rails/comments/ivo2aq/gcs_file_browserviewer_in_rails/
---
Hello, r/rails! 

I'd like to build a Google Cloud Storage bucket file browser and manager in Rails. The main functionality would be browsing files on the bucket (while recreating the folder hierarchy) and downloading one/many files from it.

How would you go about implementing that? What gems would you use etc..?

Thank you!
## [5][Validating inputs within the context of other inputs](https://www.reddit.com/r/rails/comments/ivi361/validating_inputs_within_the_context_of_other/)
- url: https://www.reddit.com/r/rails/comments/ivi361/validating_inputs_within_the_context_of_other/
---
I have three inputs that require the user to rank three resources from 1 to 3.  Validating that each input is in the range of 1..3 was easy, but I'm stumped on how to validate their uniqueness, or that the group adds up to 6.  

Any pointers would be appreciated.
## [6][Editing before Loading DateTime into Input Field](https://www.reddit.com/r/rails/comments/ivcefa/editing_before_loading_datetime_into_input_field/)
- url: https://www.reddit.com/r/rails/comments/ivcefa/editing_before_loading_datetime_into_input_field/
---
I want to have simple datetime picker, one field for date other for time. But when I am editing  it inside the form it loads the whole datetime object string date, would there be a way to format that long date and time to simple date in one field and simple time in other field.

**What I get:**

https://preview.redd.it/thu887zxfyn51.png?width=1432&amp;format=png&amp;auto=webp&amp;s=d65b4600f2b75a05f0e92ed0fa63404a60919dc4

**What I want to get:**

https://preview.redd.it/bb4aqu90gyn51.png?width=1472&amp;format=png&amp;auto=webp&amp;s=8fac06eafd7d485862960ef75b85dec8ecd05381
## [7][Please help me with my N+1 Problem](https://www.reddit.com/r/rails/comments/iv2ga7/please_help_me_with_my_n1_problem/)
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
## [8][Any number of spaces after :: is okay?](https://www.reddit.com/r/rails/comments/iv3849/any_number_of_spaces_after_is_okay/)
- url: https://www.reddit.com/r/rails/comments/iv3849/any_number_of_spaces_after_is_okay/
---
How and why does this work?   


    class User &lt; ActiveRecord::Base
      CONSTANT = 1
    end

And when I add space in between class and const name it still prints the value.   


    User::CONSTANT # prints 1
    User::                   CONSTANT # prints 1
## [9][Is this a banana solution?](https://www.reddit.com/r/rails/comments/iutmqo/is_this_a_banana_solution/)
- url: https://www.reddit.com/r/rails/comments/iutmqo/is_this_a_banana_solution/
---
A few days ago I posted a problem I had with [Nginx (also Apache)](https://www.reddit.com/r/rails/comments/ito2ou/trying_to_deploy_my_web_app_on_digitalocean/), so at the end what I did was:

1. Use standalone passenger and serve my web app on ports 5000 (HTTP) and 5001 (HTTPS)
2. redirect port `80` to `5000` and 443 to `5001` with `iptables`, example: `sudo iptables -t nat -I PREROUTING -p tcp --dport 80 -j REDIRECT --to-ports 5000`

Is this a banana solution?
## [10][webpacker breaks CSP](https://www.reddit.com/r/rails/comments/iusays/webpacker_breaks_csp/)
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
## [11][anyone successfully convert a model into a queue for ETL purposes?](https://www.reddit.com/r/rails/comments/iumw4v/anyone_successfully_convert_a_model_into_a_queue/)
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
