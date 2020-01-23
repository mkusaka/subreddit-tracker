# ruby
## [1][Sharing an online tool to generate dry-structs from JSON / JSON Schema / GraphQL definitions](https://www.reddit.com/r/ruby/comments/esrvbr/sharing_an_online_tool_to_generate_drystructs/)
- url: https://app.quicktype.io/
---

## [2][Would sorting a sorting an array linearly using multi threading speed up it's sorting time?](https://www.reddit.com/r/ruby/comments/esrrwr/would_sorting_a_sorting_an_array_linearly_using/)
- url: https://www.reddit.com/r/ruby/comments/esrrwr/would_sorting_a_sorting_an_array_linearly_using/
---
Say for instance, we have an array X. If we split X in half and then sort each of those halves concurrently using different threads and finally combine them, would the sorting time drop from O(n) to O(n/2)?

Or is O(n/2) ideally just O(n) with ½ as the constant factor?
## [3][Inside Bitcoin's Proof-of-Work / Waste 10-Minute Mining Lottery - The Ruby Edition by Example](https://www.reddit.com/r/ruby/comments/esgkpj/inside_bitcoins_proofofwork_waste_10minute_mining/)
- url: https://github.com/openblockchains/awesome-blockchains/tree/master/bitcoin_proof_of_work.rb
---

## [4][How would you learn ruby if you had access to all the programs of the startup you're interested in working for?](https://www.reddit.com/r/ruby/comments/eskqgv/how_would_you_learn_ruby_if_you_had_access_to_all/)
- url: https://www.reddit.com/r/ruby/comments/eskqgv/how_would_you_learn_ruby_if_you_had_access_to_all/
---
That might be a dumb question, but I have little experience. I've completed 3 levels of C++ courses in college, I've learned beginner programming from 'How to design programs' (scheme/dr racket), and I've read half a book on ruby.

With that said...my wife is an intern for a start up (ACH transfer platform), in a few months she is going to start learning to code on the job. Since she isn't going to start for another few months, I was hoping to (with permission) borrow their code and learn to program hands on. I know enough that I can recognize what's going on when I read code in an unfamiliar language. I know enough to use stackoverflow as a crutch if I have to. I assume I'm being unrealistic to some extent, but if this is an opportunity to learn Ruby HANDS ON, how would you best learn Ruby in my shoes?
## [5][Rails 6.1 adds query method missing to find orphan records](https://www.reddit.com/r/ruby/comments/esasfy/rails_61_adds_query_method_missing_to_find_orphan/)
- url: https://blog.saeloun.com/2020/01/21/rails-6-1-adds-query-method-missing-to-find-orphan-records
---

## [6][Implementing numbers in "pure" Ruby](https://www.reddit.com/r/ruby/comments/eser00/implementing_numbers_in_pure_ruby/)
- url: https://medium.com/carwow-product-engineering/implementing-numbers-in-pure-ruby-1d35ee53ee70
---

## [7][Rails is Fast: Optimize Your View Performance](https://www.reddit.com/r/ruby/comments/esbysh/rails_is_fast_optimize_your_view_performance/)
- url: https://blog.appsignal.com/2020/01/22/rails-is-fast-optimize-your-view-performance.html
---

## [8][I've made a new search bar for RubyGems with MeiliSearch](https://www.reddit.com/r/ruby/comments/ese6nz/ive_made_a_new_search_bar_for_rubygems_with/)
- url: https://www.reddit.com/r/ruby/comments/ese6nz/ive_made_a_new_search_bar_for_rubygems_with/
---
[https://rubygems.meilisearch.com/](https://rubygems.meilisearch.com/)

As a ruby and rails lover, I often need to find the best gem to solve my issue.

Unfortunately, the search bar in the rubygems web site is not really user-friendly. It is not instant, and despite the advanced search, it's not easy to search in and find the gem you want.

So, as a side project, I wanted to provide a better search bar: you can search by name (as "devise") or by keywords (as "coverage", "cron job"  or "web dev"). It will find the most famous rubygem matching your request.

I used MeiliSearch to implement this search bar. MeiliSearch provides a relevant and search-as-you-type engine. Plus, it's open-source.

So, here is my open-source search bar solution for our favorite open-source language!

Give me your feedback!

&amp;#x200B;

Don't hesitate to give a star to the MeiliSearch repository ⭐️[https://github.com/meilisearch/MeiliSearch](https://github.com/meilisearch/MeiliSearch)

We need your support! All the MeiliSearch team thanks you!
## [9][I made a small Ruby project with bundle gem command and rspec test framework, not sure how to require files inside lib folder in the tests](https://www.reddit.com/r/ruby/comments/esjdy7/i_made_a_small_ruby_project_with_bundle_gem/)
- url: https://www.reddit.com/r/ruby/comments/esjdy7/i_made_a_small_ruby_project_with_bundle_gem/
---
Hello, so inside the `lib` folder I have the following structure:

```
lib/
  foo.rb
  foo/
    A.rb
    B.rb
    version.rb
```

* Inside `foo.rb` I have `require foo/version`, `require foo/A`, `require foo/version`
* At the standard `gemspec` file there is a `spec.require_paths = ['lib']`
* At the standard `spec_helper` file there is a `require 'foo'`

Now I have a `spec/foo/A.rb` file which should test the functionality of `Foo::A`, but everytime I use `Foo::A` it says its an unknown constant...

I'm using `require_relative` to require the specific `foo/A.rb`, but IMO it doesn't seems right to do that... What should I do?

**EDIT**

NVM, renaming `spec/foo/A.rb` to `spec/foo/A_spec.rb` solved the issue, I feel dumb.
## [10][Rost me - I mean, what version of this code is the best for a job position (JR)?](https://www.reddit.com/r/ruby/comments/esf1fo/rost_me_i_mean_what_version_of_this_code_is_the/)
- url: https://www.reddit.com/r/ruby/comments/esf1fo/rost_me_i_mean_what_version_of_this_code_is_the/
---
So I am training for a Junior position as a Ruby Dev.

I am sure there\`s something better then my awnsers. But what I would like to ask is wich of the three is more "professional like". I do preffer the human readbility of the number 3. But it seems people is aways trying to keep stuff in one line.

In this particular example, number 3 is faster (I know I used an extra variable for it, and I hope ruby processor makes it worth :D). 

What do you take in consideration if wanting someone for a position? Wich of the 3 would you like better (or dislike less).

The software is suposed to receive an a-b range and test each number for its value powered to the position (left to right) and sum if it's equal to the original number, returning a list of this numbers. 

Form 1:  838ms

    def sum_dig_pow(a, b)
    
      (a..b).each.map { |x| x if x.to_s.split(//).each_with_index.map { |num, pos| num.to_i ** (pos.to_i + 1)}.sum == x }.compact
      
    end

&amp;#x200B;

Form 2:  863ms

    
    def sum_dig_pow(a, b)
      (a..b).each.map do |x|
        x if x.to_s.split(//).each_with_index.map { |num, pos| num.to_i ** (pos.to_i + 1)}.sum == x
       end.compact
    end

&amp;#x200B;

Form 3:  801ms

    def sum_dig_pow(a, b)
      awnser = []
      (a..b).each do |x|
        deva = x.to_s.split(//).each_with_index.map { |num, pos| num.to_i ** (pos.to_i + 1)}.sum
        if deva == x
          awnser.push(x)
        end
       end
       awnser
    end

Possible awnsers:

1.) Whathever, you\`re Junior, noone is giving you cpu and memory saving jobs

2.) Type less!

3.) Show you know the language using obscure non human readble stuff

4.) It really depends of your boss personal taste
