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
## [2][Using redux-saga with Typescipt](https://www.reddit.com/r/typescript/comments/evg6rd/using_reduxsaga_with_typescipt/)
- url: https://www.reddit.com/r/typescript/comments/evg6rd/using_reduxsaga_with_typescipt/
---
I've been using `redux-thunk` for quite a while on a project of mine, but I've recently outgrown it and have turned to other libraries to handle async actions (mostly network requests). Of the options out there, I've found `redux-saga` and `redux-observable` to be the most promising options. After some research, I found that `redux-saga` seemed the best for my use-case. 

I then installed saga and the accompanying types, but I quickly found that saga's integration with TypeScript was less than ideal. For example, using a selector returns an `any` type as opposed to the proper type which corresponds to the output of the selector.

I did some follow-up research on this topic, and from what I can see, saga doesn't support this kind of type safety -- something that `redux-observables` does support. Is there some kind of work-around or alternate `@types` library which allows saga to support real type safety with TypeScript? 

I was excited to use saga, but if proper type checks do not properly work, I will be forced to use `redux-observable`.
## [3][Searching for JavaScript/TypeScript senior dev who can share some Best Practices](https://www.reddit.com/r/typescript/comments/evbddt/searching_for_javascripttypescript_senior_dev_who/)
- url: https://www.reddit.com/r/typescript/comments/evbddt/searching_for_javascripttypescript_senior_dev_who/
---
I am searching for JavaScript/TypeScript senior dev who can share some Best Practices and answer some of my questions. I can surely pay you for it :)

Topics:

1. Runtime type checking, best practices for JSOn parsing
2. Shared code for server/client apps (interfaces, enums, locales)

Medium (or similar) articles are too basic...
## [4][Get available properties of interface without initializing](https://www.reddit.com/r/typescript/comments/ev806n/get_available_properties_of_interface_without/)
- url: https://stackoverflow.com/questions/59953416/get-available-properties-of-interface-without-initializing
---

## [5][TypeScript advice](https://www.reddit.com/r/typescript/comments/ev6sd5/typescript_advice/)
- url: https://medium.com/@Armandotrue/typescript-advice-4a46b6ae5a7
---

## [6][Understanding any and unknown in TypeScript. Difference between never and void](https://www.reddit.com/r/typescript/comments/euoyos/understanding_any_and_unknown_in_typescript/)
- url: https://wanago.io/2020/01/27/understanding-any-and-unknown-in-typescript-difference-between-never-and-void/
---

## [7][Reinventing APIs](https://www.reddit.com/r/typescript/comments/euxwv7/reinventing_apis/)
- url: https://www.reddit.com/r/typescript/comments/euxwv7/reinventing_apis/
---
I made a new TypeScript library to replace GraphQL in my projects. I would like to know your sincere opinion about this.

Repo: [https://github.com/yosbelms/remote-func](https://github.com/yosbelms/remote-func)
## [8][Parcel 2 alpha TypeScript RFC: Parcel 2 alpha's default behavior is to strip out TypeScript with babel and treat it like plain JavaScript. No type checking, no tsconfig.json, etc. Want first-class default support for TypeScript in Parcel 2? Comment in this RFC.](https://www.reddit.com/r/typescript/comments/euu1ml/parcel_2_alpha_typescript_rfc_parcel_2_alphas/)
- url: https://github.com/parcel-bundler/parcel/issues/4022
---

## [9][A Simplified Jira Clone Built with React/Babel and Node/TypeScript](https://www.reddit.com/r/typescript/comments/euoiek/a_simplified_jira_clone_built_with_reactbabel_and/)
- url: https://github.com/oldboyxx/jira_clone
---

## [10][TS2741: Property 'limitToLast' is missing in type 'FirebaseFirestore.Query' but required in type 'firebase.firestore.Query'.](https://www.reddit.com/r/typescript/comments/eua5yw/ts2741_property_limittolast_is_missing_in_type/)
- url: https://www.reddit.com/r/typescript/comments/eua5yw/ts2741_property_limittolast_is_missing_in_type/
---
    // https://firebase.google.com/docs/reference/js/firebase.firestore.CollectionReference.html#where
    where
    where(fieldPath: string | FieldPath, opStr: WhereFilterOp, value: any): Query&lt;T&gt;
    
    // https://firebase.google.com/docs/reference/js/firebase.firestore.Query.html
    firebase. firestore. Query &lt; T &gt;
    A Query refers to a Query which you can read or listen to. You can also construct refined Query objects by adding filters and ordering.

I think .where() should return firebase.firestore.Query based on the above. Can someone tell me why this is not the case or what else may be causing the type error below?

    import {DocumentTarget, RecordTypes, WhereDefinition} from "../../../shared";
    import admin, {firestore} from "firebase-admin";
    import firebase from "firebase";
    
    export default class FirestoreConnection {
      public async getSome(collection: string, whereClauses: WhereDefinition[], limit: number, orderBy: string, offset: number): Promise&lt;object|boolean&gt; {
    // TS2741: Property 'limitToLast' is missing in type 'FirebaseFirestore.Query' but required in type 'firebase.firestore.Query'.  index.d.ts(8156, 5): 'limitToLast' is declared here.
    
        const baseQuery: firebase.firestore.Query = this.database
          .collection(collection)
          .where("storeId", "==", this.shopDomain);
## [11][Fast Pipelines with Generators in TypeScript](https://www.reddit.com/r/typescript/comments/etqvgz/fast_pipelines_with_generators_in_typescript/)
- url: https://medium.com/@wim.jongeneel1/fast-pipelines-with-generators-in-typescript-85d285ae6f51
---

