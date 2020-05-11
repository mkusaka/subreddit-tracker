# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/ghmv5w/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/ghmv5w/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - May 11, 2020](https://www.reddit.com/r/androiddev/comments/ghlel5/weekly_questions_thread_may_11_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ghlel5/weekly_questions_thread_may_11_2020/
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
## [3][Build android apps entirely in C](https://www.reddit.com/r/androiddev/comments/ghaxyb/build_android_apps_entirely_in_c/)
- url: https://github.com/cnlohr/rawdrawandroid
---

## [4][Live Data - One state vs multiple live data?](https://www.reddit.com/r/androiddev/comments/ghjtnq/live_data_one_state_vs_multiple_live_data/)
- url: https://www.reddit.com/r/androiddev/comments/ghjtnq/live_data_one_state_vs_multiple_live_data/
---
Hi,

I've been experimenting a bit with live data and I'm running into a problem regarding complex data. Right now, I try to encapsulate the entire view state into one data class object and push copies of the current state with changes to have a single source of truth. 

This works well, but as the views get more complex, it seems more and more cumbersome, and I was thinking about splitting my live data into multiple groups, eg. one only for the error messages, etc. 

My google skills failed me to find a consistent opinion online since most blogs mostly deal with very simple views ... what would you say are best practices to model complex view state with live data?
## [5][An open source minimalistic notes app (Link in comments)](https://www.reddit.com/r/androiddev/comments/ggz0ul/an_open_source_minimalistic_notes_app_link_in/)
- url: https://i.redd.it/7tk00j8sywx41.png
---

## [6][Performance advice](https://www.reddit.com/r/androiddev/comments/ghix57/performance_advice/)
- url: https://www.reddit.com/r/androiddev/comments/ghix57/performance_advice/
---
I've been developing a run tracking application and as part of that I calculate the distance of a route consisting of a bunch of coordinates.

Initially I did my distance calculations in my view model. I'm aware that isn't ideal but it's easy and trivial to move it down to the repository.

The calculation is done by retrieving a list of coordinates.

https://github.com/tonsV2/open-run-tracker/blob/master/app/src/main/java/dk/fitfit/runtracker/viewmodels/LocationUpdateViewModel.kt#L70

And when that list change I calculate the distance in meters.

https://github.com/tonsV2/open-run-tracker/blob/master/app/src/main/java/dk/fitfit/runtracker/viewmodels/LocationUpdateViewModel.kt#L74

But as my app has evolved I need the distance in other places. One of those places is the notification showing progress which I intent to implement.

So instead of recalculating the distance for the sake of the notification I thought I move the calculation to a central place and pull from that. 

Basically my idea was to continuously recalculate the distance as new coordinates arrive. And store the product in the run table of my database.

In order to avoid race conditions I thought I'd use a `WorkManager` and enqueueUniqueWork.

https://github.com/tonsV2/open-run-tracker/blob/feature/centralized_distance_calculation/app/src/main/java/dk/fitfit/runtracker/LocationUpdatesBroadcastReceiver.kt#L60

And my view model now observe changes to the current run instead of the locations table.

https://github.com/tonsV2/open-run-tracker/blob/feature/centralized_distance_calculation/app/src/main/java/dk/fitfit/runtracker/viewmodels/LocationUpdateViewModel.kt#L66

However the above has incurred a huge performace penalty! I was aware that  additional database reads/writes would slow down things. But not as much as it did. So I have the following questions.

1. Am I doing something wrong? I'm surprised things are slowing down this much

2. Is there a better way to go about this?

3. The actual calculations are not that intensive. But I don't like to do the same work in several places. However as of now that seems like to easy and most performant solution

