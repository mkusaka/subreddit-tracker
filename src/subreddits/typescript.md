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
## [2][Disable duck typing for literal types](https://www.reddit.com/r/typescript/comments/ii4vsf/disable_duck_typing_for_literal_types/)
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
## [3][file changes not updating when running TSC](https://www.reddit.com/r/typescript/comments/ii6p7m/file_changes_not_updating_when_running_tsc/)
- url: https://www.reddit.com/r/typescript/comments/ii6p7m/file_changes_not_updating_when_running_tsc/
---
I recently opened a repo I set aside for a few weeks and notice my changes to .ts file inside the `src` folder, are not making it to the .js files that in the `compiled` folder.

I added this to the script file server.js and after restarting the server, even this is not printing to terminal:

    console.log(`updated on ${new Date()}`);

There are no warnings, or any other feedback when running `tsc` in terminal within the project root.

Here is my tsconfig.json file, is there a mistake somewhere?

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
        "outDir": "compiled"
      },
      "include": [
        "src", "__mocks__"
      ],
      "exclude": ["node_modules", "**/*.test.ts"]
    }
## [4][Opinion on explicit return types?](https://www.reddit.com/r/typescript/comments/ihpnme/opinion_on_explicit_return_types/)
- url: https://www.reddit.com/r/typescript/comments/ihpnme/opinion_on_explicit_return_types/
---
Hi All,

What is your opinion on always mentioning return types on each of your method? I mean for simple methods like `checkSomething(),`it's obvious to me it should not be returning and it most likely a void. Or something like `getURL`, also seems something which will return a string. Do you guys have any strong opinions on why adding types should be absolutely necessary?
## [5][Faven: a web tool to generate favicons with alpinejs](https://www.reddit.com/r/typescript/comments/ihpw9k/faven_a_web_tool_to_generate_favicons_with/)
- url: https://faven.netlify.app/
---

## [6][Need help mapping through an array that corresponds to an objects keys](https://www.reddit.com/r/typescript/comments/ihun4q/need_help_mapping_through_an_array_that/)
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
## [7][Getting started with LitElement and TypeScript](https://www.reddit.com/r/typescript/comments/ihi6cc/getting_started_with_litelement_and_typescript/)
- url: https://labs.thisdot.co/blog/getting-started-with-litelement-and-typescript
---

## [8][Possible to infer keys of a nested object for argument type?](https://www.reddit.com/r/typescript/comments/ihrlq1/possible_to_infer_keys_of_a_nested_object_for/)
- url: https://www.reddit.com/r/typescript/comments/ihrlq1/possible_to_infer_keys_of_a_nested_object_for/
---
Hey all, looking to type out a function and can't work out one arg.

I have multiple objects storing the array of roles allowed, which are then namespaced under one object by their route/screen for the helper function. That function will take the key for the screen, and then which key under that specific permissions object to check:  Currently my union of 

    keyof typeof getAuthPermissions | keyof typeof getUserPermissions

returns ALL keys obviously when using the function like 

    checkUserPermission('user', '&lt;all-keys&gt;') // where IDE populates all possible keys here

where I want to infer just the keys from the merged object from whichever screen is passed in to the function:  


All relevant code:

    const userGroups = ['Manager', 'Reporting', 'Payments', 'Support'] as const;
    type UserGroup = typeof userGroups[number];
    
    const getUserPermissions = {
        userControls: ['Manager', 'Support', 'Payments'],
        wallet: ['Manager', 'Support']
    };
    const getAuthPermissions = {
        logoutUser: ['Manager'],
        extendUserBonus: ['Manager', 'Support']
    }
    // ... a fair few more.
    const allPermissions = {
        user: getUserPermissions,
        auth: getAuthPermissions
        // ...
    }
    type Screens = keyof typeof allPermissions;
    
    function checkUserPermission(
        screen: Screens,
        userRoles: UserGroup[],
        // this should be the keys for whichever screen was passed in, something like keyof typeof allPermissions[screen]?
        permissionToCheck: keyof typeof getAuthPermissions | keyof typeof getUserPermissions // returns ALL keys. want to infer just from screen passed in
    ): boolean {
        for (const role of userRoles) {
            if (allPermissions[screen][permissionToCheck].includes(role)) {
                return true;
            }
        }
        return false;
    }

Is this possible? I know getting nested keys can be quite tricky, but I'm hoping that with inferring from the argument passed in that this should be possible?

Thanks!
## [9][Help: Generics for nested array, nested n times](https://www.reddit.com/r/typescript/comments/ihpcj0/help_generics_for_nested_array_nested_n_times/)
- url: https://www.reddit.com/r/typescript/comments/ihpcj0/help_generics_for_nested_array_nested_n_times/
---
I am writing a little piece of TypeScript code, to reverse a chessboard (i.e. so I can view it from either black or white's perspective). The code (thanks stackoverflow), is fairly generic in that it will take an array as an argument, and recursively reverse down to the bottom. Obviously I'm dealing with a 2d array to represent the chess board.

Instead of just typing it as accepting `any[]` and returning `any[]` I'd like to try use generics so the output retains the type that was passed in... if this is possible.

This is my attempt, but it gives the commented error

    export function reverseNestedArrays&lt;T&gt;(arr: T[]): T[] {
      return arr
        .map(function reverse(item) {
          return Array.isArray(item) &amp;&amp; Array.isArray(item[0])
            ? item.map(reverse)
            : item.reverse();   // Property 'reverse' does not exist on type 'T'.ts(2339)
        })
        .reverse();
    }
    

Ideally I'd like to be able to use generics in such a way so that it retains it's "reverse to any depth" capability. But if making it only work for a 2d array would make using generics possible, then I'd be happy with that.

Any help much appreciated. I have a very tenuous grasp of generics in general :)
## [10][How to deal with dynamic model values in static typing](https://www.reddit.com/r/typescript/comments/iho80j/how_to_deal_with_dynamic_model_values_in_static/)
- url: https://www.reddit.com/r/typescript/comments/iho80j/how_to_deal_with_dynamic_model_values_in_static/
---
Maybe it's because my background is in dynamic typing but there is one concept that I'm really struggling with and that is "dynamic model values". Let me explain:

Lets assume I have a simple Book model:

`class Book`

`name`

`author (Foreignkey)`

Depending on the use case I now want to return the author as an ID or a full instance from the API. I also may decide to add a field to the model that is not in the database but generated for a specific view.

Of course, I could add a union type to e. g. the author field type `author = number | IAuthor` with this approach I have to a lot of additional checks to tell the compiler that this particular component uses an author with the full instance attached to it. This approach has added a lot of clutter to my components and I don't like it.

So far I defined two different interfaces for both use cases e. g. `IBook` and `IBookWithAuthor.`

Are there any better ways of solving this?
## [11][Created VSCode launch config for debugging Mocha tests that written in TypeScript](https://www.reddit.com/r/typescript/comments/ihjn57/created_vscode_launch_config_for_debugging_mocha/)
- url: https://www.reddit.com/r/typescript/comments/ihjn57/created_vscode_launch_config_for_debugging_mocha/
---
It works fine with ES6 modules. Tested only in Windows 10 machine.

[https://gist.github.com/artem-mangilev/2f45ea18e4f5a38339212f551fa6a85c](https://gist.github.com/artem-mangilev/2f45ea18e4f5a38339212f551fa6a85c)
