# androiddev
## [1][Mod Announcements: Updated Rules](https://www.reddit.com/r/androiddev/comments/gj7xr4/mod_announcements_updated_rules/)
- url: https://www.reddit.com/r/androiddev/comments/gj7xr4/mod_announcements_updated_rules/
---
Hello everyone! It's been a while since we had a moderation update, so we've recently started taking a closer look at some things we can do to improve things around here.

One thing that has come up often is that the subreddit is seen as a dead-end. This comes in a few different facets, but namely:

* High number of low-quality tech support posts creating a low signal to noise ratio
* Certain types of posts becoming a honeypot for anti-Google vents, witch hunts, or other strange conspiracy theories.
* A high volume of bad faith engagement in comments that discourages people from contributing/returning.

Historically, we’ve taken a pretty hands-off approach to the subreddit other than a few ground rules. To try to address some of these friction points though, we’re announcing some updates to the subreddit rules today.

Details are sectioned, changes highlighted in bold where appropriate in the quoted descriptions.

---

**Rule 1: Must be related to Android Development**

No real change, but clarify in the description that this doesn’t include custom ROM development (should head to XDA)

&gt; You may post about code, design, distribution, marketing, hiring, etc. as long as it does not conflict with any other rules. We only ask that it has something to do with Android development. **Please note that this is not the place for ROM development, and should be taken to more appropriate venues like https://xda-developers.com/**

---

**Rule 2: No “help me” posts**

Previously: “No easily searched/specific dev questions”

This is to simplify the rule, but also broaden it. Asking for technical help with your specific problem does not merit a dedicated thread, nor do quick questions.

&gt; Soliciting general discussion about architecture, performance optimizations, or design is fine. **Asking for technical help with your specific problem is not** and you must redirect them to StackOverflow or the Daily Questions Thread stickied to the subreddit. **This also includes “which/what/how should I learn/do” threads.**

---

**Rule 4: No app takedown/Play Store vent posts**

Previously: “No “app takedown posts without desc of app”

We realize this will be unpopular with a vocal minority, but we feel this is for the best (detailed in the rule wording below).

&gt; **We’re sorry your app was taken down or your account was suspended or you couldn’t get a bad review removed. This subreddit cannot help you. We tried allowing these in the past under certain conditions, but in practice these threads were unproductive, often omitting key details showing they were justified, and often just became vent threads.**

---

**Rule 8: No paywalled submissions**

No title change, but we’ve ha’ a number of posts that try to advertise paywalled medium posts and use “friend links” as a copout.

&gt; Any posts linking paywalled articles are not allowed. **This includes private URLs, such as “friend” links on paywalled Medium articles**

---

**Rule 10: Be respectful and engage in good faith**

This is new, and usually one that doesn’t need writing but needs to be written now. Reasonable people can disagree and you can articulate your thoughts without condescending to others. We’re not here to be the word police, but we’re also not going to be welcoming hosts to those that would regularly break this rule.

