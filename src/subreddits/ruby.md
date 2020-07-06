# ruby
## [1][rubocop-graphql—a rubocop plugin that enforces graphql-ruby best practices](https://www.reddit.com/r/ruby/comments/hm2mrz/rubocopgraphqla_rubocop_plugin_that_enforces/)
- url: https://www.reddit.com/r/ruby/comments/hm2mrz/rubocopgraphqla_rubocop_plugin_that_enforces/
---
Source code is [here](https://github.com/DmitryTsepelev/rubocop-graphql).

For now, the gem can do following things:

- makes sure fields and arguments have snake–cased names;
- suggests to add descriptions to entities;
- forces grouping field definitions together/with their resolvers;
- suggests using input types in complex mutations;
- proposes replacing long resolver methods with resolver objects.
## [2][Contributing to Ruby MRI](https://www.reddit.com/r/ruby/comments/hm0pql/contributing_to_ruby_mri/)
- url: https://kirshatrov.com/2020/01/11/contributing-to-mri/
---

## [3][Soft Delete with Discard](https://www.reddit.com/r/ruby/comments/hm6wju/soft_delete_with_discard/)
- url: https://www.driftingruby.com/episodes/soft-delete-with-discard?utm_medium=social&amp;utm_campaign=weekly_episode&amp;utm_source=reddit
---

## [4][A 65 LOC Instagram scraper powered by Kimurai](https://www.reddit.com/r/ruby/comments/hm5mh5/a_65_loc_instagram_scraper_powered_by_kimurai/)
- url: https://github.com/glaucocustodio/inspider
---

## [5][Fun Facts about Ruby #9 - Fibonacci using Hash.new and recursion](https://www.reddit.com/r/ruby/comments/hlizb2/fun_facts_about_ruby_9_fibonacci_using_hashnew/)
- url: https://i.redd.it/p1go4tyxwz851.png
---

## [6][Simple it's un-google-able.](https://www.reddit.com/r/ruby/comments/hlskdr/simple_its_ungoogleable/)
- url: https://www.reddit.com/r/ruby/comments/hlskdr/simple_its_ungoogleable/
---
Hi folks! I'm new to Ruby and programming in general and I've run into a problem. Unfortunately due to lack of experience, I don't know if it's a syntax or a do-it-differently type problem.

I've been working through a few Ruby, OOP, and Rails books and putting my own twist on the excersizes where I can. [In the dragon class description here](https://pine.fm/LearnToProgram/chap_09.html) it shows calling pet.toss to activate the toss method for that specific dragon instance. Okay fine, I can do that.

For some reason though, this won't work:

    while true
      puts "What would you like to do with #{name}?"
      action = gets.chomp.downcase
    
      if ["feed", "walk", "toss", "rock", "bedtime"].include?(action) do
        # pet(action)
        # pet.(action)
        # pet.action
        pet.toss
        
      end

I was trying to figure out how to pass action through as a variable to decide the method (which I can't find anything online for), but even commenting that out and just having pet.toss doesn't activate the toss method.

      def toss
        puts 'You toss ' + @name + ' up into the air.'
        puts 'He giggles, which singes your eyebrows.'
        passageOfTime
      end

When I input "toss" it just loops back to the "What would you like to do..." without singing my eyebrows.

I know I'm missing something fundamental here. Any help would be appreciated.
## [7][File upload in GraphQL-ruby using Active Storage Direct Upload](https://www.reddit.com/r/ruby/comments/hllknb/file_upload_in_graphqlruby_using_active_storage/)
- url: https://www.abhaynikam.me/posts/active-storage-direct-file-upload-graphql-ruby/
---

## [8][Published the new chapter of API-only ruby on rails course (instapaper clone): Implementing user creation API endpoint](https://www.reddit.com/r/ruby/comments/hlk625/published_the_new_chapter_of_apionly_ruby_on/)
- url: https://www.reddit.com/r/ruby/comments/hlk625/published_the_new_chapter_of_apionly_ruby_on/
---
Hey ruby lovers,

I published the new chapter of **API-only** ruby on rails course. In this chapter, we're building the user creation endpoint. You can check it out from [duetcode.io](https://duetcode.io/rails-api-only-course/implement-user-creation-api-endpoint).  


Last week, I posted the [first two chapters](https://www.reddit.com/r/ruby/comments/hhb1ib/just_started_to_publish_apionly_ruby_on_rails/) to this subreddit, I got some instructive feedback. Please let me know, what do you think about the new chapter or the website itself.

Cheers!
## [9][Open source platformer game - Super Bombinhas](https://www.reddit.com/r/ruby/comments/hln387/open_source_platformer_game_super_bombinhas/)
- url: https://www.reddit.com/r/ruby/comments/hln387/open_source_platformer_game_super_bombinhas/
---
Hi!

Version 0.5.0 of my open source platformer game Super Bombinhas is available!

[https://victords.itch.io/super-bombinhas](https://victords.itch.io/super-bombinhas)

Hope you enjoy.

Thanks!
## [10][A Beginner's Guide to Abstraction](https://www.reddit.com/r/ruby/comments/hle9hx/a_beginners_guide_to_abstraction/)
- url: https://jesseduffield.com/beginners-guide-to-abstraction/
---

