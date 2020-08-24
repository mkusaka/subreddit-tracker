# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/ifo98o/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/ifo98o/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - August 24, 2020](https://www.reddit.com/r/androiddev/comments/ifmsyv/weekly_questions_thread_august_24_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ifmsyv/weekly_questions_thread_august_24_2020/
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
## [3][loggy.sh - Minimalist Live Log Tail for Android, WIP](https://www.reddit.com/r/androiddev/comments/ifcfw6/loggysh_minimalist_live_log_tail_for_android_wip/)
- url: https://v.redd.it/g001rzifsti51
---

## [4][Starting as an Android Developer.](https://www.reddit.com/r/androiddev/comments/ifmj1h/starting_as_an_android_developer/)
- url: https://www.reddit.com/r/androiddev/comments/ifmj1h/starting_as_an_android_developer/
---
Where should I start? What would you recommend to someone who has never dealt with making android apps before? Are android apps still written in Java or are they being written in Kotlin now? Where do I begin?
## [5][Best practices for a Production/Test Environment with Android/Firebase/Firestore](https://www.reddit.com/r/androiddev/comments/ifnhea/best_practices_for_a_productiontest_environment/)
- url: https://www.reddit.com/r/androiddev/comments/ifnhea/best_practices_for_a_productiontest_environment/
---
Hello,

first of all, I'm a newb self taught programmer. Into this journey about a year (overall learning) and have been working on a project (for about 6-7 months) for my father's mid sized construction company. (specialty contractor where there are no good-fit solutions in existence). Thank you for reading and any input/feedback is much appreciated.

The app is built in android/java and uses Firebase including Firestore, Authentication, Cloud Functions (NodeJS) and Storage.

I'm nearing the point where (what I call) my "Sales Module" is ready for some (live) testing to collect some user feedback. The app features a login system (Authentication) and it's database tracks users, employees, sales items,  customers, sales leads, appointments and ultimately estimates and proposals. (and more) Now before I jump into the next phase of development (creating projects, work orders, inventory/material management, time clock/sheets, etc.)

I need to get this module tested and collect some feedback on it. And of course make sure everything works as intended so far, as there are probably a good amount of untested case scenarios, that are overlooked. We have singled out a couple users that will test drive what we have so far.

&amp;#x200B;

So far I've only been working on this all by myself, so far so good, but....

a) I need to create a live running version  (production) in which I can also fix errors and bugs.

b) I would also like to continue developing on future features without interfering with the "production environment"

&amp;#x200B;

