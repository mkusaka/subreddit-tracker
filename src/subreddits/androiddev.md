# androiddev
## [1][App Feedback Thread - May 02, 2020](https://www.reddit.com/r/androiddev/comments/gc5ncv/app_feedback_thread_may_02_2020/)
- url: https://www.reddit.com/r/androiddev/comments/gc5ncv/app_feedback_thread_may_02_2020/
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
## [2][Weekly Questions Thread - April 27, 2020](https://www.reddit.com/r/androiddev/comments/g8yltx/weekly_questions_thread_april_27_2020/)
- url: https://www.reddit.com/r/androiddev/comments/g8yltx/weekly_questions_thread_april_27_2020/
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
## [3][A reminder that Single Activity App Architecture has been the official Google recommendation since 2 years ago (May 9, 2018)](https://www.reddit.com/r/androiddev/comments/gc0yw3/a_reminder_that_single_activity_app_architecture/)
- url: https://www.reddit.com/r/androiddev/comments/8i73ic/its_official_google_officially_recommends_single/
---

## [4][Play console just stopped counting downloads yesterday at 11 PM, ok I guess?](https://www.reddit.com/r/androiddev/comments/gc38hc/play_console_just_stopped_counting_downloads/)
- url: https://i.redd.it/9smwxmtgebw41.jpg
---

## [5][After 2¾ years, today is my last day at @Google](https://www.reddit.com/r/androiddev/comments/gbhtum/after_2¾_years_today_is_my_last_day_at_google/)
- url: https://twitter.com/JakeWharton/status/1256207195957575681
---

## [6][With the Android Development world moving so fast, how do you cope up with fear of getting left behind? Are you always catching up on new libs/frameworks, feeling burned out and wondering you don't know enough?](https://www.reddit.com/r/androiddev/comments/gc2p3s/with_the_android_development_world_moving_so_fast/)
- url: https://www.reddit.com/r/androiddev/comments/gc2p3s/with_the_android_development_world_moving_so_fast/
---

## [7][I want to publish an app that uses SEND_SMS, READ_PHONE_STATE, RECEIVE_SMS permissions for alarms kit automation](https://www.reddit.com/r/androiddev/comments/gc3zp7/i_want_to_publish_an_app_that_uses_send_sms_read/)
- url: https://www.reddit.com/r/androiddev/comments/gc3zp7/i_want_to_publish_an_app_that_uses_send_sms_read/
---
Hi

I have developed an android application that sends and receive SMS for control some alarms kit. I need READ\_PHONE\_STATE to be able to receive a call from the alarms kit.

