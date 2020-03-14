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
## [2][C++ Events Affected By Coronavirus](https://www.reddit.com/r/cpp/comments/fidita/c_events_affected_by_coronavirus/)
- url: https://www.reddit.com/r/cpp/comments/fidita/c_events_affected_by_coronavirus/
---
**Don't talk about the coronavirus. r/cpp is not the place to discuss or debate the coronavirus.**

# Purpose

Many C++ conferences, meetups, and even the next committee meeting have been affected by the ongoing public health crisis caused by the coronavirus.

I think it will be helpful to everyone to collect information about C++ events that have been or may be affected by the coronavirus. That includes not only cancellations/postponements, but also events that are planning to move forward but may be impacted in the future

**This is intended to be a living document. Please help be suggesting changes or additions.**


# Lexicon

**Cancelled** == The event will definitely not be rescheduled.

**Postponed** == The event will not happen on the planned dates, but may happen later this year.

**Still planned** == The event has not announced any changes to the planned dates.

# Conferences

- CppCon 2020 is still planned.
- C++Now 2020 is cancelled.
- ACCU 2020 is cancelled.
- [Cpp On Sea 2020](https://cpponsea.uk/) is still on for June 7-10.
- Core C++ 2020 is postponed from May until Fall.
- [Cpp Russia Moscow 2020](https://cppconf.ru/en/) is postponed from April 27-28 until June.
- [Advanced Developer's Conference++ 2020](https://adcpp.de/20/default) is postponed from May until October.

# C++ User Groups

- ACCU Bay Area is cancelled until further notice.

# ISO C++ Committee Meetings

- ISO C++ Varna 2020 is cancelled.
- ISO C++ New York 2020 is status quo.
## [3][A simple Star War game made purely out of C++ &amp; SFML with code](https://www.reddit.com/r/cpp/comments/fih7qs/a_simple_star_war_game_made_purely_out_of_c_sfml/)
- url: https://www.youtube.com/watch?v=3t0JbsZ6X4I
---

## [4][Proposal: Shadowing within the same scope](https://www.reddit.com/r/cpp/comments/fi0g7v/proposal_shadowing_within_the_same_scope/)
- url: https://www.reddit.com/r/cpp/comments/fi0g7v/proposal_shadowing_within_the_same_scope/
---
I'd like to gather some feedback on a short proposal I've written: [https://github.com/mgrech/cpp-proposals/blob/master/shadowing.md](https://github.com/mgrech/cpp-proposals/blob/master/shadowing.md)

I've seen this feature in Rust, where my initial reaction was that it looks very dangerous. After using it for a while though, I've found it to not only solve ergonomics issues in my code, but also prevent bugs, which is why I've decided to draft a proposal.

I also know that the committee doesn't like proposals that change fundamentals in the language, but I think this a small change with huge upsides and no backwards-compatibility issues. I've personally hit a nasty bug in my C++ code that could have been avoided with this very feature recently.
## [5][C++Now 2020 is Canceled](https://www.reddit.com/r/cpp/comments/fhqzw4/cnow_2020_is_canceled/)
- url: http://cppnow.org/announcements/2020/03/Canceled/
---

## [6][What's the reliable and FOSS way to compile C++ code for Windows (on Linux)?](https://www.reddit.com/r/cpp/comments/fhnlwh/whats_the_reliable_and_foss_way_to_compile_c_code/)
- url: https://www.reddit.com/r/cpp/comments/fhnlwh/whats_the_reliable_and_foss_way_to_compile_c_code/
---
I primarily develop my C++ applications (and everything else) on Linux, where there are two extremely good, FOSS, and maintained C++ compilers (GCC and clang).

Unfortunately, enough people use Windows that it's desirable to be able to produce Windows binaries. On windows, Microsoft's MSVC seems to have a similar level of support as GCC and Clang on Linux, but it has other issues. It's not FOSS. It doesn't work on Linux (forcing me to maintain 2 OS's for any sort of automated building). Besides not being FOSS, it has an unfavorable license that forces me to keep track of revenue, and eventually pay them. Perhaps these trade-offs are worth it, but that depends on what alternatives exist.

Clang, being extremely well supported, seems to be working on a Windows mode, but, unfortunately, that doesn't seem to be ready yet and still relies on MSVC's toolchain in key places, creating licensing issues. There may be some way to make it work, but I tried and failed to find it.

Mingw-w64 exists, works great, and has a good license, but, unfortunately, it seems to be not nearly as professionally run as Linux GCC (which has the Free Software Foundation, a serious and respected organization, behind it, and countless multi-billion dollar companies rely on it — it's not going anywhere). As evidence, consider that the Arch Linux packages (in AUR) are a full GCC version behind the native GCC packages, and that [their website](https://mingw-w64.org/doku.php) (for at least a week or so now) has had an expired TLS certificate (despite the fact that this fix can be automated quite easily). It looks like no one is actually putting the time in to maintain it, which makes me very cautious before relying on it in any way, and may make it worth it to pay Microsoft and deal with separate build infrastructure, although, there are even more flaws with that approach (then, anyone else who wants to work with my code needs to do the same, since I need ABI compatibility with 3rd party code in most of my projects).

One glimmer of hope is that the Arch Linux packages (in AUR)'s build scripts seem to pull their code from the main GCC repository. Is the necessary C++ code to build a working MinGW toolchain maintained with the GCC code in its main repository, in such a way that I can presumably learn how to build it whenever I want a maintained MinGw version?

So, how do you all address these concerns in your projects where you need to distribute Windows binaries? Is there a solution that addresses all concerns well?
## [7][Trip Report: C++ Standards Meeting in Prague, February 2020](https://www.reddit.com/r/cpp/comments/fhkeg3/trip_report_c_standards_meeting_in_prague/)
- url: https://botondballo.wordpress.com/2020/03/12/trip-report-c-standards-meeting-in-prague-february-2020/
---

## [8][Creating an intermediate GUI with dear imgui is not that difficult, but offers lots of flexibility that other GUI libraries miss out on.](https://www.reddit.com/r/cpp/comments/fhpm4t/creating_an_intermediate_gui_with_dear_imgui_is/)
- url: https://ruurdsdevlog.wordpress.com/2020/03/07/c-desktop-application-with-dear-imgui/
---

## [9][Varna ISO meeting postponed](https://www.reddit.com/r/cpp/comments/fhldku/varna_iso_meeting_postponed/)
- url: https://www.reddit.com/r/cpp/comments/fhldku/varna_iso_meeting_postponed/
---
Via [Bryce Lelbach on Twitter](https://twitter.com/blelbach/status/1238133359043174400).

The ACCU conference due to be held in Bristol, UK later this month [has also been cancelled](https://conference.accu.org/news/2020/202003121205_accu2020cancelled.html).

Update: C++ Now has been [cancelled too](http://cppnow.org/announcements/2020/03/Canceled/)
## [10][CppCast: PVS-Studio Static Analysis](https://www.reddit.com/r/cpp/comments/fhtjxd/cppcast_pvsstudio_static_analysis/)
- url: https://cppcast.com/yuri-minaev-static-analysis/
---

## [11][Question: Websocket in STL](https://www.reddit.com/r/cpp/comments/fh1vfb/question_websocket_in_stl/)
- url: https://www.reddit.com/r/cpp/comments/fh1vfb/question_websocket_in_stl/
---
I've found a number of libraries that do websockets in C++

* [https://docs.websocketpp.org/index.html](https://docs.websocketpp.org/index.html)
* [https://www.boost.org/doc/libs/1\_72\_0/libs/beast/doc/html/index.html](https://www.boost.org/doc/libs/1_72_0/libs/beast/doc/html/index.html)
* [https://libwebsockets.org](https://libwebsockets.org)

I was wondering what people used for this/is there an STL way of doing WebSockets easily?

The library needs to be permissive (Apache, MIT, BSD) because it's for an open source project.

Edit: I'm consuming a WebSocket API, not creating one.
## [12][Why there is no wchar_t variant of to_chars and from_chars?](https://www.reddit.com/r/cpp/comments/fh5h8q/why_there_is_no_wchar_t_variant_of_to_chars_and/)
- url: https://www.reddit.com/r/cpp/comments/fh5h8q/why_there_is_no_wchar_t_variant_of_to_chars_and/
---
It is especially painful on Windows where most text is `wchar_t`. `to_chars` and `from_chars` operate on `char*` buffers. But what is `char` anyway? It could be binary blob or it could be text. If it is text, what encoding is it in? It could be ASCII, UTF-8, ANSI (Windows codepage such as cp1250 or Shift-JIS), EBCDIC or something completely different. On some platforms, `char` is 16bits wide. This is bad for inter-compatibility. We already have `char` / `wchar_t` / `char8_t` / `char16_t` / `char32_t` types in Standard and also `std::string` is using them. Is there a reason why `to_chars` and `from_chars` didn't receive a bit of Unicode love? Is there a paper where I can read the rationale why this feature was not included?
