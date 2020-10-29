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
## [2][Tracking down specific TypeScript source code files/function that is causing high CPU usage for TypeScript IDE/editor language services](https://www.reddit.com/r/typescript/comments/jk7avy/tracking_down_specific_typescript_source_code/)
- url: https://www.reddit.com/r/typescript/comments/jk7avy/tracking_down_specific_typescript_source_code/
---
I've got a TypeScript/Node project where I'm getting very high CPU usage while the project is open in vscode or phpstorm... quite often even when I'm not editing any files and the project is just open in a window in the background.  

My fans have been whirring up and down heaps over the last couple of months.  They're cranking right now while I've been typing this up, as they are most of the time lately.

It's just the one project with the problem, my other typescript project don't seem to have it - so it's not due to low end hardware (I've got a i7 7700 CPU, 32GB RAM and all files are on SSDs).

I have a few clues on which files it might be, but there's at least 10 of them, and they all contain a lot of code.  Some of it is big code-generated discriminated unions, which I think is a likely cause... but not sure which one(s).  I'm been programming in this style for longer than the CPU issue has been happening, but maybe I'm going too far with it lately given that typescript uses duck typing.

Are there any methods of finding out which source code files / functions might be causing the high CPU usage to the typescript service running in either vscode/phpstorm?  Or just by running `tsc` or some other tool outside of any editor?

I'm not sure exactly what it might involve... I'm hoping for maybe some kind of benchmarking/reporting thing that could point to specific `.ts` files, and ideally lines of code or type/function names that might be hogging the typescript service.

Last resort is to just start removing code by trial and error, but I have a feeling that approach might take me weeks, so I'm hopeful that maybe there's some better way.
## [3][VS Code extension for organizing classes/interfaces](https://www.reddit.com/r/typescript/comments/jk9vnq/vs_code_extension_for_organizing_classesinterfaces/)
- url: https://www.reddit.com/r/typescript/comments/jk9vnq/vs_code_extension_for_organizing_classesinterfaces/
---
I use linters to help expose the ordering of private/public instances, endpoints, etc. But I still find myself spending more time than I need just getting classes organized and structured well. Are there any extensions you have used that help auto-format in terms of organization?

Thanks!
## [4][I'm having trouble understanding how to polyfill nodejs dependencies with webpack 5. Are there any resources other than the docs that can help me understand this whole 'alias' thing?](https://www.reddit.com/r/typescript/comments/jjr8l0/im_having_trouble_understanding_how_to_polyfill/)
- url: https://www.reddit.com/r/typescript/comments/jjr8l0/im_having_trouble_understanding_how_to_polyfill/
---

## [5][What is late binding?](https://www.reddit.com/r/typescript/comments/jjm0kj/what_is_late_binding/)
- url: https://www.reddit.com/r/typescript/comments/jjm0kj/what_is_late_binding/
---
I only know TS and no other language. Does it make sense to talk about late binding in TS and if yes what is late binding?
## [6][const _ = this;](https://www.reddit.com/r/typescript/comments/jj0zlr/const_this/)
- url: https://www.reddit.com/r/typescript/comments/jj0zlr/const_this/
---
.Net background and I need to know why this code exists in some libraries. What good does it do?

const \_ = this;  
\_.localVariable = &lt;some arbitrary value&gt;;

Would it not be the same as just using this.localVariable?
## [7][JSDoc 'as const' equivalent?](https://www.reddit.com/r/typescript/comments/jj58lh/jsdoc_as_const_equivalent/)
- url: https://www.reddit.com/r/typescript/comments/jj58lh/jsdoc_as_const_equivalent/
---
I'm in an old project that is too huge to easily convert to Typescript, so I've been using JSDoc instead. The Typescript feature that I can't figure out how to replicate in JSDoc is using `as const` to fully type the property names **and** values of an object.

```ts
// In Typescript
const anObject = {hello: 'world'} as const;
// (type shows as {hello:'world'} instead of {hello:string}
```

Is there any equivalent for this in JSDoc? I've been completely unable to find anything that does this, so instead I have to basically copy-paste any static object as a type to be able to properly type these cases, which certainly isn't DRY.
## [8][Easing into Typescript?](https://www.reddit.com/r/typescript/comments/jj3o1q/easing_into_typescript/)
- url: https://www.reddit.com/r/typescript/comments/jj3o1q/easing_into_typescript/
---
Hi Typescript fans,

I have a moderately sized Express app (\~185 files, 8,700 LOC) written in Javascript with a small amount of Babel processing. As it grows in scope, I'd like to get some of the benefits of TypeScript, especially in taking advantage of what VSCode includes in its editor.

I've tried and failed to migrate to Typescript before but had a bunch of configuration headaches and reverted to the (working) Javascript. **Any advice for partially migrating an in-progress project?**

Here are some more details:

* I'm using import rather than require, and babel to transpile the production code. I should transpile Typescript to Javascript for production, right?
* Can anyone point me toward the most loose configuration possible that just lets me throw typescript when I want, and stick with Any wherever I don't want to get into it? You can tell me why this is a bad idea, but I don't think I have the patience for a full transition to super-strict TS right away.
* I'll probably be living with JS that I want to use with ESLint rules and prettier, and that needs to be transpiled to run on Node, PLUS TS that will have its own ESLint config and transpile rules. Is this going to burn me out?
## [9][Rocketcrab: Party Games for Phones](https://www.reddit.com/r/typescript/comments/jjaapt/rocketcrab_party_games_for_phones/)
- url: https://rocketcrab.com/
---

## [10][Need help with type-graphql and typegoose](https://www.reddit.com/r/typescript/comments/jiyhvh/need_help_with_typegraphql_and_typegoose/)
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
## [11][Dendron – A Hierarchical Tool for Thought](https://www.reddit.com/r/typescript/comments/jijqyl/dendron_a_hierarchical_tool_for_thought/)
- url: https://github.com/dendronhq/dendron
---

