# androiddev
## [1][App Feedback Thread - March 28, 2020](https://www.reddit.com/r/androiddev/comments/fqjai3/app_feedback_thread_march_28_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fqjai3/app_feedback_thread_march_28_2020/
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
## [2][Weekly Questions Thread - March 23, 2020](https://www.reddit.com/r/androiddev/comments/fni20d/weekly_questions_thread_march_23_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fni20d/weekly_questions_thread_march_23_2020/
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
## [3][[Part 2] This guy decompiled my app, so I reported this app 14 days ago, and Google took down this app, but now this app is live again and still using my code. And this time it's worst he is showing ads in my own ranked keywords.](https://www.reddit.com/r/androiddev/comments/fq4mz7/part_2_this_guy_decompiled_my_app_so_i_reported/)
- url: https://i.redd.it/22jpa7rtu9p41.jpg
---

## [4][Navigation Component Using Custom Views](https://www.reddit.com/r/androiddev/comments/fqgob0/navigation_component_using_custom_views/)
- url: https://sasikanth.dev/posts/navigaiton-component-using-custom-views
---

## [5][What is the most efficient way of passing arguments into fragments](https://www.reddit.com/r/androiddev/comments/fqi8zx/what_is_the_most_efficient_way_of_passing/)
- url: https://www.reddit.com/r/androiddev/comments/fqi8zx/what_is_the_most_efficient_way_of_passing/
---
In my current project we follow the normal practice of passing arguments like so  
```
    companion object {
        const val TAG = "variableAmountFragment"
        const val CURRENCY_DETAILS_TAG = "currencyDetails"
        const val MIN_MEALS = "minMeals"
        const val TYPE = "variableAmountType"
        fun newInstance(currencyDetails: CurrencyDetails, minMeals: Int, variableAmountType: VariableAmountType) = VariableAmountFragment().apply {
            val bundle = Bundle().apply {
                putParcelable(CURRENCY_DETAILS_TAG, currencyDetails)
                putInt(MIN_MEALS, minMeals)
                putSerializable(TYPE, variableAmountType)
            }
            arguments = bundle
        }
    }
```

To retrieve the arguments we do this
```
private fun setState(savedInstanceState: Bundle?) {
    val bundle = savedInstanceState ?: requireArguments()
    currencyDetails = bundle.getParcelable(CURRENCY_DETAILS_TAG)!!
    minMeals = bundle.getInt(MIN_MEALS)
    type = bundle.getSerializable(TYPE) as VariableAmountType
}
```

To retain the state we do
```
override fun onSaveInstanceState(outState: Bundle) {
    outState.putParcelable(CURRENCY_DETAILS_TAG, currencyDetails)
    outState.putInt(MIN_MEALS, minMeals)
    outState.putSerializable(TYPE, type)
    super.onSaveInstanceState(outState)
}
```

Thats 23 LOC to pass in 3 arguments. In a Kotlin class it would be 1 line of a constructor. This process seems unreasonable and error prone. Are there no better options in this day and age?
## [6][Anyone successfully implemented FAN sandwich model?](https://www.reddit.com/r/androiddev/comments/fqk93v/anyone_successfully_implemented_fan_sandwich_model/)
- url: https://www.reddit.com/r/androiddev/comments/fqk93v/anyone_successfully_implemented_fan_sandwich_model/
---
I've been trying to wrap my head around facebooks sandwich model and how I can implement it through MoPub? I attended their webinar few days ago where they recommended doing this to increase ad revenue.

You can find info about it here -&gt;  [https://www.facebook.com/help/publisher/821011588003559](https://www.facebook.com/help/publisher/821011588003559) 

Basically you set up for example 3 placements. One which has very high CPM, another one with lower CPM and so on. You're supposed to get higher CPM if you set it up like that? But I'm unable to figure out how I can map those multiple ad units to a single one on MoPub?
## [7][Benefits if learning deprecated technology](https://www.reddit.com/r/androiddev/comments/fqevju/benefits_if_learning_deprecated_technology/)
- url: https://www.reddit.com/r/androiddev/comments/fqevju/benefits_if_learning_deprecated_technology/
---
Am I wasting my time learning AsyncTask, Loaders, Content Providers, etc? Are these useful for the event that I might run into them at a company in the future? Are they helpful to better understand how their evolved predecessors came to be? We're being taught a lot of obsolete libraries at Udacity, but I'm trying to motivate myself with a justification to get down with them and start coding them into my projects. I definitely won't be using them for my future apps, but I still wanna know what you veterans think about this; is there any benefit for me to get experience with obsolete libraries?
## [8][Stay Logged In Across Life Cycle](https://www.reddit.com/r/androiddev/comments/fqil67/stay_logged_in_across_life_cycle/)
- url: https://youtu.be/h1iXyPhRkHA
---

## [9][The Night version of my drawables are applied to a Light theme when I set the system Night mode to ON (Android 9). How to avoid this behavior?](https://www.reddit.com/r/androiddev/comments/fqi3vy/the_night_version_of_my_drawables_are_applied_to/)
- url: https://www.reddit.com/r/androiddev/comments/fqi3vy/the_night_version_of_my_drawables_are_applied_to/
---
The result is white images on white background. Is there a clean way to avoid resolving of -night resources in case of Light theme? That defeats the whole idea of having the "-night" qualifier!
## [10][[It's 2020] Does anyone still use Genymotion ?](https://www.reddit.com/r/androiddev/comments/fpyhxg/its_2020_does_anyone_still_use_genymotion/)
- url: https://www.reddit.com/r/androiddev/comments/fpyhxg/its_2020_does_anyone_still_use_genymotion/
---
I use Docker to deploy and test my web app, but Docker use Hyper-V, which would never work with Virtual Box (used by Genymotion). So i am planning to switch back to Android Studio Emulator.

Is this a good solution ?, how about the performance of AS Emulator compared to Genymotion ? Which emulator are you using ?

Thanks in advance

&amp;#x200B;

[View Poll](https://www.reddit.com/poll/fpyhxg)
## [11][Google store update times](https://www.reddit.com/r/androiddev/comments/fq7lkr/google_store_update_times/)
- url: https://www.reddit.com/r/androiddev/comments/fq7lkr/google_store_update_times/
---
Google said to expect delays of up to 7 days or more - they took about an hour tonight!
## [12][How to capture heap dump from Android App?](https://www.reddit.com/r/androiddev/comments/fqel5f/how_to_capture_heap_dump_from_android_app/)
- url: https://blog.heaphero.io/2018/06/04/how-to-capture-heap-dump-from-android-app-3-options/
---

