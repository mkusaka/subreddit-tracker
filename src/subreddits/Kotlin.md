# Kotlin
## [1][AndroidBites | 5 Easy Steps to LRU Cache in Repository.](https://www.reddit.com/r/Kotlin/comments/io0ygt/androidbites_5_easy_steps_to_lru_cache_in/)
- url: https://chetangupta.net/cache-repository/
---

## [2][I get empty Aggregated Heart Rate Data from Google Fitness Api.](https://www.reddit.com/r/Kotlin/comments/io7fw2/i_get_empty_aggregated_heart_rate_data_from/)
- url: https://www.reddit.com/r/Kotlin/comments/io7fw2/i_get_empty_aggregated_heart_rate_data_from/
---
Hey guys, I am writing my thesis and I have this task to get aggregated Heart Rate data from my Google Fit Account and print them. My code works if I want to print heart rate data from an hour during the day. But I cannot for the love of god print the values of the aggregated heart rate for each day of the week. I am a novice as far as the Google Fit Api is concerned so bare with me. I will post my builder and data dumper. If you need any more code lmk.

**Builder Function:**

    private fun queryFitnessData2(): DataReadRequest { 
    val calendar = Calendar.getInstance(TimeZone.getTimeZone("UTC")) 
    val now = Date()
    calendar.time = now 
    val endTime = calendar.timeInMillis calendar.add(Calendar.WEEK_OF_YEAR, -1) 
    val startTime = calendar.timeInMillis
    Log.i(TAG, "Range Start: ${dateFormat.format(startTime)}") 
    Log.i(TAG, "Range End: ${dateFormat.format(endTime)}")
    
    return DataReadRequest.Builder()
    
                    .enableServerQueries()
                    .bucketByTime(1, TimeUnit.DAYS)
                    .setTimeRange(startTime, endTime, TimeUnit.MILLISECONDS)
                    .aggregate(DataType.TYPE_HEART_RATE_BPM, 
                    DataType.AGGREGATE_HEART_RATE_SUMMARY)
                    .build();
    }

**Data Printer:**

    private fun printData(dataReadResult: DataReadResponse) {
            // [START parse_read_data_result]
            // If the DataReadRequest object specified aggregated data, dataReadResult will be returned
            // as buckets containing DataSets, instead of just DataSets.
            if (dataReadResult.buckets.isNotEmpty()) {
                Log.i(TAG, "Number of returned buckets of DataSets is: " + dataReadResult.buckets.size)
                for (bucket in dataReadResult.buckets) {
                    val dataSets = bucket.dataSets
                    bucket.dataSets.forEach { dumpDataSet(it) }
                }
            } else if (dataReadResult.dataSets.isNotEmpty()) {
                Log.i(TAG, "Number of returned DataSets is: " + dataReadResult.dataSets.size)
                dataReadResult.dataSets.forEach { dumpDataSet(it) }
            }
            // [END parse_read_data_result]
        }

**Dataset Parser Function:**

&amp;#x200B;

    // [START parse_dataset]
    private fun dumpDataSet(dataSet: DataSet) { 
        Log.i(TAG, "Data returned for Data type: ${dataSet.dataType.name}") 
        val dateFormat: DateFormat = getTimeInstance()
        //INTERESTING PART 
        dataSet.dataType.fields.forEach { 
           Log.i(TAG, "\tField: ${it.name} Value:${dataSet.dataPoints.getValue(it)}") 
        } 
    }
    // [END parse_dataset]

