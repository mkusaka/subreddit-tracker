# typescript
## [1][Who's hiring Typescript developers - October](https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/)
- url: https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/
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
## [2][const _ = this;](https://www.reddit.com/r/typescript/comments/jj0zlr/const_this/)
- url: https://www.reddit.com/r/typescript/comments/jj0zlr/const_this/
---
.Net background and I need to know why this code exists in some libraries. What good does it do?

const \_ = this;  
\_.localVariable = &lt;some arbitrary value&gt;;

Would it not be the same as just using this.localVariable?
## [3][Need help with type-graphql and typegoose](https://www.reddit.com/r/typescript/comments/jiyhvh/need_help_with_typegraphql_and_typegoose/)
- url: https://www.reddit.com/r/typescript/comments/jiyhvh/need_help_with_typegraphql_and_typegoose/
---
I'm using `typegoose` with `type-graphql` and when I try to use nested `@InputType()`, the nested object gets converted into `mongoose.Types.ObjectId()`. How do I handle nested InputTypes

This is my code (`GrandChild` is not a mongoose document, it's a simple object-type)

```ts
@ObjectType()
export class GrandChild {
  @Field()
  name: string;
}

@ObjectType()
export class Child {
  @prop()
  @Field()
  name: string;

  @prop({ ref: GrandChild })
  @Field(() =&gt; GrandChild)
  grandChild: Ref&lt;GrandChild&gt;;
}

@ObjectType()
export class Parent {
  @prop()
  @Field()
  name: string;

  @prop({ ref: Child })
  @Field(() =&gt; Child)
  child: Ref&lt;Child&gt;;
}

@InputType()
export class GrandChildInput {
  @Field()
  name: string;
}

@InputType()
export class ChildInput {
  @Field()
  name: string;

  @Field(() =&gt; GrandChild)
  grandChild: GrandChildInput;
}

@InputType()
export class Parent {
  @Field()
  name: string;

  @Field(() =&gt; Child)
  child: ChildInput;
}
```

---

Example input:

```json
{
  "name": "Parent A",
  "child": {
    "name": "Child A",
    "grandChild": {
      "name": "GrandChild A"
    }
  }
}
```

`parent` query:

```gql
{
  parent {
    name
    child {
      name
      grandChild {
        name
      }
    }
  }
}
```

When I run the `parent` query, I get the following output

```bash
Cannot read property "Child.grandChild" of undefined.
```

I used mongo bash to get the parent document and this is what I got:

db.parent.find():
```json
{
  name: "Parent A",
  child: ObjectId("some-object-id-here")
}
```

db.child.find():
```json
{
  name: "Child A",
  grandChild: ObjectId("some-object-id-here")  // &lt;-- This is not supposed to be a document
}
```

How do fix this?
## [4][Dendron – A Hierarchical Tool for Thought](https://www.reddit.com/r/typescript/comments/jijqyl/dendron_a_hierarchical_tool_for_thought/)
- url: https://github.com/dendronhq/dendron
---

## [5][How can I make an inline assertion?](https://www.reddit.com/r/typescript/comments/jiqo2b/how_can_i_make_an_inline_assertion/)
- url: https://www.reddit.com/r/typescript/comments/jiqo2b/how_can_i_make_an_inline_assertion/
---
My larger question is essentially, what's the best way to get the typing checker to be happy with the following code snippet?

```
function foo(param: number | null | undefined): number | null { 
  const val = param !== undefined;
  return val ? param : null;
}
```

I want to avoid casting. I'd also like to avoid changing the ternary operator with an `if/else` branch.



My code has access to an assert function typed as so:

```
declare function assert(value: any): asserts value;
```

My initial thought was to call this inline using the `,` operator like so:

```
function foo(param: number | null | undefined): number | null { 
  const val = param !== undefined;
  return val ? (assert(param !== undefined), param) : null;
}
```

But `param` is still typed as `number | null | undefined` despite the assertion throwing if `param === undefined`.

Does anyone know a better way to call the assertion keeping my constraints I mentioned above? Or a better way to make the typing checker happy?
## [6][Using Parser5](https://www.reddit.com/r/typescript/comments/jio9h4/using_parser5/)
- url: https://www.reddit.com/r/typescript/comments/jio9h4/using_parser5/
---
Is there anyway to use parser5 in a typescript project? I noticed that it was written in javascript, and uses const parse5 = require('parse5') to import it. If not is there anything that is very similar to it? I need the parser to keep track of the objects location in the DOM
## [7][Continuations in TypeScript](https://www.reddit.com/r/typescript/comments/jiapps/continuations_in_typescript/)
- url: https://medium.com/@wim_jongeneel/continuations-in-typescript-db18402010bc?source=friends_link&amp;sk=207a12570b354ebe74d7bfa220461ea5
---

## [8][This week’s open source newsletter just went out! This one had some seriously interesting projects, like a gaming framework written in Typescript released by Microsoft and a timeline component for React!](https://www.reddit.com/r/typescript/comments/jhu66o/this_weeks_open_source_newsletter_just_went_out/)
- url: https://console.substack.com/p/console-24?r=3cbez&amp;utm_campaign=post&amp;utm_medium=web&amp;utm_source=copy
---

## [9][How to select a type from a union, using a literal type property of said type?](https://www.reddit.com/r/typescript/comments/jhyf8c/how_to_select_a_type_from_a_union_using_a_literal/)
- url: https://www.reddit.com/r/typescript/comments/jhyf8c/how_to_select_a_type_from_a_union_using_a_literal/
---
I have a reducer in react. The action can be one of 8 types, but for simplicity, let's imagine that there's only 2 types

```typescript
type Add = {
  type: 'add';
  id: string;
  value: string;
}

type Remove = {
  type: 'remove';
  id: string;
}

type Action = Add | Remove;
```
Instead of using a switch case, I'm using an object of handlers, where each handler is a function that handles a specific action
```typescript
const handlers = {
  add: (state, action) =&gt; state,
  remove: (state, action) =&gt; state,
  default: (state, action) =&gt; state,
}

const reducer = (state, action) =&gt; {
  const handler = handlers[action.type] || handlers.default;
  return handler(state, action);
}
```
Now I want to type the `handlers` object appropriately. So the handler function should accept an action of type corresponding to its key in the `handlers` object.
```typescript
type Handlers = {
  [key in Action["type"]]: (state: State, action: Action) =&gt; State
//                                                ↑this here should be the action
//                                                which has type matching to it's key.
//                                                So when the key is 'add', it should
//                                                be of type Add, and so on.
}
```
All I could think of is to explicitly state the key and the matching action type. Is There a way to 'pick' the type from the union, according to the value of the key?

[link to stackoverflow question](https://stackoverflow.com/questions/64527150/in-typescript-how-to-select-a-type-from-a-union-using-a-literal-type-property)
## [10][Microsoft TileCode](https://www.reddit.com/r/typescript/comments/jhjpq4/microsoft_tilecode/)
- url: https://github.com/microsoft/tilecode
---

## [11][Best services for RBAC and authorization](https://www.reddit.com/r/typescript/comments/jhk048/best_services_for_rbac_and_authorization/)
- url: https://www.reddit.com/r/typescript/comments/jhk048/best_services_for_rbac_and_authorization/
---
Hey guys,

I’m starting work on a project and had some questions about Auth0, Cognito and the likes. 

I understand that they can provide RBAC. But do I have to manually check and prevent if a particular user can edit a resource? Or can I leverage these services?

For example: I have a REST API `/user/:id ` to which you can POST to update information. How do I ensure that (with a JWT or something) the user logged in can only update resources they have access to?
