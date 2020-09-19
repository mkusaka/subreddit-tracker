# androiddev
## [1][App Feedback Thread - September 19, 2020](https://www.reddit.com/r/androiddev/comments/ivrob6/app_feedback_thread_september_19_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ivrob6/app_feedback_thread_september_19_2020/
---
This thread is for getting feedback on your own apps.

####Developers:

- must **provide feedback** for others
- must include **Play Store**, **GitHub**, or **BitBucket** link
- must make top level comment
- must make effort to respond to questions and feedback from commenters
- may be open or closed source

####Commenters:

- must give **constructive feedback** in replies to top level comments
- must not include links to other apps

To cut down on spam, accounts who are too young or do not have enough karma to post will be removed. Please make an effort to contribute to the community before asking for feedback.

As always, the mod team is only a small group of people, and we rely on the readers to help us maintain this subreddit. Please report any rule breakers. Thank you.

\- Da Mods
## [2][Weekly Questions Thread - September 14, 2020](https://www.reddit.com/r/androiddev/comments/isiooc/weekly_questions_thread_september_14_2020/)
- url: https://www.reddit.com/r/androiddev/comments/isiooc/weekly_questions_thread_september_14_2020/
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
## [3][Google rewrites Google Pay app in Flutter](https://www.reddit.com/r/androiddev/comments/ivlxm1/google_rewrites_google_pay_app_in_flutter/)
- url: https://twitter.com/googledevs/status/1307126032349224960?s=19
---

## [4][What are effective approaches for Android dialogs in 2020?](https://www.reddit.com/r/androiddev/comments/ivf6xm/what_are_effective_approaches_for_android_dialogs/)
- url: https://www.reddit.com/r/androiddev/comments/ivf6xm/what_are_effective_approaches_for_android_dialogs/
---
We have seen many large changes to Android development since 2017 (Kotlin, JetPack, AAC etc), but what about good old dialogs? I remember dialogs being in a state of chaos on Android. Use Dialog class? No, use AlertDialog. Then Google gave us DialogFragment where you can inflate a view like a standard fragment or override onCreateDialog(). Which was for a long time an improvement. Then came Material Design, and it seemed like if we wanted Material dialogs we better implement them ourselves with custom layouts. Maybe that is why a library like afollestad's Material Dialogs has almost 20,000 stars on GitHub -- it provides a simple API for material styled dialogs. This is what everybody wants but of course Google does not make it easy. So now developers choose between hand-rolling custom fragments versus a third party library. And then of course others are writing dialogs as custom views and probably enjoy the benefits of flexibility at the cost of a higher learning curve or just having to know how to "do it right". And then there's the topic of dialogs and how they fit into app architecture, which can vary based on how you model your dialogs. I get the sense it is still a wild west on this one? Are we converging on something?
## [5][Jump Into Chrome Custom Tabs](https://www.reddit.com/r/androiddev/comments/ivrrkw/jump_into_chrome_custom_tabs/)
- url: https://www.reddit.com/r/androiddev/comments/ivrrkw/jump_into_chrome_custom_tabs/
---
In this article, you'll jump into the Chrome Custom Tabs and learn performance tips to speed up ⚡ URL loading almost by 50%.

Enjoy 💚[https://vladsonkin.com/jump-into-chrome-custom-tabs/](https://vladsonkin.com/jump-into-chrome-custom-tabs/)
## [6][What is the best way to make a mobile app?](https://www.reddit.com/r/androiddev/comments/ivrd8w/what_is_the_best_way_to_make_a_mobile_app/)
- url: https://www.reddit.com/r/androiddev/comments/ivrd8w/what_is_the_best_way_to_make_a_mobile_app/
---
I recently started learning python, made it all the way to classes and object orienting, then I suddenly had this very great idea, but it needs to be a mobile app, so which is the best way to make one? ( I prefer if you guys might guide me to a way that includes python and not a new language )
## [7][A small Disney project using Jetpack Compose, Dagger-Hilt based on MVVM architecture.](https://www.reddit.com/r/androiddev/comments/iv1w2g/a_small_disney_project_using_jetpack_compose/)
- url: https://v.redd.it/rs0v48rx3vn51
---

## [8][Is there any point in having a repository that just delegates method calls?](https://www.reddit.com/r/androiddev/comments/ivojht/is_there_any_point_in_having_a_repository_that/)
- url: https://www.reddit.com/r/androiddev/comments/ivojht/is_there_any_point_in_having_a_repository_that/
---
Is there any benefit in having a repository that just has a bunch of these methods and no own logic?

     suspend fun insert(task: Task) = taskDao.insert(task)
    
     suspend fun update(task: Task) = taskDao.update(task)
    
     suspend fun deleteTask(task: Task) = taskDao.deleteTask(task)
## [9][Exploring Jetpack Compose: Column](https://www.reddit.com/r/androiddev/comments/ivbj6d/exploring_jetpack_compose_column/)
- url: https://joebirch.co/android/exploring-jetpack-compose-column/
---

## [10][Paging jetpack with multiple module clean architecture](https://www.reddit.com/r/androiddev/comments/ivmfu2/paging_jetpack_with_multiple_module_clean/)
- url: https://www.reddit.com/r/androiddev/comments/ivmfu2/paging_jetpack_with_multiple_module_clean/
---
I am trying to put jetpack pagination and clean architecture together in the demo project. I stuck in domain module which is pure java/kotlin causes paging need to return flow paging data in the repository. I want domain separating from android component. How can I handle this?
## [11][Which language are you using for Android App Development ?](https://www.reddit.com/r/androiddev/comments/ivoa7h/which_language_are_you_using_for_android_app/)
- url: https://www.reddit.com/r/androiddev/comments/ivoa7h/which_language_are_you_using_for_android_app/
---


[View Poll](https://www.reddit.com/poll/ivoa7h)
## [12][Are you going to add sellers.json to your Admob account?](https://www.reddit.com/r/androiddev/comments/iv9k3h/are_you_going_to_add_sellersjson_to_your_admob/)
- url: https://www.reddit.com/r/androiddev/comments/iv9k3h/are_you_going_to_add_sellersjson_to_your_admob/
---
I know it's not mandatory (up until now, who know what requirements will be in the future), but there are 2 huge yellow and red warning notifications in my Admob account (also I've received 2 emails regarding to this).

Admob states that this will bring "transaprency" and " This will help advertisers to verify your inventory. If your information isn't made transparent, advertisers won't be able to see your business name, which might impact your revenue. "

I haven't added yet, but I'm thinking about adding it, however I'm not okay declaring my real/business name and address in the public json file.

Have you added sellers.json to your Admob settings?