Thanks for reading this far :-)
## [7][The Lowest quality (below zero) of Google Play's Developers Support](https://www.reddit.com/r/androiddev/comments/gh8msz/the_lowest_quality_below_zero_of_google_plays/)
- url: https://www.reddit.com/r/androiddev/comments/gh8msz/the_lowest_quality_below_zero_of_google_plays/
---
[Some time ago one of my apps was suspended by Google's AI](https://www.reddit.com/r/androiddev/comments/g6poux/google_stores_ai_suspends_apps_without_any/).

I've applied several appeals. And one of them was awarded an answer.

In my appeal, I described all the required information and ask to change the statement of the App from "Suspended" to "Deleted". 

When an app is "deleted", you could update an APK for re-moderation.

[Suspended, Unpublished, and Removed states](https://preview.redd.it/f0j625q0wzx41.png?width=1117&amp;format=png&amp;auto=webp&amp;s=35ed8c9024e337a47e76459c670461cfc57a0b5a)

When an app is "Suspended", you have no opportunities to submit an update at all.

[In suspended state no options to submit an update.](https://preview.redd.it/s500weh9wzx41.png?width=906&amp;format=png&amp;auto=webp&amp;s=132a63c03690602d0878d62655b0e4ec35ada495)

I've even uploaded a new version of APK to Google Drive and shared a link.

But a person didn't even notice, that I have no right to upload a new APK in developers panel.

He just informed me, that he had reviewed my old APK file and had found that it wasn't updated and still contains disputable assets. 

I've even sent him this screenshot with the inactive "Update" button. But have no answer...

**How people, who even don't know their product, could make decisions???**

**People, who don't have enough technical knowledge or motivation to understand your case just able to decide the future of Google Play Store.**

&amp;#x200B;

[They've reviewed not updated APK and haven't found fixes...](https://preview.redd.it/xxdcx1nfuzx41.png?width=1110&amp;format=png&amp;auto=webp&amp;s=2b2789c98cb148c2f987403771ba001025f94d42)

In the previous [tread](https://www.reddit.com/r/androiddev/comments/g6poux/google_stores_ai_suspends_apps_without_any/fpxg2yg/), one developer said, that Google's moderator asked him to prove ownership of [google.com](https://google.com)

If you are still working with Google, you trust your Business to incompetent contractors.
## [8][Play Console not reporting stats this month?](https://www.reddit.com/r/androiddev/comments/ghiqbp/play_console_not_reporting_stats_this_month/)
- url: https://www.reddit.com/r/androiddev/comments/ghiqbp/play_console_not_reporting_stats_this_month/
---
Is anyone else seeing issues with the Play Console not reporting *any* stats for May so far? 

I can see in analytics that people are downloading my app but can't see any rating/download/uninstall/revenue stats in the Play Console at all for May..

Does this happen a lot?
## [9][How do i build this kind of App in Android Studio](https://www.reddit.com/r/androiddev/comments/ghmuvq/how_do_i_build_this_kind_of_app_in_android_studio/)
- url: https://www.reddit.com/r/androiddev/comments/ghmuvq/how_do_i_build_this_kind_of_app_in_android_studio/
---
I'm new to programming and never made an App. I've watched tutorials and learning but i don't know the right way to pursue making the app i want to make

How do i make a messenger app with posts, for example like 9gag with a **messenger** feature with it. I'm not sure if there's an official name for this kind of app.
## [10][How important is hardware for development](https://www.reddit.com/r/androiddev/comments/ghmu9z/how_important_is_hardware_for_development/)
- url: https://www.reddit.com/r/androiddev/comments/ghmu9z/how_important_is_hardware_for_development/
---
Yes I did Google and research a bit before asking but I still have a vague understanding and don't want to spend more than I have to -- nor do I want to spend too little and be disappointed.

I'm a novice programmer just starting out in Kotlin and I want to purchase a laptop. Windows not Mac.



So i5 vs i7 has gotten conflicting answers. Some say i7 is significantly better others say i7 is negligible.

How much of a difference does 8gb of ram have vs 16gb? Will I notice ? Does it sloe doen compilation by seconds? Minutes? Does it affect the android emulator ?

Do you believe 240gb sad is enough for development or will I eventually run out of space easily and opt for a larger ssd? (This one's less important I guess since I can move stuff to an external)

And finally how important is a You? If I get a laptop without a dedicated gpu will it affect the android emulator ? 


Bonus UNITY question if anyone knows:
In the future I'd like to dabble in mobile game development.

Will integrated graphics suffice for some basic unity engine game making I won't be rendering 3d models on my laptop I have a PC with blender for that. I'd just like to be able to work on projects on the go if need be.
## [11][Proper replacement for AsyncTask - onProgressUpdate animation](https://www.reddit.com/r/androiddev/comments/ghjlrb/proper_replacement_for_asynctask_onprogressupdate/)
- url: https://www.reddit.com/r/androiddev/comments/ghjlrb/proper_replacement_for_asynctask_onprogressupdate/
---
Hello, 

I am sort of very new to android and one of the tasks for me is to create proper replacement for asynctask... especially to integrate our custom indeterminate "progress bar" which consist of two animations with our company logo very similar thing can be seen [here](https://tobiasahlin.com/spinkit/)... I work with object animator and two parts of animation. So my java class is basically done and just one method of this object needs to be called in interval (around 1 second) of each two parts of animation. How can i achieve this?

I came to rxjava which would have just handled the background task, but i really don't know how to integrate this animation directly into rxjava background task...  

&amp;#x200B;

    class LoadCatalogsTask implements Callable&lt;Boolean&gt;, Observer&lt;Boolean&gt; {
            @Override
            public void onSubscribe(@NonNull Disposable d) {
                ourLogo.setVisibility(View.VISIBLE);
            }
    
            @Override
            public void onNext(@NonNull Boolean aBoolean) {
                ourLogo.setVisibility(View.GONE);
                Log.d("s", "onNext: ");
            }
    
            @Override
            public void onError(@NonNull Throwable e) {
                ourLogo.setVisibility(View.GONE);
                Log.e("s", "onError: " + "s");
            }
    
            @Override
            public void onComplete() {
                Toast.makeText(MainActivity.this, "onComplete", Toast.LENGTH_SHORT).show();
            }
    
            @Override
            public Boolean call() throws Exception {
                return doAllWorkInBackground();
            }
    
        }

I would be very grateful for any suggestions eventually even
## [12][My App was suspended from the Play Store - Need help resubmitting it](https://www.reddit.com/r/androiddev/comments/ghjllo/my_app_was_suspended_from_the_play_store_need/)
- url: https://www.reddit.com/r/androiddev/comments/ghjllo/my_app_was_suspended_from_the_play_store_need/
---
Hey there,

So about 10 days ago my app with 50k active users was suspended from the Play Store because of Copyright violations. If you're interested in what happened exactly you can read it up [here](https://www.reddit.com/r/GoCompanion/comments/gau1xx/go_companion_has_been_suspended_from_the_google/), but the tldr is I was stupid and used a Pikachu artwork in one of my Play Store images.

And so although I still think it is brutal to immediately suspend my app without notice and without a chance for me to fix this issue I accepted their action as it was completely my fault.

Now my problem is the following:

Google Support told me that if my app does not violate any copyright law and only the Play Store images were the problem I can resubmit the app with a slightly different name and different package name and get my app up and running that way again.

So I made sure that my app won't violate any copyright law and did exactly that. But Google is since then rejecting every submission I make to this new app. I have changed the name, the description, I really have tried everything. But Google is only sending me one rejection after the other with the extreme wage information that my app violates "Google Play Policies". No further explanation.

My question now is, if anyone has had a similar issue and if anyone knows what I can do know in this situation?
