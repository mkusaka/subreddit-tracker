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
## [2][Cost of call: Virtual Method Call vs Non Virtual Method Call (C++)](https://www.reddit.com/r/cpp/comments/etpz7r/cost_of_call_virtual_method_call_vs_non_virtual/)
- url: https://www.reddit.com/r/cpp/comments/etpz7r/cost_of_call_virtual_method_call_vs_non_virtual/
---
Because I write program heavy dependent on interfaces, I was wondering, how much virtual method  call is slower than regular method call.

So I wrote program to check that.

Conclusion: Difference of call time between virtual methods and regular method is negligible in my opinion.

Why?

Is

    mov	ecx, DWORD PTR [esp+4]
    mov	eax, DWORD PTR [ecx]

combined by processor in one command/opcode?

Is time of receiving address of function in

    jmp	__ZN1X6MethodEv

the same (or almost) as in

    jmp	[DWORD PTR [eax+8]]

?

Or all this costs are so marginal that make them negligible?

I think it is the last one.

\##########################################################################

Measurement:

To prevent some optimizations, method calls are wrapped in to callers.

Glossary:

method - member function

Processor: AMD FX-8320 (8-Core 3.50GHz)

\##########################################################################

number of calls of each Call{...} function: 100 000 000

    1
    CallIdle:           0.163166s
    CallMethod:         0.220805s
    CallVirtualMethod:  0.213770s
    2
    CallIdle:           0.162361s
    CallMethod:         0.216204s
    CallVirtualMethod:  0.216748s
    3
    CallIdle:           0.164374s
    CallMethod:         0.215015s
    CallVirtualMethod:  0.215973s
    4
    CallIdle:           0.163159s
    CallMethod:         0.216749s
    CallVirtualMethod:  0.219072s
    5
    CallIdle:           0.162431s
    CallMethod:         0.215293s
    CallVirtualMethod:  0.215578s

Compiler:

GCC 8.1.0 (MinGW\_64)

Command

g++ -std=c++17 -Wall -Werror -pedantic -O3 -o prog.exe C.cpp main.cpp X.cpp

Command ASM:

g++ -std=c++17 -Wall -Werror -pedantic -O3 -S -masm=intel C.cpp main.cpp X.cpp

C.cpp

    #include "C.h"
    
    void CallIdle(X* x) {
        
    }
    void CallMethod(X* x) {
        x-&gt;Method();
    }
    void CallVirtualMethod(X* x) {
        x-&gt;VirtualMethod();
    }

C.s (single call of method)

    __Z8CallIdleP1X:
    LFB7:
        ret
    LFE7:
    __Z10CallMethodP1X:
    LFB8:
        mov	ecx, DWORD PTR [esp+4]
        jmp	__ZN1X6MethodEv
    LFE8:
    __Z17CallVirtualMethodP1X:
    LFB9:
        mov	ecx, DWORD PTR [esp+4]
        mov	eax, DWORD PTR [ecx]
        jmp	[DWORD PTR [eax+8]]     ; receiving address of virtual method from virtual table
    LFE9:

\##########################################################################

Added second call of method inside Call{...} function to force using 'call' instead of 'jump'.

number of calls of each Call{...} function: 100 000 000

    1
    CallIdle:           0.163293s
    CallMethod:         0.351328s
    CallVirtualMethod:  0.351897s
    2
    CallIdle:           0.166338s
    CallMethod:         0.355633s
    CallVirtualMethod:  0.352889s
    3
    CallIdle:           0.161744s
    CallMethod:         0.355947s
    CallVirtualMethod:  0.354000s
    4
    CallIdle:           0.163226s
    CallMethod:         0.355849s
    CallVirtualMethod:  0.354166s
    5
    CallIdle:           0.162738s
    CallMethod:         0.352778s
    CallVirtualMethod:  0.357627s

Compiler:

GCC 8.1.0 (MinGW\_64)

Command:

