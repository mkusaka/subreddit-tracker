# androiddev
## [1][App Feedback Thread - February 15, 2020](https://www.reddit.com/r/androiddev/comments/f48scg/app_feedback_thread_february_15_2020/)
- url: https://www.reddit.com/r/androiddev/comments/f48scg/app_feedback_thread_february_15_2020/
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
## [2][Weekly Questions Thread - February 10, 2020](https://www.reddit.com/r/androiddev/comments/f1opq1/weekly_questions_thread_february_10_2020/)
- url: https://www.reddit.com/r/androiddev/comments/f1opq1/weekly_questions_thread_february_10_2020/
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
## [3][Properly manage FragmentTransaction backstack](https://www.reddit.com/r/androiddev/comments/f4nvgt/properly_manage_fragmenttransaction_backstack/)
- url: https://www.reddit.com/r/androiddev/comments/f4nvgt/properly_manage_fragmenttransaction_backstack/
---
If I have a host activity \`MainActivity\` and I first set up a \`MainMenuFragment\` then when i click the play button it opens up the \`GameFragment\` . If I click a back button i want it to take me back to the \`MainMenuFragment\` .  I don't want to do a \`FragmentTransaction.replace()\` when I go back to the \`MainMenuFragment\` because I don't want it to reload the \`MainMenuFragment\` , i want it to preserve the state of the \`MainMenuFragment\`

&amp;#x200B;

So i tried to set this all up with calls to \`addToBackStack()\` and \`popBackStack()\` but that doesnt work because I can see that \`FragmentTransaction.backStackEntryCount\` keeps incrementing as I navigate from the \`MainMenuFragment\` to the \`GameFragment\` and back and forth.

&amp;#x200B;

Is there a straightforward approach to this?
## [4][Accidentally Released App to All Users, how to revert to closed Alpha?](https://www.reddit.com/r/androiddev/comments/f4mrmq/accidentally_released_app_to_all_users_how_to/)
- url: https://www.reddit.com/r/androiddev/comments/f4mrmq/accidentally_released_app_to_all_users_how_to/
---
yes, title.

Stupid me did this when I tried to update my alpha apk, I updated it the right way but there was this HUGE blue button saying "Release Production" so I thought I had to click it to update the apk.

Silly me it got into play Store and now I had to force unpublish it.

But now I am not sure if my alpha opt-in link is working or it is even able to update as the app says it's unpublished.

* How do I remove this Production track? I think this is what makes the APP appear on the play store

&amp;#x200B;

https://preview.redd.it/7jivlidvb8h41.png?width=975&amp;format=png&amp;auto=webp&amp;s=d006f2b81fa15b4f90cabb7535c68159f51bccf3

edit: I updated alpha app and it seems it did update and appears on play store (only for users who have installed the alpha version before) but I wonder how I am going to share the app for new users without publishing... any help is appreciated.
## [5][why google play games (for 2048 ) leader board always hacked/broken? Do developers not able to solve security issue or bug these cheater implement.](https://www.reddit.com/r/androiddev/comments/f4bdpa/why_google_play_games_for_2048_leader_board/)
- url: https://i.redd.it/wqdrfhgp24h41.jpg
---

## [6][Recyclerview parent content](https://www.reddit.com/r/androiddev/comments/f4mtkj/recyclerview_parent_content/)
- url: https://www.reddit.com/r/androiddev/comments/f4mtkj/recyclerview_parent_content/
---
Hello guys,

need some help with recyclerview. I have an object "Order" which has a Date property and a list of products for that order. The current behaviour I have is that recyclerview shows a list of products and if you tap on one of them it opens another list which shows a list of products for that order. I would like to get rid of that second screen which shows recyclerview of products. Instead, I would like to use just one recyclerview which shows a list of products. Each card shows a date of order and bellows it would output the products of that order. Any suggestions on how can I achieve that?
## [7][I've created an example (experiment?) using vanilla Dagger, @Subcomponents with Application -&gt; Activity -&gt; Fragment scoping, and ViewModel injection with SavedStateHandle using AutoFactory (to survive process death and config changes properly even with `@Inject lateinit var viewModel: MyViewModel`)](https://www.reddit.com/r/androiddev/comments/f4f8sh/ive_created_an_example_experiment_using_vanilla/)
- url: https://github.com/Zhuinden/DaggerViewModelExperiment/blob/master/app/src/main/java/com/zhuinden/daggerviewmodelexperiment/features/first/FirstFragment.kt#L20-L58
---

## [8][Help random ads keep popping in screen over any app. Deleted most recent app and they still ain't go away.](https://www.reddit.com/r/androiddev/comments/f4qcm6/help_random_ads_keep_popping_in_screen_over_any/)
- url: https://i.redd.it/x4mycfbkaah41.jpg
---

## [9][Commonsware Blog: Scoped Storage Stories: The Undocumented Documents](https://www.reddit.com/r/androiddev/comments/f4el3e/commonsware_blog_scoped_storage_stories_the/)
- url: https://commonsware.com/blog/2020/02/15/scoped-storage-stories-undocumented-documents.html
---

## [10][How is it possible for AutoCompleteTextView's item selection listener to return the wrong item when user selects an item?](https://www.reddit.com/r/androiddev/comments/f4phy4/how_is_it_possible_for_autocompletetextviews_item/)
- url: https://www.reddit.com/r/androiddev/comments/f4phy4/how_is_it_possible_for_autocompletetextviews_item/
---
Hello! I was wondering if that's ever happened to you?
## [11][I want to show images uploaded on firebase storage in google drive in custom tab. Is it even possible?](https://www.reddit.com/r/androiddev/comments/f4p72z/i_want_to_show_images_uploaded_on_firebase/)
- url: https://www.reddit.com/r/androiddev/comments/f4p72z/i_want_to_show_images_uploaded_on_firebase/
---
I am working on an app, in which, i am uploading images to firebase  in jpg. Then when user request for preview of that image. so it should  be shown in customTab of google drive as it provide interface to display  images and stuff. I want know, is this even possible or i have to  upload images to google drive directly? Any suggestion will be helpful.

Thanks in advance.
## [12][What's the difference between ApkBuilder, apkzlib and zipflinger/signflinger?](https://www.reddit.com/r/androiddev/comments/f4ozil/whats_the_difference_between_apkbuilder_apkzlib/)
- url: https://www.reddit.com/r/androiddev/comments/f4ozil/whats_the_difference_between_apkbuilder_apkzlib/
---
App Inventor currently uses ApkBuilder to add dex files and native libs to the APK file. However, as ApkBuilder is officially removed last year and in the process of upgrading the build tools that App Inventor uses to compile APK files, I am looking for replacement for ApkBuilder.

I also asked this a year ago to no avail
https://www.reddit.com/r/androiddev/comments/a9cj2z/what_is_the_alternative_to_apkbuilder/

Is Zipflinger more efficient and faster than the former two?
