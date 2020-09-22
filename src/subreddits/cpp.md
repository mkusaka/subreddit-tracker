# cpp
## [1][C++ Jobs - Q3 2020](https://www.reddit.com/r/cpp/comments/hjnaf2/c_jobs_q3_2020/)
- url: https://www.reddit.com/r/cpp/comments/hjnaf2/c_jobs_q3_2020/
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

* [C++ Jobs - Q2 2020](https://www.reddit.com/r/cpp/comments/ft77lv/c_jobs_q2_2020/)
## [2][Why is it such an abysmal pain to use libraries in C++ compared to pretty much anything else?](https://www.reddit.com/r/cpp/comments/ix9n1u/why_is_it_such_an_abysmal_pain_to_use_libraries/)
- url: https://www.reddit.com/r/cpp/comments/ix9n1u/why_is_it_such_an_abysmal_pain_to_use_libraries/
---
# I recently realized something that's been annoying me for so long
## How to add a library in JavaScript:
- Type `npm install 'library'` in a shell on your project's directory.

## How to add a library in C#:
- Type `dotnet add package 'library'` in a shell on your project's directory.

## How to add a library in Go:
- Type `go get 'library_link'` in a shell. 

## How to add a library in Rust (And this is so "C++ is compiled" isn't an excuse):
- Lookup the last version of the library.
- Type `'library' = 'library_version'` on your project file.
- Restart your editor so the language server can get the symbols from the new library.

#### If you install `cargo-edit` you can alternatively just:
- Type `cargo add 'library'` in a shell on your project's directory. `cargo-edit` will do everything for you.

## How to add a library in C++:
- Prepare two folders for header include files, and library binaries.
- Append flags to your compiler to recognize them accordingly.
- Investigate which way the library works, praying the documentation of that is actually good.
### If the library is header-only:
- Add the header's required to your include path. 
- You should probably moduralize the code or spend 30 minutes setting up precompiled headers to avoid adding a lot to your compile times.
- And also lower your warning level, because even if you put `#pragma`s around the headers editors probably wont recognize them.
### Else, if the library distributes its binaries:
- Download the `.lib` or `.a` files from the last release and put them in your library folder.
- Tell the compiler to link your libraries.
- Put the downloaded header files in your include folder.
- If the library needs a `.dll`,  download it and paste it in the folder of your compiled executable. Distribute it with your shipped application.
- If you want to keep your application as a single executable, or distribute less dependencies, or have no need for an installer, all perfectly valid reasons, library owners *usually* have a static version.
- If they don't have a static version, spend an afternoon fighting the linker figuring out how to build the library yourself.
- If you want the library to link against the static runtime, the step above is required as well.
- Make sure to select the correct runtime library, or face really weird linker errors.
### If the library *doesn't* distribute its binaries:
- Clone the repository of the library and figure out how to build it yourself. There's usually a tutorial so it's not that complicated.
- Make sure to select the correct runtime library, or face really weird linker errors.
- Do this for every platform you want to distribute in.
### If the library uses CMake ᵒʰ ᵍᵒᵈ ʷʰʸ:
- You can choose two options:
#### Use CMake too
- Abandon your project's build system and spend days learning an entirely new language that everyone complains about
- Spend a lot of time googling and [figuring out the complicated "right" ways to do what's usually pretty simple](https://qrikko.blogspot.com/2016/05/cmake-and-how-to-copy-resources-during.html)
- Probably suffer from a loss in build time
#### Generate files for your compiler
- Install CMake GUI
- Learn how to use it and configure what you desire.
- Alternatively, learn yet another command line tool, or a tiny bit of CMake syntax to change what you want.
- Generate the files according to your platform and IDE.
- Build the library with a bloated IDE, or alternatively research how to build it with the much less known and documented command line tools.
- Make sure to select the correct runtime library, or face really weird linker errors.
### If you run into linker errors when running your program (and you will):
- If it's an unresolved external symbol, most likely your library needs another dependency linked. Lookup the function's name, what library it belongs to, and link against that too.
- If the unresolved function belongs to the standard library, you messed up. Compile the library again with the runtime library your compiler wants.
- If it's something else, google the error code and spend 30 minutes staring at StackOverflow.
# Just... why?
Am I missing something? Am I stupid and doing everything wrong? I really hope that's the case so I can get back to programming instead of fighting the linker. Every single time I see I need a library I'm like "Oh fuck..." to the point sometimes I just don't bother and decide to write things myself. 
Sorry about the rant. I'm kinda tired. Do you all have any similar experiences with this? Any tips to ease on the pain a bit? Thank you.
## [3][Attending the virtual CppCon 2020](https://www.reddit.com/r/cpp/comments/ixmfrb/attending_the_virtual_cppcon_2020/)
- url: http://meetingcpp.com/blog/items/Attending-the-virtual-CppCon-2020.html
---

## [4][This feels weird, but it does work. (std::map with a std::any, get/set templates). Please tell me what you think.](https://www.reddit.com/r/cpp/comments/ixn7lq/this_feels_weird_but_it_does_work_stdmap_with_a/)
- url: https://www.reddit.com/r/cpp/comments/ixn7lq/this_feels_weird_but_it_does_work_stdmap_with_a/
---
I need opinions if this is good/bad/ugly/python. If you would do a code review and saw this, what would you think? Do you feel uncomfortable with all this dynamic auto unknown? 

    /* some legacy class which there are a lot of instances off, which sometimes need to transfer data to another instance */

    std::map&lt;std::string, std::any&gt; _tVars;

    template &lt;typename T&gt;
    T getValue(const std::string &amp;key, T defaultValue) const
    {
        auto it = _tVars.find(key);
        if (it == _tVars.end())
            return std::any_cast&lt;T&gt;(defaultValue);

        return std::any_cast&lt;T&gt;(it-&gt;second);
    };

    template &lt;typename T&gt;
    void setValue(const std::string &amp;key, T value)
    {
        _tVars[key] = value;
    };


It's used as a sort of pseudo k/v:

    exampleClass.setValue("blah", aRandomTypeInstance);
    // somewhere else
    auto blah = exampleClass.getValue("blah", defaultValue);
   // blah should be aRandomTypeInstance, if it is not set, then it's a "defaultValue"
The programmer must ensure both sides use the correct key and validate the data in it.

The data that is passed around is unknown beforehand, either comes from a sybase database or user defined json (sometimes having base64 image in it). It's used purely for passing around data between class instances.
## [5][Ideas to use C++ for to try it out](https://www.reddit.com/r/cpp/comments/ixn62j/ideas_to_use_c_for_to_try_it_out/)
- url: https://www.reddit.com/r/cpp/comments/ixn62j/ideas_to_use_c_for_to_try_it_out/
---
Hey,

I've been curious about C++ for a while. I am a developer myself knowing C# after years of working with .NET backend, but it would be fun to do something in C++, to learn the language and see how it works. However, I am a bit lost where you use this today. One area seems to be game development. I would like to do some small project on the freetime to try C++ out. I am thinking of trying out game programming, Unreal Engine, to be precise.

I've done only some simple XML 2D game back in the old day but Unreal Engine seems pretty cool. Saw some YouTube videos where they use C++ and the Engine combined to create some nice things. Anyways, what I am wondering here is if you would do some fun side project using C++, what are some nice try-out things you could do? Is game programming as I am mentioning here the way to go or is something complete different better as game programming might be too much to take in for a simple side project?
## [6][CppCon 2020 slides](https://www.reddit.com/r/cpp/comments/iwzixs/cppcon_2020_slides/)
- url: https://github.com/CppCon/CppCon2020
---

## [7][CliWidgets: a little library for create simple widgets in your terminal programs](https://www.reddit.com/r/cpp/comments/ix68a0/cliwidgets_a_little_library_for_create_simple/)
- url: https://www.reddit.com/r/cpp/comments/ix68a0/cliwidgets_a_little_library_for_create_simple/
---
Hi, I have made this library as a first project with the idea to apply my acquired knowledge about c++ during the quarentine.

The idea of the library is to help people in the part of programing menus and other "widgets" for the terminal. 

The repo link is this: [https://github.com/lazaro92/CliWidget](https://github.com/lazaro92/CliWidget)
## [8][vcpkg: Accelerate your team development environment with binary caching and manifests | C++ Team Blog](https://www.reddit.com/r/cpp/comments/ix090v/vcpkg_accelerate_your_team_development/)
- url: https://devblogs.microsoft.com/cppblog/vcpkg-accelerate-your-team-development-environment-with-binary-caching-and-manifests/
---

## [9][DAW JSON Link v2 is release](https://www.reddit.com/r/cpp/comments/ix2ld6/daw_json_link_v2_is_release/)
- url: https://www.reddit.com/r/cpp/comments/ix2ld6/daw_json_link_v2_is_release/
---
Version two marks the culmination of some nice changes in JSON Link

* Non-intrusive mapping of your C++ structures to JSON classes and parsing without an intermediary DOM type.  Because we can specify the types, we can do extra type based checks on the input data too.

* A new lazy parsing DOM for exploring data.  This can be specified in the mappings above as a member type too.  The interface, also, allows iteration over array and class types.

* A new event based(SAX) parser 

* Constexpr

* Helper support for mapping to types like variant's and JSON array's to C++ classes, the cookbook section of the project has lots of documentation on mapping your JSON to C++

* Lots of documents with real working examples

The library is https://github.com/beached/daw_json_link

And there is a sample cmake project that uses it to make a config file parser, with C++ comment support https://github.com/beached/daw_json_link_config_parser
## [10][Inheritance](https://www.reddit.com/r/cpp/comments/ixmedw/inheritance/)
- url: https://www.reddit.com/r/cpp/comments/ixmedw/inheritance/
---
Hi can someone explain me Inheritance in depth?
## [11][Code Review of toml++](https://www.reddit.com/r/cpp/comments/iwych1/code_review_of_toml/)
- url: https://medium.com/@julienjorge/code-review-of-toml-f816a6071120
---

