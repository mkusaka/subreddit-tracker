# typescript
## [1][Who's hiring Typescript developers - August](https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/)
- url: https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/
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
## [2][Getting started with LitElement and TypeScript](https://www.reddit.com/r/typescript/comments/ihi6cc/getting_started_with_litelement_and_typescript/)
- url: https://labs.thisdot.co/blog/getting-started-with-litelement-and-typescript
---

## [3][Created VSCode launch config for debugging Mocha tests that written in TypeScript](https://www.reddit.com/r/typescript/comments/ihjn57/created_vscode_launch_config_for_debugging_mocha/)
- url: https://www.reddit.com/r/typescript/comments/ihjn57/created_vscode_launch_config_for_debugging_mocha/
---
It works fine with ES6 modules. Tested only in Windows 10 machine.

[https://gist.github.com/artem-mangilev/2f45ea18e4f5a38339212f551fa6a85c](https://gist.github.com/artem-mangilev/2f45ea18e4f5a38339212f551fa6a85c)
## [4][ts-migrate: A Tool for Migrating to TypeScript at Scale](https://www.reddit.com/r/typescript/comments/ih1133/tsmigrate_a_tool_for_migrating_to_typescript_at/)
- url: https://medium.com/airbnb-engineering/ts-migrate-a-tool-for-migrating-to-typescript-at-scale-cd23bfeb5cc
---

## [5][styled-components usage with TypeScript (advice/help needed)](https://www.reddit.com/r/typescript/comments/ihkm15/styledcomponents_usage_with_typescript_advicehelp/)
- url: https://www.reddit.com/r/typescript/comments/ihkm15/styledcomponents_usage_with_typescript_advicehelp/
---
Hey (again), I'd posted yesterday regarding some general advice when creating TypeScript components with React (advice was super helpful thanks!). I normally work with styled-components when styling components and was wondering what the standard/best approach was to handle this, say you have some props on a component and only a singular one (or a selected few) will have an affect on the styling how would you go about that efficiently? I was going to create a second type to handle this, but saw the idea of **&lt;Pick&gt;**, thanks Calvin:

**ListItem.tsx (example component with some props):**

    import React from 'react'
    import { Item } from './ListItem.styles'
    
    export interface ListItemProps {
      name: string
      background: string
    }
    
    export const ListItem = ({ name, background }: ListItemProps) =&gt; (
      &lt;Item background={background}&gt;{name}&lt;/Item&gt;
    )

**ListItem.styles.tsx (file containing styles, not sure about the extension):**

    import styled from 'styled-components'
    import { ListItemProps } from './ListItem.component'
    
    export const Item = styled.li&lt;Pick&lt;ListItemProps, 'background'&gt;&gt;`
      background-color: ${props =&gt; props.background};
    `

(trivial example, with lack of keys etc)
## [6][Typing issue with async function](https://www.reddit.com/r/typescript/comments/ihk0mb/typing_issue_with_async_function/)
- url: https://www.reddit.com/r/typescript/comments/ihk0mb/typing_issue_with_async_function/
---
When I added the db query I had to turn the function async, and suddenly I had a type issue that I have been unable to resolve (after about 5 hours now). can any spot what the issue is?

    export default async function formatCallRecordsForPGPromise(
      rawCalldata: CallRecord[],
    ): Promise&lt;SaveableCallRecord[]&gt; {
      const dataSortedbyPhoneNumber: PhoneNumberKeyedCallData = lodash.groupBy(rawCalldata,
        (record: CallRecord) =&gt; record.destination);
    
      const phoneNumberAndGroupedData: [string, CallRecord[]][] = Object
        .entries(dataSortedbyPhoneNumber);
    
      const allData: Promise&lt;SaveableCallRecord[]&gt;[] = phoneNumberAndGroupedData
        .map(async (numberAndCallData: [string, CallRecord[]]) =&gt; {
          const phoneNumber = numberAndCallData[0];
    
          const selectCampaignId = 'SELECT id FROM campaign WHERE phoneNumber = $1';
    
          let campaign_id: number;
          try {
            campaign_id = await pgp.configured.one(
              selectCampaignId, phoneNumber, (queryResult: { id: number }) =&gt; queryResult.id,
            );
          } catch (error) {
            throw new Error(error);
          }
    
          const formattedDataSet: SaveableCallRecord[] = numberAndCallData[1]
            .map((record: CallRecord) =&gt; {
              const {
                date, callerid, destination, description, account, disposition, seconds, uniqueid,
              } = record;
    
              const formattedData: SaveableCallRecord = {
                unique_id: uniqueid,
                caller_id: callerid,
                date: PGPromise.as.date(new Date(date)),
                description,
                account,
                disposition,
                seconds: PGPromise.as.number(Number(seconds)),
                campaign_id: PGPromise.as.number(Number(campaign_id)),
              };
    
              return formattedData;
            });
    
          return formattedDataSet;
        });
    
      await Promise.all(allData);
    
      const allDataFlattened = allData.flat();
    
      return allDataFlattened;
    /*
    const allDataFlattened: Promise&lt;SaveableCallRecord[]&gt;[]
    Type 'Promise&lt;SaveableCallRecord[]&gt;[]' is not assignable to type 'SaveableCallRecord[]'.
      Type 'Promise&lt;SaveableCallRecord[]&gt;' is missing the following properties from type 'SaveableCallRecord': unique_id, caller_id, date, description, and 4 more.ts(2322)
    */
    }
## [7]["Sharing" declaration files between multiple projects](https://www.reddit.com/r/typescript/comments/ihhh59/sharing_declaration_files_between_multiple/)
- url: https://www.reddit.com/r/typescript/comments/ihhh59/sharing_declaration_files_between_multiple/
---
Hey, I've been running into an issue trying to share several d.ts files between my projects. I'm pretty new to TS, so bear with me.

One of my projects, Container, includes an npm package (the 'Core' project). Since we are using a few packages, including Vue, which don't include the correct declaration files, we have to add some "boilerplate" .d.ts files to each project.

For those of you who are familiar with Vue: these are the shims-tsx.d.ts and shims-vue.d.ts files, among others.

Anyway, I'm not really a fan of the idea that we have to copy these files to each new project. Which is why I've been trying to add them to the Core project, and somehow export so that they can be included in each other project "automatically" or through a main import, or something.

I've been trying to understand declaration files, and how I could achieve what I'm trying to do, but I feel like I'm hitting a wall - either because it's impossible, or because I don't understand well enough what I'm doing.

TL;DR: all of my projects need several identical d.ts files because included packages don't have typings, how can I share them or include them automatically?
## [8][Does TS breed dependence on explicit type definitions?](https://www.reddit.com/r/typescript/comments/ih6qmc/does_ts_breed_dependence_on_explicit_type/)
- url: https://www.reddit.com/r/typescript/comments/ih6qmc/does_ts_breed_dependence_on_explicit_type/
---
And if so, is that a bad thing?

Example below. I think i am dependent on annotating my way through it to get what happens. Otherwise I don't think I would be able to do it.

Honestly this is what made me fall in love with the language. But if I get a standard JS job I don't know if I could handle not having a traceable path for how everything is getting changed. 80% of my day would be spent debugging the stuff I was writing just to see if it was returning what I hope it was.

Is it possible to become "type dependent"? And if so how would you guys advise a junior to fix that issue? I guess spend time in regular JS trying to code without explicit typings?

    interface PhoneNumberKeyedCallData {
      [phoneNumber: string]: CallRecord[]
    }
    
    function formatCallRecordsForPGPromise(
      rawCalldata: CallRecord[],
    ): SaveableCallRecordFields[] {
      const dataSortedbyPhoneNumber: PhoneNumberKeyedCallData = lodash.groupBy(rawCalldata,
        (record: CallRecord) =&gt; record.destination);
    
      const phoneNumberAndGroupedData: [string, CallRecord[]][] = Object
        .entries(dataSortedbyPhoneNumber);
    
      const allDataFlattened: SaveableCallRecordFields[] = phoneNumberAndGroupedData
        .map((numberAndCallData: [string, CallRecord[]]) =&gt; {
          const phoneNumber = numberAndCallData[0];
    
          const formattedDataSet: SaveableCallRecordFields[] = numberAndCallData[1]
            .map((record: CallRecord) =&gt; {
              const {
                date, callerid, destination, description, account, disposition, seconds, uniqueid,
              } = record;
    
              const formattedData: SaveableCallRecordFields = {
                unique_id: uniqueid,
                &lt;removed more properties ...&gt;
              };
    
              return formattedData;
            });
    
          return formattedDataSet;
        })
        .flat();
    
      return allDataFlattened}
## [9][[updated benchmark] fork-ts-checker more than 2x slower with project references](https://www.reddit.com/r/typescript/comments/ih72pe/updated_benchmark_forktschecker_more_than_2x/)
- url: https://www.reddit.com/r/typescript/comments/ih72pe/updated_benchmark_forktschecker_more_than_2x/
---
after I fixed some issues and updated the benchmark (hopefully) the last the time, I learned this:

`fork-ts-checker` struggles with reference files (those which get imported) the most in a project reference structure. pure `tsc -b` is more than 2x faster[1] followed by pure `ts-loader`.

So, if you work a lot in reference files you will face significant slow-downs; however in non-reference files, so those which import reference files fork-ts-checker is the fastest setup with type checking.

What to choose? Hard to say but a well structured, DRY monorepo uses a lot references and you work a lot in those. Just think of fe components which are used by different front-end apps as one example, there are much more. fork-ts-checker seems to be really a deal-breaker here. Pure tsc shines with reference compile times but lacks on the main file compiles times. And you can't turn on `transpileOnly: true`. 

So my current gut feeling is to go with the middleground—pure ts-loader with `transpileOnly: true` and let your editor do the type checks. IDK if this will work in the long run and you would't miss type errors, any experience with this?

the new benchmark: [https://github.com/TypeStrong/ts-loader/issues/1157#issuecomment-681080972](https://github.com/TypeStrong/ts-loader/issues/1157#issuecomment-681080972)

[1] depending on the benchmark

Edit: this problem seems to be also with tsloader, i introduced a new babel variant which is by far the best (also in that benchmark)
## [10][Learning TypeScript with React Components (advice/help needed)](https://www.reddit.com/r/typescript/comments/ih013n/learning_typescript_with_react_components/)
- url: https://www.reddit.com/r/typescript/comments/ih013n/learning_typescript_with_react_components/
---
Hey, I just starting learning TypeScript (after watching a few videos this morning) and I'm curious if I'm doing this correctly (there's no bugs, but it feels weird that I'm exporting types and declaring them multiple times). Below is a very trivial React example that contains two components and some dummy data to render (in App):  


**ListItem:**

    import React from 'react'
    
    export type ListItemProps = {
      name: string
    }
    
    const ListItem = ({ name }: ListItemProps) =&gt; &lt;li&gt;{name}&lt;/li&gt;
    
    export default ListItem

**List:**

    import React from 'react'
    import ListItem, { ListItemProps } from '../ListItem/ListItem.component'
    
    type ListProps = { list: ListItemProps[] }
    
    const List = ({ list }: ListProps) =&gt; {
      return (
        &lt;ul&gt;
          {list.map(item =&gt; (
            &lt;ListItem name={item.name} /&gt;
          ))}
        &lt;/ul&gt;
      )
    }
    
    export default List

**App (contains mock data and the actual render):**  


    import React from 'react'
    import List from '../List/List.component'
    import { ListItemProps } from '../ListItem/ListItem.component'
    
    const mockData: ListItemProps[] = [
      { name: 'Item 1' },
      {
        name: 'Item 2'
      }
    ]
    
    const App = () =&gt; (
      &lt;div&gt;
        &lt;List list={mockData} /&gt;
      &lt;/div&gt;
    )
    
    export default App

I would appreciate any and all feedback, I was thinking that creating a separate file such as 'types.ts' would help tackle this (this might not even be a problem, it just feels weird at the moment!). Thanks, Calvin
## [11][Inferring mutually dependent function types?](https://www.reddit.com/r/typescript/comments/igux0f/inferring_mutually_dependent_function_types/)
- url: https://www.reddit.com/r/typescript/comments/igux0f/inferring_mutually_dependent_function_types/
---
In my project there are "converter" types, which convert between two specific other types (`A` and `B`):

    type Converter&lt;A, B&gt; = {
      convertAtoB(value: A): B;
      convertBtoA(value: B): A;
    }

They are easy to type when A and B are specified manually, like `Converter&lt;string, number&gt;`. However, I want them to be inferred. My working attempt looks like this:

    function converter&lt;A, B&gt;(data: Converter&lt;A, B&gt;) {
      return data;
    }
    
    var test = converter({
      convertAtoB: (x: string) =&gt; x, // this one returns a string...
      convertBtoA: (x: number) =&gt; x.toString() // and this accepts a number
    })

The error is correctly shown in the above example (type "string" is not assignable to "number"). So far it's working.

But now I want to define an object containing multiple such converters, and still make it impossible to write incorrect implementations, like this:

    var myConverters = converters({
      stringAndNumber: {
        convertAtoB: (x: string) =&gt; x, 
        convertBtoA: (x: number) =&gt; x.toString() // should error as above
      },
      booleanAndString: {
        convertAtoB: (x: boolean) =&gt; x.toString(),
        convertBtoA: (x: string) =&gt; x === 'true' ? true : false
      }
    })

How do I write that `converters` function, so that TypeScript reports errors just like with the singular `converter` function? I don't think I can accept A and B generic arguments in it, because every "child" converter has its own A and B.

Thanks!
