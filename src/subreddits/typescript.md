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
## [2][View the API exposed by a given module](https://www.reddit.com/r/typescript/comments/jf4c4t/view_the_api_exposed_by_a_given_module/)
- url: https://www.reddit.com/r/typescript/comments/jf4c4t/view_the_api_exposed_by_a_given_module/
---
Is there a tool or convention to view/document the API exposed by a given module?

```typescript
/// example/index.ts

export * from "./core";
export * from "./util";
```

```typescript
/// example/core/index.ts

export class Application {
  private constructor (...args: unknown[]) { … }
  
  public get id(): string { … }
  
  public async start(): Promise&lt;void&gt; { … }
  
  public static create(process: NodeJS.Process): Application {
    // …
    
    return new Application(…);
  }
}
```

```typescript
/// example/util/index.ts

export const VERSION = "0.0.0" as const;
```

For example, the above would expose a very small public API, consisting of:
- `Application`
- `Application#id`
- `Application#start()`
- `Application.create()`
- `VERSION`

Are TypeScript definition files what I'm looking for?
## [3][InversifyJS - How can I use container.get in an object literal?](https://www.reddit.com/r/typescript/comments/jf4axz/inversifyjs_how_can_i_use_containerget_in_an/)
- url: https://www.reddit.com/r/typescript/comments/jf4axz/inversifyjs_how_can_i_use_containerget_in_an/
---
I'm not sure if this is the place to ask this, but it's the best place I could find. I'm new to dependency injection with TS but did use it in C# once. When I try doing this:

    export let COMMAND_MAP = {
        "join": container.get&lt;ChannelJoiner&gt;(TYPES.ChannelJoiner)
    }

