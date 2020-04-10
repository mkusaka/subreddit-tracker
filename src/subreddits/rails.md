# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/)
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/fx6je4/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/fx6je4/personal_projects_show_off_your_own_project_andor/
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
## [3][I've launched a new newsletter with all my Ruby posts](https://www.reddit.com/r/rails/comments/fyd68r/ive_launched_a_new_newsletter_with_all_my_ruby/)
- url: https://www.reddit.com/r/rails/comments/fyd68r/ive_launched_a_new_newsletter_with_all_my_ruby/
---
I've launched a new newsletter That Weekly Tech with 3 of my Ruby on Rails posts &amp; more to come ––&gt;

[https://www.thatweeklytech.com](https://www.thatweeklytech.com/)

\- Active storage file upload behind the scenes

\- Geared pagination behind the scenes

\- JSON in Rails behind the scenes

\- More content to come...

I didn't want Medium to charge/put my content behind a paywall. So now it's, of course, all free.
## [4][From form.radio_button to form.hidden_field](https://www.reddit.com/r/rails/comments/fyd43w/from_formradio_button_to_formhidden_field/)
- url: https://www.reddit.com/r/rails/comments/fyd43w/from_formradio_button_to_formhidden_field/
---
have this form.radio\_button

    &lt;%= form.radio_button :id, language.id, selected: selected %&gt;
    &lt;%= language.name %&gt;

and the result is

    &lt;input id="language_id_7", value="7", name="language[id]", typer="radio"&gt;
    Spanish

**&gt;&gt;&gt; I want to turn it in a form.hidden\_field**

&amp;#x200B;

I edited it in this way

    &lt;%= form.hidden_field :id, language.id %&gt;
    &lt;%= language.name %&gt;

But I have this error: **Undefined method "merge" for 13:Fixium**

&amp;#x200B;

I also tried with

    &lt;%= form.hidden_field :id, :value =&gt; language.id %&gt;
    &lt;%= language.name %&gt;

But then the result is (check out the id)

    &lt;input id="language_id", value="7", name="language[id]", typer="hidden"&gt;
    Spanish

**and it doesn't work. How to solve?**
## [5][Migration depending on recently created table](https://www.reddit.com/r/rails/comments/fyctoh/migration_depending_on_recently_created_table/)
- url: https://www.reddit.com/r/rails/comments/fyctoh/migration_depending_on_recently_created_table/
---
I have two migrations, a "create users" one:

    class CreateUsers &lt; ActiveRecord::Migration[6.0]
      def change
        create_table :users do |t|
          t.string :email, null: false
          t.string :name, null: false
          t.integer :balance, default: 0
    
          t.timestamps
        end
      end
    end
    

And a "add balance constraint" one:

    class AddBalanceConstraint &lt; ActiveRecord::Migration[6.0]
      def up
        execute 'ALTER TABLE users ADD CONSTRAINT non_negative_balance CHECK(balance &gt;= 0);'
      end
    
      def down
        execute 'ALTER TABLE users DROP CONSTRAINT non_negative_balance;'
      end
    end
    

 This does not work if I run all the migrations from the start, but if I run everything **except** the constraint one, and **then** run it separately, it does work. As if it requires that the table is actually added to postgres. And I know I can validate this in the User model class, but I want it to be safer and have it one the lowest level possible. Any ideas?
## [6][Quick styling questions from a newbie.](https://www.reddit.com/r/rails/comments/fy8ap8/quick_styling_questions_from_a_newbie/)
- url: https://www.reddit.com/r/rails/comments/fy8ap8/quick_styling_questions_from_a_newbie/
---
Seen a few different styling gems from a few different sources of tutorial vids. Wondering, what are the gold standards for styling and how often do you use more than one styling gem at once? 

&amp;#x200B;

So far, I'm liking semantic and TailWind, but I'm wondering if I'm missing out on something else.
## [7][Does this look like a good strategy for Rails model relationships?](https://www.reddit.com/r/rails/comments/fxqy28/does_this_look_like_a_good_strategy_for_rails/)
- url: https://www.reddit.com/r/rails/comments/fxqy28/does_this_look_like_a_good_strategy_for_rails/
---
Here's the ERD: 

&amp;#x200B;

[https://drive.google.com/file/d/18oCL\_OdTcpVzX7EUY6vT6\_PZWOVGH0d2/view?usp=sharing](https://drive.google.com/file/d/18oCL_OdTcpVzX7EUY6vT6_PZWOVGH0d2/view?usp=sharing)

&amp;#x200B;

When a service goes down, an Outage is created. Users (employees) can write notes about the outage. When the service goes back up, the Outage gets an end\_time. A service can have many outages through time, not at once. A single Outage has one Service associated with it. Users can have many Notes per Outage. Is this ERD fine for my purposes or does it violate some basics of db relational rules.
## [8][Code Review: Sidekiq Error Service](https://www.reddit.com/r/rails/comments/fxxh58/code_review_sidekiq_error_service/)
- url: https://www.reddit.com/r/rails/comments/fxxh58/code_review_sidekiq_error_service/
---
Hey everyone,

I've got a simple app that uses Sidekiq. Most of the error handling services on heroku just aren't a great fit so I'm rolling my own. All I need is for the errors to show up in slack in our #ops channel. Then, someone on the team claims the error and we track the resolution in slack.

My big question is if I should use class variables in the `SlackErrorService` to reuse the `request` between messages. 

```
#config/initializers/sidekiq.rb

Sidekiq.configure_server do |config|
  config.error_handlers &lt;&lt; Proc.new {|exception, context_hash| SlackErrorService.notify(exception, context_hash) }
end
```

```
#app/services/slack_error_service.rb

require 'net/http'
require 'uri'
require 'json'

class SlackErrorService
  def self.notify(exception, context_hash)
    uri = URI.parse("#{ENV['SLACK_URI']}")
    request = Net::HTTP::Post.new(uri)

    request.content_type = "application/json"
    request.body = JSON.dump({
      "text" =&gt; exception,
      "blocks" =&gt; [
        {
          "type": "section",
          "block_id": context_hash[:job]["jid"],
          "fields": [
            {
              "type": "mrkdwn",
              "text": "*Context*\n#{context_hash[:context]}"
            },
            {
              "type": "mrkdwn",
              "text": "*Class*\n#{context_hash[:job]['class']}"
            },
            {
              "type": "mrkdwn",
              "text": "*Queue*\n#{context_hash[:job]['queue']}"
            },
            {
              "type": "mrkdwn",
              "text": "*Retry Count*\n#{context_hash[:job]['retry_count']}"
            },
            {
              "type": "mrkdwn",
            "text": "*Error*\n#{context_hash[:job]['error_message']}"
            }
          ]
        }
      ]
    }) 

    req_options = {
      use_ssl: uri.scheme == "https",
    }

    Net::HTTP.start(uri.hostname, uri.port, req_options) do |http|
      http.request(request)
    end
  end
end
```
## [9][Best practices to learn the structure in the big project with rails for new dev](https://www.reddit.com/r/rails/comments/fxrf86/best_practices_to_learn_the_structure_in_the_big/)
- url: https://www.reddit.com/r/rails/comments/fxrf86/best_practices_to_learn_the_structure_in_the_big/
---
I currently new in Ruby and Rails, but I have learn and practice to create some apps , what I do before made the project
Know how rails architecture is and basic of Ruby 
But more focus about modules and classes for this 
This is good for me to make reusable methods to not DRY 
I am in big project with rails now 
I have review some codes of this project and get confused and don't want to scroll down and up when I am on some files with tens of thousands lines LOL because I don't know much the structures of that project and need time to learn it  what I asked

How do you know quickly the methods were called/invoke when we in the page view on it ?
And how do you know this service on that files ? How you do best practices wit Rails console ?
## [10][Can I run a spec file with a different set of inputs in RSpec](https://www.reddit.com/r/rails/comments/fxru5a/can_i_run_a_spec_file_with_a_different_set_of/)
- url: https://www.reddit.com/r/rails/comments/fxru5a/can_i_run_a_spec_file_with_a_different_set_of/
---
Suppose there is a class User for which there are tests like below.  I want to run this test multiple times but with a different set of input. 

Can I do that?

    describe User do 
    all tests
    end

I know inside the test cases I can write a loop and make it run multiple times but I have lots of test cases and this will not be scalable.
## [11][Need help with rake development:seed not working](https://www.reddit.com/r/rails/comments/fxoim4/need_help_with_rake_developmentseed_not_working/)
- url: https://www.reddit.com/r/rails/comments/fxoim4/need_help_with_rake_developmentseed_not_working/
---
I have a rails server running in docker that does graphql API things.  
I used to be able to login to our website but right now it is not working.  
When I check inside whether there are any Users with `rails c`, it returns nil  


So i need to run `rake development:seed`

But this error below appears. the command used to work, and add all the User accounts, with their data in. I am pulling directly from my senior's API on git and it works for them. So I don't think i should touch the API's files itself? but now I am at loss for what to do.

(dunno about rails, I'm just a React developer, really. I couldn't get back to my tasks because of this error... )  


    docker-compose run --rm referem-api bundle exec rake development:seed
    Starting referem-api_postgres_1 ... done
    Cleaning up database...
    Cleaning database completed
    seeding Users Please wait...
    rake aborted!
    ActiveSupport::MessageEncryptor::InvalidMessage: ActiveSupport::MessageEncryptor::InvalidMessage
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/message_encryptor.rb:206:in `rescue in _decrypt'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/message_encryptor.rb:183:in `_decrypt'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/message_encryptor.rb:157:in `decrypt_and_verify'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/messages/rotator.rb:21:in `decrypt_and_verify'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/encrypted_file.rb:79:in `decrypt'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/encrypted_file.rb:42:in `read'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/encrypted_configuration.rb:21:in `read'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/encrypted_configuration.rb:33:in `config'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/encrypted_configuration.rb:38:in `options'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/core_ext/module/delegation.rb:271:in `method_missing'
    (erb):28:in `&lt;main&gt;'
    /usr/local/bundle/gems/activestorage-5.2.4.1/lib/active_storage/engine.rb:95:in `block (2 levels) in &lt;class:Engine&gt;'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/lazy_load_hooks.rb:71:in `instance_eval'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/lazy_load_hooks.rb:71:in `block in execute_hook'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/lazy_load_hooks.rb:62:in `with_execution_control'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/lazy_load_hooks.rb:67:in `execute_hook'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/lazy_load_hooks.rb:52:in `block in run_load_hooks'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/lazy_load_hooks.rb:51:in `each'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/lazy_load_hooks.rb:51:in `run_load_hooks'
    /usr/local/bundle/gems/activestorage-5.2.4.1/app/models/active_storage/blob.rb:235:in `&lt;class:Blob&gt;'
    /usr/local/bundle/gems/activestorage-5.2.4.1/app/models/active_storage/blob.rb:16:in `&lt;main&gt;'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:54:in `load'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:54:in `load'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:476:in `block in load_file'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:661:in `new_constants_in'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:475:in `load_file'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:373:in `block in require_or_load'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:37:in `block in load_interlock'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies/interlock.rb:14:in `block in loading'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/concurrency/share_lock.rb:151:in `exclusive'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies/interlock.rb:13:in `loading'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:37:in `load_interlock'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:356:in `require_or_load'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/active_support.rb:48:in `block in require_or_load'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/active_support.rb:16:in `allow_bootsnap_retry'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/active_support.rb:47:in `require_or_load'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:510:in `load_missing_constant'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/active_support.rb:60:in `block in load_missing_constant'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/active_support.rb:16:in `allow_bootsnap_retry'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/active_support.rb:59:in `load_missing_constant'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:195:in `const_missing'
    /usr/local/bundle/gems/activestorage-5.2.4.1/lib/active_storage/attached.rb:20:in `create_blob_from'
    /usr/local/bundle/gems/activestorage-5.2.4.1/lib/active_storage/attached/one.rb:24:in `attach'
    /referem-api/app/models/seller.rb:65:in `avatar_default'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:426:in `block in make_lambda'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:198:in `block (2 levels) in halting'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:606:in `block (2 levels) in default_terminator'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:605:in `catch'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:605:in `block in default_terminator'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:199:in `block in halting'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:513:in `block in invoke_before'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:513:in `each'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:513:in `invoke_before'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:131:in `run_callbacks'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:816:in `_run_create_callbacks'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/callbacks.rb:346:in `_create_record'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/timestamp.rb:102:in `_create_record'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/persistence.rb:705:in `create_or_update'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/callbacks.rb:342:in `block in create_or_update'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:132:in `run_callbacks'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:816:in `_run_save_callbacks'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/callbacks.rb:342:in `create_or_update'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/persistence.rb:308:in `save!'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/validations.rb:52:in `save!'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/transactions.rb:315:in `block in save!'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/transactions.rb:387:in `block in with_transaction_returning_status'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/connection_adapters/abstract/database_statements.rb:267:in `block in transaction'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/connection_adapters/abstract/transaction.rb:239:in `block in within_new_transaction'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/connection_adapters/abstract/transaction.rb:236:in `within_new_transaction'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/connection_adapters/abstract/database_statements.rb:267:in `transaction'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/transactions.rb:212:in `transaction'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/transactions.rb:385:in `with_transaction_returning_status'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/transactions.rb:315:in `save!'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/suppressor.rb:48:in `save!'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/persistence.rb:53:in `create!'
    /referem-api/db/seeds/users.rb:14:in `&lt;main&gt;'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:54:in `load'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:54:in `load'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:285:in `block in load'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:257:in `load_dependency'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:285:in `load'
    /referem-api/lib/tasks/development.rake:16:in `block (3 levels) in &lt;main&gt;'
    /referem-api/lib/tasks/development.rake:12:in `each'
    /referem-api/lib/tasks/development.rake:12:in `block (2 levels) in &lt;main&gt;'
    /usr/local/bundle/gems/rake-13.0.1/exe/rake:27:in `&lt;top (required)&gt;'
    /usr/local/bin/bundle:23:in `load'
    /usr/local/bin/bundle:23:in `&lt;main&gt;'
    
    Caused by:
    ArgumentError: key must be 16 bytes
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/message_encryptor.rb:193:in `key='
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/message_encryptor.rb:193:in `_decrypt'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/message_encryptor.rb:157:in `decrypt_and_verify'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/messages/rotator.rb:21:in `decrypt_and_verify'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/encrypted_file.rb:79:in `decrypt'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/encrypted_file.rb:42:in `read'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/encrypted_configuration.rb:21:in `read'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/encrypted_configuration.rb:33:in `config'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/encrypted_configuration.rb:38:in `options'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/core_ext/module/delegation.rb:271:in `method_missing'
    (erb):28:in `&lt;main&gt;'
    /usr/local/bundle/gems/activestorage-5.2.4.1/lib/active_storage/engine.rb:95:in `block (2 levels) in &lt;class:Engine&gt;'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/lazy_load_hooks.rb:71:in `instance_eval'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/lazy_load_hooks.rb:71:in `block in execute_hook'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/lazy_load_hooks.rb:62:in `with_execution_control'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/lazy_load_hooks.rb:67:in `execute_hook'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/lazy_load_hooks.rb:52:in `block in run_load_hooks'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/lazy_load_hooks.rb:51:in `each'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/lazy_load_hooks.rb:51:in `run_load_hooks'
    /usr/local/bundle/gems/activestorage-5.2.4.1/app/models/active_storage/blob.rb:235:in `&lt;class:Blob&gt;'
    /usr/local/bundle/gems/activestorage-5.2.4.1/app/models/active_storage/blob.rb:16:in `&lt;main&gt;'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:54:in `load'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:54:in `load'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:476:in `block in load_file'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:661:in `new_constants_in'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:475:in `load_file'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:373:in `block in require_or_load'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:37:in `block in load_interlock'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies/interlock.rb:14:in `block in loading'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/concurrency/share_lock.rb:151:in `exclusive'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies/interlock.rb:13:in `loading'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:37:in `load_interlock'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:356:in `require_or_load'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/active_support.rb:48:in `block in require_or_load'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/active_support.rb:16:in `allow_bootsnap_retry'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/active_support.rb:47:in `require_or_load'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:510:in `load_missing_constant'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/active_support.rb:60:in `block in load_missing_constant'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/active_support.rb:16:in `allow_bootsnap_retry'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/active_support.rb:59:in `load_missing_constant'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:195:in `const_missing'
    /usr/local/bundle/gems/activestorage-5.2.4.1/lib/active_storage/attached.rb:20:in `create_blob_from'
    /usr/local/bundle/gems/activestorage-5.2.4.1/lib/active_storage/attached/one.rb:24:in `attach'
    /referem-api/app/models/seller.rb:65:in `avatar_default'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:426:in `block in make_lambda'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:198:in `block (2 levels) in halting'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:606:in `block (2 levels) in default_terminator'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:605:in `catch'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:605:in `block in default_terminator'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:199:in `block in halting'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:513:in `block in invoke_before'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:513:in `each'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:513:in `invoke_before'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:131:in `run_callbacks'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:816:in `_run_create_callbacks'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/callbacks.rb:346:in `_create_record'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/timestamp.rb:102:in `_create_record'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/persistence.rb:705:in `create_or_update'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/callbacks.rb:342:in `block in create_or_update'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:132:in `run_callbacks'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/callbacks.rb:816:in `_run_save_callbacks'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/callbacks.rb:342:in `create_or_update'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/persistence.rb:308:in `save!'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/validations.rb:52:in `save!'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/transactions.rb:315:in `block in save!'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/transactions.rb:387:in `block in with_transaction_returning_status'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/connection_adapters/abstract/database_statements.rb:267:in `block in transaction'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/connection_adapters/abstract/transaction.rb:239:in `block in within_new_transaction'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/connection_adapters/abstract/transaction.rb:236:in `within_new_transaction'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/connection_adapters/abstract/database_statements.rb:267:in `transaction'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/transactions.rb:212:in `transaction'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/transactions.rb:385:in `with_transaction_returning_status'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/transactions.rb:315:in `save!'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/suppressor.rb:48:in `save!'
    /usr/local/bundle/gems/activerecord-5.2.4.1/lib/active_record/persistence.rb:53:in `create!'
    /referem-api/db/seeds/users.rb:14:in `&lt;main&gt;'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:54:in `load'
    /usr/local/bundle/gems/bootsnap-1.4.5/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:54:in `load'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:285:in `block in load'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:257:in `load_dependency'
    /usr/local/bundle/gems/activesupport-5.2.4.1/lib/active_support/dependencies.rb:285:in `load'
    /referem-api/lib/tasks/development.rake:16:in `block (3 levels) in &lt;main&gt;'
    /referem-api/lib/tasks/development.rake:12:in `each'
    /referem-api/lib/tasks/development.rake:12:in `block (2 levels) in &lt;main&gt;'
    /usr/local/bundle/gems/rake-13.0.1/exe/rake:27:in `&lt;top (required)&gt;'
    /usr/local/bin/bundle:23:in `load'
    /usr/local/bin/bundle:23:in `&lt;main&gt;'
    Tasks: TOP =&gt; development:seed
    (See full trace by running task with --trace)
## [12][Problems installing sqlite3](https://www.reddit.com/r/rails/comments/fxctvq/problems_installing_sqlite3/)
- url: https://www.reddit.com/r/rails/comments/fxctvq/problems_installing_sqlite3/
---
 I just created a new project. It's just a blog and I'm following a tutorial on Medium. Just going through the process of setting up the project and when trying to set up the rails server, it's having trouble installing sqlite3. The Gemfile that was automatically created when I created the project had the gem defined like this automatically: gem 'sqlite3', '1.4.2' (I think, while trying to fix it, I've changed it).

I've tried just having gem 'sqlite3' with no version next to it. I've tried doing older versions. It just gives me an error saying that the bundler can't continue. Can someone help me with it? I'll copy and paste the whole output from the terminal below.

\[ENV\]:/vagrant/src/blog $ bundle install

The dependency tzinfo-data (&gt;= 0) will be unused by any of the platforms Bundler is installing for. Bundler is installing for ruby but the dependency is only for x86-mingw32, x86-mswin32, x64-mingw32, java. To add those platforms to the bundle, run \`bundle lock --add-platform x86-mingw32 x86-mswin32 x64-mingw32 java\`.

Fetching gem metadata from [https://rubygems.org/](https://rubygems.org/)............

Fetching gem metadata from [https://rubygems.org/](https://rubygems.org/).

Resolving dependencies...

Using rake 13.0.1

Using concurrent-ruby 1.1.6

Using i18n 1.8.2

Using minitest 5.14.0

Using thread\_safe 0.3.6

Using tzinfo 1.2.7

Using zeitwerk 2.3.0

Using activesupport [6.0.2.2](https://6.0.2.2/)

Using builder 3.2.4

Using erubi 1.9.0

Using mini\_portile2 2.4.0

Using nokogiri 1.10.9

Using rails-dom-testing 2.0.3

Using crass 1.0.6

Using loofah 2.5.0

Using rails-html-sanitizer 1.3.0

Using actionview [6.0.2.2](https://6.0.2.2/)

Using rack 2.2.2

Using rack-test 1.1.0

Using actionpack [6.0.2.2](https://6.0.2.2/)

Using nio4r 2.5.2

Using websocket-extensions 0.1.4

Using websocket-driver 0.7.1

Using actioncable [6.0.2.2](https://6.0.2.2/)

Using globalid 0.4.2

Using activejob [6.0.2.2](https://6.0.2.2/)

Using activemodel [6.0.2.2](https://6.0.2.2/)

Using activerecord [6.0.2.2](https://6.0.2.2/)

Using mimemagic 0.3.4

Using marcel 0.3.3

Using activestorage [6.0.2.2](https://6.0.2.2/)

Using mini\_mime 1.0.2

Using mail 2.7.1

Using actionmailbox [6.0.2.2](https://6.0.2.2/)

Using actionmailer [6.0.2.2](https://6.0.2.2/)

Using actiontext [6.0.2.2](https://6.0.2.2/)

Using public\_suffix 4.0.4

Using addressable 2.7.0

Using bindex 0.8.1

Using msgpack 1.3.3

Using bootsnap 1.4.6

Using bundler 1.17.1

Using byebug 11.1.1

Using regexp\_parser 1.7.0

Using xpath 3.2.0

Using capybara 3.32.1

Using childprocess 3.0.0

Using ffi 1.12.2

Using jbuilder 2.10.0

Using rb-fsevent 0.10.3

Using rb-inotify 0.10.1

Using ruby\_dep 1.5.0

Using listen 3.1.5

Using method\_source 1.0.0

Using puma 4.3.3

Using rack-proxy 0.6.5

Using thor 1.0.1

Using railties [6.0.2.2](https://6.0.2.2/)

Using sprockets 4.0.0

Using sprockets-rails 3.2.1

Using rails [6.0.2.2](https://6.0.2.2/)

Using rubyzip 2.3.0

Using sassc 2.2.1

Using tilt 2.0.10

Using sassc-rails 2.1.2

Using sass-rails 6.0.0

Using selenium-webdriver 3.142.7

Using spring 2.1.0

Using spring-watcher-listen 2.0.1

Fetching sqlite3 1.4.2

Installing sqlite3 1.4.2 with native extensions

Gem::Ext::BuildError: ERROR: Failed to build gem native extension.

current directory:

/home/vagrant/.rbenv/versions/2.5.3/lib/ruby/gems/2.5.0/gems/sqlite3-1.4.2/ext/sqlite3

/home/vagrant/.rbenv/versions/2.5.3/bin/ruby -r ./siteconf20200408-7495-j5nn4o.rb

extconf.rb

checking for sqlite3.h... no

sqlite3.h is missing. Try 'brew install sqlite3',

'yum install sqlite-devel' or 'apt-get install libsqlite3-dev'

and check your shared library search path (the

location where your sqlite3 shared library is located).

\*\*\* extconf.rb failed \*\*\*

Could not create Makefile due to some reason, probably lack of necessary

libraries and/or headers. Check the mkmf.log file for more details. You may

need configuration options.

Provided configuration options:

\--with-opt-dir

\--without-opt-dir

\--with-opt-include

\--without-opt-include=${opt-dir}/include

\--with-opt-lib

\--without-opt-lib=${opt-dir}/lib

\--with-make-prog

\--without-make-prog

\--srcdir=.

\--curdir

\--ruby=/home/vagrant/.rbenv/versions/2.5.3/bin/$(RUBY\_BASE\_NAME)

\--with-sqlcipher

\--without-sqlcipher

\--with-sqlite3-config

\--without-sqlite3-config

\--with-pkg-config

\--without-pkg-config

\--with-sqlcipher

\--without-sqlcipher

\--with-sqlite3-dir

\--without-sqlite3-dir

\--with-sqlite3-include

\--without-sqlite3-include=${sqlite3-dir}/include

\--with-sqlite3-lib

\--without-sqlite3-lib=${sqlite3-dir}/lib

To see why this extension failed to compile, please check the mkmf.log which can be

found here:

/home/vagrant/.rbenv/versions/2.5.3/lib/ruby/gems/2.5.0/extensions/x86\_64-linux/2.5.0-static/sqlite3-1.4.2/mkmf.log

extconf failed, exit code 1

Gem files will remain installed in

/home/vagrant/.rbenv/versions/2.5.3/lib/ruby/gems/2.5.0/gems/sqlite3-1.4.2 for

inspection.

Results logged to

/home/vagrant/.rbenv/versions/2.5.3/lib/ruby/gems/2.5.0/extensions/x86\_64-linux/2.5.0-static/sqlite3-1.4.2/gem\_make.out

An error occurred while installing sqlite3 (1.4.2), and Bundler cannot

continue.

Make sure that \`gem install sqlite3 -v '1.4.2' --source '[https://rubygems.org/'\`](https://rubygems.org/'%60)

succeeds before bundling.

In Gemfile:

sqlite3
