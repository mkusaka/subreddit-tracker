# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/
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
## [2][Need help: Rails credentials in Github actions pipeline](https://www.reddit.com/r/rails/comments/jaa7vc/need_help_rails_credentials_in_github_actions/)
- url: https://www.reddit.com/r/rails/comments/jaa7vc/need_help_rails_credentials_in_github_actions/
---
Hello dear Rails community, I am relatively new to Rails and to learn I am building a new application which I want to deploy to my Ubuntu v-server via Github actions. I am really struggling with deploying the app since using this particular stack (Rails / Docker / Github actions) is new to me.

I dockerized the app and I want to push it to Dockerhub, so I can retrieve the image from Dockerhub on my server. The image pushing and pulling works fine, but the app wouldn't build on the server with the image because I guess I didn't (and still don't) fully understand the process. This was when I posted this question on Stackoverflow:

[https://stackoverflow.com/questions/64239706/cannot-start-dockerized-rails-application-when-pulling-it-from-docker-hub](https://stackoverflow.com/questions/64239706/cannot-start-dockerized-rails-application-when-pulling-it-from-docker-hub)

After studying on how to set up a Rails app in Github actions I have come really far imho and currently I get an error in the pipeline that googling the error suggests something is wrong with my credentials:

    ActiveSupport::MessageEncryptor::InvalidMessage: Cannot load database configuration:
    ActiveSupport::MessageEncryptor::InvalidMessage
    ... (I removed the rest of the stack to save space)
    Caused by:
    OpenSSL::Cipher::CipherError: 

I hope I provided all the necessary info and I would appreciate any help I can get.
## [3][Way to test your application under big traffic - jMeter](https://www.reddit.com/r/rails/comments/j9nrkh/way_to_test_your_application_under_big_traffic/)
- url: https://www.reddit.com/r/rails/comments/j9nrkh/way_to_test_your_application_under_big_traffic/
---
If you would like to check where are endpoints "to-be-improved" in your application, before users will encounter it during big traffic - there you have quick blogpost about our experience apache-jmeter [https://www.2n.pl/blog/apache-jmeter](https://www.2n.pl/blog/apache-jmeter) . Feedback / other approaches would be appreciated!
## [4][DEPLOYED MY FIRST APPLICATION!! (That actually went right.)](https://www.reddit.com/r/rails/comments/j9mdqg/deployed_my_first_application_that_actually_went/)
- url: https://www.reddit.com/r/rails/comments/j9mdqg/deployed_my_first_application_that_actually_went/
---
Hey guys, just wanted to thank all the people of the Rails community for all the help and guidance y’all gave me over the past months.

Finally deployed a working application and planning on keeping it up for a long time!

Next mini project: Deploying Wordpress to a subdomain.

Literally spent hours looking through documentation and it feels so satisfying. Goodnight! (It’s 3 AM and I have work in 4 hours. 😂)

Thanks again! 🙏
## [5][How to notify IE users that their browser is not supported?](https://www.reddit.com/r/rails/comments/j9qfug/how_to_notify_ie_users_that_their_browser_is_not/)
- url: https://www.reddit.com/r/rails/comments/j9qfug/how_to_notify_ie_users_that_their_browser_is_not/
---
My current project is not supporting IE at all. I've looked at progressive enhancement but IE is just so far behind modern browser that it doesn't even support important features for this project. Instead of having to deal with confused IE users that just see a broken site I want to automatically redirect them to a page that tells them to use a modern browser. Now I'm wondering what would be the best way to achieve this?

I've thought of three ways so far and wanted to get some opinions on them:

1. Look at useragent on rails side and use a redirect. Basically locking IE users out of the app completely because every request would be redirected.

2. Look at useragent on frontend via JS and redirect that way. Same as the first solution but would instead be happening on the users device. This would mean rails wouldn't have to check the useragent on every request.

3. Look at useragent on frontend via JS and show a notification. This would allow the user to remove the notification and "use" the application (even though IE users literally can't use it). Probably the user friendliest.

Would love to get some opinions on those options. Maybe there is also a way that I haven't even thought about?
## [6][Add library (not as gem) in Rails 6](https://www.reddit.com/r/rails/comments/j9qu2m/add_library_not_as_gem_in_rails_6/)
- url: https://www.reddit.com/r/rails/comments/j9qu2m/add_library_not_as_gem_in_rails_6/
---


I need to add this library: https://photon.komoot.de/ in my Ruby on Rails application. I should only use it in one page (view). I wouldn't want to set up complicated stuff like Webpacker. Is there any way to add a js library quickly?
## [7][Is it necessary to install Kibana while using elastic-search](https://www.reddit.com/r/rails/comments/j9nnjk/is_it_necessary_to_install_kibana_while_using/)
- url: https://www.reddit.com/r/rails/comments/j9nnjk/is_it_necessary_to_install_kibana_while_using/
---
I'm using elastic-search by following this tutorial [https://iridakos.com/programming/2017/12/03/elasticsearch-and-rails-tutorial](https://iridakos.com/programming/2017/12/03/elasticsearch-and-rails-tutorial) 

but by running  `User.import(force: true) I` face  an error `uninitialized constant Faraday::Error::ConnectionFailed`  i don't know the reason my be it occurs because I skip the kabana installation but I don't think so.

Need help from seniors.
## [8][Rails + Hacktoberfest meetup tonight! JWTs for Rails!](https://www.reddit.com/r/rails/comments/j9q575/rails_hacktoberfest_meetup_tonight_jwts_for_rails/)
- url: https://www.reddit.com/r/rails/comments/j9q575/rails_hacktoberfest_meetup_tonight_jwts_for_rails/
---
🎟[https://live.remo.co/e/israelrb-hacktoberfest/register](https://live.remo.co/e/israelrb-hacktoberfest/register)  
⏰ 19:00 Israel / 17:00 London / 12:00 NYC / 09:00 SF  


Join us for the 3rd israel.rb virtual meetup!

This time it'll be a Hacktoberfest edition with special guest Dan Moore from FusionAuth discussing JSON Web Tokens (JWTs) and what Rails developers need to know about them.

We'll also have lightning talks from the community! If you're interested in giving a 5-10 minute lightning talk on an open source theme, let us know!

https://preview.redd.it/96znkrxqsns51.jpg?width=820&amp;format=pjpg&amp;auto=webp&amp;s=53294b0be8ba64ced90afeb8676aeffa971b67c8
## [9][Fill PDF form with user data [Prawn]](https://www.reddit.com/r/rails/comments/j984ay/fill_pdf_form_with_user_data_prawn/)
- url: https://www.reddit.com/r/rails/comments/j984ay/fill_pdf_form_with_user_data_prawn/
---
My goal is to fill the existing PDF interactive form with user data.

Requirements for this are:

*   it SHOULD be able to insert data into text fields;
*   it SHOULD be commercial free 
*   it MAY be able to insert an image on the XY position.

Based on research, I chose [Prawn-Rails gem](https://github.com/cortiz/prawn-rails). I'm a little confused with documentation, how can fill form fields on the existing PDF interactive form? 

&amp;#x200B;

The process is like this:

User fills in the form, the service collects data from the form and generates pdf based on that. 

The generated pdf is always the same, only user data is changing.
## [10][Return Forms from backend ?](https://www.reddit.com/r/rails/comments/j957e4/return_forms_from_backend/)
- url: https://www.reddit.com/r/rails/comments/j957e4/return_forms_from_backend/
---
Hello folks, any one tried to return forms as json and in frontend just check the type of field if it is an file input or a text area or a text , number ? any one tried that before ? any guide ?

backend rails 

frontend using react
## [11][Building a Linktree clone with Rails in the live stream](https://www.reddit.com/r/rails/comments/j92bpp/building_a_linktree_clone_with_rails_in_the_live/)
- url: https://youtu.be/A9vTag8-2JY
---

