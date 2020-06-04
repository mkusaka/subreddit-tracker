# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/guk02c/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/guk02c/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - June 01, 2020](https://www.reddit.com/r/androiddev/comments/guij23/weekly_questions_thread_june_01_2020/)
- url: https://www.reddit.com/r/androiddev/comments/guij23/weekly_questions_thread_june_01_2020/
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
## [3][With all these new restrictive rules, what is this sub for? Serious question.](https://www.reddit.com/r/androiddev/comments/gweyu0/with_all_these_new_restrictive_rules_what_is_this/)
- url: https://www.reddit.com/r/androiddev/comments/gweyu0/with_all_these_new_restrictive_rules_what_is_this/
---
With all these new rules, what can you post in this sub? Who thought these new rules would be helpful/good for community?   


Since new rules have been applied the quality and usefulness of this sub including number of posts has been reduced drastically, now it's just a generic links sharing hub with occasional lib discussion.  


No app/takedown posts?   
No help me posts?   
What the actual fuck?   


 `News for Android developers with the who, what, where when and how of the Android community.`   
Maybe rename this sub into AndroidHelpPage? Because its no longer useful.   
I have been following this sub for a while now, and ever since new rules have been applied, this sub went from very useful to meh.
## [4][cmdline-tools package is replacing the old tools package: was I the only one that missed this change?](https://www.reddit.com/r/androiddev/comments/gwe1w2/cmdlinetools_package_is_replacing_the_old_tools/)
- url: https://www.reddit.com/r/androiddev/comments/gwe1w2/cmdlinetools_package_is_replacing_the_old_tools/
---
So, yesterday I just noticed that there is a new SDK Tool package available, `cmdline-tools`:

&amp;#x200B;

