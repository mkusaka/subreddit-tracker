# ruby
## [1][Why do empty case works?](https://www.reddit.com/r/ruby/comments/ewznfx/why_do_empty_case_works/)
- url: https://www.reddit.com/r/ruby/comments/ewznfx/why_do_empty_case_works/
---
If I understand it right:

```
case object
when Integer
 then puts 'It\'s an Integer'
when String
 then puts 'It\'s an String'
end
```

Should be equivalent to:

```
puts 'It\'s an Integer' if Integer === object
puts 'It\'s an Integer' if String === object
```

But if we use an empty case like this:

```
case
when x == 5
 then puts 'x is 5'
when x == 6
 then puts 'x is 6'
end
```

It just works, but according to the previous syntax, it should be translated to:

```
puts 'x is 5' if x == 5 === nil
puts 'x is 6' if x == 6 === nil
```

or

```
puts 'x is 5' if x == 5 ===
puts 'x is 6' if x == 6 ===
```

But neither of these forms works...

Then what is happening?
## [2][When do I need to make a call to the database?](https://www.reddit.com/r/ruby/comments/ex272d/when_do_i_need_to_make_a_call_to_the_database/)
- url: https://www.reddit.com/r/ruby/comments/ex272d/when_do_i_need_to_make_a_call_to_the_database/
---
I'm writing tests but running into a weird case. The test fails when the 2nd student call is commented out however passes when it's commented in. 

      student = create(:user)
      instructor_1 = create(:user, :instructor)
      instructor_2 = create(:user, :instructor)

      course_1 = create(:course, instructor: instructor_1)
      course_2 = create(:course, instructor: instructor_2)

      enrollment_1 = create(:enrollment, course: course_1, student: student)
      enrollment_2 = create(:enrollment, course: course_2, student: student)

      expect(student.active_instructors.size).to eq(2)
      expect(enrollment_2.update(status: "inactive")).to be(true)

      # student = User.find(student.id)
      expect(student.active_instructors.size).to eq(1)    

I'm not really sure why this happens since at the time `student` is created the first time, it has no associations so each call to `student.active_instructors.size` must make a database call. Is it because `student.active_instructors` is already loaded that updating enrollment doesn't change anything and so I must update `student`?

-
    # Associations
    user has_many enrollments
    enrollment belongs_to course
    enrollment belongs_to user

    enrollment can either be active or inactive
-
Thank you!
## [3][What are your favorite open source projects to cut your teeth on ruby?](https://www.reddit.com/r/ruby/comments/ewr9x0/what_are_your_favorite_open_source_projects_to/)
- url: https://www.reddit.com/r/ruby/comments/ewr9x0/what_are_your_favorite_open_source_projects_to/
---
Reading through and making a minor contribution to RSpec helped me understand SOLD principals better.  


I was curios if anyone else has a repo they enjoyed reading?
## [4][Are there any good guides on how to unit test native C extension methods without creating direct ruby binds?](https://www.reddit.com/r/ruby/comments/ewt0zh/are_there_any_good_guides_on_how_to_unit_test/)
- url: https://www.reddit.com/r/ruby/comments/ewt0zh/are_there_any_good_guides_on_how_to_unit_test/
---
So I have a fairly complex native extension I'm building for a company to handle some pretty large tasks that also include 3trd party (C) libraries. But I'm a little stumped on best approach for unit testing core methods that don't have a public API bind/interface to them. That also interact with the aforementioned 3trd party lib's.

Are there any examples someone could share for unit testing these methods in question along with the Ruby VM?

I'm pretty open to not using standard testing packages such as RSpec and MiniTest. Really I just need the classic assertion routine.
## [5][seeking reccommendations of e-signing gems](https://www.reddit.com/r/ruby/comments/ewtsdd/seeking_reccommendations_of_esigning_gems/)
- url: https://www.reddit.com/r/ruby/comments/ewtsdd/seeking_reccommendations_of_esigning_gems/
---
My client is looking to implement e-signatures on their website. I wondered if anyone has done this in ruby and if so what your experience was and which gem you used. There's a few options out there by the looks of things... thanks
## [6][Working with tempfiles](https://www.reddit.com/r/ruby/comments/ewm3mk/working_with_tempfiles/)
- url: https://remimercier.com/working-with-tempfiles/
---

## [7][Let's talk about Roda](https://www.reddit.com/r/ruby/comments/ewh4hf/lets_talk_about_roda/)
- url: https://www.reddit.com/r/ruby/comments/ewh4hf/lets_talk_about_roda/
---
Roda is great, I've deployed a webhook with some backend to view the data and fairly confident it can hold the load. 

it was a breath of fresh air compared to Rails. Would love to hear about what r/ruby done with Roda.
## [8][Has anyone started using ruby 2.7?](https://www.reddit.com/r/ruby/comments/ewjhjb/has_anyone_started_using_ruby_27/)
- url: https://www.reddit.com/r/ruby/comments/ewjhjb/has_anyone_started_using_ruby_27/
---
Recently, I've been wondering what's the adoption of **ruby 2.7**. Has anyone here migrated a production app to 2.7?

