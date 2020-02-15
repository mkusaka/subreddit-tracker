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
## [2][How to refactor this code using Optional Chaining?](https://www.reddit.com/r/typescript/comments/f48m23/how_to_refactor_this_code_using_optional_chaining/)
- url: https://www.reddit.com/r/typescript/comments/f48m23/how_to_refactor_this_code_using_optional_chaining/
---
`res.data.items ? res.data.items[0].id : undefined`

I've tried this but didn't work

`res.data.items[0]?.id`

`res.data.items?[0].id`
## [3][I have created a new TypeScript library to handle common WebSocket problems while keeping the size of the library to be as small as possible. Let me know what you think.](https://www.reddit.com/r/typescript/comments/f3svfg/i_have_created_a_new_typescript_library_to_handle/)
- url: https://github.com/mat-sz/typesocket
---

## [4][Store - A beautifully-simple framework-agnostic modern state management library](https://www.reddit.com/r/typescript/comments/f3vsz9/store_a_beautifullysimple_frameworkagnostic/)
- url: https://github.com/fabiospampinato/store
---

## [5][What tools do you use to generate docs for TS project?](https://www.reddit.com/r/typescript/comments/f3yb4v/what_tools_do_you_use_to_generate_docs_for_ts/)
- url: https://www.reddit.com/r/typescript/comments/f3yb4v/what_tools_do_you_use_to_generate_docs_for_ts/
---
The one I used in my projects are:
- typedoc: the ugly one, but does the job
- storybook: only for react components 

Are there any other options? Are there options that handle React and normal TS well?
## [6][GitHub: run-exclusive:⚡🔒Generate functions that do not allow parallel executions🔒 ⚡](https://www.reddit.com/r/typescript/comments/f3wakk/github_runexclusivegenerate_functions_that_do_not/)
- url: https://github.com/garronej/run-exclusive
---

## [7][Doorbot.ts – Let me in the building, with Typescript](https://www.reddit.com/r/typescript/comments/f3hggg/doorbotts_let_me_in_the_building_with_typescript/)
- url: http://www.wumpus-cave.net/2020/02/13/doorbot-ts-let-me-in-the-building-with-typescript/
---

## [8][extract discriminants into type](https://www.reddit.com/r/typescript/comments/f3awxt/extract_discriminants_into_type/)
- url: https://www.reddit.com/r/typescript/comments/f3awxt/extract_discriminants_into_type/
---
Hi, i need a new type that is a union of all the disciminants in a discriminated uniton.

    interface Square {
        kind: "square";
        size: number;
    }
    
    interface Rectangle {
        kind: "rectangle";
        width: number;
        height: number;
    }
    
    interface Circle {
        kind: "circle";
        radius: number;
    }
    
    type Shape = Square | Rectangle | Circle;
    
    type KindList = SomeMagic&lt;Shape&gt; // "square" | "rectangle" | "circle"

can someone help me?
## [9][Just blogged: TypeScript types cheat sheet](https://www.reddit.com/r/typescript/comments/f33on2/just_blogged_typescript_types_cheat_sheet/)
- url: https://medium.com/@maximzhukov_dev/typescript-types-cheat-sheet-5da7c6ee083d
---

## [10][Unit tests for d.ts files in JS project](https://www.reddit.com/r/typescript/comments/f2wbaf/unit_tests_for_dts_files_in_js_project/)
- url: https://github.com/ai/check-dts
---

## [11][Angular 9: What’s New?](https://www.reddit.com/r/typescript/comments/f2tsre/angular_9_whats_new/)
- url: https://auth0.com/blog/angular-9-whats-new/?utm_source=twitter&amp;utm_medium=sc&amp;utm_campaign=angular_9
---

