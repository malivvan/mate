package editor

import (
	"embed"
	"io/fs"
	"strings"
)

type AssetList []*Asset

func (a AssetList) List() []string {
	names := make([]string, len(a))
	for i, asset := range a {
		names[i] = asset.Name
	}
	return names
}

func (a AssetList) Get(name string) *Asset {
	for _, asset := range a {
		if asset.Name == name {
			return asset
		}
	}
	return nil
}

type Asset struct {
	Name string
	Data []byte
}

//go:embed assets/colorschemes/*.micro
var colorschemeFS embed.FS

//go:embed assets/syntax/*.yaml
var syntaxFS embed.FS

var SyntaxAssets = func() AssetList {
	var assets []*Asset
	err := fs.WalkDir(syntaxFS, "assets/syntax", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		data, err := fs.ReadFile(syntaxFS, path)
		if err != nil {
			return err
		}
		assets = append(assets, &Asset{
			Name: strings.TrimSuffix(strings.TrimPrefix(path, "assets/syntax/"), ".yaml"),
			Data: data,
		})
		return nil
	})
	if err != nil {
		panic(err)
	}
	return assets
}()

var ColorschemeAssets = func() AssetList {
	var assets []*Asset
	err := fs.WalkDir(colorschemeFS, "assets/colorschemes", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		data, err := fs.ReadFile(colorschemeFS, path)
		if err != nil {
			return err
		}
		assets = append(assets, &Asset{
			Name: strings.TrimSuffix(strings.TrimPrefix(path, "assets/colorschemes/"), ".micro"),
			Data: data,
		})
		return nil
	})
	if err != nil {
		panic(err)
	}
	return assets
}
