# typescript
## [1][Who's hiring Typescript developers - March](https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/)
- url: https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/
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
## [2][Elm style JSON decoder for TypeScript, providing type safety compile- and run-time](https://www.reddit.com/r/typescript/comments/fptbvb/elm_style_json_decoder_for_typescript_providing/)
- url: https://github.com/venil7/json-decoder
---

## [3][How to loop through a specific field of all documents and match it against a value?](https://www.reddit.com/r/typescript/comments/fpw42t/how_to_loop_through_a_specific_field_of_all/)
- url: https://www.reddit.com/r/typescript/comments/fpw42t/how_to_loop_through_a_specific_field_of_all/
---
Im reading and saving an NFC tag id into my firebase as a hexstring. I want to check all the tag ids against the one I scan with my phone and return add the matching product to my cart.. any advice?

&amp;#x200B;

[firebase](https://preview.redd.it/zydt47wml7p41.png?width=974&amp;format=png&amp;auto=webp&amp;s=d42e71e7aaf8bd0c2bc885147dc814a44b6eea53)

&amp;#x200B;

[ts](https://preview.redd.it/mnbhynfol7p41.png?width=677&amp;format=png&amp;auto=webp&amp;s=8ac975ee0b9d60ce6527f8dd81b15d47d0ef3fc1)
## [4][Approximating haskell's do syntax in Typescript](https://www.reddit.com/r/typescript/comments/fpvsxu/approximating_haskells_do_syntax_in_typescript/)
- url: https://paulgray.net/do-syntax-in-typescript/
---

## [5][TypeScript + React: Typing custom hooks with tuple types](https://www.reddit.com/r/typescript/comments/fpt5nf/typescript_react_typing_custom_hooks_with_tuple/)
- url: https://fettblog.eu/typescript-react-typeing-custom-hooks/
---

## [6][Building Vue Enterprise Application: Part 2. Services](https://www.reddit.com/r/typescript/comments/fpjcut/building_vue_enterprise_application_part_2/)
- url: https://medium.com/@gregsolo/building-vue-enterprise-application-part-2-services-f7ec400190e7
---

## [7][How to convert this to ts](https://www.reddit.com/r/typescript/comments/fptqhj/how_to_convert_this_to_ts/)
- url: https://www.reddit.com/r/typescript/comments/fptqhj/how_to_convert_this_to_ts/
---
Hello guys, any suggestions on how to convert this to ts?
const mutations = {
 [AUTH_REQUEST]: (state) =&gt; {
   state.status = ""
}}
## [8][How do I turn Typescript types into a markdown table?](https://www.reddit.com/r/typescript/comments/fph6m0/how_do_i_turn_typescript_types_into_a_markdown/)
- url: https://www.reddit.com/r/typescript/comments/fph6m0/how_do_i_turn_typescript_types_into_a_markdown/
---
I'm writing docs for a library and would like to find a way to turn my \`props\` into a simple markdown table. Every example I've seen so far feels too complex, and creates a whole docs website for me.

Is there a way to turn this:

```typescript
type Props = {
  /*
   * this is a required color!
   */
  color: string
}
```

into this:

| Prop  | description  | type  | Required  |
|---|---|---|---|
|  color | this is a required color!  | string  | yes |
## [9][Am I using the right RxJS operator?](https://www.reddit.com/r/typescript/comments/fpco9r/am_i_using_the_right_rxjs_operator/)
- url: https://www.reddit.com/r/typescript/comments/fpco9r/am_i_using_the_right_rxjs_operator/
---
Hi guys!

I am struggling with understanding when to use the different operators. I have one method where I first subscribe to check if a service is online if it is then I subscribe to another observable. If not I show a dialogue saying service is down. What I have done is something like this:

    public methodName(parameter: parameterName): any {
            this.nameOfService
                .isServiceOnline()
                .pipe(
                    concatMap(isOnline =&gt; {
                        if (!isOnline) {
                            this.showServiceDownDialog();
                            return;
                        }
    
                        do stuff
    
                        return this.nameOfService.searchInService(parameter);
                    })
                )
                .subscribe(
                    data =&gt; {
                        do more stuff
                    }
                );
        }

Does this look ok? I used concatMap because I guess I want to make sure the service is online before I try to use the search method? Or is there any other operator I should use here?
## [10][Making a 3rd Party Interface more strict](https://www.reddit.com/r/typescript/comments/fpc17j/making_a_3rd_party_interface_more_strict/)
- url: https://www.reddit.com/r/typescript/comments/fpc17j/making_a_3rd_party_interface_more_strict/
---
Hi. I am importing a 3rd party library (xml-parser) that returns an instance of the following interface from a function

    export interface Attributes {
        [name: string]: string;
    }

In my code I'd like to be able to project into a stricter type because I know the attributes returned will have certain property names e.g.

    type UserInfo = {
      username: string
    }

However assigning something of interface Attributes to a const of type UserInfo, rightly, doesn't compile. 

Is there a utility type that I can use to say this Attribute interface can be utilised like UserInfo? I tried Pick&lt;Attributes, 'username'&gt; but again username doesn't exist on the Attribute type.
## [11][Enums, string/number types of interfaces/classes cannot be used as index signatures](https://www.reddit.com/r/typescript/comments/fp5wfk/enums_stringnumber_types_of_interfacesclasses/)
- url: https://www.reddit.com/r/typescript/comments/fp5wfk/enums_stringnumber_types_of_interfacesclasses/
---
I just came across this post [https://github.com/microsoft/TypeScript/issues/37448](https://github.com/microsoft/TypeScript/issues/37448)...

Does anyone have the problem here? Are there any strategies you'd found to overcome it?
