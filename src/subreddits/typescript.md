# typescript
## [1][Who's hiring Typescript developers - June](https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/)
- url: https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/
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
## [2][A lightweight yet powerful dependency injection framework using modern TypeScript and decorators](https://www.reddit.com/r/typescript/comments/h0ogq0/a_lightweight_yet_powerful_dependency_injection/)
- url: https://github.com/H1ghBre4k3r/dependory
---

## [3][Foal TS - June release (version 1.9) - Auth with MongoDB/TypeORM, auth with REST APIs, handy shortcuts, CLI improvements](https://www.reddit.com/r/typescript/comments/h0tzhe/foal_ts_june_release_version_19_auth_with/)
- url: https://www.reddit.com/r/typescript/comments/h0tzhe/foal_ts_june_release_version_19_auth_with/
---
Foal TS framework version 1.9 is officially released!

[Foal TS - June release \(version 1.9\)](https://preview.redd.it/cbw3ao9078451.png?width=1250&amp;format=png&amp;auto=webp&amp;s=f1281bba1a03e446043f4cc901173ea9cf6a4016)

This version brings many improvements and fixes bugs.

1. **MongoDB can now be fully used with TypeORM**. The new `fetchMongoDBUser` allows you to retrieve the MongoDB user whether you use sessions or stateless JWTs.

[Authentication with JWT\/MongoDB\/TypeORM](https://preview.redd.it/st1vyyzi88451.png?width=363&amp;format=png&amp;auto=webp&amp;s=43047d0979f8707e848186f8142ea81109080ee4)

Docs: [https://foalts.gitbook.io/docs/topic-guides/authentication-and-access-control/quick-start](https://foalts.gitbook.io/docs/topic-guides/authentication-and-access-control/quick-start)

2. **\[Shortcut\] The request params and body** are passed as **second** and **third arguments** of the controllers.

[Quick access to the request body and params](https://preview.redd.it/y12146c998451.png?width=525&amp;format=png&amp;auto=webp&amp;s=1bd65f3567295334f29622b0346c1a3fe6f585d6)

Docs: [https://foalts.gitbook.io/docs/topic-guides/architecture/controllers#the-controller-method-arguments](https://foalts.gitbook.io/docs/topic-guides/architecture/controllers#the-controller-method-arguments)

3. \[CLI\] The CLI `generate` command has been extended to support subdirectories. You can now fully generate your controller (or service) architecture from the command line

[CLI commands](https://preview.redd.it/rbctunjr98451.png?width=649&amp;format=png&amp;auto=webp&amp;s=4f98058b17d5efce6efd22fbd12e0b4d87aeab26)

[Architecture generated](https://preview.redd.it/5vfxegbt98451.png?width=651&amp;format=png&amp;auto=webp&amp;s=3173ecd8a3925847c0f2f38cc572045345192821)

[app.controller.ts updated](https://preview.redd.it/hnjls0vq98451.png?width=645&amp;format=png&amp;auto=webp&amp;s=71cd1772b42804a0981cbdb8aa7ceae981d0fa30)

[api.controller.ts generated](https://preview.redd.it/gbiqvaos98451.png?width=644&amp;format=png&amp;auto=webp&amp;s=4ab3e7f480a60c528236976fce39170725c2e743)

Docs: [https://foalts.gitbook.io/docs/topic-guides/cli-and-development-environment/code-generation#create-a-controller](https://foalts.gitbook.io/docs/topic-guides/cli-and-development-environment/code-generation#create-a-controller)

Foal, in a few words, it's a **Node.js framework**:

* written in **TypeScript**
* provided **with batteries included** (Auth, OpenAPI, GraphQL, ORM, CLI, scripts, Cloud file storage, etc)
* and with a **simple and intuitive architecture** (no magic, no over-engineering).

And the must: it has more than 11,000 lines of documentation.

[https://foalts.org](https://foalts.org/)

[https://github.com/FoalTS/foal](https://github.com/FoalTS/foal)

[https://foalts.gitbook.io/docs/](https://foalts.gitbook.io/docs/)

\#TypeScript #JavaScript #NodeJS #FoalTS #CLI #Auth #Authentification #JWT
## [4][Passing generic to function without calling it (1st class function)](https://www.reddit.com/r/typescript/comments/h0qrqx/passing_generic_to_function_without_calling_it/)
- url: https://www.reddit.com/r/typescript/comments/h0qrqx/passing_generic_to_function_without_calling_it/
---
I'm trying to pass a 1st class function to a another function while have the generic inferred without calling it, but I can't get typescript to infer the type, only if I explicitly call the function.

    sql.parseTable = &lt;T&gt;(res: QueryResult): T[] =&gt;
        pipe(sqlResRows, map(convertToCamelCase))(res)
    
    const songs = await sql.query({
        sql: SELECT_SONGS,
        parse: sql.parseTable, // in this line the generic can't be inferred
        parse: res =&gt; sql.parseTable(res), // the generic correctly inferred
    })

I would prefer the first one as it's cleaner and would mean i don't have to change all my graphql resolvers, but the generic can't be inferred.

Any ideas?

TS: 3.9.5, Node: 14.4.0
## [5][Can i use typescript in express.js with functional programming, will everything works fine ?](https://www.reddit.com/r/typescript/comments/h0eqsx/can_i_use_typescript_in_expressjs_with_functional/)
- url: https://www.reddit.com/r/typescript/comments/h0eqsx/can_i_use_typescript_in_expressjs_with_functional/
---

## [6][Help with TS-Mongoose](https://www.reddit.com/r/typescript/comments/h0e3ff/help_with_tsmongoose/)
- url: https://www.reddit.com/r/typescript/comments/h0e3ff/help_with_tsmongoose/
---
The code:

The main file:

`server.ts`

    import * as Fastify from "fastify";
    import * as Dotenv from "dotenv";
    import * as mongoose from "mongoose";
    import routes from "./routes";
    import fastify = require("fastify");
    
    Dotenv.config();
    
    const server = Fastify({ logger: true });
    
    server.register(routes);
    
    const connectToMongo = async (): Promise&lt;void&gt; =&gt; {
      try {
        await mongoose.connect(String(process.env.MONGODB_URI), {
          useNewUrlParser: true,
          useUnifiedTopology: true,
          useCreateIndex: true,
        });
      } catch (err) {
        throw err;
      }
    };
    
    const startServer = async (): Promise&lt;void&gt; =&gt; {
      try {
        await server.listen(Number(process.env.PORT), process.env.ADDR);
        server.log.info(`Started Server @ ${process.env.ADDR}:${process.env.PORT}`);
        await connectToMongo();
        server.log.info("Connected to Mongo...");
      } catch (err) {
        server.log.error(err);
        process.exit(-1);
      }
    };
    
    startServer();

A sample schema:

`schema/task.ts`

    import { createSchema, Type, typedModel } from "ts-mongoose";
    
    const taskSchema = createSchema(
      {
        entryDate: Type.date({ required: true }),
        particulars: Type.string({ required: true }),
        finishDate: Type.date({ required: false }),
      },
      { collection: "Task", timestamps: true }
    );
    
    export default taskSchema;

Models:

`models.ts`

    import { typedModel } from "ts-mongoose";
    import aggregateSchema from "./schemata/aggregate";
    import taskSchema from "./schemata/task";
    import transactionSchema from "./schemata/transaction";
    
    const models = {
      Aggregate: typedModel("Aggregate", aggregateSchema),
      Task: typedModel("Task", taskSchema),
      Transaction: typedModel("Transaction", transactionSchema),
    };
    
    export default models;

Routes file (WIP):

`route.ts`

    import * as Fastify from "fastify";
    import { Server, IncomingMessage, ServerResponse } from "http";
    import handlers from "./handlers";
    
    const routes = (
      server: Fastify.FastifyInstance&lt;Server, IncomingMessage, ServerResponse&gt;,
      opts: Fastify.RegisterOptions&lt;Server, IncomingMessage, ServerResponse&gt;,
      done: (err?: Fastify.FastifyError | undefined) =&gt; void
    ): void =&gt; {
      server.get("/transaction", opts, handlers.getTransactionHandler);
    
      server.post("/transaction", opts, async (req, res) =&gt; {});
    
      server.get("/aggregate", opts, async () =&gt; {});
    
      server.get("/task", opts, async () =&gt; {});
      server.post("/task", opts, async () =&gt; {});
      server.put("/task/:id", opts, async () =&gt; {});
    
      done();
    };
    
    export default routes;

Handlers file (WIP):

`handlers.ts`

    import { FastifyRequest, FastifyReply } from "fastify";
    import { ServerResponse } from "http";
    
    import models from "./models";
    const { Aggregate, Task, Transaction } = models;
    
    const handlers = {
      getTransactionHandler: async (
        req: FastifyRequest,
        _: FastifyReply&lt;ServerResponse&gt;
      ) =&gt; {
        const x = new Task({ entryDate: Date() });
        await x.save();
        return "Hello";
      },
    };
    
    export default handlers;

In `handlers.ts`, when I instantiate a `Task` object without one of the required properties, no error is thrown at compile time. Error is thrown at runtime. Also, intellisense/autocomplete of VS Code does not work with the property names. Is that expected behaviour? I expected almost class-like typechecking behaviour, like when we instantiate a new object using constructor.
## [7][Help understanding why these type assertions aren't working](https://www.reddit.com/r/typescript/comments/h0p61x/help_understanding_why_these_type_assertions/)
- url: https://www.reddit.com/r/typescript/comments/h0p61x/help_understanding_why_these_type_assertions/
---
    // works
    const configData = this.steps.reduce((accumulator: Partial&lt;ConfigData&gt;,
     currentStepInstance: WizardSteps): Partial&lt;ConfigData&gt; =&gt; {...}, {});
    
    return configData as ConfigData;
       
    
    // doesn't work
    const configData: ConfigData = this.steps.reduce((accumulator: Partial&lt;ConfigData&gt;, 
    currentStepInstance: WizardSteps): Partial&lt;ConfigData&gt; =&gt; {...}, {});
    
    return configData;
    
    
    // doesn't work
    const configData as ConfigData = this.steps.reduce((accumulator: Partial&lt;ConfigData&gt;, 
    currentStepInstance: WizardSteps): Partial&lt;ConfigData&gt; =&gt; {...}, {});
    
    return configData;
    
    
    // doesn't work
    const configData&lt;ConfigData&gt; = this.steps.reduce((accumulator: Partial&lt;ConfigData&gt;, 
    currentStepInstance: WizardSteps): Partial&lt;ConfigData&gt; =&gt; {...}, {});
    
    return configData;

My understanding of type assertions is not complete.

I initially thought the 2nd version should work, but the error message made sense, it looks like `reduce` will be returning a Partial of the type with no explicit acknowledgement that all k/v pairs will be there.

Ok, so a type assertion should work. But both `as` and `&lt;&gt;` syntax have failed and I don't understand why.

Only the first example worked where the type was initially allowed to be inferred, only being asserted later as the object is returned. Personally I don't like that, I want to be explicit earlier on if possible.

Thanks for any feedback on this.
## [8][Can I code npm package with Typescript? If I can, how to do it?](https://www.reddit.com/r/typescript/comments/h0kbef/can_i_code_npm_package_with_typescript_if_i_can/)
- url: https://www.reddit.com/r/typescript/comments/h0kbef/can_i_code_npm_package_with_typescript_if_i_can/
---
I want to code a package about styled-components. I haven't published package before and using Typescript generally. I know installing packages with Typescript is little bit trying. How about coding a npm package with Typescript ? If you want to contribute contact w/me
## [9][What's a good charting library that I can use with TypeStyle &amp; React?](https://www.reddit.com/r/typescript/comments/h02b26/whats_a_good_charting_library_that_i_can_use_with/)
- url: https://www.reddit.com/r/typescript/comments/h02b26/whats_a_good_charting_library_that_i_can_use_with/
---
My understanding is that most charting libraries rely on traditional style rules, so wouldn't play well with TypeStyle (or styled-components) - is that accurate? If so, what recommendations do people have?
## [10][Does typescript works with pnpm?](https://www.reddit.com/r/typescript/comments/h043pc/does_typescript_works_with_pnpm/)
- url: https://www.reddit.com/r/typescript/comments/h043pc/does_typescript_works_with_pnpm/
---
Pnpm is a package manager that uses symlinks for node modules. Yesterday I started learning typescript but the code doesn't runs with pnpm. Using the npm instead of pnpm works. I stumbled upon few stackoverflow and github questions but none had the answer.
## [11][Help with return values out of `reduce` method](https://www.reddit.com/r/typescript/comments/h06rlf/help_with_return_values_out_of_reduce_method/)
- url: https://www.reddit.com/r/typescript/comments/h06rlf/help_with_return_values_out_of_reduce_method/
---
I have an interface `ConfigData` that identifies a couple of key/value pairs (`languageCode: string`, and `numberOfRepeats: number`).

Not every step object being reduced will contribute to the accumulator. Of 5 steps passing through, only 2 will add a k/v pair.

I want `const configData` in the example below to receive a `ConfigData` object after the reduce method evaluates. But my code is incorrect as is in (at least) two places:

       run(): ConfigData {
    
          const configData = this.steps.reduce((accumulator: Partial&lt;ConfigData&gt;, currentStep: WizardSteps): ConfigData =&gt; {
             // instruct and ask for feedback
             currentStep.explain();
             const rawUserResponse: string = currentStep.prompt();
    
             if (currentStep.needsInputValidation) {
                const validatedInput: string = currentStep.validate(rawUserResponse);
                
                return {
                   [currentStep.configDataKey] : validatedInput
    // const validatedInput: string
    // Type '{ [x: string]: string; }' is missing the following 
    // properties from type 'ConfigData': languageCode, 
    // numberOfRepeats ts(2739)
                };
             }
    
             if (currentStep.needsFileValidation) {
                // blocks until file is validated
                currentStep.validate();
             }
    
    
          }, {});
    
          return configData;
    // const configData: Partial&lt;ConfigData&gt;
    // Type 'Partial&lt;ConfigData&gt;' is not assignable to type 'ConfigData'.
    // Property 'languageCode' is optional in type 'Partial&lt;ConfigData&gt;' 
    // but required in type 'ConfigData'. ts(2322)
       }

I'm having a hard time with both of these. The second one is at least comprehensible but I'm lost on how to fix both. Anyone know how?
