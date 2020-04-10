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
## [2][If x has a type of T, then what's the type of JSON.parse(JSON.stringify(x))?](https://www.reddit.com/r/typescript/comments/fxub0f/if_x_has_a_type_of_t_then_whats_the_type_of/)
- url: https://effectivetypescript.com/2020/04/09/jsonify/
---

## [3][Porting to TypeScript Solved Our API Woes](https://www.reddit.com/r/typescript/comments/fy37h9/porting_to_typescript_solved_our_api_woes/)
- url: https://www.executeprogram.com/blog/porting-to-typescript-solved-our-api-woes
---

## [4][typescript object literal problem](https://www.reddit.com/r/typescript/comments/fy5v3r/typescript_object_literal_problem/)
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
## [5][Linkdash - Generate a handy dashboard of links in seconds!](https://www.reddit.com/r/typescript/comments/fxpx9q/linkdash_generate_a_handy_dashboard_of_links_in/)
- url: https://i.redd.it/3mgoes21prr41.gif
---

## [6][Debugging typescript dependencies from a non typescript node project](https://www.reddit.com/r/typescript/comments/fy81ch/debugging_typescript_dependencies_from_a_non/)
- url: https://www.reddit.com/r/typescript/comments/fy81ch/debugging_typescript_dependencies_from_a_non/
---
I've seen setups for debugging Frontend code with source maps, and complex configurations for debugging entire typescript projects in vscode with sourcemaps.  


However, what if I have a project that is in Vanilla Node.js, but some of my dependencies are in Typescript? How do I debug those in vscode or chrome?
## [7][Typescript problem using GraphQL Nexus](https://www.reddit.com/r/typescript/comments/fxz4f6/typescript_problem_using_graphql_nexus/)
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
## [8][Typescript problem with a GraphQl Nexus resolver](https://www.reddit.com/r/typescript/comments/fy1iaz/typescript_problem_with_a_graphql_nexus_resolver/)
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
## [9][TypeScript Tutorial For Beginners](https://www.reddit.com/r/typescript/comments/fxo95b/typescript_tutorial_for_beginners/)
- url: https://afteracademy.com/blog/typescript-tutorial-for-beginners
---

## [10][Instantly find code using intuitive search and code rank — live demo (metacode.app)](https://www.reddit.com/r/typescript/comments/fxfeih/instantly_find_code_using_intuitive_search_and/)
- url: https://v.redd.it/ythfwk0dxnr41
---

## [11][I'm forced to discriminate a union type to then perform the same operation (Visitor Pattern for ASTs)](https://www.reddit.com/r/typescript/comments/fxolgd/im_forced_to_discriminate_a_union_type_to_then/)
- url: https://stackoverflow.com/questions/61111469/forced-to-discriminate-a-union-type-to-then-perform-the-same-operation-visitor?stw=2
---

