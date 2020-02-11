# typescript
## [1][Who's hiring Typescript developers - February](https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/)
- url: https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/
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
## [2][What's the most advanced thing you know within TypeScript?](https://www.reddit.com/r/typescript/comments/f25ux0/whats_the_most_advanced_thing_you_know_within/)
- url: https://www.reddit.com/r/typescript/comments/f25ux0/whats_the_most_advanced_thing_you_know_within/
---
Hey,

I'm planning to give a presentation to provide insights about TypeScript at work soon and I was wondering what's the most advanced thing you can do with TypeScript?

So this is about TS itself and not an advanced project/app using it. A search on Google or the official docs gave me some examples but I was wondering if someone knows some handy things :) Or maybe some solution for a problem you would have needed a bigger workaround without TypeScript, etc.
## [3][Visual Debugging in VS Code](https://www.reddit.com/r/typescript/comments/f1shn6/visual_debugging_in_vs_code/)
- url: https://i.redd.it/pdi1ucn2f4g41.gif
---

## [4][Is there any ORM that has dynamically typed query results?](https://www.reddit.com/r/typescript/comments/f1zky5/is_there_any_orm_that_has_dynamically_typed_query/)
- url: https://www.reddit.com/r/typescript/comments/f1zky5/is_there_any_orm_that_has_dynamically_typed_query/
---
Is there any ORM (or query builder) that provides dynamic typing for the result of a query? For example, if I call `findOne` in TypeORM like this:

```
const user = await connection.getRepository(User).findOne({ select: ['firstName'] })
```

it would be nice if the `user` variable was typed to include only the `firstName` field and not all fields associated with the model. Ditto for eagerly loaded associations. However, that doesn't appear to be the case. Are the any libraries that do in fact do that?
## [5][utility type that equals a specific key within an interface given the type which that key is mapped to](https://www.reddit.com/r/typescript/comments/f23wx8/utility_type_that_equals_a_specific_key_within_an/)
- url: https://www.reddit.com/r/typescript/comments/f23wx8/utility_type_that_equals_a_specific_key_within_an/
---
I'm trying to find or write a utility type which would have the behavior of the hypothetical `ReverseLookup` type in this example:

```
interface User {
  name: string
}

interface Article {
  title: string
}

interface Shapes {
  author: User,
  editor: User,
  article: Article,
}

// these would compile
const a: ReverseLookup&lt;Shapes, User&gt; = 'author'; // ✓
const b: ReverseLookup&lt;Shapes, User&gt; = 'editor'; // ✓

// these would cause an error
const c: ReverseLookup&lt;Shapes, User&gt; = 'article'; // ✗ ts error
const d: ReverseLookup&lt;Shapes, Article&gt; = 'editor'; // ✗ ts error
```

Where `ReverseLookup` equals a specific key within an interface given the type which that key is mapped to. So given `Shapes` and `User`, `ReverseLookup` is any key within `Shapes` which is mapped to type `User` .
## [6][Conditional types for a dynamic React component in Typescript?](https://www.reddit.com/r/typescript/comments/f1x8gp/conditional_types_for_a_dynamic_react_component/)
- url: https://www.reddit.com/r/typescript/comments/f1x8gp/conditional_types_for_a_dynamic_react_component/
---
Hello typscript community! I have a question. Is there anyway for typescript to conditionally match types? I have the following dynamic react component, and I want **parameters** to be set based on the **type** property in the **SidebarProps** interface. Currently I can only escape an error by passing **any** to parameters which is obviously not ideal.

```typescript

type T = keyof typeof components;

type propertiesMap = {
  default: any;
  job: AddEditJobsProps;
  customer: AddEditCustomersProps;
};

interface SidebarProps {
  type: keyof typeof components,     
  parameters: propertiesMap[T] | any; 
} 

const components = {
  default: Loading,   
  job: AddEditJobs,   
  customer: AddEditCustomers 
}; 

export default function Sidebar(props: SidebarProps) { 
  const { open, toggle } = useToggle(); 
  const SelectComponent = components[props.type]; 

  return ( 
    &lt;RightSidebar open={open} toggleDrawer={toggle}&gt; 
      &lt;SelectComponent {...props.props} /&gt; 
    &lt;/RightSidebar&gt; 
  ); 
}
```
## [7][TypeScript Express tutorial #13. Using Mongoose virtuals to populate documents](https://www.reddit.com/r/typescript/comments/f1r5sg/typescript_express_tutorial_13_using_mongoose/)
- url: https://wanago.io/2020/02/10/typescript-express-tutorial-13-using-mongoose-virtuals-to-populate-documents/
---

## [8][TSConf:EU - Europe's TypeScript Conference](https://www.reddit.com/r/typescript/comments/f1msjk/tsconfeu_europes_typescript_conference/)
- url: https://tsconf.eu/
---

## [9][sql-lint - A command line SQL linter written in my favourite language &lt;3](https://www.reddit.com/r/typescript/comments/f1t1sr/sqllint_a_command_line_sql_linter_written_in_my/)
- url: https://github.com/joereynolds/sql-lint
---

## [10][trouble building typescript project on heroku (not the normal problem!)](https://www.reddit.com/r/typescript/comments/f1u20o/trouble_building_typescript_project_on_heroku_not/)
- url: https://www.reddit.com/r/typescript/comments/f1u20o/trouble_building_typescript_project_on_heroku_not/
---
Hey all, 

I am still getting used to typescript and am running into a rather strange issue...

&amp;#x200B;

When I run \`tsc\` in the project locally, I my project builds with no errors. However, when the Heroku deploy process runs \`tsc\`. I am getting build errors. It is only 3 or 4 errors for a medium sized project. It almost looks like some sort of config on the Heroku side is just a bit... off.

&amp;#x200B;

Has anyone experienced something similar to this? I am have these two scripts in my \`package.json\`:

&amp;#x200B;

"postinstall": "npm run compile",

"heroku-prebuild": "npm install --only=dev",

&amp;#x200B;

I have also already tried setting \`NPM\_CONFIG\_PRODUCTION=false\` as a heroku environment variable.

&amp;#x200B;

Anyone have any advice?

&amp;#x200B;

Many thanks for any input!
## [11][Weekend project: GraphQL-like queries and resolvers written in TypeScript](https://www.reddit.com/r/typescript/comments/f1fyar/weekend_project_graphqllike_queries_and_resolvers/)
- url: https://www.reddit.com/r/typescript/comments/f1fyar/weekend_project_graphqllike_queries_and_resolvers/
---
Hey folks, a month or two ago I asked myself, "what would it look like if I wanted to get GraphQL-like queries and resolvers in pure TypeScript?" I started hacking away and today I published the proof of concept under the title TypedQL. It allows you to define a graph query API and write queries against it. The real hero here is TypeScript Language Services, as seen in this [gif](https://imgur.com/fcUrI58).

Check it out [here](https://github.com/gregoryfabry/TypedQL); the README has some examples that show what the library can do. This is very much a proof of concept and there's definitely some black magic in the typings, but it can be downloaded and experimented with. I would love to hear your feedback, as I definitely think the concept has value and might be worth pursuing further.
