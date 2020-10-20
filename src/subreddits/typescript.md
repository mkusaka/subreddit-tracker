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
## [2][Generate Google Presentation from Wikipedia article](https://www.reddit.com/r/typescript/comments/je6qg8/generate_google_presentation_from_wikipedia/)
- url: https://www.reddit.com/r/typescript/comments/je6qg8/generate_google_presentation_from_wikipedia/
---
[GSlides Maker](https://github.com/vilmacio/gslides-maker) is a open-source project for create google presentations from wikipedia content like a robot. Try it [CLI version](https://github.com/vilmacio/gslides-maker) and help me with the [web application](https://gslidesmaker.com/). I hope you like it.

https://preview.redd.it/98yuu3ceb3u51.png?width=681&amp;format=png&amp;auto=webp&amp;s=5259f69c148faeab52c9a10af6490d366bfec5c5
## [3][Is this an appropriate time to assert? (Array concat overload mismatch)](https://www.reddit.com/r/typescript/comments/je8q40/is_this_an_appropriate_time_to_assert_array/)
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
## [4][How to make ConstructorParameters&lt;typeof Test&gt;[0] generic?](https://www.reddit.com/r/typescript/comments/je6kpb/how_to_make_constructorparameterstypeof_test0/)
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
## [5][Resolve out union types?](https://www.reddit.com/r/typescript/comments/je3n3i/resolve_out_union_types/)
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
## [6][Type check if return object is response object or error object](https://www.reddit.com/r/typescript/comments/jdhr9x/type_check_if_return_object_is_response_object_or/)
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
## [7][TypeORM Sucks!! Something I wanted to talk about since long!](https://www.reddit.com/r/typescript/comments/jcw28f/typeorm_sucks_something_i_wanted_to_talk_about/)
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
## [8][Type with typescript - cheat sheet with samples](https://www.reddit.com/r/typescript/comments/jd8pze/type_with_typescript_cheat_sheet_with_samples/)
- url: http://goorep.se:1001/changelog/report/rSelect/PAGE_result.htm?alias=guest&amp;set=api&amp;query=Book+pages&amp;$$TArticleBook1.ArticleBookK=7096&amp;link=%5B%5B%229F1E006D78894848838A0970E2FF0BE9zoom%22,%22Object1%22,7096%5D,%5B%229F60D5FEDE8E41CC986C10147F0AD2F7zoom%22,%22TArticleBook1%22,7096%5D,%5B%2271C91DEE3C5A4FDC8EC1114C7C18033Bzoom%22,%22TArticleBook1%22,7096%5D%5D&amp;rows=50
---

## [9][Compile-time regular expressions in TS](https://www.reddit.com/r/typescript/comments/jcruy8/compiletime_regular_expressions_in_ts/)
- url: https://www.reddit.com/r/typescript/comments/jcruy8/compiletime_regular_expressions_in_ts/
---
I'm sorry, I got too excited with TS 4.1 and made compile-time regular expressions

[https://github.com/microsoft/TypeScript/issues/6579#issuecomment-710776922](https://github.com/microsoft/TypeScript/issues/6579#issuecomment-710776922)
## [10][Ultimate GraphQL + Typescript Setup?](https://www.reddit.com/r/typescript/comments/jcmyms/ultimate_graphql_typescript_setup/)
- url: https://www.reddit.com/r/typescript/comments/jcmyms/ultimate_graphql_typescript_setup/
---
Hey! I've just started using graphQL and would advice on whats the best framework setup for Typescript on Frontend (Vue 3) + Backend (Node).

I'm currently using **AWS Amplify** to create the graphQL API and love how it generates the mutations / queries and subscriptions automatically just from a basic schema like:

&gt;type Todo u/model {id: ID!name: String!description: String}

Although it seems lacking in types and I'd ideally like to not get locked into the AWS ecosystem.

I've looked into Vue Apollo but couldn't figure out how to automatically mutations / queries and subscriptions and ideally would like to have this feature on frontend and backend.

I'd rather not have to manually set up all the the mutations / queries and if needed I could then customise them further down the line once I'm more familiar with graphQL. I'd love a solution where I can define schema and then just have default methods for each one that can be accessed like this.

`await graphQL.createTodo( {name: 'Demo todo'})`

`await graphQL.subscribeTodo( {id: '123'})`

Anyway would love to hear opinions on what frameworks are best, so I can get all that juicy goodness from having a predefined schema.

EDIT: To clarify I'll  be using a hosted graphQL instance and Node backend will also be a client. 
## [11][Online ouija board built with typescript and pub/sub - https://spiritboard.ably.dev](https://www.reddit.com/r/typescript/comments/jc8lqq/online_ouija_board_built_with_typescript_and/)
- url: https://www.reddit.com/r/typescript/comments/jc8lqq/online_ouija_board_built_with_typescript_and/
---
I made an online Ouija board!  [https://spiritboard.ably.dev/](https://spiritboard.ably.dev/)   


It is open source so you can edit it and make it your own. [https://github.com/ably/ouija](https://github.com/ably/ouija) I've also written up a very thorough readme which goes over how it all works which is in the repo.
