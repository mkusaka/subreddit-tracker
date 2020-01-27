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
## [2][TS2741: Property 'limitToLast' is missing in type 'FirebaseFirestore.Query' but required in type 'firebase.firestore.Query'.](https://www.reddit.com/r/typescript/comments/eua5yw/ts2741_property_limittolast_is_missing_in_type/)
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
## [3][Fast Pipelines with Generators in TypeScript](https://www.reddit.com/r/typescript/comments/etqvgz/fast_pipelines_with_generators_in_typescript/)
- url: https://medium.com/@wim.jongeneel1/fast-pipelines-with-generators-in-typescript-85d285ae6f51
---

## [4][Tool to detect unused class methods?](https://www.reddit.com/r/typescript/comments/etvbdz/tool_to_detect_unused_class_methods/)
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
## [5][Why isn't union type in object field extracted to out of the object?](https://www.reddit.com/r/typescript/comments/etrhcs/why_isnt_union_type_in_object_field_extracted_to/)
- url: https://www.reddit.com/r/typescript/comments/etrhcs/why_isnt_union_type_in_object_field_extracted_to/
---
I think these type should be equal:

    { a: string | undefined }
    { a: string } | { a: undefined }

These mean same set of values. And if these are equals, the following works:

    Exclude&lt;{ a: string | undefined }, { a: string }&gt;   // =&gt; { a: undefined }

But it don't work as I expected. Because `{ a: string | undefined }` is not `{ a: string } | { a: undefined }`.

If it works, Exclude&lt;A, B&gt; could be considered as subtraction of set in mathematical context. I think there is big deals if these types are equals. Why not? And please tell me related GitHub issues if you know. Thanks!
## [6][Why does TS consider this code valid?](https://www.reddit.com/r/typescript/comments/etrrh9/why_does_ts_consider_this_code_valid/)
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
## [7][TypeScript: X doesn't not exist on type Y](https://www.reddit.com/r/typescript/comments/etmblc/typescript_x_doesnt_not_exist_on_type_y/)
- url: https://www.reddit.com/r/typescript/comments/etmblc/typescript_x_doesnt_not_exist_on_type_y/
---
 

I Tried to use nodemon for auto restart server on changes but  Typescript throws the following error even though extended the type.  Please find the code attached below.

Typings work fine with VS Code but the errors while compiling by nodemon. Error attached below code. 

index.ts

    import * as fastify from 'fastify'; 
    import { Server, IncomingMessage, ServerResponse } from "http"; 
    const server: FastifyInstance&lt;Server, IncomingMessage, ServerResponse&gt; = fastify(); server.ready(err =&gt; { if (err) throw err;     server.swagger(); }); 
    const start = async () =&gt; { 
        try { 
            await server.listen(3000, "0.0.0.0");         
            server.log.info("Server running on 3000");         
            console.log("Server running on 3000"); 
        } catch (error) {         
            server.log.error(error); 
        } 
    }  
    start();

src/typings/fastify.d.ts

    import * as fastify from 'fastify'; 
    import * as http from "http";  
    declare module 'fastify' { 
        export interface FastifyInstance&lt; HttpServer = http.Server, HttpRequest = http.IncomingMessage, HttpResponse = http.ServerResponse &gt; {         
        swagger(): void 
        } 
    }

tsconfig.json

    { 
        "compileOnSave": false, 
        "compilerOptions": { 
            "baseUrl": ".", 
            "lib": [ "es2016" ], 
            "paths": { "*": [ "src/*" ] }, 
            "allowJs": true, 
            "target": "es2016", 
            "module": "commonjs", 
            "moduleResolution": "node", 
            "sourceMap": true, 
            "outDir": "./dist", 
            "typeRoots": [ "src/typings", "node_modules/@types" ] 
        }, 
        "include": [ "./src/**/*.ts" ], 
        "exclude": [ "node_modules" ], 
        "rules": {} 
    } 

[Error Thrown while compiling with NodeMon](https://i.stack.imgur.com/o71Xn.png)
## [8][Question: React onPress function type](https://www.reddit.com/r/typescript/comments/etdhan/question_react_onpress_function_type/)
- url: https://www.reddit.com/r/typescript/comments/etdhan/question_react_onpress_function_type/
---
Hey!

I just ported my react-native project to typescript and have a question about functions as props

im passing:

`&lt;DisplayCardsWithLikesdata={testData}likes={500}`**onPress={() =&gt; this.props.navigation.navigate("CardDetailScreen")}**`/&gt;`

&amp;#x200B;

to

`type Props = {onPress: Function}`

&amp;#x200B;

`const FloatingActionButtonSimple = (props:Props) =&gt; {const {onPress} = propsreturn (&lt;View style={styles.containerFab}&gt;&lt;TouchableOpacity style={styles.fab} onPress={onPress}&gt;&lt;Icon name="plus" size={16} color={"white"} /&gt;&lt;/TouchableOpacity&gt;&lt;/View&gt;);};`

&amp;#x200B;

So nothing crazy, but if i define onPress as a function it doesn't work. Does anybody of you have an idea on how to define onPress?

Thanks a lot!
## [9][Proper syntax for having a numeric variable in a string](https://www.reddit.com/r/typescript/comments/etb08f/proper_syntax_for_having_a_numeric_variable_in_a/)
- url: https://www.reddit.com/r/typescript/comments/etb08f/proper_syntax_for_having_a_numeric_variable_in_a/
---
So I have variables declared and everything, and I want to add it to a printed string. I want to have something like "Your X plants have been watered", and X is a variable rather than a string. I know I would have print("Your X plants have been watered") if I wanted it to say X, but how do I make it so its a numeric variable?
## [10][Set Up a Typescript React Redux Project](https://www.reddit.com/r/typescript/comments/et3gev/set_up_a_typescript_react_redux_project/)
- url: https://typeofnan.dev/setup-a-typescript-react-redux-project/
---

## [11][Is schemats, still maintained](https://www.reddit.com/r/typescript/comments/et56d4/is_schemats_still_maintained/)
- url: https://www.reddit.com/r/typescript/comments/et56d4/is_schemats_still_maintained/
---
On the GitHub the last commit was almost 2 years ago, does it work well?
Also does it work fine with both pg and pg-promise?

https://github.com/SweetIQ/schemats
