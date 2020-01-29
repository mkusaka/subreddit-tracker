# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/eunf8c/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/eunf8c/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - January 27, 2020](https://www.reddit.com/r/androiddev/comments/eumgji/weekly_questions_thread_january_27_2020/)
- url: https://www.reddit.com/r/androiddev/comments/eumgji/weekly_questions_thread_january_27_2020/
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
## [3][FragmentContainerView is now the recommended View to use in order to host Fragments instead of the previously common way of using FrameLayouts.](https://www.reddit.com/r/androiddev/comments/evkj7s/fragmentcontainerview_is_now_the_recommended_view/)
- url: https://twitter.com/iamyogeshtelang/status/1222414493919002624?s=20
---

## [4][Google Suspended our App and Threatened to Delete all our Google Accounts Over Nothing](https://www.reddit.com/r/androiddev/comments/evc397/google_suspended_our_app_and_threatened_to_delete/)
- url: https://medium.com/@thedrbits/google-has-threatened-to-delete-all-our-google-accounts-over-nothing-13a05a31a55a
---

## [5][Developer support [bot] is just getting *even* worse](https://www.reddit.com/r/androiddev/comments/ev8jns/developer_support_bot_is_just_getting_even_worse/)
- url: https://www.reddit.com/r/androiddev/comments/ev8jns/developer_support_bot_is_just_getting_even_worse/
---
Brace yourselves... I thought it was impossible but now even the bots are failing to fill the email about policies violation.

May the gods help the android developers... \\o/

