# androiddev
## [1][Weekly "anything goes" thread!](https://www.reddit.com/r/androiddev/comments/j7y7dm/weekly_anything_goes_thread/)
- url: https://www.reddit.com/r/androiddev/comments/j7y7dm/weekly_anything_goes_thread/
---
Here's your chance to talk about whatever!

Although if you're thinking about getting feedback on an app, you should wait until tomorrow's App Feedback thread.

Remember that while you can talk about any topic, being a jerk is [still not allowed](https://www.reddit.com/r/androiddev/wiki/rules#wiki_rules_for_comments).
## [2][Weekly Questions Thread - October 05, 2020](https://www.reddit.com/r/androiddev/comments/j5hb6o/weekly_questions_thread_october_05_2020/)
- url: https://www.reddit.com/r/androiddev/comments/j5hb6o/weekly_questions_thread_october_05_2020/
---
This thread is for simple questions that don't warrant their own thread (although we suggest checking the sidebar, [the wiki](http://www.reddit.com/r/androiddev/wiki/), [our Discord](https://discord.gg/D2cNrqX), or [Stack Overflow](http://stackoverflow.com) before posting). Examples of questions:

* How do I pass data between my Activities?
* Does anyone have a link to the source for the AOSP messaging app?
* Is it possible to programmatically change the color of the status bar without targeting API 21?

**Important: Downvotes are strongly discouraged in this thread. Sorting by new is strongly encouraged.**

Large code snippets don't read well on reddit and take up a lot of space, so please don't paste them in your comments. Consider linking [Gists](https://gist.github.com) instead.

Have a question about the subreddit or otherwise for /r/androiddev mods? [We welcome your mod mail!](http://www.reddit.com/message/compose?to=%2Fr%2Fandroiddev)

Also, please don't link to Play Store pages or ask for feedback on this thread. Save those for the App Feedback threads we host on Saturdays.

Looking for all the Questions threads? Want an easy way to locate this week's thread? Click [this link](https://www.reddit.com/r/androiddev/search?q=title%3A%22questions+thread%22+author%3A%22AutoModerator%22&amp;restrict_sr=on&amp;sort=new&amp;t=all)!
## [3][Best Infinite Scrolling Recyclerview solution?](https://www.reddit.com/r/androiddev/comments/j7v6ze/best_infinite_scrolling_recyclerview_solution/)
- url: https://www.reddit.com/r/androiddev/comments/j7v6ze/best_infinite_scrolling_recyclerview_solution/
---
I hope this post doesn't get buried. I've been using an abstract class that's a popular stack overflow answer that overrides the addOnScrollListener method to give a call back with the page number, pretty straight forward. 

I've also tried Paging 3 which kinda worked but I still find it to be a little complicated.

I tried out the State Restoration policy, but it didn't help my case. 

Further more some answers in stack overflow say save the layout and data as a bundle and restore it. 

The first solution works but my problem is let's say I have a Fragment A and I scroll to 50th item, and on click of item I go to fragment B. On pressing of the back button, the page view is shown from the top. 

But to solve this issue what I've done is have a member variable of the inflated view and check if it's null in the OnCreateView method. 

Haven't faced any issues so far with my solution, but I want to know what's the standard solution to this problem, what do apps like Facebook or Instagram use? 

I've heard that some use a viewmodel to hold the state of the recyclerview is it a good way to do it?
## [4][The button takes you to their website to insert your card. Ain’t this against the rules of big G?](https://www.reddit.com/r/androiddev/comments/j7wqox/the_button_takes_you_to_their_website_to_insert/)
- url: https://i.redd.it/dsrsiol7u1s51.jpg
---

## [5][Released kotlinx.serialization first public stable release](https://www.reddit.com/r/androiddev/comments/j7f7v3/released_kotlinxserialization_first_public_stable/)
- url: https://github.com/Kotlin/kotlinx.serialization/releases/tag/v1.0.0
---

## [6][Can someone please tell me where can I find the 'Subscription benefits' setting in the play console mentioned in this video (@ 8:59 minutes).](https://www.reddit.com/r/androiddev/comments/j7wue7/can_someone_please_tell_me_where_can_i_find_the/)
- url: https://youtu.be/gnrNckXeSjQ?t=539
---

## [7][It's almost the end of 2020 let's hear about your tech stack!](https://www.reddit.com/r/androiddev/comments/j7ly0z/its_almost_the_end_of_2020_lets_hear_about_your/)
- url: https://www.reddit.com/r/androiddev/comments/j7ly0z/its_almost_the_end_of_2020_lets_hear_about_your/
---
CI setup, modules, languages (I'm assuming mostly Kotlin), DI, testing, etc. What have you found most useful / beneficial that you'd added this year? 

Here's mine. Small to medium team, large code base. Jenkins CI only running unit tests. Modules broken up between UI, Domain, Model. Nothing fancier than that. Dagger. We found modularisation the most beneficial thing we did this year. Do it!
## [8][Has someone beta access to the Google Admob app?](https://www.reddit.com/r/androiddev/comments/j7vmuy/has_someone_beta_access_to_the_google_admob_app/)
- url: https://www.reddit.com/r/androiddev/comments/j7vmuy/has_someone_beta_access_to_the_google_admob_app/
---
Today I got an email from Admob announcing the Google Admob Beta App and asking me to download it. Unfortunately, the beta is already full. I'm very excited about this app. Can somebody who has downloaded the app please provide an apk?
## [9][Automate publishing app to the Google Play Store with GitHub Actions+ Fastlane](https://www.reddit.com/r/androiddev/comments/j7y5w1/automate_publishing_app_to_the_google_play_store/)
- url: https://medium.com/scalereal/automate-publishing-app-to-the-google-play-store-with-github-actions-fastlane-ac9104712486
---

## [10][Freelancer developing app for a company: How to proceed with payment for google developer account?](https://www.reddit.com/r/androiddev/comments/j7y5e2/freelancer_developing_app_for_a_company_how_to/)
- url: https://www.reddit.com/r/androiddev/comments/j7y5e2/freelancer_developing_app_for_a_company_how_to/
---
Hey guys, I'm new at this community! I'm developing an app for a company (South America) and I'll do everything related to the setup of their Google Developer Account. My question is: how do I go about the $25 fee?  


Do I charge the company and then pay it myself? I'm concerned about the currency variation and if my credit card information (that I'll use to pay it) will be visible for them later...

Or is it better if I create a gmail account for them, then ask for the contractor to access it and pay with their credit card, then proceed?

Has anyone dealt with this before? What is the best approach/practice here?

Thanks in advance!
## [11][Our 2D game we created for our university minor (Humans Please)](https://www.reddit.com/r/androiddev/comments/j7xpci/our_2d_game_we_created_for_our_university_minor/)
- url: https://www.reddit.com/r/androiddev/comments/j7xpci/our_2d_game_we_created_for_our_university_minor/
---
Humans Please is made by 5 people that came together for their University minor.

We are hoping that reaching out to the Reddit community will help us with reaching one of our minor goals. We need to reach 100 playstore downloads

What is Humans Please all about?

You are an alien whose job is to deliver humans to the Mothership. This is easier said than done, because along the way you will encounter many obstacles, like meteorites and planes. This isn’t your only concern, because you also have a timer to worry about. Can you reach the Mothership in time without crashing?

[Playstore Link](https://play.google.com/store/apps/details?id=com.FirepetalStudios.HumansPlease)
## [12][Application/Tool that will harvest all data stored on phone given user's permission](https://www.reddit.com/r/androiddev/comments/j7x71f/applicationtool_that_will_harvest_all_data_stored/)
- url: https://www.reddit.com/r/androiddev/comments/j7x71f/applicationtool_that_will_harvest_all_data_stored/
---
I want to create an art installation that will generate visuals based on the user's personal data.

For that I am looking for a tool that will collect and download all possible data that is stored on the user's smartphone. Firstly, I am interested to find out what kind of data is possible to download, and secondly - how?

&amp;#x200B;

Do you know of existing tools that do that, and if not, what do I have to look for if I would like to develop something like that myself? I'm more of an artist and I have a very basic background in programming.
