# ruby
## [1][Convert STDIN to human reable error?](https://www.reddit.com/r/ruby/comments/exj2g4/convert_stdin_to_human_reable_error/)
- url: https://www.reddit.com/r/ruby/comments/exj2g4/convert_stdin_to_human_reable_error/
---
I have a command line program that provides a prompt. I'm adding 'begin' - 'rescue' code as a means of error control. Not sure if that's a good way to do it or not, but I'm learning. When an error occurs, I'd like to print that error in plain english rather than the hexodecimal (?) output that I'm currently getting. Ex. Insted of #&lt;IO:0x0000000002abd540&gt; print "Error here"

I know the below just prints the error twice, but I was experimenting. I also the the recent TTY-exit post, I'm looking at that now also, but want to know how to do it without first. 

What I have (snip):

    def execute(stdin, stdout)
    begin
       [spawn(*@args, 0 =&gt; stdin, 1 =&gt; stdout) ]
          rescue
              spawn = STDIN
              puts ("#{spawn}#{STDIN}").to_s
              main
        end
      end
    end
## [2][tty-exit - Terminal exit codes for humans and machines.](https://www.reddit.com/r/ruby/comments/ex6872/ttyexit_terminal_exit_codes_for_humans_and/)
- url: https://github.com/piotrmurach/tty-exit
---

## [3][When do I need to make a call to the database?](https://www.reddit.com/r/ruby/comments/ex272d/when_do_i_need_to_make_a_call_to_the_database/)
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
## [4][Why do empty case works?](https://www.reddit.com/r/ruby/comments/ewznfx/why_do_empty_case_works/)
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
## [5][What are your favorite open source projects to cut your teeth on ruby?](https://www.reddit.com/r/ruby/comments/ewr9x0/what_are_your_favorite_open_source_projects_to/)
- url: https://www.reddit.com/r/ruby/comments/ewr9x0/what_are_your_favorite_open_source_projects_to/
---
Reading through and making a minor contribution to RSpec helped me understand SOLD principals better.  


I was curios if anyone else has a repo they enjoyed reading?
## [6][Are there any good guides on how to unit test native C extension methods without creating direct ruby binds?](https://www.reddit.com/r/ruby/comments/ewt0zh/are_there_any_good_guides_on_how_to_unit_test/)
- url: https://www.reddit.com/r/ruby/comments/ewt0zh/are_there_any_good_guides_on_how_to_unit_test/
---
So I have a fairly complex native extension I'm building for a company to handle some pretty large tasks that also include 3trd party (C) libraries. But I'm a little stumped on best approach for unit testing core methods that don't have a public API bind/interface to them. That also interact with the aforementioned 3trd party lib's.

Are there any examples someone could share for unit testing these methods in question along with the Ruby VM?

I'm pretty open to not using standard testing packages such as RSpec and MiniTest. Really I just need the classic assertion routine.
## [7][seeking reccommendations of e-signing gems](https://www.reddit.com/r/ruby/comments/ewtsdd/seeking_reccommendations_of_esigning_gems/)
- url: https://www.reddit.com/r/ruby/comments/ewtsdd/seeking_reccommendations_of_esigning_gems/
---
My client is looking to implement e-signatures on their website. I wondered if anyone has done this in ruby and if so what your experience was and which gem you used. There's a few options out there by the looks of things... thanks
## [8][Working with tempfiles](https://www.reddit.com/r/ruby/comments/ewm3mk/working_with_tempfiles/)
- url: https://remimercier.com/working-with-tempfiles/
---

## [9][Let's talk about Roda](https://www.reddit.com/r/ruby/comments/ewh4hf/lets_talk_about_roda/)
- url: https://www.reddit.com/r/ruby/comments/ewh4hf/lets_talk_about_roda/
---
Roda is great, I've deployed a webhook with some backend to view the data and fairly confident it can hold the load. 

it was a breath of fresh air compared to Rails. Would love to hear about what r/ruby done with Roda.
## [10][Has anyone started using ruby 2.7?](https://www.reddit.com/r/ruby/comments/ewjhjb/has_anyone_started_using_ruby_27/)
- url: https://www.reddit.com/r/ruby/comments/ewjhjb/has_anyone_started_using_ruby_27/
---
Recently, I've been wondering what's the adoption of **ruby 2.7**. Has anyone here migrated a production app to 2.7?

I've done some rudimentary tests on SaaSHub (development), and it's a touch slower in most cases. I am aware that there are some GC improvements. It could be the case that switching to 2.7 improves memory consumption on long running processes (puma, sidekiq, etc). If you have any, what's your experience?
