# androiddev
## [1][App Feedback Thread - March 14, 2020](https://www.reddit.com/r/androiddev/comments/fihm43/app_feedback_thread_march_14_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fihm43/app_feedback_thread_march_14_2020/
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
## [2][Weekly Questions Thread - March 09, 2020](https://www.reddit.com/r/androiddev/comments/ffsx46/weekly_questions_thread_march_09_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ffsx46/weekly_questions_thread_march_09_2020/
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
## [3][I have created this Motion Toast Library for Android Kotlin. It has various modes for creating beautiful toast.](https://www.reddit.com/r/androiddev/comments/fiyfav/i_have_created_this_motion_toast_library_for/)
- url: https://github.com/Spikeysanju/MotionToast
---

## [4][I made a customisable toggle button group library in Google's FlexboxLayout. It handles single/multi select, text + icons, smooth animations and more.](https://www.reddit.com/r/androiddev/comments/fizywq/i_made_a_customisable_toggle_button_group_library/)
- url: https://v.redd.it/87b5l43imtm41
---

## [5][Experience Requirements for Developing Android 🤔](https://www.reddit.com/r/androiddev/comments/fiz5n2/experience_requirements_for_developing_android/)
- url: https://www.reddit.com/r/androiddev/comments/fiz5n2/experience_requirements_for_developing_android/
---
I'm really interested in android dev so I wanted to know the experience requirements to start building a custom OS from android source such as should I know java and XML or should I learn assembler for the hardware compatibility and which tools are used that made all the OSs we see like One UI, MIUI , EMUI and others .
And thanks 😉
## [6][Is there an easy to understand tutorial for saving images from the Camera app and displaying them in my app?](https://www.reddit.com/r/androiddev/comments/fiuahx/is_there_an_easy_to_understand_tutorial_for/)
- url: https://www.reddit.com/r/androiddev/comments/fiuahx/is_there_an_easy_to_understand_tutorial_for/
---
I think this boils down to my not understanding URI's, or FileProvider.  I've read up on them a fair amount, but something just is not clicking for me.  Hopefully someone can point me to some resources that really dumb it down.

My current goal is to open the Camera app with a button in my application, take a picture, and when I return to my application it's written to storage, and shows up in a gridview.  As I take more pictures, they too are added to the gridview.  And since they're written to storage, they will persist when I re-open the application.

**What I can do:**

\- Have my app open the Camera app, and display the resulting image in a single ImageView.  I've accomplished this by opening the Camera app with Intent(MediaStore.ACTION\_IMAGE\_CAPTURE) and adding an EXTRA\_OUTPUT of my URI.  Then startActivityForResult.  Then onActivityResult I use BitmapFactory.decodeFile to get a Bitmap, then set it to the ImageView.

\- Display images from a collection of images in a gridview in a fragment in my activity.  Think an incredibly limited Google Photos album page, but only showing images taken within this application.

\-Save images to file in an external folder.

**What I'm not sure how to do:**

\- Find out what images (if any) exist in the storage of my application onCreate.

\- How to get the URI's from said images in storage.

\- Rework my displaying images in the gridview from just displaying images I've forced into the program to ones pulled from the URI's of stored images.  This one may actually not be very difficult, but until I've worked out the previous two items I won't know.
## [7][Any way to contact Amazon support for their App Store that isn't their useless contact form?](https://www.reddit.com/r/androiddev/comments/fikx5k/any_way_to_contact_amazon_support_for_their_app/)
- url: https://www.reddit.com/r/androiddev/comments/fikx5k/any_way_to_contact_amazon_support_for_their_app/
---
I have a Fire TV app that is totally free but now when I try to update it, Amazon tells me that I have purchases in there and should use their IAP. I emailed them 3 or 4 days ago about this and have yet to hear anything about it. Base on my past experiences with them, I expect to just get a canned response.

&amp;#x200B;