g++ -std=c++17 -Wall -Werror -pedantic -O3 -o prog.exe C.cpp main.cpp X.cpp

Command ASM:

g++ -std=c++17 -Wall -Werror -pedantic -O3 -S -masm=intel C.cpp main.cpp X.cpp

C.cpp

    #include "C.h"
    
    void CallIdle(X* x) {
        
    }
    void CallMethod(X* x) {
        x-&gt;Method();
        x-&gt;Method();
    }
    void CallVirtualMethod(X* x) {
        x-&gt;VirtualMethod();
        x-&gt;VirtualMethod();
    }

C.s (double call of method, forces using 'call' instead of 'jump')

    __Z8CallIdleP1X:
    LFB7:
        ret
    LFE7:
    __Z10CallMethodP1X:
    LFB8:
        push	ebx
        sub	esp, 8
        mov	ebx, DWORD PTR [esp+16]
        mov	ecx, ebx
        call	__ZN1X6MethodEv
        add	esp, 8
        mov	ecx, ebx
        pop	ebx
        jmp	__ZN1X6MethodEv
    LFE8:
    __Z17CallVirtualMethodP1X:
    LFB9:
        push	ebx
        sub	esp, 8
        mov	ebx, DWORD PTR [esp+16]
        mov	eax, DWORD PTR [ebx]
        mov	ecx, ebx
        call	[DWORD PTR [eax+8]]
        mov	eax, DWORD PTR [ebx]
        mov	ecx, ebx
        mov	eax, DWORD PTR [eax+8]
        add	esp, 8
        pop	ebx
        jmp	eax
    LFE9:

\##########################################################################

number of calls of each Call{...} function: 100 000 000

    1
    CallIdle:           0.162881s
    CallMethod:         0.214449s
    CallVirtualMethod:  0.238884s
    2
    CallIdle:           0.159846s
    CallMethod:         0.217064s
    CallVirtualMethod:  0.246562s
    3
    CallIdle:           0.159290s
    CallMethod:         0.214154s
    CallVirtualMethod:  0.240654s
    4
    CallIdle:           0.161967s
    CallMethod:         0.215867s
    CallVirtualMethod:  0.241240s
    5
    CallIdle:           0.158927s
    CallMethod:         0.212364s
    CallVirtualMethod:  0.241170s

Compiler:

MSVC 2013

Command:

cl /nologo /D\_USING\_V110\_SDK71\_=1 C.cpp main.cpp X.cpp /EHs /GR /O2 /MD /MP /Fomsvc\_make\_temp\\ /Fdmsvc\_make\_temp\\ /link /out:prog.exe

Command ASM:

cl /nologo /D\_USING\_V110\_SDK71\_=1 C.cpp main.cpp X.cpp /FAs /EHs /GR /O2 /MD /MP

