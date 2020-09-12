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
## [2][I made a ts template repo on github that has CI/CD actions configured, husky lint-staged, test running, eslint, prettier, and nodemon for people's usage in hacktoberfest.](https://www.reddit.com/r/typescript/comments/ir66mv/i_made_a_ts_template_repo_on_github_that_has_cicd/)
- url: https://github.com/jakehamtexas/ts-lib-starter
---

## [3][What are some of the best engineered simple CRUD application using Typescript, Redux, Axios and implementing localization?](https://www.reddit.com/r/typescript/comments/irc69u/what_are_some_of_the_best_engineered_simple_crud/)
- url: https://www.reddit.com/r/typescript/comments/irc69u/what_are_some_of_the_best_engineered_simple_crud/
---
Trying to find a simple example, and improve upon it. It's the best way to learn. Finding the application that's the most well-made and looking at the best practices.
## [4][Could I get some advice about an error I'm getting?](https://www.reddit.com/r/typescript/comments/irahqn/could_i_get_some_advice_about_an_error_im_getting/)
- url: https://www.reddit.com/r/typescript/comments/irahqn/could_i_get_some_advice_about_an_error_im_getting/
---
I'm writing a small number sequences game for my son to play - I've written a library in typescript with a buddy of mine - all it does is generate number sequences;

[https://github.com/mikeyhogarth/number-sequences/tree/v1.2.1](https://github.com/mikeyhogarth/number-sequences/tree/v1.2.1)

I then started creating a UI to use that library;

[https://github.com/mikeyhogarth/number-sequences-game](https://github.com/mikeyhogarth/number-sequences-game) (it's deployed if anyone wants to take a look, but don't expect much at the moment!).

When I first imported the \`number-sequences\` library, the compiler was throwing an error;

    Could not find a declaration file for module 'number-sequences'. '/Users/mikey/Development/number-sequences-game/node_modules/number-sequences/dist/index.js' implicitly has an 'any' type.
      Try `npm install @types/number-sequences` if it exists or add a new declaration (.d.ts) file containing `declare module 'number-sequences';`ts(7016)

I got around this by adding the following line to \`react-app.d.ts\`

**declare module "number-sequences";**

But that feels like something I shouldn't have had to do - or is it? Is there something I could have done in the original library to prevent having to do this?
## [5][What are some useful configuration that you can do to make Typescript development easier on Visual Studio Code?](https://www.reddit.com/r/typescript/comments/iqr21h/what_are_some_useful_configuration_that_you_can/)
- url: https://www.reddit.com/r/typescript/comments/iqr21h/what_are_some_useful_configuration_that_you_can/
---
What are some useful configuration that you can do to make Typescript development easier on Visual Studio Code? I am wondering if you can catch all errors before transpilation and before runtime. Some coworker told me that there are type errors appearing at runtime for some reason on his side.
## [6][Building a Game With TypeScript. Drawing Grid 5/5](https://www.reddit.com/r/typescript/comments/iqn5nn/building_a_game_with_typescript_drawing_grid_55/)
- url: https://medium.com/gregsolo/building-a-game-with-typescript-drawing-grid-5-5-49454917b3af?source=friends_link&amp;sk=84924dff6c15ff16f3a01f6d35fe48d1
---

## [7][Adding a property switches an object from one type to another. Should I override this interpreter warning?](https://www.reddit.com/r/typescript/comments/iqtik7/adding_a_property_switches_an_object_from_one/)
- url: https://www.reddit.com/r/typescript/comments/iqtik7/adding_a_property_switches_an_object_from_one/
---
    campaignCallData.map((record: IJoinedCampaignCall) =&gt; {
          const weekDate: Moment = this.extractWeek(record.date);
          const stringWeekDate: string = weekDate.format('YYYY-MM-DD');
          (record as IFormattedCampaignCall).weekDate = stringWeekDate;
          return record;
        });
    
    /*
    Conversion of type 'IJoinedCampaignCall' to type 'IFormattedCampaignCall' may be a mistake because neither type sufficiently overlaps with the other. If this was intentional, convert the expression to 'unknown' first.
      Property '"weekDate"' is missing in type 'IJoinedCampaignCall' but required in type 'IFormattedCampaignCall'.ts(2352)
    IFormattedCampaignCall.ts(11, 3): '"weekDate"' is declared here.
    */

If I follow the lint suggestion, this will work:

     (record as unknown as IFormattedCampaignCall).weekDate = stringWeekDate;
    

My question is whether this is an appropriate place to override the interpreter using `unknown`? 

The "no overlapping type" explanation is confusing since I'm obviously asseting from one type to a different one. 

I lean towards yes, seems logical if  `IFormattedCampaignCall` has the additional property, and everything else is the same, it should be valid.
## [8][Enums vs Object.freeze](https://www.reddit.com/r/typescript/comments/iqc2av/enums_vs_objectfreeze/)
- url: https://www.reddit.com/r/typescript/comments/iqc2av/enums_vs_objectfreeze/
---
I have read sentiment on here that Enums may have been a directional misfit. 

And because I tend to use string Enums to avoid string arguments, I think I may do what one poster said and revert to `Object.freeze` again. 

They are essentially the same right? Or is there a strong case for one over the other?
## [9][Any React libraries you would recommend avoiding when working with Typescript?](https://www.reddit.com/r/typescript/comments/iqfq41/any_react_libraries_you_would_recommend_avoiding/)
- url: https://www.reddit.com/r/typescript/comments/iqfq41/any_react_libraries_you_would_recommend_avoiding/
---
I had an issue with some library not long ago, and I was wondering if there were any other libraries that may become a pain in the ass because it's not well integrated with Typescript, or it's poorly designed and doesn't really use Typescript's features very well at all.
## [10][Avoid Export Default](https://www.reddit.com/r/typescript/comments/ipou7q/avoid_export_default/)
- url: https://basarat.gitbook.io/typescript/main-1/defaultisbad
---

## [11][Writing a schema parser in typescript](https://www.reddit.com/r/typescript/comments/iq003l/writing_a_schema_parser_in_typescript/)
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
