package assets

import "github.com/moisespsena/go-assetfs/assetfsapi"

func TemplateFS(fs assetfsapi.Interface) assetfsapi.Interface {
	return fs.NameSpace("templates")
}

func StaticFS(fs assetfsapi.Interface) assetfsapi.Interface {
	return fs.NameSpace("static")
}

func LocaleFS(fs assetfsapi.Interface) assetfsapi.Interface {
	return fs.NameSpace("locale")
}