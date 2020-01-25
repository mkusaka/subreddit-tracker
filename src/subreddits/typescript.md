# typescript
## [1][Who's hiring Typescript developers - January](https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/)
- url: https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][TypeScript: X doesn't not exist on type Y](https://www.reddit.com/r/typescript/comments/etmblc/typescript_x_doesnt_not_exist_on_type_y/)
- url: https://www.reddit.com/r/typescript/comments/etmblc/typescript_x_doesnt_not_exist_on_type_y/
---
 

I Tried to use nodemon for auto restart server on changes but  Typescript throws the following error even though extended the type.  Please find the code attached below.

Typings work fine with VS Code but the errors while compiling by nodemon. Error attached below code. 

index.ts

    import * as fastify from 'fastify'; 
    import { Server, IncomingMessage, ServerResponse } from "http"; 
    const server: FastifyInstance&lt;Server, IncomingMessage, ServerResponse&gt; = fastify(); server.ready(err =&gt; { if (err) throw err;     server.swagger(); }); 
    const start = async () =&gt; { 
        try { 
            await server.listen(3000, "0.0.0.0");         
            server.log.info("Server running on 3000");         
            console.log("Server running on 3000"); 
        } catch (error) {         
            server.log.error(error); 
        } 
    }  
    start();

src/typings/fastify.d.ts

    import * as fastify from 'fastify'; 
    import * as http from "http";  
    declare module 'fastify' { 
        export interface FastifyInstance&lt; HttpServer = http.Server, HttpRequest = http.IncomingMessage, HttpResponse = http.ServerResponse &gt; {         
        swagger(): void 
        } 
    }

tsconfig.json

    { 
        "compileOnSave": false, 
        "compilerOptions": { 
            "baseUrl": ".", 
            "lib": [ "es2016" ], 
            "paths": { "*": [ "src/*" ] }, 
            "allowJs": true, 
            "target": "es2016", 
            "module": "commonjs", 
            "moduleResolution": "node", 
            "sourceMap": true, 
            "outDir": "./dist", 
            "typeRoots": [ "src/typings", "node_modules/@types" ] 
        }, 
        "include": [ "./src/**/*.ts" ], 
        "exclude": [ "node_modules" ], 
        "rules": {} 
    } 

[Error Thrown while compiling with NodeMon](https://i.stack.imgur.com/o71Xn.png)
## [3][Question: React onPress function type](https://www.reddit.com/r/typescript/comments/etdhan/question_react_onpress_function_type/)
- url: https://www.reddit.com/r/typescript/comments/etdhan/question_react_onpress_function_type/
---
Hey!

I just ported my react-native project to typescript and have a question about functions as props

im passing:

`&lt;DisplayCardsWithLikesdata={testData}likes={500}`**onPress={() =&gt; this.props.navigation.navigate("CardDetailScreen")}**`/&gt;`

&amp;#x200B;

to

`type Props = {onPress: Function}`

&amp;#x200B;

`const FloatingActionButtonSimple = (props:Props) =&gt; {const {onPress} = propsreturn (&lt;View style={styles.containerFab}&gt;&lt;TouchableOpacity style={styles.fab} onPress={onPress}&gt;&lt;Icon name="plus" size={16} color={"white"} /&gt;&lt;/TouchableOpacity&gt;&lt;/View&gt;);};`

&amp;#x200B;

So nothing crazy, but if i define onPress as a function it doesn't work. Does anybody of you have an idea on how to define onPress?

Thanks a lot!
## [4][Proper syntax for having a numeric variable in a string](https://www.reddit.com/r/typescript/comments/etb08f/proper_syntax_for_having_a_numeric_variable_in_a/)
- url: https://www.reddit.com/r/typescript/comments/etb08f/proper_syntax_for_having_a_numeric_variable_in_a/
---
So I have variables declared and everything, and I want to add it to a printed string. I want to have something like "Your X plants have been watered", and X is a variable rather than a string. I know I would have print("Your X plants have been watered") if I wanted it to say X, but how do I make it so its a numeric variable?
## [5][Set Up a Typescript React Redux Project](https://www.reddit.com/r/typescript/comments/et3gev/set_up_a_typescript_react_redux_project/)
- url: https://typeofnan.dev/setup-a-typescript-react-redux-project/
---

## [6][Is schemats, still maintained](https://www.reddit.com/r/typescript/comments/et56d4/is_schemats_still_maintained/)
- url: https://www.reddit.com/r/typescript/comments/et56d4/is_schemats_still_maintained/
---
On the GitHub the last commit was almost 2 years ago, does it work well?
Also does it work fine with both pg and pg-promise?

https://github.com/SweetIQ/schemats
## [7][I finally understand how to infer types!](https://www.reddit.com/r/typescript/comments/et0irh/i_finally_understand_how_to_infer_types/)
- url: https://www.reddit.com/r/typescript/comments/et0irh/i_finally_understand_how_to_infer_types/
---
&amp;#x200B;

https://preview.redd.it/6xjpi8f6slc41.png?width=913&amp;format=png&amp;auto=webp&amp;s=65540e7469a6f7abe5ed0facb700e0ade505db77

Type inference has always stumped me, forcing me to use explicit types / generics. I finally figured it out!

&amp;#x200B;

[The dropdown includes values that are all fields in the \`data\` array passed to \`useSearchItems\`.](https://preview.redd.it/hoc218ubslc41.png?width=760&amp;format=png&amp;auto=webp&amp;s=ba41378ccf7732168eadb5b97e2294e9a71b57fe)

If there are any experts on how typescript's \`infer\` keyword works, I think it could make for a super useful Medium article.
## [8][Question about Loops/Output not printing](https://www.reddit.com/r/typescript/comments/eszchn/question_about_loopsoutput_not_printing/)
- url: https://www.reddit.com/r/typescript/comments/eszchn/question_about_loopsoutput_not_printing/
---
So I'm trying to have a counter that's like Day: n for how many times the loop runs. So it would be like starting at day 0 and every time it goes through the loop just add one day. I've declared and initialized all my variables I'm dealing with. The code for my loop is: 

while (isHarvestTime === false &amp;&amp; alive === true) {  
 print("Day: " + day);  
}

&amp;#x200B;

All I'm getting as output is a blank screen so I'm not really sure how to get the day counter to print.
## [9][Drash - A microframework for deno](https://www.reddit.com/r/typescript/comments/esnv1t/drash_a_microframework_for_deno/)
- url: https://www.reddit.com/r/typescript/comments/esnv1t/drash_a_microframework_for_deno/
---
Would love to get feedback on this:

[https://www.reddit.com/r/Deno/comments/esazl4/drash\_a\_microframework\_for\_deno/](https://www.reddit.com/r/Deno/comments/esazl4/drash_a_microframework_for_deno/)

Thanks!
## [10][Unexpected Lessons from 100% Test Coverage](https://www.reddit.com/r/typescript/comments/eso0pv/unexpected_lessons_from_100_test_coverage/)
- url: https://medium.com/@EyasSH/unexpected-lessons-from-100-test-coverage-eebeee211b7a
---

## [11][Issues Converting .NET Cryptography to Typescript (3DES)](https://www.reddit.com/r/typescript/comments/esqigt/issues_converting_net_cryptography_to_typescript/)
- url: https://www.reddit.com/r/typescript/comments/esqigt/issues_converting_net_cryptography_to_typescript/
---
Hi r/typescript

I'm really stuck with converting a code to integrate with a .NET legacy code. I've tried a lot of libraries, such as crypto-js, crypto-es, but could not make this work.

The original code that I need to integrate with is the following, and the tries I've made are in the bottom. If anyone have any idea of what mistakes I'm doing, could give me some insights?

    
    string passPhrase = "test0001";
    
    // get the md5 hash of passPhrase as byte array
    byte[] md5PassPhraseHash = null;
    MD5CryptoServiceProvider md5 = new MD5CryptoServiceProvider();
    md5PassPhraseHash = md5.ComputeHash(Encoding.UTF8.GetBytes(passPhrase));
    byte[] decodedSecret = Convert.FromBase64String(secretAsBase64);
    
    TripleDESCryptoServiceProvider tdes_dec = new TripleDESCryptoServiceProvider();
    tdes_dec.Key = md5PassPhraseHash;
    byte[] cIV_dec = new byte[8];
    Array.Copy(md5PassPhraseHash, cIV_dec, 8);
    tdes_dec.IV = cIV_dec;
    
    MemoryStream memstream_dec = new MemoryStream();
    CryptoStream cs_dec = new CryptoStream(memstream_dec, tdes_dec.CreateDecryptor(), CryptoStreamMode.Write);
    cs_dec.Write(decodedSecret, 0, decodedSecret.Length);
    cs_dec.FlushFinalBlock();
    cs_dec.Close();
    byte[] decryptedSecret = memstream_dec.ToArray();
    Int64 secret = Convert.ToInt64(Encoding.UTF8.GetString(decryptedSecret));
    // increment secret
    secret = secret + 1;
    byte[] encodedSecret = Encoding.UTF8.GetBytes(Convert.ToString(secret));
    
    TripleDESCryptoServiceProvider tdes_enc = new TripleDESCryptoServiceProvider();
    tdes_enc.Key = md5PassPhraseHash;
    byte[] cIV_enc = new byte[8];
    Array.Copy(md5PassPhraseHash, cIV_enc, 8);
    tdes_enc.IV = cIV_enc;
    
    MemoryStream memstream = new MemoryStream();
    CryptoStream cs = new CryptoStream(memstream, tdes_enc.CreateEncryptor(), CryptoStreamMode.Write);
    cs.Write(encodedSecret, 0, encodedSecret.Length);
    cs.FlushFinalBlock();
    cs.Close();
    byte[] encryptedSecret = memstream.ToArray();
    string secret_encrypted = Convert.ToBase64String(encryptedSecret);
    

I've tried a lot of libraries and the native "crypto" package, but could not make it work. Can someone help me? Either I get invalid key length or even iv length issues.

&amp;#x200B;

Here's one of the tries:

    static decrypt(passPhrase: string, secret: string): string {
        const md5PassPhraseHash = crypto
          .createHash('md5')
          .update(passPhrase)
          .digest();
    
        const decodedSecret = Buffer.from(secret, 'base64').toString();
        const iv = Buffer.from(md5PassPhraseHash.slice(0, 8));
    
        const decipher = crypto.createDecipheriv('des-ede3-cbc', md5PassPhraseHash, iv);
        decipher.update(decodedSecret, 'base64', 'utf8');
    
        const result = decipher.final('utf-8');
        console.log({ result });
    
        return result.toString();
    }
