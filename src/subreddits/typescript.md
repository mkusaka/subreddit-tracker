# typescript
## [1][Who's hiring Typescript developers - April](https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/)
- url: https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/
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
## [2][Isotope - fast, simple &amp; lightweight UI library](https://www.reddit.com/r/typescript/comments/fykjph/isotope_fast_simple_lightweight_ui_library/)
- url: https://github.com/Isotope-js/core
---

## [3][What's the best practice for data modeling, shared by both front-end and back-end?](https://www.reddit.com/r/typescript/comments/fyj2u6/whats_the_best_practice_for_data_modeling_shared/)
- url: https://www.reddit.com/r/typescript/comments/fyj2u6/whats_the_best_practice_for_data_modeling_shared/
---
Hi,

I wonder what's the best way of sharing the same data types between the client (React) and the server (Express + Socket.IO).

In my game I have different rooms, each room saves the current status, something like:

```js
class GameRoom {
    players: Player[];
    started: boolean;
    currentPlayerTurn; Player;
    dices: [number, number];

    constructor({players = [], started = false, currentPlayerTurn = null, dices = [1,1]) {
        this.players = players;
        this.started = started;
        this.currentPlayerTurn = currentPlayerTurn;
        this.dices = dices;
    }

    startGame() {
        this.currentPlayerTurn = this.players[0];
        this.started = true;
    }
    
    // etc..
}
```

The room is being generated in the server, being sent to the client as JSON, and then rebuilt in the client. I sync the data with socket events, and everything's perfect.

