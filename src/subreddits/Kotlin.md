# Kotlin
## [1][An Ergonomics Review of Using Kotlin from Swift](https://www.reddit.com/r/Kotlin/comments/ib73br/an_ergonomics_review_of_using_kotlin_from_swift/)
- url: https://benasher.co/kotlin-ios-ergonomics/
---

## [2][kotlinx.serialization Release 1.0.0-RC](https://www.reddit.com/r/Kotlin/comments/ibdz0r/kotlinxserialization_release_100rc/)
- url: https://github.com/Kotlin/kotlinx.serialization/releases/tag/1.0.0-RC
---

## [3][Best beginner book to build android apps using Kotlin?](https://www.reddit.com/r/Kotlin/comments/ib3gam/best_beginner_book_to_build_android_apps_using/)
- url: https://www.reddit.com/r/Kotlin/comments/ib3gam/best_beginner_book_to_build_android_apps_using/
---
Hi, I would like to purchase a book from Amazon that teaches me Kotlin and building android apps for beginners. There's alot of options on Amazon, but not sure which one to buy.
## [4][Why do we have emptyCollection() but no emptyMutableCollection()?](https://www.reddit.com/r/Kotlin/comments/iaw6h2/why_do_we_have_emptycollection_but_no/)
- url: https://www.reddit.com/r/Kotlin/comments/iaw6h2/why_do_we_have_emptycollection_but_no/
---
Kotlin provides  `emptyList()` but not `emptyMutableList()` . Is there a reason behind this?



`val myList: MutableList&lt;String&gt; = mutableListOf()`

vs.

`val myList: List&lt;String&gt; = emptyList()`
## [5][Running Kotlin Microservice on Google Kubernetes Engine](https://www.reddit.com/r/Kotlin/comments/iavp21/running_kotlin_microservice_on_google_kubernetes/)
- url: https://piotrminkowski.com/2020/04/03/running-kotlin-microservice-on-google-kubernetes-engine/
---

## [6][Creating a custom serializer that encodes an object as a ByteArray](https://www.reddit.com/r/Kotlin/comments/iauhhn/creating_a_custom_serializer_that_encodes_an/)
- url: https://www.reddit.com/r/Kotlin/comments/iauhhn/creating_a_custom_serializer_that_encodes_an/
---
I have a class called `ByteArraySegment`, which is a simple immutable wrapper around a `ByteArray`.  I want to make it `@Serializable` and for its serialized form to be a plain `ByteArray`.

I've been trying to follow [the docs](https://github.com/Kotlin/kotlinx.serialization/blob/master/docs/custom_serializers.md#representing-classes-as-a-single-value) to do this, and here is my attempt so far:

    @Serializer(forClass = ByteArraySegment::class)
    companion object : KSerializer&lt;ByteArraySegment&gt; {
        override val descriptor: SerialDescriptor
            get() = ByteArraySerializer().descriptor

        override fun serialize(encoder: Encoder, value: ByteArraySegment) {
            encoder.encode(ByteArraySerializer(), value.asArray)
        }

        override fun deserialize(decoder: Decoder): ByteArraySegment {
            return ByteArraySegment(decoder.decode(ByteArraySerializer()))
        }
    }

Unfortunately I get the following error when I try to serialize a `ByteArraySegment`:

    java.lang.ClassCastException: class locutus.tools.ByteArraySegment cannot be cast to 
    class [B (locutus.tools.ByteArraySegment is in unnamed module of loader 'app'; [B is in 
    module java.base of loader 'bootstrap')`

Would appreciate it if anyone has any ideas on what I'm doing wrong, as I've been blocked on this problem for a while.
## [7][AndroidBites | Init Blocks will never haunt you again |](https://www.reddit.com/r/Kotlin/comments/ib648x/androidbites_init_blocks_will_never_haunt_you/)
- url: https://chetangupta.net/init-blocks
---

## [8][Kotlin Multiplatform project with multi on demand modules](https://www.reddit.com/r/Kotlin/comments/iahz5r/kotlin_multiplatform_project_with_multi_on_demand/)
- url: https://www.reddit.com/r/Kotlin/comments/iahz5r/kotlin_multiplatform_project_with_multi_on_demand/
---
[https://stackoverflow.com/questions/63429543/kotlin-multiplatform-project-with-multi-on-demand-modules](https://stackoverflow.com/questions/63429543/kotlin-multiplatform-project-with-multi-on-demand-modules)
## [9][Waiting for Kotlin 1.4.0 Official Announcement be like](https://www.reddit.com/r/Kotlin/comments/i9wygq/waiting_for_kotlin_140_official_announcement_be/)
- url: https://i.imgur.com/MjWpJ6M.gif
---

## [10][LifecycleService as a Foreground service || Kotlin || LiveData || UI state management through service](https://www.reddit.com/r/Kotlin/comments/iacyhf/lifecycleservice_as_a_foreground_service_kotlin/)
- url: https://youtu.be/1euA5LeyLIg
---

