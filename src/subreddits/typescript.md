# typescript
## [1][Who's hiring Typescript developers - January](https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/)
- url: https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][Announcing TypeScript 3.8 Beta](https://www.reddit.com/r/typescript/comments/eph9wx/announcing_typescript_38_beta/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-3-8-beta/
---

## [3][Foal TS - January release (version 1.5) - LinkedIn and Github social support, customizable exception handler, higher CSRF protection](https://www.reddit.com/r/typescript/comments/ephuc9/foal_ts_january_release_version_15_linkedin_and/)
- url: https://www.reddit.com/r/typescript/comments/ephuc9/foal_ts_january_release_version_15_linkedin_and/
---
Foal TS framework version 1.5 is officially released!

LinkedIn and Github social support, customizable exception handler, higher CSRF protection, discover the new features of version 1.5 here: [https://github.com/FoalTS/foal/releases/tag/v1.5.0](https://github.com/FoalTS/foal/releases/tag/v1.5.0)

&amp;#x200B;

[Foal TS - January release \(version 1.5\)](https://preview.redd.it/jlt8b9k9g4b41.png?width=1054&amp;format=png&amp;auto=webp&amp;s=3b7b4c2bd0897a5e7ea14e6bfdd5a1bdce2d79b8)

Foal, in a few words, it's a Node.js framework:

* written in TypeScript
* supplied with batteries included (Auth, OpenAPI, GraphQL, ORM, CLI, ...)
* and with a simple and intuitive architecture (no over-engineering).

And the must: it has more than 11,000 lines of documentation.

[Example of error handling](https://preview.redd.it/mzo3jo2cg4b41.png?width=1804&amp;format=png&amp;auto=webp&amp;s=df1aff998feef849ea8a25eefe32adcc1c539ca9)

[https://foalts.org](https://foalts.org/)

[https://github.com/FoalTS/foal](https://github.com/FoalTS/foal)

[https://foalts.gitbook.io/docs/](https://foalts.gitbook.io/docs/)

\#TypeScript #NodeJS #javascript #LinkedIn #Github #webdevelopment
## [4][[Help] ECMA6 + VScode](https://www.reddit.com/r/typescript/comments/epj6gh/help_ecma6_vscode/)
- url: https://www.reddit.com/r/typescript/comments/epj6gh/help_ecma6_vscode/
---
I started to test my logic skills in CodeWars, but CodeWars don't accept latest version of ecma. There is a way to my eslint or typescript shows me a warning or error when I write a function or somethint that doesn't exist or is not accepted in ecma6?

&amp;#x200B;

To answer curiosity I'm using VSCODE.
## [5][Is there a way to deconstruct params AND define the interface in a single call?](https://www.reddit.com/r/typescript/comments/epen4f/is_there_a_way_to_deconstruct_params_and_define/)
- url: https://www.reddit.com/r/typescript/comments/epen4f/is_there_a_way_to_deconstruct_params_and_define/
---
No destructuring

```

function printName(param: {firstName: string, lastName: string}) {

console.log(param.firstName + ' ' + param.lastName)

}

```

&amp;#x200B;

Define interface

```interface Name {

firstName: string,

lastName: string

}

function printName(param: Name) {

console.log(param.firstName + ' ' + param.lastName)

}

```

&amp;#x200B;

Destructuring + Inline Interface

```

function printName({ firstName, lastName }: { firstName: string; lastName: string }) {

console.log(firstName + ' ' + lastName)

}

```

&amp;#x200B;

Now I'm wondering, is there a way to have destructuring so that I don't need to call param.&lt;something&gt;, but at the same time having it inline? Perhaps something like this:

```

function printName({ firstName: string; lastName: string }) {

console.log(firstName + ' ' + lastName)

}

```

Destructuring + Inline Interface is kinda annoying because you need to redeclare everything twice. Is there anything I can write that does destructuring but also defines the type at the same time?
## [6][I created a Fitbit app using TypeScript](https://www.reddit.com/r/typescript/comments/ep2rye/i_created_a_fitbit_app_using_typescript/)
- url: https://www.reddit.com/r/typescript/comments/ep2rye/i_created_a_fitbit_app_using_typescript/
---
As the title says, I want to share [this small Fitbit app](https://github.com/Shpota/zeit) with you. I posted it in this sub when I initially created it, but since then I changed the app, applied improvements, released it to the store, and got feedback from users. 

It is my pet project and I don't use TypeScript at work that's why I was glad that I could implement applications like this with TypeScript.

I would love to get any kind of feedback

https://github.com/Shpota/zeit
## [7][One hidden benefit of TS I haven't seen discussed](https://www.reddit.com/r/typescript/comments/ep8h1v/one_hidden_benefit_of_ts_i_havent_seen_discussed/)
- url: https://www.reddit.com/r/typescript/comments/ep8h1v/one_hidden_benefit_of_ts_i_havent_seen_discussed/
---
Example documentation:

    get(options?: GetOptions): Promise&lt;DocumentSnapshot&lt;T&gt;&gt;

In all the courses I took and all the toy projects I built my first 2 years, none of them ever explained what the above was. I sort of understood what something like `.set(path [, options])` after much time and frustration. But the above was another level.

Once could guess the meaning of ? and probably pick up the type hinting after learning a 2nd language, but that final bit with the Promise and generic was always a source of pain and confusion.

Feels like a large gap in the current educational landscape to be honest. TS fills it perfectly because by some magical coincidence it seems to match most complex documentation's syntax exactly.

Finally, I can read!
## [8][What is the correct way to do type checking here](https://www.reddit.com/r/typescript/comments/ep55ak/what_is_the_correct_way_to_do_type_checking_here/)
- url: https://www.reddit.com/r/typescript/comments/ep55ak/what_is_the_correct_way_to_do_type_checking_here/
---
I thought I would know this, having worked with TS for 2 months now. But I'm running into issues:

    export enum RuleTypes {global, inventory, product, blacklist}
    
    export interface GlobalRuleSchema {
      storeId: string
      , globalPercent : number
    }
    
    // inside a class
      verifySchemaIsCorrect(type: RuleTypes, recordData: unknown) {
        let schemaMatches: boolean;
        
        switch (+type) { // + needed for number coercion as per SO
          case type.global: // TS2339: Property 'global' does not exist on type 'RuleTypes'.
            schemaMatches = typeof recordData === GlobalRuleSchema; // TS2693: 'GlobalRuleSchema' only refers to a type, but is being used as a value here
            break;
          default:
            schemaMatches = false;
        }    
    
        if (Boolean(schemaMatches)) {
          return true;
        }
        
        throw Error("recordData's shape is wrong against its schema");
      }

What is the correct way to ensure the interface on `recordData` matches the desired enum value on `type` ?
## [9][Hopa - zero config CLI for runs JavaScript and TypeScript](https://www.reddit.com/r/typescript/comments/ep18kf/hopa_zero_config_cli_for_runs_javascript_and/)
- url: https://krasimirtsonev.com/blog/article/hopa-javascript-typescript-runner
---

## [10][How to group related interfaces?](https://www.reddit.com/r/typescript/comments/ep13fr/how_to_group_related_interfaces/)
- url: https://www.reddit.com/r/typescript/comments/ep13fr/how_to_group_related_interfaces/
---
I'm creating some schemas for Cloud Firestore and was curious if there are best practices for grouping interfaces. 

I don't think they can be grouped inside a class so is there any other way to encapsulate them together? 

Also, at the file level what do you guys do?

I'm currently putting them in /shared/interfaces.ts. My project's scope is relatively small and has both the back and front end in other folders. I thought it would be good for both ends to share the interface file so that they adhere to the same type definitions.
## [11][TypeScript Tip of the Week](https://www.reddit.com/r/typescript/comments/epb8ur/typescript_tip_of_the_week/)
- url: https://medium.com/@sredmond/typescript-tip-of-the-week-using-classes-interfaces-6a0570f46750
---

