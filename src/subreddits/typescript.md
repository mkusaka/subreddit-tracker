# typescript
## [1][Who's hiring Typescript developers - August](https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/)
- url: https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/
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
## [2][Typesafe alternatives for redux-saga?](https://www.reddit.com/r/typescript/comments/i3gvul/typesafe_alternatives_for_reduxsaga/)
- url: https://www.reddit.com/r/typescript/comments/i3gvul/typesafe_alternatives_for_reduxsaga/
---
I have been using redux-saga in couple TypeScript projects but big problem with redux-saga is generators which lose type safety when you \`yield\` any value.  


I am looking for other library to handle side-effects without losing type safety. Please suggest libraries that you have used to handle side effects with TypeScript and tell what did you use. It would be also nice if the library works with Redux :)
## [3][The missing TS+Vue Auth0 tutorial](https://www.reddit.com/r/typescript/comments/i3jj3q/the_missing_tsvue_auth0_tutorial/)
- url: https://www.reddit.com/r/typescript/comments/i3jj3q/the_missing_tsvue_auth0_tutorial/
---
I love Vue, and I find it nice how it works with TS, but the lack of tutorials with Vue+TS is frustrating. So if I already had to recreate the code of one of the Auth0 docs in TS, I thought I might as well publish it. So here it is, if you are setting up a project with Vue+TS using Auth0, you might wanna take a look at this.  


[https://blog.risingstack.com/auth0-vue-typescript-quickstart-docs/](https://blog.risingstack.com/auth0-vue-typescript-quickstart-docs/)  


Feedback is always welcome.
## [4][React HOC](https://www.reddit.com/r/typescript/comments/i3irl4/react_hoc/)
- url: https://www.reddit.com/r/typescript/comments/i3irl4/react_hoc/
---
Hey Guys!

&amp;#x200B;

I tried to search for this but didn't really find what I was looking for, hopefully you can help me!

&amp;#x200B;

I am looking for a HOC example that injects a Prop(s) and forces the Wrapped Component Props to extends the Injected Props.

&amp;#x200B;

I saw HOC that did this here [https://github.com/realm/react-realm-context](https://github.com/realm/react-realm-context)  (the withRealm HOC)

&amp;#x200B;

What I have so far ...

    interface WrappedProps {}
    
    const Title: React.FC&lt;WrappedProps&gt; = () =&gt; {
      return &lt;Text&gt;Title&lt;/Text&gt;;
    };
    
    interface InjectedProps {
      title: string;
    }
    
    function withCoolTitle&lt;P&gt;(Component: React.ComponentType&lt;P &amp; NewProps&gt;) {
      return (props: P) =&gt; &lt;Component {...(props as P)} title="Cool title" /&gt;;
    }
    
    export default withCoolTitle(Title);

A very basic example which doesn't really make sense but I didn't want a bloated example with a User example ...

&amp;#x200B;

So I want that WrappedProps needs to extends InjectedProps AND when I use the Title Comp that I don't need to include InjectedProps because it is injected by the HOC.

In this example WrappedProps does not need to extends InjectedProps nor do I need to inject title when I use &lt;Title /&gt; ...

&amp;#x200B;

Your help would be greatly appreciated!
## [5][What's the best bundler to deal with TypeScripts broken .mjs output?](https://www.reddit.com/r/typescript/comments/i2y54h/whats_the_best_bundler_to_deal_with_typescripts/)
- url: https://www.reddit.com/r/typescript/comments/i2y54h/whats_the_best_bundler_to_deal_with_typescripts/
---
So TypeScript cannot output `.mjs` files and [this creates a massive headache for many years](https://github.com/microsoft/TypeScript/issues/18442#issuecomment-666190219). What's the best bundler which can deal with this mess and is close to a `tsc` experience re building and hot reloading?

*Edit: to the downvoters, this is no rant, I just look for help. So rather reply and tell me what to do instead of mindless downvoting, thx.*
## [6][Error with generics when assigning function](https://www.reddit.com/r/typescript/comments/i3e416/error_with_generics_when_assigning_function/)
- url: https://www.reddit.com/r/typescript/comments/i3e416/error_with_generics_when_assigning_function/
---
I'm getting an error and I don't know why
Sample code:

    const fn: &lt; T &gt; () =&gt; Promise &lt;T&gt; = () =&gt; new Promise (resolve =&gt; resolve (2));

I get error 2345, type 5 cannot be assigned to type T or promiseLike&lt;T&gt; or undefined
Does anyone know how to fix this
My actual use case involves passing the function as a callback but this is the minimum amount of code to produce the error
Removing the generic from the declaration gives an error that Promise needs a type argument
## [7][NGX-Translate and Server-Side Rendering - Angular Universal](https://www.reddit.com/r/typescript/comments/i3320a/ngxtranslate_and_serverside_rendering_angular/)
- url: https://andremonteiro.pt/ngx-translate-angular-universal-ssr/
---

## [8][Reflecting business logic rules into the domain models using typescript by Mohsen Sareminia](https://www.reddit.com/r/typescript/comments/i2j9tg/reflecting_business_logic_rules_into_the_domain/)
- url: https://www.reddit.com/r/typescript/comments/i2j9tg/reflecting_business_logic_rules_into_the_domain/
---
Typescript is a great and powerful tool for type checking but it could bother you if you don’t define your type according to business logic rules. If your models don’t reflect business logic rules, after some time it will create a gap between what business logic says and how your code behaves. A three part series explores this topic.

* [Part 1](https://medium.com/@mohsen.sareminia/reflecting-business-logic-rules-into-the-domain-models-using-typescript-part-1-e696773ae4ae?source=friends_link&amp;sk=76085024d8b1a731ed47cabb761d0e3a) Describes how to create self-documented types and models.

* [Part 2](https://medium.com/@mohsen.sareminia/reflecting-business-logic-rules-into-the-domain-models-using-typescript-part-2-61a19fba069d?source=friends_link&amp;sk=d596490ba47c075a0648a31294b5fa33) Explains how to write some code that creates and validates those types.

* [Part 3](https://medium.com/@mohsen.sareminia/reflecting-business-logic-rules-into-the-domain-models-using-typescript-part-3-aa8998bc6d29?source=friends_link&amp;sk=bdec956a522701f196bd8a83a5a50838) Uses io-ts library to automate these boring tasks.
## [9][Is there a smarter way to filter out duplicates?](https://www.reddit.com/r/typescript/comments/i29ou1/is_there_a_smarter_way_to_filter_out_duplicates/)
- url: https://www.reddit.com/r/typescript/comments/i29ou1/is_there_a_smarter_way_to_filter_out_duplicates/
---
I want to remove objects from an array if the `playerId` is already contained in that array.

Is there a smarter way to do this in TypeScript than my current implementation?

    private filterDuplicateLeaderBoardLines(result: Result): Result {
      const leaderBoardLines: LeaderBoardLineResult[] = [];
      const playerIds = new Set&lt;string&gt;();
    
      result.sessionResult.leaderBoardLines.forEach((leaderBoardLine: LeaderBoardLineResult) =&gt; {
        const playerId = leaderBoardLine.currentDriver.playerId;
        playerIds.add(playerId);
      });
    
      playerIds.forEach((playerId) =&gt; {
        leaderBoardLines.push(
          result.sessionResult.leaderBoardLines.find(
            (leaderBoardLine) =&gt; leaderBoardLine.currentDriver.playerId === playerId
          )
        );
      });
    
      result.sessionResult.leaderBoardLines = leaderBoardLines;
    
      return result;
    }
## [10][[Question] ReactJS + TypeScript = larger bundle size](https://www.reddit.com/r/typescript/comments/i2ck8c/question_reactjs_typescript_larger_bundle_size/)
- url: https://www.reddit.com/r/typescript/comments/i2ck8c/question_reactjs_typescript_larger_bundle_size/
---
I am looking forward to using TypeScript + ReactJS in my next project. I know TS adds a lot of good stuff into the app but I wonder about the bundled size of the project.

How much bigger in a comparison between a normal ReactJS and ReactJS + TypeScript?

Thank you.
## [11][Help with getting IDE support working for a .d.ts for a JS project, in another JS project?](https://www.reddit.com/r/typescript/comments/i2fqvw/help_with_getting_ide_support_working_for_a_dts/)
- url: https://www.reddit.com/r/typescript/comments/i2fqvw/help_with_getting_ide_support_working_for_a_dts/
---
Situation:

I have a node module, pure node javascript, no bundler, packager, build step etc.  I've written a .d.ts file for it.

I want to use it from another node javascript project, but I want IDE support to work correctly.

Difficulty:  When I do const x=require('lib'); ... vs code tells me that "x" only has one property: default .  If I do 'x.default.' then I get proper intellisense, however, x.default is not how the library is exported, so that code doesn't actually work.

If I change it to "import x from 'lib';" then intellisense works properly, but of course, that's not valid usage from a CJS node package, so that also doesn't work.

Any help how to configure tsconfig / jsconfig / VSC into such a way that this will act how I expect it to?  

Someone has suggested that I add a line
    /** @type {typeof import("lib").default} */
to each place where it's imported, which apparently forces TS to understand it properly... but that's not exactly sustainable across a large project..