On the play store exists some similar application that has the same functions ( [https://play.google.com/store/apps/details?id=com.sicurit.sicuritsms&amp;hl=en](https://play.google.com/store/apps/details?id=com.sicurit.sicuritsms&amp;hl=en) ), but when I try to publish mine Google always reject the update.

I have contacted via email the Google developer support, but I receive always the same standard reply: see our documentation, remove the permissions, see our policy, etc.

But my app for works correctly needs that permissions, I don't use this permission for stole or get personal information of the users, and I don't understand how the other app that sends SMS have resolved this issue!!

Initially, when I have published the first version of the apps I have filled a little questionnaire about why I want to use these permissions, but now, when I try to submit an update, this questionnaire is not shown.

Does someone have experience with those problems? Google support is terrible about this thematic!!
## [8][Google play console stats acting weird.](https://www.reddit.com/r/androiddev/comments/gc5t0h/google_play_console_stats_acting_weird/)
- url: https://www.reddit.com/r/androiddev/comments/gc5t0h/google_play_console_stats_acting_weird/
---
Today all installs are touched zero. I guess something is wrong with stats or counting. Anyone having same issue?
## [9][Should we use Navigation Component?](https://www.reddit.com/r/androiddev/comments/gc5ceh/should_we_use_navigation_component/)
- url: https://www.reddit.com/r/androiddev/comments/gc5ceh/should_we_use_navigation_component/
---
Hi, I just wanted to have your opinion on Navigation Component.

I read some Medium explaining how great it is and some other explaining how some companies that implemented it went back to traditional navigation with activities.

So I am curious to know your opinion about it, do you recommend it?

Thanks everyone!
## [10][How should I merge multiple Async UseCases into a single one?](https://www.reddit.com/r/androiddev/comments/gc2o9z/how_should_i_merge_multiple_async_usecases_into_a/)
- url: https://www.reddit.com/r/androiddev/comments/gc2o9z/how_should_i_merge_multiple_async_usecases_into_a/
---
Hello, I am refactoring parts of our code and feel unsure about my reasoning for a specific problem. Our business requirements are the following:

* We have to call 4 network requests separately from different views.
* We have to wait for 4 network requests **to finish** and **start a new network request**.

Previously all that was handled in a single class, which was shared among 4 different views. However, I am unsure if my reasoning was correct. I did it in the following way:

1. I would initialize **4 replay subjects** for each network request to track when network requests will be finished. (I know I could use RxRelay for this kind of logic, but it is not currently an issue to me)

&amp;#8203;

    fun init(){
        firstTrigger : ReplaySubject&lt;Boolean&gt; = ReplaySubject.create()
        secondTrigger : ReplaySubject&lt;Boolean&gt; = ReplaySubject.create()
        thirdTrigger :  ReplaySubject&lt;Boolean&gt; = ReplaySubject.create()
        forthTrigger : ReplaySubject&lt;Boolean&gt; = ReplaySubject.create()
    }

&amp;#x200B;

2. Here is an example code of a one from 4 network requests. I trigger the subject after success.

&amp;#x200B;

    fun networkRequestNrOne() { 
        firstObservable = apiService.networkRequestNrOne() 
    
        compositeDisposable.add(firstObservable
        .retryWhenError(5, 1)
        .subscribeOn(Schedulers.io()) 
        .observeOn(AndroidSchedulers.mainThread()) 
        .subscribe({ 
            success -&gt; first.onNext(true) }, 
            { throwable -&gt; handleError(throwable) })
        ) 
    }

&amp;#x200B;

3. Then I have a method, which is called in init of this class. It uses zip operator from RxJava to combine all network requests

    fun observeWhenToFire(): Observable&lt;Boolean&gt; {
            return Observable.zip(firstTrigger, secondTrigger, thirdTrigger,
                    forthTrigger,
                    Function4&lt;Boolean, Boolean, Boolean, Boolean, Boolean&gt; { first, second, third,
                    forth -&gt;
                        //handle checking if all are true for example.
                    })
        }

Now I would like to refactor it and create 4 different UseCases for such cases and a new one UseCase, which would include all 4 UseCases and return a merged response.

However, I would still need to use Rx Subjects to handle such a scenario. Every UseCase would take a subject tied to that UseCase **as a parameter**, while UseCase, **which zips all other** UseCases, would **include 4\*1 subjects**.

Is there any better way? I even thought about using regular callbacks and it might be even more readable code.  I do use coroutines as well, however this time I don't want to refactor this logic to coroutines, because of our c**ustom retryWhenError extension**, which is quite unique :)
## [11][Can I use python on my app?](https://www.reddit.com/r/androiddev/comments/gc165y/can_i_use_python_on_my_app/)
- url: https://www.reddit.com/r/androiddev/comments/gc165y/can_i_use_python_on_my_app/
---
Here's the thing, I considero myself a decent python coder and a beginer in kotlin, I have the app idea, I know how I want it to work but I want to use more what i already know to get an MVP faster.

Is There a way to write some python 3 modules with the functionality i need and compile it with the rest of the app?
## [12][SocialDistance - BLE Contact Tracing](https://www.reddit.com/r/androiddev/comments/gbzyjw/socialdistance_ble_contact_tracing/)
- url: https://www.reddit.com/r/androiddev/comments/gbzyjw/socialdistance_ble_contact_tracing/
---
Two weeks before google and apple announced that they were working on a contact tracing app using Bluetooth Low Energy, I was working on the same.

My app was denied to the play store the day after it was announced by them. I decided to make it open source.

[https://github.com/rufolangus/SocialDistance](https://github.com/rufolangus/SocialDistance)

&amp;#x200B;

I have not had time to refactor. It uses firebase as a database and google AdMob, I removed all the keys so it probably will not build. If it gets traction I will do proper refactoring.
