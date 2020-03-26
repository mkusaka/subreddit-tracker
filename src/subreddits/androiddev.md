# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/fnjcmy/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/fnjcmy/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
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
## [3][Create awesome app screenshots for PlayStore](https://www.reddit.com/r/androiddev/comments/fp98il/create_awesome_app_screenshots_for_playstore/)
- url: https://www.reddit.com/r/androiddev/comments/fp98il/create_awesome_app_screenshots_for_playstore/
---
I know you all are being extra productive in this difficult time and have multiple mobile apps lined up to be released, so please don't upvote me too fast in excitement, let me pitch something first.

We made the daunting work of creating screenshots for GooglePlay a little bit more enjoyable. You might think, well, there were other ones already. Yes, we thought so too, but they didn't have the styles nor the devices we wanted and it's also cool to do the project yourself.

So do use our service if it comes handy in your next app. Go to [previewed.app](https://previewed.app), select a device, choose a style, upload a screenshot and whoalla, pretty PlayStore.

You end up with something that looks like 

[Please use PC to keep your sanity.](https://preview.redd.it/4w4v1oi3yzo41.png?width=2368&amp;format=png&amp;auto=webp&amp;s=85d2c795f41203eb07c498ca2ffea1cf6b81a405)

P.S. Thanks for the downvotes!
## [4][I guess I should feel flattered but not sure this is fair.](https://www.reddit.com/r/androiddev/comments/foyiho/i_guess_i_should_feel_flattered_but_not_sure_this/)
- url: https://i.redd.it/ehd5kzs3yvo41.png
---

## [5][How does my application work on an Android device?](https://www.reddit.com/r/androiddev/comments/fp7p4o/how_does_my_application_work_on_an_android_device/)
- url: https://www.reddit.com/r/androiddev/comments/fp7p4o/how_does_my_application_work_on_an_android_device/
---
Hello, everyone! I want to ask you about what's going on under the hood when you created or you're creating your own mobile application for Android. So, let's get started from the underlying technologies.

1. On you computer you have JDK. You wrote some code and it was compiled by javac into bytecode. Then does my byte code go to .apk file or it should be converted into machine code by JVM and only then into .apk file?
2. We have our application on an Android device. Imagine that I want to run it. How does it work? Should it compile my code again or .apk contains machine code already? If it doesn't contain, what will compile my code on an Android device? Kernel?
## [6][Mocking is not practical — Use fakes](https://www.reddit.com/r/androiddev/comments/fovoch/mocking_is_not_practical_use_fakes/)
- url: https://medium.com/@june.pravin/mocking-is-not-practical-use-fakes-e30cc6eaaf4e
---

## [7][AutoCompleteTextView - Disable dropdown behavior](https://www.reddit.com/r/androiddev/comments/fpbbc8/autocompletetextview_disable_dropdown_behavior/)
- url: https://www.reddit.com/r/androiddev/comments/fpbbc8/autocompletetextview_disable_dropdown_behavior/
---
So I have subclassed autocompletetextview to meet my requirements but I'm stuck with one issue with the dropdown behavior, whenever I click on the view the dropdown fluctuates,  the number of times I click on the view, it fluctuates every time, I want it to be persistent and dismiss only on outside touch or back press.

here is an example [https://imgur.com/PmVBANa](https://imgur.com/PmVBANa)

&amp;#x200B;

    package com.momentsnap.android.searchableSpinner;
    
    import android.content.Context;
    import android.util.AttributeSet;
    import android.view.MotionEvent;
    
    import androidx.appcompat.widget.AppCompatAutoCompleteTextView;
    
    import com.momentsnap.android.teamlist.ui.entity.OnSpinnerEventsListener;
    
    import java.util.Calendar;
    
    import timber.log.Timber;
    
    
    public class MyAutoCompleteTextView extends AppCompatAutoCompleteTextView {
    
        private static final int MAX_CLICK_DURATION = 200;
        private long startClickTime;
        private OnSpinnerEventsListener mListener;
        private boolean mOpenInitiated = false;
    
    
        public MyAutoCompleteTextView(Context context) {
            super(context);
        }
    
        public MyAutoCompleteTextView(Context context, AttributeSet attrs) {
            super(context, attrs);
        }
    
        public MyAutoCompleteTextView(Context context, AttributeSet attrs, int defStyleAttr) {
            super(context, attrs, defStyleAttr);
        }
    
        public void performFiltering(CharSequence text) {
            super.performFiltering(text, 0);
        }
    
        public void setSpinnerEventsListener(OnSpinnerEventsListener onSpinnerEventsListener) {
            mListener = onSpinnerEventsListener;
        }
    
        //Override the parent class method, so that the drop-down box selection is also generated
        // when 0 input is possible. By default, at least 1 character is required for input.
        @Override
        public boolean enoughToFilter() {
            if (!isEnabled() || getAdapter() == null) {
                return false;
            }
            if (getText().length() == 0) {
                return true;
            } else {
                return super.enoughToFilter();
            }
        }
    
        @Override
        public boolean performClick() {
            // register that the Spinner was opened so we have a status
            // indicator for the activity(which may lose focus for some other
            // reasons)
            mOpenInitiated = true;
            if (mListener != null) {
                mListener.onSpinnerOpened();
            }
            return super.performClick();
        }
    
        public void performClosedEvent() {
            mOpenInitiated = false;
            if (mListener != null) {
                mListener.onSpinnerClosed();
            }
        }
    
        /**
         * A boolean flag indicating that the Spinner triggered an open event.
         *
         * @return true for opened Spinner
         */
        public boolean hasBeenOpened() {
            return mOpenInitiated;
        }
    
        @Override
        public void onWindowFocusChanged(boolean hasWindowFocus) {
            Timber.d("onWindowFocusChanged");
            super.onWindowFocusChanged(hasWindowFocus);
            if (hasBeenOpened() &amp;&amp; hasWindowFocus) {
                Timber.d("closing popup");
                performClosedEvent();
            }
        }
    
        //Increase the chance of drop-down box display.
        @Override
        public boolean onTouchEvent(MotionEvent event) {
    
            if (!isEnabled() || getAdapter() == null) {
                return false;
            }
    
            switch (event.getAction()) {
                case MotionEvent.ACTION_DOWN: {
                    startClickTime = Calendar.getInstance().getTimeInMillis();
                    break;
                }
                case MotionEvent.ACTION_UP: {
                    long clickDuration = Calendar.getInstance().getTimeInMillis() - startClickTime;
                    if (clickDuration &lt; MAX_CLICK_DURATION) {
                        if (!isPopupShowing()) {
                            performFiltering(getText());
                            showDropDown();
                        }
                    }
                    break;
                }
            }
            return super.onTouchEvent(event);
        }
    }
## [8][ADB: Cant connect to emulator but can issue commands](https://www.reddit.com/r/androiddev/comments/fpb2r1/adb_cant_connect_to_emulator_but_can_issue/)
- url: https://www.reddit.com/r/androiddev/comments/fpb2r1/adb_cant_connect_to_emulator_but_can_issue/
---
Hi guys,

&amp;#x200B;

Just started working with adb and emulated androids.

I have an android emulator with Android studio.

If I do "adb devices" it shows up as "emulator-5554" , but second column just says "device".

I've tried to install an apk and check the packages and it all seemed fine, I just did "adb -e &lt;command&gt;"

But I can't do "adb connect emulator-5554:5555", should I be worried or is that expected behaviour from emulators? 

&amp;#x200B;

Thank you
## [9][ART and Davlik under the hood?](https://www.reddit.com/r/androiddev/comments/fpaoxk/art_and_davlik_under_the_hood/)
- url: https://www.reddit.com/r/androiddev/comments/fpaoxk/art_and_davlik_under_the_hood/
---
Hello, everyone! I asked the question here and you helped me where I should dig on that topic. So, now I've read some articles about this topic but I found a lot of unnecessary information about that. I need to get only superficially knowledge which will be useful.

As I understand, when you run your Java code on PC you have JVM but when Google created Android OS he couldn't afford use JVM in their project free. Hence, they created their own VM's for running Android applications. Once upon a time we had Davlik. It translates your Java bytecode into executable .dex files during your application is running. Now we have ART. It translates your Java bytecode into executable ART files when you're installing your app and move the code into your storage. Every application is running on ART can be run on Davlik but not every application is running on Davlik can be run on ART. ART is faster because you compiled the source code only one time when you were installing the app. It uses less performance of you device.

I hope you can tell me if something's wrong.

*Updated:* therefore, every application is running on its own virtual machine. It helps in debugging and also it prevents malicious software from their impact on the Linux kernel of your device.
## [10][Emulator 30.0.5 Stable](https://www.reddit.com/r/androiddev/comments/fowk50/emulator_3005_stable/)
- url: https://androidstudio.googleblog.com/2020/03/emulator-3005-stable.html
---

## [11][Before diving into this projects: What's your opinion on writing Python code for Android?](https://www.reddit.com/r/androiddev/comments/fp9oru/before_diving_into_this_projects_whats_your/)
- url: https://www.reddit.com/r/androiddev/comments/fp9oru/before_diving_into_this_projects_whats_your/
---
 tldr: Would you reccomend to write Android Apps in Python instead of Java?

Hallo fellow redditors!

I hope you can help me here with some advice.

A  friend and I planned on programming Alexa Skills using Python. I just  began learning programming and after I told my friend, who is a software  developer, he got hooked with my idea and we decided to write the Alexa  Skill in Python.

Now, after a  month however, we realised that the Amazon Web Services suck. Amazon is  very restrictive in terms of what you are allowed to do and how to do it  when programming Alexa Skills.

So we decided to work on an Android App instead.

Only  problem is: Thus far I only know the basics of Python and my friends  thinks that it would be way, way  easier to code an App in Java. Which  means that I'd have to start learning Java now.... :/

(and after taking a glimpse it looks scary...haha!)

So...  what do you think? I definately want to keep up studying Python, but  would you say that it is better to code in Java when it comes to Apps?
## [12][D8 Optimization: Assertions – Jake Wharton](https://www.reddit.com/r/androiddev/comments/for73y/d8_optimization_assertions_jake_wharton/)
- url: https://jakewharton.com/d8-optimization-assertions/
---

