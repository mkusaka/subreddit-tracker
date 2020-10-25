# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][imgcat - a tool to output images as RGB ANSI graphics on the terminal](https://www.reddit.com/r/golang/comments/jhms24/imgcat_a_tool_to_output_images_as_rgb_ansi/)
- url: https://github.com/trashhalo/imgcat
---

## [3][Lightweight UI library recommendation](https://www.reddit.com/r/golang/comments/jhr4ft/lightweight_ui_library_recommendation/)
- url: https://www.reddit.com/r/golang/comments/jhr4ft/lightweight_ui_library_recommendation/
---
I am building very simple UI app to run on Raspberry Pi Zero.  The UI showcases static infographics operable(zoom, scroll) through IR remote.

Which Golang UI library will you recommend?

* [Fyne](https://fyne.io/) 11k
* [Qt](https://therecipe.github.io/qt/) 7.9k
* [webview](https://github.com/webview/webview) 7k
* [Gio](https://gioui.org/) 
* [Lorca](https://github.com/zserge/lorca) 5.9k
* [go-astilectron](https://github.com/asticode/go-astilectron) 3.5k
* [wails](https://github.com/wailsapp/wails) 2.6k
* [gotk3](https://github.com/gotk3/gotk3) 1.3k
## [4][Dagger - A zero dependency, concurrency safe, in-memory directed graph database](https://www.reddit.com/r/golang/comments/jhk1ti/dagger_a_zero_dependency_concurrency_safe/)
- url: https://www.reddit.com/r/golang/comments/jhk1ti/dagger_a_zero_dependency_concurrency_safe/
---
Godoc: [https://godoc.org/github.com/autom8ter/dagger](https://godoc.org/github.com/autom8ter/dagger)

Github: [https://github.com/autom8ter/dagger](https://github.com/autom8ter/dagger)

    import "github.com/autom8ter/dagger"

&amp;#x200B;

[Directed Graph](https://preview.redd.it/4fle7kx835v51.png?width=128&amp;format=png&amp;auto=webp&amp;s=cc83217afbfab3f805c259f4c4a6dc2d8ef1cd0b)

[What is a directed Graph?](https://en.wikipedia.org/wiki/Directed_graph)

## Design:

* flexibility
* global state
   * see [primitive](https://godoc.org/github.com/autom8ter/dagger/primitive) to manage graph state manually
* concurrency safe
* high performance
* simple api

## Features

* \[x\] native graph objects(nodes/edges)
* \[x\] typed graph objects(ex: user/pet)
* \[x\] labelled nodes &amp; edges
* \[x\] depth first search
* \[x\] breadth first search
* \[x\] concurrency safe
* \[ \] import graph from JSON blob
* \[ \] export graph to JSON blob

## Example

       coleman = dagger.NewNode("user", fmt.Sprintf("cword_%v", time.Now().UnixNano()), map[string]interface{}{
       		"name": "coleman",
       	})
       	tyler = dagger.NewNode("user", fmt.Sprintf("twash_%v", time.Now().UnixNano()), map[string]interface{}{
       		"name": "tyler",
       	})
       	sarah = dagger.NewNode("user", fmt.Sprintf("swash_%v", time.Now().UnixNano()), map[string]interface{}{
       		"name": "sarah",
       	})
       	lacee = dagger.NewNode("user", fmt.Sprintf("ljans_%v", time.Now().UnixNano()), map[string]interface{}{
       		"name": "lacee",
       	})
       	// random id will be generated if one isn't provided
       	charlie = dagger.NewNode("dog", "", map[string]interface{}{
       		"name":   "charlie",
       		"weight": 25,
       	})
       
       	if err := coleman.Connect(tyler, "friend", true); err != nil {
       		exitErr(err)
       	}
       	if err := sarah.Connect(lacee, "friend", true); err != nil {
       		exitErr(err)
       	}
       	if err := coleman.Connect(lacee, "fiance", true); err != nil {
       		exitErr(err)
       	}
       	if err := tyler.Connect(sarah, "wife", true); err != nil {
       		exitErr(err)
       	}
       	if err := coleman.Connect(charlie, "pet", false); err != nil {
       		exitErr(err)
       	}
       	if err := lacee.Connect(charlie, "pet", false); err != nil {
       		exitErr(err)
       	}
       	if err := charlie.Connect(lacee, "owner", false); err != nil {
       		exitErr(err)
       	}
       	if err := charlie.Connect(coleman, "owner", false); err != nil {
       		exitErr(err)
       	}
       	charlie.Patch(map[string]interface{}{
       		"weight": 19,
       	})
       	if charlie.GetInt("weight") != 19 {
       		exit("expected charlie's weight to be 19!")
       	}
       	// check to make sure edge is patched
       	coleman.EdgesFrom(func(e *dagger.Edge) bool {
       		if e.Type() == "pet" &amp;&amp; e.GetString("name") == "charlie" {
       			if e.To().GetInt("weight") != 19 {
       				exit("failed to patch charlie's weight")
       			}
       		}
       		return true
       	})
       	if coleman.GetString("name") != "coleman" {
       		exit("expected name to be coleman")
       	}
       	// remove from graph
       	charlie.Remove()
       	// no longer in graph
       	if dagger.HasNode(charlie) {
       		exit("failed to delete node - (charlie)")
       	}
       	// check to make sure edge no longer exists(cascade)
       	coleman.EdgesFrom(func(e *dagger.Edge) bool {
       		if e.Type() == "pet" &amp;&amp; e.GetString("name") == "charlie" {
       			exit("failed to delete node - (charlie)")
       		}
       		return true
       	})
       	// check to make sure edge no longer exists(cascade)
       	lacee.EdgesFrom(func(e *dagger.Edge) bool {
       		if e.Type() == "pet" &amp;&amp; e.GetString("name") == "charlie" {
       			exit("failed to delete node - (charlie)")
       		}
       		return true
       	})
       	fmt.Printf("registered node types = %v", dagger.NodeTypes())
       	fmt.Printf("registered edge types = %v", dagger.EdgeTypes())
       	dagger.RangeNodes(func(n *dagger.Node) bool {
       		bits, err := n.JSON()
       		if err != nil {
       			exitErr(err)
       		}
       		fmt.Printf("\nfound node = %v\n", string(bits))
       		n.EdgesFrom(func(e *dagger.Edge) bool {
       			bits, err := e.JSON()
       			if err != nil {
       				exitErr(err)
       			}
       			fmt.Println(string(bits))
       			return true
       		})
       		n.EdgesTo(func(e *dagger.Edge) bool {
       			bits, err := e.JSON()
       			if err != nil {
       				exitErr(err)
       			}
       			fmt.Println(string(bits))
       			return true
       		})
       		return true
       	}
## [5][Go Proxy built with Cloudflare Workers](https://www.reddit.com/r/golang/comments/jhixs6/go_proxy_built_with_cloudflare_workers/)
- url: https://www.reddit.com/r/golang/comments/jhixs6/go_proxy_built_with_cloudflare_workers/
---
Would anyone be willing to help me test out a Go Proxy? I built it using Cloudflare Workers so it should be fast. The domain also uses Argo, which should help with cache hits. 

https://goproxy.dev/

https://workers.cloudflare.com/

https://www.cloudflare.com/network/

Thanks!
## [6][Discussing the browser compatibility of WASM produced by the official Go compiler](https://www.reddit.com/r/golang/comments/jhqltd/discussing_the_browser_compatibility_of_wasm/)
- url: https://www.reddit.com/r/golang/comments/jhqltd/discussing_the_browser_compatibility_of_wasm/
---
There currently is an open discussion on browser compatibility for Go-based WASM that I believe could use a wider audience: [https://github.com/golang/go/issues/28360](https://github.com/golang/go/issues/28360). It could have a big impact on the usefulness of Go for targeting the future web (such as supporting the Safari web browser, or not). So please feel free to write down your thoughts and ideas!
## [7][Wide Go Snippets Extension for VS Code 😋 based vim-go](https://www.reddit.com/r/golang/comments/jhl7ot/wide_go_snippets_extension_for_vs_code_based_vimgo/)
- url: https://marketplace.visualstudio.com/items?itemName=umutbasal.go-snippets-vscode
---

## [8][taskflow - Create your build pipeline in Go](https://www.reddit.com/r/golang/comments/jhsksy/taskflow_create_your_build_pipeline_in_go/)
- url: https://github.com/pellared/taskflow
---

## [9][Implementing a b-tree and being to clever](https://www.reddit.com/r/golang/comments/jhrtjn/implementing_a_btree_and_being_to_clever/)
- url: https://www.reddit.com/r/golang/comments/jhrtjn/implementing_a_btree_and_being_to_clever/
---
Hey,

I started working on a persistent key-value store named [keyver](https://github.com/j0holo/keyver). Just as a hobby project.

Currently I'm working on implementing a b-tree for indexing, but I have trouble implementing it. [As you can see in the code](https://github.com/j0holo/keyver/blob/master/cmd/server/b_tree.go) I tried to be clever (too clever?) by creating an interface which is implemented by nodes that contains the pointers and the nodes that hold the actual values.

How [Wikipedia](https://en.wikipedia.org/wiki/B-tree#Performance) describes the implementation is somewhat clear but at the same time I don't understand how I can "easily" translate this to code.

btw, I know about [https://github.com/google/btree](https://github.com/google/btree) but they use a lot of reflection, I only need to store strings for both the key and value.

Am I being too clever? How should I approach this?
## [10][What’s the best way to learn Go?](https://www.reddit.com/r/golang/comments/jhrska/whats_the_best_way_to_learn_go/)
- url: https://www.reddit.com/r/golang/comments/jhrska/whats_the_best_way_to_learn_go/
---
I’ve finished my final college semester and I’ve got some some spare time that I’m looking to invest in learning Go

What’s the best resoucres or projects you found useful when developing your go knowledge ? 

Anything you wish you could’ve told your yourself when you were learning ?
## [11][ion-sfu: simple, scalable webrtc sessions](https://www.reddit.com/r/golang/comments/jh9d0s/ionsfu_simple_scalable_webrtc_sessions/)
- url: https://github.com/pion/ion-sfu
---

