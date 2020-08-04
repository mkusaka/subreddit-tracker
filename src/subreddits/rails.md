# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/
---
In this thread you can showcase your personal pet project to other redditors.

Need help with a specific problem or just wanna have some extra eyeballs on your code? Ask away!

A suggested format to get you started:

1. **Name of your project**
2. **A short description**
3. **Application stack**
4. **Link to Live app**
5. **Link to GitHub**
6. **You experience level**
7. **Other information or areas that you would like advice on**

 

^(Many thanks to Kritnc for getting the ball rolling.)
## [2][Running tests in containers with docker-compose](https://www.reddit.com/r/rails/comments/i3fzbz/running_tests_in_containers_with_dockercompose/)
- url: https://www.reddit.com/r/rails/comments/i3fzbz/running_tests_in_containers_with_dockercompose/
---
The main advantages of this way are to have an independent environment for the tests running and to reduce the complexity of the test environment setup.

What we want to achieve:

* Running the tests should be easy.
* Test runs should be isolated and repeatable.
* Test environment should be as close to the production environment as possible.

How to setup and use docker-compose for Ruby on Rails tests you can find in the article: [https://jtway.co/running-tests-in-containers-with-docker-compose-97480726c1e3](https://jtway.co/running-tests-in-containers-with-docker-compose-97480726c1e3)
## [3][Sending Thousands of Request at the Same Time?](https://www.reddit.com/r/rails/comments/i3iyu5/sending_thousands_of_request_at_the_same_time/)
- url: https://www.reddit.com/r/rails/comments/i3iyu5/sending_thousands_of_request_at_the_same_time/
---
I need to develop an endpoint that will send (possibly) hundreds/thousands of request, process it then return/render the json to user/client.

I've tried using thread pool with the size of 5, but it took forever, but when I tried increasing the size to the number of request, it threw \`ThreadError: can't create Thread: Resource temporarily unavailable\` exception.

I don't think I can use background job/worker for this because I should return the result.

So what should I do?

I was thinking that I should wrap the process in 20sec timeout so it doesn't reach rails 30sec limit, and if it's still not finished in 20sec, it will return the unfinished result. It goes like this [https://i.imgur.com/FolcJV6.png](https://i.imgur.com/FolcJV6.png)

But it's still not working, the process still keep going even after it timeout.
## [4][Run slow RSpec files on Github Actions with parallel jobs using an auto split of the spec file by test examples](https://www.reddit.com/r/rails/comments/i36nrc/run_slow_rspec_files_on_github_actions_with/)
- url: https://www.reddit.com/r/rails/comments/i36nrc/run_slow_rspec_files_on_github_actions_with/
---
The article [shows](https://itnext.io/run-slow-rspec-files-on-github-actions-with-parallel-jobs-using-an-auto-split-of-the-spec-file-by-6bee53e31fb4?source=friends_link&amp;sk=76cf49c59df928f1fa76ded5cc9b1be0) how to address the problem of slow test files negatively impacting the whole build times.
## [5][Heroku high memory tips needed](https://www.reddit.com/r/rails/comments/i39mxf/heroku_high_memory_tips_needed/)
- url: https://www.reddit.com/r/rails/comments/i39mxf/heroku_high_memory_tips_needed/
---
I've recently run into a problem where I'm constantly getting R14 errors and then entering swapped memory. I'm running a single Hobby dyno with Puma at 2 workers,1 thread each (was trying a bunch of configurations).

https://preview.redd.it/wu2puh19vve51.png?width=1209&amp;format=png&amp;auto=webp&amp;s=0dd3db0c3de3f48644c98a980f75ade0ba4f6742

It's hard to tell when this started but I know it's been within the past month and not too much longer before that. As you can see here my number of requests hasn't really gone up to show for an increase in memory usage.  


https://preview.redd.it/yrcz0sw9wve51.png?width=578&amp;format=png&amp;auto=webp&amp;s=7d48a0ac502645f1538c6f1b47117a5129e8b115

I've felt like I've tried a million things by now but cannot find the source of this issue. If anyone has any advice on where I could look it'd be very helpful!
## [6][Tips for creating a multiselect form?](https://www.reddit.com/r/rails/comments/i34z8i/tips_for_creating_a_multiselect_form/)
- url: https://www.reddit.com/r/rails/comments/i34z8i/tips_for_creating_a_multiselect_form/
---
I'm trying to create a form with a similar multiselect as the image below. (Different concept, just using this screenshot to give a visual format of what I'm trying to do.)

For the life of me, I can't seem for figure out the "Rails way" of creating the form using Bootstrap's grid and button groups while preserving the multiselect option.

For context, I've got a model that takes a String attribute that can only be "xxx" or "yyy." The user should be able to select one or both of the options presented to them on that form before hitting the Submit button. The result of that submission should be an Array of the string(s) selected.

Any pointers or examples of how you've tackled this?

&amp;#x200B;

https://preview.redd.it/8up3te2qkue51.png?width=1168&amp;format=png&amp;auto=webp&amp;s=b9a6a1d0dac1ceac995a6494ad798715ee6be036
## [7][Tools for website organization](https://www.reddit.com/r/rails/comments/i3492z/tools_for_website_organization/)
- url: https://www.reddit.com/r/rails/comments/i3492z/tools_for_website_organization/
---
When you're planning a site or are working on one and and find it difficult to hold it all in your head, what is the tool you turn to?

I used draw.io for an entity relationship diagram recently and really liked it.  It feels a little awkward trying to do a site map with it though.
## [8][Moving to Rails in 2020](https://www.reddit.com/r/rails/comments/i2tg0x/moving_to_rails_in_2020/)
- url: https://www.reddit.com/r/rails/comments/i2tg0x/moving_to_rails_in_2020/
---
Hey people. I'm a backend developer working primarily with PHP7.x and TypeScript, with some experience with Go, but in all honesty, I always had a thing for Ruby.   


I've been thinking about learning Ruby and then Rails properly and trying to land a job that would allow me to use them. The only thing that pulls me back is a question of whether or not new stuff is actually created with Ruby/Rails or its mostly legacy code at this point. I'm honestly tired of PHP (mostly community, the language is just fine) and TypeScript doesn't bring to much joy to work with (although it is a great language).   


To sum up my question, are there new Rails apps coming out and, with 5-6 years of experience would I have to look for a junior position, given that I've got experience with other languages and know a fair bit of system design, best practices and design patterns?
## [9][belongs_to associatio name and column name same](https://www.reddit.com/r/rails/comments/i305cy/belongs_to_associatio_name_and_column_name_same/)
- url: https://www.reddit.com/r/rails/comments/i305cy/belongs_to_associatio_name_and_column_name_same/
---
Hi,

&amp;#x200B;

I am upgrading a rails application from [5.1.6.2](https://5.1.6.2) to [5.2.4.3](https://5.2.4.3) and the following that worked previously now throwing errors like 'Segmentation Fault' while running rspec.

belongs\_to :created\_by,  class\_name: 'User', foreign\_key:  :created\_by

If using the column name as association name is discouraged but used to work before, how can I address this by patching or by updating the associations. Any help is appreicated.
## [10][Hooks in ActionView::TestCase](https://www.reddit.com/r/rails/comments/i2wgd8/hooks_in_actionviewtestcase/)
- url: https://www.reddit.com/r/rails/comments/i2wgd8/hooks_in_actionviewtestcase/
---
I am modifying test cases in a test file that looks like this. 

    class TestClass &lt; ActionView
      def setup # called before each test
      end
      def teardown # called after each test
      end
      def test_this
     end
    end

I'm looking for a hook that will run after all the tests complete running. 

I am not sure where to check the documentation for the hooks available in this class. Any help would be great. It's kinda old code and I couldn't find anything on googling. 

Please do not suggest gems to achieve this.
## [11][Credentials failure from NGINX](https://www.reddit.com/r/rails/comments/i2u8bb/credentials_failure_from_nginx/)
- url: https://www.reddit.com/r/rails/comments/i2u8bb/credentials_failure_from_nginx/
---
Hosting my app on DO vps. Have rails native credentials stored. (Rails.application.credentials)

They working fine from `rails c` console but NGINX doesn't see them at all, logging `undefined method [] for NilClass` for them. Can you help me out guys?
