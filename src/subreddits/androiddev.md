# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/ey6mqi/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/ey6mqi/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - February 03, 2020](https://www.reddit.com/r/androiddev/comments/ey5han/weekly_questions_thread_february_03_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ey5han/weekly_questions_thread_february_03_2020/
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
## [3][Complex UI/Animations on Android — featuring MotionLayout](https://www.reddit.com/r/androiddev/comments/ey5lu3/complex_uianimations_on_android_featuring/)
- url: https://medium.com/@nikhilpanju22/complex-ui-animations-on-android-featuring-motionlayout-aa82d83b8660?source=friends_link&amp;sk=5b924ea26bc2ae4735483760f3c62409
---

## [4][Understand your players better with affinity learnings](https://www.reddit.com/r/androiddev/comments/ey49pz/understand_your_players_better_with_affinity/)
- url: https://medium.com/googleplaydev/understand-your-players-better-with-affinity-learnings-55a49532d163
---

## [5][Android 10: open file in Download folder (received via Nearby Connections API)](https://www.reddit.com/r/androiddev/comments/ey3zkf/android_10_open_file_in_download_folder_received/)
- url: https://www.reddit.com/r/androiddev/comments/ey3zkf/android_10_open_file_in_download_folder_received/
---
Hi.

I have several Android 10 devices with my app targeting api 29 and I want to capture an image from one or more devices and send it to a 'master' device which will then open the images and upload to a server, and I figured Nearby Connections API would be the perfect fit - or so I thought.

I am following [this guide](https://developers.google.com/nearby/connections/android/exchange-data) and the connection and transferring of the image file works just fine. Trouble is the file is placed in a 'Neaby' folder in 'Downloads'. I figured I should use `MediaStore.Downloads` and I would then be able to access files created by my app, but I don't get any results when doing the following:

```
requireActivity().applicationContext.contentResolver.query(
            MediaStore.Downloads.EXTERNAL_CONTENT_URI,
            arrayOf(MediaStore.Downloads.DISPLAY_NAME),
            null,
            null,
            null
        ).use { cursor -&gt;
           while (cursor.moveToNext()) {
              //Not reaching this point :/
           }
        }
```

I use `"com.google.android.gms:play-services-nearby:17.0.0"`.

Lowering target api or setting `android:requestLegacyExternalStorage="true"` are not options ;)

Am I doing something wrong or is my goal simply not possible with the current state of `play-services-nearby`?
## [6][20 Best Android Development Books of 2020](https://www.reddit.com/r/androiddev/comments/ey5sk3/20_best_android_development_books_of_2020/)
- url: https://booksawesome.com/20-best-android-development-books-of-2020/
---

## [7][Can you check the activity stack if it contains a certain activity?](https://www.reddit.com/r/androiddev/comments/ey3sru/can_you_check_the_activity_stack_if_it_contains_a/)
- url: https://www.reddit.com/r/androiddev/comments/ey3sru/can_you_check_the_activity_stack_if_it_contains_a/
---
Hello!   


I'm wondering if it is possible to check if the activity stack contains a certain activity  to know if this activity has already been launched before? Is it possible to distinguish between activities already launched and activities that have not been launched yet? 

Thank you.
## [8][Have the guidelines for back vs. up navigation changed?](https://www.reddit.com/r/androiddev/comments/ey5ebr/have_the_guidelines_for_back_vs_up_navigation/)
- url: https://www.reddit.com/r/androiddev/comments/ey5ebr/have_the_guidelines_for_back_vs_up_navigation/
---
The guidelines for back and up navigation used to look[ like this. (Internet Archive)] (https://web.archive.org/web/20190204225959/https://developer.android.com/training/design-navigation/ancestral-temporal)

However, that same page now redirects to a new page:
https://developer.android.com/training/design-navigation/ancestral-temporal

There is also this new page:

https://developer.android.com/guide/navigation/navigation-principles#up_and_back_are_identical_within_your_apps_task


Previously, it was recommended that the back button and the up button should not always have the same behaviour. However, the new advice is:
&gt; Within your app's task, the Up and Back buttons behave identically.

This seems like a change for the better, but I can't find any announcement for it, which seems strange for such a big change.
## [9][LocallyDynamic: Dynamic Delivery without Google Play - download and install your split APK:s on any device without having to upload app bundles to Google Play](https://www.reddit.com/r/androiddev/comments/exn19l/locallydynamic_dynamic_delivery_without_google/)
- url: https://github.com/jeppeman/LocallyDynamic
---

## [10][App development (IOS + Android)](https://www.reddit.com/r/androiddev/comments/exzy7f/app_development_ios_android/)
- url: https://www.reddit.com/r/androiddev/comments/exzy7f/app_development_ios_android/
---

I am new to programming. I want to build an app for both IOS and Android for business purposes. I know there are platforms like flutter where you can build apps for both ios and android. But, my concern is how sustainable would be using platforms like flutter instead developing same app for ios and Android separately. I appreciate your feedback
## [11][Best way to manage data in Firestore as an admin ?](https://www.reddit.com/r/androiddev/comments/ey554f/best_way_to_manage_data_in_firestore_as_an_admin/)
- url: https://www.reddit.com/r/androiddev/comments/ey554f/best_way_to_manage_data_in_firestore_as_an_admin/
---
Say I have an app all is setup I just want to filter some data before it goes into the main set manually. Should one make an admin version of the app or use a script on a PC or use cloud functions but then where comes the human decision in functions. Any insight would be appreciated.
## [12][Is Picasso Library abandoned?](https://www.reddit.com/r/androiddev/comments/expyvb/is_picasso_library_abandoned/)
- url: https://www.reddit.com/r/androiddev/comments/expyvb/is_picasso_library_abandoned/
---
The last Picasso library release is about [two years old](https://github.com/square/picasso/releases)

A long period considering Square updates its libraries very often

I'm considering to use Coil in new projects (migrating the older projects is not necessary because Picasso works very fine for me)
