# ruby
## [1][Stack Overflow Developer Survey 2020](https://www.reddit.com/r/ruby/comments/gs2zjc/stack_overflow_developer_survey_2020/)
- url: https://insights.stackoverflow.com/survey/2020#technology-how-technologies-are-connected
---

## [2][Rails 6.1 adds support for signed ids to Active Record](https://www.reddit.com/r/ruby/comments/grmymc/rails_61_adds_support_for_signed_ids_to_active/)
- url: https://blog.saeloun.com/2020/05/20/rails-6-1-adds-support-for-signed-ids-to-active-record.html
---

## [3][The Ruby Blend Episode 14: StimulusReflex, BridgetownRB, RailsBytes, AppLocale.dev, and more!](https://www.reddit.com/r/ruby/comments/gs65w5/the_ruby_blend_episode_14_stimulusreflex/)
- url: https://fireside.fm/s/ouBAUjGy+7QL69WzH
---

## [4][How To Prevent ERP Implementation Failure](https://www.reddit.com/r/ruby/comments/gs4qcb/how_to_prevent_erp_implementation_failure/)
- url: https://www.odooimplementation.com/blog/erp_implementation_failure
---

## [5][How we achieved 2× faster application run with only 1/3 of the servers by tuning auto scaling rules and switching to puma threads](https://www.reddit.com/r/ruby/comments/grnbb4/how_we_achieved_2_faster_application_run_with/)
- url: https://bytes.babbel.com/en/articles/2020-05-27-how-to-do-more-with-fewer-servers.html
---

## [6][Kaminari security update](https://www.reddit.com/r/ruby/comments/grswlf/kaminari_security_update/)
- url: https://www.reddit.com/r/ruby/comments/grswlf/kaminari_security_update/
---
I discovered a security issue in the kaminari gem. I would advise you all to update to the latest version.
## [7][How to Add jQuery in Rails 6 Using Webpacker](https://www.reddit.com/r/ruby/comments/gs1689/how_to_add_jquery_in_rails_6_using_webpacker/)
- url: https://youtu.be/dCjI0ZyvThA
---

## [8][Two Commonly Used Rails Upgrade Strategies](https://www.reddit.com/r/ruby/comments/grni7g/two_commonly_used_rails_upgrade_strategies/)
- url: https://www.fastruby.io/blog/rails/upgrades/rails-upgrade-strategies.html
---

## [9][How to securely merge action controller parameters](https://www.reddit.com/r/ruby/comments/grrko7/how_to_securely_merge_action_controller_parameters/)
- url: https://www.youtube.com/watch?v=SDe_AX6hAsU
---

## [10][Client facing token replacements](https://www.reddit.com/r/ruby/comments/grjxu9/client_facing_token_replacements/)
- url: https://www.reddit.com/r/ruby/comments/grjxu9/client_facing_token_replacements/
---
Hey there!

First of all - the story below is super simplified version of what is really needed. I simplified it to make it easier to understand what i need and to make this post shorter.

I'm developing an app, in which user can upload a CSV file (there is no any standard format of it) and records imported this way are saved in a database as a \`User\` record.

User of that system is able to configure which csv columns should be matched to which database columns of a record. Right now, i'm just storing such config as a key-value hash in a following format:  


    {
      'db_column_1' =&gt; 'csv_column_1',
      'db_column_2' =&gt; 'csv_column_2',
    }

as it works pretty well and is really simple, i need to extend that by adding a feature which would let user define a more complicated mappings, like multiple CSV columns put in the same record column. What first came to my mind was using jinja/handlebars-like engine so config hash would look as follows:  


    {
      'db_column_1' =&gt; '{{ csv_column_1 }}',
      'db_column_2' =&gt; '{{ csv_column_2 }} {{ csv_column_3 }}',
    }

Using jinja/handlebars would be an obvious overkill as a regular \`gsub\` will be enough but, there is one more thing i need - filters. (as an example) sometimes CSV has a timestamp column and i would need some sort of modifier feature to make it able to convert incoming values (pseudocode):

    {
      'db_column_1' =&gt; '{{ csv_column_1 | timestamp_to_datetime }}'
    }

Given eveyrthing above - is there any go-to and well known lib for doing such smart "token" replacements in a handlebars/jinja style? I feel like jinja/handlebars is an overkill, especially that sometimes such CSV file has 60k (!!) records in it.  


Thanks in advance for any clues!
