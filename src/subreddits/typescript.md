# typescript
## [1][Who's hiring Typescript developers - June](https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/)
- url: https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/
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
## [2][My validation library written with type inference in mind](https://www.reddit.com/r/typescript/comments/gx6c8l/my_validation_library_written_with_type_inference/)
- url: https://www.reddit.com/r/typescript/comments/gx6c8l/my_validation_library_written_with_type_inference/
---
Hi all!

I just finished work on README for my first user input validation library.

This is still in alpha but I already use it in my production projects and I hope it would be useful for you too.

This is my first step to open source, so feel free to comment, open issues and create pull requests.

Any critics is welcome.

I know there are many user input validating libraries, but I did not found one with type inference. And type inference is a one of the greatest features of TypeScript.

NPM: [https://www.npmjs.com/package/treat-like](https://www.npmjs.com/package/treat-like)

GitHub: [https://github.com/atomAltera/treat-like](https://github.com/atomAltera/treat-like)
## [3][Looking for something more advanced? Learn Data Structures and Algorithms in Typescript](https://www.reddit.com/r/typescript/comments/gx67lc/looking_for_something_more_advanced_learn_data/)
- url: https://www.reddit.com/r/typescript/comments/gx67lc/looking_for_something_more_advanced_learn_data/
---
https://github.com/jeffzh4ng/dsa-ts

MIT/Harvard/Stanford's online DS/algos courses are great for theory and understanding. Real life application and real code matters too though. You learn best when you actually implement these DSA for yourself.

This repository is a collection of DS/algorithms implemented in Typescript. If you also need video lectures, they're there too.

I finished the first section of data structures. sequences.

- static/dynamic arrays
- linked lists
- stacks
- queues
- double-ended queues
- circular buffers

Next week we'll cover the Priority Queue ADT and Heaps.

- binary heaps
- indexed binary heaps
- binomial heaps
- fibonnaci heaps
- pairing heap
- soft heap

The repository and series is just getting started! The plan is to go over all classic data structures. Hashing, self balancing trees, tries, graphs.

And algorithms: sorting, searching, backtracking, dynamic, greedy, graph theory, minimum spanning trees, and more.

Videos and code here: https://github.com/jeffzh4ng/dsa-ts
## [4][How to indicate error status in a return value.](https://www.reddit.com/r/typescript/comments/gxcrh5/how_to_indicate_error_status_in_a_return_value/)
- url: https://www.reddit.com/r/typescript/comments/gxcrh5/how_to_indicate_error_status_in_a_return_value/
---
I am learning typescript, and I don't really have a mentor. I am struggling with a lot of new concepts. 

I have a module that exports a function that returns a result. I want to indicate in the return value if there was an error, and provide feedback to be relayed to the user.

Is this a good way to implement this?

`type ProcessResult = { error: false; results: number[] } | { error: true; message: string };`  


`export function processFile(filePath: string): ProcessResult {`      
`// process the file, and produce the result, or return a descriptive error message`  
`}`

My gut instinct is it's better to check a boolean in the return value, than check typeof the return value. But maybe that's not the best way to do it. Also, that may be because I'm coming from Python, where checking the type of anything is considered bad form.
## [5][Use types or interfaces to constrain generics?](https://www.reddit.com/r/typescript/comments/gxertm/use_types_or_interfaces_to_constrain_generics/)
- url: https://www.reddit.com/r/typescript/comments/gxertm/use_types_or_interfaces_to_constrain_generics/
---
I thought interfaces should be limited to class implementations. However:

    // https://www.typescriptlang.org/docs/handbook/generics.html
    
    interface Lengthwise {
        length: number;
    }
    
    function loggingIdentity&lt;T extends Lengthwise&gt;(arg: T): T {
        console.log(arg.length);  // Now we know it has a .length property, so no more error
        return arg;
    }

I checked this code and it also compiles.

    type Lengthwise = {
        length: number;
    }
    
    function loggingIdentity&lt;T extends Lengthwise&gt;(arg: T): T {
        console.log(arg.length); 
        return arg;
    }

Can you guys tell me which you would prefer? I would think the type definition should almost always be preferred unless the generic is being used to create a class factory.
## [6][Accessing an item in an Object that is possibly null](https://www.reddit.com/r/typescript/comments/gxbeea/accessing_an_item_in_an_object_that_is_possibly/)
- url: https://www.reddit.com/r/typescript/comments/gxbeea/accessing_an_item_in_an_object_that_is_possibly/
---
Hi, i'm currently testing out firebase and its Google Authentication ..  


I'm using firebase and firebaseui in the dependencies.

&amp;#x200B;

in my App.tsx i have my authenticated state that is set via onAuthStateChanged(() =&gt; {....} in useEffect

That state i pass down to the child component AuthScreen. In the AuthScreen i want to display the username from firebase with firebase.auth().currentUser.displayName .. but heres the problem.. typescript complains that currentUser is possibly null so i cant access displayName. And i thought i can check this with firebase.auth().currentuser?.displayName or at lease with if(....currentUser !== null) .. but both doesn' work.. does someone has any idea? I would appreciate! 

&amp;#x200B;

Here is some code example from my project: 

`function App() {`  
 `const [authenticated, setAuthenticated] = useState&lt;boolean&gt;(false);`  
 `const handleLogout = () =&gt; {`  
 `firebase.auth().signOut();`  
  `};`  
 `useEffect(() =&gt; {`  
 `firebase.auth().onAuthStateChanged(user =&gt; {`  
 `setAuthenticated(!!user);`  
`});`  
  `}, []);`  
 `return (`  
 `&lt;AppContextProvider value={{ authenticated: !authenticated }}&gt;`  
 `&lt;Styled.LoginBody&gt;`  
 `&lt;Suspense fallback={&lt;div&gt;Loading...&lt;/div&gt;}&gt;`  
 `&lt;AuthScreenComponent`  
 `authenticated={authenticated}`  
 `handleLogout={handleLogout}`  
 `/&gt;`  
 `&lt;/Suspense&gt;`  
 `&lt;/Styled.LoginBody&gt;`  
 `&lt;/AppContextProvider&gt;`  
  `);`  
`}`

&amp;#x200B;

`interface IProps {`  
 `authenticated: boolean;`  
 `handleLogout: () =&gt; void;`  
`}`  
`function AuthScreen(props: IProps) {`  
 `const context = useContext(AppContext);`  
 `const uiConfig = {`  
 `signInFlow: 'popup',`  
 `signInOptions: [firebase.auth.GoogleAuthProvider.PROVIDER_ID],`  
 `callbacks: {`  
 `// Avoid redirects after sign-in.`  
 `signInSuccessWithAuthResult: () =&gt; false`  
`}`  
  `};`  
 `return (`  
 `&lt;Styled.LoginForm&gt;`  
 `&lt;&gt;`  
 `{props.authenticated ? (`  
 `&lt;&gt;`  
 `&lt;h1&gt;`  
`Welcome{' '}`  
 `{console.log(`  
 `firebase.auth().currentUser !== null`  
`? firebase.auth().currentUser.displayName`  
`: console.log('user null')`  
`)}`  
 `&lt;/h1&gt;`  
 `&lt;button onClick={props.handleLogout}&gt;Logout&lt;/button&gt;`  
 `&lt;/&gt;`  
`) : (`  
 `&lt;StyledFirebaseAuth`  
 `uiConfig={uiConfig}`  
 `firebaseAuth={firebase.auth()}`  
 `/&gt;`  
`)}`  
 `&lt;/&gt;`  
 `&lt;/Styled.LoginForm&gt;`  
  `);`  
`}`

&amp;#x200B;

https://preview.redd.it/tj1zijfe95351.png?width=1080&amp;format=png&amp;auto=webp&amp;s=3e9e669bc23511b076fca9fa8d9ff6455e388c1c
## [7][Trying to reset a key value pair in an object where the value is a type to a string](https://www.reddit.com/r/typescript/comments/gxdqin/trying_to_reset_a_key_value_pair_in_an_object/)
- url: https://www.reddit.com/r/typescript/comments/gxdqin/trying_to_reset_a_key_value_pair_in_an_object/
---
Sorry if the title is confusing, but I have a set of classes that look like this.

    export class Account {
    	id = "";
    	name? = "";
    }
    
    export class Service {
    	id = "";
    	name? = "";
    }
    
    export class LineItem {
    	account: Account = new Account();
    	service: Service = new Service();
    	type = "";
    	amount = 0;
    }

I populate LineItem and need to send it to my API. Now my issue is that my API needs `account` and `service` to be strings (the id of the type), not a type. But my state management needs to store these as types (using Vuejs &amp; Vuex). My thought process here is to store/commit the objects with the types to the state, and then right before I send this value to API, set `account` and `service` to be the value of the `id`. instead of the objects.. Any thoughts? I can try explaining a little better if needed too
## [8][Good typescript function programming library](https://www.reddit.com/r/typescript/comments/gwvw9q/good_typescript_function_programming_library/)
- url: https://www.reddit.com/r/typescript/comments/gwvw9q/good_typescript_function_programming_library/
---
Hello, I'm finding a good functional programming library for typescript, currently I found [purify](https://gigobyte.github.io/purify) and [fp-ts](https://gcanti.github.io/fp-ts) and wondering which one I should learn or are there any other libraries that worth looking at. Oh and I'm just finding a library to learn functional programming, not working on any projects. Thank you.
## [9][Getting to a maintainable + enjoying code architecture](https://www.reddit.com/r/typescript/comments/gx6xeo/getting_to_a_maintainable_enjoying_code/)
- url: https://www.reddit.com/r/typescript/comments/gx6xeo/getting_to_a_maintainable_enjoying_code/
---
Learned today from `type-graphql`'s maintainer [https://github.com/MichalLytek/type-graphql/issues/646#issuecomment-639537771](https://github.com/MichalLytek/type-graphql/issues/646#issuecomment-639537771) how a flexible reference architecture for typical GraphQL/CRUD apps could look like. Feedback/improvements are welcome!

https://preview.redd.it/u6uw2tf434351.png?width=1248&amp;format=png&amp;auto=webp&amp;s=eb18091c549398fd5dfb9825660322e67f7dd270
## [10][Is it possible to use path alias for a multi-repo project?](https://www.reddit.com/r/typescript/comments/gx4xmc/is_it_possible_to_use_path_alias_for_a_multirepo/)
- url: https://www.reddit.com/r/typescript/comments/gx4xmc/is_it_possible_to_use_path_alias_for_a_multirepo/
---
So, I'll try to explain the scenario the best I possibly can, and the reason why I'm here.

We will launch a new site pretty soon, we have an old legacy site that we're replacing completely with never technologies, etc.

My first idea was to create a Lerna mono-repo using yarn workspaces, but that didn't work out because Nuxt doesn't work that well with Yarn Workspaces (which we needed).

So I ditched that idéa and instead went with multiple repositories, basically migrated package/folder into its own repo.

So, now we have 3 different repositories:  
\* SCL - Component lib (Shared Vue components)  
\* App-A (Nuxt app)  
\* App-B (Vue app)

The SCL repo takes advantage of the path alias feature of TS, so the imports are readable, something like the following:

    {
       "paths": {
          "@/*": ["./src/*"]
       }
    }

Meaning I can reference components like the following inside the project.  
`import MyComponent from '@/components/MyComponent.vue'`  


The issue now is that I thought I could get away without building and publish this repo each time I make changes since all the packages depending on this shared package are also using TypeScript...  


So, when I reference a component from the shared library inside my other "app projects"

    import { MyComponent } from '@org/scl'

I get the following error because it's defined in the SCL repo.

    These dependencies were not found:
    @/components/SubFolder/AnotherComponent.vue

It's pretty obvious because the tsconfig.json inside the App-A repo doesn't know about the path aliases inside the SCL-repo.

Is there a pretty way to go around this, or is the best to make sure to build/transpile the TS to JS and always work against the output?

Let me know if I need to clarify something.

Br,  
Hopeless dude
## [11][Build Mobile Cross Platform Volunteer Delivery App in React Native](https://www.reddit.com/r/typescript/comments/gx6wyv/build_mobile_cross_platform_volunteer_delivery/)
- url: https://www.reddit.com/r/typescript/comments/gx6wyv/build_mobile_cross_platform_volunteer_delivery/
---
[https://www.education-ecosystem.com/glenettn/2bOj9-how-to-create-volunteer-delivery-app/jYAEb-00-intro/](https://www.education-ecosystem.com/glenettn/2bOj9-how-to-create-volunteer-delivery-app/jYAEb-00-intro/)