C.asm (single call of method)

    ; Listing generated by Microsoft (R) Optimizing Compiler Version 18.00.40629.0 
    
        TITLE	G:\work\code\RandomMinGW\examples\example_093_virtual_vs_non_virtual\C.cpp
        .686P
        .XMM
        include listing.inc
        .model	flat
    
    INCLUDELIB MSVCRT
    INCLUDELIB OLDNAMES
    
    PUBLIC	?CallIdle@@YAXPAVX@@@Z				; CallIdle
    PUBLIC	?CallMethod@@YAXPAVX@@@Z			; CallMethod
    PUBLIC	?CallVirtualMethod@@YAXPAVX@@@Z			; CallVirtualMethod
    EXTRN	?Method@X@@QAEXXZ:PROC				; X::Method
    ; Function compile flags: /Ogtpy
    ; File g:\work\code\randommingw\examples\example_093_virtual_vs_non_virtual\c.cpp
    ;	COMDAT ?CallVirtualMethod@@YAXPAVX@@@Z
    _TEXT	SEGMENT
    _x$ = 8							; size = 4
    ?CallVirtualMethod@@YAXPAVX@@@Z PROC			; CallVirtualMethod, COMDAT
    
    ; 11   :     x-&gt;VirtualMethod();
    
        mov	ecx, DWORD PTR _x$[esp-4]
        mov	eax, DWORD PTR [ecx]
        jmp	DWORD PTR [eax+4]
    ?CallVirtualMethod@@YAXPAVX@@@Z ENDP			; CallVirtualMethod
    _TEXT	ENDS
    ; Function compile flags: /Ogtpy
    ; File g:\work\code\randommingw\examples\example_093_virtual_vs_non_virtual\c.cpp
    ;	COMDAT ?CallMethod@@YAXPAVX@@@Z
    _TEXT	SEGMENT
    _x$ = 8							; size = 4
    ?CallMethod@@YAXPAVX@@@Z PROC				; CallMethod, COMDAT
    
    ; 7    :     x-&gt;Method();
    
        mov	ecx, DWORD PTR _x$[esp-4]
        jmp	?Method@X@@QAEXXZ			; X::Method
    ?CallMethod@@YAXPAVX@@@Z ENDP				; CallMethod
    _TEXT	ENDS
    ; Function compile flags: /Ogtpy
    ; File g:\work\code\randommingw\examples\example_093_virtual_vs_non_virtual\c.cpp
    ;	COMDAT ?CallIdle@@YAXPAVX@@@Z
    _TEXT	SEGMENT
    _x$ = 8							; size = 4
    ?CallIdle@@YAXPAVX@@@Z PROC				; CallIdle, COMDAT

\##########################################################################

Added second call of method inside Call{...} function to force using 'call' insted of 'jump'.

number of calls of each Call{...} function: 100 000 000

    1
    CallIdle:           0.162718s
    CallMethod:         0.321905s
    CallVirtualMethod:  0.346607s
    2
    CallIdle:           0.165081s
    CallMethod:         0.326066s
    CallVirtualMethod:  0.350423s
    3
    CallIdle:           0.160939s
    CallMethod:         0.323516s
    CallVirtualMethod:  0.346217s
    4
    CallIdle:           0.159470s
    CallMethod:         0.321954s
    CallVirtualMethod:  0.350222s
    5
    CallIdle:           0.159132s
    CallMethod:         0.320411s
    CallVirtualMethod:  0.350353s

