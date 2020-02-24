# typescript
## [1][Who's hiring Typescript developers - February](https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/)
- url: https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

Posting top level comments that aren't job postings, [that's a paddlin](https://i.imgur.com/FxMKfnY.jpg)

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][Gentle introduction into compilers. Part 1: Lexical analysis and Scanner in TypeScript](https://www.reddit.com/r/typescript/comments/f8ojhr/gentle_introduction_into_compilers_part_1_lexical/)
- url: https://indepth.dev/gentle-introduction-into-compilers-part-1-lexical-analysis-and-scanner/
---

## [3][More advanced types with TypeScript generics](https://www.reddit.com/r/typescript/comments/f8qytm/more_advanced_types_with_typescript_generics/)
- url: https://wanago.io/2020/02/24/more-advanced-types-with-typescript-generics/
---

## [4][definitelytyped.org domain has expired?](https://www.reddit.com/r/typescript/comments/f8j174/definitelytypedorg_domain_has_expired/)
- url: https://www.reddit.com/r/typescript/comments/f8j174/definitelytypedorg_domain_has_expired/
---
Just an FYI to... Microsoft I guess?

https://github.com/DefinitelyTyped/DefinitelyTyped

[Click the link in the repo description and see this.](https://i.imgur.com/F9szy1E.png)

[I made a GitHub issue as well.](https://github.com/DefinitelyTyped/DefinitelyTyped/issues/42576)
## [5][Has someone used TS with Redux Toolkit?](https://www.reddit.com/r/typescript/comments/f8nlqo/has_someone_used_ts_with_redux_toolkit/)
- url: https://www.reddit.com/r/typescript/comments/f8nlqo/has_someone_used_ts_with_redux_toolkit/
---
Hello all,

I have started to port my Redux code to Redux Toolkit as it is the recommended way from the Redux community and also I have started converting my code into TS, however I am struggling with the below code:

[https://codesandbox.io/s/redux-toolkit-with-typescript-1ks8s](https://codesandbox.io/s/redux-toolkit-with-typescript-1ks8s)  


Particularly,  typescript-eslint complains about "Missing return type on function" on the reducers functions as well as the action creators i.e.  loginLocalUser  


I have tried to several things but unsuccessfully.   


Anyone has done something similar by any chance?  


Thank you in advance and regards.
## [6][Type is missing properties from itself?](https://www.reddit.com/r/typescript/comments/f8e2sc/type_is_missing_properties_from_itself/)
- url: https://www.reddit.com/r/typescript/comments/f8e2sc/type_is_missing_properties_from_itself/
---
I'm facing a weird in TypeScript (version used is 3.7.4). For whatever reason, I'm unable to get a type to register as itself. I get the error:

&gt; Type 'Channel&lt;V&gt;' is missing the following properties from type 'Channel&lt;V&gt;': [messages], [putters], [takers], [racers]

For context, here's some of the code. I've reduced a lot of it to remove as many unnecessary pieces as I could while retaining the error. The `Channel&lt;V&gt;` there is the class of importance, and has a few symbol fields, which show up in the error message:

```TypeScript
export class Channel&lt;T&gt; {
  // `messages`, `putters`, `takers`, and `racers` are all symbols
  private [messages]: T[];
  private [putters]: Array&lt;() =&gt; void&gt;;
  private [takers]: Array&lt;(msg: T) =&gt; void&gt;;
  private [racers]: Array&lt;(chan: Channel&lt;T&gt;) =&gt; void&gt;;

  ...

  private static _foreach = &lt;V&gt;(
    arr: ChannelArray&lt;V&gt;,
    fn: (chan: Channel&lt;V&gt;) =&gt; void
  ) =&gt; {
    arr.forEach(fn); // Complains about type of `fn` here
  }
```

The `_forEach` static method is used to iterate over any collection of Channels, though here I've reduced it to only one type (array) to show off the error. The `ChannelArray&lt;V&gt;` type is just `Array&lt;Channel&lt;V&gt;&gt;` and is defined in a separate file. The weird thing is that if I define `type ChannelArray&lt;V&gt; = Array&lt;Channel&lt;V&gt;&gt;;` inside this file, the error goes away. Why would defining the helper type elsewhere cause this?

Edit: Clarified keys
## [7][OOP Polymorphism vs Discriminated Unions. When and why would you use them? Example included.](https://www.reddit.com/r/typescript/comments/f8grj0/oop_polymorphism_vs_discriminated_unions_when_and/)
- url: https://www.reddit.com/r/typescript/comments/f8grj0/oop_polymorphism_vs_discriminated_unions_when_and/
---
I am thinking about the best way my app can deal with classic programming design problems like the following:

https://www.typescriptlang.org/play/?ssl=1&amp;ssc=1&amp;pln=69&amp;pc=89#code/PQKhFgCgAIWgxArgOwMYBcCWB7ZBDAGylmCk2XQFMAnAMz1UugGUBHRPapgb2OgGtyAEwBc0AEQBndp0riA3H0mYAXpTHJEAWwBGNRZAC+ZCjXqNoAJUoY8yAOYEefQclESuth04V8A7phC6AAWGtp61AbQ0MGUmPbB6GG6+lDGkFDoAJ4ADkzMwXh50AC8LDJc0AA+Vjbodo6UBlC0KBg4yNCyeAAUkmIFRZQAlNC8MNCSAeiowdB9AHSuQqPj0dGoeJJMUhVyInzr0FzoiNSdkgvKarCTV6pNh9Cb2x51DT4HE0cnZxcLsXiiVulwCQWCUWg6XSUFQuEk6EmewGe1KYwEwjEuw4XHEABpJg8xABGAAMUIMcOQCOO7286lqXkaaO4GLcWM89Xp+OgYJCJNJBMBCSS0AATOTDM1IN0+nthvJoMBgNAyaSoLLOR8RorleLSeqMpBQBAYHAAPLmgAKxBApEg5CodAY+UKeQAogAPPBaHJOMZ8eyUdAAQS4vWGyQiBhhkFQBC2knKOMo5uQTEwvqcWkoFCTgw93qzzgmVIR1EQGGw1B6OWomAAbngqIS1FGaKt0tEcogdARMKhoEHQ+Geqsnr9ztAQphLtcmHAZ3OHpDoVBYQnJEnrEynGmM8Wc3mWG7KF6fX6Sxt4egK1Wa3XG82mHzQtBNClqATH02W8LEu21CdnwPZ9gOQ7BmGlARgG3y0qcU5Lgsr63Eh-7oKuaTrnGN5IimxJoumfjJrI+49GSCqwrhWr0gRZREYyXKNGRZIEhKlGQNI+ELMOUERrqKpqlANGNMSPGQaOCrQEAA

I have done some research, and this was the best quote on the subject I could find, but it was not specific to TypeScript:
https://softwareengineering.stackexchange.com/questions/367510/sum-types-vs-polymorphism

&gt; "It is important to recognize that interfaces and unions are kind of the opposite of each other: an interface defines some stuff the type has to implement, and the union defines some stuff the
&gt; consumer has to consider. If you add a method to an interface, you have changed that contract, and now every type that previously implemented it needs to be updated. If you add a new type to a union, &gt; you have changed that contract, and now every exhaustive pattern matching over the union has to be updated. They fill different roles, and while it may sometimes be possible to implement a system 
&gt; 'either way', which you go with is a design decision: neither is inherently better."


When you compile away the TypeScript the discriminated union version is about a third of the size, but I am not sure it 'feels as good' as the OOP version. Maybe I have missed an optimization? (I know the discriminated union exhaustive pattern matching with never trick too, but left that out of the example for a more direct comparison)

What do you make of these two, I can't decide which I like best.
## [8][Please help! I can't figure this out](https://www.reddit.com/r/typescript/comments/f8g3oy/please_help_i_cant_figure_this_out/)
- url: https://www.reddit.com/r/typescript/comments/f8g3oy/please_help_i_cant_figure_this_out/
---
Hi everybody, I recently started a new job in a small company where everyone is computer illiterate, and I'm trying to make a good impression by streamlining their workflow and making it more efficient.
One of the ideas I had was to build an Excel table with all current projects, where they can each write updates on their progress or add notes. A very important part of that is keeping track of when a project was last updated, so I'm looking for a way to have a cell automatically add a timestamp when another cell's value has changed.

I know there's a way to do that using VBA in Excel, but they only use Office 365 which only supports TypeScript. I don't know a thing about coding, and after a few hours [going](https://docs.microsoft.com/en-gb/office/dev/scripts/overview/excel) [over](https://docs.microsoft.com/en-gb/office/dev/scripts/tutorials/excel-tutorial) [the](https://docs.microsoft.com/en-gb/office/dev/scripts/develop/scripting-fundamentals) [documentation](https://docs.microsoft.com/en-gb/javascript/api/office-scripts/excel/excel.worksheet), I couldn't make heads or tails of it and decided it's time to ask for help.

tl;dr I need help writing a macro for Excel for web that automatically inputs the date and time when a cell has been changed.

Any assistance at all will be very much appreciated!
Thank you!
## [9][Typescript template to build a community app for all platforms](https://www.reddit.com/r/typescript/comments/f8a3mn/typescript_template_to_build_a_community_app_for/)
- url: https://www.reddit.com/r/typescript/comments/f8a3mn/typescript_template_to_build_a_community_app_for/
---
Hey everyone. I know that there are tons of TypeScript template but decided to share mine.

[https://github.com/shanhuiyang/TypeScript-MERN-Starter](https://github.com/shanhuiyang/TypeScript-MERN-Starter)

This template build backend, website, mobile apps in 100% Typescript.

Behind components are Express + MongoDB + React + ReactNative.

Hope you like it and look forward to your comments/issues.
## [10][Webpack/TypeScript/React starter kit as of 2020](https://www.reddit.com/r/typescript/comments/f866km/webpacktypescriptreact_starter_kit_as_of_2020/)
- url: https://www.reddit.com/r/typescript/comments/f866km/webpacktypescriptreact_starter_kit_as_of_2020/
---
Hey everyone. I know that there are tons of TypeScript starter kits but decided to share mine.

[https://krasimirtsonev.com/blog/article/beginning](https://krasimirtsonev.com/blog/article/beginning)

I did it because it's the very minimum setup and can be installed super quickly with just `npx beginning`.
## [11][JSBI; but for libraries, rather than application code](https://www.reddit.com/r/typescript/comments/f8a4cm/jsbi_but_for_libraries_rather_than_application/)
- url: https://github.com/AnyhowStep/bigint-lib#readme
---

