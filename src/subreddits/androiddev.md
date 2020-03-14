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
## [3][Is there a way to draw a shadow like the one in the image?](https://www.reddit.com/r/androiddev/comments/fihmy4/is_there_a_way_to_draw_a_shadow_like_the_one_in/)
- url: https://www.reddit.com/r/androiddev/comments/fihmy4/is_there_a_way_to_draw_a_shadow_like_the_one_in/
---
&amp;#x200B;

https://preview.redd.it/1bv9ksz4umm41.png?width=944&amp;format=png&amp;auto=webp&amp;s=584fdc3575219c1ef1019db04487b05fe58630b0
## [4][Update firebase queries without changing RecyclerView adapter using FirebaseUI - Android](https://www.reddit.com/r/androiddev/comments/fie4mk/update_firebase_queries_without_changing/)
- url: https://link.medium.com/pqIF5SiaP4
---

## [5][How do you set up two CountUp timers in one class?](https://www.reddit.com/r/androiddev/comments/figumw/how_do_you_set_up_two_countup_timers_in_one_class/)
- url: https://www.reddit.com/r/androiddev/comments/figumw/how_do_you_set_up_two_countup_timers_in_one_class/
---
Hi, I'm trying to set up 2 CountUp Timers in my service class, but switching between the two (Pausing one and starting the other and vise versa) keeps throwing the timer off, and adding minutes worth in seconds.

Is there a simple way to implement a CountUp timer to have the two be independent of each other?

Here's the code for reference:

    private Handler customHandler = new Handler();
    
        private Runnable updateTimerThread = new Runnable() {
    
            public void run() {
    
                timeInMilliseconds = SystemClock.elapsedRealtime() - startTime;
    
                updatedTime = timeSwapBuff + timeInMilliseconds;
    
                secs = (int) (updatedTime / 1000);
                mins = secs / 60;
                secs = secs % 60;
                milliseconds = (int) (updatedTime % 1000);
                //Log.d(TAG, "Secs: " + secs + " Min: " +  mins);
                customHandler.postDelayed(this, 0);
            }
        };

&amp;#x200B;

    public void startTimer(){
            if(firsttime){
                startTime = SystemClock.uptimeMillis();
                customHandler.postDelayed(updateTimerThread, 0);
                firsttime = false;
            }
        }

&amp;#x200B;

    public void pauseTimer(){
            timeSwapBuff += timeInMilliseconds;
            customHandler.removeCallbacks(updateTimerThread);
            firsttime = true;
        }

I made 2 of each to get the two timers, but switching between them is cause count problems.

Thank you for your help.
## [6][Connect two devices via the internet](https://www.reddit.com/r/androiddev/comments/fihxca/connect_two_devices_via_the_internet/)
- url: https://www.reddit.com/r/androiddev/comments/fihxca/connect_two_devices_via_the_internet/
---
I want to do this for a project: One android phone is at my home and connected to the internet. I want to remotely send a command to the phone, for example take a picture and send it back. Since this is only for a small project I want the cheapest option (preferably free). The devices should still be able to use doze. What would I need for that? Sorry if this is the wrong sub, I'd be glad to be directed to a better suited one.
## [7][Where is best to read dev chat on Fintech, IoT and AI?](https://www.reddit.com/r/androiddev/comments/fiapmw/where_is_best_to_read_dev_chat_on_fintech_iot_and/)
- url: https://www.reddit.com/r/androiddev/comments/fiapmw/where_is_best_to_read_dev_chat_on_fintech_iot_and/
---
As per title really! 
Taking this work from home period to start learning more about android (Thanks Covid).
Wondered if anyone could share their favourite place to get involved in discussion around Fintech, IoT and AI on the platform?

I'm already a lurker on here, but looking to diversify my reading!
## [8][Composable way to listen for onActivityResult.](https://www.reddit.com/r/androiddev/comments/fi4t85/composable_way_to_listen_for_onactivityresult/)
- url: https://www.reddit.com/r/androiddev/comments/fi4t85/composable_way_to_listen_for_onactivityresult/
---
Finally this is coming.. [https://issuetracker.google.com/issues/125158199](https://issuetracker.google.com/issues/125158199)
## [9][How to handle equal movement on all screen sizes](https://www.reddit.com/r/androiddev/comments/fii7c6/how_to_handle_equal_movement_on_all_screen_sizes/)
- url: https://www.reddit.com/r/androiddev/comments/fii7c6/how_to_handle_equal_movement_on_all_screen_sizes/
---
So I have a game where many things moves.

The normal way that I would do this is to assign a speed to a object and change it's position with the speed every frame.

Then comes the the devices that cannot keep a constant time frame, so I add Delta and multiply the speed based on fps, so now it moves the same distance over time no matter the fps.

Now then Android exists on a infinite number of devices with different resolutions, screen sizes etc. I have not really had to handle this problem before. How to handle this without making the game too hard or too easy on some devices.

Say that I have a object moving from the far left edge of the screen into the middle, and another object should come from the top and hit middle after some time has passed.

On my phone it would work nice since it is timed on a device where I know the approx time of left object and then I spawn the top object when it seems right.

But on much bigger screens the top object might arrive before the left object, or at the same time etc.

You get the problem, so how could I solve this in a nice way?

cheers
## [10][OxygenOS Open Beta 10 for OnePlus 7 Pro brings Feb patch](https://www.reddit.com/r/androiddev/comments/fihibf/oxygenos_open_beta_10_for_oneplus_7_pro_brings/)
- url: https://9to5google.com/2020/03/14/oxygenos-open-beta-10-oneplus-7-pro/
---

## [11][HELP - APDU VISA MSD!](https://www.reddit.com/r/androiddev/comments/fih0cv/help_apdu_visa_msd/)
- url: https://www.reddit.com/r/androiddev/comments/fih0cv/help_apdu_visa_msd/
---
Hello, could someone help me update the payment protocol of an NFC application that is using a discontinued protocol?


When making the transaction, an error occurs at the POS terminal: Communication error with the card.
## [12][Proper way to get my GPS coordinates using MVVM pattern?](https://www.reddit.com/r/androiddev/comments/fi5lra/proper_way_to_get_my_gps_coordinates_using_mvvm/)
- url: https://www.reddit.com/r/androiddev/comments/fi5lra/proper_way_to_get_my_gps_coordinates_using_mvvm/
---
Hi, I'm trying to develop a map app that should be capable of getting coordinates from an API as well as displaying my own location on an OSM MapView. Right now I'm using FusedLocationProvider to acomplish that however I'm having problems understanding where should I check periodically for my location in order to keep with the MVVM pattern.

I understand that I shouldn't reference the my view on my viewmodel because that can lead to leaks, same with my coordinates repository and since FusedLocationProvider needs a context to get the coordinates I can't use it on my viewmodel or repository but leaving it in my view seems wrong.

What would be the correct way to do this?