C.asm (double call of method, forces using 'call' instead of 'jump')

    ; Listing generated by Microsoft (R) Optimizing Compiler Version 18.00.40629.0 
    
        TITLE	G:\work\code\RandomMinGW\examples\example_093_virtual_vs_non_virtual\C.cpp
        .686P
        .XMM
        include listing.inc
        .model	flat
    
    INCLUDELIB MSVCRT
    INCLUDELIB OLDNAMES
    
    PUBLIC	?CallIdle@@YAXPAVX@@@Z				; CallIdle
    PUBLIC	?CallMethod@@YAXPAVX@@@Z			; CallMethod
    PUBLIC	?CallVirtualMethod@@YAXPAVX@@@Z			; CallVirtualMethod
    EXTRN	?Method@X@@QAEXXZ:PROC				; X::Method
    ; Function compile flags: /Ogtpy
    ; File g:\work\code\randommingw\examples\example_093_virtual_vs_non_virtual\c.cpp
    ;	COMDAT ?CallVirtualMethod@@YAXPAVX@@@Z
    _TEXT	SEGMENT
    _x$ = 8							; size = 4
    ?CallVirtualMethod@@YAXPAVX@@@Z PROC			; CallVirtualMethod, COMDAT
    
    ; 10   : void CallVirtualMethod(X* x) {
    
        push	esi
    
    ; 11   :     x-&gt;VirtualMethod();
    
        mov	esi, DWORD PTR _x$[esp]
        mov	ecx, esi
        mov	eax, DWORD PTR [esi]
        call	DWORD PTR [eax+4]
    
    ; 12   :     x-&gt;VirtualMethod();
    
        mov	eax, DWORD PTR [esi]
        mov	ecx, esi
        pop	esi
        jmp	DWORD PTR [eax+4]
    ?CallVirtualMethod@@YAXPAVX@@@Z ENDP			; CallVirtualMethod
    _TEXT	ENDS
    ; Function compile flags: /Ogtpy
    ; File g:\work\code\randommingw\examples\example_093_virtual_vs_non_virtual\c.cpp
    ;	COMDAT ?CallMethod@@YAXPAVX@@@Z
    _TEXT	SEGMENT
    _x$ = 8							; size = 4
    ?CallMethod@@YAXPAVX@@@Z PROC				; CallMethod, COMDAT
    
    ; 7    :     x-&gt;Method();
    
        mov	ecx, DWORD PTR _x$[esp-4]
        call	?Method@X@@QAEXXZ			; X::Method
    
    ; 8    :     x-&gt;Method();
    
        mov	ecx, DWORD PTR _x$[esp-4]
        jmp	?Method@X@@QAEXXZ			; X::Method
    ?CallMethod@@YAXPAVX@@@Z ENDP				; CallMethod
    _TEXT	ENDS
    ; Function compile flags: /Ogtpy
    ; File g:\work\code\randommingw\examples\example_093_virtual_vs_non_virtual\c.cpp
    ;	COMDAT ?CallIdle@@YAXPAVX@@@Z
    _TEXT	SEGMENT
    _x$ = 8							; size = 4
    ?CallIdle@@YAXPAVX@@@Z PROC				; CallIdle, COMDAT
    
    ; 4    :     
    ; 5    : }
    
        ret	0
    ?CallIdle@@YAXPAVX@@@Z ENDP				; CallIdle
    _TEXT	ENDS
    END

\##########################################################################

Whole code:

main.cpp

    #include &lt;windows.h&gt;
    #include &lt;stdio.h&gt;
    #include "C.h"
    
    #define for_ix(number) for (unsigned ix = 0; ix &lt; unsigned(number); ++ix)
        
    //------------------------------------------------------------------------------
        
    namespace {
    LARGE_INTEGER s_freq;   // time in performance ticks
    LARGE_INTEGER s_start;  // time in performance ticks
    LARGE_INTEGER s_stop;   // time in performance ticks
    double s_diff;          // time in seconds
    } // namespace
    
    void Initialize();
    void Start();
    void Stop();
    double Get(); // time in seconds
    
    //------------------------------------------------------------------------------
    
    int main() {
        Initialize();
        
        X* x = new X();
        
        enum { TIMES = 100000000 };
        
        Start();
        for_ix (TIMES) {
            CallIdle(x);
        }
        Stop();
        printf("CallIdle:           %fs\n", Get());
        
        Start();
        for_ix (TIMES) {
            CallMethod(x);
        }
        Stop();
        printf("CallMethod:         %fs\n", Get());
        
        Start();
        for_ix (TIMES) {
            CallVirtualMethod(x);
        }
        Stop();
        printf("CallVirtualMethod:  %fs\n", Get());
        
        delete x;
        
        return 0;
    }
    
    //------------------------------------------------------------------------------
    
    inline void Initialize() {
        QueryPerformanceFrequency(&amp;s_freq);
    }
    
    inline void Start() {
        QueryPerformanceCounter(&amp;s_start);
    }
    
    inline void Stop() {
        QueryPerformanceCounter(&amp;s_stop);
        s_diff = double(s_stop.QuadPart - s_start.QuadPart) / s_freq.QuadPart;
    }
    
    inline double Get() { 
        return s_diff; 
    } 

X.h

    #ifndef X_H_
    #define X_H_
    
    class X {
    public:
        X() {}
        virtual ~X() {}
        
        void Method();
        virtual void VirtualMethod();
    };
    
    #endif // X_H_

