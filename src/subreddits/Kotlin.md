# Kotlin
## [1][Multiplatform Kotlin library containing observable data structures, such as ObservableList, ObservableSet and ObservableMap](https://www.reddit.com/r/Kotlin/comments/fbevgc/multiplatform_kotlin_library_containing/)
- url: https://github.com/pearxteam/okservable
---

## [2][Simple kotlin library app](https://www.reddit.com/r/Kotlin/comments/fbr6fb/simple_kotlin_library_app/)
- url: https://youtu.be/Ta9pU3hiduk
---

## [3][Java Interoperability Not Working](https://www.reddit.com/r/Kotlin/comments/fbbr8j/java_interoperability_not_working/)
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
## [4][What's the use of this. ?](https://www.reddit.com/r/Kotlin/comments/fbh70i/whats_the_use_of_this/)
- url: https://www.reddit.com/r/Kotlin/comments/fbh70i/whats_the_use_of_this/
---
I can refer to variabeles and methods in the same class using this.method() or this.variabele . But the code also works when I don't use this. what's the differrence, and why should(n't) I use this.  ? thanks.
## [5][Function composition help](https://www.reddit.com/r/Kotlin/comments/fb4hpu/function_composition_help/)
- url: https://www.reddit.com/r/Kotlin/comments/fb4hpu/function_composition_help/
---
`fun compose(g: (Int) -&gt; Int, h: (Int) -&gt; Int): (Int) -&gt; Int {`

`}`

I'm supposed to return a function g(h(x)) but am having trouble putting it to code.  I've been returning single functions with the double colon notation, but that doesn't seem right in this case. Actually just having trouble wrapping my head around higher order functions in general but I guess that's why I'm here! Any help appreciated.
## [6][Conference for Kotliners 2020 - Official lineup announced](https://www.reddit.com/r/Kotlin/comments/fartgg/conference_for_kotliners_2020_official_lineup/)
- url: https://kotliners.com/conference
---

## [7][Migrating to Kotlin—what to look out for](https://www.reddit.com/r/Kotlin/comments/faut61/migrating_to_kotlinwhat_to_look_out_for/)
- url: https://engineering.autotrader.co.uk/2020/02/21/migrating-to-kotlin-what-to-look-out-for.html
---

## [8][Optimizing application performance with Amazon CodeGuru Profiler](https://www.reddit.com/r/Kotlin/comments/far74h/optimizing_application_performance_with_amazon/)
- url: https://aws.amazon.com/blogs/machine-learning/optimizing-application-performance-with-amazon-codeguru-profiler/
---

## [9][Writing Microservices in Kotlin with Ktor—a Multiplatform Framework for Connected System](https://www.reddit.com/r/Kotlin/comments/fa8zti/writing_microservices_in_kotlin_with_ktora/)
- url: https://www.infoq.com/articles/microservices-kotlin-ktor/
---

## [10][Learn Kotlin Channels in a Minute!](https://www.reddit.com/r/Kotlin/comments/faq0nf/learn_kotlin_channels_in_a_minute/)
- url: https://youtu.be/0PE71S80n7E
---

