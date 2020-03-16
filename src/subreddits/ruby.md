# ruby
## [1][Ruby among the most active programming languages in GitHub](https://www.reddit.com/r/ruby/comments/fj6rl4/ruby_among_the_most_active_programming_languages/)
- url: https://learnworthy.net/10-most-active-programming-languages-in-github/
---

## [2][Can please someone explain me the code](https://www.reddit.com/r/ruby/comments/fjgsla/can_please_someone_explain_me_the_code/)
- url: https://www.reddit.com/r/ruby/comments/fjgsla/can_please_someone_explain_me_the_code/
---
	f = open("test.txt" , 'a+')

&amp;#x200B;

	f.write("hello ruby")

	f.close



	f =  open("test.txt" , 'a+')



	f.write(". happy bday")

		

	f.close



	f = open("test.txt" , 'a+')



	content = f.sysread(7)

			

	puts content

			

	f.syswrite("fresh")	

	f = open("test.txt" , 'r')		

	print [f.read](https://f.read)

	f.close

&amp;#x200B;

	IN this code , the text is getting printed multiple times . Please help me with this code and the reason for multiple printing of the text.
## [3][Adding typescript to a ember.js + rails 2.3 project - should I still use typescript-rails?](https://www.reddit.com/r/ruby/comments/fjhrwn/adding_typescript_to_a_emberjs_rails_23_project/)
- url: https://www.reddit.com/r/ruby/comments/fjhrwn/adding_typescript_to_a_emberjs_rails_23_project/
---
EDIT: My bad - its a rails 4.2 project, confused it with ruby version.

I've looked around for how to add TS to rails and a lot of newer articles suggest using webpacker which rails 2 doesn't have.

I found [typescript-rails](https://github.com/typescript-ruby/typescript-rails) which can be added to the asset pipeline, but it was last updated on Nov 2016 :/.

So, those on older rails versions, how would you suggest getting TS on rails?
## [4][Is there a Jupyter Notebook equivalent for ruby?](https://www.reddit.com/r/ruby/comments/fj1o5z/is_there_a_jupyter_notebook_equivalent_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fj1o5z/is_there_a_jupyter_notebook_equivalent_for_ruby/
---
I know that in python, Jupyter is used when writing papers or data science, so I was wondering if there is a Ruby equivalent?
## [5][Help extending Math module](https://www.reddit.com/r/ruby/comments/fj7532/help_extending_math_module/)
- url: https://www.reddit.com/r/ruby/comments/fj7532/help_extending_math_module/
---
Hi all,

I'm trying to extend the Math module to include a method, `to_rad` that converts degrees to radians. 

Firstly I'm not entirely sure where to put that file. I've tried in /lib and I can't seem to get the rest of my engine, (models) to recognize that I've adding a method to that module like so[1]. Is there a better place to put it?

I'm sure there's just a misunderstanding of how Ruby namespacing and such works. I've tried googling some stuff and it isn't very helpful in this case for some reason.

[1]

    module Math
      def self.to_rad(angle)
        angle/180 * Math::PI
      end
    end    

In my model (/app/models/model.rb) I'd just like to do `Math.to_rad(x)`

Thanks, and sorry if this isn't the right place for this.
## [6][Github - red-data-tools/GR.rb - GR framework - the graphics library for visualization - for Ruby](https://www.reddit.com/r/ruby/comments/fj1i93/github_reddatatoolsgrrb_gr_framework_the_graphics/)
- url: https://www.reddit.com/r/ruby/comments/fj1i93/github_reddatatoolsgrrb_gr_framework_the_graphics/
---
https://github.com/red-data-tools/GR.rb
## [7][Compact: Automating the contract -collaboration test correspondence](https://www.reddit.com/r/ruby/comments/fja0hl/compact_automating_the_contract_collaboration/)
- url: https://www.reddit.com/r/ruby/comments/fja0hl/compact_automating_the_contract_collaboration/
---
So I wrote my my first Ruby gem and thought it might be nice if someone else knew it existed: [https://github.com/robwold/compact](https://github.com/robwold/compact)

It's very much at the proof-of-concept stage, but I figured this showcased enough of what I had in mind that it could form a starting point for a useful discussion. I'd be really keen to hear from anyone interested in the problem it's trying to solve about how it could be developed into something really useful. If any more experienced OSS maintainers have any feedback on on the code, documentation or whatever, that would be gratefully received too.

For some more background on what problem it's trying to solve:

I watched Joe Rainsberger's talk [Integration tests are a scam](https://www.infoq.com/presentations/integration-tests-scam/) and became interested in the thinking outlined there. (Gary Bernhardt discusses some advantages of this approach in his well-known [Boundaries](https://destroyallsoftware.com/talks/boundaries) talk.) The TL;DR: 

1. If you want to prove the 'basic correctness' of your code with true unit tests, then when you use a test double (mock/stub) in a 'collaboration' test, write a 'contract' test that some class in your code really behaves the way you've said your mock does.
2. This approach is actually more likely to throughly test the basic correctness of your code than the corresponding integration tests, because integration tests are hard to write and slow to run, so nobody writes enough of them to actually exercise all the possible paths of execution in your code.

He said it would be nice to have a tool that helps you automate this correspondence. Sandi Metz mentioned some tools intended to help with this in her talk [The magic tricks of testing](https://www.youtube.com/watch?v=URSWYvyc42M), but they seem to have largely fallen by the wayside. Have a look at the README if you're interested in how I'm trying to approach the problem.

Thanks for reading!
## [8][Gem markdown_helper Reaches 20,000 Downloads](https://www.reddit.com/r/ruby/comments/fj6r07/gem_markdown_helper_reaches_20000_downloads/)
- url: https://www.reddit.com/r/ruby/comments/fj6r07/gem_markdown_helper_reaches_20000_downloads/
---
My gem \[markdown\_helper\]([https://github.com/BurdetteLamar/markdown\_helper#markdown-helper](https://github.com/BurdetteLamar/markdown_helper#markdown-helper)) has reached 20,000 downloads!

Thanks to all users!

Has a \[CLI\]([https://github.com/BurdetteLamar/markdown\_helper#cli](https://github.com/BurdetteLamar/markdown_helper#cli)), btw, so you need not write Ruby code to use it.
## [9][Threads acting strange](https://www.reddit.com/r/ruby/comments/fj2pap/threads_acting_strange/)
- url: https://www.reddit.com/r/ruby/comments/fj2pap/threads_acting_strange/
---
I wanted to add some multitasking to my project and faced a problem : ruby 2.7 executes threads on their declaration, not waiting for .join() command.

that's my test code for threads:

`a =` [`Thread.new`](https://Thread.new)`{sleep 2;puts 'thread'}`  
`sleep 3`  
`puts 'waiting'`  
`sleep 2`  
`puts 'waiting'`  
`a.join`

what this code DESIRED to do:

`waiting`  
`waiting`  
`thread`

what it does for me:

`thread`  
`waiting`  
`waiting`

For some reason threads in new version execute on declaration.  
What am I doing wrong?
## [10][A gem for generating canned data](https://www.reddit.com/r/ruby/comments/fj2ocz/a_gem_for_generating_canned_data/)
- url: https://www.reddit.com/r/ruby/comments/fj2ocz/a_gem_for_generating_canned_data/
---
Currently in quarantine and so have dusted off an old project.

It's a DSL for generating fixture data where the timestamp of the created items is important.  An example of this could be generating test data for a cohort analysis tool.

Anyway, I've been struggling with explaining succinctly how it works and could do with another set of eyes. 

Any feedback on the documentation, on the code or on the concept as a whole will be greatly appreciated!

[https://github.com/tmfnll/bonito](https://github.com/tmfnll/bonito)