X.cpp

    #include "X.h"
    #include &lt;stdio.h&gt;
    
    void X::Method() {
        //puts("1");
    }
    
    void X::VirtualMethod() {
        //puts("2");
    }

C.h

    #ifndef C_H_
    #define C_H_
    
    #include "X.h"
    
    void CallIdle(X* x);
    void CallMethod(X* x);
    void CallVirtualMethod(X* x);
    
    #endif // C_H_

C.cpp

    #include "C.h"
    
    void CallIdle(X* x) {
        
    }
    void CallMethod(X* x) {
        x-&gt;Method();
        x-&gt;Method();
    }
    void CallVirtualMethod(X* x) {
        x-&gt;VirtualMethod();
        x-&gt;VirtualMethod();
    }
    

EDIT: Minor fixes to description.
## [3][The C++ Lifetime Profile: How It Plans to Make C++ Code Safer](https://www.reddit.com/r/cpp/comments/eta01f/the_c_lifetime_profile_how_it_plans_to_make_c/)
- url: https://pspdfkit.com/blog/2020/the-cpp-lifetime-profile/
---

## [4][To Bind and Loose a Reference](https://www.reddit.com/r/cpp/comments/et6eht/to_bind_and_loose_a_reference/)
- url: https://thephd.github.io/to-bind-and-loose-a-reference-optional
---

## [5][Hidden reinterpret_casts](https://www.reddit.com/r/cpp/comments/esz058/hidden_reinterpret_casts/)
- url: https://quuxplusone.github.io/blog/2020/01/22/expression-list-in-functional-cast/
---

## [6][CppCast: Circle](https://www.reddit.com/r/cpp/comments/et4dxp/cppcast_circle/)
- url: https://cppcast.com/circle-language/
---

## [7][I "Finished" Creating A Game Engine Programming In C++ Tutorial Series And Would Like Feedback For The Postmortem And Future Series](https://www.reddit.com/r/cpp/comments/et0jeg/i_finished_creating_a_game_engine_programming_in/)
- url: https://www.reddit.com/r/cpp/comments/et0jeg/i_finished_creating_a_game_engine_programming_in/
---
Back in July of 2017 I started posting a series of videos under the premise that "teaching is the best way to learn."  The premise what that I would read through the book "Game Programming In C++ Start To Finish" written by Erik Yuzwa and then adapt the information from said book into a video format. I now know that this was a terrible, terrible idea.

The code in the book was practically unusable., the software was out of date, and the version of OpenGL that it was using was already depreciated by the time I started my videos. I knew this going in. I further knew that attempting to get into game development via learning to program a game engine is the dumbest and hardest possible way to do so.

But I thought then, and I still think now, that my game ideas don't properly fit in any framework provided by current engines. I now know that I could have simply make a scaled down version of my ideas, refined them, learned how to code, and THEN learned how to make an engine that fit their needs, but at this point I'm 10 years in and I'm not backing down now. 

So here's what I'm here to ask. If you have some spare time, could you give some or all the my tutorial series a watch, and tell me just how badly I messed everything up?

I'm looking for critiques like "this part was to fast" "this part was to slow" "To little information given here" "Should have mentioned where you got this or that resource" "Should have sited your sources for this" etc.

All the criticisms will be cataloged along with my own criticisms of myself into a postmortem video. Then. Well. Then I'm going to try again. With a new book this time. One that I think will actually help me this time: Game Engine Architecture 2nd Edition by Jason Gregory. Also, if you have other books that you know to be excellent resources for building a good cross-platform game engine in C++, let me know so I can take a look at them. Also if there are any other series that cover the same topics, let me know about that too. I'll also be looking at the lists of books and series that already exist on this subreddit and others. Thanks in advance for any help you can provide!

