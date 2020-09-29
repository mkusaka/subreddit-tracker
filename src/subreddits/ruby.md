# ruby
## [1][Ruby one-liners cookbook with hundreds of examples and exercises](https://www.reddit.com/r/ruby/comments/j1ph69/ruby_oneliners_cookbook_with_hundreds_of_examples/)
- url: https://learnbyexample.github.io/learn_ruby_oneliners/one-liner-introduction.html
---

## [2][How to Fix PostgreSQL Performance Issues with PG Extras](https://www.reddit.com/r/ruby/comments/j1vsmk/how_to_fix_postgresql_performance_issues_with_pg/)
- url: https://pawelurbanek.com/postgresql-fix-performance
---

## [3][Building a Ractor based logger that will work with non-Ractor compatible code](https://www.reddit.com/r/ruby/comments/j1z2o2/building_a_ractor_based_logger_that_will_work/)
- url: https://mensfeld.pl/2020/09/building-a-ractor-based-logger-that-will-work-with-non-ractor-compatible-code/
---

## [4][Parsby – Parser combinator library for Ruby, based on Haskell's Parsec](https://www.reddit.com/r/ruby/comments/j1hkhl/parsby_parser_combinator_library_for_ruby_based/)
- url: https://github.com/jolmg/parsby
---

## [5][TIL you can do simple destructuring assignment in Ruby 3 with the `in` keyword](https://www.reddit.com/r/ruby/comments/j1mu31/til_you_can_do_simple_destructuring_assignment_in/)
- url: https://www.reddit.com/r/ruby/comments/j1mu31/til_you_can_do_simple_destructuring_assignment_in/
---
I've been following the ruby-lang atom feed lately and these gems (ha) keep coming up. This time it was a code snippet that showed up in [this discussion](https://bugs.ruby-lang.org/issues/16986#note-45)

    X = Struct.new(:foo, :bar, keyword_init: true)
    obj = X.new(foo: 1, bar: 2)

    obj in {foo:, bar:}
    foo # =&gt; 1
    bar # =&gt; 2

I tested it with array destructuring as well and it worked pretty handily (ruby 3.0.0-preview1)

    a = %i{a b}
    a in [this, that]
    this =&gt; :a
    that =&gt; :b

I thought the `in` keyword was limited to `case/in/end` but apparently not!
## [6][[HIRING] Any senior Ruby + React Full-Stack Engineers around Bangalore / okay to relocate to Bangalore?](https://www.reddit.com/r/ruby/comments/j1y8pp/hiring_any_senior_ruby_react_fullstack_engineers/)
- url: https://www.reddit.com/r/ruby/comments/j1y8pp/hiring_any_senior_ruby_react_fullstack_engineers/
---
Would love to have a chat. DM.  
CTC is good at around 40 LPA.
## [7][The Complexity of Active Record Transactions](https://www.reddit.com/r/ruby/comments/j1a97j/the_complexity_of_active_record_transactions/)
- url: https://janko.io/the-complexity-of-activerecord-transactions/
---

## [8][Rails Monolith towards Engines spike - The experience of a ten-people team](https://www.reddit.com/r/ruby/comments/j1gzd2/rails_monolith_towards_engines_spike_the/)
- url: https://withatwist.dev/our-rails-engines-spike.html
---

## [9][Rails Sidekiq configuration for micro services on reverse proxy.](https://www.reddit.com/r/ruby/comments/j1bsas/rails_sidekiq_configuration_for_micro_services_on/)
- url: https://blog.joshsoftware.com/2019/07/31/rails-sidekiq-configuration-for-micro-services-on-reverse-proxy/
---

## [10][Sandwich Scheduling: a look at Ruby’s Array cycle method](https://www.reddit.com/r/ruby/comments/j195jp/sandwich_scheduling_a_look_at_rubys_array_cycle/)
- url: https://www.mesomorphic.co.uk/blog/2020/09/28/sandwich-scheduling/
---

