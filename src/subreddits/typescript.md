# typescript
## [1][Who's hiring Typescript developers - September](https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/)
- url: https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/
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
## [2][Optional types if selection in enum is chosen](https://www.reddit.com/r/typescript/comments/iol6lc/optional_types_if_selection_in_enum_is_chosen/)
- url: https://www.reddit.com/r/typescript/comments/iol6lc/optional_types_if_selection_in_enum_is_chosen/
---
I'm working on an auth service and im quite new to TS. Looking to write a type that will say "if the provider a user chooses to sign in with is 'email' than the 'email' and 'password' params are no longer optional"

An example of what i mean:

`type SocialSignInParams = {provider: 'apple' | 'facebook' | 'google';email?: string;password?: string;};type EmailSignInParams = {provider: 'email';email: string;password: string;};`

`const signIn = async ({ provider, email, password }: SignInParams) =&gt; {if (provider === 'email') {return await signInWithEmail({ email, password });}};`

Thanks!  


Edit: Appreciate all the responses.
## [3][expressive: a lightweight library for creating express routers using decorators](https://www.reddit.com/r/typescript/comments/inujj0/expressive_a_lightweight_library_for_creating/)
- url: https://github.com/BitMountain/expressive
---

## [4][How to return proper method's type based on a class constructor argument type](https://www.reddit.com/r/typescript/comments/io38en/how_to_return_proper_methods_type_based_on_a/)
- url: https://www.reddit.com/r/typescript/comments/io38en/how_to_return_proper_methods_type_based_on_a/
---
Hello, i have this class (minimum reproducible case):

    class MyClass {
      constructor(public bar: string | string[]) {}
    
      foo(): string | string[] {
        if (Array.isArray(this.bar)) {
          return "example";
        }
    
        return ["more", "example"];
      }
    }
    
    new MyClass("example").foo(); // `string | string[]` can't we infer the right type based on `bar`?

