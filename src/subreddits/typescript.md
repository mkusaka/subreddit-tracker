# typescript
## [1][Who's hiring Typescript developers - September](https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/)
- url: https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/
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
## [2][First impressions of Prisma](https://www.reddit.com/r/typescript/comments/imapom/first_impressions_of_prisma/)
- url: https://www.reddit.com/r/typescript/comments/imapom/first_impressions_of_prisma/
---
I've been messing around in Prisma for some hours now. The syntax for defining your  schema is compact, all of the tables are defined in a single schema file.  This schema file can either be used to create the tables in the DB, or you can introspect an existing DB to generate the models for the schema.

The syntax for queries is easy to remember and autocompletion makes fields in the queries discoverable.

I guess I would just encourage you to try it out and see how you feel. You can add it to an existing project and start running queries in like 10 minutes (by introspecting your DB).

Most of my experience is TypeORM in NestJS. I was sort of discouraged from doing much on the back end because of the complexity of TypeORM. Now I feel like I can build an application that I am willing to maintain.

I remember being frustrated waiting for Prisma 2 to come out and frustrated that Prisma 1 wasn't something that I could use moving forward after going through a lengthy tutorial on Prisma 1. Regardless, I am glad the creators had ambitions for a better product.

I hope this project is well maintained moving forward because I want to keep using it. Which is why I am sort of evangelizing on here. I want to keep using it because of the syntax and autocompletion.
## [3][TypeScript ORM with no query builder, supporting full SQL queries](https://www.reddit.com/r/typescript/comments/iltfzj/typescript_orm_with_no_query_builder_supporting/)
- url: https://github.com/Seb-C/kiss-orm#kiss-orm
---

## [4][How do I extend a class prototype with TypeScript type checking?](https://www.reddit.com/r/typescript/comments/imdds1/how_do_i_extend_a_class_prototype_with_typescript/)
- url: https://stackoverflow.com/questions/63739143/how-do-i-extend-a-class-prototype-with-typescript-type-checking
---

## [5][TypeORM: affected, generatedMaps and raw property of UpdateResult is undefinde](https://www.reddit.com/r/typescript/comments/imchtw/typeorm_affected_generatedmaps_and_raw_property/)
- url: https://www.reddit.com/r/typescript/comments/imchtw/typeorm_affected_generatedmaps_and_raw_property/
---
I use TypeORM with PostgreSQL and when I make an update request I can't find out if the operation was successful, but the data is updated. What could be the reason?
## [6][How can I inherit a type from the value of a parameter from an object?](https://www.reddit.com/r/typescript/comments/im87di/how_can_i_inherit_a_type_from_the_value_of_a/)
- url: https://www.reddit.com/r/typescript/comments/im87di/how_can_i_inherit_a_type_from_the_value_of_a/
---
My sample function:

    const myFunc = (options: {
      args?: { [s: string]: any },
      callback: (args: { [s: string]: any }) =&gt; void,
    }): void =&gt; {
      const args = options.hasOwnProperty(‘args’) ? options.args : {};
      options.callback(args);
    };

Calling my sample function (See the comment below):

    myFunc({
      args: {
        test: 1,
      },
      callback: (args): void =&gt; {
        // I want `args` to inherit the same type as the `args` parameter above: `{test: number}`
      },
    });

Edit: Had a small type in my code
## [7][Greenfield Typescript project, what are current best practices?](https://www.reddit.com/r/typescript/comments/ilugvz/greenfield_typescript_project_what_are_current/)
- url: https://www.reddit.com/r/typescript/comments/ilugvz/greenfield_typescript_project_what_are_current/
---
I'm starting a Typescript project from scratch. The project will definitely become complex, so I'd like to architect it well upfront. I want to make sure I'm using the full capabilities of the language and libraries/frameworks if needed.

I'm looking at things like dependency injection, composition over inheritance, functional patterns etc. Testability is a huge must. I like dependency injection patterns because they make unit testing with mocks very easy.

