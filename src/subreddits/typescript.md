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
## [2][Why is any nonsense assignable to InstanceType&lt;...&gt;[] here?](https://www.reddit.com/r/typescript/comments/iiqcv9/why_is_any_nonsense_assignable_to_instancetype/)
- url: https://www.reddit.com/r/typescript/comments/iiqcv9/why_is_any_nonsense_assignable_to_instancetype/
---
I want to create a function which accepts a class and a callback, which also accepts this class and returns its instances. It's not very practical, but it is a simplified example of my problem.

My code looks like this ([playground link](https://www.typescriptlang.org/play?ts=4.0.2#code/C4TwDgpgBAwg9gOwM7AE4FcDGw6qgXigQgHcAKAOioENUBzJALimoRAG0BdASgID4WbANwAoEQDN0CbAEtEUYBBQAeACpQIAD0UIAJkliIUGbLj5ladZqoA0USQmYX613vgEBJZMFaYIq8Ag1Pi5eAG8RKCioVAhgdFQEeylnOm5RAF8xTAAbaiQDAFkQGDyCqDCskQB6aqhVAAtoGQQdXQhdKHQkajpmgzgAayoKEUUUMmLS-KQ7AH1cmf4odmISKAWypDJuOzWNxYKdnlEauqoAI3RgKA8oTFYWbHRqHJyQGLiEpNpUahABuJBB8EEYIMhoOJUHAALbJBAAQjO9Sa9jgbzgJBadCgSAacHQOU6eMxgg0qGhqDsVxuMhuujgSgQAHJgEjxsBJiUtvNDgZ3CsAIwAJgAzHZmeI4HALrRmScgA)):

    type Constructor = new(...args: any[]) =&gt; any;

    function test&lt;T extends Constructor&gt;(arg: T, fun: (arg: T) =&gt; InstanceType&lt;T&gt;[]) 
    {
        return fun(arg);
    }

    class MyClass {}

    // The intended usage is ok...
    test(MyClass, _class =&gt; [new _class(), new _class()]);

    // ..but I can actually return arrays of any nonsense from fun!
    // The following should show an error, but it doesn't!
    test(MyClass, _class =&gt; [123, 'foobar']);

The second call to `test` allows me to return anything from the callback, despite its type clearly inferred as `(arg: typeof MyClass) =&gt; MyClass[]`. It seems wrong to me that a `(string | number)[]` return type is assignable to `MyClass[]`.

Is there a way to fix this and make it report the correct error?

Thanks!
## [3][Maybe&lt;string&gt; vs string | null](https://www.reddit.com/r/typescript/comments/ii88m3/maybestring_vs_string_null/)
- url: https://www.reddit.com/r/typescript/comments/ii88m3/maybestring_vs_string_null/
---
Hi there! Quick question to the functional buffs out there:

What's the advantage of using a `Maybe&lt;string&gt;` type over using `string | null`? With strict null checking you still need to handle the null case, but you avoid using external libraries and having to introduce new concepts to the other developers on the team.

Same for `Either&lt;Error, string&gt;` and `string | Error`.
## [4][Writing testable code that is easy to reason with . A function that executes a sequence of injected functions (17 functions injected in total, that none has a dependency injected). Should I make it a class with no dependency injection , and expose the functions as methods (so they are testable)?](https://www.reddit.com/r/typescript/comments/iie9f4/writing_testable_code_that_is_easy_to_reason_with/)
- url: https://www.reddit.com/r/typescript/comments/iie9f4/writing_testable_code_that_is_easy_to_reason_with/
---
So I have a function (lets call it parent) which has the single responsibility to accept a big list of  arguments and then provide the appropriate parameters to a sequence of functions that are executed (17 functions in total,lets call them children) . There is also some minor logic inside parent that involves some returned values from child functions ,some type checking and argument initialization, but it is nothing special or complex .

I think injecting 17 functions as parameters is an anti pattern in general . Also these child functions do not depend on injected dependencies and have a clear and simple single responsibility .

What I thought would be an easy to reason with , but also testable solution ,is to gather all those functions in a big class as methods . The class instances will expose those methods so they will be testable . But the class will expose also the parent function which executes all the child functions .

So I can go write a test file that can test each child function in isolation but also in combination (in fact whatever combination I like) with the other child functions .

What would you do in this case ? Does the big class sound good ? I am unable to reason with dependency injection here . I find that it does more harm than good. The child functions are really specific for the parent function and I believe they will never be reused anywhere else .
## [5][ts-expect-error . Is it justified to use it when in tests ?](https://www.reddit.com/r/typescript/comments/iic8jx/tsexpecterror_is_it_justified_to_use_it_when_in/)
- url: https://www.reddit.com/r/typescript/comments/iic8jx/tsexpecterror_is_it_justified_to_use_it_when_in/
---
More specifically I happen to test a function and I give to it some parameters of wrong type .
## [6][what else besides "sourceMap": true is needed to setup source mapping?](https://www.reddit.com/r/typescript/comments/iidy9b/what_else_besides_sourcemap_true_is_needed_to/)
- url: https://www.reddit.com/r/typescript/comments/iidy9b/what_else_besides_sourcemap_true_is_needed_to/
---
I'm trying to setup source mapping so that VsCode can attach a debugger session to my TS source files but I can tell from the stack traces that the node process is still pointing to the JS files inside `/compiled`

What else needs to be done before a debugger can map to TS source files? The process I'm running is at `./compiled/server.js`

My tsconfig.json:

    {
      "compilerOptions": {
        "target": "es5",
        "lib": [
          "dom",
          "dom.iterable",
          "esnext"
        ],
        "allowJs": true,
        "skipLibCheck": true,
        "esModuleInterop": true,
        "allowSyntheticDefaultImports": true,
        "strict": true,
        "forceConsistentCasingInFileNames": true,
        "module": "CommonJS",
        "moduleResolution": "node",
        "resolveJsonModule": true,
        "isolatedModules": true,
        "sourceMap": true,
        "rootDir": "src",
        "outDir": "compiled"
      },
      "include": [
        "src", "src/__mocks__"
      ],
      "exclude": ["node_modules", "**/*.test.ts", "compiled"]
    }
## [7][Disable duck typing for literal types](https://www.reddit.com/r/typescript/comments/ii4vsf/disable_duck_typing_for_literal_types/)
- url: https://www.reddit.com/r/typescript/comments/ii4vsf/disable_duck_typing_for_literal_types/
---
Is there a way to get TypeScript to warn me about the following code?

```typescript
type Name = string
function f (_: Name): void {}

const s: string = 'test'
// This should ideally give me an error:
// Argument of type 'string' is not assignable to parameter of type 'Name'
f(s)
```

I know that in most cases you would not want to get an error, but in this case it would allow me to combine it with a type guard like:

```typescript
function isName (name: string): name is Name {
  return name.length &lt; 30
}
```

I think there was a discussion on this topic in the TypeScript repo, but I cannot find it anymore.
## [8][Opinion on explicit return types?](https://www.reddit.com/r/typescript/comments/ihpnme/opinion_on_explicit_return_types/)
- url: https://www.reddit.com/r/typescript/comments/ihpnme/opinion_on_explicit_return_types/
---
Hi All,

What is your opinion on always mentioning return types on each of your method? I mean for simple methods like `checkSomething(),`it's obvious to me it should not be returning and it most likely a void. Or something like `getURL`, also seems something which will return a string. Do you guys have any strong opinions on why adding types should be absolutely necessary?
## [9][Faven: a web tool to generate favicons with alpinejs](https://www.reddit.com/r/typescript/comments/ihpw9k/faven_a_web_tool_to_generate_favicons_with/)
- url: https://faven.netlify.app/
---

## [10][Need help mapping through an array that corresponds to an objects keys](https://www.reddit.com/r/typescript/comments/ihun4q/need_help_mapping_through_an_array_that/)
- url: https://www.reddit.com/r/typescript/comments/ihun4q/need_help_mapping_through_an_array_that/
---
\*\*\*SOLVED!\*\*\*

I have an object which contains a list of objects. Each object is a game with parameters like game name and the game options.

    interface LobbyItem {
        name: string
        boardSize: number;
        host: Player;
        hostReady: boolean;
        hostTurn: null | boolean;
        opponent: null | Player;
        opponentReady: boolean;
    }

where player is another Object with more parameters. In order to display these games in my React component, i am creating an array with the objects keys and tying to map through like so:

    const gameList = Object.keys(lobbyState.gameList).map((game, index) =&gt; {
        return (
        &lt;span className="info-banner info-box settings-button" onClick={joinGame}&gt;
           lobbyState.gameList[game].name
        &lt;/span&gt;
        )
      }) 

&amp;#x200B;

But i'm getting a TypeScript error when I try to index lobbyState.gameList\[game\]:

    Element implicitly has an 'any' type because expression of type 'string' can't be used to index type 'LobbyItem'.
      No index signature with a parameter of type 'string' was found on type 'LobbyItem'.  TS7053

New to typescript, would someone be able to help me out?  isn't the index signature 'name' a string? I should mention that I'm using nested objects instead of an array so that I can delete the games when they are done.

Thanks in advance!
## [11][Getting started with LitElement and TypeScript](https://www.reddit.com/r/typescript/comments/ihi6cc/getting_started_with_litelement_and_typescript/)
- url: https://labs.thisdot.co/blog/getting-started-with-litelement-and-typescript
---