In the interesting part section, I can log the Min, Max, Average sections but I fail to get the values from those fields. Any help would be greatly appreciated.
## [3][Because there is no Multiplatform Dependency Injection library, i've written a dependency injection library that can be used in Kotlin Native/JS/JVM](https://www.reddit.com/r/Kotlin/comments/inon8t/because_there_is_no_multiplatform_dependency/)
- url: https://www.reddit.com/r/Kotlin/comments/inon8t/because_there_is_no_multiplatform_dependency/
---
[https://github.com/corbella83/PopKorn](https://github.com/corbella83/PopKorn)
## [4][Is Kotlin good for games?](https://www.reddit.com/r/Kotlin/comments/injh8b/is_kotlin_good_for_games/)
- url: https://www.reddit.com/r/Kotlin/comments/injh8b/is_kotlin_good_for_games/
---
Hi

I've had a casual interest in game making my whole life and I want to try building mobile games. I don't have much experience in programming, just a little bit of python from 2 years back and some basic BATCH from a few years back.

Is Kotlin powerful enough to build small arcady games with?
## [5][What Are the Differences Between Kotlin &amp; Flutter?](https://www.reddit.com/r/Kotlin/comments/io49l4/what_are_the_differences_between_kotlin_flutter/)
- url: https://technostacks.com/blog/kotlin-vs-flutter
---

## [6][Kotlin vs Java: For Android Apps and Beyond in 2021...Any Views Guys?](https://www.reddit.com/r/Kotlin/comments/insuoq/kotlin_vs_java_for_android_apps_and_beyond_in/)
- url: https://eduwyre.com/article/kotlin-vs-java-selecting-the-better-one
---

## [7][What this syntax named in kotlin?](https://www.reddit.com/r/Kotlin/comments/inmwg1/what_this_syntax_named_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/inmwg1/what_this_syntax_named_in_kotlin/
---
foo { .... }
## [8][Learning kotlin server side](https://www.reddit.com/r/Kotlin/comments/inh1eh/learning_kotlin_server_side/)
- url: https://www.reddit.com/r/Kotlin/comments/inh1eh/learning_kotlin_server_side/
---
Hi everyone,

I've had an ambition to learn to make apps. However with covid and everything the android dev market here in the Netherlands is saturated. So I thought I'd learn to code server side in kotlin since I've spent 100 hours in android kotlin tutorials. When the Android market grows again I'll at least have some experience with the language.

I don't know where to start though. Anyone have any ideas for a beginning developer?

Edit: wow, didn't expect this many comments. Thanks everyone!
## [9][I should have known better](https://www.reddit.com/r/Kotlin/comments/ingqgc/i_should_have_known_better/)
- url: https://www.reddit.com/r/Kotlin/comments/ingqgc/i_should_have_known_better/
---
What am I missing that would make my life easier?

So, I write C++ and occasionally some Java (a desktop Swing application) for a living. Been thinking about trying out Kotlin for a while, finally found a reason: Compile &amp; patch [https://github.com/manuelsc/HDDFs-MPG-Dumper](https://github.com/manuelsc/HDDFs-MPG-Dumper)

I got to say, Kotlin kicked my ass today. It wasn't even the language that did it for me. It's Maven and the build tools. The IDEs did not help.

So I tried Eclipse, loaded in the Kotlin plugin, made a Kotlin project, set it as a Maven thingy and BAM, Kotlin dependencies disappeared. WTF. Mind you I've never used Maven (I'm quite happy with Ant, and I loath Make, MsBuild is pretty bad). Googled for help, found a terrible tutorial. Eclipse will crap out Ant scripts on demand, it should be able to do the same for Maven. I decided that maybe Eclipse's shitty Kotlin-Maven support was a sign I should try something else. 

So I looked at the .gitignore file, it mentioned IDEA project files, so I'm like "welp it's about time I tried out Intellij again" I've used Eclipse for all of my Java development (it's the only thing Eclipse is good at). The nice thing about Eclipse is you can mostly ignore dependency problem, just put the jars in a folder and pull them into the project dependencies. What you see is what you get. Unless I completely missed it Intellij, completely lacks anything remotely useful to help new users out with Maven. I've been thrown into the deep end of Maven. Finally dropped the dependencies in and... The Maven Central site Intellij tried to pulled the dependencies from didn't have it. I couldn't even figure out what to google for that one.

So I went back to Eclipse, made a new project, not a Maven Project, and downloaded the jars from Maven Central, added them as dependencies to the Java project, Eclipse was happy. No errors, no unhappy squiggles. And then I tried to launch it. org.jetbrains.kotlin.ui.console is like "ERROR: Unresolved reference: java (3, 8)" along with pages more of errors like it can't find jack shit. Regular console is like "Error: Could not find or load main class MainKt"

Anyway, suggestions?
## [10][Where to ACTUALLY learn kotlin?](https://www.reddit.com/r/Kotlin/comments/innjib/where_to_actually_learn_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/innjib/where_to_actually_learn_kotlin/
---
Hey i wanna learn kotlin i tried udemy but non of the courses are good enough like they don’t actually teach you why they use this or that plus i want a good course that the lecturer’s first language is English! Any help ?