I also have another app which they keep refusing to list on Fire devices despite me removing pretty much all Google Play Services (still have Crashlytics but they usually allow that) and their support always sends me canned answers in reply. This is an app that on the Play Store has over 10 million downloads.

&amp;#x200B;

So is there a better way to contact Amazon? I don't care if I have to pay to get support, I just want a human to give me some answers.
## [8][A simple method to add fade-in and out effect to change theme without needing a restart (maybe someone can take a look and let me know if it's alright)](https://www.reddit.com/r/androiddev/comments/fj0cne/a_simple_method_to_add_fadein_and_out_effect_to/)
- url: https://www.reddit.com/r/androiddev/comments/fj0cne/a_simple_method_to_add_fadein_and_out_effect_to/
---
Let me start by saying I might be doing something wrong here, but some more experienced people can correct me if my method has any problems.

Here's how it looks:

[Forgive my designing skills](https://reddit.com/link/fj0cne/video/u3tz51iurtm41/player)

1. Added `getWindow().setWindowAnimations(R.style.WindowAnimation);` in the `onCreate()` of parent activity which I wanted to switch the theme for.
2. Where `R.style.WindowAnimation` is

&amp;#8203;

    &lt;style name="WindowAnimation"&gt;
    &lt;item name="android:windowEnterAnimation"&gt;@android:anim/fade_in&lt;/item&gt;
    &lt;item name="android:windowExitAnimation"&gt;@android:anim/fade_out&lt;/item&gt;
    &lt;/style&gt;  


3. Modified the recreate method of activity like below:

    @Override
        public void recreate() {
            super.recreate();
            if (getCurrentTheme() == 1) {
             AppCompatDelegate.setDefaultNightMode(AppCompatDelegate.MODE_NIGHT_YES);
            } else {
              AppCompatDelegate.setDefaultNightMode(AppCompatDelegate.MODE_NIGHT_NO);
            }
        }

4. Then I recreated the activity using `getActivity().recreate();` in that fragment whenever I wanted to switch the activity (in my case choosing the theme using RadioGroup).

And that added the fade effect to theme switching. I don't know if overriding `recreate()` method will have any bad consequences. Because it somehow only applies theme to some parts of the app if I don't override the `recreate()` like above. Maybe someone can let me know about that.
## [9][Simple introduction to Dagger DI by a Life way!](https://www.reddit.com/r/androiddev/comments/fizwho/simple_introduction_to_dagger_di_by_a_life_way/)
- url: https://link.medium.com/eH7FMnO0R4
---

## [10][Is there common standard practice to implement a custom navigation drawer menu?](https://www.reddit.com/r/androiddev/comments/fitrcl/is_there_common_standard_practice_to_implement_a/)
- url: https://www.reddit.com/r/androiddev/comments/fitrcl/is_there_common_standard_practice_to_implement_a/
---
Previously, all the while, I'm implementing "standard" navigation drawer menu, by supplying proper menu file to app:menu

        &lt;?xml version="1.0" encoding="utf-8"?&gt;
        &lt;androidx.drawerlayout.widget.DrawerLayout xmlns:android="http://schemas.android.com/apk/res/android"
            xmlns:app="http://schemas.android.com/apk/res-auto"
            xmlns:tools="http://schemas.android.com/tools"
            android:id="@+id/drawer_layout"
            android:layout_width="match_parent"
            android:layout_height="match_parent"
            android:fitsSystemWindows="true"
            tools:openDrawer="start"&gt;
    
            &lt;include
                layout="@layout/app_bar_main"
                android:layout_width="match_parent"
                android:layout_height="match_parent" /&gt;
    
            &lt;com.google.android.material.navigation.NavigationView
                android:id="@+id/nav_view"
                android:layout_width="wrap_content"
                android:layout_height="match_parent"
                android:layout_gravity="start"
                android:fitsSystemWindows="true"
                app:itemIconTint="?attr/mainMenuIconSelector"
                app:itemTextColor="?attr/mainMenuTextSelector"
                app:menu="@menu/activity_main_drawer" /&gt;
    
        &lt;/androidx.drawerlayout.widget.DrawerLayout&gt;

Recently, I need to construct a custom navigation drawer menu dynamically, by reading from DB and perform some logic manipulation. I first thought this is a pretty common pattern, and I can get some good code example/ documentation on how to do so. But, so far, I can't get a common standard practice on how to do so.

My plan is as follow so far (I'm not sure whether it is a good plan yet)

1. Have a DB reading and logic manipulation, within a Fragment call MenuFragment.
2. MenuFragment will use RecyclerView to show all menu items.
3. Add MenuFragment to NavigationView.

I'm getting close but I'm don't know how I can handle android:fitsSystemWindows="true". Currently, I'm stuck at that stage - [https://stackoverflow.com/questions/60688716/androidfitssystemwindows-true-doesnt-work-well-for-fragment-used-for-custom](https://stackoverflow.com/questions/60688716/androidfitssystemwindows-true-doesnt-work-well-for-fragment-used-for-custom)

Does anyone know how to implement custom navigation drawer menu in a proper way (By able to support android:fitsSystemWindows="true")? Do you mind to share the technique with me?

Thank you.
## [11][Control media playback through KeyEvents](https://www.reddit.com/r/androiddev/comments/fitcsa/control_media_playback_through_keyevents/)
- url: https://www.reddit.com/r/androiddev/comments/fitcsa/control_media_playback_through_keyevents/
---
Hey everyone! I'm developing an Android app which has a button whose purpose is to play and pause any music that's playing on the phone, regardless of the source of the audio, be it Spotify, Deezer, YT Music, or any other app. After a lot of googling I managed to find the following which does in fact pause and play music, but only affects the default android music player, no other media players.

    private static void pauseMusic(Context context, int keyCode) {
            KeyEvent keyEvent = new KeyEvent(KeyEvent.ACTION_DOWN, keyCode);
            Intent intent = new Intent(Intent.ACTION_MEDIA_BUTTON);
            intent.putExtra(Intent.EXTRA_KEY_EVENT, keyEvent);
            context.sendOrderedBroadcast(intent, null);
    
            keyEvent = new KeyEvent(KeyEvent.ACTION_UP, keyCode);
            intent = new Intent(Intent.ACTION_MEDIA_BUTTON);
            intent.putExtra(Intent.EXTRA_KEY_EVENT, keyEvent);
            context.sendOrderedBroadcast(intent, null);
        }
        
        public void pauseMusic(View v) {
            pauseMusic(getApplicationContext(), KeyEvent.KEYCODE_MEDIA_PLAY_PAUSE);
        }

Android's documentation on media buttons says the following:

&gt;Prior to Android 8.0 (API level 26), the system tries to send the event to an active media session. If there are multiple active media sessions, Android tries to choose a media session that is preparing to play (buffering/connecting), playing, or paused, rather than one that is stopped.

The minimum API I set for my project is Android 6.0. I have tried the app whilst Spotify is playing audio and instead of stopping Spotify, it stops Spotify and plays a song from the default player. I've also tried to resume Spotify and the same thing happens. Is there someting I'm missing? Or something I'm doing wrong? Any pointers to the right direction or better alternatives would be greatly appreciated!
## [12][If my phone cant get GPS location in a remote village, will a USB/bluetooth external GPS fix this?](https://www.reddit.com/r/androiddev/comments/fiu6lz/if_my_phone_cant_get_gps_location_in_a_remote/)
- url: https://www.reddit.com/r/androiddev/comments/fiu6lz/if_my_phone_cant_get_gps_location_in_a_remote/
---
I dont know what the issue is. Is it that the GPS chip is low quality or is it the satellite its using low quality? Im wondering if the solution is to get a external device that works with the Android phone that has better service which I guess is a combination of the chip and the satellite it uses...
