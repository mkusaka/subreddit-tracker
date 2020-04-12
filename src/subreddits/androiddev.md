# androiddev
## [1][App Feedback Thread - April 11, 2020](https://www.reddit.com/r/androiddev/comments/fz46be/app_feedback_thread_april_11_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fz46be/app_feedback_thread_april_11_2020/
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
## [3][Featured videos on app listing page now plays within the PS app instead of opening the YT app](https://www.reddit.com/r/androiddev/comments/fzu07d/featured_videos_on_app_listing_page_now_plays/)
- url: https://i.redd.it/obfi2oqzzcs41.png
---

## [4][Adventures in Compose - The Doom fire effect](https://www.reddit.com/r/androiddev/comments/fzva59/adventures_in_compose_the_doom_fire_effect/)
- url: https://adambennett.dev/2020/04/adventures-in-compose-the-doom-fire-effect/
---

## [5][These coding games could definitely improve your programming skills](https://www.reddit.com/r/androiddev/comments/fzk3vk/these_coding_games_could_definitely_improve_your/)
- url: https://hurmaniqbal.blogspot.com/2020/04/list-of-coding-games-to-practice.html
---

## [6][CS student looking for a free android course](https://www.reddit.com/r/androiddev/comments/fzs2s7/cs_student_looking_for_a_free_android_course/)
- url: https://www.reddit.com/r/androiddev/comments/fzs2s7/cs_student_looking_for_a_free_android_course/
---
 Hello,  I'm a CS student and I'm looking for a free course that teaches the  Android Development using Android Studio. I know how to code (especially  in Java) but I don't know how to build a decent/good android app.

I don't even know where to search, I tried udemy and youtube and I haven't found something that convinced me....

I  found a bunch of 5 year old courses, should they still be good? Also,  is it a standard to use the Kotlin language (when creating the project,  as it is the default) ? Or is Java still good?
## [7][ViewBinding thing you probably don't know about](https://www.reddit.com/r/androiddev/comments/fzw6e5/viewbinding_thing_you_probably_dont_know_about/)
- url: https://www.reddit.com/r/androiddev/comments/fzw6e5/viewbinding_thing_you_probably_dont_know_about/
---
If you need to, you can add custom layout for you main layout while using ViewBinding.

Initialize you custom layout just like you initializing your main fragment/activity layout:

// MainLayout ViewBinding variables

private var mBinding
: FragmentRegistrationFirstBinding? = null  

private val mainBinding get() = mBinding!!  

// CustomLayout ViewBinding variables  

private var mapsBinding
:RegistrationGoogleMapsLayoutBinding? = null  

private val addressBinding get() = mapsBinding!!

&amp;#x200B;

// Main Layout ViewBinding initialization  
mBinding = FragmentRegistrationFirstBinding.inflate(inflater, container, false)

// CustomLayout ViewBinding initialization  
mapsBinding = RegistrationGoogleMapsLayoutBinding.inflate(inflater, container, false)  


And then just add it to MainLayout ViewBinding, just like you do with a simple view, or remove it:

// For adding  

mainBinding.mainLayout.addView(addressBinding.*root*)  


//For removing

mainBinding.mainLayout.removeView(addressBinding.*root*)

&amp;#x200B;

where mainLayout is one of layouts in your MainLayout ViewBinding
## [8][How to get a working ViewModel in an Activity like the dev guide example?](https://www.reddit.com/r/androiddev/comments/fzt2wo/how_to_get_a_working_viewmodel_in_an_activity/)
- url: https://www.reddit.com/r/androiddev/comments/fzt2wo/how_to_get_a_working_viewmodel_in_an_activity/
---
I'm struggling through the [ViewModel tutorial](https://developer.android.com/topic/libraries/architecture/viewmodel#implement). Google's example code says to do this:

    class MyActivity : AppCompatActivity() {
    
    override fun onCreate(savedInstanceState: Bundle?) {
        // Create a ViewModel the first time the system calls an activity's onCreate() method.
        // Re-created activities receive the same MyViewModel instance created by the first activity.
    
        // Use the 'by viewModels()' Kotlin property delegate
        val model: MyViewModel by viewModels()
        model.getUsers().observe(this, Observer&lt;List&lt;User&gt;&gt;{ users -&gt;
            // update UI
        })
    } 
    }

