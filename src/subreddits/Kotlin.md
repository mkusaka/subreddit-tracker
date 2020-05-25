# Kotlin
## [1][viewModels&lt;&gt;() vs activityViewModels&lt;&gt;() vs createViewModelLazy()](https://www.reddit.com/r/Kotlin/comments/gq5wjk/viewmodels_vs_activityviewmodels_vs/)
- url: https://www.reddit.com/r/Kotlin/comments/gq5wjk/viewmodels_vs_activityviewmodels_vs/
---
Recently I've been going thorough the **KTX** libraries and I found several ways to instantiate VM using ktx. 

I tried following up on the differences between by **viewModels&lt;&gt;()**, **activityViewModels&lt;&gt;()** and **createViewModelLazy()** but couldn't figure out much.  

While all this time, I've been using **KOIN viewmodel artifact** for initializing my VMs, I am just keen to know the differences between the above mentioned.
## [2][Need help with this confusion in data binding](https://www.reddit.com/r/Kotlin/comments/gq6r2e/need_help_with_this_confusion_in_data_binding/)
- url: https://www.reddit.com/r/Kotlin/comments/gq6r2e/need_help_with_this_confusion_in_data_binding/
---
So i am still learning about Android app development with Kotlin using a Udacity course.

I was learning about data binding in the course and they used the following code;

&amp;#x200B;

My name.kt

    package com.example.aboutme
    data class MyName(var name: String = "",var nickname: String = "")

&amp;#x200B;

