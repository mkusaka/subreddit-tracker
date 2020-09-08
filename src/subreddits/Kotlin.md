# Kotlin
## [1][A Cake for Kotlin: Exploring framework-less dependency injection for Kotlin](https://www.reddit.com/r/Kotlin/comments/iot8oo/a_cake_for_kotlin_exploring_frameworkless/)
- url: https://knowitlabs.no/a-cake-for-kotlin-exploring-framework-less-dependency-injection-for-kotlin-3f93b076aa8c
---

## [2][http4k blog: A retrospective on http4k v3](https://www.reddit.com/r/Kotlin/comments/iobqsq/http4k_blog_a_retrospective_on_http4k_v3/)
- url: https://www.http4k.org/blog/retrospective_v3/
---

## [3][How to create a List of Strings out of a List&lt;CharArray&gt;?](https://www.reddit.com/r/Kotlin/comments/iogw7i/how_to_create_a_list_of_strings_out_of_a/)
- url: https://www.reddit.com/r/Kotlin/comments/iogw7i/how_to_create_a_list_of_strings_out_of_a/
---
Is it possible to create a list of strings given a list of CharArray, in which each string is formed  by the elements that occupy the index i in each CharArray in the list? Like this:
List&lt;CharArray&gt; = {[A,B,C],[D,E,F]} then the resulting List&lt;String&gt; = {“AD”, “BE”, “CF”}. If it is possible to do this using as many collections functions as possible. I wish to understand them as much as I can, here is where Kotlin shines (besides its concise and pretty syntax)
## [4][Is there an unsigned byte literal?](https://www.reddit.com/r/Kotlin/comments/iodvoe/is_there_an_unsigned_byte_literal/)
- url: https://www.reddit.com/r/Kotlin/comments/iodvoe/is_there_an_unsigned_byte_literal/
---
I have a function that returns a `UByte`. I cannot use an equality operator because I do not know how to create a `UByte` literal:

```kotlin
fun foo(): UByte = 0u

fun bar(): Boolean = foo() == 0u
```
(This reports that the `==` operator cannot compare a `UByte` and a `UInt`)

My best workaround is

```kotlin
fun foo(): UByte = 0u

fun bar(): Boolean = foo() == run {
    val x: UByte = 0u
    x
}
```

but this looks odd.

With `(0u as UByte)` IDEA reports that the cast will never succeed.

If there isn't a UByte literal, is there a cleaner way to do this?
## [5][Hibernate cascading with no cascade types set (hibernate+kotlin)](https://www.reddit.com/r/Kotlin/comments/io9sxn/hibernate_cascading_with_no_cascade_types_set/)
- url: https://www.reddit.com/r/Kotlin/comments/io9sxn/hibernate_cascading_with_no_cascade_types_set/
---
Hey guys, I'm having an issue where hibernate is cascading the `merge` operation to child entities when I don't want it to. I have my entities pasted at the bottom of this post. 

I'm trying to merge the `Munch` entity without merging any `Swipes` since they're handled in a different part of the application. My understanding is that hibernate by default should `cascade` none of the DB operations for a `@OneToMany` collection or a `@ManyToOne` object unless `CascadeTypes` are explicitly specified. 


Given the entities at the bottom of the post, when I add a `Swipe` to `munch.swipes` and run the following code, the munch is updated if any of its fields have changed and the added swipe is merged into the db:
```
    fun mergeMunch(
        munch: Munch
    ) = databaseExecutor.executeAndRollbackOnFailure { entityManager -&gt;
        entityManager.merge(munch)
        entityManager.transaction.commit()
    }
```

If anyone could shed some light on either what I'm misunderstanding or misconfiguring it would be much appreciated. 


The `executeAndRollbackOnFailure()` function just in case its useful:
```
    fun &lt;T&gt; executeAndRollbackOnFailure(
        task: (EntityManager) -&gt; T
    ): T {
        val em = emf.createEntityManager()
        return try {
            em.transaction.begin()
            task.invoke(em)
        } catch (e: Exception) {
            em.transaction.rollback()
            throw e
        } finally {
            em.close()
        }
    }
```