[\&lt; INSERT EXAMPLE TEXT \&gt;](https://preview.redd.it/p10hzi871kd41.png?width=1222&amp;format=png&amp;auto=webp&amp;s=970d3639c735f18cf1b299b133a1c5d6c8b6f4ab)
## [6][Coroutines and Java Synchronization Don't Mix](https://www.reddit.com/r/androiddev/comments/ev9lio/coroutines_and_java_synchronization_dont_mix/)
- url: https://blog.danlew.net/2020/01/28/coroutines-and-java-synchronization-dont-mix/
---

## [7][Recyclerview multiple viewtypes with Room, Livedata, Paging](https://www.reddit.com/r/androiddev/comments/evkvnc/recyclerview_multiple_viewtypes_with_room/)
- url: https://www.reddit.com/r/androiddev/comments/evkvnc/recyclerview_multiple_viewtypes_with_room/
---
Anyone using Room along with Livedata, Paging (Jetpack)?
Have a query.
Use case: Think like Facebook news feed. It has story bar (horizontal recyclerview) at top. Then posts. In between posts, sometimes they display the story bar again!
My app will have somewhat similar scenario.
How to implement it using recyclerview, room, paging?
Previously, while using dummy data, I created recyclerview with recyclerview.adapter, the adapter will handle a generic feed item type class. Then both posts and storybar will extend feeditem. While passing the list of posts to the recyclerview, I'd add those storybars in appropriate position. In the recyclerview adapter, onbindviewholder, I'll check their types, and populate them accordingly. It was working fine as long as I didn't try to implement database (room).
When using room and paging, recyclerview adapter must be PageListAdapter. How to inject those storybars in the adapter items?
Please note that if possible, I don't want to store storybar as say dummy post (with all actual post fields empty) in the database. If this is the only solution you think, please mention that also.
Don't know if I clarified properly. If anyone is interested, and wants further clarification, please say so.
## [8][How to Integrate Facebook Login Feature in Android Application?](https://www.reddit.com/r/androiddev/comments/evmcnf/how_to_integrate_facebook_login_feature_in/)
- url: https://www.metizsoft.com/blog/facebook-login-integration-in-android
---

## [9][Android Studio project behaving strangely after recovering from apk.](https://www.reddit.com/r/androiddev/comments/evk70e/android_studio_project_behaving_strangely_after/)
- url: https://www.reddit.com/r/androiddev/comments/evk70e/android_studio_project_behaving_strangely_after/
---
Hello, 
Android Studio messed up my project and overwrote 90% of my files with useless code. All layout files except for content-main and all activities except for the MainActivity are affected, even the grade settings. During this recovery maybe something got messed up in the textViews are sometimes displayed too small, text inside of the buttons is stuck to the top left corner, all of this cannot be changed, the setting for that is just not available. All constraint layouts are now transparent, they used to be white. I hope someone has seen this behavior before. Just to make it easier for you and me I have already posted this on [Stackoverflow](https://stackoverflow.com/questions/59957277/project-behaving-strangely-after-recovering-it-from-apk-and-swapping-out-the-con) with source code and screenshots.

Thanks in advance!
## [10][With a shifting Android roadmap on Storage, is it better for independent Android developers to sit out the Storage Apocalypse (and is iOS a better long term platform for independent developers) ?](https://www.reddit.com/r/androiddev/comments/evle74/with_a_shifting_android_roadmap_on_storage_is_it/)
- url: https://www.reddit.com/r/androiddev/comments/evle74/with_a_shifting_android_roadmap_on_storage_is_it/
---
**Android is now the dominant mobile platform**

The need to develop Android apps will remain for companies that want app versions for both iOS/Android - thus company and contract work will always remain.

And Android is a robust platform which is now more dominant than iOS across the world.

HOWEVER, for independent developers the last few years have been bumpy.  The lack of a stable roadmap, combined with a reversal of Android's long standing "old apps will run on new Android versions" mantra means that independent developers now face a greater burden of support - having to look back at apps that were mature, having to reengineer workarounds, and handle support calls.  All without extra revenue.  And being distracted from working on newer apps and new revenue streams.

The backtracking in their code to fit new rules that apply retroactively, elimination of entire niches of apps (which may render them suddenly economically unsustainable to continue with Android), and the increased support costs as earlier users who were happy, now complain that now their app has failed them.

The last year or two has been especially distracting - it has distracted certain classes of developers (those niches that got affected) so they could not focus as much on new development, but wound up working to bring old users into line with new Android policies and restrictions - this was all work which did not bring added revenue (more work, without return, or just to stay in place).

&amp;nbsp;

**Android as a bait-and-switch "open" platform**

Whether this is a result of a colossal bait-and-switch is open to debate - whether the "open platform" and all that was just to bring users and developers into Android vs. iOS during Android's early years, with intent to switch to an iOS-like format eventually.

However, the practical impact of this is that Android has a roadmap that is not reassuring for independent developers.

Google's policy changes have affected whole classes of apps (Call/SMS affecting call recorder apps, offline SMS backup apps, U.N. polio monitoring apps).

Google's changes for Storage promise a disruption of bigger scale than before - for example, compared to their Call/SMS fiasco.

&amp;nbsp;

**Android is not a separate entity**

What is worse, Android is not a separate company, but serves as the **subservient unit to Google's wider interests in search and advertising**.  Thus Android's direction as a mobile platform of global import is compromised to the unrelated needs of Google's other needs.  This may be one reason why Android seems to shoot itself in the foot sometimes - harming developers with their changes, and users one year later as those changes appear without challenge on user devices.

**Users and developers are merely participants without a say in it's eventual direction.**  This is untenable for a mobile platform that increasingly has strategic value to the world.

&amp;nbsp;

**Android is not iOS**

Given the vast diversity of devices and android versions which don't update in lockstep (like they do for iOS), has always made Android a more difficult platform to support for independent android developers.

However, that was with the Android mantra of "old apps always work on newer Android versions".

Now that Google has dispensed with this mantra in the last few years - Call/SMS changes being one example where earlier apps stopped working (call recorder apps could not work as before), and Storage being the next big one (and will be of greater impact) - it raises new questions about the viability of Android as a platform for independent developers.

Independent developers for the last few years have had to deal with:

- app niches which disappear at a whim (which leads to revenue shortfalls the disrupt independent developers' viability).

- bot-driven app bans and unresponsive support that is not answerable to anyone.  And the outrageous concept of some abstract threshold for irreversible lifetime career ban  on a developer create a climate of master over slave.  While early Android needed developers to drive their store, now Android seems to relish in the ways it can extend it's cruelty to developers.  The more arbitrary the punitive punishments the more "power" a rule seems to enjoy.  Meanwhile in the real world, the "3 strikes and you are out" rules are recognized as oppressive policy - that hasn't filtered into the think tanks within Google.

- reversal of age old mantra "old apps will work on newer Android versions" - leading developers to now be responsible for previous work, reversals, and it's attendant support costs.  This means independent **developers now have less time to devote to build new apps**, but are now expected to keep changing old apps (without extra revenue expectations), and have less time for new apps and the attendant revenue from new efforts.  Google management is exercising a hostile attitude towards independent developers that reflects a class-conscious mentality - where Google has stopped thinking how it feels to be an independent developer.

- Google's power allows them to change such basic things as Storage and it doesn't get advertised to users (most users still are unaware what Android 10 has in store).  **Google's power means they do not fear harming developers or even violating regulatory practices in the short term**.  Everything is curable by paying some regulatory penalty.  **Penalties which would be prohibitive for Android if it was a standalone organization become easy to pay for Google the search/advertising behemoth**.  This is a powerful reason for the separation of Android from Google - until that happens, the world's mobile ecosystem will not develop in a rational way - that benefits the mobile ecosystem (and not some other entity primarily).

- Developers have no voice in Android roadmap or about how inconsiderate these changes may be for users and developers (also related to above - if Android is not standalone it will never be responsive to mobile ecosystem concerns).

- Users never find out about changes in the pipeline until they are a fait accompli - mainstream media only covers the positives of Android 10 for example, and repeat it continuously for effect.  Again, if it was Android as standalone company, it would not exercise the same power.  But Google the behemoth can.  Negative stores on call recording vanishing, or Storage going away are occasional and easily forgotten.  The next Android version gets delivered and becomes a fait accompli.

- Google does not answer for user concerns - all **ire from bad Android decisions falls on developers**.  Google is an entity that cannot be questioned about Android, and does not have to conform to any expectations of users.  Users buy what appears on the market every year, and they will complain to developers (who did not cause those changes).  This is a system designed to to not be fixed - the circle of responsibility does not come back to Google.

- Android always has had a greater support burden, because of device diversity, and Google unwillingness to enforce consistency where it mattered for the ecosystem.  Instead their primary concern is that manufacturers include their Google search, and Google Play Store and other such services.  As Commonsware reports, one of the replacements for file storage, the Storage Access Framework (SAF) is inconsistently enforced across manufacturers (rendering that "alternative" less effective).  These types of random variations multiply the support burden for independent developers (who have to do both development and support).

&amp;nbsp;

**iOS support burden vs. Android - for indep devs**

It could be argued that the same thing happens on iOS as the OS evolves.

However, the less discussed problem with Android is it's API stability and consistency.  What is advertised for Android, may not be what is delivered a few months later.  Worse, Oreo 8.0 may deliver the first promise, and Oreo 8.1 may deliver a changed promise (as happened with the new audio engine for Oreo) - and BOTH versions remain in the market for months (unlike iOS which has intrinsic capability of updating devices in lockstep - a far greater percentage of devices run the latest iOS than the latest Android OS).

Take the Storage changes - it is quite possible that Google reverses some of the changes if problems emerge.

&amp;nbsp;

**Constrained iOS may be stabler than a slowly-constraining Android**

Now compare to iOS - where Storage has always been more closed.  While viewed negatively by most Android developers as a constrained system, it could be now argued that THAT is a far better system, because that constrained system appears as a standard.  Reducing the developer burden.

Meanwhile on Android, things are still in flux.  It is not clear where Storage policies will end up in the near future, of when finally delivered.

With iOS you at least know how things are to be done, and there are no hodge-podge alternatives that half-work or may not for Storage for Android:

- where there is a pre-built expectation among users that their files are persistent - developers will have to deal with how to placate these users (Google won't deal with these users directly) when users encounter different behavior.  Some of our users reported their call recordings were now all silence after upgrading to Android 10, or for storage they will find that files are in a sandbox, and not really where they used to be.

- Pre-existing code will no longer work as before - without performance penalties (with SAF) and sometimes requiring redesigning (to suit SAFs workflows, or limitations that you can't use fopen() in your C code) - **all this extra work will not earn independent developers any more revenue**.  And users will not value this as something to pay extra for - since all that work still produces a less efficient than before app!

- there will be no ONE way to save files - but multiple intersecting ways (and with it the user expectation of how file access works or should work) - save to app-specific storage, some MediaStore stuff, or SAF - all adding to extra confusion during support calls.

- there is a cloud hanging over SAF - whether file manager apps too will need to fill a Permissions Declaration Form - and what about apps which are not "primarily" file manager apps (similar situation as Call/SMS fiasco - where apps had to be either dialer handlers, or SMS handlers)

- it is unclear how the Storage story will settle - and if Google will reverse some of the changes if problems emerge.  In such a climate is it better to continue avoiding the Storage changes as long as possible, or better to restructure your apps on a still nascent roadmap ?  Unevenness in implementation is also an issue - how well will the promised Google solutions find currency in the wild ?  Commonsware is reporting that Storage Access Framework (SAF) is unevenly implemented across manufacturers.  Google never bothered to make conformance to one behavior part of Android certification.

&amp;nbsp;


**iOS and Android - which is now more difficult for independent developers ?**

Under these conditions, I wonder if the already-constrained world of iOS (in terms of Storage) is a vastly more stable development environment ?

While iOS may change from version to version, they at least have fewer of the device diversity, and manufacturer variation-in-implementation issues (since iOS devices are updated to newer iOS versions in near lockstep).;

When Android makes radical changes, developers should expect to face a few years of fallout as all the "diversity of devices" is multiplied by the "diversity in API behavior" - and appears as support issues.

Would it be wise for independent developers to sit out the Storage Apocalypse until it's roadmap is clear ?  Or to devote more effort on their iOS projects ?

As Android looks set to be a support s**tshow for the next year or two - due to Storage.

&amp;nbsp;


----

**References:**

- [Android apps will start to lose ability to access local persistent storage and ext SD card - in a move which will boost Google's cloud storage for app data backup](https://www.reddit.com/r/androiddev/comments/e5fb9m/android_apps_will_start_to_lose_ability_to_access/)

- [The Death of External Storage: The End of the Saga(?)](https://commonsware.com/blog/2019/06/07/death-external-storage-end-saga.html)

- [Hey, Google. Where is your roadmap ? Why commercial viability for indie devs is going down, and Google Play is dead for indie developers](https://www.reddit.com/r/androiddev/comments/b24i3d/hey_google_where_is_your_roadmap_why_commercial/)

- [What's up with Android APIs - why are they getting more convoluted ?](https://www.reddit.com/r/androiddev/comments/8rpu45/whats_up_with_android_apis_why_are_they_getting/)

- [Advice to a young Android developer - or things the tutorials don't tell you
](https://www.reddit.com/r/androiddev/comments/d9y376/advice_to_a_young_android_developer_or_things_the/)
## [11][Custom video player application](https://www.reddit.com/r/androiddev/comments/evgv87/custom_video_player_application/)
- url: https://www.reddit.com/r/androiddev/comments/evgv87/custom_video_player_application/
---
Hi I'm trying to make a custom video player application. I have an MP4 file with 2 videos and 1 audio.
I have 2 extractors and 2 decoders. They extract the two videos and renders to surface using surface view.
I can change which video to render to surface using

&gt;decoder.configure(format,null,null,0);
to
\n&gt;

&gt;decoder.configure(format,surface,null,0);

How to make this change on the fly while running the application using click of a button or something or is it possible?
## [12][Is ASO crucial for the success of your app/game?](https://www.reddit.com/r/androiddev/comments/ev9u8h/is_aso_crucial_for_the_success_of_your_appgame/)
- url: https://www.reddit.com/r/androiddev/comments/ev9u8h/is_aso_crucial_for_the_success_of_your_appgame/
---
From chatting to a few android developers, I learned they don't know much about ASO. That sounds crazy when you hear that over **60% of all downloads are generated by ASO.**

Thus I decided to write a whole series about ASO this is just a first part, my goals with this series are:

1. Help you get more downloads and grow your app/game.
2. And to simplify learning of ASO.

**Why is ASO important?**

* For the majority of apps, ASO is the number one driver of all downloads.
* As of today, there are over 2.7 million apps in the Google Play app store, If potential users/customers can't find your app it doesn't matter how good your app really is. ASO helps you with that.
* It is irrelevant how you get downloads, as long as you use Google Play, ASO will impact your growth.

&amp;#x200B;

**What is ASO - App Store Optimization?**                                                                                                              ASO is the process that has two main goals:

1. To get you more downloads by improving the conversion of traffic.
2. To get more traffic to your app listing page.

&amp;#x200B;

**1. Conversion of Traffic**

CT answers the question: **Why would somebody visit your app listing page and download app/game?**

[This is part of CT](https://preview.redd.it/yd4n7p827kd41.jpg?width=476&amp;format=pjpg&amp;auto=webp&amp;s=3041089fe1c38fb59475ddd2afb1afa6bc5f1e81)

* CT helps you stand out.
* CT invites people to visit your app listing page.

&amp;#x200B;

[Here you can read more about CT.](https://www.reddit.com/r/GrowingAppsAndGames/comments/ev8ht0/ct_what_makes_people_click_and_download_your/)

&amp;#x200B;

**CRO**

[App Listing Page](https://preview.redd.it/8881zqhd7kd41.jpg?width=720&amp;format=pjpg&amp;auto=webp&amp;s=87633171d14dfbc29ab6c051dad31c0b7d5e23f8)

The process of improving the app listing page in order to get more downloads is called CRO.

&amp;#x200B;

**2. How To Get More Traffic**

There are three main ways of getting traffic inside of Google Play.

**2.1. Being Featured**

Google Play team selects apps they think are a good representation of quality or trend.

* Editors' Choice

**Charts**

* Top Free
* Top Paid
* Top Grossing
* Trending

&amp;#x200B;

**2.2. Search And Ranking**

Google Play ranks keywords from your app based on factors like app ratings, android vitals (crash rate, battery usage, the average time a user spends using your app) and many other factors.

&amp;#x200B;

How higher ranking impacts downloads.

Ranking for an app is the same as the location for a shop.

https://preview.redd.it/a3kqupumdkd41.png?width=3085&amp;format=png&amp;auto=webp&amp;s=974dfea206f293b47f633cda4ea0279dbd5f2ab5

Of course, this is an oversimplified version. But you should get a point.

&amp;#x200B;

More downloads = higher ranking in search. And a higher ranking in search means even more downloads.

This creates a positive loop. Therefore even a small improvement in CT and CRO of 10% can be huge for your app/game. **Success breeds success.**

&amp;#x200B;

**The better location gives you credibility. The higher ranking does too.**

&amp;#x200B;

**Let's look at google search data:**

* The first page gets 95% of all traffics.
* The very first organic result on the first search engine results page sees about 32.5% of overall search traffic in terms of clicks. The 2nd domain sees 17.6%, while the 7th only sees 3.5%

Google search is similar to Google Play, the top 8 apps will get a large portion of traffic and the first app gets the highest % of traffic.

In addition, being first in search means your app is seen as the "best" and the majority of people use your app.

&amp;#x200B;

That's why the first search result is the default choice for most people.

**A great way to change that is to improve** [your CT.](https://www.reddit.com/r/GrowingAppsAndGames/comments/ev8ht0/ct_what_makes_people_click_and_download_your/)

&amp;#x200B;

**2.3. Search Ads**

They are great ASO/Search tool.

[Search Ad](https://preview.redd.it/r43u22gi7kd41.png?width=720&amp;format=png&amp;auto=webp&amp;s=e21891133456e65cecc4fe325bf914c8463ec434)

&amp;#x200B;

[Mind map of what you learned today + two types of searches.](https://preview.redd.it/gf591zkk7kd41.jpg?width=1609&amp;format=pjpg&amp;auto=webp&amp;s=b5a1723518adb841be0f778f76c2a96ab6608b00)

&amp;#x200B;

Thank you for reading. In this post, you learned some of the ASO fundamentals.

**I am writing a series about growing and marketing android apps/games here on Reddit.**

[You can check it here and read more about CT](https://www.reddit.com/r/GrowingAppsAndGames/)
