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
## [2][Do I need both the @types modules, and the normal module?](https://www.reddit.com/r/typescript/comments/fkmpw3/do_i_need_both_the_types_modules_and_the_normal/)
- url: https://www.reddit.com/r/typescript/comments/fkmpw3/do_i_need_both_the_types_modules_and_the_normal/
---
Hey, I know this may seem like a silly question, but I'm brand new to TS, and a bit confused about this.

I have it so that my typescript compiles into javascript in a dist folder, where a copy of the node_modules folder is placed.

The server then runs out of dist, and I noticed that, despite me using @types modules exclusively, the javascript compiled errors when I do not have the non @types module installed.

For example, I have a module @types/winston installed, so I uninstalled the normal winston module, and the entire project ceased to function. Typescript of course gives no errors about this, as it's the javascript it compiled causing the issues.

Is this normal behavior, or behavior I could fix?
## [3][Curveball — March updates](https://www.reddit.com/r/typescript/comments/fkmcuk/curveball_march_updates/)
- url: https://medium.com//curveball-march-updates-10e0ea96f808?source=friends_link&amp;sk=10a5292d65d01264385fcd2eaa89e0d2
---

## [4][Curveball - Starter template, CLI colors and Content-Negotiation](https://www.reddit.com/r/typescript/comments/fkim3o/curveball_starter_template_cli_colors_and/)
- url: https://evertpot.com/curveball-march-updates/
---

## [5][Bring some structure to your file and directory names with ls-lint](https://www.reddit.com/r/typescript/comments/fk4rc3/bring_some_structure_to_your_file_and_directory/)
- url: https://www.reddit.com/r/typescript/comments/fk4rc3/bring_some_structure_to_your_file_and_directory/
---
Hey, worked on this a couple of weeks:

