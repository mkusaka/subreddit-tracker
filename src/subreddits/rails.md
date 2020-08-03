# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/
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
## [2][Moving to Rails in 2020](https://www.reddit.com/r/rails/comments/i2tg0x/moving_to_rails_in_2020/)
- url: https://www.reddit.com/r/rails/comments/i2tg0x/moving_to_rails_in_2020/
---
Hey people. I'm a backend developer working primarily with PHP7.x and TypeScript, with some experience with Go, but in all honesty, I always had a thing for Ruby.   


I've been thinking about learning Ruby and then Rails properly and trying to land a job that would allow me to use them. The only thing that pulls me back is a question of whether or not new stuff is actually created with Ruby/Rails or its mostly legacy code at this point. I'm honestly tired of PHP (mostly community, the language is just fine) and TypeScript doesn't bring to much joy to work with (although it is a great language).   


To sum up my question, are there new Rails apps coming out and, with 5-6 years of experience would I have to look for a junior position, given that I've got experience with other languages and know a fair bit of system design, best practices and design patterns?
## [3][Same association twice for one single model](https://www.reddit.com/r/rails/comments/i2v2z0/same_association_twice_for_one_single_model/)
- url: https://www.reddit.com/r/rails/comments/i2v2z0/same_association_twice_for_one_single_model/
---
Hello, I have a table with synonyms linked to a word and another word ("word\_linked") (which is the synonym of the first one).

Table 1 : Word

Table 2 : Synonym (word\_id, word\_linked)  


How can I create an association with both words?

&amp;#x200B;

For now I have in my Synonym model :

`belongs_to :word`

But I would also need something like this:

`has_one :word, :through =&gt; :word_linked`

&amp;#x200B;

The goal is to be able to load the word association when I do Synonym.find(1).word\_linked  


Is it possible to do that? Thanks.
## [4][Backing up and Restoring ActiveStorage files](https://www.reddit.com/r/rails/comments/i2ucgd/backing_up_and_restoring_activestorage_files/)
- url: https://www.reddit.com/r/rails/comments/i2ucgd/backing_up_and_restoring_activestorage_files/
---
I am building an app where a user simply uploads some files to a Rails app and I want add some kind of backup functionality built into the app.

The app should be able to auto-backup the database and the files, compress them into a `zip` file and then upload them to a custom or third party service.

I am thinking that I will serialize the user's database records e.g. the `User` instance and his `posts` (where the files are being referenced to and uploaded) but I don't know what can I do with `ActiveStorage` files.

Should I serialize the `active_storage_*` tables and backup the `/storage` directory as is ? Could restoring them into a new Rails installation and referencing the *same* user work?
## [5][Credentials failure from NGINX](https://www.reddit.com/r/rails/comments/i2u8bb/credentials_failure_from_nginx/)
- url: https://www.reddit.com/r/rails/comments/i2u8bb/credentials_failure_from_nginx/
---
Hosting my app on DO vps. Have rails native credentials stored. (Rails.application.credentials)

They working fine from `rails c` console but NGINX doesn't see them at all, logging `undefined method [] for NilClass` for them. Can you help me out guys?
## [6][Hooks in ActionView::TestCase](https://www.reddit.com/r/rails/comments/i2wgd8/hooks_in_actionviewtestcase/)
- url: https://www.reddit.com/r/rails/comments/i2wgd8/hooks_in_actionviewtestcase/
---
I am modifying test cases in a test file that looks like this. 

    class TestClass &lt; ActionView
      def setup # called before each test
      end
      def teardown # called after each test
      end
      def test_this
     end
    end

I'm looking for a hook that will run after all the tests complete running. 

I am not sure where to check the documentation for the hooks available in this class. Any help would be great. It's kinda old code and I couldn't find anything on googling. 

Please do not suggest gems to achieve this.
## [7][nonce for style-src added to script-src by gem - how to circumvent?](https://www.reddit.com/r/rails/comments/i2oehp/nonce_for_stylesrc_added_to_scriptsrc_by_gem_how/)
- url: https://www.reddit.com/r/rails/comments/i2oehp/nonce_for_stylesrc_added_to_scriptsrc_by_gem_how/
---
I have a single inline style tag in a Rails app, and I'm trying to whitelist it with a nonce.

The error: `Content Security Policy: The page’s settings observed the loading of a resource at inline (“style-src”). A CSP report is being sent.`

The style gets injected by a helper when using the [invisible\_captcha](https://github.com/markets/invisible_captcha) gem:

`&lt;%= invisible_captcha nonce: true %&gt;`

Comparing the view and the header shows that the nonce \*is\* being added to both, but it's being added to `script-src` instead of `style-src`.

View: `&lt;style media="screen" nonce="5Cq/QyoJ5Co+LdarO1uvrg=="&gt;`

Header: `Content-Security-Policy-Report-Only script-src 'self' https: 'nonce-5Cq/QyoJ5Co+LdarO1uvrg=='; style-src 'self' https:;`

This is my **content\_security\_policy.rb**:

    Rails.application.config.content_security_policy do |config|
     config.script_src  :self, :https    
    end
    
    Rails.application.config.content_security_policy_nonce_generator = -&gt; request { SecureRandom.base64(16) }
    Rails.application.config.content_security_policy_nonce_directives = %w(style-src script-src)

Why is it going to the wrong directive, and how can I make it go to `style-src`? The helper adds a random class name, so hashing the injected code is out of the question.
## [8][Creating Barcode Inventory Labels](https://www.reddit.com/r/rails/comments/i2epfu/creating_barcode_inventory_labels/)
- url: https://www.reddit.com/r/rails/comments/i2epfu/creating_barcode_inventory_labels/
---
I would love to know if anyone has done this before and any specifics if you did. I was able create some Avery labels with the help of Prawn Label gem. But I would really like to print the labels by rotating the font 90% but could never get the alignment correct. Next I think I will try a raw zpl file to print to a zebra printer.
## [9][Need help with bundle](https://www.reddit.com/r/rails/comments/i2iiyd/need_help_with_bundle/)
- url: https://www.reddit.com/r/rails/comments/i2iiyd/need_help_with_bundle/
---
Earlier i was trying to $bundle update, when it gave an error message saying "An error occured while trying to install puma (version) and Bundler cant continue. Make sure gem install puma -v (version) --source https//:rubygems.org succeeds before bundling" And when i try to do also that, it gives another error.

Does anyone know how to fix this?
## [10][Are you gonna stick to Mac after Apple Silicon?](https://www.reddit.com/r/rails/comments/i2qo7k/are_you_gonna_stick_to_mac_after_apple_silicon/)
- url: https://www.reddit.com/r/rails/comments/i2qo7k/are_you_gonna_stick_to_mac_after_apple_silicon/
---


[View Poll](https://www.reddit.com/poll/i2qo7k)
## [11][Rails.env.production?](https://www.reddit.com/r/rails/comments/i2faqf/railsenvproduction/)
- url: https://www.reddit.com/r/rails/comments/i2faqf/railsenvproduction/
---
Hi guys! I am a front-end. New on rails.

I'm noticing that the previous back end developer added this in the head

    &lt;% if Rails.env.production? %&gt;
    	&lt;style type="text/css" media="all"&gt;&lt;%= inline_asset('application.css') %&gt;&lt;/style&gt;
    &lt;% else %&gt;
    	&lt;%= stylesheet_link_tag 'application', media: 'all' %&gt;
    &lt;% end %&gt;

Why if `Rails.env.production?`? What is it?
