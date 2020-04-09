# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/fvy3fg/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/fvy3fg/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
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
## [3][COVID-19: Downloads Of Health &amp; Fitness Apps Increased By 40%](https://www.reddit.com/r/androiddev/comments/fxoycm/covid19_downloads_of_health_fitness_apps/)
- url: https://mobilemarketingreads.com/covid-19-downloads-of-health-fitness-apps-increased-by-40/
---

## [4][RelativeTime: a library that lets you define the language of relative time, e.g. "5 hours ago" or "1 month from now" and so on](https://www.reddit.com/r/androiddev/comments/fxek3t/relativetime_a_library_that_lets_you_define_the/)
- url: https://github.com/nihk/RelativeTime
---

## [5][Just found out that i can drag expressions into watches list in Android Studio .. very handy!](https://www.reddit.com/r/androiddev/comments/fwzhra/just_found_out_that_i_can_drag_expressions_into/)
- url: https://i.redd.it/9ld40zbzqir41.gif
---

## [6][Here are some handy Android Studio keyboard shortcuts that will speed up our development process](https://www.reddit.com/r/androiddev/comments/fxi410/here_are_some_handy_android_studio_keyboard/)
- url: https://themobilecoder.com/the-10-must-know-shortcuts-in-android-studio/
---

## [7][Error populating relational table with SQLite](https://www.reddit.com/r/androiddev/comments/fxrshm/error_populating_relational_table_with_sqlite/)
- url: https://www.reddit.com/r/androiddev/comments/fxrshm/error_populating_relational_table_with_sqlite/
---
I'm trying to populate an SQL table but I keep running into errors.
I have a cocktails table and an Ingredients table which are set up fine and have data in them. I've also created a cocktailIngredients table which is showing in DB Browswer fine. I'm then trying to populate this table with the cocktail ID and the ingredient ID from the two other tables with the following code:

    Cursor c = db.rawQuery("SELECT cocktailName FROM " + CocktailsTable.TABLE_NAME, null);

    if(c != null){
       if(c.moveToFirst()){
       do{
            ContentValues cv = new ContentValues();
            cv.put(CocktailIngredientTable.COLUMN_COCKTAIL_ID, c.getString(c.getColumnIndex(CocktailsTable._ID)));
            db.insert(CocktailIngredientTable.TABLE_NAME, null, cv);
       } while (c.moveToNext());
    }
    c.close();

I'm getting the error:  Couldn't read row 0, col -1 from CursorWindow

When I don't put in the db.insert function, the error doesn't come up, but the table doesn't populate either so I think it's something to do with this? I don't see how the cursor itself could be the issue since i'm checking that its not null first.

Can anyone help with this?
## [8][Unable to run two or more application in parallel using android studio 3.6](https://www.reddit.com/r/androiddev/comments/fxrlaf/unable_to_run_two_or_more_application_in_parallel/)
- url: https://www.reddit.com/r/androiddev/comments/fxrlaf/unable_to_run_two_or_more_application_in_parallel/
---
I've just upgraded android studio from 3.4 to 3.6, and found out that I can not run 2 instances of the same application in parallel (which is very important for development).

The first run starts without any problems, but the second one will hang until the first one is terminated, then it starts, which is not very useful. (see image attached)

https://preview.redd.it/v4q2boabdsr41.png?width=2126&amp;format=png&amp;auto=webp&amp;s=b29c7d9e6114b4e62306336da9511abdd3796e51

This is an 'Application' run configuration for a gradle project, and I have 'Allow parallel run' checked.

I never had this issue on android studio 3.4... and it doesn't make any sense why this is not possible.

Also, I started seeing this issue after the upgrade [https://stackoverflow.com/questions/60982242/how-to-fix-nosuchelementexception-scanner](https://stackoverflow.com/questions/60982242/how-to-fix-nosuchelementexception-scanner) I'm sure its related somehow...

Any help fixing this issue would be appreciated, thanks!
## [9][Don't use AppLovin if you're trying out new ad networks in your app](https://www.reddit.com/r/androiddev/comments/fxr857/dont_use_applovin_if_youre_trying_out_new_ad/)
- url: https://www.reddit.com/r/androiddev/comments/fxr857/dont_use_applovin_if_youre_trying_out_new_ad/
---
I've been trying to increase fill rate in my app, and so trying out new ad networks. But whenever I google about them I don't find much info about them from people who actually tried them out.

So just in case you're trying to add AppLovin by looking at some fake ass paid reviews online, don't do it.

I tried them out for 3 days, they have incredibly high CTR, but they only pay you per conversion. I might not have huge impressions here because it was only for 3 days, but you can still get the idea of how $hit it is in terms of eCPM.

6100 impressions and $0.73 in revenue? $0.12 eCPM? That's lower than what banner ad from India pays me at FAN.

https://preview.redd.it/xb4xdeig6sr41.jpg?width=1521&amp;format=pjpg&amp;auto=webp&amp;s=dccceba3d8678ba395c8aec30e5a7e51b99e14fa

https://preview.redd.it/r3ryv0ks5sr41.jpg?width=1500&amp;format=pjpg&amp;auto=webp&amp;s=a05eb5b26dc28c34a30618884ae78bb14aeb3123

So you're saying 75 people from USA clicked on the ad, and I get $0.61?

I know 6200 impressions might not be enough data, but an ad network that gives $0.4 eCPM for US traffic isn't worth using imo.
## [10][Android TextView Styling](https://www.reddit.com/r/androiddev/comments/fxr35c/android_textview_styling/)
- url: https://www.reddit.com/r/androiddev/comments/fxr35c/android_textview_styling/
---
I have this following string res

`&lt;string name="description"&gt;Your trial ends on`

`&lt;annotation font="semi_bold_text"&gt;%s&lt;/annotation&gt;.`

`You will be charged after your trial ends`

`&lt;/string&gt;`

I am using custom font(Montserrat) in my application. I want to use montserrat-semibold for the expiry date and montserrat-regular for the remaining text.

&amp;#x200B;

I will set the expiry date using databinding

`android:text="@{@string/description(expiry_date)}`

or by using resources getString() method with varArgs

`val expiryDate = "Apr 9, 2020"`

  `context.resources.getString(R.string.description, expiryDate)`

But according to documentation,

[https://developer.android.com/guide/topics/resources/string-resource#StylingWithAnnotations](https://developer.android.com/guide/topics/resources/string-resource#StylingWithAnnotations) 

they used annotations for static string alone.

Is there anyway we can combine string formatting with annotations?

Note: We cannot able to use Html.fromHtml() method as they don't support custom fonts as per android official documentation.

[https://developer.android.com/guide/topics/resources/string-resource#StylingWithHTML](https://developer.android.com/guide/topics/resources/string-resource#StylingWithHTML)

&amp;#x200B;

I raised the same issue in stack overflow. Here is the link.

[https://stackoverflow.com/questions/61120366/android-textview-multiple-fonts-and-string-formatting](https://stackoverflow.com/questions/61120366/android-textview-multiple-fonts-and-string-formatting)
## [11][What are the best practices for admob ad placement to increase revenue?](https://www.reddit.com/r/androiddev/comments/fxqybu/what_are_the_best_practices_for_admob_ad/)
- url: https://www.reddit.com/r/androiddev/comments/fxqybu/what_are_the_best_practices_for_admob_ad/
---
I want to know how can i increase my ad revenue with more optimized ad scheme, i am currently using an interstitial on splash and it is giving me a good Return on my investment. i was wondering if i could get some suggestions of how to use ads to be more profitable. Currently i am using Natives in my activities and banners as well, and 1 interstitial on Splash
## [12][D8, R8 and enums - Kotlin Vocabulary](https://www.reddit.com/r/androiddev/comments/fxqgh6/d8_r8_and_enums_kotlin_vocabulary/)
- url: https://www.youtube.com/watch?v=lTo03M2HzFY
---

