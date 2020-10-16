# ruby
## [1][Guide to Reactive Rails](https://www.reddit.com/r/ruby/comments/jc5k8g/guide_to_reactive_rails/)
- url: https://github.com/obie/guide-to-reactive-rails
---

## [2][Looking for feedback on a gem I have been working on: active_snapshot](https://www.reddit.com/r/ruby/comments/jbzqju/looking_for_feedback_on_a_gem_i_have_been_working/)
- url: https://www.reddit.com/r/ruby/comments/jbzqju/looking_for_feedback_on_a_gem_i_have_been_working/
---
I am looking for feedback on the design of a new "gem" that I have been working on called [`active_snapshot`](https://github.com/westonganger/active_snapshot)

**Gem:** [`active_snapshot`](https://github.com/westonganger/active_snapshot)

**Description:** Simplified snapshots and restoration for ActiveRecord models and associations with a transparent white-box implementation.

I appreciate your feedback and help.
## [3][Job Opportunities for Rubyists](https://www.reddit.com/r/ruby/comments/jbrdu3/job_opportunities_for_rubyists/)
- url: https://www.reddit.com/r/ruby/comments/jbrdu3/job_opportunities_for_rubyists/
---
Hi! I'm Billy from Kin Insurance. We're a Chicago-based Series B start-up that is changing home insurance from what it is to what it should be. We've build our platform from the ground-up using Ruby on Rails, and we're growing. If anyone is interested in full-time engineering roles or knows someone on the market, feel free to visit our careers page for more information: [https://www.kin.com/careers](https://www.kin.com/careers).

Also, feel free to post questions in comments, and I'll do my best to respond promptly.
## [4][github analytics with the hubba-reports gem - What are your most used languages (in char / bytes count)?](https://www.reddit.com/r/ruby/comments/jc7zvy/github_analytics_with_the_hubbareports_gem_what/)
- url: https://www.reddit.com/r/ruby/comments/jc7zvy/github_analytics_with_the_hubbareports_gem_what/
---
Hello,  I have split the hubba github analytics gem into two, that is, [hubba](https://github.com/rubycoco/git/tree/master/hubba) and [hubba-reports](https://github.com/rubycoco/git/tree/master/hubba-reports) for easier (re)use and split the data gathering / collecting via github api calls and the report generation. Anyways, I have added a [new language report](https://github.com/rubycoco/git/blob/master/hubba-reports/lib/hubba/reports/reports/languages.rb) that lists all your languages used by char / bytes count and by number of repos. See [LANGUAGES.md](https://github.com/yorobot/backup/blob/master/LANGUAGES.md) as an example. Happy data crunching with ruby. Cheers. Prost. 

PS: Know any other alternative github scripts / gems, please tell.
## [5][RSpec return contain empty array](https://www.reddit.com/r/ruby/comments/jc2e0z/rspec_return_contain_empty_array/)
- url: https://www.reddit.com/r/ruby/comments/jc2e0z/rspec_return_contain_empty_array/
---
Sorry if this is not suitable in here, please guide me to the correct sub cause I don't know where to ask.

I have a search method in Post model:
   
    def self.search(search, category_id)
      if search.strip.empty?
        []
      elsif category_id.empty?
        Post.approved.where('lower(title) LIKE ?', "%#{search.downcase.strip}%")
      else
        @category = Category.find_by('id = ?', category_id.to_i)
        @category.posts.approved.where('lower(title) LIKE ?', "%#{search.downcase.strip}%")
      end
    end
This is my post_spec.rb file:
    
    require 'rails helper'
    
    RSpec.describe Post, type: :model do
      let(:user) { create(:user) }
      let(:category) { create(:category) }
      let(:comments) { create(:comment) }
      let(:likes) { create(:like) }

      describe 'methods' do
        describe 'search post by keyword or by category' do
          let(:post_1) { create(:post, title: 'some example title') }
          let(:post_2) { create(:post, title: 'another title') }

          it 'return post with title like keyword' do
            expect(described_class.search('example', category.id.to_s)).to contain_exactly(post_1)
            expect(described_class.search('EXAMPLE', category.id.to_s)).to contain_exactly(post_1)
            expect(described_class.search('title', category.id.to_s)).to contain_exactly(post_1, post_2)
            expect(described_class.search('foo', category.id.to_s)).to be_empty
          end
        end
      end
    end

I don't know why my actual collection is an empty array but not the post

    1) Post methods search post by keyword or by category return post with title like keyword
      Failure/Error: expect(described_class.search('example', category.id.to_s)).to contain_exactly(post_1)
      expected collection contained:  [#&lt;Post id: 152, title: "some example title"]
      actual collection contained:    []
      the missing elements were:      [#&lt;Post id: 152, title: "some example title"]    

Any help is appreciated cause this is really give me a headache.
## [6][Unexpected performance characteristics when exploring migrating a Rails app to Heroku](https://www.reddit.com/r/ruby/comments/jbvdjs/unexpected_performance_characteristics_when/)
- url: https://bibwild.wordpress.com/2020/10/15/unexpected-performance-characteristics-when-exploring-migrating-a-rails-app-to-heroku/
---

## [7][My personal tests with ractors](https://www.reddit.com/r/ruby/comments/jc1loz/my_personal_tests_with_ractors/)
- url: https://www.reddit.com/r/ruby/comments/jc1loz/my_personal_tests_with_ractors/
---
I was playing with ractors the other day, and I thought my experiments were cool enough to warrent a post...  


```
pi@raspberrypi:~ $ /usr/local/bin/ruby yo.rb 
reminder, benchmark results are "user, system, total, (real)"
without ractors
 37.712969   0.016303  37.729272 ( 38.104503)
with ractors
&lt;internal:ractor&gt;:38: warning: Ractor is experimental, and the behavior may change in future versions of Ruby! Also there are many implementation issues.
 43.812031   4.104823  47.916854 ( 16.301058)
```
```
require 'benchmark'
def do_the_work(x)
  ### Perform meaningless work
  bob = 0 
  50_000.times do
    bob += x
  end
  (1..x).map {"bunny rabbits"}.uniq 
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
puts "with ractors"
puts(Benchmark.measure do
  do_it_with_ractors
end)
```
## [8][How to create a cross-plathorm SDK library?](https://www.reddit.com/r/ruby/comments/jbl93u/how_to_create_a_crossplathorm_sdk_library/)
- url: https://www.reddit.com/r/ruby/comments/jbl93u/how_to_create_a_crossplathorm_sdk_library/
---
Let's say there's a library written in C++ for which I want to create an SDK or gem in Ruby.

&amp;#x200B;

A C++ library provides 3 plathorm-specific pre-compiled libraries:

\* windows, \*.dll

\* linux, \*.so

\* macOS, \*.dylib

&amp;#x200B;

I'll use FFI gem and refer to a particular type of a precompiled library, from my Ruby SDK/gem, depending on a plathorm of the end user.

&amp;#x200B;

(1) What's a conventional way to do it? Where, how is it done?

Namely:

     main_precompiled_lib = 
      if plathorm == Windows
        "./win-lib.dll"
      elseif plathorm == Linux
        "./linux-lib.so"
      elseif ....

(2) How can I provide a way to the user to choose his plathorm?
## [9][Best intermediate - expert training tools and resources?](https://www.reddit.com/r/ruby/comments/jbapcq/best_intermediate_expert_training_tools_and/)
- url: https://www.reddit.com/r/ruby/comments/jbapcq/best_intermediate_expert_training_tools_and/
---
Been using rails and ruby for a while now I'm fairly competent but like to step my game up to the next level with the quality of apps I create. Has anyone found any or have any favourite resources for doing so?
## [10][OpenStruct in Ruby](https://www.reddit.com/r/ruby/comments/jbig5f/openstruct_in_ruby/)
- url: https://medium.com/rubycademy/openstruct-in-ruby-ab6ba3aff9a4
---

