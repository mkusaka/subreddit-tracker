# Kotlin
## [1][Ktor Roadmap for 2020-2021 – Ktor Blog](https://www.reddit.com/r/Kotlin/comments/i70zbu/ktor_roadmap_for_20202021_ktor_blog/)
- url: https://blog.jetbrains.com/ktor/2020/08/10/ktor-roadmap-for-2020-2021/
---

## [2][Are inline classes considered an "over-engineering" thing if used a lot?](https://www.reddit.com/r/Kotlin/comments/i732j2/are_inline_classes_considered_an_overengineering/)
- url: https://www.reddit.com/r/Kotlin/comments/i732j2/are_inline_classes_considered_an_overengineering/
---
For example, let's say we have a customer entity and its using inline classes:

&amp;#x200B;

`data class Customer(`

   `val customerId: CustomerID,`

   `val customerName: CustomerName,`

   `val  customerEmail: CustomerEmail,`

   `val customerPassword: CustomerPassword`

`)`

&amp;#x200B;

Is it okay to be implemented this way or is it considered over-engineering? 

Just curious.
## [3][Collection of Kotlin online courses and tutorials for rookies](https://www.reddit.com/r/Kotlin/comments/i74irs/collection_of_kotlin_online_courses_and_tutorials/)
- url: https://www.reddit.com/r/Kotlin/comments/i74irs/collection_of_kotlin_online_courses_and_tutorials/
---
Created an curated amazing list of all the top-rated [Kotlin courses](https://blog.coursesity.com/best-kotlin-tutorials?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=redditPost&amp;utm_term=best-kotlin-js) of all time.

Many of these courses are very helpful to Kotlin for the beginners.  

Highly recommend.
## [4][📚 Sample project about Android Components Architecture on a Modular Word focused on the scalability, testability and maintainability written in Kotlin, following best practices using Jetpack](https://www.reddit.com/r/Kotlin/comments/i6zrl7/sample_project_about_android_components/)
- url: https://www.reddit.com/r/Kotlin/comments/i6zrl7/sample_project_about_android_components/
---
&amp;#x200B;

https://preview.redd.it/lnc36424e4g51.jpg?width=800&amp;format=pjpg&amp;auto=webp&amp;s=c1c81ca3c9a096d02f52288a9ae6e4b6fe997e99

The goal of the project is to demonstrate best practices, provide a set of guidelines, and present modern Android application architecture. This application may look simple, but it has all of these small details that will set the rock-solid foundation of the larger app suitable for bigger teams and long application lifecycle management.

Application is based, apply and strictly complies with each of the following 5 points:

* A **single-activity architecture**, using the Navigation component to manage fragment operations.
* **Android architecture components,** part of Android Jetpack give to project a robust design, testable and maintainable.
* Pattern **Model-View-ViewModel** (MVVM) facilitating separation of development of the graphical user interface.
* **S.O.L.I.D**  principles intended to make software designs more understandable, flexible and maintainable.
* **Modular app architecture** allows being developed features in isolation, independently from other features.

Of course, nothing is perfect and this project isn’t an exception, there are a lot of things to improve, but the iteration is the key to my point of view. But it's a good point to start and I hope that some of the architecture decisions taken help you to apply/inspire to your current or next project.

Thank you ❤️

Project link: [https://github.com/VMadalin/android-modular-architecture](https://github.com/VMadalin/android-modular-architecture)
## [5][Is there a way to define a type explicit and anonymous?](https://www.reddit.com/r/Kotlin/comments/i6qtfz/is_there_a_way_to_define_a_type_explicit_and/)
- url: https://www.reddit.com/r/Kotlin/comments/i6qtfz/is_there_a_way_to_define_a_type_explicit_and/
---
Hi!

I have a simple question but I do not know whether there is an answer.

When defining an anonymous class you can let the object inherit from multiple interfaces. The object has then both types like every (non anonymous) class that inherits from more then one interface.

    interface First {
     fun runFirst()
    }
    
    interface Second {
     fun runSecond()
    }
    
    fun test() {
     val both = object : First, Second {
            override fun runFirst() {
              println("first")
            }
    
            override fun runSecond() {
              println("second")
            }
        }
    
     // MainKt$test$both$1
     println(both.javaClass.typeName)
     
     // first
     // second
     both.runFirst()
     both.runSecond()
    }

My variable `both` in the given example is implicitly typed. But what is the Type? IntelliJ says `&lt;anonymous object: First, Second&gt;`. The JVM type name seems to be `MainKt$test$both$1`. The obvious way to solve this problem is to introduce a new interface:

    interface Both: First, Second

But then I could simply make it a non anonymous class instead of an interface.

Is there even a way to explicitly declare "anonymous types"? I am excited to hear your thoughts on this.
## [6][Method references and lambdas in lazy properties](https://www.reddit.com/r/Kotlin/comments/i70b7a/method_references_and_lambdas_in_lazy_properties/)
- url: https://blog.kotlin-academy.com/method-references-and-lambdas-in-lazy-properties-371dbbea857b
---

## [7][Creating a new entity in Exposed](https://www.reddit.com/r/Kotlin/comments/i6rpgh/creating_a_new_entity_in_exposed/)
- url: https://www.reddit.com/r/Kotlin/comments/i6rpgh/creating_a_new_entity_in_exposed/
---
I have been trying to use Exposed in my new project and have been struggling to find a way to create an empty entity.  I am trying to first create an entity and then insert that entity in one go, instead of defining each field at the time of insertion as done in the documentation.  So basically if I had a customer entity as follows:

    Class Customer(id: EntityID&lt;Int&gt;): IntEntity(id){
        companion object: IntEntityClass&lt;Customer&gt;(Customers)
        var name     by Customers.name
        var address  by Customers.address
        var phone    by Customers.phone
    }

I would want to do something like this:

    val id = EntityID(0, Customers)
    val newCustomer = Customer(id)
    newCustomer.name = "Jim"
    newCustomer.address = "123 S Main St"
    newCustomer.phone = "123-456-7890"
    
    Customers.insert(newCustomer)

instead of manually entering each field name every time I wanted to insert or update as shown in the documentation:

    Customers.insert{
        it[name] = "Jim"
        it[address] = "123 S Main St"
        it[phone] = "123-456-7890"
    }

I would prefer something like this because it would allow for me to write more generic code and not require hard coding for each of the many tables that I have in my application.
## [8][Guide to Quarkus with Kotlin](https://www.reddit.com/r/Kotlin/comments/i6g3pg/guide_to_quarkus_with_kotlin/)
- url: https://piotrminkowski.com/2020/08/09/guide-to-quarkus-with-kotlin/
---

## [9][Are there good resources for understanding the concept of delegation?](https://www.reddit.com/r/Kotlin/comments/i6ozgh/are_there_good_resources_for_understanding_the/)
- url: https://www.reddit.com/r/Kotlin/comments/i6ozgh/are_there_good_resources_for_understanding_the/
---
I've already read the docs, but they weren't very clear, and I couldn't find another good source. Do you have a recommendation?

Edit: it would be nice if it was free
## [10][Why JetBrains created Kotlin? - YouTube - Credits to all the articles from around 2011](https://www.reddit.com/r/Kotlin/comments/i6fkjq/why_jetbrains_created_kotlin_youtube_credits_to/)
- url: https://youtu.be/zTUZ7PQgWsI
---