[Android SDK Command-line Tools package](https://preview.redd.it/sqam76znzu251.png?width=586&amp;format=png&amp;auto=webp&amp;s=4cd58e939ef3e561a5deb0d9159456e27429fa21)

I decided to download and install it and, to my surprise, contains some of the same binaries available inside the good old `tools` directory:

    $ find . -name sdkmanager 
    ./cmdline-tools/latest/bin/sdkmanager
    ./tools/bin/sdkmanager
    
    $ find . -name avdmanager 
    ./cmdline-tools/latest/bin/avdmanager
    ./tools/bin/avdmanager

But the interesting part is that the binaries are different:

    $ find . -name sdkmanager -exec sha256sum {} \;
    42e587cd6795924403943856f8d5d62200272cc121ae89e5c91681967449c6e0  ./cmdline-tools/latest/bin/sdkmanager
    ca5dce516a93bea2c070ba6cc63d3d45436ce1c9a27f1173256b2aa98ae1cffa  ./tools/bin/sdkmanager
    
    $ find . -name avdmanager -exec sha256sum {} \;
    a5ca2a53294ea0b4af3d91fb59b3c8b7878c1d9eb7ecfc9d521c896ecbf7653e  ./cmdline-tools/latest/bin/avdmanager
    a98933db7c4d8dd376d67208316732ccf3a4917d0aea109302feda86bfc5a906  ./tools/bin/avdmanager

and the crazy part is that the version numbers are very different:

    $ find . -name sdkmanager -exec  {} --version \;
    4.0.0
    26.1.1

Please note that the new `cmdline-tools` package is the one using version 4.0.0!

I tried to search about this change but I can't find a single place where this is announced.

The official SDK Manager page still refers to the version from tools and the version number 25.2.3 or higher:

&amp;#x200B;

[sdkmanager](https://preview.redd.it/ixt9lxqr0v251.png?width=878&amp;format=png&amp;auto=webp&amp;s=7bf18a65326189610a185138047ac9fb60ce02e2)

But the Command line tools page advise instead to use the new `cmdline-tools` version:

[Command line tools](https://preview.redd.it/vjaoyp5z0v251.png?width=869&amp;format=png&amp;auto=webp&amp;s=c72072b028a78d3d9afcfd6618574c5ceec5168b)

Does anybody have more info on this `cmdline-tools` package? Is it safe to migrate to it or is it better to stick with the `tools` version? 

Why did the version number of the `sdkmanager` go from 26.x back to 4.x?
## [5][Kotlin 1.4-M2 Released](https://www.reddit.com/r/androiddev/comments/gwfyw5/kotlin_14m2_released/)
- url: https://blog.jetbrains.com/kotlin/2020/06/kotlin-1-4-m2-released
---

## [6][Coroutines: Suspending State Machines](https://www.reddit.com/r/androiddev/comments/gwchzr/coroutines_suspending_state_machines/)
- url: https://medium.com/@ragdroid/coroutines-suspending-state-machines-36b189f8aa60
---

## [7][Securing Network Data Tutorial for Android](https://www.reddit.com/r/androiddev/comments/gw20zi/securing_network_data_tutorial_for_android/)
- url: https://www.raywenderlich.com/10056112-securing-network-data-tutorial-for-android
---

## [8][does PDA important when flashing your phone ?](https://www.reddit.com/r/androiddev/comments/gwh95n/does_pda_important_when_flashing_your_phone/)
- url: https://www.reddit.com/r/androiddev/comments/gwh95n/does_pda_important_when_flashing_your_phone/
---
 i have a Samsung j5 stuck on Samsung logo so i am gonna flash it but its PDA is j500HXXS1APE1  
and i did not find the exact same firmware  , is it important ? or i can flash it with a firmware  with another PDA
## [9][Android Gradle Plugin can have better performance with Graalvm](https://www.reddit.com/r/androiddev/comments/gwb4mp/android_gradle_plugin_can_have_better_performance/)
- url: https://issuetracker.google.com/issues/157893911#comment3
---

## [10][Mobile and web application costs in 2020](https://www.reddit.com/r/androiddev/comments/gwg0jq/mobile_and_web_application_costs_in_2020/)
- url: https://blog.codemagic.io/mobile-and-web-app-costs-2020/
---

## [11][RxJava onNext only works with the mainThread, not newThread](https://www.reddit.com/r/androiddev/comments/gwfr8x/rxjava_onnext_only_works_with_the_mainthread_not/)
- url: https://www.reddit.com/r/androiddev/comments/gwfr8x/rxjava_onnext_only_works_with_the_mainthread_not/
---
When I switch from running the observeOn on the Main thread to a newThread, the onNext only runs once.

That only way I can get it to work is to keep it on the Main thread, but then I get that it does too much work on the Main Thread.

Without the setErrorHandler it just crashes (Also, can I use doOnError instead of RxJavaPlugins?)

PS. It works on the emulator fine, but it's on the physical device that the issue comes up.

    public void updatePie() {
    
        RxJavaPlugins.setErrorHandler(Functions.&lt;Throwable&gt;emptyConsumer());
    
        Observable&lt;Long&gt; intervalObservable = Observable
                .interval(1, TimeUnit.SECONDS)
              //.doOnError(Functions.&lt;Throwable&gt;emptyConsumer())
                .subscribeOn(Schedulers.io())
                .takeWhile(new Predicate&lt;Long&gt;() {
                    @Override
                    public boolean test(Long aLong) throws Exception {
    
                        if (isMyServiceRunning(MyService.class) == false) {
                            RxB = false;
                        }
                        return RxB;
                    }
                })
                .observeOn(Schedulers.newThread());
    
        intervalObservable.subscribe(new io.reactivex.Observer&lt;Long&gt;() {
            @Override
            public void onSubscribe(Disposable d) {
    
            }
    
            @Override
            public void onNext(Long aLong) {
    
                triple = mService.Time;
                entries.set(0, new PieEntry(mService.Time, "kronk"));
                entries.set(1, new PieEntry(mService.Time2, "notre dame"));
                pie_chart.notifyDataSetChanged();
                pie_chart.invalidate();
    
            }
    
            @Override
            public void onError(Throwable e) {
                Log.d(TAG, "Pie Update " + e.toString());
            }
    
            @Override
            public void onComplete() {
                Log.d(TAG, "Pie Update completed");
            }
        });
    }
## [12][Fuchsia is not a science experiment](https://www.reddit.com/r/androiddev/comments/gwfht5/fuchsia_is_not_a_science_experiment/)
- url: https://fuchsia.dev/fuchsia-src/concepts
---

