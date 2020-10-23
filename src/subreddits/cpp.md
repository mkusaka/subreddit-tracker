# cpp
## [1][C++ Jobs - Q4 2020](https://www.reddit.com/r/cpp/comments/j3qems/c_jobs_q4_2020/)
- url: https://www.reddit.com/r/cpp/comments/j3qems/c_jobs_q4_2020/
---
Rules For Individuals
---------------------

* **Don't** create top-level comments - those are for employers.
* Feel free to reply to top-level comments with **on-topic** questions.
* I will create top-level comments for **meta** discussion and **individuals looking for work**.

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

\*\*Remote:\*\* [Do you offer the option of working remotely (permanently, or for the duration of the pandemic)? If so, do you require employees to live in certain areas or time zones?]

&amp;nbsp;

\*\*Visa Sponsorship:\*\* [Does your company sponsor visas?]

&amp;nbsp;

\*\*Technologies:\*\* [Required: do you mainly use C++98/03, C++11, C++14, C++17, or C++20? Optional: do you use Linux/Mac/Windows, are there languages you use in addition to C++, are there technologies like OpenGL or libraries like Boost that you need/want/like experience with, etc.]

&amp;nbsp;

\*\*Contact:\*\* [How do you want to be contacted? Email, reddit PM, telepathy, gravitational waves?]

---

Previous Post
--------------

