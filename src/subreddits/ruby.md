# ruby
## [1][Webpacker 5.0 released](https://www.reddit.com/r/ruby/comments/fp6r12/webpacker_50_released/)
- url: https://prathamesh.tech/2020/03/25/webpacker-5-0-released/
---

## [2][Glimmer (JRuby Desktop UI DSL + Data-Binding)](https://www.reddit.com/r/ruby/comments/fp62o1/glimmer_jruby_desktop_ui_dsl_databinding/)
- url: https://www.reddit.com/r/ruby/comments/fp62o1/glimmer_jruby_desktop_ui_dsl_databinding/
---
[https://github.com/AndyObtiva/glimmer](https://github.com/AndyObtiva/glimmer)

Glimmer (JRuby Desktop UI DSL + Data-Binding) has been undergoing new development.

Examples:

https://preview.redd.it/56f1oatgqyo41.png?width=135&amp;format=png&amp;auto=webp&amp;s=00d9663fd6bc51de0652b892e00d68d8aa8d5d7c

    include Glimmer
    
    shell {
      text "Glimmer"
      label {
        text "Hello World!"
      }
    }.open

&amp;#x200B;

https://preview.redd.it/022q03llqyo41.png?width=159&amp;format=png&amp;auto=webp&amp;s=f2c23d985d7497c32c19534e9950fa8de6572be0

    shell {
      text "Tic-Tac-Toe"
      composite {
        grid_layout 3, true
        (1..3).each { |row|
          (1..3).each { |column|
            button {
              layout_data :fill, :fill, true, true
              text        bind(@tic_tac_toe_board[row, column], :sign)
              enabled     bind(@tic_tac_toe_board[row, column], :empty)
              on_widget_selected {
                @tic_tac_toe_board.mark_box(row, column)
              }
            }
          }
        }
      }
    }

Recent changes/additions:
- Nested/indexed property data-binding
- SWT layout and layout data DSL
- SWT Color and Font DSL
- Ability to add Glimmer Custom Widgets
- SWT Browser widget support
- Automatic cleanup of data-binding/observers upon disposing a widget
- Easier SWT style syntax via GSWT class
- Inclusion of SWT library jars in Ruby gem
- Improved glimmer/girb commands for running on Windows/Linux/Mac
## [3][Why and How to Host your Rails 6 App with AWS ElasticBeanstalk and RDS](https://www.reddit.com/r/ruby/comments/fowm39/why_and_how_to_host_your_rails_6_app_with_aws/)
- url: https://www.reddit.com/r/ruby/comments/fowm39/why_and_how_to_host_your_rails_6_app_with_aws/
---
When you deploy a new Rails app, you typically face a double-bind. If you use an easy platform like Heroku, you could create problems for yourself as your application scales. If you use a more fully-featured platform, you risk wasting time on ops that could be spent on your product. What if you could have both: an easy deployment option that is easy to scale? In this article, Amos Omondi argues that AWS Elastic Beanstalk gives us both, then he shows us everything we need to know to get a Rails 6 app up and running on EB. [https://www.honeybadger.io/blog/rails-6-aws-elastic-beanstalk/](https://www.honeybadger.io/blog/rails-6-aws-elastic-beanstalk/)
## [4][Sending data from Ruby to Frontend without Rails?](https://www.reddit.com/r/ruby/comments/fp3djx/sending_data_from_ruby_to_frontend_without_rails/)
- url: https://www.reddit.com/r/ruby/comments/fp3djx/sending_data_from_ruby_to_frontend_without_rails/
---
I was just wondering if anyone had any resources or references on how one would use just Ruby without a framework such as rails to use as a backend and pass data to the front end?
## [5][Improving Net::HTTP Concurrency by Samuel Williams](https://www.reddit.com/r/ruby/comments/foos7v/improving_nethttp_concurrency_by_samuel_williams/)
- url: https://youtube.com/watch?v=uU8ziRoJ2Z8&amp;feature=youtu.be
---

## [6][Setup Ruby on Rails Active Storage for Transfer Accelerated S3](https://www.reddit.com/r/ruby/comments/fp5jln/setup_ruby_on_rails_active_storage_for_transfer/)
- url: https://www.reddit.com/r/ruby/comments/fp5jln/setup_ruby_on_rails_active_storage_for_transfer/
---
Transfer Accelerated S3 bucket allows you to upload/download faster. The setup should be simple in theory:

The endpoint changes from:

`mybucket.s3.us-east-1.amazonaws.com`

to:

`mybucket.s3-accelerate.amazonaws.com`

In Ruby on Rails config/storage.yml the environment variables look like this:

`amazon:`

`service: S3`

`access_key_id: &lt;%= ENV['AWS_ACCESS_KEY_ID'] %&gt;`

`secret_access_key: &lt;%= ENV['AWS_SECRET_ACCESS_KEY'] %&gt;`

`region: &lt;%= ENV['AWS_REGION'] %&gt;`

`bucket: &lt;%= ENV['AWS_BUCKET'] %&gt;`

The problem is the pattern is different, so I can't just change `AWS_REGION`. There is an extra `.s3` in there. Do you know how to implement transfer acceleration?
## [7][2 word file](https://www.reddit.com/r/ruby/comments/fp0df0/2_word_file/)
- url: https://www.reddit.com/r/ruby/comments/fp0df0/2_word_file/
---
Having some trouble running 2 worded .rb files. I am using the GNOME text editor and running the code from the Linux terminal. Does anyone know why it refuses to run? Sorry for the beginner question! 
## [8][render_async Fires Events By Default](https://www.reddit.com/r/ruby/comments/fosp1m/render_async_fires_events_by_default/)
- url: https://pragmaticpineapple.com/render-async-fires-events-by-default/
---

## [9][The return Keyword in Ruby](https://www.reddit.com/r/ruby/comments/foq6ww/the_return_keyword_in_ruby/)
- url: https://medium.com/rubycademy/the-return-keyword-in-ruby-df0a7f578fcb
---

## [10][Why does Hash.new(0) allow for unique values but Hash.new("some text") makes all keys' values point to the same place in memory?](https://www.reddit.com/r/ruby/comments/fod8gx/why_does_hashnew0_allow_for_unique_values_but/)
- url: https://www.reddit.com/r/ruby/comments/fod8gx/why_does_hashnew0_allow_for_unique_values_but/
---
To illustrate my question:


```
h = Hash.new("some text")
h["a"].upcase!
h["b"] # =&gt; "SOME TEXT"
```

Modifying one key's value will apply to all other keys' values because it's the same place in memory. That's why we use `h = Hash.new { |hash, key| hash[key] = "some text: #{key}" }` for default values that are unique to each key. However,

```
h = Hash.new(0)
h["a"] += 1
h["b"] # =&gt; 0
```

I didn't have to pass the block `{ |hash, key| hash[key] = ... }` into `Hash.new` and yet `h["b"]` remained as 0 while `h["a"]` changed. After seeing the behavior of `Hash.new("some text")` I would expect `h["b"]` to also be `1`.

Thanks for any explanations or pointing out any holes in my understanding!
