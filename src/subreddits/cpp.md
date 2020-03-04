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
## [2][Super compact serialisation of C++ classes](https://www.reddit.com/r/cpp/comments/fd8jnf/super_compact_serialisation_of_c_classes/)
- url: https://www.reddit.com/r/cpp/comments/fd8jnf/super_compact_serialisation_of_c_classes/
---
When needing to save many different classes to disk into a human readable format and load them back (a pretty common but very boring task), I figured out this trick, which is probably the shortest way to do it without macros, working with any standard-compliant C++14 compiler (plus MSVC).

    struct Device : SerialisableBrief {
    	int timeout = key("timeout") = 1000;
    	std::string address = key("adress") = "192.168.32.28";
    	bool enabled = key("enabled") = false;
    	std::vector&lt;int&gt; ports = key("ports");
    }

With the inheritance, it gets methods `save()` and `load()` that allow saving it in JSON format as an object with keys `timeout`, `address`, `enabled` and `ports`.

Article how it works: [https://lordsof.tech/programming/super-compact-serialisation-of-c-classes/](https://lordsof.tech/programming/super-compact-serialisation-of-c-classes/)

Full code: [https://github.com/Dugy/serialisable/blob/master/serialisable\_brief.hpp](https://github.com/Dugy/serialisable/blob/master/serialisable_brief.hpp)
## [3]["Actor Model and C++: what, why and how?" March 2020 Edition](https://www.reddit.com/r/cpp/comments/fdbozc/actor_model_and_c_what_why_and_how_march_2020/)
- url: https://www.reddit.com/r/cpp/comments/fdbozc/actor_model_and_c_what_why_and_how_march_2020/
---
I posted the first version of this presentation here a couple of years ago. But it seems that there still is a great misunderstanding of the Actor Model and its benefits between C++ developers. I saw many examples of such misunderstanding in various discussions, the latest examples can be found in [this fresh topic on HackerNews](https://news.ycombinator.com/item?id=22457554). Sometimes C++ developers simply don't know that there are ready to use, stable and mature tools that implement the Actor Model in C++. And last but not least: some things have changed from the time I published the first version in 2017.

So I've updated my slides and [published it as "March 2020 Edition"](https://www.slideshare.net/YauheniAkhotnikau/actor-model-and-c-what-why-and-how-march-2020-edition). I hope it will be useful for someone.

Updated slides can be found on [SlideShare](https://www.slideshare.net/YauheniAkhotnikau/actor-model-and-c-what-why-and-how-march-2020-edition), or can be downloaded from [here](https://sourceforge.net/projects/sobjectizer/files/sobjectizer/Slides/Actor_Model_and_Cpp_what_why_and_how_%28March_2020%29.pdf).
## [4][The Performance Benefits of Final Classes | C++ Team Blog](https://www.reddit.com/r/cpp/comments/fcvlx2/the_performance_benefits_of_final_classes_c_team/)
- url: https://devblogs.microsoft.com/cppblog/the-performance-benefits-of-final-classes/?WT.mc_id=social-reddit-marouill
---

## [5][Do not believe anyone who says that you can’t use Valgrind on Windows](https://www.reddit.com/r/cpp/comments/fdchqk/do_not_believe_anyone_who_says_that_you_cant_use/)
- url: https://www.albertgao.xyz/2016/09/28/how-to-use-valgrind-on-windows/
---

## [6][Will calling “free” or “delete” in C/C++ release the memory to the system?](https://www.reddit.com/r/cpp/comments/fd1f5w/will_calling_free_or_delete_in_cc_release_the/)
- url: https://lemire.me/blog/2020/03/03/calling-free-or-delete/
---

## [7][C++ Quizzes](https://www.reddit.com/r/cpp/comments/fcqruv/c_quizzes/)
- url: https://www.reddit.com/r/cpp/comments/fcqruv/c_quizzes/
---
Here are some sites that have decent free C++ Quizzes:

* http://cppquiz.org/
* http://www.mycppquiz.com/
* http://www.interqiew.com/tests?type=cpp
* http://www.interqiew.com/ask?ta=tqdp02&amp;qn=1
* https://developers.google.com/edu/c++/quiz
* http://www.pvv.org/~oma/PubQuiz_ACCU_Apr2012.pdf
* http://www.pvv.org/~oma/PubQuiz_ACCU_Apr2013.pdf
* http://www.pvv.org/~oma/PubQuiz_ACCU_Apr2014.pdf
* http://www.pvv.org/~oma/PubQuiz_ACCU_Apr2016.pdf

If there are any others worth mentioning please post them here
## [8][Binlog - A high performance C++ log library to produce structured binary logs](https://www.reddit.com/r/cpp/comments/fcruym/binlog_a_high_performance_c_log_library_to/)
- url: https://github.com/Morgan-Stanley/binlog
---

## [9][[c++23 and beyond] Structured Binding extension ideas](https://www.reddit.com/r/cpp/comments/fcmm2s/c23_and_beyond_structured_binding_extension_ideas/)
- url: https://www.reddit.com/r/cpp/comments/fcmm2s/c23_and_beyond_structured_binding_extension_ideas/
---
I like structured bindings from c++17, I have some ideas for improving them for c++23 that I'd love to get preliminary feedback on before potentially getting involved in a paper.

I've seen a few referenced here and there, so I thought I'd unify them all in one post. Some of these ideas I've seen in [stack overflow](https://stackoverflow.com/questions/45541334/can-the-structured-bindings-syntax-be-used-in-polymorphic-lambdas) or inspired by the pattern matching paper ([P1371](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1371r1.pdf)) . If other people have proposed these same ideas in active papers, I'd love to know.

Let's get started

# Structured bindings as an argument (Edit: [P0931](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2018/p0931r0.pdf))

Wouldn't it be nice if we could declare a structured binding as a parameter in a function definition?

This is a pattern that comes up in range code or code that uses lambdas with tuple/pair args.

    std::unordered_map&lt;string, int&gt; myMap = /* ... */;
    auto transformed = myMap
        | ranges::views::transform([] (const auto&amp; keyValuePair) {
             const auto&amp; [key, value] = keyValuePair;
            // return something from key and value;
          })
        | ranges::to&lt;std::vector&gt;;

I would love to get rid of the `const auto&amp; [k, v] = keyValuePair;` and just write:

    std::unordered_map&lt;string, int&gt; myMap = /* ... */;
    auto transformed = myMap
        | ranges::views::transform([] (const auto&amp; [key, value]) {
            // return something from key and value
          })
        | ranges::to&lt;std::vector&gt;;

To me has a various obvious meaning, and I think it's not an ambiguous parse AFAIK.

This would also apply to normal functions, not just lambdas. So let's say we had a structured bindings compatible type:

    struct Point {
       int x;
       int y;
    };

And the declaration of the function taking that type 

    void function(Point p);

In the definition it would be nice to write:

    void function(Point [x, y]) {
       // do something with x and y
    }
instead of

    void function(Point p) {
       auto [x, y] = p;
       // do something with x and y
    }

This works as long as `Point` or similar is structured bindings decomposable. The structured binding decomposition would not have to appear in the function declaration, just the definition since this is basically an implementation detail.

And of course, with template/concept all the following would be valid (assuming valid type substitution):

    template &lt;typename T&gt;
    void function(T [x, y]) {}
    void function(auto [x, y]) {}
    void function(Concept auto [x, y]) {}

# Variadic bindings (Edit: [P1061](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1061r1.html))
Inspired from [here](https://www.reddit.com/r/cpp/comments/8abgg8/variadic_structured_binding/), but it'd be nice to be able to write expressions like the following:

    auto [first, second, ...rest] = unpackable;

Where `first` refs the first element, and `rest` packs the remaining and is usable just like a parameter pack:

    callSomethingVariadic(rest...);

You could also just omit the name `rest` if you didn't care:

    auto [first, second, ...] = unpackable;

This of course could be combined with bindings as parameters above:

    void function(SomeType [first, second, ...rest]) {}

# Named bindings
Currently structured bindings are positional, wouldn't it be cool if we could unpack fields by their name?

    struct Datum {
       string name;
       int id;
       bool isCool;
    } datum;
    
    auto [.isCool] = datum;
    // equivalent to
    auto isCool = datum.isCool;

This is intended to be very similar to what is possible in [Javascript](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment), and is similar to proposals in the pattern matching paper.
I use `.&lt;identifier&gt;` syntax to be consistent with designator initialization syntax. So that following would be consistent:

    auto [.name, .id, .isCool] = Datum{.name = "bob", .id = 2, .isCool = true};

 But ordinality restrictions should be lifted on decomposition, since initialization order does not really apply to structured bindings based on what they really are.

    auto [.id, isCool, .name] = datum;

This would combine well with "structured bindings as a parameter" so that you could accept a whole struct as a param (to be future proof) define exactly which args your function needs in its current impl:

    void doSomething(const Datum&amp; [.id, .name]) {
       // do something with name and id
    }

Or write code that concisely genericly expresses expectation of a certain field being present:

    void doSomethingGeneric(const auto&amp; [.foo]) {
       // use foo field of passed in object
    }

## Named unpack with rename (Edit: [P1371](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1371r1.pdf) uses `auto [.name: newName] = val`)
Named unpack with rename could be supported though I'm not 100% sold on it, e.g.:

    auto [.isCool, newName = .name] = datum;

Instead of:

    auto isCool = datum.isCool; (edited)
    auto newName = datum.name;

This feature would only be for renames. I would want arbitrary expressions to be disallowed here due to order of evaluation concerns and maintaining structured binding fields as aliases rather than independent variables. So at this point I think the following should be illegal:

    auto [newId = .id + 1] = datum; // illegal

Definitely initialization dependent on prior fields should be illegal:

    auto [
      newId = .id,
      newName = .name + std::to_string(newId)
    ] = datum; // illegal

One problem here is that `[newName = .name]` syntax totally implies that maybe an arbitrary expression can be substituted (as in lambdas). So perhaps we need a different syntax here. Javascript uses a colon for this:

    const {originalName: newName} = obj

But I don't think colon carries the same semantic meaning in c++, so the following would look a little strange in c++

    auto [.name: newName] = datum;

Another option could be fat or skinny arrow as in patter matching:

    auto [.name -&gt; newName] = datum;
    auto [.name =&gt; newName] = datum;

Which by themselves look fine but do not correspond with any other patterns. With this analysis I'm most partial to
    
    auto [newName = .name] = datum;

which is why I presented it first. But this problem gets even hairier when we talk about nesting...

## Combination with ordinals (probably don't allow this)
In general this would be mostly disallowed in combination with ordinal bindings:

    auto [.isCool, id] = datum; // disallowed

Use one or the other, not both. One exception could be if the named bindings follow all the ordinal ones:

    auto [newName, ..., .id] = datum

meaning, `datum.name` binds to first positional as `newName`, ignore all other positionals and bind `id` as `datum.id`. I don't see a use for this and its very presence suggests structuring a data type so that it has both an ordinal and non-ordinal (named) structure. So my perspective is we should probably just disallow this.

## Variadic named capture? (probably don't allow this)
What about the following:

    auto [.id, .isCool, ...rest] = datum;

On some level you understand what `rest` represents, a data structure that has all the fields of datum except for `.id` and `.isCool` (so just name). I don't really think this is a particularly useful object and we get a lot of hard questions as to what type the rest object actually has and how you're allowed to use it.

EDIT: This is allowed in javascript as nested object capture:

        const {field1, field2, ...rest] = obj;

Where `rest` is an object containing the same data as obj just without `field1` and `field2`. This is fine in javascript since objects are so dynamic in that language, but in c++ static typing would force rest to be an object of a new type that has no precedent (a struct alias without certain fields)?

## named capture (but with a function) (Edit: [P1371](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1371r1.pdf) has a solution for this)
So far named capture is pretty limited to simple structs (that which can be constructed by designated initializer). What if we had something more powerful:

    std::vector&lt;int&gt; vec; // some integer range
    auto [.begin, .end] = vec;
    // equivalent to:
    auto begin = vec.begin();
    auto end = vec.end();

If we expressed member capture as the rule that:

    auto [.identifier] = val;

is equivalent to:

    auto identifier = std::invoke(&amp;decltype(val)::identifier, val);

We get both forms automatically! This is a pretty radical idea though and breaks a lot of the rules associated with structured bindings, but I'm throwing it out there anyway...

# Nested Bindings (Edit: also referenced in [P1371](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1371r1.pdf))

This has beeen requested by a few but I wanted to reiterate that it works here and fits (kind of) well with the above. Nested bindings allow you to do the following:

    struct Datum {
       int first;

       struct Inner {
          double intensity;
          char code;
       } config;

       std::string color;
    } datum;

    auto [first, [intensity, code], color] = datum;

Of course all of the above mesh with this.
As an argument:

    auto function(Datum [first, [intensity, code], color] datum) {}

Variadics with named capture:

    auto [first, [.intensity], ...] = datum;

Fully nested named:

    auto [[.intensity] = .config] = datum;

We can see that the rename syntax doesn't work great with nesting. Consider the following type:

    struct Outer {
      struct Middle {
        struct Inner {
          int x;
          int y;
        };

        Innter inner;
      };

      Middle middle;
    } val;

We have several options to decompose this and get `int x, int y` in the end:

    auto [[[x, y]]] = val; // pure postional nested
    auto [[[.y, .x]]] = val; // positional nested -&gt; named
    auto [[[.x, .y] = .inner] = .mid] = val; // fully renamed
    auto [.x, .y] = val.middle.inner; // non-nested

The lesson here is that this:

    auto [[[.x, .y] = .inner] = .mid] = val;

Is a pretty terrible solution. No one can read that and immediately understand. It reads and writes in the completely wrong direction: you have to start with "[[[" which means you basically have to know your target variable depth before even writing.
    
Let's recall how javascript does this:

    const {middle: {inner: {x, y}}} = val;

If I'm honest, I still find this highly unreadable, maybe because in javascript the syntax makes me think I'm declaring a dictionary long before it makes me realize I'm referencing the `x` and `y` fields of `val`.

If we c++ this with arrows:

    auto [.middle =&gt; [.inner =&gt; [.x, .y]]] = val;

To me it's still not intuitive what the heck this does from an outsider perspective, but at least it's easier to write than:

    auto [[[.x, .y] = .inner] = .mid] = val;

If we look at "=&gt;" from a pattern matching perspective, then some intuition arises. We can describe the following:

    auto [.middle =&gt; [.inner =&gt; [.x, .y]]] = val;

As "match `val` against having field `.middle`, take result and match against having field `.inner` then take result and match against fields `.x and .y` capturing them."

# Conclusion/Summary

So I don't know what to make of this. The "rename" syntax as well as how it would apply with nesting is probably the hardest piece to wrangle here and has questionable value, but in my perspective the others would be pretty useful and intuitive. Reminders:
  
  Structured Binding as a param:

     [](auto [k, v]) {}
     // same as [](auto p) {auto [k, v] = p;}

  Variadic Bindings:

     auto [a, b, ...] = s;
     // no simple equivalent. for tuples:
     // auto&amp; a = std::get&lt;0&gt;(s);
     // auto&amp; b = std::get&lt;1&gt;(s);

  Structured binding by field name

     auto [.x] = s;
     // same as auto x = s.x

  Nested structured bindings (positional syntax):

     auto [a, [x, y]] = s;
     // same as: auto [a, tmp] = s; auto [x, y] = tmp;

  Combination:
     [] (auto [[.x, .y], ...]) {}
     // same as: [] (auto s) {
     // auto [tmp, ...] = s;
     // auto x = tmp.x;
     // auto y = tmp.y;}

I think having them would allow us to allow for some fresh and interesting programming paradigms. I'd love to hear your thoughts on some of these components as well as references to any papers that are currently proposing some of these ideas! I would love if c++23 brought with it a super powered update to structured bindings, since c++20 did very little to improve them.
## [10][Verbose syntax for small function/lambdas](https://www.reddit.com/r/cpp/comments/fcrgqs/verbose_syntax_for_small_functionlambdas/)
- url: https://www.reddit.com/r/cpp/comments/fcrgqs/verbose_syntax_for_small_functionlambdas/
---
    	using Op = std::function&lt;double(double, double)&gt; ;
    	Op fns = [](double x, double y) {return x + y ;} ;
    	Op fnm = [](double x, double y) {return abs(x) + abs(y) ;} ;
    	Op fnv = [](double x, double y) {return hypot(x, y) ;} ;
            ... etc

Is there a more concise way to define these small lambas? Also, the definition Op std::function&lt;double(double, double)&gt; tells us already what the signatures following are. There are dependencies to update if the definition of Op is changed.

Of course, we can use a macro to do the following instead:-

            using Op = std::function&lt;double(double, double)&gt;
            defop(fns, x + y) ; defop(fnm, abs(x) + abs(y)) ; defop(fnv, hypot(x, y)) ;

Which is concise, with fewer dependencies, possibly fewer still if we could extract the signature from Op. Just wondering if there is a tidier C++ way (without relying on macros)?
## [11][Quill - An asynchronous low latency logging library (C++14)](https://www.reddit.com/r/cpp/comments/fcbflb/quill_an_asynchronous_low_latency_logging_library/)
- url: https://github.com/odygrd/quill
---

