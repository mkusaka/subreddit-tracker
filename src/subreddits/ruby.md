# ruby
## [1][ruby on arm?](https://www.reddit.com/r/ruby/comments/hdnmn9/ruby_on_arm/)
- url: https://www.reddit.com/r/ruby/comments/hdnmn9/ruby_on_arm/
---
given that mac will move to arm procs
aws has his on arm processor
and somehow the performance is good(https://www.anandtech.com/show/15578/cloud-clash-amazon-graviton2-arm-against-intel-and-amd/10)

- what is the status of ruby on arm?
- can ruby leverage here, or again we are years behind all?
- what the community could do to improve it?
- what is the story with jruby, truffleruby on arm?
## [2][Graceful Request Retries in Ruby Applications](https://www.reddit.com/r/ruby/comments/hdr91t/graceful_request_retries_in_ruby_applications/)
- url: https://medium.com/@kirill_shevch/graceful-request-retries-in-ruby-applications-7bbeac5ebd40
---

## [3][JRuby](https://www.reddit.com/r/ruby/comments/hdd1dw/jruby/)
- url: https://www.reddit.com/r/ruby/comments/hdd1dw/jruby/
---
So if I interpreted this correctly (no pun intended) I can use JRuby to write both Java and Ruby code?
## [4][What are People Using for Localizing Date/Time?](https://www.reddit.com/r/ruby/comments/hdkw5t/what_are_people_using_for_localizing_datetime/)
- url: /r/rails/comments/hdkurj/what_are_people_using_for_localizing_datetime/
---

## [5][How we feel about ruby](https://www.reddit.com/r/ruby/comments/hcuv7q/how_we_feel_about_ruby/)
- url: https://v.redd.it/wbby6pa035651
---

## [6][An efficient way of downloading multiple images and send it to s3 ?](https://www.reddit.com/r/ruby/comments/hd37ep/an_efficient_way_of_downloading_multiple_images/)
- url: https://www.reddit.com/r/ruby/comments/hd37ep/an_efficient_way_of_downloading_multiple_images/
---
I have a json file

&amp;#x200B;

    {
      "album1": {
        "image1": "url",
        "image2": "url",
        "image3": "url",
        "image4": "url"
      }
    }

I need to save those images to S3.   


I am thinking of looping over the hash, then download and send it to s3 and repeat. This task will happen in the background in sidekiq.  


The problem is, this approach is very slow and I am wondering if there are better ways to do it
## [7][How to use React components in Rails 6](https://www.reddit.com/r/ruby/comments/hd205i/how_to_use_react_components_in_rails_6/)
- url: https://medium.com/the-side-hustler/how-to-use-react-components-in-rails-6-7ef460894be
---

## [8][Open Source Progress Report - Event loops, documentation, source code analysis, Markdown &amp; turbo tests - April &amp; May 2020](https://www.reddit.com/r/ruby/comments/hcwuey/open_source_progress_report_event_loops/)
- url: https://www.codeotaku.com/journal/2020-06/open-source-progress-report/index
---

## [9][Ruby on Rails multitenancy in 2020](https://www.reddit.com/r/ruby/comments/hcntnf/ruby_on_rails_multitenancy_in_2020/)
- url: https://realptsdengineer.com/ruby-on-rails-multitenancy-in-2020/
---

## [10][Glimmer DSL for Opal (Web GUI Adapter for Desktop Apps)](https://www.reddit.com/r/ruby/comments/hcs7s8/glimmer_dsl_for_opal_web_gui_adapter_for_desktop/)
- url: https://www.reddit.com/r/ruby/comments/hcs7s8/glimmer_dsl_for_opal_web_gui_adapter_for_desktop/
---
[https://github.com/AndyObtiva/glimmer-dsl-opal](https://github.com/AndyObtiva/glimmer-dsl-opal)

[Glimmer DSL for Opal](https://github.com/AndyObtiva/glimmer-dsl-opal) is an early alpha experimental gem for webifying desktop apps built with [Glimmer](https://github.com/AndyObtiva/glimmer) without changing a line of code. Just switch the \`glimmer-dsl-swt\` gem with \`glimmer-dsl-opal\` and your desktop app can now run in Rails with [Opal](https://opalrb.com/). Glimmer Opal apps may then be custom-styled via standard CSS.

    include Glimmer
       
    shell {
      text 'Glimmer'
      label {
        text 'Hello, World!'
      }
    }

&amp;#x200B;

[Hello, World! glimmer-dsl-swt](https://preview.redd.it/c58vxljk94651.png?width=132&amp;format=png&amp;auto=webp&amp;s=358371e622381f6f1bca1e0429fc6b0f3a902836)

&amp;#x200B;

[Hello, World! glimmer-dsl-opal](https://preview.redd.it/r5d1ihyl94651.png?width=501&amp;format=png&amp;auto=webp&amp;s=46a341eac6fde48d55e839c69f26408c3248a201)

See more examples like [computed data-binding](https://github.com/AndyObtiva/glimmer-dsl-opal#hello-computed) and [Tic Tac Toe](https://github.com/AndyObtiva/glimmer-dsl-opal#tic-tac-toe) at: [https://github.com/AndyObtiva/glimmer-dsl-opal](https://github.com/AndyObtiva/glimmer-dsl-opal)
