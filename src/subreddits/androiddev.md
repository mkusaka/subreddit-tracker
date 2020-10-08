# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/j5ipv0/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/j5ipv0/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - October 05, 2020](https://www.reddit.com/r/androiddev/comments/j5hb6o/weekly_questions_thread_october_05_2020/)
- url: https://www.reddit.com/r/androiddev/comments/j5hb6o/weekly_questions_thread_october_05_2020/
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
## [3][QUESTION: Taking too long to get results of Google Android Associate developer test](https://www.reddit.com/r/androiddev/comments/j795rd/question_taking_too_long_to_get_results_of_google/)
- url: https://www.reddit.com/r/androiddev/comments/j795rd/question_taking_too_long_to_get_results_of_google/
---
Hello Everyone,

I completed my Google Android Associate Developer assessment and exit Interview on the 15th of May 2020, and my assessment as been being graded since then, its about 5months now. I have reached out to Google support but they have not provided any actual solutions or reasons.   
Is this a reasonable amount of time for the results to still be processing even  with Covid-19 restrictions, can anyone with more information help me out.

Thank you
## [4][Collate all categories in SQLite database for "All questions"](https://www.reddit.com/r/androiddev/comments/j7a4jp/collate_all_categories_in_sqlite_database_for_all/)
- url: https://www.reddit.com/r/androiddev/comments/j7a4jp/collate_all_categories_in_sqlite_database_for_all/
---
Hello all,

I recently threw myself into the deep end and set out to develop an android app with very little experience in coding. I've followed advice given in this subreddit on my previous post and have been following a YouTube series by Coding in flow as [this](https://youtu.be/mONNvMIaulY) is exactly what I'm looking to do and feel like I'm learning a lot. I have adapted the code for my needs but is mostly the same (can be seen in his description box).

  


I'm struggling to get my head around the following: how do I make an "ALL" section which collates all questions, regardless of category. I have a feeling I'm overthinking this and would appreciate any input.

  


On a unrelated note, I've used his code for the spinner and it seems offset to the right despite being aligned to the start of the START QUIZ button. Any ideas?

  


I'm not really sure where to post these kind of basic questions. I unfortunately don't know anyone in the field who could help me out. I tried at stack overflow in code review but my post got removed.

  


Thanks!
## [5][Is Kotlin the actual standard best choice to develop Android apps? Is it still Java? Will Kotlin in the future be the standard choice?](https://www.reddit.com/r/androiddev/comments/j7a48g/is_kotlin_the_actual_standard_best_choice_to/)
- url: https://www.reddit.com/r/androiddev/comments/j7a48g/is_kotlin_the_actual_standard_best_choice_to/
---

## [6][Is it possible to load more items from a listing API before reaching the bottom of the list?](https://www.reddit.com/r/androiddev/comments/j791wh/is_it_possible_to_load_more_items_from_a_listing/)
- url: https://www.reddit.com/r/androiddev/comments/j791wh/is_it_possible_to_load_more_items_from_a_listing/
---
Hello!

I am building an app that uses Paging 3 and a REST API through Retrofit, to fetch items. I have a recyclerview and I want to load more items and add them to the list before the user reaches the bottom. I'm also using LiveData and a combination of Kotlin/Java but I don't think it really matters.

Any help would be greatly appreciated,
Thanks.
## [7][IntelliJ find and replace with regex](https://www.reddit.com/r/androiddev/comments/j7cqmx/intellij_find_and_replace_with_regex/)
- url: https://medium.com/@vashisthg/intellij-find-and-replace-with-regex-a3f5aa885659
---

## [8][Using Github Actions to Automate Our Release Process](https://www.reddit.com/r/androiddev/comments/j6q5k1/using_github_actions_to_automate_our_release/)
- url: https://riggaroo.dev/using-github-actions-to-automate-our-release-process/
---

## [9][Consume REST API from Kotlin?](https://www.reddit.com/r/androiddev/comments/j7b3wr/consume_rest_api_from_kotlin/)
- url: https://www.reddit.com/r/androiddev/comments/j7b3wr/consume_rest_api_from_kotlin/
---
Hi all,

No matter what method I try (googling like a demon), I cannot get this to work.

Nearly every single example I have found online doesn't compile - and the couple that do, simply don't work, or crash the app out with no error the moment I attempt to connect.

Someone here must know the easiest way to make a simple GET request to a third-party API and iterate through the JSON result?

I'm starting to wonder if it's possible - despite, I should think, nearly every app out there needing to do something similar.

So far, this is the closest I've got:

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        val text = "Hello toast!"
        val duration = Toast.LENGTH_SHORT
        val toast = Toast.makeText(applicationContext, text, duration)
        toast.show()

        val button = findViewById&lt;Button&gt;(R.id.button)
        button?.setOnClickListener()
        {
            Toast.makeText(
                this@MainActivity,
                "hello again", Toast.LENGTH_LONG
            ).show()
            sentGet()
        }
    }

    private fun sentGet() {
        val url = URL("http://www.google.com/")

        with(url.openConnection() as HttpURLConnection) {
            requestMethod = "GET"  // optional default is GET

            println("\nSent 'GET' request to URL : $url; Response Code : $responseCode")

        }
    }

But when I click the button, the app simply disappears without a word.
## [10][Worth nagging users for a review?](https://www.reddit.com/r/androiddev/comments/j70wr4/worth_nagging_users_for_a_review/)
- url: https://www.reddit.com/r/androiddev/comments/j70wr4/worth_nagging_users_for_a_review/
---
I think it's a bad user experience to ask a user to rate the app or show an in-app rating. We interrupt the user and then ask for a favor.

Wanted to find out if that was just my personal opinion. Ran a small survey and Android users agree it is annoying... but users also said they wouldn't give a rating otherwise.

How do you balance the need for reviews without irritating users?

Has anything you've tried backfired?

Data: [https://datastudio.google.com/s/luyhD0TXrVs](https://datastudio.google.com/s/luyhD0TXrVs)
## [11][Can force stop phone calling in flutter?](https://www.reddit.com/r/androiddev/comments/j7ax95/can_force_stop_phone_calling_in_flutter/)
- url: https://www.reddit.com/r/androiddev/comments/j7ax95/can_force_stop_phone_calling_in_flutter/
---

## [12][Billing V3 problems. IAP being consumed or expiring](https://www.reddit.com/r/androiddev/comments/j77z8v/billing_v3_problems_iap_being_consumed_or_expiring/)
- url: https://www.reddit.com/r/androiddev/comments/j77z8v/billing_v3_problems_iap_being_consumed_or_expiring/
---
Been running v1 billing for years, no trouble, forced to move to V3 and having some big problems. My beta build appears to be expiring IAP. They work for a while while then disappear seemingly for good, like they are a consumable.  I haven't changed anything store side and my purchasing code is mostly the same, aside for the minor changes to support the updated library (mostly changes to BillingResult).

The purchase workflow appears to be working, after buying, I get my feature, but soon afterwards it's gone again.

Got some Beta testers that have lost their purchases, and not overly happy..  I could give them codes to get their purchases back, bit they are going to disappear again until I fix this.

What have I got wrong?

Thanks.