I get a runtime error that "get" is not defined. What am I doing wrong? Am I allowed to do this? I can send other related files if necessary.
## [4][Good book or other resource for learning how to define your models?](https://www.reddit.com/r/typescript/comments/jex71u/good_book_or_other_resource_for_learning_how_to/)
- url: https://www.reddit.com/r/typescript/comments/jex71u/good_book_or_other_resource_for_learning_how_to/
---
I’m comming from a c# background and currently have a huge model that is hard to understand and does nothing else but set the types. My model is a huge collection of classes inside classes that don’t have anything else but some properties(no constructor, methods getters or setters). I’m trying to learn how to improve this but can’t seem to find a good book or article
## [5][How to map an interface to another one](https://www.reddit.com/r/typescript/comments/jexpj2/how_to_map_an_interface_to_another_one/)
- url: https://www.reddit.com/r/typescript/comments/jexpj2/how_to_map_an_interface_to_another_one/
---
I'm sorry if this is improperly formatted, I copied this from another subreddit I posted it.
Okay folks, Idk where to ask (I got downvoted in [stack](https://stackoverflow.com/questions/64443221/how-to-map-different-interfaces-to-each-other-in-typescript)) but the question is clear there, it's more of a typescript question, I'm working on a store which has \`ICategory\`, \`IProduct\`, ...etc and I generated the GraphQL schema with graphql codegen, the schema has \`Category\`, \`Product\`, ...etc.

The ICategory is on the following format

```

export interface IBaseCategory {

id: string;

type: string;

name: string;

slug: string;

image: string|null;

items?: number;

parent?: this | null;

children?: this\[\];

customFields: ICustomFields; }
```

    and the GraphQl one is on the format of the server, names of attributes are slightly different, but it has children with has edges which have nodes, etc (GraphQL basic schema).
    Now, I've created a class for each type that queries that server and returns a `Category`, `Product`, ..etc (GraphQL format). and I want them to be mapped to the other format (The theme format, `ICategory, ...) so that I don't remove or alter the theme interfaces except as little as possible, details are in stackoverflow link.
    
    Thanks in advance.
## [6][Generate Google Presentation from Wikipedia article](https://www.reddit.com/r/typescript/comments/je6qg8/generate_google_presentation_from_wikipedia/)
- url: https://www.reddit.com/r/typescript/comments/je6qg8/generate_google_presentation_from_wikipedia/
---
[GSlides Maker](https://github.com/vilmacio/gslides-maker) is a open-source project for create google presentations from wikipedia content like a robot. Try it [CLI version](https://github.com/vilmacio/gslides-maker) and help me with the [web application](https://gslidesmaker.com/). I hope you like it.

https://preview.redd.it/98yuu3ceb3u51.png?width=681&amp;format=png&amp;auto=webp&amp;s=5259f69c148faeab52c9a10af6490d366bfec5c5
## [7][Is this an appropriate time to assert? (Array concat overload mismatch)](https://www.reddit.com/r/typescript/comments/je8q40/is_this_an_appropriate_time_to_assert_array/)
- url: https://www.reddit.com/r/typescript/comments/je8q40/is_this_an_appropriate_time_to_assert_array/
---
.flat seems not to work in my project so I have been concating 2d arrays to flatten them. In this case I think an assertion might be needed:

      const flattened: ITask[] = [].concat(...resolvedTasks);
    
    /*
    const resolvedTasks: ITask[][]
    No overload matches this call.
      Overload 1 of 2, '(...items: ConcatArray&lt;never&gt;[]): never[]', gave the following error.
        Argument of type 'ITask[]' is not assignable to parameter of type 'ConcatArray&lt;never&gt;'.
          The types returned by 'slice(...)' are incompatible between these types.
            Type 'ITask[]' is not assignable to type 'never[]'.
              Type 'ITask' is not assignable to type 'never'.
    */
    
    // but this passes lint
      const flattened: ITask[] = ([] as ITask[]).concat(...resolvedTasks);

Can someone break down what `ConcatArray&lt;never&gt;` means and whether my assertion should suit this type? My best read is it representations a type `ConcatArray`. But the `never` type argument negates its transformed concrete type value. Which is hard for me to understand.
## [8][How to make ConstructorParameters&lt;typeof Test&gt;[0] generic?](https://www.reddit.com/r/typescript/comments/je6kpb/how_to_make_constructorparameterstypeof_test0/)
- url: https://www.reddit.com/r/typescript/comments/je6kpb/how_to_make_constructorparameterstypeof_test0/
---
```
type TestParams = ConstructorParameters&lt;typeof Test&gt;[0];

class Test {
  x;
  y;
  constructor({ x, y }: { x: number; y: number }) {
    this.x = x;
    this.y = y;
  }
}

const factory = (numberOf: number, { x, y }: TestParams) =&gt; {
  console.log(numberOf);
  return new Test({ x, y });
};

factory(5, { x: 2, y: 11 });
```
I would like to make `TestParams` generic for other factory function, but:

```
type TestParams&lt;T&gt; = ConstructorParameters&lt;typeof T&gt;[0];
const factory = (numberOf: number, { x, y }: TestParams&lt;Test&gt;) =&gt; {...

```
or any permutation of it does not seem to work. Any ideas?

https://www.typescriptlang.org/play?#code/C4TwDgpgBAKhDOwAKBDATigtvKBeKAwgPYB2iaArgMbBFqoaYTARrwA8okRAZrAsAB8AbQAMAXQDcAWABQcqgBsU8HHERQA3nKhQAHjNm6Qh3VVLlqtNAApN+gDRQQUAL4AuLfs8kKmAEasks4+foFobgCUWjq6UMAAFgCW8AB0enj6pnGJKaku+CaxrnIl8rLmZMBQPCg0dAVQNr4BrADyPKGtaE72ek4uHvyIDFjw0biCMUZQlfBEihCpikQA5s1h7TyR2WjMFGgkUCQQAO7DwHaOzlGGroZytfVoIDYArL3eUABMA54AjP9bnIgA
## [9][Resolve out union types?](https://www.reddit.com/r/typescript/comments/je3n3i/resolve_out_union_types/)
- url: https://www.reddit.com/r/typescript/comments/je3n3i/resolve_out_union_types/
---
I'm trying to write some code where I start off possibly accepting a union, but if the instance is not of a specific type, I then resolve it so it is. Specifically, I pass in an object of a sequelize model type, say `Post`, which might be an actual post model, or it might be the object required to create a new post.

This is used for generating content - so I can arbitrarily create a new comment in testing, and optionally have it create a new post to put the comment on.

this is documented essentially like this.

    export const createTestComment = async (comment: CommentData, post: Post | PostData) : Promise&lt;Comment&gt; {
      if (!(post instanceof Post)) post = await createTestPost(post);
      return await post.createComment(comment);
    } 

The above doesn't work. Even though this is set so that post can NO LONGER BE of type PostData, typescript refuses to allow it because the type of the instance is still a union. Is it possible to change the type programatically to remove the union? I tried `createPost(post) as Post` but it didn't help. 

Am I better off creating a new variable without the type ambiguity?
## [10][Type check if return object is response object or error object](https://www.reddit.com/r/typescript/comments/jdhr9x/type_check_if_return_object_is_response_object_or/)
- url: https://www.reddit.com/r/typescript/comments/jdhr9x/type_check_if_return_object_is_response_object_or/
---
So I have an API service that returns either data object in case of success, or an object containing error object.

Which means returned object can have two different types, depending on success. Is there a way to make typescript recognise which one it is with an if statement?

    const getStuff = async () =&gt; {
      try {
        return (await axios.get('url')).data as { title: 'title' }
      } catch (e) {
        return { error: e, status: 400 }
      }
    }
    
    const returnedObject = await getStuff()
    
    returnedObject.status // maybe undefined, maybe 400
    
    if (returnedObject.error){
      returnedObject.status // 400
    } else {
      returnedObject.title // 'whatever'
      returnedObject.status // error, status doesn't exists on returnedObject
    }
## [11][TypeORM Sucks!! Something I wanted to talk about since long!](https://www.reddit.com/r/typescript/comments/jcw28f/typeorm_sucks_something_i_wanted_to_talk_about/)
- url: https://www.reddit.com/r/typescript/comments/jcw28f/typeorm_sucks_something_i_wanted_to_talk_about/
---
Let's address it, TypeORM sucks but the community doesn't seem to do anything about it!! So I would like to take the step forward:

1. Querybuilder is very bad, chaotic

2. Relational query is not supported - subquery on relation etc

3. Eager loading is not supported

4. Cannot extend query builder

5. Repository is just sugar coating over bad querybuilder

6. Migration is super ugly and bad, have to write raw alter table queries 50% of the times, don't run migrations in batch

7. Seeders are not supported

8. Documentation asks to enable the db-model syncronization in development, but says to create migrations for production - What the hell?

9. Entity Manager is of no use

10. Polymorphic relations not supported

11. ORM cannot spot basic difference between where() and andWhere()

&amp;#x200B;

Thesre are my two cents! Your feedbacks are most welcome! 

We need better alternative ORM in typescript!
