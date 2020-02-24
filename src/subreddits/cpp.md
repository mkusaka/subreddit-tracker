# cpp
## [1][C++ Jobs - Q1 2020](https://www.reddit.com/r/cpp/comments/eiila4/c_jobs_q1_2020/)
- url: https://www.reddit.com/r/cpp/comments/eiila4/c_jobs_q1_2020/
---
Rules For Individuals
---------------------

* **Don't** create top-level comments - those are for employers.
* Feel free to reply to top-level comments with **on-topic** questions.
* I will create one top-level comment for **meta** discussion.
* I will create another top-level comment for **individuals looking for work** and **community groups looking for sponsors**.

Rules For Employers
---------------------

* You must be hiring **directly**. No third-party recruiters.
* **One** top-level comment per employer. If you have multiple job openings, that's great, but please consolidate their descriptions or mention them in replies to your own top-level comment.
* **Don't** use URL shorteners. [reddiquette](https://www.reddithelp.com/en/categories/reddit-101/reddit-basics/reddiquette) forbids them because they're opaque to the spam filter.
* Templates are awesome. Please **use** the following template. As the "formatting help" says, use \*\*two stars\*\* to **bold text**. Use empty lines to separate sections.
* **Proofread** your comment after posting it, and edit any formatting mistakes.

---