Can anyone recommend good libraries for this or good open-source projects to look at which are a good example of e2e? This project will be an isomorphic library so there's no built-in server or overarching framework.
## [8][error TS2749: 'IRequestWithUser' refers to a value, but is being used as a type here](https://www.reddit.com/r/typescript/comments/ilxkki/error_ts2749_irequestwithuser_refers_to_a_value/)
- url: https://www.reddit.com/r/typescript/comments/ilxkki/error_ts2749_irequestwithuser_refers_to_a_value/
---
I am extending \`Request\` in \`express\` library to contain a user property:

    import { Request } from 'express';
    
    export default interface RequestWithUser extends Request {
      user: {
        user_name: string;
        password: string;
      }
    }

The title error appears in the first parameter annotation:

    import IRequestWithUser from '../shared/interfaces/isRequestWithUser';
    
    const router: Router = Router();
    
    router.post('/myRoute', async (req: IRequestWithUser, res: Response) =&gt; {
    
    /*
    error TS2749: 'IRequestWithUser' refers to a value, 
    but is being used as a type here. Did you mean 
    'typeof IRequestWithUser'?
    */

I don't believe interfaces are values. They should purely be types. So what is causing this error?

**Also tried**

`typeof IRequestWithUser` This results in \`No overload matches this call\`
## [9][Is it wrong to use namespaces? ES5 imports feel inferior to me compared to namespaces](https://www.reddit.com/r/typescript/comments/illv8o/is_it_wrong_to_use_namespaces_es5_imports_feel/)
- url: https://www.reddit.com/r/typescript/comments/illv8o/is_it_wrong_to_use_namespaces_es5_imports_feel/
---
I have the following code that works perfectly.

    export namespace Vid {
      export const SPLASH = require('./Splash.mp4')
      export const SEQUENCE_1 = require('./Sequence01.mp4')
      export const SEQUENCE_2 = require('./Sequence02.mp4')
      export const SEQUENCE_3 = require('./Sequence03.mp4')
    }

So if I want to import a Video in my file, I type `Vid.`, and I get autocomplete for all my available videos. This is great for me.

_____

Unfortunately namespaces are deprecated for some reason and I'm supposed to export the const directly.

      export const SPLASH = require('./Splash.mp4')
      export const SEQUENCE_1 = require('./Sequence01.mp4')
      export const SEQUENCE_2 = require('./Sequence02.mp4')
      export const SEQUENCE_3 = require('./Sequence03.mp4')

This is NOT ideal. Because now when I'm importing files, I need to remember the file name itself instead of just typing `Vid.` and being able to loop through my files. There will also be higher namespace collisions with other exported consts.

_____

I've seen someone suggest to do the following to replicate namespaces.

    const SPLASH = require('./Splash.mp4')
    const SEQUENCE_1 = require('./Sequence01.mp4')
    const SEQUENCE_2 = require('./Sequence02.mp4')
    const SEQUENCE_3 = require('./Sequence03.mp4')

    export const Vid = { SPLASH, SEQUENCE_1, SEQUENCE_2, SEQUENCE_3 }

However after using it for awhile, I still find it inferior.

1) `Go to Declaration` no longer works properly. When I try to lookup the declaration, it points me to `export const Vid`, then I have to look up the declaration a second time to jump to my code. This is really annoying.

2) It no longer prompts me which consts are unused. With namespaces, I'm informed which ones are used and unused. With consts, it doesn't seem to know anymore.

- [Namespaces](https://i.imgur.com/ChJwD9K.png)
- [ES5](https://i.imgur.com/4f1qyMP.png)

_____

I've reverted to using namespaces for the time being, because in my opinion they are superior. Can someone tell me what potential problems I could have by continuing to use namespaces? Does anyone know how long more they will be supported for?

**EDIT: Too many people are answering by telling me the pros of ES imports. That is not the question. The question is:**

- **What is wrong with using namespaces?**

_____

EDIT 2: After 65 comments, someone has finally provided the right answer. **My first statement is wrong and namespaces are not being deprecated. Therefore there is nothing wrong with using namespaces.**

https://www.reddit.com/r/typescript/comments/illv8o/is_it_wrong_to_use_namespaces_es5_imports_feel/g3u6fu1/
## [10][Need help to understand how assignments in arrays/objects are done](https://www.reddit.com/r/typescript/comments/ilkrk4/need_help_to_understand_how_assignments_in/)
- url: https://www.reddit.com/r/typescript/comments/ilkrk4/need_help_to_understand_how_assignments_in/
---
Hello,

I have a problem that I think is pretty simple, but I can't get my head around the actual mechanics and why it does not work. Here is my issue:

I have an object (pModel) that is passed (by reference, I assume) in a function. That object contains an array of objects (geometries) and those objects have a member variable called isHighlighted (boolean). So basically, pModel.geometries\[i\].isHighlighted = true or false.

What I am doing is that I am looping through the geometries to toggle the isHighlighted variable. I am currently doing it this way:

    for (let geometry of pModel.geometries)
    {
        console.log(geometry.isHighlighted); // returns false
        console.log(geometry);  // isHighlight inside object is false
        geometry.isHighlighted = !geometry.isHighlighted;
        console.log(geometry.isHighlighted); // returns true
        console.log(geometry);  // isHighlighted inside object is still false
    }

I don't understand why isHightlighted is not changed in the second console.log(geometry). Also, if I write geometry.isHighlighted = true, it works. the second console.log(geometry) will have isHighlighted = true. I don't understand that either...
## [11][Make third party package as global visible to all other files](https://www.reddit.com/r/typescript/comments/ilowzn/make_third_party_package_as_global_visible_to_all/)
- url: https://www.reddit.com/r/typescript/comments/ilowzn/make_third_party_package_as_global_visible_to_all/
---
Hello,

I installed momentjs and want to the project automatically access to the function everywhere without explicitly importing. Is this possible in typescript?

Thanks
