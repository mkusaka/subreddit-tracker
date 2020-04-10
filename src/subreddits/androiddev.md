# androiddev
## [1][Weekly "anything goes" thread!](https://www.reddit.com/r/androiddev/comments/fyf6ck/weekly_anything_goes_thread/)
- url: https://www.reddit.com/r/androiddev/comments/fyf6ck/weekly_anything_goes_thread/
---
Here's your chance to talk about whatever!

Although if you're thinking about getting feedback on an app, you should wait until tomorrow's App Feedback thread.

Remember that while you can talk about any topic, being a jerk is [still not allowed](https://www.reddit.com/r/androiddev/wiki/rules#wiki_rules_for_comments).
## [2][Weekly Questions Thread - April 06, 2020](https://www.reddit.com/r/androiddev/comments/fvwq7t/weekly_questions_thread_april_06_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fvwq7t/weekly_questions_thread_april_06_2020/
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
## [3][It bothers me that most tutorials and or architectural guides don’t include any kind of automated testing](https://www.reddit.com/r/androiddev/comments/fychih/it_bothers_me_that_most_tutorials_and_or/)
- url: https://www.reddit.com/r/androiddev/comments/fychih/it_bothers_me_that_most_tutorials_and_or/
---
No unit tests, no UI tests, nothing. It seems nobody cares :-(

It also makes it hard to learn more about the subject and I don’t know about you guys but I feel like knowing how to test stuff is what separates you from a median developer.
## [4][What is wrong with you Samsung?](https://www.reddit.com/r/androiddev/comments/fydhp4/what_is_wrong_with_you_samsung/)
- url: https://www.reddit.com/r/androiddev/comments/fydhp4/what_is_wrong_with_you_samsung/
---
I just want to share this unbelievable story...

As usual, I've recently submitted an update of my apps on the Galaxy Apps Store.

&amp;#x200B;

But I just got an email from Samsung to inform me that after the review, my apps have been suspended / removed.

It's not only the update, it's a complete removing from the store.

&amp;#x200B;

Ok, well, let's see the reason of the suspension...

\&gt;&gt; Unable to attach contents in Facebook.

&amp;#x200B;

It's so funny, they even sent me a video of the problem.

[Unable to attach contents in Facebook](https://reddit.com/link/fydhp4/video/qy1iigx0wyr41/player)

The export feature is just a text sharing through a simple Intent.

Do I must test all the apps on the market to check if they can handle Intents correctly ?

&amp;#x200B;

I don't know if I must laugh or cry...

I'm trying to stay polite, but what a joke!!!

&amp;#x200B;

I re-submitted the apps today with absolutely no changes at all.

I just hope the review will be handled by someone else, someone really qualified for the job...

&amp;#x200B;

EDIT:

As I thought, my re-submit has been approved by someone else and now my apps are back online on the Galaxy Apps Store.

&amp;#x200B;

Thanks to [yo\_asakura](https://www.reddit.com/user/yo_asakura/) for the clarification about the Facebook SDK.
## [5][How could this be?](https://www.reddit.com/r/androiddev/comments/fxx5ed/how_could_this_be/)
- url: https://i.redd.it/q4pupretztr41.jpg
---

## [6][New Database Inspector in Android Studio 4.1 Canary 5](https://www.reddit.com/r/androiddev/comments/fxz3v7/new_database_inspector_in_android_studio_41/)
- url: https://developer.android.com/studio/preview/features#database-inspector
---

## [7][Google Play has suspended our air quality monitoring app saying it's related to Covid-19](https://www.reddit.com/r/androiddev/comments/fxvupb/google_play_has_suspended_our_air_quality/)
- url: https://www.reddit.com/r/androiddev/comments/fxvupb/google_play_has_suspended_our_air_quality/
---
Our app is not related to Covid-19 whatsoever. It's an air quality and weather conditions monitoring app. Yet, Google has suspended it(for the second time now) citing their sensitive events policy. Has this happened to anyone else? Can anyone here suggest what we can do to get our app reinstated?
## [8][Do you / Should I use ViewBinding over kotlin synthetics?](https://www.reddit.com/r/androiddev/comments/fyc754/do_you_should_i_use_viewbinding_over_kotlin/)
- url: https://www.reddit.com/r/androiddev/comments/fyc754/do_you_should_i_use_viewbinding_over_kotlin/
---

## [9][Google is still recommending AsyncTasks in their documentation](https://www.reddit.com/r/androiddev/comments/fy7gdo/google_is_still_recommending_asynctasks_in_their/)
- url: https://www.reddit.com/r/androiddev/comments/fy7gdo/google_is_still_recommending_asynctasks_in_their/
---
I know Google's documentation has never been great but I found it especially egregious that they are still recommending using an AsyncTask to avoid ANRs:

https://developer.android.com/training/articles/perf-anr.html
## [10][EditText and focus](https://www.reddit.com/r/androiddev/comments/fye8yq/edittext_and_focus/)
- url: https://www.reddit.com/r/androiddev/comments/fye8yq/edittext_and_focus/
---
Whats the simplist way to implement a `EditText` that loses focus consistently? 

I've run into a problem where an `EditText` isn't able to lose focus which gives me no way to indicate when to remove the soft keyboard, this is by design as I understand it. 

My current application state is an `Activity` with 2 `Fragments` side-by-side and a `Toolbar` along the top, in the right side `Fragment` I have a `RecyclerView` which contains an `EditText` field, when I click/tap on the `EditText` a soft keyboard appears, this is working as intened.

The problem is that whenever I try to click/tap on any other `View` the focus on the `EditText` persists and I am therefore unable to hide the soft keyboard when the focus changes.

I am only able to set another view to have focus if I do the following on every `View` in the `Fragment`; 

    v.setOnClickListener {
    v.isFocusable = true
    v.isFocusableInTouchMode = true
    v.requestFocus()
    }

This doesn't help if a user clicks/taps on the other `Fragment` or the `Toolbar`, I would have to add this method to every `View` in the `Activity` for this behvaiour to propagate. This seems like an incredibly inefficient way to handle my problem.

Is there something simple that im missing to enable transfering of focus on `EditText` `Views`?
## [11][From your experience, is android dev more fun than web dev?](https://www.reddit.com/r/androiddev/comments/fy05lw/from_your_experience_is_android_dev_more_fun_than/)
- url: https://www.reddit.com/r/androiddev/comments/fy05lw/from_your_experience_is_android_dev_more_fun_than/
---

## [12][Android video conferencing, and rarely discussed pitfalls: Picture-in-picture and bluetooth headset](https://www.reddit.com/r/androiddev/comments/fydv1y/android_video_conferencing_and_rarely_discussed/)
- url: https://halilibo.com/2020/incident-command-center.html
---

