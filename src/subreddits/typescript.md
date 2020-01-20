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
## [2][Question, considering learning Typescript](https://www.reddit.com/r/typescript/comments/er8fjv/question_considering_learning_typescript/)
- url: https://www.reddit.com/r/typescript/comments/er8fjv/question_considering_learning_typescript/
---
Hi everyone,

I'm considering learning Typescript as I'm hearing lots of great things about it however my dilemma is that I'm half way through a React Native project, are there any downsides to using Typescript for only part of you code base instead of the entire project?

I should mention that if I like typescript I would consider rewriting the code base overtime

Thanks!
## [3][Getter and Setter on String Array](https://www.reddit.com/r/typescript/comments/er9opl/getter_and_setter_on_string_array/)
- url: https://www.reddit.com/r/typescript/comments/er9opl/getter_and_setter_on_string_array/
---
Hey guys!

Okay, not sure about this:

I am using a simple **getter** and **setter** as shown below:  
If I use different names for the getting, then it works but if I use the same name, then it fails:

`private _person: string[] = [];`  
 `get person(){`  
 `return this._person;`  
  `}`  
 `set person(value: string){`  
 `this._person.push(name);`  
  `}`  


I get the following error:  
 ***Type 'string\[\]' is not assignable to type 'string'.ts(2322)***
## [4][Syntax checking with tide and emacspeak?](https://www.reddit.com/r/typescript/comments/er6345/syntax_checking_with_tide_and_emacspeak/)
- url: https://www.reddit.com/r/typescript/comments/er6345/syntax_checking_with_tide_and_emacspeak/
---
When I use js2-mode and make a syntax error with js2-mode enabled emacspeak will say so, for example if I forget to close a bracket.
This doesn't happen with tide though, I've tried enabling js2-minor-mode too along side tide but no luck.
Does anyone know how to achieve this?
I'm also happy to just use js2-mode but no matter what I do I can't get company-lsp to work properly even with a minimal configuration so I must stick with tide for now
## [5][A Workerpool From Scratch in TypeScript and Node](https://www.reddit.com/r/typescript/comments/eqx58t/a_workerpool_from_scratch_in_typescript_and_node/)
- url: https://medium.com//a-workerpool-from-scratch-in-typescript-and-node-c4352106ffde?source=friends_link&amp;sk=96f50c4fab122cf141287276af7e9ea8
---

## [6][Node.js tips and tricks to make your application more robust and secure. Please share your thoughts?](https://www.reddit.com/r/typescript/comments/eqzvrs/nodejs_tips_and_tricks_to_make_your_application/)
- url: https://www.inkoop.io/blog/nodejs-tips-tricks-that-can-help-you-deliver-more-secure-and-robust-application?source=reddit
---

## [7][Where to see basic docs for types](https://www.reddit.com/r/typescript/comments/eqs63d/where_to_see_basic_docs_for_types/)
- url: https://www.reddit.com/r/typescript/comments/eqs63d/where_to_see_basic_docs_for_types/
---
Where are the docs for stuff like @types/express or @types/pg
I know the JavaScript docs for the most part but not the names of the specific types.
Is looking in the specific files in node_modules the only way to know?
## [8][Trouble finding this return type](https://www.reddit.com/r/typescript/comments/eqtiha/trouble_finding_this_return_type/)
- url: https://www.reddit.com/r/typescript/comments/eqtiha/trouble_finding_this_return_type/
---
There are 3 modules I'm working with: firebase, firebase-admin, and firebase-testing. admin handles authentication and has access to some classes and types. The core module does too.

Testing is needed for testing but is making things extra complicated at the moment.

Take this example:

    import * as firebase from "firebase"; // also tried without * aliasing
    import {firestore} from "firebase-admin";
    import firebaseTesting from "@firebase/testing";
    
    // https://firebase.google.com/docs/reference/js/firebase.firestore.Firestore
    // TS2740: Type '(credentials: any) =&gt; Firestore' is missing the following properties from type 'Firestore': settings, collection, doc, collectionGroup, and 4 more.
    const createTestDatabase: firestore.Firestore = credentials =&gt; {
      return firebaseTesting
        .initializeTestApp({
          projectId: 'testProject',
          auth: credentials
        })
        .firestore();
    };
    
    // Another try...
    // https://firebase.google.com/docs/reference/js/firebase.html#initializeapp
    // TS2740: Type '(credentials: any) =&gt; Firestore' is missing the following properties from type 'App': auth, database, delete, installations, and 8 more.
    const createTestDatabase: firebase.app.App = credentials =&gt; {
      return firebaseTesting
        .initializeTestApp({
          projectId: 'testProject',
          auth: credentials
        })
        .firestore();
    };

When faced with tough typings do you guys just do the easy thing and mark it as "any" or is there a process you follow to find the right typing? I tried following the docs to find the right return type but the two above haven't worked.

&amp;#x200B;

PS: What I would really like is if the IDE could help me deduce the correct type. I hear VSCode outright tells you. I'm using Phpstorm. It seems to know the shape but 1) doesn't give the full member list 2) doesn't make any type suggestions even though Phpstorm scans all libraries deeply.
## [9][Become a Better Software Architect](https://www.reddit.com/r/typescript/comments/eqinar/become_a_better_software_architect/)
- url: https://github.com/justinamiller/SoftwareArchitect
---

## [10][From flow to monorepo in Typescript](https://www.reddit.com/r/typescript/comments/eqg5w3/from_flow_to_monorepo_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/eqg5w3/from_flow_to_monorepo_in_typescript/
---
Hello, I have recently moved my multiple repositories javascript packages to a single monorepo in typescript. It's my first-time complete Typescript project and would love it if you had some time to give me some feedback. 
I had a lot of trouble finding some types of external libraries. Does anyone have experience with typing webpack and gulp plugin?

[Repository](https://github.com/FullHuman/purgecss)
## [11][Is it possible to Compile and Update the Browser at the same time?](https://www.reddit.com/r/typescript/comments/eqcv0e/is_it_possible_to_compile_and_update_the_browser/)
- url: https://www.reddit.com/r/typescript/comments/eqcv0e/is_it_possible_to_compile_and_update_the_browser/
---
Hello:  
I am using **lite server** to watch and run my code.  
However, each time I make a code change to a typescript file nothing happens, the only time the browser refreshes is when I make changes to the index.html file.

To see any updates I have to stop **lite server,**  compiled the code and then run the server again.

I attempted watching the code with **tsc --watch** but I have to run the code on a server.  


Are there solutions where you can run a server which compile on any code change and update the browser?
