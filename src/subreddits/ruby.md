# ruby
## [1][My Covid coding project](https://www.reddit.com/r/ruby/comments/gpp37y/my_covid_coding_project/)
- url: https://wishuponaclover.herokuapp.com/
---

## [2][I'm planning to build a rails course, would you like to see it as API-only?](https://www.reddit.com/r/ruby/comments/gpebe0/im_planning_to_build_a_rails_course_would_you/)
- url: https://www.reddit.com/r/ruby/comments/gpebe0/im_planning_to_build_a_rails_course_would_you/
---
Hey ruby lovers,  


I'm creating a ruby on rails course but I'm planning to have an api-only backend application in the course. In my last jobs, I used rails application with modern front-end frameworks like react and vue, so I thought it would have also some potential for rails learners to have api-only course but still wanted to get your feedback about it.  


There was another [post](https://www.reddit.com/r/ruby/comments/gi9wcg/im_creating_a_ruby_on_rails_6_course_what_would/) regarding a similar topic but it's difficult to see the demand of an api-only backend application, that's why I wanted to create this one.  


PS: I have been using reddit without pro-actively interacting with the community, it's nice to have a chance to give something back to both reddit &amp; ruby communities with building a free course.
## [3][Thoughts on Pragmatic Studio for advancing Ruby Skill?](https://www.reddit.com/r/ruby/comments/gpjl92/thoughts_on_pragmatic_studio_for_advancing_ruby/)
- url: https://www.reddit.com/r/ruby/comments/gpjl92/thoughts_on_pragmatic_studio_for_advancing_ruby/
---
So, I know the basics of both Rails and Ruby but I am a bit rusty. I've heard good things about pragmatic studio (both ruby and rails course).

However I finished TheOdinProject awhile back (including the Facebook clone, I never went past the rails course fwiw).

Will these courses be too beginner for me? I need to refresh but they are also quite expensive. I am more concerned with the Ruby one than rails since I feel like my Ruby fundamentals is lacking pretty bad (I still understand the basics though)
## [4][RSA Blind signature scheme in Ruby](https://www.reddit.com/r/ruby/comments/gp16ma/rsa_blind_signature_scheme_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/gp16ma/rsa_blind_signature_scheme_in_ruby/
---
Hey, I'm a bit of lost here

Is there a ruby library that implements blind signatures scheme (preferable with RSA)?

I've obviously looked at openssl but the RSA implementation doesn't have (as far as I can tell) blind signing and unblinding methods implemented.

If there is no ruby implementations, what would you suggest to use instead? C library and import that somehow to Ruby?
## [5][How big is Ruby in Japan? What is the japanese community around Ruby up to these days?](https://www.reddit.com/r/ruby/comments/goqwyt/how_big_is_ruby_in_japan_what_is_the_japanese/)
- url: https://www.reddit.com/r/ruby/comments/goqwyt/how_big_is_ruby_in_japan_what_is_the_japanese/
---

## [6][Demo: Ruby 2.7 Pattern Matching on YAML](https://www.reddit.com/r/ruby/comments/goei88/demo_ruby_27_pattern_matching_on_yaml/)
- url: https://www.youtube.com/watch?v=vtJyl2DIZcA&amp;feature=share
---

## [7][5 Ways to Splat in Ruby](https://www.reddit.com/r/ruby/comments/gnve20/5_ways_to_splat_in_ruby/)
- url: https://hint.io/blog/5-ways-to-splat-in-ruby
---

## [8][New to ruby need recommend a tutorial](https://www.reddit.com/r/ruby/comments/go5zx3/new_to_ruby_need_recommend_a_tutorial/)
- url: https://www.reddit.com/r/ruby/comments/go5zx3/new_to_ruby_need_recommend_a_tutorial/
---
hello am a beginner in programming (altought i already used other programming languages but meh), so i never used ruby before, just wondering if you guys can recommend me a good tutorial to follow, thanks.
## [9][Error Handling in GraphQL-Ruby](https://www.reddit.com/r/ruby/comments/go2uoq/error_handling_in_graphqlruby/)
- url: https://www.abhaynikam.me/posts/error-handling-in-graphql-ruby/
---

## [10][Which is the preferred method here](https://www.reddit.com/r/ruby/comments/go38t3/which_is_the_preferred_method_here/)
- url: https://www.reddit.com/r/ruby/comments/go38t3/which_is_the_preferred_method_here/
---
Hey everyone! So I have a preferred syntax question. Which is the most preferable method here between .count, .length, and .size? I see they all work for this particular code. 

I am new to this sub and so far it has been a wealth of knowledge! 

So thank you all!! 


data = [ [1,2],[5,10] ]

def find_num(array)
    i = 0
    new_array = []
    while i &lt; array.length 
        j = 0
        temp_min = array[i][j]
        
        while j &lt; array[i].length 
            if array[i][j] &lt; temp_min
                temp_min = array[i][j] 
            end 
                j += 1      
        end
          new_array &lt;&lt; temp_min
          i += 1 
    end
      
      return new_array
end

p find_num(data)
