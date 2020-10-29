# ruby
## [1][Backing Rails enums with Postgres ENUM type](https://www.reddit.com/r/ruby/comments/jk77cb/backing_rails_enums_with_postgres_enum_type/)
- url: https://www.reddit.com/r/ruby/comments/jk77cb/backing_rails_enums_with_postgres_enum_type/
---
[https://diffickens.com/rails/2020/10/02/back-your-enums-with-enums.html](https://diffickens.com/rails/2020/10/02/back-your-enums-with-enums.html)

I haven't blogged since ages, but there was a discussion regarding how to store Rails enums, so I decided to give writing another try. :-)
## [2][Right-hand assignment has been converted into a one-line pattern matching](https://www.reddit.com/r/ruby/comments/jk8icl/righthand_assignment_has_been_converted_into_a/)
- url: https://github.com/ruby/ruby/pull/3703
---

## [3][Ruby's Proposed Software Transactional Memory](https://www.reddit.com/r/ruby/comments/jjsayz/rubys_proposed_software_transactional_memory/)
- url: https://chrisseaton.com/truffleruby/ruby-stm/
---

## [4][More benchmarks, Ractors vs Forks + Redis + Marshal](https://www.reddit.com/r/ruby/comments/jk9zvz/more_benchmarks_ractors_vs_forks_redis_marshal/)
- url: https://www.reddit.com/r/ruby/comments/jk9zvz/more_benchmarks_ractors_vs_forks_redis_marshal/
---
The comments I got on the last post of this sort I made were... underwhelming.  Oh well, it's interesting to me.

This time I rewrote the ractor logic to use forks/redis/marshal instead to see what the difference would be.  Turns out that ractors are substantially slower than using the triumvirate above.  Another interesting thing to note is how similar the fork/redis code is to the ractor code, if you squint at it it's almost identical, and the fork/redis/marshal trick is something we've been able to do as far back as the 1.8 days (and possibly before).

Here are the results...
```
pi@raspberrypi:~ $ /usr/local/bin/ruby hi.rb 
reminder, benchmark results are "user, system, total, (real)"
without ractors
 37.051258   0.001146  37.052404 ( 37.084260)
with forks and redis
  1.091023   0.369672   1.460695 (  9.401046)
with ractors
&lt;internal:ractor&gt;:38: warning: Ractor is experimental, and the behavior may change in future versions of Ruby! Also there are many implementation issues.
 40.912575   2.273917  43.186492 ( 14.284483)
```

and here is the code...

```
require 'redis'
require 'benchmark'

def do_the_work(x)
  ### Perform meaningless work
  bob = 0 
  50_000.times do
    bob += x
  end
  (1..x).map {"bunny rabbits"}.uniq 
end

def do_it_with_forks_and_redis
  pids = (1..8).map do
    fork do
      r = Redis.new
      loop do
        x = Marshal.load(r.blpop("bob")[1])
        r.rpush("answer", Marshal.dump(do_the_work(x)))
      end
    end
  end
  
  pids.each {|pid| Process.detach(pid)}

  r = Redis.new
  a = Array.new
  (1..4_000).each do |x|
    r.rpush("bob", Marshal.dump(x))
  end

  (1..4_000).each do
    a.push(Marshal.load(r.blpop("answer")[1]))
  end
end


def do_it_with_ractors
  ractors = (1..8).map do
    Ractor.new do
      loop do
        x = Ractor.recv
        Ractor.yield(do_the_work(x))
      end
    end
  end
  a = Array.new
  (1..4_000).each do |x|
    ractors.sample.send(x)
  end
  (1..4_000).each do
    a.push(Ractor.select(*ractors)[1])
  end
end

def do_it_without_ractors
  a = Array.new
  (1..4_000).each do |x|
    a.push(do_the_work(x))
  end
end
puts "reminder, benchmark results are \"user, system, total, (real)\""
puts "without ractors"
puts(Benchmark.measure do 
  do_it_without_ractors
end)
puts "with forks and redis"
puts(Benchmark.measure do
  do_it_with_forks_and_redis
end)
puts "with ractors"
puts(Benchmark.measure do
  do_it_with_ractors
end)
```
## [5][What's the Best EC2 Instance Type for Rails Apps?](https://www.reddit.com/r/ruby/comments/jjo4lf/whats_the_best_ec2_instance_type_for_rails_apps/)
- url: https://www.fastruby.io/blog/rails/performance/ruby/best-ec2-instance-for-rails-apps.html
---

## [6][Struct class in Ruby](https://www.reddit.com/r/ruby/comments/jjs5n7/struct_class_in_ruby/)
- url: https://www.sandipmane.dev/struct-class-in-ruby
---

## [7][Database tasks can skip test database with an environment variable | BigBinary Blog](https://www.reddit.com/r/ruby/comments/jji66z/database_tasks_can_skip_test_database_with_an/)
- url: https://blog.bigbinary.com/2020/10/27/database-tasks-can-skip_test_database-with-an-environment-variable.html
---

## [8][How to use ActionText and install Trix Editor on Rails 6 Application](https://www.reddit.com/r/ruby/comments/jjgnqw/how_to_use_actiontext_and_install_trix_editor_on/)
- url: https://www.youtube.com/watch?v=SMsqFRc25gw&amp;t=57s
---

## [9][4 Mistakes to Avoid as a RoR Engineer](https://www.reddit.com/r/ruby/comments/jjbp7c/4_mistakes_to_avoid_as_a_ror_engineer/)
- url: https://thepaulo.medium.com/4-mistakes-to-avoid-as-a-ror-engineer-714374df933f
---

## [10][Squash N+1 queries early with n_plus_one_control test matchers for Ruby and Rails](https://www.reddit.com/r/ruby/comments/jj2ob8/squash_n1_queries_early_with_n_plus_one_control/)
- url: https://evilmartians.com/chronicles/squash-n-plus-one-queries-early-with-n-plus-one-control-test-matchers-for-ruby-and-rails
---

