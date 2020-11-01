# ruby
## [1][monofile gem - read in / parse monorepo / mono source tree definitions - a list of git (and github) projects, and more](https://www.reddit.com/r/ruby/comments/jlgkvd/monofile_gem_read_in_parse_monorepo_mono_source/)
- url: https://github.com/rubycoco/monos/tree/master/monofile
---

## [2][Nobody signs Ruby gems: proof in 1-tweet-size code](https://www.reddit.com/r/ruby/comments/jl29xm/nobody_signs_ruby_gems_proof_in_1tweetsize_code/)
- url: https://twitter.com/razum2um/status/1322243197008052225
---

## [3][Supporting RBS type checkinng in httpx](https://www.reddit.com/r/ruby/comments/jl6l8z/supporting_rbs_type_checkinng_in_httpx/)
- url: https://honeyryderchuck.gitlab.io/httpx/2020/10/16/rbs-duck-typing-at-httpx.html
---

## [4][RDoc for Ruby CSV](https://www.reddit.com/r/ruby/comments/jl130l/rdoc_for_ruby_csv/)
- url: https://www.reddit.com/r/ruby/comments/jl130l/rdoc_for_ruby_csv/
---
I've completed a major revision of the RDoc for Ruby class [CSV](https://ruby.github.io/csv/CSV.html), along with its nested classes [CSV::Table](https://ruby.github.io/csv/CSV/Table.html) and [CSV::Row](https://ruby.github.io/csv/CSV/Row.html).  This is not yet merged into the master branch at ruby/ruby, but may be seen on ruby.github.io at the links above.

I've also been experimenting with [CSV recipes](https://ruby.github.io/csv/doc/csv/recipes/recipes_rdoc.html), which focus on specific user tasks rather than on API capabilities.

Oh, and here's the [existing RDoc](https://ruby-doc.org/stdlib-2.7.2/libdoc/csv/rdoc/CSV.html).
## [5][READMEs on Ruby Toolbox project pages](https://www.reddit.com/r/ruby/comments/jkzf7x/readmes_on_ruby_toolbox_project_pages/)
- url: https://www.ruby-toolbox.com/blog/2020-10-30/project-readmes
---

## [6][Packwerk at Shopify](https://www.reddit.com/r/ruby/comments/jkwp16/packwerk_at_shopify/)
- url: https://youtube.com/watch?v=olEA157z7kU
---

## [7][Ruby on Rails in a week](https://www.reddit.com/r/ruby/comments/jkpvwf/ruby_on_rails_in_a_week/)
- url: https://www.simplethread.com/ruby-on-rails-in-a-week/
---

## [8][how is this converted to a hash in a function call?](https://www.reddit.com/r/ruby/comments/jktr6s/how_is_this_converted_to_a_hash_in_a_function_call/)
- url: https://www.reddit.com/r/ruby/comments/jktr6s/how_is_this_converted_to_a_hash_in_a_function_call/
---
Suppose I have a method like this 

    def fun(a, options = {})
      puts [a, options].inspect
    end
    fun(1, a: 3, b: 4) =&gt;  [1, {:a=&gt;3, :b=&gt;4}]

I don't understand how this gets converted to a hash. Also, is it possible to have this in the middle somewhere? like this  


    def fun(a, options = {}, c)
      puts [a, options].inspect
    end
    fun(1, a: 3, b: 4, 10)

without have to wrap \`a: 3, b: 4\` into  {}?
## [9][Future of Ruby - The Journey to One Million Websocket Connections](https://www.reddit.com/r/ruby/comments/jki55j/future_of_ruby_the_journey_to_one_million/)
- url: https://www.youtube.com/watch?v=Dtn9Uudw4Mo
---

## [10][Why I love ruby, a complete RPC system using redis in 57 lines of code :)](https://www.reddit.com/r/ruby/comments/jki6b7/why_i_love_ruby_a_complete_rpc_system_using_redis/)
- url: https://www.reddit.com/r/ruby/comments/jki6b7/why_i_love_ruby_a_complete_rpc_system_using_redis/
---
So this RPC system has some advantages...
1. I programmed it as a toy (which was fun for me)
2. It's simple, if there is a bug in 57 lines of code, it's easy to find and fix
3. It uses redis, which I like, is simple, and fast
4. It handles exceptions which is probably necessary if anyone ever wanted to actually use it
5. there are no timeouts for long running code, cause sometimes timeouts are frustrating
6. there are timeouts for return values because those should be fast anyway
7. usage is simple so you don't have to understand what's going on to use it
8. if it's not already obvious rrecv and rsend invocations can occur on different processes (or with a little more configuration on different machines)  (all classes used as arguments should be available to all processes since we're using Marshaling)

```
rrecv(:name_of_thing_that_can_be_called) do |args|
#do whatever processing you want
#whatever is returned by this block is returned to the caller
end
```
```
result = rsend(:name_of_thing_that_can_be_called, &lt;whatever args you want&gt;, &lt;can be many&gt;)
```
every time you define a new callable thing with rrecv, it will spin up a new thread to respond to those requests

future work, add the ability to control how many workers rrecv creates, and possibly allow other types of workers other than just threads.



```
require 'securerandom'
require 'redis'

class Message
  attr_reader :payload, :guid

  def initialize(*args)
    @guid = SecureRandom.uuid
    @payload = args
  end
end

class RPCExceptionContainer
  attr_reader :saved_exception

  def initialize(exception)
    @saved_exception = exception
  end
end

def rsend(key, *args)
  @r ||= Redis.new
  Message.new(*args).tap do |message|
    @r.rpush(key, Marshal.dump(message))
    Marshal.load(@r.blpop(message.guid)[1]).tap do |return_value|
      if return_value.class == RPCExceptionContainer
        raise return_value.saved_exception
      else
        return return_value
      end
    end
  end
end

def rrecv(key)
  @keys ||= Hash.new
  if(!@keys.has_key?(key))
    r = Redis.new
    Thread.new do
      loop do
        Marshal.load(r.blpop(key)[1]).tap do |message|
          begin
            return_value = yield(message.payload)
            r.rpush(message.guid, Marshal.dump(return_value))
            r.expire(message.guid, 30)
          rescue
            r.rpush(message.guid, Marshal.dump(RPCExceptionContainer.new($!)))
            r.expire(message.guid, 30)
          end
        end
      end
    end
    @keys[key] = true
  else
    raise "You Called rrecv twice on the same key, this is almost certainly not what you intended"
  end
end

rrecv(:hello) do |args|
  puts "Hello #{args.reduce(:+)}"
  raise "what what"
  "success"
end
x = rsend(:hello, "hi there", "nobody") rescue "bobby farquard"
puts x

rrecv(:there) do |args|
  puts "There #{args.reduce(:+)}"
  "success"
end
y = rsend(:there, "something else")
puts y


x = rsend(:hello, "hi there", "nobody") rescue "bobby farquard again"
puts x
```
