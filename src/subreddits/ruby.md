# ruby
## [1]["relations" using Mongoid.](https://www.reddit.com/r/ruby/comments/ey5pkp/relations_using_mongoid/)
- url: https://www.reddit.com/r/ruby/comments/ey5pkp/relations_using_mongoid/
---
First, I should say I know mongodb is not a relational database. 

Second, I want to make a program which helps me classify my books and my music collection a bit more accessible, so I have a set of authors, and each author has a bunch of books. In classic SQL, there's a definition of Foreign Key, but here ... I'm confused! I only know how I can make a connection to mongodb-server with mongoid and don't really know how can make these relations between tables.
## [2][Emojis from Scratch](https://www.reddit.com/r/ruby/comments/ey6nhh/emojis_from_scratch/)
- url: https://www.driftingruby.com/episodes/emojis-from-scratch
---

## [3][Error handling with Monads in Ruby](https://www.reddit.com/r/ruby/comments/exv4sw/error_handling_with_monads_in_ruby/)
- url: http://nywkap.com/programming/either-monads-ruby.html
---

## [4][How to speed up Ruby and JavaScript tests with static and dynamic split using CI parallelisation](https://www.reddit.com/r/ruby/comments/ey3opt/how_to_speed_up_ruby_and_javascript_tests_with/)
- url: https://docs.knapsackpro.com/2020/how-to-speed-up-ruby-and-javascript-tests-with-ci-parallelisation
---

## [5][Biggest number in Ruby?](https://www.reddit.com/r/ruby/comments/exso3l/biggest_number_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/exso3l/biggest_number_in_ruby/
---
I just fed Ruby a 163,072 digits long number to test the Collatz conjecture, which is insane. Like I'm pretty sure there are less atoms in the universe than a number that big. What's the biggest number it will accept honestly?
## [6][Finally built my first basic ruby project and I need your input](https://www.reddit.com/r/ruby/comments/exp1ng/finally_built_my_first_basic_ruby_project_and_i/)
- url: https://www.reddit.com/r/ruby/comments/exp1ng/finally_built_my_first_basic_ruby_project_and_i/
---
[https://www.reddit.com/r/ruby/comments/en9f8o/planning\_to\_build\_a\_simple\_twitter\_bot/?utm\_source=share&amp;utm\_medium=web2x](https://www.reddit.com/r/ruby/comments/en9f8o/planning_to_build_a_simple_twitter_bot/?utm_source=share&amp;utm_medium=web2x) I mentioned I would be building a bot a while back. So I learned how to use  the twitter gem to build a bot that liked and retweeted stuff about mental health. I called it HudumaBot (Huduma \[hoo-doo-ma\] is swahili for assistance/help)

So I wrote the first script this weekend, and posted it on my github linked below. 

[https://github.com/W3NDO/HudumaBot](https://github.com/W3NDO/HudumaBot)

It's not much, but its honest work. I have some ideas on how to make it better. Planning to add in webscrapping functionality with Nokogiri so it can post threads gathered from mental health websites. 

Lmk what you think and how I can make it better. Cheers
## [7][Need help understanding code](https://www.reddit.com/r/ruby/comments/extoqd/need_help_understanding_code/)
- url: https://www.reddit.com/r/ruby/comments/extoqd/need_help_understanding_code/
---
(Ruby 2.6.5) Code:

\`\`\`f=-&gt;n{n&lt;2?0:1+f\[n\*3/(6-5\*w=n%2)+w\]}\`\`\`

Found it on a code-golf post about coding the collatz conjecture. What I mainly don't understand is why it is typed as f\[...\] and what w is. I am trying to convert it into python but this uses many ruby concepts that I don't understand. What really is going on when this function is called?
## [8][Convert STDIN to human reable error?](https://www.reddit.com/r/ruby/comments/exj2g4/convert_stdin_to_human_reable_error/)
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
## [9][tty-exit - Terminal exit codes for humans and machines.](https://www.reddit.com/r/ruby/comments/ex6872/ttyexit_terminal_exit_codes_for_humans_and/)
- url: https://github.com/piotrmurach/tty-exit
---

## [10][When do I need to make a call to the database?](https://www.reddit.com/r/ruby/comments/ex272d/when_do_i_need_to_make_a_call_to_the_database/)
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