* [C++ Jobs - Q3 2020](https://www.reddit.com/r/cpp/comments/hjnaf2/c_jobs_q3_2020/)
## [2][Has a release() method for std::vector been proposed?](https://www.reddit.com/r/cpp/comments/jgilit/has_a_release_method_for_stdvector_been_proposed/)
- url: https://www.reddit.com/r/cpp/comments/jgilit/has_a_release_method_for_stdvector_been_proposed/
---
A method akin to

    auto std::vector&lt;T, Alloc&gt;::release() -&gt; std::unique_ptr&lt;T[], _deleter_t&gt;;

where `_deteter_t` is a deleter, that destructs in an appropriate way. The state of the vector after a call to release would be, as if it was moved from.

One often works with libraries that adopt raw pointers on construction of their internals. When coding the algorithms constructing the data to be handed over, `vector` is a convenient class to use, with the api it provides. Moreover, often one uses other libraries for said construction, and they do return a `vector`. But currently, one has to copy the data over to a new pointer before handing it over to be adopted by another structure. Such a `release` method would allow easier data transfer in described situations. Of course, when the `_deleter_t` does something special, one still could not simply hand the pointer over, but this would be implied by returning a `unique_ptr`.

Has anything like this been proposed? If yes, has it been rejected due to drawbacks? If no, is this a method deemed not useful by most?

Edit:

Returning `unique_ptr` instead of the raw pointer, is to be explicit about ownership transfer, and ensure safety until something else adopts ownership over the memory.
## [3][Modern C++ Features - a few features that you may know from Python or JavaScript that you can use in C++ too! (For beginners)](https://www.reddit.com/r/cpp/comments/jg5oip/modern_c_features_a_few_features_that_you_may/)
- url: https://blog.yuvv.xyz/modern-cpp-features
---

## [4][CppCast: Programming History, JIT Compilations and Generic Algorithms](https://www.reddit.com/r/cpp/comments/jgfz75/cppcast_programming_history_jit_compilations_and/)
- url: https://cppcast.com/ben-deane-jit-history/
---

## [5][A live chrome demo! C++, ncnnRay ported to HTML / Webassembly via emscripten](https://www.reddit.com/r/cpp/comments/jgm54o/a_live_chrome_demo_c_ncnnray_ported_to_html/)
- url: https://www.reddit.com/r/cpp/comments/jgm54o/a_live_chrome_demo_c_ncnnray_ported_to_html/
---
 Finally a live demo in Chrome! 

Written in C++, u/ncnn \+ u/emscripten \+ u/WASM ported to u/HTML / u/Webassembly. The AI models are memory eaters, so be careful ... the best thing to do is to upload a face and click "RETINAF" which is the fastest model there. 

[A live demo](https://reddit.com/link/jgm54o/video/iqtwe3cg8uu51/player)

[https://quantscientist.github.io/ncnnRay/](https://quantscientist.github.io/ncnnRay/)

u/AI u/vulkan u/machinelearning  u/generativeart u/deeplearning  u/raylib
## [6][I've rewritten the compiler for my programming language into C++](https://www.reddit.com/r/cpp/comments/jglz0x/ive_rewritten_the_compiler_for_my_programming/)
- url: https://www.reddit.com/r/cpp/comments/jglz0x/ive_rewritten_the_compiler_for_my_programming/
---
My largest C++ project, by far, is the new compiler for my programming language (called AEC), targeting WebAssembly: [https://github.com/FlatAssembler/AECforWebAssembly](https://github.com/FlatAssembler/AECforWebAssembly)

My  previous compiler was targeting x86 and was written in JavaScript.  However, I got a bit dissatisfied with JavaScript. When programming in  JavaScript, you spend a lot of your time debugging errors that, in most  other programming languages, would have been caught by the compiler and  the compiler would issue a warning (or refuse to compile the program).  For instance, in JavaScript, even in the strict mode, if you mistype the  name of a property of some object, no syntax error or reference error  would happen, your program will just continue to go and probably  misbehave in hard-to-debug ways. Dynamic typing, while it can make your  code shorter sometimes, opens a room for a whole class of bugs. In most  programming languages, supplying a different number of arguments to a  function than a function expects will lead to a compile-time warning or  even an error. In JavaScript, the JIT-compiler compiles your code into  something that crashes in hard-to-debug ways. In fact, it will even  accept obviously wrong code such as  \`functions\_that\_returns\_an\_integer()()\` (and, of course, produce  assembly code that crashes). I knew C++ has improved a lot since the  time I first learned it. In old versions of C++, for example, you needed  to write lots of code to convert a number to string, while today there  is a \`std::to\_string\` template function. So, I decided to write my new  compiler in C++. One of the things I like about C++ is that it makes it  very easy to do deep copying, which is something you need a lot when  writing a compiler, and you don't want to waste time thinking about how  you will do that (or so I thought, C++ has some surprising undefined behaviors regarding the default copy constructors: [https://stackoverflow.com/questions/63951270/using-default-copy-constructor-corrupts-a-tree-in-c](https://stackoverflow.com/questions/63951270/using-default-copy-constructor-corrupts-a-tree-in-c)). I knew about Rust, however, in my experience, Rust is  annoying in that it often refuses to compile code in the name of safety,  and sometimes there is no obvious way to do what you want in the way  Rust would consider safe. Rust has many false positives when trying to  detect semantically invalid code, JavaScript has many false negatives.  C++ appears, to me, to be the best in that regard.

The specification for my programming language is available here: [https://flatassembler.github.io/AEC\_specification.html](https://flatassembler.github.io/AEC_specification.html)

The example program I like the most is the Analog Clock in AEC: [https://flatassembler.github.io/analogClock.html](https://flatassembler.github.io/analogClock.html)

So, what do you think about my work?
## [7][Fast, typo tolerant instant search engine written in C++](https://www.reddit.com/r/cpp/comments/jfddui/fast_typo_tolerant_instant_search_engine_written/)
- url: https://github.com/typesense/typesense
---

## [8][clang-tidy support in VSCode is possible](https://www.reddit.com/r/cpp/comments/jfktsf/clangtidy_support_in_vscode_is_possible/)
- url: https://www.reddit.com/r/cpp/comments/jfktsf/clangtidy_support_in_vscode_is_possible/
---
if you're an exceptionally attractive individual like me, the only things preventing you from retiring clion and adopting vscode (for everything) are **clang-tidy** and **cppcheck**. we can get clang-tidy by upvoting this feature request for adding clang-tidy support to the cpp extension for vscode: [https://github.com/microsoft/vscode-cpptools/issues/2908](https://github.com/microsoft/vscode-cpptools/issues/2908)

clang-tidy good. more clang-tidy.
## [9][Should I bother proposing that shared_ptr gets this ‘Next’ constructor argument for a future C++ standard?](https://www.reddit.com/r/cpp/comments/jg5jdp/should_i_bother_proposing_that_shared_ptr_gets/)
- url: https://www.reddit.com/r/cpp/comments/jg5jdp/should_i_bother_proposing_that_shared_ptr_gets/
---
I have an idea for a proposal. I want to try a reddit post that can get quickly shot down before putting effort into anything formal that will get instantly shredded by the committee.

Motivating example:

Say I have a linked list connected by shared_ptr/atomic&lt;shared_ptr&gt;. Calling the destructor for the list’s head pointer deletes the objects *recursively* causing stack overflow. 

Shared pointers therefore can’t generally be used for pointer based data structures. Sub-optimal hacks like while(pop()) plaster over the cracks, otherwise we’re back to using new/delete.

A solution:

    struct node {
        T data;
        shared_ptr&lt;node&gt; ptr_next_node;
    };

If shared_ptr&lt;node&gt; was aware that ‘node’ had the member ‘ptr_next_node’, it could delete chained nodes *iteratively*.

My idea is to have a Next argument, similar to the Deleter and Allocator arguments, a concept requiring the member function:

    shared_ptr&lt;node&gt;* next() noexcept;

In my example, this returns a pointer to ptr_next_node.

The destructor of shared_ptr can now call next() iteratively:
    
    auto next_ptr = next();
    delete this-&gt;ptr;
    delete this-&gt;control_block;
    while(next_ptr &amp;&amp; next_ptr-&gt;use_count == 1) {
         auto temp = next_ptr-&gt;next();
         delete next_ptr-&gt;ptr;
         delete next_ptr-&gt;control_block;
         next_ptr = temp;
    }

Problems:

What if a node has multiple children?

The implementation above needs a rework and the return type of next() changes to:

    std::vector&lt;shared_ptr&lt;node&gt;&gt; next();

The implementation would involve a loop of depth first searches picking the first non-nullptr shared_ptr at each node until everything was nullptr. This involves no allocations or recursion. Cycles are impossible because one node in a cycle must have a use_count() of at least 2.

What if there are circular references?

Now that we can iterate around the data structure with next() this problem is solvable but hard. For now cycles leak. A further proposal can include an overload of next() with the argument tag std::delete_cycles_v.

    std::vector&lt;shared_ptr&lt;node&gt;*&gt; next(std::delete_cycles_t);

What if different nodes in a data structure have different types?

I think it’s reasonable to require they share a base class with a virtual destructor if anyone really needs nodes of different types.

What if nodes have different allocators?

I would want to see a reasonable data structure where different nodes required different allocators before considering this a problem. std::vector treats specialisations with different allocators as essentially unrelated types.

What about reflection?

With reflection, a shared_ptr could generate its own Next argument. Would this be an ABI break? We could define a std::iterative_next object of type Next that uses reflection to infer which objects in a node are of type shared_ptr&lt;node&gt;. You could then just call

     std::shared_ptr&lt;T&gt; sp = make_shared_iterative&lt;T&gt;(args...);

What about weak_ptr?

Awkward. Control blocks would need their own copy of the next pointers.

Roast me.
## [10][Qt and idiomatic smart pointer usage](https://www.reddit.com/r/cpp/comments/jfoeht/qt_and_idiomatic_smart_pointer_usage/)
- url: https://www.reddit.com/r/cpp/comments/jfoeht/qt_and_idiomatic_smart_pointer_usage/
---
One of the projects I work on is a gui application based on a c++ library which is written in c++17 style. The gui use Qt, but the library itself doesn't use any Qt classes at all except for optional model objects which inherit QAbstractItemModel. Any pointers that are part of the library api are std::unique_ptr or (rarely) std::shared_ptr.

This creates some friction because Qt's ownership model doesn't mesh very well with modern coding styles. Even though other patterns are grudgingly tolerated, Qt wants you to create widget objects on the stack with new and pass in a parent pointer which will take ownership of the object. This feels like a major step backwards when the non-gui parts of the project have successfully eliminated usage of raw new / delete usage.

The solution we came up with is based on a helper class called ScopeGuard (*):

    class ScopeGuard
    {
    public:
        using Callback = std::function&lt;void()&gt;;

        ScopeGuard(Callback cb) noexcept;
            : cb_(cb)
        {
        }
        ~ScopeGuard()
        {
            if (cb_) { cb_(); }
        }

    private:
        const SimpleCallback cb_;
    };

With that class available, it's possible to write code for creating and displaying a modal dialog that looks like this:

    {
        ...
        auto dialog = std::make_unique&lt;MyDialog&gt;(this);
        auto postcondition = ScopeGuard{[&amp;]() {
            dialog-&gt;deleteLater();
            dialog.release();
        }};
        connect(dialog.get(), &amp;MyDialog::signal, this, &amp;MyType::slot)
        ...
        dialog-&gt;exec();
    }

Using this pattern I still allow Qt to control object lifetime on its own terms. In particular I don't need to worry about whether the signal/slot connections will be cleaned up before the dialog object.

At the same time, raw usage of new is avoided and the owership semantics are more clearly conveyed to anyone reading the code. Any coder looking at this function may not realize **why** ownership of the object is being given up in this way, but it is clear that what is happening is deliberate even to someone unfamiliar with Qt's ownership model.

(*) It's probably obvious but the class name "ScopeGuard" was invented by a team member who is a fan of Dlang.
## [11][I created a technical tutorial demonstrating how to design a language-agnostic cross-platform computer Vision SDK (written in C++).](https://www.reddit.com/r/cpp/comments/jfn1a8/i_created_a_technical_tutorial_demonstrating_how/)
- url: https://www.reddit.com/r/cpp/comments/jfn1a8/i_created_a_technical_tutorial_demonstrating_how/
---
In the tutorial, I explain how to:

* Build a basic computer vision library in C++
* Compile and cross-compile the library for AMD64, ARM64, and ARM32
* Package the library and all the dependencies as a single static library
* Automate unit testing
* Set up a continuous integration (CI) pipeline
* Write python bindings for our library
* Generate documentation directly from our API

The tutorial has a corresponding video explanation and all the code is open source an available on github. 

Check it out, I'm open to any feedback or suggestions too.

Hope it helps you get started with your next big C++ project!

[https://medium.com/trueface-ai/how-to-design-a-language-agnostic-cross-platform-computer-vision-sdk-e437ecac8b4e](https://medium.com/trueface-ai/how-to-design-a-language-agnostic-cross-platform-computer-vision-sdk-e437ecac8b4e)