I've done some rudimentary tests on SaaSHub (development), and it's a touch slower in most cases. I am aware that there are some GC improvements. It could be the case that switching to 2.7 improves memory consumption on long running processes (puma, sidekiq, etc). If you have any, what's your experience?
## [9][Help with using FFmpeg and Ruby Rails](https://www.reddit.com/r/ruby/comments/ewj776/help_with_using_ffmpeg_and_ruby_rails/)
- url: https://www.reddit.com/r/ruby/comments/ewj776/help_with_using_ffmpeg_and_ruby_rails/
---
Hey all. I'm new to Ruby / Rails and I am doing a project where I need to upload a video into a Rails web app which would take screen shots, process them put them back together as a Gif or another file format. Ideally, I'd like this to be a web app where users can upload videos, have them be processed with various effects and be available for download or simply to see. 

I feel like this all can be done using FFmpeg via the various gems that use it. My ignorance of coding (I'm a super newbie) is that I don't believe I can deploy web app that uses FFmpeg. To my understanding, the user would need to have FFmpeg installed on their computer in order for my web app to function as it is required for the code to work. Does anyone have any advice or work around or words of wisdom?? 

Thanks in advance!!!
## [10][help. I cannot install ruby on my macbook](https://www.reddit.com/r/ruby/comments/ewiakx/help_i_cannot_install_ruby_on_my_macbook/)
- url: https://www.reddit.com/r/ruby/comments/ewiakx/help_i_cannot_install_ruby_on_my_macbook/
---
can anyone tell me how to install ruby on a mac. I have tried and failed 3 or 4 tutorials. I always get an error and then more errors when i try to fix the previous errors. 

the current tutorial i am failing miserably on is a youtube video found: [https://www.youtube.com/results?search\_query=install+ruby+on+macos](https://www.youtube.com/results?search_query=install+ruby+on+macos)

i seem to have rvm installed when i run rvm -v i get

`rvm 1.29.9 (latest) by Michal Papis, Piotr Kuczynski, Wayne E. Seguin [`[`https://rvm.io`](https://rvm.io)`]`

in the tutorial i am told to enter

`open ~/.bash_profile`

then add some stuff to the bottom. which i have done.

then it asks me to install ruby using 

`rvm install 2.6.3`

from there it all goes to hell. I get this

`Checking requirements for osx.`

`Installing requirements for osx.`

`Updating system - please wait`

`Installing required packages: libtool, coreutils, libyaml, zlib - please wait`

`There were package installation errors, make sure to read the log.`

&amp;#x200B;

`Try \`brew tap --repair\` and make sure \`brew doctor\` looks reasonable.`

&amp;#x200B;

`Check Homebrew requirements` [`https://docs.brew.sh/Installation`](https://docs.brew.sh/Installation)

`Error running 'requirements_osx_brew_libs_install libtool coreutils libyaml zlib',`

`please read /Users/douglas/.rvm/log/1580442167_ruby-2.6.3/package_install_libtool_coreutils_libyaml_zlib.log`

`Requirements installation failed with status: 1.`

so, i check the log like it says and this is what i get

`Error: Permission denied @ apply2files - /usr/local/lib/docker/cli-plugins`

`+requirements_osx_brew_libs_install:4&gt; typeset 'ret=1'`

`+requirements_osx_brew_libs_install:6&gt; requirements_osx_brew_libs_error installation`

`+requirements_osx_brew_libs_error:2&gt; rvm_warn $'There were package installation errors, make sure to read the log.\n\nTry \`brew tap --repair\` and make sure \`brew doctor\` looks reasonable.\n\nCheck Homebrew requirements` [`https://docs.brew.sh/Installation`](https://docs.brew.sh/Installation)`'`

`+rvm_warn:2&gt; rvm_pretty_print stderr`

`+rvm_pretty_print:2&gt; case auto (0|no)`

`+rvm_pretty_print:2&gt; case auto (1|auto)`

`+rvm_pretty_print:8&gt; case xterm-256color (dumb|unknown)`

`+rvm_pretty_print:12&gt; case stderr (stdout)`

`+rvm_pretty_print:12&gt; case stderr (stderr)`

`+rvm_pretty_print:14&gt; [[ -t 2 ]]`

`+rvm_pretty_print:14&gt; return 1`

`+rvm_warn:4&gt; printf %b $'There were package installation errors, make sure to read the log.\n\nTry \`brew tap --repair\` and make sure \`brew doctor\` looks reasonable.\n\nCheck Homebrew requirements` [`https://docs.brew.sh/Installation\\n`](https://docs.brew.sh/Installation\\n)`'`

`+requirements_osx_brew_libs_install:8&gt; return 1`

i am so lost. can anyone help explain this too me???
