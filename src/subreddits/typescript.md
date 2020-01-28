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
## [2][Understanding any and unknown in TypeScript. Difference between never and void](https://www.reddit.com/r/typescript/comments/euoyos/understanding_any_and_unknown_in_typescript/)
- url: https://wanago.io/2020/01/27/understanding-any-and-unknown-in-typescript-difference-between-never-and-void/
---

## [3][Parcel 2 alpha TypeScript RFC: Parcel 2 alpha's default behavior is to strip out TypeScript with babel and treat it like plain JavaScript. No type checking, no tsconfig.json, etc. Want first-class default support for TypeScript in Parcel 2? Comment in this RFC.](https://www.reddit.com/r/typescript/comments/euu1ml/parcel_2_alpha_typescript_rfc_parcel_2_alphas/)
- url: https://github.com/parcel-bundler/parcel/issues/4022
---

## [4][Reinventing APIs](https://www.reddit.com/r/typescript/comments/euxwv7/reinventing_apis/)
- url: https://www.reddit.com/r/typescript/comments/euxwv7/reinventing_apis/
---
I made a new TypeScript library to replace GraphQL in my projects. I would like to know your sincere opinion about this.

Repo: [https://github.com/yosbelms/remote-func](https://github.com/yosbelms/remote-func)
## [5][A Simplified Jira Clone Built with React/Babel and Node/TypeScript](https://www.reddit.com/r/typescript/comments/euoiek/a_simplified_jira_clone_built_with_reactbabel_and/)
- url: https://github.com/oldboyxx/jira_clone
---

## [6][kahoot](https://www.reddit.com/r/typescript/comments/euz4t2/kahoot/)
- url: https://www.reddit.com/r/typescript/comments/euz4t2/kahoot/
---
Anyone know a kahoot tampermonkey script to use?
## [7][TS2741: Property 'limitToLast' is missing in type 'FirebaseFirestore.Query' but required in type 'firebase.firestore.Query'.](https://www.reddit.com/r/typescript/comments/eua5yw/ts2741_property_limittolast_is_missing_in_type/)
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
## [8][Fast Pipelines with Generators in TypeScript](https://www.reddit.com/r/typescript/comments/etqvgz/fast_pipelines_with_generators_in_typescript/)
- url: https://medium.com/@wim.jongeneel1/fast-pipelines-with-generators-in-typescript-85d285ae6f51
---

## [9][Tool to detect unused class methods?](https://www.reddit.com/r/typescript/comments/etvbdz/tool_to_detect_unused_class_methods/)
- url: https://www.reddit.com/r/typescript/comments/etvbdz/tool_to_detect_unused_class_methods/
---
Hi,I recently discovered `ts-prune`, a tool which detects a subset of deadcode by highlighting any exports that are not used anywhere in the project. This information was a slightly useful to me on my current project, but unfortunately this tool did not highlight unused class methods.

Example:
```
    // src/Foo.ts
    export default class Foo{
        counter: number = 0;
        public incrCounter(){
            this.counter++
        }
        public printCounter(){
            //dead code
            console.log(`The counter is ${counter}`)    
        }
    }
    export function alsoDead(){
        //also dead code
        console.log("Fish")
    }
```
```
    // src/index.ts
    import Foo from './Foo'
    const foo = new Foo();
    foo.incrCounter();
    console.log(foo.counter)
```

In this case `ts-prune` would detect that the exported function `alsoDead` is not used anywhere in the project. However, it would not detect that the method printCounter is never used.

I am vaguely aware that using OOP-flavoured programming with typescript is not fully recommended (?), so I wouldn't be entirely surprised if there was no existing tool that helped solve this problem. 

And if worst comes to worst, I suppose I can resort to deleting each class member in turn and seeing if it fails to compile!!
## [10][Why isn't union type in object field extracted to out of the object?](https://www.reddit.com/r/typescript/comments/etrhcs/why_isnt_union_type_in_object_field_extracted_to/)
- url: https://www.reddit.com/r/typescript/comments/etrhcs/why_isnt_union_type_in_object_field_extracted_to/
---
I think these type should be equal:

    { a: string | undefined }
    { a: string } | { a: undefined }

These mean same set of values. And if these are equals, the following works:

    Exclude&lt;{ a: string | undefined }, { a: string }&gt;   // =&gt; { a: undefined }

But it don't work as I expected. Because `{ a: string | undefined }` is not `{ a: string } | { a: undefined }`.

If it works, Exclude&lt;A, B&gt; could be considered as subtraction of set in mathematical context. I think there is big deals if these types are equals. Why not? And please tell me related GitHub issues if you know. Thanks!
## [11][Why does TS consider this code valid?](https://www.reddit.com/r/typescript/comments/etrrh9/why_does_ts_consider_this_code_valid/)
- url: https://www.reddit.com/r/typescript/comments/etrrh9/why_does_ts_consider_this_code_valid/
---
Hi all,

&amp;#x200B;

I'm a little confused as to why the compiler is letting me get away with the following code. I would normally be expecting a noImplicitAny error inside of the catch param. Please note, this is also happening for the 'then' block - I have just chosen to manually type it.

&amp;#x200B;

I understand on a basic level that's it to do with the return type of request, and that typescript is doing some inference under the covers, but could someone explain why in this case you wouldn't be forced to explicitly type it? Could you explain on a deeper level what is happening? How do I fix this? What other scenarios does this happen in?

&amp;#x200B;

I'm using the axios types inside of their index.d.ts module, so there doesn't appear to be a nice way to get around this. This is obviously a problem in a shared codebase. As you can see in the 'catch' block, I can basically do anything I want at compile time. Obviously at runtime, the response.couldLiterallyHaveAnything here obviously won't exist - and IMO should be something I can pick up at compile time (if you were forced to type it).

&amp;#x200B;

Code sandbox: [https://codesandbox.io/s/typescript-playground-export-lrwfn?fontsize=14&amp;hidenavigation=1&amp;theme=dark](https://codesandbox.io/s/typescript-playground-export-lrwfn?fontsize=14&amp;hidenavigation=1&amp;theme=dark)

&amp;#x200B;

Code snippet:

    import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios';
    
    export const APIService = {
      getAll: async () =&gt; {
        const requestConfig: AxiosRequestConfig = {
          method: 'get',
          url: `someUrl`,
        };
    
        //Why is the compiler letting me have untyped/implicit anys as shown in catch block
    
        return axios
          .request(requestConfig)
          .catch((response) =&gt; Promise.reject(response.couldLiterallyHaveAnythingHere))
          .then((response: AxiosResponse) =&gt; response.data.result);
      },
    };