[Link to the series!](https://www.youtube.com/playlist?list=PLx25SxVo-qjn5eWRaBY96khBJSLy2wBty)
## [8][Refactoring generates so many new source files](https://www.reddit.com/r/cpp/comments/et0t56/refactoring_generates_so_many_new_source_files/)
- url: https://www.reddit.com/r/cpp/comments/et0t56/refactoring_generates_so_many_new_source_files/
---
I'm adding unit tests to a couple of codebases which haven't had them before.  To break dependencies I'm factoring out interfaces (base classes with only pure virtual functions) for each class's collaborators.

**This generates so. many. new. files.**

C++ already averages 2x as many files as classes because of the .h/.cpp split.  Now I have 4x as many files as (concrete) classes.  I'm not even counting test mocks, I'm talking just in the core project.

I don't really want to put the interfaces in the same .h file as the concrete classes, because declaring the interface in its own file allows me to strip out #include's if they were only necessary to the concrete implementation but not needed by the interface class.

I could eliminate the .cpp file for the interface, but I'm using those to define the pure virtual destructors.  Also to follow John Lakos's advice that every class / header should be included in a corresponding .cpp file even if the .cpp file otherwise has nothing to do.  (It provides assurances the header is correct and self-sufficient.)

One of my projects is a legacy MFC app.  In order to disentangle the business logic from the presentation logic from the Lovecraftian horror of Windows programming I'm splitting every dialog class into a view, presenter, and model.  But that's 3x the files.  Most need interfaces, so 2x3.  Count .h/.cpp and you have 2x2x3 = 12 files per dialog.

Twelve.

Times 50 dialogs.

**How am I supposed to keep this shit organized?**

I mean, it is what it is.  I'm just looking for assurances before I go on a tear doing a refactoring that turns 100 code files into 600 code files that I'm not missing something here.
## [9][Concise offline reference](https://www.reddit.com/r/cpp/comments/estu90/concise_offline_reference/)
- url: https://www.reddit.com/r/cpp/comments/estu90/concise_offline_reference/
---
I am learning C++. In C, I can type `man function_name` and get a concise but helpful reference of a function.

I tried `stdman` which is a man page version of cppreference website. It is nowhere near C man pages which neatly give function signature, arguments and return value. (maybe partly because C functions are simple and no templating). Compared to them, stdman or whatever C++ reference is horrible. Moreover, it is hard to get complete information about a function like std::string::copy() (Maybe I am missing something here though).

similarly, in python also, documentation was one help() away in interpreter. Is there an analogue for C++? I don't need to see all possibilities with the function or heck, not even complete type signature. For some example:

    std::find(iterator m, iterator n, value) - find given element in given collection

I will not be learning function from scratch every time, I just need it to remind me.

if you know something like that, please comment.
## [10][How to introduce exceptions into historically exception free code base?](https://www.reddit.com/r/cpp/comments/espg2t/how_to_introduce_exceptions_into_historically/)
- url: https://www.reddit.com/r/cpp/comments/espg2t/how_to_introduce_exceptions_into_historically/
---
We have a massive C++ code base where much of the core was written in the early 2000's using error codes and a custom template library to support an exception free environment.

Over the years we've lost support in the team to continue to maintain and debug the custom template library and there's been a desire to "modernize" our code.  One aspect being the use of the STL.  However, this creates some issues as we can now have functions that can return error codes and throw exceptions which makes things increasingly difficult to reason about.

Has anyone has experience going from exception free code to introducing something like the STL?  If so, what was your approach?

A couple things I could think of off the top of my head.  Have one master catch statement in our runtime code.  Or perhaps create macro that wraps every STL call that can throw in a try catch (though this feels cumbersome).  I'll admit, I don't have much experience with exceptions as I've mainly worked in this codebase for the past couple years.
## [11][C++ Packaging and Design Rules | Large-Scale C++ Volume I: Process and Architecture](https://www.reddit.com/r/cpp/comments/esxoq7/c_packaging_and_design_rules_largescale_c_volume/)
- url: http://www.informit.com/articles/article.aspx?p=2995361
---

