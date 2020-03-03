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
## [2][C++ Quizzes](https://www.reddit.com/r/cpp/comments/fcqruv/c_quizzes/)
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
## [3][[c++23 and beyond] Structured Binding extension ideas](https://www.reddit.com/r/cpp/comments/fcmm2s/c23_and_beyond_structured_binding_extension_ideas/)
- url: https://www.reddit.com/r/cpp/comments/fcmm2s/c23_and_beyond_structured_binding_extension_ideas/
---
I like structured bindings from c++17, I have some ideas for improving them for c++23 that I'd love to get preliminary feedback on before potentially getting involved in a paper.

I've seen a few referenced here and there, so I thought I'd unify them all in one post. Some of these ideas I've seen in [stack overflow](https://stackoverflow.com/questions/45541334/can-the-structured-bindings-syntax-be-used-in-polymorphic-lambdas) or inspired by the [pattern matching paper](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1371r1.pdf). If other people have proposed these same ideas in active papers, I'd love to know.

Let's get started

# Structured bindings as an argument

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

EDIT: I have learned that this was proposed in the following [paper](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2018/p0931r0.pdf)

# Variadic bindings
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

## Named unpack with rename
Named unpack with rename could be supported though I'm not 100% sold on it, e.g.:

    auto [.isCool, newName = .name] = datum;

Instead of:

    auto isCool = datum.name;
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

## Combination with ordinals
In general this would be mostly disallowed in combination with ordinal bindings:

    auto [.isCool, id] = datum; // disallowed

Use one or the other, not both. One exception could be if the named bindings follow all the ordinal ones:

    auto [newName, ..., .id] = datum

meaning, `datum.name` binds to first positional as `newName`, ignore all other positionals and bind `id` as `datum.id`. I don't see a use for this and its very presence suggests structuring a data type so that it has both an ordinal and non-ordinal (named) structure. So my perspective is we should probably just disallow this.

## Variadic named capture?
What about the following:

    auto [.id, .isCool, ...rest] = datum;

On some level you understand what `rest` represents, a data structure that has all the fields of datum except for `.id` and `.isCool` (so just name). I don't really think this is a particularly useful object and we get a lot of hard questions as to what type the rest object actually has and how you're allowed to use it.

## named capture (but with a function)
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

# Nested Bindings

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
## [4][Binlog - A high performance C++ log library to produce structured binary logs](https://www.reddit.com/r/cpp/comments/fcruym/binlog_a_high_performance_c_log_library_to/)
- url: https://github.com/Morgan-Stanley/binlog
---

## [5][Quill - An asynchronous low latency logging library (C++14)](https://www.reddit.com/r/cpp/comments/fcbflb/quill_an_asynchronous_low_latency_logging_library/)
- url: https://github.com/odygrd/quill
---

## [6][Verbose syntax for small function/lambdas](https://www.reddit.com/r/cpp/comments/fcrgqs/verbose_syntax_for_small_functionlambdas/)
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
## [7][In-class Member Initialisation: From C++11 to C++20](https://www.reddit.com/r/cpp/comments/fc9iiz/inclass_member_initialisation_from_c11_to_c20/)
- url: https://www.bfilipek.com/2015/02/non-static-data-members-initialization.html
---

## [8][WG21 in Prague - (Partial) Trip Report](https://www.reddit.com/r/cpp/comments/fcfbhj/wg21_in_prague_partial_trip_report/)
- url: /user/InbalL/comments/f5ftop/wg21_in_prague_partial_trip_report/
---

## [9][ABI Breaks: Not just about rebuilding](https://www.reddit.com/r/cpp/comments/fc2qqv/abi_breaks_not_just_about_rebuilding/)
- url: https://www.reddit.com/r/cpp/comments/fc2qqv/abi_breaks_not_just_about_rebuilding/
---
Related reading:

[What is ABI, and What Should WG21 Do About It?](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/p2028r0.pdf)