But there's a problem with the React side of the story: changing `GameRoom` properties won't cause a rerender. That means I have to `forceRerender()` each time something is edited, or listen to class changes. Both options are a mess and I [described it deeply on Stack Overflow](https://stackoverflow.com/questions/61142978/make-a-react-component-rerender-when-data-class-property-change/).

This mess made me think maybe classes are not the best way to go. Using interface will solve this problem entirely, but I do lose instance functions like `GameRoom.startGame()`, that will have to be turned into utility functions, like:

```js
export function startGame(gameRoom: GameRoom) {
    gameRoom.currentPlayerTurn = gameRoom.players[0];
    gameRoom.started = true;
}
```

which is another mess, since they're hidden in code, and the developer **needs to know they exist**, and not edit `gameRoom` directly.

If you guys have any idea on how to model my data types, I'd be more than happy to hear.\
Thanks!
## [4][If x has a type of T, then what's the type of JSON.parse(JSON.stringify(x))?](https://www.reddit.com/r/typescript/comments/fxub0f/if_x_has_a_type_of_t_then_whats_the_type_of/)
- url: https://effectivetypescript.com/2020/04/09/jsonify/
---

## [5][Porting to TypeScript Solved Our API Woes](https://www.reddit.com/r/typescript/comments/fy37h9/porting_to_typescript_solved_our_api_woes/)
- url: https://www.executeprogram.com/blog/porting-to-typescript-solved-our-api-woes
---

## [6][typescript object literal problem](https://www.reddit.com/r/typescript/comments/fy5v3r/typescript_object_literal_problem/)
- url: https://www.reddit.com/r/typescript/comments/fy5v3r/typescript_object_literal_problem/
---
Hi,

so in typescript, once we define an object like this 

```
const parsedData = {
            serviceType,
            stops,
            deliveries,
            requesterContact: {
                name: pickup.personInChargeName,
                phone: merchantPhoneValidator(pickup.personInChargePhoneNum),
            },
            specialRequests,
        };
```
Then we cannot dynamically add more fields like javascript
```
// error 
if (isPreOrder) {
       parsedData.scheduleAt = new Date(pickup.startAt).toISOString();
}
```
Nor we can add dummy field just letting typescript know we have this field
```
const parsedData = {
            scheduleAt : undefined  // error
            serviceType,
            stops,
            deliveries,
            requesterContact: {
                name: pickup.personInChargeName,
                phone: merchantPhoneValidator(pickup.personInChargePhoneNum),
            },
            specialRequests,
        };
```

Any better way to get around with this except disabling typescript? Thanks!
## [7][Linkdash - Generate a handy dashboard of links in seconds!](https://www.reddit.com/r/typescript/comments/fxpx9q/linkdash_generate_a_handy_dashboard_of_links_in/)
- url: https://i.redd.it/3mgoes21prr41.gif
---

## [8][Debugging typescript dependencies from a non typescript node project](https://www.reddit.com/r/typescript/comments/fy81ch/debugging_typescript_dependencies_from_a_non/)
- url: https://www.reddit.com/r/typescript/comments/fy81ch/debugging_typescript_dependencies_from_a_non/
---
I've seen setups for debugging Frontend code with source maps, and complex configurations for debugging entire typescript projects in vscode with sourcemaps.  


However, what if I have a project that is in Vanilla Node.js, but some of my dependencies are in Typescript? How do I debug those in vscode or chrome?
## [9][Typescript problem using GraphQL Nexus](https://www.reddit.com/r/typescript/comments/fxz4f6/typescript_problem_using_graphql_nexus/)
- url: https://www.reddit.com/r/typescript/comments/fxz4f6/typescript_problem_using_graphql_nexus/
---
So I just started using GraphQL Nexus today and I'm just doing some basic setup and noodling around with some stuff. I coded a simple registerUser mutation like so:

`export const Mutation = objectType({`  
  `name: 'Mutation',`  
 `definition(t) {`  
 `t.boolean('register', {`  
`args: {`  
`email: stringArg({ required: true }),`  
`password: stringArg({ required: true }),`  
 `},`  
`async resolve(_, { email, password }) {`  
 `console.log(password, email)`  
`return true`  
 `},`  
 `})`  
 `},`  
`})`

But when I replaced the arguments to 'register' with a RegisterInput input like so:

`export const RegisterInput = inputObjectType({`  
  `name: 'RegisterInput',`  
 `definition(t) {`  
 `t.string('email', { required: true })`  
 `t.string('password', { required: true })`  
 `},`  
`})`  


`export const Mutation = objectType({`  
  `name: 'Mutation',`  
 `definition(t) {`  
 `t.boolean('register', {`  
`args: {`  
`input: arg({ type: RegisterInput }),`  
 `},`  
`async resolve(_, { input: { email, password } }) {`  
 `console.log(password, email)`  
`return true`  
 `},`  
 `})`  
 `},`  
`})`

Everything works, but he resolver input argument `{ input: { email, password } }` gets typescript a little complain - Property 'email'/'password' does not exist on type '{ email: string; password: string; } | null | undefined'? How do I get those annoying red squiggles to go away?
## [10][Typescript problem with a GraphQl Nexus resolver](https://www.reddit.com/r/typescript/comments/fy1iaz/typescript_problem_with_a_graphql_nexus_resolver/)
- url: https://www.reddit.com/r/typescript/comments/fy1iaz/typescript_problem_with_a_graphql_nexus_resolver/
---
I'm creating a new MongoDB document in a GraphQL Nexus resolver:

`async resolve(_, { input: { email, password } }, { userData }) {`  
 `const hashedPassword = await bcrypt.hash(password, 10)`  
 `let newUser = {`  
  `email,`  
  `password: hashedPassword,`  
 `}`  
 `const { insertedId } = await userData.insertOne(newUser)`  
 `newUser._id = insertedId`  
  `return newUser`  
 `},`

but typescript says property '\_id' does not exist on type '{ email: string; password: string; }'. I'm stuck because this is how I create a new document in mongo using plain JS. What is the proper way to add the insertedId?
## [11][Why does typescript think my sort comparator is invalid?](https://www.reddit.com/r/typescript/comments/fy0qtz/why_does_typescript_think_my_sort_comparator_is/)
- url: https://www.reddit.com/r/typescript/comments/fy0qtz/why_does_typescript_think_my_sort_comparator_is/
---
[https://pastebin.com/hNgyeeBx](https://pastebin.com/hNgyeeBx)

I'm using a Map to match up the possible values of a \`Select\` to different comparators. When the user changes the dropdown, they can sort the elements differently. The code runs fine, but typescript doesn't like it.

I've tried a lot of different solutions to this: enum, casting, guard statement, and I'm out of ideas.

      if (comparator) {
        return SKILLS.sort(comparator); // `comparator` has a typescript error here:
        // TS2345: Argument of type 'Function' is not assignable to parameter
        // of type '(a: Skill, b: Skill) =&gt; number'. Type 'Function' provides
        // no match for the signature '(a: Skill, b: Skill): number'.
      }

Thanks in advance for any help!