When called it is of type union. Could it be otherwise? Still learning TS :(

For more info on what it is i'm trying to do: I would like an class that can query a distant api and if the arg `bar` is an array, it will query the api using multiple account and thus returning an array on responses. if `bar` is a string, it returns a single response, not in an array.

Thx people of the internet!
## [5][Running your unit tests written in TypeScript faster with zero overhead!](https://www.reddit.com/r/typescript/comments/ingrc2/running_your_unit_tests_written_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/ingrc2/running_your_unit_tests_written_in_typescript/
---
[swc-node](https://github.com/Brooooooklyn/swc-node) is a TypeScript/JavaScript compiler which is a NodeJS native addon for [swc](https://github.com/swc-project/swc).

Swc compiler is 7\~40x faster than babel, swc-node implement [@swc-node/jest](https://github.com/Brooooooklyn/swc-node/tree/master/packages/jest) and [@swc-node/register](https://github.com/Brooooooklyn/swc-node/tree/master/packages/register) top on it.

swc-node is also built on top of [napi-rs](https://github.com/napi-rs/napi-rs) , which has some killer features rather than the other packages which built from native code:

* Minimal third part dependencies, no need for `node-gyp rust/C++ toolchains.` You can just install it in your Docker environment and use it.
* No additional binary files need to download in **postinstall** scripts. Everything is available on `npm`.
* Compatible to all NodeJS versions after `v8.9.0,` No need to reinstall after NodeJS upgraded.

And thanks to deno ecosystem(swc is heavily used by deno ecosystem), swc is stable and support latest TypeScript/JavaScript feature now:

* TypeScript `experimentalDecorators`
* TypeScript  `emitDecoratorMetadata`
* `jsx pargma`
* ES5 \~ ES2020 language features

More over the swc featues, I implement `jest-hoist-transformer` in  [@swc-node/jest](https://github.com/Brooooooklyn/swc-node/tree/master/packages/jest), which would help you migrate from ts-jest smoothly.

Many projects in [Bytedance](https://bytedance.com/en/) has running with [@swc-node/jest](https://github.com/Brooooooklyn/swc-node/tree/master/packages/jest) (jest tests) and  [@swc-node/register](https://github.com/Brooooooklyn/swc-node/tree/master/packages/register) (ava tests).

&amp;#x200B;

https://preview.redd.it/ggdv42g1ugl51.png?width=1778&amp;format=png&amp;auto=webp&amp;s=630ccb99bdd6cb575f2327a9275e3ef97f2a83ca

So try it and enjoy the zero overhead performance improvement!
## [6][How do I assign a value to an object in a for loop?](https://www.reddit.com/r/typescript/comments/inhw12/how_do_i_assign_a_value_to_an_object_in_a_for_loop/)
- url: https://www.reddit.com/r/typescript/comments/inhw12/how_do_i_assign_a_value_to_an_object_in_a_for_loop/
---
Sorry if this is a really basic question. I am relatively new to javascript (and after this typescript).

I am making an angular app, which has 2 arrays of objects. The first is fetched from a REST service, and the second is a locally stored  JSON.

I made a model as follows

    export interface MenuAction {
    menuId: string;
    actionId: string;
    actionType: string;
    updIndic: string; 
    description: string; 
    longDesc?:string;
    image?: string; }

My local Json looks like this.

&amp;#x200B;

    [{
    "actionId" : "0010", "image" : "policy.jpg", "text" : "Actions related to policy system configuration. For example, configuration of affinity agreements"
    },
    { "actionId" : "0020",
    "image" : "account.jpg",
    "text" : "Account system configuration" 
    }]

&amp;#x200B;

Then in my component.ts, I have done the following.

    ------------
    import menuDetails from '../../../data/menu_details.json'; 
     ------
    menuActions : MenuAction[]; 

So basically, I want to loop through the "details" object, and when I find a matching actionId, I want to "enrich" the "menuActions" object with the additional information (so I can use it all in a card on the html)

I tried this in ngInit, but I am missing something.

&amp;#x200B;

    ngOnInit(): void {
    
    this._menuaction.getMenuActions('0002').subscribe (data =&gt; { this.data = data;  this.menuActions = this.data.content.userMenu.nbaUserMenuViewArray; 
    
    console.log (menuDetails);  // Loop all of the menuactions
    
    let index; for (let action of this.menuActions) {
    // Find the right details object  
    index = menuDetails.findIndex(x =&gt; x.actionId === action.actionId); 
    console.log ('Found at '+ index); 
    this.menuActions[action].longDesc = menuDetails[index].text;  // This line fails
    }          
    });

Can anyone point me in the right direction?
## [7][Strict null checks (and function hoisting?) causing issues .](https://www.reddit.com/r/typescript/comments/in31qc/strict_null_checks_and_function_hoisting_causing/)
- url: https://www.reddit.com/r/typescript/comments/in31qc/strict_null_checks_and_function_hoisting_causing/
---
[Here](https://www.typescriptlang.org/play?#code/GYVwdgxgLglg9mABMOcAUEAWBDATgfgC5EBvRAEwFMAHKTIxMEAWwCNLdEB6AKmblyVEcWvDDYANomp5szKBwDOiHlwC+ASlIAoAJAxgiDDk4Bec4nBVgMMJXJaseRKdJqA3HomUopCjTpENRdEJ1xPfUM0KlpMFwsrShs7B39YkIAGT0Qc7i5+QWFRBElpWWYfJURbGFhJGAAve2rDOkoAT0Q8IUTk+z0DIyh26ko4QxjAgEILACImNg5ZrTpcOAB3RABRXDXcNFnl7NyufIEhEVgSqRlcOUrcZW7QzEoIAGtmlE42mB+R5qDNqdZ7UNYANxgVHIiFYnTalkUHD0EAQil8wCQrjQWlMAD4dLpdJM4gBqRAARiyp3WmE6tWqynRMAkNzgikUMFY3k47EsYGstma+G0uTFeQKF2K4ikzxASPIeg82jUQA) is the playground link .

Why `depth` is still possibly undefined inside `fn` ? and how do I make it not be without adding an unnecessary `if` clause ?
## [8][How much OOP paradigm is supported by TypeScript, compare to Java and C#?](https://www.reddit.com/r/typescript/comments/in31g0/how_much_oop_paradigm_is_supported_by_typescript/)
- url: https://www.reddit.com/r/typescript/comments/in31g0/how_much_oop_paradigm_is_supported_by_typescript/
---
I'm JavaScript React developer who is migrating to TypeScript. I would like to know that how much OOP paradigm supported by TypeScript compare to Java and C#? Anyone who came from Java or C# background? Can you explain this please?
## [9][What `{type: T}` does here?](https://www.reddit.com/r/typescript/comments/imwrt7/what_type_t_does_here/)
- url: https://www.reddit.com/r/typescript/comments/imwrt7/what_type_t_does_here/
---
\`\`\`  
type LookUp&lt;U, T extends string&gt; = { \[K in T\]: U extends { type: T } ? U : never }\[T\]  
\`\`\`  


I don't understand this helper type at all, because I don't know what {type: T} does here, also, this is not in the docs. So could you please explain this to me, thank you!
## [10][Why my lexer is reading '=' twice?](https://www.reddit.com/r/typescript/comments/in6cee/why_my_lexer_is_reading_twice/)
- url: https://www.reddit.com/r/typescript/comments/in6cee/why_my_lexer_is_reading_twice/
---
Github repo: [https://github.com/Mdsp9070/monkeylanguage](https://github.com/Mdsp9070/monkeylanguage)

&amp;#x200B;

I'm building an interpreter with Typescript all covered by tests. Recently I added some extra methods, like isLetter, isDigit, skipWhitespace and so on!

&amp;#x200B;

My last commit: [https://github.com/Mdsp9070/monkeylanguage/commit/687f56cf921fff9cfcc0862c6df773f6035945b8](https://github.com/Mdsp9070/monkeylanguage/commit/687f56cf921fff9cfcc0862c6df773f6035945b8)

&amp;#x200B;

However my test is failing but I don't know why.

That's my input:

      const input = `let five = 5;
    let ten = 10;
    let add = fn(x, y) {
    x + y;
    };
    let result = add(five, ten);
    `;

And that's my test result:

    yarn run v1.22.4
    $ jest
     FAIL  interpreter/src/lexer/lexer.test.ts
      test lexer
        ✕ should return the right token type and literal (71 ms)
    
      ● test lexer › should return the right token type and literal
    
        expect(received).toEqual(expected) // deep equality
    
        Expected: "INT"
        Received: "ASSIGN"
    
          57 | 
          58 |       console.log(token)
        &gt; 59 |       expect(token.type).toEqual(test.type);
             |                          ^
          60 |       expect(token.literal).toEqual(test.literal);
          61 |     }
          62 |   });
    
          at Object.&lt;anonymous&gt; (interpreter/src/lexer/lexer.test.ts:59:26)
    
      console.log
        l
    
          at Lexer.nextToken (interpreter/src/lexer/lexer.ts:71:13)
    
      console.log
        { literal: 'let', type: 'LET' }
    
          at Object.&lt;anonymous&gt; (interpreter/src/lexer/lexer.test.ts:58:15)
    
      console.log
        f
    
          at Lexer.nextToken (interpreter/src/lexer/lexer.ts:71:13)
    
      console.log
        { literal: 'five', type: 'IDENT' }
    
          at Object.&lt;anonymous&gt; (interpreter/src/lexer/lexer.test.ts:58:15)
    
      console.log
        =
    
          at Lexer.nextToken (interpreter/src/lexer/lexer.ts:71:13)
    
      console.log
        { type: 'ASSIGN', literal: '=' }
    
          at Object.&lt;anonymous&gt; (interpreter/src/lexer/lexer.test.ts:58:15)
    
      console.log
        =
    
          at Lexer.nextToken (interpreter/src/lexer/lexer.ts:71:13)
    
      console.log
        { type: 'ASSIGN', literal: '=' }
    
          at Object.&lt;anonymous&gt; (interpreter/src/lexer/lexer.test.ts:58:15)

Why is '=' being reading twice?
## [11][Does TypeScript check types at runtime?](https://www.reddit.com/r/typescript/comments/in51pa/does_typescript_check_types_at_runtime/)
- url: https://www.reddit.com/r/typescript/comments/in51pa/does_typescript_check_types_at_runtime/
---
Hey guys, I have recently implemented TypeScript in my React app. I was wondering if TypeScript checks types at runtime or not, since I have not re-tested all my features, and I'm afraid that my app will break when a user is using it and TypeScript finds a type error