[The Day The Standard Library Died](https://cor3ntin.github.io/posts/abi/)

# Q: What does the C++ committee need to do to fix large swaths of ABI problems?

## A: Absolutely nothing

On current implementations, `std::unique_ptr`'s calling convention causes some inefficiencies compared to raw pointers.  The standard doesn't dictate the calling convention of `std::unique_ptr`, so implementers could change that if they chose to.

On current implementations, `std::hash` will return the same result for the same input, even across program invocations.  This makes it vulnerable to cache poisoning attacks.  Nothing in the standard requires that different instances of a program produce the same output.  An implementation could choose to have a global variable with a per-program-instance seed in it, and have `std::hash` mix that in.

On current implementations, `std::regex` is extremely slow.  Allegedly, this could be improved substantially without changing the API of `std::regex`, though most implementations don't change `std::regex` due to ABI concerns.  An implementation could change if it wanted to though.  However, very few people have waded into the guts of `std::regex` and provided a faster implementation, ABI breaking or otherwise.  Declaring an ABI break won't make such an implementation appear.

None of these issues are things that the C++ committee claims to have any control over.  They are dictated by vendors and by the customers of the vendors.  A new vendor could come along and have a better implementation.  For customers that prioritize QoI over ABI stability, they could switch and recompile everything.

Even better, the most common standard library implementations are all open source now.  You could fork the standard library, tweak the mangling, and be your own vendor.  You can then be in control of your own ~~destiny~~ ABI, and without taking the large up-front cost of reinventing the parts of the standard library that you are satisfied with.  libc++ has a [LIBCXX\_ABI\_UNSTABLE](https://github.com/llvm/llvm-project/blob/master/libcxx/CMakeLists.txt#L125) configuration flag, so that you always get the latest and greatest optimizations.  libstdc++ has a `--enable-symvers=gnu-versioned-namespace` configuration flag that is ABI unstable, and it goes a long way towards allowing multiple libstdc++ instances coexist simultaneously.  Currently the libc++ and libstdc++ unstable ABI branches don't have many new optimizations because there aren't many contributions and few people use it.  I will choose to be optimistic, and assume that they are unused because people were not aware of them.

If your only concern is ABI, and not API, then vendors and developers can fix this on their own without negatively affecting code portability or conformance.  If the QoI gains from an ABI break are worth a few days / weeks to you, then that option is available *today*.

# Q: What aspects of ABI makes things difficult for the C++ committee.

## A: API and semantic changes that would require changes to the ABI are difficult for the C++ committee to deal with.

There are a lot of things that you can do to a type or function to make it ABI incompatible with the old type.  The C++ committee is reluctant to make these kinds of changes, as they have a substantially higher cost.  Changing the layout of a type, adding virtual methods to an existing class, and changing template parameters are the most common operations that run afoul of ABI.

# Q: Are ABI changes difficult for toolchain vendors to deal with?

## A1: For major vendors, they difficulty varies depending on the magnitude of the break.

Since GCC 5 dealt with the std::string ABI break, [GCC has broken the language ABI 6 other times](https://gcc.gnu.org/onlinedocs/gcc/C_002b_002b-Dialect-Options.html), and most people didn't even notice.  There were several library ABI breaks (notably return type changes for std::complex and associative container erase) that went smoothly as well.  Quite a few people noticed the GCC 5 `std::string` ABI changes though.

In some cases, there are compiler heroics that can be done to mitigate an library ABI change.  You will get varying responses as to whether this is a worthwhile thing to do, depending on the vendor and the change.

If the language ABI changes in a large way, then it can cause substantially more pain.  GCC had a major language ABI change in GCC 3.4, and that rippled out into the library.  Dealing with libstdc++.so.5 and libstdc++.so.6 was a major hassle for many people, myself included.

## A2: For smaller vendors, the difficulty of an ABI break depends on their customer base.

These days, it's easier than ever to be your own toolchain vendor.  That makes you a vendor with excellent insight into how difficult an ABI change would be.

# Q: Why don't you just rebuild after an ABI change?

## A1: Are you rebuilding the standard library too?

Many people will recommend not passing standard library types around, and not throwing exceptions across shared library boundaries.  They often forget that at least one very commonly used shared library does exactly that... your C++ standard library.

On many platforms, there is usually a system C++ standard library.  If you want to use that, then you need to deal with standard library types and exceptions going across shared library boundaries.  If OS version N+1 breaks ABI in the system C++ standard library, the program you shipped and tested with for OS version N will not work on the upgraded OS until you rebuild.

## A2: Sometimes, rebuilding isn't enough

Suppose your company distributes pre-built programs to customers, and this program supports plugins (e.g. Wireshark dissector plugins).  If the plugin ABI changes, in the pre-built program, then all of the plugins need to rebuild.  The customer that upgrades the program is unlikely to be the one that does the rebuilding, but they will be responsible for upgrading all the plugins as well.  The customer cannot effectively upgrade until the entire ecosystem has responded to the ABI break.  At best, that takes a lot of time.  More likely, some parts of the ecosystem have become unresponsive, and won't ever upgrade.

This also requires upgrading large swaths of a system at once.  In certain industries, it is very difficult to convince a customer to upgrade anything at all, and upgrading an entire system would be right out.

Imagine breaking ABI on a system library on a phone.  Just getting all of the apps that your company owns upgraded and deployed at the same time as the system library would be a herculean effort, much less getting all the third party apps to upgrade as well.

There are things you can do to mitigate these problems, at least for library and C++ language breaks on Windows, but it's hard to mitigate this if you are relying on a system C++ standard library.  Also, these mitigations usually involve writing more error prone code that is less expressive and less efficient than if you just passed around C++ standard library types.

## A3: Sometimes you can't rebuild everything.

Sometimes, business models revolve around selling pre-built binaries to other people.  It is difficult to coordinate ABI changes across these businesses.

Sometimes, there is a pre-built binary, and the company that provided that binary is no longer able to provide updates, possibly because the company no longer exists.

Sometimes, there is a pre-built binary that is a shared dependency among many companies (e.g. OpenSSL).  Breaking ABI on an upgrade of such a binary will cause substantial issues.

# Q: What tools do we have for managing ABI changes?

## A: Several, but they all have substantial trade-offs.

The most direct tool is to just make a new thing and leave the old one alone.  Don't like `std::unordered_map`? Then make `std::open_addressed_hash_map`.  This technique allows new and old worlds to intermix, but the translations between new and old must be done explicitly.  You don't get to just rebuild your program and get the benefits of the new type.  Naming the new things becomes increasingly difficult, at least if you decide to not do the "lazy" thing and just name the new class `std::unordered_map2` or `std2::unordered_map`.  Personally, I'm fine with slapping a version number on most of these classes, as it gives a strong clue to users that we may need to revise this thing again in the future, and it would mean we might get an incrementally better hash map without needing to wait for hashing research to cease.

`inline` namespaces are another tool that can be used, but they solve far fewer ABI problems than many think.  Upgrading a type like `std::string` or `std::unordered_map` via `inline` namespaces generally wouldn't work, as user types holding the upgraded types would also change, breaking those ABIs.  `inline` namespaces can probably help add / change parameters to functions, and may even extend to updating empty callable objects, but neither of those are issues that have caused many problems in the C++ committee in the past.

Adding a layer of indirection, similar to COM, does a lot to address stability *and* extensibility, at a large cost to performance.  However, one area that the C++ committee hasn't explored much in the past is to look at the places where we already have a layer of indirection, and using COM-like techniques to allow us to add methods in the future.  Right now, I don't have a good understanding of the performance trade-offs between the different plug-in / indirect call techniques that we could use for things like `std::pmr::memory_resource` and `std::error_category`.

# Q: What can I do if I don't want to pay the costs for ABI stability?

## A: Be your own toolchain vendor, using the existing open-source libraries and tools.

If you are able to rebuild all your source, then you can point all your builds at a custom standard library, and turn on (or even make your own) ABI breaking changes.  You now have a competitive advantage, and you didn't even need to amend an international treaty (the C++ standard) to make it happen!  If your changes were only ABI breaking and not API breaking, then you haven't even given up on code portability.

Note that libc++ didn't need to get libstdc++'s permission in order to coexist on Linux.  You can have multiple standard libraries at the same time, though there are some technical challenges created when you do that.

# Q: What can I do if I want to change the standard in a way that is ABI breaking?

## A1: Consider doing things in a non-breaking way.

## A2: Talk to compiler vendors and the ABI Review Group (ARG) to see if there is a way to mitigate the ABI break.

## A3: Demonstrate that your change is so valuable that the benefit outweighs the cost, or that the cost isn't necessarily that high.

# Assorted points to make before people in the comments get them wrong

* I'm neither advocating to freeze ABI, nor am I advocating to break ABI.  In fact, I think those questions are too broad to even be useful.
* Fixing `std::unordered_map`'s performance woes would require an API break, as well as an ABI break.
* I have my doubts that `std::vector` could be made substantially faster with only an ABI break.  I can believe it if it is also coupled with an API break in the form of different exception safety guarantees.  Others are free to prove me wrong though.
* Making `&lt;cstring&gt;` `constexpr` will probably be fine.  The ABI issues were raised and addressed for `constexpr` `&lt;cmath&gt;`, and that paper is waiting in LWG.
* Filters on recursive\_directory\_iterators had additional concerns beyond ABI, and there wasn't consensus to pursue, even if we chose a different name.
* Making destructors implicitly `virtual` in polymorphic classes would be a massive cross-language ABI break, and not just a C++ ABI break, thanks to COM.  You'd be breaking the entire Windows ecosystem.  At a minimum, you'd need a way to opt out of `virtual` destructors.
* Are you sure that you are arguing against ABI stability?  Maybe you are arguing against backwards compatibility in general.
## [10][Kigs framework - A free, Open Source, Multi-Purpose and Cross-platform framework](https://www.reddit.com/r/cpp/comments/fccii8/kigs_framework_a_free_open_source_multipurpose/)
- url: https://www.reddit.com/r/cpp/comments/fccii8/kigs_framework_a_free_open_source_multipurpose/
---
Hello [r/cpp readers](https://www.reddit.com/r/cpp),

This is a short presentation of a newly Open Sourced C++ framework: [Kigs framework](https://kigs-framework.org/), [GitHub project](https://github.com/assoria/kigs).

Kigs framework is a lightweight, fast, scalable framework that [we](https://assoria.com/) have ported to several different platforms, and we used as a base to build all our projects from Nintendo DSi games to industrial robots simulator (have a look [here](https://kigs-framework.org/Projects)). 

Kigs framework gives access to high-level architecture and functionalities ( serialization, reflection, signals/slots, data-driven applications...) while maintaining low-level control (optimizations, platform-specific customization, easy C/C++ extern libraries usage ...).

## Why Use the Kigs Framework?

* You look for a free, MIT licensed framework, offering high-level functionalities such as serialization, instance factory, signals/slots management, scenegraph/rendering, Lua scripting... using C++.
* You want to learn game/application development using C++ and Lua scripting.
* You want to experiment with new ideas without starting from scratch.
* You are curious to see how we implemented this or that feature and perhaps want to help improve the framework.

We would be happy if others take over our framework, improve it and adapt it to their desires and their needs.

## What is available today?

Windows (x86, x64, UWP OpenGL or D3D ) and HTML5 (Emscripten, [here](https://kigs-framework.org/Samples) are the built samples) platforms are already released. Android platform should be released soon, and we would be happy to get some help to clean up and release iOS platform, or to add new platforms (Linux and MacOS ?).  

We also released a series of introductory articles on [CodeProject](https://www.codeproject.com/Articles/5253209/Kigs-Framework-Introduction-1-8).

I hope you will find this interesting,

    thanks,

          Stéphane
## [11][For those who want a simpler single-header asynchronous logger at the cost of slightly more latency (c++17)](https://www.reddit.com/r/cpp/comments/fcfd5i/for_those_who_want_a_simpler_singleheader/)
- url: https://gist.github.com/jrandom/5fab0cda11105e0e62c1f33c7a372b5e
---