&gt; The Android developer community is a warm and friendly field and /r/androiddev strives to continue this. Engage in good-faith discussion and be respectful of others’ opinions, privacy, and intentions. Threads that violate this will be removed at mods’ discretion. This rule is intentionally broad, as toxic behavior comes in a variety of different forms. Examples: ad hominem, [sealioning](https://twitter.com/sbarolo/status/1036782685547622401), targeted attacks on others’ work, edgelording, and other keyboard warrior behavior.

---

We hope these changes help. We have more ideas in the pipeline too (particularly around how to implement these changes meaningfully, possibly with more moderation hands on deck) and will keep you all posted when we have more to share!
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
## [3][It is not just us! None can escape vague terms of Google policies.](https://www.reddit.com/r/androiddev/comments/gj4t26/it_is_not_just_us_none_can_escape_vague_terms_of/)
- url: https://blog.pushbullet.com/2020/05/13/lets-guess-what-google-requires-in-14-days-or-they-kill-our-extension
---

## [4][Create stunning Google Play screenshots - What features do you need for your app?](https://www.reddit.com/r/androiddev/comments/gjlcti/create_stunning_google_play_screenshots_what/)
- url: https://www.reddit.com/r/androiddev/comments/gjlcti/create_stunning_google_play_screenshots_what/
---
I [shared our initial tool](https://www.reddit.com/r/androiddev/comments/fp98il/create_awesome_app_screenshots_for_playstore/) over a month ago, and received a lot of positive feedback. Tbh, I've never had so much upvote karma, so decided to spend every evening of the next month re-writing &amp; improving the tool. Hence, I'm ready to share the new version, and ask you to **flood me with features requests**. 

Link: [https://previewed.app/screenshots/googleplay](https://previewed.app/screenshots/googleplay)

We decided to do a weekend Hackathon, where we implement all of the requests (Realistically possible). This could be:

* Adding different designs
* Devices
* Anything that comes to your mind!

\------------------ Question for Devs -----------------

We had a problem of adding new devices. Since people have different devices and a lot of Android phones with custom screenshot dimensions, there is no phone that fits all. So we decided to create a variable frame, which kinda resembles the phone. Do you think it's a good idea? Or do you prefer to have a set phone instead? E.g. Samsung galaxy s20 / Google Pixel etc.? 

\------------------ End of Question --------------------

**TL DR:**

Simple tool, to create app visuals for Google Play. In a few clicks, you can get something like this:

&amp;#x200B;

https://preview.redd.it/nczspbux3qy41.png?width=1776&amp;format=png&amp;auto=webp&amp;s=3f24888e590d435e8d5b18028951f2f485318ec4
## [5][ConstraintLayout 2.0.0 beta 6](https://www.reddit.com/r/androiddev/comments/gjbbls/constraintlayout_200_beta_6/)
- url: https://androidstudio.googleblog.com/2020/05/constraintlayout-200-beta-6.html
---

## [6][getting frames from usb camera](https://www.reddit.com/r/androiddev/comments/gjmcxv/getting_frames_from_usb_camera/)
- url: https://www.reddit.com/r/androiddev/comments/gjmcxv/getting_frames_from_usb_camera/
---
[https://developer.android.com/guide/topics/connectivity/usb/host](https://developer.android.com/guide/topics/connectivity/usb/host)

There are questions about this all over the place, but thought I'd try here. I'm trying to use a usb camera in my app. I can connect to it, see that it USB\_CLASS\_VIDEO interface. I just don't know how to ask it for frames. The Android tutorial really falls short with how to communicate with the device. It says to use UsbDeviceConnection.bulkTransfer, but I'm not getting anything out of it.

Is this possible without getting into libraries like libuvc?
## [7][A lightweight Android network response API for handling data and error response with transformation extensions.](https://www.reddit.com/r/androiddev/comments/gjlroh/a_lightweight_android_network_response_api_for/)
- url: https://github.com/skydoves/Sandwich
---

## [8][Room crash: "Pre-packaged database has an invalid schema"](https://www.reddit.com/r/androiddev/comments/gjlqpg/room_crash_prepackaged_database_has_an_invalid/)
- url: https://www.reddit.com/r/androiddev/comments/gjlqpg/room_crash_prepackaged_database_has_an_invalid/
---
Hi :)  
After upgrading from room 2.1.0 to 2.2.5 my app crashes generating this stacktrace:   


&gt; java.lang.IllegalStateException: Pre-packaged database has an invalid schema: MyTestTable(com.testcomponent.room.entities.MyTestTable).  
&gt;  
&gt;Expected:  
&gt;  
&gt;TableInfo{name='MyTestTable', columns={Params=Column{name='Params', type='TEXT', affinity='2', notNull=true, primaryKeyPosition=2, defaultValue='null'}, Rank=Column{name='Rank', type='INTEGER', affinity='3', notNull=false, primaryKeyPosition=0, defaultValue='null'}, Id=Column{name='Id', type='INTEGER', affinity='3', notNull=true, primaryKeyPosition=1, defaultValue='null'}, Date=Column{name='Date', type='TEXT', affinity='2', notNull=false, primaryKeyPosition=0, defaultValue='null'}}, foreignKeys=\[\], indices=\[\]}  
&gt;  
&gt;Found:  
&gt;  
&gt;TableInfo{name='MyTestTable', columns={Params=Column{name='Params', type='NVARCHAR', affinity='2', notNull=true, primaryKeyPosition=2, defaultValue='null'}, Rank=Column{name='Rank', type='INT', affinity='3', notNull=false, primaryKeyPosition=0, defaultValue='null'}, Id=Column{name='Id', type='BIGINT', affinity='3', notNull=true, primaryKeyPosition=1, defaultValue='null'}, Date=Column{name='Date', type='DATETIME', affinity='1', notNull=false, primaryKeyPosition=0, defaultValue='null'}}, foreignKeys=\[\], indices=\[\]}

  
I'm wondering: Can room handles a prepackaged db containing a DATETIME column ?  
This is the SQL create statement:

    CREATE TABLE "MyTestTable" (
    	"Id"	[BIGINT] NOT NULL,
    	"Params"	[NVARCHAR](250) NOT NULL,
    	"Rank"	[INT],
    	"Date"	[DATETIME],
    	PRIMARY KEY("Id","Params")
    )

And i'm loaded an example project on this [github repo](https://github.com/aeroxr1/room-sample-project.git) on master branch
## [9][How to create resizing handles on ImageView](https://www.reddit.com/r/androiddev/comments/gjlmjg/how_to_create_resizing_handles_on_imageview/)
- url: https://www.reddit.com/r/androiddev/comments/gjlmjg/how_to_create_resizing_handles_on_imageview/
---
https://i.imgur.com/Rm9cgQ0.png

How would I go on creating these resizing squares? Is there a library that Android has?
## [10][Can someone please explain to me the "right way" to use MutableLiveData and coroutines outside of Main scope to do internet API requests?](https://www.reddit.com/r/androiddev/comments/gjla8q/can_someone_please_explain_to_me_the_right_way_to/)
- url: https://www.reddit.com/r/androiddev/comments/gjla8q/can_someone_please_explain_to_me_the_right_way_to/
---
I'm trying to learn how all the pieces fit together to request a list of `Foo` in JSON from a public API (with Moshi) and get it into a `MutableLiveData&lt;List&gt;`. Here's the basic outline of what I've got so far, and a list of questions follow below the code:

    private var job = Job()
    private val fooScope = CoroutineScope(job + Dispatchers.IO)
    private val fooSerializedList = MutableLiveData&lt;List&lt;FooSerialized&gt;&gt;()
    
    // Caller must handle any exceptions this function throws (eg network errors)
    private fun getFooSerialized(): List&lt;FooSerialized&gt; {
        val getFooDeferred = FooApi.retrofitService.getFoo()
        var listResult: List&lt;FooSerialized&gt;
        fooScope.async {
            // await the completion of our Retrofit request
            listResult = getFooDeferred.await()
            if(listResult.isEmpty()) Log.wtf(TAG, "Empty listResult")
        }
        return listResult
    }

And the caller call this function inside a try/catch block to handle network errors:

    fooScope.async {
        try {
            fooSerializedList.postValue(getFooSerialized())
            // should I do stuff with fooSerializedList here or outside coroutine?
        } catch (e: Exception) {
            // handle error
        }
    }
    // Should I do stuff with fooSerializedList here and will it be populated?

**Questions:**

1. `getFooSerialized` doesn't work because `listResult` isn't initialized, what's The Right Way to handle this?
2. Coroutines in IO scope: I keep having to use `postValue` for LiveData, rather than `setValue` (i.e. the `.value` property) because it doesn't work when not in Main scope. In my second code block, where should I access `fooSerializedList`, and how do I access it in a way such that it will be correctly populated by the time I access it?
3. Any other advice on how I could be doing this better + example code?
## [11][Android Dialogs: Android Makers 2020](https://www.reddit.com/r/androiddev/comments/gjaufy/android_dialogs_android_makers_2020/)
- url: https://youtu.be/hl_Br5wiLhY
---

## [12][Recap: ASO CONFERENCE 2020](https://www.reddit.com/r/androiddev/comments/gjk0wc/recap_aso_conference_2020/)
- url: https://www.reddit.com/r/androiddev/comments/gjk0wc/recap_aso_conference_2020/
---
Hi guys! 👋

Yesterday the ASO conference 2020 was held, a conference by ASO practitioners for ASO practitioners. If you missed this event, we have collected key points from the speeches for you.

Read all the details here 👇  
[https://asodesk.com/blog/recap-aso-conference-2020/](https://asodesk.com/blog/recap-aso-conference-2020/)
