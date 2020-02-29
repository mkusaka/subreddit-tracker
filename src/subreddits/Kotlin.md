# Kotlin
## [1][Java Interoperability Not Working](https://www.reddit.com/r/Kotlin/comments/fbbr8j/java_interoperability_not_working/)
- url: https://www.reddit.com/r/Kotlin/comments/fbbr8j/java_interoperability_not_working/
---
I'm studying using "Kotlin Programming: The Big Nerd Ranch Guide" (2018 edition) and on Java Interoperability chapter, the example is to use Java private fields as an ordinary property when calling inside Kotlin. No need to call getter/setter method for this Java private field. Here is the sample code:

`Jhava.java`

    public class Jhava {
      private String greeting = "BLARGH";
      public void setGreeting(String greeting) {
        this.greeting = greeting;
      }
    }

`Khotlin.kt`

    fun main() {
      val jhava = Jhava()
      jhava.greeting = "Hello"
    }

But `jhava.greeting = "Hello"` IntelliJ gives me an error `Cannot access 'greeting': it is private in 'Jhava'`. I'm not sure if the book is outdated or I'm doing something wrong here.
## [2][Function composition help](https://www.reddit.com/r/Kotlin/comments/fb4hpu/function_composition_help/)
- url: https://www.reddit.com/r/Kotlin/comments/fb4hpu/function_composition_help/
---
`fun compose(g: (Int) -&gt; Int, h: (Int) -&gt; Int): (Int) -&gt; Int {`

`}`

I'm supposed to return a function g(h(x)) but am having trouble putting it to code.  I've been returning single functions with the double colon notation, but that doesn't seem right in this case. Actually just having trouble wrapping my head around higher order functions in general but I guess that's why I'm here! Any help appreciated.
## [3][Conference for Kotliners 2020 - Official lineup announced](https://www.reddit.com/r/Kotlin/comments/fartgg/conference_for_kotliners_2020_official_lineup/)
- url: https://kotliners.com/conference
---

## [4][Migrating to Kotlin—what to look out for](https://www.reddit.com/r/Kotlin/comments/faut61/migrating_to_kotlinwhat_to_look_out_for/)
- url: https://engineering.autotrader.co.uk/2020/02/21/migrating-to-kotlin-what-to-look-out-for.html
---

## [5][Optimizing application performance with Amazon CodeGuru Profiler](https://www.reddit.com/r/Kotlin/comments/far74h/optimizing_application_performance_with_amazon/)
- url: https://aws.amazon.com/blogs/machine-learning/optimizing-application-performance-with-amazon-codeguru-profiler/
---

## [6][Writing Microservices in Kotlin with Ktor—a Multiplatform Framework for Connected System](https://www.reddit.com/r/Kotlin/comments/fa8zti/writing_microservices_in_kotlin_with_ktora/)
- url: https://www.infoq.com/articles/microservices-kotlin-ktor/
---

## [7][Learn Kotlin Channels in a Minute!](https://www.reddit.com/r/Kotlin/comments/faq0nf/learn_kotlin_channels_in_a_minute/)
- url: https://youtu.be/0PE71S80n7E
---

## [8][New to Kotlin, was reading Android Development Guide ... is this snippet a mistake?](https://www.reddit.com/r/Kotlin/comments/fae0k3/new_to_kotlin_was_reading_android_development/)
- url: https://www.reddit.com/r/Kotlin/comments/fae0k3/new_to_kotlin_was_reading_android_development/
---
[https://developer.android.com/guide/topics/connectivity/bluetooth-le#setup](https://developer.android.com/guide/topics/connectivity/bluetooth-le#setup)

Second snippet starts with this:

    private val BluetoothAdapter.isDisabled: Boolean
    get() = !isEnabled 

Where does `get()` come from and why is a function being assigned a Boolean value?
## [9][Kotlin Development Plan](https://www.reddit.com/r/Kotlin/comments/faccbt/kotlin_development_plan/)
- url: https://www.intuit.com/blog/uncategorized/kotlin-development-plan/
---

## [10][Composition over inheritance (and Kotlin)](https://www.reddit.com/r/Kotlin/comments/fa1v39/composition_over_inheritance_and_kotlin/)
- url: https://www.rockandnull.com/composition-over-inheridance/
---