\*\*Company:\*\* [Company name; also, use the "formatting help" to make it a link to your company's website, or a specific careers page if you have one.]

&amp;nbsp;

\*\*Type:\*\* [Full time, part time, internship, contract, etc.]

&amp;nbsp;

\*\*Description:\*\* [What does your company do, and what are you hiring C++ devs for? How much experience are you looking for, and what seniority levels are you hiring for? The more details you provide, the better.]

&amp;nbsp;

\*\*Location:\*\* [Where's your office - or if you're hiring at multiple offices, list them. If your workplace language isn't English, please specify it.]

&amp;nbsp;

\*\*Remote:\*\* [Do you offer the option of working remotely? If so, do you require employees to live in certain areas or time zones?]

&amp;nbsp;

\*\*Visa Sponsorship:\*\* [Does your company sponsor visas?]

&amp;nbsp;

\*\*Technologies:\*\* [Required: do you mainly use C++98/03, C++11, C++14, C++17, or the C++20 working draft? Optional: do you use Linux/Mac/Windows, are there languages you use in addition to C++, are there technologies like OpenGL or libraries like Boost that you need/want/like experience with, etc.]

&amp;nbsp;

\*\*Contact:\*\* [How do you want to be contacted? Email, reddit PM, telepathy, gravitational waves?]

---

Previous Post
--------------

* [C++ Jobs - Q4 2019](https://www.reddit.com/r/cpp/comments/dbqgbw/c_jobs_q4_2019/)
## [2][69-year-old father of C++: C++ has survived for decades, what keeps its vitality?](https://www.reddit.com/r/cpp/comments/f8jcuj/69yearold_father_of_c_c_has_survived_for_decades/)
- url: https://www.youtube.com/watch?v=mWtrLdR56oA
---

## [3][References, simply](https://www.reddit.com/r/cpp/comments/f8jrfk/references_simply/)
- url: https://herbsutter.com/2020/02/23/references-simply/
---

## [4][Basic concepts about building C++ software](https://www.reddit.com/r/cpp/comments/f8aj36/basic_concepts_about_building_c_software/)
- url: http://felipepiovezan.gitlab.io/blog/build_system_basics/
---

## [5][xmake v2.3.1 released, Seamless integration with other build systems](https://www.reddit.com/r/cpp/comments/f8jxj6/xmake_v231_released_seamless_integration_with/)
- url: https://tboox.org/2020/02/23/xmake-update-v2.3.1/
---

## [6][Harald Achitz: Const should be a type, give me operator .](https://www.reddit.com/r/cpp/comments/f8e1k1/harald_achitz_const_should_be_a_type_give_me/)
- url: https://youtu.be/oqGxNd5MPoM
---

## [7][C++ is NOT a superset of C: tentative definitions, implicit conversions, implicit declarations &amp;more](https://www.reddit.com/r/cpp/comments/f86rdq/c_is_not_a_superset_of_c_tentative_definitions/)
- url: https://www.youtube.com/watch?v=s3Cv0-U5bXc
---

## [8][Creating std::variant based on index at runtime](https://www.reddit.com/r/cpp/comments/f8cbzs/creating_stdvariant_based_on_index_at_runtime/)
- url: https://www.reddit.com/r/cpp/comments/f8cbzs/creating_stdvariant_based_on_index_at_runtime/
---
Hi everyone.

Here is a problem I faced recently: 

You have an array of std::variants you want to save to file and load later. Saving is easy - you just save the index of value's type in std::variant and then object data itself. Loading is a bit more tricky. I couldn't find a simple way to create a std::variant with value based on an index I read from file. Basically I wanted this:

    std::size_t typeIndexReadFromFile;
    std::variant&lt;A, B, C&gt; v = make_variant&lt;A, B, C&gt;(typeIndexReadFromFile);
    
    // Now I can run std::visit() on `v`, so it loads its data from file.

Finally I came up with this: [https://godbolt.org/z/sert3m](https://godbolt.org/z/sert3m)

Is there a simpler way to do it? Maybe I missed something in the standard library?
## [9][Q:I'm relativly new to c++ (just studied about threads and processes).my queation is: is there a simple solution to determine wether a thread finished an iteration or not](https://www.reddit.com/r/cpp/comments/f8ni2j/qim_relativly_new_to_c_just_studied_about_threads/)
- url: https://www.reddit.com/r/cpp/comments/f8ni2j/qim_relativly_new_to_c_just_studied_about_threads/
---

## [10][EnTT v3.3.0 is out: Gaming meets Modern C++](https://www.reddit.com/r/cpp/comments/f7tzbd/entt_v330_is_out_gaming_meets_modern_c/)
- url: https://www.reddit.com/r/cpp/comments/f7tzbd/entt_v330_is_out_gaming_meets_modern_c/
---
# What's EnTT

[`EnTT`](https://github.com/skypjack/entt) is a header-only library written in **modern C++**.

It's mainly known for its innovative and very performing [entity-component-system](https://en.wikipedia.org/wiki/Entity_component_system) (ECS) model. However, it offers also [many other things](https://github.com/skypjack/entt#introduction) useful for development, from flexible tools for managing signals to an integrated reflection system and so on. That's not all: some new modules are under development and will soon become part of the library (or at least I hope so).

`EnTT` is also a production-ready, fully documented and battle-tested library with a 100% coverage. [Among others](https://github.com/skypjack/entt/wiki/EnTT-in-Action), It's currently used in **Minecraft** by Mojang and the **ArcGIS Runtime SDK** by Esri.

# Why it's important for you

You should care because the **best feedback** I've ever received is perhaps when someone told me that this library was a great source for studying modern C++. Maybe he was right, maybe not. It's not for me to say, but it may still interest you.

In any case, I hope I've given other material to those interested in it in this sense and I'm open to any feedback, because ours is a language in which you never stop learning. So, go for it!

# What else

In my free time I'm running the [ECS back and forth](https://skypjack.github.io/tags/#ecs) series (along with [other posts](https://skypjack.github.io/)). I'd like to start something new about the internals of this library and the development of a software with `EnTT`. The latter should serve as a step by step guide to using all the feature offered by the library. I cannot really set a deadline for this but I'd like to know if there is interest in such a series before to start, so any feedback is appreciated!

If you are using `EnTT` and want to tell me *hello* or which of your products relies on it, do not hesitate to [contact me](https://github.com/skypjack)! For everyone else interested in the library, the [wiki](https://github.com/skypjack/entt/wiki) contains more than what I've said here and the [gitter](https://gitter.im/skypjack/entt) channel is a great place to come and ask your first question!

I'm looking forward to hearing from you. :)
## [11][How much could we gain by completely breaking ABI?](https://www.reddit.com/r/cpp/comments/f7zls8/how_much_could_we_gain_by_completely_breaking_abi/)
- url: https://www.reddit.com/r/cpp/comments/f7zls8/how_much_could_we_gain_by_completely_breaking_abi/
---
Quoting from 2020-02 Prague Meeting report:
&gt; Although there was strong interest in exploring how to evolve ABI in the future, we are not pursuing making C++23 a clean ABI breaking release at this time.

Is there any big list of things that could be improved, if it was decided for C++23 to be an ABI-breaking release?

I found a good paper mentioning some:
http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/p2028r0.pdf

But it only listed a few, and gave just vague estimates on what would be real gains.

Is there a bigger list, listing more of existing problems, with measured possible speed-ups or memory gains?

I encountered one myself, that is not widely known: 
- `sizeof(std::random_device)` on GCC is enormous, cause it is a union with Mersenne Twister. It was MT on Windows until it got implemented as indeterminate recently.

And I'm sure there are many many more.

I'd love to know more about it. Understand, where along the way there were made some small mistakes, and how big is the technical debt we are dragging behind us.

How much are we losing, and for how long can we afford not to break ABI? Are these changes big enough, that we can expect someone to fork the language just for performance reasons?
