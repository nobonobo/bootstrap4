# Bootstrap4 components for Vecty

## Define

```go
type ComponentName struct {
    vecty.Core
    Bold bool `vecty:"prop"`
    ...
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

func (c *ComponentName) Render() vecty.ComponentOrHTML {
    return vecty.Tag(elementName,
		vecty.Markup(
            vecty.ClassMap{
                "bold": c.Bold,
            }
            ...
        ),
        c.Markup,
        c.Children,
    )
}
```

## Usage

### Basic

```go
    &ComponentName{
        Bold: true,
        Markup:vecty.Markup(vecty.Attribute("id","name")),
        Children: vecty.Text("name"),
    }
```

### Multiple Child

```go
    &ComponentName{
        Bold: true,
        Markup:vecty.Markup(vecty.Attribute("id","name")),
        Children: vecty.List(
            vecty.Text("bra..bra.."),
            elem.Span(vecty.Text("memo")),
        ),
    }
```

### Nested Component

```go
    &ComponentName1{
        Bold: true,
        Markup:vecty.Markup(vecty.Attribute("id","name")),
        Children: &ComponentName2{
            ...
        },
    }
```

### Event Handling

logging and original behavier
```go
&ComponentName{
    Href:    "#/",
    Markup: vecty.Markup(event.Click(func(event *vecty.Event) {
        log.Println("click: link")
    })),
    Children: vecty.Text("link"),
},
```

logging only
```go
&ComponentName{
    Href:    "#/",
    Markup: vecty.Markup(event.Click(func(event *vecty.Event) {
        log.Println("click: link")
    }).PreventDefault()),
    Children: vecty.Text("link"),
},
```
