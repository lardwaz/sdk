# Lardwaz Go SDK

The Go implementation of Lardwaz SDK.

---
- [Lardwaz Go SDK](#lardwaz-go-sdk)
  - [Getting Started](#getting-started)
  - [Example Plugin](#example-plugin)
  - [Plugin Symbols](#plugin-symbols)
      - [GetPluginInfo](#getplugininfo)
      - [OnLoaded](#onloaded)
      - [OnInstall](#oninstall)
      - [OnUninstall](#onuninstall)
      - [GetEntities](#getentities)
      - [GetAuthProvider](#getauthprovider)
      - [GetCommandHook](#getcommandhook)
      - [GetHTTPHooks](#gethttphooks)
      - [GetSchemaHooks](#getschemahooks)
      - [GetResolverHooks](#getresolverhooks)
      - [GetGlobalResolverHooks](#getglobalresolverhooks)

---

## Getting Started
To install the library, run:

```sh
go get go.lsl.digital/lardwaz/sdk
```

## Example Plugin
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
		),
	})

	return entities
}
```

---

## Plugin Symbols

An overview of the various plugin symbols exposed by the sdk.

#### GetPluginInfo

Meta information about the plugin.
 - Name
 - Machine name
 - Description
 - Logo
 - Author
 - Maintainers
 - Version
 - Dependencies

```go
func GetPluginInfo() info.Info {
	return info.New("Plugin Demo", "plugin-demo", "A demo plugin for development", "chrome_reader_mode", "LSL Digital", nil, "v0.0.1")
}
```

#### OnLoaded

Code to execute when plugin has loaded

```go
func OnLoaded() hook.LoadedHook {
	loadedHook := hook.NewLoadedHook()

	loadedHook.RegisterNewHook(func(appConf config.Config, schematics schematic.Schematics, configPath string) error {
        log.Printf("Configuration from lardwaz: %+v\n", appConf)

        return nil
	})

	return loadedHook
}
```

#### OnInstall

Code to execute when plugin is installed.

```go
func OnInstall() hook.InstallHook {
	installHook := hook.NewInstallHook()

	installHook.RegisterNewHook(func() error {
        log.Println("Plugin has been installed")

		return nil
	})

	return installHook
}
```

#### OnUninstall

Code to execute when plugin is uninstalled.

```go
func OnUninstall() hook.UninstallHook {
	uninstallHook := hook.NewUninstallHook()

	uninstallHook.RegisterNewHook(func() error {
        log.Println("Plugin has been uninstalled")

		return nil
	})

	return uninstallHook
}
```

#### GetEntities

Entities provided by the plugin.

In this example, we return the entity `Tag` with fields `ID` and `Name`. 

```go
type Tag struct {
	ID       string    `json:"id" label:"Actions"`
	Name     string    `json:"name" label:"Name" sortable:"true" filterable:"true" widget_list_type:"any" widget_edit_type:"any"`
}

func GetEntities() []entity.Entity {
	var entities []entity.Entity

	entities = append(entities, struct {
		entity.Entity
		content.Contentable
		resolver.Resolvable
		types.Seedable
	}{
        // The instance of the entity itself is defined here
        entity.New(Tag{}),
        // Used by any backoffice front-end system
        content.New("icon", "show:all", "Tag", "name", "A tag that describes things", "content", 10),
        // Defines what resolvers we need { `single`, `listing`, `create`, `update`, `delete` }
        resolver.New("A tag for the articles", resolver.ALL),
        // Defines how the entity should be seeded
		types.NewSeedable(1,
			types.Field{Name: "ID", Type: seeder.Random(util.UUID)},
			types.Field{Name: "Name", Type: seeder.Random(util.Word)},
		),
	})

	return entities
}
```

#### GetAuthProvider

Defines a valid `auth.Provider` implementation.

```go
func GetAuthProvider() authsdk.ProviderFn {
	return func() authsdk.Provider {
		creds := auth.NewCredentials(cfg.Credentials)

		provider = jwt.New(creds)

		return provider
	}
}
```

#### GetCommandHook

Registers new cobra commands.

```go
func GetCommandHook() hook.CommandHook {
	hooks := hook.NewCommandHook()

	demoCmd := &cobra.Command{
		Use:   "demo",
		Short: "A demo command from plugin-demo",
		Long:  "This command was created in plugin-demo and does nothing but say hello",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Hello from plugin-demo!")
		},
	}

	hooks.RegisterNewHook(func(cmd *cobra.Command) error {
		cmd.AddCommand(demoCmd)
		return nil
	})

	return hooks
}
```

#### GetHTTPHooks

Exposes the lardwaz `mux.Router` for obvious routing usages.

```go
func GetHTTPHooks() hook.HTTPHooks {
	hooks := hook.NewHTTPHooks()

	hooks.RegisterNewHook(func(router *mux.Router) error {
		router.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello from plugin-demo"))
		})
		return nil
	})

	return hooks
}
```

#### GetSchemaHooks

Allows for registering of new GraphQL `Queries` and `Mutations`.

```go
func GetSchemaHooks() hook.SchemaHooks {
	hooks := hook.NewSchemaHooks()

	hooks.RegisterNewQueryHook(func(queries graphql.Fields) error {
		queries["hello"] = &graphql.Field{
			Type:        graphql.String,
			Description: "A demo query field from plugin-demo",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello from plugin-demo", nil
			},
		}
		return nil
	})

	hooks.RegisterNewMutationHook(func(mutations graphql.Fields) error {
		mutations["hello_name"] = &graphql.Field{
			Type:        graphql.String,
			Description: "A demo mutation field from plugin-demo",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := "unnamed"

				// Checks args
				if v, ok := p.Args["name"]; ok {
					if vs, ok := v.(string); ok {
						name = vs
					}
				}

				return "Hello " + name + " from plugin-demo", nil
			},
		}
		return nil
	})

	return hooks
}
```

#### GetResolverHooks

Allows for registering entity-level `pre` and `post` GraphQL `resolver` hooks.

```go
func GetResolverHooks() hook.ResolverHooks {
	hooks := hook.NewResolverHooks()

	hooks.RegisterNewPostListingHook("demo", func(entityname string, p *graphql.ResolveParams, vals []interface{}) ([]interface{}, error) {
		// Override vals just for testing purposes
		result := []Demo{
			Demo{ID: uuid.NewV4().String(), Output: "Resolver Hook Test Demo-Plugin #0"},
			Demo{ID: uuid.NewV4().String(), Output: "Resolver Hook Test Demo-Plugin #1"},
			Demo{ID: uuid.NewV4().String(), Output: "Resolver Hook Test Demo-Plugin #2"},
		}

		var res []interface{}
		for _, r := range result {
			res = append(res, r)
		}

		return res, nil
	})

	return hooks
}
```

#### GetGlobalResolverHooks

Allows for registering resolver-level `pre` and `post` GraphQL `resolver` hooks.

```go
func GetGlobalResolverHooks() hook.GlobalResolverHooks {
	hooks := hook.NewGlobalResolverHooks()

	hooks.RegisterNewGlobalPreSingleHook(func(entityname string, p *graphql.ResolveParams) error {
        log.Printf("The single resolver for '%s' has been called\n", entityname)

        return nil
    })

	return hooks
}
```