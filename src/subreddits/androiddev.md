# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/ibdg88/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/ibdg88/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - August 17, 2020](https://www.reddit.com/r/androiddev/comments/ibbyvc/weekly_questions_thread_august_17_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ibbyvc/weekly_questions_thread_august_17_2020/
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
## [3][Android 11 is taking away the camera picker, forcing people to only use the built-in camera](https://www.reddit.com/r/androiddev/comments/ici93p/android_11_is_taking_away_the_camera_picker/)
- url: https://www.reddit.com/r/androiddev/comments/ici93p/android_11_is_taking_away_the_camera_picker/
---
Articles:

- https://www.androidpolice.com/2020/08/18/android-11-camera-apps-chooser/

- https://www.androidauthority.com/android-11-camera-restrictions-1149232/

- https://9to5google.com/2020/08/18/android-11-default-camera-app-changes/

CommonsWare article:

- https://commonsware.com/blog/2020/08/16/action-image-capture-android-r.html

Saw this recently on reddit too:

- https://www.reddit.com/r/Android/comments/icbcgj/android_11_is_taking_away_the_camera_picker/g22nh2d/?context=3

- https://www.reddit.com/r/android_devs/comments/iba7p7/action_image_capture_and_android_r/


Related issues on the issue tracker (consider starring if you want to revert this change ) :

- https://issuetracker.google.com/issues/164794531
- https://issuetracker.google.com/issues/165337020
- https://issuetracker.google.com/issues/165313053
- https://issuetracker.google.com/issues/165313060
## [4][Handling strings (placeholders) - Phrase (square lib) still recommended?](https://www.reddit.com/r/androiddev/comments/ick0js/handling_strings_placeholders_phrase_square_lib/)
- url: https://www.reddit.com/r/androiddev/comments/ick0js/handling_strings_placeholders_phrase_square_lib/
---
Hi everyone,

we stumbled across a problem with string placeholders and translations.we wanted to have a string like "15min 30sec"so we'd define `&lt;string name="duration"&gt;%1$smin %2$ssec&lt;/string&gt;`

our translation tool is not capable of properly identifying placeholders if the do not have a trailing whitespace and are followed by an alphabetical character. so i found [https://developer.android.com/guide/topics/resources/localization#mark-message-parts](https://developer.android.com/guide/topics/resources/localization#mark-message-parts)  
maybe it's interesting to some of you which are dealing with a similar problem.  


next to that documentation i found that there's a more convenient way of handling placeholders with the lib by square: [https://github.com/square/phrase](https://github.com/square/phrase)  
you can define strings like: `&lt;string name="duration"&gt;{minutes}min {seconds}sec&lt;/string&gt;`  
and the bind the named params in code. also i placeholders are much easier to digest for external people (translators)

as it's not heavily maintained I wanted to ask if sbdy is (still) using it.
## [5][Activities are not what you think](https://www.reddit.com/r/androiddev/comments/icdle7/activities_are_not_what_you_think/)
- url: https://gabrielfv.com/android/architecture/en-us/2020/08/18/activities-are-not-what-you-think-pt-1.html
---

## [6][Is RealmRecyclerViewAdapter considered usable in terms of clean code/ separation of concerns?](https://www.reddit.com/r/androiddev/comments/iclw15/is_realmrecyclerviewadapter_considered_usable_in/)
- url: https://www.reddit.com/r/androiddev/comments/iclw15/is_realmrecyclerviewadapter_considered_usable_in/
---
Hi, 

I'm creating app in which I use Realm as my (currently) only data source. I was wondering if using RealmRVAdapter is good practice for displaying list of objects. I'm concerned about it preventing me from closing realm instance after performing query because of RealmChangeListeners used in RealmRV. 

So here few questions:

Does anyone use RealmRV in apps with MVVM and Repository pattern? If so how do you manage your Realm instance? Is keeping Realm open is considered good/passable option? Or should I ditch RealmRV and handle updates in RecyclerView on my own?
## [7][Game Translation Fails — How to Avoid Them ?](https://www.reddit.com/r/androiddev/comments/icnl63/game_translation_fails_how_to_avoid_them/)
- url: https://www.leveluptranslation.com/single-post/game-translation-fails-causes-and-how-to-avoid-them
---

## [8][`pagingDataAdapterOf` builder function for `PagingDataAdapter` in Kiel 1.2.0](https://www.reddit.com/r/androiddev/comments/icktiq/pagingdataadapterof_builder_function_for/)
- url: https://www.reddit.com/r/androiddev/comments/icktiq/pagingdataadapterof_builder_function_for/
---
Hi Android Fellas,

`pagingDataAdapterOf` builder function for `PagingDataAdapter`,which is alpha now, has been introduced in the version 1.2.0 of kiel.

github: https://github.com/ibrahimyilmaz/kiel

Example usage: 
```kt
    private val adapter = pagingDataAdapterOf&lt;EntityListItem&gt; {
         diff(
             areItemsTheSame = { old, new -&gt; old.entityData.id == new.entityData.id },
             areContentsTheSame = { old, new -&gt; old == new }
         )
         register(
             layoutResource = R.layout.item_entity,
             viewHolder = ::EntityListItemViewHolder,
             onBindBindViewHolder = { viewHolder, _, item -&gt;
                 viewHolder.itemView.setOnClickListener { viewModel.select(item) }
                 viewHolder.binding.entityTitleTextView.text = item.title

                 viewHolder.binding.entityRemoveRelationAction.isVisible = item.isRemoveAvailable
                 if (item.isRemoveAvailable) {
                     viewHolder.binding.entityRemoveRelationAction.setOnClickListener {
                         viewModel.removeRelation(item)
                     }
                 }
             }
         )
     }
```

Any feedback, idea or PR is always welcome

Sincerely
Ibra
## [9][New language features and more in Kotlin 1.4](https://www.reddit.com/r/androiddev/comments/ibxm5o/new_language_features_and_more_in_kotlin_14/)
- url: https://android-developers.googleblog.com/2020/08/new-language-features-and-more-in.html
---

## [10][Android Activity Life Cycle - Some important scenarios!](https://www.reddit.com/r/androiddev/comments/ichv1q/android_activity_life_cycle_some_important/)
- url: https://ayusch.com/android-activity-life-cycle-the-complete-guide/
---

## [11][What Role Does Mobile App Development Company Play in Boosting The Start-Ups?](https://www.reddit.com/r/androiddev/comments/icjxxh/what_role_does_mobile_app_development_company/)
- url: https://www.metizsoft.com/blog/role-of-mobile-app-development-company
---

## [12][Firebase - could you please help me which modules to use for my use cases?](https://www.reddit.com/r/androiddev/comments/icjpmq/firebase_could_you_please_help_me_which_modules/)
- url: https://www.reddit.com/r/androiddev/comments/icjpmq/firebase_could_you_please_help_me_which_modules/
---
I am starting to explore firebase, but it seems very robust, complex system with many modules and I dont even know where to start.

Could you please point me in the right direction? **What firebase modules do I need for these use cases:**

* once in few weeks I need to send an important message to user's devices (for example about a launched discount, or about some critical update) *I want to have a full control over message style (sometimes I want to show just a Toast, for more importane messages simple dialog and in very urgent cases status bar notification)* - **should I use firebase in-app messaging, or...?**

*  my app has highly customizable look, so i would like to periodically sync user settings, so app always look same on all devices - **should I use firebase remote config?**

* I would like to make some simple polls with A/B/C options once in a while (ask users what features should I add, or other feedback) - **no idea about this...what module to use?**

Then of course I wanna use analytics and crashlytics, but thats easy to figure out.