GitHub: [ls-lint](https://github.com/loeffel-io/ls-lint)

From time to time its hard to keep a clean naming structure in your projects. [ls-lint](https://github.com/loeffel-io/ls-lint) would solves this issue easily for all of your files and directories. With a simple .ls-lint.yml file, ls-lint makes sure that all your files are at the right place.

**Benefits:**

\- Works for directory and file names (all extensions supported)

\- Incredibly fast

\- Full unicode support

\- Almost zero third-party dependencies (only go-yaml)

**Example &amp; How-to (vue.js)**

    # .ls-lint.yml
    
    ls:
      .dir: regex:[a-z0-9\-]+
      .js: kebab-case
      .css: kebab-case
      .html: kebab-case
      .json: kebab-case
      .ts: kebab-case
      .sh: kebab-case
      .dev.js: kebab-case
      .prod.js: kebab-case
      .d.ts: kebab-case
      .vdom.js: kebab-case
      .spec.js: kebab-case
    
      dist:
        .js: point.case
    
      benchmarks/ssr:
        .js: camelCase
    
    ignore:
      - test
      - benchmarks/dbmon/ENV.js
      - .babelrc.js
      - .eslintrc.js
      - .github
      - .circleci
      - .git

Would love to get some feedback :-)

Pull requests are welcome!
## [6][Trying to figure out how to define multiple generic types, with an unknown number of them](https://www.reddit.com/r/typescript/comments/fkdtgz/trying_to_figure_out_how_to_define_multiple/)
- url: https://www.reddit.com/r/typescript/comments/fkdtgz/trying_to_figure_out_how_to_define_multiple/
---
Ok, so the code will make this more obvious. Here is what I have now:

    interface CreatorArgs&lt;Ctx&gt; {
    	context: {
    		type: Context&lt;Ctx&gt;;
    		value: Ctx;
    	};
    }

This interface has a "context" property, whose properties depend on the generic parameter Ctx. However, I have a situation where I need to be able to support a variable number of Contexts. So I'm trying to adjust this interface accordingly. Here is what I have so far:

    interface CreatorArgs&lt;Ctx&gt; {
    	contexts: {
    		[ctxKey: string]: {
    			type: Context&lt;Ctx&gt;;
    			value: Ctx;
    		};
    	};
    }

So each context will have a key which will contain the same structure from above. There will be an unknown number of these Contexts. My problem is the generic Ctx parameter, it will not be consistent in its value across all Contexts. 

So I'm wondering what the option is for how to type this. I'm not super experienced with TypeScript, I'm a longtime JS developer who is trying to get a better feel for TS. Any guidance, including telling me I'm approaching this all wrong, is appreciated.
## [7][Converting TypeScript in Javascript](https://www.reddit.com/r/typescript/comments/fkb1su/converting_typescript_in_javascript/)
- url: https://www.reddit.com/r/typescript/comments/fkb1su/converting_typescript_in_javascript/
---
Hi everyone,

I'm currently working on this project [https://github.com/matteobruni/tsparticles](https://github.com/matteobruni/tsparticles) and you can see the webpack config in the root.

Webpack is giving me warnings about the javascript output because it's becoming big. 

I'm thinking to split the options classes from the main file but I've read someone saying that webpack is not a great choice with TypeScript, so I was wondering about what's the best solution for having a javascript output.

ES5 is not required, if the solution requires ES6 it's still welcome.

If you have tips for the code they're welcome too, I've started some months ago with TypeScript
## [8][How do I assign a type to a property that is supposed to accept any React component](https://www.reddit.com/r/typescript/comments/fjufi8/how_do_i_assign_a_type_to_a_property_that_is/)
- url: https://www.reddit.com/r/typescript/comments/fjufi8/how_do_i_assign_a_type_to_a_property_that_is/
---
    type AnyComponent&lt;Props,State&gt; = FC&lt;Props&gt; | Component&lt;Props,State&gt;;
    
    interface CreatorArgs {
    	component: AnyComponent&lt;object, object&gt;; // I'll get more specific here once this works
    }
    
    // Much later in the code, creatorArgs: CreatorArgs
    
    const Component: AnyComponent&lt;object, object&gt; = creatorArgs.component;
    
    const output = () =&gt; (
    	&lt;Component /&gt; // JSX element type 'Component' does not have any construct or call signatures.
    );

So the code sample is above. Basically I'm writing a function, one of the arguments to this function is a React component. This should be able to be either a class component or a functional component. I tried creating a union type for them, but I'm getting the error in the comment above, where it doesn't have any constructor or call signature.

I'm an experienced JS dev trying to boost my TS skills here. Looking for guidance. Thanks.
## [9][What's up with language servers?](https://www.reddit.com/r/typescript/comments/fk6zte/whats_up_with_language_servers/)
- url: https://www.reddit.com/r/typescript/comments/fk6zte/whats_up_with_language_servers/
---
While the adoption of TypeScript and also of LSP grows, somehow TypeScript language servers seem to have been dying a slow death. [SourceGraph's server](https://github.com/sourcegraph/javascript-typescript-langserver/) has not seen real updates in at least a year and while everybody has migrated off to [TypeFox' one](https://github.com/theia-ide/typescript-language-server), it has not been updated for several months either and is actively (but apparently not very successfully) [looking for maintainers](https://github.com/theia-ide/typescript-language-server/issues/141).

So what's going on, has everyone migrated to VSCode and not using language servers anymore? ~~Is [Microsoft's LS](https://github.com/microsoft/vscode-languageserver-node) any good for using without VSCode?~~ (EDIT: It's a framework, not an implementation of TS language server)
## [10][[HELP] Bundle CSS / SCSS modules library with rollup](https://www.reddit.com/r/typescript/comments/fk1td1/help_bundle_css_scss_modules_library_with_rollup/)
- url: https://www.reddit.com/r/typescript/comments/fk1td1/help_bundle_css_scss_modules_library_with_rollup/
---
As the title says,

I've spent the last week poking around and trying different starter, but I am still unable to compile the library correctly. It is driving me crazy and I don't understand why I can't compile it.

Error:
```
[!] (plugin Typescript) TS2307: Cannot find module './index.module.css'.
src/index.tsx (4:25)

import * as styles from './index.module.css';
```

tsconfig.json
```json
{
	"compilerOptions": {
		"outDir": "dist",
		"module": "esnext",
		"target": "es5",
		"lib": ["es6", "dom", "es2016", "es2017"],
		"sourceMap": true,
		"allowJs": false,
		"jsx": "react",
		"declaration": true,
		"moduleResolution": "node",
		"allowSyntheticDefaultImports": true,
		"esModuleInterop": true,
		"plugins": [{ "name": "typescript-plugin-css-modules" }]
	},
	"include": ["src/**/*"],
	"exclude": ["node_modules", "dist"]
}

```


rollup.config.js

```ts
import commonjs from 'rollup-plugin-commonjs';
import del from 'rollup-plugin-delete';
import external from 'rollup-plugin-peer-deps-external';
import postcss from 'rollup-plugin-postcss';
import resolve from 'rollup-plugin-node-resolve';
import ts from '@wessberg/rollup-plugin-ts';

import pkg from './package.json';

export default {
	input: 'src/index.tsx',
	output: [
		{
			file: pkg.module,
			format: 'es',
			exports: 'named',
			sourcemap: true
		}
	],
	plugins: [
		del({ targets: 'dist/*' }),

		external(),
		postcss({
			extract: false,
			writeDefinitions: true,
			modules: true,
			namedExports: true
		}),

		ts({
			transpiler: 'babel',
			exclude: ['**/__tests__/**']
		}),

		resolve(),
		commonjs({
			exclude: ['**/*.stories.js']
		})
	]
};

```

index.tsx
```ts
// some code and then

import * as styles from './surface.module.css';
// import styles from './surface.module.css'; // also this doesn't work

// exports ...
```
## [11][Trouble in finding modules with '@' paths, getting "cannot find module '@/common/utils'" despite the file definitely existing and being recognised by my IDE.](https://www.reddit.com/r/typescript/comments/fjiqzr/trouble_in_finding_modules_with_paths_getting/)
- url: https://www.reddit.com/r/typescript/comments/fjiqzr/trouble_in_finding_modules_with_paths_getting/
---
I need help, `tsc` and `ts-node` do not recognise the custom path aliases specified in my `tsconfig.json` file. My file is as follows:

    {
      "compilerOptions": {
        "target": "es5",
        "module": "commonjs",
        "strict": true,
        "jsx": "preserve",
        "importHelpers": true,
        "moduleResolution": "node",
        "experimentalDecorators": true,
        "esModuleInterop": true,
        "allowSyntheticDefaultImports": true,
        "sourceMap": true,
        "baseUrl": ".",
        "types": ["webpack-env", "jest", "node"],
        "paths": {
          "@/*": ["src/*"]
        },
        "lib": ["esnext", "dom", "dom.iterable", "scripthost"]
      },
      "include": [
        "src/**/*.ts",
        "src/**/*.tsx",
        "src/**/*.vue",
        "tests/**/*.ts",
        "tests/**/*.tsx"
      ],
      "exclude": ["node_modules"]
    }

My IDE recognises the paths and I'm able to use the imported values as expected but when I compile any TS file or use `ts-node` on them, I get the error that the module is not recognised. Any help would be much appreciated, thanks fellas.

EDIT: **SOLVED** thanks to /u/moremattymattmatt, I had to install the `tsconfig-paths` npm package to use path aliasing.
