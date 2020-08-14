# golang
## [1][I also push directly to prod!](https://www.reddit.com/r/golang/comments/i91ynq/i_also_push_directly_to_prod/)
- url: https://i.redd.it/ql3zpki3gsg51.png
---

## [2][From Python to Go: migrating our entire API](https://www.reddit.com/r/golang/comments/i9fo8f/from_python_to_go_migrating_our_entire_api/)
- url: https://www.repustate.com/blog/migrating-entire-api-go-python/
---

## [3][We can stop adding debug log lines! we can now inject Golang tracepoints via eBPF](https://www.reddit.com/r/golang/comments/i9cs68/we_can_stop_adding_debug_log_lines_we_can_now/)
- url: https://www.reddit.com/r/golang/comments/i9cs68/we_can_stop_adding_debug_log_lines_we_can_now/
---
We've been hacking ways to get live  debug data using eBPF and have finally cracked it.

Here's a quick guide on how to inject tracepoints and dynamically log running Golang code: [https://docs.pixielabs.ai/tutorials/simple-go-tracing/](https://docs.pixielabs.ai/tutorials/simple-go-tracing/).

Would be great to get feedback from fellow gophers :)

&amp;#x200B;

https://preview.redd.it/a9jdmjy4fvg51.png?width=1710&amp;format=png&amp;auto=webp&amp;s=4893e3ec525e6a3e385d7c8aa7c823d372d28050

Here's a 2 min video walking through the simple tutorial: [https://www.youtube.com/watch?v=aH7PHSsiIPM&amp;feature=youtu.be](https://www.youtube.com/watch?v=aH7PHSsiIPM&amp;feature=youtu.be)
## [4][go report PDF file (mutilate language)](https://www.reddit.com/r/golang/comments/i9gr6t/go_report_pdf_file_mutilate_language/)
- url: https://www.reddit.com/r/golang/comments/i9gr6t/go_report_pdf_file_mutilate_language/
---
# gopdf Introduction

