# ruby
## [1][How we feel about ruby](https://www.reddit.com/r/ruby/comments/hcuv7q/how_we_feel_about_ruby/)
- url: https://v.redd.it/wbby6pa035651
---

## [2][An efficient way of downloading multiple images and send it to s3 ?](https://www.reddit.com/r/ruby/comments/hd37ep/an_efficient_way_of_downloading_multiple_images/)
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
## [3][Open Source Progress Report - Event loops, documentation, source code analysis, Markdown &amp; turbo tests - April &amp; May 2020](https://www.reddit.com/r/ruby/comments/hcwuey/open_source_progress_report_event_loops/)
- url: https://www.codeotaku.com/journal/2020-06/open-source-progress-report/index
---

## [4][How to use React components in Rails 6](https://www.reddit.com/r/ruby/comments/hd205i/how_to_use_react_components_in_rails_6/)
- url: https://medium.com/the-side-hustler/how-to-use-react-components-in-rails-6-7ef460894be
---

## [5][Ruby on Rails multitenancy in 2020](https://www.reddit.com/r/ruby/comments/hcntnf/ruby_on_rails_multitenancy_in_2020/)
- url: https://realptsdengineer.com/ruby-on-rails-multitenancy-in-2020/
---

## [6][Glimmer DSL for Opal (Web GUI Adapter for Desktop Apps)](https://www.reddit.com/r/ruby/comments/hcs7s8/glimmer_dsl_for_opal_web_gui_adapter_for_desktop/)
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
## [7][Practical resources to learn about ruby and rails security topics](https://www.reddit.com/r/ruby/comments/hcudz3/practical_resources_to_learn_about_ruby_and_rails/)
- url: https://www.reddit.com/r/ruby/comments/hcudz3/practical_resources_to_learn_about_ruby_and_rails/
---
With my new job, I'm diving into ruby/rails at the moment and it's a great place to be I think :)

I'd love to align my learning along with security topics. I have a basic understanding of security, having only implemented a few authentication systems so far (1 or 2 from scratch in simple ROR and Flask projects; with and without amplify in AWS serverless apps; and a few others with 3rd party authentication services).

So I'm looking for resources to better understand and practice relevant security topics, ideally within a ruby/rails context.

**My little research gave me this so far:**

* [https://github.com/pxlpnk/awesome-ruby-security#educational](https://github.com/pxlpnk/awesome-ruby-security#educational) 

looks like it has the fundamental theory covered. Also, there are a few "insecure playground projects" which sounds very practical! But I wouldn't have much of an idea where to start exploiting the projects...

* Found a few tips in this Ruby Security ["handbook"](https://www.sqreen.com/resources/ruby-security-handbook?utm_medium=Paid_Search&amp;utm_source=Google&amp;utm_campaign=Search_Google_EMEA_Non-Brand_ACQ_Phrase&amp;utm_content=Non-Brand_KWs&amp;gclid=CjwKCAjw57b3BRBlEiwA1ImytujMmZYSFBKv7cMfaN6-gXwDrQvqRu32JkKwxa6Yd_BWInlHQlN2qhoCXDcQAvD_BwE). Theoretical but probably actionable.
* "[A week with rails security strategy](https://bauland42.com/a-week-with-a-rails-security-strategy/)" and the according newsletter by the good person who wrote the [rails security guide](https://guides.rubyonrails.org/security.html). 

Do you know of any **practical resources or courses** where I can dig deeper into and maybe practice the fundamental security topics? 

Or would you recommend some of the resources above and somehow practice it on my own applications?
## [8][Library to write, in pure Ruby, web apps that are automagically accessible from all the Internet (link to source code and online demos in comment)](https://www.reddit.com/r/ruby/comments/hch8e6/library_to_write_in_pure_ruby_web_apps_that_are/)
- url: https://i.redd.it/n9wu9gsrd0651.gif
---

## [9][Ruby Notes for Professionals "by the beautiful people at Stack Overflow"](https://www.reddit.com/r/ruby/comments/hc8btp/ruby_notes_for_professionals_by_the_beautiful/)
- url: https://www.dbooks.org/ruby-notes-for-professionals-5592543428/read/
---

## [10][Rails Vs Sinatra](https://www.reddit.com/r/ruby/comments/hcgsb7/rails_vs_sinatra/)
- url: https://www.reddit.com/r/ruby/comments/hcgsb7/rails_vs_sinatra/
---
Hey all!

I'm a long time ruby dev who's never gone into rails or Sinatra(mostly been on the Ops side), I've read some very old threads explaining that as people have scaled they've moved away from Sinatra into rails, but none of them actually go into much detail. Does anyone here have experience with reasons to migrate to rails? Thus far I've had zero problems with Sinatra so I intend on using it but am wondering what I may run into down the line.

Appreciate any advice!
