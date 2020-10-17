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
## [2][Compile-time regular expressions in TS](https://www.reddit.com/r/typescript/comments/jcruy8/compiletime_regular_expressions_in_ts/)
- url: https://www.reddit.com/r/typescript/comments/jcruy8/compiletime_regular_expressions_in_ts/
---
I'm sorry, I got too excited with TS 4.1 and made compile-time regular expressions

[https://github.com/microsoft/TypeScript/issues/6579#issuecomment-710776922](https://github.com/microsoft/TypeScript/issues/6579#issuecomment-710776922)
## [3][Ultimate GraphQL + Typescript Setup?](https://www.reddit.com/r/typescript/comments/jcmyms/ultimate_graphql_typescript_setup/)
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
## [4][Online ouija board built with typescript and pub/sub - https://spiritboard.ably.dev](https://www.reddit.com/r/typescript/comments/jc8lqq/online_ouija_board_built_with_typescript_and/)
- url: https://www.reddit.com/r/typescript/comments/jc8lqq/online_ouija_board_built_with_typescript_and/
---
I made an online Ouija board!  [https://spiritboard.ably.dev/](https://spiritboard.ably.dev/)   


It is open source so you can edit it and make it your own. [https://github.com/ably/ouija](https://github.com/ably/ouija) I've also written up a very thorough readme which goes over how it all works which is in the repo.
## [5][I cannot get callback function arguments typed for the life of me](https://www.reddit.com/r/typescript/comments/jce7ik/i_cannot_get_callback_function_arguments_typed/)
- url: https://www.reddit.com/r/typescript/comments/jce7ik/i_cannot_get_callback_function_arguments_typed/
---
    type ConfigCallback = (...args: never[]) =&gt; AxiosRequestConfig
    
    class ApiService {
      private ctx: Context
      constructor(ctx: Context) {
        this.ctx = ctx
      }
    
      buildRequest&lt;T&gt;(config: ConfigCallback) {
        return async (...ar: Parameters&lt;typeof config&gt;): Promise&lt;T&gt; =&gt; {
          const resolvedConfig =
            typeof config === 'function' ? config.apply(this, ar) : config
          try {
            return await this.ctx.$axios.$request&lt;T&gt;(resolvedConfig)
          } catch (e) {
            throw e
          }
        }
      }
    
      commons: this.buildRequest&lt;Commons[]&gt;((type: string) =&gt; ({
        url: versioned('commons'),
        method: 'get',
        params: { type },
      }))
    
      someMethod(){
        // Argument of type '"string"' is not assignable to type 'never'
        this.commons('someType')
      }
    
    }

Some might say I should do (type: string) =&gt; this.buildRequest&lt;&gt;() and remove all the trickery in buildRequest(), however "commons" or something similar doesn't have to be a function, so I want to avoid calling it like this.commons('params')().
## [6][swc with Jest vs swc with Babel](https://www.reddit.com/r/typescript/comments/jcdie9/swc_with_jest_vs_swc_with_babel/)
- url: https://www.reddit.com/r/typescript/comments/jcdie9/swc_with_jest_vs_swc_with_babel/
---
So I've recently been seeing some impressive claims about various pieces of JS tooling written in languages other than JavaScript, so I was keen to give some of them a try.

Today I spent some of my work 10% time experimenting with swapping babel out of the toolchain of a current work project (a small typescript react application that was originally bootstrapped with Create React App) and replacing it with SWC.

Although test times are at present not that bad, anticipating the project getting bigger, I was particularly interested in seeing whether I could improve the test times (also jest runs are easy to benchmark).

The result was a little bit disappointing. Using `yarn add jest-benchmark` I found that our tests ran at an average of about 6.75 seconds per run using `babel` (average over 10 runs), compared to 6.5 using `swc`. Now I appreciate that transpilation isn't the only bottleneck in a test run, but given the wild claims I've seen everywhere about these tools (and all the blog posts which talk about it significantly speeding up your test runs), is this really what I'd expect to see or have I misconfigured things somehow? I've already raised an issue on the `swc` Github repository already only to have my issue immediately closed with a spectacularly unhelpful reply ([https://github.com/swc-project/swc/issues/1172](https://github.com/swc-project/swc/issues/1172)).

To configure `swc` with `jest`, I installed `jest-swc` and changed the `jestConfig` transform block to the following:

    "transform": {
      "^.+\\.(js|jsx|ts|tsx)$": "swc-jest",
      "^.+\\.css$": "&lt;rootDir&gt;/config/jest/cssTransform.js",
      "^(?!.*\\.(js|jsx|ts|tsx|css|json)$)": "&lt;rootDir&gt;/config/jest/fileTransform.js"
    },

I then added the following `.swcrc` to the root of the repository:

    {
      "jsc": {
        "target": "es2017",
        "parser": {
          "syntax": "typescript",
          "tsx": true,
          "decorators": false,
          "dynamicImport": false
        },
        "transform": {
          "react": {
            "pragma": "React.createElement",
            "pragmaFrag": "React.Fragment",
            "throwIfNamespace": true,
            "development": false,
            "useBuiltins": false
          },
          "hidden": {
            "jest": true
          }
        }
      },
      "module": {
        "type": "commonjs",
        "strict": false,
        "strictMode": true,
        "lazy": false,
        "noInterop": false
      }
    }
## [7][Updating relationship objects (using typeorm)](https://www.reddit.com/r/typescript/comments/jcct5w/updating_relationship_objects_using_typeorm/)
- url: https://www.reddit.com/r/typescript/comments/jcct5w/updating_relationship_objects_using_typeorm/
---
I have two entities:

    @Entity({ name: 'problem' })
    export class Problem {
      @PrimaryGeneratedColumn('uuid')
      id: string;
       
      @Column({ type: 'boolean', default: false })
      isSolved: boolean;
    
      @ManyToOne(() =&gt; Student, student =&gt; student.problems,  { nullable: true })
      student: Student;
    }
    
    @Entity({ name: 'student' })
    export class Student {
      @PrimaryGeneratedColumn('uuid')
      id: string;
    
      @OneToMany(() =&gt; Problem, problem =&gt; problem.student, { cascade: ['insert', 'update'], nullable: true }) 
      problems: Problem[];
    }

Each student and problem can exist independently; when creating a new problem, I want to update the student relationship 'problems':

 

    const students : Student[] = await getConnection().
 getRepository(Student).find({ relations: ['problems'] });
    
        students.forEach(async function (student, index) {
          // Student is free can solve problems
          if (student.problems === undefined || student.problems.length == 0) {        
            newProblem.student = student;
            // RELATIONSHIP NOT UPDATED TO STUDENT EVEN IF A PROBLEM WAS ASSIGNED
            await getConnection().getRepository(Problem).save(newProblem).then(e =&gt; ProblemDTO.fromEntity(e));
            return newProblem;
          }

How can I update both sides (student + problem)?
## [8][Bringing TypeScript types at runtime with TypeOnly](https://www.reddit.com/r/typescript/comments/jbpbnp/bringing_typescript_types_at_runtime_with_typeonly/)
- url: https://medium.com/@paleo.said/c317e9dd8880
---

## [9][What is the proper way to extend custom headers on a fetch request?](https://www.reddit.com/r/typescript/comments/jbxqh1/what_is_the_proper_way_to_extend_custom_headers/)
- url: https://www.reddit.com/r/typescript/comments/jbxqh1/what_is_the_proper_way_to_extend_custom_headers/
---
I was getting "No overload matches this call" until I used `new Header()` to sneak my custom headers in via object.

Another option would be to use a global declaration file to extend the fetch global variable.

Should one turn to declaration files for custom headers in this case?
## [10][Is this a bug in Typescript on VS code?](https://www.reddit.com/r/typescript/comments/jc4zao/is_this_a_bug_in_typescript_on_vs_code/)
- url: https://www.reddit.com/r/typescript/comments/jc4zao/is_this_a_bug_in_typescript_on_vs_code/
---
Hi, Is this a bug in Typescript on VS code?

or do I have to configure on my editor ?

&amp;#x200B;

basically, the type returned by the \`requestsHistory\` function should return the type of

DeliveryRequestHistoryResponseData (right).

but on the left side, it looks like my editor isn't working properly since the type inference is lost for some reason.

&amp;#x200B;

[the issue](https://preview.redd.it/v48bpqkzket51.png?width=1368&amp;format=png&amp;auto=webp&amp;s=4bd3bb6637388f9a1190014fcf7ac0b9b67b8a72)

EDIT:

    export interface DeliveryRequestModel {
    	... // ommited for brevity
    }
    
    export interface DeliveryRequestHistoryResponseData {
    	data: {
    		delivery_requests: DeliveryRequestModel[]
    	},
    	success: boolean
    }

&amp;#x200B;

tsconfig.json that was generated by the nativescript-cli

&amp;#x200B;

[tsconfig.json](https://preview.redd.it/am18hhd2oet51.png?width=402&amp;format=png&amp;auto=webp&amp;s=abb5bc4094ac1ec80b87b1ae8ef7f420c0cfed96)
## [11][How to use an existing type or empty string?](https://www.reddit.com/r/typescript/comments/jbynnz/how_to_use_an_existing_type_or_empty_string/)
- url: https://www.reddit.com/r/typescript/comments/jbynnz/how_to_use_an_existing_type_or_empty_string/
---
I'm using angular v9 w/ typescript v 3.5.3

I feel like I run into this issue often and I never really understand the typescript rule for these cases. The typescript error I'm getting:

&gt;error TS2345: Argument of type '"" | "Learn More" | "Be Healthy" | "Be Fit"' is not assignable to parameter of type 'Goal'.

&amp;#x200B;

Some example code:

    // this type is used else where throughtout the code
    export type Goal = 'Learn More' | 'Be Healthy' | 'Be Fit';
    
    // some of the component code
    ...
    class SomeComponent {
      goalList: Goal[] = ['Learn More', 'Be Healthy', 'Be Fit'];
      // 
      goalSelection$ = new BehaviorSubject&lt;Goal|''&gt;('');
    
      ...
    
      // error occurs in this method
      goalSelectIsValid() {
        const goal = this.goalSelection$.getValue();
        return this.goalList.includes(goal);
      }
    }
    
    ...
    // template example
    &lt;app-dropdown #goalDropdown&gt;
      &lt;div currentlySelected&gt;
        {{ (goalSelection$ | async) }}
      &lt;/div&gt;
      &lt;div dropdown&gt;
        &lt;div class="layout vertical options"&gt;
          &lt;button mat-button *ngFor="let item of goalList" (click)="updateGoalSelection(item); goalDropdown.open = false"&gt;
            {{ item }}
          &lt;/button&gt;
        &lt;/div&gt;
      &lt;/div&gt;
    &lt;/app-dropdown&gt;
    &lt;button type="submit" [disabled]="!goalSelectIsValid()"&gt;

So when I'm writing this out, my thought process is that `this.goalList.includes` is just validating the value of `this.goalSelection$.getValue()` .. I know that `goalSelection$` can be an empty string, but I'm not setting any value that invalidates the typing of `this.goalList`.

This always trips me up, and I know it's a lack of understanding of typescript. 

In these scenarios, how should I better approach typing these variables?
