# Lardwaz Go SDK

The Go implementation of Lardwaz SDK.

------

### Getting Started
To install the library, run:

```
go get go.lsl.digital/lardwaz/sdk
```

The following is a very simple plugin implemented using Lardwaz Go SDK.

```go
package main

import (
	"time"

	"go.lsl.digital/lardwaz/sdk/content"
	"go.lsl.digital/lardwaz/sdk/entity"
	"go.lsl.digital/lardwaz/sdk/info"
	"go.lsl.digital/lardwaz/sdk/resolver"
	"go.lsl.digital/lardwaz/sdk/schematic"
	"go.lsl.digital/lardwaz/sdk/seeder"
	"go.lsl.digital/lardwaz/sdk/seeder/types"
	"go.lsl.digital/lardwaz/sdk/seeder/util"
)

// since package is main, needs an entry point
func main() {}

// GetPluginInfo identifies this plugin in Lardwaz
func GetPluginInfo() info.Info {
	return info.New("Plugin Demo", "plugin-demo", "A demo plugin for development", "chrome_reader_mode", "LSL Digital", nil, "v0.0.1")
}

//Article represents an article on the website
type Article struct {
	ID        string    `json:"id" label:"Actions"`
	Title     string    `json:"title" label:"Title" sortable:"true" filterable:"true" widget_list_type:"any" widget_edit_type:"any"`
	Body      string    `json:"body" label:"Body" widget_list_hide:"true" widget_edit_type:"builder_content"`
	Content   string    `json:"content" label:"Content" widget_list_hide:"true"  widget_edit_type:"builder_content"`
	Extra     string    `json:"extra" label:"Extra" widget_list_hide:"true"  widget_edit_type:"builder_simple"`
	Tags      []Tag     `json:"tags" label:"Tags"`
	CreatedAt time.Time `json:"created_at" label:"Created At" sortable:"true" filterable:"true" widget_list_type:"date" widget_edit_type:"date-readonly"`
	UpdatedAt time.Time `json:"updated_at" label:"Updated At" sortable:"true" widget_list_type:"date"  widget_edit_type:"date-readonly"`
	Status    string    `json:"status" label:"Status" widget_list_type:"status"  widget_edit_type:"status" widget_edit_options:"Draft, Saved, Published, Cancelled"`
}
//Tag represents a means of categorization of articles
type Tag struct {
	ID       string    `json:"id" label:"Actions"`
	Name     string    `json:"name" label:"Name" sortable:"true" filterable:"true" widget_list_type:"any" widget_edit_type:"any"`
	Articles []Article `json:"articles" widget_list_hide:"true" label:"Articles" widget_list_type:"any" widget_edit_type:"any"`
}

//GetEntities returns entities defined by the plugin
func GetEntities() []entity.Entity {
	var entities []entity.Entity

	entities = append(entities, struct {
		entity.Entity
		content.Contentable
		resolver.Resolvable
		types.Seedable
	}{
		entity.New(Article{}),
		content.New("icon", "show:all", "Article", "title", "An article on the website", "content", 10),
		resolver.New("An article on the website", resolver.ALL),
		types.NewSeedable(2,
			types.Field{Name: "ID", Type: seeder.Random(util.UUID)},
			types.Field{Name: "Title", Type: seeder.Random(util.Title)},
			types.Field{Name: "Body", Type: types.NewParagraphsN(15)},
			types.Field{Name: "Author", Type: seeder.Random(util.UUID)},
			types.Field{Name: "Tags", Type: seeder.Random(util.UUID)},
			types.Field{Name: "CreatedAt", Type: types.NewDatetime(true, 0, 0, 0, 0, "", "")},
			types.Field{Name: "UpdatedAt", Type: types.NewDatetime(true, 0, 0, 0, 0, "", "")},
		),
	})

	entities = append(entities, struct {
		entity.Entity
		content.Contentable
		resolver.Resolvable
		types.Seedable
	}{
		entity.New(Tag{}),
		content.New("icon", "show:all", "Tag", "name", "A tag that describes things", "content", 10),
		resolver.New("A tag for the articles", resolver.ALL),
		types.NewSeedable(1,
			types.Field{Name: "ID", Type: seeder.Random(util.UUID)},
			types.Field{Name: "Name", Type: seeder.Random(util.Word)},
			types.Field{Name: "Articles", Type: seeder.Random(util.UUID)},
		),
	})

	return entities
}
```

### Plugin Symbols

An overview of the various plugin symbols exposed by the sdk.

#### GetEntities

Returns meta information about the plugin
 - Name
 - Machine name
 - Description
 - Logo
 - Author
 - Maintainers
 - Version
 - Dependencies

##### Usage

```go
func GetPluginInfo() info.Info {
	return info.New("Plugin Demo", "plugin-demo", "A demo plugin for development", "chrome_reader_mode", "LSL Digital", nil, "v0.0.1")
}
```