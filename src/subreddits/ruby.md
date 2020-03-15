# ruby
## [1][Generate colorfull Triangle patterns](https://www.reddit.com/r/ruby/comments/fimzrs/generate_colorfull_triangle_patterns/)
- url: https://www.reddit.com/r/ruby/comments/fimzrs/generate_colorfull_triangle_patterns/
---
[TrianglePattern](https://github.com/hendrixfan/triangle)

I made a Gem to generate colorfull triangle patterns. You can use the resulting SVG as background image for your Project.
## [2][Ruby Conferences Updates - Status: CANCELD, NEW DATE, ONSITE CANCELED / ONLINE ONLY](https://www.reddit.com/r/ruby/comments/fi3etc/ruby_conferences_updates_status_canceld_new_date/)
- url: https://www.reddit.com/r/ruby/comments/fi3etc/ruby_conferences_updates_status_canceld_new_date/
---
FYI: I try to keep the [Ruby Conference Calendar](https://planetruby.github.io/calendar/2020) up-to-date with Corona Covid-19 cancellations. So far the list includes:
CANCELED:

- Wrocław &lt;3 Ruby (wroclove.rb) @ Poland
- RailsConf @ United States
- Balkan Ruby @ Bulgaria
- RubyNess @ Scottland
- RubyConf India
- RedDotRubyConf @ Singapore 

ONSITE CANCELED / ONLINE ONLY:

- Ruby Wine 2.0 @ Moldova
- Brighton RubyConf @ England

POSTPONED:

- RubyConf Belarus (July)
- RubyKaigi @ Japan (September)
- Ruby Day Italy (September)

PS: Cancellation or postponement or going all online is the responsible thing to do to try to contain (or mitigate) the covid epidemic. See [Coronavirus: why you must act NOW](https://medium.com/@tomaspueyo/coronavirus-act-today-or-people-will-die-f4d3d9cd99ca) for example.
## [3][Adding documentation to a docker image](https://www.reddit.com/r/ruby/comments/fia0et/adding_documentation_to_a_docker_image/)
- url: https://www.reddit.com/r/ruby/comments/fia0et/adding_documentation_to_a_docker_image/
---
TL; DR -- If ruby is installed with the --disable-install-doc flag and then the source is deleted, is there a simple way to go back and install the `ri` documentation?

The Docker container for ruby is slimmed down and they do the configure with the --disable-install-doc option and then they remove the source directory and the tar ball.  This makes sense for a production image.

But during development, I'd like to have the `ri` documentation "local".  One choice is to put it on my laptop outside of the Docker images which is the path I'm currently taking.  But I'm wondering if there is a simple way to go back and create the documentation.

The third alternative I can think of is to use an entirely different image (e.g. an Ubuntu Linux image with the ruby package installed) and use that during development but that implies that my development and testing images will be different from my production images.  And... for RoR in particular, that is somewhat typical since the development and testing systems tend to have additional gems installed (e.g. rspec) that are not installed in production.
## [4][HAPPY FRIDAY THE 13TH Y'ALL](https://www.reddit.com/r/ruby/comments/fi2dlg/happy_friday_the_13th_yall/)
- url: https://i.redd.it/h2id8j4rvgm41.png
---

## [5][need help...](https://www.reddit.com/r/ruby/comments/fiekc0/need_help/)
- url: https://www.reddit.com/r/ruby/comments/fiekc0/need_help/
---
module Modu

	cons = 10

&amp;#x200B;

	def Modu.add(x,y)

		puts x+y

	end

	def sub(x,y)

		puts x-y

	end

end

&amp;#x200B;

class Cls

	include Modu

	def mul(x,y)

		puts x\*y

	end

end

&amp;#x200B;

obj = [Cls.new](https://Cls.new)

puts "enter two numbers "

x = gets.chomp.to\_i

y = gets.chomp.to\_i

obj.mul(x,y)

obj.sub(x,y)

Modu.add(x,y)

puts Modu::cons

&amp;#x200B;

The code works fine except the fact that it doesn't print the cons value and throws a "no\_method\_error"
## [6][Released a string metrics/distance gem written in Rust!](https://www.reddit.com/r/ruby/comments/fhu07l/released_a_string_metricsdistance_gem_written_in/)
- url: https://github.com/anirbanmu/str_metrics
---

## [7][Build a Linux CLI tool like glances in Ruby](https://www.reddit.com/r/ruby/comments/fhzkng/build_a_linux_cli_tool_like_glances_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fhzkng/build_a_linux_cli_tool_like_glances_in_ruby/
---
I was thinking of using ruby to build a CLI tool like the glances or Top commands in linux . Ideally it's just a fun project for myself to learn about how ruby can interact with the system to get info like memory usage and CPU load. What do you guys think of this venture?
## [8][i need help for a ruby script that i made](https://www.reddit.com/r/ruby/comments/fhy31c/i_need_help_for_a_ruby_script_that_i_made/)
- url: https://www.reddit.com/r/ruby/comments/fhy31c/i_need_help_for_a_ruby_script_that_i_made/
---
 

its for school and it keeps getting this notice

(repl):24: syntax error, unexpected end-of-input, expecting keyword\_end begginings()             \^

&amp;#x200B;

&amp;#x200B;

this is the code

&amp;#x200B;

def greeting  
puts "Hello adventurer!  Please type your name: "  
name = gets  
puts "\\n" + "It is nice to meet you #{name}"

return name  
end  
greeting()  
def begginings  
puts "\\n" "you spawn on a mysteriouse island where will you go? 1 for north 2 for east 3 for south 4 for west"  
direction = gets  
operation\_selection = gets.to\_i  
if operation\_selection == 4  
return "you encounter a pack of wild dogs"  
else  
return "you encounter a pack of wild dogs"  
if operation\_selection == 1  
return "you find a tree"  
end  
end  
begginings()
## [9][Announcing RubyGems.org Stats - RubyGems Blog](https://www.reddit.com/r/ruby/comments/fhr6t1/announcing_rubygemsorg_stats_rubygems_blog/)
- url: https://blog.rubygems.org/2020/03/09/announcing-rubygems-stats.html
---

## [10][Cypress vs. Capybara](https://www.reddit.com/r/ruby/comments/fhhd1t/cypress_vs_capybara/)
- url: https://blog.testdouble.com/talks/2020-03-12-cypress-vs-capybara/
---