[gopdf](https://www.github.com/tiechui1994/gopdf) is a relatively complete **PDF** export library, which integrates third-party libraries, making it more convenient for users to develop and use. It has the following features:

* Support Unicode characters (including Chinese, Japanese, Korean, etc.)
* Automatic pagination of PDF documents
* Automatic wrapping of PDF documents
* Automatic global positioning of PDF documents, no need for users to manually locate
* PDF default configuration options are simple, and several commonly used methods have been built-in
* PDF documents adopt the attribute settings similar to html pages, which are easy to understand
* PDF supports image insertion, the format can be PNG or JPEG, the image will be properly compressed
* PDF supports document compression
* PDF conversion unit built-in processing
* Executor can be nested

github link: [https://github.com/tiechui1994/gopdf](https://github.com/tiechui1994/gopdf)

## Example:

    const (
        TABLE_IG = "IPAexG"
        TABLE_MD = "MPBOLD"
        TABLE_MY = "微软雅黑"
    )
    
    func ManyTableReportWithData() {
        r := core.CreateReport()
        font1 := core.FontMap{
        	FontName: TABLE_IG,
        	FileName: "example//ttf/ipaexg.ttf",
        }
        font2 := core.FontMap{
        	FontName: TABLE_MD,
        	FileName: "example//ttf/mplus-1p-bold.ttf",
        }
        font3 := core.FontMap{
        	FontName: TABLE_MY,
        	FileName: "example//ttf/microsoft.ttf",
        }
        r.SetFonts([]*core.FontMap{&amp;font1, &amp;font2, &amp;font3})
        r.SetPage("A4", "mm", "P")
        
        r.RegisterExecutor(core.Executor(ManyTableReportWithDataExecutor), core.Detail)
        
        r.Execute("many_table_data.pdf")
        r.SaveAtomicCellText("many_table_data.txt")
        }
        
        func ManyTableReportWithDataExecutor(report *core.Report) {
        unit := report.GetUnit()
        
        lineSpace := 0.01 * unit
        lineHeight := 2 * unit
        
        rows, cols := 800, 5
        table := NewTable(cols, rows, 80*unit, lineHeight, report)
        table.SetMargin(core.Scope{0, 0, 0, 0})
        
        for i := 0; i &lt; rows; i += 5 {
        	key := rand.Intn(3)
        	//key := (i+1)%2 + 1
        	f1 := core.Font{Family: TABLE_MY, Size: 10}
        	border := core.NewScope(0.5*unit, 0.5*unit, 0, 0)
        
        	switch key {
        	case 0:
        	   for row := 0; row &lt; 5; row++ {
                  for col := 0; col &lt; cols; col++ {
            	  conent := fmt.Sprintf("%v-(%v,%v)", 0, i+row, col)
            	  cell := table.NewCell()
            	  txt := NewTextCell(table.GetColWidth(i+row, col), lineHeight, lineSpace, report)
            	  txt.SetFont(f1).SetBorder(border).SetContent(conent + GetRandStr(1))
            	  cell.SetElement(txt)    
            	}
        	}
        
        	case 1:
    	  c00 := table.NewCellByRange(1, 5) 
    	  c01 := table.NewCellByRange(2, 2) 
    	  c03 := table.NewCellByRange(2, 3)
    	  c21 := table.NewCellByRange(2, 1)
    	  c31 := table.NewCellByRange(4, 1)
    	  c41 := table.NewCellByRange(4, 1)
        
        	  t00 := NewTextCell(table.GetColWidth(i+0, 0), lineHeight, lineSpace, report)
        	  t01 := NewTextCell(table.GetColWidth(i+0, 1), lineHeight, lineSpace, report)
        	  t03 := NewTextCell(table.GetColWidth(i+0, 3), lineHeight, lineSpace, report)
        	  t21 := NewTextCell(table.GetColWidth(i+2, 1), lineHeight, lineSpace, report)
        	  t31 := NewTextCell(table.GetColWidth(i+3, 1), lineHeight, lineSpace, report)
        	  t41 := NewTextCell(table.GetColWidth(i+4, 1), lineHeight, lineSpace, report)
        
        	  t00.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 1, i+0, 0) + GetRandStr(5))
        	  t01.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 1, i+0, 1) + GetRandStr(4))
        	  t03.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 1, i+0, 3) + GetRandStr(6))
        	  t21.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 1, i+2, 1) + GetRandStr(2))
        	  t31.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 1, i+3, 1) + GetRandStr(4))
        	  t41.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 1, i+4, 1) + GetRandStr(4))
        
        	  c00.SetElement(t00)
        	  c01.SetElement(t01)
        	  c03.SetElement(t03)
        	  c21.SetElement(t21)
        	  c31.SetElement(t31)
        	  c41.SetElement(t41)
        
        	case 2:
        	  c00 := table.NewCellByRange(3, 2)
    	  c03 := table.NewCellByRange(2, 3)
    	  c20 := table.NewCellByRange(1, 2)
    	  c21 := table.NewCellByRange(2, 3)
    	  c33 := table.NewCellByRange(2, 2)
    	  c40 := table.NewCellByRange(1, 1)
        
    	  t00 := NewTextCell(table.GetColWidth(i+0, 0), lineHeight, lineSpace, report)
    	  t03 := NewTextCell(table.GetColWidth(i+0, 3), lineHeight, lineSpace, report)
    	  t20 := NewTextCell(table.GetColWidth(i+2, 0), lineHeight, lineSpace, report)
    	  t21 := NewTextCell(table.GetColWidth(i+2, 1), lineHeight, lineSpace, report)
    	  t33 := NewTextCell(table.GetColWidth(i+3, 3), lineHeight, lineSpace, report)
    	  t40 := NewTextCell(table.GetColWidth(i+4, 0), lineHeight, lineSpace, report)
        
    	  t00.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 2, i+0, 0) + GetRandStr(6))
    	  t03.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 2, i+0, 3) + GetRandStr(6))
    	  t20.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 2, i+2, 0) + GetRandStr(2))
    	  t21.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 2, i+2, 1) + GetRandStr(6))
    	  t33.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 2, i+3, 3) + GetRandStr(4))
    	  t40.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 2, i+4, 0) + GetRandStr(1))
        
    	  c00.SetElement(t00)
    	  c03.SetElement(t03)
    	  c20.SetElement(t20)
    	  c21.SetElement(t21)
    	  c33.SetElement(t33)
    	  c40.SetElement(t40)
        	}    
        }
        
        table.GenerateAtomicCell()
    }
    
    
    func GetRandStr(l ...int) string {
        seed := rand.New(rand.NewSource(time.Now().UnixNano()))
        str := "0123456789ABCDEFGHIGKLMNOPQRSTUVWXYZ"
        l = append(l, 8)
        r := seed.Intn(l[0]*11) + 8
        data := strings.Repeat(str, r/36+1)
        return data[:r] + "---"
    }
    
    func TestTable(t *testing.T) {
        ManyTableReportWithData()
    }

