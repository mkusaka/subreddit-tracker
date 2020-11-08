# androiddev
## [1][App Feedback Thread - November 07, 2020](https://www.reddit.com/r/androiddev/comments/jpq1lo/app_feedback_thread_november_07_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jpq1lo/app_feedback_thread_november_07_2020/
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
## [2][Weekly Questions Thread - November 02, 2020](https://www.reddit.com/r/androiddev/comments/jmley8/weekly_questions_thread_november_02_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jmley8/weekly_questions_thread_november_02_2020/
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
## [3][Is it possible to host my own notification server and *NOT* use FCM?](https://www.reddit.com/r/androiddev/comments/jq9gkc/is_it_possible_to_host_my_own_notification_server/)
- url: https://www.reddit.com/r/androiddev/comments/jq9gkc/is_it_possible_to_host_my_own_notification_server/
---
Wondering if it is possible to create my own notification server and avoid having to use Firebase. Seems like the universe wants everyone to use Google, and folk happily obliges, but for sake of decentralization I would like to explore alternatives.

Are there any self-hosted FOSS notification servers for Android that will allow me to send every type of notification that Firebase can?

Or perhaps a spec so I can (try to) implement one myself.
## [4][does anybody know why do facebook messenger request the captive portal page for ?](https://www.reddit.com/r/androiddev/comments/jq1x3g/does_anybody_know_why_do_facebook_messenger/)
- url: https://i.redd.it/cvznkwz9vwx51.png
---

## [5][Learn the shortcomings of windows floating over other apps on Android. It took me 7 years to go through all of it.](https://www.reddit.com/r/androiddev/comments/jq9vgd/learn_the_shortcomings_of_windows_floating_over/)
- url: https://localazy.com/blog/floating-windows-on-android-9-shortcomings
---

## [6][Unable to attach file to Gmail app from my app in Android 11](https://www.reddit.com/r/androiddev/comments/jq78y5/unable_to_attach_file_to_gmail_app_from_my_app_in/)
- url: https://www.reddit.com/r/androiddev/comments/jq78y5/unable_to_attach_file_to_gmail_app_from_my_app_in/
---
Hello.

I have a logs file that is attached to an email client app when a user wants to contact our support team from within the app.From android 11 it stopped working, so I added a FileProvider which seemed to work.However, since I upgrade the compile Sdk version to 30, it stopped working again.I read that there were some storage changes that were introduced in Android 11, but I couldn't figure out what is it exactly that I need to do know if I want to attach a simple internal (cache or external cache) file (not in shared storage) like a logs file.

Activity :

    val logFile: File = File(globalContext.externalCacheDir, "MyLogFile.log")
    
    fun Activity.openMail(
        type: String = "text/plain"
    ): Boolean {
        val intent = Intent(Intent.ACTION_SENDTO)
        intent.putExtra(Intent.EXTRA_EMAIL, arrayOf("&lt;some address...&gt;"))
        intent.putExtra(Intent.EXTRA_SUBJECT, "Android Logs")
        intent.putExtra(Intent.EXTRA_TEXT, "some text")
    
    
        val uri = FileProvider.getUriForFile(this, "${BuildConfig.APPLICATION_ID}.fileprovider", logFile)
        intent.putExtra(Intent.EXTRA_STREAM, uri)
    
        intent.addFlags(Intent.FLAG_GRANT_READ_URI_PERMISSION)
        intent.addFlags(Intent.FLAG_GRANT_WRITE_URI_PERMISSION)
        intent.type = type
        intent.data = Uri.parse("mailto:")
    
        if (intent.resolveActivity(packageManager) != null) {
            startActivity(Intent.createChooser(intent, "Send email via:"))
            return true
        } else {
            return false
        }
    }

Manifest :

    &lt;uses-permission android:name="android.permission.WRITE_EXTERNAL_STORAGE" /&gt;
    &lt;uses-permission android:name="android.permission.READ_EXTERNAL_STORAGE" /&gt;
    
    &lt;provider
        android:name="androidx.core.content.FileProvider"
        android:authorities="${applicationId}.fileprovider"
        android:exported="false"
        android:grantUriPermissions="true"&gt;
            &lt;meta-data
                android:name="android.support.FILE_PROVIDER_PATHS"
                android:resource="@xml/provider_paths" /&gt;
    &lt;/provider&gt;

provider\_paths :

    &lt;?xml version="1.0" encoding="utf-8"?&gt;
    &lt;paths&gt;
        &lt;external-cache-path name="external_cache" path="."/&gt;
    &lt;/paths&gt;
