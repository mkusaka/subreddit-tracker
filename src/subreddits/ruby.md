# ruby
## [1][hubba gem - little github api helper for (auto-)generating statistics reports (summary, stars, timeline, updates)](https://www.reddit.com/r/ruby/comments/j8fz13/hubba_gem_little_github_api_helper_for/)
- url: https://github.com/rubycoco/git/tree/master/hubba
---

## [2][Remote Ruby: Adam Wathan joins to talk TailwindCSS, Tailwind UI, and View Components](https://www.reddit.com/r/ruby/comments/j7zouw/remote_ruby_adam_wathan_joins_to_talk_tailwindcss/)
- url: https://share.transistor.fm/s/7bd5b843
---

## [3][Can someone help me out with what this does rand(999_999)](https://www.reddit.com/r/ruby/comments/j7vkhg/can_someone_help_me_out_with_what_this_does/)
- url: https://www.reddit.com/r/ruby/comments/j7vkhg/can_someone_help_me_out_with_what_this_does/
---
I'm trying to generate a fixed digit random number, and I came across this code. Could someone help me with what is happening here?  


I'm not sure i get much from the documentation for rand either.
## [4][Ruby 2.7.2 has revised Rdoc for Hash](https://www.reddit.com/r/ruby/comments/j7l51e/ruby_272_has_revised_rdoc_for_hash/)
- url: https://www.reddit.com/r/ruby/comments/j7l51e/ruby_272_has_revised_rdoc_for_hash/
---
I've spent a good part of this year working on revisions to the RDoc for Hash.  The revisions include all methods documentation in addition to much of the introductory text (class documentation).  Lots more example code.

Check it out:

\- 2.7.2: [https://ruby-doc.org/core-2.7.2/Hash.html](https://ruby-doc.org/core-2.7.2/Hash.html)

\- 2.7.1: [https://ruby-doc.org/core-2.7.1/Hash.html](https://ruby-doc.org/core-2.7.1/Hash.html)
## [5][Multithreading -- Parallelizing loops](https://www.reddit.com/r/ruby/comments/j80sz0/multithreading_parallelizing_loops/)
- url: https://www.reddit.com/r/ruby/comments/j80sz0/multithreading_parallelizing_loops/
---
I am looking for something like Julia's `Threads.@threads` macro, which allows to parallelize a for loop within multithreading, now I am looking for something like this for Ruby.
## [6][Fastest way to generating a n digit number](https://www.reddit.com/r/ruby/comments/j7w5xn/fastest_way_to_generating_a_n_digit_number/)
- url: https://www.reddit.com/r/ruby/comments/j7w5xn/fastest_way_to_generating_a_n_digit_number/
---
I found this code to generate 6 digit number 

    rand(100000...999999)

Is this approach fine?   
I'm not sure what the two dots(..) or three dots(...) mean and I hope it doesn't expand into an array and then rand selects a values
## [7][Speeding up Rails with Memoization](https://www.reddit.com/r/ruby/comments/j7v6l2/speeding_up_rails_with_memoization/)
- url: https://www.honeybadger.io/blog/ruby-rails-memoization/?utm_source=rubyweekly&amp;utm_medium=email&amp;utm_campaign=ruby
---

## [8][Beginner need some help understaind data pipeline](https://www.reddit.com/r/ruby/comments/j7vlbx/beginner_need_some_help_understaind_data_pipeline/)
- url: https://www.reddit.com/r/ruby/comments/j7vlbx/beginner_need_some_help_understaind_data_pipeline/
---
Context:  
 I am pulling in data from a social media site and storing it in my database. For eg. Everday I make an api req to fetch all tweets of a user. Then make another api req for each tweet to fetch its public\_metrics like tweets, likes,etc.   


If I am using Sidekiq, how should I build my data pipeline?  


Right now, I have a worker who just makes both the api calls from the same worker.  
Should I separate, "fetching all tweets" into a seprate worker and "fetch tweets metrics" into a separate wrker?
## [9]["Ruby 2.7.2 ... contains intentional incompatibility"](https://www.reddit.com/r/ruby/comments/j7fk83/ruby_272_contains_intentional_incompatibility/)
- url: https://www.reddit.com/r/ruby/comments/j7fk83/ruby_272_contains_intentional_incompatibility/
---
[https://www.ruby-lang.org/en/news/2020/10/02/ruby-2-7-2-released/](https://www.ruby-lang.org/en/news/2020/10/02/ruby-2-7-2-released/)

Anybody know what the incompatibility is? I'm glad they mention it as an issue, but I couldn't find any details on that page.
## [10][Help to create the shortest hexadecimal string from two strings](https://www.reddit.com/r/ruby/comments/j7juoj/help_to_create_the_shortest_hexadecimal_string/)
- url: https://www.reddit.com/r/ruby/comments/j7juoj/help_to_create_the_shortest_hexadecimal_string/
---
I have two strings say first = "say" and second = "henlo". I want to find the shortest possible user-friendly string(a-z, 0-9) that can encode this information. This string isn't any encrypted value I want it to be as human-readable as possible.

This is how I do it currently. I combine the two words with space since both the words can't contain space like "first second" word = "say henlo" and compute the hexdigest using SHA1 and I get a value like so

    require 'faker'
    
    all = {}
    iterations = 5000000
    word_size = 10
    count, duplicates = 0, 0
    
    iterations.times do
        word = Faker::Lorem.sentence(80, true, 15).split(' ').sample(2).join(' ')
        digest = Digest::SHA1.hexdigest(word)[0..word_size-1]
        if all[digest] == word
            duplicates += 1
            next
        end
        if all.include?(digest) 
            puts [digest, all[digest], word].inspect
            count += 1
        end
        all[digest] = word
    end
    puts [iterations, word_size, duplicates, count, (count.to_f / (iterations - duplicates)) * 100].inspect

And this is the result I got

    [1000000, 8, 363802, 64, 0.01]
    [1000000, 10, 363723, 0, 0.0]
    [5000000, 10, 3769185, 0, 0.0]

Now, I have a couple of questions.

1. Is there any other way to shrink the hexdigest likes into something shorter, and user-readable string without increasing the chance of collision?
2. What is the probability that taking the first n characters is unique for a set of words like 1M words? How could I reliability calculate it to compute the trade-offs

EDIT:The purpose of the short string isn't to substitute saving both the string themselves but to generate a short string for the user which they can use in their API calls.  
EDIT 2: The reason I can't use the "human-friend" words as it is, is because it can be languages with accents(French and unicode ones also)  and API key is ASCII only. 
