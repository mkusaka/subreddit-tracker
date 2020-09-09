# Kotlin
## [1][Played around with WebGL and a bug gave me this.](https://www.reddit.com/r/Kotlin/comments/ipedbe/played_around_with_webgl_and_a_bug_gave_me_this/)
- url: https://i.redd.it/pfe9csik14m51.png
---

## [2][Simple &amp; effective G1 GC tuning tips](https://www.reddit.com/r/Kotlin/comments/ipawvp/simple_effective_g1_gc_tuning_tips/)
- url: https://blog.gceasy.io/2020/06/02/simple-effective-g1-gc-tuning-tips/
---

## [3][How we built Localazy CLI: Kotlin MPP and Github Actions](https://www.reddit.com/r/Kotlin/comments/ipce9o/how_we_built_localazy_cli_kotlin_mpp_and_github/)
- url: https://www.reddit.com/r/Kotlin/comments/ipce9o/how_we_built_localazy_cli_kotlin_mpp_and_github/
---
**How we built Localazy CLI, command-line friendly software localization tool for smart&amp;lazy developers.** 

We thought about the right solution that would enable other developers to use Localazy on all platforms. We wanted to keep the developer-friendly / developer-centric approach and CLI, aka the command line interface, was the logical answer. 

We wanted **Localazy CLI** to be supported on all major platforms. Ideally, we wanted no dependency on Java or Node as those usually mean a huge binary or giant virtual machine. Small and fast binary for a given operating system looked like a perfect solution.

Kotlin is already heavily used at Localazy, and it powers the whole backend. And so [*Kotlin MPP*](https://kotlinlang.org/docs/reference/mpp-intro.html) was a logical choice to look at for this task. It allowed us to write a single codebase for native Linux, macOS, Windows and Java in a programming language we know well.  


Almost all the code is shared between platforms and written in pure Kotlin. There were only two platform-dependent features - filesystem access and HTTP communication.

The filesystem access is quite simple, and we just create an actual implementation for Java (using java.io.File and streams), Linux/macOS (using POSIX) and Windows (slightly modified POSIX as there are small differences on Windows).

For HTTP communication, good old *HTTPUrlConnection* is used for Java and *ktor client* for Linux, macOS and Windows. The prototype of CLI was ready in a couple of days, and with a bit of fine-tuning, we were able to compile it to Java’s JAR and native binary for Linux, macOS and Windows.

&amp;#x200B;

&gt;Thanks to JetBrains for Kotlin, it was fun to write the CLI, and it works smoothly on all platforms.

**Everything seemed to be great, except it wasn’t…read more at the link below.**

[https://localazy.com/blog/localazy-cli-sofware-localization-tool-kotlin-github](https://localazy.com/blog/localazy-cli-sofware-localization-tool-kotlin-github)
## [4][A Cake for Kotlin: Exploring framework-less dependency injection for Kotlin](https://www.reddit.com/r/Kotlin/comments/iot8oo/a_cake_for_kotlin_exploring_frameworkless/)
- url: https://knowitlabs.no/a-cake-for-kotlin-exploring-framework-less-dependency-injection-for-kotlin-3f93b076aa8c
---

## [5][http4k blog: A retrospective on http4k v3](https://www.reddit.com/r/Kotlin/comments/iobqsq/http4k_blog_a_retrospective_on_http4k_v3/)
- url: https://www.http4k.org/blog/retrospective_v3/
---

## [6][select image from gallery for button and display guys please help](https://www.reddit.com/r/Kotlin/comments/iowdnw/select_image_from_gallery_for_button_and_display/)
- url: https://stackoverflow.com/questions/63778471/select-image-from-gallery-for-button-and-display-kotlin-android-studio
---

## [7][How to create a List of Strings out of a List&lt;CharArray&gt;?](https://www.reddit.com/r/Kotlin/comments/iogw7i/how_to_create_a_list_of_strings_out_of_a/)
- url: https://www.reddit.com/r/Kotlin/comments/iogw7i/how_to_create_a_list_of_strings_out_of_a/
---
Is it possible to create a list of strings given a list of CharArray, in which each string is formed  by the elements that occupy the index i in each CharArray in the list? Like this:
List&lt;CharArray&gt; = {[A,B,C],[D,E,F]} then the resulting List&lt;String&gt; = {“AD”, “BE”, “CF”}. If it is possible to do this using as many collections functions as possible. I wish to understand them as much as I can, here is where Kotlin shines (besides its concise and pretty syntax)
## [8][Is there an unsigned byte literal?](https://www.reddit.com/r/Kotlin/comments/iodvoe/is_there_an_unsigned_byte_literal/)
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
## [9][Hibernate cascading with no cascade types set (hibernate+kotlin)](https://www.reddit.com/r/Kotlin/comments/io9sxn/hibernate_cascading_with_no_cascade_types_set/)
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
## [10][Setting and getting an explanation](https://www.reddit.com/r/Kotlin/comments/iohzmx/setting_and_getting_an_explanation/)
- url: https://www.reddit.com/r/Kotlin/comments/iohzmx/setting_and_getting_an_explanation/
---
Hey guys,

I recently started to learn Kotlin, and I'm having some troubles understanding the usage and the meaning of get and set methods in classes creation. I do not have any Java background (this is killing me, tbh) and I recently found them while starting to learn OOP and classes management. 

The online class I'm following is trying to explain me the importance of get and set, especially while checking for correct parameter inputs in classes but I can't seem to get the point.

Why using set and get methods if I can check for a parameter value directly inside the init method? Is there something I'm missing? 

I also looked on the documentation about them, but it shows how they works, not for what or when they are used.

Can someone better explain to me the meaning and the usage of these two methods? A brief explanations would be great if it is a huge thematic to talk about.

Thank you a lot! :D
