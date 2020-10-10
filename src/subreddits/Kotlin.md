# Kotlin
## [1][How to serialize an array of different types of value with `kotlinx.serialization`.](https://www.reddit.com/r/Kotlin/comments/j8gjt8/how_to_serialize_an_array_of_different_types_of/)
- url: https://www.reddit.com/r/Kotlin/comments/j8gjt8/how_to_serialize_an_array_of_different_types_of/
---
I try the following code.

\`\`\` kotlin

import kotlinx.serialization.json.Json

Json.encodeToString(arrayOf("gid", arrayOf("status", "c", "d"), 3))

\`\`\`

&amp;#x200B;

It gives the error message.

&amp;#x200B;

\`\`\`

Exception in thread "main" kotlinx.serialization.SerializationException: Class 'String' is not registered for polymorphic serialization in the scope of 'Serializable'.

Mark the base class as 'sealed' or register the serializer explicitly.

\`\`\`
## [2][New Release Cadence for Kotlin and the IntelliJ Kotlin Plugin – Kotlin Blog](https://www.reddit.com/r/Kotlin/comments/j81868/new_release_cadence_for_kotlin_and_the_intellij/)
- url: https://blog.jetbrains.com/kotlin/2020/10/new-release-cadence-for-kotlin-and-the-intellij-kotlin-plugin/
---

## [3][What is most advanced web framework/libraries in kotlin?](https://www.reddit.com/r/Kotlin/comments/j8gy4u/what_is_most_advanced_web_frameworklibraries_in/)
- url: https://www.reddit.com/r/Kotlin/comments/j8gy4u/what_is_most_advanced_web_frameworklibraries_in/
---
I am looking for something that has a lot of things working out-of-the-box. Something like symphony in PHP, where you have many ready to work solutions and you just adjust them. Is there something like this in kotlin? What is important for me is quickly building front-end and back-end so it will be more like building blocks. Spring is too complex, i want something easy. Ktor is very good but doesnt have solutions out of the box, you must create everything by yourself.
## [4][How to pass data between MainActivity and a Fragment?](https://www.reddit.com/r/Kotlin/comments/j8iczk/how_to_pass_data_between_mainactivity_and_a/)
- url: https://www.reddit.com/r/Kotlin/comments/j8iczk/how_to_pass_data_between_mainactivity_and_a/
---
Hi.

I've created a barcode scanner that once scanned writes the barcode to a string variable.

I now want access to that variable back on the MainActivity.

I can't see how to do it at all. 

I've followed an example online (https://johncodeos.com/how-to-pass-data-between-fragments-in-android-using-kotlin/), but as I'm finding quite often with these tutorials is that they are out of date, or simply don't work.
This one for instance, highlights things like "content_id" in red and won't compile.

So, before I go down yet another rabbit hole...is there a simple way of passing say String data back to the MainActivity from a Fragment?


I have this:

    val thread = Thread(Runnable {
                    try {
                        var json = URL("https://world.openfoodfacts.org/api/v0/product/" + result.contents.toString() + ".json").readText()
                        val mapper = jacksonObjectMapper()
                        mapper.disable(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES)
                        var products: Json4Kotlin_Base = mapper.readValue(json)
                        var product_name = products.product!!.product_name
                        var calories = products.product!!.nutriments?.energy_value
                        var barcode = result.contents.toString()
                        Thread(Runnable {
                            this@MainActivity.runOnUiThread(java.lang.Runnable {
                                this.text_food.text = product_name
                                this.text_calories.text = calories.toString() + " calories per 100g"
                                this.button_add_food.visibility = View.VISIBLE
                                this.text_barcode.text = barcode
                                
       // need to put variables here: barcode/calories/percent
 
                            })
                        }).start()

                    } catch (e: Exception) {
                        e.printStackTrace()
                    }
                })
                thread.start()

...in MainActivity.kt

A bit later - after some user input on HomeFragment, the user clicks a button and, I have:


            button_add_food.setOnClickListener {
                val thread = Thread(Runnable {
                    try {

                // and access them here:
                        var json = URL("http://www.myurl.co.uk:5001/api/food/1/$barcode/$calories/$percent").readText()

     ........
## [5][MDC Weekly Blend — Edition #9 - with loads of Android and Kotlin curated articles.](https://www.reddit.com/r/Kotlin/comments/j8icak/mdc_weekly_blend_edition_9_with_loads_of_android/)
- url: https://mobiledeveloperscafe.substack.com/p/mdc-weekly-blend-edition-9
---

## [6][SpringBoot alternative for Identity Server 4](https://www.reddit.com/r/Kotlin/comments/j8hwla/springboot_alternative_for_identity_server_4/)
- url: https://www.reddit.com/r/Kotlin/comments/j8hwla/springboot_alternative_for_identity_server_4/
---
Hi, is there something like Asp.Net Identity Server 4 for SpringBoot?
## [7][100-off-java-programming-complete-beginner-to-advanced-4/](https://www.reddit.com/r/Kotlin/comments/j8fk3e/100offjavaprogrammingcompletebeginnertoadvanced4/)
- url: https://unitedaca.com/100-off-java-programming-complete-beginner-to-advanced-4/
---

## [8][Seek help configuring kotlin + gradle](https://www.reddit.com/r/Kotlin/comments/j7ykhw/seek_help_configuring_kotlin_gradle/)
- url: https://www.reddit.com/r/Kotlin/comments/j7ykhw/seek_help_configuring_kotlin_gradle/
---
Build is failing because local jar dependency is not being used (see pic).  
Please see repo with code -   
[https://github.com/sanskrit-coders/jyotisha-kotlin/blob/master/build.gradle](https://github.com/sanskrit-coders/jyotisha-kotlin/blob/master/build.gradle)  


https://preview.redd.it/mlxvj8ili2s51.png?width=1147&amp;format=png&amp;auto=webp&amp;s=c699643a5178e88636516b2c23678505dfb5b17d
## [9][Convert JSON to Kotlin classes.](https://www.reddit.com/r/Kotlin/comments/j7xz7x/convert_json_to_kotlin_classes/)
- url: https://www.reddit.com/r/Kotlin/comments/j7xz7x/convert_json_to_kotlin_classes/
---
Hi all,

I have a chunk of json from here (https://world.openfoodfacts.org/api/v0/product/5010358111810.json), that I need to be able to parse.

I wrote a quick prototype in C# where I simply pasted the JSON in to Visual Studio and it converts the json into C# classes.

I realise that you can't do that in Android Studio and I tried an online json--&gt;Kotlin class converter tool, which appeared to do the job...until I imported them into Android Studio. 

It didn't like some of the names it created, so I set about changing them manually, then realised that when it comes to deserialising it, the deserializer won't know what the names are to match.

How do you all go about creating the classes for a JSON string this size?

Failing that, is it possible to convert it to a generic type and iterate around the result?
## [10][Synchronising access to a shared resource for asynchronous calls to a Spring Webserver does not seem to work with @Synchronized.... can someone help out?](https://www.reddit.com/r/Kotlin/comments/j81kaw/synchronising_access_to_a_shared_resource_for/)
- url: https://stackoverflow.com/questions/64282162/synchronising-access-to-a-shared-resource-for-asynchronous-calls-to-a-spring-web
---

