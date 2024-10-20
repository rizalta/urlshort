// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Home() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>URL short</title><script src=\"/assets/js/htmx.min.js\"></script><link href=\"/assets/css/output.css\" type=\"text/css\" rel=\"stylesheet\"></head><body class=\"bg-black\"><div class=\"flex flex-col items-center justify-center gap-3\"><nav class=\"flex items-baseline justify-center w-full bg-black h-20 gap-2 pt-5 border-b border-b-white\"><a href=\"/\"><h1 class=\"text-3xl text-white font-mono\">URL short</h1></a></nav><div class=\"flex flex-col justify-center items-center w-full gap-10 pt-32\"><form hx-post=\"/shorten\" hx-swap=\"outerHTML\" class=\"flex flex-col justify-center items-center gap-6 w-full\"><p class=\"text-white text-2xl font-mono\">Shorten your urls</p><input type=\"text\" name=\"url\" class=\"border-none rounded-md shadow-md w-1/3 h-16 text-blue-800 text-2xl py-1 px-2 placeholder:text-blue-300\" placeholder=\"Enter url\"> <button type=\"submit\" class=\"px-4 py-2 w-36 h-16 text-2xl font-mono text-black bg-white rounded-md hover:bg-slate-100 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-slate-400\">Shorten</button></form></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate