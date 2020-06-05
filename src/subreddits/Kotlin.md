# Kotlin
## [1][Introduction to event-driven microservices with Spring Cloud Stream and Kotlin](https://www.reddit.com/r/Kotlin/comments/gx16b0/introduction_to_eventdriven_microservices_with/)
- url: https://piotrminkowski.com/2020/06/05/introduction-to-event-driven-microservices-with-spring-cloud-stream/
---

## [2][Rob Percival's Complete Kotlin Developer Course](https://www.reddit.com/r/Kotlin/comments/gwz46t/rob_percivals_complete_kotlin_developer_course/)
- url: http://google.com/amp/s/onlinecoursesgalore.com/complete-kotlin-developer-course-rob-percival/amp/
---

## [3][Why You Should Use Kotlin Over Java?](https://www.reddit.com/r/Kotlin/comments/gwz8pl/why_you_should_use_kotlin_over_java/)
- url: https://www.reddit.com/r/Kotlin/comments/gwz8pl/why_you_should_use_kotlin_over_java/
---
 

Those days are gone when you blindly prefer Java for android based projects. As time passed, Kotlin emerged as the tough competitor of Java for android development. Now it’s time to switch over Kotlin. It is not just about the popularity amongst the developers but also the advantages over Java. Before comparing Kotlin with java, Let see some basics of Kotlin-

### What is Kotlin?

 

Kotlin is developed by JetBrains, and was first released in 2016. Immediately after the launch, it has gained a lot of popularity among the android community. And it gained momentum when google announced Kotlin as a first-class language for developing Android apps at the I/O 2017.

When it comes to its usability, Kotlin has a set of tools and frameworks that work with Java. Its advanced compiler can check the errors during compile time and run time also. It also reduces the lines of code to a great extent. Kotlin has everything that developers need and couldn’t find in Java or any other alternative like JRuby or Groovy. It is concise and completely interoperable with Java. Also, Kotlin came with higher compatibility with Android Studio in which IntelliJ IDEA, JetBrains’ own Java IDE serves as the backbone. Know the best Kotlin server-side frameworks at- [**5 Best Kotlin Server-side frameworks**](https://solaceinfotech.com/blog/5-best-kotlin-server-side-frameworks/).

## What Is The Problem With Java?

Java was released 20 years ago. And hence each new version of Java has to be compatible with the previous version. It becomes a challenge to compete with newer, more concise, expressive programming languages. Java is still the official programming language for android development but it just starts to show its limitations and its age.

## Reasons To Use Kotlin Over Java-

 

### 1. Easy Code-

Kotlin programs are easy to code as it does not need semicolons in their program. Also, they have smart casts and string templates. Java is not a concise language and it increases the chances of bugs. Kotlin is a concise language that means having fewer chances of runtime and compile-time errors. Kotlin provides an easy way to use mutable and immutable declarations for different data structures.

### 2. Completely Interoperable With Java-

One of the biggest advantages of Kotlin is, it is compatible with java. With multiple ways, it can easily use and exchange information from Java. Code of Kotlin and Java can co-exist in the same project. Kotlin can perform well with Java programming language. You can use multiple java libraries in Kotlin projects and this makes it more compatible. Compatibility is not just about the libraries but also with some advanced frameworks. You can use easily use Kotlin from Java by just converting a complete project to Kotlin. This is the biggest advantage for developers that they do not need to learn a new language. Anyone who is proficient with Java can easily code in Kotlin.

### 3. More Concise Than Java-

Every Kotlin programmer is happy with this feature. Kotlin is a statically-typed language that is very easy to read and write. You can solve the same problem with fewer lines of code and this reduces the number of bugs and crashes on the UX side. Due to the fewer lines of code, it becomes-

* Easy to maintain, 
* Easy to read and 
* Simple to apply changes when required.

There are some Kotlin features that are responsible for its code conciseness are- Data classes,  Smart casts, Type interface, Properties

### 4. Smart And Safe Compiler-

One of the main goals of Kotlin’s development team was to add a good compiler. And here are some important aspects of the compiler in Kotlin-

* It performs a lot of efficient checks, decreasing runtime errors, and the bugs in code.
* It detects errors at compile time not at runtime, by using the “fail-fast” principle.

### 5. Easy To Maintain-

Kotlin supports a lot of IDEs including Android Studio hence It is a “one-stop language” for all application development. You can freely use all those you have already tried and tested development tools that you’re comfortable to maintain code. As the Kotlin possesses fewer lines of code, it is easy to maintain and work upon. This is one of the best advantages over Java.

### 6. Eliminating Null References:

The biggest advantage over Java is null references. Accessing a member of null reference results in a null reference exception and this is a significant drawback of Java called as NullPointerException or NPE. The type system of Kotlin is dedicated to eliminate NullPointerException from code.

The possible causes of NPE’s may be:

1. An explicit call to throw NullPointerException();
2. Usage of the !! operator;
3. Java interoperation:

* When try to access a member on a null reference of a platform type;
* Generic types used for Java interoperation with incorrect nullability, e.g. a piece of Java code might add null into a Kotlin MutableList&lt;String&gt;, meaning that MutableList&lt;String?&gt; should be used for working with it;

4. Some data inconsistency with initialization, when:

An uninitialized this available in a constructor is passed and used somewhere (“leaking this”);

A superclass constructor calls an open member whose implementation in the derived class uses uninitialized state;

### 7. Solution To Java’s Issue-

Old java has a lot of issues. It hinders a well-known issue of null pointer. Kotlin tried to solve these issues. To overcome difficulties of java, it has adopted things from multiple languages like C# and Scala. It has instances from the Pascal language and is considered as very influential in Kotlin’s development. Some of the elements such as parameter lists and variable declarations with data type following a variable could be in Kotlin. It eliminates the error by removing the boilerplate code. It has some features like late initializations, delegations and also addresses a type safety. Kotlin was easy to type wrongly typed variables to a list before generics came along. This would typically lead to blowing up during the run time because the compiler does not detect it.

### 8. Improves Productivity-

Kotlin has been built with developer productivity in mind and is the key advantage of Kotlin over Java. Also, it’s a given that improved productivity returns to concise code itself, including its natural syntax and its general clean language design. It will need less time to write new code, deploy it, and maintain it in Kotlin. Concise and clear code in Kotlin helps to improve the developers’ productivity. Kotlin has a lot of powerful features that accelerate development tasks. These features are-

* Object declarations
* Parameter values
* Extension functions

You can also know the use of Kotlin for android development at- [**Why you should use Kotlin for android development?**](https://solaceinfotech.com/blog/why-you-should-use-kotlin-for-android-development/)
## [4][Kotlin 1.4-M2 Released](https://www.reddit.com/r/Kotlin/comments/gwfyn3/kotlin_14m2_released/)
- url: https://blog.jetbrains.com/kotlin/2020/06/kotlin-1-4-m2-released
---

## [5][Why doesn't Kotlin support multiple inheritance?](https://www.reddit.com/r/Kotlin/comments/gx0ham/why_doesnt_kotlin_support_multiple_inheritance/)
- url: https://www.reddit.com/r/Kotlin/comments/gx0ham/why_doesnt_kotlin_support_multiple_inheritance/
---
https://i.imgur.com/1LNHYwM.jpg
I know about this diamond problem but that can easily be solved. As a language feature if a class inherits multiple things that override the same thing it should throw a compile time error and just require the class itself to override it to avoid ambiguity. There could be a ClassName.super syntax to it so if you wanna use code from a certain class your class inherits from you can seamlessly do so. This is all that's needed to be able to seamlessly use multiple inheritance, this one problem that has such an easy solution isn't a valid reason not to support multiple inheritance.

If you agree with me and would like to see Multiple Inheritance come to Kotlin give this YouTrack issue a thumbs up: https://youtrack.jetbrains.com/issue/KT-39406 it would be greatly appreciated 🙂
## [6][Introducing GraphQL Kotlin Client](https://www.reddit.com/r/Kotlin/comments/gwitbq/introducing_graphql_kotlin_client/)
- url: /r/graphql/comments/gwis3f/introducing_graphql_kotlin_client/
---

## [7][Can Kotlin Classes be Split Across Multiple Files?](https://www.reddit.com/r/Kotlin/comments/gwvxd1/can_kotlin_classes_be_split_across_multiple_files/)
- url: https://www.reddit.com/r/Kotlin/comments/gwvxd1/can_kotlin_classes_be_split_across_multiple_files/
---
I haven't used Kotlin in a while and I was wondering if this was possible. C# is able to do so using the partial keyword and I was wondering if Kotlin had something similar.
## [8][Koge the new 100 % pure Kotlin OpenGL Game Engine](https://www.reddit.com/r/Kotlin/comments/gw259e/koge_the_new_100_pure_kotlin_opengl_game_engine/)
- url: https://www.reddit.com/r/Kotlin/comments/gw259e/koge_the_new_100_pure_kotlin_opengl_game_engine/
---
Koge (Kotlin OpenGL Game Engine) is a 2D game framework developed in Kotlin that works in Windows, Linux and Mac OS X. It is very simple and intuitive to use. For more information check [Koge on github](https://github.com/KogeLabs/Koge/blob/master/README.md)
## [9][Kotlin if statements and exception handling](https://www.reddit.com/r/Kotlin/comments/gwcusx/kotlin_if_statements_and_exception_handling/)
- url: https://www.reddit.com/r/Kotlin/comments/gwcusx/kotlin_if_statements_and_exception_handling/
---
A coworker and I noticed an interesting behaviour today, where the `get` operator on a `List`, if used inside an `if` statement, would never result in a crash due to an `OutOfBoundsException`, even if that index was definitely out of bounds. The statement will just return false instead.

This makes sense, in my opinion, but was wondering if anyone has any insight or documentation they can point me to regarding this, just out of interest and intrigue. I couldn't find anything with a quick google search.
## [10][Is Kotlin enough similar to Java to let me study concepts from java books and apply them in a kt app easily?](https://www.reddit.com/r/Kotlin/comments/gvy4ec/is_kotlin_enough_similar_to_java_to_let_me_study/)
- url: https://www.reddit.com/r/Kotlin/comments/gvy4ec/is_kotlin_enough_similar_to_java_to_let_me_study/
---
I'm a java programmer. I'm curious about Kotlin and I'd like to give it a try, if I could transfer my java knowledge easily. Especially the things regarding network programming and parallel programming.  
Any opinion/suggestions?
