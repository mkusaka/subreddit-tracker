# ruby
## [1][Improving Net::HTTP Concurrency by Samuel Williams](https://www.reddit.com/r/ruby/comments/foos7v/improving_nethttp_concurrency_by_samuel_williams/)
- url: https://youtube.com/watch?v=uU8ziRoJ2Z8&amp;feature=youtu.be
---

## [2][Why does Hash.new(0) allow for unique values but Hash.new("some text") makes all keys' values point to the same place in memory?](https://www.reddit.com/r/ruby/comments/fod8gx/why_does_hashnew0_allow_for_unique_values_but/)
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
## [3][The return Keyword in Ruby](https://www.reddit.com/r/ruby/comments/foq6ww/the_return_keyword_in_ruby/)
- url: https://medium.com/rubycademy/the-return-keyword-in-ruby-df0a7f578fcb
---

## [4][RuboCoping with legacy: Bring your Ruby code up to Standard](https://www.reddit.com/r/ruby/comments/fo57nt/rubocoping_with_legacy_bring_your_ruby_code_up_to/)
- url: https://evilmartians.com/chronicles/rubocoping-with-legacy-bring-your-ruby-code-up-to-standard
---

## [5][Six-month release cadence is moving @OpenJDK forward at a fast pace.](https://www.reddit.com/r/ruby/comments/fobwg8/sixmonth_release_cadence_is_moving_openjdk/)
- url: https://twitter.com/bellsoftware/status/1240762836042878977
---

## [6][How to Develop and Debug Ruby Applications in Kubernetes](https://www.reddit.com/r/ruby/comments/fo2w02/how_to_develop_and_debug_ruby_applications_in/)
- url: https://www.reddit.com/r/ruby/comments/fo2w02/how_to_develop_and_debug_ruby_applications_in/
---
Develop your Ruby applications in Kubernetes with instant hot reload as you type code and using your favorite IDE debugger:

[https://okteto.com/blog/how-to-develop-ruby-apps-in-kubernetes/](https://okteto.com/blog/how-to-develop-ruby-apps-in-kubernetes/)
## [7][When Ruby 2.7.1 is coming out?](https://www.reddit.com/r/ruby/comments/foehws/when_ruby_271_is_coming_out/)
- url: https://www.reddit.com/r/ruby/comments/foehws/when_ruby_271_is_coming_out/
---

## [8][So I have started planning for this project and there is a lot I do not know. For starters, is there a tool I can use to get system info like memory or CPU usage?](https://www.reddit.com/r/ruby/comments/fo4gk4/so_i_have_started_planning_for_this_project_and/)
- url: /r/ruby/comments/fhzkng/build_a_linux_cli_tool_like_glances_in_ruby/
---

## [9][Using UUIDs to your new Rails 6 application](https://www.reddit.com/r/ruby/comments/fnmqj5/using_uuids_to_your_new_rails_6_application/)
- url: https://www.reddit.com/r/ruby/comments/fnmqj5/using_uuids_to_your_new_rails_6_application/
---

[The first part](https://itnext.io/why-working-with-uuids-instead-of-ids-is-better-b60d22caf601?source=friends_link&amp;sk=ed026c73e0cb06d153c63814c95746c3) of this series describes why working with UUIDs instead of IDs is better
[The second part](https://medium.com/@guillaumeocculy/using-uuids-to-your-rails-6-application-6438f4eeafdf?source=friends_link&amp;sk=cf6c4b885e89ce382f6adea703adea77) explains the power of UUIDs on a brand new Rails application.
## [10][How to do multi-step forms in Rails](https://www.reddit.com/r/ruby/comments/fnpbru/how_to_do_multistep_forms_in_rails/)
- url: https://www.codewithjason.com/rails-multi-step-forms/
---

