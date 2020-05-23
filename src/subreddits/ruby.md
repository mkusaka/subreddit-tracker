# ruby
## [1][How big is Ruby in Japan? What is the japanese community around Ruby up to these days?](https://www.reddit.com/r/ruby/comments/goqwyt/how_big_is_ruby_in_japan_what_is_the_japanese/)
- url: https://www.reddit.com/r/ruby/comments/goqwyt/how_big_is_ruby_in_japan_what_is_the_japanese/
---

## [2][RSA Blind signature scheme in Ruby](https://www.reddit.com/r/ruby/comments/gp16ma/rsa_blind_signature_scheme_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/gp16ma/rsa_blind_signature_scheme_in_ruby/
---
Hey, I'm a bit of lost here

Is there a ruby library that implements blind signatures scheme (preferable with RSA)?

I've obviously looked at openssl but the RSA implementation doesn't have (as far as I can tell) blind signing and unblinding methods implemented.

If there is no ruby implementations, what would you suggest to use instead? C library and import that somehow to Ruby?
## [3][Demo: Ruby 2.7 Pattern Matching on YAML](https://www.reddit.com/r/ruby/comments/goei88/demo_ruby_27_pattern_matching_on_yaml/)
- url: https://www.youtube.com/watch?v=vtJyl2DIZcA&amp;feature=share
---

## [4][5 Ways to Splat in Ruby](https://www.reddit.com/r/ruby/comments/gnve20/5_ways_to_splat_in_ruby/)
- url: https://hint.io/blog/5-ways-to-splat-in-ruby
---

## [5][New to ruby need recommend a tutorial](https://www.reddit.com/r/ruby/comments/go5zx3/new_to_ruby_need_recommend_a_tutorial/)
- url: https://www.reddit.com/r/ruby/comments/go5zx3/new_to_ruby_need_recommend_a_tutorial/
---
hello am a beginner in programming (altought i already used other programming languages but meh), so i never used ruby before, just wondering if you guys can recommend me a good tutorial to follow, thanks.
## [6][Error Handling in GraphQL-Ruby](https://www.reddit.com/r/ruby/comments/go2uoq/error_handling_in_graphqlruby/)
- url: https://www.abhaynikam.me/posts/error-handling-in-graphql-ruby/
---

## [7][Which is the preferred method here](https://www.reddit.com/r/ruby/comments/go38t3/which_is_the_preferred_method_here/)
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
## [8][Having a hard time using ruby bundler on manjaro](https://www.reddit.com/r/ruby/comments/go359m/having_a_hard_time_using_ruby_bundler_on_manjaro/)
- url: https://www.reddit.com/r/ruby/comments/go359m/having_a_hard_time_using_ruby_bundler_on_manjaro/
---
I've been trying to get this [program](https://github.com/tmking/reddit_archiver) to run properly but I've got an error on the last step. I don't know ruby well enough to figure it out but it get the error : 

    bash: bundle: command not found

when I run the command

    $ bundle exec reddit_archiver

It seems to be a problem with bundler which I have uninstalled an reinstalled to no avail. Any help would be greatly appreciated
## [9][A prototype of an interactive tutorial environment for building games with Ruby (link in the comments).](https://www.reddit.com/r/ruby/comments/gnpvp8/a_prototype_of_an_interactive_tutorial/)
- url: https://v.redd.it/w3w1mpldl1051
---

## [10][Live Talks: Going live when this post is 3 hours old](https://www.reddit.com/r/ruby/comments/gnxdvy/live_talks_going_live_when_this_post_is_3_hours/)
- url: https://www.reddit.com/r/ruby/comments/gnxdvy/live_talks_going_live_when_this_post_is_3_hours/
---
Hey everyone!

I hope this doesn't break any community rules. I wanted to share a project me and my colleagues have been doing during this time of pandemic which makes us super proud. We call it Subvisual Live Talks and it's essentially one of us having a chat with an interesting guest, usually well-known within their communities.

A fair number of us have been frequently speaking at and attending multiple conferences for the past few years. We also organise Ruby Conf Portugal, Mirror Conf and local programming and design meetups. As a result, we made a lot of friends in tech communities. Now that we're all at home, we had this weird idea "what if we livestream the conversations we've been having with some of them?"

It's all very laid back and informal. We've been doing it on [Twitch](https://twitch.tv/wearesubvisual) every Thursday at 6 PM GMT+1. If you want to ask questions, we have - erm - _someone_ whose sole job is to go through both the Twitch chat and the #subvisualtalks hashtag on Twitter and write down any questions you want to ask the guest. We'll do our best to have them answered.

The VODs stay available for 7 days on Twitch and end up going to YouTube eventually (depending on how available that - erm - *someone* is).

So far we've had conversations with [Josh Clayton](https://www.youtube.com/watch?v=fponvB8i9us) (managing director @ thoughtbot), [Saša Jurić](https://www.youtube.com/watch?v=4P9WmnUJ2vo) (author of Elixir in Action), [Chris Keathley](https://www.youtube.com/watch?v=wIjtQvj2C9g) (renowed alchemist), [Fred Oliveira](https://www.youtube.com/watch?v=eQo4_AnfAV0) (CTO @ Union) and [Aaron Patterson](https://www.youtube.com/watch?v=ZZ-8O6_nUjE) (Ruby core team, Rails core team, who I'm sure you all know). All of them available on [Youtube](https://www.youtube.com/channel/UCs2aM7E1bqul_KZqK0Sq2vg).

This has been a blast and the conversations tend to have amazingly insightful content. Josh speaking about his challenges in the Marines, Aaron with his adventures through the Ruby garbage collector. And today we have Miles from Cookpad, who was at one point sound engineer for Bob Dylan! It starts at 6 PM GMT+1 on our Twitch channel.

We really hope you enjoy it! We're just doing it for the fun, as more of a hobbie so if you have any feedback or suggestions for the guest or the project, please do let us know.
