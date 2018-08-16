package assets

import "github.com/moisespsena/go-assetfs/api"

func TemplateFS(fs api.Interface) api.Interface {
	return fs.NameSpace("templates")
}

func StaticFS(fs api.Interface) api.Interface {
	return fs.NameSpace("static")
}

func LocaleFS(fs api.Interface) api.Interface {
	return fs.NameSpace("locale")
}