The above case shows a complex and irregular case, the rendering is as follows:

https://preview.redd.it/jem73ljozwg51.png?width=1157&amp;format=png&amp;auto=webp&amp;s=bc531d44594e59176bed41f8a7108afd479896fb

## Future plan

1. Planning to develop the part of `Markdown` to PDF
2. Optimize components

To update the detailed usage, please refer to the open source library [https://github.com/tiechui1994/gopdf](https://github.com/tiechui1994/gopdf), everyone is welcome to contribute your own strength.
## [5][Go 1.15's interface optimization for small integers is invisible to Go programs](https://www.reddit.com/r/golang/comments/i9ll67/go_115s_interface_optimization_for_small_integers/)
- url: https://utcc.utoronto.ca/~cks/space/blog/programming/Go115InterfaceSmallIntsII
---

## [6][Proposal: Register-based Go calling convention](https://www.reddit.com/r/golang/comments/i8x4xe/proposal_registerbased_go_calling_convention/)
- url: https://go.googlesource.com/proposal/+/refs/changes/78/248178/1/design/40724-register-calling.md
---

## [7][Implementing Language Server Protocol in go - for other languages. Advice please](https://www.reddit.com/r/golang/comments/i9ixbr/implementing_language_server_protocol_in_go_for/)
- url: https://www.reddit.com/r/golang/comments/i9ixbr/implementing_language_server_protocol_in_go_for/
---
I have looked at the source graph and gopls code and and am looking for a solution to implement LSP for “language-x” without doing the JSON rpc and request response mapping from scratch.
Can anyone who has worked in this space suggest how to approach this, or what parts of existing projects are general purpose enough to reuse?
## [8][When an Interface Depends on Another Interface in Go](https://www.reddit.com/r/golang/comments/i9iqw7/when_an_interface_depends_on_another_interface_in/)
- url: https://medium.com/@orenrosenblum/when-an-interface-depends-on-another-interface-in-go-a32d988cd21e?sk=6755fe1ad1a55e47237dcc2867935800
---

## [9][GraphQL and Authorization](https://www.reddit.com/r/golang/comments/i9drqp/graphql_and_authorization/)
- url: https://www.reddit.com/r/golang/comments/i9drqp/graphql_and_authorization/
---
I am using gqlgen and I was wondering if any of you guys have experience implementing authorization (not authentication) with gqlgen or any other Go based graphql library? If so, which layer did you put that logic in? I was thinking of using Casbin to help with the authorization process, but I don't know if I should put my logic in a middleware, context, or elsewhere. Any examples would be greatly appreciated.
## [10][Can anyone help with golang.org/x/crypto/ssh?](https://www.reddit.com/r/golang/comments/i9i5i5/can_anyone_help_with_golangorgxcryptossh/)
- url: https://www.reddit.com/r/golang/comments/i9i5i5/can_anyone_help_with_golangorgxcryptossh/
---
Hey guys, it's not often that I ask questions here, but this time, I'm lost.

So, I have to make an app that (in the background) connects to an SSH session, starts a shell, runs some commands, gets the whole output back and somehow deals with it.

The problem is, the server to which I am connecting can be rather slow in responding and simply cuts off the output if I'm going ahead and entering new commands without waiting for the output of the previous ones.

Is there some way to wait until the server sends me the full response for a command in shell and only then continuing with another command? Any insights on this would be very appreaciated!