MainActivity.kt

    import androidx.databinding.DataBindingUtil
    import com.rackaapps.aboutme.databinding.ActivityMainBinding
    
    class MainActivity : AppCompatActivity() {
    
        private lateinit var binding: ActivityMainBinding
        private val myName = MyName("Racka Legit")

&amp;#x200B;

activity\_main.xml (for a text view that's supposed to show the name &amp; one for nickname)

    &lt;TextView
                android:id="@+id/name_text"
                style="@style/NameStyle"
                android:layout_width="match_parent"
                android:layout_height="wrap_content"
                android:text="@={myName.name}"
                android:textAlignment="center" /&gt;
    
    &lt;TextView
                android:id="@+id/nickname_text"
                style="@style/NameStyle"
                android:layout_width="match_parent"
                android:layout_height="wrap_content"
                android:textAlignment="center"
                android:visibility="gone"
                android:text="@={myName.nickname}"/&gt;

&amp;#x200B;

Now with this code i get an error saying `"cannot find symbol class ActivityMainBindingImpl"`

And its telling me to `"import com.example.aboutme.databinding.ActivityMainBindingImpl"` which gives an error too.

**Note that,**

    dataBinding {enabled = true}

was enabled too.

&amp;#x200B;

**After some Googling i found out that i should use;**

    android:text="@{myName.name}"

&amp;

    android:text="@{myName.nickname}"

**Instead of;**

    android:text="@={myName.name}"

&amp;

    android:text="@={myName.nickname}"

and it worked like charm!

&amp;#x200B;

Now i want to know what difference does the equal (=) make in data binding in the context of the code above?

And why did the code used in the course work for them but didn't work for me?

I didn't find any info about the use of this yet.
## [3][Annotation processor Extensions](https://www.reddit.com/r/Kotlin/comments/gppuj9/annotation_processor_extensions/)
- url: https://www.reddit.com/r/Kotlin/comments/gppuj9/annotation_processor_extensions/
---
I created a repo with the intention of gathering useful extensions/utils for anyone working with Kotlin annotation processing and code generation. Have a look.

https://github.com/iFanie/KTAP

Also, what extension/utils do you find useful for Kotlin annotation processor projects?
## [4][Understand Kotlin Collection Function Past Tense](https://www.reddit.com/r/Kotlin/comments/gpk2w7/understand_kotlin_collection_function_past_tense/)
- url: https://medium.com/@elye.project/understand-kotlin-collection-function-past-tense-59f592af9436?source=friends_link&amp;sk=e22ccd272ebd28ef6f417e7b455b5b4f
---

## [5][How to secure Ktor web app with Keycloak Jetty 9.x Adapters?](https://www.reddit.com/r/Kotlin/comments/gpnxq8/how_to_secure_ktor_web_app_with_keycloak_jetty_9x/)
- url: https://www.reddit.com/r/Kotlin/comments/gpnxq8/how_to_secure_ktor_web_app_with_keycloak_jetty_9x/
---
Hi all 

I would like to secure the Ktor webapp, that is run on top of Jetty server with [Keycloak Jetty 9.x Adapters][1]. 

Ktor provides a [hook][2] for Jetty server initialization and maybe it is the right place in integrate the Keycloak Jetty 9.x Adapters.  

How to integrate Keycloak Jetty 9.x Adapters into Ktor app?  

Thanks


  [1]: https://www.keycloak.org/docs/latest/securing_apps/index.html#_jetty9_adapter
  [2]: https://ktor.io/servers/configuration.html#jetty
## [6][I want to search in multiple columns at the same time](https://www.reddit.com/r/Kotlin/comments/gpqxs5/i_want_to_search_in_multiple_columns_at_the_same/)
- url: https://www.reddit.com/r/Kotlin/comments/gpqxs5/i_want_to_search_in_multiple_columns_at_the_same/
---
So I have an SQLite DB with a table that contains the following columns: ID, Title, Description, Date. My current query only searches in the Title column, but I want it to also look for the matching string in the Description and Date columns and return the result.

I have one more question, what is the difference between that type of query and those like "SELECT \* FROM..."?

Here is my query fun [https://pastebin.com/mqbJhkh2](https://pastebin.com/mqbJhkh2)

Edit: I found a solution. I created another function just like the other one and passed the search string as 3 parameters and used this query `"Title like ? OR Description like ? OR Date like ? ",selectionArgs"`.
## [7][Harmony - A process-safe SharedPreference library](https://www.reddit.com/r/Kotlin/comments/gpbmci/harmony_a_processsafe_sharedpreference_library/)
- url: /r/android_devs/comments/gpbkyp/harmony_a_processsafe_sharedpreference_library/
---

## [8][Kotlin native performance compared to JVM - Opinions wanted](https://www.reddit.com/r/Kotlin/comments/gpdnt7/kotlin_native_performance_compared_to_jvm/)
- url: https://www.reddit.com/r/Kotlin/comments/gpdnt7/kotlin_native_performance_compared_to_jvm/
---
So my experience with Kotlin native has been less than ideal, bit of a rant (sorry!)

Essentially, after 3 years of Kotlin, went into the native field, thinking i would get better performance or memory usage for large scale applications.  


Although the latter is true, especially with millions of objects, the same cannot be said about the former.

First example. Java's File class vs native code using C functions such as fread(). I did some benchmarks (Loading a 55MB text file and load all the lines into an array. JVM does this in 10ms, where native code does it in 2500ms. How is it that much slower!?

Second example, loading loads of classes. JVM does the whole thing about 1000x faster, when converting lines in arrays to objects using a parser. Admittedly, under memory pressure, Kotlin native does perform more consistently, where JVM will pause to do a GC sweep. But still, total performance difference is about 1000x!

Has anyone else had similar experiences? - Or do you reckon i am doing something wrong here?
## [9][Trying to understand something about Kotlin's Collections interface and Java](https://www.reddit.com/r/Kotlin/comments/gp1hil/trying_to_understand_something_about_kotlins/)
- url: https://www.reddit.com/r/Kotlin/comments/gp1hil/trying_to_understand_something_about_kotlins/
---
I'm using Data Binding on Android and I'm trying to understand why I have to import `java.util.List` into my data binding layouts even tho the rest of the code is using Kotlin:

 [https://imgur.com/a/jFiIfy9](https://imgur.com/a/jFiIfy9) 

I tried replacing this for `kotlin.collections.List` but the code doesn't compile. Online I read something about that the Kotlin collection interfaces "don't actually exist". Can someone shed some light on this?
## [10][Spring Framework returns 404 for some URLs even when there are controllers for them.](https://www.reddit.com/r/Kotlin/comments/gpgcsh/spring_framework_returns_404_for_some_urls_even/)
- url: https://www.reddit.com/r/Kotlin/comments/gpgcsh/spring_framework_returns_404_for_some_urls_even/
---
I two controllers for my site written in Spring Boot and Kotlin, `IndexController` and `AboutController`. Both with a single function with a `@GetMapping` annotation to `/` and `/about`, respectively. However, the index page works find but the about page returns a 404 error. Even when I move the about page controller method to the IndexController, it still doesn't work.

In another test project, I initially couldn't even get the index page to route correctly, but that problem randomly fixed itself for some reason.

What could be going on, and how do I fix it?