Here are my entities:

`Munch`
```

@Entity
data class Munch(
    @Column
    val name: String,
    @OneToMany(
        fetch = FetchType.LAZY,
        mappedBy = "munch",
    )
    val swipes: MutableList&lt;Swipe&gt; = mutableListOf(),
) {
    @Id
    @GenericGenerator(name = "generator", strategy = "uuid")
    @GeneratedValue(generator = "generator")
    lateinit var munchId: String

    fun addSwipe(swipe: Swipe) {
        swipes.add(swipe)
        swipe.munch = this
    }
}
```

`Swipe`

```
@Entity
data class Swipe(
    @EmbeddedId
    val swipeIdKey: SwipeIdKey,
    @Column(nullable = true)
    val liked: Boolean,
) : Serializable {
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "munchId")
    @MapsId("munchId")
    lateinit var munch: Munch

    @Transient
    var updated = false
```

`SwipeIdKey`

```
@Embeddable
class SwipeIdKey : Serializable {

    @Column(nullable = false)
    lateinit var restaurantId: String

    @Column(nullable = true)
    lateinit var userId: String

    @Column(nullable = true)
    var munchId: String? = null
}
```
## [6][Best Programming Languages for Web Development](https://www.reddit.com/r/Kotlin/comments/ioppbr/best_programming_languages_for_web_development/)
- url: https://tekkiwebsolutions.com/blog/best-programming-languages-for-web-development/
---

## [7][Where to publish packages?](https://www.reddit.com/r/Kotlin/comments/ioiccq/where_to_publish_packages/)
- url: https://www.reddit.com/r/Kotlin/comments/ioiccq/where_to_publish_packages/
---
Jcenter? Maven central? Someplace else?

With Python/.NET - there's one obvious place to publish. What's the recommendation for JVM land?
## [8][Setting and getting an explanation](https://www.reddit.com/r/Kotlin/comments/iohzmx/setting_and_getting_an_explanation/)
- url: https://www.reddit.com/r/Kotlin/comments/iohzmx/setting_and_getting_an_explanation/
---
Hey guys,

I recently started to learn Kotlin, and I'm having some troubles understanding the usage and the meaning of get and set methods in classes creation. I do not have any Java background (this is killing me, tbh) and I recently found them while starting to learn OOP and classes management. 

The online class I'm following is trying to explain me the importance of get and set, especially while checking for correct parameter inputs in classes but I can't seem to get the point.

Why using set and get methods if I can check for a parameter value directly inside the init method? Is there something I'm missing? 

I also looked on the documentation about them, but it shows how they works, not for what or when they are used.

Can someone better explain to me the meaning and the usage of these two methods? A brief explanations would be great if it is a huge thematic to talk about.

Thank you a lot! :D
## [9][Which "createOrAdd" style would you prefer? Do you know a better one?](https://www.reddit.com/r/Kotlin/comments/ioffd3/which_createoradd_style_would_you_prefer_do_you/)
- url: https://www.reddit.com/r/Kotlin/comments/ioffd3/which_createoradd_style_would_you_prefer_do_you/
---
Hey fellow Kotliners!

I'd like to hear out your opinions. I'm basically implementing a good ol' add peer to room logic. If the room doesn't exist create it on the fly, otherwise just add it to the room.

Which ones seems more readable to you? Which is the "more kotlin" way? Is there a better one? I'm using Spring Data JPA btw. Thanks in advance!

&amp;#x200B;

    roomRepository.findByIdOrNull(roomId)?.apply {
        addPeerToExistingRoom(this, peerId)
    } ?: addPeerToNewRoom(roomId, peerId)

&amp;#x200B;

    val room: Room? = roomRepository.findByIdOrNull(roomId)
    val responseType = if (room == null) {
        addPeerToNewRoom(roomId, peerId)
    } else {
        addPeerToExistingRoom(room, peerId)
    }
## [10][I get empty Aggregated Heart Rate Data from Google Fitness Api.](https://www.reddit.com/r/Kotlin/comments/io7fw2/i_get_empty_aggregated_heart_rate_data_from/)
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
