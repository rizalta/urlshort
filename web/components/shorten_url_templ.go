// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func ShortenURL(url string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"flex flex-col justify-center items-center gap-6 w-full\"><input type=\"text\" name=\"url\" id=\"shorturl\" class=\"border-none rounded-md shadow-md w-1/3 h-16 text-blue-800 text-2xl py-1 px-2\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(url)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/components/shorten_url.templ`, Line: 10, Col: 14}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" disabled><div class=\"flex flex-row gap-3\"><button type=\"button\" class=\"px-4 py-2 w-36 h-16 text-2xl font-mono text-black bg-white rounded-md hover:bg-slate-100 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-slate-400\" onclick=\"copyLink()\" id=\"btn\">Copy</button> <button class=\"px-4 py-2 w-36 h-16 text-2xl font-mono text-black bg-white rounded-md hover:bg-slate-100 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-slate-400\" type=\"button\" hx-get=\"/\" hx-target=\"body\">New</button></div></form><script>\n\tconst copyLink = () => {\n\t\tconst link = document.getElementById(\"shorturl\")\n\t\tconst button = document.getElementById(\"btn\")\n\t\tlink.select()\n\t\tlink.setSelectionRange(0, 99999)\n\n\t\tnavigator.clipboard.writeText(link.value)\n\t\t\t.then(() => {\n\t\t\t\tbutton.textContent = \"Copied\";\n\t\t\t\tsetTimeout(() => {\n\t\t\t\t\tbutton.textContent = \"Copy\";\n\t\t\t\t}, 2000);\n\t\t\t})\n\t\t\t.catch(err => {\n\t\t\t\tconsole.error('Failed to copy: ', err);\n\t\t\t\tbuttonElement.textContent = \"Error\";\n\t\t\t})\n\t}\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