## [7][1 — 7 November Android Newsletter](https://www.reddit.com/r/androiddev/comments/jqac9w/1_7_november_android_newsletter/)
- url: https://www.reddit.com/r/androiddev/comments/jqac9w/1_7_november_android_newsletter/
---
Stay up to date with Android development, in this week's edition:   
😍 View Binding Benefits  
🏗️ Build UI with Jetpack Compose  
🤐 Private library repository  
🤓 RecyclerView &amp; Sealed Classes  
and much more!

Read it here 👉 [https://vladsonkin.com/android-newsletter-19/](https://vladsonkin.com/android-newsletter-19/?fbclid=IwAR3CjVL55KJxUSF7xRdohBbPh8DTAh7g2XhK8wRaEtXGnvy0XtQMxVF-Tg0)  
What's your favorite one?

🔥Featuring [@NixonTherRipper](https://twitter.com/NixonTherRipper) [@ibrahimsn98](https://twitter.com/ibrahimsn98) [@SG5202](https://twitter.com/SG5202) [@theDroidLady](https://twitter.com/theDroidLady) [@RimGazzeh](https://twitter.com/RimGazzeh) [@FMuntenescu](https://twitter.com/FMuntenescu) and many other great authors!

💚 Subscribe and receive new editions directly to your email. Weekly, no spam, unsub anytime.  
Here is an example: [https://mailchi.mp/1790973b5a2e/android-newsletter-19](https://mailchi.mp/1790973b5a2e/android-newsletter-19?fbclid=IwAR0nvzgRPHySnIn0U4KV8ZoveCXQHCMS4-rpri-r4Nsqwu5-URMn4q_ONJ4)
## [8][An introduction to Constraintlayout 2.0](https://www.reddit.com/r/androiddev/comments/jqbpgv/an_introduction_to_constraintlayout_20/)
- url: https://link.medium.com/TvhnHpMnfbb
---

## [9][Android biometric authentication](https://www.reddit.com/r/androiddev/comments/jq8awi/android_biometric_authentication/)
- url: https://www.reddit.com/r/androiddev/comments/jq8awi/android_biometric_authentication/
---
Hi, I work on an e-payment app and i want to lock app with a finger print or a pin, 
I look at github and found a lib i edited it to work with my code but its only uses pin and i cannot add a biometric to it i tried but not succeed.
Does anyone know how to do it or know a working libraries have both ?


The library work by making a parent activity and all activities on my app must inherit from it, is this the best way
## [10][View layout shows up as expected in Design view, but emulator isn't using the full screen length (image inside)](https://www.reddit.com/r/androiddev/comments/jq553e/view_layout_shows_up_as_expected_in_design_view/)
- url: https://www.reddit.com/r/androiddev/comments/jq553e/view_layout_shows_up_as_expected_in_design_view/
---
This is a screen grab of the emulator (left) beside the design view.

https://imgur.com/7tiwZCY

The design view fits the background over the entire screen, as intended.  When I run the app, however, you can see the white rectangle at the bottom of the emulator screen where the background isn't being scaled.  Is this a known issue with the emulator, or perhaps a different issue you're familiar with?

I'd appreciate any insight you might have.  The only thing I can find via google is an issue where the entire view (vertical *and* horizontal) is scaled down.
## [11][Weird, is there any way to clear cache and storage once a user updates my app.](https://www.reddit.com/r/androiddev/comments/jq5ijy/weird_is_there_any_way_to_clear_cache_and_storage/)
- url: https://www.reddit.com/r/androiddev/comments/jq5ijy/weird_is_there_any_way_to_clear_cache_and_storage/
---
Hey all

I've made some major internal logic changes in v5 of my app that will most likely cause my app for a user who updates through the play store to crash.
They will need to clear the cache and storage from within android settings for my app.
Is there any way to have this done automatically when they update? Or whats the best way to go about doing this.


Has anyone faced this problem before?
Thanks
## [12][Built some Toolbar + CollapsingToolbar, playground for scroll flags, and material design elements focused templates. If you know how Window Insets consumed, dispatched and effects with or without or wish to contribute new templates it's more than welcome. Looking for detailed inset guide.](https://www.reddit.com/r/androiddev/comments/jpqzwv/built_some_toolbar_collapsingtoolbar_playground/)
- url: https://github.com/SmartToolFactory/Toolbar-Samples
---

