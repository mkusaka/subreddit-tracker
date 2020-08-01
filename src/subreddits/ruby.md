# ruby
## [1][rodauth-oauth: a rack toolkit for building OAuth/OpenID servers](https://www.reddit.com/r/ruby/comments/i1iyia/rodauthoauth_a_rack_toolkit_for_building/)
- url: https://honeyryderchuck.gitlab.io/rodauth-oauth/
---

## [2][Iterating Over Hashes](https://www.reddit.com/r/ruby/comments/i1bzao/iterating_over_hashes/)
- url: https://www.reddit.com/r/ruby/comments/i1bzao/iterating_over_hashes/
---
Hey y'all

If anyone could help me I'd really appreciate it - I've been trying to learn some Ruby and I've been using Codecademy's 'Learn Ruby' course. It's been going okay I guess, but I'm completely lost on iterating over a hash. I get what a hash is, and it's not making new hashes that fucks me up, it's trying to do anything with the damn thing. 

Here's a thing (I'm a n00b, can you tell?) of code straight from the lesson. I don't think the last line is right, but it looks the same as the hint given on the page??? Everything under the frequencies = [Hash.new](https://Hash.new) line is gibberish and I had to have it solved for me. I've tried googling my question but I'm coming up blank. This code is supposed to take an input from the user and arrange the words by frequency. I really don't understand the | lines or the placeholder thing or well... anything! I'm so lost and I feel like giving up. Halp???

rb:

puts "Enter a phrase you'd like to analyze: "  
text = gets.chomp  
words = text.split  
frequencies = Hash.new(0)  
words.each { |word| frequencies\[word\] += 1 }  
frequencies = frequencies.sort\_by do |word, count|  
 count  
end  
frequencies.reverse!  
frequencies.each do | word, frequency|  
 puts word + " " + frequency.to\_s  
end
## [3][LOWREZ JAM 2020 starts this Saturday and is a comfortable 15 days long (no unhealthy all nighters). Here’s your chance to make a video game with Ruby. I’m giving away DragonRuby Game Toolkit licenses for the event :-)](https://www.reddit.com/r/ruby/comments/i0ym7m/lowrez_jam_2020_starts_this_saturday_and_is_a/)
- url: http://lowrezjam.dragonruby.org
---

## [4][This applies to Ruby as well](https://www.reddit.com/r/ruby/comments/i0v9to/this_applies_to_ruby_as_well/)
- url: /r/rails/comments/i0f8dx/shopify_on_rails_is_doing_well/
---

## [5][Autoprefixer Rails was deprecated](https://www.reddit.com/r/ruby/comments/i0txom/autoprefixer_rails_was_deprecated/)
- url: https://twitter.com/Autoprefixer/status/1288935261842309122
---

## [6][The State of Ruby 3 Typing](https://www.reddit.com/r/ruby/comments/i0lfdw/the_state_of_ruby_3_typing/)
- url: https://developer.squareup.com/blog/the-state-of-ruby-3-typing/
---

## [7][How does this code work and why is it here ?](https://www.reddit.com/r/ruby/comments/i0qlba/how_does_this_code_work_and_why_is_it_here/)
- url: https://www.reddit.com/r/ruby/comments/i0qlba/how_does_this_code_work_and_why_is_it_here/
---
The following lines of code come from the httparty gem. I do not understand the first line. What does it do and what does it resolve to ?


    dir = File.expand_path(File.join(File.dirname(__FILE__), '..', 'lib')) # The problem is here

    require File.join(dir, 'httparty')
    require 'pp'
    config = YAML.load(File.read(File.join(ENV['HOME'], '.twitter')))

 It tried to run it on irb but got the message:   

**NameError (undefined local variable or method `File' for main:Object)**
## [8][Shopify development as a side hustle?](https://www.reddit.com/r/ruby/comments/i0q4p6/shopify_development_as_a_side_hustle/)
- url: /r/cscareerquestions/comments/i0n9sg/shopify_development_as_a_side_hustle/
---

## [9][Ruby for pen-testing](https://www.reddit.com/r/ruby/comments/i0m97d/ruby_for_pentesting/)
- url: https://www.reddit.com/r/ruby/comments/i0m97d/ruby_for_pentesting/
---
Hello guys i have learned ruby and I want to learn some libs for penetration testing,what are the most known or the best libs for hacking ?
## [10][The State of Ruby 3 Typing](https://www.reddit.com/r/ruby/comments/i122g5/the_state_of_ruby_3_typing/)
- url: https://developer.squareup.com/blog/the-state-of-ruby-3-typing/
---

