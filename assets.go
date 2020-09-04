package assets

import "github.com/moisespsena-go/assetfs/assetfsapi"

const (
	NS_LOCALE    = "@locale"
	NS_PRIVATE   = "@private"
	NS_STATIC    = "static"
	NS_TEMPLATES = "templates"
)

func LocaleFS(fs assetfsapi.Interface) assetfsapi.Interface {
	return fs.NameSpace(NS_LOCALE)
}

func PrivateFS(fs assetfsapi.Interface) assetfsapi.Interface {
	return fs.NameSpace(NS_PRIVATE)
}

func StaticFS(fs assetfsapi.Interface) assetfsapi.Interface {
	return fs.NameSpace(NS_STATIC)
}

func TemplateFS(fs assetfsapi.Interface) assetfsapi.Interface {
	return fs.NameSpace(NS_TEMPLATES)
}
