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
## [2][Building a Game With TypeScript. Drawing Grid 5/5](https://www.reddit.com/r/typescript/comments/iqn5nn/building_a_game_with_typescript_drawing_grid_55/)
- url: https://medium.com/gregsolo/building-a-game-with-typescript-drawing-grid-5-5-49454917b3af?source=friends_link&amp;sk=84924dff6c15ff16f3a01f6d35fe48d1
---

## [3][Enums vs Object.freeze](https://www.reddit.com/r/typescript/comments/iqc2av/enums_vs_objectfreeze/)
- url: https://www.reddit.com/r/typescript/comments/iqc2av/enums_vs_objectfreeze/
---
I have read sentiment on here that Enums may have been a directional misfit. 

And because I tend to use string Enums to avoid string arguments, I think I may do what one poster said and revert to `Object.freeze` again. 

They are essentially the same right? Or is there a strong case for one over the other?
## [4][Any React libraries you would recommend avoiding when working with Typescript?](https://www.reddit.com/r/typescript/comments/iqfq41/any_react_libraries_you_would_recommend_avoiding/)
- url: https://www.reddit.com/r/typescript/comments/iqfq41/any_react_libraries_you_would_recommend_avoiding/
---
I had an issue with some library not long ago, and I was wondering if there were any other libraries that may become a pain in the ass because it's not well integrated with Typescript, or it's poorly designed and doesn't really use Typescript's features very well at all.
## [5][Avoid Export Default](https://www.reddit.com/r/typescript/comments/ipou7q/avoid_export_default/)
- url: https://basarat.gitbook.io/typescript/main-1/defaultisbad
---

## [6][Writing a schema parser in typescript](https://www.reddit.com/r/typescript/comments/iq003l/writing_a_schema_parser_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/iq003l/writing_a_schema_parser_in_typescript/
---
Hi,

In python I can write the following code:

	class Schema(dict):
		table = 
		schema = {}

		def __init__(self, **kwargs):
			super().__init__()

			if not isinstance(self.table, str):
				raise TypeError('Table should be a string name')

			if not isinstance(self.schema, dict):
				raise TypeError('Malformed schema.')
			// Write schema validation here

			for key, keyProps in self.schema.items():

				if key in kwargs:
					if isinstance(kwargs[key], keyProps['type']):
						self[key] = kwargs[key]
					else:
						raise TypeError(f"Mismatched type for key '{key}': Expected '{keyProps['type']}' Got '{type(kwargs[key])}'")
				else:
					if keyProps['required'] or keyProps['default']:
						self[key] = keyProps.get('default') or keyProps['type']()

	class Entity(dict):
		schema = {
			'name': {
				'type': str,
				'required': True,
			},
			'count': {
				'type': int,
				'required': True,
				'default': 10
			},
		}

This allows me to model various entities as per some well defined schema, as well as ensure that the schema is followed by all instances of entities. Additionally, I can add helper functions to reuse functionality for reading/writing to databases/streams/event queues etc

How do I achieve something similar in typescript while leveraging type checking at compile time?

	interface GlobalSchema {
		name: string, 
		table: string, 
		attributes: {[key: string]: {
			type: string | number | boolean,
			required?: boolean,
			default?: string | number | boolean,
		}
	}; 

	const SampleSchema: GlobalSchema = {
		name: "Sample",
		table: "commonTable",
		attributes: {
			id: {
				type: string,
			},
			fname: {
				type: string,
				required: false,
			}, 
			lname: {
				type: string,
				default: "someName",
			}, 
			active: {
				type: boolean,
				default: true
			}
		}
	}

	const sampleEntity: SampleSchema.attributes = {
		id: "10", 
		fname: "John", 
		lname: "Doe", 
		active: true
	}
## [7][Extending enzyme to support wait](https://www.reddit.com/r/typescript/comments/ipzqwl/extending_enzyme_to_support_wait/)
- url: https://www.reddit.com/r/typescript/comments/ipzqwl/extending_enzyme_to_support_wait/
---
    import * as enzyme from 'enzyme';
    import Adapter from 'enzyme-adapter-react-16';
    import { act } from 'react-dom/test-utils';
    import wait from 'waait';
    
    enzyme.configure({ adapter: new Adapter() });
    
    declare module 'enzyme' {
      interface ReactWrapper {
        waitForUpdate: () =&gt; void;
      }
    }
    
    enzyme.ReactWrapper.prototype.waitForUpdate = async function waitForUpdate() {
      return act(async () =&gt; {
        await wait(0);
        this.setProps({});
      });
    };
    

and in test

    it('should', async() =&gt; {
    //do something
    await wrapper.waitForUpdate()
    //assert
    ))

any problem with this approach ?
## [8][Need a "little" help with some Typings](https://www.reddit.com/r/typescript/comments/ipm5sn/need_a_little_help_with_some_typings/)
- url: https://www.reddit.com/r/typescript/comments/ipm5sn/need_a_little_help_with_some_typings/
---
So I have some Typescript I want to use to create DOM-Elements easily but I just can't get the Type Inference to do the Correct thing I would be happy for Ideas what I could change.

[https://gist.github.com/mio991/c54e940e4ae599edec79488d30846db7](https://gist.github.com/mio991/c54e940e4ae599edec79488d30846db7)

The End usage should look like this:

    const templateFactory = div({})(
        div({})(
            h1({innerText: "title"})()
            h3({innerText: "subtitle"})()
        ),
        div({})(
            p({innerText: "article"})()
        )
    )
    
    // templateFactory should have the type: (context:{title:string; subtitle:string; article:string;})=&gt; HTMLDivElement
## [9][Cleaning up this mapped type](https://www.reddit.com/r/typescript/comments/ipqaj6/cleaning_up_this_mapped_type/)
- url: https://www.reddit.com/r/typescript/comments/ipqaj6/cleaning_up_this_mapped_type/
---
I got it working but it's not pretty. Any improvements welcome. I especially want to get rid of that empty string `''`

    interface myInt {
      "p1": string;
      "p2": string; 
    } 
    
    export default class X {
      private myMethod &lt;Record extends { [key in keyof myInt]: string }&gt;(matchPattern: RegExp, record: Record): boolean {
        const matchFound: boolean = matchPattern.test(record['' as keyof myInt]);
        return matchFound;
      }
## [10][Why does TS have a return type for void instead of using undefined?](https://www.reddit.com/r/typescript/comments/ipowhp/why_does_ts_have_a_return_type_for_void_instead/)
- url: https://www.reddit.com/r/typescript/comments/ipowhp/why_does_ts_have_a_return_type_for_void_instead/
---
Is this a carryover from Java, or some older language? Or is there a good reason to differentiate the default return `undefined` from `void`?
## [11][Angle syntax before and after the parameter definitions](https://www.reddit.com/r/typescript/comments/iposqh/angle_syntax_before_and_after_the_parameter/)
- url: https://www.reddit.com/r/typescript/comments/iposqh/angle_syntax_before_and_after_the_parameter/
---
In Typescript do the angle brackets before and after the parameter definitions repesent different things?

    async function a&lt;g extends {}&gt;(objs: Array&lt;g&gt;): Promise&lt;g&gt; {...}

In the first angle brackets I see a generic type argument declaration

In the 2nd and third angle brackets i see type constructors acting on a type argument `g`

I think these have about the same level of connection as the {} that represents an object literal and the {} that represents a function body. Basically none. Correct?