However, this causes the following compiler error that I don't understand:

&gt;Unresolved reference. None of the following candidates is applicable because of receiver type mismatch:  
&gt;  
&gt;MainThread public inline fun &lt;reified VM : ViewModel&gt; Fragment.viewModels(noinline ownerProducer: () -&gt; ViewModelStoreOwner = ..., noinline factoryProducer: (() -&gt; ViewModelProvider.Factory)? = ...): Lazy&lt;MyViewModel&gt; defined in [androidx.fragment.app](https://androidx.fragment.app)

So, the Google example is different to [the docs](https://developer.android.com/reference/kotlin/androidx/fragment/app/package-summary#viewmodels), where `by viewModels()` delegate is scoped to a Fragment instead of an activity. I don't know if this is significant, because ... well, the documentation is pretty far removed from the reality I encounter when trying to reproduce something that remotely works like their example.

Is their example even valid and would it actually work?
## [9][How do you build a professional UI?](https://www.reddit.com/r/androiddev/comments/fz2gua/how_do_you_build_a_professional_ui/)
- url: https://www.reddit.com/r/androiddev/comments/fz2gua/how_do_you_build_a_professional_ui/
---
I started teaching myself to code about 2-3 years ago and started with Android studio. I could almost get any answer I needed from tutorials or resources.

Since I started I've begun countless projects but only finished a handful. And I've started to realise that it's because I can never make an app look as good as I like. 

Since the quarantine I've started a project I'm really invested in, but I still can't get the UI to look the way I want.
With coding, it's like I know what I want to do and I just need to find out how.
But with UI it's like I don't even know what I want to do.

Does anyone know any good tutorials or learning resources on how to build polished UIs?

Edit: been a busy day but come back to lots of answers thanks all! Will be reading through tonight, really appreciate it !
## [10][Is it okay to ask for donations in a medium outside the Play Store](https://www.reddit.com/r/androiddev/comments/fzv8cd/is_it_okay_to_ask_for_donations_in_a_medium/)
- url: https://www.reddit.com/r/androiddev/comments/fzv8cd/is_it_okay_to_ask_for_donations_in_a_medium/
---
I've got apps I don't plan on ever charging money for. However, I'd like to offer fans of these apps the opportunity to give me money if they wish to show appreciation in that way.

My question is: Would Google consider that (for instance, in the form of a "support me" link in an "About" fragment) an attempt at circumventing the App Store Payment system, or are they okay with kinda-sorta-noncommercial content like this?
## [11][Dualshock2 dpad problem](https://www.reddit.com/r/androiddev/comments/fzv3qj/dualshock2_dpad_problem/)
- url: https://www.reddit.com/r/androiddev/comments/fzv3qj/dualshock2_dpad_problem/
---
So I'm using an original dualshock2 with a usb converter and a otg adapter on my android phone and I can't get the dpad to work I've tried countless remapping/gamepad controller apps but it seems like my phone doesn't recive any inputs from dpad presses but they work I use em on the PS2 console do you guys have any idea how to fix this
(Sorry if this is against the rules I'm new here)
## [12][I would love some assistance](https://www.reddit.com/r/androiddev/comments/fzun28/i_would_love_some_assistance/)
- url: https://www.reddit.com/r/androiddev/comments/fzun28/i_would_love_some_assistance/
---
Noob HS student here trying to go about with my first app, so please go easy on me.

I finally figured out how to start Alpha testing so I grabbed a bunch of my friends mails and started it. It says 'full rollout'  in the App releases (closed tracks)  but the status of the application is 'Pending Publication'. I thought Pending Publication is only for if you are putting your app on google Play, which I am not. Also, my friends haven't really received any mail so far (Ik this as I put my normal mail account also a tester).

Have I done something wrong somewhere? Did I press some wrong buttons by mistake? Any help will be much appreciated! Thanks.