I assume I need to create an additional Google Cloud Project (a la "Production Version/Dev Version')? In theory that's no problem, but I am worried about code inconsistencies (when starting to troubleshoot in the production environment).  Keep in mind I have my Android App Code, but also my Firebase stuff including Firestore DB, Cloud Functions, Storage, etc.

The code in my "Dev Version" should not be released until those features are ready. Or should I just keep one single development code and make newer features inaccessible but still compiling/building them? Please pardon my lack of knowledge and use of proper terms when it comes to this, as I have no experience whatsoever in this particular department. Any help/ideas/suggestions are greatly appreciated, especially articles/lecture/tutorials, etc.

Thank you!

&amp;#x200B;

TL;DR:

&amp;#x200B;

How do I (or do I even need to?) separate my production and development environment when creating an android app (Java) with a fairly complex Google Cloud Firebase project (Firestore, Cloud Functions, Authentication, Storage, etc.) ?
## [6][Google Play ratings are not average based and lack transparency](https://www.reddit.com/r/androiddev/comments/if5l7h/google_play_ratings_are_not_average_based_and/)
- url: https://www.reddit.com/r/androiddev/comments/if5l7h/google_play_ratings_are_not_average_based_and/
---
It's beeen a while since I've been monitoring the Google Play rating on one of my apps (the rating that is displayed to users on Google Play).

&amp;#x200B;

My Google Play rating according to the Google Play Dev Console is **4.448** which rounds to **4.4** on Google Play.

&amp;#x200B;

As defined by Google, "the Google Play rating is calculated based on your most recent ratings".

&amp;#x200B;

Here are the average ratings of my app for all the selectable periods on Google Play dev console:

**Lifetime: 4.609**

**Last 365 days: 4.469**

**Last 180 days: 4.561**

**Last 30 days: 4.485**

**Last 7 days: 4.580**

**Today: 4.556**

&amp;#x200B;

As you can see, for Any selected period, the average of all the ratings on that period is higher than the Google Play rating.

&amp;#x200B;

It's also very strange that the app's Google Play rating (4.448) has not changed more than 0.001 up or down in the last couple of months even if it receives a decent amount of daily ratings (around 10 per day) and "the most recent ratings" definitely have varied.

&amp;#x200B;

The only logical explanation I can think of (other than that the whole rating system is bugged) is that Google is weighting the ratings so that more "helpful"/upvoted ratings have more weight on the average. But even if that was the case, how can the number be so stable?

&amp;#x200B;

I think the lack of transparency/clarity on how Google Play ratings are calculated is an issue and contributes to raising concerns on potential manipulation.
## [7][5 Awesome Open-source Android Kotlin Apps with MVVM, Dagger, Coroutines, Retrofit, and Room!](https://www.reddit.com/r/androiddev/comments/ifdikf/5_awesome_opensource_android_kotlin_apps_with/)
- url: https://medium.com/@androiddevnotes/5-awesome-open-source-android-kotlin-apps-with-mvvm-dagger-coroutines-retrofit-and-room-98b65296b199
---

## [8][Example based learning similar to Hacking with Swift?](https://www.reddit.com/r/androiddev/comments/iffq24/example_based_learning_similar_to_hacking_with/)
- url: https://www.reddit.com/r/androiddev/comments/iffq24/example_based_learning_similar_to_hacking_with/
---
I'm starting a new position as an Android developer and I was wondering if there were any similar resources to [Hacking with Swift](https://www.hackingwithswift.com/100/swiftui). They basically just handhold you through a bunch of projects, and you get a brief overview of what iOS is capable of. I actually got pretty far down the Android dev rabbit hole ~5 years ago, but that was back when it was still using Java and I'm sure many things have changed. Are there any resources similar to that (text preferred)? 

Thanks!
## [9][What are the different foregroundServiceTypes?](https://www.reddit.com/r/androiddev/comments/ifo5e7/what_are_the_different_foregroundservicetypes/)
- url: https://www.reddit.com/r/androiddev/comments/ifo5e7/what_are_the_different_foregroundservicetypes/
---
Most tutorials tells me to set it to "location", but I also see "dataSync", "connectedDevice" and others. What are the implications for applying these types to my service, and what are implications of NOT applying these types?

Will I fx not be able to make phone calls if I don't use the "phoneCall" service type?
## [10][Request Android permissions using annotations only](https://www.reddit.com/r/androiddev/comments/if0tka/request_android_permissions_using_annotations_only/)
- url: https://www.reddit.com/r/androiddev/comments/if0tka/request_android_permissions_using_annotations_only/
---
[https://github.com/LikeTheSalad/aaper](https://github.com/LikeTheSalad/aaper)
## [11][What costs come along a world chat for a game?](https://www.reddit.com/r/androiddev/comments/ifn83o/what_costs_come_along_a_world_chat_for_a_game/)
- url: https://www.reddit.com/r/androiddev/comments/ifn83o/what_costs_come_along_a_world_chat_for_a_game/
---
Talking about having a world chat players can write in, what continuous costs come along this? Does it cost e.g. server payment to have it running?
## [12][Instant apps](https://www.reddit.com/r/androiddev/comments/ifkj6j/instant_apps/)
- url: https://www.reddit.com/r/androiddev/comments/ifkj6j/instant_apps/
---
I need to make an instant app but I cannot find any good information. When instant app launches firstly, lots of people wrote articles about how to write instant app. But all of them are deprecated.  I asked a working example in stack overflow but no body replied. Why? Is instant apps dead? Why no body mention about it?
