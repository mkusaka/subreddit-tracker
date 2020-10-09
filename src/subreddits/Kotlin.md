# Kotlin
## [1][Released kotlinx.serialization first public stable release](https://www.reddit.com/r/Kotlin/comments/j7f6dk/released_kotlinxserialization_first_public_stable/)
- url: https://github.com/Kotlin/kotlinx.serialization/releases/tag/v1.0.0
---

## [2][Seek help configuring kotlin + gradle](https://www.reddit.com/r/Kotlin/comments/j7ykhw/seek_help_configuring_kotlin_gradle/)
- url: https://www.reddit.com/r/Kotlin/comments/j7ykhw/seek_help_configuring_kotlin_gradle/
---
Build is failing because local jar dependency is not being used (see pic).  
Please see repo with code -   
[https://github.com/sanskrit-coders/jyotisha-kotlin/blob/master/build.gradle](https://github.com/sanskrit-coders/jyotisha-kotlin/blob/master/build.gradle)  


https://preview.redd.it/mlxvj8ili2s51.png?width=1147&amp;format=png&amp;auto=webp&amp;s=c699643a5178e88636516b2c23678505dfb5b17d
## [3][Convert JSON to Kotlin classes.](https://www.reddit.com/r/Kotlin/comments/j7xz7x/convert_json_to_kotlin_classes/)
- url: https://www.reddit.com/r/Kotlin/comments/j7xz7x/convert_json_to_kotlin_classes/
---
Hi all,

I have a chunk of json from here (https://world.openfoodfacts.org/api/v0/product/5010358111810.json), that I need to be able to parse.

I wrote a quick prototype in C# where I simply pasted the JSON in to Visual Studio and it converts the json into C# classes.

I realise that you can't do that in Android Studio and I tried an online json--&gt;Kotlin class converter tool, which appeared to do the job...until I imported them into Android Studio. 

It didn't like some of the names it created, so I set about changing them manually, then realised that when it comes to deserialising it, the deserializer won't know what the names are to match.

How do you all go about creating the classes for a JSON string this size?

Failing that, is it possible to convert it to a generic type and iterate around the result?
## [4][kotlinx.serialization 1.0 released – Kotlin Blog](https://www.reddit.com/r/Kotlin/comments/j7h8cp/kotlinxserialization_10_released_kotlin_blog/)
- url: https://blog.jetbrains.com/kotlin/2020/10/kotlinx-serialization-1-0-released/
---

## [5][A beginner question](https://www.reddit.com/r/Kotlin/comments/j7s3n0/a_beginner_question/)
- url: https://www.reddit.com/r/Kotlin/comments/j7s3n0/a_beginner_question/
---
I have dragged a number of buttons into a view. I have used the attributes panel to give these buttons properties. My question is under these circumstances is it required that I create a variable for these buttons in the .kt  file. I am thinking that under these circumstances the IDE has already accomplished this. Is this the case?
Thanks
## [6][Kotlin Multiplatform Mobile (KMM) goes Alpha](https://www.reddit.com/r/Kotlin/comments/j79eyu/kotlin_multiplatform_mobile_kmm_goes_alpha/)
- url: https://thetechjourno.blogspot.com/2020/10/kotlin-multiplatform-mobile-kmm-goes.html
---

## [7][Interested in writing Kotlin articles, get paid to write tutorials, follow the link.](https://www.reddit.com/r/Kotlin/comments/j7b5lp/interested_in_writing_kotlin_articles_get_paid_to/)
- url: https://developers.decoded.africa/get-paid-to-write-tutorials/
---

## [8][Calling REST API from Kotlin?](https://www.reddit.com/r/Kotlin/comments/j7binh/calling_rest_api_from_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/j7binh/calling_rest_api_from_kotlin/
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
## [9][Help! Can't even get a "toast" pop up working.](https://www.reddit.com/r/Kotlin/comments/j79qd5/help_cant_even_get_a_toast_pop_up_working/)
- url: https://www.reddit.com/r/Kotlin/comments/j79qd5/help_cant_even_get_a_toast_pop_up_working/
---
Hi all,

I've been going at this for over an hour now. I'm just running through tutorials etc..., but nothing seems to be working at all.

I found this:

https://developer.android.com/guide/topics/ui/notifiers/toasts

And then tried this:

    view.findViewById&lt;Button&gt;(R.id.button_toast).setOnClickListener {
            val text = "Hello toast!"
            val duration = Toast.LENGTH_SHORT
            val toast = Toast.makeText(this,text, duration)
            toast.show()
        }


Which then moaned:

    one of the following functions can be called with the arguments supplied: 
    public open fun makeText(p0: Context!, p1: CharSequence!, p2: Int): Toast! defined in 
        android.widget.Toast
    public open fun makeText(p0: Context!, p1: Int, p2: Int): Toast! defined in android.widget.Toast


So, I tried this:

    fun Context.toast(message: CharSequence) =
            Toast.makeText(this, "hello", Toast.LENGTH_SHORT).show()

Which doesn't moan about anything, but just doesn't work. I can step over/into it and there's no error or anything, just no popup.

What do you have to do to get a popup working in Kotlin?

TIA.
## [10][I want to learn kotlin, do you have any advice?](https://www.reddit.com/r/Kotlin/comments/j7ef0a/i_want_to_learn_kotlin_do_you_have_any_advice/)
- url: https://www.reddit.com/r/Kotlin/comments/j7ef0a/i_want_to_learn_kotlin_do_you_have_any_advice/
---
Hello, I want to learn kotlin. I can't use IntelliJ IDEA because I have 1GB RAM and Intel Pentium 2 T3200 laptop. I didn't know any languages ​​before, but so far I have learned the basic syntax. The Kotlin playground is insufficient, do you have any program suggestions for applying and improving what I have learned? Also do you have any advice for learning?
