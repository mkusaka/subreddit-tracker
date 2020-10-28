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
## [2][What is late binding?](https://www.reddit.com/r/typescript/comments/jjm0kj/what_is_late_binding/)
- url: https://www.reddit.com/r/typescript/comments/jjm0kj/what_is_late_binding/
---
I only know TS and no other language. Does it make sense to talk about late binding in TS and if yes what is late binding?
## [3][const _ = this;](https://www.reddit.com/r/typescript/comments/jj0zlr/const_this/)
- url: https://www.reddit.com/r/typescript/comments/jj0zlr/const_this/
---
.Net background and I need to know why this code exists in some libraries. What good does it do?

const \_ = this;  
\_.localVariable = &lt;some arbitrary value&gt;;

Would it not be the same as just using this.localVariable?
## [4][Easing into Typescript?](https://www.reddit.com/r/typescript/comments/jj3o1q/easing_into_typescript/)
- url: https://www.reddit.com/r/typescript/comments/jj3o1q/easing_into_typescript/
---
Hi Typescript fans,

I have a moderately sized Express app (\~185 files, 8,700 LOC) written in Javascript with a small amount of Babel processing. As it grows in scope, I'd like to get some of the benefits of TypeScript, especially in taking advantage of what VSCode includes in its editor.

I've tried and failed to migrate to Typescript before but had a bunch of configuration headaches and reverted to the (working) Javascript. **Any advice for partially migrating an in-progress project?**

Here are some more details:

* I'm using import rather than require, and babel to transpile the production code. I should transpile Typescript to Javascript for production, right?
* Can anyone point me toward the most loose configuration possible that just lets me throw typescript when I want, and stick with Any wherever I don't want to get into it? You can tell me why this is a bad idea, but I don't think I have the patience for a full transition to super-strict TS right away.
* I'll probably be living with JS that I want to use with ESLint rules and prettier, and that needs to be transpiled to run on Node, PLUS TS that will have its own ESLint config and transpile rules. Is this going to burn me out?
## [5][JSDoc 'as const' equivalent?](https://www.reddit.com/r/typescript/comments/jj58lh/jsdoc_as_const_equivalent/)
- url: https://www.reddit.com/r/typescript/comments/jj58lh/jsdoc_as_const_equivalent/
---
I'm in an old project that is too huge to easily convert to Typescript, so I've been using JSDoc instead. The Typescript feature that I can't figure out how to replicate in JSDoc is using `as const` to fully type the property names **and** values of an object.

```ts
// In Typescript
const anObject = {hello: 'world'} as const;
// (type shows as {hello:'world'} instead of {hello:string}
```

Is there any equivalent for this in JSDoc? I've been completely unable to find anything that does this, so instead I have to basically copy-paste any static object as a type to be able to properly type these cases, which certainly isn't DRY.
## [6][Rocketcrab: Party Games for Phones](https://www.reddit.com/r/typescript/comments/jjaapt/rocketcrab_party_games_for_phones/)
- url: https://rocketcrab.com/
---

## [7][Need help with type-graphql and typegoose](https://www.reddit.com/r/typescript/comments/jiyhvh/need_help_with_typegraphql_and_typegoose/)
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
## [8][Dendron – A Hierarchical Tool for Thought](https://www.reddit.com/r/typescript/comments/jijqyl/dendron_a_hierarchical_tool_for_thought/)
- url: https://github.com/dendronhq/dendron
---

## [9][Great method to test execution time of TypeScript code](https://www.reddit.com/r/typescript/comments/jj4wcn/great_method_to_test_execution_time_of_typescript/)
- url: https://i.redd.it/3jtms5lo3ov51.jpg
---

## [10][How can I make an inline assertion?](https://www.reddit.com/r/typescript/comments/jiqo2b/how_can_i_make_an_inline_assertion/)
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
## [11][Using Parser5](https://www.reddit.com/r/typescript/comments/jio9h4/using_parser5/)
- url: https://www.reddit.com/r/typescript/comments/jio9h4/using_parser5/
---
Is there anyway to use parser5 in a typescript project? I noticed that it was written in javascript, and uses const parse5 = require('parse5') to import it. If not is there anything that is very similar to it? I need the parser to keep track of the objects location in the DOM
